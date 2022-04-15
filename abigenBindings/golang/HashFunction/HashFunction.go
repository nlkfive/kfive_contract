// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package HashFunction

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

// HashFunctionMetaData contains all meta data concerning the HashFunction contract.
var HashFunctionMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_input\",\"type\":\"bytes\"}],\"name\":\"hash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_text\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_num\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"encodePacked\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_num\",\"type\":\"uint256\"}],\"name\":\"encodePacked2\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"_fake\",\"type\":\"bool\"},{\"internalType\":\"bytes32\",\"name\":\"_secret\",\"type\":\"bytes32\"}],\"name\":\"encodePacked3\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"hour2uint256\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]",
}

// HashFunctionABI is the input ABI used to generate the binding from.
// Deprecated: Use HashFunctionMetaData.ABI instead.
var HashFunctionABI = HashFunctionMetaData.ABI

// HashFunction is an auto generated Go binding around an Ethereum contract.
type HashFunction struct {
	HashFunctionCaller     // Read-only binding to the contract
	HashFunctionTransactor // Write-only binding to the contract
	HashFunctionFilterer   // Log filterer for contract events
}

// HashFunctionCaller is an auto generated read-only Go binding around an Ethereum contract.
type HashFunctionCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HashFunctionTransactor is an auto generated write-only Go binding around an Ethereum contract.
type HashFunctionTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HashFunctionFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type HashFunctionFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HashFunctionSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type HashFunctionSession struct {
	Contract     *HashFunction     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// HashFunctionCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type HashFunctionCallerSession struct {
	Contract *HashFunctionCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// HashFunctionTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type HashFunctionTransactorSession struct {
	Contract     *HashFunctionTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// HashFunctionRaw is an auto generated low-level Go binding around an Ethereum contract.
type HashFunctionRaw struct {
	Contract *HashFunction // Generic contract binding to access the raw methods on
}

// HashFunctionCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type HashFunctionCallerRaw struct {
	Contract *HashFunctionCaller // Generic read-only contract binding to access the raw methods on
}

// HashFunctionTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type HashFunctionTransactorRaw struct {
	Contract *HashFunctionTransactor // Generic write-only contract binding to access the raw methods on
}

