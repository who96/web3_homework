// SPDX-License-Identifier: MIT
pragma solidity ^0.8.10;

/**
 * @title AbiDecode
 * @dev 演示Solidity中ABI编码和解码的基本用法
 * @notice ABI (Application Binary Interface) 是智能合约与外部世界交互的标准接口
 * @author 基于Linus的"好品味"原则：简单、直接、无特殊情况
 */
contract AbiDecode {
    /**
     * @dev 自定义结构体，用于演示复杂数据类型的编码/解码
     * @notice 结构体包含字符串和固定长度数组，代表典型的业务数据
     */
    struct MyStruct {
        string name;    // 名称字段
        uint[2] nums;   // 固定长度为2的uint数组
    }

    /**
     * @dev 将多个不同类型的参数编码为字节数组
     * @param x 无符号整数
     * @param addr 以太坊地址
     * @param arr 动态长度uint数组
     * @param myStruct 自定义结构体
     * @return 编码后的字节数组，可用于跨合约调用或数据存储
     * 
     * @notice 实际生产应用场景：
     * 1. 跨合约调用：将函数参数编码后通过call()传递给其他合约
     * 2. 事件日志：将复杂数据编码后存储在事件中
     * 3. 数据存储：将多个字段打包存储以节省gas
     * 4. 代理模式：在代理合约中编码调用数据
     */
    function encode(
        uint x,                    // 基础类型：无符号整数
        address addr,              // 基础类型：以太坊地址
        uint[] calldata arr,       // 动态数组：使用calldata节省gas
        MyStruct calldata myStruct // 结构体：使用calldata节省gas
    ) external pure returns (bytes memory) {
        // 使用abi.encode()将多个参数编码为字节数组
        // 这是Solidity内置的序列化函数，遵循ABI标准
        return abi.encode(x, addr, arr, myStruct);
    }

    /**
     * @dev 将编码的字节数组解码为原始参数
     * @param data 待解码的字节数组
     * @return x 解码后的无符号整数
     * @return addr 解码后的以太坊地址
     * @return arr 解码后的动态uint数组
     * @return myStruct 解码后的自定义结构体
     * 
     * @notice 实际生产应用场景：
     * 1. 接收跨合约调用：解码其他合约传递的参数
     * 2. 事件解析：从链下解析事件中的编码数据
     * 3. 数据恢复：从存储中恢复打包的数据
     * 4. 代理模式：在实现合约中解码代理传递的调用数据
     * 5. 批量操作：解码包含多个操作的批量数据
     */
    function decode(bytes calldata data)
        external
        pure
        returns (
            uint x,                    // 返回值：无符号整数
            address addr,              // 返回值：以太坊地址
            uint[] memory arr,         // 返回值：动态数组（使用memory）
            MyStruct memory myStruct   // 返回值：结构体（使用memory）
        )
    {
        // 使用abi.decode()将字节数组解码为指定类型
        // 必须指定完整的类型元组，顺序必须与编码时一致
        // 注释掉的代码展示了另一种声明方式
        // (uint x, address addr, uint[] memory arr, MyStruct myStruct) = ...
        (x, addr, arr, myStruct) = abi.decode(data, (uint, address, uint[], MyStruct));
    }
}
