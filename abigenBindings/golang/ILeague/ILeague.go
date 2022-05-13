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
	ABI: "[{\"inputs\":[],\"name\":\"CannotCancel\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidContract\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidSlot\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RaceExisted\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RaceNotExisted\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"}],\"name\":\"TooEarly\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"}],\"name\":\"TooLate\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"}],\"name\":\"RaceCancelled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"slots\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"startAt\",\"type\":\"uint256\"}],\"name\":\"RaceCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"result\",\"type\":\"bytes32\"}],\"name\":\"RaceResultUpdated\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"slots\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startAt\",\"type\":\"uint256\"}],\"name\":\"createRace\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"raceId\",\"type\":\"bytes32\"}],\"name\":\"raceSlots\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"slots\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"raceId\",\"type\":\"bytes32\"}],\"name\":\"raceStartAt\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"startAt\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"raceId\",\"type\":\"bytes32\"}],\"name\":\"raceResult\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"result\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"leagueName\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"}],\"name\":\"cancelRace\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"register\",\"type\":\"address\"}],\"name\":\"registerIsValid\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"raceId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"result\",\"type\":\"bytes32\"}],\"name\":\"updateRaceResult\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
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

// RaceResult is a free data retrieval call binding the contract method 0x6d93b958.
//
// Solidity: function raceResult(bytes32 raceId) view returns(bytes32 result)
func (_ILeague *ILeagueCaller) RaceResult(opts *bind.CallOpts, raceId [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _ILeague.contract.Call(opts, &out, "raceResult", raceId)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// RaceResult is a free data retrieval call binding the contract method 0x6d93b958.
//
// Solidity: function raceResult(bytes32 raceId) view returns(bytes32 result)
func (_ILeague *ILeagueSession) RaceResult(raceId [32]byte) ([32]byte, error) {
	return _ILeague.Contract.RaceResult(&_ILeague.CallOpts, raceId)
}

// RaceResult is a free data retrieval call binding the contract method 0x6d93b958.
//
// Solidity: function raceResult(bytes32 raceId) view returns(bytes32 result)
func (_ILeague *ILeagueCallerSession) RaceResult(raceId [32]byte) ([32]byte, error) {
	return _ILeague.Contract.RaceResult(&_ILeague.CallOpts, raceId)
}

// RaceSlots is a free data retrieval call binding the contract method 0xd0cbb2a4.
//
// Solidity: function raceSlots(bytes32 raceId) view returns(uint256 slots)
func (_ILeague *ILeagueCaller) RaceSlots(opts *bind.CallOpts, raceId [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _ILeague.contract.Call(opts, &out, "raceSlots", raceId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RaceSlots is a free data retrieval call binding the contract method 0xd0cbb2a4.
//
// Solidity: function raceSlots(bytes32 raceId) view returns(uint256 slots)
func (_ILeague *ILeagueSession) RaceSlots(raceId [32]byte) (*big.Int, error) {
	return _ILeague.Contract.RaceSlots(&_ILeague.CallOpts, raceId)
}

// RaceSlots is a free data retrieval call binding the contract method 0xd0cbb2a4.
//
// Solidity: function raceSlots(bytes32 raceId) view returns(uint256 slots)
func (_ILeague *ILeagueCallerSession) RaceSlots(raceId [32]byte) (*big.Int, error) {
	return _ILeague.Contract.RaceSlots(&_ILeague.CallOpts, raceId)
}

// RaceStartAt is a free data retrieval call binding the contract method 0x1d88e1bd.
//
// Solidity: function raceStartAt(bytes32 raceId) view returns(uint256 startAt)
func (_ILeague *ILeagueCaller) RaceStartAt(opts *bind.CallOpts, raceId [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _ILeague.contract.Call(opts, &out, "raceStartAt", raceId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RaceStartAt is a free data retrieval call binding the contract method 0x1d88e1bd.
//
// Solidity: function raceStartAt(bytes32 raceId) view returns(uint256 startAt)
func (_ILeague *ILeagueSession) RaceStartAt(raceId [32]byte) (*big.Int, error) {
	return _ILeague.Contract.RaceStartAt(&_ILeague.CallOpts, raceId)
}

// RaceStartAt is a free data retrieval call binding the contract method 0x1d88e1bd.
//
// Solidity: function raceStartAt(bytes32 raceId) view returns(uint256 startAt)
func (_ILeague *ILeagueCallerSession) RaceStartAt(raceId [32]byte) (*big.Int, error) {
	return _ILeague.Contract.RaceStartAt(&_ILeague.CallOpts, raceId)
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

// CancelRace is a paid mutator transaction binding the contract method 0xc05c107d.
//
// Solidity: function cancelRace(bytes32 id) returns()
func (_ILeague *ILeagueTransactor) CancelRace(opts *bind.TransactOpts, id [32]byte) (*types.Transaction, error) {
	return _ILeague.contract.Transact(opts, "cancelRace", id)
}

// CancelRace is a paid mutator transaction binding the contract method 0xc05c107d.
//
// Solidity: function cancelRace(bytes32 id) returns()
func (_ILeague *ILeagueSession) CancelRace(id [32]byte) (*types.Transaction, error) {
	return _ILeague.Contract.CancelRace(&_ILeague.TransactOpts, id)
}

// CancelRace is a paid mutator transaction binding the contract method 0xc05c107d.
//
// Solidity: function cancelRace(bytes32 id) returns()
func (_ILeague *ILeagueTransactorSession) CancelRace(id [32]byte) (*types.Transaction, error) {
	return _ILeague.Contract.CancelRace(&_ILeague.TransactOpts, id)
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

// RegisterIsValid is a paid mutator transaction binding the contract method 0x53eeece1.
//
// Solidity: function registerIsValid(address register) returns(bool)
func (_ILeague *ILeagueTransactor) RegisterIsValid(opts *bind.TransactOpts, register common.Address) (*types.Transaction, error) {
	return _ILeague.contract.Transact(opts, "registerIsValid", register)
}

// RegisterIsValid is a paid mutator transaction binding the contract method 0x53eeece1.
//
// Solidity: function registerIsValid(address register) returns(bool)
func (_ILeague *ILeagueSession) RegisterIsValid(register common.Address) (*types.Transaction, error) {
	return _ILeague.Contract.RegisterIsValid(&_ILeague.TransactOpts, register)
}

// RegisterIsValid is a paid mutator transaction binding the contract method 0x53eeece1.
//
// Solidity: function registerIsValid(address register) returns(bool)
func (_ILeague *ILeagueTransactorSession) RegisterIsValid(register common.Address) (*types.Transaction, error) {
	return _ILeague.Contract.RegisterIsValid(&_ILeague.TransactOpts, register)
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
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRaceCreated is a free log retrieval operation binding the contract event 0x3d4c26760d0e4825f9a3f6ccba234d54aeeae3e6471639b8366708a183f1e837.
//
// Solidity: event RaceCreated(bytes32 id, uint256 slots, uint256 startAt)
func (_ILeague *ILeagueFilterer) FilterRaceCreated(opts *bind.FilterOpts) (*ILeagueRaceCreatedIterator, error) {

	logs, sub, err := _ILeague.contract.FilterLogs(opts, "RaceCreated")
	if err != nil {
		return nil, err
	}
	return &ILeagueRaceCreatedIterator{contract: _ILeague.contract, event: "RaceCreated", logs: logs, sub: sub}, nil
}

// WatchRaceCreated is a free log subscription operation binding the contract event 0x3d4c26760d0e4825f9a3f6ccba234d54aeeae3e6471639b8366708a183f1e837.
//
// Solidity: event RaceCreated(bytes32 id, uint256 slots, uint256 startAt)
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

// ParseRaceCreated is a log parse operation binding the contract event 0x3d4c26760d0e4825f9a3f6ccba234d54aeeae3e6471639b8366708a183f1e837.
//
// Solidity: event RaceCreated(bytes32 id, uint256 slots, uint256 startAt)
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
