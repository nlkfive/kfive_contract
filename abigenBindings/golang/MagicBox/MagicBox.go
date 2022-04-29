// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package MagicBox

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

// MagicBoxMetaData contains all meta data concerning the MagicBox contract.
var MagicBoxMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"acceptToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"vrfCoordinator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"magicBoxReward\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"endedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"subscriptionId\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"keyHash\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"BoxRequested\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidContract\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoReward\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"have\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"want\",\"type\":\"address\"}],\"name\":\"OnlyCoordinatorCanFulfill\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OpenFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RandomnessNotGenerated\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"}],\"name\":\"TooLate\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"requestId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"randomness\",\"type\":\"uint256\"}],\"name\":\"BoxReceived\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"enumMagicBox.BoxType\",\"name\":\"boxType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"requestId\",\"type\":\"uint256\"}],\"name\":\"RequestBoxSuccessful\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"enumMagicBox.BoxType\",\"name\":\"boxType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"rewardIndex\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nftRewardId\",\"type\":\"uint256\"}],\"name\":\"RewardAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"rewardIndex\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"RewardTransfered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PAUSER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getRoleMember\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleMemberCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"requestId\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"randomWords\",\"type\":\"uint256[]\"}],\"name\":\"rawFulfillRandomWords\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumMagicBox.BoxType\",\"name\":\"boxType\",\"type\":\"uint8\"}],\"name\":\"requestBox\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"requestId\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"requestId\",\"type\":\"uint256\"}],\"name\":\"openBox\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumMagicBox.BoxType\",\"name\":\"boxType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"nftRewardId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"tokenURI\",\"type\":\"string\"}],\"name\":\"addReward\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumMagicBox.BoxType\",\"name\":\"boxType\",\"type\":\"uint8\"}],\"name\":\"remainReward\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumMagicBox.BoxType\",\"name\":\"boxType\",\"type\":\"uint8\"}],\"name\":\"listReward\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"getRequestId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getEndedAt\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC721Received\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]",
}

// MagicBoxABI is the input ABI used to generate the binding from.
// Deprecated: Use MagicBoxMetaData.ABI instead.
var MagicBoxABI = MagicBoxMetaData.ABI

// MagicBox is an auto generated Go binding around an Ethereum contract.
type MagicBox struct {
	MagicBoxCaller     // Read-only binding to the contract
	MagicBoxTransactor // Write-only binding to the contract
	MagicBoxFilterer   // Log filterer for contract events
}

