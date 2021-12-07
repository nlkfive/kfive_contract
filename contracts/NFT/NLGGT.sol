// SPDX-License-Identifier: MIT

pragma solidity 0.8.4;

import "./KfiveNFT.sol";

contract NLGGT is KfiveNFT {
    constructor(address _marketplaceStorage, string memory baseTokenURI)
        KfiveNFT(
            _marketplaceStorage,
            "Ngoc Linh Ginseng Garden Token",
            "NLGGT",
            baseTokenURI
        )
        _requireMarketplaceStorage(_marketplaceStorage)
    {}
}
