// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package TokenRegistry

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

// TokenRegistryMetaData contains all meta data concerning the TokenRegistry contract.
var TokenRegistryMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"offchain\",\"type\":\"bytes32\"}],\"name\":\"__addNewToken\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"a\",\"type\":\"address\"}],\"name\":\"addAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"admin\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"a\",\"type\":\"address\"}],\"name\":\"removeAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"tokenMapping\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"decimals\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"valid\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"decimals\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"offchain\",\"type\":\"bytes32\"}],\"name\":\"addNewToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"tokenIsExisted\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"getBalanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"getTokenByAddr\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// TokenRegistryABI is the input ABI used to generate the binding from.
// Deprecated: Use TokenRegistryMetaData.ABI instead.
var TokenRegistryABI = TokenRegistryMetaData.ABI

// TokenRegistry is an auto generated Go binding around an Ethereum contract.
type TokenRegistry struct {
	TokenRegistryCaller     // Read-only binding to the contract
	TokenRegistryTransactor // Write-only binding to the contract
	TokenRegistryFilterer   // Log filterer for contract events
}

// TokenRegistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type TokenRegistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenRegistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TokenRegistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenRegistryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TokenRegistryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenRegistrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TokenRegistrySession struct {
	Contract     *TokenRegistry    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TokenRegistryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TokenRegistryCallerSession struct {
	Contract *TokenRegistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// TokenRegistryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TokenRegistryTransactorSession struct {
	Contract     *TokenRegistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// TokenRegistryRaw is an auto generated low-level Go binding around an Ethereum contract.
type TokenRegistryRaw struct {
	Contract *TokenRegistry // Generic contract binding to access the raw methods on
}

// TokenRegistryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TokenRegistryCallerRaw struct {
	Contract *TokenRegistryCaller // Generic read-only contract binding to access the raw methods on
}

// TokenRegistryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TokenRegistryTransactorRaw struct {
	Contract *TokenRegistryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTokenRegistry creates a new instance of TokenRegistry, bound to a specific deployed contract.
func NewTokenRegistry(address common.Address, backend bind.ContractBackend) (*TokenRegistry, error) {
	contract, err := bindTokenRegistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TokenRegistry{TokenRegistryCaller: TokenRegistryCaller{contract: contract}, TokenRegistryTransactor: TokenRegistryTransactor{contract: contract}, TokenRegistryFilterer: TokenRegistryFilterer{contract: contract}}, nil
}

// NewTokenRegistryCaller creates a new read-only instance of TokenRegistry, bound to a specific deployed contract.
func NewTokenRegistryCaller(address common.Address, caller bind.ContractCaller) (*TokenRegistryCaller, error) {
	contract, err := bindTokenRegistry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TokenRegistryCaller{contract: contract}, nil
}

// NewTokenRegistryTransactor creates a new write-only instance of TokenRegistry, bound to a specific deployed contract.
func NewTokenRegistryTransactor(address common.Address, transactor bind.ContractTransactor) (*TokenRegistryTransactor, error) {
	contract, err := bindTokenRegistry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TokenRegistryTransactor{contract: contract}, nil
}

// NewTokenRegistryFilterer creates a new log filterer instance of TokenRegistry, bound to a specific deployed contract.
func NewTokenRegistryFilterer(address common.Address, filterer bind.ContractFilterer) (*TokenRegistryFilterer, error) {
	contract, err := bindTokenRegistry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TokenRegistryFilterer{contract: contract}, nil
}

// bindTokenRegistry binds a generic wrapper to an already deployed contract.
func bindTokenRegistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TokenRegistryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TokenRegistry *TokenRegistryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TokenRegistry.Contract.TokenRegistryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TokenRegistry *TokenRegistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TokenRegistry.Contract.TokenRegistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TokenRegistry *TokenRegistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TokenRegistry.Contract.TokenRegistryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TokenRegistry *TokenRegistryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TokenRegistry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TokenRegistry *TokenRegistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TokenRegistry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TokenRegistry *TokenRegistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TokenRegistry.Contract.contract.Transact(opts, method, params...)
}

