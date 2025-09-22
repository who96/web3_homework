// SPDX-License-Identifier: MIT
pragma solidity ^0.8.7;

import "forge-std/Test.sol";
import "forge-std/console.sol";
import "../../src/upgradeable/SimpleAuctionUpgradeable.sol";
import "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";
import "@openzeppelin/contracts/proxy/transparent/ProxyAdmin.sol";
import "../../src/AuctionPriceFeed.sol";
import "../../src/SimpleNFT.sol";

/**
 * @title BasicProxyTest
 * @notice Simple test to verify proxy architecture works
 */
contract BasicProxyTest is Test {
    SimpleNFT public nft;
    AuctionPriceFeed public priceFeed;
    SimpleAuctionUpgradeable public auctionImpl;
    ProxyAdmin public proxyAdmin;
    TransparentUpgradeableProxy public auctionProxy;
    SimpleAuctionUpgradeable public auction; // Proxy interface

    address public owner = address(0x1);
    address public platformWallet = address(0x4);

    // Mock price feed for testing
    address mockPriceFeed = address(0x5);

    function setUp() public {
        vm.startPrank(owner);

        // Deploy NFT
        nft = new SimpleNFT();

        // Deploy mock price feed for testing
        vm.mockCall(
            mockPriceFeed,
            abi.encodeWithSignature("latestRoundData()"),
            abi.encode(uint80(1), int256(400000000000), uint256(block.timestamp), uint256(block.timestamp), uint80(1))
        );
        vm.mockCall(
            mockPriceFeed,
            abi.encodeWithSignature("decimals()"),
            abi.encode(uint8(8))
        );

        priceFeed = new AuctionPriceFeed(mockPriceFeed);

        // Deploy implementation
        auctionImpl = new SimpleAuctionUpgradeable();

        // Deploy proxy admin
        proxyAdmin = new ProxyAdmin(owner);

        // Prepare initialization data
        bytes memory initData = abi.encodeWithSelector(
            SimpleAuctionUpgradeable.initialize.selector,
            address(priceFeed),
            platformWallet
        );

        // Deploy proxy
        auctionProxy = new TransparentUpgradeableProxy(
            address(auctionImpl),
            address(proxyAdmin),
            initData
        );

        // Get proxy interface
        auction = SimpleAuctionUpgradeable(address(auctionProxy));

        vm.stopPrank();
    }

    function testInitialDeployment() public view {
        // Verify initial state
        assertEq(auction.auctionCounter(), 0);
        assertEq(auction.PLATFORM_FEE(), 300);
        assertEq(auction.MIN_BID_INCREMENT(), 500);
        assertEq(auction.MIN_DURATION(), 120);
        assertEq(auction.owner(), owner);
        assertEq(auction.platformWallet(), platformWallet);
    }

    function testProxyAdminOwnership() public view {
        assertEq(proxyAdmin.owner(), owner);
    }

    function testBasicAuctionFunctionality() public {
        vm.startPrank(owner);

        // Mint NFT
        uint256 tokenId = nft.mint(owner);

        // Approve auction contract
        nft.setApprovalForAll(address(auction), true);

        // Create auction
        uint256 auctionId = auction.createAuction(
            address(nft),
            tokenId,
            3600, // 1 hour
            0.1 ether
        );

        assertEq(auction.auctionCounter(), 1);
        assertTrue(auction.isAuctionActive(auctionId));

        vm.stopPrank();
    }

    function testVersionFunction() public view {
        string memory version = auction.getVersion();
        assertEq(version, "1.0.0");
    }

    function testPlatformWalletSetter() public {
        address newPlatformWallet = address(0x999);

        vm.prank(owner);
        auction.setPlatformWallet(newPlatformWallet);

        assertEq(auction.platformWallet(), newPlatformWallet);
    }
}