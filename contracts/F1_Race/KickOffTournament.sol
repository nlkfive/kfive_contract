// SPDX-License-Identifier: MIT
pragma solidity 0.8.4;

import "./Tournament.sol";
import "../NFT/IKfiveNFT.sol";

contract KickOffTournament is Tournament
{
    using Address for address;
    IKfiveNFT private _nlggt;

    constructor(uint8 noRace, address nlggtAddr, address raceReward, string memory _tournamentName) Tournament(noRace, raceReward, _tournamentName){
        if(!nlggtAddr.isContract()) revert InvalidContract();
        _nlggt = IKfiveNFT(nlggtAddr);
    }

    function checkValidRegister() internal override view {
        if (_nlggt.balanceOf(_msgSender()) > 0) revert InvalidRegister();
    }
}

