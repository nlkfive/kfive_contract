// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package StorageLock

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// StorageLockMetaData contains all meta data concerning the StorageLock contract.
var StorageLockMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_marketplaceStorage\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_marketplaceStorage\",\"type\":\"address\"}],\"name\":\"MarketplaceStorageUpdated\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"marketplaceStorage\",\"outputs\":[{\"internalType\":\"contractIMarketplaceStorage\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_marketplaceStorage\",\"type\":\"address\"}],\"name\":\"updateStorageAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b5060405161048938038061048983398181016040528101906100329190610120565b8061005c8173ffffffffffffffffffffffffffffffffffffffff166100e860201b6101131760201c565b61009b576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016100929061016c565b60405180910390fd5b6000819050826000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050505061020f565b6000808273ffffffffffffffffffffffffffffffffffffffff163b119050919050565b60008151905061011a816101f8565b92915050565b60006020828403121561013257600080fd5b60006101408482850161010b565b91505092915050565b600061015660108361018c565b9150610161826101cf565b602082019050919050565b6000602082019050818103600083015261018581610149565b9050919050565b600082825260208201905092915050565b60006101a8826101af565b9050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b7f496e76616c696420636f6e747261637400000000000000000000000000000000600082015250565b6102018161019d565b811461020c57600080fd5b50565b61026b8061021e6000396000f3fe608060405234801561001057600080fd5b50600436106100365760003560e01c806346b3aec61461003b578063889e212914610059575b600080fd5b610043610075565b60405161005091906101ad565b60405180910390f35b610073600480360381019061006e919061014b565b610099565b005b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055507f19e07a962d0e1cab8b7d4be06b77a7225ca7ebb744994cb6546fdd815590dc9d816040516101089190610192565b60405180910390a150565b6000808273ffffffffffffffffffffffffffffffffffffffff163b119050919050565b6000813590506101458161021e565b92915050565b60006020828403121561015d57600080fd5b600061016b84828501610136565b91505092915050565b61017d816101c8565b82525050565b61018c816101fa565b82525050565b60006020820190506101a76000830184610174565b92915050565b60006020820190506101c26000830184610183565b92915050565b60006101d3826101da565b9050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b60006102058261020c565b9050919050565b6000610217826101da565b9050919050565b610227816101c8565b811461023257600080fd5b5056fea264697066735822122021ea03c82e57e9ef47acefaa9096ea0fb4e63c2e5e377e1ee0f4d378501d129964736f6c63430008040033",
}

// StorageLockABI is the input ABI used to generate the binding from.
// Deprecated: Use StorageLockMetaData.ABI instead.
var StorageLockABI = StorageLockMetaData.ABI

// StorageLockBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use StorageLockMetaData.Bin instead.
var StorageLockBin = StorageLockMetaData.Bin

// DeployStorageLock deploys a new Ethereum contract, binding an instance of StorageLock to it.
func DeployStorageLock(auth *bind.TransactOpts, backend bind.ContractBackend, _marketplaceStorage common.Address) (common.Address, *types.Transaction, *StorageLock, error) {
	parsed, err := StorageLockMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(StorageLockBin), backend, _marketplaceStorage)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &StorageLock{StorageLockCaller: StorageLockCaller{contract: contract}, StorageLockTransactor: StorageLockTransactor{contract: contract}, StorageLockFilterer: StorageLockFilterer{contract: contract}}, nil
}

// StorageLock is an auto generated Go binding around an Ethereum contract.
type StorageLock struct {
	StorageLockCaller     // Read-only binding to the contract
	StorageLockTransactor // Write-only binding to the contract
	StorageLockFilterer   // Log filterer for contract events
}

