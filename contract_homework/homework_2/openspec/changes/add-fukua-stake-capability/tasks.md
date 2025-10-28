## 0. 基础准备
- [x] 0.1 将 `Advanced2-contract-stake/stake-contract/contracts/MetaNode*.sol` 迁移到 `contract_homework/homework_2/contracts/`
- [x] 0.2 将迁移的合约/文件名统一为 `FukuaStakeToken.sol`、`FukuaStake.sol` 等以反映项目命名，并同步更新 pragma import/引用
- [x] 0.3 基于 Foundry 初始化项目结构（`foundry.toml`、`lib/`、`script/`、`test/`、`Makefile`/`package.json` 等）并安装 OZ 依赖
- [x] 0.4 调整本地脚本与 .env 引用路径，确认在 `homework_2` 根目录能执行 `forge build`、`forge test`、`forge script`

## 1. 规范落地
- [x] 1.1 对照 `project.md` 与 `specs/staking-platform/spec.md`，确认所有字段/事件被文档覆盖
- [x] 1.2 若存在遗漏（例如前端展示字段、暂停语义），在 proposal/spec 中追加补丁并重新验证
- [x] 1.3 执行 `openspec validate add-fukua-stake-capability --strict` 记录校验结果

## 2. 智能合约实现
- [x] 2.1 在迁移后的 `FukuaStake.sol` 中实现全局 `pause()/unpause()` 并保证所有入口受 `whenNotPaused` 控制
  - [x] 2.1.1 编写暂停/恢复相关测试（含未授权调用失败、暂停后入口报错）
- [x] 2.2 调整奖励、权重、锁定提现逻辑，确保 `massUpdatePools`、`pendingMetaNode` 与事件遵循规范
  - [x] 2.2.1 为奖励计算与权重调整编写正向/边界测试（含 `getMultiplier` 区间裁剪、`poolWeight` 变更）
- [x] 2.3 校验 AccessControl 角色：`DEFAULT_ADMIN_ROLE`、`ADMIN_ROLE`、`UPGRADE_ROLE` 分工明确并覆盖所有管理员方法
  - [x] 2.3.1 添加角色控制测试，验证未授权调用会回滚、授权路径可执行
- [x] 2.4 覆盖安全转账、余额不足、锁定队列边界等防御逻辑（含 `_safeFukuaTransfer`、`_safeEthTransfer`）
  - [x] 2.4.1 编写异常场景测试：奖励余额不足、锁定队列顺序、部分解质押
- [x] 2.5 使用 Foundry 编写用户流程测试：质押/解质押/提现、奖励领取、暂停开关联动
- [ ] 2.6 补充部署或运营脚本（如 `script/Deploy.s.sol`）并提供脚本级 smoke test（如 `forge script --dry-run`）
- [ ] 2.7 持续运行 `forge test -vvv`、`forge fmt`、`forge build` 并根据需要记录 gas 报告

## 3. 前端与运维
- [x] 3.1 前端/脚本获取全局状态（`paused`、`claimPaused`、`withdrawPaused`、奖励参数）并渲染到 UI
- [x] 3.2 实现质押、解质押、提现、领取奖励交互，并展示锁定队列与倒计时
- [x] 3.3 管理面板支持池增删、参数调整、暂停控制，并在界面上基于 AccessControl 做权限判断
- [x] 3.4 撰写部署与运营手册：奖励代币预充、角色授予、熔断流程、升级步骤
