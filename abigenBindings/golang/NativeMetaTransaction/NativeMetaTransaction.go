// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package NativeMetaTransaction

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

// NativeMetaTransactionMetaData contains all meta data concerning the NativeMetaTransaction contract.
var NativeMetaTransactionMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"userAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"relayerAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"functionSignature\",\"type\":\"bytes\"}],\"name\":\"MetaTransactionExecuted\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"domainSeparator\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getChainId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"userAddress\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"functionSignature\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"sigR\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"sigS\",\"type\":\"bytes32\"},{\"internalType\":\"uint8\",\"name\":\"sigV\",\"type\":\"uint8\"}],\"name\":\"executeMetaTransaction\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"getNonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50610e33806100206000396000f3fe60806040526004361061003f5760003560e01c80630c53c51c146100445780632d0335ab146100745780633408e470146100b1578063f698da25146100dc575b600080fd5b61005e60048036038101906100599190610641565b610107565b60405161006b9190610989565b60405180910390f35b34801561008057600080fd5b5061009b60048036038101906100969190610618565b610371565b6040516100a89190610a0b565b60405180910390f35b3480156100bd57600080fd5b506100c66103ba565b6040516100d39190610a0b565b60405180910390f35b3480156100e857600080fd5b506100f16103c7565b6040516100fe91906108e4565b60405180910390f35b606060006040518060600160405280600160008a73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205481526020018873ffffffffffffffffffffffffffffffffffffffff16815260200187815250905061018a87828787876103cd565b6101c9576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016101c0906109ab565b60405180910390fd5b60018060008973ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020546102149190610abf565b600160008973ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055507f5845892132946850460bff5a0083f71031bc5bf9aadcd40f1de79423eac9b10b87338860405161028a939291906108a6565b60405180910390a16000803073ffffffffffffffffffffffffffffffffffffffff16888a6040516020016102bf929190610847565b6040516020818303038152906040526040516102db9190610830565b6000604051808303816000865af19150503d8060008114610318576040519150601f19603f3d011682016040523d82523d6000602084013e61031d565b606091505b509150915081610362576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610359906109cb565b60405180910390fd5b80935050505095945050505050565b6000600160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020549050919050565b6000804690508091505090565b60005481565b60008073ffffffffffffffffffffffffffffffffffffffff168673ffffffffffffffffffffffffffffffffffffffff16141561043e576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610435906109eb565b60405180910390fd5b600161045161044c876104d6565b61053e565b838686604051600081526020016040526040516104719493929190610944565b6020604051602081039080840390855afa158015610493573d6000803e3d6000fd5b5050506020604051035173ffffffffffffffffffffffffffffffffffffffff168673ffffffffffffffffffffffffffffffffffffffff1614905095945050505050565b6000604051806080016040528060438152602001610dbb60439139805190602001208260000151836020015184604001518051906020012060405160200161052194939291906108ff565b604051602081830303815290604052805190602001209050919050565b600080548260405160200161055492919061086f565b604051602081830303815290604052805190602001209050919050565b600061058461057f84610a4b565b610a26565b90508281526020810184848401111561059c57600080fd5b6105a7848285610b68565b509392505050565b6000813590506105be81610d75565b92915050565b6000813590506105d381610d8c565b92915050565b600082601f8301126105ea57600080fd5b81356105fa848260208601610571565b91505092915050565b60008135905061061281610da3565b92915050565b60006020828403121561062a57600080fd5b6000610638848285016105af565b91505092915050565b600080600080600060a0868803121561065957600080fd5b6000610667888289016105af565b955050602086013567ffffffffffffffff81111561068457600080fd5b610690888289016105d9565b94505060406106a1888289016105c4565b93505060606106b2888289016105c4565b92505060806106c388828901610603565b9150509295509295909350565b6106d981610b15565b82525050565b6106f06106eb82610b15565b610bdb565b82525050565b6106ff81610b27565b82525050565b61071661071182610b27565b610bed565b82525050565b600061072782610a7c565b6107318185610a87565b9350610741818560208601610b77565b61074a81610c67565b840191505092915050565b600061076082610a7c565b61076a8185610a98565b935061077a818560208601610b77565b80840191505092915050565b6000610793600283610ab4565b915061079e82610c85565b600282019050919050565b60006107b6603d83610aa3565b91506107c182610cae565b604082019050919050565b60006107d9602783610aa3565b91506107e482610cfd565b604082019050919050565b60006107fc601a83610aa3565b915061080782610d4c565b602082019050919050565b61081b81610b51565b82525050565b61082a81610b5b565b82525050565b600061083c8284610755565b915081905092915050565b60006108538285610755565b915061085f82846106df565b6014820191508190509392505050565b600061087a82610786565b91506108868285610705565b6020820191506108968284610705565b6020820191508190509392505050565b60006060820190506108bb60008301866106d0565b6108c860208301856106d0565b81810360408301526108da818461071c565b9050949350505050565b60006020820190506108f960008301846106f6565b92915050565b600060808201905061091460008301876106f6565b6109216020830186610812565b61092e60408301856106d0565b61093b60608301846106f6565b95945050505050565b600060808201905061095960008301876106f6565b6109666020830186610821565b61097360408301856106f6565b61098060608301846106f6565b95945050505050565b600060208201905081810360008301526109a3818461071c565b905092915050565b600060208201905081810360008301526109c4816107a9565b9050919050565b600060208201905081810360008301526109e4816107cc565b9050919050565b60006020820190508181036000830152610a04816107ef565b9050919050565b6000602082019050610a206000830184610812565b92915050565b6000610a30610a41565b9050610a3c8282610baa565b919050565b6000604051905090565b600067ffffffffffffffff821115610a6657610a65610c38565b5b610a6f82610c67565b9050602081019050919050565b600081519050919050565b600082825260208201905092915050565b600081905092915050565b600082825260208201905092915050565b600081905092915050565b6000610aca82610b51565b9150610ad583610b51565b9250827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff03821115610b0a57610b09610c09565b5b828201905092915050565b6000610b2082610b31565b9050919050565b6000819050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b600060ff82169050919050565b82818337600083830152505050565b60005b83811015610b95578082015181840152602081019050610b7a565b83811115610ba4576000848401525b50505050565b610bb382610c67565b810181811067ffffffffffffffff82111715610bd257610bd1610c38565b5b80604052505050565b6000610be682610bf7565b9050919050565b6000819050919050565b6000610c0282610c78565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6000601f19601f8301169050919050565b60008160601b9050919050565b7f1901000000000000000000000000000000000000000000000000000000000000600082015250565b7f4e4d5423657865637574654d6574615472616e73616374696f6e3a205349474e60008201527f45525f414e445f5349474e41545552455f444f5f4e4f545f4d41544348000000602082015250565b7f4e4d5423657865637574654d6574615472616e73616374696f6e3a2043414c4c60008201527f5f4641494c454400000000000000000000000000000000000000000000000000602082015250565b7f4e4d54237665726966793a20494e56414c49445f5349474e4552000000000000600082015250565b610d7e81610b15565b8114610d8957600080fd5b50565b610d9581610b27565b8114610da057600080fd5b50565b610dac81610b5b565b8114610db757600080fd5b5056fe4d6574615472616e73616374696f6e2875696e74323536206e6f6e63652c616464726573732066726f6d2c62797465732066756e6374696f6e5369676e617475726529a2646970667358221220e10b9d1e087b407359ca9afb527ff07f673a0abbc0a3adce4719da05173680ad64736f6c63430008040033",
}

