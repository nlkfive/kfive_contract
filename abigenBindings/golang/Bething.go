// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package smc

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

// SmcMetaData contains all meta data concerning the Smc contract.
var SmcMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"acceptedToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"raceList\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AlreadyReceived\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BetFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FundFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidContract\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidSlot\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RaceNotExisted\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"}],\"name\":\"TooEarly\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"}],\"name\":\"TooLate\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"acceptToken\",\"type\":\"address\"}],\"name\":\"AcceptTokenUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"slotId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"raceId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"betValue\",\"type\":\"uint256\"}],\"name\":\"BetSuccessful\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"raceId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"claimValue\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"ClaimCommission\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"raceId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"claimValue\",\"type\":\"uint256\"}],\"name\":\"ClaimSuccessful\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"raceId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fundValue\",\"type\":\"uint256\"}],\"name\":\"FundSuccessful\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"race\",\"type\":\"address\"}],\"name\":\"RaceListUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PAUSER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_acceptedToken\",\"outputs\":[{\"internalType\":\"contractIBEP20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_raceList\",\"outputs\":[{\"internalType\":\"contractIRaceList\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getRoleMember\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleMemberCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"slotId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"raceId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"betValue\",\"type\":\"uint256\"}],\"name\":\"bet\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"raceId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"fundValue\",\"type\":\"uint256\"}],\"name\":\"fundRace\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"raceId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"slotId\",\"type\":\"uint256\"}],\"name\":\"claim\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"raceId\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"claimCommission\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"raceId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"slotId\",\"type\":\"uint256\"}],\"name\":\"totalSlotBet\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"raceId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"slotId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"userSlotBet\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"raceId\",\"type\":\"bytes32\"}],\"name\":\"totalRaceBet\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"raceId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"slotId\",\"type\":\"uint256\"}],\"name\":\"getSlotPosition\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"raceList\",\"type\":\"address\"}],\"name\":\"updateRaceAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"acceptToken\",\"type\":\"address\"}],\"name\":\"updateAcceptTokenAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"_to\",\"type\":\"address\"}],\"name\":\"destroySmartContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// SmcABI is the input ABI used to generate the binding from.
// Deprecated: Use SmcMetaData.ABI instead.
var SmcABI = SmcMetaData.ABI

// Smc is an auto generated Go binding around an Ethereum contract.
type Smc struct {
	SmcCaller     // Read-only binding to the contract
	SmcTransactor // Write-only binding to the contract
	SmcFilterer   // Log filterer for contract events
}

// SmcCaller is an auto generated read-only Go binding around an Ethereum contract.
type SmcCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SmcTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SmcTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SmcFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SmcFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SmcSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SmcSession struct {
	Contract     *Smc              // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SmcCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SmcCallerSession struct {
	Contract *SmcCaller    // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// SmcTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SmcTransactorSession struct {
	Contract     *SmcTransactor    // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SmcRaw is an auto generated low-level Go binding around an Ethereum contract.
type SmcRaw struct {
	Contract *Smc // Generic contract binding to access the raw methods on
}

// SmcCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SmcCallerRaw struct {
	Contract *SmcCaller // Generic read-only contract binding to access the raw methods on
}

// SmcTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SmcTransactorRaw struct {
	Contract *SmcTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSmc creates a new instance of Smc, bound to a specific deployed contract.