// StorageLockCaller is an auto generated read-only Go binding around an Ethereum contract.
type StorageLockCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StorageLockTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StorageLockTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StorageLockFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StorageLockFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StorageLockSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StorageLockSession struct {
	Contract     *StorageLock      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StorageLockCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StorageLockCallerSession struct {
	Contract *StorageLockCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// StorageLockTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StorageLockTransactorSession struct {
	Contract     *StorageLockTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// StorageLockRaw is an auto generated low-level Go binding around an Ethereum contract.
type StorageLockRaw struct {
	Contract *StorageLock // Generic contract binding to access the raw methods on
}

// StorageLockCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StorageLockCallerRaw struct {
	Contract *StorageLockCaller // Generic read-only contract binding to access the raw methods on
}

// StorageLockTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StorageLockTransactorRaw struct {
	Contract *StorageLockTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStorageLock creates a new instance of StorageLock, bound to a specific deployed contract.
func NewStorageLock(address common.Address, backend bind.ContractBackend) (*StorageLock, error) {
	contract, err := bindStorageLock(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &StorageLock{StorageLockCaller: StorageLockCaller{contract: contract}, StorageLockTransactor: StorageLockTransactor{contract: contract}, StorageLockFilterer: StorageLockFilterer{contract: contract}}, nil
}

// NewStorageLockCaller creates a new read-only instance of StorageLock, bound to a specific deployed contract.
func NewStorageLockCaller(address common.Address, caller bind.ContractCaller) (*StorageLockCaller, error) {
	contract, err := bindStorageLock(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StorageLockCaller{contract: contract}, nil
}

// NewStorageLockTransactor creates a new write-only instance of StorageLock, bound to a specific deployed contract.
func NewStorageLockTransactor(address common.Address, transactor bind.ContractTransactor) (*StorageLockTransactor, error) {
	contract, err := bindStorageLock(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StorageLockTransactor{contract: contract}, nil
}

// NewStorageLockFilterer creates a new log filterer instance of StorageLock, bound to a specific deployed contract.
func NewStorageLockFilterer(address common.Address, filterer bind.ContractFilterer) (*StorageLockFilterer, error) {
	contract, err := bindStorageLock(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StorageLockFilterer{contract: contract}, nil
}

// bindStorageLock binds a generic wrapper to an already deployed contract.
func bindStorageLock(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(StorageLockABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StorageLock *StorageLockRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StorageLock.Contract.StorageLockCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StorageLock *StorageLockRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StorageLock.Contract.StorageLockTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StorageLock *StorageLockRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StorageLock.Contract.StorageLockTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StorageLock *StorageLockCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StorageLock.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StorageLock *StorageLockTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StorageLock.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StorageLock *StorageLockTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StorageLock.Contract.contract.Transact(opts, method, params...)
}

// MarketplaceStorage is a free data retrieval call binding the contract method 0x46b3aec6.
//
// Solidity: function marketplaceStorage() view returns(address)
func (_StorageLock *StorageLockCaller) MarketplaceStorage(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StorageLock.contract.Call(opts, &out, "marketplaceStorage")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// MarketplaceStorage is a free data retrieval call binding the contract method 0x46b3aec6.
//
// Solidity: function marketplaceStorage() view returns(address)
func (_StorageLock *StorageLockSession) MarketplaceStorage() (common.Address, error) {
	return _StorageLock.Contract.MarketplaceStorage(&_StorageLock.CallOpts)
}

// MarketplaceStorage is a free data retrieval call binding the contract method 0x46b3aec6.
//
// Solidity: function marketplaceStorage() view returns(address)
func (_StorageLock *StorageLockCallerSession) MarketplaceStorage() (common.Address, error) {
	return _StorageLock.Contract.MarketplaceStorage(&_StorageLock.CallOpts)
}

// UpdateStorageAddress is a paid mutator transaction binding the contract method 0x889e2129.
//
// Solidity: function updateStorageAddress(address _marketplaceStorage) returns()
func (_StorageLock *StorageLockTransactor) UpdateStorageAddress(opts *bind.TransactOpts, _marketplaceStorage common.Address) (*types.Transaction, error) {
	return _StorageLock.contract.Transact(opts, "updateStorageAddress", _marketplaceStorage)
}

// UpdateStorageAddress is a paid mutator transaction binding the contract method 0x889e2129.
//
// Solidity: function updateStorageAddress(address _marketplaceStorage) returns()
func (_StorageLock *StorageLockSession) UpdateStorageAddress(_marketplaceStorage common.Address) (*types.Transaction, error) {
	return _StorageLock.Contract.UpdateStorageAddress(&_StorageLock.TransactOpts, _marketplaceStorage)
}

// UpdateStorageAddress is a paid mutator transaction binding the contract method 0x889e2129.
//
// Solidity: function updateStorageAddress(address _marketplaceStorage) returns()
func (_StorageLock *StorageLockTransactorSession) UpdateStorageAddress(_marketplaceStorage common.Address) (*types.Transaction, error) {
	return _StorageLock.Contract.UpdateStorageAddress(&_StorageLock.TransactOpts, _marketplaceStorage)
}

// StorageLockMarketplaceStorageUpdatedIterator is returned from FilterMarketplaceStorageUpdated and is used to iterate over the raw logs and unpacked data for MarketplaceStorageUpdated events raised by the StorageLock contract.
type StorageLockMarketplaceStorageUpdatedIterator struct {
	Event *StorageLockMarketplaceStorageUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *StorageLockMarketplaceStorageUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StorageLockMarketplaceStorageUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(StorageLockMarketplaceStorageUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *StorageLockMarketplaceStorageUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StorageLockMarketplaceStorageUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StorageLockMarketplaceStorageUpdated represents a MarketplaceStorageUpdated event raised by the StorageLock contract.
type StorageLockMarketplaceStorageUpdated struct {
	MarketplaceStorage common.Address
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterMarketplaceStorageUpdated is a free log retrieval operation binding the contract event 0x19e07a962d0e1cab8b7d4be06b77a7225ca7ebb744994cb6546fdd815590dc9d.
//
// Solidity: event MarketplaceStorageUpdated(address _marketplaceStorage)
func (_StorageLock *StorageLockFilterer) FilterMarketplaceStorageUpdated(opts *bind.FilterOpts) (*StorageLockMarketplaceStorageUpdatedIterator, error) {

	logs, sub, err := _StorageLock.contract.FilterLogs(opts, "MarketplaceStorageUpdated")
	if err != nil {
		return nil, err
	}
	return &StorageLockMarketplaceStorageUpdatedIterator{contract: _StorageLock.contract, event: "MarketplaceStorageUpdated", logs: logs, sub: sub}, nil
}

// WatchMarketplaceStorageUpdated is a free log subscription operation binding the contract event 0x19e07a962d0e1cab8b7d4be06b77a7225ca7ebb744994cb6546fdd815590dc9d.
//
// Solidity: event MarketplaceStorageUpdated(address _marketplaceStorage)
func (_StorageLock *StorageLockFilterer) WatchMarketplaceStorageUpdated(opts *bind.WatchOpts, sink chan<- *StorageLockMarketplaceStorageUpdated) (event.Subscription, error) {

	logs, sub, err := _StorageLock.contract.WatchLogs(opts, "MarketplaceStorageUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StorageLockMarketplaceStorageUpdated)
				if err := _StorageLock.contract.UnpackLog(event, "MarketplaceStorageUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseMarketplaceStorageUpdated is a log parse operation binding the contract event 0x19e07a962d0e1cab8b7d4be06b77a7225ca7ebb744994cb6546fdd815590dc9d.
//
// Solidity: event MarketplaceStorageUpdated(address _marketplaceStorage)
func (_StorageLock *StorageLockFilterer) ParseMarketplaceStorageUpdated(log types.Log) (*StorageLockMarketplaceStorageUpdated, error) {
	event := new(StorageLockMarketplaceStorageUpdated)
	if err := _StorageLock.contract.UnpackLog(event, "MarketplaceStorageUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
