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
	ABI: "[{\"inputs\":[],\"name\":\"AlreadyRegistered\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AlreadySelected\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidContract\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidSender\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidSlot\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MaximumReached\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotNLGGTHolder\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RaceNotExisted\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RewardIsExisted\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RewardNotExistedOrReceived\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"}],\"name\":\"TooEarly\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"}],\"name\":\"TooLate\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"nlggt\",\"type\":\"address\"}],\"name\":\"NlggtUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"requestId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"raceId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"randomness\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"selected\",\"type\":\"bytes32\"}],\"name\":\"ParticipantsSelected\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"race\",\"type\":\"address\"}],\"name\":\"RaceListUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"raceId\",\"type\":\"bytes32\"}],\"name\":\"RandomInProgress\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"slotId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"participant\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"raceId\",\"type\":\"bytes32\"}],\"name\":\"Registered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"raceId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nftRewardId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes1\",\"name\":\"resultIndex\",\"type\":\"bytes1\"}],\"name\":\"RewardAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"raceId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"slotId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nftRewardId\",\"type\":\"uint256\"}],\"name\":\"RewardReceived\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"slotId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"raceId\",\"type\":\"bytes32\"}],\"name\":\"register\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"raceId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"nftRewardId\",\"type\":\"uint256\"},{\"internalType\":\"bytes1\",\"name\":\"resultIndex\",\"type\":\"bytes1\"}],\"name\":\"addReward\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"raceId\",\"type\":\"bytes32\"}],\"name\":\"selectParticipant\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"requestId\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"raceList\",\"type\":\"address\"}],\"name\":\"updateRaceAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"nlggt\",\"type\":\"address\"}],\"name\":\"updateNlggtAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"raceId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"slotId\",\"type\":\"uint256\"}],\"name\":\"receiveReward\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
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

// AddReward is a paid mutator transaction binding the contract method 0x9a789d92.
//
// Solidity: function addReward(bytes32 raceId, uint256 nftRewardId, bytes1 resultIndex) returns()
func (_Smc *SmcTransactor) AddReward(opts *bind.TransactOpts, raceId [32]byte, nftRewardId *big.Int, resultIndex [1]byte) (*types.Transaction, error) {
	return _Smc.contract.Transact(opts, "addReward", raceId, nftRewardId, resultIndex)
}

// AddReward is a paid mutator transaction binding the contract method 0x9a789d92.
//
// Solidity: function addReward(bytes32 raceId, uint256 nftRewardId, bytes1 resultIndex) returns()
func (_Smc *SmcSession) AddReward(raceId [32]byte, nftRewardId *big.Int, resultIndex [1]byte) (*types.Transaction, error) {
	return _Smc.Contract.AddReward(&_Smc.TransactOpts, raceId, nftRewardId, resultIndex)
}

// AddReward is a paid mutator transaction binding the contract method 0x9a789d92.
//
// Solidity: function addReward(bytes32 raceId, uint256 nftRewardId, bytes1 resultIndex) returns()
func (_Smc *SmcTransactorSession) AddReward(raceId [32]byte, nftRewardId *big.Int, resultIndex [1]byte) (*types.Transaction, error) {
	return _Smc.Contract.AddReward(&_Smc.TransactOpts, raceId, nftRewardId, resultIndex)
}

// ReceiveReward is a paid mutator transaction binding the contract method 0x1a0fc715.
//
// Solidity: function receiveReward(bytes32 raceId, uint256 slotId) returns()
func (_Smc *SmcTransactor) ReceiveReward(opts *bind.TransactOpts, raceId [32]byte, slotId *big.Int) (*types.Transaction, error) {
	return _Smc.contract.Transact(opts, "receiveReward", raceId, slotId)
}

// ReceiveReward is a paid mutator transaction binding the contract method 0x1a0fc715.
//
// Solidity: function receiveReward(bytes32 raceId, uint256 slotId) returns()
func (_Smc *SmcSession) ReceiveReward(raceId [32]byte, slotId *big.Int) (*types.Transaction, error) {
	return _Smc.Contract.ReceiveReward(&_Smc.TransactOpts, raceId, slotId)
}

// ReceiveReward is a paid mutator transaction binding the contract method 0x1a0fc715.
//
// Solidity: function receiveReward(bytes32 raceId, uint256 slotId) returns()
func (_Smc *SmcTransactorSession) ReceiveReward(raceId [32]byte, slotId *big.Int) (*types.Transaction, error) {
	return _Smc.Contract.ReceiveReward(&_Smc.TransactOpts, raceId, slotId)
}

