// SPDX-License-Identifier: MIT
pragma solidity 0.8.4;

import "./Marketplace.sol";
import "@openzeppelin/contracts/utils/math/SafeMath.sol";

contract OrderMarketplace is Marketplace {
    using SafeMath for uint256;

    constructor(address _acceptedToken, uint256 _ownerCutPerMillion)
        Marketplace(_acceptedToken, _ownerCutPerMillion)
    {}

    /**
     * @dev Creates a new order
     * @param nftAddress - Non fungible registry address
     * @param assetId - ID of the published NFT
     * @param priceInWei - Price in Wei for the supported coin
     * @param expiresAt - Duration of the order (in hours)
     */
    function createOrder(
        address nftAddress,
        uint256 assetId,
        uint256 priceInWei,
        uint256 expiresAt
    ) external whenNotPaused {
        _createOrder(nftAddress, assetId, priceInWei, expiresAt);
    }

    /**
     * @dev Cancel an already published order
     *  can only be canceled by seller or the contract owner
     * @param nftAddress - Address of the NFT registry
     * @param assetId - ID of the published NFT
     */
    function cancelOrder(address nftAddress, uint256 assetId)
        public
        whenNotPaused
    {
        _cancelOrder(nftAddress, assetId);
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
        uint256 expiresAt
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
            expiresAt > block.timestamp.add(1 minutes),
            "Publication should be more than 1 minute in the future"
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

        orderByAssetId[nftAddress][assetId] = Order({
            id: orderId,
            seller: assetOwner,
            nftAddress: nftAddress,
            price: priceInWei,
            expiresAt: expiresAt
        });

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
            expiresAt
        );
    }

    function _cancelOrder(address nftAddress, uint256 assetId)
        internal
        returns (Order memory)
    {
        address sender = _msgSender();
        Order memory order = orderByAssetId[nftAddress][assetId];

        require(order.id != 0, "Asset not published");
        require(
            order.seller == sender || sender == owner(),
            "Unauthorized user"
        );

        bytes32 orderId = order.id;
        address orderSeller = order.seller;
        address orderNftAddress = order.nftAddress;
        delete orderByAssetId[nftAddress][assetId];

        emit OrderCancelled(orderId, assetId, orderSeller, orderNftAddress);

        return order;
    }

    function _executeOrder(
        address nftAddress,
        uint256 assetId,
        uint256 price
    ) internal _requireERC721(nftAddress) returns (Order memory) {
        address sender = _msgSender();

        IERC721 nftRegistry = IERC721(nftAddress);
        Order memory order = orderByAssetId[nftAddress][assetId];
        require(order.id != 0, "Asset not published");

        address seller = order.seller;
        require(seller != address(0), "Invalid address");
        require(seller != sender, "Unauthorized user");
        require(order.price == price, "The price is not correct");
        require(block.timestamp < order.expiresAt, "The order expired");
        require(
            seller == nftRegistry.ownerOf(assetId),
            "The seller is no longer the owner"
        );

        bytes32 orderId = order.id;
        {
            uint256 saleShareAmount = 0;
            delete orderByAssetId[nftAddress][assetId];

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

        emit OrderSuccessful(
            orderId,
            assetId,
            seller,
            nftAddress,
            price,
            sender
        );

        return order;
    }
}
