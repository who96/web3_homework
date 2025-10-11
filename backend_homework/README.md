# Go-Ethereum 交易生命周期学习项目

使用 go-ethereum 与 Sepolia 测试链交互，深入理解交易处理流程。

## 核心功能

### 1. 交易生命周期演示
完整展示一笔交易的五个阶段：
- **构建交易**：设置 nonce、gas price、gas limit
- **钱包签名**：使用私钥对交易进行 EIP155 签名
- **交易广播**：发送到网络节点
- **矿工打包**：等待被包含到区块
- **区块广播与执行**：交易执行并获取回执

### 2. 基础功能
- 查询账户余额
- 订阅新区块
- 监控交易状态

## 项目结构

```
.
├── config/         # 配置管理
├── tx/            # 交易管理核心代码
├── contracts/     # Foundry 智能合约测试
│   ├── src/       # 合约源码
│   └── test/      # 合约测试
└── main.go        # 入口文件
```

## 运行

### 查询余额
```bash
go run main.go
```

### 运行 Foundry 测试
```bash
make forge-test
```

### 构建项目
```bash
make build
```

## 交易处理流程

1. **私钥 → 地址**：从私钥推导出公钥和地址
2. **构建原始交易**：设置接收方、金额、gas 参数
3. **签名**：生成 v, r, s 签名值
4. **广播**：通过 RPC 发送到节点
5. **等待确认**：轮询获取交易回执
6. **验证状态**：检查交易是否成功

## 关键数据结构

- `Transaction`: 包含 nonce、to、value、gas 等字段
- `Receipt`: 包含区块号、gas 使用量、执行状态
- `Block`: 包含区块哈希、时间戳、交易列表

## Makefile 命令

- `make build` - 构建二进制
- `make run` - 运行程序
- `make test` - Go 测试
- `make forge-test` - 合约测试
- `make forge-build` - 编译合约
- `make clean` - 清理构建文件