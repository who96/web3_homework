// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import {Script, console2} from "forge-std/Script.sol";
import {FukuaStake} from "../contracts/FukuaStake.sol";

contract ShowStatus is Script {
    function run() external view {
        address stakeAddr = vm.envAddress("FUKUA_STAKE_ADDRESS");
        FukuaStake stake = FukuaStake(stakeAddr);

        console2.log("startBlock", stake.startBlock());
        console2.log("endBlock", stake.endBlock());
        console2.log("fukuaPerBlock", stake.fukuaPerBlock());
        console2.log("totalPoolWeight", stake.totalPoolWeight());
        console2.log("paused", stake.paused());
        console2.log("claimPaused", stake.claimPaused());
        console2.log("withdrawPaused", stake.withdrawPaused());

        uint256 pools = stake.poolLength();
        console2.log("poolLength", pools);
        for (uint256 pid = 0; pid < pools; pid++) {
            (address stToken, uint256 weight, uint256 lastRewardBlock, uint256 accFukuaPerShare, uint256 stTokenAmount,
                uint256 minDepositAmount, uint256 lockedBlocks) = stake.pool(pid);
            console2.log("--- Pool", pid);
            console2.log("staking token", stToken);
            console2.log("weight", weight);
            console2.log("lastRewardBlock", lastRewardBlock);
            console2.log("accFukuaPerShare", accFukuaPerShare);
            console2.log("stTokenAmount", stTokenAmount);
            console2.log("minDeposit", minDepositAmount);
            console2.log("lockedBlocks", lockedBlocks);
        }
    }
}
