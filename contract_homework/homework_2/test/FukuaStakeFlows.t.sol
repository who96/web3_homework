// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import {Test} from "forge-std/Test.sol";
import {IERC20} from "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import {IAccessControl} from "@openzeppelin/contracts/access/IAccessControl.sol";
import {PausableUpgradeable} from "@openzeppelin/contracts-upgradeable/utils/PausableUpgradeable.sol";
import {FukuaStake} from "../contracts/FukuaStake.sol";
import {FukuaStakeToken} from "../contracts/FukuaStakeToken.sol";

contract FukuaStakeFlowHarness is FukuaStake {
    function exposedAuthorizeUpgrade(address newImplementation) external {
        _authorizeUpgrade(newImplementation);
    }
}

contract FukuaStakeFlowsTest is Test {
    FukuaStakeFlowHarness internal stake;
    FukuaStakeToken internal rewardToken;

    address internal constant ADMIN_WALLET = 0x5D4a329B262AC7a1d9Ae0f4C54171dF61E2c0b69;
    address internal constant USER = 0xF4076C4a38971D71812B298A6aA9213C5425fa51;
    address internal constant USER_TWO = 0x3b8802408e3da17Bc66aA6a8dAb537FD49403355;

    function setUp() public {
        rewardToken = new FukuaStakeToken();
        stake = new FukuaStakeFlowHarness();

        vm.prank(ADMIN_WALLET);
        stake.initialize(IERC20(address(rewardToken)), block.number, block.number + 1000, 1 ether);

        // 第一个池为 ETH，另建一个 ERC20 池用于多资产路径
        vm.prank(ADMIN_WALLET);
        stake.addPool(address(0), 100, 0, 10, false);
        vm.prank(ADMIN_WALLET);
        stake.addPool(address(rewardToken), 200, 0, 20, false);

        vm.deal(USER, 50 ether);
        vm.deal(USER_TWO, 50 ether);

        // 为第二个用户准备 ERC20 质押资产
        rewardToken.transfer(USER_TWO, 1_000 ether);

        // 合约预充奖励
        rewardToken.transfer(address(stake), 5_000 ether);
    }

    function testFullHappyPath() public {
        // USER 质押 ETH
        vm.prank(USER);
        stake.depositEth{value: 5 ether}();

        // USER_TWO 质押 ERC20
        vm.prank(USER_TWO);
        rewardToken.approve(address(stake), 500 ether);
        vm.prank(USER_TWO);
        stake.deposit(1, 500 ether);

        // 推进区块，累积奖励
        vm.roll(block.number + 20);

        // USER 领取奖励
        vm.prank(USER);
        stake.claim(0);
        uint256 balanceUser = rewardToken.balanceOf(USER);
        assertGt(balanceUser, 0, "user should receive rewards");

        // USER_TWO 解质押一半并在解锁后提现
        vm.prank(USER_TWO);
        stake.unstake(1, 200 ether);
        vm.roll(block.number + 25);
        uint256 userTwoBalanceBefore = rewardToken.balanceOf(USER_TWO);
        vm.prank(USER_TWO);
        stake.withdraw(1);
        uint256 userTwoBalanceAfter = rewardToken.balanceOf(USER_TWO);
        assertEq(userTwoBalanceAfter - userTwoBalanceBefore, 200 ether, "withdraw should return unstaked amount");

        // 管理员暂停并验证入口被阻断
        vm.prank(ADMIN_WALLET);
        stake.pause();

        vm.prank(USER);
        vm.expectRevert(PausableUpgradeable.EnforcedPause.selector);
        stake.depositEth{value: 1 ether}();

        // 管理员解除暂停
        vm.prank(ADMIN_WALLET);
        stake.unpause();

        // 再次质押确认恢复
        vm.prank(USER);
        stake.depositEth{value: 1 ether}();

        // 校验合约状态
        assertTrue(stake.paused() == false, "contract should be active");
        (uint256 queued, uint256 unlocked) = stake.withdrawAmount(1, USER_TWO);
        assertEq(queued, 0, "queue should be empty after withdraw");
        assertEq(unlocked, 0, "no unlocked funds remaining");
    }

    function testAdminEmergencyWithdrawPausedState() public {
        // 授予 USER 管理员角色后触发暂停
        vm.startPrank(ADMIN_WALLET);
        stake.grantRole(stake.ADMIN_ROLE(), USER);
        vm.stopPrank();

        vm.prank(USER);
        stake.pauseWithdraw();

        vm.prank(USER);
        vm.expectRevert("withdraw is paused");
        stake.withdraw(0);

        vm.prank(USER);
        stake.unpauseWithdraw();

        vm.prank(USER);
        // 无请求不会抛错
        stake.withdraw(0);
    }

    function testUpgradeRoleCheckInFlow() public {
        // 新账户尝试升级必须失败
        vm.prank(USER);
        vm.expectRevert(
            abi.encodeWithSelector(IAccessControl.AccessControlUnauthorizedAccount.selector, USER, stake.UPGRADE_ROLE())
        );
        vm.prank(USER);
        stake.exposedAuthorizeUpgrade(address(0xbeef));

        // 管理员授予后成功
        vm.startPrank(ADMIN_WALLET);
        stake.grantRole(stake.UPGRADE_ROLE(), USER);
        vm.stopPrank();

        vm.prank(USER);
        stake.exposedAuthorizeUpgrade(address(0xdead));
    }
}
