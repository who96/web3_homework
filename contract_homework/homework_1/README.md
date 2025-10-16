# Fukua (FK) - SHIB风格Meme代币

一个完全符合ERC20标准的Meme代币，支持交易税、白名单机制和防操纵保护。

## 特性

✅ **标准ERC20** - 完全兼容ERC20接口
✅ **交易税机制** - 可配置1-5%税率
✅ **白名单系统** - owner和LP免税
✅ **交易限制** - 单笔最大额度限制
✅ **Gas优化** - 白名单转账<30k gas, 含税转账<65k gas
✅ **安全可靠** - 基于OpenZeppelin, 100%测试覆盖
✅ **DEX兼容** - 完美支持Uniswap V2

## 代币信息

| 属性 | 值 |
|------|-----|
| 名称 | Fukua |
| 符号 | FK |
| 总供应量 | 21,000,000 FK |
| 初始税率 | 3% (可调) |
| 最大税率 | 5% |
| 单笔限额 | 420,000 FK (2%, 可调) |

## 快速开始

### 安装依赖

```bash
forge install
```

### 编译合约

```bash
forge build
```

### 运行测试

```bash
# 单元测试
forge test -vvv

# 集成测试 (需要 mainnet RPC)
forge test --match-contract MemeTokenIntegrationTest --fork-url $MAINNET_RPC_URL -vv
```

**测试结果:**
- 单元测试: 33/33 passed ✓
- 集成测试: 9/9 passed ✓
- **Slither审计: 0个问题** ✓

### 部署

```bash
# 测试网
forge script script/DeployMemeToken.s.sol \
  --rpc-url $SEPOLIA_RPC_URL \
  --broadcast \
  --verify

# 本地测试
anvil
forge script script/DeployMemeToken.s.sol \
  --rpc-url http://localhost:8545 \
  --broadcast
```

## 部署信息

