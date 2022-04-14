// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package Bething

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

// BethingMetaData contains all meta data concerning the Bething contract.
var BethingMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"acceptedToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"raceList\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AlreadyReceived\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BetFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FundFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidContract\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidSlot\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RaceNotExisted\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"}],\"name\":\"TooEarly\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"}],\"name\":\"TooLate\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"acceptToken\",\"type\":\"address\"}],\"name\":\"AcceptTokenUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"slotId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"raceId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"betValue\",\"type\":\"uint256\"}],\"name\":\"BetSuccessful\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"raceId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"claimValue\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"ClaimCommission\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"raceId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"claimValue\",\"type\":\"uint256\"}],\"name\":\"ClaimSuccessful\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"raceId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fundValue\",\"type\":\"uint256\"}],\"name\":\"FundSuccessful\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"race\",\"type\":\"address\"}],\"name\":\"RaceListUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PAUSER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_acceptedToken\",\"outputs\":[{\"internalType\":\"contractIBEP20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_raceList\",\"outputs\":[{\"internalType\":\"contractIRaceList\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getRoleMember\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleMemberCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"slotId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"raceId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"betValue\",\"type\":\"uint256\"}],\"name\":\"bet\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"raceId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"fundValue\",\"type\":\"uint256\"}],\"name\":\"fundRace\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"raceId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"slotId\",\"type\":\"uint256\"}],\"name\":\"claim\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"raceId\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"claimCommission\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"raceId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"slotId\",\"type\":\"uint256\"}],\"name\":\"totalSlotBet\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"raceId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"slotId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"userSlotBet\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"raceId\",\"type\":\"bytes32\"}],\"name\":\"totalRaceBet\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"raceId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"slotId\",\"type\":\"uint256\"}],\"name\":\"getSlotPosition\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"raceList\",\"type\":\"address\"}],\"name\":\"updateRaceAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"acceptToken\",\"type\":\"address\"}],\"name\":\"updateAcceptTokenAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"_to\",\"type\":\"address\"}],\"name\":\"destroySmartContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// BethingABI is the input ABI used to generate the binding from.
// Deprecated: Use BethingMetaData.ABI instead.
var BethingABI = BethingMetaData.ABI

// Bething is an auto generated Go binding around an Ethereum contract.
type Bething struct {
	BethingCaller     // Read-only binding to the contract
	BethingTransactor // Write-only binding to the contract
	BethingFilterer   // Log filterer for contract events
}

