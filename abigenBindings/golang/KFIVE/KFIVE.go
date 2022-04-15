// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package KFIVE

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

// KFIVEMetaData contains all meta data concerning the KFIVE contract.
var KFIVEMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"AddedBlackList\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_blackListedUser\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_balance\",\"type\":\"uint256\"}],\"name\":\"DestroyedBlackFunds\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"RemovedBlackList\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"offchain\",\"type\":\"bytes32\"}],\"name\":\"__issue\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"offchain\",\"type\":\"bytes32\"}],\"name\":\"__redeem\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"offchain\",\"type\":\"bytes32\"}],\"name\":\"__transferByAdmin\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"a\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"maxIssuingTokenPerTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxTotalIssuingToken\",\"type\":\"uint256\"}],\"name\":\"addAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_evilUser\",\"type\":\"address\"}],\"name\":\"addBlackList\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"admin\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"status\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"maxIssuingTokenPerTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxTotalIssuingToken\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"remainingIssuingToken\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_maker\",\"type\":\"address\"}],\"name\":\"getBlackListStatus\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"isBlackListed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"a\",\"type\":\"address\"}],\"name\":\"removeAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_clearedUser\",\"type\":\"address\"}],\"name\":\"removeBlackList\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"offchain\",\"type\":\"bytes32\"}],\"name\":\"transferByAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"issuer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"offchain\",\"type\":\"bytes32\"}],\"name\":\"issue\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"issuer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"offchain\",\"type\":\"bytes32\"}],\"name\":\"issueByAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_blackListedUser\",\"type\":\"address\"}],\"name\":\"destroyBlackFunds\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"redeemer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"offchain\",\"type\":\"bytes32\"}],\"name\":\"redeem\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// KFIVEABI is the input ABI used to generate the binding from.
// Deprecated: Use KFIVEMetaData.ABI instead.
var KFIVEABI = KFIVEMetaData.ABI

// KFIVE is an auto generated Go binding around an Ethereum contract.
type KFIVE struct {
	KFIVECaller     // Read-only binding to the contract
	KFIVETransactor // Write-only binding to the contract
	KFIVEFilterer   // Log filterer for contract events
}

// KFIVECaller is an auto generated read-only Go binding around an Ethereum contract.
type KFIVECaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KFIVETransactor is an auto generated write-only Go binding around an Ethereum contract.
type KFIVETransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KFIVEFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type KFIVEFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KFIVESession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type KFIVESession struct {
	Contract     *KFIVE            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// KFIVECallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type KFIVECallerSession struct {
	Contract *KFIVECaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// KFIVETransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type KFIVETransactorSession struct {
	Contract     *KFIVETransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// KFIVERaw is an auto generated low-level Go binding around an Ethereum contract.
type KFIVERaw struct {
	Contract *KFIVE // Generic contract binding to access the raw methods on
}

// KFIVECallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type KFIVECallerRaw struct {
	Contract *KFIVECaller // Generic read-only contract binding to access the raw methods on
}

// KFIVETransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type KFIVETransactorRaw struct {
	Contract *KFIVETransactor // Generic write-only contract binding to access the raw methods on
}