🚀 **Sepolia测试网**:
- **MemeToken**: `0x61a33158B1541AD0fc87DF41075ac6A40CC52498`
  - [Etherscan验证](https://sepolia.etherscan.io/address/0x61a33158b1541ad0fc87df41075ac6a40cc52498) ✅
- **MockDEX**: `0x4961dDb55265Bcd4E230B2aceaf257A745e73de0`
  - 流动性: 210,000 FK + 0.1 ETH ✅
  - [Swap测试](https://sepolia.etherscan.io/tx/0x139c316ce0c74364a935f6601cb570226fc7f37c3fa88b3fc22c50c3a3585595) ✅
- [完整部署记录](docs/sepolia-deployment.md)

## 文档

📖 [完整部署指南](docs/deployment-guide.md)
🧪 [集成测试指南](docs/integration-testing.md)
📋 [Sepolia部署记录](docs/sepolia-deployment.md)
📋 [OpenSpec文档](openspec/changes/add-meme-token-contract/)
🧪 [单元测试](test/MemeToken.t.sol) | [集成测试](test/MemeToken.integration.t.sol)

## 项目结构

```
.
├── src/
│   ├── MemeToken.sol             # 核心合约 (239行)
│   └── MockDEX.sol               # 测试DEX (180行)
├── test/
│   ├── MemeToken.t.sol           # 单元测试 (33 tests)
│   └── MemeToken.integration.t.sol # 集成测试 (9 tests)
├── script/
│   ├── DeployMemeToken.s.sol     # MemeToken部署
│   ├── DeployMockDEX.s.sol       # MockDEX部署
│   └── AddLiquidity.s.sol        # Uniswap流动性
├── docs/
│   ├── deployment-guide.md       # 部署指南
│   ├── integration-testing.md    # 集成测试指南
│   └── sepolia-deployment.md     # Sepolia部署记录
├── openspec/
│   ├── project.md                # 项目上下文
│   └── changes/
│       └── add-meme-token-contract/
│           ├── proposal.md       # 变更提案
│           ├── design.md         # 设计文档
│           ├── tasks.md          # 任务清单 (107/109完成)
│           └── specs/            # 规格说明
└── foundry.toml                  # Foundry配置
```

## 主要文件说明

### 📄 核心合约 (`src/`)

#### `MemeToken.sol` (239行)
**功能**: Fukua代币的核心智能合约

**主要特性**:
- 继承OpenZeppelin的`ERC20`和`Ownable`
- 实现交易税机制（1-5%可配置）
- 白名单系统（税收豁免和交易限制豁免）
- 单笔最大交易额度限制
- 完整的NatSpec文档注释

**核心方法**:
```solidity
// 税收管理
function setTaxRate(uint256 newRate) external onlyOwner
function setTaxRecipient(address newRecipient) external onlyOwner
function calculateTax(uint256 amount) external view returns (uint256)

// 白名单管理
function setTaxExempt(address account, bool exempt) external onlyOwner
function setLimitExempt(address account, bool exempt) external onlyOwner

// 交易限制
function setMaxTransactionAmount(uint256 newAmount) external onlyOwner

// 查询函数
function taxRate() external view returns (uint256)
function taxRecipient() external view returns (address)
function isExemptFromTax(address account) external view returns (bool)
```

**关键设计**:
- 使用`_update()`覆盖点统一处理所有转账
- Basis points系统（10000 = 100%）实现精确税率
- Mapping实现O(1)白名单查询
- 所有状态变更都触发事件

**Sepolia部署**: `0x61a33158B1541AD0fc87DF41075ac6A40CC52498`

---

#### `MockDEX.sol` (180行)
**功能**: 简化的去中心化交易所，用于测试网流动性演示

**主要特性**:
- 恒定乘积AMM算法 (x * y = k)
- 支持添加/移除流动性
- 双向Swap交易（ETH ⇄ Token）
- 0.3%交易手续费（模拟Uniswap）
- LP份额管理

**核心方法**:
```solidity
// 流动性管理
function addLiquidity(uint256 tokenAmount) external payable returns (uint256 shares)
function removeLiquidity(uint256 shares) external returns (uint256 tokenAmount, uint256 ethAmount)

// 交易功能
function swapEthForToken(uint256 minTokenOut) external payable returns (uint256 tokenOut)
function swapTokenForEth(uint256 tokenIn, uint256 minEthOut) external returns (uint256 ethOut)

// 查询功能
function getTokenOut(uint256 ethIn) external view returns (uint256)
function getEthOut(uint256 tokenIn) external view returns (uint256)
```

**使用场景**: Sepolia测试网没有官方Uniswap V2，此合约用于演示MemeToken的DEX兼容性

**Sepolia部署**: `0x4961dDb55265Bcd4E230B2aceaf257A745e73de0`

---

### 🧪 测试文件 (`test/`)

#### `MemeToken.t.sol` (418行, 33个测试)
**功能**: MemeToken合约的完整单元测试套件

**测试覆盖**:
1. **基础功能测试** (7个)
   - 初始状态验证
   - ERC20标准功能（transfer、approve、transferFrom）
   - Owner权限验证

2. **税收机制测试** (9个)
   - 普通转账收税验证
   - 白名单免税验证
   - 税率修改和权限控制
   - 边界情况（税率为0、税费计算精度）

3. **交易限制测试** (4个)
   - 超限转账拒绝
   - 白名单豁免验证
   - 限额修改功能

4. **边界情况测试** (5个)
   - 零额转账
   - 自我转账
   - 全额转账
   - 小额税费舍入

5. **Fuzz测试** (2个)
   - 随机金额转账税费正确性
   - 随机税率设置

6. **Gas基准测试** (2个)
   - 白名单转账: ~29,700 gas
   - 含税转账: ~64,713 gas

**运行方式**:
```bash
forge test -vvv                           # 所有测试
forge test --match-test test_TransferWithTax  # 单个测试
forge test --gas-report                   # Gas报告
```

**测试结果**: ✅ 33/33 passed

---

#### `MemeToken.integration.t.sol` (327行, 9个测试)
**功能**: 与真实Uniswap V2的集成测试（mainnet fork）

**测试场景**:
1. **流动性管理**
   - 添加流动性到Uniswap V2
   - 验证Pair地址创建
   - LP tokens发放验证

2. **交易测试**
   - 买入代币（ETH → FK），验证税费收取
   - 卖出代币（FK → ETH），验证税费收取
   - 白名单Pair地址免税验证

3. **交易限制测试**
   - 验证maxTransactionAmount限制生效
   - 白名单地址豁免验证

4. **Gas基准测试**
   - Uniswap Swap + 税收机制总Gas消耗

**运行方式**:
```bash
# 需要mainnet RPC URL
forge test --match-contract MemeTokenIntegrationTest \
  --fork-url $MAINNET_RPC_URL -vv
```

**测试结果**: ✅ 9/9 passed

**重要性**: 验证合约与真实Uniswap V2的兼容性，确保主网部署后能正常工作

---

### 🚀 部署脚本 (`script/`)

#### `DeployMemeToken.s.sol` (74行)
**功能**: MemeToken合约的自动化部署脚本

**部署流程**:
1. 读取环境变量（私钥、税收地址）
2. 部署MemeToken合约
3. 验证部署状态（owner、税率、白名单）
4. 输出部署信息和后续步骤

**环境变量需求**:
```bash
PRIVATE_KEY=0x...           # 部署者私钥
WALLET1=0x...               # 税收接收地址
SEPOLIA_RPC_URL=https://... # RPC URL
ETHERSCAN_API_KEY=...       # Etherscan验证密钥
```

**使用方式**:
```bash
# 测试网部署
forge script script/DeployMemeToken.s.sol \
  --rpc-url $SEPOLIA_RPC_URL \
  --broadcast \
  --verify

# 本地Anvil测试
anvil
forge script script/DeployMemeToken.s.sol \
  --rpc-url http://localhost:8545 \
  --broadcast
```

**输出示例**:
```
Token deployed at: 0x61a33158B1541AD0fc87DF41075ac6A40CC52498
Total supply: 21,000,000 FK
Max transaction amount: 420,000 FK
```

---

#### `DeployMockDEX.s.sol` (90行)
**功能**: 在Sepolia部署MockDEX并自动添加流动性

**部署流程**:
1. 读取已部署的MemeToken地址
2. 部署MockDEX合约
3. 授权DEX使用代币
4. 添加流动性（210,000 FK + 0.1 ETH）
5. 输出白名单设置命令

**环境变量需求**:
```bash
PRIVATE_KEY=0x...
SEPOLIA_RPC_URL=https://...
```

**使用方式**:
```bash
forge script script/DeployMockDEX.s.sol \
  --rpc-url $SEPOLIA_RPC_URL \
  --broadcast -vvv
```

**部署后操作**: 脚本会输出两条命令，用于将DEX地址加入白名单

---

#### `AddLiquidity.s.sol` (130行)
**功能**: 向Uniswap V2添加流动性的辅助脚本（主网/测试网）

**功能说明**:
- 适配真实Uniswap V2 Router地址
- 自动创建WETH-MemeToken交易对
- 支持滑点设置（默认5%）
- 自动输出Pair地址和白名单命令

**使用方式**:
```bash
# 需要先设置环境变量
export MEME_TOKEN_ADDRESS=0x...

forge script script/AddLiquidity.s.sol \
  --rpc-url $MAINNET_RPC_URL \
  --broadcast
```

**注意**: 仅在有真实Uniswap V2部署的网络上使用（主网、某些测试网）

---

### 📚 文档文件 (`docs/`)

#### `deployment-guide.md` (10KB)
**功能**: 完整的部署和使用指南

**主要内容**:
1. **环境准备**
   - Foundry安装和配置
   - 私钥和RPC URL配置
   - 依赖安装

2. **本地测试部署**
   - Anvil本地节点启动
   - 部署步骤和验证

3. **测试网部署**
   - Sepolia部署完整流程
   - Etherscan合约验证
   - 测试交易执行

4. **主网部署准备**
   - 安全检查清单
   - Gas费用估算
   - 部署步骤详解

5. **Uniswap流动性添加**
   - 代币授权
   - 添加流动性命令
   - Pair地址白名单设置

6. **使用指南**
   - 代币转账
   - 税费查询
   - 白名单管理
   - Owner权限操作

7. **FAQ**
   - 常见问题解答
   - 故障排查

---

#### `integration-testing.md` (8KB)
**功能**: Mainnet fork集成测试详细指南

**主要内容**:
1. **测试概述**
   - Fork测试原理
   - 测试用例介绍

2. **运行集成测试**
   - 前置条件（RPC URL获取）
   - 运行命令详解
   - 参数说明

3. **测试用例详解** (9个)
   - 每个测试的目的、步骤、预期结果
   - 测试覆盖的功能点

4. **常见问题**
   - RPC限制处理
   - 性能优化
   - 故障排查

5. **Gas基准测试结果**
   - 各操作Gas消耗统计

6. **安全审计结果**
   - Slither静态分析结果
   - 0个问题的证明

---

#### `sepolia-deployment.md` (12KB)
**功能**: Sepolia测试网部署的完整记录

**主要内容**:
1. **部署信息**
   - MemeToken合约地址和交易哈希
   - MockDEX合约地址和流动性信息
   - Etherscan链接

2. **合约配置**
   - 代币参数（名称、符号、供应量）
   - 税率和限制设置
   - Owner地址

3. **功能测试结果** (5个测试场景)
   - 白名单免税转账（TX: 0xbb6467...）
   - 含税转账（TX: 0x74d254...）
   - 交易限制验证
   - Owner权限测试（TX: 0xb24ed4...）
   - DEX流动性测试（TX: 0x139c31...）

4. **Gas消耗统计**
   - 部署、转账、Swap各操作的实际Gas消耗

5. **测试覆盖总结**
   - 功能覆盖清单
   - 测试类型统计

6. **与合约交互示例**
   - Cast命令示例（查询和交易）

7. **实施方案说明**
   - MockDEX解决方案
   - 测试网限制的应对

8. **安全建议**
   - 中心化风险提示
   - 主网部署建议

---

### 📋 OpenSpec文档 (`openspec/`)

#### `project.md`
**功能**: 项目上下文和技术规范

**内容**:
- 技术栈说明（Foundry、Solidity、OpenZeppelin）
- 代币经济学设计
- DEX集成方案
- 编码规范（命名、注释、测试要求）

---

#### `changes/add-meme-token-contract/proposal.md`
**功能**: 变更提案文档

**内容**:
- 为什么需要这个Meme代币
- 要实现什么功能
- 对项目的影响评估

---

#### `changes/add-meme-token-contract/design.md`
**功能**: 技术设计文档

**内容**:
- 5个关键技术决策
  1. 税收机制在`_update()`中实现
  2. 使用mapping实现白名单
  3. Basis points税率系统
  4. 单笔交易限制（不做每日限制）
  5. 完全ERC20兼容

---

#### `changes/add-meme-token-contract/tasks.md` (107/109完成)
**功能**: 详细任务清单和进度跟踪

**内容**:
- 10个阶段共109个子任务
- 每个任务的完成状态标记 [x]
- Gas基准测试结果
- 测试网部署记录（TX哈希、地址、Gas消耗）

**当前状态**: 98.2%完成（107/109）

---

#### `changes/add-meme-token-contract/specs/meme-token/spec.md`
**功能**: 详细规格说明（Requirements）

**内容**:
- 10个需求（Requirements）
- 每个需求包含多个场景（Scenarios）
- 使用SHALL/MUST规范语言
- 40+个测试场景定义

---

### ⚙️ 配置文件

#### `foundry.toml`
**功能**: Foundry框架配置

**主要配置**:
```toml
[profile.default]
src = "src"
out = "out"
libs = ["lib"]
solc_version = "0.8.20"
optimizer = true
optimizer_runs = 200

remappings = [
    "@openzeppelin/=lib/openzeppelin-contracts/",
    "forge-std/=lib/forge-std/src/"
]
```

---

#### `.env`
**功能**: 环境变量配置（不提交到Git）

**必需变量**:
```bash
PRIVATE_KEY=0x...              # 部署者私钥
WALLET1=0x...                  # 税收接收地址
SEPOLIA_RPC_URL=https://...    # Sepolia RPC
MAINNET_RPC_URL=https://...    # 主网RPC（fork测试用）
ETHERSCAN_API_KEY=...          # 合约验证密钥
```

---

#### `.gitignore`
**功能**: Git忽略文件配置

**忽略内容**:
- `out/` - 编译输出
- `cache/` - Foundry缓存
- `broadcast/` - 部署记录
- `.env` - 私钥等敏感信息
- `node_modules/` - 依赖

---

## 文件使用流程

### 1️⃣ 本地开发流程
```bash
# 1. 阅读规格
openspec/changes/add-meme-token-contract/specs/

# 2. 实现合约
src/MemeToken.sol

# 3. 编写测试
test/MemeToken.t.sol

# 4. 运行测试
forge test -vvv

# 5. 集成测试
test/MemeToken.integration.t.sol
forge test --fork-url $MAINNET_RPC_URL
```

### 2️⃣ 部署流程
```bash
# 1. 配置环境
.env

# 2. 部署合约
script/DeployMemeToken.s.sol

# 3. 部署DEX（测试网）
script/DeployMockDEX.s.sol

# 4. 添加流动性（主网）
script/AddLiquidity.s.sol

# 5. 查看记录
docs/sepolia-deployment.md
```

### 3️⃣ 使用流程
```bash
# 1. 阅读使用指南
docs/deployment-guide.md

# 2. 查看部署记录
docs/sepolia-deployment.md

# 3. 与合约交互
cast call/send ...
```

---

## 核心功能

### 1. 交易税机制

每笔转账自动扣除可配置的税费（默认3%），发送到指定的税收地址。

```solidity
// 查询税率
uint256 rate = token.taxRate(); // 返回300 (3%)

// 计算税费
uint256 tax = token.calculateTax(1000 ether); // 30 ether
```

### 2. 白名单系统

Owner、流动性池、合约本身默认免税。Owner可以添加其他地址。

```solidity
// 检查白名单
bool exempt = token.isExemptFromTax(address);

// 添加白名单 (仅owner)
token.setTaxExempt(address, true);
```

### 3. 交易限制

防止大额抛售，单笔转账不能超过限额（默认420,000 FK）。

```solidity
// 查询限额
uint256 maxTx = token.maxTransactionAmount();

// 修改限额 (仅owner)
token.setMaxTransactionAmount(newAmount);
```

## 安全特性

- ✅ 基于OpenZeppelin v5.4.0成熟合约
- ✅ Solidity 0.8.20内置溢出保护
- ✅ 税率上限5%硬编码限制
- ✅ 所有admin函数仅owner可调用
- ✅ 完整的事件记录
- ✅ 详细的NatSpec文档
- ✅ 100%测试覆盖 + Fuzz测试
- ✅ **Slither静态分析: 0个问题**

## Gas成本

| 操作 | Gas成本 |
|------|---------|
| 白名单转账 | ~29,700 gas |
| 含税转账 | ~64,713 gas |
| 部署 | ~2,500,000 gas |

## 技术栈

- **语言:** Solidity ^0.8.20
- **框架:** Foundry
- **测试:** Forge + Fuzz Testing
- **基础合约:** OpenZeppelin Contracts v5.4.0
- **DEX:** Uniswap V2

## 开发指南

### 添加新功能

1. 创建OpenSpec变更提案
2. 编写规格说明和设计文档
3. 实现功能
4. 编写测试 (100%覆盖)
5. 更新文档

### 运行特定测试

```bash
# 单个测试
forge test --match-test test_TransferWithTax -vvv

# 测试合约
forge test --match-contract MemeTokenTest

# Fuzz测试
forge test --match-test testFuzz -vvv

# Gas报告
forge test --gas-report
```

### 代码格式化

```bash
forge fmt
```

### 安全审计

```bash
# 安装Slither
python3 -m venv .venv
source .venv/bin/activate
pip install slither-analyzer

# 运行审计
slither src/MemeToken.sol --solc-remaps "@openzeppelin/=lib/openzeppelin-contracts/" --filter-paths "lib/"
```

## 常见问题

**Q: 为什么转账失败?**
A: 检查是否超过单笔限额或余额不足（含税费）。

**Q: 如何查看税费历史?**
A: 在Etherscan查看`TaxCollected`事件。

**Q: Owner可以修改哪些参数?**
A: 税率(≤5%)、税收地址、交易限额、白名单。

**Q: 合约可以升级吗?**
A: 不可以，这是不可变合约，保证去中心化。

更多问题请查看 [部署指南FAQ](docs/deployment-guide.md#faq)。

## License

MIT

## 免责声明

本代币仅用于教育和演示目的。部署和使用需自行承担风险。请遵守当地法规。

---

**Built with ❤️ using Foundry**
