// SPDX-License-Identifier: MIT
pragma solidity ^0.8.7;

import "forge-std/Test.sol";
import "forge-std/console.sol";
import "../../src/SimpleAuction.sol";
import "../../src/SimpleNFT.sol";
import "../../src/AuctionPriceFeed.sol";

contract MockPriceFeed {
    int256 public price = 2000e8; // $2000 with 8 decimals

    function latestRoundData() external view returns (
        uint80 roundId,
        int256 answer,
        uint256 startedAt,
        uint256 updatedAt_,
        uint80 answeredInRound
    ) {
        return (1, price, block.timestamp, block.timestamp, 1);
    }

    function decimals() external pure returns (uint8) {
        return 8;
    }

    function setPrice(int256 _price) external {
        price = _price;
    }
}

contract AuctionIntegrationTest is Test {
    SimpleAuction public auction;
    SimpleNFT public nft;
    AuctionPriceFeed public priceFeed;
    MockPriceFeed public mockPriceFeed;

    address public deployer;
    address public alice; // Seller
    address public bob; // Bidder 1
    address public charlie; // Bidder 2
    address public platformWallet;

    uint256 public constant DURATION = 7200; // 2 hours
    uint256 public constant RESERVE_PRICE = 0.5 ether;

    function setUp() public {
        deployer = address(this);
        alice = address(0x100);
        bob = address(0x200);
        charlie = address(0x300);
        platformWallet = address(0x999);

        // Deploy contracts
        mockPriceFeed = new MockPriceFeed();
        nft = new SimpleNFT();
        priceFeed = new AuctionPriceFeed(address(mockPriceFeed));
        auction = new SimpleAuction(address(priceFeed), platformWallet);

        // Setup users with ETH
        vm.deal(alice, 10 ether);
        vm.deal(bob, 10 ether);
        vm.deal(charlie, 10 ether);

        // Alice mints NFTs
        vm.startPrank(alice);
        for (uint256 i = 0; i < 3; i++) {
            nft.mint(alice);
        }
        nft.setApprovalForAll(address(auction), true);
        vm.stopPrank();
    }

    function testCompleteAuctionWorkflowSuccessful() public {
        // === Phase 1: Create Auction ===
        vm.prank(alice);
        uint256 auctionId = auction.createAuction(
            address(nft),
            0, // tokenId
            DURATION,
            RESERVE_PRICE
        );

        // Verify auction creation
        assertEq(nft.ownerOf(0), address(auction));
        assertTrue(auction.isAuctionActive(auctionId));

        SimpleAuction.Auction memory auctionInfo = auction.getAuction(auctionId);
        assertEq(auctionInfo.seller, alice);
        assertEq(auctionInfo.reservePrice, RESERVE_PRICE);
        assertEq(auctionInfo.highestBidder, address(0));

        // === Phase 2: Bidding Phase ===

        // Bob places first bid at reserve price
        vm.prank(bob);
        auction.placeBid{value: RESERVE_PRICE}(auctionId);

        auctionInfo = auction.getAuction(auctionId);
        assertEq(auctionInfo.highestBidder, bob);
        assertEq(auctionInfo.highestBid, RESERVE_PRICE);

        // Charlie outbids Bob
        uint256 charlieBid = RESERVE_PRICE + (RESERVE_PRICE * 500 / 10000); // 5% increment
        vm.prank(charlie);
        auction.placeBid{value: charlieBid}(auctionId);

        auctionInfo = auction.getAuction(auctionId);
        assertEq(auctionInfo.highestBidder, charlie);
        assertEq(auctionInfo.highestBid, charlieBid);

        // Bob has pending refund
        assertEq(auction.pendingRefunds(bob), RESERVE_PRICE);

        // Bob outbids Charlie
        uint256 bobSecondBid = charlieBid + (charlieBid * 500 / 10000);
        vm.prank(bob);
        auction.placeBid{value: bobSecondBid}(auctionId);

        auctionInfo = auction.getAuction(auctionId);
        assertEq(auctionInfo.highestBidder, bob);
        assertEq(auctionInfo.highestBid, bobSecondBid);

        // Charlie withdraws refund
        uint256 charlieBalanceBefore = charlie.balance;
        vm.prank(charlie);
        auction.withdrawRefund();
        assertEq(charlie.balance, charlieBalanceBefore + charlieBid);

        // === Phase 3: Auction End ===

        // Wait for auction to end
        vm.warp(block.timestamp + DURATION + 1);

        // End auction
        auction.endAuction(auctionId);

        auctionInfo = auction.getAuction(auctionId);
        assertTrue(auctionInfo.ended);
        assertFalse(auction.isAuctionActive(auctionId));

        // === Phase 4: Claiming Phase ===

        // Bob (winner) claims NFT
        vm.prank(bob);
        auction.claimNFT(auctionId);
        assertEq(nft.ownerOf(0), bob);

        // Alice (seller) claims funds
        uint256 platformFee = (bobSecondBid * 300) / 10000; // 3%
        uint256 sellerAmount = bobSecondBid - platformFee;

        uint256 aliceBalanceBefore = alice.balance;
        uint256 platformBalanceBefore = platformWallet.balance;

        vm.prank(alice);
        auction.claimFunds(auctionId);

        assertEq(alice.balance, aliceBalanceBefore + sellerAmount);
        assertEq(platformWallet.balance, platformBalanceBefore + platformFee);

        // === Final Verification ===
        auctionInfo = auction.getAuction(auctionId);
        assertTrue(auctionInfo.claimed);
        assertEq(auctionInfo.highestBid, 0); // Reset after claiming

        // Bob still has pending refund from first bid
        uint256 bobFinalBalance = bob.balance;
        vm.prank(bob);
        auction.withdrawRefund();
        assertEq(bob.balance, bobFinalBalance + RESERVE_PRICE);
    }

    function testCompleteAuctionWorkflowNoBids() public {
        // === Phase 1: Create Auction ===
        vm.prank(alice);
        uint256 auctionId = auction.createAuction(
            address(nft),
            1, // tokenId
            DURATION,
            RESERVE_PRICE
        );

        assertEq(nft.ownerOf(1), address(auction));

        // === Phase 2: No Bidding ===
        // Skip directly to end

        // === Phase 3: Auction End ===
        vm.warp(block.timestamp + DURATION + 1);
        auction.endAuction(auctionId);

        SimpleAuction.Auction memory auctionInfo = auction.getAuction(auctionId);
        assertTrue(auctionInfo.ended);
        assertEq(auctionInfo.highestBidder, address(0));

        // === Phase 4: Claiming Phase (No Winner) ===

        // Alice can claim back her NFT
        vm.prank(alice);
        auction.claimNFT(auctionId);
        assertEq(nft.ownerOf(1), alice);

        // No funds to claim (auction failed)
        vm.prank(alice);
        vm.expectRevert(SimpleAuction.InvalidAuction.selector);
        auction.claimFunds(auctionId);
    }

    function testMultipleSimultaneousAuctions() public {
        // Create 3 simultaneous auctions
        vm.startPrank(alice);
        uint256 auction1 = auction.createAuction(address(nft), 0, DURATION, 0.5 ether);
        uint256 auction2 = auction.createAuction(address(nft), 1, DURATION, 1.0 ether);
        uint256 auction3 = auction.createAuction(address(nft), 2, DURATION, 1.5 ether);
        vm.stopPrank();

        // Bob bids on auction 1
        vm.prank(bob);
        auction.placeBid{value: 0.5 ether}(auction1);

        // Charlie bids on auction 2
        vm.prank(charlie);
        auction.placeBid{value: 1.0 ether}(auction2);

        // Bob also bids on auction 3
        vm.prank(bob);
        auction.placeBid{value: 1.5 ether}(auction3);

        // Verify independent state
        SimpleAuction.Auction memory info1 = auction.getAuction(auction1);
        SimpleAuction.Auction memory info2 = auction.getAuction(auction2);
        SimpleAuction.Auction memory info3 = auction.getAuction(auction3);

        assertEq(info1.highestBidder, bob);
        assertEq(info1.highestBid, 0.5 ether);
        assertEq(info2.highestBidder, charlie);
        assertEq(info2.highestBid, 1.0 ether);
        assertEq(info3.highestBidder, bob);
        assertEq(info3.highestBid, 1.5 ether);

        // End all auctions
        vm.warp(block.timestamp + DURATION + 1);
        auction.endAuction(auction1);
        auction.endAuction(auction2);
        auction.endAuction(auction3);

        // Winners claim NFTs
        vm.prank(bob);
        auction.claimNFT(auction1);
        assertEq(nft.ownerOf(0), bob);

        vm.prank(charlie);
        auction.claimNFT(auction2);
        assertEq(nft.ownerOf(1), charlie);

        vm.prank(bob);
        auction.claimNFT(auction3);
        assertEq(nft.ownerOf(2), bob);

        // Alice claims all funds
        uint256 aliceBalanceBefore = alice.balance;

        vm.prank(alice);
        auction.claimFunds(auction1);

        vm.prank(alice);
        auction.claimFunds(auction2);

        vm.prank(alice);
        auction.claimFunds(auction3);

        // Calculate total expected
        uint256 totalBids = 0.5 ether + 1.0 ether + 1.5 ether;
        uint256 totalFees = (totalBids * 300) / 10000;
        uint256 aliceExpected = totalBids - totalFees;

        assertEq(alice.balance, aliceBalanceBefore + aliceExpected);
    }

    function testPriceVsETHIntegration() public {
        // Set price to $1500 per ETH
        mockPriceFeed.setPrice(1500e8);

        vm.prank(alice);
        uint256 auctionId = auction.createAuction(address(nft), 0, DURATION, 2 ether);

        // Check USD price display
        (bool success, uint256 usdPrice) = auction.getAuctionPriceInUSD(auctionId);
        assertTrue(success);
        assertEq(usdPrice, 3000e8); // 2 ETH * $1500 = $3000

        // Bid 3 ETH
        vm.prank(bob);
        auction.placeBid{value: 3 ether}(auctionId);

        // Check new USD price
        (success, usdPrice) = auction.getAuctionPriceInUSD(auctionId);
        assertTrue(success);
        assertEq(usdPrice, 4500e8); // 3 ETH * $1500 = $4500

        // Price changes during auction
        mockPriceFeed.setPrice(1800e8);

        (success, usdPrice) = auction.getAuctionPriceInUSD(auctionId);
        assertTrue(success);
        assertEq(usdPrice, 5400e8); // 3 ETH * $1800 = $5400
    }

    function testCompleteWorkflowWithPlatformWalletChange() public {
        // Create auction
        vm.prank(alice);
        uint256 auctionId = auction.createAuction(address(nft), 0, DURATION, 1 ether);

        // Bid
        vm.prank(bob);
        auction.placeBid{value: 1 ether}(auctionId);

        // Change platform wallet mid-auction
        address newPlatformWallet = address(0x888);
        auction.setPlatformWallet(newPlatformWallet);

        // End auction
        vm.warp(block.timestamp + DURATION + 1);
        auction.endAuction(auctionId);

        // Claim funds - should go to NEW platform wallet
        uint256 platformFee = (1 ether * 300) / 10000;

        vm.prank(alice);
        auction.claimFunds(auctionId);

        assertEq(newPlatformWallet.balance, platformFee);
        assertEq(platformWallet.balance, 0); // Old wallet gets nothing
    }

    function testStressTestBiddingWar() public {
        vm.prank(alice);
        uint256 auctionId = auction.createAuction(address(nft), 0, DURATION, 0.1 ether);

        uint256 currentBid = 0.1 ether;
        address currentWinner = bob;

        // Simulate 20 rounds of bidding between Bob and Charlie
        for (uint256 i = 0; i < 20; i++) {
            uint256 nextBid = currentBid + (currentBid * 500 / 10000); // 5% increment

            if (i % 2 == 0) {
                // Bob's turn
                vm.prank(bob);
                auction.placeBid{value: nextBid}(auctionId);
                currentWinner = bob;
            } else {
                // Charlie's turn
                vm.prank(charlie);
                auction.placeBid{value: nextBid}(auctionId);
                currentWinner = charlie;
            }

            currentBid = nextBid;
        }

        // Verify final state
        SimpleAuction.Auction memory auctionInfo = auction.getAuction(auctionId);
        assertEq(auctionInfo.highestBidder, currentWinner);
        assertEq(auctionInfo.highestBid, currentBid);

        // End and complete auction
        vm.warp(block.timestamp + DURATION + 1);
        auction.endAuction(auctionId);

        vm.prank(currentWinner);
        auction.claimNFT(auctionId);
        assertEq(nft.ownerOf(0), currentWinner);

        // Both bidders should have refunds to withdraw
        assertTrue(auction.pendingRefunds(bob) > 0 || auction.pendingRefunds(charlie) > 0);
    }

    function testRevertScenarios() public {
        // Test various failure scenarios in integration context

        // 1. Try to bid on ended auction
        vm.prank(alice);
        uint256 auctionId = auction.createAuction(address(nft), 0, DURATION, 1 ether);

        vm.warp(block.timestamp + DURATION + 1);

        vm.prank(bob);
        vm.expectRevert(SimpleAuction.AuctionAlreadyEnded.selector);
        auction.placeBid{value: 1 ether}(auctionId);

        // 2. Try to claim NFT before auction ends
        vm.warp(block.timestamp - DURATION); // Go back

        vm.prank(bob);
        auction.placeBid{value: 1 ether}(auctionId);

        vm.prank(bob);
        vm.expectRevert(SimpleAuction.AuctionStillActive.selector);
        auction.claimNFT(auctionId);

        // 3. Try to claim funds before ending
        vm.prank(alice);
        vm.expectRevert(SimpleAuction.AuctionStillActive.selector);
        auction.claimFunds(auctionId);

        // 4. Try to end auction early
        vm.expectRevert(SimpleAuction.AuctionStillActive.selector);
        auction.endAuction(auctionId);
    }
}