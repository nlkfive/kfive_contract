// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package IPublicAuction

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

// IPublicAuctionMetaData contains all meta data concerning the IPublicAuction contract.
var IPublicAuctionMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"bidder\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"publicAuctionId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"bidValue\",\"type\":\"uint256\"}],\"name\":\"PublicAuctionBidSuccessful\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"canceller\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"publicAuctionId\",\"type\":\"bytes32\"}],\"name\":\"PublicAuctionCancelled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"nftAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"publicAuctionId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"assetId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"biddingEnd\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"startPriceInWei\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"minIncrement\",\"type\":\"uint256\"}],\"name\":\"PublicAuctionCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"publicAuctionId\",\"type\":\"bytes32\"}],\"name\":\"PublicAuctionEnded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"bidder\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"publicAuctionId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"PublicAuctionRefund\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"buyer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"publicAuctionId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalPrice\",\"type\":\"uint256\"}],\"name\":\"PublicAuctionSuccessful\",\"type\":\"event\"}]",
}

// IPublicAuctionABI is the input ABI used to generate the binding from.
// Deprecated: Use IPublicAuctionMetaData.ABI instead.
var IPublicAuctionABI = IPublicAuctionMetaData.ABI

// IPublicAuction is an auto generated Go binding around an Ethereum contract.
type IPublicAuction struct {
	IPublicAuctionCaller     // Read-only binding to the contract
	IPublicAuctionTransactor // Write-only binding to the contract
	IPublicAuctionFilterer   // Log filterer for contract events
}

