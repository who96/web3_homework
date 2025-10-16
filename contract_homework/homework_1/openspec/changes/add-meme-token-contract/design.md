# 设计文档: Meme代币合约

## Context（背景）
实现一个SHIB风格的Meme代币，需要在标准ERC20基础上添加交易税和防操纵机制。

**约束:**
- 必须保持ERC20接口兼容性（DEX要求）
- Gas成本不能过高（每笔转账应<200k gas）
- 部署后不可变（不使用代理模式）

**利益相关者:**
- 代币持有者：希望低gas费、无隐藏税费
- 流动性提供者：希望添加/移除流动性时免税
- 项目方：希望通过税收支持营销和开发

## Goals / Non-Goals（目标与非目标）

**Goals:**
- 实现可配置的交易税（1-5%范围）
- 白名单机制，O(1)查询复杂度
- 单笔交易额度限制
- 100% ERC20标准兼容
- 清晰的NatSpec文档

**Non-Goals:**
- 链上每日交易次数追踪（成本太高）
- 内置流动性池操作（DEX已提供）
- 合约升级能力（保持简单）
- 多链部署（先专注以太坊）

## Decisions（关键决策）

### 决策1: 税收逻辑在`_update()`中实现
**为什么:**
- OpenZeppelin ERC20的`_update()`是所有转账的统一入口点
- 避免在`transfer()`, `transferFrom()`, `mint()`, `burn()`中重复代码
- 更容易维护和审计

**替代方案:**
- ❌ 重写`transfer()`和`transferFrom()` - 代码重复，容易出错
- ❌ 使用hook模式 - 增加复杂度，没有实际好处

**Trade-offs:**
- ✅ 优点: 简洁、统一、不易出错
- ⚠️ 缺点: 需要仔细处理铸币/销毁时的税收逻辑（这些操作应该免税）

### 决策2: 白名单用mapping而非数组
**为什么:**
- mapping查询是O(1)，数组是O(n)
- 每笔转账都要检查白名单，必须高效
- 添加/删除白名单地址更简单

**数据结构:**
```solidity
mapping(address => bool) private _isExemptFromTax;
```

**替代方案:**
- ❌ 数组 - O(n)查询，gas爆炸
- ❌ EnumerableSet - 增加复杂度，没有必要（不需要枚举）

### 决策3: 税率用basis points (10000 = 100%)
**为什么:**
- 避免浮点数问题
- 支持精确到0.01%的税率（如250 = 2.5%）
- 标准DeFi做法，与Uniswap等一致

**实现:**
```solidity
uint256 public constant MAX_TAX_RATE = 500; // 5%
uint256 public constant DENOMINATOR = 10000;
uint256 private _taxRate; // e.g., 300 = 3%

// 税费计算
uint256 taxAmount = (amount * _taxRate) / DENOMINATOR;
```

**替代方案:**
- ❌ 百分比（100 = 100%）- 精度不够
- ❌ 直接小数 - Solidity不支持浮点数

### 决策4: 交易限制只用单笔最大额度
**为什么:**
- 每日交易次数需要存储每个地址的交易历史，gas成本爆炸
- 用户可以通过多个地址绕过次数限制
- 单笔最大额度已经足够防止大额抛售

**实现:**
```solidity
uint256 public maxTransactionAmount; // e.g., 1% of total supply

// 在_update()中检查
require(amount <= maxTransactionAmount || _isExemptFromLimit[from], "Exceeds max tx");
```

**替代方案:**
- ❌ 每日次数限制 - 成本高、容易绕过
- ❌ 滑动窗口限制 - 复杂度高、gas贵

## Architecture（架构）

### 合约继承结构
```
MemeToken
  ├─ ERC20 (OpenZeppelin)
  └─ Ownable (OpenZeppelin)
```

