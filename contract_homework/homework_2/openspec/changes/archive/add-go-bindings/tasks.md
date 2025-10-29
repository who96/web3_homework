## Go 绑定生成与脚本验证（已完成）
- [x] 1. 导出 `FukuaStake` 合约 ABI / 字节码供 `abigen` 使用
- [x] 2. 安装或确认 `abigen` 可用，生成 Go 绑定代码
- [x] 3. 编写 Go 脚本：
  - 连接 Sepolia，调用 `depositEth` 质押 1 ETH，并在等待 10 区块后调用 `claim`
  - 启动 goroutine 订阅 `Deposit` / `Claim` / `RequestUnstake` / `Withdraw` 事件，输出“用户/操作/币种/金额/资金流向”信息
- [x] 4. 运行脚本输出结果，记录交互调用与事件监听日志
