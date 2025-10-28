# Ops Runbook

## Role Provisioning
- 默认管理员：部署初始化调用者 (`DEFAULT_ADMIN_ROLE`)
- 授予管理员：`grantRole(ADMIN_ROLE, account)`
- 授予升级权限：`grantRole(UPGRADE_ROLE, account)`
- 推荐将运营账户与升级账户区分，并记录多签地址

## Reward Funding
1. 确定奖励速率 `fukuaPerBlock` 与时间窗口 (`startBlock`, `endBlock`)
2. 计算总奖励需求：`(endBlock - startBlock) * fukuaPerBlock`
3. 将所需量的 FUKUA 充值到质押合约地址
4. 定期检查 `fukuaToken.balanceOf(stake)`，不足时补充

## Emergency Procedures
- 全局暂停：`pause()` -> 禁用所有入口
- 个别暂停：`pauseWithdraw()` / `pauseClaim()`
- 恢复：`unpause()` / `unpauseWithdraw()` / `unpauseClaim()`
- 在升级或修复漏洞前，建议先全局暂停

## Upgrade Checklist
1. 确认新实现兼容存储布局
2. 部署新逻辑合约并记录地址
3. 使用持有 `UPGRADE_ROLE` 的账户调用 `upgradeTo(newImplementation)`（通过代理脚本或前端）
4. 升级后运行 `massUpdatePools()` 同步奖励快照
5. 重新启用暂停功能（如有）

## Monitoring
- 监听 `Claim`、`Withdraw`、`RequestUnstake` 事件以统计资金流量
- 监控 `Withdraw` 失败情况（通常来自锁定队列未清或熔断）
- 记录 `SetFukuaPerBlock`、`SetPoolWeight` 等参数变更

## Incident Logging
- 每次暂停/恢复、参数调整需记录时间、操作者、原因
- 若遇到奖励余额不足，通知运营补充并记录补充数量
