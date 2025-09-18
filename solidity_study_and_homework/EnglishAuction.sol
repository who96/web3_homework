// SPDX-License-Identifier: MIT
pragma solidity ^0.8.10;

interface IERC721 {
    function transferFrom(address from, address to, uint256 tokenId) external;
}

contract EnglishAuction {
    IERC721 public immutable nft;
    uint256 public nftId;
    address payable public immutable seller;
    uint32 public endAt;
    bool public started;
    bool public ended;
    //拍卖持续时间
    uint256 public immutable duration;

    event Started();
    event Bid(address indexed bidder, uint256 amount);
    event Withdraw(address indexed bidder, uint256 amount);
    event Ended(address indexed bidder, uint256 amount);

    address public highestBidder;
    uint256 public highestBid;
    mapping(address => uint256) public bids;

    constructor(address _nft, uint256 _nftId, uint256 _startingPrice, uint256 _endAt, uint256 _duration) {
        nft = IERC721(_nft);
        nftId = _nftId;
        seller = payable(msg.sender);
        endAt = uint32(_endAt);
        duration = _duration;
        highestBid = _startingPrice;
    }

    function start() public {
        require(!started, "Auction already started");
        started = true;
        endAt = uint32(block.timestamp + duration);
        nft.transferFrom(seller, address(this), nftId);
        emit Started();
    }

    function bid() public payable {
        require(started, "Auction not started");
        require(block.timestamp < endAt, "Auction already ended");
        require(msg.value > highestBid, "There already is a higher bid");
        if (highestBid != 0) {
            bids[highestBidder] += highestBid;
        }
        highestBidder = msg.sender;
        highestBid = msg.value;
        emit Bid(msg.sender, msg.value);
    }

    function withdraw() public payable returns (bool) {
        uint256 amount = bids[msg.sender];
        if (amount > 0) {
            bids[msg.sender] = 0;   
            (bool success, ) = payable(msg.sender).call{value: amount}("");
            require(success, "Transfer failed");
            emit Withdraw(msg.sender, amount);
            return true;
        }
        return false;
    }

    function end() public {
        require(block.timestamp >= endAt, "Auction not ended");
        require(!ended, "Auction already ended");
        ended = true;
        if (highestBidder != address(0)) {
        nft.transferFrom(address(this), highestBidder, nftId);
            (bool success, ) = seller.call{value: highestBid}("");
            require(success, "Transfer failed");
        }else{
            nft.transferFrom(address(this), seller, nftId);
        }
        emit Ended(highestBidder, highestBid);
    }
}