## ADDED Requirements

### Requirement: 多池配置
质押平台 MUST 支持多个质押池，并对首个池、质押资产类型以及每池参数施加确定性约束。

#### Scenario: 添加首个 ETH 池
- **GIVEN** 平台当前没有任何质押池
- **WHEN** 管理员调用 `addPool(address(0), weight, minDeposit, lockBlocks, withUpdate)`
- **THEN** 必须创建 PID 0 的池，将 `stTokenAddress` 设为 `address(0)`，`lastRewardBlock` 设为 `max(block.number, startBlock)`，并把 `totalPoolWeight` 增加 `weight`

#### Scenario: 添加 ERC20 池
- **GIVEN** 平台已经存在至少一个质押池
- **WHEN** 管理员以非零质押代币地址且 `unstakeLockedBlocks > 0` 调用 `addPool`
- **THEN** 若质押代币地址为零必须回滚；当 `withUpdate == true` 时必须先执行 `massUpdatePools()`；并且必须触发 `AddPool` 事件记录配置

#### Scenario: 更新池参数
- **WHEN** 管理员修改某池的 `minDepositAmount` 或 `unstakeLockedBlocks`
- **THEN** 新值必须被写入存储，并且触发 `UpdatePoolInfo` 事件携带最新限制

### Requirement: 用户质押
平台 MUST 允许用户在满足池约束的情况下质押 ETH 或 ERC20 资产，并产出准确的会计事件。

#### Scenario: 向 PID 0 质押 ETH
- **GIVEN** PID 0 的 `stTokenAddress == address(0)`
- **WHEN** 用户调用 `depositEth()` 且 `msg.value >= minDepositAmount`
- **THEN** 合约必须增加用户 `stAmount`，更新池总量，结算待发奖励，并以 ETH 数额触发 `Deposit` 事件

#### Scenario: 向非零 PID 质押 ERC20
- **WHEN** 用户在 ERC20 质押池调用 `deposit(pid, amount)`
- **THEN** 必须要求 `pid != 0`、`amount > minDepositAmount`，并通过 `safeTransferFrom` 转入资产，随后经 `_deposit` 更新记账并触发 `Deposit`

#### Scenario: 拒绝不足额质押
- **WHEN** 质押金额小于等于池的最小限制（依据当前不等号规则）
- **THEN** 交易必须以 `deposit amount is too small` 回滚

### Requirement: 奖励累积与领取
平台 MUST 在配置的区块窗口内按区块累积 Fukua 奖励，追踪用户应得额度，并在领取时安全转账。

#### Scenario: 交互时更新池奖励
- **GIVEN** 某池存在质押余额且 `block.number > lastRewardBlock`
- **WHEN** `deposit`、`unstake` 或 `claim` 触发 `updatePool(pid)`
- **THEN** `getMultiplier(lastRewardBlock, block.number)` 必须把区间限制在 `[startBlock, endBlock]`，按 `poolWeight / totalPoolWeight` 计算池奖励，更新 `accFukuaPerShare`，并把 `lastRewardBlock` 设为当前区块

#### Scenario: 计算待领取奖励
- **WHEN** 查询 `pendingFukua(pid, user)`
- **THEN** 必须返回 `user.stAmount * accFukuaPerShare / 1e18 - finishedFukua + pendingFukua`，必要时模拟到指定区块的新增奖励

#### Scenario: 领取奖励
- **WHEN** 用户在 `claimPaused == false` 时调用 `claim(pid)`
- **THEN** 合约必须将待领取额度清零，使用 `_safeFukuaTransfer` 在可用余额范围内转账，刷新 `finishedFukua`，并触发 `Claim` 事件

### Requirement: 解质押锁定队列
平台 MUST 通过排队机制实现解质押锁定，只在到达解锁区块后释放资金。

#### Scenario: 创建解质押请求
- **WHEN** 用户以 `amount <= stAmount` 调用 `unstake(pid, amount)`
- **THEN** 合约必须结算待发奖励，减少 `stAmount`，追加 `UnstakeRequest{amount, block.number + unstakeLockedBlocks}`，并触发 `RequestUnstake`

