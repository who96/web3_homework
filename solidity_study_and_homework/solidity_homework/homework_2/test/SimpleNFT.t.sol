// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "forge-std/Test.sol";
import "../contracts/SimpleNFT.sol";

contract SimpleNFTTest is Test {
    SimpleNFT public nft;
    address public testUser = address(0x123);
    address public owner;

    function setUp() public {
        owner = address(this);
        nft = new SimpleNFT("TestNFT", "TNFT");
    }

    function test_InitialState() public view {
        assertEq(nft.name(), "TestNFT", "Name should be TestNFT");
        assertEq(nft.symbol(), "TNFT", "Symbol should be TNFT");
        assertEq(nft.owner(), owner, "Owner should be test contract");
    }

    function test_MintNFT() public {
        string memory tokenURI = "ipfs://test-uri";
        uint256 tokenId = nft.mintNFT(testUser, tokenURI);

        assertEq(nft.ownerOf(tokenId), testUser, "Owner should be testUser");
        assertEq(nft.tokenURI(tokenId), tokenURI, "URI should match");
    }

    function test_MultipleMints() public {
        uint256 id1 = nft.mintNFT(testUser, "ipfs://1");
        uint256 id2 = nft.mintNFT(testUser, "ipfs://2");

        assertEq(id1, 0, "First token ID should be 0");
        assertEq(id2, 1, "Second token ID should be 1");
        assertEq(nft.balanceOf(testUser), 2, "Should own 2 NFTs");
    }

    function test_TransferOwnership() public {
        address newOwner = address(0x456);
        nft.transferOwnership(newOwner);
        assertEq(nft.owner(), newOwner, "Ownership should be transferred");
    }

    function testFuzz_MintMultipleNFTs(uint8 count) public {
        vm.assume(count > 0 && count < 100); // 合理的范围

        for (uint256 i = 0; i < count; i++) {
            nft.mintNFT(testUser, string(abi.encodePacked("ipfs://", vm.toString(i))));
        }

        assertEq(nft.balanceOf(testUser), count, "Balance should match mint count");
    }

    function test_RevertWhen_TokenDoesNotExist() public {
        vm.expectRevert();
        nft.ownerOf(999);
    }

    function test_Events() public {
        vm.expectEmit(true, true, true, true, address(nft));
        emit Transfer(address(0), testUser, 0);

        nft.mintNFT(testUser, "ipfs://test");
    }

    event Transfer(address indexed from, address indexed to, uint256 indexed tokenId);
}