// SPDX-License-Identifier: MIT
pragma solidity 0.8.4;

import "./Marketplace.sol";
import "@openzeppelin/contracts/utils/math/SafeMath.sol";

contract AuctionMarketplace is BlindAuctionStorage, Marketplace {
    using SafeMath for uint256;

    constructor(address _acceptedToken, uint256 _ownerCutPerMillion)
        Marketplace(_acceptedToken, _ownerCutPerMillion)
    {}

    /**
     * @dev Creates a new auction
     * @param nftAddress - Non fungible registry address
     * @param assetId - ID of the published NFT
     * @param startPriceInWei - Price in Wei for the supported coin
     * @param biddingTime - Duration of the bidding (in seconds)
     * @param revealTime - Duration of the reveal (in seconds)
     */
    function createAuction(
        address nftAddress,
        uint256 assetId,
        uint256 startPriceInWei,
        uint256 biddingTime,
        uint256 revealTime
    ) external _requireERC721(nftAddress) whenNotPaused {
        _createAuction(
            nftAddress,
            assetId,
            startPriceInWei,
            biddingTime,
            revealTime
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
        auctionExisted(nftAsset, auctionId)
        auctionIsRunning(nftAsset, auctionId)
    {
        _cancelAuction(nftAsset, auctionId);
    }

    /**
     * @dev Bid the auction for a published NFT
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
    )
        external
        whenNotPaused
        auctionExisted(nftAsset, auctionId)
        auctionIsRunning(nftAsset, auctionId)
    {
        _bidAuction(nftAsset, auctionId, blindedBid, deposit);
    }

    /**
     * @dev Reveal your blinded bids. You will get a refund for all
     * correctly blinded invalid bids and for all bids except for
     * the totally highest.
     * @param nftAsset - keccak256(abi.encodePacked(nftAddress, assetId))
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
    )
        external
        whenNotPaused
        auctionExisted(nftAsset, auctionId)
        auctionIsRunning(nftAsset, auctionId)
    {
        _revealBid(nftAsset, auctionId, _values, _fake, _secret);
    }

    /**
     * @dev Withdraw a bid that was overbid.
     * @param nftAsset - keccak256(abi.encodePacked(nftAddress, assetId))
     * @param auctionId - ID of the auction
     */
    function withdraw(bytes32 nftAsset, bytes32 auctionId)
        external
        whenNotPaused
        auctionExisted(nftAsset, auctionId)
    {
        _withdraw(nftAsset, auctionId);
    }

    /**
     * @dev Withdraw a bid that was overbid.
     * @param nftAsset - keccak256(abi.encodePacked(nftAddress, assetId))
     * @param auctionId - ID of the auction
     */
    function pendingReturns(bytes32 nftAsset, bytes32 auctionId)
        external
        view
        auctionExisted(nftAsset, auctionId)
        returns (uint256)
    {
        address sender = _msgSender();
        Auction storage _auction = nftAuctions[nftAsset].auctions[auctionId];
        uint256 amount = _auction.pendingReturns[sender];
        return amount;
    }

    /**
     * @dev End the auction and send the highest bid to the beneficiary.
     * @param nftAddress - NFT Address
     * @param assetId - ID of NFT address
     * @param auctionId - ID of the auction
     */
    function auctionEnd(
        address nftAddress,
        uint256 assetId,
        bytes32 auctionId
    ) external whenNotPaused {
        bytes32 nftAsset = keccak256(abi.encodePacked(nftAddress, assetId));
        _auctionEnd(nftAddress, assetId, nftAsset, auctionId);
    }

    function _createAuction(
        address nftAddress,
        uint256 assetId,
        uint256 startPriceInWei,
        uint256 biddingTime,
        uint256 revealTime
    ) internal _requireERC721(nftAddress) {
        bytes32 nftAsset = keccak256(abi.encodePacked(nftAddress, assetId));

        // Validate input
        address assetOwner;
        {
            require(startPriceInWei > 0, "Price should be bigger than 0");
            address sender = _msgSender();
            {
                // Check owner
                IERC721 nftRegistry = IERC721(nftAddress);
                assetOwner = nftRegistry.ownerOf(assetId);
                require(
                    sender == assetOwner,
                    "Only the owner can create auctions"
                );

                // Check permission
                require(
                    nftRegistry.getApproved(assetId) == address(this) ||
                        nftRegistry.isApprovedForAll(assetOwner, address(this)),
                    "The contract is not authorized to manage the asset"
                );

                // Check if there isn't any running auction
                require(
                    nftAuctions[nftAsset].runningAuction == bytes32(0),
                    "This asset has a running auction"
                );
            }
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

        // Update current running before actually create
        nftAuctions[nftAsset].runningAuction = auctionId;

        // Create new auction
        Auction storage auction = nftAuctions[nftAsset].auctions[auctionId];
        {
            uint256 biddingEnd = block.timestamp.add(biddingTime);
            uint256 revealEnd = biddingEnd.add(revealTime);
            auction.id = auctionId;
            auction.seller = assetOwner;
            auction.biddingEnd = biddingEnd;
            auction.revealEnd = revealEnd;
            auction.startPrice = startPriceInWei;
        }

        emit AuctionCreated(
            assetOwner,
            nftAddress,
            auctionId,
            assetId,
            auction.biddingEnd,
            auction.revealEnd,
            startPriceInWei
        );
    }

    function _cancelAuction(bytes32 nftAsset, bytes32 auctionId) internal {
        address sender = _msgSender();

        Auction storage auction = nftAuctions[nftAsset].auctions[auctionId];
        require(auction.id != 0, "Asset not published");
        require(
            auction.seller == sender || sender == owner(),
            "Unauthorized user"
        );

        nftAuctions[nftAsset].runningAuction = bytes32(0);
        emit AuctionCancelled(auctionId);
    }

    /// Place a blinded bid with `_blindedBid` =
    /// keccak256(abi.encodePacked(value, fake, secret)).
    /// The sent ether is only refunded if the bid is correctly
    /// revealed in the revealing phase. The bid is valid if the
    /// ether sent together with the bid is at least "value" and
    /// "fake" is not true. Setting "fake" to true and sending
    /// not the exact amount are ways to hide the real bid but
    /// still make the required deposit. The same address can
    /// place multiple bids.
    function _bidAuction(
        bytes32 nftAsset,
        bytes32 auctionId,
        bytes32 _blindedBid,
        uint256 deposit
    ) internal {
        Auction storage _auction = nftAuctions[nftAsset].auctions[auctionId];
        onlyBefore(_auction.biddingEnd);
        address sender = _msgSender();

        require(
            acceptedToken.transferFrom(sender, address(this), deposit),
            "Deposit the bid to the Marketplace contract failed"
        );
        require(
            _auction.bids[sender][_blindedBid] != 0,
            "You have already bidded this value"
        );

        _auction.bids[sender][_blindedBid] = deposit;
        emit BidSuccessful(sender, auctionId, _blindedBid);
    }

    function _placeBid(
        Auction storage _auction,
        address bidder,
        uint256 value
    ) internal returns (bool success) {
        if (value <= _auction.startPrice || value <= _auction.highestBid) {
            return false;
        }
        if (_auction.highestBidder != address(0)) {
            // Refund the previously highest bidder.
            _auction.pendingReturns[_auction.highestBidder] += _auction
                .highestBid;
        }
        _auction.highestBid = value;
        _auction.highestBidder = bidder;
        return true;
    }

    function _revealBid(
        bytes32 nftAsset,
        bytes32 auctionId,
        uint256[] memory _values,
        bool[] memory _fake,
        bytes32[] memory _secret
    ) internal {
        Auction storage _auction = nftAuctions[nftAsset].auctions[auctionId];
        onlyAfter(_auction.biddingEnd);
        onlyBefore(_auction.revealEnd);
        address sender = _msgSender();
        uint256 length = _values.length;
        require(_fake.length == length);
        require(_secret.length == length);

        uint256 refund;
        for (uint256 i = 0; i < length; i++) {
            (uint256 value, bool fake, bytes32 secret) = (
                _values[i],
                _fake[i],
                _secret[i]
            );
            bytes32 blindedBid = keccak256(
                abi.encodePacked(value, fake, secret)
            );
            uint256 deposit = _auction.bids[sender][blindedBid];
            if (deposit == 0) {
                // Incorrect bid parameter
                continue;
            }

            // Make it impossible for the sender to re-claim
            // the same deposit.
            _auction.bids[sender][blindedBid] = 0;

            // Add to refund
            refund += deposit;

            // If this is the highest bid then save it
            if (
                !fake && deposit >= value && _placeBid(_auction, sender, value)
            ) {
                refund -= value;
            }

            emit RevealSuccessful(fake, auctionId, value);
        }
        if (refund > 0) {
            acceptedToken.transfer(sender, refund);
            emit AuctionRefund(sender, auctionId, refund);
        }
    }

    function _withdraw(bytes32 nftAsset, bytes32 auctionId) internal {
        address sender = _msgSender();
        Auction storage _auction = nftAuctions[nftAsset].auctions[auctionId];
        uint256 amount = _auction.pendingReturns[sender];
        if (amount > 0) {
            // It is important to set this to zero because the recipient
            // can call this function again as part of the receiving call
            // before `transfer` returns (see the remark above about
            // conditions -> effects -> interaction).
            _auction.pendingReturns[sender] = 0;
            acceptedToken.transfer(sender, amount);
            emit AuctionRefund(sender, auctionId, amount);
        }
    }

    function _auctionEnd(
        address nftAddress,
        uint256 assetId,
        bytes32 nftAsset,
        bytes32 auctionId
    )
        internal
        auctionExisted(nftAsset, auctionId)
        auctionIsRunning(nftAsset, auctionId)
    {
        Auction storage _auction = nftAuctions[nftAsset].auctions[auctionId];
        onlyAfter(_auction.revealEnd);

        address auctionHighestBidder = _auction.highestBidder;
        address seller = _auction.seller;
        uint256 auctionHighestBid = _auction.highestBid;

        require(_auction.id != 0, "Asset not published");
        require(seller != address(0), "Invalid seller address");
        require(
            auctionHighestBidder != address(0),
            "Invalid highest bidder address"
        );
        IERC721 nftRegistry = IERC721(nftAddress);
        require(
            seller == nftRegistry.ownerOf(assetId),
            "The seller is no longer the owner"
        );
        nftAuctions[nftAsset].runningAuction = bytes32(0);

        uint256 saleShareAmount = 0;

        if (ownerCutPerMillion > 0) {
            // Calculate sale share
            saleShareAmount = auctionHighestBid.mul(ownerCutPerMillion).div(
                1000000
            );

            // Transfer share amount for marketplace Owner
            require(
                acceptedToken.transfer(owner(), saleShareAmount),
                "Transfering the cut to the Marketplace owner failed"
            );
        }

        // Transfer sale amount to seller
        require(
            acceptedToken.transfer(
                seller,
                auctionHighestBid.sub(saleShareAmount)
            ),
            "Transfering the sale amount to the seller failed"
        );

        // Transfer asset owner
        nftRegistry.safeTransferFrom(seller, auctionHighestBidder, assetId);

        emit AuctionSuccessful(
            seller,
            auctionHighestBidder,
            auctionId,
            auctionHighestBid
        );
    }
}
