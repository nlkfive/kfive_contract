// SPDX-License-Identifier: MIT

pragma solidity 0.8.4;

import "./KfiveNFT.sol";

contract MagicBoxReward is KfiveNFT {
    constructor(address _marketplaceStorage, string memory baseTokenURI)
        KfiveNFT(
            _marketplaceStorage,
            "Magic Box Token",
            "MB5",
            baseTokenURI
        )
        _requireMarketplaceStorage(_marketplaceStorage)
    {}
}

