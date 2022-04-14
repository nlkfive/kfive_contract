// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package smc

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

// SmcMetaData contains all meta data concerning the Smc contract.
var SmcMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_input\",\"type\":\"bytes\"}],\"name\":\"hash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_text\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_num\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"encodePacked\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_num\",\"type\":\"uint256\"}],\"name\":\"encodePacked2\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"_fake\",\"type\":\"bool\"},{\"internalType\":\"bytes32\",\"name\":\"_secret\",\"type\":\"bytes32\"}],\"name\":\"encodePacked3\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"hour2uint256\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]",
}

// SmcABI is the input ABI used to generate the binding from.
// Deprecated: Use SmcMetaData.ABI instead.
var SmcABI = SmcMetaData.ABI

// Smc is an auto generated Go binding around an Ethereum contract.
type Smc struct {
	SmcCaller     // Read-only binding to the contract
	SmcTransactor // Write-only binding to the contract
	SmcFilterer   // Log filterer for contract events
}

// SmcCaller is an auto generated read-only Go binding around an Ethereum contract.
type SmcCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SmcTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SmcTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SmcFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SmcFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SmcSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SmcSession struct {
	Contract     *Smc              // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SmcCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SmcCallerSession struct {
	Contract *SmcCaller    // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// SmcTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SmcTransactorSession struct {
	Contract     *SmcTransactor    // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SmcRaw is an auto generated low-level Go binding around an Ethereum contract.
type SmcRaw struct {
	Contract *Smc // Generic contract binding to access the raw methods on
}

// SmcCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SmcCallerRaw struct {
	Contract *SmcCaller // Generic read-only contract binding to access the raw methods on
}

// SmcTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SmcTransactorRaw struct {
	Contract *SmcTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSmc creates a new instance of Smc, bound to a specific deployed contract.
