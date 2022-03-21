// SPDX-License-Identifier: MIT
pragma solidity 0.8.4;

import "./IRaceList.sol";
import "../common/AccessControl.sol";
import "@openzeppelin/contracts/utils/introspection/ERC165.sol";
import "@chainlink/contracts/src/v0.8/ChainlinkClient.sol";
import "@openzeppelin/contracts/utils/Address.sol";

contract RaceList is
    KfiveAccessControl,
    IRaceList,
    ChainlinkClient
{
    using Address for address;
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

    // https://market.link/jobs/5f9f57b4-afcb-4f3c-957d-63a7bb91469d
    // jobId: ce039030ac69497cbb9615b9b1d7b363 -> 0x2063653033393033306163363934393763626239363135623962316437623336
    // oracle: 0x77a5310E41F0B9FE35E95239Fa5624390fadFbBA
    // fee: 0.04 * 10 ** 18 LINK -> 40000000000000000
    constructor(address chainlink, address _oracle, bytes32 _jobId, uint256 _fee) {
        if(!chainlink.isContract()) revert InvalidContract();
        setChainlinkToken(chainlink);
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
        // TODO: add api get race result
        request.add("get", "https://staging.kfive.ru/race/result");
        request.addBytes("raceId", abi.encodePacked(id));
        bytes32 requestId = sendChainlinkRequestTo(oracle, request, fee);
        apiRequestMap[requestId] = id;
    }

    function fulfill(bytes32 _requestId, bytes32 _data) public recordChainlinkFulfillment(_requestId) {
        bytes32 raceId = apiRequestMap[_requestId];
        Race storage race = races[raceId];
        if (race.betStarted == 0) revert RaceNotExisted();
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