// NewHashFunction creates a new instance of HashFunction, bound to a specific deployed contract.
func NewHashFunction(address common.Address, backend bind.ContractBackend) (*HashFunction, error) {
	contract, err := bindHashFunction(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &HashFunction{HashFunctionCaller: HashFunctionCaller{contract: contract}, HashFunctionTransactor: HashFunctionTransactor{contract: contract}, HashFunctionFilterer: HashFunctionFilterer{contract: contract}}, nil
}

// NewHashFunctionCaller creates a new read-only instance of HashFunction, bound to a specific deployed contract.
func NewHashFunctionCaller(address common.Address, caller bind.ContractCaller) (*HashFunctionCaller, error) {
	contract, err := bindHashFunction(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &HashFunctionCaller{contract: contract}, nil
}

// NewHashFunctionTransactor creates a new write-only instance of HashFunction, bound to a specific deployed contract.
func NewHashFunctionTransactor(address common.Address, transactor bind.ContractTransactor) (*HashFunctionTransactor, error) {
	contract, err := bindHashFunction(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &HashFunctionTransactor{contract: contract}, nil
}

// NewHashFunctionFilterer creates a new log filterer instance of HashFunction, bound to a specific deployed contract.
func NewHashFunctionFilterer(address common.Address, filterer bind.ContractFilterer) (*HashFunctionFilterer, error) {
	contract, err := bindHashFunction(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &HashFunctionFilterer{contract: contract}, nil
}

// bindHashFunction binds a generic wrapper to an already deployed contract.
func bindHashFunction(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(HashFunctionABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_HashFunction *HashFunctionRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _HashFunction.Contract.HashFunctionCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_HashFunction *HashFunctionRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HashFunction.Contract.HashFunctionTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_HashFunction *HashFunctionRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _HashFunction.Contract.HashFunctionTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_HashFunction *HashFunctionCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _HashFunction.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_HashFunction *HashFunctionTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HashFunction.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_HashFunction *HashFunctionTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _HashFunction.Contract.contract.Transact(opts, method, params...)
}

// EncodePacked is a free data retrieval call binding the contract method 0x4a52893c.
//
// Solidity: function encodePacked(string _text, uint256 _num, address _addr) pure returns(bytes)
func (_HashFunction *HashFunctionCaller) EncodePacked(opts *bind.CallOpts, _text string, _num *big.Int, _addr common.Address) ([]byte, error) {
	var out []interface{}
	err := _HashFunction.contract.Call(opts, &out, "encodePacked", _text, _num, _addr)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// EncodePacked is a free data retrieval call binding the contract method 0x4a52893c.
//
// Solidity: function encodePacked(string _text, uint256 _num, address _addr) pure returns(bytes)
func (_HashFunction *HashFunctionSession) EncodePacked(_text string, _num *big.Int, _addr common.Address) ([]byte, error) {
	return _HashFunction.Contract.EncodePacked(&_HashFunction.CallOpts, _text, _num, _addr)
}

// EncodePacked is a free data retrieval call binding the contract method 0x4a52893c.
//
// Solidity: function encodePacked(string _text, uint256 _num, address _addr) pure returns(bytes)
func (_HashFunction *HashFunctionCallerSession) EncodePacked(_text string, _num *big.Int, _addr common.Address) ([]byte, error) {
	return _HashFunction.Contract.EncodePacked(&_HashFunction.CallOpts, _text, _num, _addr)
}

// EncodePacked2 is a free data retrieval call binding the contract method 0xe995f5a9.
//
// Solidity: function encodePacked2(address _address, uint256 _num) pure returns(bytes)
func (_HashFunction *HashFunctionCaller) EncodePacked2(opts *bind.CallOpts, _address common.Address, _num *big.Int) ([]byte, error) {
	var out []interface{}
	err := _HashFunction.contract.Call(opts, &out, "encodePacked2", _address, _num)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// EncodePacked2 is a free data retrieval call binding the contract method 0xe995f5a9.
//
// Solidity: function encodePacked2(address _address, uint256 _num) pure returns(bytes)
func (_HashFunction *HashFunctionSession) EncodePacked2(_address common.Address, _num *big.Int) ([]byte, error) {
	return _HashFunction.Contract.EncodePacked2(&_HashFunction.CallOpts, _address, _num)
}

// EncodePacked2 is a free data retrieval call binding the contract method 0xe995f5a9.
//
// Solidity: function encodePacked2(address _address, uint256 _num) pure returns(bytes)
func (_HashFunction *HashFunctionCallerSession) EncodePacked2(_address common.Address, _num *big.Int) ([]byte, error) {
	return _HashFunction.Contract.EncodePacked2(&_HashFunction.CallOpts, _address, _num)
}

// EncodePacked3 is a free data retrieval call binding the contract method 0x003adb58.
//
// Solidity: function encodePacked3(uint256 _value, bool _fake, bytes32 _secret) pure returns(bytes)
func (_HashFunction *HashFunctionCaller) EncodePacked3(opts *bind.CallOpts, _value *big.Int, _fake bool, _secret [32]byte) ([]byte, error) {
	var out []interface{}
	err := _HashFunction.contract.Call(opts, &out, "encodePacked3", _value, _fake, _secret)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// EncodePacked3 is a free data retrieval call binding the contract method 0x003adb58.
//
// Solidity: function encodePacked3(uint256 _value, bool _fake, bytes32 _secret) pure returns(bytes)
func (_HashFunction *HashFunctionSession) EncodePacked3(_value *big.Int, _fake bool, _secret [32]byte) ([]byte, error) {
	return _HashFunction.Contract.EncodePacked3(&_HashFunction.CallOpts, _value, _fake, _secret)
}

// EncodePacked3 is a free data retrieval call binding the contract method 0x003adb58.
//
// Solidity: function encodePacked3(uint256 _value, bool _fake, bytes32 _secret) pure returns(bytes)
func (_HashFunction *HashFunctionCallerSession) EncodePacked3(_value *big.Int, _fake bool, _secret [32]byte) ([]byte, error) {
	return _HashFunction.Contract.EncodePacked3(&_HashFunction.CallOpts, _value, _fake, _secret)
}

// Hash is a free data retrieval call binding the contract method 0xaa1e84de.
//
// Solidity: function hash(bytes _input) pure returns(bytes32)
func (_HashFunction *HashFunctionCaller) Hash(opts *bind.CallOpts, _input []byte) ([32]byte, error) {
	var out []interface{}
	err := _HashFunction.contract.Call(opts, &out, "hash", _input)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// Hash is a free data retrieval call binding the contract method 0xaa1e84de.
//
// Solidity: function hash(bytes _input) pure returns(bytes32)
func (_HashFunction *HashFunctionSession) Hash(_input []byte) ([32]byte, error) {
	return _HashFunction.Contract.Hash(&_HashFunction.CallOpts, _input)
}

// Hash is a free data retrieval call binding the contract method 0xaa1e84de.
//
// Solidity: function hash(bytes _input) pure returns(bytes32)
func (_HashFunction *HashFunctionCallerSession) Hash(_input []byte) ([32]byte, error) {
	return _HashFunction.Contract.Hash(&_HashFunction.CallOpts, _input)
}

// Hour2uint256 is a free data retrieval call binding the contract method 0xfe8754ce.
//
// Solidity: function hour2uint256() pure returns(uint256)
func (_HashFunction *HashFunctionCaller) Hour2uint256(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _HashFunction.contract.Call(opts, &out, "hour2uint256")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Hour2uint256 is a free data retrieval call binding the contract method 0xfe8754ce.
//
// Solidity: function hour2uint256() pure returns(uint256)
func (_HashFunction *HashFunctionSession) Hour2uint256() (*big.Int, error) {
	return _HashFunction.Contract.Hour2uint256(&_HashFunction.CallOpts)
}

// Hour2uint256 is a free data retrieval call binding the contract method 0xfe8754ce.
//
// Solidity: function hour2uint256() pure returns(uint256)
func (_HashFunction *HashFunctionCallerSession) Hour2uint256() (*big.Int, error) {
	return _HashFunction.Contract.Hour2uint256(&_HashFunction.CallOpts)
}
