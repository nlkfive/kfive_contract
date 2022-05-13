// SPDX-License-Identifier: MIT
pragma solidity 0.8.4;

import "./interfaces/ILeague.sol";
import "../common/AccessControl.sol";
import "@openzeppelin/contracts/utils/introspection/ERC165.sol";
import "@openzeppelin/contracts/token/ERC721/IERC721Receiver.sol";
import "../NFT/IKfiveNFT.sol";
import "@openzeppelin/contracts/utils/Address.sol";

abstract contract League is
    KfiveAccessControl,
    ILeague,
    IERC721Receiver
{
    using Address for address;

    string private _leagueName;
    uint256 private _totalRace;
    IKfiveNFT private _f1Reward;
    uint256 private _currentRaceNo = 1;

    // Race info
    mapping(bytes32 => uint256) private _raceSlot;
    mapping(bytes32 => uint256) private _raceStartAt;
    mapping(bytes32 => bytes32) private _raceResult;
    // WinnerIndex => NFTID
    mapping(uint256 => uint256) private _listRewards;
    // Raceid -> User Address -> isRegistered
    mapping(bytes32 => mapping(address => bool)) private registeredSlot;
    // RaceId -> slotId -> list registers
    mapping(bytes32 => mapping(uint256 => address[])) private registrationList;

    constructor(
        address raceReward,
        uint256 totalRace,
        string memory name
    ) {
        if(!raceReward.isContract()) revert InvalidContract();
        _f1Reward = IKfiveNFT(raceReward);
        _totalRace = totalRace;
        _leagueName = name;
    }
    
    function supportsInterface(bytes4 interfaceId)
        public
        view
        virtual
        override(AccessControlEnumerable, IERC165)
        returns (bool)
    {
        return
            interfaceId == type(ILeague).interfaceId ||
            super.supportsInterface(interfaceId);
    }

    function raceInfo(bytes32 raceId)
        external
        view
        override
        returns (uint256 slots, uint256 startAt, bytes32 result)
    {
        slots = _raceSlot[raceId];
        startAt = _raceStartAt[raceId];
        result = _raceResult[raceId];
    }

    function leagueName() 
        external 
        view 
        override
        returns (string memory) 
    {
        return _leagueName;
    }

    function registerRace(
        bytes32 raceId,
        uint256 slotId
    )
        external 
        override
        whenNotPaused
    {
        onlyBefore(_raceStartAt[raceId]);
        checkValidRegister();
        if(slotId >= _raceSlot[raceId]) revert InvalidSlot();
        if(registeredSlot[raceId][_msgSender()]) revert AlreadyRegistered();

        registeredSlot[raceId][_msgSender()] = true;
        registrationList[raceId][slotId].push(_msgSender());
        
        emit Registered(slotId, _msgSender(), raceId);
    }

    /**
     * @dev Create new race
     */
    function createRace(
        uint256 slots,
        uint256 startAt
    ) 
        external 
        override 
        whenNotPaused
        onlyRole(ADMIN_ROLE) 
        returns (bytes32)
    {
        if(slots >= 32) revert InvalidSlot();
        if (_currentRaceNo > _totalRace) revert CannotCreateMoreRace(_currentRaceNo, _totalRace);
        onlyBefore(startAt);

        bytes32 raceId = keccak256(
            abi.encodePacked(
                _currentRaceNo,
                slots,
                startAt
            )
        );
        if (_raceStartAt[raceId] != 0) revert RaceExisted();

        _raceSlot[raceId] = slots;
        _raceStartAt[raceId] = startAt;

        emit RaceCreated(
            raceId,
            slots,
            startAt,
            _currentRaceNo
        );
        _currentRaceNo += 1;
        return raceId;
    }

    /**
     * @dev Delete race by id.
     */
    function cancelRace(bytes32 raceId)
        external
        override
        whenNotPaused
        onlyRole(ADMIN_ROLE)
    {
        if (_raceStartAt[raceId] == 0) revert RaceNotExisted();
        if (_raceResult[raceId] != 0) revert CannotCancel();
        delete _raceStartAt[raceId];
        delete _raceSlot[raceId];
        emit RaceCancelled(raceId);
    }

    /**
     * @dev Update race result.
     */
    function updateRaceResult(
        bytes32 raceId,
        bytes32 result
    ) 
        external 
        override 
        whenNotPaused
        onlyRole(ADMIN_ROLE) 
    {
        if (_raceStartAt[raceId] == 0) revert RaceNotExisted();
        onlyAfter(_raceStartAt[raceId]);
        _raceResult[raceId] = result;
        emit RaceResultUpdated(raceId, result);
    }

    /**
     * @dev Add reward.
     */
    function addRewardByTransfer(bytes32 raceId, uint256 nftRewardId, uint256 winnerIndex) 
        external 
        override
        whenNotPaused
        onlyRole(ADMIN_ROLE)
    {
        if (_listRewards[winnerIndex] != 0) revert RewardIsExisted();

        _f1Reward.safeTransferFrom(
            _msgSender(),
            address(this),
            nftRewardId
        );
        _listRewards[winnerIndex] = nftRewardId;
        emit RewardAdded(raceId, nftRewardId, winnerIndex);
    }

    /**
     * @dev Add reward.
     */
    function addRewardByMint(bytes32 raceId, uint256 nftRewardId, uint256 winnerIndex, string memory tokenURI) 
        external 
        override
        whenNotPaused
        onlyRole(ADMIN_ROLE)
    {
        if (_listRewards[winnerIndex] != 0) revert RewardIsExisted();

        _listRewards[winnerIndex] = nftRewardId;
        _f1Reward.mint(
            address(this),
            nftRewardId,
            tokenURI
        );
        emit RewardAdded(raceId, nftRewardId, winnerIndex);
    }

    /**
     * @dev Receive reward after race ended.
     */
    function receiveReward(bytes32 raceId, uint256 slotId) 
        external 
        override
        whenNotPaused
    {
    }

    function removeReward(
        bytes32 raceId, uint256 winnerIndex
    )  
        external 
        override
        whenNotPaused
        onlyRole(ADMIN_ROLE) {

        if (_listRewards[winnerIndex] == 0) revert RewardIsNotExisted();

        _f1Reward.safeTransferFrom(
            address(this),
            _msgSender(),
            _listRewards[winnerIndex]
        );
        _listRewards[winnerIndex] = 0;
        emit RewardRemoved(raceId, _listRewards[winnerIndex], winnerIndex);
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

    function checkValidRegister() internal view virtual;

    function onlyAfter(uint256 _time) internal view {
        if (block.timestamp <= _time) revert TooEarly(_time);
    }

    function onlyBefore(uint256 _time) internal view {
        if (block.timestamp >= _time) revert TooLate(_time);
    }
}