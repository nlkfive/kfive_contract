// SPDX-License-Identifier: GPL-3.0
pragma solidity 0.8.4;

contract BlindAuctionStorage {
    struct Auction {
        // Owner of the NFT
        address seller;
        // Highest bidder
        address highestBidder;
        // NFT registry address
        address nftAddress;
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
        // List of bids (user => _blindedBid => deposit value)
        mapping(address => mapping(bytes32 => uint256)) bids;
        // List of pending returns
        mapping(address => uint256) pendingReturns;
    }

    struct NftAuction {
        bytes32 runningAuction;
        mapping(bytes32 => Auction) auctions;
    }

    // From ERC721 registry keccak256(abi.encodePacked(nft,assetId)) to NftAuction (to avoid asset collision)
    mapping(bytes32 => NftAuction) public nftAuctions;

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

    modifier auctionExisted(bytes32 nftAsset, bytes32 auctionId) {
        require(
            nftAuctions[nftAsset].auctions[auctionId].id != 0,
            "Auction is not existed"
        );
        _;
    }

    modifier auctionIsRunning(bytes32 nftAsset, bytes32 auctionId) {
        require(
            nftAuctions[nftAsset].runningAuction == auctionId,
            "Auction is not running"
        );
        _;
    }
}
