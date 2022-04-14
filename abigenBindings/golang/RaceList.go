// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package RaceList

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

// IRaceRace is an auto generated low-level Go binding around an user-defined struct.
type IRaceRace struct {
	Commission *big.Int
	RewardRate *big.Int
	Slots      *big.Int
	BetStarted *big.Int
	BetEnded   *big.Int
	Result     [32]byte
}

// RaceListMetaData contains all meta data concerning the RaceList contract.
var RaceListMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"InvalidCommission\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidRewardRate\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidSlot\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RaceNotExisted\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"}],\"name\":\"TooEarly\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"}],\"name\":\"TooLate\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"}],\"name\":\"RaceCancelled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"commission\",\"type\":\"uint256\"}],\"name\":\"RaceCommissionUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"slots\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"betStarted\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"betEnded\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"commission\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"rewardRate\",\"type\":\"uint256\"}],\"name\":\"RaceCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"result\",\"type\":\"bytes32\"}],\"name\":\"RaceResultUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"rewardRate\",\"type\":\"uint256\"}],\"name\":\"RaceRewardRateUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PAUSER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getRoleMember\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleMemberCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"raceId\",\"type\":\"bytes32\"}],\"name\":\"raceIsExisted\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"raceId\",\"type\":\"bytes32\"}],\"name\":\"raceResult\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"slots\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"betStarted\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"betEnded\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"commission\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rewardRate\",\"type\":\"uint256\"}],\"name\":\"createRace\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"raceId\",\"type\":\"bytes32\"}],\"name\":\"getRace\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"commission\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rewardRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"slots\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"betStarted\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"betEnded\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"result\",\"type\":\"bytes32\"}],\"internalType\":\"structIRace.Race\",\"name\":\"race\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"}],\"name\":\"cancelRace\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"result\",\"type\":\"bytes32\"}],\"name\":\"updateResult\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"commission\",\"type\":\"uint256\"}],\"name\":\"updateCommission\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"rewardRate\",\"type\":\"uint256\"}],\"name\":\"updateRewardRate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// RaceListABI is the input ABI used to generate the binding from.
// Deprecated: Use RaceListMetaData.ABI instead.
var RaceListABI = RaceListMetaData.ABI

// RaceList is an auto generated Go binding around an Ethereum contract.
type RaceList struct {
	RaceListCaller     // Read-only binding to the contract
	RaceListTransactor // Write-only binding to the contract
	RaceListFilterer   // Log filterer for contract events
}

