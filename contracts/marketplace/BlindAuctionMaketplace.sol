// SPDX-License-Identifier: MIT
pragma solidity 0.8.4;

import "./Marketplace.sol";
import "./storage/IBlindAuction.sol";
import "@openzeppelin/contracts/utils/math/SafeMath.sol";

contract BlindAuctionMarketplace is IBlindAuction, Marketplace {
    using SafeMath for uint256;

    // List of bids (bid => (user => _blindedBid[]))
    mapping(bytes32 => mapping(address => bytes32[])) private _blindedBids;

    // List of max deposits (auction => (user => max deposit))
    mapping(bytes32 => mapping(address => uint256)) private _deposits;

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
     * @param revealEnd - Timestamp when reveal end (in seconds)
     */
    function createBlindAuction(
        address nftAddress,
        uint256 assetId,
        uint256 startPriceInWei,
        uint256 biddingEnd,
        uint256 revealEnd
    )
        external 
        _requireERC721(nftAddress) 
        whenNotPaused
    {
        // Validate input
        address assetOwner;
        {
            if(biddingEnd < block.timestamp.add(minStageDuration))
                revert InvalidBiddingEnd();
            if(startPriceInWei == 0) revert InvalidPrice();
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

        bytes32 blindAuctionId = keccak256(
            abi.encodePacked(
                block.timestamp,
                assetOwner,
                assetId,
                nftAddress,
                startPriceInWei
            )
        );

        marketplaceStorage.createBlindAuction(
            assetOwner,
            nftAddress,
            assetId,
            blindAuctionId,
            biddingEnd,
            revealEnd,
            startPriceInWei
        );

        emit BlindAuctionCreatedSuccessful(
            assetOwner,
            nftAddress,
            blindAuctionId,
            assetId,
            biddingEnd,
            revealEnd,
            startPriceInWei
        );
    }

    /**
     * @dev Cancel an already published auction
     *  can only be canceled by seller or the contract owner
     * @param nftAsset - keccak256(abi.encodePacked(nftAddress, assetId))
     * @param blindAuctionId - ID of the auction
     */
    function cancelBlindAuction(bytes32 nftAsset, bytes32 blindAuctionId)
        external
        whenNotPaused
    {
        checkRunning(nftAsset, blindAuctionId);

        BlindAuction memory _blindAuction = marketplaceStorage.getBlindAuction(blindAuctionId);

        onlyBefore(_blindAuction.biddingEnd);

        address sender = _msgSender();

        if(!(_blindAuction.seller == sender || hasRole(CANCEL_ROLE, _msgSender()))) revert Unauthorized();

        marketplaceStorage.endBlindAuction(nftAsset);
        emit AuctionCancelledSuccessful(sender, blindAuctionId);
    }

    /**
     * @dev Bid the auction for a published NFT
     * @param nftAsset - keccak256(abi.encodePacked(nftAddress, assetId))
     * @param blindAuctionId - ID of the auction
     * @param blindedBid - Blinded bid
     * @param deposit - Deposit value
     */
    function bidBlindAuction(
        bytes32 nftAsset,
        bytes32 blindAuctionId,
        bytes32 blindedBid,
        uint256 deposit
    ) external whenNotPaused {
        checkRunning(nftAsset, blindAuctionId);

        BlindAuction memory _blindAuction = marketplaceStorage.getBlindAuction(blindAuctionId);

        // auction have not ended yet
        onlyBefore(_blindAuction.biddingEnd);

        address sender = _msgSender();

        if(deposit <= _blindAuction.startPrice) {
            revert InvalidBiddingPrice();
        }

        if(deposit > _deposits[blindAuctionId][sender]) {
            acceptedToken.transferFrom(sender, address(this), deposit - _deposits[blindAuctionId][sender]);
            _deposits[blindAuctionId][sender] = deposit;
        }

        _blindedBids[blindAuctionId][sender].push(blindedBid);

        emit BlindAuctionBidSuccessful(sender, blindAuctionId, blindedBid);
    }

    /**
     * @dev Withdraw a bid that was overbid.
     * @param blindAuctionId - ID of the auction
     * @param assetId - ID of the nft
     * @param nftAddress - Nft address
     */
    function withdraw(
        bytes32 blindAuctionId, 
        uint256 assetId, 
        address nftAddress, 
        bytes32 nftAsset) external whenNotPaused {
        checkExisted(blindAuctionId);

        BlindAuction memory _blindAuction = marketplaceStorage.getBlindAuction(blindAuctionId);

        if(!marketplaceStorage.blindAuctionIsEnded(nftAsset, blindAuctionId)) {
            marketplaceStorage.endBlindAuction(nftAsset);
        }

        address sender = _msgSender();
        uint256 refund = _deposits[blindAuctionId][sender];
        if(refund == 0) {
            revert InvalidWithdraw();
        }

        // handle if sender is highest bidder
        if(sender == _blindAuction.highestBidder) {
            refund = refund - _blindAuction.highestBid;
            _grantReward(nftAddress, assetId, blindAuctionId);
        }

        // refund
        acceptedToken.transfer(sender, refund);

        // update deposit value to zero
        _deposits[blindAuctionId][sender] = 0;

        emit AuctionRefundSuccessful(sender, blindAuctionId, refund);
    }

    function _grantReward(
        address nftAddress,
        uint256 assetId,
        bytes32 blindAuctionId
    ) internal {
        BlindAuction memory _auction = marketplaceStorage.getBlindAuction(blindAuctionId);

        address auctionHighestBidder = _auction.highestBidder;
        address seller = _auction.seller;
        uint256 auctionHighestBid = _auction.highestBid;

        IERC721 nftRegistry = IERC721(nftAddress);
        if(seller != nftRegistry.ownerOf(assetId)) revert Unauthorized();

        if (auctionHighestBidder == address(0) || auctionHighestBid == 0){
            revert RewardGranted();
        }

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
        marketplaceStorage.updateHighestBidBlindAuction(address(0), 0, blindAuctionId);

        emit GrantAuctionRewardSuccessful(
            auctionHighestBidder,
            blindAuctionId,
            assetId
        );
    }

    /**
     * @dev Reveal your blinded bids. You will get a refund for all
     * correctly blinded invalid bids and for all bids except for
     * the totally highest.
     * @param blindAuctionId - ID of the auction
     * @param _values - Array of bid values
     * @param secret - The secret values used for verify the encoded bid values in bidding stage
    */
    function RevealBid(
        bytes32 blindAuctionId,
        bytes32 nftAsset,
        uint256[] memory _values,
        bytes32 secret
    ) external whenNotPaused  {
        checkRunning(nftAsset, blindAuctionId);

        BlindAuction memory _blindAuction = marketplaceStorage.getBlindAuction(blindAuctionId);

        onlyAfter(_blindAuction.biddingEnd);
        onlyBefore(_blindAuction.revealEnd);

        if(marketplaceStorage.blindAuctionIsEnded(nftAsset, blindAuctionId)) {
            revert AuctionEnded();
        }

        address sender = _msgSender();
        uint256 maxDeposit = _deposits[blindAuctionId][sender];
        uint256 length = _values.length;

        // sender has not bid yet
        if (maxDeposit == 0 || length == 0) {
            revert NotBidYet();
        }

        for (uint256 i = 0; i < length; i++) {
            if(_verifyBid(_values[i], secret, _blindedBids[blindAuctionId][sender][i])) {
                if (maxDeposit >= _values[i] && _blindAuction.highestBid < _values[i]) {
                    marketplaceStorage.updateHighestBidBlindAuction(sender, _values[i], blindAuctionId);
                }
            }
        }
        emit RevealSuccessful(sender, blindAuctionId);
    }

    function _verifyBid(
        uint256 value,
        bytes32 secret,
        bytes32 blindedBid
    ) internal pure returns (bool) {
        if(_encodeBid(value, secret) == blindedBid) {
            return true;
        }
        return false;
    }

    function _encodeBid(uint256 value, bytes32 secret) internal pure returns (bytes32) {
        return keccak256(abi.encodePacked(value, secret));
    }

    function getBlindBid(bytes32 auctionId, address bidder) public view returns (bytes32[] memory) {
        return _blindedBids[auctionId][bidder];
    }

    function onlyBefore(uint256 _time) internal view {
        if (block.timestamp >= _time) revert TooLate(_time);
    }

    function onlyAfter(uint256 _time) internal view {
        if (block.timestamp <= _time) revert TooEarly(_time);
    }

    function checkExisted(bytes32 auctionId) public view {
        if (!marketplaceStorage.blindAuctionIsExisted(auctionId))
            revert NotExisted();
    }

    function checkRunning(bytes32 nftAsset, bytes32 auctionId) public view {
        checkExisted(auctionId);
        if (!marketplaceStorage.blindAuctionIsRunning(nftAsset, auctionId))
            revert NotRunning();
    }

}
