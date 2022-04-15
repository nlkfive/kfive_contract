// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package NativeMetaTransaction

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

// NativeMetaTransactionMetaData contains all meta data concerning the NativeMetaTransaction contract.
var NativeMetaTransactionMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"userAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"relayerAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"functionSignature\",\"type\":\"bytes\"}],\"name\":\"MetaTransactionExecuted\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"domainSeparator\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getChainId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"userAddress\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"functionSignature\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"sigR\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"sigS\",\"type\":\"bytes32\"},{\"internalType\":\"uint8\",\"name\":\"sigV\",\"type\":\"uint8\"}],\"name\":\"executeMetaTransaction\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"getNonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// NativeMetaTransactionABI is the input ABI used to generate the binding from.
// Deprecated: Use NativeMetaTransactionMetaData.ABI instead.
var NativeMetaTransactionABI = NativeMetaTransactionMetaData.ABI

// NativeMetaTransaction is an auto generated Go binding around an Ethereum contract.
type NativeMetaTransaction struct {
	NativeMetaTransactionCaller     // Read-only binding to the contract
	NativeMetaTransactionTransactor // Write-only binding to the contract
	NativeMetaTransactionFilterer   // Log filterer for contract events
}

// NativeMetaTransactionCaller is an auto generated read-only Go binding around an Ethereum contract.
type NativeMetaTransactionCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NativeMetaTransactionTransactor is an auto generated write-only Go binding around an Ethereum contract.
type NativeMetaTransactionTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NativeMetaTransactionFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type NativeMetaTransactionFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NativeMetaTransactionSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type NativeMetaTransactionSession struct {
	Contract     *NativeMetaTransaction // Generic contract binding to set the session for
	CallOpts     bind.CallOpts          // Call options to use throughout this session
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// NativeMetaTransactionCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type NativeMetaTransactionCallerSession struct {
	Contract *NativeMetaTransactionCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                // Call options to use throughout this session
}

// NativeMetaTransactionTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type NativeMetaTransactionTransactorSession struct {
	Contract     *NativeMetaTransactionTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                // Transaction auth options to use throughout this session
}

// NativeMetaTransactionRaw is an auto generated low-level Go binding around an Ethereum contract.
type NativeMetaTransactionRaw struct {
	Contract *NativeMetaTransaction // Generic contract binding to access the raw methods on
}

// NativeMetaTransactionCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type NativeMetaTransactionCallerRaw struct {
	Contract *NativeMetaTransactionCaller // Generic read-only contract binding to access the raw methods on
}

// NativeMetaTransactionTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type NativeMetaTransactionTransactorRaw struct {
	Contract *NativeMetaTransactionTransactor // Generic write-only contract binding to access the raw methods on
}

