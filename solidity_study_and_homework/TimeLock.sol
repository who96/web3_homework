// SPDX-License-Identifier: MIT
pragma solidity ^0.8.10;

/**
 * @title TimeLock
 * @dev 简单的时间锁合约，实现延迟执行机制
 * @notice 基于Linus的"好品味"原则：简单、安全、无特殊情况
 * @author 时间锁是区块链安全的基础组件，就像Linux内核中的权限控制
 */
contract TimeLock {
    // 事件定义
    event QueueTransaction(
        bytes32 indexed txHash,
        address indexed target,
        uint256 value,
        string signature,
        bytes data,
        uint256 eta
    );
    
    event ExecuteTransaction(
        bytes32 indexed txHash,
        address indexed target,
        uint256 value,
        string signature,
        bytes data,
        uint256 eta
    );
    
    event CancelTransaction(bytes32 indexed txHash);
    
    // 状态变量
    uint256 public constant MIN_DELAY = 2 days;        // 最小延迟时间
    uint256 public constant MAX_DELAY = 30 days;       // 最大延迟时间
    uint256 public constant GRACE_PERIOD = 14 days;    // 宽限期
    
    address public admin;                              // 管理员地址
    uint256 public delay;                              // 当前延迟时间
    
    mapping(bytes32 => bool) public queuedTransactions; // 待执行交易映射
    
    // 修饰符
    modifier onlyAdmin() {
        require(msg.sender == admin, "TimeLock: caller is not admin");
        _;
    }
    
    modifier onlyThisContract() {
        require(msg.sender == address(this), "TimeLock: caller is not this contract");
        _;
    }
    
    /**
     * @dev 构造函数
     * @param _admin 管理员地址
     * @param _delay 延迟时间（秒）
     */
    constructor(address _admin, uint256 _delay) {
        require(_admin != address(0), "TimeLock: invalid admin");
        require(_delay >= MIN_DELAY, "TimeLock: delay must exceed minimum delay");
        require(_delay <= MAX_DELAY, "TimeLock: delay must not exceed maximum delay");
        
        admin = _admin;
        delay = _delay;
    }
    
    /**
     * @dev 接收以太币
     */
    receive() external payable {}
    
    /**
     * @dev 将交易加入队列
     * @param target 目标合约地址
     * @param value 发送的以太币数量
     * @param signature 函数签名
     * @param data 调用数据
     * @param eta 执行时间戳
     */
    function queueTransaction(
        address target,
        uint256 value,
        string memory signature,
        bytes memory data,
        uint256 eta
    ) external onlyAdmin returns (bytes32) {
        // 验证执行时间
        require(eta >= block.timestamp + delay, "TimeLock: estimated execution time must satisfy delay");
        require(eta <= block.timestamp + delay + GRACE_PERIOD, "TimeLock: execution time too far in future");
        
        // 计算交易哈希
        bytes32 txHash = keccak256(abi.encode(target, value, signature, data, eta));
        
        // 检查交易是否已存在
        require(!queuedTransactions[txHash], "TimeLock: transaction already queued");
        
        // 将交易加入队列
        queuedTransactions[txHash] = true;
        
        emit QueueTransaction(txHash, target, value, signature, data, eta);
        
        return txHash;
    }
    
    /**
     * @dev 执行队列中的交易
     * @param target 目标合约地址
     * @param value 发送的以太币数量
     * @param signature 函数签名
     * @param data 调用数据
     * @param eta 执行时间戳
     */
    function executeTransaction(
        address target,
        uint256 value,
        string memory signature,
        bytes memory data,
        uint256 eta
    ) external onlyAdmin returns (bytes memory) {
        // 计算交易哈希
        bytes32 txHash = keccak256(abi.encode(target, value, signature, data, eta));
        
        // 验证交易在队列中
        require(queuedTransactions[txHash], "TimeLock: transaction not queued");
        
        // 验证执行时间
        require(block.timestamp >= eta, "TimeLock: transaction not yet ready");
        require(block.timestamp <= eta + GRACE_PERIOD, "TimeLock: transaction is stale");
        
        // 从队列中移除交易
        queuedTransactions[txHash] = false;
        
        // 执行交易
        bytes memory returnData;
        if (bytes(signature).length == 0) {
            // 无函数签名，直接调用
            (bool success, bytes memory retData) = target.call{value: value}(data);
            require(success, "TimeLock: transaction execution reverted");
            returnData = retData;
        } else {
            // 有函数签名，构造完整调用数据
            bytes memory callData = abi.encodePacked(bytes4(keccak256(bytes(signature))), data);
            (bool success, bytes memory retData) = target.call{value: value}(callData);
            require(success, "TimeLock: transaction execution reverted");
            returnData = retData;
        }
        
        emit ExecuteTransaction(txHash, target, value, signature, data, eta);
        
        return returnData;
    }
    
    /**
     * @dev 取消队列中的交易
     * @param target 目标合约地址
     * @param value 发送的以太币数量
     * @param signature 函数签名
     * @param data 调用数据
     * @param eta 执行时间戳
     */
    function cancelTransaction(
        address target,
        uint256 value,
        string memory signature,
        bytes memory data,
        uint256 eta
    ) external onlyAdmin {
        // 计算交易哈希
        bytes32 txHash = keccak256(abi.encode(target, value, signature, data, eta));
        
        // 验证交易在队列中
        require(queuedTransactions[txHash], "TimeLock: transaction not queued");
        
        // 从队列中移除交易
        queuedTransactions[txHash] = false;
        
        emit CancelTransaction(txHash);
    }
    
    /**
     * @dev 更新延迟时间（只能通过时间锁本身执行，且需要管理员权限）
     * @param newDelay 新的延迟时间
     * @notice 这个函数只能通过时间锁本身调用，且调用者必须是管理员
     */
    function updateDelay(uint256 newDelay) external onlyThisContract {
        // 验证调用者是否是管理员
        // 注意：当通过时间锁调用时，msg.sender是时间锁合约地址
        // 我们需要验证原始调用者（tx.origin）是否是管理员
        require(tx.origin == admin, "TimeLock: caller is not admin");
        
        require(newDelay >= MIN_DELAY, "TimeLock: delay must exceed minimum delay");
        require(newDelay <= MAX_DELAY, "TimeLock: delay must not exceed maximum delay");
        
        delay = newDelay;
    }
    
    /**
     * @dev 更新管理员地址（只能通过时间锁本身执行，且需要管理员权限）
     * @param newAdmin 新的管理员地址
     * @notice 这个函数只能通过时间锁本身调用，且调用者必须是管理员
     */
    function updateAdmin(address newAdmin) external onlyThisContract {
        // 验证调用者是否是管理员
        // 注意：当通过时间锁调用时，msg.sender是时间锁合约地址
        // 我们需要验证原始调用者（tx.origin）是否是管理员
        require(tx.origin == admin, "TimeLock: caller is not admin");
        
        require(newAdmin != address(0), "TimeLock: invalid admin");
        
        admin = newAdmin;
    }
    
    /**
     * @dev 获取合约余额
     */
    function getBalance() external view returns (uint256) {
        return address(this).balance;
    }
}
