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
	ABI: "[{\"inputs\":[],\"name\":\"AlreadyReceived\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BetFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FundFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidContract\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidSlot\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RaceNotExisted\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"}],\"name\":\"TooEarly\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"}],\"name\":\"TooLate\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"acceptToken\",\"type\":\"address\"}],\"name\":\"AcceptTokenUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"slotId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"raceId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"betValue\",\"type\":\"uint256\"}],\"name\":\"BetSuccessful\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"raceId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"claimValue\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"ClaimCommission\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"raceId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"claimValue\",\"type\":\"uint256\"}],\"name\":\"ClaimSuccessful\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"raceId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fundValue\",\"type\":\"uint256\"}],\"name\":\"FundSuccessful\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"race\",\"type\":\"address\"}],\"name\":\"RaceListUpdated\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"slotId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"raceId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"betValue\",\"type\":\"uint256\"}],\"name\":\"bet\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"raceId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"fundValue\",\"type\":\"uint256\"}],\"name\":\"fundRace\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"raceId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"slotId\",\"type\":\"uint256\"}],\"name\":\"claim\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"raceId\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"claimCommission\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"raceId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"slotId\",\"type\":\"uint256\"}],\"name\":\"totalSlotBet\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"raceId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"slotId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"userSlotBet\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"raceId\",\"type\":\"bytes32\"}],\"name\":\"totalRaceBet\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"raceId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"slotId\",\"type\":\"uint256\"}],\"name\":\"getSlotPosition\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"raceList\",\"type\":\"address\"}],\"name\":\"updateRaceAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"acceptToken\",\"type\":\"address\"}],\"name\":\"updateAcceptTokenAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
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

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Smc *SmcCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Smc.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Smc *SmcSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Smc.Contract.SupportsInterface(&_Smc.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Smc *SmcCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Smc.Contract.SupportsInterface(&_Smc.CallOpts, interfaceId)
}

// Bet is a paid mutator transaction binding the contract method 0x94fc859b.
//
// Solidity: function bet(uint256 slotId, bytes32 raceId, uint256 betValue) returns()
func (_Smc *SmcTransactor) Bet(opts *bind.TransactOpts, slotId *big.Int, raceId [32]byte, betValue *big.Int) (*types.Transaction, error) {
	return _Smc.contract.Transact(opts, "bet", slotId, raceId, betValue)
}

// Bet is a paid mutator transaction binding the contract method 0x94fc859b.
//
// Solidity: function bet(uint256 slotId, bytes32 raceId, uint256 betValue) returns()
func (_Smc *SmcSession) Bet(slotId *big.Int, raceId [32]byte, betValue *big.Int) (*types.Transaction, error) {
	return _Smc.Contract.Bet(&_Smc.TransactOpts, slotId, raceId, betValue)
}

// Bet is a paid mutator transaction binding the contract method 0x94fc859b.
//
// Solidity: function bet(uint256 slotId, bytes32 raceId, uint256 betValue) returns()
func (_Smc *SmcTransactorSession) Bet(slotId *big.Int, raceId [32]byte, betValue *big.Int) (*types.Transaction, error) {
	return _Smc.Contract.Bet(&_Smc.TransactOpts, slotId, raceId, betValue)
}

// Claim is a paid mutator transaction binding the contract method 0x63f44968.
//
// Solidity: function claim(bytes32 raceId, uint256 slotId) returns()
func (_Smc *SmcTransactor) Claim(opts *bind.TransactOpts, raceId [32]byte, slotId *big.Int) (*types.Transaction, error) {
	return _Smc.contract.Transact(opts, "claim", raceId, slotId)
}

// Claim is a paid mutator transaction binding the contract method 0x63f44968.
//
// Solidity: function claim(bytes32 raceId, uint256 slotId) returns()
func (_Smc *SmcSession) Claim(raceId [32]byte, slotId *big.Int) (*types.Transaction, error) {
	return _Smc.Contract.Claim(&_Smc.TransactOpts, raceId, slotId)
}

// Claim is a paid mutator transaction binding the contract method 0x63f44968.
//
// Solidity: function claim(bytes32 raceId, uint256 slotId) returns()
func (_Smc *SmcTransactorSession) Claim(raceId [32]byte, slotId *big.Int) (*types.Transaction, error) {
	return _Smc.Contract.Claim(&_Smc.TransactOpts, raceId, slotId)
}

