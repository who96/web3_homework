// 配置管理包
// 功能：从 .env 文件加载配置信息，统一管理应用配置
package config

import (
	"log" // 日志记录
	"os"  // 操作系统接口，用于读取环境变量

	"github.com/joho/godotenv" // 第三方库，用于加载 .env 文件
)

// Config 结构体定义了应用程序需要的所有配置项
type Config struct {
	PrivateKey string // 以太坊账户私钥（16进制格式，包含0x前缀）
	RPCUrl     string // HTTP RPC 节点地址（https://开头，用于常规查询和交易）
	WSUrl      string // WebSocket RPC 节点地址（wss://开头，用于事件订阅）
	ChainID    int64  // 区块链网络ID（用于EIP-155签名，防止跨链重放攻击）
}

// Load 函数加载并返回配置
// 返回值：*Config 指针，包含所有配置信息
// 如果 .env 文件加载失败，程序会直接退出
func Load() *Config {
	// 使用 godotenv 库加载当前目录下的 .env 文件
	// .env 文件格式：KEY=VALUE，每行一个配置项
	err := godotenv.Load()
	if err != nil {
		// 加载失败时记录错误并退出程序
		// 这是一个关键配置，没有它程序无法运行
		log.Fatal("Error loading .env file")
	}

	// 创建并返回 Config 结构体实例
	// os.Getenv 从环境变量中读取值
	return &Config{
		// 私钥：用于签名交易
		// 格式：0x开头的64个十六进制字符
		// 安全提示：永远不要将私钥提交到代码仓库
		PrivateKey: os.Getenv("PRIVATE_KEY"),

		// HTTP RPC URL：用于与以太坊节点通信
		// 支持的操作：查询余额、发送交易、获取区块等
		// 格式：https://ethereum-sepolia-rpc.publicnode.com
		RPCUrl: os.Getenv("SEPOLIA_RPC_URL"),

		// WebSocket URL：用于订阅事件
		// 支持的操作：订阅新区块、监听事件日志等
		// 格式：wss://ethereum-sepolia-rpc.publicnode.com
		// 注意：HTTP 不支持订阅，必须使用 WebSocket
		WSUrl: os.Getenv("SEPOLIA_WS_URL"),

		// 链ID：Sepolia 测试网的 ID 是 11155111
		// 其他常见链ID：
		// - 主网(Mainnet): 1
		// - Goerli测试网: 5
		// - Sepolia测试网: 11155111
		// 用途：EIP-155 要求在签名中包含链ID，防止交易在不同链之间重放
		ChainID: 11155111, // Sepolia chain ID
	}
}