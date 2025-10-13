package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/gagliardetto/solana-go"
	"github.com/huluobo/solana-homework/config"
	"github.com/huluobo/solana-homework/pkg/chain"
	"github.com/huluobo/solana-homework/pkg/events"
)

func main() {
	ctx := context.Background()
	cfg := config.New()

	fmt.Println("=== Solana 事件监听演示 ===")
	fmt.Println("这个演示会:")
	fmt.Println("1. 发送一笔 SOL 转账")
	fmt.Println("2. 通过 WebSocket 实时监听交易确认\n")

	// 提示用户输入私钥
	fmt.Println("需要私钥才能发送交易。")
	fmt.Println("如果没有私钥，程序将只演示如何使用事件监听器。\n")

	// 演示事件监听器的使用方式
	fmt.Println("【事件监听器使用示例】")
	fmt.Println(`
// 1. 创建监听器
listener, err := events.NewListener(cfg.WSURL)
if err != nil {
    log.Fatal(err)
}
defer listener.Close()

// 2. 订阅特定交易签名
ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
defer cancel()

err = listener.SubscribeSignature(ctx, transactionSignature)
if err != nil {
    log.Printf("Transaction failed or timeout: %v", err)
}

// 3. 监听账户变化
err = listener.SubscribeAccount(ctx, accountPublicKey)
if err != nil {
    log.Printf("Subscription ended: %v", err)
}
`)

	// 如果用户提供了私钥，可以执行真实的转账+监听演示
	fmt.Println("\n【实际测试】")
	fmt.Println("要测试实际的事件监听，请:")
	fmt.Println("1. 准备一个有余额的 Devnet 钱包")
	fmt.Println("2. 在代码中硬编码私钥（仅用于测试）")
	fmt.Println("3. 取消注释下面的代码并重新编译\n")

	// 下面是实际测试代码的模板（默认注释掉）
	_ = demonstrateEventListening
}

// demonstrateEventListening shows a complete example of sending a transaction
// and listening for its confirmation via WebSocket
func demonstrateEventListening() {
	ctx := context.Background()
	cfg := config.New()

	// Step 1: Create RPC client
	client, err := chain.NewClient(cfg.RPCURL)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}
	defer client.Close()

	// Step 2: Load private key (REPLACE WITH YOUR KEY)
	privateKeyStr := "YOUR_PRIVATE_KEY_BASE58_HERE"
	privateKey, err := solana.PrivateKeyFromBase58(privateKeyStr)
	if err != nil {
		log.Fatalf("Invalid private key: %v", err)
	}

	// Step 3: Set target address and amount
	toAddr := solana.MustPublicKeyFromBase58("11111111111111111111111111111111")
	amount := uint64(1000) // 0.000001 SOL

	// Step 4: Create event listener BEFORE sending transaction
	listener, err := events.NewListener(cfg.WSURL)
	if err != nil {
		log.Fatalf("Failed to create listener: %v", err)
	}
	defer listener.Close()

	// Step 5: Send transaction
	fmt.Println("Sending transaction...")
	sig, err := client.TransferSOL(ctx, privateKey, toAddr, amount)
	if err != nil {
		log.Fatalf("Transfer failed: %v", err)
	}

	fmt.Printf("Transaction sent: %s\n", sig)

	// Step 6: Subscribe to transaction confirmation
	confirmCtx, cancel := context.WithTimeout(ctx, 60*time.Second)
	defer cancel()

	err = listener.SubscribeSignature(confirmCtx, sig)
	if err != nil {
		log.Fatalf("Failed to confirm transaction: %v", err)
	}

	fmt.Println("Transaction confirmed successfully!")
}
