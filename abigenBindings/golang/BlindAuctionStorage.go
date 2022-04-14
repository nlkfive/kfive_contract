// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package BlindAuctionStorage

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

// BlindAuctionStorageMetaData contains all meta data concerning the BlindAuctionStorage contract.
var BlindAuctionStorageMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"AuctionEndAlreadyCalled\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"}],\"name\":\"TooEarly\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"}],\"name\":\"TooLate\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"assetId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"nftAddress\",\"type\":\"address\"}],\"name\":\"AuctionCancelled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"assetId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"nftAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"biddingEnd\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"revealEnd\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"startPriceInWei\",\"type\":\"uint256\"}],\"name\":\"AuctionCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"nftAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"assetId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"deposit\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"bidder\",\"type\":\"address\"}],\"name\":\"AuctionRefund\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"assetId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"nftAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalPrice\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"buyer\",\"type\":\"address\"}],\"name\":\"AuctionSuccessful\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"bidder\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"nftAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"assetId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"blindedBid\",\"type\":\"bytes32\"}],\"name\":\"BidSuccessful\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"auctionByAssetId\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"ended\",\"type\":\"bool\"},{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"highestBidder\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"nftAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"biddingEnd\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"revealEnd\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"highestBid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startPrice\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// BlindAuctionStorageABI is the input ABI used to generate the binding from.
// Deprecated: Use BlindAuctionStorageMetaData.ABI instead.
var BlindAuctionStorageABI = BlindAuctionStorageMetaData.ABI

// BlindAuctionStorage is an auto generated Go binding around an Ethereum contract.
type BlindAuctionStorage struct {
	BlindAuctionStorageCaller     // Read-only binding to the contract
	BlindAuctionStorageTransactor // Write-only binding to the contract
	BlindAuctionStorageFilterer   // Log filterer for contract events
}

// BlindAuctionStorageCaller is an auto generated read-only Go binding around an Ethereum contract.
type BlindAuctionStorageCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BlindAuctionStorageTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BlindAuctionStorageTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BlindAuctionStorageFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BlindAuctionStorageFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BlindAuctionStorageSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BlindAuctionStorageSession struct {
	Contract     *BlindAuctionStorage // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// BlindAuctionStorageCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BlindAuctionStorageCallerSession struct {
	Contract *BlindAuctionStorageCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// BlindAuctionStorageTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BlindAuctionStorageTransactorSession struct {
	Contract     *BlindAuctionStorageTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// BlindAuctionStorageRaw is an auto generated low-level Go binding around an Ethereum contract.
type BlindAuctionStorageRaw struct {
	Contract *BlindAuctionStorage // Generic contract binding to access the raw methods on
}

// BlindAuctionStorageCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BlindAuctionStorageCallerRaw struct {
	Contract *BlindAuctionStorageCaller // Generic read-only contract binding to access the raw methods on
}

// BlindAuctionStorageTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BlindAuctionStorageTransactorRaw struct {
	Contract *BlindAuctionStorageTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBlindAuctionStorage creates a new instance of BlindAuctionStorage, bound to a specific deployed contract.
