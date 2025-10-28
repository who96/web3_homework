# Frontend State Guide

本文件列出了前端需要读取并展示的链上状态，便于界面渲染与数据同步。

## 全局状态

| 字段 | 来源 | 描述 |
| ---- | ---- | ---- |
| `startBlock` | `FukuaStake.startBlock()` | 奖励开始区块 |
| `endBlock` | `FukuaStake.endBlock()` | 奖励结束区块 |
| `fukuaPerBlock` | `FukuaStake.fukuaPerBlock()` | 每区块奖励数量 |
| `totalPoolWeight` | `FukuaStake.totalPoolWeight()` | 所有池权重之和 |
| `claimPaused` | `FukuaStake.claimPaused()` | 领取奖励是否暂停 |
| `withdrawPaused` | `FukuaStake.withdrawPaused()` | 提款是否暂停 |
| `paused()` | `FukuaStake.paused()` | 全局暂停开关 |
| `fukuaToken` | `FukuaStake.fukuaToken()` | 奖励代币地址 |

## 池级状态

1. 调用 `poolLength()` 获取池数量
2. 对每个 `pid` 调用 `pool(pid)` 读取：
   - `stTokenAddress`
   - `poolWeight`
   - `lastRewardBlock`
   - `accFukuaPerShare`
   - `stTokenAmount`
   - `minDepositAmount`
   - `unstakeLockedBlocks`

## 用户状态

- `stakingBalance(pid, user)`：当前质押量
- `pendingFukua(pid, user)`：待领取奖励（含缓存）
- `withdrawAmount(pid, user)`：返回 `requestAmount` 与 `pendingWithdrawAmount`
- `user(pid, user).requests`：锁定队列，可通过多次调用或事件索引展示明细

## 事件订阅建议

| 事件 | 用途 |
| ---- | ---- |
| `Deposit` | 更新池/用户质押显示 |
| `RequestUnstake` | 更新锁定队列 |
| `Withdraw` | 更新余额和历史记录 |
| `Claim` | 更新奖励余额 |
| `SetFukuaPerBlock`/`SetFukuaToken` | 更新全局参数 |
| `Pause`/`Unpause` | 同步全局暂停状态 |
| `PauseWithdraw`/`UnpauseWithdraw` | 同步提款开关 |
| `PauseClaim`/`UnpauseClaim` | 同步领取开关 |
