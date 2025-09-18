// SPDX-License-Identifier: MIT
pragma solidity ^0.8.10;

contract PiggyBank {
    address public owner;
    uint256 public balance;

    event Deposit(address indexed depositor, uint256 amount);
    event Withdraw(address indexed owner, uint256 amount);
    
    constructor() payable {
        owner = msg.sender;
        balance = msg.value;
    }
    
    receive() external payable {
        emit Deposit(msg.sender, msg.value);  // 现在可以正常使用
        balance += msg.value;
    }
    
    function withdraw(uint256 amount) public {
        require(msg.sender == owner, "Only owner can withdraw");
        (bool ok, ) = address(owner).call{value: amount}("");
        require(ok, "Withdraw failed");
        balance -= amount;
        emit Withdraw(msg.sender, amount);  // 也可以记录提款
    }
    function getBalance() public view returns (uint256) {
        return balance;
    }

    function getOwner() public view returns (address) {
        return owner;
    }
}