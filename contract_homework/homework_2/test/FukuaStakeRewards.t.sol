// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import {Test} from "forge-std/Test.sol";
import {IERC20} from "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import {ERC20} from "@openzeppelin/contracts/token/ERC20/ERC20.sol";
import {FukuaStake} from "../contracts/FukuaStake.sol";
import {FukuaStakeToken} from "../contracts/FukuaStakeToken.sol";

contract MockStakeToken is ERC20 {
    constructor() ERC20("Mock Stake Token", "MST") {}

    function mint(address to, uint256 amount) external {
        _mint(to, amount);
    }
}

contract FukuaStakeRewardsTest is Test {
    FukuaStake internal stake;
    FukuaStakeToken internal rewardToken;
    address internal admin = address(this);
    address internal alice = address(0xA11CE);
    address internal bob = address(0xB0B);

    function setUp() public {
        rewardToken = new FukuaStakeToken();
        stake = new FukuaStake();

        uint256 startBlock = block.number;
        stake.initialize(IERC20(address(rewardToken)), startBlock, startBlock + 1000, 1 ether);

        // 初始化 ETH 池
        stake.addPool(address(0), 100, 0, 10, false);

        vm.deal(alice, 100 ether);
        vm.deal(bob, 100 ether);
    }

    function testRewardAccrualSinglePool() public {
        rewardToken.transfer(address(stake), 20 ether);

        vm.prank(alice);
        stake.depositEth{value: 1 ether}();

        vm.roll(block.number + 10);

        uint256 pending = stake.pendingFukua(0, alice);
        assertEq(pending, 10 ether, "pending reward mismatch after 10 blocks");

        vm.prank(alice);
        stake.claim(0);

        assertEq(rewardToken.balanceOf(alice), 10 ether, "claimed amount should equal pending reward");
    }

    function testPoolWeightAdjustmentAffectsRewards() public {
        MockStakeToken stakeToken = new MockStakeToken();
        stake.addPool(address(stakeToken), 100, 0, 10, false);
        stakeToken.mint(bob, 10 ether);

        vm.prank(bob);
        stakeToken.approve(address(stake), type(uint256).max);

        vm.prank(alice);
        stake.depositEth{value: 1 ether}();

        vm.prank(bob);
        stake.deposit(1, 10 ether);

        vm.roll(block.number + 10);

        uint256 pendingAliceBefore = stake.pendingFukua(0, alice);
        uint256 pendingBobBefore = stake.pendingFukua(1, bob);
        assertApproxEqRel(pendingAliceBefore, pendingBobBefore, 1e15, "equal weights should split rewards evenly");

        // 将池 0 权重调高，并更新全部池奖励
        stake.setPoolWeight(0, 200, true);

        vm.roll(block.number + 10);

        uint256 pendingAliceAfter = stake.pendingFukua(0, alice);
        uint256 pendingBobAfter = stake.pendingFukua(1, bob);

        // 计算新增奖励（扣除之前已累积部分）
        uint256 aliceIncrement = pendingAliceAfter - pendingAliceBefore;
        uint256 bobIncrement = pendingBobAfter - pendingBobBefore;

        // 调整后的比率应约为 2:1
        assertApproxEqRel(aliceIncrement, bobIncrement * 2, 1e15, "updated weight should give pool0 twice the rewards");
    }

    function testClaimHandlesInsufficientRewardBalance() public {
        vm.prank(alice);
        stake.depositEth{value: 1 ether}();

        vm.roll(block.number + 10);

        uint256 pendingBefore = stake.pendingFukua(0, alice);
        assertGt(pendingBefore, 0, "expected pending rewards");

        // 合约中没有奖励代币，首次领取应失败但保留待领取金额
        vm.prank(alice);
        stake.claim(0);

        assertEq(rewardToken.balanceOf(alice), 0, "no rewards should be paid when balance is zero");

        uint256 pendingAfterFirstClaim = stake.pendingFukua(0, alice);
        assertEq(pendingAfterFirstClaim, pendingBefore, "pending reward should remain when balance is insufficient");

        // 仅补充一半奖励，再次领取应转出补充部分并保留剩余
        rewardToken.transfer(address(stake), pendingBefore / 2);

        vm.prank(alice);
        stake.claim(0);

        assertEq(rewardToken.balanceOf(alice), pendingBefore / 2, "should receive funded half");

        uint256 pendingAfterSecondClaim = stake.pendingFukua(0, alice);
        assertEq(pendingAfterSecondClaim, pendingBefore - (pendingBefore / 2), "remaining reward should stay pending");
    }

    function testUnstakeAndWithdrawQueue() public {
        vm.prank(alice);
        stake.depositEth{value: 1 ether}();

        vm.prank(alice);
        stake.unstake(0, 0.4 ether);

        (uint256 totalQueued, uint256 unlocked) = stake.withdrawAmount(0, alice);
        assertEq(totalQueued, 0.4 ether, "queued amount mismatch");
        assertEq(unlocked, 0, "nothing should be unlocked yet");

        vm.prank(alice);
        stake.withdraw(0);

        (totalQueued, unlocked) = stake.withdrawAmount(0, alice);
        assertEq(totalQueued, 0.4 ether, "queue should remain unchanged before unlock");
        assertEq(unlocked, 0, "no funds should be unlocked before reaching unlock height");

        uint256 unlockBlock = block.number + 10;
        vm.roll(unlockBlock);

        uint256 balanceBefore = alice.balance;
        vm.prank(alice);
        stake.withdraw(0);
        uint256 balanceAfter = alice.balance;

        assertEq(balanceAfter - balanceBefore, 0.4 ether, "withdraw after unlock should transfer queued amount");

        (totalQueued, unlocked) = stake.withdrawAmount(0, alice);
        assertEq(totalQueued, 0, "queue should be empty after withdraw");
        assertEq(unlocked, 0, "no unlocked funds should remain");
    }
}
