// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package LuckyVault

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

// LuckyVaultMetaData contains all meta data concerning the LuckyVault contract.
var LuckyVaultMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tradeToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"depositToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"rewardToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"endedAt\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AdminWithdrawFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"DepositFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InsufficientBalance\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidAmount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidContract\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidWithdrawAmount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoReward\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RewardFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SendDepositTokenFailed\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"}],\"name\":\"TooEarly\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"}],\"name\":\"TooLate\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TradeFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"WithdrawFailed\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"AdminWithdrawSuccessful\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"rewardAmount\",\"type\":\"uint256\"}],\"name\":\"DepositSuccessful\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tradeAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"rewardAmount\",\"type\":\"uint256\"}],\"name\":\"TradeSuccessful\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"WithdrawSuccessful\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PAUSER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getRoleMember\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleMemberCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"trade\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdrawByAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"getDeposit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTradeRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTradeRewardRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getDepositRewardRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getEndedAt\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// LuckyVaultABI is the input ABI used to generate the binding from.
// Deprecated: Use LuckyVaultMetaData.ABI instead.
var LuckyVaultABI = LuckyVaultMetaData.ABI

// LuckyVault is an auto generated Go binding around an Ethereum contract.
type LuckyVault struct {
	LuckyVaultCaller     // Read-only binding to the contract
	LuckyVaultTransactor // Write-only binding to the contract
	LuckyVaultFilterer   // Log filterer for contract events
}

// LuckyVaultCaller is an auto generated read-only Go binding around an Ethereum contract.
type LuckyVaultCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LuckyVaultTransactor is an auto generated write-only Go binding around an Ethereum contract.
type LuckyVaultTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LuckyVaultFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type LuckyVaultFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LuckyVaultSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type LuckyVaultSession struct {
	Contract     *LuckyVault       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// LuckyVaultCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type LuckyVaultCallerSession struct {
	Contract *LuckyVaultCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// LuckyVaultTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type LuckyVaultTransactorSession struct {
	Contract     *LuckyVaultTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// LuckyVaultRaw is an auto generated low-level Go binding around an Ethereum contract.
type LuckyVaultRaw struct {
	Contract *LuckyVault // Generic contract binding to access the raw methods on
}

// LuckyVaultCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type LuckyVaultCallerRaw struct {
	Contract *LuckyVaultCaller // Generic read-only contract binding to access the raw methods on
}

// LuckyVaultTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type LuckyVaultTransactorRaw struct {
	Contract *LuckyVaultTransactor // Generic write-only contract binding to access the raw methods on
}

