// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package IBething

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

// IBethingMetaData contains all meta data concerning the IBething contract.
var IBethingMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"AlreadyReceived\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BetFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FundFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidContract\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidSlot\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RaceNotExisted\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"}],\"name\":\"TooEarly\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"}],\"name\":\"TooLate\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"acceptToken\",\"type\":\"address\"}],\"name\":\"AcceptTokenUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"slotId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"raceId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"betValue\",\"type\":\"uint256\"}],\"name\":\"BetSuccessful\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"raceId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"claimValue\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"ClaimCommission\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"raceId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"claimValue\",\"type\":\"uint256\"}],\"name\":\"ClaimSuccessful\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"raceId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fundValue\",\"type\":\"uint256\"}],\"name\":\"FundSuccessful\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"race\",\"type\":\"address\"}],\"name\":\"RaceListUpdated\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"slotId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"raceId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"betValue\",\"type\":\"uint256\"}],\"name\":\"bet\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"raceId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"fundValue\",\"type\":\"uint256\"}],\"name\":\"fundRace\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"raceId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"slotId\",\"type\":\"uint256\"}],\"name\":\"claim\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"raceId\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"claimCommission\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"raceId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"slotId\",\"type\":\"uint256\"}],\"name\":\"totalSlotBet\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"raceId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"slotId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"userSlotBet\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"raceId\",\"type\":\"bytes32\"}],\"name\":\"totalRaceBet\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"raceId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"slotId\",\"type\":\"uint256\"}],\"name\":\"getSlotPosition\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"raceList\",\"type\":\"address\"}],\"name\":\"updateRaceAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"acceptToken\",\"type\":\"address\"}],\"name\":\"updateAcceptTokenAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// IBethingABI is the input ABI used to generate the binding from.
// Deprecated: Use IBethingMetaData.ABI instead.
var IBethingABI = IBethingMetaData.ABI

// IBething is an auto generated Go binding around an Ethereum contract.
type IBething struct {
	IBethingCaller     // Read-only binding to the contract
	IBethingTransactor // Write-only binding to the contract
	IBethingFilterer   // Log filterer for contract events
}

