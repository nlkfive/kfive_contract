// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package KickOffLeague

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

// KickOffLeagueMetaData contains all meta data concerning the KickOffLeague contract.
var KickOffLeagueMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_leagueName\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"nlggtAddr\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"CannotCancel\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidContract\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidSlot\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RaceExisted\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RaceNotExisted\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"}],\"name\":\"TooEarly\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"}],\"name\":\"TooLate\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"}],\"name\":\"RaceCancelled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"slots\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"startAt\",\"type\":\"uint256\"}],\"name\":\"RaceCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"result\",\"type\":\"bytes32\"}],\"name\":\"RaceResultUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PAUSER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"raceId\",\"type\":\"bytes32\"}],\"name\":\"cancelRace\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"slots\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startAt\",\"type\":\"uint256\"}],\"name\":\"createRace\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getRoleMember\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleMemberCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"leagueName\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"raceId\",\"type\":\"bytes32\"}],\"name\":\"raceResult\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"result\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"raceId\",\"type\":\"bytes32\"}],\"name\":\"raceSlots\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"slots\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"raceId\",\"type\":\"bytes32\"}],\"name\":\"raceStartAt\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"startAt\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"raceId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"result\",\"type\":\"bytes32\"}],\"name\":\"updateRaceResult\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"register\",\"type\":\"address\"}],\"name\":\"registerIsValid\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// KickOffLeagueABI is the input ABI used to generate the binding from.
// Deprecated: Use KickOffLeagueMetaData.ABI instead.
var KickOffLeagueABI = KickOffLeagueMetaData.ABI

// KickOffLeague is an auto generated Go binding around an Ethereum contract.
type KickOffLeague struct {
	KickOffLeagueCaller     // Read-only binding to the contract
	KickOffLeagueTransactor // Write-only binding to the contract
	KickOffLeagueFilterer   // Log filterer for contract events
}

// KickOffLeagueCaller is an auto generated read-only Go binding around an Ethereum contract.
type KickOffLeagueCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KickOffLeagueTransactor is an auto generated write-only Go binding around an Ethereum contract.
type KickOffLeagueTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KickOffLeagueFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type KickOffLeagueFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KickOffLeagueSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type KickOffLeagueSession struct {
	Contract     *KickOffLeague    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// KickOffLeagueCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type KickOffLeagueCallerSession struct {
	Contract *KickOffLeagueCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// KickOffLeagueTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type KickOffLeagueTransactorSession struct {
	Contract     *KickOffLeagueTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// KickOffLeagueRaw is an auto generated low-level Go binding around an Ethereum contract.
type KickOffLeagueRaw struct {
	Contract *KickOffLeague // Generic contract binding to access the raw methods on
}

// KickOffLeagueCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type KickOffLeagueCallerRaw struct {
	Contract *KickOffLeagueCaller // Generic read-only contract binding to access the raw methods on
}

// KickOffLeagueTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type KickOffLeagueTransactorRaw struct {
	Contract *KickOffLeagueTransactor // Generic write-only contract binding to access the raw methods on
}