// NewKFIVE creates a new instance of KFIVE, bound to a specific deployed contract.
func NewKFIVE(address common.Address, backend bind.ContractBackend) (*KFIVE, error) {
	contract, err := bindKFIVE(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &KFIVE{KFIVECaller: KFIVECaller{contract: contract}, KFIVETransactor: KFIVETransactor{contract: contract}, KFIVEFilterer: KFIVEFilterer{contract: contract}}, nil
}

// NewKFIVECaller creates a new read-only instance of KFIVE, bound to a specific deployed contract.
func NewKFIVECaller(address common.Address, caller bind.ContractCaller) (*KFIVECaller, error) {
	contract, err := bindKFIVE(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &KFIVECaller{contract: contract}, nil
}

// NewKFIVETransactor creates a new write-only instance of KFIVE, bound to a specific deployed contract.
func NewKFIVETransactor(address common.Address, transactor bind.ContractTransactor) (*KFIVETransactor, error) {
	contract, err := bindKFIVE(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &KFIVETransactor{contract: contract}, nil
}

// NewKFIVEFilterer creates a new log filterer instance of KFIVE, bound to a specific deployed contract.
func NewKFIVEFilterer(address common.Address, filterer bind.ContractFilterer) (*KFIVEFilterer, error) {
	contract, err := bindKFIVE(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &KFIVEFilterer{contract: contract}, nil
}

// bindKFIVE binds a generic wrapper to an already deployed contract.
func bindKFIVE(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(KFIVEABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_KFIVE *KFIVERaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _KFIVE.Contract.KFIVECaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_KFIVE *KFIVERaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _KFIVE.Contract.KFIVETransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_KFIVE *KFIVERaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _KFIVE.Contract.KFIVETransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_KFIVE *KFIVECallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _KFIVE.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_KFIVE *KFIVETransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _KFIVE.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_KFIVE *KFIVETransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _KFIVE.Contract.contract.Transact(opts, method, params...)
}

// Admin is a free data retrieval call binding the contract method 0x63a846f8.
//
// Solidity: function admin(address ) view returns(bool status, uint256 maxIssuingTokenPerTime, uint256 maxTotalIssuingToken, uint256 remainingIssuingToken)
func (_KFIVE *KFIVECaller) Admin(opts *bind.CallOpts, arg0 common.Address) (struct {
	Status                 bool
	MaxIssuingTokenPerTime *big.Int
	MaxTotalIssuingToken   *big.Int
	RemainingIssuingToken  *big.Int
}, error) {
	var out []interface{}
	err := _KFIVE.contract.Call(opts, &out, "admin", arg0)

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
func (_KFIVE *KFIVESession) Admin(arg0 common.Address) (struct {
	Status                 bool
	MaxIssuingTokenPerTime *big.Int
	MaxTotalIssuingToken   *big.Int
	RemainingIssuingToken  *big.Int
}, error) {
	return _KFIVE.Contract.Admin(&_KFIVE.CallOpts, arg0)
}

// Admin is a free data retrieval call binding the contract method 0x63a846f8.
//
// Solidity: function admin(address ) view returns(bool status, uint256 maxIssuingTokenPerTime, uint256 maxTotalIssuingToken, uint256 remainingIssuingToken)
func (_KFIVE *KFIVECallerSession) Admin(arg0 common.Address) (struct {
	Status                 bool
	MaxIssuingTokenPerTime *big.Int
	MaxTotalIssuingToken   *big.Int
	RemainingIssuingToken  *big.Int
}, error) {
	return _KFIVE.Contract.Admin(&_KFIVE.CallOpts, arg0)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_KFIVE *KFIVECaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _KFIVE.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_KFIVE *KFIVESession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _KFIVE.Contract.Allowance(&_KFIVE.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_KFIVE *KFIVECallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _KFIVE.Contract.Allowance(&_KFIVE.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_KFIVE *KFIVECaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _KFIVE.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_KFIVE *KFIVESession) BalanceOf(account common.Address) (*big.Int, error) {
	return _KFIVE.Contract.BalanceOf(&_KFIVE.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_KFIVE *KFIVECallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _KFIVE.Contract.BalanceOf(&_KFIVE.CallOpts, account)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_KFIVE *KFIVECaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _KFIVE.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_KFIVE *KFIVESession) Decimals() (uint8, error) {
	return _KFIVE.Contract.Decimals(&_KFIVE.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_KFIVE *KFIVECallerSession) Decimals() (uint8, error) {
	return _KFIVE.Contract.Decimals(&_KFIVE.CallOpts)
}

// GetBlackListStatus is a free data retrieval call binding the contract method 0x59bf1abe.
//
// Solidity: function getBlackListStatus(address _maker) view returns(bool)
func (_KFIVE *KFIVECaller) GetBlackListStatus(opts *bind.CallOpts, _maker common.Address) (bool, error) {
	var out []interface{}
	err := _KFIVE.contract.Call(opts, &out, "getBlackListStatus", _maker)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// GetBlackListStatus is a free data retrieval call binding the contract method 0x59bf1abe.
//
// Solidity: function getBlackListStatus(address _maker) view returns(bool)
func (_KFIVE *KFIVESession) GetBlackListStatus(_maker common.Address) (bool, error) {
	return _KFIVE.Contract.GetBlackListStatus(&_KFIVE.CallOpts, _maker)
}

// GetBlackListStatus is a free data retrieval call binding the contract method 0x59bf1abe.
//
// Solidity: function getBlackListStatus(address _maker) view returns(bool)
func (_KFIVE *KFIVECallerSession) GetBlackListStatus(_maker common.Address) (bool, error) {
	return _KFIVE.Contract.GetBlackListStatus(&_KFIVE.CallOpts, _maker)
}

// GetOwner is a free data retrieval call binding the contract method 0x893d20e8.
//
// Solidity: function getOwner() view returns(address)
func (_KFIVE *KFIVECaller) GetOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _KFIVE.contract.Call(opts, &out, "getOwner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetOwner is a free data retrieval call binding the contract method 0x893d20e8.
//
// Solidity: function getOwner() view returns(address)
func (_KFIVE *KFIVESession) GetOwner() (common.Address, error) {
	return _KFIVE.Contract.GetOwner(&_KFIVE.CallOpts)
}

// GetOwner is a free data retrieval call binding the contract method 0x893d20e8.
//
// Solidity: function getOwner() view returns(address)
func (_KFIVE *KFIVECallerSession) GetOwner() (common.Address, error) {
	return _KFIVE.Contract.GetOwner(&_KFIVE.CallOpts)
}

// IsBlackListed is a free data retrieval call binding the contract method 0xe47d6060.
//
// Solidity: function isBlackListed(address ) view returns(bool)
func (_KFIVE *KFIVECaller) IsBlackListed(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _KFIVE.contract.Call(opts, &out, "isBlackListed", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsBlackListed is a free data retrieval call binding the contract method 0xe47d6060.
//
// Solidity: function isBlackListed(address ) view returns(bool)
func (_KFIVE *KFIVESession) IsBlackListed(arg0 common.Address) (bool, error) {
	return _KFIVE.Contract.IsBlackListed(&_KFIVE.CallOpts, arg0)
}

// IsBlackListed is a free data retrieval call binding the contract method 0xe47d6060.
//
// Solidity: function isBlackListed(address ) view returns(bool)
func (_KFIVE *KFIVECallerSession) IsBlackListed(arg0 common.Address) (bool, error) {
	return _KFIVE.Contract.IsBlackListed(&_KFIVE.CallOpts, arg0)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_KFIVE *KFIVECaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _KFIVE.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_KFIVE *KFIVESession) Name() (string, error) {
	return _KFIVE.Contract.Name(&_KFIVE.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_KFIVE *KFIVECallerSession) Name() (string, error) {
	return _KFIVE.Contract.Name(&_KFIVE.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_KFIVE *KFIVECaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _KFIVE.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_KFIVE *KFIVESession) Owner() (common.Address, error) {
	return _KFIVE.Contract.Owner(&_KFIVE.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_KFIVE *KFIVECallerSession) Owner() (common.Address, error) {
	return _KFIVE.Contract.Owner(&_KFIVE.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_KFIVE *KFIVECaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _KFIVE.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_KFIVE *KFIVESession) Paused() (bool, error) {
	return _KFIVE.Contract.Paused(&_KFIVE.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_KFIVE *KFIVECallerSession) Paused() (bool, error) {
	return _KFIVE.Contract.Paused(&_KFIVE.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_KFIVE *KFIVECaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _KFIVE.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_KFIVE *KFIVESession) Symbol() (string, error) {
	return _KFIVE.Contract.Symbol(&_KFIVE.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_KFIVE *KFIVECallerSession) Symbol() (string, error) {
	return _KFIVE.Contract.Symbol(&_KFIVE.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_KFIVE *KFIVECaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _KFIVE.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_KFIVE *KFIVESession) TotalSupply() (*big.Int, error) {
	return _KFIVE.Contract.TotalSupply(&_KFIVE.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_KFIVE *KFIVECallerSession) TotalSupply() (*big.Int, error) {
	return _KFIVE.Contract.TotalSupply(&_KFIVE.CallOpts)
}

// AddAdmin is a paid mutator transaction binding the contract method 0x46f2ed2e.
//
// Solidity: function addAdmin(address a, uint256 maxIssuingTokenPerTime, uint256 maxTotalIssuingToken) returns()
func (_KFIVE *KFIVETransactor) AddAdmin(opts *bind.TransactOpts, a common.Address, maxIssuingTokenPerTime *big.Int, maxTotalIssuingToken *big.Int) (*types.Transaction, error) {
	return _KFIVE.contract.Transact(opts, "addAdmin", a, maxIssuingTokenPerTime, maxTotalIssuingToken)
}

// AddAdmin is a paid mutator transaction binding the contract method 0x46f2ed2e.
//
// Solidity: function addAdmin(address a, uint256 maxIssuingTokenPerTime, uint256 maxTotalIssuingToken) returns()
func (_KFIVE *KFIVESession) AddAdmin(a common.Address, maxIssuingTokenPerTime *big.Int, maxTotalIssuingToken *big.Int) (*types.Transaction, error) {
	return _KFIVE.Contract.AddAdmin(&_KFIVE.TransactOpts, a, maxIssuingTokenPerTime, maxTotalIssuingToken)
}

// AddAdmin is a paid mutator transaction binding the contract method 0x46f2ed2e.
//
// Solidity: function addAdmin(address a, uint256 maxIssuingTokenPerTime, uint256 maxTotalIssuingToken) returns()
func (_KFIVE *KFIVETransactorSession) AddAdmin(a common.Address, maxIssuingTokenPerTime *big.Int, maxTotalIssuingToken *big.Int) (*types.Transaction, error) {
	return _KFIVE.Contract.AddAdmin(&_KFIVE.TransactOpts, a, maxIssuingTokenPerTime, maxTotalIssuingToken)
}

// AddBlackList is a paid mutator transaction binding the contract method 0x0ecb93c0.
//
// Solidity: function addBlackList(address _evilUser) returns()
func (_KFIVE *KFIVETransactor) AddBlackList(opts *bind.TransactOpts, _evilUser common.Address) (*types.Transaction, error) {
	return _KFIVE.contract.Transact(opts, "addBlackList", _evilUser)
}

// AddBlackList is a paid mutator transaction binding the contract method 0x0ecb93c0.
//
// Solidity: function addBlackList(address _evilUser) returns()
func (_KFIVE *KFIVESession) AddBlackList(_evilUser common.Address) (*types.Transaction, error) {
	return _KFIVE.Contract.AddBlackList(&_KFIVE.TransactOpts, _evilUser)
}

// AddBlackList is a paid mutator transaction binding the contract method 0x0ecb93c0.
//
// Solidity: function addBlackList(address _evilUser) returns()
func (_KFIVE *KFIVETransactorSession) AddBlackList(_evilUser common.Address) (*types.Transaction, error) {
	return _KFIVE.Contract.AddBlackList(&_KFIVE.TransactOpts, _evilUser)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_KFIVE *KFIVETransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _KFIVE.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_KFIVE *KFIVESession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _KFIVE.Contract.Approve(&_KFIVE.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_KFIVE *KFIVETransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _KFIVE.Contract.Approve(&_KFIVE.TransactOpts, spender, amount)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_KFIVE *KFIVETransactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _KFIVE.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_KFIVE *KFIVESession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _KFIVE.Contract.DecreaseAllowance(&_KFIVE.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_KFIVE *KFIVETransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _KFIVE.Contract.DecreaseAllowance(&_KFIVE.TransactOpts, spender, subtractedValue)
}

// DestroyBlackFunds is a paid mutator transaction binding the contract method 0xf3bdc228.
//
// Solidity: function destroyBlackFunds(address _blackListedUser) returns()
func (_KFIVE *KFIVETransactor) DestroyBlackFunds(opts *bind.TransactOpts, _blackListedUser common.Address) (*types.Transaction, error) {
	return _KFIVE.contract.Transact(opts, "destroyBlackFunds", _blackListedUser)
}

// DestroyBlackFunds is a paid mutator transaction binding the contract method 0xf3bdc228.
//
// Solidity: function destroyBlackFunds(address _blackListedUser) returns()
func (_KFIVE *KFIVESession) DestroyBlackFunds(_blackListedUser common.Address) (*types.Transaction, error) {
	return _KFIVE.Contract.DestroyBlackFunds(&_KFIVE.TransactOpts, _blackListedUser)
}

// DestroyBlackFunds is a paid mutator transaction binding the contract method 0xf3bdc228.
//
// Solidity: function destroyBlackFunds(address _blackListedUser) returns()
func (_KFIVE *KFIVETransactorSession) DestroyBlackFunds(_blackListedUser common.Address) (*types.Transaction, error) {
	return _KFIVE.Contract.DestroyBlackFunds(&_KFIVE.TransactOpts, _blackListedUser)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_KFIVE *KFIVETransactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _KFIVE.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_KFIVE *KFIVESession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _KFIVE.Contract.IncreaseAllowance(&_KFIVE.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_KFIVE *KFIVETransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _KFIVE.Contract.IncreaseAllowance(&_KFIVE.TransactOpts, spender, addedValue)
}

// Issue is a paid mutator transaction binding the contract method 0xea3f068d.
//
// Solidity: function issue(address issuer, uint256 value, bytes32 offchain) returns()
func (_KFIVE *KFIVETransactor) Issue(opts *bind.TransactOpts, issuer common.Address, value *big.Int, offchain [32]byte) (*types.Transaction, error) {
	return _KFIVE.contract.Transact(opts, "issue", issuer, value, offchain)
}

// Issue is a paid mutator transaction binding the contract method 0xea3f068d.
//
// Solidity: function issue(address issuer, uint256 value, bytes32 offchain) returns()
func (_KFIVE *KFIVESession) Issue(issuer common.Address, value *big.Int, offchain [32]byte) (*types.Transaction, error) {
	return _KFIVE.Contract.Issue(&_KFIVE.TransactOpts, issuer, value, offchain)
}

// Issue is a paid mutator transaction binding the contract method 0xea3f068d.
//
// Solidity: function issue(address issuer, uint256 value, bytes32 offchain) returns()
func (_KFIVE *KFIVETransactorSession) Issue(issuer common.Address, value *big.Int, offchain [32]byte) (*types.Transaction, error) {
	return _KFIVE.Contract.Issue(&_KFIVE.TransactOpts, issuer, value, offchain)
}

// IssueByAdmin is a paid mutator transaction binding the contract method 0xd10601ec.
//
// Solidity: function issueByAdmin(address issuer, uint256 value, bytes32 offchain) returns()
func (_KFIVE *KFIVETransactor) IssueByAdmin(opts *bind.TransactOpts, issuer common.Address, value *big.Int, offchain [32]byte) (*types.Transaction, error) {
	return _KFIVE.contract.Transact(opts, "issueByAdmin", issuer, value, offchain)
}

// IssueByAdmin is a paid mutator transaction binding the contract method 0xd10601ec.
//
// Solidity: function issueByAdmin(address issuer, uint256 value, bytes32 offchain) returns()
func (_KFIVE *KFIVESession) IssueByAdmin(issuer common.Address, value *big.Int, offchain [32]byte) (*types.Transaction, error) {
	return _KFIVE.Contract.IssueByAdmin(&_KFIVE.TransactOpts, issuer, value, offchain)
}

// IssueByAdmin is a paid mutator transaction binding the contract method 0xd10601ec.
//
// Solidity: function issueByAdmin(address issuer, uint256 value, bytes32 offchain) returns()
func (_KFIVE *KFIVETransactorSession) IssueByAdmin(issuer common.Address, value *big.Int, offchain [32]byte) (*types.Transaction, error) {
	return _KFIVE.Contract.IssueByAdmin(&_KFIVE.TransactOpts, issuer, value, offchain)
}

// Mint is a paid mutator transaction binding the contract method 0xa0712d68.
//
// Solidity: function mint(uint256 amount) returns(bool)
func (_KFIVE *KFIVETransactor) Mint(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _KFIVE.contract.Transact(opts, "mint", amount)
}

// Mint is a paid mutator transaction binding the contract method 0xa0712d68.
//
// Solidity: function mint(uint256 amount) returns(bool)
func (_KFIVE *KFIVESession) Mint(amount *big.Int) (*types.Transaction, error) {
	return _KFIVE.Contract.Mint(&_KFIVE.TransactOpts, amount)
}

// Mint is a paid mutator transaction binding the contract method 0xa0712d68.
//
// Solidity: function mint(uint256 amount) returns(bool)
func (_KFIVE *KFIVETransactorSession) Mint(amount *big.Int) (*types.Transaction, error) {
	return _KFIVE.Contract.Mint(&_KFIVE.TransactOpts, amount)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_KFIVE *KFIVETransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _KFIVE.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_KFIVE *KFIVESession) Pause() (*types.Transaction, error) {
	return _KFIVE.Contract.Pause(&_KFIVE.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_KFIVE *KFIVETransactorSession) Pause() (*types.Transaction, error) {
	return _KFIVE.Contract.Pause(&_KFIVE.TransactOpts)
}

// Redeem is a paid mutator transaction binding the contract method 0x992c3e4b.
//
// Solidity: function redeem(address redeemer, uint256 value, bytes32 offchain) returns()
func (_KFIVE *KFIVETransactor) Redeem(opts *bind.TransactOpts, redeemer common.Address, value *big.Int, offchain [32]byte) (*types.Transaction, error) {
	return _KFIVE.contract.Transact(opts, "redeem", redeemer, value, offchain)
}

// Redeem is a paid mutator transaction binding the contract method 0x992c3e4b.
//
// Solidity: function redeem(address redeemer, uint256 value, bytes32 offchain) returns()
func (_KFIVE *KFIVESession) Redeem(redeemer common.Address, value *big.Int, offchain [32]byte) (*types.Transaction, error) {
	return _KFIVE.Contract.Redeem(&_KFIVE.TransactOpts, redeemer, value, offchain)
}

// Redeem is a paid mutator transaction binding the contract method 0x992c3e4b.
//
// Solidity: function redeem(address redeemer, uint256 value, bytes32 offchain) returns()
func (_KFIVE *KFIVETransactorSession) Redeem(redeemer common.Address, value *big.Int, offchain [32]byte) (*types.Transaction, error) {
	return _KFIVE.Contract.Redeem(&_KFIVE.TransactOpts, redeemer, value, offchain)
}

// RemoveAdmin is a paid mutator transaction binding the contract method 0x1785f53c.
//
// Solidity: function removeAdmin(address a) returns()
func (_KFIVE *KFIVETransactor) RemoveAdmin(opts *bind.TransactOpts, a common.Address) (*types.Transaction, error) {
	return _KFIVE.contract.Transact(opts, "removeAdmin", a)
}

// RemoveAdmin is a paid mutator transaction binding the contract method 0x1785f53c.
//
// Solidity: function removeAdmin(address a) returns()
func (_KFIVE *KFIVESession) RemoveAdmin(a common.Address) (*types.Transaction, error) {
	return _KFIVE.Contract.RemoveAdmin(&_KFIVE.TransactOpts, a)
}

// RemoveAdmin is a paid mutator transaction binding the contract method 0x1785f53c.
//
// Solidity: function removeAdmin(address a) returns()
func (_KFIVE *KFIVETransactorSession) RemoveAdmin(a common.Address) (*types.Transaction, error) {
	return _KFIVE.Contract.RemoveAdmin(&_KFIVE.TransactOpts, a)
}

// RemoveBlackList is a paid mutator transaction binding the contract method 0xe4997dc5.
//
// Solidity: function removeBlackList(address _clearedUser) returns()
func (_KFIVE *KFIVETransactor) RemoveBlackList(opts *bind.TransactOpts, _clearedUser common.Address) (*types.Transaction, error) {
	return _KFIVE.contract.Transact(opts, "removeBlackList", _clearedUser)
}

// RemoveBlackList is a paid mutator transaction binding the contract method 0xe4997dc5.
//
// Solidity: function removeBlackList(address _clearedUser) returns()
func (_KFIVE *KFIVESession) RemoveBlackList(_clearedUser common.Address) (*types.Transaction, error) {
	return _KFIVE.Contract.RemoveBlackList(&_KFIVE.TransactOpts, _clearedUser)
}

// RemoveBlackList is a paid mutator transaction binding the contract method 0xe4997dc5.
//
// Solidity: function removeBlackList(address _clearedUser) returns()
func (_KFIVE *KFIVETransactorSession) RemoveBlackList(_clearedUser common.Address) (*types.Transaction, error) {
	return _KFIVE.Contract.RemoveBlackList(&_KFIVE.TransactOpts, _clearedUser)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_KFIVE *KFIVETransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _KFIVE.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_KFIVE *KFIVESession) RenounceOwnership() (*types.Transaction, error) {
	return _KFIVE.Contract.RenounceOwnership(&_KFIVE.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_KFIVE *KFIVETransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _KFIVE.Contract.RenounceOwnership(&_KFIVE.TransactOpts)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns(bool)
func (_KFIVE *KFIVETransactor) Transfer(opts *bind.TransactOpts, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _KFIVE.contract.Transact(opts, "transfer", _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns(bool)
func (_KFIVE *KFIVESession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _KFIVE.Contract.Transfer(&_KFIVE.TransactOpts, _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns(bool)
func (_KFIVE *KFIVETransactorSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _KFIVE.Contract.Transfer(&_KFIVE.TransactOpts, _to, _value)
}

// TransferByAdmin is a paid mutator transaction binding the contract method 0xf9ba884f.
//
// Solidity: function transferByAdmin(address from, address to, uint256 value, bytes32 offchain) returns()
func (_KFIVE *KFIVETransactor) TransferByAdmin(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int, offchain [32]byte) (*types.Transaction, error) {
	return _KFIVE.contract.Transact(opts, "transferByAdmin", from, to, value, offchain)
}

// TransferByAdmin is a paid mutator transaction binding the contract method 0xf9ba884f.
//
// Solidity: function transferByAdmin(address from, address to, uint256 value, bytes32 offchain) returns()
func (_KFIVE *KFIVESession) TransferByAdmin(from common.Address, to common.Address, value *big.Int, offchain [32]byte) (*types.Transaction, error) {
	return _KFIVE.Contract.TransferByAdmin(&_KFIVE.TransactOpts, from, to, value, offchain)
}

// TransferByAdmin is a paid mutator transaction binding the contract method 0xf9ba884f.
//
// Solidity: function transferByAdmin(address from, address to, uint256 value, bytes32 offchain) returns()
func (_KFIVE *KFIVETransactorSession) TransferByAdmin(from common.Address, to common.Address, value *big.Int, offchain [32]byte) (*types.Transaction, error) {
	return _KFIVE.Contract.TransferByAdmin(&_KFIVE.TransactOpts, from, to, value, offchain)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _value) returns(bool)
func (_KFIVE *KFIVETransactor) TransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _KFIVE.contract.Transact(opts, "transferFrom", _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _value) returns(bool)
func (_KFIVE *KFIVESession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _KFIVE.Contract.TransferFrom(&_KFIVE.TransactOpts, _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _value) returns(bool)
func (_KFIVE *KFIVETransactorSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _KFIVE.Contract.TransferFrom(&_KFIVE.TransactOpts, _from, _to, _value)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_KFIVE *KFIVETransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _KFIVE.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_KFIVE *KFIVESession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _KFIVE.Contract.TransferOwnership(&_KFIVE.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_KFIVE *KFIVETransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _KFIVE.Contract.TransferOwnership(&_KFIVE.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_KFIVE *KFIVETransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _KFIVE.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_KFIVE *KFIVESession) Unpause() (*types.Transaction, error) {
	return _KFIVE.Contract.Unpause(&_KFIVE.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_KFIVE *KFIVETransactorSession) Unpause() (*types.Transaction, error) {
	return _KFIVE.Contract.Unpause(&_KFIVE.TransactOpts)
}

// KFIVEAddedBlackListIterator is returned from FilterAddedBlackList and is used to iterate over the raw logs and unpacked data for AddedBlackList events raised by the KFIVE contract.
type KFIVEAddedBlackListIterator struct {
	Event *KFIVEAddedBlackList // Event containing the contract specifics and raw log

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
func (it *KFIVEAddedBlackListIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KFIVEAddedBlackList)
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
		it.Event = new(KFIVEAddedBlackList)
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
func (it *KFIVEAddedBlackListIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KFIVEAddedBlackListIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KFIVEAddedBlackList represents a AddedBlackList event raised by the KFIVE contract.
type KFIVEAddedBlackList struct {
	User common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterAddedBlackList is a free log retrieval operation binding the contract event 0x42e160154868087d6bfdc0ca23d96a1c1cfa32f1b72ba9ba27b69b98a0d819dc.
//
// Solidity: event AddedBlackList(address _user)
func (_KFIVE *KFIVEFilterer) FilterAddedBlackList(opts *bind.FilterOpts) (*KFIVEAddedBlackListIterator, error) {

	logs, sub, err := _KFIVE.contract.FilterLogs(opts, "AddedBlackList")
	if err != nil {
		return nil, err
	}
	return &KFIVEAddedBlackListIterator{contract: _KFIVE.contract, event: "AddedBlackList", logs: logs, sub: sub}, nil
}

// WatchAddedBlackList is a free log subscription operation binding the contract event 0x42e160154868087d6bfdc0ca23d96a1c1cfa32f1b72ba9ba27b69b98a0d819dc.
//
// Solidity: event AddedBlackList(address _user)
func (_KFIVE *KFIVEFilterer) WatchAddedBlackList(opts *bind.WatchOpts, sink chan<- *KFIVEAddedBlackList) (event.Subscription, error) {

	logs, sub, err := _KFIVE.contract.WatchLogs(opts, "AddedBlackList")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KFIVEAddedBlackList)
				if err := _KFIVE.contract.UnpackLog(event, "AddedBlackList", log); err != nil {
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
func (_KFIVE *KFIVEFilterer) ParseAddedBlackList(log types.Log) (*KFIVEAddedBlackList, error) {
	event := new(KFIVEAddedBlackList)
	if err := _KFIVE.contract.UnpackLog(event, "AddedBlackList", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// KFIVEApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the KFIVE contract.
type KFIVEApprovalIterator struct {
	Event *KFIVEApproval // Event containing the contract specifics and raw log

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
func (it *KFIVEApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KFIVEApproval)
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
		it.Event = new(KFIVEApproval)
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
func (it *KFIVEApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KFIVEApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KFIVEApproval represents a Approval event raised by the KFIVE contract.
type KFIVEApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_KFIVE *KFIVEFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*KFIVEApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _KFIVE.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &KFIVEApprovalIterator{contract: _KFIVE.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_KFIVE *KFIVEFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *KFIVEApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _KFIVE.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KFIVEApproval)
				if err := _KFIVE.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_KFIVE *KFIVEFilterer) ParseApproval(log types.Log) (*KFIVEApproval, error) {
	event := new(KFIVEApproval)
	if err := _KFIVE.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// KFIVEDestroyedBlackFundsIterator is returned from FilterDestroyedBlackFunds and is used to iterate over the raw logs and unpacked data for DestroyedBlackFunds events raised by the KFIVE contract.
type KFIVEDestroyedBlackFundsIterator struct {
	Event *KFIVEDestroyedBlackFunds // Event containing the contract specifics and raw log

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
func (it *KFIVEDestroyedBlackFundsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KFIVEDestroyedBlackFunds)
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
		it.Event = new(KFIVEDestroyedBlackFunds)
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
func (it *KFIVEDestroyedBlackFundsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KFIVEDestroyedBlackFundsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KFIVEDestroyedBlackFunds represents a DestroyedBlackFunds event raised by the KFIVE contract.
type KFIVEDestroyedBlackFunds struct {
	BlackListedUser common.Address
	Balance         *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterDestroyedBlackFunds is a free log retrieval operation binding the contract event 0x61e6e66b0d6339b2980aecc6ccc0039736791f0ccde9ed512e789a7fbdd698c6.
//
// Solidity: event DestroyedBlackFunds(address _blackListedUser, uint256 _balance)
func (_KFIVE *KFIVEFilterer) FilterDestroyedBlackFunds(opts *bind.FilterOpts) (*KFIVEDestroyedBlackFundsIterator, error) {

	logs, sub, err := _KFIVE.contract.FilterLogs(opts, "DestroyedBlackFunds")
	if err != nil {
		return nil, err
	}
	return &KFIVEDestroyedBlackFundsIterator{contract: _KFIVE.contract, event: "DestroyedBlackFunds", logs: logs, sub: sub}, nil
}

// WatchDestroyedBlackFunds is a free log subscription operation binding the contract event 0x61e6e66b0d6339b2980aecc6ccc0039736791f0ccde9ed512e789a7fbdd698c6.
//
// Solidity: event DestroyedBlackFunds(address _blackListedUser, uint256 _balance)
func (_KFIVE *KFIVEFilterer) WatchDestroyedBlackFunds(opts *bind.WatchOpts, sink chan<- *KFIVEDestroyedBlackFunds) (event.Subscription, error) {

	logs, sub, err := _KFIVE.contract.WatchLogs(opts, "DestroyedBlackFunds")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KFIVEDestroyedBlackFunds)
				if err := _KFIVE.contract.UnpackLog(event, "DestroyedBlackFunds", log); err != nil {
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
func (_KFIVE *KFIVEFilterer) ParseDestroyedBlackFunds(log types.Log) (*KFIVEDestroyedBlackFunds, error) {
	event := new(KFIVEDestroyedBlackFunds)
	if err := _KFIVE.contract.UnpackLog(event, "DestroyedBlackFunds", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// KFIVEOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the KFIVE contract.
type KFIVEOwnershipTransferredIterator struct {
	Event *KFIVEOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *KFIVEOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KFIVEOwnershipTransferred)
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
		it.Event = new(KFIVEOwnershipTransferred)
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
func (it *KFIVEOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KFIVEOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KFIVEOwnershipTransferred represents a OwnershipTransferred event raised by the KFIVE contract.
type KFIVEOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_KFIVE *KFIVEFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*KFIVEOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _KFIVE.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &KFIVEOwnershipTransferredIterator{contract: _KFIVE.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_KFIVE *KFIVEFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *KFIVEOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _KFIVE.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KFIVEOwnershipTransferred)
				if err := _KFIVE.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_KFIVE *KFIVEFilterer) ParseOwnershipTransferred(log types.Log) (*KFIVEOwnershipTransferred, error) {
	event := new(KFIVEOwnershipTransferred)
	if err := _KFIVE.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// KFIVEPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the KFIVE contract.
type KFIVEPausedIterator struct {
	Event *KFIVEPaused // Event containing the contract specifics and raw log

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
func (it *KFIVEPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KFIVEPaused)
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
		it.Event = new(KFIVEPaused)
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
func (it *KFIVEPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KFIVEPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KFIVEPaused represents a Paused event raised by the KFIVE contract.
type KFIVEPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_KFIVE *KFIVEFilterer) FilterPaused(opts *bind.FilterOpts) (*KFIVEPausedIterator, error) {

	logs, sub, err := _KFIVE.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &KFIVEPausedIterator{contract: _KFIVE.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_KFIVE *KFIVEFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *KFIVEPaused) (event.Subscription, error) {

	logs, sub, err := _KFIVE.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KFIVEPaused)
				if err := _KFIVE.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_KFIVE *KFIVEFilterer) ParsePaused(log types.Log) (*KFIVEPaused, error) {
	event := new(KFIVEPaused)
	if err := _KFIVE.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// KFIVERemovedBlackListIterator is returned from FilterRemovedBlackList and is used to iterate over the raw logs and unpacked data for RemovedBlackList events raised by the KFIVE contract.
type KFIVERemovedBlackListIterator struct {
	Event *KFIVERemovedBlackList // Event containing the contract specifics and raw log

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
func (it *KFIVERemovedBlackListIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KFIVERemovedBlackList)
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
		it.Event = new(KFIVERemovedBlackList)
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
func (it *KFIVERemovedBlackListIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KFIVERemovedBlackListIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KFIVERemovedBlackList represents a RemovedBlackList event raised by the KFIVE contract.
type KFIVERemovedBlackList struct {
	User common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterRemovedBlackList is a free log retrieval operation binding the contract event 0xd7e9ec6e6ecd65492dce6bf513cd6867560d49544421d0783ddf06e76c24470c.
//
// Solidity: event RemovedBlackList(address _user)
func (_KFIVE *KFIVEFilterer) FilterRemovedBlackList(opts *bind.FilterOpts) (*KFIVERemovedBlackListIterator, error) {

	logs, sub, err := _KFIVE.contract.FilterLogs(opts, "RemovedBlackList")
	if err != nil {
		return nil, err
	}
	return &KFIVERemovedBlackListIterator{contract: _KFIVE.contract, event: "RemovedBlackList", logs: logs, sub: sub}, nil
}

// WatchRemovedBlackList is a free log subscription operation binding the contract event 0xd7e9ec6e6ecd65492dce6bf513cd6867560d49544421d0783ddf06e76c24470c.
//
// Solidity: event RemovedBlackList(address _user)
func (_KFIVE *KFIVEFilterer) WatchRemovedBlackList(opts *bind.WatchOpts, sink chan<- *KFIVERemovedBlackList) (event.Subscription, error) {

	logs, sub, err := _KFIVE.contract.WatchLogs(opts, "RemovedBlackList")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KFIVERemovedBlackList)
				if err := _KFIVE.contract.UnpackLog(event, "RemovedBlackList", log); err != nil {
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
func (_KFIVE *KFIVEFilterer) ParseRemovedBlackList(log types.Log) (*KFIVERemovedBlackList, error) {
	event := new(KFIVERemovedBlackList)
	if err := _KFIVE.contract.UnpackLog(event, "RemovedBlackList", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// KFIVETransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the KFIVE contract.
type KFIVETransferIterator struct {
	Event *KFIVETransfer // Event containing the contract specifics and raw log

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
func (it *KFIVETransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KFIVETransfer)
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
		it.Event = new(KFIVETransfer)
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
func (it *KFIVETransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KFIVETransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KFIVETransfer represents a Transfer event raised by the KFIVE contract.
type KFIVETransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_KFIVE *KFIVEFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*KFIVETransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _KFIVE.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &KFIVETransferIterator{contract: _KFIVE.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_KFIVE *KFIVEFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *KFIVETransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _KFIVE.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KFIVETransfer)
				if err := _KFIVE.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_KFIVE *KFIVEFilterer) ParseTransfer(log types.Log) (*KFIVETransfer, error) {
	event := new(KFIVETransfer)
	if err := _KFIVE.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// KFIVEUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the KFIVE contract.
type KFIVEUnpausedIterator struct {
	Event *KFIVEUnpaused // Event containing the contract specifics and raw log

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
func (it *KFIVEUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KFIVEUnpaused)
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
		it.Event = new(KFIVEUnpaused)
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
func (it *KFIVEUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KFIVEUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KFIVEUnpaused represents a Unpaused event raised by the KFIVE contract.
type KFIVEUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_KFIVE *KFIVEFilterer) FilterUnpaused(opts *bind.FilterOpts) (*KFIVEUnpausedIterator, error) {

	logs, sub, err := _KFIVE.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &KFIVEUnpausedIterator{contract: _KFIVE.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_KFIVE *KFIVEFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *KFIVEUnpaused) (event.Subscription, error) {

	logs, sub, err := _KFIVE.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KFIVEUnpaused)
				if err := _KFIVE.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
func (_KFIVE *KFIVEFilterer) ParseUnpaused(log types.Log) (*KFIVEUnpaused, error) {
	event := new(KFIVEUnpaused)
	if err := _KFIVE.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// KFIVEIssueIterator is returned from FilterIssue and is used to iterate over the raw logs and unpacked data for Issue events raised by the KFIVE contract.
type KFIVEIssueIterator struct {
	Event *KFIVEIssue // Event containing the contract specifics and raw log

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
func (it *KFIVEIssueIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KFIVEIssue)
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
		it.Event = new(KFIVEIssue)
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
func (it *KFIVEIssueIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KFIVEIssueIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KFIVEIssue represents a Issue event raised by the KFIVE contract.
type KFIVEIssue struct {
	Offchain [32]byte
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterIssue is a free log retrieval operation binding the contract event 0x182a88d3fbfcb38ffaddf26e6d63ecd93b4348141de1eb744c43c4fbb1ccfb22.
//
// Solidity: event __issue(bytes32 offchain)
func (_KFIVE *KFIVEFilterer) FilterIssue(opts *bind.FilterOpts) (*KFIVEIssueIterator, error) {

	logs, sub, err := _KFIVE.contract.FilterLogs(opts, "__issue")
	if err != nil {
		return nil, err
	}
	return &KFIVEIssueIterator{contract: _KFIVE.contract, event: "__issue", logs: logs, sub: sub}, nil
}

// WatchIssue is a free log subscription operation binding the contract event 0x182a88d3fbfcb38ffaddf26e6d63ecd93b4348141de1eb744c43c4fbb1ccfb22.
//
// Solidity: event __issue(bytes32 offchain)
func (_KFIVE *KFIVEFilterer) WatchIssue(opts *bind.WatchOpts, sink chan<- *KFIVEIssue) (event.Subscription, error) {

	logs, sub, err := _KFIVE.contract.WatchLogs(opts, "__issue")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KFIVEIssue)
				if err := _KFIVE.contract.UnpackLog(event, "__issue", log); err != nil {
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
func (_KFIVE *KFIVEFilterer) ParseIssue(log types.Log) (*KFIVEIssue, error) {
	event := new(KFIVEIssue)
	if err := _KFIVE.contract.UnpackLog(event, "__issue", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// KFIVERedeemIterator is returned from FilterRedeem and is used to iterate over the raw logs and unpacked data for Redeem events raised by the KFIVE contract.
type KFIVERedeemIterator struct {
	Event *KFIVERedeem // Event containing the contract specifics and raw log

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
func (it *KFIVERedeemIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KFIVERedeem)
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
		it.Event = new(KFIVERedeem)
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
func (it *KFIVERedeemIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KFIVERedeemIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KFIVERedeem represents a Redeem event raised by the KFIVE contract.
type KFIVERedeem struct {
	Offchain [32]byte
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterRedeem is a free log retrieval operation binding the contract event 0xb0de879351469d2741406aafc9ba1f44eb957cf44ee3391e59a7a9097050c927.
//
// Solidity: event __redeem(bytes32 offchain)
func (_KFIVE *KFIVEFilterer) FilterRedeem(opts *bind.FilterOpts) (*KFIVERedeemIterator, error) {

	logs, sub, err := _KFIVE.contract.FilterLogs(opts, "__redeem")
	if err != nil {
		return nil, err
	}
	return &KFIVERedeemIterator{contract: _KFIVE.contract, event: "__redeem", logs: logs, sub: sub}, nil
}

// WatchRedeem is a free log subscription operation binding the contract event 0xb0de879351469d2741406aafc9ba1f44eb957cf44ee3391e59a7a9097050c927.
//
// Solidity: event __redeem(bytes32 offchain)
func (_KFIVE *KFIVEFilterer) WatchRedeem(opts *bind.WatchOpts, sink chan<- *KFIVERedeem) (event.Subscription, error) {

	logs, sub, err := _KFIVE.contract.WatchLogs(opts, "__redeem")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KFIVERedeem)
				if err := _KFIVE.contract.UnpackLog(event, "__redeem", log); err != nil {
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
func (_KFIVE *KFIVEFilterer) ParseRedeem(log types.Log) (*KFIVERedeem, error) {
	event := new(KFIVERedeem)
	if err := _KFIVE.contract.UnpackLog(event, "__redeem", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// KFIVETransferByAdminIterator is returned from FilterTransferByAdmin and is used to iterate over the raw logs and unpacked data for TransferByAdmin events raised by the KFIVE contract.
type KFIVETransferByAdminIterator struct {
	Event *KFIVETransferByAdmin // Event containing the contract specifics and raw log

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
func (it *KFIVETransferByAdminIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KFIVETransferByAdmin)
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
		it.Event = new(KFIVETransferByAdmin)
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
func (it *KFIVETransferByAdminIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KFIVETransferByAdminIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KFIVETransferByAdmin represents a TransferByAdmin event raised by the KFIVE contract.
type KFIVETransferByAdmin struct {
	Offchain [32]byte
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterTransferByAdmin is a free log retrieval operation binding the contract event 0x3b36ee6b35325f38e95938557be92853c842b7a9a19fd7ac4931a6d24db52682.
//
// Solidity: event __transferByAdmin(bytes32 offchain)
func (_KFIVE *KFIVEFilterer) FilterTransferByAdmin(opts *bind.FilterOpts) (*KFIVETransferByAdminIterator, error) {

	logs, sub, err := _KFIVE.contract.FilterLogs(opts, "__transferByAdmin")
	if err != nil {
		return nil, err
	}
	return &KFIVETransferByAdminIterator{contract: _KFIVE.contract, event: "__transferByAdmin", logs: logs, sub: sub}, nil
}

// WatchTransferByAdmin is a free log subscription operation binding the contract event 0x3b36ee6b35325f38e95938557be92853c842b7a9a19fd7ac4931a6d24db52682.
//
// Solidity: event __transferByAdmin(bytes32 offchain)
func (_KFIVE *KFIVEFilterer) WatchTransferByAdmin(opts *bind.WatchOpts, sink chan<- *KFIVETransferByAdmin) (event.Subscription, error) {

	logs, sub, err := _KFIVE.contract.WatchLogs(opts, "__transferByAdmin")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KFIVETransferByAdmin)
				if err := _KFIVE.contract.UnpackLog(event, "__transferByAdmin", log); err != nil {
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
func (_KFIVE *KFIVEFilterer) ParseTransferByAdmin(log types.Log) (*KFIVETransferByAdmin, error) {
	event := new(KFIVETransferByAdmin)
	if err := _KFIVE.contract.UnpackLog(event, "__transferByAdmin", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
