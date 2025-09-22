// SPDX-License-Identifier: MIT
pragma solidity ^0.8.7;

import "forge-std/Script.sol";
import "forge-std/console.sol";
import "../../src/upgradeable/SimpleAuctionUpgradeable.sol";
import "../../src/proxy/AuctionProxy.sol";
import "../../src/proxy/AuctionProxyAdmin.sol";
import "../../src/AuctionPriceFeed.sol";
import "../../src/SimpleNFT.sol";
import "../HelperConfig.sol";

/**
 * @title DeployProxy
 * @notice Deploy upgradeable auction system with proxy architecture
 */
contract DeployProxy is Script {
    SimpleNFT public simpleNFT;
    AuctionPriceFeed public priceFeed;
    SimpleAuctionUpgradeable public auctionImpl;
    AuctionProxyAdmin public proxyAdmin;
    AuctionProxy public auctionProxy;
    HelperConfig public helperConfig;

    function run() external returns (
        SimpleNFT,
        AuctionPriceFeed,
        SimpleAuctionUpgradeable,
        AuctionProxyAdmin,
        AuctionProxy
    ) {
        helperConfig = new HelperConfig();
        (
            , // oracle
            , // jobId
            , // chainlinkFee
            , // link
            , // updateInterval
            address chainlinkPriceFeed,
            , // subscriptionId
            , // vrfCoordinator
              // keyHash
        ) = helperConfig.activeNetworkConfig();

        address platformWallet = msg.sender;

        vm.startBroadcast();

        // 1. Deploy SimpleNFT (non-upgradeable)
        simpleNFT = new SimpleNFT();
        console.log("SimpleNFT deployed at:", address(simpleNFT));

        // 2. Deploy AuctionPriceFeed (non-upgradeable)
        priceFeed = new AuctionPriceFeed(chainlinkPriceFeed);
        console.log("AuctionPriceFeed deployed at:", address(priceFeed));

        // 3. Deploy SimpleAuctionUpgradeable implementation
        auctionImpl = new SimpleAuctionUpgradeable();
        console.log("SimpleAuctionUpgradeable implementation deployed at:", address(auctionImpl));

        // 4. Deploy ProxyAdmin
        proxyAdmin = new AuctionProxyAdmin(msg.sender);
        console.log("AuctionProxyAdmin deployed at:", address(proxyAdmin));

        // 5. Prepare initialization data
        bytes memory initData = abi.encodeWithSelector(
            SimpleAuctionUpgradeable.initialize.selector,
            address(priceFeed),
            platformWallet
        );

        // 6. Deploy proxy pointing to implementation
        auctionProxy = new AuctionProxy(
            address(auctionImpl),
            address(proxyAdmin),
            initData
        );
        console.log("AuctionProxy deployed at:", address(auctionProxy));

        vm.stopBroadcast();

        // Verify deployment
        _verifyProxyDeployment();

        return (simpleNFT, priceFeed, auctionImpl, proxyAdmin, auctionProxy);
    }

    function _verifyProxyDeployment() internal view {
        console.log("\n=== Proxy Deployment Verification ===");

        // Verify SimpleNFT
        require(address(simpleNFT) != address(0), "SimpleNFT deployment failed");
        console.log("+ SimpleNFT verification passed");

        // Verify AuctionPriceFeed
        require(address(priceFeed) != address(0), "AuctionPriceFeed deployment failed");
        console.log("+ AuctionPriceFeed verification passed");

        // Verify implementation
        require(address(auctionImpl) != address(0), "Implementation deployment failed");
        console.log("+ Implementation verification passed");

        // Verify ProxyAdmin
        require(address(proxyAdmin) != address(0), "ProxyAdmin deployment failed");
        require(proxyAdmin.owner() == msg.sender, "ProxyAdmin owner should be deployer");
        console.log("+ ProxyAdmin verification passed");

        // Verify Proxy
        require(address(auctionProxy) != address(0), "Proxy deployment failed");
        console.log("+ Proxy verification passed");

        // Verify proxy functionality through interface
        SimpleAuctionUpgradeable auction = SimpleAuctionUpgradeable(address(auctionProxy));
        require(auction.auctionCounter() == 0, "Auction counter should be 0");
        require(auction.PLATFORM_FEE() == 300, "Platform fee should be 3%");
        require(auction.MIN_BID_INCREMENT() == 500, "Min bid increment should be 5%");
        require(auction.MIN_DURATION() == 120, "Min duration should be 2 minutes");
        console.log("+ Proxy functionality verification passed");

        // Verify version
        string memory version = auction.getVersion();
        console.log("+ Contract version:", version);

        console.log("SUCCESS: All proxy contracts deployed and verified!");
    }

    /**
     * @notice Helper function to upgrade the proxy to a new implementation
     * @param newImplementation Address of the new implementation contract
     */
    function upgradeProxy(address newImplementation) external {
        vm.startBroadcast();

        proxyAdmin.upgradeProxy(
            address(auctionProxy),
            newImplementation
        );

        vm.stopBroadcast();

        console.log("Proxy upgraded to new implementation:", newImplementation);
    }

    /**
     * @notice Helper function to upgrade with initialization call
     * @param newImplementation Address of the new implementation contract
     * @param data Initialization data to call on new implementation
     */
    function upgradeProxyAndCall(address newImplementation, bytes calldata data) external {
        vm.startBroadcast();

        proxyAdmin.upgradeProxyAndCall(
            address(auctionProxy),
            newImplementation,
            data
        );

        vm.stopBroadcast();

        console.log("Proxy upgraded with call to new implementation:", newImplementation);
    }
}