// ClaimCommission is a paid mutator transaction binding the contract method 0x5cce2b42.
//
// Solidity: function claimCommission(bytes32 raceId, address receiver) returns()
func (_Smc *SmcTransactor) ClaimCommission(opts *bind.TransactOpts, raceId [32]byte, receiver common.Address) (*types.Transaction, error) {
	return _Smc.contract.Transact(opts, "claimCommission", raceId, receiver)
}

// ClaimCommission is a paid mutator transaction binding the contract method 0x5cce2b42.
//
// Solidity: function claimCommission(bytes32 raceId, address receiver) returns()
func (_Smc *SmcSession) ClaimCommission(raceId [32]byte, receiver common.Address) (*types.Transaction, error) {
	return _Smc.Contract.ClaimCommission(&_Smc.TransactOpts, raceId, receiver)
}

// ClaimCommission is a paid mutator transaction binding the contract method 0x5cce2b42.
//
// Solidity: function claimCommission(bytes32 raceId, address receiver) returns()
func (_Smc *SmcTransactorSession) ClaimCommission(raceId [32]byte, receiver common.Address) (*types.Transaction, error) {
	return _Smc.Contract.ClaimCommission(&_Smc.TransactOpts, raceId, receiver)
}

// FundRace is a paid mutator transaction binding the contract method 0x2ebe4493.
//
// Solidity: function fundRace(bytes32 raceId, uint256 fundValue) returns()
func (_Smc *SmcTransactor) FundRace(opts *bind.TransactOpts, raceId [32]byte, fundValue *big.Int) (*types.Transaction, error) {
	return _Smc.contract.Transact(opts, "fundRace", raceId, fundValue)
}

// FundRace is a paid mutator transaction binding the contract method 0x2ebe4493.
//
// Solidity: function fundRace(bytes32 raceId, uint256 fundValue) returns()
func (_Smc *SmcSession) FundRace(raceId [32]byte, fundValue *big.Int) (*types.Transaction, error) {
	return _Smc.Contract.FundRace(&_Smc.TransactOpts, raceId, fundValue)
}

// FundRace is a paid mutator transaction binding the contract method 0x2ebe4493.
//
// Solidity: function fundRace(bytes32 raceId, uint256 fundValue) returns()
func (_Smc *SmcTransactorSession) FundRace(raceId [32]byte, fundValue *big.Int) (*types.Transaction, error) {
	return _Smc.Contract.FundRace(&_Smc.TransactOpts, raceId, fundValue)
}

// GetSlotPosition is a paid mutator transaction binding the contract method 0xd0b57c69.
//
// Solidity: function getSlotPosition(bytes32 raceId, uint256 slotId) returns(uint8)
func (_Smc *SmcTransactor) GetSlotPosition(opts *bind.TransactOpts, raceId [32]byte, slotId *big.Int) (*types.Transaction, error) {
	return _Smc.contract.Transact(opts, "getSlotPosition", raceId, slotId)
}

// GetSlotPosition is a paid mutator transaction binding the contract method 0xd0b57c69.
//
// Solidity: function getSlotPosition(bytes32 raceId, uint256 slotId) returns(uint8)
func (_Smc *SmcSession) GetSlotPosition(raceId [32]byte, slotId *big.Int) (*types.Transaction, error) {
	return _Smc.Contract.GetSlotPosition(&_Smc.TransactOpts, raceId, slotId)
}

// GetSlotPosition is a paid mutator transaction binding the contract method 0xd0b57c69.
//
// Solidity: function getSlotPosition(bytes32 raceId, uint256 slotId) returns(uint8)
func (_Smc *SmcTransactorSession) GetSlotPosition(raceId [32]byte, slotId *big.Int) (*types.Transaction, error) {
	return _Smc.Contract.GetSlotPosition(&_Smc.TransactOpts, raceId, slotId)
}

// TotalRaceBet is a paid mutator transaction binding the contract method 0x2123c15e.
//
// Solidity: function totalRaceBet(bytes32 raceId) returns(uint256)
func (_Smc *SmcTransactor) TotalRaceBet(opts *bind.TransactOpts, raceId [32]byte) (*types.Transaction, error) {
	return _Smc.contract.Transact(opts, "totalRaceBet", raceId)
}