#### Scenario: 提取已解锁资金
- **WHEN** 用户在一个或多个请求满足 `unlockBlocks <= block.number` 后调用 `withdraw(pid)`
- **THEN** 合约必须按顺序累计已解锁金额，原地压缩队列，依据池资产类型转出 ETH 或 ERC20，并触发 `Withdraw` 事件记录金额

#### Scenario: 保留未解锁请求
- **WHEN** 调用 `withdraw(pid)` 但队首请求尚未解锁
- **THEN** 函数必须保持队列不变且不转出任何代币

### Requirement: 基于角色的管理
平台 MUST 将特权操作限定在 AccessControl 角色之内，并对配置变更发出可审计事件。

- **WHEN** 调用 `setFukuaToken`、`setStartBlock`、`setEndBlock`、`setFukuaPerBlock`、`addPool`、`updatePool`、`setPoolWeight`、`pauseWithdraw`、`unpauseWithdraw`、`pauseClaim` 或 `unpauseClaim`
- **THEN** 调用者必须拥有 `ADMIN_ROLE`，否则应回滚

#### Scenario: 升级授权
- **WHEN** 在 UUPS 升级过程中执行 `_authorizeUpgrade`
- **THEN** 必须要求调用者具备 `UPGRADE_ROLE`，以确保只有批准账户能部署新逻辑

#### Scenario: 事件记录
- **WHEN** 管理员操作修改配置
- **THEN** 必须触发对应事件（`SetFukuaToken`、`SetStartBlock`、`SetEndBlock`、`SetFukuaTokenPerBlock`、`AddPool`、`UpdatePoolInfo`、`SetPoolWeight`、`PauseWithdraw`、`UnpauseWithdraw`、`PauseClaim`、`UnpauseClaim`）并携带更新值

### Requirement: 紧急控制
平台 MUST 同时提供全局与细粒度的暂停机制，以在不破坏账目前提下应对实时事件。

#### Scenario: 全局暂停
- **WHEN** 合约未暂停且管理员调用 `pause()`
- **THEN** `PausableUpgradeable` 必须进入暂停状态，使后续 `depositETH`、`deposit`、`unstake`、`withdraw`、`claim` 因 `whenNotPaused` 而回滚

#### Scenario: 恢复运营
- **WHEN** 事件解除且管理员调用 `unpause()`
- **THEN** 所有流程必须恢复，且用户余额与排队请求不得被篡改

#### Scenario: 提款细粒度暂停
- **WHEN** 调用 `pauseWithdraw()`
- **THEN** `unstake` 必须继续可用，但 `withdraw` 必须以 `withdraw is paused` 回滚，直至调用 `unpauseWithdraw()`

#### Scenario: 领取细粒度暂停
- **WHEN** 调用 `pauseClaim()`
- **THEN** 用户仍可创建解质押请求，但 `claim` 必须以 `claim is paused` 回滚，直到 `unpauseClaim()` 成功

### Requirement: 状态可视化与工具
平台 MUST 暴露足够状态以便客户端与运营方展示池指标、用户持仓和暂停状态。

- **WHEN** 客户端读取 `startBlock`、`endBlock`、`fukuaPerBlock`、`totalPoolWeight`、`claimPaused`、`withdrawPaused`、`paused()`
- **THEN** 合约必须返回反映最新管理员动作的实时值

#### Scenario: 查询池数据
- **WHEN** 客户端检查 `pool[pid]` 与 `poolLength()`
- **THEN** 必须获得质押资产地址、权重、最小质押、锁定区块、累计奖励以及池总质押量，以渲染池卡片

#### Scenario: 查询用户持仓
- **WHEN** 客户端调用 `stakingBalance`、`pendingFukua`、`withdrawAmount`
- **THEN** 合约必须提供当前质押、待领取奖励（含排队奖励）、累计提现请求以及已解锁金额，用于展示用户面板
