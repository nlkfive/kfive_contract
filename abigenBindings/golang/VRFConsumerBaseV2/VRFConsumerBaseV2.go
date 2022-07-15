// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package VRFConsumerBaseV2

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

// VRFConsumerBaseV2MetaData contains all meta data concerning the VRFConsumerBaseV2 contract.
var VRFConsumerBaseV2MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"have\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"want\",\"type\":\"address\"}],\"name\":\"OnlyCoordinatorCanFulfill\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"requestId\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"randomWords\",\"type\":\"uint256[]\"}],\"name\":\"rawFulfillRandomWords\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// VRFConsumerBaseV2ABI is the input ABI used to generate the binding from.
// Deprecated: Use VRFConsumerBaseV2MetaData.ABI instead.
var VRFConsumerBaseV2ABI = VRFConsumerBaseV2MetaData.ABI

// VRFConsumerBaseV2 is an auto generated Go binding around an Ethereum contract.
type VRFConsumerBaseV2 struct {
	VRFConsumerBaseV2Caller     // Read-only binding to the contract
	VRFConsumerBaseV2Transactor // Write-only binding to the contract
	VRFConsumerBaseV2Filterer   // Log filterer for contract events
}

// VRFConsumerBaseV2Caller is an auto generated read-only Go binding around an Ethereum contract.
type VRFConsumerBaseV2Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VRFConsumerBaseV2Transactor is an auto generated write-only Go binding around an Ethereum contract.
type VRFConsumerBaseV2Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VRFConsumerBaseV2Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type VRFConsumerBaseV2Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VRFConsumerBaseV2Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type VRFConsumerBaseV2Session struct {
	Contract     *VRFConsumerBaseV2 // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// VRFConsumerBaseV2CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type VRFConsumerBaseV2CallerSession struct {
	Contract *VRFConsumerBaseV2Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// VRFConsumerBaseV2TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type VRFConsumerBaseV2TransactorSession struct {
	Contract     *VRFConsumerBaseV2Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// VRFConsumerBaseV2Raw is an auto generated low-level Go binding around an Ethereum contract.
type VRFConsumerBaseV2Raw struct {
	Contract *VRFConsumerBaseV2 // Generic contract binding to access the raw methods on
}

// VRFConsumerBaseV2CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type VRFConsumerBaseV2CallerRaw struct {
	Contract *VRFConsumerBaseV2Caller // Generic read-only contract binding to access the raw methods on
}

// VRFConsumerBaseV2TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type VRFConsumerBaseV2TransactorRaw struct {
	Contract *VRFConsumerBaseV2Transactor // Generic write-only contract binding to access the raw methods on
}

// NewVRFConsumerBaseV2 creates a new instance of VRFConsumerBaseV2, bound to a specific deployed contract.
func NewVRFConsumerBaseV2(address common.Address, backend bind.ContractBackend) (*VRFConsumerBaseV2, error) {
	contract, err := bindVRFConsumerBaseV2(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &VRFConsumerBaseV2{VRFConsumerBaseV2Caller: VRFConsumerBaseV2Caller{contract: contract}, VRFConsumerBaseV2Transactor: VRFConsumerBaseV2Transactor{contract: contract}, VRFConsumerBaseV2Filterer: VRFConsumerBaseV2Filterer{contract: contract}}, nil
}

// NewVRFConsumerBaseV2Caller creates a new read-only instance of VRFConsumerBaseV2, bound to a specific deployed contract.
func NewVRFConsumerBaseV2Caller(address common.Address, caller bind.ContractCaller) (*VRFConsumerBaseV2Caller, error) {
	contract, err := bindVRFConsumerBaseV2(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &VRFConsumerBaseV2Caller{contract: contract}, nil
}

// NewVRFConsumerBaseV2Transactor creates a new write-only instance of VRFConsumerBaseV2, bound to a specific deployed contract.
func NewVRFConsumerBaseV2Transactor(address common.Address, transactor bind.ContractTransactor) (*VRFConsumerBaseV2Transactor, error) {
	contract, err := bindVRFConsumerBaseV2(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &VRFConsumerBaseV2Transactor{contract: contract}, nil
}

// NewVRFConsumerBaseV2Filterer creates a new log filterer instance of VRFConsumerBaseV2, bound to a specific deployed contract.
func NewVRFConsumerBaseV2Filterer(address common.Address, filterer bind.ContractFilterer) (*VRFConsumerBaseV2Filterer, error) {
	contract, err := bindVRFConsumerBaseV2(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &VRFConsumerBaseV2Filterer{contract: contract}, nil
}

// bindVRFConsumerBaseV2 binds a generic wrapper to an already deployed contract.
func bindVRFConsumerBaseV2(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(VRFConsumerBaseV2ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VRFConsumerBaseV2 *VRFConsumerBaseV2Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VRFConsumerBaseV2.Contract.VRFConsumerBaseV2Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VRFConsumerBaseV2 *VRFConsumerBaseV2Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VRFConsumerBaseV2.Contract.VRFConsumerBaseV2Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VRFConsumerBaseV2 *VRFConsumerBaseV2Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VRFConsumerBaseV2.Contract.VRFConsumerBaseV2Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VRFConsumerBaseV2 *VRFConsumerBaseV2CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VRFConsumerBaseV2.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VRFConsumerBaseV2 *VRFConsumerBaseV2TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VRFConsumerBaseV2.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VRFConsumerBaseV2 *VRFConsumerBaseV2TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VRFConsumerBaseV2.Contract.contract.Transact(opts, method, params...)
}

// RawFulfillRandomWords is a paid mutator transaction binding the contract method 0x1fe543e3.
//
// Solidity: function rawFulfillRandomWords(uint256 requestId, uint256[] randomWords) returns()
func (_VRFConsumerBaseV2 *VRFConsumerBaseV2Transactor) RawFulfillRandomWords(opts *bind.TransactOpts, requestId *big.Int, randomWords []*big.Int) (*types.Transaction, error) {
	return _VRFConsumerBaseV2.contract.Transact(opts, "rawFulfillRandomWords", requestId, randomWords)
}

// RawFulfillRandomWords is a paid mutator transaction binding the contract method 0x1fe543e3.
//
// Solidity: function rawFulfillRandomWords(uint256 requestId, uint256[] randomWords) returns()
func (_VRFConsumerBaseV2 *VRFConsumerBaseV2Session) RawFulfillRandomWords(requestId *big.Int, randomWords []*big.Int) (*types.Transaction, error) {
	return _VRFConsumerBaseV2.Contract.RawFulfillRandomWords(&_VRFConsumerBaseV2.TransactOpts, requestId, randomWords)
}

// RawFulfillRandomWords is a paid mutator transaction binding the contract method 0x1fe543e3.
//
// Solidity: function rawFulfillRandomWords(uint256 requestId, uint256[] randomWords) returns()
func (_VRFConsumerBaseV2 *VRFConsumerBaseV2TransactorSession) RawFulfillRandomWords(requestId *big.Int, randomWords []*big.Int) (*types.Transaction, error) {
	return _VRFConsumerBaseV2.Contract.RawFulfillRandomWords(&_VRFConsumerBaseV2.TransactOpts, requestId, randomWords)
}
