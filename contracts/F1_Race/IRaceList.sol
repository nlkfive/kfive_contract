// SPDX-License-Identifier: MIT
pragma solidity 0.8.4;

import "@openzeppelin/contracts/utils/introspection/IERC165.sol";
import "./IRace.sol";

interface IRaceList is IRace, IERC165 {
    /**
     * @dev Create new auction
     */
    function createRace(
        uint256 slots,
        uint256 betStarted,
        uint256 betEnded,
        uint256 commission,
        uint256 rewardRate
    ) external;

    /**
     * @dev Get race by id.
     */
    function getRace(bytes32 id)
        external
        view
        returns (Race memory race);

    /**
     * @dev Cancel race by id.
     */
    function cancelRace(bytes32 id) external;

    /**
     * @dev Update race result.
     */
    function updateResult(
        bytes32 id
    ) external;

    /**
     * @dev Update commission.
     */
    function updateCommission(
        bytes32 id,
        uint256 commission
    ) external;

    /**
     * @dev Update rewardRate.
     */
    function updateRewardRate(
        bytes32 id,
        uint256 rewardRate
    ) external;

    function raceIsExisted(bytes32 raceId)
        external
        view
        returns (bool);

    function raceResult(bytes32 raceId)
        external
        view
        returns (bytes32);

    event RaceCreated(
        bytes32 id,
        uint256 slots,
        uint256 betStarted,
        uint256 betEnded,
        uint256 commission,
        uint256 rewardRate
    );
    event RaceResultUpdated(bytes32 id, bytes32 result);
    event RaceCancelled(bytes32 id);
    event RaceCommissionUpdated(bytes32 id, uint256 commission);
    event RaceRewardRateUpdated(bytes32 id, uint256 rewardRate);

    error TooLate(uint256 time);
    error TooEarly(uint256 time);
}
