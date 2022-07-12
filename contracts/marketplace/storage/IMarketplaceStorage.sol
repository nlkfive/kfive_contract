// SPDX-License-Identifier: MIT
pragma solidity 0.8.4;

import "@openzeppelin/contracts/utils/introspection/IERC165.sol";
import "./IPublicAuction.sol";
import "./IBlindAuction.sol";
import "./IOrder.sol";

interface IMarketplaceStorage is IERC165, IOrder, IPublicAuction, IBlindAuction {
    // Error
    error InvalidMkpSender();
    error AssetNotExisted();
    error AssetUnvailable();

    function assetIsAvailable(bytes32 nftAsset) external view returns (bool);
    ////////////////////////////////////////////////////////////
    // Public Auction storage
    ////////////////////////////////////////////////////////////

    /**
     * @dev Returns if the public auction of this asset is existed or not.
     */
    function publicAuctionIsExisted(bytes32 publicAuctionId)
        external
        view
        returns (bool existed);

    /**
     * @dev Returns if the public auction is ended or not.
     */
    function publicAuctionIsEnded(bytes32 nftAsset, bytes32 publicAuctionId)
        external
        view
        returns (bool ended);

    /**
     * @dev Returns if the public auction of this asset is running or not.
     */
    function publicAuctionIsRunning(bytes32 nftAsset, bytes32 publicAuctionId)
        external
        view
        returns (bool existed);

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
    ) external;

    /**
     * @dev Get public auction by nftAsset.
     */
    function getPublicAuction(bytes32 publicAuctionId)
        external
        view
        returns (PublicAuction memory publicAuction);

    /**
     * @dev Delete public auction by nftAsset.
     */
    function endPublicAuction(bytes32 nftAsset) external;

    /**
     * @dev Delete update public auction highest bid.
     */
    function updateHighestBidPublicAuction(address bidder, uint256 highestBid, bytes32 publicAuctionId) external;


    ////////////////////////////////////////////////////////////
    // Blind Auction storage
    ////////////////////////////////////////////////////////////

    /**
     * @dev Returns if the blind auction of this asset is existed or not.
     */
    function blindAuctionIsExisted(bytes32 blindAuctionId)
        external
        view
        returns (bool existed);

    /**
     * @dev Returns if the blind auction is ended or not.
     */
    function blindAuctionIsEnded(bytes32 nftAsset, bytes32 blindAuctionId)
        external
        view
        returns (bool ended);

    /**
     * @dev Returns if the blind auction of this asset is running or not.
     */
    function blindAuctionIsRunning(bytes32 nftAsset, bytes32 blindAuctionId)
        external
        view
        returns (bool existed);

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
    ) external;

    /**
     * @dev Get blind auction by nftAsset.
     */
    function getBlindAuction(bytes32 blindAuctionId)
        external
        view
        returns (BlindAuction memory blindAuction);

    /**
     * @dev Delete blind auction by nftAsset.
     */
    function endBlindAuction(bytes32 nftAsset) external;

    /**
     * @dev Delete update blind auction highest bid.
     */
    function updateHighestBidBlindAuction(address bidder, uint256 highestBid, bytes32 blindAuctionId) external;

    ////////////////////////////////////////////////////////////
    // Order storage
    ////////////////////////////////////////////////////////////

    /**
     * @dev Create new order
     */
    function createOrder(
        address seller,
        address nftAddress,
        uint256 assetId,
        bytes32 orderId,
        uint256 priceInWei,
        uint256 expiredAt
    ) external;

    /**
     * @dev Get order by nftAsset.
     */
    function getOrder(bytes32 nftAsset)
        external
        view
        returns (Order memory order);

    /**
     * @dev Delete order by nftAsset.
     */
    function deleteOrder(bytes32 nftAsset) external;
}
