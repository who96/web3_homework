// SPDX-License-Identifier: MIT
pragma solidity ^0.8.7;

import "forge-std/Script.sol";
import "forge-std/console.sol";
import "../src/SimpleNFT.sol";
import "../src/AuctionPriceFeed.sol";
import "../src/SimpleAuction.sol";
import "./HelperConfig.sol";

contract DeployContracts is Script {
    SimpleNFT public simpleNFT;
    AuctionPriceFeed public priceFeed;
    SimpleAuction public auction;
    HelperConfig public helperConfig;

    function run() external returns (SimpleNFT, AuctionPriceFeed, SimpleAuction) {
        helperConfig = new HelperConfig();
        (
            address oracle,
            bytes32 jobId,
            uint256 chainlinkFee,
            address link,
            uint256 updateInterval,
            address chainlinkPriceFeed,
            uint64 subscriptionId,
            address vrfCoordinator,
            bytes32 keyHash
        ) = helperConfig.activeNetworkConfig();

        address platformWallet = msg.sender; // Use deployer as platform wallet

        vm.startBroadcast();

        // 1. Deploy SimpleNFT contract
        simpleNFT = new SimpleNFT();
        console.log("SimpleNFT deployed at:", address(simpleNFT));

        // 2. Deploy AuctionPriceFeed contract with Chainlink price feed
        priceFeed = new AuctionPriceFeed(chainlinkPriceFeed);
        console.log("AuctionPriceFeed deployed at:", address(priceFeed));
        console.log("Using price feed address:", chainlinkPriceFeed);

        // 3. Deploy SimpleAuction contract
        auction = new SimpleAuction(address(priceFeed), platformWallet);
        console.log("SimpleAuction deployed at:", address(auction));
        console.log("Platform wallet set to:", platformWallet);

        vm.stopBroadcast();

        // Verify deployment
        _verifyDeployment();

        return (simpleNFT, priceFeed, auction);
    }

    function _verifyDeployment() internal view {
        console.log("\n=== Deployment Verification ===");

        // Verify SimpleNFT
        require(address(simpleNFT) != address(0), "SimpleNFT deployment failed");
        require(simpleNFT.totalSupply() == 0, "SimpleNFT initial supply should be 0");
        console.log("+ SimpleNFT verification passed");

        // Verify AuctionPriceFeed
        require(address(priceFeed) != address(0), "AuctionPriceFeed deployment failed");
        require(priceFeed.DECIMALS() == 8, "Price feed decimals should be 8");
        console.log("+ AuctionPriceFeed verification passed");

        // Verify SimpleAuction
        require(address(auction) != address(0), "SimpleAuction deployment failed");
        require(auction.auctionCounter() == 0, "Auction counter should be 0");
        require(auction.PLATFORM_FEE() == 300, "Platform fee should be 3%");
        require(auction.MIN_BID_INCREMENT() == 500, "Min bid increment should be 5%");
        require(auction.MIN_DURATION() == 120, "Min duration should be 2 minutes");
        console.log("+ SimpleAuction verification passed");

        console.log("SUCCESS: All contracts deployed and verified!");
    }
}