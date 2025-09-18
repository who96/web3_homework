// SPDX-License-Identifier: MIT
pragma solidity ^0.8.10;

contract Receiver {
    event Log(address to, uint256 amount, bytes data);

    function gerSelector(string calldata funcName) public pure returns (bytes4) {
        return bytes4(keccak256(bytes(funcName)));
    }
    
    function  transfer(address to, uint256 amount) public {
        emit Log(to, amount, msg.data);
        //0xa9059cbb
        //000000000000000000000000ab8483f64d9c6d1ecf9b849ae677dd3315835cb2
        //000000000000000000000000000000000000000000000000000000000000000b
    }
}