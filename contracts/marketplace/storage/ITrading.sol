// SPDX-License-Identifier: GPL-3.0
pragma solidity 0.8.4;

interface ITrading {
    ////////////////////////////////////////////////////////////
    // ERRORS
    ////////////////////////////////////////////////////////////
    error InvalidExpiredTime();
    error TradingExpired();
    ////////////////////////////////////////////////////////////
    // EVENTS
    ////////////////////////////////////////////////////////////
    event TradingCreated(
        bytes32 tradingId,
        uint256 assetId,
        address seller,
        address nftAddress,
        uint256 priceInWei,
        uint256 expiredAt
    );
    event TradingSuccessful(bytes32 id, address buyer, address seller);
    event TradingCancelled(address who, bytes32 id);
    ////////////////////////////////////////////////////////////
    // STORAGE
    ////////////////////////////////////////////////////////////
    struct Trading {
        // Trading ID
        bytes32 tradingId;
        // Owner of the NFT
        address seller;
        // NFT registry address
        address nftAddress;
        // Price (in wei) for the published item
        uint256 price;
        // Time when this sale ends
        uint256 expiredAt;
    }
}