func NewSmc(address common.Address, backend bind.ContractBackend) (*Smc, error) {
	contract, err := bindSmc(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Smc{SmcCaller: SmcCaller{contract: contract}, SmcTransactor: SmcTransactor{contract: contract}, SmcFilterer: SmcFilterer{contract: contract}}, nil
}

// NewSmcCaller creates a new read-only instance of Smc, bound to a specific deployed contract.
func NewSmcCaller(address common.Address, caller bind.ContractCaller) (*SmcCaller, error) {
	contract, err := bindSmc(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SmcCaller{contract: contract}, nil
}

// NewSmcTransactor creates a new write-only instance of Smc, bound to a specific deployed contract.
func NewSmcTransactor(address common.Address, transactor bind.ContractTransactor) (*SmcTransactor, error) {
	contract, err := bindSmc(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SmcTransactor{contract: contract}, nil
}

// NewSmcFilterer creates a new log filterer instance of Smc, bound to a specific deployed contract.
func NewSmcFilterer(address common.Address, filterer bind.ContractFilterer) (*SmcFilterer, error) {
	contract, err := bindSmc(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SmcFilterer{contract: contract}, nil
}

// bindSmc binds a generic wrapper to an already deployed contract.
func bindSmc(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SmcABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Smc *SmcRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Smc.Contract.SmcCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Smc *SmcRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Smc.Contract.SmcTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Smc *SmcRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Smc.Contract.SmcTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Smc *SmcCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Smc.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Smc *SmcTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Smc.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Smc *SmcTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Smc.Contract.contract.Transact(opts, method, params...)
}

// EncodePacked is a free data retrieval call binding the contract method 0x4a52893c.
//
// Solidity: function encodePacked(string _text, uint256 _num, address _addr) pure returns(bytes)
func (_Smc *SmcCaller) EncodePacked(opts *bind.CallOpts, _text string, _num *big.Int, _addr common.Address) ([]byte, error) {
	var out []interface{}
	err := _Smc.contract.Call(opts, &out, "encodePacked", _text, _num, _addr)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// EncodePacked is a free data retrieval call binding the contract method 0x4a52893c.
//
// Solidity: function encodePacked(string _text, uint256 _num, address _addr) pure returns(bytes)
func (_Smc *SmcSession) EncodePacked(_text string, _num *big.Int, _addr common.Address) ([]byte, error) {
	return _Smc.Contract.EncodePacked(&_Smc.CallOpts, _text, _num, _addr)
}

// EncodePacked is a free data retrieval call binding the contract method 0x4a52893c.
//
// Solidity: function encodePacked(string _text, uint256 _num, address _addr) pure returns(bytes)
func (_Smc *SmcCallerSession) EncodePacked(_text string, _num *big.Int, _addr common.Address) ([]byte, error) {
	return _Smc.Contract.EncodePacked(&_Smc.CallOpts, _text, _num, _addr)
}

// EncodePacked2 is a free data retrieval call binding the contract method 0xe995f5a9.
//
// Solidity: function encodePacked2(address _address, uint256 _num) pure returns(bytes)
func (_Smc *SmcCaller) EncodePacked2(opts *bind.CallOpts, _address common.Address, _num *big.Int) ([]byte, error) {
	var out []interface{}
	err := _Smc.contract.Call(opts, &out, "encodePacked2", _address, _num)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// EncodePacked2 is a free data retrieval call binding the contract method 0xe995f5a9.
//
// Solidity: function encodePacked2(address _address, uint256 _num) pure returns(bytes)
func (_Smc *SmcSession) EncodePacked2(_address common.Address, _num *big.Int) ([]byte, error) {
	return _Smc.Contract.EncodePacked2(&_Smc.CallOpts, _address, _num)
}

// EncodePacked2 is a free data retrieval call binding the contract method 0xe995f5a9.
//
// Solidity: function encodePacked2(address _address, uint256 _num) pure returns(bytes)
func (_Smc *SmcCallerSession) EncodePacked2(_address common.Address, _num *big.Int) ([]byte, error) {
	return _Smc.Contract.EncodePacked2(&_Smc.CallOpts, _address, _num)
}

// EncodePacked3 is a free data retrieval call binding the contract method 0x003adb58.
//
// Solidity: function encodePacked3(uint256 _value, bool _fake, bytes32 _secret) pure returns(bytes)
func (_Smc *SmcCaller) EncodePacked3(opts *bind.CallOpts, _value *big.Int, _fake bool, _secret [32]byte) ([]byte, error) {
	var out []interface{}
	err := _Smc.contract.Call(opts, &out, "encodePacked3", _value, _fake, _secret)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// EncodePacked3 is a free data retrieval call binding the contract method 0x003adb58.
//
// Solidity: function encodePacked3(uint256 _value, bool _fake, bytes32 _secret) pure returns(bytes)
func (_Smc *SmcSession) EncodePacked3(_value *big.Int, _fake bool, _secret [32]byte) ([]byte, error) {
	return _Smc.Contract.EncodePacked3(&_Smc.CallOpts, _value, _fake, _secret)
}

// EncodePacked3 is a free data retrieval call binding the contract method 0x003adb58.
//
// Solidity: function encodePacked3(uint256 _value, bool _fake, bytes32 _secret) pure returns(bytes)
func (_Smc *SmcCallerSession) EncodePacked3(_value *big.Int, _fake bool, _secret [32]byte) ([]byte, error) {
	return _Smc.Contract.EncodePacked3(&_Smc.CallOpts, _value, _fake, _secret)
}

// Hash is a free data retrieval call binding the contract method 0xaa1e84de.
//
// Solidity: function hash(bytes _input) pure returns(bytes32)
func (_Smc *SmcCaller) Hash(opts *bind.CallOpts, _input []byte) ([32]byte, error) {
	var out []interface{}
	err := _Smc.contract.Call(opts, &out, "hash", _input)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// Hash is a free data retrieval call binding the contract method 0xaa1e84de.
//
// Solidity: function hash(bytes _input) pure returns(bytes32)
func (_Smc *SmcSession) Hash(_input []byte) ([32]byte, error) {
	return _Smc.Contract.Hash(&_Smc.CallOpts, _input)
}

// Hash is a free data retrieval call binding the contract method 0xaa1e84de.
//
// Solidity: function hash(bytes _input) pure returns(bytes32)
func (_Smc *SmcCallerSession) Hash(_input []byte) ([32]byte, error) {
	return _Smc.Contract.Hash(&_Smc.CallOpts, _input)
}

// Hour2uint256 is a free data retrieval call binding the contract method 0xfe8754ce.
//
// Solidity: function hour2uint256() pure returns(uint256)
func (_Smc *SmcCaller) Hour2uint256(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Smc.contract.Call(opts, &out, "hour2uint256")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Hour2uint256 is a free data retrieval call binding the contract method 0xfe8754ce.
//
// Solidity: function hour2uint256() pure returns(uint256)
func (_Smc *SmcSession) Hour2uint256() (*big.Int, error) {
	return _Smc.Contract.Hour2uint256(&_Smc.CallOpts)
}

// Hour2uint256 is a free data retrieval call binding the contract method 0xfe8754ce.
//
// Solidity: function hour2uint256() pure returns(uint256)
func (_Smc *SmcCallerSession) Hour2uint256() (*big.Int, error) {
	return _Smc.Contract.Hour2uint256(&_Smc.CallOpts)
}
