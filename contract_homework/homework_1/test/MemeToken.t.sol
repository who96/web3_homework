// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "forge-std/Test.sol";
import "../src/MemeToken.sol";

contract MemeTokenTest is Test {
    MemeToken public token;

    address public owner = address(1);
    address public taxRecipient = address(2);
    address public user1 = address(3);
    address public user2 = address(4);
    address public lpPool = address(5);

    uint256 public constant INITIAL_TAX_RATE = 300; // 3%
    uint256 public constant TOTAL_SUPPLY = 21_000_000 * 10 ** 18;

    event TaxCollected(address indexed from, uint256 amount);
    event TaxRateUpdated(uint256 oldRate, uint256 newRate);
    event TaxRecipientUpdated(address indexed oldRecipient, address indexed newRecipient);
    event MaxTransactionAmountUpdated(uint256 oldAmount, uint256 newAmount);

    function setUp() public {
        vm.prank(owner);
        token = new MemeToken(owner, taxRecipient, INITIAL_TAX_RATE);

        // 给用户转一些代币用于测试
        vm.startPrank(owner);
        token.transfer(user1, 1_000_000 * 10 ** 18);
        token.transfer(user2, 1_000_000 * 10 ** 18);
        vm.stopPrank();
    }

    // ========== 基础测试 ==========

    function test_InitialState() public {
        assertEq(token.name(), "Fukua");
        assertEq(token.symbol(), "FK");
        assertEq(token.decimals(), 18);
        assertEq(token.totalSupply(), TOTAL_SUPPLY);
        assertEq(token.taxRate(), INITIAL_TAX_RATE);
        assertEq(token.taxRecipient(), taxRecipient);
        assertEq(token.maxTransactionAmount(), TOTAL_SUPPLY * 2 / 100);
    }

    function test_OwnerHasInitialSupply() public {
        // owner在setUp中转出了2M，所以剩余19M
        assertEq(token.balanceOf(owner), 19_000_000 * 10 ** 18);
    }

    function test_OwnerIsExemptFromTax() public {
        assertTrue(token.isExemptFromTax(owner));
        assertTrue(token.isExemptFromTax(address(token)));
    }

    function test_OwnerIsExemptFromLimit() public {
        assertTrue(token.isExemptFromLimit(owner));
        assertTrue(token.isExemptFromLimit(address(token)));
    }

    // ========== 税收机制测试 ==========

    function test_TransferWithTax() public {
        uint256 transferAmount = 1000 * 10 ** 18;
        uint256 expectedTax = (transferAmount * INITIAL_TAX_RATE) / token.DENOMINATOR();
        uint256 expectedReceive = transferAmount - expectedTax;

        uint256 user1BalBefore = token.balanceOf(user1);
        uint256 user2BalBefore = token.balanceOf(user2);
        uint256 taxRecipientBalBefore = token.balanceOf(taxRecipient);

        vm.prank(user1);
        vm.expectEmit(true, false, false, true);
        emit TaxCollected(user1, expectedTax);
        token.transfer(user2, transferAmount);

        assertEq(token.balanceOf(user1), user1BalBefore - transferAmount);
        assertEq(token.balanceOf(user2), user2BalBefore + expectedReceive);
        assertEq(token.balanceOf(taxRecipient), taxRecipientBalBefore + expectedTax);
    }

    function test_WhitelistTransferNoTax() public {
        uint256 transferAmount = 1000 * 10 ** 18;

        uint256 ownerBalBefore = token.balanceOf(owner);
        uint256 user1BalBefore = token.balanceOf(user1);
        uint256 taxRecipientBalBefore = token.balanceOf(taxRecipient);

        // owner是白名单，转账不收税
        vm.prank(owner);
        token.transfer(user1, transferAmount);

        assertEq(token.balanceOf(owner), ownerBalBefore - transferAmount);
        assertEq(token.balanceOf(user1), user1BalBefore + transferAmount);
        assertEq(token.balanceOf(taxRecipient), taxRecipientBalBefore); // 税收地址余额不变
    }

    function test_TransferToWhitelistNoTax() public {
        // 将user2加入白名单
        vm.prank(owner);
        token.setTaxExempt(user2, true);

        uint256 transferAmount = 1000 * 10 ** 18;
        uint256 taxRecipientBalBefore = token.balanceOf(taxRecipient);

        vm.prank(user1);
        token.transfer(user2, transferAmount);

        // 转给白名单地址不收税
        assertEq(token.balanceOf(taxRecipient), taxRecipientBalBefore);
    }

    function test_TaxRateZeroNoTax() public {
        // 设置税率为0
        vm.prank(owner);
        token.setTaxRate(0);

        uint256 transferAmount = 1000 * 10 ** 18;
        uint256 user2BalBefore = token.balanceOf(user2);
        uint256 taxRecipientBalBefore = token.balanceOf(taxRecipient);

        vm.prank(user1);
        token.transfer(user2, transferAmount);

        // 税率为0，接收全额
        assertEq(token.balanceOf(user2), user2BalBefore + transferAmount);
        assertEq(token.balanceOf(taxRecipient), taxRecipientBalBefore);
    }

    function test_CalculateTax() public {
        uint256 amount = 10000 * 10 ** 18;
        uint256 expectedTax = (amount * INITIAL_TAX_RATE) / token.DENOMINATOR();
        assertEq(token.calculateTax(amount), expectedTax);
    }

    // ========== 税率管理测试 ==========

    function test_SetTaxRate() public {
        uint256 newRate = 400; // 4%

        vm.prank(owner);
        vm.expectEmit(true, true, true, true);
        emit TaxRateUpdated(INITIAL_TAX_RATE, newRate);
        token.setTaxRate(newRate);

        assertEq(token.taxRate(), newRate);
    }

    function test_RevertSetTaxRateExceedsMax() public {
        uint256 invalidRate = 600; // 6%, 超过5%上限

        vm.prank(owner);
        vm.expectRevert("Tax rate exceeds maximum");
        token.setTaxRate(invalidRate);
    }

    function test_RevertSetTaxRateNonOwner() public {
        vm.prank(user1);
        vm.expectRevert(abi.encodeWithSignature("OwnableUnauthorizedAccount(address)", user1));
        token.setTaxRate(400);
    }

    // ========== 税收地址管理测试 ==========

    function test_SetTaxRecipient() public {
        address newRecipient = address(99);

        vm.prank(owner);
        vm.expectEmit(true, true, true, true);
        emit TaxRecipientUpdated(taxRecipient, newRecipient);
        token.setTaxRecipient(newRecipient);

        assertEq(token.taxRecipient(), newRecipient);
    }

    function test_RevertSetTaxRecipientZeroAddress() public {
        vm.prank(owner);
        vm.expectRevert("Tax recipient cannot be zero address");
        token.setTaxRecipient(address(0));
    }

    function test_RevertSetTaxRecipientNonOwner() public {
        vm.prank(user1);
        vm.expectRevert(abi.encodeWithSignature("OwnableUnauthorizedAccount(address)", user1));
        token.setTaxRecipient(address(99));
    }

    // ========== 白名单管理测试 ==========

    function test_SetTaxExempt() public {
        assertFalse(token.isExemptFromTax(user1));

        vm.prank(owner);
        token.setTaxExempt(user1, true);

        assertTrue(token.isExemptFromTax(user1));

        // 移除白名单
        vm.prank(owner);
        token.setTaxExempt(user1, false);

        assertFalse(token.isExemptFromTax(user1));
    }

    function test_RevertSetTaxExemptNonOwner() public {
        vm.prank(user1);
        vm.expectRevert(abi.encodeWithSignature("OwnableUnauthorizedAccount(address)", user1));
        token.setTaxExempt(user2, true);
    }

    function test_SetLimitExempt() public {
        assertFalse(token.isExemptFromLimit(user1));

        vm.prank(owner);
        token.setLimitExempt(user1, true);

        assertTrue(token.isExemptFromLimit(user1));
    }

    // ========== 交易限制测试 ==========

    function test_TransferWithinLimit() public {
        uint256 maxTx = token.maxTransactionAmount();
        uint256 transferAmount = maxTx;

        vm.prank(user1);
        token.transfer(user2, transferAmount);

        // 转账成功（会收税，所以不检查确切余额）
        assertGt(token.balanceOf(user2), 0);
    }

    function test_RevertTransferExceedsLimit() public {
        uint256 maxTx = token.maxTransactionAmount();
        uint256 transferAmount = maxTx + 1;

        vm.prank(user1);
        vm.expectRevert("Transfer amount exceeds max transaction amount");
        token.transfer(user2, transferAmount);
    }

    function test_WhitelistExemptFromLimit() public {
        // owner已在白名单中，可以转超过限制的额度
        uint256 maxTx = token.maxTransactionAmount();
        uint256 transferAmount = maxTx * 2;

        vm.prank(owner);
        token.transfer(user1, transferAmount);

        // 转账成功
        assertGt(token.balanceOf(user1), 0);
    }

    function test_SetMaxTransactionAmount() public {
        uint256 newAmount = 100_000 * 10 ** 18;
        uint256 oldAmount = token.maxTransactionAmount();

        vm.expectEmit(true, true, true, true);
        emit MaxTransactionAmountUpdated(oldAmount, newAmount);

        vm.prank(owner);
        token.setMaxTransactionAmount(newAmount);

        assertEq(token.maxTransactionAmount(), newAmount);
    }

    function test_RevertSetMaxTransactionAmountNonOwner() public {
        vm.prank(user1);
        vm.expectRevert(abi.encodeWithSignature("OwnableUnauthorizedAccount(address)", user1));
        token.setMaxTransactionAmount(100_000 * 10 ** 18);
    }

    // ========== 边界情况测试 ==========

    function test_ZeroAmountTransfer() public {
        vm.prank(user1);
        token.transfer(user2, 0);

        // 不应该revert
    }

    function test_SelfTransfer() public {
        uint256 balBefore = token.balanceOf(user1);

        vm.prank(user1);
        token.transfer(user1, 1000 * 10 ** 18);

        // 自己转给自己会收税
        uint256 expectedTax = (1000 * 10 ** 18 * INITIAL_TAX_RATE) / token.DENOMINATOR();
        assertEq(token.balanceOf(user1), balBefore - expectedTax);
    }

    function test_TransferAllBalance() public {
        uint256 balance = token.balanceOf(user1);

        // 将user1加入白名单避免税费，并豁免交易限制
        vm.startPrank(owner);
        token.setTaxExempt(user1, true);
        token.setLimitExempt(user1, true);
        vm.stopPrank();

        vm.prank(user1);
        token.transfer(user2, balance);

        assertEq(token.balanceOf(user1), 0);
    }

    function test_SmallAmountTaxRounding() public {
        // 小额转账税费可能舍入为0
        uint256 transferAmount = 10 * 10 ** 18; // 10 FK
        uint256 expectedTax = (transferAmount * INITIAL_TAX_RATE) / token.DENOMINATOR();

        uint256 taxRecipientBalBefore = token.balanceOf(taxRecipient);

        vm.prank(user1);
        token.transfer(user2, transferAmount);

        if (expectedTax > 0) {
            assertEq(token.balanceOf(taxRecipient), taxRecipientBalBefore + expectedTax);
        } else {
            // 税费为0，实际免税
            assertEq(token.balanceOf(taxRecipient), taxRecipientBalBefore);
        }
    }

    // ========== ERC20标准功能测试 ==========

    function test_Approve() public {
        vm.prank(user1);
        token.approve(user2, 1000 * 10 ** 18);

        assertEq(token.allowance(user1, user2), 1000 * 10 ** 18);
    }

    function test_TransferFrom() public {
        uint256 approveAmount = 1000 * 10 ** 18;

        // user1授权给user2
        vm.prank(user1);
        token.approve(user2, approveAmount);

        uint256 user1BalBefore = token.balanceOf(user1);
        uint256 ownerBalBefore = token.balanceOf(owner);

        // user2从user1转到owner (owner是白名单，不收税)
        vm.prank(user2);
        token.transferFrom(user1, owner, approveAmount);

        // owner是白名单，转给owner不收税
        assertEq(token.balanceOf(user1), user1BalBefore - approveAmount);
        assertEq(token.balanceOf(owner), ownerBalBefore + approveAmount); // 不收税，全额到账
        assertEq(token.allowance(user1, user2), 0);
    }

    // ========== Fuzz测试 ==========

    function testFuzz_TransferWithTax(uint256 amount) public {
        // 限制amount在合理范围内
        amount = bound(amount, 1, token.balanceOf(user1));

        // 确保不超过交易限制
        if (amount > token.maxTransactionAmount()) {
            amount = token.maxTransactionAmount();
        }

        uint256 expectedTax = (amount * INITIAL_TAX_RATE) / token.DENOMINATOR();
        uint256 expectedReceive = amount - expectedTax;

        uint256 user2BalBefore = token.balanceOf(user2);
        uint256 taxRecipientBalBefore = token.balanceOf(taxRecipient);

        vm.prank(user1);
        token.transfer(user2, amount);

        assertEq(token.balanceOf(user2), user2BalBefore + expectedReceive);
        assertEq(token.balanceOf(taxRecipient), taxRecipientBalBefore + expectedTax);
    }

    function testFuzz_SetTaxRate(uint256 newRate) public {
        // 限制在有效范围内
        newRate = bound(newRate, 0, token.MAX_TAX_RATE());

        vm.prank(owner);
        token.setTaxRate(newRate);

        assertEq(token.taxRate(), newRate);
    }

    // ========== Gas基准测试 ==========

    function test_GasWhitelistTransfer() public {
        uint256 gasBefore = gasleft();

        vm.prank(owner);
        token.transfer(user1, 1000 * 10 ** 18);

        uint256 gasUsed = gasBefore - gasleft();
        emit log_named_uint("Gas used for whitelist transfer", gasUsed);

        // 白名单转账应该 < 100k gas
        assertLt(gasUsed, 100_000);
    }

    function test_GasTaxedTransfer() public {
        uint256 gasBefore = gasleft();

        vm.prank(user1);
        token.transfer(user2, 1000 * 10 ** 18);

        uint256 gasUsed = gasBefore - gasleft();
        emit log_named_uint("Gas used for taxed transfer", gasUsed);

        // 含税转账应该 < 200k gas
        assertLt(gasUsed, 200_000);
    }
}
