// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package BlackList

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

// BlackListMetaData contains all meta data concerning the BlackList contract.
var BlackListMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"AddedBlackList\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"RemovedBlackList\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"isBlackListed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_maker\",\"type\":\"address\"}],\"name\":\"getBlackListStatus\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_evilUser\",\"type\":\"address\"}],\"name\":\"addBlackList\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_clearedUser\",\"type\":\"address\"}],\"name\":\"removeBlackList\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b5061002d61002261003260201b60201c565b61003a60201b60201c565b6100fe565b600033905090565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050816000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b6108908061010d6000396000f3fe608060405234801561001057600080fd5b506004361061007d5760003560e01c80638da5cb5b1161005b5780638da5cb5b146100d8578063e47d6060146100f6578063e4997dc514610126578063f2fde38b146101425761007d565b80630ecb93c01461008257806359bf1abe1461009e578063715018a6146100ce575b600080fd5b61009c60048036038101906100979190610679565b61015e565b005b6100b860048036038101906100b39190610679565b61026b565b6040516100c59190610721565b60405180910390f35b6100d66102c1565b005b6100e0610349565b6040516100ed9190610706565b60405180910390f35b610110600480360381019061010b9190610679565b610372565b60405161011d9190610721565b60405180910390f35b610140600480360381019061013b9190610679565b610392565b005b61015c60048036038101906101579190610679565b6104a0565b005b610166610598565b73ffffffffffffffffffffffffffffffffffffffff16610184610349565b73ffffffffffffffffffffffffffffffffffffffff16146101da576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016101d19061075c565b60405180910390fd5b60018060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055507f42e160154868087d6bfdc0ca23d96a1c1cfa32f1b72ba9ba27b69b98a0d819dc816040516102609190610706565b60405180910390a150565b6000600160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff169050919050565b6102c9610598565b73ffffffffffffffffffffffffffffffffffffffff166102e7610349565b73ffffffffffffffffffffffffffffffffffffffff161461033d576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016103349061075c565b60405180910390fd5b61034760006105a0565b565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b60016020528060005260406000206000915054906101000a900460ff1681565b61039a610598565b73ffffffffffffffffffffffffffffffffffffffff166103b8610349565b73ffffffffffffffffffffffffffffffffffffffff161461040e576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016104059061075c565b60405180910390fd5b6000600160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055507fd7e9ec6e6ecd65492dce6bf513cd6867560d49544421d0783ddf06e76c24470c816040516104959190610706565b60405180910390a150565b6104a8610598565b73ffffffffffffffffffffffffffffffffffffffff166104c6610349565b73ffffffffffffffffffffffffffffffffffffffff161461051c576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016105139061075c565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16141561058c576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016105839061073c565b60405180910390fd5b610595816105a0565b50565b600033905090565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050816000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b60008135905061067381610843565b92915050565b60006020828403121561068b57600080fd5b600061069984828501610664565b91505092915050565b6106ab8161078d565b82525050565b6106ba8161079f565b82525050565b60006106cd60268361077c565b91506106d8826107cb565b604082019050919050565b60006106f060208361077c565b91506106fb8261081a565b602082019050919050565b600060208201905061071b60008301846106a2565b92915050565b600060208201905061073660008301846106b1565b92915050565b60006020820190508181036000830152610755816106c0565b9050919050565b60006020820190508181036000830152610775816106e3565b9050919050565b600082825260208201905092915050565b6000610798826107ab565b9050919050565b60008115159050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b7f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160008201527f6464726573730000000000000000000000000000000000000000000000000000602082015250565b7f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572600082015250565b61084c8161078d565b811461085757600080fd5b5056fea26469706673582212200d3bdec767e76c4a6a0195c66c6f6b2bfb8273e9d71162255f0c3dee580a3be564736f6c63430008040033",
}

// BlackListABI is the input ABI used to generate the binding from.
// Deprecated: Use BlackListMetaData.ABI instead.
var BlackListABI = BlackListMetaData.ABI

// BlackListBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use BlackListMetaData.Bin instead.
var BlackListBin = BlackListMetaData.Bin

