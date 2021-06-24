// SPDX-License-Identifier: GPL-3.0
pragma solidity ^0.8.4;

contract BlindAuction {
    struct Bid {
        bytes32 blindedBid;
        uint256 deposit;
    }

    // Errors that describe failures.

    /// The function has been called too early.
    /// Try again at `time`.
    error TooEarly(uint256 time);
    /// The function has been called too late.
    /// It cannot be called after `time`.
    error TooLate(uint256 time);
    /// The function auctionEnd has already been called.
    error AuctionEndAlreadyCalled();

    event BidSuccessful(
        address indexed bidder,
        address nftAddress,
        uint256 indexed assetId,
        bytes32 blindedBid
    );

    event AuctionRefund(
        address nftAddress,
        uint256 indexed assetId,
        uint256 deposit,
        address indexed bidder
    );

    // Modifiers are a convenient way to validate inputs to
    // functions. `onlyBefore` is applied to `bid` below:
    // The new function body is the modifier's body where
    // `_` is replaced by the old function body.
    function onlyBefore(uint256 _time) internal view {
        if (block.timestamp >= _time) revert TooLate(_time);
    }

    function onlyAfter(uint256 _time) internal view {
        if (block.timestamp <= _time) revert TooEarly(_time);
    }
}
