// 合约部署示例
// 功能：将 SimpleStorage 合约部署到 Sepolia 测试网
package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
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

	// 获取 nonce
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal("Failed to get nonce:", err)
	}

	// 获取 gas price
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal("Failed to suggest gas price:", err)
	}

	// 获取 chain ID
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
	auth.Value = big.NewInt(0)     // 不发送ETH
	auth.GasLimit = uint64(300000) // Gas 限制
	auth.GasPrice = gasPrice

	// 部署合约
	// 构造函数参数：_initialValue = 42
	initialValue := big.NewInt(42)

	fmt.Println("=== 部署 SimpleStorage 合约 ===")
	fmt.Printf("从地址: %s\n", fromAddress.Hex())
	fmt.Printf("初始值: %s\n", initialValue.String())
	fmt.Printf("Gas Price: %s wei\n", gasPrice.String())
	fmt.Println("\n正在部署...")

	address, tx, instance, err := bindings.DeploySimpleStorage(auth, client, initialValue)
	if err != nil {
		log.Fatal("Failed to deploy contract:", err)
	}

	fmt.Printf("\n✅ 合约部署成功！\n")
	fmt.Printf("合约地址: %s\n", address.Hex())
	fmt.Printf("交易哈希: %s\n", tx.Hash().Hex())
	fmt.Printf("\n等待交易确认...\n")

	// 等待交易被打包
	receipt, err := bind.WaitMined(context.Background(), client, tx)
	if err != nil {
		log.Fatal("Failed to wait for transaction:", err)
	}

	if receipt.Status == 0 {
		log.Fatal("Transaction failed")
	}

	fmt.Printf("\n✅ 交易已确认！\n")
	fmt.Printf("区块号: %d\n", receipt.BlockNumber.Uint64())
	fmt.Printf("Gas 使用: %d\n", receipt.GasUsed)

	// 读取初始值验证部署
	value, err := instance.Get(&bind.CallOpts{})
	if err != nil {
		log.Fatal("Failed to read value:", err)
	}

	fmt.Printf("\n验证: 合约存储的值 = %s\n", value.String())
	fmt.Printf("\n📝 请保存合约地址用于后续交互: %s\n", address.Hex())
}
