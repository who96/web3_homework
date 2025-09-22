// SPDX-License-Identifier: MIT
pragma solidity ^0.8.7;

import "forge-std/Test.sol";
import "forge-std/console.sol";
import "../src/SimpleNFT.sol";

contract SimpleNFTTest is Test {
    SimpleNFT public nft;
    address public owner;
    address public user1;
    address public user2;

    event Transfer(address indexed from, address indexed to, uint256 indexed tokenId);

    function setUp() public {
        owner = address(this);
        user1 = address(0x1);
        user2 = address(0x2);

        nft = new SimpleNFT();
    }

    function testInitialState() public {
        assertEq(nft.name(), "SimpleNFT");
        assertEq(nft.symbol(), "SNFT");
        assertEq(nft.totalSupply(), 0);
        assertEq(nft.tokenCounter(), 0);
        assertEq(nft.owner(), owner);
    }

    function testMintToUser() public {
        uint256 tokenId = nft.mint(user1);

        assertEq(tokenId, 0);
        assertEq(nft.ownerOf(tokenId), user1);
        assertEq(nft.balanceOf(user1), 1);
        assertEq(nft.totalSupply(), 1);
        assertEq(nft.tokenCounter(), 1);
    }

    function testMintMultiple() public {
        uint256 tokenId1 = nft.mint(user1);
        uint256 tokenId2 = nft.mint(user2);
        uint256 tokenId3 = nft.mint(user1);

        assertEq(tokenId1, 0);
        assertEq(tokenId2, 1);
        assertEq(tokenId3, 2);

        assertEq(nft.ownerOf(0), user1);
        assertEq(nft.ownerOf(1), user2);
        assertEq(nft.ownerOf(2), user1);

        assertEq(nft.balanceOf(user1), 2);
        assertEq(nft.balanceOf(user2), 1);
        assertEq(nft.totalSupply(), 3);
        assertEq(nft.tokenCounter(), 3);
    }

    function testMintEmitsTransferEvent() public {
        vm.expectEmit(true, true, true, true);
        emit Transfer(address(0), user1, 0);

        nft.mint(user1);
    }

    function testMintToZeroAddress() public {
        vm.expectRevert();
        nft.mint(address(0));
    }

    function testApproveAndTransfer() public {
        uint256 tokenId = nft.mint(user1);

        vm.prank(user1);
        nft.approve(user2, tokenId);

        assertEq(nft.getApproved(tokenId), user2);

        vm.prank(user2);
        nft.transferFrom(user1, user2, tokenId);

        assertEq(nft.ownerOf(tokenId), user2);
        assertEq(nft.balanceOf(user1), 0);
        assertEq(nft.balanceOf(user2), 1);
    }

    function testSetApprovalForAll() public {
        nft.mint(user1);
        uint256 tokenId2 = nft.mint(user1);

        vm.prank(user1);
        nft.setApprovalForAll(user2, true);

        assertTrue(nft.isApprovedForAll(user1, user2));

        vm.prank(user2);
        nft.transferFrom(user1, user2, 0);

        vm.prank(user2);
        nft.transferFrom(user1, user2, tokenId2);

        assertEq(nft.balanceOf(user1), 0);
        assertEq(nft.balanceOf(user2), 2);
    }

    function testUnauthorizedTransfer() public {
        uint256 tokenId = nft.mint(user1);

        vm.prank(user2);
        vm.expectRevert();
        nft.transferFrom(user1, user2, tokenId);
    }

    function testTransferToZeroAddress() public {
        uint256 tokenId = nft.mint(user1);

        vm.prank(user1);
        vm.expectRevert();
        nft.transferFrom(user1, address(0), tokenId);
    }

    function testTransferNonexistentToken() public {
        vm.expectRevert();
        nft.transferFrom(user1, user2, 999);
    }

    function testSupportsInterface() public {
        assertTrue(nft.supportsInterface(0x80ac58cd)); // ERC721
        assertTrue(nft.supportsInterface(0x5b5e139f)); // ERC721Metadata
        assertTrue(nft.supportsInterface(0x01ffc9a7)); // ERC165
        assertFalse(nft.supportsInterface(0x12345678)); // Random interface
    }

    function testOwnershipFunctions() public {
        assertEq(nft.owner(), owner);

        address newOwner = address(0x999);
        nft.transferOwnership(newOwner);

        assertEq(nft.owner(), newOwner);
    }

    function testTokenURINotImplemented() public {
        uint256 tokenId = nft.mint(user1);

        // Since we didn't implement tokenURI, it should return empty string
        // or revert depending on OpenZeppelin version
        try nft.tokenURI(tokenId) returns (string memory uri) {
            assertEq(uri, "");
        } catch {
            // Expected if tokenURI is not implemented
        }
    }

    function testFuzzMint(address to, uint8 count) public {
        vm.assume(to != address(0));
        vm.assume(to.code.length == 0); // Only EOA addresses
        vm.assume(count > 0 && count <= 100); // Reasonable bounds

        for (uint256 i = 0; i < count; i++) {
            uint256 tokenId = nft.mint(to);
            assertEq(tokenId, i);
            assertEq(nft.ownerOf(tokenId), to);
        }

        assertEq(nft.balanceOf(to), count);
        assertEq(nft.totalSupply(), count);
        assertEq(nft.tokenCounter(), count);
    }
}