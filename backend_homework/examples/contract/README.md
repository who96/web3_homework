# SimpleStorage 合约交互示例

这是一个完整的智能合约交互示例，展示如何使用 `abigen` 生成的 Go 绑定与 Sepolia 测试网上的 SimpleStorage 合约交互。

## 快速开始

### 1. 生成 Go 绑定代码

```bash
make abigen
```

这会：
- 编译 `SimpleStorage.sol` 合约
- 提取 ABI 和 Bytecode
- 生成 `contracts/bindings/SimpleStorage.go`

### 2. 部署合约到 Sepolia

```bash
make contract-deploy
```

输出示例：
```
=== 部署 SimpleStorage 合约 ===
从地址: 0x5D4a329B262AC7a1d9Ae0f4C54171dF61E2c0b69
初始值: 42
Gas Price: 1000000000 wei

正在部署...

✅ 合约部署成功！
合约地址: 0x1234567890abcdef1234567890abcdef12345678
交易哈希: 0xabcd...

📝 请保存合约地址用于后续交互: 0x1234...
```

**重要：保存合约地址！**

### 3. 与合约交互

使用步骤2中获得的合约地址：

```bash
make contract-interact ADDR=0x1234567890abcdef1234567890abcdef12345678
```

这会执行以下操作：
1. 读取当前存储值
2. 调用 `increment()` 增加计数器
3. 读取更新后的值
4. 解析 `ValueChanged` 事件

输出示例：
```
=== SimpleStorage 合约交互演示 ===
合约地址: 0x1234...
调用者地址: 0x5D4a...

--- 1. 读取当前存储值 ---
当前值: 42

--- 2. 调用 increment() 增加计数器 ---
交易已发送: 0xdef...
等待交易确认...
✅ 交易已确认 (区块 12345, Gas 使用: 50000)

--- 3. 读取更新后的值 ---
新值: 43
变化: 42 -> 43 (+1)

--- 4. 解析 ValueChanged 事件 ---
事件列表:
  事件 #1:
    旧值: 42
    新值: 43
    调用者: 0x5D4a329B262AC7a1d9Ae0f4C54171dF61E2c0b69
    交易: 0xdef...

✅ 交互演示完成！
```

## 手动运行（不使用 Makefile）

### 部署

```bash
cd examples/contract
go run deploy.go
```

### 交互

```bash
cd examples/contract
go run interact.go 0x<合约地址>
```

## 文件说明

- **deploy.go** - 部署合约到 Sepolia 测试网
  - 加载私钥和RPC配置
  - 创建交易签名器
  - 部署合约（初始值=42）
  - 等待交易确认
  - 验证部署成功

- **interact.go** - 与已部署合约交互
  - 读取当前值（view 函数，不消耗gas）
  - 调用 increment()（写入函数，需要gas）
  - 读取更新后的值
  - 解析 ValueChanged 事件日志

## 合约方法

SimpleStorage 合约提供以下方法：

```solidity
// 构造函数
constructor(uint256 _initialValue)

// 读取函数（免费，不需要gas）
function get() public view returns (uint256)

// 写入函数（需要gas）
function set(uint256 _value) public
function increment() public

// 事件
event ValueChanged(uint256 oldValue, uint256 newValue, address indexed changer)
```

## Go 绑定使用示例

### 部署合约

```go
auth := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
initialValue := big.NewInt(42)

address, tx, instance, err := bindings.DeploySimpleStorage(auth, client, initialValue)
```

### 连接已部署合约

```go
contractAddress := common.HexToAddress("0x...")
instance, err := bindings.NewSimpleStorage(contractAddress, client)
```

### 调用只读方法

```go
value, err := instance.Get(&bind.CallOpts{})
fmt.Println("Current value:", value)
```

### 调用写入方法

```go
auth := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
tx, err := instance.Increment(auth)
receipt, err := bind.WaitMined(context.Background(), client, tx)
```

### 监听事件

```go
filterOpts := &bind.FilterOpts{
    Start: startBlock,
    End:   &endBlock,
}

iter, err := instance.FilterValueChanged(filterOpts, nil)
for iter.Next() {
    event := iter.Event
    fmt.Println("Old:", event.OldValue)
    fmt.Println("New:", event.NewValue)
    fmt.Println("Changer:", event.Changer.Hex())
}
```

## 前置要求

1. **Sepolia 测试网 ETH** - 需要少量 SepoliaETH 用于支付 gas
   - 从水龙头获取: https://sepoliafaucet.com/

2. **配置 .env 文件** - 确保以下环境变量已设置：
   ```
   PRIVATE_KEY=0x...
   SEPOLIA_RPC_URL=https://ethereum-sepolia-rpc.publicnode.com
   ```

3. **安装依赖**：
   ```bash
   go mod tidy
   ```

## 故障排除

### 编译错误："main redeclared"

这是正常的！`deploy.go` 和 `interact.go` 都有 `main` 函数，不要同时编译它们。
分别运行：
- `go run deploy.go`
- `go run interact.go <address>`

### Gas 不足

增加 `auth.GasLimit`:
```go
auth.GasLimit = uint64(300000)  // 增加到 300k
```

### RPC 错误

尝试使用不同的 RPC 节点：
- https://rpc.sepolia.org
- https://ethereum-sepolia-rpc.publicnode.com
- https://sepolia.infura.io/v3/YOUR_KEY

## 相关资源

- **abigen 文档**: https://geth.ethereum.org/docs/tools/abigen
- **go-ethereum 文档**: https://pkg.go.dev/github.com/ethereum/go-ethereum
- **Sepolia 浏览器**: https://sepolia.etherscan.io/
