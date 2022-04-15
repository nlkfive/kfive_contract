// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package BEP20Crowdsale

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

// BEP20CrowdsaleMetaData contains all meta data concerning the BEP20Crowdsale contract.
var BEP20CrowdsaleMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"__rate\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"__wallet\",\"type\":\"address\"},{\"internalType\":\"contractIBEP20\",\"name\":\"__token\",\"type\":\"address\"},{\"internalType\":\"contractIBEP20\",\"name\":\"__acceptToken\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"purchaser\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"beneficiary\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"TokensPurchased\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token\",\"outputs\":[{\"internalType\":\"contractIBEP20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"wallet\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"weiRaised\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"_to\",\"type\":\"address\"}],\"name\":\"destroySmartContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"newReceiver\",\"type\":\"address\"}],\"name\":\"changeWallet\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"beneficiary\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"buyTokens\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
}

// BEP20CrowdsaleABI is the input ABI used to generate the binding from.
// Deprecated: Use BEP20CrowdsaleMetaData.ABI instead.
var BEP20CrowdsaleABI = BEP20CrowdsaleMetaData.ABI

// BEP20Crowdsale is an auto generated Go binding around an Ethereum contract.
type BEP20Crowdsale struct {
	BEP20CrowdsaleCaller     // Read-only binding to the contract
	BEP20CrowdsaleTransactor // Write-only binding to the contract
	BEP20CrowdsaleFilterer   // Log filterer for contract events
}