// Register is a paid mutator transaction binding the contract method 0x2965d809.
//
// Solidity: function register(uint256 slotId, bytes32 raceId) returns()
func (_Smc *SmcTransactor) Register(opts *bind.TransactOpts, slotId *big.Int, raceId [32]byte) (*types.Transaction, error) {
	return _Smc.contract.Transact(opts, "register", slotId, raceId)
}

// Register is a paid mutator transaction binding the contract method 0x2965d809.
//
// Solidity: function register(uint256 slotId, bytes32 raceId) returns()
func (_Smc *SmcSession) Register(slotId *big.Int, raceId [32]byte) (*types.Transaction, error) {
	return _Smc.Contract.Register(&_Smc.TransactOpts, slotId, raceId)
}

// Register is a paid mutator transaction binding the contract method 0x2965d809.
//
// Solidity: function register(uint256 slotId, bytes32 raceId) returns()
func (_Smc *SmcTransactorSession) Register(slotId *big.Int, raceId [32]byte) (*types.Transaction, error) {
	return _Smc.Contract.Register(&_Smc.TransactOpts, slotId, raceId)
}

// SelectParticipant is a paid mutator transaction binding the contract method 0x7d0f79b7.
//
// Solidity: function selectParticipant(bytes32 raceId) returns(uint256 requestId)
func (_Smc *SmcTransactor) SelectParticipant(opts *bind.TransactOpts, raceId [32]byte) (*types.Transaction, error) {
	return _Smc.contract.Transact(opts, "selectParticipant", raceId)
}

// SelectParticipant is a paid mutator transaction binding the contract method 0x7d0f79b7.
//
// Solidity: function selectParticipant(bytes32 raceId) returns(uint256 requestId)
func (_Smc *SmcSession) SelectParticipant(raceId [32]byte) (*types.Transaction, error) {
	return _Smc.Contract.SelectParticipant(&_Smc.TransactOpts, raceId)
}

// SelectParticipant is a paid mutator transaction binding the contract method 0x7d0f79b7.
//
// Solidity: function selectParticipant(bytes32 raceId) returns(uint256 requestId)
func (_Smc *SmcTransactorSession) SelectParticipant(raceId [32]byte) (*types.Transaction, error) {
	return _Smc.Contract.SelectParticipant(&_Smc.TransactOpts, raceId)
}

// UpdateNlggtAddress is a paid mutator transaction binding the contract method 0xd6641f06.
//
// Solidity: function updateNlggtAddress(address nlggt) returns()
func (_Smc *SmcTransactor) UpdateNlggtAddress(opts *bind.TransactOpts, nlggt common.Address) (*types.Transaction, error) {
	return _Smc.contract.Transact(opts, "updateNlggtAddress", nlggt)
}

// UpdateNlggtAddress is a paid mutator transaction binding the contract method 0xd6641f06.
//
// Solidity: function updateNlggtAddress(address nlggt) returns()
func (_Smc *SmcSession) UpdateNlggtAddress(nlggt common.Address) (*types.Transaction, error) {
	return _Smc.Contract.UpdateNlggtAddress(&_Smc.TransactOpts, nlggt)
}

