# Meme代币规格说明

## ADDED Requirements

### Requirement: ERC20基础功能
合约 SHALL 实现完整的ERC20标准接口，包括`transfer`, `transferFrom`, `approve`, `balanceOf`, `totalSupply`等函数。

#### Scenario: 标准转账成功
- **WHEN** 用户调用`transfer(recipient, amount)`且余额充足
- **THEN** `amount`代币从发送者转移到接收者
- **AND** 触发`Transfer(from, to, amount)`事件

#### Scenario: 授权转账成功
- **WHEN** 用户A授权用户B额度，用户B调用`transferFrom(A, C, amount)`
- **THEN** `amount`代币从A转移到C
- **AND** B的授权额度减少`amount`
- **AND** 触发`Transfer(A, C, amount)`事件

#### Scenario: 余额不足时revert
- **WHEN** 用户余额少于转账金额
- **THEN** 交易revert，提示"ERC20: transfer amount exceeds balance"

### Requirement: 交易税机制
合约 SHALL 对每笔代币转账征收可配置的税费，并将税费转移到指定的税收接收地址。

#### Scenario: 普通转账收取税费
- **WHEN** 用户A转账1000代币给用户B，税率为3%（300 basis points）
- **THEN** 用户A减少1000代币
- **AND** 用户B收到970代币（1000 - 30税费）
- **AND** 税收地址收到30代币
- **AND** 触发`TaxCollected(A, 30)`事件
- **AND** 触发两个`Transfer`事件: `Transfer(A, B, 970)`和`Transfer(A, taxRecipient, 30)`

#### Scenario: 白名单地址免税
- **WHEN** 白名单地址A转账1000代币给用户B
- **THEN** 用户B收到完整的1000代币
- **AND** 税收地址余额不变
- **AND** 不触发`TaxCollected`事件

#### Scenario: 转账给白名单地址免税
- **WHEN** 用户A转账1000代币给白名单地址B
- **THEN** 白名单地址B收到完整的1000代币
- **AND** 税收地址余额不变

#### Scenario: 税率为零时不收税
- **WHEN** 税率设置为0
- **THEN** 所有转账不收取税费，完整金额转移

#### Scenario: Owner修改税率成功
- **WHEN** Owner调用`setTaxRate(400)`（4%）
- **THEN** 税率更新为400 basis points
- **AND** 触发`TaxRateUpdated(oldRate, 400)`事件

#### Scenario: 非Owner修改税率失败
- **WHEN** 非Owner地址调用`setTaxRate(400)`
- **THEN** 交易revert，提示"Ownable: caller is not the owner"

#### Scenario: 税率超过上限时revert
- **WHEN** Owner尝试设置税率为600（6%），超过MAX_TAX_RATE（500）
- **THEN** 交易revert，提示"Tax rate exceeds maximum"

### Requirement: 白名单管理
合约 SHALL 提供白名单机制，允许Owner豁免特定地址的交易税。

#### Scenario: 添加地址到税收白名单
- **WHEN** Owner调用`setTaxExempt(addressA, true)`
- **THEN** `isExemptFromTax(addressA)`返回`true`
- **AND** addressA的后续转账免税

#### Scenario: 移除地址出税收白名单
- **WHEN** Owner调用`setTaxExempt(addressA, false)`
- **THEN** `isExemptFromTax(addressA)`返回`false`
- **AND** addressA的后续转账需要收税

#### Scenario: 合约部署时Owner自动加入白名单
- **WHEN** 合约部署完成
- **THEN** `isExemptFromTax(owner)`返回`true`
- **AND** `isExemptFromTax(address(this))`返回`true`（合约自己）

#### Scenario: 非Owner无法管理白名单
- **WHEN** 非Owner地址调用`setTaxExempt(addressA, true)`
- **THEN** 交易revert，提示"Ownable: caller is not the owner"

### Requirement: 交易限制
合约 SHALL 限制单笔交易的最大金额，防止大额抛售操纵市场。白名单地址不受限制。

#### Scenario: 转账金额在限制内成功
- **WHEN** 用户转账金额 <= `maxTransactionAmount`
- **THEN** 转账成功执行（含税费逻辑）

#### Scenario: 转账金额超过限制时revert
- **WHEN** 用户转账金额 > `maxTransactionAmount`
- **AND** 用户不在限制豁免白名单中
- **THEN** 交易revert，提示"Transfer amount exceeds max transaction amount"

#### Scenario: 白名单地址不受交易限制
- **WHEN** 白名单地址转账金额 > `maxTransactionAmount`
- **THEN** 转账成功执行

