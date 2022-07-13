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
	ABI: "[{\"inputs\":[],\"name\":\"AuctionEnded\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidBiddingEnd\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidBiddingPrice\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidReveal\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidWithdraw\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotBidYet\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotRunning\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RewardGranted\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"TooEarly\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"TooLate\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"canceller\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"auctionId\",\"type\":\"bytes32\"}],\"name\":\"AuctionCancelledSuccessful\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"bidder\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"auctionId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"AuctionRefundSuccessful\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"blindAuctionId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"blindedBid\",\"type\":\"bytes32\"}],\"name\":\"BlindAuctionBidSuccessful\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"assetOwner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"nftAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"blindAuctionId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"assetId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"biddingEnd\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"revealEnd\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"startPriceInWei\",\"type\":\"uint256\"}],\"name\":\"BlindAuctionCreatedSuccessful\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"auctionHighestBidder\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"auctionId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"assetId\",\"type\":\"uint256\"}],\"name\":\"GrantAuctionRewardSuccessful\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"blindAuctionId\",\"type\":\"bytes32\"}],\"name\":\"RevealSuccessful\",\"type\":\"event\"}]",
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

// IBlindAuctionAuctionCancelledSuccessfulIterator is returned from FilterAuctionCancelledSuccessful and is used to iterate over the raw logs and unpacked data for AuctionCancelledSuccessful events raised by the IBlindAuction contract.
type IBlindAuctionAuctionCancelledSuccessfulIterator struct {
	Event *IBlindAuctionAuctionCancelledSuccessful // Event containing the contract specifics and raw log

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
func (it *IBlindAuctionAuctionCancelledSuccessfulIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IBlindAuctionAuctionCancelledSuccessful)
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
		it.Event = new(IBlindAuctionAuctionCancelledSuccessful)
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
func (it *IBlindAuctionAuctionCancelledSuccessfulIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IBlindAuctionAuctionCancelledSuccessfulIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IBlindAuctionAuctionCancelledSuccessful represents a AuctionCancelledSuccessful event raised by the IBlindAuction contract.
type IBlindAuctionAuctionCancelledSuccessful struct {
	Canceller common.Address
	AuctionId [32]byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterAuctionCancelledSuccessful is a free log retrieval operation binding the contract event 0xa021239bb373fef19aa7c7ef798961ab038ebf95bd85d546b731d25855db190e.
//
// Solidity: event AuctionCancelledSuccessful(address canceller, bytes32 auctionId)
func (_IBlindAuction *IBlindAuctionFilterer) FilterAuctionCancelledSuccessful(opts *bind.FilterOpts) (*IBlindAuctionAuctionCancelledSuccessfulIterator, error) {

	logs, sub, err := _IBlindAuction.contract.FilterLogs(opts, "AuctionCancelledSuccessful")
	if err != nil {
		return nil, err
	}
	return &IBlindAuctionAuctionCancelledSuccessfulIterator{contract: _IBlindAuction.contract, event: "AuctionCancelledSuccessful", logs: logs, sub: sub}, nil
}

// WatchAuctionCancelledSuccessful is a free log subscription operation binding the contract event 0xa021239bb373fef19aa7c7ef798961ab038ebf95bd85d546b731d25855db190e.
//
// Solidity: event AuctionCancelledSuccessful(address canceller, bytes32 auctionId)
func (_IBlindAuction *IBlindAuctionFilterer) WatchAuctionCancelledSuccessful(opts *bind.WatchOpts, sink chan<- *IBlindAuctionAuctionCancelledSuccessful) (event.Subscription, error) {

	logs, sub, err := _IBlindAuction.contract.WatchLogs(opts, "AuctionCancelledSuccessful")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IBlindAuctionAuctionCancelledSuccessful)
				if err := _IBlindAuction.contract.UnpackLog(event, "AuctionCancelledSuccessful", log); err != nil {
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
func (_IBlindAuction *IBlindAuctionFilterer) ParseAuctionCancelledSuccessful(log types.Log) (*IBlindAuctionAuctionCancelledSuccessful, error) {
	event := new(IBlindAuctionAuctionCancelledSuccessful)
	if err := _IBlindAuction.contract.UnpackLog(event, "AuctionCancelledSuccessful", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IBlindAuctionAuctionRefundSuccessfulIterator is returned from FilterAuctionRefundSuccessful and is used to iterate over the raw logs and unpacked data for AuctionRefundSuccessful events raised by the IBlindAuction contract.
type IBlindAuctionAuctionRefundSuccessfulIterator struct {
	Event *IBlindAuctionAuctionRefundSuccessful // Event containing the contract specifics and raw log

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
func (it *IBlindAuctionAuctionRefundSuccessfulIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IBlindAuctionAuctionRefundSuccessful)
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
		it.Event = new(IBlindAuctionAuctionRefundSuccessful)
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
func (it *IBlindAuctionAuctionRefundSuccessfulIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IBlindAuctionAuctionRefundSuccessfulIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IBlindAuctionAuctionRefundSuccessful represents a AuctionRefundSuccessful event raised by the IBlindAuction contract.
type IBlindAuctionAuctionRefundSuccessful struct {
	Bidder    common.Address
	AuctionId [32]byte
	Value     *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterAuctionRefundSuccessful is a free log retrieval operation binding the contract event 0x5f4b88832be3746851e4d0ce6129f89fe39ddcf5a4b0204699724c459baf7d9d.
//
// Solidity: event AuctionRefundSuccessful(address bidder, bytes32 auctionId, uint256 value)
func (_IBlindAuction *IBlindAuctionFilterer) FilterAuctionRefundSuccessful(opts *bind.FilterOpts) (*IBlindAuctionAuctionRefundSuccessfulIterator, error) {

	logs, sub, err := _IBlindAuction.contract.FilterLogs(opts, "AuctionRefundSuccessful")
	if err != nil {
		return nil, err
	}
	return &IBlindAuctionAuctionRefundSuccessfulIterator{contract: _IBlindAuction.contract, event: "AuctionRefundSuccessful", logs: logs, sub: sub}, nil
}

// WatchAuctionRefundSuccessful is a free log subscription operation binding the contract event 0x5f4b88832be3746851e4d0ce6129f89fe39ddcf5a4b0204699724c459baf7d9d.
//
// Solidity: event AuctionRefundSuccessful(address bidder, bytes32 auctionId, uint256 value)
func (_IBlindAuction *IBlindAuctionFilterer) WatchAuctionRefundSuccessful(opts *bind.WatchOpts, sink chan<- *IBlindAuctionAuctionRefundSuccessful) (event.Subscription, error) {

	logs, sub, err := _IBlindAuction.contract.WatchLogs(opts, "AuctionRefundSuccessful")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IBlindAuctionAuctionRefundSuccessful)
				if err := _IBlindAuction.contract.UnpackLog(event, "AuctionRefundSuccessful", log); err != nil {
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
func (_IBlindAuction *IBlindAuctionFilterer) ParseAuctionRefundSuccessful(log types.Log) (*IBlindAuctionAuctionRefundSuccessful, error) {
	event := new(IBlindAuctionAuctionRefundSuccessful)
	if err := _IBlindAuction.contract.UnpackLog(event, "AuctionRefundSuccessful", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
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
	Sender         common.Address
	BlindAuctionId [32]byte
	BlindedBid     [32]byte
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterBlindAuctionBidSuccessful is a free log retrieval operation binding the contract event 0x42412bd3e1b349d0fc9c3518c85d09f611398d80842c5d3e36449b76dca01b6b.
//
// Solidity: event BlindAuctionBidSuccessful(address sender, bytes32 blindAuctionId, bytes32 blindedBid)
func (_IBlindAuction *IBlindAuctionFilterer) FilterBlindAuctionBidSuccessful(opts *bind.FilterOpts) (*IBlindAuctionBlindAuctionBidSuccessfulIterator, error) {

	logs, sub, err := _IBlindAuction.contract.FilterLogs(opts, "BlindAuctionBidSuccessful")
	if err != nil {
		return nil, err
	}
	return &IBlindAuctionBlindAuctionBidSuccessfulIterator{contract: _IBlindAuction.contract, event: "BlindAuctionBidSuccessful", logs: logs, sub: sub}, nil
}

// WatchBlindAuctionBidSuccessful is a free log subscription operation binding the contract event 0x42412bd3e1b349d0fc9c3518c85d09f611398d80842c5d3e36449b76dca01b6b.
//
// Solidity: event BlindAuctionBidSuccessful(address sender, bytes32 blindAuctionId, bytes32 blindedBid)
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
// Solidity: event BlindAuctionBidSuccessful(address sender, bytes32 blindAuctionId, bytes32 blindedBid)
func (_IBlindAuction *IBlindAuctionFilterer) ParseBlindAuctionBidSuccessful(log types.Log) (*IBlindAuctionBlindAuctionBidSuccessful, error) {
	event := new(IBlindAuctionBlindAuctionBidSuccessful)
	if err := _IBlindAuction.contract.UnpackLog(event, "BlindAuctionBidSuccessful", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IBlindAuctionBlindAuctionCreatedSuccessfulIterator is returned from FilterBlindAuctionCreatedSuccessful and is used to iterate over the raw logs and unpacked data for BlindAuctionCreatedSuccessful events raised by the IBlindAuction contract.
type IBlindAuctionBlindAuctionCreatedSuccessfulIterator struct {
	Event *IBlindAuctionBlindAuctionCreatedSuccessful // Event containing the contract specifics and raw log

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
func (it *IBlindAuctionBlindAuctionCreatedSuccessfulIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IBlindAuctionBlindAuctionCreatedSuccessful)
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
		it.Event = new(IBlindAuctionBlindAuctionCreatedSuccessful)
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
func (it *IBlindAuctionBlindAuctionCreatedSuccessfulIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IBlindAuctionBlindAuctionCreatedSuccessfulIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IBlindAuctionBlindAuctionCreatedSuccessful represents a BlindAuctionCreatedSuccessful event raised by the IBlindAuction contract.
type IBlindAuctionBlindAuctionCreatedSuccessful struct {
	AssetOwner      common.Address
	NftAddress      common.Address
	BlindAuctionId  [32]byte
	AssetId         *big.Int
	BiddingEnd      *big.Int
	RevealEnd       *big.Int
	StartPriceInWei *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterBlindAuctionCreatedSuccessful is a free log retrieval operation binding the contract event 0xde549e5fcbd4d262f76b1f3a6362777a2e669ce8a2a7b5460e00c7a38341baf7.
//
// Solidity: event BlindAuctionCreatedSuccessful(address assetOwner, address nftAddress, bytes32 blindAuctionId, uint256 assetId, uint256 biddingEnd, uint256 revealEnd, uint256 startPriceInWei)
func (_IBlindAuction *IBlindAuctionFilterer) FilterBlindAuctionCreatedSuccessful(opts *bind.FilterOpts) (*IBlindAuctionBlindAuctionCreatedSuccessfulIterator, error) {

	logs, sub, err := _IBlindAuction.contract.FilterLogs(opts, "BlindAuctionCreatedSuccessful")
	if err != nil {
		return nil, err
	}
	return &IBlindAuctionBlindAuctionCreatedSuccessfulIterator{contract: _IBlindAuction.contract, event: "BlindAuctionCreatedSuccessful", logs: logs, sub: sub}, nil
}

// WatchBlindAuctionCreatedSuccessful is a free log subscription operation binding the contract event 0xde549e5fcbd4d262f76b1f3a6362777a2e669ce8a2a7b5460e00c7a38341baf7.
//
// Solidity: event BlindAuctionCreatedSuccessful(address assetOwner, address nftAddress, bytes32 blindAuctionId, uint256 assetId, uint256 biddingEnd, uint256 revealEnd, uint256 startPriceInWei)
func (_IBlindAuction *IBlindAuctionFilterer) WatchBlindAuctionCreatedSuccessful(opts *bind.WatchOpts, sink chan<- *IBlindAuctionBlindAuctionCreatedSuccessful) (event.Subscription, error) {

	logs, sub, err := _IBlindAuction.contract.WatchLogs(opts, "BlindAuctionCreatedSuccessful")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IBlindAuctionBlindAuctionCreatedSuccessful)
				if err := _IBlindAuction.contract.UnpackLog(event, "BlindAuctionCreatedSuccessful", log); err != nil {
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

// ParseBlindAuctionCreatedSuccessful is a log parse operation binding the contract event 0xde549e5fcbd4d262f76b1f3a6362777a2e669ce8a2a7b5460e00c7a38341baf7.
//
// Solidity: event BlindAuctionCreatedSuccessful(address assetOwner, address nftAddress, bytes32 blindAuctionId, uint256 assetId, uint256 biddingEnd, uint256 revealEnd, uint256 startPriceInWei)
func (_IBlindAuction *IBlindAuctionFilterer) ParseBlindAuctionCreatedSuccessful(log types.Log) (*IBlindAuctionBlindAuctionCreatedSuccessful, error) {
	event := new(IBlindAuctionBlindAuctionCreatedSuccessful)
	if err := _IBlindAuction.contract.UnpackLog(event, "BlindAuctionCreatedSuccessful", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IBlindAuctionGrantAuctionRewardSuccessfulIterator is returned from FilterGrantAuctionRewardSuccessful and is used to iterate over the raw logs and unpacked data for GrantAuctionRewardSuccessful events raised by the IBlindAuction contract.
type IBlindAuctionGrantAuctionRewardSuccessfulIterator struct {
	Event *IBlindAuctionGrantAuctionRewardSuccessful // Event containing the contract specifics and raw log

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
func (it *IBlindAuctionGrantAuctionRewardSuccessfulIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IBlindAuctionGrantAuctionRewardSuccessful)
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
		it.Event = new(IBlindAuctionGrantAuctionRewardSuccessful)
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
func (it *IBlindAuctionGrantAuctionRewardSuccessfulIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IBlindAuctionGrantAuctionRewardSuccessfulIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IBlindAuctionGrantAuctionRewardSuccessful represents a GrantAuctionRewardSuccessful event raised by the IBlindAuction contract.
type IBlindAuctionGrantAuctionRewardSuccessful struct {
	AuctionHighestBidder common.Address
	AuctionId            [32]byte
	AssetId              *big.Int
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterGrantAuctionRewardSuccessful is a free log retrieval operation binding the contract event 0xa307d88e125ed19583b02332c5bd74f0323774248b3f9cac5047c2433d9d8ddc.
//
// Solidity: event GrantAuctionRewardSuccessful(address auctionHighestBidder, bytes32 auctionId, uint256 assetId)
func (_IBlindAuction *IBlindAuctionFilterer) FilterGrantAuctionRewardSuccessful(opts *bind.FilterOpts) (*IBlindAuctionGrantAuctionRewardSuccessfulIterator, error) {

	logs, sub, err := _IBlindAuction.contract.FilterLogs(opts, "GrantAuctionRewardSuccessful")
	if err != nil {
		return nil, err
	}
	return &IBlindAuctionGrantAuctionRewardSuccessfulIterator{contract: _IBlindAuction.contract, event: "GrantAuctionRewardSuccessful", logs: logs, sub: sub}, nil
}

// WatchGrantAuctionRewardSuccessful is a free log subscription operation binding the contract event 0xa307d88e125ed19583b02332c5bd74f0323774248b3f9cac5047c2433d9d8ddc.
//
// Solidity: event GrantAuctionRewardSuccessful(address auctionHighestBidder, bytes32 auctionId, uint256 assetId)
func (_IBlindAuction *IBlindAuctionFilterer) WatchGrantAuctionRewardSuccessful(opts *bind.WatchOpts, sink chan<- *IBlindAuctionGrantAuctionRewardSuccessful) (event.Subscription, error) {

	logs, sub, err := _IBlindAuction.contract.WatchLogs(opts, "GrantAuctionRewardSuccessful")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IBlindAuctionGrantAuctionRewardSuccessful)
				if err := _IBlindAuction.contract.UnpackLog(event, "GrantAuctionRewardSuccessful", log); err != nil {
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
func (_IBlindAuction *IBlindAuctionFilterer) ParseGrantAuctionRewardSuccessful(log types.Log) (*IBlindAuctionGrantAuctionRewardSuccessful, error) {
	event := new(IBlindAuctionGrantAuctionRewardSuccessful)
	if err := _IBlindAuction.contract.UnpackLog(event, "GrantAuctionRewardSuccessful", log); err != nil {
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
	Sender         common.Address
	BlindAuctionId [32]byte
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterRevealSuccessful is a free log retrieval operation binding the contract event 0x61d4ab60470c97b371c5f91be47873283386bb3af4de728e8795ef5318e013d6.
//
// Solidity: event RevealSuccessful(address sender, bytes32 blindAuctionId)
func (_IBlindAuction *IBlindAuctionFilterer) FilterRevealSuccessful(opts *bind.FilterOpts) (*IBlindAuctionRevealSuccessfulIterator, error) {

	logs, sub, err := _IBlindAuction.contract.FilterLogs(opts, "RevealSuccessful")
	if err != nil {
		return nil, err
	}
	return &IBlindAuctionRevealSuccessfulIterator{contract: _IBlindAuction.contract, event: "RevealSuccessful", logs: logs, sub: sub}, nil
}

// WatchRevealSuccessful is a free log subscription operation binding the contract event 0x61d4ab60470c97b371c5f91be47873283386bb3af4de728e8795ef5318e013d6.
//
// Solidity: event RevealSuccessful(address sender, bytes32 blindAuctionId)
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
// Solidity: event RevealSuccessful(address sender, bytes32 blindAuctionId)
func (_IBlindAuction *IBlindAuctionFilterer) ParseRevealSuccessful(log types.Log) (*IBlindAuctionRevealSuccessful, error) {
	event := new(IBlindAuctionRevealSuccessful)
	if err := _IBlindAuction.contract.UnpackLog(event, "RevealSuccessful", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
