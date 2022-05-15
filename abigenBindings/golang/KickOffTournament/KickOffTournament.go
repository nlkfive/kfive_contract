// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package KickOffTournament

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

// KickOffTournamentMetaData contains all meta data concerning the KickOffTournament contract.
var KickOffTournamentMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"noRace\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"nlggtAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"raceReward\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_tournamentName\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AlreadyRegistered\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_currentRaceNo\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalRace\",\"type\":\"uint256\"}],\"name\":\"CannotCreateMoreRace\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidContract\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidRegister\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidSender\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidSlot\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotEndYet\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RaceExisted\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RaceNotExisted\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RaceWasUpdated\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"}],\"name\":\"TooEarly\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"}],\"name\":\"TooLate\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"}],\"name\":\"RaceCancelled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"noSlot\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"raceNo\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"startAt\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"}],\"name\":\"RaceCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes27\",\"name\":\"result\",\"type\":\"bytes27\"}],\"name\":\"RaceResultUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"slotId\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"participant\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"raceId\",\"type\":\"bytes32\"}],\"name\":\"Registered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"winnerIndex\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"score\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"winner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nftRewardId\",\"type\":\"uint256\"}],\"name\":\"RewardGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PAUSER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"raceId\",\"type\":\"bytes32\"}],\"name\":\"cancelRace\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"noSlot\",\"type\":\"uint8\"},{\"internalType\":\"uint32\",\"name\":\"startAt\",\"type\":\"uint32\"}],\"name\":\"createRace\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getRoleMember\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleMemberCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"register\",\"type\":\"address\"}],\"name\":\"getTotalScore\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"score\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"winnerIndex\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"winner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nftRewardId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"tokenURI\",\"type\":\"string\"}],\"name\":\"grandReward\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC721Received\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"raceId\",\"type\":\"bytes32\"}],\"name\":\"raceInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"uint8\",\"name\":\"noSlot\",\"type\":\"uint8\"},{\"internalType\":\"uint32\",\"name\":\"startAt\",\"type\":\"uint32\"},{\"internalType\":\"bytes27\",\"name\":\"result\",\"type\":\"bytes27\"}],\"internalType\":\"structITournament.Race\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"slotId\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"raceId\",\"type\":\"bytes32\"}],\"name\":\"registerRace\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"register\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"raceId\",\"type\":\"bytes32\"}],\"name\":\"registeredSlot\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tournamentInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"uint8\",\"name\":\"totalRace\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"createdRace\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"endedRace\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"tournamentName\",\"type\":\"string\"}],\"internalType\":\"structITournament.TournamentInfo\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"raceId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes27\",\"name\":\"result\",\"type\":\"bytes27\"}],\"name\":\"updateRaceResult\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// KickOffTournamentABI is the input ABI used to generate the binding from.
// Deprecated: Use KickOffTournamentMetaData.ABI instead.
var KickOffTournamentABI = KickOffTournamentMetaData.ABI

// KickOffTournament is an auto generated Go binding around an Ethereum contract.
type KickOffTournament struct {
	KickOffTournamentCaller     // Read-only binding to the contract
	KickOffTournamentTransactor // Write-only binding to the contract
	KickOffTournamentFilterer   // Log filterer for contract events
}

// KickOffTournamentCaller is an auto generated read-only Go binding around an Ethereum contract.
type KickOffTournamentCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KickOffTournamentTransactor is an auto generated write-only Go binding around an Ethereum contract.
type KickOffTournamentTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KickOffTournamentFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type KickOffTournamentFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KickOffTournamentSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type KickOffTournamentSession struct {
	Contract     *KickOffTournament // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// KickOffTournamentCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type KickOffTournamentCallerSession struct {
	Contract *KickOffTournamentCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// KickOffTournamentTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type KickOffTournamentTransactorSession struct {
	Contract     *KickOffTournamentTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// KickOffTournamentRaw is an auto generated low-level Go binding around an Ethereum contract.
type KickOffTournamentRaw struct {
	Contract *KickOffTournament // Generic contract binding to access the raw methods on
}

// KickOffTournamentCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type KickOffTournamentCallerRaw struct {
	Contract *KickOffTournamentCaller // Generic read-only contract binding to access the raw methods on
}

// KickOffTournamentTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type KickOffTournamentTransactorRaw struct {
	Contract *KickOffTournamentTransactor // Generic write-only contract binding to access the raw methods on
}