func NewSmc(address common.Address, backend bind.ContractBackend) (*Smc, error) {
	contract, err := bindSmc(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Smc{SmcCaller: SmcCaller{contract: contract}, SmcTransactor: SmcTransactor{contract: contract}, SmcFilterer: SmcFilterer{contract: contract}}, nil
}

// NewSmcCaller creates a new read-only instance of Smc, bound to a specific deployed contract.
func NewSmcCaller(address common.Address, caller bind.ContractCaller) (*SmcCaller, error) {
	contract, err := bindSmc(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SmcCaller{contract: contract}, nil
}

// NewSmcTransactor creates a new write-only instance of Smc, bound to a specific deployed contract.
func NewSmcTransactor(address common.Address, transactor bind.ContractTransactor) (*SmcTransactor, error) {
	contract, err := bindSmc(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SmcTransactor{contract: contract}, nil
}

// NewSmcFilterer creates a new log filterer instance of Smc, bound to a specific deployed contract.
func NewSmcFilterer(address common.Address, filterer bind.ContractFilterer) (*SmcFilterer, error) {
	contract, err := bindSmc(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SmcFilterer{contract: contract}, nil
}

// bindSmc binds a generic wrapper to an already deployed contract.
func bindSmc(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SmcABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Smc *SmcRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Smc.Contract.SmcCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Smc *SmcRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Smc.Contract.SmcTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Smc *SmcRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Smc.Contract.SmcTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Smc *SmcCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Smc.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Smc *SmcTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Smc.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Smc *SmcTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Smc.Contract.contract.Transact(opts, method, params...)
}

// ADMINROLE is a free data retrieval call binding the contract method 0x75b238fc.
//
// Solidity: function ADMIN_ROLE() view returns(bytes32)
func (_Smc *SmcCaller) ADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Smc.contract.Call(opts, &out, "ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ADMINROLE is a free data retrieval call binding the contract method 0x75b238fc.
//
// Solidity: function ADMIN_ROLE() view returns(bytes32)
func (_Smc *SmcSession) ADMINROLE() ([32]byte, error) {
	return _Smc.Contract.ADMINROLE(&_Smc.CallOpts)
}

// ADMINROLE is a free data retrieval call binding the contract method 0x75b238fc.
//
// Solidity: function ADMIN_ROLE() view returns(bytes32)
func (_Smc *SmcCallerSession) ADMINROLE() ([32]byte, error) {
	return _Smc.Contract.ADMINROLE(&_Smc.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Smc *SmcCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Smc.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Smc *SmcSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Smc.Contract.DEFAULTADMINROLE(&_Smc.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Smc *SmcCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Smc.Contract.DEFAULTADMINROLE(&_Smc.CallOpts)
}

// PAUSERROLE is a free data retrieval call binding the contract method 0xe63ab1e9.
//
// Solidity: function PAUSER_ROLE() view returns(bytes32)
func (_Smc *SmcCaller) PAUSERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Smc.contract.Call(opts, &out, "PAUSER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// PAUSERROLE is a free data retrieval call binding the contract method 0xe63ab1e9.
//
// Solidity: function PAUSER_ROLE() view returns(bytes32)
func (_Smc *SmcSession) PAUSERROLE() ([32]byte, error) {
	return _Smc.Contract.PAUSERROLE(&_Smc.CallOpts)
}

// PAUSERROLE is a free data retrieval call binding the contract method 0xe63ab1e9.
//
// Solidity: function PAUSER_ROLE() view returns(bytes32)
func (_Smc *SmcCallerSession) PAUSERROLE() ([32]byte, error) {
	return _Smc.Contract.PAUSERROLE(&_Smc.CallOpts)
}

// AcceptedToken is a free data retrieval call binding the contract method 0x4e5b3c7c.
//
// Solidity: function _acceptedToken() view returns(address)
func (_Smc *SmcCaller) AcceptedToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Smc.contract.Call(opts, &out, "_acceptedToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AcceptedToken is a free data retrieval call binding the contract method 0x4e5b3c7c.
//
// Solidity: function _acceptedToken() view returns(address)
func (_Smc *SmcSession) AcceptedToken() (common.Address, error) {
	return _Smc.Contract.AcceptedToken(&_Smc.CallOpts)
}

// AcceptedToken is a free data retrieval call binding the contract method 0x4e5b3c7c.
//
// Solidity: function _acceptedToken() view returns(address)
func (_Smc *SmcCallerSession) AcceptedToken() (common.Address, error) {
	return _Smc.Contract.AcceptedToken(&_Smc.CallOpts)
}

// RaceList is a free data retrieval call binding the contract method 0x6c76f620.
//
// Solidity: function _raceList() view returns(address)
func (_Smc *SmcCaller) RaceList(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Smc.contract.Call(opts, &out, "_raceList")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RaceList is a free data retrieval call binding the contract method 0x6c76f620.
//
// Solidity: function _raceList() view returns(address)
func (_Smc *SmcSession) RaceList() (common.Address, error) {
	return _Smc.Contract.RaceList(&_Smc.CallOpts)
}

// RaceList is a free data retrieval call binding the contract method 0x6c76f620.
//
// Solidity: function _raceList() view returns(address)
func (_Smc *SmcCallerSession) RaceList() (common.Address, error) {
	return _Smc.Contract.RaceList(&_Smc.CallOpts)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Smc *SmcCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _Smc.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Smc *SmcSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Smc.Contract.GetRoleAdmin(&_Smc.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Smc *SmcCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Smc.Contract.GetRoleAdmin(&_Smc.CallOpts, role)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_Smc *SmcCaller) GetRoleMember(opts *bind.CallOpts, role [32]byte, index *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Smc.contract.Call(opts, &out, "getRoleMember", role, index)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_Smc *SmcSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _Smc.Contract.GetRoleMember(&_Smc.CallOpts, role, index)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_Smc *SmcCallerSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _Smc.Contract.GetRoleMember(&_Smc.CallOpts, role, index)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_Smc *SmcCaller) GetRoleMemberCount(opts *bind.CallOpts, role [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _Smc.contract.Call(opts, &out, "getRoleMemberCount", role)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_Smc *SmcSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _Smc.Contract.GetRoleMemberCount(&_Smc.CallOpts, role)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_Smc *SmcCallerSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _Smc.Contract.GetRoleMemberCount(&_Smc.CallOpts, role)
}

// GetSlotPosition is a free data retrieval call binding the contract method 0xd0b57c69.
//
// Solidity: function getSlotPosition(bytes32 raceId, uint256 slotId) view returns(uint8)
func (_Smc *SmcCaller) GetSlotPosition(opts *bind.CallOpts, raceId [32]byte, slotId *big.Int) (uint8, error) {
	var out []interface{}
	err := _Smc.contract.Call(opts, &out, "getSlotPosition", raceId, slotId)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// GetSlotPosition is a free data retrieval call binding the contract method 0xd0b57c69.
//
// Solidity: function getSlotPosition(bytes32 raceId, uint256 slotId) view returns(uint8)
func (_Smc *SmcSession) GetSlotPosition(raceId [32]byte, slotId *big.Int) (uint8, error) {
	return _Smc.Contract.GetSlotPosition(&_Smc.CallOpts, raceId, slotId)
}

// GetSlotPosition is a free data retrieval call binding the contract method 0xd0b57c69.
//
// Solidity: function getSlotPosition(bytes32 raceId, uint256 slotId) view returns(uint8)
func (_Smc *SmcCallerSession) GetSlotPosition(raceId [32]byte, slotId *big.Int) (uint8, error) {
	return _Smc.Contract.GetSlotPosition(&_Smc.CallOpts, raceId, slotId)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Smc *SmcCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _Smc.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Smc *SmcSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Smc.Contract.HasRole(&_Smc.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Smc *SmcCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Smc.Contract.HasRole(&_Smc.CallOpts, role, account)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Smc *SmcCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Smc.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Smc *SmcSession) Owner() (common.Address, error) {
	return _Smc.Contract.Owner(&_Smc.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Smc *SmcCallerSession) Owner() (common.Address, error) {
	return _Smc.Contract.Owner(&_Smc.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Smc *SmcCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Smc.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Smc *SmcSession) Paused() (bool, error) {
	return _Smc.Contract.Paused(&_Smc.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Smc *SmcCallerSession) Paused() (bool, error) {
	return _Smc.Contract.Paused(&_Smc.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Smc *SmcCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Smc.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Smc *SmcSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Smc.Contract.SupportsInterface(&_Smc.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Smc *SmcCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Smc.Contract.SupportsInterface(&_Smc.CallOpts, interfaceId)
}

// TotalRaceBet is a free data retrieval call binding the contract method 0x2123c15e.
//
// Solidity: function totalRaceBet(bytes32 raceId) view returns(uint256)
func (_Smc *SmcCaller) TotalRaceBet(opts *bind.CallOpts, raceId [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _Smc.contract.Call(opts, &out, "totalRaceBet", raceId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalRaceBet is a free data retrieval call binding the contract method 0x2123c15e.
//
// Solidity: function totalRaceBet(bytes32 raceId) view returns(uint256)
func (_Smc *SmcSession) TotalRaceBet(raceId [32]byte) (*big.Int, error) {
	return _Smc.Contract.TotalRaceBet(&_Smc.CallOpts, raceId)
}

// TotalRaceBet is a free data retrieval call binding the contract method 0x2123c15e.
//
// Solidity: function totalRaceBet(bytes32 raceId) view returns(uint256)
func (_Smc *SmcCallerSession) TotalRaceBet(raceId [32]byte) (*big.Int, error) {
	return _Smc.Contract.TotalRaceBet(&_Smc.CallOpts, raceId)
}

// TotalSlotBet is a free data retrieval call binding the contract method 0xe740582a.
//
// Solidity: function totalSlotBet(bytes32 raceId, uint256 slotId) view returns(uint256)
func (_Smc *SmcCaller) TotalSlotBet(opts *bind.CallOpts, raceId [32]byte, slotId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Smc.contract.Call(opts, &out, "totalSlotBet", raceId, slotId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSlotBet is a free data retrieval call binding the contract method 0xe740582a.
//
// Solidity: function totalSlotBet(bytes32 raceId, uint256 slotId) view returns(uint256)
func (_Smc *SmcSession) TotalSlotBet(raceId [32]byte, slotId *big.Int) (*big.Int, error) {
	return _Smc.Contract.TotalSlotBet(&_Smc.CallOpts, raceId, slotId)
}

// TotalSlotBet is a free data retrieval call binding the contract method 0xe740582a.
//
// Solidity: function totalSlotBet(bytes32 raceId, uint256 slotId) view returns(uint256)
func (_Smc *SmcCallerSession) TotalSlotBet(raceId [32]byte, slotId *big.Int) (*big.Int, error) {
	return _Smc.Contract.TotalSlotBet(&_Smc.CallOpts, raceId, slotId)
}

// UserSlotBet is a free data retrieval call binding the contract method 0x2e279801.
//
// Solidity: function userSlotBet(bytes32 raceId, uint256 slotId, address user) view returns(uint256)
func (_Smc *SmcCaller) UserSlotBet(opts *bind.CallOpts, raceId [32]byte, slotId *big.Int, user common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Smc.contract.Call(opts, &out, "userSlotBet", raceId, slotId, user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UserSlotBet is a free data retrieval call binding the contract method 0x2e279801.
//
// Solidity: function userSlotBet(bytes32 raceId, uint256 slotId, address user) view returns(uint256)
func (_Smc *SmcSession) UserSlotBet(raceId [32]byte, slotId *big.Int, user common.Address) (*big.Int, error) {
	return _Smc.Contract.UserSlotBet(&_Smc.CallOpts, raceId, slotId, user)
}

// UserSlotBet is a free data retrieval call binding the contract method 0x2e279801.
//
// Solidity: function userSlotBet(bytes32 raceId, uint256 slotId, address user) view returns(uint256)
func (_Smc *SmcCallerSession) UserSlotBet(raceId [32]byte, slotId *big.Int, user common.Address) (*big.Int, error) {
	return _Smc.Contract.UserSlotBet(&_Smc.CallOpts, raceId, slotId, user)
}

// Bet is a paid mutator transaction binding the contract method 0x94fc859b.
//
// Solidity: function bet(uint256 slotId, bytes32 raceId, uint256 betValue) returns()
func (_Smc *SmcTransactor) Bet(opts *bind.TransactOpts, slotId *big.Int, raceId [32]byte, betValue *big.Int) (*types.Transaction, error) {
	return _Smc.contract.Transact(opts, "bet", slotId, raceId, betValue)
}

// Bet is a paid mutator transaction binding the contract method 0x94fc859b.
//
// Solidity: function bet(uint256 slotId, bytes32 raceId, uint256 betValue) returns()
func (_Smc *SmcSession) Bet(slotId *big.Int, raceId [32]byte, betValue *big.Int) (*types.Transaction, error) {
	return _Smc.Contract.Bet(&_Smc.TransactOpts, slotId, raceId, betValue)
}

// Bet is a paid mutator transaction binding the contract method 0x94fc859b.
//
// Solidity: function bet(uint256 slotId, bytes32 raceId, uint256 betValue) returns()
func (_Smc *SmcTransactorSession) Bet(slotId *big.Int, raceId [32]byte, betValue *big.Int) (*types.Transaction, error) {
	return _Smc.Contract.Bet(&_Smc.TransactOpts, slotId, raceId, betValue)
}

// Claim is a paid mutator transaction binding the contract method 0x63f44968.
//
// Solidity: function claim(bytes32 raceId, uint256 slotId) returns()
func (_Smc *SmcTransactor) Claim(opts *bind.TransactOpts, raceId [32]byte, slotId *big.Int) (*types.Transaction, error) {
	return _Smc.contract.Transact(opts, "claim", raceId, slotId)
}

// Claim is a paid mutator transaction binding the contract method 0x63f44968.
//
// Solidity: function claim(bytes32 raceId, uint256 slotId) returns()
func (_Smc *SmcSession) Claim(raceId [32]byte, slotId *big.Int) (*types.Transaction, error) {
	return _Smc.Contract.Claim(&_Smc.TransactOpts, raceId, slotId)
}

// Claim is a paid mutator transaction binding the contract method 0x63f44968.
//
// Solidity: function claim(bytes32 raceId, uint256 slotId) returns()
func (_Smc *SmcTransactorSession) Claim(raceId [32]byte, slotId *big.Int) (*types.Transaction, error) {
	return _Smc.Contract.Claim(&_Smc.TransactOpts, raceId, slotId)
}

// ClaimCommission is a paid mutator transaction binding the contract method 0x5cce2b42.
//
// Solidity: function claimCommission(bytes32 raceId, address receiver) returns()
func (_Smc *SmcTransactor) ClaimCommission(opts *bind.TransactOpts, raceId [32]byte, receiver common.Address) (*types.Transaction, error) {
	return _Smc.contract.Transact(opts, "claimCommission", raceId, receiver)
}

// ClaimCommission is a paid mutator transaction binding the contract method 0x5cce2b42.
//
// Solidity: function claimCommission(bytes32 raceId, address receiver) returns()
func (_Smc *SmcSession) ClaimCommission(raceId [32]byte, receiver common.Address) (*types.Transaction, error) {
	return _Smc.Contract.ClaimCommission(&_Smc.TransactOpts, raceId, receiver)
}

// ClaimCommission is a paid mutator transaction binding the contract method 0x5cce2b42.
//
// Solidity: function claimCommission(bytes32 raceId, address receiver) returns()
func (_Smc *SmcTransactorSession) ClaimCommission(raceId [32]byte, receiver common.Address) (*types.Transaction, error) {
	return _Smc.Contract.ClaimCommission(&_Smc.TransactOpts, raceId, receiver)
}

// DestroySmartContract is a paid mutator transaction binding the contract method 0x39df43ff.
//
// Solidity: function destroySmartContract(address _to) returns()
func (_Smc *SmcTransactor) DestroySmartContract(opts *bind.TransactOpts, _to common.Address) (*types.Transaction, error) {
	return _Smc.contract.Transact(opts, "destroySmartContract", _to)
}

// DestroySmartContract is a paid mutator transaction binding the contract method 0x39df43ff.
//
// Solidity: function destroySmartContract(address _to) returns()
func (_Smc *SmcSession) DestroySmartContract(_to common.Address) (*types.Transaction, error) {
	return _Smc.Contract.DestroySmartContract(&_Smc.TransactOpts, _to)
}

// DestroySmartContract is a paid mutator transaction binding the contract method 0x39df43ff.
//
// Solidity: function destroySmartContract(address _to) returns()
func (_Smc *SmcTransactorSession) DestroySmartContract(_to common.Address) (*types.Transaction, error) {
	return _Smc.Contract.DestroySmartContract(&_Smc.TransactOpts, _to)
}

// FundRace is a paid mutator transaction binding the contract method 0x2ebe4493.
//
// Solidity: function fundRace(bytes32 raceId, uint256 fundValue) returns()
func (_Smc *SmcTransactor) FundRace(opts *bind.TransactOpts, raceId [32]byte, fundValue *big.Int) (*types.Transaction, error) {
	return _Smc.contract.Transact(opts, "fundRace", raceId, fundValue)
}

// FundRace is a paid mutator transaction binding the contract method 0x2ebe4493.
//
// Solidity: function fundRace(bytes32 raceId, uint256 fundValue) returns()
func (_Smc *SmcSession) FundRace(raceId [32]byte, fundValue *big.Int) (*types.Transaction, error) {
	return _Smc.Contract.FundRace(&_Smc.TransactOpts, raceId, fundValue)
}

// FundRace is a paid mutator transaction binding the contract method 0x2ebe4493.
//
// Solidity: function fundRace(bytes32 raceId, uint256 fundValue) returns()
func (_Smc *SmcTransactorSession) FundRace(raceId [32]byte, fundValue *big.Int) (*types.Transaction, error) {
	return _Smc.Contract.FundRace(&_Smc.TransactOpts, raceId, fundValue)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Smc *SmcTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Smc.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Smc *SmcSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Smc.Contract.GrantRole(&_Smc.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Smc *SmcTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Smc.Contract.GrantRole(&_Smc.TransactOpts, role, account)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Smc *SmcTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Smc.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Smc *SmcSession) Pause() (*types.Transaction, error) {
	return _Smc.Contract.Pause(&_Smc.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Smc *SmcTransactorSession) Pause() (*types.Transaction, error) {
	return _Smc.Contract.Pause(&_Smc.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Smc *SmcTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Smc.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Smc *SmcSession) RenounceOwnership() (*types.Transaction, error) {
	return _Smc.Contract.RenounceOwnership(&_Smc.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Smc *SmcTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Smc.Contract.RenounceOwnership(&_Smc.TransactOpts)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Smc *SmcTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Smc.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Smc *SmcSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Smc.Contract.RenounceRole(&_Smc.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Smc *SmcTransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Smc.Contract.RenounceRole(&_Smc.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Smc *SmcTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Smc.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Smc *SmcSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Smc.Contract.RevokeRole(&_Smc.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Smc *SmcTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Smc.Contract.RevokeRole(&_Smc.TransactOpts, role, account)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Smc *SmcTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Smc.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Smc *SmcSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Smc.Contract.TransferOwnership(&_Smc.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Smc *SmcTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Smc.Contract.TransferOwnership(&_Smc.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Smc *SmcTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Smc.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Smc *SmcSession) Unpause() (*types.Transaction, error) {
	return _Smc.Contract.Unpause(&_Smc.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Smc *SmcTransactorSession) Unpause() (*types.Transaction, error) {
	return _Smc.Contract.Unpause(&_Smc.TransactOpts)
}

// UpdateAcceptTokenAddress is a paid mutator transaction binding the contract method 0xb4a2fad0.
//
// Solidity: function updateAcceptTokenAddress(address acceptToken) returns()
func (_Smc *SmcTransactor) UpdateAcceptTokenAddress(opts *bind.TransactOpts, acceptToken common.Address) (*types.Transaction, error) {
	return _Smc.contract.Transact(opts, "updateAcceptTokenAddress", acceptToken)
}

// UpdateAcceptTokenAddress is a paid mutator transaction binding the contract method 0xb4a2fad0.
//
// Solidity: function updateAcceptTokenAddress(address acceptToken) returns()
func (_Smc *SmcSession) UpdateAcceptTokenAddress(acceptToken common.Address) (*types.Transaction, error) {
	return _Smc.Contract.UpdateAcceptTokenAddress(&_Smc.TransactOpts, acceptToken)
}

// UpdateAcceptTokenAddress is a paid mutator transaction binding the contract method 0xb4a2fad0.
//
// Solidity: function updateAcceptTokenAddress(address acceptToken) returns()
func (_Smc *SmcTransactorSession) UpdateAcceptTokenAddress(acceptToken common.Address) (*types.Transaction, error) {
	return _Smc.Contract.UpdateAcceptTokenAddress(&_Smc.TransactOpts, acceptToken)
}

// UpdateRaceAddress is a paid mutator transaction binding the contract method 0xd7c49938.
//
// Solidity: function updateRaceAddress(address raceList) returns()
func (_Smc *SmcTransactor) UpdateRaceAddress(opts *bind.TransactOpts, raceList common.Address) (*types.Transaction, error) {
	return _Smc.contract.Transact(opts, "updateRaceAddress", raceList)
}

// UpdateRaceAddress is a paid mutator transaction binding the contract method 0xd7c49938.
//
// Solidity: function updateRaceAddress(address raceList) returns()
func (_Smc *SmcSession) UpdateRaceAddress(raceList common.Address) (*types.Transaction, error) {
	return _Smc.Contract.UpdateRaceAddress(&_Smc.TransactOpts, raceList)
}

// UpdateRaceAddress is a paid mutator transaction binding the contract method 0xd7c49938.
//
// Solidity: function updateRaceAddress(address raceList) returns()
func (_Smc *SmcTransactorSession) UpdateRaceAddress(raceList common.Address) (*types.Transaction, error) {
	return _Smc.Contract.UpdateRaceAddress(&_Smc.TransactOpts, raceList)
}

// SmcAcceptTokenUpdatedIterator is returned from FilterAcceptTokenUpdated and is used to iterate over the raw logs and unpacked data for AcceptTokenUpdated events raised by the Smc contract.
type SmcAcceptTokenUpdatedIterator struct {
	Event *SmcAcceptTokenUpdated // Event containing the contract specifics and raw log

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
func (it *SmcAcceptTokenUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SmcAcceptTokenUpdated)
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
		it.Event = new(SmcAcceptTokenUpdated)
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
func (it *SmcAcceptTokenUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SmcAcceptTokenUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SmcAcceptTokenUpdated represents a AcceptTokenUpdated event raised by the Smc contract.
type SmcAcceptTokenUpdated struct {
	AcceptToken common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterAcceptTokenUpdated is a free log retrieval operation binding the contract event 0x44c3b193081dd50818bc919dc1963ae03762174c9a89af5fbbce54ee71a13f97.
//
// Solidity: event AcceptTokenUpdated(address acceptToken)
func (_Smc *SmcFilterer) FilterAcceptTokenUpdated(opts *bind.FilterOpts) (*SmcAcceptTokenUpdatedIterator, error) {

	logs, sub, err := _Smc.contract.FilterLogs(opts, "AcceptTokenUpdated")
	if err != nil {
		return nil, err
	}
	return &SmcAcceptTokenUpdatedIterator{contract: _Smc.contract, event: "AcceptTokenUpdated", logs: logs, sub: sub}, nil
}

// WatchAcceptTokenUpdated is a free log subscription operation binding the contract event 0x44c3b193081dd50818bc919dc1963ae03762174c9a89af5fbbce54ee71a13f97.
//
// Solidity: event AcceptTokenUpdated(address acceptToken)
func (_Smc *SmcFilterer) WatchAcceptTokenUpdated(opts *bind.WatchOpts, sink chan<- *SmcAcceptTokenUpdated) (event.Subscription, error) {

	logs, sub, err := _Smc.contract.WatchLogs(opts, "AcceptTokenUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SmcAcceptTokenUpdated)
				if err := _Smc.contract.UnpackLog(event, "AcceptTokenUpdated", log); err != nil {
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

// ParseAcceptTokenUpdated is a log parse operation binding the contract event 0x44c3b193081dd50818bc919dc1963ae03762174c9a89af5fbbce54ee71a13f97.
//
// Solidity: event AcceptTokenUpdated(address acceptToken)
func (_Smc *SmcFilterer) ParseAcceptTokenUpdated(log types.Log) (*SmcAcceptTokenUpdated, error) {
	event := new(SmcAcceptTokenUpdated)
	if err := _Smc.contract.UnpackLog(event, "AcceptTokenUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SmcBetSuccessfulIterator is returned from FilterBetSuccessful and is used to iterate over the raw logs and unpacked data for BetSuccessful events raised by the Smc contract.
type SmcBetSuccessfulIterator struct {
	Event *SmcBetSuccessful // Event containing the contract specifics and raw log

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
func (it *SmcBetSuccessfulIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SmcBetSuccessful)
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
		it.Event = new(SmcBetSuccessful)
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
func (it *SmcBetSuccessfulIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SmcBetSuccessfulIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SmcBetSuccessful represents a BetSuccessful event raised by the Smc contract.
type SmcBetSuccessful struct {
	SlotId   *big.Int
	RaceId   [32]byte
	BetValue *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterBetSuccessful is a free log retrieval operation binding the contract event 0x90340911801212a3d95efe0dca9cdbd74549020af8cf33fc0861b74d82fe6dbd.
//
// Solidity: event BetSuccessful(uint256 slotId, bytes32 raceId, uint256 betValue)
func (_Smc *SmcFilterer) FilterBetSuccessful(opts *bind.FilterOpts) (*SmcBetSuccessfulIterator, error) {

	logs, sub, err := _Smc.contract.FilterLogs(opts, "BetSuccessful")
	if err != nil {
		return nil, err
	}
	return &SmcBetSuccessfulIterator{contract: _Smc.contract, event: "BetSuccessful", logs: logs, sub: sub}, nil
}

// WatchBetSuccessful is a free log subscription operation binding the contract event 0x90340911801212a3d95efe0dca9cdbd74549020af8cf33fc0861b74d82fe6dbd.
//
// Solidity: event BetSuccessful(uint256 slotId, bytes32 raceId, uint256 betValue)
func (_Smc *SmcFilterer) WatchBetSuccessful(opts *bind.WatchOpts, sink chan<- *SmcBetSuccessful) (event.Subscription, error) {

	logs, sub, err := _Smc.contract.WatchLogs(opts, "BetSuccessful")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SmcBetSuccessful)
				if err := _Smc.contract.UnpackLog(event, "BetSuccessful", log); err != nil {
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

// ParseBetSuccessful is a log parse operation binding the contract event 0x90340911801212a3d95efe0dca9cdbd74549020af8cf33fc0861b74d82fe6dbd.
//
// Solidity: event BetSuccessful(uint256 slotId, bytes32 raceId, uint256 betValue)
func (_Smc *SmcFilterer) ParseBetSuccessful(log types.Log) (*SmcBetSuccessful, error) {
	event := new(SmcBetSuccessful)
	if err := _Smc.contract.UnpackLog(event, "BetSuccessful", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SmcClaimCommissionIterator is returned from FilterClaimCommission and is used to iterate over the raw logs and unpacked data for ClaimCommission events raised by the Smc contract.
type SmcClaimCommissionIterator struct {
	Event *SmcClaimCommission // Event containing the contract specifics and raw log

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
func (it *SmcClaimCommissionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SmcClaimCommission)
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
		it.Event = new(SmcClaimCommission)
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
func (it *SmcClaimCommissionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SmcClaimCommissionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SmcClaimCommission represents a ClaimCommission event raised by the Smc contract.
type SmcClaimCommission struct {
	RaceId     [32]byte
	ClaimValue *big.Int
	Receiver   common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterClaimCommission is a free log retrieval operation binding the contract event 0x94fe36569bb216f6cb564fee8983d035c7cacba5679023a4ccabe33fdf6d4d88.
//
// Solidity: event ClaimCommission(bytes32 raceId, uint256 claimValue, address receiver)
func (_Smc *SmcFilterer) FilterClaimCommission(opts *bind.FilterOpts) (*SmcClaimCommissionIterator, error) {

	logs, sub, err := _Smc.contract.FilterLogs(opts, "ClaimCommission")
	if err != nil {
		return nil, err
	}
	return &SmcClaimCommissionIterator{contract: _Smc.contract, event: "ClaimCommission", logs: logs, sub: sub}, nil
}

// WatchClaimCommission is a free log subscription operation binding the contract event 0x94fe36569bb216f6cb564fee8983d035c7cacba5679023a4ccabe33fdf6d4d88.
//
// Solidity: event ClaimCommission(bytes32 raceId, uint256 claimValue, address receiver)
func (_Smc *SmcFilterer) WatchClaimCommission(opts *bind.WatchOpts, sink chan<- *SmcClaimCommission) (event.Subscription, error) {

	logs, sub, err := _Smc.contract.WatchLogs(opts, "ClaimCommission")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SmcClaimCommission)
				if err := _Smc.contract.UnpackLog(event, "ClaimCommission", log); err != nil {
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

// ParseClaimCommission is a log parse operation binding the contract event 0x94fe36569bb216f6cb564fee8983d035c7cacba5679023a4ccabe33fdf6d4d88.
//
// Solidity: event ClaimCommission(bytes32 raceId, uint256 claimValue, address receiver)
func (_Smc *SmcFilterer) ParseClaimCommission(log types.Log) (*SmcClaimCommission, error) {
	event := new(SmcClaimCommission)
	if err := _Smc.contract.UnpackLog(event, "ClaimCommission", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SmcClaimSuccessfulIterator is returned from FilterClaimSuccessful and is used to iterate over the raw logs and unpacked data for ClaimSuccessful events raised by the Smc contract.
type SmcClaimSuccessfulIterator struct {
	Event *SmcClaimSuccessful // Event containing the contract specifics and raw log

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
func (it *SmcClaimSuccessfulIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SmcClaimSuccessful)
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
		it.Event = new(SmcClaimSuccessful)
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
func (it *SmcClaimSuccessfulIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SmcClaimSuccessfulIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SmcClaimSuccessful represents a ClaimSuccessful event raised by the Smc contract.
type SmcClaimSuccessful struct {
	RaceId     [32]byte
	ClaimValue *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterClaimSuccessful is a free log retrieval operation binding the contract event 0x78d224c5b360f08891a0d6c82ec1fceea8a4e5b39dc51c1373ddfaa13d848df0.
//
// Solidity: event ClaimSuccessful(bytes32 raceId, uint256 claimValue)
func (_Smc *SmcFilterer) FilterClaimSuccessful(opts *bind.FilterOpts) (*SmcClaimSuccessfulIterator, error) {

	logs, sub, err := _Smc.contract.FilterLogs(opts, "ClaimSuccessful")
	if err != nil {
		return nil, err
	}
	return &SmcClaimSuccessfulIterator{contract: _Smc.contract, event: "ClaimSuccessful", logs: logs, sub: sub}, nil
}

// WatchClaimSuccessful is a free log subscription operation binding the contract event 0x78d224c5b360f08891a0d6c82ec1fceea8a4e5b39dc51c1373ddfaa13d848df0.
//
// Solidity: event ClaimSuccessful(bytes32 raceId, uint256 claimValue)
func (_Smc *SmcFilterer) WatchClaimSuccessful(opts *bind.WatchOpts, sink chan<- *SmcClaimSuccessful) (event.Subscription, error) {

	logs, sub, err := _Smc.contract.WatchLogs(opts, "ClaimSuccessful")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SmcClaimSuccessful)
				if err := _Smc.contract.UnpackLog(event, "ClaimSuccessful", log); err != nil {
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

// ParseClaimSuccessful is a log parse operation binding the contract event 0x78d224c5b360f08891a0d6c82ec1fceea8a4e5b39dc51c1373ddfaa13d848df0.
//
// Solidity: event ClaimSuccessful(bytes32 raceId, uint256 claimValue)
func (_Smc *SmcFilterer) ParseClaimSuccessful(log types.Log) (*SmcClaimSuccessful, error) {
	event := new(SmcClaimSuccessful)
	if err := _Smc.contract.UnpackLog(event, "ClaimSuccessful", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SmcFundSuccessfulIterator is returned from FilterFundSuccessful and is used to iterate over the raw logs and unpacked data for FundSuccessful events raised by the Smc contract.
type SmcFundSuccessfulIterator struct {
	Event *SmcFundSuccessful // Event containing the contract specifics and raw log

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
func (it *SmcFundSuccessfulIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SmcFundSuccessful)
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
		it.Event = new(SmcFundSuccessful)
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
func (it *SmcFundSuccessfulIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SmcFundSuccessfulIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SmcFundSuccessful represents a FundSuccessful event raised by the Smc contract.
type SmcFundSuccessful struct {
	RaceId    [32]byte
	FundValue *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterFundSuccessful is a free log retrieval operation binding the contract event 0x18828fa7a168b0db08e3e72ef21348f4c4fa1ee1b46803878b1195794aaad4c3.
//
// Solidity: event FundSuccessful(bytes32 raceId, uint256 fundValue)
func (_Smc *SmcFilterer) FilterFundSuccessful(opts *bind.FilterOpts) (*SmcFundSuccessfulIterator, error) {

	logs, sub, err := _Smc.contract.FilterLogs(opts, "FundSuccessful")
	if err != nil {
		return nil, err
	}
	return &SmcFundSuccessfulIterator{contract: _Smc.contract, event: "FundSuccessful", logs: logs, sub: sub}, nil
}

// WatchFundSuccessful is a free log subscription operation binding the contract event 0x18828fa7a168b0db08e3e72ef21348f4c4fa1ee1b46803878b1195794aaad4c3.
//
// Solidity: event FundSuccessful(bytes32 raceId, uint256 fundValue)
func (_Smc *SmcFilterer) WatchFundSuccessful(opts *bind.WatchOpts, sink chan<- *SmcFundSuccessful) (event.Subscription, error) {

	logs, sub, err := _Smc.contract.WatchLogs(opts, "FundSuccessful")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SmcFundSuccessful)
				if err := _Smc.contract.UnpackLog(event, "FundSuccessful", log); err != nil {
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

// ParseFundSuccessful is a log parse operation binding the contract event 0x18828fa7a168b0db08e3e72ef21348f4c4fa1ee1b46803878b1195794aaad4c3.
//
// Solidity: event FundSuccessful(bytes32 raceId, uint256 fundValue)
func (_Smc *SmcFilterer) ParseFundSuccessful(log types.Log) (*SmcFundSuccessful, error) {
	event := new(SmcFundSuccessful)
	if err := _Smc.contract.UnpackLog(event, "FundSuccessful", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SmcOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Smc contract.
type SmcOwnershipTransferredIterator struct {
	Event *SmcOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *SmcOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SmcOwnershipTransferred)
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
		it.Event = new(SmcOwnershipTransferred)
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
func (it *SmcOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SmcOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SmcOwnershipTransferred represents a OwnershipTransferred event raised by the Smc contract.
type SmcOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Smc *SmcFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*SmcOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Smc.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &SmcOwnershipTransferredIterator{contract: _Smc.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Smc *SmcFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *SmcOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Smc.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SmcOwnershipTransferred)
				if err := _Smc.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Smc *SmcFilterer) ParseOwnershipTransferred(log types.Log) (*SmcOwnershipTransferred, error) {
	event := new(SmcOwnershipTransferred)
	if err := _Smc.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SmcPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the Smc contract.
type SmcPausedIterator struct {
	Event *SmcPaused // Event containing the contract specifics and raw log

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
func (it *SmcPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SmcPaused)
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
		it.Event = new(SmcPaused)
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
func (it *SmcPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SmcPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SmcPaused represents a Paused event raised by the Smc contract.
type SmcPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Smc *SmcFilterer) FilterPaused(opts *bind.FilterOpts) (*SmcPausedIterator, error) {

	logs, sub, err := _Smc.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &SmcPausedIterator{contract: _Smc.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Smc *SmcFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *SmcPaused) (event.Subscription, error) {

	logs, sub, err := _Smc.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SmcPaused)
				if err := _Smc.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_Smc *SmcFilterer) ParsePaused(log types.Log) (*SmcPaused, error) {
	event := new(SmcPaused)
	if err := _Smc.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SmcRaceListUpdatedIterator is returned from FilterRaceListUpdated and is used to iterate over the raw logs and unpacked data for RaceListUpdated events raised by the Smc contract.
type SmcRaceListUpdatedIterator struct {
	Event *SmcRaceListUpdated // Event containing the contract specifics and raw log

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
func (it *SmcRaceListUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SmcRaceListUpdated)
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
		it.Event = new(SmcRaceListUpdated)
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
func (it *SmcRaceListUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SmcRaceListUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SmcRaceListUpdated represents a RaceListUpdated event raised by the Smc contract.
type SmcRaceListUpdated struct {
	Race common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterRaceListUpdated is a free log retrieval operation binding the contract event 0x7a659342c2b03e1c14729e3f1c86f414756d13bd8ff0636713cc22246182b42d.
//
// Solidity: event RaceListUpdated(address race)
func (_Smc *SmcFilterer) FilterRaceListUpdated(opts *bind.FilterOpts) (*SmcRaceListUpdatedIterator, error) {

	logs, sub, err := _Smc.contract.FilterLogs(opts, "RaceListUpdated")
	if err != nil {
		return nil, err
	}
	return &SmcRaceListUpdatedIterator{contract: _Smc.contract, event: "RaceListUpdated", logs: logs, sub: sub}, nil
}

// WatchRaceListUpdated is a free log subscription operation binding the contract event 0x7a659342c2b03e1c14729e3f1c86f414756d13bd8ff0636713cc22246182b42d.
//
// Solidity: event RaceListUpdated(address race)
func (_Smc *SmcFilterer) WatchRaceListUpdated(opts *bind.WatchOpts, sink chan<- *SmcRaceListUpdated) (event.Subscription, error) {

	logs, sub, err := _Smc.contract.WatchLogs(opts, "RaceListUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SmcRaceListUpdated)
				if err := _Smc.contract.UnpackLog(event, "RaceListUpdated", log); err != nil {
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

// ParseRaceListUpdated is a log parse operation binding the contract event 0x7a659342c2b03e1c14729e3f1c86f414756d13bd8ff0636713cc22246182b42d.
//
// Solidity: event RaceListUpdated(address race)
func (_Smc *SmcFilterer) ParseRaceListUpdated(log types.Log) (*SmcRaceListUpdated, error) {
	event := new(SmcRaceListUpdated)
	if err := _Smc.contract.UnpackLog(event, "RaceListUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SmcRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the Smc contract.
type SmcRoleAdminChangedIterator struct {
	Event *SmcRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *SmcRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SmcRoleAdminChanged)
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
		it.Event = new(SmcRoleAdminChanged)
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
func (it *SmcRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SmcRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SmcRoleAdminChanged represents a RoleAdminChanged event raised by the Smc contract.
type SmcRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Smc *SmcFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*SmcRoleAdminChangedIterator, error) {

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

	logs, sub, err := _Smc.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &SmcRoleAdminChangedIterator{contract: _Smc.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Smc *SmcFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *SmcRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _Smc.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SmcRoleAdminChanged)
				if err := _Smc.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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
func (_Smc *SmcFilterer) ParseRoleAdminChanged(log types.Log) (*SmcRoleAdminChanged, error) {
	event := new(SmcRoleAdminChanged)
	if err := _Smc.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SmcRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the Smc contract.
type SmcRoleGrantedIterator struct {
	Event *SmcRoleGranted // Event containing the contract specifics and raw log

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
func (it *SmcRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SmcRoleGranted)
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
		it.Event = new(SmcRoleGranted)
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
func (it *SmcRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SmcRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SmcRoleGranted represents a RoleGranted event raised by the Smc contract.
type SmcRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Smc *SmcFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*SmcRoleGrantedIterator, error) {

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

	logs, sub, err := _Smc.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &SmcRoleGrantedIterator{contract: _Smc.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Smc *SmcFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *SmcRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _Smc.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SmcRoleGranted)
				if err := _Smc.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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
func (_Smc *SmcFilterer) ParseRoleGranted(log types.Log) (*SmcRoleGranted, error) {
	event := new(SmcRoleGranted)
	if err := _Smc.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SmcRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the Smc contract.
type SmcRoleRevokedIterator struct {
	Event *SmcRoleRevoked // Event containing the contract specifics and raw log

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
func (it *SmcRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SmcRoleRevoked)
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
		it.Event = new(SmcRoleRevoked)
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
func (it *SmcRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SmcRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SmcRoleRevoked represents a RoleRevoked event raised by the Smc contract.
type SmcRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Smc *SmcFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*SmcRoleRevokedIterator, error) {

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

	logs, sub, err := _Smc.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &SmcRoleRevokedIterator{contract: _Smc.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Smc *SmcFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *SmcRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _Smc.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SmcRoleRevoked)
				if err := _Smc.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
func (_Smc *SmcFilterer) ParseRoleRevoked(log types.Log) (*SmcRoleRevoked, error) {
	event := new(SmcRoleRevoked)
	if err := _Smc.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SmcUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the Smc contract.
type SmcUnpausedIterator struct {
	Event *SmcUnpaused // Event containing the contract specifics and raw log

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
func (it *SmcUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SmcUnpaused)
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
		it.Event = new(SmcUnpaused)
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
func (it *SmcUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SmcUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SmcUnpaused represents a Unpaused event raised by the Smc contract.
type SmcUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Smc *SmcFilterer) FilterUnpaused(opts *bind.FilterOpts) (*SmcUnpausedIterator, error) {

	logs, sub, err := _Smc.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &SmcUnpausedIterator{contract: _Smc.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Smc *SmcFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *SmcUnpaused) (event.Subscription, error) {

	logs, sub, err := _Smc.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SmcUnpaused)
				if err := _Smc.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
func (_Smc *SmcFilterer) ParseUnpaused(log types.Log) (*SmcUnpaused, error) {
	event := new(SmcUnpaused)
	if err := _Smc.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
