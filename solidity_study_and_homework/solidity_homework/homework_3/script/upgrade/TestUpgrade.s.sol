// SPDX-License-Identifier: MIT
pragma solidity ^0.8.7;

import "forge-std/Script.sol";
import "forge-std/console.sol";
import "../../src/upgradeable/SimpleAuctionUpgradeable.sol";
import "../../src/upgradeable/SimpleAuctionUpgradeableV2.sol";
import "@openzeppelin/contracts/proxy/transparent/ProxyAdmin.sol";
import "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";
import "../../src/SimpleNFT.sol";

/**
 * @title TestUpgrade
 * @notice Test upgrade functionality on Sepolia with simplified proxy architecture
 */
contract TestUpgrade is Script {
    // Deployed contract addresses from DeployProxySimple
    address constant SIMPLE_NFT = 0x6d61687CDE7F12A9c31aD833b305EF0c65EA830b;
    address constant AUCTION_PRICE_FEED = 0x15c8CED44bbdc1fec603BB9147EA51Dcc0809d7a;
    address constant AUCTION_PROXY = 0x687265FBABe670a18D0274478984D6c9a03CCCb6;
    address constant INTERNAL_PROXY_ADMIN = 0x17D768939362FdfE8e3EC586A15e258E270d7BE7;

    SimpleNFT public nft;
    ProxyAdmin public proxyAdmin;
    SimpleAuctionUpgradeable public auction;
    SimpleAuctionUpgradeableV2 public auctionV2Impl;

    function run() external {
        nft = SimpleNFT(SIMPLE_NFT);
        auction = SimpleAuctionUpgradeable(AUCTION_PROXY);

        // Use the known internal ProxyAdmin address from DeployProxySimple
        proxyAdmin = ProxyAdmin(INTERNAL_PROXY_ADMIN);
        console.log("Using internal ProxyAdmin address:", address(proxyAdmin));

        vm.startBroadcast();

        console.log("=== Upgrade Test on Sepolia ===");

        // Step 1: Check current version
        _checkCurrentVersion();

        // Step 2: Create auction before upgrade to test data persistence
        uint256 auctionId = _createTestAuction();

        // Step 3: Deploy V2 implementation
        _deployV2Implementation();

        // Step 4: Perform upgrade
        _performUpgrade();

        // Step 5: Verify upgrade and data persistence
        _verifyUpgrade(auctionId);

        // Step 6: Test new V2 functionality
        _testV2Functionality();

        vm.stopBroadcast();

        console.log("SUCCESS: Upgrade test completed successfully!");
    }

    function _checkCurrentVersion() internal view {
        string memory currentVersion = auction.getVersion();
        console.log("Current version:", currentVersion);
        require(
            keccak256(abi.encodePacked(currentVersion)) == keccak256(abi.encodePacked("1.0.0")),
            "Expected version 1.0.0"
        );
    }

    function _createTestAuction() internal returns (uint256) {
        console.log("Creating test auction before upgrade...");

        // Mint NFT to deployer
        uint256 tokenId = nft.mint(msg.sender);
        console.log("Minted NFT with tokenId:", tokenId);

        // Approve auction contract to transfer NFT
        nft.setApprovalForAll(address(auction), true);

        // Create auction
        uint256 auctionId = auction.createAuction(
            address(nft),
            tokenId,
            3600, // 1 hour
            0.1 ether
        );

        console.log("Created auction with ID:", auctionId);
        console.log("Auction counter before upgrade:", auction.auctionCounter());

        return auctionId;
    }

    function _deployV2Implementation() internal {
        console.log("Deploying V2 implementation...");
        auctionV2Impl = new SimpleAuctionUpgradeableV2();
        console.log("V2 implementation deployed at:", address(auctionV2Impl));
    }

    function _performUpgrade() internal {
        console.log("Performing upgrade to V2...");

        proxyAdmin.upgradeAndCall(
            ITransparentUpgradeableProxy(address(auction)),
            address(auctionV2Impl),
            ""
        );

        console.log("Upgrade completed!");
    }

    function _verifyUpgrade(uint256 auctionId) internal view {
        console.log("Verifying upgrade and data persistence...");

        // Check version is now 2.0.0
        SimpleAuctionUpgradeableV2 auctionV2 = SimpleAuctionUpgradeableV2(address(auction));
        string memory newVersion = auctionV2.getVersion();
        console.log("New version:", newVersion);
        require(
            keccak256(abi.encodePacked(newVersion)) == keccak256(abi.encodePacked("2.0.0")),
            "Expected version 2.0.0"
        );

        // Check auction data is preserved
        console.log("Auction counter after upgrade:", auction.auctionCounter());
        require(auction.auctionCounter() == 1, "Auction counter should be preserved");

        // Check specific auction data
        SimpleAuctionUpgradeable.Auction memory auctionData = auction.getAuction(auctionId);
        require(auctionData.nftContract == address(nft), "NFT contract should be preserved");
        require(auctionData.reservePrice == 0.1 ether, "Reserve price should be preserved");
        require(auctionData.seller == msg.sender, "Seller should be preserved");

        console.log("+ Data persistence verified");
    }

    function _testV2Functionality() internal {
        console.log("Testing V2 new functionality...");

        SimpleAuctionUpgradeableV2 auctionV2 = SimpleAuctionUpgradeableV2(address(auction));

        // Check initial emergency pause state
        bool initialPauseState = auctionV2.isEmergencyPaused();
        console.log("Initial emergency pause state:", initialPauseState);
        require(!initialPauseState, "Should not be paused initially");

        // Toggle emergency pause
        auctionV2.toggleEmergencyPause();
        bool newPauseState = auctionV2.isEmergencyPaused();
        console.log("Emergency pause state after toggle:", newPauseState);
        require(newPauseState, "Should be paused after toggle");

        // Try to create auction while paused (should fail)
        uint256 newTokenId = nft.mint(msg.sender);
        nft.setApprovalForAll(address(auction), true);

        try auctionV2.createAuction(address(nft), newTokenId, 3600, 0.1 ether) {
            revert("Should have failed due to emergency pause");
        } catch Error(string memory reason) {
            console.log("Expected failure:", reason);
            require(
                keccak256(abi.encodePacked(reason)) == keccak256(abi.encodePacked("Contract is emergency paused")),
                "Should fail with emergency pause message"
            );
        }

        // Toggle pause off and try again
        auctionV2.toggleEmergencyPause();
        require(!auctionV2.isEmergencyPaused(), "Should not be paused after second toggle");

        // Now creating auction should work
        uint256 newAuctionId = auctionV2.createAuction(address(nft), newTokenId, 3600, 0.1 ether);
        console.log("Successfully created auction after unpausing:", newAuctionId);

        console.log("+ V2 functionality verified");
    }
}