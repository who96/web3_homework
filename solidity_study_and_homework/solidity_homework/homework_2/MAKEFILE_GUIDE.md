# Makefile 使用指南

## 基础命令

### 编译和测试
```bash
make build              # 编译所有合约
make test               # 运行所有测试
make test-match TEST=test_MintNFT  # 运行特定测试
make gas                # 生成Gas报告
make coverage           # 生成测试覆盖率报告
make format             # 格式化代码
make clean              # 清理编译产物
```

### BeggingContract专用命令
```bash
make test-beggar        # 运行BeggingContract本地测试
make deploy-beggar      # 部署BeggingContract到测试网
make deploy-beggar-fast # 快速部署（不验证合约）

# 需要先设置合约地址环境变量
export CONTRACT_ADDRESS=0x1234...
make test-beggar-sepolia  # Sepolia交互测试（包含提现测试）
make beggar-info         # 查看合约基本信息
make withdraw-beggar     # 提现合约资金（仅owner，有确认提示）
```

## 部署命令

### 本地部署（Anvil）
```bash
# 启动本地节点（一个终端）
make anvil

# 部署到本地（另一个终端）
make deploy-local
```

### 测试网部署
```bash
# 默认部署到Sepolia
make deploy-testnet

# 指定网络
make deploy-testnet NETWORK=goerli

# BeggingContract部署
make deploy-beggar NETWORK=sepolia
```

### 主网部署（危险⚠️）
```bash
make deploy-mainnet     # 有确认提示
```

## 环境变量

### 必需变量
创建`.env`文件：
```bash
# 私钥（不含0x前缀）
PRIVATE_KEY=your_private_key_here

# Etherscan API Key（用于验证）
ETHERSCAN_API_KEY=your_api_key_here
```

### 可选变量
```bash
# 自定义RPC URLs
SEPOLIA_RPC_URL=https://your-rpc-url
MAINNET_RPC_URL=https://your-mainnet-rpc

# 指定网络（默认sepolia）
NETWORK=sepolia

# 指定合约地址（用于交互测试）
CONTRACT_ADDRESS=0x1234...
```

## 完整工作流程

### SimpleNFT部署流程
```bash
# 1. 编译和测试
make build
make test

# 2. 部署到Sepolia
make deploy-testnet

# 3. 记录合约地址并验证
# (部署脚本会自动显示合约地址)
```

### BeggingContract部署流程
```bash
# 1. 本地测试
make test-beggar

# 2. 部署到Sepolia
make deploy-beggar

# 3. 设置合约地址并测试
export CONTRACT_ADDRESS=0xABC123...
make test-beggar-sepolia

# 4. 查看合约信息
make beggar-info
```

## 故障排除

### 编译错误
```bash
# 清理缓存重新编译
make clean
make build

# 检查Solidity版本
cat foundry.toml | grep solc
```

### 部署失败
```bash
# 检查环境变量
cat .env

# 检查账户余额（需要测试ETH）
cast balance $(cast wallet address --private-key $PRIVATE_KEY) --rpc-url sepolia

# 手动指定gas limit
forge script script/Deploy.s.sol --rpc-url sepolia --broadcast --gas-limit 500000
```

### 测试失败
```bash
# 详细输出模式
make test -vvvv

# 运行特定测试
make test-match TEST=test_specific_function
```

## 实用技巧

### 1. 快速部署测试循环
```bash
# 一键编译+测试+部署
make build && make test && make deploy-beggar
```

### 2. 监控合约
```bash
# 设置合约地址后查看状态
export CONTRACT_ADDRESS=0x1234...
make beggar-info

# 实时监控事件
cast logs --rpc-url sepolia --address $CONTRACT_ADDRESS
```

### 3. Gas优化
```bash
# 查看gas使用报告
make gas

# 启用优化器编译
# 在 foundry.toml 中设置 optimizer = true
```

### 4. 批量操作
```bash
# 部署多个合约
make deploy-testnet && make deploy-beggar

# 运行所有测试
make test && make test-beggar
```

## 命令速查

| 命令 | 功能 | 备注 |
|------|------|------|
| `make build` | 编译合约 | 基础操作 |
| `make test` | 运行测试 | 本地测试 |
| `make deploy-testnet` | 部署SimpleNFT | 需要.env |
| `make deploy-beggar` | 部署BeggingContract | 需要.env |
| `make test-beggar` | BeggingContract测试 | 本地 |
| `make test-beggar-sepolia` | Sepolia完整测试 | 需要CONTRACT_ADDRESS |
| `make beggar-info` | 查看合约信息 | 需要CONTRACT_ADDRESS |
| `make withdraw-beggar` | 提现合约资金 | 仅owner，有确认 |
| `make anvil` | 启动本地节点 | 开发调试 |
| `make clean` | 清理缓存 | 故障排除 |

记住：**测试先于部署，本地先于远程**！