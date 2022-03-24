// SPDX-License-Identifier: MIT
pragma solidity 0.8.4;

import "@openzeppelin/contracts/utils/introspection/IERC165.sol";

interface IRegistrationList is IERC165 {
    /**
     * @dev Register to be a participant
     */
    function register(
        uint256 slotId,
        bytes32 raceId
    ) external;

    /**
     * @dev Add reward.
     */
    function addReward(bytes32 raceId, uint256 nftRewardId, bytes1 resultIndex) 
        external;

    /**
     * @dev Select random participant for race into slot from registered list for each slot
     */
    function selectParticipant(bytes32 raceId)
        external returns (uint256 requestId);

    /**
     * @dev Update race address
     */
    function updateRaceAddress(address raceList)
        external;

    /**
     * @dev Update nlggt address
     */
    function updateNlggtAddress(address nlggt)
        external;

    /**
     * @dev Receive reward after race ended.
     */
    function receiveReward(bytes32 raceId, uint256 slotId) external;

    event Registered(
        uint256 slotId,
        address participant,
        bytes32 raceId
    );

    event RandomInProgress(
        bytes32 raceId
    );
    event ParticipantsSelected(
        uint256 requestId,
        bytes32 raceId,
        bytes32 randomness,
        bytes32 selected
    );

    event RaceListUpdated(address race);
    event NlggtUpdated(address nlggt);
    event RewardAdded(bytes32 raceId, uint256 nftRewardId, bytes1 resultIndex);
    event RewardReceived(bytes32 raceId, uint256 slotId, uint256 nftRewardId);

    error TooEarly(uint256 time);
    error TooLate(uint256 time);
    error NotNLGGTHolder();
    error RaceNotExisted();
    error InvalidSlot();
    error MaximumReached();
    error InvalidContract();
    error AlreadySelected();
    error AlreadyRegistered();
    error InvalidSender();
    error RewardNotExistedOrReceived();
    error RewardIsExisted();
}
