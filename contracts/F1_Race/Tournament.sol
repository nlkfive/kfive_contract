// SPDX-License-Identifier: MIT
pragma solidity 0.8.4;

import "./interfaces/ITournament.sol";
import "../common/AccessControl.sol";
import "@openzeppelin/contracts/utils/introspection/ERC165.sol";
import "@openzeppelin/contracts/token/ERC721/IERC721Receiver.sol";
import "../NFT/IKfiveNFT.sol";
import "@openzeppelin/contracts/utils/Address.sol";
import "@openzeppelin/contracts/utils/math/SafeMath.sol";

abstract contract Tournament is
    KfiveAccessControl,
    ITournament,
    IERC721Receiver
{
    using Address for address;
    using SafeMath for uint256;

    TournamentInfo private _tournamentInfo;
    IKfiveNFT private _f1Reward;

    // Race info
    mapping(bytes32 => Race) private _raceInfo;
    // RaceId -> User address -> slotId
    mapping(bytes32 => mapping(address => uint8)) private _registeredSlot;
    // RaceNo -> RaceId
    mapping(uint8 => bytes32) private _listRaces;
    // UserAddress -> Score
    mapping(address => uint8) private _userScore;

    constructor(
        uint8 totalRace,
        address raceReward,
        string memory name
    ) {
        if(!raceReward.isContract()) revert InvalidContract();
        _f1Reward = IKfiveNFT(raceReward);
        _tournamentInfo = TournamentInfo({
            totalRace: totalRace,
            createdRace: 0,
            endedRace: 0,
            tournamentName: name
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
            interfaceId == type(ITournament).interfaceId ||
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

    function tournamentInfo() 
        external 
        view 
        override
        returns (TournamentInfo memory) 
    {
        return _tournamentInfo;
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
        for (uint8 i = 0; i < _tournamentInfo.createdRace; i++) {
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
        uint8 raceNo = _tournamentInfo.createdRace + 1;
        if (raceNo > _tournamentInfo.totalRace) revert CannotCreateMoreRace(raceNo, _tournamentInfo.totalRace);
        onlyBefore(startAt);

        bytes32 raceId = keccak256(
            abi.encodePacked(
                raceNo,
                noSlot,
                startAt
            )
        );
        if (_raceInfo[raceId].startAt == 0) revert RaceNotExisted();
        _tournamentInfo.createdRace += 1;
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
        _tournamentInfo.endedRace = _tournamentInfo.endedRace - 1;
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
        _tournamentInfo.endedRace = _tournamentInfo.endedRace + 1;
        emit RaceResultUpdated(raceId, result);
    }

    /**
     * @dev Mint nft to the winner
     * Admin check the winner list offchain and announce the winner to onchain by using this function
     */
    function grandReward(uint8 winnerIndex, address winner, uint256 nftRewardId, string memory tokenURI) 
        external 
        override
        whenNotPaused
        onlyRole(ADMIN_ROLE) 
    {
        if (_tournamentInfo.endedRace < _tournamentInfo.totalRace) revert NotEndYet();
        _f1Reward.mint(
            address(this),
            nftRewardId,
            tokenURI
        );
        emit RewardGranted(winnerIndex, getTotalScore(winner), winner, nftRewardId);
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