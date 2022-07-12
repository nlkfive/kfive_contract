// SPDX-License-Identifier: MIT
pragma solidity 0.8.4;

import "./Marketplace.sol";
import "./storage/IPublicAuction.sol";
import "@openzeppelin/contracts/utils/math/SafeMath.sol";

contract PublicAuctionMarketplace is IPublicAuction, Marketplace {
    using SafeMath for uint256;

    constructor(
        address acceptedToken,
        address marketplaceStorage,
        address beneficary,
        uint256 ownerCutPerMillion
    ) Marketplace(acceptedToken, beneficary, marketplaceStorage, ownerCutPerMillion) {
    }

    /**
     * @dev Creates a new auction
     * @param nftAddress - Non fungible registry address
     * @param assetId - ID of the published NFT
     * @param startPriceInWei - Price in Wei for the supported coin
     * @param biddingEnd - Timestamp when bidding end (in seconds)
     * @param minIncrement - Min bid increment price in wei
     */
    function createAuction(
        address nftAddress,
        uint256 assetId,
        uint256 startPriceInWei,
        uint256 biddingEnd,
        uint256 minIncrement
    )
        external 
        _requireERC721(nftAddress) 
        whenNotPaused
    {
        // Validate input
        address assetOwner;
        {
            if(biddingEnd < block.timestamp.add(minStageDuration)) revert InvalidBiddingEnd();
            if(startPriceInWei == 0 || minIncrement == 0) revert InvalidPrice();
            address sender = _msgSender();
            {
                // Check owner
                IERC721 nftRegistry = IERC721(nftAddress);
                assetOwner = nftRegistry.ownerOf(assetId);
                if(sender != assetOwner || 
                    !(
                        nftRegistry.getApproved(assetId) == address(this) ||
                        nftRegistry.isApprovedForAll(assetOwner, address(this))
                    )
                ) revert Unauthorized();

                // Check if asset is available
                if (!marketplaceStorage.assetIsAvailable(
                    keccak256( // nftAsset
                        abi.encodePacked(nftAddress, assetId)
                    )
                )) {
                    revert Unavailable();
                }
            }
            // Check if there's a publication fee and
            // transfer the amount to marketplace owner
            if (publicationFeeInWei > 0) {
                acceptedToken.transferFrom(
                    sender,
                    beneficary,
                    publicationFeeInWei
                );
            }
        }

        bytes32 publicAuctionId = keccak256(
            abi.encodePacked(
                block.timestamp,
                assetOwner,
                assetId,
                nftAddress,
                startPriceInWei
            )
        );

        marketplaceStorage.createPublicAuction(
            assetOwner,
            nftAddress,
            assetId,
            publicAuctionId,
            biddingEnd,
            startPriceInWei,
            minIncrement
        );

        emit PublicAuctionCreatedSuccessful(
            assetOwner,
            nftAddress,
            publicAuctionId,
            assetId,
            biddingEnd,
            startPriceInWei,
            minIncrement
        );
    }

    /**
     * @dev Cancel an already published auction
     *  can only be canceled by seller or the contract owner
     * @param nftAsset - keccak256(abi.encodePacked(nftAddress, assetId))
     * @param publicAuctionId - ID of the public auction
     */
    function cancelAuction(
        bytes32 nftAsset, 
        bytes32 publicAuctionId
    )
        external
        whenNotPaused
    {
        checkRunning(nftAsset, publicAuctionId);

        PublicAuction memory _publicAuction = marketplaceStorage.getPublicAuction(publicAuctionId);

        // Auction cannot be canceled if after bid time
        onlyBefore(_publicAuction.biddingEnd);

        address sender = _msgSender();

        if(!(_publicAuction.seller == sender || hasRole(CANCEL_ROLE, _msgSender()))) {
            revert Unauthorized();
        }

        // Refund to the current highest bid
        if(_publicAuction.highestBidder != address(0)) {
            acceptedToken.transfer(_publicAuction.highestBidder, _publicAuction.highestBid);
            emit AuctionRefundSuccessful(_publicAuction.highestBidder, publicAuctionId, _publicAuction.highestBid);
            marketplaceStorage.updateHighestBidPublicAuction(address(0), 0, publicAuctionId);
        }

        marketplaceStorage.endPublicAuction(nftAsset);
        emit AuctionCancelledSuccessful(sender, publicAuctionId);
    }

    /**
     * @dev Bid the auction for a published NFT
     * @param nftAsset - keccak256(abi.encodePacked(nftAddress, assetId))
     * @param publicAuctionId - ID of the auction
     * @param bidValue - Bid value
     */
    function bid(
        bytes32 nftAsset,
        bytes32 publicAuctionId,
        uint256 bidValue
    ) external whenNotPaused {
        checkRunning(nftAsset, publicAuctionId);

        PublicAuction memory _publicAuction = marketplaceStorage.getPublicAuction(publicAuctionId);

        // auction have not ended yet
        onlyBefore(_publicAuction.biddingEnd);

        if(bidValue <= _publicAuction.startPrice + _publicAuction.minIncrement
            || bidValue <= _publicAuction.highestBid + _publicAuction.minIncrement) {
            revert InvalidBiddingPrice();
        }

        address sender = _msgSender();

        if(_publicAuction.highestBidder == address(0)) {
            // first bid
            acceptedToken.transferFrom(sender, address(this), bidValue);
        } else {
            if(sender == _publicAuction.highestBidder) {
                acceptedToken.transferFrom(sender, address(this), bidValue - _publicAuction.highestBid);
            } else {
                acceptedToken.transferFrom(sender, address(this), bidValue);
                // Refund to the previous highestBidder
                acceptedToken.transfer(_publicAuction.highestBidder, _publicAuction.highestBid);
                emit AuctionRefundSuccessful(_publicAuction.highestBidder, publicAuctionId, _publicAuction.highestBid);
            }
        }

        // Update highest bid
        marketplaceStorage.updateHighestBidPublicAuction(sender, bidValue, publicAuctionId);

        emit PublicAuctionBidSuccessful(sender, publicAuctionId, bidValue);
    }

    /**
     * @dev Winner receive auction reward.
     * @param publicAuctionId - ID of the auction
     * @param nftAsset - ID of the nft
     * @param assetId - ID of the nft
     * @param nftAddress - Nft address
     */
    function receiveReward(
        bytes32 publicAuctionId, 
        bytes32 nftAsset, 
        uint256 assetId, 
        address nftAddress
    ) external whenNotPaused {
        checkExisted(publicAuctionId);

        PublicAuction memory _publicAuction = marketplaceStorage.getPublicAuction(publicAuctionId);

        // after bid time
        onlyAfter(_publicAuction.biddingEnd);

        address sender = _msgSender();
        address auctionHighestBidder = _publicAuction.highestBidder;
        address seller = _publicAuction.seller;
        uint256 auctionHighestBid = _publicAuction.highestBid;

        // check sender is winner
        if(sender != auctionHighestBidder) {
            revert NotWinner();
        }

        // Check if reward has been granted
        if(auctionHighestBidder == address(0) ||
            auctionHighestBid == 0) {
            revert RewardGranted();
        }

        IERC721 nftRegistry = IERC721(nftAddress);
        if(seller != nftRegistry.ownerOf(assetId)) revert Unauthorized();

        // end auction storage
        marketplaceStorage.endPublicAuction(nftAsset);
        marketplaceStorage.updateHighestBidPublicAuction(address(0), 0, publicAuctionId);

        uint256 saleShareAmount = 0;

        if (ownerCutPerMillion > 0) {
            // Calculate sale share
            saleShareAmount = auctionHighestBid.mul(ownerCutPerMillion).div(
                1000000
            );

            // Transfer share amount for marketplace Owner
            acceptedToken.transfer(beneficary, saleShareAmount);
        }

        // Transfer sale amount to seller
        acceptedToken.transfer(
            seller,
            auctionHighestBid.sub(saleShareAmount)
        );

        // Transfer asset owner
        nftRegistry.safeTransferFrom(seller, auctionHighestBidder, assetId);

        emit GrantAuctionRewardSuccessful(
            auctionHighestBidder,
            publicAuctionId,
            assetId
        );
    } 

    function onlyBefore(uint256 _time) internal view {
        if (block.timestamp >= _time) revert TooLate(_time);
    }

    function onlyAfter(uint256 _time) internal view {
        if (block.timestamp <= _time) revert TooEarly(_time);
    }

    function checkExisted(bytes32 publicAuctionId) public view {
        if (!marketplaceStorage.publicAuctionIsExisted(publicAuctionId))
            revert NotExisted();
    }

    function checkRunning(bytes32 nftAsset, bytes32 publicAuctionId) public view {
        checkExisted(publicAuctionId);
        if (!marketplaceStorage.publicAuctionIsRunning(nftAsset, publicAuctionId))
            revert NotRunning();
    }
}
