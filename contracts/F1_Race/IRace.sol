// SPDX-License-Identifier: MIT
pragma solidity 0.8.4;

interface IRace {
    struct Race {
        // Race commission percentage
        uint256 commission;
        // Reward rate for winner bet
        uint256 rewardRate;
        // Number of slots
        uint256 slots;
        // Bething start time
        uint256 betStarted;
        // Bething end time
        uint256 betEnded;
        // Race result
        bytes32 result;
    }
}