// DeployBlackList deploys a new Ethereum contract, binding an instance of BlackList to it.
func DeployBlackList(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *BlackList, error) {
	parsed, err := BlackListMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(BlackListBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &BlackList{BlackListCaller: BlackListCaller{contract: contract}, BlackListTransactor: BlackListTransactor{contract: contract}, BlackListFilterer: BlackListFilterer{contract: contract}}, nil
}

// BlackList is an auto generated Go binding around an Ethereum contract.
type BlackList struct {
	BlackListCaller     // Read-only binding to the contract
	BlackListTransactor // Write-only binding to the contract
	BlackListFilterer   // Log filterer for contract events
}

// BlackListCaller is an auto generated read-only Go binding around an Ethereum contract.
type BlackListCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BlackListTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BlackListTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BlackListFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BlackListFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BlackListSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BlackListSession struct {
	Contract     *BlackList        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BlackListCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BlackListCallerSession struct {
	Contract *BlackListCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// BlackListTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BlackListTransactorSession struct {
	Contract     *BlackListTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// BlackListRaw is an auto generated low-level Go binding around an Ethereum contract.
type BlackListRaw struct {
	Contract *BlackList // Generic contract binding to access the raw methods on
}

// BlackListCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BlackListCallerRaw struct {
	Contract *BlackListCaller // Generic read-only contract binding to access the raw methods on
}

// BlackListTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BlackListTransactorRaw struct {
	Contract *BlackListTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBlackList creates a new instance of BlackList, bound to a specific deployed contract.
func NewBlackList(address common.Address, backend bind.ContractBackend) (*BlackList, error) {
	contract, err := bindBlackList(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BlackList{BlackListCaller: BlackListCaller{contract: contract}, BlackListTransactor: BlackListTransactor{contract: contract}, BlackListFilterer: BlackListFilterer{contract: contract}}, nil
}

// NewBlackListCaller creates a new read-only instance of BlackList, bound to a specific deployed contract.
func NewBlackListCaller(address common.Address, caller bind.ContractCaller) (*BlackListCaller, error) {
	contract, err := bindBlackList(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BlackListCaller{contract: contract}, nil
}

// NewBlackListTransactor creates a new write-only instance of BlackList, bound to a specific deployed contract.
func NewBlackListTransactor(address common.Address, transactor bind.ContractTransactor) (*BlackListTransactor, error) {
	contract, err := bindBlackList(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BlackListTransactor{contract: contract}, nil
}

// NewBlackListFilterer creates a new log filterer instance of BlackList, bound to a specific deployed contract.
func NewBlackListFilterer(address common.Address, filterer bind.ContractFilterer) (*BlackListFilterer, error) {
	contract, err := bindBlackList(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BlackListFilterer{contract: contract}, nil
}

// bindBlackList binds a generic wrapper to an already deployed contract.
func bindBlackList(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BlackListABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BlackList *BlackListRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BlackList.Contract.BlackListCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BlackList *BlackListRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BlackList.Contract.BlackListTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BlackList *BlackListRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BlackList.Contract.BlackListTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BlackList *BlackListCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BlackList.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BlackList *BlackListTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BlackList.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BlackList *BlackListTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BlackList.Contract.contract.Transact(opts, method, params...)
}

// GetBlackListStatus is a free data retrieval call binding the contract method 0x59bf1abe.
//
// Solidity: function getBlackListStatus(address _maker) view returns(bool)
func (_BlackList *BlackListCaller) GetBlackListStatus(opts *bind.CallOpts, _maker common.Address) (bool, error) {
	var out []interface{}
	err := _BlackList.contract.Call(opts, &out, "getBlackListStatus", _maker)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// GetBlackListStatus is a free data retrieval call binding the contract method 0x59bf1abe.
//
// Solidity: function getBlackListStatus(address _maker) view returns(bool)
func (_BlackList *BlackListSession) GetBlackListStatus(_maker common.Address) (bool, error) {
	return _BlackList.Contract.GetBlackListStatus(&_BlackList.CallOpts, _maker)
}

// GetBlackListStatus is a free data retrieval call binding the contract method 0x59bf1abe.
//
// Solidity: function getBlackListStatus(address _maker) view returns(bool)
func (_BlackList *BlackListCallerSession) GetBlackListStatus(_maker common.Address) (bool, error) {
	return _BlackList.Contract.GetBlackListStatus(&_BlackList.CallOpts, _maker)
}

// IsBlackListed is a free data retrieval call binding the contract method 0xe47d6060.
//
// Solidity: function isBlackListed(address ) view returns(bool)
func (_BlackList *BlackListCaller) IsBlackListed(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _BlackList.contract.Call(opts, &out, "isBlackListed", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsBlackListed is a free data retrieval call binding the contract method 0xe47d6060.
//
// Solidity: function isBlackListed(address ) view returns(bool)
func (_BlackList *BlackListSession) IsBlackListed(arg0 common.Address) (bool, error) {
	return _BlackList.Contract.IsBlackListed(&_BlackList.CallOpts, arg0)
}

// IsBlackListed is a free data retrieval call binding the contract method 0xe47d6060.
//
// Solidity: function isBlackListed(address ) view returns(bool)
func (_BlackList *BlackListCallerSession) IsBlackListed(arg0 common.Address) (bool, error) {
	return _BlackList.Contract.IsBlackListed(&_BlackList.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BlackList *BlackListCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BlackList.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BlackList *BlackListSession) Owner() (common.Address, error) {
	return _BlackList.Contract.Owner(&_BlackList.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BlackList *BlackListCallerSession) Owner() (common.Address, error) {
	return _BlackList.Contract.Owner(&_BlackList.CallOpts)
}

// AddBlackList is a paid mutator transaction binding the contract method 0x0ecb93c0.
//
// Solidity: function addBlackList(address _evilUser) returns()
func (_BlackList *BlackListTransactor) AddBlackList(opts *bind.TransactOpts, _evilUser common.Address) (*types.Transaction, error) {
	return _BlackList.contract.Transact(opts, "addBlackList", _evilUser)
}

// AddBlackList is a paid mutator transaction binding the contract method 0x0ecb93c0.
//
// Solidity: function addBlackList(address _evilUser) returns()
func (_BlackList *BlackListSession) AddBlackList(_evilUser common.Address) (*types.Transaction, error) {
	return _BlackList.Contract.AddBlackList(&_BlackList.TransactOpts, _evilUser)
}

// AddBlackList is a paid mutator transaction binding the contract method 0x0ecb93c0.
//
// Solidity: function addBlackList(address _evilUser) returns()
func (_BlackList *BlackListTransactorSession) AddBlackList(_evilUser common.Address) (*types.Transaction, error) {
	return _BlackList.Contract.AddBlackList(&_BlackList.TransactOpts, _evilUser)
}

// RemoveBlackList is a paid mutator transaction binding the contract method 0xe4997dc5.
//
// Solidity: function removeBlackList(address _clearedUser) returns()
func (_BlackList *BlackListTransactor) RemoveBlackList(opts *bind.TransactOpts, _clearedUser common.Address) (*types.Transaction, error) {
	return _BlackList.contract.Transact(opts, "removeBlackList", _clearedUser)
}

// RemoveBlackList is a paid mutator transaction binding the contract method 0xe4997dc5.
//
// Solidity: function removeBlackList(address _clearedUser) returns()
func (_BlackList *BlackListSession) RemoveBlackList(_clearedUser common.Address) (*types.Transaction, error) {
	return _BlackList.Contract.RemoveBlackList(&_BlackList.TransactOpts, _clearedUser)
}

// RemoveBlackList is a paid mutator transaction binding the contract method 0xe4997dc5.
//
// Solidity: function removeBlackList(address _clearedUser) returns()
func (_BlackList *BlackListTransactorSession) RemoveBlackList(_clearedUser common.Address) (*types.Transaction, error) {
	return _BlackList.Contract.RemoveBlackList(&_BlackList.TransactOpts, _clearedUser)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_BlackList *BlackListTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BlackList.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_BlackList *BlackListSession) RenounceOwnership() (*types.Transaction, error) {
	return _BlackList.Contract.RenounceOwnership(&_BlackList.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_BlackList *BlackListTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _BlackList.Contract.RenounceOwnership(&_BlackList.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_BlackList *BlackListTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _BlackList.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_BlackList *BlackListSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _BlackList.Contract.TransferOwnership(&_BlackList.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_BlackList *BlackListTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _BlackList.Contract.TransferOwnership(&_BlackList.TransactOpts, newOwner)
}

// BlackListAddedBlackListIterator is returned from FilterAddedBlackList and is used to iterate over the raw logs and unpacked data for AddedBlackList events raised by the BlackList contract.
type BlackListAddedBlackListIterator struct {
	Event *BlackListAddedBlackList // Event containing the contract specifics and raw log

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
func (it *BlackListAddedBlackListIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BlackListAddedBlackList)
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
		it.Event = new(BlackListAddedBlackList)
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
func (it *BlackListAddedBlackListIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BlackListAddedBlackListIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BlackListAddedBlackList represents a AddedBlackList event raised by the BlackList contract.
type BlackListAddedBlackList struct {
	User common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterAddedBlackList is a free log retrieval operation binding the contract event 0x42e160154868087d6bfdc0ca23d96a1c1cfa32f1b72ba9ba27b69b98a0d819dc.
//
// Solidity: event AddedBlackList(address _user)
func (_BlackList *BlackListFilterer) FilterAddedBlackList(opts *bind.FilterOpts) (*BlackListAddedBlackListIterator, error) {

	logs, sub, err := _BlackList.contract.FilterLogs(opts, "AddedBlackList")
	if err != nil {
		return nil, err
	}
	return &BlackListAddedBlackListIterator{contract: _BlackList.contract, event: "AddedBlackList", logs: logs, sub: sub}, nil
}

// WatchAddedBlackList is a free log subscription operation binding the contract event 0x42e160154868087d6bfdc0ca23d96a1c1cfa32f1b72ba9ba27b69b98a0d819dc.
//
// Solidity: event AddedBlackList(address _user)
func (_BlackList *BlackListFilterer) WatchAddedBlackList(opts *bind.WatchOpts, sink chan<- *BlackListAddedBlackList) (event.Subscription, error) {

	logs, sub, err := _BlackList.contract.WatchLogs(opts, "AddedBlackList")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BlackListAddedBlackList)
				if err := _BlackList.contract.UnpackLog(event, "AddedBlackList", log); err != nil {
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

// ParseAddedBlackList is a log parse operation binding the contract event 0x42e160154868087d6bfdc0ca23d96a1c1cfa32f1b72ba9ba27b69b98a0d819dc.
//
// Solidity: event AddedBlackList(address _user)
func (_BlackList *BlackListFilterer) ParseAddedBlackList(log types.Log) (*BlackListAddedBlackList, error) {
	event := new(BlackListAddedBlackList)
	if err := _BlackList.contract.UnpackLog(event, "AddedBlackList", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BlackListOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the BlackList contract.
type BlackListOwnershipTransferredIterator struct {
	Event *BlackListOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *BlackListOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BlackListOwnershipTransferred)
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
		it.Event = new(BlackListOwnershipTransferred)
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
func (it *BlackListOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BlackListOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BlackListOwnershipTransferred represents a OwnershipTransferred event raised by the BlackList contract.
type BlackListOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_BlackList *BlackListFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*BlackListOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _BlackList.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &BlackListOwnershipTransferredIterator{contract: _BlackList.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_BlackList *BlackListFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *BlackListOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _BlackList.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BlackListOwnershipTransferred)
				if err := _BlackList.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_BlackList *BlackListFilterer) ParseOwnershipTransferred(log types.Log) (*BlackListOwnershipTransferred, error) {
	event := new(BlackListOwnershipTransferred)
	if err := _BlackList.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BlackListRemovedBlackListIterator is returned from FilterRemovedBlackList and is used to iterate over the raw logs and unpacked data for RemovedBlackList events raised by the BlackList contract.
type BlackListRemovedBlackListIterator struct {
	Event *BlackListRemovedBlackList // Event containing the contract specifics and raw log

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
func (it *BlackListRemovedBlackListIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BlackListRemovedBlackList)
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
		it.Event = new(BlackListRemovedBlackList)
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
func (it *BlackListRemovedBlackListIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BlackListRemovedBlackListIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BlackListRemovedBlackList represents a RemovedBlackList event raised by the BlackList contract.
type BlackListRemovedBlackList struct {
	User common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterRemovedBlackList is a free log retrieval operation binding the contract event 0xd7e9ec6e6ecd65492dce6bf513cd6867560d49544421d0783ddf06e76c24470c.
//
// Solidity: event RemovedBlackList(address _user)
func (_BlackList *BlackListFilterer) FilterRemovedBlackList(opts *bind.FilterOpts) (*BlackListRemovedBlackListIterator, error) {

	logs, sub, err := _BlackList.contract.FilterLogs(opts, "RemovedBlackList")
	if err != nil {
		return nil, err
	}
	return &BlackListRemovedBlackListIterator{contract: _BlackList.contract, event: "RemovedBlackList", logs: logs, sub: sub}, nil
}

// WatchRemovedBlackList is a free log subscription operation binding the contract event 0xd7e9ec6e6ecd65492dce6bf513cd6867560d49544421d0783ddf06e76c24470c.
//
// Solidity: event RemovedBlackList(address _user)
func (_BlackList *BlackListFilterer) WatchRemovedBlackList(opts *bind.WatchOpts, sink chan<- *BlackListRemovedBlackList) (event.Subscription, error) {

	logs, sub, err := _BlackList.contract.WatchLogs(opts, "RemovedBlackList")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BlackListRemovedBlackList)
				if err := _BlackList.contract.UnpackLog(event, "RemovedBlackList", log); err != nil {
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

// ParseRemovedBlackList is a log parse operation binding the contract event 0xd7e9ec6e6ecd65492dce6bf513cd6867560d49544421d0783ddf06e76c24470c.
//
// Solidity: event RemovedBlackList(address _user)
func (_BlackList *BlackListFilterer) ParseRemovedBlackList(log types.Log) (*BlackListRemovedBlackList, error) {
	event := new(BlackListRemovedBlackList)
	if err := _BlackList.contract.UnpackLog(event, "RemovedBlackList", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
