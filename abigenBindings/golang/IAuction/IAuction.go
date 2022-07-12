// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package IAuction

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

// IAuctionMetaData contains all meta data concerning the IAuction contract.
var IAuctionMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"AuctionEnded\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidBiddingEnd\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidBiddingPrice\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidWithdraw\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotRunning\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotWinner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RewardGranted\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"TooEarly\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"TooLate\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"canceller\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"auctionId\",\"type\":\"bytes32\"}],\"name\":\"AuctionCancelledSuccessful\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"bidder\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"auctionId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"AuctionRefundSuccessful\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"auctionHighestBidder\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"auctionId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"assetId\",\"type\":\"uint256\"}],\"name\":\"GrantAuctionRewardSuccessful\",\"type\":\"event\"}]",
}

// IAuctionABI is the input ABI used to generate the binding from.
// Deprecated: Use IAuctionMetaData.ABI instead.
var IAuctionABI = IAuctionMetaData.ABI

// IAuction is an auto generated Go binding around an Ethereum contract.
type IAuction struct {
	IAuctionCaller     // Read-only binding to the contract
	IAuctionTransactor // Write-only binding to the contract
	IAuctionFilterer   // Log filterer for contract events
}

// IAuctionCaller is an auto generated read-only Go binding around an Ethereum contract.
type IAuctionCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IAuctionTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IAuctionTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IAuctionFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IAuctionFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IAuctionSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IAuctionSession struct {
	Contract     *IAuction         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IAuctionCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IAuctionCallerSession struct {
	Contract *IAuctionCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// IAuctionTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IAuctionTransactorSession struct {
	Contract     *IAuctionTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// IAuctionRaw is an auto generated low-level Go binding around an Ethereum contract.
type IAuctionRaw struct {
	Contract *IAuction // Generic contract binding to access the raw methods on
}

// IAuctionCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IAuctionCallerRaw struct {
	Contract *IAuctionCaller // Generic read-only contract binding to access the raw methods on
}

// IAuctionTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IAuctionTransactorRaw struct {
	Contract *IAuctionTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIAuction creates a new instance of IAuction, bound to a specific deployed contract.
func NewIAuction(address common.Address, backend bind.ContractBackend) (*IAuction, error) {
	contract, err := bindIAuction(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IAuction{IAuctionCaller: IAuctionCaller{contract: contract}, IAuctionTransactor: IAuctionTransactor{contract: contract}, IAuctionFilterer: IAuctionFilterer{contract: contract}}, nil
}

// NewIAuctionCaller creates a new read-only instance of IAuction, bound to a specific deployed contract.
func NewIAuctionCaller(address common.Address, caller bind.ContractCaller) (*IAuctionCaller, error) {
	contract, err := bindIAuction(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IAuctionCaller{contract: contract}, nil
}

// NewIAuctionTransactor creates a new write-only instance of IAuction, bound to a specific deployed contract.
func NewIAuctionTransactor(address common.Address, transactor bind.ContractTransactor) (*IAuctionTransactor, error) {
	contract, err := bindIAuction(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IAuctionTransactor{contract: contract}, nil
}

// NewIAuctionFilterer creates a new log filterer instance of IAuction, bound to a specific deployed contract.
func NewIAuctionFilterer(address common.Address, filterer bind.ContractFilterer) (*IAuctionFilterer, error) {
	contract, err := bindIAuction(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IAuctionFilterer{contract: contract}, nil
}

// bindIAuction binds a generic wrapper to an already deployed contract.
func bindIAuction(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IAuctionABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IAuction *IAuctionRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IAuction.Contract.IAuctionCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IAuction *IAuctionRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IAuction.Contract.IAuctionTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IAuction *IAuctionRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IAuction.Contract.IAuctionTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IAuction *IAuctionCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IAuction.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IAuction *IAuctionTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IAuction.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IAuction *IAuctionTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IAuction.Contract.contract.Transact(opts, method, params...)
}

// IAuctionAuctionCancelledSuccessfulIterator is returned from FilterAuctionCancelledSuccessful and is used to iterate over the raw logs and unpacked data for AuctionCancelledSuccessful events raised by the IAuction contract.
type IAuctionAuctionCancelledSuccessfulIterator struct {
	Event *IAuctionAuctionCancelledSuccessful // Event containing the contract specifics and raw log

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
func (it *IAuctionAuctionCancelledSuccessfulIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IAuctionAuctionCancelledSuccessful)
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
		it.Event = new(IAuctionAuctionCancelledSuccessful)
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
func (it *IAuctionAuctionCancelledSuccessfulIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IAuctionAuctionCancelledSuccessfulIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IAuctionAuctionCancelledSuccessful represents a AuctionCancelledSuccessful event raised by the IAuction contract.
type IAuctionAuctionCancelledSuccessful struct {
	Canceller common.Address
	AuctionId [32]byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterAuctionCancelledSuccessful is a free log retrieval operation binding the contract event 0xa021239bb373fef19aa7c7ef798961ab038ebf95bd85d546b731d25855db190e.
//
// Solidity: event AuctionCancelledSuccessful(address canceller, bytes32 auctionId)
func (_IAuction *IAuctionFilterer) FilterAuctionCancelledSuccessful(opts *bind.FilterOpts) (*IAuctionAuctionCancelledSuccessfulIterator, error) {

	logs, sub, err := _IAuction.contract.FilterLogs(opts, "AuctionCancelledSuccessful")
	if err != nil {
		return nil, err
	}
	return &IAuctionAuctionCancelledSuccessfulIterator{contract: _IAuction.contract, event: "AuctionCancelledSuccessful", logs: logs, sub: sub}, nil
}

// WatchAuctionCancelledSuccessful is a free log subscription operation binding the contract event 0xa021239bb373fef19aa7c7ef798961ab038ebf95bd85d546b731d25855db190e.
//
// Solidity: event AuctionCancelledSuccessful(address canceller, bytes32 auctionId)
func (_IAuction *IAuctionFilterer) WatchAuctionCancelledSuccessful(opts *bind.WatchOpts, sink chan<- *IAuctionAuctionCancelledSuccessful) (event.Subscription, error) {

	logs, sub, err := _IAuction.contract.WatchLogs(opts, "AuctionCancelledSuccessful")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IAuctionAuctionCancelledSuccessful)
				if err := _IAuction.contract.UnpackLog(event, "AuctionCancelledSuccessful", log); err != nil {
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

// ParseAuctionCancelledSuccessful is a log parse operation binding the contract event 0xa021239bb373fef19aa7c7ef798961ab038ebf95bd85d546b731d25855db190e.
//
// Solidity: event AuctionCancelledSuccessful(address canceller, bytes32 auctionId)
func (_IAuction *IAuctionFilterer) ParseAuctionCancelledSuccessful(log types.Log) (*IAuctionAuctionCancelledSuccessful, error) {
	event := new(IAuctionAuctionCancelledSuccessful)
	if err := _IAuction.contract.UnpackLog(event, "AuctionCancelledSuccessful", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IAuctionAuctionRefundSuccessfulIterator is returned from FilterAuctionRefundSuccessful and is used to iterate over the raw logs and unpacked data for AuctionRefundSuccessful events raised by the IAuction contract.
type IAuctionAuctionRefundSuccessfulIterator struct {
	Event *IAuctionAuctionRefundSuccessful // Event containing the contract specifics and raw log

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
func (it *IAuctionAuctionRefundSuccessfulIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IAuctionAuctionRefundSuccessful)
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
		it.Event = new(IAuctionAuctionRefundSuccessful)
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
func (it *IAuctionAuctionRefundSuccessfulIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IAuctionAuctionRefundSuccessfulIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IAuctionAuctionRefundSuccessful represents a AuctionRefundSuccessful event raised by the IAuction contract.
type IAuctionAuctionRefundSuccessful struct {
	Bidder    common.Address
	AuctionId [32]byte
	Value     *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterAuctionRefundSuccessful is a free log retrieval operation binding the contract event 0x5f4b88832be3746851e4d0ce6129f89fe39ddcf5a4b0204699724c459baf7d9d.
//
// Solidity: event AuctionRefundSuccessful(address bidder, bytes32 auctionId, uint256 value)
func (_IAuction *IAuctionFilterer) FilterAuctionRefundSuccessful(opts *bind.FilterOpts) (*IAuctionAuctionRefundSuccessfulIterator, error) {

	logs, sub, err := _IAuction.contract.FilterLogs(opts, "AuctionRefundSuccessful")
	if err != nil {
		return nil, err
	}
	return &IAuctionAuctionRefundSuccessfulIterator{contract: _IAuction.contract, event: "AuctionRefundSuccessful", logs: logs, sub: sub}, nil
}

// WatchAuctionRefundSuccessful is a free log subscription operation binding the contract event 0x5f4b88832be3746851e4d0ce6129f89fe39ddcf5a4b0204699724c459baf7d9d.
//
// Solidity: event AuctionRefundSuccessful(address bidder, bytes32 auctionId, uint256 value)
func (_IAuction *IAuctionFilterer) WatchAuctionRefundSuccessful(opts *bind.WatchOpts, sink chan<- *IAuctionAuctionRefundSuccessful) (event.Subscription, error) {

	logs, sub, err := _IAuction.contract.WatchLogs(opts, "AuctionRefundSuccessful")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IAuctionAuctionRefundSuccessful)
				if err := _IAuction.contract.UnpackLog(event, "AuctionRefundSuccessful", log); err != nil {
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

// ParseAuctionRefundSuccessful is a log parse operation binding the contract event 0x5f4b88832be3746851e4d0ce6129f89fe39ddcf5a4b0204699724c459baf7d9d.
//
// Solidity: event AuctionRefundSuccessful(address bidder, bytes32 auctionId, uint256 value)
func (_IAuction *IAuctionFilterer) ParseAuctionRefundSuccessful(log types.Log) (*IAuctionAuctionRefundSuccessful, error) {
	event := new(IAuctionAuctionRefundSuccessful)
	if err := _IAuction.contract.UnpackLog(event, "AuctionRefundSuccessful", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IAuctionGrantAuctionRewardSuccessfulIterator is returned from FilterGrantAuctionRewardSuccessful and is used to iterate over the raw logs and unpacked data for GrantAuctionRewardSuccessful events raised by the IAuction contract.
type IAuctionGrantAuctionRewardSuccessfulIterator struct {
	Event *IAuctionGrantAuctionRewardSuccessful // Event containing the contract specifics and raw log

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
func (it *IAuctionGrantAuctionRewardSuccessfulIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IAuctionGrantAuctionRewardSuccessful)
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
		it.Event = new(IAuctionGrantAuctionRewardSuccessful)
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
func (it *IAuctionGrantAuctionRewardSuccessfulIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IAuctionGrantAuctionRewardSuccessfulIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IAuctionGrantAuctionRewardSuccessful represents a GrantAuctionRewardSuccessful event raised by the IAuction contract.
type IAuctionGrantAuctionRewardSuccessful struct {
	AuctionHighestBidder common.Address
	AuctionId            [32]byte
	AssetId              *big.Int
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterGrantAuctionRewardSuccessful is a free log retrieval operation binding the contract event 0xa307d88e125ed19583b02332c5bd74f0323774248b3f9cac5047c2433d9d8ddc.
//
// Solidity: event GrantAuctionRewardSuccessful(address auctionHighestBidder, bytes32 auctionId, uint256 assetId)
func (_IAuction *IAuctionFilterer) FilterGrantAuctionRewardSuccessful(opts *bind.FilterOpts) (*IAuctionGrantAuctionRewardSuccessfulIterator, error) {

	logs, sub, err := _IAuction.contract.FilterLogs(opts, "GrantAuctionRewardSuccessful")
	if err != nil {
		return nil, err
	}
	return &IAuctionGrantAuctionRewardSuccessfulIterator{contract: _IAuction.contract, event: "GrantAuctionRewardSuccessful", logs: logs, sub: sub}, nil
}

// WatchGrantAuctionRewardSuccessful is a free log subscription operation binding the contract event 0xa307d88e125ed19583b02332c5bd74f0323774248b3f9cac5047c2433d9d8ddc.
//
// Solidity: event GrantAuctionRewardSuccessful(address auctionHighestBidder, bytes32 auctionId, uint256 assetId)
func (_IAuction *IAuctionFilterer) WatchGrantAuctionRewardSuccessful(opts *bind.WatchOpts, sink chan<- *IAuctionGrantAuctionRewardSuccessful) (event.Subscription, error) {

	logs, sub, err := _IAuction.contract.WatchLogs(opts, "GrantAuctionRewardSuccessful")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IAuctionGrantAuctionRewardSuccessful)
				if err := _IAuction.contract.UnpackLog(event, "GrantAuctionRewardSuccessful", log); err != nil {
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

// ParseGrantAuctionRewardSuccessful is a log parse operation binding the contract event 0xa307d88e125ed19583b02332c5bd74f0323774248b3f9cac5047c2433d9d8ddc.
//
// Solidity: event GrantAuctionRewardSuccessful(address auctionHighestBidder, bytes32 auctionId, uint256 assetId)
func (_IAuction *IAuctionFilterer) ParseGrantAuctionRewardSuccessful(log types.Log) (*IAuctionGrantAuctionRewardSuccessful, error) {
	event := new(IAuctionGrantAuctionRewardSuccessful)
	if err := _IAuction.contract.UnpackLog(event, "GrantAuctionRewardSuccessful", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
