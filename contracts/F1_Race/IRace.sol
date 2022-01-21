// SPDX-License-Identifier: MIT
pragma solidity 0.8.4;

interface IRace {
    struct Race {
        // Number of slots
        uint256 slots;
        // Registration start time
        uint256 registerAt;
        // Bething start time
        uint256 betStarted;
        // Bething end time
        uint256 betEnded;
        // Race result
        bytes32 result;
    }
}