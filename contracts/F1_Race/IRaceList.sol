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
        uint256 registerAt,
        uint256 betStarted,
        uint256 betEnded
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
    function updateRaceResult(
        bytes32 id
    ) external;

    function raceIsExisted(bytes32 raceId)
        external
        view
        returns (bool);

    event RaceCreated(
        bytes32 id,
        uint256 slots,
        uint256 registerAt,
        uint256 betStarted,
        uint256 betEnded
    );

    error TooEarly(uint256 time);
}
