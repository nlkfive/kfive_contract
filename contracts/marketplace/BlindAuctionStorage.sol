// SPDX-License-Identifier: GPL-3.0
pragma solidity 0.8.4;

contract BlindAuctionStorage {
    struct Bid {
        bytes32 blindedBid;
        uint256 deposit;
    }

    struct Auction {
        bool ended;
        // Auction ID
        bytes32 id;
        // Owner of the NFT
        address seller;
        // Highest bidder
        address highestBidder;
        // NFT registry address
        address nftAddress;
        // Bidding end time
        uint256 biddingEnd;
        // Reveal end time
        uint256 revealEnd;
        // Highest bid
        uint256 highestBid;
        // Start Price (in wei) for the published item
        uint256 startPrice;
        // List of bids
        mapping(address => Bid[]) bids;
        // List of pending returns
        mapping(address => uint256) pendingReturns;
    }

    // From ERC721 registry assetId to Auction (to avoid asset collision)
    mapping(address => mapping(uint256 => mapping(bytes32 => Auction)))
        public auctionByAssetId;

    // ERRORS
    error TooEarly(uint256 time);
    error TooLate(uint256 time);
    error AuctionEndAlreadyCalled();

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
        uint256 biddingEnd,
        uint256 revealEnd,
        uint256 startPriceInWei
    );
    event AuctionSuccessful(
        address indexed seller,
        address indexed buyer,
        bytes32 indexed auctionId,
        uint256 totalPrice
    );
    event AuctionCancelled(bytes32 indexed id);

    function onlyBefore(uint256 _time) internal view {
        if (block.timestamp >= _time) revert TooLate(_time);
    }

    function onlyAfter(uint256 _time) internal view {
        if (block.timestamp <= _time) revert TooEarly(_time);
    }
}
