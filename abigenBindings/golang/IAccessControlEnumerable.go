// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package IAccessControlEnumerable

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

// IAccessControlEnumerableMetaData contains all meta data concerning the IAccessControlEnumerable contract.
var IAccessControlEnumerableMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getRoleMember\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleMemberCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// IAccessControlEnumerableABI is the input ABI used to generate the binding from.
// Deprecated: Use IAccessControlEnumerableMetaData.ABI instead.
var IAccessControlEnumerableABI = IAccessControlEnumerableMetaData.ABI

// IAccessControlEnumerable is an auto generated Go binding around an Ethereum contract.
type IAccessControlEnumerable struct {
	IAccessControlEnumerableCaller     // Read-only binding to the contract
	IAccessControlEnumerableTransactor // Write-only binding to the contract
	IAccessControlEnumerableFilterer   // Log filterer for contract events
}

// IAccessControlEnumerableCaller is an auto generated read-only Go binding around an Ethereum contract.
type IAccessControlEnumerableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IAccessControlEnumerableTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IAccessControlEnumerableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IAccessControlEnumerableFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IAccessControlEnumerableFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IAccessControlEnumerableSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IAccessControlEnumerableSession struct {
	Contract     *IAccessControlEnumerable // Generic contract binding to set the session for
	CallOpts     bind.CallOpts             // Call options to use throughout this session
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// IAccessControlEnumerableCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IAccessControlEnumerableCallerSession struct {
	Contract *IAccessControlEnumerableCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                   // Call options to use throughout this session
}

// IAccessControlEnumerableTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IAccessControlEnumerableTransactorSession struct {
	Contract     *IAccessControlEnumerableTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                   // Transaction auth options to use throughout this session
}

// IAccessControlEnumerableRaw is an auto generated low-level Go binding around an Ethereum contract.
type IAccessControlEnumerableRaw struct {
	Contract *IAccessControlEnumerable // Generic contract binding to access the raw methods on
}

// IAccessControlEnumerableCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IAccessControlEnumerableCallerRaw struct {
	Contract *IAccessControlEnumerableCaller // Generic read-only contract binding to access the raw methods on
}

// IAccessControlEnumerableTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IAccessControlEnumerableTransactorRaw struct {
	Contract *IAccessControlEnumerableTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIAccessControlEnumerable creates a new instance of IAccessControlEnumerable, bound to a specific deployed contract.
func NewIAccessControlEnumerable(address common.Address, backend bind.ContractBackend) (*IAccessControlEnumerable, error) {
	contract, err := bindIAccessControlEnumerable(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IAccessControlEnumerable{IAccessControlEnumerableCaller: IAccessControlEnumerableCaller{contract: contract}, IAccessControlEnumerableTransactor: IAccessControlEnumerableTransactor{contract: contract}, IAccessControlEnumerableFilterer: IAccessControlEnumerableFilterer{contract: contract}}, nil
}

// NewIAccessControlEnumerableCaller creates a new read-only instance of IAccessControlEnumerable, bound to a specific deployed contract.
func NewIAccessControlEnumerableCaller(address common.Address, caller bind.ContractCaller) (*IAccessControlEnumerableCaller, error) {
	contract, err := bindIAccessControlEnumerable(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IAccessControlEnumerableCaller{contract: contract}, nil
}

// NewIAccessControlEnumerableTransactor creates a new write-only instance of IAccessControlEnumerable, bound to a specific deployed contract.
func NewIAccessControlEnumerableTransactor(address common.Address, transactor bind.ContractTransactor) (*IAccessControlEnumerableTransactor, error) {
	contract, err := bindIAccessControlEnumerable(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IAccessControlEnumerableTransactor{contract: contract}, nil
}

// NewIAccessControlEnumerableFilterer creates a new log filterer instance of IAccessControlEnumerable, bound to a specific deployed contract.
func NewIAccessControlEnumerableFilterer(address common.Address, filterer bind.ContractFilterer) (*IAccessControlEnumerableFilterer, error) {
	contract, err := bindIAccessControlEnumerable(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IAccessControlEnumerableFilterer{contract: contract}, nil
}

// bindIAccessControlEnumerable binds a generic wrapper to an already deployed contract.
func bindIAccessControlEnumerable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IAccessControlEnumerableABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IAccessControlEnumerable *IAccessControlEnumerableRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IAccessControlEnumerable.Contract.IAccessControlEnumerableCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IAccessControlEnumerable *IAccessControlEnumerableRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IAccessControlEnumerable.Contract.IAccessControlEnumerableTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IAccessControlEnumerable *IAccessControlEnumerableRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IAccessControlEnumerable.Contract.IAccessControlEnumerableTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IAccessControlEnumerable *IAccessControlEnumerableCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IAccessControlEnumerable.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IAccessControlEnumerable *IAccessControlEnumerableTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IAccessControlEnumerable.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IAccessControlEnumerable *IAccessControlEnumerableTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IAccessControlEnumerable.Contract.contract.Transact(opts, method, params...)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_IAccessControlEnumerable *IAccessControlEnumerableCaller) GetRoleMember(opts *bind.CallOpts, role [32]byte, index *big.Int) (common.Address, error) {
	var out []interface{}
	err := _IAccessControlEnumerable.contract.Call(opts, &out, "getRoleMember", role, index)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_IAccessControlEnumerable *IAccessControlEnumerableSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _IAccessControlEnumerable.Contract.GetRoleMember(&_IAccessControlEnumerable.CallOpts, role, index)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_IAccessControlEnumerable *IAccessControlEnumerableCallerSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _IAccessControlEnumerable.Contract.GetRoleMember(&_IAccessControlEnumerable.CallOpts, role, index)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_IAccessControlEnumerable *IAccessControlEnumerableCaller) GetRoleMemberCount(opts *bind.CallOpts, role [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _IAccessControlEnumerable.contract.Call(opts, &out, "getRoleMemberCount", role)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_IAccessControlEnumerable *IAccessControlEnumerableSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _IAccessControlEnumerable.Contract.GetRoleMemberCount(&_IAccessControlEnumerable.CallOpts, role)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_IAccessControlEnumerable *IAccessControlEnumerableCallerSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _IAccessControlEnumerable.Contract.GetRoleMemberCount(&_IAccessControlEnumerable.CallOpts, role)
}