// NewKickOffLeague creates a new instance of KickOffLeague, bound to a specific deployed contract.
func NewKickOffLeague(address common.Address, backend bind.ContractBackend) (*KickOffLeague, error) {
	contract, err := bindKickOffLeague(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &KickOffLeague{KickOffLeagueCaller: KickOffLeagueCaller{contract: contract}, KickOffLeagueTransactor: KickOffLeagueTransactor{contract: contract}, KickOffLeagueFilterer: KickOffLeagueFilterer{contract: contract}}, nil
}

// NewKickOffLeagueCaller creates a new read-only instance of KickOffLeague, bound to a specific deployed contract.
func NewKickOffLeagueCaller(address common.Address, caller bind.ContractCaller) (*KickOffLeagueCaller, error) {
	contract, err := bindKickOffLeague(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &KickOffLeagueCaller{contract: contract}, nil
}

// NewKickOffLeagueTransactor creates a new write-only instance of KickOffLeague, bound to a specific deployed contract.
func NewKickOffLeagueTransactor(address common.Address, transactor bind.ContractTransactor) (*KickOffLeagueTransactor, error) {
	contract, err := bindKickOffLeague(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &KickOffLeagueTransactor{contract: contract}, nil
}

// NewKickOffLeagueFilterer creates a new log filterer instance of KickOffLeague, bound to a specific deployed contract.
func NewKickOffLeagueFilterer(address common.Address, filterer bind.ContractFilterer) (*KickOffLeagueFilterer, error) {
	contract, err := bindKickOffLeague(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &KickOffLeagueFilterer{contract: contract}, nil
}

// bindKickOffLeague binds a generic wrapper to an already deployed contract.
func bindKickOffLeague(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(KickOffLeagueABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_KickOffLeague *KickOffLeagueRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _KickOffLeague.Contract.KickOffLeagueCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_KickOffLeague *KickOffLeagueRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _KickOffLeague.Contract.KickOffLeagueTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_KickOffLeague *KickOffLeagueRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _KickOffLeague.Contract.KickOffLeagueTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_KickOffLeague *KickOffLeagueCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _KickOffLeague.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_KickOffLeague *KickOffLeagueTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _KickOffLeague.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_KickOffLeague *KickOffLeagueTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _KickOffLeague.Contract.contract.Transact(opts, method, params...)
}

// ADMINROLE is a free data retrieval call binding the contract method 0x75b238fc.
//
// Solidity: function ADMIN_ROLE() view returns(bytes32)
func (_KickOffLeague *KickOffLeagueCaller) ADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _KickOffLeague.contract.Call(opts, &out, "ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ADMINROLE is a free data retrieval call binding the contract method 0x75b238fc.
//
// Solidity: function ADMIN_ROLE() view returns(bytes32)
func (_KickOffLeague *KickOffLeagueSession) ADMINROLE() ([32]byte, error) {
	return _KickOffLeague.Contract.ADMINROLE(&_KickOffLeague.CallOpts)
}

// ADMINROLE is a free data retrieval call binding the contract method 0x75b238fc.
//
// Solidity: function ADMIN_ROLE() view returns(bytes32)
func (_KickOffLeague *KickOffLeagueCallerSession) ADMINROLE() ([32]byte, error) {
	return _KickOffLeague.Contract.ADMINROLE(&_KickOffLeague.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_KickOffLeague *KickOffLeagueCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _KickOffLeague.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_KickOffLeague *KickOffLeagueSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _KickOffLeague.Contract.DEFAULTADMINROLE(&_KickOffLeague.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_KickOffLeague *KickOffLeagueCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _KickOffLeague.Contract.DEFAULTADMINROLE(&_KickOffLeague.CallOpts)
}

// PAUSERROLE is a free data retrieval call binding the contract method 0xe63ab1e9.
//
// Solidity: function PAUSER_ROLE() view returns(bytes32)
func (_KickOffLeague *KickOffLeagueCaller) PAUSERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _KickOffLeague.contract.Call(opts, &out, "PAUSER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// PAUSERROLE is a free data retrieval call binding the contract method 0xe63ab1e9.
//
// Solidity: function PAUSER_ROLE() view returns(bytes32)
func (_KickOffLeague *KickOffLeagueSession) PAUSERROLE() ([32]byte, error) {
	return _KickOffLeague.Contract.PAUSERROLE(&_KickOffLeague.CallOpts)
}

// PAUSERROLE is a free data retrieval call binding the contract method 0xe63ab1e9.
//
// Solidity: function PAUSER_ROLE() view returns(bytes32)
func (_KickOffLeague *KickOffLeagueCallerSession) PAUSERROLE() ([32]byte, error) {
	return _KickOffLeague.Contract.PAUSERROLE(&_KickOffLeague.CallOpts)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_KickOffLeague *KickOffLeagueCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _KickOffLeague.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_KickOffLeague *KickOffLeagueSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _KickOffLeague.Contract.GetRoleAdmin(&_KickOffLeague.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_KickOffLeague *KickOffLeagueCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _KickOffLeague.Contract.GetRoleAdmin(&_KickOffLeague.CallOpts, role)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_KickOffLeague *KickOffLeagueCaller) GetRoleMember(opts *bind.CallOpts, role [32]byte, index *big.Int) (common.Address, error) {
	var out []interface{}
	err := _KickOffLeague.contract.Call(opts, &out, "getRoleMember", role, index)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_KickOffLeague *KickOffLeagueSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _KickOffLeague.Contract.GetRoleMember(&_KickOffLeague.CallOpts, role, index)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_KickOffLeague *KickOffLeagueCallerSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _KickOffLeague.Contract.GetRoleMember(&_KickOffLeague.CallOpts, role, index)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_KickOffLeague *KickOffLeagueCaller) GetRoleMemberCount(opts *bind.CallOpts, role [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _KickOffLeague.contract.Call(opts, &out, "getRoleMemberCount", role)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_KickOffLeague *KickOffLeagueSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _KickOffLeague.Contract.GetRoleMemberCount(&_KickOffLeague.CallOpts, role)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_KickOffLeague *KickOffLeagueCallerSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _KickOffLeague.Contract.GetRoleMemberCount(&_KickOffLeague.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_KickOffLeague *KickOffLeagueCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _KickOffLeague.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_KickOffLeague *KickOffLeagueSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _KickOffLeague.Contract.HasRole(&_KickOffLeague.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_KickOffLeague *KickOffLeagueCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _KickOffLeague.Contract.HasRole(&_KickOffLeague.CallOpts, role, account)
}

// LeagueName is a free data retrieval call binding the contract method 0x841abf6e.
//
// Solidity: function leagueName() view returns(string)
func (_KickOffLeague *KickOffLeagueCaller) LeagueName(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _KickOffLeague.contract.Call(opts, &out, "leagueName")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// LeagueName is a free data retrieval call binding the contract method 0x841abf6e.
//
// Solidity: function leagueName() view returns(string)
func (_KickOffLeague *KickOffLeagueSession) LeagueName() (string, error) {
	return _KickOffLeague.Contract.LeagueName(&_KickOffLeague.CallOpts)
}

// LeagueName is a free data retrieval call binding the contract method 0x841abf6e.
//
// Solidity: function leagueName() view returns(string)
func (_KickOffLeague *KickOffLeagueCallerSession) LeagueName() (string, error) {
	return _KickOffLeague.Contract.LeagueName(&_KickOffLeague.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_KickOffLeague *KickOffLeagueCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _KickOffLeague.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_KickOffLeague *KickOffLeagueSession) Owner() (common.Address, error) {
	return _KickOffLeague.Contract.Owner(&_KickOffLeague.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_KickOffLeague *KickOffLeagueCallerSession) Owner() (common.Address, error) {
	return _KickOffLeague.Contract.Owner(&_KickOffLeague.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_KickOffLeague *KickOffLeagueCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _KickOffLeague.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_KickOffLeague *KickOffLeagueSession) Paused() (bool, error) {
	return _KickOffLeague.Contract.Paused(&_KickOffLeague.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_KickOffLeague *KickOffLeagueCallerSession) Paused() (bool, error) {
	return _KickOffLeague.Contract.Paused(&_KickOffLeague.CallOpts)
}

// RaceResult is a free data retrieval call binding the contract method 0x6d93b958.
//
// Solidity: function raceResult(bytes32 raceId) view returns(bytes32 result)
func (_KickOffLeague *KickOffLeagueCaller) RaceResult(opts *bind.CallOpts, raceId [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _KickOffLeague.contract.Call(opts, &out, "raceResult", raceId)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// RaceResult is a free data retrieval call binding the contract method 0x6d93b958.
//
// Solidity: function raceResult(bytes32 raceId) view returns(bytes32 result)
func (_KickOffLeague *KickOffLeagueSession) RaceResult(raceId [32]byte) ([32]byte, error) {
	return _KickOffLeague.Contract.RaceResult(&_KickOffLeague.CallOpts, raceId)
}

// RaceResult is a free data retrieval call binding the contract method 0x6d93b958.
//
// Solidity: function raceResult(bytes32 raceId) view returns(bytes32 result)
func (_KickOffLeague *KickOffLeagueCallerSession) RaceResult(raceId [32]byte) ([32]byte, error) {
	return _KickOffLeague.Contract.RaceResult(&_KickOffLeague.CallOpts, raceId)
}

// RaceSlots is a free data retrieval call binding the contract method 0xd0cbb2a4.
//
// Solidity: function raceSlots(bytes32 raceId) view returns(uint256 slots)
func (_KickOffLeague *KickOffLeagueCaller) RaceSlots(opts *bind.CallOpts, raceId [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _KickOffLeague.contract.Call(opts, &out, "raceSlots", raceId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RaceSlots is a free data retrieval call binding the contract method 0xd0cbb2a4.
//
// Solidity: function raceSlots(bytes32 raceId) view returns(uint256 slots)
func (_KickOffLeague *KickOffLeagueSession) RaceSlots(raceId [32]byte) (*big.Int, error) {
	return _KickOffLeague.Contract.RaceSlots(&_KickOffLeague.CallOpts, raceId)
}

// RaceSlots is a free data retrieval call binding the contract method 0xd0cbb2a4.
//
// Solidity: function raceSlots(bytes32 raceId) view returns(uint256 slots)
func (_KickOffLeague *KickOffLeagueCallerSession) RaceSlots(raceId [32]byte) (*big.Int, error) {
	return _KickOffLeague.Contract.RaceSlots(&_KickOffLeague.CallOpts, raceId)
}

// RaceStartAt is a free data retrieval call binding the contract method 0x1d88e1bd.
//
// Solidity: function raceStartAt(bytes32 raceId) view returns(uint256 startAt)
func (_KickOffLeague *KickOffLeagueCaller) RaceStartAt(opts *bind.CallOpts, raceId [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _KickOffLeague.contract.Call(opts, &out, "raceStartAt", raceId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RaceStartAt is a free data retrieval call binding the contract method 0x1d88e1bd.
//
// Solidity: function raceStartAt(bytes32 raceId) view returns(uint256 startAt)
func (_KickOffLeague *KickOffLeagueSession) RaceStartAt(raceId [32]byte) (*big.Int, error) {
	return _KickOffLeague.Contract.RaceStartAt(&_KickOffLeague.CallOpts, raceId)
}

// RaceStartAt is a free data retrieval call binding the contract method 0x1d88e1bd.
//
// Solidity: function raceStartAt(bytes32 raceId) view returns(uint256 startAt)
func (_KickOffLeague *KickOffLeagueCallerSession) RaceStartAt(raceId [32]byte) (*big.Int, error) {
	return _KickOffLeague.Contract.RaceStartAt(&_KickOffLeague.CallOpts, raceId)
}

// RegisterIsValid is a free data retrieval call binding the contract method 0x53eeece1.
//
// Solidity: function registerIsValid(address register) view returns(bool)
func (_KickOffLeague *KickOffLeagueCaller) RegisterIsValid(opts *bind.CallOpts, register common.Address) (bool, error) {
	var out []interface{}
	err := _KickOffLeague.contract.Call(opts, &out, "registerIsValid", register)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// RegisterIsValid is a free data retrieval call binding the contract method 0x53eeece1.
//
// Solidity: function registerIsValid(address register) view returns(bool)
func (_KickOffLeague *KickOffLeagueSession) RegisterIsValid(register common.Address) (bool, error) {
	return _KickOffLeague.Contract.RegisterIsValid(&_KickOffLeague.CallOpts, register)
}

// RegisterIsValid is a free data retrieval call binding the contract method 0x53eeece1.
//
// Solidity: function registerIsValid(address register) view returns(bool)
func (_KickOffLeague *KickOffLeagueCallerSession) RegisterIsValid(register common.Address) (bool, error) {
	return _KickOffLeague.Contract.RegisterIsValid(&_KickOffLeague.CallOpts, register)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_KickOffLeague *KickOffLeagueCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _KickOffLeague.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_KickOffLeague *KickOffLeagueSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _KickOffLeague.Contract.SupportsInterface(&_KickOffLeague.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_KickOffLeague *KickOffLeagueCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _KickOffLeague.Contract.SupportsInterface(&_KickOffLeague.CallOpts, interfaceId)
}

// CancelRace is a paid mutator transaction binding the contract method 0xc05c107d.
//
// Solidity: function cancelRace(bytes32 raceId) returns()
func (_KickOffLeague *KickOffLeagueTransactor) CancelRace(opts *bind.TransactOpts, raceId [32]byte) (*types.Transaction, error) {
	return _KickOffLeague.contract.Transact(opts, "cancelRace", raceId)
}

// CancelRace is a paid mutator transaction binding the contract method 0xc05c107d.
//
// Solidity: function cancelRace(bytes32 raceId) returns()
func (_KickOffLeague *KickOffLeagueSession) CancelRace(raceId [32]byte) (*types.Transaction, error) {
	return _KickOffLeague.Contract.CancelRace(&_KickOffLeague.TransactOpts, raceId)
}

// CancelRace is a paid mutator transaction binding the contract method 0xc05c107d.
//
// Solidity: function cancelRace(bytes32 raceId) returns()
func (_KickOffLeague *KickOffLeagueTransactorSession) CancelRace(raceId [32]byte) (*types.Transaction, error) {
	return _KickOffLeague.Contract.CancelRace(&_KickOffLeague.TransactOpts, raceId)
}

// CreateRace is a paid mutator transaction binding the contract method 0x420f0351.
//
// Solidity: function createRace(uint256 slots, uint256 startAt) returns(bytes32)
func (_KickOffLeague *KickOffLeagueTransactor) CreateRace(opts *bind.TransactOpts, slots *big.Int, startAt *big.Int) (*types.Transaction, error) {
	return _KickOffLeague.contract.Transact(opts, "createRace", slots, startAt)
}

// CreateRace is a paid mutator transaction binding the contract method 0x420f0351.
//
// Solidity: function createRace(uint256 slots, uint256 startAt) returns(bytes32)
func (_KickOffLeague *KickOffLeagueSession) CreateRace(slots *big.Int, startAt *big.Int) (*types.Transaction, error) {
	return _KickOffLeague.Contract.CreateRace(&_KickOffLeague.TransactOpts, slots, startAt)
}

// CreateRace is a paid mutator transaction binding the contract method 0x420f0351.
//
// Solidity: function createRace(uint256 slots, uint256 startAt) returns(bytes32)
func (_KickOffLeague *KickOffLeagueTransactorSession) CreateRace(slots *big.Int, startAt *big.Int) (*types.Transaction, error) {
	return _KickOffLeague.Contract.CreateRace(&_KickOffLeague.TransactOpts, slots, startAt)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_KickOffLeague *KickOffLeagueTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _KickOffLeague.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_KickOffLeague *KickOffLeagueSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _KickOffLeague.Contract.GrantRole(&_KickOffLeague.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_KickOffLeague *KickOffLeagueTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _KickOffLeague.Contract.GrantRole(&_KickOffLeague.TransactOpts, role, account)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_KickOffLeague *KickOffLeagueTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _KickOffLeague.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_KickOffLeague *KickOffLeagueSession) Pause() (*types.Transaction, error) {
	return _KickOffLeague.Contract.Pause(&_KickOffLeague.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_KickOffLeague *KickOffLeagueTransactorSession) Pause() (*types.Transaction, error) {
	return _KickOffLeague.Contract.Pause(&_KickOffLeague.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_KickOffLeague *KickOffLeagueTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _KickOffLeague.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_KickOffLeague *KickOffLeagueSession) RenounceOwnership() (*types.Transaction, error) {
	return _KickOffLeague.Contract.RenounceOwnership(&_KickOffLeague.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_KickOffLeague *KickOffLeagueTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _KickOffLeague.Contract.RenounceOwnership(&_KickOffLeague.TransactOpts)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_KickOffLeague *KickOffLeagueTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _KickOffLeague.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_KickOffLeague *KickOffLeagueSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _KickOffLeague.Contract.RenounceRole(&_KickOffLeague.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_KickOffLeague *KickOffLeagueTransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _KickOffLeague.Contract.RenounceRole(&_KickOffLeague.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_KickOffLeague *KickOffLeagueTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _KickOffLeague.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_KickOffLeague *KickOffLeagueSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _KickOffLeague.Contract.RevokeRole(&_KickOffLeague.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_KickOffLeague *KickOffLeagueTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _KickOffLeague.Contract.RevokeRole(&_KickOffLeague.TransactOpts, role, account)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_KickOffLeague *KickOffLeagueTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _KickOffLeague.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_KickOffLeague *KickOffLeagueSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _KickOffLeague.Contract.TransferOwnership(&_KickOffLeague.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_KickOffLeague *KickOffLeagueTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _KickOffLeague.Contract.TransferOwnership(&_KickOffLeague.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_KickOffLeague *KickOffLeagueTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _KickOffLeague.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_KickOffLeague *KickOffLeagueSession) Unpause() (*types.Transaction, error) {
	return _KickOffLeague.Contract.Unpause(&_KickOffLeague.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_KickOffLeague *KickOffLeagueTransactorSession) Unpause() (*types.Transaction, error) {
	return _KickOffLeague.Contract.Unpause(&_KickOffLeague.TransactOpts)
}

// UpdateRaceResult is a paid mutator transaction binding the contract method 0x8c9d6dbd.
//
// Solidity: function updateRaceResult(bytes32 raceId, bytes32 result) returns()
func (_KickOffLeague *KickOffLeagueTransactor) UpdateRaceResult(opts *bind.TransactOpts, raceId [32]byte, result [32]byte) (*types.Transaction, error) {
	return _KickOffLeague.contract.Transact(opts, "updateRaceResult", raceId, result)
}

// UpdateRaceResult is a paid mutator transaction binding the contract method 0x8c9d6dbd.
//
// Solidity: function updateRaceResult(bytes32 raceId, bytes32 result) returns()
func (_KickOffLeague *KickOffLeagueSession) UpdateRaceResult(raceId [32]byte, result [32]byte) (*types.Transaction, error) {
	return _KickOffLeague.Contract.UpdateRaceResult(&_KickOffLeague.TransactOpts, raceId, result)
}

// UpdateRaceResult is a paid mutator transaction binding the contract method 0x8c9d6dbd.
//
// Solidity: function updateRaceResult(bytes32 raceId, bytes32 result) returns()
func (_KickOffLeague *KickOffLeagueTransactorSession) UpdateRaceResult(raceId [32]byte, result [32]byte) (*types.Transaction, error) {
	return _KickOffLeague.Contract.UpdateRaceResult(&_KickOffLeague.TransactOpts, raceId, result)
}

// KickOffLeagueOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the KickOffLeague contract.
type KickOffLeagueOwnershipTransferredIterator struct {
	Event *KickOffLeagueOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *KickOffLeagueOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KickOffLeagueOwnershipTransferred)
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
		it.Event = new(KickOffLeagueOwnershipTransferred)
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
func (it *KickOffLeagueOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KickOffLeagueOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KickOffLeagueOwnershipTransferred represents a OwnershipTransferred event raised by the KickOffLeague contract.
type KickOffLeagueOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_KickOffLeague *KickOffLeagueFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*KickOffLeagueOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _KickOffLeague.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &KickOffLeagueOwnershipTransferredIterator{contract: _KickOffLeague.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_KickOffLeague *KickOffLeagueFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *KickOffLeagueOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _KickOffLeague.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KickOffLeagueOwnershipTransferred)
				if err := _KickOffLeague.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_KickOffLeague *KickOffLeagueFilterer) ParseOwnershipTransferred(log types.Log) (*KickOffLeagueOwnershipTransferred, error) {
	event := new(KickOffLeagueOwnershipTransferred)
	if err := _KickOffLeague.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// KickOffLeaguePausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the KickOffLeague contract.
type KickOffLeaguePausedIterator struct {
	Event *KickOffLeaguePaused // Event containing the contract specifics and raw log

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
func (it *KickOffLeaguePausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KickOffLeaguePaused)
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
		it.Event = new(KickOffLeaguePaused)
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
func (it *KickOffLeaguePausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KickOffLeaguePausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KickOffLeaguePaused represents a Paused event raised by the KickOffLeague contract.
type KickOffLeaguePaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_KickOffLeague *KickOffLeagueFilterer) FilterPaused(opts *bind.FilterOpts) (*KickOffLeaguePausedIterator, error) {

	logs, sub, err := _KickOffLeague.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &KickOffLeaguePausedIterator{contract: _KickOffLeague.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_KickOffLeague *KickOffLeagueFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *KickOffLeaguePaused) (event.Subscription, error) {

	logs, sub, err := _KickOffLeague.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KickOffLeaguePaused)
				if err := _KickOffLeague.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_KickOffLeague *KickOffLeagueFilterer) ParsePaused(log types.Log) (*KickOffLeaguePaused, error) {
	event := new(KickOffLeaguePaused)
	if err := _KickOffLeague.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// KickOffLeagueRaceCancelledIterator is returned from FilterRaceCancelled and is used to iterate over the raw logs and unpacked data for RaceCancelled events raised by the KickOffLeague contract.
type KickOffLeagueRaceCancelledIterator struct {
	Event *KickOffLeagueRaceCancelled // Event containing the contract specifics and raw log

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
func (it *KickOffLeagueRaceCancelledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KickOffLeagueRaceCancelled)
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
		it.Event = new(KickOffLeagueRaceCancelled)
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
func (it *KickOffLeagueRaceCancelledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KickOffLeagueRaceCancelledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KickOffLeagueRaceCancelled represents a RaceCancelled event raised by the KickOffLeague contract.
type KickOffLeagueRaceCancelled struct {
	Id  [32]byte
	Raw types.Log // Blockchain specific contextual infos
}

// FilterRaceCancelled is a free log retrieval operation binding the contract event 0x3f3669b11af9b00fb2a95f8ded36a79a2e1d5374efbc5ec6b2b5b96ceaf9ac6a.
//
// Solidity: event RaceCancelled(bytes32 id)
func (_KickOffLeague *KickOffLeagueFilterer) FilterRaceCancelled(opts *bind.FilterOpts) (*KickOffLeagueRaceCancelledIterator, error) {

	logs, sub, err := _KickOffLeague.contract.FilterLogs(opts, "RaceCancelled")
	if err != nil {
		return nil, err
	}
	return &KickOffLeagueRaceCancelledIterator{contract: _KickOffLeague.contract, event: "RaceCancelled", logs: logs, sub: sub}, nil
}

// WatchRaceCancelled is a free log subscription operation binding the contract event 0x3f3669b11af9b00fb2a95f8ded36a79a2e1d5374efbc5ec6b2b5b96ceaf9ac6a.
//
// Solidity: event RaceCancelled(bytes32 id)
func (_KickOffLeague *KickOffLeagueFilterer) WatchRaceCancelled(opts *bind.WatchOpts, sink chan<- *KickOffLeagueRaceCancelled) (event.Subscription, error) {

	logs, sub, err := _KickOffLeague.contract.WatchLogs(opts, "RaceCancelled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KickOffLeagueRaceCancelled)
				if err := _KickOffLeague.contract.UnpackLog(event, "RaceCancelled", log); err != nil {
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

// ParseRaceCancelled is a log parse operation binding the contract event 0x3f3669b11af9b00fb2a95f8ded36a79a2e1d5374efbc5ec6b2b5b96ceaf9ac6a.
//
// Solidity: event RaceCancelled(bytes32 id)
func (_KickOffLeague *KickOffLeagueFilterer) ParseRaceCancelled(log types.Log) (*KickOffLeagueRaceCancelled, error) {
	event := new(KickOffLeagueRaceCancelled)
	if err := _KickOffLeague.contract.UnpackLog(event, "RaceCancelled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// KickOffLeagueRaceCreatedIterator is returned from FilterRaceCreated and is used to iterate over the raw logs and unpacked data for RaceCreated events raised by the KickOffLeague contract.
type KickOffLeagueRaceCreatedIterator struct {
	Event *KickOffLeagueRaceCreated // Event containing the contract specifics and raw log

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
func (it *KickOffLeagueRaceCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KickOffLeagueRaceCreated)
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
		it.Event = new(KickOffLeagueRaceCreated)
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
func (it *KickOffLeagueRaceCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KickOffLeagueRaceCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KickOffLeagueRaceCreated represents a RaceCreated event raised by the KickOffLeague contract.
type KickOffLeagueRaceCreated struct {
	Id      [32]byte
	Slots   *big.Int
	StartAt *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRaceCreated is a free log retrieval operation binding the contract event 0x3d4c26760d0e4825f9a3f6ccba234d54aeeae3e6471639b8366708a183f1e837.
//
// Solidity: event RaceCreated(bytes32 id, uint256 slots, uint256 startAt)
func (_KickOffLeague *KickOffLeagueFilterer) FilterRaceCreated(opts *bind.FilterOpts) (*KickOffLeagueRaceCreatedIterator, error) {

	logs, sub, err := _KickOffLeague.contract.FilterLogs(opts, "RaceCreated")
	if err != nil {
		return nil, err
	}
	return &KickOffLeagueRaceCreatedIterator{contract: _KickOffLeague.contract, event: "RaceCreated", logs: logs, sub: sub}, nil
}

// WatchRaceCreated is a free log subscription operation binding the contract event 0x3d4c26760d0e4825f9a3f6ccba234d54aeeae3e6471639b8366708a183f1e837.
//
// Solidity: event RaceCreated(bytes32 id, uint256 slots, uint256 startAt)
func (_KickOffLeague *KickOffLeagueFilterer) WatchRaceCreated(opts *bind.WatchOpts, sink chan<- *KickOffLeagueRaceCreated) (event.Subscription, error) {

	logs, sub, err := _KickOffLeague.contract.WatchLogs(opts, "RaceCreated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KickOffLeagueRaceCreated)
				if err := _KickOffLeague.contract.UnpackLog(event, "RaceCreated", log); err != nil {
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

// ParseRaceCreated is a log parse operation binding the contract event 0x3d4c26760d0e4825f9a3f6ccba234d54aeeae3e6471639b8366708a183f1e837.
//
// Solidity: event RaceCreated(bytes32 id, uint256 slots, uint256 startAt)
func (_KickOffLeague *KickOffLeagueFilterer) ParseRaceCreated(log types.Log) (*KickOffLeagueRaceCreated, error) {
	event := new(KickOffLeagueRaceCreated)
	if err := _KickOffLeague.contract.UnpackLog(event, "RaceCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// KickOffLeagueRaceResultUpdatedIterator is returned from FilterRaceResultUpdated and is used to iterate over the raw logs and unpacked data for RaceResultUpdated events raised by the KickOffLeague contract.
type KickOffLeagueRaceResultUpdatedIterator struct {
	Event *KickOffLeagueRaceResultUpdated // Event containing the contract specifics and raw log

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
func (it *KickOffLeagueRaceResultUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KickOffLeagueRaceResultUpdated)
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
		it.Event = new(KickOffLeagueRaceResultUpdated)
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
func (it *KickOffLeagueRaceResultUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KickOffLeagueRaceResultUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KickOffLeagueRaceResultUpdated represents a RaceResultUpdated event raised by the KickOffLeague contract.
type KickOffLeagueRaceResultUpdated struct {
	Id     [32]byte
	Result [32]byte
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRaceResultUpdated is a free log retrieval operation binding the contract event 0xd82663592968d73ce1995154e44c793fcb46a4006abfb7438656fb0d7ba5ff49.
//
// Solidity: event RaceResultUpdated(bytes32 id, bytes32 result)
func (_KickOffLeague *KickOffLeagueFilterer) FilterRaceResultUpdated(opts *bind.FilterOpts) (*KickOffLeagueRaceResultUpdatedIterator, error) {

	logs, sub, err := _KickOffLeague.contract.FilterLogs(opts, "RaceResultUpdated")
	if err != nil {
		return nil, err
	}
	return &KickOffLeagueRaceResultUpdatedIterator{contract: _KickOffLeague.contract, event: "RaceResultUpdated", logs: logs, sub: sub}, nil
}

// WatchRaceResultUpdated is a free log subscription operation binding the contract event 0xd82663592968d73ce1995154e44c793fcb46a4006abfb7438656fb0d7ba5ff49.
//
// Solidity: event RaceResultUpdated(bytes32 id, bytes32 result)
func (_KickOffLeague *KickOffLeagueFilterer) WatchRaceResultUpdated(opts *bind.WatchOpts, sink chan<- *KickOffLeagueRaceResultUpdated) (event.Subscription, error) {

	logs, sub, err := _KickOffLeague.contract.WatchLogs(opts, "RaceResultUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KickOffLeagueRaceResultUpdated)
				if err := _KickOffLeague.contract.UnpackLog(event, "RaceResultUpdated", log); err != nil {
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

// ParseRaceResultUpdated is a log parse operation binding the contract event 0xd82663592968d73ce1995154e44c793fcb46a4006abfb7438656fb0d7ba5ff49.
//
// Solidity: event RaceResultUpdated(bytes32 id, bytes32 result)
func (_KickOffLeague *KickOffLeagueFilterer) ParseRaceResultUpdated(log types.Log) (*KickOffLeagueRaceResultUpdated, error) {
	event := new(KickOffLeagueRaceResultUpdated)
	if err := _KickOffLeague.contract.UnpackLog(event, "RaceResultUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// KickOffLeagueRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the KickOffLeague contract.
type KickOffLeagueRoleAdminChangedIterator struct {
	Event *KickOffLeagueRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *KickOffLeagueRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KickOffLeagueRoleAdminChanged)
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
		it.Event = new(KickOffLeagueRoleAdminChanged)
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
func (it *KickOffLeagueRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KickOffLeagueRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KickOffLeagueRoleAdminChanged represents a RoleAdminChanged event raised by the KickOffLeague contract.
type KickOffLeagueRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_KickOffLeague *KickOffLeagueFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*KickOffLeagueRoleAdminChangedIterator, error) {

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

	logs, sub, err := _KickOffLeague.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &KickOffLeagueRoleAdminChangedIterator{contract: _KickOffLeague.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_KickOffLeague *KickOffLeagueFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *KickOffLeagueRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _KickOffLeague.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KickOffLeagueRoleAdminChanged)
				if err := _KickOffLeague.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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
func (_KickOffLeague *KickOffLeagueFilterer) ParseRoleAdminChanged(log types.Log) (*KickOffLeagueRoleAdminChanged, error) {
	event := new(KickOffLeagueRoleAdminChanged)
	if err := _KickOffLeague.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// KickOffLeagueRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the KickOffLeague contract.
type KickOffLeagueRoleGrantedIterator struct {
	Event *KickOffLeagueRoleGranted // Event containing the contract specifics and raw log

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
func (it *KickOffLeagueRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KickOffLeagueRoleGranted)
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
		it.Event = new(KickOffLeagueRoleGranted)
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
func (it *KickOffLeagueRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KickOffLeagueRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KickOffLeagueRoleGranted represents a RoleGranted event raised by the KickOffLeague contract.
type KickOffLeagueRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_KickOffLeague *KickOffLeagueFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*KickOffLeagueRoleGrantedIterator, error) {

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

	logs, sub, err := _KickOffLeague.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &KickOffLeagueRoleGrantedIterator{contract: _KickOffLeague.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_KickOffLeague *KickOffLeagueFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *KickOffLeagueRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _KickOffLeague.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KickOffLeagueRoleGranted)
				if err := _KickOffLeague.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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
func (_KickOffLeague *KickOffLeagueFilterer) ParseRoleGranted(log types.Log) (*KickOffLeagueRoleGranted, error) {
	event := new(KickOffLeagueRoleGranted)
	if err := _KickOffLeague.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// KickOffLeagueRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the KickOffLeague contract.
type KickOffLeagueRoleRevokedIterator struct {
	Event *KickOffLeagueRoleRevoked // Event containing the contract specifics and raw log

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
func (it *KickOffLeagueRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KickOffLeagueRoleRevoked)
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
		it.Event = new(KickOffLeagueRoleRevoked)
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
func (it *KickOffLeagueRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KickOffLeagueRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KickOffLeagueRoleRevoked represents a RoleRevoked event raised by the KickOffLeague contract.
type KickOffLeagueRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_KickOffLeague *KickOffLeagueFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*KickOffLeagueRoleRevokedIterator, error) {

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

	logs, sub, err := _KickOffLeague.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &KickOffLeagueRoleRevokedIterator{contract: _KickOffLeague.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_KickOffLeague *KickOffLeagueFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *KickOffLeagueRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _KickOffLeague.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KickOffLeagueRoleRevoked)
				if err := _KickOffLeague.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
func (_KickOffLeague *KickOffLeagueFilterer) ParseRoleRevoked(log types.Log) (*KickOffLeagueRoleRevoked, error) {
	event := new(KickOffLeagueRoleRevoked)
	if err := _KickOffLeague.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// KickOffLeagueUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the KickOffLeague contract.
type KickOffLeagueUnpausedIterator struct {
	Event *KickOffLeagueUnpaused // Event containing the contract specifics and raw log

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
func (it *KickOffLeagueUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KickOffLeagueUnpaused)
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
		it.Event = new(KickOffLeagueUnpaused)
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
func (it *KickOffLeagueUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KickOffLeagueUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KickOffLeagueUnpaused represents a Unpaused event raised by the KickOffLeague contract.
type KickOffLeagueUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_KickOffLeague *KickOffLeagueFilterer) FilterUnpaused(opts *bind.FilterOpts) (*KickOffLeagueUnpausedIterator, error) {

	logs, sub, err := _KickOffLeague.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &KickOffLeagueUnpausedIterator{contract: _KickOffLeague.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_KickOffLeague *KickOffLeagueFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *KickOffLeagueUnpaused) (event.Subscription, error) {

	logs, sub, err := _KickOffLeague.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KickOffLeagueUnpaused)
				if err := _KickOffLeague.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
func (_KickOffLeague *KickOffLeagueFilterer) ParseUnpaused(log types.Log) (*KickOffLeagueUnpaused, error) {
	event := new(KickOffLeagueUnpaused)
	if err := _KickOffLeague.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
