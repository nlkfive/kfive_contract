// SPDX-License-Identifier: GPL-3.0
pragma solidity 0.8.4;

interface IAuction {
    // ERRORS
    error TooEarly(uint256 time);
    error TooLate(uint256 time);

    // AUCTION EVENTS
    event BidSuccessful(
        address indexed bidder,
        bytes32 indexed auctionId,
        bytes32 blindedBid
    );
    event RevealSuccessful(bool fake, bytes32 indexed auctionId, uint256 value);
    event AuctionRefund(
        address indexed bidder,
        bytes32 auctionId,
        uint256 deposit
    );
    event AuctionCreated(
        address indexed seller,
        address nftAddress,
        bytes32 indexed auctionId,
        uint256 assetId,
        uint256 biddingTime,
        uint256 revealTime,
        uint256 startPriceInWei
    );
    event AuctionSuccessful(
        address indexed seller,
        address indexed buyer,
        bytes32 indexed auctionId,
        uint256 totalPrice
    );
    event AuctionCancelled(bytes32 indexed id);

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
