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
     * @dev Get race info by id.
     */
    function raceInfo(bytes32 raceId)
        external
        view
        returns (uint256 slots, uint256 startAt, bytes32 result);

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
    function cancelRace(bytes32 raceId) external;

    /**
     * @dev Register race by id.
     */
    function registerRace(bytes32 raceId, uint256 slotId) external;

    /**
     * @dev Update race result.
     */
    function updateRaceResult(
        bytes32 raceId,
        bytes32 result
    ) external;

    /**
     * @dev Add reward by transfer.
     */
    function addRewardByTransfer(bytes32 raceId, uint256 nftRewardId, uint256 winnerIndex) 
        external;

     /**
     * @dev Add reward by mint.
     */
    function addRewardByMint(bytes32 raceId, uint256 nftRewardId, uint256 winnerIndex, string memory tokenURI) 
        external;

    /**
     * @dev Remove added reward.
     */
    function removeReward(
        bytes32 raceId, uint256 winnerIndex
    ) external;

    /**
     * @dev Receive reward after race ended.
     */
    function receiveReward(bytes32 raceId, uint256 slotId) external;

    event RaceCreated(
        bytes32 id,
        uint256 slots,
        uint256 startAt,
        uint256 raceNo
    );
    event RaceResultUpdated(bytes32 id, bytes32 result);
    event RaceCancelled(bytes32 id);
    event RewardAdded(bytes32 raceId, uint256 nftRewardId, uint256 winnerIndex);
    event RewardReceived(bytes32 raceId, uint256 slotId, uint256 nftRewardId);
    event RewardRemoved(bytes32 raceId, uint256 nftRewardId, uint256 winnerIndex);
    event Registered(uint256 slotId, address participant, bytes32 raceId);

    error TooLate(uint256 time);
    error TooEarly(uint256 time);
    error InvalidContract();
    error InvalidSlot();
    error RaceNotExisted();
    error RaceExisted();
    error CannotCancel();
    error RewardIsNotExisted();
    error RewardIsExisted();
    error InvalidSender();
    error CannotCreateMoreRace(uint256 _currentRaceNo, uint256 totalRace);
    error AlreadyRegistered();
    error InvalidRegister();
}