// IBethingCaller is an auto generated read-only Go binding around an Ethereum contract.
type IBethingCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IBethingTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IBethingTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IBethingFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IBethingFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IBethingSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IBethingSession struct {
	Contract     *IBething         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IBethingCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IBethingCallerSession struct {
	Contract *IBethingCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// IBethingTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IBethingTransactorSession struct {
	Contract     *IBethingTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// IBethingRaw is an auto generated low-level Go binding around an Ethereum contract.
type IBethingRaw struct {
	Contract *IBething // Generic contract binding to access the raw methods on
}

// IBethingCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IBethingCallerRaw struct {
	Contract *IBethingCaller // Generic read-only contract binding to access the raw methods on
}

// IBethingTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IBethingTransactorRaw struct {
	Contract *IBethingTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIBething creates a new instance of IBething, bound to a specific deployed contract.
func NewIBething(address common.Address, backend bind.ContractBackend) (*IBething, error) {
	contract, err := bindIBething(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IBething{IBethingCaller: IBethingCaller{contract: contract}, IBethingTransactor: IBethingTransactor{contract: contract}, IBethingFilterer: IBethingFilterer{contract: contract}}, nil
}

// NewIBethingCaller creates a new read-only instance of IBething, bound to a specific deployed contract.
func NewIBethingCaller(address common.Address, caller bind.ContractCaller) (*IBethingCaller, error) {
	contract, err := bindIBething(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IBethingCaller{contract: contract}, nil
}

// NewIBethingTransactor creates a new write-only instance of IBething, bound to a specific deployed contract.
func NewIBethingTransactor(address common.Address, transactor bind.ContractTransactor) (*IBethingTransactor, error) {
	contract, err := bindIBething(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IBethingTransactor{contract: contract}, nil
}

// NewIBethingFilterer creates a new log filterer instance of IBething, bound to a specific deployed contract.
func NewIBethingFilterer(address common.Address, filterer bind.ContractFilterer) (*IBethingFilterer, error) {
	contract, err := bindIBething(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IBethingFilterer{contract: contract}, nil
}

// bindIBething binds a generic wrapper to an already deployed contract.
func bindIBething(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IBethingABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IBething *IBethingRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IBething.Contract.IBethingCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IBething *IBethingRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IBething.Contract.IBethingTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IBething *IBethingRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IBething.Contract.IBethingTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IBething *IBethingCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IBething.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IBething *IBethingTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IBething.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IBething *IBethingTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IBething.Contract.contract.Transact(opts, method, params...)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_IBething *IBethingCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _IBething.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_IBething *IBethingSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _IBething.Contract.SupportsInterface(&_IBething.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_IBething *IBethingCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _IBething.Contract.SupportsInterface(&_IBething.CallOpts, interfaceId)
}

// Bet is a paid mutator transaction binding the contract method 0x94fc859b.
//
// Solidity: function bet(uint256 slotId, bytes32 raceId, uint256 betValue) returns()
func (_IBething *IBethingTransactor) Bet(opts *bind.TransactOpts, slotId *big.Int, raceId [32]byte, betValue *big.Int) (*types.Transaction, error) {
	return _IBething.contract.Transact(opts, "bet", slotId, raceId, betValue)
}

// Bet is a paid mutator transaction binding the contract method 0x94fc859b.
//
// Solidity: function bet(uint256 slotId, bytes32 raceId, uint256 betValue) returns()
func (_IBething *IBethingSession) Bet(slotId *big.Int, raceId [32]byte, betValue *big.Int) (*types.Transaction, error) {
	return _IBething.Contract.Bet(&_IBething.TransactOpts, slotId, raceId, betValue)
}

// Bet is a paid mutator transaction binding the contract method 0x94fc859b.
//
// Solidity: function bet(uint256 slotId, bytes32 raceId, uint256 betValue) returns()
func (_IBething *IBethingTransactorSession) Bet(slotId *big.Int, raceId [32]byte, betValue *big.Int) (*types.Transaction, error) {
	return _IBething.Contract.Bet(&_IBething.TransactOpts, slotId, raceId, betValue)
}

// Claim is a paid mutator transaction binding the contract method 0x63f44968.
//
// Solidity: function claim(bytes32 raceId, uint256 slotId) returns()
func (_IBething *IBethingTransactor) Claim(opts *bind.TransactOpts, raceId [32]byte, slotId *big.Int) (*types.Transaction, error) {
	return _IBething.contract.Transact(opts, "claim", raceId, slotId)
}

// Claim is a paid mutator transaction binding the contract method 0x63f44968.
//
// Solidity: function claim(bytes32 raceId, uint256 slotId) returns()
func (_IBething *IBethingSession) Claim(raceId [32]byte, slotId *big.Int) (*types.Transaction, error) {
	return _IBething.Contract.Claim(&_IBething.TransactOpts, raceId, slotId)
}

// Claim is a paid mutator transaction binding the contract method 0x63f44968.
//
// Solidity: function claim(bytes32 raceId, uint256 slotId) returns()
func (_IBething *IBethingTransactorSession) Claim(raceId [32]byte, slotId *big.Int) (*types.Transaction, error) {
	return _IBething.Contract.Claim(&_IBething.TransactOpts, raceId, slotId)
}

// ClaimCommission is a paid mutator transaction binding the contract method 0x5cce2b42.
//
// Solidity: function claimCommission(bytes32 raceId, address receiver) returns()
func (_IBething *IBethingTransactor) ClaimCommission(opts *bind.TransactOpts, raceId [32]byte, receiver common.Address) (*types.Transaction, error) {
	return _IBething.contract.Transact(opts, "claimCommission", raceId, receiver)
}

// ClaimCommission is a paid mutator transaction binding the contract method 0x5cce2b42.
//
// Solidity: function claimCommission(bytes32 raceId, address receiver) returns()
func (_IBething *IBethingSession) ClaimCommission(raceId [32]byte, receiver common.Address) (*types.Transaction, error) {
	return _IBething.Contract.ClaimCommission(&_IBething.TransactOpts, raceId, receiver)
}

// ClaimCommission is a paid mutator transaction binding the contract method 0x5cce2b42.
//
// Solidity: function claimCommission(bytes32 raceId, address receiver) returns()
func (_IBething *IBethingTransactorSession) ClaimCommission(raceId [32]byte, receiver common.Address) (*types.Transaction, error) {
	return _IBething.Contract.ClaimCommission(&_IBething.TransactOpts, raceId, receiver)
}

// FundRace is a paid mutator transaction binding the contract method 0x2ebe4493.
//
// Solidity: function fundRace(bytes32 raceId, uint256 fundValue) returns()
func (_IBething *IBethingTransactor) FundRace(opts *bind.TransactOpts, raceId [32]byte, fundValue *big.Int) (*types.Transaction, error) {
	return _IBething.contract.Transact(opts, "fundRace", raceId, fundValue)
}

// FundRace is a paid mutator transaction binding the contract method 0x2ebe4493.
//
// Solidity: function fundRace(bytes32 raceId, uint256 fundValue) returns()
func (_IBething *IBethingSession) FundRace(raceId [32]byte, fundValue *big.Int) (*types.Transaction, error) {
	return _IBething.Contract.FundRace(&_IBething.TransactOpts, raceId, fundValue)
}

// FundRace is a paid mutator transaction binding the contract method 0x2ebe4493.
//
// Solidity: function fundRace(bytes32 raceId, uint256 fundValue) returns()
func (_IBething *IBethingTransactorSession) FundRace(raceId [32]byte, fundValue *big.Int) (*types.Transaction, error) {
	return _IBething.Contract.FundRace(&_IBething.TransactOpts, raceId, fundValue)
}

// GetSlotPosition is a paid mutator transaction binding the contract method 0xd0b57c69.
//
// Solidity: function getSlotPosition(bytes32 raceId, uint256 slotId) returns(uint8)
func (_IBething *IBethingTransactor) GetSlotPosition(opts *bind.TransactOpts, raceId [32]byte, slotId *big.Int) (*types.Transaction, error) {
	return _IBething.contract.Transact(opts, "getSlotPosition", raceId, slotId)
}

// GetSlotPosition is a paid mutator transaction binding the contract method 0xd0b57c69.
//
// Solidity: function getSlotPosition(bytes32 raceId, uint256 slotId) returns(uint8)
func (_IBething *IBethingSession) GetSlotPosition(raceId [32]byte, slotId *big.Int) (*types.Transaction, error) {
	return _IBething.Contract.GetSlotPosition(&_IBething.TransactOpts, raceId, slotId)
}

// GetSlotPosition is a paid mutator transaction binding the contract method 0xd0b57c69.
//
// Solidity: function getSlotPosition(bytes32 raceId, uint256 slotId) returns(uint8)
func (_IBething *IBethingTransactorSession) GetSlotPosition(raceId [32]byte, slotId *big.Int) (*types.Transaction, error) {
	return _IBething.Contract.GetSlotPosition(&_IBething.TransactOpts, raceId, slotId)
}

// TotalRaceBet is a paid mutator transaction binding the contract method 0x2123c15e.
//
// Solidity: function totalRaceBet(bytes32 raceId) returns(uint256)
func (_IBething *IBethingTransactor) TotalRaceBet(opts *bind.TransactOpts, raceId [32]byte) (*types.Transaction, error) {
	return _IBething.contract.Transact(opts, "totalRaceBet", raceId)
}

// TotalRaceBet is a paid mutator transaction binding the contract method 0x2123c15e.
//
// Solidity: function totalRaceBet(bytes32 raceId) returns(uint256)
func (_IBething *IBethingSession) TotalRaceBet(raceId [32]byte) (*types.Transaction, error) {
	return _IBething.Contract.TotalRaceBet(&_IBething.TransactOpts, raceId)
}

// TotalRaceBet is a paid mutator transaction binding the contract method 0x2123c15e.
//
// Solidity: function totalRaceBet(bytes32 raceId) returns(uint256)
func (_IBething *IBethingTransactorSession) TotalRaceBet(raceId [32]byte) (*types.Transaction, error) {
	return _IBething.Contract.TotalRaceBet(&_IBething.TransactOpts, raceId)
}

// TotalSlotBet is a paid mutator transaction binding the contract method 0xe740582a.
//
// Solidity: function totalSlotBet(bytes32 raceId, uint256 slotId) returns(uint256)
func (_IBething *IBethingTransactor) TotalSlotBet(opts *bind.TransactOpts, raceId [32]byte, slotId *big.Int) (*types.Transaction, error) {
	return _IBething.contract.Transact(opts, "totalSlotBet", raceId, slotId)
}

// TotalSlotBet is a paid mutator transaction binding the contract method 0xe740582a.
//
// Solidity: function totalSlotBet(bytes32 raceId, uint256 slotId) returns(uint256)
func (_IBething *IBethingSession) TotalSlotBet(raceId [32]byte, slotId *big.Int) (*types.Transaction, error) {
	return _IBething.Contract.TotalSlotBet(&_IBething.TransactOpts, raceId, slotId)
}

// TotalSlotBet is a paid mutator transaction binding the contract method 0xe740582a.
//
// Solidity: function totalSlotBet(bytes32 raceId, uint256 slotId) returns(uint256)
func (_IBething *IBethingTransactorSession) TotalSlotBet(raceId [32]byte, slotId *big.Int) (*types.Transaction, error) {
	return _IBething.Contract.TotalSlotBet(&_IBething.TransactOpts, raceId, slotId)
}

// UpdateAcceptTokenAddress is a paid mutator transaction binding the contract method 0xb4a2fad0.
//
// Solidity: function updateAcceptTokenAddress(address acceptToken) returns()
func (_IBething *IBethingTransactor) UpdateAcceptTokenAddress(opts *bind.TransactOpts, acceptToken common.Address) (*types.Transaction, error) {
	return _IBething.contract.Transact(opts, "updateAcceptTokenAddress", acceptToken)
}

// UpdateAcceptTokenAddress is a paid mutator transaction binding the contract method 0xb4a2fad0.
//
// Solidity: function updateAcceptTokenAddress(address acceptToken) returns()
func (_IBething *IBethingSession) UpdateAcceptTokenAddress(acceptToken common.Address) (*types.Transaction, error) {
	return _IBething.Contract.UpdateAcceptTokenAddress(&_IBething.TransactOpts, acceptToken)
}

// UpdateAcceptTokenAddress is a paid mutator transaction binding the contract method 0xb4a2fad0.
//
// Solidity: function updateAcceptTokenAddress(address acceptToken) returns()
func (_IBething *IBethingTransactorSession) UpdateAcceptTokenAddress(acceptToken common.Address) (*types.Transaction, error) {
	return _IBething.Contract.UpdateAcceptTokenAddress(&_IBething.TransactOpts, acceptToken)
}

// UpdateRaceAddress is a paid mutator transaction binding the contract method 0xd7c49938.
//
// Solidity: function updateRaceAddress(address raceList) returns()
func (_IBething *IBethingTransactor) UpdateRaceAddress(opts *bind.TransactOpts, raceList common.Address) (*types.Transaction, error) {
	return _IBething.contract.Transact(opts, "updateRaceAddress", raceList)
}

// UpdateRaceAddress is a paid mutator transaction binding the contract method 0xd7c49938.
//
// Solidity: function updateRaceAddress(address raceList) returns()
func (_IBething *IBethingSession) UpdateRaceAddress(raceList common.Address) (*types.Transaction, error) {
	return _IBething.Contract.UpdateRaceAddress(&_IBething.TransactOpts, raceList)
}

// UpdateRaceAddress is a paid mutator transaction binding the contract method 0xd7c49938.
//
// Solidity: function updateRaceAddress(address raceList) returns()
func (_IBething *IBethingTransactorSession) UpdateRaceAddress(raceList common.Address) (*types.Transaction, error) {
	return _IBething.Contract.UpdateRaceAddress(&_IBething.TransactOpts, raceList)
}

// UserSlotBet is a paid mutator transaction binding the contract method 0x2e279801.
//
// Solidity: function userSlotBet(bytes32 raceId, uint256 slotId, address user) returns(uint256)
func (_IBething *IBethingTransactor) UserSlotBet(opts *bind.TransactOpts, raceId [32]byte, slotId *big.Int, user common.Address) (*types.Transaction, error) {
	return _IBething.contract.Transact(opts, "userSlotBet", raceId, slotId, user)
}

// UserSlotBet is a paid mutator transaction binding the contract method 0x2e279801.
//
// Solidity: function userSlotBet(bytes32 raceId, uint256 slotId, address user) returns(uint256)
func (_IBething *IBethingSession) UserSlotBet(raceId [32]byte, slotId *big.Int, user common.Address) (*types.Transaction, error) {
	return _IBething.Contract.UserSlotBet(&_IBething.TransactOpts, raceId, slotId, user)
}

// UserSlotBet is a paid mutator transaction binding the contract method 0x2e279801.
//
// Solidity: function userSlotBet(bytes32 raceId, uint256 slotId, address user) returns(uint256)
func (_IBething *IBethingTransactorSession) UserSlotBet(raceId [32]byte, slotId *big.Int, user common.Address) (*types.Transaction, error) {
	return _IBething.Contract.UserSlotBet(&_IBething.TransactOpts, raceId, slotId, user)
}

// IBethingAcceptTokenUpdatedIterator is returned from FilterAcceptTokenUpdated and is used to iterate over the raw logs and unpacked data for AcceptTokenUpdated events raised by the IBething contract.
type IBethingAcceptTokenUpdatedIterator struct {
	Event *IBethingAcceptTokenUpdated // Event containing the contract specifics and raw log

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
func (it *IBethingAcceptTokenUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IBethingAcceptTokenUpdated)
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
		it.Event = new(IBethingAcceptTokenUpdated)
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
func (it *IBethingAcceptTokenUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IBethingAcceptTokenUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IBethingAcceptTokenUpdated represents a AcceptTokenUpdated event raised by the IBething contract.
type IBethingAcceptTokenUpdated struct {
	AcceptToken common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterAcceptTokenUpdated is a free log retrieval operation binding the contract event 0x44c3b193081dd50818bc919dc1963ae03762174c9a89af5fbbce54ee71a13f97.
//
// Solidity: event AcceptTokenUpdated(address acceptToken)
func (_IBething *IBethingFilterer) FilterAcceptTokenUpdated(opts *bind.FilterOpts) (*IBethingAcceptTokenUpdatedIterator, error) {

	logs, sub, err := _IBething.contract.FilterLogs(opts, "AcceptTokenUpdated")
	if err != nil {
		return nil, err
	}
	return &IBethingAcceptTokenUpdatedIterator{contract: _IBething.contract, event: "AcceptTokenUpdated", logs: logs, sub: sub}, nil
}

// WatchAcceptTokenUpdated is a free log subscription operation binding the contract event 0x44c3b193081dd50818bc919dc1963ae03762174c9a89af5fbbce54ee71a13f97.
//
// Solidity: event AcceptTokenUpdated(address acceptToken)
func (_IBething *IBethingFilterer) WatchAcceptTokenUpdated(opts *bind.WatchOpts, sink chan<- *IBethingAcceptTokenUpdated) (event.Subscription, error) {

	logs, sub, err := _IBething.contract.WatchLogs(opts, "AcceptTokenUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IBethingAcceptTokenUpdated)
				if err := _IBething.contract.UnpackLog(event, "AcceptTokenUpdated", log); err != nil {
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
func (_IBething *IBethingFilterer) ParseAcceptTokenUpdated(log types.Log) (*IBethingAcceptTokenUpdated, error) {
	event := new(IBethingAcceptTokenUpdated)
	if err := _IBething.contract.UnpackLog(event, "AcceptTokenUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IBethingBetSuccessfulIterator is returned from FilterBetSuccessful and is used to iterate over the raw logs and unpacked data for BetSuccessful events raised by the IBething contract.
type IBethingBetSuccessfulIterator struct {
	Event *IBethingBetSuccessful // Event containing the contract specifics and raw log

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
func (it *IBethingBetSuccessfulIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IBethingBetSuccessful)
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
		it.Event = new(IBethingBetSuccessful)
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
func (it *IBethingBetSuccessfulIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IBethingBetSuccessfulIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IBethingBetSuccessful represents a BetSuccessful event raised by the IBething contract.
type IBethingBetSuccessful struct {
	SlotId   *big.Int
	RaceId   [32]byte
	BetValue *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterBetSuccessful is a free log retrieval operation binding the contract event 0x90340911801212a3d95efe0dca9cdbd74549020af8cf33fc0861b74d82fe6dbd.
//
// Solidity: event BetSuccessful(uint256 slotId, bytes32 raceId, uint256 betValue)
func (_IBething *IBethingFilterer) FilterBetSuccessful(opts *bind.FilterOpts) (*IBethingBetSuccessfulIterator, error) {

	logs, sub, err := _IBething.contract.FilterLogs(opts, "BetSuccessful")
	if err != nil {
		return nil, err
	}
	return &IBethingBetSuccessfulIterator{contract: _IBething.contract, event: "BetSuccessful", logs: logs, sub: sub}, nil
}

// WatchBetSuccessful is a free log subscription operation binding the contract event 0x90340911801212a3d95efe0dca9cdbd74549020af8cf33fc0861b74d82fe6dbd.
//
// Solidity: event BetSuccessful(uint256 slotId, bytes32 raceId, uint256 betValue)
func (_IBething *IBethingFilterer) WatchBetSuccessful(opts *bind.WatchOpts, sink chan<- *IBethingBetSuccessful) (event.Subscription, error) {

	logs, sub, err := _IBething.contract.WatchLogs(opts, "BetSuccessful")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IBethingBetSuccessful)
				if err := _IBething.contract.UnpackLog(event, "BetSuccessful", log); err != nil {
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
func (_IBething *IBethingFilterer) ParseBetSuccessful(log types.Log) (*IBethingBetSuccessful, error) {
	event := new(IBethingBetSuccessful)
	if err := _IBething.contract.UnpackLog(event, "BetSuccessful", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IBethingClaimCommissionIterator is returned from FilterClaimCommission and is used to iterate over the raw logs and unpacked data for ClaimCommission events raised by the IBething contract.
type IBethingClaimCommissionIterator struct {
	Event *IBethingClaimCommission // Event containing the contract specifics and raw log

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
func (it *IBethingClaimCommissionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IBethingClaimCommission)
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
		it.Event = new(IBethingClaimCommission)
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
func (it *IBethingClaimCommissionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IBethingClaimCommissionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IBethingClaimCommission represents a ClaimCommission event raised by the IBething contract.
type IBethingClaimCommission struct {
	RaceId     [32]byte
	ClaimValue *big.Int
	Receiver   common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterClaimCommission is a free log retrieval operation binding the contract event 0x94fe36569bb216f6cb564fee8983d035c7cacba5679023a4ccabe33fdf6d4d88.
//
// Solidity: event ClaimCommission(bytes32 raceId, uint256 claimValue, address receiver)
func (_IBething *IBethingFilterer) FilterClaimCommission(opts *bind.FilterOpts) (*IBethingClaimCommissionIterator, error) {

	logs, sub, err := _IBething.contract.FilterLogs(opts, "ClaimCommission")
	if err != nil {
		return nil, err
	}
	return &IBethingClaimCommissionIterator{contract: _IBething.contract, event: "ClaimCommission", logs: logs, sub: sub}, nil
}

// WatchClaimCommission is a free log subscription operation binding the contract event 0x94fe36569bb216f6cb564fee8983d035c7cacba5679023a4ccabe33fdf6d4d88.
//
// Solidity: event ClaimCommission(bytes32 raceId, uint256 claimValue, address receiver)
func (_IBething *IBethingFilterer) WatchClaimCommission(opts *bind.WatchOpts, sink chan<- *IBethingClaimCommission) (event.Subscription, error) {

	logs, sub, err := _IBething.contract.WatchLogs(opts, "ClaimCommission")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IBethingClaimCommission)
				if err := _IBething.contract.UnpackLog(event, "ClaimCommission", log); err != nil {
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
func (_IBething *IBethingFilterer) ParseClaimCommission(log types.Log) (*IBethingClaimCommission, error) {
	event := new(IBethingClaimCommission)
	if err := _IBething.contract.UnpackLog(event, "ClaimCommission", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IBethingClaimSuccessfulIterator is returned from FilterClaimSuccessful and is used to iterate over the raw logs and unpacked data for ClaimSuccessful events raised by the IBething contract.
type IBethingClaimSuccessfulIterator struct {
	Event *IBethingClaimSuccessful // Event containing the contract specifics and raw log

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
func (it *IBethingClaimSuccessfulIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IBethingClaimSuccessful)
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
		it.Event = new(IBethingClaimSuccessful)
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
func (it *IBethingClaimSuccessfulIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IBethingClaimSuccessfulIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IBethingClaimSuccessful represents a ClaimSuccessful event raised by the IBething contract.
type IBethingClaimSuccessful struct {
	RaceId     [32]byte
	ClaimValue *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterClaimSuccessful is a free log retrieval operation binding the contract event 0x78d224c5b360f08891a0d6c82ec1fceea8a4e5b39dc51c1373ddfaa13d848df0.
//
// Solidity: event ClaimSuccessful(bytes32 raceId, uint256 claimValue)
func (_IBething *IBethingFilterer) FilterClaimSuccessful(opts *bind.FilterOpts) (*IBethingClaimSuccessfulIterator, error) {

	logs, sub, err := _IBething.contract.FilterLogs(opts, "ClaimSuccessful")
	if err != nil {
		return nil, err
	}
	return &IBethingClaimSuccessfulIterator{contract: _IBething.contract, event: "ClaimSuccessful", logs: logs, sub: sub}, nil
}

// WatchClaimSuccessful is a free log subscription operation binding the contract event 0x78d224c5b360f08891a0d6c82ec1fceea8a4e5b39dc51c1373ddfaa13d848df0.
//
// Solidity: event ClaimSuccessful(bytes32 raceId, uint256 claimValue)
func (_IBething *IBethingFilterer) WatchClaimSuccessful(opts *bind.WatchOpts, sink chan<- *IBethingClaimSuccessful) (event.Subscription, error) {

	logs, sub, err := _IBething.contract.WatchLogs(opts, "ClaimSuccessful")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IBethingClaimSuccessful)
				if err := _IBething.contract.UnpackLog(event, "ClaimSuccessful", log); err != nil {
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
func (_IBething *IBethingFilterer) ParseClaimSuccessful(log types.Log) (*IBethingClaimSuccessful, error) {
	event := new(IBethingClaimSuccessful)
	if err := _IBething.contract.UnpackLog(event, "ClaimSuccessful", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IBethingFundSuccessfulIterator is returned from FilterFundSuccessful and is used to iterate over the raw logs and unpacked data for FundSuccessful events raised by the IBething contract.
type IBethingFundSuccessfulIterator struct {
	Event *IBethingFundSuccessful // Event containing the contract specifics and raw log

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
func (it *IBethingFundSuccessfulIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IBethingFundSuccessful)
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
		it.Event = new(IBethingFundSuccessful)
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
func (it *IBethingFundSuccessfulIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IBethingFundSuccessfulIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IBethingFundSuccessful represents a FundSuccessful event raised by the IBething contract.
type IBethingFundSuccessful struct {
	RaceId    [32]byte
	FundValue *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterFundSuccessful is a free log retrieval operation binding the contract event 0x18828fa7a168b0db08e3e72ef21348f4c4fa1ee1b46803878b1195794aaad4c3.
//
// Solidity: event FundSuccessful(bytes32 raceId, uint256 fundValue)
func (_IBething *IBethingFilterer) FilterFundSuccessful(opts *bind.FilterOpts) (*IBethingFundSuccessfulIterator, error) {

	logs, sub, err := _IBething.contract.FilterLogs(opts, "FundSuccessful")
	if err != nil {
		return nil, err
	}
	return &IBethingFundSuccessfulIterator{contract: _IBething.contract, event: "FundSuccessful", logs: logs, sub: sub}, nil
}

// WatchFundSuccessful is a free log subscription operation binding the contract event 0x18828fa7a168b0db08e3e72ef21348f4c4fa1ee1b46803878b1195794aaad4c3.
//
// Solidity: event FundSuccessful(bytes32 raceId, uint256 fundValue)
func (_IBething *IBethingFilterer) WatchFundSuccessful(opts *bind.WatchOpts, sink chan<- *IBethingFundSuccessful) (event.Subscription, error) {

	logs, sub, err := _IBething.contract.WatchLogs(opts, "FundSuccessful")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IBethingFundSuccessful)
				if err := _IBething.contract.UnpackLog(event, "FundSuccessful", log); err != nil {
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
func (_IBething *IBethingFilterer) ParseFundSuccessful(log types.Log) (*IBethingFundSuccessful, error) {
	event := new(IBethingFundSuccessful)
	if err := _IBething.contract.UnpackLog(event, "FundSuccessful", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IBethingRaceListUpdatedIterator is returned from FilterRaceListUpdated and is used to iterate over the raw logs and unpacked data for RaceListUpdated events raised by the IBething contract.
type IBethingRaceListUpdatedIterator struct {
	Event *IBethingRaceListUpdated // Event containing the contract specifics and raw log

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
func (it *IBethingRaceListUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IBethingRaceListUpdated)
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
		it.Event = new(IBethingRaceListUpdated)
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
func (it *IBethingRaceListUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IBethingRaceListUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IBethingRaceListUpdated represents a RaceListUpdated event raised by the IBething contract.
type IBethingRaceListUpdated struct {
	Race common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterRaceListUpdated is a free log retrieval operation binding the contract event 0x7a659342c2b03e1c14729e3f1c86f414756d13bd8ff0636713cc22246182b42d.
//
// Solidity: event RaceListUpdated(address race)
func (_IBething *IBethingFilterer) FilterRaceListUpdated(opts *bind.FilterOpts) (*IBethingRaceListUpdatedIterator, error) {

	logs, sub, err := _IBething.contract.FilterLogs(opts, "RaceListUpdated")
	if err != nil {
		return nil, err
	}
	return &IBethingRaceListUpdatedIterator{contract: _IBething.contract, event: "RaceListUpdated", logs: logs, sub: sub}, nil
}

// WatchRaceListUpdated is a free log subscription operation binding the contract event 0x7a659342c2b03e1c14729e3f1c86f414756d13bd8ff0636713cc22246182b42d.
//
// Solidity: event RaceListUpdated(address race)
func (_IBething *IBethingFilterer) WatchRaceListUpdated(opts *bind.WatchOpts, sink chan<- *IBethingRaceListUpdated) (event.Subscription, error) {

	logs, sub, err := _IBething.contract.WatchLogs(opts, "RaceListUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IBethingRaceListUpdated)
				if err := _IBething.contract.UnpackLog(event, "RaceListUpdated", log); err != nil {
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
func (_IBething *IBethingFilterer) ParseRaceListUpdated(log types.Log) (*IBethingRaceListUpdated, error) {
	event := new(IBethingRaceListUpdated)
	if err := _IBething.contract.UnpackLog(event, "RaceListUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
