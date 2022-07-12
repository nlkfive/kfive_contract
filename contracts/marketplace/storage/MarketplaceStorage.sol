// SPDX-License-Identifier: MIT
pragma solidity 0.8.4;

import "./IPublicAuction.sol";
import "./IOrder.sol";
import "./IMarketplaceStorage.sol";
import "../../common/AccessControl.sol";
import "@openzeppelin/contracts/token/ERC721/IERC721.sol";
import "@openzeppelin/contracts/utils/introspection/ERC165.sol";

contract MarketplaceStorage is
    KfiveAccessControl,
    IMarketplaceStorage
{
    address orderMarketplace;
    address publicAuctionMarketplace;
    address blindAuctionMarketplace;

    error TooEarly(uint256);
    error TooLate(uint256);

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
     * @dev Update public auction marketplace address.
     *
     * Requirements:
     *
     * - the caller must have the `ADMIN_ROLE`.
     */
    function updatePublicAuctionMarketplace(address _publicAuctionMarketplace)
        external
        whenNotPaused
        onlyRole(ADMIN_ROLE)
    {
        publicAuctionMarketplace = _publicAuctionMarketplace;
    }

     /**
     * @dev Update blind auction marketplace address.
     *
     * Requirements:
     *
     * - the caller must have the `ADMIN_ROLE`.
     */
    function updateBlindAuctionMarketplace(address _blindAuctionMarketplace)
        external
        whenNotPaused
        onlyRole(ADMIN_ROLE)
    {
        blindAuctionMarketplace = _blindAuctionMarketplace;
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
            runningPublicAuctionIds[nftAsset] == bytes32(0) &&
            runningBlindAuctionIds[nftAsset] == bytes32(0) &&
            orderByNftAsset[nftAsset].orderId == bytes32(0);
    }

    ////////////////////////////////////////////////////////////
    // Public Auction storage
    ////////////////////////////////////////////////////////////
    // From running Id to public auction
    mapping(bytes32 => PublicAuction) private publicAuctions;
    // From ERC721 registry keccak256(abi.encodePacked(nft,assetId)) to running Public Auction ID
    mapping(bytes32 => bytes32) private runningPublicAuctionIds;

    function publicAuctionIsExisted(bytes32 publicAuctionId)
        external
        view
        override
        returns (bool)
    {
        return publicAuctions[publicAuctionId].id != bytes32(0);
    }

    function publicAuctionIsEnded(bytes32 nftAsset, bytes32 publicAuctionId)
    external
        view
        override
        returns (bool)
    {
        return runningPublicAuctionIds[nftAsset] != publicAuctionId;
    }

    function publicAuctionIsRunning(bytes32 nftAsset, bytes32 publicAuctionId)
        external
        view
        override
        returns (bool)
    {
        return runningPublicAuctionIds[nftAsset] == publicAuctionId;
    }

    /**
     * @dev Create new public auction
     */
    function createPublicAuction(
        address assetOwner,
        address nftAddress,
        uint256 assetId,
        bytes32 publicAuctionId,
        uint256 biddingEnd,
        uint256 startPriceInWei,
        uint256 minIncrement
    ) 
        external 
        override 
        whenNotPaused 
        onlyFrom(publicAuctionMarketplace)
    {
        bytes32 nftAsset = keccak256(abi.encodePacked(nftAddress, assetId));
        if(!assetIsAvailable(nftAsset)) revert AssetUnvailable();

        // Update current running before actually create
        runningPublicAuctionIds[nftAsset] = publicAuctionId;

        // Create new public auction
        PublicAuction storage publicAuction = publicAuctions[publicAuctionId];
        {
            publicAuction.id = publicAuctionId;
            publicAuction.seller = assetOwner;
            publicAuction.biddingEnd = biddingEnd;
            publicAuction.startPrice = startPriceInWei;
            publicAuction.minIncrement = minIncrement;
        }
    }

    /**
     * @dev Get public auction by id.
     */
    function getPublicAuction(bytes32 publicAuctionId)
        external
        view
        override
        returns (PublicAuction memory publicAuction)
    {
        return publicAuctions[publicAuctionId];
    }

    /**
     * @dev Delete public auction by nftAsset.
     */
    function endPublicAuction(bytes32 nftAsset)
        external
        override
        whenNotPaused
        onlyFrom(publicAuctionMarketplace)
    {
        if (runningPublicAuctionIds[nftAsset] == 0) revert AuctionAlreadyEnded();
        delete runningPublicAuctionIds[nftAsset];
    }

    /**
     * @dev Update highest bid auction by public auction ID.
     */
    function updateHighestBidPublicAuction(
        address bidder,
        uint256 highestBid,
        bytes32 publicAuctionId
    ) 
        external 
        override 
        whenNotPaused
        onlyFrom(publicAuctionMarketplace)
    {
        PublicAuction storage publicAuction = publicAuctions[publicAuctionId];
        onlyBefore(publicAuction.biddingEnd);

        publicAuction.highestBid = highestBid;
        publicAuction.highestBidder = bidder;
    }

    ////////////////////////////////////////////////////////////
    // Blind Auction storage
    ////////////////////////////////////////////////////////////
    // From running Id to blind auction
    mapping(bytes32 => BlindAuction) private blindAuctions;
    // From ERC721 registry keccak256(abi.encodePacked(nft,assetId)) to running Blind Auction ID
    mapping(bytes32 => bytes32) private runningBlindAuctionIds;

    function blindAuctionIsExisted(bytes32 blindAuctionId)
        external
        view
        override
        returns (bool)
    {
        return blindAuctions[blindAuctionId].id != bytes32(0);
    }

    function blindAuctionIsEnded(bytes32 nftAsset, bytes32 blindAuctionId)
    external
        view
        override
        returns (bool)
    {
        return runningBlindAuctionIds[nftAsset] != blindAuctionId;
    }

    function blindAuctionIsRunning(bytes32 nftAsset, bytes32 blindAuctionId)
        external
        view
        override
        returns (bool)
    {
        return runningBlindAuctionIds[nftAsset] == blindAuctionId;
    }

    /**
     * @dev Create new blind auction
     */
    function createBlindAuction(
        address assetOwner,
        address nftAddress,
        uint256 assetId,
        bytes32 blindAuctionId,
        uint256 biddingEnd,
        uint256 revealEnd,
        uint256 startPriceInWei
    ) 
        external 
        override 
        whenNotPaused 
        onlyFrom(blindAuctionMarketplace)
    {
        bytes32 nftAsset = keccak256(abi.encodePacked(nftAddress, assetId));
        if(!assetIsAvailable(nftAsset)) revert AssetUnvailable();

        // Update current running before actually create
        runningBlindAuctionIds[nftAsset] = blindAuctionId;

        // Create new blind auction
        BlindAuction storage blindAuction = blindAuctions[blindAuctionId];
        {
            blindAuction.id = blindAuctionId;
            blindAuction.seller = assetOwner;
            blindAuction.biddingEnd = biddingEnd;
            blindAuction.revealEnd = revealEnd;
            blindAuction.startPrice = startPriceInWei;
        }
    }

    /**
     * @dev Get blind auction by id.
     */
    function getBlindAuction(bytes32 blindAuctionId)
        external
        view
        override
        returns (BlindAuction memory blindAuction)
    {
        return blindAuctions[blindAuctionId];
    }

    /**
     * @dev Delete blind auction by nftAsset.
     */
    function endBlindAuction(bytes32 nftAsset)
        external
        override
        whenNotPaused
        onlyFrom(blindAuctionMarketplace)
    {
        if (runningBlindAuctionIds[nftAsset] == 0) revert AuctionAlreadyEnded();
        delete runningBlindAuctionIds[nftAsset];
    }

    /**
     * @dev Update highest bid auction by public auction ID.
     */
    function updateHighestBidBlindAuction(
        address bidder,
        uint256 highestBid,
        bytes32 blindAuctionId
    ) 
        external 
        override 
        whenNotPaused
        onlyFrom(publicAuctionMarketplace)
    {
        BlindAuction storage blindAuction = blindAuctions[blindAuctionId];
        onlyBefore(blindAuction.biddingEnd);

        blindAuction.highestBid = highestBid;
        blindAuction.highestBidder = bidder;
    }

    ////////////////////////////////////////////////////////////
    // Order storage
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
