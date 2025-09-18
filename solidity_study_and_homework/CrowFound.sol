// SPDX-License-Identifier: MIT // 许可证标识：采用 MIT 开源许可证
pragma solidity ^0.8.10; // 指定 Solidity 编译器版本为 0.8.10 及以上

interface IERC20 { // ERC20 代币接口，供本合约与代币交互
    function transfer(address, uint) external returns (bool); // 从合约将代币转给某地址

    function transferFrom(
        address, // 代币转出方（所有者）
        address, // 代币接收方
        uint     // 转账数量
    ) external returns (bool); // 从转出方代扣转账（需事先授权）
}

contract CrowdFund { // 众筹合约主体
    event Launch(
        uint id,                   // 活动 ID
        address indexed creator,   // 发起人地址（可索引）
        uint goal,                 // 募集目标金额
        uint32 startAt,            // 开始时间（时间戳）
        uint32 endAt               // 结束时间（时间戳）
    );
    event Cancel(uint id); // 活动取消事件
    event Pledge(uint indexed id, address indexed caller, uint amount); // 质押（出资）事件
    event Unpledge(uint indexed id, address indexed caller, uint amount); // 取消质押事件
    event Claim(uint id); // 创作者提取资金事件
    event Refund(uint id, address indexed caller, uint amount); // 出资人退款事件

    struct Campaign { // 众筹活动数据结构
        // Creator of campaign // 活动发起者地址
        address creator;
        // Amount of tokens to raise // 目标募集代币数量
        uint goal;
        // Total amount pledged // 当前已质押（出资）总额
        uint pledged;
        // Timestamp of start of campaign // 活动开始时间戳
        uint32 startAt;
        // Timestamp of end of campaign // 活动结束时间戳
        uint32 endAt;
        // True if goal was reached and creator has claimed the tokens. // 是否已被发起者提取
        bool claimed;
    }

    IERC20 public immutable token; // 参与众筹使用的 ERC20 代币合约地址（不可变）
    // Total count of campaigns created.
    // It is also used to generate id for new campaigns.
    uint public count; // 已创建的活动总数，同时用于生成下一个活动的 ID
    // Mapping from id to Campaign
    mapping(uint => Campaign) public campaigns; // 活动 ID => 活动详情
    // Mapping from campaign id => pledger => amount pledged
    mapping(uint => mapping(address => uint)) public pledgedAmount; // 活动 ID => 出资人地址 => 出资金额

    constructor(address _token) { // 部署时传入 ERC20 代币地址
        token = IERC20(_token); // 记录代币接口实例
    }

    function launch(
        uint _goal,        // 目标金额
        uint32 _startAt,   // 开始时间
        uint32 _endAt      // 结束时间
    ) external { // 创建新的众筹活动，只能由外部账户/合约调用
        require(_startAt >= block.timestamp, "start at < now"); // 要求开始时间不早于当前时间
        require(_endAt >= _startAt, "end at < start at"); // 要求结束时间不早于开始时间
        require(_endAt <= block.timestamp + 90 days, "end at > max duration"); // 活动最长 90 天

        count += 1; // 递增活动计数，作为新活动 ID
        campaigns[count] = Campaign({
            creator: msg.sender, // 记录发起者
            goal: _goal,         // 记录目标金额
            pledged: 0,          // 初始质押金额为 0
            startAt: _startAt,   // 设置开始时间
            endAt: _endAt,       // 设置结束时间
            claimed: false       // 初始未提取
        });

        emit Launch(count, msg.sender, _goal, _startAt, _endAt); // 触发活动创建事件
    }

    function cancel(uint _id) external { // 取消未开始的活动
        Campaign memory campaign = campaigns[_id]; // 读取活动数据到内存
        require(campaign.creator == msg.sender, "not creator"); // 只有发起者可以取消
        require(block.timestamp < campaign.startAt, "started"); // 活动未开始才能取消

        delete campaigns[_id]; // 删除活动记录
        emit Cancel(_id); // 触发取消事件
    }

    function pledge(uint _id, uint _amount) external { // 对指定活动出资质押
        Campaign storage campaign = campaigns[_id]; // 引用存储中的活动数据
        require(block.timestamp >= campaign.startAt, "not started"); // 活动已开始
        require(block.timestamp <= campaign.endAt, "ended"); // 活动未结束

        campaign.pledged += _amount; // 增加活动总质押额
        pledgedAmount[_id][msg.sender] += _amount; // 记录出资人对该活动的出资额
        token.transferFrom(msg.sender, address(this), _amount); // 从出资人转代币到合约

        emit Pledge(_id, msg.sender, _amount); // 触发质押事件
    }

    function unpledge(uint _id, uint _amount) external { // 在活动结束前撤回部分/全部出资
        Campaign storage campaign = campaigns[_id]; // 引用活动数据
        require(block.timestamp <= campaign.endAt, "ended"); // 必须在活动结束前

        campaign.pledged -= _amount; // 减少活动总质押额
        pledgedAmount[_id][msg.sender] -= _amount; // 减少该出资人的记录
        token.transfer(msg.sender, _amount); // 将相应代币退还给出资人

        emit Unpledge(_id, msg.sender, _amount); // 触发撤回事件
    }

    function claim(uint _id) external { // 活动结束且达标后，发起者提取资金
        Campaign storage campaign = campaigns[_id]; // 引用活动数据
        require(campaign.creator == msg.sender, "not creator"); // 只有发起者可提取
        require(block.timestamp > campaign.endAt, "not ended"); // 活动已结束
        require(campaign.pledged >= campaign.goal, "pledged < goal"); // 募集达到目标
        require(!campaign.claimed, "claimed"); // 尚未提取过

        campaign.claimed = true; // 标记已提取
        token.transfer(campaign.creator, campaign.pledged); // 将全部质押金额转给发起者

        emit Claim(_id); // 触发提取事件
    }

    function refund(uint _id) external { // 活动未达标，出资人可在结束后退款
        Campaign memory campaign = campaigns[_id]; // 读取活动数据到内存（只读）
        require(block.timestamp > campaign.endAt, "not ended"); // 活动已结束
        require(campaign.pledged < campaign.goal, "pledged >= goal"); // 未达到目标

        uint bal = pledgedAmount[_id][msg.sender]; // 读取调用者在该活动中的出资额
        pledgedAmount[_id][msg.sender] = 0; // 清零调用者的出资记录，防止重复退款
        token.transfer(msg.sender, bal); // 退回代币

        emit Refund(_id, msg.sender, bal); // 触发退款事件
    }
}
