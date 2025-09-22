// SPDX-License-Identifier: MIT
pragma solidity ^0.8.7;

import "@chainlink/contracts/src/v0.8/shared/interfaces/AggregatorV3Interface.sol";

contract AuctionPriceFeed {
    AggregatorV3Interface internal immutable PRICE_FEED;
    uint8 public constant DECIMALS = 8;
    uint256 public constant STALE_PRICE_DELAY = 3600; // 1 hour

    error StalePriceData();
    error InvalidPriceData();

    constructor(address _priceFeed) {
        PRICE_FEED = AggregatorV3Interface(_priceFeed);
    }

    function getLatestPrice() public view returns (int256) {
        (
            uint80 roundId,
            int256 price,
            ,
            uint256 updatedAt,
            uint80 answeredInRound
        ) = PRICE_FEED.latestRoundData();

        if (price <= 0) {
            revert InvalidPriceData();
        }

        if (updatedAt == 0 || block.timestamp - updatedAt > STALE_PRICE_DELAY) {
            revert StalePriceData();
        }

        if (roundId != answeredInRound) {
            revert InvalidPriceData();
        }

        return price;
    }

    function getEthUsdPrice(uint256 ethAmount) external view returns (uint256) {
        int256 price = getLatestPrice();
        return (ethAmount * uint256(price)) / (10**18);
    }

    function tryGetLatestPrice() external view returns (bool success, int256 price) {
        try this.getLatestPrice() returns (int256 result) {
            return (true, result);
        } catch {
            return (false, 0);
        }
    }

    function tryGetEthUsdPrice(uint256 ethAmount) external view returns (bool success, uint256 usdAmount) {
        try this.getEthUsdPrice(ethAmount) returns (uint256 result) {
            return (true, result);
        } catch {
            return (false, 0);
        }
    }

    function getPriceFeed() external view returns (AggregatorV3Interface) {
        return PRICE_FEED;
    }

    function getDecimals() external view returns (uint8) {
        return PRICE_FEED.decimals();
    }
}