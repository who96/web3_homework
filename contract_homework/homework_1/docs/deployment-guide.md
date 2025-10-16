# Fukua (FK) Meme代币部署指南

这是一份完整的SHIB风格Meme代币部署和使用指南。

## 目录
- [环境准备](#环境准备)
- [快速开始](#快速开始)
- [部署到测试网](#部署到测试网)
- [添加流动性](#添加流动性)
- [合约验证](#合约验证)
- [使用指南](#使用指南)
- [FAQ](#faq)

---

## 环境准备

### 1. 安装Foundry

```bash
curl -L https://foundry.paradigm.xyz | bash
foundryup
```

验证安装:
```bash
forge --version
cast --version
```

### 2. 配置环境变量

复制`.env.example`为`.env`并填写:

```bash
# 部署者私钥 (不要泄露!)
PRIVATE_KEY=your_private_key_here

# WALLET1地址 (税收接收地址,99%代币发送地址)
WALLET1=0x...

# RPC URLs
SEPOLIA_RPC_URL=https://sepolia.infura.io/v3/YOUR_KEY
MAINNET_RPC_URL=https://mainnet.infura.io/v3/YOUR_KEY

# Etherscan API Key (用于合约验证)
ETHERSCAN_API_KEY=your_etherscan_key
```

**⚠️ 安全提醒:**
- 永远不要提交`.env`文件到git
- 不要在主网部署前使用测试私钥
- 建议使用硬件钱包或多签钱包

### 3. 安装依赖

```bash
forge install
```

### 4. 编译合约

```bash
forge build
```

### 5. 运行测试

```bash
forge test -vvv
```

所有33个测试应该通过。

---

## 快速开始

### 本地测试部署

```bash
# 1. 启动本地节点
anvil

# 2. 在新终端部署
forge script script/DeployMemeToken.s.sol \
  --rpc-url http://localhost:8545 \
  --broadcast
```

---

## 部署到测试网

### 步骤1: 获取测试ETH

访问 [Sepolia Faucet](https://sepoliafaucet.com/) 获取测试ETH。

### 步骤2: 部署合约

```bash
forge script script/DeployMemeToken.s.sol \
  --rpc-url $SEPOLIA_RPC_URL \
  --broadcast \
  --verify \
  -vvvv
```

**输出示例:**
```
===========================================
Deploying MemeToken...
===========================================
Deployer: 0x123...
Tax Recipient (WALLET1): 0x456...
Initial Tax Rate: 300 basis points (3%)
===========================================
Token deployed at: 0x789...
Token name: Fukua
Token symbol: FK
Total supply: 21000000 FK
Max transaction amount: 420000 FK
===========================================
```

**记录合约地址** - 你会需要它！

### 步骤3: 验证部署

```bash
# 查询代币信息
cast call <TOKEN_ADDRESS> "name()(string)" --rpc-url $SEPOLIA_RPC_URL
cast call <TOKEN_ADDRESS> "symbol()(string)" --rpc-url $SEPOLIA_RPC_URL
cast call <TOKEN_ADDRESS> "totalSupply()(uint256)" --rpc-url $SEPOLIA_RPC_URL

# 查询税率
cast call <TOKEN_ADDRESS> "taxRate()(uint256)" --rpc-url $SEPOLIA_RPC_URL

# 查询税收地址
cast call <TOKEN_ADDRESS> "taxRecipient()(address)" --rpc-url $SEPOLIA_RPC_URL
```

---

## 添加流动性

### ⚠️ 重要说明

**Sepolia测试网**没有官方Uniswap V2部署。因此我们提供两种方案：
- **方案A (推荐测试网使用)**: 部署MockDEX - 简化的DEX用于演示
- **方案B (主网使用)**: 使用真实Uniswap V2

---

### 方案A: 使用MockDEX (Sepolia测试网)

#### 1. 部署MockDEX并自动添加流动性

```bash
forge script script/DeployMockDEX.s.sol \
  --rpc-url $SEPOLIA_RPC_URL \
  --broadcast \
  -vvv
```

**脚本会自动完成**:
1. 部署MockDEX合约
2. 授权DEX使用210,000 FK
3. 添加流动性 (210,000 FK + 0.1 ETH)
4. 输出DEX地址和白名单命令

**输出示例**:
```
MockDEX deployed at: 0x4961dDb55265Bcd4E230B2aceaf257A745e73de0
Liquidity added!
- Token amount: 210000 FK
- ETH amount: 0.1 ETH
- LP shares: 144913767461894385737

==========================================
IMPORTANT: Add DEX to whitelist!
==========================================
Run the following commands as owner:

cast send 0x61a33158B1541AD0fc87DF41075ac6A40CC52498 \
  'setTaxExempt(address,bool)' \
  0x4961dDb55265Bcd4E230B2aceaf257A745e73de0 true \
  --rpc-url $SEPOLIA_RPC_URL --private-key $PRIVATE_KEY

cast send 0x61a33158B1541AD0fc87DF41075ac6A40CC52498 \
  'setLimitExempt(address,bool)' \
  0x4961dDb55265Bcd4E230B2aceaf257A745e73de0 true \
  --rpc-url $SEPOLIA_RPC_URL --private-key $PRIVATE_KEY
==========================================
```

#### 2. 将MockDEX加入白名单

**复制脚本输出的命令并执行：**

```bash
# 税收豁免
cast send <TOKEN_ADDRESS> \
  "setTaxExempt(address,bool)" \
  <MOCKDEX_ADDRESS> \
  true \
  --rpc-url $SEPOLIA_RPC_URL \
  --private-key $PRIVATE_KEY

# 交易限制豁免
cast send <TOKEN_ADDRESS> \
  "setLimitExempt(address,bool)" \
  <MOCKDEX_ADDRESS> \
  true \
  --rpc-url $SEPOLIA_RPC_URL \
  --private-key $PRIVATE_KEY
```

#### 3. 测试Swap交易

```bash
# ETH 换 Token
cast send <MOCKDEX_ADDRESS> \
  "swapEthForToken(uint256)" \
  0 \
  --value 0.01ether \
  --rpc-url $SEPOLIA_RPC_URL \
  --private-key $PRIVATE_KEY

# Token 换 ETH (需要先授权)
# 先授权
cast send <TOKEN_ADDRESS> \
  "approve(address,uint256)" \
  <MOCKDEX_ADDRESS> \
  $(cast to-wei 1000) \
  --rpc-url $SEPOLIA_RPC_URL \
  --private-key $PRIVATE_KEY

# 再swap
cast send <MOCKDEX_ADDRESS> \
  "swapTokenForEth(uint256,uint256)" \
  $(cast to-wei 1000) \
  0 \
  --rpc-url $SEPOLIA_RPC_URL \
  --private-key $PRIVATE_KEY
```

#### 4. 查询MockDEX信息

```bash
# 查询储备
cast call <MOCKDEX_ADDRESS> "tokenReserve()(uint256)" --rpc-url $SEPOLIA_RPC_URL
cast call <MOCKDEX_ADDRESS> "ethReserve()(uint256)" --rpc-url $SEPOLIA_RPC_URL

# 查询价格 (0.01 ETH能换多少FK)
cast call <MOCKDEX_ADDRESS> \
  "getTokenOut(uint256)(uint256)" \
  10000000000000000 \
  --rpc-url $SEPOLIA_RPC_URL

# 查询LP份额
cast call <MOCKDEX_ADDRESS> \
  "liquidityShares(address)(uint256)" \
  <YOUR_ADDRESS> \
  --rpc-url $SEPOLIA_RPC_URL
```

**MockDEX特性**:
- 恒定乘积AMM (x * y = k)
- 0.3%交易手续费
- 支持双向swap
- 支持添加/移除流动性

---

### 方案B: 使用Uniswap V2 (主网或支持Uniswap的测试网)

#### 1. 使用自动化脚本

```bash
# 1. 设置合约地址
export MEME_TOKEN_ADDRESS=<YOUR_TOKEN_ADDRESS>

# 2. 运行流动性脚本
forge script script/AddLiquidity.s.sol \
  --rpc-url $MAINNET_RPC_URL \
  --broadcast \
  -vvvv
```

脚本会:
1. 授权Uniswap Router使用210,000 FK
2. 添加 210,000 FK + 0.1 ETH 到Uniswap V2
3. 输出Pair地址

#### 2. 手动添加流动性

##### 步骤1: 授权Uniswap Router

```bash
# Uniswap V2 Router地址
# Mainnet: 0x7a250d5630B4cF539739dF2C5dAcb4c659F2488D

cast send <TOKEN_ADDRESS> \
  "approve(address,uint256)" \
  0x7a250d5630B4cF539739dF2C5dAcb4c659F2488D \
  $(cast to-wei 210000) \
  --rpc-url $MAINNET_RPC_URL \
  --private-key $PRIVATE_KEY
```

##### 步骤2: 添加流动性

通过Uniswap界面或使用`cast`调用`addLiquidityETH`:

```bash
cast send 0x7a250d5630B4cF539739dF2C5dAcb4c659F2488D \
  "addLiquidityETH(address,uint256,uint256,uint256,address,uint256)" \
  <TOKEN_ADDRESS> \
  $(cast to-wei 210000) \
  $(cast to-wei 199500) \
  $(cast to-wei 0.095) \
  <YOUR_ADDRESS> \
  $(($(date +%s) + 300)) \
  --value 0.1ether \
  --rpc-url $MAINNET_RPC_URL \
  --private-key $PRIVATE_KEY
```

##### 步骤3: 获取Pair地址

```bash
# Uniswap Factory地址 (Mainnet)
# 0x5C69bEe701ef814a2B6a3EDD4B1652CB9cc5aA6f

cast call 0x5C69bEe701ef814a2B6a3EDD4B1652CB9cc5aA6f \
  "getPair(address,address)(address)" \
  <TOKEN_ADDRESS> \
  0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2 \
  --rpc-url $MAINNET_RPC_URL
```

##### 步骤4: 将Pair加入白名单

**这一步至关重要！** 否则LP操作会被收税。

```bash
# 税收豁免
cast send <TOKEN_ADDRESS> \
  "setTaxExempt(address,bool)" \
  <PAIR_ADDRESS> \
  true \
  --rpc-url $MAINNET_RPC_URL \
  --private-key $PRIVATE_KEY

# 交易限制豁免
cast send <TOKEN_ADDRESS> \
  "setLimitExempt(address,bool)" \
  <PAIR_ADDRESS> \
  true \
  --rpc-url $MAINNET_RPC_URL \
  --private-key $PRIVATE_KEY
```

验证白名单:
```bash
cast call <TOKEN_ADDRESS> \
  "isExemptFromTax(address)(bool)" \
  <PAIR_ADDRESS> \
  --rpc-url $MAINNET_RPC_URL
```

---

### 流动性方案对比

| 特性 | MockDEX (Sepolia) | Uniswap V2 (主网) |
|------|-------------------|-------------------|
| 部署网络 | Sepolia测试网 | Ethereum主网 |
| 合约地址 | 0x4961...73de0 | 0x7a25...2488D |
| AMM算法 | 恒定乘积 (x*y=k) | 恒定乘积 (x*y=k) |
| 手续费 | 0.3% | 0.3% |
| 流动性 | 210k FK + 0.1 ETH | 自定义 |
| 使用场景 | 测试演示 | 生产环境 |
| 安全性 | 未审计 | 久经考验 |
| 部署脚本 | DeployMockDEX.s.sol | AddLiquidity.s.sol |

---

## 合约验证

### 自动验证 (推荐)

部署时添加`--verify`标志:

```bash
forge script script/DeployMemeToken.s.sol \
  --rpc-url $SEPOLIA_RPC_URL \
  --broadcast \
  --verify
```

### 手动验证

```bash
forge verify-contract \
  <TOKEN_ADDRESS> \
  src/MemeToken.sol:MemeToken \
  --chain sepolia \
  --constructor-args $(cast abi-encode "constructor(address,address,uint256)" <OWNER> <TAX_RECIPIENT> 300)
```

验证后可以在Etherscan上看到源代码和读写合约功能。

---

## 使用指南

### 1. 代币转账

#### 使用MetaMask或其他钱包

1. 添加代币到钱包: 使用合约地址导入
2. 发送转账
3. **注意:** 普通用户之间转账会收取3%税费

#### 使用cast命令

```bash
# 转账1000 FK
cast send <TOKEN_ADDRESS> \
  "transfer(address,uint256)" \
  <RECIPIENT> \
  $(cast to-wei 1000) \
  --rpc-url $SEPOLIA_RPC_URL \
  --private-key $PRIVATE_KEY
```

### 2. 查看税费信息

```bash
# 当前税率 (basis points, 300 = 3%)
cast call <TOKEN_ADDRESS> "taxRate()(uint256)" --rpc-url $SEPOLIA_RPC_URL

# 计算税费
cast call <TOKEN_ADDRESS> \
  "calculateTax(uint256)(uint256)" \
  $(cast to-wei 1000) \
  --rpc-url $SEPOLIA_RPC_URL

# 税收接收地址
cast call <TOKEN_ADDRESS> "taxRecipient()(address)" --rpc-url $SEPOLIA_RPC_URL
```

### 3. 管理白名单 (仅Owner)

#### 查询白名单状态

```bash
# 税收豁免
cast call <TOKEN_ADDRESS> \
  "isExemptFromTax(address)(bool)" \
  <ADDRESS> \
  --rpc-url $SEPOLIA_RPC_URL

# 交易限制豁免
cast call <TOKEN_ADDRESS> \
  "isExemptFromLimit(address)(bool)" \
  <ADDRESS> \
  --rpc-url $SEPOLIA_RPC_URL
```

#### 添加到白名单

```bash
# 税收豁免
cast send <TOKEN_ADDRESS> \
  "setTaxExempt(address,bool)" \
  <ADDRESS> \
  true \
  --rpc-url $SEPOLIA_RPC_URL \
  --private-key $PRIVATE_KEY

# 交易限制豁免
cast send <TOKEN_ADDRESS> \
  "setLimitExempt(address,bool)" \
  <ADDRESS> \
  true \
  --rpc-url $SEPOLIA_RPC_URL \
  --private-key $PRIVATE_KEY
```

#### 从白名单移除

将上述命令中的`true`改为`false`。

### 4. 修改税率 (仅Owner)

```bash
# 修改税率为2% (200 basis points)
cast send <TOKEN_ADDRESS> \
  "setTaxRate(uint256)" \
  200 \
  --rpc-url $SEPOLIA_RPC_URL \
  --private-key $PRIVATE_KEY
```

**限制:** 税率不能超过5% (500 basis points)

### 5. 修改税收地址 (仅Owner)

```bash
cast send <TOKEN_ADDRESS> \
  "setTaxRecipient(address)" \
  <NEW_ADDRESS> \
  --rpc-url $SEPOLIA_RPC_URL \
  --private-key $PRIVATE_KEY
```

### 6. 修改交易限制 (仅Owner)

```bash
# 修改最大单笔交易为100,000 FK
cast send <TOKEN_ADDRESS> \
  "setMaxTransactionAmount(uint256)" \
  $(cast to-wei 100000) \
  --rpc-url $SEPOLIA_RPC_URL \
  --private-key $PRIVATE_KEY
```

### 7. Uniswap交易

#### 在Uniswap界面交易

1. 访问 [Uniswap](https://app.uniswap.org)
2. 连接钱包
3. 选择Sepolia网络
4. 导入FK代币 (使用合约地址)
5. Swap ETH <-> FK

**注意:**
- 买入FK: 从Pair转出到买家时收取3%税费
- 卖出FK: 从卖家转入Pair时收取3%税费
- LP操作: 因为Pair在白名单，不收税

#### 使用cast交易

```bash
# 通过Router swap
cast send <UNISWAP_ROUTER> \
  "swapExactETHForTokens(uint256,address[],address,uint256)" \
  $(cast to-wei 90) \
  "[<WETH>,<TOKEN_ADDRESS>]" \
  <YOUR_ADDRESS> \
  $(cast --to-hex $(date +%s + 300)) \
  --value 0.1ether \
  --rpc-url $SEPOLIA_RPC_URL \
  --private-key $PRIVATE_KEY
```

---

## FAQ

### Q1: 为什么我的转账失败了?

**A:** 检查以下几点:
1. 是否超过单笔最大交易限制 (默认420,000 FK)?
2. 余额是否充足 (包括税费)?
3. 是否有足够的gas?

### Q2: 如何查看当前交易限制?

```bash
cast call <TOKEN_ADDRESS> "maxTransactionAmount()(uint256)" --rpc-url $SEPOLIA_RPC_URL
```

### Q3: Owner如何放弃所有权?

```bash
cast send <TOKEN_ADDRESS> \
  "renounceOwnership()" \
  --rpc-url $SEPOLIA_RPC_URL \
  --private-key $PRIVATE_KEY
```

**警告:** 放弃所有权后无法修改税率、白名单等参数！

### Q4: 如何在Uniswap移除流动性?

1. 访问 [Uniswap Pool](https://app.uniswap.org/#/pool)
2. 找到FK-ETH池子
3. 点击"Remove"
4. 选择移除的比例
5. 确认交易

**因为Pair在白名单，移除流动性不会被收税。**

### Q5: Gas费大约多少?

根据测试:
- 白名单转账: ~30k gas
- 含税转账: ~65k gas
- 部署: ~2.5M gas
- 添加流动性: ~200k gas

### Q6: 如何查看税费收取历史?

在Etherscan上查看`TaxCollected`事件:
1. 访问合约页面
2. 点击"Events"标签
3. 过滤`TaxCollected`事件

### Q7: 可以修改代币名称或符号吗?

**不可以。** 名称(Fukua)和符号(FK)在部署时固定,无法修改。

### Q8: 为什么部署后要把20,790,000 FK转给WALLET1?

这是设计的分配方案:
- 210,000 FK (1%) → Uniswap流动性池
- 20,790,000 FK (99%) → WALLET1 (项目方/营销钱包)

### Q9: 如何在主网部署?

1. **再三确认所有参数**
2. 使用硬件钱包或多签
3. 先在测试网完整测试
4. 准备足够的ETH (部署 + 流动性 + gas)
5. 替换RPC URL为mainnet
6. 执行部署脚本

```bash
forge script script/DeployMemeToken.s.sol \
  --rpc-url $MAINNET_RPC_URL \
  --broadcast \
  --verify \
  -vvvv
```

**主网部署是不可逆的,请务必小心！**

---

## 相关资源

- [Foundry Book](https://book.getfoundry.sh/)
- [OpenZeppelin Contracts](https://docs.openzeppelin.com/contracts/)
- [Uniswap V2 Docs](https://docs.uniswap.org/contracts/v2/overview)
- [Etherscan](https://etherscan.io/)
- [Sepolia Faucet](https://sepoliafaucet.com/)

---

## 支持

如有问题,请查看:
1. 合约源代码: `src/MemeToken.sol`
2. 测试用例: `test/MemeToken.t.sol`
3. OpenSpec文档: `openspec/changes/add-meme-token-contract/`

---

**免责声明:** 本代币仅用于教育目的。部署和使用需自行承担风险。请遵守当地法规。
