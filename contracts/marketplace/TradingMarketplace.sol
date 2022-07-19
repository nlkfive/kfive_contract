// SPDX-License-Identifier: MIT
pragma solidity 0.8.4;

import "./Marketplace.sol";
import "./storage/ITrading.sol";
import "@openzeppelin/contracts/utils/math/SafeMath.sol";

contract TradingMarketplace is ITrading, Marketplace {
    using SafeMath for uint256;

    constructor(
        address _acceptedToken,
        address _beneficary,
        address _marketplaceStorage,
        uint256 _ownerCutPerMillion
    ) Marketplace(_acceptedToken, _beneficary, _marketplaceStorage, _ownerCutPerMillion) { }

    /**
     * @dev Creates a new trading
     * @param nftAddress - Non fungible registry address
     * @param assetId - ID of the published NFT
     * @param priceInWei - Price in Wei for the supported coin
     * @param expiredAt - Duration of the trading (in hours)
     */
    function createTrading(
        address nftAddress,
        uint256 assetId,
        uint256 priceInWei,
        uint256 expiredAt
    ) external whenNotPaused {
        _createTrading(nftAddress, assetId, priceInWei, expiredAt);
    }

    /**
     * @dev Cancel an already published trading
     *  can only be canceled by seller or the contract owner
     * @param nftAsset - keccak256(abi.encodePacked(nftAddress, assetId))
     */
    function cancelTrading(bytes32 nftAsset) public whenNotPaused {
        _cancelTrading(nftAsset);
    }

    /**
     * @dev Executes the sale for a published NFT
     * @param nftAddress - Address of the NFT registry
     * @param assetId - ID of the published NFT
     * @param price - Trading price
     */
    function executeTrading(
        address nftAddress,
        uint256 assetId,
        uint256 price
    ) external whenNotPaused {
        _executeTrading(nftAddress, assetId, price);
    }

    function _createTrading(
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
        if(expiredAt < block.timestamp.add(minStageDuration))
            revert InvalidExpiredTime();
        bytes32 nftAsset = keccak256(abi.encodePacked(nftAddress, assetId));
        if(!marketplaceStorage.assetIsAvailable(nftAsset)) revert Unavailable();

        bytes32 tradingId = keccak256(
            abi.encodePacked(
                block.timestamp,
                assetOwner,
                assetId,
                nftAddress,
                priceInWei
            )
        );

        marketplaceStorage.createTrading(
            assetOwner,
            nftAddress,
            assetId,
            tradingId,
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

        emit TradingCreated(
            tradingId,
            assetId,
            assetOwner,
            nftAddress,
            priceInWei,
            expiredAt
        );
    }

    function _cancelTrading(bytes32 nftAsset) internal {
        address sender = _msgSender();
        Trading memory trading = marketplaceStorage.getTrading(nftAsset);
        bytes32 tradingId = trading.tradingId;

        if(tradingId == 0) revert NotExisted();
        if(
            !(trading.seller == sender || hasRole(CANCEL_ROLE, _msgSender()))
        ) revert Unauthorized();

        marketplaceStorage.deleteTrading(nftAsset);
        emit TradingCancelled(sender, tradingId);
    }

    function _executeTrading(
        address nftAddress,
        uint256 assetId,
        uint256 price
    ) internal returns (Trading memory) {
        bytes32 nftAsset = keccak256(abi.encodePacked(nftAddress, assetId));
        Trading memory trading = marketplaceStorage.getTrading(nftAsset);
        if(trading.tradingId == 0) revert NotExisted();

        address seller = trading.seller;
        address sender = _msgSender();
        if(seller == sender) revert Unauthorized();
        if(trading.price != price) revert InvalidPrice();
        if(block.timestamp > trading.expiredAt) revert TradingExpired();

        IERC721 nftRegistry = IERC721(nftAddress);
        if(seller != nftRegistry.ownerOf(assetId)) revert Unauthorized();

        bytes32 tradingId = trading.tradingId;
        {
            uint256 saleShareAmount = 0;
            marketplaceStorage.deleteTrading(nftAsset);

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

        emit TradingSuccessful(tradingId, sender, seller);

        return trading;
    }
}
