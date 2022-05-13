// SPDX-License-Identifier: MIT
pragma solidity 0.8.4;

import "@openzeppelin/contracts/utils/introspection/IERC165.sol";

interface ILeague is IERC165 {
    /**
     * @dev Create new auction
     */
    function createRace(
        uint256 slots,
        uint256 startAt
    ) external returns (bytes32);

    /**
     * @dev Get race slots by id.
     */
    function raceSlots(bytes32 raceId)
        external
        view
        returns (uint256 slots);

    /**
     * @dev Get race start at by id.
     */
    function raceStartAt(bytes32 raceId)
        external
        view
        returns (uint256 startAt);

    /**
     * @dev Get race result by id.
     */
    function raceResult(bytes32 raceId)
        external
        view
        returns (bytes32 result);
        
    /**
     * @dev Get league name.
     */
    function leagueName() 
        external 
        view 
        returns (string memory);

    /**
     * @dev Cancel race by id.
     */
    function cancelRace(bytes32 id) external;

    /**
     * @dev Update race result.
     */
    function updateRaceResult(
        bytes32 raceId,
        bytes32 result
    ) external;

    event RaceCreated(
        bytes32 id,
        uint256 slots,
        uint256 startAt
    );
    event RaceResultUpdated(bytes32 id, bytes32 result);
    event RaceCancelled(bytes32 id);

    error TooLate(uint256 time);
    error TooEarly(uint256 time);
    error InvalidSlot();
    error RaceNotExisted();
    error RaceExisted();
    error CannotCancel();
}
