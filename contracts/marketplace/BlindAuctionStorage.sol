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
        address nftAddress,
        uint256 indexed assetId,
        bytes32 blindedBid
    );

    event AuctionRefund(
        address nftAddress,
        uint256 indexed assetId,
        uint256 deposit,
        address indexed bidder
    );
    event AuctionCreated(
        bytes32 id,
        uint256 indexed assetId,
        address indexed seller,
        address nftAddress,
        uint256 biddingEnd,
        uint256 revealEnd,
        uint256 startPriceInWei
    );
    event AuctionSuccessful(
        bytes32 id,
        uint256 indexed assetId,
        address indexed seller,
        address nftAddress,
        uint256 totalPrice,
        address indexed buyer
    );
    event AuctionCancelled(
        bytes32 id,
        uint256 indexed assetId,
        address indexed seller,
        address nftAddress
    );

    function onlyBefore(uint256 _time) internal view {
        if (block.timestamp >= _time) revert TooLate(_time);
    }

    function onlyAfter(uint256 _time) internal view {
        if (block.timestamp <= _time) revert TooEarly(_time);
    }
}
