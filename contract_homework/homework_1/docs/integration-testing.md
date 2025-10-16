# 集成测试指南

## 概述

本项目包含完整的集成测试套件，用于验证 MemeToken 与 Uniswap V2 的真实交互。集成测试使用 Foundry 的 **mainnet fork** 功能，在本地环境中模拟主网状态。

## 测试文件

- **位置**: `test/MemeToken.integration.t.sol`
- **测试用例数**: 9个
- **测试内容**:
  - 添加 Uniswap V2 流动性
  - 通过 Uniswap 买入代币（验证税费收取）
  - 通过 Uniswap 卖出代币（验证税费收取）
  - 白名单机制（Pair 地址免税）
  - 交易限制验证
  - Gas 基准测试

## 运行集成测试

### 前置条件

1. **获取 Mainnet RPC URL**

   推荐使用免费的 RPC 提供商：
   - [Alchemy](https://www.alchemy.com/) - 免费额度 300M compute units/月
   - [Infura](https://www.infura.io/) - 免费额度 100,000 requests/天
   - [QuickNode](https://www.quicknode.com/) - 免费额度有限

   注册后获取 Mainnet Ethereum RPC URL，格式：
   ```
   https://eth-mainnet.g.alchemy.com/v2/YOUR_API_KEY
   ```

2. **设置环境变量**

   在 `.env` 文件中添加：
   ```bash
   MAINNET_RPC_URL=https://eth-mainnet.g.alchemy.com/v2/YOUR_API_KEY
   ```

   或者在命令行中直接指定：
   ```bash
   export MAINNET_RPC_URL=https://eth-mainnet.g.alchemy.com/v2/YOUR_API_KEY
   ```

### 运行命令

#### 1. 运行所有集成测试

```bash
forge test --match-contract MemeTokenIntegrationTest --fork-url $MAINNET_RPC_URL -vv
```

**参数说明**:
- `--match-contract MemeTokenIntegrationTest`: 只运行集成测试合约
- `--fork-url $MAINNET_RPC_URL`: 指定要 fork 的主网 RPC
- `-vv`: 显示详细日志（包括 console.log 输出）

#### 2. 运行特定测试用例

```bash
# 只测试添加流动性
forge test --match-test test_AddLiquidityToUniswapV2 --fork-url $MAINNET_RPC_URL -vv

# 只测试买入代币
forge test --match-test test_BuyTokensWithTax --fork-url $MAINNET_RPC_URL -vv

# 只测试卖出代币
forge test --match-test test_SellTokensWithTax --fork-url $MAINNET_RPC_URL -vv
```

#### 3. 运行所有测试（单元测试 + 集成测试）

```bash
forge test --fork-url $MAINNET_RPC_URL -vv
```

## 测试用例详解

### 1. `test_AddLiquidityToUniswapV2()`

**功能**: 验证添加 Uniswap V2 流动性

**步骤**:
1. Owner 转账 210,000 FK 给流动性提供者
2. 流动性提供者授权 Router 使用代币
3. 调用 `addLiquidityETH()` 添加流动性（210,000 FK + 0.1 ETH）
4. 验证 Pair 地址创建成功
5. 验证 LP tokens 发放

**预期结果**:
- 流动性添加成功
- 获得 Pair 地址
- 收到 LP tokens

---

### 2. `test_BuyTokensWithTax()`

**功能**: 验证通过 Uniswap 买入代币时收取税费

**步骤**:
1. 添加流动性
2. 将 Pair 地址加入白名单（避免流动性操作收税）
3. User1 使用 0.01 ETH 通过 `swapExactETHForTokens()` 买入代币
4. 验证 User1 收到代币（扣除税费后）
5. 验证税收地址收到 3% 税费

**预期结果**:
- User1 收到代币
- 税收地址余额增加约 3%
- 税费计算精确

---

### 3. `test_SellTokensWithTax()`

**功能**: 验证通过 Uniswap 卖出代币时收取税费

**步骤**:
1. 先执行买入操作（获得代币）
2. User1 卖出 50% 的代币
3. 验证税收地址再次收到税费

**预期结果**:
- 卖出成功
- 税收地址余额再次增加

---

### 4. `test_PairWhitelistNoTax()`

**功能**: 验证 Pair 地址在白名单中时不收税

**步骤**:
1. 添加流动性
2. 将 Pair 地址加入税收白名单
3. User1 直接转账给 Pair 地址
4. 验证税收地址余额不变

**预期结果**:
- 转账成功
- 不收取税费

---

### 5. `test_MaxTransactionLimitOnSwap()`

**功能**: 验证交易限制功能

**步骤**:
1. 添加流动性
2. 尝试转账超过 `maxTransactionAmount` 的代币
3. 验证交易被拒绝

**预期结果**:
- 交易 revert，提示 "Transfer amount exceeds max transaction amount"

---

### 6. `test_GasSwapWithTax()`

**功能**: 测量 Uniswap swap 的 Gas 消耗

**步骤**:
1. 添加流动性
2. 执行 swap 交易
3. 记录 Gas 使用量

**预期结果**:
- Gas < 300,000（包含 Uniswap 逻辑 + 税收逻辑）

---

## 常见问题

### Q1: RPC 请求超限怎么办？

**A**: 集成测试会消耗 RPC 配额。如果遇到限制：
1. 使用多个 RPC 提供商轮流使用
2. 等待配额恢复（通常按天重置）
3. 升级到付费计划（不推荐，测试不应该花钱）

### Q2: Fork 测试很慢？

**A**: 这是正常的，因为需要从主网同步状态。优化方法：
1. 使用更快的 RPC 提供商（如 Alchemy）
2. 只运行特定测试用例而不是全部
3. 考虑本地运行 Geth 归档节点（仅适合高级用户）

### Q3: 测试失败："revert: Transaction reverted"？

**A**: 可能原因：
1. RPC URL 不正确或无效
2. 网络连接问题
3. Uniswap 合约地址错误（检查是否使用主网地址）

解决方法：
```bash
# 测试 RPC 连接
cast block-number --rpc-url $MAINNET_RPC_URL

# 验证 Uniswap Router 地址
cast code 0x7a250d5630B4cF539739dF2C5dAcb4c659F2488D --rpc-url $MAINNET_RPC_URL
```

### Q4: 需要真实 ETH 吗？

**A**: **不需要**！Fork 测试在本地环境运行，所有账户余额都是模拟的。`vm.deal()` 会给测试账户分配任意数量的 ETH。

### Q5: 能否在测试网运行集成测试？

**A**: 可以，但需要：
1. 修改 Uniswap 地址为测试网地址（Sepolia 等）
2. 确保测试网有真实的 Uniswap V2 部署
3. 准备测试网 ETH

不推荐，因为 mainnet fork 更接近真实环境。

---

## Gas 基准测试结果

| 操作 | Gas 消耗 | 说明 |
|------|---------|------|
| 白名单转账（无税） | ~29,700 | 基础 ERC20 转账 + 白名单检查 |
| 含税转账 | ~64,713 | 基础转账 + 税费计算 + 双次 transfer |
| Uniswap Swap（含税） | <300,000 | Uniswap 逻辑 + 我们的税收机制 |

---

## 安全审计结果

**Slither 静态分析**: ✅ **0个问题**

```bash
# 运行 Slither
source .venv/bin/activate
slither src/MemeToken.sol --solc-remaps "@openzeppelin/=lib/openzeppelin-contracts/" --filter-paths "lib/"
```

**输出**:
```
INFO:Slither:src/MemeToken.sol analyzed (9 contracts with 100 detectors), 0 result(s) found
```

---

## 下一步

集成测试已完成，如需部署到测试网：

1. 参考 [`docs/deployment-guide.md`](./deployment-guide.md)
2. 准备 Sepolia 测试网 ETH
3. 配置 `.env` 文件（私钥、RPC URL）
4. 运行部署脚本：
   ```bash
   forge script script/DeployMemeToken.s.sol --rpc-url $SEPOLIA_RPC_URL --broadcast
   ```

---

## Linus 的建议

> "测试不是为了证明代码没有 bug，而是为了在用户发现之前找到 bug。"

**集成测试的价值**:
1. **真实环境验证**: 单元测试无法覆盖与外部协议（Uniswap）的交互
2. **边界情况发现**: 真实的链上状态可能暴露意外行为
3. **Gas 成本确认**: 实际 gas 消耗可能与理论计算不同

**但请记住**:
- 不要过度测试：测试覆盖率 100% 不等于没有 bug
- 保持测试简单：复杂的测试代码本身就是 bug 的来源
- 优先测试核心逻辑：税收机制 > 边界情况 > 美化输出

> "Talk is cheap. Show me the code."
> —— Linus Torvalds
