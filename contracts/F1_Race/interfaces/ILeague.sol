// SPDX-License-Identifier: MIT
pragma solidity 0.8.4;

import "@openzeppelin/contracts/utils/introspection/IERC165.sol";

interface ILeague is IERC165 {

    struct Race {
        uint8 noSlot; // Max 26
        uint32 startAt; // Max 4294967295 -> Sunday, 7 February 2106 06:28:15
        bytes27 result; // 27 bytes - Remove first bytes
    }

    struct LeagueInfo {
        uint8 totalRace; // Max 255 races
        uint8 createdRace; // Max totalRace
        uint8 endedRace; // Max totalRace
        string leagueName;
    }

    /**
     * @dev Create new auction
     */
    function createRace(
        uint8 noSlot,
        uint32 startAt
    ) external returns (bytes32);

    /**
     * @dev Get race info by id.
     */
    function raceInfo(bytes32 raceId)
        external
        view
        returns (Race memory);

    /**
     * @dev Check registered slot.
     */
    function registeredSlot(address register, bytes32 raceId) 
        external 
        view 
    returns (uint8);

    /**
     * @dev Check total score of this league.
     */
    function getTotalScore(address register) 
        external 
        view 
        returns (uint8);

    /**
     * @dev Get league info.
     */
    function leagueInfo() 
        external 
        view 
        returns (LeagueInfo memory);

    /**
     * @dev Cancel race by id.
     */
    function cancelRace(bytes32 raceId) external;

    /**
     * @dev Register race by id.
     */
    function registerRace(uint8 slotId, bytes32 raceId) external;

    /**
     * @dev Update race result.
     */
    function updateRaceResult(
        bytes32 raceId,
        bytes27 result
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
    function receiveReward(uint8 slotId, bytes32 raceId) external;

    event RaceCreated(
        uint8 noSlot,
        uint8 raceNo,
        uint32 startAt,
        bytes32 id
    );
    event RaceResultUpdated(bytes32 id, bytes27 result);
    event RaceCancelled(bytes32 id);
    event RewardAdded(bytes32 raceId, uint256 nftRewardId, uint256 winnerIndex);
    event RewardReceived(uint8 slotId, bytes32 raceId, uint256 nftRewardId);
    event RewardRemoved(bytes32 raceId, uint256 nftRewardId, uint256 winnerIndex);
    event Registered(uint8 slotId, address participant, bytes32 raceId);

    error TooLate(uint256 time);
    error TooEarly(uint256 time);
    error InvalidContract();
    error InvalidSlot();
    error RaceNotExisted();
    error RaceExisted();
    error RewardIsNotExisted();
    error RewardIsExisted();
    error InvalidSender();
    error CannotCreateMoreRace(uint256 _currentRaceNo, uint256 totalRace);
    error AlreadyRegistered();
    error InvalidRegister();
    error NotEndYet();
    error RaceWasUpdated();
}
