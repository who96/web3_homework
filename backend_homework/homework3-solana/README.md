# Solana Go 开发实战作业 - 基础链交互

## 项目结构

```
homework3-solana/
├── cmd/
│   └── chain-demo/     # 基础链交互演示程序
├── pkg/
│   └── chain/          # 核心链交互逻辑
│       ├── client.go   # RPC 客户端封装
│       └── transfer.go # 转账功能
├── config/             # 配置管理
│   └── config.go       # Devnet 配置
└── README.md
```

## 已实现功能（基础链交互 40%）

- ✅ 查询最新区块哈希 (`GetLatestBlockhash`)
- ✅ 查询账户余额 (`GetBalance`)
- ✅ SOL 转账 (`TransferSOL`)
- ✅ 交易确认等待 (`WaitForConfirmation`)

## 快速开始

### 1. 编译并运行演示程序

```bash
cd homework3-solana
go run cmd/chain-demo/main.go
```

### 2. 在 Devnet 上测试转账

**前置条件：**
- 安装 Solana CLI: https://docs.solana.com/cli/install-solana-cli-tools
- 生成密钥对: `solana-keygen new --outfile ~/.config/solana/devnet.json`
- 领取测试币: `solana airdrop 1 --url devnet`

**修改代码启用转账：**

在 `cmd/chain-demo/main.go` 中取消注释转账代码，并替换为你的私钥。

## 代码设计原则

### 1. 简洁性
- 每个函数只做一件事
- 不使用接口（没必要）
- 不做过度抽象

### 2. 错误处理
- 所有函数都返回 error
- 调用者必须处理错误
- 不使用 panic 或忽略错误

### 3. 数据结构优先
- Solana 的核心数据：Account, Transaction, Instruction
- 理解数据模型 > 堆砌代码

## Solana 关键概念

### Blockhash 有效期
Solana 的 blockhash 只有 **60-90 秒**有效期（约 150 个区块）。
- ❌ 不要缓存 blockhash
- ✅ 每次构造交易时重新获取

### 账户模型
```
Solana Account:
├── Address (32 bytes)
├── Lamports (balance, u64)
├── Data (bytes)
├── Owner (program ID)
└── Executable (bool)
```

对比以太坊:
```
EVM Account:
├── Address (20 bytes)
├── Balance (uint256)
├── Nonce (uint64)
└── Code (只有合约账户)
```

### 交易确认级别
1. **Processed**: 交易已包含在区块中（可能回滚）
2. **Confirmed**: 区块已被集群确认（乐观确认）
3. **Finalized**: 区块已最终确定（不可回滚）

**生产环境永远使用 Finalized**

## 下一步

- [ ] 智能合约开发（token-swap）
- [ ] 事件监听（WebSocket 订阅）
- [ ] 技术报告

## 参考资料

- Solana 官方文档: https://docs.solana.com
- Solana Go SDK: https://github.com/gagliardetto/solana-go
- Devnet Explorer: https://explorer.solana.com/?cluster=devnet
