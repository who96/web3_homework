# Fukua Stake 项目指南

本文档介绍 Fukua Stake 系统的整体结构、部署流程、合约交互方式，以及前端界面的使用步骤。项目已经部署于 **Sepolia** 测试网，可直接使用下述合约地址进行体验。

## 1. 系统概览

- **FukuaStakeToken (FUKUA)**：奖励代币，遵循 ERC20 标准。
- **FukuaStake**：质押主合约，支持多质押池、按区块发放奖励、解质押锁定队列以及安全熔断。
- **前端控制台**：本地静态页面，直接与合约交互。

### 已部署合约地址（Sepolia）

| 合约 | 地址 |
| ---- | ---- |
| FukuaStake | `0xFF16fD29A0138E432A49ba7A68dE689c41D43239` |
| FukuaStakeToken (FUKUA) | `0x8F18b27F3d91b258a3a9242D2Bd4D9C219EEcE1E` |

部署脚本已自动添加了 **PID 0 的 ETH 池**（权重 100，锁定期 10 区块）并为 FukuaStake 合约预存奖励。

---

## 2. 目录结构

```
contract_homework/homework_2/
├─ contracts/                 # Solidity 合约
│  ├─ FukuaStake.sol
│  └─ FukuaStakeToken.sol
├─ script/
│  └─ Deploy.s.sol            # Foundry 脚本：部署奖励代币和质押合约并初始化配置
├─ test/                      # Foundry 测试（Pause、Rewards、Roles、流程等）
├─ frontend/                  # 前端静态文件
│  ├─ index.html
│  ├─ styles.css
│  └─ app.js
├─ docs/                      # 前端和运维说明
│  ├─ frontend/state-guide.md
│  ├─ frontend/interactions.md
│  └─ operations/runbook.md
├─ start_frontend.sh          # 启动本地前端服务器（自动清理 5173 端口）
├─ .env                       # 私钥 / RPC 配置（请确保安全）
├─ foundry.toml / foundry.lock
└─ README.md                  # 本文件
```

---

## 3. 部署指南

以下步骤默认基于 Foundry 环境，RPC 与私钥由 `.env` 提供。

### 3.1 环境准备

```bash
cd /Users/huluobo/web3_project/web3_homework/contract_homework/homework_2
forge build          # 编译
forge test           # 运行所有测试
```

确保 `.env` 格式如下（示例）：

```env
PRIVATE_KEY=0x...               # 管理员私钥
SEPOLIA_RPC_URL=https://...
```

### 3.2 部署奖励代币与质押合约

项目已经写好脚本 `script/Deploy.s.sol`，它会：
- 部署 FukuaStakeToken
- 部署 FukuaStake
- 注册 PID 0 的 ETH 池
- 将一部分 FUKUA 转入 staking 合约作为奖励池

运行：

```bash
source .env
forge script script/Deploy.s.sol \
  --rpc-url $SEPOLIA_RPC_URL \
  --private-key $PRIVATE_KEY \
  --broadcast
```

执行成功后，合约地址会记录在 `broadcast/Deploy.s.sol/11155111/run-latest.json`。本文前言已列出我们部署好的地址，可直接复用。

### 3.3 其它管理操作（可选）

使用 `cast send` 可进行其它初始化：

- 添加新池（示例为 ERC20 池）：

  ```bash
  cast send $FUKUA_STAKE_ADDRESS \
    "addPool(address,uint256,uint256,uint256,bool)" \
    <stakingToken> 100 0 20 false \
    --rpc-url $SEPOLIA_RPC_URL \
    --private-key $PRIVATE_KEY
  ```

- 为新池补充奖励代币：

  ```bash
  cast send $FUKUA_TOKEN_ADDRESS \
    "transfer(address,uint256)" \
    $FUKUA_STAKE_ADDRESS 1000000000000000000000 \
    --rpc-url $SEPOLIA_RPC_URL \
    --private-key $PRIVATE_KEY
  ```

---

## 4. 合约交互说明

主要函数与行为：

