// SPDX-License-Identifier: GPL-3.0
pragma solidity 0.8.4;

interface IOrder {
    ////////////////////////////////////////////////////////////
    // EVENTS
    ////////////////////////////////////////////////////////////
    event OrderCreated(
        bytes32 orderId,
        uint256 assetId,
        address seller,
        address nftAddress,
        uint256 priceInWei,
        uint256 expiredAt
    );
    event OrderSuccessful(bytes32 id, address buyer, address seller);
    event OrderCancelled(address who, bytes32 id);

    ////////////////////////////////////////////////////////////
    // STORAGE
    ////////////////////////////////////////////////////////////
    struct Order {
        // Order ID
        bytes32 orderId;
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
