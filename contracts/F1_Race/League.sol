// SPDX-License-Identifier: MIT
pragma solidity 0.8.4;

import "./interfaces/ILeague.sol";
import "../common/AccessControl.sol";
import "@openzeppelin/contracts/utils/introspection/ERC165.sol";
import "@openzeppelin/contracts/token/ERC721/IERC721Receiver.sol";
import "../NFT/IKfiveNFT.sol";
import "@openzeppelin/contracts/utils/Address.sol";
import "@openzeppelin/contracts/utils/math/SafeMath.sol";

abstract contract League is
    KfiveAccessControl,
    ILeague,
    IERC721Receiver
{
    using Address for address;
    using SafeMath for uint256;

    LeagueInfo private _leagueInfo;
    IKfiveNFT private _f1Reward;

    // Race info
    mapping(bytes32 => Race) private _raceInfo;
    // WinnerIndex => NFTID
    mapping(uint256 => uint256) private _listRewards;
    // RaceId -> User address -> slotId
    mapping(bytes32 => mapping(address => uint8)) private _registeredSlot;
    // RaceNo -> RaceId
    mapping(uint8 => bytes32) private _listRaces;

    constructor(
        uint8 totalRace,
        address raceReward,
        string memory name
    ) {
        if(!raceReward.isContract()) revert InvalidContract();
        _f1Reward = IKfiveNFT(raceReward);
        _leagueInfo = LeagueInfo({
            totalRace: totalRace,
            createdRace: 0,
            endedRace: 0,
            leagueName: name
        });
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
        returns (Race memory)
    {
        return _raceInfo[raceId];
    }

    function leagueInfo() 
        external 
        view 
        override
        returns (LeagueInfo memory) 
    {
        return _leagueInfo;
    }

    function registeredSlot(address register, bytes32 raceId)
        external 
        view 
        override
        returns (uint8) 
    {
        return _registeredSlot[raceId][register];
    }

    function getTotalScore(address register)
        public 
        view 
        override
        returns (uint8 score) 
    {
        for (uint8 i = 0; i < _leagueInfo.createdRace; i++) {
            bytes32 raceId = _listRaces[i];
            uint8 slotId = _registeredSlot[raceId][register];
            score += uint8(_raceInfo[raceId].result[slotId]);
        }
    }

    function registerRace(
        uint8 slotId,
        bytes32 raceId
    )
        external 
        override
        whenNotPaused
    {
        Race memory _race = _raceInfo[raceId];
        onlyBefore(_race.startAt);
        checkValidRegister();
        if(slotId > _race.noSlot || slotId == 0) revert InvalidSlot();
        if(_registeredSlot[raceId][_msgSender()] != 0) revert AlreadyRegistered();

        _registeredSlot[raceId][_msgSender()] = slotId;
        
        emit Registered(slotId, _msgSender(), raceId);
    }

    /**
     * @dev Create new race
     */
    function createRace(
        uint8 noSlot,
        uint32 startAt
    ) 
        external 
        override 
        whenNotPaused
        onlyRole(ADMIN_ROLE) 
        returns (bytes32)
    {
        if(noSlot > 26) revert InvalidSlot();
        uint8 raceNo = _leagueInfo.createdRace + 1;
        if (raceNo > _leagueInfo.totalRace) revert CannotCreateMoreRace(raceNo, _leagueInfo.totalRace);
        onlyBefore(startAt);

        bytes32 raceId = keccak256(
            abi.encodePacked(
                raceNo,
                noSlot,
                startAt
            )
        );
        if (_raceInfo[raceId].startAt == 0) revert RaceNotExisted();
        _leagueInfo.createdRace += 1;
        _listRaces[raceNo] = raceId;

        _raceInfo[raceId] = Race({
            noSlot: noSlot,
            startAt: startAt,
            result: 0
        });

        emit RaceCreated(
            noSlot,
            raceNo,
            startAt,
            raceId
        );
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
        Race memory _race = _raceInfo[raceId];
        if (_race.startAt == 0) revert RaceNotExisted();
        if (_race.result != 0) revert RaceWasUpdated();
        _leagueInfo.endedRace = _leagueInfo.endedRace - 1;
        delete _raceInfo[raceId];
        emit RaceCancelled(raceId);
    }

    /**
     * @dev Update race result.
     * Highest score each race = 5 -> 51 races -> max score 255
     * Highest score each race = 15 -> 17 races -> max score 255
     * Result : Ignore first bytes because of slotId start at 1, 1 -> 26
     */
    function updateRaceResult(
        bytes32 raceId,
        bytes27 result
    ) 
        external 
        override
        whenNotPaused
        onlyRole(ADMIN_ROLE) 
    {
        Race storage _race = _raceInfo[raceId];
        if (_race.startAt == 0) revert RaceNotExisted();
        if (_race.result != 0) revert RaceWasUpdated();
        onlyAfter(_race.startAt);
        _race.result = result;
        _leagueInfo.endedRace = _leagueInfo.endedRace + 1;
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
    function receiveReward(uint8 slotId, bytes32 raceId) 
        external 
        override
        whenNotPaused
    {
        if (_leagueInfo.endedRace < _leagueInfo.totalRace) revert NotEndYet();

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