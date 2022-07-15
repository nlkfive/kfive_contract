// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package VRFCoordinatorV2Interface

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

// VRFCoordinatorV2InterfaceMetaData contains all meta data concerning the VRFCoordinatorV2Interface contract.
var VRFCoordinatorV2InterfaceMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"getRequestConfig\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"},{\"internalType\":\"bytes32[]\",\"name\":\"\",\"type\":\"bytes32[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"keyHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint64\",\"name\":\"subId\",\"type\":\"uint64\"},{\"internalType\":\"uint16\",\"name\":\"minimumRequestConfirmations\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"callbackGasLimit\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"numWords\",\"type\":\"uint32\"}],\"name\":\"requestRandomWords\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"requestId\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"createSubscription\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"subId\",\"type\":\"uint64\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"subId\",\"type\":\"uint64\"}],\"name\":\"getSubscription\",\"outputs\":[{\"internalType\":\"uint96\",\"name\":\"balance\",\"type\":\"uint96\"},{\"internalType\":\"uint64\",\"name\":\"reqCount\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"consumers\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"subId\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"requestSubscriptionOwnerTransfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"subId\",\"type\":\"uint64\"}],\"name\":\"acceptSubscriptionOwnerTransfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"subId\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"consumer\",\"type\":\"address\"}],\"name\":\"addConsumer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"subId\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"consumer\",\"type\":\"address\"}],\"name\":\"removeConsumer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"subId\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"cancelSubscription\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// VRFCoordinatorV2InterfaceABI is the input ABI used to generate the binding from.
// Deprecated: Use VRFCoordinatorV2InterfaceMetaData.ABI instead.
var VRFCoordinatorV2InterfaceABI = VRFCoordinatorV2InterfaceMetaData.ABI

// VRFCoordinatorV2Interface is an auto generated Go binding around an Ethereum contract.
type VRFCoordinatorV2Interface struct {
	VRFCoordinatorV2InterfaceCaller     // Read-only binding to the contract
	VRFCoordinatorV2InterfaceTransactor // Write-only binding to the contract
	VRFCoordinatorV2InterfaceFilterer   // Log filterer for contract events
}

// VRFCoordinatorV2InterfaceCaller is an auto generated read-only Go binding around an Ethereum contract.
type VRFCoordinatorV2InterfaceCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VRFCoordinatorV2InterfaceTransactor is an auto generated write-only Go binding around an Ethereum contract.
type VRFCoordinatorV2InterfaceTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VRFCoordinatorV2InterfaceFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type VRFCoordinatorV2InterfaceFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VRFCoordinatorV2InterfaceSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type VRFCoordinatorV2InterfaceSession struct {
	Contract     *VRFCoordinatorV2Interface // Generic contract binding to set the session for
	CallOpts     bind.CallOpts              // Call options to use throughout this session
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// VRFCoordinatorV2InterfaceCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type VRFCoordinatorV2InterfaceCallerSession struct {
	Contract *VRFCoordinatorV2InterfaceCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                    // Call options to use throughout this session
}

// VRFCoordinatorV2InterfaceTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type VRFCoordinatorV2InterfaceTransactorSession struct {
	Contract     *VRFCoordinatorV2InterfaceTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                    // Transaction auth options to use throughout this session
}

// VRFCoordinatorV2InterfaceRaw is an auto generated low-level Go binding around an Ethereum contract.
type VRFCoordinatorV2InterfaceRaw struct {
	Contract *VRFCoordinatorV2Interface // Generic contract binding to access the raw methods on
}

// VRFCoordinatorV2InterfaceCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type VRFCoordinatorV2InterfaceCallerRaw struct {
	Contract *VRFCoordinatorV2InterfaceCaller // Generic read-only contract binding to access the raw methods on
}

// VRFCoordinatorV2InterfaceTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type VRFCoordinatorV2InterfaceTransactorRaw struct {
	Contract *VRFCoordinatorV2InterfaceTransactor // Generic write-only contract binding to access the raw methods on
}

