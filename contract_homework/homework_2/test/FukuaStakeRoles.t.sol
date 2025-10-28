// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import {Test} from "forge-std/Test.sol";
import {IERC20} from "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import {IAccessControl} from "@openzeppelin/contracts/access/IAccessControl.sol";
import {AccessControlUpgradeable} from "@openzeppelin/contracts-upgradeable/access/AccessControlUpgradeable.sol";
import {FukuaStake} from "../contracts/FukuaStake.sol";
import {FukuaStakeToken} from "../contracts/FukuaStakeToken.sol";

contract FukuaStakeHarness is FukuaStake {
    function exposedAuthorizeUpgrade(address newImplementation) external {
        _authorizeUpgrade(newImplementation);
    }
}

contract FukuaStakeRolesTest is Test {
    FukuaStakeHarness internal stake;
    FukuaStakeToken internal rewardToken;

    address internal constant ADMIN_WALLET = 0x5D4a329B262AC7a1d9Ae0f4C54171dF61E2c0b69;
    address internal constant OPERATOR_WALLET = 0xF4076C4a38971D71812B298A6aA9213C5425fa51;
    address internal constant UPGRADE_WALLET = 0x3b8802408e3da17Bc66aA6a8dAb537FD49403355;
    address internal constant RANDOM_USER = 0x1C43bBDd79c85f72DeD3bE8Dc8A3Bb3395E2bAA0;

    function setUp() public {
        rewardToken = new FukuaStakeToken();
        stake = new FukuaStakeHarness();

        vm.prank(ADMIN_WALLET);
        stake.initialize(IERC20(address(rewardToken)), block.number, block.number + 1000, 1 ether);
    }

    function testDefaultAdminCanGrantRoles() public {
        bytes32 adminRole = stake.ADMIN_ROLE();

        vm.prank(ADMIN_WALLET);
        stake.grantRole(adminRole, OPERATOR_WALLET);

        vm.prank(OPERATOR_WALLET);
        stake.pause();

        vm.prank(OPERATOR_WALLET);
        stake.unpause();
    }

    function testNonAdminCannotAccessAdminFunctions() public {
        vm.expectRevert(
            abi.encodeWithSelector(
                IAccessControl.AccessControlUnauthorizedAccount.selector, RANDOM_USER, stake.ADMIN_ROLE()
            )
        );
        vm.prank(RANDOM_USER);
        stake.pause();

        vm.expectRevert(
            abi.encodeWithSelector(
                IAccessControl.AccessControlUnauthorizedAccount.selector, RANDOM_USER, stake.ADMIN_ROLE()
            )
        );
        vm.prank(RANDOM_USER);
        stake.setFukuaPerBlock(2 ether);
    }

    function testUpgradeRoleRestriction() public {
        bytes32 upgradeRole = stake.UPGRADE_ROLE();

        vm.prank(ADMIN_WALLET);
        stake.grantRole(upgradeRole, UPGRADE_WALLET);

        vm.prank(UPGRADE_WALLET);
        stake.exposedAuthorizeUpgrade(address(0xbeef));

        vm.expectRevert(
            abi.encodeWithSelector(IAccessControl.AccessControlUnauthorizedAccount.selector, RANDOM_USER, upgradeRole)
        );
        vm.prank(RANDOM_USER);
        stake.exposedAuthorizeUpgrade(address(0xdead));
    }
}
