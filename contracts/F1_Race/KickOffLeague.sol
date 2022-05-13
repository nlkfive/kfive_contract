// SPDX-License-Identifier: MIT
pragma solidity 0.8.4;

import "./League.sol";
import "../NFT/IKfiveNFT.sol";

contract KickOffLeague is League
{
    using Address for address;
    IKfiveNFT private _nlggt;

    constructor(address nlggtAddr, address raceReward, uint256 noRace, string memory _leagueName) League(raceReward, noRace, _leagueName){
        if(!nlggtAddr.isContract()) revert InvalidContract();
        _nlggt = IKfiveNFT(nlggtAddr);
    }

    function checkValidRegister() internal override view {
        if (_nlggt.balanceOf(_msgSender()) > 0) revert InvalidRegister();
    }
}