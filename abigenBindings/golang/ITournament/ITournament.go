// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ITournament

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

// ITournamentRace is an auto generated low-level Go binding around an user-defined struct.
type ITournamentRace struct {
	NoSlot  uint8
	StartAt uint32
	Result  [27]byte
}

// ITournamentTournamentInfo is an auto generated low-level Go binding around an user-defined struct.
type ITournamentTournamentInfo struct {
	TotalRace      uint8
	CreatedRace    uint8
	EndedRace      uint8
	TournamentName string
}

// ITournamentMetaData contains all meta data concerning the ITournament contract.
var ITournamentMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"AlreadyRegistered\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_currentRaceNo\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalRace\",\"type\":\"uint256\"}],\"name\":\"CannotCreateMoreRace\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidContract\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidRegister\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidSender\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidSlot\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotEndYet\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RaceExisted\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RaceNotExisted\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RaceWasUpdated\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"}],\"name\":\"TooEarly\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"}],\"name\":\"TooLate\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"}],\"name\":\"RaceCancelled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"noSlot\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"raceNo\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"startAt\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"}],\"name\":\"RaceCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes27\",\"name\":\"result\",\"type\":\"bytes27\"}],\"name\":\"RaceResultUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"slotId\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"participant\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"raceId\",\"type\":\"bytes32\"}],\"name\":\"Registered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"winnerIndex\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"score\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"winner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nftRewardId\",\"type\":\"uint256\"}],\"name\":\"RewardGranted\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"noSlot\",\"type\":\"uint8\"},{\"internalType\":\"uint32\",\"name\":\"startAt\",\"type\":\"uint32\"}],\"name\":\"createRace\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"raceId\",\"type\":\"bytes32\"}],\"name\":\"raceInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"uint8\",\"name\":\"noSlot\",\"type\":\"uint8\"},{\"internalType\":\"uint32\",\"name\":\"startAt\",\"type\":\"uint32\"},{\"internalType\":\"bytes27\",\"name\":\"result\",\"type\":\"bytes27\"}],\"internalType\":\"structITournament.Race\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"register\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"raceId\",\"type\":\"bytes32\"}],\"name\":\"registeredSlot\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"register\",\"type\":\"address\"}],\"name\":\"getTotalScore\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tournamentInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"uint8\",\"name\":\"totalRace\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"createdRace\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"endedRace\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"tournamentName\",\"type\":\"string\"}],\"internalType\":\"structITournament.TournamentInfo\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"raceId\",\"type\":\"bytes32\"}],\"name\":\"cancelRace\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"slotId\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"raceId\",\"type\":\"bytes32\"}],\"name\":\"registerRace\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"raceId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes27\",\"name\":\"result\",\"type\":\"bytes27\"}],\"name\":\"updateRaceResult\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"winnerIndex\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"winner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nftRewardId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"tokenURI\",\"type\":\"string\"}],\"name\":\"grandReward\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// ITournamentABI is the input ABI used to generate the binding from.
// Deprecated: Use ITournamentMetaData.ABI instead.
var ITournamentABI = ITournamentMetaData.ABI

// ITournament is an auto generated Go binding around an Ethereum contract.
type ITournament struct {
	ITournamentCaller     // Read-only binding to the contract
	ITournamentTransactor // Write-only binding to the contract
	ITournamentFilterer   // Log filterer for contract events
}

// ITournamentCaller is an auto generated read-only Go binding around an Ethereum contract.
type ITournamentCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ITournamentTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ITournamentTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ITournamentFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ITournamentFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ITournamentSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ITournamentSession struct {
	Contract     *ITournament      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ITournamentCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ITournamentCallerSession struct {
	Contract *ITournamentCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// ITournamentTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ITournamentTransactorSession struct {
	Contract     *ITournamentTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// ITournamentRaw is an auto generated low-level Go binding around an Ethereum contract.
type ITournamentRaw struct {
	Contract *ITournament // Generic contract binding to access the raw methods on
}

// ITournamentCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ITournamentCallerRaw struct {
	Contract *ITournamentCaller // Generic read-only contract binding to access the raw methods on
}

// ITournamentTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ITournamentTransactorRaw struct {
	Contract *ITournamentTransactor // Generic write-only contract binding to access the raw methods on
}

