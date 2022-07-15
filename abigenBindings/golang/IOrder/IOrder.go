// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package IOrder

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

// IOrderMetaData contains all meta data concerning the IOrder contract.
var IOrderMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"InvalidExpiredTime\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OrderExpired\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"who\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"}],\"name\":\"OrderCancelled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"orderId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"assetId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"nftAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"priceInWei\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"expiredAt\",\"type\":\"uint256\"}],\"name\":\"OrderCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"buyer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"}],\"name\":\"OrderSuccessful\",\"type\":\"event\"}]",
}

// IOrderABI is the input ABI used to generate the binding from.
// Deprecated: Use IOrderMetaData.ABI instead.
var IOrderABI = IOrderMetaData.ABI

// IOrder is an auto generated Go binding around an Ethereum contract.
type IOrder struct {
	IOrderCaller     // Read-only binding to the contract
	IOrderTransactor // Write-only binding to the contract
	IOrderFilterer   // Log filterer for contract events
}

// IOrderCaller is an auto generated read-only Go binding around an Ethereum contract.
type IOrderCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IOrderTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IOrderTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IOrderFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IOrderFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IOrderSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IOrderSession struct {
	Contract     *IOrder           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IOrderCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IOrderCallerSession struct {
	Contract *IOrderCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// IOrderTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IOrderTransactorSession struct {
	Contract     *IOrderTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IOrderRaw is an auto generated low-level Go binding around an Ethereum contract.
type IOrderRaw struct {
	Contract *IOrder // Generic contract binding to access the raw methods on
}

// IOrderCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IOrderCallerRaw struct {
	Contract *IOrderCaller // Generic read-only contract binding to access the raw methods on
}

// IOrderTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IOrderTransactorRaw struct {
	Contract *IOrderTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIOrder creates a new instance of IOrder, bound to a specific deployed contract.
func NewIOrder(address common.Address, backend bind.ContractBackend) (*IOrder, error) {
	contract, err := bindIOrder(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IOrder{IOrderCaller: IOrderCaller{contract: contract}, IOrderTransactor: IOrderTransactor{contract: contract}, IOrderFilterer: IOrderFilterer{contract: contract}}, nil
}

// NewIOrderCaller creates a new read-only instance of IOrder, bound to a specific deployed contract.
func NewIOrderCaller(address common.Address, caller bind.ContractCaller) (*IOrderCaller, error) {
	contract, err := bindIOrder(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IOrderCaller{contract: contract}, nil
}

// NewIOrderTransactor creates a new write-only instance of IOrder, bound to a specific deployed contract.
func NewIOrderTransactor(address common.Address, transactor bind.ContractTransactor) (*IOrderTransactor, error) {
	contract, err := bindIOrder(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IOrderTransactor{contract: contract}, nil
}

// NewIOrderFilterer creates a new log filterer instance of IOrder, bound to a specific deployed contract.
func NewIOrderFilterer(address common.Address, filterer bind.ContractFilterer) (*IOrderFilterer, error) {
	contract, err := bindIOrder(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IOrderFilterer{contract: contract}, nil
}

// bindIOrder binds a generic wrapper to an already deployed contract.
func bindIOrder(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IOrderABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IOrder *IOrderRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IOrder.Contract.IOrderCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IOrder *IOrderRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IOrder.Contract.IOrderTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IOrder *IOrderRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IOrder.Contract.IOrderTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IOrder *IOrderCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IOrder.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IOrder *IOrderTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IOrder.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IOrder *IOrderTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IOrder.Contract.contract.Transact(opts, method, params...)
}

// IOrderOrderCancelledIterator is returned from FilterOrderCancelled and is used to iterate over the raw logs and unpacked data for OrderCancelled events raised by the IOrder contract.
type IOrderOrderCancelledIterator struct {
	Event *IOrderOrderCancelled // Event containing the contract specifics and raw log

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
func (it *IOrderOrderCancelledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IOrderOrderCancelled)
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
		it.Event = new(IOrderOrderCancelled)
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
func (it *IOrderOrderCancelledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IOrderOrderCancelledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IOrderOrderCancelled represents a OrderCancelled event raised by the IOrder contract.
type IOrderOrderCancelled struct {
	Who common.Address
	Id  [32]byte
	Raw types.Log // Blockchain specific contextual infos
}

// FilterOrderCancelled is a free log retrieval operation binding the contract event 0x35974c4230d53fb4c6e8553fd900c88ba92747dbc689a79bcd6ba755cb936985.
//
// Solidity: event OrderCancelled(address who, bytes32 id)
func (_IOrder *IOrderFilterer) FilterOrderCancelled(opts *bind.FilterOpts) (*IOrderOrderCancelledIterator, error) {

	logs, sub, err := _IOrder.contract.FilterLogs(opts, "OrderCancelled")
	if err != nil {
		return nil, err
	}
	return &IOrderOrderCancelledIterator{contract: _IOrder.contract, event: "OrderCancelled", logs: logs, sub: sub}, nil
}

// WatchOrderCancelled is a free log subscription operation binding the contract event 0x35974c4230d53fb4c6e8553fd900c88ba92747dbc689a79bcd6ba755cb936985.
//
// Solidity: event OrderCancelled(address who, bytes32 id)
func (_IOrder *IOrderFilterer) WatchOrderCancelled(opts *bind.WatchOpts, sink chan<- *IOrderOrderCancelled) (event.Subscription, error) {

	logs, sub, err := _IOrder.contract.WatchLogs(opts, "OrderCancelled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IOrderOrderCancelled)
				if err := _IOrder.contract.UnpackLog(event, "OrderCancelled", log); err != nil {
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

// ParseOrderCancelled is a log parse operation binding the contract event 0x35974c4230d53fb4c6e8553fd900c88ba92747dbc689a79bcd6ba755cb936985.
//
// Solidity: event OrderCancelled(address who, bytes32 id)
func (_IOrder *IOrderFilterer) ParseOrderCancelled(log types.Log) (*IOrderOrderCancelled, error) {
	event := new(IOrderOrderCancelled)
	if err := _IOrder.contract.UnpackLog(event, "OrderCancelled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IOrderOrderCreatedIterator is returned from FilterOrderCreated and is used to iterate over the raw logs and unpacked data for OrderCreated events raised by the IOrder contract.
type IOrderOrderCreatedIterator struct {
	Event *IOrderOrderCreated // Event containing the contract specifics and raw log

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
func (it *IOrderOrderCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IOrderOrderCreated)
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
		it.Event = new(IOrderOrderCreated)
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
func (it *IOrderOrderCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IOrderOrderCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IOrderOrderCreated represents a OrderCreated event raised by the IOrder contract.
type IOrderOrderCreated struct {
	OrderId    [32]byte
	AssetId    *big.Int
	Seller     common.Address
	NftAddress common.Address
	PriceInWei *big.Int
	ExpiredAt  *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterOrderCreated is a free log retrieval operation binding the contract event 0x84c66c3f7ba4b390e20e8e8233e2a516f3ce34a72749e4f12bd010dfba238039.
//
// Solidity: event OrderCreated(bytes32 orderId, uint256 assetId, address seller, address nftAddress, uint256 priceInWei, uint256 expiredAt)
func (_IOrder *IOrderFilterer) FilterOrderCreated(opts *bind.FilterOpts) (*IOrderOrderCreatedIterator, error) {

	logs, sub, err := _IOrder.contract.FilterLogs(opts, "OrderCreated")
	if err != nil {
		return nil, err
	}
	return &IOrderOrderCreatedIterator{contract: _IOrder.contract, event: "OrderCreated", logs: logs, sub: sub}, nil
}

// WatchOrderCreated is a free log subscription operation binding the contract event 0x84c66c3f7ba4b390e20e8e8233e2a516f3ce34a72749e4f12bd010dfba238039.
//
// Solidity: event OrderCreated(bytes32 orderId, uint256 assetId, address seller, address nftAddress, uint256 priceInWei, uint256 expiredAt)
func (_IOrder *IOrderFilterer) WatchOrderCreated(opts *bind.WatchOpts, sink chan<- *IOrderOrderCreated) (event.Subscription, error) {

	logs, sub, err := _IOrder.contract.WatchLogs(opts, "OrderCreated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IOrderOrderCreated)
				if err := _IOrder.contract.UnpackLog(event, "OrderCreated", log); err != nil {
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

// ParseOrderCreated is a log parse operation binding the contract event 0x84c66c3f7ba4b390e20e8e8233e2a516f3ce34a72749e4f12bd010dfba238039.
//
// Solidity: event OrderCreated(bytes32 orderId, uint256 assetId, address seller, address nftAddress, uint256 priceInWei, uint256 expiredAt)
func (_IOrder *IOrderFilterer) ParseOrderCreated(log types.Log) (*IOrderOrderCreated, error) {
	event := new(IOrderOrderCreated)
	if err := _IOrder.contract.UnpackLog(event, "OrderCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IOrderOrderSuccessfulIterator is returned from FilterOrderSuccessful and is used to iterate over the raw logs and unpacked data for OrderSuccessful events raised by the IOrder contract.
type IOrderOrderSuccessfulIterator struct {
	Event *IOrderOrderSuccessful // Event containing the contract specifics and raw log

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
func (it *IOrderOrderSuccessfulIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IOrderOrderSuccessful)
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
		it.Event = new(IOrderOrderSuccessful)
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
func (it *IOrderOrderSuccessfulIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IOrderOrderSuccessfulIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IOrderOrderSuccessful represents a OrderSuccessful event raised by the IOrder contract.
type IOrderOrderSuccessful struct {
	Id     [32]byte
	Buyer  common.Address
	Seller common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterOrderSuccessful is a free log retrieval operation binding the contract event 0xa3fe2a6ffb1ad0de1c1ee8e3513d205b99980c40af9de82e502095245debabbb.
//
// Solidity: event OrderSuccessful(bytes32 id, address buyer, address seller)
func (_IOrder *IOrderFilterer) FilterOrderSuccessful(opts *bind.FilterOpts) (*IOrderOrderSuccessfulIterator, error) {

	logs, sub, err := _IOrder.contract.FilterLogs(opts, "OrderSuccessful")
	if err != nil {
		return nil, err
	}
	return &IOrderOrderSuccessfulIterator{contract: _IOrder.contract, event: "OrderSuccessful", logs: logs, sub: sub}, nil
}

// WatchOrderSuccessful is a free log subscription operation binding the contract event 0xa3fe2a6ffb1ad0de1c1ee8e3513d205b99980c40af9de82e502095245debabbb.
//
// Solidity: event OrderSuccessful(bytes32 id, address buyer, address seller)
func (_IOrder *IOrderFilterer) WatchOrderSuccessful(opts *bind.WatchOpts, sink chan<- *IOrderOrderSuccessful) (event.Subscription, error) {

	logs, sub, err := _IOrder.contract.WatchLogs(opts, "OrderSuccessful")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IOrderOrderSuccessful)
				if err := _IOrder.contract.UnpackLog(event, "OrderSuccessful", log); err != nil {
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

// ParseOrderSuccessful is a log parse operation binding the contract event 0xa3fe2a6ffb1ad0de1c1ee8e3513d205b99980c40af9de82e502095245debabbb.
//
// Solidity: event OrderSuccessful(bytes32 id, address buyer, address seller)
func (_IOrder *IOrderFilterer) ParseOrderSuccessful(log types.Log) (*IOrderOrderSuccessful, error) {
	event := new(IOrderOrderSuccessful)
	if err := _IOrder.contract.UnpackLog(event, "OrderSuccessful", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
