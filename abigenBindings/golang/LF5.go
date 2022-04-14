// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package LF5

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

// LF5MetaData contains all meta data concerning the LF5 contract.
var LF5MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"AddedBlackList\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_blackListedUser\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_balance\",\"type\":\"uint256\"}],\"name\":\"DestroyedBlackFunds\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"RemovedBlackList\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"offchain\",\"type\":\"bytes32\"}],\"name\":\"__issue\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"offchain\",\"type\":\"bytes32\"}],\"name\":\"__redeem\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"offchain\",\"type\":\"bytes32\"}],\"name\":\"__transferByAdmin\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"a\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"maxIssuingTokenPerTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxTotalIssuingToken\",\"type\":\"uint256\"}],\"name\":\"addAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_evilUser\",\"type\":\"address\"}],\"name\":\"addBlackList\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"admin\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"status\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"maxIssuingTokenPerTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxTotalIssuingToken\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"remainingIssuingToken\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_maker\",\"type\":\"address\"}],\"name\":\"getBlackListStatus\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"isBlackListed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"a\",\"type\":\"address\"}],\"name\":\"removeAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_clearedUser\",\"type\":\"address\"}],\"name\":\"removeBlackList\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"offchain\",\"type\":\"bytes32\"}],\"name\":\"transferByAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"issuer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"offchain\",\"type\":\"bytes32\"}],\"name\":\"issue\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"issuer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"offchain\",\"type\":\"bytes32\"}],\"name\":\"issueByAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_blackListedUser\",\"type\":\"address\"}],\"name\":\"destroyBlackFunds\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"redeemer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"offchain\",\"type\":\"bytes32\"}],\"name\":\"redeem\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// LF5ABI is the input ABI used to generate the binding from.
// Deprecated: Use LF5MetaData.ABI instead.
var LF5ABI = LF5MetaData.ABI

// LF5 is an auto generated Go binding around an Ethereum contract.
type LF5 struct {
	LF5Caller     // Read-only binding to the contract
	LF5Transactor // Write-only binding to the contract
	LF5Filterer   // Log filterer for contract events
}

// LF5Caller is an auto generated read-only Go binding around an Ethereum contract.
type LF5Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LF5Transactor is an auto generated write-only Go binding around an Ethereum contract.
type LF5Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LF5Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type LF5Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LF5Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type LF5Session struct {
	Contract     *LF5              // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// LF5CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type LF5CallerSession struct {
	Contract *LF5Caller    // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// LF5TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type LF5TransactorSession struct {
	Contract     *LF5Transactor    // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// LF5Raw is an auto generated low-level Go binding around an Ethereum contract.
type LF5Raw struct {
	Contract *LF5 // Generic contract binding to access the raw methods on
}

// LF5CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type LF5CallerRaw struct {
	Contract *LF5Caller // Generic read-only contract binding to access the raw methods on
}

// LF5TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type LF5TransactorRaw struct {
	Contract *LF5Transactor // Generic write-only contract binding to access the raw methods on
}

