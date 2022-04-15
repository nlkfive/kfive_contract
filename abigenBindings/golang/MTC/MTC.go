// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package MTC

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

// MTCMetaData contains all meta data concerning the MTC contract.
var MTCMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_kfiveAddress\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"AddedBlackList\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_blackListedUser\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_balance\",\"type\":\"uint256\"}],\"name\":\"DestroyedBlackFunds\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"ExchangeFromKFIVE\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"kfiveAddress\",\"type\":\"address\"}],\"name\":\"KfiveAddressChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldReceiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newReceiver\",\"type\":\"address\"}],\"name\":\"KfiveReceiverTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"RemovedBlackList\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"__issue\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"__redeem\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"__transferByAdmin\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"KFIVE2MTC\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"a\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"maxIssuingTokenPerTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxTotalIssuingToken\",\"type\":\"uint256\"}],\"name\":\"addAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_evilUser\",\"type\":\"address\"}],\"name\":\"addBlackList\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"admin\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"status\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"maxIssuingTokenPerTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxTotalIssuingToken\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"remainingIssuingToken\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_maker\",\"type\":\"address\"}],\"name\":\"getBlackListStatus\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"isBlackListed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"kfiveToken\",\"outputs\":[{\"internalType\":\"contractIBEP20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"a\",\"type\":\"address\"}],\"name\":\"removeAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_clearedUser\",\"type\":\"address\"}],\"name\":\"removeBlackList\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transferByAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"issuer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"issue\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"issuer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"issueByAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_blackListedUser\",\"type\":\"address\"}],\"name\":\"destroyBlackFunds\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"redeemer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"redeem\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newReceiver\",\"type\":\"address\"}],\"name\":\"transferKfiveReceiver\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newAddress\",\"type\":\"address\"}],\"name\":\"updateKfiveAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"exchangeFromKfive\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// MTCABI is the input ABI used to generate the binding from.
// Deprecated: Use MTCMetaData.ABI instead.
var MTCABI = MTCMetaData.ABI

// MTC is an auto generated Go binding around an Ethereum contract.
type MTC struct {
	MTCCaller     // Read-only binding to the contract
	MTCTransactor // Write-only binding to the contract
	MTCFilterer   // Log filterer for contract events
}

// MTCCaller is an auto generated read-only Go binding around an Ethereum contract.
type MTCCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MTCTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MTCTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MTCFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MTCFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MTCSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MTCSession struct {
	Contract     *MTC              // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MTCCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MTCCallerSession struct {
	Contract *MTCCaller    // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// MTCTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MTCTransactorSession struct {
	Contract     *MTCTransactor    // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MTCRaw is an auto generated low-level Go binding around an Ethereum contract.
type MTCRaw struct {
	Contract *MTC // Generic contract binding to access the raw methods on
}

// MTCCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MTCCallerRaw struct {
	Contract *MTCCaller // Generic read-only contract binding to access the raw methods on
}

// MTCTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MTCTransactorRaw struct {
	Contract *MTCTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMTC creates a new instance of MTC, bound to a specific deployed contract.
