// SPDX-License-Identifier: MIT
pragma solidity 0.8.4;

import "./IRaceList.sol";
import "../common/AccessControl.sol";
import "@openzeppelin/contracts/utils/introspection/ERC165.sol";
import "@chainlink/contracts/src/v0.8/ChainlinkClient.sol";

contract RaceList is
    KfiveAccessControl,
    IRaceList,
    ChainlinkClient
{
    using Chainlink for Chainlink.Request;
    
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
    // From api request id to running raceId
    mapping(bytes32 => bytes32) public apiRequestMap;

    address private oracle;
    bytes32 private jobId;
    uint256 private fee;

    constructor(address _oracle, bytes32 _jobId, uint256 _fee) {
        setPublicChainlinkToken();
        oracle = _oracle;
        jobId = _jobId;
        fee = _fee;
    }

    function raceIsExisted(bytes32 raceId)
        external
        view
        override
        returns (bool)
    {
        return races[raceId].registerAt != 0;
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
        uint256 registerAt,
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
        require(slots < 32, "Invalid slots");
        require(betStarted < betEnded, "Invalid bet time");
        require(registerAt < betEnded, "Invalid register time");
        // These two value will be divide to 1000
        // Commission will be share to the bething holder
        // Its value (percentage) must be lower than 1
        require(commission < 1000, "Invalid commission");
        // Reward rate will be multiply to the base reward for 1st and 2nd bettors
        // Its value (percentage) must be bigger than 1
        require(rewardRate >= 1000, "Invalid rewardRate");
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
            race.commission = commission;
            race.rewardRate = rewardRate;
        }

        emit RaceCreated(
            raceId,
            slots,
            registerAt,
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
        require(races[id].registerAt != 0, "Not existed");
        onlyBefore(races[id].betStarted);
        delete races[id];
    }

    /**
     * @dev Update race result.
     */
    function updateResult(
        bytes32 id
    ) 
        external 
        override 
        whenNotPaused
        onlyRole(ADMIN_ROLE) 
    {
        Race storage race = races[id];
        onlyAfter(race.betEnded);

        Chainlink.Request memory request = buildChainlinkRequest(jobId, address(this), this.fulfill.selector);
        request.addBytes("id", abi.encodePacked(id));
        bytes32 requestId = sendChainlinkRequestTo(oracle, request, fee);
        apiRequestMap[requestId] = id;
    }

    function fulfill(bytes32 _requestId, bytes32 _data) public recordChainlinkFulfillment(_requestId) {
        bytes32 raceId = apiRequestMap[_requestId];
        Race storage race = races[raceId];
        require(race.registerAt != 0, "Not existed");
        race.result = _data;
        emit RaceResultUpdated(raceId, _data);
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
