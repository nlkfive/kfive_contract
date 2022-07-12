// SPDX-License-Identifier: GPL-3.0
pragma solidity 0.8.4;

import "./IAuction.sol";

interface IBlindAuction is IAuction {
    ////////////////////////////////////////////////////////////
    // EVENTS
    ////////////////////////////////////////////////////////////
    event BlindAuctionCreatedSuccessful(
        address assetOwner,
        address nftAddress,
        bytes32 blindAuctionId,
        uint256 assetId,
        uint256 biddingEnd,
        uint256 revealEnd,
        uint256 startPriceInWei
    );
    event BlindAuctionBidSuccessful(
        address sender, 
        bytes32 blindAuctionId, 
        bytes32 blindedBid
    );
    event RevealSuccessful(
        address sender, 
        bytes32 blindAuctionId
    );
    ////////////////////////////////////////////////////////////
    // STORAGE
    ////////////////////////////////////////////////////////////
    struct BlindAuction {
        // Blind Auction ID
        bytes32 id;
        // Owner of the NFT
        address seller;
        // Bidding end time
        uint256 biddingEnd;
        // Reveal end time
        uint256 revealEnd;
        // Start Price (in wei) for the published item
        uint256 startPrice;
        // Highest bid
        uint256 highestBid;
        // Highest bidder
        address highestBidder;
    }
}