func NewMTC(address common.Address, backend bind.ContractBackend) (*MTC, error) {
	contract, err := bindMTC(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MTC{MTCCaller: MTCCaller{contract: contract}, MTCTransactor: MTCTransactor{contract: contract}, MTCFilterer: MTCFilterer{contract: contract}}, nil
}

// NewMTCCaller creates a new read-only instance of MTC, bound to a specific deployed contract.
func NewMTCCaller(address common.Address, caller bind.ContractCaller) (*MTCCaller, error) {
	contract, err := bindMTC(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MTCCaller{contract: contract}, nil
}

// NewMTCTransactor creates a new write-only instance of MTC, bound to a specific deployed contract.
func NewMTCTransactor(address common.Address, transactor bind.ContractTransactor) (*MTCTransactor, error) {
	contract, err := bindMTC(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MTCTransactor{contract: contract}, nil
}

// NewMTCFilterer creates a new log filterer instance of MTC, bound to a specific deployed contract.
func NewMTCFilterer(address common.Address, filterer bind.ContractFilterer) (*MTCFilterer, error) {
	contract, err := bindMTC(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MTCFilterer{contract: contract}, nil
}

// bindMTC binds a generic wrapper to an already deployed contract.
func bindMTC(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MTCABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MTC *MTCRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MTC.Contract.MTCCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MTC *MTCRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MTC.Contract.MTCTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MTC *MTCRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MTC.Contract.MTCTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MTC *MTCCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MTC.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MTC *MTCTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MTC.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MTC *MTCTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MTC.Contract.contract.Transact(opts, method, params...)
}

// KFIVE2MTC is a free data retrieval call binding the contract method 0x62e8db23.
//
// Solidity: function KFIVE2MTC() view returns(uint256)
func (_MTC *MTCCaller) KFIVE2MTC(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MTC.contract.Call(opts, &out, "KFIVE2MTC")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// KFIVE2MTC is a free data retrieval call binding the contract method 0x62e8db23.
//
// Solidity: function KFIVE2MTC() view returns(uint256)
func (_MTC *MTCSession) KFIVE2MTC() (*big.Int, error) {
	return _MTC.Contract.KFIVE2MTC(&_MTC.CallOpts)
}

// KFIVE2MTC is a free data retrieval call binding the contract method 0x62e8db23.
//
// Solidity: function KFIVE2MTC() view returns(uint256)
func (_MTC *MTCCallerSession) KFIVE2MTC() (*big.Int, error) {
	return _MTC.Contract.KFIVE2MTC(&_MTC.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0x63a846f8.
//
// Solidity: function admin(address ) view returns(bool status, uint256 maxIssuingTokenPerTime, uint256 maxTotalIssuingToken, uint256 remainingIssuingToken)
func (_MTC *MTCCaller) Admin(opts *bind.CallOpts, arg0 common.Address) (struct {
	Status                 bool
	MaxIssuingTokenPerTime *big.Int
	MaxTotalIssuingToken   *big.Int
	RemainingIssuingToken  *big.Int
}, error) {
	var out []interface{}
	err := _MTC.contract.Call(opts, &out, "admin", arg0)

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
func (_MTC *MTCSession) Admin(arg0 common.Address) (struct {
	Status                 bool
	MaxIssuingTokenPerTime *big.Int
	MaxTotalIssuingToken   *big.Int
	RemainingIssuingToken  *big.Int
}, error) {
	return _MTC.Contract.Admin(&_MTC.CallOpts, arg0)
}

// Admin is a free data retrieval call binding the contract method 0x63a846f8.
//
// Solidity: function admin(address ) view returns(bool status, uint256 maxIssuingTokenPerTime, uint256 maxTotalIssuingToken, uint256 remainingIssuingToken)
func (_MTC *MTCCallerSession) Admin(arg0 common.Address) (struct {
	Status                 bool
	MaxIssuingTokenPerTime *big.Int
	MaxTotalIssuingToken   *big.Int
	RemainingIssuingToken  *big.Int
}, error) {
	return _MTC.Contract.Admin(&_MTC.CallOpts, arg0)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_MTC *MTCCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _MTC.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_MTC *MTCSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _MTC.Contract.Allowance(&_MTC.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_MTC *MTCCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _MTC.Contract.Allowance(&_MTC.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_MTC *MTCCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _MTC.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_MTC *MTCSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _MTC.Contract.BalanceOf(&_MTC.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_MTC *MTCCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _MTC.Contract.BalanceOf(&_MTC.CallOpts, account)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_MTC *MTCCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _MTC.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_MTC *MTCSession) Decimals() (uint8, error) {
	return _MTC.Contract.Decimals(&_MTC.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_MTC *MTCCallerSession) Decimals() (uint8, error) {
	return _MTC.Contract.Decimals(&_MTC.CallOpts)
}

// GetBlackListStatus is a free data retrieval call binding the contract method 0x59bf1abe.
//
// Solidity: function getBlackListStatus(address _maker) view returns(bool)
func (_MTC *MTCCaller) GetBlackListStatus(opts *bind.CallOpts, _maker common.Address) (bool, error) {
	var out []interface{}
	err := _MTC.contract.Call(opts, &out, "getBlackListStatus", _maker)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// GetBlackListStatus is a free data retrieval call binding the contract method 0x59bf1abe.
//
// Solidity: function getBlackListStatus(address _maker) view returns(bool)
func (_MTC *MTCSession) GetBlackListStatus(_maker common.Address) (bool, error) {
	return _MTC.Contract.GetBlackListStatus(&_MTC.CallOpts, _maker)
}

// GetBlackListStatus is a free data retrieval call binding the contract method 0x59bf1abe.
//
// Solidity: function getBlackListStatus(address _maker) view returns(bool)
func (_MTC *MTCCallerSession) GetBlackListStatus(_maker common.Address) (bool, error) {
	return _MTC.Contract.GetBlackListStatus(&_MTC.CallOpts, _maker)
}

// GetOwner is a free data retrieval call binding the contract method 0x893d20e8.
//
// Solidity: function getOwner() view returns(address)
func (_MTC *MTCCaller) GetOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MTC.contract.Call(opts, &out, "getOwner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetOwner is a free data retrieval call binding the contract method 0x893d20e8.
//
// Solidity: function getOwner() view returns(address)
func (_MTC *MTCSession) GetOwner() (common.Address, error) {
	return _MTC.Contract.GetOwner(&_MTC.CallOpts)
}

// GetOwner is a free data retrieval call binding the contract method 0x893d20e8.
//
// Solidity: function getOwner() view returns(address)
func (_MTC *MTCCallerSession) GetOwner() (common.Address, error) {
	return _MTC.Contract.GetOwner(&_MTC.CallOpts)
}

// IsBlackListed is a free data retrieval call binding the contract method 0xe47d6060.
//
// Solidity: function isBlackListed(address ) view returns(bool)
func (_MTC *MTCCaller) IsBlackListed(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _MTC.contract.Call(opts, &out, "isBlackListed", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsBlackListed is a free data retrieval call binding the contract method 0xe47d6060.
//
// Solidity: function isBlackListed(address ) view returns(bool)
func (_MTC *MTCSession) IsBlackListed(arg0 common.Address) (bool, error) {
	return _MTC.Contract.IsBlackListed(&_MTC.CallOpts, arg0)
}

// IsBlackListed is a free data retrieval call binding the contract method 0xe47d6060.
//
// Solidity: function isBlackListed(address ) view returns(bool)
func (_MTC *MTCCallerSession) IsBlackListed(arg0 common.Address) (bool, error) {
	return _MTC.Contract.IsBlackListed(&_MTC.CallOpts, arg0)
}

// KfiveToken is a free data retrieval call binding the contract method 0xe1c558c1.
//
// Solidity: function kfiveToken() view returns(address)
func (_MTC *MTCCaller) KfiveToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MTC.contract.Call(opts, &out, "kfiveToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// KfiveToken is a free data retrieval call binding the contract method 0xe1c558c1.
//
// Solidity: function kfiveToken() view returns(address)
func (_MTC *MTCSession) KfiveToken() (common.Address, error) {
	return _MTC.Contract.KfiveToken(&_MTC.CallOpts)
}

// KfiveToken is a free data retrieval call binding the contract method 0xe1c558c1.
//
// Solidity: function kfiveToken() view returns(address)
func (_MTC *MTCCallerSession) KfiveToken() (common.Address, error) {
	return _MTC.Contract.KfiveToken(&_MTC.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_MTC *MTCCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _MTC.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_MTC *MTCSession) Name() (string, error) {
	return _MTC.Contract.Name(&_MTC.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_MTC *MTCCallerSession) Name() (string, error) {
	return _MTC.Contract.Name(&_MTC.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_MTC *MTCCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MTC.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_MTC *MTCSession) Owner() (common.Address, error) {
	return _MTC.Contract.Owner(&_MTC.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_MTC *MTCCallerSession) Owner() (common.Address, error) {
	return _MTC.Contract.Owner(&_MTC.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_MTC *MTCCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _MTC.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_MTC *MTCSession) Paused() (bool, error) {
	return _MTC.Contract.Paused(&_MTC.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_MTC *MTCCallerSession) Paused() (bool, error) {
	return _MTC.Contract.Paused(&_MTC.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_MTC *MTCCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _MTC.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_MTC *MTCSession) Symbol() (string, error) {
	return _MTC.Contract.Symbol(&_MTC.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_MTC *MTCCallerSession) Symbol() (string, error) {
	return _MTC.Contract.Symbol(&_MTC.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_MTC *MTCCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MTC.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_MTC *MTCSession) TotalSupply() (*big.Int, error) {
	return _MTC.Contract.TotalSupply(&_MTC.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_MTC *MTCCallerSession) TotalSupply() (*big.Int, error) {
	return _MTC.Contract.TotalSupply(&_MTC.CallOpts)
}

// AddAdmin is a paid mutator transaction binding the contract method 0x46f2ed2e.
//
// Solidity: function addAdmin(address a, uint256 maxIssuingTokenPerTime, uint256 maxTotalIssuingToken) returns()
func (_MTC *MTCTransactor) AddAdmin(opts *bind.TransactOpts, a common.Address, maxIssuingTokenPerTime *big.Int, maxTotalIssuingToken *big.Int) (*types.Transaction, error) {
	return _MTC.contract.Transact(opts, "addAdmin", a, maxIssuingTokenPerTime, maxTotalIssuingToken)
}

// AddAdmin is a paid mutator transaction binding the contract method 0x46f2ed2e.
//
// Solidity: function addAdmin(address a, uint256 maxIssuingTokenPerTime, uint256 maxTotalIssuingToken) returns()
func (_MTC *MTCSession) AddAdmin(a common.Address, maxIssuingTokenPerTime *big.Int, maxTotalIssuingToken *big.Int) (*types.Transaction, error) {
	return _MTC.Contract.AddAdmin(&_MTC.TransactOpts, a, maxIssuingTokenPerTime, maxTotalIssuingToken)
}

// AddAdmin is a paid mutator transaction binding the contract method 0x46f2ed2e.
//
// Solidity: function addAdmin(address a, uint256 maxIssuingTokenPerTime, uint256 maxTotalIssuingToken) returns()
func (_MTC *MTCTransactorSession) AddAdmin(a common.Address, maxIssuingTokenPerTime *big.Int, maxTotalIssuingToken *big.Int) (*types.Transaction, error) {
	return _MTC.Contract.AddAdmin(&_MTC.TransactOpts, a, maxIssuingTokenPerTime, maxTotalIssuingToken)
}

// AddBlackList is a paid mutator transaction binding the contract method 0x0ecb93c0.
//
// Solidity: function addBlackList(address _evilUser) returns()
func (_MTC *MTCTransactor) AddBlackList(opts *bind.TransactOpts, _evilUser common.Address) (*types.Transaction, error) {
	return _MTC.contract.Transact(opts, "addBlackList", _evilUser)
}

// AddBlackList is a paid mutator transaction binding the contract method 0x0ecb93c0.
//
// Solidity: function addBlackList(address _evilUser) returns()
func (_MTC *MTCSession) AddBlackList(_evilUser common.Address) (*types.Transaction, error) {
	return _MTC.Contract.AddBlackList(&_MTC.TransactOpts, _evilUser)
}

// AddBlackList is a paid mutator transaction binding the contract method 0x0ecb93c0.
//
// Solidity: function addBlackList(address _evilUser) returns()
func (_MTC *MTCTransactorSession) AddBlackList(_evilUser common.Address) (*types.Transaction, error) {
	return _MTC.Contract.AddBlackList(&_MTC.TransactOpts, _evilUser)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_MTC *MTCTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MTC.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_MTC *MTCSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MTC.Contract.Approve(&_MTC.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_MTC *MTCTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MTC.Contract.Approve(&_MTC.TransactOpts, spender, amount)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_MTC *MTCTransactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _MTC.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_MTC *MTCSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _MTC.Contract.DecreaseAllowance(&_MTC.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_MTC *MTCTransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _MTC.Contract.DecreaseAllowance(&_MTC.TransactOpts, spender, subtractedValue)
}

// DestroyBlackFunds is a paid mutator transaction binding the contract method 0xf3bdc228.
//
// Solidity: function destroyBlackFunds(address _blackListedUser) returns()
func (_MTC *MTCTransactor) DestroyBlackFunds(opts *bind.TransactOpts, _blackListedUser common.Address) (*types.Transaction, error) {
	return _MTC.contract.Transact(opts, "destroyBlackFunds", _blackListedUser)
}

// DestroyBlackFunds is a paid mutator transaction binding the contract method 0xf3bdc228.
//
// Solidity: function destroyBlackFunds(address _blackListedUser) returns()
func (_MTC *MTCSession) DestroyBlackFunds(_blackListedUser common.Address) (*types.Transaction, error) {
	return _MTC.Contract.DestroyBlackFunds(&_MTC.TransactOpts, _blackListedUser)
}

// DestroyBlackFunds is a paid mutator transaction binding the contract method 0xf3bdc228.
//
// Solidity: function destroyBlackFunds(address _blackListedUser) returns()
func (_MTC *MTCTransactorSession) DestroyBlackFunds(_blackListedUser common.Address) (*types.Transaction, error) {
	return _MTC.Contract.DestroyBlackFunds(&_MTC.TransactOpts, _blackListedUser)
}

// ExchangeFromKfive is a paid mutator transaction binding the contract method 0x29fb503b.
//
// Solidity: function exchangeFromKfive(uint256 _value) returns()
func (_MTC *MTCTransactor) ExchangeFromKfive(opts *bind.TransactOpts, _value *big.Int) (*types.Transaction, error) {
	return _MTC.contract.Transact(opts, "exchangeFromKfive", _value)
}

// ExchangeFromKfive is a paid mutator transaction binding the contract method 0x29fb503b.
//
// Solidity: function exchangeFromKfive(uint256 _value) returns()
func (_MTC *MTCSession) ExchangeFromKfive(_value *big.Int) (*types.Transaction, error) {
	return _MTC.Contract.ExchangeFromKfive(&_MTC.TransactOpts, _value)
}

// ExchangeFromKfive is a paid mutator transaction binding the contract method 0x29fb503b.
//
// Solidity: function exchangeFromKfive(uint256 _value) returns()
func (_MTC *MTCTransactorSession) ExchangeFromKfive(_value *big.Int) (*types.Transaction, error) {
	return _MTC.Contract.ExchangeFromKfive(&_MTC.TransactOpts, _value)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_MTC *MTCTransactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _MTC.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_MTC *MTCSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _MTC.Contract.IncreaseAllowance(&_MTC.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_MTC *MTCTransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _MTC.Contract.IncreaseAllowance(&_MTC.TransactOpts, spender, addedValue)
}

// Issue is a paid mutator transaction binding the contract method 0x867904b4.
//
// Solidity: function issue(address issuer, uint256 value) returns()
func (_MTC *MTCTransactor) Issue(opts *bind.TransactOpts, issuer common.Address, value *big.Int) (*types.Transaction, error) {
	return _MTC.contract.Transact(opts, "issue", issuer, value)
}

// Issue is a paid mutator transaction binding the contract method 0x867904b4.
//
// Solidity: function issue(address issuer, uint256 value) returns()
func (_MTC *MTCSession) Issue(issuer common.Address, value *big.Int) (*types.Transaction, error) {
	return _MTC.Contract.Issue(&_MTC.TransactOpts, issuer, value)
}

// Issue is a paid mutator transaction binding the contract method 0x867904b4.
//
// Solidity: function issue(address issuer, uint256 value) returns()
func (_MTC *MTCTransactorSession) Issue(issuer common.Address, value *big.Int) (*types.Transaction, error) {
	return _MTC.Contract.Issue(&_MTC.TransactOpts, issuer, value)
}

// IssueByAdmin is a paid mutator transaction binding the contract method 0xfae2b447.
//
// Solidity: function issueByAdmin(address issuer, uint256 value) returns()
func (_MTC *MTCTransactor) IssueByAdmin(opts *bind.TransactOpts, issuer common.Address, value *big.Int) (*types.Transaction, error) {
	return _MTC.contract.Transact(opts, "issueByAdmin", issuer, value)
}

// IssueByAdmin is a paid mutator transaction binding the contract method 0xfae2b447.
//
// Solidity: function issueByAdmin(address issuer, uint256 value) returns()
func (_MTC *MTCSession) IssueByAdmin(issuer common.Address, value *big.Int) (*types.Transaction, error) {
	return _MTC.Contract.IssueByAdmin(&_MTC.TransactOpts, issuer, value)
}

// IssueByAdmin is a paid mutator transaction binding the contract method 0xfae2b447.
//
// Solidity: function issueByAdmin(address issuer, uint256 value) returns()
func (_MTC *MTCTransactorSession) IssueByAdmin(issuer common.Address, value *big.Int) (*types.Transaction, error) {
	return _MTC.Contract.IssueByAdmin(&_MTC.TransactOpts, issuer, value)
}

// Mint is a paid mutator transaction binding the contract method 0xa0712d68.
//
// Solidity: function mint(uint256 amount) returns(bool)
func (_MTC *MTCTransactor) Mint(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _MTC.contract.Transact(opts, "mint", amount)
}

// Mint is a paid mutator transaction binding the contract method 0xa0712d68.
//
// Solidity: function mint(uint256 amount) returns(bool)
func (_MTC *MTCSession) Mint(amount *big.Int) (*types.Transaction, error) {
	return _MTC.Contract.Mint(&_MTC.TransactOpts, amount)
}

// Mint is a paid mutator transaction binding the contract method 0xa0712d68.
//
// Solidity: function mint(uint256 amount) returns(bool)
func (_MTC *MTCTransactorSession) Mint(amount *big.Int) (*types.Transaction, error) {
	return _MTC.Contract.Mint(&_MTC.TransactOpts, amount)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_MTC *MTCTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MTC.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_MTC *MTCSession) Pause() (*types.Transaction, error) {
	return _MTC.Contract.Pause(&_MTC.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_MTC *MTCTransactorSession) Pause() (*types.Transaction, error) {
	return _MTC.Contract.Pause(&_MTC.TransactOpts)
}

// Redeem is a paid mutator transaction binding the contract method 0x1e9a6950.
//
// Solidity: function redeem(address redeemer, uint256 value) returns()
func (_MTC *MTCTransactor) Redeem(opts *bind.TransactOpts, redeemer common.Address, value *big.Int) (*types.Transaction, error) {
	return _MTC.contract.Transact(opts, "redeem", redeemer, value)
}

// Redeem is a paid mutator transaction binding the contract method 0x1e9a6950.
//
// Solidity: function redeem(address redeemer, uint256 value) returns()
func (_MTC *MTCSession) Redeem(redeemer common.Address, value *big.Int) (*types.Transaction, error) {
	return _MTC.Contract.Redeem(&_MTC.TransactOpts, redeemer, value)
}

// Redeem is a paid mutator transaction binding the contract method 0x1e9a6950.
//
// Solidity: function redeem(address redeemer, uint256 value) returns()
func (_MTC *MTCTransactorSession) Redeem(redeemer common.Address, value *big.Int) (*types.Transaction, error) {
	return _MTC.Contract.Redeem(&_MTC.TransactOpts, redeemer, value)
}

// RemoveAdmin is a paid mutator transaction binding the contract method 0x1785f53c.
//
// Solidity: function removeAdmin(address a) returns()
func (_MTC *MTCTransactor) RemoveAdmin(opts *bind.TransactOpts, a common.Address) (*types.Transaction, error) {
	return _MTC.contract.Transact(opts, "removeAdmin", a)
}

// RemoveAdmin is a paid mutator transaction binding the contract method 0x1785f53c.
//
// Solidity: function removeAdmin(address a) returns()
func (_MTC *MTCSession) RemoveAdmin(a common.Address) (*types.Transaction, error) {
	return _MTC.Contract.RemoveAdmin(&_MTC.TransactOpts, a)
}

// RemoveAdmin is a paid mutator transaction binding the contract method 0x1785f53c.
//
// Solidity: function removeAdmin(address a) returns()
func (_MTC *MTCTransactorSession) RemoveAdmin(a common.Address) (*types.Transaction, error) {
	return _MTC.Contract.RemoveAdmin(&_MTC.TransactOpts, a)
}

// RemoveBlackList is a paid mutator transaction binding the contract method 0xe4997dc5.
//
// Solidity: function removeBlackList(address _clearedUser) returns()
func (_MTC *MTCTransactor) RemoveBlackList(opts *bind.TransactOpts, _clearedUser common.Address) (*types.Transaction, error) {
	return _MTC.contract.Transact(opts, "removeBlackList", _clearedUser)
}

// RemoveBlackList is a paid mutator transaction binding the contract method 0xe4997dc5.
//
// Solidity: function removeBlackList(address _clearedUser) returns()
func (_MTC *MTCSession) RemoveBlackList(_clearedUser common.Address) (*types.Transaction, error) {
	return _MTC.Contract.RemoveBlackList(&_MTC.TransactOpts, _clearedUser)
}

// RemoveBlackList is a paid mutator transaction binding the contract method 0xe4997dc5.
//
// Solidity: function removeBlackList(address _clearedUser) returns()
func (_MTC *MTCTransactorSession) RemoveBlackList(_clearedUser common.Address) (*types.Transaction, error) {
	return _MTC.Contract.RemoveBlackList(&_MTC.TransactOpts, _clearedUser)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_MTC *MTCTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MTC.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_MTC *MTCSession) RenounceOwnership() (*types.Transaction, error) {
	return _MTC.Contract.RenounceOwnership(&_MTC.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_MTC *MTCTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _MTC.Contract.RenounceOwnership(&_MTC.TransactOpts)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns(bool)
func (_MTC *MTCTransactor) Transfer(opts *bind.TransactOpts, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _MTC.contract.Transact(opts, "transfer", _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns(bool)
func (_MTC *MTCSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _MTC.Contract.Transfer(&_MTC.TransactOpts, _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns(bool)
func (_MTC *MTCTransactorSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _MTC.Contract.Transfer(&_MTC.TransactOpts, _to, _value)
}

// TransferByAdmin is a paid mutator transaction binding the contract method 0x8eb17dfe.
//
// Solidity: function transferByAdmin(address from, address to, uint256 value) returns()
func (_MTC *MTCTransactor) TransferByAdmin(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _MTC.contract.Transact(opts, "transferByAdmin", from, to, value)
}

// TransferByAdmin is a paid mutator transaction binding the contract method 0x8eb17dfe.
//
// Solidity: function transferByAdmin(address from, address to, uint256 value) returns()
func (_MTC *MTCSession) TransferByAdmin(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _MTC.Contract.TransferByAdmin(&_MTC.TransactOpts, from, to, value)
}

// TransferByAdmin is a paid mutator transaction binding the contract method 0x8eb17dfe.
//
// Solidity: function transferByAdmin(address from, address to, uint256 value) returns()
func (_MTC *MTCTransactorSession) TransferByAdmin(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _MTC.Contract.TransferByAdmin(&_MTC.TransactOpts, from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _value) returns(bool)
func (_MTC *MTCTransactor) TransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _MTC.contract.Transact(opts, "transferFrom", _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _value) returns(bool)
func (_MTC *MTCSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _MTC.Contract.TransferFrom(&_MTC.TransactOpts, _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _value) returns(bool)
func (_MTC *MTCTransactorSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _MTC.Contract.TransferFrom(&_MTC.TransactOpts, _from, _to, _value)
}

// TransferKfiveReceiver is a paid mutator transaction binding the contract method 0x33738f96.
//
// Solidity: function transferKfiveReceiver(address newReceiver) returns()
func (_MTC *MTCTransactor) TransferKfiveReceiver(opts *bind.TransactOpts, newReceiver common.Address) (*types.Transaction, error) {
	return _MTC.contract.Transact(opts, "transferKfiveReceiver", newReceiver)
}

// TransferKfiveReceiver is a paid mutator transaction binding the contract method 0x33738f96.
//
// Solidity: function transferKfiveReceiver(address newReceiver) returns()
func (_MTC *MTCSession) TransferKfiveReceiver(newReceiver common.Address) (*types.Transaction, error) {
	return _MTC.Contract.TransferKfiveReceiver(&_MTC.TransactOpts, newReceiver)
}

// TransferKfiveReceiver is a paid mutator transaction binding the contract method 0x33738f96.
//
// Solidity: function transferKfiveReceiver(address newReceiver) returns()
func (_MTC *MTCTransactorSession) TransferKfiveReceiver(newReceiver common.Address) (*types.Transaction, error) {
	return _MTC.Contract.TransferKfiveReceiver(&_MTC.TransactOpts, newReceiver)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_MTC *MTCTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _MTC.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_MTC *MTCSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _MTC.Contract.TransferOwnership(&_MTC.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_MTC *MTCTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _MTC.Contract.TransferOwnership(&_MTC.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_MTC *MTCTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MTC.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_MTC *MTCSession) Unpause() (*types.Transaction, error) {
	return _MTC.Contract.Unpause(&_MTC.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_MTC *MTCTransactorSession) Unpause() (*types.Transaction, error) {
	return _MTC.Contract.Unpause(&_MTC.TransactOpts)
}

// UpdateKfiveAddress is a paid mutator transaction binding the contract method 0x1ff171c6.
//
// Solidity: function updateKfiveAddress(address newAddress) returns()
func (_MTC *MTCTransactor) UpdateKfiveAddress(opts *bind.TransactOpts, newAddress common.Address) (*types.Transaction, error) {
	return _MTC.contract.Transact(opts, "updateKfiveAddress", newAddress)
}

// UpdateKfiveAddress is a paid mutator transaction binding the contract method 0x1ff171c6.
//
// Solidity: function updateKfiveAddress(address newAddress) returns()
func (_MTC *MTCSession) UpdateKfiveAddress(newAddress common.Address) (*types.Transaction, error) {
	return _MTC.Contract.UpdateKfiveAddress(&_MTC.TransactOpts, newAddress)
}

// UpdateKfiveAddress is a paid mutator transaction binding the contract method 0x1ff171c6.
//
// Solidity: function updateKfiveAddress(address newAddress) returns()
func (_MTC *MTCTransactorSession) UpdateKfiveAddress(newAddress common.Address) (*types.Transaction, error) {
	return _MTC.Contract.UpdateKfiveAddress(&_MTC.TransactOpts, newAddress)
}

// MTCAddedBlackListIterator is returned from FilterAddedBlackList and is used to iterate over the raw logs and unpacked data for AddedBlackList events raised by the MTC contract.
type MTCAddedBlackListIterator struct {
	Event *MTCAddedBlackList // Event containing the contract specifics and raw log

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
func (it *MTCAddedBlackListIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MTCAddedBlackList)
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
		it.Event = new(MTCAddedBlackList)
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
func (it *MTCAddedBlackListIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MTCAddedBlackListIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MTCAddedBlackList represents a AddedBlackList event raised by the MTC contract.
type MTCAddedBlackList struct {
	User common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterAddedBlackList is a free log retrieval operation binding the contract event 0x42e160154868087d6bfdc0ca23d96a1c1cfa32f1b72ba9ba27b69b98a0d819dc.
//
// Solidity: event AddedBlackList(address _user)
func (_MTC *MTCFilterer) FilterAddedBlackList(opts *bind.FilterOpts) (*MTCAddedBlackListIterator, error) {

	logs, sub, err := _MTC.contract.FilterLogs(opts, "AddedBlackList")
	if err != nil {
		return nil, err
	}
	return &MTCAddedBlackListIterator{contract: _MTC.contract, event: "AddedBlackList", logs: logs, sub: sub}, nil
}

// WatchAddedBlackList is a free log subscription operation binding the contract event 0x42e160154868087d6bfdc0ca23d96a1c1cfa32f1b72ba9ba27b69b98a0d819dc.
//
// Solidity: event AddedBlackList(address _user)
func (_MTC *MTCFilterer) WatchAddedBlackList(opts *bind.WatchOpts, sink chan<- *MTCAddedBlackList) (event.Subscription, error) {

	logs, sub, err := _MTC.contract.WatchLogs(opts, "AddedBlackList")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MTCAddedBlackList)
				if err := _MTC.contract.UnpackLog(event, "AddedBlackList", log); err != nil {
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
func (_MTC *MTCFilterer) ParseAddedBlackList(log types.Log) (*MTCAddedBlackList, error) {
	event := new(MTCAddedBlackList)
	if err := _MTC.contract.UnpackLog(event, "AddedBlackList", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MTCApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the MTC contract.
type MTCApprovalIterator struct {
	Event *MTCApproval // Event containing the contract specifics and raw log

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
func (it *MTCApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MTCApproval)
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
		it.Event = new(MTCApproval)
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
func (it *MTCApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MTCApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MTCApproval represents a Approval event raised by the MTC contract.
type MTCApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_MTC *MTCFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*MTCApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _MTC.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &MTCApprovalIterator{contract: _MTC.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_MTC *MTCFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *MTCApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _MTC.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MTCApproval)
				if err := _MTC.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_MTC *MTCFilterer) ParseApproval(log types.Log) (*MTCApproval, error) {
	event := new(MTCApproval)
	if err := _MTC.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MTCDestroyedBlackFundsIterator is returned from FilterDestroyedBlackFunds and is used to iterate over the raw logs and unpacked data for DestroyedBlackFunds events raised by the MTC contract.
type MTCDestroyedBlackFundsIterator struct {
	Event *MTCDestroyedBlackFunds // Event containing the contract specifics and raw log

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
func (it *MTCDestroyedBlackFundsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MTCDestroyedBlackFunds)
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
		it.Event = new(MTCDestroyedBlackFunds)
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
func (it *MTCDestroyedBlackFundsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MTCDestroyedBlackFundsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MTCDestroyedBlackFunds represents a DestroyedBlackFunds event raised by the MTC contract.
type MTCDestroyedBlackFunds struct {
	BlackListedUser common.Address
	Balance         *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterDestroyedBlackFunds is a free log retrieval operation binding the contract event 0x61e6e66b0d6339b2980aecc6ccc0039736791f0ccde9ed512e789a7fbdd698c6.
//
// Solidity: event DestroyedBlackFunds(address _blackListedUser, uint256 _balance)
func (_MTC *MTCFilterer) FilterDestroyedBlackFunds(opts *bind.FilterOpts) (*MTCDestroyedBlackFundsIterator, error) {

	logs, sub, err := _MTC.contract.FilterLogs(opts, "DestroyedBlackFunds")
	if err != nil {
		return nil, err
	}
	return &MTCDestroyedBlackFundsIterator{contract: _MTC.contract, event: "DestroyedBlackFunds", logs: logs, sub: sub}, nil
}

// WatchDestroyedBlackFunds is a free log subscription operation binding the contract event 0x61e6e66b0d6339b2980aecc6ccc0039736791f0ccde9ed512e789a7fbdd698c6.
//
// Solidity: event DestroyedBlackFunds(address _blackListedUser, uint256 _balance)
func (_MTC *MTCFilterer) WatchDestroyedBlackFunds(opts *bind.WatchOpts, sink chan<- *MTCDestroyedBlackFunds) (event.Subscription, error) {

	logs, sub, err := _MTC.contract.WatchLogs(opts, "DestroyedBlackFunds")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MTCDestroyedBlackFunds)
				if err := _MTC.contract.UnpackLog(event, "DestroyedBlackFunds", log); err != nil {
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
func (_MTC *MTCFilterer) ParseDestroyedBlackFunds(log types.Log) (*MTCDestroyedBlackFunds, error) {
	event := new(MTCDestroyedBlackFunds)
	if err := _MTC.contract.UnpackLog(event, "DestroyedBlackFunds", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MTCExchangeFromKFIVEIterator is returned from FilterExchangeFromKFIVE and is used to iterate over the raw logs and unpacked data for ExchangeFromKFIVE events raised by the MTC contract.
type MTCExchangeFromKFIVEIterator struct {
	Event *MTCExchangeFromKFIVE // Event containing the contract specifics and raw log

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
func (it *MTCExchangeFromKFIVEIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MTCExchangeFromKFIVE)
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
		it.Event = new(MTCExchangeFromKFIVE)
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
func (it *MTCExchangeFromKFIVEIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MTCExchangeFromKFIVEIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MTCExchangeFromKFIVE represents a ExchangeFromKFIVE event raised by the MTC contract.
type MTCExchangeFromKFIVE struct {
	User  common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterExchangeFromKFIVE is a free log retrieval operation binding the contract event 0xb43ac86e09c923c3e090221d36f594325e6ae96c37063fedab869c0b900a85ea.
//
// Solidity: event ExchangeFromKFIVE(address user, uint256 value)
func (_MTC *MTCFilterer) FilterExchangeFromKFIVE(opts *bind.FilterOpts) (*MTCExchangeFromKFIVEIterator, error) {

	logs, sub, err := _MTC.contract.FilterLogs(opts, "ExchangeFromKFIVE")
	if err != nil {
		return nil, err
	}
	return &MTCExchangeFromKFIVEIterator{contract: _MTC.contract, event: "ExchangeFromKFIVE", logs: logs, sub: sub}, nil
}

// WatchExchangeFromKFIVE is a free log subscription operation binding the contract event 0xb43ac86e09c923c3e090221d36f594325e6ae96c37063fedab869c0b900a85ea.
//
// Solidity: event ExchangeFromKFIVE(address user, uint256 value)
func (_MTC *MTCFilterer) WatchExchangeFromKFIVE(opts *bind.WatchOpts, sink chan<- *MTCExchangeFromKFIVE) (event.Subscription, error) {

	logs, sub, err := _MTC.contract.WatchLogs(opts, "ExchangeFromKFIVE")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MTCExchangeFromKFIVE)
				if err := _MTC.contract.UnpackLog(event, "ExchangeFromKFIVE", log); err != nil {
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

// ParseExchangeFromKFIVE is a log parse operation binding the contract event 0xb43ac86e09c923c3e090221d36f594325e6ae96c37063fedab869c0b900a85ea.
//
// Solidity: event ExchangeFromKFIVE(address user, uint256 value)
func (_MTC *MTCFilterer) ParseExchangeFromKFIVE(log types.Log) (*MTCExchangeFromKFIVE, error) {
	event := new(MTCExchangeFromKFIVE)
	if err := _MTC.contract.UnpackLog(event, "ExchangeFromKFIVE", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MTCKfiveAddressChangedIterator is returned from FilterKfiveAddressChanged and is used to iterate over the raw logs and unpacked data for KfiveAddressChanged events raised by the MTC contract.
type MTCKfiveAddressChangedIterator struct {
	Event *MTCKfiveAddressChanged // Event containing the contract specifics and raw log

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
func (it *MTCKfiveAddressChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MTCKfiveAddressChanged)
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
		it.Event = new(MTCKfiveAddressChanged)
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
func (it *MTCKfiveAddressChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MTCKfiveAddressChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MTCKfiveAddressChanged represents a KfiveAddressChanged event raised by the MTC contract.
type MTCKfiveAddressChanged struct {
	KfiveAddress common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterKfiveAddressChanged is a free log retrieval operation binding the contract event 0x4c5132ff0e1cf6de4e01db06e2ed201f844477488336708093ca61092d28ee9e.
//
// Solidity: event KfiveAddressChanged(address kfiveAddress)
func (_MTC *MTCFilterer) FilterKfiveAddressChanged(opts *bind.FilterOpts) (*MTCKfiveAddressChangedIterator, error) {

	logs, sub, err := _MTC.contract.FilterLogs(opts, "KfiveAddressChanged")
	if err != nil {
		return nil, err
	}
	return &MTCKfiveAddressChangedIterator{contract: _MTC.contract, event: "KfiveAddressChanged", logs: logs, sub: sub}, nil
}

// WatchKfiveAddressChanged is a free log subscription operation binding the contract event 0x4c5132ff0e1cf6de4e01db06e2ed201f844477488336708093ca61092d28ee9e.
//
// Solidity: event KfiveAddressChanged(address kfiveAddress)
func (_MTC *MTCFilterer) WatchKfiveAddressChanged(opts *bind.WatchOpts, sink chan<- *MTCKfiveAddressChanged) (event.Subscription, error) {

	logs, sub, err := _MTC.contract.WatchLogs(opts, "KfiveAddressChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MTCKfiveAddressChanged)
				if err := _MTC.contract.UnpackLog(event, "KfiveAddressChanged", log); err != nil {
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

// ParseKfiveAddressChanged is a log parse operation binding the contract event 0x4c5132ff0e1cf6de4e01db06e2ed201f844477488336708093ca61092d28ee9e.
//
// Solidity: event KfiveAddressChanged(address kfiveAddress)
func (_MTC *MTCFilterer) ParseKfiveAddressChanged(log types.Log) (*MTCKfiveAddressChanged, error) {
	event := new(MTCKfiveAddressChanged)
	if err := _MTC.contract.UnpackLog(event, "KfiveAddressChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MTCKfiveReceiverTransferredIterator is returned from FilterKfiveReceiverTransferred and is used to iterate over the raw logs and unpacked data for KfiveReceiverTransferred events raised by the MTC contract.
type MTCKfiveReceiverTransferredIterator struct {
	Event *MTCKfiveReceiverTransferred // Event containing the contract specifics and raw log

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
func (it *MTCKfiveReceiverTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MTCKfiveReceiverTransferred)
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
		it.Event = new(MTCKfiveReceiverTransferred)
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
func (it *MTCKfiveReceiverTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MTCKfiveReceiverTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MTCKfiveReceiverTransferred represents a KfiveReceiverTransferred event raised by the MTC contract.
type MTCKfiveReceiverTransferred struct {
	OldReceiver common.Address
	NewReceiver common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterKfiveReceiverTransferred is a free log retrieval operation binding the contract event 0x628d49ba701aa7aeb6c972bd501fe788482a30eeaad3736ab7a3eb1d5ce19f5a.
//
// Solidity: event KfiveReceiverTransferred(address oldReceiver, address newReceiver)
func (_MTC *MTCFilterer) FilterKfiveReceiverTransferred(opts *bind.FilterOpts) (*MTCKfiveReceiverTransferredIterator, error) {

	logs, sub, err := _MTC.contract.FilterLogs(opts, "KfiveReceiverTransferred")
	if err != nil {
		return nil, err
	}
	return &MTCKfiveReceiverTransferredIterator{contract: _MTC.contract, event: "KfiveReceiverTransferred", logs: logs, sub: sub}, nil
}

// WatchKfiveReceiverTransferred is a free log subscription operation binding the contract event 0x628d49ba701aa7aeb6c972bd501fe788482a30eeaad3736ab7a3eb1d5ce19f5a.
//
// Solidity: event KfiveReceiverTransferred(address oldReceiver, address newReceiver)
func (_MTC *MTCFilterer) WatchKfiveReceiverTransferred(opts *bind.WatchOpts, sink chan<- *MTCKfiveReceiverTransferred) (event.Subscription, error) {

	logs, sub, err := _MTC.contract.WatchLogs(opts, "KfiveReceiverTransferred")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MTCKfiveReceiverTransferred)
				if err := _MTC.contract.UnpackLog(event, "KfiveReceiverTransferred", log); err != nil {
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

// ParseKfiveReceiverTransferred is a log parse operation binding the contract event 0x628d49ba701aa7aeb6c972bd501fe788482a30eeaad3736ab7a3eb1d5ce19f5a.
//
// Solidity: event KfiveReceiverTransferred(address oldReceiver, address newReceiver)
func (_MTC *MTCFilterer) ParseKfiveReceiverTransferred(log types.Log) (*MTCKfiveReceiverTransferred, error) {
	event := new(MTCKfiveReceiverTransferred)
	if err := _MTC.contract.UnpackLog(event, "KfiveReceiverTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MTCOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the MTC contract.
type MTCOwnershipTransferredIterator struct {
	Event *MTCOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *MTCOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MTCOwnershipTransferred)
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
		it.Event = new(MTCOwnershipTransferred)
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
func (it *MTCOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MTCOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MTCOwnershipTransferred represents a OwnershipTransferred event raised by the MTC contract.
type MTCOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_MTC *MTCFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*MTCOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _MTC.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &MTCOwnershipTransferredIterator{contract: _MTC.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_MTC *MTCFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *MTCOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _MTC.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MTCOwnershipTransferred)
				if err := _MTC.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_MTC *MTCFilterer) ParseOwnershipTransferred(log types.Log) (*MTCOwnershipTransferred, error) {
	event := new(MTCOwnershipTransferred)
	if err := _MTC.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MTCPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the MTC contract.
type MTCPausedIterator struct {
	Event *MTCPaused // Event containing the contract specifics and raw log

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
func (it *MTCPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MTCPaused)
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
		it.Event = new(MTCPaused)
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
func (it *MTCPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MTCPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MTCPaused represents a Paused event raised by the MTC contract.
type MTCPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_MTC *MTCFilterer) FilterPaused(opts *bind.FilterOpts) (*MTCPausedIterator, error) {

	logs, sub, err := _MTC.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &MTCPausedIterator{contract: _MTC.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_MTC *MTCFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *MTCPaused) (event.Subscription, error) {

	logs, sub, err := _MTC.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MTCPaused)
				if err := _MTC.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_MTC *MTCFilterer) ParsePaused(log types.Log) (*MTCPaused, error) {
	event := new(MTCPaused)
	if err := _MTC.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MTCRemovedBlackListIterator is returned from FilterRemovedBlackList and is used to iterate over the raw logs and unpacked data for RemovedBlackList events raised by the MTC contract.
type MTCRemovedBlackListIterator struct {
	Event *MTCRemovedBlackList // Event containing the contract specifics and raw log

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
func (it *MTCRemovedBlackListIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MTCRemovedBlackList)
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
		it.Event = new(MTCRemovedBlackList)
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
func (it *MTCRemovedBlackListIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MTCRemovedBlackListIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MTCRemovedBlackList represents a RemovedBlackList event raised by the MTC contract.
type MTCRemovedBlackList struct {
	User common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterRemovedBlackList is a free log retrieval operation binding the contract event 0xd7e9ec6e6ecd65492dce6bf513cd6867560d49544421d0783ddf06e76c24470c.
//
// Solidity: event RemovedBlackList(address _user)
func (_MTC *MTCFilterer) FilterRemovedBlackList(opts *bind.FilterOpts) (*MTCRemovedBlackListIterator, error) {

	logs, sub, err := _MTC.contract.FilterLogs(opts, "RemovedBlackList")
	if err != nil {
		return nil, err
	}
	return &MTCRemovedBlackListIterator{contract: _MTC.contract, event: "RemovedBlackList", logs: logs, sub: sub}, nil
}

// WatchRemovedBlackList is a free log subscription operation binding the contract event 0xd7e9ec6e6ecd65492dce6bf513cd6867560d49544421d0783ddf06e76c24470c.
//
// Solidity: event RemovedBlackList(address _user)
func (_MTC *MTCFilterer) WatchRemovedBlackList(opts *bind.WatchOpts, sink chan<- *MTCRemovedBlackList) (event.Subscription, error) {

	logs, sub, err := _MTC.contract.WatchLogs(opts, "RemovedBlackList")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MTCRemovedBlackList)
				if err := _MTC.contract.UnpackLog(event, "RemovedBlackList", log); err != nil {
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
func (_MTC *MTCFilterer) ParseRemovedBlackList(log types.Log) (*MTCRemovedBlackList, error) {
	event := new(MTCRemovedBlackList)
	if err := _MTC.contract.UnpackLog(event, "RemovedBlackList", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MTCTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the MTC contract.
type MTCTransferIterator struct {
	Event *MTCTransfer // Event containing the contract specifics and raw log

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
func (it *MTCTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MTCTransfer)
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
		it.Event = new(MTCTransfer)
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
func (it *MTCTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MTCTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MTCTransfer represents a Transfer event raised by the MTC contract.
type MTCTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_MTC *MTCFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*MTCTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _MTC.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &MTCTransferIterator{contract: _MTC.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_MTC *MTCFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *MTCTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _MTC.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MTCTransfer)
				if err := _MTC.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_MTC *MTCFilterer) ParseTransfer(log types.Log) (*MTCTransfer, error) {
	event := new(MTCTransfer)
	if err := _MTC.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MTCUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the MTC contract.
type MTCUnpausedIterator struct {
	Event *MTCUnpaused // Event containing the contract specifics and raw log

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
func (it *MTCUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MTCUnpaused)
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
		it.Event = new(MTCUnpaused)
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
func (it *MTCUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MTCUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MTCUnpaused represents a Unpaused event raised by the MTC contract.
type MTCUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_MTC *MTCFilterer) FilterUnpaused(opts *bind.FilterOpts) (*MTCUnpausedIterator, error) {

	logs, sub, err := _MTC.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &MTCUnpausedIterator{contract: _MTC.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_MTC *MTCFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *MTCUnpaused) (event.Subscription, error) {

	logs, sub, err := _MTC.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MTCUnpaused)
				if err := _MTC.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
func (_MTC *MTCFilterer) ParseUnpaused(log types.Log) (*MTCUnpaused, error) {
	event := new(MTCUnpaused)
	if err := _MTC.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MTCIssueIterator is returned from FilterIssue and is used to iterate over the raw logs and unpacked data for Issue events raised by the MTC contract.
type MTCIssueIterator struct {
	Event *MTCIssue // Event containing the contract specifics and raw log

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
func (it *MTCIssueIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MTCIssue)
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
		it.Event = new(MTCIssue)
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
func (it *MTCIssueIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MTCIssueIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MTCIssue represents a Issue event raised by the MTC contract.
type MTCIssue struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterIssue is a free log retrieval operation binding the contract event 0x44b237bbd5491bb03d9d735219b4226fac4cb366696bbd63a7dc6f608af4e547.
//
// Solidity: event __issue()
func (_MTC *MTCFilterer) FilterIssue(opts *bind.FilterOpts) (*MTCIssueIterator, error) {

	logs, sub, err := _MTC.contract.FilterLogs(opts, "__issue")
	if err != nil {
		return nil, err
	}
	return &MTCIssueIterator{contract: _MTC.contract, event: "__issue", logs: logs, sub: sub}, nil
}

// WatchIssue is a free log subscription operation binding the contract event 0x44b237bbd5491bb03d9d735219b4226fac4cb366696bbd63a7dc6f608af4e547.
//
// Solidity: event __issue()
func (_MTC *MTCFilterer) WatchIssue(opts *bind.WatchOpts, sink chan<- *MTCIssue) (event.Subscription, error) {

	logs, sub, err := _MTC.contract.WatchLogs(opts, "__issue")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MTCIssue)
				if err := _MTC.contract.UnpackLog(event, "__issue", log); err != nil {
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

// ParseIssue is a log parse operation binding the contract event 0x44b237bbd5491bb03d9d735219b4226fac4cb366696bbd63a7dc6f608af4e547.
//
// Solidity: event __issue()
func (_MTC *MTCFilterer) ParseIssue(log types.Log) (*MTCIssue, error) {
	event := new(MTCIssue)
	if err := _MTC.contract.UnpackLog(event, "__issue", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MTCRedeemIterator is returned from FilterRedeem and is used to iterate over the raw logs and unpacked data for Redeem events raised by the MTC contract.
type MTCRedeemIterator struct {
	Event *MTCRedeem // Event containing the contract specifics and raw log

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
func (it *MTCRedeemIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MTCRedeem)
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
		it.Event = new(MTCRedeem)
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
func (it *MTCRedeemIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MTCRedeemIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MTCRedeem represents a Redeem event raised by the MTC contract.
type MTCRedeem struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterRedeem is a free log retrieval operation binding the contract event 0xd41c85faca8de216b4dbc5c57d0a60237745d8ea20e32ac116e8dd4f95fae3f9.
//
// Solidity: event __redeem()
func (_MTC *MTCFilterer) FilterRedeem(opts *bind.FilterOpts) (*MTCRedeemIterator, error) {

	logs, sub, err := _MTC.contract.FilterLogs(opts, "__redeem")
	if err != nil {
		return nil, err
	}
	return &MTCRedeemIterator{contract: _MTC.contract, event: "__redeem", logs: logs, sub: sub}, nil
}

// WatchRedeem is a free log subscription operation binding the contract event 0xd41c85faca8de216b4dbc5c57d0a60237745d8ea20e32ac116e8dd4f95fae3f9.
//
// Solidity: event __redeem()
func (_MTC *MTCFilterer) WatchRedeem(opts *bind.WatchOpts, sink chan<- *MTCRedeem) (event.Subscription, error) {

	logs, sub, err := _MTC.contract.WatchLogs(opts, "__redeem")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MTCRedeem)
				if err := _MTC.contract.UnpackLog(event, "__redeem", log); err != nil {
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

// ParseRedeem is a log parse operation binding the contract event 0xd41c85faca8de216b4dbc5c57d0a60237745d8ea20e32ac116e8dd4f95fae3f9.
//
// Solidity: event __redeem()
func (_MTC *MTCFilterer) ParseRedeem(log types.Log) (*MTCRedeem, error) {
	event := new(MTCRedeem)
	if err := _MTC.contract.UnpackLog(event, "__redeem", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MTCTransferByAdminIterator is returned from FilterTransferByAdmin and is used to iterate over the raw logs and unpacked data for TransferByAdmin events raised by the MTC contract.
type MTCTransferByAdminIterator struct {
	Event *MTCTransferByAdmin // Event containing the contract specifics and raw log

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
func (it *MTCTransferByAdminIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MTCTransferByAdmin)
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
		it.Event = new(MTCTransferByAdmin)
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
func (it *MTCTransferByAdminIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MTCTransferByAdminIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MTCTransferByAdmin represents a TransferByAdmin event raised by the MTC contract.
type MTCTransferByAdmin struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterTransferByAdmin is a free log retrieval operation binding the contract event 0x8c32870a771c16ccafda4603ff1d2ce8d72fb7613fd5502d71466b0f7dd8b442.
//
// Solidity: event __transferByAdmin()
func (_MTC *MTCFilterer) FilterTransferByAdmin(opts *bind.FilterOpts) (*MTCTransferByAdminIterator, error) {

	logs, sub, err := _MTC.contract.FilterLogs(opts, "__transferByAdmin")
	if err != nil {
		return nil, err
	}
	return &MTCTransferByAdminIterator{contract: _MTC.contract, event: "__transferByAdmin", logs: logs, sub: sub}, nil
}

// WatchTransferByAdmin is a free log subscription operation binding the contract event 0x8c32870a771c16ccafda4603ff1d2ce8d72fb7613fd5502d71466b0f7dd8b442.
//
// Solidity: event __transferByAdmin()
func (_MTC *MTCFilterer) WatchTransferByAdmin(opts *bind.WatchOpts, sink chan<- *MTCTransferByAdmin) (event.Subscription, error) {

	logs, sub, err := _MTC.contract.WatchLogs(opts, "__transferByAdmin")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MTCTransferByAdmin)
				if err := _MTC.contract.UnpackLog(event, "__transferByAdmin", log); err != nil {
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

// ParseTransferByAdmin is a log parse operation binding the contract event 0x8c32870a771c16ccafda4603ff1d2ce8d72fb7613fd5502d71466b0f7dd8b442.
//
// Solidity: event __transferByAdmin()
func (_MTC *MTCFilterer) ParseTransferByAdmin(log types.Log) (*MTCTransferByAdmin, error) {
	event := new(MTCTransferByAdmin)
	if err := _MTC.contract.UnpackLog(event, "__transferByAdmin", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
