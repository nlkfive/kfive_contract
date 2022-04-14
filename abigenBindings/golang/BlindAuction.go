// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package smc

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

// SmcMetaData contains all meta data concerning the Smc contract.
var SmcMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"AuctionEndAlreadyCalled\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"}],\"name\":\"TooEarly\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"}],\"name\":\"TooLate\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"nftAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"assetId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"deposit\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"bidder\",\"type\":\"address\"}],\"name\":\"AuctionRefund\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"bidder\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"nftAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"assetId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"blindedBid\",\"type\":\"bytes32\"}],\"name\":\"BidSuccessful\",\"type\":\"event\"}]",
}

// SmcABI is the input ABI used to generate the binding from.
// Deprecated: Use SmcMetaData.ABI instead.
var SmcABI = SmcMetaData.ABI

// Smc is an auto generated Go binding around an Ethereum contract.
type Smc struct {
	SmcCaller     // Read-only binding to the contract
	SmcTransactor // Write-only binding to the contract
	SmcFilterer   // Log filterer for contract events
}

// SmcCaller is an auto generated read-only Go binding around an Ethereum contract.
type SmcCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SmcTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SmcTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SmcFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SmcFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SmcSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SmcSession struct {
	Contract     *Smc              // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SmcCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SmcCallerSession struct {
	Contract *SmcCaller    // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// SmcTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SmcTransactorSession struct {
	Contract     *SmcTransactor    // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SmcRaw is an auto generated low-level Go binding around an Ethereum contract.
type SmcRaw struct {
	Contract *Smc // Generic contract binding to access the raw methods on
}

// SmcCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SmcCallerRaw struct {
	Contract *SmcCaller // Generic read-only contract binding to access the raw methods on
}

// SmcTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SmcTransactorRaw struct {
	Contract *SmcTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSmc creates a new instance of Smc, bound to a specific deployed contract.
