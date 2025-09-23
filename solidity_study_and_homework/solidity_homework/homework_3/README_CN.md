# NFT拍卖市场 - 可升级智能合约系统

<div align="center">

![Solidity](https://img.shields.io/badge/Solidity-0.8.7+-blue)
![Foundry](https://img.shields.io/badge/Foundry-Ready-green)
![OpenZeppelin](https://img.shields.io/badge/OpenZeppelin-Upgradeable-orange)
![Chainlink](https://img.shields.io/badge/Chainlink-Price_Feed-yellow)
![License](https://img.shields.io/badge/License-MIT-red)

**基于透明代理模式的可升级NFT拍卖平台**

[English](./README.md) | 中文

</div>

## 📋 目录

- [项目概述](#项目概述)
- [技术架构](#技术架构)
- [项目结构](#项目结构)
- [核心功能](#核心功能)
- [安装与配置](#安装与配置)
- [部署指南](#部署指南)
- [使用说明](#使用说明)
- [测试框架](#测试框架)
- [合约地址](#合约地址)
- [E2E测试](#e2e测试)
- [升级指南](#升级指南)
- [开发指南](#开发指南)
- [故障排除](#故障排除)
- [许可证](#许可证)

## 🎯 项目概述

这是一个功能完整的NFT拍卖市场智能合约系统，基于**透明代理模式**实现合约升级能力。项目集成了Chainlink价格预言机，提供安全可靠的链上拍卖体验。

### 核心特性

- 🏛️ **透明代理升级架构** - 支持合约逻辑升级，数据永不丢失
- ⚡ **完整拍卖机制** - 创建、竞价、结算全流程
- 💰 **Chainlink价格集成** - 实时ETH/USD价格显示
- 🛡️ **企业级安全** - 重入攻击防护、权限控制、紧急暂停
- 🧪 **全面测试覆盖** - 单元测试、集成测试、E2E测试
- 🚀 **自动化部署** - Foundry脚本 + Makefile工作流
- 📊 **生产环境验证** - Sepolia测试网完整部署

### 业务场景

- NFT数字藏品拍卖
- 游戏道具交易市场
- 艺术品数字化拍卖
- 版权IP交易平台

## 🏗️ 技术架构

### 系统架构图

```
┌─────────────────┐    ┌─────────────────┐    ┌─────────────────┐
│   用户界面      │    │   价格预言机    │    │   管理控制台    │
│   (Web3 DApp)  │    │  (Chainlink)    │    │  (Owner Only)   │
└─────────┬───────┘    └─────────┬───────┘    └─────────┬───────┘
          │                      │                      │
          ▼                      ▼                      ▼
┌─────────────────────────────────────────────────────────────────┐
│                        智能合约层                               │
├─────────────────┬─────────────────┬─────────────────────────────┤
│   透明代理      │   NFT合约       │      价格Feed合约           │
│ (Proxy Layer)   │ (SimpleNFT)     │  (AuctionPriceFeed)         │
├─────────────────┼─────────────────┼─────────────────────────────┤
│   实现合约      │   拍卖逻辑      │      升级管理               │
│(Implementation) │(Auction Logic)  │   (ProxyAdmin)              │
└─────────────────┴─────────────────┴─────────────────────────────┘
          │                      │                      │
          ▼                      ▼                      ▼
┌─────────────────────────────────────────────────────────────────┐
│                      区块链网络                                 │
│              (Ethereum Sepolia Testnet)                        │
└─────────────────────────────────────────────────────────────────┘
```

### 核心技术栈

- **智能合约框架**: Solidity 0.8.7+
- **开发工具链**: Foundry (Forge + Cast + Anvil)
- **升级模式**: OpenZeppelin透明代理 (TransparentUpgradeableProxy)
- **预言机**: Chainlink Price Feeds
- **测试网络**: Ethereum Sepolia
- **安全库**: OpenZeppelin Contracts (ReentrancyGuard, Ownable, Initializable)

### 升级架构详解

```solidity
用户调用 → 代理合约 → delegatecall → 实现合约
  (固定地址)  (数据存储)      (逻辑处理)   (可升级)

存储层 (代理合约):
- 用户永远交互的地址
- 所有状态变量存储
- 升级权限控制

逻辑层 (实现合约):
- 业务逻辑处理
- 可独立升级
- 兼容存储布局
```

## 📁 项目结构

```
homework_3/
├── 📁 src/                          # 智能合约源码
│   ├── SimpleNFT.sol               # ERC721 NFT合约
│   ├── SimpleAuction.sol           # 标准拍卖合约
│   ├── AuctionPriceFeed.sol        # Chainlink价格Feed
│   ├── PriceFeedConsumer.sol       # 价格消费者基类
│   ├── 📁 upgradeable/
│   │   ├── SimpleAuctionUpgradeable.sol    # 可升级拍卖合约 V1
│   │   └── SimpleAuctionUpgradeableV2.sol  # 可升级拍卖合约 V2 (演示升级)
│   └── 📁 proxy/
│       ├── AuctionProxy.sol        # 透明代理合约
│       └── AuctionProxyAdmin.sol   # 代理管理合约
├── 📁 test/                         # 测试文件
│   ├── SimpleNFT.t.sol             # NFT合约测试
│   ├── SimpleAuction.t.sol         # 拍卖合约测试
│   ├── 📁 integration/
│   │   └── AuctionIntegration.t.sol # 集成测试
│   ├── 📁 upgradeable/
│   │   ├── BasicProxyTest.t.sol    # 代理基础测试
│   │   └── ProxyUpgrade.t.sol      # 升级功能测试
│   └── 📁 utils/
│       └── Cheats.sol              # 测试工具合约
├── 📁 script/                       # 部署脚本
│   ├── DeployContracts.s.sol       # 标准部署脚本
│   ├── HelperConfig.sol            # 网络配置助手
│   ├── E2EAuctionTestEN.s.sol      # E2E测试脚本
│   └── 📁 upgrade/
│       ├── DeployProxy.s.sol       # 代理部署脚本
│       ├── DeployProxySimple.s.sol # 简化代理部署
│       ├── TestUpgrade.s.sol       # 升级测试脚本
│       └── CheckOwner.s.sol        # 权限检查脚本
├── 📄 Makefile                      # 自动化命令集
├── 📄 foundry.toml                  # Foundry配置
├── 📄 .env.example                  # 环境变量模板
├── 📄 E2E_TESTING.md               # E2E测试文档
├── 📄 UPGRADE_TESTING.md           # 升级测试文档
└── 📄 README_CN.md                 # 中文文档 (本文档)
```

## ⚡ 核心功能

### 🎯 拍卖系统核心功能

#### 1. 拍卖创建
- **NFT授权检查** - 自动验证NFT所有权和授权状态
- **参数验证** - 持续时间、保留价格合法性检查
- **NFT托管** - 拍卖期间NFT安全托管到合约
- **事件发布** - 链上事件通知拍卖创建

#### 2. 竞价系统
- **增量竞价** - 最低5%增量保证公平竞争
- **自动退款** - 被超越竞价者资金自动入退款池
- **时间验证** - 拍卖时间窗口严格控制
- **金额验证** - 竞价金额合规性检查

#### 3. 结算机制
- **拍卖结束** - 时间到期自动可结算
- **NFT转移** - 获胜者自动获得NFT所有权
- **资金分配** - 卖家获得97%，平台获得3%手续费
- **退款处理** - 失败竞价者可主动提取退款

#### 4. 管理功能
- **平台钱包设置** - 动态修改手续费收取地址
- **紧急暂停** - 紧急情况下暂停合约操作 (V2功能)
- **权限控制** - 基于OpenZeppelin的访问控制

### 🔗 Chainlink价格集成

#### 实时价格获取
```solidity
// 获取当前ETH/USD价格
function getEthUsdPrice() external view returns (uint256)

// 获取拍卖价格的USD等值
function getAuctionPriceInUSD(uint256 auctionId) external view returns (uint256, uint256)
```

#### 价格显示特性
- **实时汇率** - Chainlink预言机提供可信价格数据
- **多币种显示** - ETH和USD双重价格展示
- **精度处理** - 正确处理价格精度和小数位
- **异常处理** - 价格源异常时的降级策略

### 🛡️ 安全机制

#### 重入攻击防护
```solidity
// 使用OpenZeppelin ReentrancyGuard
modifier nonReentrant() {
    require(_status != _ENTERED, "ReentrancyGuard: reentrant call");
    _status = _ENTERED;
    _;
    _status = _NOT_ENTERED;
}
```

#### 权限控制系统
- **合约所有者** - 拥有管理权限的账户
- **平台钱包** - 接收手续费的指定地址
- **代理管理员** - 控制合约升级的特殊账户
- **多重验证** - 关键操作需要权限验证

## 🔧 合约API详解

### SimpleNFT.sol - NFT合约

```solidity
contract SimpleNFT is ERC721, Ownable {
    // 铸造NFT (仅所有者)
    function mint(address to) external onlyOwner returns (uint256)

    // 批量铸造NFT
    function batchMint(address to, uint256 amount) external onlyOwner

    // 设置基础URI
    function setBaseURI(string memory newBaseURI) external onlyOwner

    // 获取代币URI
    function tokenURI(uint256 tokenId) public view override returns (string memory)
}
```

### SimpleAuctionUpgradeable.sol - 核心拍卖合约

#### 关键状态变量
```solidity
struct Auction {
    address nftContract;    // NFT合约地址
    uint256 tokenId;       // NFT token ID
    address seller;        // 卖家地址
    uint256 startTime;     // 拍卖开始时间
    uint256 duration;      // 拍卖持续时间
    uint256 reservePrice;  // 保留价格
    address highestBidder; // 最高出价者
    uint256 highestBid;    // 最高出价
    bool ended;           // 是否已结束
    bool claimed;         // 是否已领取
}
```

#### 核心函数接口
```solidity
// 创建拍卖
function createAuction(
    address nftContract,
    uint256 tokenId,
    uint256 duration,      // 秒为单位
    uint256 reservePrice   // wei为单位
) external nonReentrant returns (uint256 auctionId)

// 参与竞价
function placeBid(uint256 auctionId)
    external payable nonReentrant

// 结束拍卖
function endAuction(uint256 auctionId)
    external nonReentrant

// 获胜者领取NFT
function claimNFT(uint256 auctionId)
    external nonReentrant

// 卖家领取资金
function claimFunds(uint256 auctionId)
    external nonReentrant

// 失败竞价者提取退款
function withdrawRefund()
    external nonReentrant

// 查询函数
function getAuction(uint256 auctionId)
    external view returns (Auction memory)

function isAuctionActive(uint256 auctionId)
    external view returns (bool)

function getPendingRefund(address bidder)
    external view returns (uint256)
```

#### V2版本新功能 (SimpleAuctionUpgradeableV2.sol)
```solidity
// 紧急暂停功能
function toggleEmergencyPause() external onlyOwner

function isEmergencyPaused() external view returns (bool)

// 带暂停检查的重写函数
function createAuction(...) external override nonReentrant notEmergencyPaused

function placeBid(uint256 auctionId) external payable override nonReentrant notEmergencyPaused
```

### AuctionPriceFeed.sol - 价格预言机集成

```solidity
contract AuctionPriceFeed {
    // 获取ETH/USD价格
    function getEthUsdPrice() external view returns (uint256)

    // 将ETH金额转换为USD
    function ethToUsd(uint256 ethAmount) external view returns (uint256)

    // 获取拍卖的USD价格信息
    function getAuctionPriceInUSD(uint256 auctionId)
        external view returns (uint256 ethPrice, uint256 usdPrice)

    // 更新价格Feed地址 (仅所有者)
    function updatePriceFeed(address newPriceFeed) external onlyOwner
}
```

### 代理升级系统

#### TransparentUpgradeableProxy
- **用户入口** - 所有用户调用的固定地址
- **存储保持** - 升级过程中数据完全保留
- **权限隔离** - 管理员和普通用户调用隔离

#### ProxyAdmin
```solidity
// 升级实现合约
function upgradeAndCall(
    ITransparentUpgradeableProxy proxy,
    address implementation,
    bytes memory data
) external onlyOwner

// 获取当前实现地址
function getProxyImplementation(address proxy)
    external view returns (address)

// 获取代理管理员
function getProxyAdmin(address proxy)
    external view returns (address)
```

## 🔍 事件系统

### 拍卖事件
```solidity
event AuctionCreated(
    uint256 indexed auctionId,
    address indexed nftContract,
    uint256 indexed tokenId,
    address seller,
    uint256 startTime,
    uint256 duration,
    uint256 reservePrice
);

event BidPlaced(
    uint256 indexed auctionId,
    address indexed bidder,
    uint256 amount
);

event AuctionEnded(
    uint256 indexed auctionId,
    address indexed winner,
    uint256 winningBid
);

event NFTClaimed(
    uint256 indexed auctionId,
    address indexed winner
);

event FundsClaimed(
    uint256 indexed auctionId,
    address indexed seller,
    uint256 amount
);

event RefundWithdrawn(
    address indexed bidder,
    uint256 amount
);
```

### 管理事件
```solidity
event PlatformWalletChanged(
    address indexed oldWallet,
    address indexed newWallet
);

event EmergencyPausedChanged(bool paused); // V2事件
```

## 🚀 安装与配置

### 环境要求

- **Node.js**: >= 16.0.0
- **Git**: 最新版本
- **Foundry**: 最新版本
- **钱包**: MetaMask或其他Web3钱包
- **测试ETH**: Sepolia测试网ETH (通过水龙头获取)

### 安装Foundry

```bash
# 安装Foundry
curl -L https://foundry.paradigm.xyz | bash
foundryup

# 验证安装
forge --version
cast --version
anvil --version
```

### 克隆项目

```bash
# 克隆仓库
git clone https://github.com/who96/web3_homework.git
cd solidity_study_and_homework/solidity_homework/homework_3

# 安装依赖 (Foundry会自动处理)
forge install
```

### 环境配置

1. **复制环境变量模板**
```bash
cp .env.example .env
```

2. **配置.env文件**
```bash
# 网络配置
SEPOLIA_RPC_URL=https://ethereum-sepolia-rpc.publicnode.com

# 私钥配置 (测试环境专用，勿在生产环境使用)
PRIVATE_KEY=your_private_key_here
PRIVATE_KEY_2=wallet2_private_key
PRIVATE_KEY_3=wallet3_private_key
PRIVATE_KEY_4=wallet4_private_key
PRIVATE_KEY_5=wallet5_private_key
PRIVATE_KEY_6=wallet6_private_key

# Etherscan API (用于合约验证)
ETHERSCAN_API_KEY=your_etherscan_api_key
```

3. **Sepolia测试网配置**
- 网络名称: Sepolia
- RPC URL: https://ethereum-sepolia-rpc.publicnode.com
- 链ID: 11155111
- 货币符号: ETH
- 区块浏览器: https://sepolia.etherscan.io

### 获取测试ETH

访问以下水龙头获取Sepolia测试ETH:
- https://sepoliafaucet.com/
- https://www.alchemy.com/faucets/ethereum-sepolia
- https://faucets.chain.link/sepolia

## 📦 部署指南

### 快速部署 (推荐)

```bash
# 1. 编译合约
make build

# 2. 部署到Sepolia (包含代理架构)
make deploy

# 3. 验证部署
make verify-deployment
```

### 分步部署

#### 1. 编译检查
```bash
# 编译所有合约
forge build

# 检查合约大小
forge build --sizes

# 运行静态分析
forge test --gas-report
```

#### 2. 本地测试部署
```bash
# 启动本地节点
anvil

# 在新终端中部署到本地
forge script script/DeployContracts.s.sol:DeployContracts --rpc-url http://localhost:8545 --private-key 0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80 --broadcast
```

#### 3. Sepolia测试网部署
```bash
# 标准部署 (非代理模式)
forge script script/DeployContracts.s.sol:DeployContracts \
  --rpc-url $SEPOLIA_RPC_URL \
  --private-key $PRIVATE_KEY \
  --broadcast \
  --verify \
  --etherscan-api-key $ETHERSCAN_API_KEY

# 代理模式部署 (推荐)
forge script script/upgrade/DeployProxySimple.s.sol:DeployProxySimple \
  --rpc-url $SEPOLIA_RPC_URL \
  --private-key $PRIVATE_KEY \
  --broadcast \
  --verify \
  --etherscan-api-key $ETHERSCAN_API_KEY
```

### 部署后验证

```bash
# 验证合约功能
make verify-contract-version
make verify-implementation
make verify-proxy-admin

# 检查合约余额
make check-balances

# 显示部署地址
make show-wallets
```

## 💼 使用说明

### Makefile命令速查

```bash
# 📊 帮助和信息
make help                    # 显示所有可用命令
make show-wallets           # 显示钱包地址
make check-balances         # 检查余额

# 🔨 构建和部署
make build                  # 编译合约
make deploy                 # 部署到Sepolia
make clean                  # 清理构建产物

# 🧪 测试命令
make test                   # 运行所有测试
make test-e2e              # 运行E2E测试
make test-integration      # 运行集成测试

# 🎯 E2E测试阶段
make phase1                # 创建NFT和拍卖
make phase2                # 执行竞价阶段
make phase3                # 测试价格feed
make phase4                # 测试平台钱包
make phase5                # 结束拍卖和结算
make phase6                # 测试合约升级

# 🔍 验证命令
make verify-deployment     # 验证部署状态
make verify-contract-version  # 检查合约版本
make verify-emergency-pause   # 检查紧急暂停状态
```

### 基本使用流程

#### 1. 创建NFT拍卖
```bash
# 方法1: 使用Makefile (推荐)
make phase1

# 方法2: 使用cast命令
cast send $NFT_CONTRACT "mint(address)" $YOUR_ADDRESS --private-key $PRIVATE_KEY --rpc-url $SEPOLIA_RPC_URL
cast send $NFT_CONTRACT "setApprovalForAll(address,bool)" $AUCTION_CONTRACT true --private-key $PRIVATE_KEY --rpc-url $SEPOLIA_RPC_URL
cast send $AUCTION_CONTRACT "createAuction(address,uint256,uint256,uint256)" $NFT_CONTRACT 0 3600 0.001ether --private-key $PRIVATE_KEY --rpc-url $SEPOLIA_RPC_URL
```

#### 2. 参与竞价
```bash
# 使用Makefile
make phase2

# 使用cast命令
cast send $AUCTION_CONTRACT "placeBid(uint256)" 0 --value 0.001ether --private-key $BIDDER_PRIVATE_KEY --rpc-url $SEPOLIA_RPC_URL
```

#### 3. 拍卖结算
```bash
# 等待拍卖结束后
make phase5

# 或手动结算
cast send $AUCTION_CONTRACT "endAuction(uint256)" 0 --private-key $PRIVATE_KEY --rpc-url $SEPOLIA_RPC_URL
cast send $AUCTION_CONTRACT "claimNFT(uint256)" 0 --private-key $WINNER_PRIVATE_KEY --rpc-url $SEPOLIA_RPC_URL
cast send $AUCTION_CONTRACT "claimFunds(uint256)" 0 --private-key $SELLER_PRIVATE_KEY --rpc-url $SEPOLIA_RPC_URL
```

#### 4. 查询拍卖状态
```bash
# 查询拍卖信息
cast call $AUCTION_CONTRACT "getAuction(uint256)" 0 --rpc-url $SEPOLIA_RPC_URL

# 查询是否活跃
cast call $AUCTION_CONTRACT "isAuctionActive(uint256)" 0 --rpc-url $SEPOLIA_RPC_URL

# 查询待退款金额
cast call $AUCTION_CONTRACT "getPendingRefund(address)" $YOUR_ADDRESS --rpc-url $SEPOLIA_RPC_URL

# 查询当前ETH价格
cast call $PRICE_FEED_CONTRACT "getEthUsdPrice()" --rpc-url $SEPOLIA_RPC_URL
```

### Web3集成示例

#### JavaScript/TypeScript集成
```javascript
// 使用ethers.js
import { ethers } from 'ethers';

const provider = new ethers.providers.JsonRpcProvider(process.env.SEPOLIA_RPC_URL);
const signer = new ethers.Wallet(process.env.PRIVATE_KEY, provider);

// 合约实例
const auctionContract = new ethers.Contract(AUCTION_ADDRESS, AUCTION_ABI, signer);
const nftContract = new ethers.Contract(NFT_ADDRESS, NFT_ABI, signer);

// 创建拍卖
async function createAuction(tokenId, duration, reservePrice) {
    // 1. 授权NFT
    await nftContract.setApprovalForAll(AUCTION_ADDRESS, true);

    // 2. 创建拍卖
    const tx = await auctionContract.createAuction(
        NFT_ADDRESS,
        tokenId,
        duration,
        ethers.utils.parseEther(reservePrice.toString())
    );

    const receipt = await tx.wait();
    const auctionId = receipt.events[1].args.auctionId;
    return auctionId;
}

// 参与竞价
async function placeBid(auctionId, bidAmount) {
    const tx = await auctionContract.placeBid(auctionId, {
        value: ethers.utils.parseEther(bidAmount.toString())
    });
    return await tx.wait();
}

// 监听事件
auctionContract.on("BidPlaced", (auctionId, bidder, amount) => {
    console.log(`新出价: 拍卖${auctionId}, 出价者${bidder}, 金额${ethers.utils.formatEther(amount)} ETH`);
});
```

## 🧪 测试框架

### 测试分类

#### 1. 单元测试
```bash
# 运行所有单元测试
forge test

# 运行特定合约测试
forge test --match-contract SimpleNFTTest
forge test --match-contract SimpleAuctionTest

# 运行特定函数测试
forge test --match-test testCreateAuction
```

#### 2. 集成测试
```bash
# 运行集成测试
forge test --match-contract AuctionIntegrationTest

# 包含gas报告
forge test --gas-report
```

#### 3. 升级测试
```bash
# 代理基础测试
forge test --match-contract BasicProxyTest

# 升级功能测试
forge test --match-contract ProxyUpgradeTest

# 本地升级测试
make test-upgrade-local
```

#### 4. E2E测试
```bash
# 完整E2E测试流程
make test-e2e

# 手动分阶段测试
make test-e2e-manual
```

### 测试覆盖率

```bash
# 生成覆盖率报告
forge coverage

# 生成详细HTML报告
forge coverage --report lcov
genhtml lcov.info --output-directory coverage

# 查看覆盖率
open coverage/index.html
```

### 模糊测试

```bash
# 运行模糊测试
forge test --fuzz-runs 1000

# 运行不变性测试
forge test --invariant-runs 100
```

## 📍 合约地址

### Sepolia测试网部署 (当前)

```bash
# 核心合约
SimpleNFT            = 0x6d61687CDE7F12A9c31aD833b305EF0c65EA830b
AuctionPriceFeed     = 0x15c8CED44bbdc1fec603BB9147EA51Dcc0809d7a

# 代理架构
TransparentUpgradeableProxy = 0x687265FBABe670a18D0274478984D6c9a03CCCb6
ProxyAdmin                  = 0x17D768939362FdfE8e3EC586A15e258E270d7BE7
Implementation_V1           = 0x9b1e5223Cb5BDb82bA05F75130330c55f3445169
Implementation_V2           = 0x4e8229e669b3c45adFa6Ce1421d002e013EdB122

# 用户交互地址 (代理合约)
AuctionContract = 0x687265FBABe670a18D0274478984D6c9a03CCCb6
```

### 区块链浏览器链接

- **NFT合约**: https://sepolia.etherscan.io/address/0x6d61687CDE7F12A9c31aD833b305EF0c65EA830b
- **拍卖合约**: https://sepolia.etherscan.io/address/0x687265FBABe670a18D0274478984D6c9a03CCCb6
- **价格Feed**: https://sepolia.etherscan.io/address/0x15c8CED44bbdc1fec603BB9147EA51Dcc0809d7a

## 🔄 E2E测试

### 自动化E2E测试

详细E2E测试流程请参考: [E2E_TESTING.md](./E2E_TESTING.md)

```bash
# 完整6阶段E2E测试
make test-e2e

# 执行结果验证
Phase 1: ✅ NFT创建和拍卖设置
Phase 2: ✅ 竞价流程 (4轮竞价)
Phase 3: ✅ Chainlink价格Feed集成
Phase 4: ✅ 平台钱包管理
Phase 5: ✅ 拍卖结算和资金分配
Phase 6: ✅ 合约升级验证
```

### 测试场景覆盖

- **正常流程**: 创建→竞价→结算→领取
- **边界条件**: 最小竞价增幅、时间边界、权限验证
- **异常处理**: 重入攻击、无效参数、授权失败
- **升级场景**: 数据保持、功能扩展、权限控制

## 🔧 升级指南

### 合约升级流程

详细升级测试请参考: [UPGRADE_TESTING.md](./UPGRADE_TESTING.md)

#### 1. 准备V2实现合约
```bash
# 部署新实现合约
forge script script/upgrade/DeployV2Implementation.s.sol \
  --rpc-url $SEPOLIA_RPC_URL \
  --private-key $PRIVATE_KEY \
  --broadcast
```

#### 2. 执行升级
```bash
# 通过ProxyAdmin升级
cast send $PROXY_ADMIN "upgradeAndCall(address,address,bytes)" \
  $AUCTION_PROXY $NEW_IMPLEMENTATION "0x" \
  --private-key $PRIVATE_KEY \
  --rpc-url $SEPOLIA_RPC_URL
```

#### 3. 验证升级
```bash
# 检查版本号
make verify-contract-version

# 测试新功能
make verify-emergency-pause

# 验证数据完整性
cast call $AUCTION_CONTRACT "auctionCounter()" --rpc-url $SEPOLIA_RPC_URL
```

### 升级安全要点

- ✅ **存储布局兼容** - 新变量只能追加，不能修改现有变量
- ✅ **权限验证** - 只有ProxyAdmin owner可以执行升级
- ✅ **数据保持** - 所有历史数据在升级后完整保留
- ✅ **回滚能力** - 可以回滚到之前的实现版本

## 🔧 开发指南

### 开发环境设置

```bash
# 安装开发依赖
npm install -g @openzeppelin/contracts
npm install -g @chainlink/contracts

# 配置IDE (VSCode推荐插件)
- Solidity
- Hardhat Solidity
- GitLens
```

### 代码规范

#### 1. 命名规范
```solidity
// 合约名: PascalCase
contract SimpleAuction {}

// 函数名: camelCase
function createAuction() {}

// 变量名: camelCase
uint256 auctionCounter;

// 常量: UPPER_SNAKE_CASE
uint256 constant MIN_DURATION = 120;

// 私有变量: _前缀
uint256 private _status;
```

#### 2. 注释规范
```solidity
/**
 * @title 合约标题
 * @notice 合约功能简述
 * @dev 开发者注释
 */
contract Example {
    /**
     * @notice 函数功能描述
     * @param param1 参数1说明
     * @param param2 参数2说明
     * @return 返回值说明
     */
    function exampleFunction(uint256 param1, address param2)
        external
        returns (uint256)
    {
        // 实现逻辑
    }
}
```

#### 3. 安全最佳实践
- 总是使用`nonReentrant`修饰器防止重入攻击
- 使用`require`进行参数验证
- 遵循"Checks-Effects-Interactions"模式
- 避免使用`tx.origin`，使用`msg.sender`
- 正确处理整数溢出 (Solidity 0.8+自动检查)

### 贡献指南

1. **Fork项目** 到你的GitHub账户
2. **创建功能分支** `git checkout -b feature/new-feature`
3. **编写测试** 确保新功能有充分的测试覆盖
4. **运行测试** `make test` 确保所有测试通过
5. **提交代码** `git commit -m "feat: add new feature"`
6. **推送分支** `git push origin feature/new-feature`
7. **创建Pull Request** 详细描述变更内容

### 调试技巧

#### 1. 使用console.log调试
```solidity
import "forge-std/console.sol";

function debugFunction() external {
    console.log("Debug: auction counter =", auctionCounter);
    console.log("Debug: msg.sender =", msg.sender);
}
```

#### 2. 事件调试
```solidity
event Debug(string message, uint256 value);

function debugWithEvent() external {
    emit Debug("Checkpoint 1", block.timestamp);
}
```

#### 3. 本地网络调试
```bash
# 启动本地节点 (详细日志)
anvil --host 0.0.0.0 --port 8545 -v

# 使用cast调试
cast call $CONTRACT "function()" --rpc-url http://localhost:8545
```

## 🚨 故障排除

### 常见问题及解决方案

#### 1. 编译错误

**问题**: `ParserError: Source not found`
```bash
# 解决方案: 更新依赖
forge update
forge install
```

**问题**: `DeclarationError: Identifier not found`
```bash
# 解决方案: 检查导入路径
forge remappings > remappings.txt
```

#### 2. 部署失败

**问题**: `insufficient funds for intrinsic transaction cost`
```bash
# 解决方案: 获取更多测试ETH
# 访问 https://sepoliafaucet.com/
```

**问题**: `nonce too low`
```bash
# 解决方案: 重置账户nonce
cast nonce $YOUR_ADDRESS --rpc-url $SEPOLIA_RPC_URL
```

#### 3. 交易失败

**问题**: `execution reverted: NFTNotApproved`
```bash
# 解决方案: 检查NFT授权
cast call $NFT_CONTRACT "isApprovedForAll(address,address)" $OWNER $AUCTION_CONTRACT --rpc-url $SEPOLIA_RPC_URL

# 如果返回false，执行授权
cast send $NFT_CONTRACT "setApprovalForAll(address,bool)" $AUCTION_CONTRACT true --private-key $PRIVATE_KEY --rpc-url $SEPOLIA_RPC_URL
```

**问题**: `execution reverted: InsufficientBid`
```bash
# 解决方案: 检查最低竞价要求
cast call $AUCTION_CONTRACT "getAuction(uint256)" 0 --rpc-url $SEPOLIA_RPC_URL
# 新竞价需要比当前最高价高出至少5%
```

#### 4. 升级问题

**问题**: `ERC1967InvalidImplementation`
```bash
# 解决方案: 检查实现合约
cast code $IMPLEMENTATION_ADDRESS --rpc-url $SEPOLIA_RPC_URL

# 如果返回0x，说明合约部署失败，需要重新部署
```

**问题**: `OwnableUnauthorizedAccount`
```bash
# 解决方案: 检查ProxyAdmin所有者
cast call $PROXY_ADMIN "owner()" --rpc-url $SEPOLIA_RPC_URL

# 确保使用正确的私钥
```

### 环境检查清单

```bash
# 1. 检查Foundry版本
forge --version
cast --version

# 2. 检查网络连接
cast block-number --rpc-url $SEPOLIA_RPC_URL

# 3. 检查余额
cast balance $YOUR_ADDRESS --rpc-url $SEPOLIA_RPC_URL

# 4. 检查合约状态
cast call $AUCTION_CONTRACT "getVersion()" --rpc-url $SEPOLIA_RPC_URL

# 5. 检查权限
cast call $PROXY_ADMIN "owner()" --rpc-url $SEPOLIA_RPC_URL
```

### 紧急操作

#### 暂停合约 (V2版本)
```bash
# 启用紧急暂停
cast send $AUCTION_CONTRACT "toggleEmergencyPause()" \
  --private-key $OWNER_PRIVATE_KEY \
  --rpc-url $SEPOLIA_RPC_URL

# 检查暂停状态
cast call $AUCTION_CONTRACT "isEmergencyPaused()" --rpc-url $SEPOLIA_RPC_URL
```

#### 回滚升级
```bash
# 回滚到之前版本
cast send $PROXY_ADMIN "upgradeAndCall(address,address,bytes)" \
  $AUCTION_PROXY $OLD_IMPLEMENTATION "0x" \
  --private-key $PRIVATE_KEY \
  --rpc-url $SEPOLIA_RPC_URL
```

## 📞 支持与社区

### 获取帮助

- **文档**: 查阅本README和相关文档
- **测试**: 运行`make test`检查问题
- **日志**: 查看Foundry和cast的详细输出
- **社区**: Foundry Discord, OpenZeppelin Forum

### 报告问题

1. 使用`make help`查看可用命令
2. 运行相关测试重现问题
3. 提供详细的错误日志
4. 描述期望行为和实际行为差异