// BEP20CrowdsaleCaller is an auto generated read-only Go binding around an Ethereum contract.
type BEP20CrowdsaleCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BEP20CrowdsaleTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BEP20CrowdsaleTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BEP20CrowdsaleFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BEP20CrowdsaleFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BEP20CrowdsaleSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BEP20CrowdsaleSession struct {
	Contract     *BEP20Crowdsale   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BEP20CrowdsaleCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BEP20CrowdsaleCallerSession struct {
	Contract *BEP20CrowdsaleCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// BEP20CrowdsaleTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BEP20CrowdsaleTransactorSession struct {
	Contract     *BEP20CrowdsaleTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// BEP20CrowdsaleRaw is an auto generated low-level Go binding around an Ethereum contract.
type BEP20CrowdsaleRaw struct {
	Contract *BEP20Crowdsale // Generic contract binding to access the raw methods on
}

// BEP20CrowdsaleCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BEP20CrowdsaleCallerRaw struct {
	Contract *BEP20CrowdsaleCaller // Generic read-only contract binding to access the raw methods on
}

// BEP20CrowdsaleTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BEP20CrowdsaleTransactorRaw struct {
	Contract *BEP20CrowdsaleTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBEP20Crowdsale creates a new instance of BEP20Crowdsale, bound to a specific deployed contract.
func NewBEP20Crowdsale(address common.Address, backend bind.ContractBackend) (*BEP20Crowdsale, error) {
	contract, err := bindBEP20Crowdsale(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BEP20Crowdsale{BEP20CrowdsaleCaller: BEP20CrowdsaleCaller{contract: contract}, BEP20CrowdsaleTransactor: BEP20CrowdsaleTransactor{contract: contract}, BEP20CrowdsaleFilterer: BEP20CrowdsaleFilterer{contract: contract}}, nil
}

// NewBEP20CrowdsaleCaller creates a new read-only instance of BEP20Crowdsale, bound to a specific deployed contract.
func NewBEP20CrowdsaleCaller(address common.Address, caller bind.ContractCaller) (*BEP20CrowdsaleCaller, error) {
	contract, err := bindBEP20Crowdsale(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BEP20CrowdsaleCaller{contract: contract}, nil
}

// NewBEP20CrowdsaleTransactor creates a new write-only instance of BEP20Crowdsale, bound to a specific deployed contract.
func NewBEP20CrowdsaleTransactor(address common.Address, transactor bind.ContractTransactor) (*BEP20CrowdsaleTransactor, error) {
	contract, err := bindBEP20Crowdsale(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BEP20CrowdsaleTransactor{contract: contract}, nil
}

// NewBEP20CrowdsaleFilterer creates a new log filterer instance of BEP20Crowdsale, bound to a specific deployed contract.
func NewBEP20CrowdsaleFilterer(address common.Address, filterer bind.ContractFilterer) (*BEP20CrowdsaleFilterer, error) {
	contract, err := bindBEP20Crowdsale(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BEP20CrowdsaleFilterer{contract: contract}, nil
}

// bindBEP20Crowdsale binds a generic wrapper to an already deployed contract.
func bindBEP20Crowdsale(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BEP20CrowdsaleABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BEP20Crowdsale *BEP20CrowdsaleRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BEP20Crowdsale.Contract.BEP20CrowdsaleCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BEP20Crowdsale *BEP20CrowdsaleRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BEP20Crowdsale.Contract.BEP20CrowdsaleTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BEP20Crowdsale *BEP20CrowdsaleRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BEP20Crowdsale.Contract.BEP20CrowdsaleTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BEP20Crowdsale *BEP20CrowdsaleCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BEP20Crowdsale.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BEP20Crowdsale *BEP20CrowdsaleTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BEP20Crowdsale.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BEP20Crowdsale *BEP20CrowdsaleTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BEP20Crowdsale.Contract.contract.Transact(opts, method, params...)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BEP20Crowdsale *BEP20CrowdsaleCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BEP20Crowdsale.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BEP20Crowdsale *BEP20CrowdsaleSession) Owner() (common.Address, error) {
	return _BEP20Crowdsale.Contract.Owner(&_BEP20Crowdsale.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BEP20Crowdsale *BEP20CrowdsaleCallerSession) Owner() (common.Address, error) {
	return _BEP20Crowdsale.Contract.Owner(&_BEP20Crowdsale.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_BEP20Crowdsale *BEP20CrowdsaleCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _BEP20Crowdsale.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_BEP20Crowdsale *BEP20CrowdsaleSession) Paused() (bool, error) {
	return _BEP20Crowdsale.Contract.Paused(&_BEP20Crowdsale.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_BEP20Crowdsale *BEP20CrowdsaleCallerSession) Paused() (bool, error) {
	return _BEP20Crowdsale.Contract.Paused(&_BEP20Crowdsale.CallOpts)
}

// Rate is a free data retrieval call binding the contract method 0x2c4e722e.
//
// Solidity: function rate() view returns(uint256)
func (_BEP20Crowdsale *BEP20CrowdsaleCaller) Rate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BEP20Crowdsale.contract.Call(opts, &out, "rate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Rate is a free data retrieval call binding the contract method 0x2c4e722e.
//
// Solidity: function rate() view returns(uint256)
func (_BEP20Crowdsale *BEP20CrowdsaleSession) Rate() (*big.Int, error) {
	return _BEP20Crowdsale.Contract.Rate(&_BEP20Crowdsale.CallOpts)
}

// Rate is a free data retrieval call binding the contract method 0x2c4e722e.
//
// Solidity: function rate() view returns(uint256)
func (_BEP20Crowdsale *BEP20CrowdsaleCallerSession) Rate() (*big.Int, error) {
	return _BEP20Crowdsale.Contract.Rate(&_BEP20Crowdsale.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_BEP20Crowdsale *BEP20CrowdsaleCaller) Token(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BEP20Crowdsale.contract.Call(opts, &out, "token")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_BEP20Crowdsale *BEP20CrowdsaleSession) Token() (common.Address, error) {
	return _BEP20Crowdsale.Contract.Token(&_BEP20Crowdsale.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_BEP20Crowdsale *BEP20CrowdsaleCallerSession) Token() (common.Address, error) {
	return _BEP20Crowdsale.Contract.Token(&_BEP20Crowdsale.CallOpts)
}

// Wallet is a free data retrieval call binding the contract method 0x521eb273.
//
// Solidity: function wallet() view returns(address)
func (_BEP20Crowdsale *BEP20CrowdsaleCaller) Wallet(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BEP20Crowdsale.contract.Call(opts, &out, "wallet")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Wallet is a free data retrieval call binding the contract method 0x521eb273.
//
// Solidity: function wallet() view returns(address)
func (_BEP20Crowdsale *BEP20CrowdsaleSession) Wallet() (common.Address, error) {
	return _BEP20Crowdsale.Contract.Wallet(&_BEP20Crowdsale.CallOpts)
}

// Wallet is a free data retrieval call binding the contract method 0x521eb273.
//
// Solidity: function wallet() view returns(address)
func (_BEP20Crowdsale *BEP20CrowdsaleCallerSession) Wallet() (common.Address, error) {
	return _BEP20Crowdsale.Contract.Wallet(&_BEP20Crowdsale.CallOpts)
}

// WeiRaised is a free data retrieval call binding the contract method 0x4042b66f.
//
// Solidity: function weiRaised() view returns(uint256)
func (_BEP20Crowdsale *BEP20CrowdsaleCaller) WeiRaised(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BEP20Crowdsale.contract.Call(opts, &out, "weiRaised")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WeiRaised is a free data retrieval call binding the contract method 0x4042b66f.
//
// Solidity: function weiRaised() view returns(uint256)
func (_BEP20Crowdsale *BEP20CrowdsaleSession) WeiRaised() (*big.Int, error) {
	return _BEP20Crowdsale.Contract.WeiRaised(&_BEP20Crowdsale.CallOpts)
}

// WeiRaised is a free data retrieval call binding the contract method 0x4042b66f.
//
// Solidity: function weiRaised() view returns(uint256)
func (_BEP20Crowdsale *BEP20CrowdsaleCallerSession) WeiRaised() (*big.Int, error) {
	return _BEP20Crowdsale.Contract.WeiRaised(&_BEP20Crowdsale.CallOpts)
}

// BuyTokens is a paid mutator transaction binding the contract method 0x0752881a.
//
// Solidity: function buyTokens(address beneficiary, uint256 value) payable returns()
func (_BEP20Crowdsale *BEP20CrowdsaleTransactor) BuyTokens(opts *bind.TransactOpts, beneficiary common.Address, value *big.Int) (*types.Transaction, error) {
	return _BEP20Crowdsale.contract.Transact(opts, "buyTokens", beneficiary, value)
}

// BuyTokens is a paid mutator transaction binding the contract method 0x0752881a.
//
// Solidity: function buyTokens(address beneficiary, uint256 value) payable returns()
func (_BEP20Crowdsale *BEP20CrowdsaleSession) BuyTokens(beneficiary common.Address, value *big.Int) (*types.Transaction, error) {
	return _BEP20Crowdsale.Contract.BuyTokens(&_BEP20Crowdsale.TransactOpts, beneficiary, value)
}

// BuyTokens is a paid mutator transaction binding the contract method 0x0752881a.
//
// Solidity: function buyTokens(address beneficiary, uint256 value) payable returns()
func (_BEP20Crowdsale *BEP20CrowdsaleTransactorSession) BuyTokens(beneficiary common.Address, value *big.Int) (*types.Transaction, error) {
	return _BEP20Crowdsale.Contract.BuyTokens(&_BEP20Crowdsale.TransactOpts, beneficiary, value)
}

// ChangeWallet is a paid mutator transaction binding the contract method 0x98b9a2dc.
//
// Solidity: function changeWallet(address newReceiver) returns()
func (_BEP20Crowdsale *BEP20CrowdsaleTransactor) ChangeWallet(opts *bind.TransactOpts, newReceiver common.Address) (*types.Transaction, error) {
	return _BEP20Crowdsale.contract.Transact(opts, "changeWallet", newReceiver)
}

// ChangeWallet is a paid mutator transaction binding the contract method 0x98b9a2dc.
//
// Solidity: function changeWallet(address newReceiver) returns()
func (_BEP20Crowdsale *BEP20CrowdsaleSession) ChangeWallet(newReceiver common.Address) (*types.Transaction, error) {
	return _BEP20Crowdsale.Contract.ChangeWallet(&_BEP20Crowdsale.TransactOpts, newReceiver)
}

// ChangeWallet is a paid mutator transaction binding the contract method 0x98b9a2dc.
//
// Solidity: function changeWallet(address newReceiver) returns()
func (_BEP20Crowdsale *BEP20CrowdsaleTransactorSession) ChangeWallet(newReceiver common.Address) (*types.Transaction, error) {
	return _BEP20Crowdsale.Contract.ChangeWallet(&_BEP20Crowdsale.TransactOpts, newReceiver)
}

// DestroySmartContract is a paid mutator transaction binding the contract method 0x39df43ff.
//
// Solidity: function destroySmartContract(address _to) returns()
func (_BEP20Crowdsale *BEP20CrowdsaleTransactor) DestroySmartContract(opts *bind.TransactOpts, _to common.Address) (*types.Transaction, error) {
	return _BEP20Crowdsale.contract.Transact(opts, "destroySmartContract", _to)
}

// DestroySmartContract is a paid mutator transaction binding the contract method 0x39df43ff.
//
// Solidity: function destroySmartContract(address _to) returns()
func (_BEP20Crowdsale *BEP20CrowdsaleSession) DestroySmartContract(_to common.Address) (*types.Transaction, error) {
	return _BEP20Crowdsale.Contract.DestroySmartContract(&_BEP20Crowdsale.TransactOpts, _to)
}

// DestroySmartContract is a paid mutator transaction binding the contract method 0x39df43ff.
//
// Solidity: function destroySmartContract(address _to) returns()
func (_BEP20Crowdsale *BEP20CrowdsaleTransactorSession) DestroySmartContract(_to common.Address) (*types.Transaction, error) {
	return _BEP20Crowdsale.Contract.DestroySmartContract(&_BEP20Crowdsale.TransactOpts, _to)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_BEP20Crowdsale *BEP20CrowdsaleTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BEP20Crowdsale.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_BEP20Crowdsale *BEP20CrowdsaleSession) Pause() (*types.Transaction, error) {
	return _BEP20Crowdsale.Contract.Pause(&_BEP20Crowdsale.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_BEP20Crowdsale *BEP20CrowdsaleTransactorSession) Pause() (*types.Transaction, error) {
	return _BEP20Crowdsale.Contract.Pause(&_BEP20Crowdsale.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_BEP20Crowdsale *BEP20CrowdsaleTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BEP20Crowdsale.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_BEP20Crowdsale *BEP20CrowdsaleSession) RenounceOwnership() (*types.Transaction, error) {
	return _BEP20Crowdsale.Contract.RenounceOwnership(&_BEP20Crowdsale.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_BEP20Crowdsale *BEP20CrowdsaleTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _BEP20Crowdsale.Contract.RenounceOwnership(&_BEP20Crowdsale.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_BEP20Crowdsale *BEP20CrowdsaleTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _BEP20Crowdsale.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_BEP20Crowdsale *BEP20CrowdsaleSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _BEP20Crowdsale.Contract.TransferOwnership(&_BEP20Crowdsale.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_BEP20Crowdsale *BEP20CrowdsaleTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _BEP20Crowdsale.Contract.TransferOwnership(&_BEP20Crowdsale.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_BEP20Crowdsale *BEP20CrowdsaleTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BEP20Crowdsale.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_BEP20Crowdsale *BEP20CrowdsaleSession) Unpause() (*types.Transaction, error) {
	return _BEP20Crowdsale.Contract.Unpause(&_BEP20Crowdsale.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_BEP20Crowdsale *BEP20CrowdsaleTransactorSession) Unpause() (*types.Transaction, error) {
	return _BEP20Crowdsale.Contract.Unpause(&_BEP20Crowdsale.TransactOpts)
}

// BEP20CrowdsaleOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the BEP20Crowdsale contract.
type BEP20CrowdsaleOwnershipTransferredIterator struct {
	Event *BEP20CrowdsaleOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *BEP20CrowdsaleOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BEP20CrowdsaleOwnershipTransferred)
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
		it.Event = new(BEP20CrowdsaleOwnershipTransferred)
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
func (it *BEP20CrowdsaleOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BEP20CrowdsaleOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BEP20CrowdsaleOwnershipTransferred represents a OwnershipTransferred event raised by the BEP20Crowdsale contract.
type BEP20CrowdsaleOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_BEP20Crowdsale *BEP20CrowdsaleFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*BEP20CrowdsaleOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _BEP20Crowdsale.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &BEP20CrowdsaleOwnershipTransferredIterator{contract: _BEP20Crowdsale.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_BEP20Crowdsale *BEP20CrowdsaleFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *BEP20CrowdsaleOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _BEP20Crowdsale.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BEP20CrowdsaleOwnershipTransferred)
				if err := _BEP20Crowdsale.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_BEP20Crowdsale *BEP20CrowdsaleFilterer) ParseOwnershipTransferred(log types.Log) (*BEP20CrowdsaleOwnershipTransferred, error) {
	event := new(BEP20CrowdsaleOwnershipTransferred)
	if err := _BEP20Crowdsale.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BEP20CrowdsalePausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the BEP20Crowdsale contract.
type BEP20CrowdsalePausedIterator struct {
	Event *BEP20CrowdsalePaused // Event containing the contract specifics and raw log

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
func (it *BEP20CrowdsalePausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BEP20CrowdsalePaused)
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
		it.Event = new(BEP20CrowdsalePaused)
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
func (it *BEP20CrowdsalePausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BEP20CrowdsalePausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BEP20CrowdsalePaused represents a Paused event raised by the BEP20Crowdsale contract.
type BEP20CrowdsalePaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_BEP20Crowdsale *BEP20CrowdsaleFilterer) FilterPaused(opts *bind.FilterOpts) (*BEP20CrowdsalePausedIterator, error) {

	logs, sub, err := _BEP20Crowdsale.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &BEP20CrowdsalePausedIterator{contract: _BEP20Crowdsale.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_BEP20Crowdsale *BEP20CrowdsaleFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *BEP20CrowdsalePaused) (event.Subscription, error) {

	logs, sub, err := _BEP20Crowdsale.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BEP20CrowdsalePaused)
				if err := _BEP20Crowdsale.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_BEP20Crowdsale *BEP20CrowdsaleFilterer) ParsePaused(log types.Log) (*BEP20CrowdsalePaused, error) {
	event := new(BEP20CrowdsalePaused)
	if err := _BEP20Crowdsale.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BEP20CrowdsaleTokensPurchasedIterator is returned from FilterTokensPurchased and is used to iterate over the raw logs and unpacked data for TokensPurchased events raised by the BEP20Crowdsale contract.
type BEP20CrowdsaleTokensPurchasedIterator struct {
	Event *BEP20CrowdsaleTokensPurchased // Event containing the contract specifics and raw log

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
func (it *BEP20CrowdsaleTokensPurchasedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BEP20CrowdsaleTokensPurchased)
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
		it.Event = new(BEP20CrowdsaleTokensPurchased)
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
func (it *BEP20CrowdsaleTokensPurchasedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BEP20CrowdsaleTokensPurchasedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BEP20CrowdsaleTokensPurchased represents a TokensPurchased event raised by the BEP20Crowdsale contract.
type BEP20CrowdsaleTokensPurchased struct {
	Purchaser   common.Address
	Beneficiary common.Address
	Value       *big.Int
	Amount      *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterTokensPurchased is a free log retrieval operation binding the contract event 0x6faf93231a456e552dbc9961f58d9713ee4f2e69d15f1975b050ef0911053a7b.
//
// Solidity: event TokensPurchased(address indexed purchaser, address indexed beneficiary, uint256 value, uint256 amount)
func (_BEP20Crowdsale *BEP20CrowdsaleFilterer) FilterTokensPurchased(opts *bind.FilterOpts, purchaser []common.Address, beneficiary []common.Address) (*BEP20CrowdsaleTokensPurchasedIterator, error) {

	var purchaserRule []interface{}
	for _, purchaserItem := range purchaser {
		purchaserRule = append(purchaserRule, purchaserItem)
	}
	var beneficiaryRule []interface{}
	for _, beneficiaryItem := range beneficiary {
		beneficiaryRule = append(beneficiaryRule, beneficiaryItem)
	}

	logs, sub, err := _BEP20Crowdsale.contract.FilterLogs(opts, "TokensPurchased", purchaserRule, beneficiaryRule)
	if err != nil {
		return nil, err
	}
	return &BEP20CrowdsaleTokensPurchasedIterator{contract: _BEP20Crowdsale.contract, event: "TokensPurchased", logs: logs, sub: sub}, nil
}

// WatchTokensPurchased is a free log subscription operation binding the contract event 0x6faf93231a456e552dbc9961f58d9713ee4f2e69d15f1975b050ef0911053a7b.
//
// Solidity: event TokensPurchased(address indexed purchaser, address indexed beneficiary, uint256 value, uint256 amount)
func (_BEP20Crowdsale *BEP20CrowdsaleFilterer) WatchTokensPurchased(opts *bind.WatchOpts, sink chan<- *BEP20CrowdsaleTokensPurchased, purchaser []common.Address, beneficiary []common.Address) (event.Subscription, error) {

	var purchaserRule []interface{}
	for _, purchaserItem := range purchaser {
		purchaserRule = append(purchaserRule, purchaserItem)
	}
	var beneficiaryRule []interface{}
	for _, beneficiaryItem := range beneficiary {
		beneficiaryRule = append(beneficiaryRule, beneficiaryItem)
	}

	logs, sub, err := _BEP20Crowdsale.contract.WatchLogs(opts, "TokensPurchased", purchaserRule, beneficiaryRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BEP20CrowdsaleTokensPurchased)
				if err := _BEP20Crowdsale.contract.UnpackLog(event, "TokensPurchased", log); err != nil {
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
func (_BEP20Crowdsale *BEP20CrowdsaleFilterer) ParseTokensPurchased(log types.Log) (*BEP20CrowdsaleTokensPurchased, error) {
	event := new(BEP20CrowdsaleTokensPurchased)
	if err := _BEP20Crowdsale.contract.UnpackLog(event, "TokensPurchased", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BEP20CrowdsaleUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the BEP20Crowdsale contract.
type BEP20CrowdsaleUnpausedIterator struct {
	Event *BEP20CrowdsaleUnpaused // Event containing the contract specifics and raw log

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
func (it *BEP20CrowdsaleUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BEP20CrowdsaleUnpaused)
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
		it.Event = new(BEP20CrowdsaleUnpaused)
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
func (it *BEP20CrowdsaleUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BEP20CrowdsaleUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BEP20CrowdsaleUnpaused represents a Unpaused event raised by the BEP20Crowdsale contract.
type BEP20CrowdsaleUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_BEP20Crowdsale *BEP20CrowdsaleFilterer) FilterUnpaused(opts *bind.FilterOpts) (*BEP20CrowdsaleUnpausedIterator, error) {

	logs, sub, err := _BEP20Crowdsale.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &BEP20CrowdsaleUnpausedIterator{contract: _BEP20Crowdsale.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_BEP20Crowdsale *BEP20CrowdsaleFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *BEP20CrowdsaleUnpaused) (event.Subscription, error) {

	logs, sub, err := _BEP20Crowdsale.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BEP20CrowdsaleUnpaused)
				if err := _BEP20Crowdsale.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
func (_BEP20Crowdsale *BEP20CrowdsaleFilterer) ParseUnpaused(log types.Log) (*BEP20CrowdsaleUnpaused, error) {
	event := new(BEP20CrowdsaleUnpaused)
	if err := _BEP20Crowdsale.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
