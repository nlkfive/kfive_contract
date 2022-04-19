// SPDX-License-Identifier: MIT
pragma solidity 0.8.4;

import "../BEP20/IBEP20.sol";
import "../BEP20/BEP20.sol";
import "@openzeppelin/contracts/security/Pausable.sol";
import "../common/BlackList.sol";
import "../common/TokenAdmin.sol";

contract LF5 is BEP20("Lucky K5", "LF5", 0), Pausable, BlackList, TokenAdmin {
    // events to track onchain-offchain relationships
    event __transferByAdmin(bytes32 offchain);
    event __issue(bytes32 offchain);
    event __redeem(bytes32 offchain);

    // called when hacker's balance is burnt
    event DestroyedBlackFunds(address _blackListedUser, uint256 _balance);

    function cap() internal pure override returns (uint256) {
        return 1080000;
    }

    constructor() {}

    /**
     * @dev function to transfer LF5
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
     * @dev function to purchase new LF5
     * @param issuer the address that will receive the newly minted LF5
     * @param value the amount of LF5 to mint
     */
    function _issue(
        address issuer,
        uint256 value,
        bytes32 offchain
    ) private {
        _mint(issuer, value);
        emit __issue(offchain);
    }

    /**
     * @dev function to burn LF5 of hacker
     * @param _blackListedUser the account whose LF5 will be burnt
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
     * @dev function to burn LF5
     * @param redeemer the account whose LF5 will be burnt
     * @param value the amount of LF5 to be burnt
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

    function pause() public onlyOwner {
        _pause();
    }

    function unpause() public onlyOwner {
        _unpause();
    }
}
