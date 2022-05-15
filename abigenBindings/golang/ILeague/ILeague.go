// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ILeague

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

// ILeagueMetaData contains all meta data concerning the ILeague contract.
var ILeagueMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"AlreadyRegistered\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CannotCancel\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_currentRaceNo\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalRace\",\"type\":\"uint256\"}],\"name\":\"CannotCreateMoreRace\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidContract\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidRegister\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidSender\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidSlot\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotEndYet\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RaceExisted\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RaceNotExisted\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RewardIsExisted\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RewardIsNotExisted\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"}],\"name\":\"TooEarly\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"}],\"name\":\"TooLate\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"}],\"name\":\"RaceCancelled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"slots\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"startAt\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"raceNo\",\"type\":\"uint256\"}],\"name\":\"RaceCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"result\",\"type\":\"bytes32\"}],\"name\":\"RaceResultUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"slotId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"participant\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"raceId\",\"type\":\"bytes32\"}],\"name\":\"Registered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"raceId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nftRewardId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"winnerIndex\",\"type\":\"uint256\"}],\"name\":\"RewardAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"raceId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"slotId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nftRewardId\",\"type\":\"uint256\"}],\"name\":\"RewardReceived\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"raceId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nftRewardId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"winnerIndex\",\"type\":\"uint256\"}],\"name\":\"RewardRemoved\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"slots\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startAt\",\"type\":\"uint256\"}],\"name\":\"createRace\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"raceId\",\"type\":\"bytes32\"}],\"name\":\"raceInfo\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"slots\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startAt\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"result\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"register\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"raceId\",\"type\":\"bytes32\"}],\"name\":\"registeredSlot\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"register\",\"type\":\"address\"}],\"name\":\"getTotalScore\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ended\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"leagueName\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"raceId\",\"type\":\"bytes32\"}],\"name\":\"cancelRace\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"raceId\",\"type\":\"bytes32\"},{\"internalType\":\"uint8\",\"name\":\"slotId\",\"type\":\"uint8\"}],\"name\":\"registerRace\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"raceId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"result\",\"type\":\"bytes32\"}],\"name\":\"updateRaceResult\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"raceId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"nftRewardId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"winnerIndex\",\"type\":\"uint256\"}],\"name\":\"addRewardByTransfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"raceId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"nftRewardId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"winnerIndex\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"tokenURI\",\"type\":\"string\"}],\"name\":\"addRewardByMint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"raceId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"winnerIndex\",\"type\":\"uint256\"}],\"name\":\"removeReward\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"raceId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"slotId\",\"type\":\"uint256\"}],\"name\":\"receiveReward\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// ILeagueABI is the input ABI used to generate the binding from.
// Deprecated: Use ILeagueMetaData.ABI instead.
var ILeagueABI = ILeagueMetaData.ABI

// ILeague is an auto generated Go binding around an Ethereum contract.
type ILeague struct {
	ILeagueCaller     // Read-only binding to the contract
	ILeagueTransactor // Write-only binding to the contract
	ILeagueFilterer   // Log filterer for contract events
}

// ILeagueCaller is an auto generated read-only Go binding around an Ethereum contract.
type ILeagueCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ILeagueTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ILeagueTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ILeagueFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ILeagueFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ILeagueSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ILeagueSession struct {
	Contract     *ILeague          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ILeagueCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ILeagueCallerSession struct {
	Contract *ILeagueCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// ILeagueTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ILeagueTransactorSession struct {
	Contract     *ILeagueTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// ILeagueRaw is an auto generated low-level Go binding around an Ethereum contract.
type ILeagueRaw struct {
	Contract *ILeague // Generic contract binding to access the raw methods on
}

// ILeagueCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ILeagueCallerRaw struct {
	Contract *ILeagueCaller // Generic read-only contract binding to access the raw methods on
}

// ILeagueTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ILeagueTransactorRaw struct {
	Contract *ILeagueTransactor // Generic write-only contract binding to access the raw methods on
}