// RaceListCaller is an auto generated read-only Go binding around an Ethereum contract.
type RaceListCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RaceListTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RaceListTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RaceListFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RaceListFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RaceListSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RaceListSession struct {
	Contract     *RaceList         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RaceListCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RaceListCallerSession struct {
	Contract *RaceListCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// RaceListTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RaceListTransactorSession struct {
	Contract     *RaceListTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// RaceListRaw is an auto generated low-level Go binding around an Ethereum contract.
type RaceListRaw struct {
	Contract *RaceList // Generic contract binding to access the raw methods on
}

// RaceListCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RaceListCallerRaw struct {
	Contract *RaceListCaller // Generic read-only contract binding to access the raw methods on
}

// RaceListTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RaceListTransactorRaw struct {
	Contract *RaceListTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRaceList creates a new instance of RaceList, bound to a specific deployed contract.
func NewRaceList(address common.Address, backend bind.ContractBackend) (*RaceList, error) {
	contract, err := bindRaceList(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &RaceList{RaceListCaller: RaceListCaller{contract: contract}, RaceListTransactor: RaceListTransactor{contract: contract}, RaceListFilterer: RaceListFilterer{contract: contract}}, nil
}

// NewRaceListCaller creates a new read-only instance of RaceList, bound to a specific deployed contract.
func NewRaceListCaller(address common.Address, caller bind.ContractCaller) (*RaceListCaller, error) {
	contract, err := bindRaceList(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RaceListCaller{contract: contract}, nil
}

// NewRaceListTransactor creates a new write-only instance of RaceList, bound to a specific deployed contract.
func NewRaceListTransactor(address common.Address, transactor bind.ContractTransactor) (*RaceListTransactor, error) {
	contract, err := bindRaceList(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RaceListTransactor{contract: contract}, nil
}

// NewRaceListFilterer creates a new log filterer instance of RaceList, bound to a specific deployed contract.
func NewRaceListFilterer(address common.Address, filterer bind.ContractFilterer) (*RaceListFilterer, error) {
	contract, err := bindRaceList(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RaceListFilterer{contract: contract}, nil
}

// bindRaceList binds a generic wrapper to an already deployed contract.
func bindRaceList(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(RaceListABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RaceList *RaceListRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RaceList.Contract.RaceListCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RaceList *RaceListRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RaceList.Contract.RaceListTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RaceList *RaceListRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RaceList.Contract.RaceListTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RaceList *RaceListCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RaceList.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RaceList *RaceListTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RaceList.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RaceList *RaceListTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RaceList.Contract.contract.Transact(opts, method, params...)
}

// ADMINROLE is a free data retrieval call binding the contract method 0x75b238fc.
//
// Solidity: function ADMIN_ROLE() view returns(bytes32)
func (_RaceList *RaceListCaller) ADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _RaceList.contract.Call(opts, &out, "ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ADMINROLE is a free data retrieval call binding the contract method 0x75b238fc.
//
// Solidity: function ADMIN_ROLE() view returns(bytes32)
func (_RaceList *RaceListSession) ADMINROLE() ([32]byte, error) {
	return _RaceList.Contract.ADMINROLE(&_RaceList.CallOpts)
}

// ADMINROLE is a free data retrieval call binding the contract method 0x75b238fc.
//
// Solidity: function ADMIN_ROLE() view returns(bytes32)
func (_RaceList *RaceListCallerSession) ADMINROLE() ([32]byte, error) {
	return _RaceList.Contract.ADMINROLE(&_RaceList.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_RaceList *RaceListCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _RaceList.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_RaceList *RaceListSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _RaceList.Contract.DEFAULTADMINROLE(&_RaceList.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_RaceList *RaceListCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _RaceList.Contract.DEFAULTADMINROLE(&_RaceList.CallOpts)
}

// PAUSERROLE is a free data retrieval call binding the contract method 0xe63ab1e9.
//
// Solidity: function PAUSER_ROLE() view returns(bytes32)
func (_RaceList *RaceListCaller) PAUSERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _RaceList.contract.Call(opts, &out, "PAUSER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// PAUSERROLE is a free data retrieval call binding the contract method 0xe63ab1e9.
//
// Solidity: function PAUSER_ROLE() view returns(bytes32)
func (_RaceList *RaceListSession) PAUSERROLE() ([32]byte, error) {
	return _RaceList.Contract.PAUSERROLE(&_RaceList.CallOpts)
}

// PAUSERROLE is a free data retrieval call binding the contract method 0xe63ab1e9.
//
// Solidity: function PAUSER_ROLE() view returns(bytes32)
func (_RaceList *RaceListCallerSession) PAUSERROLE() ([32]byte, error) {
	return _RaceList.Contract.PAUSERROLE(&_RaceList.CallOpts)
}

// GetRace is a free data retrieval call binding the contract method 0xf70125db.
//
// Solidity: function getRace(bytes32 raceId) view returns((uint256,uint256,uint256,uint256,uint256,bytes32) race)
func (_RaceList *RaceListCaller) GetRace(opts *bind.CallOpts, raceId [32]byte) (IRaceRace, error) {
	var out []interface{}
	err := _RaceList.contract.Call(opts, &out, "getRace", raceId)

	if err != nil {
		return *new(IRaceRace), err
	}

	out0 := *abi.ConvertType(out[0], new(IRaceRace)).(*IRaceRace)

	return out0, err

}

// GetRace is a free data retrieval call binding the contract method 0xf70125db.
//
// Solidity: function getRace(bytes32 raceId) view returns((uint256,uint256,uint256,uint256,uint256,bytes32) race)
func (_RaceList *RaceListSession) GetRace(raceId [32]byte) (IRaceRace, error) {
	return _RaceList.Contract.GetRace(&_RaceList.CallOpts, raceId)
}

// GetRace is a free data retrieval call binding the contract method 0xf70125db.
//
// Solidity: function getRace(bytes32 raceId) view returns((uint256,uint256,uint256,uint256,uint256,bytes32) race)
func (_RaceList *RaceListCallerSession) GetRace(raceId [32]byte) (IRaceRace, error) {
	return _RaceList.Contract.GetRace(&_RaceList.CallOpts, raceId)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_RaceList *RaceListCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _RaceList.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_RaceList *RaceListSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _RaceList.Contract.GetRoleAdmin(&_RaceList.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_RaceList *RaceListCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _RaceList.Contract.GetRoleAdmin(&_RaceList.CallOpts, role)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_RaceList *RaceListCaller) GetRoleMember(opts *bind.CallOpts, role [32]byte, index *big.Int) (common.Address, error) {
	var out []interface{}
	err := _RaceList.contract.Call(opts, &out, "getRoleMember", role, index)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_RaceList *RaceListSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _RaceList.Contract.GetRoleMember(&_RaceList.CallOpts, role, index)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_RaceList *RaceListCallerSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _RaceList.Contract.GetRoleMember(&_RaceList.CallOpts, role, index)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_RaceList *RaceListCaller) GetRoleMemberCount(opts *bind.CallOpts, role [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _RaceList.contract.Call(opts, &out, "getRoleMemberCount", role)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_RaceList *RaceListSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _RaceList.Contract.GetRoleMemberCount(&_RaceList.CallOpts, role)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_RaceList *RaceListCallerSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _RaceList.Contract.GetRoleMemberCount(&_RaceList.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_RaceList *RaceListCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _RaceList.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_RaceList *RaceListSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _RaceList.Contract.HasRole(&_RaceList.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_RaceList *RaceListCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _RaceList.Contract.HasRole(&_RaceList.CallOpts, role, account)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_RaceList *RaceListCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RaceList.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_RaceList *RaceListSession) Owner() (common.Address, error) {
	return _RaceList.Contract.Owner(&_RaceList.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_RaceList *RaceListCallerSession) Owner() (common.Address, error) {
	return _RaceList.Contract.Owner(&_RaceList.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_RaceList *RaceListCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _RaceList.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_RaceList *RaceListSession) Paused() (bool, error) {
	return _RaceList.Contract.Paused(&_RaceList.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_RaceList *RaceListCallerSession) Paused() (bool, error) {
	return _RaceList.Contract.Paused(&_RaceList.CallOpts)
}

// RaceIsExisted is a free data retrieval call binding the contract method 0x2869ad12.
//
// Solidity: function raceIsExisted(bytes32 raceId) view returns(bool)
func (_RaceList *RaceListCaller) RaceIsExisted(opts *bind.CallOpts, raceId [32]byte) (bool, error) {
	var out []interface{}
	err := _RaceList.contract.Call(opts, &out, "raceIsExisted", raceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// RaceIsExisted is a free data retrieval call binding the contract method 0x2869ad12.
//
// Solidity: function raceIsExisted(bytes32 raceId) view returns(bool)
func (_RaceList *RaceListSession) RaceIsExisted(raceId [32]byte) (bool, error) {
	return _RaceList.Contract.RaceIsExisted(&_RaceList.CallOpts, raceId)
}

// RaceIsExisted is a free data retrieval call binding the contract method 0x2869ad12.
//
// Solidity: function raceIsExisted(bytes32 raceId) view returns(bool)
func (_RaceList *RaceListCallerSession) RaceIsExisted(raceId [32]byte) (bool, error) {
	return _RaceList.Contract.RaceIsExisted(&_RaceList.CallOpts, raceId)
}

// RaceResult is a free data retrieval call binding the contract method 0x6d93b958.
//
// Solidity: function raceResult(bytes32 raceId) view returns(bytes32)
func (_RaceList *RaceListCaller) RaceResult(opts *bind.CallOpts, raceId [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _RaceList.contract.Call(opts, &out, "raceResult", raceId)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// RaceResult is a free data retrieval call binding the contract method 0x6d93b958.
//
// Solidity: function raceResult(bytes32 raceId) view returns(bytes32)
func (_RaceList *RaceListSession) RaceResult(raceId [32]byte) ([32]byte, error) {
	return _RaceList.Contract.RaceResult(&_RaceList.CallOpts, raceId)
}

// RaceResult is a free data retrieval call binding the contract method 0x6d93b958.
//
// Solidity: function raceResult(bytes32 raceId) view returns(bytes32)
func (_RaceList *RaceListCallerSession) RaceResult(raceId [32]byte) ([32]byte, error) {
	return _RaceList.Contract.RaceResult(&_RaceList.CallOpts, raceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_RaceList *RaceListCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _RaceList.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_RaceList *RaceListSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _RaceList.Contract.SupportsInterface(&_RaceList.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_RaceList *RaceListCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _RaceList.Contract.SupportsInterface(&_RaceList.CallOpts, interfaceId)
}

// CancelRace is a paid mutator transaction binding the contract method 0xc05c107d.
//
// Solidity: function cancelRace(bytes32 id) returns()
func (_RaceList *RaceListTransactor) CancelRace(opts *bind.TransactOpts, id [32]byte) (*types.Transaction, error) {
	return _RaceList.contract.Transact(opts, "cancelRace", id)
}

// CancelRace is a paid mutator transaction binding the contract method 0xc05c107d.
//
// Solidity: function cancelRace(bytes32 id) returns()
func (_RaceList *RaceListSession) CancelRace(id [32]byte) (*types.Transaction, error) {
	return _RaceList.Contract.CancelRace(&_RaceList.TransactOpts, id)
}

// CancelRace is a paid mutator transaction binding the contract method 0xc05c107d.
//
// Solidity: function cancelRace(bytes32 id) returns()
func (_RaceList *RaceListTransactorSession) CancelRace(id [32]byte) (*types.Transaction, error) {
	return _RaceList.Contract.CancelRace(&_RaceList.TransactOpts, id)
}

// CreateRace is a paid mutator transaction binding the contract method 0xd65b869a.
//
// Solidity: function createRace(uint256 slots, uint256 betStarted, uint256 betEnded, uint256 commission, uint256 rewardRate) returns()
func (_RaceList *RaceListTransactor) CreateRace(opts *bind.TransactOpts, slots *big.Int, betStarted *big.Int, betEnded *big.Int, commission *big.Int, rewardRate *big.Int) (*types.Transaction, error) {
	return _RaceList.contract.Transact(opts, "createRace", slots, betStarted, betEnded, commission, rewardRate)
}

// CreateRace is a paid mutator transaction binding the contract method 0xd65b869a.
//
// Solidity: function createRace(uint256 slots, uint256 betStarted, uint256 betEnded, uint256 commission, uint256 rewardRate) returns()
func (_RaceList *RaceListSession) CreateRace(slots *big.Int, betStarted *big.Int, betEnded *big.Int, commission *big.Int, rewardRate *big.Int) (*types.Transaction, error) {
	return _RaceList.Contract.CreateRace(&_RaceList.TransactOpts, slots, betStarted, betEnded, commission, rewardRate)
}

// CreateRace is a paid mutator transaction binding the contract method 0xd65b869a.
//
// Solidity: function createRace(uint256 slots, uint256 betStarted, uint256 betEnded, uint256 commission, uint256 rewardRate) returns()
func (_RaceList *RaceListTransactorSession) CreateRace(slots *big.Int, betStarted *big.Int, betEnded *big.Int, commission *big.Int, rewardRate *big.Int) (*types.Transaction, error) {
	return _RaceList.Contract.CreateRace(&_RaceList.TransactOpts, slots, betStarted, betEnded, commission, rewardRate)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_RaceList *RaceListTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _RaceList.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_RaceList *RaceListSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _RaceList.Contract.GrantRole(&_RaceList.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_RaceList *RaceListTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _RaceList.Contract.GrantRole(&_RaceList.TransactOpts, role, account)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_RaceList *RaceListTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RaceList.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_RaceList *RaceListSession) Pause() (*types.Transaction, error) {
	return _RaceList.Contract.Pause(&_RaceList.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_RaceList *RaceListTransactorSession) Pause() (*types.Transaction, error) {
	return _RaceList.Contract.Pause(&_RaceList.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_RaceList *RaceListTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RaceList.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_RaceList *RaceListSession) RenounceOwnership() (*types.Transaction, error) {
	return _RaceList.Contract.RenounceOwnership(&_RaceList.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_RaceList *RaceListTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _RaceList.Contract.RenounceOwnership(&_RaceList.TransactOpts)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_RaceList *RaceListTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _RaceList.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_RaceList *RaceListSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _RaceList.Contract.RenounceRole(&_RaceList.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_RaceList *RaceListTransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _RaceList.Contract.RenounceRole(&_RaceList.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_RaceList *RaceListTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _RaceList.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_RaceList *RaceListSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _RaceList.Contract.RevokeRole(&_RaceList.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_RaceList *RaceListTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _RaceList.Contract.RevokeRole(&_RaceList.TransactOpts, role, account)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_RaceList *RaceListTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _RaceList.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_RaceList *RaceListSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _RaceList.Contract.TransferOwnership(&_RaceList.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_RaceList *RaceListTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _RaceList.Contract.TransferOwnership(&_RaceList.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_RaceList *RaceListTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RaceList.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_RaceList *RaceListSession) Unpause() (*types.Transaction, error) {
	return _RaceList.Contract.Unpause(&_RaceList.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_RaceList *RaceListTransactorSession) Unpause() (*types.Transaction, error) {
	return _RaceList.Contract.Unpause(&_RaceList.TransactOpts)
}

// UpdateCommission is a paid mutator transaction binding the contract method 0xc58f80a7.
//
// Solidity: function updateCommission(bytes32 id, uint256 commission) returns()
func (_RaceList *RaceListTransactor) UpdateCommission(opts *bind.TransactOpts, id [32]byte, commission *big.Int) (*types.Transaction, error) {
	return _RaceList.contract.Transact(opts, "updateCommission", id, commission)
}

// UpdateCommission is a paid mutator transaction binding the contract method 0xc58f80a7.
//
// Solidity: function updateCommission(bytes32 id, uint256 commission) returns()
func (_RaceList *RaceListSession) UpdateCommission(id [32]byte, commission *big.Int) (*types.Transaction, error) {
	return _RaceList.Contract.UpdateCommission(&_RaceList.TransactOpts, id, commission)
}

// UpdateCommission is a paid mutator transaction binding the contract method 0xc58f80a7.
//
// Solidity: function updateCommission(bytes32 id, uint256 commission) returns()
func (_RaceList *RaceListTransactorSession) UpdateCommission(id [32]byte, commission *big.Int) (*types.Transaction, error) {
	return _RaceList.Contract.UpdateCommission(&_RaceList.TransactOpts, id, commission)
}

// UpdateResult is a paid mutator transaction binding the contract method 0xb6d940fc.
//
// Solidity: function updateResult(bytes32 id, bytes32 result) returns()
func (_RaceList *RaceListTransactor) UpdateResult(opts *bind.TransactOpts, id [32]byte, result [32]byte) (*types.Transaction, error) {
	return _RaceList.contract.Transact(opts, "updateResult", id, result)
}

// UpdateResult is a paid mutator transaction binding the contract method 0xb6d940fc.
//
// Solidity: function updateResult(bytes32 id, bytes32 result) returns()
func (_RaceList *RaceListSession) UpdateResult(id [32]byte, result [32]byte) (*types.Transaction, error) {
	return _RaceList.Contract.UpdateResult(&_RaceList.TransactOpts, id, result)
}

// UpdateResult is a paid mutator transaction binding the contract method 0xb6d940fc.
//
// Solidity: function updateResult(bytes32 id, bytes32 result) returns()
func (_RaceList *RaceListTransactorSession) UpdateResult(id [32]byte, result [32]byte) (*types.Transaction, error) {
	return _RaceList.Contract.UpdateResult(&_RaceList.TransactOpts, id, result)
}

// UpdateRewardRate is a paid mutator transaction binding the contract method 0x68b2c40c.
//
// Solidity: function updateRewardRate(bytes32 id, uint256 rewardRate) returns()
func (_RaceList *RaceListTransactor) UpdateRewardRate(opts *bind.TransactOpts, id [32]byte, rewardRate *big.Int) (*types.Transaction, error) {
	return _RaceList.contract.Transact(opts, "updateRewardRate", id, rewardRate)
}

// UpdateRewardRate is a paid mutator transaction binding the contract method 0x68b2c40c.
//
// Solidity: function updateRewardRate(bytes32 id, uint256 rewardRate) returns()
func (_RaceList *RaceListSession) UpdateRewardRate(id [32]byte, rewardRate *big.Int) (*types.Transaction, error) {
	return _RaceList.Contract.UpdateRewardRate(&_RaceList.TransactOpts, id, rewardRate)
}

// UpdateRewardRate is a paid mutator transaction binding the contract method 0x68b2c40c.
//
// Solidity: function updateRewardRate(bytes32 id, uint256 rewardRate) returns()
func (_RaceList *RaceListTransactorSession) UpdateRewardRate(id [32]byte, rewardRate *big.Int) (*types.Transaction, error) {
	return _RaceList.Contract.UpdateRewardRate(&_RaceList.TransactOpts, id, rewardRate)
}

// RaceListOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the RaceList contract.
type RaceListOwnershipTransferredIterator struct {
	Event *RaceListOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *RaceListOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RaceListOwnershipTransferred)
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
		it.Event = new(RaceListOwnershipTransferred)
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
func (it *RaceListOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RaceListOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RaceListOwnershipTransferred represents a OwnershipTransferred event raised by the RaceList contract.
type RaceListOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_RaceList *RaceListFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*RaceListOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _RaceList.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &RaceListOwnershipTransferredIterator{contract: _RaceList.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_RaceList *RaceListFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *RaceListOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _RaceList.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RaceListOwnershipTransferred)
				if err := _RaceList.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_RaceList *RaceListFilterer) ParseOwnershipTransferred(log types.Log) (*RaceListOwnershipTransferred, error) {
	event := new(RaceListOwnershipTransferred)
	if err := _RaceList.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RaceListPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the RaceList contract.
type RaceListPausedIterator struct {
	Event *RaceListPaused // Event containing the contract specifics and raw log

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
func (it *RaceListPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RaceListPaused)
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
		it.Event = new(RaceListPaused)
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
func (it *RaceListPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RaceListPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RaceListPaused represents a Paused event raised by the RaceList contract.
type RaceListPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_RaceList *RaceListFilterer) FilterPaused(opts *bind.FilterOpts) (*RaceListPausedIterator, error) {

	logs, sub, err := _RaceList.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &RaceListPausedIterator{contract: _RaceList.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_RaceList *RaceListFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *RaceListPaused) (event.Subscription, error) {

	logs, sub, err := _RaceList.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RaceListPaused)
				if err := _RaceList.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_RaceList *RaceListFilterer) ParsePaused(log types.Log) (*RaceListPaused, error) {
	event := new(RaceListPaused)
	if err := _RaceList.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RaceListRaceCancelledIterator is returned from FilterRaceCancelled and is used to iterate over the raw logs and unpacked data for RaceCancelled events raised by the RaceList contract.
type RaceListRaceCancelledIterator struct {
	Event *RaceListRaceCancelled // Event containing the contract specifics and raw log

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
func (it *RaceListRaceCancelledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RaceListRaceCancelled)
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
		it.Event = new(RaceListRaceCancelled)
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
func (it *RaceListRaceCancelledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RaceListRaceCancelledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RaceListRaceCancelled represents a RaceCancelled event raised by the RaceList contract.
type RaceListRaceCancelled struct {
	Id  [32]byte
	Raw types.Log // Blockchain specific contextual infos
}

// FilterRaceCancelled is a free log retrieval operation binding the contract event 0x3f3669b11af9b00fb2a95f8ded36a79a2e1d5374efbc5ec6b2b5b96ceaf9ac6a.
//
// Solidity: event RaceCancelled(bytes32 id)
func (_RaceList *RaceListFilterer) FilterRaceCancelled(opts *bind.FilterOpts) (*RaceListRaceCancelledIterator, error) {

	logs, sub, err := _RaceList.contract.FilterLogs(opts, "RaceCancelled")
	if err != nil {
		return nil, err
	}
	return &RaceListRaceCancelledIterator{contract: _RaceList.contract, event: "RaceCancelled", logs: logs, sub: sub}, nil
}

// WatchRaceCancelled is a free log subscription operation binding the contract event 0x3f3669b11af9b00fb2a95f8ded36a79a2e1d5374efbc5ec6b2b5b96ceaf9ac6a.
//
// Solidity: event RaceCancelled(bytes32 id)
func (_RaceList *RaceListFilterer) WatchRaceCancelled(opts *bind.WatchOpts, sink chan<- *RaceListRaceCancelled) (event.Subscription, error) {

	logs, sub, err := _RaceList.contract.WatchLogs(opts, "RaceCancelled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RaceListRaceCancelled)
				if err := _RaceList.contract.UnpackLog(event, "RaceCancelled", log); err != nil {
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
func (_RaceList *RaceListFilterer) ParseRaceCancelled(log types.Log) (*RaceListRaceCancelled, error) {
	event := new(RaceListRaceCancelled)
	if err := _RaceList.contract.UnpackLog(event, "RaceCancelled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RaceListRaceCommissionUpdatedIterator is returned from FilterRaceCommissionUpdated and is used to iterate over the raw logs and unpacked data for RaceCommissionUpdated events raised by the RaceList contract.
type RaceListRaceCommissionUpdatedIterator struct {
	Event *RaceListRaceCommissionUpdated // Event containing the contract specifics and raw log

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
func (it *RaceListRaceCommissionUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RaceListRaceCommissionUpdated)
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
		it.Event = new(RaceListRaceCommissionUpdated)
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
func (it *RaceListRaceCommissionUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RaceListRaceCommissionUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RaceListRaceCommissionUpdated represents a RaceCommissionUpdated event raised by the RaceList contract.
type RaceListRaceCommissionUpdated struct {
	Id         [32]byte
	Commission *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterRaceCommissionUpdated is a free log retrieval operation binding the contract event 0x9ff8c0f5ebec7767df2725581a4945a5a800a1f929ec51da017dc56d31a2b3a0.
//
// Solidity: event RaceCommissionUpdated(bytes32 id, uint256 commission)
func (_RaceList *RaceListFilterer) FilterRaceCommissionUpdated(opts *bind.FilterOpts) (*RaceListRaceCommissionUpdatedIterator, error) {

	logs, sub, err := _RaceList.contract.FilterLogs(opts, "RaceCommissionUpdated")
	if err != nil {
		return nil, err
	}
	return &RaceListRaceCommissionUpdatedIterator{contract: _RaceList.contract, event: "RaceCommissionUpdated", logs: logs, sub: sub}, nil
}

// WatchRaceCommissionUpdated is a free log subscription operation binding the contract event 0x9ff8c0f5ebec7767df2725581a4945a5a800a1f929ec51da017dc56d31a2b3a0.
//
// Solidity: event RaceCommissionUpdated(bytes32 id, uint256 commission)
func (_RaceList *RaceListFilterer) WatchRaceCommissionUpdated(opts *bind.WatchOpts, sink chan<- *RaceListRaceCommissionUpdated) (event.Subscription, error) {

	logs, sub, err := _RaceList.contract.WatchLogs(opts, "RaceCommissionUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RaceListRaceCommissionUpdated)
				if err := _RaceList.contract.UnpackLog(event, "RaceCommissionUpdated", log); err != nil {
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

// ParseRaceCommissionUpdated is a log parse operation binding the contract event 0x9ff8c0f5ebec7767df2725581a4945a5a800a1f929ec51da017dc56d31a2b3a0.
//
// Solidity: event RaceCommissionUpdated(bytes32 id, uint256 commission)
func (_RaceList *RaceListFilterer) ParseRaceCommissionUpdated(log types.Log) (*RaceListRaceCommissionUpdated, error) {
	event := new(RaceListRaceCommissionUpdated)
	if err := _RaceList.contract.UnpackLog(event, "RaceCommissionUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RaceListRaceCreatedIterator is returned from FilterRaceCreated and is used to iterate over the raw logs and unpacked data for RaceCreated events raised by the RaceList contract.
type RaceListRaceCreatedIterator struct {
	Event *RaceListRaceCreated // Event containing the contract specifics and raw log

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
func (it *RaceListRaceCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RaceListRaceCreated)
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
		it.Event = new(RaceListRaceCreated)
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
func (it *RaceListRaceCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RaceListRaceCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RaceListRaceCreated represents a RaceCreated event raised by the RaceList contract.
type RaceListRaceCreated struct {
	Id         [32]byte
	Slots      *big.Int
	BetStarted *big.Int
	BetEnded   *big.Int
	Commission *big.Int
	RewardRate *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterRaceCreated is a free log retrieval operation binding the contract event 0x1da8002ad97980e4380172489894c1e0e2d36896f491ff1d1773613fac4c8a22.
//
// Solidity: event RaceCreated(bytes32 id, uint256 slots, uint256 betStarted, uint256 betEnded, uint256 commission, uint256 rewardRate)
func (_RaceList *RaceListFilterer) FilterRaceCreated(opts *bind.FilterOpts) (*RaceListRaceCreatedIterator, error) {

	logs, sub, err := _RaceList.contract.FilterLogs(opts, "RaceCreated")
	if err != nil {
		return nil, err
	}
	return &RaceListRaceCreatedIterator{contract: _RaceList.contract, event: "RaceCreated", logs: logs, sub: sub}, nil
}

// WatchRaceCreated is a free log subscription operation binding the contract event 0x1da8002ad97980e4380172489894c1e0e2d36896f491ff1d1773613fac4c8a22.
//
// Solidity: event RaceCreated(bytes32 id, uint256 slots, uint256 betStarted, uint256 betEnded, uint256 commission, uint256 rewardRate)
func (_RaceList *RaceListFilterer) WatchRaceCreated(opts *bind.WatchOpts, sink chan<- *RaceListRaceCreated) (event.Subscription, error) {

	logs, sub, err := _RaceList.contract.WatchLogs(opts, "RaceCreated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RaceListRaceCreated)
				if err := _RaceList.contract.UnpackLog(event, "RaceCreated", log); err != nil {
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

// ParseRaceCreated is a log parse operation binding the contract event 0x1da8002ad97980e4380172489894c1e0e2d36896f491ff1d1773613fac4c8a22.
//
// Solidity: event RaceCreated(bytes32 id, uint256 slots, uint256 betStarted, uint256 betEnded, uint256 commission, uint256 rewardRate)
func (_RaceList *RaceListFilterer) ParseRaceCreated(log types.Log) (*RaceListRaceCreated, error) {
	event := new(RaceListRaceCreated)
	if err := _RaceList.contract.UnpackLog(event, "RaceCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RaceListRaceResultUpdatedIterator is returned from FilterRaceResultUpdated and is used to iterate over the raw logs and unpacked data for RaceResultUpdated events raised by the RaceList contract.
type RaceListRaceResultUpdatedIterator struct {
	Event *RaceListRaceResultUpdated // Event containing the contract specifics and raw log

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
func (it *RaceListRaceResultUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RaceListRaceResultUpdated)
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
		it.Event = new(RaceListRaceResultUpdated)
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
func (it *RaceListRaceResultUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RaceListRaceResultUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RaceListRaceResultUpdated represents a RaceResultUpdated event raised by the RaceList contract.
type RaceListRaceResultUpdated struct {
	Id     [32]byte
	Result [32]byte
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRaceResultUpdated is a free log retrieval operation binding the contract event 0xd82663592968d73ce1995154e44c793fcb46a4006abfb7438656fb0d7ba5ff49.
//
// Solidity: event RaceResultUpdated(bytes32 id, bytes32 result)
func (_RaceList *RaceListFilterer) FilterRaceResultUpdated(opts *bind.FilterOpts) (*RaceListRaceResultUpdatedIterator, error) {

	logs, sub, err := _RaceList.contract.FilterLogs(opts, "RaceResultUpdated")
	if err != nil {
		return nil, err
	}
	return &RaceListRaceResultUpdatedIterator{contract: _RaceList.contract, event: "RaceResultUpdated", logs: logs, sub: sub}, nil
}

// WatchRaceResultUpdated is a free log subscription operation binding the contract event 0xd82663592968d73ce1995154e44c793fcb46a4006abfb7438656fb0d7ba5ff49.
//
// Solidity: event RaceResultUpdated(bytes32 id, bytes32 result)
func (_RaceList *RaceListFilterer) WatchRaceResultUpdated(opts *bind.WatchOpts, sink chan<- *RaceListRaceResultUpdated) (event.Subscription, error) {

	logs, sub, err := _RaceList.contract.WatchLogs(opts, "RaceResultUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RaceListRaceResultUpdated)
				if err := _RaceList.contract.UnpackLog(event, "RaceResultUpdated", log); err != nil {
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
func (_RaceList *RaceListFilterer) ParseRaceResultUpdated(log types.Log) (*RaceListRaceResultUpdated, error) {
	event := new(RaceListRaceResultUpdated)
	if err := _RaceList.contract.UnpackLog(event, "RaceResultUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RaceListRaceRewardRateUpdatedIterator is returned from FilterRaceRewardRateUpdated and is used to iterate over the raw logs and unpacked data for RaceRewardRateUpdated events raised by the RaceList contract.
type RaceListRaceRewardRateUpdatedIterator struct {
	Event *RaceListRaceRewardRateUpdated // Event containing the contract specifics and raw log

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
func (it *RaceListRaceRewardRateUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RaceListRaceRewardRateUpdated)
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
		it.Event = new(RaceListRaceRewardRateUpdated)
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
func (it *RaceListRaceRewardRateUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RaceListRaceRewardRateUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RaceListRaceRewardRateUpdated represents a RaceRewardRateUpdated event raised by the RaceList contract.
type RaceListRaceRewardRateUpdated struct {
	Id         [32]byte
	RewardRate *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterRaceRewardRateUpdated is a free log retrieval operation binding the contract event 0x25a0545ede30f4217593809347e09822194f6a9fe74a5b3f2c5b6e843375f422.
//
// Solidity: event RaceRewardRateUpdated(bytes32 id, uint256 rewardRate)
func (_RaceList *RaceListFilterer) FilterRaceRewardRateUpdated(opts *bind.FilterOpts) (*RaceListRaceRewardRateUpdatedIterator, error) {

	logs, sub, err := _RaceList.contract.FilterLogs(opts, "RaceRewardRateUpdated")
	if err != nil {
		return nil, err
	}
	return &RaceListRaceRewardRateUpdatedIterator{contract: _RaceList.contract, event: "RaceRewardRateUpdated", logs: logs, sub: sub}, nil
}

// WatchRaceRewardRateUpdated is a free log subscription operation binding the contract event 0x25a0545ede30f4217593809347e09822194f6a9fe74a5b3f2c5b6e843375f422.
//
// Solidity: event RaceRewardRateUpdated(bytes32 id, uint256 rewardRate)
func (_RaceList *RaceListFilterer) WatchRaceRewardRateUpdated(opts *bind.WatchOpts, sink chan<- *RaceListRaceRewardRateUpdated) (event.Subscription, error) {

	logs, sub, err := _RaceList.contract.WatchLogs(opts, "RaceRewardRateUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RaceListRaceRewardRateUpdated)
				if err := _RaceList.contract.UnpackLog(event, "RaceRewardRateUpdated", log); err != nil {
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

// ParseRaceRewardRateUpdated is a log parse operation binding the contract event 0x25a0545ede30f4217593809347e09822194f6a9fe74a5b3f2c5b6e843375f422.
//
// Solidity: event RaceRewardRateUpdated(bytes32 id, uint256 rewardRate)
func (_RaceList *RaceListFilterer) ParseRaceRewardRateUpdated(log types.Log) (*RaceListRaceRewardRateUpdated, error) {
	event := new(RaceListRaceRewardRateUpdated)
	if err := _RaceList.contract.UnpackLog(event, "RaceRewardRateUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RaceListRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the RaceList contract.
type RaceListRoleAdminChangedIterator struct {
	Event *RaceListRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *RaceListRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RaceListRoleAdminChanged)
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
		it.Event = new(RaceListRoleAdminChanged)
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
func (it *RaceListRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RaceListRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RaceListRoleAdminChanged represents a RoleAdminChanged event raised by the RaceList contract.
type RaceListRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_RaceList *RaceListFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*RaceListRoleAdminChangedIterator, error) {

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

	logs, sub, err := _RaceList.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &RaceListRoleAdminChangedIterator{contract: _RaceList.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_RaceList *RaceListFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *RaceListRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _RaceList.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RaceListRoleAdminChanged)
				if err := _RaceList.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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
func (_RaceList *RaceListFilterer) ParseRoleAdminChanged(log types.Log) (*RaceListRoleAdminChanged, error) {
	event := new(RaceListRoleAdminChanged)
	if err := _RaceList.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RaceListRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the RaceList contract.
type RaceListRoleGrantedIterator struct {
	Event *RaceListRoleGranted // Event containing the contract specifics and raw log

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
func (it *RaceListRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RaceListRoleGranted)
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
		it.Event = new(RaceListRoleGranted)
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
func (it *RaceListRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RaceListRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RaceListRoleGranted represents a RoleGranted event raised by the RaceList contract.
type RaceListRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_RaceList *RaceListFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*RaceListRoleGrantedIterator, error) {

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

	logs, sub, err := _RaceList.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &RaceListRoleGrantedIterator{contract: _RaceList.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_RaceList *RaceListFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *RaceListRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _RaceList.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RaceListRoleGranted)
				if err := _RaceList.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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
func (_RaceList *RaceListFilterer) ParseRoleGranted(log types.Log) (*RaceListRoleGranted, error) {
	event := new(RaceListRoleGranted)
	if err := _RaceList.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RaceListRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the RaceList contract.
type RaceListRoleRevokedIterator struct {
	Event *RaceListRoleRevoked // Event containing the contract specifics and raw log

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
func (it *RaceListRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RaceListRoleRevoked)
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
		it.Event = new(RaceListRoleRevoked)
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
func (it *RaceListRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RaceListRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RaceListRoleRevoked represents a RoleRevoked event raised by the RaceList contract.
type RaceListRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_RaceList *RaceListFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*RaceListRoleRevokedIterator, error) {

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

	logs, sub, err := _RaceList.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &RaceListRoleRevokedIterator{contract: _RaceList.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_RaceList *RaceListFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *RaceListRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _RaceList.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RaceListRoleRevoked)
				if err := _RaceList.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
func (_RaceList *RaceListFilterer) ParseRoleRevoked(log types.Log) (*RaceListRoleRevoked, error) {
	event := new(RaceListRoleRevoked)
	if err := _RaceList.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RaceListUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the RaceList contract.
type RaceListUnpausedIterator struct {
	Event *RaceListUnpaused // Event containing the contract specifics and raw log

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
func (it *RaceListUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RaceListUnpaused)
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
		it.Event = new(RaceListUnpaused)
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
func (it *RaceListUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RaceListUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RaceListUnpaused represents a Unpaused event raised by the RaceList contract.
type RaceListUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_RaceList *RaceListFilterer) FilterUnpaused(opts *bind.FilterOpts) (*RaceListUnpausedIterator, error) {

	logs, sub, err := _RaceList.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &RaceListUnpausedIterator{contract: _RaceList.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_RaceList *RaceListFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *RaceListUnpaused) (event.Subscription, error) {

	logs, sub, err := _RaceList.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RaceListUnpaused)
				if err := _RaceList.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
func (_RaceList *RaceListFilterer) ParseUnpaused(log types.Log) (*RaceListUnpaused, error) {
	event := new(RaceListUnpaused)
	if err := _RaceList.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
