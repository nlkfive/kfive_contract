// SPDX-License-Identifier: MIT

pragma solidity 0.8.4;

import "./BEP20Crowdsale.sol";
import "./validation/CappedCrowdsale.sol";
import "./validation/TimedCrowdsale.sol";

contract BEP20KFIVECrowdsale is
    BEP20Crowdsale,
    CappedCrowdsale,
    TimedCrowdsale
{
    constructor(
        uint256 rate, // rate, in TKNbits
        address payable wallet, // wallet to send Ether
        IBEP20 token, // the token
        IBEP20 acceptToken, // the token
        uint256 cap, // total cap, in wei
        uint256 openingTime, // opening time in unix epoch seconds
        uint256 closingTime // closing time in unix epoch seconds
    )
        CappedCrowdsale(cap)
        TimedCrowdsale(openingTime, closingTime)
        BEP20Crowdsale(rate, wallet, token, acceptToken)
    {
        // nice, we just created a crowdsale that's only open
        // for a certain amount of time
        // and stops accepting contributions once it reaches `cap`
    }

    function _preValidatePurchase(address beneficiary, uint256 weiAmount)
        internal
        view
        override(BEP20Crowdsale, CappedCrowdsale, TimedCrowdsale)
    {
        super._preValidatePurchase(beneficiary, weiAmount);
    }
}
