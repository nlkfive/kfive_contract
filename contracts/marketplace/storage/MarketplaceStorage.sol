// SPDX-License-Identifier: MIT
pragma solidity 0.8.4;

import "./IAuction.sol";
import "./IOrder.sol";
import "./IMarketplaceStorage.sol";
import "../../common/AccessControl.sol";
import "@openzeppelin/contracts/token/ERC721/IERC721.sol";
import "@openzeppelin/contracts/utils/introspection/ERC165.sol";

contract MarketplaceStorage is
    KfiveAccessControl,
    IMarketplaceStorage
{
    address auctionMarketplace;
    address orderMarketplace;

    /**
     * @dev See {IERC165-supportsInterface}.
     */
    function supportsInterface(bytes4 interfaceId)
        public
        view
        virtual
        override(AccessControlEnumerable, IERC165)
        returns (bool)
    {
        return
            interfaceId == type(IMarketplaceStorage).interfaceId ||
            super.supportsInterface(interfaceId);
    }

    /**
     * @dev Update auction marketplace address.
     *
     * Requirements:
     *
     * - the caller must have the `ADMIN_ROLE`.
     */
    function updateAuctionMarketplace(address _auctionMarketplace)
        external
        whenNotPaused
        onlyRole(ADMIN_ROLE)
    {
        auctionMarketplace = _auctionMarketplace;
    }

    /**
     * @dev Update order marketplace address.
     *
     * Requirements:
     *
     * - the caller must have the `ADMIN_ROLE`.
     */
    function updateOrderMarketplace(address _orderMarketplace)
        external
        whenNotPaused
        onlyRole(ADMIN_ROLE)
    {
        orderMarketplace = _orderMarketplace;
    }

    ////////////////////////////////////////////////////////////
    // Availability
    ////////////////////////////////////////////////////////////
    function assetIsAvailable(bytes32 nftAsset)
        public
        view
        override
        returns (bool)
    {
        return
            runningActionIds[nftAsset] == bytes32(0) &&
            orderByNftAsset[nftAsset].orderId == bytes32(0);
    }

    ////////////////////////////////////////////////////////////
    // Auction storage
    ////////////////////////////////////////////////////////////
    // From running auction Id to auction
    mapping(bytes32 => Auction) private auctions;
    // From ERC721 registry keccak256(abi.encodePacked(nft,assetId)) to running Auction ID
    mapping(bytes32 => bytes32) private runningActionIds;

    function auctionIsExisted(bytes32 auctionId)
        external
        view
        override
        returns (bool)
    {
        return auctions[auctionId].id != bytes32(0);
    }

    function auctionIsEnded(bytes32 nftAsset, bytes32 auctionId)
    external
        view
        override
        returns (bool)
    {
        return runningActionIds[nftAsset] != auctionId;
    }

    function auctionIsRunning(bytes32 nftAsset, bytes32 auctionId)
        external
        view
        override
        returns (bool)
    {
        return runningActionIds[nftAsset] == auctionId;
    }

    /**
     * @dev Create new auction
     */
    function createAuction(
        address assetOwner,
        address nftAddress,
        uint256 assetId,
        bytes32 auctionId,
        uint256 biddingEnd,
        uint256 revealEnd,
        uint256 startPriceInWei
    ) 
        external 
        override 
        whenNotPaused 
        onlyFrom(auctionMarketplace) 
    {
        bytes32 nftAsset = keccak256(abi.encodePacked(nftAddress, assetId));
        if(!assetIsAvailable(nftAsset)) revert AssetUnvailable();

        // Update current running before actually create
        runningActionIds[nftAsset] = auctionId;

        // Create new auction
        Auction storage auction = auctions[auctionId];
        {
            auction.id = auctionId;
            auction.seller = assetOwner;
            auction.biddingEnd = biddingEnd;
            auction.revealEnd = revealEnd;
            auction.startPrice = startPriceInWei;
        }
    }

    /**
     * @dev Get auction by nftAsset.
     */
    function getAuction(bytes32 auctionId)
        external
        view
        override
        returns (Auction memory order)
    {
        return auctions[auctionId];
    }

    /**
     * @dev Delete auction by nftAsset.
     */
    function auctionEnded(bytes32 nftAsset)
        external
        override
        whenNotPaused
        onlyFrom(auctionMarketplace)
    {
        if (runningActionIds[nftAsset] == 0) revert AuctionAlreadyEnded();
        delete runningActionIds[nftAsset];
    }

    /**
     * @dev Update highest bid auction by auction ID.
     */
    function updateHighestBid(
        address bidder,
        uint256 highestBid,
        bytes32 auctionId
    ) 
        external 
        override 
        whenNotPaused
        onlyFrom(auctionMarketplace) 
    {
        Auction storage auction = auctions[auctionId];
        onlyBefore(auction.revealEnd);

        auction.highestBid = highestBid;
        auction.highestBidder = bidder;
    }

    ////////////////////////////////////////////////////////////
    // Trading storage
    ////////////////////////////////////////////////////////////
    // From ERC721 registry assetId to Order (to avoid asset collision)
    mapping(bytes32 => Order) public orderByNftAsset;

    function createOrder(
        address seller,
        address nftAddress,
        uint256 assetId,
        bytes32 orderId,
        uint256 priceInWei,
        uint256 expiredAt
    ) 
        external 
        override 
        onlyFrom(orderMarketplace) 
        whenNotPaused
    {
        bytes32 nftAsset = keccak256(abi.encodePacked(nftAddress, assetId));
        if(!assetIsAvailable(nftAsset)) revert AssetUnvailable();

        orderByNftAsset[nftAsset] = Order({
            orderId: orderId,
            seller: seller,
            nftAddress: nftAddress,
            price: priceInWei,
            expiredAt: expiredAt
        });
    }

    function getOrder(bytes32 nftAsset)
        external
        view
        override
        returns (Order memory order)
    {
        return orderByNftAsset[nftAsset];
    }

    function deleteOrder(bytes32 nftAsset)
        external
        override
        whenNotPaused
        onlyFrom(orderMarketplace)
    {
        if(orderByNftAsset[nftAsset].orderId == bytes32(0)) revert AssetNotExisted();
        delete orderByNftAsset[nftAsset];
    }

    function onlyBefore(uint256 _time) internal view {
        if (block.timestamp >= _time) revert TooLate(_time);
    }

    function onlyAfter(uint256 _time) internal view {
        if (block.timestamp <= _time) revert TooEarly(_time);
    }

    modifier onlyFrom(address contractAddr) {
        if(_msgSender() != contractAddr) revert InvalidMkpSender();
        _;
    }
}
