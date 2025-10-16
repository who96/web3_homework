# 项目上下文

## 项目目的
在以太坊上实现一个SHIB风格的Meme代币智能合约，包含交易税机制、流动性管理支持和防操纵保护机制。

**核心目标:**
- 创建一个内置交易税的ERC20代币（每笔转账收取手续费）
- 与DEX流动性池无缝兼容（Uniswap V2/V3）
- 通过交易限制防止市场操纵
- 提供完整的部署和使用文档

## 技术栈
- **开发语言:** Solidity ^0.8.20
- **开发框架:** Foundry（开发、测试、部署）
- **基础合约:** OpenZeppelin Contracts（ERC20, Ownable, ReentrancyGuard）
- **DEX集成:** Uniswap V2 Router接口
- **测试框架:** Foundry (forge test)
- **网络:** 以太坊主网 / 测试网（Sepolia, Goerli）

## 项目规范

### 代码风格
- **Solidity风格指南:** 遵循官方Solidity风格指南和OpenZeppelin模式
- **命名规范:**
  - 函数: camelCase (`transferWithTax`, `setTaxRate`)
  - 状态变量: camelCase，内部变量加下划线前缀 (`_taxRate`, `_maxTxAmount`)
  - 常量: UPPER_SNAKE_CASE (`MAX_TAX_RATE`, `DENOMINATOR`)
  - 事件: PascalCase (`TaxCollected`, `LimitUpdated`)
- **文档注释:** 所有public/external函数必须有NatSpec注释 (@notice, @dev, @param, @return)
- **行长度:** 最大120字符
- **缩进:** 4个空格（不使用tab）

### 架构模式
- **继承:** 优先使用OpenZeppelin的成熟合约，不要自己造轮子
- **访问控制:** 使用Ownable管理admin函数，后续升级到多签钱包
- **重入保护:** 对所有进行外部调用的函数应用ReentrancyGuard
- **Gas优化:** 最小化storage写入，使用事件做链下索引
- **简洁优先:** 避免不必要的抽象层。从单文件实现开始。

**关键设计决策:**
1. **税收机制:** 在 `_update()` 重写中实现，不在单独的transfer函数中
2. **白名单模式:** 使用mapping实现O(1)豁免检查（owner、流动性池、合约本身）
3. **不做链上每日限制:** 追踪每日交易次数成本太高，只使用单笔最大额度限制
4. **DEX集成:** 合约是标准ERC20，流动性操作由外部DEX router处理

### 测试策略
- **单元测试:** 100%覆盖所有税收逻辑、限制和边界情况
- **集成测试:** 使用forked主网测试真实Uniswap合约
- **必须包含的测试用例:**
  - 税费计算精度
  - 白名单豁免（owner、LP、合约）
  - 交易限制执行
  - 边界情况: 零额转账、自己转给自己、max uint256
- **Gas基准测试:** 测量和记录常用操作的gas成本
- **安全分析:** 部署前运行Slither/Mythril静态分析

### Git工作流
- **分支策略:**
  - `main` - 生产就绪代码
  - `develop` - 集成分支
  - 功能分支: `feature/tax-mechanism`, `feature/tx-limits`
- **提交规范:** Conventional Commits格式
  - `feat: 添加交易税机制`
  - `fix: 防止白名单地址被收税`
  - `docs: 添加部署指南`
  - `test: 添加零额转账边界测试`
- **PR要求:** 测试通过、包含gas基准、NatSpec注释完整

## 领域知识

### Meme代币经济学
- **交易税:** 通常每笔转账1-10%，分配给:
  - 营销钱包
  - 开发基金
  - 销毁地址（通缩机制）
  - 流动性供应
- **流动性锁定:** 锁定初始流动性是建立信誉的标准做法
- **反鲸鱼机制:**
  - 最大交易规模（如总量的1%）
  - 最大钱包规模（可选，本实现不包含）

### DEX机制
- **Uniswap V2:** 自动做市商（AMM），使用恒定乘积公式 (x * y = k)
- **流动性供应:** 用户存入等值的ETH和代币以获得LP代币
- **税费考虑:** LP操作（添加/移除流动性）必须免税，否则会破坏数学公式
- **Router vs Pair:**
  - Router: 面向用户的合约，用于交换和流动性操作
  - Pair: 实际持有代币的池子（这个地址需要加入白名单）

### 安全考虑
- **重入攻击:** 税费转账可能触发恶意代币的回调
- **整数溢出:** Solidity 0.8+内置检查，但仍需验证税费计算
- **中心化风险:** Owner可以修改税率和限制；计划放弃所有权或改用多签
- **蜜罐风险:** 避免阻止卖出的模式（如买卖税不对称且不披露）

## 重要约束

### 技术约束
- **ERC20合规:** 必须严格遵循ERC20标准（IERC20接口）以兼容DEX
- **Gas限制:** 区块gas限制约30M；复杂转账逻辑必须保持在每笔交易1M gas以下
- **存储成本:** 每个存储槽（32字节）初始化需20k gas；最小化状态变量
- **不可变性:** 部署后合约代码无法更改（除非使用代理模式，我们避免这样做）

### 业务约束
- **税率范围:** 税率必须合理（建议最高1-5%）以避免用户反感
- **无后门:** 不能有隐藏铸币、不能暂停转账（初始发布期除外）
- **透明度:** 所有参数（税率、限制）必须公开可读
- **公平发射:** 不能有未披露的团队预挖

### 监管约束
- **非投资建议:** 代码和文档必须包含免责声明
- **无证券声明:** 避免暗示投资回报的语言
- **合规性:** 注意当地关于代币发行的法规

## 外部依赖

### 必需的库
- **OpenZeppelin Contracts:** `@openzeppelin/contracts` v5.0.0+
  - `ERC20.sol` - 基础代币实现
  - `Ownable.sol` - 访问控制
  - `ReentrancyGuard.sol` - 重入保护（如需要）
- **Foundry:** 最新版本
  - `forge` - 测试和编译
  - `anvil` - 本地节点
  - `cast` - 链交互工具
  - `forge-std` - 标准测试库

### 外部合约（仅接口）
- **Uniswap V2 Router:** `0x7a250d5630B4cF539739dF2C5dAcb4c659F2488D` (以太坊主网)
- **Uniswap V2 Factory:** 用于编程方式计算pair地址
- **WETH:** `0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2`

### 开发工具
- **Anvil:** Foundry本地区块链，用于测试
- **Tenderly/Etherscan:** 合约验证和监控
- **Slither:** 安全漏洞静态分析
- **Remix IDE:** 快速原型和测试（可选）

### API（可选）
- **Etherscan API:** 用于合约验证
- **CoinGecko/CoinMarketCap:** 用于价格追踪（发布后）
- **TheGraph:** 用于索引转账事件和税费收集（可选）
