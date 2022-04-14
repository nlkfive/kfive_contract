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
	ABI: "[{\"inputs\":[],\"name\":\"AuctionEndAlreadyCalled\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"}],\"name\":\"TooEarly\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"}],\"name\":\"TooLate\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"assetId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"nftAddress\",\"type\":\"address\"}],\"name\":\"AuctionCancelled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"assetId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"nftAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"biddingEnd\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"revealEnd\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"startPriceInWei\",\"type\":\"uint256\"}],\"name\":\"AuctionCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"nftAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"assetId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"deposit\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"bidder\",\"type\":\"address\"}],\"name\":\"AuctionRefund\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"assetId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"nftAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalPrice\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"buyer\",\"type\":\"address\"}],\"name\":\"AuctionSuccessful\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"bidder\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"nftAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"assetId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"blindedBid\",\"type\":\"bytes32\"}],\"name\":\"BidSuccessful\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"auctionByAssetId\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"ended\",\"type\":\"bool\"},{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"highestBidder\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"nftAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"biddingEnd\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"revealEnd\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"highestBid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startPrice\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
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

// AuctionByAssetId is a free data retrieval call binding the contract method 0xf4e58618.
//
// Solidity: function auctionByAssetId(address , uint256 , bytes32 ) view returns(bool ended, bytes32 id, address seller, address highestBidder, address nftAddress, uint256 biddingEnd, uint256 revealEnd, uint256 highestBid, uint256 startPrice)
func (_Smc *SmcCaller) AuctionByAssetId(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int, arg2 [32]byte) (struct {
	Ended         bool
	Id            [32]byte
	Seller        common.Address
	HighestBidder common.Address
	NftAddress    common.Address
	BiddingEnd    *big.Int
	RevealEnd     *big.Int
	HighestBid    *big.Int
	StartPrice    *big.Int
}, error) {
	var out []interface{}
	err := _Smc.contract.Call(opts, &out, "auctionByAssetId", arg0, arg1, arg2)

	outstruct := new(struct {
		Ended         bool
		Id            [32]byte
		Seller        common.Address
		HighestBidder common.Address
		NftAddress    common.Address
		BiddingEnd    *big.Int
		RevealEnd     *big.Int
		HighestBid    *big.Int
		StartPrice    *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Ended = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.Id = *abi.ConvertType(out[1], new([32]byte)).(*[32]byte)
	outstruct.Seller = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)
	outstruct.HighestBidder = *abi.ConvertType(out[3], new(common.Address)).(*common.Address)
	outstruct.NftAddress = *abi.ConvertType(out[4], new(common.Address)).(*common.Address)
	outstruct.BiddingEnd = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	outstruct.RevealEnd = *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)
	outstruct.HighestBid = *abi.ConvertType(out[7], new(*big.Int)).(**big.Int)
	outstruct.StartPrice = *abi.ConvertType(out[8], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// AuctionByAssetId is a free data retrieval call binding the contract method 0xf4e58618.
//
// Solidity: function auctionByAssetId(address , uint256 , bytes32 ) view returns(bool ended, bytes32 id, address seller, address highestBidder, address nftAddress, uint256 biddingEnd, uint256 revealEnd, uint256 highestBid, uint256 startPrice)
func (_Smc *SmcSession) AuctionByAssetId(arg0 common.Address, arg1 *big.Int, arg2 [32]byte) (struct {
	Ended         bool
	Id            [32]byte
	Seller        common.Address
	HighestBidder common.Address
	NftAddress    common.Address
	BiddingEnd    *big.Int
	RevealEnd     *big.Int
	HighestBid    *big.Int
	StartPrice    *big.Int
}, error) {
	return _Smc.Contract.AuctionByAssetId(&_Smc.CallOpts, arg0, arg1, arg2)
}

// AuctionByAssetId is a free data retrieval call binding the contract method 0xf4e58618.
//
// Solidity: function auctionByAssetId(address , uint256 , bytes32 ) view returns(bool ended, bytes32 id, address seller, address highestBidder, address nftAddress, uint256 biddingEnd, uint256 revealEnd, uint256 highestBid, uint256 startPrice)
func (_Smc *SmcCallerSession) AuctionByAssetId(arg0 common.Address, arg1 *big.Int, arg2 [32]byte) (struct {
	Ended         bool
	Id            [32]byte
	Seller        common.Address
	HighestBidder common.Address
	NftAddress    common.Address
	BiddingEnd    *big.Int
	RevealEnd     *big.Int
	HighestBid    *big.Int
	StartPrice    *big.Int
}, error) {
	return _Smc.Contract.AuctionByAssetId(&_Smc.CallOpts, arg0, arg1, arg2)
}

// SmcAuctionCancelledIterator is returned from FilterAuctionCancelled and is used to iterate over the raw logs and unpacked data for AuctionCancelled events raised by the Smc contract.
type SmcAuctionCancelledIterator struct {
	Event *SmcAuctionCancelled // Event containing the contract specifics and raw log

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
func (it *SmcAuctionCancelledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SmcAuctionCancelled)
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
		it.Event = new(SmcAuctionCancelled)
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
func (it *SmcAuctionCancelledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SmcAuctionCancelledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SmcAuctionCancelled represents a AuctionCancelled event raised by the Smc contract.
type SmcAuctionCancelled struct {
	Id         [32]byte
	AssetId    *big.Int
	Seller     common.Address
	NftAddress common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterAuctionCancelled is a free log retrieval operation binding the contract event 0xc8d42db0d3cbb28bf5c6069f43e43364babcf8b257b1054e8b35bb2f46f0ab1e.
//
// Solidity: event AuctionCancelled(bytes32 id, uint256 indexed assetId, address indexed seller, address nftAddress)
func (_Smc *SmcFilterer) FilterAuctionCancelled(opts *bind.FilterOpts, assetId []*big.Int, seller []common.Address) (*SmcAuctionCancelledIterator, error) {

	var assetIdRule []interface{}
	for _, assetIdItem := range assetId {
		assetIdRule = append(assetIdRule, assetIdItem)
	}
	var sellerRule []interface{}
	for _, sellerItem := range seller {
		sellerRule = append(sellerRule, sellerItem)
	}

	logs, sub, err := _Smc.contract.FilterLogs(opts, "AuctionCancelled", assetIdRule, sellerRule)
	if err != nil {
		return nil, err
	}
	return &SmcAuctionCancelledIterator{contract: _Smc.contract, event: "AuctionCancelled", logs: logs, sub: sub}, nil
}

// WatchAuctionCancelled is a free log subscription operation binding the contract event 0xc8d42db0d3cbb28bf5c6069f43e43364babcf8b257b1054e8b35bb2f46f0ab1e.
//
// Solidity: event AuctionCancelled(bytes32 id, uint256 indexed assetId, address indexed seller, address nftAddress)
func (_Smc *SmcFilterer) WatchAuctionCancelled(opts *bind.WatchOpts, sink chan<- *SmcAuctionCancelled, assetId []*big.Int, seller []common.Address) (event.Subscription, error) {

	var assetIdRule []interface{}
	for _, assetIdItem := range assetId {
		assetIdRule = append(assetIdRule, assetIdItem)
	}
	var sellerRule []interface{}
	for _, sellerItem := range seller {
		sellerRule = append(sellerRule, sellerItem)
	}

	logs, sub, err := _Smc.contract.WatchLogs(opts, "AuctionCancelled", assetIdRule, sellerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SmcAuctionCancelled)
				if err := _Smc.contract.UnpackLog(event, "AuctionCancelled", log); err != nil {
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

// ParseAuctionCancelled is a log parse operation binding the contract event 0xc8d42db0d3cbb28bf5c6069f43e43364babcf8b257b1054e8b35bb2f46f0ab1e.
//
// Solidity: event AuctionCancelled(bytes32 id, uint256 indexed assetId, address indexed seller, address nftAddress)
func (_Smc *SmcFilterer) ParseAuctionCancelled(log types.Log) (*SmcAuctionCancelled, error) {
	event := new(SmcAuctionCancelled)
	if err := _Smc.contract.UnpackLog(event, "AuctionCancelled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SmcAuctionCreatedIterator is returned from FilterAuctionCreated and is used to iterate over the raw logs and unpacked data for AuctionCreated events raised by the Smc contract.
type SmcAuctionCreatedIterator struct {
	Event *SmcAuctionCreated // Event containing the contract specifics and raw log

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
func (it *SmcAuctionCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SmcAuctionCreated)
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
		it.Event = new(SmcAuctionCreated)
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
func (it *SmcAuctionCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SmcAuctionCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SmcAuctionCreated represents a AuctionCreated event raised by the Smc contract.
type SmcAuctionCreated struct {
	Id              [32]byte
	AssetId         *big.Int
	Seller          common.Address
	NftAddress      common.Address
	BiddingEnd      *big.Int
	RevealEnd       *big.Int
	StartPriceInWei *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterAuctionCreated is a free log retrieval operation binding the contract event 0xba862ff741f4b79389e946bd999bc64728b41fc3dc7ab36af5696ecc2e7ac292.
//
// Solidity: event AuctionCreated(bytes32 id, uint256 indexed assetId, address indexed seller, address nftAddress, uint256 biddingEnd, uint256 revealEnd, uint256 startPriceInWei)
func (_Smc *SmcFilterer) FilterAuctionCreated(opts *bind.FilterOpts, assetId []*big.Int, seller []common.Address) (*SmcAuctionCreatedIterator, error) {

	var assetIdRule []interface{}
	for _, assetIdItem := range assetId {
		assetIdRule = append(assetIdRule, assetIdItem)
	}
	var sellerRule []interface{}
	for _, sellerItem := range seller {
		sellerRule = append(sellerRule, sellerItem)
	}

	logs, sub, err := _Smc.contract.FilterLogs(opts, "AuctionCreated", assetIdRule, sellerRule)
	if err != nil {
		return nil, err
	}
	return &SmcAuctionCreatedIterator{contract: _Smc.contract, event: "AuctionCreated", logs: logs, sub: sub}, nil
}

// WatchAuctionCreated is a free log subscription operation binding the contract event 0xba862ff741f4b79389e946bd999bc64728b41fc3dc7ab36af5696ecc2e7ac292.
//
// Solidity: event AuctionCreated(bytes32 id, uint256 indexed assetId, address indexed seller, address nftAddress, uint256 biddingEnd, uint256 revealEnd, uint256 startPriceInWei)
func (_Smc *SmcFilterer) WatchAuctionCreated(opts *bind.WatchOpts, sink chan<- *SmcAuctionCreated, assetId []*big.Int, seller []common.Address) (event.Subscription, error) {

	var assetIdRule []interface{}
	for _, assetIdItem := range assetId {
		assetIdRule = append(assetIdRule, assetIdItem)
	}
	var sellerRule []interface{}
	for _, sellerItem := range seller {
		sellerRule = append(sellerRule, sellerItem)
	}

	logs, sub, err := _Smc.contract.WatchLogs(opts, "AuctionCreated", assetIdRule, sellerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SmcAuctionCreated)
				if err := _Smc.contract.UnpackLog(event, "AuctionCreated", log); err != nil {
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

// ParseAuctionCreated is a log parse operation binding the contract event 0xba862ff741f4b79389e946bd999bc64728b41fc3dc7ab36af5696ecc2e7ac292.
//
// Solidity: event AuctionCreated(bytes32 id, uint256 indexed assetId, address indexed seller, address nftAddress, uint256 biddingEnd, uint256 revealEnd, uint256 startPriceInWei)
func (_Smc *SmcFilterer) ParseAuctionCreated(log types.Log) (*SmcAuctionCreated, error) {
	event := new(SmcAuctionCreated)
	if err := _Smc.contract.UnpackLog(event, "AuctionCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
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

// SmcAuctionSuccessfulIterator is returned from FilterAuctionSuccessful and is used to iterate over the raw logs and unpacked data for AuctionSuccessful events raised by the Smc contract.
type SmcAuctionSuccessfulIterator struct {
	Event *SmcAuctionSuccessful // Event containing the contract specifics and raw log

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
func (it *SmcAuctionSuccessfulIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SmcAuctionSuccessful)
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
		it.Event = new(SmcAuctionSuccessful)
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
func (it *SmcAuctionSuccessfulIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SmcAuctionSuccessfulIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SmcAuctionSuccessful represents a AuctionSuccessful event raised by the Smc contract.
type SmcAuctionSuccessful struct {
	Id         [32]byte
	AssetId    *big.Int
	Seller     common.Address
	NftAddress common.Address
	TotalPrice *big.Int
	Buyer      common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterAuctionSuccessful is a free log retrieval operation binding the contract event 0x5dc66c4eeb72a3000f560dd1dcdf051cccf7c0baba37a3d77c844afd7a9d26f1.
//
// Solidity: event AuctionSuccessful(bytes32 id, uint256 indexed assetId, address indexed seller, address nftAddress, uint256 totalPrice, address indexed buyer)
func (_Smc *SmcFilterer) FilterAuctionSuccessful(opts *bind.FilterOpts, assetId []*big.Int, seller []common.Address, buyer []common.Address) (*SmcAuctionSuccessfulIterator, error) {

	var assetIdRule []interface{}
	for _, assetIdItem := range assetId {
		assetIdRule = append(assetIdRule, assetIdItem)
	}
	var sellerRule []interface{}
	for _, sellerItem := range seller {
		sellerRule = append(sellerRule, sellerItem)
	}

	var buyerRule []interface{}
	for _, buyerItem := range buyer {
		buyerRule = append(buyerRule, buyerItem)
	}

	logs, sub, err := _Smc.contract.FilterLogs(opts, "AuctionSuccessful", assetIdRule, sellerRule, buyerRule)
	if err != nil {
		return nil, err
	}
	return &SmcAuctionSuccessfulIterator{contract: _Smc.contract, event: "AuctionSuccessful", logs: logs, sub: sub}, nil
}

// WatchAuctionSuccessful is a free log subscription operation binding the contract event 0x5dc66c4eeb72a3000f560dd1dcdf051cccf7c0baba37a3d77c844afd7a9d26f1.
//
// Solidity: event AuctionSuccessful(bytes32 id, uint256 indexed assetId, address indexed seller, address nftAddress, uint256 totalPrice, address indexed buyer)
func (_Smc *SmcFilterer) WatchAuctionSuccessful(opts *bind.WatchOpts, sink chan<- *SmcAuctionSuccessful, assetId []*big.Int, seller []common.Address, buyer []common.Address) (event.Subscription, error) {

	var assetIdRule []interface{}
	for _, assetIdItem := range assetId {
		assetIdRule = append(assetIdRule, assetIdItem)
	}
	var sellerRule []interface{}
	for _, sellerItem := range seller {
		sellerRule = append(sellerRule, sellerItem)
	}

	var buyerRule []interface{}
	for _, buyerItem := range buyer {
		buyerRule = append(buyerRule, buyerItem)
	}

	logs, sub, err := _Smc.contract.WatchLogs(opts, "AuctionSuccessful", assetIdRule, sellerRule, buyerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SmcAuctionSuccessful)
				if err := _Smc.contract.UnpackLog(event, "AuctionSuccessful", log); err != nil {
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

// ParseAuctionSuccessful is a log parse operation binding the contract event 0x5dc66c4eeb72a3000f560dd1dcdf051cccf7c0baba37a3d77c844afd7a9d26f1.
//
// Solidity: event AuctionSuccessful(bytes32 id, uint256 indexed assetId, address indexed seller, address nftAddress, uint256 totalPrice, address indexed buyer)
func (_Smc *SmcFilterer) ParseAuctionSuccessful(log types.Log) (*SmcAuctionSuccessful, error) {
	event := new(SmcAuctionSuccessful)
	if err := _Smc.contract.UnpackLog(event, "AuctionSuccessful", log); err != nil {
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
