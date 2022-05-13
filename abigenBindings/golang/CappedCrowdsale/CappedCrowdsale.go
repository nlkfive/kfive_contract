// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package CappedCrowdsale

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

// CappedCrowdsaleMetaData contains all meta data concerning the CappedCrowdsale contract.
var CappedCrowdsaleMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"purchaser\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"beneficiary\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"TokensPurchased\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"beneficiary\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"buyTokens\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"newReceiver\",\"type\":\"address\"}],\"name\":\"changeWallet\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"_to\",\"type\":\"address\"}],\"name\":\"destroySmartContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token\",\"outputs\":[{\"internalType\":\"contractIBEP20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"wallet\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"weiRaised\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"cap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"capReached\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// CappedCrowdsaleABI is the input ABI used to generate the binding from.
// Deprecated: Use CappedCrowdsaleMetaData.ABI instead.
var CappedCrowdsaleABI = CappedCrowdsaleMetaData.ABI

// CappedCrowdsale is an auto generated Go binding around an Ethereum contract.
type CappedCrowdsale struct {
	CappedCrowdsaleCaller     // Read-only binding to the contract
	CappedCrowdsaleTransactor // Write-only binding to the contract
	CappedCrowdsaleFilterer   // Log filterer for contract events
}

// CappedCrowdsaleCaller is an auto generated read-only Go binding around an Ethereum contract.
type CappedCrowdsaleCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CappedCrowdsaleTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CappedCrowdsaleTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CappedCrowdsaleFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CappedCrowdsaleFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CappedCrowdsaleSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CappedCrowdsaleSession struct {
	Contract     *CappedCrowdsale  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CappedCrowdsaleCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CappedCrowdsaleCallerSession struct {
	Contract *CappedCrowdsaleCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// CappedCrowdsaleTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CappedCrowdsaleTransactorSession struct {
	Contract     *CappedCrowdsaleTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// CappedCrowdsaleRaw is an auto generated low-level Go binding around an Ethereum contract.
type CappedCrowdsaleRaw struct {
	Contract *CappedCrowdsale // Generic contract binding to access the raw methods on
}

// CappedCrowdsaleCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CappedCrowdsaleCallerRaw struct {
	Contract *CappedCrowdsaleCaller // Generic read-only contract binding to access the raw methods on
}

// CappedCrowdsaleTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CappedCrowdsaleTransactorRaw struct {
	Contract *CappedCrowdsaleTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCappedCrowdsale creates a new instance of CappedCrowdsale, bound to a specific deployed contract.
func NewCappedCrowdsale(address common.Address, backend bind.ContractBackend) (*CappedCrowdsale, error) {
	contract, err := bindCappedCrowdsale(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CappedCrowdsale{CappedCrowdsaleCaller: CappedCrowdsaleCaller{contract: contract}, CappedCrowdsaleTransactor: CappedCrowdsaleTransactor{contract: contract}, CappedCrowdsaleFilterer: CappedCrowdsaleFilterer{contract: contract}}, nil
}

// NewCappedCrowdsaleCaller creates a new read-only instance of CappedCrowdsale, bound to a specific deployed contract.
func NewCappedCrowdsaleCaller(address common.Address, caller bind.ContractCaller) (*CappedCrowdsaleCaller, error) {
	contract, err := bindCappedCrowdsale(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CappedCrowdsaleCaller{contract: contract}, nil
}

// NewCappedCrowdsaleTransactor creates a new write-only instance of CappedCrowdsale, bound to a specific deployed contract.
func NewCappedCrowdsaleTransactor(address common.Address, transactor bind.ContractTransactor) (*CappedCrowdsaleTransactor, error) {
	contract, err := bindCappedCrowdsale(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CappedCrowdsaleTransactor{contract: contract}, nil
}

// NewCappedCrowdsaleFilterer creates a new log filterer instance of CappedCrowdsale, bound to a specific deployed contract.
func NewCappedCrowdsaleFilterer(address common.Address, filterer bind.ContractFilterer) (*CappedCrowdsaleFilterer, error) {
	contract, err := bindCappedCrowdsale(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CappedCrowdsaleFilterer{contract: contract}, nil
}

// bindCappedCrowdsale binds a generic wrapper to an already deployed contract.
func bindCappedCrowdsale(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(CappedCrowdsaleABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CappedCrowdsale *CappedCrowdsaleRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CappedCrowdsale.Contract.CappedCrowdsaleCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CappedCrowdsale *CappedCrowdsaleRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CappedCrowdsale.Contract.CappedCrowdsaleTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CappedCrowdsale *CappedCrowdsaleRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CappedCrowdsale.Contract.CappedCrowdsaleTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CappedCrowdsale *CappedCrowdsaleCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CappedCrowdsale.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CappedCrowdsale *CappedCrowdsaleTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CappedCrowdsale.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CappedCrowdsale *CappedCrowdsaleTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CappedCrowdsale.Contract.contract.Transact(opts, method, params...)
}

// Cap is a free data retrieval call binding the contract method 0x355274ea.
//
// Solidity: function cap() view returns(uint256)
func (_CappedCrowdsale *CappedCrowdsaleCaller) Cap(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CappedCrowdsale.contract.Call(opts, &out, "cap")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Cap is a free data retrieval call binding the contract method 0x355274ea.
//
// Solidity: function cap() view returns(uint256)
func (_CappedCrowdsale *CappedCrowdsaleSession) Cap() (*big.Int, error) {
	return _CappedCrowdsale.Contract.Cap(&_CappedCrowdsale.CallOpts)
}

// Cap is a free data retrieval call binding the contract method 0x355274ea.
//
// Solidity: function cap() view returns(uint256)
func (_CappedCrowdsale *CappedCrowdsaleCallerSession) Cap() (*big.Int, error) {
	return _CappedCrowdsale.Contract.Cap(&_CappedCrowdsale.CallOpts)
}

// CapReached is a free data retrieval call binding the contract method 0x4f935945.
//
// Solidity: function capReached() view returns(bool)
func (_CappedCrowdsale *CappedCrowdsaleCaller) CapReached(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _CappedCrowdsale.contract.Call(opts, &out, "capReached")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CapReached is a free data retrieval call binding the contract method 0x4f935945.
//
// Solidity: function capReached() view returns(bool)
func (_CappedCrowdsale *CappedCrowdsaleSession) CapReached() (bool, error) {
	return _CappedCrowdsale.Contract.CapReached(&_CappedCrowdsale.CallOpts)
}

// CapReached is a free data retrieval call binding the contract method 0x4f935945.
//
// Solidity: function capReached() view returns(bool)
func (_CappedCrowdsale *CappedCrowdsaleCallerSession) CapReached() (bool, error) {
	return _CappedCrowdsale.Contract.CapReached(&_CappedCrowdsale.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_CappedCrowdsale *CappedCrowdsaleCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CappedCrowdsale.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_CappedCrowdsale *CappedCrowdsaleSession) Owner() (common.Address, error) {
	return _CappedCrowdsale.Contract.Owner(&_CappedCrowdsale.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_CappedCrowdsale *CappedCrowdsaleCallerSession) Owner() (common.Address, error) {
	return _CappedCrowdsale.Contract.Owner(&_CappedCrowdsale.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_CappedCrowdsale *CappedCrowdsaleCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _CappedCrowdsale.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_CappedCrowdsale *CappedCrowdsaleSession) Paused() (bool, error) {
	return _CappedCrowdsale.Contract.Paused(&_CappedCrowdsale.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_CappedCrowdsale *CappedCrowdsaleCallerSession) Paused() (bool, error) {
	return _CappedCrowdsale.Contract.Paused(&_CappedCrowdsale.CallOpts)
}

// Rate is a free data retrieval call binding the contract method 0x2c4e722e.
//
// Solidity: function rate() view returns(uint256)
func (_CappedCrowdsale *CappedCrowdsaleCaller) Rate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CappedCrowdsale.contract.Call(opts, &out, "rate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Rate is a free data retrieval call binding the contract method 0x2c4e722e.
//
// Solidity: function rate() view returns(uint256)
func (_CappedCrowdsale *CappedCrowdsaleSession) Rate() (*big.Int, error) {
	return _CappedCrowdsale.Contract.Rate(&_CappedCrowdsale.CallOpts)
}

// Rate is a free data retrieval call binding the contract method 0x2c4e722e.
//
// Solidity: function rate() view returns(uint256)
func (_CappedCrowdsale *CappedCrowdsaleCallerSession) Rate() (*big.Int, error) {
	return _CappedCrowdsale.Contract.Rate(&_CappedCrowdsale.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_CappedCrowdsale *CappedCrowdsaleCaller) Token(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CappedCrowdsale.contract.Call(opts, &out, "token")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_CappedCrowdsale *CappedCrowdsaleSession) Token() (common.Address, error) {
	return _CappedCrowdsale.Contract.Token(&_CappedCrowdsale.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_CappedCrowdsale *CappedCrowdsaleCallerSession) Token() (common.Address, error) {
	return _CappedCrowdsale.Contract.Token(&_CappedCrowdsale.CallOpts)
}

// Wallet is a free data retrieval call binding the contract method 0x521eb273.
//
// Solidity: function wallet() view returns(address)
func (_CappedCrowdsale *CappedCrowdsaleCaller) Wallet(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CappedCrowdsale.contract.Call(opts, &out, "wallet")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Wallet is a free data retrieval call binding the contract method 0x521eb273.
//
// Solidity: function wallet() view returns(address)
func (_CappedCrowdsale *CappedCrowdsaleSession) Wallet() (common.Address, error) {
	return _CappedCrowdsale.Contract.Wallet(&_CappedCrowdsale.CallOpts)
}

// Wallet is a free data retrieval call binding the contract method 0x521eb273.
//
// Solidity: function wallet() view returns(address)
func (_CappedCrowdsale *CappedCrowdsaleCallerSession) Wallet() (common.Address, error) {
	return _CappedCrowdsale.Contract.Wallet(&_CappedCrowdsale.CallOpts)
}

// WeiRaised is a free data retrieval call binding the contract method 0x4042b66f.
//
// Solidity: function weiRaised() view returns(uint256)
func (_CappedCrowdsale *CappedCrowdsaleCaller) WeiRaised(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CappedCrowdsale.contract.Call(opts, &out, "weiRaised")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WeiRaised is a free data retrieval call binding the contract method 0x4042b66f.
//
// Solidity: function weiRaised() view returns(uint256)
func (_CappedCrowdsale *CappedCrowdsaleSession) WeiRaised() (*big.Int, error) {
	return _CappedCrowdsale.Contract.WeiRaised(&_CappedCrowdsale.CallOpts)
}

// WeiRaised is a free data retrieval call binding the contract method 0x4042b66f.
//
// Solidity: function weiRaised() view returns(uint256)
func (_CappedCrowdsale *CappedCrowdsaleCallerSession) WeiRaised() (*big.Int, error) {
	return _CappedCrowdsale.Contract.WeiRaised(&_CappedCrowdsale.CallOpts)
}

// BuyTokens is a paid mutator transaction binding the contract method 0x0752881a.
//
// Solidity: function buyTokens(address beneficiary, uint256 value) payable returns()
func (_CappedCrowdsale *CappedCrowdsaleTransactor) BuyTokens(opts *bind.TransactOpts, beneficiary common.Address, value *big.Int) (*types.Transaction, error) {
	return _CappedCrowdsale.contract.Transact(opts, "buyTokens", beneficiary, value)
}

// BuyTokens is a paid mutator transaction binding the contract method 0x0752881a.
//
// Solidity: function buyTokens(address beneficiary, uint256 value) payable returns()
func (_CappedCrowdsale *CappedCrowdsaleSession) BuyTokens(beneficiary common.Address, value *big.Int) (*types.Transaction, error) {
	return _CappedCrowdsale.Contract.BuyTokens(&_CappedCrowdsale.TransactOpts, beneficiary, value)
}

// BuyTokens is a paid mutator transaction binding the contract method 0x0752881a.
//
// Solidity: function buyTokens(address beneficiary, uint256 value) payable returns()
func (_CappedCrowdsale *CappedCrowdsaleTransactorSession) BuyTokens(beneficiary common.Address, value *big.Int) (*types.Transaction, error) {
	return _CappedCrowdsale.Contract.BuyTokens(&_CappedCrowdsale.TransactOpts, beneficiary, value)
}

// ChangeWallet is a paid mutator transaction binding the contract method 0x98b9a2dc.
//
// Solidity: function changeWallet(address newReceiver) returns()
func (_CappedCrowdsale *CappedCrowdsaleTransactor) ChangeWallet(opts *bind.TransactOpts, newReceiver common.Address) (*types.Transaction, error) {
	return _CappedCrowdsale.contract.Transact(opts, "changeWallet", newReceiver)
}

// ChangeWallet is a paid mutator transaction binding the contract method 0x98b9a2dc.
//
// Solidity: function changeWallet(address newReceiver) returns()
func (_CappedCrowdsale *CappedCrowdsaleSession) ChangeWallet(newReceiver common.Address) (*types.Transaction, error) {
	return _CappedCrowdsale.Contract.ChangeWallet(&_CappedCrowdsale.TransactOpts, newReceiver)
}

// ChangeWallet is a paid mutator transaction binding the contract method 0x98b9a2dc.
//
// Solidity: function changeWallet(address newReceiver) returns()
func (_CappedCrowdsale *CappedCrowdsaleTransactorSession) ChangeWallet(newReceiver common.Address) (*types.Transaction, error) {
	return _CappedCrowdsale.Contract.ChangeWallet(&_CappedCrowdsale.TransactOpts, newReceiver)
}

// DestroySmartContract is a paid mutator transaction binding the contract method 0x39df43ff.
//
// Solidity: function destroySmartContract(address _to) returns()
func (_CappedCrowdsale *CappedCrowdsaleTransactor) DestroySmartContract(opts *bind.TransactOpts, _to common.Address) (*types.Transaction, error) {
	return _CappedCrowdsale.contract.Transact(opts, "destroySmartContract", _to)
}

// DestroySmartContract is a paid mutator transaction binding the contract method 0x39df43ff.
//
// Solidity: function destroySmartContract(address _to) returns()
func (_CappedCrowdsale *CappedCrowdsaleSession) DestroySmartContract(_to common.Address) (*types.Transaction, error) {
	return _CappedCrowdsale.Contract.DestroySmartContract(&_CappedCrowdsale.TransactOpts, _to)
}

// DestroySmartContract is a paid mutator transaction binding the contract method 0x39df43ff.
//
// Solidity: function destroySmartContract(address _to) returns()
func (_CappedCrowdsale *CappedCrowdsaleTransactorSession) DestroySmartContract(_to common.Address) (*types.Transaction, error) {
	return _CappedCrowdsale.Contract.DestroySmartContract(&_CappedCrowdsale.TransactOpts, _to)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_CappedCrowdsale *CappedCrowdsaleTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CappedCrowdsale.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_CappedCrowdsale *CappedCrowdsaleSession) Pause() (*types.Transaction, error) {
	return _CappedCrowdsale.Contract.Pause(&_CappedCrowdsale.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_CappedCrowdsale *CappedCrowdsaleTransactorSession) Pause() (*types.Transaction, error) {
	return _CappedCrowdsale.Contract.Pause(&_CappedCrowdsale.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_CappedCrowdsale *CappedCrowdsaleTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CappedCrowdsale.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_CappedCrowdsale *CappedCrowdsaleSession) RenounceOwnership() (*types.Transaction, error) {
	return _CappedCrowdsale.Contract.RenounceOwnership(&_CappedCrowdsale.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_CappedCrowdsale *CappedCrowdsaleTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _CappedCrowdsale.Contract.RenounceOwnership(&_CappedCrowdsale.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_CappedCrowdsale *CappedCrowdsaleTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _CappedCrowdsale.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_CappedCrowdsale *CappedCrowdsaleSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _CappedCrowdsale.Contract.TransferOwnership(&_CappedCrowdsale.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_CappedCrowdsale *CappedCrowdsaleTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _CappedCrowdsale.Contract.TransferOwnership(&_CappedCrowdsale.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_CappedCrowdsale *CappedCrowdsaleTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CappedCrowdsale.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_CappedCrowdsale *CappedCrowdsaleSession) Unpause() (*types.Transaction, error) {
	return _CappedCrowdsale.Contract.Unpause(&_CappedCrowdsale.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_CappedCrowdsale *CappedCrowdsaleTransactorSession) Unpause() (*types.Transaction, error) {
	return _CappedCrowdsale.Contract.Unpause(&_CappedCrowdsale.TransactOpts)
}

// CappedCrowdsaleOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the CappedCrowdsale contract.
type CappedCrowdsaleOwnershipTransferredIterator struct {
	Event *CappedCrowdsaleOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *CappedCrowdsaleOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CappedCrowdsaleOwnershipTransferred)
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
		it.Event = new(CappedCrowdsaleOwnershipTransferred)
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
func (it *CappedCrowdsaleOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CappedCrowdsaleOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CappedCrowdsaleOwnershipTransferred represents a OwnershipTransferred event raised by the CappedCrowdsale contract.
type CappedCrowdsaleOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_CappedCrowdsale *CappedCrowdsaleFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*CappedCrowdsaleOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _CappedCrowdsale.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &CappedCrowdsaleOwnershipTransferredIterator{contract: _CappedCrowdsale.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_CappedCrowdsale *CappedCrowdsaleFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *CappedCrowdsaleOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _CappedCrowdsale.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CappedCrowdsaleOwnershipTransferred)
				if err := _CappedCrowdsale.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_CappedCrowdsale *CappedCrowdsaleFilterer) ParseOwnershipTransferred(log types.Log) (*CappedCrowdsaleOwnershipTransferred, error) {
	event := new(CappedCrowdsaleOwnershipTransferred)
	if err := _CappedCrowdsale.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CappedCrowdsalePausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the CappedCrowdsale contract.
type CappedCrowdsalePausedIterator struct {
	Event *CappedCrowdsalePaused // Event containing the contract specifics and raw log

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
func (it *CappedCrowdsalePausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CappedCrowdsalePaused)
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
		it.Event = new(CappedCrowdsalePaused)
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
func (it *CappedCrowdsalePausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CappedCrowdsalePausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CappedCrowdsalePaused represents a Paused event raised by the CappedCrowdsale contract.
type CappedCrowdsalePaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_CappedCrowdsale *CappedCrowdsaleFilterer) FilterPaused(opts *bind.FilterOpts) (*CappedCrowdsalePausedIterator, error) {

	logs, sub, err := _CappedCrowdsale.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &CappedCrowdsalePausedIterator{contract: _CappedCrowdsale.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_CappedCrowdsale *CappedCrowdsaleFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *CappedCrowdsalePaused) (event.Subscription, error) {

	logs, sub, err := _CappedCrowdsale.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CappedCrowdsalePaused)
				if err := _CappedCrowdsale.contract.UnpackLog(event, "Paused", log); err != nil {
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

// ParsePaused is a log parse operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_CappedCrowdsale *CappedCrowdsaleFilterer) ParsePaused(log types.Log) (*CappedCrowdsalePaused, error) {
	event := new(CappedCrowdsalePaused)
	if err := _CappedCrowdsale.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CappedCrowdsaleTokensPurchasedIterator is returned from FilterTokensPurchased and is used to iterate over the raw logs and unpacked data for TokensPurchased events raised by the CappedCrowdsale contract.
type CappedCrowdsaleTokensPurchasedIterator struct {
	Event *CappedCrowdsaleTokensPurchased // Event containing the contract specifics and raw log

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
func (it *CappedCrowdsaleTokensPurchasedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CappedCrowdsaleTokensPurchased)
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
		it.Event = new(CappedCrowdsaleTokensPurchased)
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
func (it *CappedCrowdsaleTokensPurchasedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CappedCrowdsaleTokensPurchasedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CappedCrowdsaleTokensPurchased represents a TokensPurchased event raised by the CappedCrowdsale contract.
type CappedCrowdsaleTokensPurchased struct {
	Purchaser   common.Address
	Beneficiary common.Address
	Value       *big.Int
	Amount      *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterTokensPurchased is a free log retrieval operation binding the contract event 0x6faf93231a456e552dbc9961f58d9713ee4f2e69d15f1975b050ef0911053a7b.
//
// Solidity: event TokensPurchased(address indexed purchaser, address indexed beneficiary, uint256 value, uint256 amount)
func (_CappedCrowdsale *CappedCrowdsaleFilterer) FilterTokensPurchased(opts *bind.FilterOpts, purchaser []common.Address, beneficiary []common.Address) (*CappedCrowdsaleTokensPurchasedIterator, error) {

	var purchaserRule []interface{}
	for _, purchaserItem := range purchaser {
		purchaserRule = append(purchaserRule, purchaserItem)
	}
	var beneficiaryRule []interface{}
	for _, beneficiaryItem := range beneficiary {
		beneficiaryRule = append(beneficiaryRule, beneficiaryItem)
	}

	logs, sub, err := _CappedCrowdsale.contract.FilterLogs(opts, "TokensPurchased", purchaserRule, beneficiaryRule)
	if err != nil {
		return nil, err
	}
	return &CappedCrowdsaleTokensPurchasedIterator{contract: _CappedCrowdsale.contract, event: "TokensPurchased", logs: logs, sub: sub}, nil
}

// WatchTokensPurchased is a free log subscription operation binding the contract event 0x6faf93231a456e552dbc9961f58d9713ee4f2e69d15f1975b050ef0911053a7b.
//
// Solidity: event TokensPurchased(address indexed purchaser, address indexed beneficiary, uint256 value, uint256 amount)
func (_CappedCrowdsale *CappedCrowdsaleFilterer) WatchTokensPurchased(opts *bind.WatchOpts, sink chan<- *CappedCrowdsaleTokensPurchased, purchaser []common.Address, beneficiary []common.Address) (event.Subscription, error) {

	var purchaserRule []interface{}
	for _, purchaserItem := range purchaser {
		purchaserRule = append(purchaserRule, purchaserItem)
	}
	var beneficiaryRule []interface{}
	for _, beneficiaryItem := range beneficiary {
		beneficiaryRule = append(beneficiaryRule, beneficiaryItem)
	}

	logs, sub, err := _CappedCrowdsale.contract.WatchLogs(opts, "TokensPurchased", purchaserRule, beneficiaryRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CappedCrowdsaleTokensPurchased)
				if err := _CappedCrowdsale.contract.UnpackLog(event, "TokensPurchased", log); err != nil {
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

// ParseTokensPurchased is a log parse operation binding the contract event 0x6faf93231a456e552dbc9961f58d9713ee4f2e69d15f1975b050ef0911053a7b.
//
// Solidity: event TokensPurchased(address indexed purchaser, address indexed beneficiary, uint256 value, uint256 amount)
func (_CappedCrowdsale *CappedCrowdsaleFilterer) ParseTokensPurchased(log types.Log) (*CappedCrowdsaleTokensPurchased, error) {
	event := new(CappedCrowdsaleTokensPurchased)
	if err := _CappedCrowdsale.contract.UnpackLog(event, "TokensPurchased", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CappedCrowdsaleUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the CappedCrowdsale contract.
type CappedCrowdsaleUnpausedIterator struct {
	Event *CappedCrowdsaleUnpaused // Event containing the contract specifics and raw log

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
func (it *CappedCrowdsaleUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CappedCrowdsaleUnpaused)
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
		it.Event = new(CappedCrowdsaleUnpaused)
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
func (it *CappedCrowdsaleUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CappedCrowdsaleUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CappedCrowdsaleUnpaused represents a Unpaused event raised by the CappedCrowdsale contract.
type CappedCrowdsaleUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_CappedCrowdsale *CappedCrowdsaleFilterer) FilterUnpaused(opts *bind.FilterOpts) (*CappedCrowdsaleUnpausedIterator, error) {

	logs, sub, err := _CappedCrowdsale.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &CappedCrowdsaleUnpausedIterator{contract: _CappedCrowdsale.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_CappedCrowdsale *CappedCrowdsaleFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *CappedCrowdsaleUnpaused) (event.Subscription, error) {

	logs, sub, err := _CappedCrowdsale.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CappedCrowdsaleUnpaused)
				if err := _CappedCrowdsale.contract.UnpackLog(event, "Unpaused", log); err != nil {
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

// ParseUnpaused is a log parse operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_CappedCrowdsale *CappedCrowdsaleFilterer) ParseUnpaused(log types.Log) (*CappedCrowdsaleUnpaused, error) {
	event := new(CappedCrowdsaleUnpaused)
	if err := _CappedCrowdsale.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
