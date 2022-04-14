// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package BUSD

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

// BUSDMetaData contains all meta data concerning the BUSD contract.
var BUSDMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"AddedBlackList\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_blackListedUser\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_balance\",\"type\":\"uint256\"}],\"name\":\"DestroyedBlackFunds\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"RemovedBlackList\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"offchain\",\"type\":\"bytes32\"}],\"name\":\"__issue\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"offchain\",\"type\":\"bytes32\"}],\"name\":\"__redeem\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"offchain\",\"type\":\"bytes32\"}],\"name\":\"__transferByAdmin\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"a\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"maxIssuingTokenPerTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxTotalIssuingToken\",\"type\":\"uint256\"}],\"name\":\"addAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_evilUser\",\"type\":\"address\"}],\"name\":\"addBlackList\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"admin\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"status\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"maxIssuingTokenPerTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxTotalIssuingToken\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"remainingIssuingToken\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_maker\",\"type\":\"address\"}],\"name\":\"getBlackListStatus\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"isBlackListed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"a\",\"type\":\"address\"}],\"name\":\"removeAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_clearedUser\",\"type\":\"address\"}],\"name\":\"removeBlackList\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"offchain\",\"type\":\"bytes32\"}],\"name\":\"transferByAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"issuer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"offchain\",\"type\":\"bytes32\"}],\"name\":\"issue\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"issuer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"offchain\",\"type\":\"bytes32\"}],\"name\":\"issueByAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_blackListedUser\",\"type\":\"address\"}],\"name\":\"destroyBlackFunds\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"redeemer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"offchain\",\"type\":\"bytes32\"}],\"name\":\"redeem\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// BUSDABI is the input ABI used to generate the binding from.
// Deprecated: Use BUSDMetaData.ABI instead.
var BUSDABI = BUSDMetaData.ABI

// BUSD is an auto generated Go binding around an Ethereum contract.
type BUSD struct {
	BUSDCaller     // Read-only binding to the contract
	BUSDTransactor // Write-only binding to the contract
	BUSDFilterer   // Log filterer for contract events
}

