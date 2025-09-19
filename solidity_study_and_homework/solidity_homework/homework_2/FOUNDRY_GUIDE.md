# Foundry 使用指南

## 安装

1. 运行安装脚本：
```bash
./INSTALL_FOUNDRY.sh
source ~/.bashrc  # 或 source ~/.zshrc
foundryup  # 更新到最新版本
```

## 项目结构

```
├── contracts/       # 合约源码（兼容 Hardhat 结构）
├── test/           # 测试文件（.t.sol）
├── script/         # 部署脚本（.s.sol）
├── out/            # 编译输出
├── foundry.toml    # 配置文件
└── Makefile        # 快捷命令
```

## 常用命令

### 基础操作
```bash
# 编译
forge build
# 或
make build

# 测试
forge test
# 或
make test

# 详细测试输出
forge test -vvvv

# 测试特定函数
forge test --match-test test_MintNFT
```

### 高级功能
```bash
# Gas 报告
make gas

# 测试覆盖率
make coverage

# 模糊测试（自动生成测试输入）
forge test --match-test testFuzz

# 格式化代码
make format
```

### 部署

1. **本地部署**（使用 Anvil）
```bash
# 启动本地节点
anvil

# 另一个终端部署
make deploy-local
```

2. **Sepolia 测试网部署详解**

**步骤1：准备测试环境**
```bash
# 确保 .env 文件配置正确
cat .env
# 应包含：
# PRIVATE_KEY=your_private_key_without_0x
# SEPOLIA_RPC_URL=https://eth-sepolia.public.blastapi.io
# ETHERSCAN_API_KEY=your_etherscan_api_key
```

**步骤2：获取测试ETH**
```bash
# 方法1：使用水龙头
# 访问 https://sepoliafaucet.com/
# 输入你的钱包地址获取测试ETH

# 方法2：检查余额
cast balance <your_address> --rpc-url sepolia
```

**步骤3：部署合约**
```bash
# 部署SimpleNFT合约
make deploy-testnet NETWORK=sepolia

# 或者直接使用forge命令
forge script script/Deploy.s.sol:Deploy --rpc-url sepolia --broadcast --verify

# 部署BeggingContract合约
forge create --rpc-url sepolia --private-key $PRIVATE_KEY contracts/getMoney.sol:BeggingContract --verify
```

**步骤4：合约验证**
```bash
# 如果部署时没有验证，可以单独验证
forge verify-contract <contract_address> contracts/getMoney.sol:BeggingContract --etherscan-api-key $ETHERSCAN_API_KEY --chain sepolia
```

**步骤5：合约交互测试**
```bash
# 调用只读函数（以BeggingContract为例）
cast call <contract_address> "owner()" --rpc-url sepolia

# 调用写函数（需要在工作时间UTC+8 9-18点）
cast send <contract_address> "donate()" --value 0.1ether --rpc-url sepolia --private-key $PRIVATE_KEY

# 查询捐赠金额
cast call <contract_address> "getDonation(address)" <your_address> --rpc-url sepolia

# 查询排行榜
cast call <contract_address> "getTopDonors()" --rpc-url sepolia
```

**步骤6：部署验证checklist**
- ✅ 合约地址已生成
- ✅ 在Etherscan上已验证
- ✅ 基本函数调用成功
- ✅ 事件正确触发
- ✅ 权限控制正常工作

3. **主网**（谨慎！）
```bash
make deploy-mainnet
```

## 环境变量

创建 `.env` 文件：
```env
# 私钥（不含 0x）
PRIVATE_KEY=your_private_key_here

# RPC URLs
SEPOLIA_RPC_URL=https://eth-sepolia.public.blastapi.io
MAINNET_RPC_URL=https://eth.llamarpc.com

# Etherscan API（用于验证）
ETHERSCAN_API_KEY=your_key_here
```

## Foundry vs Hardhat

### 优势
- ✅ **速度快**：编译和测试快 10-100 倍
- ✅ **原生 Solidity**：测试用 Solidity 写，不用学 JavaScript
- ✅ **模糊测试**：自动生成边界测试用例
- ✅ **简单直接**：没有复杂的插件系统
- ✅ **更好的错误信息**：调试更容易

### 测试对比

**Hardhat (JavaScript)**:
```javascript
it("should mint NFT", async () => {
    const tx = await nft.mintNFT(user, uri);
    expect(await nft.ownerOf(0)).to.equal(user);
});
```

**Foundry (Solidity)**:
```solidity
function test_MintNFT() public {
    uint256 tokenId = nft.mintNFT(user, uri);
    assertEq(nft.ownerOf(tokenId), user);
}
```

## 调试技巧

