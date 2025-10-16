// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "@openzeppelin/contracts/token/ERC20/ERC20.sol";
import "@openzeppelin/contracts/access/Ownable.sol";

/**
 * @title MemeToken (Fukua)
 * @notice SHIB风格的Meme代币，支持交易税和防操纵机制
 * @dev 继承OpenZeppelin的ERC20和Ownable，在_update()中实现税收逻辑
 *
 * 核心功能:
 * - 交易税机制：每笔转账收取可配置的税费(1-5%)
 * - 白名单系统：owner、流动性池、合约本身免税
 * - 交易限制：单笔最大交易额度，防止大额抛售
 * - 完全兼容ERC20标准，可在Uniswap等DEX交易
 */
contract MemeToken is ERC20, Ownable {
    /// @notice 税率分母，10000 = 100%，支持精确到0.01%的税率
    uint256 public constant DENOMINATOR = 10000;

    /// @notice 最大税率限制为5% (500 basis points)
    uint256 public constant MAX_TAX_RATE = 500;

    /// @notice 当前税率 (basis points，例如 300 = 3%)
    uint256 private _taxRate;

    /// @notice 税费接收地址
    address private _taxRecipient;

    /// @notice 单笔最大交易额度
    uint256 public maxTransactionAmount;

    /// @notice 税收豁免白名单
    mapping(address => bool) private _isExemptFromTax;

    /// @notice 交易限制豁免白名单
    mapping(address => bool) private _isExemptFromLimit;

    // ========== Events ==========

    /// @notice 税费收取事件
    /// @param from 转账发送者
    /// @param amount 税费金额
    event TaxCollected(address indexed from, uint256 amount);

    /// @notice 税率更新事件
    /// @param oldRate 旧税率
    /// @param newRate 新税率
    event TaxRateUpdated(uint256 oldRate, uint256 newRate);

    /// @notice 税收接收地址更新事件
    /// @param oldRecipient 旧地址
    /// @param newRecipient 新地址
    event TaxRecipientUpdated(address indexed oldRecipient, address indexed newRecipient);

    /// @notice 最大交易额度更新事件
    /// @param oldAmount 旧额度
    /// @param newAmount 新额度
    event MaxTransactionAmountUpdated(uint256 oldAmount, uint256 newAmount);

    // ========== Constructor ==========

    /**
     * @notice 构造函数
     * @param initialOwner 初始owner地址
     * @param taxRecipient_ 税收接收地址
     * @param taxRate_ 初始税率 (basis points)
     */
    constructor(address initialOwner, address taxRecipient_, uint256 taxRate_)
        ERC20("Fukua", "FK")
        Ownable(initialOwner)
    {
        require(taxRecipient_ != address(0), "Tax recipient cannot be zero address");
        require(taxRate_ <= MAX_TAX_RATE, "Tax rate exceeds maximum");

        _taxRecipient = taxRecipient_;
        _taxRate = taxRate_;

        // 总供应量: 21,000,000 FK
        uint256 totalSupply = 21_000_000 * 10 ** decimals();

        // 设置单笔最大交易额度为总量的2% (420,000 FK)
        maxTransactionAmount = totalSupply * 2 / 100;

        // 将owner、合约自身加入白名单
        _isExemptFromTax[initialOwner] = true;
        _isExemptFromTax[address(this)] = true;
        _isExemptFromLimit[initialOwner] = true;
        _isExemptFromLimit[address(this)] = true;

        // 铸造全部代币给owner
        _mint(initialOwner, totalSupply);
    }

    // ========== Core Override ==========

    /**
     * @dev 重写_update函数，在所有转账中统一处理税收和限制逻辑
     * @param from 发送者地址 (address(0)表示铸币)
     * @param to 接收者地址 (address(0)表示销毁)
     * @param value 转账金额
     */
    function _update(address from, address to, uint256 value) internal virtual override {
        // 铸币和销毁操作直接执行，不收税也不限制
        if (from == address(0) || to == address(0)) {
            super._update(from, to, value);
            return;
        }

        // 检查交易限制（发送者不在豁免名单中）
        if (!_isExemptFromLimit[from]) {
            require(value <= maxTransactionAmount, "Transfer amount exceeds max transaction amount");
        }

        // 检查是否需要收税（发送者和接收者都不在豁免名单中）
        bool shouldTax = !_isExemptFromTax[from] && !_isExemptFromTax[to] && _taxRate > 0;

        if (shouldTax) {
            // 计算税费
            uint256 taxAmount = (value * _taxRate) / DENOMINATOR;
            uint256 amountAfterTax = value - taxAmount;

            // 转账扣税后的金额给接收者
            super._update(from, to, amountAfterTax);

            // 转税费给税收接收地址
            if (taxAmount > 0) {
                super._update(from, _taxRecipient, taxAmount);
                emit TaxCollected(from, taxAmount);
            }
        } else {
            // 不收税，直接转账
            super._update(from, to, value);
        }
    }

    // ========== Admin Functions ==========

    /**
     * @notice 设置税率
     * @dev 只有owner可以调用，税率不能超过MAX_TAX_RATE
     * @param newRate 新税率 (basis points)
     */
    function setTaxRate(uint256 newRate) external onlyOwner {
        require(newRate <= MAX_TAX_RATE, "Tax rate exceeds maximum");
        uint256 oldRate = _taxRate;
        _taxRate = newRate;
        emit TaxRateUpdated(oldRate, newRate);
    }

    /**
     * @notice 设置税收接收地址
     * @dev 只有owner可以调用
     * @param newRecipient 新的税收接收地址
     */
    function setTaxRecipient(address newRecipient) external onlyOwner {
        require(newRecipient != address(0), "Tax recipient cannot be zero address");
        address oldRecipient = _taxRecipient;
        _taxRecipient = newRecipient;
        emit TaxRecipientUpdated(oldRecipient, newRecipient);
    }

    /**
     * @notice 设置单笔最大交易额度
     * @dev 只有owner可以调用
     * @param newAmount 新的最大交易额度
     */
    function setMaxTransactionAmount(uint256 newAmount) external onlyOwner {
        uint256 oldAmount = maxTransactionAmount;
        maxTransactionAmount = newAmount;
        emit MaxTransactionAmountUpdated(oldAmount, newAmount);
    }

    /**
     * @notice 设置地址的税收豁免状态
     * @dev 只有owner可以调用，用于添加流动性池等地址到白名单
     * @param account 目标地址
     * @param exempt true=豁免税收, false=需要收税
     */
    function setTaxExempt(address account, bool exempt) external onlyOwner {
        _isExemptFromTax[account] = exempt;
    }

    /**
     * @notice 设置地址的交易限制豁免状态
     * @dev 只有owner可以调用
     * @param account 目标地址
     * @param exempt true=豁免限制, false=受限制
     */
    function setLimitExempt(address account, bool exempt) external onlyOwner {
        _isExemptFromLimit[account] = exempt;
    }

    // ========== View Functions ==========

    /**
     * @notice 查询当前税率
     * @return 税率 (basis points)
     */
    function taxRate() external view returns (uint256) {
        return _taxRate;
    }

    /**
     * @notice 查询税收接收地址
     * @return 税收接收地址
     */
    function taxRecipient() external view returns (address) {
        return _taxRecipient;
    }

    /**
     * @notice 查询地址是否税收豁免
     * @param account 目标地址
     * @return true=豁免, false=需要收税
     */
    function isExemptFromTax(address account) external view returns (bool) {
        return _isExemptFromTax[account];
    }

    /**
     * @notice 查询地址是否交易限制豁免
     * @param account 目标地址
     * @return true=豁免, false=受限制
     */
    function isExemptFromLimit(address account) external view returns (bool) {
        return _isExemptFromLimit[account];
    }

    /**
     * @notice 计算给定金额的税费
     * @param amount 转账金额
     * @return 税费金额
     */
    function calculateTax(uint256 amount) external view returns (uint256) {
        return (amount * _taxRate) / DENOMINATOR;
    }
}
