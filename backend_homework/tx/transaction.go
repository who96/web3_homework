// 交易管理包
// 功能：处理以太坊交易的完整生命周期，包括签名、广播、监控等
package tx

import (
	"context"       // 用于控制请求的生命周期和超时
	"crypto/ecdsa"  // 椭圆曲线数字签名算法，用于交易签名
	"fmt"           // 格式化输入输出
	"math/big"      // 大数运算，处理以太坊金额和数值
	"time"          // 时间相关操作

	"github.com/ethereum/go-ethereum"           // 以太坊核心错误类型
	"github.com/ethereum/go-ethereum/common"    // 通用类型（地址、哈希等）
	"github.com/ethereum/go-ethereum/core/types" // 核心数据类型（交易、区块、回执等）
	"github.com/ethereum/go-ethereum/crypto"     // 加密相关功能
	"github.com/ethereum/go-ethereum/ethclient"  // 以太坊客户端
)

// TransactionManager 交易管理器结构体
// 封装了与以太坊交互所需的所有核心组件
type TransactionManager struct {
	client     *ethclient.Client  // 以太坊客户端连接实例
	privateKey *ecdsa.PrivateKey  // ECDSA私钥，用于签名交易
	chainID    *big.Int           // 链ID，用于EIP-155签名防重放攻击
	address    common.Address     // 账户地址，从私钥推导而来
}

// NewTransactionManager 创建新的交易管理器实例
// 参数：
//   - rpcURL: 以太坊节点的RPC地址（HTTP或WebSocket）
//   - privateKeyHex: 16进制格式的私钥字符串（包含0x前缀）
//   - chainID: 区块链网络ID（主网=1，Sepolia=11155111）
// 返回：
//   - *TransactionManager: 交易管理器实例
//   - error: 错误信息
func NewTransactionManager(rpcURL string, privateKeyHex string, chainID int64) (*TransactionManager, error) {
	// 步骤1: 连接到以太坊节点
	// ethclient.Dial 会自动识别是 HTTP 还是 WebSocket 连接
	client, err := ethclient.Dial(rpcURL)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to Ethereum node: %v", err)
	}

	// 步骤2: 解析私钥字符串
	// crypto.HexToECDSA 需要不带0x前缀的纯16进制字符串
	// privateKeyHex[2:] 去掉前面的 "0x"
	privateKey, err := crypto.HexToECDSA(privateKeyHex[2:]) // Remove 0x prefix
	if err != nil {
		return nil, fmt.Errorf("failed to parse private key: %v", err)
	}

	// 步骤3: 从私钥推导公钥和地址
	// 以太坊地址是公钥的 Keccak-256 哈希的后20字节
	publicKey := privateKey.Public()                        // 获取公钥接口
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)      // 类型断言转换为ECDSA公钥
	if !ok {
		return nil, fmt.Errorf("failed to cast public key to ECDSA")
	}
	address := crypto.PubkeyToAddress(*publicKeyECDSA)      // 从公钥计算地址

	// 返回初始化完成的交易管理器
	return &TransactionManager{
		client:     client,
		privateKey: privateKey,
		chainID:    big.NewInt(chainID), // 转换为 big.Int 类型
		address:    address,
	}, nil
}

