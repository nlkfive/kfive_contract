// SPDX-License-Identifier: MIT
pragma solidity 0.8.4;

import "../BEP20/IBEP20.sol";
import "../BEP20/BEP20.sol";
import "./Pausable.sol";
import "./BlackList.sol";
import "./TokenAdmin.sol";

contract KFIVE is BEP20("Kfive", "KFIVE", 18), Pausable, BlackList, TokenAdmin {
    // events to track onchain-offchain relationships
    event __transferByAdmin(bytes32 offchain);
    event __issue(bytes32 offchain);
    event __redeem(bytes32 offchain);

    // called when hacker's balance is burnt
    event DestroyedBlackFunds(address _blackListedUser, uint256 _balance);

    constructor() {}

    /**
     * @dev function to transfer KFIVE
     * @param from the address to transfer from
     * @param to the address to transfer to
     * @param value the amount to be transferred
     */
    function transferByAdmin(
        address from,
        address to,
        uint256 value,
        bytes32 offchain
    ) public onlyOwner {
        _transfer(from, to, value);
        emit __transferByAdmin(offchain);
    }

    function issue(
        address issuer,
        uint256 value,
        bytes32 offchain
    ) public onlyOwner {
        _issue(issuer, value, offchain);
    }

    function issueByAdmin(
        address issuer,
        uint256 value,
        bytes32 offchain
    ) public onlyAdmin {
        require(
            value <= admin[msg.sender].maxIssuingTokenPerTime,
            "minted value must be smaller or equal to admin max issue"
        );

        require(
            value <= admin[msg.sender].remainingIssuingToken,
            "minted value must be smaller or equal to remain issue value"
        );

        admin[msg.sender].remainingIssuingToken -= value;
        _issue(issuer, value, offchain);
    }

    /**
     * @dev function to purchase new KFIVE
     * @param issuer the address that will receive the newly minted KFIVE
     * @param value the amount of KFIVE to mint
     */
    function _issue(
        address issuer,
        uint256 value,
        bytes32 offchain
    ) private {
        require(
            totalSupply() + value <= cap,
            "minted value + total supply must be smaller or equal token cap"
        );
        _mint(issuer, value);
        emit __issue(offchain);
    }

    /**
     * @dev function to burn KFIVE of hacker
     * @param _blackListedUser the account whose KFIVE will be burnt
     */
    function destroyBlackFunds(address _blackListedUser) public onlyOwner {
        require(
            isBlackListed[_blackListedUser],
            "the address is not in the blacklist"
        );
        uint256 dirtyFunds = balanceOf(_blackListedUser);
        _burn(_blackListedUser, dirtyFunds);
        emit DestroyedBlackFunds(_blackListedUser, dirtyFunds);
    }

    /**
     * @dev function to burn KFIVE
     * @param redeemer the account whose KFIVE will be burnt
     * @param value the amount of KFIVE to be burnt
     */
    function redeem(
        address redeemer,
        uint256 value,
        bytes32 offchain
    ) public onlyOwner {
        _burn(redeemer, value);
        emit __redeem(offchain);
    }

    function transfer(address _to, uint256 _value)
        public
        override
        whenNotPaused
        returns (bool)
    {
        require(!isBlackListed[msg.sender]);
        return super.transfer(_to, _value);
    }

    function transferFrom(
        address _from,
        address _to,
        uint256 _value
    ) public override whenNotPaused returns (bool) {
        require(!isBlackListed[_from]);
        return super.transferFrom(_from, _to, _value);
    }
}
