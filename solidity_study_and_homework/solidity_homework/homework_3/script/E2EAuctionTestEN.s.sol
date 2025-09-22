// SPDX-License-Identifier: MIT
pragma solidity ^0.8.7;

import "forge-std/Script.sol";
import "forge-std/console.sol";
import "../src/SimpleNFT.sol";
import "../src/SimpleAuction.sol";
import "../src/AuctionPriceFeed.sol";

/**
 * @title End-to-End Auction Test Script (English Version)
 * @notice Complete auction workflow test using 6 wallets on Sepolia testnet
 *
 * Test Flow:
 * 1. Main wallet creates NFT and starts auction (120s)
 * 2. Wallets 2-5 bid during auction period (under 0.001 ETH)
 * 3. Wallet 6 tries to bid after timeout (should fail)
 * 4. Verify Chainlink Price Feed data
 * 5. Test setPlatformWallet function
 * 6. Settle auction and verify fund distribution
 * 7. Verify refund mechanism
 */
contract E2EAuctionTestEN is Script {
    // Deployed contract addresses (with 2-minute MIN_DURATION on Sepolia)
    SimpleNFT constant NFT = SimpleNFT(0x8137e03e0d3F6C65db453Cc0EEeEE3C3c268a892);
    SimpleAuction constant AUCTION = SimpleAuction(0xD0FAa943812Fc9439639978a0599cEeD9b4AdE97);
    AuctionPriceFeed constant PRICE_FEED = AuctionPriceFeed(0x51ac07f0791C878Ce7C5f0223C7eB9F43376AC72);

    // Test parameters
    uint256 constant AUCTION_DURATION = 120; // 2 minutes auction time
    uint256 constant RESERVE_PRICE = 0.0001 ether; // Reserve price 0.0001 ETH
    uint256 constant BID_INCREMENT = 0.0001 ether; // Bid increment 0.0001 ETH

    // Wallet addresses (calculated from environment variables)
    address wallet1; // Main wallet - seller
    address wallet2; // Bidder 1
    address wallet3; // Bidder 2
    address wallet4; // Bidder 3
    address wallet5; // Bidder 4
    address wallet6; // Late bidder

    uint256 tokenId;
    uint256 auctionId;

    function run() external {
        console.log("=== Starting End-to-End Auction Test ===");

        // Calculate all wallet addresses
        _setupWallets();

        // Phase 1: Create NFT and auction
        _phase1_CreateAuction();

        // Phase 2: Bidding phase
        _phase2_BiddingPhase();

        // Phase 3: Verify Price Feed
        _phase3_VerifyPriceFeed();

        // Phase 4: Wait for auction end
        _phase4_WaitForAuctionEnd();

        // Phase 5: Test setPlatformWallet function
        _phase5_SetPlatformWallet();

        // Phase 6: Try timeout bid (should fail)
        _phase6_TimeoutBid();

        // Phase 7: Settlement phase
        _phase7_SettlementPhase();

        console.log("=== End-to-End Test Completed ===");
    }

    function _setupWallets() internal {
        wallet1 = vm.addr(vm.envUint("PRIVATE_KEY"));
        wallet2 = vm.addr(vm.envUint("PRIVATE_KEY_2"));
        wallet3 = vm.addr(vm.envUint("PRIVATE_KEY_3"));
        wallet4 = vm.addr(vm.envUint("PRIVATE_KEY_4"));
        wallet5 = vm.addr(vm.envUint("PRIVATE_KEY_5"));
        wallet6 = vm.addr(vm.envUint("PRIVATE_KEY_6"));

        console.log("Wallet addresses setup:");
        console.log("Wallet1 (seller):", wallet1);
        console.log("Wallet2 (bidder1):", wallet2);
        console.log("Wallet3 (bidder2):", wallet3);
        console.log("Wallet4 (bidder3):", wallet4);
        console.log("Wallet5 (bidder4):", wallet5);
        console.log("Wallet6 (late bidder):", wallet6);
    }

    function _phase1_CreateAuction() internal {
        console.log("\n=== Phase 1: Create NFT and Auction ===");

        vm.startBroadcast(vm.envUint("PRIVATE_KEY"));

        // 1.1 Test SimpleNFT.mint function explicitly
        console.log("1.1 Testing SimpleNFT.mint function");
        uint256 totalSupplyBefore = NFT.totalSupply();
        tokenId = NFT.mint(wallet1);
        uint256 totalSupplyAfter = NFT.totalSupply();

        require(totalSupplyAfter == totalSupplyBefore + 1, "Mint should increase total supply");
        require(NFT.ownerOf(tokenId) == wallet1, "Minted NFT should belong to specified address");
        console.log("[OK] NFT Mint function verified, TokenID:", tokenId);
        console.log("   Owner:", NFT.ownerOf(tokenId));
        console.log("   Total supply:", totalSupplyAfter);

        // 1.2 Test SimpleNFT.setApprovalForAll function explicitly
        console.log("1.2 Testing SimpleNFT.setApprovalForAll function");
        bool approvedBefore = NFT.isApprovedForAll(wallet1, address(AUCTION));
        NFT.setApprovalForAll(address(AUCTION), true);
        bool approvedAfter = NFT.isApprovedForAll(wallet1, address(AUCTION));

        require(!approvedBefore && approvedAfter, "Approval state should change from false to true");
        console.log("[OK] NFT approval function verified");
        console.log("   Auction contract authorized for all NFTs");

        // 1.3 Create auction
        console.log("1.3 Creating auction");
        auctionId = AUCTION.createAuction(
            address(NFT),
            tokenId,
            AUCTION_DURATION,
            RESERVE_PRICE
        );

        console.log("[OK] Auction created successfully, AuctionID:", auctionId);
        console.log("   Duration:", AUCTION_DURATION, "seconds");
        console.log("   Reserve price:", RESERVE_PRICE, "Wei");

        // 1.4 Verify NFT ownership transfer
        require(NFT.ownerOf(tokenId) == address(AUCTION), "NFT should be transferred to auction contract");
        require(AUCTION.isAuctionActive(auctionId), "Auction should be active");
        console.log("[OK] NFT escrowed to auction contract");
        console.log("   Auction status: Active");

        vm.stopBroadcast();
    }

    function _phase2_BiddingPhase() internal {
        console.log("\n=== Phase 2: Bidding Phase ===");

        uint256 currentBid = RESERVE_PRICE;

        // 2.1 Wallet2 first bid (reserve price)
        vm.broadcast(vm.envUint("PRIVATE_KEY_2"));
        AUCTION.placeBid{value: currentBid}(auctionId);
        console.log("Wallet2 bid:", currentBid, "Wei");

        // 2.2 Wallet3 bid (increment)
        currentBid += BID_INCREMENT;
        vm.broadcast(vm.envUint("PRIVATE_KEY_3"));
        AUCTION.placeBid{value: currentBid}(auctionId);
        console.log("Wallet3 bid:", currentBid, "Wei");

        // 2.3 Wallet4 bid (continue increment)
        currentBid += BID_INCREMENT;
        vm.broadcast(vm.envUint("PRIVATE_KEY_4"));
        AUCTION.placeBid{value: currentBid}(auctionId);
        console.log("Wallet4 bid:", currentBid, "Wei");

        // 2.4 Wallet5 final bid (highest)
        currentBid += BID_INCREMENT;
        vm.broadcast(vm.envUint("PRIVATE_KEY_5"));
        AUCTION.placeBid{value: currentBid}(auctionId);
        console.log("Wallet5 bid:", currentBid, "Wei (current highest)");

        // 2.5 Verify auction status
        SimpleAuction.Auction memory auction = AUCTION.getAuction(auctionId);
        require(auction.highestBidder == wallet5, "Wallet5 should be current highest bidder");
        require(auction.highestBid == currentBid, "Highest bid should match");
        console.log("Current highest bidder:", auction.highestBidder);
        console.log("Current highest bid:", auction.highestBid, "Wei");
    }

    function _phase3_VerifyPriceFeed() internal {
        console.log("\n=== Phase 3: Verify Chainlink Price Feed ===");

        SimpleAuction.Auction memory auction = AUCTION.getAuction(auctionId);
        uint256 currentHighestBid = auction.highestBid;

        // 3.1 Test AuctionPriceFeed.getEthUsdPrice function explicitly
        console.log("3.1 Testing AuctionPriceFeed.getEthUsdPrice function");
        console.log("   Current highest bid (Wei):", currentHighestBid);

        uint256 ethPriceFromFeed = PRICE_FEED.getEthUsdPrice(currentHighestBid);
        require(ethPriceFromFeed > 0, "ETH price should be greater than 0");

        console.log("[OK] PriceFeed.getEthUsdPrice function verified");
        console.log("   ETH to USD conversion verified");

        // 3.2 Test price calculation for different ETH amounts
        console.log("3.2 Testing price calculation for different ETH amounts");

        uint256 oneEthPrice = PRICE_FEED.getEthUsdPrice(1 ether);
        uint256 halfEthPrice = PRICE_FEED.getEthUsdPrice(0.5 ether);
        uint256 minBidPrice = PRICE_FEED.getEthUsdPrice(RESERVE_PRICE);

        console.log("   1 ETH price:", oneEthPrice);
        console.log("   0.5 ETH price:", halfEthPrice);
        console.log("   Reserve price USD:", minBidPrice);

        require(oneEthPrice == halfEthPrice * 2, "Price calculation should be linear");
        console.log("[OK] Linear price calculation verified");

        // 3.3 Verify price consistency through auction contract
        console.log("3.3 Verifying auction contract price integration");
        (bool success, uint256 auctionUsdPrice) = AUCTION.getAuctionPriceInUSD(auctionId);
        require(success, "Auction price retrieval should succeed");
        require(auctionUsdPrice == ethPriceFromFeed, "Auction contract and PriceFeed prices should match");

        console.log("[OK] Auction-PriceFeed integration verified");
        console.log("   Current auction price:", auctionUsdPrice, "USD");
    }

    function _phase4_WaitForAuctionEnd() internal {
        console.log("\n=== Phase 4: Wait for Auction End ===");

        // Note: Cannot actually wait in script, need manual wait or staged execution
        console.log("Please wait 2 minutes before continuing...");
        console.log("Or manually call endAuction function");

        // This should wait for auction time to end, but scripts cannot sleep
        // In actual testing need to execute script in two stages or wait manually
    }

    function _phase5_SetPlatformWallet() internal {
        console.log("\n=== Phase 5: Test setPlatformWallet Function ===");

        // 5.1 Record original platform wallet (set during contract deployment)
        // From deployment script, original platform wallet is 0x5D4a329B262AC7a1d9Ae0f4C54171dF61E2c0b69
        address originalPlatformWallet = 0x5D4a329B262AC7a1d9Ae0f4C54171dF61E2c0b69;

        // 5.2 Set new platform wallet (use wallet6 as new platform wallet)
        console.log("5.1 Testing SimpleAuction.setPlatformWallet function");
        console.log("   Original platform wallet:", originalPlatformWallet);
        console.log("   New platform wallet:", wallet6);

        vm.broadcast(vm.envUint("PRIVATE_KEY")); // Only owner can call
        AUCTION.setPlatformWallet(wallet6);

        console.log("[OK] Platform wallet set successfully");
        console.log("   Note: Platform fees during settlement will go to new wallet");

        // 5.3 Verify setting took effect
        // Note: Cannot directly read platformWallet state variable, can only verify during settlement
        console.log("5.2 Platform wallet change will be verified in settlement phase");
    }

    function _phase6_TimeoutBid() internal {
        console.log("\n=== Phase 6: Timeout Bid Test ===");

        // 6.1 Check if auction has ended
        if (AUCTION.isAuctionActive(auctionId)) {
            console.log("Auction still active, skipping timeout test");
            return;
        }

        // 6.2 Try timeout bid (should fail)
        console.log("Wallet6 attempting timeout bid...");
        vm.expectRevert();
        vm.broadcast(vm.envUint("PRIVATE_KEY_6"));
        AUCTION.placeBid{value: 0.001 ether}(auctionId);
        console.log("Timeout bid correctly rejected");
    }

    function _phase7_SettlementPhase() internal {
        console.log("\n=== Phase 7: Settlement Phase ===");

        SimpleAuction.Auction memory auction = AUCTION.getAuction(auctionId);
        uint256 finalBid = auction.highestBid;
        uint256 expectedPlatformFee = (finalBid * 300) / 10000; // 3% fee rate
        address originalPlatformWallet = 0x5D4a329B262AC7a1d9Ae0f4C54171dF61E2c0b69;

        console.log("7.1 Recording pre-settlement balances");
        console.log("   Final winning bid:", finalBid, "Wei");
        console.log("   Expected platform fee:", expectedPlatformFee, "Wei");
        uint256 wallet6BalanceBefore = wallet6.balance;
        uint256 originalPlatformBalanceBefore = originalPlatformWallet.balance;
        console.log("   New platform wallet balance (pre-settlement):", wallet6BalanceBefore, "Wei");
        console.log("   Original platform wallet balance (pre-settlement):", originalPlatformBalanceBefore, "Wei");

        // 7.2 End auction (if still active)
        if (AUCTION.isAuctionActive(auctionId)) {
            vm.broadcast(vm.envUint("PRIVATE_KEY"));
            AUCTION.endAuction(auctionId);
            console.log("[OK] Auction ended");
        }

        // 7.3 Winner claims NFT
        console.log("7.2 Wallet5 (winner) claiming NFT");
        vm.broadcast(vm.envUint("PRIVATE_KEY_5"));
        AUCTION.claimNFT(auctionId);

        require(NFT.ownerOf(tokenId) == wallet5, "NFT should be transferred to winner");
        console.log("[OK] NFT successfully transferred to winner:", wallet5);

        // 7.4 Seller claims funds (platform fee should go to new platform wallet)
        console.log("7.3 Wallet1 (seller) claiming funds");
        uint256 seller1BalanceBefore = wallet1.balance;

        vm.broadcast(vm.envUint("PRIVATE_KEY"));
        AUCTION.claimFunds(auctionId);

        uint256 seller1BalanceAfter = wallet1.balance;
        uint256 sellerReceived = seller1BalanceAfter - seller1BalanceBefore;
        console.log("   Seller received funds:", sellerReceived, "Wei");

        // 7.5 Verify setPlatformWallet function took effect
        console.log("7.4 Verifying setPlatformWallet function effect");
        uint256 wallet6BalanceAfter = wallet6.balance;
        uint256 originalPlatformBalanceAfter = originalPlatformWallet.balance;
        uint256 platformFeeReceived = wallet6BalanceAfter - wallet6BalanceBefore;

        console.log("   New platform wallet balance (post-settlement):", wallet6BalanceAfter, "Wei");
        console.log("   Original platform wallet balance (post-settlement):", originalPlatformBalanceAfter, "Wei");
        console.log("   New platform wallet received fee:", platformFeeReceived, "Wei");
        console.log("   Original platform wallet balance change:", originalPlatformBalanceAfter - originalPlatformBalanceBefore, "Wei");

        require(platformFeeReceived == expectedPlatformFee, "New platform wallet should receive correct platform fee");
        require(originalPlatformBalanceAfter == originalPlatformBalanceBefore, "Original platform wallet balance should not increase");

        console.log("[OK] setPlatformWallet verified");
        console.log("   Platform fee sent to new wallet");

        // 7.6 Failed bidders claim refunds
        console.log("7.5 Failed bidders claiming refunds");

        // Wallet2 claims refund
        uint256 refund2 = AUCTION.pendingRefunds(wallet2);
        if (refund2 > 0) {
            vm.broadcast(vm.envUint("PRIVATE_KEY_2"));
            AUCTION.withdrawRefund();
            console.log("   Wallet2 refund:", refund2, "Wei");
        }

        // Wallet3 claims refund
        uint256 refund3 = AUCTION.pendingRefunds(wallet3);
        if (refund3 > 0) {
            vm.broadcast(vm.envUint("PRIVATE_KEY_3"));
            AUCTION.withdrawRefund();
            console.log("   Wallet3 refund:", refund3, "Wei");
        }

        // Wallet4 claims refund
        uint256 refund4 = AUCTION.pendingRefunds(wallet4);
        if (refund4 > 0) {
            vm.broadcast(vm.envUint("PRIVATE_KEY_4"));
            AUCTION.withdrawRefund();
            console.log("   Wallet4 refund:", refund4, "Wei");
        }

        console.log("[OK] All refunds completed");
    }
}