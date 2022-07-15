// SPDX-License-Identifier: GPL-3.0
pragma solidity 0.8.4;

import "./IAuction.sol";

interface IPublicAuction is IAuction {
    ////////////////////////////////////////////////////////////
    // PUBLIC AUCTION ERRORS
    ////////////////////////////////////////////////////////////
    error NotWinner(); // 0x618c7242
    ////////////////////////////////////////////////////////////
    // PUBLIC AUCTION EVENTS
    ////////////////////////////////////////////////////////////
    event PublicAuctionCreatedSuccessful(
        address assetOwner,
        address nftAddress,
        bytes32 publicAuctionId,
        uint256 assetId,
        uint256 startTime,
        uint256 biddingEnd,
        uint256 startPriceInWei,
        uint256 minIncrement
    );
    event PublicAuctionBidSuccessful(
        address sender, 
        bytes32 publicAuctionId, 
        uint256 bidValue
    );
    ////////////////////////////////////////////////////////////
    // PUBLIC AUCTION STORAGE
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
        // Start time
        uint256 startTime;
        // Min increment bid price
        uint256 minIncrement;
    }
}
