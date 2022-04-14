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

// IRaceRace is an auto generated low-level Go binding around an user-defined struct.
type IRaceRace struct {
	Commission *big.Int
	RewardRate *big.Int
	Slots      *big.Int
	BetStarted *big.Int
	BetEnded   *big.Int
	Result     [32]byte
}

// SmcMetaData contains all meta data concerning the Smc contract.
var SmcMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"InvalidCommission\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidRewardRate\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidSlot\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RaceNotExisted\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"}],\"name\":\"TooEarly\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"}],\"name\":\"TooLate\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"}],\"name\":\"RaceCancelled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"commission\",\"type\":\"uint256\"}],\"name\":\"RaceCommissionUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"slots\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"betStarted\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"betEnded\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"commission\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"rewardRate\",\"type\":\"uint256\"}],\"name\":\"RaceCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"result\",\"type\":\"bytes32\"}],\"name\":\"RaceResultUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"rewardRate\",\"type\":\"uint256\"}],\"name\":\"RaceRewardRateUpdated\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"slots\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"betStarted\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"betEnded\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"commission\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rewardRate\",\"type\":\"uint256\"}],\"name\":\"createRace\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"}],\"name\":\"getRace\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"commission\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rewardRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"slots\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"betStarted\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"betEnded\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"result\",\"type\":\"bytes32\"}],\"internalType\":\"structIRace.Race\",\"name\":\"race\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"}],\"name\":\"cancelRace\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"result\",\"type\":\"bytes32\"}],\"name\":\"updateResult\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"commission\",\"type\":\"uint256\"}],\"name\":\"updateCommission\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"rewardRate\",\"type\":\"uint256\"}],\"name\":\"updateRewardRate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"raceId\",\"type\":\"bytes32\"}],\"name\":\"raceIsExisted\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"raceId\",\"type\":\"bytes32\"}],\"name\":\"raceResult\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
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

// GetRace is a free data retrieval call binding the contract method 0xf70125db.
//
// Solidity: function getRace(bytes32 id) view returns((uint256,uint256,uint256,uint256,uint256,bytes32) race)
func (_Smc *SmcCaller) GetRace(opts *bind.CallOpts, id [32]byte) (IRaceRace, error) {
	var out []interface{}
	err := _Smc.contract.Call(opts, &out, "getRace", id)

	if err != nil {
		return *new(IRaceRace), err
	}

	out0 := *abi.ConvertType(out[0], new(IRaceRace)).(*IRaceRace)

	return out0, err

}

// GetRace is a free data retrieval call binding the contract method 0xf70125db.
//
// Solidity: function getRace(bytes32 id) view returns((uint256,uint256,uint256,uint256,uint256,bytes32) race)
func (_Smc *SmcSession) GetRace(id [32]byte) (IRaceRace, error) {
	return _Smc.Contract.GetRace(&_Smc.CallOpts, id)
}

// GetRace is a free data retrieval call binding the contract method 0xf70125db.
//
// Solidity: function getRace(bytes32 id) view returns((uint256,uint256,uint256,uint256,uint256,bytes32) race)
func (_Smc *SmcCallerSession) GetRace(id [32]byte) (IRaceRace, error) {
	return _Smc.Contract.GetRace(&_Smc.CallOpts, id)
}

