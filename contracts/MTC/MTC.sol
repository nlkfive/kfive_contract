// SPDX-License-Identifier: MIT
pragma solidity 0.8.4;

import "../BEP20/IBEP20.sol";
import "../BEP20/BEP20.sol";
import "@openzeppelin/contracts/security/Pausable.sol";
import "./BlackList.sol";
import "./TokenAdmin.sol";
import "@openzeppelin/contracts/utils/Address.sol";

contract MTC is BEP20("MTC", "MTC", 10), Pausable, BlackList, TokenAdmin {
    using Address for address;

    event __transferByAdmin();
    event __issue();
    event __redeem();
    event KfiveReceiverTransferred(address oldReceiver, address newReceiver);
    event ExchangeFromKFIVE(address user, uint256 value);
    event KfiveAddressChanged(address kfiveAddress);

    // exchanged KFIVE to MTC will be transfer to this address
    address private kfiveReceiver;

    // exchange
    IBEP20 public kfiveToken;

    // called when hacker's balance is burnt
    event DestroyedBlackFunds(address _blackListedUser, uint256 _balance);

    constructor(address _kfiveAddress) {
        require(
            _kfiveAddress.isContract(),
            "The kfive token address must be a deployed contract"
        );
        kfiveReceiver = msg.sender;
        kfiveToken = IBEP20(_kfiveAddress);
    }

    /**
     * @dev function to transfer MTC
     * @param from the address to transfer from
     * @param to the address to transfer to
     * @param value the amount to be transferred
     */
    function transferByAdmin(
        address from,
        address to,
        uint256 value
    ) public onlyOwner {
        _transfer(from, to, value);
        emit __transferByAdmin();
    }

    function issue(address issuer, uint256 value) public onlyOwner {
        _issue(issuer, value);
    }

    function issueByAdmin(address issuer, uint256 value) public onlyAdmin {
        require(
            value <= admin[msg.sender].maxIssuingTokenPerTime,
            "minted value must be smaller or equal to admin max issue"
        );

        require(
            value <= admin[msg.sender].remainingIssuingToken,
            "minted value must be smaller or equal to remain issue value"
        );

        admin[msg.sender].remainingIssuingToken -= value;
        _issue(issuer, value);
    }

    /**
     * @dev function to purchase new MTC
     * @param issuer the address that will receive the newly minted MTC
     * @param value the amount of MTC to mint
     */
    function _issue(address issuer, uint256 value) private {
        require(
            totalSupply() + value <= cap,
            "minted value + total supply must be smaller or equal token cap"
        );
        _mint(issuer, value);
        emit __issue();
    }

    /**
     * @dev function to burn MTC of hacker
     * @param _blackListedUser the account whose MTC will be burnt
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
     * @dev function to burn MTC
     * @param redeemer the account whose MTC will be burnt
     * @param value the amount of MTC to be burnt
     */
    function redeem(address redeemer, uint256 value) public onlyOwner {
        _burn(redeemer, value);
        emit __redeem();
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

    /**
     * @dev Transfers KFIVE receiver to a new account (`newReceiver`).
     * Can only be called by the owner.
     */
    function transferKfiveReceiver(address newReceiver) public onlyOwner {
        require(
            newReceiver != address(0),
            "Ownable: new owner is the zero address"
        );
        emit KfiveReceiverTransferred(kfiveReceiver, newReceiver);
        kfiveReceiver = newReceiver;
    }

    /**
     * @dev Update KFIVE address.
     * Can only be called by the owner.
     */
    function updateKfiveAddress(address newAddress) public onlyOwner {
        require(
            newAddress.isContract(),
            "The kfive token address must be a deployed contract"
        );
        emit KfiveAddressChanged(newAddress);
        kfiveToken = IBEP20(newAddress);
    }

    function exchangeFromKfive(uint256 _value) public whenNotPaused {
        require(!isBlackListed[msg.sender]);
        require(
            msg.sender != kfiveReceiver,
            "Exchanger must be different from the kfive receiver"
        );
        kfiveToken.transferFrom(msg.sender, kfiveReceiver, _value);
        emit ExchangeFromKFIVE(msg.sender, _value);
        _issue(msg.sender, KFIVE2MTC * _value);
    }
}
