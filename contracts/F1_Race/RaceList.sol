// SPDX-License-Identifier: MIT
pragma solidity 0.8.4;

import "./IRaceList.sol";
import "../common/AccessControl.sol";
import "@openzeppelin/contracts/utils/introspection/ERC165.sol";

contract RaceList is
    KfiveAccessControl,
    IRaceList
{
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
            interfaceId == type(IRaceList).interfaceId ||
            super.supportsInterface(interfaceId);
    }

    ////////////////////////////////////////////////////////////
    // Race list
    ////////////////////////////////////////////////////////////
    // From running raceId to race
    mapping(bytes32 => Race) public races;

    function raceIsExisted(bytes32 raceId)
        external
        view
        override
        returns (bool)
    {
        return races[raceId].registerAt != 0;
    }

    /**
     * @dev Create new race
     */
    function createRace(
        uint256 slots,
        uint256 registerAt,
        uint256 betStarted,
        uint256 betEnded
    ) 
        external 
        override 
        whenNotPaused
        onlyRole(ADMIN_ROLE) 
    {
        require(slots < 32, "Invalid slots");
        require(betStarted < betEnded, "Invalid bet time");
        require(registerAt < betEnded, "Invalid register time");
        bytes32 raceId = keccak256(
            abi.encodePacked(
                slots,
                registerAt,
                betStarted,
                betEnded
            )
        );

        // Create new auction
        Race storage race = races[raceId];
        {
            race.slots = slots;
            race.registerAt = registerAt;
            race.betStarted = betStarted;
            race.betEnded = betEnded;
        }

        emit RaceCreated(
            raceId,
            slots,
            registerAt,
            betStarted,
            betEnded
        );
    }

    /**
     * @dev Get race by raceId.
     */
    function getRace(bytes32 raceId)
        external
        view
        override
        returns (Race memory race)
    {
        return races[raceId];
    }

    /**
     * @dev Delete race by id.
     */
    function cancelRace(bytes32 id)
        external
        override
        onlyRole(ADMIN_ROLE)
    {
        require(races[id].registerAt != 0, "Already ended");
        delete races[id];
    }

    /**
     * @dev Update highest bid auction by auction ID.
     */
    function updateRaceResult(
        bytes32 id
    ) 
        external 
        view
        override 
        whenNotPaused
        onlyRole(ADMIN_ROLE) 
    {
        Race storage race = races[id];
        onlyAfter(race.betEnded);
        // TODO: get result from oracle
    }

    function onlyAfter(uint256 _time) internal view {
        if (block.timestamp <= _time) revert TooEarly(_time);
    }
}