// UpdateNlggtAddress is a paid mutator transaction binding the contract method 0xd6641f06.
//
// Solidity: function updateNlggtAddress(address nlggt) returns()
func (_Smc *SmcTransactorSession) UpdateNlggtAddress(nlggt common.Address) (*types.Transaction, error) {
	return _Smc.Contract.UpdateNlggtAddress(&_Smc.TransactOpts, nlggt)
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

// SmcNlggtUpdatedIterator is returned from FilterNlggtUpdated and is used to iterate over the raw logs and unpacked data for NlggtUpdated events raised by the Smc contract.
type SmcNlggtUpdatedIterator struct {
	Event *SmcNlggtUpdated // Event containing the contract specifics and raw log

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
func (it *SmcNlggtUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SmcNlggtUpdated)
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
		it.Event = new(SmcNlggtUpdated)
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
func (it *SmcNlggtUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SmcNlggtUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SmcNlggtUpdated represents a NlggtUpdated event raised by the Smc contract.
type SmcNlggtUpdated struct {
	Nlggt common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterNlggtUpdated is a free log retrieval operation binding the contract event 0x5660a1e6d6115d84caa3bc631133061cbc77b46ff3be76d65178ae3243446b98.
//
// Solidity: event NlggtUpdated(address nlggt)
func (_Smc *SmcFilterer) FilterNlggtUpdated(opts *bind.FilterOpts) (*SmcNlggtUpdatedIterator, error) {

	logs, sub, err := _Smc.contract.FilterLogs(opts, "NlggtUpdated")
	if err != nil {
		return nil, err
	}
	return &SmcNlggtUpdatedIterator{contract: _Smc.contract, event: "NlggtUpdated", logs: logs, sub: sub}, nil
}

// WatchNlggtUpdated is a free log subscription operation binding the contract event 0x5660a1e6d6115d84caa3bc631133061cbc77b46ff3be76d65178ae3243446b98.
//
// Solidity: event NlggtUpdated(address nlggt)
func (_Smc *SmcFilterer) WatchNlggtUpdated(opts *bind.WatchOpts, sink chan<- *SmcNlggtUpdated) (event.Subscription, error) {

	logs, sub, err := _Smc.contract.WatchLogs(opts, "NlggtUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SmcNlggtUpdated)
				if err := _Smc.contract.UnpackLog(event, "NlggtUpdated", log); err != nil {
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
func (_Smc *SmcFilterer) ParseNlggtUpdated(log types.Log) (*SmcNlggtUpdated, error) {
	event := new(SmcNlggtUpdated)
	if err := _Smc.contract.UnpackLog(event, "NlggtUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SmcParticipantsSelectedIterator is returned from FilterParticipantsSelected and is used to iterate over the raw logs and unpacked data for ParticipantsSelected events raised by the Smc contract.
type SmcParticipantsSelectedIterator struct {
	Event *SmcParticipantsSelected // Event containing the contract specifics and raw log

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
func (it *SmcParticipantsSelectedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SmcParticipantsSelected)
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
		it.Event = new(SmcParticipantsSelected)
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
func (it *SmcParticipantsSelectedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SmcParticipantsSelectedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SmcParticipantsSelected represents a ParticipantsSelected event raised by the Smc contract.
type SmcParticipantsSelected struct {
	RequestId  *big.Int
	RaceId     [32]byte
	Randomness [32]byte
	Selected   [32]byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterParticipantsSelected is a free log retrieval operation binding the contract event 0xf67ae7b61862fbe089bafd38aea69def4c64db49c4bff6d29f038cfd13fdd4c2.
//
// Solidity: event ParticipantsSelected(uint256 requestId, bytes32 raceId, bytes32 randomness, bytes32 selected)
func (_Smc *SmcFilterer) FilterParticipantsSelected(opts *bind.FilterOpts) (*SmcParticipantsSelectedIterator, error) {

	logs, sub, err := _Smc.contract.FilterLogs(opts, "ParticipantsSelected")
	if err != nil {
		return nil, err
	}
	return &SmcParticipantsSelectedIterator{contract: _Smc.contract, event: "ParticipantsSelected", logs: logs, sub: sub}, nil
}

// WatchParticipantsSelected is a free log subscription operation binding the contract event 0xf67ae7b61862fbe089bafd38aea69def4c64db49c4bff6d29f038cfd13fdd4c2.
//
// Solidity: event ParticipantsSelected(uint256 requestId, bytes32 raceId, bytes32 randomness, bytes32 selected)
func (_Smc *SmcFilterer) WatchParticipantsSelected(opts *bind.WatchOpts, sink chan<- *SmcParticipantsSelected) (event.Subscription, error) {

	logs, sub, err := _Smc.contract.WatchLogs(opts, "ParticipantsSelected")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SmcParticipantsSelected)
				if err := _Smc.contract.UnpackLog(event, "ParticipantsSelected", log); err != nil {
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

// ParseParticipantsSelected is a log parse operation binding the contract event 0xf67ae7b61862fbe089bafd38aea69def4c64db49c4bff6d29f038cfd13fdd4c2.
//
// Solidity: event ParticipantsSelected(uint256 requestId, bytes32 raceId, bytes32 randomness, bytes32 selected)
func (_Smc *SmcFilterer) ParseParticipantsSelected(log types.Log) (*SmcParticipantsSelected, error) {
	event := new(SmcParticipantsSelected)
	if err := _Smc.contract.UnpackLog(event, "ParticipantsSelected", log); err != nil {
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

// SmcRandomInProgressIterator is returned from FilterRandomInProgress and is used to iterate over the raw logs and unpacked data for RandomInProgress events raised by the Smc contract.
type SmcRandomInProgressIterator struct {
	Event *SmcRandomInProgress // Event containing the contract specifics and raw log

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
func (it *SmcRandomInProgressIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SmcRandomInProgress)
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
		it.Event = new(SmcRandomInProgress)
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
func (it *SmcRandomInProgressIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SmcRandomInProgressIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SmcRandomInProgress represents a RandomInProgress event raised by the Smc contract.
type SmcRandomInProgress struct {
	RaceId [32]byte
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRandomInProgress is a free log retrieval operation binding the contract event 0x018dc112e9dad7765a03056b4091d237454ee37ae27aab3c3abd16d428900a13.
//
// Solidity: event RandomInProgress(bytes32 raceId)
func (_Smc *SmcFilterer) FilterRandomInProgress(opts *bind.FilterOpts) (*SmcRandomInProgressIterator, error) {

	logs, sub, err := _Smc.contract.FilterLogs(opts, "RandomInProgress")
	if err != nil {
		return nil, err
	}
	return &SmcRandomInProgressIterator{contract: _Smc.contract, event: "RandomInProgress", logs: logs, sub: sub}, nil
}

// WatchRandomInProgress is a free log subscription operation binding the contract event 0x018dc112e9dad7765a03056b4091d237454ee37ae27aab3c3abd16d428900a13.
//
// Solidity: event RandomInProgress(bytes32 raceId)
func (_Smc *SmcFilterer) WatchRandomInProgress(opts *bind.WatchOpts, sink chan<- *SmcRandomInProgress) (event.Subscription, error) {

	logs, sub, err := _Smc.contract.WatchLogs(opts, "RandomInProgress")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SmcRandomInProgress)
				if err := _Smc.contract.UnpackLog(event, "RandomInProgress", log); err != nil {
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
func (_Smc *SmcFilterer) ParseRandomInProgress(log types.Log) (*SmcRandomInProgress, error) {
	event := new(SmcRandomInProgress)
	if err := _Smc.contract.UnpackLog(event, "RandomInProgress", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SmcRegisteredIterator is returned from FilterRegistered and is used to iterate over the raw logs and unpacked data for Registered events raised by the Smc contract.
type SmcRegisteredIterator struct {
	Event *SmcRegistered // Event containing the contract specifics and raw log

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
func (it *SmcRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SmcRegistered)
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
		it.Event = new(SmcRegistered)
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
func (it *SmcRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SmcRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SmcRegistered represents a Registered event raised by the Smc contract.
type SmcRegistered struct {
	SlotId      *big.Int
	Participant common.Address
	RaceId      [32]byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterRegistered is a free log retrieval operation binding the contract event 0xd62c7bd2779dc6ebe426bed623399b1fc1d190c063eadefcffffb1efe6f253ef.
//
// Solidity: event Registered(uint256 slotId, address participant, bytes32 raceId)
func (_Smc *SmcFilterer) FilterRegistered(opts *bind.FilterOpts) (*SmcRegisteredIterator, error) {

	logs, sub, err := _Smc.contract.FilterLogs(opts, "Registered")
	if err != nil {
		return nil, err
	}
	return &SmcRegisteredIterator{contract: _Smc.contract, event: "Registered", logs: logs, sub: sub}, nil
}

// WatchRegistered is a free log subscription operation binding the contract event 0xd62c7bd2779dc6ebe426bed623399b1fc1d190c063eadefcffffb1efe6f253ef.
//
// Solidity: event Registered(uint256 slotId, address participant, bytes32 raceId)
func (_Smc *SmcFilterer) WatchRegistered(opts *bind.WatchOpts, sink chan<- *SmcRegistered) (event.Subscription, error) {

	logs, sub, err := _Smc.contract.WatchLogs(opts, "Registered")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SmcRegistered)
				if err := _Smc.contract.UnpackLog(event, "Registered", log); err != nil {
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
func (_Smc *SmcFilterer) ParseRegistered(log types.Log) (*SmcRegistered, error) {
	event := new(SmcRegistered)
	if err := _Smc.contract.UnpackLog(event, "Registered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SmcRewardAddedIterator is returned from FilterRewardAdded and is used to iterate over the raw logs and unpacked data for RewardAdded events raised by the Smc contract.
type SmcRewardAddedIterator struct {
	Event *SmcRewardAdded // Event containing the contract specifics and raw log

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
func (it *SmcRewardAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SmcRewardAdded)
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
		it.Event = new(SmcRewardAdded)
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
func (it *SmcRewardAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SmcRewardAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SmcRewardAdded represents a RewardAdded event raised by the Smc contract.
type SmcRewardAdded struct {
	RaceId      [32]byte
	NftRewardId *big.Int
	ResultIndex [1]byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterRewardAdded is a free log retrieval operation binding the contract event 0x2b942e265bd6d4ca5617a327f9a60e1a85f4351aaa0be2a84cda1cf9d32c8c04.
//
// Solidity: event RewardAdded(bytes32 raceId, uint256 nftRewardId, bytes1 resultIndex)
func (_Smc *SmcFilterer) FilterRewardAdded(opts *bind.FilterOpts) (*SmcRewardAddedIterator, error) {

	logs, sub, err := _Smc.contract.FilterLogs(opts, "RewardAdded")
	if err != nil {
		return nil, err
	}
	return &SmcRewardAddedIterator{contract: _Smc.contract, event: "RewardAdded", logs: logs, sub: sub}, nil
}

// WatchRewardAdded is a free log subscription operation binding the contract event 0x2b942e265bd6d4ca5617a327f9a60e1a85f4351aaa0be2a84cda1cf9d32c8c04.
//
// Solidity: event RewardAdded(bytes32 raceId, uint256 nftRewardId, bytes1 resultIndex)
func (_Smc *SmcFilterer) WatchRewardAdded(opts *bind.WatchOpts, sink chan<- *SmcRewardAdded) (event.Subscription, error) {

	logs, sub, err := _Smc.contract.WatchLogs(opts, "RewardAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SmcRewardAdded)
				if err := _Smc.contract.UnpackLog(event, "RewardAdded", log); err != nil {
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
func (_Smc *SmcFilterer) ParseRewardAdded(log types.Log) (*SmcRewardAdded, error) {
	event := new(SmcRewardAdded)
	if err := _Smc.contract.UnpackLog(event, "RewardAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SmcRewardReceivedIterator is returned from FilterRewardReceived and is used to iterate over the raw logs and unpacked data for RewardReceived events raised by the Smc contract.
type SmcRewardReceivedIterator struct {
	Event *SmcRewardReceived // Event containing the contract specifics and raw log

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
func (it *SmcRewardReceivedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SmcRewardReceived)
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
		it.Event = new(SmcRewardReceived)
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
func (it *SmcRewardReceivedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SmcRewardReceivedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SmcRewardReceived represents a RewardReceived event raised by the Smc contract.
type SmcRewardReceived struct {
	RaceId      [32]byte
	SlotId      *big.Int
	NftRewardId *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterRewardReceived is a free log retrieval operation binding the contract event 0xe2ece170b35ae79cbe45ffdaf4931b8bc45056da5d3e9aa4aea67fcdf307929a.
//
// Solidity: event RewardReceived(bytes32 raceId, uint256 slotId, uint256 nftRewardId)
func (_Smc *SmcFilterer) FilterRewardReceived(opts *bind.FilterOpts) (*SmcRewardReceivedIterator, error) {

	logs, sub, err := _Smc.contract.FilterLogs(opts, "RewardReceived")
	if err != nil {
		return nil, err
	}
	return &SmcRewardReceivedIterator{contract: _Smc.contract, event: "RewardReceived", logs: logs, sub: sub}, nil
}

// WatchRewardReceived is a free log subscription operation binding the contract event 0xe2ece170b35ae79cbe45ffdaf4931b8bc45056da5d3e9aa4aea67fcdf307929a.
//
// Solidity: event RewardReceived(bytes32 raceId, uint256 slotId, uint256 nftRewardId)
func (_Smc *SmcFilterer) WatchRewardReceived(opts *bind.WatchOpts, sink chan<- *SmcRewardReceived) (event.Subscription, error) {

	logs, sub, err := _Smc.contract.WatchLogs(opts, "RewardReceived")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SmcRewardReceived)
				if err := _Smc.contract.UnpackLog(event, "RewardReceived", log); err != nil {
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
func (_Smc *SmcFilterer) ParseRewardReceived(log types.Log) (*SmcRewardReceived, error) {
	event := new(SmcRewardReceived)
	if err := _Smc.contract.UnpackLog(event, "RewardReceived", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