// MagicBoxCaller is an auto generated read-only Go binding around an Ethereum contract.
type MagicBoxCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MagicBoxTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MagicBoxTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MagicBoxFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MagicBoxFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MagicBoxSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MagicBoxSession struct {
	Contract     *MagicBox         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MagicBoxCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MagicBoxCallerSession struct {
	Contract *MagicBoxCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// MagicBoxTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MagicBoxTransactorSession struct {
	Contract     *MagicBoxTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// MagicBoxRaw is an auto generated low-level Go binding around an Ethereum contract.
type MagicBoxRaw struct {
	Contract *MagicBox // Generic contract binding to access the raw methods on
}

// MagicBoxCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MagicBoxCallerRaw struct {
	Contract *MagicBoxCaller // Generic read-only contract binding to access the raw methods on
}

// MagicBoxTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MagicBoxTransactorRaw struct {
	Contract *MagicBoxTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMagicBox creates a new instance of MagicBox, bound to a specific deployed contract.
func NewMagicBox(address common.Address, backend bind.ContractBackend) (*MagicBox, error) {
	contract, err := bindMagicBox(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MagicBox{MagicBoxCaller: MagicBoxCaller{contract: contract}, MagicBoxTransactor: MagicBoxTransactor{contract: contract}, MagicBoxFilterer: MagicBoxFilterer{contract: contract}}, nil
}

// NewMagicBoxCaller creates a new read-only instance of MagicBox, bound to a specific deployed contract.
func NewMagicBoxCaller(address common.Address, caller bind.ContractCaller) (*MagicBoxCaller, error) {
	contract, err := bindMagicBox(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MagicBoxCaller{contract: contract}, nil
}

// NewMagicBoxTransactor creates a new write-only instance of MagicBox, bound to a specific deployed contract.
func NewMagicBoxTransactor(address common.Address, transactor bind.ContractTransactor) (*MagicBoxTransactor, error) {
	contract, err := bindMagicBox(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MagicBoxTransactor{contract: contract}, nil
}

// NewMagicBoxFilterer creates a new log filterer instance of MagicBox, bound to a specific deployed contract.
func NewMagicBoxFilterer(address common.Address, filterer bind.ContractFilterer) (*MagicBoxFilterer, error) {
	contract, err := bindMagicBox(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MagicBoxFilterer{contract: contract}, nil
}

// bindMagicBox binds a generic wrapper to an already deployed contract.
func bindMagicBox(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MagicBoxABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MagicBox *MagicBoxRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MagicBox.Contract.MagicBoxCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MagicBox *MagicBoxRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MagicBox.Contract.MagicBoxTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MagicBox *MagicBoxRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MagicBox.Contract.MagicBoxTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MagicBox *MagicBoxCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MagicBox.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MagicBox *MagicBoxTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MagicBox.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MagicBox *MagicBoxTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MagicBox.Contract.contract.Transact(opts, method, params...)
}

// ADMINROLE is a free data retrieval call binding the contract method 0x75b238fc.
//
// Solidity: function ADMIN_ROLE() view returns(bytes32)
func (_MagicBox *MagicBoxCaller) ADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _MagicBox.contract.Call(opts, &out, "ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ADMINROLE is a free data retrieval call binding the contract method 0x75b238fc.
//
// Solidity: function ADMIN_ROLE() view returns(bytes32)
func (_MagicBox *MagicBoxSession) ADMINROLE() ([32]byte, error) {
	return _MagicBox.Contract.ADMINROLE(&_MagicBox.CallOpts)
}

// ADMINROLE is a free data retrieval call binding the contract method 0x75b238fc.
//
// Solidity: function ADMIN_ROLE() view returns(bytes32)
func (_MagicBox *MagicBoxCallerSession) ADMINROLE() ([32]byte, error) {
	return _MagicBox.Contract.ADMINROLE(&_MagicBox.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_MagicBox *MagicBoxCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _MagicBox.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_MagicBox *MagicBoxSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _MagicBox.Contract.DEFAULTADMINROLE(&_MagicBox.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_MagicBox *MagicBoxCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _MagicBox.Contract.DEFAULTADMINROLE(&_MagicBox.CallOpts)
}

// PAUSERROLE is a free data retrieval call binding the contract method 0xe63ab1e9.
//
// Solidity: function PAUSER_ROLE() view returns(bytes32)
func (_MagicBox *MagicBoxCaller) PAUSERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _MagicBox.contract.Call(opts, &out, "PAUSER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// PAUSERROLE is a free data retrieval call binding the contract method 0xe63ab1e9.
//
// Solidity: function PAUSER_ROLE() view returns(bytes32)
func (_MagicBox *MagicBoxSession) PAUSERROLE() ([32]byte, error) {
	return _MagicBox.Contract.PAUSERROLE(&_MagicBox.CallOpts)
}

// PAUSERROLE is a free data retrieval call binding the contract method 0xe63ab1e9.
//
// Solidity: function PAUSER_ROLE() view returns(bytes32)
func (_MagicBox *MagicBoxCallerSession) PAUSERROLE() ([32]byte, error) {
	return _MagicBox.Contract.PAUSERROLE(&_MagicBox.CallOpts)
}

// GetEndedAt is a free data retrieval call binding the contract method 0x1906ab9b.
//
// Solidity: function getEndedAt() view returns(uint256)
func (_MagicBox *MagicBoxCaller) GetEndedAt(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MagicBox.contract.Call(opts, &out, "getEndedAt")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetEndedAt is a free data retrieval call binding the contract method 0x1906ab9b.
//
// Solidity: function getEndedAt() view returns(uint256)
func (_MagicBox *MagicBoxSession) GetEndedAt() (*big.Int, error) {
	return _MagicBox.Contract.GetEndedAt(&_MagicBox.CallOpts)
}

// GetEndedAt is a free data retrieval call binding the contract method 0x1906ab9b.
//
// Solidity: function getEndedAt() view returns(uint256)
func (_MagicBox *MagicBoxCallerSession) GetEndedAt() (*big.Int, error) {
	return _MagicBox.Contract.GetEndedAt(&_MagicBox.CallOpts)
}

// GetRequestId is a free data retrieval call binding the contract method 0x1e612da8.
//
// Solidity: function getRequestId(address sender) view returns(uint256)
func (_MagicBox *MagicBoxCaller) GetRequestId(opts *bind.CallOpts, sender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _MagicBox.contract.Call(opts, &out, "getRequestId", sender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRequestId is a free data retrieval call binding the contract method 0x1e612da8.
//
// Solidity: function getRequestId(address sender) view returns(uint256)
func (_MagicBox *MagicBoxSession) GetRequestId(sender common.Address) (*big.Int, error) {
	return _MagicBox.Contract.GetRequestId(&_MagicBox.CallOpts, sender)
}

// GetRequestId is a free data retrieval call binding the contract method 0x1e612da8.
//
// Solidity: function getRequestId(address sender) view returns(uint256)
func (_MagicBox *MagicBoxCallerSession) GetRequestId(sender common.Address) (*big.Int, error) {
	return _MagicBox.Contract.GetRequestId(&_MagicBox.CallOpts, sender)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_MagicBox *MagicBoxCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _MagicBox.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_MagicBox *MagicBoxSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _MagicBox.Contract.GetRoleAdmin(&_MagicBox.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_MagicBox *MagicBoxCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _MagicBox.Contract.GetRoleAdmin(&_MagicBox.CallOpts, role)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_MagicBox *MagicBoxCaller) GetRoleMember(opts *bind.CallOpts, role [32]byte, index *big.Int) (common.Address, error) {
	var out []interface{}
	err := _MagicBox.contract.Call(opts, &out, "getRoleMember", role, index)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_MagicBox *MagicBoxSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _MagicBox.Contract.GetRoleMember(&_MagicBox.CallOpts, role, index)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_MagicBox *MagicBoxCallerSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _MagicBox.Contract.GetRoleMember(&_MagicBox.CallOpts, role, index)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_MagicBox *MagicBoxCaller) GetRoleMemberCount(opts *bind.CallOpts, role [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _MagicBox.contract.Call(opts, &out, "getRoleMemberCount", role)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_MagicBox *MagicBoxSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _MagicBox.Contract.GetRoleMemberCount(&_MagicBox.CallOpts, role)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_MagicBox *MagicBoxCallerSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _MagicBox.Contract.GetRoleMemberCount(&_MagicBox.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_MagicBox *MagicBoxCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _MagicBox.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_MagicBox *MagicBoxSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _MagicBox.Contract.HasRole(&_MagicBox.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_MagicBox *MagicBoxCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _MagicBox.Contract.HasRole(&_MagicBox.CallOpts, role, account)
}

// ListReward is a free data retrieval call binding the contract method 0xd86eab47.
//
// Solidity: function listReward(uint8 boxType) view returns(uint256[])
func (_MagicBox *MagicBoxCaller) ListReward(opts *bind.CallOpts, boxType uint8) ([]*big.Int, error) {
	var out []interface{}
	err := _MagicBox.contract.Call(opts, &out, "listReward", boxType)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// ListReward is a free data retrieval call binding the contract method 0xd86eab47.
//
// Solidity: function listReward(uint8 boxType) view returns(uint256[])
func (_MagicBox *MagicBoxSession) ListReward(boxType uint8) ([]*big.Int, error) {
	return _MagicBox.Contract.ListReward(&_MagicBox.CallOpts, boxType)
}

// ListReward is a free data retrieval call binding the contract method 0xd86eab47.
//
// Solidity: function listReward(uint8 boxType) view returns(uint256[])
func (_MagicBox *MagicBoxCallerSession) ListReward(boxType uint8) ([]*big.Int, error) {
	return _MagicBox.Contract.ListReward(&_MagicBox.CallOpts, boxType)
}

// OnERC721Received is a free data retrieval call binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) pure returns(bytes4)
func (_MagicBox *MagicBoxCaller) OnERC721Received(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) ([4]byte, error) {
	var out []interface{}
	err := _MagicBox.contract.Call(opts, &out, "onERC721Received", arg0, arg1, arg2, arg3)

	if err != nil {
		return *new([4]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)

	return out0, err

}

// OnERC721Received is a free data retrieval call binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) pure returns(bytes4)
func (_MagicBox *MagicBoxSession) OnERC721Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) ([4]byte, error) {
	return _MagicBox.Contract.OnERC721Received(&_MagicBox.CallOpts, arg0, arg1, arg2, arg3)
}

// OnERC721Received is a free data retrieval call binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) pure returns(bytes4)
func (_MagicBox *MagicBoxCallerSession) OnERC721Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) ([4]byte, error) {
	return _MagicBox.Contract.OnERC721Received(&_MagicBox.CallOpts, arg0, arg1, arg2, arg3)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_MagicBox *MagicBoxCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MagicBox.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_MagicBox *MagicBoxSession) Owner() (common.Address, error) {
	return _MagicBox.Contract.Owner(&_MagicBox.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_MagicBox *MagicBoxCallerSession) Owner() (common.Address, error) {
	return _MagicBox.Contract.Owner(&_MagicBox.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_MagicBox *MagicBoxCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _MagicBox.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_MagicBox *MagicBoxSession) Paused() (bool, error) {
	return _MagicBox.Contract.Paused(&_MagicBox.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_MagicBox *MagicBoxCallerSession) Paused() (bool, error) {
	return _MagicBox.Contract.Paused(&_MagicBox.CallOpts)
}

// RemainReward is a free data retrieval call binding the contract method 0xa9d7fc21.
//
// Solidity: function remainReward(uint8 boxType) view returns(uint256)
func (_MagicBox *MagicBoxCaller) RemainReward(opts *bind.CallOpts, boxType uint8) (*big.Int, error) {
	var out []interface{}
	err := _MagicBox.contract.Call(opts, &out, "remainReward", boxType)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RemainReward is a free data retrieval call binding the contract method 0xa9d7fc21.
//
// Solidity: function remainReward(uint8 boxType) view returns(uint256)
func (_MagicBox *MagicBoxSession) RemainReward(boxType uint8) (*big.Int, error) {
	return _MagicBox.Contract.RemainReward(&_MagicBox.CallOpts, boxType)
}

// RemainReward is a free data retrieval call binding the contract method 0xa9d7fc21.
//
// Solidity: function remainReward(uint8 boxType) view returns(uint256)
func (_MagicBox *MagicBoxCallerSession) RemainReward(boxType uint8) (*big.Int, error) {
	return _MagicBox.Contract.RemainReward(&_MagicBox.CallOpts, boxType)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_MagicBox *MagicBoxCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _MagicBox.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_MagicBox *MagicBoxSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _MagicBox.Contract.SupportsInterface(&_MagicBox.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_MagicBox *MagicBoxCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _MagicBox.Contract.SupportsInterface(&_MagicBox.CallOpts, interfaceId)
}

// AddReward is a paid mutator transaction binding the contract method 0x9a136697.
//
// Solidity: function addReward(uint8 boxType, uint256 nftRewardId, string tokenURI) returns()
func (_MagicBox *MagicBoxTransactor) AddReward(opts *bind.TransactOpts, boxType uint8, nftRewardId *big.Int, tokenURI string) (*types.Transaction, error) {
	return _MagicBox.contract.Transact(opts, "addReward", boxType, nftRewardId, tokenURI)
}

// AddReward is a paid mutator transaction binding the contract method 0x9a136697.
//
// Solidity: function addReward(uint8 boxType, uint256 nftRewardId, string tokenURI) returns()
func (_MagicBox *MagicBoxSession) AddReward(boxType uint8, nftRewardId *big.Int, tokenURI string) (*types.Transaction, error) {
	return _MagicBox.Contract.AddReward(&_MagicBox.TransactOpts, boxType, nftRewardId, tokenURI)
}

// AddReward is a paid mutator transaction binding the contract method 0x9a136697.
//
// Solidity: function addReward(uint8 boxType, uint256 nftRewardId, string tokenURI) returns()
func (_MagicBox *MagicBoxTransactorSession) AddReward(boxType uint8, nftRewardId *big.Int, tokenURI string) (*types.Transaction, error) {
	return _MagicBox.Contract.AddReward(&_MagicBox.TransactOpts, boxType, nftRewardId, tokenURI)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_MagicBox *MagicBoxTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _MagicBox.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_MagicBox *MagicBoxSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _MagicBox.Contract.GrantRole(&_MagicBox.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_MagicBox *MagicBoxTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _MagicBox.Contract.GrantRole(&_MagicBox.TransactOpts, role, account)
}

// OpenBox is a paid mutator transaction binding the contract method 0xb1e5e2b7.
//
// Solidity: function openBox(uint256 requestId) returns()
func (_MagicBox *MagicBoxTransactor) OpenBox(opts *bind.TransactOpts, requestId *big.Int) (*types.Transaction, error) {
	return _MagicBox.contract.Transact(opts, "openBox", requestId)
}

// OpenBox is a paid mutator transaction binding the contract method 0xb1e5e2b7.
//
// Solidity: function openBox(uint256 requestId) returns()
func (_MagicBox *MagicBoxSession) OpenBox(requestId *big.Int) (*types.Transaction, error) {
	return _MagicBox.Contract.OpenBox(&_MagicBox.TransactOpts, requestId)
}

// OpenBox is a paid mutator transaction binding the contract method 0xb1e5e2b7.
//
// Solidity: function openBox(uint256 requestId) returns()
func (_MagicBox *MagicBoxTransactorSession) OpenBox(requestId *big.Int) (*types.Transaction, error) {
	return _MagicBox.Contract.OpenBox(&_MagicBox.TransactOpts, requestId)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_MagicBox *MagicBoxTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MagicBox.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_MagicBox *MagicBoxSession) Pause() (*types.Transaction, error) {
	return _MagicBox.Contract.Pause(&_MagicBox.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_MagicBox *MagicBoxTransactorSession) Pause() (*types.Transaction, error) {
	return _MagicBox.Contract.Pause(&_MagicBox.TransactOpts)
}

// RawFulfillRandomWords is a paid mutator transaction binding the contract method 0x1fe543e3.
//
// Solidity: function rawFulfillRandomWords(uint256 requestId, uint256[] randomWords) returns()
func (_MagicBox *MagicBoxTransactor) RawFulfillRandomWords(opts *bind.TransactOpts, requestId *big.Int, randomWords []*big.Int) (*types.Transaction, error) {
	return _MagicBox.contract.Transact(opts, "rawFulfillRandomWords", requestId, randomWords)
}

// RawFulfillRandomWords is a paid mutator transaction binding the contract method 0x1fe543e3.
//
// Solidity: function rawFulfillRandomWords(uint256 requestId, uint256[] randomWords) returns()
func (_MagicBox *MagicBoxSession) RawFulfillRandomWords(requestId *big.Int, randomWords []*big.Int) (*types.Transaction, error) {
	return _MagicBox.Contract.RawFulfillRandomWords(&_MagicBox.TransactOpts, requestId, randomWords)
}

// RawFulfillRandomWords is a paid mutator transaction binding the contract method 0x1fe543e3.
//
// Solidity: function rawFulfillRandomWords(uint256 requestId, uint256[] randomWords) returns()
func (_MagicBox *MagicBoxTransactorSession) RawFulfillRandomWords(requestId *big.Int, randomWords []*big.Int) (*types.Transaction, error) {
	return _MagicBox.Contract.RawFulfillRandomWords(&_MagicBox.TransactOpts, requestId, randomWords)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_MagicBox *MagicBoxTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MagicBox.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_MagicBox *MagicBoxSession) RenounceOwnership() (*types.Transaction, error) {
	return _MagicBox.Contract.RenounceOwnership(&_MagicBox.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_MagicBox *MagicBoxTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _MagicBox.Contract.RenounceOwnership(&_MagicBox.TransactOpts)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_MagicBox *MagicBoxTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _MagicBox.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_MagicBox *MagicBoxSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _MagicBox.Contract.RenounceRole(&_MagicBox.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_MagicBox *MagicBoxTransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _MagicBox.Contract.RenounceRole(&_MagicBox.TransactOpts, role, account)
}

// RequestBox is a paid mutator transaction binding the contract method 0xea0192a0.
//
// Solidity: function requestBox(uint8 boxType) returns(uint256 requestId)
func (_MagicBox *MagicBoxTransactor) RequestBox(opts *bind.TransactOpts, boxType uint8) (*types.Transaction, error) {
	return _MagicBox.contract.Transact(opts, "requestBox", boxType)
}

// RequestBox is a paid mutator transaction binding the contract method 0xea0192a0.
//
// Solidity: function requestBox(uint8 boxType) returns(uint256 requestId)
func (_MagicBox *MagicBoxSession) RequestBox(boxType uint8) (*types.Transaction, error) {
	return _MagicBox.Contract.RequestBox(&_MagicBox.TransactOpts, boxType)
}

// RequestBox is a paid mutator transaction binding the contract method 0xea0192a0.
//
// Solidity: function requestBox(uint8 boxType) returns(uint256 requestId)
func (_MagicBox *MagicBoxTransactorSession) RequestBox(boxType uint8) (*types.Transaction, error) {
	return _MagicBox.Contract.RequestBox(&_MagicBox.TransactOpts, boxType)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_MagicBox *MagicBoxTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _MagicBox.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_MagicBox *MagicBoxSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _MagicBox.Contract.RevokeRole(&_MagicBox.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_MagicBox *MagicBoxTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _MagicBox.Contract.RevokeRole(&_MagicBox.TransactOpts, role, account)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_MagicBox *MagicBoxTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _MagicBox.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_MagicBox *MagicBoxSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _MagicBox.Contract.TransferOwnership(&_MagicBox.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_MagicBox *MagicBoxTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _MagicBox.Contract.TransferOwnership(&_MagicBox.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_MagicBox *MagicBoxTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MagicBox.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_MagicBox *MagicBoxSession) Unpause() (*types.Transaction, error) {
	return _MagicBox.Contract.Unpause(&_MagicBox.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_MagicBox *MagicBoxTransactorSession) Unpause() (*types.Transaction, error) {
	return _MagicBox.Contract.Unpause(&_MagicBox.TransactOpts)
}

// MagicBoxBoxReceivedIterator is returned from FilterBoxReceived and is used to iterate over the raw logs and unpacked data for BoxReceived events raised by the MagicBox contract.
type MagicBoxBoxReceivedIterator struct {
	Event *MagicBoxBoxReceived // Event containing the contract specifics and raw log

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
func (it *MagicBoxBoxReceivedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MagicBoxBoxReceived)
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
		it.Event = new(MagicBoxBoxReceived)
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
func (it *MagicBoxBoxReceivedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MagicBoxBoxReceivedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MagicBoxBoxReceived represents a BoxReceived event raised by the MagicBox contract.
type MagicBoxBoxReceived struct {
	RequestId  *big.Int
	Randomness *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterBoxReceived is a free log retrieval operation binding the contract event 0x45b6eedd25b1b5c371d24f0338eccfea7dacaca789cf8dc0694d8fd308dc25c1.
//
// Solidity: event BoxReceived(uint256 requestId, uint256 randomness)
func (_MagicBox *MagicBoxFilterer) FilterBoxReceived(opts *bind.FilterOpts) (*MagicBoxBoxReceivedIterator, error) {

	logs, sub, err := _MagicBox.contract.FilterLogs(opts, "BoxReceived")
	if err != nil {
		return nil, err
	}
	return &MagicBoxBoxReceivedIterator{contract: _MagicBox.contract, event: "BoxReceived", logs: logs, sub: sub}, nil
}

// WatchBoxReceived is a free log subscription operation binding the contract event 0x45b6eedd25b1b5c371d24f0338eccfea7dacaca789cf8dc0694d8fd308dc25c1.
//
// Solidity: event BoxReceived(uint256 requestId, uint256 randomness)
func (_MagicBox *MagicBoxFilterer) WatchBoxReceived(opts *bind.WatchOpts, sink chan<- *MagicBoxBoxReceived) (event.Subscription, error) {

	logs, sub, err := _MagicBox.contract.WatchLogs(opts, "BoxReceived")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MagicBoxBoxReceived)
				if err := _MagicBox.contract.UnpackLog(event, "BoxReceived", log); err != nil {
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

// ParseBoxReceived is a log parse operation binding the contract event 0x45b6eedd25b1b5c371d24f0338eccfea7dacaca789cf8dc0694d8fd308dc25c1.
//
// Solidity: event BoxReceived(uint256 requestId, uint256 randomness)
func (_MagicBox *MagicBoxFilterer) ParseBoxReceived(log types.Log) (*MagicBoxBoxReceived, error) {
	event := new(MagicBoxBoxReceived)
	if err := _MagicBox.contract.UnpackLog(event, "BoxReceived", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MagicBoxOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the MagicBox contract.
type MagicBoxOwnershipTransferredIterator struct {
	Event *MagicBoxOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *MagicBoxOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MagicBoxOwnershipTransferred)
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
		it.Event = new(MagicBoxOwnershipTransferred)
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
func (it *MagicBoxOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MagicBoxOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MagicBoxOwnershipTransferred represents a OwnershipTransferred event raised by the MagicBox contract.
type MagicBoxOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_MagicBox *MagicBoxFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*MagicBoxOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _MagicBox.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &MagicBoxOwnershipTransferredIterator{contract: _MagicBox.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_MagicBox *MagicBoxFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *MagicBoxOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _MagicBox.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MagicBoxOwnershipTransferred)
				if err := _MagicBox.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_MagicBox *MagicBoxFilterer) ParseOwnershipTransferred(log types.Log) (*MagicBoxOwnershipTransferred, error) {
	event := new(MagicBoxOwnershipTransferred)
	if err := _MagicBox.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MagicBoxPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the MagicBox contract.
type MagicBoxPausedIterator struct {
	Event *MagicBoxPaused // Event containing the contract specifics and raw log

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
func (it *MagicBoxPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MagicBoxPaused)
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
		it.Event = new(MagicBoxPaused)
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
func (it *MagicBoxPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MagicBoxPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MagicBoxPaused represents a Paused event raised by the MagicBox contract.
type MagicBoxPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_MagicBox *MagicBoxFilterer) FilterPaused(opts *bind.FilterOpts) (*MagicBoxPausedIterator, error) {

	logs, sub, err := _MagicBox.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &MagicBoxPausedIterator{contract: _MagicBox.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_MagicBox *MagicBoxFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *MagicBoxPaused) (event.Subscription, error) {

	logs, sub, err := _MagicBox.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MagicBoxPaused)
				if err := _MagicBox.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_MagicBox *MagicBoxFilterer) ParsePaused(log types.Log) (*MagicBoxPaused, error) {
	event := new(MagicBoxPaused)
	if err := _MagicBox.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MagicBoxRequestBoxSuccessfulIterator is returned from FilterRequestBoxSuccessful and is used to iterate over the raw logs and unpacked data for RequestBoxSuccessful events raised by the MagicBox contract.
type MagicBoxRequestBoxSuccessfulIterator struct {
	Event *MagicBoxRequestBoxSuccessful // Event containing the contract specifics and raw log

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
func (it *MagicBoxRequestBoxSuccessfulIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MagicBoxRequestBoxSuccessful)
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
		it.Event = new(MagicBoxRequestBoxSuccessful)
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
func (it *MagicBoxRequestBoxSuccessfulIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MagicBoxRequestBoxSuccessfulIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MagicBoxRequestBoxSuccessful represents a RequestBoxSuccessful event raised by the MagicBox contract.
type MagicBoxRequestBoxSuccessful struct {
	BoxType   uint8
	Sender    common.Address
	RequestId *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterRequestBoxSuccessful is a free log retrieval operation binding the contract event 0x563346928a474bb9aa88267cb3c79d1a3520fc4d63177ad5eed057408ed5ecf5.
//
// Solidity: event RequestBoxSuccessful(uint8 boxType, address sender, uint256 requestId)
func (_MagicBox *MagicBoxFilterer) FilterRequestBoxSuccessful(opts *bind.FilterOpts) (*MagicBoxRequestBoxSuccessfulIterator, error) {

	logs, sub, err := _MagicBox.contract.FilterLogs(opts, "RequestBoxSuccessful")
	if err != nil {
		return nil, err
	}
	return &MagicBoxRequestBoxSuccessfulIterator{contract: _MagicBox.contract, event: "RequestBoxSuccessful", logs: logs, sub: sub}, nil
}

// WatchRequestBoxSuccessful is a free log subscription operation binding the contract event 0x563346928a474bb9aa88267cb3c79d1a3520fc4d63177ad5eed057408ed5ecf5.
//
// Solidity: event RequestBoxSuccessful(uint8 boxType, address sender, uint256 requestId)
func (_MagicBox *MagicBoxFilterer) WatchRequestBoxSuccessful(opts *bind.WatchOpts, sink chan<- *MagicBoxRequestBoxSuccessful) (event.Subscription, error) {

	logs, sub, err := _MagicBox.contract.WatchLogs(opts, "RequestBoxSuccessful")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MagicBoxRequestBoxSuccessful)
				if err := _MagicBox.contract.UnpackLog(event, "RequestBoxSuccessful", log); err != nil {
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

// ParseRequestBoxSuccessful is a log parse operation binding the contract event 0x563346928a474bb9aa88267cb3c79d1a3520fc4d63177ad5eed057408ed5ecf5.
//
// Solidity: event RequestBoxSuccessful(uint8 boxType, address sender, uint256 requestId)
func (_MagicBox *MagicBoxFilterer) ParseRequestBoxSuccessful(log types.Log) (*MagicBoxRequestBoxSuccessful, error) {
	event := new(MagicBoxRequestBoxSuccessful)
	if err := _MagicBox.contract.UnpackLog(event, "RequestBoxSuccessful", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MagicBoxRewardAddedIterator is returned from FilterRewardAdded and is used to iterate over the raw logs and unpacked data for RewardAdded events raised by the MagicBox contract.
type MagicBoxRewardAddedIterator struct {
	Event *MagicBoxRewardAdded // Event containing the contract specifics and raw log

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
func (it *MagicBoxRewardAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MagicBoxRewardAdded)
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
		it.Event = new(MagicBoxRewardAdded)
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
func (it *MagicBoxRewardAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MagicBoxRewardAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MagicBoxRewardAdded represents a RewardAdded event raised by the MagicBox contract.
type MagicBoxRewardAdded struct {
	BoxType     uint8
	RewardIndex *big.Int
	NftRewardId *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterRewardAdded is a free log retrieval operation binding the contract event 0x99f416e25a0f9360ca447cc71598c43e0fc939787e14faa4e7b3bf51ccf5de08.
//
// Solidity: event RewardAdded(uint8 boxType, uint256 rewardIndex, uint256 nftRewardId)
func (_MagicBox *MagicBoxFilterer) FilterRewardAdded(opts *bind.FilterOpts) (*MagicBoxRewardAddedIterator, error) {

	logs, sub, err := _MagicBox.contract.FilterLogs(opts, "RewardAdded")
	if err != nil {
		return nil, err
	}
	return &MagicBoxRewardAddedIterator{contract: _MagicBox.contract, event: "RewardAdded", logs: logs, sub: sub}, nil
}

// WatchRewardAdded is a free log subscription operation binding the contract event 0x99f416e25a0f9360ca447cc71598c43e0fc939787e14faa4e7b3bf51ccf5de08.
//
// Solidity: event RewardAdded(uint8 boxType, uint256 rewardIndex, uint256 nftRewardId)
func (_MagicBox *MagicBoxFilterer) WatchRewardAdded(opts *bind.WatchOpts, sink chan<- *MagicBoxRewardAdded) (event.Subscription, error) {

	logs, sub, err := _MagicBox.contract.WatchLogs(opts, "RewardAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MagicBoxRewardAdded)
				if err := _MagicBox.contract.UnpackLog(event, "RewardAdded", log); err != nil {
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

// ParseRewardAdded is a log parse operation binding the contract event 0x99f416e25a0f9360ca447cc71598c43e0fc939787e14faa4e7b3bf51ccf5de08.
//
// Solidity: event RewardAdded(uint8 boxType, uint256 rewardIndex, uint256 nftRewardId)
func (_MagicBox *MagicBoxFilterer) ParseRewardAdded(log types.Log) (*MagicBoxRewardAdded, error) {
	event := new(MagicBoxRewardAdded)
	if err := _MagicBox.contract.UnpackLog(event, "RewardAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MagicBoxRewardTransferedIterator is returned from FilterRewardTransfered and is used to iterate over the raw logs and unpacked data for RewardTransfered events raised by the MagicBox contract.
type MagicBoxRewardTransferedIterator struct {
	Event *MagicBoxRewardTransfered // Event containing the contract specifics and raw log

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
func (it *MagicBoxRewardTransferedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MagicBoxRewardTransfered)
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
		it.Event = new(MagicBoxRewardTransfered)
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
func (it *MagicBoxRewardTransferedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MagicBoxRewardTransferedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MagicBoxRewardTransfered represents a RewardTransfered event raised by the MagicBox contract.
type MagicBoxRewardTransfered struct {
	Receiver    common.Address
	RewardIndex *big.Int
	TokenId     *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterRewardTransfered is a free log retrieval operation binding the contract event 0xb84c51b5972dee6f0d5b879da24ecd8d4f93168b696aa2b12a5dcf9036656959.
//
// Solidity: event RewardTransfered(address receiver, uint256 rewardIndex, uint256 tokenId)
func (_MagicBox *MagicBoxFilterer) FilterRewardTransfered(opts *bind.FilterOpts) (*MagicBoxRewardTransferedIterator, error) {

	logs, sub, err := _MagicBox.contract.FilterLogs(opts, "RewardTransfered")
	if err != nil {
		return nil, err
	}
	return &MagicBoxRewardTransferedIterator{contract: _MagicBox.contract, event: "RewardTransfered", logs: logs, sub: sub}, nil
}

// WatchRewardTransfered is a free log subscription operation binding the contract event 0xb84c51b5972dee6f0d5b879da24ecd8d4f93168b696aa2b12a5dcf9036656959.
//
// Solidity: event RewardTransfered(address receiver, uint256 rewardIndex, uint256 tokenId)
func (_MagicBox *MagicBoxFilterer) WatchRewardTransfered(opts *bind.WatchOpts, sink chan<- *MagicBoxRewardTransfered) (event.Subscription, error) {

	logs, sub, err := _MagicBox.contract.WatchLogs(opts, "RewardTransfered")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MagicBoxRewardTransfered)
				if err := _MagicBox.contract.UnpackLog(event, "RewardTransfered", log); err != nil {
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

// ParseRewardTransfered is a log parse operation binding the contract event 0xb84c51b5972dee6f0d5b879da24ecd8d4f93168b696aa2b12a5dcf9036656959.
//
// Solidity: event RewardTransfered(address receiver, uint256 rewardIndex, uint256 tokenId)
func (_MagicBox *MagicBoxFilterer) ParseRewardTransfered(log types.Log) (*MagicBoxRewardTransfered, error) {
	event := new(MagicBoxRewardTransfered)
	if err := _MagicBox.contract.UnpackLog(event, "RewardTransfered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MagicBoxRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the MagicBox contract.
type MagicBoxRoleAdminChangedIterator struct {
	Event *MagicBoxRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *MagicBoxRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MagicBoxRoleAdminChanged)
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
		it.Event = new(MagicBoxRoleAdminChanged)
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
func (it *MagicBoxRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MagicBoxRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MagicBoxRoleAdminChanged represents a RoleAdminChanged event raised by the MagicBox contract.
type MagicBoxRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_MagicBox *MagicBoxFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*MagicBoxRoleAdminChangedIterator, error) {

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

	logs, sub, err := _MagicBox.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &MagicBoxRoleAdminChangedIterator{contract: _MagicBox.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_MagicBox *MagicBoxFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *MagicBoxRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _MagicBox.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MagicBoxRoleAdminChanged)
				if err := _MagicBox.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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
func (_MagicBox *MagicBoxFilterer) ParseRoleAdminChanged(log types.Log) (*MagicBoxRoleAdminChanged, error) {
	event := new(MagicBoxRoleAdminChanged)
	if err := _MagicBox.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MagicBoxRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the MagicBox contract.
type MagicBoxRoleGrantedIterator struct {
	Event *MagicBoxRoleGranted // Event containing the contract specifics and raw log

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
func (it *MagicBoxRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MagicBoxRoleGranted)
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
		it.Event = new(MagicBoxRoleGranted)
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
func (it *MagicBoxRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MagicBoxRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MagicBoxRoleGranted represents a RoleGranted event raised by the MagicBox contract.
type MagicBoxRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_MagicBox *MagicBoxFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*MagicBoxRoleGrantedIterator, error) {

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

	logs, sub, err := _MagicBox.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &MagicBoxRoleGrantedIterator{contract: _MagicBox.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_MagicBox *MagicBoxFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *MagicBoxRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _MagicBox.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MagicBoxRoleGranted)
				if err := _MagicBox.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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
func (_MagicBox *MagicBoxFilterer) ParseRoleGranted(log types.Log) (*MagicBoxRoleGranted, error) {
	event := new(MagicBoxRoleGranted)
	if err := _MagicBox.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MagicBoxRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the MagicBox contract.
type MagicBoxRoleRevokedIterator struct {
	Event *MagicBoxRoleRevoked // Event containing the contract specifics and raw log

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
func (it *MagicBoxRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MagicBoxRoleRevoked)
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
		it.Event = new(MagicBoxRoleRevoked)
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
func (it *MagicBoxRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MagicBoxRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MagicBoxRoleRevoked represents a RoleRevoked event raised by the MagicBox contract.
type MagicBoxRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_MagicBox *MagicBoxFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*MagicBoxRoleRevokedIterator, error) {

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

	logs, sub, err := _MagicBox.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &MagicBoxRoleRevokedIterator{contract: _MagicBox.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_MagicBox *MagicBoxFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *MagicBoxRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _MagicBox.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MagicBoxRoleRevoked)
				if err := _MagicBox.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
func (_MagicBox *MagicBoxFilterer) ParseRoleRevoked(log types.Log) (*MagicBoxRoleRevoked, error) {
	event := new(MagicBoxRoleRevoked)
	if err := _MagicBox.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MagicBoxUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the MagicBox contract.
type MagicBoxUnpausedIterator struct {
	Event *MagicBoxUnpaused // Event containing the contract specifics and raw log

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
func (it *MagicBoxUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MagicBoxUnpaused)
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
		it.Event = new(MagicBoxUnpaused)
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
func (it *MagicBoxUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MagicBoxUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MagicBoxUnpaused represents a Unpaused event raised by the MagicBox contract.
type MagicBoxUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_MagicBox *MagicBoxFilterer) FilterUnpaused(opts *bind.FilterOpts) (*MagicBoxUnpausedIterator, error) {

	logs, sub, err := _MagicBox.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &MagicBoxUnpausedIterator{contract: _MagicBox.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_MagicBox *MagicBoxFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *MagicBoxUnpaused) (event.Subscription, error) {

	logs, sub, err := _MagicBox.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MagicBoxUnpaused)
				if err := _MagicBox.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
func (_MagicBox *MagicBoxFilterer) ParseUnpaused(log types.Log) (*MagicBoxUnpaused, error) {
	event := new(MagicBoxUnpaused)
	if err := _MagicBox.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
