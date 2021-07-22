// SPDX-License-Identifier: MIT
pragma solidity 0.8.4;

import "@openzeppelin/contracts/utils/math/SafeMath.sol";
import "@openzeppelin/contracts/utils/Address.sol";

import "./MarketplaceStorage.sol";
import "@openzeppelin/contracts/access/Ownable.sol";
import "@openzeppelin/contracts/security/Pausable.sol";
import "../BEP20/IBEP20.sol";
import "@openzeppelin/contracts/token/ERC721/IERC721.sol";
import "@openzeppelin/contracts/utils/cryptography/ECDSA.sol";
import "@openzeppelin/contracts/utils/cryptography/draft-EIP712.sol";

contract Marketplace is
    Ownable,
    Pausable,
    MarketplaceStorage,
    EIP712("KFIVE Marketplace", "1")
{
    using SafeMath for uint256;
    using Address for address;

    /**
     * @dev Initialize this contract. Acts as a constructor
     * @param _acceptedToken - Address of the ERC20 accepted for this marketplace
     * @param _ownerCutPerMillion - owner cut per million
     */
    constructor(address _acceptedToken, uint256 _ownerCutPerMillion) {
        // Fee init
        setOwnerCutPerMillion(_ownerCutPerMillion);

        transferOwnership(_msgSender());

        require(
            _acceptedToken.isContract(),
            "The accepted token address must be a deployed contract"
        );
        acceptedToken = IBEP20(_acceptedToken);
    }

    /**
     * @dev Sets the publication fee that's charged to users to publish items
     * @param _publicationFee - Fee amount in wei this contract charges to publish an item
     */
    function setPublicationFee(uint256 _publicationFee) external onlyOwner {
        publicationFeeInWei = _publicationFee;
        emit ChangedPublicationFee(publicationFeeInWei);
    }

    /**
     * @dev Sets the share cut for the owner of the contract that's
     *  charged to the seller on a successful sale
     * @param _ownerCutPerMillion - Share amount, from 0 to 999,999
     */
    function setOwnerCutPerMillion(uint256 _ownerCutPerMillion)
        public
        onlyOwner
    {
        require(
            _ownerCutPerMillion < 1000000,
            "The owner cut should be between 0 and 999,999"
        );

        ownerCutPerMillion = _ownerCutPerMillion;
        emit ChangedOwnerCutPerMillion(ownerCutPerMillion);
    }

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
     * @dev Cancel an already published auction
     *  can only be canceled by seller or the contract owner
     * @param nftAddress - Address of the NFT registry
     * @param assetId - ID of the published NFT
     * @param auctionId - ID of the auction
     */
    function cancelAuction(
        address nftAddress,
        uint256 assetId,
        bytes32 auctionId
    ) external whenNotPaused {
        _cancelAuction(nftAddress, assetId, auctionId);
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

    /**
     * @dev Bid the auction for a published NFT
     * @param nftAddress - Address of the NFT registry
     * @param assetId - ID of the published NFT
     * @param auctionId - ID of the auction
     * @param blindedBid - Encoded Bid price
     * @param deposit - Deposit value (Must be large than bid price)
     */
    function bidAuction(
        address nftAddress,
        uint256 assetId,
        bytes32 auctionId,
        bytes32 blindedBid,
        uint256 deposit
    ) external _requireERC721(nftAddress) whenNotPaused {
        _bidAuction(nftAddress, assetId, auctionId, blindedBid, deposit);
    }

    /**
     * @dev Reveal your blinded bids. You will get a refund for all
     * correctly blinded invalid bids and for all bids except for
     * the totally highest.
     * @param nftAddress - Address of the NFT registry
     * @param assetId - ID of the published NFT
     * @param auctionId - ID of the auction
     * @param _values - Array of bid values
     * @param _fake - Array of true - false values, auction will ignore true value
     * @param _secret - Array of secret values used for verify the encoded bid values in bidding stage
     */
    function revealBid(
        address nftAddress,
        uint256 assetId,
        bytes32 auctionId,
        uint256[] memory _values,
        bool[] memory _fake,
        bytes32[] memory _secret
    ) external {
        Auction storage _auction = auctionByAssetId[nftAddress][assetId][
            auctionId
        ];
        onlyAfter(_auction.biddingEnd);
        onlyBefore(_auction.revealEnd);
        address sender = _msgSender();
        uint256 length = _auction.bids[sender].length;
        require(_values.length == length);
        require(_fake.length == length);
        require(_secret.length == length);

        uint256 refund;
        for (uint256 i = 0; i < length; i++) {
            Bid storage bidToCheck = _auction.bids[sender][i];
            (uint256 value, bool fake, bytes32 secret) = (
                _values[i],
                _fake[i],
                _secret[i]
            );
            if (
                bidToCheck.blindedBid !=
                keccak256(abi.encodePacked(value, fake, secret))
            ) {
                // Bid was not actually revealed.
                // Do not refund deposit.
                continue;
            }
            refund += bidToCheck.deposit;
            if (!fake && bidToCheck.deposit >= value) {
                if (_placeBid(_auction, sender, value)) refund -= value;
            }
            // Make it impossible for the sender to re-claim
            // the same deposit.
            bidToCheck.blindedBid = bytes32(0);
        }
        emit AuctionRefund(nftAddress, assetId, refund, sender);
        acceptedToken.transfer(sender, refund);
    }

    /**
     * @dev Withdraw a bid that was overbid.
     * @param nftAddress - Address of the NFT registry
     * @param assetId - ID of the published NFT
     * @param auctionId - ID of the auction
     */
    function withdraw(
        address nftAddress,
        uint256 assetId,
        bytes32 auctionId
    ) external {
        address sender = _msgSender();
        Auction storage _auction = auctionByAssetId[nftAddress][assetId][
            auctionId
        ];
        uint256 amount = _auction.pendingReturns[sender];
        if (amount > 0) {
            // It is important to set this to zero because the recipient
            // can call this function again as part of the receiving call
            // before `transfer` returns (see the remark above about
            // conditions -> effects -> interaction).
            _auction.pendingReturns[sender] = 0;
            acceptedToken.transfer(sender, amount);
        }
    }

    /**
     * @dev Withdraw a bid that was overbid.
     * @param nftAddress - Address of the NFT registry
     * @param assetId - ID of the published NFT
     * @param auctionId - ID of the auction
     */
    function pendingReturns(
        address nftAddress,
        uint256 assetId,
        bytes32 auctionId
    ) external view returns (uint256) {
        address sender = _msgSender();
        Auction storage _auction = auctionByAssetId[nftAddress][assetId][
            auctionId
        ];
        uint256 amount = _auction.pendingReturns[sender];
        return amount;
    }

    /**
     * @dev End the auction and send the highest bid to the beneficiary.
     * @param nftAddress - Address of the NFT registry
     * @param assetId - ID of the published NFT
     * @param auctionId - ID of the auction
     */
    function auctionEnd(
        address nftAddress,
        uint256 assetId,
        bytes32 auctionId
    ) external {
        Auction storage _auction = auctionByAssetId[nftAddress][assetId][
            auctionId
        ];
        if (_auction.ended) revert AuctionEndAlreadyCalled();
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
        IERC721 nftRegistry = IERC721(_auction.nftAddress);
        require(
            seller == nftRegistry.ownerOf(assetId),
            "The seller is no longer the owner"
        );
        _auction.ended = true;

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
            auctionId,
            assetId,
            seller,
            nftAddress,
            auctionHighestBid,
            auctionHighestBidder
        );
    }

    /**
     * @dev Get total bids.
     * @param bidder - Address of the bidder
     * @param nftAddress - Address of the NFT registry
     * @param assetId - ID of the published NFT
     * @param auctionId - ID of the auction
     */
    function totalBid(
        address bidder,
        address nftAddress,
        bytes32 auctionId,
        uint256 assetId
    ) external view returns (uint256) {
        Auction storage _auction = auctionByAssetId[nftAddress][assetId][
            auctionId
        ];
        return _auction.bids[bidder].length;
    }

    function _createAuction(
        address nftAddress,
        uint256 assetId,
        uint256 startPriceInWei,
        uint256 biddingTime,
        uint256 revealTime
    ) internal _requireERC721(nftAddress) {
        address assetOwner;
        {
            require(startPriceInWei > 0, "Price should be bigger than 0");
            address sender = _msgSender();
            {
                IERC721 nftRegistry = IERC721(nftAddress);
                assetOwner = nftRegistry.ownerOf(assetId);

                require(
                    sender == assetOwner,
                    "Only the owner can create auctions"
                );
                require(
                    nftRegistry.getApproved(assetId) == address(this) ||
                        nftRegistry.isApprovedForAll(assetOwner, address(this)),
                    "The contract is not authorized to manage the asset"
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
        Auction storage auction = auctionByAssetId[nftAddress][assetId][
            auctionId
        ];
        {
            uint256 biddingEnd = block.timestamp.add(biddingTime);
            uint256 revealEnd = biddingEnd.add(revealTime);
            auction.id = auctionId;
            auction.seller = assetOwner;
            auction.nftAddress = nftAddress;
            auction.biddingEnd = biddingEnd;
            auction.revealEnd = revealEnd;
            auction.startPrice = startPriceInWei;
            auction.ended = false;
        }

        emit AuctionCreated(
            auctionId,
            assetId,
            assetOwner,
            nftAddress,
            auction.biddingEnd,
            auction.revealEnd,
            startPriceInWei
        );
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

    function _cancelAuction(
        address nftAddress,
        uint256 assetId,
        bytes32 auctionId
    ) internal {
        address sender = _msgSender();
        Auction storage auction = auctionByAssetId[nftAddress][assetId][
            auctionId
        ];

        require(auction.id != 0, "Asset not published");
        require(
            auction.seller == sender || sender == owner(),
            "Unauthorized user"
        );

        address auctionSeller = auction.seller;
        address auctionNftAddress = auction.nftAddress;
        auction.ended = true;

        emit AuctionCancelled(
            auctionId,
            assetId,
            auctionSeller,
            auctionNftAddress
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

        uint256 saleShareAmount = 0;

        bytes32 orderId = order.id;
        delete orderByAssetId[nftAddress][assetId];

        if (ownerCutPerMillion > 0) {
            // Calculate sale share
            saleShareAmount = price.mul(ownerCutPerMillion).div(1000000);

            // Transfer share amount for marketplace Owner
            require(
                acceptedToken.transferFrom(sender, owner(), saleShareAmount),
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

    modifier _requireERC721(address nftAddress) {
        require(
            nftAddress.isContract(),
            "The NFT Address should be a contract"
        );
        IERC721 nftRegistry = IERC721(nftAddress);
        require(
            nftRegistry.supportsInterface(ERC721_Interface),
            "The NFT contract has an invalid ERC721 implementation"
        );
        _;
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
        address nftAddress,
        uint256 assetId,
        bytes32 auctionId,
        bytes32 _blindedBid,
        uint256 deposit
    ) internal {
        Auction storage _auction = auctionByAssetId[nftAddress][assetId][
            auctionId
        ];
        onlyBefore(_auction.biddingEnd);
        if (_auction.ended) revert AuctionEndAlreadyCalled();
        address sender = _msgSender();

        require(
            acceptedToken.transferFrom(sender, address(this), deposit),
            "Deposit the bid to the Marketplace contract failed"
        );

        _auction.bids[sender].push(
            Bid({blindedBid: _blindedBid, deposit: deposit})
        );
        emit BidSuccessful(sender, nftAddress, assetId, _blindedBid);
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

    function _msgSender() internal view override returns (address sender) {
        if (msg.sender == address(this)) {
            bytes memory array = msg.data;
            uint256 index = msg.data.length;
            assembly {
                // Load the 32 bytes word from memory with the address on the lower 20 bytes, and mask those.
                sender := and(
                    mload(add(array, index)),
                    0xffffffffffffffffffffffffffffffffffffffff
                )
            }
        } else {
            sender = msg.sender;
        }
        return sender;
    }
}
