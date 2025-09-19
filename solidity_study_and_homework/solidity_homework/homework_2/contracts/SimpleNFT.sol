// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "@openzeppelin/contracts/token/ERC721/extensions/ERC721URIStorage.sol";
import "@openzeppelin/contracts/access/Ownable.sol";

contract SimpleNFT is ERC721URIStorage, Ownable {
    uint256 private _tokenIdCounter;

    constructor(string memory name, string memory symbol) ERC721(name, symbol) Ownable(msg.sender) {}

    function mintNFT(address recipient, string memory tokenURI) public returns (uint256) {
        uint256 tokenId = _tokenIdCounter++;
        _safeMint(recipient, tokenId);
        _setTokenURI(tokenId, tokenURI);
        return tokenId;
    }
}