// SPDX-License-Identifier: MIT
pragma solidity ^0.8.7;

import "@openzeppelin/contracts/interfaces/IERC721.sol";
import "openzeppelin-contracts-upgradeable/contracts/utils/ReentrancyGuardUpgradeable.sol";
import "openzeppelin-contracts-upgradeable/contracts/access/OwnableUpgradeable.sol";
import "openzeppelin-contracts-upgradeable/contracts/proxy/utils/Initializable.sol";
import "../AuctionPriceFeed.sol";

contract SimpleAuctionUpgradeable is
    Initializable,
    ReentrancyGuardUpgradeable,
    OwnableUpgradeable
{
    struct Auction {
        address nftContract;
        uint256 tokenId;
        address seller;
        uint256 startTime;
        uint256 duration;
        uint256 reservePrice;
        address highestBidder;
        uint256 highestBid;
        bool ended;
        bool claimed;
    }

    mapping(uint256 => Auction) public auctions;
    mapping(address => uint256) public pendingRefunds;
    uint256 public auctionCounter;

    AuctionPriceFeed public PRICE_FEED;
    uint256 public constant PLATFORM_FEE = 300; // 3% = 300/10000
    uint256 public constant MIN_BID_INCREMENT = 500; // 5% = 500/10000
    uint256 public constant MIN_DURATION = 120; // 2 minutes
    address public platformWallet;

    event AuctionCreated(
        uint256 indexed auctionId,
        address indexed nftContract,
        uint256 indexed tokenId,
        address seller,
        uint256 startTime,
        uint256 duration,
        uint256 reservePrice
    );

    event BidPlaced(
        uint256 indexed auctionId,
        address indexed bidder,
        uint256 amount
    );

    event AuctionEnded(
        uint256 indexed auctionId,
        address indexed winner,
        uint256 amount
    );

    event RefundWithdrawn(address indexed bidder, uint256 amount);
    event NFTClaimed(uint256 indexed auctionId, address indexed winner);
    event FundsClaimed(uint256 indexed auctionId, address indexed seller, uint256 amount);

    error AuctionNotFound();
    error AuctionAlreadyEnded();
    error AuctionStillActive();
    error InsufficientBid();
    error NotAuthorized();
    error NotNFTOwner();
    error NFTNotApproved();
    error InvalidDuration();
    error NoRefundAvailable();
    error TransferFailed();
    error InvalidAuction();

    /// @custom:oz-upgrades-unsafe-allow constructor
    constructor() {
        _disableInitializers();
    }

    function initialize(
        address _priceFeed,
        address _platformWallet
    ) public initializer {
        __ReentrancyGuard_init();
        __Ownable_init(msg.sender);

        PRICE_FEED = AuctionPriceFeed(_priceFeed);
        platformWallet = _platformWallet;
        auctionCounter = 0;
    }

    function createAuction(
        address nftContract,
        uint256 tokenId,
        uint256 duration,
        uint256 reservePrice
    ) external virtual nonReentrant returns (uint256) {
        if (duration < MIN_DURATION) {
            revert InvalidDuration();
        }

        IERC721 nft = IERC721(nftContract);

        if (nft.ownerOf(tokenId) != msg.sender) {
            revert NotNFTOwner();
        }

        if (!nft.isApprovedForAll(msg.sender, address(this)) &&
            nft.getApproved(tokenId) != address(this)) {
            revert NFTNotApproved();
        }

        uint256 auctionId = auctionCounter;
        auctionCounter++;

        auctions[auctionId] = Auction({
            nftContract: nftContract,
            tokenId: tokenId,
            seller: msg.sender,
            startTime: block.timestamp,
            duration: duration,
            reservePrice: reservePrice,
            highestBidder: address(0),
            highestBid: 0,
            ended: false,
            claimed: false
        });

        nft.transferFrom(msg.sender, address(this), tokenId);

        emit AuctionCreated(
            auctionId,
            nftContract,
            tokenId,
            msg.sender,
            block.timestamp,
            duration,
            reservePrice
        );

        return auctionId;
    }

    function placeBid(uint256 auctionId) external payable virtual nonReentrant {
        Auction storage auction = auctions[auctionId];

        if (auction.seller == address(0)) {
            revert AuctionNotFound();
        }

        if (auction.ended || block.timestamp >= auction.startTime + auction.duration) {
            revert AuctionAlreadyEnded();
        }

        uint256 minBid = auction.highestBid == 0
            ? auction.reservePrice
            : auction.highestBid + (auction.highestBid * MIN_BID_INCREMENT / 10000);

        if (msg.value < minBid) {
            revert InsufficientBid();
        }

        if (auction.highestBidder != address(0)) {
            pendingRefunds[auction.highestBidder] += auction.highestBid;
        }

        auction.highestBidder = msg.sender;
        auction.highestBid = msg.value;

        emit BidPlaced(auctionId, msg.sender, msg.value);
    }

    function withdrawRefund() external nonReentrant {
        uint256 refund = pendingRefunds[msg.sender];
        if (refund == 0) {
            revert NoRefundAvailable();
        }

        pendingRefunds[msg.sender] = 0;

        (bool success, ) = payable(msg.sender).call{value: refund}("");
        if (!success) {
            pendingRefunds[msg.sender] = refund;
            revert TransferFailed();
        }

        emit RefundWithdrawn(msg.sender, refund);
    }

    function endAuction(uint256 auctionId) external nonReentrant {
        Auction storage auction = auctions[auctionId];

        if (auction.seller == address(0)) {
            revert AuctionNotFound();
        }

        if (auction.ended) {
            revert AuctionAlreadyEnded();
        }

        if (block.timestamp < auction.startTime + auction.duration) {
            revert AuctionStillActive();
        }

        auction.ended = true;

        emit AuctionEnded(auctionId, auction.highestBidder, auction.highestBid);
    }

    function claimNFT(uint256 auctionId) external nonReentrant {
        Auction storage auction = auctions[auctionId];

        if (auction.seller == address(0)) {
            revert AuctionNotFound();
        }

        if (!auction.ended) {
            revert AuctionStillActive();
        }

        if (auction.claimed) {
            revert InvalidAuction();
        }

        if (auction.highestBidder == address(0)) {
            if (msg.sender != auction.seller) {
                revert NotAuthorized();
            }
        } else {
            if (msg.sender != auction.highestBidder) {
                revert NotAuthorized();
            }
        }

        auction.claimed = true;

        IERC721(auction.nftContract).transferFrom(
            address(this),
            auction.highestBidder == address(0) ? auction.seller : auction.highestBidder,
            auction.tokenId
        );

        emit NFTClaimed(auctionId, msg.sender);
    }

    function claimFunds(uint256 auctionId) external nonReentrant {
        Auction storage auction = auctions[auctionId];

        if (auction.seller == address(0)) {
            revert AuctionNotFound();
        }

        if (!auction.ended) {
            revert AuctionStillActive();
        }

        if (msg.sender != auction.seller) {
            revert NotAuthorized();
        }

        if (auction.highestBidder == address(0)) {
            revert InvalidAuction();
        }

        uint256 totalAmount = auction.highestBid;
        uint256 platformFee = (totalAmount * PLATFORM_FEE) / 10000;
        uint256 sellerAmount = totalAmount - platformFee;

        auction.highestBid = 0;

        (bool successSeller, ) = payable(auction.seller).call{value: sellerAmount}("");
        if (!successSeller) {
            revert TransferFailed();
        }

        (bool successPlatform, ) = payable(platformWallet).call{value: platformFee}("");
        if (!successPlatform) {
            revert TransferFailed();
        }

        emit FundsClaimed(auctionId, auction.seller, sellerAmount);
    }

    function getAuction(uint256 auctionId) external view returns (Auction memory) {
        return auctions[auctionId];
    }

    function isAuctionActive(uint256 auctionId) external view returns (bool) {
        Auction memory auction = auctions[auctionId];
        return !auction.ended &&
               block.timestamp < auction.startTime + auction.duration &&
               auction.seller != address(0);
    }

    function getAuctionPriceInUSD(uint256 auctionId) external view returns (bool success, uint256 usdPrice) {
        Auction memory auction = auctions[auctionId];
        if (auction.highestBid == 0) {
            return PRICE_FEED.tryGetEthUsdPrice(auction.reservePrice);
        } else {
            return PRICE_FEED.tryGetEthUsdPrice(auction.highestBid);
        }
    }

    function setPlatformWallet(address _platformWallet) external onlyOwner {
        platformWallet = _platformWallet;
    }

    function getVersion() public pure virtual returns (string memory) {
        return "1.0.0";
    }

    uint256[44] private __gap;
}