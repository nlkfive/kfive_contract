// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package EIP712Base

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

// EIP712BaseMetaData contains all meta data concerning the EIP712Base contract.
var EIP712BaseMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"domainSeparator\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getChainId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// EIP712BaseABI is the input ABI used to generate the binding from.
// Deprecated: Use EIP712BaseMetaData.ABI instead.
var EIP712BaseABI = EIP712BaseMetaData.ABI

// EIP712Base is an auto generated Go binding around an Ethereum contract.
type EIP712Base struct {
	EIP712BaseCaller     // Read-only binding to the contract
	EIP712BaseTransactor // Write-only binding to the contract
	EIP712BaseFilterer   // Log filterer for contract events
}

// EIP712BaseCaller is an auto generated read-only Go binding around an Ethereum contract.
type EIP712BaseCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EIP712BaseTransactor is an auto generated write-only Go binding around an Ethereum contract.
type EIP712BaseTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EIP712BaseFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type EIP712BaseFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EIP712BaseSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type EIP712BaseSession struct {
	Contract     *EIP712Base       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// EIP712BaseCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type EIP712BaseCallerSession struct {
	Contract *EIP712BaseCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// EIP712BaseTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type EIP712BaseTransactorSession struct {
	Contract     *EIP712BaseTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// EIP712BaseRaw is an auto generated low-level Go binding around an Ethereum contract.
type EIP712BaseRaw struct {
	Contract *EIP712Base // Generic contract binding to access the raw methods on
}

// EIP712BaseCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type EIP712BaseCallerRaw struct {
	Contract *EIP712BaseCaller // Generic read-only contract binding to access the raw methods on
}

// EIP712BaseTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type EIP712BaseTransactorRaw struct {
	Contract *EIP712BaseTransactor // Generic write-only contract binding to access the raw methods on
}

// NewEIP712Base creates a new instance of EIP712Base, bound to a specific deployed contract.
func NewEIP712Base(address common.Address, backend bind.ContractBackend) (*EIP712Base, error) {
	contract, err := bindEIP712Base(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &EIP712Base{EIP712BaseCaller: EIP712BaseCaller{contract: contract}, EIP712BaseTransactor: EIP712BaseTransactor{contract: contract}, EIP712BaseFilterer: EIP712BaseFilterer{contract: contract}}, nil
}

// NewEIP712BaseCaller creates a new read-only instance of EIP712Base, bound to a specific deployed contract.
func NewEIP712BaseCaller(address common.Address, caller bind.ContractCaller) (*EIP712BaseCaller, error) {
	contract, err := bindEIP712Base(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &EIP712BaseCaller{contract: contract}, nil
}

// NewEIP712BaseTransactor creates a new write-only instance of EIP712Base, bound to a specific deployed contract.
func NewEIP712BaseTransactor(address common.Address, transactor bind.ContractTransactor) (*EIP712BaseTransactor, error) {
	contract, err := bindEIP712Base(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &EIP712BaseTransactor{contract: contract}, nil
}

// NewEIP712BaseFilterer creates a new log filterer instance of EIP712Base, bound to a specific deployed contract.
func NewEIP712BaseFilterer(address common.Address, filterer bind.ContractFilterer) (*EIP712BaseFilterer, error) {
	contract, err := bindEIP712Base(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &EIP712BaseFilterer{contract: contract}, nil
}

// bindEIP712Base binds a generic wrapper to an already deployed contract.
func bindEIP712Base(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(EIP712BaseABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EIP712Base *EIP712BaseRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EIP712Base.Contract.EIP712BaseCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EIP712Base *EIP712BaseRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EIP712Base.Contract.EIP712BaseTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EIP712Base *EIP712BaseRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EIP712Base.Contract.EIP712BaseTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EIP712Base *EIP712BaseCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EIP712Base.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EIP712Base *EIP712BaseTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EIP712Base.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EIP712Base *EIP712BaseTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EIP712Base.Contract.contract.Transact(opts, method, params...)
}

// DomainSeparator is a free data retrieval call binding the contract method 0xf698da25.
//
// Solidity: function domainSeparator() view returns(bytes32)
func (_EIP712Base *EIP712BaseCaller) DomainSeparator(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _EIP712Base.contract.Call(opts, &out, "domainSeparator")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DomainSeparator is a free data retrieval call binding the contract method 0xf698da25.
//
// Solidity: function domainSeparator() view returns(bytes32)
func (_EIP712Base *EIP712BaseSession) DomainSeparator() ([32]byte, error) {
	return _EIP712Base.Contract.DomainSeparator(&_EIP712Base.CallOpts)
}

// DomainSeparator is a free data retrieval call binding the contract method 0xf698da25.
//
// Solidity: function domainSeparator() view returns(bytes32)
func (_EIP712Base *EIP712BaseCallerSession) DomainSeparator() ([32]byte, error) {
	return _EIP712Base.Contract.DomainSeparator(&_EIP712Base.CallOpts)
}

// GetChainId is a free data retrieval call binding the contract method 0x3408e470.
//
// Solidity: function getChainId() view returns(uint256)
func (_EIP712Base *EIP712BaseCaller) GetChainId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _EIP712Base.contract.Call(opts, &out, "getChainId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetChainId is a free data retrieval call binding the contract method 0x3408e470.
//
// Solidity: function getChainId() view returns(uint256)
func (_EIP712Base *EIP712BaseSession) GetChainId() (*big.Int, error) {
	return _EIP712Base.Contract.GetChainId(&_EIP712Base.CallOpts)
}

// GetChainId is a free data retrieval call binding the contract method 0x3408e470.
//
// Solidity: function getChainId() view returns(uint256)
func (_EIP712Base *EIP712BaseCallerSession) GetChainId() (*big.Int, error) {
	return _EIP712Base.Contract.GetChainId(&_EIP712Base.CallOpts)
}