func NewSmc(address common.Address, backend bind.ContractBackend) (*Smc, error) {
	contract, err := bindSmc(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Smc{SmcCaller: SmcCaller{contract: contract}, SmcTransactor: SmcTransactor{contract: contract}, SmcFilterer: SmcFilterer{contract: contract}}, nil
}

// NewSmcCaller creates a new read-only instance of Smc, bound to a specific deployed contract.
func NewSmcCaller(address common.Address, caller bind.ContractCaller) (*SmcCaller, error) {
	contract, err := bindSmc(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SmcCaller{contract: contract}, nil
}

// NewSmcTransactor creates a new write-only instance of Smc, bound to a specific deployed contract.
func NewSmcTransactor(address common.Address, transactor bind.ContractTransactor) (*SmcTransactor, error) {
	contract, err := bindSmc(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SmcTransactor{contract: contract}, nil
}

// NewSmcFilterer creates a new log filterer instance of Smc, bound to a specific deployed contract.
func NewSmcFilterer(address common.Address, filterer bind.ContractFilterer) (*SmcFilterer, error) {
	contract, err := bindSmc(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SmcFilterer{contract: contract}, nil
}

// bindSmc binds a generic wrapper to an already deployed contract.
func bindSmc(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SmcABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Smc *SmcRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Smc.Contract.SmcCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Smc *SmcRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Smc.Contract.SmcTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Smc *SmcRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Smc.Contract.SmcTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Smc *SmcCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Smc.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Smc *SmcTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Smc.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Smc *SmcTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Smc.Contract.contract.Transact(opts, method, params...)
}

// SmcAuctionRefundIterator is returned from FilterAuctionRefund and is used to iterate over the raw logs and unpacked data for AuctionRefund events raised by the Smc contract.
type SmcAuctionRefundIterator struct {
	Event *SmcAuctionRefund // Event containing the contract specifics and raw log

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
func (it *SmcAuctionRefundIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SmcAuctionRefund)
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
		it.Event = new(SmcAuctionRefund)
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
func (it *SmcAuctionRefundIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SmcAuctionRefundIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SmcAuctionRefund represents a AuctionRefund event raised by the Smc contract.
type SmcAuctionRefund struct {
	NftAddress common.Address
	AssetId    *big.Int
	Deposit    *big.Int
	Bidder     common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterAuctionRefund is a free log retrieval operation binding the contract event 0xe71ce726e7a1cceb31d105962067824bc8ebd189ffaa94839b1ac96e6a8d38af.
//
// Solidity: event AuctionRefund(address nftAddress, uint256 indexed assetId, uint256 deposit, address indexed bidder)
func (_Smc *SmcFilterer) FilterAuctionRefund(opts *bind.FilterOpts, assetId []*big.Int, bidder []common.Address) (*SmcAuctionRefundIterator, error) {

	var assetIdRule []interface{}
	for _, assetIdItem := range assetId {
		assetIdRule = append(assetIdRule, assetIdItem)
	}

	var bidderRule []interface{}
	for _, bidderItem := range bidder {
		bidderRule = append(bidderRule, bidderItem)
	}

	logs, sub, err := _Smc.contract.FilterLogs(opts, "AuctionRefund", assetIdRule, bidderRule)
	if err != nil {
		return nil, err
	}
	return &SmcAuctionRefundIterator{contract: _Smc.contract, event: "AuctionRefund", logs: logs, sub: sub}, nil
}

// WatchAuctionRefund is a free log subscription operation binding the contract event 0xe71ce726e7a1cceb31d105962067824bc8ebd189ffaa94839b1ac96e6a8d38af.
//
// Solidity: event AuctionRefund(address nftAddress, uint256 indexed assetId, uint256 deposit, address indexed bidder)
func (_Smc *SmcFilterer) WatchAuctionRefund(opts *bind.WatchOpts, sink chan<- *SmcAuctionRefund, assetId []*big.Int, bidder []common.Address) (event.Subscription, error) {

	var assetIdRule []interface{}
	for _, assetIdItem := range assetId {
		assetIdRule = append(assetIdRule, assetIdItem)
	}

	var bidderRule []interface{}
	for _, bidderItem := range bidder {
		bidderRule = append(bidderRule, bidderItem)
	}

	logs, sub, err := _Smc.contract.WatchLogs(opts, "AuctionRefund", assetIdRule, bidderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SmcAuctionRefund)
				if err := _Smc.contract.UnpackLog(event, "AuctionRefund", log); err != nil {
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

// ParseAuctionRefund is a log parse operation binding the contract event 0xe71ce726e7a1cceb31d105962067824bc8ebd189ffaa94839b1ac96e6a8d38af.
//
// Solidity: event AuctionRefund(address nftAddress, uint256 indexed assetId, uint256 deposit, address indexed bidder)
func (_Smc *SmcFilterer) ParseAuctionRefund(log types.Log) (*SmcAuctionRefund, error) {
	event := new(SmcAuctionRefund)
	if err := _Smc.contract.UnpackLog(event, "AuctionRefund", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SmcBidSuccessfulIterator is returned from FilterBidSuccessful and is used to iterate over the raw logs and unpacked data for BidSuccessful events raised by the Smc contract.
type SmcBidSuccessfulIterator struct {
	Event *SmcBidSuccessful // Event containing the contract specifics and raw log

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
func (it *SmcBidSuccessfulIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SmcBidSuccessful)
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
		it.Event = new(SmcBidSuccessful)
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
func (it *SmcBidSuccessfulIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SmcBidSuccessfulIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SmcBidSuccessful represents a BidSuccessful event raised by the Smc contract.
type SmcBidSuccessful struct {
	Bidder     common.Address
	NftAddress common.Address
	AssetId    *big.Int
	BlindedBid [32]byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterBidSuccessful is a free log retrieval operation binding the contract event 0xc9b646ec968ab42359a60ba55be704c9d9c49f1db98b242275274b18763eefda.
//
// Solidity: event BidSuccessful(address indexed bidder, address nftAddress, uint256 indexed assetId, bytes32 blindedBid)
func (_Smc *SmcFilterer) FilterBidSuccessful(opts *bind.FilterOpts, bidder []common.Address, assetId []*big.Int) (*SmcBidSuccessfulIterator, error) {

	var bidderRule []interface{}
	for _, bidderItem := range bidder {
		bidderRule = append(bidderRule, bidderItem)
	}

	var assetIdRule []interface{}
	for _, assetIdItem := range assetId {
		assetIdRule = append(assetIdRule, assetIdItem)
	}

	logs, sub, err := _Smc.contract.FilterLogs(opts, "BidSuccessful", bidderRule, assetIdRule)
	if err != nil {
		return nil, err
	}
	return &SmcBidSuccessfulIterator{contract: _Smc.contract, event: "BidSuccessful", logs: logs, sub: sub}, nil
}

// WatchBidSuccessful is a free log subscription operation binding the contract event 0xc9b646ec968ab42359a60ba55be704c9d9c49f1db98b242275274b18763eefda.
//
// Solidity: event BidSuccessful(address indexed bidder, address nftAddress, uint256 indexed assetId, bytes32 blindedBid)
func (_Smc *SmcFilterer) WatchBidSuccessful(opts *bind.WatchOpts, sink chan<- *SmcBidSuccessful, bidder []common.Address, assetId []*big.Int) (event.Subscription, error) {

	var bidderRule []interface{}
	for _, bidderItem := range bidder {
		bidderRule = append(bidderRule, bidderItem)
	}

	var assetIdRule []interface{}
	for _, assetIdItem := range assetId {
		assetIdRule = append(assetIdRule, assetIdItem)
	}

	logs, sub, err := _Smc.contract.WatchLogs(opts, "BidSuccessful", bidderRule, assetIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SmcBidSuccessful)
				if err := _Smc.contract.UnpackLog(event, "BidSuccessful", log); err != nil {
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

// ParseBidSuccessful is a log parse operation binding the contract event 0xc9b646ec968ab42359a60ba55be704c9d9c49f1db98b242275274b18763eefda.
//
// Solidity: event BidSuccessful(address indexed bidder, address nftAddress, uint256 indexed assetId, bytes32 blindedBid)
func (_Smc *SmcFilterer) ParseBidSuccessful(log types.Log) (*SmcBidSuccessful, error) {
	event := new(SmcBidSuccessful)
	if err := _Smc.contract.UnpackLog(event, "BidSuccessful", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
