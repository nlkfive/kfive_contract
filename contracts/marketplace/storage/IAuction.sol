// SPDX-License-Identifier: GPL-3.0
pragma solidity 0.8.4;

interface IAuction {
    ////////////////////////////////////////////////////////////
    // COMMON AUCTION ERRORS
    ////////////////////////////////////////////////////////////
    error InvalidBiddingEnd(); // 0x18c71d1ad
    error InvalidBiddingPrice(); // 0x7532083f
    error RewardGranted(); // 0x946ef902
    error TooLate(uint256); // 0x691e5682
    error TooEarly(uint256); // 0x2a35a324
    error NotRunning(); // 0x03b5e413
    error NotStarted(); // 0x4ea9379a
    ////////////////////////////////////////////////////////////
    // COMMON AUCTION EVENTS
    ////////////////////////////////////////////////////////////
    event AuctionRefundSuccessful(
        address bidder,
        bytes32 auctionId,
        uint256 value
    );
    event AuctionCancelledSuccessful(
        address canceller, 
        bytes32 auctionId
    );
    event GrantAuctionRewardSuccessful(
        address auctionHighestBidder,
        bytes32 auctionId,
        uint256 assetId
    );
}