// NewILeague creates a new instance of ILeague, bound to a specific deployed contract.
func NewILeague(address common.Address, backend bind.ContractBackend) (*ILeague, error) {
	contract, err := bindILeague(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ILeague{ILeagueCaller: ILeagueCaller{contract: contract}, ILeagueTransactor: ILeagueTransactor{contract: contract}, ILeagueFilterer: ILeagueFilterer{contract: contract}}, nil
}

// NewILeagueCaller creates a new read-only instance of ILeague, bound to a specific deployed contract.
func NewILeagueCaller(address common.Address, caller bind.ContractCaller) (*ILeagueCaller, error) {
	contract, err := bindILeague(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ILeagueCaller{contract: contract}, nil
}

// NewILeagueTransactor creates a new write-only instance of ILeague, bound to a specific deployed contract.
func NewILeagueTransactor(address common.Address, transactor bind.ContractTransactor) (*ILeagueTransactor, error) {
	contract, err := bindILeague(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ILeagueTransactor{contract: contract}, nil
}

// NewILeagueFilterer creates a new log filterer instance of ILeague, bound to a specific deployed contract.
func NewILeagueFilterer(address common.Address, filterer bind.ContractFilterer) (*ILeagueFilterer, error) {
	contract, err := bindILeague(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ILeagueFilterer{contract: contract}, nil
}

// bindILeague binds a generic wrapper to an already deployed contract.
func bindILeague(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ILeagueABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ILeague *ILeagueRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ILeague.Contract.ILeagueCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ILeague *ILeagueRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ILeague.Contract.ILeagueTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ILeague *ILeagueRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ILeague.Contract.ILeagueTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ILeague *ILeagueCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ILeague.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ILeague *ILeagueTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ILeague.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ILeague *ILeagueTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ILeague.Contract.contract.Transact(opts, method, params...)
}

// Ended is a free data retrieval call binding the contract method 0x12fa6feb.
//
// Solidity: function ended() view returns(bool)
func (_ILeague *ILeagueCaller) Ended(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _ILeague.contract.Call(opts, &out, "ended")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Ended is a free data retrieval call binding the contract method 0x12fa6feb.
//
// Solidity: function ended() view returns(bool)
func (_ILeague *ILeagueSession) Ended() (bool, error) {
	return _ILeague.Contract.Ended(&_ILeague.CallOpts)
}

// Ended is a free data retrieval call binding the contract method 0x12fa6feb.
//
// Solidity: function ended() view returns(bool)
func (_ILeague *ILeagueCallerSession) Ended() (bool, error) {
	return _ILeague.Contract.Ended(&_ILeague.CallOpts)
}

// GetTotalScore is a free data retrieval call binding the contract method 0xd283b3c5.
//
// Solidity: function getTotalScore(address register) view returns(uint256)
func (_ILeague *ILeagueCaller) GetTotalScore(opts *bind.CallOpts, register common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ILeague.contract.Call(opts, &out, "getTotalScore", register)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTotalScore is a free data retrieval call binding the contract method 0xd283b3c5.
//
// Solidity: function getTotalScore(address register) view returns(uint256)
func (_ILeague *ILeagueSession) GetTotalScore(register common.Address) (*big.Int, error) {
	return _ILeague.Contract.GetTotalScore(&_ILeague.CallOpts, register)
}

// GetTotalScore is a free data retrieval call binding the contract method 0xd283b3c5.
//
// Solidity: function getTotalScore(address register) view returns(uint256)
func (_ILeague *ILeagueCallerSession) GetTotalScore(register common.Address) (*big.Int, error) {
	return _ILeague.Contract.GetTotalScore(&_ILeague.CallOpts, register)
}

// LeagueName is a free data retrieval call binding the contract method 0x841abf6e.
//
// Solidity: function leagueName() view returns(string)
func (_ILeague *ILeagueCaller) LeagueName(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ILeague.contract.Call(opts, &out, "leagueName")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// LeagueName is a free data retrieval call binding the contract method 0x841abf6e.
//
// Solidity: function leagueName() view returns(string)
func (_ILeague *ILeagueSession) LeagueName() (string, error) {
	return _ILeague.Contract.LeagueName(&_ILeague.CallOpts)
}

// LeagueName is a free data retrieval call binding the contract method 0x841abf6e.
//
// Solidity: function leagueName() view returns(string)
func (_ILeague *ILeagueCallerSession) LeagueName() (string, error) {
	return _ILeague.Contract.LeagueName(&_ILeague.CallOpts)
}

// RaceInfo is a free data retrieval call binding the contract method 0xef8597fd.
//
// Solidity: function raceInfo(bytes32 raceId) view returns(uint256 slots, uint256 startAt, bytes32 result)
func (_ILeague *ILeagueCaller) RaceInfo(opts *bind.CallOpts, raceId [32]byte) (struct {
	Slots   *big.Int
	StartAt *big.Int
	Result  [32]byte
}, error) {
	var out []interface{}
	err := _ILeague.contract.Call(opts, &out, "raceInfo", raceId)

	outstruct := new(struct {
		Slots   *big.Int
		StartAt *big.Int
		Result  [32]byte
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Slots = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.StartAt = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Result = *abi.ConvertType(out[2], new([32]byte)).(*[32]byte)

	return *outstruct, err

}

// RaceInfo is a free data retrieval call binding the contract method 0xef8597fd.
//
// Solidity: function raceInfo(bytes32 raceId) view returns(uint256 slots, uint256 startAt, bytes32 result)
func (_ILeague *ILeagueSession) RaceInfo(raceId [32]byte) (struct {
	Slots   *big.Int
	StartAt *big.Int
	Result  [32]byte
}, error) {
	return _ILeague.Contract.RaceInfo(&_ILeague.CallOpts, raceId)
}

// RaceInfo is a free data retrieval call binding the contract method 0xef8597fd.
//
// Solidity: function raceInfo(bytes32 raceId) view returns(uint256 slots, uint256 startAt, bytes32 result)
func (_ILeague *ILeagueCallerSession) RaceInfo(raceId [32]byte) (struct {
	Slots   *big.Int
	StartAt *big.Int
	Result  [32]byte
}, error) {
	return _ILeague.Contract.RaceInfo(&_ILeague.CallOpts, raceId)
}

// RegisteredSlot is a free data retrieval call binding the contract method 0x96fd6980.
//
// Solidity: function registeredSlot(address register, bytes32 raceId) view returns(uint256)
func (_ILeague *ILeagueCaller) RegisteredSlot(opts *bind.CallOpts, register common.Address, raceId [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _ILeague.contract.Call(opts, &out, "registeredSlot", register, raceId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RegisteredSlot is a free data retrieval call binding the contract method 0x96fd6980.
//
// Solidity: function registeredSlot(address register, bytes32 raceId) view returns(uint256)
func (_ILeague *ILeagueSession) RegisteredSlot(register common.Address, raceId [32]byte) (*big.Int, error) {
	return _ILeague.Contract.RegisteredSlot(&_ILeague.CallOpts, register, raceId)
}

// RegisteredSlot is a free data retrieval call binding the contract method 0x96fd6980.
//
// Solidity: function registeredSlot(address register, bytes32 raceId) view returns(uint256)
func (_ILeague *ILeagueCallerSession) RegisteredSlot(register common.Address, raceId [32]byte) (*big.Int, error) {
	return _ILeague.Contract.RegisteredSlot(&_ILeague.CallOpts, register, raceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_ILeague *ILeagueCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _ILeague.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_ILeague *ILeagueSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _ILeague.Contract.SupportsInterface(&_ILeague.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_ILeague *ILeagueCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _ILeague.Contract.SupportsInterface(&_ILeague.CallOpts, interfaceId)
}

// AddRewardByMint is a paid mutator transaction binding the contract method 0xa4c1564a.
//
// Solidity: function addRewardByMint(bytes32 raceId, uint256 nftRewardId, uint256 winnerIndex, string tokenURI) returns()
func (_ILeague *ILeagueTransactor) AddRewardByMint(opts *bind.TransactOpts, raceId [32]byte, nftRewardId *big.Int, winnerIndex *big.Int, tokenURI string) (*types.Transaction, error) {
	return _ILeague.contract.Transact(opts, "addRewardByMint", raceId, nftRewardId, winnerIndex, tokenURI)
}

// AddRewardByMint is a paid mutator transaction binding the contract method 0xa4c1564a.
//
// Solidity: function addRewardByMint(bytes32 raceId, uint256 nftRewardId, uint256 winnerIndex, string tokenURI) returns()
func (_ILeague *ILeagueSession) AddRewardByMint(raceId [32]byte, nftRewardId *big.Int, winnerIndex *big.Int, tokenURI string) (*types.Transaction, error) {
	return _ILeague.Contract.AddRewardByMint(&_ILeague.TransactOpts, raceId, nftRewardId, winnerIndex, tokenURI)
}

// AddRewardByMint is a paid mutator transaction binding the contract method 0xa4c1564a.
//
// Solidity: function addRewardByMint(bytes32 raceId, uint256 nftRewardId, uint256 winnerIndex, string tokenURI) returns()
func (_ILeague *ILeagueTransactorSession) AddRewardByMint(raceId [32]byte, nftRewardId *big.Int, winnerIndex *big.Int, tokenURI string) (*types.Transaction, error) {
	return _ILeague.Contract.AddRewardByMint(&_ILeague.TransactOpts, raceId, nftRewardId, winnerIndex, tokenURI)
}

// AddRewardByTransfer is a paid mutator transaction binding the contract method 0x21ad4e51.
//
// Solidity: function addRewardByTransfer(bytes32 raceId, uint256 nftRewardId, uint256 winnerIndex) returns()
func (_ILeague *ILeagueTransactor) AddRewardByTransfer(opts *bind.TransactOpts, raceId [32]byte, nftRewardId *big.Int, winnerIndex *big.Int) (*types.Transaction, error) {
	return _ILeague.contract.Transact(opts, "addRewardByTransfer", raceId, nftRewardId, winnerIndex)
}

// AddRewardByTransfer is a paid mutator transaction binding the contract method 0x21ad4e51.
//
// Solidity: function addRewardByTransfer(bytes32 raceId, uint256 nftRewardId, uint256 winnerIndex) returns()
func (_ILeague *ILeagueSession) AddRewardByTransfer(raceId [32]byte, nftRewardId *big.Int, winnerIndex *big.Int) (*types.Transaction, error) {
	return _ILeague.Contract.AddRewardByTransfer(&_ILeague.TransactOpts, raceId, nftRewardId, winnerIndex)
}

// AddRewardByTransfer is a paid mutator transaction binding the contract method 0x21ad4e51.
//
// Solidity: function addRewardByTransfer(bytes32 raceId, uint256 nftRewardId, uint256 winnerIndex) returns()
func (_ILeague *ILeagueTransactorSession) AddRewardByTransfer(raceId [32]byte, nftRewardId *big.Int, winnerIndex *big.Int) (*types.Transaction, error) {
	return _ILeague.Contract.AddRewardByTransfer(&_ILeague.TransactOpts, raceId, nftRewardId, winnerIndex)
}

// CancelRace is a paid mutator transaction binding the contract method 0xc05c107d.
//
// Solidity: function cancelRace(bytes32 raceId) returns()
func (_ILeague *ILeagueTransactor) CancelRace(opts *bind.TransactOpts, raceId [32]byte) (*types.Transaction, error) {
	return _ILeague.contract.Transact(opts, "cancelRace", raceId)
}

// CancelRace is a paid mutator transaction binding the contract method 0xc05c107d.
//
// Solidity: function cancelRace(bytes32 raceId) returns()
func (_ILeague *ILeagueSession) CancelRace(raceId [32]byte) (*types.Transaction, error) {
	return _ILeague.Contract.CancelRace(&_ILeague.TransactOpts, raceId)
}

// CancelRace is a paid mutator transaction binding the contract method 0xc05c107d.
//
// Solidity: function cancelRace(bytes32 raceId) returns()
func (_ILeague *ILeagueTransactorSession) CancelRace(raceId [32]byte) (*types.Transaction, error) {
	return _ILeague.Contract.CancelRace(&_ILeague.TransactOpts, raceId)
}

// CreateRace is a paid mutator transaction binding the contract method 0x420f0351.
//
// Solidity: function createRace(uint256 slots, uint256 startAt) returns(bytes32)
func (_ILeague *ILeagueTransactor) CreateRace(opts *bind.TransactOpts, slots *big.Int, startAt *big.Int) (*types.Transaction, error) {
	return _ILeague.contract.Transact(opts, "createRace", slots, startAt)
}

// CreateRace is a paid mutator transaction binding the contract method 0x420f0351.
//
// Solidity: function createRace(uint256 slots, uint256 startAt) returns(bytes32)
func (_ILeague *ILeagueSession) CreateRace(slots *big.Int, startAt *big.Int) (*types.Transaction, error) {
	return _ILeague.Contract.CreateRace(&_ILeague.TransactOpts, slots, startAt)
}

// CreateRace is a paid mutator transaction binding the contract method 0x420f0351.
//
// Solidity: function createRace(uint256 slots, uint256 startAt) returns(bytes32)
func (_ILeague *ILeagueTransactorSession) CreateRace(slots *big.Int, startAt *big.Int) (*types.Transaction, error) {
	return _ILeague.Contract.CreateRace(&_ILeague.TransactOpts, slots, startAt)
}

// ReceiveReward is a paid mutator transaction binding the contract method 0x1a0fc715.
//
// Solidity: function receiveReward(bytes32 raceId, uint256 slotId) returns()
func (_ILeague *ILeagueTransactor) ReceiveReward(opts *bind.TransactOpts, raceId [32]byte, slotId *big.Int) (*types.Transaction, error) {
	return _ILeague.contract.Transact(opts, "receiveReward", raceId, slotId)
}

// ReceiveReward is a paid mutator transaction binding the contract method 0x1a0fc715.
//
// Solidity: function receiveReward(bytes32 raceId, uint256 slotId) returns()
func (_ILeague *ILeagueSession) ReceiveReward(raceId [32]byte, slotId *big.Int) (*types.Transaction, error) {
	return _ILeague.Contract.ReceiveReward(&_ILeague.TransactOpts, raceId, slotId)
}

// ReceiveReward is a paid mutator transaction binding the contract method 0x1a0fc715.
//
// Solidity: function receiveReward(bytes32 raceId, uint256 slotId) returns()
func (_ILeague *ILeagueTransactorSession) ReceiveReward(raceId [32]byte, slotId *big.Int) (*types.Transaction, error) {
	return _ILeague.Contract.ReceiveReward(&_ILeague.TransactOpts, raceId, slotId)
}

// RegisterRace is a paid mutator transaction binding the contract method 0xe614880f.
//
// Solidity: function registerRace(bytes32 raceId, uint8 slotId) returns()
func (_ILeague *ILeagueTransactor) RegisterRace(opts *bind.TransactOpts, raceId [32]byte, slotId uint8) (*types.Transaction, error) {
	return _ILeague.contract.Transact(opts, "registerRace", raceId, slotId)
}

// RegisterRace is a paid mutator transaction binding the contract method 0xe614880f.
//
// Solidity: function registerRace(bytes32 raceId, uint8 slotId) returns()
func (_ILeague *ILeagueSession) RegisterRace(raceId [32]byte, slotId uint8) (*types.Transaction, error) {
	return _ILeague.Contract.RegisterRace(&_ILeague.TransactOpts, raceId, slotId)
}

// RegisterRace is a paid mutator transaction binding the contract method 0xe614880f.
//
// Solidity: function registerRace(bytes32 raceId, uint8 slotId) returns()
func (_ILeague *ILeagueTransactorSession) RegisterRace(raceId [32]byte, slotId uint8) (*types.Transaction, error) {
	return _ILeague.Contract.RegisterRace(&_ILeague.TransactOpts, raceId, slotId)
}

// RemoveReward is a paid mutator transaction binding the contract method 0x70e45679.
//
// Solidity: function removeReward(bytes32 raceId, uint256 winnerIndex) returns()
func (_ILeague *ILeagueTransactor) RemoveReward(opts *bind.TransactOpts, raceId [32]byte, winnerIndex *big.Int) (*types.Transaction, error) {
	return _ILeague.contract.Transact(opts, "removeReward", raceId, winnerIndex)
}

// RemoveReward is a paid mutator transaction binding the contract method 0x70e45679.
//
// Solidity: function removeReward(bytes32 raceId, uint256 winnerIndex) returns()
func (_ILeague *ILeagueSession) RemoveReward(raceId [32]byte, winnerIndex *big.Int) (*types.Transaction, error) {
	return _ILeague.Contract.RemoveReward(&_ILeague.TransactOpts, raceId, winnerIndex)
}

// RemoveReward is a paid mutator transaction binding the contract method 0x70e45679.
//
// Solidity: function removeReward(bytes32 raceId, uint256 winnerIndex) returns()
func (_ILeague *ILeagueTransactorSession) RemoveReward(raceId [32]byte, winnerIndex *big.Int) (*types.Transaction, error) {
	return _ILeague.Contract.RemoveReward(&_ILeague.TransactOpts, raceId, winnerIndex)
}

// UpdateRaceResult is a paid mutator transaction binding the contract method 0x8c9d6dbd.
//
// Solidity: function updateRaceResult(bytes32 raceId, bytes32 result) returns()
func (_ILeague *ILeagueTransactor) UpdateRaceResult(opts *bind.TransactOpts, raceId [32]byte, result [32]byte) (*types.Transaction, error) {
	return _ILeague.contract.Transact(opts, "updateRaceResult", raceId, result)
}

// UpdateRaceResult is a paid mutator transaction binding the contract method 0x8c9d6dbd.
//
// Solidity: function updateRaceResult(bytes32 raceId, bytes32 result) returns()
func (_ILeague *ILeagueSession) UpdateRaceResult(raceId [32]byte, result [32]byte) (*types.Transaction, error) {
	return _ILeague.Contract.UpdateRaceResult(&_ILeague.TransactOpts, raceId, result)
}

// UpdateRaceResult is a paid mutator transaction binding the contract method 0x8c9d6dbd.
//
// Solidity: function updateRaceResult(bytes32 raceId, bytes32 result) returns()
func (_ILeague *ILeagueTransactorSession) UpdateRaceResult(raceId [32]byte, result [32]byte) (*types.Transaction, error) {
	return _ILeague.Contract.UpdateRaceResult(&_ILeague.TransactOpts, raceId, result)
}

// ILeagueRaceCancelledIterator is returned from FilterRaceCancelled and is used to iterate over the raw logs and unpacked data for RaceCancelled events raised by the ILeague contract.
type ILeagueRaceCancelledIterator struct {
	Event *ILeagueRaceCancelled // Event containing the contract specifics and raw log

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
func (it *ILeagueRaceCancelledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ILeagueRaceCancelled)
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
		it.Event = new(ILeagueRaceCancelled)
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
func (it *ILeagueRaceCancelledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ILeagueRaceCancelledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ILeagueRaceCancelled represents a RaceCancelled event raised by the ILeague contract.
type ILeagueRaceCancelled struct {
	Id  [32]byte
	Raw types.Log // Blockchain specific contextual infos
}

// FilterRaceCancelled is a free log retrieval operation binding the contract event 0x3f3669b11af9b00fb2a95f8ded36a79a2e1d5374efbc5ec6b2b5b96ceaf9ac6a.
//
// Solidity: event RaceCancelled(bytes32 id)
func (_ILeague *ILeagueFilterer) FilterRaceCancelled(opts *bind.FilterOpts) (*ILeagueRaceCancelledIterator, error) {

	logs, sub, err := _ILeague.contract.FilterLogs(opts, "RaceCancelled")
	if err != nil {
		return nil, err
	}
	return &ILeagueRaceCancelledIterator{contract: _ILeague.contract, event: "RaceCancelled", logs: logs, sub: sub}, nil
}

// WatchRaceCancelled is a free log subscription operation binding the contract event 0x3f3669b11af9b00fb2a95f8ded36a79a2e1d5374efbc5ec6b2b5b96ceaf9ac6a.
//
// Solidity: event RaceCancelled(bytes32 id)
func (_ILeague *ILeagueFilterer) WatchRaceCancelled(opts *bind.WatchOpts, sink chan<- *ILeagueRaceCancelled) (event.Subscription, error) {

	logs, sub, err := _ILeague.contract.WatchLogs(opts, "RaceCancelled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ILeagueRaceCancelled)
				if err := _ILeague.contract.UnpackLog(event, "RaceCancelled", log); err != nil {
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

// ParseRaceCancelled is a log parse operation binding the contract event 0x3f3669b11af9b00fb2a95f8ded36a79a2e1d5374efbc5ec6b2b5b96ceaf9ac6a.
//
// Solidity: event RaceCancelled(bytes32 id)
func (_ILeague *ILeagueFilterer) ParseRaceCancelled(log types.Log) (*ILeagueRaceCancelled, error) {
	event := new(ILeagueRaceCancelled)
	if err := _ILeague.contract.UnpackLog(event, "RaceCancelled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ILeagueRaceCreatedIterator is returned from FilterRaceCreated and is used to iterate over the raw logs and unpacked data for RaceCreated events raised by the ILeague contract.
type ILeagueRaceCreatedIterator struct {
	Event *ILeagueRaceCreated // Event containing the contract specifics and raw log

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
func (it *ILeagueRaceCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ILeagueRaceCreated)
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
		it.Event = new(ILeagueRaceCreated)
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
func (it *ILeagueRaceCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ILeagueRaceCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ILeagueRaceCreated represents a RaceCreated event raised by the ILeague contract.
type ILeagueRaceCreated struct {
	Id      [32]byte
	Slots   *big.Int
	StartAt *big.Int
	RaceNo  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRaceCreated is a free log retrieval operation binding the contract event 0x7d96e4700661907fc6a8ea8f3b79949cb0170bd192240352fdd5f891f5c2d562.
//
// Solidity: event RaceCreated(bytes32 id, uint256 slots, uint256 startAt, uint256 raceNo)
func (_ILeague *ILeagueFilterer) FilterRaceCreated(opts *bind.FilterOpts) (*ILeagueRaceCreatedIterator, error) {

	logs, sub, err := _ILeague.contract.FilterLogs(opts, "RaceCreated")
	if err != nil {
		return nil, err
	}
	return &ILeagueRaceCreatedIterator{contract: _ILeague.contract, event: "RaceCreated", logs: logs, sub: sub}, nil
}

// WatchRaceCreated is a free log subscription operation binding the contract event 0x7d96e4700661907fc6a8ea8f3b79949cb0170bd192240352fdd5f891f5c2d562.
//
// Solidity: event RaceCreated(bytes32 id, uint256 slots, uint256 startAt, uint256 raceNo)
func (_ILeague *ILeagueFilterer) WatchRaceCreated(opts *bind.WatchOpts, sink chan<- *ILeagueRaceCreated) (event.Subscription, error) {

	logs, sub, err := _ILeague.contract.WatchLogs(opts, "RaceCreated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ILeagueRaceCreated)
				if err := _ILeague.contract.UnpackLog(event, "RaceCreated", log); err != nil {
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

// ParseRaceCreated is a log parse operation binding the contract event 0x7d96e4700661907fc6a8ea8f3b79949cb0170bd192240352fdd5f891f5c2d562.
//
// Solidity: event RaceCreated(bytes32 id, uint256 slots, uint256 startAt, uint256 raceNo)
func (_ILeague *ILeagueFilterer) ParseRaceCreated(log types.Log) (*ILeagueRaceCreated, error) {
	event := new(ILeagueRaceCreated)
	if err := _ILeague.contract.UnpackLog(event, "RaceCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ILeagueRaceResultUpdatedIterator is returned from FilterRaceResultUpdated and is used to iterate over the raw logs and unpacked data for RaceResultUpdated events raised by the ILeague contract.
type ILeagueRaceResultUpdatedIterator struct {
	Event *ILeagueRaceResultUpdated // Event containing the contract specifics and raw log

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
func (it *ILeagueRaceResultUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ILeagueRaceResultUpdated)
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
		it.Event = new(ILeagueRaceResultUpdated)
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
func (it *ILeagueRaceResultUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ILeagueRaceResultUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ILeagueRaceResultUpdated represents a RaceResultUpdated event raised by the ILeague contract.
type ILeagueRaceResultUpdated struct {
	Id     [32]byte
	Result [32]byte
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRaceResultUpdated is a free log retrieval operation binding the contract event 0xd82663592968d73ce1995154e44c793fcb46a4006abfb7438656fb0d7ba5ff49.
//
// Solidity: event RaceResultUpdated(bytes32 id, bytes32 result)
func (_ILeague *ILeagueFilterer) FilterRaceResultUpdated(opts *bind.FilterOpts) (*ILeagueRaceResultUpdatedIterator, error) {

	logs, sub, err := _ILeague.contract.FilterLogs(opts, "RaceResultUpdated")
	if err != nil {
		return nil, err
	}
	return &ILeagueRaceResultUpdatedIterator{contract: _ILeague.contract, event: "RaceResultUpdated", logs: logs, sub: sub}, nil
}

// WatchRaceResultUpdated is a free log subscription operation binding the contract event 0xd82663592968d73ce1995154e44c793fcb46a4006abfb7438656fb0d7ba5ff49.
//
// Solidity: event RaceResultUpdated(bytes32 id, bytes32 result)
func (_ILeague *ILeagueFilterer) WatchRaceResultUpdated(opts *bind.WatchOpts, sink chan<- *ILeagueRaceResultUpdated) (event.Subscription, error) {

	logs, sub, err := _ILeague.contract.WatchLogs(opts, "RaceResultUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ILeagueRaceResultUpdated)
				if err := _ILeague.contract.UnpackLog(event, "RaceResultUpdated", log); err != nil {
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

// ParseRaceResultUpdated is a log parse operation binding the contract event 0xd82663592968d73ce1995154e44c793fcb46a4006abfb7438656fb0d7ba5ff49.
//
// Solidity: event RaceResultUpdated(bytes32 id, bytes32 result)
func (_ILeague *ILeagueFilterer) ParseRaceResultUpdated(log types.Log) (*ILeagueRaceResultUpdated, error) {
	event := new(ILeagueRaceResultUpdated)
	if err := _ILeague.contract.UnpackLog(event, "RaceResultUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ILeagueRegisteredIterator is returned from FilterRegistered and is used to iterate over the raw logs and unpacked data for Registered events raised by the ILeague contract.
type ILeagueRegisteredIterator struct {
	Event *ILeagueRegistered // Event containing the contract specifics and raw log

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
func (it *ILeagueRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ILeagueRegistered)
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
		it.Event = new(ILeagueRegistered)
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
func (it *ILeagueRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ILeagueRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ILeagueRegistered represents a Registered event raised by the ILeague contract.
type ILeagueRegistered struct {
	SlotId      *big.Int
	Participant common.Address
	RaceId      [32]byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterRegistered is a free log retrieval operation binding the contract event 0xd62c7bd2779dc6ebe426bed623399b1fc1d190c063eadefcffffb1efe6f253ef.
//
// Solidity: event Registered(uint256 slotId, address participant, bytes32 raceId)
func (_ILeague *ILeagueFilterer) FilterRegistered(opts *bind.FilterOpts) (*ILeagueRegisteredIterator, error) {

	logs, sub, err := _ILeague.contract.FilterLogs(opts, "Registered")
	if err != nil {
		return nil, err
	}
	return &ILeagueRegisteredIterator{contract: _ILeague.contract, event: "Registered", logs: logs, sub: sub}, nil
}

// WatchRegistered is a free log subscription operation binding the contract event 0xd62c7bd2779dc6ebe426bed623399b1fc1d190c063eadefcffffb1efe6f253ef.
//
// Solidity: event Registered(uint256 slotId, address participant, bytes32 raceId)
func (_ILeague *ILeagueFilterer) WatchRegistered(opts *bind.WatchOpts, sink chan<- *ILeagueRegistered) (event.Subscription, error) {

	logs, sub, err := _ILeague.contract.WatchLogs(opts, "Registered")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ILeagueRegistered)
				if err := _ILeague.contract.UnpackLog(event, "Registered", log); err != nil {
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
func (_ILeague *ILeagueFilterer) ParseRegistered(log types.Log) (*ILeagueRegistered, error) {
	event := new(ILeagueRegistered)
	if err := _ILeague.contract.UnpackLog(event, "Registered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ILeagueRewardAddedIterator is returned from FilterRewardAdded and is used to iterate over the raw logs and unpacked data for RewardAdded events raised by the ILeague contract.
type ILeagueRewardAddedIterator struct {
	Event *ILeagueRewardAdded // Event containing the contract specifics and raw log

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
func (it *ILeagueRewardAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ILeagueRewardAdded)
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
		it.Event = new(ILeagueRewardAdded)
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
func (it *ILeagueRewardAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ILeagueRewardAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ILeagueRewardAdded represents a RewardAdded event raised by the ILeague contract.
type ILeagueRewardAdded struct {
	RaceId      [32]byte
	NftRewardId *big.Int
	WinnerIndex *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterRewardAdded is a free log retrieval operation binding the contract event 0xa28d22473b1a2f01a0708abe09b0bf3b7e1b0f30b0f759e2948b55c8e6fdd6e3.
//
// Solidity: event RewardAdded(bytes32 raceId, uint256 nftRewardId, uint256 winnerIndex)
func (_ILeague *ILeagueFilterer) FilterRewardAdded(opts *bind.FilterOpts) (*ILeagueRewardAddedIterator, error) {

	logs, sub, err := _ILeague.contract.FilterLogs(opts, "RewardAdded")
	if err != nil {
		return nil, err
	}
	return &ILeagueRewardAddedIterator{contract: _ILeague.contract, event: "RewardAdded", logs: logs, sub: sub}, nil
}

// WatchRewardAdded is a free log subscription operation binding the contract event 0xa28d22473b1a2f01a0708abe09b0bf3b7e1b0f30b0f759e2948b55c8e6fdd6e3.
//
// Solidity: event RewardAdded(bytes32 raceId, uint256 nftRewardId, uint256 winnerIndex)
func (_ILeague *ILeagueFilterer) WatchRewardAdded(opts *bind.WatchOpts, sink chan<- *ILeagueRewardAdded) (event.Subscription, error) {

	logs, sub, err := _ILeague.contract.WatchLogs(opts, "RewardAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ILeagueRewardAdded)
				if err := _ILeague.contract.UnpackLog(event, "RewardAdded", log); err != nil {
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

// ParseRewardAdded is a log parse operation binding the contract event 0xa28d22473b1a2f01a0708abe09b0bf3b7e1b0f30b0f759e2948b55c8e6fdd6e3.
//
// Solidity: event RewardAdded(bytes32 raceId, uint256 nftRewardId, uint256 winnerIndex)
func (_ILeague *ILeagueFilterer) ParseRewardAdded(log types.Log) (*ILeagueRewardAdded, error) {
	event := new(ILeagueRewardAdded)
	if err := _ILeague.contract.UnpackLog(event, "RewardAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ILeagueRewardReceivedIterator is returned from FilterRewardReceived and is used to iterate over the raw logs and unpacked data for RewardReceived events raised by the ILeague contract.
type ILeagueRewardReceivedIterator struct {
	Event *ILeagueRewardReceived // Event containing the contract specifics and raw log

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
func (it *ILeagueRewardReceivedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ILeagueRewardReceived)
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
		it.Event = new(ILeagueRewardReceived)
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
func (it *ILeagueRewardReceivedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ILeagueRewardReceivedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ILeagueRewardReceived represents a RewardReceived event raised by the ILeague contract.
type ILeagueRewardReceived struct {
	RaceId      [32]byte
	SlotId      *big.Int
	NftRewardId *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterRewardReceived is a free log retrieval operation binding the contract event 0xe2ece170b35ae79cbe45ffdaf4931b8bc45056da5d3e9aa4aea67fcdf307929a.
//
// Solidity: event RewardReceived(bytes32 raceId, uint256 slotId, uint256 nftRewardId)
func (_ILeague *ILeagueFilterer) FilterRewardReceived(opts *bind.FilterOpts) (*ILeagueRewardReceivedIterator, error) {

	logs, sub, err := _ILeague.contract.FilterLogs(opts, "RewardReceived")
	if err != nil {
		return nil, err
	}
	return &ILeagueRewardReceivedIterator{contract: _ILeague.contract, event: "RewardReceived", logs: logs, sub: sub}, nil
}

// WatchRewardReceived is a free log subscription operation binding the contract event 0xe2ece170b35ae79cbe45ffdaf4931b8bc45056da5d3e9aa4aea67fcdf307929a.
//
// Solidity: event RewardReceived(bytes32 raceId, uint256 slotId, uint256 nftRewardId)
func (_ILeague *ILeagueFilterer) WatchRewardReceived(opts *bind.WatchOpts, sink chan<- *ILeagueRewardReceived) (event.Subscription, error) {

	logs, sub, err := _ILeague.contract.WatchLogs(opts, "RewardReceived")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ILeagueRewardReceived)
				if err := _ILeague.contract.UnpackLog(event, "RewardReceived", log); err != nil {
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
func (_ILeague *ILeagueFilterer) ParseRewardReceived(log types.Log) (*ILeagueRewardReceived, error) {
	event := new(ILeagueRewardReceived)
	if err := _ILeague.contract.UnpackLog(event, "RewardReceived", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ILeagueRewardRemovedIterator is returned from FilterRewardRemoved and is used to iterate over the raw logs and unpacked data for RewardRemoved events raised by the ILeague contract.
type ILeagueRewardRemovedIterator struct {
	Event *ILeagueRewardRemoved // Event containing the contract specifics and raw log

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
func (it *ILeagueRewardRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ILeagueRewardRemoved)
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
		it.Event = new(ILeagueRewardRemoved)
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
func (it *ILeagueRewardRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ILeagueRewardRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ILeagueRewardRemoved represents a RewardRemoved event raised by the ILeague contract.
type ILeagueRewardRemoved struct {
	RaceId      [32]byte
	NftRewardId *big.Int
	WinnerIndex *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterRewardRemoved is a free log retrieval operation binding the contract event 0xefba7d4f0912cb3bf75c07898c0969e76416cc35a7046ba3c04902f3f3f607f3.
//
// Solidity: event RewardRemoved(bytes32 raceId, uint256 nftRewardId, uint256 winnerIndex)
func (_ILeague *ILeagueFilterer) FilterRewardRemoved(opts *bind.FilterOpts) (*ILeagueRewardRemovedIterator, error) {

	logs, sub, err := _ILeague.contract.FilterLogs(opts, "RewardRemoved")
	if err != nil {
		return nil, err
	}
	return &ILeagueRewardRemovedIterator{contract: _ILeague.contract, event: "RewardRemoved", logs: logs, sub: sub}, nil
}

// WatchRewardRemoved is a free log subscription operation binding the contract event 0xefba7d4f0912cb3bf75c07898c0969e76416cc35a7046ba3c04902f3f3f607f3.
//
// Solidity: event RewardRemoved(bytes32 raceId, uint256 nftRewardId, uint256 winnerIndex)
func (_ILeague *ILeagueFilterer) WatchRewardRemoved(opts *bind.WatchOpts, sink chan<- *ILeagueRewardRemoved) (event.Subscription, error) {

	logs, sub, err := _ILeague.contract.WatchLogs(opts, "RewardRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ILeagueRewardRemoved)
				if err := _ILeague.contract.UnpackLog(event, "RewardRemoved", log); err != nil {
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

// ParseRewardRemoved is a log parse operation binding the contract event 0xefba7d4f0912cb3bf75c07898c0969e76416cc35a7046ba3c04902f3f3f607f3.
//
// Solidity: event RewardRemoved(bytes32 raceId, uint256 nftRewardId, uint256 winnerIndex)
func (_ILeague *ILeagueFilterer) ParseRewardRemoved(log types.Log) (*ILeagueRewardRemoved, error) {
	event := new(ILeagueRewardRemoved)
	if err := _ILeague.contract.UnpackLog(event, "RewardRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