// NewNativeMetaTransaction creates a new instance of NativeMetaTransaction, bound to a specific deployed contract.
func NewNativeMetaTransaction(address common.Address, backend bind.ContractBackend) (*NativeMetaTransaction, error) {
	contract, err := bindNativeMetaTransaction(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &NativeMetaTransaction{NativeMetaTransactionCaller: NativeMetaTransactionCaller{contract: contract}, NativeMetaTransactionTransactor: NativeMetaTransactionTransactor{contract: contract}, NativeMetaTransactionFilterer: NativeMetaTransactionFilterer{contract: contract}}, nil
}

// NewNativeMetaTransactionCaller creates a new read-only instance of NativeMetaTransaction, bound to a specific deployed contract.
func NewNativeMetaTransactionCaller(address common.Address, caller bind.ContractCaller) (*NativeMetaTransactionCaller, error) {
	contract, err := bindNativeMetaTransaction(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &NativeMetaTransactionCaller{contract: contract}, nil
}

// NewNativeMetaTransactionTransactor creates a new write-only instance of NativeMetaTransaction, bound to a specific deployed contract.
func NewNativeMetaTransactionTransactor(address common.Address, transactor bind.ContractTransactor) (*NativeMetaTransactionTransactor, error) {
	contract, err := bindNativeMetaTransaction(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &NativeMetaTransactionTransactor{contract: contract}, nil
}

// NewNativeMetaTransactionFilterer creates a new log filterer instance of NativeMetaTransaction, bound to a specific deployed contract.
func NewNativeMetaTransactionFilterer(address common.Address, filterer bind.ContractFilterer) (*NativeMetaTransactionFilterer, error) {
	contract, err := bindNativeMetaTransaction(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &NativeMetaTransactionFilterer{contract: contract}, nil
}

// bindNativeMetaTransaction binds a generic wrapper to an already deployed contract.
func bindNativeMetaTransaction(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(NativeMetaTransactionABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NativeMetaTransaction *NativeMetaTransactionRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _NativeMetaTransaction.Contract.NativeMetaTransactionCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NativeMetaTransaction *NativeMetaTransactionRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NativeMetaTransaction.Contract.NativeMetaTransactionTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NativeMetaTransaction *NativeMetaTransactionRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NativeMetaTransaction.Contract.NativeMetaTransactionTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NativeMetaTransaction *NativeMetaTransactionCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _NativeMetaTransaction.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NativeMetaTransaction *NativeMetaTransactionTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NativeMetaTransaction.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NativeMetaTransaction *NativeMetaTransactionTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NativeMetaTransaction.Contract.contract.Transact(opts, method, params...)
}

// DomainSeparator is a free data retrieval call binding the contract method 0xf698da25.
//
// Solidity: function domainSeparator() view returns(bytes32)
func (_NativeMetaTransaction *NativeMetaTransactionCaller) DomainSeparator(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _NativeMetaTransaction.contract.Call(opts, &out, "domainSeparator")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DomainSeparator is a free data retrieval call binding the contract method 0xf698da25.
//
// Solidity: function domainSeparator() view returns(bytes32)
func (_NativeMetaTransaction *NativeMetaTransactionSession) DomainSeparator() ([32]byte, error) {
	return _NativeMetaTransaction.Contract.DomainSeparator(&_NativeMetaTransaction.CallOpts)
}

// DomainSeparator is a free data retrieval call binding the contract method 0xf698da25.
//
// Solidity: function domainSeparator() view returns(bytes32)
func (_NativeMetaTransaction *NativeMetaTransactionCallerSession) DomainSeparator() ([32]byte, error) {
	return _NativeMetaTransaction.Contract.DomainSeparator(&_NativeMetaTransaction.CallOpts)
}

// GetChainId is a free data retrieval call binding the contract method 0x3408e470.
//
// Solidity: function getChainId() view returns(uint256)
func (_NativeMetaTransaction *NativeMetaTransactionCaller) GetChainId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NativeMetaTransaction.contract.Call(opts, &out, "getChainId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetChainId is a free data retrieval call binding the contract method 0x3408e470.
//
// Solidity: function getChainId() view returns(uint256)
func (_NativeMetaTransaction *NativeMetaTransactionSession) GetChainId() (*big.Int, error) {
	return _NativeMetaTransaction.Contract.GetChainId(&_NativeMetaTransaction.CallOpts)
}

// GetChainId is a free data retrieval call binding the contract method 0x3408e470.
//
// Solidity: function getChainId() view returns(uint256)
func (_NativeMetaTransaction *NativeMetaTransactionCallerSession) GetChainId() (*big.Int, error) {
	return _NativeMetaTransaction.Contract.GetChainId(&_NativeMetaTransaction.CallOpts)
}

// GetNonce is a free data retrieval call binding the contract method 0x2d0335ab.
//
// Solidity: function getNonce(address user) view returns(uint256 nonce)
func (_NativeMetaTransaction *NativeMetaTransactionCaller) GetNonce(opts *bind.CallOpts, user common.Address) (*big.Int, error) {
	var out []interface{}
	err := _NativeMetaTransaction.contract.Call(opts, &out, "getNonce", user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNonce is a free data retrieval call binding the contract method 0x2d0335ab.
//
// Solidity: function getNonce(address user) view returns(uint256 nonce)
func (_NativeMetaTransaction *NativeMetaTransactionSession) GetNonce(user common.Address) (*big.Int, error) {
	return _NativeMetaTransaction.Contract.GetNonce(&_NativeMetaTransaction.CallOpts, user)
}

// GetNonce is a free data retrieval call binding the contract method 0x2d0335ab.
//
// Solidity: function getNonce(address user) view returns(uint256 nonce)
func (_NativeMetaTransaction *NativeMetaTransactionCallerSession) GetNonce(user common.Address) (*big.Int, error) {
	return _NativeMetaTransaction.Contract.GetNonce(&_NativeMetaTransaction.CallOpts, user)
}

// ExecuteMetaTransaction is a paid mutator transaction binding the contract method 0x0c53c51c.
//
// Solidity: function executeMetaTransaction(address userAddress, bytes functionSignature, bytes32 sigR, bytes32 sigS, uint8 sigV) payable returns(bytes)
func (_NativeMetaTransaction *NativeMetaTransactionTransactor) ExecuteMetaTransaction(opts *bind.TransactOpts, userAddress common.Address, functionSignature []byte, sigR [32]byte, sigS [32]byte, sigV uint8) (*types.Transaction, error) {
	return _NativeMetaTransaction.contract.Transact(opts, "executeMetaTransaction", userAddress, functionSignature, sigR, sigS, sigV)
}

// ExecuteMetaTransaction is a paid mutator transaction binding the contract method 0x0c53c51c.
//
// Solidity: function executeMetaTransaction(address userAddress, bytes functionSignature, bytes32 sigR, bytes32 sigS, uint8 sigV) payable returns(bytes)
func (_NativeMetaTransaction *NativeMetaTransactionSession) ExecuteMetaTransaction(userAddress common.Address, functionSignature []byte, sigR [32]byte, sigS [32]byte, sigV uint8) (*types.Transaction, error) {
	return _NativeMetaTransaction.Contract.ExecuteMetaTransaction(&_NativeMetaTransaction.TransactOpts, userAddress, functionSignature, sigR, sigS, sigV)
}

// ExecuteMetaTransaction is a paid mutator transaction binding the contract method 0x0c53c51c.
//
// Solidity: function executeMetaTransaction(address userAddress, bytes functionSignature, bytes32 sigR, bytes32 sigS, uint8 sigV) payable returns(bytes)
func (_NativeMetaTransaction *NativeMetaTransactionTransactorSession) ExecuteMetaTransaction(userAddress common.Address, functionSignature []byte, sigR [32]byte, sigS [32]byte, sigV uint8) (*types.Transaction, error) {
	return _NativeMetaTransaction.Contract.ExecuteMetaTransaction(&_NativeMetaTransaction.TransactOpts, userAddress, functionSignature, sigR, sigS, sigV)
}

// NativeMetaTransactionMetaTransactionExecutedIterator is returned from FilterMetaTransactionExecuted and is used to iterate over the raw logs and unpacked data for MetaTransactionExecuted events raised by the NativeMetaTransaction contract.
type NativeMetaTransactionMetaTransactionExecutedIterator struct {
	Event *NativeMetaTransactionMetaTransactionExecuted // Event containing the contract specifics and raw log

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
func (it *NativeMetaTransactionMetaTransactionExecutedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NativeMetaTransactionMetaTransactionExecuted)
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
		it.Event = new(NativeMetaTransactionMetaTransactionExecuted)
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
func (it *NativeMetaTransactionMetaTransactionExecutedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NativeMetaTransactionMetaTransactionExecutedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NativeMetaTransactionMetaTransactionExecuted represents a MetaTransactionExecuted event raised by the NativeMetaTransaction contract.
type NativeMetaTransactionMetaTransactionExecuted struct {
	UserAddress       common.Address
	RelayerAddress    common.Address
	FunctionSignature []byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterMetaTransactionExecuted is a free log retrieval operation binding the contract event 0x5845892132946850460bff5a0083f71031bc5bf9aadcd40f1de79423eac9b10b.
//
// Solidity: event MetaTransactionExecuted(address userAddress, address relayerAddress, bytes functionSignature)
func (_NativeMetaTransaction *NativeMetaTransactionFilterer) FilterMetaTransactionExecuted(opts *bind.FilterOpts) (*NativeMetaTransactionMetaTransactionExecutedIterator, error) {

	logs, sub, err := _NativeMetaTransaction.contract.FilterLogs(opts, "MetaTransactionExecuted")
	if err != nil {
		return nil, err
	}
	return &NativeMetaTransactionMetaTransactionExecutedIterator{contract: _NativeMetaTransaction.contract, event: "MetaTransactionExecuted", logs: logs, sub: sub}, nil
}

// WatchMetaTransactionExecuted is a free log subscription operation binding the contract event 0x5845892132946850460bff5a0083f71031bc5bf9aadcd40f1de79423eac9b10b.
//
// Solidity: event MetaTransactionExecuted(address userAddress, address relayerAddress, bytes functionSignature)
func (_NativeMetaTransaction *NativeMetaTransactionFilterer) WatchMetaTransactionExecuted(opts *bind.WatchOpts, sink chan<- *NativeMetaTransactionMetaTransactionExecuted) (event.Subscription, error) {

	logs, sub, err := _NativeMetaTransaction.contract.WatchLogs(opts, "MetaTransactionExecuted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NativeMetaTransactionMetaTransactionExecuted)
				if err := _NativeMetaTransaction.contract.UnpackLog(event, "MetaTransactionExecuted", log); err != nil {
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

// ParseMetaTransactionExecuted is a log parse operation binding the contract event 0x5845892132946850460bff5a0083f71031bc5bf9aadcd40f1de79423eac9b10b.
//
// Solidity: event MetaTransactionExecuted(address userAddress, address relayerAddress, bytes functionSignature)
func (_NativeMetaTransaction *NativeMetaTransactionFilterer) ParseMetaTransactionExecuted(log types.Log) (*NativeMetaTransactionMetaTransactionExecuted, error) {
	event := new(NativeMetaTransactionMetaTransactionExecuted)
	if err := _NativeMetaTransaction.contract.UnpackLog(event, "MetaTransactionExecuted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
