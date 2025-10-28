// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

// 导入 OpenZeppelin 的标准 ERC20 接口
import {IERC20} from "@openzeppelin/contracts/token/ERC20/IERC20.sol";
// 导入安全的 ERC20 操作库，防止转账失败等问题
import {SafeERC20} from "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";
// 导入地址工具库
import {Address} from "@openzeppelin/contracts/utils/Address.sol";
// 导入数学运算库，提供安全的乘除法
import {Math} from "@openzeppelin/contracts/utils/math/Math.sol";

// 导入可初始化合约基类（用于代理模式）
import {Initializable} from "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
// 导入 UUPS 可升级代理模式
import {UUPSUpgradeable} from "@openzeppelin/contracts-upgradeable/proxy/utils/UUPSUpgradeable.sol";
// 导入访问控制模块（角色权限管理）
import {AccessControlUpgradeable} from "@openzeppelin/contracts-upgradeable/access/AccessControlUpgradeable.sol";
// 导入可暂停功能模块
import {PausableUpgradeable} from "@openzeppelin/contracts-upgradeable/utils/PausableUpgradeable.sol";

// Fukua 质押合约
// 继承多个基类：可初始化、可升级、可暂停、访问控制
contract FukuaStake is Initializable, UUPSUpgradeable, PausableUpgradeable, AccessControlUpgradeable {
    // 使用 SafeERC20 库为 IERC20 类型添加安全操作
    using SafeERC20 for IERC20;
    // 使用 Address 库为 address 类型添加工具函数
    using Address for address;
    // 使用 Math 库为 uint256 添加数学运算
    using Math for uint256;

    // ************************************** 常量定义 **************************************

    // 管理员角色哈希值，用于权限控制
    bytes32 public constant ADMIN_ROLE = keccak256("admin_role");
    // 升级权限角色哈希值，控制合约升级权限
    bytes32 public constant UPGRADE_ROLE = keccak256("upgrade_role");

    // ETH 质押池的 ID 固定为 0
    uint256 public constant ETH_PID = 0;

    // ************************************** 数据结构定义 **************************************
    /*
    核心算法说明：
    在任何时刻，用户应得但未领取的 Fukua 奖励计算公式为：

    待领取奖励 = (user.stAmount * pool.accFukuaPerShare) - user.finishedFukua

    当用户存入或取出质押代币时，会发生以下步骤：
    1. 更新资金池的 accFukuaPerShare（每个质押代币累计的奖励）和 lastRewardBlock（上次奖励区块）
    2. 将用户的待领取奖励发送到其地址
    3. 更新用户的质押金额 stAmount
    4. 更新用户的已结算奖励 finishedFukua
    */

    // 资金池结构体
    struct Pool {
        // 质押代币的合约地址（address(0x0) 表示 ETH）
        address stTokenAddress;
        // 资金池权重，决定该池在总奖励中的分配比例
        uint256 poolWeight;
        // 最后一次分配奖励的区块号
        uint256 lastRewardBlock;
        // 每个质押代币累计获得的 Fukua 奖励（放大 1e18 倍以保持精度）
        uint256 accFukuaPerShare;
        // 当前池中质押代币的总量
        uint256 stTokenAmount;
        // 最小质押金额限制
        uint256 minDepositAmount;
        // 解除质押后需要锁定的区块数
        uint256 unstakeLockedBlocks;
    }

    // 解除质押请求结构体
    struct UnstakeRequest {
        // 请求提取的金额
        uint256 amount;
        // 该金额可以被提取的解锁区块号
        uint256 unlockBlocks;
    }

    // 用户信息结构体
    struct User {
        // 用户质押的代币数量
        uint256 stAmount;
        // 用户已经结算的 Fukua 奖励数量（用于计算增量奖励）
        uint256 finishedFukua;
        // 用户待领取的 Fukua 奖励数量
        uint256 pendingFukua;
        // 用户的提款请求列表（包含金额和解锁区块）
        UnstakeRequest[] requests;
    }

    // ************************************** 状态变量 **************************************

    // 质押挖矿开始的区块号
    uint256 public startBlock;
    // 质押挖矿结束的区块号
    uint256 public endBlock;
    // 每个区块产出的 Fukua 奖励数量
    uint256 public fukuaPerBlock;

    // 是否暂停提款功能
    bool public withdrawPaused;
    // 是否暂停领取奖励功能
    bool public claimPaused;

    // Fukua 奖励代币的合约接口
    IERC20 public fukuaToken;

    // 所有资金池的总权重（所有池 poolWeight 的总和）
    uint256 public totalPoolWeight;
    // 资金池数组，存储所有质押池信息
    Pool[] public pool;

    // 嵌套映射：资金池 ID => 用户地址 => 用户信息
    mapping(uint256 => mapping(address => User)) public user;

    // ************************************** 事件定义 **************************************

    // 设置 Fukua 代币地址事件
    event SetFukuaToken(IERC20 indexed token);

    // 暂停提款事件
    event PauseWithdraw();

    // 恢复提款事件
    event UnpauseWithdraw();

    // 暂停领取奖励事件
    event PauseClaim();

    // 恢复领取奖励事件
    event UnpauseClaim();

    // 设置开始区块事件
    event SetStartBlock(uint256 indexed startBlock);

    // 设置结束区块事件
    event SetEndBlock(uint256 indexed endBlock);

    // 设置每区块奖励数量事件
    event SetFukuaPerBlock(uint256 indexed newFukuaPerBlock);

    // 添加新资金池事件
    event AddPool(
        address indexed stTokenAddress,
        uint256 indexed poolWeight,
        uint256 indexed lastRewardBlock,
        uint256 minDepositAmount,
        uint256 unstakeLockedBlocks
    );

    // 更新资金池信息事件
    event UpdatePoolInfo(uint256 indexed poolId, uint256 indexed minDepositAmount, uint256 indexed unstakeLockedBlocks);

    // 设置资金池权重事件
    event SetPoolWeight(uint256 indexed poolId, uint256 indexed poolWeight, uint256 totalPoolWeight);

    // 更新资金池奖励事件
    event UpdatePool(uint256 indexed poolId, uint256 indexed lastRewardBlock, uint256 totalFukuaReward);

    // 用户存款事件
    event Deposit(address indexed user, uint256 indexed poolId, uint256 amount);

    // 用户请求解除质押事件
    event RequestUnstake(address indexed user, uint256 indexed poolId, uint256 amount);

    // 用户提款事件
    event Withdraw(address indexed user, uint256 indexed poolId, uint256 amount, uint256 indexed blockNumber);

    // 用户领取奖励事件
    event Claim(address indexed user, uint256 indexed poolId, uint256 fukuaReward);

    // ************************************** 修饰符定义 **************************************

    // 检查资金池 ID 是否有效
    modifier checkPid(uint256 _pid) {
        _checkPid(_pid);
        _;
    }

    // 检查领取功能未暂停
    modifier whenNotClaimPaused() {
        _ensureClaimNotPaused();
        _;
    }

    // 检查提款功能未暂停
    modifier whenNotWithdrawPaused() {
        _ensureWithdrawNotPaused();
        _;
    }

    function _checkPid(uint256 _pid) internal view {
        require(_pid < pool.length, "invalid pid");
    }

    function _ensureClaimNotPaused() internal view {
        require(!claimPaused, "claim is paused");
    }

    function _ensureWithdrawNotPaused() internal view {
        require(!withdrawPaused, "withdraw is paused");
    }

    /**
     * @notice 初始化合约，设置 Fukua 代币地址和基本参数
     * @dev 这是代理合约的初始化函数，只能调用一次
     * @param _fukuaToken         Fukua 奖励代币的合约地址
     * @param _startBlock       挖矿开始区块号
     * @param _endBlock         挖矿结束区块号
     * @param _fukuaPerBlock 每区块产出的 Fukua 数量
     */
    function initialize(IERC20 _fukuaToken, uint256 _startBlock, uint256 _endBlock, uint256 _fukuaPerBlock)
        public
        initializer
    {
        // 参数验证：开始区块必须小于等于结束区块，每区块奖励必须大于 0
        require(_startBlock <= _endBlock && _fukuaPerBlock > 0, "invalid parameters");

        // 初始化访问控制模块
        __AccessControl_init();
        // 初始化 UUPS 可升级模块
        __UUPSUpgradeable_init();
        // 初始化可暂停模块
        __Pausable_init();
        // 授予部署者默认管理员角色
        _grantRole(DEFAULT_ADMIN_ROLE, msg.sender);
        // 授予部署者升级权限角色
        _grantRole(UPGRADE_ROLE, msg.sender);
        // 授予部署者管理员角色
        _grantRole(ADMIN_ROLE, msg.sender);

        // 设置 Fukua 代币地址
        setFukuaToken(_fukuaToken);

        // 设置挖矿开始区块
        startBlock = _startBlock;
        // 设置挖矿结束区块
        endBlock = _endBlock;
        // 设置每区块奖励数量
        fukuaPerBlock = _fukuaPerBlock;
    }

    /**
     * @dev 授权升级函数，只有拥有 UPGRADE_ROLE 的地址才能升级合约
     * @param newImplementation 新实现合约的地址
     */
    function _authorizeUpgrade(address newImplementation)
        internal
        override
        onlyRole(UPGRADE_ROLE) // 要求调用者拥有升级角色

    {
        // 空函数体，权限检查在修饰符中完成
    }

    // ************************************** 管理员函数 **************************************

    /**
     * @notice 全局暂停所有受 `whenNotPaused` 保护的入口
     * @dev 只有管理员可以调用
     */
    function pause() public onlyRole(ADMIN_ROLE) {
        _pause();
    }

    /**
     * @notice 解除全局暂停
     * @dev 只有管理员可以调用
     */
    function unpause() public onlyRole(ADMIN_ROLE) {
        _unpause();
    }

    /**
     * @notice 设置 Fukua 奖励代币地址
     * @dev 只有管理员可以调用
     * @param _fukuaToken Fukua 代币合约地址
     */
    function setFukuaToken(IERC20 _fukuaToken) public onlyRole(ADMIN_ROLE) {
        // 更新 Fukua 代币接口
        fukuaToken = _fukuaToken;

        // 触发设置事件
        emit SetFukuaToken(fukuaToken);
    }

    /**
     * @notice 暂停提款功能
     * @dev 只有管理员可以调用，用于紧急情况
     */
    function pauseWithdraw() public onlyRole(ADMIN_ROLE) {
        // 检查提款功能是否已经暂停
        require(!withdrawPaused, "withdraw has been already paused");

        // 设置提款暂停标志为 true
        withdrawPaused = true;

        // 触发暂停提款事件
        emit PauseWithdraw();
    }

    /**
     * @notice 恢复提款功能
     * @dev 只有管理员可以调用
     */
    function unpauseWithdraw() public onlyRole(ADMIN_ROLE) {
        // 检查提款功能是否已经恢复
        require(withdrawPaused, "withdraw has been already unpaused");

        // 设置提款暂停标志为 false
        withdrawPaused = false;

        // 触发恢复提款事件
        emit UnpauseWithdraw();
    }

    /**
     * @notice 暂停领取奖励功能
     * @dev 只有管理员可以调用，用于紧急情况
     */
    function pauseClaim() public onlyRole(ADMIN_ROLE) {
        // 检查领取功能是否已经暂停
        require(!claimPaused, "claim has been already paused");

        // 设置领取暂停标志为 true
        claimPaused = true;

        // 触发暂停领取事件
        emit PauseClaim();
    }

    /**
     * @notice 恢复领取奖励功能
     * @dev 只有管理员可以调用
     */
    function unpauseClaim() public onlyRole(ADMIN_ROLE) {
        // 检查领取功能是否已经恢复
        require(claimPaused, "claim has been already unpaused");

        // 设置领取暂停标志为 false
        claimPaused = false;

        // 触发恢复领取事件
        emit UnpauseClaim();
    }

    /**
     * @notice 更新挖矿开始区块号
     * @dev 只有管理员可以调用
     * @param _startBlock 新的开始区块号
     */
    function setStartBlock(uint256 _startBlock) public onlyRole(ADMIN_ROLE) {
        // 验证开始区块必须小于等于结束区块
        require(_startBlock <= endBlock, "start block must be smaller than end block");

        // 更新开始区块号
        startBlock = _startBlock;

        // 触发设置开始区块事件
        emit SetStartBlock(_startBlock);
    }

    /**
     * @notice 更新挖矿结束区块号
     * @dev 只有管理员可以调用
     * @param _endBlock 新的结束区块号
     */
    function setEndBlock(uint256 _endBlock) public onlyRole(ADMIN_ROLE) {
        // 验证开始区块必须小于等于结束区块
        require(startBlock <= _endBlock, "start block must be smaller than end block");

        // 更新结束区块号
        endBlock = _endBlock;

        // 触发设置结束区块事件
        emit SetEndBlock(_endBlock);
    }

    /**
     * @notice 更新每区块 Fukua 奖励数量
     * @dev 只有管理员可以调用
     * @param _fukuaPerBlock 新的每区块奖励数量
     */
    function setFukuaPerBlock(uint256 _fukuaPerBlock) public onlyRole(ADMIN_ROLE) {
        // 验证奖励数量必须大于 0
        require(_fukuaPerBlock > 0, "invalid parameter");

        // 更新每区块奖励数量
        fukuaPerBlock = _fukuaPerBlock;

        // 触发设置每区块奖励事件
        emit SetFukuaPerBlock(_fukuaPerBlock);
    }

    /**
     * @notice 添加新的质押资金池
     * @dev 只有管理员可以调用
     * @dev 警告：不要添加相同的质押代币两次，否则会导致奖励计算错误
     * @param _stTokenAddress      质押代币地址（address(0x0) 表示 ETH）
     * @param _poolWeight          资金池权重
     * @param _minDepositAmount    最小质押金额
     * @param _unstakeLockedBlocks 解除质押锁定区块数
     * @param _withUpdate          是否在添加前更新所有池的奖励
     */
    function addPool(
        address _stTokenAddress,
        uint256 _poolWeight,
        uint256 _minDepositAmount,
        uint256 _unstakeLockedBlocks,
        bool _withUpdate
    ) public onlyRole(ADMIN_ROLE) {
        // 第一个池必须是 ETH 池（地址为 0x0）
        if (pool.length > 0) {
            // 如果不是第一个池，质押代币地址不能是 0x0
            require(_stTokenAddress != address(0x0), "invalid staking token address");
        } else {
            // 第一个池必须是 ETH 池，地址必须是 0x0
            require(_stTokenAddress == address(0x0), "invalid staking token address");
        }
        // 允许最小存款金额为 0
        //require(_minDepositAmount > 0, "invalid min deposit amount");
        // 验证解锁区块数必须大于 0
        require(_unstakeLockedBlocks > 0, "invalid withdraw locked blocks");
        // 验证当前区块号必须小于结束区块号
        require(block.number < endBlock, "Already ended");

        // 如果需要，先更新所有资金池的奖励
        if (_withUpdate) {
            massUpdatePools();
        }

        // 计算初始奖励区块号：如果当前区块已经超过开始区块，使用当前区块；否则使用开始区块
        uint256 lastRewardBlock = block.number > startBlock ? block.number : startBlock;
        // 累加总权重
        totalPoolWeight = totalPoolWeight + _poolWeight;

        // 将新池添加到数组
        pool.push(
            Pool({
                stTokenAddress: _stTokenAddress, // 质押代币地址
                poolWeight: _poolWeight, // 池权重
                lastRewardBlock: lastRewardBlock, // 最后奖励区块
                accFukuaPerShare: 0, // 每质押代币累计奖励（初始为 0）
                stTokenAmount: 0, // 质押代币总量（初始为 0）
                minDepositAmount: _minDepositAmount, // 最小存款金额
                unstakeLockedBlocks: _unstakeLockedBlocks // 解锁区块数
            })
        );

        // 触发添加资金池事件
        emit AddPool(_stTokenAddress, _poolWeight, lastRewardBlock, _minDepositAmount, _unstakeLockedBlocks);
    }

    /**
     * @notice 更新指定资金池的参数（最小存款金额和解锁区块数）
     * @dev 只有管理员可以调用
     * @param _pid                 资金池 ID
     * @param _minDepositAmount    新的最小存款金额
     * @param _unstakeLockedBlocks 新的解锁区块数
     */
    function updatePool(uint256 _pid, uint256 _minDepositAmount, uint256 _unstakeLockedBlocks)
        public
        onlyRole(ADMIN_ROLE)
        checkPid(_pid)
    {
        // 更新最小存款金额
        pool[_pid].minDepositAmount = _minDepositAmount;
        // 更新解锁区块数
        pool[_pid].unstakeLockedBlocks = _unstakeLockedBlocks;

        // 触发更新资金池信息事件
        emit UpdatePoolInfo(_pid, _minDepositAmount, _unstakeLockedBlocks);
    }

    /**
     * @notice 更新指定资金池的权重
     * @dev 只有管理员可以调用
     * @param _pid        资金池 ID
     * @param _poolWeight 新的池权重
     * @param _withUpdate 是否在更新前先更新所有池的奖励
     */
    function setPoolWeight(uint256 _pid, uint256 _poolWeight, bool _withUpdate)
        public
        onlyRole(ADMIN_ROLE)
        checkPid(_pid)
    {
        // 验证权重必须大于 0
        require(_poolWeight > 0, "invalid pool weight");

        // 如果需要，先更新所有资金池的奖励
        if (_withUpdate) {
            massUpdatePools();
        }

        // 更新总权重：减去旧权重，加上新权重
        totalPoolWeight = totalPoolWeight - pool[_pid].poolWeight + _poolWeight;
        // 更新池权重
        pool[_pid].poolWeight = _poolWeight;

        // 触发设置池权重事件
        emit SetPoolWeight(_pid, _poolWeight, totalPoolWeight);
    }

    // ************************************** 查询函数 **************************************

    /**
     * @notice 获取资金池的数量
     * @return 资金池数组的长度
     */
    function poolLength() external view returns (uint256) {
        return pool.length;
    }

    /**
     * @notice 计算从 _from 到 _to 区块的奖励乘数（区间为 [_from, _to)）
     * @dev 这个函数计算在有效挖矿期间内的总奖励
     * @param _from 起始区块号（包含）
     * @param _to   结束区块号（不包含）
     * @return multiplier 奖励乘数（区块数 * 每区块奖励）
     */
    function getMultiplier(uint256 _from, uint256 _to) public view returns (uint256 multiplier) {
        // 验证起始区块必须小于等于结束区块
        require(_from <= _to, "invalid block");
        // 如果起始区块小于挖矿开始区块，调整为开始区块
        if (_from < startBlock) _from = startBlock;
        // 如果结束区块大于挖矿结束区块，调整为结束区块
        if (_to > endBlock) _to = endBlock;
        // 再次验证调整后的区块范围
        require(_from <= _to, "end block must be greater than start block");
        // 用于检查乘法是否溢出
        bool success;
        // 计算区块差乘以每区块奖励，使用安全乘法防止溢出
        (success, multiplier) = (_to - _from).tryMul(fukuaPerBlock);
        // 要求乘法成功（未溢出）
        require(success, "multiplier overflow");
    }

    /**
     * @notice 获取用户在指定资金池的待领取 Fukua 奖励数量
     * @param _pid  资金池 ID
     * @param _user 用户地址
     * @return 待领取的 Fukua 数量
     */
    function pendingFukua(uint256 _pid, address _user) external view checkPid(_pid) returns (uint256) {
        // 调用按区块号查询的函数，使用当前区块号
        return pendingFukuaByBlockNumber(_pid, _user, block.number);
    }

    /**
     * @notice 根据指定区块号获取用户的待领取 Fukua 奖励
     * @dev 这个函数可以模拟未来某个区块时的奖励情况
     * @param _pid         资金池 ID
     * @param _user        用户地址
     * @param _blockNumber 查询的区块号
     * @return 待领取的 Fukua 数量
     */
    function pendingFukuaByBlockNumber(uint256 _pid, address _user, uint256 _blockNumber)
        public
        view
        checkPid(_pid)
        returns (uint256)
    {
        // 使用 storage 指针引用池数据（节省 gas）
        Pool storage pool_ = pool[_pid];
        // 使用 storage 指针引用用户数据
        User storage user_ = user[_pid][_user];
        // 获取当前每质押代币的累计奖励
        uint256 accFukuaPerShare = pool_.accFukuaPerShare;
        // 获取池中质押代币总量
        uint256 stSupply = pool_.stTokenAmount;

        // 如果查询区块号大于最后奖励区块，且池中有质押代币，需要计算新增奖励
        if (_blockNumber > pool_.lastRewardBlock && stSupply != 0) {
            // 计算从最后奖励区块到查询区块的总奖励
            uint256 multiplier = getMultiplier(pool_.lastRewardBlock, _blockNumber);
            // 计算该池应得的奖励：总奖励 * 池权重 / 总权重
            uint256 fukuaForPool = multiplier * pool_.poolWeight / totalPoolWeight;
            // 更新每质押代币的累计奖励：原累计值 + (池奖励 * 1e18 / 质押总量)
            // 乘以 1 ether (1e18) 是为了保持小数精度
            accFukuaPerShare = accFukuaPerShare + fukuaForPool * (1 ether) / stSupply;
        }

        // 返回用户待领取奖励：
        // (用户质押量 * 每质押代币累计奖励 / 1e18) - 已结算奖励 + 待领取奖励
        return user_.stAmount * accFukuaPerShare / (1 ether) - user_.finishedFukua + user_.pendingFukua;
    }

    /**
     * @notice 获取用户在指定资金池的质押余额
     * @param _pid  资金池 ID
     * @param _user 用户地址
     * @return 用户的质押代币数量
     */
    function stakingBalance(uint256 _pid, address _user) external view checkPid(_pid) returns (uint256) {
        return user[_pid][_user].stAmount;
    }

    /**
     * @notice 获取用户的提款信息
     * @dev 包括总请求金额和已解锁可提取金额
     * @param _pid  资金池 ID
     * @param _user 用户地址
     * @return requestAmount          总请求提款金额（包括已解锁和未解锁）
     * @return pendingWithdrawAmount  已解锁可立即提取的金额
     */
    function withdrawAmount(uint256 _pid, address _user)
        public
        view
        checkPid(_pid)
        returns (uint256 requestAmount, uint256 pendingWithdrawAmount)
    {
        // 获取用户信息的引用
        User storage user_ = user[_pid][_user];

        // 遍历用户的所有提款请求
        for (uint256 i = 0; i < user_.requests.length; i++) {
            // 如果该请求的解锁区块已到达，累加到可提取金额
            if (user_.requests[i].unlockBlocks <= block.number) {
                pendingWithdrawAmount = pendingWithdrawAmount + user_.requests[i].amount;
            }
            // 累加到总请求金额
            requestAmount = requestAmount + user_.requests[i].amount;
        }
    }

    // ************************************** 公共函数 **************************************

    /**
     * @notice 更新指定资金池的奖励变量到最新状态
     * @dev 这个函数会计算自上次更新以来的奖励，并更新累计奖励值
     * @param _pid 资金池 ID
     */
    function updatePool(uint256 _pid) public checkPid(_pid) {
        // 获取池的引用
        Pool storage pool_ = pool[_pid];

        // 如果当前区块号小于等于最后奖励区块，无需更新
        if (block.number <= pool_.lastRewardBlock) {
            return;
        }

        // 计算从最后奖励区块到当前区块的总奖励，并乘以池权重
        (bool success1, uint256 totalFukuaReward) =
            getMultiplier(pool_.lastRewardBlock, block.number).tryMul(pool_.poolWeight);
        // 检查乘法是否溢出
        require(success1, "overflow");

        // 除以总权重，得到该池应得的奖励
        (success1, totalFukuaReward) = totalFukuaReward.tryDiv(totalPoolWeight);
        // 检查除法是否成功
        require(success1, "overflow");

        // 获取池中的质押代币总量
        uint256 stSupply = pool_.stTokenAmount;
        // 如果池中有质押代币，更新每质押代币的累计奖励
        if (stSupply > 0) {
            // 将奖励乘以 1e18 保持精度
            (bool success2, uint256 totalFukuaRewardScaled) = totalFukuaReward.tryMul(1 ether);
            require(success2, "overflow");

            // 除以质押总量，得到每个质押代币应得的奖励增量
            (success2, totalFukuaRewardScaled) = totalFukuaRewardScaled.tryDiv(stSupply);
            require(success2, "overflow");

            // 累加到原有的每质押代币累计奖励
            (bool success3, uint256 accFukuaPerShare) = pool_.accFukuaPerShare.tryAdd(totalFukuaRewardScaled);
            require(success3, "overflow");
            // 更新池的累计奖励值
            pool_.accFukuaPerShare = accFukuaPerShare;
        }

        // 更新最后奖励区块号为当前区块
        pool_.lastRewardBlock = block.number;

        // 触发更新资金池事件
        emit UpdatePool(_pid, pool_.lastRewardBlock, totalFukuaReward);
    }

    /**
     * @notice 批量更新所有资金池的奖励变量
     * @dev 注意 gas 消耗！池越多，gas 消耗越大
     */
    function massUpdatePools() public {
        // 获取资金池总数
        uint256 length = pool.length;
        // 遍历所有资金池
        for (uint256 pid = 0; pid < length; pid++) {
            // 更新每个池的奖励
            updatePool(pid);
        }
    }

    /**
     * @notice 存入 ETH 进行质押以获取 Fukua 奖励
     * @dev 这是一个 payable 函数，需要发送 ETH
     */
    function depositEth() public payable whenNotPaused {
        // 获取 ETH 池的引用
        Pool storage pool_ = pool[ETH_PID];
        // 验证该池确实是 ETH 池（地址为 0x0）
        require(pool_.stTokenAddress == address(0x0), "invalid staking token address");

        // 获取发送的 ETH 数量
        uint256 _amount = msg.value;
        // 验证存款金额不小于最小存款限制
        require(_amount >= pool_.minDepositAmount, "deposit amount is too small");

        // 调用内部存款函数
        _deposit(ETH_PID, _amount);
    }

    /**
     * @notice 存入 ERC20 代币进行质押以获取 Fukua 奖励
     * @dev 存款前用户需要先授权本合约使用其代币
     * @param _pid    资金池 ID
     * @param _amount 存入的代币数量
     */
    function deposit(uint256 _pid, uint256 _amount) public whenNotPaused checkPid(_pid) {
        // 不能使用此函数存入 ETH（ETH 池 ID 为 0）
        require(_pid != 0, "deposit not support ETH staking");
        // 获取池的引用
        Pool storage pool_ = pool[_pid];
        // 验证存款金额大于最小存款限制
        require(_amount > pool_.minDepositAmount, "deposit amount is too small");

        // 如果存款金额大于 0，从用户账户转入代币到合约
        if (_amount > 0) {
            IERC20(pool_.stTokenAddress).safeTransferFrom(msg.sender, address(this), _amount);
        }

        // 调用内部存款函数
        _deposit(_pid, _amount);
    }

    /**
     * @notice 解除质押（申请提取质押代币）
     * @dev 提取的代币会进入锁定期，需要等待解锁后才能提款
     * @param _pid    资金池 ID
     * @param _amount 解除质押的代币数量
     */
    function unstake(uint256 _pid, uint256 _amount) public whenNotPaused checkPid(_pid) whenNotWithdrawPaused {
        // 获取池的引用
        Pool storage pool_ = pool[_pid];
        // 获取用户的引用
        User storage user_ = user[_pid][msg.sender];

        // 验证用户的质押余额足够
        require(user_.stAmount >= _amount, "Not enough staking token balance");

        // 更新池的奖励状态
        updatePool(_pid);

        // 计算用户当前应得的奖励增量
        uint256 pendingFukua_ = user_.stAmount * pool_.accFukuaPerShare / (1 ether) - user_.finishedFukua;

        // 如果有新增奖励，累加到用户的待领取奖励
        if (pendingFukua_ > 0) {
            user_.pendingFukua = user_.pendingFukua + pendingFukua_;
        }

        // 如果解除质押金额大于 0
        if (_amount > 0) {
            // 减少用户的质押金额
            user_.stAmount = user_.stAmount - _amount;
            // 创建一个新的提款请求，添加到用户的请求列表
            user_.requests
                .push(
                    UnstakeRequest({
                        amount: _amount, // 提款金额
                        unlockBlocks: block.number + pool_.unstakeLockedBlocks // 解锁区块号 = 当前区块 + 锁定区块数
                    })
                );
        }

        // 减少池中的质押代币总量
        pool_.stTokenAmount = pool_.stTokenAmount - _amount;
        // 更新用户的已结算奖励（基于新的质押金额）
        user_.finishedFukua = user_.stAmount * pool_.accFukuaPerShare / (1 ether);

        // 触发解除质押请求事件
        emit RequestUnstake(msg.sender, _pid, _amount);
    }

    /**
     * @notice 提取已解锁的质押代币
     * @dev 只能提取解锁区块已到达的提款请求
     * @param _pid 资金池 ID
     */
    function withdraw(uint256 _pid) public whenNotPaused checkPid(_pid) whenNotWithdrawPaused {
        // 获取池的引用
        Pool storage pool_ = pool[_pid];
        // 获取用户的引用
        User storage user_ = user[_pid][msg.sender];

        // 可提取的金额
        uint256 pendingWithdraw_;
        // 需要从请求列表中移除的请求数量
        uint256 popNum_;
        // 遍历用户的所有提款请求
        for (uint256 i = 0; i < user_.requests.length; i++) {
            // 如果遇到未解锁的请求，停止遍历（因为后面的请求肯定也未解锁）
            if (user_.requests[i].unlockBlocks > block.number) {
                break;
            }
            // 累加可提取金额
            pendingWithdraw_ = pendingWithdraw_ + user_.requests[i].amount;
            // 增加待移除请求计数
            popNum_++;
        }

        // 将未解锁的请求移到数组前面（覆盖已解锁的请求）
        for (uint256 i = 0; i < user_.requests.length - popNum_; i++) {
            user_.requests[i] = user_.requests[i + popNum_];
        }

        // 从数组末尾移除已处理的请求
        for (uint256 i = 0; i < popNum_; i++) {
            user_.requests.pop();
        }

        // 如果有可提取金额，执行转账
        if (pendingWithdraw_ > 0) {
            // 如果是 ETH 池，转账 ETH
            if (pool_.stTokenAddress == address(0x0)) {
                _safeEthTransfer(msg.sender, pendingWithdraw_);
            } else {
                // 否则转账 ERC20 代币
                IERC20(pool_.stTokenAddress).safeTransfer(msg.sender, pendingWithdraw_);
            }
        }

        // 触发提款事件
        emit Withdraw(msg.sender, _pid, pendingWithdraw_, block.number);
    }

    /**
     * @notice 领取 Fukua 代币奖励
     * @param _pid 资金池 ID
     */
    function claim(uint256 _pid) public whenNotPaused checkPid(_pid) whenNotClaimPaused {
        // 获取池的引用
        Pool storage pool_ = pool[_pid];
        // 获取用户的引用
        User storage user_ = user[_pid][msg.sender];

        // 更新池的奖励状态
        updatePool(_pid);

        // 计算用户总的待领取奖励：当前应得奖励 + 之前累积的待领取奖励
        uint256 pendingFukua_ =
            (user_.stAmount * pool_.accFukuaPerShare) / (1 ether) - user_.finishedFukua + user_.pendingFukua;

        uint256 transferred;

        // 如果有待领取奖励
        if (pendingFukua_ > 0) {
            // 清零用户的待领取奖励（因为即将转账，若余额不足再补回）
            user_.pendingFukua = 0;
            // 安全转账 Fukua 代币给用户
            transferred = _safeFukuaTransfer(msg.sender, pendingFukua_);

            // 如果合约余额不足，剩余部分继续保留为待领取奖励
            if (transferred < pendingFukua_) {
                user_.pendingFukua = pendingFukua_ - transferred;
            }
        }

        // 更新用户的已结算奖励（同步到最新状态）
        user_.finishedFukua = (user_.stAmount * pool_.accFukuaPerShare) / (1 ether);

        // 触发领取奖励事件
        emit Claim(msg.sender, _pid, transferred);
    }

    // ************************************** 内部函数 **************************************

    /**
     * @notice 内部存款函数，处理用户存款逻辑
     * @param _pid    资金池 ID
     * @param _amount 存入的代币数量
     */
    function _deposit(uint256 _pid, uint256 _amount) internal {
        // 获取池的引用
        Pool storage pool_ = pool[_pid];
        // 获取用户的引用
        User storage user_ = user[_pid][msg.sender];

        // 更新池的奖励状态到最新
        updatePool(_pid);

        // 如果用户已有质押，先结算之前的奖励
        if (user_.stAmount > 0) {
            // 计算用户当前的累计奖励：质押量 * 每质押代币累计奖励
            (bool success1, uint256 accStake) = user_.stAmount.tryMul(pool_.accFukuaPerShare);
            require(success1, "user stAmount mul accFukuaPerShare overflow");
            // 除以 1e18 恢复精度
            (success1, accStake) = accStake.tryDiv(1 ether);
            require(success1, "accStake div 1 ether overflow");

            // 计算新增奖励：累计奖励 - 已结算奖励
            (bool success2, uint256 pendingFukua_) = accStake.trySub(user_.finishedFukua);
            require(success2, "accStake sub finishedFukua overflow");

            // 如果有新增奖励，累加到用户的待领取奖励
            if (pendingFukua_ > 0) {
                (bool success3, uint256 _pendingFukua) = user_.pendingFukua.tryAdd(pendingFukua_);
                require(success3, "user pendingFukua overflow");
                user_.pendingFukua = _pendingFukua;
            }
        }

        // 如果存款金额大于 0，增加用户的质押金额
        if (_amount > 0) {
            (bool success4, uint256 stAmount) = user_.stAmount.tryAdd(_amount);
            require(success4, "user stAmount overflow");
            user_.stAmount = stAmount;
        }

        // 增加池中的质押代币总量
        (bool success5, uint256 stTokenAmount) = pool_.stTokenAmount.tryAdd(_amount);
        require(success5, "pool stTokenAmount overflow");
        pool_.stTokenAmount = stTokenAmount;

        // 更新用户的已结算奖励：新质押量 * 每质押代币累计奖励
        (bool success6, uint256 finishedFukua) = user_.stAmount.tryMul(pool_.accFukuaPerShare);
        require(success6, "user stAmount mul accFukuaPerShare overflow");

        // 除以 1e18 恢复精度
        (success6, finishedFukua) = finishedFukua.tryDiv(1 ether);
        require(success6, "finishedFukua div 1 ether overflow");

        // 更新已结算奖励
        user_.finishedFukua = finishedFukua;

        // 触发存款事件
        emit Deposit(msg.sender, _pid, _amount);
    }

    /**
     * @notice 安全的 Fukua 代币转账函数
     * @dev 防止因舍入误差导致合约 Fukua 余额不足的情况
     * @param _to     接收地址
     * @param _amount 转账金额
     */
    function _safeFukuaTransfer(address _to, uint256 _amount) internal returns (uint256) {
        uint256 balance = fukuaToken.balanceOf(address(this));
        uint256 transferAmount = _amount > balance ? balance : _amount;

        if (transferAmount > 0) {
            fukuaToken.safeTransfer(_to, transferAmount);
        }

        return transferAmount;
    }

    /**
     * @notice 安全的 ETH 转账函数
     * @dev 使用低级 call 函数转账，并检查返回值
     * @param _to     接收地址
     * @param _amount 转账金额（单位：wei）
     */
    function _safeEthTransfer(address _to, uint256 _amount) internal {
        // 使用 call 函数转账 ETH，并获取返回值
        (bool success, bytes memory data) = address(_to).call{value: _amount}("");

        // 要求 call 调用成功
        require(success, "ETH transfer call failed");
        // 如果有返回数据，解码并验证
        if (data.length > 0) {
            require(abi.decode(data, (bool)), "ETH transfer operation did not succeed");
        }
    }
}
