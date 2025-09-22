// SPDX-License-Identifier: MIT
pragma solidity ^0.8.7;

import "./SimpleAuctionUpgradeable.sol";

/**
 * @title SimpleAuctionUpgradeableV2
 * @notice V2 version of the upgradeable auction contract for testing upgrade functionality
 * @dev This is a test version that adds a new feature while maintaining storage compatibility
 */
contract SimpleAuctionUpgradeableV2 is SimpleAuctionUpgradeable {

    // New feature in V2: Emergency pause functionality
    bool public emergencyPaused;

    // Event for V2 functionality
    event EmergencyPauseToggled(bool paused);

    /**
     * @dev Override version to indicate this is V2
     */
    function getVersion() public pure override returns (string memory) {
        return "2.0.0";
    }

    /**
     * @dev New function in V2: Toggle emergency pause
     * Only owner can call this function
     */
    function toggleEmergencyPause() external onlyOwner {
        emergencyPaused = !emergencyPaused;
        emit EmergencyPauseToggled(emergencyPaused);
    }

    /**
     * @dev Override createAuction to respect emergency pause
     */
    function createAuction(
        address nftContract,
        uint256 tokenId,
        uint256 duration,
        uint256 reservePrice
    ) external override nonReentrant returns (uint256) {
        require(!emergencyPaused, "Contract is emergency paused");

        // Implement the same logic as parent but with emergency pause check
        return _createAuctionInternal(nftContract, tokenId, duration, reservePrice);
    }

    /**
     * @dev Override placeBid to respect emergency pause
     */
    function placeBid(uint256 auctionId) external payable override nonReentrant {
        require(!emergencyPaused, "Contract is emergency paused");

        // Implement the same logic as parent but with emergency pause check
        _placeBidInternal(auctionId);
    }

    /**
     * @dev Internal function to create auction (extracted from parent logic)
     */
    function _createAuctionInternal(
        address nftContract,
        uint256 tokenId,
        uint256 duration,
        uint256 reservePrice
    ) internal returns (uint256) {
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

    /**
     * @dev Internal function to place bid (extracted from parent logic)
     */
    function _placeBidInternal(uint256 auctionId) internal {
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

    /**
     * @dev Get emergency pause status (new function in V2)
     */
    function isEmergencyPaused() external view returns (bool) {
        return emergencyPaused;
    }

    // Storage gap adjustment: reduce by 1 for the new emergencyPaused variable
    uint256[43] private __gap;
}