// SPDX-License-Identifier: MIT

pragma solidity 0.8.4;

import "../common/AccessControl.sol";
import "@openzeppelin/contracts/token/ERC721/IERC721Receiver.sol";
import "../BEP20/IBEP20.sol";
import "../NFT/IKfiveNFT.sol";
import "@openzeppelin/contracts/utils/Address.sol";

contract LF5Exchange is 
    KfiveAccessControl, 
    IERC721Receiver
{
    using Address for address;

    IBEP20 private _rewardToken;
    IKfiveNFT private _magicBoxReward;
    enum BoxType{ GOLD, SILVER, BROZEN }
    mapping (BoxType => uint256) _price;

    error InvalidContract(); // 0x6eefed20
    error InvalidReward(); // 0x28829e82
    error NotOwner(); // 0x30cd7471
    error ReceiveLF5Failed(); // 0xa74238ad

    event ExchangeSuccess(uint256 nftId, BoxType boxType);

    constructor(
        address magicBoxReward,
        address rewardToken
    ) {
        if(!magicBoxReward.isContract()) revert InvalidContract();
        if(!rewardToken.isContract()) revert InvalidContract();

        _rewardToken = IBEP20(rewardToken);
        _magicBoxReward = IKfiveNFT(magicBoxReward);

        _price[BoxType.GOLD] = 15;
        _price[BoxType.SILVER] = 5;
        _price[BoxType.BROZEN] = 1;
    }

    function exchangeForLF5(uint256 nftId, bytes memory nftData) external {
        address assetOwner = _magicBoxReward.ownerOf(nftId);
        address sender = _msgSender();
        if (sender != assetOwner){
            revert NotOwner();
        }
        BoxType boxType = checkValidNft(nftId, nftData);
        _magicBoxReward.safeTransferFrom(
            sender,
            address(this),
            nftId
        );

        if (!_rewardToken.transfer(sender, _price[boxType])) revert ReceiveLF5Failed();

        emit ExchangeSuccess(nftId, boxType);
    }

    function checkValidNft(uint256 nftId, bytes memory nftData) internal pure returns (BoxType boxType){
        if (nftId != uint256(keccak256(nftData))){
            revert InvalidReward();
        }
        bytes32 boxBytesType = getBoxType(nftData);
        if (boxBytesType == bytes32(0x7b226e616d65223a224f70656e20476f6c6420426f78222c22676966745f636f)){
            boxType = BoxType.GOLD;
        } else if (boxBytesType == bytes32(0x7b226e616d65223a224f70656e2053696c76657220426f78222c22676966745f)) {
            boxType = BoxType.SILVER;
        } else if (boxBytesType == bytes32(0x7b226e616d65223a224f70656e2042726f6e7a6520426f78222c22676966745f)) {
            boxType = BoxType.BROZEN;
        } else {
            revert InvalidReward();
        }
    }

    function getBoxType(bytes memory nftData) internal pure returns (bytes32 boxBytesType){
        // Get 32 bytes from data
        assembly {
            boxBytesType := mload(add(nftData,0x20))
        }
    }

    function onERC721Received(
        address,
        address,
        uint256,
        bytes calldata
    ) external pure override returns (bytes4) {
        return
            bytes4(
                keccak256("onERC721Received(address,address,uint256,bytes)")
            );
    }
}