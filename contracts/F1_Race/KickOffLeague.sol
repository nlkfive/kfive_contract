// SPDX-License-Identifier: MIT
pragma solidity 0.8.4;

import "./League.sol";
import "@openzeppelin/contracts/utils/Address.sol";
import "../NFT/IKfiveNFT.sol";

contract KickOffLeague is League
{
    using Address for address;

    IKfiveNFT private _nlggt;

    constructor(string memory _leagueName, address nlggtAddr) League(_leagueName){
        if(!nlggtAddr.isContract()) revert InvalidContract();
        _nlggt = IKfiveNFT(nlggtAddr);
    }

    function registerIsValid(address register)
        external 
        override
        view
        returns (bool) {
        return _nlggt.balanceOf(register) > 0;
    }
}