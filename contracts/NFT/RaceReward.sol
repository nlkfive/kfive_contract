// SPDX-License-Identifier: MIT

pragma solidity 0.8.4;

import "./KfiveNFT.sol";

contract RaceReward is KfiveNFT {
    constructor(address _marketplaceStorage, string memory baseTokenURI)
        KfiveNFT(
            _marketplaceStorage,
            "Ngoc Linh F1 Reward Token",
            "NLGF1",
            baseTokenURI
        )
        _requireMarketplaceStorage(_marketplaceStorage)
    {}
}