// NewLuckyVault creates a new instance of LuckyVault, bound to a specific deployed contract.
func NewLuckyVault(address common.Address, backend bind.ContractBackend) (*LuckyVault, error) {
	contract, err := bindLuckyVault(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &LuckyVault{LuckyVaultCaller: LuckyVaultCaller{contract: contract}, LuckyVaultTransactor: LuckyVaultTransactor{contract: contract}, LuckyVaultFilterer: LuckyVaultFilterer{contract: contract}}, nil
}

// NewLuckyVaultCaller creates a new read-only instance of LuckyVault, bound to a specific deployed contract.
func NewLuckyVaultCaller(address common.Address, caller bind.ContractCaller) (*LuckyVaultCaller, error) {
	contract, err := bindLuckyVault(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &LuckyVaultCaller{contract: contract}, nil
}

// NewLuckyVaultTransactor creates a new write-only instance of LuckyVault, bound to a specific deployed contract.
func NewLuckyVaultTransactor(address common.Address, transactor bind.ContractTransactor) (*LuckyVaultTransactor, error) {
	contract, err := bindLuckyVault(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &LuckyVaultTransactor{contract: contract}, nil
}

// NewLuckyVaultFilterer creates a new log filterer instance of LuckyVault, bound to a specific deployed contract.
func NewLuckyVaultFilterer(address common.Address, filterer bind.ContractFilterer) (*LuckyVaultFilterer, error) {
	contract, err := bindLuckyVault(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &LuckyVaultFilterer{contract: contract}, nil
}

// bindLuckyVault binds a generic wrapper to an already deployed contract.
func bindLuckyVault(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(LuckyVaultABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LuckyVault *LuckyVaultRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LuckyVault.Contract.LuckyVaultCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LuckyVault *LuckyVaultRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LuckyVault.Contract.LuckyVaultTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LuckyVault *LuckyVaultRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LuckyVault.Contract.LuckyVaultTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LuckyVault *LuckyVaultCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LuckyVault.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LuckyVault *LuckyVaultTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LuckyVault.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LuckyVault *LuckyVaultTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LuckyVault.Contract.contract.Transact(opts, method, params...)
}

// ADMINROLE is a free data retrieval call binding the contract method 0x75b238fc.
//
// Solidity: function ADMIN_ROLE() view returns(bytes32)
func (_LuckyVault *LuckyVaultCaller) ADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _LuckyVault.contract.Call(opts, &out, "ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ADMINROLE is a free data retrieval call binding the contract method 0x75b238fc.
//
// Solidity: function ADMIN_ROLE() view returns(bytes32)
func (_LuckyVault *LuckyVaultSession) ADMINROLE() ([32]byte, error) {
	return _LuckyVault.Contract.ADMINROLE(&_LuckyVault.CallOpts)
}

// ADMINROLE is a free data retrieval call binding the contract method 0x75b238fc.
//
// Solidity: function ADMIN_ROLE() view returns(bytes32)
func (_LuckyVault *LuckyVaultCallerSession) ADMINROLE() ([32]byte, error) {
	return _LuckyVault.Contract.ADMINROLE(&_LuckyVault.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_LuckyVault *LuckyVaultCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _LuckyVault.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_LuckyVault *LuckyVaultSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _LuckyVault.Contract.DEFAULTADMINROLE(&_LuckyVault.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_LuckyVault *LuckyVaultCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _LuckyVault.Contract.DEFAULTADMINROLE(&_LuckyVault.CallOpts)
}

// PAUSERROLE is a free data retrieval call binding the contract method 0xe63ab1e9.
//
// Solidity: function PAUSER_ROLE() view returns(bytes32)
func (_LuckyVault *LuckyVaultCaller) PAUSERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _LuckyVault.contract.Call(opts, &out, "PAUSER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// PAUSERROLE is a free data retrieval call binding the contract method 0xe63ab1e9.
//
// Solidity: function PAUSER_ROLE() view returns(bytes32)
func (_LuckyVault *LuckyVaultSession) PAUSERROLE() ([32]byte, error) {
	return _LuckyVault.Contract.PAUSERROLE(&_LuckyVault.CallOpts)
}

// PAUSERROLE is a free data retrieval call binding the contract method 0xe63ab1e9.
//
// Solidity: function PAUSER_ROLE() view returns(bytes32)
func (_LuckyVault *LuckyVaultCallerSession) PAUSERROLE() ([32]byte, error) {
	return _LuckyVault.Contract.PAUSERROLE(&_LuckyVault.CallOpts)
}

// GetDeposit is a free data retrieval call binding the contract method 0xe1254fba.
//
// Solidity: function getDeposit(address owner) view returns(uint256)
func (_LuckyVault *LuckyVaultCaller) GetDeposit(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _LuckyVault.contract.Call(opts, &out, "getDeposit", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetDeposit is a free data retrieval call binding the contract method 0xe1254fba.
//
// Solidity: function getDeposit(address owner) view returns(uint256)
func (_LuckyVault *LuckyVaultSession) GetDeposit(owner common.Address) (*big.Int, error) {
	return _LuckyVault.Contract.GetDeposit(&_LuckyVault.CallOpts, owner)
}

// GetDeposit is a free data retrieval call binding the contract method 0xe1254fba.
//
// Solidity: function getDeposit(address owner) view returns(uint256)
func (_LuckyVault *LuckyVaultCallerSession) GetDeposit(owner common.Address) (*big.Int, error) {
	return _LuckyVault.Contract.GetDeposit(&_LuckyVault.CallOpts, owner)
}

// GetDepositRewardRate is a free data retrieval call binding the contract method 0x6f15f646.
//
// Solidity: function getDepositRewardRate() view returns(uint256)
func (_LuckyVault *LuckyVaultCaller) GetDepositRewardRate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LuckyVault.contract.Call(opts, &out, "getDepositRewardRate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetDepositRewardRate is a free data retrieval call binding the contract method 0x6f15f646.
//
// Solidity: function getDepositRewardRate() view returns(uint256)
func (_LuckyVault *LuckyVaultSession) GetDepositRewardRate() (*big.Int, error) {
	return _LuckyVault.Contract.GetDepositRewardRate(&_LuckyVault.CallOpts)
}

// GetDepositRewardRate is a free data retrieval call binding the contract method 0x6f15f646.
//
// Solidity: function getDepositRewardRate() view returns(uint256)
func (_LuckyVault *LuckyVaultCallerSession) GetDepositRewardRate() (*big.Int, error) {
	return _LuckyVault.Contract.GetDepositRewardRate(&_LuckyVault.CallOpts)
}

// GetEndedAt is a free data retrieval call binding the contract method 0x1906ab9b.
//
// Solidity: function getEndedAt() view returns(uint256)
func (_LuckyVault *LuckyVaultCaller) GetEndedAt(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LuckyVault.contract.Call(opts, &out, "getEndedAt")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetEndedAt is a free data retrieval call binding the contract method 0x1906ab9b.
//
// Solidity: function getEndedAt() view returns(uint256)
func (_LuckyVault *LuckyVaultSession) GetEndedAt() (*big.Int, error) {
	return _LuckyVault.Contract.GetEndedAt(&_LuckyVault.CallOpts)
}

// GetEndedAt is a free data retrieval call binding the contract method 0x1906ab9b.
//
// Solidity: function getEndedAt() view returns(uint256)
func (_LuckyVault *LuckyVaultCallerSession) GetEndedAt() (*big.Int, error) {
	return _LuckyVault.Contract.GetEndedAt(&_LuckyVault.CallOpts)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_LuckyVault *LuckyVaultCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _LuckyVault.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_LuckyVault *LuckyVaultSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _LuckyVault.Contract.GetRoleAdmin(&_LuckyVault.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_LuckyVault *LuckyVaultCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _LuckyVault.Contract.GetRoleAdmin(&_LuckyVault.CallOpts, role)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_LuckyVault *LuckyVaultCaller) GetRoleMember(opts *bind.CallOpts, role [32]byte, index *big.Int) (common.Address, error) {
	var out []interface{}
	err := _LuckyVault.contract.Call(opts, &out, "getRoleMember", role, index)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_LuckyVault *LuckyVaultSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _LuckyVault.Contract.GetRoleMember(&_LuckyVault.CallOpts, role, index)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_LuckyVault *LuckyVaultCallerSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _LuckyVault.Contract.GetRoleMember(&_LuckyVault.CallOpts, role, index)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_LuckyVault *LuckyVaultCaller) GetRoleMemberCount(opts *bind.CallOpts, role [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _LuckyVault.contract.Call(opts, &out, "getRoleMemberCount", role)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_LuckyVault *LuckyVaultSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _LuckyVault.Contract.GetRoleMemberCount(&_LuckyVault.CallOpts, role)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_LuckyVault *LuckyVaultCallerSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _LuckyVault.Contract.GetRoleMemberCount(&_LuckyVault.CallOpts, role)
}

// GetTradeRate is a free data retrieval call binding the contract method 0x9fcfe0e4.
//
// Solidity: function getTradeRate() view returns(uint256)
func (_LuckyVault *LuckyVaultCaller) GetTradeRate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LuckyVault.contract.Call(opts, &out, "getTradeRate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTradeRate is a free data retrieval call binding the contract method 0x9fcfe0e4.
//
// Solidity: function getTradeRate() view returns(uint256)
func (_LuckyVault *LuckyVaultSession) GetTradeRate() (*big.Int, error) {
	return _LuckyVault.Contract.GetTradeRate(&_LuckyVault.CallOpts)
}

// GetTradeRate is a free data retrieval call binding the contract method 0x9fcfe0e4.
//
// Solidity: function getTradeRate() view returns(uint256)
func (_LuckyVault *LuckyVaultCallerSession) GetTradeRate() (*big.Int, error) {
	return _LuckyVault.Contract.GetTradeRate(&_LuckyVault.CallOpts)
}

// GetTradeRewardRate is a free data retrieval call binding the contract method 0xe8fcc067.
//
// Solidity: function getTradeRewardRate() view returns(uint256)
func (_LuckyVault *LuckyVaultCaller) GetTradeRewardRate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LuckyVault.contract.Call(opts, &out, "getTradeRewardRate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTradeRewardRate is a free data retrieval call binding the contract method 0xe8fcc067.
//
// Solidity: function getTradeRewardRate() view returns(uint256)
func (_LuckyVault *LuckyVaultSession) GetTradeRewardRate() (*big.Int, error) {
	return _LuckyVault.Contract.GetTradeRewardRate(&_LuckyVault.CallOpts)
}

// GetTradeRewardRate is a free data retrieval call binding the contract method 0xe8fcc067.
//
// Solidity: function getTradeRewardRate() view returns(uint256)
func (_LuckyVault *LuckyVaultCallerSession) GetTradeRewardRate() (*big.Int, error) {
	return _LuckyVault.Contract.GetTradeRewardRate(&_LuckyVault.CallOpts)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_LuckyVault *LuckyVaultCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _LuckyVault.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_LuckyVault *LuckyVaultSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _LuckyVault.Contract.HasRole(&_LuckyVault.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_LuckyVault *LuckyVaultCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _LuckyVault.Contract.HasRole(&_LuckyVault.CallOpts, role, account)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_LuckyVault *LuckyVaultCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LuckyVault.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_LuckyVault *LuckyVaultSession) Owner() (common.Address, error) {
	return _LuckyVault.Contract.Owner(&_LuckyVault.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_LuckyVault *LuckyVaultCallerSession) Owner() (common.Address, error) {
	return _LuckyVault.Contract.Owner(&_LuckyVault.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_LuckyVault *LuckyVaultCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _LuckyVault.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_LuckyVault *LuckyVaultSession) Paused() (bool, error) {
	return _LuckyVault.Contract.Paused(&_LuckyVault.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_LuckyVault *LuckyVaultCallerSession) Paused() (bool, error) {
	return _LuckyVault.Contract.Paused(&_LuckyVault.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_LuckyVault *LuckyVaultCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _LuckyVault.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_LuckyVault *LuckyVaultSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _LuckyVault.Contract.SupportsInterface(&_LuckyVault.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_LuckyVault *LuckyVaultCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _LuckyVault.Contract.SupportsInterface(&_LuckyVault.CallOpts, interfaceId)
}

// Deposit is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(uint256 amount) returns()
func (_LuckyVault *LuckyVaultTransactor) Deposit(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _LuckyVault.contract.Transact(opts, "deposit", amount)
}

// Deposit is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(uint256 amount) returns()
func (_LuckyVault *LuckyVaultSession) Deposit(amount *big.Int) (*types.Transaction, error) {
	return _LuckyVault.Contract.Deposit(&_LuckyVault.TransactOpts, amount)
}

// Deposit is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(uint256 amount) returns()
func (_LuckyVault *LuckyVaultTransactorSession) Deposit(amount *big.Int) (*types.Transaction, error) {
	return _LuckyVault.Contract.Deposit(&_LuckyVault.TransactOpts, amount)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_LuckyVault *LuckyVaultTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _LuckyVault.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_LuckyVault *LuckyVaultSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _LuckyVault.Contract.GrantRole(&_LuckyVault.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_LuckyVault *LuckyVaultTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _LuckyVault.Contract.GrantRole(&_LuckyVault.TransactOpts, role, account)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_LuckyVault *LuckyVaultTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LuckyVault.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_LuckyVault *LuckyVaultSession) Pause() (*types.Transaction, error) {
	return _LuckyVault.Contract.Pause(&_LuckyVault.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_LuckyVault *LuckyVaultTransactorSession) Pause() (*types.Transaction, error) {
	return _LuckyVault.Contract.Pause(&_LuckyVault.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_LuckyVault *LuckyVaultTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LuckyVault.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_LuckyVault *LuckyVaultSession) RenounceOwnership() (*types.Transaction, error) {
	return _LuckyVault.Contract.RenounceOwnership(&_LuckyVault.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_LuckyVault *LuckyVaultTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _LuckyVault.Contract.RenounceOwnership(&_LuckyVault.TransactOpts)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_LuckyVault *LuckyVaultTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _LuckyVault.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_LuckyVault *LuckyVaultSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _LuckyVault.Contract.RenounceRole(&_LuckyVault.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_LuckyVault *LuckyVaultTransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _LuckyVault.Contract.RenounceRole(&_LuckyVault.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_LuckyVault *LuckyVaultTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _LuckyVault.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_LuckyVault *LuckyVaultSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _LuckyVault.Contract.RevokeRole(&_LuckyVault.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_LuckyVault *LuckyVaultTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _LuckyVault.Contract.RevokeRole(&_LuckyVault.TransactOpts, role, account)
}

// Trade is a paid mutator transaction binding the contract method 0xdf1dd826.
//
// Solidity: function trade(uint256 amount) returns()
func (_LuckyVault *LuckyVaultTransactor) Trade(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _LuckyVault.contract.Transact(opts, "trade", amount)
}

// Trade is a paid mutator transaction binding the contract method 0xdf1dd826.
//
// Solidity: function trade(uint256 amount) returns()
func (_LuckyVault *LuckyVaultSession) Trade(amount *big.Int) (*types.Transaction, error) {
	return _LuckyVault.Contract.Trade(&_LuckyVault.TransactOpts, amount)
}

// Trade is a paid mutator transaction binding the contract method 0xdf1dd826.
//
// Solidity: function trade(uint256 amount) returns()
func (_LuckyVault *LuckyVaultTransactorSession) Trade(amount *big.Int) (*types.Transaction, error) {
	return _LuckyVault.Contract.Trade(&_LuckyVault.TransactOpts, amount)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_LuckyVault *LuckyVaultTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _LuckyVault.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_LuckyVault *LuckyVaultSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _LuckyVault.Contract.TransferOwnership(&_LuckyVault.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_LuckyVault *LuckyVaultTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _LuckyVault.Contract.TransferOwnership(&_LuckyVault.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_LuckyVault *LuckyVaultTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LuckyVault.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_LuckyVault *LuckyVaultSession) Unpause() (*types.Transaction, error) {
	return _LuckyVault.Contract.Unpause(&_LuckyVault.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_LuckyVault *LuckyVaultTransactorSession) Unpause() (*types.Transaction, error) {
	return _LuckyVault.Contract.Unpause(&_LuckyVault.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 amount) returns()
func (_LuckyVault *LuckyVaultTransactor) Withdraw(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _LuckyVault.contract.Transact(opts, "withdraw", amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 amount) returns()
func (_LuckyVault *LuckyVaultSession) Withdraw(amount *big.Int) (*types.Transaction, error) {
	return _LuckyVault.Contract.Withdraw(&_LuckyVault.TransactOpts, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 amount) returns()
func (_LuckyVault *LuckyVaultTransactorSession) Withdraw(amount *big.Int) (*types.Transaction, error) {
	return _LuckyVault.Contract.Withdraw(&_LuckyVault.TransactOpts, amount)
}

// WithdrawByAdmin is a paid mutator transaction binding the contract method 0x487c3580.
//
// Solidity: function withdrawByAdmin(address receiver, uint256 amount) returns()
func (_LuckyVault *LuckyVaultTransactor) WithdrawByAdmin(opts *bind.TransactOpts, receiver common.Address, amount *big.Int) (*types.Transaction, error) {
	return _LuckyVault.contract.Transact(opts, "withdrawByAdmin", receiver, amount)
}

// WithdrawByAdmin is a paid mutator transaction binding the contract method 0x487c3580.
//
// Solidity: function withdrawByAdmin(address receiver, uint256 amount) returns()
func (_LuckyVault *LuckyVaultSession) WithdrawByAdmin(receiver common.Address, amount *big.Int) (*types.Transaction, error) {
	return _LuckyVault.Contract.WithdrawByAdmin(&_LuckyVault.TransactOpts, receiver, amount)
}

// WithdrawByAdmin is a paid mutator transaction binding the contract method 0x487c3580.
//
// Solidity: function withdrawByAdmin(address receiver, uint256 amount) returns()
func (_LuckyVault *LuckyVaultTransactorSession) WithdrawByAdmin(receiver common.Address, amount *big.Int) (*types.Transaction, error) {
	return _LuckyVault.Contract.WithdrawByAdmin(&_LuckyVault.TransactOpts, receiver, amount)
}

// LuckyVaultAdminWithdrawSuccessfulIterator is returned from FilterAdminWithdrawSuccessful and is used to iterate over the raw logs and unpacked data for AdminWithdrawSuccessful events raised by the LuckyVault contract.
type LuckyVaultAdminWithdrawSuccessfulIterator struct {
	Event *LuckyVaultAdminWithdrawSuccessful // Event containing the contract specifics and raw log

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
func (it *LuckyVaultAdminWithdrawSuccessfulIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LuckyVaultAdminWithdrawSuccessful)
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
		it.Event = new(LuckyVaultAdminWithdrawSuccessful)
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
func (it *LuckyVaultAdminWithdrawSuccessfulIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LuckyVaultAdminWithdrawSuccessfulIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LuckyVaultAdminWithdrawSuccessful represents a AdminWithdrawSuccessful event raised by the LuckyVault contract.
type LuckyVaultAdminWithdrawSuccessful struct {
	Sender   common.Address
	Receiver common.Address
	Amount   *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterAdminWithdrawSuccessful is a free log retrieval operation binding the contract event 0x5f1e6e75f939066676897a7e8731d3286ac6abab221d260194e7ea956e2e2eed.
//
// Solidity: event AdminWithdrawSuccessful(address sender, address receiver, uint256 amount)
func (_LuckyVault *LuckyVaultFilterer) FilterAdminWithdrawSuccessful(opts *bind.FilterOpts) (*LuckyVaultAdminWithdrawSuccessfulIterator, error) {

	logs, sub, err := _LuckyVault.contract.FilterLogs(opts, "AdminWithdrawSuccessful")
	if err != nil {
		return nil, err
	}
	return &LuckyVaultAdminWithdrawSuccessfulIterator{contract: _LuckyVault.contract, event: "AdminWithdrawSuccessful", logs: logs, sub: sub}, nil
}

// WatchAdminWithdrawSuccessful is a free log subscription operation binding the contract event 0x5f1e6e75f939066676897a7e8731d3286ac6abab221d260194e7ea956e2e2eed.
//
// Solidity: event AdminWithdrawSuccessful(address sender, address receiver, uint256 amount)
func (_LuckyVault *LuckyVaultFilterer) WatchAdminWithdrawSuccessful(opts *bind.WatchOpts, sink chan<- *LuckyVaultAdminWithdrawSuccessful) (event.Subscription, error) {

	logs, sub, err := _LuckyVault.contract.WatchLogs(opts, "AdminWithdrawSuccessful")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LuckyVaultAdminWithdrawSuccessful)
				if err := _LuckyVault.contract.UnpackLog(event, "AdminWithdrawSuccessful", log); err != nil {
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

// ParseAdminWithdrawSuccessful is a log parse operation binding the contract event 0x5f1e6e75f939066676897a7e8731d3286ac6abab221d260194e7ea956e2e2eed.
//
// Solidity: event AdminWithdrawSuccessful(address sender, address receiver, uint256 amount)
func (_LuckyVault *LuckyVaultFilterer) ParseAdminWithdrawSuccessful(log types.Log) (*LuckyVaultAdminWithdrawSuccessful, error) {
	event := new(LuckyVaultAdminWithdrawSuccessful)
	if err := _LuckyVault.contract.UnpackLog(event, "AdminWithdrawSuccessful", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LuckyVaultDepositSuccessfulIterator is returned from FilterDepositSuccessful and is used to iterate over the raw logs and unpacked data for DepositSuccessful events raised by the LuckyVault contract.
type LuckyVaultDepositSuccessfulIterator struct {
	Event *LuckyVaultDepositSuccessful // Event containing the contract specifics and raw log

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
func (it *LuckyVaultDepositSuccessfulIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LuckyVaultDepositSuccessful)
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
		it.Event = new(LuckyVaultDepositSuccessful)
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
func (it *LuckyVaultDepositSuccessfulIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LuckyVaultDepositSuccessfulIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LuckyVaultDepositSuccessful represents a DepositSuccessful event raised by the LuckyVault contract.
type LuckyVaultDepositSuccessful struct {
	Sender       common.Address
	Amount       *big.Int
	RewardAmount *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterDepositSuccessful is a free log retrieval operation binding the contract event 0x61bb8b82654398e0bdd4e0d489b5e9104b6840211fd89f7264264c96589051fd.
//
// Solidity: event DepositSuccessful(address sender, uint256 amount, uint256 rewardAmount)
func (_LuckyVault *LuckyVaultFilterer) FilterDepositSuccessful(opts *bind.FilterOpts) (*LuckyVaultDepositSuccessfulIterator, error) {

	logs, sub, err := _LuckyVault.contract.FilterLogs(opts, "DepositSuccessful")
	if err != nil {
		return nil, err
	}
	return &LuckyVaultDepositSuccessfulIterator{contract: _LuckyVault.contract, event: "DepositSuccessful", logs: logs, sub: sub}, nil
}

// WatchDepositSuccessful is a free log subscription operation binding the contract event 0x61bb8b82654398e0bdd4e0d489b5e9104b6840211fd89f7264264c96589051fd.
//
// Solidity: event DepositSuccessful(address sender, uint256 amount, uint256 rewardAmount)
func (_LuckyVault *LuckyVaultFilterer) WatchDepositSuccessful(opts *bind.WatchOpts, sink chan<- *LuckyVaultDepositSuccessful) (event.Subscription, error) {

	logs, sub, err := _LuckyVault.contract.WatchLogs(opts, "DepositSuccessful")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LuckyVaultDepositSuccessful)
				if err := _LuckyVault.contract.UnpackLog(event, "DepositSuccessful", log); err != nil {
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

// ParseDepositSuccessful is a log parse operation binding the contract event 0x61bb8b82654398e0bdd4e0d489b5e9104b6840211fd89f7264264c96589051fd.
//
// Solidity: event DepositSuccessful(address sender, uint256 amount, uint256 rewardAmount)
func (_LuckyVault *LuckyVaultFilterer) ParseDepositSuccessful(log types.Log) (*LuckyVaultDepositSuccessful, error) {
	event := new(LuckyVaultDepositSuccessful)
	if err := _LuckyVault.contract.UnpackLog(event, "DepositSuccessful", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LuckyVaultOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the LuckyVault contract.
type LuckyVaultOwnershipTransferredIterator struct {
	Event *LuckyVaultOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *LuckyVaultOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LuckyVaultOwnershipTransferred)
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
		it.Event = new(LuckyVaultOwnershipTransferred)
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
func (it *LuckyVaultOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LuckyVaultOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LuckyVaultOwnershipTransferred represents a OwnershipTransferred event raised by the LuckyVault contract.
type LuckyVaultOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_LuckyVault *LuckyVaultFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*LuckyVaultOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _LuckyVault.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &LuckyVaultOwnershipTransferredIterator{contract: _LuckyVault.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_LuckyVault *LuckyVaultFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *LuckyVaultOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _LuckyVault.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LuckyVaultOwnershipTransferred)
				if err := _LuckyVault.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_LuckyVault *LuckyVaultFilterer) ParseOwnershipTransferred(log types.Log) (*LuckyVaultOwnershipTransferred, error) {
	event := new(LuckyVaultOwnershipTransferred)
	if err := _LuckyVault.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LuckyVaultPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the LuckyVault contract.
type LuckyVaultPausedIterator struct {
	Event *LuckyVaultPaused // Event containing the contract specifics and raw log

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
func (it *LuckyVaultPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LuckyVaultPaused)
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
		it.Event = new(LuckyVaultPaused)
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
func (it *LuckyVaultPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LuckyVaultPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LuckyVaultPaused represents a Paused event raised by the LuckyVault contract.
type LuckyVaultPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_LuckyVault *LuckyVaultFilterer) FilterPaused(opts *bind.FilterOpts) (*LuckyVaultPausedIterator, error) {

	logs, sub, err := _LuckyVault.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &LuckyVaultPausedIterator{contract: _LuckyVault.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_LuckyVault *LuckyVaultFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *LuckyVaultPaused) (event.Subscription, error) {

	logs, sub, err := _LuckyVault.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LuckyVaultPaused)
				if err := _LuckyVault.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_LuckyVault *LuckyVaultFilterer) ParsePaused(log types.Log) (*LuckyVaultPaused, error) {
	event := new(LuckyVaultPaused)
	if err := _LuckyVault.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LuckyVaultRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the LuckyVault contract.
type LuckyVaultRoleAdminChangedIterator struct {
	Event *LuckyVaultRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *LuckyVaultRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LuckyVaultRoleAdminChanged)
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
		it.Event = new(LuckyVaultRoleAdminChanged)
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
func (it *LuckyVaultRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LuckyVaultRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LuckyVaultRoleAdminChanged represents a RoleAdminChanged event raised by the LuckyVault contract.
type LuckyVaultRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_LuckyVault *LuckyVaultFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*LuckyVaultRoleAdminChangedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _LuckyVault.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &LuckyVaultRoleAdminChangedIterator{contract: _LuckyVault.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_LuckyVault *LuckyVaultFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *LuckyVaultRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _LuckyVault.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LuckyVaultRoleAdminChanged)
				if err := _LuckyVault.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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

// ParseRoleAdminChanged is a log parse operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_LuckyVault *LuckyVaultFilterer) ParseRoleAdminChanged(log types.Log) (*LuckyVaultRoleAdminChanged, error) {
	event := new(LuckyVaultRoleAdminChanged)
	if err := _LuckyVault.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LuckyVaultRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the LuckyVault contract.
type LuckyVaultRoleGrantedIterator struct {
	Event *LuckyVaultRoleGranted // Event containing the contract specifics and raw log

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
func (it *LuckyVaultRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LuckyVaultRoleGranted)
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
		it.Event = new(LuckyVaultRoleGranted)
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
func (it *LuckyVaultRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LuckyVaultRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LuckyVaultRoleGranted represents a RoleGranted event raised by the LuckyVault contract.
type LuckyVaultRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_LuckyVault *LuckyVaultFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*LuckyVaultRoleGrantedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _LuckyVault.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &LuckyVaultRoleGrantedIterator{contract: _LuckyVault.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_LuckyVault *LuckyVaultFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *LuckyVaultRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _LuckyVault.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LuckyVaultRoleGranted)
				if err := _LuckyVault.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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

// ParseRoleGranted is a log parse operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_LuckyVault *LuckyVaultFilterer) ParseRoleGranted(log types.Log) (*LuckyVaultRoleGranted, error) {
	event := new(LuckyVaultRoleGranted)
	if err := _LuckyVault.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LuckyVaultRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the LuckyVault contract.
type LuckyVaultRoleRevokedIterator struct {
	Event *LuckyVaultRoleRevoked // Event containing the contract specifics and raw log

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
func (it *LuckyVaultRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LuckyVaultRoleRevoked)
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
		it.Event = new(LuckyVaultRoleRevoked)
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
func (it *LuckyVaultRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LuckyVaultRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LuckyVaultRoleRevoked represents a RoleRevoked event raised by the LuckyVault contract.
type LuckyVaultRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_LuckyVault *LuckyVaultFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*LuckyVaultRoleRevokedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _LuckyVault.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &LuckyVaultRoleRevokedIterator{contract: _LuckyVault.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_LuckyVault *LuckyVaultFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *LuckyVaultRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _LuckyVault.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LuckyVaultRoleRevoked)
				if err := _LuckyVault.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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

// ParseRoleRevoked is a log parse operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_LuckyVault *LuckyVaultFilterer) ParseRoleRevoked(log types.Log) (*LuckyVaultRoleRevoked, error) {
	event := new(LuckyVaultRoleRevoked)
	if err := _LuckyVault.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LuckyVaultTradeSuccessfulIterator is returned from FilterTradeSuccessful and is used to iterate over the raw logs and unpacked data for TradeSuccessful events raised by the LuckyVault contract.
type LuckyVaultTradeSuccessfulIterator struct {
	Event *LuckyVaultTradeSuccessful // Event containing the contract specifics and raw log

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
func (it *LuckyVaultTradeSuccessfulIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LuckyVaultTradeSuccessful)
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
		it.Event = new(LuckyVaultTradeSuccessful)
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
func (it *LuckyVaultTradeSuccessfulIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LuckyVaultTradeSuccessfulIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LuckyVaultTradeSuccessful represents a TradeSuccessful event raised by the LuckyVault contract.
type LuckyVaultTradeSuccessful struct {
	Sender       common.Address
	Amount       *big.Int
	TradeAmount  *big.Int
	RewardAmount *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterTradeSuccessful is a free log retrieval operation binding the contract event 0x333d31babb882d300a2175cd3cb272cf6c099241021c22af2699242f92f4ff4f.
//
// Solidity: event TradeSuccessful(address sender, uint256 amount, uint256 tradeAmount, uint256 rewardAmount)
func (_LuckyVault *LuckyVaultFilterer) FilterTradeSuccessful(opts *bind.FilterOpts) (*LuckyVaultTradeSuccessfulIterator, error) {

	logs, sub, err := _LuckyVault.contract.FilterLogs(opts, "TradeSuccessful")
	if err != nil {
		return nil, err
	}
	return &LuckyVaultTradeSuccessfulIterator{contract: _LuckyVault.contract, event: "TradeSuccessful", logs: logs, sub: sub}, nil
}

// WatchTradeSuccessful is a free log subscription operation binding the contract event 0x333d31babb882d300a2175cd3cb272cf6c099241021c22af2699242f92f4ff4f.
//
// Solidity: event TradeSuccessful(address sender, uint256 amount, uint256 tradeAmount, uint256 rewardAmount)
func (_LuckyVault *LuckyVaultFilterer) WatchTradeSuccessful(opts *bind.WatchOpts, sink chan<- *LuckyVaultTradeSuccessful) (event.Subscription, error) {

	logs, sub, err := _LuckyVault.contract.WatchLogs(opts, "TradeSuccessful")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LuckyVaultTradeSuccessful)
				if err := _LuckyVault.contract.UnpackLog(event, "TradeSuccessful", log); err != nil {
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

// ParseTradeSuccessful is a log parse operation binding the contract event 0x333d31babb882d300a2175cd3cb272cf6c099241021c22af2699242f92f4ff4f.
//
// Solidity: event TradeSuccessful(address sender, uint256 amount, uint256 tradeAmount, uint256 rewardAmount)
func (_LuckyVault *LuckyVaultFilterer) ParseTradeSuccessful(log types.Log) (*LuckyVaultTradeSuccessful, error) {
	event := new(LuckyVaultTradeSuccessful)
	if err := _LuckyVault.contract.UnpackLog(event, "TradeSuccessful", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LuckyVaultUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the LuckyVault contract.
type LuckyVaultUnpausedIterator struct {
	Event *LuckyVaultUnpaused // Event containing the contract specifics and raw log

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
func (it *LuckyVaultUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LuckyVaultUnpaused)
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
		it.Event = new(LuckyVaultUnpaused)
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
func (it *LuckyVaultUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LuckyVaultUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LuckyVaultUnpaused represents a Unpaused event raised by the LuckyVault contract.
type LuckyVaultUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_LuckyVault *LuckyVaultFilterer) FilterUnpaused(opts *bind.FilterOpts) (*LuckyVaultUnpausedIterator, error) {

	logs, sub, err := _LuckyVault.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &LuckyVaultUnpausedIterator{contract: _LuckyVault.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_LuckyVault *LuckyVaultFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *LuckyVaultUnpaused) (event.Subscription, error) {

	logs, sub, err := _LuckyVault.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LuckyVaultUnpaused)
				if err := _LuckyVault.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
func (_LuckyVault *LuckyVaultFilterer) ParseUnpaused(log types.Log) (*LuckyVaultUnpaused, error) {
	event := new(LuckyVaultUnpaused)
	if err := _LuckyVault.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LuckyVaultWithdrawSuccessfulIterator is returned from FilterWithdrawSuccessful and is used to iterate over the raw logs and unpacked data for WithdrawSuccessful events raised by the LuckyVault contract.
type LuckyVaultWithdrawSuccessfulIterator struct {
	Event *LuckyVaultWithdrawSuccessful // Event containing the contract specifics and raw log

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
func (it *LuckyVaultWithdrawSuccessfulIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LuckyVaultWithdrawSuccessful)
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
		it.Event = new(LuckyVaultWithdrawSuccessful)
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
func (it *LuckyVaultWithdrawSuccessfulIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LuckyVaultWithdrawSuccessfulIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LuckyVaultWithdrawSuccessful represents a WithdrawSuccessful event raised by the LuckyVault contract.
type LuckyVaultWithdrawSuccessful struct {
	Sender common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWithdrawSuccessful is a free log retrieval operation binding the contract event 0x84e742a3ea03623687a5febaa797636d728230721d4b3d1405d19fbbc16298d9.
//
// Solidity: event WithdrawSuccessful(address sender, uint256 amount)
func (_LuckyVault *LuckyVaultFilterer) FilterWithdrawSuccessful(opts *bind.FilterOpts) (*LuckyVaultWithdrawSuccessfulIterator, error) {

	logs, sub, err := _LuckyVault.contract.FilterLogs(opts, "WithdrawSuccessful")
	if err != nil {
		return nil, err
	}
	return &LuckyVaultWithdrawSuccessfulIterator{contract: _LuckyVault.contract, event: "WithdrawSuccessful", logs: logs, sub: sub}, nil
}

// WatchWithdrawSuccessful is a free log subscription operation binding the contract event 0x84e742a3ea03623687a5febaa797636d728230721d4b3d1405d19fbbc16298d9.
//
// Solidity: event WithdrawSuccessful(address sender, uint256 amount)
func (_LuckyVault *LuckyVaultFilterer) WatchWithdrawSuccessful(opts *bind.WatchOpts, sink chan<- *LuckyVaultWithdrawSuccessful) (event.Subscription, error) {

	logs, sub, err := _LuckyVault.contract.WatchLogs(opts, "WithdrawSuccessful")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LuckyVaultWithdrawSuccessful)
				if err := _LuckyVault.contract.UnpackLog(event, "WithdrawSuccessful", log); err != nil {
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

// ParseWithdrawSuccessful is a log parse operation binding the contract event 0x84e742a3ea03623687a5febaa797636d728230721d4b3d1405d19fbbc16298d9.
//
// Solidity: event WithdrawSuccessful(address sender, uint256 amount)
func (_LuckyVault *LuckyVaultFilterer) ParseWithdrawSuccessful(log types.Log) (*LuckyVaultWithdrawSuccessful, error) {
	event := new(LuckyVaultWithdrawSuccessful)
	if err := _LuckyVault.contract.UnpackLog(event, "WithdrawSuccessful", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