// NewLF5 creates a new instance of LF5, bound to a specific deployed contract.
func NewLF5(address common.Address, backend bind.ContractBackend) (*LF5, error) {
	contract, err := bindLF5(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &LF5{LF5Caller: LF5Caller{contract: contract}, LF5Transactor: LF5Transactor{contract: contract}, LF5Filterer: LF5Filterer{contract: contract}}, nil
}

// NewLF5Caller creates a new read-only instance of LF5, bound to a specific deployed contract.
func NewLF5Caller(address common.Address, caller bind.ContractCaller) (*LF5Caller, error) {
	contract, err := bindLF5(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &LF5Caller{contract: contract}, nil
}

// NewLF5Transactor creates a new write-only instance of LF5, bound to a specific deployed contract.
func NewLF5Transactor(address common.Address, transactor bind.ContractTransactor) (*LF5Transactor, error) {
	contract, err := bindLF5(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &LF5Transactor{contract: contract}, nil
}

// NewLF5Filterer creates a new log filterer instance of LF5, bound to a specific deployed contract.
func NewLF5Filterer(address common.Address, filterer bind.ContractFilterer) (*LF5Filterer, error) {
	contract, err := bindLF5(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &LF5Filterer{contract: contract}, nil
}

// bindLF5 binds a generic wrapper to an already deployed contract.
func bindLF5(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(LF5ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LF5 *LF5Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LF5.Contract.LF5Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LF5 *LF5Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LF5.Contract.LF5Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LF5 *LF5Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LF5.Contract.LF5Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LF5 *LF5CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LF5.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LF5 *LF5TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LF5.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LF5 *LF5TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LF5.Contract.contract.Transact(opts, method, params...)
}

// Admin is a free data retrieval call binding the contract method 0x63a846f8.
//
// Solidity: function admin(address ) view returns(bool status, uint256 maxIssuingTokenPerTime, uint256 maxTotalIssuingToken, uint256 remainingIssuingToken)
func (_LF5 *LF5Caller) Admin(opts *bind.CallOpts, arg0 common.Address) (struct {
	Status                 bool
	MaxIssuingTokenPerTime *big.Int
	MaxTotalIssuingToken   *big.Int
	RemainingIssuingToken  *big.Int
}, error) {
	var out []interface{}
	err := _LF5.contract.Call(opts, &out, "admin", arg0)

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
func (_LF5 *LF5Session) Admin(arg0 common.Address) (struct {
	Status                 bool
	MaxIssuingTokenPerTime *big.Int
	MaxTotalIssuingToken   *big.Int
	RemainingIssuingToken  *big.Int
}, error) {
	return _LF5.Contract.Admin(&_LF5.CallOpts, arg0)
}

// Admin is a free data retrieval call binding the contract method 0x63a846f8.
//
// Solidity: function admin(address ) view returns(bool status, uint256 maxIssuingTokenPerTime, uint256 maxTotalIssuingToken, uint256 remainingIssuingToken)
func (_LF5 *LF5CallerSession) Admin(arg0 common.Address) (struct {
	Status                 bool
	MaxIssuingTokenPerTime *big.Int
	MaxTotalIssuingToken   *big.Int
	RemainingIssuingToken  *big.Int
}, error) {
	return _LF5.Contract.Admin(&_LF5.CallOpts, arg0)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_LF5 *LF5Caller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _LF5.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_LF5 *LF5Session) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _LF5.Contract.Allowance(&_LF5.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_LF5 *LF5CallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _LF5.Contract.Allowance(&_LF5.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_LF5 *LF5Caller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _LF5.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_LF5 *LF5Session) BalanceOf(account common.Address) (*big.Int, error) {
	return _LF5.Contract.BalanceOf(&_LF5.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_LF5 *LF5CallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _LF5.Contract.BalanceOf(&_LF5.CallOpts, account)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_LF5 *LF5Caller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _LF5.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_LF5 *LF5Session) Decimals() (uint8, error) {
	return _LF5.Contract.Decimals(&_LF5.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_LF5 *LF5CallerSession) Decimals() (uint8, error) {
	return _LF5.Contract.Decimals(&_LF5.CallOpts)
}

// GetBlackListStatus is a free data retrieval call binding the contract method 0x59bf1abe.
//
// Solidity: function getBlackListStatus(address _maker) view returns(bool)
func (_LF5 *LF5Caller) GetBlackListStatus(opts *bind.CallOpts, _maker common.Address) (bool, error) {
	var out []interface{}
	err := _LF5.contract.Call(opts, &out, "getBlackListStatus", _maker)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// GetBlackListStatus is a free data retrieval call binding the contract method 0x59bf1abe.
//
// Solidity: function getBlackListStatus(address _maker) view returns(bool)
func (_LF5 *LF5Session) GetBlackListStatus(_maker common.Address) (bool, error) {
	return _LF5.Contract.GetBlackListStatus(&_LF5.CallOpts, _maker)
}

// GetBlackListStatus is a free data retrieval call binding the contract method 0x59bf1abe.
//
// Solidity: function getBlackListStatus(address _maker) view returns(bool)
func (_LF5 *LF5CallerSession) GetBlackListStatus(_maker common.Address) (bool, error) {
	return _LF5.Contract.GetBlackListStatus(&_LF5.CallOpts, _maker)
}

// GetOwner is a free data retrieval call binding the contract method 0x893d20e8.
//
// Solidity: function getOwner() view returns(address)
func (_LF5 *LF5Caller) GetOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LF5.contract.Call(opts, &out, "getOwner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetOwner is a free data retrieval call binding the contract method 0x893d20e8.
//
// Solidity: function getOwner() view returns(address)
func (_LF5 *LF5Session) GetOwner() (common.Address, error) {
	return _LF5.Contract.GetOwner(&_LF5.CallOpts)
}

// GetOwner is a free data retrieval call binding the contract method 0x893d20e8.
//
// Solidity: function getOwner() view returns(address)
func (_LF5 *LF5CallerSession) GetOwner() (common.Address, error) {
	return _LF5.Contract.GetOwner(&_LF5.CallOpts)
}

// IsBlackListed is a free data retrieval call binding the contract method 0xe47d6060.
//
// Solidity: function isBlackListed(address ) view returns(bool)
func (_LF5 *LF5Caller) IsBlackListed(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _LF5.contract.Call(opts, &out, "isBlackListed", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsBlackListed is a free data retrieval call binding the contract method 0xe47d6060.
//
// Solidity: function isBlackListed(address ) view returns(bool)
func (_LF5 *LF5Session) IsBlackListed(arg0 common.Address) (bool, error) {
	return _LF5.Contract.IsBlackListed(&_LF5.CallOpts, arg0)
}

// IsBlackListed is a free data retrieval call binding the contract method 0xe47d6060.
//
// Solidity: function isBlackListed(address ) view returns(bool)
func (_LF5 *LF5CallerSession) IsBlackListed(arg0 common.Address) (bool, error) {
	return _LF5.Contract.IsBlackListed(&_LF5.CallOpts, arg0)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_LF5 *LF5Caller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _LF5.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_LF5 *LF5Session) Name() (string, error) {
	return _LF5.Contract.Name(&_LF5.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_LF5 *LF5CallerSession) Name() (string, error) {
	return _LF5.Contract.Name(&_LF5.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_LF5 *LF5Caller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LF5.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_LF5 *LF5Session) Owner() (common.Address, error) {
	return _LF5.Contract.Owner(&_LF5.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_LF5 *LF5CallerSession) Owner() (common.Address, error) {
	return _LF5.Contract.Owner(&_LF5.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_LF5 *LF5Caller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _LF5.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_LF5 *LF5Session) Paused() (bool, error) {
	return _LF5.Contract.Paused(&_LF5.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_LF5 *LF5CallerSession) Paused() (bool, error) {
	return _LF5.Contract.Paused(&_LF5.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_LF5 *LF5Caller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _LF5.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_LF5 *LF5Session) Symbol() (string, error) {
	return _LF5.Contract.Symbol(&_LF5.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_LF5 *LF5CallerSession) Symbol() (string, error) {
	return _LF5.Contract.Symbol(&_LF5.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_LF5 *LF5Caller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LF5.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_LF5 *LF5Session) TotalSupply() (*big.Int, error) {
	return _LF5.Contract.TotalSupply(&_LF5.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_LF5 *LF5CallerSession) TotalSupply() (*big.Int, error) {
	return _LF5.Contract.TotalSupply(&_LF5.CallOpts)
}

// AddAdmin is a paid mutator transaction binding the contract method 0x46f2ed2e.
//
// Solidity: function addAdmin(address a, uint256 maxIssuingTokenPerTime, uint256 maxTotalIssuingToken) returns()
func (_LF5 *LF5Transactor) AddAdmin(opts *bind.TransactOpts, a common.Address, maxIssuingTokenPerTime *big.Int, maxTotalIssuingToken *big.Int) (*types.Transaction, error) {
	return _LF5.contract.Transact(opts, "addAdmin", a, maxIssuingTokenPerTime, maxTotalIssuingToken)
}

// AddAdmin is a paid mutator transaction binding the contract method 0x46f2ed2e.
//
// Solidity: function addAdmin(address a, uint256 maxIssuingTokenPerTime, uint256 maxTotalIssuingToken) returns()
func (_LF5 *LF5Session) AddAdmin(a common.Address, maxIssuingTokenPerTime *big.Int, maxTotalIssuingToken *big.Int) (*types.Transaction, error) {
	return _LF5.Contract.AddAdmin(&_LF5.TransactOpts, a, maxIssuingTokenPerTime, maxTotalIssuingToken)
}

// AddAdmin is a paid mutator transaction binding the contract method 0x46f2ed2e.
//
// Solidity: function addAdmin(address a, uint256 maxIssuingTokenPerTime, uint256 maxTotalIssuingToken) returns()
func (_LF5 *LF5TransactorSession) AddAdmin(a common.Address, maxIssuingTokenPerTime *big.Int, maxTotalIssuingToken *big.Int) (*types.Transaction, error) {
	return _LF5.Contract.AddAdmin(&_LF5.TransactOpts, a, maxIssuingTokenPerTime, maxTotalIssuingToken)
}

// AddBlackList is a paid mutator transaction binding the contract method 0x0ecb93c0.
//
// Solidity: function addBlackList(address _evilUser) returns()
func (_LF5 *LF5Transactor) AddBlackList(opts *bind.TransactOpts, _evilUser common.Address) (*types.Transaction, error) {
	return _LF5.contract.Transact(opts, "addBlackList", _evilUser)
}

// AddBlackList is a paid mutator transaction binding the contract method 0x0ecb93c0.
//
// Solidity: function addBlackList(address _evilUser) returns()
func (_LF5 *LF5Session) AddBlackList(_evilUser common.Address) (*types.Transaction, error) {
	return _LF5.Contract.AddBlackList(&_LF5.TransactOpts, _evilUser)
}

// AddBlackList is a paid mutator transaction binding the contract method 0x0ecb93c0.
//
// Solidity: function addBlackList(address _evilUser) returns()
func (_LF5 *LF5TransactorSession) AddBlackList(_evilUser common.Address) (*types.Transaction, error) {
	return _LF5.Contract.AddBlackList(&_LF5.TransactOpts, _evilUser)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_LF5 *LF5Transactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _LF5.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_LF5 *LF5Session) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _LF5.Contract.Approve(&_LF5.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_LF5 *LF5TransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _LF5.Contract.Approve(&_LF5.TransactOpts, spender, amount)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_LF5 *LF5Transactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _LF5.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_LF5 *LF5Session) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _LF5.Contract.DecreaseAllowance(&_LF5.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_LF5 *LF5TransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _LF5.Contract.DecreaseAllowance(&_LF5.TransactOpts, spender, subtractedValue)
}

// DestroyBlackFunds is a paid mutator transaction binding the contract method 0xf3bdc228.
//
// Solidity: function destroyBlackFunds(address _blackListedUser) returns()
func (_LF5 *LF5Transactor) DestroyBlackFunds(opts *bind.TransactOpts, _blackListedUser common.Address) (*types.Transaction, error) {
	return _LF5.contract.Transact(opts, "destroyBlackFunds", _blackListedUser)
}

// DestroyBlackFunds is a paid mutator transaction binding the contract method 0xf3bdc228.
//
// Solidity: function destroyBlackFunds(address _blackListedUser) returns()
func (_LF5 *LF5Session) DestroyBlackFunds(_blackListedUser common.Address) (*types.Transaction, error) {
	return _LF5.Contract.DestroyBlackFunds(&_LF5.TransactOpts, _blackListedUser)
}

// DestroyBlackFunds is a paid mutator transaction binding the contract method 0xf3bdc228.
//
// Solidity: function destroyBlackFunds(address _blackListedUser) returns()
func (_LF5 *LF5TransactorSession) DestroyBlackFunds(_blackListedUser common.Address) (*types.Transaction, error) {
	return _LF5.Contract.DestroyBlackFunds(&_LF5.TransactOpts, _blackListedUser)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_LF5 *LF5Transactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _LF5.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_LF5 *LF5Session) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _LF5.Contract.IncreaseAllowance(&_LF5.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_LF5 *LF5TransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _LF5.Contract.IncreaseAllowance(&_LF5.TransactOpts, spender, addedValue)
}

// Issue is a paid mutator transaction binding the contract method 0xea3f068d.
//
// Solidity: function issue(address issuer, uint256 value, bytes32 offchain) returns()
func (_LF5 *LF5Transactor) Issue(opts *bind.TransactOpts, issuer common.Address, value *big.Int, offchain [32]byte) (*types.Transaction, error) {
	return _LF5.contract.Transact(opts, "issue", issuer, value, offchain)
}

// Issue is a paid mutator transaction binding the contract method 0xea3f068d.
//
// Solidity: function issue(address issuer, uint256 value, bytes32 offchain) returns()
func (_LF5 *LF5Session) Issue(issuer common.Address, value *big.Int, offchain [32]byte) (*types.Transaction, error) {
	return _LF5.Contract.Issue(&_LF5.TransactOpts, issuer, value, offchain)
}

// Issue is a paid mutator transaction binding the contract method 0xea3f068d.
//
// Solidity: function issue(address issuer, uint256 value, bytes32 offchain) returns()
func (_LF5 *LF5TransactorSession) Issue(issuer common.Address, value *big.Int, offchain [32]byte) (*types.Transaction, error) {
	return _LF5.Contract.Issue(&_LF5.TransactOpts, issuer, value, offchain)
}

// IssueByAdmin is a paid mutator transaction binding the contract method 0xd10601ec.
//
// Solidity: function issueByAdmin(address issuer, uint256 value, bytes32 offchain) returns()
func (_LF5 *LF5Transactor) IssueByAdmin(opts *bind.TransactOpts, issuer common.Address, value *big.Int, offchain [32]byte) (*types.Transaction, error) {
	return _LF5.contract.Transact(opts, "issueByAdmin", issuer, value, offchain)
}

// IssueByAdmin is a paid mutator transaction binding the contract method 0xd10601ec.
//
// Solidity: function issueByAdmin(address issuer, uint256 value, bytes32 offchain) returns()
func (_LF5 *LF5Session) IssueByAdmin(issuer common.Address, value *big.Int, offchain [32]byte) (*types.Transaction, error) {
	return _LF5.Contract.IssueByAdmin(&_LF5.TransactOpts, issuer, value, offchain)
}

// IssueByAdmin is a paid mutator transaction binding the contract method 0xd10601ec.
//
// Solidity: function issueByAdmin(address issuer, uint256 value, bytes32 offchain) returns()
func (_LF5 *LF5TransactorSession) IssueByAdmin(issuer common.Address, value *big.Int, offchain [32]byte) (*types.Transaction, error) {
	return _LF5.Contract.IssueByAdmin(&_LF5.TransactOpts, issuer, value, offchain)
}

// Mint is a paid mutator transaction binding the contract method 0xa0712d68.
//
// Solidity: function mint(uint256 amount) returns(bool)
func (_LF5 *LF5Transactor) Mint(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _LF5.contract.Transact(opts, "mint", amount)
}

// Mint is a paid mutator transaction binding the contract method 0xa0712d68.
//
// Solidity: function mint(uint256 amount) returns(bool)
func (_LF5 *LF5Session) Mint(amount *big.Int) (*types.Transaction, error) {
	return _LF5.Contract.Mint(&_LF5.TransactOpts, amount)
}

// Mint is a paid mutator transaction binding the contract method 0xa0712d68.
//
// Solidity: function mint(uint256 amount) returns(bool)
func (_LF5 *LF5TransactorSession) Mint(amount *big.Int) (*types.Transaction, error) {
	return _LF5.Contract.Mint(&_LF5.TransactOpts, amount)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_LF5 *LF5Transactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LF5.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_LF5 *LF5Session) Pause() (*types.Transaction, error) {
	return _LF5.Contract.Pause(&_LF5.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_LF5 *LF5TransactorSession) Pause() (*types.Transaction, error) {
	return _LF5.Contract.Pause(&_LF5.TransactOpts)
}

// Redeem is a paid mutator transaction binding the contract method 0x992c3e4b.
//
// Solidity: function redeem(address redeemer, uint256 value, bytes32 offchain) returns()
func (_LF5 *LF5Transactor) Redeem(opts *bind.TransactOpts, redeemer common.Address, value *big.Int, offchain [32]byte) (*types.Transaction, error) {
	return _LF5.contract.Transact(opts, "redeem", redeemer, value, offchain)
}

// Redeem is a paid mutator transaction binding the contract method 0x992c3e4b.
//
// Solidity: function redeem(address redeemer, uint256 value, bytes32 offchain) returns()
func (_LF5 *LF5Session) Redeem(redeemer common.Address, value *big.Int, offchain [32]byte) (*types.Transaction, error) {
	return _LF5.Contract.Redeem(&_LF5.TransactOpts, redeemer, value, offchain)
}

// Redeem is a paid mutator transaction binding the contract method 0x992c3e4b.
//
// Solidity: function redeem(address redeemer, uint256 value, bytes32 offchain) returns()
func (_LF5 *LF5TransactorSession) Redeem(redeemer common.Address, value *big.Int, offchain [32]byte) (*types.Transaction, error) {
	return _LF5.Contract.Redeem(&_LF5.TransactOpts, redeemer, value, offchain)
}

// RemoveAdmin is a paid mutator transaction binding the contract method 0x1785f53c.
//
// Solidity: function removeAdmin(address a) returns()
func (_LF5 *LF5Transactor) RemoveAdmin(opts *bind.TransactOpts, a common.Address) (*types.Transaction, error) {
	return _LF5.contract.Transact(opts, "removeAdmin", a)
}

// RemoveAdmin is a paid mutator transaction binding the contract method 0x1785f53c.
//
// Solidity: function removeAdmin(address a) returns()
func (_LF5 *LF5Session) RemoveAdmin(a common.Address) (*types.Transaction, error) {
	return _LF5.Contract.RemoveAdmin(&_LF5.TransactOpts, a)
}

// RemoveAdmin is a paid mutator transaction binding the contract method 0x1785f53c.
//
// Solidity: function removeAdmin(address a) returns()
func (_LF5 *LF5TransactorSession) RemoveAdmin(a common.Address) (*types.Transaction, error) {
	return _LF5.Contract.RemoveAdmin(&_LF5.TransactOpts, a)
}

// RemoveBlackList is a paid mutator transaction binding the contract method 0xe4997dc5.
//
// Solidity: function removeBlackList(address _clearedUser) returns()
func (_LF5 *LF5Transactor) RemoveBlackList(opts *bind.TransactOpts, _clearedUser common.Address) (*types.Transaction, error) {
	return _LF5.contract.Transact(opts, "removeBlackList", _clearedUser)
}

// RemoveBlackList is a paid mutator transaction binding the contract method 0xe4997dc5.
//
// Solidity: function removeBlackList(address _clearedUser) returns()
func (_LF5 *LF5Session) RemoveBlackList(_clearedUser common.Address) (*types.Transaction, error) {
	return _LF5.Contract.RemoveBlackList(&_LF5.TransactOpts, _clearedUser)
}

// RemoveBlackList is a paid mutator transaction binding the contract method 0xe4997dc5.
//
// Solidity: function removeBlackList(address _clearedUser) returns()
func (_LF5 *LF5TransactorSession) RemoveBlackList(_clearedUser common.Address) (*types.Transaction, error) {
	return _LF5.Contract.RemoveBlackList(&_LF5.TransactOpts, _clearedUser)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_LF5 *LF5Transactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LF5.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_LF5 *LF5Session) RenounceOwnership() (*types.Transaction, error) {
	return _LF5.Contract.RenounceOwnership(&_LF5.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_LF5 *LF5TransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _LF5.Contract.RenounceOwnership(&_LF5.TransactOpts)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns(bool)
func (_LF5 *LF5Transactor) Transfer(opts *bind.TransactOpts, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _LF5.contract.Transact(opts, "transfer", _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns(bool)
func (_LF5 *LF5Session) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _LF5.Contract.Transfer(&_LF5.TransactOpts, _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns(bool)
func (_LF5 *LF5TransactorSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _LF5.Contract.Transfer(&_LF5.TransactOpts, _to, _value)
}

// TransferByAdmin is a paid mutator transaction binding the contract method 0xf9ba884f.
//
// Solidity: function transferByAdmin(address from, address to, uint256 value, bytes32 offchain) returns()
func (_LF5 *LF5Transactor) TransferByAdmin(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int, offchain [32]byte) (*types.Transaction, error) {
	return _LF5.contract.Transact(opts, "transferByAdmin", from, to, value, offchain)
}

// TransferByAdmin is a paid mutator transaction binding the contract method 0xf9ba884f.
//
// Solidity: function transferByAdmin(address from, address to, uint256 value, bytes32 offchain) returns()
func (_LF5 *LF5Session) TransferByAdmin(from common.Address, to common.Address, value *big.Int, offchain [32]byte) (*types.Transaction, error) {
	return _LF5.Contract.TransferByAdmin(&_LF5.TransactOpts, from, to, value, offchain)
}

// TransferByAdmin is a paid mutator transaction binding the contract method 0xf9ba884f.
//
// Solidity: function transferByAdmin(address from, address to, uint256 value, bytes32 offchain) returns()
func (_LF5 *LF5TransactorSession) TransferByAdmin(from common.Address, to common.Address, value *big.Int, offchain [32]byte) (*types.Transaction, error) {
	return _LF5.Contract.TransferByAdmin(&_LF5.TransactOpts, from, to, value, offchain)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _value) returns(bool)
func (_LF5 *LF5Transactor) TransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _LF5.contract.Transact(opts, "transferFrom", _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _value) returns(bool)
func (_LF5 *LF5Session) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _LF5.Contract.TransferFrom(&_LF5.TransactOpts, _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _value) returns(bool)
func (_LF5 *LF5TransactorSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _LF5.Contract.TransferFrom(&_LF5.TransactOpts, _from, _to, _value)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_LF5 *LF5Transactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _LF5.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_LF5 *LF5Session) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _LF5.Contract.TransferOwnership(&_LF5.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_LF5 *LF5TransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _LF5.Contract.TransferOwnership(&_LF5.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_LF5 *LF5Transactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LF5.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_LF5 *LF5Session) Unpause() (*types.Transaction, error) {
	return _LF5.Contract.Unpause(&_LF5.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_LF5 *LF5TransactorSession) Unpause() (*types.Transaction, error) {
	return _LF5.Contract.Unpause(&_LF5.TransactOpts)
}

// LF5AddedBlackListIterator is returned from FilterAddedBlackList and is used to iterate over the raw logs and unpacked data for AddedBlackList events raised by the LF5 contract.
type LF5AddedBlackListIterator struct {
	Event *LF5AddedBlackList // Event containing the contract specifics and raw log

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
func (it *LF5AddedBlackListIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LF5AddedBlackList)
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
		it.Event = new(LF5AddedBlackList)
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
func (it *LF5AddedBlackListIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LF5AddedBlackListIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LF5AddedBlackList represents a AddedBlackList event raised by the LF5 contract.
type LF5AddedBlackList struct {
	User common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterAddedBlackList is a free log retrieval operation binding the contract event 0x42e160154868087d6bfdc0ca23d96a1c1cfa32f1b72ba9ba27b69b98a0d819dc.
//
// Solidity: event AddedBlackList(address _user)
func (_LF5 *LF5Filterer) FilterAddedBlackList(opts *bind.FilterOpts) (*LF5AddedBlackListIterator, error) {

	logs, sub, err := _LF5.contract.FilterLogs(opts, "AddedBlackList")
	if err != nil {
		return nil, err
	}
	return &LF5AddedBlackListIterator{contract: _LF5.contract, event: "AddedBlackList", logs: logs, sub: sub}, nil
}

// WatchAddedBlackList is a free log subscription operation binding the contract event 0x42e160154868087d6bfdc0ca23d96a1c1cfa32f1b72ba9ba27b69b98a0d819dc.
//
// Solidity: event AddedBlackList(address _user)
func (_LF5 *LF5Filterer) WatchAddedBlackList(opts *bind.WatchOpts, sink chan<- *LF5AddedBlackList) (event.Subscription, error) {

	logs, sub, err := _LF5.contract.WatchLogs(opts, "AddedBlackList")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LF5AddedBlackList)
				if err := _LF5.contract.UnpackLog(event, "AddedBlackList", log); err != nil {
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

// ParseAddedBlackList is a log parse operation binding the contract event 0x42e160154868087d6bfdc0ca23d96a1c1cfa32f1b72ba9ba27b69b98a0d819dc.
//
// Solidity: event AddedBlackList(address _user)
func (_LF5 *LF5Filterer) ParseAddedBlackList(log types.Log) (*LF5AddedBlackList, error) {
	event := new(LF5AddedBlackList)
	if err := _LF5.contract.UnpackLog(event, "AddedBlackList", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LF5ApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the LF5 contract.
type LF5ApprovalIterator struct {
	Event *LF5Approval // Event containing the contract specifics and raw log

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
func (it *LF5ApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LF5Approval)
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
		it.Event = new(LF5Approval)
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
func (it *LF5ApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LF5ApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LF5Approval represents a Approval event raised by the LF5 contract.
type LF5Approval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_LF5 *LF5Filterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*LF5ApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _LF5.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &LF5ApprovalIterator{contract: _LF5.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_LF5 *LF5Filterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *LF5Approval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _LF5.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LF5Approval)
				if err := _LF5.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_LF5 *LF5Filterer) ParseApproval(log types.Log) (*LF5Approval, error) {
	event := new(LF5Approval)
	if err := _LF5.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LF5DestroyedBlackFundsIterator is returned from FilterDestroyedBlackFunds and is used to iterate over the raw logs and unpacked data for DestroyedBlackFunds events raised by the LF5 contract.
type LF5DestroyedBlackFundsIterator struct {
	Event *LF5DestroyedBlackFunds // Event containing the contract specifics and raw log

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
func (it *LF5DestroyedBlackFundsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LF5DestroyedBlackFunds)
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
		it.Event = new(LF5DestroyedBlackFunds)
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
func (it *LF5DestroyedBlackFundsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LF5DestroyedBlackFundsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LF5DestroyedBlackFunds represents a DestroyedBlackFunds event raised by the LF5 contract.
type LF5DestroyedBlackFunds struct {
	BlackListedUser common.Address
	Balance         *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterDestroyedBlackFunds is a free log retrieval operation binding the contract event 0x61e6e66b0d6339b2980aecc6ccc0039736791f0ccde9ed512e789a7fbdd698c6.
//
// Solidity: event DestroyedBlackFunds(address _blackListedUser, uint256 _balance)
func (_LF5 *LF5Filterer) FilterDestroyedBlackFunds(opts *bind.FilterOpts) (*LF5DestroyedBlackFundsIterator, error) {

	logs, sub, err := _LF5.contract.FilterLogs(opts, "DestroyedBlackFunds")
	if err != nil {
		return nil, err
	}
	return &LF5DestroyedBlackFundsIterator{contract: _LF5.contract, event: "DestroyedBlackFunds", logs: logs, sub: sub}, nil
}

// WatchDestroyedBlackFunds is a free log subscription operation binding the contract event 0x61e6e66b0d6339b2980aecc6ccc0039736791f0ccde9ed512e789a7fbdd698c6.
//
// Solidity: event DestroyedBlackFunds(address _blackListedUser, uint256 _balance)
func (_LF5 *LF5Filterer) WatchDestroyedBlackFunds(opts *bind.WatchOpts, sink chan<- *LF5DestroyedBlackFunds) (event.Subscription, error) {

	logs, sub, err := _LF5.contract.WatchLogs(opts, "DestroyedBlackFunds")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LF5DestroyedBlackFunds)
				if err := _LF5.contract.UnpackLog(event, "DestroyedBlackFunds", log); err != nil {
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

// ParseDestroyedBlackFunds is a log parse operation binding the contract event 0x61e6e66b0d6339b2980aecc6ccc0039736791f0ccde9ed512e789a7fbdd698c6.
//
// Solidity: event DestroyedBlackFunds(address _blackListedUser, uint256 _balance)
func (_LF5 *LF5Filterer) ParseDestroyedBlackFunds(log types.Log) (*LF5DestroyedBlackFunds, error) {
	event := new(LF5DestroyedBlackFunds)
	if err := _LF5.contract.UnpackLog(event, "DestroyedBlackFunds", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LF5OwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the LF5 contract.
type LF5OwnershipTransferredIterator struct {
	Event *LF5OwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *LF5OwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LF5OwnershipTransferred)
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
		it.Event = new(LF5OwnershipTransferred)
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
func (it *LF5OwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LF5OwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LF5OwnershipTransferred represents a OwnershipTransferred event raised by the LF5 contract.
type LF5OwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_LF5 *LF5Filterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*LF5OwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _LF5.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &LF5OwnershipTransferredIterator{contract: _LF5.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_LF5 *LF5Filterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *LF5OwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _LF5.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LF5OwnershipTransferred)
				if err := _LF5.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_LF5 *LF5Filterer) ParseOwnershipTransferred(log types.Log) (*LF5OwnershipTransferred, error) {
	event := new(LF5OwnershipTransferred)
	if err := _LF5.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LF5PausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the LF5 contract.
type LF5PausedIterator struct {
	Event *LF5Paused // Event containing the contract specifics and raw log

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
func (it *LF5PausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LF5Paused)
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
		it.Event = new(LF5Paused)
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
func (it *LF5PausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LF5PausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LF5Paused represents a Paused event raised by the LF5 contract.
type LF5Paused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_LF5 *LF5Filterer) FilterPaused(opts *bind.FilterOpts) (*LF5PausedIterator, error) {

	logs, sub, err := _LF5.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &LF5PausedIterator{contract: _LF5.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_LF5 *LF5Filterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *LF5Paused) (event.Subscription, error) {

	logs, sub, err := _LF5.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LF5Paused)
				if err := _LF5.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_LF5 *LF5Filterer) ParsePaused(log types.Log) (*LF5Paused, error) {
	event := new(LF5Paused)
	if err := _LF5.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LF5RemovedBlackListIterator is returned from FilterRemovedBlackList and is used to iterate over the raw logs and unpacked data for RemovedBlackList events raised by the LF5 contract.
type LF5RemovedBlackListIterator struct {
	Event *LF5RemovedBlackList // Event containing the contract specifics and raw log

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
func (it *LF5RemovedBlackListIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LF5RemovedBlackList)
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
		it.Event = new(LF5RemovedBlackList)
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
func (it *LF5RemovedBlackListIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LF5RemovedBlackListIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LF5RemovedBlackList represents a RemovedBlackList event raised by the LF5 contract.
type LF5RemovedBlackList struct {
	User common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterRemovedBlackList is a free log retrieval operation binding the contract event 0xd7e9ec6e6ecd65492dce6bf513cd6867560d49544421d0783ddf06e76c24470c.
//
// Solidity: event RemovedBlackList(address _user)
func (_LF5 *LF5Filterer) FilterRemovedBlackList(opts *bind.FilterOpts) (*LF5RemovedBlackListIterator, error) {

	logs, sub, err := _LF5.contract.FilterLogs(opts, "RemovedBlackList")
	if err != nil {
		return nil, err
	}
	return &LF5RemovedBlackListIterator{contract: _LF5.contract, event: "RemovedBlackList", logs: logs, sub: sub}, nil
}

// WatchRemovedBlackList is a free log subscription operation binding the contract event 0xd7e9ec6e6ecd65492dce6bf513cd6867560d49544421d0783ddf06e76c24470c.
//
// Solidity: event RemovedBlackList(address _user)
func (_LF5 *LF5Filterer) WatchRemovedBlackList(opts *bind.WatchOpts, sink chan<- *LF5RemovedBlackList) (event.Subscription, error) {

	logs, sub, err := _LF5.contract.WatchLogs(opts, "RemovedBlackList")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LF5RemovedBlackList)
				if err := _LF5.contract.UnpackLog(event, "RemovedBlackList", log); err != nil {
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

// ParseRemovedBlackList is a log parse operation binding the contract event 0xd7e9ec6e6ecd65492dce6bf513cd6867560d49544421d0783ddf06e76c24470c.
//
// Solidity: event RemovedBlackList(address _user)
func (_LF5 *LF5Filterer) ParseRemovedBlackList(log types.Log) (*LF5RemovedBlackList, error) {
	event := new(LF5RemovedBlackList)
	if err := _LF5.contract.UnpackLog(event, "RemovedBlackList", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LF5TransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the LF5 contract.
type LF5TransferIterator struct {
	Event *LF5Transfer // Event containing the contract specifics and raw log

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
func (it *LF5TransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LF5Transfer)
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
		it.Event = new(LF5Transfer)
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
func (it *LF5TransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LF5TransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LF5Transfer represents a Transfer event raised by the LF5 contract.
type LF5Transfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_LF5 *LF5Filterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*LF5TransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _LF5.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &LF5TransferIterator{contract: _LF5.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_LF5 *LF5Filterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *LF5Transfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _LF5.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LF5Transfer)
				if err := _LF5.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_LF5 *LF5Filterer) ParseTransfer(log types.Log) (*LF5Transfer, error) {
	event := new(LF5Transfer)
	if err := _LF5.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LF5UnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the LF5 contract.
type LF5UnpausedIterator struct {
	Event *LF5Unpaused // Event containing the contract specifics and raw log

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
func (it *LF5UnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LF5Unpaused)
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
		it.Event = new(LF5Unpaused)
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
func (it *LF5UnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LF5UnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LF5Unpaused represents a Unpaused event raised by the LF5 contract.
type LF5Unpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_LF5 *LF5Filterer) FilterUnpaused(opts *bind.FilterOpts) (*LF5UnpausedIterator, error) {

	logs, sub, err := _LF5.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &LF5UnpausedIterator{contract: _LF5.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_LF5 *LF5Filterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *LF5Unpaused) (event.Subscription, error) {

	logs, sub, err := _LF5.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LF5Unpaused)
				if err := _LF5.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
func (_LF5 *LF5Filterer) ParseUnpaused(log types.Log) (*LF5Unpaused, error) {
	event := new(LF5Unpaused)
	if err := _LF5.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LF5IssueIterator is returned from FilterIssue and is used to iterate over the raw logs and unpacked data for Issue events raised by the LF5 contract.
type LF5IssueIterator struct {
	Event *LF5Issue // Event containing the contract specifics and raw log

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
func (it *LF5IssueIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LF5Issue)
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
		it.Event = new(LF5Issue)
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
func (it *LF5IssueIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LF5IssueIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LF5Issue represents a Issue event raised by the LF5 contract.
type LF5Issue struct {
	Offchain [32]byte
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterIssue is a free log retrieval operation binding the contract event 0x182a88d3fbfcb38ffaddf26e6d63ecd93b4348141de1eb744c43c4fbb1ccfb22.
//
// Solidity: event __issue(bytes32 offchain)
func (_LF5 *LF5Filterer) FilterIssue(opts *bind.FilterOpts) (*LF5IssueIterator, error) {

	logs, sub, err := _LF5.contract.FilterLogs(opts, "__issue")
	if err != nil {
		return nil, err
	}
	return &LF5IssueIterator{contract: _LF5.contract, event: "__issue", logs: logs, sub: sub}, nil
}

// WatchIssue is a free log subscription operation binding the contract event 0x182a88d3fbfcb38ffaddf26e6d63ecd93b4348141de1eb744c43c4fbb1ccfb22.
//
// Solidity: event __issue(bytes32 offchain)
func (_LF5 *LF5Filterer) WatchIssue(opts *bind.WatchOpts, sink chan<- *LF5Issue) (event.Subscription, error) {

	logs, sub, err := _LF5.contract.WatchLogs(opts, "__issue")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LF5Issue)
				if err := _LF5.contract.UnpackLog(event, "__issue", log); err != nil {
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

// ParseIssue is a log parse operation binding the contract event 0x182a88d3fbfcb38ffaddf26e6d63ecd93b4348141de1eb744c43c4fbb1ccfb22.
//
// Solidity: event __issue(bytes32 offchain)
func (_LF5 *LF5Filterer) ParseIssue(log types.Log) (*LF5Issue, error) {
	event := new(LF5Issue)
	if err := _LF5.contract.UnpackLog(event, "__issue", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LF5RedeemIterator is returned from FilterRedeem and is used to iterate over the raw logs and unpacked data for Redeem events raised by the LF5 contract.
type LF5RedeemIterator struct {
	Event *LF5Redeem // Event containing the contract specifics and raw log

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
func (it *LF5RedeemIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LF5Redeem)
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
		it.Event = new(LF5Redeem)
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
func (it *LF5RedeemIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LF5RedeemIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LF5Redeem represents a Redeem event raised by the LF5 contract.
type LF5Redeem struct {
	Offchain [32]byte
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterRedeem is a free log retrieval operation binding the contract event 0xb0de879351469d2741406aafc9ba1f44eb957cf44ee3391e59a7a9097050c927.
//
// Solidity: event __redeem(bytes32 offchain)
func (_LF5 *LF5Filterer) FilterRedeem(opts *bind.FilterOpts) (*LF5RedeemIterator, error) {

	logs, sub, err := _LF5.contract.FilterLogs(opts, "__redeem")
	if err != nil {
		return nil, err
	}
	return &LF5RedeemIterator{contract: _LF5.contract, event: "__redeem", logs: logs, sub: sub}, nil
}

// WatchRedeem is a free log subscription operation binding the contract event 0xb0de879351469d2741406aafc9ba1f44eb957cf44ee3391e59a7a9097050c927.
//
// Solidity: event __redeem(bytes32 offchain)
func (_LF5 *LF5Filterer) WatchRedeem(opts *bind.WatchOpts, sink chan<- *LF5Redeem) (event.Subscription, error) {

	logs, sub, err := _LF5.contract.WatchLogs(opts, "__redeem")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LF5Redeem)
				if err := _LF5.contract.UnpackLog(event, "__redeem", log); err != nil {
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

// ParseRedeem is a log parse operation binding the contract event 0xb0de879351469d2741406aafc9ba1f44eb957cf44ee3391e59a7a9097050c927.
//
// Solidity: event __redeem(bytes32 offchain)
func (_LF5 *LF5Filterer) ParseRedeem(log types.Log) (*LF5Redeem, error) {
	event := new(LF5Redeem)
	if err := _LF5.contract.UnpackLog(event, "__redeem", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LF5TransferByAdminIterator is returned from FilterTransferByAdmin and is used to iterate over the raw logs and unpacked data for TransferByAdmin events raised by the LF5 contract.
type LF5TransferByAdminIterator struct {
	Event *LF5TransferByAdmin // Event containing the contract specifics and raw log

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
func (it *LF5TransferByAdminIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LF5TransferByAdmin)
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
		it.Event = new(LF5TransferByAdmin)
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
func (it *LF5TransferByAdminIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LF5TransferByAdminIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LF5TransferByAdmin represents a TransferByAdmin event raised by the LF5 contract.
type LF5TransferByAdmin struct {
	Offchain [32]byte
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterTransferByAdmin is a free log retrieval operation binding the contract event 0x3b36ee6b35325f38e95938557be92853c842b7a9a19fd7ac4931a6d24db52682.
//
// Solidity: event __transferByAdmin(bytes32 offchain)
func (_LF5 *LF5Filterer) FilterTransferByAdmin(opts *bind.FilterOpts) (*LF5TransferByAdminIterator, error) {

	logs, sub, err := _LF5.contract.FilterLogs(opts, "__transferByAdmin")
	if err != nil {
		return nil, err
	}
	return &LF5TransferByAdminIterator{contract: _LF5.contract, event: "__transferByAdmin", logs: logs, sub: sub}, nil
}

// WatchTransferByAdmin is a free log subscription operation binding the contract event 0x3b36ee6b35325f38e95938557be92853c842b7a9a19fd7ac4931a6d24db52682.
//
// Solidity: event __transferByAdmin(bytes32 offchain)
func (_LF5 *LF5Filterer) WatchTransferByAdmin(opts *bind.WatchOpts, sink chan<- *LF5TransferByAdmin) (event.Subscription, error) {

	logs, sub, err := _LF5.contract.WatchLogs(opts, "__transferByAdmin")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LF5TransferByAdmin)
				if err := _LF5.contract.UnpackLog(event, "__transferByAdmin", log); err != nil {
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

// ParseTransferByAdmin is a log parse operation binding the contract event 0x3b36ee6b35325f38e95938557be92853c842b7a9a19fd7ac4931a6d24db52682.
//
// Solidity: event __transferByAdmin(bytes32 offchain)
func (_LF5 *LF5Filterer) ParseTransferByAdmin(log types.Log) (*LF5TransferByAdmin, error) {
	event := new(LF5TransferByAdmin)
	if err := _LF5.contract.UnpackLog(event, "__transferByAdmin", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