// DemoTransactionLifecycle 演示完整的交易生命周期
// 展示从构建交易到最终确认的完整流程
// 参数：
//   - toAddress: 接收方地址（16进制字符串）
//   - amount: 转账金额（Wei单位）
// 返回：
//   - error: 执行过程中的错误
func (tm *TransactionManager) DemoTransactionLifecycle(toAddress string, amount *big.Int) error {
	ctx := context.Background()           // 创建基础上下文
	to := common.HexToAddress(toAddress)  // 将字符串地址转换为Address类型

	// 打印交易基本信息
	fmt.Println("=== 交易生命周期演示 ===\n")
	fmt.Printf("从地址: %s\n", tm.address.Hex())
	fmt.Printf("到地址: %s\n", to.Hex())
	fmt.Printf("金额: %s wei\n\n", amount.String())

	// ========== 步骤1: 构建交易 ==========
	fmt.Println("【步骤1: 构建交易】")

	// 获取 nonce（交易序号）
	// nonce 是该账户发送的交易数量，用于防止重放攻击
	// PendingNonceAt 返回下一个可用的 nonce
	nonce, err := tm.client.PendingNonceAt(ctx, tm.address)
	if err != nil {
		return fmt.Errorf("获取 nonce 失败: %v", err)
	}
	fmt.Printf("Nonce: %d\n", nonce)

	// 获取建议的 gas 价格
	// SuggestGasPrice 返回节点建议的 gas 价格
	// gas 价格决定了交易的优先级和执行速度
	gasPrice, err := tm.client.SuggestGasPrice(ctx)
	if err != nil {
		return fmt.Errorf("获取 gas price 失败: %v", err)
	}
	fmt.Printf("Gas Price: %s wei\n", gasPrice.String())

	// 设置 gas limit
	// 21000 是 ETH 转账的标准 gas 消耗量
	// 智能合约调用通常需要更多 gas
	gasLimit := uint64(21000) // ETH 转账标准 gas limit
	fmt.Printf("Gas Limit: %d\n", gasLimit)

	// 创建未签名的交易对象
	// 参数: nonce, 接收地址, 金额, gas限制, gas价格, 数据(ETH转账为nil)
	tx := types.NewTransaction(nonce, to, amount, gasLimit, gasPrice, nil)
	fmt.Printf("交易哈希(未签名): %s\n\n", tx.Hash().Hex())

	// ========== 步骤2: 钱包签名 ==========
	fmt.Println("【步骤2: 钱包签名】")

	// 使用私钥对交易进行签名
	// EIP155Signer 包含了链ID，防止跨链重放攻击
	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(tm.chainID), tm.privateKey)
	if err != nil {
		return fmt.Errorf("签名交易失败: %v", err)
	}
	fmt.Printf("签名后交易哈希: %s\n", signedTx.Hash().Hex())

	// 显示签名的三个组成部分
	// v, r, s 是 ECDSA 签名的标准组成部分
	// v 包含了恢复ID和链ID信息
	v, r, s := signedTx.RawSignatureValues()
	fmt.Printf("签名 V: %s\n", v.String())  // 恢复ID + 链ID信息
	fmt.Printf("签名 R: %s\n", r.String())  // 签名的第一部分
	fmt.Printf("签名 S: %s\n\n", s.String()) // 签名的第二部分

	// ========== 步骤3: 交易广播 ==========
	fmt.Println("【步骤3: 交易广播】")

	// 将签名后的交易发送到网络
	// 节点会验证签名并广播给其他节点
	err = tm.client.SendTransaction(ctx, signedTx)
	if err != nil {
		return fmt.Errorf("发送交易失败: %v", err)
	}
	fmt.Printf("交易已广播到网络\n")
	fmt.Printf("交易哈希: %s\n\n", signedTx.Hash().Hex())

	// ========== 步骤4: 矿工打包 ==========
	fmt.Println("【步骤4: 矿工打包】")
	fmt.Println("等待交易被打包到区块...")

	// 等待交易被矿工打包并获取回执
	// 这个过程可能需要几秒到几十秒
	receipt, err := tm.WaitForReceipt(ctx, signedTx.Hash())
	if err != nil {
		return fmt.Errorf("等待交易确认失败: %v", err)
	}

	// ========== 步骤5: 区块广播与交易执行 ==========
	fmt.Println("【步骤5: 区块广播与交易执行】")

	// 显示交易执行结果
	fmt.Printf("交易已被包含在区块: %d\n", receipt.BlockNumber.Uint64())
	fmt.Printf("交易索引: %d\n", receipt.TransactionIndex) // 交易在区块中的位置
	fmt.Printf("Gas 使用量: %d\n", receipt.GasUsed)        // 实际消耗的 gas
	fmt.Printf("交易状态: %s\n", tm.getStatusText(receipt.Status))

	// 获取包含该交易的区块信息
	block, err := tm.client.BlockByNumber(ctx, receipt.BlockNumber)
	if err == nil {
		fmt.Printf("\n区块信息:\n")
		fmt.Printf("区块哈希: %s\n", block.Hash().Hex())
		fmt.Printf("区块时间戳: %s\n", time.Unix(int64(block.Time()), 0).Format("2006-01-02 15:04:05"))
		fmt.Printf("区块内交易数量: %d\n", len(block.Transactions()))
	}

	return nil
}

