// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package NLGGT

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

// NLGGTMetaData contains all meta data concerning the NLGGT contract.
var NLGGTMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_marketplaceStorage\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"baseTokenURI\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"AddedBlackList\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_marketplaceStorage\",\"type\":\"address\"}],\"name\":\"MarketplaceStorageUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"RemovedBlackList\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"IMarketplaceStorage_Interface\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MINTER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PAUSER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_evilUser\",\"type\":\"address\"}],\"name\":\"addBlackList\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_maker\",\"type\":\"address\"}],\"name\":\"getBlackListStatus\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getRoleMember\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleMemberCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"isBlackListed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"marketplaceStorage\",\"outputs\":[{\"internalType\":\"contractIMarketplaceStorage\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"gtokenURI\",\"type\":\"string\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_clearedUser\",\"type\":\"address\"}],\"name\":\"removeBlackList\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"tokenByIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"tokenOfOwnerByIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_marketplaceStorage\",\"type\":\"address\"}],\"name\":\"updateStorageAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// NLGGTABI is the input ABI used to generate the binding from.
// Deprecated: Use NLGGTMetaData.ABI instead.
var NLGGTABI = NLGGTMetaData.ABI

// NLGGT is an auto generated Go binding around an Ethereum contract.
type NLGGT struct {
	NLGGTCaller     // Read-only binding to the contract
	NLGGTTransactor // Write-only binding to the contract
	NLGGTFilterer   // Log filterer for contract events
}

// NLGGTCaller is an auto generated read-only Go binding around an Ethereum contract.
type NLGGTCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NLGGTTransactor is an auto generated write-only Go binding around an Ethereum contract.
type NLGGTTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NLGGTFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type NLGGTFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NLGGTSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type NLGGTSession struct {
	Contract     *NLGGT            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// NLGGTCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type NLGGTCallerSession struct {
	Contract *NLGGTCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// NLGGTTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type NLGGTTransactorSession struct {
	Contract     *NLGGTTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// NLGGTRaw is an auto generated low-level Go binding around an Ethereum contract.
type NLGGTRaw struct {
	Contract *NLGGT // Generic contract binding to access the raw methods on
}

// NLGGTCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type NLGGTCallerRaw struct {
	Contract *NLGGTCaller // Generic read-only contract binding to access the raw methods on
}

// NLGGTTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type NLGGTTransactorRaw struct {
	Contract *NLGGTTransactor // Generic write-only contract binding to access the raw methods on
}

