// SPDX-License-Identifier: MIT
pragma solidity ^0.8.10;

// 多重委托调用合约 - 允许批量执行多个函数调用
contract MultiDelegatecall {
    // 自定义错误：委托调用失败
    error DelegatecallFailed();

    // 批量执行委托调用的主函数
    // data: 包含多个函数调用数据的字节数组
    // 返回值: 每个函数调用的返回结果数组
    function multiDelegatecall(bytes[] memory data)
        external
        payable
        returns (bytes[] memory results)
    {
        // 初始化结果数组，长度与输入数据数组相同
        results = new bytes[](data.length);

        // 循环执行每个函数调用
        for (uint i; i < data.length; i++) {
            // 对当前合约执行委托调用
            // delegatecall会使用当前合约的存储和上下文，但执行目标合约的代码
            // 关键：msg.sender保持不变，仍然是原始调用者
            (bool ok, bytes memory res) = address(this).delegatecall(data[i]);
            
            // 如果委托调用失败，回滚整个交易
            if (!ok) {
                revert DelegatecallFailed();
            }
            
            // 保存当前调用的返回结果
            results[i] = res;
        }
    }
}

// 为什么使用多重委托调用而不是多重普通调用？

// 普通调用：alice -> multi call --- call ---> test (msg.sender = multi call)
// 委托调用：alice -> test --- delegatecall ---> test (msg.sender = alice)
// 关键区别：委托调用保持原始调用者的身份，普通调用会改变msg.sender

// 测试多重委托调用的合约
contract TestMultiDelegatecall is MultiDelegatecall {
    // 事件：记录调用者、函数名和参数
    event Log(address caller, string func, uint i);

    // 测试函数1：简单的数学运算
    function func1(uint x, uint y) external {
        // 注意：这里的msg.sender仍然是原始调用者(alice)，不是MultiDelegatecall合约
        emit Log(msg.sender, "func1", x + y);
    }

    // 测试函数2：返回固定值
    function func2() external returns (uint) {
        // msg.sender仍然是原始调用者(alice)
        emit Log(msg.sender, "func2", 2);
        return 111;
    }

    // 余额映射：记录每个地址的代币余额
    mapping(address => uint) public balanceOf;

    // 警告：当与多重委托调用结合使用时，此代码不安全
    // 用户可以用一次msg.value的价格进行多次铸造
    function mint() external payable {
        // 危险：msg.value在整个批量操作中只支付一次，但mint可能被调用多次
        balanceOf[msg.sender] += msg.value;
    }
}

// 辅助合约：用于生成函数调用数据
contract Helper {
    // 生成func1函数的调用数据
    // 用于构造multiDelegatecall所需的bytes[]参数
    function getFunc1Data(uint x, uint y) external pure returns (bytes memory) {
        return abi.encodeWithSelector(TestMultiDelegatecall.func1.selector, x, y);
    }

    // 生成func2函数的调用数据
    function getFunc2Data() external pure returns (bytes memory) {
        return abi.encodeWithSelector(TestMultiDelegatecall.func2.selector);
    }

    // 生成mint函数的调用数据
    // 注意：这个函数存在安全风险，仅用于演示
    function getMintData() external pure returns (bytes memory) {
        return abi.encodeWithSelector(TestMultiDelegatecall.mint.selector);
    }
}
