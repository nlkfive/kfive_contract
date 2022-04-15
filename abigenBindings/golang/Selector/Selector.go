// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package Selector

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

// SelectorMetaData contains all meta data concerning the Selector contract.
var SelectorMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"calcStoreInterfaceId\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]",
}

// SelectorABI is the input ABI used to generate the binding from.
// Deprecated: Use SelectorMetaData.ABI instead.
var SelectorABI = SelectorMetaData.ABI

// Selector is an auto generated Go binding around an Ethereum contract.
type Selector struct {
	SelectorCaller     // Read-only binding to the contract
	SelectorTransactor // Write-only binding to the contract
	SelectorFilterer   // Log filterer for contract events
}

// SelectorCaller is an auto generated read-only Go binding around an Ethereum contract.
type SelectorCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SelectorTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SelectorTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SelectorFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SelectorFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SelectorSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SelectorSession struct {
	Contract     *Selector         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SelectorCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SelectorCallerSession struct {
	Contract *SelectorCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// SelectorTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SelectorTransactorSession struct {
	Contract     *SelectorTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// SelectorRaw is an auto generated low-level Go binding around an Ethereum contract.
type SelectorRaw struct {
	Contract *Selector // Generic contract binding to access the raw methods on
}

// SelectorCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SelectorCallerRaw struct {
	Contract *SelectorCaller // Generic read-only contract binding to access the raw methods on
}

// SelectorTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SelectorTransactorRaw struct {
	Contract *SelectorTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSelector creates a new instance of Selector, bound to a specific deployed contract.
func NewSelector(address common.Address, backend bind.ContractBackend) (*Selector, error) {
	contract, err := bindSelector(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Selector{SelectorCaller: SelectorCaller{contract: contract}, SelectorTransactor: SelectorTransactor{contract: contract}, SelectorFilterer: SelectorFilterer{contract: contract}}, nil
}

// NewSelectorCaller creates a new read-only instance of Selector, bound to a specific deployed contract.
func NewSelectorCaller(address common.Address, caller bind.ContractCaller) (*SelectorCaller, error) {
	contract, err := bindSelector(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SelectorCaller{contract: contract}, nil
}

// NewSelectorTransactor creates a new write-only instance of Selector, bound to a specific deployed contract.
func NewSelectorTransactor(address common.Address, transactor bind.ContractTransactor) (*SelectorTransactor, error) {
	contract, err := bindSelector(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SelectorTransactor{contract: contract}, nil
}

// NewSelectorFilterer creates a new log filterer instance of Selector, bound to a specific deployed contract.
func NewSelectorFilterer(address common.Address, filterer bind.ContractFilterer) (*SelectorFilterer, error) {
	contract, err := bindSelector(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SelectorFilterer{contract: contract}, nil
}

// bindSelector binds a generic wrapper to an already deployed contract.
func bindSelector(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SelectorABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Selector *SelectorRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Selector.Contract.SelectorCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Selector *SelectorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Selector.Contract.SelectorTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Selector *SelectorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Selector.Contract.SelectorTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Selector *SelectorCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Selector.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Selector *SelectorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Selector.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Selector *SelectorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Selector.Contract.contract.Transact(opts, method, params...)
}

// CalcStoreInterfaceId is a free data retrieval call binding the contract method 0x63d5af8d.
//
// Solidity: function calcStoreInterfaceId() pure returns(bytes4)
func (_Selector *SelectorCaller) CalcStoreInterfaceId(opts *bind.CallOpts) ([4]byte, error) {
	var out []interface{}
	err := _Selector.contract.Call(opts, &out, "calcStoreInterfaceId")

	if err != nil {
		return *new([4]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)

	return out0, err

}

// CalcStoreInterfaceId is a free data retrieval call binding the contract method 0x63d5af8d.
//
// Solidity: function calcStoreInterfaceId() pure returns(bytes4)
func (_Selector *SelectorSession) CalcStoreInterfaceId() ([4]byte, error) {
	return _Selector.Contract.CalcStoreInterfaceId(&_Selector.CallOpts)
}

// CalcStoreInterfaceId is a free data retrieval call binding the contract method 0x63d5af8d.
//
// Solidity: function calcStoreInterfaceId() pure returns(bytes4)
func (_Selector *SelectorCallerSession) CalcStoreInterfaceId() ([4]byte, error) {
	return _Selector.Contract.CalcStoreInterfaceId(&_Selector.CallOpts)
}
