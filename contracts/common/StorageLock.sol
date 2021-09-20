// SPDX-License-Identifier: MIT
pragma solidity 0.8.4;
import "../marketplace/storage/IMarketplaceStorage.sol";
import "@openzeppelin/contracts/utils/Address.sol";

contract StorageLock {
    IMarketplaceStorage public marketplaceStorage;
    bytes4 public constant IMarketplaceStorage_Interface = bytes4(0x85c93a92);
    using Address for address;

    constructor(address _marketplaceStorage)
        _requireMarketplaceStorage(_marketplaceStorage)
    {
        marketplaceStorage = IMarketplaceStorage(_marketplaceStorage);
    }

    function updateStorageAddress(address _marketplaceStorage) public virtual {
        marketplaceStorage = IMarketplaceStorage(_marketplaceStorage);
        emit MarketplaceStorageUpdated(_marketplaceStorage);
    }

    event MarketplaceStorageUpdated(address _marketplaceStorage);

    modifier _requireMarketplaceStorage(address storageAddress) {
        require(storageAddress.isContract(), "Invalid contract");
        IMarketplaceStorage storageAddressRegistry = IMarketplaceStorage(
            storageAddress
        );
        require(
            storageAddressRegistry.supportsInterface(
                IMarketplaceStorage_Interface
            ),
            "Invalid storage"
        );
        _;
    }
}
