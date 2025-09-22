// SPDX-License-Identifier: MIT
pragma solidity ^0.8.7;

import "forge-std/Script.sol";
import "forge-std/console.sol";
import "../../src/upgradeable/SimpleAuctionUpgradeable.sol";
import "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";
import "@openzeppelin/contracts/proxy/transparent/ProxyAdmin.sol";
import "../../src/AuctionPriceFeed.sol";
import "../../src/SimpleNFT.sol";
import "../HelperConfig.sol";

/**
 * @title DeployProxySimple
 * @notice Deploy upgradeable auction system with simplified OpenZeppelin standard architecture
 */
contract DeployProxySimple is Script {
    SimpleNFT public simpleNFT;
    AuctionPriceFeed public priceFeed;
    SimpleAuctionUpgradeable public auctionImpl;
    TransparentUpgradeableProxy public auctionProxy;
    ProxyAdmin public proxyAdmin;
    HelperConfig public helperConfig;

    function run() external returns (
        SimpleNFT,
        AuctionPriceFeed,
        SimpleAuctionUpgradeable,
        TransparentUpgradeableProxy,
        ProxyAdmin
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

        // 4. Prepare initialization data
        bytes memory initData = abi.encodeWithSelector(
            SimpleAuctionUpgradeable.initialize.selector,
            address(priceFeed),
            platformWallet
        );

        // 5. Deploy TransparentUpgradeableProxy (creates internal ProxyAdmin automatically)
        auctionProxy = new TransparentUpgradeableProxy(
            address(auctionImpl),
            msg.sender, // This becomes the owner of the internal ProxyAdmin
            initData
        );
        console.log("TransparentUpgradeableProxy deployed at:", address(auctionProxy));

        // 6. Get the internal ProxyAdmin address
        bytes32 adminSlot = 0xb53127684a568b3173ae13b9f8a6016e243e63b6e8ee1178d6a717850b5d6103;
        address internalProxyAdminAddr = address(uint160(uint256(vm.load(address(auctionProxy), adminSlot))));
        proxyAdmin = ProxyAdmin(internalProxyAdminAddr);
        console.log("Internal ProxyAdmin address:", address(proxyAdmin));

        vm.stopBroadcast();

        // Verify deployment
        _verifyProxyDeployment();

        return (simpleNFT, priceFeed, auctionImpl, auctionProxy, proxyAdmin);
    }

    function _verifyProxyDeployment() internal view {
        console.log("\\n=== Simplified Proxy Deployment Verification ===");

        // Verify SimpleNFT
        require(address(simpleNFT) != address(0), "SimpleNFT deployment failed");
        console.log("+ SimpleNFT verification passed");

        // Verify AuctionPriceFeed
        require(address(priceFeed) != address(0), "AuctionPriceFeed deployment failed");
        console.log("+ AuctionPriceFeed verification passed");

        // Verify implementation
        require(address(auctionImpl) != address(0), "Implementation deployment failed");
        console.log("+ Implementation verification passed");

        // Verify Proxy
        require(address(auctionProxy) != address(0), "Proxy deployment failed");
        console.log("+ Proxy verification passed");

        // Verify ProxyAdmin
        require(address(proxyAdmin) != address(0), "ProxyAdmin not found");
        require(proxyAdmin.owner() == msg.sender, "ProxyAdmin owner should be deployer");
        console.log("+ ProxyAdmin verification passed");

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

        console.log("SUCCESS: All contracts deployed and verified with simplified architecture!");
    }

    /**
     * @notice Helper function to upgrade the proxy to a new implementation
     * @param newImplementation Address of the new implementation contract
     */
    function upgradeProxy(address newImplementation) external {
        vm.startBroadcast();

        proxyAdmin.upgradeAndCall(
            ITransparentUpgradeableProxy(address(auctionProxy)),
            newImplementation,
            ""
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

        proxyAdmin.upgradeAndCall(
            ITransparentUpgradeableProxy(address(auctionProxy)),
            newImplementation,
            data
        );

        vm.stopBroadcast();

        console.log("Proxy upgraded with call to new implementation:", newImplementation);
    }
}