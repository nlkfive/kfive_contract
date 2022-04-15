// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package BlindAuction

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

// BlindAuctionMetaData contains all meta data concerning the BlindAuction contract.
var BlindAuctionMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"AuctionEndAlreadyCalled\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"}],\"name\":\"TooEarly\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"}],\"name\":\"TooLate\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"nftAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"assetId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"deposit\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"bidder\",\"type\":\"address\"}],\"name\":\"AuctionRefund\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"bidder\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"nftAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"assetId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"blindedBid\",\"type\":\"bytes32\"}],\"name\":\"BidSuccessful\",\"type\":\"event\"}]",
}

// BlindAuctionABI is the input ABI used to generate the binding from.
// Deprecated: Use BlindAuctionMetaData.ABI instead.
var BlindAuctionABI = BlindAuctionMetaData.ABI

// BlindAuction is an auto generated Go binding around an Ethereum contract.
type BlindAuction struct {
	BlindAuctionCaller     // Read-only binding to the contract
	BlindAuctionTransactor // Write-only binding to the contract
	BlindAuctionFilterer   // Log filterer for contract events
}

// BlindAuctionCaller is an auto generated read-only Go binding around an Ethereum contract.
type BlindAuctionCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BlindAuctionTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BlindAuctionTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BlindAuctionFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BlindAuctionFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BlindAuctionSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BlindAuctionSession struct {
	Contract     *BlindAuction     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BlindAuctionCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BlindAuctionCallerSession struct {
	Contract *BlindAuctionCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// BlindAuctionTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BlindAuctionTransactorSession struct {
	Contract     *BlindAuctionTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// BlindAuctionRaw is an auto generated low-level Go binding around an Ethereum contract.
type BlindAuctionRaw struct {
	Contract *BlindAuction // Generic contract binding to access the raw methods on
}

// BlindAuctionCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BlindAuctionCallerRaw struct {
	Contract *BlindAuctionCaller // Generic read-only contract binding to access the raw methods on
}

// BlindAuctionTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BlindAuctionTransactorRaw struct {
	Contract *BlindAuctionTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBlindAuction creates a new instance of BlindAuction, bound to a specific deployed contract.
func NewBlindAuction(address common.Address, backend bind.ContractBackend) (*BlindAuction, error) {
	contract, err := bindBlindAuction(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BlindAuction{BlindAuctionCaller: BlindAuctionCaller{contract: contract}, BlindAuctionTransactor: BlindAuctionTransactor{contract: contract}, BlindAuctionFilterer: BlindAuctionFilterer{contract: contract}}, nil
}

// NewBlindAuctionCaller creates a new read-only instance of BlindAuction, bound to a specific deployed contract.
func NewBlindAuctionCaller(address common.Address, caller bind.ContractCaller) (*BlindAuctionCaller, error) {
	contract, err := bindBlindAuction(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BlindAuctionCaller{contract: contract}, nil
}

// NewBlindAuctionTransactor creates a new write-only instance of BlindAuction, bound to a specific deployed contract.
func NewBlindAuctionTransactor(address common.Address, transactor bind.ContractTransactor) (*BlindAuctionTransactor, error) {
	contract, err := bindBlindAuction(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BlindAuctionTransactor{contract: contract}, nil
}

// NewBlindAuctionFilterer creates a new log filterer instance of BlindAuction, bound to a specific deployed contract.
func NewBlindAuctionFilterer(address common.Address, filterer bind.ContractFilterer) (*BlindAuctionFilterer, error) {
	contract, err := bindBlindAuction(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BlindAuctionFilterer{contract: contract}, nil
}

// bindBlindAuction binds a generic wrapper to an already deployed contract.
func bindBlindAuction(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BlindAuctionABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BlindAuction *BlindAuctionRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BlindAuction.Contract.BlindAuctionCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BlindAuction *BlindAuctionRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BlindAuction.Contract.BlindAuctionTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BlindAuction *BlindAuctionRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BlindAuction.Contract.BlindAuctionTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BlindAuction *BlindAuctionCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BlindAuction.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BlindAuction *BlindAuctionTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BlindAuction.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BlindAuction *BlindAuctionTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BlindAuction.Contract.contract.Transact(opts, method, params...)
}

// BlindAuctionAuctionRefundIterator is returned from FilterAuctionRefund and is used to iterate over the raw logs and unpacked data for AuctionRefund events raised by the BlindAuction contract.
type BlindAuctionAuctionRefundIterator struct {
	Event *BlindAuctionAuctionRefund // Event containing the contract specifics and raw log

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
func (it *BlindAuctionAuctionRefundIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BlindAuctionAuctionRefund)
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
		it.Event = new(BlindAuctionAuctionRefund)
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
func (it *BlindAuctionAuctionRefundIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BlindAuctionAuctionRefundIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BlindAuctionAuctionRefund represents a AuctionRefund event raised by the BlindAuction contract.
type BlindAuctionAuctionRefund struct {
	NftAddress common.Address
	AssetId    *big.Int
	Deposit    *big.Int
	Bidder     common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterAuctionRefund is a free log retrieval operation binding the contract event 0xe71ce726e7a1cceb31d105962067824bc8ebd189ffaa94839b1ac96e6a8d38af.
//
// Solidity: event AuctionRefund(address nftAddress, uint256 indexed assetId, uint256 deposit, address indexed bidder)
func (_BlindAuction *BlindAuctionFilterer) FilterAuctionRefund(opts *bind.FilterOpts, assetId []*big.Int, bidder []common.Address) (*BlindAuctionAuctionRefundIterator, error) {

	var assetIdRule []interface{}
	for _, assetIdItem := range assetId {
		assetIdRule = append(assetIdRule, assetIdItem)
	}

	var bidderRule []interface{}
	for _, bidderItem := range bidder {
		bidderRule = append(bidderRule, bidderItem)
	}

	logs, sub, err := _BlindAuction.contract.FilterLogs(opts, "AuctionRefund", assetIdRule, bidderRule)
	if err != nil {
		return nil, err
	}
	return &BlindAuctionAuctionRefundIterator{contract: _BlindAuction.contract, event: "AuctionRefund", logs: logs, sub: sub}, nil
}

// WatchAuctionRefund is a free log subscription operation binding the contract event 0xe71ce726e7a1cceb31d105962067824bc8ebd189ffaa94839b1ac96e6a8d38af.
//
// Solidity: event AuctionRefund(address nftAddress, uint256 indexed assetId, uint256 deposit, address indexed bidder)
func (_BlindAuction *BlindAuctionFilterer) WatchAuctionRefund(opts *bind.WatchOpts, sink chan<- *BlindAuctionAuctionRefund, assetId []*big.Int, bidder []common.Address) (event.Subscription, error) {

	var assetIdRule []interface{}
	for _, assetIdItem := range assetId {
		assetIdRule = append(assetIdRule, assetIdItem)
	}

	var bidderRule []interface{}
	for _, bidderItem := range bidder {
		bidderRule = append(bidderRule, bidderItem)
	}

	logs, sub, err := _BlindAuction.contract.WatchLogs(opts, "AuctionRefund", assetIdRule, bidderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BlindAuctionAuctionRefund)
				if err := _BlindAuction.contract.UnpackLog(event, "AuctionRefund", log); err != nil {
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
func (_BlindAuction *BlindAuctionFilterer) ParseAuctionRefund(log types.Log) (*BlindAuctionAuctionRefund, error) {
	event := new(BlindAuctionAuctionRefund)
	if err := _BlindAuction.contract.UnpackLog(event, "AuctionRefund", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BlindAuctionBidSuccessfulIterator is returned from FilterBidSuccessful and is used to iterate over the raw logs and unpacked data for BidSuccessful events raised by the BlindAuction contract.
type BlindAuctionBidSuccessfulIterator struct {
	Event *BlindAuctionBidSuccessful // Event containing the contract specifics and raw log

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
func (it *BlindAuctionBidSuccessfulIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BlindAuctionBidSuccessful)
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
		it.Event = new(BlindAuctionBidSuccessful)
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
func (it *BlindAuctionBidSuccessfulIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BlindAuctionBidSuccessfulIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BlindAuctionBidSuccessful represents a BidSuccessful event raised by the BlindAuction contract.
type BlindAuctionBidSuccessful struct {
	Bidder     common.Address
	NftAddress common.Address
	AssetId    *big.Int
	BlindedBid [32]byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterBidSuccessful is a free log retrieval operation binding the contract event 0xc9b646ec968ab42359a60ba55be704c9d9c49f1db98b242275274b18763eefda.
//
// Solidity: event BidSuccessful(address indexed bidder, address nftAddress, uint256 indexed assetId, bytes32 blindedBid)
func (_BlindAuction *BlindAuctionFilterer) FilterBidSuccessful(opts *bind.FilterOpts, bidder []common.Address, assetId []*big.Int) (*BlindAuctionBidSuccessfulIterator, error) {

	var bidderRule []interface{}
	for _, bidderItem := range bidder {
		bidderRule = append(bidderRule, bidderItem)
	}

	var assetIdRule []interface{}
	for _, assetIdItem := range assetId {
		assetIdRule = append(assetIdRule, assetIdItem)
	}

	logs, sub, err := _BlindAuction.contract.FilterLogs(opts, "BidSuccessful", bidderRule, assetIdRule)
	if err != nil {
		return nil, err
	}
	return &BlindAuctionBidSuccessfulIterator{contract: _BlindAuction.contract, event: "BidSuccessful", logs: logs, sub: sub}, nil
}

// WatchBidSuccessful is a free log subscription operation binding the contract event 0xc9b646ec968ab42359a60ba55be704c9d9c49f1db98b242275274b18763eefda.
//
// Solidity: event BidSuccessful(address indexed bidder, address nftAddress, uint256 indexed assetId, bytes32 blindedBid)
func (_BlindAuction *BlindAuctionFilterer) WatchBidSuccessful(opts *bind.WatchOpts, sink chan<- *BlindAuctionBidSuccessful, bidder []common.Address, assetId []*big.Int) (event.Subscription, error) {

	var bidderRule []interface{}
	for _, bidderItem := range bidder {
		bidderRule = append(bidderRule, bidderItem)
	}

	var assetIdRule []interface{}
	for _, assetIdItem := range assetId {
		assetIdRule = append(assetIdRule, assetIdItem)
	}

	logs, sub, err := _BlindAuction.contract.WatchLogs(opts, "BidSuccessful", bidderRule, assetIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BlindAuctionBidSuccessful)
				if err := _BlindAuction.contract.UnpackLog(event, "BidSuccessful", log); err != nil {
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
func (_BlindAuction *BlindAuctionFilterer) ParseBidSuccessful(log types.Log) (*BlindAuctionBidSuccessful, error) {
	event := new(BlindAuctionBidSuccessful)
	if err := _BlindAuction.contract.UnpackLog(event, "BidSuccessful", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
