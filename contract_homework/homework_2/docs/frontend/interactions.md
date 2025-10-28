# Frontend Interaction Checklist

## 用户路径

### 质押 ETH
1. 确认 `pid == 0` 且 `pool[0].stTokenAddress == address(0)`
2. 检查 `msg.value >= minDepositAmount`
3. 调用 `depositEth()` 并监听 `Deposit`
4. 刷新 `stakingBalance`、`pendingFukua`

### 质押 ERC20
1. 调用 `token.approve(stake, amount)`
2. 检查 `amount > minDepositAmount`
3. 调用 `deposit(pid, amount)`
4. 监听 `Deposit` 事件并更新界面

### 领取奖励
1. 调用 `pendingFukua(pid, account)` 计算预估
2. 调用 `claim(pid)`
3. 处理事件 `Claim(user, pid, amount)` 更新代币余额及历史
4. 若余额不足，提示剩余待领取奖励 `pendingFukua - amount`

### 解质押与提现
1. 调用 `unstake(pid, amount)`，监听 `RequestUnstake`
2. 读取 `withdrawAmount(pid, account)` 显示锁定队列与可立即提现金额
3. 在解锁区块到达后调用 `withdraw(pid)`
4. 展示 `Withdraw` 事件及最新余额

### 暂停提示
- `paused()` 为真：禁用全部交互
- `withdrawPaused` 为真：禁用提现
- `claimPaused` 为真：禁用领取奖励

界面需在按钮上实时提示当前状态。

## 管理员路径

### 参数调整
- `setFukuaToken(newToken)`
- `setFukuaPerBlock(newRate)`
- `setStartBlock` / `setEndBlock`
- `addPool(token, weight, minDeposit, lockBlocks, withUpdate)`
- `updatePool(pid, minDeposit, lockBlocks)`
- `setPoolWeight(pid, weight, withUpdate)`

操作流程：
1. 在提交前展示当前值 & 改动值
2. 确认调用账户拥有 `ADMIN_ROLE`
3. 成功后刷新状态并记录事件

### 熔断
- 全局：`pause()` / `unpause()`
- 提现：`pauseWithdraw()` / `unpauseWithdraw()`
- 领取：`pauseClaim()` / `unpauseClaim()`

管理员界面需给予二次确认，并在成功后刷新全局状态。

## 错误提示

| 错误 | 说明 | UI 提示 |
| --- | --- | --- |
| `invalid pid` | PID 越界 | “请选择有效质押池” |
| `deposit amount is too small` | 少于最小质押 | “金额低于最小门槛” |
| `withdraw is paused` | 提现已暂停 | “提现功能暂时关闭” |
| `claim is paused` | 奖励领取已暂停 | “领取功能暂时关闭” |
| `Pausable: paused` | 合约暂停 | “合约暂停中” |
| `AccessControlUnauthorizedAccount` | 权限不足 | “需要管理员权限” |

## 数据刷新建议

- 订阅事件 + 每个新区块轮询关键信息
- 操作成功后先乐观更新界面，再等待链上确认
- 锁定队列可展示 countdown(`unlockBlocks - block.number`)
