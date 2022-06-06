// SPDX-License-Identifier: GPL-3.0
pragma solidity 0.8.4;

interface IAuction {
    // ERRORS
    error TooEarly(uint256 time);
    error TooLate(uint256 time);
    error NotRunning();

    // AUCTION EVENTS
    event BidSuccessful(
        address bidder,
        bytes32 auctionId,
        bytes32 blindedBid
    );
    event RevealSuccessful(bool fake, address revealer, bytes32 auctionId, uint256 value, bytes32 blindedBid);
    event RevealFailed(bool fake, address revealer, bytes32 auctionId, uint256 value);
    event AuctionEnded(bytes32 auctionId);
    event AuctionRefund(
        address bidder,
        bytes32 auctionId,
        uint256 deposit
    );
    event AuctionCreated(
        address seller,
        address nftAddress,
        bytes32 auctionId,
        uint256 assetId,
        uint256 biddingEnd,
        uint256 revealEnd,
        uint256 startPriceInWei
    );
    event AuctionSuccessful(
        address seller,
        address buyer,
        bytes32 auctionId,
        uint256 totalPrice
    );
    event AuctionCancelled(address who, bytes32 auctionId);

    ////////////////////////////////////////////////////////////
    // STORAGE
    ////////////////////////////////////////////////////////////
    struct Auction {
        // Owner of the NFT
        address seller;
        // Highest bidder
        address highestBidder;
        // Auction ID
        bytes32 id;
        // Bidding end time
        uint256 biddingEnd;
        // Reveal end time
        uint256 revealEnd;
        // Highest bid
        uint256 highestBid;
        // Start Price (in wei) for the published item
        uint256 startPrice;
    }
}
