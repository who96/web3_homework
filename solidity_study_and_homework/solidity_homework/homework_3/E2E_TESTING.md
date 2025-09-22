# NFT拍卖市场 E2E测试框架

## 概述

这是一个完整的端到端测试框架，用于验证NFT拍卖市场的所有功能。设计为可复用的Makefile系统，支持可升级合约的重复测试。

## 快速开始

```bash
# 查看所有可用命令
make help

# 完整部署和测试流程
make build          # 编译合约
make deploy         # 部署到Sepolia
make test-e2e-manual # 手动分阶段测试
```

## 测试阶段详解

### Phase 1: 创建NFT和拍卖
```bash
make phase1
```
- 铸造NFT (TokenID: 0)
- 授权拍卖合约
- 创建2分钟拍卖 (储备价: 0.0001 ETH)

### Phase 2: 竞拍阶段
```bash
make phase2
```
- 钱包2: 0.0001 ETH
- 钱包3: 0.0002 ETH
- 钱包4: 0.0003 ETH
- 钱包5: 0.0004 ETH (最高出价者)

### Phase 3: 价格测试
```bash
make phase3
```
- 测试 `getEthUsdPrice()` 函数
- 验证价格计算准确性 ($4185+/ETH)
- 测试拍卖-价格源集成

### Phase 4: 平台钱包测试
```bash
make phase4
```
- 测试 `setPlatformWallet()` 函数
- 将钱包6设为新平台钱包

### Phase 5: 结算阶段
```bash
# 等待拍卖结束后执行
make phase5
```
- 结束拍卖
- 获胜者认领NFT
- 卖家认领资金 (97%)
- 平台费转给新平台钱包 (3%)
- 失败竞拍者申请退款

## 实用工具

```bash
make show-wallets     # 显示所有钱包地址
make check-balances   # 检查所有钱包余额
make verify-deployment # 验证合约部署状态
```

## 验证命令

```bash
make verify-nft-owner    # 验证NFT拥有者
make verify-auction-state # 验证拍卖状态
make verify-platform-fee  # 验证平台费率
```

## 可升级合约支持

当实现可升级拍卖合约后：

1. **更新合约地址**:
   ```bash
   # 编辑 Makefile 中的合约地址
   NFT_CONTRACT = [新地址]
   AUCTION_CONTRACT = [新地址]
   PRICE_FEED_CONTRACT = [新地址]
   ```

2. **运行相同测试**:
   ```bash
   make test-e2e-manual
   ```

3. **对比测试结果**:
   所有功能应保持向后兼容

## 钱包配置

测试使用6个钱包:
- 钱包1: 卖家/部署者
- 钱包2-5: 竞拍者
- 钱包6: 新平台钱包

所有私钥在 `.env` 文件中配置。

## 网络配置

- 测试网: Sepolia
- RPC: PublicNode (稳定可靠)
- Gas优化: 使用 `--legacy` 标志

## Linus式设计原则

1. **简洁性**: 每个阶段独立可测试
2. **可重用性**: Makefile参数化，易于修改
3. **可验证性**: 每步都有验证命令
4. **错误处理**: 退款命令使用 `-` 前缀，允许失败
5. **向后兼容**: 可升级合约必须通过相同测试

## 故障排除

- **交易失败**: 检查钱包余额和gas limit
- **拍卖超时**: 使用 `make wait-auction` 检查状态
- **私钥错误**: 确保 `.env` 中包含 `0x` 前缀