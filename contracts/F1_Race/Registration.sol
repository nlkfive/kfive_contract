// SPDX-License-Identifier: MIT
pragma solidity 0.8.4;

import "./IRegistrationList.sol";
import "../common/AccessControl.sol";
import "@openzeppelin/contracts/utils/introspection/ERC165.sol";
import "@openzeppelin/contracts/utils/Address.sol";
import "@openzeppelin/contracts/token/ERC721/extensions/IERC721Enumerable.sol";
import "@openzeppelin/contracts/token/ERC721/IERC721Receiver.sol";
import "./IRaceList.sol";
import "./IRace.sol";

import "@chainlink/contracts/src/v0.8/interfaces/VRFCoordinatorV2Interface.sol";
import "@chainlink/contracts/src/v0.8/VRFConsumerBaseV2.sol";

contract RegistrationList is
    VRFConsumerBaseV2,
    IRace,
    KfiveAccessControl,
    IRegistrationList,
    IERC721Receiver
{
    using Address for address;
    IRaceList private _raceList;
    IERC721Enumerable private _nlggt;
    IERC721Enumerable private _raceReward;
    VRFCoordinatorV2Interface private _vrfCoordinator;

    // The contract will contain multiple objects. 
    // Each oracle job has a unique key hash that identifies the tasks that it should perform. 
    // The contract will store the Key Hash that identifies Chainlink VRF and the fee amount to use in the request.
    bytes32 private s_keyHash;
    uint64 private s_subscriptionId;

    // The default is 3, but you can set this higher.
    uint16 private constant requestConfirmations = 3;

    // Depends on the number of requested values that you want sent to the
    // fulfillRandomWords() function. Storing each word costs about 20,000 gas,
    // so 100,000 is a safe default for this example contract. Test and adjust
    // this limit based on the network that you select, the size of the request,
    // and the processing of the callback request in the fulfillRandomWords()
    // function.
    uint32 private constant callbackGasLimit = 100000;

    // For this example, retrieve 2 random values in one request.
    // Cannot exceed VRFCoordinatorV2.MAX_NUM_WORDS.
    uint32 private constant numWords =  1;

    bytes32 private constant RANDOM_IN_PROGRESS = 0xffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff;


    // https://docs.chain.link/docs/vrf-contracts/
    // https://vrf.chain.link/?_ga=2.219199065.197734000.1648027927-202136979.1647858568

    // Binance Smart Chain Testnet
    // VRF Coordinator	0x6A2AAd07396B36Fe02a22b33cf443582f682c82f
    // Key Hash	0xd4bb89654db74673a187bd804519e65e3f71a52bc55f11da7601a13dcf505314

    // Binance Smart Chain Mainnet
    // VRF Coordinator	0xc587d9053cd1118f25F645F9E08BB98c9712A4EE
    // Key Hash	0x114f3da0a805b6a67d6e9cd2ec746f7028f1b7376365af575cfea3550dd1aa04
    
    constructor(
        address vrfCoordinator,
        address nlggt,
        address raceReward,
        address raceList,
        uint64 subscriptionId,
        bytes32 keyHash
    ) VRFConsumerBaseV2(vrfCoordinator) {
        if(!nlggt.isContract()) revert InvalidContract();
        if(!raceList.isContract()) revert InvalidContract();
        if(!raceReward.isContract()) revert InvalidContract();
        _raceList = IRaceList(raceList);
        _nlggt = IERC721Enumerable(nlggt);
        _raceReward = IERC721Enumerable(raceReward);
        _vrfCoordinator = VRFCoordinatorV2Interface(vrfCoordinator);
        s_keyHash = keyHash;
        s_subscriptionId = subscriptionId;
    }

    /**
     * @dev See {IERC165-supportsInterface}.
     */
    function supportsInterface(bytes4 interfaceId)
        public
        view
        virtual
        override(AccessControlEnumerable, IERC165)
        returns (bool)
    {
        return
            interfaceId == type(IRegistrationList).interfaceId ||
            super.supportsInterface(interfaceId);
    }

    ////////////////////////////////////////////////////////////
    // Participants list
    ////////////////////////////////////////////////////////////
    // From raceId => (slotId => (index => participant address))
    mapping(bytes32 => mapping(uint256 => mapping(uint8 => address))) private registrationList;
    // Track user only register one slot each race
    // From raceId => (slotId => (participant address => index))
    mapping(bytes32 => mapping(address => bool)) private registeredSlot;
    // Random number (32 bytes = 32 slots) - Using for select selected participants
    // byte index 1 => slot 1 (L2R)
    // byte index 2 => slot 2
    // ...
    // byte index 32 => slot 32
    // Selected participant at slot 1 = byte 1 modulo total registered in slot 1
    // From raceId => random number generated by chainlink
    mapping(bytes32 => bytes32) private randomSelected;
    // Count total registered (32 bytes = 32 slots) - Maximum each slot is 255
    // byte index 1 => slot 1 (L2R)
    // byte index 2 => slot 2
    // ...
    // byte index 32 => slot 32
    // From raceId => total registered bytes array
    mapping(bytes32 => bytes32) private totalRegistered;
    // From vrf request id to running raceId
    mapping(uint256 => bytes32) private vrfRandomMap;
    // From raceId => (race result index => NFT reward ID)
    mapping(bytes32 => mapping(bytes1 => uint256)) private raceRewards;

    /**
     * @dev Register to be a participant
     * Participant must be an nlggt holder
     */
    function register(
        uint256 slotId,
        bytes32 raceId
    )
        external 
        override
        whenNotPaused
    {
        if(_nlggt.balanceOf(_msgSender()) == 0) revert NotNLGGTHolder();
        Race memory _race = _raceList.getRace(raceId);
        if(_race.betStarted == 0) revert RaceNotExisted();
        if(slotId >= _race.slots) revert InvalidSlot();
        if(registeredSlot[raceId][_msgSender()]) revert AlreadyRegistered();

        onlyAfter(_race.betStarted);
        onlyBefore(_race.betEnded);
        uint8 totalRegisteredAtSlot = uint8(totalRegistered[raceId][slotId]);
        if (totalRegisteredAtSlot >= 255) revert MaximumReached();

        registeredSlot[raceId][_msgSender()] = true;
        totalRegistered[raceId] = increaseAtIndex(totalRegistered[raceId], slotId);
        registrationList[raceId][slotId][totalRegisteredAtSlot] = _msgSender();
        emit Registered(slotId, _msgSender(), raceId);
    }

    /**
     * @dev Select random participants for race from registered list for each slot
     */
    function selectParticipant(bytes32 raceId)
        external 
        override
        whenNotPaused
        onlyRole(ADMIN_ROLE) returns (uint256 requestId)
    {
        // checking participant at slot is already selected
        if(randomSelected[raceId] != 0) revert AlreadySelected();
        // check if race is existed
        Race memory _race = _raceList.getRace(raceId);
        if(_race.betStarted == 0) revert RaceNotExisted();
        onlyAfter(_race.betEnded);

        // storing raceId
        randomSelected[raceId] = RANDOM_IN_PROGRESS;
        // requesting randomness
        requestId = _vrfCoordinator.requestRandomWords(
            s_keyHash,
            s_subscriptionId,
            requestConfirmations,
            callbackGasLimit,
            numWords
        );
        // storing requestId
        vrfRandomMap[requestId] = raceId;

        emit RandomInProgress(raceId);
    }

    function fulfillRandomWords(uint256 requestId, uint256[] memory randomWords) internal override 
    {
        bytes32 raceId = vrfRandomMap[requestId];
        bytes32 randomness = bytes32(randomWords[0]);
        // checking participant at slot is already selected
        if(randomSelected[raceId] != 0 && randomSelected[raceId] != RANDOM_IN_PROGRESS) revert AlreadySelected();
        // Mark this race has selected participants
        randomSelected[raceId] = modBytes(totalRegistered[raceId], randomness);
        randomSelected[raceId] |= bytes32(uint256(0xff));
        emit ParticipantsSelected(requestId, raceId, randomness, randomSelected[raceId]);
    }

    /**
     * @dev Add reward.
     */
    function addReward(bytes32 raceId, uint256 nftRewardId, bytes1 resultIndex) 
        external 
        override
        whenNotPaused
        onlyRole(ADMIN_ROLE)
    {
        Race memory _race = _raceList.getRace(raceId);
        if (_race.betStarted == 0) revert RaceNotExisted();

        if (raceRewards[raceId][resultIndex] != 0){
            revert RewardIsExisted();
        }
        _raceReward.safeTransferFrom(
            _msgSender(),
            address(this),
            nftRewardId
        );
        raceRewards[raceId][resultIndex] = nftRewardId;
        emit RewardAdded(raceId, nftRewardId, resultIndex);
    }

    /**
     * @dev Receive reward after race ended.
     */
    function receiveReward(bytes32 raceId, uint256 slotId) 
        external 
        override
        whenNotPaused
    {
        address slotWinner = selectedParticipant(raceId, slotId);
        if(slotWinner != _msgSender()) {
            revert InvalidSender();
        }

        Race memory _race = _raceList.getRace(raceId);
        if (_race.betStarted == 0) revert RaceNotExisted();
        onlyAfter(_race.betEnded);

        uint256 nftRewardId = raceRewards[raceId][_race.result[slotId]];
        if (nftRewardId == 0){
            revert RewardNotExistedOrReceived();
        }
        raceRewards[raceId][_race.result[slotId]] = 0;
        _raceReward.safeTransferFrom(
            address(this),
            slotWinner,
            nftRewardId
        );
        emit RewardReceived(raceId, slotId, nftRewardId);
    }

    /**
     * @dev Update race address
     */
    function updateRaceAddress(address raceList)
        external 
        override
        onlyRole(ADMIN_ROLE)
        whenNotPaused
    {
        _raceList = IRaceList(raceList);
        emit RaceListUpdated(raceList);
    }

    /**
     * @dev Update nlggt address
     */
    function updateNlggtAddress(address nlggt)
        external
        override
        onlyRole(ADMIN_ROLE)
        whenNotPaused
    {
        _nlggt = IERC721Enumerable(nlggt);
        emit NlggtUpdated(nlggt);
    }

    function onlyBefore(uint256 _time) internal view {
        if (block.timestamp >= _time) revert TooLate(_time);
    }

    function onlyAfter(uint256 _time) internal view {
        if (block.timestamp <= _time) revert TooEarly(_time);
    }

    // 0 < index < 31
    function increaseAtIndex(bytes32 input, uint256 index) internal pure returns (bytes32) {
        return bytes32(uint256(input) + (1 << ((31 - index) * 8)));
    }

    // Random select index from participant bytes array - Gas = 35000
    function modBytes(bytes32 _totalRegistered, bytes32 random) internal pure returns (bytes32 result) {
        for (uint256 slotId = 0; slotId < 32; slotId++) {
            if (_totalRegistered[slotId] != bytes1(0)){
            result |= bytes32(
                uint256(
                    uint8(random[slotId]) % uint8(_totalRegistered[slotId])
                )
            ) << (8 * (31 - slotId));
            }
        }
    }

    // Gas: 33619
    function selectedParticipants(bytes32 raceId) 
        external 
        view
        returns (bytes32 result)
    {
        if (randomSelected[raceId] == bytes32(0) || randomSelected[raceId] == RANDOM_IN_PROGRESS){
            result = bytes32(0);
        } else {
            for (uint256 slotId = 0; slotId < 32; slotId++) {
                if (totalRegistered[raceId][slotId] == bytes1(0)){
                continue;
                }
                result |= bytes32(
                (uint256(uint8(randomSelected[raceId][slotId])) + 1) 
                << (8 * (31 - slotId))
                );
            }
        }
    }

    function selectedParticipant(bytes32 raceId, uint256 slotId) 
        public 
        view
        returns (address)
    {
        return registrationList[raceId][slotId][uint8(randomSelected[raceId][slotId])];
    }

    function totalRegister(bytes32 raceId) 
        external 
        view
        returns (bytes32)
    {
        return totalRegistered[raceId];
    }

    function onERC721Received(
        address,
        address,
        uint256,
        bytes calldata
    ) external pure override returns (bytes4) {
        return
            bytes4(
                keccak256("onERC721Received(address,address,uint256,bytes)")
            );
    }
}
