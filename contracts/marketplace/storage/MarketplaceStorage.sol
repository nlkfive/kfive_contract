// SPDX-License-Identifier: MIT
pragma solidity 0.8.4;

import "./IPublicAuction.sol";
import "./ITrading.sol";
import "./IMarketplaceStorage.sol";
import "../../common/AccessControl.sol";
import "@openzeppelin/contracts/token/ERC721/IERC721.sol";
import "@openzeppelin/contracts/utils/introspection/ERC165.sol";

contract MarketplaceStorage is
    KfiveAccessControl,
    IMarketplaceStorage
{
    address tradingMarketplace;
    address publicAuctionMarketplace;
    address blindAuctionMarketplace;

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
     * @dev Update trading marketplace address.
     *
     * Requirements:
     *
     * - the caller must have the `ADMIN_ROLE`.
     */
    function updateTradingMarketplace(address _tradingMarketplace)
        external
        whenNotPaused
        onlyRole(ADMIN_ROLE)
    {
        tradingMarketplace = _tradingMarketplace;
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
            tradingByNftAsset[nftAsset].tradingId == bytes32(0);
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
        bytes32 nftAsset,
        bytes32 publicAuctionId,
        uint256 biddingEnd,
        uint256 startPriceInWei,
        uint256 startTime,
        uint256 minIncrement
    ) 
        external 
        override 
        whenNotPaused 
        onlyFrom(publicAuctionMarketplace)
    {
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
            publicAuction.startTime = startTime;
            publicAuction.minIncrement = minIncrement;
            publicAuction.highestBid = startPriceInWei;
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
        if (runningPublicAuctionIds[nftAsset] == 0) revert AuctionEnded();
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
        bytes32 nftAsset,
        bytes32 blindAuctionId,
        uint256 startTime,
        uint256 biddingEnd,
        uint256 revealEnd,
        uint256 startPriceInWei
    ) 
        external 
        override 
        whenNotPaused 
        onlyFrom(blindAuctionMarketplace)
    {
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
            blindAuction.startTime = startTime;
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
        if (runningBlindAuctionIds[nftAsset] == 0) revert AuctionEnded();
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
        onlyFrom(blindAuctionMarketplace)
    {
        BlindAuction storage blindAuction = blindAuctions[blindAuctionId];
        onlyAfter(blindAuction.biddingEnd);

        blindAuction.highestBid = highestBid;
        blindAuction.highestBidder = bidder;
    }

    ////////////////////////////////////////////////////////////
    // Trading storage
    ////////////////////////////////////////////////////////////
    // From ERC721 registry assetId to Trading (to avoid asset collision)
    mapping(bytes32 => Trading) public tradingByNftAsset;

    function createTrading(
        address seller,
        address nftAddress,
        uint256 assetId,
        bytes32 tradingId,
        uint256 priceInWei,
        uint256 expiredAt
    ) 
        external 
        override 
        onlyFrom(tradingMarketplace) 
        whenNotPaused
    {
        bytes32 nftAsset = keccak256(abi.encodePacked(nftAddress, assetId));
        if(!assetIsAvailable(nftAsset)) revert AssetUnvailable();

        tradingByNftAsset[nftAsset] = Trading({
            tradingId: tradingId,
            seller: seller,
            nftAddress: nftAddress,
            price: priceInWei,
            expiredAt: expiredAt
        });
    }

    function getTrading(bytes32 nftAsset)
        external
        view
        override
        returns (Trading memory trading)
    {
        return tradingByNftAsset[nftAsset];
    }

    function deleteTrading(bytes32 nftAsset)
        external
        override
        whenNotPaused
        onlyFrom(tradingMarketplace)
    {
        if(tradingByNftAsset[nftAsset].tradingId == bytes32(0)) revert AssetNotExisted();
        delete tradingByNftAsset[nftAsset];
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
