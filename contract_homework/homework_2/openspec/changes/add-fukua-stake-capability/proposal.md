# 变更提案: 添加 Fukua Stake 能力规范

## Why
当前项目只有散落的产品描述和合约原型，缺乏系统化的 OpenSpec 约束。为了后续按规范实现/迭代，需要将质押池配置、奖励发放、解质押流程、紧急熔断以及前端展示等关键行为写成正式规范，避免开发阶段反复猜测。

## What Changes
- 建立 `staking-platform` 能力规范，覆盖多池质押、奖励分配、解质押锁定与提款、奖励领取、角色权限及紧急控制。
- 记录管理员参数调整（起止区块、每区块产出、池权重、暂停开关）的行为边界与事件要求。
- 定义前端/运营所需的全局状态展示与角色检测，确保产品方案与链上能力一致。
- 创建实现清单与设计摘要，为后续代码实现/回溯提供依据。

## Impact
- 受影响规格: `staking-platform`
- 受影响代码: `contracts/MetaNodeStake.sol`, `contracts/MetaNode.sol`, 后续前端与脚本
- 依赖: Foundry 测试、AccessControl 角色管理、UUPS 升级流程
