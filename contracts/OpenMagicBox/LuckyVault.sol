// SPDX-License-Identifier: MIT

pragma solidity 0.8.4;

import "@openzeppelin/contracts/utils/Address.sol";
import "@openzeppelin/contracts/utils/math/SafeMath.sol";
import "../common/AccessControl.sol";
import "../BEP20/IBEP20.sol";
import "@openzeppelin/contracts/security/ReentrancyGuard.sol";

contract LuckyVault is KfiveAccessControl, ReentrancyGuard {
    using Address for address;
    using SafeMath for uint256;

    IBEP20 private _tradeToken;
    IBEP20 private _depositToken;
    IBEP20 private _rewardToken;
    address _receiver;
    uint256 _tradeRate;
    uint256 _tradeRewardRate;
    uint256 _depositRewardRate;
    uint256 _endedAt;

    mapping(address => uint256) depositList;

    error InvalidContract();
    error DepositFailed();
    error RewardFailed();
    error TradeFailed();
    error SendDepositTokenFailed();
    error TooLate(uint256 time);
    error TooEarly(uint256 time);
    error InvalidWithdrawAmount();
    error WithdrawFailed();
    error AdminWithdrawFailed();
    error InsufficientBalance();
    error InvalidAmount();
    error NoReward();

    event DepositSuccessful(address sender, uint256 amount, uint256 rewardAmount);
    event TradeSuccessful(address sender, uint256 amount, uint256 tradeAmount, uint256 rewardAmount);
    event WithdrawSuccessful(address sender, uint256 amount);
    event AdminWithdrawSuccessful(address sender, address receiver, uint256 amount);

    constructor(
        address tradeToken, // BUSD
        address depositToken, // KFIVE
        address rewardToken, // LF5
        address receiver, // Address where receive tradeToken
        uint256 tradeRate, // 5.2 BUSD/KFIVE -> 0.192 => 192
        uint256 tradeRewardRate, // 10 KFIVE/LF5 -> 
        uint256 depositRewardRate,
        uint256 endedAt
    ) {
        if(!tradeToken.isContract()) revert InvalidContract();
        if(!depositToken.isContract()) revert InvalidContract();
        if(!rewardToken.isContract()) revert InvalidContract();
        if (block.timestamp >= endedAt) revert TooLate(endedAt);

        _tradeToken = IBEP20(tradeToken);
        _depositToken = IBEP20(depositToken);
        _rewardToken = IBEP20(rewardToken);
        _receiver = receiver;
        _tradeRate = tradeRate;
        _tradeRewardRate = tradeRewardRate;
        _depositRewardRate = depositRewardRate;
        _endedAt = endedAt;
    }

    /**
     * @dev Deposit lock token
     */
    function deposit(
        uint256 amount
    ) 
        external
        nonReentrant
        whenNotPaused
    {
        address sender = _msgSender();

        // KFIVE -> LF5 
        // div(10**10).div(1000) -> div(10**13)
        uint256 rewardAmount = amount.mul(_depositRewardRate).div(10**13);
        if (rewardAmount == 0) revert NoReward();

        if (!_depositToken.transferFrom(sender, address(this), amount)) revert DepositFailed();
        
        _transferReward(sender, rewardAmount);

        depositList[sender] = depositList[sender].add(amount);
        emit DepositSuccessful(sender, amount, rewardAmount);
    }

    /**
     * @dev Trade lock token
     */
    function trade(
        uint256 amount
    ) 
        external
        nonReentrant
        whenNotPaused
    {
        if (amount < 1 ether) revert InvalidAmount(); // At least 1 BUSD
        address sender = _msgSender();
        _preValidatePurchase(sender, amount);

        // BUSD -> KFIVE 
        // div(1 eth).mul(10**10).div(1000) -> div(10**11)
        uint256 tradeAmount = amount.mul(_tradeRate).div(10**11); 
        checkBalance(tradeAmount);

        // KFIVE -> LF5 
        // div(10**10).div(1000) -> div(10**13)
        uint256 rewardAmount = tradeAmount.mul(_tradeRewardRate).div(10**13);
        if (rewardAmount == 0) revert NoReward();

        if (!_tradeToken.transferFrom(sender, _receiver, amount)) revert TradeFailed(); // BUSD
        if (!_depositToken.transfer(sender, tradeAmount)) revert SendDepositTokenFailed(); // KFIVE

        _transferReward(sender, rewardAmount);

        emit TradeSuccessful(sender, amount, tradeAmount, rewardAmount);
    }

    /** 
     * @dev Withdraw lock token (User)
     */
    function withdraw(
        uint256 amount
    ) 
        external
        whenNotPaused
    {
        onlyAfter(_endedAt);

        address sender = _msgSender();
        if (depositList[sender] < amount) revert InvalidWithdrawAmount();
        checkBalance(amount);

        depositList[sender] = depositList[sender].sub(amount);
        if (!_depositToken.transfer(sender, amount)) revert WithdrawFailed();
        
        emit WithdrawSuccessful(sender, amount);
    }

    /** 
     * @dev Withdraw lock token (Admin)
     */
    function withdrawByAdmin(
        address receiver,
        uint256 amount
    ) 
        external
        whenNotPaused
        onlyRole(ADMIN_ROLE)
    {
        checkBalance(amount);
        if (!_depositToken.transfer(receiver, amount)) revert AdminWithdrawFailed();
        
        emit AdminWithdrawSuccessful(_msgSender(), receiver, amount);
    }

    function _transferReward(
        address sender,
        uint256 amount
    )
        internal
        whenNotPaused
    {
        if (!_rewardToken.transfer(sender, amount)) revert RewardFailed();
    }

    function onlyAfter(uint256 _time) internal view {
        if (block.timestamp <= _time) revert TooEarly(_time);
    }

    function checkBalance(uint256 amount) internal view {
        uint256 tokenBalance = _depositToken.balanceOf(address(this));
        if (tokenBalance < amount) revert InsufficientBalance();
    }

    /**
     * @dev Validation of an incoming purchase. Use require statements to revert state when conditions are not met.
     * Use `super` in contracts that inherit from BEP20Crowdsale to extend their validations.
     * Example from CappedCrowdsale.sol's _preValidatePurchase method:
     *     super._preValidatePurchase(beneficiary, weiAmount);
     *     require(weiRaised().add(weiAmount) <= cap);
     * @param beneficiary Address performing the token purchase
     * @param weiAmount Value in wei involved in the purchase
     */
    function _preValidatePurchase(address beneficiary, uint256 weiAmount)
        internal
        view
        virtual
    {
        require(beneficiary != address(0), "Invalid beneficiary");
        require(weiAmount != 0, "Invalid");
        this; // silence state mutability warning without generating bytecode - see https://github.com/ethereum/solidity/issues/2691
    }
}