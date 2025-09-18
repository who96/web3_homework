// SPDX-License-Identifier: MIT
pragma solidity ^0.8.10;

/**
 * 多签钱包合约 - 需要多个授权签名才能执行交易
 * 核心设计理念：安全性 > 便利性，防止单点故障
 */
contract MultiSigWallet {
    // 事件定义 - 用于链下监听和日志记录
    event Deposit(address indexed sender, uint amount, uint balance);  // 存款事件
    event SubmitTransaction(  // 提交交易事件
        address indexed owner,
        uint indexed txIndex,
        address indexed to,
        uint valtg65ttttttdeeeeeeeeeeeeeerttue,
        bytes data
    );
    event ConfirmTransaction(address indexed owner, uint indexed txIndex);  // 确认交易事件
    event RevokeConfirmation(address indexed owner, uint indexed txIndex); // 撤销确认事件
    event ExecuteTransaction(address indexed owner, uint indexed txIndex);  // 执行交易事件

    // 核心状态变量 - 这是整个合约的数据基础
    address[] public  owners;                    // 所有者地址数组 - 用于遍历
    mapping(address => bool) public isOwner;    // 地址到所有者的快速查找映射
    uint public numConfirmationsRequired;       // 执行交易需要的最小确认数 
    // 交易数据结构 - 封装所有交易相关信息
    struct Transaction {
        address to;             // 目标地址
        uint value;             // 转账金额(wei)
        bytes data;             // 调用数据(用于合约调用)
        bool executed;          // 执行状态标志
        uint numConfirmations;  // 当前确认数量
    }

    // 双重映射：交易索引 => 所有者地址 => 是否已确认
    // 这种设计避免了遍历查找，O(1)时间复杂度
    mapping(uint => mapping(address => bool)) public isConfirmed;
    
    Transaction[] public transactions;  // 所有交易的动态数组

    // 修饰符设计 - 用于消除函数内部的重复检查逻辑
    // 这是"好品味"的体现：把边界检查提取出来，让主逻辑更清晰
    modifier onlyOwner() {
        require(isOwner[msg.sender], "not owner");
        _;
    }

    modifier txExists(uint _txIndex) {
        require(_txIndex < transactions.length, "tx does not exist");
        _;
    }

    modifier notExecuted(uint _txIndex) {
        require(!transactions[_txIndex].executed, "tx already executed");
        _;
    }

    modifier notConfirmed(uint _txIndex) {
        require(!isConfirmed[_txIndex][msg.sender], "tx already confirmed");
        _;
    }

    /**
     * 构造函数 - 初始化多签钱包
     * @param _owners 所有者地址数组
     * @param _numConfirmationsRequired 执行交易需要的最小确认数
     */
    constructor(address[] memory _owners, uint _numConfirmationsRequired) {
        require(_owners.length > 0, "owners required");
        require(
            _numConfirmationsRequired > 0 &&
                _numConfirmationsRequired <= _owners.length,
            "invalid number of required confirmations"
        );

        // 初始化所有者列表 - 同时维护数组和映射两种数据结构
        for (uint i = 0; i < _owners.length; i++) {
            address owner = _owners[i];

            require(owner != address(0), "invalid owner");        // 防止零地址
            require(!isOwner[owner], "owner not unique");       // 防止重复地址

            isOwner[owner] = true;   // 设置快速查找映射
            owners.push(owner);      // 添加到数组用于遍历
        }

        numConfirmationsRequired = _numConfirmationsRequired;
    }

    // 接收以太币的回退函数 - 让钱包可以接收直接转账
    receive() external payable {
        emit Deposit(msg.sender, msg.value, address(this).balance);
    }

    /**
     * 提交新交易 - 多签流程的第一步
     * 只是创建交易，不执行。需要后续确认和执行步骤
     */
    function submitTransaction(
        address _to,
        uint _value,
        bytes memory _data
    ) public onlyOwner {
        uint txIndex = transactions.length;  // 使用数组长度作为新交易的索引

        transactions.push(
            Transaction({
                to: _to,
                value: _value,
                data: _data,
                executed: false,        // 初始状态：未执行
                numConfirmations: 0     // 初始确认数为0
            })
        );

        emit SubmitTransaction(msg.sender, txIndex, _to, _value, _data);
    }

    /**
     * 确认交易 - 多签流程的第二步
     * 每个所有者对交易进行确认，增加确认计数
     */
    function confirmTransaction(uint _txIndex)
        public
        onlyOwner           // 只有所有者可以确认
        txExists(_txIndex)  // 交易必须存在
        notExecuted(_txIndex)  // 交易不能已经执行
        notConfirmed(_txIndex) // 该所有者不能已经确认过
    {
        Transaction storage transaction = transactions[_txIndex];
        transaction.numConfirmations += 1;              // 增加确认计数
        isConfirmed[_txIndex][msg.sender] = true;       // 记录该所有者已确认

        emit ConfirmTransaction(msg.sender, _txIndex);
    }

    /**
     * 执行交易 - 多签流程的最后一步
     * 当确认数达到要求时，实际执行交易
     */
    function executeTransaction(uint _txIndex)
        public
        onlyOwner
        txExists(_txIndex)
        notExecuted(_txIndex)
    {
        Transaction storage transaction = transactions[_txIndex];

        // 检查确认数是否足够 - 这是多签的核心安全检查
        require(
            transaction.numConfirmations >= numConfirmationsRequired,
            "cannot execute tx"
        );

        transaction.executed = true;  // 标记为已执行，防止重复执行

        // 执行实际的交易调用 - 可能是转账或合约调用
        (bool success, ) = transaction.to.call{value: transaction.value}(
            transaction.data
        );
        require(success, "tx failed");  // 交易失败则回滚整个操作

        emit ExecuteTransaction(msg.sender, _txIndex);
    }

    /**
     * 撤销确认 - 允许所有者改变主意
     * 在交易执行前，所有者可以撤回自己的确认
     */
    function revokeConfirmation(uint _txIndex)
        public
        onlyOwner
        txExists(_txIndex)
        notExecuted(_txIndex)  // 只能撤销未执行的交易确认
    {
        Transaction storage transaction = transactions[_txIndex];

        require(isConfirmed[_txIndex][msg.sender], "tx not confirmed");  // 必须之前已确认

        transaction.numConfirmations -= 1;              // 减少确认计数
        isConfirmed[_txIndex][msg.sender] = false;      // 清除确认状态

        emit RevokeConfirmation(msg.sender, _txIndex);
    }

    // 查询函数 - 提供对内部状态的只读访问
    
    /** 获取所有所有者地址 */
    function getOwners() public view returns (address[] memory) {
        return owners;
    }

    /** 获取交易总数 */
    function getTransactionCount() public view returns (uint) {
        return transactions.length;
    }

    /** 
     * 获取指定交易的详细信息
     * 返回交易的所有字段，供前端显示使用
     */
    function getTransaction(uint _txIndex)
        public
        view
        returns (
            address to,
            uint value,
            bytes memory data,
            bool executed,
            uint numConfirmations
        )
    {
        Transaction storage transaction = transactions[_txIndex];

        return (
            transaction.to,
            transaction.value,
            transaction.data,
            transaction.executed,
            transaction.numConfirmations
        );
    }
}
