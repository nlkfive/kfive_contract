// SPDX-License-Identifier: GPL-3.0
pragma solidity 0.8.4;

interface IPublicAuction {
    // PUBLIC AUCTION EVENTS
    event PublicAuctionEnded(bytes32 publicAuctionId);
    event PublicAuctionRefund(
        address bidder,
        bytes32 publicAuctionId,
        uint256 value
    );
    event PublicAuctionCreated(
        address seller,
        address nftAddress,
        bytes32 publicAuctionId,
        uint256 assetId,
        uint256 biddingEnd,
        uint256 startPriceInWei,
        uint256 minIncrement
    );
    event PublicAuctionSuccessful(
        address seller,
        address buyer,
        bytes32 publicAuctionId,
        uint256 totalPrice
    );
    event PublicAuctionBidSuccessful(
        address bidder,
        bytes32 publicAuctionId,
        uint256 bidValue
    );
    event PublicAuctionCancelled(address canceller, bytes32 publicAuctionId);

    ////////////////////////////////////////////////////////////
    // STORAGE
    ////////////////////////////////////////////////////////////
    struct PublicAuction {
        // Public Auction ID
        bytes32 id;
        // Owner of the NFT
        address seller;
        // Highest bid
        uint256 highestBid;
        // Highest bidder
        address highestBidder;
        // Bidding end time
        uint256 biddingEnd;
        // Start Price (in wei) for the published item
        uint256 startPrice;
        // Min increment bid price
        uint256 minIncrement;
    }
}
