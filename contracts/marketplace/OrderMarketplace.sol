// SPDX-License-Identifier: MIT
pragma solidity 0.8.4;

import "./Marketplace.sol";
import "./storage/IOrder.sol";
import "@openzeppelin/contracts/utils/math/SafeMath.sol";

contract OrderMarketplace is IOrder, Marketplace {
    using SafeMath for uint256;

    constructor(
        address _acceptedToken,
        address _marketplaceStorage,
        uint256 _ownerCutPerMillion
    ) Marketplace(_acceptedToken, _marketplaceStorage, _ownerCutPerMillion) {}

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

        require(sender == assetOwner, "Only the owner can create orders");
        require(
            nftRegistry.getApproved(assetId) == address(this) ||
                nftRegistry.isApprovedForAll(assetOwner, address(this)),
            "The contract is not authorized to manage the asset"
        );
        require(priceInWei > 0, "Price should be bigger than 0");
        require(
            expiredAt > block.timestamp.add(1 minutes),
            "Publication should be more than 1 minute in the future"
        );
        bytes32 nftAsset = keccak256(abi.encodePacked(nftAddress, assetId));
        require(
            marketplaceStorage.assetIsAvailable(nftAsset),
            "This asset is unavailable"
        );

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
            require(
                acceptedToken.transferFrom(
                    sender,
                    owner(),
                    publicationFeeInWei
                ),
                "Transfering the publication fee to the Marketplace owner failed"
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

    function _cancelOrder(bytes32 nftAsset) internal returns (Order memory) {
        address sender = _msgSender();
        Order memory order = marketplaceStorage.getOrder(nftAsset);

        require(order.orderId != 0, "Order is not existed");
        require(
            order.seller == sender || sender == owner(),
            "Unauthorized user"
        );

        bytes32 orderId = order.orderId;
        marketplaceStorage.deleteOrder(nftAsset);

        emit OrderCancelled(orderId);
        return order;
    }

    function _executeOrder(
        address nftAddress,
        uint256 assetId,
        uint256 price
    ) internal returns (Order memory) {
        bytes32 nftAsset = keccak256(abi.encodePacked(nftAddress, assetId));
        Order memory order = marketplaceStorage.getOrder(nftAsset);
        require(order.orderId != 0, "Order is not existed");

        address seller = order.seller;
        require(seller != address(0), "Invalid address");

        address sender = _msgSender();
        require(seller != sender, "Unauthorized user");
        require(order.price == price, "The price is not correct");
        require(block.timestamp < order.expiredAt, "The order expired");

        IERC721 nftRegistry = IERC721(nftAddress);
        require(
            seller == nftRegistry.ownerOf(assetId),
            "The seller is no longer the owner"
        );

        bytes32 orderId = order.orderId;
        {
            uint256 saleShareAmount = 0;
            marketplaceStorage.deleteOrder(nftAsset);

            if (ownerCutPerMillion > 0) {
                // Calculate sale share
                saleShareAmount = price.mul(ownerCutPerMillion).div(1000000);

                // Transfer share amount for marketplace Owner
                require(
                    acceptedToken.transferFrom(
                        sender,
                        owner(),
                        saleShareAmount
                    ),
                    "Transfering the cut to the Marketplace owner failed"
                );
            }

            // Transfer sale amount to seller
            require(
                acceptedToken.transferFrom(
                    sender,
                    seller,
                    price.sub(saleShareAmount)
                ),
                "Transfering the sale amount to the seller failed"
            );
        }

        // Transfer asset owner
        nftRegistry.safeTransferFrom(seller, sender, assetId);

        emit OrderSuccessful(orderId, sender);

        return order;
    }
}
