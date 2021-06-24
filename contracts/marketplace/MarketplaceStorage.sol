// SPDX-License-Identifier: MIT
pragma solidity 0.8.4;

import "../BEP20/IBEP20.sol";
import "./BlindAuction.sol";
import "@openzeppelin/contracts/token/ERC721/IERC721.sol";

contract MarketplaceStorage is BlindAuction {
    IBEP20 public acceptedToken;

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

    uint256 public ownerCutPerMillion;
    uint256 public publicationFeeInWei;

    bytes4 public constant ERC721_Interface = bytes4(0x80ac58cd);

    // EVENTS
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

    event ChangedPublicationFee(uint256 publicationFee);
    event ChangedOwnerCutPerMillion(uint256 ownerCutPerMillion);
}
