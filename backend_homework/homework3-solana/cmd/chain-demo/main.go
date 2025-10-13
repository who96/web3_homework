package main

import (
	"context"
	"fmt"
	"log"

	"github.com/gagliardetto/solana-go"
	"github.com/huluobo/solana-homework/config"
	"github.com/huluobo/solana-homework/pkg/chain"
)

func main() {
	ctx := context.Background()

	// Initialize config
	cfg := config.New()
	fmt.Printf("=== Solana Devnet 基础链交互演示 ===\n")
	fmt.Printf("RPC Endpoint: %s\n\n", cfg.RPCURL)

	// Create RPC client
	client, err := chain.NewClient(cfg.RPCURL)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}
	defer client.Close()

	// Demo 1: Get latest blockhash
	fmt.Println("【1. 查询最新区块哈希】")
	blockhash, err := client.GetLatestBlockhash(ctx)
	if err != nil {
		log.Fatalf("Failed to get blockhash: %v", err)
	}
	fmt.Printf("Latest Blockhash: %s\n", blockhash)
	fmt.Printf("注意: Solana 的 blockhash 有效期只有 60-90 秒\n\n")

	// Demo 2: Query account balances
	fmt.Println("【2. 查询账户余额】")

	// Test addresses (Solana uses base58, not hex)
	// These are well-known Solana program addresses
	testAddresses := []string{
		"11111111111111111111111111111111",                // System Program
		"TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA",   // SPL Token Program
		"ATokenGPvbdGVxr1b2hvZbsiqW5xWH25efTNsLJA8knL", // Associated Token Program
	}

	for i, addrStr := range testAddresses {
		addr, err := solana.PublicKeyFromBase58(addrStr)
		if err != nil {
			fmt.Printf("无效地址 %s: %v\n", addrStr, err)
			continue
		}

		balance, err := client.GetBalance(ctx, addr)
		if err != nil {
			fmt.Printf("查询地址 %s 失败: %v\n", addrStr, err)
			continue
		}

		// Convert lamports to SOL (1 SOL = 1,000,000,000 lamports)
		solBalance := float64(balance) / 1e9
		fmt.Printf("地址 %d: %s\n", i+1, addrStr)
		fmt.Printf("  余额(lamports): %d\n", balance)
		fmt.Printf("  余额(SOL): %.9f\n\n", solBalance)
	}

	// Demo 3: Transfer SOL (commented out - needs private key)
	fmt.Println("【3. SOL 转账演示】")
	fmt.Println("转账功能需要私钥。使用示例:")
	fmt.Println(`
  // 从 base58 字符串加载私钥
  privateKey, err := solana.PrivateKeyFromBase58("your-private-key-here")
  if err != nil {
      log.Fatal(err)
  }

  // 目标地址
  toAddr := solana.MustPublicKeyFromBase58("target-address")

  // 转账金额 (0.01 SOL = 10,000,000 lamports)
  amount := uint64(10_000_000)

  // 执行转账
  sig, err := client.TransferSOL(ctx, privateKey, toAddr, amount)
  if err != nil {
      log.Fatalf("Transfer failed: %v", err)
  }

  fmt.Printf("Transaction signature: %s\n", sig)
  fmt.Printf("View on explorer: https://explorer.solana.com/tx/%s?cluster=devnet\n", sig)
`)

	fmt.Println("\n提示: 在 Devnet 测试前,先用 solana-keygen 生成密钥,然后用 solana airdrop 领取测试币")
	fmt.Println("命令: solana airdrop 1 <YOUR_ADDRESS> --url devnet")
}
