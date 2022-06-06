// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package Registration

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

// RegistrationMetaData contains all meta data concerning the Registration contract.
var RegistrationMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"vrfCoordinator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"raceReward\",\"type\":\"address\"},{\"internalType\":\"uint64\",\"name\":\"subscriptionId\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"keyHash\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AlreadyRegistered\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AlreadySelected\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidContract\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidSender\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidSlot\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"have\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"want\",\"type\":\"address\"}],\"name\":\"OnlyCoordinatorCanFulfill\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RaceNotExisted\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RewardIsExisted\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RewardIsNotExistedOrReceived\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"}],\"name\":\"TooEarly\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"}],\"name\":\"TooLate\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"requestId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"raceId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"randomness\",\"type\":\"uint256\"}],\"name\":\"ParticipantsSelected\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"raceId\",\"type\":\"bytes32\"}],\"name\":\"RandomInProgress\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"slotId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"participant\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"raceId\",\"type\":\"bytes32\"}],\"name\":\"Registered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"raceId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nftRewardId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes1\",\"name\":\"resultIndex\",\"type\":\"bytes1\"}],\"name\":\"RewardAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"raceId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"slotId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nftRewardId\",\"type\":\"uint256\"}],\"name\":\"RewardReceived\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"raceId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nftRewardId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes1\",\"name\":\"resultIndex\",\"type\":\"bytes1\"}],\"name\":\"RewardRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PAUSER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getRoleMember\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleMemberCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"requestId\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"randomWords\",\"type\":\"uint256[]\"}],\"name\":\"rawFulfillRandomWords\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"slotId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"raceId\",\"type\":\"bytes32\"}],\"name\":\"register\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"raceId\",\"type\":\"bytes32\"}],\"name\":\"selectParticipant\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"requestId\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"leagueAddr\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"raceId\",\"type\":\"bytes32\"}],\"name\":\"enableRegistration\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"raceId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"nftRewardId\",\"type\":\"uint256\"},{\"internalType\":\"bytes1\",\"name\":\"resultIndex\",\"type\":\"bytes1\"}],\"name\":\"addRewardByTransfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"raceId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"nftRewardId\",\"type\":\"uint256\"},{\"internalType\":\"bytes1\",\"name\":\"resultIndex\",\"type\":\"bytes1\"},{\"internalType\":\"string\",\"name\":\"tokenURI\",\"type\":\"string\"}],\"name\":\"addRewardByMint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"raceId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"slotId\",\"type\":\"uint256\"}],\"name\":\"receiveReward\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"raceId\",\"type\":\"bytes32\"}],\"name\":\"selectedParticipants\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"raceId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"slotId\",\"type\":\"uint256\"}],\"name\":\"selectedParticipant\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"raceId\",\"type\":\"bytes32\"}],\"name\":\"totalRegister\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC721Received\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"raceId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"nftRewardId\",\"type\":\"uint256\"},{\"internalType\":\"bytes1\",\"name\":\"resultIndex\",\"type\":\"bytes1\"}],\"name\":\"removeReward\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// RegistrationABI is the input ABI used to generate the binding from.
// Deprecated: Use RegistrationMetaData.ABI instead.
var RegistrationABI = RegistrationMetaData.ABI

// Registration is an auto generated Go binding around an Ethereum contract.
type Registration struct {
	RegistrationCaller     // Read-only binding to the contract
	RegistrationTransactor // Write-only binding to the contract
	RegistrationFilterer   // Log filterer for contract events
}

