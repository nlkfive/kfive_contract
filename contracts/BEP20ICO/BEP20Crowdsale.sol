// SPDX-License-Identifier: MIT

pragma solidity 0.8.4;

import "@openzeppelin/contracts/utils/Context.sol";
import "../BEP20/IBEP20.sol";
import "@openzeppelin/contracts/utils/math/SafeMath.sol";
import "@openzeppelin/contracts/security/ReentrancyGuard.sol";
import "@openzeppelin/contracts/access/Ownable.sol";
import "@openzeppelin/contracts/security/Pausable.sol";

/**
 * @title BEP20Crowdsale
 * @dev BEP20Crowdsale is a base contract for managing a token crowdsale,
 * allowing investors to purchase tokens with ether. This contract implements
 * such functionality in its most fundamental form and can be extended to provide additional
 * functionality and/or custom behavior.
 * The external interface represents the basic interface for purchasing tokens, and conforms
 * the base architecture for crowdsales. It is *not* intended to be modified / overridden.
 * The internal interface conforms the extensible and modifiable surface of crowdsales. Override
 * the methods to add functionality. Consider using 'super' where appropriate to concatenate
 * behavior.
 */
contract BEP20Crowdsale is Context, Ownable, ReentrancyGuard, Pausable {
    using SafeMath for uint256;

    // The token being sold
    IBEP20 private _token;

    // The token using for buying
    IBEP20 private _acceptToken;

    // Address where funds are collected
    address payable private _wallet;

    // How many token units a buyer gets per wei.
    // The rate is the conversion between wei and the smallest and indivisible token unit.
    // So, if you are using a rate of 1 with a ERC20Detailed token with 3 decimals called TOK
    // 1 wei will give you 1 unit, or 0.001 TOK.
    uint256 private _rate;

    // Amount of wei raised
    uint256 private _weiRaised;

    /**
     * Event for token purchase logging
     * @param purchaser who paid for the tokens
     * @param beneficiary who got the tokens
     * @param value weis paid for purchase
     * @param amount amount of tokens purchased
     */
    event TokensPurchased(
        address indexed purchaser,
        address indexed beneficiary,
        uint256 value,
        uint256 amount
    );

    /**
     * @param __rate Number of token units a buyer gets per wei
     * @dev The rate is the conversion between wei and the smallest and indivisible
     * token unit. So, if you are using a rate of 1 with a ERC20Detailed token
     * with 3 decimals called TOK, 1 wei will give you 1 unit, or 0.001 TOK.
     * @param __wallet Address where collected funds will be forwarded to
     * @param __token Address of the token being sold
     * @param __acceptToken Address of the token using for buying
     */
    constructor(
        uint256 __rate,
        address payable __wallet,
        IBEP20 __token,
        IBEP20 __acceptToken
    ) {
        require(__rate > 0, "Invalid rate");
        require(__wallet != address(0), "Invalid address");
        require(address(__token) != address(0), "Invalid token");
        require(address(__acceptToken) != address(0), "Invalid accept token");

        _rate = __rate;
        _wallet = __wallet;
        _token = __token;
        _acceptToken = __acceptToken;
    }

    /**
     * @return the token being sold.
     */
    function token() public view returns (IBEP20) {
        return _token;
    }

    /**
     * @return the address where funds are collected.
     */
    function wallet() public view returns (address payable) {
        return _wallet;
    }

    /**
     * @return the number of token units a buyer gets per wei.
     */
    function rate() public view returns (uint256) {
        return _rate;
    }

    /**
     * @return the amount of wei raised.
     */
    function weiRaised() public view returns (uint256) {
        return _weiRaised;
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

    /**
     * @dev Validation of an executed purchase. Observe state and use revert statements to undo rollback when valid
     * conditions are not met.
     * @param beneficiary Address performing the token purchase
     * @param weiAmount Value in wei involved in the purchase
     */
    function _postValidatePurchase(address beneficiary, uint256 weiAmount)
        internal
        view
    {
        // solhint-disable-previous-line no-empty-blocks
    }

    /**
     * @dev Source of tokens. Override this method to modify the way in which the crowdsale ultimately gets and sends
     * its tokens.
     * @param beneficiary Address performing the token purchase
     * @param tokenAmount Number of tokens to be emitted
     */
    function _deliverTokens(address beneficiary, uint256 tokenAmount) internal {
        _token.transfer(beneficiary, tokenAmount);
    }

    /**
     * @dev Executed when a purchase has been validated and is ready to be executed. Doesn't necessarily emit/send
     * tokens.
     * @param beneficiary Address receiving the tokens
     * @param tokenAmount Number of tokens to be purchased
     */
    function _processPurchase(address beneficiary, uint256 tokenAmount)
        internal
    {
        _deliverTokens(beneficiary, tokenAmount);
    }

    /**
     * @dev Override for extensions that require an internal state to check for validity (current user contributions,
     * etc.)
     * @param beneficiary Address receiving the tokens
     * @param weiAmount Value in wei involved in the purchase
     */
    function _updatePurchasingState(address beneficiary, uint256 weiAmount)
        internal
    {
        // solhint-disable-previous-line no-empty-blocks
    }

    /**
     * @dev Override to extend the way in which ether is converted to tokens.
     * @param weiAmount Value in wei to be converted into tokens
     * @return Number of tokens that can be purchased with the specified _weiAmount
     */
    function _getTokenAmount(uint256 weiAmount)
        internal
        view
        returns (uint256)
    {
        return weiAmount.mul(_rate);
    }

    /**
     * @dev Determines how BEP20 is stored/forwarded on purchases.
     */
    function _forwardFunds(address buyer, uint256 value) internal {
        _acceptToken.transferFrom(buyer, _wallet, value);
    }

    /**
     * @dev Destroy Smart Contracts using selfdestruct
     * @param _to all remaining funds on the address of the Smart Contract are transferred to this address
     */
    function destroySmartContract(address payable _to) public onlyOwner {
        uint256 tokenBalance = _token.balanceOf(address(this));
        _token.transfer(_to, tokenBalance);

        uint256 acceptTokenBalance = _acceptToken.balanceOf(address(this));
        _acceptToken.transfer(_to, acceptTokenBalance);

        selfdestruct(_to);
    }

    function changeWallet(address payable newReceiver) external onlyOwner {
        _wallet = newReceiver;
    }

    /**
     * @dev Pauses all token transfers.
     *
     * Requirements:
     *
     * - the caller must have the `PAUSER_ROLE`.
     */
    function pause() public whenNotPaused onlyOwner {
        _pause();
    }

    /**
     * @dev Unpauses all token transfers.
     *
     * Requirements:
     *
     * - the caller must have the `PAUSER_ROLE`.
     */
    function unpause() public whenPaused onlyOwner {
        _unpause();
    }

    /**
     * @dev low level token purchase ***DO NOT OVERRIDE***
     * This function has a non-reentrancy guard, so it shouldn't be called by
     * another `nonReentrant` function.
     * @param beneficiary Recipient of the token purchase
     */
    function buyTokens(address beneficiary, uint256 value)
        public
        payable
        nonReentrant
        whenNotPaused
    {
        require(value >= 1 ether, "Invalid amount");
        uint256 ethAmount = value.div(1 ether);
        _preValidatePurchase(beneficiary, ethAmount);

        // calculate token amount to be created
        uint256 tokens = _getTokenAmount(ethAmount);

        _forwardFunds(beneficiary, ethAmount.mul(1 ether));

        // update state
        _updatePurchasingState(beneficiary, ethAmount);
        _weiRaised = _weiRaised.add(ethAmount);

        _processPurchase(beneficiary, tokens);
        emit TokensPurchased(_msgSender(), beneficiary, ethAmount, tokens);

        _postValidatePurchase(beneficiary, ethAmount);
    }
}
