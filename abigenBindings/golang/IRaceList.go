// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package IRaceList

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

// IRaceListMetaData contains all meta data concerning the IRaceList contract.
var IRaceListMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"InvalidCommission\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidRewardRate\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidSlot\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RaceNotExisted\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"}],\"name\":\"TooEarly\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"}],\"name\":\"TooLate\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"}],\"name\":\"RaceCancelled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"commission\",\"type\":\"uint256\"}],\"name\":\"RaceCommissionUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"slots\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"betStarted\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"betEnded\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"commission\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"rewardRate\",\"type\":\"uint256\"}],\"name\":\"RaceCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"result\",\"type\":\"bytes32\"}],\"name\":\"RaceResultUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"rewardRate\",\"type\":\"uint256\"}],\"name\":\"RaceRewardRateUpdated\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"slots\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"betStarted\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"betEnded\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"commission\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rewardRate\",\"type\":\"uint256\"}],\"name\":\"createRace\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"}],\"name\":\"getRace\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"commission\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rewardRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"slots\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"betStarted\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"betEnded\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"result\",\"type\":\"bytes32\"}],\"internalType\":\"structIRace.Race\",\"name\":\"race\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"}],\"name\":\"cancelRace\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"result\",\"type\":\"bytes32\"}],\"name\":\"updateResult\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"commission\",\"type\":\"uint256\"}],\"name\":\"updateCommission\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"rewardRate\",\"type\":\"uint256\"}],\"name\":\"updateRewardRate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"raceId\",\"type\":\"bytes32\"}],\"name\":\"raceIsExisted\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"raceId\",\"type\":\"bytes32\"}],\"name\":\"raceResult\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// IRaceListABI is the input ABI used to generate the binding from.
// Deprecated: Use IRaceListMetaData.ABI instead.
var IRaceListABI = IRaceListMetaData.ABI

// IRaceList is an auto generated Go binding around an Ethereum contract.
type IRaceList struct {
	IRaceListCaller     // Read-only binding to the contract
	IRaceListTransactor // Write-only binding to the contract
	IRaceListFilterer   // Log filterer for contract events
}

