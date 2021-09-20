// SPDX-License-Identifier: MIT
pragma solidity 0.8.4;

import "@openzeppelin/contracts/utils/Address.sol";
import "@openzeppelin/contracts/security/Pausable.sol";
import "../BEP20/IBEP20.sol";
import "../common/BlackList.sol";
import "./storage/IMarketplaceStorage.sol";
import "@openzeppelin/contracts/token/ERC721/IERC721.sol";
import "@openzeppelin/contracts/utils/cryptography/ECDSA.sol";
import "@openzeppelin/contracts/access/AccessControlEnumerable.sol";

contract Marketplace is AccessControlEnumerable, Pausable, BlackList {
    IBEP20 public acceptedToken;
    IMarketplaceStorage public marketplaceStorage;

    uint256 public ownerCutPerMillion;
    uint256 public publicationFeeInWei;

    bytes4 public constant ERC721_Interface = bytes4(0x80ac58cd);
    bytes4 public constant IMarketplaceStorage_Interface = bytes4(0x85c93a92);

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
    ) _requireMarketplaceStorage(_marketplaceStorage) {
        // Fee init
        setOwnerCutPerMillion(_ownerCutPerMillion);
        transferOwnership(_msgSender());

        require(
            _acceptedToken.isContract(),
            "The accepted token address must be a deployed contract"
        );
        acceptedToken = IBEP20(_acceptedToken);
        marketplaceStorage = IMarketplaceStorage(_marketplaceStorage);

        _setupRole(DEFAULT_ADMIN_ROLE, _msgSender());

        _setupRole(CANCEL_ROLE, _msgSender());
        _setupRole(PAUSER_ROLE, _msgSender());

        _setRoleAdmin(CANCEL_ROLE, ADMIN_ROLE);
        _setRoleAdmin(PAUSER_ROLE, ADMIN_ROLE);
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
        require(
            _ownerCutPerMillion < 1000000,
            "The owner cut should be between 0 and 999,999"
        );

        ownerCutPerMillion = _ownerCutPerMillion;
        emit ChangedOwnerCutPerMillion(ownerCutPerMillion);
    }

    modifier _requireERC721(address nftAddress) {
        require(
            nftAddress.isContract(),
            "The NFT Address should be a contract"
        );
        IERC721 nftRegistry = IERC721(nftAddress);
        require(
            nftRegistry.supportsInterface(ERC721_Interface),
            "The NFT contract has an invalid ERC721 implementation"
        );
        _;
    }

    modifier _requireMarketplaceStorage(address storageAddress) {
        require(
            storageAddress.isContract(),
            "The Marketplace storage Address should be a contract"
        );
        IMarketplaceStorage storageAddressRegistry = IMarketplaceStorage(
            storageAddress
        );
        require(
            storageAddressRegistry.supportsInterface(
                IMarketplaceStorage_Interface
            ),
            "The Marketplace storage contract has an invalid IMarketplaceStorage implementation"
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
}
