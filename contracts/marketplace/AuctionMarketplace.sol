// SPDX-License-Identifier: MIT
pragma solidity 0.8.4;

import "./Marketplace.sol";
import "./storage/IAuction.sol";
import "@openzeppelin/contracts/utils/math/SafeMath.sol";

contract AuctionMarketplace is IAuction, Marketplace {
    using SafeMath for uint256;

    // List of bids (bid => user => _blindedBid => deposit value)
    mapping(bytes32 => mapping(address => mapping(bytes32 => uint256))) private _bids;
    // List of pending returns
    mapping(bytes32 => mapping(address => uint256)) private _pendingReturns;

    constructor(
        address acceptedToken,
        address marketplaceStorage,
        address beneficary,
        uint256 ownerCutPerMillion
    ) Marketplace(acceptedToken, beneficary, marketplaceStorage, ownerCutPerMillion) {
    }

    error InvalidBiddingEnd();
    error InvalidRevealEnd();
    error Bidded();

    /**
     * @dev Creates a new auction
     * @param nftAddress - Non fungible registry address
     * @param assetId - ID of the published NFT
     * @param startPriceInWei - Price in Wei for the supported coin
     * @param biddingEnd - Timestamp when bidding end (in seconds)
     * @param revealEnd - Timestamp when reveal end (in seconds)
     */
    function createAuction(
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
            if(biddingEnd < block.timestamp.add(minStageDuration)) revert InvalidBiddingEnd();
            if(revealEnd < biddingEnd.add(minStageDuration)) revert InvalidRevealEnd();
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

        bytes32 auctionId = keccak256(
            abi.encodePacked(
                block.timestamp,
                assetOwner,
                assetId,
                nftAddress,
                startPriceInWei
            )
        );

        marketplaceStorage.createAuction(
            assetOwner,
            nftAddress,
            assetId,
            auctionId,
            biddingEnd,
            revealEnd,
            startPriceInWei
        );

        emit AuctionCreated(
            assetOwner,
            nftAddress,
            auctionId,
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
     * @param auctionId - ID of the auction
     */
    function cancelAuction(bytes32 nftAsset, bytes32 auctionId)
        external
        whenNotPaused
    {
        checkExisted(auctionId);
        checkRunning(nftAsset, auctionId);
        address sender = _msgSender();
        address seller = marketplaceStorage.getAuction(auctionId).seller;

        if(
            !(seller == sender || hasRole(CANCEL_ROLE, _msgSender()))
        ) revert Unauthorized();
        marketplaceStorage.auctionEnded(nftAsset);
        emit AuctionCancelled(auctionId);
    }

    /**
     * @dev Bid the auction for a published NFT
     * Place a blinded bid with `_blindedBid` =
     * keccak256(abi.encodePacked(value, fake, secret)).
     * The sent ether is only refunded if the bid is correctly
     * revealed in the revealing phase. The bid is valid if the
     * ether sent together with the bid is at least "value" and
     * "fake" is not true. Setting "fake" to true and sending
     * not the exact amount are ways to hide the real bid but
     * still make the required deposit. The same address can
     * place multiple bids.
     * @param nftAsset - keccak256(abi.encodePacked(nftAddress, assetId))
     * @param auctionId - ID of the auction
     * @param blindedBid - Encoded Bid price
     * @param deposit - Deposit value (Must be large than bid price)
     */
    function bidAuction(
        bytes32 nftAsset,
        bytes32 auctionId,
        bytes32 blindedBid,
        uint256 deposit
    ) external whenNotPaused {
        checkRunning(nftAsset, auctionId);
        uint256 biddingEnd = marketplaceStorage
            .getAuction(auctionId)
            .biddingEnd;
        onlyBefore(biddingEnd);
        address sender = _msgSender();

        acceptedToken.transferFrom(sender, address(this), deposit);
        if(_bids[auctionId][sender][blindedBid] != 0) revert Bidded();

        _bids[auctionId][sender][blindedBid] = deposit;
        emit BidSuccessful(sender, auctionId, blindedBid);
    }

    /**
     * @dev Reveal your blinded bids. You will get a refund for all
     * correctly blinded invalid bids and for all bids except for
     * the totally highest.
     * @param auctionId - ID of the auction
     * @param _values - Array of bid values
     * @param _fake - Array of true - false values, auction will ignore true value
     * @param _secret - Array of secret values used for verify the encoded bid values in bidding stage
     */
    function revealBid(
        bytes32 nftAsset,
        bytes32 auctionId,
        uint256[] memory _values,
        bool[] memory _fake,
        bytes32[] memory _secret
    ) external whenNotPaused {
        Auction memory _auction = marketplaceStorage.getAuction(auctionId);
        onlyAfter(_auction.biddingEnd);
        address sender = _msgSender();
        uint256 length = _values.length;
        require(_fake.length == length);
        require(_secret.length == length);
        bool isEnded = marketplaceStorage.auctionIsEnded(nftAsset, auctionId);

        uint256 refund;
        for (uint256 i = 0; i < length; i++) {
            refund = refund.add(
                _verifyRevealBid(
                    _fake[i] || isEnded, 
                    sender,
                    _blindedBid(_values[i], _fake[i], _secret[i]),
                    _values[i],
                    auctionId
                )
            );
        }
        if (refund > 0) {
            acceptedToken.transfer(sender, refund);
            emit AuctionRefund(sender, auctionId, refund);
        }
    }

    function _blindedBid(uint256 value, bool fake, bytes32 secret) internal pure returns (bytes32) {
        return keccak256(abi.encodePacked(value, fake, secret));
    }

    /**
     * @dev Withdraw a bid that was overbid.
     * @param auctionId - ID of the auction
     */
    function withdraw(bytes32 auctionId) external whenNotPaused {
        uint256 amount = pendingReturns(auctionId);
        address sender = _msgSender();
        if (amount > 0) {
            // It is important to set this to zero because the recipient
            // can call this function again as part of the receiving call
            // before `transfer` returns (see the remark above about
            // conditions -> effects -> interaction).
            _pendingReturns[auctionId][sender] = 0;
            acceptedToken.transfer(sender, amount);
            emit AuctionRefund(sender, auctionId, amount);
        }
    }

    /**
     * @dev Withdraw a bid that was overbid.
     * @param auctionId - ID of the auction
     */
    function pendingReturns(bytes32 auctionId) public view returns (uint256) {
        checkExisted(auctionId);
        return _pendingReturns[auctionId][_msgSender()];
    }

    /**
     * @dev End the auction and send the highest bid to the beneficiary.
     * @param nftAddress - NFT Address
     * @param assetId - ID of NFT address
     * @param auctionId - ID of the auction
     */
    function closeAuction(
        address nftAddress,
        uint256 assetId,
        bytes32 auctionId
    ) external whenNotPaused {
        checkExisted(auctionId);
        _closeAuction(
            nftAddress,
            assetId,
            keccak256(abi.encodePacked(nftAddress, assetId)),
            auctionId
        );
    }

    function _placeBid(
        bytes32 auctionId,
        address bidder,
        uint256 value
    ) internal returns (bool success) {
        Auction memory _auction = marketplaceStorage.getAuction(auctionId);
        if (value <= _auction.startPrice || value <= _auction.highestBid) {
            return false;
        }
        if (_auction.highestBidder != address(0)) {
            // Refund the previously highest bidder.
            _pendingReturns[_auction.id][
                _auction.highestBidder
            ] = _pendingReturns[_auction.id][_auction.highestBidder].add(
                _auction.highestBid
            );
        }
        // Update highest Bid
        marketplaceStorage.updateHighestBid(bidder, value, _auction.id);
        return true;
    }

    function _verifyRevealBid(
        bool fake,
        address sender,
        bytes32 blindedBid,
        uint256 value,
        bytes32 auctionId
    ) internal returns (uint256 refund) {
        uint256 deposit = _bids[auctionId][sender][blindedBid];
        if (deposit == 0) {
            // Incorrect bid parameter
            emit RevealFailed(fake, auctionId, value);
            return 0;
        }

        // Make it impossible for the sender to re-claim
        // the same deposit.
        _bids[auctionId][sender][blindedBid] = 0;

        // Add to refund
        refund = refund.add(deposit);

        // If this is the highest bid then save it
        if (!fake && deposit >= value && _placeBid(auctionId, sender, value)) {
            refund = refund.sub(value);
        }

        emit RevealSuccessful(fake, auctionId, value, blindedBid);
        return refund;
    }

    function _closeAuction(
        address nftAddress,
        uint256 assetId,
        bytes32 nftAsset,
        bytes32 auctionId
    ) internal {
        checkRunning(nftAsset, auctionId);
        Auction memory _auction = marketplaceStorage.getAuction(auctionId);
        onlyAfter(_auction.revealEnd);

        address auctionHighestBidder = _auction.highestBidder;
        address seller = _auction.seller;
        uint256 auctionHighestBid = _auction.highestBid;

        IERC721 nftRegistry = IERC721(nftAddress);
        if(seller != nftRegistry.ownerOf(assetId)) revert Unauthorized();
        marketplaceStorage.auctionEnded(nftAsset);

        if (auctionHighestBidder != address(0)){
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

            emit AuctionSuccessful(
                seller,
                auctionHighestBidder,
                auctionId,
                auctionHighestBid
            );
        } else {
            emit AuctionEnded(auctionId);
        }
    }

    function onlyBefore(uint256 _time) internal view {
        if (block.timestamp >= _time) revert TooLate(_time);
    }

    function onlyAfter(uint256 _time) internal view {
        if (block.timestamp <= _time) revert TooEarly(_time);
    }

    function checkExisted(bytes32 auctionId) public view {
        if (!marketplaceStorage.auctionIsExisted(auctionId))
            revert NotExisted();
    }

    function checkRunning(bytes32 nftAsset, bytes32 auctionId) public view {
        checkExisted(auctionId);
        if (!marketplaceStorage.auctionIsRunning(nftAsset, auctionId))
            revert NotRunning();
    }

}
