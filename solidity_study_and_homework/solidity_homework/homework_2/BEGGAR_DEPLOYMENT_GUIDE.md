# BeggingContract Sepolia 部署测试指南

## 快速开始

### 1. 环境准备
```bash
# 确保环境变量配置
cat .env
# 应包含:
# PRIVATE_KEY=your_private_key_without_0x
# ETHERSCAN_API_KEY=your_etherscan_key

# 获取测试ETH
# 访问 https://sepoliafaucet.com/
```

### 2. 部署合约
```bash
# 方法1: 使用脚本部署 (推荐)
make deploy-beggar

# 方法2: 快速部署
make deploy-beggar-fast

# 方法3: 直接使用forge命令
forge script script/DeployBeggar.s.sol:DeployBeggar --rpc-url sepolia --broadcast --verify
```

### 3. 测试合约
```bash
# 设置合约地址环境变量
export CONTRACT_ADDRESS=0x1234...

# 运行交互测试
make test-beggar-sepolia

# 或直接运行脚本
./test_sepolia_beggar.sh 0x1234...

# 查看合约基本信息
make beggar-info
```

## 功能测试

### 时间限制测试
BeggingContract只允许在UTC+8的9:00-18:00时间段内捐赠：

```bash
# 检查当前时间是否在工作时间
date -u +%H | awk '{print ($1+8)%24}'

# 工作时间内测试捐赠
cast send $CONTRACT_ADDRESS "donate()" --value 0.01ether --rpc-url sepolia --private-key $PRIVATE_KEY

# 工作时间外会失败
# Error: Donation only 9:00-18:00 UTC+8
```

### 捐赠排行榜测试
```bash
# 查看排行榜
cast call $CONTRACT_ADDRESS "getTopDonors()" --rpc-url sepolia

# 查询个人捐赠
YOUR_ADDRESS=$(cast wallet address --private-key $PRIVATE_KEY)
cast call $CONTRACT_ADDRESS "getDonation(address)" $YOUR_ADDRESS --rpc-url sepolia
```

### 权限控制测试
```bash
# 只有owner可以提取资金
cast call $CONTRACT_ADDRESS "owner()" --rpc-url sepolia

# Owner提取资金
cast send $CONTRACT_ADDRESS "withdraw()" --rpc-url sepolia --private-key $PRIVATE_KEY

# 非owner会失败: Error: Not owner
```

## 监控和调试

### 查看事件日志
```bash
# 监控Donation事件
cast logs --rpc-url sepolia --address $CONTRACT_ADDRESS --decode-logs "Donation(address indexed,uint256)"
```

### 查看交易详情
```bash
# 获取交易receipt
cast receipt <tx_hash> --rpc-url sepolia

# 查看gas使用情况
cast tx <tx_hash> --rpc-url sepolia
```

## 常见错误处理

### 1. "Donation only 9:00-18:00 UTC+8"
**原因**: 当前时间不在工作时间内
**解决**: 等待工作时间或修改合约用于测试

### 2. "Zero donation"
**原因**: 尝试捐赠0 ETH
**解决**: 捐赠大于0的金额

### 3. "Not owner"
**原因**: 非owner尝试调用withdraw()
**解决**: 使用owner私钥调用

### 4. "insufficient funds"
**原因**: 账户ETH余额不足
**解决**: 从水龙头获取测试ETH

## 部署验证清单

- [ ] 合约成功部署并获得地址
- [ ] 在Etherscan上成功验证
- [ ] Owner地址正确设置
- [ ] 工作时间内可以正常捐赠
- [ ] 工作时间外捐赠被拒绝
- [ ] 捐赠排行榜正确更新
- [ ] Donation事件正确触发
- [ ] Owner可以成功提取资金
- [ ] 非owner无法提取资金

## 实用命令汇总

```bash
# 部署
make deploy-beggar

# 测试
export CONTRACT_ADDRESS=0x1234...
make test-beggar-sepolia

# 捐赠 (工作时间内)
cast send $CONTRACT_ADDRESS "donate()" --value 0.01ether --rpc-url sepolia --private-key $PRIVATE_KEY

# 查询排行榜
cast call $CONTRACT_ADDRESS "getTopDonors()" --rpc-url sepolia

# 提取资金 (仅owner)
cast send $CONTRACT_ADDRESS "withdraw()" --rpc-url sepolia --private-key $PRIVATE_KEY

# 查看Etherscan
echo "https://sepolia.etherscan.io/address/$CONTRACT_ADDRESS"
```

## 成功部署示例

```
=== Deployment Summary ===
Contract Address: 0xAbC123...
Owner Address: 0xDeF456...
Network: Sepolia Testnet
Working Hours: 9:00-18:00 UTC+8

=== Next Steps ===
1. Verify contract on Etherscan
2. Test donate() function during working hours
3. Check ranking with getTopDonors()
```

完成这些步骤后，你的BeggingContract就成功部署并可以在Sepolia测试网上使用了！