// TotalRaceBet is a paid mutator transaction binding the contract method 0x2123c15e.
//
// Solidity: function totalRaceBet(bytes32 raceId) returns(uint256)
func (_Smc *SmcSession) TotalRaceBet(raceId [32]byte) (*types.Transaction, error) {
	return _Smc.Contract.TotalRaceBet(&_Smc.TransactOpts, raceId)
}

// TotalRaceBet is a paid mutator transaction binding the contract method 0x2123c15e.
//
// Solidity: function totalRaceBet(bytes32 raceId) returns(uint256)
func (_Smc *SmcTransactorSession) TotalRaceBet(raceId [32]byte) (*types.Transaction, error) {
	return _Smc.Contract.TotalRaceBet(&_Smc.TransactOpts, raceId)
}

// TotalSlotBet is a paid mutator transaction binding the contract method 0xe740582a.
//
// Solidity: function totalSlotBet(bytes32 raceId, uint256 slotId) returns(uint256)
func (_Smc *SmcTransactor) TotalSlotBet(opts *bind.TransactOpts, raceId [32]byte, slotId *big.Int) (*types.Transaction, error) {
	return _Smc.contract.Transact(opts, "totalSlotBet", raceId, slotId)
}

// TotalSlotBet is a paid mutator transaction binding the contract method 0xe740582a.
//
// Solidity: function totalSlotBet(bytes32 raceId, uint256 slotId) returns(uint256)
func (_Smc *SmcSession) TotalSlotBet(raceId [32]byte, slotId *big.Int) (*types.Transaction, error) {
	return _Smc.Contract.TotalSlotBet(&_Smc.TransactOpts, raceId, slotId)
}

// TotalSlotBet is a paid mutator transaction binding the contract method 0xe740582a.
//
// Solidity: function totalSlotBet(bytes32 raceId, uint256 slotId) returns(uint256)
func (_Smc *SmcTransactorSession) TotalSlotBet(raceId [32]byte, slotId *big.Int) (*types.Transaction, error) {
	return _Smc.Contract.TotalSlotBet(&_Smc.TransactOpts, raceId, slotId)
}

// UpdateAcceptTokenAddress is a paid mutator transaction binding the contract method 0xb4a2fad0.
//
// Solidity: function updateAcceptTokenAddress(address acceptToken) returns()
func (_Smc *SmcTransactor) UpdateAcceptTokenAddress(opts *bind.TransactOpts, acceptToken common.Address) (*types.Transaction, error) {
	return _Smc.contract.Transact(opts, "updateAcceptTokenAddress", acceptToken)
}

// UpdateAcceptTokenAddress is a paid mutator transaction binding the contract method 0xb4a2fad0.
//
// Solidity: function updateAcceptTokenAddress(address acceptToken) returns()
func (_Smc *SmcSession) UpdateAcceptTokenAddress(acceptToken common.Address) (*types.Transaction, error) {
	return _Smc.Contract.UpdateAcceptTokenAddress(&_Smc.TransactOpts, acceptToken)
}

// UpdateAcceptTokenAddress is a paid mutator transaction binding the contract method 0xb4a2fad0.
//
// Solidity: function updateAcceptTokenAddress(address acceptToken) returns()
func (_Smc *SmcTransactorSession) UpdateAcceptTokenAddress(acceptToken common.Address) (*types.Transaction, error) {
	return _Smc.Contract.UpdateAcceptTokenAddress(&_Smc.TransactOpts, acceptToken)
}

// UpdateRaceAddress is a paid mutator transaction binding the contract method 0xd7c49938.
//
// Solidity: function updateRaceAddress(address raceList) returns()
func (_Smc *SmcTransactor) UpdateRaceAddress(opts *bind.TransactOpts, raceList common.Address) (*types.Transaction, error) {
	return _Smc.contract.Transact(opts, "updateRaceAddress", raceList)
}

// UpdateRaceAddress is a paid mutator transaction binding the contract method 0xd7c49938.
//
// Solidity: function updateRaceAddress(address raceList) returns()
func (_Smc *SmcSession) UpdateRaceAddress(raceList common.Address) (*types.Transaction, error) {
	return _Smc.Contract.UpdateRaceAddress(&_Smc.TransactOpts, raceList)
}