// IRaceListCaller is an auto generated read-only Go binding around an Ethereum contract.
type IRaceListCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IRaceListTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IRaceListTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IRaceListFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IRaceListFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IRaceListSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IRaceListSession struct {
	Contract     *IRaceList        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IRaceListCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IRaceListCallerSession struct {
	Contract *IRaceListCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// IRaceListTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IRaceListTransactorSession struct {
	Contract     *IRaceListTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// IRaceListRaw is an auto generated low-level Go binding around an Ethereum contract.
type IRaceListRaw struct {
	Contract *IRaceList // Generic contract binding to access the raw methods on
}

// IRaceListCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IRaceListCallerRaw struct {
	Contract *IRaceListCaller // Generic read-only contract binding to access the raw methods on
}

// IRaceListTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IRaceListTransactorRaw struct {
	Contract *IRaceListTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIRaceList creates a new instance of IRaceList, bound to a specific deployed contract.
func NewIRaceList(address common.Address, backend bind.ContractBackend) (*IRaceList, error) {
	contract, err := bindIRaceList(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IRaceList{IRaceListCaller: IRaceListCaller{contract: contract}, IRaceListTransactor: IRaceListTransactor{contract: contract}, IRaceListFilterer: IRaceListFilterer{contract: contract}}, nil
}

// NewIRaceListCaller creates a new read-only instance of IRaceList, bound to a specific deployed contract.
func NewIRaceListCaller(address common.Address, caller bind.ContractCaller) (*IRaceListCaller, error) {
	contract, err := bindIRaceList(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IRaceListCaller{contract: contract}, nil
}

// NewIRaceListTransactor creates a new write-only instance of IRaceList, bound to a specific deployed contract.
func NewIRaceListTransactor(address common.Address, transactor bind.ContractTransactor) (*IRaceListTransactor, error) {
	contract, err := bindIRaceList(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IRaceListTransactor{contract: contract}, nil
}

// NewIRaceListFilterer creates a new log filterer instance of IRaceList, bound to a specific deployed contract.
func NewIRaceListFilterer(address common.Address, filterer bind.ContractFilterer) (*IRaceListFilterer, error) {
	contract, err := bindIRaceList(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IRaceListFilterer{contract: contract}, nil
}

// bindIRaceList binds a generic wrapper to an already deployed contract.
func bindIRaceList(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IRaceListABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IRaceList *IRaceListRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IRaceList.Contract.IRaceListCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IRaceList *IRaceListRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IRaceList.Contract.IRaceListTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IRaceList *IRaceListRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IRaceList.Contract.IRaceListTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IRaceList *IRaceListCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IRaceList.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IRaceList *IRaceListTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IRaceList.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IRaceList *IRaceListTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IRaceList.Contract.contract.Transact(opts, method, params...)
}

// GetRace is a free data retrieval call binding the contract method 0xf70125db.
//
// Solidity: function getRace(bytes32 id) view returns((uint256,uint256,uint256,uint256,uint256,bytes32) race)
func (_IRaceList *IRaceListCaller) GetRace(opts *bind.CallOpts, id [32]byte) (IRaceRace, error) {
	var out []interface{}
	err := _IRaceList.contract.Call(opts, &out, "getRace", id)

	if err != nil {
		return *new(IRaceRace), err
	}

	out0 := *abi.ConvertType(out[0], new(IRaceRace)).(*IRaceRace)

	return out0, err

}

// GetRace is a free data retrieval call binding the contract method 0xf70125db.
//
// Solidity: function getRace(bytes32 id) view returns((uint256,uint256,uint256,uint256,uint256,bytes32) race)
func (_IRaceList *IRaceListSession) GetRace(id [32]byte) (IRaceRace, error) {
	return _IRaceList.Contract.GetRace(&_IRaceList.CallOpts, id)
}

// GetRace is a free data retrieval call binding the contract method 0xf70125db.
//
// Solidity: function getRace(bytes32 id) view returns((uint256,uint256,uint256,uint256,uint256,bytes32) race)
func (_IRaceList *IRaceListCallerSession) GetRace(id [32]byte) (IRaceRace, error) {
	return _IRaceList.Contract.GetRace(&_IRaceList.CallOpts, id)
}

// RaceIsExisted is a free data retrieval call binding the contract method 0x2869ad12.
//
// Solidity: function raceIsExisted(bytes32 raceId) view returns(bool)
func (_IRaceList *IRaceListCaller) RaceIsExisted(opts *bind.CallOpts, raceId [32]byte) (bool, error) {
	var out []interface{}
	err := _IRaceList.contract.Call(opts, &out, "raceIsExisted", raceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// RaceIsExisted is a free data retrieval call binding the contract method 0x2869ad12.
//
// Solidity: function raceIsExisted(bytes32 raceId) view returns(bool)
func (_IRaceList *IRaceListSession) RaceIsExisted(raceId [32]byte) (bool, error) {
	return _IRaceList.Contract.RaceIsExisted(&_IRaceList.CallOpts, raceId)
}

// RaceIsExisted is a free data retrieval call binding the contract method 0x2869ad12.
//
// Solidity: function raceIsExisted(bytes32 raceId) view returns(bool)
func (_IRaceList *IRaceListCallerSession) RaceIsExisted(raceId [32]byte) (bool, error) {
	return _IRaceList.Contract.RaceIsExisted(&_IRaceList.CallOpts, raceId)
}

// RaceResult is a free data retrieval call binding the contract method 0x6d93b958.
//
// Solidity: function raceResult(bytes32 raceId) view returns(bytes32)
func (_IRaceList *IRaceListCaller) RaceResult(opts *bind.CallOpts, raceId [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _IRaceList.contract.Call(opts, &out, "raceResult", raceId)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// RaceResult is a free data retrieval call binding the contract method 0x6d93b958.
//
// Solidity: function raceResult(bytes32 raceId) view returns(bytes32)
func (_IRaceList *IRaceListSession) RaceResult(raceId [32]byte) ([32]byte, error) {
	return _IRaceList.Contract.RaceResult(&_IRaceList.CallOpts, raceId)
}

// RaceResult is a free data retrieval call binding the contract method 0x6d93b958.
//
// Solidity: function raceResult(bytes32 raceId) view returns(bytes32)
func (_IRaceList *IRaceListCallerSession) RaceResult(raceId [32]byte) ([32]byte, error) {
	return _IRaceList.Contract.RaceResult(&_IRaceList.CallOpts, raceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_IRaceList *IRaceListCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _IRaceList.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_IRaceList *IRaceListSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _IRaceList.Contract.SupportsInterface(&_IRaceList.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_IRaceList *IRaceListCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _IRaceList.Contract.SupportsInterface(&_IRaceList.CallOpts, interfaceId)
}

// CancelRace is a paid mutator transaction binding the contract method 0xc05c107d.
//
// Solidity: function cancelRace(bytes32 id) returns()
func (_IRaceList *IRaceListTransactor) CancelRace(opts *bind.TransactOpts, id [32]byte) (*types.Transaction, error) {
	return _IRaceList.contract.Transact(opts, "cancelRace", id)
}

// CancelRace is a paid mutator transaction binding the contract method 0xc05c107d.
//
// Solidity: function cancelRace(bytes32 id) returns()
func (_IRaceList *IRaceListSession) CancelRace(id [32]byte) (*types.Transaction, error) {
	return _IRaceList.Contract.CancelRace(&_IRaceList.TransactOpts, id)
}

// CancelRace is a paid mutator transaction binding the contract method 0xc05c107d.
//
// Solidity: function cancelRace(bytes32 id) returns()
func (_IRaceList *IRaceListTransactorSession) CancelRace(id [32]byte) (*types.Transaction, error) {
	return _IRaceList.Contract.CancelRace(&_IRaceList.TransactOpts, id)
}

// CreateRace is a paid mutator transaction binding the contract method 0xd65b869a.
//
// Solidity: function createRace(uint256 slots, uint256 betStarted, uint256 betEnded, uint256 commission, uint256 rewardRate) returns()
func (_IRaceList *IRaceListTransactor) CreateRace(opts *bind.TransactOpts, slots *big.Int, betStarted *big.Int, betEnded *big.Int, commission *big.Int, rewardRate *big.Int) (*types.Transaction, error) {
	return _IRaceList.contract.Transact(opts, "createRace", slots, betStarted, betEnded, commission, rewardRate)
}

// CreateRace is a paid mutator transaction binding the contract method 0xd65b869a.
//
// Solidity: function createRace(uint256 slots, uint256 betStarted, uint256 betEnded, uint256 commission, uint256 rewardRate) returns()
func (_IRaceList *IRaceListSession) CreateRace(slots *big.Int, betStarted *big.Int, betEnded *big.Int, commission *big.Int, rewardRate *big.Int) (*types.Transaction, error) {
	return _IRaceList.Contract.CreateRace(&_IRaceList.TransactOpts, slots, betStarted, betEnded, commission, rewardRate)
}

// CreateRace is a paid mutator transaction binding the contract method 0xd65b869a.
//
// Solidity: function createRace(uint256 slots, uint256 betStarted, uint256 betEnded, uint256 commission, uint256 rewardRate) returns()
func (_IRaceList *IRaceListTransactorSession) CreateRace(slots *big.Int, betStarted *big.Int, betEnded *big.Int, commission *big.Int, rewardRate *big.Int) (*types.Transaction, error) {
	return _IRaceList.Contract.CreateRace(&_IRaceList.TransactOpts, slots, betStarted, betEnded, commission, rewardRate)
}

// UpdateCommission is a paid mutator transaction binding the contract method 0xc58f80a7.
//
// Solidity: function updateCommission(bytes32 id, uint256 commission) returns()
func (_IRaceList *IRaceListTransactor) UpdateCommission(opts *bind.TransactOpts, id [32]byte, commission *big.Int) (*types.Transaction, error) {
	return _IRaceList.contract.Transact(opts, "updateCommission", id, commission)
}

// UpdateCommission is a paid mutator transaction binding the contract method 0xc58f80a7.
//
// Solidity: function updateCommission(bytes32 id, uint256 commission) returns()
func (_IRaceList *IRaceListSession) UpdateCommission(id [32]byte, commission *big.Int) (*types.Transaction, error) {
	return _IRaceList.Contract.UpdateCommission(&_IRaceList.TransactOpts, id, commission)
}

// UpdateCommission is a paid mutator transaction binding the contract method 0xc58f80a7.
//
// Solidity: function updateCommission(bytes32 id, uint256 commission) returns()
func (_IRaceList *IRaceListTransactorSession) UpdateCommission(id [32]byte, commission *big.Int) (*types.Transaction, error) {
	return _IRaceList.Contract.UpdateCommission(&_IRaceList.TransactOpts, id, commission)
}

// UpdateResult is a paid mutator transaction binding the contract method 0xb6d940fc.
//
// Solidity: function updateResult(bytes32 id, bytes32 result) returns()
func (_IRaceList *IRaceListTransactor) UpdateResult(opts *bind.TransactOpts, id [32]byte, result [32]byte) (*types.Transaction, error) {
	return _IRaceList.contract.Transact(opts, "updateResult", id, result)
}

// UpdateResult is a paid mutator transaction binding the contract method 0xb6d940fc.
//
// Solidity: function updateResult(bytes32 id, bytes32 result) returns()
func (_IRaceList *IRaceListSession) UpdateResult(id [32]byte, result [32]byte) (*types.Transaction, error) {
	return _IRaceList.Contract.UpdateResult(&_IRaceList.TransactOpts, id, result)
}

// UpdateResult is a paid mutator transaction binding the contract method 0xb6d940fc.
//
// Solidity: function updateResult(bytes32 id, bytes32 result) returns()
func (_IRaceList *IRaceListTransactorSession) UpdateResult(id [32]byte, result [32]byte) (*types.Transaction, error) {
	return _IRaceList.Contract.UpdateResult(&_IRaceList.TransactOpts, id, result)
}

// UpdateRewardRate is a paid mutator transaction binding the contract method 0x68b2c40c.
//
// Solidity: function updateRewardRate(bytes32 id, uint256 rewardRate) returns()
func (_IRaceList *IRaceListTransactor) UpdateRewardRate(opts *bind.TransactOpts, id [32]byte, rewardRate *big.Int) (*types.Transaction, error) {
	return _IRaceList.contract.Transact(opts, "updateRewardRate", id, rewardRate)
}

// UpdateRewardRate is a paid mutator transaction binding the contract method 0x68b2c40c.
//
// Solidity: function updateRewardRate(bytes32 id, uint256 rewardRate) returns()
func (_IRaceList *IRaceListSession) UpdateRewardRate(id [32]byte, rewardRate *big.Int) (*types.Transaction, error) {
	return _IRaceList.Contract.UpdateRewardRate(&_IRaceList.TransactOpts, id, rewardRate)
}

// UpdateRewardRate is a paid mutator transaction binding the contract method 0x68b2c40c.
//
// Solidity: function updateRewardRate(bytes32 id, uint256 rewardRate) returns()
func (_IRaceList *IRaceListTransactorSession) UpdateRewardRate(id [32]byte, rewardRate *big.Int) (*types.Transaction, error) {
	return _IRaceList.Contract.UpdateRewardRate(&_IRaceList.TransactOpts, id, rewardRate)
}

// IRaceListRaceCancelledIterator is returned from FilterRaceCancelled and is used to iterate over the raw logs and unpacked data for RaceCancelled events raised by the IRaceList contract.
type IRaceListRaceCancelledIterator struct {
	Event *IRaceListRaceCancelled // Event containing the contract specifics and raw log

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
func (it *IRaceListRaceCancelledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IRaceListRaceCancelled)
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
		it.Event = new(IRaceListRaceCancelled)
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
func (it *IRaceListRaceCancelledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IRaceListRaceCancelledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IRaceListRaceCancelled represents a RaceCancelled event raised by the IRaceList contract.
type IRaceListRaceCancelled struct {
	Id  [32]byte
	Raw types.Log // Blockchain specific contextual infos
}

// FilterRaceCancelled is a free log retrieval operation binding the contract event 0x3f3669b11af9b00fb2a95f8ded36a79a2e1d5374efbc5ec6b2b5b96ceaf9ac6a.
//
// Solidity: event RaceCancelled(bytes32 id)
func (_IRaceList *IRaceListFilterer) FilterRaceCancelled(opts *bind.FilterOpts) (*IRaceListRaceCancelledIterator, error) {

	logs, sub, err := _IRaceList.contract.FilterLogs(opts, "RaceCancelled")
	if err != nil {
		return nil, err
	}
	return &IRaceListRaceCancelledIterator{contract: _IRaceList.contract, event: "RaceCancelled", logs: logs, sub: sub}, nil
}

// WatchRaceCancelled is a free log subscription operation binding the contract event 0x3f3669b11af9b00fb2a95f8ded36a79a2e1d5374efbc5ec6b2b5b96ceaf9ac6a.
//
// Solidity: event RaceCancelled(bytes32 id)
func (_IRaceList *IRaceListFilterer) WatchRaceCancelled(opts *bind.WatchOpts, sink chan<- *IRaceListRaceCancelled) (event.Subscription, error) {

	logs, sub, err := _IRaceList.contract.WatchLogs(opts, "RaceCancelled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IRaceListRaceCancelled)
				if err := _IRaceList.contract.UnpackLog(event, "RaceCancelled", log); err != nil {
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
func (_IRaceList *IRaceListFilterer) ParseRaceCancelled(log types.Log) (*IRaceListRaceCancelled, error) {
	event := new(IRaceListRaceCancelled)
	if err := _IRaceList.contract.UnpackLog(event, "RaceCancelled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IRaceListRaceCommissionUpdatedIterator is returned from FilterRaceCommissionUpdated and is used to iterate over the raw logs and unpacked data for RaceCommissionUpdated events raised by the IRaceList contract.
type IRaceListRaceCommissionUpdatedIterator struct {
	Event *IRaceListRaceCommissionUpdated // Event containing the contract specifics and raw log

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
func (it *IRaceListRaceCommissionUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IRaceListRaceCommissionUpdated)
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
		it.Event = new(IRaceListRaceCommissionUpdated)
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
func (it *IRaceListRaceCommissionUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IRaceListRaceCommissionUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IRaceListRaceCommissionUpdated represents a RaceCommissionUpdated event raised by the IRaceList contract.
type IRaceListRaceCommissionUpdated struct {
	Id         [32]byte
	Commission *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterRaceCommissionUpdated is a free log retrieval operation binding the contract event 0x9ff8c0f5ebec7767df2725581a4945a5a800a1f929ec51da017dc56d31a2b3a0.
//
// Solidity: event RaceCommissionUpdated(bytes32 id, uint256 commission)
func (_IRaceList *IRaceListFilterer) FilterRaceCommissionUpdated(opts *bind.FilterOpts) (*IRaceListRaceCommissionUpdatedIterator, error) {

	logs, sub, err := _IRaceList.contract.FilterLogs(opts, "RaceCommissionUpdated")
	if err != nil {
		return nil, err
	}
	return &IRaceListRaceCommissionUpdatedIterator{contract: _IRaceList.contract, event: "RaceCommissionUpdated", logs: logs, sub: sub}, nil
}

// WatchRaceCommissionUpdated is a free log subscription operation binding the contract event 0x9ff8c0f5ebec7767df2725581a4945a5a800a1f929ec51da017dc56d31a2b3a0.
//
// Solidity: event RaceCommissionUpdated(bytes32 id, uint256 commission)
func (_IRaceList *IRaceListFilterer) WatchRaceCommissionUpdated(opts *bind.WatchOpts, sink chan<- *IRaceListRaceCommissionUpdated) (event.Subscription, error) {

	logs, sub, err := _IRaceList.contract.WatchLogs(opts, "RaceCommissionUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IRaceListRaceCommissionUpdated)
				if err := _IRaceList.contract.UnpackLog(event, "RaceCommissionUpdated", log); err != nil {
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
func (_IRaceList *IRaceListFilterer) ParseRaceCommissionUpdated(log types.Log) (*IRaceListRaceCommissionUpdated, error) {
	event := new(IRaceListRaceCommissionUpdated)
	if err := _IRaceList.contract.UnpackLog(event, "RaceCommissionUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IRaceListRaceCreatedIterator is returned from FilterRaceCreated and is used to iterate over the raw logs and unpacked data for RaceCreated events raised by the IRaceList contract.
type IRaceListRaceCreatedIterator struct {
	Event *IRaceListRaceCreated // Event containing the contract specifics and raw log

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
func (it *IRaceListRaceCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IRaceListRaceCreated)
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
		it.Event = new(IRaceListRaceCreated)
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
func (it *IRaceListRaceCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IRaceListRaceCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IRaceListRaceCreated represents a RaceCreated event raised by the IRaceList contract.
type IRaceListRaceCreated struct {
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
func (_IRaceList *IRaceListFilterer) FilterRaceCreated(opts *bind.FilterOpts) (*IRaceListRaceCreatedIterator, error) {

	logs, sub, err := _IRaceList.contract.FilterLogs(opts, "RaceCreated")
	if err != nil {
		return nil, err
	}
	return &IRaceListRaceCreatedIterator{contract: _IRaceList.contract, event: "RaceCreated", logs: logs, sub: sub}, nil
}

// WatchRaceCreated is a free log subscription operation binding the contract event 0x1da8002ad97980e4380172489894c1e0e2d36896f491ff1d1773613fac4c8a22.
//
// Solidity: event RaceCreated(bytes32 id, uint256 slots, uint256 betStarted, uint256 betEnded, uint256 commission, uint256 rewardRate)
func (_IRaceList *IRaceListFilterer) WatchRaceCreated(opts *bind.WatchOpts, sink chan<- *IRaceListRaceCreated) (event.Subscription, error) {

	logs, sub, err := _IRaceList.contract.WatchLogs(opts, "RaceCreated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IRaceListRaceCreated)
				if err := _IRaceList.contract.UnpackLog(event, "RaceCreated", log); err != nil {
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
func (_IRaceList *IRaceListFilterer) ParseRaceCreated(log types.Log) (*IRaceListRaceCreated, error) {
	event := new(IRaceListRaceCreated)
	if err := _IRaceList.contract.UnpackLog(event, "RaceCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IRaceListRaceResultUpdatedIterator is returned from FilterRaceResultUpdated and is used to iterate over the raw logs and unpacked data for RaceResultUpdated events raised by the IRaceList contract.
type IRaceListRaceResultUpdatedIterator struct {
	Event *IRaceListRaceResultUpdated // Event containing the contract specifics and raw log

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
func (it *IRaceListRaceResultUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IRaceListRaceResultUpdated)
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
		it.Event = new(IRaceListRaceResultUpdated)
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
func (it *IRaceListRaceResultUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IRaceListRaceResultUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IRaceListRaceResultUpdated represents a RaceResultUpdated event raised by the IRaceList contract.
type IRaceListRaceResultUpdated struct {
	Id     [32]byte
	Result [32]byte
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRaceResultUpdated is a free log retrieval operation binding the contract event 0xd82663592968d73ce1995154e44c793fcb46a4006abfb7438656fb0d7ba5ff49.
//
// Solidity: event RaceResultUpdated(bytes32 id, bytes32 result)
func (_IRaceList *IRaceListFilterer) FilterRaceResultUpdated(opts *bind.FilterOpts) (*IRaceListRaceResultUpdatedIterator, error) {

	logs, sub, err := _IRaceList.contract.FilterLogs(opts, "RaceResultUpdated")
	if err != nil {
		return nil, err
	}
	return &IRaceListRaceResultUpdatedIterator{contract: _IRaceList.contract, event: "RaceResultUpdated", logs: logs, sub: sub}, nil
}

// WatchRaceResultUpdated is a free log subscription operation binding the contract event 0xd82663592968d73ce1995154e44c793fcb46a4006abfb7438656fb0d7ba5ff49.
//
// Solidity: event RaceResultUpdated(bytes32 id, bytes32 result)
func (_IRaceList *IRaceListFilterer) WatchRaceResultUpdated(opts *bind.WatchOpts, sink chan<- *IRaceListRaceResultUpdated) (event.Subscription, error) {

	logs, sub, err := _IRaceList.contract.WatchLogs(opts, "RaceResultUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IRaceListRaceResultUpdated)
				if err := _IRaceList.contract.UnpackLog(event, "RaceResultUpdated", log); err != nil {
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
func (_IRaceList *IRaceListFilterer) ParseRaceResultUpdated(log types.Log) (*IRaceListRaceResultUpdated, error) {
	event := new(IRaceListRaceResultUpdated)
	if err := _IRaceList.contract.UnpackLog(event, "RaceResultUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IRaceListRaceRewardRateUpdatedIterator is returned from FilterRaceRewardRateUpdated and is used to iterate over the raw logs and unpacked data for RaceRewardRateUpdated events raised by the IRaceList contract.
type IRaceListRaceRewardRateUpdatedIterator struct {
	Event *IRaceListRaceRewardRateUpdated // Event containing the contract specifics and raw log

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
func (it *IRaceListRaceRewardRateUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IRaceListRaceRewardRateUpdated)
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
		it.Event = new(IRaceListRaceRewardRateUpdated)
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
func (it *IRaceListRaceRewardRateUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IRaceListRaceRewardRateUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IRaceListRaceRewardRateUpdated represents a RaceRewardRateUpdated event raised by the IRaceList contract.
type IRaceListRaceRewardRateUpdated struct {
	Id         [32]byte
	RewardRate *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterRaceRewardRateUpdated is a free log retrieval operation binding the contract event 0x25a0545ede30f4217593809347e09822194f6a9fe74a5b3f2c5b6e843375f422.
//
// Solidity: event RaceRewardRateUpdated(bytes32 id, uint256 rewardRate)
func (_IRaceList *IRaceListFilterer) FilterRaceRewardRateUpdated(opts *bind.FilterOpts) (*IRaceListRaceRewardRateUpdatedIterator, error) {

	logs, sub, err := _IRaceList.contract.FilterLogs(opts, "RaceRewardRateUpdated")
	if err != nil {
		return nil, err
	}
	return &IRaceListRaceRewardRateUpdatedIterator{contract: _IRaceList.contract, event: "RaceRewardRateUpdated", logs: logs, sub: sub}, nil
}

// WatchRaceRewardRateUpdated is a free log subscription operation binding the contract event 0x25a0545ede30f4217593809347e09822194f6a9fe74a5b3f2c5b6e843375f422.
//
// Solidity: event RaceRewardRateUpdated(bytes32 id, uint256 rewardRate)
func (_IRaceList *IRaceListFilterer) WatchRaceRewardRateUpdated(opts *bind.WatchOpts, sink chan<- *IRaceListRaceRewardRateUpdated) (event.Subscription, error) {

	logs, sub, err := _IRaceList.contract.WatchLogs(opts, "RaceRewardRateUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IRaceListRaceRewardRateUpdated)
				if err := _IRaceList.contract.UnpackLog(event, "RaceRewardRateUpdated", log); err != nil {
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
func (_IRaceList *IRaceListFilterer) ParseRaceRewardRateUpdated(log types.Log) (*IRaceListRaceRewardRateUpdated, error) {
	event := new(IRaceListRaceRewardRateUpdated)
	if err := _IRaceList.contract.UnpackLog(event, "RaceRewardRateUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
