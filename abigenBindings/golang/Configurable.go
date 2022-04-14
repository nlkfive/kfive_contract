// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package Configurable

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

// ConfigurableMetaData contains all meta data concerning the Configurable contract.
var ConfigurableMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"KFIVE2MTC\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"cap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// ConfigurableABI is the input ABI used to generate the binding from.
// Deprecated: Use ConfigurableMetaData.ABI instead.
var ConfigurableABI = ConfigurableMetaData.ABI

// Configurable is an auto generated Go binding around an Ethereum contract.
type Configurable struct {
	ConfigurableCaller     // Read-only binding to the contract
	ConfigurableTransactor // Write-only binding to the contract
	ConfigurableFilterer   // Log filterer for contract events
}

// ConfigurableCaller is an auto generated read-only Go binding around an Ethereum contract.
type ConfigurableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ConfigurableTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ConfigurableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ConfigurableFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ConfigurableFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ConfigurableSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ConfigurableSession struct {
	Contract     *Configurable     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ConfigurableCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ConfigurableCallerSession struct {
	Contract *ConfigurableCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// ConfigurableTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ConfigurableTransactorSession struct {
	Contract     *ConfigurableTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// ConfigurableRaw is an auto generated low-level Go binding around an Ethereum contract.
type ConfigurableRaw struct {
	Contract *Configurable // Generic contract binding to access the raw methods on
}

// ConfigurableCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ConfigurableCallerRaw struct {
	Contract *ConfigurableCaller // Generic read-only contract binding to access the raw methods on
}

// ConfigurableTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ConfigurableTransactorRaw struct {
	Contract *ConfigurableTransactor // Generic write-only contract binding to access the raw methods on
}

// NewConfigurable creates a new instance of Configurable, bound to a specific deployed contract.
func NewConfigurable(address common.Address, backend bind.ContractBackend) (*Configurable, error) {
	contract, err := bindConfigurable(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Configurable{ConfigurableCaller: ConfigurableCaller{contract: contract}, ConfigurableTransactor: ConfigurableTransactor{contract: contract}, ConfigurableFilterer: ConfigurableFilterer{contract: contract}}, nil
}

// NewConfigurableCaller creates a new read-only instance of Configurable, bound to a specific deployed contract.
func NewConfigurableCaller(address common.Address, caller bind.ContractCaller) (*ConfigurableCaller, error) {
	contract, err := bindConfigurable(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ConfigurableCaller{contract: contract}, nil
}

// NewConfigurableTransactor creates a new write-only instance of Configurable, bound to a specific deployed contract.
func NewConfigurableTransactor(address common.Address, transactor bind.ContractTransactor) (*ConfigurableTransactor, error) {
	contract, err := bindConfigurable(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ConfigurableTransactor{contract: contract}, nil
}

// NewConfigurableFilterer creates a new log filterer instance of Configurable, bound to a specific deployed contract.
func NewConfigurableFilterer(address common.Address, filterer bind.ContractFilterer) (*ConfigurableFilterer, error) {
	contract, err := bindConfigurable(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ConfigurableFilterer{contract: contract}, nil
}

// bindConfigurable binds a generic wrapper to an already deployed contract.
func bindConfigurable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ConfigurableABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Configurable *ConfigurableRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Configurable.Contract.ConfigurableCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Configurable *ConfigurableRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Configurable.Contract.ConfigurableTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Configurable *ConfigurableRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Configurable.Contract.ConfigurableTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Configurable *ConfigurableCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Configurable.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Configurable *ConfigurableTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Configurable.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Configurable *ConfigurableTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Configurable.Contract.contract.Transact(opts, method, params...)
}

// KFIVE2MTC is a free data retrieval call binding the contract method 0x62e8db23.
//
// Solidity: function KFIVE2MTC() view returns(uint256)
func (_Configurable *ConfigurableCaller) KFIVE2MTC(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Configurable.contract.Call(opts, &out, "KFIVE2MTC")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// KFIVE2MTC is a free data retrieval call binding the contract method 0x62e8db23.
//
// Solidity: function KFIVE2MTC() view returns(uint256)
func (_Configurable *ConfigurableSession) KFIVE2MTC() (*big.Int, error) {
	return _Configurable.Contract.KFIVE2MTC(&_Configurable.CallOpts)
}

// KFIVE2MTC is a free data retrieval call binding the contract method 0x62e8db23.
//
// Solidity: function KFIVE2MTC() view returns(uint256)
func (_Configurable *ConfigurableCallerSession) KFIVE2MTC() (*big.Int, error) {
	return _Configurable.Contract.KFIVE2MTC(&_Configurable.CallOpts)
}

// Cap is a free data retrieval call binding the contract method 0x355274ea.
//
// Solidity: function cap() view returns(uint256)
func (_Configurable *ConfigurableCaller) Cap(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Configurable.contract.Call(opts, &out, "cap")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Cap is a free data retrieval call binding the contract method 0x355274ea.
//
// Solidity: function cap() view returns(uint256)
func (_Configurable *ConfigurableSession) Cap() (*big.Int, error) {
	return _Configurable.Contract.Cap(&_Configurable.CallOpts)
}

// Cap is a free data retrieval call binding the contract method 0x355274ea.
//
// Solidity: function cap() view returns(uint256)
func (_Configurable *ConfigurableCallerSession) Cap() (*big.Int, error) {
	return _Configurable.Contract.Cap(&_Configurable.CallOpts)
}
