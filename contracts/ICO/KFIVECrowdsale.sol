// SPDX-License-Identifier: MIT

pragma solidity 0.8.4;

import "./Crowdsale.sol";
import "./validation/CappedCrowdsale.sol";
import "./validation/TimedCrowdsale.sol";

contract KFIVECrowdsale is Crowdsale, CappedCrowdsale, TimedCrowdsale {
    constructor(
        uint256 rate, // rate, in TKNbits
        address payable wallet, // wallet to send Ether
        IBEP20 token, // the token
        uint256 cap, // total cap, in wei
        uint256 openingTime, // opening time in unix epoch seconds
        uint256 closingTime // closing time in unix epoch seconds
    )
        CappedCrowdsale(cap)
        TimedCrowdsale(openingTime, closingTime)
        Crowdsale(rate, wallet, token)
    {
        // nice, we just created a crowdsale that's only open
        // for a certain amount of time
        // and stops accepting contributions once it reaches `cap`
    }

    function _preValidatePurchase(address beneficiary, uint256 weiAmount)
        internal
        view
        override(Crowdsale, CappedCrowdsale, TimedCrowdsale)
    {
        super._preValidatePurchase(beneficiary, weiAmount);
    }
}