// BUSDCaller is an auto generated read-only Go binding around an Ethereum contract.
type BUSDCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BUSDTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BUSDTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BUSDFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BUSDFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BUSDSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BUSDSession struct {
	Contract     *BUSD             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BUSDCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BUSDCallerSession struct {
	Contract *BUSDCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// BUSDTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BUSDTransactorSession struct {
	Contract     *BUSDTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BUSDRaw is an auto generated low-level Go binding around an Ethereum contract.
type BUSDRaw struct {
	Contract *BUSD // Generic contract binding to access the raw methods on
}

// BUSDCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BUSDCallerRaw struct {
	Contract *BUSDCaller // Generic read-only contract binding to access the raw methods on
}

// BUSDTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BUSDTransactorRaw struct {
	Contract *BUSDTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBUSD creates a new instance of BUSD, bound to a specific deployed contract.
func NewBUSD(address common.Address, backend bind.ContractBackend) (*BUSD, error) {
	contract, err := bindBUSD(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BUSD{BUSDCaller: BUSDCaller{contract: contract}, BUSDTransactor: BUSDTransactor{contract: contract}, BUSDFilterer: BUSDFilterer{contract: contract}}, nil
}

// NewBUSDCaller creates a new read-only instance of BUSD, bound to a specific deployed contract.
func NewBUSDCaller(address common.Address, caller bind.ContractCaller) (*BUSDCaller, error) {
	contract, err := bindBUSD(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BUSDCaller{contract: contract}, nil
}

// NewBUSDTransactor creates a new write-only instance of BUSD, bound to a specific deployed contract.
func NewBUSDTransactor(address common.Address, transactor bind.ContractTransactor) (*BUSDTransactor, error) {
	contract, err := bindBUSD(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BUSDTransactor{contract: contract}, nil
}

// NewBUSDFilterer creates a new log filterer instance of BUSD, bound to a specific deployed contract.
func NewBUSDFilterer(address common.Address, filterer bind.ContractFilterer) (*BUSDFilterer, error) {
	contract, err := bindBUSD(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BUSDFilterer{contract: contract}, nil
}

// bindBUSD binds a generic wrapper to an already deployed contract.
func bindBUSD(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BUSDABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BUSD *BUSDRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BUSD.Contract.BUSDCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BUSD *BUSDRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BUSD.Contract.BUSDTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BUSD *BUSDRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BUSD.Contract.BUSDTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BUSD *BUSDCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BUSD.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BUSD *BUSDTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BUSD.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BUSD *BUSDTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BUSD.Contract.contract.Transact(opts, method, params...)
}

// Admin is a free data retrieval call binding the contract method 0x63a846f8.
//
// Solidity: function admin(address ) view returns(bool status, uint256 maxIssuingTokenPerTime, uint256 maxTotalIssuingToken, uint256 remainingIssuingToken)
func (_BUSD *BUSDCaller) Admin(opts *bind.CallOpts, arg0 common.Address) (struct {
	Status                 bool
	MaxIssuingTokenPerTime *big.Int
	MaxTotalIssuingToken   *big.Int
	RemainingIssuingToken  *big.Int
}, error) {
	var out []interface{}
	err := _BUSD.contract.Call(opts, &out, "admin", arg0)

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
func (_BUSD *BUSDSession) Admin(arg0 common.Address) (struct {
	Status                 bool
	MaxIssuingTokenPerTime *big.Int
	MaxTotalIssuingToken   *big.Int
	RemainingIssuingToken  *big.Int
}, error) {
	return _BUSD.Contract.Admin(&_BUSD.CallOpts, arg0)
}

// Admin is a free data retrieval call binding the contract method 0x63a846f8.
//
// Solidity: function admin(address ) view returns(bool status, uint256 maxIssuingTokenPerTime, uint256 maxTotalIssuingToken, uint256 remainingIssuingToken)
func (_BUSD *BUSDCallerSession) Admin(arg0 common.Address) (struct {
	Status                 bool
	MaxIssuingTokenPerTime *big.Int
	MaxTotalIssuingToken   *big.Int
	RemainingIssuingToken  *big.Int
}, error) {
	return _BUSD.Contract.Admin(&_BUSD.CallOpts, arg0)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_BUSD *BUSDCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _BUSD.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_BUSD *BUSDSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _BUSD.Contract.Allowance(&_BUSD.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_BUSD *BUSDCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _BUSD.Contract.Allowance(&_BUSD.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_BUSD *BUSDCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _BUSD.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_BUSD *BUSDSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _BUSD.Contract.BalanceOf(&_BUSD.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_BUSD *BUSDCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _BUSD.Contract.BalanceOf(&_BUSD.CallOpts, account)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_BUSD *BUSDCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _BUSD.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_BUSD *BUSDSession) Decimals() (uint8, error) {
	return _BUSD.Contract.Decimals(&_BUSD.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_BUSD *BUSDCallerSession) Decimals() (uint8, error) {
	return _BUSD.Contract.Decimals(&_BUSD.CallOpts)
}

// GetBlackListStatus is a free data retrieval call binding the contract method 0x59bf1abe.
//
// Solidity: function getBlackListStatus(address _maker) view returns(bool)
func (_BUSD *BUSDCaller) GetBlackListStatus(opts *bind.CallOpts, _maker common.Address) (bool, error) {
	var out []interface{}
	err := _BUSD.contract.Call(opts, &out, "getBlackListStatus", _maker)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// GetBlackListStatus is a free data retrieval call binding the contract method 0x59bf1abe.
//
// Solidity: function getBlackListStatus(address _maker) view returns(bool)
func (_BUSD *BUSDSession) GetBlackListStatus(_maker common.Address) (bool, error) {
	return _BUSD.Contract.GetBlackListStatus(&_BUSD.CallOpts, _maker)
}

// GetBlackListStatus is a free data retrieval call binding the contract method 0x59bf1abe.
//
// Solidity: function getBlackListStatus(address _maker) view returns(bool)
func (_BUSD *BUSDCallerSession) GetBlackListStatus(_maker common.Address) (bool, error) {
	return _BUSD.Contract.GetBlackListStatus(&_BUSD.CallOpts, _maker)
}

// GetOwner is a free data retrieval call binding the contract method 0x893d20e8.
//
// Solidity: function getOwner() view returns(address)
func (_BUSD *BUSDCaller) GetOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BUSD.contract.Call(opts, &out, "getOwner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetOwner is a free data retrieval call binding the contract method 0x893d20e8.
//
// Solidity: function getOwner() view returns(address)
func (_BUSD *BUSDSession) GetOwner() (common.Address, error) {
	return _BUSD.Contract.GetOwner(&_BUSD.CallOpts)
}

// GetOwner is a free data retrieval call binding the contract method 0x893d20e8.
//
// Solidity: function getOwner() view returns(address)
func (_BUSD *BUSDCallerSession) GetOwner() (common.Address, error) {
	return _BUSD.Contract.GetOwner(&_BUSD.CallOpts)
}

// IsBlackListed is a free data retrieval call binding the contract method 0xe47d6060.
//
// Solidity: function isBlackListed(address ) view returns(bool)
func (_BUSD *BUSDCaller) IsBlackListed(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _BUSD.contract.Call(opts, &out, "isBlackListed", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsBlackListed is a free data retrieval call binding the contract method 0xe47d6060.
//
// Solidity: function isBlackListed(address ) view returns(bool)
func (_BUSD *BUSDSession) IsBlackListed(arg0 common.Address) (bool, error) {
	return _BUSD.Contract.IsBlackListed(&_BUSD.CallOpts, arg0)
}

// IsBlackListed is a free data retrieval call binding the contract method 0xe47d6060.
//
// Solidity: function isBlackListed(address ) view returns(bool)
func (_BUSD *BUSDCallerSession) IsBlackListed(arg0 common.Address) (bool, error) {
	return _BUSD.Contract.IsBlackListed(&_BUSD.CallOpts, arg0)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_BUSD *BUSDCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _BUSD.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_BUSD *BUSDSession) Name() (string, error) {
	return _BUSD.Contract.Name(&_BUSD.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_BUSD *BUSDCallerSession) Name() (string, error) {
	return _BUSD.Contract.Name(&_BUSD.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BUSD *BUSDCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BUSD.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BUSD *BUSDSession) Owner() (common.Address, error) {
	return _BUSD.Contract.Owner(&_BUSD.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BUSD *BUSDCallerSession) Owner() (common.Address, error) {
	return _BUSD.Contract.Owner(&_BUSD.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_BUSD *BUSDCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _BUSD.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_BUSD *BUSDSession) Paused() (bool, error) {
	return _BUSD.Contract.Paused(&_BUSD.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_BUSD *BUSDCallerSession) Paused() (bool, error) {
	return _BUSD.Contract.Paused(&_BUSD.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_BUSD *BUSDCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _BUSD.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_BUSD *BUSDSession) Symbol() (string, error) {
	return _BUSD.Contract.Symbol(&_BUSD.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_BUSD *BUSDCallerSession) Symbol() (string, error) {
	return _BUSD.Contract.Symbol(&_BUSD.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_BUSD *BUSDCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BUSD.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_BUSD *BUSDSession) TotalSupply() (*big.Int, error) {
	return _BUSD.Contract.TotalSupply(&_BUSD.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_BUSD *BUSDCallerSession) TotalSupply() (*big.Int, error) {
	return _BUSD.Contract.TotalSupply(&_BUSD.CallOpts)
}

// AddAdmin is a paid mutator transaction binding the contract method 0x46f2ed2e.
//
// Solidity: function addAdmin(address a, uint256 maxIssuingTokenPerTime, uint256 maxTotalIssuingToken) returns()
func (_BUSD *BUSDTransactor) AddAdmin(opts *bind.TransactOpts, a common.Address, maxIssuingTokenPerTime *big.Int, maxTotalIssuingToken *big.Int) (*types.Transaction, error) {
	return _BUSD.contract.Transact(opts, "addAdmin", a, maxIssuingTokenPerTime, maxTotalIssuingToken)
}

// AddAdmin is a paid mutator transaction binding the contract method 0x46f2ed2e.
//
// Solidity: function addAdmin(address a, uint256 maxIssuingTokenPerTime, uint256 maxTotalIssuingToken) returns()
func (_BUSD *BUSDSession) AddAdmin(a common.Address, maxIssuingTokenPerTime *big.Int, maxTotalIssuingToken *big.Int) (*types.Transaction, error) {
	return _BUSD.Contract.AddAdmin(&_BUSD.TransactOpts, a, maxIssuingTokenPerTime, maxTotalIssuingToken)
}

// AddAdmin is a paid mutator transaction binding the contract method 0x46f2ed2e.
//
// Solidity: function addAdmin(address a, uint256 maxIssuingTokenPerTime, uint256 maxTotalIssuingToken) returns()
func (_BUSD *BUSDTransactorSession) AddAdmin(a common.Address, maxIssuingTokenPerTime *big.Int, maxTotalIssuingToken *big.Int) (*types.Transaction, error) {
	return _BUSD.Contract.AddAdmin(&_BUSD.TransactOpts, a, maxIssuingTokenPerTime, maxTotalIssuingToken)
}

// AddBlackList is a paid mutator transaction binding the contract method 0x0ecb93c0.
//
// Solidity: function addBlackList(address _evilUser) returns()
func (_BUSD *BUSDTransactor) AddBlackList(opts *bind.TransactOpts, _evilUser common.Address) (*types.Transaction, error) {
	return _BUSD.contract.Transact(opts, "addBlackList", _evilUser)
}

// AddBlackList is a paid mutator transaction binding the contract method 0x0ecb93c0.
//
// Solidity: function addBlackList(address _evilUser) returns()
func (_BUSD *BUSDSession) AddBlackList(_evilUser common.Address) (*types.Transaction, error) {
	return _BUSD.Contract.AddBlackList(&_BUSD.TransactOpts, _evilUser)
}

// AddBlackList is a paid mutator transaction binding the contract method 0x0ecb93c0.
//
// Solidity: function addBlackList(address _evilUser) returns()
func (_BUSD *BUSDTransactorSession) AddBlackList(_evilUser common.Address) (*types.Transaction, error) {
	return _BUSD.Contract.AddBlackList(&_BUSD.TransactOpts, _evilUser)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_BUSD *BUSDTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BUSD.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_BUSD *BUSDSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BUSD.Contract.Approve(&_BUSD.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_BUSD *BUSDTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BUSD.Contract.Approve(&_BUSD.TransactOpts, spender, amount)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_BUSD *BUSDTransactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _BUSD.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_BUSD *BUSDSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _BUSD.Contract.DecreaseAllowance(&_BUSD.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_BUSD *BUSDTransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _BUSD.Contract.DecreaseAllowance(&_BUSD.TransactOpts, spender, subtractedValue)
}

// DestroyBlackFunds is a paid mutator transaction binding the contract method 0xf3bdc228.
//
// Solidity: function destroyBlackFunds(address _blackListedUser) returns()
func (_BUSD *BUSDTransactor) DestroyBlackFunds(opts *bind.TransactOpts, _blackListedUser common.Address) (*types.Transaction, error) {
	return _BUSD.contract.Transact(opts, "destroyBlackFunds", _blackListedUser)
}

// DestroyBlackFunds is a paid mutator transaction binding the contract method 0xf3bdc228.
//
// Solidity: function destroyBlackFunds(address _blackListedUser) returns()
func (_BUSD *BUSDSession) DestroyBlackFunds(_blackListedUser common.Address) (*types.Transaction, error) {
	return _BUSD.Contract.DestroyBlackFunds(&_BUSD.TransactOpts, _blackListedUser)
}

// DestroyBlackFunds is a paid mutator transaction binding the contract method 0xf3bdc228.
//
// Solidity: function destroyBlackFunds(address _blackListedUser) returns()
func (_BUSD *BUSDTransactorSession) DestroyBlackFunds(_blackListedUser common.Address) (*types.Transaction, error) {
	return _BUSD.Contract.DestroyBlackFunds(&_BUSD.TransactOpts, _blackListedUser)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_BUSD *BUSDTransactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _BUSD.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_BUSD *BUSDSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _BUSD.Contract.IncreaseAllowance(&_BUSD.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_BUSD *BUSDTransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _BUSD.Contract.IncreaseAllowance(&_BUSD.TransactOpts, spender, addedValue)
}

// Issue is a paid mutator transaction binding the contract method 0xea3f068d.
//
// Solidity: function issue(address issuer, uint256 value, bytes32 offchain) returns()
func (_BUSD *BUSDTransactor) Issue(opts *bind.TransactOpts, issuer common.Address, value *big.Int, offchain [32]byte) (*types.Transaction, error) {
	return _BUSD.contract.Transact(opts, "issue", issuer, value, offchain)
}

// Issue is a paid mutator transaction binding the contract method 0xea3f068d.
//
// Solidity: function issue(address issuer, uint256 value, bytes32 offchain) returns()
func (_BUSD *BUSDSession) Issue(issuer common.Address, value *big.Int, offchain [32]byte) (*types.Transaction, error) {
	return _BUSD.Contract.Issue(&_BUSD.TransactOpts, issuer, value, offchain)
}

// Issue is a paid mutator transaction binding the contract method 0xea3f068d.
//
// Solidity: function issue(address issuer, uint256 value, bytes32 offchain) returns()
func (_BUSD *BUSDTransactorSession) Issue(issuer common.Address, value *big.Int, offchain [32]byte) (*types.Transaction, error) {
	return _BUSD.Contract.Issue(&_BUSD.TransactOpts, issuer, value, offchain)
}

// IssueByAdmin is a paid mutator transaction binding the contract method 0xd10601ec.
//
// Solidity: function issueByAdmin(address issuer, uint256 value, bytes32 offchain) returns()
func (_BUSD *BUSDTransactor) IssueByAdmin(opts *bind.TransactOpts, issuer common.Address, value *big.Int, offchain [32]byte) (*types.Transaction, error) {
	return _BUSD.contract.Transact(opts, "issueByAdmin", issuer, value, offchain)
}

// IssueByAdmin is a paid mutator transaction binding the contract method 0xd10601ec.
//
// Solidity: function issueByAdmin(address issuer, uint256 value, bytes32 offchain) returns()
func (_BUSD *BUSDSession) IssueByAdmin(issuer common.Address, value *big.Int, offchain [32]byte) (*types.Transaction, error) {
	return _BUSD.Contract.IssueByAdmin(&_BUSD.TransactOpts, issuer, value, offchain)
}

// IssueByAdmin is a paid mutator transaction binding the contract method 0xd10601ec.
//
// Solidity: function issueByAdmin(address issuer, uint256 value, bytes32 offchain) returns()
func (_BUSD *BUSDTransactorSession) IssueByAdmin(issuer common.Address, value *big.Int, offchain [32]byte) (*types.Transaction, error) {
	return _BUSD.Contract.IssueByAdmin(&_BUSD.TransactOpts, issuer, value, offchain)
}

// Mint is a paid mutator transaction binding the contract method 0xa0712d68.
//
// Solidity: function mint(uint256 amount) returns(bool)
func (_BUSD *BUSDTransactor) Mint(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _BUSD.contract.Transact(opts, "mint", amount)
}

// Mint is a paid mutator transaction binding the contract method 0xa0712d68.
//
// Solidity: function mint(uint256 amount) returns(bool)
func (_BUSD *BUSDSession) Mint(amount *big.Int) (*types.Transaction, error) {
	return _BUSD.Contract.Mint(&_BUSD.TransactOpts, amount)
}

// Mint is a paid mutator transaction binding the contract method 0xa0712d68.
//
// Solidity: function mint(uint256 amount) returns(bool)
func (_BUSD *BUSDTransactorSession) Mint(amount *big.Int) (*types.Transaction, error) {
	return _BUSD.Contract.Mint(&_BUSD.TransactOpts, amount)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_BUSD *BUSDTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BUSD.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_BUSD *BUSDSession) Pause() (*types.Transaction, error) {
	return _BUSD.Contract.Pause(&_BUSD.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_BUSD *BUSDTransactorSession) Pause() (*types.Transaction, error) {
	return _BUSD.Contract.Pause(&_BUSD.TransactOpts)
}

// Redeem is a paid mutator transaction binding the contract method 0x992c3e4b.
//
// Solidity: function redeem(address redeemer, uint256 value, bytes32 offchain) returns()
func (_BUSD *BUSDTransactor) Redeem(opts *bind.TransactOpts, redeemer common.Address, value *big.Int, offchain [32]byte) (*types.Transaction, error) {
	return _BUSD.contract.Transact(opts, "redeem", redeemer, value, offchain)
}

// Redeem is a paid mutator transaction binding the contract method 0x992c3e4b.
//
// Solidity: function redeem(address redeemer, uint256 value, bytes32 offchain) returns()
func (_BUSD *BUSDSession) Redeem(redeemer common.Address, value *big.Int, offchain [32]byte) (*types.Transaction, error) {
	return _BUSD.Contract.Redeem(&_BUSD.TransactOpts, redeemer, value, offchain)
}

// Redeem is a paid mutator transaction binding the contract method 0x992c3e4b.
//
// Solidity: function redeem(address redeemer, uint256 value, bytes32 offchain) returns()
func (_BUSD *BUSDTransactorSession) Redeem(redeemer common.Address, value *big.Int, offchain [32]byte) (*types.Transaction, error) {
	return _BUSD.Contract.Redeem(&_BUSD.TransactOpts, redeemer, value, offchain)
}

// RemoveAdmin is a paid mutator transaction binding the contract method 0x1785f53c.
//
// Solidity: function removeAdmin(address a) returns()
func (_BUSD *BUSDTransactor) RemoveAdmin(opts *bind.TransactOpts, a common.Address) (*types.Transaction, error) {
	return _BUSD.contract.Transact(opts, "removeAdmin", a)
}

// RemoveAdmin is a paid mutator transaction binding the contract method 0x1785f53c.
//
// Solidity: function removeAdmin(address a) returns()
func (_BUSD *BUSDSession) RemoveAdmin(a common.Address) (*types.Transaction, error) {
	return _BUSD.Contract.RemoveAdmin(&_BUSD.TransactOpts, a)
}

// RemoveAdmin is a paid mutator transaction binding the contract method 0x1785f53c.
//
// Solidity: function removeAdmin(address a) returns()
func (_BUSD *BUSDTransactorSession) RemoveAdmin(a common.Address) (*types.Transaction, error) {
	return _BUSD.Contract.RemoveAdmin(&_BUSD.TransactOpts, a)
}

// RemoveBlackList is a paid mutator transaction binding the contract method 0xe4997dc5.
//
// Solidity: function removeBlackList(address _clearedUser) returns()
func (_BUSD *BUSDTransactor) RemoveBlackList(opts *bind.TransactOpts, _clearedUser common.Address) (*types.Transaction, error) {
	return _BUSD.contract.Transact(opts, "removeBlackList", _clearedUser)
}

// RemoveBlackList is a paid mutator transaction binding the contract method 0xe4997dc5.
//
// Solidity: function removeBlackList(address _clearedUser) returns()
func (_BUSD *BUSDSession) RemoveBlackList(_clearedUser common.Address) (*types.Transaction, error) {
	return _BUSD.Contract.RemoveBlackList(&_BUSD.TransactOpts, _clearedUser)
}

// RemoveBlackList is a paid mutator transaction binding the contract method 0xe4997dc5.
//
// Solidity: function removeBlackList(address _clearedUser) returns()
func (_BUSD *BUSDTransactorSession) RemoveBlackList(_clearedUser common.Address) (*types.Transaction, error) {
	return _BUSD.Contract.RemoveBlackList(&_BUSD.TransactOpts, _clearedUser)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_BUSD *BUSDTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BUSD.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_BUSD *BUSDSession) RenounceOwnership() (*types.Transaction, error) {
	return _BUSD.Contract.RenounceOwnership(&_BUSD.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_BUSD *BUSDTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _BUSD.Contract.RenounceOwnership(&_BUSD.TransactOpts)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns(bool)
func (_BUSD *BUSDTransactor) Transfer(opts *bind.TransactOpts, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _BUSD.contract.Transact(opts, "transfer", _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns(bool)
func (_BUSD *BUSDSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _BUSD.Contract.Transfer(&_BUSD.TransactOpts, _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns(bool)
func (_BUSD *BUSDTransactorSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _BUSD.Contract.Transfer(&_BUSD.TransactOpts, _to, _value)
}

// TransferByAdmin is a paid mutator transaction binding the contract method 0xf9ba884f.
//
// Solidity: function transferByAdmin(address from, address to, uint256 value, bytes32 offchain) returns()
func (_BUSD *BUSDTransactor) TransferByAdmin(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int, offchain [32]byte) (*types.Transaction, error) {
	return _BUSD.contract.Transact(opts, "transferByAdmin", from, to, value, offchain)
}

// TransferByAdmin is a paid mutator transaction binding the contract method 0xf9ba884f.
//
// Solidity: function transferByAdmin(address from, address to, uint256 value, bytes32 offchain) returns()
func (_BUSD *BUSDSession) TransferByAdmin(from common.Address, to common.Address, value *big.Int, offchain [32]byte) (*types.Transaction, error) {
	return _BUSD.Contract.TransferByAdmin(&_BUSD.TransactOpts, from, to, value, offchain)
}

// TransferByAdmin is a paid mutator transaction binding the contract method 0xf9ba884f.
//
// Solidity: function transferByAdmin(address from, address to, uint256 value, bytes32 offchain) returns()
func (_BUSD *BUSDTransactorSession) TransferByAdmin(from common.Address, to common.Address, value *big.Int, offchain [32]byte) (*types.Transaction, error) {
	return _BUSD.Contract.TransferByAdmin(&_BUSD.TransactOpts, from, to, value, offchain)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _value) returns(bool)
func (_BUSD *BUSDTransactor) TransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _BUSD.contract.Transact(opts, "transferFrom", _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _value) returns(bool)
func (_BUSD *BUSDSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _BUSD.Contract.TransferFrom(&_BUSD.TransactOpts, _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _value) returns(bool)
func (_BUSD *BUSDTransactorSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _BUSD.Contract.TransferFrom(&_BUSD.TransactOpts, _from, _to, _value)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_BUSD *BUSDTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _BUSD.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_BUSD *BUSDSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _BUSD.Contract.TransferOwnership(&_BUSD.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_BUSD *BUSDTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _BUSD.Contract.TransferOwnership(&_BUSD.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_BUSD *BUSDTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BUSD.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_BUSD *BUSDSession) Unpause() (*types.Transaction, error) {
	return _BUSD.Contract.Unpause(&_BUSD.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_BUSD *BUSDTransactorSession) Unpause() (*types.Transaction, error) {
	return _BUSD.Contract.Unpause(&_BUSD.TransactOpts)
}

// BUSDAddedBlackListIterator is returned from FilterAddedBlackList and is used to iterate over the raw logs and unpacked data for AddedBlackList events raised by the BUSD contract.
type BUSDAddedBlackListIterator struct {
	Event *BUSDAddedBlackList // Event containing the contract specifics and raw log

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
func (it *BUSDAddedBlackListIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BUSDAddedBlackList)
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
		it.Event = new(BUSDAddedBlackList)
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
func (it *BUSDAddedBlackListIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BUSDAddedBlackListIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BUSDAddedBlackList represents a AddedBlackList event raised by the BUSD contract.
type BUSDAddedBlackList struct {
	User common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterAddedBlackList is a free log retrieval operation binding the contract event 0x42e160154868087d6bfdc0ca23d96a1c1cfa32f1b72ba9ba27b69b98a0d819dc.
//
// Solidity: event AddedBlackList(address _user)
func (_BUSD *BUSDFilterer) FilterAddedBlackList(opts *bind.FilterOpts) (*BUSDAddedBlackListIterator, error) {

	logs, sub, err := _BUSD.contract.FilterLogs(opts, "AddedBlackList")
	if err != nil {
		return nil, err
	}
	return &BUSDAddedBlackListIterator{contract: _BUSD.contract, event: "AddedBlackList", logs: logs, sub: sub}, nil
}

// WatchAddedBlackList is a free log subscription operation binding the contract event 0x42e160154868087d6bfdc0ca23d96a1c1cfa32f1b72ba9ba27b69b98a0d819dc.
//
// Solidity: event AddedBlackList(address _user)
func (_BUSD *BUSDFilterer) WatchAddedBlackList(opts *bind.WatchOpts, sink chan<- *BUSDAddedBlackList) (event.Subscription, error) {

	logs, sub, err := _BUSD.contract.WatchLogs(opts, "AddedBlackList")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BUSDAddedBlackList)
				if err := _BUSD.contract.UnpackLog(event, "AddedBlackList", log); err != nil {
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
func (_BUSD *BUSDFilterer) ParseAddedBlackList(log types.Log) (*BUSDAddedBlackList, error) {
	event := new(BUSDAddedBlackList)
	if err := _BUSD.contract.UnpackLog(event, "AddedBlackList", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BUSDApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the BUSD contract.
type BUSDApprovalIterator struct {
	Event *BUSDApproval // Event containing the contract specifics and raw log

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
func (it *BUSDApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BUSDApproval)
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
		it.Event = new(BUSDApproval)
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
func (it *BUSDApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BUSDApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BUSDApproval represents a Approval event raised by the BUSD contract.
type BUSDApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_BUSD *BUSDFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*BUSDApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _BUSD.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &BUSDApprovalIterator{contract: _BUSD.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_BUSD *BUSDFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *BUSDApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _BUSD.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BUSDApproval)
				if err := _BUSD.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_BUSD *BUSDFilterer) ParseApproval(log types.Log) (*BUSDApproval, error) {
	event := new(BUSDApproval)
	if err := _BUSD.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BUSDDestroyedBlackFundsIterator is returned from FilterDestroyedBlackFunds and is used to iterate over the raw logs and unpacked data for DestroyedBlackFunds events raised by the BUSD contract.
type BUSDDestroyedBlackFundsIterator struct {
	Event *BUSDDestroyedBlackFunds // Event containing the contract specifics and raw log

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
func (it *BUSDDestroyedBlackFundsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BUSDDestroyedBlackFunds)
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
		it.Event = new(BUSDDestroyedBlackFunds)
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
func (it *BUSDDestroyedBlackFundsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BUSDDestroyedBlackFundsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BUSDDestroyedBlackFunds represents a DestroyedBlackFunds event raised by the BUSD contract.
type BUSDDestroyedBlackFunds struct {
	BlackListedUser common.Address
	Balance         *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterDestroyedBlackFunds is a free log retrieval operation binding the contract event 0x61e6e66b0d6339b2980aecc6ccc0039736791f0ccde9ed512e789a7fbdd698c6.
//
// Solidity: event DestroyedBlackFunds(address _blackListedUser, uint256 _balance)
func (_BUSD *BUSDFilterer) FilterDestroyedBlackFunds(opts *bind.FilterOpts) (*BUSDDestroyedBlackFundsIterator, error) {

	logs, sub, err := _BUSD.contract.FilterLogs(opts, "DestroyedBlackFunds")
	if err != nil {
		return nil, err
	}
	return &BUSDDestroyedBlackFundsIterator{contract: _BUSD.contract, event: "DestroyedBlackFunds", logs: logs, sub: sub}, nil
}

// WatchDestroyedBlackFunds is a free log subscription operation binding the contract event 0x61e6e66b0d6339b2980aecc6ccc0039736791f0ccde9ed512e789a7fbdd698c6.
//
// Solidity: event DestroyedBlackFunds(address _blackListedUser, uint256 _balance)
func (_BUSD *BUSDFilterer) WatchDestroyedBlackFunds(opts *bind.WatchOpts, sink chan<- *BUSDDestroyedBlackFunds) (event.Subscription, error) {

	logs, sub, err := _BUSD.contract.WatchLogs(opts, "DestroyedBlackFunds")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BUSDDestroyedBlackFunds)
				if err := _BUSD.contract.UnpackLog(event, "DestroyedBlackFunds", log); err != nil {
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
func (_BUSD *BUSDFilterer) ParseDestroyedBlackFunds(log types.Log) (*BUSDDestroyedBlackFunds, error) {
	event := new(BUSDDestroyedBlackFunds)
	if err := _BUSD.contract.UnpackLog(event, "DestroyedBlackFunds", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BUSDOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the BUSD contract.
type BUSDOwnershipTransferredIterator struct {
	Event *BUSDOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *BUSDOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BUSDOwnershipTransferred)
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
		it.Event = new(BUSDOwnershipTransferred)
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
func (it *BUSDOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BUSDOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BUSDOwnershipTransferred represents a OwnershipTransferred event raised by the BUSD contract.
type BUSDOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_BUSD *BUSDFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*BUSDOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _BUSD.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &BUSDOwnershipTransferredIterator{contract: _BUSD.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_BUSD *BUSDFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *BUSDOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _BUSD.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BUSDOwnershipTransferred)
				if err := _BUSD.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_BUSD *BUSDFilterer) ParseOwnershipTransferred(log types.Log) (*BUSDOwnershipTransferred, error) {
	event := new(BUSDOwnershipTransferred)
	if err := _BUSD.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BUSDPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the BUSD contract.
type BUSDPausedIterator struct {
	Event *BUSDPaused // Event containing the contract specifics and raw log

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
func (it *BUSDPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BUSDPaused)
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
		it.Event = new(BUSDPaused)
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
func (it *BUSDPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BUSDPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BUSDPaused represents a Paused event raised by the BUSD contract.
type BUSDPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_BUSD *BUSDFilterer) FilterPaused(opts *bind.FilterOpts) (*BUSDPausedIterator, error) {

	logs, sub, err := _BUSD.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &BUSDPausedIterator{contract: _BUSD.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_BUSD *BUSDFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *BUSDPaused) (event.Subscription, error) {

	logs, sub, err := _BUSD.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BUSDPaused)
				if err := _BUSD.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_BUSD *BUSDFilterer) ParsePaused(log types.Log) (*BUSDPaused, error) {
	event := new(BUSDPaused)
	if err := _BUSD.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BUSDRemovedBlackListIterator is returned from FilterRemovedBlackList and is used to iterate over the raw logs and unpacked data for RemovedBlackList events raised by the BUSD contract.
type BUSDRemovedBlackListIterator struct {
	Event *BUSDRemovedBlackList // Event containing the contract specifics and raw log

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
func (it *BUSDRemovedBlackListIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BUSDRemovedBlackList)
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
		it.Event = new(BUSDRemovedBlackList)
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
func (it *BUSDRemovedBlackListIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BUSDRemovedBlackListIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BUSDRemovedBlackList represents a RemovedBlackList event raised by the BUSD contract.
type BUSDRemovedBlackList struct {
	User common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterRemovedBlackList is a free log retrieval operation binding the contract event 0xd7e9ec6e6ecd65492dce6bf513cd6867560d49544421d0783ddf06e76c24470c.
//
// Solidity: event RemovedBlackList(address _user)
func (_BUSD *BUSDFilterer) FilterRemovedBlackList(opts *bind.FilterOpts) (*BUSDRemovedBlackListIterator, error) {

	logs, sub, err := _BUSD.contract.FilterLogs(opts, "RemovedBlackList")
	if err != nil {
		return nil, err
	}
	return &BUSDRemovedBlackListIterator{contract: _BUSD.contract, event: "RemovedBlackList", logs: logs, sub: sub}, nil
}

// WatchRemovedBlackList is a free log subscription operation binding the contract event 0xd7e9ec6e6ecd65492dce6bf513cd6867560d49544421d0783ddf06e76c24470c.
//
// Solidity: event RemovedBlackList(address _user)
func (_BUSD *BUSDFilterer) WatchRemovedBlackList(opts *bind.WatchOpts, sink chan<- *BUSDRemovedBlackList) (event.Subscription, error) {

	logs, sub, err := _BUSD.contract.WatchLogs(opts, "RemovedBlackList")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BUSDRemovedBlackList)
				if err := _BUSD.contract.UnpackLog(event, "RemovedBlackList", log); err != nil {
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
func (_BUSD *BUSDFilterer) ParseRemovedBlackList(log types.Log) (*BUSDRemovedBlackList, error) {
	event := new(BUSDRemovedBlackList)
	if err := _BUSD.contract.UnpackLog(event, "RemovedBlackList", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BUSDTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the BUSD contract.
type BUSDTransferIterator struct {
	Event *BUSDTransfer // Event containing the contract specifics and raw log

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
func (it *BUSDTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BUSDTransfer)
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
		it.Event = new(BUSDTransfer)
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
func (it *BUSDTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BUSDTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BUSDTransfer represents a Transfer event raised by the BUSD contract.
type BUSDTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_BUSD *BUSDFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*BUSDTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _BUSD.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &BUSDTransferIterator{contract: _BUSD.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_BUSD *BUSDFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *BUSDTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _BUSD.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BUSDTransfer)
				if err := _BUSD.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_BUSD *BUSDFilterer) ParseTransfer(log types.Log) (*BUSDTransfer, error) {
	event := new(BUSDTransfer)
	if err := _BUSD.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BUSDUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the BUSD contract.
type BUSDUnpausedIterator struct {
	Event *BUSDUnpaused // Event containing the contract specifics and raw log

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
func (it *BUSDUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BUSDUnpaused)
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
		it.Event = new(BUSDUnpaused)
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
func (it *BUSDUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BUSDUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BUSDUnpaused represents a Unpaused event raised by the BUSD contract.
type BUSDUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_BUSD *BUSDFilterer) FilterUnpaused(opts *bind.FilterOpts) (*BUSDUnpausedIterator, error) {

	logs, sub, err := _BUSD.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &BUSDUnpausedIterator{contract: _BUSD.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_BUSD *BUSDFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *BUSDUnpaused) (event.Subscription, error) {

	logs, sub, err := _BUSD.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BUSDUnpaused)
				if err := _BUSD.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
func (_BUSD *BUSDFilterer) ParseUnpaused(log types.Log) (*BUSDUnpaused, error) {
	event := new(BUSDUnpaused)
	if err := _BUSD.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BUSDIssueIterator is returned from FilterIssue and is used to iterate over the raw logs and unpacked data for Issue events raised by the BUSD contract.
type BUSDIssueIterator struct {
	Event *BUSDIssue // Event containing the contract specifics and raw log

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
func (it *BUSDIssueIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BUSDIssue)
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
		it.Event = new(BUSDIssue)
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
func (it *BUSDIssueIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BUSDIssueIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BUSDIssue represents a Issue event raised by the BUSD contract.
type BUSDIssue struct {
	Offchain [32]byte
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterIssue is a free log retrieval operation binding the contract event 0x182a88d3fbfcb38ffaddf26e6d63ecd93b4348141de1eb744c43c4fbb1ccfb22.
//
// Solidity: event __issue(bytes32 offchain)
func (_BUSD *BUSDFilterer) FilterIssue(opts *bind.FilterOpts) (*BUSDIssueIterator, error) {

	logs, sub, err := _BUSD.contract.FilterLogs(opts, "__issue")
	if err != nil {
		return nil, err
	}
	return &BUSDIssueIterator{contract: _BUSD.contract, event: "__issue", logs: logs, sub: sub}, nil
}

// WatchIssue is a free log subscription operation binding the contract event 0x182a88d3fbfcb38ffaddf26e6d63ecd93b4348141de1eb744c43c4fbb1ccfb22.
//
// Solidity: event __issue(bytes32 offchain)
func (_BUSD *BUSDFilterer) WatchIssue(opts *bind.WatchOpts, sink chan<- *BUSDIssue) (event.Subscription, error) {

	logs, sub, err := _BUSD.contract.WatchLogs(opts, "__issue")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BUSDIssue)
				if err := _BUSD.contract.UnpackLog(event, "__issue", log); err != nil {
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
func (_BUSD *BUSDFilterer) ParseIssue(log types.Log) (*BUSDIssue, error) {
	event := new(BUSDIssue)
	if err := _BUSD.contract.UnpackLog(event, "__issue", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BUSDRedeemIterator is returned from FilterRedeem and is used to iterate over the raw logs and unpacked data for Redeem events raised by the BUSD contract.
type BUSDRedeemIterator struct {
	Event *BUSDRedeem // Event containing the contract specifics and raw log

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
func (it *BUSDRedeemIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BUSDRedeem)
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
		it.Event = new(BUSDRedeem)
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
func (it *BUSDRedeemIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BUSDRedeemIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BUSDRedeem represents a Redeem event raised by the BUSD contract.
type BUSDRedeem struct {
	Offchain [32]byte
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterRedeem is a free log retrieval operation binding the contract event 0xb0de879351469d2741406aafc9ba1f44eb957cf44ee3391e59a7a9097050c927.
//
// Solidity: event __redeem(bytes32 offchain)
func (_BUSD *BUSDFilterer) FilterRedeem(opts *bind.FilterOpts) (*BUSDRedeemIterator, error) {

	logs, sub, err := _BUSD.contract.FilterLogs(opts, "__redeem")
	if err != nil {
		return nil, err
	}
	return &BUSDRedeemIterator{contract: _BUSD.contract, event: "__redeem", logs: logs, sub: sub}, nil
}

// WatchRedeem is a free log subscription operation binding the contract event 0xb0de879351469d2741406aafc9ba1f44eb957cf44ee3391e59a7a9097050c927.
//
// Solidity: event __redeem(bytes32 offchain)
func (_BUSD *BUSDFilterer) WatchRedeem(opts *bind.WatchOpts, sink chan<- *BUSDRedeem) (event.Subscription, error) {

	logs, sub, err := _BUSD.contract.WatchLogs(opts, "__redeem")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BUSDRedeem)
				if err := _BUSD.contract.UnpackLog(event, "__redeem", log); err != nil {
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
func (_BUSD *BUSDFilterer) ParseRedeem(log types.Log) (*BUSDRedeem, error) {
	event := new(BUSDRedeem)
	if err := _BUSD.contract.UnpackLog(event, "__redeem", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BUSDTransferByAdminIterator is returned from FilterTransferByAdmin and is used to iterate over the raw logs and unpacked data for TransferByAdmin events raised by the BUSD contract.
type BUSDTransferByAdminIterator struct {
	Event *BUSDTransferByAdmin // Event containing the contract specifics and raw log

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
func (it *BUSDTransferByAdminIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BUSDTransferByAdmin)
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
		it.Event = new(BUSDTransferByAdmin)
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
func (it *BUSDTransferByAdminIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BUSDTransferByAdminIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BUSDTransferByAdmin represents a TransferByAdmin event raised by the BUSD contract.
type BUSDTransferByAdmin struct {
	Offchain [32]byte
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterTransferByAdmin is a free log retrieval operation binding the contract event 0x3b36ee6b35325f38e95938557be92853c842b7a9a19fd7ac4931a6d24db52682.
//
// Solidity: event __transferByAdmin(bytes32 offchain)
func (_BUSD *BUSDFilterer) FilterTransferByAdmin(opts *bind.FilterOpts) (*BUSDTransferByAdminIterator, error) {

	logs, sub, err := _BUSD.contract.FilterLogs(opts, "__transferByAdmin")
	if err != nil {
		return nil, err
	}
	return &BUSDTransferByAdminIterator{contract: _BUSD.contract, event: "__transferByAdmin", logs: logs, sub: sub}, nil
}

// WatchTransferByAdmin is a free log subscription operation binding the contract event 0x3b36ee6b35325f38e95938557be92853c842b7a9a19fd7ac4931a6d24db52682.
//
// Solidity: event __transferByAdmin(bytes32 offchain)
func (_BUSD *BUSDFilterer) WatchTransferByAdmin(opts *bind.WatchOpts, sink chan<- *BUSDTransferByAdmin) (event.Subscription, error) {

	logs, sub, err := _BUSD.contract.WatchLogs(opts, "__transferByAdmin")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BUSDTransferByAdmin)
				if err := _BUSD.contract.UnpackLog(event, "__transferByAdmin", log); err != nil {
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
func (_BUSD *BUSDFilterer) ParseTransferByAdmin(log types.Log) (*BUSDTransferByAdmin, error) {
	event := new(BUSDTransferByAdmin)
	if err := _BUSD.contract.UnpackLog(event, "__transferByAdmin", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
