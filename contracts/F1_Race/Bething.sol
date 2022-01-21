// SPDX-License-Identifier: MIT
pragma solidity 0.8.4;

import "./IBething.sol";
import "../common/AccessControl.sol";
import "./IRaceList.sol";
import "@openzeppelin/contracts/utils/introspection/ERC165.sol";
import "@openzeppelin/contracts/utils/Address.sol";
import "@openzeppelin/contracts/token/ERC721/extensions/IERC721Enumerable.sol";
import "./IRace.sol";
import "../BEP20/IBEP20.sol";
import "@openzeppelin/contracts/utils/math/SafeMath.sol";

contract Bething is
    IRace,
    KfiveAccessControl,
    IBething
{
    using Address for address;
    using SafeMath for uint256;

    IRaceList public _raceList;
    IBEP20 public _acceptedToken;

    constructor(
        address acceptedToken,
        address raceList
    ) {
        require(acceptedToken.isContract(), "Invalid token");
        require(raceList.isContract(), "Invalid RaceList");
        _raceList = IRaceList(raceList);
        _acceptedToken = IBEP20(acceptedToken);
    }

    // From raceId to (slotId to (bettor to total bet))
    mapping(bytes32 => mapping(uint256 => mapping(address => uint256))) public bettors;
    // From raceId to total bet
    mapping(bytes32 => uint256) public totalBets;

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
            interfaceId == type(IBething).interfaceId ||
            super.supportsInterface(interfaceId);
    }

    /**
     * @dev Bet a race slot
     */
    function bet(
        uint256 slotId,
        bytes32 raceId,
        uint256 betValue
    ) 
        external
        override
        whenNotPaused
    {
        Race memory _race = _raceList.getRace(raceId);
        require(_race.registerAt != 0, "Not existed");
        require(slotId < _race.slots, "Invalid slot id");
        onlyAfter(_race.betStarted);
        onlyBefore(_race.betEnded);

        require(
            _acceptedToken.transferFrom(_msgSender(), address(this), betValue),
            "Bet failed"
        );
        bettors[raceId][slotId][_msgSender()].add(betValue);
        totalBets[raceId].add(betValue);
        emit BetSuccessful(slotId, raceId, betValue);
    }

    /**
     * @dev Call by admin - Fund some reward for the winners
     */
    function fundRace(
        bytes32 raceId,
        uint256 fundValue
    ) 
        external
        override
        whenNotPaused
    {
        Race memory _race = _raceList.getRace(raceId);
        require(_race.registerAt != 0, "Not existed");
        onlyAfter(_race.betStarted);
        onlyBefore(_race.betEnded);

        require(
            _acceptedToken.transferFrom(_msgSender(), address(this), fundValue),
            "Fund failed"
        );
        totalBets[raceId].add(fundValue);
        emit FundSuccessful(raceId, fundValue);
    }

    /**
     * @dev Claim winning reward from the ended race
     */
    function claim(
        bytes32 raceId
    ) 
        external
        override
        whenNotPaused
    {
        Race memory _race = _raceList.getRace(raceId);
        require(_race.registerAt != 0, "Not existed");
        onlyAfter(_race.betEnded);

        // TODO: Calculate the reward
        emit ClaimSuccessful(raceId, 0);
    }

    /**
     * @dev Get your total slot bet
     */
    function totalSlotBet(
        bytes32 raceId,
        uint256 slotId
    ) 
        external 
        override
        view
        returns (uint256) 
    {
        return bettors[raceId][slotId][_msgSender()];
    }

    /**
     * @dev Get total race bet
     */
    function totalRaceBet(
        bytes32 raceId
    ) 
        external 
        override
        view
        returns (uint256) 
    {
        return totalBets[raceId];
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
    function updateAcceptTokenAddress(address acceptToken)
        external
        override
        onlyRole(ADMIN_ROLE)
        whenNotPaused
    {
        _acceptedToken = IBEP20(acceptToken);
        emit AcceptTokenUpdated(acceptToken);
    }

    function onlyBefore(uint256 _time) internal view {
        if (block.timestamp >= _time) revert TooLate(_time);
    }

    function onlyAfter(uint256 _time) internal view {
        if (block.timestamp <= _time) revert TooEarly(_time);
    }
}