// BethingCaller is an auto generated read-only Go binding around an Ethereum contract.
type BethingCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BethingTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BethingTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BethingFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BethingFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BethingSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BethingSession struct {
	Contract     *Bething          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BethingCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BethingCallerSession struct {
	Contract *BethingCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// BethingTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BethingTransactorSession struct {
	Contract     *BethingTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// BethingRaw is an auto generated low-level Go binding around an Ethereum contract.
type BethingRaw struct {
	Contract *Bething // Generic contract binding to access the raw methods on
}

// BethingCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BethingCallerRaw struct {
	Contract *BethingCaller // Generic read-only contract binding to access the raw methods on
}

// BethingTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BethingTransactorRaw struct {
	Contract *BethingTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBething creates a new instance of Bething, bound to a specific deployed contract.
func NewBething(address common.Address, backend bind.ContractBackend) (*Bething, error) {
	contract, err := bindBething(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Bething{BethingCaller: BethingCaller{contract: contract}, BethingTransactor: BethingTransactor{contract: contract}, BethingFilterer: BethingFilterer{contract: contract}}, nil
}

// NewBethingCaller creates a new read-only instance of Bething, bound to a specific deployed contract.
func NewBethingCaller(address common.Address, caller bind.ContractCaller) (*BethingCaller, error) {
	contract, err := bindBething(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BethingCaller{contract: contract}, nil
}

// NewBethingTransactor creates a new write-only instance of Bething, bound to a specific deployed contract.
func NewBethingTransactor(address common.Address, transactor bind.ContractTransactor) (*BethingTransactor, error) {
	contract, err := bindBething(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BethingTransactor{contract: contract}, nil
}

// NewBethingFilterer creates a new log filterer instance of Bething, bound to a specific deployed contract.
func NewBethingFilterer(address common.Address, filterer bind.ContractFilterer) (*BethingFilterer, error) {
	contract, err := bindBething(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BethingFilterer{contract: contract}, nil
}

// bindBething binds a generic wrapper to an already deployed contract.
func bindBething(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BethingABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Bething *BethingRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Bething.Contract.BethingCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Bething *BethingRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bething.Contract.BethingTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Bething *BethingRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Bething.Contract.BethingTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Bething *BethingCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Bething.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Bething *BethingTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bething.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Bething *BethingTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Bething.Contract.contract.Transact(opts, method, params...)
}

// ADMINROLE is a free data retrieval call binding the contract method 0x75b238fc.
//
// Solidity: function ADMIN_ROLE() view returns(bytes32)
func (_Bething *BethingCaller) ADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Bething.contract.Call(opts, &out, "ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ADMINROLE is a free data retrieval call binding the contract method 0x75b238fc.
//
// Solidity: function ADMIN_ROLE() view returns(bytes32)
func (_Bething *BethingSession) ADMINROLE() ([32]byte, error) {
	return _Bething.Contract.ADMINROLE(&_Bething.CallOpts)
}

// ADMINROLE is a free data retrieval call binding the contract method 0x75b238fc.
//
// Solidity: function ADMIN_ROLE() view returns(bytes32)
func (_Bething *BethingCallerSession) ADMINROLE() ([32]byte, error) {
	return _Bething.Contract.ADMINROLE(&_Bething.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Bething *BethingCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Bething.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Bething *BethingSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Bething.Contract.DEFAULTADMINROLE(&_Bething.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Bething *BethingCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Bething.Contract.DEFAULTADMINROLE(&_Bething.CallOpts)
}

// PAUSERROLE is a free data retrieval call binding the contract method 0xe63ab1e9.
//
// Solidity: function PAUSER_ROLE() view returns(bytes32)
func (_Bething *BethingCaller) PAUSERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Bething.contract.Call(opts, &out, "PAUSER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// PAUSERROLE is a free data retrieval call binding the contract method 0xe63ab1e9.
//
// Solidity: function PAUSER_ROLE() view returns(bytes32)
func (_Bething *BethingSession) PAUSERROLE() ([32]byte, error) {
	return _Bething.Contract.PAUSERROLE(&_Bething.CallOpts)
}

// PAUSERROLE is a free data retrieval call binding the contract method 0xe63ab1e9.
//
// Solidity: function PAUSER_ROLE() view returns(bytes32)
func (_Bething *BethingCallerSession) PAUSERROLE() ([32]byte, error) {
	return _Bething.Contract.PAUSERROLE(&_Bething.CallOpts)
}

// AcceptedToken is a free data retrieval call binding the contract method 0x4e5b3c7c.
//
// Solidity: function _acceptedToken() view returns(address)
func (_Bething *BethingCaller) AcceptedToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Bething.contract.Call(opts, &out, "_acceptedToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AcceptedToken is a free data retrieval call binding the contract method 0x4e5b3c7c.
//
// Solidity: function _acceptedToken() view returns(address)
func (_Bething *BethingSession) AcceptedToken() (common.Address, error) {
	return _Bething.Contract.AcceptedToken(&_Bething.CallOpts)
}

// AcceptedToken is a free data retrieval call binding the contract method 0x4e5b3c7c.
//
// Solidity: function _acceptedToken() view returns(address)
func (_Bething *BethingCallerSession) AcceptedToken() (common.Address, error) {
	return _Bething.Contract.AcceptedToken(&_Bething.CallOpts)
}

// RaceList is a free data retrieval call binding the contract method 0x6c76f620.
//
// Solidity: function _raceList() view returns(address)
func (_Bething *BethingCaller) RaceList(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Bething.contract.Call(opts, &out, "_raceList")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RaceList is a free data retrieval call binding the contract method 0x6c76f620.
//
// Solidity: function _raceList() view returns(address)
func (_Bething *BethingSession) RaceList() (common.Address, error) {
	return _Bething.Contract.RaceList(&_Bething.CallOpts)
}

// RaceList is a free data retrieval call binding the contract method 0x6c76f620.
//
// Solidity: function _raceList() view returns(address)
func (_Bething *BethingCallerSession) RaceList() (common.Address, error) {
	return _Bething.Contract.RaceList(&_Bething.CallOpts)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Bething *BethingCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _Bething.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Bething *BethingSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Bething.Contract.GetRoleAdmin(&_Bething.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Bething *BethingCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Bething.Contract.GetRoleAdmin(&_Bething.CallOpts, role)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_Bething *BethingCaller) GetRoleMember(opts *bind.CallOpts, role [32]byte, index *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Bething.contract.Call(opts, &out, "getRoleMember", role, index)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_Bething *BethingSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _Bething.Contract.GetRoleMember(&_Bething.CallOpts, role, index)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_Bething *BethingCallerSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _Bething.Contract.GetRoleMember(&_Bething.CallOpts, role, index)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_Bething *BethingCaller) GetRoleMemberCount(opts *bind.CallOpts, role [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _Bething.contract.Call(opts, &out, "getRoleMemberCount", role)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_Bething *BethingSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _Bething.Contract.GetRoleMemberCount(&_Bething.CallOpts, role)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_Bething *BethingCallerSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _Bething.Contract.GetRoleMemberCount(&_Bething.CallOpts, role)
}

// GetSlotPosition is a free data retrieval call binding the contract method 0xd0b57c69.
//
// Solidity: function getSlotPosition(bytes32 raceId, uint256 slotId) view returns(uint8)
func (_Bething *BethingCaller) GetSlotPosition(opts *bind.CallOpts, raceId [32]byte, slotId *big.Int) (uint8, error) {
	var out []interface{}
	err := _Bething.contract.Call(opts, &out, "getSlotPosition", raceId, slotId)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// GetSlotPosition is a free data retrieval call binding the contract method 0xd0b57c69.
//
// Solidity: function getSlotPosition(bytes32 raceId, uint256 slotId) view returns(uint8)
func (_Bething *BethingSession) GetSlotPosition(raceId [32]byte, slotId *big.Int) (uint8, error) {
	return _Bething.Contract.GetSlotPosition(&_Bething.CallOpts, raceId, slotId)
}

// GetSlotPosition is a free data retrieval call binding the contract method 0xd0b57c69.
//
// Solidity: function getSlotPosition(bytes32 raceId, uint256 slotId) view returns(uint8)
func (_Bething *BethingCallerSession) GetSlotPosition(raceId [32]byte, slotId *big.Int) (uint8, error) {
	return _Bething.Contract.GetSlotPosition(&_Bething.CallOpts, raceId, slotId)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Bething *BethingCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _Bething.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Bething *BethingSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Bething.Contract.HasRole(&_Bething.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Bething *BethingCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Bething.Contract.HasRole(&_Bething.CallOpts, role, account)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Bething *BethingCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Bething.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Bething *BethingSession) Owner() (common.Address, error) {
	return _Bething.Contract.Owner(&_Bething.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Bething *BethingCallerSession) Owner() (common.Address, error) {
	return _Bething.Contract.Owner(&_Bething.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Bething *BethingCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Bething.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Bething *BethingSession) Paused() (bool, error) {
	return _Bething.Contract.Paused(&_Bething.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Bething *BethingCallerSession) Paused() (bool, error) {
	return _Bething.Contract.Paused(&_Bething.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Bething *BethingCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Bething.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Bething *BethingSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Bething.Contract.SupportsInterface(&_Bething.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Bething *BethingCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Bething.Contract.SupportsInterface(&_Bething.CallOpts, interfaceId)
}

// TotalRaceBet is a free data retrieval call binding the contract method 0x2123c15e.
//
// Solidity: function totalRaceBet(bytes32 raceId) view returns(uint256)
func (_Bething *BethingCaller) TotalRaceBet(opts *bind.CallOpts, raceId [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _Bething.contract.Call(opts, &out, "totalRaceBet", raceId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalRaceBet is a free data retrieval call binding the contract method 0x2123c15e.
//
// Solidity: function totalRaceBet(bytes32 raceId) view returns(uint256)
func (_Bething *BethingSession) TotalRaceBet(raceId [32]byte) (*big.Int, error) {
	return _Bething.Contract.TotalRaceBet(&_Bething.CallOpts, raceId)
}

// TotalRaceBet is a free data retrieval call binding the contract method 0x2123c15e.
//
// Solidity: function totalRaceBet(bytes32 raceId) view returns(uint256)
func (_Bething *BethingCallerSession) TotalRaceBet(raceId [32]byte) (*big.Int, error) {
	return _Bething.Contract.TotalRaceBet(&_Bething.CallOpts, raceId)
}

// TotalSlotBet is a free data retrieval call binding the contract method 0xe740582a.
//
// Solidity: function totalSlotBet(bytes32 raceId, uint256 slotId) view returns(uint256)
func (_Bething *BethingCaller) TotalSlotBet(opts *bind.CallOpts, raceId [32]byte, slotId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Bething.contract.Call(opts, &out, "totalSlotBet", raceId, slotId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSlotBet is a free data retrieval call binding the contract method 0xe740582a.
//
// Solidity: function totalSlotBet(bytes32 raceId, uint256 slotId) view returns(uint256)
func (_Bething *BethingSession) TotalSlotBet(raceId [32]byte, slotId *big.Int) (*big.Int, error) {
	return _Bething.Contract.TotalSlotBet(&_Bething.CallOpts, raceId, slotId)
}

// TotalSlotBet is a free data retrieval call binding the contract method 0xe740582a.
//
// Solidity: function totalSlotBet(bytes32 raceId, uint256 slotId) view returns(uint256)
func (_Bething *BethingCallerSession) TotalSlotBet(raceId [32]byte, slotId *big.Int) (*big.Int, error) {
	return _Bething.Contract.TotalSlotBet(&_Bething.CallOpts, raceId, slotId)
}

// UserSlotBet is a free data retrieval call binding the contract method 0x2e279801.
//
// Solidity: function userSlotBet(bytes32 raceId, uint256 slotId, address user) view returns(uint256)
func (_Bething *BethingCaller) UserSlotBet(opts *bind.CallOpts, raceId [32]byte, slotId *big.Int, user common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Bething.contract.Call(opts, &out, "userSlotBet", raceId, slotId, user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UserSlotBet is a free data retrieval call binding the contract method 0x2e279801.
//
// Solidity: function userSlotBet(bytes32 raceId, uint256 slotId, address user) view returns(uint256)
func (_Bething *BethingSession) UserSlotBet(raceId [32]byte, slotId *big.Int, user common.Address) (*big.Int, error) {
	return _Bething.Contract.UserSlotBet(&_Bething.CallOpts, raceId, slotId, user)
}

// UserSlotBet is a free data retrieval call binding the contract method 0x2e279801.
//
// Solidity: function userSlotBet(bytes32 raceId, uint256 slotId, address user) view returns(uint256)
func (_Bething *BethingCallerSession) UserSlotBet(raceId [32]byte, slotId *big.Int, user common.Address) (*big.Int, error) {
	return _Bething.Contract.UserSlotBet(&_Bething.CallOpts, raceId, slotId, user)
}

// Bet is a paid mutator transaction binding the contract method 0x94fc859b.
//
// Solidity: function bet(uint256 slotId, bytes32 raceId, uint256 betValue) returns()
func (_Bething *BethingTransactor) Bet(opts *bind.TransactOpts, slotId *big.Int, raceId [32]byte, betValue *big.Int) (*types.Transaction, error) {
	return _Bething.contract.Transact(opts, "bet", slotId, raceId, betValue)
}

// Bet is a paid mutator transaction binding the contract method 0x94fc859b.
//
// Solidity: function bet(uint256 slotId, bytes32 raceId, uint256 betValue) returns()
func (_Bething *BethingSession) Bet(slotId *big.Int, raceId [32]byte, betValue *big.Int) (*types.Transaction, error) {
	return _Bething.Contract.Bet(&_Bething.TransactOpts, slotId, raceId, betValue)
}

// Bet is a paid mutator transaction binding the contract method 0x94fc859b.
//
// Solidity: function bet(uint256 slotId, bytes32 raceId, uint256 betValue) returns()
func (_Bething *BethingTransactorSession) Bet(slotId *big.Int, raceId [32]byte, betValue *big.Int) (*types.Transaction, error) {
	return _Bething.Contract.Bet(&_Bething.TransactOpts, slotId, raceId, betValue)
}

// Claim is a paid mutator transaction binding the contract method 0x63f44968.
//
// Solidity: function claim(bytes32 raceId, uint256 slotId) returns()
func (_Bething *BethingTransactor) Claim(opts *bind.TransactOpts, raceId [32]byte, slotId *big.Int) (*types.Transaction, error) {
	return _Bething.contract.Transact(opts, "claim", raceId, slotId)
}

// Claim is a paid mutator transaction binding the contract method 0x63f44968.
//
// Solidity: function claim(bytes32 raceId, uint256 slotId) returns()
func (_Bething *BethingSession) Claim(raceId [32]byte, slotId *big.Int) (*types.Transaction, error) {
	return _Bething.Contract.Claim(&_Bething.TransactOpts, raceId, slotId)
}

// Claim is a paid mutator transaction binding the contract method 0x63f44968.
//
// Solidity: function claim(bytes32 raceId, uint256 slotId) returns()
func (_Bething *BethingTransactorSession) Claim(raceId [32]byte, slotId *big.Int) (*types.Transaction, error) {
	return _Bething.Contract.Claim(&_Bething.TransactOpts, raceId, slotId)
}

// ClaimCommission is a paid mutator transaction binding the contract method 0x5cce2b42.
//
// Solidity: function claimCommission(bytes32 raceId, address receiver) returns()
func (_Bething *BethingTransactor) ClaimCommission(opts *bind.TransactOpts, raceId [32]byte, receiver common.Address) (*types.Transaction, error) {
	return _Bething.contract.Transact(opts, "claimCommission", raceId, receiver)
}

// ClaimCommission is a paid mutator transaction binding the contract method 0x5cce2b42.
//
// Solidity: function claimCommission(bytes32 raceId, address receiver) returns()
func (_Bething *BethingSession) ClaimCommission(raceId [32]byte, receiver common.Address) (*types.Transaction, error) {
	return _Bething.Contract.ClaimCommission(&_Bething.TransactOpts, raceId, receiver)
}

// ClaimCommission is a paid mutator transaction binding the contract method 0x5cce2b42.
//
// Solidity: function claimCommission(bytes32 raceId, address receiver) returns()
func (_Bething *BethingTransactorSession) ClaimCommission(raceId [32]byte, receiver common.Address) (*types.Transaction, error) {
	return _Bething.Contract.ClaimCommission(&_Bething.TransactOpts, raceId, receiver)
}

// DestroySmartContract is a paid mutator transaction binding the contract method 0x39df43ff.
//
// Solidity: function destroySmartContract(address _to) returns()
func (_Bething *BethingTransactor) DestroySmartContract(opts *bind.TransactOpts, _to common.Address) (*types.Transaction, error) {
	return _Bething.contract.Transact(opts, "destroySmartContract", _to)
}

// DestroySmartContract is a paid mutator transaction binding the contract method 0x39df43ff.
//
// Solidity: function destroySmartContract(address _to) returns()
func (_Bething *BethingSession) DestroySmartContract(_to common.Address) (*types.Transaction, error) {
	return _Bething.Contract.DestroySmartContract(&_Bething.TransactOpts, _to)
}

// DestroySmartContract is a paid mutator transaction binding the contract method 0x39df43ff.
//
// Solidity: function destroySmartContract(address _to) returns()
func (_Bething *BethingTransactorSession) DestroySmartContract(_to common.Address) (*types.Transaction, error) {
	return _Bething.Contract.DestroySmartContract(&_Bething.TransactOpts, _to)
}

// FundRace is a paid mutator transaction binding the contract method 0x2ebe4493.
//
// Solidity: function fundRace(bytes32 raceId, uint256 fundValue) returns()
func (_Bething *BethingTransactor) FundRace(opts *bind.TransactOpts, raceId [32]byte, fundValue *big.Int) (*types.Transaction, error) {
	return _Bething.contract.Transact(opts, "fundRace", raceId, fundValue)
}

// FundRace is a paid mutator transaction binding the contract method 0x2ebe4493.
//
// Solidity: function fundRace(bytes32 raceId, uint256 fundValue) returns()
func (_Bething *BethingSession) FundRace(raceId [32]byte, fundValue *big.Int) (*types.Transaction, error) {
	return _Bething.Contract.FundRace(&_Bething.TransactOpts, raceId, fundValue)
}

// FundRace is a paid mutator transaction binding the contract method 0x2ebe4493.
//
// Solidity: function fundRace(bytes32 raceId, uint256 fundValue) returns()
func (_Bething *BethingTransactorSession) FundRace(raceId [32]byte, fundValue *big.Int) (*types.Transaction, error) {
	return _Bething.Contract.FundRace(&_Bething.TransactOpts, raceId, fundValue)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Bething *BethingTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Bething.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Bething *BethingSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Bething.Contract.GrantRole(&_Bething.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Bething *BethingTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Bething.Contract.GrantRole(&_Bething.TransactOpts, role, account)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Bething *BethingTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bething.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Bething *BethingSession) Pause() (*types.Transaction, error) {
	return _Bething.Contract.Pause(&_Bething.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Bething *BethingTransactorSession) Pause() (*types.Transaction, error) {
	return _Bething.Contract.Pause(&_Bething.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Bething *BethingTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bething.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Bething *BethingSession) RenounceOwnership() (*types.Transaction, error) {
	return _Bething.Contract.RenounceOwnership(&_Bething.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Bething *BethingTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Bething.Contract.RenounceOwnership(&_Bething.TransactOpts)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Bething *BethingTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Bething.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Bething *BethingSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Bething.Contract.RenounceRole(&_Bething.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Bething *BethingTransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Bething.Contract.RenounceRole(&_Bething.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Bething *BethingTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Bething.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Bething *BethingSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Bething.Contract.RevokeRole(&_Bething.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Bething *BethingTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Bething.Contract.RevokeRole(&_Bething.TransactOpts, role, account)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Bething *BethingTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Bething.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Bething *BethingSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Bething.Contract.TransferOwnership(&_Bething.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Bething *BethingTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Bething.Contract.TransferOwnership(&_Bething.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Bething *BethingTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bething.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Bething *BethingSession) Unpause() (*types.Transaction, error) {
	return _Bething.Contract.Unpause(&_Bething.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Bething *BethingTransactorSession) Unpause() (*types.Transaction, error) {
	return _Bething.Contract.Unpause(&_Bething.TransactOpts)
}

// UpdateAcceptTokenAddress is a paid mutator transaction binding the contract method 0xb4a2fad0.
//
// Solidity: function updateAcceptTokenAddress(address acceptToken) returns()
func (_Bething *BethingTransactor) UpdateAcceptTokenAddress(opts *bind.TransactOpts, acceptToken common.Address) (*types.Transaction, error) {
	return _Bething.contract.Transact(opts, "updateAcceptTokenAddress", acceptToken)
}

// UpdateAcceptTokenAddress is a paid mutator transaction binding the contract method 0xb4a2fad0.
//
// Solidity: function updateAcceptTokenAddress(address acceptToken) returns()
func (_Bething *BethingSession) UpdateAcceptTokenAddress(acceptToken common.Address) (*types.Transaction, error) {
	return _Bething.Contract.UpdateAcceptTokenAddress(&_Bething.TransactOpts, acceptToken)
}

// UpdateAcceptTokenAddress is a paid mutator transaction binding the contract method 0xb4a2fad0.
//
// Solidity: function updateAcceptTokenAddress(address acceptToken) returns()
func (_Bething *BethingTransactorSession) UpdateAcceptTokenAddress(acceptToken common.Address) (*types.Transaction, error) {
	return _Bething.Contract.UpdateAcceptTokenAddress(&_Bething.TransactOpts, acceptToken)
}

// UpdateRaceAddress is a paid mutator transaction binding the contract method 0xd7c49938.
//
// Solidity: function updateRaceAddress(address raceList) returns()
func (_Bething *BethingTransactor) UpdateRaceAddress(opts *bind.TransactOpts, raceList common.Address) (*types.Transaction, error) {
	return _Bething.contract.Transact(opts, "updateRaceAddress", raceList)
}

// UpdateRaceAddress is a paid mutator transaction binding the contract method 0xd7c49938.
//
// Solidity: function updateRaceAddress(address raceList) returns()
func (_Bething *BethingSession) UpdateRaceAddress(raceList common.Address) (*types.Transaction, error) {
	return _Bething.Contract.UpdateRaceAddress(&_Bething.TransactOpts, raceList)
}

// UpdateRaceAddress is a paid mutator transaction binding the contract method 0xd7c49938.
//
// Solidity: function updateRaceAddress(address raceList) returns()
func (_Bething *BethingTransactorSession) UpdateRaceAddress(raceList common.Address) (*types.Transaction, error) {
	return _Bething.Contract.UpdateRaceAddress(&_Bething.TransactOpts, raceList)
}

// BethingAcceptTokenUpdatedIterator is returned from FilterAcceptTokenUpdated and is used to iterate over the raw logs and unpacked data for AcceptTokenUpdated events raised by the Bething contract.
type BethingAcceptTokenUpdatedIterator struct {
	Event *BethingAcceptTokenUpdated // Event containing the contract specifics and raw log

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
func (it *BethingAcceptTokenUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BethingAcceptTokenUpdated)
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
		it.Event = new(BethingAcceptTokenUpdated)
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
func (it *BethingAcceptTokenUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BethingAcceptTokenUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BethingAcceptTokenUpdated represents a AcceptTokenUpdated event raised by the Bething contract.
type BethingAcceptTokenUpdated struct {
	AcceptToken common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterAcceptTokenUpdated is a free log retrieval operation binding the contract event 0x44c3b193081dd50818bc919dc1963ae03762174c9a89af5fbbce54ee71a13f97.
//
// Solidity: event AcceptTokenUpdated(address acceptToken)
func (_Bething *BethingFilterer) FilterAcceptTokenUpdated(opts *bind.FilterOpts) (*BethingAcceptTokenUpdatedIterator, error) {

	logs, sub, err := _Bething.contract.FilterLogs(opts, "AcceptTokenUpdated")
	if err != nil {
		return nil, err
	}
	return &BethingAcceptTokenUpdatedIterator{contract: _Bething.contract, event: "AcceptTokenUpdated", logs: logs, sub: sub}, nil
}

// WatchAcceptTokenUpdated is a free log subscription operation binding the contract event 0x44c3b193081dd50818bc919dc1963ae03762174c9a89af5fbbce54ee71a13f97.
//
// Solidity: event AcceptTokenUpdated(address acceptToken)
func (_Bething *BethingFilterer) WatchAcceptTokenUpdated(opts *bind.WatchOpts, sink chan<- *BethingAcceptTokenUpdated) (event.Subscription, error) {

	logs, sub, err := _Bething.contract.WatchLogs(opts, "AcceptTokenUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BethingAcceptTokenUpdated)
				if err := _Bething.contract.UnpackLog(event, "AcceptTokenUpdated", log); err != nil {
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
func (_Bething *BethingFilterer) ParseAcceptTokenUpdated(log types.Log) (*BethingAcceptTokenUpdated, error) {
	event := new(BethingAcceptTokenUpdated)
	if err := _Bething.contract.UnpackLog(event, "AcceptTokenUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BethingBetSuccessfulIterator is returned from FilterBetSuccessful and is used to iterate over the raw logs and unpacked data for BetSuccessful events raised by the Bething contract.
type BethingBetSuccessfulIterator struct {
	Event *BethingBetSuccessful // Event containing the contract specifics and raw log

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
func (it *BethingBetSuccessfulIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BethingBetSuccessful)
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
		it.Event = new(BethingBetSuccessful)
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
func (it *BethingBetSuccessfulIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BethingBetSuccessfulIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BethingBetSuccessful represents a BetSuccessful event raised by the Bething contract.
type BethingBetSuccessful struct {
	SlotId   *big.Int
	RaceId   [32]byte
	BetValue *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterBetSuccessful is a free log retrieval operation binding the contract event 0x90340911801212a3d95efe0dca9cdbd74549020af8cf33fc0861b74d82fe6dbd.
//
// Solidity: event BetSuccessful(uint256 slotId, bytes32 raceId, uint256 betValue)
func (_Bething *BethingFilterer) FilterBetSuccessful(opts *bind.FilterOpts) (*BethingBetSuccessfulIterator, error) {

	logs, sub, err := _Bething.contract.FilterLogs(opts, "BetSuccessful")
	if err != nil {
		return nil, err
	}
	return &BethingBetSuccessfulIterator{contract: _Bething.contract, event: "BetSuccessful", logs: logs, sub: sub}, nil
}

// WatchBetSuccessful is a free log subscription operation binding the contract event 0x90340911801212a3d95efe0dca9cdbd74549020af8cf33fc0861b74d82fe6dbd.
//
// Solidity: event BetSuccessful(uint256 slotId, bytes32 raceId, uint256 betValue)
func (_Bething *BethingFilterer) WatchBetSuccessful(opts *bind.WatchOpts, sink chan<- *BethingBetSuccessful) (event.Subscription, error) {

	logs, sub, err := _Bething.contract.WatchLogs(opts, "BetSuccessful")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BethingBetSuccessful)
				if err := _Bething.contract.UnpackLog(event, "BetSuccessful", log); err != nil {
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
func (_Bething *BethingFilterer) ParseBetSuccessful(log types.Log) (*BethingBetSuccessful, error) {
	event := new(BethingBetSuccessful)
	if err := _Bething.contract.UnpackLog(event, "BetSuccessful", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BethingClaimCommissionIterator is returned from FilterClaimCommission and is used to iterate over the raw logs and unpacked data for ClaimCommission events raised by the Bething contract.
type BethingClaimCommissionIterator struct {
	Event *BethingClaimCommission // Event containing the contract specifics and raw log

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
func (it *BethingClaimCommissionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BethingClaimCommission)
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
		it.Event = new(BethingClaimCommission)
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
func (it *BethingClaimCommissionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BethingClaimCommissionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BethingClaimCommission represents a ClaimCommission event raised by the Bething contract.
type BethingClaimCommission struct {
	RaceId     [32]byte
	ClaimValue *big.Int
	Receiver   common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterClaimCommission is a free log retrieval operation binding the contract event 0x94fe36569bb216f6cb564fee8983d035c7cacba5679023a4ccabe33fdf6d4d88.
//
// Solidity: event ClaimCommission(bytes32 raceId, uint256 claimValue, address receiver)
func (_Bething *BethingFilterer) FilterClaimCommission(opts *bind.FilterOpts) (*BethingClaimCommissionIterator, error) {

	logs, sub, err := _Bething.contract.FilterLogs(opts, "ClaimCommission")
	if err != nil {
		return nil, err
	}
	return &BethingClaimCommissionIterator{contract: _Bething.contract, event: "ClaimCommission", logs: logs, sub: sub}, nil
}

// WatchClaimCommission is a free log subscription operation binding the contract event 0x94fe36569bb216f6cb564fee8983d035c7cacba5679023a4ccabe33fdf6d4d88.
//
// Solidity: event ClaimCommission(bytes32 raceId, uint256 claimValue, address receiver)
func (_Bething *BethingFilterer) WatchClaimCommission(opts *bind.WatchOpts, sink chan<- *BethingClaimCommission) (event.Subscription, error) {

	logs, sub, err := _Bething.contract.WatchLogs(opts, "ClaimCommission")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BethingClaimCommission)
				if err := _Bething.contract.UnpackLog(event, "ClaimCommission", log); err != nil {
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
func (_Bething *BethingFilterer) ParseClaimCommission(log types.Log) (*BethingClaimCommission, error) {
	event := new(BethingClaimCommission)
	if err := _Bething.contract.UnpackLog(event, "ClaimCommission", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BethingClaimSuccessfulIterator is returned from FilterClaimSuccessful and is used to iterate over the raw logs and unpacked data for ClaimSuccessful events raised by the Bething contract.
type BethingClaimSuccessfulIterator struct {
	Event *BethingClaimSuccessful // Event containing the contract specifics and raw log

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
func (it *BethingClaimSuccessfulIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BethingClaimSuccessful)
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
		it.Event = new(BethingClaimSuccessful)
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
func (it *BethingClaimSuccessfulIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BethingClaimSuccessfulIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BethingClaimSuccessful represents a ClaimSuccessful event raised by the Bething contract.
type BethingClaimSuccessful struct {
	RaceId     [32]byte
	ClaimValue *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterClaimSuccessful is a free log retrieval operation binding the contract event 0x78d224c5b360f08891a0d6c82ec1fceea8a4e5b39dc51c1373ddfaa13d848df0.
//
// Solidity: event ClaimSuccessful(bytes32 raceId, uint256 claimValue)
func (_Bething *BethingFilterer) FilterClaimSuccessful(opts *bind.FilterOpts) (*BethingClaimSuccessfulIterator, error) {

	logs, sub, err := _Bething.contract.FilterLogs(opts, "ClaimSuccessful")
	if err != nil {
		return nil, err
	}
	return &BethingClaimSuccessfulIterator{contract: _Bething.contract, event: "ClaimSuccessful", logs: logs, sub: sub}, nil
}

// WatchClaimSuccessful is a free log subscription operation binding the contract event 0x78d224c5b360f08891a0d6c82ec1fceea8a4e5b39dc51c1373ddfaa13d848df0.
//
// Solidity: event ClaimSuccessful(bytes32 raceId, uint256 claimValue)
func (_Bething *BethingFilterer) WatchClaimSuccessful(opts *bind.WatchOpts, sink chan<- *BethingClaimSuccessful) (event.Subscription, error) {

	logs, sub, err := _Bething.contract.WatchLogs(opts, "ClaimSuccessful")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BethingClaimSuccessful)
				if err := _Bething.contract.UnpackLog(event, "ClaimSuccessful", log); err != nil {
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
func (_Bething *BethingFilterer) ParseClaimSuccessful(log types.Log) (*BethingClaimSuccessful, error) {
	event := new(BethingClaimSuccessful)
	if err := _Bething.contract.UnpackLog(event, "ClaimSuccessful", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BethingFundSuccessfulIterator is returned from FilterFundSuccessful and is used to iterate over the raw logs and unpacked data for FundSuccessful events raised by the Bething contract.
type BethingFundSuccessfulIterator struct {
	Event *BethingFundSuccessful // Event containing the contract specifics and raw log

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
func (it *BethingFundSuccessfulIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BethingFundSuccessful)
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
		it.Event = new(BethingFundSuccessful)
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
func (it *BethingFundSuccessfulIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BethingFundSuccessfulIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BethingFundSuccessful represents a FundSuccessful event raised by the Bething contract.
type BethingFundSuccessful struct {
	RaceId    [32]byte
	FundValue *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterFundSuccessful is a free log retrieval operation binding the contract event 0x18828fa7a168b0db08e3e72ef21348f4c4fa1ee1b46803878b1195794aaad4c3.
//
// Solidity: event FundSuccessful(bytes32 raceId, uint256 fundValue)
func (_Bething *BethingFilterer) FilterFundSuccessful(opts *bind.FilterOpts) (*BethingFundSuccessfulIterator, error) {

	logs, sub, err := _Bething.contract.FilterLogs(opts, "FundSuccessful")
	if err != nil {
		return nil, err
	}
	return &BethingFundSuccessfulIterator{contract: _Bething.contract, event: "FundSuccessful", logs: logs, sub: sub}, nil
}

// WatchFundSuccessful is a free log subscription operation binding the contract event 0x18828fa7a168b0db08e3e72ef21348f4c4fa1ee1b46803878b1195794aaad4c3.
//
// Solidity: event FundSuccessful(bytes32 raceId, uint256 fundValue)
func (_Bething *BethingFilterer) WatchFundSuccessful(opts *bind.WatchOpts, sink chan<- *BethingFundSuccessful) (event.Subscription, error) {

	logs, sub, err := _Bething.contract.WatchLogs(opts, "FundSuccessful")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BethingFundSuccessful)
				if err := _Bething.contract.UnpackLog(event, "FundSuccessful", log); err != nil {
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
func (_Bething *BethingFilterer) ParseFundSuccessful(log types.Log) (*BethingFundSuccessful, error) {
	event := new(BethingFundSuccessful)
	if err := _Bething.contract.UnpackLog(event, "FundSuccessful", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BethingOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Bething contract.
type BethingOwnershipTransferredIterator struct {
	Event *BethingOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *BethingOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BethingOwnershipTransferred)
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
		it.Event = new(BethingOwnershipTransferred)
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
func (it *BethingOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BethingOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BethingOwnershipTransferred represents a OwnershipTransferred event raised by the Bething contract.
type BethingOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Bething *BethingFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*BethingOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Bething.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &BethingOwnershipTransferredIterator{contract: _Bething.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Bething *BethingFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *BethingOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Bething.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BethingOwnershipTransferred)
				if err := _Bething.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Bething *BethingFilterer) ParseOwnershipTransferred(log types.Log) (*BethingOwnershipTransferred, error) {
	event := new(BethingOwnershipTransferred)
	if err := _Bething.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BethingPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the Bething contract.
type BethingPausedIterator struct {
	Event *BethingPaused // Event containing the contract specifics and raw log

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
func (it *BethingPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BethingPaused)
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
		it.Event = new(BethingPaused)
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
func (it *BethingPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BethingPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BethingPaused represents a Paused event raised by the Bething contract.
type BethingPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Bething *BethingFilterer) FilterPaused(opts *bind.FilterOpts) (*BethingPausedIterator, error) {

	logs, sub, err := _Bething.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &BethingPausedIterator{contract: _Bething.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Bething *BethingFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *BethingPaused) (event.Subscription, error) {

	logs, sub, err := _Bething.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BethingPaused)
				if err := _Bething.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_Bething *BethingFilterer) ParsePaused(log types.Log) (*BethingPaused, error) {
	event := new(BethingPaused)
	if err := _Bething.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BethingRaceListUpdatedIterator is returned from FilterRaceListUpdated and is used to iterate over the raw logs and unpacked data for RaceListUpdated events raised by the Bething contract.
type BethingRaceListUpdatedIterator struct {
	Event *BethingRaceListUpdated // Event containing the contract specifics and raw log

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
func (it *BethingRaceListUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BethingRaceListUpdated)
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
		it.Event = new(BethingRaceListUpdated)
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
func (it *BethingRaceListUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BethingRaceListUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BethingRaceListUpdated represents a RaceListUpdated event raised by the Bething contract.
type BethingRaceListUpdated struct {
	Race common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterRaceListUpdated is a free log retrieval operation binding the contract event 0x7a659342c2b03e1c14729e3f1c86f414756d13bd8ff0636713cc22246182b42d.
//
// Solidity: event RaceListUpdated(address race)
func (_Bething *BethingFilterer) FilterRaceListUpdated(opts *bind.FilterOpts) (*BethingRaceListUpdatedIterator, error) {

	logs, sub, err := _Bething.contract.FilterLogs(opts, "RaceListUpdated")
	if err != nil {
		return nil, err
	}
	return &BethingRaceListUpdatedIterator{contract: _Bething.contract, event: "RaceListUpdated", logs: logs, sub: sub}, nil
}

// WatchRaceListUpdated is a free log subscription operation binding the contract event 0x7a659342c2b03e1c14729e3f1c86f414756d13bd8ff0636713cc22246182b42d.
//
// Solidity: event RaceListUpdated(address race)
func (_Bething *BethingFilterer) WatchRaceListUpdated(opts *bind.WatchOpts, sink chan<- *BethingRaceListUpdated) (event.Subscription, error) {

	logs, sub, err := _Bething.contract.WatchLogs(opts, "RaceListUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BethingRaceListUpdated)
				if err := _Bething.contract.UnpackLog(event, "RaceListUpdated", log); err != nil {
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
func (_Bething *BethingFilterer) ParseRaceListUpdated(log types.Log) (*BethingRaceListUpdated, error) {
	event := new(BethingRaceListUpdated)
	if err := _Bething.contract.UnpackLog(event, "RaceListUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BethingRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the Bething contract.
type BethingRoleAdminChangedIterator struct {
	Event *BethingRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *BethingRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BethingRoleAdminChanged)
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
		it.Event = new(BethingRoleAdminChanged)
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
func (it *BethingRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BethingRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BethingRoleAdminChanged represents a RoleAdminChanged event raised by the Bething contract.
type BethingRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Bething *BethingFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*BethingRoleAdminChangedIterator, error) {

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

	logs, sub, err := _Bething.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &BethingRoleAdminChangedIterator{contract: _Bething.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Bething *BethingFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *BethingRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _Bething.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BethingRoleAdminChanged)
				if err := _Bething.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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
func (_Bething *BethingFilterer) ParseRoleAdminChanged(log types.Log) (*BethingRoleAdminChanged, error) {
	event := new(BethingRoleAdminChanged)
	if err := _Bething.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BethingRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the Bething contract.
type BethingRoleGrantedIterator struct {
	Event *BethingRoleGranted // Event containing the contract specifics and raw log

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
func (it *BethingRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BethingRoleGranted)
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
		it.Event = new(BethingRoleGranted)
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
func (it *BethingRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BethingRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BethingRoleGranted represents a RoleGranted event raised by the Bething contract.
type BethingRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Bething *BethingFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*BethingRoleGrantedIterator, error) {

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

	logs, sub, err := _Bething.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &BethingRoleGrantedIterator{contract: _Bething.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Bething *BethingFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *BethingRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _Bething.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BethingRoleGranted)
				if err := _Bething.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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
func (_Bething *BethingFilterer) ParseRoleGranted(log types.Log) (*BethingRoleGranted, error) {
	event := new(BethingRoleGranted)
	if err := _Bething.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BethingRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the Bething contract.
type BethingRoleRevokedIterator struct {
	Event *BethingRoleRevoked // Event containing the contract specifics and raw log

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
func (it *BethingRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BethingRoleRevoked)
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
		it.Event = new(BethingRoleRevoked)
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
func (it *BethingRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BethingRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BethingRoleRevoked represents a RoleRevoked event raised by the Bething contract.
type BethingRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Bething *BethingFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*BethingRoleRevokedIterator, error) {

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

	logs, sub, err := _Bething.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &BethingRoleRevokedIterator{contract: _Bething.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Bething *BethingFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *BethingRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _Bething.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BethingRoleRevoked)
				if err := _Bething.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
func (_Bething *BethingFilterer) ParseRoleRevoked(log types.Log) (*BethingRoleRevoked, error) {
	event := new(BethingRoleRevoked)
	if err := _Bething.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BethingUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the Bething contract.
type BethingUnpausedIterator struct {
	Event *BethingUnpaused // Event containing the contract specifics and raw log

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
func (it *BethingUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BethingUnpaused)
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
		it.Event = new(BethingUnpaused)
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
func (it *BethingUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BethingUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BethingUnpaused represents a Unpaused event raised by the Bething contract.
type BethingUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Bething *BethingFilterer) FilterUnpaused(opts *bind.FilterOpts) (*BethingUnpausedIterator, error) {

	logs, sub, err := _Bething.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &BethingUnpausedIterator{contract: _Bething.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Bething *BethingFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *BethingUnpaused) (event.Subscription, error) {

	logs, sub, err := _Bething.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BethingUnpaused)
				if err := _Bething.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
func (_Bething *BethingFilterer) ParseUnpaused(log types.Log) (*BethingUnpaused, error) {
	event := new(BethingUnpaused)
	if err := _Bething.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
