// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ITrading

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

// ITradingMetaData contains all meta data concerning the ITrading contract.
var ITradingMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"InvalidExpiredTime\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TradingExpired\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"who\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"}],\"name\":\"TradingCancelled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"tradingId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"assetId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"nftAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"priceInWei\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"expiredAt\",\"type\":\"uint256\"}],\"name\":\"TradingCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"buyer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"}],\"name\":\"TradingSuccessful\",\"type\":\"event\"}]",
}

// ITradingABI is the input ABI used to generate the binding from.
// Deprecated: Use ITradingMetaData.ABI instead.
var ITradingABI = ITradingMetaData.ABI

// ITrading is an auto generated Go binding around an Ethereum contract.
type ITrading struct {
	ITradingCaller     // Read-only binding to the contract
	ITradingTransactor // Write-only binding to the contract
	ITradingFilterer   // Log filterer for contract events
}

// ITradingCaller is an auto generated read-only Go binding around an Ethereum contract.
type ITradingCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ITradingTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ITradingTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ITradingFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ITradingFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ITradingSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ITradingSession struct {
	Contract     *ITrading         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ITradingCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ITradingCallerSession struct {
	Contract *ITradingCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// ITradingTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ITradingTransactorSession struct {
	Contract     *ITradingTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// ITradingRaw is an auto generated low-level Go binding around an Ethereum contract.
type ITradingRaw struct {
	Contract *ITrading // Generic contract binding to access the raw methods on
}

// ITradingCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ITradingCallerRaw struct {
	Contract *ITradingCaller // Generic read-only contract binding to access the raw methods on
}

// ITradingTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ITradingTransactorRaw struct {
	Contract *ITradingTransactor // Generic write-only contract binding to access the raw methods on
}