// NativeMetaTransactionABI is the input ABI used to generate the binding from.
// Deprecated: Use NativeMetaTransactionMetaData.ABI instead.
var NativeMetaTransactionABI = NativeMetaTransactionMetaData.ABI

// NativeMetaTransactionBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use NativeMetaTransactionMetaData.Bin instead.
var NativeMetaTransactionBin = NativeMetaTransactionMetaData.Bin

// DeployNativeMetaTransaction deploys a new Ethereum contract, binding an instance of NativeMetaTransaction to it.
func DeployNativeMetaTransaction(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *NativeMetaTransaction, error) {
	parsed, err := NativeMetaTransactionMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(NativeMetaTransactionBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &NativeMetaTransaction{NativeMetaTransactionCaller: NativeMetaTransactionCaller{contract: contract}, NativeMetaTransactionTransactor: NativeMetaTransactionTransactor{contract: contract}, NativeMetaTransactionFilterer: NativeMetaTransactionFilterer{contract: contract}}, nil
}

// NativeMetaTransaction is an auto generated Go binding around an Ethereum contract.
type NativeMetaTransaction struct {
	NativeMetaTransactionCaller     // Read-only binding to the contract
	NativeMetaTransactionTransactor // Write-only binding to the contract
	NativeMetaTransactionFilterer   // Log filterer for contract events
}

// NativeMetaTransactionCaller is an auto generated read-only Go binding around an Ethereum contract.
type NativeMetaTransactionCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NativeMetaTransactionTransactor is an auto generated write-only Go binding around an Ethereum contract.
type NativeMetaTransactionTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NativeMetaTransactionFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type NativeMetaTransactionFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NativeMetaTransactionSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type NativeMetaTransactionSession struct {
	Contract     *NativeMetaTransaction // Generic contract binding to set the session for
	CallOpts     bind.CallOpts          // Call options to use throughout this session
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// NativeMetaTransactionCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type NativeMetaTransactionCallerSession struct {
	Contract *NativeMetaTransactionCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                // Call options to use throughout this session
}

// NativeMetaTransactionTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type NativeMetaTransactionTransactorSession struct {
	Contract     *NativeMetaTransactionTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                // Transaction auth options to use throughout this session
}

