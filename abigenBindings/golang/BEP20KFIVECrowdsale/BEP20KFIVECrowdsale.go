// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package BEP20KFIVECrowdsale

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

// BEP20KFIVECrowdsaleMetaData contains all meta data concerning the BEP20KFIVECrowdsale contract.
var BEP20KFIVECrowdsaleMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"rate\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"wallet\",\"type\":\"address\"},{\"internalType\":\"contractIBEP20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"contractIBEP20\",\"name\":\"acceptToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"cap\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"openingTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"closingTime\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"prevClosingTime\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newClosingTime\",\"type\":\"uint256\"}],\"name\":\"TimedCrowdsaleExtended\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"purchaser\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"beneficiary\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"TokensPurchased\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"beneficiary\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"buyTokens\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"cap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"capReached\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"newReceiver\",\"type\":\"address\"}],\"name\":\"changeWallet\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"closingTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"_to\",\"type\":\"address\"}],\"name\":\"destroySmartContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"hasClosed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isOpen\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"openingTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token\",\"outputs\":[{\"internalType\":\"contractIBEP20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"wallet\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"weiRaised\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// BEP20KFIVECrowdsaleABI is the input ABI used to generate the binding from.
// Deprecated: Use BEP20KFIVECrowdsaleMetaData.ABI instead.
var BEP20KFIVECrowdsaleABI = BEP20KFIVECrowdsaleMetaData.ABI

// BEP20KFIVECrowdsale is an auto generated Go binding around an Ethereum contract.
type BEP20KFIVECrowdsale struct {
	BEP20KFIVECrowdsaleCaller     // Read-only binding to the contract
	BEP20KFIVECrowdsaleTransactor // Write-only binding to the contract
	BEP20KFIVECrowdsaleFilterer   // Log filterer for contract events
}

// BEP20KFIVECrowdsaleCaller is an auto generated read-only Go binding around an Ethereum contract.
type BEP20KFIVECrowdsaleCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BEP20KFIVECrowdsaleTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BEP20KFIVECrowdsaleTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BEP20KFIVECrowdsaleFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BEP20KFIVECrowdsaleFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BEP20KFIVECrowdsaleSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BEP20KFIVECrowdsaleSession struct {
	Contract     *BEP20KFIVECrowdsale // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// BEP20KFIVECrowdsaleCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BEP20KFIVECrowdsaleCallerSession struct {
	Contract *BEP20KFIVECrowdsaleCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// BEP20KFIVECrowdsaleTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BEP20KFIVECrowdsaleTransactorSession struct {
	Contract     *BEP20KFIVECrowdsaleTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// BEP20KFIVECrowdsaleRaw is an auto generated low-level Go binding around an Ethereum contract.
type BEP20KFIVECrowdsaleRaw struct {
	Contract *BEP20KFIVECrowdsale // Generic contract binding to access the raw methods on
}

// BEP20KFIVECrowdsaleCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BEP20KFIVECrowdsaleCallerRaw struct {
	Contract *BEP20KFIVECrowdsaleCaller // Generic read-only contract binding to access the raw methods on
}

// BEP20KFIVECrowdsaleTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BEP20KFIVECrowdsaleTransactorRaw struct {
	Contract *BEP20KFIVECrowdsaleTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBEP20KFIVECrowdsale creates a new instance of BEP20KFIVECrowdsale, bound to a specific deployed contract.
