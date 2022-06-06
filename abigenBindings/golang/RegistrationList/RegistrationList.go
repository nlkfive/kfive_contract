// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package RegistrationList

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

// RegistrationListMetaData contains all meta data concerning the RegistrationList contract.
var RegistrationListMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"vrfCoordinator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"nlggt\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"raceReward\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"raceList\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"subscriptionId\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"keyHash\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AlreadyRegistered\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AlreadySelected\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidContract\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidSender\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidSlot\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotNLGGTHolder\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"have\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"want\",\"type\":\"address\"}],\"name\":\"OnlyCoordinatorCanFulfill\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RaceNotExisted\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RewardIsExisted\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RewardIsNotExisted\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RewardNotExistedOrReceived\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"}],\"name\":\"TooEarly\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"}],\"name\":\"TooLate\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"nlggt\",\"type\":\"address\"}],\"name\":\"NlggtUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"requestId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"raceId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"randomness\",\"type\":\"uint256\"}],\"name\":\"ParticipantsSelected\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"race\",\"type\":\"address\"}],\"name\":\"RaceListUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"raceId\",\"type\":\"bytes32\"}],\"name\":\"RandomInProgress\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"slotId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"participant\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"raceId\",\"type\":\"bytes32\"}],\"name\":\"Registered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"raceId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nftRewardId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes1\",\"name\":\"resultIndex\",\"type\":\"bytes1\"}],\"name\":\"RewardAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"raceId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"slotId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nftRewardId\",\"type\":\"uint256\"}],\"name\":\"RewardReceived\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"raceId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nftRewardId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes1\",\"name\":\"resultIndex\",\"type\":\"bytes1\"}],\"name\":\"RewardRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PAUSER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getRoleMember\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleMemberCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"requestId\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"randomWords\",\"type\":\"uint256[]\"}],\"name\":\"rawFulfillRandomWords\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"slotId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"raceId\",\"type\":\"bytes32\"}],\"name\":\"register\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"raceId\",\"type\":\"bytes32\"}],\"name\":\"selectParticipant\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"requestId\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"raceId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"nftRewardId\",\"type\":\"uint256\"},{\"internalType\":\"bytes1\",\"name\":\"resultIndex\",\"type\":\"bytes1\"}],\"name\":\"addRewardByTransfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"raceId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"nftRewardId\",\"type\":\"uint256\"},{\"internalType\":\"bytes1\",\"name\":\"resultIndex\",\"type\":\"bytes1\"},{\"internalType\":\"string\",\"name\":\"tokenURI\",\"type\":\"string\"}],\"name\":\"addRewardByMint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"raceId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"slotId\",\"type\":\"uint256\"}],\"name\":\"receiveReward\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"raceList\",\"type\":\"address\"}],\"name\":\"updateRaceAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"nlggt\",\"type\":\"address\"}],\"name\":\"updateNlggtAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"raceId\",\"type\":\"bytes32\"}],\"name\":\"selectedParticipants\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"raceId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"slotId\",\"type\":\"uint256\"}],\"name\":\"selectedParticipant\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"raceId\",\"type\":\"bytes32\"}],\"name\":\"totalRegister\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC721Received\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"raceId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"nftRewardId\",\"type\":\"uint256\"},{\"internalType\":\"bytes1\",\"name\":\"resultIndex\",\"type\":\"bytes1\"}],\"name\":\"removeReward\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// RegistrationListABI is the input ABI used to generate the binding from.
// Deprecated: Use RegistrationListMetaData.ABI instead.
var RegistrationListABI = RegistrationListMetaData.ABI

// RegistrationList is an auto generated Go binding around an Ethereum contract.
type RegistrationList struct {
	RegistrationListCaller     // Read-only binding to the contract
	RegistrationListTransactor // Write-only binding to the contract
	RegistrationListFilterer   // Log filterer for contract events
}

// RegistrationListCaller is an auto generated read-only Go binding around an Ethereum contract.
type RegistrationListCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RegistrationListTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RegistrationListTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RegistrationListFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RegistrationListFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RegistrationListSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RegistrationListSession struct {
	Contract     *RegistrationList // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RegistrationListCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RegistrationListCallerSession struct {
	Contract *RegistrationListCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// RegistrationListTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RegistrationListTransactorSession struct {
	Contract     *RegistrationListTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// RegistrationListRaw is an auto generated low-level Go binding around an Ethereum contract.
type RegistrationListRaw struct {
	Contract *RegistrationList // Generic contract binding to access the raw methods on
}

// RegistrationListCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RegistrationListCallerRaw struct {
	Contract *RegistrationListCaller // Generic read-only contract binding to access the raw methods on
}

// RegistrationListTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RegistrationListTransactorRaw struct {
	Contract *RegistrationListTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRegistrationList creates a new instance of RegistrationList, bound to a specific deployed contract.
