# 智能合约部署原理

## 核心概念

### 1. 合约地址计算
合约地址 = keccak256(部署者地址 + nonce)

- 部署者地址：你的钱包地址
- nonce：该地址的交易计数

每次部署 nonce 增加，所以地址永远不同。

### 2. 多次部署示例

```
第一次部署：
- 部署者：0x1804c8AB1F12E6bbf3894d4083f33e07309d1f38
- nonce：5
- 合约地址：0x1e05FD7d24Cb492a76C79c3E4F565F44D8bf1691

第二次部署：
- 部署者：0x1804c8AB1F12E6bbf3894d4083f33e07309d1f38
- nonce：6
- 合约地址：0x完全不同的新地址
```

### 3. 实际影响

**不会冲突的原因**：
- ✅ 不同地址 = 不同存储空间
- ✅ 不同状态 = 独立的 NFT ID 计数
- ✅ 不同余额 = 独立的资产

**需要注意的**：
- ⚠️ 用户需要知道新地址
- ⚠️ 旧合约的 NFT 还在旧地址
- ⚠️ 前端需要更新合约地址

## 实践建议

### 开发阶段
随便部署，每次都是新的：
```bash
make deploy-testnet  # 部署新实例
```

### 生产阶段
1. **代理模式**（可升级）
   ```solidity
   // 使用 OpenZeppelin 的代理合约
   // 逻辑可更新，地址不变
   ```

2. **工厂模式**（多实例）
   ```solidity
   contract NFTFactory {
       function createNFT() returns (address) {
           return address(new SimpleNFT());
       }
   }
   ```

## 常见误解

❌ **错误**："重新部署会覆盖旧合约"
✅ **正确**：每次部署都创建新的独立合约

❌ **错误**："需要先删除旧合约"
✅ **正确**：区块链上的合约无法删除（除非有自毁功能）

❌ **错误**："两个合约会共享数据"
✅ **正确**：每个合约有独立的存储空间

## 测试多次部署

```bash
# 部署第一个
forge script script/Deploy.s.sol --rpc-url sepolia --broadcast
# 输出：合约 A 地址

# 立即再部署
forge script script/Deploy.s.sol --rpc-url sepolia --broadcast
# 输出：合约 B 地址（完全不同）

# 两个合约独立运行
cast call <合约A> "totalSupply()"  # 返回 1
cast call <合约B> "totalSupply()"  # 返回 1
```

## 总结

**区块链 = 只增不删的账本**

每次部署都是在账本上添加新页，而不是修改旧页。这是区块链不可变性的体现。