// RaceIsExisted is a free data retrieval call binding the contract method 0x2869ad12.
//
// Solidity: function raceIsExisted(bytes32 raceId) view returns(bool)
func (_Smc *SmcCaller) RaceIsExisted(opts *bind.CallOpts, raceId [32]byte) (bool, error) {
	var out []interface{}
	err := _Smc.contract.Call(opts, &out, "raceIsExisted", raceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// RaceIsExisted is a free data retrieval call binding the contract method 0x2869ad12.
//
// Solidity: function raceIsExisted(bytes32 raceId) view returns(bool)
func (_Smc *SmcSession) RaceIsExisted(raceId [32]byte) (bool, error) {
	return _Smc.Contract.RaceIsExisted(&_Smc.CallOpts, raceId)
}

// RaceIsExisted is a free data retrieval call binding the contract method 0x2869ad12.
//
// Solidity: function raceIsExisted(bytes32 raceId) view returns(bool)
func (_Smc *SmcCallerSession) RaceIsExisted(raceId [32]byte) (bool, error) {
	return _Smc.Contract.RaceIsExisted(&_Smc.CallOpts, raceId)
}

// RaceResult is a free data retrieval call binding the contract method 0x6d93b958.
//
// Solidity: function raceResult(bytes32 raceId) view returns(bytes32)
func (_Smc *SmcCaller) RaceResult(opts *bind.CallOpts, raceId [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _Smc.contract.Call(opts, &out, "raceResult", raceId)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// RaceResult is a free data retrieval call binding the contract method 0x6d93b958.
//
// Solidity: function raceResult(bytes32 raceId) view returns(bytes32)
func (_Smc *SmcSession) RaceResult(raceId [32]byte) ([32]byte, error) {
	return _Smc.Contract.RaceResult(&_Smc.CallOpts, raceId)
}

// RaceResult is a free data retrieval call binding the contract method 0x6d93b958.
//
// Solidity: function raceResult(bytes32 raceId) view returns(bytes32)
func (_Smc *SmcCallerSession) RaceResult(raceId [32]byte) ([32]byte, error) {
	return _Smc.Contract.RaceResult(&_Smc.CallOpts, raceId)
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

// CancelRace is a paid mutator transaction binding the contract method 0xc05c107d.
//
// Solidity: function cancelRace(bytes32 id) returns()
func (_Smc *SmcTransactor) CancelRace(opts *bind.TransactOpts, id [32]byte) (*types.Transaction, error) {
	return _Smc.contract.Transact(opts, "cancelRace", id)
}

// CancelRace is a paid mutator transaction binding the contract method 0xc05c107d.
//
// Solidity: function cancelRace(bytes32 id) returns()
func (_Smc *SmcSession) CancelRace(id [32]byte) (*types.Transaction, error) {
	return _Smc.Contract.CancelRace(&_Smc.TransactOpts, id)
}

// CancelRace is a paid mutator transaction binding the contract method 0xc05c107d.
//
// Solidity: function cancelRace(bytes32 id) returns()
func (_Smc *SmcTransactorSession) CancelRace(id [32]byte) (*types.Transaction, error) {
	return _Smc.Contract.CancelRace(&_Smc.TransactOpts, id)
}

// CreateRace is a paid mutator transaction binding the contract method 0xd65b869a.
//
// Solidity: function createRace(uint256 slots, uint256 betStarted, uint256 betEnded, uint256 commission, uint256 rewardRate) returns()
func (_Smc *SmcTransactor) CreateRace(opts *bind.TransactOpts, slots *big.Int, betStarted *big.Int, betEnded *big.Int, commission *big.Int, rewardRate *big.Int) (*types.Transaction, error) {
	return _Smc.contract.Transact(opts, "createRace", slots, betStarted, betEnded, commission, rewardRate)
}

// CreateRace is a paid mutator transaction binding the contract method 0xd65b869a.
//
// Solidity: function createRace(uint256 slots, uint256 betStarted, uint256 betEnded, uint256 commission, uint256 rewardRate) returns()
func (_Smc *SmcSession) CreateRace(slots *big.Int, betStarted *big.Int, betEnded *big.Int, commission *big.Int, rewardRate *big.Int) (*types.Transaction, error) {
	return _Smc.Contract.CreateRace(&_Smc.TransactOpts, slots, betStarted, betEnded, commission, rewardRate)
}

// CreateRace is a paid mutator transaction binding the contract method 0xd65b869a.
//
// Solidity: function createRace(uint256 slots, uint256 betStarted, uint256 betEnded, uint256 commission, uint256 rewardRate) returns()
func (_Smc *SmcTransactorSession) CreateRace(slots *big.Int, betStarted *big.Int, betEnded *big.Int, commission *big.Int, rewardRate *big.Int) (*types.Transaction, error) {
	return _Smc.Contract.CreateRace(&_Smc.TransactOpts, slots, betStarted, betEnded, commission, rewardRate)
}

// UpdateCommission is a paid mutator transaction binding the contract method 0xc58f80a7.
//
// Solidity: function updateCommission(bytes32 id, uint256 commission) returns()
func (_Smc *SmcTransactor) UpdateCommission(opts *bind.TransactOpts, id [32]byte, commission *big.Int) (*types.Transaction, error) {
	return _Smc.contract.Transact(opts, "updateCommission", id, commission)
}

// UpdateCommission is a paid mutator transaction binding the contract method 0xc58f80a7.
//
// Solidity: function updateCommission(bytes32 id, uint256 commission) returns()
func (_Smc *SmcSession) UpdateCommission(id [32]byte, commission *big.Int) (*types.Transaction, error) {
	return _Smc.Contract.UpdateCommission(&_Smc.TransactOpts, id, commission)
}

// UpdateCommission is a paid mutator transaction binding the contract method 0xc58f80a7.
//
// Solidity: function updateCommission(bytes32 id, uint256 commission) returns()
func (_Smc *SmcTransactorSession) UpdateCommission(id [32]byte, commission *big.Int) (*types.Transaction, error) {
	return _Smc.Contract.UpdateCommission(&_Smc.TransactOpts, id, commission)
}

// UpdateResult is a paid mutator transaction binding the contract method 0xb6d940fc.
//
// Solidity: function updateResult(bytes32 id, bytes32 result) returns()
func (_Smc *SmcTransactor) UpdateResult(opts *bind.TransactOpts, id [32]byte, result [32]byte) (*types.Transaction, error) {
	return _Smc.contract.Transact(opts, "updateResult", id, result)
}

// UpdateResult is a paid mutator transaction binding the contract method 0xb6d940fc.
//
// Solidity: function updateResult(bytes32 id, bytes32 result) returns()
func (_Smc *SmcSession) UpdateResult(id [32]byte, result [32]byte) (*types.Transaction, error) {
	return _Smc.Contract.UpdateResult(&_Smc.TransactOpts, id, result)
}

// UpdateResult is a paid mutator transaction binding the contract method 0xb6d940fc.
//
// Solidity: function updateResult(bytes32 id, bytes32 result) returns()
func (_Smc *SmcTransactorSession) UpdateResult(id [32]byte, result [32]byte) (*types.Transaction, error) {
	return _Smc.Contract.UpdateResult(&_Smc.TransactOpts, id, result)
}

// UpdateRewardRate is a paid mutator transaction binding the contract method 0x68b2c40c.
//
// Solidity: function updateRewardRate(bytes32 id, uint256 rewardRate) returns()
func (_Smc *SmcTransactor) UpdateRewardRate(opts *bind.TransactOpts, id [32]byte, rewardRate *big.Int) (*types.Transaction, error) {
	return _Smc.contract.Transact(opts, "updateRewardRate", id, rewardRate)
}

// UpdateRewardRate is a paid mutator transaction binding the contract method 0x68b2c40c.
//
// Solidity: function updateRewardRate(bytes32 id, uint256 rewardRate) returns()
func (_Smc *SmcSession) UpdateRewardRate(id [32]byte, rewardRate *big.Int) (*types.Transaction, error) {
	return _Smc.Contract.UpdateRewardRate(&_Smc.TransactOpts, id, rewardRate)
}

// UpdateRewardRate is a paid mutator transaction binding the contract method 0x68b2c40c.
//
// Solidity: function updateRewardRate(bytes32 id, uint256 rewardRate) returns()
func (_Smc *SmcTransactorSession) UpdateRewardRate(id [32]byte, rewardRate *big.Int) (*types.Transaction, error) {
	return _Smc.Contract.UpdateRewardRate(&_Smc.TransactOpts, id, rewardRate)
}

// SmcRaceCancelledIterator is returned from FilterRaceCancelled and is used to iterate over the raw logs and unpacked data for RaceCancelled events raised by the Smc contract.
type SmcRaceCancelledIterator struct {
	Event *SmcRaceCancelled // Event containing the contract specifics and raw log

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
func (it *SmcRaceCancelledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SmcRaceCancelled)
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
		it.Event = new(SmcRaceCancelled)
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
func (it *SmcRaceCancelledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SmcRaceCancelledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SmcRaceCancelled represents a RaceCancelled event raised by the Smc contract.
type SmcRaceCancelled struct {
	Id  [32]byte
	Raw types.Log // Blockchain specific contextual infos
}

// FilterRaceCancelled is a free log retrieval operation binding the contract event 0x3f3669b11af9b00fb2a95f8ded36a79a2e1d5374efbc5ec6b2b5b96ceaf9ac6a.
//
// Solidity: event RaceCancelled(bytes32 id)
func (_Smc *SmcFilterer) FilterRaceCancelled(opts *bind.FilterOpts) (*SmcRaceCancelledIterator, error) {

	logs, sub, err := _Smc.contract.FilterLogs(opts, "RaceCancelled")
	if err != nil {
		return nil, err
	}
	return &SmcRaceCancelledIterator{contract: _Smc.contract, event: "RaceCancelled", logs: logs, sub: sub}, nil
}

// WatchRaceCancelled is a free log subscription operation binding the contract event 0x3f3669b11af9b00fb2a95f8ded36a79a2e1d5374efbc5ec6b2b5b96ceaf9ac6a.
//
// Solidity: event RaceCancelled(bytes32 id)
func (_Smc *SmcFilterer) WatchRaceCancelled(opts *bind.WatchOpts, sink chan<- *SmcRaceCancelled) (event.Subscription, error) {

	logs, sub, err := _Smc.contract.WatchLogs(opts, "RaceCancelled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SmcRaceCancelled)
				if err := _Smc.contract.UnpackLog(event, "RaceCancelled", log); err != nil {
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
func (_Smc *SmcFilterer) ParseRaceCancelled(log types.Log) (*SmcRaceCancelled, error) {
	event := new(SmcRaceCancelled)
	if err := _Smc.contract.UnpackLog(event, "RaceCancelled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SmcRaceCommissionUpdatedIterator is returned from FilterRaceCommissionUpdated and is used to iterate over the raw logs and unpacked data for RaceCommissionUpdated events raised by the Smc contract.
type SmcRaceCommissionUpdatedIterator struct {
	Event *SmcRaceCommissionUpdated // Event containing the contract specifics and raw log

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
func (it *SmcRaceCommissionUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SmcRaceCommissionUpdated)
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
		it.Event = new(SmcRaceCommissionUpdated)
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
func (it *SmcRaceCommissionUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SmcRaceCommissionUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SmcRaceCommissionUpdated represents a RaceCommissionUpdated event raised by the Smc contract.
type SmcRaceCommissionUpdated struct {
	Id         [32]byte
	Commission *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterRaceCommissionUpdated is a free log retrieval operation binding the contract event 0x9ff8c0f5ebec7767df2725581a4945a5a800a1f929ec51da017dc56d31a2b3a0.
//
// Solidity: event RaceCommissionUpdated(bytes32 id, uint256 commission)
func (_Smc *SmcFilterer) FilterRaceCommissionUpdated(opts *bind.FilterOpts) (*SmcRaceCommissionUpdatedIterator, error) {

	logs, sub, err := _Smc.contract.FilterLogs(opts, "RaceCommissionUpdated")
	if err != nil {
		return nil, err
	}
	return &SmcRaceCommissionUpdatedIterator{contract: _Smc.contract, event: "RaceCommissionUpdated", logs: logs, sub: sub}, nil
}

// WatchRaceCommissionUpdated is a free log subscription operation binding the contract event 0x9ff8c0f5ebec7767df2725581a4945a5a800a1f929ec51da017dc56d31a2b3a0.
//
// Solidity: event RaceCommissionUpdated(bytes32 id, uint256 commission)
func (_Smc *SmcFilterer) WatchRaceCommissionUpdated(opts *bind.WatchOpts, sink chan<- *SmcRaceCommissionUpdated) (event.Subscription, error) {

	logs, sub, err := _Smc.contract.WatchLogs(opts, "RaceCommissionUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SmcRaceCommissionUpdated)
				if err := _Smc.contract.UnpackLog(event, "RaceCommissionUpdated", log); err != nil {
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

// ParseRaceCommissionUpdated is a log parse operation binding the contract event 0x9ff8c0f5ebec7767df2725581a4945a5a800a1f929ec51da017dc56d31a2b3a0.
//
// Solidity: event RaceCommissionUpdated(bytes32 id, uint256 commission)
func (_Smc *SmcFilterer) ParseRaceCommissionUpdated(log types.Log) (*SmcRaceCommissionUpdated, error) {
	event := new(SmcRaceCommissionUpdated)
	if err := _Smc.contract.UnpackLog(event, "RaceCommissionUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SmcRaceCreatedIterator is returned from FilterRaceCreated and is used to iterate over the raw logs and unpacked data for RaceCreated events raised by the Smc contract.
type SmcRaceCreatedIterator struct {
	Event *SmcRaceCreated // Event containing the contract specifics and raw log

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
func (it *SmcRaceCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SmcRaceCreated)
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
		it.Event = new(SmcRaceCreated)
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
func (it *SmcRaceCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SmcRaceCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SmcRaceCreated represents a RaceCreated event raised by the Smc contract.
type SmcRaceCreated struct {
	Id         [32]byte
	Slots      *big.Int
	BetStarted *big.Int
	BetEnded   *big.Int
	Commission *big.Int
	RewardRate *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterRaceCreated is a free log retrieval operation binding the contract event 0x1da8002ad97980e4380172489894c1e0e2d36896f491ff1d1773613fac4c8a22.
//
// Solidity: event RaceCreated(bytes32 id, uint256 slots, uint256 betStarted, uint256 betEnded, uint256 commission, uint256 rewardRate)
func (_Smc *SmcFilterer) FilterRaceCreated(opts *bind.FilterOpts) (*SmcRaceCreatedIterator, error) {

	logs, sub, err := _Smc.contract.FilterLogs(opts, "RaceCreated")
	if err != nil {
		return nil, err
	}
	return &SmcRaceCreatedIterator{contract: _Smc.contract, event: "RaceCreated", logs: logs, sub: sub}, nil
}

// WatchRaceCreated is a free log subscription operation binding the contract event 0x1da8002ad97980e4380172489894c1e0e2d36896f491ff1d1773613fac4c8a22.
//
// Solidity: event RaceCreated(bytes32 id, uint256 slots, uint256 betStarted, uint256 betEnded, uint256 commission, uint256 rewardRate)
func (_Smc *SmcFilterer) WatchRaceCreated(opts *bind.WatchOpts, sink chan<- *SmcRaceCreated) (event.Subscription, error) {

	logs, sub, err := _Smc.contract.WatchLogs(opts, "RaceCreated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SmcRaceCreated)
				if err := _Smc.contract.UnpackLog(event, "RaceCreated", log); err != nil {
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

// ParseRaceCreated is a log parse operation binding the contract event 0x1da8002ad97980e4380172489894c1e0e2d36896f491ff1d1773613fac4c8a22.
//
// Solidity: event RaceCreated(bytes32 id, uint256 slots, uint256 betStarted, uint256 betEnded, uint256 commission, uint256 rewardRate)
func (_Smc *SmcFilterer) ParseRaceCreated(log types.Log) (*SmcRaceCreated, error) {
	event := new(SmcRaceCreated)
	if err := _Smc.contract.UnpackLog(event, "RaceCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SmcRaceResultUpdatedIterator is returned from FilterRaceResultUpdated and is used to iterate over the raw logs and unpacked data for RaceResultUpdated events raised by the Smc contract.
type SmcRaceResultUpdatedIterator struct {
	Event *SmcRaceResultUpdated // Event containing the contract specifics and raw log

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
func (it *SmcRaceResultUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SmcRaceResultUpdated)
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
		it.Event = new(SmcRaceResultUpdated)
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
func (it *SmcRaceResultUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SmcRaceResultUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SmcRaceResultUpdated represents a RaceResultUpdated event raised by the Smc contract.
type SmcRaceResultUpdated struct {
	Id     [32]byte
	Result [32]byte
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRaceResultUpdated is a free log retrieval operation binding the contract event 0xd82663592968d73ce1995154e44c793fcb46a4006abfb7438656fb0d7ba5ff49.
//
// Solidity: event RaceResultUpdated(bytes32 id, bytes32 result)
func (_Smc *SmcFilterer) FilterRaceResultUpdated(opts *bind.FilterOpts) (*SmcRaceResultUpdatedIterator, error) {

	logs, sub, err := _Smc.contract.FilterLogs(opts, "RaceResultUpdated")
	if err != nil {
		return nil, err
	}
	return &SmcRaceResultUpdatedIterator{contract: _Smc.contract, event: "RaceResultUpdated", logs: logs, sub: sub}, nil
}

// WatchRaceResultUpdated is a free log subscription operation binding the contract event 0xd82663592968d73ce1995154e44c793fcb46a4006abfb7438656fb0d7ba5ff49.
//
// Solidity: event RaceResultUpdated(bytes32 id, bytes32 result)
func (_Smc *SmcFilterer) WatchRaceResultUpdated(opts *bind.WatchOpts, sink chan<- *SmcRaceResultUpdated) (event.Subscription, error) {

	logs, sub, err := _Smc.contract.WatchLogs(opts, "RaceResultUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SmcRaceResultUpdated)
				if err := _Smc.contract.UnpackLog(event, "RaceResultUpdated", log); err != nil {
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
func (_Smc *SmcFilterer) ParseRaceResultUpdated(log types.Log) (*SmcRaceResultUpdated, error) {
	event := new(SmcRaceResultUpdated)
	if err := _Smc.contract.UnpackLog(event, "RaceResultUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SmcRaceRewardRateUpdatedIterator is returned from FilterRaceRewardRateUpdated and is used to iterate over the raw logs and unpacked data for RaceRewardRateUpdated events raised by the Smc contract.
type SmcRaceRewardRateUpdatedIterator struct {
	Event *SmcRaceRewardRateUpdated // Event containing the contract specifics and raw log

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
func (it *SmcRaceRewardRateUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SmcRaceRewardRateUpdated)
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
		it.Event = new(SmcRaceRewardRateUpdated)
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
func (it *SmcRaceRewardRateUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SmcRaceRewardRateUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SmcRaceRewardRateUpdated represents a RaceRewardRateUpdated event raised by the Smc contract.
type SmcRaceRewardRateUpdated struct {
	Id         [32]byte
	RewardRate *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterRaceRewardRateUpdated is a free log retrieval operation binding the contract event 0x25a0545ede30f4217593809347e09822194f6a9fe74a5b3f2c5b6e843375f422.
//
// Solidity: event RaceRewardRateUpdated(bytes32 id, uint256 rewardRate)
func (_Smc *SmcFilterer) FilterRaceRewardRateUpdated(opts *bind.FilterOpts) (*SmcRaceRewardRateUpdatedIterator, error) {

	logs, sub, err := _Smc.contract.FilterLogs(opts, "RaceRewardRateUpdated")
	if err != nil {
		return nil, err
	}
	return &SmcRaceRewardRateUpdatedIterator{contract: _Smc.contract, event: "RaceRewardRateUpdated", logs: logs, sub: sub}, nil
}

// WatchRaceRewardRateUpdated is a free log subscription operation binding the contract event 0x25a0545ede30f4217593809347e09822194f6a9fe74a5b3f2c5b6e843375f422.
//
// Solidity: event RaceRewardRateUpdated(bytes32 id, uint256 rewardRate)
func (_Smc *SmcFilterer) WatchRaceRewardRateUpdated(opts *bind.WatchOpts, sink chan<- *SmcRaceRewardRateUpdated) (event.Subscription, error) {

	logs, sub, err := _Smc.contract.WatchLogs(opts, "RaceRewardRateUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SmcRaceRewardRateUpdated)
				if err := _Smc.contract.UnpackLog(event, "RaceRewardRateUpdated", log); err != nil {
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

// ParseRaceRewardRateUpdated is a log parse operation binding the contract event 0x25a0545ede30f4217593809347e09822194f6a9fe74a5b3f2c5b6e843375f422.
//
// Solidity: event RaceRewardRateUpdated(bytes32 id, uint256 rewardRate)
func (_Smc *SmcFilterer) ParseRaceRewardRateUpdated(log types.Log) (*SmcRaceRewardRateUpdated, error) {
	event := new(SmcRaceRewardRateUpdated)
	if err := _Smc.contract.UnpackLog(event, "RaceRewardRateUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
