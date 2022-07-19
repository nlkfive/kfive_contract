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
     * @param blindAuctionId - keccak256(abi.encodePacked(nftAddress, assetId, startPriceInWei, biddingEnd, revealEnd, time.now))
     * @param nftAsset - keccak256(abi.encodePacked(nftAddress, assetId))
     * @param assetId - ID of the published NFT
     * @param startPriceInWei - Price in Wei for the supported coin
     * @param biddingEnd - Timestamp when bidding end (in seconds)
     * @param revealEnd - Timestamp when reveal end (in seconds)
     */
    function createAuction(
        address nftAddress,
        bytes32 blindAuctionId,
        bytes32 nftAsset,
        uint256 assetId,
        uint256 startPriceInWei,
        uint256 startTime,
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
            if(startTime < block.timestamp || startTime.add(minStageDuration) > biddingEnd) 
                revert InvalidTime();
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
                if (!marketplaceStorage.assetIsAvailable(nftAsset)) {
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

        marketplaceStorage.createBlindAuction(
            assetOwner,
            nftAsset,
            blindAuctionId,
            startTime,
            biddingEnd,
            revealEnd,
            startPriceInWei
        );

        emit BlindAuctionCreatedSuccessful(
            assetOwner,
            nftAddress,
            blindAuctionId,
            assetId,
            startTime,
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
    function cancelAuction(
        bytes32 nftAsset, 
        bytes32 blindAuctionId
    )
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
    function bid(
        bytes32 nftAsset,
        bytes32 blindAuctionId,
        bytes32 blindedBid,
        uint256 deposit
    ) external whenNotPaused {
        checkRunning(nftAsset, blindAuctionId);

        BlindAuction memory _blindAuction = marketplaceStorage.getBlindAuction(blindAuctionId);

        onlyAfter(_blindAuction.startTime);
        onlyBefore(_blindAuction.biddingEnd);

        address sender = _msgSender();

        // deposit >= startPrice
        if(deposit < _blindAuction.startPrice) {
            revert InvalidBiddingPrice();
        }

        // send max deposit to smc
        if(deposit > _deposits[blindAuctionId][sender]) {
            acceptedToken.transferFrom(sender, address(this), deposit - _deposits[blindAuctionId][sender]);
            _deposits[blindAuctionId][sender] = deposit;
        }

        // save blind value bid
        _blindedBids[blindAuctionId][sender].push(blindedBid);

        emit BlindAuctionBidSuccessful(sender, blindAuctionId, blindedBid, deposit);
    }

    /**
     * @dev Withdraw a bid that was overbid.
     * @param blindAuctionId - ID of the auction
     * @param assetId - ID of the nft
     * @param nftAddress - Nft address
     */
    function withdraw(
        address nftAddress, 
        bytes32 nftAsset,
        bytes32 blindAuctionId, 
        uint256 assetId
    ) external whenNotPaused {
        checkExisted(blindAuctionId);

        BlindAuction memory _blindAuction = marketplaceStorage.getBlindAuction(blindAuctionId);

        if(!marketplaceStorage.blindAuctionIsEnded(nftAsset, blindAuctionId)) {
            onlyAfter(_blindAuction.revealEnd);
            marketplaceStorage.endBlindAuction(nftAsset);
        }

        address sender = _msgSender();
        // check deposit is exists
        uint256 refund = _deposits[blindAuctionId][sender];
        if(refund == 0) {
            revert InvalidWithdraw();
        }

        // handle if sender is highest bidder
        if(sender == _blindAuction.highestBidder) {
            refund = refund - _blindAuction.highestBid;
            _grantReward(nftAddress, blindAuctionId, assetId);
        }

        // refund
        acceptedToken.transfer(sender, refund);

        // update deposit value to zero
        _deposits[blindAuctionId][sender] = 0;

        emit AuctionRefundSuccessful(sender, blindAuctionId, refund);
    }

    function _grantReward(
        address nftAddress,
        bytes32 blindAuctionId,
        uint256 assetId
    ) internal {
        BlindAuction memory _auction = marketplaceStorage.getBlindAuction(blindAuctionId);

        address seller = _auction.seller;

        IERC721 nftRegistry = IERC721(nftAddress);
        if(seller != nftRegistry.ownerOf(assetId)) revert Unauthorized();

        address auctionHighestBidder = _auction.highestBidder;

        if (auctionHighestBidder == address(0)) {
            revert RewardGranted();
        }

        uint256 saleShareAmount = 0;
        uint256 auctionHighestBid = _auction.highestBid;

        if (ownerCutPerMillion > 0) {
            // Calculate sale share
            saleShareAmount = auctionHighestBid.mul(ownerCutPerMillion).div(1000000);

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
            blindAuctionId,
            assetId
        );
    }

    /**
     * @dev Reveal your blinded bids. You will get a refund for all
     * correctly blinded invalid bids and for all bids except for
     * the totally highest.
     * @param blindAuctionId - ID of the auction
     * @param nftAsset - keccak256(abi.encodePacked(nftAddress, assetId))
     * @param secret - The secret values used for verify the encoded bid values in bidding stage
     * @param _values - Array of bid values
    */
    function reveal(
        bytes32 blindAuctionId,
        bytes32 nftAsset,
        bytes32 secret,
        uint256[] memory _values
    ) external whenNotPaused  {
        checkRunning(nftAsset, blindAuctionId);

        BlindAuction memory _blindAuction = marketplaceStorage.getBlindAuction(blindAuctionId);

        onlyAfter(_blindAuction.biddingEnd);
        onlyBefore(_blindAuction.revealEnd);

        address sender = _msgSender();
        uint256 maxDeposit = _deposits[blindAuctionId][sender];
        uint256 length = _values.length;

        // sender has not bid yet
        if (maxDeposit == 0) {
            revert NotBidYet();
        }

        bytes32[] memory blindedBids = _blindedBids[blindAuctionId][sender];

        if(!(length == blindedBids.length)) {
            revert InvalidReveal();
        }

        uint256 maxValue = 0;

        // get the highest bid
        for (uint i = 0; i < length; i = _unsafe_inc(i)) {
            if(
                keccak256(abi.encodePacked(_values[i], secret)) == blindedBids[i] && 
                maxValue < _values[i]
            ) {
                maxValue = _values[i];
            }
            emit RevealSuccessful(sender, blindAuctionId, blindedBids[i], _values[i]);
        }

        // update if bid is less than or equal to max deposit and greater than current highest bid
        if(!(maxDeposit < maxValue) && maxValue > _blindAuction.highestBid) {
            marketplaceStorage.updateHighestBidBlindAuction(sender, maxValue, blindAuctionId);
        }

        
    }

    function _unsafe_inc(uint x) private pure returns (uint) {
        unchecked { return x + 1; }
    }

    // show list blind bid of user
    function getBlindBid(bytes32 blindAuctionId, address bidder) public view returns (bytes32[] memory) {
        return _blindedBids[blindAuctionId][bidder];
    }

    function onlyBefore(uint256 _time) internal view {
        if (!(block.timestamp < _time)) revert TooLate(_time); // 0x691e5682
    }

    function onlyAfter(uint256 _time) internal view {
        if (!(block.timestamp > _time)) revert TooEarly(_time); // 0x2a35a324
    }

    function checkExisted(bytes32 blindAuctionId) public view {
        if (!marketplaceStorage.blindAuctionIsExisted(blindAuctionId))
            revert NotExisted(); // 0xafdd4890
    }

    function checkRunning(bytes32 nftAsset, bytes32 blindAuctionId) public view {
        checkExisted(blindAuctionId);
        if (!marketplaceStorage.blindAuctionIsRunning(nftAsset, blindAuctionId))
            revert NotRunning(); // 0xfb18c2ce
    }

}
