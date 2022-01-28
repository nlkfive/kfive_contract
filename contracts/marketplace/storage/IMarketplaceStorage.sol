// SPDX-License-Identifier: MIT
pragma solidity 0.8.4;

import "@openzeppelin/contracts/utils/introspection/IERC165.sol";
import "./IAuction.sol";
import "./IOrder.sol";

interface IMarketplaceStorage is IERC165, IOrder, IAuction {
    function assetIsAvailable(bytes32 nftAsset) external view returns (bool);

    ////////////////////////////////////////////////////////////
    // Auction storage
    ////////////////////////////////////////////////////////////

    /**
     * @dev Returns if the auction of this asset is existed or not.
     */
    function auctionIsExisted(bytes32 auctionId)
        external
        view
        returns (bool existed);

    /**
     * @dev Returns if the auction is ended or not.
     */
    function auctionIsEnded(bytes32 nftAsset, bytes32 auctionId)
        external
        view
        returns (bool ended);

    /**
     * @dev Returns if the auction of this asset is running or not.
     */
    function auctionIsRunning(bytes32 nftAsset, bytes32 auctionId)
        external
        view
        returns (bool existed);

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
    ) external;

    /**
     * @dev Get auction by nftAsset.
     */
    function getAuction(bytes32 auctionId)
        external
        view
        returns (Auction memory order);

    /**
     * @dev Delete auction by nftAsset.
     */
    function auctionEnded(bytes32 nftAsset) external;

    /**
     * @dev Delete auction by nftAsset.
     */
    function updateHighestBid(
        address bidder,
        uint256 highestBid,
        bytes32 auctionId
    ) external;

    ////////////////////////////////////////////////////////////
    // Trading storage
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
