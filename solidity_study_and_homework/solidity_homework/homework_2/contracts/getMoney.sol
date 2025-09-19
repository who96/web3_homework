// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;
// 创建一个名为 BeggingContract 的合约。
// 合约应包含以下功能：
// 一个 mapping 来记录每个捐赠者的捐赠金额。
// 一个 donate 函数，允许用户向合约发送以太币，并记录捐赠信息。
// 一个 withdraw 函数，允许合约所有者提取所有资金。
// 一个 getDonation 函数，允许查询某个地址的捐赠金额。
// 使用 payable 修饰符和 address.transfer 实现支付和提款。
// 使用 onlyOwner 修饰符限制 withdraw 函数只能由合约所有者调用。
// 捐赠事件：添加 Donation 事件，记录每次捐赠的地址和金额。
// 捐赠排行榜：实现一个功能，显示捐赠金额最多的前 3 个地址。
// 时间限制：添加一个时间限制，只有在特定时间段内才能捐赠(比如只有在每天的9点到18点之间才能进行捐献，其他时间不能进行捐献)

contract BeggingContract {
    // 合约所有者地址
    address public owner;
    // 记录每个地址的捐赠总额
    mapping(address => uint256) public donations;
    // 前三名捐赠者地址
    address[3] public topDonors;
    // 前三名捐赠金额
    uint256[3] public topAmounts;

    // 捐赠事件
    event Donation(address indexed donor, uint256 amount);

    // 限制只有所有者可调用
    modifier onlyOwner() {
        require(msg.sender == owner, "Not owner");
        _;
    }

    // 限制工作时间内捐赠（UTC+8 9:00-18:00）
    modifier duringWorkHours() {
        uint256 hour = (block.timestamp / 3600 + 8) % 24;
        require(hour >= 9 && hour < 18, "Donation only 9:00-18:00 UTC+8 ");
        _;
    }
    // 构造函数，部署者成为所有者
    constructor() {
        owner = msg.sender;
    }

    // 捐赠函数
    function donate() external payable duringWorkHours {
        require(msg.value > 0, "Zero donation");

        donations[msg.sender] += msg.value;
        emit Donation(msg.sender, msg.value);

        _updateTopDonors(msg.sender, donations[msg.sender]);
    }

    // 提取合约所有资金
    function withdraw() external onlyOwner {
        payable(owner).transfer(address(this).balance);
    }

    // 查询指定地址的捐赠总额
    function getDonation(address donor) external view returns (uint256) {
        return donations[donor];
    }

    // 获取捐赠排行榜前三名
    function getTopDonors() external view returns (address[3] memory, uint256[3] memory) {
        return (topDonors, topAmounts);
    }

    // 更新排行榜（内部函数）
    function _updateTopDonors(address donor, uint256 totalAmount) private {
        for (uint i = 0; i < 3; i++) {
            if (totalAmount > topAmounts[i]) {
                // 向后移动排名较低的捐赠者
                for (uint j = 2; j > i; j--) {
                    if (topDonors[j-1] != donor) {
                        topDonors[j] = topDonors[j-1];
                        topAmounts[j] = topAmounts[j-1];
                    }
                }
                // 插入新的排名
                topDonors[i] = donor;
                topAmounts[i] = totalAmount;
                break;
            }
        }
    }
}