// NativeMetaTransactionRaw is an auto generated low-level Go binding around an Ethereum contract.
type NativeMetaTransactionRaw struct {
	Contract *NativeMetaTransaction // Generic contract binding to access the raw methods on
}

// NativeMetaTransactionCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type NativeMetaTransactionCallerRaw struct {
	Contract *NativeMetaTransactionCaller // Generic read-only contract binding to access the raw methods on
}

// NativeMetaTransactionTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type NativeMetaTransactionTransactorRaw struct {
	Contract *NativeMetaTransactionTransactor // Generic write-only contract binding to access the raw methods on
}

// NewNativeMetaTransaction creates a new instance of NativeMetaTransaction, bound to a specific deployed contract.
func NewNativeMetaTransaction(address common.Address, backend bind.ContractBackend) (*NativeMetaTransaction, error) {
	contract, err := bindNativeMetaTransaction(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &NativeMetaTransaction{NativeMetaTransactionCaller: NativeMetaTransactionCaller{contract: contract}, NativeMetaTransactionTransactor: NativeMetaTransactionTransactor{contract: contract}, NativeMetaTransactionFilterer: NativeMetaTransactionFilterer{contract: contract}}, nil
}

// NewNativeMetaTransactionCaller creates a new read-only instance of NativeMetaTransaction, bound to a specific deployed contract.
func NewNativeMetaTransactionCaller(address common.Address, caller bind.ContractCaller) (*NativeMetaTransactionCaller, error) {
	contract, err := bindNativeMetaTransaction(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &NativeMetaTransactionCaller{contract: contract}, nil
}

// NewNativeMetaTransactionTransactor creates a new write-only instance of NativeMetaTransaction, bound to a specific deployed contract.
func NewNativeMetaTransactionTransactor(address common.Address, transactor bind.ContractTransactor) (*NativeMetaTransactionTransactor, error) {
	contract, err := bindNativeMetaTransaction(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &NativeMetaTransactionTransactor{contract: contract}, nil
}

// NewNativeMetaTransactionFilterer creates a new log filterer instance of NativeMetaTransaction, bound to a specific deployed contract.
func NewNativeMetaTransactionFilterer(address common.Address, filterer bind.ContractFilterer) (*NativeMetaTransactionFilterer, error) {
	contract, err := bindNativeMetaTransaction(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &NativeMetaTransactionFilterer{contract: contract}, nil
}

// bindNativeMetaTransaction binds a generic wrapper to an already deployed contract.
func bindNativeMetaTransaction(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(NativeMetaTransactionABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NativeMetaTransaction *NativeMetaTransactionRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _NativeMetaTransaction.Contract.NativeMetaTransactionCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NativeMetaTransaction *NativeMetaTransactionRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NativeMetaTransaction.Contract.NativeMetaTransactionTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NativeMetaTransaction *NativeMetaTransactionRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NativeMetaTransaction.Contract.NativeMetaTransactionTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NativeMetaTransaction *NativeMetaTransactionCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _NativeMetaTransaction.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NativeMetaTransaction *NativeMetaTransactionTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NativeMetaTransaction.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NativeMetaTransaction *NativeMetaTransactionTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NativeMetaTransaction.Contract.contract.Transact(opts, method, params...)
}

// DomainSeparator is a free data retrieval call binding the contract method 0xf698da25.
//
// Solidity: function domainSeparator() view returns(bytes32)
func (_NativeMetaTransaction *NativeMetaTransactionCaller) DomainSeparator(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _NativeMetaTransaction.contract.Call(opts, &out, "domainSeparator")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DomainSeparator is a free data retrieval call binding the contract method 0xf698da25.
//
// Solidity: function domainSeparator() view returns(bytes32)
func (_NativeMetaTransaction *NativeMetaTransactionSession) DomainSeparator() ([32]byte, error) {
	return _NativeMetaTransaction.Contract.DomainSeparator(&_NativeMetaTransaction.CallOpts)
}

// DomainSeparator is a free data retrieval call binding the contract method 0xf698da25.
//
// Solidity: function domainSeparator() view returns(bytes32)
func (_NativeMetaTransaction *NativeMetaTransactionCallerSession) DomainSeparator() ([32]byte, error) {
	return _NativeMetaTransaction.Contract.DomainSeparator(&_NativeMetaTransaction.CallOpts)
}

// GetChainId is a free data retrieval call binding the contract method 0x3408e470.
//
// Solidity: function getChainId() view returns(uint256)
func (_NativeMetaTransaction *NativeMetaTransactionCaller) GetChainId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NativeMetaTransaction.contract.Call(opts, &out, "getChainId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetChainId is a free data retrieval call binding the contract method 0x3408e470.
//
// Solidity: function getChainId() view returns(uint256)
func (_NativeMetaTransaction *NativeMetaTransactionSession) GetChainId() (*big.Int, error) {
	return _NativeMetaTransaction.Contract.GetChainId(&_NativeMetaTransaction.CallOpts)
}

// GetChainId is a free data retrieval call binding the contract method 0x3408e470.
//
// Solidity: function getChainId() view returns(uint256)
func (_NativeMetaTransaction *NativeMetaTransactionCallerSession) GetChainId() (*big.Int, error) {
	return _NativeMetaTransaction.Contract.GetChainId(&_NativeMetaTransaction.CallOpts)
}

// GetNonce is a free data retrieval call binding the contract method 0x2d0335ab.
//
// Solidity: function getNonce(address user) view returns(uint256 nonce)
func (_NativeMetaTransaction *NativeMetaTransactionCaller) GetNonce(opts *bind.CallOpts, user common.Address) (*big.Int, error) {
	var out []interface{}
	err := _NativeMetaTransaction.contract.Call(opts, &out, "getNonce", user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNonce is a free data retrieval call binding the contract method 0x2d0335ab.
//
// Solidity: function getNonce(address user) view returns(uint256 nonce)
func (_NativeMetaTransaction *NativeMetaTransactionSession) GetNonce(user common.Address) (*big.Int, error) {
	return _NativeMetaTransaction.Contract.GetNonce(&_NativeMetaTransaction.CallOpts, user)
}

// GetNonce is a free data retrieval call binding the contract method 0x2d0335ab.
//
// Solidity: function getNonce(address user) view returns(uint256 nonce)
func (_NativeMetaTransaction *NativeMetaTransactionCallerSession) GetNonce(user common.Address) (*big.Int, error) {
	return _NativeMetaTransaction.Contract.GetNonce(&_NativeMetaTransaction.CallOpts, user)
}

// ExecuteMetaTransaction is a paid mutator transaction binding the contract method 0x0c53c51c.
//
// Solidity: function executeMetaTransaction(address userAddress, bytes functionSignature, bytes32 sigR, bytes32 sigS, uint8 sigV) payable returns(bytes)
func (_NativeMetaTransaction *NativeMetaTransactionTransactor) ExecuteMetaTransaction(opts *bind.TransactOpts, userAddress common.Address, functionSignature []byte, sigR [32]byte, sigS [32]byte, sigV uint8) (*types.Transaction, error) {
	return _NativeMetaTransaction.contract.Transact(opts, "executeMetaTransaction", userAddress, functionSignature, sigR, sigS, sigV)
}

// ExecuteMetaTransaction is a paid mutator transaction binding the contract method 0x0c53c51c.
//
// Solidity: function executeMetaTransaction(address userAddress, bytes functionSignature, bytes32 sigR, bytes32 sigS, uint8 sigV) payable returns(bytes)
func (_NativeMetaTransaction *NativeMetaTransactionSession) ExecuteMetaTransaction(userAddress common.Address, functionSignature []byte, sigR [32]byte, sigS [32]byte, sigV uint8) (*types.Transaction, error) {
	return _NativeMetaTransaction.Contract.ExecuteMetaTransaction(&_NativeMetaTransaction.TransactOpts, userAddress, functionSignature, sigR, sigS, sigV)
}

// ExecuteMetaTransaction is a paid mutator transaction binding the contract method 0x0c53c51c.
//
// Solidity: function executeMetaTransaction(address userAddress, bytes functionSignature, bytes32 sigR, bytes32 sigS, uint8 sigV) payable returns(bytes)
func (_NativeMetaTransaction *NativeMetaTransactionTransactorSession) ExecuteMetaTransaction(userAddress common.Address, functionSignature []byte, sigR [32]byte, sigS [32]byte, sigV uint8) (*types.Transaction, error) {
	return _NativeMetaTransaction.Contract.ExecuteMetaTransaction(&_NativeMetaTransaction.TransactOpts, userAddress, functionSignature, sigR, sigS, sigV)
}

// NativeMetaTransactionMetaTransactionExecutedIterator is returned from FilterMetaTransactionExecuted and is used to iterate over the raw logs and unpacked data for MetaTransactionExecuted events raised by the NativeMetaTransaction contract.
type NativeMetaTransactionMetaTransactionExecutedIterator struct {
	Event *NativeMetaTransactionMetaTransactionExecuted // Event containing the contract specifics and raw log

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
func (it *NativeMetaTransactionMetaTransactionExecutedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NativeMetaTransactionMetaTransactionExecuted)
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
		it.Event = new(NativeMetaTransactionMetaTransactionExecuted)
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
func (it *NativeMetaTransactionMetaTransactionExecutedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NativeMetaTransactionMetaTransactionExecutedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NativeMetaTransactionMetaTransactionExecuted represents a MetaTransactionExecuted event raised by the NativeMetaTransaction contract.
type NativeMetaTransactionMetaTransactionExecuted struct {
	UserAddress       common.Address
	RelayerAddress    common.Address
	FunctionSignature []byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterMetaTransactionExecuted is a free log retrieval operation binding the contract event 0x5845892132946850460bff5a0083f71031bc5bf9aadcd40f1de79423eac9b10b.
//
// Solidity: event MetaTransactionExecuted(address userAddress, address relayerAddress, bytes functionSignature)
func (_NativeMetaTransaction *NativeMetaTransactionFilterer) FilterMetaTransactionExecuted(opts *bind.FilterOpts) (*NativeMetaTransactionMetaTransactionExecutedIterator, error) {

	logs, sub, err := _NativeMetaTransaction.contract.FilterLogs(opts, "MetaTransactionExecuted")
	if err != nil {
		return nil, err
	}
	return &NativeMetaTransactionMetaTransactionExecutedIterator{contract: _NativeMetaTransaction.contract, event: "MetaTransactionExecuted", logs: logs, sub: sub}, nil
}

// WatchMetaTransactionExecuted is a free log subscription operation binding the contract event 0x5845892132946850460bff5a0083f71031bc5bf9aadcd40f1de79423eac9b10b.
//
// Solidity: event MetaTransactionExecuted(address userAddress, address relayerAddress, bytes functionSignature)
func (_NativeMetaTransaction *NativeMetaTransactionFilterer) WatchMetaTransactionExecuted(opts *bind.WatchOpts, sink chan<- *NativeMetaTransactionMetaTransactionExecuted) (event.Subscription, error) {

	logs, sub, err := _NativeMetaTransaction.contract.WatchLogs(opts, "MetaTransactionExecuted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NativeMetaTransactionMetaTransactionExecuted)
				if err := _NativeMetaTransaction.contract.UnpackLog(event, "MetaTransactionExecuted", log); err != nil {
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

// ParseMetaTransactionExecuted is a log parse operation binding the contract event 0x5845892132946850460bff5a0083f71031bc5bf9aadcd40f1de79423eac9b10b.
//
// Solidity: event MetaTransactionExecuted(address userAddress, address relayerAddress, bytes functionSignature)
func (_NativeMetaTransaction *NativeMetaTransactionFilterer) ParseMetaTransactionExecuted(log types.Log) (*NativeMetaTransactionMetaTransactionExecuted, error) {
	event := new(NativeMetaTransactionMetaTransactionExecuted)
	if err := _NativeMetaTransaction.contract.UnpackLog(event, "MetaTransactionExecuted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
