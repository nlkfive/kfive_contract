// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package OrderStorage

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

// OrderStorageMetaData contains all meta data concerning the OrderStorage contract.
var OrderStorageMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"assetId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"nftAddress\",\"type\":\"address\"}],\"name\":\"OrderCancelled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"assetId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"nftAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"priceInWei\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"expiresAt\",\"type\":\"uint256\"}],\"name\":\"OrderCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"assetId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"nftAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalPrice\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"buyer\",\"type\":\"address\"}],\"name\":\"OrderSuccessful\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"orderByAssetId\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"nftAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expiresAt\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50610277806100206000396000f3fe608060405234801561001057600080fd5b506004361061002b5760003560e01c8063e61f385114610030575b600080fd5b61004a60048036038101906100459190610111565b610064565b60405161005b95949392919061017a565b60405180910390f35b6000602052816000526040600020602052806000526040600020600091509150508060000154908060010160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060020160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060030154908060040154905085565b6000813590506100f681610213565b92915050565b60008135905061010b8161022a565b92915050565b6000806040838503121561012457600080fd5b6000610132858286016100e7565b9250506020610143858286016100fc565b9150509250929050565b610156816101cd565b82525050565b610165816101df565b82525050565b61017481610209565b82525050565b600060a08201905061018f600083018861015c565b61019c602083018761014d565b6101a9604083018661014d565b6101b6606083018561016b565b6101c3608083018461016b565b9695505050505050565b60006101d8826101e9565b9050919050565b6000819050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b61021c816101cd565b811461022757600080fd5b50565b61023381610209565b811461023e57600080fd5b5056fea2646970667358221220a00f26721a1cae34ce8ddad30ed6fab66e997749a47ccac87bb617ef9eca578f64736f6c63430008040033",
}

// OrderStorageABI is the input ABI used to generate the binding from.
// Deprecated: Use OrderStorageMetaData.ABI instead.
var OrderStorageABI = OrderStorageMetaData.ABI

// OrderStorageBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use OrderStorageMetaData.Bin instead.
var OrderStorageBin = OrderStorageMetaData.Bin

