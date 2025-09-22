// SPDX-License-Identifier: MIT
pragma solidity ^0.8.7;

import "forge-std/Test.sol";
import "forge-std/console.sol";
import "../../src/upgradeable/SimpleAuctionUpgradeable.sol";
import "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";
import "@openzeppelin/contracts/proxy/transparent/ProxyAdmin.sol";
import "@openzeppelin/contracts/interfaces/IERC1967.sol";
import "../../src/AuctionPriceFeed.sol";
import "../../src/SimpleNFT.sol";

contract ProxyUpgradeTest is Test {
    SimpleNFT public nft;
    AuctionPriceFeed public priceFeed;
    SimpleAuctionUpgradeable public auctionImpl;
    SimpleAuctionUpgradeable public auctionImplV2;
    ProxyAdmin public proxyAdmin;
    TransparentUpgradeableProxy public auctionProxy;
    SimpleAuctionUpgradeable public auction; // Proxy interface

    address public owner = address(0x1);
    address public user1 = address(0x2);
    address public user2 = address(0x3);
    address public platformWallet = address(0x4);

    // Mock price feed for testing
    address mockPriceFeed = address(0x5);

    function setUp() public {
        vm.startPrank(owner);

        // Deploy NFT and price feed
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

        // Prepare initialization data
        bytes memory initData = abi.encodeWithSelector(
            SimpleAuctionUpgradeable.initialize.selector,
            address(priceFeed),
            platformWallet
        );

        // Deploy proxy (TransparentUpgradeableProxy creates its own ProxyAdmin internally)
        auctionProxy = new TransparentUpgradeableProxy(
            address(auctionImpl),
            owner, // This is the initialOwner for the internally-created ProxyAdmin
            initData
        );

        // Get the ProxyAdmin that was created internally by the proxy
        // We need to use ERC1967 admin slot to get the actual admin
        bytes32 adminSlot = 0xb53127684a568b3173ae13b9f8a6016e243e63b6e8ee1178d6a717850b5d6103;
        proxyAdmin = ProxyAdmin(address(uint160(uint256(vm.load(address(auctionProxy), adminSlot)))));

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

    function testProxyStoragePersistence() public {
        vm.startPrank(owner);

        // Create some auction data
        uint256 tokenId = nft.mint(owner);
        nft.setApprovalForAll(address(auction), true);
        uint256 auctionId = auction.createAuction(
            address(nft),
            tokenId,
            3600,
            0.1 ether
        );

        // Store original data
        uint256 originalCounter = auction.auctionCounter();
        SimpleAuctionUpgradeable.Auction memory originalAuction = auction.getAuction(auctionId);

        vm.stopPrank();

        // Deploy new implementation (same contract for testing)
        auctionImplV2 = new SimpleAuctionUpgradeable();

        // Upgrade proxy
        vm.prank(owner);
        proxyAdmin.upgradeAndCall(
            ITransparentUpgradeableProxy(address(auctionProxy)),
            address(auctionImplV2),
            ""
        );

        // Verify data persistence after upgrade
        assertEq(auction.auctionCounter(), originalCounter);

        SimpleAuctionUpgradeable.Auction memory upgradedAuction = auction.getAuction(auctionId);
        assertEq(upgradedAuction.nftContract, originalAuction.nftContract);
        assertEq(upgradedAuction.tokenId, originalAuction.tokenId);
        assertEq(upgradedAuction.seller, originalAuction.seller);
        assertEq(upgradedAuction.reservePrice, originalAuction.reservePrice);
    }

    function testUpgradeAuthorization() public {
        // Deploy new implementation
        auctionImplV2 = new SimpleAuctionUpgradeable();

        // Try to upgrade from non-owner account (should fail)
        vm.expectRevert();
        vm.prank(user1);
        proxyAdmin.upgradeAndCall(
            ITransparentUpgradeableProxy(address(auctionProxy)),
            address(auctionImplV2),
            ""
        );

        // Upgrade from owner should succeed
        vm.prank(owner);
        proxyAdmin.upgradeAndCall(
            ITransparentUpgradeableProxy(address(auctionProxy)),
            address(auctionImplV2),
            ""
        );
    }

    function testVersionAfterUpgrade() public {
        // Check initial version
        string memory initialVersion = auction.getVersion();
        assertEq(initialVersion, "1.0.0");

        // Deploy upgraded implementation with modified version
        auctionImplV2 = new SimpleAuctionUpgradeableV2();

        // Upgrade
        vm.prank(owner);
        proxyAdmin.upgradeAndCall(
            ITransparentUpgradeableProxy(address(auctionProxy)),
            address(auctionImplV2),
            ""
        );

        // Check new version
        string memory newVersion = SimpleAuctionUpgradeableV2(address(auction)).getVersion();
        assertEq(newVersion, "2.0.0");
    }

    function testFunctionalityAfterUpgrade() public {
        vm.startPrank(owner);

        // Create auction before upgrade
        uint256 tokenId1 = nft.mint(owner);
        nft.setApprovalForAll(address(auction), true);
        uint256 auctionId1 = auction.createAuction(
            address(nft),
            tokenId1,
            3600,
            0.1 ether
        );

        // Upgrade
        auctionImplV2 = new SimpleAuctionUpgradeable();
        proxyAdmin.upgradeAndCall(
            ITransparentUpgradeableProxy(address(auctionProxy)),
            address(auctionImplV2),
            ""
        );

        // Verify old auction still works
        assertTrue(auction.isAuctionActive(auctionId1));

        // Create new auction after upgrade
        uint256 tokenId2 = nft.mint(owner);
        uint256 auctionId2 = auction.createAuction(
            address(nft),
            tokenId2,
            3600,
            0.1 ether
        );

        // Verify both auctions exist
        assertTrue(auction.isAuctionActive(auctionId1));
        assertTrue(auction.isAuctionActive(auctionId2));
        assertEq(auction.auctionCounter(), 2);

        vm.stopPrank();
    }

    function testProxyAdminTransferOwnership() public {
        address newOwner = address(0x999);

        vm.prank(owner);
        proxyAdmin.transferOwnership(newOwner);

        assertEq(proxyAdmin.owner(), newOwner);

        // Old owner cannot upgrade anymore
        auctionImplV2 = new SimpleAuctionUpgradeable();
        vm.expectRevert();
        vm.prank(owner);
        proxyAdmin.upgradeAndCall(
            ITransparentUpgradeableProxy(address(auctionProxy)),
            address(auctionImplV2),
            ""
        );

        // New owner can upgrade
        vm.prank(newOwner);
        proxyAdmin.upgradeAndCall(
            ITransparentUpgradeableProxy(address(auctionProxy)),
            address(auctionImplV2),
            ""
        );
    }
}

// Mock V2 contract for version testing
contract SimpleAuctionUpgradeableV2 is SimpleAuctionUpgradeable {
    function getVersion() public pure override returns (string memory) {
        return "2.0.0";
    }
}