// NewNLGGT creates a new instance of NLGGT, bound to a specific deployed contract.
func NewNLGGT(address common.Address, backend bind.ContractBackend) (*NLGGT, error) {
	contract, err := bindNLGGT(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &NLGGT{NLGGTCaller: NLGGTCaller{contract: contract}, NLGGTTransactor: NLGGTTransactor{contract: contract}, NLGGTFilterer: NLGGTFilterer{contract: contract}}, nil
}

// NewNLGGTCaller creates a new read-only instance of NLGGT, bound to a specific deployed contract.
func NewNLGGTCaller(address common.Address, caller bind.ContractCaller) (*NLGGTCaller, error) {
	contract, err := bindNLGGT(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &NLGGTCaller{contract: contract}, nil
}

// NewNLGGTTransactor creates a new write-only instance of NLGGT, bound to a specific deployed contract.
func NewNLGGTTransactor(address common.Address, transactor bind.ContractTransactor) (*NLGGTTransactor, error) {
	contract, err := bindNLGGT(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &NLGGTTransactor{contract: contract}, nil
}

// NewNLGGTFilterer creates a new log filterer instance of NLGGT, bound to a specific deployed contract.
func NewNLGGTFilterer(address common.Address, filterer bind.ContractFilterer) (*NLGGTFilterer, error) {
	contract, err := bindNLGGT(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &NLGGTFilterer{contract: contract}, nil
}

// bindNLGGT binds a generic wrapper to an already deployed contract.
func bindNLGGT(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(NLGGTABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NLGGT *NLGGTRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _NLGGT.Contract.NLGGTCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NLGGT *NLGGTRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NLGGT.Contract.NLGGTTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NLGGT *NLGGTRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NLGGT.Contract.NLGGTTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NLGGT *NLGGTCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _NLGGT.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NLGGT *NLGGTTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NLGGT.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NLGGT *NLGGTTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NLGGT.Contract.contract.Transact(opts, method, params...)
}

// ADMINROLE is a free data retrieval call binding the contract method 0x75b238fc.
//
// Solidity: function ADMIN_ROLE() view returns(bytes32)
func (_NLGGT *NLGGTCaller) ADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _NLGGT.contract.Call(opts, &out, "ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ADMINROLE is a free data retrieval call binding the contract method 0x75b238fc.
//
// Solidity: function ADMIN_ROLE() view returns(bytes32)
func (_NLGGT *NLGGTSession) ADMINROLE() ([32]byte, error) {
	return _NLGGT.Contract.ADMINROLE(&_NLGGT.CallOpts)
}

// ADMINROLE is a free data retrieval call binding the contract method 0x75b238fc.
//
// Solidity: function ADMIN_ROLE() view returns(bytes32)
func (_NLGGT *NLGGTCallerSession) ADMINROLE() ([32]byte, error) {
	return _NLGGT.Contract.ADMINROLE(&_NLGGT.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_NLGGT *NLGGTCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _NLGGT.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_NLGGT *NLGGTSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _NLGGT.Contract.DEFAULTADMINROLE(&_NLGGT.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_NLGGT *NLGGTCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _NLGGT.Contract.DEFAULTADMINROLE(&_NLGGT.CallOpts)
}

// IMarketplaceStorageInterface is a free data retrieval call binding the contract method 0x8bafe559.
//
// Solidity: function IMarketplaceStorage_Interface() view returns(bytes4)
func (_NLGGT *NLGGTCaller) IMarketplaceStorageInterface(opts *bind.CallOpts) ([4]byte, error) {
	var out []interface{}
	err := _NLGGT.contract.Call(opts, &out, "IMarketplaceStorage_Interface")

	if err != nil {
		return *new([4]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)

	return out0, err

}

// IMarketplaceStorageInterface is a free data retrieval call binding the contract method 0x8bafe559.
//
// Solidity: function IMarketplaceStorage_Interface() view returns(bytes4)
func (_NLGGT *NLGGTSession) IMarketplaceStorageInterface() ([4]byte, error) {
	return _NLGGT.Contract.IMarketplaceStorageInterface(&_NLGGT.CallOpts)
}

// IMarketplaceStorageInterface is a free data retrieval call binding the contract method 0x8bafe559.
//
// Solidity: function IMarketplaceStorage_Interface() view returns(bytes4)
func (_NLGGT *NLGGTCallerSession) IMarketplaceStorageInterface() ([4]byte, error) {
	return _NLGGT.Contract.IMarketplaceStorageInterface(&_NLGGT.CallOpts)
}

// MINTERROLE is a free data retrieval call binding the contract method 0xd5391393.
//
// Solidity: function MINTER_ROLE() view returns(bytes32)
func (_NLGGT *NLGGTCaller) MINTERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _NLGGT.contract.Call(opts, &out, "MINTER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// MINTERROLE is a free data retrieval call binding the contract method 0xd5391393.
//
// Solidity: function MINTER_ROLE() view returns(bytes32)
func (_NLGGT *NLGGTSession) MINTERROLE() ([32]byte, error) {
	return _NLGGT.Contract.MINTERROLE(&_NLGGT.CallOpts)
}

// MINTERROLE is a free data retrieval call binding the contract method 0xd5391393.
//
// Solidity: function MINTER_ROLE() view returns(bytes32)
func (_NLGGT *NLGGTCallerSession) MINTERROLE() ([32]byte, error) {
	return _NLGGT.Contract.MINTERROLE(&_NLGGT.CallOpts)
}

// PAUSERROLE is a free data retrieval call binding the contract method 0xe63ab1e9.
//
// Solidity: function PAUSER_ROLE() view returns(bytes32)
func (_NLGGT *NLGGTCaller) PAUSERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _NLGGT.contract.Call(opts, &out, "PAUSER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// PAUSERROLE is a free data retrieval call binding the contract method 0xe63ab1e9.
//
// Solidity: function PAUSER_ROLE() view returns(bytes32)
func (_NLGGT *NLGGTSession) PAUSERROLE() ([32]byte, error) {
	return _NLGGT.Contract.PAUSERROLE(&_NLGGT.CallOpts)
}

// PAUSERROLE is a free data retrieval call binding the contract method 0xe63ab1e9.
//
// Solidity: function PAUSER_ROLE() view returns(bytes32)
func (_NLGGT *NLGGTCallerSession) PAUSERROLE() ([32]byte, error) {
	return _NLGGT.Contract.PAUSERROLE(&_NLGGT.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_NLGGT *NLGGTCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _NLGGT.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_NLGGT *NLGGTSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _NLGGT.Contract.BalanceOf(&_NLGGT.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_NLGGT *NLGGTCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _NLGGT.Contract.BalanceOf(&_NLGGT.CallOpts, owner)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_NLGGT *NLGGTCaller) GetApproved(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _NLGGT.contract.Call(opts, &out, "getApproved", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_NLGGT *NLGGTSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _NLGGT.Contract.GetApproved(&_NLGGT.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_NLGGT *NLGGTCallerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _NLGGT.Contract.GetApproved(&_NLGGT.CallOpts, tokenId)
}

// GetBlackListStatus is a free data retrieval call binding the contract method 0x59bf1abe.
//
// Solidity: function getBlackListStatus(address _maker) view returns(bool)
func (_NLGGT *NLGGTCaller) GetBlackListStatus(opts *bind.CallOpts, _maker common.Address) (bool, error) {
	var out []interface{}
	err := _NLGGT.contract.Call(opts, &out, "getBlackListStatus", _maker)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// GetBlackListStatus is a free data retrieval call binding the contract method 0x59bf1abe.
//
// Solidity: function getBlackListStatus(address _maker) view returns(bool)
func (_NLGGT *NLGGTSession) GetBlackListStatus(_maker common.Address) (bool, error) {
	return _NLGGT.Contract.GetBlackListStatus(&_NLGGT.CallOpts, _maker)
}

// GetBlackListStatus is a free data retrieval call binding the contract method 0x59bf1abe.
//
// Solidity: function getBlackListStatus(address _maker) view returns(bool)
func (_NLGGT *NLGGTCallerSession) GetBlackListStatus(_maker common.Address) (bool, error) {
	return _NLGGT.Contract.GetBlackListStatus(&_NLGGT.CallOpts, _maker)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_NLGGT *NLGGTCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _NLGGT.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_NLGGT *NLGGTSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _NLGGT.Contract.GetRoleAdmin(&_NLGGT.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_NLGGT *NLGGTCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _NLGGT.Contract.GetRoleAdmin(&_NLGGT.CallOpts, role)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_NLGGT *NLGGTCaller) GetRoleMember(opts *bind.CallOpts, role [32]byte, index *big.Int) (common.Address, error) {
	var out []interface{}
	err := _NLGGT.contract.Call(opts, &out, "getRoleMember", role, index)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_NLGGT *NLGGTSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _NLGGT.Contract.GetRoleMember(&_NLGGT.CallOpts, role, index)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_NLGGT *NLGGTCallerSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _NLGGT.Contract.GetRoleMember(&_NLGGT.CallOpts, role, index)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_NLGGT *NLGGTCaller) GetRoleMemberCount(opts *bind.CallOpts, role [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _NLGGT.contract.Call(opts, &out, "getRoleMemberCount", role)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_NLGGT *NLGGTSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _NLGGT.Contract.GetRoleMemberCount(&_NLGGT.CallOpts, role)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_NLGGT *NLGGTCallerSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _NLGGT.Contract.GetRoleMemberCount(&_NLGGT.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_NLGGT *NLGGTCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _NLGGT.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_NLGGT *NLGGTSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _NLGGT.Contract.HasRole(&_NLGGT.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_NLGGT *NLGGTCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _NLGGT.Contract.HasRole(&_NLGGT.CallOpts, role, account)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_NLGGT *NLGGTCaller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _NLGGT.contract.Call(opts, &out, "isApprovedForAll", owner, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_NLGGT *NLGGTSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _NLGGT.Contract.IsApprovedForAll(&_NLGGT.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_NLGGT *NLGGTCallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _NLGGT.Contract.IsApprovedForAll(&_NLGGT.CallOpts, owner, operator)
}

// IsBlackListed is a free data retrieval call binding the contract method 0xe47d6060.
//
// Solidity: function isBlackListed(address ) view returns(bool)
func (_NLGGT *NLGGTCaller) IsBlackListed(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _NLGGT.contract.Call(opts, &out, "isBlackListed", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsBlackListed is a free data retrieval call binding the contract method 0xe47d6060.
//
// Solidity: function isBlackListed(address ) view returns(bool)
func (_NLGGT *NLGGTSession) IsBlackListed(arg0 common.Address) (bool, error) {
	return _NLGGT.Contract.IsBlackListed(&_NLGGT.CallOpts, arg0)
}

// IsBlackListed is a free data retrieval call binding the contract method 0xe47d6060.
//
// Solidity: function isBlackListed(address ) view returns(bool)
func (_NLGGT *NLGGTCallerSession) IsBlackListed(arg0 common.Address) (bool, error) {
	return _NLGGT.Contract.IsBlackListed(&_NLGGT.CallOpts, arg0)
}

// MarketplaceStorage is a free data retrieval call binding the contract method 0x46b3aec6.
//
// Solidity: function marketplaceStorage() view returns(address)
func (_NLGGT *NLGGTCaller) MarketplaceStorage(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _NLGGT.contract.Call(opts, &out, "marketplaceStorage")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// MarketplaceStorage is a free data retrieval call binding the contract method 0x46b3aec6.
//
// Solidity: function marketplaceStorage() view returns(address)
func (_NLGGT *NLGGTSession) MarketplaceStorage() (common.Address, error) {
	return _NLGGT.Contract.MarketplaceStorage(&_NLGGT.CallOpts)
}

// MarketplaceStorage is a free data retrieval call binding the contract method 0x46b3aec6.
//
// Solidity: function marketplaceStorage() view returns(address)
func (_NLGGT *NLGGTCallerSession) MarketplaceStorage() (common.Address, error) {
	return _NLGGT.Contract.MarketplaceStorage(&_NLGGT.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_NLGGT *NLGGTCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _NLGGT.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_NLGGT *NLGGTSession) Name() (string, error) {
	return _NLGGT.Contract.Name(&_NLGGT.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_NLGGT *NLGGTCallerSession) Name() (string, error) {
	return _NLGGT.Contract.Name(&_NLGGT.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_NLGGT *NLGGTCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _NLGGT.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_NLGGT *NLGGTSession) Owner() (common.Address, error) {
	return _NLGGT.Contract.Owner(&_NLGGT.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_NLGGT *NLGGTCallerSession) Owner() (common.Address, error) {
	return _NLGGT.Contract.Owner(&_NLGGT.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_NLGGT *NLGGTCaller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _NLGGT.contract.Call(opts, &out, "ownerOf", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_NLGGT *NLGGTSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _NLGGT.Contract.OwnerOf(&_NLGGT.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_NLGGT *NLGGTCallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _NLGGT.Contract.OwnerOf(&_NLGGT.CallOpts, tokenId)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_NLGGT *NLGGTCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _NLGGT.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_NLGGT *NLGGTSession) Paused() (bool, error) {
	return _NLGGT.Contract.Paused(&_NLGGT.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_NLGGT *NLGGTCallerSession) Paused() (bool, error) {
	return _NLGGT.Contract.Paused(&_NLGGT.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_NLGGT *NLGGTCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _NLGGT.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_NLGGT *NLGGTSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _NLGGT.Contract.SupportsInterface(&_NLGGT.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_NLGGT *NLGGTCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _NLGGT.Contract.SupportsInterface(&_NLGGT.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_NLGGT *NLGGTCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _NLGGT.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_NLGGT *NLGGTSession) Symbol() (string, error) {
	return _NLGGT.Contract.Symbol(&_NLGGT.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_NLGGT *NLGGTCallerSession) Symbol() (string, error) {
	return _NLGGT.Contract.Symbol(&_NLGGT.CallOpts)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_NLGGT *NLGGTCaller) TokenByIndex(opts *bind.CallOpts, index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _NLGGT.contract.Call(opts, &out, "tokenByIndex", index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_NLGGT *NLGGTSession) TokenByIndex(index *big.Int) (*big.Int, error) {
	return _NLGGT.Contract.TokenByIndex(&_NLGGT.CallOpts, index)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_NLGGT *NLGGTCallerSession) TokenByIndex(index *big.Int) (*big.Int, error) {
	return _NLGGT.Contract.TokenByIndex(&_NLGGT.CallOpts, index)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_NLGGT *NLGGTCaller) TokenOfOwnerByIndex(opts *bind.CallOpts, owner common.Address, index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _NLGGT.contract.Call(opts, &out, "tokenOfOwnerByIndex", owner, index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_NLGGT *NLGGTSession) TokenOfOwnerByIndex(owner common.Address, index *big.Int) (*big.Int, error) {
	return _NLGGT.Contract.TokenOfOwnerByIndex(&_NLGGT.CallOpts, owner, index)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_NLGGT *NLGGTCallerSession) TokenOfOwnerByIndex(owner common.Address, index *big.Int) (*big.Int, error) {
	return _NLGGT.Contract.TokenOfOwnerByIndex(&_NLGGT.CallOpts, owner, index)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_NLGGT *NLGGTCaller) TokenURI(opts *bind.CallOpts, tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _NLGGT.contract.Call(opts, &out, "tokenURI", tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_NLGGT *NLGGTSession) TokenURI(tokenId *big.Int) (string, error) {
	return _NLGGT.Contract.TokenURI(&_NLGGT.CallOpts, tokenId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_NLGGT *NLGGTCallerSession) TokenURI(tokenId *big.Int) (string, error) {
	return _NLGGT.Contract.TokenURI(&_NLGGT.CallOpts, tokenId)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_NLGGT *NLGGTCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NLGGT.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_NLGGT *NLGGTSession) TotalSupply() (*big.Int, error) {
	return _NLGGT.Contract.TotalSupply(&_NLGGT.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_NLGGT *NLGGTCallerSession) TotalSupply() (*big.Int, error) {
	return _NLGGT.Contract.TotalSupply(&_NLGGT.CallOpts)
}

// AddBlackList is a paid mutator transaction binding the contract method 0x0ecb93c0.
//
// Solidity: function addBlackList(address _evilUser) returns()
func (_NLGGT *NLGGTTransactor) AddBlackList(opts *bind.TransactOpts, _evilUser common.Address) (*types.Transaction, error) {
	return _NLGGT.contract.Transact(opts, "addBlackList", _evilUser)
}

// AddBlackList is a paid mutator transaction binding the contract method 0x0ecb93c0.
//
// Solidity: function addBlackList(address _evilUser) returns()
func (_NLGGT *NLGGTSession) AddBlackList(_evilUser common.Address) (*types.Transaction, error) {
	return _NLGGT.Contract.AddBlackList(&_NLGGT.TransactOpts, _evilUser)
}

// AddBlackList is a paid mutator transaction binding the contract method 0x0ecb93c0.
//
// Solidity: function addBlackList(address _evilUser) returns()
func (_NLGGT *NLGGTTransactorSession) AddBlackList(_evilUser common.Address) (*types.Transaction, error) {
	return _NLGGT.Contract.AddBlackList(&_NLGGT.TransactOpts, _evilUser)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_NLGGT *NLGGTTransactor) Approve(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _NLGGT.contract.Transact(opts, "approve", to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_NLGGT *NLGGTSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _NLGGT.Contract.Approve(&_NLGGT.TransactOpts, to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_NLGGT *NLGGTTransactorSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _NLGGT.Contract.Approve(&_NLGGT.TransactOpts, to, tokenId)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 tokenId) returns()
func (_NLGGT *NLGGTTransactor) Burn(opts *bind.TransactOpts, tokenId *big.Int) (*types.Transaction, error) {
	return _NLGGT.contract.Transact(opts, "burn", tokenId)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 tokenId) returns()
func (_NLGGT *NLGGTSession) Burn(tokenId *big.Int) (*types.Transaction, error) {
	return _NLGGT.Contract.Burn(&_NLGGT.TransactOpts, tokenId)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 tokenId) returns()
func (_NLGGT *NLGGTTransactorSession) Burn(tokenId *big.Int) (*types.Transaction, error) {
	return _NLGGT.Contract.Burn(&_NLGGT.TransactOpts, tokenId)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_NLGGT *NLGGTTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _NLGGT.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_NLGGT *NLGGTSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _NLGGT.Contract.GrantRole(&_NLGGT.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_NLGGT *NLGGTTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _NLGGT.Contract.GrantRole(&_NLGGT.TransactOpts, role, account)
}

// Mint is a paid mutator transaction binding the contract method 0xd3fc9864.
//
// Solidity: function mint(address to, uint256 tokenId, string gtokenURI) returns()
func (_NLGGT *NLGGTTransactor) Mint(opts *bind.TransactOpts, to common.Address, tokenId *big.Int, gtokenURI string) (*types.Transaction, error) {
	return _NLGGT.contract.Transact(opts, "mint", to, tokenId, gtokenURI)
}

// Mint is a paid mutator transaction binding the contract method 0xd3fc9864.
//
// Solidity: function mint(address to, uint256 tokenId, string gtokenURI) returns()
func (_NLGGT *NLGGTSession) Mint(to common.Address, tokenId *big.Int, gtokenURI string) (*types.Transaction, error) {
	return _NLGGT.Contract.Mint(&_NLGGT.TransactOpts, to, tokenId, gtokenURI)
}

// Mint is a paid mutator transaction binding the contract method 0xd3fc9864.
//
// Solidity: function mint(address to, uint256 tokenId, string gtokenURI) returns()
func (_NLGGT *NLGGTTransactorSession) Mint(to common.Address, tokenId *big.Int, gtokenURI string) (*types.Transaction, error) {
	return _NLGGT.Contract.Mint(&_NLGGT.TransactOpts, to, tokenId, gtokenURI)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_NLGGT *NLGGTTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NLGGT.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_NLGGT *NLGGTSession) Pause() (*types.Transaction, error) {
	return _NLGGT.Contract.Pause(&_NLGGT.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_NLGGT *NLGGTTransactorSession) Pause() (*types.Transaction, error) {
	return _NLGGT.Contract.Pause(&_NLGGT.TransactOpts)
}

// RemoveBlackList is a paid mutator transaction binding the contract method 0xe4997dc5.
//
// Solidity: function removeBlackList(address _clearedUser) returns()
func (_NLGGT *NLGGTTransactor) RemoveBlackList(opts *bind.TransactOpts, _clearedUser common.Address) (*types.Transaction, error) {
	return _NLGGT.contract.Transact(opts, "removeBlackList", _clearedUser)
}

// RemoveBlackList is a paid mutator transaction binding the contract method 0xe4997dc5.
//
// Solidity: function removeBlackList(address _clearedUser) returns()
func (_NLGGT *NLGGTSession) RemoveBlackList(_clearedUser common.Address) (*types.Transaction, error) {
	return _NLGGT.Contract.RemoveBlackList(&_NLGGT.TransactOpts, _clearedUser)
}

// RemoveBlackList is a paid mutator transaction binding the contract method 0xe4997dc5.
//
// Solidity: function removeBlackList(address _clearedUser) returns()
func (_NLGGT *NLGGTTransactorSession) RemoveBlackList(_clearedUser common.Address) (*types.Transaction, error) {
	return _NLGGT.Contract.RemoveBlackList(&_NLGGT.TransactOpts, _clearedUser)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_NLGGT *NLGGTTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NLGGT.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_NLGGT *NLGGTSession) RenounceOwnership() (*types.Transaction, error) {
	return _NLGGT.Contract.RenounceOwnership(&_NLGGT.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_NLGGT *NLGGTTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _NLGGT.Contract.RenounceOwnership(&_NLGGT.TransactOpts)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_NLGGT *NLGGTTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _NLGGT.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_NLGGT *NLGGTSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _NLGGT.Contract.RenounceRole(&_NLGGT.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_NLGGT *NLGGTTransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _NLGGT.Contract.RenounceRole(&_NLGGT.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_NLGGT *NLGGTTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _NLGGT.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_NLGGT *NLGGTSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _NLGGT.Contract.RevokeRole(&_NLGGT.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_NLGGT *NLGGTTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _NLGGT.Contract.RevokeRole(&_NLGGT.TransactOpts, role, account)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_NLGGT *NLGGTTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _NLGGT.contract.Transact(opts, "safeTransferFrom", from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_NLGGT *NLGGTSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _NLGGT.Contract.SafeTransferFrom(&_NLGGT.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_NLGGT *NLGGTTransactorSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _NLGGT.Contract.SafeTransferFrom(&_NLGGT.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_NLGGT *NLGGTTransactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _NLGGT.contract.Transact(opts, "safeTransferFrom0", from, to, tokenId, _data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_NLGGT *NLGGTSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _NLGGT.Contract.SafeTransferFrom0(&_NLGGT.TransactOpts, from, to, tokenId, _data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_NLGGT *NLGGTTransactorSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _NLGGT.Contract.SafeTransferFrom0(&_NLGGT.TransactOpts, from, to, tokenId, _data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_NLGGT *NLGGTTransactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _NLGGT.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_NLGGT *NLGGTSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _NLGGT.Contract.SetApprovalForAll(&_NLGGT.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_NLGGT *NLGGTTransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _NLGGT.Contract.SetApprovalForAll(&_NLGGT.TransactOpts, operator, approved)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_NLGGT *NLGGTTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _NLGGT.contract.Transact(opts, "transferFrom", from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_NLGGT *NLGGTSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _NLGGT.Contract.TransferFrom(&_NLGGT.TransactOpts, from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_NLGGT *NLGGTTransactorSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _NLGGT.Contract.TransferFrom(&_NLGGT.TransactOpts, from, to, tokenId)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_NLGGT *NLGGTTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _NLGGT.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_NLGGT *NLGGTSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _NLGGT.Contract.TransferOwnership(&_NLGGT.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_NLGGT *NLGGTTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _NLGGT.Contract.TransferOwnership(&_NLGGT.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_NLGGT *NLGGTTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NLGGT.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_NLGGT *NLGGTSession) Unpause() (*types.Transaction, error) {
	return _NLGGT.Contract.Unpause(&_NLGGT.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_NLGGT *NLGGTTransactorSession) Unpause() (*types.Transaction, error) {
	return _NLGGT.Contract.Unpause(&_NLGGT.TransactOpts)
}

// UpdateStorageAddress is a paid mutator transaction binding the contract method 0x889e2129.
//
// Solidity: function updateStorageAddress(address _marketplaceStorage) returns()
func (_NLGGT *NLGGTTransactor) UpdateStorageAddress(opts *bind.TransactOpts, _marketplaceStorage common.Address) (*types.Transaction, error) {
	return _NLGGT.contract.Transact(opts, "updateStorageAddress", _marketplaceStorage)
}

// UpdateStorageAddress is a paid mutator transaction binding the contract method 0x889e2129.
//
// Solidity: function updateStorageAddress(address _marketplaceStorage) returns()
func (_NLGGT *NLGGTSession) UpdateStorageAddress(_marketplaceStorage common.Address) (*types.Transaction, error) {
	return _NLGGT.Contract.UpdateStorageAddress(&_NLGGT.TransactOpts, _marketplaceStorage)
}

// UpdateStorageAddress is a paid mutator transaction binding the contract method 0x889e2129.
//
// Solidity: function updateStorageAddress(address _marketplaceStorage) returns()
func (_NLGGT *NLGGTTransactorSession) UpdateStorageAddress(_marketplaceStorage common.Address) (*types.Transaction, error) {
	return _NLGGT.Contract.UpdateStorageAddress(&_NLGGT.TransactOpts, _marketplaceStorage)
}

// NLGGTAddedBlackListIterator is returned from FilterAddedBlackList and is used to iterate over the raw logs and unpacked data for AddedBlackList events raised by the NLGGT contract.
type NLGGTAddedBlackListIterator struct {
	Event *NLGGTAddedBlackList // Event containing the contract specifics and raw log

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
func (it *NLGGTAddedBlackListIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NLGGTAddedBlackList)
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
		it.Event = new(NLGGTAddedBlackList)
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
func (it *NLGGTAddedBlackListIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NLGGTAddedBlackListIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NLGGTAddedBlackList represents a AddedBlackList event raised by the NLGGT contract.
type NLGGTAddedBlackList struct {
	User common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterAddedBlackList is a free log retrieval operation binding the contract event 0x42e160154868087d6bfdc0ca23d96a1c1cfa32f1b72ba9ba27b69b98a0d819dc.
//
// Solidity: event AddedBlackList(address _user)
func (_NLGGT *NLGGTFilterer) FilterAddedBlackList(opts *bind.FilterOpts) (*NLGGTAddedBlackListIterator, error) {

	logs, sub, err := _NLGGT.contract.FilterLogs(opts, "AddedBlackList")
	if err != nil {
		return nil, err
	}
	return &NLGGTAddedBlackListIterator{contract: _NLGGT.contract, event: "AddedBlackList", logs: logs, sub: sub}, nil
}

// WatchAddedBlackList is a free log subscription operation binding the contract event 0x42e160154868087d6bfdc0ca23d96a1c1cfa32f1b72ba9ba27b69b98a0d819dc.
//
// Solidity: event AddedBlackList(address _user)
func (_NLGGT *NLGGTFilterer) WatchAddedBlackList(opts *bind.WatchOpts, sink chan<- *NLGGTAddedBlackList) (event.Subscription, error) {

	logs, sub, err := _NLGGT.contract.WatchLogs(opts, "AddedBlackList")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NLGGTAddedBlackList)
				if err := _NLGGT.contract.UnpackLog(event, "AddedBlackList", log); err != nil {
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
func (_NLGGT *NLGGTFilterer) ParseAddedBlackList(log types.Log) (*NLGGTAddedBlackList, error) {
	event := new(NLGGTAddedBlackList)
	if err := _NLGGT.contract.UnpackLog(event, "AddedBlackList", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NLGGTApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the NLGGT contract.
type NLGGTApprovalIterator struct {
	Event *NLGGTApproval // Event containing the contract specifics and raw log

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
func (it *NLGGTApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NLGGTApproval)
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
		it.Event = new(NLGGTApproval)
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
func (it *NLGGTApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NLGGTApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NLGGTApproval represents a Approval event raised by the NLGGT contract.
type NLGGTApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_NLGGT *NLGGTFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*NLGGTApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _NLGGT.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &NLGGTApprovalIterator{contract: _NLGGT.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_NLGGT *NLGGTFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *NLGGTApproval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _NLGGT.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NLGGTApproval)
				if err := _NLGGT.contract.UnpackLog(event, "Approval", log); err != nil {
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
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_NLGGT *NLGGTFilterer) ParseApproval(log types.Log) (*NLGGTApproval, error) {
	event := new(NLGGTApproval)
	if err := _NLGGT.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NLGGTApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the NLGGT contract.
type NLGGTApprovalForAllIterator struct {
	Event *NLGGTApprovalForAll // Event containing the contract specifics and raw log

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
func (it *NLGGTApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NLGGTApprovalForAll)
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
		it.Event = new(NLGGTApprovalForAll)
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
func (it *NLGGTApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NLGGTApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NLGGTApprovalForAll represents a ApprovalForAll event raised by the NLGGT contract.
type NLGGTApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_NLGGT *NLGGTFilterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*NLGGTApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _NLGGT.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &NLGGTApprovalForAllIterator{contract: _NLGGT.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_NLGGT *NLGGTFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *NLGGTApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _NLGGT.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NLGGTApprovalForAll)
				if err := _NLGGT.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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

// ParseApprovalForAll is a log parse operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_NLGGT *NLGGTFilterer) ParseApprovalForAll(log types.Log) (*NLGGTApprovalForAll, error) {
	event := new(NLGGTApprovalForAll)
	if err := _NLGGT.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NLGGTMarketplaceStorageUpdatedIterator is returned from FilterMarketplaceStorageUpdated and is used to iterate over the raw logs and unpacked data for MarketplaceStorageUpdated events raised by the NLGGT contract.
type NLGGTMarketplaceStorageUpdatedIterator struct {
	Event *NLGGTMarketplaceStorageUpdated // Event containing the contract specifics and raw log

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
func (it *NLGGTMarketplaceStorageUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NLGGTMarketplaceStorageUpdated)
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
		it.Event = new(NLGGTMarketplaceStorageUpdated)
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
func (it *NLGGTMarketplaceStorageUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NLGGTMarketplaceStorageUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NLGGTMarketplaceStorageUpdated represents a MarketplaceStorageUpdated event raised by the NLGGT contract.
type NLGGTMarketplaceStorageUpdated struct {
	MarketplaceStorage common.Address
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterMarketplaceStorageUpdated is a free log retrieval operation binding the contract event 0x19e07a962d0e1cab8b7d4be06b77a7225ca7ebb744994cb6546fdd815590dc9d.
//
// Solidity: event MarketplaceStorageUpdated(address _marketplaceStorage)
func (_NLGGT *NLGGTFilterer) FilterMarketplaceStorageUpdated(opts *bind.FilterOpts) (*NLGGTMarketplaceStorageUpdatedIterator, error) {

	logs, sub, err := _NLGGT.contract.FilterLogs(opts, "MarketplaceStorageUpdated")
	if err != nil {
		return nil, err
	}
	return &NLGGTMarketplaceStorageUpdatedIterator{contract: _NLGGT.contract, event: "MarketplaceStorageUpdated", logs: logs, sub: sub}, nil
}

// WatchMarketplaceStorageUpdated is a free log subscription operation binding the contract event 0x19e07a962d0e1cab8b7d4be06b77a7225ca7ebb744994cb6546fdd815590dc9d.
//
// Solidity: event MarketplaceStorageUpdated(address _marketplaceStorage)
func (_NLGGT *NLGGTFilterer) WatchMarketplaceStorageUpdated(opts *bind.WatchOpts, sink chan<- *NLGGTMarketplaceStorageUpdated) (event.Subscription, error) {

	logs, sub, err := _NLGGT.contract.WatchLogs(opts, "MarketplaceStorageUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NLGGTMarketplaceStorageUpdated)
				if err := _NLGGT.contract.UnpackLog(event, "MarketplaceStorageUpdated", log); err != nil {
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

// ParseMarketplaceStorageUpdated is a log parse operation binding the contract event 0x19e07a962d0e1cab8b7d4be06b77a7225ca7ebb744994cb6546fdd815590dc9d.
//
// Solidity: event MarketplaceStorageUpdated(address _marketplaceStorage)
func (_NLGGT *NLGGTFilterer) ParseMarketplaceStorageUpdated(log types.Log) (*NLGGTMarketplaceStorageUpdated, error) {
	event := new(NLGGTMarketplaceStorageUpdated)
	if err := _NLGGT.contract.UnpackLog(event, "MarketplaceStorageUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NLGGTOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the NLGGT contract.
type NLGGTOwnershipTransferredIterator struct {
	Event *NLGGTOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *NLGGTOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NLGGTOwnershipTransferred)
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
		it.Event = new(NLGGTOwnershipTransferred)
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
func (it *NLGGTOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NLGGTOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NLGGTOwnershipTransferred represents a OwnershipTransferred event raised by the NLGGT contract.
type NLGGTOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_NLGGT *NLGGTFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*NLGGTOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _NLGGT.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &NLGGTOwnershipTransferredIterator{contract: _NLGGT.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_NLGGT *NLGGTFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *NLGGTOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _NLGGT.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NLGGTOwnershipTransferred)
				if err := _NLGGT.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_NLGGT *NLGGTFilterer) ParseOwnershipTransferred(log types.Log) (*NLGGTOwnershipTransferred, error) {
	event := new(NLGGTOwnershipTransferred)
	if err := _NLGGT.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NLGGTPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the NLGGT contract.
type NLGGTPausedIterator struct {
	Event *NLGGTPaused // Event containing the contract specifics and raw log

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
func (it *NLGGTPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NLGGTPaused)
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
		it.Event = new(NLGGTPaused)
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
func (it *NLGGTPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NLGGTPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NLGGTPaused represents a Paused event raised by the NLGGT contract.
type NLGGTPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_NLGGT *NLGGTFilterer) FilterPaused(opts *bind.FilterOpts) (*NLGGTPausedIterator, error) {

	logs, sub, err := _NLGGT.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &NLGGTPausedIterator{contract: _NLGGT.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_NLGGT *NLGGTFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *NLGGTPaused) (event.Subscription, error) {

	logs, sub, err := _NLGGT.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NLGGTPaused)
				if err := _NLGGT.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_NLGGT *NLGGTFilterer) ParsePaused(log types.Log) (*NLGGTPaused, error) {
	event := new(NLGGTPaused)
	if err := _NLGGT.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NLGGTRemovedBlackListIterator is returned from FilterRemovedBlackList and is used to iterate over the raw logs and unpacked data for RemovedBlackList events raised by the NLGGT contract.
type NLGGTRemovedBlackListIterator struct {
	Event *NLGGTRemovedBlackList // Event containing the contract specifics and raw log

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
func (it *NLGGTRemovedBlackListIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NLGGTRemovedBlackList)
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
		it.Event = new(NLGGTRemovedBlackList)
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
func (it *NLGGTRemovedBlackListIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NLGGTRemovedBlackListIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NLGGTRemovedBlackList represents a RemovedBlackList event raised by the NLGGT contract.
type NLGGTRemovedBlackList struct {
	User common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterRemovedBlackList is a free log retrieval operation binding the contract event 0xd7e9ec6e6ecd65492dce6bf513cd6867560d49544421d0783ddf06e76c24470c.
//
// Solidity: event RemovedBlackList(address _user)
func (_NLGGT *NLGGTFilterer) FilterRemovedBlackList(opts *bind.FilterOpts) (*NLGGTRemovedBlackListIterator, error) {

	logs, sub, err := _NLGGT.contract.FilterLogs(opts, "RemovedBlackList")
	if err != nil {
		return nil, err
	}
	return &NLGGTRemovedBlackListIterator{contract: _NLGGT.contract, event: "RemovedBlackList", logs: logs, sub: sub}, nil
}

// WatchRemovedBlackList is a free log subscription operation binding the contract event 0xd7e9ec6e6ecd65492dce6bf513cd6867560d49544421d0783ddf06e76c24470c.
//
// Solidity: event RemovedBlackList(address _user)
func (_NLGGT *NLGGTFilterer) WatchRemovedBlackList(opts *bind.WatchOpts, sink chan<- *NLGGTRemovedBlackList) (event.Subscription, error) {

	logs, sub, err := _NLGGT.contract.WatchLogs(opts, "RemovedBlackList")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NLGGTRemovedBlackList)
				if err := _NLGGT.contract.UnpackLog(event, "RemovedBlackList", log); err != nil {
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
func (_NLGGT *NLGGTFilterer) ParseRemovedBlackList(log types.Log) (*NLGGTRemovedBlackList, error) {
	event := new(NLGGTRemovedBlackList)
	if err := _NLGGT.contract.UnpackLog(event, "RemovedBlackList", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NLGGTRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the NLGGT contract.
type NLGGTRoleAdminChangedIterator struct {
	Event *NLGGTRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *NLGGTRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NLGGTRoleAdminChanged)
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
		it.Event = new(NLGGTRoleAdminChanged)
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
func (it *NLGGTRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NLGGTRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NLGGTRoleAdminChanged represents a RoleAdminChanged event raised by the NLGGT contract.
type NLGGTRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_NLGGT *NLGGTFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*NLGGTRoleAdminChangedIterator, error) {

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

	logs, sub, err := _NLGGT.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &NLGGTRoleAdminChangedIterator{contract: _NLGGT.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_NLGGT *NLGGTFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *NLGGTRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _NLGGT.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NLGGTRoleAdminChanged)
				if err := _NLGGT.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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
func (_NLGGT *NLGGTFilterer) ParseRoleAdminChanged(log types.Log) (*NLGGTRoleAdminChanged, error) {
	event := new(NLGGTRoleAdminChanged)
	if err := _NLGGT.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NLGGTRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the NLGGT contract.
type NLGGTRoleGrantedIterator struct {
	Event *NLGGTRoleGranted // Event containing the contract specifics and raw log

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
func (it *NLGGTRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NLGGTRoleGranted)
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
		it.Event = new(NLGGTRoleGranted)
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
func (it *NLGGTRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NLGGTRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NLGGTRoleGranted represents a RoleGranted event raised by the NLGGT contract.
type NLGGTRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_NLGGT *NLGGTFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*NLGGTRoleGrantedIterator, error) {

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

	logs, sub, err := _NLGGT.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &NLGGTRoleGrantedIterator{contract: _NLGGT.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_NLGGT *NLGGTFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *NLGGTRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _NLGGT.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NLGGTRoleGranted)
				if err := _NLGGT.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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
func (_NLGGT *NLGGTFilterer) ParseRoleGranted(log types.Log) (*NLGGTRoleGranted, error) {
	event := new(NLGGTRoleGranted)
	if err := _NLGGT.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NLGGTRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the NLGGT contract.
type NLGGTRoleRevokedIterator struct {
	Event *NLGGTRoleRevoked // Event containing the contract specifics and raw log

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
func (it *NLGGTRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NLGGTRoleRevoked)
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
		it.Event = new(NLGGTRoleRevoked)
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
func (it *NLGGTRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NLGGTRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NLGGTRoleRevoked represents a RoleRevoked event raised by the NLGGT contract.
type NLGGTRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_NLGGT *NLGGTFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*NLGGTRoleRevokedIterator, error) {

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

	logs, sub, err := _NLGGT.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &NLGGTRoleRevokedIterator{contract: _NLGGT.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_NLGGT *NLGGTFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *NLGGTRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _NLGGT.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NLGGTRoleRevoked)
				if err := _NLGGT.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
func (_NLGGT *NLGGTFilterer) ParseRoleRevoked(log types.Log) (*NLGGTRoleRevoked, error) {
	event := new(NLGGTRoleRevoked)
	if err := _NLGGT.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NLGGTTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the NLGGT contract.
type NLGGTTransferIterator struct {
	Event *NLGGTTransfer // Event containing the contract specifics and raw log

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
func (it *NLGGTTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NLGGTTransfer)
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
		it.Event = new(NLGGTTransfer)
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
func (it *NLGGTTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NLGGTTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NLGGTTransfer represents a Transfer event raised by the NLGGT contract.
type NLGGTTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_NLGGT *NLGGTFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*NLGGTTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _NLGGT.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &NLGGTTransferIterator{contract: _NLGGT.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_NLGGT *NLGGTFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *NLGGTTransfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _NLGGT.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NLGGTTransfer)
				if err := _NLGGT.contract.UnpackLog(event, "Transfer", log); err != nil {
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
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_NLGGT *NLGGTFilterer) ParseTransfer(log types.Log) (*NLGGTTransfer, error) {
	event := new(NLGGTTransfer)
	if err := _NLGGT.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NLGGTUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the NLGGT contract.
type NLGGTUnpausedIterator struct {
	Event *NLGGTUnpaused // Event containing the contract specifics and raw log

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
func (it *NLGGTUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NLGGTUnpaused)
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
		it.Event = new(NLGGTUnpaused)
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
func (it *NLGGTUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NLGGTUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NLGGTUnpaused represents a Unpaused event raised by the NLGGT contract.
type NLGGTUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_NLGGT *NLGGTFilterer) FilterUnpaused(opts *bind.FilterOpts) (*NLGGTUnpausedIterator, error) {

	logs, sub, err := _NLGGT.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &NLGGTUnpausedIterator{contract: _NLGGT.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_NLGGT *NLGGTFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *NLGGTUnpaused) (event.Subscription, error) {

	logs, sub, err := _NLGGT.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NLGGTUnpaused)
				if err := _NLGGT.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
func (_NLGGT *NLGGTFilterer) ParseUnpaused(log types.Log) (*NLGGTUnpaused, error) {
	event := new(NLGGTUnpaused)
	if err := _NLGGT.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
