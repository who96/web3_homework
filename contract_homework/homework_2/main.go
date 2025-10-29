package main

import (
	"context"
	"crypto/ecdsa"
	"encoding/json"
	"fmt"
	"log"
	"math/big"
	"os"
	"os/signal"
	"strings"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"

	fukuastake "github.com/huluobo/fukua-stake-client/bindings"
)

const (
	stakeAddressStr = "0xFF16fD29A0138E432A49ba7A68dE689c41D43239"
	rewardTokenStr  = "0x8F18b27F3d91b258a3a9242D2Bd4D9C219EEcE1E"
	waitBlocks      = 10
)

type poolMeta struct {
	token      common.Address
	lockBlocks *big.Int
}

var poolCache = struct {
	sync.RWMutex
	m map[uint64]poolMeta
}{m: make(map[uint64]poolMeta)}

func main() {
	rpcURL := mustEnv("SEPOLIA_RPC_URL")
	wsURL := os.Getenv("SEPOLIA_WS_URL")
	privKeyHex := mustEnv("PRIVATE_KEY")

	client, err := ethclient.Dial(rpcURL)
	if err != nil {
		log.Fatalf("连接 RPC 失败: %v", err)
	}
	defer client.Close()

	var logClient *ethclient.Client
	if wsURL != "" {
		if logClient, err = ethclient.Dial(wsURL); err != nil {
			log.Fatalf("连接 WebSocket 失败: %v", err)
		}
	} else {
		log.Fatal("SEPOLIA_WS_URL 未设置，事件监听需要 WebSocket 端点")
	}
	defer logClient.Close()

	privateKey, err := crypto.HexToECDSA(strings.TrimPrefix(privKeyHex, "0x"))
	if err != nil {
		log.Fatalf("解析私钥失败: %v", err)
	}

	chainID, err := client.ChainID(context.Background())
	if err != nil {
		log.Fatalf("获取 chain id 失败: %v", err)
	}

	stakeAddr := common.HexToAddress(stakeAddressStr)
	rewardAddr := common.HexToAddress(rewardTokenStr)

	contract, err := fukuastake.NewFukuastake(stakeAddr, client)
	if err != nil {
		log.Fatalf("绑定合约失败: %v", err)
	}
	filterer, _ := fukuastake.NewFukuastakeFilterer(stakeAddr, client)
	caller, _ := fukuastake.NewFukuastakeCaller(stakeAddr, client)

	parsedABI, err := fukuastake.FukuastakeMetaData.GetAbi()
	if err != nil {
		log.Fatalf("解析 ABI 失败: %v", err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go func() {
		query := ethereum.FilterQuery{Addresses: []common.Address{stakeAddr}}
	outer:
		for {
			logsCh := make(chan types.Log)
			sub, err := logClient.SubscribeFilterLogs(ctx, query, logsCh)
			if err != nil {
				log.Printf("订阅事件失败: %v", err)
				time.Sleep(5 * time.Second)
				continue
			}
			for {
				select {
				case err := <-sub.Err():
					log.Printf("监听错误: %v，重试订阅", err)
					sub.Unsubscribe()
					time.Sleep(5 * time.Second)
					continue outer
				case vLog := <-logsCh:
					handleEvent(filterer, caller, parsedABI, rewardAddr, vLog)
				}
			}
		}
	}()

	// 发送质押 1 ETH
	log.Println("发送质押交易...")
	stakeAuth := newTransactor(privateKey, chainID)
	stakeAuth.Value = big.NewInt(0).Mul(big.NewInt(1e18), big.NewInt(1))
	stakeTx, err := contract.DepositEth(stakeAuth)
	if err != nil {
		log.Fatalf("质押失败: %v", err)
	}
	log.Printf("质押 tx hash: %s", stakeTx.Hash().Hex())
	waitForReceipt(client, stakeTx.Hash())

	// 等待锁定期区块
	waitForBlocks(client, stakeTx.Hash(), waitBlocks)

	// 调用领取奖励
	log.Println("领取奖励...")
	claimAuth := newTransactor(privateKey, chainID)
	claimTx, err := contract.Claim(claimAuth, big.NewInt(0))
	if err != nil {
		log.Fatalf("领取失败: %v", err)
	}
	log.Printf("领取 tx hash: %s", claimTx.Hash().Hex())
	waitForReceipt(client, claimTx.Hash())

	log.Println("脚本执行完毕，等待事件日志...")
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt)
	select {
	case <-time.After(30 * time.Second):
	case <-sig:
	}
}

func newTransactor(privateKey *ecdsa.PrivateKey, chainID *big.Int) *bind.TransactOpts {
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		log.Fatalf("创建交易对象失败: %v", err)
	}
	auth.Context = context.Background()
	return auth
}

func waitForReceipt(client *ethclient.Client, hash common.Hash) {
	ctx := context.Background()
	for {
		receipt, err := client.TransactionReceipt(ctx, hash)
		if err == nil {
			log.Printf("交易 %s 已确认 (区块 %s)", hash.Hex(), receipt.BlockNumber)
			return
		}
		time.Sleep(4 * time.Second)
	}
}

