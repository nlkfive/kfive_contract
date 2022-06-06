// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package DivideFunction

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

// DivideFunctionMetaData contains all meta data concerning the DivideFunction contract.
var DivideFunctionMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"weiValue\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rate\",\"type\":\"uint256\"}],\"name\":\"weiToKfive\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b5061034a806100206000396000f3fe608060405234801561001057600080fd5b506004361061002b5760003560e01c8063d7a65f0b14610030575b600080fd5b61004a60048036038101906100459190610127565b610060565b60405161005791906101b5565b60405180910390f35b6000633b9aca008310156100a9576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016100a090610195565b60405180910390fd5b60006100c2633b9aca00856100e690919063ffffffff16565b905060006100d984836100fc90919063ffffffff16565b9050809250505092915050565b600081836100f491906101e1565b905092915050565b6000818361010a9190610212565b905092915050565b600081359050610121816102fd565b92915050565b6000806040838503121561013a57600080fd5b600061014885828601610112565b925050602061015985828601610112565b9150509250929050565b6000610170601e836101d0565b915061017b826102d4565b602082019050919050565b61018f8161026c565b82525050565b600060208201905081810360008301526101ae81610163565b9050919050565b60006020820190506101ca6000830184610186565b92915050565b600082825260208201905092915050565b60006101ec8261026c565b91506101f78361026c565b925082610207576102066102a5565b5b828204905092915050565b600061021d8261026c565b91506102288361026c565b9250817fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff048311821515161561026157610260610276565b5b828202905092915050565b6000819050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b7f4d696e696d756d2076616c75652073686f756c64206265203120677765690000600082015250565b6103068161026c565b811461031157600080fd5b5056fea26469706673582212200fc7c09c5ba60850b4d85ac7a0418820e17979678e0b731773d7b1afb19f22d364736f6c63430008040033",
}

// DivideFunctionABI is the input ABI used to generate the binding from.
// Deprecated: Use DivideFunctionMetaData.ABI instead.
var DivideFunctionABI = DivideFunctionMetaData.ABI

// DivideFunctionBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use DivideFunctionMetaData.Bin instead.
var DivideFunctionBin = DivideFunctionMetaData.Bin

// DeployDivideFunction deploys a new Ethereum contract, binding an instance of DivideFunction to it.
func DeployDivideFunction(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *DivideFunction, error) {
	parsed, err := DivideFunctionMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(DivideFunctionBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &DivideFunction{DivideFunctionCaller: DivideFunctionCaller{contract: contract}, DivideFunctionTransactor: DivideFunctionTransactor{contract: contract}, DivideFunctionFilterer: DivideFunctionFilterer{contract: contract}}, nil
}

// DivideFunction is an auto generated Go binding around an Ethereum contract.
type DivideFunction struct {
	DivideFunctionCaller     // Read-only binding to the contract
	DivideFunctionTransactor // Write-only binding to the contract
	DivideFunctionFilterer   // Log filterer for contract events
}

// DivideFunctionCaller is an auto generated read-only Go binding around an Ethereum contract.
type DivideFunctionCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DivideFunctionTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DivideFunctionTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DivideFunctionFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DivideFunctionFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DivideFunctionSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DivideFunctionSession struct {
	Contract     *DivideFunction   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DivideFunctionCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DivideFunctionCallerSession struct {
	Contract *DivideFunctionCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// DivideFunctionTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DivideFunctionTransactorSession struct {
	Contract     *DivideFunctionTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// DivideFunctionRaw is an auto generated low-level Go binding around an Ethereum contract.
type DivideFunctionRaw struct {
	Contract *DivideFunction // Generic contract binding to access the raw methods on
}

// DivideFunctionCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DivideFunctionCallerRaw struct {
	Contract *DivideFunctionCaller // Generic read-only contract binding to access the raw methods on
}

// DivideFunctionTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DivideFunctionTransactorRaw struct {
	Contract *DivideFunctionTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDivideFunction creates a new instance of DivideFunction, bound to a specific deployed contract.
func NewDivideFunction(address common.Address, backend bind.ContractBackend) (*DivideFunction, error) {
	contract, err := bindDivideFunction(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DivideFunction{DivideFunctionCaller: DivideFunctionCaller{contract: contract}, DivideFunctionTransactor: DivideFunctionTransactor{contract: contract}, DivideFunctionFilterer: DivideFunctionFilterer{contract: contract}}, nil
}

// NewDivideFunctionCaller creates a new read-only instance of DivideFunction, bound to a specific deployed contract.
func NewDivideFunctionCaller(address common.Address, caller bind.ContractCaller) (*DivideFunctionCaller, error) {
	contract, err := bindDivideFunction(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DivideFunctionCaller{contract: contract}, nil
}

// NewDivideFunctionTransactor creates a new write-only instance of DivideFunction, bound to a specific deployed contract.
func NewDivideFunctionTransactor(address common.Address, transactor bind.ContractTransactor) (*DivideFunctionTransactor, error) {
	contract, err := bindDivideFunction(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DivideFunctionTransactor{contract: contract}, nil
}

// NewDivideFunctionFilterer creates a new log filterer instance of DivideFunction, bound to a specific deployed contract.
func NewDivideFunctionFilterer(address common.Address, filterer bind.ContractFilterer) (*DivideFunctionFilterer, error) {
	contract, err := bindDivideFunction(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DivideFunctionFilterer{contract: contract}, nil
}

// bindDivideFunction binds a generic wrapper to an already deployed contract.
func bindDivideFunction(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DivideFunctionABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DivideFunction *DivideFunctionRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DivideFunction.Contract.DivideFunctionCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DivideFunction *DivideFunctionRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DivideFunction.Contract.DivideFunctionTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DivideFunction *DivideFunctionRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DivideFunction.Contract.DivideFunctionTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DivideFunction *DivideFunctionCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DivideFunction.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DivideFunction *DivideFunctionTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DivideFunction.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DivideFunction *DivideFunctionTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DivideFunction.Contract.contract.Transact(opts, method, params...)
}

// WeiToKfive is a free data retrieval call binding the contract method 0xd7a65f0b.
//
// Solidity: function weiToKfive(uint256 weiValue, uint256 rate) pure returns(uint256)
func (_DivideFunction *DivideFunctionCaller) WeiToKfive(opts *bind.CallOpts, weiValue *big.Int, rate *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _DivideFunction.contract.Call(opts, &out, "weiToKfive", weiValue, rate)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WeiToKfive is a free data retrieval call binding the contract method 0xd7a65f0b.
//
// Solidity: function weiToKfive(uint256 weiValue, uint256 rate) pure returns(uint256)
func (_DivideFunction *DivideFunctionSession) WeiToKfive(weiValue *big.Int, rate *big.Int) (*big.Int, error) {
	return _DivideFunction.Contract.WeiToKfive(&_DivideFunction.CallOpts, weiValue, rate)
}

// WeiToKfive is a free data retrieval call binding the contract method 0xd7a65f0b.
//
// Solidity: function weiToKfive(uint256 weiValue, uint256 rate) pure returns(uint256)
func (_DivideFunction *DivideFunctionCallerSession) WeiToKfive(weiValue *big.Int, rate *big.Int) (*big.Int, error) {
	return _DivideFunction.Contract.WeiToKfive(&_DivideFunction.CallOpts, weiValue, rate)
}
