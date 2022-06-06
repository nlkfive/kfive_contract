// SPDX-License-Identifier: MIT
pragma solidity 0.8.4;

import "@openzeppelin/contracts/utils/introspection/IERC165.sol";

interface ITournament is IERC165 {

    struct Race {
        uint8 noSlot; // Max 26
        uint32 startAt; // Max 4294967295 -> Sunday, 7 February 2106 06:28:15
        bytes27 result; // 27 bytes - Remove first bytes
    }

    struct TournamentInfo {
        uint8 totalRace; // Max 255 races
        uint8 createdRace; // Max totalRace
        uint8 endedRace; // Max totalRace
        string tournamentName;
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
     * @dev Check total score of this tournament.
     */
    function getTotalScore(address register) 
        external 
        view 
        returns (uint8);

    /**
     * @dev Get tournament info.
     */
    function tournamentInfo() 
        external 
        view 
        returns (TournamentInfo memory);

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
     * @dev Grand reward after league ended.
     */
    function grandReward(uint8 winnerIndex, address winner, uint256 nftRewardId, string memory tokenURI) external;

    event RaceCreated(
        uint8 noSlot,
        uint8 raceNo,
        uint32 startAt,
        bytes32 id
    );
    event RaceResultUpdated(bytes32 id, bytes27 result);
    event RaceCancelled(bytes32 id);
    event RewardGranted(uint8 winnerIndex, uint8 score, address winner, uint256 nftRewardId);
    event Registered(uint8 slotId, address participant, bytes32 raceId);

    error TooLate(uint256 time); //691e5682
    error TooEarly(uint256 time); //2a35a324
    error InvalidContract(); //6eefed20
    error InvalidSlot(); //1258e4438
    error RaceNotExisted(); //e2a9931a
    error RaceExisted(); //4868480f
    error InvalidSender(); //ddb5de5e
    error CannotCreateMoreRace(uint256 _currentRaceNo, uint256 totalRace);//df15abb8
    error AlreadyRegistered(); //3a81d6fc
    error InvalidRegister(); //14ddd536
    error NotEndYet(); //e0219957
    error RaceWasUpdated(); //6404a619
}