// RegistrationCaller is an auto generated read-only Go binding around an Ethereum contract.
type RegistrationCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RegistrationTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RegistrationTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RegistrationFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RegistrationFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RegistrationSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RegistrationSession struct {
	Contract     *Registration     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RegistrationCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RegistrationCallerSession struct {
	Contract *RegistrationCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// RegistrationTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RegistrationTransactorSession struct {
	Contract     *RegistrationTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// RegistrationRaw is an auto generated low-level Go binding around an Ethereum contract.
type RegistrationRaw struct {
	Contract *Registration // Generic contract binding to access the raw methods on
}

// RegistrationCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RegistrationCallerRaw struct {
	Contract *RegistrationCaller // Generic read-only contract binding to access the raw methods on
}

// RegistrationTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RegistrationTransactorRaw struct {
	Contract *RegistrationTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRegistration creates a new instance of Registration, bound to a specific deployed contract.
func NewRegistration(address common.Address, backend bind.ContractBackend) (*Registration, error) {
	contract, err := bindRegistration(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Registration{RegistrationCaller: RegistrationCaller{contract: contract}, RegistrationTransactor: RegistrationTransactor{contract: contract}, RegistrationFilterer: RegistrationFilterer{contract: contract}}, nil
}

// NewRegistrationCaller creates a new read-only instance of Registration, bound to a specific deployed contract.
func NewRegistrationCaller(address common.Address, caller bind.ContractCaller) (*RegistrationCaller, error) {
	contract, err := bindRegistration(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RegistrationCaller{contract: contract}, nil
}

// NewRegistrationTransactor creates a new write-only instance of Registration, bound to a specific deployed contract.
func NewRegistrationTransactor(address common.Address, transactor bind.ContractTransactor) (*RegistrationTransactor, error) {
	contract, err := bindRegistration(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RegistrationTransactor{contract: contract}, nil
}

// NewRegistrationFilterer creates a new log filterer instance of Registration, bound to a specific deployed contract.
func NewRegistrationFilterer(address common.Address, filterer bind.ContractFilterer) (*RegistrationFilterer, error) {
	contract, err := bindRegistration(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RegistrationFilterer{contract: contract}, nil
}

// bindRegistration binds a generic wrapper to an already deployed contract.
func bindRegistration(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(RegistrationABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Registration *RegistrationRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Registration.Contract.RegistrationCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Registration *RegistrationRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Registration.Contract.RegistrationTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Registration *RegistrationRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Registration.Contract.RegistrationTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Registration *RegistrationCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Registration.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Registration *RegistrationTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Registration.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Registration *RegistrationTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Registration.Contract.contract.Transact(opts, method, params...)
}

// ADMINROLE is a free data retrieval call binding the contract method 0x75b238fc.
//
// Solidity: function ADMIN_ROLE() view returns(bytes32)
func (_Registration *RegistrationCaller) ADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Registration.contract.Call(opts, &out, "ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ADMINROLE is a free data retrieval call binding the contract method 0x75b238fc.
//
// Solidity: function ADMIN_ROLE() view returns(bytes32)
func (_Registration *RegistrationSession) ADMINROLE() ([32]byte, error) {
	return _Registration.Contract.ADMINROLE(&_Registration.CallOpts)
}

// ADMINROLE is a free data retrieval call binding the contract method 0x75b238fc.
//
// Solidity: function ADMIN_ROLE() view returns(bytes32)
func (_Registration *RegistrationCallerSession) ADMINROLE() ([32]byte, error) {
	return _Registration.Contract.ADMINROLE(&_Registration.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Registration *RegistrationCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Registration.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Registration *RegistrationSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Registration.Contract.DEFAULTADMINROLE(&_Registration.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Registration *RegistrationCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Registration.Contract.DEFAULTADMINROLE(&_Registration.CallOpts)
}

// PAUSERROLE is a free data retrieval call binding the contract method 0xe63ab1e9.
//
// Solidity: function PAUSER_ROLE() view returns(bytes32)
func (_Registration *RegistrationCaller) PAUSERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Registration.contract.Call(opts, &out, "PAUSER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// PAUSERROLE is a free data retrieval call binding the contract method 0xe63ab1e9.
//
// Solidity: function PAUSER_ROLE() view returns(bytes32)
func (_Registration *RegistrationSession) PAUSERROLE() ([32]byte, error) {
	return _Registration.Contract.PAUSERROLE(&_Registration.CallOpts)
}

// PAUSERROLE is a free data retrieval call binding the contract method 0xe63ab1e9.
//
// Solidity: function PAUSER_ROLE() view returns(bytes32)
func (_Registration *RegistrationCallerSession) PAUSERROLE() ([32]byte, error) {
	return _Registration.Contract.PAUSERROLE(&_Registration.CallOpts)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Registration *RegistrationCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _Registration.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Registration *RegistrationSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Registration.Contract.GetRoleAdmin(&_Registration.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Registration *RegistrationCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Registration.Contract.GetRoleAdmin(&_Registration.CallOpts, role)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_Registration *RegistrationCaller) GetRoleMember(opts *bind.CallOpts, role [32]byte, index *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Registration.contract.Call(opts, &out, "getRoleMember", role, index)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_Registration *RegistrationSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _Registration.Contract.GetRoleMember(&_Registration.CallOpts, role, index)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_Registration *RegistrationCallerSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _Registration.Contract.GetRoleMember(&_Registration.CallOpts, role, index)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_Registration *RegistrationCaller) GetRoleMemberCount(opts *bind.CallOpts, role [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _Registration.contract.Call(opts, &out, "getRoleMemberCount", role)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_Registration *RegistrationSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _Registration.Contract.GetRoleMemberCount(&_Registration.CallOpts, role)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_Registration *RegistrationCallerSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _Registration.Contract.GetRoleMemberCount(&_Registration.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Registration *RegistrationCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _Registration.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Registration *RegistrationSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Registration.Contract.HasRole(&_Registration.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Registration *RegistrationCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Registration.Contract.HasRole(&_Registration.CallOpts, role, account)
}

// OnERC721Received is a free data retrieval call binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) pure returns(bytes4)
func (_Registration *RegistrationCaller) OnERC721Received(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) ([4]byte, error) {
	var out []interface{}
	err := _Registration.contract.Call(opts, &out, "onERC721Received", arg0, arg1, arg2, arg3)

	if err != nil {
		return *new([4]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)

	return out0, err

}

// OnERC721Received is a free data retrieval call binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) pure returns(bytes4)
func (_Registration *RegistrationSession) OnERC721Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) ([4]byte, error) {
	return _Registration.Contract.OnERC721Received(&_Registration.CallOpts, arg0, arg1, arg2, arg3)
}

// OnERC721Received is a free data retrieval call binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) pure returns(bytes4)
func (_Registration *RegistrationCallerSession) OnERC721Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) ([4]byte, error) {
	return _Registration.Contract.OnERC721Received(&_Registration.CallOpts, arg0, arg1, arg2, arg3)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Registration *RegistrationCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Registration.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Registration *RegistrationSession) Owner() (common.Address, error) {
	return _Registration.Contract.Owner(&_Registration.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Registration *RegistrationCallerSession) Owner() (common.Address, error) {
	return _Registration.Contract.Owner(&_Registration.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Registration *RegistrationCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Registration.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Registration *RegistrationSession) Paused() (bool, error) {
	return _Registration.Contract.Paused(&_Registration.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Registration *RegistrationCallerSession) Paused() (bool, error) {
	return _Registration.Contract.Paused(&_Registration.CallOpts)
}

// SelectedParticipant is a free data retrieval call binding the contract method 0x81a2deb4.
//
// Solidity: function selectedParticipant(bytes32 raceId, uint256 slotId) view returns(address)
func (_Registration *RegistrationCaller) SelectedParticipant(opts *bind.CallOpts, raceId [32]byte, slotId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Registration.contract.Call(opts, &out, "selectedParticipant", raceId, slotId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SelectedParticipant is a free data retrieval call binding the contract method 0x81a2deb4.
//
// Solidity: function selectedParticipant(bytes32 raceId, uint256 slotId) view returns(address)
func (_Registration *RegistrationSession) SelectedParticipant(raceId [32]byte, slotId *big.Int) (common.Address, error) {
	return _Registration.Contract.SelectedParticipant(&_Registration.CallOpts, raceId, slotId)
}

// SelectedParticipant is a free data retrieval call binding the contract method 0x81a2deb4.
//
// Solidity: function selectedParticipant(bytes32 raceId, uint256 slotId) view returns(address)
func (_Registration *RegistrationCallerSession) SelectedParticipant(raceId [32]byte, slotId *big.Int) (common.Address, error) {
	return _Registration.Contract.SelectedParticipant(&_Registration.CallOpts, raceId, slotId)
}

// SelectedParticipants is a free data retrieval call binding the contract method 0x0506f9ef.
//
// Solidity: function selectedParticipants(bytes32 raceId) view returns(address[])
func (_Registration *RegistrationCaller) SelectedParticipants(opts *bind.CallOpts, raceId [32]byte) ([]common.Address, error) {
	var out []interface{}
	err := _Registration.contract.Call(opts, &out, "selectedParticipants", raceId)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// SelectedParticipants is a free data retrieval call binding the contract method 0x0506f9ef.
//
// Solidity: function selectedParticipants(bytes32 raceId) view returns(address[])
func (_Registration *RegistrationSession) SelectedParticipants(raceId [32]byte) ([]common.Address, error) {
	return _Registration.Contract.SelectedParticipants(&_Registration.CallOpts, raceId)
}

// SelectedParticipants is a free data retrieval call binding the contract method 0x0506f9ef.
//
// Solidity: function selectedParticipants(bytes32 raceId) view returns(address[])
func (_Registration *RegistrationCallerSession) SelectedParticipants(raceId [32]byte) ([]common.Address, error) {
	return _Registration.Contract.SelectedParticipants(&_Registration.CallOpts, raceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Registration *RegistrationCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Registration.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Registration *RegistrationSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Registration.Contract.SupportsInterface(&_Registration.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Registration *RegistrationCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Registration.Contract.SupportsInterface(&_Registration.CallOpts, interfaceId)
}

// TotalRegister is a free data retrieval call binding the contract method 0xbe7cb219.
//
// Solidity: function totalRegister(bytes32 raceId) view returns(uint256[])
func (_Registration *RegistrationCaller) TotalRegister(opts *bind.CallOpts, raceId [32]byte) ([]*big.Int, error) {
	var out []interface{}
	err := _Registration.contract.Call(opts, &out, "totalRegister", raceId)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// TotalRegister is a free data retrieval call binding the contract method 0xbe7cb219.
//
// Solidity: function totalRegister(bytes32 raceId) view returns(uint256[])
func (_Registration *RegistrationSession) TotalRegister(raceId [32]byte) ([]*big.Int, error) {
	return _Registration.Contract.TotalRegister(&_Registration.CallOpts, raceId)
}

// TotalRegister is a free data retrieval call binding the contract method 0xbe7cb219.
//
// Solidity: function totalRegister(bytes32 raceId) view returns(uint256[])
func (_Registration *RegistrationCallerSession) TotalRegister(raceId [32]byte) ([]*big.Int, error) {
	return _Registration.Contract.TotalRegister(&_Registration.CallOpts, raceId)
}

// AddRewardByMint is a paid mutator transaction binding the contract method 0x6bef76da.
//
// Solidity: function addRewardByMint(bytes32 raceId, uint256 nftRewardId, bytes1 resultIndex, string tokenURI) returns()
func (_Registration *RegistrationTransactor) AddRewardByMint(opts *bind.TransactOpts, raceId [32]byte, nftRewardId *big.Int, resultIndex [1]byte, tokenURI string) (*types.Transaction, error) {
	return _Registration.contract.Transact(opts, "addRewardByMint", raceId, nftRewardId, resultIndex, tokenURI)
}

// AddRewardByMint is a paid mutator transaction binding the contract method 0x6bef76da.
//
// Solidity: function addRewardByMint(bytes32 raceId, uint256 nftRewardId, bytes1 resultIndex, string tokenURI) returns()
func (_Registration *RegistrationSession) AddRewardByMint(raceId [32]byte, nftRewardId *big.Int, resultIndex [1]byte, tokenURI string) (*types.Transaction, error) {
	return _Registration.Contract.AddRewardByMint(&_Registration.TransactOpts, raceId, nftRewardId, resultIndex, tokenURI)
}

// AddRewardByMint is a paid mutator transaction binding the contract method 0x6bef76da.
//
// Solidity: function addRewardByMint(bytes32 raceId, uint256 nftRewardId, bytes1 resultIndex, string tokenURI) returns()
func (_Registration *RegistrationTransactorSession) AddRewardByMint(raceId [32]byte, nftRewardId *big.Int, resultIndex [1]byte, tokenURI string) (*types.Transaction, error) {
	return _Registration.Contract.AddRewardByMint(&_Registration.TransactOpts, raceId, nftRewardId, resultIndex, tokenURI)
}

// AddRewardByTransfer is a paid mutator transaction binding the contract method 0x81357309.
//
// Solidity: function addRewardByTransfer(bytes32 raceId, uint256 nftRewardId, bytes1 resultIndex) returns()
func (_Registration *RegistrationTransactor) AddRewardByTransfer(opts *bind.TransactOpts, raceId [32]byte, nftRewardId *big.Int, resultIndex [1]byte) (*types.Transaction, error) {
	return _Registration.contract.Transact(opts, "addRewardByTransfer", raceId, nftRewardId, resultIndex)
}

// AddRewardByTransfer is a paid mutator transaction binding the contract method 0x81357309.
//
// Solidity: function addRewardByTransfer(bytes32 raceId, uint256 nftRewardId, bytes1 resultIndex) returns()
func (_Registration *RegistrationSession) AddRewardByTransfer(raceId [32]byte, nftRewardId *big.Int, resultIndex [1]byte) (*types.Transaction, error) {
	return _Registration.Contract.AddRewardByTransfer(&_Registration.TransactOpts, raceId, nftRewardId, resultIndex)
}

// AddRewardByTransfer is a paid mutator transaction binding the contract method 0x81357309.
//
// Solidity: function addRewardByTransfer(bytes32 raceId, uint256 nftRewardId, bytes1 resultIndex) returns()
func (_Registration *RegistrationTransactorSession) AddRewardByTransfer(raceId [32]byte, nftRewardId *big.Int, resultIndex [1]byte) (*types.Transaction, error) {
	return _Registration.Contract.AddRewardByTransfer(&_Registration.TransactOpts, raceId, nftRewardId, resultIndex)
}

// EnableRegistration is a paid mutator transaction binding the contract method 0x52f3dcfb.
//
// Solidity: function enableRegistration(address leagueAddr, bytes32 raceId) returns()
func (_Registration *RegistrationTransactor) EnableRegistration(opts *bind.TransactOpts, leagueAddr common.Address, raceId [32]byte) (*types.Transaction, error) {
	return _Registration.contract.Transact(opts, "enableRegistration", leagueAddr, raceId)
}

// EnableRegistration is a paid mutator transaction binding the contract method 0x52f3dcfb.
//
// Solidity: function enableRegistration(address leagueAddr, bytes32 raceId) returns()
func (_Registration *RegistrationSession) EnableRegistration(leagueAddr common.Address, raceId [32]byte) (*types.Transaction, error) {
	return _Registration.Contract.EnableRegistration(&_Registration.TransactOpts, leagueAddr, raceId)
}

// EnableRegistration is a paid mutator transaction binding the contract method 0x52f3dcfb.
//
// Solidity: function enableRegistration(address leagueAddr, bytes32 raceId) returns()
func (_Registration *RegistrationTransactorSession) EnableRegistration(leagueAddr common.Address, raceId [32]byte) (*types.Transaction, error) {
	return _Registration.Contract.EnableRegistration(&_Registration.TransactOpts, leagueAddr, raceId)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Registration *RegistrationTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Registration.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Registration *RegistrationSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Registration.Contract.GrantRole(&_Registration.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Registration *RegistrationTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Registration.Contract.GrantRole(&_Registration.TransactOpts, role, account)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Registration *RegistrationTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Registration.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Registration *RegistrationSession) Pause() (*types.Transaction, error) {
	return _Registration.Contract.Pause(&_Registration.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Registration *RegistrationTransactorSession) Pause() (*types.Transaction, error) {
	return _Registration.Contract.Pause(&_Registration.TransactOpts)
}

// RawFulfillRandomWords is a paid mutator transaction binding the contract method 0x1fe543e3.
//
// Solidity: function rawFulfillRandomWords(uint256 requestId, uint256[] randomWords) returns()
func (_Registration *RegistrationTransactor) RawFulfillRandomWords(opts *bind.TransactOpts, requestId *big.Int, randomWords []*big.Int) (*types.Transaction, error) {
	return _Registration.contract.Transact(opts, "rawFulfillRandomWords", requestId, randomWords)
}

// RawFulfillRandomWords is a paid mutator transaction binding the contract method 0x1fe543e3.
//
// Solidity: function rawFulfillRandomWords(uint256 requestId, uint256[] randomWords) returns()
func (_Registration *RegistrationSession) RawFulfillRandomWords(requestId *big.Int, randomWords []*big.Int) (*types.Transaction, error) {
	return _Registration.Contract.RawFulfillRandomWords(&_Registration.TransactOpts, requestId, randomWords)
}

// RawFulfillRandomWords is a paid mutator transaction binding the contract method 0x1fe543e3.
//
// Solidity: function rawFulfillRandomWords(uint256 requestId, uint256[] randomWords) returns()
func (_Registration *RegistrationTransactorSession) RawFulfillRandomWords(requestId *big.Int, randomWords []*big.Int) (*types.Transaction, error) {
	return _Registration.Contract.RawFulfillRandomWords(&_Registration.TransactOpts, requestId, randomWords)
}

// ReceiveReward is a paid mutator transaction binding the contract method 0x1a0fc715.
//
// Solidity: function receiveReward(bytes32 raceId, uint256 slotId) returns()
func (_Registration *RegistrationTransactor) ReceiveReward(opts *bind.TransactOpts, raceId [32]byte, slotId *big.Int) (*types.Transaction, error) {
	return _Registration.contract.Transact(opts, "receiveReward", raceId, slotId)
}

// ReceiveReward is a paid mutator transaction binding the contract method 0x1a0fc715.
//
// Solidity: function receiveReward(bytes32 raceId, uint256 slotId) returns()
func (_Registration *RegistrationSession) ReceiveReward(raceId [32]byte, slotId *big.Int) (*types.Transaction, error) {
	return _Registration.Contract.ReceiveReward(&_Registration.TransactOpts, raceId, slotId)
}

// ReceiveReward is a paid mutator transaction binding the contract method 0x1a0fc715.
//
// Solidity: function receiveReward(bytes32 raceId, uint256 slotId) returns()
func (_Registration *RegistrationTransactorSession) ReceiveReward(raceId [32]byte, slotId *big.Int) (*types.Transaction, error) {
	return _Registration.Contract.ReceiveReward(&_Registration.TransactOpts, raceId, slotId)
}

// Register is a paid mutator transaction binding the contract method 0x2965d809.
//
// Solidity: function register(uint256 slotId, bytes32 raceId) returns()
func (_Registration *RegistrationTransactor) Register(opts *bind.TransactOpts, slotId *big.Int, raceId [32]byte) (*types.Transaction, error) {
	return _Registration.contract.Transact(opts, "register", slotId, raceId)
}

// Register is a paid mutator transaction binding the contract method 0x2965d809.
//
// Solidity: function register(uint256 slotId, bytes32 raceId) returns()
func (_Registration *RegistrationSession) Register(slotId *big.Int, raceId [32]byte) (*types.Transaction, error) {
	return _Registration.Contract.Register(&_Registration.TransactOpts, slotId, raceId)
}

// Register is a paid mutator transaction binding the contract method 0x2965d809.
//
// Solidity: function register(uint256 slotId, bytes32 raceId) returns()
func (_Registration *RegistrationTransactorSession) Register(slotId *big.Int, raceId [32]byte) (*types.Transaction, error) {
	return _Registration.Contract.Register(&_Registration.TransactOpts, slotId, raceId)
}

// RemoveReward is a paid mutator transaction binding the contract method 0x0fa445f5.
//
// Solidity: function removeReward(bytes32 raceId, uint256 nftRewardId, bytes1 resultIndex) returns()
func (_Registration *RegistrationTransactor) RemoveReward(opts *bind.TransactOpts, raceId [32]byte, nftRewardId *big.Int, resultIndex [1]byte) (*types.Transaction, error) {
	return _Registration.contract.Transact(opts, "removeReward", raceId, nftRewardId, resultIndex)
}

// RemoveReward is a paid mutator transaction binding the contract method 0x0fa445f5.
//
// Solidity: function removeReward(bytes32 raceId, uint256 nftRewardId, bytes1 resultIndex) returns()
func (_Registration *RegistrationSession) RemoveReward(raceId [32]byte, nftRewardId *big.Int, resultIndex [1]byte) (*types.Transaction, error) {
	return _Registration.Contract.RemoveReward(&_Registration.TransactOpts, raceId, nftRewardId, resultIndex)
}

// RemoveReward is a paid mutator transaction binding the contract method 0x0fa445f5.
//
// Solidity: function removeReward(bytes32 raceId, uint256 nftRewardId, bytes1 resultIndex) returns()
func (_Registration *RegistrationTransactorSession) RemoveReward(raceId [32]byte, nftRewardId *big.Int, resultIndex [1]byte) (*types.Transaction, error) {
	return _Registration.Contract.RemoveReward(&_Registration.TransactOpts, raceId, nftRewardId, resultIndex)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Registration *RegistrationTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Registration.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Registration *RegistrationSession) RenounceOwnership() (*types.Transaction, error) {
	return _Registration.Contract.RenounceOwnership(&_Registration.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Registration *RegistrationTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Registration.Contract.RenounceOwnership(&_Registration.TransactOpts)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Registration *RegistrationTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Registration.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Registration *RegistrationSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Registration.Contract.RenounceRole(&_Registration.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Registration *RegistrationTransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Registration.Contract.RenounceRole(&_Registration.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Registration *RegistrationTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Registration.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Registration *RegistrationSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Registration.Contract.RevokeRole(&_Registration.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Registration *RegistrationTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Registration.Contract.RevokeRole(&_Registration.TransactOpts, role, account)
}

// SelectParticipant is a paid mutator transaction binding the contract method 0x7d0f79b7.
//
// Solidity: function selectParticipant(bytes32 raceId) returns(uint256 requestId)
func (_Registration *RegistrationTransactor) SelectParticipant(opts *bind.TransactOpts, raceId [32]byte) (*types.Transaction, error) {
	return _Registration.contract.Transact(opts, "selectParticipant", raceId)
}

// SelectParticipant is a paid mutator transaction binding the contract method 0x7d0f79b7.
//
// Solidity: function selectParticipant(bytes32 raceId) returns(uint256 requestId)
func (_Registration *RegistrationSession) SelectParticipant(raceId [32]byte) (*types.Transaction, error) {
	return _Registration.Contract.SelectParticipant(&_Registration.TransactOpts, raceId)
}

// SelectParticipant is a paid mutator transaction binding the contract method 0x7d0f79b7.
//
// Solidity: function selectParticipant(bytes32 raceId) returns(uint256 requestId)
func (_Registration *RegistrationTransactorSession) SelectParticipant(raceId [32]byte) (*types.Transaction, error) {
	return _Registration.Contract.SelectParticipant(&_Registration.TransactOpts, raceId)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Registration *RegistrationTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Registration.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Registration *RegistrationSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Registration.Contract.TransferOwnership(&_Registration.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Registration *RegistrationTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Registration.Contract.TransferOwnership(&_Registration.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Registration *RegistrationTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Registration.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Registration *RegistrationSession) Unpause() (*types.Transaction, error) {
	return _Registration.Contract.Unpause(&_Registration.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Registration *RegistrationTransactorSession) Unpause() (*types.Transaction, error) {
	return _Registration.Contract.Unpause(&_Registration.TransactOpts)
}

// RegistrationOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Registration contract.
type RegistrationOwnershipTransferredIterator struct {
	Event *RegistrationOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *RegistrationOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RegistrationOwnershipTransferred)
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
		it.Event = new(RegistrationOwnershipTransferred)
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
func (it *RegistrationOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RegistrationOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RegistrationOwnershipTransferred represents a OwnershipTransferred event raised by the Registration contract.
type RegistrationOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Registration *RegistrationFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*RegistrationOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Registration.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &RegistrationOwnershipTransferredIterator{contract: _Registration.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Registration *RegistrationFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *RegistrationOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Registration.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RegistrationOwnershipTransferred)
				if err := _Registration.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Registration *RegistrationFilterer) ParseOwnershipTransferred(log types.Log) (*RegistrationOwnershipTransferred, error) {
	event := new(RegistrationOwnershipTransferred)
	if err := _Registration.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RegistrationParticipantsSelectedIterator is returned from FilterParticipantsSelected and is used to iterate over the raw logs and unpacked data for ParticipantsSelected events raised by the Registration contract.
type RegistrationParticipantsSelectedIterator struct {
	Event *RegistrationParticipantsSelected // Event containing the contract specifics and raw log

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
func (it *RegistrationParticipantsSelectedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RegistrationParticipantsSelected)
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
		it.Event = new(RegistrationParticipantsSelected)
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
func (it *RegistrationParticipantsSelectedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RegistrationParticipantsSelectedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RegistrationParticipantsSelected represents a ParticipantsSelected event raised by the Registration contract.
type RegistrationParticipantsSelected struct {
	RequestId  *big.Int
	RaceId     [32]byte
	Randomness *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterParticipantsSelected is a free log retrieval operation binding the contract event 0x5a1c9864d9d35edf19fb83e7a07ad46d28d80307bd248fb165e8814a1d1cfe2b.
//
// Solidity: event ParticipantsSelected(uint256 requestId, bytes32 raceId, uint256 randomness)
func (_Registration *RegistrationFilterer) FilterParticipantsSelected(opts *bind.FilterOpts) (*RegistrationParticipantsSelectedIterator, error) {

	logs, sub, err := _Registration.contract.FilterLogs(opts, "ParticipantsSelected")
	if err != nil {
		return nil, err
	}
	return &RegistrationParticipantsSelectedIterator{contract: _Registration.contract, event: "ParticipantsSelected", logs: logs, sub: sub}, nil
}

// WatchParticipantsSelected is a free log subscription operation binding the contract event 0x5a1c9864d9d35edf19fb83e7a07ad46d28d80307bd248fb165e8814a1d1cfe2b.
//
// Solidity: event ParticipantsSelected(uint256 requestId, bytes32 raceId, uint256 randomness)
func (_Registration *RegistrationFilterer) WatchParticipantsSelected(opts *bind.WatchOpts, sink chan<- *RegistrationParticipantsSelected) (event.Subscription, error) {

	logs, sub, err := _Registration.contract.WatchLogs(opts, "ParticipantsSelected")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RegistrationParticipantsSelected)
				if err := _Registration.contract.UnpackLog(event, "ParticipantsSelected", log); err != nil {
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
func (_Registration *RegistrationFilterer) ParseParticipantsSelected(log types.Log) (*RegistrationParticipantsSelected, error) {
	event := new(RegistrationParticipantsSelected)
	if err := _Registration.contract.UnpackLog(event, "ParticipantsSelected", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RegistrationPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the Registration contract.
type RegistrationPausedIterator struct {
	Event *RegistrationPaused // Event containing the contract specifics and raw log

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
func (it *RegistrationPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RegistrationPaused)
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
		it.Event = new(RegistrationPaused)
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
func (it *RegistrationPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RegistrationPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RegistrationPaused represents a Paused event raised by the Registration contract.
type RegistrationPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Registration *RegistrationFilterer) FilterPaused(opts *bind.FilterOpts) (*RegistrationPausedIterator, error) {

	logs, sub, err := _Registration.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &RegistrationPausedIterator{contract: _Registration.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Registration *RegistrationFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *RegistrationPaused) (event.Subscription, error) {

	logs, sub, err := _Registration.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RegistrationPaused)
				if err := _Registration.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_Registration *RegistrationFilterer) ParsePaused(log types.Log) (*RegistrationPaused, error) {
	event := new(RegistrationPaused)
	if err := _Registration.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RegistrationRandomInProgressIterator is returned from FilterRandomInProgress and is used to iterate over the raw logs and unpacked data for RandomInProgress events raised by the Registration contract.
type RegistrationRandomInProgressIterator struct {
	Event *RegistrationRandomInProgress // Event containing the contract specifics and raw log

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
func (it *RegistrationRandomInProgressIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RegistrationRandomInProgress)
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
		it.Event = new(RegistrationRandomInProgress)
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
func (it *RegistrationRandomInProgressIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RegistrationRandomInProgressIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RegistrationRandomInProgress represents a RandomInProgress event raised by the Registration contract.
type RegistrationRandomInProgress struct {
	RaceId [32]byte
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRandomInProgress is a free log retrieval operation binding the contract event 0x018dc112e9dad7765a03056b4091d237454ee37ae27aab3c3abd16d428900a13.
//
// Solidity: event RandomInProgress(bytes32 raceId)
func (_Registration *RegistrationFilterer) FilterRandomInProgress(opts *bind.FilterOpts) (*RegistrationRandomInProgressIterator, error) {

	logs, sub, err := _Registration.contract.FilterLogs(opts, "RandomInProgress")
	if err != nil {
		return nil, err
	}
	return &RegistrationRandomInProgressIterator{contract: _Registration.contract, event: "RandomInProgress", logs: logs, sub: sub}, nil
}

// WatchRandomInProgress is a free log subscription operation binding the contract event 0x018dc112e9dad7765a03056b4091d237454ee37ae27aab3c3abd16d428900a13.
//
// Solidity: event RandomInProgress(bytes32 raceId)
func (_Registration *RegistrationFilterer) WatchRandomInProgress(opts *bind.WatchOpts, sink chan<- *RegistrationRandomInProgress) (event.Subscription, error) {

	logs, sub, err := _Registration.contract.WatchLogs(opts, "RandomInProgress")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RegistrationRandomInProgress)
				if err := _Registration.contract.UnpackLog(event, "RandomInProgress", log); err != nil {
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
func (_Registration *RegistrationFilterer) ParseRandomInProgress(log types.Log) (*RegistrationRandomInProgress, error) {
	event := new(RegistrationRandomInProgress)
	if err := _Registration.contract.UnpackLog(event, "RandomInProgress", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RegistrationRegisteredIterator is returned from FilterRegistered and is used to iterate over the raw logs and unpacked data for Registered events raised by the Registration contract.
type RegistrationRegisteredIterator struct {
	Event *RegistrationRegistered // Event containing the contract specifics and raw log

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
func (it *RegistrationRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RegistrationRegistered)
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
		it.Event = new(RegistrationRegistered)
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
func (it *RegistrationRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RegistrationRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RegistrationRegistered represents a Registered event raised by the Registration contract.
type RegistrationRegistered struct {
	SlotId      *big.Int
	Participant common.Address
	RaceId      [32]byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterRegistered is a free log retrieval operation binding the contract event 0xd62c7bd2779dc6ebe426bed623399b1fc1d190c063eadefcffffb1efe6f253ef.
//
// Solidity: event Registered(uint256 slotId, address participant, bytes32 raceId)
func (_Registration *RegistrationFilterer) FilterRegistered(opts *bind.FilterOpts) (*RegistrationRegisteredIterator, error) {

	logs, sub, err := _Registration.contract.FilterLogs(opts, "Registered")
	if err != nil {
		return nil, err
	}
	return &RegistrationRegisteredIterator{contract: _Registration.contract, event: "Registered", logs: logs, sub: sub}, nil
}

// WatchRegistered is a free log subscription operation binding the contract event 0xd62c7bd2779dc6ebe426bed623399b1fc1d190c063eadefcffffb1efe6f253ef.
//
// Solidity: event Registered(uint256 slotId, address participant, bytes32 raceId)
func (_Registration *RegistrationFilterer) WatchRegistered(opts *bind.WatchOpts, sink chan<- *RegistrationRegistered) (event.Subscription, error) {

	logs, sub, err := _Registration.contract.WatchLogs(opts, "Registered")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RegistrationRegistered)
				if err := _Registration.contract.UnpackLog(event, "Registered", log); err != nil {
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
func (_Registration *RegistrationFilterer) ParseRegistered(log types.Log) (*RegistrationRegistered, error) {
	event := new(RegistrationRegistered)
	if err := _Registration.contract.UnpackLog(event, "Registered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RegistrationRewardAddedIterator is returned from FilterRewardAdded and is used to iterate over the raw logs and unpacked data for RewardAdded events raised by the Registration contract.
type RegistrationRewardAddedIterator struct {
	Event *RegistrationRewardAdded // Event containing the contract specifics and raw log

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
func (it *RegistrationRewardAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RegistrationRewardAdded)
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
		it.Event = new(RegistrationRewardAdded)
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
func (it *RegistrationRewardAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RegistrationRewardAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RegistrationRewardAdded represents a RewardAdded event raised by the Registration contract.
type RegistrationRewardAdded struct {
	RaceId      [32]byte
	NftRewardId *big.Int
	ResultIndex [1]byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterRewardAdded is a free log retrieval operation binding the contract event 0x2b942e265bd6d4ca5617a327f9a60e1a85f4351aaa0be2a84cda1cf9d32c8c04.
//
// Solidity: event RewardAdded(bytes32 raceId, uint256 nftRewardId, bytes1 resultIndex)
func (_Registration *RegistrationFilterer) FilterRewardAdded(opts *bind.FilterOpts) (*RegistrationRewardAddedIterator, error) {

	logs, sub, err := _Registration.contract.FilterLogs(opts, "RewardAdded")
	if err != nil {
		return nil, err
	}
	return &RegistrationRewardAddedIterator{contract: _Registration.contract, event: "RewardAdded", logs: logs, sub: sub}, nil
}

// WatchRewardAdded is a free log subscription operation binding the contract event 0x2b942e265bd6d4ca5617a327f9a60e1a85f4351aaa0be2a84cda1cf9d32c8c04.
//
// Solidity: event RewardAdded(bytes32 raceId, uint256 nftRewardId, bytes1 resultIndex)
func (_Registration *RegistrationFilterer) WatchRewardAdded(opts *bind.WatchOpts, sink chan<- *RegistrationRewardAdded) (event.Subscription, error) {

	logs, sub, err := _Registration.contract.WatchLogs(opts, "RewardAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RegistrationRewardAdded)
				if err := _Registration.contract.UnpackLog(event, "RewardAdded", log); err != nil {
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
func (_Registration *RegistrationFilterer) ParseRewardAdded(log types.Log) (*RegistrationRewardAdded, error) {
	event := new(RegistrationRewardAdded)
	if err := _Registration.contract.UnpackLog(event, "RewardAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RegistrationRewardReceivedIterator is returned from FilterRewardReceived and is used to iterate over the raw logs and unpacked data for RewardReceived events raised by the Registration contract.
type RegistrationRewardReceivedIterator struct {
	Event *RegistrationRewardReceived // Event containing the contract specifics and raw log

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
func (it *RegistrationRewardReceivedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RegistrationRewardReceived)
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
		it.Event = new(RegistrationRewardReceived)
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
func (it *RegistrationRewardReceivedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RegistrationRewardReceivedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RegistrationRewardReceived represents a RewardReceived event raised by the Registration contract.
type RegistrationRewardReceived struct {
	RaceId      [32]byte
	SlotId      *big.Int
	NftRewardId *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterRewardReceived is a free log retrieval operation binding the contract event 0xe2ece170b35ae79cbe45ffdaf4931b8bc45056da5d3e9aa4aea67fcdf307929a.
//
// Solidity: event RewardReceived(bytes32 raceId, uint256 slotId, uint256 nftRewardId)
func (_Registration *RegistrationFilterer) FilterRewardReceived(opts *bind.FilterOpts) (*RegistrationRewardReceivedIterator, error) {

	logs, sub, err := _Registration.contract.FilterLogs(opts, "RewardReceived")
	if err != nil {
		return nil, err
	}
	return &RegistrationRewardReceivedIterator{contract: _Registration.contract, event: "RewardReceived", logs: logs, sub: sub}, nil
}

// WatchRewardReceived is a free log subscription operation binding the contract event 0xe2ece170b35ae79cbe45ffdaf4931b8bc45056da5d3e9aa4aea67fcdf307929a.
//
// Solidity: event RewardReceived(bytes32 raceId, uint256 slotId, uint256 nftRewardId)
func (_Registration *RegistrationFilterer) WatchRewardReceived(opts *bind.WatchOpts, sink chan<- *RegistrationRewardReceived) (event.Subscription, error) {

	logs, sub, err := _Registration.contract.WatchLogs(opts, "RewardReceived")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RegistrationRewardReceived)
				if err := _Registration.contract.UnpackLog(event, "RewardReceived", log); err != nil {
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
func (_Registration *RegistrationFilterer) ParseRewardReceived(log types.Log) (*RegistrationRewardReceived, error) {
	event := new(RegistrationRewardReceived)
	if err := _Registration.contract.UnpackLog(event, "RewardReceived", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RegistrationRewardRemovedIterator is returned from FilterRewardRemoved and is used to iterate over the raw logs and unpacked data for RewardRemoved events raised by the Registration contract.
type RegistrationRewardRemovedIterator struct {
	Event *RegistrationRewardRemoved // Event containing the contract specifics and raw log

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
func (it *RegistrationRewardRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RegistrationRewardRemoved)
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
		it.Event = new(RegistrationRewardRemoved)
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
func (it *RegistrationRewardRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RegistrationRewardRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RegistrationRewardRemoved represents a RewardRemoved event raised by the Registration contract.
type RegistrationRewardRemoved struct {
	RaceId      [32]byte
	NftRewardId *big.Int
	ResultIndex [1]byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterRewardRemoved is a free log retrieval operation binding the contract event 0xddda89b96dee97fe43f9803c50584f6dd667bedc8b02a7c75407e6906bf31ead.
//
// Solidity: event RewardRemoved(bytes32 raceId, uint256 nftRewardId, bytes1 resultIndex)
func (_Registration *RegistrationFilterer) FilterRewardRemoved(opts *bind.FilterOpts) (*RegistrationRewardRemovedIterator, error) {

	logs, sub, err := _Registration.contract.FilterLogs(opts, "RewardRemoved")
	if err != nil {
		return nil, err
	}
	return &RegistrationRewardRemovedIterator{contract: _Registration.contract, event: "RewardRemoved", logs: logs, sub: sub}, nil
}

// WatchRewardRemoved is a free log subscription operation binding the contract event 0xddda89b96dee97fe43f9803c50584f6dd667bedc8b02a7c75407e6906bf31ead.
//
// Solidity: event RewardRemoved(bytes32 raceId, uint256 nftRewardId, bytes1 resultIndex)
func (_Registration *RegistrationFilterer) WatchRewardRemoved(opts *bind.WatchOpts, sink chan<- *RegistrationRewardRemoved) (event.Subscription, error) {

	logs, sub, err := _Registration.contract.WatchLogs(opts, "RewardRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RegistrationRewardRemoved)
				if err := _Registration.contract.UnpackLog(event, "RewardRemoved", log); err != nil {
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
func (_Registration *RegistrationFilterer) ParseRewardRemoved(log types.Log) (*RegistrationRewardRemoved, error) {
	event := new(RegistrationRewardRemoved)
	if err := _Registration.contract.UnpackLog(event, "RewardRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RegistrationRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the Registration contract.
type RegistrationRoleAdminChangedIterator struct {
	Event *RegistrationRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *RegistrationRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RegistrationRoleAdminChanged)
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
		it.Event = new(RegistrationRoleAdminChanged)
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
func (it *RegistrationRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RegistrationRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RegistrationRoleAdminChanged represents a RoleAdminChanged event raised by the Registration contract.
type RegistrationRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Registration *RegistrationFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*RegistrationRoleAdminChangedIterator, error) {

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

	logs, sub, err := _Registration.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &RegistrationRoleAdminChangedIterator{contract: _Registration.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Registration *RegistrationFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *RegistrationRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _Registration.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RegistrationRoleAdminChanged)
				if err := _Registration.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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
func (_Registration *RegistrationFilterer) ParseRoleAdminChanged(log types.Log) (*RegistrationRoleAdminChanged, error) {
	event := new(RegistrationRoleAdminChanged)
	if err := _Registration.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RegistrationRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the Registration contract.
type RegistrationRoleGrantedIterator struct {
	Event *RegistrationRoleGranted // Event containing the contract specifics and raw log

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
func (it *RegistrationRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RegistrationRoleGranted)
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
		it.Event = new(RegistrationRoleGranted)
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
func (it *RegistrationRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RegistrationRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RegistrationRoleGranted represents a RoleGranted event raised by the Registration contract.
type RegistrationRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Registration *RegistrationFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*RegistrationRoleGrantedIterator, error) {

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

	logs, sub, err := _Registration.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &RegistrationRoleGrantedIterator{contract: _Registration.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Registration *RegistrationFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *RegistrationRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _Registration.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RegistrationRoleGranted)
				if err := _Registration.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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
func (_Registration *RegistrationFilterer) ParseRoleGranted(log types.Log) (*RegistrationRoleGranted, error) {
	event := new(RegistrationRoleGranted)
	if err := _Registration.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RegistrationRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the Registration contract.
type RegistrationRoleRevokedIterator struct {
	Event *RegistrationRoleRevoked // Event containing the contract specifics and raw log

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
func (it *RegistrationRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RegistrationRoleRevoked)
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
		it.Event = new(RegistrationRoleRevoked)
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
func (it *RegistrationRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RegistrationRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RegistrationRoleRevoked represents a RoleRevoked event raised by the Registration contract.
type RegistrationRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Registration *RegistrationFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*RegistrationRoleRevokedIterator, error) {

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

	logs, sub, err := _Registration.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &RegistrationRoleRevokedIterator{contract: _Registration.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Registration *RegistrationFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *RegistrationRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _Registration.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RegistrationRoleRevoked)
				if err := _Registration.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
func (_Registration *RegistrationFilterer) ParseRoleRevoked(log types.Log) (*RegistrationRoleRevoked, error) {
	event := new(RegistrationRoleRevoked)
	if err := _Registration.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RegistrationUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the Registration contract.
type RegistrationUnpausedIterator struct {
	Event *RegistrationUnpaused // Event containing the contract specifics and raw log

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
func (it *RegistrationUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RegistrationUnpaused)
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
		it.Event = new(RegistrationUnpaused)
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
func (it *RegistrationUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RegistrationUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RegistrationUnpaused represents a Unpaused event raised by the Registration contract.
type RegistrationUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Registration *RegistrationFilterer) FilterUnpaused(opts *bind.FilterOpts) (*RegistrationUnpausedIterator, error) {

	logs, sub, err := _Registration.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &RegistrationUnpausedIterator{contract: _Registration.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Registration *RegistrationFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *RegistrationUnpaused) (event.Subscription, error) {

	logs, sub, err := _Registration.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RegistrationUnpaused)
				if err := _Registration.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
func (_Registration *RegistrationFilterer) ParseUnpaused(log types.Log) (*RegistrationUnpaused, error) {
	event := new(RegistrationUnpaused)
	if err := _Registration.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
