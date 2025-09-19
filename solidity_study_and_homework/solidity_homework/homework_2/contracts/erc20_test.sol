// SPDX-License-Identifier: MIT
pragma solidity ^0.8.10;

// 任务：参考 openzeppelin-contracts/contracts/token/ERC20/IERC20.sol实现一个简单的 ERC20 代币合约。要求：
// 合约包含以下标准 ERC20 功能：
// balanceOf：查询账户余额。
// transfer：转账。
// approve 和 transferFrom：授权和代扣转账。
// 使用 event 记录转账和授权操作。
// 提供 mint 函数，允许合约所有者增发代币。


interface IERC20 {
    function totalSupply() external view returns (uint256);
    function balanceOf(address account) external view returns (uint256);
    function transfer(address to, uint256 amount) external returns (bool);
    function allowance(address owner, address spender) external view returns (uint256);
    function approve(address spender, uint256 amount) external returns (bool);
    function transferFrom(address from, address to, uint256 amount) external returns (bool);
}

contract ERC20Test is IERC20 {
    event Transfer(address indexed from, address indexed to, uint256 amount);
    event Approval(address indexed owner, address indexed spender, uint256 amount);
    mapping(address => uint256) private _balances;
    mapping(address => mapping(address => uint256)) private _allowances;
    uint256 private _totalSupply;
    string public name = "wtwtMoon";
    string public symbol = "WTM";
    uint8 public decimals = 18;
    address public owner;

    constructor(uint256 initialSupply) {
        _balances[msg.sender] = initialSupply;
        _totalSupply = initialSupply;
        owner = msg.sender;
        emit Transfer(address(0), msg.sender, initialSupply);
    }

    function mint(uint256 amount) external {
        require(msg.sender == owner, "Only owner can mint");
        _balances[msg.sender] += amount;
        _totalSupply += amount;
        emit Transfer(address(0), msg.sender, amount);
    }


    function totalSupply() external view override returns (uint256) {
        return _totalSupply;
    }

    function balanceOf(address account) external view override returns (uint256) {
        return _balances[account];
    }

    function allowance(address owner_, address spender_) external view override returns (uint256) {
        return _allowances[owner_][spender_];
    }

    function approve(address spender, uint256 amount) external override returns (bool) {
        require(spender != address(0), "Approve to zero address");
        _allowances[msg.sender][spender] = amount;
        emit Approval(msg.sender, spender, amount);
        return true;
    }

    function transfer(address to, uint256 amount) external override returns (bool) {
        require(to != address(0), "Transfer to zero address");
        require(_balances[msg.sender] >= amount, "Insufficient balance");

        _balances[msg.sender] -= amount;
        _balances[to] += amount;
        emit Transfer(msg.sender, to, amount);
        return true;
    }
    function transferFrom(address from, address to, uint256 amount) external override returns (bool) {
        require(to != address(0), "Transfer to zero address");
        require(_balances[from] >= amount, "Insufficient balance");
        require(_allowances[from][msg.sender] >= amount, "Insufficient allowance");

        _balances[from] -= amount;
        _balances[to] += amount;
        _allowances[from][msg.sender] -= amount;

        emit Transfer(from, to, amount);
        return true;
    }
}
