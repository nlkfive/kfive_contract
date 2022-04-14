// SPDX-License-Identifier: MIT

pragma solidity 0.8.4;

import "./KfiveNFT.sol";

contract NLGST is KfiveNFT {
    constructor(address _marketplaceStorage, string memory baseTokenURI)
        KfiveNFT(
            _marketplaceStorage,
            "Ngoc Linh Ginseng Token",
            "NLGST",
            baseTokenURI
        )
        _requireMarketplaceStorage(_marketplaceStorage)
    {}
}
