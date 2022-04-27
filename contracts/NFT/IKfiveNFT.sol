// SPDX-License-Identifier: MIT

pragma solidity ^0.8.4;

import "@openzeppelin/contracts/token/ERC721/extensions/IERC721Enumerable.sol";

interface IKfiveNFT is IERC721Enumerable {
    function mint(
        address to,
        uint256 tokenId,
        string memory gtokenURI
    ) external;

    function burn(uint256 tokenId) external;
}
