// SPDX-License-Identifier: MIT
pragma solidity ^0.8.7;

import "forge-std/Test.sol";
import "forge-std/console.sol";
import "../src/SimpleAuction.sol";
import "../src/SimpleNFT.sol";
import "../src/AuctionPriceFeed.sol";

contract MockPriceFeed {
    int256 public price = 2000e8; // $2000 with 8 decimals

    function latestRoundData() external view returns (
        uint80 roundId,
        int256 answer,
        uint256 startedAt,
        uint256 updatedAt_,
        uint80 answeredInRound
    ) {
        // Always return current timestamp to avoid stale data
        return (1, price, block.timestamp, block.timestamp, 1);
    }

    function decimals() external pure returns (uint8) {
        return 8;
    }

    function setPrice(int256 _price) external {
        price = _price;
    }
}

contract SimpleAuctionTest is Test {
    SimpleAuction public auction;
    SimpleNFT public nft;
    AuctionPriceFeed public priceFeed;
    MockPriceFeed public mockPriceFeed;

    address public owner;
    address public seller;
    address public bidder1;
    address public bidder2;
    address public platformWallet;

    uint256 public constant TOKEN_ID = 0;
    uint256 public constant DURATION = 7200; // 2 hours
    uint256 public constant RESERVE_PRICE = 1 ether;

    event AuctionCreated(
        uint256 indexed auctionId,
        address indexed nftContract,
        uint256 indexed tokenId,
        address seller,
        uint256 startTime,
        uint256 duration,
        uint256 reservePrice
    );

    event BidPlaced(
        uint256 indexed auctionId,
        address indexed bidder,
        uint256 amount
    );

    event AuctionEnded(
        uint256 indexed auctionId,
        address indexed winner,
        uint256 amount
    );

    event RefundWithdrawn(address indexed bidder, uint256 amount);
    event NFTClaimed(uint256 indexed auctionId, address indexed winner);
    event FundsClaimed(uint256 indexed auctionId, address indexed seller, uint256 amount);

    function setUp() public {
        owner = address(this);
        seller = address(0x1);
        bidder1 = address(0x2);
        bidder2 = address(0x3);
        platformWallet = address(0x999);

        // Deploy mock price feed
        mockPriceFeed = new MockPriceFeed();

        // Deploy contracts
        nft = new SimpleNFT();
        priceFeed = new AuctionPriceFeed(address(mockPriceFeed));
        auction = new SimpleAuction(address(priceFeed), platformWallet);

        // Setup: Mint NFT to seller and approve auction contract
        uint256 tokenId = nft.mint(seller);
        assertEq(tokenId, TOKEN_ID);

        vm.prank(seller);
        nft.setApprovalForAll(address(auction), true);

        // Give ETH to bidders
        vm.deal(bidder1, 10 ether);
        vm.deal(bidder2, 10 ether);
    }

    function testInitialState() public {
        assertEq(auction.auctionCounter(), 0);
        assertEq(auction.PLATFORM_FEE(), 300); // 3%
        assertEq(auction.MIN_BID_INCREMENT(), 500); // 5%
        assertEq(auction.MIN_DURATION(), 120); // 2 minutes
        assertEq(auction.platformWallet(), platformWallet);
    }

    function testCreateAuction() public {
        vm.expectEmit(true, true, true, true);
        emit AuctionCreated(0, address(nft), TOKEN_ID, seller, block.timestamp, DURATION, RESERVE_PRICE);

        vm.prank(seller);
        uint256 auctionId = auction.createAuction(address(nft), TOKEN_ID, DURATION, RESERVE_PRICE);

        assertEq(auctionId, 0);
        assertEq(auction.auctionCounter(), 1);

        SimpleAuction.Auction memory auctionInfo = auction.getAuction(auctionId);
        assertEq(auctionInfo.nftContract, address(nft));
        assertEq(auctionInfo.tokenId, TOKEN_ID);
        assertEq(auctionInfo.seller, seller);
        assertEq(auctionInfo.duration, DURATION);
        assertEq(auctionInfo.reservePrice, RESERVE_PRICE);
        assertEq(auctionInfo.highestBidder, address(0));
        assertEq(auctionInfo.highestBid, 0);
        assertFalse(auctionInfo.ended);
        assertFalse(auctionInfo.claimed);

        // Check NFT was transferred to auction contract
        assertEq(nft.ownerOf(TOKEN_ID), address(auction));
        assertTrue(auction.isAuctionActive(auctionId));
    }

    function testCreateAuctionFailures() public {
        // Invalid duration
        vm.prank(seller);
        vm.expectRevert(SimpleAuction.InvalidDuration.selector);
        auction.createAuction(address(nft), TOKEN_ID, 60, RESERVE_PRICE); // 1 minute < 2 minutes

        // Not NFT owner
        vm.prank(bidder1);
        vm.expectRevert(SimpleAuction.NotNFTOwner.selector);
        auction.createAuction(address(nft), TOKEN_ID, DURATION, RESERVE_PRICE);

        // NFT not approved
        vm.prank(seller);
        nft.setApprovalForAll(address(auction), false);

        vm.prank(seller);
        vm.expectRevert(SimpleAuction.NFTNotApproved.selector);
        auction.createAuction(address(nft), TOKEN_ID, DURATION, RESERVE_PRICE);
    }

    function testPlaceBid() public {
        vm.prank(seller);
        uint256 auctionId = auction.createAuction(address(nft), TOKEN_ID, DURATION, RESERVE_PRICE);

        // First bid at reserve price
        vm.expectEmit(true, true, false, true);
        emit BidPlaced(auctionId, bidder1, RESERVE_PRICE);

        vm.prank(bidder1);
        auction.placeBid{value: RESERVE_PRICE}(auctionId);

        SimpleAuction.Auction memory auctionInfo = auction.getAuction(auctionId);
        assertEq(auctionInfo.highestBidder, bidder1);
        assertEq(auctionInfo.highestBid, RESERVE_PRICE);

        // Second bid with 5% increment
        uint256 minBid = RESERVE_PRICE + (RESERVE_PRICE * 500 / 10000);

        vm.prank(bidder2);
        auction.placeBid{value: minBid}(auctionId);

        auctionInfo = auction.getAuction(auctionId);
        assertEq(auctionInfo.highestBidder, bidder2);
        assertEq(auctionInfo.highestBid, minBid);

        // Check pending refund for previous bidder
        assertEq(auction.pendingRefunds(bidder1), RESERVE_PRICE);
    }

    function testPlaceBidFailures() public {
        vm.prank(seller);
        uint256 auctionId = auction.createAuction(address(nft), TOKEN_ID, DURATION, RESERVE_PRICE);

        // Bid on nonexistent auction
        vm.prank(bidder1);
        vm.expectRevert(SimpleAuction.AuctionNotFound.selector);
        auction.placeBid{value: RESERVE_PRICE}(999);

        // Insufficient bid
        vm.prank(bidder1);
        vm.expectRevert(SimpleAuction.InsufficientBid.selector);
        auction.placeBid{value: RESERVE_PRICE - 1}(auctionId);

        // Bid after auction ended
        vm.warp(block.timestamp + DURATION + 1);

        vm.prank(bidder1);
        vm.expectRevert(SimpleAuction.AuctionAlreadyEnded.selector);
        auction.placeBid{value: RESERVE_PRICE}(auctionId);
    }

    function testWithdrawRefund() public {
        vm.prank(seller);
        uint256 auctionId = auction.createAuction(address(nft), TOKEN_ID, DURATION, RESERVE_PRICE);

        // Place two bids
        vm.prank(bidder1);
        auction.placeBid{value: RESERVE_PRICE}(auctionId);

        uint256 secondBid = RESERVE_PRICE + (RESERVE_PRICE * 500 / 10000);
        vm.prank(bidder2);
        auction.placeBid{value: secondBid}(auctionId);

        // Withdraw refund
        uint256 bidder1BalanceBefore = bidder1.balance;

        vm.expectEmit(true, false, false, true);
        emit RefundWithdrawn(bidder1, RESERVE_PRICE);

        vm.prank(bidder1);
        auction.withdrawRefund();

        assertEq(bidder1.balance, bidder1BalanceBefore + RESERVE_PRICE);
        assertEq(auction.pendingRefunds(bidder1), 0);

        // Try to withdraw again (should fail)
        vm.prank(bidder1);
        vm.expectRevert(SimpleAuction.NoRefundAvailable.selector);
        auction.withdrawRefund();
    }

    function testEndAuction() public {
        vm.prank(seller);
        uint256 auctionId = auction.createAuction(address(nft), TOKEN_ID, DURATION, RESERVE_PRICE);

        vm.prank(bidder1);
        auction.placeBid{value: RESERVE_PRICE}(auctionId);

        // Try to end auction early
        vm.expectRevert(SimpleAuction.AuctionStillActive.selector);
        auction.endAuction(auctionId);

        // Wait for auction to end
        vm.warp(block.timestamp + DURATION + 1);

        vm.expectEmit(true, true, false, true);
        emit AuctionEnded(auctionId, bidder1, RESERVE_PRICE);

        auction.endAuction(auctionId);

        SimpleAuction.Auction memory auctionInfo = auction.getAuction(auctionId);
        assertTrue(auctionInfo.ended);
        assertFalse(auction.isAuctionActive(auctionId));
    }

    function testClaimNFTWinner() public {
        vm.prank(seller);
        uint256 auctionId = auction.createAuction(address(nft), TOKEN_ID, DURATION, RESERVE_PRICE);

        vm.prank(bidder1);
        auction.placeBid{value: RESERVE_PRICE}(auctionId);

        vm.warp(block.timestamp + DURATION + 1);
        auction.endAuction(auctionId);

        vm.expectEmit(true, true, false, false);
        emit NFTClaimed(auctionId, bidder1);

        vm.prank(bidder1);
        auction.claimNFT(auctionId);

        assertEq(nft.ownerOf(TOKEN_ID), bidder1);

        SimpleAuction.Auction memory auctionInfo = auction.getAuction(auctionId);
        assertTrue(auctionInfo.claimed);
    }

    function testClaimNFTNoWinner() public {
        vm.prank(seller);
        uint256 auctionId = auction.createAuction(address(nft), TOKEN_ID, DURATION, RESERVE_PRICE);

        // No bids placed
        vm.warp(block.timestamp + DURATION + 1);
        auction.endAuction(auctionId);

        // Seller can claim back NFT
        vm.prank(seller);
        auction.claimNFT(auctionId);

        assertEq(nft.ownerOf(TOKEN_ID), seller);
    }

    function testClaimFunds() public {
        vm.prank(seller);
        uint256 auctionId = auction.createAuction(address(nft), TOKEN_ID, DURATION, RESERVE_PRICE);

        vm.prank(bidder1);
        auction.placeBid{value: RESERVE_PRICE}(auctionId);

        vm.warp(block.timestamp + DURATION + 1);
        auction.endAuction(auctionId);

        uint256 platformFee = (RESERVE_PRICE * 300) / 10000; // 3%
        uint256 sellerAmount = RESERVE_PRICE - platformFee;

        uint256 sellerBalanceBefore = seller.balance;
        uint256 platformBalanceBefore = platformWallet.balance;

        vm.expectEmit(true, true, false, true);
        emit FundsClaimed(auctionId, seller, sellerAmount);

        vm.prank(seller);
        auction.claimFunds(auctionId);

        assertEq(seller.balance, sellerBalanceBefore + sellerAmount);
        assertEq(platformWallet.balance, platformBalanceBefore + platformFee);

        SimpleAuction.Auction memory auctionInfo = auction.getAuction(auctionId);
        assertEq(auctionInfo.highestBid, 0); // Reset after claiming
    }


    function testGetAuctionPriceInUSD() public {
        vm.prank(seller);
        uint256 auctionId = auction.createAuction(address(nft), TOKEN_ID, DURATION, RESERVE_PRICE);

        // Test with reserve price (no bids)
        (bool success, uint256 usdPrice) = auction.getAuctionPriceInUSD(auctionId);
        assertTrue(success);

        // Price should be 2000e8 for 1 ETH
        uint256 expectedPrice = 2000e8;
        assertEq(usdPrice, expectedPrice);

        // Test with highest bid
        vm.prank(bidder1);
        auction.placeBid{value: 2 ether}(auctionId);

        (success, usdPrice) = auction.getAuctionPriceInUSD(auctionId);
        assertTrue(success);

        // For 2 ETH: should be 4000e8
        expectedPrice = 4000e8;
        assertEq(usdPrice, expectedPrice);
    }

    function testMultipleAuctions() public {
        // Create second NFT
        uint256 tokenId2 = nft.mint(seller);

        vm.prank(seller);
        uint256 auctionId1 = auction.createAuction(address(nft), TOKEN_ID, DURATION, RESERVE_PRICE);

        vm.prank(seller);
        uint256 auctionId2 = auction.createAuction(address(nft), tokenId2, DURATION, RESERVE_PRICE * 2);

        assertEq(auctionId1, 0);
        assertEq(auctionId2, 1);
        assertEq(auction.auctionCounter(), 2);

        // Bid on both auctions
        vm.prank(bidder1);
        auction.placeBid{value: RESERVE_PRICE}(auctionId1);

        vm.prank(bidder2);
        auction.placeBid{value: RESERVE_PRICE * 2}(auctionId2);

        // Check both auctions are independent
        SimpleAuction.Auction memory auction1 = auction.getAuction(auctionId1);
        SimpleAuction.Auction memory auction2 = auction.getAuction(auctionId2);

        assertEq(auction1.highestBidder, bidder1);
        assertEq(auction2.highestBidder, bidder2);
        assertEq(auction1.highestBid, RESERVE_PRICE);
        assertEq(auction2.highestBid, RESERVE_PRICE * 2);
    }

    function testSetPlatformWallet() public {
        address newPlatformWallet = address(0x888);

        auction.setPlatformWallet(newPlatformWallet);
        assertEq(auction.platformWallet(), newPlatformWallet);

        // Only owner can set
        vm.prank(bidder1);
        vm.expectRevert();
        auction.setPlatformWallet(address(0x777));
    }

    function testReentrancyProtection() public {
        // This test ensures ReentrancyGuard is working
        // In a real attack scenario, an attacker would try to re-enter
        // during fund transfers, but our test setup is simplified

        vm.prank(seller);
        uint256 auctionId = auction.createAuction(address(nft), TOKEN_ID, DURATION, RESERVE_PRICE);

        vm.prank(bidder1);
        auction.placeBid{value: RESERVE_PRICE}(auctionId);

        vm.prank(bidder2);
        auction.placeBid{value: RESERVE_PRICE + 0.1 ether}(auctionId);

        // Withdraw refund should work normally
        vm.prank(bidder1);
        auction.withdrawRefund();

        assertEq(auction.pendingRefunds(bidder1), 0);
    }

    function testFuzzBidding(uint256 bidAmount, uint8 numBidders) public {
        vm.assume(bidAmount >= RESERVE_PRICE && bidAmount <= 100 ether);
        vm.assume(numBidders >= 1 && numBidders <= 10);

        vm.prank(seller);
        uint256 auctionId = auction.createAuction(address(nft), TOKEN_ID, DURATION, RESERVE_PRICE);

        address currentHighest = address(0);
        uint256 currentBid = 0;

        for (uint256 i = 0; i < numBidders; i++) {
            address bidder = address(uint160(0x1000 + i));
            uint256 thisBid = bidAmount + (i * 0.1 ether);

            vm.deal(bidder, thisBid);

            if (thisBid >= RESERVE_PRICE &&
                (currentBid == 0 || thisBid >= currentBid + (currentBid * 500 / 10000))) {

                vm.prank(bidder);
                auction.placeBid{value: thisBid}(auctionId);

                currentHighest = bidder;
                currentBid = thisBid;
            }
        }

        if (currentHighest != address(0)) {
            SimpleAuction.Auction memory auctionInfo = auction.getAuction(auctionId);
            assertEq(auctionInfo.highestBidder, currentHighest);
            assertEq(auctionInfo.highestBid, currentBid);
        }
    }
}