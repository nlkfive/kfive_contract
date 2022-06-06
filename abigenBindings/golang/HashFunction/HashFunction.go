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
	Bin: "0x608060405234801561001057600080fd5b50610900806100206000396000f3fe608060405234801561001057600080fd5b50600436106100565760003560e01c80623adb581461005b5780634a52893c1461008b578063aa1e84de146100bb578063e995f5a9146100eb578063fe8754ce1461011b575b600080fd5b610075600480360381019061007091906103eb565b610139565b60405161008291906105fd565b60405180910390f35b6100a560048036038101906100a09190610384565b610168565b6040516100b291906105fd565b60405180910390f35b6100d560048036038101906100d09190610343565b61019c565b6040516100e291906105e2565b60405180910390f35b61010560048036038101906101009190610307565b6101ad565b60405161011291906105fd565b60405180910390f35b6101236101d9565b604051610130919061061f565b60405180910390f35b6060838383604051602001610150939291906105a5565b60405160208183030381529060405290509392505050565b606083838460008560405160200161018495949392919061054a565b60405160208183030381529060405290509392505050565b600081805190602001209050919050565b606082826040516020016101c292919061051e565b604051602081830303815290604052905092915050565b6000610e10905090565b60006101f66101f18461065f565b61063a565b90508281526020810184848401111561020e57600080fd5b610219848285610745565b509392505050565b600061023461022f84610690565b61063a565b90508281526020810184848401111561024c57600080fd5b610257848285610745565b509392505050565b60008135905061026e8161086e565b92915050565b60008135905061028381610885565b92915050565b6000813590506102988161089c565b92915050565b600082601f8301126102af57600080fd5b81356102bf8482602086016101e3565b91505092915050565b600082601f8301126102d957600080fd5b81356102e9848260208601610221565b91505092915050565b600081359050610301816108b3565b92915050565b6000806040838503121561031a57600080fd5b60006103288582860161025f565b9250506020610339858286016102f2565b9150509250929050565b60006020828403121561035557600080fd5b600082013567ffffffffffffffff81111561036f57600080fd5b61037b8482850161029e565b91505092915050565b60008060006060848603121561039957600080fd5b600084013567ffffffffffffffff8111156103b357600080fd5b6103bf868287016102c8565b93505060206103d0868287016102f2565b92505060406103e18682870161025f565b9150509250925092565b60008060006060848603121561040057600080fd5b600061040e868287016102f2565b935050602061041f86828701610274565b925050604061043086828701610289565b9150509250925092565b61044b610446826106f3565b6107b8565b82525050565b61046261045d82610705565b6107ca565b82525050565b61047181610711565b82525050565b61048861048382610711565b6107dc565b82525050565b6000610499826106c1565b6104a381856106d7565b93506104b3818560208601610754565b6104bc81610843565b840191505092915050565b60006104d2826106cc565b6104dc81856106e8565b93506104ec818560208601610754565b80840191505092915050565b6105018161073b565b82525050565b6105186105138261073b565b6107f8565b82525050565b600061052a828561043a565b60148201915061053a8284610507565b6020820191508190509392505050565b600061055682886104c7565b91506105628287610507565b6020820191506105728286610507565b6020820191506105828285610451565b600182019150610592828461043a565b6014820191508190509695505050505050565b60006105b18286610507565b6020820191506105c18285610451565b6001820191506105d18284610477565b602082019150819050949350505050565b60006020820190506105f76000830184610468565b92915050565b60006020820190508181036000830152610617818461048e565b905092915050565b600060208201905061063460008301846104f8565b92915050565b6000610644610655565b90506106508282610787565b919050565b6000604051905090565b600067ffffffffffffffff82111561067a57610679610814565b5b61068382610843565b9050602081019050919050565b600067ffffffffffffffff8211156106ab576106aa610814565b5b6106b482610843565b9050602081019050919050565b600081519050919050565b600081519050919050565b600082825260208201905092915050565b600081905092915050565b60006106fe8261071b565b9050919050565b60008115159050919050565b6000819050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b82818337600083830152505050565b60005b83811015610772578082015181840152602081019050610757565b83811115610781576000848401525b50505050565b61079082610843565b810181811067ffffffffffffffff821117156107af576107ae610814565b5b80604052505050565b60006107c3826107e6565b9050919050565b60006107d582610802565b9050919050565b6000819050919050565b60006107f182610861565b9050919050565b6000819050919050565b600061080d82610854565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6000601f19601f8301169050919050565b60008160f81b9050919050565b60008160601b9050919050565b610877816106f3565b811461088257600080fd5b50565b61088e81610705565b811461089957600080fd5b50565b6108a581610711565b81146108b057600080fd5b50565b6108bc8161073b565b81146108c757600080fd5b5056fea2646970667358221220f0374ee1ac3b4b1386366f72de6481ec19650dfed169502052e84a611537700b64736f6c63430008040033",
}

// HashFunctionABI is the input ABI used to generate the binding from.
// Deprecated: Use HashFunctionMetaData.ABI instead.
var HashFunctionABI = HashFunctionMetaData.ABI

// HashFunctionBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use HashFunctionMetaData.Bin instead.
var HashFunctionBin = HashFunctionMetaData.Bin

// DeployHashFunction deploys a new Ethereum contract, binding an instance of HashFunction to it.
func DeployHashFunction(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *HashFunction, error) {
	parsed, err := HashFunctionMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(HashFunctionBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &HashFunction{HashFunctionCaller: HashFunctionCaller{contract: contract}, HashFunctionTransactor: HashFunctionTransactor{contract: contract}, HashFunctionFilterer: HashFunctionFilterer{contract: contract}}, nil
}

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