func NewRegistrationList(address common.Address, backend bind.ContractBackend) (*RegistrationList, error) {
	contract, err := bindRegistrationList(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &RegistrationList{RegistrationListCaller: RegistrationListCaller{contract: contract}, RegistrationListTransactor: RegistrationListTransactor{contract: contract}, RegistrationListFilterer: RegistrationListFilterer{contract: contract}}, nil
}

// NewRegistrationListCaller creates a new read-only instance of RegistrationList, bound to a specific deployed contract.
func NewRegistrationListCaller(address common.Address, caller bind.ContractCaller) (*RegistrationListCaller, error) {
	contract, err := bindRegistrationList(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RegistrationListCaller{contract: contract}, nil
}

// NewRegistrationListTransactor creates a new write-only instance of RegistrationList, bound to a specific deployed contract.
func NewRegistrationListTransactor(address common.Address, transactor bind.ContractTransactor) (*RegistrationListTransactor, error) {
	contract, err := bindRegistrationList(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RegistrationListTransactor{contract: contract}, nil
}

// NewRegistrationListFilterer creates a new log filterer instance of RegistrationList, bound to a specific deployed contract.
func NewRegistrationListFilterer(address common.Address, filterer bind.ContractFilterer) (*RegistrationListFilterer, error) {
	contract, err := bindRegistrationList(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RegistrationListFilterer{contract: contract}, nil
}

// bindRegistrationList binds a generic wrapper to an already deployed contract.
func bindRegistrationList(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(RegistrationListABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RegistrationList *RegistrationListRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RegistrationList.Contract.RegistrationListCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RegistrationList *RegistrationListRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RegistrationList.Contract.RegistrationListTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RegistrationList *RegistrationListRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RegistrationList.Contract.RegistrationListTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RegistrationList *RegistrationListCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RegistrationList.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RegistrationList *RegistrationListTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RegistrationList.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RegistrationList *RegistrationListTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RegistrationList.Contract.contract.Transact(opts, method, params...)
}

// ADMINROLE is a free data retrieval call binding the contract method 0x75b238fc.
//
// Solidity: function ADMIN_ROLE() view returns(bytes32)
func (_RegistrationList *RegistrationListCaller) ADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _RegistrationList.contract.Call(opts, &out, "ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ADMINROLE is a free data retrieval call binding the contract method 0x75b238fc.
//
// Solidity: function ADMIN_ROLE() view returns(bytes32)
func (_RegistrationList *RegistrationListSession) ADMINROLE() ([32]byte, error) {
	return _RegistrationList.Contract.ADMINROLE(&_RegistrationList.CallOpts)
}

// ADMINROLE is a free data retrieval call binding the contract method 0x75b238fc.
//
// Solidity: function ADMIN_ROLE() view returns(bytes32)
func (_RegistrationList *RegistrationListCallerSession) ADMINROLE() ([32]byte, error) {
	return _RegistrationList.Contract.ADMINROLE(&_RegistrationList.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_RegistrationList *RegistrationListCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _RegistrationList.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_RegistrationList *RegistrationListSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _RegistrationList.Contract.DEFAULTADMINROLE(&_RegistrationList.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_RegistrationList *RegistrationListCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _RegistrationList.Contract.DEFAULTADMINROLE(&_RegistrationList.CallOpts)
}

// PAUSERROLE is a free data retrieval call binding the contract method 0xe63ab1e9.
//
// Solidity: function PAUSER_ROLE() view returns(bytes32)
func (_RegistrationList *RegistrationListCaller) PAUSERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _RegistrationList.contract.Call(opts, &out, "PAUSER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// PAUSERROLE is a free data retrieval call binding the contract method 0xe63ab1e9.
//
// Solidity: function PAUSER_ROLE() view returns(bytes32)
func (_RegistrationList *RegistrationListSession) PAUSERROLE() ([32]byte, error) {
	return _RegistrationList.Contract.PAUSERROLE(&_RegistrationList.CallOpts)
}

// PAUSERROLE is a free data retrieval call binding the contract method 0xe63ab1e9.
//
// Solidity: function PAUSER_ROLE() view returns(bytes32)
func (_RegistrationList *RegistrationListCallerSession) PAUSERROLE() ([32]byte, error) {
	return _RegistrationList.Contract.PAUSERROLE(&_RegistrationList.CallOpts)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_RegistrationList *RegistrationListCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _RegistrationList.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_RegistrationList *RegistrationListSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _RegistrationList.Contract.GetRoleAdmin(&_RegistrationList.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_RegistrationList *RegistrationListCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _RegistrationList.Contract.GetRoleAdmin(&_RegistrationList.CallOpts, role)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_RegistrationList *RegistrationListCaller) GetRoleMember(opts *bind.CallOpts, role [32]byte, index *big.Int) (common.Address, error) {
	var out []interface{}
	err := _RegistrationList.contract.Call(opts, &out, "getRoleMember", role, index)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_RegistrationList *RegistrationListSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _RegistrationList.Contract.GetRoleMember(&_RegistrationList.CallOpts, role, index)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_RegistrationList *RegistrationListCallerSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _RegistrationList.Contract.GetRoleMember(&_RegistrationList.CallOpts, role, index)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_RegistrationList *RegistrationListCaller) GetRoleMemberCount(opts *bind.CallOpts, role [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _RegistrationList.contract.Call(opts, &out, "getRoleMemberCount", role)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_RegistrationList *RegistrationListSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _RegistrationList.Contract.GetRoleMemberCount(&_RegistrationList.CallOpts, role)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_RegistrationList *RegistrationListCallerSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _RegistrationList.Contract.GetRoleMemberCount(&_RegistrationList.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_RegistrationList *RegistrationListCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _RegistrationList.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_RegistrationList *RegistrationListSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _RegistrationList.Contract.HasRole(&_RegistrationList.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_RegistrationList *RegistrationListCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _RegistrationList.Contract.HasRole(&_RegistrationList.CallOpts, role, account)
}

// OnERC721Received is a free data retrieval call binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) pure returns(bytes4)
func (_RegistrationList *RegistrationListCaller) OnERC721Received(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) ([4]byte, error) {
	var out []interface{}
	err := _RegistrationList.contract.Call(opts, &out, "onERC721Received", arg0, arg1, arg2, arg3)

	if err != nil {
		return *new([4]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)

	return out0, err

}

// OnERC721Received is a free data retrieval call binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) pure returns(bytes4)
func (_RegistrationList *RegistrationListSession) OnERC721Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) ([4]byte, error) {
	return _RegistrationList.Contract.OnERC721Received(&_RegistrationList.CallOpts, arg0, arg1, arg2, arg3)
}

// OnERC721Received is a free data retrieval call binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) pure returns(bytes4)
func (_RegistrationList *RegistrationListCallerSession) OnERC721Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) ([4]byte, error) {
	return _RegistrationList.Contract.OnERC721Received(&_RegistrationList.CallOpts, arg0, arg1, arg2, arg3)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_RegistrationList *RegistrationListCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RegistrationList.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_RegistrationList *RegistrationListSession) Owner() (common.Address, error) {
	return _RegistrationList.Contract.Owner(&_RegistrationList.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_RegistrationList *RegistrationListCallerSession) Owner() (common.Address, error) {
	return _RegistrationList.Contract.Owner(&_RegistrationList.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_RegistrationList *RegistrationListCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _RegistrationList.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_RegistrationList *RegistrationListSession) Paused() (bool, error) {
	return _RegistrationList.Contract.Paused(&_RegistrationList.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_RegistrationList *RegistrationListCallerSession) Paused() (bool, error) {
	return _RegistrationList.Contract.Paused(&_RegistrationList.CallOpts)
}

// SelectedParticipant is a free data retrieval call binding the contract method 0x81a2deb4.
//
// Solidity: function selectedParticipant(bytes32 raceId, uint256 slotId) view returns(address)
func (_RegistrationList *RegistrationListCaller) SelectedParticipant(opts *bind.CallOpts, raceId [32]byte, slotId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _RegistrationList.contract.Call(opts, &out, "selectedParticipant", raceId, slotId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SelectedParticipant is a free data retrieval call binding the contract method 0x81a2deb4.
//
// Solidity: function selectedParticipant(bytes32 raceId, uint256 slotId) view returns(address)
func (_RegistrationList *RegistrationListSession) SelectedParticipant(raceId [32]byte, slotId *big.Int) (common.Address, error) {
	return _RegistrationList.Contract.SelectedParticipant(&_RegistrationList.CallOpts, raceId, slotId)
}

// SelectedParticipant is a free data retrieval call binding the contract method 0x81a2deb4.
//
// Solidity: function selectedParticipant(bytes32 raceId, uint256 slotId) view returns(address)
func (_RegistrationList *RegistrationListCallerSession) SelectedParticipant(raceId [32]byte, slotId *big.Int) (common.Address, error) {
	return _RegistrationList.Contract.SelectedParticipant(&_RegistrationList.CallOpts, raceId, slotId)
}

// SelectedParticipants is a free data retrieval call binding the contract method 0x0506f9ef.
//
// Solidity: function selectedParticipants(bytes32 raceId) view returns(address[])
func (_RegistrationList *RegistrationListCaller) SelectedParticipants(opts *bind.CallOpts, raceId [32]byte) ([]common.Address, error) {
	var out []interface{}
	err := _RegistrationList.contract.Call(opts, &out, "selectedParticipants", raceId)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// SelectedParticipants is a free data retrieval call binding the contract method 0x0506f9ef.
//
// Solidity: function selectedParticipants(bytes32 raceId) view returns(address[])
func (_RegistrationList *RegistrationListSession) SelectedParticipants(raceId [32]byte) ([]common.Address, error) {
	return _RegistrationList.Contract.SelectedParticipants(&_RegistrationList.CallOpts, raceId)
}

// SelectedParticipants is a free data retrieval call binding the contract method 0x0506f9ef.
//
// Solidity: function selectedParticipants(bytes32 raceId) view returns(address[])
func (_RegistrationList *RegistrationListCallerSession) SelectedParticipants(raceId [32]byte) ([]common.Address, error) {
	return _RegistrationList.Contract.SelectedParticipants(&_RegistrationList.CallOpts, raceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_RegistrationList *RegistrationListCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _RegistrationList.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_RegistrationList *RegistrationListSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _RegistrationList.Contract.SupportsInterface(&_RegistrationList.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_RegistrationList *RegistrationListCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _RegistrationList.Contract.SupportsInterface(&_RegistrationList.CallOpts, interfaceId)
}

// TotalRegister is a free data retrieval call binding the contract method 0xbe7cb219.
//
// Solidity: function totalRegister(bytes32 raceId) view returns(uint256[])
func (_RegistrationList *RegistrationListCaller) TotalRegister(opts *bind.CallOpts, raceId [32]byte) ([]*big.Int, error) {
	var out []interface{}
	err := _RegistrationList.contract.Call(opts, &out, "totalRegister", raceId)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// TotalRegister is a free data retrieval call binding the contract method 0xbe7cb219.
//
// Solidity: function totalRegister(bytes32 raceId) view returns(uint256[])
func (_RegistrationList *RegistrationListSession) TotalRegister(raceId [32]byte) ([]*big.Int, error) {
	return _RegistrationList.Contract.TotalRegister(&_RegistrationList.CallOpts, raceId)
}

// TotalRegister is a free data retrieval call binding the contract method 0xbe7cb219.
//
// Solidity: function totalRegister(bytes32 raceId) view returns(uint256[])
func (_RegistrationList *RegistrationListCallerSession) TotalRegister(raceId [32]byte) ([]*big.Int, error) {
	return _RegistrationList.Contract.TotalRegister(&_RegistrationList.CallOpts, raceId)
}

// AddRewardByMint is a paid mutator transaction binding the contract method 0x6bef76da.
//
// Solidity: function addRewardByMint(bytes32 raceId, uint256 nftRewardId, bytes1 resultIndex, string tokenURI) returns()
func (_RegistrationList *RegistrationListTransactor) AddRewardByMint(opts *bind.TransactOpts, raceId [32]byte, nftRewardId *big.Int, resultIndex [1]byte, tokenURI string) (*types.Transaction, error) {
	return _RegistrationList.contract.Transact(opts, "addRewardByMint", raceId, nftRewardId, resultIndex, tokenURI)
}

// AddRewardByMint is a paid mutator transaction binding the contract method 0x6bef76da.
//
// Solidity: function addRewardByMint(bytes32 raceId, uint256 nftRewardId, bytes1 resultIndex, string tokenURI) returns()
func (_RegistrationList *RegistrationListSession) AddRewardByMint(raceId [32]byte, nftRewardId *big.Int, resultIndex [1]byte, tokenURI string) (*types.Transaction, error) {
	return _RegistrationList.Contract.AddRewardByMint(&_RegistrationList.TransactOpts, raceId, nftRewardId, resultIndex, tokenURI)
}

// AddRewardByMint is a paid mutator transaction binding the contract method 0x6bef76da.
//
// Solidity: function addRewardByMint(bytes32 raceId, uint256 nftRewardId, bytes1 resultIndex, string tokenURI) returns()
func (_RegistrationList *RegistrationListTransactorSession) AddRewardByMint(raceId [32]byte, nftRewardId *big.Int, resultIndex [1]byte, tokenURI string) (*types.Transaction, error) {
	return _RegistrationList.Contract.AddRewardByMint(&_RegistrationList.TransactOpts, raceId, nftRewardId, resultIndex, tokenURI)
}

// AddRewardByTransfer is a paid mutator transaction binding the contract method 0x81357309.
//
// Solidity: function addRewardByTransfer(bytes32 raceId, uint256 nftRewardId, bytes1 resultIndex) returns()
func (_RegistrationList *RegistrationListTransactor) AddRewardByTransfer(opts *bind.TransactOpts, raceId [32]byte, nftRewardId *big.Int, resultIndex [1]byte) (*types.Transaction, error) {
	return _RegistrationList.contract.Transact(opts, "addRewardByTransfer", raceId, nftRewardId, resultIndex)
}

// AddRewardByTransfer is a paid mutator transaction binding the contract method 0x81357309.
//
// Solidity: function addRewardByTransfer(bytes32 raceId, uint256 nftRewardId, bytes1 resultIndex) returns()
func (_RegistrationList *RegistrationListSession) AddRewardByTransfer(raceId [32]byte, nftRewardId *big.Int, resultIndex [1]byte) (*types.Transaction, error) {
	return _RegistrationList.Contract.AddRewardByTransfer(&_RegistrationList.TransactOpts, raceId, nftRewardId, resultIndex)
}

// AddRewardByTransfer is a paid mutator transaction binding the contract method 0x81357309.
//
// Solidity: function addRewardByTransfer(bytes32 raceId, uint256 nftRewardId, bytes1 resultIndex) returns()
func (_RegistrationList *RegistrationListTransactorSession) AddRewardByTransfer(raceId [32]byte, nftRewardId *big.Int, resultIndex [1]byte) (*types.Transaction, error) {
	return _RegistrationList.Contract.AddRewardByTransfer(&_RegistrationList.TransactOpts, raceId, nftRewardId, resultIndex)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_RegistrationList *RegistrationListTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _RegistrationList.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_RegistrationList *RegistrationListSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _RegistrationList.Contract.GrantRole(&_RegistrationList.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_RegistrationList *RegistrationListTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _RegistrationList.Contract.GrantRole(&_RegistrationList.TransactOpts, role, account)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_RegistrationList *RegistrationListTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RegistrationList.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_RegistrationList *RegistrationListSession) Pause() (*types.Transaction, error) {
	return _RegistrationList.Contract.Pause(&_RegistrationList.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_RegistrationList *RegistrationListTransactorSession) Pause() (*types.Transaction, error) {
	return _RegistrationList.Contract.Pause(&_RegistrationList.TransactOpts)
}

// RawFulfillRandomWords is a paid mutator transaction binding the contract method 0x1fe543e3.
//
// Solidity: function rawFulfillRandomWords(uint256 requestId, uint256[] randomWords) returns()
func (_RegistrationList *RegistrationListTransactor) RawFulfillRandomWords(opts *bind.TransactOpts, requestId *big.Int, randomWords []*big.Int) (*types.Transaction, error) {
	return _RegistrationList.contract.Transact(opts, "rawFulfillRandomWords", requestId, randomWords)
}

// RawFulfillRandomWords is a paid mutator transaction binding the contract method 0x1fe543e3.
//
// Solidity: function rawFulfillRandomWords(uint256 requestId, uint256[] randomWords) returns()
func (_RegistrationList *RegistrationListSession) RawFulfillRandomWords(requestId *big.Int, randomWords []*big.Int) (*types.Transaction, error) {
	return _RegistrationList.Contract.RawFulfillRandomWords(&_RegistrationList.TransactOpts, requestId, randomWords)
}

// RawFulfillRandomWords is a paid mutator transaction binding the contract method 0x1fe543e3.
//
// Solidity: function rawFulfillRandomWords(uint256 requestId, uint256[] randomWords) returns()
func (_RegistrationList *RegistrationListTransactorSession) RawFulfillRandomWords(requestId *big.Int, randomWords []*big.Int) (*types.Transaction, error) {
	return _RegistrationList.Contract.RawFulfillRandomWords(&_RegistrationList.TransactOpts, requestId, randomWords)
}

// ReceiveReward is a paid mutator transaction binding the contract method 0x1a0fc715.
//
// Solidity: function receiveReward(bytes32 raceId, uint256 slotId) returns()
func (_RegistrationList *RegistrationListTransactor) ReceiveReward(opts *bind.TransactOpts, raceId [32]byte, slotId *big.Int) (*types.Transaction, error) {
	return _RegistrationList.contract.Transact(opts, "receiveReward", raceId, slotId)
}

// ReceiveReward is a paid mutator transaction binding the contract method 0x1a0fc715.
//
// Solidity: function receiveReward(bytes32 raceId, uint256 slotId) returns()
func (_RegistrationList *RegistrationListSession) ReceiveReward(raceId [32]byte, slotId *big.Int) (*types.Transaction, error) {
	return _RegistrationList.Contract.ReceiveReward(&_RegistrationList.TransactOpts, raceId, slotId)
}

// ReceiveReward is a paid mutator transaction binding the contract method 0x1a0fc715.
//
// Solidity: function receiveReward(bytes32 raceId, uint256 slotId) returns()
func (_RegistrationList *RegistrationListTransactorSession) ReceiveReward(raceId [32]byte, slotId *big.Int) (*types.Transaction, error) {
	return _RegistrationList.Contract.ReceiveReward(&_RegistrationList.TransactOpts, raceId, slotId)
}

// Register is a paid mutator transaction binding the contract method 0x2965d809.
//
// Solidity: function register(uint256 slotId, bytes32 raceId) returns()
func (_RegistrationList *RegistrationListTransactor) Register(opts *bind.TransactOpts, slotId *big.Int, raceId [32]byte) (*types.Transaction, error) {
	return _RegistrationList.contract.Transact(opts, "register", slotId, raceId)
}

// Register is a paid mutator transaction binding the contract method 0x2965d809.
//
// Solidity: function register(uint256 slotId, bytes32 raceId) returns()
func (_RegistrationList *RegistrationListSession) Register(slotId *big.Int, raceId [32]byte) (*types.Transaction, error) {
	return _RegistrationList.Contract.Register(&_RegistrationList.TransactOpts, slotId, raceId)
}

// Register is a paid mutator transaction binding the contract method 0x2965d809.
//
// Solidity: function register(uint256 slotId, bytes32 raceId) returns()
func (_RegistrationList *RegistrationListTransactorSession) Register(slotId *big.Int, raceId [32]byte) (*types.Transaction, error) {
	return _RegistrationList.Contract.Register(&_RegistrationList.TransactOpts, slotId, raceId)
}

// RemoveReward is a paid mutator transaction binding the contract method 0x0fa445f5.
//
// Solidity: function removeReward(bytes32 raceId, uint256 nftRewardId, bytes1 resultIndex) returns()
func (_RegistrationList *RegistrationListTransactor) RemoveReward(opts *bind.TransactOpts, raceId [32]byte, nftRewardId *big.Int, resultIndex [1]byte) (*types.Transaction, error) {
	return _RegistrationList.contract.Transact(opts, "removeReward", raceId, nftRewardId, resultIndex)
}

// RemoveReward is a paid mutator transaction binding the contract method 0x0fa445f5.
//
// Solidity: function removeReward(bytes32 raceId, uint256 nftRewardId, bytes1 resultIndex) returns()
func (_RegistrationList *RegistrationListSession) RemoveReward(raceId [32]byte, nftRewardId *big.Int, resultIndex [1]byte) (*types.Transaction, error) {
	return _RegistrationList.Contract.RemoveReward(&_RegistrationList.TransactOpts, raceId, nftRewardId, resultIndex)
}

// RemoveReward is a paid mutator transaction binding the contract method 0x0fa445f5.
//
// Solidity: function removeReward(bytes32 raceId, uint256 nftRewardId, bytes1 resultIndex) returns()
func (_RegistrationList *RegistrationListTransactorSession) RemoveReward(raceId [32]byte, nftRewardId *big.Int, resultIndex [1]byte) (*types.Transaction, error) {
	return _RegistrationList.Contract.RemoveReward(&_RegistrationList.TransactOpts, raceId, nftRewardId, resultIndex)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_RegistrationList *RegistrationListTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RegistrationList.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_RegistrationList *RegistrationListSession) RenounceOwnership() (*types.Transaction, error) {
	return _RegistrationList.Contract.RenounceOwnership(&_RegistrationList.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_RegistrationList *RegistrationListTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _RegistrationList.Contract.RenounceOwnership(&_RegistrationList.TransactOpts)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_RegistrationList *RegistrationListTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _RegistrationList.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_RegistrationList *RegistrationListSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _RegistrationList.Contract.RenounceRole(&_RegistrationList.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_RegistrationList *RegistrationListTransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _RegistrationList.Contract.RenounceRole(&_RegistrationList.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_RegistrationList *RegistrationListTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _RegistrationList.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_RegistrationList *RegistrationListSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _RegistrationList.Contract.RevokeRole(&_RegistrationList.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_RegistrationList *RegistrationListTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _RegistrationList.Contract.RevokeRole(&_RegistrationList.TransactOpts, role, account)
}

// SelectParticipant is a paid mutator transaction binding the contract method 0x7d0f79b7.
//
// Solidity: function selectParticipant(bytes32 raceId) returns(uint256 requestId)
func (_RegistrationList *RegistrationListTransactor) SelectParticipant(opts *bind.TransactOpts, raceId [32]byte) (*types.Transaction, error) {
	return _RegistrationList.contract.Transact(opts, "selectParticipant", raceId)
}

// SelectParticipant is a paid mutator transaction binding the contract method 0x7d0f79b7.
//
// Solidity: function selectParticipant(bytes32 raceId) returns(uint256 requestId)
func (_RegistrationList *RegistrationListSession) SelectParticipant(raceId [32]byte) (*types.Transaction, error) {
	return _RegistrationList.Contract.SelectParticipant(&_RegistrationList.TransactOpts, raceId)
}

// SelectParticipant is a paid mutator transaction binding the contract method 0x7d0f79b7.
//
// Solidity: function selectParticipant(bytes32 raceId) returns(uint256 requestId)
func (_RegistrationList *RegistrationListTransactorSession) SelectParticipant(raceId [32]byte) (*types.Transaction, error) {
	return _RegistrationList.Contract.SelectParticipant(&_RegistrationList.TransactOpts, raceId)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_RegistrationList *RegistrationListTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _RegistrationList.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_RegistrationList *RegistrationListSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _RegistrationList.Contract.TransferOwnership(&_RegistrationList.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_RegistrationList *RegistrationListTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _RegistrationList.Contract.TransferOwnership(&_RegistrationList.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_RegistrationList *RegistrationListTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RegistrationList.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_RegistrationList *RegistrationListSession) Unpause() (*types.Transaction, error) {
	return _RegistrationList.Contract.Unpause(&_RegistrationList.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_RegistrationList *RegistrationListTransactorSession) Unpause() (*types.Transaction, error) {
	return _RegistrationList.Contract.Unpause(&_RegistrationList.TransactOpts)
}

// UpdateNlggtAddress is a paid mutator transaction binding the contract method 0xd6641f06.
//
// Solidity: function updateNlggtAddress(address nlggt) returns()
func (_RegistrationList *RegistrationListTransactor) UpdateNlggtAddress(opts *bind.TransactOpts, nlggt common.Address) (*types.Transaction, error) {
	return _RegistrationList.contract.Transact(opts, "updateNlggtAddress", nlggt)
}

// UpdateNlggtAddress is a paid mutator transaction binding the contract method 0xd6641f06.
//
// Solidity: function updateNlggtAddress(address nlggt) returns()
func (_RegistrationList *RegistrationListSession) UpdateNlggtAddress(nlggt common.Address) (*types.Transaction, error) {
	return _RegistrationList.Contract.UpdateNlggtAddress(&_RegistrationList.TransactOpts, nlggt)
}

// UpdateNlggtAddress is a paid mutator transaction binding the contract method 0xd6641f06.
//
// Solidity: function updateNlggtAddress(address nlggt) returns()
func (_RegistrationList *RegistrationListTransactorSession) UpdateNlggtAddress(nlggt common.Address) (*types.Transaction, error) {
	return _RegistrationList.Contract.UpdateNlggtAddress(&_RegistrationList.TransactOpts, nlggt)
}

// UpdateRaceAddress is a paid mutator transaction binding the contract method 0xd7c49938.
//
// Solidity: function updateRaceAddress(address raceList) returns()
func (_RegistrationList *RegistrationListTransactor) UpdateRaceAddress(opts *bind.TransactOpts, raceList common.Address) (*types.Transaction, error) {
	return _RegistrationList.contract.Transact(opts, "updateRaceAddress", raceList)
}

// UpdateRaceAddress is a paid mutator transaction binding the contract method 0xd7c49938.
//
// Solidity: function updateRaceAddress(address raceList) returns()
func (_RegistrationList *RegistrationListSession) UpdateRaceAddress(raceList common.Address) (*types.Transaction, error) {
	return _RegistrationList.Contract.UpdateRaceAddress(&_RegistrationList.TransactOpts, raceList)
}

// UpdateRaceAddress is a paid mutator transaction binding the contract method 0xd7c49938.
//
// Solidity: function updateRaceAddress(address raceList) returns()
func (_RegistrationList *RegistrationListTransactorSession) UpdateRaceAddress(raceList common.Address) (*types.Transaction, error) {
	return _RegistrationList.Contract.UpdateRaceAddress(&_RegistrationList.TransactOpts, raceList)
}

// RegistrationListNlggtUpdatedIterator is returned from FilterNlggtUpdated and is used to iterate over the raw logs and unpacked data for NlggtUpdated events raised by the RegistrationList contract.
type RegistrationListNlggtUpdatedIterator struct {
	Event *RegistrationListNlggtUpdated // Event containing the contract specifics and raw log

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
func (it *RegistrationListNlggtUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RegistrationListNlggtUpdated)
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
		it.Event = new(RegistrationListNlggtUpdated)
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
func (it *RegistrationListNlggtUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RegistrationListNlggtUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RegistrationListNlggtUpdated represents a NlggtUpdated event raised by the RegistrationList contract.
type RegistrationListNlggtUpdated struct {
	Nlggt common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterNlggtUpdated is a free log retrieval operation binding the contract event 0x5660a1e6d6115d84caa3bc631133061cbc77b46ff3be76d65178ae3243446b98.
//
// Solidity: event NlggtUpdated(address nlggt)
func (_RegistrationList *RegistrationListFilterer) FilterNlggtUpdated(opts *bind.FilterOpts) (*RegistrationListNlggtUpdatedIterator, error) {

	logs, sub, err := _RegistrationList.contract.FilterLogs(opts, "NlggtUpdated")
	if err != nil {
		return nil, err
	}
	return &RegistrationListNlggtUpdatedIterator{contract: _RegistrationList.contract, event: "NlggtUpdated", logs: logs, sub: sub}, nil
}

// WatchNlggtUpdated is a free log subscription operation binding the contract event 0x5660a1e6d6115d84caa3bc631133061cbc77b46ff3be76d65178ae3243446b98.
//
// Solidity: event NlggtUpdated(address nlggt)
func (_RegistrationList *RegistrationListFilterer) WatchNlggtUpdated(opts *bind.WatchOpts, sink chan<- *RegistrationListNlggtUpdated) (event.Subscription, error) {

	logs, sub, err := _RegistrationList.contract.WatchLogs(opts, "NlggtUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RegistrationListNlggtUpdated)
				if err := _RegistrationList.contract.UnpackLog(event, "NlggtUpdated", log); err != nil {
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

// ParseNlggtUpdated is a log parse operation binding the contract event 0x5660a1e6d6115d84caa3bc631133061cbc77b46ff3be76d65178ae3243446b98.
//
// Solidity: event NlggtUpdated(address nlggt)
func (_RegistrationList *RegistrationListFilterer) ParseNlggtUpdated(log types.Log) (*RegistrationListNlggtUpdated, error) {
	event := new(RegistrationListNlggtUpdated)
	if err := _RegistrationList.contract.UnpackLog(event, "NlggtUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RegistrationListOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the RegistrationList contract.
type RegistrationListOwnershipTransferredIterator struct {
	Event *RegistrationListOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *RegistrationListOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RegistrationListOwnershipTransferred)
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
		it.Event = new(RegistrationListOwnershipTransferred)
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
func (it *RegistrationListOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RegistrationListOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RegistrationListOwnershipTransferred represents a OwnershipTransferred event raised by the RegistrationList contract.
type RegistrationListOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_RegistrationList *RegistrationListFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*RegistrationListOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _RegistrationList.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &RegistrationListOwnershipTransferredIterator{contract: _RegistrationList.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_RegistrationList *RegistrationListFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *RegistrationListOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _RegistrationList.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RegistrationListOwnershipTransferred)
				if err := _RegistrationList.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_RegistrationList *RegistrationListFilterer) ParseOwnershipTransferred(log types.Log) (*RegistrationListOwnershipTransferred, error) {
	event := new(RegistrationListOwnershipTransferred)
	if err := _RegistrationList.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RegistrationListParticipantsSelectedIterator is returned from FilterParticipantsSelected and is used to iterate over the raw logs and unpacked data for ParticipantsSelected events raised by the RegistrationList contract.
type RegistrationListParticipantsSelectedIterator struct {
	Event *RegistrationListParticipantsSelected // Event containing the contract specifics and raw log

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
func (it *RegistrationListParticipantsSelectedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RegistrationListParticipantsSelected)
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
		it.Event = new(RegistrationListParticipantsSelected)
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
func (it *RegistrationListParticipantsSelectedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RegistrationListParticipantsSelectedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RegistrationListParticipantsSelected represents a ParticipantsSelected event raised by the RegistrationList contract.
type RegistrationListParticipantsSelected struct {
	RequestId  *big.Int
	RaceId     [32]byte
	Randomness *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterParticipantsSelected is a free log retrieval operation binding the contract event 0x5a1c9864d9d35edf19fb83e7a07ad46d28d80307bd248fb165e8814a1d1cfe2b.
//
// Solidity: event ParticipantsSelected(uint256 requestId, bytes32 raceId, uint256 randomness)
func (_RegistrationList *RegistrationListFilterer) FilterParticipantsSelected(opts *bind.FilterOpts) (*RegistrationListParticipantsSelectedIterator, error) {

	logs, sub, err := _RegistrationList.contract.FilterLogs(opts, "ParticipantsSelected")
	if err != nil {
		return nil, err
	}
	return &RegistrationListParticipantsSelectedIterator{contract: _RegistrationList.contract, event: "ParticipantsSelected", logs: logs, sub: sub}, nil
}

// WatchParticipantsSelected is a free log subscription operation binding the contract event 0x5a1c9864d9d35edf19fb83e7a07ad46d28d80307bd248fb165e8814a1d1cfe2b.
//
// Solidity: event ParticipantsSelected(uint256 requestId, bytes32 raceId, uint256 randomness)
func (_RegistrationList *RegistrationListFilterer) WatchParticipantsSelected(opts *bind.WatchOpts, sink chan<- *RegistrationListParticipantsSelected) (event.Subscription, error) {

	logs, sub, err := _RegistrationList.contract.WatchLogs(opts, "ParticipantsSelected")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RegistrationListParticipantsSelected)
				if err := _RegistrationList.contract.UnpackLog(event, "ParticipantsSelected", log); err != nil {
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

// ParseParticipantsSelected is a log parse operation binding the contract event 0x5a1c9864d9d35edf19fb83e7a07ad46d28d80307bd248fb165e8814a1d1cfe2b.
//
// Solidity: event ParticipantsSelected(uint256 requestId, bytes32 raceId, uint256 randomness)
func (_RegistrationList *RegistrationListFilterer) ParseParticipantsSelected(log types.Log) (*RegistrationListParticipantsSelected, error) {
	event := new(RegistrationListParticipantsSelected)
	if err := _RegistrationList.contract.UnpackLog(event, "ParticipantsSelected", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RegistrationListPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the RegistrationList contract.
type RegistrationListPausedIterator struct {
	Event *RegistrationListPaused // Event containing the contract specifics and raw log

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
func (it *RegistrationListPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RegistrationListPaused)
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
		it.Event = new(RegistrationListPaused)
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
func (it *RegistrationListPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RegistrationListPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RegistrationListPaused represents a Paused event raised by the RegistrationList contract.
type RegistrationListPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_RegistrationList *RegistrationListFilterer) FilterPaused(opts *bind.FilterOpts) (*RegistrationListPausedIterator, error) {

	logs, sub, err := _RegistrationList.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &RegistrationListPausedIterator{contract: _RegistrationList.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_RegistrationList *RegistrationListFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *RegistrationListPaused) (event.Subscription, error) {

	logs, sub, err := _RegistrationList.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RegistrationListPaused)
				if err := _RegistrationList.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_RegistrationList *RegistrationListFilterer) ParsePaused(log types.Log) (*RegistrationListPaused, error) {
	event := new(RegistrationListPaused)
	if err := _RegistrationList.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RegistrationListRaceListUpdatedIterator is returned from FilterRaceListUpdated and is used to iterate over the raw logs and unpacked data for RaceListUpdated events raised by the RegistrationList contract.
type RegistrationListRaceListUpdatedIterator struct {
	Event *RegistrationListRaceListUpdated // Event containing the contract specifics and raw log

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
func (it *RegistrationListRaceListUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RegistrationListRaceListUpdated)
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
		it.Event = new(RegistrationListRaceListUpdated)
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
func (it *RegistrationListRaceListUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RegistrationListRaceListUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RegistrationListRaceListUpdated represents a RaceListUpdated event raised by the RegistrationList contract.
type RegistrationListRaceListUpdated struct {
	Race common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterRaceListUpdated is a free log retrieval operation binding the contract event 0x7a659342c2b03e1c14729e3f1c86f414756d13bd8ff0636713cc22246182b42d.
//
// Solidity: event RaceListUpdated(address race)
func (_RegistrationList *RegistrationListFilterer) FilterRaceListUpdated(opts *bind.FilterOpts) (*RegistrationListRaceListUpdatedIterator, error) {

	logs, sub, err := _RegistrationList.contract.FilterLogs(opts, "RaceListUpdated")
	if err != nil {
		return nil, err
	}
	return &RegistrationListRaceListUpdatedIterator{contract: _RegistrationList.contract, event: "RaceListUpdated", logs: logs, sub: sub}, nil
}

// WatchRaceListUpdated is a free log subscription operation binding the contract event 0x7a659342c2b03e1c14729e3f1c86f414756d13bd8ff0636713cc22246182b42d.
//
// Solidity: event RaceListUpdated(address race)
func (_RegistrationList *RegistrationListFilterer) WatchRaceListUpdated(opts *bind.WatchOpts, sink chan<- *RegistrationListRaceListUpdated) (event.Subscription, error) {

	logs, sub, err := _RegistrationList.contract.WatchLogs(opts, "RaceListUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RegistrationListRaceListUpdated)
				if err := _RegistrationList.contract.UnpackLog(event, "RaceListUpdated", log); err != nil {
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
func (_RegistrationList *RegistrationListFilterer) ParseRaceListUpdated(log types.Log) (*RegistrationListRaceListUpdated, error) {
	event := new(RegistrationListRaceListUpdated)
	if err := _RegistrationList.contract.UnpackLog(event, "RaceListUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RegistrationListRandomInProgressIterator is returned from FilterRandomInProgress and is used to iterate over the raw logs and unpacked data for RandomInProgress events raised by the RegistrationList contract.
type RegistrationListRandomInProgressIterator struct {
	Event *RegistrationListRandomInProgress // Event containing the contract specifics and raw log

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
func (it *RegistrationListRandomInProgressIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RegistrationListRandomInProgress)
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
		it.Event = new(RegistrationListRandomInProgress)
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
func (it *RegistrationListRandomInProgressIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RegistrationListRandomInProgressIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RegistrationListRandomInProgress represents a RandomInProgress event raised by the RegistrationList contract.
type RegistrationListRandomInProgress struct {
	RaceId [32]byte
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRandomInProgress is a free log retrieval operation binding the contract event 0x018dc112e9dad7765a03056b4091d237454ee37ae27aab3c3abd16d428900a13.
//
// Solidity: event RandomInProgress(bytes32 raceId)
func (_RegistrationList *RegistrationListFilterer) FilterRandomInProgress(opts *bind.FilterOpts) (*RegistrationListRandomInProgressIterator, error) {

	logs, sub, err := _RegistrationList.contract.FilterLogs(opts, "RandomInProgress")
	if err != nil {
		return nil, err
	}
	return &RegistrationListRandomInProgressIterator{contract: _RegistrationList.contract, event: "RandomInProgress", logs: logs, sub: sub}, nil
}

// WatchRandomInProgress is a free log subscription operation binding the contract event 0x018dc112e9dad7765a03056b4091d237454ee37ae27aab3c3abd16d428900a13.
//
// Solidity: event RandomInProgress(bytes32 raceId)
func (_RegistrationList *RegistrationListFilterer) WatchRandomInProgress(opts *bind.WatchOpts, sink chan<- *RegistrationListRandomInProgress) (event.Subscription, error) {

	logs, sub, err := _RegistrationList.contract.WatchLogs(opts, "RandomInProgress")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RegistrationListRandomInProgress)
				if err := _RegistrationList.contract.UnpackLog(event, "RandomInProgress", log); err != nil {
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

// ParseRandomInProgress is a log parse operation binding the contract event 0x018dc112e9dad7765a03056b4091d237454ee37ae27aab3c3abd16d428900a13.
//
// Solidity: event RandomInProgress(bytes32 raceId)
func (_RegistrationList *RegistrationListFilterer) ParseRandomInProgress(log types.Log) (*RegistrationListRandomInProgress, error) {
	event := new(RegistrationListRandomInProgress)
	if err := _RegistrationList.contract.UnpackLog(event, "RandomInProgress", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RegistrationListRegisteredIterator is returned from FilterRegistered and is used to iterate over the raw logs and unpacked data for Registered events raised by the RegistrationList contract.
type RegistrationListRegisteredIterator struct {
	Event *RegistrationListRegistered // Event containing the contract specifics and raw log

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
func (it *RegistrationListRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RegistrationListRegistered)
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
		it.Event = new(RegistrationListRegistered)
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
func (it *RegistrationListRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RegistrationListRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RegistrationListRegistered represents a Registered event raised by the RegistrationList contract.
type RegistrationListRegistered struct {
	SlotId      *big.Int
	Participant common.Address
	RaceId      [32]byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterRegistered is a free log retrieval operation binding the contract event 0xd62c7bd2779dc6ebe426bed623399b1fc1d190c063eadefcffffb1efe6f253ef.
//
// Solidity: event Registered(uint256 slotId, address participant, bytes32 raceId)
func (_RegistrationList *RegistrationListFilterer) FilterRegistered(opts *bind.FilterOpts) (*RegistrationListRegisteredIterator, error) {

	logs, sub, err := _RegistrationList.contract.FilterLogs(opts, "Registered")
	if err != nil {
		return nil, err
	}
	return &RegistrationListRegisteredIterator{contract: _RegistrationList.contract, event: "Registered", logs: logs, sub: sub}, nil
}

// WatchRegistered is a free log subscription operation binding the contract event 0xd62c7bd2779dc6ebe426bed623399b1fc1d190c063eadefcffffb1efe6f253ef.
//
// Solidity: event Registered(uint256 slotId, address participant, bytes32 raceId)
func (_RegistrationList *RegistrationListFilterer) WatchRegistered(opts *bind.WatchOpts, sink chan<- *RegistrationListRegistered) (event.Subscription, error) {

	logs, sub, err := _RegistrationList.contract.WatchLogs(opts, "Registered")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RegistrationListRegistered)
				if err := _RegistrationList.contract.UnpackLog(event, "Registered", log); err != nil {
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

// ParseRegistered is a log parse operation binding the contract event 0xd62c7bd2779dc6ebe426bed623399b1fc1d190c063eadefcffffb1efe6f253ef.
//
// Solidity: event Registered(uint256 slotId, address participant, bytes32 raceId)
func (_RegistrationList *RegistrationListFilterer) ParseRegistered(log types.Log) (*RegistrationListRegistered, error) {
	event := new(RegistrationListRegistered)
	if err := _RegistrationList.contract.UnpackLog(event, "Registered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RegistrationListRewardAddedIterator is returned from FilterRewardAdded and is used to iterate over the raw logs and unpacked data for RewardAdded events raised by the RegistrationList contract.
type RegistrationListRewardAddedIterator struct {
	Event *RegistrationListRewardAdded // Event containing the contract specifics and raw log

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
func (it *RegistrationListRewardAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RegistrationListRewardAdded)
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
		it.Event = new(RegistrationListRewardAdded)
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
func (it *RegistrationListRewardAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RegistrationListRewardAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RegistrationListRewardAdded represents a RewardAdded event raised by the RegistrationList contract.
type RegistrationListRewardAdded struct {
	RaceId      [32]byte
	NftRewardId *big.Int
	ResultIndex [1]byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterRewardAdded is a free log retrieval operation binding the contract event 0x2b942e265bd6d4ca5617a327f9a60e1a85f4351aaa0be2a84cda1cf9d32c8c04.
//
// Solidity: event RewardAdded(bytes32 raceId, uint256 nftRewardId, bytes1 resultIndex)
func (_RegistrationList *RegistrationListFilterer) FilterRewardAdded(opts *bind.FilterOpts) (*RegistrationListRewardAddedIterator, error) {

	logs, sub, err := _RegistrationList.contract.FilterLogs(opts, "RewardAdded")
	if err != nil {
		return nil, err
	}
	return &RegistrationListRewardAddedIterator{contract: _RegistrationList.contract, event: "RewardAdded", logs: logs, sub: sub}, nil
}

// WatchRewardAdded is a free log subscription operation binding the contract event 0x2b942e265bd6d4ca5617a327f9a60e1a85f4351aaa0be2a84cda1cf9d32c8c04.
//
// Solidity: event RewardAdded(bytes32 raceId, uint256 nftRewardId, bytes1 resultIndex)
func (_RegistrationList *RegistrationListFilterer) WatchRewardAdded(opts *bind.WatchOpts, sink chan<- *RegistrationListRewardAdded) (event.Subscription, error) {

	logs, sub, err := _RegistrationList.contract.WatchLogs(opts, "RewardAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RegistrationListRewardAdded)
				if err := _RegistrationList.contract.UnpackLog(event, "RewardAdded", log); err != nil {
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

// ParseRewardAdded is a log parse operation binding the contract event 0x2b942e265bd6d4ca5617a327f9a60e1a85f4351aaa0be2a84cda1cf9d32c8c04.
//
// Solidity: event RewardAdded(bytes32 raceId, uint256 nftRewardId, bytes1 resultIndex)
func (_RegistrationList *RegistrationListFilterer) ParseRewardAdded(log types.Log) (*RegistrationListRewardAdded, error) {
	event := new(RegistrationListRewardAdded)
	if err := _RegistrationList.contract.UnpackLog(event, "RewardAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RegistrationListRewardReceivedIterator is returned from FilterRewardReceived and is used to iterate over the raw logs and unpacked data for RewardReceived events raised by the RegistrationList contract.
type RegistrationListRewardReceivedIterator struct {
	Event *RegistrationListRewardReceived // Event containing the contract specifics and raw log

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
func (it *RegistrationListRewardReceivedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RegistrationListRewardReceived)
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
		it.Event = new(RegistrationListRewardReceived)
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
func (it *RegistrationListRewardReceivedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RegistrationListRewardReceivedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RegistrationListRewardReceived represents a RewardReceived event raised by the RegistrationList contract.
type RegistrationListRewardReceived struct {
	RaceId      [32]byte
	SlotId      *big.Int
	NftRewardId *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterRewardReceived is a free log retrieval operation binding the contract event 0xe2ece170b35ae79cbe45ffdaf4931b8bc45056da5d3e9aa4aea67fcdf307929a.
//
// Solidity: event RewardReceived(bytes32 raceId, uint256 slotId, uint256 nftRewardId)
func (_RegistrationList *RegistrationListFilterer) FilterRewardReceived(opts *bind.FilterOpts) (*RegistrationListRewardReceivedIterator, error) {

	logs, sub, err := _RegistrationList.contract.FilterLogs(opts, "RewardReceived")
	if err != nil {
		return nil, err
	}
	return &RegistrationListRewardReceivedIterator{contract: _RegistrationList.contract, event: "RewardReceived", logs: logs, sub: sub}, nil
}

// WatchRewardReceived is a free log subscription operation binding the contract event 0xe2ece170b35ae79cbe45ffdaf4931b8bc45056da5d3e9aa4aea67fcdf307929a.
//
// Solidity: event RewardReceived(bytes32 raceId, uint256 slotId, uint256 nftRewardId)
func (_RegistrationList *RegistrationListFilterer) WatchRewardReceived(opts *bind.WatchOpts, sink chan<- *RegistrationListRewardReceived) (event.Subscription, error) {

	logs, sub, err := _RegistrationList.contract.WatchLogs(opts, "RewardReceived")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RegistrationListRewardReceived)
				if err := _RegistrationList.contract.UnpackLog(event, "RewardReceived", log); err != nil {
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

// ParseRewardReceived is a log parse operation binding the contract event 0xe2ece170b35ae79cbe45ffdaf4931b8bc45056da5d3e9aa4aea67fcdf307929a.
//
// Solidity: event RewardReceived(bytes32 raceId, uint256 slotId, uint256 nftRewardId)
func (_RegistrationList *RegistrationListFilterer) ParseRewardReceived(log types.Log) (*RegistrationListRewardReceived, error) {
	event := new(RegistrationListRewardReceived)
	if err := _RegistrationList.contract.UnpackLog(event, "RewardReceived", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RegistrationListRewardRemovedIterator is returned from FilterRewardRemoved and is used to iterate over the raw logs and unpacked data for RewardRemoved events raised by the RegistrationList contract.
type RegistrationListRewardRemovedIterator struct {
	Event *RegistrationListRewardRemoved // Event containing the contract specifics and raw log

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
func (it *RegistrationListRewardRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RegistrationListRewardRemoved)
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
		it.Event = new(RegistrationListRewardRemoved)
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
func (it *RegistrationListRewardRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RegistrationListRewardRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RegistrationListRewardRemoved represents a RewardRemoved event raised by the RegistrationList contract.
type RegistrationListRewardRemoved struct {
	RaceId      [32]byte
	NftRewardId *big.Int
	ResultIndex [1]byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterRewardRemoved is a free log retrieval operation binding the contract event 0xddda89b96dee97fe43f9803c50584f6dd667bedc8b02a7c75407e6906bf31ead.
//
// Solidity: event RewardRemoved(bytes32 raceId, uint256 nftRewardId, bytes1 resultIndex)
func (_RegistrationList *RegistrationListFilterer) FilterRewardRemoved(opts *bind.FilterOpts) (*RegistrationListRewardRemovedIterator, error) {

	logs, sub, err := _RegistrationList.contract.FilterLogs(opts, "RewardRemoved")
	if err != nil {
		return nil, err
	}
	return &RegistrationListRewardRemovedIterator{contract: _RegistrationList.contract, event: "RewardRemoved", logs: logs, sub: sub}, nil
}

// WatchRewardRemoved is a free log subscription operation binding the contract event 0xddda89b96dee97fe43f9803c50584f6dd667bedc8b02a7c75407e6906bf31ead.
//
// Solidity: event RewardRemoved(bytes32 raceId, uint256 nftRewardId, bytes1 resultIndex)
func (_RegistrationList *RegistrationListFilterer) WatchRewardRemoved(opts *bind.WatchOpts, sink chan<- *RegistrationListRewardRemoved) (event.Subscription, error) {

	logs, sub, err := _RegistrationList.contract.WatchLogs(opts, "RewardRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RegistrationListRewardRemoved)
				if err := _RegistrationList.contract.UnpackLog(event, "RewardRemoved", log); err != nil {
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

// ParseRewardRemoved is a log parse operation binding the contract event 0xddda89b96dee97fe43f9803c50584f6dd667bedc8b02a7c75407e6906bf31ead.
//
// Solidity: event RewardRemoved(bytes32 raceId, uint256 nftRewardId, bytes1 resultIndex)
func (_RegistrationList *RegistrationListFilterer) ParseRewardRemoved(log types.Log) (*RegistrationListRewardRemoved, error) {
	event := new(RegistrationListRewardRemoved)
	if err := _RegistrationList.contract.UnpackLog(event, "RewardRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RegistrationListRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the RegistrationList contract.
type RegistrationListRoleAdminChangedIterator struct {
	Event *RegistrationListRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *RegistrationListRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RegistrationListRoleAdminChanged)
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
		it.Event = new(RegistrationListRoleAdminChanged)
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
func (it *RegistrationListRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RegistrationListRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RegistrationListRoleAdminChanged represents a RoleAdminChanged event raised by the RegistrationList contract.
type RegistrationListRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_RegistrationList *RegistrationListFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*RegistrationListRoleAdminChangedIterator, error) {

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

	logs, sub, err := _RegistrationList.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &RegistrationListRoleAdminChangedIterator{contract: _RegistrationList.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_RegistrationList *RegistrationListFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *RegistrationListRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _RegistrationList.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RegistrationListRoleAdminChanged)
				if err := _RegistrationList.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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
func (_RegistrationList *RegistrationListFilterer) ParseRoleAdminChanged(log types.Log) (*RegistrationListRoleAdminChanged, error) {
	event := new(RegistrationListRoleAdminChanged)
	if err := _RegistrationList.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RegistrationListRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the RegistrationList contract.
type RegistrationListRoleGrantedIterator struct {
	Event *RegistrationListRoleGranted // Event containing the contract specifics and raw log

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
func (it *RegistrationListRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RegistrationListRoleGranted)
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
		it.Event = new(RegistrationListRoleGranted)
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
func (it *RegistrationListRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RegistrationListRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RegistrationListRoleGranted represents a RoleGranted event raised by the RegistrationList contract.
type RegistrationListRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_RegistrationList *RegistrationListFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*RegistrationListRoleGrantedIterator, error) {

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

	logs, sub, err := _RegistrationList.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &RegistrationListRoleGrantedIterator{contract: _RegistrationList.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_RegistrationList *RegistrationListFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *RegistrationListRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _RegistrationList.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RegistrationListRoleGranted)
				if err := _RegistrationList.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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
func (_RegistrationList *RegistrationListFilterer) ParseRoleGranted(log types.Log) (*RegistrationListRoleGranted, error) {
	event := new(RegistrationListRoleGranted)
	if err := _RegistrationList.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RegistrationListRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the RegistrationList contract.
type RegistrationListRoleRevokedIterator struct {
	Event *RegistrationListRoleRevoked // Event containing the contract specifics and raw log

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
func (it *RegistrationListRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RegistrationListRoleRevoked)
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
		it.Event = new(RegistrationListRoleRevoked)
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
func (it *RegistrationListRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RegistrationListRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RegistrationListRoleRevoked represents a RoleRevoked event raised by the RegistrationList contract.
type RegistrationListRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_RegistrationList *RegistrationListFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*RegistrationListRoleRevokedIterator, error) {

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

	logs, sub, err := _RegistrationList.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &RegistrationListRoleRevokedIterator{contract: _RegistrationList.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_RegistrationList *RegistrationListFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *RegistrationListRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _RegistrationList.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RegistrationListRoleRevoked)
				if err := _RegistrationList.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
func (_RegistrationList *RegistrationListFilterer) ParseRoleRevoked(log types.Log) (*RegistrationListRoleRevoked, error) {
	event := new(RegistrationListRoleRevoked)
	if err := _RegistrationList.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RegistrationListUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the RegistrationList contract.
type RegistrationListUnpausedIterator struct {
	Event *RegistrationListUnpaused // Event containing the contract specifics and raw log

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
func (it *RegistrationListUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RegistrationListUnpaused)
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
		it.Event = new(RegistrationListUnpaused)
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
func (it *RegistrationListUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RegistrationListUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RegistrationListUnpaused represents a Unpaused event raised by the RegistrationList contract.
type RegistrationListUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_RegistrationList *RegistrationListFilterer) FilterUnpaused(opts *bind.FilterOpts) (*RegistrationListUnpausedIterator, error) {

	logs, sub, err := _RegistrationList.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &RegistrationListUnpausedIterator{contract: _RegistrationList.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_RegistrationList *RegistrationListFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *RegistrationListUnpaused) (event.Subscription, error) {

	logs, sub, err := _RegistrationList.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RegistrationListUnpaused)
				if err := _RegistrationList.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
func (_RegistrationList *RegistrationListFilterer) ParseUnpaused(log types.Log) (*RegistrationListUnpaused, error) {
	event := new(RegistrationListUnpaused)
	if err := _RegistrationList.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
