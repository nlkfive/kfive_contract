// SPDX-License-Identifier: MIT
pragma solidity 0.8.4;

import "./IRaceList.sol";
import "../common/AccessControl.sol";
import "@openzeppelin/contracts/utils/introspection/ERC165.sol";
import "@openzeppelin/contracts/utils/Address.sol";

contract RaceList is
    KfiveAccessControl,
    IRaceList
{
    using Address for address;
    
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
    mapping(bytes32 => Race) private races;

    function raceIsExisted(bytes32 raceId)
        external
        view
        override
        returns (bool)
    {
        return races[raceId].betStarted != 0;
    }

    function raceResult(bytes32 raceId)
        external
        view
        override
        returns (bytes32)
    {
        return races[raceId].result;
    }

    /**
     * @dev Create new race
     */
    function createRace(
        uint256 slots,
        uint256 betStarted,
        uint256 betEnded,
        uint256 commission,
        uint256 rewardRate
    ) 
        external 
        override 
        whenNotPaused
        onlyRole(ADMIN_ROLE) 
    {
        if(slots >= 32) revert InvalidSlot();
        if (block.timestamp >= betStarted) revert TooLate(betStarted);
        if (betStarted >= betEnded) revert TooLate(betEnded);
        // These two value will be divide to 1000
        // Commission will be share to the bething holder
        // Its value (percentage) must be lower than 1
        if (commission >= 1000) revert InvalidCommission();
        
        // Reward rate will be multiply to the base reward for 1st and 2nd bettors
        // Its value (percentage) must be bigger than 1
        if (rewardRate < 1000) revert InvalidRewardRate();

        bytes32 raceId = keccak256(
            abi.encodePacked(
                slots,
                betStarted,
                betEnded
            )
        );

        // Create new auction
        Race storage race = races[raceId];
        {
            race.slots = slots;
            race.betStarted = betStarted;
            race.betEnded = betEnded;
            race.commission = commission;
            race.rewardRate = rewardRate;
        }

        emit RaceCreated(
            raceId,
            slots,
            betStarted,
            betEnded,
            commission,
            rewardRate
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
        if (races[id].betStarted == 0) revert RaceNotExisted();
        onlyBefore(races[id].betStarted);
        delete races[id];
        emit RaceCancelled(id);
    }

    /**
     * @dev Update race result.
     */
    function updateResult(
        bytes32 id,
        bytes32 result
    ) 
        external 
        override 
        whenNotPaused
        onlyRole(ADMIN_ROLE) 
    {
        Race storage race = races[id];
        if (race.betStarted == 0) revert RaceNotExisted();
        onlyAfter(race.betEnded);

        race.result = result;
        emit RaceResultUpdated(id, result);
    }

    /**
     * @dev Update commission.
     */
    function updateCommission(
        bytes32 id,
        uint256 commission
    ) 
        external 
        override 
        whenNotPaused
        onlyRole(ADMIN_ROLE) 
    {
        Race storage race = races[id];
        onlyBefore(race.betEnded);
        // These two value will be divide to 1000
        // Commission will be share to the bething holder
        // Its value (percentage) must be lower than 1
        if (commission >= 1000) revert InvalidCommission();
        race.commission = commission;
        emit RaceCommissionUpdated(id, commission);
    }

    /**
     * @dev Update rewardRate.
     */
    function updateRewardRate(
        bytes32 id,
        uint256 rewardRate
    )
        external 
        override 
        whenNotPaused
        onlyRole(ADMIN_ROLE) 
    {
        Race storage race = races[id];
        onlyBefore(race.betEnded);
        // Reward rate will be multiply to the base reward for 1st and 2nd bettors
        // Its value (percentage) must be bigger than 1
        if (rewardRate < 1000) revert InvalidRewardRate();

        race.rewardRate = rewardRate;
        emit RaceRewardRateUpdated(id, rewardRate);
    }

    function onlyAfter(uint256 _time) internal view {
        if (block.timestamp <= _time) revert TooEarly(_time);
    }

    function onlyBefore(uint256 _time) internal view {
        if (block.timestamp >= _time) revert TooLate(_time);
    }
}