// UpdateRaceAddress is a paid mutator transaction binding the contract method 0xd7c49938.
//
// Solidity: function updateRaceAddress(address raceList) returns()
func (_Smc *SmcTransactorSession) UpdateRaceAddress(raceList common.Address) (*types.Transaction, error) {
	return _Smc.Contract.UpdateRaceAddress(&_Smc.TransactOpts, raceList)
}

// UserSlotBet is a paid mutator transaction binding the contract method 0x2e279801.
//
// Solidity: function userSlotBet(bytes32 raceId, uint256 slotId, address user) returns(uint256)
func (_Smc *SmcTransactor) UserSlotBet(opts *bind.TransactOpts, raceId [32]byte, slotId *big.Int, user common.Address) (*types.Transaction, error) {
	return _Smc.contract.Transact(opts, "userSlotBet", raceId, slotId, user)
}

// UserSlotBet is a paid mutator transaction binding the contract method 0x2e279801.
//
// Solidity: function userSlotBet(bytes32 raceId, uint256 slotId, address user) returns(uint256)
func (_Smc *SmcSession) UserSlotBet(raceId [32]byte, slotId *big.Int, user common.Address) (*types.Transaction, error) {
	return _Smc.Contract.UserSlotBet(&_Smc.TransactOpts, raceId, slotId, user)
}

// UserSlotBet is a paid mutator transaction binding the contract method 0x2e279801.
//
// Solidity: function userSlotBet(bytes32 raceId, uint256 slotId, address user) returns(uint256)
func (_Smc *SmcTransactorSession) UserSlotBet(raceId [32]byte, slotId *big.Int, user common.Address) (*types.Transaction, error) {
	return _Smc.Contract.UserSlotBet(&_Smc.TransactOpts, raceId, slotId, user)
}