// NewKickOffTournament creates a new instance of KickOffTournament, bound to a specific deployed contract.
func NewKickOffTournament(address common.Address, backend bind.ContractBackend) (*KickOffTournament, error) {
	contract, err := bindKickOffTournament(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &KickOffTournament{KickOffTournamentCaller: KickOffTournamentCaller{contract: contract}, KickOffTournamentTransactor: KickOffTournamentTransactor{contract: contract}, KickOffTournamentFilterer: KickOffTournamentFilterer{contract: contract}}, nil
}

// NewKickOffTournamentCaller creates a new read-only instance of KickOffTournament, bound to a specific deployed contract.
func NewKickOffTournamentCaller(address common.Address, caller bind.ContractCaller) (*KickOffTournamentCaller, error) {
	contract, err := bindKickOffTournament(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &KickOffTournamentCaller{contract: contract}, nil
}

// NewKickOffTournamentTransactor creates a new write-only instance of KickOffTournament, bound to a specific deployed contract.
func NewKickOffTournamentTransactor(address common.Address, transactor bind.ContractTransactor) (*KickOffTournamentTransactor, error) {
	contract, err := bindKickOffTournament(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &KickOffTournamentTransactor{contract: contract}, nil
}

// NewKickOffTournamentFilterer creates a new log filterer instance of KickOffTournament, bound to a specific deployed contract.
func NewKickOffTournamentFilterer(address common.Address, filterer bind.ContractFilterer) (*KickOffTournamentFilterer, error) {
	contract, err := bindKickOffTournament(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &KickOffTournamentFilterer{contract: contract}, nil
}

// bindKickOffTournament binds a generic wrapper to an already deployed contract.
func bindKickOffTournament(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(KickOffTournamentABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_KickOffTournament *KickOffTournamentRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _KickOffTournament.Contract.KickOffTournamentCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_KickOffTournament *KickOffTournamentRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _KickOffTournament.Contract.KickOffTournamentTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_KickOffTournament *KickOffTournamentRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _KickOffTournament.Contract.KickOffTournamentTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_KickOffTournament *KickOffTournamentCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _KickOffTournament.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_KickOffTournament *KickOffTournamentTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _KickOffTournament.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_KickOffTournament *KickOffTournamentTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _KickOffTournament.Contract.contract.Transact(opts, method, params...)
}

// ADMINROLE is a free data retrieval call binding the contract method 0x75b238fc.
//
// Solidity: function ADMIN_ROLE() view returns(bytes32)
func (_KickOffTournament *KickOffTournamentCaller) ADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _KickOffTournament.contract.Call(opts, &out, "ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ADMINROLE is a free data retrieval call binding the contract method 0x75b238fc.
//
// Solidity: function ADMIN_ROLE() view returns(bytes32)
func (_KickOffTournament *KickOffTournamentSession) ADMINROLE() ([32]byte, error) {
	return _KickOffTournament.Contract.ADMINROLE(&_KickOffTournament.CallOpts)
}

// ADMINROLE is a free data retrieval call binding the contract method 0x75b238fc.
//
// Solidity: function ADMIN_ROLE() view returns(bytes32)
func (_KickOffTournament *KickOffTournamentCallerSession) ADMINROLE() ([32]byte, error) {
	return _KickOffTournament.Contract.ADMINROLE(&_KickOffTournament.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_KickOffTournament *KickOffTournamentCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _KickOffTournament.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_KickOffTournament *KickOffTournamentSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _KickOffTournament.Contract.DEFAULTADMINROLE(&_KickOffTournament.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_KickOffTournament *KickOffTournamentCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _KickOffTournament.Contract.DEFAULTADMINROLE(&_KickOffTournament.CallOpts)
}

// PAUSERROLE is a free data retrieval call binding the contract method 0xe63ab1e9.
//
// Solidity: function PAUSER_ROLE() view returns(bytes32)
func (_KickOffTournament *KickOffTournamentCaller) PAUSERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _KickOffTournament.contract.Call(opts, &out, "PAUSER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// PAUSERROLE is a free data retrieval call binding the contract method 0xe63ab1e9.
//
// Solidity: function PAUSER_ROLE() view returns(bytes32)
func (_KickOffTournament *KickOffTournamentSession) PAUSERROLE() ([32]byte, error) {
	return _KickOffTournament.Contract.PAUSERROLE(&_KickOffTournament.CallOpts)
}

// PAUSERROLE is a free data retrieval call binding the contract method 0xe63ab1e9.
//
// Solidity: function PAUSER_ROLE() view returns(bytes32)
func (_KickOffTournament *KickOffTournamentCallerSession) PAUSERROLE() ([32]byte, error) {
	return _KickOffTournament.Contract.PAUSERROLE(&_KickOffTournament.CallOpts)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_KickOffTournament *KickOffTournamentCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _KickOffTournament.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_KickOffTournament *KickOffTournamentSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _KickOffTournament.Contract.GetRoleAdmin(&_KickOffTournament.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_KickOffTournament *KickOffTournamentCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _KickOffTournament.Contract.GetRoleAdmin(&_KickOffTournament.CallOpts, role)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_KickOffTournament *KickOffTournamentCaller) GetRoleMember(opts *bind.CallOpts, role [32]byte, index *big.Int) (common.Address, error) {
	var out []interface{}
	err := _KickOffTournament.contract.Call(opts, &out, "getRoleMember", role, index)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_KickOffTournament *KickOffTournamentSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _KickOffTournament.Contract.GetRoleMember(&_KickOffTournament.CallOpts, role, index)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_KickOffTournament *KickOffTournamentCallerSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _KickOffTournament.Contract.GetRoleMember(&_KickOffTournament.CallOpts, role, index)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_KickOffTournament *KickOffTournamentCaller) GetRoleMemberCount(opts *bind.CallOpts, role [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _KickOffTournament.contract.Call(opts, &out, "getRoleMemberCount", role)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_KickOffTournament *KickOffTournamentSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _KickOffTournament.Contract.GetRoleMemberCount(&_KickOffTournament.CallOpts, role)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_KickOffTournament *KickOffTournamentCallerSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _KickOffTournament.Contract.GetRoleMemberCount(&_KickOffTournament.CallOpts, role)
}

// GetTotalScore is a free data retrieval call binding the contract method 0xd283b3c5.
//
// Solidity: function getTotalScore(address register) view returns(uint8 score)
func (_KickOffTournament *KickOffTournamentCaller) GetTotalScore(opts *bind.CallOpts, register common.Address) (uint8, error) {
	var out []interface{}
	err := _KickOffTournament.contract.Call(opts, &out, "getTotalScore", register)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// GetTotalScore is a free data retrieval call binding the contract method 0xd283b3c5.
//
// Solidity: function getTotalScore(address register) view returns(uint8 score)
func (_KickOffTournament *KickOffTournamentSession) GetTotalScore(register common.Address) (uint8, error) {
	return _KickOffTournament.Contract.GetTotalScore(&_KickOffTournament.CallOpts, register)
}

// GetTotalScore is a free data retrieval call binding the contract method 0xd283b3c5.
//
// Solidity: function getTotalScore(address register) view returns(uint8 score)
func (_KickOffTournament *KickOffTournamentCallerSession) GetTotalScore(register common.Address) (uint8, error) {
	return _KickOffTournament.Contract.GetTotalScore(&_KickOffTournament.CallOpts, register)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_KickOffTournament *KickOffTournamentCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _KickOffTournament.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_KickOffTournament *KickOffTournamentSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _KickOffTournament.Contract.HasRole(&_KickOffTournament.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_KickOffTournament *KickOffTournamentCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _KickOffTournament.Contract.HasRole(&_KickOffTournament.CallOpts, role, account)
}

// OnERC721Received is a free data retrieval call binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) pure returns(bytes4)
func (_KickOffTournament *KickOffTournamentCaller) OnERC721Received(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) ([4]byte, error) {
	var out []interface{}
	err := _KickOffTournament.contract.Call(opts, &out, "onERC721Received", arg0, arg1, arg2, arg3)

	if err != nil {
		return *new([4]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)

	return out0, err

}

// OnERC721Received is a free data retrieval call binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) pure returns(bytes4)
func (_KickOffTournament *KickOffTournamentSession) OnERC721Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) ([4]byte, error) {
	return _KickOffTournament.Contract.OnERC721Received(&_KickOffTournament.CallOpts, arg0, arg1, arg2, arg3)
}

// OnERC721Received is a free data retrieval call binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) pure returns(bytes4)
func (_KickOffTournament *KickOffTournamentCallerSession) OnERC721Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) ([4]byte, error) {
	return _KickOffTournament.Contract.OnERC721Received(&_KickOffTournament.CallOpts, arg0, arg1, arg2, arg3)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_KickOffTournament *KickOffTournamentCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _KickOffTournament.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_KickOffTournament *KickOffTournamentSession) Owner() (common.Address, error) {
	return _KickOffTournament.Contract.Owner(&_KickOffTournament.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_KickOffTournament *KickOffTournamentCallerSession) Owner() (common.Address, error) {
	return _KickOffTournament.Contract.Owner(&_KickOffTournament.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_KickOffTournament *KickOffTournamentCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _KickOffTournament.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_KickOffTournament *KickOffTournamentSession) Paused() (bool, error) {
	return _KickOffTournament.Contract.Paused(&_KickOffTournament.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_KickOffTournament *KickOffTournamentCallerSession) Paused() (bool, error) {
	return _KickOffTournament.Contract.Paused(&_KickOffTournament.CallOpts)
}

// RaceInfo is a free data retrieval call binding the contract method 0xef8597fd.
//
// Solidity: function raceInfo(bytes32 raceId) view returns((uint8,uint32,bytes27))
func (_KickOffTournament *KickOffTournamentCaller) RaceInfo(opts *bind.CallOpts, raceId [32]byte) (ITournamentRace, error) {
	var out []interface{}
	err := _KickOffTournament.contract.Call(opts, &out, "raceInfo", raceId)

	if err != nil {
		return *new(ITournamentRace), err
	}

	out0 := *abi.ConvertType(out[0], new(ITournamentRace)).(*ITournamentRace)

	return out0, err

}

// RaceInfo is a free data retrieval call binding the contract method 0xef8597fd.
//
// Solidity: function raceInfo(bytes32 raceId) view returns((uint8,uint32,bytes27))
func (_KickOffTournament *KickOffTournamentSession) RaceInfo(raceId [32]byte) (ITournamentRace, error) {
	return _KickOffTournament.Contract.RaceInfo(&_KickOffTournament.CallOpts, raceId)
}

// RaceInfo is a free data retrieval call binding the contract method 0xef8597fd.
//
// Solidity: function raceInfo(bytes32 raceId) view returns((uint8,uint32,bytes27))
func (_KickOffTournament *KickOffTournamentCallerSession) RaceInfo(raceId [32]byte) (ITournamentRace, error) {
	return _KickOffTournament.Contract.RaceInfo(&_KickOffTournament.CallOpts, raceId)
}

// RegisteredSlot is a free data retrieval call binding the contract method 0x96fd6980.
//
// Solidity: function registeredSlot(address register, bytes32 raceId) view returns(uint8)
func (_KickOffTournament *KickOffTournamentCaller) RegisteredSlot(opts *bind.CallOpts, register common.Address, raceId [32]byte) (uint8, error) {
	var out []interface{}
	err := _KickOffTournament.contract.Call(opts, &out, "registeredSlot", register, raceId)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// RegisteredSlot is a free data retrieval call binding the contract method 0x96fd6980.
//
// Solidity: function registeredSlot(address register, bytes32 raceId) view returns(uint8)
func (_KickOffTournament *KickOffTournamentSession) RegisteredSlot(register common.Address, raceId [32]byte) (uint8, error) {
	return _KickOffTournament.Contract.RegisteredSlot(&_KickOffTournament.CallOpts, register, raceId)
}

// RegisteredSlot is a free data retrieval call binding the contract method 0x96fd6980.
//
// Solidity: function registeredSlot(address register, bytes32 raceId) view returns(uint8)
func (_KickOffTournament *KickOffTournamentCallerSession) RegisteredSlot(register common.Address, raceId [32]byte) (uint8, error) {
	return _KickOffTournament.Contract.RegisteredSlot(&_KickOffTournament.CallOpts, register, raceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_KickOffTournament *KickOffTournamentCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _KickOffTournament.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_KickOffTournament *KickOffTournamentSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _KickOffTournament.Contract.SupportsInterface(&_KickOffTournament.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_KickOffTournament *KickOffTournamentCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _KickOffTournament.Contract.SupportsInterface(&_KickOffTournament.CallOpts, interfaceId)
}

// TournamentInfo is a free data retrieval call binding the contract method 0xc1fab611.
//
// Solidity: function tournamentInfo() view returns((uint8,uint8,uint8,string))
func (_KickOffTournament *KickOffTournamentCaller) TournamentInfo(opts *bind.CallOpts) (ITournamentTournamentInfo, error) {
	var out []interface{}
	err := _KickOffTournament.contract.Call(opts, &out, "tournamentInfo")

	if err != nil {
		return *new(ITournamentTournamentInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(ITournamentTournamentInfo)).(*ITournamentTournamentInfo)

	return out0, err

}

// TournamentInfo is a free data retrieval call binding the contract method 0xc1fab611.
//
// Solidity: function tournamentInfo() view returns((uint8,uint8,uint8,string))
func (_KickOffTournament *KickOffTournamentSession) TournamentInfo() (ITournamentTournamentInfo, error) {
	return _KickOffTournament.Contract.TournamentInfo(&_KickOffTournament.CallOpts)
}

// TournamentInfo is a free data retrieval call binding the contract method 0xc1fab611.
//
// Solidity: function tournamentInfo() view returns((uint8,uint8,uint8,string))
func (_KickOffTournament *KickOffTournamentCallerSession) TournamentInfo() (ITournamentTournamentInfo, error) {
	return _KickOffTournament.Contract.TournamentInfo(&_KickOffTournament.CallOpts)
}

// CancelRace is a paid mutator transaction binding the contract method 0xc05c107d.
//
// Solidity: function cancelRace(bytes32 raceId) returns()
func (_KickOffTournament *KickOffTournamentTransactor) CancelRace(opts *bind.TransactOpts, raceId [32]byte) (*types.Transaction, error) {
	return _KickOffTournament.contract.Transact(opts, "cancelRace", raceId)
}

// CancelRace is a paid mutator transaction binding the contract method 0xc05c107d.
//
// Solidity: function cancelRace(bytes32 raceId) returns()
func (_KickOffTournament *KickOffTournamentSession) CancelRace(raceId [32]byte) (*types.Transaction, error) {
	return _KickOffTournament.Contract.CancelRace(&_KickOffTournament.TransactOpts, raceId)
}

// CancelRace is a paid mutator transaction binding the contract method 0xc05c107d.
//
// Solidity: function cancelRace(bytes32 raceId) returns()
func (_KickOffTournament *KickOffTournamentTransactorSession) CancelRace(raceId [32]byte) (*types.Transaction, error) {
	return _KickOffTournament.Contract.CancelRace(&_KickOffTournament.TransactOpts, raceId)
}

// CreateRace is a paid mutator transaction binding the contract method 0x928e03da.
//
// Solidity: function createRace(uint8 noSlot, uint32 startAt) returns(bytes32)
func (_KickOffTournament *KickOffTournamentTransactor) CreateRace(opts *bind.TransactOpts, noSlot uint8, startAt uint32) (*types.Transaction, error) {
	return _KickOffTournament.contract.Transact(opts, "createRace", noSlot, startAt)
}

// CreateRace is a paid mutator transaction binding the contract method 0x928e03da.
//
// Solidity: function createRace(uint8 noSlot, uint32 startAt) returns(bytes32)
func (_KickOffTournament *KickOffTournamentSession) CreateRace(noSlot uint8, startAt uint32) (*types.Transaction, error) {
	return _KickOffTournament.Contract.CreateRace(&_KickOffTournament.TransactOpts, noSlot, startAt)
}

// CreateRace is a paid mutator transaction binding the contract method 0x928e03da.
//
// Solidity: function createRace(uint8 noSlot, uint32 startAt) returns(bytes32)
func (_KickOffTournament *KickOffTournamentTransactorSession) CreateRace(noSlot uint8, startAt uint32) (*types.Transaction, error) {
	return _KickOffTournament.Contract.CreateRace(&_KickOffTournament.TransactOpts, noSlot, startAt)
}

// GrandReward is a paid mutator transaction binding the contract method 0xda0df24f.
//
// Solidity: function grandReward(uint8 winnerIndex, address winner, uint256 nftRewardId, string tokenURI) returns()
func (_KickOffTournament *KickOffTournamentTransactor) GrandReward(opts *bind.TransactOpts, winnerIndex uint8, winner common.Address, nftRewardId *big.Int, tokenURI string) (*types.Transaction, error) {
	return _KickOffTournament.contract.Transact(opts, "grandReward", winnerIndex, winner, nftRewardId, tokenURI)
}

// GrandReward is a paid mutator transaction binding the contract method 0xda0df24f.
//
// Solidity: function grandReward(uint8 winnerIndex, address winner, uint256 nftRewardId, string tokenURI) returns()
func (_KickOffTournament *KickOffTournamentSession) GrandReward(winnerIndex uint8, winner common.Address, nftRewardId *big.Int, tokenURI string) (*types.Transaction, error) {
	return _KickOffTournament.Contract.GrandReward(&_KickOffTournament.TransactOpts, winnerIndex, winner, nftRewardId, tokenURI)
}

// GrandReward is a paid mutator transaction binding the contract method 0xda0df24f.
//
// Solidity: function grandReward(uint8 winnerIndex, address winner, uint256 nftRewardId, string tokenURI) returns()
func (_KickOffTournament *KickOffTournamentTransactorSession) GrandReward(winnerIndex uint8, winner common.Address, nftRewardId *big.Int, tokenURI string) (*types.Transaction, error) {
	return _KickOffTournament.Contract.GrandReward(&_KickOffTournament.TransactOpts, winnerIndex, winner, nftRewardId, tokenURI)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_KickOffTournament *KickOffTournamentTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _KickOffTournament.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_KickOffTournament *KickOffTournamentSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _KickOffTournament.Contract.GrantRole(&_KickOffTournament.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_KickOffTournament *KickOffTournamentTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _KickOffTournament.Contract.GrantRole(&_KickOffTournament.TransactOpts, role, account)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_KickOffTournament *KickOffTournamentTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _KickOffTournament.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_KickOffTournament *KickOffTournamentSession) Pause() (*types.Transaction, error) {
	return _KickOffTournament.Contract.Pause(&_KickOffTournament.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_KickOffTournament *KickOffTournamentTransactorSession) Pause() (*types.Transaction, error) {
	return _KickOffTournament.Contract.Pause(&_KickOffTournament.TransactOpts)
}

// RegisterRace is a paid mutator transaction binding the contract method 0x1e621ebe.
//
// Solidity: function registerRace(uint8 slotId, bytes32 raceId) returns()
func (_KickOffTournament *KickOffTournamentTransactor) RegisterRace(opts *bind.TransactOpts, slotId uint8, raceId [32]byte) (*types.Transaction, error) {
	return _KickOffTournament.contract.Transact(opts, "registerRace", slotId, raceId)
}

// RegisterRace is a paid mutator transaction binding the contract method 0x1e621ebe.
//
// Solidity: function registerRace(uint8 slotId, bytes32 raceId) returns()
func (_KickOffTournament *KickOffTournamentSession) RegisterRace(slotId uint8, raceId [32]byte) (*types.Transaction, error) {
	return _KickOffTournament.Contract.RegisterRace(&_KickOffTournament.TransactOpts, slotId, raceId)
}

// RegisterRace is a paid mutator transaction binding the contract method 0x1e621ebe.
//
// Solidity: function registerRace(uint8 slotId, bytes32 raceId) returns()
func (_KickOffTournament *KickOffTournamentTransactorSession) RegisterRace(slotId uint8, raceId [32]byte) (*types.Transaction, error) {
	return _KickOffTournament.Contract.RegisterRace(&_KickOffTournament.TransactOpts, slotId, raceId)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_KickOffTournament *KickOffTournamentTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _KickOffTournament.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_KickOffTournament *KickOffTournamentSession) RenounceOwnership() (*types.Transaction, error) {
	return _KickOffTournament.Contract.RenounceOwnership(&_KickOffTournament.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_KickOffTournament *KickOffTournamentTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _KickOffTournament.Contract.RenounceOwnership(&_KickOffTournament.TransactOpts)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_KickOffTournament *KickOffTournamentTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _KickOffTournament.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_KickOffTournament *KickOffTournamentSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _KickOffTournament.Contract.RenounceRole(&_KickOffTournament.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_KickOffTournament *KickOffTournamentTransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _KickOffTournament.Contract.RenounceRole(&_KickOffTournament.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_KickOffTournament *KickOffTournamentTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _KickOffTournament.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_KickOffTournament *KickOffTournamentSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _KickOffTournament.Contract.RevokeRole(&_KickOffTournament.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_KickOffTournament *KickOffTournamentTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _KickOffTournament.Contract.RevokeRole(&_KickOffTournament.TransactOpts, role, account)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_KickOffTournament *KickOffTournamentTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _KickOffTournament.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_KickOffTournament *KickOffTournamentSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _KickOffTournament.Contract.TransferOwnership(&_KickOffTournament.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_KickOffTournament *KickOffTournamentTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _KickOffTournament.Contract.TransferOwnership(&_KickOffTournament.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_KickOffTournament *KickOffTournamentTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _KickOffTournament.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_KickOffTournament *KickOffTournamentSession) Unpause() (*types.Transaction, error) {
	return _KickOffTournament.Contract.Unpause(&_KickOffTournament.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_KickOffTournament *KickOffTournamentTransactorSession) Unpause() (*types.Transaction, error) {
	return _KickOffTournament.Contract.Unpause(&_KickOffTournament.TransactOpts)
}

// UpdateRaceResult is a paid mutator transaction binding the contract method 0xdd0b3778.
//
// Solidity: function updateRaceResult(bytes32 raceId, bytes27 result) returns()
func (_KickOffTournament *KickOffTournamentTransactor) UpdateRaceResult(opts *bind.TransactOpts, raceId [32]byte, result [27]byte) (*types.Transaction, error) {
	return _KickOffTournament.contract.Transact(opts, "updateRaceResult", raceId, result)
}

// UpdateRaceResult is a paid mutator transaction binding the contract method 0xdd0b3778.
//
// Solidity: function updateRaceResult(bytes32 raceId, bytes27 result) returns()
func (_KickOffTournament *KickOffTournamentSession) UpdateRaceResult(raceId [32]byte, result [27]byte) (*types.Transaction, error) {
	return _KickOffTournament.Contract.UpdateRaceResult(&_KickOffTournament.TransactOpts, raceId, result)
}

// UpdateRaceResult is a paid mutator transaction binding the contract method 0xdd0b3778.
//
// Solidity: function updateRaceResult(bytes32 raceId, bytes27 result) returns()
func (_KickOffTournament *KickOffTournamentTransactorSession) UpdateRaceResult(raceId [32]byte, result [27]byte) (*types.Transaction, error) {
	return _KickOffTournament.Contract.UpdateRaceResult(&_KickOffTournament.TransactOpts, raceId, result)
}

// KickOffTournamentOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the KickOffTournament contract.
type KickOffTournamentOwnershipTransferredIterator struct {
	Event *KickOffTournamentOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *KickOffTournamentOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KickOffTournamentOwnershipTransferred)
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
		it.Event = new(KickOffTournamentOwnershipTransferred)
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
func (it *KickOffTournamentOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KickOffTournamentOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KickOffTournamentOwnershipTransferred represents a OwnershipTransferred event raised by the KickOffTournament contract.
type KickOffTournamentOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_KickOffTournament *KickOffTournamentFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*KickOffTournamentOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _KickOffTournament.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &KickOffTournamentOwnershipTransferredIterator{contract: _KickOffTournament.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_KickOffTournament *KickOffTournamentFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *KickOffTournamentOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _KickOffTournament.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KickOffTournamentOwnershipTransferred)
				if err := _KickOffTournament.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_KickOffTournament *KickOffTournamentFilterer) ParseOwnershipTransferred(log types.Log) (*KickOffTournamentOwnershipTransferred, error) {
	event := new(KickOffTournamentOwnershipTransferred)
	if err := _KickOffTournament.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// KickOffTournamentPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the KickOffTournament contract.
type KickOffTournamentPausedIterator struct {
	Event *KickOffTournamentPaused // Event containing the contract specifics and raw log

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
func (it *KickOffTournamentPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KickOffTournamentPaused)
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
		it.Event = new(KickOffTournamentPaused)
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
func (it *KickOffTournamentPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KickOffTournamentPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KickOffTournamentPaused represents a Paused event raised by the KickOffTournament contract.
type KickOffTournamentPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_KickOffTournament *KickOffTournamentFilterer) FilterPaused(opts *bind.FilterOpts) (*KickOffTournamentPausedIterator, error) {

	logs, sub, err := _KickOffTournament.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &KickOffTournamentPausedIterator{contract: _KickOffTournament.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_KickOffTournament *KickOffTournamentFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *KickOffTournamentPaused) (event.Subscription, error) {

	logs, sub, err := _KickOffTournament.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KickOffTournamentPaused)
				if err := _KickOffTournament.contract.UnpackLog(event, "Paused", log); err != nil {
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

// ParsePaused is a log parse operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_KickOffTournament *KickOffTournamentFilterer) ParsePaused(log types.Log) (*KickOffTournamentPaused, error) {
	event := new(KickOffTournamentPaused)
	if err := _KickOffTournament.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// KickOffTournamentRaceCancelledIterator is returned from FilterRaceCancelled and is used to iterate over the raw logs and unpacked data for RaceCancelled events raised by the KickOffTournament contract.
type KickOffTournamentRaceCancelledIterator struct {
	Event *KickOffTournamentRaceCancelled // Event containing the contract specifics and raw log

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
func (it *KickOffTournamentRaceCancelledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KickOffTournamentRaceCancelled)
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
		it.Event = new(KickOffTournamentRaceCancelled)
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
func (it *KickOffTournamentRaceCancelledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KickOffTournamentRaceCancelledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KickOffTournamentRaceCancelled represents a RaceCancelled event raised by the KickOffTournament contract.
type KickOffTournamentRaceCancelled struct {
	Id  [32]byte
	Raw types.Log // Blockchain specific contextual infos
}

// FilterRaceCancelled is a free log retrieval operation binding the contract event 0x3f3669b11af9b00fb2a95f8ded36a79a2e1d5374efbc5ec6b2b5b96ceaf9ac6a.
//
// Solidity: event RaceCancelled(bytes32 id)
func (_KickOffTournament *KickOffTournamentFilterer) FilterRaceCancelled(opts *bind.FilterOpts) (*KickOffTournamentRaceCancelledIterator, error) {

	logs, sub, err := _KickOffTournament.contract.FilterLogs(opts, "RaceCancelled")
	if err != nil {
		return nil, err
	}
	return &KickOffTournamentRaceCancelledIterator{contract: _KickOffTournament.contract, event: "RaceCancelled", logs: logs, sub: sub}, nil
}

// WatchRaceCancelled is a free log subscription operation binding the contract event 0x3f3669b11af9b00fb2a95f8ded36a79a2e1d5374efbc5ec6b2b5b96ceaf9ac6a.
//
// Solidity: event RaceCancelled(bytes32 id)
func (_KickOffTournament *KickOffTournamentFilterer) WatchRaceCancelled(opts *bind.WatchOpts, sink chan<- *KickOffTournamentRaceCancelled) (event.Subscription, error) {

	logs, sub, err := _KickOffTournament.contract.WatchLogs(opts, "RaceCancelled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KickOffTournamentRaceCancelled)
				if err := _KickOffTournament.contract.UnpackLog(event, "RaceCancelled", log); err != nil {
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
func (_KickOffTournament *KickOffTournamentFilterer) ParseRaceCancelled(log types.Log) (*KickOffTournamentRaceCancelled, error) {
	event := new(KickOffTournamentRaceCancelled)
	if err := _KickOffTournament.contract.UnpackLog(event, "RaceCancelled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// KickOffTournamentRaceCreatedIterator is returned from FilterRaceCreated and is used to iterate over the raw logs and unpacked data for RaceCreated events raised by the KickOffTournament contract.
type KickOffTournamentRaceCreatedIterator struct {
	Event *KickOffTournamentRaceCreated // Event containing the contract specifics and raw log

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
func (it *KickOffTournamentRaceCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KickOffTournamentRaceCreated)
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
		it.Event = new(KickOffTournamentRaceCreated)
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
func (it *KickOffTournamentRaceCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KickOffTournamentRaceCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KickOffTournamentRaceCreated represents a RaceCreated event raised by the KickOffTournament contract.
type KickOffTournamentRaceCreated struct {
	NoSlot  uint8
	RaceNo  uint8
	StartAt uint32
	Id      [32]byte
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRaceCreated is a free log retrieval operation binding the contract event 0x58e30581b52e563d6f294b197724e473591cf7668760f3bae9a2678502011536.
//
// Solidity: event RaceCreated(uint8 noSlot, uint8 raceNo, uint32 startAt, bytes32 id)
func (_KickOffTournament *KickOffTournamentFilterer) FilterRaceCreated(opts *bind.FilterOpts) (*KickOffTournamentRaceCreatedIterator, error) {

	logs, sub, err := _KickOffTournament.contract.FilterLogs(opts, "RaceCreated")
	if err != nil {
		return nil, err
	}
	return &KickOffTournamentRaceCreatedIterator{contract: _KickOffTournament.contract, event: "RaceCreated", logs: logs, sub: sub}, nil
}

// WatchRaceCreated is a free log subscription operation binding the contract event 0x58e30581b52e563d6f294b197724e473591cf7668760f3bae9a2678502011536.
//
// Solidity: event RaceCreated(uint8 noSlot, uint8 raceNo, uint32 startAt, bytes32 id)
func (_KickOffTournament *KickOffTournamentFilterer) WatchRaceCreated(opts *bind.WatchOpts, sink chan<- *KickOffTournamentRaceCreated) (event.Subscription, error) {

	logs, sub, err := _KickOffTournament.contract.WatchLogs(opts, "RaceCreated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KickOffTournamentRaceCreated)
				if err := _KickOffTournament.contract.UnpackLog(event, "RaceCreated", log); err != nil {
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
func (_KickOffTournament *KickOffTournamentFilterer) ParseRaceCreated(log types.Log) (*KickOffTournamentRaceCreated, error) {
	event := new(KickOffTournamentRaceCreated)
	if err := _KickOffTournament.contract.UnpackLog(event, "RaceCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// KickOffTournamentRaceResultUpdatedIterator is returned from FilterRaceResultUpdated and is used to iterate over the raw logs and unpacked data for RaceResultUpdated events raised by the KickOffTournament contract.
type KickOffTournamentRaceResultUpdatedIterator struct {
	Event *KickOffTournamentRaceResultUpdated // Event containing the contract specifics and raw log

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
func (it *KickOffTournamentRaceResultUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KickOffTournamentRaceResultUpdated)
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
		it.Event = new(KickOffTournamentRaceResultUpdated)
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
func (it *KickOffTournamentRaceResultUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KickOffTournamentRaceResultUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KickOffTournamentRaceResultUpdated represents a RaceResultUpdated event raised by the KickOffTournament contract.
type KickOffTournamentRaceResultUpdated struct {
	Id     [32]byte
	Result [27]byte
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRaceResultUpdated is a free log retrieval operation binding the contract event 0x2ab6bf279c23ac56ddfab0e48cd46e6576b30742d689be7cabf12d8f279c7870.
//
// Solidity: event RaceResultUpdated(bytes32 id, bytes27 result)
func (_KickOffTournament *KickOffTournamentFilterer) FilterRaceResultUpdated(opts *bind.FilterOpts) (*KickOffTournamentRaceResultUpdatedIterator, error) {

	logs, sub, err := _KickOffTournament.contract.FilterLogs(opts, "RaceResultUpdated")
	if err != nil {
		return nil, err
	}
	return &KickOffTournamentRaceResultUpdatedIterator{contract: _KickOffTournament.contract, event: "RaceResultUpdated", logs: logs, sub: sub}, nil
}

// WatchRaceResultUpdated is a free log subscription operation binding the contract event 0x2ab6bf279c23ac56ddfab0e48cd46e6576b30742d689be7cabf12d8f279c7870.
//
// Solidity: event RaceResultUpdated(bytes32 id, bytes27 result)
func (_KickOffTournament *KickOffTournamentFilterer) WatchRaceResultUpdated(opts *bind.WatchOpts, sink chan<- *KickOffTournamentRaceResultUpdated) (event.Subscription, error) {

	logs, sub, err := _KickOffTournament.contract.WatchLogs(opts, "RaceResultUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KickOffTournamentRaceResultUpdated)
				if err := _KickOffTournament.contract.UnpackLog(event, "RaceResultUpdated", log); err != nil {
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
func (_KickOffTournament *KickOffTournamentFilterer) ParseRaceResultUpdated(log types.Log) (*KickOffTournamentRaceResultUpdated, error) {
	event := new(KickOffTournamentRaceResultUpdated)
	if err := _KickOffTournament.contract.UnpackLog(event, "RaceResultUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// KickOffTournamentRegisteredIterator is returned from FilterRegistered and is used to iterate over the raw logs and unpacked data for Registered events raised by the KickOffTournament contract.
type KickOffTournamentRegisteredIterator struct {
	Event *KickOffTournamentRegistered // Event containing the contract specifics and raw log

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
func (it *KickOffTournamentRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KickOffTournamentRegistered)
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
		it.Event = new(KickOffTournamentRegistered)
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
func (it *KickOffTournamentRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KickOffTournamentRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KickOffTournamentRegistered represents a Registered event raised by the KickOffTournament contract.
type KickOffTournamentRegistered struct {
	SlotId      uint8
	Participant common.Address
	RaceId      [32]byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterRegistered is a free log retrieval operation binding the contract event 0xad76e1cb5a0536a10bd2c282409cd1301bc63cff1eeed6b269642f1c24eaaa8f.
//
// Solidity: event Registered(uint8 slotId, address participant, bytes32 raceId)
func (_KickOffTournament *KickOffTournamentFilterer) FilterRegistered(opts *bind.FilterOpts) (*KickOffTournamentRegisteredIterator, error) {

	logs, sub, err := _KickOffTournament.contract.FilterLogs(opts, "Registered")
	if err != nil {
		return nil, err
	}
	return &KickOffTournamentRegisteredIterator{contract: _KickOffTournament.contract, event: "Registered", logs: logs, sub: sub}, nil
}

// WatchRegistered is a free log subscription operation binding the contract event 0xad76e1cb5a0536a10bd2c282409cd1301bc63cff1eeed6b269642f1c24eaaa8f.
//
// Solidity: event Registered(uint8 slotId, address participant, bytes32 raceId)
func (_KickOffTournament *KickOffTournamentFilterer) WatchRegistered(opts *bind.WatchOpts, sink chan<- *KickOffTournamentRegistered) (event.Subscription, error) {

	logs, sub, err := _KickOffTournament.contract.WatchLogs(opts, "Registered")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KickOffTournamentRegistered)
				if err := _KickOffTournament.contract.UnpackLog(event, "Registered", log); err != nil {
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
func (_KickOffTournament *KickOffTournamentFilterer) ParseRegistered(log types.Log) (*KickOffTournamentRegistered, error) {
	event := new(KickOffTournamentRegistered)
	if err := _KickOffTournament.contract.UnpackLog(event, "Registered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// KickOffTournamentRewardGrantedIterator is returned from FilterRewardGranted and is used to iterate over the raw logs and unpacked data for RewardGranted events raised by the KickOffTournament contract.
type KickOffTournamentRewardGrantedIterator struct {
	Event *KickOffTournamentRewardGranted // Event containing the contract specifics and raw log

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
func (it *KickOffTournamentRewardGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KickOffTournamentRewardGranted)
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
		it.Event = new(KickOffTournamentRewardGranted)
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
func (it *KickOffTournamentRewardGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KickOffTournamentRewardGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KickOffTournamentRewardGranted represents a RewardGranted event raised by the KickOffTournament contract.
type KickOffTournamentRewardGranted struct {
	WinnerIndex uint8
	Score       uint8
	Winner      common.Address
	NftRewardId *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterRewardGranted is a free log retrieval operation binding the contract event 0x665f81dad11e5f7be344a68d526bf2281aac8b94951d346d51391531bfcae85e.
//
// Solidity: event RewardGranted(uint8 winnerIndex, uint8 score, address winner, uint256 nftRewardId)
func (_KickOffTournament *KickOffTournamentFilterer) FilterRewardGranted(opts *bind.FilterOpts) (*KickOffTournamentRewardGrantedIterator, error) {

	logs, sub, err := _KickOffTournament.contract.FilterLogs(opts, "RewardGranted")
	if err != nil {
		return nil, err
	}
	return &KickOffTournamentRewardGrantedIterator{contract: _KickOffTournament.contract, event: "RewardGranted", logs: logs, sub: sub}, nil
}

// WatchRewardGranted is a free log subscription operation binding the contract event 0x665f81dad11e5f7be344a68d526bf2281aac8b94951d346d51391531bfcae85e.
//
// Solidity: event RewardGranted(uint8 winnerIndex, uint8 score, address winner, uint256 nftRewardId)
func (_KickOffTournament *KickOffTournamentFilterer) WatchRewardGranted(opts *bind.WatchOpts, sink chan<- *KickOffTournamentRewardGranted) (event.Subscription, error) {

	logs, sub, err := _KickOffTournament.contract.WatchLogs(opts, "RewardGranted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KickOffTournamentRewardGranted)
				if err := _KickOffTournament.contract.UnpackLog(event, "RewardGranted", log); err != nil {
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
func (_KickOffTournament *KickOffTournamentFilterer) ParseRewardGranted(log types.Log) (*KickOffTournamentRewardGranted, error) {
	event := new(KickOffTournamentRewardGranted)
	if err := _KickOffTournament.contract.UnpackLog(event, "RewardGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// KickOffTournamentRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the KickOffTournament contract.
type KickOffTournamentRoleAdminChangedIterator struct {
	Event *KickOffTournamentRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *KickOffTournamentRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KickOffTournamentRoleAdminChanged)
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
		it.Event = new(KickOffTournamentRoleAdminChanged)
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
func (it *KickOffTournamentRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KickOffTournamentRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KickOffTournamentRoleAdminChanged represents a RoleAdminChanged event raised by the KickOffTournament contract.
type KickOffTournamentRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_KickOffTournament *KickOffTournamentFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*KickOffTournamentRoleAdminChangedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _KickOffTournament.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &KickOffTournamentRoleAdminChangedIterator{contract: _KickOffTournament.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_KickOffTournament *KickOffTournamentFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *KickOffTournamentRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _KickOffTournament.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KickOffTournamentRoleAdminChanged)
				if err := _KickOffTournament.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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

// ParseRoleAdminChanged is a log parse operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_KickOffTournament *KickOffTournamentFilterer) ParseRoleAdminChanged(log types.Log) (*KickOffTournamentRoleAdminChanged, error) {
	event := new(KickOffTournamentRoleAdminChanged)
	if err := _KickOffTournament.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// KickOffTournamentRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the KickOffTournament contract.
type KickOffTournamentRoleGrantedIterator struct {
	Event *KickOffTournamentRoleGranted // Event containing the contract specifics and raw log

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
func (it *KickOffTournamentRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KickOffTournamentRoleGranted)
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
		it.Event = new(KickOffTournamentRoleGranted)
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
func (it *KickOffTournamentRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KickOffTournamentRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KickOffTournamentRoleGranted represents a RoleGranted event raised by the KickOffTournament contract.
type KickOffTournamentRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_KickOffTournament *KickOffTournamentFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*KickOffTournamentRoleGrantedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _KickOffTournament.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &KickOffTournamentRoleGrantedIterator{contract: _KickOffTournament.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_KickOffTournament *KickOffTournamentFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *KickOffTournamentRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _KickOffTournament.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KickOffTournamentRoleGranted)
				if err := _KickOffTournament.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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

// ParseRoleGranted is a log parse operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_KickOffTournament *KickOffTournamentFilterer) ParseRoleGranted(log types.Log) (*KickOffTournamentRoleGranted, error) {
	event := new(KickOffTournamentRoleGranted)
	if err := _KickOffTournament.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// KickOffTournamentRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the KickOffTournament contract.
type KickOffTournamentRoleRevokedIterator struct {
	Event *KickOffTournamentRoleRevoked // Event containing the contract specifics and raw log

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
func (it *KickOffTournamentRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KickOffTournamentRoleRevoked)
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
		it.Event = new(KickOffTournamentRoleRevoked)
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
func (it *KickOffTournamentRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KickOffTournamentRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KickOffTournamentRoleRevoked represents a RoleRevoked event raised by the KickOffTournament contract.
type KickOffTournamentRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_KickOffTournament *KickOffTournamentFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*KickOffTournamentRoleRevokedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _KickOffTournament.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &KickOffTournamentRoleRevokedIterator{contract: _KickOffTournament.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_KickOffTournament *KickOffTournamentFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *KickOffTournamentRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _KickOffTournament.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KickOffTournamentRoleRevoked)
				if err := _KickOffTournament.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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

// ParseRoleRevoked is a log parse operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_KickOffTournament *KickOffTournamentFilterer) ParseRoleRevoked(log types.Log) (*KickOffTournamentRoleRevoked, error) {
	event := new(KickOffTournamentRoleRevoked)
	if err := _KickOffTournament.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// KickOffTournamentUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the KickOffTournament contract.
type KickOffTournamentUnpausedIterator struct {
	Event *KickOffTournamentUnpaused // Event containing the contract specifics and raw log

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
func (it *KickOffTournamentUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KickOffTournamentUnpaused)
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
		it.Event = new(KickOffTournamentUnpaused)
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
func (it *KickOffTournamentUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KickOffTournamentUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KickOffTournamentUnpaused represents a Unpaused event raised by the KickOffTournament contract.
type KickOffTournamentUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_KickOffTournament *KickOffTournamentFilterer) FilterUnpaused(opts *bind.FilterOpts) (*KickOffTournamentUnpausedIterator, error) {

	logs, sub, err := _KickOffTournament.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &KickOffTournamentUnpausedIterator{contract: _KickOffTournament.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_KickOffTournament *KickOffTournamentFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *KickOffTournamentUnpaused) (event.Subscription, error) {

	logs, sub, err := _KickOffTournament.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KickOffTournamentUnpaused)
				if err := _KickOffTournament.contract.UnpackLog(event, "Unpaused", log); err != nil {
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

// ParseUnpaused is a log parse operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_KickOffTournament *KickOffTournamentFilterer) ParseUnpaused(log types.Log) (*KickOffTournamentUnpaused, error) {
	event := new(KickOffTournamentUnpaused)
	if err := _KickOffTournament.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
