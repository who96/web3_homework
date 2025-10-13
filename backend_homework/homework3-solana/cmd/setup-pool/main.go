package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/gagliardetto/solana-go"
	"github.com/huluobo/solana-homework/config"
	"github.com/huluobo/solana-homework/pkg/chain"
	"github.com/huluobo/solana-homework/pkg/token"
)

const (
	TokenA = "H3C1Bkk1YzjcrXP4eJZWPM1cdMrorty9oWRPCutzp1EJ" // a_t mint
	TokenB = "DQhU9TdEL8FW9ALTExNbt3NzR64rj5ZUKqyh3BQbwfay" // b_t mint
)

func main() {
	fmt.Println("==========================================================")
	fmt.Println("   Token Swap Pool 初始化")
	fmt.Println("==========================================================\n")

	ctx := context.Background()
	cfg := config.New()

	// Load keypair (will be pool authority)
	keypairPath := chain.GetDefaultKeypairPath()
	fmt.Printf("📁 加载密钥: %s\n", keypairPath)

	privateKey, err := chain.LoadKeypairFromFile(keypairPath)
	if err != nil {
		log.Fatalf("❌ 加载密钥失败: %v", err)
	}

	poolAuthority := privateKey.PublicKey()
	fmt.Printf("✅ Pool Authority: %s\n\n", poolAuthority)

	// Create token client
	tokenClient, err := token.NewClient(cfg.RPCURL)
	if err != nil {
		log.Fatalf("❌ 创建 Token 客户端失败: %v", err)
	}

	mintA := solana.MustPublicKeyFromBase58(TokenA)
	mintB := solana.MustPublicKeyFromBase58(TokenB)

	// Step 1: Create pool's a_t token account
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println("📌 Step 1: 创建 Pool 的 a_t Token 账户")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━\n")

	poolAAccount, err := tokenClient.CreateTokenAccount(ctx, privateKey, mintA, poolAuthority)
	if err != nil {
		log.Fatalf("❌ 创建失败: %v", err)
	}
	fmt.Printf("✅ Pool's a_t 账户: %s\n\n", poolAAccount)

	// Step 2: Create pool's b_t token account
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println("📌 Step 2: 创建 Pool 的 b_t Token 账户")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━\n")

	poolBAccount, err := tokenClient.CreateTokenAccount(ctx, privateKey, mintB, poolAuthority)
	if err != nil {
		log.Fatalf("❌ 创建失败: %v", err)
	}
	fmt.Printf("✅ Pool's b_t 账户: %s\n\n", poolBAccount)

	// Step 3: Get existing token accounts for minting
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println("📌 Step 3: 查询现有 Token 账户")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━\n")

	accounts, err := tokenClient.GetTokenAccountsByOwner(ctx, poolAuthority)
	if err != nil {
		log.Fatalf("❌ 查询失败: %v", err)
	}

	fmt.Printf("找到 %d 个 Token 账户:\n", len(accounts))
	for i, acc := range accounts {
		balance, _ := tokenClient.GetTokenBalance(ctx, acc)
		fmt.Printf("   %d. %s (余额: %.2f)\n", i+1, acc, balance)
	}
	fmt.Println()

	// Step 4: Save configuration
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println("📌 Step 4: 保存配置")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━\n")

	configFile := "swap_pool_config.txt"
	content := fmt.Sprintf(`Token Swap Pool Configuration
Generated: %s

Pool Authority: %s

Token Accounts:
- Pool's a_t Account: %s
- Pool's b_t Account: %s

Token Mints:
- a_t Mint: %s
- b_t Mint: %s

Exchange Rate: 1 a_t = 100 b_t

Next Steps:
1. Mint b_t tokens to pool's b_t account for liquidity:
   spl-token mint %s 10000 %s

2. Create user token accounts:
   spl-token create-account %s
   spl-token create-account %s

3. Mint some a_t to user for testing:
   spl-token mint %s <amount> <user_a_t_account>

4. Run swap demo:
   go run cmd/swap-demo/main.go
`,
		"2025-10-12",
		poolAuthority,
		poolAAccount,
		poolBAccount,
		TokenA,
		TokenB,
		TokenB, poolBAccount,
		TokenA, TokenB,
		TokenA,
	)

	err = os.WriteFile(configFile, []byte(content), 0644)
	if err != nil {
		log.Fatalf("❌ 保存配置失败: %v", err)
	}

	fmt.Printf("✅ 配置已保存到: %s\n\n", configFile)

	// Summary
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println("🎉 Pool 初始化完成!")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━\n")

	fmt.Println("📋 Pool 信息:")
	fmt.Printf("   Authority: %s\n", poolAuthority)
	fmt.Printf("   a_t 账户: %s\n", poolAAccount)
	fmt.Printf("   b_t 账户: %s\n", poolBAccount)
	fmt.Println()

	fmt.Println("📝 下一步操作:")
	fmt.Println("   1. 向 pool 的 b_t 账户充值 10000 b_t")
	fmt.Printf("      spl-token mint %s 10000 %s\n", TokenB, poolBAccount)
	fmt.Println()
	fmt.Println("   2. 创建你的 token 账户并充值 a_t 用于测试")
	fmt.Printf("      spl-token create-account %s\n", TokenA)
	fmt.Printf("      spl-token mint %s 10 <your_a_t_account>\n", TokenA)
	fmt.Println()
}