1. **使用 console.log**
```solidity
import "forge-std/console.sol";

function test_Debug() public {
    console.log("Value:", someValue);
    console.log("Address:", someAddress);
}
```

2. **查看调用栈**
```bash
forge test -vvvv  # 最详细的输出
```

3. **Fork 主网测试**
```bash
# Fork 主网状态进行测试
forge test --fork-url mainnet
```

## Cheatcodes（作弊码）

Foundry 提供强大的测试功能：

```solidity
// 模拟用户
vm.prank(alice);
nft.mintNFT(alice, "uri");

// 设置区块时间
vm.warp(block.timestamp + 1 days);

// 期望回滚
vm.expectRevert("Not owner");
nft.burn(tokenId);

// 修改余额
vm.deal(alice, 100 ether);

// 模拟调用
vm.mockCall(address, abi.encodeWithSelector(...), returnData);
```

## Sepolia部署脚本示例

**创建BeggingContract部署脚本** `script/DeployBeggar.s.sol`:
```solidity
// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "forge-std/Script.sol";
import "../contracts/getMoney.sol";

contract DeployBeggar is Script {
    function run() external {
        uint256 deployerPrivateKey = vm.envUint("PRIVATE_KEY");

        vm.startBroadcast(deployerPrivateKey);

        BeggingContract beggar = new BeggingContract();
        console.log("BeggingContract deployed to:", address(beggar));
        console.log("Owner:", beggar.owner());

        vm.stopBroadcast();
    }
}
```

**部署命令**:
```bash
forge script script/DeployBeggar.s.sol:DeployBeggar --rpc-url sepolia --broadcast --verify
```

## 快速测试脚本

**创建测试脚本** `test_sepolia.sh`:
```bash
#!/bin/bash

CONTRACT_ADDRESS="your_contract_address_here"
RPC_URL="sepolia"

echo "=== 测试BeggingContract ==="
echo "1. 检查owner:"
cast call $CONTRACT_ADDRESS "owner()" --rpc-url $RPC_URL

echo "2. 检查当前UTC+8时间是否在工作时间:"
TIMESTAMP=$(date +%s)
echo "当前时间戳: $TIMESTAMP"

echo "3. 尝试捐赠 0.01 ETH (如果在工作时间):"
cast send $CONTRACT_ADDRESS "donate()" --value 0.01ether --rpc-url $RPC_URL --private-key $PRIVATE_KEY

echo "4. 查询我的捐赠:"
YOUR_ADDRESS=$(cast wallet address --private-key $PRIVATE_KEY)
cast call $CONTRACT_ADDRESS "getDonation(address)" $YOUR_ADDRESS --rpc-url $RPC_URL

echo "5. 查询排行榜:"
cast call $CONTRACT_ADDRESS "getTopDonors()" --rpc-url $RPC_URL
```

## 常见问题

1. **找不到 forge 命令**
   - 运行 `source ~/.bashrc` 或重启终端
   - 确保 `~/.foundry/bin` 在 PATH 中

2. **编译错误**
   - 检查 `foundry.toml` 中的 solc 版本
   - 运行 `forge clean` 清理缓存

3. **测试失败**
   - 使用 `-vvvv` 查看详细信息
   - 检查 `vm.prank` 和权限问题

4. **Sepolia部署问题**
   - **余额不足**: 去 https://sepoliafaucet.com/ 获取测试ETH
   - **私钥格式错误**: 确保不包含 `0x` 前缀
   - **RPC超时**: 换用其他公共RPC或Alchemy/Infura
   - **Gas估算失败**: 手动设置 `--gas-limit` 参数

5. **合约验证失败**
   - **API Key无效**: 检查Etherscan API Key
   - **源码不匹配**: 确保编译器版本一致
   - **依赖库问题**: 使用 `--constructor-args` 指定参数

6. **时间相关错误(BeggingContract)**
   - **"Donation only 9:00-18:00 UTC+8"**: 当前不在工作时间
   - **解决方法**: 等待工作时间或修改合约时间逻辑用于测试

## 监控和调试

**查看交易详情**:
```bash
# 查看交易receipt
cast receipt <tx_hash> --rpc-url sepolia

# 查看交易详细信息
cast tx <tx_hash> --rpc-url sepolia

# 查看合约代码
cast code <contract_address> --rpc-url sepolia

# 查看存储槽
cast storage <contract_address> <slot> --rpc-url sepolia
```

**实时监控事件**:
```bash
# 监控Donation事件
cast logs --rpc-url sepolia --address <contract_address> --decode-logs "Donation(address indexed,uint256)"
```

## 总结

Foundry 是现代 Solidity 开发的最佳选择：
- 更快的开发体验
- 更强大的测试能力
- 更简单的工具链
- 更接近 EVM 本质

告别 JavaScript，拥抱 Solidity！