#### Scenario: Owner修改最大交易额度
- **WHEN** Owner调用`setMaxTransactionAmount(newAmount)`
- **THEN** `maxTransactionAmount`更新为`newAmount`
- **AND** 触发`MaxTransactionAmountUpdated(oldAmount, newAmount)`事件

#### Scenario: 非Owner无法修改交易限制
- **WHEN** 非Owner地址调用`setMaxTransactionAmount(newAmount)`
- **THEN** 交易revert，提示"Ownable: caller is not the owner"

### Requirement: 税费计算精度
合约 SHALL 使用basis points（10000 = 100%）计算税费，确保精度和一致性。

#### Scenario: 税费计算精确到0.01%
- **WHEN** 税率为250 basis points（2.5%），转账1000代币
- **THEN** 税费为25代币（1000 * 250 / 10000）
- **AND** 接收者收到975代币

#### Scenario: 小额转账税费舍入处理
- **WHEN** 税率为300（3%），转账10代币
- **THEN** 税费为0代币（10 * 300 / 10000 = 0，Solidity整数除法向下舍入）
- **AND** 接收者收到10代币（小额转账实际免税）

#### Scenario: 查询税费计算结果
- **WHEN** 调用`calculateTax(1000)`，税率为300
- **THEN** 返回30代币

### Requirement: 税收接收地址管理
合约 SHALL 允许Owner配置和修改税收接收地址。

#### Scenario: Owner修改税收接收地址
- **WHEN** Owner调用`setTaxRecipient(newAddress)`
- **THEN** 税收接收地址更新为`newAddress`
- **AND** 触发`TaxRecipientUpdated(oldAddress, newAddress)`事件

#### Scenario: 后续转账税费发送到新地址
- **WHEN** 税收地址修改后，用户进行转账
- **THEN** 税费发送到新的税收接收地址

#### Scenario: 税收地址不能为零地址
- **WHEN** Owner尝试设置税收地址为`address(0)`
- **THEN** 交易revert，提示"Tax recipient cannot be zero address"

### Requirement: Gas效率
合约的转账操作 SHALL 保持合理的gas成本，不影响用户体验。

#### Scenario: 普通转账（无税）gas成本低于100k
- **WHEN** 白名单地址之间转账
- **THEN** gas成本 < 100,000

#### Scenario: 含税转账gas成本低于200k
- **WHEN** 普通地址之间转账，触发税费逻辑
- **THEN** gas成本 < 200,000

### Requirement: 事件记录
合约 SHALL 为关键操作触发事件，便于链下索引和监控。

#### Scenario: 税费收取触发事件
- **WHEN** 转账收取税费
- **THEN** 触发`TaxCollected(address indexed from, uint256 amount)`事件

#### Scenario: 参数修改触发事件
- **WHEN** Owner修改税率、税收地址或交易限制
- **THEN** 触发对应的事件（`TaxRateUpdated`, `TaxRecipientUpdated`, `MaxTransactionAmountUpdated`）

### Requirement: 文档完整性
合约代码 SHALL 包含完整的NatSpec注释，解释每个函数和变量的作用。

#### Scenario: 所有public函数有@notice注释
- **WHEN** 查看合约源代码
- **THEN** 所有public/external函数包含`@notice`说明用途

#### Scenario: 复杂逻辑有@dev注释
- **WHEN** 查看`_update()`等核心函数
- **THEN** 包含`@dev`注释解释实现细节

#### Scenario: 参数和返回值有完整说明
- **WHEN** 查看函数签名
- **THEN** 所有参数有`@param`说明，返回值有`@return`说明

### Requirement: Uniswap兼容性
合约 SHALL 与Uniswap V2协议完全兼容，支持流动性操作。

#### Scenario: 添加流动性成功
- **WHEN** Owner通过Uniswap Router添加ETH-MemeToken流动性
- **AND** Uniswap Pair地址已加入白名单
- **THEN** 流动性添加成功，不收取税费
- **AND** Owner收到LP代币

#### Scenario: Swap交易收取税费
- **WHEN** 用户通过Uniswap将ETH换成MemeToken
- **THEN** MemeToken从Pair转移到用户时收取税费
- **AND** 用户收到扣税后的代币

#### Scenario: 移除流动性免税
- **WHEN** LP持有者移除流动性，销毁LP代币
- **AND** Pair地址在白名单中
- **THEN** MemeToken从Pair转移到LP持有者时免税
