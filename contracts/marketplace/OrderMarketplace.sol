// SPDX-License-Identifier: MIT
pragma solidity 0.8.4;

import "./Marketplace.sol";
import "./storage/IOrder.sol";
import "@openzeppelin/contracts/utils/math/SafeMath.sol";

contract OrderMarketplace is IOrder, Marketplace {
    using SafeMath for uint256;
    
    error InvalidExpiredTime();
    error OrderExpired();

    constructor(
        address _acceptedToken,
        address _beneficary,
        address _marketplaceStorage,
        uint256 _ownerCutPerMillion
    ) Marketplace(_acceptedToken, _beneficary, _marketplaceStorage, _ownerCutPerMillion) { }

    /**
     * @dev Creates a new order
     * @param nftAddress - Non fungible registry address
     * @param assetId - ID of the published NFT
     * @param priceInWei - Price in Wei for the supported coin
     * @param expiredAt - Duration of the order (in hours)
     */
    function createOrder(
        address nftAddress,
        uint256 assetId,
        uint256 priceInWei,
        uint256 expiredAt
    ) external whenNotPaused {
        _createOrder(nftAddress, assetId, priceInWei, expiredAt);
    }

    /**
     * @dev Cancel an already published order
     *  can only be canceled by seller or the contract owner
     * @param nftAsset - keccak256(abi.encodePacked(nftAddress, assetId))
     */
    function cancelOrder(bytes32 nftAsset) public whenNotPaused {
        _cancelOrder(nftAsset);
    }

    /**
     * @dev Executes the sale for a published NFT
     * @param nftAddress - Address of the NFT registry
     * @param assetId - ID of the published NFT
     * @param price - Order price
     */
    function executeOrder(
        address nftAddress,
        uint256 assetId,
        uint256 price
    ) external whenNotPaused {
        _executeOrder(nftAddress, assetId, price);
    }

    function _createOrder(
        address nftAddress,
        uint256 assetId,
        uint256 priceInWei,
        uint256 expiredAt
    ) internal _requireERC721(nftAddress) {
        address sender = _msgSender();

        IERC721 nftRegistry = IERC721(nftAddress);
        address assetOwner = nftRegistry.ownerOf(assetId);

        if(sender != assetOwner || 
            !(
                nftRegistry.getApproved(assetId) == address(this) ||
                nftRegistry.isApprovedForAll(assetOwner, address(this))
            )
        ) revert Unauthorized();
        if(priceInWei == 0) revert InvalidPrice();
        if(expiredAt < block.timestamp.add(minStageDuration)) revert InvalidExpiredTime();
        bytes32 nftAsset = keccak256(abi.encodePacked(nftAddress, assetId));
        if(!marketplaceStorage.assetIsAvailable(nftAsset)) revert Unavailable();

        bytes32 orderId = keccak256(
            abi.encodePacked(
                block.timestamp,
                assetOwner,
                assetId,
                nftAddress,
                priceInWei
            )
        );

        marketplaceStorage.createOrder(
            assetOwner,
            nftAddress,
            assetId,
            orderId,
            priceInWei,
            expiredAt
        );

        // Check if there's a publication fee and
        // transfer the amount to marketplace owner
        if (publicationFeeInWei > 0) {
            acceptedToken.transferFrom(
                sender,
                beneficary,
                publicationFeeInWei
            );
        }

        emit OrderCreated(
            orderId,
            assetId,
            assetOwner,
            nftAddress,
            priceInWei,
            expiredAt
        );
    }

    function _cancelOrder(bytes32 nftAsset) internal {
        address sender = _msgSender();
        Order memory order = marketplaceStorage.getOrder(nftAsset);
        bytes32 orderId = order.orderId;

        if(orderId == 0) revert NotExisted();
        if(
            !(order.seller == sender || hasRole(CANCEL_ROLE, _msgSender()))
        ) revert Unauthorized();

        marketplaceStorage.deleteOrder(nftAsset);
        emit OrderCancelled(orderId);
    }

    function _executeOrder(
        address nftAddress,
        uint256 assetId,
        uint256 price
    ) internal returns (Order memory) {
        bytes32 nftAsset = keccak256(abi.encodePacked(nftAddress, assetId));
        Order memory order = marketplaceStorage.getOrder(nftAsset);
        if(order.orderId == 0) revert NotExisted();

        address seller = order.seller;
        address sender = _msgSender();
        if(seller == sender) revert Unauthorized();
        if(order.price != price) revert InvalidPrice();
        if(block.timestamp > order.expiredAt) revert OrderExpired();

        IERC721 nftRegistry = IERC721(nftAddress);
        if(seller != nftRegistry.ownerOf(assetId)) revert Unauthorized();

        bytes32 orderId = order.orderId;
        {
            uint256 saleShareAmount = 0;
            marketplaceStorage.deleteOrder(nftAsset);

            if (ownerCutPerMillion > 0) {
                // Calculate sale share
                saleShareAmount = price.mul(ownerCutPerMillion).div(1000000);

                // Transfer share amount for marketplace Owner
                acceptedToken.transferFrom(
                    sender,
                    beneficary,
                    saleShareAmount
                );
            }

            // Transfer sale amount to seller
            acceptedToken.transferFrom(
                sender,
                seller,
                price.sub(saleShareAmount)
            );
        }

        // Transfer asset owner
        nftRegistry.safeTransferFrom(seller, sender, assetId);

        emit OrderSuccessful(orderId, sender);

        return order;
    }
}