// Admin is a free data retrieval call binding the contract method 0x63a846f8.
//
// Solidity: function admin(address ) view returns(bool)
func (_TokenRegistry *TokenRegistryCaller) Admin(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _TokenRegistry.contract.Call(opts, &out, "admin", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Admin is a free data retrieval call binding the contract method 0x63a846f8.
//
// Solidity: function admin(address ) view returns(bool)
func (_TokenRegistry *TokenRegistrySession) Admin(arg0 common.Address) (bool, error) {
	return _TokenRegistry.Contract.Admin(&_TokenRegistry.CallOpts, arg0)
}

// Admin is a free data retrieval call binding the contract method 0x63a846f8.
//
// Solidity: function admin(address ) view returns(bool)
func (_TokenRegistry *TokenRegistryCallerSession) Admin(arg0 common.Address) (bool, error) {
	return _TokenRegistry.Contract.Admin(&_TokenRegistry.CallOpts, arg0)
}

// GetBalanceOf is a free data retrieval call binding the contract method 0x53290b44.
//
// Solidity: function getBalanceOf(address token, address addr) view returns(uint256)
func (_TokenRegistry *TokenRegistryCaller) GetBalanceOf(opts *bind.CallOpts, token common.Address, addr common.Address) (*big.Int, error) {
	var out []interface{}
	err := _TokenRegistry.contract.Call(opts, &out, "getBalanceOf", token, addr)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBalanceOf is a free data retrieval call binding the contract method 0x53290b44.
//
// Solidity: function getBalanceOf(address token, address addr) view returns(uint256)
func (_TokenRegistry *TokenRegistrySession) GetBalanceOf(token common.Address, addr common.Address) (*big.Int, error) {
	return _TokenRegistry.Contract.GetBalanceOf(&_TokenRegistry.CallOpts, token, addr)
}

// GetBalanceOf is a free data retrieval call binding the contract method 0x53290b44.
//
// Solidity: function getBalanceOf(address token, address addr) view returns(uint256)
func (_TokenRegistry *TokenRegistryCallerSession) GetBalanceOf(token common.Address, addr common.Address) (*big.Int, error) {
	return _TokenRegistry.Contract.GetBalanceOf(&_TokenRegistry.CallOpts, token, addr)
}

// GetTokenByAddr is a free data retrieval call binding the contract method 0x70571828.
//
// Solidity: function getTokenByAddr(address token) view returns(string, string, uint8)
func (_TokenRegistry *TokenRegistryCaller) GetTokenByAddr(opts *bind.CallOpts, token common.Address) (string, string, uint8, error) {
	var out []interface{}
	err := _TokenRegistry.contract.Call(opts, &out, "getTokenByAddr", token)

	if err != nil {
		return *new(string), *new(string), *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)
	out1 := *abi.ConvertType(out[1], new(string)).(*string)
	out2 := *abi.ConvertType(out[2], new(uint8)).(*uint8)

	return out0, out1, out2, err

}

// GetTokenByAddr is a free data retrieval call binding the contract method 0x70571828.
//
// Solidity: function getTokenByAddr(address token) view returns(string, string, uint8)
func (_TokenRegistry *TokenRegistrySession) GetTokenByAddr(token common.Address) (string, string, uint8, error) {
	return _TokenRegistry.Contract.GetTokenByAddr(&_TokenRegistry.CallOpts, token)
}

// GetTokenByAddr is a free data retrieval call binding the contract method 0x70571828.
//
// Solidity: function getTokenByAddr(address token) view returns(string, string, uint8)
func (_TokenRegistry *TokenRegistryCallerSession) GetTokenByAddr(token common.Address) (string, string, uint8, error) {
	return _TokenRegistry.Contract.GetTokenByAddr(&_TokenRegistry.CallOpts, token)
}

// TokenIsExisted is a free data retrieval call binding the contract method 0x61eeeb8c.
//
// Solidity: function tokenIsExisted(address token) view returns(bool)
func (_TokenRegistry *TokenRegistryCaller) TokenIsExisted(opts *bind.CallOpts, token common.Address) (bool, error) {
	var out []interface{}
	err := _TokenRegistry.contract.Call(opts, &out, "tokenIsExisted", token)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// TokenIsExisted is a free data retrieval call binding the contract method 0x61eeeb8c.
//
// Solidity: function tokenIsExisted(address token) view returns(bool)
func (_TokenRegistry *TokenRegistrySession) TokenIsExisted(token common.Address) (bool, error) {
	return _TokenRegistry.Contract.TokenIsExisted(&_TokenRegistry.CallOpts, token)
}

// TokenIsExisted is a free data retrieval call binding the contract method 0x61eeeb8c.
//
// Solidity: function tokenIsExisted(address token) view returns(bool)
func (_TokenRegistry *TokenRegistryCallerSession) TokenIsExisted(token common.Address) (bool, error) {
	return _TokenRegistry.Contract.TokenIsExisted(&_TokenRegistry.CallOpts, token)
}

// TokenMapping is a free data retrieval call binding the contract method 0xba27f50b.
//
// Solidity: function tokenMapping(address ) view returns(address token, string symbol, string name, uint8 decimals, bool valid)
func (_TokenRegistry *TokenRegistryCaller) TokenMapping(opts *bind.CallOpts, arg0 common.Address) (struct {
	Token    common.Address
	Symbol   string
	Name     string
	Decimals uint8
	Valid    bool
}, error) {
	var out []interface{}
	err := _TokenRegistry.contract.Call(opts, &out, "tokenMapping", arg0)

	outstruct := new(struct {
		Token    common.Address
		Symbol   string
		Name     string
		Decimals uint8
		Valid    bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Token = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Symbol = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.Name = *abi.ConvertType(out[2], new(string)).(*string)
	outstruct.Decimals = *abi.ConvertType(out[3], new(uint8)).(*uint8)
	outstruct.Valid = *abi.ConvertType(out[4], new(bool)).(*bool)

	return *outstruct, err

}

// TokenMapping is a free data retrieval call binding the contract method 0xba27f50b.
//
// Solidity: function tokenMapping(address ) view returns(address token, string symbol, string name, uint8 decimals, bool valid)
func (_TokenRegistry *TokenRegistrySession) TokenMapping(arg0 common.Address) (struct {
	Token    common.Address
	Symbol   string
	Name     string
	Decimals uint8
	Valid    bool
}, error) {
	return _TokenRegistry.Contract.TokenMapping(&_TokenRegistry.CallOpts, arg0)
}

// TokenMapping is a free data retrieval call binding the contract method 0xba27f50b.
//
// Solidity: function tokenMapping(address ) view returns(address token, string symbol, string name, uint8 decimals, bool valid)
func (_TokenRegistry *TokenRegistryCallerSession) TokenMapping(arg0 common.Address) (struct {
	Token    common.Address
	Symbol   string
	Name     string
	Decimals uint8
	Valid    bool
}, error) {
	return _TokenRegistry.Contract.TokenMapping(&_TokenRegistry.CallOpts, arg0)
}

// AddAdmin is a paid mutator transaction binding the contract method 0x70480275.
//
// Solidity: function addAdmin(address a) returns()
func (_TokenRegistry *TokenRegistryTransactor) AddAdmin(opts *bind.TransactOpts, a common.Address) (*types.Transaction, error) {
	return _TokenRegistry.contract.Transact(opts, "addAdmin", a)
}

// AddAdmin is a paid mutator transaction binding the contract method 0x70480275.
//
// Solidity: function addAdmin(address a) returns()
func (_TokenRegistry *TokenRegistrySession) AddAdmin(a common.Address) (*types.Transaction, error) {
	return _TokenRegistry.Contract.AddAdmin(&_TokenRegistry.TransactOpts, a)
}

// AddAdmin is a paid mutator transaction binding the contract method 0x70480275.
//
// Solidity: function addAdmin(address a) returns()
func (_TokenRegistry *TokenRegistryTransactorSession) AddAdmin(a common.Address) (*types.Transaction, error) {
	return _TokenRegistry.Contract.AddAdmin(&_TokenRegistry.TransactOpts, a)
}

// AddNewToken is a paid mutator transaction binding the contract method 0x99a32482.
//
// Solidity: function addNewToken(address token, string symbol, string name, uint8 decimals, bytes32 offchain) returns()
func (_TokenRegistry *TokenRegistryTransactor) AddNewToken(opts *bind.TransactOpts, token common.Address, symbol string, name string, decimals uint8, offchain [32]byte) (*types.Transaction, error) {
	return _TokenRegistry.contract.Transact(opts, "addNewToken", token, symbol, name, decimals, offchain)
}

// AddNewToken is a paid mutator transaction binding the contract method 0x99a32482.
//
// Solidity: function addNewToken(address token, string symbol, string name, uint8 decimals, bytes32 offchain) returns()
func (_TokenRegistry *TokenRegistrySession) AddNewToken(token common.Address, symbol string, name string, decimals uint8, offchain [32]byte) (*types.Transaction, error) {
	return _TokenRegistry.Contract.AddNewToken(&_TokenRegistry.TransactOpts, token, symbol, name, decimals, offchain)
}

// AddNewToken is a paid mutator transaction binding the contract method 0x99a32482.
//
// Solidity: function addNewToken(address token, string symbol, string name, uint8 decimals, bytes32 offchain) returns()
func (_TokenRegistry *TokenRegistryTransactorSession) AddNewToken(token common.Address, symbol string, name string, decimals uint8, offchain [32]byte) (*types.Transaction, error) {
	return _TokenRegistry.Contract.AddNewToken(&_TokenRegistry.TransactOpts, token, symbol, name, decimals, offchain)
}

// RemoveAdmin is a paid mutator transaction binding the contract method 0x1785f53c.
//
// Solidity: function removeAdmin(address a) returns()
func (_TokenRegistry *TokenRegistryTransactor) RemoveAdmin(opts *bind.TransactOpts, a common.Address) (*types.Transaction, error) {
	return _TokenRegistry.contract.Transact(opts, "removeAdmin", a)
}

// RemoveAdmin is a paid mutator transaction binding the contract method 0x1785f53c.
//
// Solidity: function removeAdmin(address a) returns()
func (_TokenRegistry *TokenRegistrySession) RemoveAdmin(a common.Address) (*types.Transaction, error) {
	return _TokenRegistry.Contract.RemoveAdmin(&_TokenRegistry.TransactOpts, a)
}

// RemoveAdmin is a paid mutator transaction binding the contract method 0x1785f53c.
//
// Solidity: function removeAdmin(address a) returns()
func (_TokenRegistry *TokenRegistryTransactorSession) RemoveAdmin(a common.Address) (*types.Transaction, error) {
	return _TokenRegistry.Contract.RemoveAdmin(&_TokenRegistry.TransactOpts, a)
}

// TransferToken is a paid mutator transaction binding the contract method 0x2c54de4f.
//
// Solidity: function transferToken(address token, address from, address to, uint256 amount) returns()
func (_TokenRegistry *TokenRegistryTransactor) TransferToken(opts *bind.TransactOpts, token common.Address, from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TokenRegistry.contract.Transact(opts, "transferToken", token, from, to, amount)
}

// TransferToken is a paid mutator transaction binding the contract method 0x2c54de4f.
//
// Solidity: function transferToken(address token, address from, address to, uint256 amount) returns()
func (_TokenRegistry *TokenRegistrySession) TransferToken(token common.Address, from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TokenRegistry.Contract.TransferToken(&_TokenRegistry.TransactOpts, token, from, to, amount)
}

// TransferToken is a paid mutator transaction binding the contract method 0x2c54de4f.
//
// Solidity: function transferToken(address token, address from, address to, uint256 amount) returns()
func (_TokenRegistry *TokenRegistryTransactorSession) TransferToken(token common.Address, from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TokenRegistry.Contract.TransferToken(&_TokenRegistry.TransactOpts, token, from, to, amount)
}

// TokenRegistryAddNewTokenIterator is returned from FilterAddNewToken and is used to iterate over the raw logs and unpacked data for AddNewToken events raised by the TokenRegistry contract.
type TokenRegistryAddNewTokenIterator struct {
	Event *TokenRegistryAddNewToken // Event containing the contract specifics and raw log

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
func (it *TokenRegistryAddNewTokenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenRegistryAddNewToken)
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
		it.Event = new(TokenRegistryAddNewToken)
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
func (it *TokenRegistryAddNewTokenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenRegistryAddNewTokenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenRegistryAddNewToken represents a AddNewToken event raised by the TokenRegistry contract.
type TokenRegistryAddNewToken struct {
	Offchain [32]byte
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterAddNewToken is a free log retrieval operation binding the contract event 0x662298be6d0168dcaa626ce9f521c3992f79a1ae8a49dd23a143dd4ebf7d4b4a.
//
// Solidity: event __addNewToken(bytes32 offchain)
func (_TokenRegistry *TokenRegistryFilterer) FilterAddNewToken(opts *bind.FilterOpts) (*TokenRegistryAddNewTokenIterator, error) {

	logs, sub, err := _TokenRegistry.contract.FilterLogs(opts, "__addNewToken")
	if err != nil {
		return nil, err
	}
	return &TokenRegistryAddNewTokenIterator{contract: _TokenRegistry.contract, event: "__addNewToken", logs: logs, sub: sub}, nil
}

// WatchAddNewToken is a free log subscription operation binding the contract event 0x662298be6d0168dcaa626ce9f521c3992f79a1ae8a49dd23a143dd4ebf7d4b4a.
//
// Solidity: event __addNewToken(bytes32 offchain)
func (_TokenRegistry *TokenRegistryFilterer) WatchAddNewToken(opts *bind.WatchOpts, sink chan<- *TokenRegistryAddNewToken) (event.Subscription, error) {

	logs, sub, err := _TokenRegistry.contract.WatchLogs(opts, "__addNewToken")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenRegistryAddNewToken)
				if err := _TokenRegistry.contract.UnpackLog(event, "__addNewToken", log); err != nil {
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

// ParseAddNewToken is a log parse operation binding the contract event 0x662298be6d0168dcaa626ce9f521c3992f79a1ae8a49dd23a143dd4ebf7d4b4a.
//
// Solidity: event __addNewToken(bytes32 offchain)
func (_TokenRegistry *TokenRegistryFilterer) ParseAddNewToken(log types.Log) (*TokenRegistryAddNewToken, error) {
	event := new(TokenRegistryAddNewToken)
	if err := _TokenRegistry.contract.UnpackLog(event, "__addNewToken", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
