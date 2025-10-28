// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import {Test} from "forge-std/Test.sol";
import {IERC20} from "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import {PausableUpgradeable} from "@openzeppelin/contracts-upgradeable/utils/PausableUpgradeable.sol";
import {FukuaStake} from "../contracts/FukuaStake.sol";
import {FukuaStakeToken} from "../contracts/FukuaStakeToken.sol";

contract FukuaStakePauseTest is Test {
    FukuaStake internal stake;
    FukuaStakeToken internal rewardToken;
    address internal admin = address(this);
    address internal user = address(0xBEEF);

    function setUp() public {
        rewardToken = new FukuaStakeToken();
        stake = new FukuaStake();
        uint256 startBlock = block.number;
        stake.initialize(IERC20(address(rewardToken)), startBlock, startBlock + 1000, 1 ether);

        // 初始化 ETH 质押池，最小质押为 0，锁定 10 个区块
        stake.addPool(address(0), 100, 0, 10, false);
        vm.deal(user, 10 ether);
    }

    function testPauseByAdmin() public {
        stake.pause();
        assertTrue(stake.paused(), "contract should be paused");

        stake.unpause();
        assertFalse(stake.paused(), "contract should be unpaused");
    }

    function testPauseFailsForNonAdmin() public {
        vm.prank(user);
        vm.expectRevert(); // AccessControl revert
        stake.pause();
    }

    function testPausedBlocksDeposits() public {
        stake.pause();

        vm.prank(user);
        vm.expectRevert(PausableUpgradeable.EnforcedPause.selector);
        stake.depositEth{value: 0.1 ether}();
    }

    function testUnpauseAllowsDeposits() public {
        stake.pause();
        stake.unpause();

        vm.prank(user);
        stake.depositEth{value: 0.2 ether}();

        assertEq(stake.stakingBalance(0, user), 0.2 ether, "deposit should succeed after unpause");
    }
}