// WaitForReceipt 等待交易回执
// 轮询检查交易是否被打包到区块中
// 参数：
//   - ctx: 上下文
//   - txHash: 交易哈希
// 返回：
//   - *types.Receipt: 交易回执
//   - error: 错误信息
func (tm *TransactionManager) WaitForReceipt(ctx context.Context, txHash common.Hash) (*types.Receipt, error) {
	// 最多等待60秒
	for i := 0; i < 60; i++ {
		// 尝试获取交易回执
		// 如果交易已被打包，会返回回执对象
		receipt, err := tm.client.TransactionReceipt(ctx, txHash)
		if err == nil {
			return receipt, nil // 成功获取回执
		}

		// 如果错误是"未找到"，说明交易还未被打包
		if err == ethereum.NotFound {
			fmt.Print(".") // 打印进度点
			time.Sleep(1 * time.Second) // 等待1秒后重试
			continue
		}

		// 其他错误直接返回
		return nil, err
	}

	// 超时返回错误
	return nil, fmt.Errorf("timeout waiting for transaction receipt")
}

// SubscribeNewBlocks 订阅新区块事件
// 使用 WebSocket 连接实时接收新区块通知
// 参数：
//   - wsURL: WebSocket RPC地址（必须是wss://）
// 返回：
//   - error: 错误信息
func (tm *TransactionManager) SubscribeNewBlocks(wsURL string) error {
	// 创建专门的 WebSocket 连接用于订阅
	// 订阅功能只能通过 WebSocket 实现，不能用 HTTP
	wsClient, err := ethclient.Dial(wsURL)
	if err != nil {
		return fmt.Errorf("连接 WebSocket 失败: %v", err)
	}
	defer wsClient.Close() // 函数退出时关闭连接

	// 创建接收区块头的通道
	headers := make(chan *types.Header)

	// 订阅新区块头事件
	// SubscribeNewHead 会在每个新区块产生时推送区块头
	sub, err := wsClient.SubscribeNewHead(context.Background(), headers)
	if err != nil {
		return fmt.Errorf("订阅新区块失败: %v", err)
	}

	fmt.Println("开始订阅新区块...")
	fmt.Println("按 Ctrl+C 退出订阅")

	// 事件循环
	for {
		select {
		// 监听订阅错误
		case err := <-sub.Err():
			return err

		// 接收新区块头
		case header := <-headers:
			// 打印区块基本信息
			fmt.Printf("新区块: #%d, 哈希: %s, 时间: %s\n",
				header.Number.Uint64(),     // 区块号
				header.Hash().Hex(),         // 区块哈希
				time.Unix(int64(header.Time), 0).Format("15:04:05")) // 时间戳

			// 获取完整区块信息（可选）
			block, err := wsClient.BlockByHash(context.Background(), header.Hash())
			if err == nil {
				fmt.Printf("  交易数量: %d, Gas Used: %d\n",
					len(block.Transactions()), // 区块中的交易数
					header.GasUsed)             // 区块消耗的总 gas
			}
		}
	}
}

