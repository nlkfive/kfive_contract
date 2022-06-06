// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package IRegistrationList

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

// IRegistrationListMetaData contains all meta data concerning the IRegistrationList contract.
var IRegistrationListMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"AlreadyRegistered\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AlreadySelected\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidContract\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidSender\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidSlot\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotNLGGTHolder\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RaceNotExisted\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RewardIsExisted\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RewardIsNotExisted\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RewardNotExistedOrReceived\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"}],\"name\":\"TooEarly\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"}],\"name\":\"TooLate\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"nlggt\",\"type\":\"address\"}],\"name\":\"NlggtUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"requestId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"raceId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"randomness\",\"type\":\"uint256\"}],\"name\":\"ParticipantsSelected\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"race\",\"type\":\"address\"}],\"name\":\"RaceListUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"raceId\",\"type\":\"bytes32\"}],\"name\":\"RandomInProgress\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"slotId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"participant\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"raceId\",\"type\":\"bytes32\"}],\"name\":\"Registered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"raceId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nftRewardId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes1\",\"name\":\"resultIndex\",\"type\":\"bytes1\"}],\"name\":\"RewardAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"raceId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"slotId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nftRewardId\",\"type\":\"uint256\"}],\"name\":\"RewardReceived\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"raceId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nftRewardId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes1\",\"name\":\"resultIndex\",\"type\":\"bytes1\"}],\"name\":\"RewardRemoved\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"slotId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"raceId\",\"type\":\"bytes32\"}],\"name\":\"register\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"raceId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"nftRewardId\",\"type\":\"uint256\"},{\"internalType\":\"bytes1\",\"name\":\"resultIndex\",\"type\":\"bytes1\"}],\"name\":\"addRewardByTransfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"raceId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"nftRewardId\",\"type\":\"uint256\"},{\"internalType\":\"bytes1\",\"name\":\"resultIndex\",\"type\":\"bytes1\"},{\"internalType\":\"string\",\"name\":\"tokenURI\",\"type\":\"string\"}],\"name\":\"addRewardByMint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"raceId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"nftRewardId\",\"type\":\"uint256\"},{\"internalType\":\"bytes1\",\"name\":\"resultIndex\",\"type\":\"bytes1\"}],\"name\":\"removeReward\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"raceId\",\"type\":\"bytes32\"}],\"name\":\"selectParticipant\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"requestId\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"raceList\",\"type\":\"address\"}],\"name\":\"updateRaceAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"nlggt\",\"type\":\"address\"}],\"name\":\"updateNlggtAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"raceId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"slotId\",\"type\":\"uint256\"}],\"name\":\"receiveReward\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// IRegistrationListABI is the input ABI used to generate the binding from.
// Deprecated: Use IRegistrationListMetaData.ABI instead.
var IRegistrationListABI = IRegistrationListMetaData.ABI

// IRegistrationList is an auto generated Go binding around an Ethereum contract.
type IRegistrationList struct {
	IRegistrationListCaller     // Read-only binding to the contract
	IRegistrationListTransactor // Write-only binding to the contract
	IRegistrationListFilterer   // Log filterer for contract events
}

// IRegistrationListCaller is an auto generated read-only Go binding around an Ethereum contract.
type IRegistrationListCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IRegistrationListTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IRegistrationListTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IRegistrationListFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IRegistrationListFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IRegistrationListSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IRegistrationListSession struct {
	Contract     *IRegistrationList // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// IRegistrationListCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IRegistrationListCallerSession struct {
	Contract *IRegistrationListCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// IRegistrationListTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IRegistrationListTransactorSession struct {
	Contract     *IRegistrationListTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// IRegistrationListRaw is an auto generated low-level Go binding around an Ethereum contract.
type IRegistrationListRaw struct {
	Contract *IRegistrationList // Generic contract binding to access the raw methods on
}

// IRegistrationListCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IRegistrationListCallerRaw struct {
	Contract *IRegistrationListCaller // Generic read-only contract binding to access the raw methods on
}

// IRegistrationListTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IRegistrationListTransactorRaw struct {
	Contract *IRegistrationListTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIRegistrationList creates a new instance of IRegistrationList, bound to a specific deployed contract.
