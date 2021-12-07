// SPDX-License-Identifier: MIT

pragma solidity 0.8.4;

import "@openzeppelin/contracts/token/ERC721/ERC721.sol";
import "@openzeppelin/contracts/token/ERC721/extensions/ERC721Enumerable.sol";
import "@openzeppelin/contracts/token/ERC721/extensions/ERC721Pausable.sol";
import "@openzeppelin/contracts/token/ERC721/extensions/ERC721URIStorage.sol";
import "@openzeppelin/contracts/access/AccessControlEnumerable.sol";
import "../common/BlackList.sol";
import "../common/StorageLock.sol";

contract KfiveNFT is
    StorageLock,
    AccessControlEnumerable,
    ERC721Enumerable,
    ERC721Pausable,
    ERC721URIStorage,
    BlackList
{
    bytes32 public constant ADMIN_ROLE = keccak256("ADMIN_ROLE");
    bytes32 public constant MINTER_ROLE = keccak256("MINTER_ROLE");
    bytes32 public constant PAUSER_ROLE = keccak256("PAUSER_ROLE");

    string private _baseTokenURI;

    /**
     * @dev Grants `DEFAULT_ADMIN_ROLE`, `MINTER_ROLE` and `PAUSER_ROLE` to the
     * account that deploys the contract.
     *
     * Token URI format is '`baseTokenURI/QmHash`'.
     * See {ERC721-tokenURI}.
     */
    constructor(
        address _marketplaceStorage,
        string memory tokenName,
        string memory tokenSymbol,
        string memory baseTokenURI
    ) StorageLock(_marketplaceStorage) ERC721(tokenName, tokenSymbol) {
        _baseTokenURI = baseTokenURI;

        _grantRole(DEFAULT_ADMIN_ROLE, _msgSender());
        _grantRole(ADMIN_ROLE, _msgSender());

        _setRoleAdmin(MINTER_ROLE, ADMIN_ROLE);
        _setRoleAdmin(PAUSER_ROLE, ADMIN_ROLE);

        marketplaceStorage = IMarketplaceStorage(_marketplaceStorage);
    }

    function _baseURI() internal view virtual override returns (string memory) {
        return _baseTokenURI;
    }

    /**
     * @dev Creates a new token for `to`. Its token ID will be automatically
     * assigned (and available on the emitted {IERC721-Transfer} event), and the token
     * URI autogenerated based on the base URI passed at construction.
     *
     * See {ERC721-_mint}.
     *
     * Requirements:
     *
     * - the caller must have the `MINTER_ROLE`.
     */
    function mint(
        address to,
        uint256 tokenId,
        string memory gtokenURI
    ) public virtual onlyRole(MINTER_ROLE) {
        // We cannot just use balanceOf to create the new tokenId because tokens
        // can be burned (destroyed), so we need a separate counter.
        _safeMint(to, tokenId);
        _setTokenURI(tokenId, gtokenURI);
    }

    /**
     * @dev Pauses all token transfers.
     *
     * See {ERC721Pausable} and {Pausable-_pause}.
     *
     * Requirements:
     *
     * - the caller must have the `PAUSER_ROLE`.
     */
    function pause() public virtual onlyRole(PAUSER_ROLE) {
        _pause();
    }

    /**
     * @dev Unpauses all token transfers.
     *
     * See {ERC721Pausable} and {Pausable-_unpause}.
     *
     * Requirements:
     *
     * - the caller must have the `PAUSER_ROLE`.
     */
    function unpause() public virtual onlyRole(PAUSER_ROLE) {
        _unpause();
    }

    function _beforeTokenTransfer(
        address from,
        address to,
        uint256 tokenId
    ) internal virtual override(ERC721, ERC721Enumerable, ERC721Pausable) {
        require(!isBlackListed[from]);
        bytes32 nftAsset = keccak256(abi.encodePacked(address(this), tokenId));
        require(marketplaceStorage.assetIsAvailable(nftAsset), "Unavailable");
        super._beforeTokenTransfer(from, to, tokenId);
    }

    /**
     * @dev See {IERC165-supportsInterface}.
     */
    function supportsInterface(bytes4 interfaceId)
        public
        view
        virtual
        override(AccessControlEnumerable, ERC721, ERC721Enumerable)
        returns (bool)
    {
        return super.supportsInterface(interfaceId);
    }

    function _burn(uint256 tokenId)
        internal
        virtual
        override(ERC721, ERC721URIStorage)
    {
        super._burn(tokenId);
    }

    function burn(uint256 tokenId) public onlyRole(ADMIN_ROLE) {
        _burn(tokenId);
    }

    function tokenURI(uint256 tokenId)
        public
        view
        virtual
        override(ERC721, ERC721URIStorage)
        returns (string memory)
    {
        return super.tokenURI(tokenId);
    }

    function transferOwnership(address newOwner) public override onlyOwner {
        _grantRole(ADMIN_ROLE, newOwner);
        _grantRole(DEFAULT_ADMIN_ROLE, newOwner);

        _revokeRole(ADMIN_ROLE, owner());
        _revokeRole(DEFAULT_ADMIN_ROLE, owner());

        super.transferOwnership(newOwner);
    }

    function _grantRole(bytes32 role, address account) internal override {
        super._grantRole(role, account);
        if (role == ADMIN_ROLE) {
            super._grantRole(MINTER_ROLE, account);
            super._grantRole(PAUSER_ROLE, account);
        }
    }

    function _revokeRole(bytes32 role, address account) internal override {
        super._revokeRole(role, account);
        if (role == ADMIN_ROLE) {
            super._revokeRole(MINTER_ROLE, account);
            super._revokeRole(PAUSER_ROLE, account);
        }
    }

    function updateStorageAddress(address _marketplaceStorage)
        public
        override
        onlyRole(ADMIN_ROLE)
    {
        super.updateStorageAddress(_marketplaceStorage);
    }
}
