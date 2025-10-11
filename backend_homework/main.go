// 主程序入口文件
// 功能：演示 go-ethereum 与 Sepolia 测试链的交互操作
package main

import (
	"fmt"      // 格式化输入输出
	"log"      // 日志记录和错误处理
	"math/big" // 大数运算库，处理以太坊金额

	"github.com/huluobo/go-eth-learning/config" // 配置管理模块
	"github.com/huluobo/go-eth-learning/tx"     // 交易管理模块
)

func main() {
	// ========== 初始化阶段 ==========
	// 从 .env 文件加载配置信息（私钥、RPC地址、链ID等）
	cfg := config.Load()

	// 创建交易管理器实例
	// 参数1: RPC URL - 用于连接以太坊节点
	// 参数2: 私钥 - 用于签名交易
	// 参数3: 链ID - Sepolia 测试网是 11155111
	txManager, err := tx.NewTransactionManager(cfg.RPCUrl, cfg.PrivateKey, cfg.ChainID)
	if err != nil {
		// log.Fatal 会打印错误并调用 os.Exit(1) 退出程序
		log.Fatal("创建交易管理器失败:", err)
	}
	// defer 确保程序退出时关闭与以太坊节点的连接
	defer txManager.Close()

	// ========== 功能1: 查询账户余额 ==========
	fmt.Println("=== 查询余额 ===")
	// 定义要查询余额的钱包地址列表
	balances := []string{
		"0x5D4a329B262AC7a1d9Ae0f4C54171dF61E2c0b69", // WALLET1 - 第一个测试钱包
		"0xF4076C4a38971D71812B298A6aA9213C5425fa51", // WALLET2 - 第二个测试钱包
	}

	// 遍历地址列表，逐个查询余额
	for _, addr := range balances {
		// 调用 GetBalance 方法查询指定地址的余额
		// 返回值是 Wei 单位的大整数
		balance, err := txManager.GetBalance(addr)
		if err != nil {
			// 查询失败时打印错误信息但不退出，继续查询下一个地址
			fmt.Printf("查询 %s 余额失败: %v\n", addr, err)
			continue // 跳过当前循环，继续下一个
		}

		// 单位转换：从 Wei 转换为 ETH
		// 1 ETH = 10^18 Wei
		etherBalance := new(big.Float).SetInt(balance)        // 先将 big.Int 转为 big.Float
		etherBalance.Quo(etherBalance, big.NewFloat(1e18))    // 除以 10^18

		// 格式化输出查询结果
		fmt.Printf("地址 %s:\n", addr)
		fmt.Printf("  余额(Wei): %s\n", balance.String())      // Wei 是最小单位
		fmt.Printf("  余额(ETH): %s\n\n", etherBalance.Text('f', 6)) // 保留6位小数
	}

	// ========== 功能2: 演示完整的交易生命周期 ==========
	fmt.Println("\n是否演示交易生命周期？(需要消耗 gas) [y/n]: ")
	var response string      // 用于存储用户输入
	fmt.Scanln(&response)    // 从标准输入读取用户响应

	// 用户确认后执行交易演示
	if response == "y" {
		// 设置转账金额：0.0001 ETH
		// 必须使用 Wei 单位，所以是 0.0001 * 10^18 = 100000000000000
		amount := big.NewInt(100000000000000) // 0.0001 ETH in wei

		// 设置接收地址（这里使用 WALLET2）
		toAddress := "0xF4076C4a38971D71812B298A6aA9213C5425fa51" // WALLET2

		// 执行交易生命周期演示
		// 包含：构建交易、签名、广播、等待打包、获取回执等完整流程
		err := txManager.DemoTransactionLifecycle(toAddress, amount)
		if err != nil {
			// 交易失败时记录错误但不退出程序
			log.Printf("交易演示失败: %v\n", err)
		}
	}

	// ========== 功能3: 智能合约交互（SimpleStorage） ==========
	fmt.Println("\n是否演示智能合约交互？[y/n]: ")
	var contractChoice string
	fmt.Scanln(&contractChoice)

	if contractChoice == "y" {
		fmt.Println("\n=== SimpleStorage 合约交互演示 ===")
		fmt.Println("提示：请使用以下 Makefile 命令进行合约操作：")
		fmt.Println("  1. 生成 Go 绑定:  make abigen")
		fmt.Println("  2. 部署合约:      make contract-deploy")
		fmt.Println("  3. 交互合约:      make contract-interact ADDR=<合约地址>")
		fmt.Println("\n完整示例流程请查看 examples/contract/ 目录")
	}

	// ========== 功能4: 订阅新区块（需要 WebSocket） ==========
	fmt.Println("\n选择订阅模式:")
	fmt.Println("1. 不订阅新区块")
	fmt.Println("2. 常规订阅新区块")
	fmt.Println("3. 订阅新区块并监控大额转账（>0.5 ETH）")
	fmt.Print("请输入选项 [1/2/3]: ")

	var choice string
	fmt.Scanln(&choice)    // 读取用户选择

	switch choice {
	case "1":
		// 用户选择不订阅
		fmt.Println("跳过订阅功能")

	case "2":
		// 常规订阅：只显示新区块信息
		if cfg.WSUrl == "" {
			log.Printf("WebSocket URL 未配置，无法订阅。请在 .env 中设置 SEPOLIA_WS_URL\n")
		} else {
			fmt.Println("\n开始常规订阅...")
			err := txManager.SubscribeNewBlocks(cfg.WSUrl)
			if err != nil {
				log.Printf("订阅失败: %v\n", err)
			}
		}

	case "3":
		// 高级订阅：监控大额转账
		if cfg.WSUrl == "" {
			log.Printf("WebSocket URL 未配置，无法订阅。请在 .env 中设置 SEPOLIA_WS_URL\n")
		} else {
			// 设置监控阈值：0.5 ETH
			threshold := new(big.Int)
			threshold.SetString("500000000000000000", 10) // 0.5 ETH in Wei

			fmt.Println("\n开始监控大额转账（>0.5 ETH）...")
			err := txManager.SubscribeWithMonitor(cfg.WSUrl, threshold)
			if err != nil {
				log.Printf("订阅失败: %v\n", err)
			}
		}

	default:
		fmt.Println("无效选项，跳过订阅功能")
	}
}