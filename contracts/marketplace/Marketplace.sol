// SPDX-License-Identifier: MIT
pragma solidity 0.8.4;

import "@openzeppelin/contracts/security/Pausable.sol";
import "../BEP20/IBEP20.sol";
import "../common/StorageLock.sol";
import "../common/BlackList.sol";
import "./storage/IMarketplaceStorage.sol";
import "@openzeppelin/contracts/token/ERC721/IERC721.sol";
import "@openzeppelin/contracts/access/AccessControlEnumerable.sol";

contract Marketplace is
    AccessControlEnumerable,
    StorageLock,
    Pausable,
    BlackList
{
    IBEP20 public acceptedToken;

    uint256 public ownerCutPerMillion;
    uint256 public publicationFeeInWei;

    bytes4 public constant ERC721_Interface = bytes4(0x80ac58cd);

    event ChangedPublicationFee(uint256 publicationFee);
    event ChangedOwnerCutPerMillion(uint256 ownerCutPerMillion);

    bytes32 public constant ADMIN_ROLE = keccak256("ADMIN_ROLE");
    bytes32 public constant CANCEL_ROLE = keccak256("CANCEL_ROLE");
    bytes32 public constant PAUSER_ROLE = keccak256("PAUSER_ROLE");

    using Address for address;

    /**
     * @dev Initialize this contract. Acts as a constructor
     * @param _acceptedToken - Address of the ERC20 accepted for this marketplace
     * @param _ownerCutPerMillion - owner cut per million
     */
    constructor(
        address _acceptedToken,
        address _marketplaceStorage,
        uint256 _ownerCutPerMillion
    ) StorageLock(_marketplaceStorage) {
        transferOwnership(_msgSender());

        require(_acceptedToken.isContract(), "Invalid contract");
        acceptedToken = IBEP20(_acceptedToken);
        marketplaceStorage = IMarketplaceStorage(_marketplaceStorage);

        _setupRole(DEFAULT_ADMIN_ROLE, _msgSender());

        _setupRole(ADMIN_ROLE, _msgSender());
        _setupRole(CANCEL_ROLE, _msgSender());
        _setupRole(PAUSER_ROLE, _msgSender());

        _setRoleAdmin(CANCEL_ROLE, ADMIN_ROLE);
        _setRoleAdmin(PAUSER_ROLE, ADMIN_ROLE);

        // Fee init
        setOwnerCutPerMillion(_ownerCutPerMillion);
    }

    /**
     * @dev Sets the publication fee that's charged to users to publish items
     * @param _publicationFee - Fee amount in wei this contract charges to publish an item
     */
    function setPublicationFee(uint256 _publicationFee)
        external
        onlyRole(ADMIN_ROLE)
    {
        publicationFeeInWei = _publicationFee;
        emit ChangedPublicationFee(publicationFeeInWei);
    }

    /**
     * @dev Sets the share cut for the owner of the contract that's
     *  charged to the seller on a successful sale
     * @param _ownerCutPerMillion - Share amount, from 0 to 999,999
     */
    function setOwnerCutPerMillion(uint256 _ownerCutPerMillion)
        public
        onlyRole(ADMIN_ROLE)
    {
        require(_ownerCutPerMillion < 1000000, "Invalid cut");
        ownerCutPerMillion = _ownerCutPerMillion;
        emit ChangedOwnerCutPerMillion(ownerCutPerMillion);
    }

    modifier _requireERC721(address nftAddress) {
        require(nftAddress.isContract(), "Invalid contract");
        require(
            IERC721(nftAddress).supportsInterface(ERC721_Interface),
            "Invalid erc721"
        );
        _;
    }

    function _msgSender() internal view override returns (address sender) {
        if (msg.sender == address(this)) {
            bytes memory array = msg.data;
            uint256 index = msg.data.length;
            assembly {
                // Load the 32 bytes word from memory with the address on the lower 20 bytes, and mask those.
                sender := and(
                    mload(add(array, index)),
                    0xffffffffffffffffffffffffffffffffffffffff
                )
            }
        } else {
            sender = msg.sender;
        }
        return sender;
    }

    function pause() public onlyRole(PAUSER_ROLE) {
        _pause();
    }

    function unpause() public onlyRole(PAUSER_ROLE) {
        _unpause();
    }

    function transferOwnership(address newOwner) public override onlyOwner {
        _setupRole(CANCEL_ROLE, newOwner);
        _setupRole(PAUSER_ROLE, newOwner);
        _setupRole(DEFAULT_ADMIN_ROLE, newOwner);

        revokeRole(CANCEL_ROLE, owner());
        revokeRole(PAUSER_ROLE, owner());
        revokeRole(DEFAULT_ADMIN_ROLE, owner());

        super.transferOwnership(newOwner);
    }

    function updateStorageAddress(address _marketplaceStorage)
        public
        override
        onlyRole(ADMIN_ROLE)
    {
        super.updateStorageAddress(_marketplaceStorage);
    }
}
