// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package IRegistration

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

// IRegistrationMetaData contains all meta data concerning the IRegistration contract.
var IRegistrationMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"AlreadyRegistered\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AlreadySelected\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidContract\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidSender\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidSlot\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RaceNotExisted\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RewardIsExisted\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RewardIsNotExistedOrReceived\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"}],\"name\":\"TooEarly\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"}],\"name\":\"TooLate\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"requestId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"raceId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"randomness\",\"type\":\"uint256\"}],\"name\":\"ParticipantsSelected\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"raceId\",\"type\":\"bytes32\"}],\"name\":\"RandomInProgress\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"slotId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"participant\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"raceId\",\"type\":\"bytes32\"}],\"name\":\"Registered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"raceId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nftRewardId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes1\",\"name\":\"resultIndex\",\"type\":\"bytes1\"}],\"name\":\"RewardAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"raceId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"slotId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nftRewardId\",\"type\":\"uint256\"}],\"name\":\"RewardReceived\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"raceId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nftRewardId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes1\",\"name\":\"resultIndex\",\"type\":\"bytes1\"}],\"name\":\"RewardRemoved\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"slotId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"raceId\",\"type\":\"bytes32\"}],\"name\":\"register\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"leagueAddr\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"raceId\",\"type\":\"bytes32\"}],\"name\":\"enableRegistration\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"raceId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"nftRewardId\",\"type\":\"uint256\"},{\"internalType\":\"bytes1\",\"name\":\"resultIndex\",\"type\":\"bytes1\"}],\"name\":\"addRewardByTransfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"raceId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"nftRewardId\",\"type\":\"uint256\"},{\"internalType\":\"bytes1\",\"name\":\"resultIndex\",\"type\":\"bytes1\"},{\"internalType\":\"string\",\"name\":\"tokenURI\",\"type\":\"string\"}],\"name\":\"addRewardByMint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"raceId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"nftRewardId\",\"type\":\"uint256\"},{\"internalType\":\"bytes1\",\"name\":\"resultIndex\",\"type\":\"bytes1\"}],\"name\":\"removeReward\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"raceId\",\"type\":\"bytes32\"}],\"name\":\"selectParticipant\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"requestId\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"raceId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"slotId\",\"type\":\"uint256\"}],\"name\":\"receiveReward\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// IRegistrationABI is the input ABI used to generate the binding from.
// Deprecated: Use IRegistrationMetaData.ABI instead.
var IRegistrationABI = IRegistrationMetaData.ABI

// IRegistration is an auto generated Go binding around an Ethereum contract.
type IRegistration struct {
	IRegistrationCaller     // Read-only binding to the contract
	IRegistrationTransactor // Write-only binding to the contract
	IRegistrationFilterer   // Log filterer for contract events
}