// DeployOrderStorage deploys a new Ethereum contract, binding an instance of OrderStorage to it.
func DeployOrderStorage(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *OrderStorage, error) {
	parsed, err := OrderStorageMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(OrderStorageBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &OrderStorage{OrderStorageCaller: OrderStorageCaller{contract: contract}, OrderStorageTransactor: OrderStorageTransactor{contract: contract}, OrderStorageFilterer: OrderStorageFilterer{contract: contract}}, nil
}

// OrderStorage is an auto generated Go binding around an Ethereum contract.
type OrderStorage struct {
	OrderStorageCaller     // Read-only binding to the contract
	OrderStorageTransactor // Write-only binding to the contract
	OrderStorageFilterer   // Log filterer for contract events
}

// OrderStorageCaller is an auto generated read-only Go binding around an Ethereum contract.
type OrderStorageCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OrderStorageTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OrderStorageTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OrderStorageFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OrderStorageFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OrderStorageSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OrderStorageSession struct {
	Contract     *OrderStorage     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OrderStorageCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OrderStorageCallerSession struct {
	Contract *OrderStorageCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// OrderStorageTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OrderStorageTransactorSession struct {
	Contract     *OrderStorageTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// OrderStorageRaw is an auto generated low-level Go binding around an Ethereum contract.
type OrderStorageRaw struct {
	Contract *OrderStorage // Generic contract binding to access the raw methods on
}

// OrderStorageCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OrderStorageCallerRaw struct {
	Contract *OrderStorageCaller // Generic read-only contract binding to access the raw methods on
}

// OrderStorageTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OrderStorageTransactorRaw struct {
	Contract *OrderStorageTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOrderStorage creates a new instance of OrderStorage, bound to a specific deployed contract.
func NewOrderStorage(address common.Address, backend bind.ContractBackend) (*OrderStorage, error) {
	contract, err := bindOrderStorage(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &OrderStorage{OrderStorageCaller: OrderStorageCaller{contract: contract}, OrderStorageTransactor: OrderStorageTransactor{contract: contract}, OrderStorageFilterer: OrderStorageFilterer{contract: contract}}, nil
}

// NewOrderStorageCaller creates a new read-only instance of OrderStorage, bound to a specific deployed contract.
func NewOrderStorageCaller(address common.Address, caller bind.ContractCaller) (*OrderStorageCaller, error) {
	contract, err := bindOrderStorage(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OrderStorageCaller{contract: contract}, nil
}

// NewOrderStorageTransactor creates a new write-only instance of OrderStorage, bound to a specific deployed contract.
func NewOrderStorageTransactor(address common.Address, transactor bind.ContractTransactor) (*OrderStorageTransactor, error) {
	contract, err := bindOrderStorage(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OrderStorageTransactor{contract: contract}, nil
}

// NewOrderStorageFilterer creates a new log filterer instance of OrderStorage, bound to a specific deployed contract.
func NewOrderStorageFilterer(address common.Address, filterer bind.ContractFilterer) (*OrderStorageFilterer, error) {
	contract, err := bindOrderStorage(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OrderStorageFilterer{contract: contract}, nil
}

// bindOrderStorage binds a generic wrapper to an already deployed contract.
func bindOrderStorage(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(OrderStorageABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OrderStorage *OrderStorageRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OrderStorage.Contract.OrderStorageCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OrderStorage *OrderStorageRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OrderStorage.Contract.OrderStorageTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OrderStorage *OrderStorageRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OrderStorage.Contract.OrderStorageTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OrderStorage *OrderStorageCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OrderStorage.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OrderStorage *OrderStorageTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OrderStorage.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OrderStorage *OrderStorageTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OrderStorage.Contract.contract.Transact(opts, method, params...)
}

// OrderByAssetId is a free data retrieval call binding the contract method 0xe61f3851.
//
// Solidity: function orderByAssetId(address , uint256 ) view returns(bytes32 id, address seller, address nftAddress, uint256 price, uint256 expiresAt)
func (_OrderStorage *OrderStorageCaller) OrderByAssetId(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (struct {
	Id         [32]byte
	Seller     common.Address
	NftAddress common.Address
	Price      *big.Int
	ExpiresAt  *big.Int
}, error) {
	var out []interface{}
	err := _OrderStorage.contract.Call(opts, &out, "orderByAssetId", arg0, arg1)

	outstruct := new(struct {
		Id         [32]byte
		Seller     common.Address
		NftAddress common.Address
		Price      *big.Int
		ExpiresAt  *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Id = *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	outstruct.Seller = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.NftAddress = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)
	outstruct.Price = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.ExpiresAt = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// OrderByAssetId is a free data retrieval call binding the contract method 0xe61f3851.
//
// Solidity: function orderByAssetId(address , uint256 ) view returns(bytes32 id, address seller, address nftAddress, uint256 price, uint256 expiresAt)
func (_OrderStorage *OrderStorageSession) OrderByAssetId(arg0 common.Address, arg1 *big.Int) (struct {
	Id         [32]byte
	Seller     common.Address
	NftAddress common.Address
	Price      *big.Int
	ExpiresAt  *big.Int
}, error) {
	return _OrderStorage.Contract.OrderByAssetId(&_OrderStorage.CallOpts, arg0, arg1)
}

// OrderByAssetId is a free data retrieval call binding the contract method 0xe61f3851.
//
// Solidity: function orderByAssetId(address , uint256 ) view returns(bytes32 id, address seller, address nftAddress, uint256 price, uint256 expiresAt)
func (_OrderStorage *OrderStorageCallerSession) OrderByAssetId(arg0 common.Address, arg1 *big.Int) (struct {
	Id         [32]byte
	Seller     common.Address
	NftAddress common.Address
	Price      *big.Int
	ExpiresAt  *big.Int
}, error) {
	return _OrderStorage.Contract.OrderByAssetId(&_OrderStorage.CallOpts, arg0, arg1)
}

// OrderStorageOrderCancelledIterator is returned from FilterOrderCancelled and is used to iterate over the raw logs and unpacked data for OrderCancelled events raised by the OrderStorage contract.
type OrderStorageOrderCancelledIterator struct {
	Event *OrderStorageOrderCancelled // Event containing the contract specifics and raw log

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
func (it *OrderStorageOrderCancelledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OrderStorageOrderCancelled)
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
		it.Event = new(OrderStorageOrderCancelled)
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
func (it *OrderStorageOrderCancelledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OrderStorageOrderCancelledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OrderStorageOrderCancelled represents a OrderCancelled event raised by the OrderStorage contract.
type OrderStorageOrderCancelled struct {
	Id         [32]byte
	AssetId    *big.Int
	Seller     common.Address
	NftAddress common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterOrderCancelled is a free log retrieval operation binding the contract event 0x0325426328de5b91ae4ad8462ad4076de4bcaf4551e81556185cacde5a425c6b.
//
// Solidity: event OrderCancelled(bytes32 id, uint256 indexed assetId, address indexed seller, address nftAddress)
func (_OrderStorage *OrderStorageFilterer) FilterOrderCancelled(opts *bind.FilterOpts, assetId []*big.Int, seller []common.Address) (*OrderStorageOrderCancelledIterator, error) {

	var assetIdRule []interface{}
	for _, assetIdItem := range assetId {
		assetIdRule = append(assetIdRule, assetIdItem)
	}
	var sellerRule []interface{}
	for _, sellerItem := range seller {
		sellerRule = append(sellerRule, sellerItem)
	}

	logs, sub, err := _OrderStorage.contract.FilterLogs(opts, "OrderCancelled", assetIdRule, sellerRule)
	if err != nil {
		return nil, err
	}
	return &OrderStorageOrderCancelledIterator{contract: _OrderStorage.contract, event: "OrderCancelled", logs: logs, sub: sub}, nil
}

// WatchOrderCancelled is a free log subscription operation binding the contract event 0x0325426328de5b91ae4ad8462ad4076de4bcaf4551e81556185cacde5a425c6b.
//
// Solidity: event OrderCancelled(bytes32 id, uint256 indexed assetId, address indexed seller, address nftAddress)
func (_OrderStorage *OrderStorageFilterer) WatchOrderCancelled(opts *bind.WatchOpts, sink chan<- *OrderStorageOrderCancelled, assetId []*big.Int, seller []common.Address) (event.Subscription, error) {

	var assetIdRule []interface{}
	for _, assetIdItem := range assetId {
		assetIdRule = append(assetIdRule, assetIdItem)
	}
	var sellerRule []interface{}
	for _, sellerItem := range seller {
		sellerRule = append(sellerRule, sellerItem)
	}

	logs, sub, err := _OrderStorage.contract.WatchLogs(opts, "OrderCancelled", assetIdRule, sellerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OrderStorageOrderCancelled)
				if err := _OrderStorage.contract.UnpackLog(event, "OrderCancelled", log); err != nil {
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

// ParseOrderCancelled is a log parse operation binding the contract event 0x0325426328de5b91ae4ad8462ad4076de4bcaf4551e81556185cacde5a425c6b.
//
// Solidity: event OrderCancelled(bytes32 id, uint256 indexed assetId, address indexed seller, address nftAddress)
func (_OrderStorage *OrderStorageFilterer) ParseOrderCancelled(log types.Log) (*OrderStorageOrderCancelled, error) {
	event := new(OrderStorageOrderCancelled)
	if err := _OrderStorage.contract.UnpackLog(event, "OrderCancelled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OrderStorageOrderCreatedIterator is returned from FilterOrderCreated and is used to iterate over the raw logs and unpacked data for OrderCreated events raised by the OrderStorage contract.
type OrderStorageOrderCreatedIterator struct {
	Event *OrderStorageOrderCreated // Event containing the contract specifics and raw log

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
func (it *OrderStorageOrderCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OrderStorageOrderCreated)
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
		it.Event = new(OrderStorageOrderCreated)
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
func (it *OrderStorageOrderCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OrderStorageOrderCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OrderStorageOrderCreated represents a OrderCreated event raised by the OrderStorage contract.
type OrderStorageOrderCreated struct {
	Id         [32]byte
	AssetId    *big.Int
	Seller     common.Address
	NftAddress common.Address
	PriceInWei *big.Int
	ExpiresAt  *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterOrderCreated is a free log retrieval operation binding the contract event 0x84c66c3f7ba4b390e20e8e8233e2a516f3ce34a72749e4f12bd010dfba238039.
//
// Solidity: event OrderCreated(bytes32 id, uint256 indexed assetId, address indexed seller, address nftAddress, uint256 priceInWei, uint256 expiresAt)
func (_OrderStorage *OrderStorageFilterer) FilterOrderCreated(opts *bind.FilterOpts, assetId []*big.Int, seller []common.Address) (*OrderStorageOrderCreatedIterator, error) {

	var assetIdRule []interface{}
	for _, assetIdItem := range assetId {
		assetIdRule = append(assetIdRule, assetIdItem)
	}
	var sellerRule []interface{}
	for _, sellerItem := range seller {
		sellerRule = append(sellerRule, sellerItem)
	}

	logs, sub, err := _OrderStorage.contract.FilterLogs(opts, "OrderCreated", assetIdRule, sellerRule)
	if err != nil {
		return nil, err
	}
	return &OrderStorageOrderCreatedIterator{contract: _OrderStorage.contract, event: "OrderCreated", logs: logs, sub: sub}, nil
}

// WatchOrderCreated is a free log subscription operation binding the contract event 0x84c66c3f7ba4b390e20e8e8233e2a516f3ce34a72749e4f12bd010dfba238039.
//
// Solidity: event OrderCreated(bytes32 id, uint256 indexed assetId, address indexed seller, address nftAddress, uint256 priceInWei, uint256 expiresAt)
func (_OrderStorage *OrderStorageFilterer) WatchOrderCreated(opts *bind.WatchOpts, sink chan<- *OrderStorageOrderCreated, assetId []*big.Int, seller []common.Address) (event.Subscription, error) {

	var assetIdRule []interface{}
	for _, assetIdItem := range assetId {
		assetIdRule = append(assetIdRule, assetIdItem)
	}
	var sellerRule []interface{}
	for _, sellerItem := range seller {
		sellerRule = append(sellerRule, sellerItem)
	}

	logs, sub, err := _OrderStorage.contract.WatchLogs(opts, "OrderCreated", assetIdRule, sellerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OrderStorageOrderCreated)
				if err := _OrderStorage.contract.UnpackLog(event, "OrderCreated", log); err != nil {
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
// Solidity: event OrderCreated(bytes32 id, uint256 indexed assetId, address indexed seller, address nftAddress, uint256 priceInWei, uint256 expiresAt)
func (_OrderStorage *OrderStorageFilterer) ParseOrderCreated(log types.Log) (*OrderStorageOrderCreated, error) {
	event := new(OrderStorageOrderCreated)
	if err := _OrderStorage.contract.UnpackLog(event, "OrderCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OrderStorageOrderSuccessfulIterator is returned from FilterOrderSuccessful and is used to iterate over the raw logs and unpacked data for OrderSuccessful events raised by the OrderStorage contract.
type OrderStorageOrderSuccessfulIterator struct {
	Event *OrderStorageOrderSuccessful // Event containing the contract specifics and raw log

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
func (it *OrderStorageOrderSuccessfulIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OrderStorageOrderSuccessful)
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
		it.Event = new(OrderStorageOrderSuccessful)
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
func (it *OrderStorageOrderSuccessfulIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OrderStorageOrderSuccessfulIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OrderStorageOrderSuccessful represents a OrderSuccessful event raised by the OrderStorage contract.
type OrderStorageOrderSuccessful struct {
	Id         [32]byte
	AssetId    *big.Int
	Seller     common.Address
	NftAddress common.Address
	TotalPrice *big.Int
	Buyer      common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterOrderSuccessful is a free log retrieval operation binding the contract event 0x695ec315e8a642a74d450a4505eeea53df699b47a7378c7d752e97d5b16eb9bb.
//
// Solidity: event OrderSuccessful(bytes32 id, uint256 indexed assetId, address indexed seller, address nftAddress, uint256 totalPrice, address indexed buyer)
func (_OrderStorage *OrderStorageFilterer) FilterOrderSuccessful(opts *bind.FilterOpts, assetId []*big.Int, seller []common.Address, buyer []common.Address) (*OrderStorageOrderSuccessfulIterator, error) {

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

	logs, sub, err := _OrderStorage.contract.FilterLogs(opts, "OrderSuccessful", assetIdRule, sellerRule, buyerRule)
	if err != nil {
		return nil, err
	}
	return &OrderStorageOrderSuccessfulIterator{contract: _OrderStorage.contract, event: "OrderSuccessful", logs: logs, sub: sub}, nil
}

// WatchOrderSuccessful is a free log subscription operation binding the contract event 0x695ec315e8a642a74d450a4505eeea53df699b47a7378c7d752e97d5b16eb9bb.
//
// Solidity: event OrderSuccessful(bytes32 id, uint256 indexed assetId, address indexed seller, address nftAddress, uint256 totalPrice, address indexed buyer)
func (_OrderStorage *OrderStorageFilterer) WatchOrderSuccessful(opts *bind.WatchOpts, sink chan<- *OrderStorageOrderSuccessful, assetId []*big.Int, seller []common.Address, buyer []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _OrderStorage.contract.WatchLogs(opts, "OrderSuccessful", assetIdRule, sellerRule, buyerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OrderStorageOrderSuccessful)
				if err := _OrderStorage.contract.UnpackLog(event, "OrderSuccessful", log); err != nil {
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

// ParseOrderSuccessful is a log parse operation binding the contract event 0x695ec315e8a642a74d450a4505eeea53df699b47a7378c7d752e97d5b16eb9bb.
//
// Solidity: event OrderSuccessful(bytes32 id, uint256 indexed assetId, address indexed seller, address nftAddress, uint256 totalPrice, address indexed buyer)
func (_OrderStorage *OrderStorageFilterer) ParseOrderSuccessful(log types.Log) (*OrderStorageOrderSuccessful, error) {
	event := new(OrderStorageOrderSuccessful)
	if err := _OrderStorage.contract.UnpackLog(event, "OrderSuccessful", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