**为什么不用ReentrancyGuard:**
- 税费转账的目标地址是受信任的（owner控制的钱包）
- ERC20标准转账没有回调钩子（除非目标是恶意合约）
- 增加不必要的gas成本
- 如果后续发现需要，可以在特定函数上添加

### 状态变量设计
```solidity
// 税收相关
uint256 private _taxRate;              // 当前税率 (basis points)
address private _taxRecipient;          // 税收接收地址
mapping(address => bool) private _isExemptFromTax;  // 税收豁免白名单

// 交易限制
uint256 public maxTransactionAmount;    // 单笔最大额度
mapping(address => bool) private _isExemptFromLimit; // 限制豁免白名单

// 常量
uint256 public constant MAX_TAX_RATE = 500;    // 5%
uint256 public constant DENOMINATOR = 10000;
```

**存储优化:**
- uint256占满32字节槽位，不需要packing
- mapping单独占槽位，无法优化
- 常量不占用storage

### 核心逻辑流程
```
用户调用 transfer(to, amount)
  ↓
ERC20.transfer() 调用 _update(from, to, amount)
  ↓
MemeToken._update() 重写:
  1. 检查交易限制 (if !_isExemptFromLimit[from])
  2. 检查是否需要收税 (if !_isExemptFromTax[from] && !_isExemptFromTax[to])
  3. 如果需要收税:
     - 计算税费: taxAmount = amount * _taxRate / DENOMINATOR
     - 实际转账: amount - taxAmount
     - 转税费给_taxRecipient: taxAmount
  4. 调用 super._update() 完成转账
```

## Risks / Trade-offs（风险与权衡）

### 风险1: 税收逻辑增加gas成本
**影响:** 每笔转账需要额外计算和转账税费
**缓解:**
- 白名单豁免owner和LP，他们的高频操作不受影响
- 使用basis points避免复杂计算
- Foundry gas基准测试，确保<200k gas

### 风险2: Owner权力过大
**影响:** Owner可以随时修改税率、白名单、交易限制
**缓解:**
- 设置MAX_TAX_RATE上限（5%）
- 文档中明确说明后续会转移到多签或DAO
- 提供`renounceOwnership()`选项

### 风险3: 白名单配置错误
**影响:** 如果忘记把Uniswap Pair加入白名单，LP操作会失败
**缓解:**
- 部署脚本中自动添加关键地址到白名单
- 测试中覆盖白名单场景
- 文档中详细说明白名单设置步骤

### 风险4: ERC20兼容性问题
**影响:** 如果税收逻辑破坏标准，无法在DEX上交易
**缓解:**
- 严格遵循OpenZeppelin ERC20实现
- 只重写`_update()`，不改变公开接口
- 使用Uniswap forked mainnet测试

## Migration Plan（迁移计划）
N/A - 这是全新合约，不涉及迁移

## Deployment Checklist（部署检查清单）
1. ✅ Foundry测试全部通过（包括fuzz tests）
2. ✅ Slither静态分析无高危漏洞
3. ✅ Gas基准测试: 普通转账 < 100k gas, 含税转账 < 200k gas
4. ✅ 在测试网部署并验证:
   - 普通转账工作正常
   - 税费计算准确
   - 白名单豁免生效
   - Uniswap添加流动性成功
5. ✅ 合约验证（Etherscan）
6. ✅ 文档完整（NatSpec + deployment guide）

## Open Questions（待解决问题）
- [x] 税费接收地址是单个地址还是需要分配到多个地址（营销/开发/销毁）？
  - **决定:** 单个地址（.env中的WALLET1），保持简单
- [x] 是否需要"买入税"和"卖出税"分别设置？
  - **决定:** 不需要，统一税率更简单透明，避免被标记为蜜罐
- [x] 初始供应量和分配方案？
  - **决定:**
    - 代币名称: Fukua (FK)
    - 总供应量: 21,000,000 FK
    - 流动性池: 210,000 FK (1%) + 0.1 ETH → Uniswap V2
    - WALLET1: 20,790,000 FK (99%)