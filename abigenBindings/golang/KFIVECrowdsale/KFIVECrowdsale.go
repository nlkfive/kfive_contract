// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package KFIVECrowdsale

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

// KFIVECrowdsaleMetaData contains all meta data concerning the KFIVECrowdsale contract.
var KFIVECrowdsaleMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"rate\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"wallet\",\"type\":\"address\"},{\"internalType\":\"contractIBEP20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"cap\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"openingTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"closingTime\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"prevClosingTime\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newClosingTime\",\"type\":\"uint256\"}],\"name\":\"TimedCrowdsaleExtended\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"purchaser\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"beneficiary\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"TokensPurchased\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"beneficiary\",\"type\":\"address\"}],\"name\":\"buyTokens\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"cap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"capReached\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"newReceiver\",\"type\":\"address\"}],\"name\":\"changeWallet\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"closingTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"_to\",\"type\":\"address\"}],\"name\":\"destroySmartContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"hasClosed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isOpen\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"openingTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token\",\"outputs\":[{\"internalType\":\"contractIBEP20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"wallet\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"weiRaised\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// KFIVECrowdsaleABI is the input ABI used to generate the binding from.
// Deprecated: Use KFIVECrowdsaleMetaData.ABI instead.
var KFIVECrowdsaleABI = KFIVECrowdsaleMetaData.ABI

// KFIVECrowdsale is an auto generated Go binding around an Ethereum contract.
type KFIVECrowdsale struct {
	KFIVECrowdsaleCaller     // Read-only binding to the contract
	KFIVECrowdsaleTransactor // Write-only binding to the contract
	KFIVECrowdsaleFilterer   // Log filterer for contract events
}

// KFIVECrowdsaleCaller is an auto generated read-only Go binding around an Ethereum contract.
type KFIVECrowdsaleCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KFIVECrowdsaleTransactor is an auto generated write-only Go binding around an Ethereum contract.
type KFIVECrowdsaleTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KFIVECrowdsaleFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type KFIVECrowdsaleFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KFIVECrowdsaleSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type KFIVECrowdsaleSession struct {
	Contract     *KFIVECrowdsale   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// KFIVECrowdsaleCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type KFIVECrowdsaleCallerSession struct {
	Contract *KFIVECrowdsaleCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// KFIVECrowdsaleTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type KFIVECrowdsaleTransactorSession struct {
	Contract     *KFIVECrowdsaleTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// KFIVECrowdsaleRaw is an auto generated low-level Go binding around an Ethereum contract.
type KFIVECrowdsaleRaw struct {
	Contract *KFIVECrowdsale // Generic contract binding to access the raw methods on
}

// KFIVECrowdsaleCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type KFIVECrowdsaleCallerRaw struct {
	Contract *KFIVECrowdsaleCaller // Generic read-only contract binding to access the raw methods on
}

// KFIVECrowdsaleTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type KFIVECrowdsaleTransactorRaw struct {
	Contract *KFIVECrowdsaleTransactor // Generic write-only contract binding to access the raw methods on
}