// IRegistrationCaller is an auto generated read-only Go binding around an Ethereum contract.
type IRegistrationCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IRegistrationTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IRegistrationTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IRegistrationFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IRegistrationFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IRegistrationSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IRegistrationSession struct {
	Contract     *IRegistration    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IRegistrationCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IRegistrationCallerSession struct {
	Contract *IRegistrationCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// IRegistrationTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IRegistrationTransactorSession struct {
	Contract     *IRegistrationTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// IRegistrationRaw is an auto generated low-level Go binding around an Ethereum contract.
type IRegistrationRaw struct {
	Contract *IRegistration // Generic contract binding to access the raw methods on
}

// IRegistrationCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IRegistrationCallerRaw struct {
	Contract *IRegistrationCaller // Generic read-only contract binding to access the raw methods on
}

// IRegistrationTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IRegistrationTransactorRaw struct {
	Contract *IRegistrationTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIRegistration creates a new instance of IRegistration, bound to a specific deployed contract.
func NewIRegistration(address common.Address, backend bind.ContractBackend) (*IRegistration, error) {
	contract, err := bindIRegistration(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IRegistration{IRegistrationCaller: IRegistrationCaller{contract: contract}, IRegistrationTransactor: IRegistrationTransactor{contract: contract}, IRegistrationFilterer: IRegistrationFilterer{contract: contract}}, nil
}

// NewIRegistrationCaller creates a new read-only instance of IRegistration, bound to a specific deployed contract.
func NewIRegistrationCaller(address common.Address, caller bind.ContractCaller) (*IRegistrationCaller, error) {
	contract, err := bindIRegistration(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IRegistrationCaller{contract: contract}, nil
}

// NewIRegistrationTransactor creates a new write-only instance of IRegistration, bound to a specific deployed contract.
func NewIRegistrationTransactor(address common.Address, transactor bind.ContractTransactor) (*IRegistrationTransactor, error) {
	contract, err := bindIRegistration(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IRegistrationTransactor{contract: contract}, nil
}

// NewIRegistrationFilterer creates a new log filterer instance of IRegistration, bound to a specific deployed contract.
func NewIRegistrationFilterer(address common.Address, filterer bind.ContractFilterer) (*IRegistrationFilterer, error) {
	contract, err := bindIRegistration(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IRegistrationFilterer{contract: contract}, nil
}

// bindIRegistration binds a generic wrapper to an already deployed contract.
func bindIRegistration(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IRegistrationABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IRegistration *IRegistrationRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IRegistration.Contract.IRegistrationCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IRegistration *IRegistrationRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IRegistration.Contract.IRegistrationTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IRegistration *IRegistrationRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IRegistration.Contract.IRegistrationTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IRegistration *IRegistrationCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IRegistration.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IRegistration *IRegistrationTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IRegistration.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IRegistration *IRegistrationTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IRegistration.Contract.contract.Transact(opts, method, params...)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_IRegistration *IRegistrationCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _IRegistration.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_IRegistration *IRegistrationSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _IRegistration.Contract.SupportsInterface(&_IRegistration.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_IRegistration *IRegistrationCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _IRegistration.Contract.SupportsInterface(&_IRegistration.CallOpts, interfaceId)
}

// AddRewardByMint is a paid mutator transaction binding the contract method 0x6bef76da.
//
// Solidity: function addRewardByMint(bytes32 raceId, uint256 nftRewardId, bytes1 resultIndex, string tokenURI) returns()
func (_IRegistration *IRegistrationTransactor) AddRewardByMint(opts *bind.TransactOpts, raceId [32]byte, nftRewardId *big.Int, resultIndex [1]byte, tokenURI string) (*types.Transaction, error) {
	return _IRegistration.contract.Transact(opts, "addRewardByMint", raceId, nftRewardId, resultIndex, tokenURI)
}

// AddRewardByMint is a paid mutator transaction binding the contract method 0x6bef76da.
//
// Solidity: function addRewardByMint(bytes32 raceId, uint256 nftRewardId, bytes1 resultIndex, string tokenURI) returns()
func (_IRegistration *IRegistrationSession) AddRewardByMint(raceId [32]byte, nftRewardId *big.Int, resultIndex [1]byte, tokenURI string) (*types.Transaction, error) {
	return _IRegistration.Contract.AddRewardByMint(&_IRegistration.TransactOpts, raceId, nftRewardId, resultIndex, tokenURI)
}

// AddRewardByMint is a paid mutator transaction binding the contract method 0x6bef76da.
//
// Solidity: function addRewardByMint(bytes32 raceId, uint256 nftRewardId, bytes1 resultIndex, string tokenURI) returns()
func (_IRegistration *IRegistrationTransactorSession) AddRewardByMint(raceId [32]byte, nftRewardId *big.Int, resultIndex [1]byte, tokenURI string) (*types.Transaction, error) {
	return _IRegistration.Contract.AddRewardByMint(&_IRegistration.TransactOpts, raceId, nftRewardId, resultIndex, tokenURI)
}

// AddRewardByTransfer is a paid mutator transaction binding the contract method 0x81357309.
//
// Solidity: function addRewardByTransfer(bytes32 raceId, uint256 nftRewardId, bytes1 resultIndex) returns()
func (_IRegistration *IRegistrationTransactor) AddRewardByTransfer(opts *bind.TransactOpts, raceId [32]byte, nftRewardId *big.Int, resultIndex [1]byte) (*types.Transaction, error) {
	return _IRegistration.contract.Transact(opts, "addRewardByTransfer", raceId, nftRewardId, resultIndex)
}

// AddRewardByTransfer is a paid mutator transaction binding the contract method 0x81357309.
//
// Solidity: function addRewardByTransfer(bytes32 raceId, uint256 nftRewardId, bytes1 resultIndex) returns()
func (_IRegistration *IRegistrationSession) AddRewardByTransfer(raceId [32]byte, nftRewardId *big.Int, resultIndex [1]byte) (*types.Transaction, error) {
	return _IRegistration.Contract.AddRewardByTransfer(&_IRegistration.TransactOpts, raceId, nftRewardId, resultIndex)
}

// AddRewardByTransfer is a paid mutator transaction binding the contract method 0x81357309.
//
// Solidity: function addRewardByTransfer(bytes32 raceId, uint256 nftRewardId, bytes1 resultIndex) returns()
func (_IRegistration *IRegistrationTransactorSession) AddRewardByTransfer(raceId [32]byte, nftRewardId *big.Int, resultIndex [1]byte) (*types.Transaction, error) {
	return _IRegistration.Contract.AddRewardByTransfer(&_IRegistration.TransactOpts, raceId, nftRewardId, resultIndex)
}

// EnableRegistration is a paid mutator transaction binding the contract method 0x52f3dcfb.
//
// Solidity: function enableRegistration(address leagueAddr, bytes32 raceId) returns()
func (_IRegistration *IRegistrationTransactor) EnableRegistration(opts *bind.TransactOpts, leagueAddr common.Address, raceId [32]byte) (*types.Transaction, error) {
	return _IRegistration.contract.Transact(opts, "enableRegistration", leagueAddr, raceId)
}

// EnableRegistration is a paid mutator transaction binding the contract method 0x52f3dcfb.
//
// Solidity: function enableRegistration(address leagueAddr, bytes32 raceId) returns()
func (_IRegistration *IRegistrationSession) EnableRegistration(leagueAddr common.Address, raceId [32]byte) (*types.Transaction, error) {
	return _IRegistration.Contract.EnableRegistration(&_IRegistration.TransactOpts, leagueAddr, raceId)
}

// EnableRegistration is a paid mutator transaction binding the contract method 0x52f3dcfb.
//
// Solidity: function enableRegistration(address leagueAddr, bytes32 raceId) returns()
func (_IRegistration *IRegistrationTransactorSession) EnableRegistration(leagueAddr common.Address, raceId [32]byte) (*types.Transaction, error) {
	return _IRegistration.Contract.EnableRegistration(&_IRegistration.TransactOpts, leagueAddr, raceId)
}

// ReceiveReward is a paid mutator transaction binding the contract method 0x1a0fc715.
//
// Solidity: function receiveReward(bytes32 raceId, uint256 slotId) returns()
func (_IRegistration *IRegistrationTransactor) ReceiveReward(opts *bind.TransactOpts, raceId [32]byte, slotId *big.Int) (*types.Transaction, error) {
	return _IRegistration.contract.Transact(opts, "receiveReward", raceId, slotId)
}

// ReceiveReward is a paid mutator transaction binding the contract method 0x1a0fc715.
//
// Solidity: function receiveReward(bytes32 raceId, uint256 slotId) returns()
func (_IRegistration *IRegistrationSession) ReceiveReward(raceId [32]byte, slotId *big.Int) (*types.Transaction, error) {
	return _IRegistration.Contract.ReceiveReward(&_IRegistration.TransactOpts, raceId, slotId)
}

// ReceiveReward is a paid mutator transaction binding the contract method 0x1a0fc715.
//
// Solidity: function receiveReward(bytes32 raceId, uint256 slotId) returns()
func (_IRegistration *IRegistrationTransactorSession) ReceiveReward(raceId [32]byte, slotId *big.Int) (*types.Transaction, error) {
	return _IRegistration.Contract.ReceiveReward(&_IRegistration.TransactOpts, raceId, slotId)
}

// Register is a paid mutator transaction binding the contract method 0x2965d809.
//
// Solidity: function register(uint256 slotId, bytes32 raceId) returns()
func (_IRegistration *IRegistrationTransactor) Register(opts *bind.TransactOpts, slotId *big.Int, raceId [32]byte) (*types.Transaction, error) {
	return _IRegistration.contract.Transact(opts, "register", slotId, raceId)
}

// Register is a paid mutator transaction binding the contract method 0x2965d809.
//
// Solidity: function register(uint256 slotId, bytes32 raceId) returns()
func (_IRegistration *IRegistrationSession) Register(slotId *big.Int, raceId [32]byte) (*types.Transaction, error) {
	return _IRegistration.Contract.Register(&_IRegistration.TransactOpts, slotId, raceId)
}

// Register is a paid mutator transaction binding the contract method 0x2965d809.
//
// Solidity: function register(uint256 slotId, bytes32 raceId) returns()
func (_IRegistration *IRegistrationTransactorSession) Register(slotId *big.Int, raceId [32]byte) (*types.Transaction, error) {
	return _IRegistration.Contract.Register(&_IRegistration.TransactOpts, slotId, raceId)
}

// RemoveReward is a paid mutator transaction binding the contract method 0x0fa445f5.
//
// Solidity: function removeReward(bytes32 raceId, uint256 nftRewardId, bytes1 resultIndex) returns()
func (_IRegistration *IRegistrationTransactor) RemoveReward(opts *bind.TransactOpts, raceId [32]byte, nftRewardId *big.Int, resultIndex [1]byte) (*types.Transaction, error) {
	return _IRegistration.contract.Transact(opts, "removeReward", raceId, nftRewardId, resultIndex)
}

// RemoveReward is a paid mutator transaction binding the contract method 0x0fa445f5.
//
// Solidity: function removeReward(bytes32 raceId, uint256 nftRewardId, bytes1 resultIndex) returns()
func (_IRegistration *IRegistrationSession) RemoveReward(raceId [32]byte, nftRewardId *big.Int, resultIndex [1]byte) (*types.Transaction, error) {
	return _IRegistration.Contract.RemoveReward(&_IRegistration.TransactOpts, raceId, nftRewardId, resultIndex)
}

// RemoveReward is a paid mutator transaction binding the contract method 0x0fa445f5.
//
// Solidity: function removeReward(bytes32 raceId, uint256 nftRewardId, bytes1 resultIndex) returns()
func (_IRegistration *IRegistrationTransactorSession) RemoveReward(raceId [32]byte, nftRewardId *big.Int, resultIndex [1]byte) (*types.Transaction, error) {
	return _IRegistration.Contract.RemoveReward(&_IRegistration.TransactOpts, raceId, nftRewardId, resultIndex)
}

// SelectParticipant is a paid mutator transaction binding the contract method 0x7d0f79b7.
//
// Solidity: function selectParticipant(bytes32 raceId) returns(uint256 requestId)
func (_IRegistration *IRegistrationTransactor) SelectParticipant(opts *bind.TransactOpts, raceId [32]byte) (*types.Transaction, error) {
	return _IRegistration.contract.Transact(opts, "selectParticipant", raceId)
}

// SelectParticipant is a paid mutator transaction binding the contract method 0x7d0f79b7.
//
// Solidity: function selectParticipant(bytes32 raceId) returns(uint256 requestId)
func (_IRegistration *IRegistrationSession) SelectParticipant(raceId [32]byte) (*types.Transaction, error) {
	return _IRegistration.Contract.SelectParticipant(&_IRegistration.TransactOpts, raceId)
}

// SelectParticipant is a paid mutator transaction binding the contract method 0x7d0f79b7.
//
// Solidity: function selectParticipant(bytes32 raceId) returns(uint256 requestId)
func (_IRegistration *IRegistrationTransactorSession) SelectParticipant(raceId [32]byte) (*types.Transaction, error) {
	return _IRegistration.Contract.SelectParticipant(&_IRegistration.TransactOpts, raceId)
}

// IRegistrationParticipantsSelectedIterator is returned from FilterParticipantsSelected and is used to iterate over the raw logs and unpacked data for ParticipantsSelected events raised by the IRegistration contract.
type IRegistrationParticipantsSelectedIterator struct {
	Event *IRegistrationParticipantsSelected // Event containing the contract specifics and raw log

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
func (it *IRegistrationParticipantsSelectedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IRegistrationParticipantsSelected)
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
		it.Event = new(IRegistrationParticipantsSelected)
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
func (it *IRegistrationParticipantsSelectedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IRegistrationParticipantsSelectedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IRegistrationParticipantsSelected represents a ParticipantsSelected event raised by the IRegistration contract.
type IRegistrationParticipantsSelected struct {
	RequestId  *big.Int
	RaceId     [32]byte
	Randomness *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterParticipantsSelected is a free log retrieval operation binding the contract event 0x5a1c9864d9d35edf19fb83e7a07ad46d28d80307bd248fb165e8814a1d1cfe2b.
//
// Solidity: event ParticipantsSelected(uint256 requestId, bytes32 raceId, uint256 randomness)
func (_IRegistration *IRegistrationFilterer) FilterParticipantsSelected(opts *bind.FilterOpts) (*IRegistrationParticipantsSelectedIterator, error) {

	logs, sub, err := _IRegistration.contract.FilterLogs(opts, "ParticipantsSelected")
	if err != nil {
		return nil, err
	}
	return &IRegistrationParticipantsSelectedIterator{contract: _IRegistration.contract, event: "ParticipantsSelected", logs: logs, sub: sub}, nil
}

// WatchParticipantsSelected is a free log subscription operation binding the contract event 0x5a1c9864d9d35edf19fb83e7a07ad46d28d80307bd248fb165e8814a1d1cfe2b.
//
// Solidity: event ParticipantsSelected(uint256 requestId, bytes32 raceId, uint256 randomness)
func (_IRegistration *IRegistrationFilterer) WatchParticipantsSelected(opts *bind.WatchOpts, sink chan<- *IRegistrationParticipantsSelected) (event.Subscription, error) {

	logs, sub, err := _IRegistration.contract.WatchLogs(opts, "ParticipantsSelected")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IRegistrationParticipantsSelected)
				if err := _IRegistration.contract.UnpackLog(event, "ParticipantsSelected", log); err != nil {
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
func (_IRegistration *IRegistrationFilterer) ParseParticipantsSelected(log types.Log) (*IRegistrationParticipantsSelected, error) {
	event := new(IRegistrationParticipantsSelected)
	if err := _IRegistration.contract.UnpackLog(event, "ParticipantsSelected", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IRegistrationRandomInProgressIterator is returned from FilterRandomInProgress and is used to iterate over the raw logs and unpacked data for RandomInProgress events raised by the IRegistration contract.
type IRegistrationRandomInProgressIterator struct {
	Event *IRegistrationRandomInProgress // Event containing the contract specifics and raw log

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
func (it *IRegistrationRandomInProgressIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IRegistrationRandomInProgress)
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
		it.Event = new(IRegistrationRandomInProgress)
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
func (it *IRegistrationRandomInProgressIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IRegistrationRandomInProgressIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IRegistrationRandomInProgress represents a RandomInProgress event raised by the IRegistration contract.
type IRegistrationRandomInProgress struct {
	RaceId [32]byte
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRandomInProgress is a free log retrieval operation binding the contract event 0x018dc112e9dad7765a03056b4091d237454ee37ae27aab3c3abd16d428900a13.
//
// Solidity: event RandomInProgress(bytes32 raceId)
func (_IRegistration *IRegistrationFilterer) FilterRandomInProgress(opts *bind.FilterOpts) (*IRegistrationRandomInProgressIterator, error) {

	logs, sub, err := _IRegistration.contract.FilterLogs(opts, "RandomInProgress")
	if err != nil {
		return nil, err
	}
	return &IRegistrationRandomInProgressIterator{contract: _IRegistration.contract, event: "RandomInProgress", logs: logs, sub: sub}, nil
}

// WatchRandomInProgress is a free log subscription operation binding the contract event 0x018dc112e9dad7765a03056b4091d237454ee37ae27aab3c3abd16d428900a13.
//
// Solidity: event RandomInProgress(bytes32 raceId)
func (_IRegistration *IRegistrationFilterer) WatchRandomInProgress(opts *bind.WatchOpts, sink chan<- *IRegistrationRandomInProgress) (event.Subscription, error) {

	logs, sub, err := _IRegistration.contract.WatchLogs(opts, "RandomInProgress")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IRegistrationRandomInProgress)
				if err := _IRegistration.contract.UnpackLog(event, "RandomInProgress", log); err != nil {
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
func (_IRegistration *IRegistrationFilterer) ParseRandomInProgress(log types.Log) (*IRegistrationRandomInProgress, error) {
	event := new(IRegistrationRandomInProgress)
	if err := _IRegistration.contract.UnpackLog(event, "RandomInProgress", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IRegistrationRegisteredIterator is returned from FilterRegistered and is used to iterate over the raw logs and unpacked data for Registered events raised by the IRegistration contract.
type IRegistrationRegisteredIterator struct {
	Event *IRegistrationRegistered // Event containing the contract specifics and raw log

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
func (it *IRegistrationRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IRegistrationRegistered)
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
		it.Event = new(IRegistrationRegistered)
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
func (it *IRegistrationRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IRegistrationRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IRegistrationRegistered represents a Registered event raised by the IRegistration contract.
type IRegistrationRegistered struct {
	SlotId      *big.Int
	Participant common.Address
	RaceId      [32]byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterRegistered is a free log retrieval operation binding the contract event 0xd62c7bd2779dc6ebe426bed623399b1fc1d190c063eadefcffffb1efe6f253ef.
//
// Solidity: event Registered(uint256 slotId, address participant, bytes32 raceId)
func (_IRegistration *IRegistrationFilterer) FilterRegistered(opts *bind.FilterOpts) (*IRegistrationRegisteredIterator, error) {

	logs, sub, err := _IRegistration.contract.FilterLogs(opts, "Registered")
	if err != nil {
		return nil, err
	}
	return &IRegistrationRegisteredIterator{contract: _IRegistration.contract, event: "Registered", logs: logs, sub: sub}, nil
}

// WatchRegistered is a free log subscription operation binding the contract event 0xd62c7bd2779dc6ebe426bed623399b1fc1d190c063eadefcffffb1efe6f253ef.
//
// Solidity: event Registered(uint256 slotId, address participant, bytes32 raceId)
func (_IRegistration *IRegistrationFilterer) WatchRegistered(opts *bind.WatchOpts, sink chan<- *IRegistrationRegistered) (event.Subscription, error) {

	logs, sub, err := _IRegistration.contract.WatchLogs(opts, "Registered")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IRegistrationRegistered)
				if err := _IRegistration.contract.UnpackLog(event, "Registered", log); err != nil {
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
func (_IRegistration *IRegistrationFilterer) ParseRegistered(log types.Log) (*IRegistrationRegistered, error) {
	event := new(IRegistrationRegistered)
	if err := _IRegistration.contract.UnpackLog(event, "Registered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IRegistrationRewardAddedIterator is returned from FilterRewardAdded and is used to iterate over the raw logs and unpacked data for RewardAdded events raised by the IRegistration contract.
type IRegistrationRewardAddedIterator struct {
	Event *IRegistrationRewardAdded // Event containing the contract specifics and raw log

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
func (it *IRegistrationRewardAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IRegistrationRewardAdded)
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
		it.Event = new(IRegistrationRewardAdded)
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
func (it *IRegistrationRewardAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IRegistrationRewardAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IRegistrationRewardAdded represents a RewardAdded event raised by the IRegistration contract.
type IRegistrationRewardAdded struct {
	RaceId      [32]byte
	NftRewardId *big.Int
	ResultIndex [1]byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterRewardAdded is a free log retrieval operation binding the contract event 0x2b942e265bd6d4ca5617a327f9a60e1a85f4351aaa0be2a84cda1cf9d32c8c04.
//
// Solidity: event RewardAdded(bytes32 raceId, uint256 nftRewardId, bytes1 resultIndex)
func (_IRegistration *IRegistrationFilterer) FilterRewardAdded(opts *bind.FilterOpts) (*IRegistrationRewardAddedIterator, error) {

	logs, sub, err := _IRegistration.contract.FilterLogs(opts, "RewardAdded")
	if err != nil {
		return nil, err
	}
	return &IRegistrationRewardAddedIterator{contract: _IRegistration.contract, event: "RewardAdded", logs: logs, sub: sub}, nil
}

// WatchRewardAdded is a free log subscription operation binding the contract event 0x2b942e265bd6d4ca5617a327f9a60e1a85f4351aaa0be2a84cda1cf9d32c8c04.
//
// Solidity: event RewardAdded(bytes32 raceId, uint256 nftRewardId, bytes1 resultIndex)
func (_IRegistration *IRegistrationFilterer) WatchRewardAdded(opts *bind.WatchOpts, sink chan<- *IRegistrationRewardAdded) (event.Subscription, error) {

	logs, sub, err := _IRegistration.contract.WatchLogs(opts, "RewardAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IRegistrationRewardAdded)
				if err := _IRegistration.contract.UnpackLog(event, "RewardAdded", log); err != nil {
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
func (_IRegistration *IRegistrationFilterer) ParseRewardAdded(log types.Log) (*IRegistrationRewardAdded, error) {
	event := new(IRegistrationRewardAdded)
	if err := _IRegistration.contract.UnpackLog(event, "RewardAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IRegistrationRewardReceivedIterator is returned from FilterRewardReceived and is used to iterate over the raw logs and unpacked data for RewardReceived events raised by the IRegistration contract.
type IRegistrationRewardReceivedIterator struct {
	Event *IRegistrationRewardReceived // Event containing the contract specifics and raw log

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
func (it *IRegistrationRewardReceivedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IRegistrationRewardReceived)
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
		it.Event = new(IRegistrationRewardReceived)
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
func (it *IRegistrationRewardReceivedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IRegistrationRewardReceivedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IRegistrationRewardReceived represents a RewardReceived event raised by the IRegistration contract.
type IRegistrationRewardReceived struct {
	RaceId      [32]byte
	SlotId      *big.Int
	NftRewardId *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterRewardReceived is a free log retrieval operation binding the contract event 0xe2ece170b35ae79cbe45ffdaf4931b8bc45056da5d3e9aa4aea67fcdf307929a.
//
// Solidity: event RewardReceived(bytes32 raceId, uint256 slotId, uint256 nftRewardId)
func (_IRegistration *IRegistrationFilterer) FilterRewardReceived(opts *bind.FilterOpts) (*IRegistrationRewardReceivedIterator, error) {

	logs, sub, err := _IRegistration.contract.FilterLogs(opts, "RewardReceived")
	if err != nil {
		return nil, err
	}
	return &IRegistrationRewardReceivedIterator{contract: _IRegistration.contract, event: "RewardReceived", logs: logs, sub: sub}, nil
}

// WatchRewardReceived is a free log subscription operation binding the contract event 0xe2ece170b35ae79cbe45ffdaf4931b8bc45056da5d3e9aa4aea67fcdf307929a.
//
// Solidity: event RewardReceived(bytes32 raceId, uint256 slotId, uint256 nftRewardId)
func (_IRegistration *IRegistrationFilterer) WatchRewardReceived(opts *bind.WatchOpts, sink chan<- *IRegistrationRewardReceived) (event.Subscription, error) {

	logs, sub, err := _IRegistration.contract.WatchLogs(opts, "RewardReceived")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IRegistrationRewardReceived)
				if err := _IRegistration.contract.UnpackLog(event, "RewardReceived", log); err != nil {
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
func (_IRegistration *IRegistrationFilterer) ParseRewardReceived(log types.Log) (*IRegistrationRewardReceived, error) {
	event := new(IRegistrationRewardReceived)
	if err := _IRegistration.contract.UnpackLog(event, "RewardReceived", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IRegistrationRewardRemovedIterator is returned from FilterRewardRemoved and is used to iterate over the raw logs and unpacked data for RewardRemoved events raised by the IRegistration contract.
type IRegistrationRewardRemovedIterator struct {
	Event *IRegistrationRewardRemoved // Event containing the contract specifics and raw log

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
func (it *IRegistrationRewardRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IRegistrationRewardRemoved)
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
		it.Event = new(IRegistrationRewardRemoved)
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
func (it *IRegistrationRewardRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IRegistrationRewardRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IRegistrationRewardRemoved represents a RewardRemoved event raised by the IRegistration contract.
type IRegistrationRewardRemoved struct {
	RaceId      [32]byte
	NftRewardId *big.Int
	ResultIndex [1]byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterRewardRemoved is a free log retrieval operation binding the contract event 0xddda89b96dee97fe43f9803c50584f6dd667bedc8b02a7c75407e6906bf31ead.
//
// Solidity: event RewardRemoved(bytes32 raceId, uint256 nftRewardId, bytes1 resultIndex)
func (_IRegistration *IRegistrationFilterer) FilterRewardRemoved(opts *bind.FilterOpts) (*IRegistrationRewardRemovedIterator, error) {

	logs, sub, err := _IRegistration.contract.FilterLogs(opts, "RewardRemoved")
	if err != nil {
		return nil, err
	}
	return &IRegistrationRewardRemovedIterator{contract: _IRegistration.contract, event: "RewardRemoved", logs: logs, sub: sub}, nil
}

// WatchRewardRemoved is a free log subscription operation binding the contract event 0xddda89b96dee97fe43f9803c50584f6dd667bedc8b02a7c75407e6906bf31ead.
//
// Solidity: event RewardRemoved(bytes32 raceId, uint256 nftRewardId, bytes1 resultIndex)
func (_IRegistration *IRegistrationFilterer) WatchRewardRemoved(opts *bind.WatchOpts, sink chan<- *IRegistrationRewardRemoved) (event.Subscription, error) {

	logs, sub, err := _IRegistration.contract.WatchLogs(opts, "RewardRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IRegistrationRewardRemoved)
				if err := _IRegistration.contract.UnpackLog(event, "RewardRemoved", log); err != nil {
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
func (_IRegistration *IRegistrationFilterer) ParseRewardRemoved(log types.Log) (*IRegistrationRewardRemoved, error) {
	event := new(IRegistrationRewardRemoved)
	if err := _IRegistration.contract.UnpackLog(event, "RewardRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