// SmcAcceptTokenUpdatedIterator is returned from FilterAcceptTokenUpdated and is used to iterate over the raw logs and unpacked data for AcceptTokenUpdated events raised by the Smc contract.
type SmcAcceptTokenUpdatedIterator struct {
	Event *SmcAcceptTokenUpdated // Event containing the contract specifics and raw log

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
func (it *SmcAcceptTokenUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SmcAcceptTokenUpdated)
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
		it.Event = new(SmcAcceptTokenUpdated)
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
func (it *SmcAcceptTokenUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SmcAcceptTokenUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SmcAcceptTokenUpdated represents a AcceptTokenUpdated event raised by the Smc contract.
type SmcAcceptTokenUpdated struct {
	AcceptToken common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterAcceptTokenUpdated is a free log retrieval operation binding the contract event 0x44c3b193081dd50818bc919dc1963ae03762174c9a89af5fbbce54ee71a13f97.
//
// Solidity: event AcceptTokenUpdated(address acceptToken)
func (_Smc *SmcFilterer) FilterAcceptTokenUpdated(opts *bind.FilterOpts) (*SmcAcceptTokenUpdatedIterator, error) {

	logs, sub, err := _Smc.contract.FilterLogs(opts, "AcceptTokenUpdated")
	if err != nil {
		return nil, err
	}
	return &SmcAcceptTokenUpdatedIterator{contract: _Smc.contract, event: "AcceptTokenUpdated", logs: logs, sub: sub}, nil
}

// WatchAcceptTokenUpdated is a free log subscription operation binding the contract event 0x44c3b193081dd50818bc919dc1963ae03762174c9a89af5fbbce54ee71a13f97.
//
// Solidity: event AcceptTokenUpdated(address acceptToken)
func (_Smc *SmcFilterer) WatchAcceptTokenUpdated(opts *bind.WatchOpts, sink chan<- *SmcAcceptTokenUpdated) (event.Subscription, error) {

	logs, sub, err := _Smc.contract.WatchLogs(opts, "AcceptTokenUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SmcAcceptTokenUpdated)
				if err := _Smc.contract.UnpackLog(event, "AcceptTokenUpdated", log); err != nil {
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

// ParseAcceptTokenUpdated is a log parse operation binding the contract event 0x44c3b193081dd50818bc919dc1963ae03762174c9a89af5fbbce54ee71a13f97.
//
// Solidity: event AcceptTokenUpdated(address acceptToken)
func (_Smc *SmcFilterer) ParseAcceptTokenUpdated(log types.Log) (*SmcAcceptTokenUpdated, error) {
	event := new(SmcAcceptTokenUpdated)
	if err := _Smc.contract.UnpackLog(event, "AcceptTokenUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SmcBetSuccessfulIterator is returned from FilterBetSuccessful and is used to iterate over the raw logs and unpacked data for BetSuccessful events raised by the Smc contract.
type SmcBetSuccessfulIterator struct {
	Event *SmcBetSuccessful // Event containing the contract specifics and raw log

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
func (it *SmcBetSuccessfulIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SmcBetSuccessful)
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
		it.Event = new(SmcBetSuccessful)
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
func (it *SmcBetSuccessfulIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SmcBetSuccessfulIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SmcBetSuccessful represents a BetSuccessful event raised by the Smc contract.
type SmcBetSuccessful struct {
	SlotId   *big.Int
	RaceId   [32]byte
	BetValue *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterBetSuccessful is a free log retrieval operation binding the contract event 0x90340911801212a3d95efe0dca9cdbd74549020af8cf33fc0861b74d82fe6dbd.
//
// Solidity: event BetSuccessful(uint256 slotId, bytes32 raceId, uint256 betValue)
func (_Smc *SmcFilterer) FilterBetSuccessful(opts *bind.FilterOpts) (*SmcBetSuccessfulIterator, error) {

	logs, sub, err := _Smc.contract.FilterLogs(opts, "BetSuccessful")
	if err != nil {
		return nil, err
	}
	return &SmcBetSuccessfulIterator{contract: _Smc.contract, event: "BetSuccessful", logs: logs, sub: sub}, nil
}

// WatchBetSuccessful is a free log subscription operation binding the contract event 0x90340911801212a3d95efe0dca9cdbd74549020af8cf33fc0861b74d82fe6dbd.
//
// Solidity: event BetSuccessful(uint256 slotId, bytes32 raceId, uint256 betValue)
func (_Smc *SmcFilterer) WatchBetSuccessful(opts *bind.WatchOpts, sink chan<- *SmcBetSuccessful) (event.Subscription, error) {

	logs, sub, err := _Smc.contract.WatchLogs(opts, "BetSuccessful")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SmcBetSuccessful)
				if err := _Smc.contract.UnpackLog(event, "BetSuccessful", log); err != nil {
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

// ParseBetSuccessful is a log parse operation binding the contract event 0x90340911801212a3d95efe0dca9cdbd74549020af8cf33fc0861b74d82fe6dbd.
//
// Solidity: event BetSuccessful(uint256 slotId, bytes32 raceId, uint256 betValue)
func (_Smc *SmcFilterer) ParseBetSuccessful(log types.Log) (*SmcBetSuccessful, error) {
	event := new(SmcBetSuccessful)
	if err := _Smc.contract.UnpackLog(event, "BetSuccessful", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SmcClaimCommissionIterator is returned from FilterClaimCommission and is used to iterate over the raw logs and unpacked data for ClaimCommission events raised by the Smc contract.
type SmcClaimCommissionIterator struct {
	Event *SmcClaimCommission // Event containing the contract specifics and raw log

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
func (it *SmcClaimCommissionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SmcClaimCommission)
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
		it.Event = new(SmcClaimCommission)
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
func (it *SmcClaimCommissionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SmcClaimCommissionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SmcClaimCommission represents a ClaimCommission event raised by the Smc contract.
type SmcClaimCommission struct {
	RaceId     [32]byte
	ClaimValue *big.Int
	Receiver   common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterClaimCommission is a free log retrieval operation binding the contract event 0x94fe36569bb216f6cb564fee8983d035c7cacba5679023a4ccabe33fdf6d4d88.
//
// Solidity: event ClaimCommission(bytes32 raceId, uint256 claimValue, address receiver)
func (_Smc *SmcFilterer) FilterClaimCommission(opts *bind.FilterOpts) (*SmcClaimCommissionIterator, error) {

	logs, sub, err := _Smc.contract.FilterLogs(opts, "ClaimCommission")
	if err != nil {
		return nil, err
	}
	return &SmcClaimCommissionIterator{contract: _Smc.contract, event: "ClaimCommission", logs: logs, sub: sub}, nil
}

// WatchClaimCommission is a free log subscription operation binding the contract event 0x94fe36569bb216f6cb564fee8983d035c7cacba5679023a4ccabe33fdf6d4d88.
//
// Solidity: event ClaimCommission(bytes32 raceId, uint256 claimValue, address receiver)
func (_Smc *SmcFilterer) WatchClaimCommission(opts *bind.WatchOpts, sink chan<- *SmcClaimCommission) (event.Subscription, error) {

	logs, sub, err := _Smc.contract.WatchLogs(opts, "ClaimCommission")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SmcClaimCommission)
				if err := _Smc.contract.UnpackLog(event, "ClaimCommission", log); err != nil {
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

// ParseClaimCommission is a log parse operation binding the contract event 0x94fe36569bb216f6cb564fee8983d035c7cacba5679023a4ccabe33fdf6d4d88.
//
// Solidity: event ClaimCommission(bytes32 raceId, uint256 claimValue, address receiver)
func (_Smc *SmcFilterer) ParseClaimCommission(log types.Log) (*SmcClaimCommission, error) {
	event := new(SmcClaimCommission)
	if err := _Smc.contract.UnpackLog(event, "ClaimCommission", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SmcClaimSuccessfulIterator is returned from FilterClaimSuccessful and is used to iterate over the raw logs and unpacked data for ClaimSuccessful events raised by the Smc contract.
type SmcClaimSuccessfulIterator struct {
	Event *SmcClaimSuccessful // Event containing the contract specifics and raw log

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
func (it *SmcClaimSuccessfulIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SmcClaimSuccessful)
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
		it.Event = new(SmcClaimSuccessful)
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
func (it *SmcClaimSuccessfulIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SmcClaimSuccessfulIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SmcClaimSuccessful represents a ClaimSuccessful event raised by the Smc contract.
type SmcClaimSuccessful struct {
	RaceId     [32]byte
	ClaimValue *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterClaimSuccessful is a free log retrieval operation binding the contract event 0x78d224c5b360f08891a0d6c82ec1fceea8a4e5b39dc51c1373ddfaa13d848df0.
//
// Solidity: event ClaimSuccessful(bytes32 raceId, uint256 claimValue)
func (_Smc *SmcFilterer) FilterClaimSuccessful(opts *bind.FilterOpts) (*SmcClaimSuccessfulIterator, error) {

	logs, sub, err := _Smc.contract.FilterLogs(opts, "ClaimSuccessful")
	if err != nil {
		return nil, err
	}
	return &SmcClaimSuccessfulIterator{contract: _Smc.contract, event: "ClaimSuccessful", logs: logs, sub: sub}, nil
}

// WatchClaimSuccessful is a free log subscription operation binding the contract event 0x78d224c5b360f08891a0d6c82ec1fceea8a4e5b39dc51c1373ddfaa13d848df0.
//
// Solidity: event ClaimSuccessful(bytes32 raceId, uint256 claimValue)
func (_Smc *SmcFilterer) WatchClaimSuccessful(opts *bind.WatchOpts, sink chan<- *SmcClaimSuccessful) (event.Subscription, error) {

	logs, sub, err := _Smc.contract.WatchLogs(opts, "ClaimSuccessful")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SmcClaimSuccessful)
				if err := _Smc.contract.UnpackLog(event, "ClaimSuccessful", log); err != nil {
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

// ParseClaimSuccessful is a log parse operation binding the contract event 0x78d224c5b360f08891a0d6c82ec1fceea8a4e5b39dc51c1373ddfaa13d848df0.
//
// Solidity: event ClaimSuccessful(bytes32 raceId, uint256 claimValue)
func (_Smc *SmcFilterer) ParseClaimSuccessful(log types.Log) (*SmcClaimSuccessful, error) {
	event := new(SmcClaimSuccessful)
	if err := _Smc.contract.UnpackLog(event, "ClaimSuccessful", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SmcFundSuccessfulIterator is returned from FilterFundSuccessful and is used to iterate over the raw logs and unpacked data for FundSuccessful events raised by the Smc contract.
type SmcFundSuccessfulIterator struct {
	Event *SmcFundSuccessful // Event containing the contract specifics and raw log

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
func (it *SmcFundSuccessfulIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SmcFundSuccessful)
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
		it.Event = new(SmcFundSuccessful)
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
func (it *SmcFundSuccessfulIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SmcFundSuccessfulIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SmcFundSuccessful represents a FundSuccessful event raised by the Smc contract.
type SmcFundSuccessful struct {
	RaceId    [32]byte
	FundValue *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterFundSuccessful is a free log retrieval operation binding the contract event 0x18828fa7a168b0db08e3e72ef21348f4c4fa1ee1b46803878b1195794aaad4c3.
//
// Solidity: event FundSuccessful(bytes32 raceId, uint256 fundValue)
func (_Smc *SmcFilterer) FilterFundSuccessful(opts *bind.FilterOpts) (*SmcFundSuccessfulIterator, error) {

	logs, sub, err := _Smc.contract.FilterLogs(opts, "FundSuccessful")
	if err != nil {
		return nil, err
	}
	return &SmcFundSuccessfulIterator{contract: _Smc.contract, event: "FundSuccessful", logs: logs, sub: sub}, nil
}

// WatchFundSuccessful is a free log subscription operation binding the contract event 0x18828fa7a168b0db08e3e72ef21348f4c4fa1ee1b46803878b1195794aaad4c3.
//
// Solidity: event FundSuccessful(bytes32 raceId, uint256 fundValue)
func (_Smc *SmcFilterer) WatchFundSuccessful(opts *bind.WatchOpts, sink chan<- *SmcFundSuccessful) (event.Subscription, error) {

	logs, sub, err := _Smc.contract.WatchLogs(opts, "FundSuccessful")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SmcFundSuccessful)
				if err := _Smc.contract.UnpackLog(event, "FundSuccessful", log); err != nil {
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

// ParseFundSuccessful is a log parse operation binding the contract event 0x18828fa7a168b0db08e3e72ef21348f4c4fa1ee1b46803878b1195794aaad4c3.
//
// Solidity: event FundSuccessful(bytes32 raceId, uint256 fundValue)
func (_Smc *SmcFilterer) ParseFundSuccessful(log types.Log) (*SmcFundSuccessful, error) {
	event := new(SmcFundSuccessful)
	if err := _Smc.contract.UnpackLog(event, "FundSuccessful", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SmcRaceListUpdatedIterator is returned from FilterRaceListUpdated and is used to iterate over the raw logs and unpacked data for RaceListUpdated events raised by the Smc contract.
type SmcRaceListUpdatedIterator struct {
	Event *SmcRaceListUpdated // Event containing the contract specifics and raw log

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
func (it *SmcRaceListUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SmcRaceListUpdated)
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
		it.Event = new(SmcRaceListUpdated)
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
func (it *SmcRaceListUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SmcRaceListUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SmcRaceListUpdated represents a RaceListUpdated event raised by the Smc contract.
type SmcRaceListUpdated struct {
	Race common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterRaceListUpdated is a free log retrieval operation binding the contract event 0x7a659342c2b03e1c14729e3f1c86f414756d13bd8ff0636713cc22246182b42d.
//
// Solidity: event RaceListUpdated(address race)
func (_Smc *SmcFilterer) FilterRaceListUpdated(opts *bind.FilterOpts) (*SmcRaceListUpdatedIterator, error) {

	logs, sub, err := _Smc.contract.FilterLogs(opts, "RaceListUpdated")
	if err != nil {
		return nil, err
	}
	return &SmcRaceListUpdatedIterator{contract: _Smc.contract, event: "RaceListUpdated", logs: logs, sub: sub}, nil
}

// WatchRaceListUpdated is a free log subscription operation binding the contract event 0x7a659342c2b03e1c14729e3f1c86f414756d13bd8ff0636713cc22246182b42d.
//
// Solidity: event RaceListUpdated(address race)
func (_Smc *SmcFilterer) WatchRaceListUpdated(opts *bind.WatchOpts, sink chan<- *SmcRaceListUpdated) (event.Subscription, error) {

	logs, sub, err := _Smc.contract.WatchLogs(opts, "RaceListUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SmcRaceListUpdated)
				if err := _Smc.contract.UnpackLog(event, "RaceListUpdated", log); err != nil {
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
func (_Smc *SmcFilterer) ParseRaceListUpdated(log types.Log) (*SmcRaceListUpdated, error) {
	event := new(SmcRaceListUpdated)
	if err := _Smc.contract.UnpackLog(event, "RaceListUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
