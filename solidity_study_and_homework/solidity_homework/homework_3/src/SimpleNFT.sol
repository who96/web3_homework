// SPDX-License-Identifier: MIT
pragma solidity ^0.8.7;

import "@openzeppelin/contracts/token/ERC721/ERC721.sol";
import "@openzeppelin/contracts/access/Ownable.sol";

contract SimpleNFT is ERC721, Ownable {
    uint256 private _tokenCounter;

    constructor() ERC721("SimpleNFT", "SNFT") Ownable(msg.sender) {
        _tokenCounter = 0;
    }

    function mint(address to) external returns (uint256) {
        uint256 tokenId = _tokenCounter;
        _tokenCounter += 1;
        _safeMint(to, tokenId);
        return tokenId;
    }

    function tokenCounter() external view returns (uint256) {
        return _tokenCounter;
    }

    function totalSupply() external view returns (uint256) {
        return _tokenCounter;
    }
}