// NewKFIVECrowdsale creates a new instance of KFIVECrowdsale, bound to a specific deployed contract.
func NewKFIVECrowdsale(address common.Address, backend bind.ContractBackend) (*KFIVECrowdsale, error) {
	contract, err := bindKFIVECrowdsale(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &KFIVECrowdsale{KFIVECrowdsaleCaller: KFIVECrowdsaleCaller{contract: contract}, KFIVECrowdsaleTransactor: KFIVECrowdsaleTransactor{contract: contract}, KFIVECrowdsaleFilterer: KFIVECrowdsaleFilterer{contract: contract}}, nil
}

// NewKFIVECrowdsaleCaller creates a new read-only instance of KFIVECrowdsale, bound to a specific deployed contract.
func NewKFIVECrowdsaleCaller(address common.Address, caller bind.ContractCaller) (*KFIVECrowdsaleCaller, error) {
	contract, err := bindKFIVECrowdsale(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &KFIVECrowdsaleCaller{contract: contract}, nil
}

// NewKFIVECrowdsaleTransactor creates a new write-only instance of KFIVECrowdsale, bound to a specific deployed contract.
func NewKFIVECrowdsaleTransactor(address common.Address, transactor bind.ContractTransactor) (*KFIVECrowdsaleTransactor, error) {
	contract, err := bindKFIVECrowdsale(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &KFIVECrowdsaleTransactor{contract: contract}, nil
}

// NewKFIVECrowdsaleFilterer creates a new log filterer instance of KFIVECrowdsale, bound to a specific deployed contract.
func NewKFIVECrowdsaleFilterer(address common.Address, filterer bind.ContractFilterer) (*KFIVECrowdsaleFilterer, error) {
	contract, err := bindKFIVECrowdsale(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &KFIVECrowdsaleFilterer{contract: contract}, nil
}

// bindKFIVECrowdsale binds a generic wrapper to an already deployed contract.
func bindKFIVECrowdsale(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(KFIVECrowdsaleABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_KFIVECrowdsale *KFIVECrowdsaleRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _KFIVECrowdsale.Contract.KFIVECrowdsaleCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_KFIVECrowdsale *KFIVECrowdsaleRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _KFIVECrowdsale.Contract.KFIVECrowdsaleTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_KFIVECrowdsale *KFIVECrowdsaleRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _KFIVECrowdsale.Contract.KFIVECrowdsaleTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_KFIVECrowdsale *KFIVECrowdsaleCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _KFIVECrowdsale.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_KFIVECrowdsale *KFIVECrowdsaleTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _KFIVECrowdsale.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_KFIVECrowdsale *KFIVECrowdsaleTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _KFIVECrowdsale.Contract.contract.Transact(opts, method, params...)
}

// Cap is a free data retrieval call binding the contract method 0x355274ea.
//
// Solidity: function cap() view returns(uint256)
func (_KFIVECrowdsale *KFIVECrowdsaleCaller) Cap(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _KFIVECrowdsale.contract.Call(opts, &out, "cap")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Cap is a free data retrieval call binding the contract method 0x355274ea.
//
// Solidity: function cap() view returns(uint256)
func (_KFIVECrowdsale *KFIVECrowdsaleSession) Cap() (*big.Int, error) {
	return _KFIVECrowdsale.Contract.Cap(&_KFIVECrowdsale.CallOpts)
}

// Cap is a free data retrieval call binding the contract method 0x355274ea.
//
// Solidity: function cap() view returns(uint256)
func (_KFIVECrowdsale *KFIVECrowdsaleCallerSession) Cap() (*big.Int, error) {
	return _KFIVECrowdsale.Contract.Cap(&_KFIVECrowdsale.CallOpts)
}

// CapReached is a free data retrieval call binding the contract method 0x4f935945.
//
// Solidity: function capReached() view returns(bool)
func (_KFIVECrowdsale *KFIVECrowdsaleCaller) CapReached(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _KFIVECrowdsale.contract.Call(opts, &out, "capReached")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CapReached is a free data retrieval call binding the contract method 0x4f935945.
//
// Solidity: function capReached() view returns(bool)
func (_KFIVECrowdsale *KFIVECrowdsaleSession) CapReached() (bool, error) {
	return _KFIVECrowdsale.Contract.CapReached(&_KFIVECrowdsale.CallOpts)
}

// CapReached is a free data retrieval call binding the contract method 0x4f935945.
//
// Solidity: function capReached() view returns(bool)
func (_KFIVECrowdsale *KFIVECrowdsaleCallerSession) CapReached() (bool, error) {
	return _KFIVECrowdsale.Contract.CapReached(&_KFIVECrowdsale.CallOpts)
}

// ClosingTime is a free data retrieval call binding the contract method 0x4b6753bc.
//
// Solidity: function closingTime() view returns(uint256)
func (_KFIVECrowdsale *KFIVECrowdsaleCaller) ClosingTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _KFIVECrowdsale.contract.Call(opts, &out, "closingTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ClosingTime is a free data retrieval call binding the contract method 0x4b6753bc.
//
// Solidity: function closingTime() view returns(uint256)
func (_KFIVECrowdsale *KFIVECrowdsaleSession) ClosingTime() (*big.Int, error) {
	return _KFIVECrowdsale.Contract.ClosingTime(&_KFIVECrowdsale.CallOpts)
}

// ClosingTime is a free data retrieval call binding the contract method 0x4b6753bc.
//
// Solidity: function closingTime() view returns(uint256)
func (_KFIVECrowdsale *KFIVECrowdsaleCallerSession) ClosingTime() (*big.Int, error) {
	return _KFIVECrowdsale.Contract.ClosingTime(&_KFIVECrowdsale.CallOpts)
}

// HasClosed is a free data retrieval call binding the contract method 0x1515bc2b.
//
// Solidity: function hasClosed() view returns(bool)
func (_KFIVECrowdsale *KFIVECrowdsaleCaller) HasClosed(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _KFIVECrowdsale.contract.Call(opts, &out, "hasClosed")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasClosed is a free data retrieval call binding the contract method 0x1515bc2b.
//
// Solidity: function hasClosed() view returns(bool)
func (_KFIVECrowdsale *KFIVECrowdsaleSession) HasClosed() (bool, error) {
	return _KFIVECrowdsale.Contract.HasClosed(&_KFIVECrowdsale.CallOpts)
}

// HasClosed is a free data retrieval call binding the contract method 0x1515bc2b.
//
// Solidity: function hasClosed() view returns(bool)
func (_KFIVECrowdsale *KFIVECrowdsaleCallerSession) HasClosed() (bool, error) {
	return _KFIVECrowdsale.Contract.HasClosed(&_KFIVECrowdsale.CallOpts)
}

// IsOpen is a free data retrieval call binding the contract method 0x47535d7b.
//
// Solidity: function isOpen() view returns(bool)
func (_KFIVECrowdsale *KFIVECrowdsaleCaller) IsOpen(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _KFIVECrowdsale.contract.Call(opts, &out, "isOpen")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsOpen is a free data retrieval call binding the contract method 0x47535d7b.
//
// Solidity: function isOpen() view returns(bool)
func (_KFIVECrowdsale *KFIVECrowdsaleSession) IsOpen() (bool, error) {
	return _KFIVECrowdsale.Contract.IsOpen(&_KFIVECrowdsale.CallOpts)
}

// IsOpen is a free data retrieval call binding the contract method 0x47535d7b.
//
// Solidity: function isOpen() view returns(bool)
func (_KFIVECrowdsale *KFIVECrowdsaleCallerSession) IsOpen() (bool, error) {
	return _KFIVECrowdsale.Contract.IsOpen(&_KFIVECrowdsale.CallOpts)
}

// OpeningTime is a free data retrieval call binding the contract method 0xb7a8807c.
//
// Solidity: function openingTime() view returns(uint256)
func (_KFIVECrowdsale *KFIVECrowdsaleCaller) OpeningTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _KFIVECrowdsale.contract.Call(opts, &out, "openingTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// OpeningTime is a free data retrieval call binding the contract method 0xb7a8807c.
//
// Solidity: function openingTime() view returns(uint256)
func (_KFIVECrowdsale *KFIVECrowdsaleSession) OpeningTime() (*big.Int, error) {
	return _KFIVECrowdsale.Contract.OpeningTime(&_KFIVECrowdsale.CallOpts)
}

// OpeningTime is a free data retrieval call binding the contract method 0xb7a8807c.
//
// Solidity: function openingTime() view returns(uint256)
func (_KFIVECrowdsale *KFIVECrowdsaleCallerSession) OpeningTime() (*big.Int, error) {
	return _KFIVECrowdsale.Contract.OpeningTime(&_KFIVECrowdsale.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_KFIVECrowdsale *KFIVECrowdsaleCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _KFIVECrowdsale.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_KFIVECrowdsale *KFIVECrowdsaleSession) Owner() (common.Address, error) {
	return _KFIVECrowdsale.Contract.Owner(&_KFIVECrowdsale.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_KFIVECrowdsale *KFIVECrowdsaleCallerSession) Owner() (common.Address, error) {
	return _KFIVECrowdsale.Contract.Owner(&_KFIVECrowdsale.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_KFIVECrowdsale *KFIVECrowdsaleCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _KFIVECrowdsale.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_KFIVECrowdsale *KFIVECrowdsaleSession) Paused() (bool, error) {
	return _KFIVECrowdsale.Contract.Paused(&_KFIVECrowdsale.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_KFIVECrowdsale *KFIVECrowdsaleCallerSession) Paused() (bool, error) {
	return _KFIVECrowdsale.Contract.Paused(&_KFIVECrowdsale.CallOpts)
}

// Rate is a free data retrieval call binding the contract method 0x2c4e722e.
//
// Solidity: function rate() view returns(uint256)
func (_KFIVECrowdsale *KFIVECrowdsaleCaller) Rate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _KFIVECrowdsale.contract.Call(opts, &out, "rate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Rate is a free data retrieval call binding the contract method 0x2c4e722e.
//
// Solidity: function rate() view returns(uint256)
func (_KFIVECrowdsale *KFIVECrowdsaleSession) Rate() (*big.Int, error) {
	return _KFIVECrowdsale.Contract.Rate(&_KFIVECrowdsale.CallOpts)
}

// Rate is a free data retrieval call binding the contract method 0x2c4e722e.
//
// Solidity: function rate() view returns(uint256)
func (_KFIVECrowdsale *KFIVECrowdsaleCallerSession) Rate() (*big.Int, error) {
	return _KFIVECrowdsale.Contract.Rate(&_KFIVECrowdsale.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_KFIVECrowdsale *KFIVECrowdsaleCaller) Token(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _KFIVECrowdsale.contract.Call(opts, &out, "token")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_KFIVECrowdsale *KFIVECrowdsaleSession) Token() (common.Address, error) {
	return _KFIVECrowdsale.Contract.Token(&_KFIVECrowdsale.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_KFIVECrowdsale *KFIVECrowdsaleCallerSession) Token() (common.Address, error) {
	return _KFIVECrowdsale.Contract.Token(&_KFIVECrowdsale.CallOpts)
}

// Wallet is a free data retrieval call binding the contract method 0x521eb273.
//
// Solidity: function wallet() view returns(address)
func (_KFIVECrowdsale *KFIVECrowdsaleCaller) Wallet(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _KFIVECrowdsale.contract.Call(opts, &out, "wallet")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Wallet is a free data retrieval call binding the contract method 0x521eb273.
//
// Solidity: function wallet() view returns(address)
func (_KFIVECrowdsale *KFIVECrowdsaleSession) Wallet() (common.Address, error) {
	return _KFIVECrowdsale.Contract.Wallet(&_KFIVECrowdsale.CallOpts)
}

// Wallet is a free data retrieval call binding the contract method 0x521eb273.
//
// Solidity: function wallet() view returns(address)
func (_KFIVECrowdsale *KFIVECrowdsaleCallerSession) Wallet() (common.Address, error) {
	return _KFIVECrowdsale.Contract.Wallet(&_KFIVECrowdsale.CallOpts)
}

// WeiRaised is a free data retrieval call binding the contract method 0x4042b66f.
//
// Solidity: function weiRaised() view returns(uint256)
func (_KFIVECrowdsale *KFIVECrowdsaleCaller) WeiRaised(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _KFIVECrowdsale.contract.Call(opts, &out, "weiRaised")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WeiRaised is a free data retrieval call binding the contract method 0x4042b66f.
//
// Solidity: function weiRaised() view returns(uint256)
func (_KFIVECrowdsale *KFIVECrowdsaleSession) WeiRaised() (*big.Int, error) {
	return _KFIVECrowdsale.Contract.WeiRaised(&_KFIVECrowdsale.CallOpts)
}

// WeiRaised is a free data retrieval call binding the contract method 0x4042b66f.
//
// Solidity: function weiRaised() view returns(uint256)
func (_KFIVECrowdsale *KFIVECrowdsaleCallerSession) WeiRaised() (*big.Int, error) {
	return _KFIVECrowdsale.Contract.WeiRaised(&_KFIVECrowdsale.CallOpts)
}

// BuyTokens is a paid mutator transaction binding the contract method 0xec8ac4d8.
//
// Solidity: function buyTokens(address beneficiary) payable returns()
func (_KFIVECrowdsale *KFIVECrowdsaleTransactor) BuyTokens(opts *bind.TransactOpts, beneficiary common.Address) (*types.Transaction, error) {
	return _KFIVECrowdsale.contract.Transact(opts, "buyTokens", beneficiary)
}

// BuyTokens is a paid mutator transaction binding the contract method 0xec8ac4d8.
//
// Solidity: function buyTokens(address beneficiary) payable returns()
func (_KFIVECrowdsale *KFIVECrowdsaleSession) BuyTokens(beneficiary common.Address) (*types.Transaction, error) {
	return _KFIVECrowdsale.Contract.BuyTokens(&_KFIVECrowdsale.TransactOpts, beneficiary)
}

// BuyTokens is a paid mutator transaction binding the contract method 0xec8ac4d8.
//
// Solidity: function buyTokens(address beneficiary) payable returns()
func (_KFIVECrowdsale *KFIVECrowdsaleTransactorSession) BuyTokens(beneficiary common.Address) (*types.Transaction, error) {
	return _KFIVECrowdsale.Contract.BuyTokens(&_KFIVECrowdsale.TransactOpts, beneficiary)
}

// ChangeWallet is a paid mutator transaction binding the contract method 0x98b9a2dc.
//
// Solidity: function changeWallet(address newReceiver) returns()
func (_KFIVECrowdsale *KFIVECrowdsaleTransactor) ChangeWallet(opts *bind.TransactOpts, newReceiver common.Address) (*types.Transaction, error) {
	return _KFIVECrowdsale.contract.Transact(opts, "changeWallet", newReceiver)
}

// ChangeWallet is a paid mutator transaction binding the contract method 0x98b9a2dc.
//
// Solidity: function changeWallet(address newReceiver) returns()
func (_KFIVECrowdsale *KFIVECrowdsaleSession) ChangeWallet(newReceiver common.Address) (*types.Transaction, error) {
	return _KFIVECrowdsale.Contract.ChangeWallet(&_KFIVECrowdsale.TransactOpts, newReceiver)
}

// ChangeWallet is a paid mutator transaction binding the contract method 0x98b9a2dc.
//
// Solidity: function changeWallet(address newReceiver) returns()
func (_KFIVECrowdsale *KFIVECrowdsaleTransactorSession) ChangeWallet(newReceiver common.Address) (*types.Transaction, error) {
	return _KFIVECrowdsale.Contract.ChangeWallet(&_KFIVECrowdsale.TransactOpts, newReceiver)
}

// DestroySmartContract is a paid mutator transaction binding the contract method 0x39df43ff.
//
// Solidity: function destroySmartContract(address _to) returns()
func (_KFIVECrowdsale *KFIVECrowdsaleTransactor) DestroySmartContract(opts *bind.TransactOpts, _to common.Address) (*types.Transaction, error) {
	return _KFIVECrowdsale.contract.Transact(opts, "destroySmartContract", _to)
}

// DestroySmartContract is a paid mutator transaction binding the contract method 0x39df43ff.
//
// Solidity: function destroySmartContract(address _to) returns()
func (_KFIVECrowdsale *KFIVECrowdsaleSession) DestroySmartContract(_to common.Address) (*types.Transaction, error) {
	return _KFIVECrowdsale.Contract.DestroySmartContract(&_KFIVECrowdsale.TransactOpts, _to)
}

// DestroySmartContract is a paid mutator transaction binding the contract method 0x39df43ff.
//
// Solidity: function destroySmartContract(address _to) returns()
func (_KFIVECrowdsale *KFIVECrowdsaleTransactorSession) DestroySmartContract(_to common.Address) (*types.Transaction, error) {
	return _KFIVECrowdsale.Contract.DestroySmartContract(&_KFIVECrowdsale.TransactOpts, _to)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_KFIVECrowdsale *KFIVECrowdsaleTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _KFIVECrowdsale.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_KFIVECrowdsale *KFIVECrowdsaleSession) Pause() (*types.Transaction, error) {
	return _KFIVECrowdsale.Contract.Pause(&_KFIVECrowdsale.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_KFIVECrowdsale *KFIVECrowdsaleTransactorSession) Pause() (*types.Transaction, error) {
	return _KFIVECrowdsale.Contract.Pause(&_KFIVECrowdsale.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_KFIVECrowdsale *KFIVECrowdsaleTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _KFIVECrowdsale.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_KFIVECrowdsale *KFIVECrowdsaleSession) RenounceOwnership() (*types.Transaction, error) {
	return _KFIVECrowdsale.Contract.RenounceOwnership(&_KFIVECrowdsale.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_KFIVECrowdsale *KFIVECrowdsaleTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _KFIVECrowdsale.Contract.RenounceOwnership(&_KFIVECrowdsale.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_KFIVECrowdsale *KFIVECrowdsaleTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _KFIVECrowdsale.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_KFIVECrowdsale *KFIVECrowdsaleSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _KFIVECrowdsale.Contract.TransferOwnership(&_KFIVECrowdsale.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_KFIVECrowdsale *KFIVECrowdsaleTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _KFIVECrowdsale.Contract.TransferOwnership(&_KFIVECrowdsale.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_KFIVECrowdsale *KFIVECrowdsaleTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _KFIVECrowdsale.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_KFIVECrowdsale *KFIVECrowdsaleSession) Unpause() (*types.Transaction, error) {
	return _KFIVECrowdsale.Contract.Unpause(&_KFIVECrowdsale.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_KFIVECrowdsale *KFIVECrowdsaleTransactorSession) Unpause() (*types.Transaction, error) {
	return _KFIVECrowdsale.Contract.Unpause(&_KFIVECrowdsale.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_KFIVECrowdsale *KFIVECrowdsaleTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _KFIVECrowdsale.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_KFIVECrowdsale *KFIVECrowdsaleSession) Receive() (*types.Transaction, error) {
	return _KFIVECrowdsale.Contract.Receive(&_KFIVECrowdsale.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_KFIVECrowdsale *KFIVECrowdsaleTransactorSession) Receive() (*types.Transaction, error) {
	return _KFIVECrowdsale.Contract.Receive(&_KFIVECrowdsale.TransactOpts)
}

// KFIVECrowdsaleOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the KFIVECrowdsale contract.
type KFIVECrowdsaleOwnershipTransferredIterator struct {
	Event *KFIVECrowdsaleOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *KFIVECrowdsaleOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KFIVECrowdsaleOwnershipTransferred)
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
		it.Event = new(KFIVECrowdsaleOwnershipTransferred)
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
func (it *KFIVECrowdsaleOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KFIVECrowdsaleOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KFIVECrowdsaleOwnershipTransferred represents a OwnershipTransferred event raised by the KFIVECrowdsale contract.
type KFIVECrowdsaleOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_KFIVECrowdsale *KFIVECrowdsaleFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*KFIVECrowdsaleOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _KFIVECrowdsale.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &KFIVECrowdsaleOwnershipTransferredIterator{contract: _KFIVECrowdsale.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_KFIVECrowdsale *KFIVECrowdsaleFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *KFIVECrowdsaleOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _KFIVECrowdsale.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KFIVECrowdsaleOwnershipTransferred)
				if err := _KFIVECrowdsale.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_KFIVECrowdsale *KFIVECrowdsaleFilterer) ParseOwnershipTransferred(log types.Log) (*KFIVECrowdsaleOwnershipTransferred, error) {
	event := new(KFIVECrowdsaleOwnershipTransferred)
	if err := _KFIVECrowdsale.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// KFIVECrowdsalePausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the KFIVECrowdsale contract.
type KFIVECrowdsalePausedIterator struct {
	Event *KFIVECrowdsalePaused // Event containing the contract specifics and raw log

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
func (it *KFIVECrowdsalePausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KFIVECrowdsalePaused)
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
		it.Event = new(KFIVECrowdsalePaused)
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
func (it *KFIVECrowdsalePausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KFIVECrowdsalePausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KFIVECrowdsalePaused represents a Paused event raised by the KFIVECrowdsale contract.
type KFIVECrowdsalePaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_KFIVECrowdsale *KFIVECrowdsaleFilterer) FilterPaused(opts *bind.FilterOpts) (*KFIVECrowdsalePausedIterator, error) {

	logs, sub, err := _KFIVECrowdsale.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &KFIVECrowdsalePausedIterator{contract: _KFIVECrowdsale.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_KFIVECrowdsale *KFIVECrowdsaleFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *KFIVECrowdsalePaused) (event.Subscription, error) {

	logs, sub, err := _KFIVECrowdsale.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KFIVECrowdsalePaused)
				if err := _KFIVECrowdsale.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_KFIVECrowdsale *KFIVECrowdsaleFilterer) ParsePaused(log types.Log) (*KFIVECrowdsalePaused, error) {
	event := new(KFIVECrowdsalePaused)
	if err := _KFIVECrowdsale.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// KFIVECrowdsaleTimedCrowdsaleExtendedIterator is returned from FilterTimedCrowdsaleExtended and is used to iterate over the raw logs and unpacked data for TimedCrowdsaleExtended events raised by the KFIVECrowdsale contract.
type KFIVECrowdsaleTimedCrowdsaleExtendedIterator struct {
	Event *KFIVECrowdsaleTimedCrowdsaleExtended // Event containing the contract specifics and raw log

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
func (it *KFIVECrowdsaleTimedCrowdsaleExtendedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KFIVECrowdsaleTimedCrowdsaleExtended)
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
		it.Event = new(KFIVECrowdsaleTimedCrowdsaleExtended)
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
func (it *KFIVECrowdsaleTimedCrowdsaleExtendedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KFIVECrowdsaleTimedCrowdsaleExtendedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KFIVECrowdsaleTimedCrowdsaleExtended represents a TimedCrowdsaleExtended event raised by the KFIVECrowdsale contract.
type KFIVECrowdsaleTimedCrowdsaleExtended struct {
	PrevClosingTime *big.Int
	NewClosingTime  *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterTimedCrowdsaleExtended is a free log retrieval operation binding the contract event 0x46711e222f558a07afd26e5e71b48ecb0a8b2cdcd40faeb1323e05e2c76a2f32.
//
// Solidity: event TimedCrowdsaleExtended(uint256 prevClosingTime, uint256 newClosingTime)
func (_KFIVECrowdsale *KFIVECrowdsaleFilterer) FilterTimedCrowdsaleExtended(opts *bind.FilterOpts) (*KFIVECrowdsaleTimedCrowdsaleExtendedIterator, error) {

	logs, sub, err := _KFIVECrowdsale.contract.FilterLogs(opts, "TimedCrowdsaleExtended")
	if err != nil {
		return nil, err
	}
	return &KFIVECrowdsaleTimedCrowdsaleExtendedIterator{contract: _KFIVECrowdsale.contract, event: "TimedCrowdsaleExtended", logs: logs, sub: sub}, nil
}

// WatchTimedCrowdsaleExtended is a free log subscription operation binding the contract event 0x46711e222f558a07afd26e5e71b48ecb0a8b2cdcd40faeb1323e05e2c76a2f32.
//
// Solidity: event TimedCrowdsaleExtended(uint256 prevClosingTime, uint256 newClosingTime)
func (_KFIVECrowdsale *KFIVECrowdsaleFilterer) WatchTimedCrowdsaleExtended(opts *bind.WatchOpts, sink chan<- *KFIVECrowdsaleTimedCrowdsaleExtended) (event.Subscription, error) {

	logs, sub, err := _KFIVECrowdsale.contract.WatchLogs(opts, "TimedCrowdsaleExtended")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KFIVECrowdsaleTimedCrowdsaleExtended)
				if err := _KFIVECrowdsale.contract.UnpackLog(event, "TimedCrowdsaleExtended", log); err != nil {
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

// ParseTimedCrowdsaleExtended is a log parse operation binding the contract event 0x46711e222f558a07afd26e5e71b48ecb0a8b2cdcd40faeb1323e05e2c76a2f32.
//
// Solidity: event TimedCrowdsaleExtended(uint256 prevClosingTime, uint256 newClosingTime)
func (_KFIVECrowdsale *KFIVECrowdsaleFilterer) ParseTimedCrowdsaleExtended(log types.Log) (*KFIVECrowdsaleTimedCrowdsaleExtended, error) {
	event := new(KFIVECrowdsaleTimedCrowdsaleExtended)
	if err := _KFIVECrowdsale.contract.UnpackLog(event, "TimedCrowdsaleExtended", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// KFIVECrowdsaleTokensPurchasedIterator is returned from FilterTokensPurchased and is used to iterate over the raw logs and unpacked data for TokensPurchased events raised by the KFIVECrowdsale contract.
type KFIVECrowdsaleTokensPurchasedIterator struct {
	Event *KFIVECrowdsaleTokensPurchased // Event containing the contract specifics and raw log

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
func (it *KFIVECrowdsaleTokensPurchasedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KFIVECrowdsaleTokensPurchased)
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
		it.Event = new(KFIVECrowdsaleTokensPurchased)
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
func (it *KFIVECrowdsaleTokensPurchasedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KFIVECrowdsaleTokensPurchasedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KFIVECrowdsaleTokensPurchased represents a TokensPurchased event raised by the KFIVECrowdsale contract.
type KFIVECrowdsaleTokensPurchased struct {
	Purchaser   common.Address
	Beneficiary common.Address
	Value       *big.Int
	Amount      *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterTokensPurchased is a free log retrieval operation binding the contract event 0x6faf93231a456e552dbc9961f58d9713ee4f2e69d15f1975b050ef0911053a7b.
//
// Solidity: event TokensPurchased(address indexed purchaser, address indexed beneficiary, uint256 value, uint256 amount)
func (_KFIVECrowdsale *KFIVECrowdsaleFilterer) FilterTokensPurchased(opts *bind.FilterOpts, purchaser []common.Address, beneficiary []common.Address) (*KFIVECrowdsaleTokensPurchasedIterator, error) {

	var purchaserRule []interface{}
	for _, purchaserItem := range purchaser {
		purchaserRule = append(purchaserRule, purchaserItem)
	}
	var beneficiaryRule []interface{}
	for _, beneficiaryItem := range beneficiary {
		beneficiaryRule = append(beneficiaryRule, beneficiaryItem)
	}

	logs, sub, err := _KFIVECrowdsale.contract.FilterLogs(opts, "TokensPurchased", purchaserRule, beneficiaryRule)
	if err != nil {
		return nil, err
	}
	return &KFIVECrowdsaleTokensPurchasedIterator{contract: _KFIVECrowdsale.contract, event: "TokensPurchased", logs: logs, sub: sub}, nil
}

// WatchTokensPurchased is a free log subscription operation binding the contract event 0x6faf93231a456e552dbc9961f58d9713ee4f2e69d15f1975b050ef0911053a7b.
//
// Solidity: event TokensPurchased(address indexed purchaser, address indexed beneficiary, uint256 value, uint256 amount)
func (_KFIVECrowdsale *KFIVECrowdsaleFilterer) WatchTokensPurchased(opts *bind.WatchOpts, sink chan<- *KFIVECrowdsaleTokensPurchased, purchaser []common.Address, beneficiary []common.Address) (event.Subscription, error) {

	var purchaserRule []interface{}
	for _, purchaserItem := range purchaser {
		purchaserRule = append(purchaserRule, purchaserItem)
	}
	var beneficiaryRule []interface{}
	for _, beneficiaryItem := range beneficiary {
		beneficiaryRule = append(beneficiaryRule, beneficiaryItem)
	}

	logs, sub, err := _KFIVECrowdsale.contract.WatchLogs(opts, "TokensPurchased", purchaserRule, beneficiaryRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KFIVECrowdsaleTokensPurchased)
				if err := _KFIVECrowdsale.contract.UnpackLog(event, "TokensPurchased", log); err != nil {
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
func (_KFIVECrowdsale *KFIVECrowdsaleFilterer) ParseTokensPurchased(log types.Log) (*KFIVECrowdsaleTokensPurchased, error) {
	event := new(KFIVECrowdsaleTokensPurchased)
	if err := _KFIVECrowdsale.contract.UnpackLog(event, "TokensPurchased", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// KFIVECrowdsaleUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the KFIVECrowdsale contract.
type KFIVECrowdsaleUnpausedIterator struct {
	Event *KFIVECrowdsaleUnpaused // Event containing the contract specifics and raw log

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
func (it *KFIVECrowdsaleUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KFIVECrowdsaleUnpaused)
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
		it.Event = new(KFIVECrowdsaleUnpaused)
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
func (it *KFIVECrowdsaleUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KFIVECrowdsaleUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KFIVECrowdsaleUnpaused represents a Unpaused event raised by the KFIVECrowdsale contract.
type KFIVECrowdsaleUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_KFIVECrowdsale *KFIVECrowdsaleFilterer) FilterUnpaused(opts *bind.FilterOpts) (*KFIVECrowdsaleUnpausedIterator, error) {

	logs, sub, err := _KFIVECrowdsale.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &KFIVECrowdsaleUnpausedIterator{contract: _KFIVECrowdsale.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_KFIVECrowdsale *KFIVECrowdsaleFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *KFIVECrowdsaleUnpaused) (event.Subscription, error) {

	logs, sub, err := _KFIVECrowdsale.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KFIVECrowdsaleUnpaused)
				if err := _KFIVECrowdsale.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
func (_KFIVECrowdsale *KFIVECrowdsaleFilterer) ParseUnpaused(log types.Log) (*KFIVECrowdsaleUnpaused, error) {
	event := new(KFIVECrowdsaleUnpaused)
	if err := _KFIVECrowdsale.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