func NewBEP20KFIVECrowdsale(address common.Address, backend bind.ContractBackend) (*BEP20KFIVECrowdsale, error) {
	contract, err := bindBEP20KFIVECrowdsale(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BEP20KFIVECrowdsale{BEP20KFIVECrowdsaleCaller: BEP20KFIVECrowdsaleCaller{contract: contract}, BEP20KFIVECrowdsaleTransactor: BEP20KFIVECrowdsaleTransactor{contract: contract}, BEP20KFIVECrowdsaleFilterer: BEP20KFIVECrowdsaleFilterer{contract: contract}}, nil
}

// NewBEP20KFIVECrowdsaleCaller creates a new read-only instance of BEP20KFIVECrowdsale, bound to a specific deployed contract.
func NewBEP20KFIVECrowdsaleCaller(address common.Address, caller bind.ContractCaller) (*BEP20KFIVECrowdsaleCaller, error) {
	contract, err := bindBEP20KFIVECrowdsale(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BEP20KFIVECrowdsaleCaller{contract: contract}, nil
}

// NewBEP20KFIVECrowdsaleTransactor creates a new write-only instance of BEP20KFIVECrowdsale, bound to a specific deployed contract.
func NewBEP20KFIVECrowdsaleTransactor(address common.Address, transactor bind.ContractTransactor) (*BEP20KFIVECrowdsaleTransactor, error) {
	contract, err := bindBEP20KFIVECrowdsale(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BEP20KFIVECrowdsaleTransactor{contract: contract}, nil
}

// NewBEP20KFIVECrowdsaleFilterer creates a new log filterer instance of BEP20KFIVECrowdsale, bound to a specific deployed contract.
func NewBEP20KFIVECrowdsaleFilterer(address common.Address, filterer bind.ContractFilterer) (*BEP20KFIVECrowdsaleFilterer, error) {
	contract, err := bindBEP20KFIVECrowdsale(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BEP20KFIVECrowdsaleFilterer{contract: contract}, nil
}

// bindBEP20KFIVECrowdsale binds a generic wrapper to an already deployed contract.
func bindBEP20KFIVECrowdsale(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BEP20KFIVECrowdsaleABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BEP20KFIVECrowdsale *BEP20KFIVECrowdsaleRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BEP20KFIVECrowdsale.Contract.BEP20KFIVECrowdsaleCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BEP20KFIVECrowdsale *BEP20KFIVECrowdsaleRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BEP20KFIVECrowdsale.Contract.BEP20KFIVECrowdsaleTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BEP20KFIVECrowdsale *BEP20KFIVECrowdsaleRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BEP20KFIVECrowdsale.Contract.BEP20KFIVECrowdsaleTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BEP20KFIVECrowdsale *BEP20KFIVECrowdsaleCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BEP20KFIVECrowdsale.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BEP20KFIVECrowdsale *BEP20KFIVECrowdsaleTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BEP20KFIVECrowdsale.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BEP20KFIVECrowdsale *BEP20KFIVECrowdsaleTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BEP20KFIVECrowdsale.Contract.contract.Transact(opts, method, params...)
}

// Cap is a free data retrieval call binding the contract method 0x355274ea.
//
// Solidity: function cap() view returns(uint256)
func (_BEP20KFIVECrowdsale *BEP20KFIVECrowdsaleCaller) Cap(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BEP20KFIVECrowdsale.contract.Call(opts, &out, "cap")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Cap is a free data retrieval call binding the contract method 0x355274ea.
//
// Solidity: function cap() view returns(uint256)
func (_BEP20KFIVECrowdsale *BEP20KFIVECrowdsaleSession) Cap() (*big.Int, error) {
	return _BEP20KFIVECrowdsale.Contract.Cap(&_BEP20KFIVECrowdsale.CallOpts)
}

// Cap is a free data retrieval call binding the contract method 0x355274ea.
//
// Solidity: function cap() view returns(uint256)
func (_BEP20KFIVECrowdsale *BEP20KFIVECrowdsaleCallerSession) Cap() (*big.Int, error) {
	return _BEP20KFIVECrowdsale.Contract.Cap(&_BEP20KFIVECrowdsale.CallOpts)
}

// CapReached is a free data retrieval call binding the contract method 0x4f935945.
//
// Solidity: function capReached() view returns(bool)
func (_BEP20KFIVECrowdsale *BEP20KFIVECrowdsaleCaller) CapReached(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _BEP20KFIVECrowdsale.contract.Call(opts, &out, "capReached")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CapReached is a free data retrieval call binding the contract method 0x4f935945.
//
// Solidity: function capReached() view returns(bool)
func (_BEP20KFIVECrowdsale *BEP20KFIVECrowdsaleSession) CapReached() (bool, error) {
	return _BEP20KFIVECrowdsale.Contract.CapReached(&_BEP20KFIVECrowdsale.CallOpts)
}

// CapReached is a free data retrieval call binding the contract method 0x4f935945.
//
// Solidity: function capReached() view returns(bool)
func (_BEP20KFIVECrowdsale *BEP20KFIVECrowdsaleCallerSession) CapReached() (bool, error) {
	return _BEP20KFIVECrowdsale.Contract.CapReached(&_BEP20KFIVECrowdsale.CallOpts)
}

// ClosingTime is a free data retrieval call binding the contract method 0x4b6753bc.
//
// Solidity: function closingTime() view returns(uint256)
func (_BEP20KFIVECrowdsale *BEP20KFIVECrowdsaleCaller) ClosingTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BEP20KFIVECrowdsale.contract.Call(opts, &out, "closingTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ClosingTime is a free data retrieval call binding the contract method 0x4b6753bc.
//
// Solidity: function closingTime() view returns(uint256)
func (_BEP20KFIVECrowdsale *BEP20KFIVECrowdsaleSession) ClosingTime() (*big.Int, error) {
	return _BEP20KFIVECrowdsale.Contract.ClosingTime(&_BEP20KFIVECrowdsale.CallOpts)
}

// ClosingTime is a free data retrieval call binding the contract method 0x4b6753bc.
//
// Solidity: function closingTime() view returns(uint256)
func (_BEP20KFIVECrowdsale *BEP20KFIVECrowdsaleCallerSession) ClosingTime() (*big.Int, error) {
	return _BEP20KFIVECrowdsale.Contract.ClosingTime(&_BEP20KFIVECrowdsale.CallOpts)
}

// HasClosed is a free data retrieval call binding the contract method 0x1515bc2b.
//
// Solidity: function hasClosed() view returns(bool)
func (_BEP20KFIVECrowdsale *BEP20KFIVECrowdsaleCaller) HasClosed(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _BEP20KFIVECrowdsale.contract.Call(opts, &out, "hasClosed")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasClosed is a free data retrieval call binding the contract method 0x1515bc2b.
//
// Solidity: function hasClosed() view returns(bool)
func (_BEP20KFIVECrowdsale *BEP20KFIVECrowdsaleSession) HasClosed() (bool, error) {
	return _BEP20KFIVECrowdsale.Contract.HasClosed(&_BEP20KFIVECrowdsale.CallOpts)
}

// HasClosed is a free data retrieval call binding the contract method 0x1515bc2b.
//
// Solidity: function hasClosed() view returns(bool)
func (_BEP20KFIVECrowdsale *BEP20KFIVECrowdsaleCallerSession) HasClosed() (bool, error) {
	return _BEP20KFIVECrowdsale.Contract.HasClosed(&_BEP20KFIVECrowdsale.CallOpts)
}

// IsOpen is a free data retrieval call binding the contract method 0x47535d7b.
//
// Solidity: function isOpen() view returns(bool)
func (_BEP20KFIVECrowdsale *BEP20KFIVECrowdsaleCaller) IsOpen(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _BEP20KFIVECrowdsale.contract.Call(opts, &out, "isOpen")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsOpen is a free data retrieval call binding the contract method 0x47535d7b.
//
// Solidity: function isOpen() view returns(bool)
func (_BEP20KFIVECrowdsale *BEP20KFIVECrowdsaleSession) IsOpen() (bool, error) {
	return _BEP20KFIVECrowdsale.Contract.IsOpen(&_BEP20KFIVECrowdsale.CallOpts)
}

// IsOpen is a free data retrieval call binding the contract method 0x47535d7b.
//
// Solidity: function isOpen() view returns(bool)
func (_BEP20KFIVECrowdsale *BEP20KFIVECrowdsaleCallerSession) IsOpen() (bool, error) {
	return _BEP20KFIVECrowdsale.Contract.IsOpen(&_BEP20KFIVECrowdsale.CallOpts)
}

// OpeningTime is a free data retrieval call binding the contract method 0xb7a8807c.
//
// Solidity: function openingTime() view returns(uint256)
func (_BEP20KFIVECrowdsale *BEP20KFIVECrowdsaleCaller) OpeningTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BEP20KFIVECrowdsale.contract.Call(opts, &out, "openingTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// OpeningTime is a free data retrieval call binding the contract method 0xb7a8807c.
//
// Solidity: function openingTime() view returns(uint256)
func (_BEP20KFIVECrowdsale *BEP20KFIVECrowdsaleSession) OpeningTime() (*big.Int, error) {
	return _BEP20KFIVECrowdsale.Contract.OpeningTime(&_BEP20KFIVECrowdsale.CallOpts)
}

// OpeningTime is a free data retrieval call binding the contract method 0xb7a8807c.
//
// Solidity: function openingTime() view returns(uint256)
func (_BEP20KFIVECrowdsale *BEP20KFIVECrowdsaleCallerSession) OpeningTime() (*big.Int, error) {
	return _BEP20KFIVECrowdsale.Contract.OpeningTime(&_BEP20KFIVECrowdsale.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BEP20KFIVECrowdsale *BEP20KFIVECrowdsaleCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BEP20KFIVECrowdsale.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BEP20KFIVECrowdsale *BEP20KFIVECrowdsaleSession) Owner() (common.Address, error) {
	return _BEP20KFIVECrowdsale.Contract.Owner(&_BEP20KFIVECrowdsale.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BEP20KFIVECrowdsale *BEP20KFIVECrowdsaleCallerSession) Owner() (common.Address, error) {
	return _BEP20KFIVECrowdsale.Contract.Owner(&_BEP20KFIVECrowdsale.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_BEP20KFIVECrowdsale *BEP20KFIVECrowdsaleCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _BEP20KFIVECrowdsale.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_BEP20KFIVECrowdsale *BEP20KFIVECrowdsaleSession) Paused() (bool, error) {
	return _BEP20KFIVECrowdsale.Contract.Paused(&_BEP20KFIVECrowdsale.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_BEP20KFIVECrowdsale *BEP20KFIVECrowdsaleCallerSession) Paused() (bool, error) {
	return _BEP20KFIVECrowdsale.Contract.Paused(&_BEP20KFIVECrowdsale.CallOpts)
}

// Rate is a free data retrieval call binding the contract method 0x2c4e722e.
//
// Solidity: function rate() view returns(uint256)
func (_BEP20KFIVECrowdsale *BEP20KFIVECrowdsaleCaller) Rate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BEP20KFIVECrowdsale.contract.Call(opts, &out, "rate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Rate is a free data retrieval call binding the contract method 0x2c4e722e.
//
// Solidity: function rate() view returns(uint256)
func (_BEP20KFIVECrowdsale *BEP20KFIVECrowdsaleSession) Rate() (*big.Int, error) {
	return _BEP20KFIVECrowdsale.Contract.Rate(&_BEP20KFIVECrowdsale.CallOpts)
}

// Rate is a free data retrieval call binding the contract method 0x2c4e722e.
//
// Solidity: function rate() view returns(uint256)
func (_BEP20KFIVECrowdsale *BEP20KFIVECrowdsaleCallerSession) Rate() (*big.Int, error) {
	return _BEP20KFIVECrowdsale.Contract.Rate(&_BEP20KFIVECrowdsale.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_BEP20KFIVECrowdsale *BEP20KFIVECrowdsaleCaller) Token(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BEP20KFIVECrowdsale.contract.Call(opts, &out, "token")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_BEP20KFIVECrowdsale *BEP20KFIVECrowdsaleSession) Token() (common.Address, error) {
	return _BEP20KFIVECrowdsale.Contract.Token(&_BEP20KFIVECrowdsale.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_BEP20KFIVECrowdsale *BEP20KFIVECrowdsaleCallerSession) Token() (common.Address, error) {
	return _BEP20KFIVECrowdsale.Contract.Token(&_BEP20KFIVECrowdsale.CallOpts)
}

// Wallet is a free data retrieval call binding the contract method 0x521eb273.
//
// Solidity: function wallet() view returns(address)
func (_BEP20KFIVECrowdsale *BEP20KFIVECrowdsaleCaller) Wallet(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BEP20KFIVECrowdsale.contract.Call(opts, &out, "wallet")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Wallet is a free data retrieval call binding the contract method 0x521eb273.
//
// Solidity: function wallet() view returns(address)
func (_BEP20KFIVECrowdsale *BEP20KFIVECrowdsaleSession) Wallet() (common.Address, error) {
	return _BEP20KFIVECrowdsale.Contract.Wallet(&_BEP20KFIVECrowdsale.CallOpts)
}

// Wallet is a free data retrieval call binding the contract method 0x521eb273.
//
// Solidity: function wallet() view returns(address)
func (_BEP20KFIVECrowdsale *BEP20KFIVECrowdsaleCallerSession) Wallet() (common.Address, error) {
	return _BEP20KFIVECrowdsale.Contract.Wallet(&_BEP20KFIVECrowdsale.CallOpts)
}

// WeiRaised is a free data retrieval call binding the contract method 0x4042b66f.
//
// Solidity: function weiRaised() view returns(uint256)
func (_BEP20KFIVECrowdsale *BEP20KFIVECrowdsaleCaller) WeiRaised(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BEP20KFIVECrowdsale.contract.Call(opts, &out, "weiRaised")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WeiRaised is a free data retrieval call binding the contract method 0x4042b66f.
//
// Solidity: function weiRaised() view returns(uint256)
func (_BEP20KFIVECrowdsale *BEP20KFIVECrowdsaleSession) WeiRaised() (*big.Int, error) {
	return _BEP20KFIVECrowdsale.Contract.WeiRaised(&_BEP20KFIVECrowdsale.CallOpts)
}

// WeiRaised is a free data retrieval call binding the contract method 0x4042b66f.
//
// Solidity: function weiRaised() view returns(uint256)
func (_BEP20KFIVECrowdsale *BEP20KFIVECrowdsaleCallerSession) WeiRaised() (*big.Int, error) {
	return _BEP20KFIVECrowdsale.Contract.WeiRaised(&_BEP20KFIVECrowdsale.CallOpts)
}

// BuyTokens is a paid mutator transaction binding the contract method 0x0752881a.
//
// Solidity: function buyTokens(address beneficiary, uint256 value) payable returns()
func (_BEP20KFIVECrowdsale *BEP20KFIVECrowdsaleTransactor) BuyTokens(opts *bind.TransactOpts, beneficiary common.Address, value *big.Int) (*types.Transaction, error) {
	return _BEP20KFIVECrowdsale.contract.Transact(opts, "buyTokens", beneficiary, value)
}

// BuyTokens is a paid mutator transaction binding the contract method 0x0752881a.
//
// Solidity: function buyTokens(address beneficiary, uint256 value) payable returns()
func (_BEP20KFIVECrowdsale *BEP20KFIVECrowdsaleSession) BuyTokens(beneficiary common.Address, value *big.Int) (*types.Transaction, error) {
	return _BEP20KFIVECrowdsale.Contract.BuyTokens(&_BEP20KFIVECrowdsale.TransactOpts, beneficiary, value)
}

// BuyTokens is a paid mutator transaction binding the contract method 0x0752881a.
//
// Solidity: function buyTokens(address beneficiary, uint256 value) payable returns()
func (_BEP20KFIVECrowdsale *BEP20KFIVECrowdsaleTransactorSession) BuyTokens(beneficiary common.Address, value *big.Int) (*types.Transaction, error) {
	return _BEP20KFIVECrowdsale.Contract.BuyTokens(&_BEP20KFIVECrowdsale.TransactOpts, beneficiary, value)
}

// ChangeWallet is a paid mutator transaction binding the contract method 0x98b9a2dc.
//
// Solidity: function changeWallet(address newReceiver) returns()
func (_BEP20KFIVECrowdsale *BEP20KFIVECrowdsaleTransactor) ChangeWallet(opts *bind.TransactOpts, newReceiver common.Address) (*types.Transaction, error) {
	return _BEP20KFIVECrowdsale.contract.Transact(opts, "changeWallet", newReceiver)
}

// ChangeWallet is a paid mutator transaction binding the contract method 0x98b9a2dc.
//
// Solidity: function changeWallet(address newReceiver) returns()
func (_BEP20KFIVECrowdsale *BEP20KFIVECrowdsaleSession) ChangeWallet(newReceiver common.Address) (*types.Transaction, error) {
	return _BEP20KFIVECrowdsale.Contract.ChangeWallet(&_BEP20KFIVECrowdsale.TransactOpts, newReceiver)
}

// ChangeWallet is a paid mutator transaction binding the contract method 0x98b9a2dc.
//
// Solidity: function changeWallet(address newReceiver) returns()
func (_BEP20KFIVECrowdsale *BEP20KFIVECrowdsaleTransactorSession) ChangeWallet(newReceiver common.Address) (*types.Transaction, error) {
	return _BEP20KFIVECrowdsale.Contract.ChangeWallet(&_BEP20KFIVECrowdsale.TransactOpts, newReceiver)
}

// DestroySmartContract is a paid mutator transaction binding the contract method 0x39df43ff.
//
// Solidity: function destroySmartContract(address _to) returns()
func (_BEP20KFIVECrowdsale *BEP20KFIVECrowdsaleTransactor) DestroySmartContract(opts *bind.TransactOpts, _to common.Address) (*types.Transaction, error) {
	return _BEP20KFIVECrowdsale.contract.Transact(opts, "destroySmartContract", _to)
}

// DestroySmartContract is a paid mutator transaction binding the contract method 0x39df43ff.
//
// Solidity: function destroySmartContract(address _to) returns()
func (_BEP20KFIVECrowdsale *BEP20KFIVECrowdsaleSession) DestroySmartContract(_to common.Address) (*types.Transaction, error) {
	return _BEP20KFIVECrowdsale.Contract.DestroySmartContract(&_BEP20KFIVECrowdsale.TransactOpts, _to)
}

// DestroySmartContract is a paid mutator transaction binding the contract method 0x39df43ff.
//
// Solidity: function destroySmartContract(address _to) returns()
func (_BEP20KFIVECrowdsale *BEP20KFIVECrowdsaleTransactorSession) DestroySmartContract(_to common.Address) (*types.Transaction, error) {
	return _BEP20KFIVECrowdsale.Contract.DestroySmartContract(&_BEP20KFIVECrowdsale.TransactOpts, _to)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_BEP20KFIVECrowdsale *BEP20KFIVECrowdsaleTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BEP20KFIVECrowdsale.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_BEP20KFIVECrowdsale *BEP20KFIVECrowdsaleSession) Pause() (*types.Transaction, error) {
	return _BEP20KFIVECrowdsale.Contract.Pause(&_BEP20KFIVECrowdsale.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_BEP20KFIVECrowdsale *BEP20KFIVECrowdsaleTransactorSession) Pause() (*types.Transaction, error) {
	return _BEP20KFIVECrowdsale.Contract.Pause(&_BEP20KFIVECrowdsale.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_BEP20KFIVECrowdsale *BEP20KFIVECrowdsaleTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BEP20KFIVECrowdsale.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_BEP20KFIVECrowdsale *BEP20KFIVECrowdsaleSession) RenounceOwnership() (*types.Transaction, error) {
	return _BEP20KFIVECrowdsale.Contract.RenounceOwnership(&_BEP20KFIVECrowdsale.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_BEP20KFIVECrowdsale *BEP20KFIVECrowdsaleTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _BEP20KFIVECrowdsale.Contract.RenounceOwnership(&_BEP20KFIVECrowdsale.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_BEP20KFIVECrowdsale *BEP20KFIVECrowdsaleTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _BEP20KFIVECrowdsale.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_BEP20KFIVECrowdsale *BEP20KFIVECrowdsaleSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _BEP20KFIVECrowdsale.Contract.TransferOwnership(&_BEP20KFIVECrowdsale.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_BEP20KFIVECrowdsale *BEP20KFIVECrowdsaleTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _BEP20KFIVECrowdsale.Contract.TransferOwnership(&_BEP20KFIVECrowdsale.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_BEP20KFIVECrowdsale *BEP20KFIVECrowdsaleTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BEP20KFIVECrowdsale.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_BEP20KFIVECrowdsale *BEP20KFIVECrowdsaleSession) Unpause() (*types.Transaction, error) {
	return _BEP20KFIVECrowdsale.Contract.Unpause(&_BEP20KFIVECrowdsale.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_BEP20KFIVECrowdsale *BEP20KFIVECrowdsaleTransactorSession) Unpause() (*types.Transaction, error) {
	return _BEP20KFIVECrowdsale.Contract.Unpause(&_BEP20KFIVECrowdsale.TransactOpts)
}

// BEP20KFIVECrowdsaleOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the BEP20KFIVECrowdsale contract.
type BEP20KFIVECrowdsaleOwnershipTransferredIterator struct {
	Event *BEP20KFIVECrowdsaleOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *BEP20KFIVECrowdsaleOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BEP20KFIVECrowdsaleOwnershipTransferred)
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
		it.Event = new(BEP20KFIVECrowdsaleOwnershipTransferred)
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
func (it *BEP20KFIVECrowdsaleOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BEP20KFIVECrowdsaleOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BEP20KFIVECrowdsaleOwnershipTransferred represents a OwnershipTransferred event raised by the BEP20KFIVECrowdsale contract.
type BEP20KFIVECrowdsaleOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_BEP20KFIVECrowdsale *BEP20KFIVECrowdsaleFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*BEP20KFIVECrowdsaleOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _BEP20KFIVECrowdsale.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &BEP20KFIVECrowdsaleOwnershipTransferredIterator{contract: _BEP20KFIVECrowdsale.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_BEP20KFIVECrowdsale *BEP20KFIVECrowdsaleFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *BEP20KFIVECrowdsaleOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _BEP20KFIVECrowdsale.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BEP20KFIVECrowdsaleOwnershipTransferred)
				if err := _BEP20KFIVECrowdsale.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_BEP20KFIVECrowdsale *BEP20KFIVECrowdsaleFilterer) ParseOwnershipTransferred(log types.Log) (*BEP20KFIVECrowdsaleOwnershipTransferred, error) {
	event := new(BEP20KFIVECrowdsaleOwnershipTransferred)
	if err := _BEP20KFIVECrowdsale.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BEP20KFIVECrowdsalePausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the BEP20KFIVECrowdsale contract.
type BEP20KFIVECrowdsalePausedIterator struct {
	Event *BEP20KFIVECrowdsalePaused // Event containing the contract specifics and raw log

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
func (it *BEP20KFIVECrowdsalePausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BEP20KFIVECrowdsalePaused)
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
		it.Event = new(BEP20KFIVECrowdsalePaused)
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
func (it *BEP20KFIVECrowdsalePausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BEP20KFIVECrowdsalePausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BEP20KFIVECrowdsalePaused represents a Paused event raised by the BEP20KFIVECrowdsale contract.
type BEP20KFIVECrowdsalePaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_BEP20KFIVECrowdsale *BEP20KFIVECrowdsaleFilterer) FilterPaused(opts *bind.FilterOpts) (*BEP20KFIVECrowdsalePausedIterator, error) {

	logs, sub, err := _BEP20KFIVECrowdsale.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &BEP20KFIVECrowdsalePausedIterator{contract: _BEP20KFIVECrowdsale.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_BEP20KFIVECrowdsale *BEP20KFIVECrowdsaleFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *BEP20KFIVECrowdsalePaused) (event.Subscription, error) {

	logs, sub, err := _BEP20KFIVECrowdsale.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BEP20KFIVECrowdsalePaused)
				if err := _BEP20KFIVECrowdsale.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_BEP20KFIVECrowdsale *BEP20KFIVECrowdsaleFilterer) ParsePaused(log types.Log) (*BEP20KFIVECrowdsalePaused, error) {
	event := new(BEP20KFIVECrowdsalePaused)
	if err := _BEP20KFIVECrowdsale.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BEP20KFIVECrowdsaleTimedCrowdsaleExtendedIterator is returned from FilterTimedCrowdsaleExtended and is used to iterate over the raw logs and unpacked data for TimedCrowdsaleExtended events raised by the BEP20KFIVECrowdsale contract.
type BEP20KFIVECrowdsaleTimedCrowdsaleExtendedIterator struct {
	Event *BEP20KFIVECrowdsaleTimedCrowdsaleExtended // Event containing the contract specifics and raw log

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
func (it *BEP20KFIVECrowdsaleTimedCrowdsaleExtendedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BEP20KFIVECrowdsaleTimedCrowdsaleExtended)
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
		it.Event = new(BEP20KFIVECrowdsaleTimedCrowdsaleExtended)
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
func (it *BEP20KFIVECrowdsaleTimedCrowdsaleExtendedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BEP20KFIVECrowdsaleTimedCrowdsaleExtendedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BEP20KFIVECrowdsaleTimedCrowdsaleExtended represents a TimedCrowdsaleExtended event raised by the BEP20KFIVECrowdsale contract.
type BEP20KFIVECrowdsaleTimedCrowdsaleExtended struct {
	PrevClosingTime *big.Int
	NewClosingTime  *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterTimedCrowdsaleExtended is a free log retrieval operation binding the contract event 0x46711e222f558a07afd26e5e71b48ecb0a8b2cdcd40faeb1323e05e2c76a2f32.
//
// Solidity: event TimedCrowdsaleExtended(uint256 prevClosingTime, uint256 newClosingTime)
func (_BEP20KFIVECrowdsale *BEP20KFIVECrowdsaleFilterer) FilterTimedCrowdsaleExtended(opts *bind.FilterOpts) (*BEP20KFIVECrowdsaleTimedCrowdsaleExtendedIterator, error) {

	logs, sub, err := _BEP20KFIVECrowdsale.contract.FilterLogs(opts, "TimedCrowdsaleExtended")
	if err != nil {
		return nil, err
	}
	return &BEP20KFIVECrowdsaleTimedCrowdsaleExtendedIterator{contract: _BEP20KFIVECrowdsale.contract, event: "TimedCrowdsaleExtended", logs: logs, sub: sub}, nil
}

// WatchTimedCrowdsaleExtended is a free log subscription operation binding the contract event 0x46711e222f558a07afd26e5e71b48ecb0a8b2cdcd40faeb1323e05e2c76a2f32.
//
// Solidity: event TimedCrowdsaleExtended(uint256 prevClosingTime, uint256 newClosingTime)
func (_BEP20KFIVECrowdsale *BEP20KFIVECrowdsaleFilterer) WatchTimedCrowdsaleExtended(opts *bind.WatchOpts, sink chan<- *BEP20KFIVECrowdsaleTimedCrowdsaleExtended) (event.Subscription, error) {

	logs, sub, err := _BEP20KFIVECrowdsale.contract.WatchLogs(opts, "TimedCrowdsaleExtended")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BEP20KFIVECrowdsaleTimedCrowdsaleExtended)
				if err := _BEP20KFIVECrowdsale.contract.UnpackLog(event, "TimedCrowdsaleExtended", log); err != nil {
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
func (_BEP20KFIVECrowdsale *BEP20KFIVECrowdsaleFilterer) ParseTimedCrowdsaleExtended(log types.Log) (*BEP20KFIVECrowdsaleTimedCrowdsaleExtended, error) {
	event := new(BEP20KFIVECrowdsaleTimedCrowdsaleExtended)
	if err := _BEP20KFIVECrowdsale.contract.UnpackLog(event, "TimedCrowdsaleExtended", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BEP20KFIVECrowdsaleTokensPurchasedIterator is returned from FilterTokensPurchased and is used to iterate over the raw logs and unpacked data for TokensPurchased events raised by the BEP20KFIVECrowdsale contract.
type BEP20KFIVECrowdsaleTokensPurchasedIterator struct {
	Event *BEP20KFIVECrowdsaleTokensPurchased // Event containing the contract specifics and raw log

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
func (it *BEP20KFIVECrowdsaleTokensPurchasedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BEP20KFIVECrowdsaleTokensPurchased)
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
		it.Event = new(BEP20KFIVECrowdsaleTokensPurchased)
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
func (it *BEP20KFIVECrowdsaleTokensPurchasedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BEP20KFIVECrowdsaleTokensPurchasedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BEP20KFIVECrowdsaleTokensPurchased represents a TokensPurchased event raised by the BEP20KFIVECrowdsale contract.
type BEP20KFIVECrowdsaleTokensPurchased struct {
	Purchaser   common.Address
	Beneficiary common.Address
	Value       *big.Int
	Amount      *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterTokensPurchased is a free log retrieval operation binding the contract event 0x6faf93231a456e552dbc9961f58d9713ee4f2e69d15f1975b050ef0911053a7b.
//
// Solidity: event TokensPurchased(address indexed purchaser, address indexed beneficiary, uint256 value, uint256 amount)
func (_BEP20KFIVECrowdsale *BEP20KFIVECrowdsaleFilterer) FilterTokensPurchased(opts *bind.FilterOpts, purchaser []common.Address, beneficiary []common.Address) (*BEP20KFIVECrowdsaleTokensPurchasedIterator, error) {

	var purchaserRule []interface{}
	for _, purchaserItem := range purchaser {
		purchaserRule = append(purchaserRule, purchaserItem)
	}
	var beneficiaryRule []interface{}
	for _, beneficiaryItem := range beneficiary {
		beneficiaryRule = append(beneficiaryRule, beneficiaryItem)
	}

	logs, sub, err := _BEP20KFIVECrowdsale.contract.FilterLogs(opts, "TokensPurchased", purchaserRule, beneficiaryRule)
	if err != nil {
		return nil, err
	}
	return &BEP20KFIVECrowdsaleTokensPurchasedIterator{contract: _BEP20KFIVECrowdsale.contract, event: "TokensPurchased", logs: logs, sub: sub}, nil
}

// WatchTokensPurchased is a free log subscription operation binding the contract event 0x6faf93231a456e552dbc9961f58d9713ee4f2e69d15f1975b050ef0911053a7b.
//
// Solidity: event TokensPurchased(address indexed purchaser, address indexed beneficiary, uint256 value, uint256 amount)
func (_BEP20KFIVECrowdsale *BEP20KFIVECrowdsaleFilterer) WatchTokensPurchased(opts *bind.WatchOpts, sink chan<- *BEP20KFIVECrowdsaleTokensPurchased, purchaser []common.Address, beneficiary []common.Address) (event.Subscription, error) {

	var purchaserRule []interface{}
	for _, purchaserItem := range purchaser {
		purchaserRule = append(purchaserRule, purchaserItem)
	}
	var beneficiaryRule []interface{}
	for _, beneficiaryItem := range beneficiary {
		beneficiaryRule = append(beneficiaryRule, beneficiaryItem)
	}

	logs, sub, err := _BEP20KFIVECrowdsale.contract.WatchLogs(opts, "TokensPurchased", purchaserRule, beneficiaryRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BEP20KFIVECrowdsaleTokensPurchased)
				if err := _BEP20KFIVECrowdsale.contract.UnpackLog(event, "TokensPurchased", log); err != nil {
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
func (_BEP20KFIVECrowdsale *BEP20KFIVECrowdsaleFilterer) ParseTokensPurchased(log types.Log) (*BEP20KFIVECrowdsaleTokensPurchased, error) {
	event := new(BEP20KFIVECrowdsaleTokensPurchased)
	if err := _BEP20KFIVECrowdsale.contract.UnpackLog(event, "TokensPurchased", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BEP20KFIVECrowdsaleUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the BEP20KFIVECrowdsale contract.
type BEP20KFIVECrowdsaleUnpausedIterator struct {
	Event *BEP20KFIVECrowdsaleUnpaused // Event containing the contract specifics and raw log

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
func (it *BEP20KFIVECrowdsaleUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BEP20KFIVECrowdsaleUnpaused)
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
		it.Event = new(BEP20KFIVECrowdsaleUnpaused)
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
func (it *BEP20KFIVECrowdsaleUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BEP20KFIVECrowdsaleUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BEP20KFIVECrowdsaleUnpaused represents a Unpaused event raised by the BEP20KFIVECrowdsale contract.
type BEP20KFIVECrowdsaleUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_BEP20KFIVECrowdsale *BEP20KFIVECrowdsaleFilterer) FilterUnpaused(opts *bind.FilterOpts) (*BEP20KFIVECrowdsaleUnpausedIterator, error) {

	logs, sub, err := _BEP20KFIVECrowdsale.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &BEP20KFIVECrowdsaleUnpausedIterator{contract: _BEP20KFIVECrowdsale.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_BEP20KFIVECrowdsale *BEP20KFIVECrowdsaleFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *BEP20KFIVECrowdsaleUnpaused) (event.Subscription, error) {

	logs, sub, err := _BEP20KFIVECrowdsale.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BEP20KFIVECrowdsaleUnpaused)
				if err := _BEP20KFIVECrowdsale.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
func (_BEP20KFIVECrowdsale *BEP20KFIVECrowdsaleFilterer) ParseUnpaused(log types.Log) (*BEP20KFIVECrowdsaleUnpaused, error) {
	event := new(BEP20KFIVECrowdsaleUnpaused)
	if err := _BEP20KFIVECrowdsale.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
