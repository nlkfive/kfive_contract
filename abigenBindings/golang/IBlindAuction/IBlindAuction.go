// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package IBlindAuction

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

// IBlindAuctionMetaData contains all meta data concerning the IBlindAuction contract.
var IBlindAuctionMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"bidder\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"blindAuctionId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"blindedBid\",\"type\":\"bytes32\"}],\"name\":\"BlindAuctionBidSuccessful\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"canceller\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"blindAuctionId\",\"type\":\"bytes32\"}],\"name\":\"BlindAuctionCancelled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"nftAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"blindAuctionId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"assetId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"biddingEnd\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"revealEnd\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"startPriceInWei\",\"type\":\"uint256\"}],\"name\":\"BlindAuctionCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"blindAuctionId\",\"type\":\"bytes32\"}],\"name\":\"BlindAuctionEnded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"bidder\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"blindAuctionId\",\"type\":\"bytes32\"}],\"name\":\"BlindAuctionRefund\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"buyer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"blindAuctionId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalPrice\",\"type\":\"uint256\"}],\"name\":\"BlindAuctionSuccessful\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"blindAuctionId\",\"type\":\"bytes32\"}],\"name\":\"RevealSuccessful\",\"type\":\"event\"}]",
}

// IBlindAuctionABI is the input ABI used to generate the binding from.
// Deprecated: Use IBlindAuctionMetaData.ABI instead.
var IBlindAuctionABI = IBlindAuctionMetaData.ABI

// IBlindAuction is an auto generated Go binding around an Ethereum contract.
type IBlindAuction struct {
	IBlindAuctionCaller     // Read-only binding to the contract
	IBlindAuctionTransactor // Write-only binding to the contract
	IBlindAuctionFilterer   // Log filterer for contract events
}