// NewITournament creates a new instance of ITournament, bound to a specific deployed contract.
func NewITournament(address common.Address, backend bind.ContractBackend) (*ITournament, error) {
	contract, err := bindITournament(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ITournament{ITournamentCaller: ITournamentCaller{contract: contract}, ITournamentTransactor: ITournamentTransactor{contract: contract}, ITournamentFilterer: ITournamentFilterer{contract: contract}}, nil
}

// NewITournamentCaller creates a new read-only instance of ITournament, bound to a specific deployed contract.
func NewITournamentCaller(address common.Address, caller bind.ContractCaller) (*ITournamentCaller, error) {
	contract, err := bindITournament(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ITournamentCaller{contract: contract}, nil
}

// NewITournamentTransactor creates a new write-only instance of ITournament, bound to a specific deployed contract.
func NewITournamentTransactor(address common.Address, transactor bind.ContractTransactor) (*ITournamentTransactor, error) {
	contract, err := bindITournament(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ITournamentTransactor{contract: contract}, nil
}

// NewITournamentFilterer creates a new log filterer instance of ITournament, bound to a specific deployed contract.
func NewITournamentFilterer(address common.Address, filterer bind.ContractFilterer) (*ITournamentFilterer, error) {
	contract, err := bindITournament(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ITournamentFilterer{contract: contract}, nil
}

// bindITournament binds a generic wrapper to an already deployed contract.
func bindITournament(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ITournamentABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ITournament *ITournamentRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ITournament.Contract.ITournamentCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ITournament *ITournamentRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ITournament.Contract.ITournamentTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ITournament *ITournamentRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ITournament.Contract.ITournamentTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ITournament *ITournamentCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ITournament.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ITournament *ITournamentTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ITournament.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ITournament *ITournamentTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ITournament.Contract.contract.Transact(opts, method, params...)
}

// GetTotalScore is a free data retrieval call binding the contract method 0xd283b3c5.
//
// Solidity: function getTotalScore(address register) view returns(uint8)
func (_ITournament *ITournamentCaller) GetTotalScore(opts *bind.CallOpts, register common.Address) (uint8, error) {
	var out []interface{}
	err := _ITournament.contract.Call(opts, &out, "getTotalScore", register)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// GetTotalScore is a free data retrieval call binding the contract method 0xd283b3c5.
//
// Solidity: function getTotalScore(address register) view returns(uint8)
func (_ITournament *ITournamentSession) GetTotalScore(register common.Address) (uint8, error) {
	return _ITournament.Contract.GetTotalScore(&_ITournament.CallOpts, register)
}

// GetTotalScore is a free data retrieval call binding the contract method 0xd283b3c5.
//
// Solidity: function getTotalScore(address register) view returns(uint8)
func (_ITournament *ITournamentCallerSession) GetTotalScore(register common.Address) (uint8, error) {
	return _ITournament.Contract.GetTotalScore(&_ITournament.CallOpts, register)
}

// RaceInfo is a free data retrieval call binding the contract method 0xef8597fd.
//
// Solidity: function raceInfo(bytes32 raceId) view returns((uint8,uint32,bytes27))
func (_ITournament *ITournamentCaller) RaceInfo(opts *bind.CallOpts, raceId [32]byte) (ITournamentRace, error) {
	var out []interface{}
	err := _ITournament.contract.Call(opts, &out, "raceInfo", raceId)

	if err != nil {
		return *new(ITournamentRace), err
	}

	out0 := *abi.ConvertType(out[0], new(ITournamentRace)).(*ITournamentRace)

	return out0, err

}

// RaceInfo is a free data retrieval call binding the contract method 0xef8597fd.
//
// Solidity: function raceInfo(bytes32 raceId) view returns((uint8,uint32,bytes27))
func (_ITournament *ITournamentSession) RaceInfo(raceId [32]byte) (ITournamentRace, error) {
	return _ITournament.Contract.RaceInfo(&_ITournament.CallOpts, raceId)
}

// RaceInfo is a free data retrieval call binding the contract method 0xef8597fd.
//
// Solidity: function raceInfo(bytes32 raceId) view returns((uint8,uint32,bytes27))
func (_ITournament *ITournamentCallerSession) RaceInfo(raceId [32]byte) (ITournamentRace, error) {
	return _ITournament.Contract.RaceInfo(&_ITournament.CallOpts, raceId)
}

// RegisteredSlot is a free data retrieval call binding the contract method 0x96fd6980.
//
// Solidity: function registeredSlot(address register, bytes32 raceId) view returns(uint8)
func (_ITournament *ITournamentCaller) RegisteredSlot(opts *bind.CallOpts, register common.Address, raceId [32]byte) (uint8, error) {
	var out []interface{}
	err := _ITournament.contract.Call(opts, &out, "registeredSlot", register, raceId)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// RegisteredSlot is a free data retrieval call binding the contract method 0x96fd6980.
//
// Solidity: function registeredSlot(address register, bytes32 raceId) view returns(uint8)
func (_ITournament *ITournamentSession) RegisteredSlot(register common.Address, raceId [32]byte) (uint8, error) {
	return _ITournament.Contract.RegisteredSlot(&_ITournament.CallOpts, register, raceId)
}

// RegisteredSlot is a free data retrieval call binding the contract method 0x96fd6980.
//
// Solidity: function registeredSlot(address register, bytes32 raceId) view returns(uint8)
func (_ITournament *ITournamentCallerSession) RegisteredSlot(register common.Address, raceId [32]byte) (uint8, error) {
	return _ITournament.Contract.RegisteredSlot(&_ITournament.CallOpts, register, raceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_ITournament *ITournamentCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _ITournament.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_ITournament *ITournamentSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _ITournament.Contract.SupportsInterface(&_ITournament.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_ITournament *ITournamentCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _ITournament.Contract.SupportsInterface(&_ITournament.CallOpts, interfaceId)
}

// TournamentInfo is a free data retrieval call binding the contract method 0xc1fab611.
//
// Solidity: function tournamentInfo() view returns((uint8,uint8,uint8,string))
func (_ITournament *ITournamentCaller) TournamentInfo(opts *bind.CallOpts) (ITournamentTournamentInfo, error) {
	var out []interface{}
	err := _ITournament.contract.Call(opts, &out, "tournamentInfo")

	if err != nil {
		return *new(ITournamentTournamentInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(ITournamentTournamentInfo)).(*ITournamentTournamentInfo)

	return out0, err

}

// TournamentInfo is a free data retrieval call binding the contract method 0xc1fab611.
//
// Solidity: function tournamentInfo() view returns((uint8,uint8,uint8,string))
func (_ITournament *ITournamentSession) TournamentInfo() (ITournamentTournamentInfo, error) {
	return _ITournament.Contract.TournamentInfo(&_ITournament.CallOpts)
}

// TournamentInfo is a free data retrieval call binding the contract method 0xc1fab611.
//
// Solidity: function tournamentInfo() view returns((uint8,uint8,uint8,string))
func (_ITournament *ITournamentCallerSession) TournamentInfo() (ITournamentTournamentInfo, error) {
	return _ITournament.Contract.TournamentInfo(&_ITournament.CallOpts)
}

// CancelRace is a paid mutator transaction binding the contract method 0xc05c107d.
//
// Solidity: function cancelRace(bytes32 raceId) returns()
func (_ITournament *ITournamentTransactor) CancelRace(opts *bind.TransactOpts, raceId [32]byte) (*types.Transaction, error) {
	return _ITournament.contract.Transact(opts, "cancelRace", raceId)
}

// CancelRace is a paid mutator transaction binding the contract method 0xc05c107d.
//
// Solidity: function cancelRace(bytes32 raceId) returns()
func (_ITournament *ITournamentSession) CancelRace(raceId [32]byte) (*types.Transaction, error) {
	return _ITournament.Contract.CancelRace(&_ITournament.TransactOpts, raceId)
}

// CancelRace is a paid mutator transaction binding the contract method 0xc05c107d.
//
// Solidity: function cancelRace(bytes32 raceId) returns()
func (_ITournament *ITournamentTransactorSession) CancelRace(raceId [32]byte) (*types.Transaction, error) {
	return _ITournament.Contract.CancelRace(&_ITournament.TransactOpts, raceId)
}

// CreateRace is a paid mutator transaction binding the contract method 0x928e03da.
//
// Solidity: function createRace(uint8 noSlot, uint32 startAt) returns(bytes32)
func (_ITournament *ITournamentTransactor) CreateRace(opts *bind.TransactOpts, noSlot uint8, startAt uint32) (*types.Transaction, error) {
	return _ITournament.contract.Transact(opts, "createRace", noSlot, startAt)
}

// CreateRace is a paid mutator transaction binding the contract method 0x928e03da.
//
// Solidity: function createRace(uint8 noSlot, uint32 startAt) returns(bytes32)
func (_ITournament *ITournamentSession) CreateRace(noSlot uint8, startAt uint32) (*types.Transaction, error) {
	return _ITournament.Contract.CreateRace(&_ITournament.TransactOpts, noSlot, startAt)
}

// CreateRace is a paid mutator transaction binding the contract method 0x928e03da.
//
// Solidity: function createRace(uint8 noSlot, uint32 startAt) returns(bytes32)
func (_ITournament *ITournamentTransactorSession) CreateRace(noSlot uint8, startAt uint32) (*types.Transaction, error) {
	return _ITournament.Contract.CreateRace(&_ITournament.TransactOpts, noSlot, startAt)
}

// GrandReward is a paid mutator transaction binding the contract method 0xda0df24f.
//
// Solidity: function grandReward(uint8 winnerIndex, address winner, uint256 nftRewardId, string tokenURI) returns()
func (_ITournament *ITournamentTransactor) GrandReward(opts *bind.TransactOpts, winnerIndex uint8, winner common.Address, nftRewardId *big.Int, tokenURI string) (*types.Transaction, error) {
	return _ITournament.contract.Transact(opts, "grandReward", winnerIndex, winner, nftRewardId, tokenURI)
}

// GrandReward is a paid mutator transaction binding the contract method 0xda0df24f.
//
// Solidity: function grandReward(uint8 winnerIndex, address winner, uint256 nftRewardId, string tokenURI) returns()
func (_ITournament *ITournamentSession) GrandReward(winnerIndex uint8, winner common.Address, nftRewardId *big.Int, tokenURI string) (*types.Transaction, error) {
	return _ITournament.Contract.GrandReward(&_ITournament.TransactOpts, winnerIndex, winner, nftRewardId, tokenURI)
}

// GrandReward is a paid mutator transaction binding the contract method 0xda0df24f.
//
// Solidity: function grandReward(uint8 winnerIndex, address winner, uint256 nftRewardId, string tokenURI) returns()
func (_ITournament *ITournamentTransactorSession) GrandReward(winnerIndex uint8, winner common.Address, nftRewardId *big.Int, tokenURI string) (*types.Transaction, error) {
	return _ITournament.Contract.GrandReward(&_ITournament.TransactOpts, winnerIndex, winner, nftRewardId, tokenURI)
}

// RegisterRace is a paid mutator transaction binding the contract method 0x1e621ebe.
//
// Solidity: function registerRace(uint8 slotId, bytes32 raceId) returns()
func (_ITournament *ITournamentTransactor) RegisterRace(opts *bind.TransactOpts, slotId uint8, raceId [32]byte) (*types.Transaction, error) {
	return _ITournament.contract.Transact(opts, "registerRace", slotId, raceId)
}

// RegisterRace is a paid mutator transaction binding the contract method 0x1e621ebe.
//
// Solidity: function registerRace(uint8 slotId, bytes32 raceId) returns()
func (_ITournament *ITournamentSession) RegisterRace(slotId uint8, raceId [32]byte) (*types.Transaction, error) {
	return _ITournament.Contract.RegisterRace(&_ITournament.TransactOpts, slotId, raceId)
}

// RegisterRace is a paid mutator transaction binding the contract method 0x1e621ebe.
//
// Solidity: function registerRace(uint8 slotId, bytes32 raceId) returns()
func (_ITournament *ITournamentTransactorSession) RegisterRace(slotId uint8, raceId [32]byte) (*types.Transaction, error) {
	return _ITournament.Contract.RegisterRace(&_ITournament.TransactOpts, slotId, raceId)
}

// UpdateRaceResult is a paid mutator transaction binding the contract method 0xdd0b3778.
//
// Solidity: function updateRaceResult(bytes32 raceId, bytes27 result) returns()
func (_ITournament *ITournamentTransactor) UpdateRaceResult(opts *bind.TransactOpts, raceId [32]byte, result [27]byte) (*types.Transaction, error) {
	return _ITournament.contract.Transact(opts, "updateRaceResult", raceId, result)
}

// UpdateRaceResult is a paid mutator transaction binding the contract method 0xdd0b3778.
//
// Solidity: function updateRaceResult(bytes32 raceId, bytes27 result) returns()
func (_ITournament *ITournamentSession) UpdateRaceResult(raceId [32]byte, result [27]byte) (*types.Transaction, error) {
	return _ITournament.Contract.UpdateRaceResult(&_ITournament.TransactOpts, raceId, result)
}

// UpdateRaceResult is a paid mutator transaction binding the contract method 0xdd0b3778.
//
// Solidity: function updateRaceResult(bytes32 raceId, bytes27 result) returns()
func (_ITournament *ITournamentTransactorSession) UpdateRaceResult(raceId [32]byte, result [27]byte) (*types.Transaction, error) {
	return _ITournament.Contract.UpdateRaceResult(&_ITournament.TransactOpts, raceId, result)
}

// ITournamentRaceCancelledIterator is returned from FilterRaceCancelled and is used to iterate over the raw logs and unpacked data for RaceCancelled events raised by the ITournament contract.
type ITournamentRaceCancelledIterator struct {
	Event *ITournamentRaceCancelled // Event containing the contract specifics and raw log

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
func (it *ITournamentRaceCancelledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ITournamentRaceCancelled)
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
		it.Event = new(ITournamentRaceCancelled)
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
func (it *ITournamentRaceCancelledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ITournamentRaceCancelledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ITournamentRaceCancelled represents a RaceCancelled event raised by the ITournament contract.
type ITournamentRaceCancelled struct {
	Id  [32]byte
	Raw types.Log // Blockchain specific contextual infos
}

// FilterRaceCancelled is a free log retrieval operation binding the contract event 0x3f3669b11af9b00fb2a95f8ded36a79a2e1d5374efbc5ec6b2b5b96ceaf9ac6a.
//
// Solidity: event RaceCancelled(bytes32 id)
func (_ITournament *ITournamentFilterer) FilterRaceCancelled(opts *bind.FilterOpts) (*ITournamentRaceCancelledIterator, error) {

	logs, sub, err := _ITournament.contract.FilterLogs(opts, "RaceCancelled")
	if err != nil {
		return nil, err
	}
	return &ITournamentRaceCancelledIterator{contract: _ITournament.contract, event: "RaceCancelled", logs: logs, sub: sub}, nil
}

// WatchRaceCancelled is a free log subscription operation binding the contract event 0x3f3669b11af9b00fb2a95f8ded36a79a2e1d5374efbc5ec6b2b5b96ceaf9ac6a.
//
// Solidity: event RaceCancelled(bytes32 id)
func (_ITournament *ITournamentFilterer) WatchRaceCancelled(opts *bind.WatchOpts, sink chan<- *ITournamentRaceCancelled) (event.Subscription, error) {

	logs, sub, err := _ITournament.contract.WatchLogs(opts, "RaceCancelled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ITournamentRaceCancelled)
				if err := _ITournament.contract.UnpackLog(event, "RaceCancelled", log); err != nil {
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
func (_ITournament *ITournamentFilterer) ParseRaceCancelled(log types.Log) (*ITournamentRaceCancelled, error) {
	event := new(ITournamentRaceCancelled)
	if err := _ITournament.contract.UnpackLog(event, "RaceCancelled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ITournamentRaceCreatedIterator is returned from FilterRaceCreated and is used to iterate over the raw logs and unpacked data for RaceCreated events raised by the ITournament contract.
type ITournamentRaceCreatedIterator struct {
	Event *ITournamentRaceCreated // Event containing the contract specifics and raw log

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
func (it *ITournamentRaceCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ITournamentRaceCreated)
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
		it.Event = new(ITournamentRaceCreated)
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
func (it *ITournamentRaceCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ITournamentRaceCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ITournamentRaceCreated represents a RaceCreated event raised by the ITournament contract.
type ITournamentRaceCreated struct {
	NoSlot  uint8
	RaceNo  uint8
	StartAt uint32
	Id      [32]byte
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRaceCreated is a free log retrieval operation binding the contract event 0x58e30581b52e563d6f294b197724e473591cf7668760f3bae9a2678502011536.
//
// Solidity: event RaceCreated(uint8 noSlot, uint8 raceNo, uint32 startAt, bytes32 id)
func (_ITournament *ITournamentFilterer) FilterRaceCreated(opts *bind.FilterOpts) (*ITournamentRaceCreatedIterator, error) {

	logs, sub, err := _ITournament.contract.FilterLogs(opts, "RaceCreated")
	if err != nil {
		return nil, err
	}
	return &ITournamentRaceCreatedIterator{contract: _ITournament.contract, event: "RaceCreated", logs: logs, sub: sub}, nil
}

// WatchRaceCreated is a free log subscription operation binding the contract event 0x58e30581b52e563d6f294b197724e473591cf7668760f3bae9a2678502011536.
//
// Solidity: event RaceCreated(uint8 noSlot, uint8 raceNo, uint32 startAt, bytes32 id)
func (_ITournament *ITournamentFilterer) WatchRaceCreated(opts *bind.WatchOpts, sink chan<- *ITournamentRaceCreated) (event.Subscription, error) {

	logs, sub, err := _ITournament.contract.WatchLogs(opts, "RaceCreated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ITournamentRaceCreated)
				if err := _ITournament.contract.UnpackLog(event, "RaceCreated", log); err != nil {
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

// ParseRaceCreated is a log parse operation binding the contract event 0x58e30581b52e563d6f294b197724e473591cf7668760f3bae9a2678502011536.
//
// Solidity: event RaceCreated(uint8 noSlot, uint8 raceNo, uint32 startAt, bytes32 id)
func (_ITournament *ITournamentFilterer) ParseRaceCreated(log types.Log) (*ITournamentRaceCreated, error) {
	event := new(ITournamentRaceCreated)
	if err := _ITournament.contract.UnpackLog(event, "RaceCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ITournamentRaceResultUpdatedIterator is returned from FilterRaceResultUpdated and is used to iterate over the raw logs and unpacked data for RaceResultUpdated events raised by the ITournament contract.
type ITournamentRaceResultUpdatedIterator struct {
	Event *ITournamentRaceResultUpdated // Event containing the contract specifics and raw log

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
func (it *ITournamentRaceResultUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ITournamentRaceResultUpdated)
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
		it.Event = new(ITournamentRaceResultUpdated)
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
func (it *ITournamentRaceResultUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ITournamentRaceResultUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ITournamentRaceResultUpdated represents a RaceResultUpdated event raised by the ITournament contract.
type ITournamentRaceResultUpdated struct {
	Id     [32]byte
	Result [27]byte
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRaceResultUpdated is a free log retrieval operation binding the contract event 0x2ab6bf279c23ac56ddfab0e48cd46e6576b30742d689be7cabf12d8f279c7870.
//
// Solidity: event RaceResultUpdated(bytes32 id, bytes27 result)
func (_ITournament *ITournamentFilterer) FilterRaceResultUpdated(opts *bind.FilterOpts) (*ITournamentRaceResultUpdatedIterator, error) {

	logs, sub, err := _ITournament.contract.FilterLogs(opts, "RaceResultUpdated")
	if err != nil {
		return nil, err
	}
	return &ITournamentRaceResultUpdatedIterator{contract: _ITournament.contract, event: "RaceResultUpdated", logs: logs, sub: sub}, nil
}

// WatchRaceResultUpdated is a free log subscription operation binding the contract event 0x2ab6bf279c23ac56ddfab0e48cd46e6576b30742d689be7cabf12d8f279c7870.
//
// Solidity: event RaceResultUpdated(bytes32 id, bytes27 result)
func (_ITournament *ITournamentFilterer) WatchRaceResultUpdated(opts *bind.WatchOpts, sink chan<- *ITournamentRaceResultUpdated) (event.Subscription, error) {

	logs, sub, err := _ITournament.contract.WatchLogs(opts, "RaceResultUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ITournamentRaceResultUpdated)
				if err := _ITournament.contract.UnpackLog(event, "RaceResultUpdated", log); err != nil {
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

// ParseRaceResultUpdated is a log parse operation binding the contract event 0x2ab6bf279c23ac56ddfab0e48cd46e6576b30742d689be7cabf12d8f279c7870.
//
// Solidity: event RaceResultUpdated(bytes32 id, bytes27 result)
func (_ITournament *ITournamentFilterer) ParseRaceResultUpdated(log types.Log) (*ITournamentRaceResultUpdated, error) {
	event := new(ITournamentRaceResultUpdated)
	if err := _ITournament.contract.UnpackLog(event, "RaceResultUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ITournamentRegisteredIterator is returned from FilterRegistered and is used to iterate over the raw logs and unpacked data for Registered events raised by the ITournament contract.
type ITournamentRegisteredIterator struct {
	Event *ITournamentRegistered // Event containing the contract specifics and raw log

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
func (it *ITournamentRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ITournamentRegistered)
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
		it.Event = new(ITournamentRegistered)
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
func (it *ITournamentRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ITournamentRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ITournamentRegistered represents a Registered event raised by the ITournament contract.
type ITournamentRegistered struct {
	SlotId      uint8
	Participant common.Address
	RaceId      [32]byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterRegistered is a free log retrieval operation binding the contract event 0xad76e1cb5a0536a10bd2c282409cd1301bc63cff1eeed6b269642f1c24eaaa8f.
//
// Solidity: event Registered(uint8 slotId, address participant, bytes32 raceId)
func (_ITournament *ITournamentFilterer) FilterRegistered(opts *bind.FilterOpts) (*ITournamentRegisteredIterator, error) {

	logs, sub, err := _ITournament.contract.FilterLogs(opts, "Registered")
	if err != nil {
		return nil, err
	}
	return &ITournamentRegisteredIterator{contract: _ITournament.contract, event: "Registered", logs: logs, sub: sub}, nil
}

// WatchRegistered is a free log subscription operation binding the contract event 0xad76e1cb5a0536a10bd2c282409cd1301bc63cff1eeed6b269642f1c24eaaa8f.
//
// Solidity: event Registered(uint8 slotId, address participant, bytes32 raceId)
func (_ITournament *ITournamentFilterer) WatchRegistered(opts *bind.WatchOpts, sink chan<- *ITournamentRegistered) (event.Subscription, error) {

	logs, sub, err := _ITournament.contract.WatchLogs(opts, "Registered")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ITournamentRegistered)
				if err := _ITournament.contract.UnpackLog(event, "Registered", log); err != nil {
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

// ParseRegistered is a log parse operation binding the contract event 0xad76e1cb5a0536a10bd2c282409cd1301bc63cff1eeed6b269642f1c24eaaa8f.
//
// Solidity: event Registered(uint8 slotId, address participant, bytes32 raceId)
func (_ITournament *ITournamentFilterer) ParseRegistered(log types.Log) (*ITournamentRegistered, error) {
	event := new(ITournamentRegistered)
	if err := _ITournament.contract.UnpackLog(event, "Registered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ITournamentRewardGrantedIterator is returned from FilterRewardGranted and is used to iterate over the raw logs and unpacked data for RewardGranted events raised by the ITournament contract.
type ITournamentRewardGrantedIterator struct {
	Event *ITournamentRewardGranted // Event containing the contract specifics and raw log

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
func (it *ITournamentRewardGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ITournamentRewardGranted)
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
		it.Event = new(ITournamentRewardGranted)
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
func (it *ITournamentRewardGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ITournamentRewardGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ITournamentRewardGranted represents a RewardGranted event raised by the ITournament contract.
type ITournamentRewardGranted struct {
	WinnerIndex uint8
	Score       uint8
	Winner      common.Address
	NftRewardId *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterRewardGranted is a free log retrieval operation binding the contract event 0x665f81dad11e5f7be344a68d526bf2281aac8b94951d346d51391531bfcae85e.
//
// Solidity: event RewardGranted(uint8 winnerIndex, uint8 score, address winner, uint256 nftRewardId)
func (_ITournament *ITournamentFilterer) FilterRewardGranted(opts *bind.FilterOpts) (*ITournamentRewardGrantedIterator, error) {

	logs, sub, err := _ITournament.contract.FilterLogs(opts, "RewardGranted")
	if err != nil {
		return nil, err
	}
	return &ITournamentRewardGrantedIterator{contract: _ITournament.contract, event: "RewardGranted", logs: logs, sub: sub}, nil
}

// WatchRewardGranted is a free log subscription operation binding the contract event 0x665f81dad11e5f7be344a68d526bf2281aac8b94951d346d51391531bfcae85e.
//
// Solidity: event RewardGranted(uint8 winnerIndex, uint8 score, address winner, uint256 nftRewardId)
func (_ITournament *ITournamentFilterer) WatchRewardGranted(opts *bind.WatchOpts, sink chan<- *ITournamentRewardGranted) (event.Subscription, error) {

	logs, sub, err := _ITournament.contract.WatchLogs(opts, "RewardGranted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ITournamentRewardGranted)
				if err := _ITournament.contract.UnpackLog(event, "RewardGranted", log); err != nil {
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

// ParseRewardGranted is a log parse operation binding the contract event 0x665f81dad11e5f7be344a68d526bf2281aac8b94951d346d51391531bfcae85e.
//
// Solidity: event RewardGranted(uint8 winnerIndex, uint8 score, address winner, uint256 nftRewardId)
func (_ITournament *ITournamentFilterer) ParseRewardGranted(log types.Log) (*ITournamentRewardGranted, error) {
	event := new(ITournamentRewardGranted)
	if err := _ITournament.contract.UnpackLog(event, "RewardGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
