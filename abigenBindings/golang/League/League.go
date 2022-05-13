// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package League

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

// LeagueMetaData contains all meta data concerning the League contract.
var LeagueMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"CannotCancel\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidSlot\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RaceExisted\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RaceNotExisted\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"}],\"name\":\"TooEarly\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"}],\"name\":\"TooLate\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"}],\"name\":\"RaceCancelled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"slots\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"startAt\",\"type\":\"uint256\"}],\"name\":\"RaceCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"result\",\"type\":\"bytes32\"}],\"name\":\"RaceResultUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PAUSER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getRoleMember\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleMemberCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"raceId\",\"type\":\"bytes32\"}],\"name\":\"raceSlots\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"slots\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"raceId\",\"type\":\"bytes32\"}],\"name\":\"raceStartAt\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"startAt\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"raceId\",\"type\":\"bytes32\"}],\"name\":\"raceResult\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"result\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"leagueName\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"slots\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startAt\",\"type\":\"uint256\"}],\"name\":\"createRace\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"raceId\",\"type\":\"bytes32\"}],\"name\":\"cancelRace\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"raceId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"result\",\"type\":\"bytes32\"}],\"name\":\"updateRaceResult\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// LeagueABI is the input ABI used to generate the binding from.
// Deprecated: Use LeagueMetaData.ABI instead.
var LeagueABI = LeagueMetaData.ABI

// League is an auto generated Go binding around an Ethereum contract.
type League struct {
	LeagueCaller     // Read-only binding to the contract
	LeagueTransactor // Write-only binding to the contract
	LeagueFilterer   // Log filterer for contract events
}

// LeagueCaller is an auto generated read-only Go binding around an Ethereum contract.
type LeagueCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LeagueTransactor is an auto generated write-only Go binding around an Ethereum contract.
type LeagueTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LeagueFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type LeagueFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LeagueSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type LeagueSession struct {
	Contract     *League           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// LeagueCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type LeagueCallerSession struct {
	Contract *LeagueCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// LeagueTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type LeagueTransactorSession struct {
	Contract     *LeagueTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// LeagueRaw is an auto generated low-level Go binding around an Ethereum contract.
type LeagueRaw struct {
	Contract *League // Generic contract binding to access the raw methods on
}

// LeagueCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type LeagueCallerRaw struct {
	Contract *LeagueCaller // Generic read-only contract binding to access the raw methods on
}

// LeagueTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type LeagueTransactorRaw struct {
	Contract *LeagueTransactor // Generic write-only contract binding to access the raw methods on
}

