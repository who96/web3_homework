// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import {Script} from "forge-std/Script.sol";
import {IERC20} from "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import {FukuaStake} from "../contracts/FukuaStake.sol";
import {FukuaStakeToken} from "../contracts/FukuaStakeToken.sol";

contract DeployFukuaStake is Script {
    function run() external {
        vm.startBroadcast();

        // 预部署示例：先部署奖励代币再部署质押合约
        FukuaStakeToken rewardToken = new FukuaStakeToken();

        // 默认参数仅作为演示，请在实际部署前根据需求调整
        FukuaStake stake = new FukuaStake();
        stake.initialize(IERC20(address(rewardToken)), block.number, block.number + 100_000, 1 ether);

        vm.stopBroadcast();

        // 防止编译器优化掉变量
        rewardToken;
        stake;
    }
}
