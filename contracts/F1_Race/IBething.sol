// SPDX-License-Identifier: MIT
pragma solidity 0.8.4;

import "@openzeppelin/contracts/utils/introspection/IERC165.sol";

interface IBething is IERC165 {
    /**
     * @dev Bet a race slot
     */
    function bet(
        uint256 slotId,
        bytes32 raceId,
        uint256 betValue
    ) external;

    /**
     * @dev Fund for the race
     */
    function fundRace(
        bytes32 raceId,
        uint256 fundValue
    ) external;

    /**
     * @dev Claim winning reward from the ended race
     */
    function claim(
        bytes32 raceId
    ) external;

    /**
     * @dev Claim commission from ended race - Admin only
     */
    function claimCommission(
        bytes32 raceId,
        address receiver
    ) external;

    /**
     * @dev Get your total slot bet
     */
    function totalSlotBet(
        bytes32 raceId,
        uint256 slotId
    ) external returns (uint256);

    /**
     * @dev Get total race bet
     */
    function totalRaceBet(
        bytes32 raceId
    ) external returns (uint256);

    /**
     * @dev Get slot position after race ended
     */
    function getSlotPosition(
        bytes32 raceId,
        uint256 slotId
    ) external returns (uint256);

    /**
     * @dev Update race address
     */
    function updateRaceAddress(address raceList)
        external;

    /**
     * @dev Update accept token address
     */
    function updateAcceptTokenAddress(address acceptToken)
        external;

    event BetSuccessful(uint256 slotId, bytes32 raceId, uint256 betValue);
    event FundSuccessful(bytes32 raceId, uint256 fundValue);
    event ClaimSuccessful(bytes32 raceId, uint256 claimValue);
    event ClaimCommission(bytes32 raceId, uint256 claimValue, address receiver);
    event RaceListUpdated(address race);
    event AcceptTokenUpdated(address acceptToken);

    error TooEarly(uint256 time);
    error TooLate(uint256 time);
}