// IPublicAuctionCaller is an auto generated read-only Go binding around an Ethereum contract.
type IPublicAuctionCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IPublicAuctionTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IPublicAuctionTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IPublicAuctionFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IPublicAuctionFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IPublicAuctionSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IPublicAuctionSession struct {
	Contract     *IPublicAuction   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IPublicAuctionCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IPublicAuctionCallerSession struct {
	Contract *IPublicAuctionCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// IPublicAuctionTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IPublicAuctionTransactorSession struct {
	Contract     *IPublicAuctionTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// IPublicAuctionRaw is an auto generated low-level Go binding around an Ethereum contract.
type IPublicAuctionRaw struct {
	Contract *IPublicAuction // Generic contract binding to access the raw methods on
}

// IPublicAuctionCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IPublicAuctionCallerRaw struct {
	Contract *IPublicAuctionCaller // Generic read-only contract binding to access the raw methods on
}

// IPublicAuctionTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IPublicAuctionTransactorRaw struct {
	Contract *IPublicAuctionTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIPublicAuction creates a new instance of IPublicAuction, bound to a specific deployed contract.
func NewIPublicAuction(address common.Address, backend bind.ContractBackend) (*IPublicAuction, error) {
	contract, err := bindIPublicAuction(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IPublicAuction{IPublicAuctionCaller: IPublicAuctionCaller{contract: contract}, IPublicAuctionTransactor: IPublicAuctionTransactor{contract: contract}, IPublicAuctionFilterer: IPublicAuctionFilterer{contract: contract}}, nil
}

// NewIPublicAuctionCaller creates a new read-only instance of IPublicAuction, bound to a specific deployed contract.
func NewIPublicAuctionCaller(address common.Address, caller bind.ContractCaller) (*IPublicAuctionCaller, error) {
	contract, err := bindIPublicAuction(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IPublicAuctionCaller{contract: contract}, nil
}

// NewIPublicAuctionTransactor creates a new write-only instance of IPublicAuction, bound to a specific deployed contract.
func NewIPublicAuctionTransactor(address common.Address, transactor bind.ContractTransactor) (*IPublicAuctionTransactor, error) {
	contract, err := bindIPublicAuction(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IPublicAuctionTransactor{contract: contract}, nil
}

// NewIPublicAuctionFilterer creates a new log filterer instance of IPublicAuction, bound to a specific deployed contract.
func NewIPublicAuctionFilterer(address common.Address, filterer bind.ContractFilterer) (*IPublicAuctionFilterer, error) {
	contract, err := bindIPublicAuction(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IPublicAuctionFilterer{contract: contract}, nil
}

// bindIPublicAuction binds a generic wrapper to an already deployed contract.
func bindIPublicAuction(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IPublicAuctionABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IPublicAuction *IPublicAuctionRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IPublicAuction.Contract.IPublicAuctionCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IPublicAuction *IPublicAuctionRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IPublicAuction.Contract.IPublicAuctionTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IPublicAuction *IPublicAuctionRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IPublicAuction.Contract.IPublicAuctionTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IPublicAuction *IPublicAuctionCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IPublicAuction.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IPublicAuction *IPublicAuctionTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IPublicAuction.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IPublicAuction *IPublicAuctionTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IPublicAuction.Contract.contract.Transact(opts, method, params...)
}

// IPublicAuctionPublicAuctionBidSuccessfulIterator is returned from FilterPublicAuctionBidSuccessful and is used to iterate over the raw logs and unpacked data for PublicAuctionBidSuccessful events raised by the IPublicAuction contract.
type IPublicAuctionPublicAuctionBidSuccessfulIterator struct {
	Event *IPublicAuctionPublicAuctionBidSuccessful // Event containing the contract specifics and raw log

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
func (it *IPublicAuctionPublicAuctionBidSuccessfulIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IPublicAuctionPublicAuctionBidSuccessful)
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
		it.Event = new(IPublicAuctionPublicAuctionBidSuccessful)
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
func (it *IPublicAuctionPublicAuctionBidSuccessfulIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IPublicAuctionPublicAuctionBidSuccessfulIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IPublicAuctionPublicAuctionBidSuccessful represents a PublicAuctionBidSuccessful event raised by the IPublicAuction contract.
type IPublicAuctionPublicAuctionBidSuccessful struct {
	Bidder          common.Address
	PublicAuctionId [32]byte
	BidValue        *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterPublicAuctionBidSuccessful is a free log retrieval operation binding the contract event 0x9f541988b3aac74ac43a6f4b330ceed668244b74461b90850d199a38a172f7c2.
//
// Solidity: event PublicAuctionBidSuccessful(address bidder, bytes32 publicAuctionId, uint256 bidValue)
func (_IPublicAuction *IPublicAuctionFilterer) FilterPublicAuctionBidSuccessful(opts *bind.FilterOpts) (*IPublicAuctionPublicAuctionBidSuccessfulIterator, error) {

	logs, sub, err := _IPublicAuction.contract.FilterLogs(opts, "PublicAuctionBidSuccessful")
	if err != nil {
		return nil, err
	}
	return &IPublicAuctionPublicAuctionBidSuccessfulIterator{contract: _IPublicAuction.contract, event: "PublicAuctionBidSuccessful", logs: logs, sub: sub}, nil
}

// WatchPublicAuctionBidSuccessful is a free log subscription operation binding the contract event 0x9f541988b3aac74ac43a6f4b330ceed668244b74461b90850d199a38a172f7c2.
//
// Solidity: event PublicAuctionBidSuccessful(address bidder, bytes32 publicAuctionId, uint256 bidValue)
func (_IPublicAuction *IPublicAuctionFilterer) WatchPublicAuctionBidSuccessful(opts *bind.WatchOpts, sink chan<- *IPublicAuctionPublicAuctionBidSuccessful) (event.Subscription, error) {

	logs, sub, err := _IPublicAuction.contract.WatchLogs(opts, "PublicAuctionBidSuccessful")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IPublicAuctionPublicAuctionBidSuccessful)
				if err := _IPublicAuction.contract.UnpackLog(event, "PublicAuctionBidSuccessful", log); err != nil {
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

// ParsePublicAuctionBidSuccessful is a log parse operation binding the contract event 0x9f541988b3aac74ac43a6f4b330ceed668244b74461b90850d199a38a172f7c2.
//
// Solidity: event PublicAuctionBidSuccessful(address bidder, bytes32 publicAuctionId, uint256 bidValue)
func (_IPublicAuction *IPublicAuctionFilterer) ParsePublicAuctionBidSuccessful(log types.Log) (*IPublicAuctionPublicAuctionBidSuccessful, error) {
	event := new(IPublicAuctionPublicAuctionBidSuccessful)
	if err := _IPublicAuction.contract.UnpackLog(event, "PublicAuctionBidSuccessful", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IPublicAuctionPublicAuctionCancelledIterator is returned from FilterPublicAuctionCancelled and is used to iterate over the raw logs and unpacked data for PublicAuctionCancelled events raised by the IPublicAuction contract.
type IPublicAuctionPublicAuctionCancelledIterator struct {
	Event *IPublicAuctionPublicAuctionCancelled // Event containing the contract specifics and raw log

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
func (it *IPublicAuctionPublicAuctionCancelledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IPublicAuctionPublicAuctionCancelled)
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
		it.Event = new(IPublicAuctionPublicAuctionCancelled)
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
func (it *IPublicAuctionPublicAuctionCancelledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IPublicAuctionPublicAuctionCancelledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IPublicAuctionPublicAuctionCancelled represents a PublicAuctionCancelled event raised by the IPublicAuction contract.
type IPublicAuctionPublicAuctionCancelled struct {
	Canceller       common.Address
	PublicAuctionId [32]byte
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterPublicAuctionCancelled is a free log retrieval operation binding the contract event 0x87cbb8dceaaacfbe7c9d99abf2a9c85f0249d860e329f60869079b204c7ad00d.
//
// Solidity: event PublicAuctionCancelled(address canceller, bytes32 publicAuctionId)
func (_IPublicAuction *IPublicAuctionFilterer) FilterPublicAuctionCancelled(opts *bind.FilterOpts) (*IPublicAuctionPublicAuctionCancelledIterator, error) {

	logs, sub, err := _IPublicAuction.contract.FilterLogs(opts, "PublicAuctionCancelled")
	if err != nil {
		return nil, err
	}
	return &IPublicAuctionPublicAuctionCancelledIterator{contract: _IPublicAuction.contract, event: "PublicAuctionCancelled", logs: logs, sub: sub}, nil
}

// WatchPublicAuctionCancelled is a free log subscription operation binding the contract event 0x87cbb8dceaaacfbe7c9d99abf2a9c85f0249d860e329f60869079b204c7ad00d.
//
// Solidity: event PublicAuctionCancelled(address canceller, bytes32 publicAuctionId)
func (_IPublicAuction *IPublicAuctionFilterer) WatchPublicAuctionCancelled(opts *bind.WatchOpts, sink chan<- *IPublicAuctionPublicAuctionCancelled) (event.Subscription, error) {

	logs, sub, err := _IPublicAuction.contract.WatchLogs(opts, "PublicAuctionCancelled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IPublicAuctionPublicAuctionCancelled)
				if err := _IPublicAuction.contract.UnpackLog(event, "PublicAuctionCancelled", log); err != nil {
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

// ParsePublicAuctionCancelled is a log parse operation binding the contract event 0x87cbb8dceaaacfbe7c9d99abf2a9c85f0249d860e329f60869079b204c7ad00d.
//
// Solidity: event PublicAuctionCancelled(address canceller, bytes32 publicAuctionId)
func (_IPublicAuction *IPublicAuctionFilterer) ParsePublicAuctionCancelled(log types.Log) (*IPublicAuctionPublicAuctionCancelled, error) {
	event := new(IPublicAuctionPublicAuctionCancelled)
	if err := _IPublicAuction.contract.UnpackLog(event, "PublicAuctionCancelled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IPublicAuctionPublicAuctionCreatedIterator is returned from FilterPublicAuctionCreated and is used to iterate over the raw logs and unpacked data for PublicAuctionCreated events raised by the IPublicAuction contract.
type IPublicAuctionPublicAuctionCreatedIterator struct {
	Event *IPublicAuctionPublicAuctionCreated // Event containing the contract specifics and raw log

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
func (it *IPublicAuctionPublicAuctionCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IPublicAuctionPublicAuctionCreated)
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
		it.Event = new(IPublicAuctionPublicAuctionCreated)
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
func (it *IPublicAuctionPublicAuctionCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IPublicAuctionPublicAuctionCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IPublicAuctionPublicAuctionCreated represents a PublicAuctionCreated event raised by the IPublicAuction contract.
type IPublicAuctionPublicAuctionCreated struct {
	Seller          common.Address
	NftAddress      common.Address
	PublicAuctionId [32]byte
	AssetId         *big.Int
	BiddingEnd      *big.Int
	StartPriceInWei *big.Int
	MinIncrement    *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterPublicAuctionCreated is a free log retrieval operation binding the contract event 0xbdca6148a24d8e6e2d4ced0a6e168095869e61ea26b82332620abe8ba3ba5bd9.
//
// Solidity: event PublicAuctionCreated(address seller, address nftAddress, bytes32 publicAuctionId, uint256 assetId, uint256 biddingEnd, uint256 startPriceInWei, uint256 minIncrement)
func (_IPublicAuction *IPublicAuctionFilterer) FilterPublicAuctionCreated(opts *bind.FilterOpts) (*IPublicAuctionPublicAuctionCreatedIterator, error) {

	logs, sub, err := _IPublicAuction.contract.FilterLogs(opts, "PublicAuctionCreated")
	if err != nil {
		return nil, err
	}
	return &IPublicAuctionPublicAuctionCreatedIterator{contract: _IPublicAuction.contract, event: "PublicAuctionCreated", logs: logs, sub: sub}, nil
}

// WatchPublicAuctionCreated is a free log subscription operation binding the contract event 0xbdca6148a24d8e6e2d4ced0a6e168095869e61ea26b82332620abe8ba3ba5bd9.
//
// Solidity: event PublicAuctionCreated(address seller, address nftAddress, bytes32 publicAuctionId, uint256 assetId, uint256 biddingEnd, uint256 startPriceInWei, uint256 minIncrement)
func (_IPublicAuction *IPublicAuctionFilterer) WatchPublicAuctionCreated(opts *bind.WatchOpts, sink chan<- *IPublicAuctionPublicAuctionCreated) (event.Subscription, error) {

	logs, sub, err := _IPublicAuction.contract.WatchLogs(opts, "PublicAuctionCreated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IPublicAuctionPublicAuctionCreated)
				if err := _IPublicAuction.contract.UnpackLog(event, "PublicAuctionCreated", log); err != nil {
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

// ParsePublicAuctionCreated is a log parse operation binding the contract event 0xbdca6148a24d8e6e2d4ced0a6e168095869e61ea26b82332620abe8ba3ba5bd9.
//
// Solidity: event PublicAuctionCreated(address seller, address nftAddress, bytes32 publicAuctionId, uint256 assetId, uint256 biddingEnd, uint256 startPriceInWei, uint256 minIncrement)
func (_IPublicAuction *IPublicAuctionFilterer) ParsePublicAuctionCreated(log types.Log) (*IPublicAuctionPublicAuctionCreated, error) {
	event := new(IPublicAuctionPublicAuctionCreated)
	if err := _IPublicAuction.contract.UnpackLog(event, "PublicAuctionCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IPublicAuctionPublicAuctionEndedIterator is returned from FilterPublicAuctionEnded and is used to iterate over the raw logs and unpacked data for PublicAuctionEnded events raised by the IPublicAuction contract.
type IPublicAuctionPublicAuctionEndedIterator struct {
	Event *IPublicAuctionPublicAuctionEnded // Event containing the contract specifics and raw log

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
func (it *IPublicAuctionPublicAuctionEndedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IPublicAuctionPublicAuctionEnded)
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
		it.Event = new(IPublicAuctionPublicAuctionEnded)
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
func (it *IPublicAuctionPublicAuctionEndedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IPublicAuctionPublicAuctionEndedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IPublicAuctionPublicAuctionEnded represents a PublicAuctionEnded event raised by the IPublicAuction contract.
type IPublicAuctionPublicAuctionEnded struct {
	PublicAuctionId [32]byte
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterPublicAuctionEnded is a free log retrieval operation binding the contract event 0x156c2754a62667a51625de4ceb04df6640d97d06de8c89bd0bbb33307f144e42.
//
// Solidity: event PublicAuctionEnded(bytes32 publicAuctionId)
func (_IPublicAuction *IPublicAuctionFilterer) FilterPublicAuctionEnded(opts *bind.FilterOpts) (*IPublicAuctionPublicAuctionEndedIterator, error) {

	logs, sub, err := _IPublicAuction.contract.FilterLogs(opts, "PublicAuctionEnded")
	if err != nil {
		return nil, err
	}
	return &IPublicAuctionPublicAuctionEndedIterator{contract: _IPublicAuction.contract, event: "PublicAuctionEnded", logs: logs, sub: sub}, nil
}

// WatchPublicAuctionEnded is a free log subscription operation binding the contract event 0x156c2754a62667a51625de4ceb04df6640d97d06de8c89bd0bbb33307f144e42.
//
// Solidity: event PublicAuctionEnded(bytes32 publicAuctionId)
func (_IPublicAuction *IPublicAuctionFilterer) WatchPublicAuctionEnded(opts *bind.WatchOpts, sink chan<- *IPublicAuctionPublicAuctionEnded) (event.Subscription, error) {

	logs, sub, err := _IPublicAuction.contract.WatchLogs(opts, "PublicAuctionEnded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IPublicAuctionPublicAuctionEnded)
				if err := _IPublicAuction.contract.UnpackLog(event, "PublicAuctionEnded", log); err != nil {
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

// ParsePublicAuctionEnded is a log parse operation binding the contract event 0x156c2754a62667a51625de4ceb04df6640d97d06de8c89bd0bbb33307f144e42.
//
// Solidity: event PublicAuctionEnded(bytes32 publicAuctionId)
func (_IPublicAuction *IPublicAuctionFilterer) ParsePublicAuctionEnded(log types.Log) (*IPublicAuctionPublicAuctionEnded, error) {
	event := new(IPublicAuctionPublicAuctionEnded)
	if err := _IPublicAuction.contract.UnpackLog(event, "PublicAuctionEnded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IPublicAuctionPublicAuctionRefundIterator is returned from FilterPublicAuctionRefund and is used to iterate over the raw logs and unpacked data for PublicAuctionRefund events raised by the IPublicAuction contract.
type IPublicAuctionPublicAuctionRefundIterator struct {
	Event *IPublicAuctionPublicAuctionRefund // Event containing the contract specifics and raw log

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
func (it *IPublicAuctionPublicAuctionRefundIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IPublicAuctionPublicAuctionRefund)
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
		it.Event = new(IPublicAuctionPublicAuctionRefund)
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
func (it *IPublicAuctionPublicAuctionRefundIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IPublicAuctionPublicAuctionRefundIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IPublicAuctionPublicAuctionRefund represents a PublicAuctionRefund event raised by the IPublicAuction contract.
type IPublicAuctionPublicAuctionRefund struct {
	Bidder          common.Address
	PublicAuctionId [32]byte
	Value           *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterPublicAuctionRefund is a free log retrieval operation binding the contract event 0x4fac96064e645402e0d4854b9549caba0169e3252c0e08a305eb60f498c88911.
//
// Solidity: event PublicAuctionRefund(address bidder, bytes32 publicAuctionId, uint256 value)
func (_IPublicAuction *IPublicAuctionFilterer) FilterPublicAuctionRefund(opts *bind.FilterOpts) (*IPublicAuctionPublicAuctionRefundIterator, error) {

	logs, sub, err := _IPublicAuction.contract.FilterLogs(opts, "PublicAuctionRefund")
	if err != nil {
		return nil, err
	}
	return &IPublicAuctionPublicAuctionRefundIterator{contract: _IPublicAuction.contract, event: "PublicAuctionRefund", logs: logs, sub: sub}, nil
}

// WatchPublicAuctionRefund is a free log subscription operation binding the contract event 0x4fac96064e645402e0d4854b9549caba0169e3252c0e08a305eb60f498c88911.
//
// Solidity: event PublicAuctionRefund(address bidder, bytes32 publicAuctionId, uint256 value)
func (_IPublicAuction *IPublicAuctionFilterer) WatchPublicAuctionRefund(opts *bind.WatchOpts, sink chan<- *IPublicAuctionPublicAuctionRefund) (event.Subscription, error) {

	logs, sub, err := _IPublicAuction.contract.WatchLogs(opts, "PublicAuctionRefund")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IPublicAuctionPublicAuctionRefund)
				if err := _IPublicAuction.contract.UnpackLog(event, "PublicAuctionRefund", log); err != nil {
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

// ParsePublicAuctionRefund is a log parse operation binding the contract event 0x4fac96064e645402e0d4854b9549caba0169e3252c0e08a305eb60f498c88911.
//
// Solidity: event PublicAuctionRefund(address bidder, bytes32 publicAuctionId, uint256 value)
func (_IPublicAuction *IPublicAuctionFilterer) ParsePublicAuctionRefund(log types.Log) (*IPublicAuctionPublicAuctionRefund, error) {
	event := new(IPublicAuctionPublicAuctionRefund)
	if err := _IPublicAuction.contract.UnpackLog(event, "PublicAuctionRefund", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IPublicAuctionPublicAuctionSuccessfulIterator is returned from FilterPublicAuctionSuccessful and is used to iterate over the raw logs and unpacked data for PublicAuctionSuccessful events raised by the IPublicAuction contract.
type IPublicAuctionPublicAuctionSuccessfulIterator struct {
	Event *IPublicAuctionPublicAuctionSuccessful // Event containing the contract specifics and raw log

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
func (it *IPublicAuctionPublicAuctionSuccessfulIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IPublicAuctionPublicAuctionSuccessful)
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
		it.Event = new(IPublicAuctionPublicAuctionSuccessful)
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
func (it *IPublicAuctionPublicAuctionSuccessfulIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IPublicAuctionPublicAuctionSuccessfulIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IPublicAuctionPublicAuctionSuccessful represents a PublicAuctionSuccessful event raised by the IPublicAuction contract.
type IPublicAuctionPublicAuctionSuccessful struct {
	Seller          common.Address
	Buyer           common.Address
	PublicAuctionId [32]byte
	TotalPrice      *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterPublicAuctionSuccessful is a free log retrieval operation binding the contract event 0xd5133f9e5285bd0635ccfe53ae8978b96ef34a7bf16a0a696a411dc669322cb6.
//
// Solidity: event PublicAuctionSuccessful(address seller, address buyer, bytes32 publicAuctionId, uint256 totalPrice)
func (_IPublicAuction *IPublicAuctionFilterer) FilterPublicAuctionSuccessful(opts *bind.FilterOpts) (*IPublicAuctionPublicAuctionSuccessfulIterator, error) {

	logs, sub, err := _IPublicAuction.contract.FilterLogs(opts, "PublicAuctionSuccessful")
	if err != nil {
		return nil, err
	}
	return &IPublicAuctionPublicAuctionSuccessfulIterator{contract: _IPublicAuction.contract, event: "PublicAuctionSuccessful", logs: logs, sub: sub}, nil
}

// WatchPublicAuctionSuccessful is a free log subscription operation binding the contract event 0xd5133f9e5285bd0635ccfe53ae8978b96ef34a7bf16a0a696a411dc669322cb6.
//
// Solidity: event PublicAuctionSuccessful(address seller, address buyer, bytes32 publicAuctionId, uint256 totalPrice)
func (_IPublicAuction *IPublicAuctionFilterer) WatchPublicAuctionSuccessful(opts *bind.WatchOpts, sink chan<- *IPublicAuctionPublicAuctionSuccessful) (event.Subscription, error) {

	logs, sub, err := _IPublicAuction.contract.WatchLogs(opts, "PublicAuctionSuccessful")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IPublicAuctionPublicAuctionSuccessful)
				if err := _IPublicAuction.contract.UnpackLog(event, "PublicAuctionSuccessful", log); err != nil {
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

// ParsePublicAuctionSuccessful is a log parse operation binding the contract event 0xd5133f9e5285bd0635ccfe53ae8978b96ef34a7bf16a0a696a411dc669322cb6.
//
// Solidity: event PublicAuctionSuccessful(address seller, address buyer, bytes32 publicAuctionId, uint256 totalPrice)
func (_IPublicAuction *IPublicAuctionFilterer) ParsePublicAuctionSuccessful(log types.Log) (*IPublicAuctionPublicAuctionSuccessful, error) {
	event := new(IPublicAuctionPublicAuctionSuccessful)
	if err := _IPublicAuction.contract.UnpackLog(event, "PublicAuctionSuccessful", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
