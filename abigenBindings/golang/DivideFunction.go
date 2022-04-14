// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package DivideFunction

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

// DivideFunctionMetaData contains all meta data concerning the DivideFunction contract.
var DivideFunctionMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"weiValue\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rate\",\"type\":\"uint256\"}],\"name\":\"weiToKfive\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]",
}

// DivideFunctionABI is the input ABI used to generate the binding from.
// Deprecated: Use DivideFunctionMetaData.ABI instead.
var DivideFunctionABI = DivideFunctionMetaData.ABI

// DivideFunction is an auto generated Go binding around an Ethereum contract.
type DivideFunction struct {
	DivideFunctionCaller     // Read-only binding to the contract
	DivideFunctionTransactor // Write-only binding to the contract
	DivideFunctionFilterer   // Log filterer for contract events
}

// DivideFunctionCaller is an auto generated read-only Go binding around an Ethereum contract.
type DivideFunctionCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DivideFunctionTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DivideFunctionTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DivideFunctionFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DivideFunctionFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DivideFunctionSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DivideFunctionSession struct {
	Contract     *DivideFunction   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DivideFunctionCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DivideFunctionCallerSession struct {
	Contract *DivideFunctionCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// DivideFunctionTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DivideFunctionTransactorSession struct {
	Contract     *DivideFunctionTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// DivideFunctionRaw is an auto generated low-level Go binding around an Ethereum contract.
type DivideFunctionRaw struct {
	Contract *DivideFunction // Generic contract binding to access the raw methods on
}

// DivideFunctionCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DivideFunctionCallerRaw struct {
	Contract *DivideFunctionCaller // Generic read-only contract binding to access the raw methods on
}

// DivideFunctionTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DivideFunctionTransactorRaw struct {
	Contract *DivideFunctionTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDivideFunction creates a new instance of DivideFunction, bound to a specific deployed contract.
func NewDivideFunction(address common.Address, backend bind.ContractBackend) (*DivideFunction, error) {
	contract, err := bindDivideFunction(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DivideFunction{DivideFunctionCaller: DivideFunctionCaller{contract: contract}, DivideFunctionTransactor: DivideFunctionTransactor{contract: contract}, DivideFunctionFilterer: DivideFunctionFilterer{contract: contract}}, nil
}

// NewDivideFunctionCaller creates a new read-only instance of DivideFunction, bound to a specific deployed contract.
func NewDivideFunctionCaller(address common.Address, caller bind.ContractCaller) (*DivideFunctionCaller, error) {
	contract, err := bindDivideFunction(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DivideFunctionCaller{contract: contract}, nil
}

// NewDivideFunctionTransactor creates a new write-only instance of DivideFunction, bound to a specific deployed contract.
func NewDivideFunctionTransactor(address common.Address, transactor bind.ContractTransactor) (*DivideFunctionTransactor, error) {
	contract, err := bindDivideFunction(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DivideFunctionTransactor{contract: contract}, nil
}

// NewDivideFunctionFilterer creates a new log filterer instance of DivideFunction, bound to a specific deployed contract.
func NewDivideFunctionFilterer(address common.Address, filterer bind.ContractFilterer) (*DivideFunctionFilterer, error) {
	contract, err := bindDivideFunction(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DivideFunctionFilterer{contract: contract}, nil
}

// bindDivideFunction binds a generic wrapper to an already deployed contract.
func bindDivideFunction(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DivideFunctionABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DivideFunction *DivideFunctionRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DivideFunction.Contract.DivideFunctionCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DivideFunction *DivideFunctionRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DivideFunction.Contract.DivideFunctionTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DivideFunction *DivideFunctionRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DivideFunction.Contract.DivideFunctionTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DivideFunction *DivideFunctionCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DivideFunction.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DivideFunction *DivideFunctionTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DivideFunction.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DivideFunction *DivideFunctionTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DivideFunction.Contract.contract.Transact(opts, method, params...)
}

// WeiToKfive is a free data retrieval call binding the contract method 0xd7a65f0b.
//
// Solidity: function weiToKfive(uint256 weiValue, uint256 rate) pure returns(uint256)
func (_DivideFunction *DivideFunctionCaller) WeiToKfive(opts *bind.CallOpts, weiValue *big.Int, rate *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _DivideFunction.contract.Call(opts, &out, "weiToKfive", weiValue, rate)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WeiToKfive is a free data retrieval call binding the contract method 0xd7a65f0b.
//
// Solidity: function weiToKfive(uint256 weiValue, uint256 rate) pure returns(uint256)
func (_DivideFunction *DivideFunctionSession) WeiToKfive(weiValue *big.Int, rate *big.Int) (*big.Int, error) {
	return _DivideFunction.Contract.WeiToKfive(&_DivideFunction.CallOpts, weiValue, rate)
}

// WeiToKfive is a free data retrieval call binding the contract method 0xd7a65f0b.
//
// Solidity: function weiToKfive(uint256 weiValue, uint256 rate) pure returns(uint256)
func (_DivideFunction *DivideFunctionCallerSession) WeiToKfive(weiValue *big.Int, rate *big.Int) (*big.Int, error) {
	return _DivideFunction.Contract.WeiToKfive(&_DivideFunction.CallOpts, weiValue, rate)
}
