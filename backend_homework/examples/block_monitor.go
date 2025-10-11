// 独立运行: go run examples/block_monitor.go
// 功能: 通过 WebSocket 订阅以太坊新区块，实时监控大额交易
package main

import (
	"context"    // 用于管理请求的生命周期
	"fmt"        // 格式化输出
	"log"        // 日志和错误处理
	"math/big"   // 大数运算（以太坊金额都是大整数）
	"os"         // 读取环境变量
	"time"       // 时间格式化

	"github.com/ethereum/go-ethereum/core/types" // 以太坊核心数据类型（区块、交易等）
	"github.com/ethereum/go-ethereum/ethclient"  // 以太坊客户端连接库
	"github.com/joho/godotenv"                   // 加载 .env 文件
)

func main() {
	// 步骤1: 加载环境配置
	// 从父目录的 .env 文件加载环境变量（因为这个文件在 examples 子目录）
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal("Error loading .env file") // 加载失败直接退出
	}

	// 步骤2: 获取 WebSocket URL
	// WebSocket 是持久连接，支持服务器主动推送数据
	wsURL := os.Getenv("SEPOLIA_WS_URL") // 从环境变量读取 wss:// 地址
	if wsURL == "" {
		log.Fatal("SEPOLIA_WS_URL 未配置") // 必须是 wss:// 而不是 https://
	}

	// 步骤3: 建立 WebSocket 连接
	// ethclient.Dial 会自动识别是 HTTP 还是 WebSocket
	wsClient, err := ethclient.Dial(wsURL)
	if err != nil {
		log.Fatal("连接失败:", err)
	}
	defer wsClient.Close() // 程序退出时关闭连接

	// 步骤4: 创建区块头通道
	// Go 的 channel 用于协程间通信，这里接收新区块头
	headers := make(chan *types.Header) // 无缓冲通道，同步接收

	// 步骤5: 订阅新区块头事件
	// SubscribeNewHead 会在每个新区块产生时推送区块头信息
	sub, err := wsClient.SubscribeNewHead(
		context.Background(), // 使用默认上下文
		headers,              // 区块头会被推送到这个通道
	)
	if err != nil {
		log.Fatal("订阅失败:", err) // WebSocket 必须支持 eth_subscribe
	}

	fmt.Println("=== 监控 Sepolia 大额交易 ===")
	fmt.Println("监控标准: > 0.1 ETH")
	fmt.Println("")

	// 步骤6: 设置监控阈值
	// 0.1 ETH = 100000000000000000 Wei（1 ETH = 10^18 Wei）
	threshold := new(big.Int).Mul(
		big.NewInt(100000000000000000), // 0.1 ETH in Wei
		big.NewInt(1),
	)
	chainID := big.NewInt(11155111) // Sepolia 测试网的链 ID

	// 步骤7: 主事件循环
	for { // 无限循环，持续监听
		select { // Go 的 select 用于监听多个通道

		// 情况1: 订阅出错
		case err := <-sub.Err():
			log.Fatal(err) // 网络断开或其他错误，程序退出

		// 情况2: 收到新区块头
		case header := <-headers:
			// 步骤7.1: 获取完整区块数据
			// 区块头只有元信息，需要获取完整区块才能看到交易
			block, err := wsClient.BlockByHash(
				context.Background(),
				header.Hash(), // 用区块哈希获取完整区块
			)
			if err != nil {
				continue // 获取失败跳过这个区块
			}

			// 步骤7.2: 遍历区块中的所有交易
			for _, tx := range block.Transactions() {
				// 步骤7.3: 检查交易金额是否超过阈值
				// tx.Value() 返回交易金额（Wei）
				// Cmp 比较两个大数：返回 1(大于) 0(等于) -1(小于)
				if tx.Value().Cmp(threshold) > 0 {
					// 步骤7.4: 转换金额单位
					// 从 Wei 转换为 ETH（除以 10^18）
					ethValue := new(big.Float).SetInt(tx.Value()) // 先转为浮点数
					ethValue.Quo(ethValue, big.NewFloat(1e18))     // 除以 10^18

					// 步骤7.5: 输出交易信息
					fmt.Printf("🚨 大额交易发现!\n")
					fmt.Printf("  区块: #%d\n", block.Number().Uint64())
					fmt.Printf("  交易: %s\n", tx.Hash().Hex())
					fmt.Printf("  金额: %.4f ETH\n", ethValue)

					// 将 Unix 时间戳转换为可读时间
					fmt.Printf("  时间: %s\n",
						time.Unix(int64(block.Time()), 0).Format("15:04:05"))

					// 步骤7.6: 从签名恢复发送方地址
					// 以太坊交易不直接包含发送方地址，需要从签名推导
					from, err := types.Sender(
						types.LatestSignerForChainID(chainID), // EIP-155 签名器
						tx,                                     // 交易对象
					)
					if err == nil {
						fmt.Printf("  从: %s\n", from.Hex())
					}

					// 步骤7.7: 获取接收方地址
					// tx.To() 可能为 nil（合约创建交易）
					if tx.To() != nil {
						fmt.Printf("  到: %s\n", tx.To().Hex())
					}
					fmt.Println("") // 空行分隔
				}
			}
		}
	}
}