// NewITrading creates a new instance of ITrading, bound to a specific deployed contract.
func NewITrading(address common.Address, backend bind.ContractBackend) (*ITrading, error) {
	contract, err := bindITrading(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ITrading{ITradingCaller: ITradingCaller{contract: contract}, ITradingTransactor: ITradingTransactor{contract: contract}, ITradingFilterer: ITradingFilterer{contract: contract}}, nil
}

// NewITradingCaller creates a new read-only instance of ITrading, bound to a specific deployed contract.
func NewITradingCaller(address common.Address, caller bind.ContractCaller) (*ITradingCaller, error) {
	contract, err := bindITrading(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ITradingCaller{contract: contract}, nil
}

// NewITradingTransactor creates a new write-only instance of ITrading, bound to a specific deployed contract.
func NewITradingTransactor(address common.Address, transactor bind.ContractTransactor) (*ITradingTransactor, error) {
	contract, err := bindITrading(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ITradingTransactor{contract: contract}, nil
}

// NewITradingFilterer creates a new log filterer instance of ITrading, bound to a specific deployed contract.
func NewITradingFilterer(address common.Address, filterer bind.ContractFilterer) (*ITradingFilterer, error) {
	contract, err := bindITrading(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ITradingFilterer{contract: contract}, nil
}

// bindITrading binds a generic wrapper to an already deployed contract.
func bindITrading(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ITradingABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ITrading *ITradingRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ITrading.Contract.ITradingCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ITrading *ITradingRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ITrading.Contract.ITradingTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ITrading *ITradingRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ITrading.Contract.ITradingTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ITrading *ITradingCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ITrading.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ITrading *ITradingTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ITrading.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ITrading *ITradingTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ITrading.Contract.contract.Transact(opts, method, params...)
}

// ITradingTradingCancelledIterator is returned from FilterTradingCancelled and is used to iterate over the raw logs and unpacked data for TradingCancelled events raised by the ITrading contract.
type ITradingTradingCancelledIterator struct {
	Event *ITradingTradingCancelled // Event containing the contract specifics and raw log

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
func (it *ITradingTradingCancelledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ITradingTradingCancelled)
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
		it.Event = new(ITradingTradingCancelled)
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
func (it *ITradingTradingCancelledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ITradingTradingCancelledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ITradingTradingCancelled represents a TradingCancelled event raised by the ITrading contract.
type ITradingTradingCancelled struct {
	Who common.Address
	Id  [32]byte
	Raw types.Log // Blockchain specific contextual infos
}

// FilterTradingCancelled is a free log retrieval operation binding the contract event 0x559020ef5970f017b414f50b16b6cacd96c9bbb8fd863b9d71ab3d3222d781fa.
//
// Solidity: event TradingCancelled(address who, bytes32 id)
func (_ITrading *ITradingFilterer) FilterTradingCancelled(opts *bind.FilterOpts) (*ITradingTradingCancelledIterator, error) {

	logs, sub, err := _ITrading.contract.FilterLogs(opts, "TradingCancelled")
	if err != nil {
		return nil, err
	}
	return &ITradingTradingCancelledIterator{contract: _ITrading.contract, event: "TradingCancelled", logs: logs, sub: sub}, nil
}

// WatchTradingCancelled is a free log subscription operation binding the contract event 0x559020ef5970f017b414f50b16b6cacd96c9bbb8fd863b9d71ab3d3222d781fa.
//
// Solidity: event TradingCancelled(address who, bytes32 id)
func (_ITrading *ITradingFilterer) WatchTradingCancelled(opts *bind.WatchOpts, sink chan<- *ITradingTradingCancelled) (event.Subscription, error) {

	logs, sub, err := _ITrading.contract.WatchLogs(opts, "TradingCancelled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ITradingTradingCancelled)
				if err := _ITrading.contract.UnpackLog(event, "TradingCancelled", log); err != nil {
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

// ParseTradingCancelled is a log parse operation binding the contract event 0x559020ef5970f017b414f50b16b6cacd96c9bbb8fd863b9d71ab3d3222d781fa.
//
// Solidity: event TradingCancelled(address who, bytes32 id)
func (_ITrading *ITradingFilterer) ParseTradingCancelled(log types.Log) (*ITradingTradingCancelled, error) {
	event := new(ITradingTradingCancelled)
	if err := _ITrading.contract.UnpackLog(event, "TradingCancelled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ITradingTradingCreatedIterator is returned from FilterTradingCreated and is used to iterate over the raw logs and unpacked data for TradingCreated events raised by the ITrading contract.
type ITradingTradingCreatedIterator struct {
	Event *ITradingTradingCreated // Event containing the contract specifics and raw log

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
func (it *ITradingTradingCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ITradingTradingCreated)
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
		it.Event = new(ITradingTradingCreated)
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
func (it *ITradingTradingCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ITradingTradingCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ITradingTradingCreated represents a TradingCreated event raised by the ITrading contract.
type ITradingTradingCreated struct {
	TradingId  [32]byte
	AssetId    *big.Int
	Seller     common.Address
	NftAddress common.Address
	PriceInWei *big.Int
	ExpiredAt  *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterTradingCreated is a free log retrieval operation binding the contract event 0x80d7cd14a57c43d79381c3567166d2e66efc27aa3d0dca4b41e7a3fd736569c6.
//
// Solidity: event TradingCreated(bytes32 tradingId, uint256 assetId, address seller, address nftAddress, uint256 priceInWei, uint256 expiredAt)
func (_ITrading *ITradingFilterer) FilterTradingCreated(opts *bind.FilterOpts) (*ITradingTradingCreatedIterator, error) {

	logs, sub, err := _ITrading.contract.FilterLogs(opts, "TradingCreated")
	if err != nil {
		return nil, err
	}
	return &ITradingTradingCreatedIterator{contract: _ITrading.contract, event: "TradingCreated", logs: logs, sub: sub}, nil
}

// WatchTradingCreated is a free log subscription operation binding the contract event 0x80d7cd14a57c43d79381c3567166d2e66efc27aa3d0dca4b41e7a3fd736569c6.
//
// Solidity: event TradingCreated(bytes32 tradingId, uint256 assetId, address seller, address nftAddress, uint256 priceInWei, uint256 expiredAt)
func (_ITrading *ITradingFilterer) WatchTradingCreated(opts *bind.WatchOpts, sink chan<- *ITradingTradingCreated) (event.Subscription, error) {

	logs, sub, err := _ITrading.contract.WatchLogs(opts, "TradingCreated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ITradingTradingCreated)
				if err := _ITrading.contract.UnpackLog(event, "TradingCreated", log); err != nil {
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

// ParseTradingCreated is a log parse operation binding the contract event 0x80d7cd14a57c43d79381c3567166d2e66efc27aa3d0dca4b41e7a3fd736569c6.
//
// Solidity: event TradingCreated(bytes32 tradingId, uint256 assetId, address seller, address nftAddress, uint256 priceInWei, uint256 expiredAt)
func (_ITrading *ITradingFilterer) ParseTradingCreated(log types.Log) (*ITradingTradingCreated, error) {
	event := new(ITradingTradingCreated)
	if err := _ITrading.contract.UnpackLog(event, "TradingCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ITradingTradingSuccessfulIterator is returned from FilterTradingSuccessful and is used to iterate over the raw logs and unpacked data for TradingSuccessful events raised by the ITrading contract.
type ITradingTradingSuccessfulIterator struct {
	Event *ITradingTradingSuccessful // Event containing the contract specifics and raw log

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
func (it *ITradingTradingSuccessfulIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ITradingTradingSuccessful)
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
		it.Event = new(ITradingTradingSuccessful)
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
func (it *ITradingTradingSuccessfulIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ITradingTradingSuccessfulIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ITradingTradingSuccessful represents a TradingSuccessful event raised by the ITrading contract.
type ITradingTradingSuccessful struct {
	Id     [32]byte
	Buyer  common.Address
	Seller common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterTradingSuccessful is a free log retrieval operation binding the contract event 0x5cb1fd7912ab2064b667d5d47f54a6a7e64fab88792e1859853ac3c6a3dd9f63.
//
// Solidity: event TradingSuccessful(bytes32 id, address buyer, address seller)
func (_ITrading *ITradingFilterer) FilterTradingSuccessful(opts *bind.FilterOpts) (*ITradingTradingSuccessfulIterator, error) {

	logs, sub, err := _ITrading.contract.FilterLogs(opts, "TradingSuccessful")
	if err != nil {
		return nil, err
	}
	return &ITradingTradingSuccessfulIterator{contract: _ITrading.contract, event: "TradingSuccessful", logs: logs, sub: sub}, nil
}

// WatchTradingSuccessful is a free log subscription operation binding the contract event 0x5cb1fd7912ab2064b667d5d47f54a6a7e64fab88792e1859853ac3c6a3dd9f63.
//
// Solidity: event TradingSuccessful(bytes32 id, address buyer, address seller)
func (_ITrading *ITradingFilterer) WatchTradingSuccessful(opts *bind.WatchOpts, sink chan<- *ITradingTradingSuccessful) (event.Subscription, error) {

	logs, sub, err := _ITrading.contract.WatchLogs(opts, "TradingSuccessful")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ITradingTradingSuccessful)
				if err := _ITrading.contract.UnpackLog(event, "TradingSuccessful", log); err != nil {
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

// ParseTradingSuccessful is a log parse operation binding the contract event 0x5cb1fd7912ab2064b667d5d47f54a6a7e64fab88792e1859853ac3c6a3dd9f63.
//
// Solidity: event TradingSuccessful(bytes32 id, address buyer, address seller)
func (_ITrading *ITradingFilterer) ParseTradingSuccessful(log types.Log) (*ITradingTradingSuccessful, error) {
	event := new(ITradingTradingSuccessful)
	if err := _ITrading.contract.UnpackLog(event, "TradingSuccessful", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
