// SPDX-License-Identifier: MIT
pragma solidity ^0.8.10;

interface IERC20 {
    //转账 由msg.sender向to地址转账amount
    function transfer(address to, uint256 amount) external returns (bool);
    //转账from地址的余额到to地址
    function transferFrom(address from, address to, uint256 amount) external returns (bool);
    //批准 由msg.sender向spender地址批准amount
    function approve(address spender, uint256 amount) external returns (bool);
    //批准额度 查询由owner地址向spender地址批准的额度
    function allowance(address owner, address spender) external view returns (uint256);
    //查询owner地址的余额
    function balanceOf(address owner) external view returns (uint256);
    //查询总供应量
    function totalSupply() external view returns (uint256);
}


contract ERC20Test is IERC20 {
    mapping(address => uint256) public balances;
    mapping(address => mapping(address => uint256)) public allowances;
    uint256 public totalSupply;
    //名字
    string public name = "dogMoon";
    //符号(缩写)
    string public symbol = "DM";
    //小数位
    uint8 public decimals = 18;

    constructor(uint256 initialSupply) {
        balances[msg.sender] = initialSupply;
        totalSupply = initialSupply;
    }
    //转账 由msg.sender向to地址转账amount
    function transfer(address to, uint256 amount) external returns (bool) {
        balances[msg.sender] -= amount;
        balances[to] += amount;
        return true;
    }
    //转账from地址的余额到to地址
    function transferFrom(address from, address to, uint256 amount) external returns (bool) {
        balances[from] -= amount;
        balances[to] += amount;
        allowances[from][msg.sender] -= amount;
        return true;
    }
    //批准 由msg.sender向spender地址批准amount
    function approve(address spender, uint256 amount) external returns (bool) {
        allowances[msg.sender][spender] = amount;
        return true;
    }
    //批准额度 查询由owner地址向spender地址批准的额度
    function allowance(address owner, address spender) external view returns (uint256) {
        return allowances[owner][spender];
    }
    //查询owner地址的余额
    function balanceOf(address owner) external view returns (uint256) {
        return balances[owner];
    }
    //查询总供应量
    function totalSupply() external view returns (uint256) {
        return totalSupply;
    }
}