// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package TokenAdmin

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

// TokenAdminMetaData contains all meta data concerning the TokenAdmin contract.
var TokenAdminMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"admin\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"status\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"maxIssuingTokenPerTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxTotalIssuingToken\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"remainingIssuingToken\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"a\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"maxIssuingTokenPerTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxTotalIssuingToken\",\"type\":\"uint256\"}],\"name\":\"addAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"a\",\"type\":\"address\"}],\"name\":\"removeAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// TokenAdminABI is the input ABI used to generate the binding from.
// Deprecated: Use TokenAdminMetaData.ABI instead.
var TokenAdminABI = TokenAdminMetaData.ABI

// TokenAdmin is an auto generated Go binding around an Ethereum contract.
type TokenAdmin struct {
	TokenAdminCaller     // Read-only binding to the contract
	TokenAdminTransactor // Write-only binding to the contract
	TokenAdminFilterer   // Log filterer for contract events
}

// TokenAdminCaller is an auto generated read-only Go binding around an Ethereum contract.
type TokenAdminCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenAdminTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TokenAdminTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenAdminFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TokenAdminFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenAdminSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TokenAdminSession struct {
	Contract     *TokenAdmin       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TokenAdminCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TokenAdminCallerSession struct {
	Contract *TokenAdminCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// TokenAdminTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TokenAdminTransactorSession struct {
	Contract     *TokenAdminTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// TokenAdminRaw is an auto generated low-level Go binding around an Ethereum contract.
type TokenAdminRaw struct {
	Contract *TokenAdmin // Generic contract binding to access the raw methods on
}

// TokenAdminCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TokenAdminCallerRaw struct {
	Contract *TokenAdminCaller // Generic read-only contract binding to access the raw methods on
}

// TokenAdminTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TokenAdminTransactorRaw struct {
	Contract *TokenAdminTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTokenAdmin creates a new instance of TokenAdmin, bound to a specific deployed contract.
func NewTokenAdmin(address common.Address, backend bind.ContractBackend) (*TokenAdmin, error) {
	contract, err := bindTokenAdmin(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TokenAdmin{TokenAdminCaller: TokenAdminCaller{contract: contract}, TokenAdminTransactor: TokenAdminTransactor{contract: contract}, TokenAdminFilterer: TokenAdminFilterer{contract: contract}}, nil
}

// NewTokenAdminCaller creates a new read-only instance of TokenAdmin, bound to a specific deployed contract.
func NewTokenAdminCaller(address common.Address, caller bind.ContractCaller) (*TokenAdminCaller, error) {
	contract, err := bindTokenAdmin(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TokenAdminCaller{contract: contract}, nil
}

// NewTokenAdminTransactor creates a new write-only instance of TokenAdmin, bound to a specific deployed contract.
func NewTokenAdminTransactor(address common.Address, transactor bind.ContractTransactor) (*TokenAdminTransactor, error) {
	contract, err := bindTokenAdmin(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TokenAdminTransactor{contract: contract}, nil
}

// NewTokenAdminFilterer creates a new log filterer instance of TokenAdmin, bound to a specific deployed contract.
func NewTokenAdminFilterer(address common.Address, filterer bind.ContractFilterer) (*TokenAdminFilterer, error) {
	contract, err := bindTokenAdmin(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TokenAdminFilterer{contract: contract}, nil
}

// bindTokenAdmin binds a generic wrapper to an already deployed contract.
func bindTokenAdmin(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TokenAdminABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TokenAdmin *TokenAdminRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TokenAdmin.Contract.TokenAdminCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TokenAdmin *TokenAdminRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TokenAdmin.Contract.TokenAdminTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TokenAdmin *TokenAdminRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TokenAdmin.Contract.TokenAdminTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TokenAdmin *TokenAdminCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TokenAdmin.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TokenAdmin *TokenAdminTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TokenAdmin.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TokenAdmin *TokenAdminTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TokenAdmin.Contract.contract.Transact(opts, method, params...)
}

// Admin is a free data retrieval call binding the contract method 0x63a846f8.
//
// Solidity: function admin(address ) view returns(bool status, uint256 maxIssuingTokenPerTime, uint256 maxTotalIssuingToken, uint256 remainingIssuingToken)
func (_TokenAdmin *TokenAdminCaller) Admin(opts *bind.CallOpts, arg0 common.Address) (struct {
	Status                 bool
	MaxIssuingTokenPerTime *big.Int
	MaxTotalIssuingToken   *big.Int
	RemainingIssuingToken  *big.Int
}, error) {
	var out []interface{}
	err := _TokenAdmin.contract.Call(opts, &out, "admin", arg0)

	outstruct := new(struct {
		Status                 bool
		MaxIssuingTokenPerTime *big.Int
		MaxTotalIssuingToken   *big.Int
		RemainingIssuingToken  *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Status = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.MaxIssuingTokenPerTime = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.MaxTotalIssuingToken = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.RemainingIssuingToken = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Admin is a free data retrieval call binding the contract method 0x63a846f8.
//
// Solidity: function admin(address ) view returns(bool status, uint256 maxIssuingTokenPerTime, uint256 maxTotalIssuingToken, uint256 remainingIssuingToken)
func (_TokenAdmin *TokenAdminSession) Admin(arg0 common.Address) (struct {
	Status                 bool
	MaxIssuingTokenPerTime *big.Int
	MaxTotalIssuingToken   *big.Int
	RemainingIssuingToken  *big.Int
}, error) {
	return _TokenAdmin.Contract.Admin(&_TokenAdmin.CallOpts, arg0)
}

// Admin is a free data retrieval call binding the contract method 0x63a846f8.
//
// Solidity: function admin(address ) view returns(bool status, uint256 maxIssuingTokenPerTime, uint256 maxTotalIssuingToken, uint256 remainingIssuingToken)
func (_TokenAdmin *TokenAdminCallerSession) Admin(arg0 common.Address) (struct {
	Status                 bool
	MaxIssuingTokenPerTime *big.Int
	MaxTotalIssuingToken   *big.Int
	RemainingIssuingToken  *big.Int
}, error) {
	return _TokenAdmin.Contract.Admin(&_TokenAdmin.CallOpts, arg0)
}

// AddAdmin is a paid mutator transaction binding the contract method 0x46f2ed2e.
//
// Solidity: function addAdmin(address a, uint256 maxIssuingTokenPerTime, uint256 maxTotalIssuingToken) returns()
func (_TokenAdmin *TokenAdminTransactor) AddAdmin(opts *bind.TransactOpts, a common.Address, maxIssuingTokenPerTime *big.Int, maxTotalIssuingToken *big.Int) (*types.Transaction, error) {
	return _TokenAdmin.contract.Transact(opts, "addAdmin", a, maxIssuingTokenPerTime, maxTotalIssuingToken)
}

// AddAdmin is a paid mutator transaction binding the contract method 0x46f2ed2e.
//
// Solidity: function addAdmin(address a, uint256 maxIssuingTokenPerTime, uint256 maxTotalIssuingToken) returns()
func (_TokenAdmin *TokenAdminSession) AddAdmin(a common.Address, maxIssuingTokenPerTime *big.Int, maxTotalIssuingToken *big.Int) (*types.Transaction, error) {
	return _TokenAdmin.Contract.AddAdmin(&_TokenAdmin.TransactOpts, a, maxIssuingTokenPerTime, maxTotalIssuingToken)
}

// AddAdmin is a paid mutator transaction binding the contract method 0x46f2ed2e.
//
// Solidity: function addAdmin(address a, uint256 maxIssuingTokenPerTime, uint256 maxTotalIssuingToken) returns()
func (_TokenAdmin *TokenAdminTransactorSession) AddAdmin(a common.Address, maxIssuingTokenPerTime *big.Int, maxTotalIssuingToken *big.Int) (*types.Transaction, error) {
	return _TokenAdmin.Contract.AddAdmin(&_TokenAdmin.TransactOpts, a, maxIssuingTokenPerTime, maxTotalIssuingToken)
}

// RemoveAdmin is a paid mutator transaction binding the contract method 0x1785f53c.
//
// Solidity: function removeAdmin(address a) returns()
func (_TokenAdmin *TokenAdminTransactor) RemoveAdmin(opts *bind.TransactOpts, a common.Address) (*types.Transaction, error) {
	return _TokenAdmin.contract.Transact(opts, "removeAdmin", a)
}

// RemoveAdmin is a paid mutator transaction binding the contract method 0x1785f53c.
//
// Solidity: function removeAdmin(address a) returns()
func (_TokenAdmin *TokenAdminSession) RemoveAdmin(a common.Address) (*types.Transaction, error) {
	return _TokenAdmin.Contract.RemoveAdmin(&_TokenAdmin.TransactOpts, a)
}

// RemoveAdmin is a paid mutator transaction binding the contract method 0x1785f53c.
//
// Solidity: function removeAdmin(address a) returns()
func (_TokenAdmin *TokenAdminTransactorSession) RemoveAdmin(a common.Address) (*types.Transaction, error) {
	return _TokenAdmin.Contract.RemoveAdmin(&_TokenAdmin.TransactOpts, a)
}
