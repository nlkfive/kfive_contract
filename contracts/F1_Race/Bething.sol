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
    mapping(bytes32 => mapping(uint256 => mapping(address => uint256))) private bettors;
    // From raceId to (slotId to total bet)
    mapping(bytes32 => mapping(uint256 => uint256)) private totalSlotBets;
    // From raceId to total bet
    mapping(bytes32 => uint256) private totalBets;
    // From raceId to commission is received
    mapping(bytes32 => bool) private commissionReceived;

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
        require(_race.betStarted != 0, "Not existed");
        require(slotId < _race.slots, "Invalid slot id");
        onlyAfter(_race.betStarted);
        onlyBefore(_race.betEnded);

        require(
            _acceptedToken.transferFrom(_msgSender(), address(this), betValue),
            "Bet failed"
        );
        bettors[raceId][slotId][_msgSender()].add(betValue);
        totalSlotBets[raceId][slotId].add(betValue);
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
        require(_race.betStarted != 0, "Not existed");
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
     * 1st = total1stBet * rewardRate^2 * [totalRaceBet - %commission] 
     * div (total1stBet*rewardRate^2 + total2ndBet*rewardRate + total3rdBet)
     ***
     * 2nd = total2ndBet * rewardRate * [totalRaceBet - %commission] 
     * div (total1stBet*rewardRate^2 + total2ndBet*rewardRate + total3rdBet)
     ***
     * 3rd = total3rdBet * [totalRaceBet - %commission] 
     * div (total1stBet*rewardRate^2 + total2ndBet*rewardRate + total3rdBet)
     */
    function claim(
        bytes32 raceId,
        uint256 slotId
    ) 
        external
        override
        whenNotPaused
    {
        Race memory _race = _raceList.getRace(raceId);
        require(_race.betStarted != 0, "Not existed");
        onlyAfter(_race.betEnded);

        uint256 totalTop3BetReward = _totalTop3BetReward(
            raceId,
            _race.slots,
            _race.result,
            _race.rewardRate
        );
        uint256 reward = _calculateReward(
            raceId, 
            slotId,
            _race.rewardRate,
            _race.commission,
            totalTop3BetReward
        );
        bettors[raceId][slotId][_msgSender()] = 0;
        _acceptedToken.transfer(_msgSender(), reward);
        emit ClaimSuccessful(raceId, reward);
    }

    /**
     * @dev Claim commission from ended race - Admin only
     */
    function claimCommission(
        bytes32 raceId,
        address receiver
    ) 
        external 
        override
    {
        require(!commissionReceived[raceId], "Already received");
        Race memory _race = _raceList.getRace(raceId);
        require(_race.betStarted != 0, "Not existed");
        onlyAfter(_race.betEnded);
        uint256 commission = _calculateCommission(raceId, _race.commission);
        commissionReceived[raceId] = true;
        _acceptedToken.transfer(receiver, commission);
        emit ClaimCommission(raceId, commission, receiver);
    }

    /**
     * @dev Get your total slot bet
     */
    function totalSlotBet(
        bytes32 raceId,
        uint256 slotId
    ) 
        public 
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
        public 
        override
        view
        returns (uint256) 
    {
        return totalBets[raceId];
    }

    /**
     * @dev Get total race bet
     */
    function getSlotPosition(
        bytes32 raceId,
        uint256 slotId
    ) 
        public 
        override
        view
        returns (uint256) 
    {
        Race memory _race = _raceList.getRace(raceId);
        return uint256(_race.result[slotId]);
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

    // = totalRaceBet * (1000 - commission)/1000
    function _calculateCommission(bytes32 raceId, uint256 commission) 
        private 
        returns (uint256)
    {
        return totalRaceBet(raceId).mul(1000 - commission).div(1000);
    }

    function _totalTop3BetReward(
        bytes32 raceId, 
        uint256 slots,
        bytes32 result,
        uint256 rewardRate
    ) 
        private
        returns (uint256 totalTop3BetReward)
    {
        for (uint256 slotId = 0; slotId < slots; slotId++) {
            if (result[index] == 1){
                totalTop3BetReward.add(
                    totalSlotBets[raceId][slotId].mul(rewardRate).mul(rewardRate)
                );
            } else if (result[index] == 2) {
                totalTop3BetReward.add(
                    totalSlotBets[raceId][slotId].mul(rewardRate)
                );
            } else if (result[index] == 3){
                totalTop3BetReward.add(
                    totalSlotBets[raceId][slotId]
                );
            }
        }
    }

    // * Top 3 position:
    // * slotReward = totalSlotBet * rewardRate^(3-position) * [totalRaceBet - %commission] 
    // * div (total1stBet*rewardRate^2 + total2ndBet*rewardRate + total3rdBet)
    function _calculateReward(
        bytes32 raceId, 
        uint256 slotId,
        uint256 rewardRate,
        uint256 commission,
        uint256 totalTop3BetReward
    ) 
        private 
        returns (uint256 reward)
    {
        // if the race isn't ended or you are out of top 3
        uint256 position = getSlotPosition(raceId, slotId);
        if (position > 0 && position < 3){
            reward = totalSlotBet(raceId, slotId).mul(
                totalRaceBet(raceId) - _calculateCommission(raceId, commission)
            ).div(
                totalTop3BetReward
            );

            for (uint256 count = 0; count < position; count++) {
                reward = reward.mul(rewardRate);
            }
        }
    }

    function onlyBefore(uint256 _time) internal view {
        if (block.timestamp >= _time) revert TooLate(_time);
    }

    function onlyAfter(uint256 _time) internal view {
        if (block.timestamp <= _time) revert TooEarly(_time);
    }
}