// SubscribeWithMonitor 订阅新区块并监控大额交易
// 增强版订阅：不仅显示新区块，还监控超过阈值的交易
// 参数：
//   - wsURL: WebSocket RPC地址（必须是wss://）
//   - threshold: 监控阈值（Wei单位），超过此金额的交易会被特别标记
// 返回：
//   - error: 错误信息
func (tm *TransactionManager) SubscribeWithMonitor(wsURL string, threshold *big.Int) error {
	// 创建 WebSocket 连接
	wsClient, err := ethclient.Dial(wsURL)
	if err != nil {
		return fmt.Errorf("连接 WebSocket 失败: %v", err)
	}
	defer wsClient.Close()

	// 创建区块头通道
	headers := make(chan *types.Header)

	// 订阅新区块头事件
	sub, err := wsClient.SubscribeNewHead(context.Background(), headers)
	if err != nil {
		return fmt.Errorf("订阅新区块失败: %v", err)
	}

	// 将阈值从 Wei 转换为 ETH 用于显示
	thresholdETH := new(big.Float).SetInt(threshold)
	thresholdETH.Quo(thresholdETH, big.NewFloat(1e18))

	fmt.Println("====== 监控模式已启动 ======")
	fmt.Printf("监控阈值: %.4f ETH\n", thresholdETH)
	fmt.Println("按 Ctrl+C 退出监控")
	fmt.Println("=============================")

	// 交易计数器
	totalLargeTransactions := 0

	// 事件循环
	for {
		select {
		case err := <-sub.Err():
			return err

		case header := <-headers:
			// 显示新区块信息
			fmt.Printf("📦 新区块: #%d | 时间: %s\n",
				header.Number.Uint64(),
				time.Unix(int64(header.Time), 0).Format("15:04:05"))

			// 获取完整区块
			block, err := wsClient.BlockByHash(context.Background(), header.Hash())
			if err != nil {
				fmt.Printf("  ⚠️  获取区块详情失败: %v\n", err)
				continue
			}

			// 统计本区块的大额交易
			largeTransactionsInBlock := 0

			// 遍历区块中的所有交易
			for _, tx := range block.Transactions() {
				// 检查交易金额是否超过阈值
				if tx.Value().Cmp(threshold) > 0 {
					largeTransactionsInBlock++
					totalLargeTransactions++

					// 转换金额为 ETH
					ethValue := new(big.Float).SetInt(tx.Value())
					ethValue.Quo(ethValue, big.NewFloat(1e18))

					// 获取发送方地址
					from, err := types.Sender(types.LatestSignerForChainID(tm.chainID), tx)
					fromStr := "未知"
					if err == nil {
						fromStr = from.Hex()
					}

					// 获取接收方地址
					toStr := "合约创建"
					if tx.To() != nil {
						toStr = tx.To().Hex()
					}

					// 打印大额交易警报
					fmt.Println("\n🚨 ===== 大额交易警报 =====")
					fmt.Printf("💰 金额: %.4f ETH\n", ethValue)
					fmt.Printf("📤 从: %s\n", fromStr)
					fmt.Printf("📥 到: %s\n", toStr)
					fmt.Printf("🔗 交易哈希: %s\n", tx.Hash().Hex())
					fmt.Println("==========================")
				}
			}

			// 区块统计信息
			fmt.Printf("  📊 区块统计: 总交易数 %d | 大额交易 %d\n",
				len(block.Transactions()),
				largeTransactionsInBlock)

			if largeTransactionsInBlock > 0 {
				fmt.Printf("  💡 累计发现大额交易: %d 笔\n", totalLargeTransactions)
			}

			fmt.Println("------------------------------")
		}
	}
}

// GetBalance 查询指定地址的余额
// 参数：
//   - address: 要查询的地址（16进制字符串）
// 返回：
//   - *big.Int: 余额（Wei单位）
//   - error: 错误信息
func (tm *TransactionManager) GetBalance(address string) (*big.Int, error) {
	// 将字符串地址转换为 Address 类型
	addr := common.HexToAddress(address)

	// 查询最新区块的余额
	// 第二个参数 nil 表示查询最新区块
	balance, err := tm.client.BalanceAt(context.Background(), addr, nil)
	if err != nil {
		return nil, err
	}
	return balance, nil
}

// getStatusText 将交易状态码转换为可读文本
// 内部辅助函数
// 参数：
//   - status: 状态码（1=成功，0=失败）
// 返回：
//   - string: 状态文本
func (tm *TransactionManager) getStatusText(status uint64) string {
	if status == 1 {
		return "成功 ✅"
	}
	return "失败 ❌"
}

// Close 关闭与以太坊节点的连接
// 应该在程序退出前调用
func (tm *TransactionManager) Close() {
	tm.client.Close()
}