// NewLeague creates a new instance of League, bound to a specific deployed contract.
func NewLeague(address common.Address, backend bind.ContractBackend) (*League, error) {
	contract, err := bindLeague(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &League{LeagueCaller: LeagueCaller{contract: contract}, LeagueTransactor: LeagueTransactor{contract: contract}, LeagueFilterer: LeagueFilterer{contract: contract}}, nil
}

// NewLeagueCaller creates a new read-only instance of League, bound to a specific deployed contract.
func NewLeagueCaller(address common.Address, caller bind.ContractCaller) (*LeagueCaller, error) {
	contract, err := bindLeague(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &LeagueCaller{contract: contract}, nil
}

// NewLeagueTransactor creates a new write-only instance of League, bound to a specific deployed contract.
func NewLeagueTransactor(address common.Address, transactor bind.ContractTransactor) (*LeagueTransactor, error) {
	contract, err := bindLeague(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &LeagueTransactor{contract: contract}, nil
}

// NewLeagueFilterer creates a new log filterer instance of League, bound to a specific deployed contract.
func NewLeagueFilterer(address common.Address, filterer bind.ContractFilterer) (*LeagueFilterer, error) {
	contract, err := bindLeague(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &LeagueFilterer{contract: contract}, nil
}

// bindLeague binds a generic wrapper to an already deployed contract.
func bindLeague(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(LeagueABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_League *LeagueRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _League.Contract.LeagueCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_League *LeagueRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _League.Contract.LeagueTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_League *LeagueRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _League.Contract.LeagueTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_League *LeagueCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _League.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_League *LeagueTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _League.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_League *LeagueTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _League.Contract.contract.Transact(opts, method, params...)
}

// ADMINROLE is a free data retrieval call binding the contract method 0x75b238fc.
//
// Solidity: function ADMIN_ROLE() view returns(bytes32)
func (_League *LeagueCaller) ADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _League.contract.Call(opts, &out, "ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ADMINROLE is a free data retrieval call binding the contract method 0x75b238fc.
//
// Solidity: function ADMIN_ROLE() view returns(bytes32)
func (_League *LeagueSession) ADMINROLE() ([32]byte, error) {
	return _League.Contract.ADMINROLE(&_League.CallOpts)
}

// ADMINROLE is a free data retrieval call binding the contract method 0x75b238fc.
//
// Solidity: function ADMIN_ROLE() view returns(bytes32)
func (_League *LeagueCallerSession) ADMINROLE() ([32]byte, error) {
	return _League.Contract.ADMINROLE(&_League.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_League *LeagueCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _League.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_League *LeagueSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _League.Contract.DEFAULTADMINROLE(&_League.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_League *LeagueCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _League.Contract.DEFAULTADMINROLE(&_League.CallOpts)
}

// PAUSERROLE is a free data retrieval call binding the contract method 0xe63ab1e9.
//
// Solidity: function PAUSER_ROLE() view returns(bytes32)
func (_League *LeagueCaller) PAUSERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _League.contract.Call(opts, &out, "PAUSER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// PAUSERROLE is a free data retrieval call binding the contract method 0xe63ab1e9.
//
// Solidity: function PAUSER_ROLE() view returns(bytes32)
func (_League *LeagueSession) PAUSERROLE() ([32]byte, error) {
	return _League.Contract.PAUSERROLE(&_League.CallOpts)
}

// PAUSERROLE is a free data retrieval call binding the contract method 0xe63ab1e9.
//
// Solidity: function PAUSER_ROLE() view returns(bytes32)
func (_League *LeagueCallerSession) PAUSERROLE() ([32]byte, error) {
	return _League.Contract.PAUSERROLE(&_League.CallOpts)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_League *LeagueCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _League.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_League *LeagueSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _League.Contract.GetRoleAdmin(&_League.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_League *LeagueCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _League.Contract.GetRoleAdmin(&_League.CallOpts, role)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_League *LeagueCaller) GetRoleMember(opts *bind.CallOpts, role [32]byte, index *big.Int) (common.Address, error) {
	var out []interface{}
	err := _League.contract.Call(opts, &out, "getRoleMember", role, index)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_League *LeagueSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _League.Contract.GetRoleMember(&_League.CallOpts, role, index)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_League *LeagueCallerSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _League.Contract.GetRoleMember(&_League.CallOpts, role, index)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_League *LeagueCaller) GetRoleMemberCount(opts *bind.CallOpts, role [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _League.contract.Call(opts, &out, "getRoleMemberCount", role)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_League *LeagueSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _League.Contract.GetRoleMemberCount(&_League.CallOpts, role)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_League *LeagueCallerSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _League.Contract.GetRoleMemberCount(&_League.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_League *LeagueCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _League.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_League *LeagueSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _League.Contract.HasRole(&_League.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_League *LeagueCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _League.Contract.HasRole(&_League.CallOpts, role, account)
}

// LeagueName is a free data retrieval call binding the contract method 0x841abf6e.
//
// Solidity: function leagueName() view returns(string)
func (_League *LeagueCaller) LeagueName(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _League.contract.Call(opts, &out, "leagueName")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// LeagueName is a free data retrieval call binding the contract method 0x841abf6e.
//
// Solidity: function leagueName() view returns(string)
func (_League *LeagueSession) LeagueName() (string, error) {
	return _League.Contract.LeagueName(&_League.CallOpts)
}

// LeagueName is a free data retrieval call binding the contract method 0x841abf6e.
//
// Solidity: function leagueName() view returns(string)
func (_League *LeagueCallerSession) LeagueName() (string, error) {
	return _League.Contract.LeagueName(&_League.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_League *LeagueCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _League.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_League *LeagueSession) Owner() (common.Address, error) {
	return _League.Contract.Owner(&_League.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_League *LeagueCallerSession) Owner() (common.Address, error) {
	return _League.Contract.Owner(&_League.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_League *LeagueCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _League.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_League *LeagueSession) Paused() (bool, error) {
	return _League.Contract.Paused(&_League.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_League *LeagueCallerSession) Paused() (bool, error) {
	return _League.Contract.Paused(&_League.CallOpts)
}

// RaceResult is a free data retrieval call binding the contract method 0x6d93b958.
//
// Solidity: function raceResult(bytes32 raceId) view returns(bytes32 result)
func (_League *LeagueCaller) RaceResult(opts *bind.CallOpts, raceId [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _League.contract.Call(opts, &out, "raceResult", raceId)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// RaceResult is a free data retrieval call binding the contract method 0x6d93b958.
//
// Solidity: function raceResult(bytes32 raceId) view returns(bytes32 result)
func (_League *LeagueSession) RaceResult(raceId [32]byte) ([32]byte, error) {
	return _League.Contract.RaceResult(&_League.CallOpts, raceId)
}

// RaceResult is a free data retrieval call binding the contract method 0x6d93b958.
//
// Solidity: function raceResult(bytes32 raceId) view returns(bytes32 result)
func (_League *LeagueCallerSession) RaceResult(raceId [32]byte) ([32]byte, error) {
	return _League.Contract.RaceResult(&_League.CallOpts, raceId)
}

// RaceSlots is a free data retrieval call binding the contract method 0xd0cbb2a4.
//
// Solidity: function raceSlots(bytes32 raceId) view returns(uint256 slots)
func (_League *LeagueCaller) RaceSlots(opts *bind.CallOpts, raceId [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _League.contract.Call(opts, &out, "raceSlots", raceId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RaceSlots is a free data retrieval call binding the contract method 0xd0cbb2a4.
//
// Solidity: function raceSlots(bytes32 raceId) view returns(uint256 slots)
func (_League *LeagueSession) RaceSlots(raceId [32]byte) (*big.Int, error) {
	return _League.Contract.RaceSlots(&_League.CallOpts, raceId)
}

// RaceSlots is a free data retrieval call binding the contract method 0xd0cbb2a4.
//
// Solidity: function raceSlots(bytes32 raceId) view returns(uint256 slots)
func (_League *LeagueCallerSession) RaceSlots(raceId [32]byte) (*big.Int, error) {
	return _League.Contract.RaceSlots(&_League.CallOpts, raceId)
}

// RaceStartAt is a free data retrieval call binding the contract method 0x1d88e1bd.
//
// Solidity: function raceStartAt(bytes32 raceId) view returns(uint256 startAt)
func (_League *LeagueCaller) RaceStartAt(opts *bind.CallOpts, raceId [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _League.contract.Call(opts, &out, "raceStartAt", raceId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RaceStartAt is a free data retrieval call binding the contract method 0x1d88e1bd.
//
// Solidity: function raceStartAt(bytes32 raceId) view returns(uint256 startAt)
func (_League *LeagueSession) RaceStartAt(raceId [32]byte) (*big.Int, error) {
	return _League.Contract.RaceStartAt(&_League.CallOpts, raceId)
}

// RaceStartAt is a free data retrieval call binding the contract method 0x1d88e1bd.
//
// Solidity: function raceStartAt(bytes32 raceId) view returns(uint256 startAt)
func (_League *LeagueCallerSession) RaceStartAt(raceId [32]byte) (*big.Int, error) {
	return _League.Contract.RaceStartAt(&_League.CallOpts, raceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_League *LeagueCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _League.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_League *LeagueSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _League.Contract.SupportsInterface(&_League.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_League *LeagueCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _League.Contract.SupportsInterface(&_League.CallOpts, interfaceId)
}

// CancelRace is a paid mutator transaction binding the contract method 0xc05c107d.
//
// Solidity: function cancelRace(bytes32 raceId) returns()
func (_League *LeagueTransactor) CancelRace(opts *bind.TransactOpts, raceId [32]byte) (*types.Transaction, error) {
	return _League.contract.Transact(opts, "cancelRace", raceId)
}

// CancelRace is a paid mutator transaction binding the contract method 0xc05c107d.
//
// Solidity: function cancelRace(bytes32 raceId) returns()
func (_League *LeagueSession) CancelRace(raceId [32]byte) (*types.Transaction, error) {
	return _League.Contract.CancelRace(&_League.TransactOpts, raceId)
}

// CancelRace is a paid mutator transaction binding the contract method 0xc05c107d.
//
// Solidity: function cancelRace(bytes32 raceId) returns()
func (_League *LeagueTransactorSession) CancelRace(raceId [32]byte) (*types.Transaction, error) {
	return _League.Contract.CancelRace(&_League.TransactOpts, raceId)
}

// CreateRace is a paid mutator transaction binding the contract method 0x420f0351.
//
// Solidity: function createRace(uint256 slots, uint256 startAt) returns(bytes32)
func (_League *LeagueTransactor) CreateRace(opts *bind.TransactOpts, slots *big.Int, startAt *big.Int) (*types.Transaction, error) {
	return _League.contract.Transact(opts, "createRace", slots, startAt)
}

// CreateRace is a paid mutator transaction binding the contract method 0x420f0351.
//
// Solidity: function createRace(uint256 slots, uint256 startAt) returns(bytes32)
func (_League *LeagueSession) CreateRace(slots *big.Int, startAt *big.Int) (*types.Transaction, error) {
	return _League.Contract.CreateRace(&_League.TransactOpts, slots, startAt)
}

// CreateRace is a paid mutator transaction binding the contract method 0x420f0351.
//
// Solidity: function createRace(uint256 slots, uint256 startAt) returns(bytes32)
func (_League *LeagueTransactorSession) CreateRace(slots *big.Int, startAt *big.Int) (*types.Transaction, error) {
	return _League.Contract.CreateRace(&_League.TransactOpts, slots, startAt)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_League *LeagueTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _League.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_League *LeagueSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _League.Contract.GrantRole(&_League.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_League *LeagueTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _League.Contract.GrantRole(&_League.TransactOpts, role, account)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_League *LeagueTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _League.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_League *LeagueSession) Pause() (*types.Transaction, error) {
	return _League.Contract.Pause(&_League.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_League *LeagueTransactorSession) Pause() (*types.Transaction, error) {
	return _League.Contract.Pause(&_League.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_League *LeagueTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _League.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_League *LeagueSession) RenounceOwnership() (*types.Transaction, error) {
	return _League.Contract.RenounceOwnership(&_League.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_League *LeagueTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _League.Contract.RenounceOwnership(&_League.TransactOpts)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_League *LeagueTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _League.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_League *LeagueSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _League.Contract.RenounceRole(&_League.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_League *LeagueTransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _League.Contract.RenounceRole(&_League.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_League *LeagueTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _League.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_League *LeagueSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _League.Contract.RevokeRole(&_League.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_League *LeagueTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _League.Contract.RevokeRole(&_League.TransactOpts, role, account)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_League *LeagueTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _League.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_League *LeagueSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _League.Contract.TransferOwnership(&_League.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_League *LeagueTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _League.Contract.TransferOwnership(&_League.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_League *LeagueTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _League.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_League *LeagueSession) Unpause() (*types.Transaction, error) {
	return _League.Contract.Unpause(&_League.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_League *LeagueTransactorSession) Unpause() (*types.Transaction, error) {
	return _League.Contract.Unpause(&_League.TransactOpts)
}

// UpdateRaceResult is a paid mutator transaction binding the contract method 0x8c9d6dbd.
//
// Solidity: function updateRaceResult(bytes32 raceId, bytes32 result) returns()
func (_League *LeagueTransactor) UpdateRaceResult(opts *bind.TransactOpts, raceId [32]byte, result [32]byte) (*types.Transaction, error) {
	return _League.contract.Transact(opts, "updateRaceResult", raceId, result)
}

// UpdateRaceResult is a paid mutator transaction binding the contract method 0x8c9d6dbd.
//
// Solidity: function updateRaceResult(bytes32 raceId, bytes32 result) returns()
func (_League *LeagueSession) UpdateRaceResult(raceId [32]byte, result [32]byte) (*types.Transaction, error) {
	return _League.Contract.UpdateRaceResult(&_League.TransactOpts, raceId, result)
}

// UpdateRaceResult is a paid mutator transaction binding the contract method 0x8c9d6dbd.
//
// Solidity: function updateRaceResult(bytes32 raceId, bytes32 result) returns()
func (_League *LeagueTransactorSession) UpdateRaceResult(raceId [32]byte, result [32]byte) (*types.Transaction, error) {
	return _League.Contract.UpdateRaceResult(&_League.TransactOpts, raceId, result)
}

// LeagueOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the League contract.
type LeagueOwnershipTransferredIterator struct {
	Event *LeagueOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *LeagueOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LeagueOwnershipTransferred)
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
		it.Event = new(LeagueOwnershipTransferred)
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
func (it *LeagueOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LeagueOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LeagueOwnershipTransferred represents a OwnershipTransferred event raised by the League contract.
type LeagueOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_League *LeagueFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*LeagueOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _League.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &LeagueOwnershipTransferredIterator{contract: _League.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_League *LeagueFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *LeagueOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _League.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LeagueOwnershipTransferred)
				if err := _League.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_League *LeagueFilterer) ParseOwnershipTransferred(log types.Log) (*LeagueOwnershipTransferred, error) {
	event := new(LeagueOwnershipTransferred)
	if err := _League.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LeaguePausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the League contract.
type LeaguePausedIterator struct {
	Event *LeaguePaused // Event containing the contract specifics and raw log

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
func (it *LeaguePausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LeaguePaused)
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
		it.Event = new(LeaguePaused)
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
func (it *LeaguePausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LeaguePausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LeaguePaused represents a Paused event raised by the League contract.
type LeaguePaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_League *LeagueFilterer) FilterPaused(opts *bind.FilterOpts) (*LeaguePausedIterator, error) {

	logs, sub, err := _League.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &LeaguePausedIterator{contract: _League.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_League *LeagueFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *LeaguePaused) (event.Subscription, error) {

	logs, sub, err := _League.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LeaguePaused)
				if err := _League.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_League *LeagueFilterer) ParsePaused(log types.Log) (*LeaguePaused, error) {
	event := new(LeaguePaused)
	if err := _League.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LeagueRaceCancelledIterator is returned from FilterRaceCancelled and is used to iterate over the raw logs and unpacked data for RaceCancelled events raised by the League contract.
type LeagueRaceCancelledIterator struct {
	Event *LeagueRaceCancelled // Event containing the contract specifics and raw log

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
func (it *LeagueRaceCancelledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LeagueRaceCancelled)
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
		it.Event = new(LeagueRaceCancelled)
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
func (it *LeagueRaceCancelledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LeagueRaceCancelledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LeagueRaceCancelled represents a RaceCancelled event raised by the League contract.
type LeagueRaceCancelled struct {
	Id  [32]byte
	Raw types.Log // Blockchain specific contextual infos
}

// FilterRaceCancelled is a free log retrieval operation binding the contract event 0x3f3669b11af9b00fb2a95f8ded36a79a2e1d5374efbc5ec6b2b5b96ceaf9ac6a.
//
// Solidity: event RaceCancelled(bytes32 id)
func (_League *LeagueFilterer) FilterRaceCancelled(opts *bind.FilterOpts) (*LeagueRaceCancelledIterator, error) {

	logs, sub, err := _League.contract.FilterLogs(opts, "RaceCancelled")
	if err != nil {
		return nil, err
	}
	return &LeagueRaceCancelledIterator{contract: _League.contract, event: "RaceCancelled", logs: logs, sub: sub}, nil
}

// WatchRaceCancelled is a free log subscription operation binding the contract event 0x3f3669b11af9b00fb2a95f8ded36a79a2e1d5374efbc5ec6b2b5b96ceaf9ac6a.
//
// Solidity: event RaceCancelled(bytes32 id)
func (_League *LeagueFilterer) WatchRaceCancelled(opts *bind.WatchOpts, sink chan<- *LeagueRaceCancelled) (event.Subscription, error) {

	logs, sub, err := _League.contract.WatchLogs(opts, "RaceCancelled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LeagueRaceCancelled)
				if err := _League.contract.UnpackLog(event, "RaceCancelled", log); err != nil {
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
func (_League *LeagueFilterer) ParseRaceCancelled(log types.Log) (*LeagueRaceCancelled, error) {
	event := new(LeagueRaceCancelled)
	if err := _League.contract.UnpackLog(event, "RaceCancelled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LeagueRaceCreatedIterator is returned from FilterRaceCreated and is used to iterate over the raw logs and unpacked data for RaceCreated events raised by the League contract.
type LeagueRaceCreatedIterator struct {
	Event *LeagueRaceCreated // Event containing the contract specifics and raw log

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
func (it *LeagueRaceCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LeagueRaceCreated)
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
		it.Event = new(LeagueRaceCreated)
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
func (it *LeagueRaceCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LeagueRaceCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LeagueRaceCreated represents a RaceCreated event raised by the League contract.
type LeagueRaceCreated struct {
	Id      [32]byte
	Slots   *big.Int
	StartAt *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRaceCreated is a free log retrieval operation binding the contract event 0x3d4c26760d0e4825f9a3f6ccba234d54aeeae3e6471639b8366708a183f1e837.
//
// Solidity: event RaceCreated(bytes32 id, uint256 slots, uint256 startAt)
func (_League *LeagueFilterer) FilterRaceCreated(opts *bind.FilterOpts) (*LeagueRaceCreatedIterator, error) {

	logs, sub, err := _League.contract.FilterLogs(opts, "RaceCreated")
	if err != nil {
		return nil, err
	}
	return &LeagueRaceCreatedIterator{contract: _League.contract, event: "RaceCreated", logs: logs, sub: sub}, nil
}

// WatchRaceCreated is a free log subscription operation binding the contract event 0x3d4c26760d0e4825f9a3f6ccba234d54aeeae3e6471639b8366708a183f1e837.
//
// Solidity: event RaceCreated(bytes32 id, uint256 slots, uint256 startAt)
func (_League *LeagueFilterer) WatchRaceCreated(opts *bind.WatchOpts, sink chan<- *LeagueRaceCreated) (event.Subscription, error) {

	logs, sub, err := _League.contract.WatchLogs(opts, "RaceCreated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LeagueRaceCreated)
				if err := _League.contract.UnpackLog(event, "RaceCreated", log); err != nil {
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
func (_League *LeagueFilterer) ParseRaceCreated(log types.Log) (*LeagueRaceCreated, error) {
	event := new(LeagueRaceCreated)
	if err := _League.contract.UnpackLog(event, "RaceCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LeagueRaceResultUpdatedIterator is returned from FilterRaceResultUpdated and is used to iterate over the raw logs and unpacked data for RaceResultUpdated events raised by the League contract.
type LeagueRaceResultUpdatedIterator struct {
	Event *LeagueRaceResultUpdated // Event containing the contract specifics and raw log

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
func (it *LeagueRaceResultUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LeagueRaceResultUpdated)
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
		it.Event = new(LeagueRaceResultUpdated)
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
func (it *LeagueRaceResultUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LeagueRaceResultUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LeagueRaceResultUpdated represents a RaceResultUpdated event raised by the League contract.
type LeagueRaceResultUpdated struct {
	Id     [32]byte
	Result [32]byte
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRaceResultUpdated is a free log retrieval operation binding the contract event 0xd82663592968d73ce1995154e44c793fcb46a4006abfb7438656fb0d7ba5ff49.
//
// Solidity: event RaceResultUpdated(bytes32 id, bytes32 result)
func (_League *LeagueFilterer) FilterRaceResultUpdated(opts *bind.FilterOpts) (*LeagueRaceResultUpdatedIterator, error) {

	logs, sub, err := _League.contract.FilterLogs(opts, "RaceResultUpdated")
	if err != nil {
		return nil, err
	}
	return &LeagueRaceResultUpdatedIterator{contract: _League.contract, event: "RaceResultUpdated", logs: logs, sub: sub}, nil
}

// WatchRaceResultUpdated is a free log subscription operation binding the contract event 0xd82663592968d73ce1995154e44c793fcb46a4006abfb7438656fb0d7ba5ff49.
//
// Solidity: event RaceResultUpdated(bytes32 id, bytes32 result)
func (_League *LeagueFilterer) WatchRaceResultUpdated(opts *bind.WatchOpts, sink chan<- *LeagueRaceResultUpdated) (event.Subscription, error) {

	logs, sub, err := _League.contract.WatchLogs(opts, "RaceResultUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LeagueRaceResultUpdated)
				if err := _League.contract.UnpackLog(event, "RaceResultUpdated", log); err != nil {
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
func (_League *LeagueFilterer) ParseRaceResultUpdated(log types.Log) (*LeagueRaceResultUpdated, error) {
	event := new(LeagueRaceResultUpdated)
	if err := _League.contract.UnpackLog(event, "RaceResultUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LeagueRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the League contract.
type LeagueRoleAdminChangedIterator struct {
	Event *LeagueRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *LeagueRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LeagueRoleAdminChanged)
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
		it.Event = new(LeagueRoleAdminChanged)
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
func (it *LeagueRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LeagueRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LeagueRoleAdminChanged represents a RoleAdminChanged event raised by the League contract.
type LeagueRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_League *LeagueFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*LeagueRoleAdminChangedIterator, error) {

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

	logs, sub, err := _League.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &LeagueRoleAdminChangedIterator{contract: _League.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_League *LeagueFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *LeagueRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _League.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LeagueRoleAdminChanged)
				if err := _League.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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
func (_League *LeagueFilterer) ParseRoleAdminChanged(log types.Log) (*LeagueRoleAdminChanged, error) {
	event := new(LeagueRoleAdminChanged)
	if err := _League.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LeagueRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the League contract.
type LeagueRoleGrantedIterator struct {
	Event *LeagueRoleGranted // Event containing the contract specifics and raw log

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
func (it *LeagueRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LeagueRoleGranted)
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
		it.Event = new(LeagueRoleGranted)
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
func (it *LeagueRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LeagueRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LeagueRoleGranted represents a RoleGranted event raised by the League contract.
type LeagueRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_League *LeagueFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*LeagueRoleGrantedIterator, error) {

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

	logs, sub, err := _League.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &LeagueRoleGrantedIterator{contract: _League.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_League *LeagueFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *LeagueRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _League.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LeagueRoleGranted)
				if err := _League.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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
func (_League *LeagueFilterer) ParseRoleGranted(log types.Log) (*LeagueRoleGranted, error) {
	event := new(LeagueRoleGranted)
	if err := _League.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LeagueRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the League contract.
type LeagueRoleRevokedIterator struct {
	Event *LeagueRoleRevoked // Event containing the contract specifics and raw log

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
func (it *LeagueRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LeagueRoleRevoked)
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
		it.Event = new(LeagueRoleRevoked)
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
func (it *LeagueRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LeagueRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LeagueRoleRevoked represents a RoleRevoked event raised by the League contract.
type LeagueRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_League *LeagueFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*LeagueRoleRevokedIterator, error) {

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

	logs, sub, err := _League.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &LeagueRoleRevokedIterator{contract: _League.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_League *LeagueFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *LeagueRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _League.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LeagueRoleRevoked)
				if err := _League.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
func (_League *LeagueFilterer) ParseRoleRevoked(log types.Log) (*LeagueRoleRevoked, error) {
	event := new(LeagueRoleRevoked)
	if err := _League.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LeagueUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the League contract.
type LeagueUnpausedIterator struct {
	Event *LeagueUnpaused // Event containing the contract specifics and raw log

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
func (it *LeagueUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LeagueUnpaused)
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
		it.Event = new(LeagueUnpaused)
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
func (it *LeagueUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LeagueUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LeagueUnpaused represents a Unpaused event raised by the League contract.
type LeagueUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_League *LeagueFilterer) FilterUnpaused(opts *bind.FilterOpts) (*LeagueUnpausedIterator, error) {

	logs, sub, err := _League.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &LeagueUnpausedIterator{contract: _League.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_League *LeagueFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *LeagueUnpaused) (event.Subscription, error) {

	logs, sub, err := _League.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LeagueUnpaused)
				if err := _League.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
func (_League *LeagueFilterer) ParseUnpaused(log types.Log) (*LeagueUnpaused, error) {
	event := new(LeagueUnpaused)
	if err := _League.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
