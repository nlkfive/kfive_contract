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
	ABI: "[{\"inputs\":[],\"name\":\"InvalidBiddingEnd\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidBiddingPrice\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotRunning\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotStarted\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotWinner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RewardGranted\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"TooEarly\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"TooLate\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"canceller\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"auctionId\",\"type\":\"bytes32\"}],\"name\":\"AuctionCancelledSuccessful\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"bidder\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"auctionId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"AuctionRefundSuccessful\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"auctionHighestBidder\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"auctionId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"assetId\",\"type\":\"uint256\"}],\"name\":\"GrantAuctionRewardSuccessful\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"publicAuctionId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"bidValue\",\"type\":\"uint256\"}],\"name\":\"PublicAuctionBidSuccessful\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"assetOwner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"nftAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"publicAuctionId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"assetId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"biddingEnd\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"startPriceInWei\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"minIncrement\",\"type\":\"uint256\"}],\"name\":\"PublicAuctionCreatedSuccessful\",\"type\":\"event\"}]",
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

// IPublicAuctionAuctionCancelledSuccessfulIterator is returned from FilterAuctionCancelledSuccessful and is used to iterate over the raw logs and unpacked data for AuctionCancelledSuccessful events raised by the IPublicAuction contract.
type IPublicAuctionAuctionCancelledSuccessfulIterator struct {
	Event *IPublicAuctionAuctionCancelledSuccessful // Event containing the contract specifics and raw log

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
func (it *IPublicAuctionAuctionCancelledSuccessfulIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IPublicAuctionAuctionCancelledSuccessful)
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
		it.Event = new(IPublicAuctionAuctionCancelledSuccessful)
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
func (it *IPublicAuctionAuctionCancelledSuccessfulIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IPublicAuctionAuctionCancelledSuccessfulIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IPublicAuctionAuctionCancelledSuccessful represents a AuctionCancelledSuccessful event raised by the IPublicAuction contract.
type IPublicAuctionAuctionCancelledSuccessful struct {
	Canceller common.Address
	AuctionId [32]byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterAuctionCancelledSuccessful is a free log retrieval operation binding the contract event 0xa021239bb373fef19aa7c7ef798961ab038ebf95bd85d546b731d25855db190e.
//
// Solidity: event AuctionCancelledSuccessful(address canceller, bytes32 auctionId)
func (_IPublicAuction *IPublicAuctionFilterer) FilterAuctionCancelledSuccessful(opts *bind.FilterOpts) (*IPublicAuctionAuctionCancelledSuccessfulIterator, error) {

	logs, sub, err := _IPublicAuction.contract.FilterLogs(opts, "AuctionCancelledSuccessful")
	if err != nil {
		return nil, err
	}
	return &IPublicAuctionAuctionCancelledSuccessfulIterator{contract: _IPublicAuction.contract, event: "AuctionCancelledSuccessful", logs: logs, sub: sub}, nil
}

// WatchAuctionCancelledSuccessful is a free log subscription operation binding the contract event 0xa021239bb373fef19aa7c7ef798961ab038ebf95bd85d546b731d25855db190e.
//
// Solidity: event AuctionCancelledSuccessful(address canceller, bytes32 auctionId)
func (_IPublicAuction *IPublicAuctionFilterer) WatchAuctionCancelledSuccessful(opts *bind.WatchOpts, sink chan<- *IPublicAuctionAuctionCancelledSuccessful) (event.Subscription, error) {

	logs, sub, err := _IPublicAuction.contract.WatchLogs(opts, "AuctionCancelledSuccessful")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IPublicAuctionAuctionCancelledSuccessful)
				if err := _IPublicAuction.contract.UnpackLog(event, "AuctionCancelledSuccessful", log); err != nil {
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
func (_IPublicAuction *IPublicAuctionFilterer) ParseAuctionCancelledSuccessful(log types.Log) (*IPublicAuctionAuctionCancelledSuccessful, error) {
	event := new(IPublicAuctionAuctionCancelledSuccessful)
	if err := _IPublicAuction.contract.UnpackLog(event, "AuctionCancelledSuccessful", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IPublicAuctionAuctionRefundSuccessfulIterator is returned from FilterAuctionRefundSuccessful and is used to iterate over the raw logs and unpacked data for AuctionRefundSuccessful events raised by the IPublicAuction contract.
type IPublicAuctionAuctionRefundSuccessfulIterator struct {
	Event *IPublicAuctionAuctionRefundSuccessful // Event containing the contract specifics and raw log

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
func (it *IPublicAuctionAuctionRefundSuccessfulIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IPublicAuctionAuctionRefundSuccessful)
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
		it.Event = new(IPublicAuctionAuctionRefundSuccessful)
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
func (it *IPublicAuctionAuctionRefundSuccessfulIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IPublicAuctionAuctionRefundSuccessfulIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IPublicAuctionAuctionRefundSuccessful represents a AuctionRefundSuccessful event raised by the IPublicAuction contract.
type IPublicAuctionAuctionRefundSuccessful struct {
	Bidder    common.Address
	AuctionId [32]byte
	Value     *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterAuctionRefundSuccessful is a free log retrieval operation binding the contract event 0x5f4b88832be3746851e4d0ce6129f89fe39ddcf5a4b0204699724c459baf7d9d.
//
// Solidity: event AuctionRefundSuccessful(address bidder, bytes32 auctionId, uint256 value)
func (_IPublicAuction *IPublicAuctionFilterer) FilterAuctionRefundSuccessful(opts *bind.FilterOpts) (*IPublicAuctionAuctionRefundSuccessfulIterator, error) {

	logs, sub, err := _IPublicAuction.contract.FilterLogs(opts, "AuctionRefundSuccessful")
	if err != nil {
		return nil, err
	}
	return &IPublicAuctionAuctionRefundSuccessfulIterator{contract: _IPublicAuction.contract, event: "AuctionRefundSuccessful", logs: logs, sub: sub}, nil
}

// WatchAuctionRefundSuccessful is a free log subscription operation binding the contract event 0x5f4b88832be3746851e4d0ce6129f89fe39ddcf5a4b0204699724c459baf7d9d.
//
// Solidity: event AuctionRefundSuccessful(address bidder, bytes32 auctionId, uint256 value)
func (_IPublicAuction *IPublicAuctionFilterer) WatchAuctionRefundSuccessful(opts *bind.WatchOpts, sink chan<- *IPublicAuctionAuctionRefundSuccessful) (event.Subscription, error) {

	logs, sub, err := _IPublicAuction.contract.WatchLogs(opts, "AuctionRefundSuccessful")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IPublicAuctionAuctionRefundSuccessful)
				if err := _IPublicAuction.contract.UnpackLog(event, "AuctionRefundSuccessful", log); err != nil {
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
func (_IPublicAuction *IPublicAuctionFilterer) ParseAuctionRefundSuccessful(log types.Log) (*IPublicAuctionAuctionRefundSuccessful, error) {
	event := new(IPublicAuctionAuctionRefundSuccessful)
	if err := _IPublicAuction.contract.UnpackLog(event, "AuctionRefundSuccessful", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IPublicAuctionGrantAuctionRewardSuccessfulIterator is returned from FilterGrantAuctionRewardSuccessful and is used to iterate over the raw logs and unpacked data for GrantAuctionRewardSuccessful events raised by the IPublicAuction contract.
type IPublicAuctionGrantAuctionRewardSuccessfulIterator struct {
	Event *IPublicAuctionGrantAuctionRewardSuccessful // Event containing the contract specifics and raw log

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
func (it *IPublicAuctionGrantAuctionRewardSuccessfulIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IPublicAuctionGrantAuctionRewardSuccessful)
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
		it.Event = new(IPublicAuctionGrantAuctionRewardSuccessful)
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
func (it *IPublicAuctionGrantAuctionRewardSuccessfulIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IPublicAuctionGrantAuctionRewardSuccessfulIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IPublicAuctionGrantAuctionRewardSuccessful represents a GrantAuctionRewardSuccessful event raised by the IPublicAuction contract.
type IPublicAuctionGrantAuctionRewardSuccessful struct {
	AuctionHighestBidder common.Address
	AuctionId            [32]byte
	AssetId              *big.Int
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterGrantAuctionRewardSuccessful is a free log retrieval operation binding the contract event 0xa307d88e125ed19583b02332c5bd74f0323774248b3f9cac5047c2433d9d8ddc.
//
// Solidity: event GrantAuctionRewardSuccessful(address auctionHighestBidder, bytes32 auctionId, uint256 assetId)
func (_IPublicAuction *IPublicAuctionFilterer) FilterGrantAuctionRewardSuccessful(opts *bind.FilterOpts) (*IPublicAuctionGrantAuctionRewardSuccessfulIterator, error) {

	logs, sub, err := _IPublicAuction.contract.FilterLogs(opts, "GrantAuctionRewardSuccessful")
	if err != nil {
		return nil, err
	}
	return &IPublicAuctionGrantAuctionRewardSuccessfulIterator{contract: _IPublicAuction.contract, event: "GrantAuctionRewardSuccessful", logs: logs, sub: sub}, nil
}

// WatchGrantAuctionRewardSuccessful is a free log subscription operation binding the contract event 0xa307d88e125ed19583b02332c5bd74f0323774248b3f9cac5047c2433d9d8ddc.
//
// Solidity: event GrantAuctionRewardSuccessful(address auctionHighestBidder, bytes32 auctionId, uint256 assetId)
func (_IPublicAuction *IPublicAuctionFilterer) WatchGrantAuctionRewardSuccessful(opts *bind.WatchOpts, sink chan<- *IPublicAuctionGrantAuctionRewardSuccessful) (event.Subscription, error) {

	logs, sub, err := _IPublicAuction.contract.WatchLogs(opts, "GrantAuctionRewardSuccessful")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IPublicAuctionGrantAuctionRewardSuccessful)
				if err := _IPublicAuction.contract.UnpackLog(event, "GrantAuctionRewardSuccessful", log); err != nil {
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
func (_IPublicAuction *IPublicAuctionFilterer) ParseGrantAuctionRewardSuccessful(log types.Log) (*IPublicAuctionGrantAuctionRewardSuccessful, error) {
	event := new(IPublicAuctionGrantAuctionRewardSuccessful)
	if err := _IPublicAuction.contract.UnpackLog(event, "GrantAuctionRewardSuccessful", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
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
	Sender          common.Address
	PublicAuctionId [32]byte
	BidValue        *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterPublicAuctionBidSuccessful is a free log retrieval operation binding the contract event 0x9f541988b3aac74ac43a6f4b330ceed668244b74461b90850d199a38a172f7c2.
//
// Solidity: event PublicAuctionBidSuccessful(address sender, bytes32 publicAuctionId, uint256 bidValue)
func (_IPublicAuction *IPublicAuctionFilterer) FilterPublicAuctionBidSuccessful(opts *bind.FilterOpts) (*IPublicAuctionPublicAuctionBidSuccessfulIterator, error) {

	logs, sub, err := _IPublicAuction.contract.FilterLogs(opts, "PublicAuctionBidSuccessful")
	if err != nil {
		return nil, err
	}
	return &IPublicAuctionPublicAuctionBidSuccessfulIterator{contract: _IPublicAuction.contract, event: "PublicAuctionBidSuccessful", logs: logs, sub: sub}, nil
}

// WatchPublicAuctionBidSuccessful is a free log subscription operation binding the contract event 0x9f541988b3aac74ac43a6f4b330ceed668244b74461b90850d199a38a172f7c2.
//
// Solidity: event PublicAuctionBidSuccessful(address sender, bytes32 publicAuctionId, uint256 bidValue)
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
// Solidity: event PublicAuctionBidSuccessful(address sender, bytes32 publicAuctionId, uint256 bidValue)
func (_IPublicAuction *IPublicAuctionFilterer) ParsePublicAuctionBidSuccessful(log types.Log) (*IPublicAuctionPublicAuctionBidSuccessful, error) {
	event := new(IPublicAuctionPublicAuctionBidSuccessful)
	if err := _IPublicAuction.contract.UnpackLog(event, "PublicAuctionBidSuccessful", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IPublicAuctionPublicAuctionCreatedSuccessfulIterator is returned from FilterPublicAuctionCreatedSuccessful and is used to iterate over the raw logs and unpacked data for PublicAuctionCreatedSuccessful events raised by the IPublicAuction contract.
type IPublicAuctionPublicAuctionCreatedSuccessfulIterator struct {
	Event *IPublicAuctionPublicAuctionCreatedSuccessful // Event containing the contract specifics and raw log

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
func (it *IPublicAuctionPublicAuctionCreatedSuccessfulIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IPublicAuctionPublicAuctionCreatedSuccessful)
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
		it.Event = new(IPublicAuctionPublicAuctionCreatedSuccessful)
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
func (it *IPublicAuctionPublicAuctionCreatedSuccessfulIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IPublicAuctionPublicAuctionCreatedSuccessfulIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IPublicAuctionPublicAuctionCreatedSuccessful represents a PublicAuctionCreatedSuccessful event raised by the IPublicAuction contract.
type IPublicAuctionPublicAuctionCreatedSuccessful struct {
	AssetOwner      common.Address
	NftAddress      common.Address
	PublicAuctionId [32]byte
	AssetId         *big.Int
	BiddingEnd      *big.Int
	StartPriceInWei *big.Int
	StartTime       *big.Int
	MinIncrement    *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterPublicAuctionCreatedSuccessful is a free log retrieval operation binding the contract event 0xd3c25d0f5cc0da484b0fc13f082a5e7da800d235500c5904281dbdbb9e22ba90.
//
// Solidity: event PublicAuctionCreatedSuccessful(address assetOwner, address nftAddress, bytes32 publicAuctionId, uint256 assetId, uint256 biddingEnd, uint256 startPriceInWei, uint256 startTime, uint256 minIncrement)
func (_IPublicAuction *IPublicAuctionFilterer) FilterPublicAuctionCreatedSuccessful(opts *bind.FilterOpts) (*IPublicAuctionPublicAuctionCreatedSuccessfulIterator, error) {

	logs, sub, err := _IPublicAuction.contract.FilterLogs(opts, "PublicAuctionCreatedSuccessful")
	if err != nil {
		return nil, err
	}
	return &IPublicAuctionPublicAuctionCreatedSuccessfulIterator{contract: _IPublicAuction.contract, event: "PublicAuctionCreatedSuccessful", logs: logs, sub: sub}, nil
}

// WatchPublicAuctionCreatedSuccessful is a free log subscription operation binding the contract event 0xd3c25d0f5cc0da484b0fc13f082a5e7da800d235500c5904281dbdbb9e22ba90.
//
// Solidity: event PublicAuctionCreatedSuccessful(address assetOwner, address nftAddress, bytes32 publicAuctionId, uint256 assetId, uint256 biddingEnd, uint256 startPriceInWei, uint256 startTime, uint256 minIncrement)
func (_IPublicAuction *IPublicAuctionFilterer) WatchPublicAuctionCreatedSuccessful(opts *bind.WatchOpts, sink chan<- *IPublicAuctionPublicAuctionCreatedSuccessful) (event.Subscription, error) {

	logs, sub, err := _IPublicAuction.contract.WatchLogs(opts, "PublicAuctionCreatedSuccessful")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IPublicAuctionPublicAuctionCreatedSuccessful)
				if err := _IPublicAuction.contract.UnpackLog(event, "PublicAuctionCreatedSuccessful", log); err != nil {
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

// ParsePublicAuctionCreatedSuccessful is a log parse operation binding the contract event 0xd3c25d0f5cc0da484b0fc13f082a5e7da800d235500c5904281dbdbb9e22ba90.
//
// Solidity: event PublicAuctionCreatedSuccessful(address assetOwner, address nftAddress, bytes32 publicAuctionId, uint256 assetId, uint256 biddingEnd, uint256 startPriceInWei, uint256 startTime, uint256 minIncrement)
func (_IPublicAuction *IPublicAuctionFilterer) ParsePublicAuctionCreatedSuccessful(log types.Log) (*IPublicAuctionPublicAuctionCreatedSuccessful, error) {
	event := new(IPublicAuctionPublicAuctionCreatedSuccessful)
	if err := _IPublicAuction.contract.UnpackLog(event, "PublicAuctionCreatedSuccessful", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
