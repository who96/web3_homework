// SPDX-License-Identifier: MIT
pragma solidity ^0.8.10;

interface IERC721 {
    function transferFrom(address from, address to, uint256 tokenId) external;
}

contract DutchAuction {
    IERC721 public immutable nft;
    uint256 public nftId;
    uint private constant DURATION = 7 days;
    address payable public immutable seller;
    uint256 public immutable startingPrice;
    uint256 public immutable startedAt;
    uint256 public immutable expiresAt;
    uint256 public immutable discountRate;

    constructor(address _nft, uint256 _nftId, uint256 _startingPrice, uint256 _discountRate) {
        nft = IERC721(_nft);
        nftId = _nftId;
        seller = payable(msg.sender);
        startingPrice = _startingPrice;
        discountRate = _discountRate;
        startedAt = block.timestamp;
        expiresAt = block.timestamp + DURATION;

        require(_startingPrice >= _discountRate * DURATION, "Starting price is too low");
    }

    function getPrice() public view returns (uint256) {
        uint256 timeElapsed = block.timestamp - startedAt;
        uint256 discount = discountRate * timeElapsed;
        return startingPrice - discount;
    }

    function buy() public payable {
        require(block.timestamp < expiresAt, "Auction already ended");
        uint256 price = getPrice();
        require(msg.value >= price, "Insufficient bid amount");
        nft.transferFrom(seller, msg.sender, nftId);
        uint256 refund = msg.value - price;
        if (refund > 0) {
            payable(msg.sender).transfer(refund);
        }
        //发送ether给卖家
        (bool success, ) = seller.call{value: price}("");
        require(success, "Transfer failed");
        
    }
}