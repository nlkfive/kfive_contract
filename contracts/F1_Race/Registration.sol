// SPDX-License-Identifier: MIT
pragma solidity 0.8.4;

import "./interfaces/IRegistration.sol";
import "../common/AccessControl.sol";
import "@openzeppelin/contracts/utils/introspection/ERC165.sol";
import "@openzeppelin/contracts/utils/Address.sol";
import "@openzeppelin/contracts/token/ERC721/extensions/IERC721Enumerable.sol";
import "@openzeppelin/contracts/token/ERC721/IERC721Receiver.sol";
import "./interfaces/ILeague.sol";
import "../NFT/IKfiveNFT.sol";
import "@openzeppelin/contracts/utils/math/SafeMath.sol";

import "@chainlink/contracts/src/v0.8/interfaces/VRFCoordinatorV2Interface.sol";
import "@chainlink/contracts/src/v0.8/VRFConsumerBaseV2.sol";

contract Registration is
    VRFConsumerBaseV2,
    KfiveAccessControl,
    IRegistration,
    IERC721Receiver
{
    using SafeMath for uint256;
    using Address for address;
    IKfiveNFT private _raceReward;
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
        address raceReward,
        uint64 subscriptionId,
        bytes32 keyHash
    ) VRFConsumerBaseV2(vrfCoordinator) {
        if(!raceReward.isContract()) revert InvalidContract();
        _raceReward = IKfiveNFT(raceReward);
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
            interfaceId == type(IRegistration).interfaceId ||
            super.supportsInterface(interfaceId);
    }

    // RaceID -> Randomness
    mapping(bytes32 => address) private leagueRace;
    // RaceId -> slotId -> list registers
    mapping(bytes32 => mapping(uint256 => address[])) private registrationList;
    // Raceid -> User Address -> isRegistered
    mapping(bytes32 => mapping(address => bool)) private registeredSlot;
    // RaceID -> Randomness
    mapping(bytes32 => uint256) private randomSelected;
    // VRFrequestID -> raceId
    mapping(uint256 => bytes32) private vrfRandomMap;
    // RaceId => RaceIndex => NFTID
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
        (uint256 _startAt, uint256 _slots,) = checkRaceExisted(raceId);
        onlyBefore(_startAt);
        if(slotId >= _slots) revert InvalidSlot();

        if(registeredSlot[raceId][_msgSender()]) revert AlreadyRegistered();

        registeredSlot[raceId][_msgSender()] = true;
        registrationList[raceId][slotId].push(_msgSender());
        
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
        (uint256 _startAt,,) = checkRaceExisted(raceId);
        onlyAfter(_startAt);

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
        uint256 randomness = randomWords[0];
        // checking participant at slot is already selected
        if(randomSelected[raceId] != 0) revert AlreadySelected();
        // Mark this race has selected participants
        randomSelected[raceId] = randomness;
        emit ParticipantsSelected(requestId, raceId, randomness);
    }

    /**
     * @dev Enable registration.
     */
    function enableRegistration(address leagueAddr, bytes32 raceId) 
        external
        override
        whenNotPaused
        onlyRole(ADMIN_ROLE)
    {
        if(!leagueAddr.isContract()) revert InvalidContract();
        ILeague _leagueAddr = ILeague(leagueAddr);
        uint256 _startAt = _leagueAddr.raceStartAt(raceId);
        if(_startAt == 0) revert RaceNotExisted();
        leagueRace[raceId] = leagueAddr;
    }

    /**
     * @dev Add reward.
     */
    function addRewardByTransfer(bytes32 raceId, uint256 nftRewardId, bytes1 resultIndex) 
        external 
        override
        whenNotPaused
        onlyRole(ADMIN_ROLE)
    {
        checkRaceExisted(raceId);

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
     * @dev Add reward.
     */
    function addRewardByMint(bytes32 raceId, uint256 nftRewardId, bytes1 resultIndex, string memory tokenURI) 
        external 
        override
        whenNotPaused
        onlyRole(ADMIN_ROLE)
    {
        checkRaceExisted(raceId);

        if (raceRewards[raceId][resultIndex] != 0){
            revert RewardIsExisted();
        }
        _raceReward.mint(
            address(this),
            nftRewardId,
            tokenURI
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

        (uint256 _startAt,, bytes32 _result) = checkRaceExisted(raceId);
        onlyAfter(_startAt);

        uint256 nftRewardId = raceRewards[raceId][_result[slotId]];
        if (nftRewardId == 0){
            revert RewardIsNotExistedOrReceived();
        }
        raceRewards[raceId][_result[slotId]] = 0;
        _raceReward.safeTransferFrom(
            address(this),
            slotWinner,
            nftRewardId
        );
        emit RewardReceived(raceId, slotId, nftRewardId);
    }

    function onlyBefore(uint256 _time) internal view {
        if (block.timestamp >= _time) revert TooLate(_time);
    }

    function onlyAfter(uint256 _time) internal view {
        if (block.timestamp <= _time) revert TooEarly(_time);
    }

    function selectedParticipants(bytes32 raceId) 
        external 
        view
        returns (address[] memory)
    {
        (, uint256 _slots,) = checkRaceExisted(raceId);
        address[] memory results = new address[](_slots);
        if (randomSelected[raceId] != 0){
            for (uint256 slotId = 0; slotId < _slots; slotId++) {
                results[slotId] = selectedParticipant(raceId, slotId);
            }
        }
        return results;
    }

    function selectedParticipant(bytes32 raceId, uint256 slotId) 
        public 
        view
        returns (address)
    {
        if(randomSelected[raceId] == 0 || registrationList[raceId][slotId].length == 0) return address(0);
        uint256 selectedIndex = randomSelected[raceId].mod(registrationList[raceId][slotId].length);
        return registrationList[raceId][slotId][selectedIndex];
    }

    function totalRegister(bytes32 raceId) 
        external 
        view
        returns (uint256[] memory)
    {
        (, uint256 _slots,) = checkRaceExisted(raceId);
        uint256[] memory results = new uint256[](_slots);
        for (uint256 slotId = 0; slotId < _slots; slotId++) {
            results[slotId] = registrationList[raceId][slotId].length;
        }
        return results;
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

    function removeReward(
        bytes32 raceId, uint256 nftRewardId, bytes1 resultIndex
    )  
        external 
        override
        whenNotPaused
        onlyRole(ADMIN_ROLE) {

        checkRaceExisted(raceId);

        if (raceRewards[raceId][resultIndex] != nftRewardId){
            revert RewardIsNotExistedOrReceived();
        }
        _raceReward.safeTransferFrom(
            address(this),
            _msgSender(),
            nftRewardId
        );
        raceRewards[raceId][resultIndex] = 0;
        emit RewardRemoved(raceId, nftRewardId, resultIndex);
    } 

    function checkRaceExisted(bytes32 raceId) internal view returns (uint256, uint256, bytes32) {
        ILeague _leagueAddr = ILeague(leagueRace[raceId]);
        uint256 _startAt = _leagueAddr.raceStartAt(raceId);
        if(_startAt == 0) revert RaceNotExisted();
        uint256 _slots = _leagueAddr.raceSlots(raceId);
        bytes32 _result = _leagueAddr.raceResult(raceId);
        return (_startAt, _slots, _result);
    }
}
