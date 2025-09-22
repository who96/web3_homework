# 可升级合约测试指南

## 实现可升级拍卖合约后的测试流程

### 1. 更新Makefile配置

部署新的可升级合约后，更新 `Makefile` 中的合约地址：

```bash
# 在Makefile中更新这些变量
NFT_CONTRACT = [新的NFT合约地址]
AUCTION_CONTRACT = [新的可升级拍卖合约地址]
PRICE_FEED_CONTRACT = [新的价格源合约地址]
```

### 2. 向后兼容性测试

```bash
# 完整的回归测试
make verify-deployment  # 验证新合约部署
make phase1            # 测试基本功能
make phase2            # 测试竞拍逻辑
make phase3            # 验证价格源集成
make phase4            # 测试管理员功能
# 等待120秒...
make phase5            # 测试结算逻辑
```

### 3. 升级特有功能测试

为可升级合约添加专门的测试目标：

```makefile
# 在Makefile中添加升级测试
test-upgrade:
	@echo "🚀 Testing upgrade functionality..."
	# 测试升级逻辑
	cast send $(AUCTION_CONTRACT) "upgrade()" \
		--private-key $(PRIVATE_KEY) --rpc-url $(RPC_URL) --legacy

test-storage:
	@echo "📦 Testing storage compatibility..."
	# 验证升级后存储槽不冲突
	cast call $(AUCTION_CONTRACT) "getStorageLayout()" --rpc-url $(RPC_URL)
```

### 4. 性能对比

```bash
# 记录gas消耗
make phase1 | grep "gasUsed"
make phase2 | grep "gasUsed"
# 对比升级前后的gas成本
```

### 5. Linus式质量检查清单

- [ ] **Never break userspace**: 所有现有功能必须正常工作
- [ ] **Good taste**: 升级逻辑简洁，无特殊情况
- [ ] **实用主义**: 解决真实问题，不是理论完美
- [ ] **简洁性**: 升级不增加复杂度

### 6. 示例测试脚本

创建 `test-upgrade.sh` 来自动化对比测试：

```bash
#!/bin/bash
echo "=== Pre-upgrade Test ==="
# 使用旧合约地址测试
OLD_AUCTION=0xOLD_ADDRESS make test-e2e-manual

echo "=== Post-upgrade Test ==="
# 使用新合约地址测试
NEW_AUCTION=0xNEW_ADDRESS make test-e2e-manual

echo "=== Comparison ==="
# 对比结果
diff pre-upgrade.log post-upgrade.log
```

这套框架确保可升级合约实现后，所有现有功能保持完全兼容。