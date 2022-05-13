// SPDX-License-Identifier: MIT
pragma solidity 0.8.4;

import "./interfaces/ILeague.sol";
import "../common/AccessControl.sol";
import "@openzeppelin/contracts/utils/introspection/ERC165.sol";
import "@openzeppelin/contracts/utils/Address.sol";

contract League is
    KfiveAccessControl,
    ILeague
{
    using Address for address;
    string private _leagueName;

    // Race info
    mapping(bytes32 => uint256) private _raceSlot;
    mapping(bytes32 => uint256) private _raceStartAt;
    mapping(bytes32 => bytes32) private _raceResult;

    constructor(
        string memory name
    ) {
        _leagueName = name;
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
            interfaceId == type(ILeague).interfaceId ||
            super.supportsInterface(interfaceId);
    }

    function raceSlots(bytes32 raceId)
        external
        view
        override
        returns (uint256 slots)
    {
        slots = _raceSlot[raceId];
    }

    function raceStartAt(bytes32 raceId)
        external
        view
        override
        returns (uint256 startAt)
    {
        startAt = _raceStartAt[raceId];
    }

    function raceResult(bytes32 raceId)
        external
        view
        override
        returns (bytes32 result)
    {
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
        onlyBefore(startAt);

        bytes32 raceId = keccak256(
            abi.encodePacked(
                _leagueName,
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
            startAt
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

    function onlyAfter(uint256 _time) internal view {
        if (block.timestamp <= _time) revert TooEarly(_time);
    }

    function onlyBefore(uint256 _time) internal view {
        if (block.timestamp >= _time) revert TooLate(_time);
    }
}