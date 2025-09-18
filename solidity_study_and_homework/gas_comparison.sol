// SPDX-License-Identifier: MIT
pragma solidity ^0.8.10;

// 普通多次调用的示例
contract NormalCalls {
    function callThreeContracts() external {
        // 需要3个交易
        // 每个交易的基础费用：21,000 gas
        // 总基础费用：63,000 gas
    }
}

// 多重调用的示例
contract MultiCall {
    function multiCall(address[] calldata targets, bytes[] calldata data)
        external
        view
        returns (bytes[] memory)
    {
        require(targets.length == data.length, "target length != data length");
        
        bytes[] memory results = new bytes[](data.length);
        
        for (uint i; i < targets.length; i++) {
            (bool success, bytes memory result) = targets[i].staticcall(data[i]);
            require(success, "call failed");
            results[i] = result;
        }
        
        return results;
    }
}

// Gas消耗对比
contract GasComparison {
    // 普通多次调用：3个交易
    // 交易1：21,000 + 函数调用gas
    // 交易2：21,000 + 函数调用gas  
    // 交易3：21,000 + 函数调用gas
    // 总计：63,000 + 函数调用gas
    
    // 多重调用：1个交易
    // 交易1：21,000 + 函数调用gas + 循环开销
    // 总计：21,000 + 函数调用gas + 循环开销
    
    // 节省：42,000 gas - 循环开销
    // 循环开销通常 < 1,000 gas
    // 净节省：约 41,000 gas
}
