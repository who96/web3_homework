// 合约交互示例
// 功能：与已部署的 SimpleStorage 合约交互（increment, get, 事件监听）
package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/huluobo/go-eth-learning/contracts/bindings"
	"github.com/joho/godotenv"
)

func main() {
	// 加载 .env 配置
	if err := godotenv.Load("../../.env"); err != nil {
		log.Fatal("Error loading .env file:", err)
	}

	// 从命令行参数获取合约地址
	if len(os.Args) < 2 {
		log.Fatal("Usage: go run interact.go <contract_address>")
	}
	contractAddressHex := os.Args[1]
	if !common.IsHexAddress(contractAddressHex) {
		log.Fatal("Invalid contract address")
	}
	contractAddress := common.HexToAddress(contractAddressHex)

	// 连接到 Sepolia 测试网
	rpcURL := os.Getenv("SEPOLIA_RPC_URL")
	client, err := ethclient.Dial(rpcURL)
	if err != nil {
		log.Fatal("Failed to connect to Sepolia:", err)
	}
	defer client.Close()

	// 加载私钥
	privateKeyHex := os.Getenv("PRIVATE_KEY")
	if len(privateKeyHex) > 2 && privateKeyHex[:2] == "0x" {
		privateKeyHex = privateKeyHex[2:]
	}
	privateKey, err := crypto.HexToECDSA(privateKeyHex)
	if err != nil {
		log.Fatal("Failed to load private key:", err)
	}

	// 获取公钥和地址
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("Failed to cast public key to ECDSA")
	}
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	// 连接到已部署的合约
	instance, err := bindings.NewSimpleStorage(contractAddress, client)
	if err != nil {
		log.Fatal("Failed to connect to contract:", err)
	}

	fmt.Println("=== SimpleStorage 合约交互演示 ===")
	fmt.Printf("合约地址: %s\n", contractAddress.Hex())
	fmt.Printf("调用者地址: %s\n\n", fromAddress.Hex())

	// ========== 1. 读取当前值 ==========
	fmt.Println("--- 1. 读取当前存储值 ---")
	currentValue, err := instance.Get(&bind.CallOpts{})
	if err != nil {
		log.Fatal("Failed to get value:", err)
	}
	fmt.Printf("当前值: %s\n\n", currentValue.String())

	// ========== 2. 调用 increment() 增加计数器 ==========
	fmt.Println("--- 2. 调用 increment() 增加计数器 ---")

	// 获取交易参数
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal("Failed to get nonce:", err)
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal("Failed to suggest gas price:", err)
	}

	chainID, err := client.ChainID(context.Background())
	if err != nil {
		log.Fatal("Failed to get chain ID:", err)
	}

	// 创建交易签名器
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		log.Fatal("Failed to create transactor:", err)
	}

	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)
	auth.GasLimit = uint64(100000)
	auth.GasPrice = gasPrice

	// 调用 increment()
	tx, err := instance.Increment(auth)
	if err != nil {
		log.Fatal("Failed to call increment:", err)
	}

	fmt.Printf("交易已发送: %s\n", tx.Hash().Hex())
	fmt.Println("等待交易确认...")

	// 等待交易确认
	receipt, err := bind.WaitMined(context.Background(), client, tx)
	if err != nil {
		log.Fatal("Failed to wait for transaction:", err)
	}

	if receipt.Status == 0 {
		log.Fatal("Transaction failed")
	}

	fmt.Printf("✅ 交易已确认 (区块 %d, Gas 使用: %d)\n\n", receipt.BlockNumber.Uint64(), receipt.GasUsed)

	// ========== 3. 读取新值 ==========
	fmt.Println("--- 3. 读取更新后的值 ---")
	newValue, err := instance.Get(&bind.CallOpts{})
	if err != nil {
		log.Fatal("Failed to get new value:", err)
	}
	fmt.Printf("新值: %s\n", newValue.String())
	fmt.Printf("变化: %s -> %s (+%d)\n\n", currentValue.String(), newValue.String(), new(big.Int).Sub(newValue, currentValue).Int64())

	// ========== 4. 解析 ValueChanged 事件 ==========
	fmt.Println("--- 4. 解析 ValueChanged 事件 ---")

	// 过滤事件日志
	blockNum := receipt.BlockNumber.Uint64()
	filterOpts := &bind.FilterOpts{
		Start:   blockNum,
		End:     &blockNum,
		Context: context.Background(),
	}

	// 查询 ValueChanged 事件
	iter, err := instance.FilterValueChanged(filterOpts, nil)
	if err != nil {
		log.Fatal("Failed to filter events:", err)
	}
	defer iter.Close()

	fmt.Println("事件列表:")
	eventCount := 0
	for iter.Next() {
		event := iter.Event
		eventCount++
		fmt.Printf("  事件 #%d:\n", eventCount)
		fmt.Printf("    旧值: %s\n", event.OldValue.String())
		fmt.Printf("    新值: %s\n", event.NewValue.String())
		fmt.Printf("    调用者: %s\n", event.Changer.Hex())
		fmt.Printf("    交易: %s\n", event.Raw.TxHash.Hex())
	}

	if err := iter.Error(); err != nil {
		log.Fatal("Error iterating events:", err)
	}

	if eventCount == 0 {
		fmt.Println("  (未找到事件)")
	}

	fmt.Println("\n✅ 交互演示完成！")
}