// IBlindAuctionCaller is an auto generated read-only Go binding around an Ethereum contract.
type IBlindAuctionCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IBlindAuctionTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IBlindAuctionTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IBlindAuctionFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IBlindAuctionFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IBlindAuctionSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IBlindAuctionSession struct {
	Contract     *IBlindAuction    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IBlindAuctionCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IBlindAuctionCallerSession struct {
	Contract *IBlindAuctionCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// IBlindAuctionTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IBlindAuctionTransactorSession struct {
	Contract     *IBlindAuctionTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// IBlindAuctionRaw is an auto generated low-level Go binding around an Ethereum contract.
type IBlindAuctionRaw struct {
	Contract *IBlindAuction // Generic contract binding to access the raw methods on
}

// IBlindAuctionCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IBlindAuctionCallerRaw struct {
	Contract *IBlindAuctionCaller // Generic read-only contract binding to access the raw methods on
}

// IBlindAuctionTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IBlindAuctionTransactorRaw struct {
	Contract *IBlindAuctionTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIBlindAuction creates a new instance of IBlindAuction, bound to a specific deployed contract.
func NewIBlindAuction(address common.Address, backend bind.ContractBackend) (*IBlindAuction, error) {
	contract, err := bindIBlindAuction(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IBlindAuction{IBlindAuctionCaller: IBlindAuctionCaller{contract: contract}, IBlindAuctionTransactor: IBlindAuctionTransactor{contract: contract}, IBlindAuctionFilterer: IBlindAuctionFilterer{contract: contract}}, nil
}

// NewIBlindAuctionCaller creates a new read-only instance of IBlindAuction, bound to a specific deployed contract.
func NewIBlindAuctionCaller(address common.Address, caller bind.ContractCaller) (*IBlindAuctionCaller, error) {
	contract, err := bindIBlindAuction(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IBlindAuctionCaller{contract: contract}, nil
}

// NewIBlindAuctionTransactor creates a new write-only instance of IBlindAuction, bound to a specific deployed contract.
func NewIBlindAuctionTransactor(address common.Address, transactor bind.ContractTransactor) (*IBlindAuctionTransactor, error) {
	contract, err := bindIBlindAuction(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IBlindAuctionTransactor{contract: contract}, nil
}

// NewIBlindAuctionFilterer creates a new log filterer instance of IBlindAuction, bound to a specific deployed contract.
func NewIBlindAuctionFilterer(address common.Address, filterer bind.ContractFilterer) (*IBlindAuctionFilterer, error) {
	contract, err := bindIBlindAuction(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IBlindAuctionFilterer{contract: contract}, nil
}

// bindIBlindAuction binds a generic wrapper to an already deployed contract.
func bindIBlindAuction(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IBlindAuctionABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IBlindAuction *IBlindAuctionRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IBlindAuction.Contract.IBlindAuctionCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IBlindAuction *IBlindAuctionRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IBlindAuction.Contract.IBlindAuctionTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IBlindAuction *IBlindAuctionRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IBlindAuction.Contract.IBlindAuctionTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IBlindAuction *IBlindAuctionCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IBlindAuction.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IBlindAuction *IBlindAuctionTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IBlindAuction.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IBlindAuction *IBlindAuctionTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IBlindAuction.Contract.contract.Transact(opts, method, params...)
}

// IBlindAuctionBlindAuctionBidSuccessfulIterator is returned from FilterBlindAuctionBidSuccessful and is used to iterate over the raw logs and unpacked data for BlindAuctionBidSuccessful events raised by the IBlindAuction contract.
type IBlindAuctionBlindAuctionBidSuccessfulIterator struct {
	Event *IBlindAuctionBlindAuctionBidSuccessful // Event containing the contract specifics and raw log

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
func (it *IBlindAuctionBlindAuctionBidSuccessfulIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IBlindAuctionBlindAuctionBidSuccessful)
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
		it.Event = new(IBlindAuctionBlindAuctionBidSuccessful)
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
func (it *IBlindAuctionBlindAuctionBidSuccessfulIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IBlindAuctionBlindAuctionBidSuccessfulIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IBlindAuctionBlindAuctionBidSuccessful represents a BlindAuctionBidSuccessful event raised by the IBlindAuction contract.
type IBlindAuctionBlindAuctionBidSuccessful struct {
	Bidder         common.Address
	BlindAuctionId [32]byte
	BlindedBid     [32]byte
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterBlindAuctionBidSuccessful is a free log retrieval operation binding the contract event 0x42412bd3e1b349d0fc9c3518c85d09f611398d80842c5d3e36449b76dca01b6b.
//
// Solidity: event BlindAuctionBidSuccessful(address bidder, bytes32 blindAuctionId, bytes32 blindedBid)
func (_IBlindAuction *IBlindAuctionFilterer) FilterBlindAuctionBidSuccessful(opts *bind.FilterOpts) (*IBlindAuctionBlindAuctionBidSuccessfulIterator, error) {

	logs, sub, err := _IBlindAuction.contract.FilterLogs(opts, "BlindAuctionBidSuccessful")
	if err != nil {
		return nil, err
	}
	return &IBlindAuctionBlindAuctionBidSuccessfulIterator{contract: _IBlindAuction.contract, event: "BlindAuctionBidSuccessful", logs: logs, sub: sub}, nil
}

// WatchBlindAuctionBidSuccessful is a free log subscription operation binding the contract event 0x42412bd3e1b349d0fc9c3518c85d09f611398d80842c5d3e36449b76dca01b6b.
//
// Solidity: event BlindAuctionBidSuccessful(address bidder, bytes32 blindAuctionId, bytes32 blindedBid)
func (_IBlindAuction *IBlindAuctionFilterer) WatchBlindAuctionBidSuccessful(opts *bind.WatchOpts, sink chan<- *IBlindAuctionBlindAuctionBidSuccessful) (event.Subscription, error) {

	logs, sub, err := _IBlindAuction.contract.WatchLogs(opts, "BlindAuctionBidSuccessful")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IBlindAuctionBlindAuctionBidSuccessful)
				if err := _IBlindAuction.contract.UnpackLog(event, "BlindAuctionBidSuccessful", log); err != nil {
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

// ParseBlindAuctionBidSuccessful is a log parse operation binding the contract event 0x42412bd3e1b349d0fc9c3518c85d09f611398d80842c5d3e36449b76dca01b6b.
//
// Solidity: event BlindAuctionBidSuccessful(address bidder, bytes32 blindAuctionId, bytes32 blindedBid)
func (_IBlindAuction *IBlindAuctionFilterer) ParseBlindAuctionBidSuccessful(log types.Log) (*IBlindAuctionBlindAuctionBidSuccessful, error) {
	event := new(IBlindAuctionBlindAuctionBidSuccessful)
	if err := _IBlindAuction.contract.UnpackLog(event, "BlindAuctionBidSuccessful", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IBlindAuctionBlindAuctionCancelledIterator is returned from FilterBlindAuctionCancelled and is used to iterate over the raw logs and unpacked data for BlindAuctionCancelled events raised by the IBlindAuction contract.
type IBlindAuctionBlindAuctionCancelledIterator struct {
	Event *IBlindAuctionBlindAuctionCancelled // Event containing the contract specifics and raw log

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
func (it *IBlindAuctionBlindAuctionCancelledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IBlindAuctionBlindAuctionCancelled)
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
		it.Event = new(IBlindAuctionBlindAuctionCancelled)
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
func (it *IBlindAuctionBlindAuctionCancelledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IBlindAuctionBlindAuctionCancelledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IBlindAuctionBlindAuctionCancelled represents a BlindAuctionCancelled event raised by the IBlindAuction contract.
type IBlindAuctionBlindAuctionCancelled struct {
	Canceller      common.Address
	BlindAuctionId [32]byte
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterBlindAuctionCancelled is a free log retrieval operation binding the contract event 0x620e9f21de284210359a265815fc84d3f6f4686df7a3841c7ef048a50ea80216.
//
// Solidity: event BlindAuctionCancelled(address canceller, bytes32 blindAuctionId)
func (_IBlindAuction *IBlindAuctionFilterer) FilterBlindAuctionCancelled(opts *bind.FilterOpts) (*IBlindAuctionBlindAuctionCancelledIterator, error) {

	logs, sub, err := _IBlindAuction.contract.FilterLogs(opts, "BlindAuctionCancelled")
	if err != nil {
		return nil, err
	}
	return &IBlindAuctionBlindAuctionCancelledIterator{contract: _IBlindAuction.contract, event: "BlindAuctionCancelled", logs: logs, sub: sub}, nil
}

// WatchBlindAuctionCancelled is a free log subscription operation binding the contract event 0x620e9f21de284210359a265815fc84d3f6f4686df7a3841c7ef048a50ea80216.
//
// Solidity: event BlindAuctionCancelled(address canceller, bytes32 blindAuctionId)
func (_IBlindAuction *IBlindAuctionFilterer) WatchBlindAuctionCancelled(opts *bind.WatchOpts, sink chan<- *IBlindAuctionBlindAuctionCancelled) (event.Subscription, error) {

	logs, sub, err := _IBlindAuction.contract.WatchLogs(opts, "BlindAuctionCancelled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IBlindAuctionBlindAuctionCancelled)
				if err := _IBlindAuction.contract.UnpackLog(event, "BlindAuctionCancelled", log); err != nil {
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

// ParseBlindAuctionCancelled is a log parse operation binding the contract event 0x620e9f21de284210359a265815fc84d3f6f4686df7a3841c7ef048a50ea80216.
//
// Solidity: event BlindAuctionCancelled(address canceller, bytes32 blindAuctionId)
func (_IBlindAuction *IBlindAuctionFilterer) ParseBlindAuctionCancelled(log types.Log) (*IBlindAuctionBlindAuctionCancelled, error) {
	event := new(IBlindAuctionBlindAuctionCancelled)
	if err := _IBlindAuction.contract.UnpackLog(event, "BlindAuctionCancelled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IBlindAuctionBlindAuctionCreatedIterator is returned from FilterBlindAuctionCreated and is used to iterate over the raw logs and unpacked data for BlindAuctionCreated events raised by the IBlindAuction contract.
type IBlindAuctionBlindAuctionCreatedIterator struct {
	Event *IBlindAuctionBlindAuctionCreated // Event containing the contract specifics and raw log

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
func (it *IBlindAuctionBlindAuctionCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IBlindAuctionBlindAuctionCreated)
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
		it.Event = new(IBlindAuctionBlindAuctionCreated)
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
func (it *IBlindAuctionBlindAuctionCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IBlindAuctionBlindAuctionCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IBlindAuctionBlindAuctionCreated represents a BlindAuctionCreated event raised by the IBlindAuction contract.
type IBlindAuctionBlindAuctionCreated struct {
	Seller          common.Address
	NftAddress      common.Address
	BlindAuctionId  [32]byte
	AssetId         *big.Int
	BiddingEnd      *big.Int
	RevealEnd       *big.Int
	StartPriceInWei *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterBlindAuctionCreated is a free log retrieval operation binding the contract event 0xe64fd4ba0702f46ae298a3be33f16f223a5e06556977c42dee952f976abfed83.
//
// Solidity: event BlindAuctionCreated(address seller, address nftAddress, bytes32 blindAuctionId, uint256 assetId, uint256 biddingEnd, uint256 revealEnd, uint256 startPriceInWei)
func (_IBlindAuction *IBlindAuctionFilterer) FilterBlindAuctionCreated(opts *bind.FilterOpts) (*IBlindAuctionBlindAuctionCreatedIterator, error) {

	logs, sub, err := _IBlindAuction.contract.FilterLogs(opts, "BlindAuctionCreated")
	if err != nil {
		return nil, err
	}
	return &IBlindAuctionBlindAuctionCreatedIterator{contract: _IBlindAuction.contract, event: "BlindAuctionCreated", logs: logs, sub: sub}, nil
}

// WatchBlindAuctionCreated is a free log subscription operation binding the contract event 0xe64fd4ba0702f46ae298a3be33f16f223a5e06556977c42dee952f976abfed83.
//
// Solidity: event BlindAuctionCreated(address seller, address nftAddress, bytes32 blindAuctionId, uint256 assetId, uint256 biddingEnd, uint256 revealEnd, uint256 startPriceInWei)
func (_IBlindAuction *IBlindAuctionFilterer) WatchBlindAuctionCreated(opts *bind.WatchOpts, sink chan<- *IBlindAuctionBlindAuctionCreated) (event.Subscription, error) {

	logs, sub, err := _IBlindAuction.contract.WatchLogs(opts, "BlindAuctionCreated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IBlindAuctionBlindAuctionCreated)
				if err := _IBlindAuction.contract.UnpackLog(event, "BlindAuctionCreated", log); err != nil {
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

// ParseBlindAuctionCreated is a log parse operation binding the contract event 0xe64fd4ba0702f46ae298a3be33f16f223a5e06556977c42dee952f976abfed83.
//
// Solidity: event BlindAuctionCreated(address seller, address nftAddress, bytes32 blindAuctionId, uint256 assetId, uint256 biddingEnd, uint256 revealEnd, uint256 startPriceInWei)
func (_IBlindAuction *IBlindAuctionFilterer) ParseBlindAuctionCreated(log types.Log) (*IBlindAuctionBlindAuctionCreated, error) {
	event := new(IBlindAuctionBlindAuctionCreated)
	if err := _IBlindAuction.contract.UnpackLog(event, "BlindAuctionCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IBlindAuctionBlindAuctionEndedIterator is returned from FilterBlindAuctionEnded and is used to iterate over the raw logs and unpacked data for BlindAuctionEnded events raised by the IBlindAuction contract.
type IBlindAuctionBlindAuctionEndedIterator struct {
	Event *IBlindAuctionBlindAuctionEnded // Event containing the contract specifics and raw log

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
func (it *IBlindAuctionBlindAuctionEndedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IBlindAuctionBlindAuctionEnded)
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
		it.Event = new(IBlindAuctionBlindAuctionEnded)
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
func (it *IBlindAuctionBlindAuctionEndedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IBlindAuctionBlindAuctionEndedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IBlindAuctionBlindAuctionEnded represents a BlindAuctionEnded event raised by the IBlindAuction contract.
type IBlindAuctionBlindAuctionEnded struct {
	BlindAuctionId [32]byte
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterBlindAuctionEnded is a free log retrieval operation binding the contract event 0xbe3b74f797d49d96897f669a8b62fc0d77fa4cd6cf999a1b08bc31e7f4faf219.
//
// Solidity: event BlindAuctionEnded(bytes32 blindAuctionId)
func (_IBlindAuction *IBlindAuctionFilterer) FilterBlindAuctionEnded(opts *bind.FilterOpts) (*IBlindAuctionBlindAuctionEndedIterator, error) {

	logs, sub, err := _IBlindAuction.contract.FilterLogs(opts, "BlindAuctionEnded")
	if err != nil {
		return nil, err
	}
	return &IBlindAuctionBlindAuctionEndedIterator{contract: _IBlindAuction.contract, event: "BlindAuctionEnded", logs: logs, sub: sub}, nil
}

// WatchBlindAuctionEnded is a free log subscription operation binding the contract event 0xbe3b74f797d49d96897f669a8b62fc0d77fa4cd6cf999a1b08bc31e7f4faf219.
//
// Solidity: event BlindAuctionEnded(bytes32 blindAuctionId)
func (_IBlindAuction *IBlindAuctionFilterer) WatchBlindAuctionEnded(opts *bind.WatchOpts, sink chan<- *IBlindAuctionBlindAuctionEnded) (event.Subscription, error) {

	logs, sub, err := _IBlindAuction.contract.WatchLogs(opts, "BlindAuctionEnded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IBlindAuctionBlindAuctionEnded)
				if err := _IBlindAuction.contract.UnpackLog(event, "BlindAuctionEnded", log); err != nil {
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

// ParseBlindAuctionEnded is a log parse operation binding the contract event 0xbe3b74f797d49d96897f669a8b62fc0d77fa4cd6cf999a1b08bc31e7f4faf219.
//
// Solidity: event BlindAuctionEnded(bytes32 blindAuctionId)
func (_IBlindAuction *IBlindAuctionFilterer) ParseBlindAuctionEnded(log types.Log) (*IBlindAuctionBlindAuctionEnded, error) {
	event := new(IBlindAuctionBlindAuctionEnded)
	if err := _IBlindAuction.contract.UnpackLog(event, "BlindAuctionEnded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IBlindAuctionBlindAuctionRefundIterator is returned from FilterBlindAuctionRefund and is used to iterate over the raw logs and unpacked data for BlindAuctionRefund events raised by the IBlindAuction contract.
type IBlindAuctionBlindAuctionRefundIterator struct {
	Event *IBlindAuctionBlindAuctionRefund // Event containing the contract specifics and raw log

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
func (it *IBlindAuctionBlindAuctionRefundIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IBlindAuctionBlindAuctionRefund)
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
		it.Event = new(IBlindAuctionBlindAuctionRefund)
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
func (it *IBlindAuctionBlindAuctionRefundIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IBlindAuctionBlindAuctionRefundIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IBlindAuctionBlindAuctionRefund represents a BlindAuctionRefund event raised by the IBlindAuction contract.
type IBlindAuctionBlindAuctionRefund struct {
	Bidder         common.Address
	BlindAuctionId [32]byte
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterBlindAuctionRefund is a free log retrieval operation binding the contract event 0xadada2afe0e9306e1b93faf56e74b041719e68f722de4f550d5a57e39610b404.
//
// Solidity: event BlindAuctionRefund(address bidder, bytes32 blindAuctionId)
func (_IBlindAuction *IBlindAuctionFilterer) FilterBlindAuctionRefund(opts *bind.FilterOpts) (*IBlindAuctionBlindAuctionRefundIterator, error) {

	logs, sub, err := _IBlindAuction.contract.FilterLogs(opts, "BlindAuctionRefund")
	if err != nil {
		return nil, err
	}
	return &IBlindAuctionBlindAuctionRefundIterator{contract: _IBlindAuction.contract, event: "BlindAuctionRefund", logs: logs, sub: sub}, nil
}

// WatchBlindAuctionRefund is a free log subscription operation binding the contract event 0xadada2afe0e9306e1b93faf56e74b041719e68f722de4f550d5a57e39610b404.
//
// Solidity: event BlindAuctionRefund(address bidder, bytes32 blindAuctionId)
func (_IBlindAuction *IBlindAuctionFilterer) WatchBlindAuctionRefund(opts *bind.WatchOpts, sink chan<- *IBlindAuctionBlindAuctionRefund) (event.Subscription, error) {

	logs, sub, err := _IBlindAuction.contract.WatchLogs(opts, "BlindAuctionRefund")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IBlindAuctionBlindAuctionRefund)
				if err := _IBlindAuction.contract.UnpackLog(event, "BlindAuctionRefund", log); err != nil {
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

// ParseBlindAuctionRefund is a log parse operation binding the contract event 0xadada2afe0e9306e1b93faf56e74b041719e68f722de4f550d5a57e39610b404.
//
// Solidity: event BlindAuctionRefund(address bidder, bytes32 blindAuctionId)
func (_IBlindAuction *IBlindAuctionFilterer) ParseBlindAuctionRefund(log types.Log) (*IBlindAuctionBlindAuctionRefund, error) {
	event := new(IBlindAuctionBlindAuctionRefund)
	if err := _IBlindAuction.contract.UnpackLog(event, "BlindAuctionRefund", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IBlindAuctionBlindAuctionSuccessfulIterator is returned from FilterBlindAuctionSuccessful and is used to iterate over the raw logs and unpacked data for BlindAuctionSuccessful events raised by the IBlindAuction contract.
type IBlindAuctionBlindAuctionSuccessfulIterator struct {
	Event *IBlindAuctionBlindAuctionSuccessful // Event containing the contract specifics and raw log

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
func (it *IBlindAuctionBlindAuctionSuccessfulIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IBlindAuctionBlindAuctionSuccessful)
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
		it.Event = new(IBlindAuctionBlindAuctionSuccessful)
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
func (it *IBlindAuctionBlindAuctionSuccessfulIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IBlindAuctionBlindAuctionSuccessfulIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IBlindAuctionBlindAuctionSuccessful represents a BlindAuctionSuccessful event raised by the IBlindAuction contract.
type IBlindAuctionBlindAuctionSuccessful struct {
	Seller         common.Address
	Buyer          common.Address
	BlindAuctionId [32]byte
	TotalPrice     *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterBlindAuctionSuccessful is a free log retrieval operation binding the contract event 0x2258179acb0769137f4ff532afd72c91199ae597abded9936acb61a2e66fc142.
//
// Solidity: event BlindAuctionSuccessful(address seller, address buyer, bytes32 blindAuctionId, uint256 totalPrice)
func (_IBlindAuction *IBlindAuctionFilterer) FilterBlindAuctionSuccessful(opts *bind.FilterOpts) (*IBlindAuctionBlindAuctionSuccessfulIterator, error) {

	logs, sub, err := _IBlindAuction.contract.FilterLogs(opts, "BlindAuctionSuccessful")
	if err != nil {
		return nil, err
	}
	return &IBlindAuctionBlindAuctionSuccessfulIterator{contract: _IBlindAuction.contract, event: "BlindAuctionSuccessful", logs: logs, sub: sub}, nil
}

// WatchBlindAuctionSuccessful is a free log subscription operation binding the contract event 0x2258179acb0769137f4ff532afd72c91199ae597abded9936acb61a2e66fc142.
//
// Solidity: event BlindAuctionSuccessful(address seller, address buyer, bytes32 blindAuctionId, uint256 totalPrice)
func (_IBlindAuction *IBlindAuctionFilterer) WatchBlindAuctionSuccessful(opts *bind.WatchOpts, sink chan<- *IBlindAuctionBlindAuctionSuccessful) (event.Subscription, error) {

	logs, sub, err := _IBlindAuction.contract.WatchLogs(opts, "BlindAuctionSuccessful")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IBlindAuctionBlindAuctionSuccessful)
				if err := _IBlindAuction.contract.UnpackLog(event, "BlindAuctionSuccessful", log); err != nil {
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

// ParseBlindAuctionSuccessful is a log parse operation binding the contract event 0x2258179acb0769137f4ff532afd72c91199ae597abded9936acb61a2e66fc142.
//
// Solidity: event BlindAuctionSuccessful(address seller, address buyer, bytes32 blindAuctionId, uint256 totalPrice)
func (_IBlindAuction *IBlindAuctionFilterer) ParseBlindAuctionSuccessful(log types.Log) (*IBlindAuctionBlindAuctionSuccessful, error) {
	event := new(IBlindAuctionBlindAuctionSuccessful)
	if err := _IBlindAuction.contract.UnpackLog(event, "BlindAuctionSuccessful", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IBlindAuctionRevealSuccessfulIterator is returned from FilterRevealSuccessful and is used to iterate over the raw logs and unpacked data for RevealSuccessful events raised by the IBlindAuction contract.
type IBlindAuctionRevealSuccessfulIterator struct {
	Event *IBlindAuctionRevealSuccessful // Event containing the contract specifics and raw log

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
func (it *IBlindAuctionRevealSuccessfulIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IBlindAuctionRevealSuccessful)
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
		it.Event = new(IBlindAuctionRevealSuccessful)
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
func (it *IBlindAuctionRevealSuccessfulIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IBlindAuctionRevealSuccessfulIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IBlindAuctionRevealSuccessful represents a RevealSuccessful event raised by the IBlindAuction contract.
type IBlindAuctionRevealSuccessful struct {
	Seller         common.Address
	BlindAuctionId [32]byte
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterRevealSuccessful is a free log retrieval operation binding the contract event 0x61d4ab60470c97b371c5f91be47873283386bb3af4de728e8795ef5318e013d6.
//
// Solidity: event RevealSuccessful(address seller, bytes32 blindAuctionId)
func (_IBlindAuction *IBlindAuctionFilterer) FilterRevealSuccessful(opts *bind.FilterOpts) (*IBlindAuctionRevealSuccessfulIterator, error) {

	logs, sub, err := _IBlindAuction.contract.FilterLogs(opts, "RevealSuccessful")
	if err != nil {
		return nil, err
	}
	return &IBlindAuctionRevealSuccessfulIterator{contract: _IBlindAuction.contract, event: "RevealSuccessful", logs: logs, sub: sub}, nil
}

// WatchRevealSuccessful is a free log subscription operation binding the contract event 0x61d4ab60470c97b371c5f91be47873283386bb3af4de728e8795ef5318e013d6.
//
// Solidity: event RevealSuccessful(address seller, bytes32 blindAuctionId)
func (_IBlindAuction *IBlindAuctionFilterer) WatchRevealSuccessful(opts *bind.WatchOpts, sink chan<- *IBlindAuctionRevealSuccessful) (event.Subscription, error) {

	logs, sub, err := _IBlindAuction.contract.WatchLogs(opts, "RevealSuccessful")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IBlindAuctionRevealSuccessful)
				if err := _IBlindAuction.contract.UnpackLog(event, "RevealSuccessful", log); err != nil {
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

// ParseRevealSuccessful is a log parse operation binding the contract event 0x61d4ab60470c97b371c5f91be47873283386bb3af4de728e8795ef5318e013d6.
//
// Solidity: event RevealSuccessful(address seller, bytes32 blindAuctionId)
func (_IBlindAuction *IBlindAuctionFilterer) ParseRevealSuccessful(log types.Log) (*IBlindAuctionRevealSuccessful, error) {
	event := new(IBlindAuctionRevealSuccessful)
	if err := _IBlindAuction.contract.UnpackLog(event, "RevealSuccessful", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