func NewBlindAuctionStorage(address common.Address, backend bind.ContractBackend) (*BlindAuctionStorage, error) {
	contract, err := bindBlindAuctionStorage(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BlindAuctionStorage{BlindAuctionStorageCaller: BlindAuctionStorageCaller{contract: contract}, BlindAuctionStorageTransactor: BlindAuctionStorageTransactor{contract: contract}, BlindAuctionStorageFilterer: BlindAuctionStorageFilterer{contract: contract}}, nil
}

// NewBlindAuctionStorageCaller creates a new read-only instance of BlindAuctionStorage, bound to a specific deployed contract.
func NewBlindAuctionStorageCaller(address common.Address, caller bind.ContractCaller) (*BlindAuctionStorageCaller, error) {
	contract, err := bindBlindAuctionStorage(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BlindAuctionStorageCaller{contract: contract}, nil
}

// NewBlindAuctionStorageTransactor creates a new write-only instance of BlindAuctionStorage, bound to a specific deployed contract.
func NewBlindAuctionStorageTransactor(address common.Address, transactor bind.ContractTransactor) (*BlindAuctionStorageTransactor, error) {
	contract, err := bindBlindAuctionStorage(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BlindAuctionStorageTransactor{contract: contract}, nil
}

// NewBlindAuctionStorageFilterer creates a new log filterer instance of BlindAuctionStorage, bound to a specific deployed contract.
func NewBlindAuctionStorageFilterer(address common.Address, filterer bind.ContractFilterer) (*BlindAuctionStorageFilterer, error) {
	contract, err := bindBlindAuctionStorage(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BlindAuctionStorageFilterer{contract: contract}, nil
}

// bindBlindAuctionStorage binds a generic wrapper to an already deployed contract.
func bindBlindAuctionStorage(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BlindAuctionStorageABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BlindAuctionStorage *BlindAuctionStorageRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BlindAuctionStorage.Contract.BlindAuctionStorageCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BlindAuctionStorage *BlindAuctionStorageRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BlindAuctionStorage.Contract.BlindAuctionStorageTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BlindAuctionStorage *BlindAuctionStorageRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BlindAuctionStorage.Contract.BlindAuctionStorageTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BlindAuctionStorage *BlindAuctionStorageCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BlindAuctionStorage.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BlindAuctionStorage *BlindAuctionStorageTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BlindAuctionStorage.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BlindAuctionStorage *BlindAuctionStorageTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BlindAuctionStorage.Contract.contract.Transact(opts, method, params...)
}

// AuctionByAssetId is a free data retrieval call binding the contract method 0xf4e58618.
//
// Solidity: function auctionByAssetId(address , uint256 , bytes32 ) view returns(bool ended, bytes32 id, address seller, address highestBidder, address nftAddress, uint256 biddingEnd, uint256 revealEnd, uint256 highestBid, uint256 startPrice)
func (_BlindAuctionStorage *BlindAuctionStorageCaller) AuctionByAssetId(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int, arg2 [32]byte) (struct {
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
	err := _BlindAuctionStorage.contract.Call(opts, &out, "auctionByAssetId", arg0, arg1, arg2)

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
func (_BlindAuctionStorage *BlindAuctionStorageSession) AuctionByAssetId(arg0 common.Address, arg1 *big.Int, arg2 [32]byte) (struct {
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
	return _BlindAuctionStorage.Contract.AuctionByAssetId(&_BlindAuctionStorage.CallOpts, arg0, arg1, arg2)
}

// AuctionByAssetId is a free data retrieval call binding the contract method 0xf4e58618.
//
// Solidity: function auctionByAssetId(address , uint256 , bytes32 ) view returns(bool ended, bytes32 id, address seller, address highestBidder, address nftAddress, uint256 biddingEnd, uint256 revealEnd, uint256 highestBid, uint256 startPrice)
func (_BlindAuctionStorage *BlindAuctionStorageCallerSession) AuctionByAssetId(arg0 common.Address, arg1 *big.Int, arg2 [32]byte) (struct {
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
	return _BlindAuctionStorage.Contract.AuctionByAssetId(&_BlindAuctionStorage.CallOpts, arg0, arg1, arg2)
}

// BlindAuctionStorageAuctionCancelledIterator is returned from FilterAuctionCancelled and is used to iterate over the raw logs and unpacked data for AuctionCancelled events raised by the BlindAuctionStorage contract.
type BlindAuctionStorageAuctionCancelledIterator struct {
	Event *BlindAuctionStorageAuctionCancelled // Event containing the contract specifics and raw log

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
func (it *BlindAuctionStorageAuctionCancelledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BlindAuctionStorageAuctionCancelled)
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
		it.Event = new(BlindAuctionStorageAuctionCancelled)
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
func (it *BlindAuctionStorageAuctionCancelledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BlindAuctionStorageAuctionCancelledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BlindAuctionStorageAuctionCancelled represents a AuctionCancelled event raised by the BlindAuctionStorage contract.
type BlindAuctionStorageAuctionCancelled struct {
	Id         [32]byte
	AssetId    *big.Int
	Seller     common.Address
	NftAddress common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterAuctionCancelled is a free log retrieval operation binding the contract event 0xc8d42db0d3cbb28bf5c6069f43e43364babcf8b257b1054e8b35bb2f46f0ab1e.
//
// Solidity: event AuctionCancelled(bytes32 id, uint256 indexed assetId, address indexed seller, address nftAddress)
func (_BlindAuctionStorage *BlindAuctionStorageFilterer) FilterAuctionCancelled(opts *bind.FilterOpts, assetId []*big.Int, seller []common.Address) (*BlindAuctionStorageAuctionCancelledIterator, error) {

	var assetIdRule []interface{}
	for _, assetIdItem := range assetId {
		assetIdRule = append(assetIdRule, assetIdItem)
	}
	var sellerRule []interface{}
	for _, sellerItem := range seller {
		sellerRule = append(sellerRule, sellerItem)
	}

	logs, sub, err := _BlindAuctionStorage.contract.FilterLogs(opts, "AuctionCancelled", assetIdRule, sellerRule)
	if err != nil {
		return nil, err
	}
	return &BlindAuctionStorageAuctionCancelledIterator{contract: _BlindAuctionStorage.contract, event: "AuctionCancelled", logs: logs, sub: sub}, nil
}

// WatchAuctionCancelled is a free log subscription operation binding the contract event 0xc8d42db0d3cbb28bf5c6069f43e43364babcf8b257b1054e8b35bb2f46f0ab1e.
//
// Solidity: event AuctionCancelled(bytes32 id, uint256 indexed assetId, address indexed seller, address nftAddress)
func (_BlindAuctionStorage *BlindAuctionStorageFilterer) WatchAuctionCancelled(opts *bind.WatchOpts, sink chan<- *BlindAuctionStorageAuctionCancelled, assetId []*big.Int, seller []common.Address) (event.Subscription, error) {

	var assetIdRule []interface{}
	for _, assetIdItem := range assetId {
		assetIdRule = append(assetIdRule, assetIdItem)
	}
	var sellerRule []interface{}
	for _, sellerItem := range seller {
		sellerRule = append(sellerRule, sellerItem)
	}

	logs, sub, err := _BlindAuctionStorage.contract.WatchLogs(opts, "AuctionCancelled", assetIdRule, sellerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BlindAuctionStorageAuctionCancelled)
				if err := _BlindAuctionStorage.contract.UnpackLog(event, "AuctionCancelled", log); err != nil {
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
func (_BlindAuctionStorage *BlindAuctionStorageFilterer) ParseAuctionCancelled(log types.Log) (*BlindAuctionStorageAuctionCancelled, error) {
	event := new(BlindAuctionStorageAuctionCancelled)
	if err := _BlindAuctionStorage.contract.UnpackLog(event, "AuctionCancelled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BlindAuctionStorageAuctionCreatedIterator is returned from FilterAuctionCreated and is used to iterate over the raw logs and unpacked data for AuctionCreated events raised by the BlindAuctionStorage contract.
type BlindAuctionStorageAuctionCreatedIterator struct {
	Event *BlindAuctionStorageAuctionCreated // Event containing the contract specifics and raw log

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
func (it *BlindAuctionStorageAuctionCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BlindAuctionStorageAuctionCreated)
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
		it.Event = new(BlindAuctionStorageAuctionCreated)
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
func (it *BlindAuctionStorageAuctionCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BlindAuctionStorageAuctionCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BlindAuctionStorageAuctionCreated represents a AuctionCreated event raised by the BlindAuctionStorage contract.
type BlindAuctionStorageAuctionCreated struct {
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
func (_BlindAuctionStorage *BlindAuctionStorageFilterer) FilterAuctionCreated(opts *bind.FilterOpts, assetId []*big.Int, seller []common.Address) (*BlindAuctionStorageAuctionCreatedIterator, error) {

	var assetIdRule []interface{}
	for _, assetIdItem := range assetId {
		assetIdRule = append(assetIdRule, assetIdItem)
	}
	var sellerRule []interface{}
	for _, sellerItem := range seller {
		sellerRule = append(sellerRule, sellerItem)
	}

	logs, sub, err := _BlindAuctionStorage.contract.FilterLogs(opts, "AuctionCreated", assetIdRule, sellerRule)
	if err != nil {
		return nil, err
	}
	return &BlindAuctionStorageAuctionCreatedIterator{contract: _BlindAuctionStorage.contract, event: "AuctionCreated", logs: logs, sub: sub}, nil
}

// WatchAuctionCreated is a free log subscription operation binding the contract event 0xba862ff741f4b79389e946bd999bc64728b41fc3dc7ab36af5696ecc2e7ac292.
//
// Solidity: event AuctionCreated(bytes32 id, uint256 indexed assetId, address indexed seller, address nftAddress, uint256 biddingEnd, uint256 revealEnd, uint256 startPriceInWei)
func (_BlindAuctionStorage *BlindAuctionStorageFilterer) WatchAuctionCreated(opts *bind.WatchOpts, sink chan<- *BlindAuctionStorageAuctionCreated, assetId []*big.Int, seller []common.Address) (event.Subscription, error) {

	var assetIdRule []interface{}
	for _, assetIdItem := range assetId {
		assetIdRule = append(assetIdRule, assetIdItem)
	}
	var sellerRule []interface{}
	for _, sellerItem := range seller {
		sellerRule = append(sellerRule, sellerItem)
	}

	logs, sub, err := _BlindAuctionStorage.contract.WatchLogs(opts, "AuctionCreated", assetIdRule, sellerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BlindAuctionStorageAuctionCreated)
				if err := _BlindAuctionStorage.contract.UnpackLog(event, "AuctionCreated", log); err != nil {
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
func (_BlindAuctionStorage *BlindAuctionStorageFilterer) ParseAuctionCreated(log types.Log) (*BlindAuctionStorageAuctionCreated, error) {
	event := new(BlindAuctionStorageAuctionCreated)
	if err := _BlindAuctionStorage.contract.UnpackLog(event, "AuctionCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BlindAuctionStorageAuctionRefundIterator is returned from FilterAuctionRefund and is used to iterate over the raw logs and unpacked data for AuctionRefund events raised by the BlindAuctionStorage contract.
type BlindAuctionStorageAuctionRefundIterator struct {
	Event *BlindAuctionStorageAuctionRefund // Event containing the contract specifics and raw log

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
func (it *BlindAuctionStorageAuctionRefundIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BlindAuctionStorageAuctionRefund)
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
		it.Event = new(BlindAuctionStorageAuctionRefund)
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
func (it *BlindAuctionStorageAuctionRefundIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BlindAuctionStorageAuctionRefundIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BlindAuctionStorageAuctionRefund represents a AuctionRefund event raised by the BlindAuctionStorage contract.
type BlindAuctionStorageAuctionRefund struct {
	NftAddress common.Address
	AssetId    *big.Int
	Deposit    *big.Int
	Bidder     common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterAuctionRefund is a free log retrieval operation binding the contract event 0xe71ce726e7a1cceb31d105962067824bc8ebd189ffaa94839b1ac96e6a8d38af.
//
// Solidity: event AuctionRefund(address nftAddress, uint256 indexed assetId, uint256 deposit, address indexed bidder)
func (_BlindAuctionStorage *BlindAuctionStorageFilterer) FilterAuctionRefund(opts *bind.FilterOpts, assetId []*big.Int, bidder []common.Address) (*BlindAuctionStorageAuctionRefundIterator, error) {

	var assetIdRule []interface{}
	for _, assetIdItem := range assetId {
		assetIdRule = append(assetIdRule, assetIdItem)
	}

	var bidderRule []interface{}
	for _, bidderItem := range bidder {
		bidderRule = append(bidderRule, bidderItem)
	}

	logs, sub, err := _BlindAuctionStorage.contract.FilterLogs(opts, "AuctionRefund", assetIdRule, bidderRule)
	if err != nil {
		return nil, err
	}
	return &BlindAuctionStorageAuctionRefundIterator{contract: _BlindAuctionStorage.contract, event: "AuctionRefund", logs: logs, sub: sub}, nil
}

// WatchAuctionRefund is a free log subscription operation binding the contract event 0xe71ce726e7a1cceb31d105962067824bc8ebd189ffaa94839b1ac96e6a8d38af.
//
// Solidity: event AuctionRefund(address nftAddress, uint256 indexed assetId, uint256 deposit, address indexed bidder)
func (_BlindAuctionStorage *BlindAuctionStorageFilterer) WatchAuctionRefund(opts *bind.WatchOpts, sink chan<- *BlindAuctionStorageAuctionRefund, assetId []*big.Int, bidder []common.Address) (event.Subscription, error) {

	var assetIdRule []interface{}
	for _, assetIdItem := range assetId {
		assetIdRule = append(assetIdRule, assetIdItem)
	}

	var bidderRule []interface{}
	for _, bidderItem := range bidder {
		bidderRule = append(bidderRule, bidderItem)
	}

	logs, sub, err := _BlindAuctionStorage.contract.WatchLogs(opts, "AuctionRefund", assetIdRule, bidderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BlindAuctionStorageAuctionRefund)
				if err := _BlindAuctionStorage.contract.UnpackLog(event, "AuctionRefund", log); err != nil {
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
func (_BlindAuctionStorage *BlindAuctionStorageFilterer) ParseAuctionRefund(log types.Log) (*BlindAuctionStorageAuctionRefund, error) {
	event := new(BlindAuctionStorageAuctionRefund)
	if err := _BlindAuctionStorage.contract.UnpackLog(event, "AuctionRefund", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BlindAuctionStorageAuctionSuccessfulIterator is returned from FilterAuctionSuccessful and is used to iterate over the raw logs and unpacked data for AuctionSuccessful events raised by the BlindAuctionStorage contract.
type BlindAuctionStorageAuctionSuccessfulIterator struct {
	Event *BlindAuctionStorageAuctionSuccessful // Event containing the contract specifics and raw log

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
func (it *BlindAuctionStorageAuctionSuccessfulIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BlindAuctionStorageAuctionSuccessful)
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
		it.Event = new(BlindAuctionStorageAuctionSuccessful)
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
func (it *BlindAuctionStorageAuctionSuccessfulIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BlindAuctionStorageAuctionSuccessfulIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BlindAuctionStorageAuctionSuccessful represents a AuctionSuccessful event raised by the BlindAuctionStorage contract.
type BlindAuctionStorageAuctionSuccessful struct {
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
func (_BlindAuctionStorage *BlindAuctionStorageFilterer) FilterAuctionSuccessful(opts *bind.FilterOpts, assetId []*big.Int, seller []common.Address, buyer []common.Address) (*BlindAuctionStorageAuctionSuccessfulIterator, error) {

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

	logs, sub, err := _BlindAuctionStorage.contract.FilterLogs(opts, "AuctionSuccessful", assetIdRule, sellerRule, buyerRule)
	if err != nil {
		return nil, err
	}
	return &BlindAuctionStorageAuctionSuccessfulIterator{contract: _BlindAuctionStorage.contract, event: "AuctionSuccessful", logs: logs, sub: sub}, nil
}

// WatchAuctionSuccessful is a free log subscription operation binding the contract event 0x5dc66c4eeb72a3000f560dd1dcdf051cccf7c0baba37a3d77c844afd7a9d26f1.
//
// Solidity: event AuctionSuccessful(bytes32 id, uint256 indexed assetId, address indexed seller, address nftAddress, uint256 totalPrice, address indexed buyer)
func (_BlindAuctionStorage *BlindAuctionStorageFilterer) WatchAuctionSuccessful(opts *bind.WatchOpts, sink chan<- *BlindAuctionStorageAuctionSuccessful, assetId []*big.Int, seller []common.Address, buyer []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _BlindAuctionStorage.contract.WatchLogs(opts, "AuctionSuccessful", assetIdRule, sellerRule, buyerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BlindAuctionStorageAuctionSuccessful)
				if err := _BlindAuctionStorage.contract.UnpackLog(event, "AuctionSuccessful", log); err != nil {
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
func (_BlindAuctionStorage *BlindAuctionStorageFilterer) ParseAuctionSuccessful(log types.Log) (*BlindAuctionStorageAuctionSuccessful, error) {
	event := new(BlindAuctionStorageAuctionSuccessful)
	if err := _BlindAuctionStorage.contract.UnpackLog(event, "AuctionSuccessful", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BlindAuctionStorageBidSuccessfulIterator is returned from FilterBidSuccessful and is used to iterate over the raw logs and unpacked data for BidSuccessful events raised by the BlindAuctionStorage contract.
type BlindAuctionStorageBidSuccessfulIterator struct {
	Event *BlindAuctionStorageBidSuccessful // Event containing the contract specifics and raw log

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
func (it *BlindAuctionStorageBidSuccessfulIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BlindAuctionStorageBidSuccessful)
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
		it.Event = new(BlindAuctionStorageBidSuccessful)
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
func (it *BlindAuctionStorageBidSuccessfulIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BlindAuctionStorageBidSuccessfulIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BlindAuctionStorageBidSuccessful represents a BidSuccessful event raised by the BlindAuctionStorage contract.
type BlindAuctionStorageBidSuccessful struct {
	Bidder     common.Address
	NftAddress common.Address
	AssetId    *big.Int
	BlindedBid [32]byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterBidSuccessful is a free log retrieval operation binding the contract event 0xc9b646ec968ab42359a60ba55be704c9d9c49f1db98b242275274b18763eefda.
//
// Solidity: event BidSuccessful(address indexed bidder, address nftAddress, uint256 indexed assetId, bytes32 blindedBid)
func (_BlindAuctionStorage *BlindAuctionStorageFilterer) FilterBidSuccessful(opts *bind.FilterOpts, bidder []common.Address, assetId []*big.Int) (*BlindAuctionStorageBidSuccessfulIterator, error) {

	var bidderRule []interface{}
	for _, bidderItem := range bidder {
		bidderRule = append(bidderRule, bidderItem)
	}

	var assetIdRule []interface{}
	for _, assetIdItem := range assetId {
		assetIdRule = append(assetIdRule, assetIdItem)
	}

	logs, sub, err := _BlindAuctionStorage.contract.FilterLogs(opts, "BidSuccessful", bidderRule, assetIdRule)
	if err != nil {
		return nil, err
	}
	return &BlindAuctionStorageBidSuccessfulIterator{contract: _BlindAuctionStorage.contract, event: "BidSuccessful", logs: logs, sub: sub}, nil
}

// WatchBidSuccessful is a free log subscription operation binding the contract event 0xc9b646ec968ab42359a60ba55be704c9d9c49f1db98b242275274b18763eefda.
//
// Solidity: event BidSuccessful(address indexed bidder, address nftAddress, uint256 indexed assetId, bytes32 blindedBid)
func (_BlindAuctionStorage *BlindAuctionStorageFilterer) WatchBidSuccessful(opts *bind.WatchOpts, sink chan<- *BlindAuctionStorageBidSuccessful, bidder []common.Address, assetId []*big.Int) (event.Subscription, error) {

	var bidderRule []interface{}
	for _, bidderItem := range bidder {
		bidderRule = append(bidderRule, bidderItem)
	}

	var assetIdRule []interface{}
	for _, assetIdItem := range assetId {
		assetIdRule = append(assetIdRule, assetIdItem)
	}

	logs, sub, err := _BlindAuctionStorage.contract.WatchLogs(opts, "BidSuccessful", bidderRule, assetIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BlindAuctionStorageBidSuccessful)
				if err := _BlindAuctionStorage.contract.UnpackLog(event, "BidSuccessful", log); err != nil {
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
func (_BlindAuctionStorage *BlindAuctionStorageFilterer) ParseBidSuccessful(log types.Log) (*BlindAuctionStorageBidSuccessful, error) {
	event := new(BlindAuctionStorageBidSuccessful)
	if err := _BlindAuctionStorage.contract.UnpackLog(event, "BidSuccessful", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