func NewIRegistrationList(address common.Address, backend bind.ContractBackend) (*IRegistrationList, error) {
	contract, err := bindIRegistrationList(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IRegistrationList{IRegistrationListCaller: IRegistrationListCaller{contract: contract}, IRegistrationListTransactor: IRegistrationListTransactor{contract: contract}, IRegistrationListFilterer: IRegistrationListFilterer{contract: contract}}, nil
}

// NewIRegistrationListCaller creates a new read-only instance of IRegistrationList, bound to a specific deployed contract.
func NewIRegistrationListCaller(address common.Address, caller bind.ContractCaller) (*IRegistrationListCaller, error) {
	contract, err := bindIRegistrationList(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IRegistrationListCaller{contract: contract}, nil
}

// NewIRegistrationListTransactor creates a new write-only instance of IRegistrationList, bound to a specific deployed contract.
func NewIRegistrationListTransactor(address common.Address, transactor bind.ContractTransactor) (*IRegistrationListTransactor, error) {
	contract, err := bindIRegistrationList(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IRegistrationListTransactor{contract: contract}, nil
}

// NewIRegistrationListFilterer creates a new log filterer instance of IRegistrationList, bound to a specific deployed contract.
func NewIRegistrationListFilterer(address common.Address, filterer bind.ContractFilterer) (*IRegistrationListFilterer, error) {
	contract, err := bindIRegistrationList(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IRegistrationListFilterer{contract: contract}, nil
}

// bindIRegistrationList binds a generic wrapper to an already deployed contract.
func bindIRegistrationList(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IRegistrationListABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IRegistrationList *IRegistrationListRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IRegistrationList.Contract.IRegistrationListCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IRegistrationList *IRegistrationListRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IRegistrationList.Contract.IRegistrationListTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IRegistrationList *IRegistrationListRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IRegistrationList.Contract.IRegistrationListTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IRegistrationList *IRegistrationListCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IRegistrationList.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IRegistrationList *IRegistrationListTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IRegistrationList.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IRegistrationList *IRegistrationListTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IRegistrationList.Contract.contract.Transact(opts, method, params...)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_IRegistrationList *IRegistrationListCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _IRegistrationList.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_IRegistrationList *IRegistrationListSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _IRegistrationList.Contract.SupportsInterface(&_IRegistrationList.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_IRegistrationList *IRegistrationListCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _IRegistrationList.Contract.SupportsInterface(&_IRegistrationList.CallOpts, interfaceId)
}

// AddRewardByMint is a paid mutator transaction binding the contract method 0x6bef76da.
//
// Solidity: function addRewardByMint(bytes32 raceId, uint256 nftRewardId, bytes1 resultIndex, string tokenURI) returns()
func (_IRegistrationList *IRegistrationListTransactor) AddRewardByMint(opts *bind.TransactOpts, raceId [32]byte, nftRewardId *big.Int, resultIndex [1]byte, tokenURI string) (*types.Transaction, error) {
	return _IRegistrationList.contract.Transact(opts, "addRewardByMint", raceId, nftRewardId, resultIndex, tokenURI)
}

// AddRewardByMint is a paid mutator transaction binding the contract method 0x6bef76da.
//
// Solidity: function addRewardByMint(bytes32 raceId, uint256 nftRewardId, bytes1 resultIndex, string tokenURI) returns()
func (_IRegistrationList *IRegistrationListSession) AddRewardByMint(raceId [32]byte, nftRewardId *big.Int, resultIndex [1]byte, tokenURI string) (*types.Transaction, error) {
	return _IRegistrationList.Contract.AddRewardByMint(&_IRegistrationList.TransactOpts, raceId, nftRewardId, resultIndex, tokenURI)
}

// AddRewardByMint is a paid mutator transaction binding the contract method 0x6bef76da.
//
// Solidity: function addRewardByMint(bytes32 raceId, uint256 nftRewardId, bytes1 resultIndex, string tokenURI) returns()
func (_IRegistrationList *IRegistrationListTransactorSession) AddRewardByMint(raceId [32]byte, nftRewardId *big.Int, resultIndex [1]byte, tokenURI string) (*types.Transaction, error) {
	return _IRegistrationList.Contract.AddRewardByMint(&_IRegistrationList.TransactOpts, raceId, nftRewardId, resultIndex, tokenURI)
}

// AddRewardByTransfer is a paid mutator transaction binding the contract method 0x81357309.
//
// Solidity: function addRewardByTransfer(bytes32 raceId, uint256 nftRewardId, bytes1 resultIndex) returns()
func (_IRegistrationList *IRegistrationListTransactor) AddRewardByTransfer(opts *bind.TransactOpts, raceId [32]byte, nftRewardId *big.Int, resultIndex [1]byte) (*types.Transaction, error) {
	return _IRegistrationList.contract.Transact(opts, "addRewardByTransfer", raceId, nftRewardId, resultIndex)
}

// AddRewardByTransfer is a paid mutator transaction binding the contract method 0x81357309.
//
// Solidity: function addRewardByTransfer(bytes32 raceId, uint256 nftRewardId, bytes1 resultIndex) returns()
func (_IRegistrationList *IRegistrationListSession) AddRewardByTransfer(raceId [32]byte, nftRewardId *big.Int, resultIndex [1]byte) (*types.Transaction, error) {
	return _IRegistrationList.Contract.AddRewardByTransfer(&_IRegistrationList.TransactOpts, raceId, nftRewardId, resultIndex)
}

// AddRewardByTransfer is a paid mutator transaction binding the contract method 0x81357309.
//
// Solidity: function addRewardByTransfer(bytes32 raceId, uint256 nftRewardId, bytes1 resultIndex) returns()
func (_IRegistrationList *IRegistrationListTransactorSession) AddRewardByTransfer(raceId [32]byte, nftRewardId *big.Int, resultIndex [1]byte) (*types.Transaction, error) {
	return _IRegistrationList.Contract.AddRewardByTransfer(&_IRegistrationList.TransactOpts, raceId, nftRewardId, resultIndex)
}

// ReceiveReward is a paid mutator transaction binding the contract method 0x1a0fc715.
//
// Solidity: function receiveReward(bytes32 raceId, uint256 slotId) returns()
func (_IRegistrationList *IRegistrationListTransactor) ReceiveReward(opts *bind.TransactOpts, raceId [32]byte, slotId *big.Int) (*types.Transaction, error) {
	return _IRegistrationList.contract.Transact(opts, "receiveReward", raceId, slotId)
}

// ReceiveReward is a paid mutator transaction binding the contract method 0x1a0fc715.
//
// Solidity: function receiveReward(bytes32 raceId, uint256 slotId) returns()
func (_IRegistrationList *IRegistrationListSession) ReceiveReward(raceId [32]byte, slotId *big.Int) (*types.Transaction, error) {
	return _IRegistrationList.Contract.ReceiveReward(&_IRegistrationList.TransactOpts, raceId, slotId)
}

// ReceiveReward is a paid mutator transaction binding the contract method 0x1a0fc715.
//
// Solidity: function receiveReward(bytes32 raceId, uint256 slotId) returns()
func (_IRegistrationList *IRegistrationListTransactorSession) ReceiveReward(raceId [32]byte, slotId *big.Int) (*types.Transaction, error) {
	return _IRegistrationList.Contract.ReceiveReward(&_IRegistrationList.TransactOpts, raceId, slotId)
}

// Register is a paid mutator transaction binding the contract method 0x2965d809.
//
// Solidity: function register(uint256 slotId, bytes32 raceId) returns()
func (_IRegistrationList *IRegistrationListTransactor) Register(opts *bind.TransactOpts, slotId *big.Int, raceId [32]byte) (*types.Transaction, error) {
	return _IRegistrationList.contract.Transact(opts, "register", slotId, raceId)
}

// Register is a paid mutator transaction binding the contract method 0x2965d809.
//
// Solidity: function register(uint256 slotId, bytes32 raceId) returns()
func (_IRegistrationList *IRegistrationListSession) Register(slotId *big.Int, raceId [32]byte) (*types.Transaction, error) {
	return _IRegistrationList.Contract.Register(&_IRegistrationList.TransactOpts, slotId, raceId)
}

// Register is a paid mutator transaction binding the contract method 0x2965d809.
//
// Solidity: function register(uint256 slotId, bytes32 raceId) returns()
func (_IRegistrationList *IRegistrationListTransactorSession) Register(slotId *big.Int, raceId [32]byte) (*types.Transaction, error) {
	return _IRegistrationList.Contract.Register(&_IRegistrationList.TransactOpts, slotId, raceId)
}

// RemoveReward is a paid mutator transaction binding the contract method 0x0fa445f5.
//
// Solidity: function removeReward(bytes32 raceId, uint256 nftRewardId, bytes1 resultIndex) returns()
func (_IRegistrationList *IRegistrationListTransactor) RemoveReward(opts *bind.TransactOpts, raceId [32]byte, nftRewardId *big.Int, resultIndex [1]byte) (*types.Transaction, error) {
	return _IRegistrationList.contract.Transact(opts, "removeReward", raceId, nftRewardId, resultIndex)
}

// RemoveReward is a paid mutator transaction binding the contract method 0x0fa445f5.
//
// Solidity: function removeReward(bytes32 raceId, uint256 nftRewardId, bytes1 resultIndex) returns()
func (_IRegistrationList *IRegistrationListSession) RemoveReward(raceId [32]byte, nftRewardId *big.Int, resultIndex [1]byte) (*types.Transaction, error) {
	return _IRegistrationList.Contract.RemoveReward(&_IRegistrationList.TransactOpts, raceId, nftRewardId, resultIndex)
}

// RemoveReward is a paid mutator transaction binding the contract method 0x0fa445f5.
//
// Solidity: function removeReward(bytes32 raceId, uint256 nftRewardId, bytes1 resultIndex) returns()
func (_IRegistrationList *IRegistrationListTransactorSession) RemoveReward(raceId [32]byte, nftRewardId *big.Int, resultIndex [1]byte) (*types.Transaction, error) {
	return _IRegistrationList.Contract.RemoveReward(&_IRegistrationList.TransactOpts, raceId, nftRewardId, resultIndex)
}

// SelectParticipant is a paid mutator transaction binding the contract method 0x7d0f79b7.
//
// Solidity: function selectParticipant(bytes32 raceId) returns(uint256 requestId)
func (_IRegistrationList *IRegistrationListTransactor) SelectParticipant(opts *bind.TransactOpts, raceId [32]byte) (*types.Transaction, error) {
	return _IRegistrationList.contract.Transact(opts, "selectParticipant", raceId)
}

// SelectParticipant is a paid mutator transaction binding the contract method 0x7d0f79b7.
//
// Solidity: function selectParticipant(bytes32 raceId) returns(uint256 requestId)
func (_IRegistrationList *IRegistrationListSession) SelectParticipant(raceId [32]byte) (*types.Transaction, error) {
	return _IRegistrationList.Contract.SelectParticipant(&_IRegistrationList.TransactOpts, raceId)
}

// SelectParticipant is a paid mutator transaction binding the contract method 0x7d0f79b7.
//
// Solidity: function selectParticipant(bytes32 raceId) returns(uint256 requestId)
func (_IRegistrationList *IRegistrationListTransactorSession) SelectParticipant(raceId [32]byte) (*types.Transaction, error) {
	return _IRegistrationList.Contract.SelectParticipant(&_IRegistrationList.TransactOpts, raceId)
}

// UpdateNlggtAddress is a paid mutator transaction binding the contract method 0xd6641f06.
//
// Solidity: function updateNlggtAddress(address nlggt) returns()
func (_IRegistrationList *IRegistrationListTransactor) UpdateNlggtAddress(opts *bind.TransactOpts, nlggt common.Address) (*types.Transaction, error) {
	return _IRegistrationList.contract.Transact(opts, "updateNlggtAddress", nlggt)
}

// UpdateNlggtAddress is a paid mutator transaction binding the contract method 0xd6641f06.
//
// Solidity: function updateNlggtAddress(address nlggt) returns()
func (_IRegistrationList *IRegistrationListSession) UpdateNlggtAddress(nlggt common.Address) (*types.Transaction, error) {
	return _IRegistrationList.Contract.UpdateNlggtAddress(&_IRegistrationList.TransactOpts, nlggt)
}

// UpdateNlggtAddress is a paid mutator transaction binding the contract method 0xd6641f06.
//
// Solidity: function updateNlggtAddress(address nlggt) returns()
func (_IRegistrationList *IRegistrationListTransactorSession) UpdateNlggtAddress(nlggt common.Address) (*types.Transaction, error) {
	return _IRegistrationList.Contract.UpdateNlggtAddress(&_IRegistrationList.TransactOpts, nlggt)
}

// UpdateRaceAddress is a paid mutator transaction binding the contract method 0xd7c49938.
//
// Solidity: function updateRaceAddress(address raceList) returns()
func (_IRegistrationList *IRegistrationListTransactor) UpdateRaceAddress(opts *bind.TransactOpts, raceList common.Address) (*types.Transaction, error) {
	return _IRegistrationList.contract.Transact(opts, "updateRaceAddress", raceList)
}

// UpdateRaceAddress is a paid mutator transaction binding the contract method 0xd7c49938.
//
// Solidity: function updateRaceAddress(address raceList) returns()
func (_IRegistrationList *IRegistrationListSession) UpdateRaceAddress(raceList common.Address) (*types.Transaction, error) {
	return _IRegistrationList.Contract.UpdateRaceAddress(&_IRegistrationList.TransactOpts, raceList)
}

// UpdateRaceAddress is a paid mutator transaction binding the contract method 0xd7c49938.
//
// Solidity: function updateRaceAddress(address raceList) returns()
func (_IRegistrationList *IRegistrationListTransactorSession) UpdateRaceAddress(raceList common.Address) (*types.Transaction, error) {
	return _IRegistrationList.Contract.UpdateRaceAddress(&_IRegistrationList.TransactOpts, raceList)
}

// IRegistrationListNlggtUpdatedIterator is returned from FilterNlggtUpdated and is used to iterate over the raw logs and unpacked data for NlggtUpdated events raised by the IRegistrationList contract.
type IRegistrationListNlggtUpdatedIterator struct {
	Event *IRegistrationListNlggtUpdated // Event containing the contract specifics and raw log

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
func (it *IRegistrationListNlggtUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IRegistrationListNlggtUpdated)
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
		it.Event = new(IRegistrationListNlggtUpdated)
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
func (it *IRegistrationListNlggtUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IRegistrationListNlggtUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IRegistrationListNlggtUpdated represents a NlggtUpdated event raised by the IRegistrationList contract.
type IRegistrationListNlggtUpdated struct {
	Nlggt common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterNlggtUpdated is a free log retrieval operation binding the contract event 0x5660a1e6d6115d84caa3bc631133061cbc77b46ff3be76d65178ae3243446b98.
//
// Solidity: event NlggtUpdated(address nlggt)
func (_IRegistrationList *IRegistrationListFilterer) FilterNlggtUpdated(opts *bind.FilterOpts) (*IRegistrationListNlggtUpdatedIterator, error) {

	logs, sub, err := _IRegistrationList.contract.FilterLogs(opts, "NlggtUpdated")
	if err != nil {
		return nil, err
	}
	return &IRegistrationListNlggtUpdatedIterator{contract: _IRegistrationList.contract, event: "NlggtUpdated", logs: logs, sub: sub}, nil
}

// WatchNlggtUpdated is a free log subscription operation binding the contract event 0x5660a1e6d6115d84caa3bc631133061cbc77b46ff3be76d65178ae3243446b98.
//
// Solidity: event NlggtUpdated(address nlggt)
func (_IRegistrationList *IRegistrationListFilterer) WatchNlggtUpdated(opts *bind.WatchOpts, sink chan<- *IRegistrationListNlggtUpdated) (event.Subscription, error) {

	logs, sub, err := _IRegistrationList.contract.WatchLogs(opts, "NlggtUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IRegistrationListNlggtUpdated)
				if err := _IRegistrationList.contract.UnpackLog(event, "NlggtUpdated", log); err != nil {
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

// ParseNlggtUpdated is a log parse operation binding the contract event 0x5660a1e6d6115d84caa3bc631133061cbc77b46ff3be76d65178ae3243446b98.
//
// Solidity: event NlggtUpdated(address nlggt)
func (_IRegistrationList *IRegistrationListFilterer) ParseNlggtUpdated(log types.Log) (*IRegistrationListNlggtUpdated, error) {
	event := new(IRegistrationListNlggtUpdated)
	if err := _IRegistrationList.contract.UnpackLog(event, "NlggtUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IRegistrationListParticipantsSelectedIterator is returned from FilterParticipantsSelected and is used to iterate over the raw logs and unpacked data for ParticipantsSelected events raised by the IRegistrationList contract.
type IRegistrationListParticipantsSelectedIterator struct {
	Event *IRegistrationListParticipantsSelected // Event containing the contract specifics and raw log

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
func (it *IRegistrationListParticipantsSelectedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IRegistrationListParticipantsSelected)
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
		it.Event = new(IRegistrationListParticipantsSelected)
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
func (it *IRegistrationListParticipantsSelectedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IRegistrationListParticipantsSelectedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IRegistrationListParticipantsSelected represents a ParticipantsSelected event raised by the IRegistrationList contract.
type IRegistrationListParticipantsSelected struct {
	RequestId  *big.Int
	RaceId     [32]byte
	Randomness *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterParticipantsSelected is a free log retrieval operation binding the contract event 0x5a1c9864d9d35edf19fb83e7a07ad46d28d80307bd248fb165e8814a1d1cfe2b.
//
// Solidity: event ParticipantsSelected(uint256 requestId, bytes32 raceId, uint256 randomness)
func (_IRegistrationList *IRegistrationListFilterer) FilterParticipantsSelected(opts *bind.FilterOpts) (*IRegistrationListParticipantsSelectedIterator, error) {

	logs, sub, err := _IRegistrationList.contract.FilterLogs(opts, "ParticipantsSelected")
	if err != nil {
		return nil, err
	}
	return &IRegistrationListParticipantsSelectedIterator{contract: _IRegistrationList.contract, event: "ParticipantsSelected", logs: logs, sub: sub}, nil
}

// WatchParticipantsSelected is a free log subscription operation binding the contract event 0x5a1c9864d9d35edf19fb83e7a07ad46d28d80307bd248fb165e8814a1d1cfe2b.
//
// Solidity: event ParticipantsSelected(uint256 requestId, bytes32 raceId, uint256 randomness)
func (_IRegistrationList *IRegistrationListFilterer) WatchParticipantsSelected(opts *bind.WatchOpts, sink chan<- *IRegistrationListParticipantsSelected) (event.Subscription, error) {

	logs, sub, err := _IRegistrationList.contract.WatchLogs(opts, "ParticipantsSelected")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IRegistrationListParticipantsSelected)
				if err := _IRegistrationList.contract.UnpackLog(event, "ParticipantsSelected", log); err != nil {
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
func (_IRegistrationList *IRegistrationListFilterer) ParseParticipantsSelected(log types.Log) (*IRegistrationListParticipantsSelected, error) {
	event := new(IRegistrationListParticipantsSelected)
	if err := _IRegistrationList.contract.UnpackLog(event, "ParticipantsSelected", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IRegistrationListRaceListUpdatedIterator is returned from FilterRaceListUpdated and is used to iterate over the raw logs and unpacked data for RaceListUpdated events raised by the IRegistrationList contract.
type IRegistrationListRaceListUpdatedIterator struct {
	Event *IRegistrationListRaceListUpdated // Event containing the contract specifics and raw log

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
func (it *IRegistrationListRaceListUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IRegistrationListRaceListUpdated)
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
		it.Event = new(IRegistrationListRaceListUpdated)
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
func (it *IRegistrationListRaceListUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IRegistrationListRaceListUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IRegistrationListRaceListUpdated represents a RaceListUpdated event raised by the IRegistrationList contract.
type IRegistrationListRaceListUpdated struct {
	Race common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterRaceListUpdated is a free log retrieval operation binding the contract event 0x7a659342c2b03e1c14729e3f1c86f414756d13bd8ff0636713cc22246182b42d.
//
// Solidity: event RaceListUpdated(address race)
func (_IRegistrationList *IRegistrationListFilterer) FilterRaceListUpdated(opts *bind.FilterOpts) (*IRegistrationListRaceListUpdatedIterator, error) {

	logs, sub, err := _IRegistrationList.contract.FilterLogs(opts, "RaceListUpdated")
	if err != nil {
		return nil, err
	}
	return &IRegistrationListRaceListUpdatedIterator{contract: _IRegistrationList.contract, event: "RaceListUpdated", logs: logs, sub: sub}, nil
}

// WatchRaceListUpdated is a free log subscription operation binding the contract event 0x7a659342c2b03e1c14729e3f1c86f414756d13bd8ff0636713cc22246182b42d.
//
// Solidity: event RaceListUpdated(address race)
func (_IRegistrationList *IRegistrationListFilterer) WatchRaceListUpdated(opts *bind.WatchOpts, sink chan<- *IRegistrationListRaceListUpdated) (event.Subscription, error) {

	logs, sub, err := _IRegistrationList.contract.WatchLogs(opts, "RaceListUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IRegistrationListRaceListUpdated)
				if err := _IRegistrationList.contract.UnpackLog(event, "RaceListUpdated", log); err != nil {
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

// ParseRaceListUpdated is a log parse operation binding the contract event 0x7a659342c2b03e1c14729e3f1c86f414756d13bd8ff0636713cc22246182b42d.
//
// Solidity: event RaceListUpdated(address race)
func (_IRegistrationList *IRegistrationListFilterer) ParseRaceListUpdated(log types.Log) (*IRegistrationListRaceListUpdated, error) {
	event := new(IRegistrationListRaceListUpdated)
	if err := _IRegistrationList.contract.UnpackLog(event, "RaceListUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IRegistrationListRandomInProgressIterator is returned from FilterRandomInProgress and is used to iterate over the raw logs and unpacked data for RandomInProgress events raised by the IRegistrationList contract.
type IRegistrationListRandomInProgressIterator struct {
	Event *IRegistrationListRandomInProgress // Event containing the contract specifics and raw log

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
func (it *IRegistrationListRandomInProgressIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IRegistrationListRandomInProgress)
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
		it.Event = new(IRegistrationListRandomInProgress)
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
func (it *IRegistrationListRandomInProgressIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IRegistrationListRandomInProgressIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IRegistrationListRandomInProgress represents a RandomInProgress event raised by the IRegistrationList contract.
type IRegistrationListRandomInProgress struct {
	RaceId [32]byte
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRandomInProgress is a free log retrieval operation binding the contract event 0x018dc112e9dad7765a03056b4091d237454ee37ae27aab3c3abd16d428900a13.
//
// Solidity: event RandomInProgress(bytes32 raceId)
func (_IRegistrationList *IRegistrationListFilterer) FilterRandomInProgress(opts *bind.FilterOpts) (*IRegistrationListRandomInProgressIterator, error) {

	logs, sub, err := _IRegistrationList.contract.FilterLogs(opts, "RandomInProgress")
	if err != nil {
		return nil, err
	}
	return &IRegistrationListRandomInProgressIterator{contract: _IRegistrationList.contract, event: "RandomInProgress", logs: logs, sub: sub}, nil
}

// WatchRandomInProgress is a free log subscription operation binding the contract event 0x018dc112e9dad7765a03056b4091d237454ee37ae27aab3c3abd16d428900a13.
//
// Solidity: event RandomInProgress(bytes32 raceId)
func (_IRegistrationList *IRegistrationListFilterer) WatchRandomInProgress(opts *bind.WatchOpts, sink chan<- *IRegistrationListRandomInProgress) (event.Subscription, error) {

	logs, sub, err := _IRegistrationList.contract.WatchLogs(opts, "RandomInProgress")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IRegistrationListRandomInProgress)
				if err := _IRegistrationList.contract.UnpackLog(event, "RandomInProgress", log); err != nil {
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
func (_IRegistrationList *IRegistrationListFilterer) ParseRandomInProgress(log types.Log) (*IRegistrationListRandomInProgress, error) {
	event := new(IRegistrationListRandomInProgress)
	if err := _IRegistrationList.contract.UnpackLog(event, "RandomInProgress", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IRegistrationListRegisteredIterator is returned from FilterRegistered and is used to iterate over the raw logs and unpacked data for Registered events raised by the IRegistrationList contract.
type IRegistrationListRegisteredIterator struct {
	Event *IRegistrationListRegistered // Event containing the contract specifics and raw log

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
func (it *IRegistrationListRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IRegistrationListRegistered)
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
		it.Event = new(IRegistrationListRegistered)
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
func (it *IRegistrationListRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IRegistrationListRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IRegistrationListRegistered represents a Registered event raised by the IRegistrationList contract.
type IRegistrationListRegistered struct {
	SlotId      *big.Int
	Participant common.Address
	RaceId      [32]byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterRegistered is a free log retrieval operation binding the contract event 0xd62c7bd2779dc6ebe426bed623399b1fc1d190c063eadefcffffb1efe6f253ef.
//
// Solidity: event Registered(uint256 slotId, address participant, bytes32 raceId)
func (_IRegistrationList *IRegistrationListFilterer) FilterRegistered(opts *bind.FilterOpts) (*IRegistrationListRegisteredIterator, error) {

	logs, sub, err := _IRegistrationList.contract.FilterLogs(opts, "Registered")
	if err != nil {
		return nil, err
	}
	return &IRegistrationListRegisteredIterator{contract: _IRegistrationList.contract, event: "Registered", logs: logs, sub: sub}, nil
}

// WatchRegistered is a free log subscription operation binding the contract event 0xd62c7bd2779dc6ebe426bed623399b1fc1d190c063eadefcffffb1efe6f253ef.
//
// Solidity: event Registered(uint256 slotId, address participant, bytes32 raceId)
func (_IRegistrationList *IRegistrationListFilterer) WatchRegistered(opts *bind.WatchOpts, sink chan<- *IRegistrationListRegistered) (event.Subscription, error) {

	logs, sub, err := _IRegistrationList.contract.WatchLogs(opts, "Registered")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IRegistrationListRegistered)
				if err := _IRegistrationList.contract.UnpackLog(event, "Registered", log); err != nil {
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
func (_IRegistrationList *IRegistrationListFilterer) ParseRegistered(log types.Log) (*IRegistrationListRegistered, error) {
	event := new(IRegistrationListRegistered)
	if err := _IRegistrationList.contract.UnpackLog(event, "Registered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IRegistrationListRewardAddedIterator is returned from FilterRewardAdded and is used to iterate over the raw logs and unpacked data for RewardAdded events raised by the IRegistrationList contract.
type IRegistrationListRewardAddedIterator struct {
	Event *IRegistrationListRewardAdded // Event containing the contract specifics and raw log

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
func (it *IRegistrationListRewardAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IRegistrationListRewardAdded)
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
		it.Event = new(IRegistrationListRewardAdded)
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
func (it *IRegistrationListRewardAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IRegistrationListRewardAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IRegistrationListRewardAdded represents a RewardAdded event raised by the IRegistrationList contract.
type IRegistrationListRewardAdded struct {
	RaceId      [32]byte
	NftRewardId *big.Int
	ResultIndex [1]byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterRewardAdded is a free log retrieval operation binding the contract event 0x2b942e265bd6d4ca5617a327f9a60e1a85f4351aaa0be2a84cda1cf9d32c8c04.
//
// Solidity: event RewardAdded(bytes32 raceId, uint256 nftRewardId, bytes1 resultIndex)
func (_IRegistrationList *IRegistrationListFilterer) FilterRewardAdded(opts *bind.FilterOpts) (*IRegistrationListRewardAddedIterator, error) {

	logs, sub, err := _IRegistrationList.contract.FilterLogs(opts, "RewardAdded")
	if err != nil {
		return nil, err
	}
	return &IRegistrationListRewardAddedIterator{contract: _IRegistrationList.contract, event: "RewardAdded", logs: logs, sub: sub}, nil
}

// WatchRewardAdded is a free log subscription operation binding the contract event 0x2b942e265bd6d4ca5617a327f9a60e1a85f4351aaa0be2a84cda1cf9d32c8c04.
//
// Solidity: event RewardAdded(bytes32 raceId, uint256 nftRewardId, bytes1 resultIndex)
func (_IRegistrationList *IRegistrationListFilterer) WatchRewardAdded(opts *bind.WatchOpts, sink chan<- *IRegistrationListRewardAdded) (event.Subscription, error) {

	logs, sub, err := _IRegistrationList.contract.WatchLogs(opts, "RewardAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IRegistrationListRewardAdded)
				if err := _IRegistrationList.contract.UnpackLog(event, "RewardAdded", log); err != nil {
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
func (_IRegistrationList *IRegistrationListFilterer) ParseRewardAdded(log types.Log) (*IRegistrationListRewardAdded, error) {
	event := new(IRegistrationListRewardAdded)
	if err := _IRegistrationList.contract.UnpackLog(event, "RewardAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IRegistrationListRewardReceivedIterator is returned from FilterRewardReceived and is used to iterate over the raw logs and unpacked data for RewardReceived events raised by the IRegistrationList contract.
type IRegistrationListRewardReceivedIterator struct {
	Event *IRegistrationListRewardReceived // Event containing the contract specifics and raw log

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
func (it *IRegistrationListRewardReceivedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IRegistrationListRewardReceived)
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
		it.Event = new(IRegistrationListRewardReceived)
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
func (it *IRegistrationListRewardReceivedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IRegistrationListRewardReceivedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IRegistrationListRewardReceived represents a RewardReceived event raised by the IRegistrationList contract.
type IRegistrationListRewardReceived struct {
	RaceId      [32]byte
	SlotId      *big.Int
	NftRewardId *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterRewardReceived is a free log retrieval operation binding the contract event 0xe2ece170b35ae79cbe45ffdaf4931b8bc45056da5d3e9aa4aea67fcdf307929a.
//
// Solidity: event RewardReceived(bytes32 raceId, uint256 slotId, uint256 nftRewardId)
func (_IRegistrationList *IRegistrationListFilterer) FilterRewardReceived(opts *bind.FilterOpts) (*IRegistrationListRewardReceivedIterator, error) {

	logs, sub, err := _IRegistrationList.contract.FilterLogs(opts, "RewardReceived")
	if err != nil {
		return nil, err
	}
	return &IRegistrationListRewardReceivedIterator{contract: _IRegistrationList.contract, event: "RewardReceived", logs: logs, sub: sub}, nil
}

// WatchRewardReceived is a free log subscription operation binding the contract event 0xe2ece170b35ae79cbe45ffdaf4931b8bc45056da5d3e9aa4aea67fcdf307929a.
//
// Solidity: event RewardReceived(bytes32 raceId, uint256 slotId, uint256 nftRewardId)
func (_IRegistrationList *IRegistrationListFilterer) WatchRewardReceived(opts *bind.WatchOpts, sink chan<- *IRegistrationListRewardReceived) (event.Subscription, error) {

	logs, sub, err := _IRegistrationList.contract.WatchLogs(opts, "RewardReceived")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IRegistrationListRewardReceived)
				if err := _IRegistrationList.contract.UnpackLog(event, "RewardReceived", log); err != nil {
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
func (_IRegistrationList *IRegistrationListFilterer) ParseRewardReceived(log types.Log) (*IRegistrationListRewardReceived, error) {
	event := new(IRegistrationListRewardReceived)
	if err := _IRegistrationList.contract.UnpackLog(event, "RewardReceived", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IRegistrationListRewardRemovedIterator is returned from FilterRewardRemoved and is used to iterate over the raw logs and unpacked data for RewardRemoved events raised by the IRegistrationList contract.
type IRegistrationListRewardRemovedIterator struct {
	Event *IRegistrationListRewardRemoved // Event containing the contract specifics and raw log

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
func (it *IRegistrationListRewardRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IRegistrationListRewardRemoved)
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
		it.Event = new(IRegistrationListRewardRemoved)
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
func (it *IRegistrationListRewardRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IRegistrationListRewardRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IRegistrationListRewardRemoved represents a RewardRemoved event raised by the IRegistrationList contract.
type IRegistrationListRewardRemoved struct {
	RaceId      [32]byte
	NftRewardId *big.Int
	ResultIndex [1]byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterRewardRemoved is a free log retrieval operation binding the contract event 0xddda89b96dee97fe43f9803c50584f6dd667bedc8b02a7c75407e6906bf31ead.
//
// Solidity: event RewardRemoved(bytes32 raceId, uint256 nftRewardId, bytes1 resultIndex)
func (_IRegistrationList *IRegistrationListFilterer) FilterRewardRemoved(opts *bind.FilterOpts) (*IRegistrationListRewardRemovedIterator, error) {

	logs, sub, err := _IRegistrationList.contract.FilterLogs(opts, "RewardRemoved")
	if err != nil {
		return nil, err
	}
	return &IRegistrationListRewardRemovedIterator{contract: _IRegistrationList.contract, event: "RewardRemoved", logs: logs, sub: sub}, nil
}

// WatchRewardRemoved is a free log subscription operation binding the contract event 0xddda89b96dee97fe43f9803c50584f6dd667bedc8b02a7c75407e6906bf31ead.
//
// Solidity: event RewardRemoved(bytes32 raceId, uint256 nftRewardId, bytes1 resultIndex)
func (_IRegistrationList *IRegistrationListFilterer) WatchRewardRemoved(opts *bind.WatchOpts, sink chan<- *IRegistrationListRewardRemoved) (event.Subscription, error) {

	logs, sub, err := _IRegistrationList.contract.WatchLogs(opts, "RewardRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IRegistrationListRewardRemoved)
				if err := _IRegistrationList.contract.UnpackLog(event, "RewardRemoved", log); err != nil {
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
func (_IRegistrationList *IRegistrationListFilterer) ParseRewardRemoved(log types.Log) (*IRegistrationListRewardRemoved, error) {
	event := new(IRegistrationListRewardRemoved)
	if err := _IRegistrationList.contract.UnpackLog(event, "RewardRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
