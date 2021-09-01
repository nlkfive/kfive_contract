// SPDX-License-Identifier: MIT
pragma solidity 0.8.4;

import "../BEP20/IBEP20.sol";
import "./BlindAuctionStorage.sol";
import "./OrderStorage.sol";
import "@openzeppelin/contracts/token/ERC721/IERC721.sol";

contract MarketplaceStorage {
    IBEP20 public acceptedToken;

    uint256 public ownerCutPerMillion;
    uint256 public publicationFeeInWei;

    bytes4 public constant ERC721_Interface = bytes4(0x80ac58cd);

    event ChangedPublicationFee(uint256 publicationFee);
    event ChangedOwnerCutPerMillion(uint256 ownerCutPerMillion);
}
