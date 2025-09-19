# BeggingContract 提现功能测试指南

## 📋 测试场景

### 场景1：Owner提现测试
```bash
# 1. 确保你是合约owner并且合约有余额
export CONTRACT_ADDRESS=0xYourContractAddress
make beggar-info

# 2. 运行完整测试（包含提现）
make test-beggar-sepolia

# 3. 单独执行提现（谨慎！）
make withdraw-beggar
```

### 场景2：非Owner拒绝测试
```bash
# 使用非owner账户的私钥测试
export PRIVATE_KEY=non_owner_private_key
make test-beggar-sepolia
# 应该看到"✅ 正确拒绝非owner的提现请求"
```

## 🧪 测试脚本详解

测试脚本会自动执行以下检查：

### 1. 地址格式转换
```bash
# ABI编码地址转换为标准格式
OWNER_RAW=0x0000000000000000000000005d4a329b262ac7a1d9ae0f4c54171df61e2c0b69
OWNER=0x5d4a329b262ac7a1d9ae0f4c54171df61e2c0b69  # 提取后40个字符
```

### 2. 权限验证
```bash
# 大小写不敏感比较
OWNER_LOWER=0x5d4a329b262ac7a1d9ae0f4c54171df61e2c0b69
YOUR_ADDRESS_LOWER=0x5d4a329b262ac7a1d9ae0f4c54171df61e2c0b69
# 匹配 -> "你是合约owner"
```

### 3. 提现逻辑测试

#### 如果你是Owner:
```bash
=== 💸 提现功能测试 ===
合约余额: 0.05 ether
💰 合约有余额，测试提现功能...
Owner提现前余额: 1.234 ether
执行withdraw()...
✅ 提现成功! 交易hash: 0xabc123...
Owner提现后余额: 1.284 ether  # 增加了0.05减去gas费
合约提现后余额: 0 ether
✅ 合约余额已清空
```

#### 如果你不是Owner:
```bash
=== 💸 提现功能测试 ===
🔒 非owner账户，测试拒绝访问...
尝试以非owner身份调用withdraw()（应该失败）...
✅ 正确拒绝非owner的提现请求
```

## 📊 测试结果解读

### ✅ 正常情况
- **Owner提现**: 合约余额转移到owner地址
- **非Owner拒绝**: 交易被回滚，显示"Not owner"错误
- **余额验证**: 提现后合约余额为0

### ⚠️ 异常情况
- **合约无余额**: 显示"合约余额为0，无法测试提现功能"
- **Gas不足**: 交易失败，余额未变化
- **网络错误**: 显示RPC超时或连接错误

## 🔧 手动提现命令

### 安全提现（推荐）
```bash
# 1. 检查合约状态
export CONTRACT_ADDRESS=0xYourAddress
make beggar-info

# 2. 执行提现（有确认提示）
make withdraw-beggar

# 3. 验证结果
make beggar-info  # 检查合约余额应为0
```

### 直接提现（高级用户）
```bash
# 一键提现
cast send $CONTRACT_ADDRESS "withdraw()" --rpc-url sepolia --private-key $PRIVATE_KEY

# 检查交易结果
cast receipt <tx_hash> --rpc-url sepolia
```

## 🎯 测试检查清单

部署并测试提现功能时，确保：

- [ ] 合约成功部署
- [ ] Owner地址正确设置
- [ ] 合约有捐赠余额（通过donate()函数）
- [ ] Owner可以成功提现
- [ ] 非Owner提现被拒绝
- [ ] 提现后合约余额为0
- [ ] Owner余额增加（减去gas费用）

## 📝 常见问题

### Q: 提现后为什么Owner余额没增加预期的数量？
A: Gas费用会从提现金额中扣除，实际到账 = 提现金额 - Gas费用

### Q: 为什么非Owner测试时没有报"Not owner"错误？
A: 某些RPC可能返回通用的revert错误，脚本会识别任何revert作为正确的拒绝行为

### Q: 合约余额为0时如何测试？
A: 先在工作时间内调用donate()函数向合约捐赠一些ETH

### Q: 测试脚本显示"意外结果"怎么办？
A: 检查网络连接、私钥格式、合约地址是否正确

## 🚀 完整测试流程

```bash
# 1. 准备环境
export CONTRACT_ADDRESS=0xYourContractAddress
export PRIVATE_KEY=your_private_key

# 2. 在工作时间内捐赠（可选）
cast send $CONTRACT_ADDRESS "donate()" --value 0.1ether --rpc-url sepolia --private-key $PRIVATE_KEY

# 3. 运行完整测试
make test-beggar-sepolia

# 4. 检查结果
make beggar-info
```

提现功能现在完全集成到测试脚本中，会根据你的权限自动进行相应的测试！