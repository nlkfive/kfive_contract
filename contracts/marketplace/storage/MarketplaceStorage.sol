// SPDX-License-Identifier: MIT
pragma solidity 0.8.4;

import "./IAuction.sol";
import "./IOrder.sol";
import "./IMarketplaceStorage.sol";
import "@openzeppelin/contracts/token/ERC721/IERC721.sol";
import "@openzeppelin/contracts/utils/introspection/ERC165.sol";
import "@openzeppelin/contracts/access/Ownable.sol";
import "@openzeppelin/contracts/access/AccessControlEnumerable.sol";
import "@openzeppelin/contracts/security/Pausable.sol";
import "@openzeppelin/contracts/utils/math/SafeMath.sol";

contract MarketplaceStorage is
    AccessControlEnumerable,
    Ownable,
    Pausable,
    IMarketplaceStorage
{
    using SafeMath for uint256;

    bytes32 public constant ADMIN_ROLE = keccak256("ADMIN_ROLE");
    bytes32 public constant PAUSER_ROLE = keccak256("PAUSER_ROLE");

    address auctionMarketplace;
    address orderMarketplace;

    constructor() {
        _setupRole(DEFAULT_ADMIN_ROLE, _msgSender());
        _setupRole(ADMIN_ROLE, _msgSender());
        _setupRole(PAUSER_ROLE, _msgSender());

        _setRoleAdmin(PAUSER_ROLE, ADMIN_ROLE);
    }

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
        onlyRole(ADMIN_ROLE)
    {
        orderMarketplace = _orderMarketplace;
    }

    /**
     * @dev Pauses all token transfers.
     *
     * Requirements:
     *
     * - the caller must have the `PAUSER_ROLE`.
     */
    function pause() public onlyRole(PAUSER_ROLE) {
        _pause();
    }

    /**
     * @dev Unpauses all token transfers.
     *
     * Requirements:
     *
     * - the caller must have the `PAUSER_ROLE`.
     */
    function unpause() public onlyRole(PAUSER_ROLE) {
        _unpause();
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
    mapping(bytes32 => Auction) public auctions;
    // From ERC721 registry keccak256(abi.encodePacked(nft,assetId)) to running Auction ID
    mapping(bytes32 => bytes32) public runningActionIds;

    function auctionIsExisted(bytes32 auctionId)
        external
        view
        override
        returns (bool)
    {
        return auctions[auctionId].id != bytes32(0);
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
        uint256 biddingTime,
        uint256 revealTime,
        uint256 startPriceInWei
    ) external override {
        require(
            _msgSender() == auctionMarketplace,
            "Must be called from auction marketplace contract"
        );
        bytes32 nftAsset = keccak256(abi.encodePacked(nftAddress, assetId));
        require(assetIsAvailable(nftAsset), "This asset is unvailable");

        // Update current running before actually create
        runningActionIds[nftAsset] = auctionId;

        // Create new auction
        Auction storage auction = auctions[auctionId];
        {
            uint256 biddingEnd = block.timestamp.add(biddingTime);
            uint256 revealEnd = biddingEnd.add(revealTime);
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
    function auctionEnded(bytes32 nftAsset) external override {
        require(
            _msgSender() == auctionMarketplace,
            "Must be called from auction marketplace contract"
        );
        require(
            runningActionIds[nftAsset] != 0,
            "Auction has been already ended"
        );
        delete runningActionIds[nftAsset];
    }

    /**
     * @dev Update highest bid auction by auction ID.
     */
    function updateHighestBid(
        address bidder,
        uint256 highestBid,
        bytes32 auctionId
    ) external override {
        require(
            _msgSender() == auctionMarketplace,
            "Must be called from auction marketplace contract"
        );
        Auction storage auction = auctions[auctionId];
        onlyBefore(auction.biddingEnd);

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
    ) external override {
        require(
            _msgSender() == orderMarketplace,
            "Must be called from order marketplace contract"
        );
        bytes32 nftAsset = keccak256(abi.encodePacked(nftAddress, assetId));
        require(assetIsAvailable(nftAsset), "This asset is unvailable");

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

    function deleteOrder(bytes32 nftAsset) external override {
        require(
            _msgSender() == orderMarketplace,
            "Must be called from order marketplace contract"
        );
        require(
            orderByNftAsset[nftAsset].orderId != bytes32(0),
            "Order is not existed"
        );
        delete orderByNftAsset[nftAsset];
    }

    function onlyBefore(uint256 _time) internal view {
        if (block.timestamp >= _time) revert TooLate(_time);
    }

    function onlyAfter(uint256 _time) internal view {
        if (block.timestamp <= _time) revert TooEarly(_time);
    }
}