func waitForBlocks(client *ethclient.Client, hash common.Hash, step int64) {
	ctx := context.Background()
	receipt, err := client.TransactionReceipt(ctx, hash)
	if err != nil {
		log.Fatalf("获取交易回执失败: %v", err)
	}
	target := new(big.Int).Add(receipt.BlockNumber, big.NewInt(step))
	for {
		header, err := client.HeaderByNumber(ctx, nil)
		if err != nil {
			log.Fatalf("读取区块头失败: %v", err)
		}
		if header.Number.Cmp(target) >= 0 {
			log.Printf("当前区块 %s 已达到目标 %s", header.Number, target)
			return
		}
		time.Sleep(12 * time.Second)
	}
}

func handleEvent(filterer *fukuastake.FukuastakeFilterer, caller *fukuastake.FukuastakeCaller, parsedABI *abi.ABI, rewardToken common.Address, evt types.Log) {
	switch evt.Topics[0] {
	case parsedABI.Events["Deposit"].ID:
		ev, err := filterer.ParseDeposit(evt)
		if err != nil {
			log.Printf("解析 Deposit 事件失败: %v", err)
			return
		}
		meta := poolMetaInfo(caller, ev.PoolId)
		info := map[string]any{
			"event":  "Deposit",
			"tx":     evt.TxHash.Hex(),
			"user":   ev.User.Hex(),
			"pool":   ev.PoolId.String(),
			"amount": humanAmount(ev.Amount),
			"token":  tokenLabel(meta.token),
			"flow":   fmt.Sprintf("%s -> %s", ev.User.Hex(), evt.Address.Hex()),
		}
		printJSON(info)
	case parsedABI.Events["Claim"].ID:
		ev, err := filterer.ParseClaim(evt)
		if err != nil {
			log.Printf("解析 Claim 事件失败: %v", err)
			return
		}
		info := map[string]any{
			"event":  "Claim",
			"tx":     evt.TxHash.Hex(),
			"user":   ev.User.Hex(),
			"pool":   ev.PoolId.String(),
			"amount": humanAmount(ev.FukuaReward),
			"token":  tokenLabel(rewardToken),
			"flow":   fmt.Sprintf("%s -> %s", evt.Address.Hex(), ev.User.Hex()),
		}
		printJSON(info)
	case parsedABI.Events["RequestUnstake"].ID:
		ev, err := filterer.ParseRequestUnstake(evt)
		if err != nil {
			log.Printf("解析 RequestUnstake 事件失败: %v", err)
			return
		}
		meta := poolMetaInfo(caller, ev.PoolId)
		unlock := new(big.Int).SetUint64(evt.BlockNumber)
		if meta.lockBlocks != nil {
			unlock.Add(unlock, meta.lockBlocks)
		}
		info := map[string]any{
			"event":  "RequestUnstake",
			"tx":     evt.TxHash.Hex(),
			"user":   ev.User.Hex(),
			"pool":   ev.PoolId.String(),
			"amount": humanAmount(ev.Amount),
			"token":  tokenLabel(meta.token),
			"unlock": unlock.String(),
			"flow":   fmt.Sprintf("%s -> %s (锁定队列)", ev.User.Hex(), evt.Address.Hex()),
		}
		printJSON(info)
	case parsedABI.Events["Withdraw"].ID:
		ev, err := filterer.ParseWithdraw(evt)
		if err != nil {
			log.Printf("解析 Withdraw 事件失败: %v", err)
			return
		}
		meta := poolMetaInfo(caller, ev.PoolId)
		info := map[string]any{
			"event":  "Withdraw",
			"tx":     evt.TxHash.Hex(),
			"user":   ev.User.Hex(),
			"pool":   ev.PoolId.String(),
			"amount": humanAmount(ev.Amount),
			"token":  tokenLabel(meta.token),
			"flow":   fmt.Sprintf("%s -> %s", evt.Address.Hex(), ev.User.Hex()),
		}
		printJSON(info)
	}
}

func poolMetaInfo(caller *fukuastake.FukuastakeCaller, pid *big.Int) poolMeta {
	id := pid.Uint64()
	poolCache.RLock()
	if meta, ok := poolCache.m[id]; ok {
		poolCache.RUnlock()
		return meta
	}
	poolCache.RUnlock()

	pool, err := caller.Pool(&bind.CallOpts{}, pid)
	if err != nil {
		log.Printf("读取池信息失败: %v", err)
		return poolMeta{}
	}
	meta := poolMeta{token: pool.StTokenAddress, lockBlocks: new(big.Int).Set(pool.UnstakeLockedBlocks)}
	poolCache.Lock()
	poolCache.m[id] = meta
	poolCache.Unlock()
	return meta
}

func humanAmount(v *big.Int) string {
	f := new(big.Float).SetInt(v)
	ethVal := new(big.Float).Quo(f, big.NewFloat(1e18))
	str := ethVal.Text('f', 6)
	return str
}

func tokenLabel(addr common.Address) string {
	if addr == (common.Address{}) {
		return "ETH"
	}
	if strings.EqualFold(addr.Hex(), rewardTokenStr) {
		return "FUKUA"
	}
	return fmt.Sprintf("Token(%s)", addr.Hex())
}

func printJSON(v any) {
	b, _ := json.MarshalIndent(v, "", "  ")
	log.Println(string(b))
}

func mustEnv(key string) string {
	v := os.Getenv(key)
	if v == "" {
		log.Fatalf("环境变量 %s 未设置", key)
	}
	return v
}
