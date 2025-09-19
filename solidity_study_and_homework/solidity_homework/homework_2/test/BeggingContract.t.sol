// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "forge-std/Test.sol";
import "../contracts/getMoney.sol";

contract BeggingContractTest is Test {
    BeggingContract public beggingContract;
    address public owner;
    address public donor1;
    address public donor2;
    address public donor3;
    address public donor4;

    event Donation(address indexed donor, uint256 amount);

    function setUp() public {
        owner = makeAddr("owner");
        donor1 = makeAddr("donor1");
        donor2 = makeAddr("donor2");
        donor3 = makeAddr("donor3");
        donor4 = makeAddr("donor4");

        vm.deal(donor1, 10 ether);
        vm.deal(donor2, 10 ether);
        vm.deal(donor3, 10 ether);
        vm.deal(donor4, 10 ether);

        vm.prank(owner);
        beggingContract = new BeggingContract();
    }

    function test_Constructor() public {
        assertEq(beggingContract.owner(), owner, "Owner should be deployer");
    }

    function test_DonateInWorkHours() public {
        uint256 workHourTimestamp = _getWorkHourTimestamp();
        vm.warp(workHourTimestamp);

        uint256 donationAmount = 1 ether;

        vm.expectEmit(true, false, false, true);
        emit Donation(donor1, donationAmount);

        vm.prank(donor1);
        beggingContract.donate{value: donationAmount}();

        assertEq(beggingContract.getDonation(donor1), donationAmount, "Donation amount should be recorded");
        assertEq(address(beggingContract).balance, donationAmount, "Contract balance should increase");
    }

    function test_DonateOutsideWorkHours() public {
        uint256 offHourTimestamp = _getOffHourTimestamp();
        vm.warp(offHourTimestamp);

        vm.prank(donor1);
        vm.expectRevert("Donation only 9:00-18:00 UTC+8 ");
        beggingContract.donate{value: 1 ether}();
    }

    function test_DonateZeroAmount() public {
        uint256 workHourTimestamp = _getWorkHourTimestamp();
        vm.warp(workHourTimestamp);

        vm.prank(donor1);
        vm.expectRevert("Zero donation");
        beggingContract.donate{value: 0}();
    }

    function test_MultipleDonationsFromSameUser() public {
        uint256 workHourTimestamp = _getWorkHourTimestamp();
        vm.warp(workHourTimestamp);

        vm.prank(donor1);
        beggingContract.donate{value: 1 ether}();

        vm.prank(donor1);
        beggingContract.donate{value: 2 ether}();

        assertEq(beggingContract.getDonation(donor1), 3 ether, "Total donations should accumulate");
    }

    function test_TopDonorsRanking() public {
        uint256 workHourTimestamp = _getWorkHourTimestamp();
        vm.warp(workHourTimestamp);

        vm.prank(donor1);
        beggingContract.donate{value: 3 ether}();

        vm.prank(donor2);
        beggingContract.donate{value: 5 ether}();

        vm.prank(donor3);
        beggingContract.donate{value: 1 ether}();

        (address[3] memory topDonors, uint256[3] memory topAmounts) = beggingContract.getTopDonors();

        assertEq(topDonors[0], donor2, "Top donor should be donor2");
        assertEq(topAmounts[0], 5 ether, "Top amount should be 5 ether");

        assertEq(topDonors[1], donor1, "Second donor should be donor1");
        assertEq(topAmounts[1], 3 ether, "Second amount should be 3 ether");

        assertEq(topDonors[2], donor3, "Third donor should be donor3");
        assertEq(topAmounts[2], 1 ether, "Third amount should be 1 ether");
    }

    function test_TopDonorsWithMoreThanThreeDonors() public {
        uint256 workHourTimestamp = _getWorkHourTimestamp();
        vm.warp(workHourTimestamp);

        vm.prank(donor1);
        beggingContract.donate{value: 1 ether}();

        vm.prank(donor2);
        beggingContract.donate{value: 2 ether}();

        vm.prank(donor3);
        beggingContract.donate{value: 3 ether}();

        vm.prank(donor4);
        beggingContract.donate{value: 4 ether}();

        (address[3] memory topDonors, uint256[3] memory topAmounts) = beggingContract.getTopDonors();

        assertEq(topDonors[0], donor4, "Top donor should be donor4");
        assertEq(topAmounts[0], 4 ether, "Top amount should be 4 ether");

        assertEq(topDonors[1], donor3, "Second donor should be donor3");
        assertEq(topAmounts[1], 3 ether, "Second amount should be 3 ether");

        assertEq(topDonors[2], donor2, "Third donor should be donor2");
        assertEq(topAmounts[2], 2 ether, "Third amount should be 2 ether");
    }

    function test_SameDonorUpdatingRanking() public {
        uint256 workHourTimestamp = _getWorkHourTimestamp();
        vm.warp(workHourTimestamp);

        vm.prank(donor1);
        beggingContract.donate{value: 1 ether}();

        vm.prank(donor2);
        beggingContract.donate{value: 2 ether}();

        vm.prank(donor1);
        beggingContract.donate{value: 3 ether}();

        (address[3] memory topDonors, uint256[3] memory topAmounts) = beggingContract.getTopDonors();

        assertEq(topDonors[0], donor1, "Top donor should be donor1");
        assertEq(topAmounts[0], 4 ether, "Top amount should be 4 ether total");

        assertEq(topDonors[1], donor2, "Second donor should be donor2");
        assertEq(topAmounts[1], 2 ether, "Second amount should be 2 ether");
    }

    function test_WithdrawOnlyOwner() public {
        uint256 workHourTimestamp = _getWorkHourTimestamp();
        vm.warp(workHourTimestamp);

        vm.prank(donor1);
        beggingContract.donate{value: 5 ether}();

        uint256 ownerBalanceBefore = owner.balance;

        vm.prank(owner);
        beggingContract.withdraw();

        assertEq(owner.balance, ownerBalanceBefore + 5 ether, "Owner should receive withdrawn funds");
        assertEq(address(beggingContract).balance, 0, "Contract balance should be zero after withdrawal");
    }

    function test_WithdrawNotOwner() public {
        uint256 workHourTimestamp = _getWorkHourTimestamp();
        vm.warp(workHourTimestamp);

        vm.prank(donor1);
        beggingContract.donate{value: 5 ether}();

        vm.prank(donor1);
        vm.expectRevert("Not owner");
        beggingContract.withdraw();
    }

    function test_GetDonationForNonDonor() public {
        assertEq(beggingContract.getDonation(donor1), 0, "Non-donor should have zero donations");
    }

    function test_DonationEvents() public {
        uint256 workHourTimestamp = _getWorkHourTimestamp();
        vm.warp(workHourTimestamp);

        vm.expectEmit(true, false, false, true);
        emit Donation(donor1, 1 ether);

        vm.prank(donor1);
        beggingContract.donate{value: 1 ether}();

        vm.expectEmit(true, false, false, true);
        emit Donation(donor1, 2 ether);

        vm.prank(donor1);
        beggingContract.donate{value: 2 ether}();
    }

    function test_WorkHoursTimeCalculation() public {
        uint256 dayStart = 86400 * 100;

        // UTC+8时间9点 = UTC时间1点
        uint256 utc9AM = dayStart + 1 * 3600;
        // UTC+8时间17:59 = UTC时间9:59 (工作时间最后一分钟)
        uint256 utc6PM_minus1 = dayStart + 9 * 3600 + 59 * 60;
        // UTC+8时间8点 = UTC时间0点
        uint256 utc8AM = dayStart + 0 * 3600;
        // UTC+8时间18点 = UTC时间10点 (工作时间结束)
        uint256 utc6PM = dayStart + 10 * 3600;

        // 测试工作时间开始（9点）可以捐赠
        vm.warp(utc9AM);
        vm.prank(donor1);
        beggingContract.donate{value: 1 ether}();

        // 测试工作时间结束前（17:59）可以捐赠
        vm.warp(utc6PM_minus1);
        vm.prank(donor2);
        beggingContract.donate{value: 1 ether}();

        // 测试工作时间前（8点）不能捐赠
        vm.warp(utc8AM);
        vm.prank(donor3);
        vm.expectRevert("Donation only 9:00-18:00 UTC+8 ");
        beggingContract.donate{value: 1 ether}();

        // 测试工作时间后（18点）不能捐赠
        vm.warp(utc6PM);
        vm.prank(donor4);
        vm.expectRevert("Donation only 9:00-18:00 UTC+8 ");
        beggingContract.donate{value: 1 ether}();
    }

    function testFuzz_DonateValidAmounts(uint256 amount) public {
        vm.assume(amount > 0 && amount <= 1000 ether);

        uint256 workHourTimestamp = _getWorkHourTimestamp();
        vm.warp(workHourTimestamp);

        vm.deal(donor1, amount);

        vm.prank(donor1);
        beggingContract.donate{value: amount}();

        assertEq(beggingContract.getDonation(donor1), amount, "Fuzz: donation amount should be recorded");
    }

    function _getWorkHourTimestamp() private pure returns (uint256) {
        // UTC时间1点 = UTC+8时间9点
        uint256 utcHour = 1;
        uint256 dayStart = 86400 * 100; // 任意一天的开始
        return dayStart + utcHour * 3600;
    }

    function _getOffHourTimestamp() private pure returns (uint256) {
        // UTC时间12点 = UTC+8时间20点 (工作时间外)
        uint256 utcHour = 12;
        uint256 dayStart = 86400 * 100; // 任意一天的开始
        return dayStart + utcHour * 3600;
    }
}