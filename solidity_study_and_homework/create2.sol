// This is the older way of doing it using assembly
  // 下面是使用内联汇编的旧方式
  contract FactoryAssembly {                       // 合约：FactoryAssembly，通过内联汇编手动执行 CREATE2
      event Deployed(address addr, uint salt);     // 事件：记录部署出的地址和 salt（此处类型为 uint，建议与 bytes32 统一）

      // 1. Get bytecode of contract to be deployed
      // NOTE: _owner and _foo are arguments of the TestContract's constructor
      // 1. 获取待部署合约的字节码；_owner/_foo 是 TestContract 构造参数
      function getBytecode(address _owner, uint _foo) public pure returns (bytes memory) {
          bytes memory bytecode = type(TestContract).creationCode; // 取得 TestContract 的创建字节码（不含构造参数）

          return abi.encodePacked(bytecode, abi.encode(_owner, _foo)); // 拼接构造参数，得到完整部署字节码
      }

      // 2. Compute the address of the contract to be deployed
      // NOTE: _salt is a random number used to create an address
      // 2. 预测将要部署的合约地址；_salt 参与地址计算（应为 32 字节值）
      function getAddress(bytes memory bytecode, uint _salt)
          public
          view
          returns (address)
      {
          bytes32 hash = keccak256(
              abi.encodePacked(bytes1(0xff), address(this), _salt, keccak256(bytecode))
          );                                       // 按 CREATE2 公式计算哈希：keccak256(0xff ++ 部署者 ++ salt ++ keccak256(bytecode))

          // NOTE: cast last 20 bytes of hash to address
          // 取哈希的后 20 字节作为地址
          return address(uint160(uint(hash)));
      }

      // 3. Deploy the contract
      // NOTE:
      // Check the event log Deployed which contains the address of the deployed TestContract.
      // The address in the log should equal the address computed from above.
      // 3. 实际部署合约
      // 注意：通过 Deployed 事件中的地址与 getAddress 计算值进行核对，应完全一致
      function deploy(bytes memory bytecode, uint _salt) public payable {
          address addr;                             // 用于接收新合约地址

          /*
          NOTE: How to call create2
          // 说明：CREATE2 的调用方式和地址计算规则

          create2(v, p, n, s)
          create new contract with code at memory p to p + n
          and send v wei
          and return the new address
          where new address = first 20 bytes of keccak256(0xff + address(this) + s + keccak256(mem[p…(p+n)))
                s = big-endian 256-bit value

          // 中文：create2(v, p, n, s)
          // 在内存 [p, p+n) 位置的代码被用来创建新合约，并转发 v wei
          // 新地址 = keccak256(0xff ++ address(this) ++ s ++ keccak256(mem[p…p+n])) 的后 20 字节
          // 其中 s 是 256-bit（大端）盐值
          */
          assembly {
              addr := create2(
                  callvalue(),            // 转发当前调用携带的以太（msg.value）
                  add(bytecode, 0x20),    // 跳过前 32 字节长度头，指向真实代码起始位置
                  mload(bytecode),        // 从前 32 字节读取代码长度
                  _salt                   // salt 参数（与 getAddress 中的一致）
              )

              if iszero(extcodesize(addr)) { // 如果部署失败或地址处没有代码（0），回退
                  revert(0, 0)
              }
          }

          emit Deployed(addr, _salt);         // 发出事件，便于 off-chain 索引与确认
      }
  }

  contract TestContract {                      // 被部署的目标合约
      address public owner;                    // 状态变量：owner（公开可读）
      uint public foo;                         // 状态变量：foo（公开可读）

      constructor(address _owner, uint _foo) payable {
          owner = _owner;                      // 构造时设置 owner
          foo = _foo;                          // 构造时设置 foo
      }

      function getBalance() public view returns (uint) {
          return address(this).balance;        // 返回合约自身余额
      }
  }