// NewVRFCoordinatorV2Interface creates a new instance of VRFCoordinatorV2Interface, bound to a specific deployed contract.
func NewVRFCoordinatorV2Interface(address common.Address, backend bind.ContractBackend) (*VRFCoordinatorV2Interface, error) {
	contract, err := bindVRFCoordinatorV2Interface(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &VRFCoordinatorV2Interface{VRFCoordinatorV2InterfaceCaller: VRFCoordinatorV2InterfaceCaller{contract: contract}, VRFCoordinatorV2InterfaceTransactor: VRFCoordinatorV2InterfaceTransactor{contract: contract}, VRFCoordinatorV2InterfaceFilterer: VRFCoordinatorV2InterfaceFilterer{contract: contract}}, nil
}

// NewVRFCoordinatorV2InterfaceCaller creates a new read-only instance of VRFCoordinatorV2Interface, bound to a specific deployed contract.
func NewVRFCoordinatorV2InterfaceCaller(address common.Address, caller bind.ContractCaller) (*VRFCoordinatorV2InterfaceCaller, error) {
	contract, err := bindVRFCoordinatorV2Interface(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &VRFCoordinatorV2InterfaceCaller{contract: contract}, nil
}

// NewVRFCoordinatorV2InterfaceTransactor creates a new write-only instance of VRFCoordinatorV2Interface, bound to a specific deployed contract.
func NewVRFCoordinatorV2InterfaceTransactor(address common.Address, transactor bind.ContractTransactor) (*VRFCoordinatorV2InterfaceTransactor, error) {
	contract, err := bindVRFCoordinatorV2Interface(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &VRFCoordinatorV2InterfaceTransactor{contract: contract}, nil
}

// NewVRFCoordinatorV2InterfaceFilterer creates a new log filterer instance of VRFCoordinatorV2Interface, bound to a specific deployed contract.
func NewVRFCoordinatorV2InterfaceFilterer(address common.Address, filterer bind.ContractFilterer) (*VRFCoordinatorV2InterfaceFilterer, error) {
	contract, err := bindVRFCoordinatorV2Interface(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &VRFCoordinatorV2InterfaceFilterer{contract: contract}, nil
}

// bindVRFCoordinatorV2Interface binds a generic wrapper to an already deployed contract.
func bindVRFCoordinatorV2Interface(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(VRFCoordinatorV2InterfaceABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VRFCoordinatorV2Interface *VRFCoordinatorV2InterfaceRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VRFCoordinatorV2Interface.Contract.VRFCoordinatorV2InterfaceCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VRFCoordinatorV2Interface *VRFCoordinatorV2InterfaceRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VRFCoordinatorV2Interface.Contract.VRFCoordinatorV2InterfaceTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VRFCoordinatorV2Interface *VRFCoordinatorV2InterfaceRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VRFCoordinatorV2Interface.Contract.VRFCoordinatorV2InterfaceTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VRFCoordinatorV2Interface *VRFCoordinatorV2InterfaceCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VRFCoordinatorV2Interface.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VRFCoordinatorV2Interface *VRFCoordinatorV2InterfaceTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VRFCoordinatorV2Interface.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VRFCoordinatorV2Interface *VRFCoordinatorV2InterfaceTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VRFCoordinatorV2Interface.Contract.contract.Transact(opts, method, params...)
}

// GetRequestConfig is a free data retrieval call binding the contract method 0x00012291.
//
// Solidity: function getRequestConfig() view returns(uint16, uint32, bytes32[])
func (_VRFCoordinatorV2Interface *VRFCoordinatorV2InterfaceCaller) GetRequestConfig(opts *bind.CallOpts) (uint16, uint32, [][32]byte, error) {
	var out []interface{}
	err := _VRFCoordinatorV2Interface.contract.Call(opts, &out, "getRequestConfig")

	if err != nil {
		return *new(uint16), *new(uint32), *new([][32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)
	out1 := *abi.ConvertType(out[1], new(uint32)).(*uint32)
	out2 := *abi.ConvertType(out[2], new([][32]byte)).(*[][32]byte)

	return out0, out1, out2, err

}

// GetRequestConfig is a free data retrieval call binding the contract method 0x00012291.
//
// Solidity: function getRequestConfig() view returns(uint16, uint32, bytes32[])
func (_VRFCoordinatorV2Interface *VRFCoordinatorV2InterfaceSession) GetRequestConfig() (uint16, uint32, [][32]byte, error) {
	return _VRFCoordinatorV2Interface.Contract.GetRequestConfig(&_VRFCoordinatorV2Interface.CallOpts)
}

// GetRequestConfig is a free data retrieval call binding the contract method 0x00012291.
//
// Solidity: function getRequestConfig() view returns(uint16, uint32, bytes32[])
func (_VRFCoordinatorV2Interface *VRFCoordinatorV2InterfaceCallerSession) GetRequestConfig() (uint16, uint32, [][32]byte, error) {
	return _VRFCoordinatorV2Interface.Contract.GetRequestConfig(&_VRFCoordinatorV2Interface.CallOpts)
}

// GetSubscription is a free data retrieval call binding the contract method 0xa47c7696.
//
// Solidity: function getSubscription(uint64 subId) view returns(uint96 balance, uint64 reqCount, address owner, address[] consumers)
func (_VRFCoordinatorV2Interface *VRFCoordinatorV2InterfaceCaller) GetSubscription(opts *bind.CallOpts, subId uint64) (struct {
	Balance   *big.Int
	ReqCount  uint64
	Owner     common.Address
	Consumers []common.Address
}, error) {
	var out []interface{}
	err := _VRFCoordinatorV2Interface.contract.Call(opts, &out, "getSubscription", subId)

	outstruct := new(struct {
		Balance   *big.Int
		ReqCount  uint64
		Owner     common.Address
		Consumers []common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Balance = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.ReqCount = *abi.ConvertType(out[1], new(uint64)).(*uint64)
	outstruct.Owner = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)
	outstruct.Consumers = *abi.ConvertType(out[3], new([]common.Address)).(*[]common.Address)

	return *outstruct, err

}

// GetSubscription is a free data retrieval call binding the contract method 0xa47c7696.
//
// Solidity: function getSubscription(uint64 subId) view returns(uint96 balance, uint64 reqCount, address owner, address[] consumers)
func (_VRFCoordinatorV2Interface *VRFCoordinatorV2InterfaceSession) GetSubscription(subId uint64) (struct {
	Balance   *big.Int
	ReqCount  uint64
	Owner     common.Address
	Consumers []common.Address
}, error) {
	return _VRFCoordinatorV2Interface.Contract.GetSubscription(&_VRFCoordinatorV2Interface.CallOpts, subId)
}

// GetSubscription is a free data retrieval call binding the contract method 0xa47c7696.
//
// Solidity: function getSubscription(uint64 subId) view returns(uint96 balance, uint64 reqCount, address owner, address[] consumers)
func (_VRFCoordinatorV2Interface *VRFCoordinatorV2InterfaceCallerSession) GetSubscription(subId uint64) (struct {
	Balance   *big.Int
	ReqCount  uint64
	Owner     common.Address
	Consumers []common.Address
}, error) {
	return _VRFCoordinatorV2Interface.Contract.GetSubscription(&_VRFCoordinatorV2Interface.CallOpts, subId)
}

// AcceptSubscriptionOwnerTransfer is a paid mutator transaction binding the contract method 0x82359740.
//
// Solidity: function acceptSubscriptionOwnerTransfer(uint64 subId) returns()
func (_VRFCoordinatorV2Interface *VRFCoordinatorV2InterfaceTransactor) AcceptSubscriptionOwnerTransfer(opts *bind.TransactOpts, subId uint64) (*types.Transaction, error) {
	return _VRFCoordinatorV2Interface.contract.Transact(opts, "acceptSubscriptionOwnerTransfer", subId)
}

// AcceptSubscriptionOwnerTransfer is a paid mutator transaction binding the contract method 0x82359740.
//
// Solidity: function acceptSubscriptionOwnerTransfer(uint64 subId) returns()
func (_VRFCoordinatorV2Interface *VRFCoordinatorV2InterfaceSession) AcceptSubscriptionOwnerTransfer(subId uint64) (*types.Transaction, error) {
	return _VRFCoordinatorV2Interface.Contract.AcceptSubscriptionOwnerTransfer(&_VRFCoordinatorV2Interface.TransactOpts, subId)
}

// AcceptSubscriptionOwnerTransfer is a paid mutator transaction binding the contract method 0x82359740.
//
// Solidity: function acceptSubscriptionOwnerTransfer(uint64 subId) returns()
func (_VRFCoordinatorV2Interface *VRFCoordinatorV2InterfaceTransactorSession) AcceptSubscriptionOwnerTransfer(subId uint64) (*types.Transaction, error) {
	return _VRFCoordinatorV2Interface.Contract.AcceptSubscriptionOwnerTransfer(&_VRFCoordinatorV2Interface.TransactOpts, subId)
}

// AddConsumer is a paid mutator transaction binding the contract method 0x7341c10c.
//
// Solidity: function addConsumer(uint64 subId, address consumer) returns()
func (_VRFCoordinatorV2Interface *VRFCoordinatorV2InterfaceTransactor) AddConsumer(opts *bind.TransactOpts, subId uint64, consumer common.Address) (*types.Transaction, error) {
	return _VRFCoordinatorV2Interface.contract.Transact(opts, "addConsumer", subId, consumer)
}

// AddConsumer is a paid mutator transaction binding the contract method 0x7341c10c.
//
// Solidity: function addConsumer(uint64 subId, address consumer) returns()
func (_VRFCoordinatorV2Interface *VRFCoordinatorV2InterfaceSession) AddConsumer(subId uint64, consumer common.Address) (*types.Transaction, error) {
	return _VRFCoordinatorV2Interface.Contract.AddConsumer(&_VRFCoordinatorV2Interface.TransactOpts, subId, consumer)
}

// AddConsumer is a paid mutator transaction binding the contract method 0x7341c10c.
//
// Solidity: function addConsumer(uint64 subId, address consumer) returns()
func (_VRFCoordinatorV2Interface *VRFCoordinatorV2InterfaceTransactorSession) AddConsumer(subId uint64, consumer common.Address) (*types.Transaction, error) {
	return _VRFCoordinatorV2Interface.Contract.AddConsumer(&_VRFCoordinatorV2Interface.TransactOpts, subId, consumer)
}

// CancelSubscription is a paid mutator transaction binding the contract method 0xd7ae1d30.
//
// Solidity: function cancelSubscription(uint64 subId, address to) returns()
func (_VRFCoordinatorV2Interface *VRFCoordinatorV2InterfaceTransactor) CancelSubscription(opts *bind.TransactOpts, subId uint64, to common.Address) (*types.Transaction, error) {
	return _VRFCoordinatorV2Interface.contract.Transact(opts, "cancelSubscription", subId, to)
}

// CancelSubscription is a paid mutator transaction binding the contract method 0xd7ae1d30.
//
// Solidity: function cancelSubscription(uint64 subId, address to) returns()
func (_VRFCoordinatorV2Interface *VRFCoordinatorV2InterfaceSession) CancelSubscription(subId uint64, to common.Address) (*types.Transaction, error) {
	return _VRFCoordinatorV2Interface.Contract.CancelSubscription(&_VRFCoordinatorV2Interface.TransactOpts, subId, to)
}

// CancelSubscription is a paid mutator transaction binding the contract method 0xd7ae1d30.
//
// Solidity: function cancelSubscription(uint64 subId, address to) returns()
func (_VRFCoordinatorV2Interface *VRFCoordinatorV2InterfaceTransactorSession) CancelSubscription(subId uint64, to common.Address) (*types.Transaction, error) {
	return _VRFCoordinatorV2Interface.Contract.CancelSubscription(&_VRFCoordinatorV2Interface.TransactOpts, subId, to)
}

// CreateSubscription is a paid mutator transaction binding the contract method 0xa21a23e4.
//
// Solidity: function createSubscription() returns(uint64 subId)
func (_VRFCoordinatorV2Interface *VRFCoordinatorV2InterfaceTransactor) CreateSubscription(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VRFCoordinatorV2Interface.contract.Transact(opts, "createSubscription")
}

// CreateSubscription is a paid mutator transaction binding the contract method 0xa21a23e4.
//
// Solidity: function createSubscription() returns(uint64 subId)
func (_VRFCoordinatorV2Interface *VRFCoordinatorV2InterfaceSession) CreateSubscription() (*types.Transaction, error) {
	return _VRFCoordinatorV2Interface.Contract.CreateSubscription(&_VRFCoordinatorV2Interface.TransactOpts)
}

// CreateSubscription is a paid mutator transaction binding the contract method 0xa21a23e4.
//
// Solidity: function createSubscription() returns(uint64 subId)
func (_VRFCoordinatorV2Interface *VRFCoordinatorV2InterfaceTransactorSession) CreateSubscription() (*types.Transaction, error) {
	return _VRFCoordinatorV2Interface.Contract.CreateSubscription(&_VRFCoordinatorV2Interface.TransactOpts)
}

// RemoveConsumer is a paid mutator transaction binding the contract method 0x9f87fad7.
//
// Solidity: function removeConsumer(uint64 subId, address consumer) returns()
func (_VRFCoordinatorV2Interface *VRFCoordinatorV2InterfaceTransactor) RemoveConsumer(opts *bind.TransactOpts, subId uint64, consumer common.Address) (*types.Transaction, error) {
	return _VRFCoordinatorV2Interface.contract.Transact(opts, "removeConsumer", subId, consumer)
}

// RemoveConsumer is a paid mutator transaction binding the contract method 0x9f87fad7.
//
// Solidity: function removeConsumer(uint64 subId, address consumer) returns()
func (_VRFCoordinatorV2Interface *VRFCoordinatorV2InterfaceSession) RemoveConsumer(subId uint64, consumer common.Address) (*types.Transaction, error) {
	return _VRFCoordinatorV2Interface.Contract.RemoveConsumer(&_VRFCoordinatorV2Interface.TransactOpts, subId, consumer)
}

// RemoveConsumer is a paid mutator transaction binding the contract method 0x9f87fad7.
//
// Solidity: function removeConsumer(uint64 subId, address consumer) returns()
func (_VRFCoordinatorV2Interface *VRFCoordinatorV2InterfaceTransactorSession) RemoveConsumer(subId uint64, consumer common.Address) (*types.Transaction, error) {
	return _VRFCoordinatorV2Interface.Contract.RemoveConsumer(&_VRFCoordinatorV2Interface.TransactOpts, subId, consumer)
}

// RequestRandomWords is a paid mutator transaction binding the contract method 0x5d3b1d30.
//
// Solidity: function requestRandomWords(bytes32 keyHash, uint64 subId, uint16 minimumRequestConfirmations, uint32 callbackGasLimit, uint32 numWords) returns(uint256 requestId)
func (_VRFCoordinatorV2Interface *VRFCoordinatorV2InterfaceTransactor) RequestRandomWords(opts *bind.TransactOpts, keyHash [32]byte, subId uint64, minimumRequestConfirmations uint16, callbackGasLimit uint32, numWords uint32) (*types.Transaction, error) {
	return _VRFCoordinatorV2Interface.contract.Transact(opts, "requestRandomWords", keyHash, subId, minimumRequestConfirmations, callbackGasLimit, numWords)
}

// RequestRandomWords is a paid mutator transaction binding the contract method 0x5d3b1d30.
//
// Solidity: function requestRandomWords(bytes32 keyHash, uint64 subId, uint16 minimumRequestConfirmations, uint32 callbackGasLimit, uint32 numWords) returns(uint256 requestId)
func (_VRFCoordinatorV2Interface *VRFCoordinatorV2InterfaceSession) RequestRandomWords(keyHash [32]byte, subId uint64, minimumRequestConfirmations uint16, callbackGasLimit uint32, numWords uint32) (*types.Transaction, error) {
	return _VRFCoordinatorV2Interface.Contract.RequestRandomWords(&_VRFCoordinatorV2Interface.TransactOpts, keyHash, subId, minimumRequestConfirmations, callbackGasLimit, numWords)
}

// RequestRandomWords is a paid mutator transaction binding the contract method 0x5d3b1d30.
//
// Solidity: function requestRandomWords(bytes32 keyHash, uint64 subId, uint16 minimumRequestConfirmations, uint32 callbackGasLimit, uint32 numWords) returns(uint256 requestId)
func (_VRFCoordinatorV2Interface *VRFCoordinatorV2InterfaceTransactorSession) RequestRandomWords(keyHash [32]byte, subId uint64, minimumRequestConfirmations uint16, callbackGasLimit uint32, numWords uint32) (*types.Transaction, error) {
	return _VRFCoordinatorV2Interface.Contract.RequestRandomWords(&_VRFCoordinatorV2Interface.TransactOpts, keyHash, subId, minimumRequestConfirmations, callbackGasLimit, numWords)
}

// RequestSubscriptionOwnerTransfer is a paid mutator transaction binding the contract method 0x04c357cb.
//
// Solidity: function requestSubscriptionOwnerTransfer(uint64 subId, address newOwner) returns()
func (_VRFCoordinatorV2Interface *VRFCoordinatorV2InterfaceTransactor) RequestSubscriptionOwnerTransfer(opts *bind.TransactOpts, subId uint64, newOwner common.Address) (*types.Transaction, error) {
	return _VRFCoordinatorV2Interface.contract.Transact(opts, "requestSubscriptionOwnerTransfer", subId, newOwner)
}

// RequestSubscriptionOwnerTransfer is a paid mutator transaction binding the contract method 0x04c357cb.
//
// Solidity: function requestSubscriptionOwnerTransfer(uint64 subId, address newOwner) returns()
func (_VRFCoordinatorV2Interface *VRFCoordinatorV2InterfaceSession) RequestSubscriptionOwnerTransfer(subId uint64, newOwner common.Address) (*types.Transaction, error) {
	return _VRFCoordinatorV2Interface.Contract.RequestSubscriptionOwnerTransfer(&_VRFCoordinatorV2Interface.TransactOpts, subId, newOwner)
}

// RequestSubscriptionOwnerTransfer is a paid mutator transaction binding the contract method 0x04c357cb.
//
// Solidity: function requestSubscriptionOwnerTransfer(uint64 subId, address newOwner) returns()
func (_VRFCoordinatorV2Interface *VRFCoordinatorV2InterfaceTransactorSession) RequestSubscriptionOwnerTransfer(subId uint64, newOwner common.Address) (*types.Transaction, error) {
	return _VRFCoordinatorV2Interface.Contract.RequestSubscriptionOwnerTransfer(&_VRFCoordinatorV2Interface.TransactOpts, subId, newOwner)
}