- `depositEth()` / `deposit(pid, amount)`：质押 ETH 或 ERC20。
- `unstake(pid, amount)`：发起解质押（进入锁定队列）。
- `withdraw(pid)`：锁定期满后提取解质押金额。
- `claim(pid)`：领取 FUKUA 奖励。
- `setFukuaPerBlock`、`setPoolWeight`、`pause/unpause` 等需 `ADMIN_ROLE`。
- `upgradeTo` 由 `UPGRADE_ROLE` 控制（脚本已经授予管理员）。

锁定期逻辑：`unstakeLockedBlocks` 定义了等待多少个区块后才能 `withdraw`；可通过 `pool(pid)` 查询。

---

## 5. 前端使用步骤

### 5.1 启动前端

```bash
./start_frontend.sh
```

脚本会自动清理占用 5173 端口的旧进程，并在 `http://localhost:5173` 提供静态网页。浏览器访问即可。

### 5.2 连接与加载合约

1. 点击右上角 **“连接钱包”**，在 MetaMask 中确认。当前仅支持浏览器注入钱包。
2. 在 “合约连接” 面板：
   - 填写 FukuaStake 地址（例如 `0xFF16fD29A0138E432A49ba7A68dE689c41D43239`）。
   - “奖励代币地址（可选）”可留空；加载后会自动填入 `fukuaToken()` 返回值。
3. 点击 **“加载合约”**。成功后：
   - “全局状态 / 质押池 / 我的资产 / 管理员工具”会显示真实数据。
   - “Messages” 区域输出 `[INFO] 已连接 FukuaStake 合约：...`。

### 5.3 普通用户操作

1. **质押 ETH**
   - “选择池 ID” 选 `PID 0`。
   - “质押数量”输入 `0.1`（示例）。
   - 点击 **“质押”**，在钱包确认。
   - Messages 显示“质押成功”， `质押余额`、`总质押` 更新。

2. **领取奖励**
   - 等待几个区块（或切换其它账户触发区块）。
   - 点击 **“领取奖励”**。
   - Messages 显示“领取成功”，`待领取奖励` 归零，钱包获得 FUKUA。

3. **解质押并提现**
   - “解质押数量”输入 `0.05`，点击 **“发起解质押”**。
   - 等待锁定期（PID 0 默认 10 区块）。
   - 点击 **“领取解锁质押”**，Messages 表明提现成功。

### 5.4 管理员功能（需 `ADMIN_ROLE`）

1. **调整每区块奖励**
   - 在“管理员工具”输入新的数值（如 1.5），点击“更新奖励速率”。
   - `fukuaPerBlock` 字段更新，Messages 显示成功。

2. **调整池权重**
   - 输入 PID、权重，并决定是否 `withUpdate`。
   - 点击“更新权重”，表格中 `权重` 改变。

3. **熔断**
   - “全局暂停/取消暂停”：禁用或恢复所有入口。
   - “暂停/恢复领取”“暂停/恢复提现”：仅影响特定功能。

Messages 区域会实时显示执行结果。

---

## 6. 常见问题与排查

- **地址校验失败**：前端会尝试标准化地址，小写 / 大写均可。若仍报错，检查是否录入了多余字符。
- **连接钱包后仍提示“请先加载合约”**：确认已经在合约输入框填入正确地址并点击“加载合约”；检查 `Messages` 是否有失败原因。
- **权限不足**：管理员操作需要 `ADMIN_ROLE`，请切换至授权账户或通过合约授予角色。
- **奖励余额不足导致领取失败**：合约会自动把实际支付额返回给用户，未支付部分保留在 `pendingFukua` 中；管理员需补充奖励代币。

更多前端控件与运维细节请参考 `docs/frontend/*` 与 `docs/operations/runbook.md`。

---

## 7. 附录：测试与脚本

- `forge test -vv`：包含暂停、奖励、锁定队列、角色权限、流程等 14 个用例。
- `script/ShowStatus.s.sol`：可快速读取链上全局状态（设置 `FUKUA_STAKE_ADDRESS` 环境变量后运行）。

示例：

```bash
export FUKUA_STAKE_ADDRESS=0xFF16fD29A0138E432A49ba7A68dE689c41D43239
forge script script/ShowStatus.s.sol --sig run --fork-url $SEPOLIA_RPC_URL
```

---

至此，即可完整地部署、交互并验证 Fukua Stake 项目。若需在其它网络部署，仅需更换 RPC 与私钥配置，运行上述同样的脚本步骤即可。
