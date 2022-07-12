// SPDX-License-Identifier: GPL-3.0
pragma solidity 0.8.4;

interface IBlindAuction {
    // BLIND AUCTION EVENTS
    event BlindAuctionEnded(bytes32 blindAuctionId);
    event BlindAuctionRefund(
        address bidder,
        bytes32 blindAuctionId
    );
    event BlindAuctionCreated(
        address seller,
        address nftAddress,
        bytes32 blindAuctionId,
        uint256 assetId,
        uint256 biddingEnd,
        uint256 revealEnd,
        uint256 startPriceInWei
    );
    event RevealSuccessful(
        address seller,
        bytes32 blindAuctionId
    );
    event BlindAuctionSuccessful(
        address seller,
        address buyer,
        bytes32 blindAuctionId,
        uint256 totalPrice
    );
    event BlindAuctionBidSuccessful(
        address bidder,
        bytes32 blindAuctionId,
        bytes32 blindedBid
    );
    event BlindAuctionCancelled(address canceller, bytes32 blindAuctionId);

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
