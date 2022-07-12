// SPDX-License-Identifier: GPL-3.0
pragma solidity 0.8.4;

interface IAuction {
    ////////////////////////////////////////////////////////////
    // ERRORS
    ////////////////////////////////////////////////////////////
    error AuctionCancelled();
    error InvalidBiddingEnd();
    error InvalidBiddingPrice();
    error AuctionRefund();
    error NotWinner();
    error RewardGranted();
    error TooLate(uint256);
    error TooEarly(uint256);
    error NotRunning();
    error InvalidWithdraw();
    error AuctionEnded();
    ////////////////////////////////////////////////////////////
    // EVENTS
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
