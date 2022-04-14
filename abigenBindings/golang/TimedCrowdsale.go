// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package TimedCrowdsale

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

// TimedCrowdsaleMetaData contains all meta data concerning the TimedCrowdsale contract.
var TimedCrowdsaleMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"prevClosingTime\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newClosingTime\",\"type\":\"uint256\"}],\"name\":\"TimedCrowdsaleExtended\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"purchaser\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"beneficiary\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"TokensPurchased\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"beneficiary\",\"type\":\"address\"}],\"name\":\"buyTokens\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"newReceiver\",\"type\":\"address\"}],\"name\":\"changeWallet\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"_to\",\"type\":\"address\"}],\"name\":\"destroySmartContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token\",\"outputs\":[{\"internalType\":\"contractIBEP20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"wallet\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"weiRaised\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"},{\"inputs\":[],\"name\":\"openingTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"closingTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isOpen\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"hasClosed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// TimedCrowdsaleABI is the input ABI used to generate the binding from.
// Deprecated: Use TimedCrowdsaleMetaData.ABI instead.
var TimedCrowdsaleABI = TimedCrowdsaleMetaData.ABI

// TimedCrowdsale is an auto generated Go binding around an Ethereum contract.
type TimedCrowdsale struct {
	TimedCrowdsaleCaller     // Read-only binding to the contract
	TimedCrowdsaleTransactor // Write-only binding to the contract
	TimedCrowdsaleFilterer   // Log filterer for contract events
}

// TimedCrowdsaleCaller is an auto generated read-only Go binding around an Ethereum contract.
type TimedCrowdsaleCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TimedCrowdsaleTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TimedCrowdsaleTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TimedCrowdsaleFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TimedCrowdsaleFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TimedCrowdsaleSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TimedCrowdsaleSession struct {
	Contract     *TimedCrowdsale   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TimedCrowdsaleCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TimedCrowdsaleCallerSession struct {
	Contract *TimedCrowdsaleCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// TimedCrowdsaleTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TimedCrowdsaleTransactorSession struct {
	Contract     *TimedCrowdsaleTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// TimedCrowdsaleRaw is an auto generated low-level Go binding around an Ethereum contract.
type TimedCrowdsaleRaw struct {
	Contract *TimedCrowdsale // Generic contract binding to access the raw methods on
}

// TimedCrowdsaleCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TimedCrowdsaleCallerRaw struct {
	Contract *TimedCrowdsaleCaller // Generic read-only contract binding to access the raw methods on
}

// TimedCrowdsaleTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TimedCrowdsaleTransactorRaw struct {
	Contract *TimedCrowdsaleTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTimedCrowdsale creates a new instance of TimedCrowdsale, bound to a specific deployed contract.
func NewTimedCrowdsale(address common.Address, backend bind.ContractBackend) (*TimedCrowdsale, error) {
	contract, err := bindTimedCrowdsale(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TimedCrowdsale{TimedCrowdsaleCaller: TimedCrowdsaleCaller{contract: contract}, TimedCrowdsaleTransactor: TimedCrowdsaleTransactor{contract: contract}, TimedCrowdsaleFilterer: TimedCrowdsaleFilterer{contract: contract}}, nil
}

// NewTimedCrowdsaleCaller creates a new read-only instance of TimedCrowdsale, bound to a specific deployed contract.
func NewTimedCrowdsaleCaller(address common.Address, caller bind.ContractCaller) (*TimedCrowdsaleCaller, error) {
	contract, err := bindTimedCrowdsale(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TimedCrowdsaleCaller{contract: contract}, nil
}

// NewTimedCrowdsaleTransactor creates a new write-only instance of TimedCrowdsale, bound to a specific deployed contract.
func NewTimedCrowdsaleTransactor(address common.Address, transactor bind.ContractTransactor) (*TimedCrowdsaleTransactor, error) {
	contract, err := bindTimedCrowdsale(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TimedCrowdsaleTransactor{contract: contract}, nil
}

// NewTimedCrowdsaleFilterer creates a new log filterer instance of TimedCrowdsale, bound to a specific deployed contract.
func NewTimedCrowdsaleFilterer(address common.Address, filterer bind.ContractFilterer) (*TimedCrowdsaleFilterer, error) {
	contract, err := bindTimedCrowdsale(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TimedCrowdsaleFilterer{contract: contract}, nil
}

// bindTimedCrowdsale binds a generic wrapper to an already deployed contract.
func bindTimedCrowdsale(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TimedCrowdsaleABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TimedCrowdsale *TimedCrowdsaleRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TimedCrowdsale.Contract.TimedCrowdsaleCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TimedCrowdsale *TimedCrowdsaleRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TimedCrowdsale.Contract.TimedCrowdsaleTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TimedCrowdsale *TimedCrowdsaleRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TimedCrowdsale.Contract.TimedCrowdsaleTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TimedCrowdsale *TimedCrowdsaleCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TimedCrowdsale.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TimedCrowdsale *TimedCrowdsaleTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TimedCrowdsale.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TimedCrowdsale *TimedCrowdsaleTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TimedCrowdsale.Contract.contract.Transact(opts, method, params...)
}

// ClosingTime is a free data retrieval call binding the contract method 0x4b6753bc.
//
// Solidity: function closingTime() view returns(uint256)
func (_TimedCrowdsale *TimedCrowdsaleCaller) ClosingTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TimedCrowdsale.contract.Call(opts, &out, "closingTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ClosingTime is a free data retrieval call binding the contract method 0x4b6753bc.
//
// Solidity: function closingTime() view returns(uint256)
func (_TimedCrowdsale *TimedCrowdsaleSession) ClosingTime() (*big.Int, error) {
	return _TimedCrowdsale.Contract.ClosingTime(&_TimedCrowdsale.CallOpts)
}

// ClosingTime is a free data retrieval call binding the contract method 0x4b6753bc.
//
// Solidity: function closingTime() view returns(uint256)
func (_TimedCrowdsale *TimedCrowdsaleCallerSession) ClosingTime() (*big.Int, error) {
	return _TimedCrowdsale.Contract.ClosingTime(&_TimedCrowdsale.CallOpts)
}

// HasClosed is a free data retrieval call binding the contract method 0x1515bc2b.
//
// Solidity: function hasClosed() view returns(bool)
func (_TimedCrowdsale *TimedCrowdsaleCaller) HasClosed(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _TimedCrowdsale.contract.Call(opts, &out, "hasClosed")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasClosed is a free data retrieval call binding the contract method 0x1515bc2b.
//
// Solidity: function hasClosed() view returns(bool)
func (_TimedCrowdsale *TimedCrowdsaleSession) HasClosed() (bool, error) {
	return _TimedCrowdsale.Contract.HasClosed(&_TimedCrowdsale.CallOpts)
}

// HasClosed is a free data retrieval call binding the contract method 0x1515bc2b.
//
// Solidity: function hasClosed() view returns(bool)
func (_TimedCrowdsale *TimedCrowdsaleCallerSession) HasClosed() (bool, error) {
	return _TimedCrowdsale.Contract.HasClosed(&_TimedCrowdsale.CallOpts)
}

// IsOpen is a free data retrieval call binding the contract method 0x47535d7b.
//
// Solidity: function isOpen() view returns(bool)
func (_TimedCrowdsale *TimedCrowdsaleCaller) IsOpen(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _TimedCrowdsale.contract.Call(opts, &out, "isOpen")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsOpen is a free data retrieval call binding the contract method 0x47535d7b.
//
// Solidity: function isOpen() view returns(bool)
func (_TimedCrowdsale *TimedCrowdsaleSession) IsOpen() (bool, error) {
	return _TimedCrowdsale.Contract.IsOpen(&_TimedCrowdsale.CallOpts)
}

// IsOpen is a free data retrieval call binding the contract method 0x47535d7b.
//
// Solidity: function isOpen() view returns(bool)
func (_TimedCrowdsale *TimedCrowdsaleCallerSession) IsOpen() (bool, error) {
	return _TimedCrowdsale.Contract.IsOpen(&_TimedCrowdsale.CallOpts)
}

// OpeningTime is a free data retrieval call binding the contract method 0xb7a8807c.
//
// Solidity: function openingTime() view returns(uint256)
func (_TimedCrowdsale *TimedCrowdsaleCaller) OpeningTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TimedCrowdsale.contract.Call(opts, &out, "openingTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// OpeningTime is a free data retrieval call binding the contract method 0xb7a8807c.
//
// Solidity: function openingTime() view returns(uint256)
func (_TimedCrowdsale *TimedCrowdsaleSession) OpeningTime() (*big.Int, error) {
	return _TimedCrowdsale.Contract.OpeningTime(&_TimedCrowdsale.CallOpts)
}

// OpeningTime is a free data retrieval call binding the contract method 0xb7a8807c.
//
// Solidity: function openingTime() view returns(uint256)
func (_TimedCrowdsale *TimedCrowdsaleCallerSession) OpeningTime() (*big.Int, error) {
	return _TimedCrowdsale.Contract.OpeningTime(&_TimedCrowdsale.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_TimedCrowdsale *TimedCrowdsaleCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TimedCrowdsale.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_TimedCrowdsale *TimedCrowdsaleSession) Owner() (common.Address, error) {
	return _TimedCrowdsale.Contract.Owner(&_TimedCrowdsale.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_TimedCrowdsale *TimedCrowdsaleCallerSession) Owner() (common.Address, error) {
	return _TimedCrowdsale.Contract.Owner(&_TimedCrowdsale.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_TimedCrowdsale *TimedCrowdsaleCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _TimedCrowdsale.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_TimedCrowdsale *TimedCrowdsaleSession) Paused() (bool, error) {
	return _TimedCrowdsale.Contract.Paused(&_TimedCrowdsale.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_TimedCrowdsale *TimedCrowdsaleCallerSession) Paused() (bool, error) {
	return _TimedCrowdsale.Contract.Paused(&_TimedCrowdsale.CallOpts)
}

// Rate is a free data retrieval call binding the contract method 0x2c4e722e.
//
// Solidity: function rate() view returns(uint256)
func (_TimedCrowdsale *TimedCrowdsaleCaller) Rate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TimedCrowdsale.contract.Call(opts, &out, "rate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Rate is a free data retrieval call binding the contract method 0x2c4e722e.
//
// Solidity: function rate() view returns(uint256)
func (_TimedCrowdsale *TimedCrowdsaleSession) Rate() (*big.Int, error) {
	return _TimedCrowdsale.Contract.Rate(&_TimedCrowdsale.CallOpts)
}

// Rate is a free data retrieval call binding the contract method 0x2c4e722e.
//
// Solidity: function rate() view returns(uint256)
func (_TimedCrowdsale *TimedCrowdsaleCallerSession) Rate() (*big.Int, error) {
	return _TimedCrowdsale.Contract.Rate(&_TimedCrowdsale.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_TimedCrowdsale *TimedCrowdsaleCaller) Token(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TimedCrowdsale.contract.Call(opts, &out, "token")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_TimedCrowdsale *TimedCrowdsaleSession) Token() (common.Address, error) {
	return _TimedCrowdsale.Contract.Token(&_TimedCrowdsale.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_TimedCrowdsale *TimedCrowdsaleCallerSession) Token() (common.Address, error) {
	return _TimedCrowdsale.Contract.Token(&_TimedCrowdsale.CallOpts)
}

// Wallet is a free data retrieval call binding the contract method 0x521eb273.
//
// Solidity: function wallet() view returns(address)
func (_TimedCrowdsale *TimedCrowdsaleCaller) Wallet(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TimedCrowdsale.contract.Call(opts, &out, "wallet")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Wallet is a free data retrieval call binding the contract method 0x521eb273.
//
// Solidity: function wallet() view returns(address)
func (_TimedCrowdsale *TimedCrowdsaleSession) Wallet() (common.Address, error) {
	return _TimedCrowdsale.Contract.Wallet(&_TimedCrowdsale.CallOpts)
}

// Wallet is a free data retrieval call binding the contract method 0x521eb273.
//
// Solidity: function wallet() view returns(address)
func (_TimedCrowdsale *TimedCrowdsaleCallerSession) Wallet() (common.Address, error) {
	return _TimedCrowdsale.Contract.Wallet(&_TimedCrowdsale.CallOpts)
}

// WeiRaised is a free data retrieval call binding the contract method 0x4042b66f.
//
// Solidity: function weiRaised() view returns(uint256)
func (_TimedCrowdsale *TimedCrowdsaleCaller) WeiRaised(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TimedCrowdsale.contract.Call(opts, &out, "weiRaised")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WeiRaised is a free data retrieval call binding the contract method 0x4042b66f.
//
// Solidity: function weiRaised() view returns(uint256)
func (_TimedCrowdsale *TimedCrowdsaleSession) WeiRaised() (*big.Int, error) {
	return _TimedCrowdsale.Contract.WeiRaised(&_TimedCrowdsale.CallOpts)
}

// WeiRaised is a free data retrieval call binding the contract method 0x4042b66f.
//
// Solidity: function weiRaised() view returns(uint256)
func (_TimedCrowdsale *TimedCrowdsaleCallerSession) WeiRaised() (*big.Int, error) {
	return _TimedCrowdsale.Contract.WeiRaised(&_TimedCrowdsale.CallOpts)
}

// BuyTokens is a paid mutator transaction binding the contract method 0xec8ac4d8.
//
// Solidity: function buyTokens(address beneficiary) payable returns()
func (_TimedCrowdsale *TimedCrowdsaleTransactor) BuyTokens(opts *bind.TransactOpts, beneficiary common.Address) (*types.Transaction, error) {
	return _TimedCrowdsale.contract.Transact(opts, "buyTokens", beneficiary)
}

// BuyTokens is a paid mutator transaction binding the contract method 0xec8ac4d8.
//
// Solidity: function buyTokens(address beneficiary) payable returns()
func (_TimedCrowdsale *TimedCrowdsaleSession) BuyTokens(beneficiary common.Address) (*types.Transaction, error) {
	return _TimedCrowdsale.Contract.BuyTokens(&_TimedCrowdsale.TransactOpts, beneficiary)
}

// BuyTokens is a paid mutator transaction binding the contract method 0xec8ac4d8.
//
// Solidity: function buyTokens(address beneficiary) payable returns()
func (_TimedCrowdsale *TimedCrowdsaleTransactorSession) BuyTokens(beneficiary common.Address) (*types.Transaction, error) {
	return _TimedCrowdsale.Contract.BuyTokens(&_TimedCrowdsale.TransactOpts, beneficiary)
}

// ChangeWallet is a paid mutator transaction binding the contract method 0x98b9a2dc.
//
// Solidity: function changeWallet(address newReceiver) returns()
func (_TimedCrowdsale *TimedCrowdsaleTransactor) ChangeWallet(opts *bind.TransactOpts, newReceiver common.Address) (*types.Transaction, error) {
	return _TimedCrowdsale.contract.Transact(opts, "changeWallet", newReceiver)
}

// ChangeWallet is a paid mutator transaction binding the contract method 0x98b9a2dc.
//
// Solidity: function changeWallet(address newReceiver) returns()
func (_TimedCrowdsale *TimedCrowdsaleSession) ChangeWallet(newReceiver common.Address) (*types.Transaction, error) {
	return _TimedCrowdsale.Contract.ChangeWallet(&_TimedCrowdsale.TransactOpts, newReceiver)
}

// ChangeWallet is a paid mutator transaction binding the contract method 0x98b9a2dc.
//
// Solidity: function changeWallet(address newReceiver) returns()
func (_TimedCrowdsale *TimedCrowdsaleTransactorSession) ChangeWallet(newReceiver common.Address) (*types.Transaction, error) {
	return _TimedCrowdsale.Contract.ChangeWallet(&_TimedCrowdsale.TransactOpts, newReceiver)
}

// DestroySmartContract is a paid mutator transaction binding the contract method 0x39df43ff.
//
// Solidity: function destroySmartContract(address _to) returns()
func (_TimedCrowdsale *TimedCrowdsaleTransactor) DestroySmartContract(opts *bind.TransactOpts, _to common.Address) (*types.Transaction, error) {
	return _TimedCrowdsale.contract.Transact(opts, "destroySmartContract", _to)
}

// DestroySmartContract is a paid mutator transaction binding the contract method 0x39df43ff.
//
// Solidity: function destroySmartContract(address _to) returns()
func (_TimedCrowdsale *TimedCrowdsaleSession) DestroySmartContract(_to common.Address) (*types.Transaction, error) {
	return _TimedCrowdsale.Contract.DestroySmartContract(&_TimedCrowdsale.TransactOpts, _to)
}

// DestroySmartContract is a paid mutator transaction binding the contract method 0x39df43ff.
//
// Solidity: function destroySmartContract(address _to) returns()
func (_TimedCrowdsale *TimedCrowdsaleTransactorSession) DestroySmartContract(_to common.Address) (*types.Transaction, error) {
	return _TimedCrowdsale.Contract.DestroySmartContract(&_TimedCrowdsale.TransactOpts, _to)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_TimedCrowdsale *TimedCrowdsaleTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TimedCrowdsale.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_TimedCrowdsale *TimedCrowdsaleSession) Pause() (*types.Transaction, error) {
	return _TimedCrowdsale.Contract.Pause(&_TimedCrowdsale.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_TimedCrowdsale *TimedCrowdsaleTransactorSession) Pause() (*types.Transaction, error) {
	return _TimedCrowdsale.Contract.Pause(&_TimedCrowdsale.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_TimedCrowdsale *TimedCrowdsaleTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TimedCrowdsale.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_TimedCrowdsale *TimedCrowdsaleSession) RenounceOwnership() (*types.Transaction, error) {
	return _TimedCrowdsale.Contract.RenounceOwnership(&_TimedCrowdsale.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_TimedCrowdsale *TimedCrowdsaleTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _TimedCrowdsale.Contract.RenounceOwnership(&_TimedCrowdsale.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_TimedCrowdsale *TimedCrowdsaleTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _TimedCrowdsale.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_TimedCrowdsale *TimedCrowdsaleSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _TimedCrowdsale.Contract.TransferOwnership(&_TimedCrowdsale.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_TimedCrowdsale *TimedCrowdsaleTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _TimedCrowdsale.Contract.TransferOwnership(&_TimedCrowdsale.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_TimedCrowdsale *TimedCrowdsaleTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TimedCrowdsale.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_TimedCrowdsale *TimedCrowdsaleSession) Unpause() (*types.Transaction, error) {
	return _TimedCrowdsale.Contract.Unpause(&_TimedCrowdsale.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_TimedCrowdsale *TimedCrowdsaleTransactorSession) Unpause() (*types.Transaction, error) {
	return _TimedCrowdsale.Contract.Unpause(&_TimedCrowdsale.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_TimedCrowdsale *TimedCrowdsaleTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TimedCrowdsale.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_TimedCrowdsale *TimedCrowdsaleSession) Receive() (*types.Transaction, error) {
	return _TimedCrowdsale.Contract.Receive(&_TimedCrowdsale.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_TimedCrowdsale *TimedCrowdsaleTransactorSession) Receive() (*types.Transaction, error) {
	return _TimedCrowdsale.Contract.Receive(&_TimedCrowdsale.TransactOpts)
}

// TimedCrowdsaleOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the TimedCrowdsale contract.
type TimedCrowdsaleOwnershipTransferredIterator struct {
	Event *TimedCrowdsaleOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *TimedCrowdsaleOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TimedCrowdsaleOwnershipTransferred)
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
		it.Event = new(TimedCrowdsaleOwnershipTransferred)
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
func (it *TimedCrowdsaleOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TimedCrowdsaleOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TimedCrowdsaleOwnershipTransferred represents a OwnershipTransferred event raised by the TimedCrowdsale contract.
type TimedCrowdsaleOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_TimedCrowdsale *TimedCrowdsaleFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*TimedCrowdsaleOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _TimedCrowdsale.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &TimedCrowdsaleOwnershipTransferredIterator{contract: _TimedCrowdsale.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_TimedCrowdsale *TimedCrowdsaleFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *TimedCrowdsaleOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _TimedCrowdsale.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TimedCrowdsaleOwnershipTransferred)
				if err := _TimedCrowdsale.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_TimedCrowdsale *TimedCrowdsaleFilterer) ParseOwnershipTransferred(log types.Log) (*TimedCrowdsaleOwnershipTransferred, error) {
	event := new(TimedCrowdsaleOwnershipTransferred)
	if err := _TimedCrowdsale.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TimedCrowdsalePausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the TimedCrowdsale contract.
type TimedCrowdsalePausedIterator struct {
	Event *TimedCrowdsalePaused // Event containing the contract specifics and raw log

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
func (it *TimedCrowdsalePausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TimedCrowdsalePaused)
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
		it.Event = new(TimedCrowdsalePaused)
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
func (it *TimedCrowdsalePausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TimedCrowdsalePausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TimedCrowdsalePaused represents a Paused event raised by the TimedCrowdsale contract.
type TimedCrowdsalePaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_TimedCrowdsale *TimedCrowdsaleFilterer) FilterPaused(opts *bind.FilterOpts) (*TimedCrowdsalePausedIterator, error) {

	logs, sub, err := _TimedCrowdsale.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &TimedCrowdsalePausedIterator{contract: _TimedCrowdsale.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_TimedCrowdsale *TimedCrowdsaleFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *TimedCrowdsalePaused) (event.Subscription, error) {

	logs, sub, err := _TimedCrowdsale.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TimedCrowdsalePaused)
				if err := _TimedCrowdsale.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_TimedCrowdsale *TimedCrowdsaleFilterer) ParsePaused(log types.Log) (*TimedCrowdsalePaused, error) {
	event := new(TimedCrowdsalePaused)
	if err := _TimedCrowdsale.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TimedCrowdsaleTimedCrowdsaleExtendedIterator is returned from FilterTimedCrowdsaleExtended and is used to iterate over the raw logs and unpacked data for TimedCrowdsaleExtended events raised by the TimedCrowdsale contract.
type TimedCrowdsaleTimedCrowdsaleExtendedIterator struct {
	Event *TimedCrowdsaleTimedCrowdsaleExtended // Event containing the contract specifics and raw log

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
func (it *TimedCrowdsaleTimedCrowdsaleExtendedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TimedCrowdsaleTimedCrowdsaleExtended)
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
		it.Event = new(TimedCrowdsaleTimedCrowdsaleExtended)
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
func (it *TimedCrowdsaleTimedCrowdsaleExtendedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TimedCrowdsaleTimedCrowdsaleExtendedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TimedCrowdsaleTimedCrowdsaleExtended represents a TimedCrowdsaleExtended event raised by the TimedCrowdsale contract.
type TimedCrowdsaleTimedCrowdsaleExtended struct {
	PrevClosingTime *big.Int
	NewClosingTime  *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterTimedCrowdsaleExtended is a free log retrieval operation binding the contract event 0x46711e222f558a07afd26e5e71b48ecb0a8b2cdcd40faeb1323e05e2c76a2f32.
//
// Solidity: event TimedCrowdsaleExtended(uint256 prevClosingTime, uint256 newClosingTime)
func (_TimedCrowdsale *TimedCrowdsaleFilterer) FilterTimedCrowdsaleExtended(opts *bind.FilterOpts) (*TimedCrowdsaleTimedCrowdsaleExtendedIterator, error) {

	logs, sub, err := _TimedCrowdsale.contract.FilterLogs(opts, "TimedCrowdsaleExtended")
	if err != nil {
		return nil, err
	}
	return &TimedCrowdsaleTimedCrowdsaleExtendedIterator{contract: _TimedCrowdsale.contract, event: "TimedCrowdsaleExtended", logs: logs, sub: sub}, nil
}

// WatchTimedCrowdsaleExtended is a free log subscription operation binding the contract event 0x46711e222f558a07afd26e5e71b48ecb0a8b2cdcd40faeb1323e05e2c76a2f32.
//
// Solidity: event TimedCrowdsaleExtended(uint256 prevClosingTime, uint256 newClosingTime)
func (_TimedCrowdsale *TimedCrowdsaleFilterer) WatchTimedCrowdsaleExtended(opts *bind.WatchOpts, sink chan<- *TimedCrowdsaleTimedCrowdsaleExtended) (event.Subscription, error) {

	logs, sub, err := _TimedCrowdsale.contract.WatchLogs(opts, "TimedCrowdsaleExtended")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TimedCrowdsaleTimedCrowdsaleExtended)
				if err := _TimedCrowdsale.contract.UnpackLog(event, "TimedCrowdsaleExtended", log); err != nil {
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
func (_TimedCrowdsale *TimedCrowdsaleFilterer) ParseTimedCrowdsaleExtended(log types.Log) (*TimedCrowdsaleTimedCrowdsaleExtended, error) {
	event := new(TimedCrowdsaleTimedCrowdsaleExtended)
	if err := _TimedCrowdsale.contract.UnpackLog(event, "TimedCrowdsaleExtended", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TimedCrowdsaleTokensPurchasedIterator is returned from FilterTokensPurchased and is used to iterate over the raw logs and unpacked data for TokensPurchased events raised by the TimedCrowdsale contract.
type TimedCrowdsaleTokensPurchasedIterator struct {
	Event *TimedCrowdsaleTokensPurchased // Event containing the contract specifics and raw log

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
func (it *TimedCrowdsaleTokensPurchasedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TimedCrowdsaleTokensPurchased)
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
		it.Event = new(TimedCrowdsaleTokensPurchased)
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
func (it *TimedCrowdsaleTokensPurchasedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TimedCrowdsaleTokensPurchasedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TimedCrowdsaleTokensPurchased represents a TokensPurchased event raised by the TimedCrowdsale contract.
type TimedCrowdsaleTokensPurchased struct {
	Purchaser   common.Address
	Beneficiary common.Address
	Value       *big.Int
	Amount      *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterTokensPurchased is a free log retrieval operation binding the contract event 0x6faf93231a456e552dbc9961f58d9713ee4f2e69d15f1975b050ef0911053a7b.
//
// Solidity: event TokensPurchased(address indexed purchaser, address indexed beneficiary, uint256 value, uint256 amount)
func (_TimedCrowdsale *TimedCrowdsaleFilterer) FilterTokensPurchased(opts *bind.FilterOpts, purchaser []common.Address, beneficiary []common.Address) (*TimedCrowdsaleTokensPurchasedIterator, error) {

	var purchaserRule []interface{}
	for _, purchaserItem := range purchaser {
		purchaserRule = append(purchaserRule, purchaserItem)
	}
	var beneficiaryRule []interface{}
	for _, beneficiaryItem := range beneficiary {
		beneficiaryRule = append(beneficiaryRule, beneficiaryItem)
	}

	logs, sub, err := _TimedCrowdsale.contract.FilterLogs(opts, "TokensPurchased", purchaserRule, beneficiaryRule)
	if err != nil {
		return nil, err
	}
	return &TimedCrowdsaleTokensPurchasedIterator{contract: _TimedCrowdsale.contract, event: "TokensPurchased", logs: logs, sub: sub}, nil
}

// WatchTokensPurchased is a free log subscription operation binding the contract event 0x6faf93231a456e552dbc9961f58d9713ee4f2e69d15f1975b050ef0911053a7b.
//
// Solidity: event TokensPurchased(address indexed purchaser, address indexed beneficiary, uint256 value, uint256 amount)
func (_TimedCrowdsale *TimedCrowdsaleFilterer) WatchTokensPurchased(opts *bind.WatchOpts, sink chan<- *TimedCrowdsaleTokensPurchased, purchaser []common.Address, beneficiary []common.Address) (event.Subscription, error) {

	var purchaserRule []interface{}
	for _, purchaserItem := range purchaser {
		purchaserRule = append(purchaserRule, purchaserItem)
	}
	var beneficiaryRule []interface{}
	for _, beneficiaryItem := range beneficiary {
		beneficiaryRule = append(beneficiaryRule, beneficiaryItem)
	}

	logs, sub, err := _TimedCrowdsale.contract.WatchLogs(opts, "TokensPurchased", purchaserRule, beneficiaryRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TimedCrowdsaleTokensPurchased)
				if err := _TimedCrowdsale.contract.UnpackLog(event, "TokensPurchased", log); err != nil {
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
func (_TimedCrowdsale *TimedCrowdsaleFilterer) ParseTokensPurchased(log types.Log) (*TimedCrowdsaleTokensPurchased, error) {
	event := new(TimedCrowdsaleTokensPurchased)
	if err := _TimedCrowdsale.contract.UnpackLog(event, "TokensPurchased", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TimedCrowdsaleUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the TimedCrowdsale contract.
type TimedCrowdsaleUnpausedIterator struct {
	Event *TimedCrowdsaleUnpaused // Event containing the contract specifics and raw log

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
func (it *TimedCrowdsaleUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TimedCrowdsaleUnpaused)
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
		it.Event = new(TimedCrowdsaleUnpaused)
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
func (it *TimedCrowdsaleUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TimedCrowdsaleUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TimedCrowdsaleUnpaused represents a Unpaused event raised by the TimedCrowdsale contract.
type TimedCrowdsaleUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_TimedCrowdsale *TimedCrowdsaleFilterer) FilterUnpaused(opts *bind.FilterOpts) (*TimedCrowdsaleUnpausedIterator, error) {

	logs, sub, err := _TimedCrowdsale.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &TimedCrowdsaleUnpausedIterator{contract: _TimedCrowdsale.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_TimedCrowdsale *TimedCrowdsaleFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *TimedCrowdsaleUnpaused) (event.Subscription, error) {

	logs, sub, err := _TimedCrowdsale.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TimedCrowdsaleUnpaused)
				if err := _TimedCrowdsale.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
func (_TimedCrowdsale *TimedCrowdsaleFilterer) ParseUnpaused(log types.Log) (*TimedCrowdsaleUnpaused, error) {
	event := new(TimedCrowdsaleUnpaused)
	if err := _TimedCrowdsale.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
