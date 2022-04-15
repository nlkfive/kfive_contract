// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package IMarketplaceStorage

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

// IAuctionAuction is an auto generated low-level Go binding around an user-defined struct.
type IAuctionAuction struct {
	Seller        common.Address
	HighestBidder common.Address
	Id            [32]byte
	BiddingEnd    *big.Int
	RevealEnd     *big.Int
	HighestBid    *big.Int
	StartPrice    *big.Int
}

// IOrderOrder is an auto generated low-level Go binding around an user-defined struct.
type IOrderOrder struct {
	OrderId    [32]byte
	Seller     common.Address
	NftAddress common.Address
	Price      *big.Int
	ExpiredAt  *big.Int
}

// IMarketplaceStorageMetaData contains all meta data concerning the IMarketplaceStorage contract.
var IMarketplaceStorageMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"NotExisted\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotRunning\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"}],\"name\":\"TooEarly\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"}],\"name\":\"TooLate\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"auctionId\",\"type\":\"bytes32\"}],\"name\":\"AuctionCancelled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"nftAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"auctionId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"assetId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"biddingEnd\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"revealEnd\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"startPriceInWei\",\"type\":\"uint256\"}],\"name\":\"AuctionCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"auctionId\",\"type\":\"bytes32\"}],\"name\":\"AuctionEnded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"bidder\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"auctionId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"deposit\",\"type\":\"uint256\"}],\"name\":\"AuctionRefund\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"buyer\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"auctionId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalPrice\",\"type\":\"uint256\"}],\"name\":\"AuctionSuccessful\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"bidder\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"auctionId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"blindedBid\",\"type\":\"bytes32\"}],\"name\":\"BidSuccessful\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"}],\"name\":\"OrderCancelled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"orderId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"assetId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"nftAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"priceInWei\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"expiredAt\",\"type\":\"uint256\"}],\"name\":\"OrderCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"buyer\",\"type\":\"address\"}],\"name\":\"OrderSuccessful\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"fake\",\"type\":\"bool\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"auctionId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"RevealFailed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"fake\",\"type\":\"bool\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"auctionId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"blindedBid\",\"type\":\"bytes32\"}],\"name\":\"RevealSuccessful\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"nftAsset\",\"type\":\"bytes32\"}],\"name\":\"assetIsAvailable\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"auctionId\",\"type\":\"bytes32\"}],\"name\":\"auctionIsExisted\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"existed\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"nftAsset\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"auctionId\",\"type\":\"bytes32\"}],\"name\":\"auctionIsEnded\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"ended\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"nftAsset\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"auctionId\",\"type\":\"bytes32\"}],\"name\":\"auctionIsRunning\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"existed\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"assetOwner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"nftAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"assetId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"auctionId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"biddingEnd\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"revealEnd\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startPriceInWei\",\"type\":\"uint256\"}],\"name\":\"createAuction\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"auctionId\",\"type\":\"bytes32\"}],\"name\":\"getAuction\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"highestBidder\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"biddingEnd\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"revealEnd\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"highestBid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startPrice\",\"type\":\"uint256\"}],\"internalType\":\"structIAuction.Auction\",\"name\":\"order\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"nftAsset\",\"type\":\"bytes32\"}],\"name\":\"auctionEnded\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"bidder\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"highestBid\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"auctionId\",\"type\":\"bytes32\"}],\"name\":\"updateHighestBid\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"nftAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"assetId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"orderId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"priceInWei\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expiredAt\",\"type\":\"uint256\"}],\"name\":\"createOrder\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"nftAsset\",\"type\":\"bytes32\"}],\"name\":\"getOrder\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"orderId\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"nftAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expiredAt\",\"type\":\"uint256\"}],\"internalType\":\"structIOrder.Order\",\"name\":\"order\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"nftAsset\",\"type\":\"bytes32\"}],\"name\":\"deleteOrder\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// IMarketplaceStorageABI is the input ABI used to generate the binding from.
// Deprecated: Use IMarketplaceStorageMetaData.ABI instead.
var IMarketplaceStorageABI = IMarketplaceStorageMetaData.ABI

// IMarketplaceStorage is an auto generated Go binding around an Ethereum contract.
type IMarketplaceStorage struct {
	IMarketplaceStorageCaller     // Read-only binding to the contract
	IMarketplaceStorageTransactor // Write-only binding to the contract
	IMarketplaceStorageFilterer   // Log filterer for contract events
}

// IMarketplaceStorageCaller is an auto generated read-only Go binding around an Ethereum contract.
type IMarketplaceStorageCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IMarketplaceStorageTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IMarketplaceStorageTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IMarketplaceStorageFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IMarketplaceStorageFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IMarketplaceStorageSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IMarketplaceStorageSession struct {
	Contract     *IMarketplaceStorage // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// IMarketplaceStorageCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IMarketplaceStorageCallerSession struct {
	Contract *IMarketplaceStorageCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// IMarketplaceStorageTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IMarketplaceStorageTransactorSession struct {
	Contract     *IMarketplaceStorageTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// IMarketplaceStorageRaw is an auto generated low-level Go binding around an Ethereum contract.
type IMarketplaceStorageRaw struct {
	Contract *IMarketplaceStorage // Generic contract binding to access the raw methods on
}

// IMarketplaceStorageCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IMarketplaceStorageCallerRaw struct {
	Contract *IMarketplaceStorageCaller // Generic read-only contract binding to access the raw methods on
}

// IMarketplaceStorageTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IMarketplaceStorageTransactorRaw struct {
	Contract *IMarketplaceStorageTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIMarketplaceStorage creates a new instance of IMarketplaceStorage, bound to a specific deployed contract.
func NewIMarketplaceStorage(address common.Address, backend bind.ContractBackend) (*IMarketplaceStorage, error) {
	contract, err := bindIMarketplaceStorage(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IMarketplaceStorage{IMarketplaceStorageCaller: IMarketplaceStorageCaller{contract: contract}, IMarketplaceStorageTransactor: IMarketplaceStorageTransactor{contract: contract}, IMarketplaceStorageFilterer: IMarketplaceStorageFilterer{contract: contract}}, nil
}

// NewIMarketplaceStorageCaller creates a new read-only instance of IMarketplaceStorage, bound to a specific deployed contract.
func NewIMarketplaceStorageCaller(address common.Address, caller bind.ContractCaller) (*IMarketplaceStorageCaller, error) {
	contract, err := bindIMarketplaceStorage(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IMarketplaceStorageCaller{contract: contract}, nil
}

// NewIMarketplaceStorageTransactor creates a new write-only instance of IMarketplaceStorage, bound to a specific deployed contract.
func NewIMarketplaceStorageTransactor(address common.Address, transactor bind.ContractTransactor) (*IMarketplaceStorageTransactor, error) {
	contract, err := bindIMarketplaceStorage(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IMarketplaceStorageTransactor{contract: contract}, nil
}

// NewIMarketplaceStorageFilterer creates a new log filterer instance of IMarketplaceStorage, bound to a specific deployed contract.
func NewIMarketplaceStorageFilterer(address common.Address, filterer bind.ContractFilterer) (*IMarketplaceStorageFilterer, error) {
	contract, err := bindIMarketplaceStorage(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IMarketplaceStorageFilterer{contract: contract}, nil
}

// bindIMarketplaceStorage binds a generic wrapper to an already deployed contract.
func bindIMarketplaceStorage(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IMarketplaceStorageABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IMarketplaceStorage *IMarketplaceStorageRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IMarketplaceStorage.Contract.IMarketplaceStorageCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IMarketplaceStorage *IMarketplaceStorageRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IMarketplaceStorage.Contract.IMarketplaceStorageTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IMarketplaceStorage *IMarketplaceStorageRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IMarketplaceStorage.Contract.IMarketplaceStorageTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IMarketplaceStorage *IMarketplaceStorageCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IMarketplaceStorage.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IMarketplaceStorage *IMarketplaceStorageTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IMarketplaceStorage.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IMarketplaceStorage *IMarketplaceStorageTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IMarketplaceStorage.Contract.contract.Transact(opts, method, params...)
}

// AssetIsAvailable is a free data retrieval call binding the contract method 0x0a7f8977.
//
// Solidity: function assetIsAvailable(bytes32 nftAsset) view returns(bool)
func (_IMarketplaceStorage *IMarketplaceStorageCaller) AssetIsAvailable(opts *bind.CallOpts, nftAsset [32]byte) (bool, error) {
	var out []interface{}
	err := _IMarketplaceStorage.contract.Call(opts, &out, "assetIsAvailable", nftAsset)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// AssetIsAvailable is a free data retrieval call binding the contract method 0x0a7f8977.
//
// Solidity: function assetIsAvailable(bytes32 nftAsset) view returns(bool)
func (_IMarketplaceStorage *IMarketplaceStorageSession) AssetIsAvailable(nftAsset [32]byte) (bool, error) {
	return _IMarketplaceStorage.Contract.AssetIsAvailable(&_IMarketplaceStorage.CallOpts, nftAsset)
}

// AssetIsAvailable is a free data retrieval call binding the contract method 0x0a7f8977.
//
// Solidity: function assetIsAvailable(bytes32 nftAsset) view returns(bool)
func (_IMarketplaceStorage *IMarketplaceStorageCallerSession) AssetIsAvailable(nftAsset [32]byte) (bool, error) {
	return _IMarketplaceStorage.Contract.AssetIsAvailable(&_IMarketplaceStorage.CallOpts, nftAsset)
}

// AuctionIsEnded is a free data retrieval call binding the contract method 0xc064bc50.
//
// Solidity: function auctionIsEnded(bytes32 nftAsset, bytes32 auctionId) view returns(bool ended)
func (_IMarketplaceStorage *IMarketplaceStorageCaller) AuctionIsEnded(opts *bind.CallOpts, nftAsset [32]byte, auctionId [32]byte) (bool, error) {
	var out []interface{}
	err := _IMarketplaceStorage.contract.Call(opts, &out, "auctionIsEnded", nftAsset, auctionId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// AuctionIsEnded is a free data retrieval call binding the contract method 0xc064bc50.
//
// Solidity: function auctionIsEnded(bytes32 nftAsset, bytes32 auctionId) view returns(bool ended)
func (_IMarketplaceStorage *IMarketplaceStorageSession) AuctionIsEnded(nftAsset [32]byte, auctionId [32]byte) (bool, error) {
	return _IMarketplaceStorage.Contract.AuctionIsEnded(&_IMarketplaceStorage.CallOpts, nftAsset, auctionId)
}

// AuctionIsEnded is a free data retrieval call binding the contract method 0xc064bc50.
//
// Solidity: function auctionIsEnded(bytes32 nftAsset, bytes32 auctionId) view returns(bool ended)
func (_IMarketplaceStorage *IMarketplaceStorageCallerSession) AuctionIsEnded(nftAsset [32]byte, auctionId [32]byte) (bool, error) {
	return _IMarketplaceStorage.Contract.AuctionIsEnded(&_IMarketplaceStorage.CallOpts, nftAsset, auctionId)
}

// AuctionIsExisted is a free data retrieval call binding the contract method 0x1a8b4169.
//
// Solidity: function auctionIsExisted(bytes32 auctionId) view returns(bool existed)
func (_IMarketplaceStorage *IMarketplaceStorageCaller) AuctionIsExisted(opts *bind.CallOpts, auctionId [32]byte) (bool, error) {
	var out []interface{}
	err := _IMarketplaceStorage.contract.Call(opts, &out, "auctionIsExisted", auctionId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// AuctionIsExisted is a free data retrieval call binding the contract method 0x1a8b4169.
//
// Solidity: function auctionIsExisted(bytes32 auctionId) view returns(bool existed)
func (_IMarketplaceStorage *IMarketplaceStorageSession) AuctionIsExisted(auctionId [32]byte) (bool, error) {
	return _IMarketplaceStorage.Contract.AuctionIsExisted(&_IMarketplaceStorage.CallOpts, auctionId)
}

// AuctionIsExisted is a free data retrieval call binding the contract method 0x1a8b4169.
//
// Solidity: function auctionIsExisted(bytes32 auctionId) view returns(bool existed)
func (_IMarketplaceStorage *IMarketplaceStorageCallerSession) AuctionIsExisted(auctionId [32]byte) (bool, error) {
	return _IMarketplaceStorage.Contract.AuctionIsExisted(&_IMarketplaceStorage.CallOpts, auctionId)
}

// AuctionIsRunning is a free data retrieval call binding the contract method 0xbc66c29f.
//
// Solidity: function auctionIsRunning(bytes32 nftAsset, bytes32 auctionId) view returns(bool existed)
func (_IMarketplaceStorage *IMarketplaceStorageCaller) AuctionIsRunning(opts *bind.CallOpts, nftAsset [32]byte, auctionId [32]byte) (bool, error) {
	var out []interface{}
	err := _IMarketplaceStorage.contract.Call(opts, &out, "auctionIsRunning", nftAsset, auctionId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// AuctionIsRunning is a free data retrieval call binding the contract method 0xbc66c29f.
//
// Solidity: function auctionIsRunning(bytes32 nftAsset, bytes32 auctionId) view returns(bool existed)
func (_IMarketplaceStorage *IMarketplaceStorageSession) AuctionIsRunning(nftAsset [32]byte, auctionId [32]byte) (bool, error) {
	return _IMarketplaceStorage.Contract.AuctionIsRunning(&_IMarketplaceStorage.CallOpts, nftAsset, auctionId)
}

// AuctionIsRunning is a free data retrieval call binding the contract method 0xbc66c29f.
//
// Solidity: function auctionIsRunning(bytes32 nftAsset, bytes32 auctionId) view returns(bool existed)
func (_IMarketplaceStorage *IMarketplaceStorageCallerSession) AuctionIsRunning(nftAsset [32]byte, auctionId [32]byte) (bool, error) {
	return _IMarketplaceStorage.Contract.AuctionIsRunning(&_IMarketplaceStorage.CallOpts, nftAsset, auctionId)
}

// GetAuction is a free data retrieval call binding the contract method 0x15924b5b.
//
// Solidity: function getAuction(bytes32 auctionId) view returns((address,address,bytes32,uint256,uint256,uint256,uint256) order)
func (_IMarketplaceStorage *IMarketplaceStorageCaller) GetAuction(opts *bind.CallOpts, auctionId [32]byte) (IAuctionAuction, error) {
	var out []interface{}
	err := _IMarketplaceStorage.contract.Call(opts, &out, "getAuction", auctionId)

	if err != nil {
		return *new(IAuctionAuction), err
	}

	out0 := *abi.ConvertType(out[0], new(IAuctionAuction)).(*IAuctionAuction)

	return out0, err

}

// GetAuction is a free data retrieval call binding the contract method 0x15924b5b.
//
// Solidity: function getAuction(bytes32 auctionId) view returns((address,address,bytes32,uint256,uint256,uint256,uint256) order)
func (_IMarketplaceStorage *IMarketplaceStorageSession) GetAuction(auctionId [32]byte) (IAuctionAuction, error) {
	return _IMarketplaceStorage.Contract.GetAuction(&_IMarketplaceStorage.CallOpts, auctionId)
}

// GetAuction is a free data retrieval call binding the contract method 0x15924b5b.
//
// Solidity: function getAuction(bytes32 auctionId) view returns((address,address,bytes32,uint256,uint256,uint256,uint256) order)
func (_IMarketplaceStorage *IMarketplaceStorageCallerSession) GetAuction(auctionId [32]byte) (IAuctionAuction, error) {
	return _IMarketplaceStorage.Contract.GetAuction(&_IMarketplaceStorage.CallOpts, auctionId)
}

// GetOrder is a free data retrieval call binding the contract method 0x5778472a.
//
// Solidity: function getOrder(bytes32 nftAsset) view returns((bytes32,address,address,uint256,uint256) order)
func (_IMarketplaceStorage *IMarketplaceStorageCaller) GetOrder(opts *bind.CallOpts, nftAsset [32]byte) (IOrderOrder, error) {
	var out []interface{}
	err := _IMarketplaceStorage.contract.Call(opts, &out, "getOrder", nftAsset)

	if err != nil {
		return *new(IOrderOrder), err
	}

	out0 := *abi.ConvertType(out[0], new(IOrderOrder)).(*IOrderOrder)

	return out0, err

}

// GetOrder is a free data retrieval call binding the contract method 0x5778472a.
//
// Solidity: function getOrder(bytes32 nftAsset) view returns((bytes32,address,address,uint256,uint256) order)
func (_IMarketplaceStorage *IMarketplaceStorageSession) GetOrder(nftAsset [32]byte) (IOrderOrder, error) {
	return _IMarketplaceStorage.Contract.GetOrder(&_IMarketplaceStorage.CallOpts, nftAsset)
}

// GetOrder is a free data retrieval call binding the contract method 0x5778472a.
//
// Solidity: function getOrder(bytes32 nftAsset) view returns((bytes32,address,address,uint256,uint256) order)
func (_IMarketplaceStorage *IMarketplaceStorageCallerSession) GetOrder(nftAsset [32]byte) (IOrderOrder, error) {
	return _IMarketplaceStorage.Contract.GetOrder(&_IMarketplaceStorage.CallOpts, nftAsset)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_IMarketplaceStorage *IMarketplaceStorageCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _IMarketplaceStorage.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_IMarketplaceStorage *IMarketplaceStorageSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _IMarketplaceStorage.Contract.SupportsInterface(&_IMarketplaceStorage.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_IMarketplaceStorage *IMarketplaceStorageCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _IMarketplaceStorage.Contract.SupportsInterface(&_IMarketplaceStorage.CallOpts, interfaceId)
}

// AuctionEnded is a paid mutator transaction binding the contract method 0xac84b7bc.
//
// Solidity: function auctionEnded(bytes32 nftAsset) returns()
func (_IMarketplaceStorage *IMarketplaceStorageTransactor) AuctionEnded(opts *bind.TransactOpts, nftAsset [32]byte) (*types.Transaction, error) {
	return _IMarketplaceStorage.contract.Transact(opts, "auctionEnded", nftAsset)
}

// AuctionEnded is a paid mutator transaction binding the contract method 0xac84b7bc.
//
// Solidity: function auctionEnded(bytes32 nftAsset) returns()
func (_IMarketplaceStorage *IMarketplaceStorageSession) AuctionEnded(nftAsset [32]byte) (*types.Transaction, error) {
	return _IMarketplaceStorage.Contract.AuctionEnded(&_IMarketplaceStorage.TransactOpts, nftAsset)
}

// AuctionEnded is a paid mutator transaction binding the contract method 0xac84b7bc.
//
// Solidity: function auctionEnded(bytes32 nftAsset) returns()
func (_IMarketplaceStorage *IMarketplaceStorageTransactorSession) AuctionEnded(nftAsset [32]byte) (*types.Transaction, error) {
	return _IMarketplaceStorage.Contract.AuctionEnded(&_IMarketplaceStorage.TransactOpts, nftAsset)
}

// CreateAuction is a paid mutator transaction binding the contract method 0x84b04aa9.
//
// Solidity: function createAuction(address assetOwner, address nftAddress, uint256 assetId, bytes32 auctionId, uint256 biddingEnd, uint256 revealEnd, uint256 startPriceInWei) returns()
func (_IMarketplaceStorage *IMarketplaceStorageTransactor) CreateAuction(opts *bind.TransactOpts, assetOwner common.Address, nftAddress common.Address, assetId *big.Int, auctionId [32]byte, biddingEnd *big.Int, revealEnd *big.Int, startPriceInWei *big.Int) (*types.Transaction, error) {
	return _IMarketplaceStorage.contract.Transact(opts, "createAuction", assetOwner, nftAddress, assetId, auctionId, biddingEnd, revealEnd, startPriceInWei)
}

// CreateAuction is a paid mutator transaction binding the contract method 0x84b04aa9.
//
// Solidity: function createAuction(address assetOwner, address nftAddress, uint256 assetId, bytes32 auctionId, uint256 biddingEnd, uint256 revealEnd, uint256 startPriceInWei) returns()
func (_IMarketplaceStorage *IMarketplaceStorageSession) CreateAuction(assetOwner common.Address, nftAddress common.Address, assetId *big.Int, auctionId [32]byte, biddingEnd *big.Int, revealEnd *big.Int, startPriceInWei *big.Int) (*types.Transaction, error) {
	return _IMarketplaceStorage.Contract.CreateAuction(&_IMarketplaceStorage.TransactOpts, assetOwner, nftAddress, assetId, auctionId, biddingEnd, revealEnd, startPriceInWei)
}

// CreateAuction is a paid mutator transaction binding the contract method 0x84b04aa9.
//
// Solidity: function createAuction(address assetOwner, address nftAddress, uint256 assetId, bytes32 auctionId, uint256 biddingEnd, uint256 revealEnd, uint256 startPriceInWei) returns()
func (_IMarketplaceStorage *IMarketplaceStorageTransactorSession) CreateAuction(assetOwner common.Address, nftAddress common.Address, assetId *big.Int, auctionId [32]byte, biddingEnd *big.Int, revealEnd *big.Int, startPriceInWei *big.Int) (*types.Transaction, error) {
	return _IMarketplaceStorage.Contract.CreateAuction(&_IMarketplaceStorage.TransactOpts, assetOwner, nftAddress, assetId, auctionId, biddingEnd, revealEnd, startPriceInWei)
}

// CreateOrder is a paid mutator transaction binding the contract method 0x930ca094.
//
// Solidity: function createOrder(address seller, address nftAddress, uint256 assetId, bytes32 orderId, uint256 priceInWei, uint256 expiredAt) returns()
func (_IMarketplaceStorage *IMarketplaceStorageTransactor) CreateOrder(opts *bind.TransactOpts, seller common.Address, nftAddress common.Address, assetId *big.Int, orderId [32]byte, priceInWei *big.Int, expiredAt *big.Int) (*types.Transaction, error) {
	return _IMarketplaceStorage.contract.Transact(opts, "createOrder", seller, nftAddress, assetId, orderId, priceInWei, expiredAt)
}

// CreateOrder is a paid mutator transaction binding the contract method 0x930ca094.
//
// Solidity: function createOrder(address seller, address nftAddress, uint256 assetId, bytes32 orderId, uint256 priceInWei, uint256 expiredAt) returns()
func (_IMarketplaceStorage *IMarketplaceStorageSession) CreateOrder(seller common.Address, nftAddress common.Address, assetId *big.Int, orderId [32]byte, priceInWei *big.Int, expiredAt *big.Int) (*types.Transaction, error) {
	return _IMarketplaceStorage.Contract.CreateOrder(&_IMarketplaceStorage.TransactOpts, seller, nftAddress, assetId, orderId, priceInWei, expiredAt)
}

// CreateOrder is a paid mutator transaction binding the contract method 0x930ca094.
//
// Solidity: function createOrder(address seller, address nftAddress, uint256 assetId, bytes32 orderId, uint256 priceInWei, uint256 expiredAt) returns()
func (_IMarketplaceStorage *IMarketplaceStorageTransactorSession) CreateOrder(seller common.Address, nftAddress common.Address, assetId *big.Int, orderId [32]byte, priceInWei *big.Int, expiredAt *big.Int) (*types.Transaction, error) {
	return _IMarketplaceStorage.Contract.CreateOrder(&_IMarketplaceStorage.TransactOpts, seller, nftAddress, assetId, orderId, priceInWei, expiredAt)
}

// DeleteOrder is a paid mutator transaction binding the contract method 0x87a61cbd.
//
// Solidity: function deleteOrder(bytes32 nftAsset) returns()
func (_IMarketplaceStorage *IMarketplaceStorageTransactor) DeleteOrder(opts *bind.TransactOpts, nftAsset [32]byte) (*types.Transaction, error) {
	return _IMarketplaceStorage.contract.Transact(opts, "deleteOrder", nftAsset)
}

// DeleteOrder is a paid mutator transaction binding the contract method 0x87a61cbd.
//
// Solidity: function deleteOrder(bytes32 nftAsset) returns()
func (_IMarketplaceStorage *IMarketplaceStorageSession) DeleteOrder(nftAsset [32]byte) (*types.Transaction, error) {
	return _IMarketplaceStorage.Contract.DeleteOrder(&_IMarketplaceStorage.TransactOpts, nftAsset)
}

// DeleteOrder is a paid mutator transaction binding the contract method 0x87a61cbd.
//
// Solidity: function deleteOrder(bytes32 nftAsset) returns()
func (_IMarketplaceStorage *IMarketplaceStorageTransactorSession) DeleteOrder(nftAsset [32]byte) (*types.Transaction, error) {
	return _IMarketplaceStorage.Contract.DeleteOrder(&_IMarketplaceStorage.TransactOpts, nftAsset)
}

// UpdateHighestBid is a paid mutator transaction binding the contract method 0x572f7d5e.
//
// Solidity: function updateHighestBid(address bidder, uint256 highestBid, bytes32 auctionId) returns()
func (_IMarketplaceStorage *IMarketplaceStorageTransactor) UpdateHighestBid(opts *bind.TransactOpts, bidder common.Address, highestBid *big.Int, auctionId [32]byte) (*types.Transaction, error) {
	return _IMarketplaceStorage.contract.Transact(opts, "updateHighestBid", bidder, highestBid, auctionId)
}

// UpdateHighestBid is a paid mutator transaction binding the contract method 0x572f7d5e.
//
// Solidity: function updateHighestBid(address bidder, uint256 highestBid, bytes32 auctionId) returns()
func (_IMarketplaceStorage *IMarketplaceStorageSession) UpdateHighestBid(bidder common.Address, highestBid *big.Int, auctionId [32]byte) (*types.Transaction, error) {
	return _IMarketplaceStorage.Contract.UpdateHighestBid(&_IMarketplaceStorage.TransactOpts, bidder, highestBid, auctionId)
}

// UpdateHighestBid is a paid mutator transaction binding the contract method 0x572f7d5e.
//
// Solidity: function updateHighestBid(address bidder, uint256 highestBid, bytes32 auctionId) returns()
func (_IMarketplaceStorage *IMarketplaceStorageTransactorSession) UpdateHighestBid(bidder common.Address, highestBid *big.Int, auctionId [32]byte) (*types.Transaction, error) {
	return _IMarketplaceStorage.Contract.UpdateHighestBid(&_IMarketplaceStorage.TransactOpts, bidder, highestBid, auctionId)
}

// IMarketplaceStorageAuctionCancelledIterator is returned from FilterAuctionCancelled and is used to iterate over the raw logs and unpacked data for AuctionCancelled events raised by the IMarketplaceStorage contract.
type IMarketplaceStorageAuctionCancelledIterator struct {
	Event *IMarketplaceStorageAuctionCancelled // Event containing the contract specifics and raw log

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
func (it *IMarketplaceStorageAuctionCancelledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IMarketplaceStorageAuctionCancelled)
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
		it.Event = new(IMarketplaceStorageAuctionCancelled)
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
func (it *IMarketplaceStorageAuctionCancelledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IMarketplaceStorageAuctionCancelledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IMarketplaceStorageAuctionCancelled represents a AuctionCancelled event raised by the IMarketplaceStorage contract.
type IMarketplaceStorageAuctionCancelled struct {
	AuctionId [32]byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterAuctionCancelled is a free log retrieval operation binding the contract event 0xd88fab4b08bf76f15cf2d6e03e382acf1edd6790ab82967e406abac37db20288.
//
// Solidity: event AuctionCancelled(bytes32 indexed auctionId)
func (_IMarketplaceStorage *IMarketplaceStorageFilterer) FilterAuctionCancelled(opts *bind.FilterOpts, auctionId [][32]byte) (*IMarketplaceStorageAuctionCancelledIterator, error) {

	var auctionIdRule []interface{}
	for _, auctionIdItem := range auctionId {
		auctionIdRule = append(auctionIdRule, auctionIdItem)
	}

	logs, sub, err := _IMarketplaceStorage.contract.FilterLogs(opts, "AuctionCancelled", auctionIdRule)
	if err != nil {
		return nil, err
	}
	return &IMarketplaceStorageAuctionCancelledIterator{contract: _IMarketplaceStorage.contract, event: "AuctionCancelled", logs: logs, sub: sub}, nil
}

// WatchAuctionCancelled is a free log subscription operation binding the contract event 0xd88fab4b08bf76f15cf2d6e03e382acf1edd6790ab82967e406abac37db20288.
//
// Solidity: event AuctionCancelled(bytes32 indexed auctionId)
func (_IMarketplaceStorage *IMarketplaceStorageFilterer) WatchAuctionCancelled(opts *bind.WatchOpts, sink chan<- *IMarketplaceStorageAuctionCancelled, auctionId [][32]byte) (event.Subscription, error) {

	var auctionIdRule []interface{}
	for _, auctionIdItem := range auctionId {
		auctionIdRule = append(auctionIdRule, auctionIdItem)
	}

	logs, sub, err := _IMarketplaceStorage.contract.WatchLogs(opts, "AuctionCancelled", auctionIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IMarketplaceStorageAuctionCancelled)
				if err := _IMarketplaceStorage.contract.UnpackLog(event, "AuctionCancelled", log); err != nil {
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

// ParseAuctionCancelled is a log parse operation binding the contract event 0xd88fab4b08bf76f15cf2d6e03e382acf1edd6790ab82967e406abac37db20288.
//
// Solidity: event AuctionCancelled(bytes32 indexed auctionId)
func (_IMarketplaceStorage *IMarketplaceStorageFilterer) ParseAuctionCancelled(log types.Log) (*IMarketplaceStorageAuctionCancelled, error) {
	event := new(IMarketplaceStorageAuctionCancelled)
	if err := _IMarketplaceStorage.contract.UnpackLog(event, "AuctionCancelled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IMarketplaceStorageAuctionCreatedIterator is returned from FilterAuctionCreated and is used to iterate over the raw logs and unpacked data for AuctionCreated events raised by the IMarketplaceStorage contract.
type IMarketplaceStorageAuctionCreatedIterator struct {
	Event *IMarketplaceStorageAuctionCreated // Event containing the contract specifics and raw log

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
func (it *IMarketplaceStorageAuctionCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IMarketplaceStorageAuctionCreated)
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
		it.Event = new(IMarketplaceStorageAuctionCreated)
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
func (it *IMarketplaceStorageAuctionCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IMarketplaceStorageAuctionCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IMarketplaceStorageAuctionCreated represents a AuctionCreated event raised by the IMarketplaceStorage contract.
type IMarketplaceStorageAuctionCreated struct {
	Seller          common.Address
	NftAddress      common.Address
	AuctionId       [32]byte
	AssetId         *big.Int
	BiddingEnd      *big.Int
	RevealEnd       *big.Int
	StartPriceInWei *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterAuctionCreated is a free log retrieval operation binding the contract event 0xaf4bad306f14d5c908e5b871fa54296ed7d4f29b8092bf5062de6813c247e54b.
//
// Solidity: event AuctionCreated(address indexed seller, address nftAddress, bytes32 indexed auctionId, uint256 assetId, uint256 biddingEnd, uint256 revealEnd, uint256 startPriceInWei)
func (_IMarketplaceStorage *IMarketplaceStorageFilterer) FilterAuctionCreated(opts *bind.FilterOpts, seller []common.Address, auctionId [][32]byte) (*IMarketplaceStorageAuctionCreatedIterator, error) {

	var sellerRule []interface{}
	for _, sellerItem := range seller {
		sellerRule = append(sellerRule, sellerItem)
	}

	var auctionIdRule []interface{}
	for _, auctionIdItem := range auctionId {
		auctionIdRule = append(auctionIdRule, auctionIdItem)
	}

	logs, sub, err := _IMarketplaceStorage.contract.FilterLogs(opts, "AuctionCreated", sellerRule, auctionIdRule)
	if err != nil {
		return nil, err
	}
	return &IMarketplaceStorageAuctionCreatedIterator{contract: _IMarketplaceStorage.contract, event: "AuctionCreated", logs: logs, sub: sub}, nil
}

// WatchAuctionCreated is a free log subscription operation binding the contract event 0xaf4bad306f14d5c908e5b871fa54296ed7d4f29b8092bf5062de6813c247e54b.
//
// Solidity: event AuctionCreated(address indexed seller, address nftAddress, bytes32 indexed auctionId, uint256 assetId, uint256 biddingEnd, uint256 revealEnd, uint256 startPriceInWei)
func (_IMarketplaceStorage *IMarketplaceStorageFilterer) WatchAuctionCreated(opts *bind.WatchOpts, sink chan<- *IMarketplaceStorageAuctionCreated, seller []common.Address, auctionId [][32]byte) (event.Subscription, error) {

	var sellerRule []interface{}
	for _, sellerItem := range seller {
		sellerRule = append(sellerRule, sellerItem)
	}

	var auctionIdRule []interface{}
	for _, auctionIdItem := range auctionId {
		auctionIdRule = append(auctionIdRule, auctionIdItem)
	}

	logs, sub, err := _IMarketplaceStorage.contract.WatchLogs(opts, "AuctionCreated", sellerRule, auctionIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IMarketplaceStorageAuctionCreated)
				if err := _IMarketplaceStorage.contract.UnpackLog(event, "AuctionCreated", log); err != nil {
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

// ParseAuctionCreated is a log parse operation binding the contract event 0xaf4bad306f14d5c908e5b871fa54296ed7d4f29b8092bf5062de6813c247e54b.
//
// Solidity: event AuctionCreated(address indexed seller, address nftAddress, bytes32 indexed auctionId, uint256 assetId, uint256 biddingEnd, uint256 revealEnd, uint256 startPriceInWei)
func (_IMarketplaceStorage *IMarketplaceStorageFilterer) ParseAuctionCreated(log types.Log) (*IMarketplaceStorageAuctionCreated, error) {
	event := new(IMarketplaceStorageAuctionCreated)
	if err := _IMarketplaceStorage.contract.UnpackLog(event, "AuctionCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IMarketplaceStorageAuctionEndedIterator is returned from FilterAuctionEnded and is used to iterate over the raw logs and unpacked data for AuctionEnded events raised by the IMarketplaceStorage contract.
type IMarketplaceStorageAuctionEndedIterator struct {
	Event *IMarketplaceStorageAuctionEnded // Event containing the contract specifics and raw log

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
func (it *IMarketplaceStorageAuctionEndedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IMarketplaceStorageAuctionEnded)
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
		it.Event = new(IMarketplaceStorageAuctionEnded)
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
func (it *IMarketplaceStorageAuctionEndedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IMarketplaceStorageAuctionEndedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IMarketplaceStorageAuctionEnded represents a AuctionEnded event raised by the IMarketplaceStorage contract.
type IMarketplaceStorageAuctionEnded struct {
	AuctionId [32]byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterAuctionEnded is a free log retrieval operation binding the contract event 0x447d0298e6e2716a4343195759c35b277a60b76e628284d93262fa9869db9921.
//
// Solidity: event AuctionEnded(bytes32 indexed auctionId)
func (_IMarketplaceStorage *IMarketplaceStorageFilterer) FilterAuctionEnded(opts *bind.FilterOpts, auctionId [][32]byte) (*IMarketplaceStorageAuctionEndedIterator, error) {

	var auctionIdRule []interface{}
	for _, auctionIdItem := range auctionId {
		auctionIdRule = append(auctionIdRule, auctionIdItem)
	}

	logs, sub, err := _IMarketplaceStorage.contract.FilterLogs(opts, "AuctionEnded", auctionIdRule)
	if err != nil {
		return nil, err
	}
	return &IMarketplaceStorageAuctionEndedIterator{contract: _IMarketplaceStorage.contract, event: "AuctionEnded", logs: logs, sub: sub}, nil
}

// WatchAuctionEnded is a free log subscription operation binding the contract event 0x447d0298e6e2716a4343195759c35b277a60b76e628284d93262fa9869db9921.
//
// Solidity: event AuctionEnded(bytes32 indexed auctionId)
func (_IMarketplaceStorage *IMarketplaceStorageFilterer) WatchAuctionEnded(opts *bind.WatchOpts, sink chan<- *IMarketplaceStorageAuctionEnded, auctionId [][32]byte) (event.Subscription, error) {

	var auctionIdRule []interface{}
	for _, auctionIdItem := range auctionId {
		auctionIdRule = append(auctionIdRule, auctionIdItem)
	}

	logs, sub, err := _IMarketplaceStorage.contract.WatchLogs(opts, "AuctionEnded", auctionIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IMarketplaceStorageAuctionEnded)
				if err := _IMarketplaceStorage.contract.UnpackLog(event, "AuctionEnded", log); err != nil {
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

// ParseAuctionEnded is a log parse operation binding the contract event 0x447d0298e6e2716a4343195759c35b277a60b76e628284d93262fa9869db9921.
//
// Solidity: event AuctionEnded(bytes32 indexed auctionId)
func (_IMarketplaceStorage *IMarketplaceStorageFilterer) ParseAuctionEnded(log types.Log) (*IMarketplaceStorageAuctionEnded, error) {
	event := new(IMarketplaceStorageAuctionEnded)
	if err := _IMarketplaceStorage.contract.UnpackLog(event, "AuctionEnded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IMarketplaceStorageAuctionRefundIterator is returned from FilterAuctionRefund and is used to iterate over the raw logs and unpacked data for AuctionRefund events raised by the IMarketplaceStorage contract.
type IMarketplaceStorageAuctionRefundIterator struct {
	Event *IMarketplaceStorageAuctionRefund // Event containing the contract specifics and raw log

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
func (it *IMarketplaceStorageAuctionRefundIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IMarketplaceStorageAuctionRefund)
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
		it.Event = new(IMarketplaceStorageAuctionRefund)
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
func (it *IMarketplaceStorageAuctionRefundIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IMarketplaceStorageAuctionRefundIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IMarketplaceStorageAuctionRefund represents a AuctionRefund event raised by the IMarketplaceStorage contract.
type IMarketplaceStorageAuctionRefund struct {
	Bidder    common.Address
	AuctionId [32]byte
	Deposit   *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterAuctionRefund is a free log retrieval operation binding the contract event 0x6be22eb0241a0575f8b866ff176552cf71078977d6d30e501dd2d326c8d57114.
//
// Solidity: event AuctionRefund(address indexed bidder, bytes32 auctionId, uint256 deposit)
func (_IMarketplaceStorage *IMarketplaceStorageFilterer) FilterAuctionRefund(opts *bind.FilterOpts, bidder []common.Address) (*IMarketplaceStorageAuctionRefundIterator, error) {

	var bidderRule []interface{}
	for _, bidderItem := range bidder {
		bidderRule = append(bidderRule, bidderItem)
	}

	logs, sub, err := _IMarketplaceStorage.contract.FilterLogs(opts, "AuctionRefund", bidderRule)
	if err != nil {
		return nil, err
	}
	return &IMarketplaceStorageAuctionRefundIterator{contract: _IMarketplaceStorage.contract, event: "AuctionRefund", logs: logs, sub: sub}, nil
}

// WatchAuctionRefund is a free log subscription operation binding the contract event 0x6be22eb0241a0575f8b866ff176552cf71078977d6d30e501dd2d326c8d57114.
//
// Solidity: event AuctionRefund(address indexed bidder, bytes32 auctionId, uint256 deposit)
func (_IMarketplaceStorage *IMarketplaceStorageFilterer) WatchAuctionRefund(opts *bind.WatchOpts, sink chan<- *IMarketplaceStorageAuctionRefund, bidder []common.Address) (event.Subscription, error) {

	var bidderRule []interface{}
	for _, bidderItem := range bidder {
		bidderRule = append(bidderRule, bidderItem)
	}

	logs, sub, err := _IMarketplaceStorage.contract.WatchLogs(opts, "AuctionRefund", bidderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IMarketplaceStorageAuctionRefund)
				if err := _IMarketplaceStorage.contract.UnpackLog(event, "AuctionRefund", log); err != nil {
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

// ParseAuctionRefund is a log parse operation binding the contract event 0x6be22eb0241a0575f8b866ff176552cf71078977d6d30e501dd2d326c8d57114.
//
// Solidity: event AuctionRefund(address indexed bidder, bytes32 auctionId, uint256 deposit)
func (_IMarketplaceStorage *IMarketplaceStorageFilterer) ParseAuctionRefund(log types.Log) (*IMarketplaceStorageAuctionRefund, error) {
	event := new(IMarketplaceStorageAuctionRefund)
	if err := _IMarketplaceStorage.contract.UnpackLog(event, "AuctionRefund", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IMarketplaceStorageAuctionSuccessfulIterator is returned from FilterAuctionSuccessful and is used to iterate over the raw logs and unpacked data for AuctionSuccessful events raised by the IMarketplaceStorage contract.
type IMarketplaceStorageAuctionSuccessfulIterator struct {
	Event *IMarketplaceStorageAuctionSuccessful // Event containing the contract specifics and raw log

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
func (it *IMarketplaceStorageAuctionSuccessfulIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IMarketplaceStorageAuctionSuccessful)
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
		it.Event = new(IMarketplaceStorageAuctionSuccessful)
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
func (it *IMarketplaceStorageAuctionSuccessfulIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IMarketplaceStorageAuctionSuccessfulIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IMarketplaceStorageAuctionSuccessful represents a AuctionSuccessful event raised by the IMarketplaceStorage contract.
type IMarketplaceStorageAuctionSuccessful struct {
	Seller     common.Address
	Buyer      common.Address
	AuctionId  [32]byte
	TotalPrice *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterAuctionSuccessful is a free log retrieval operation binding the contract event 0x5ed769e966210f5d06d68e75831871131255cee237cb8edf78565aae170475ea.
//
// Solidity: event AuctionSuccessful(address indexed seller, address indexed buyer, bytes32 indexed auctionId, uint256 totalPrice)
func (_IMarketplaceStorage *IMarketplaceStorageFilterer) FilterAuctionSuccessful(opts *bind.FilterOpts, seller []common.Address, buyer []common.Address, auctionId [][32]byte) (*IMarketplaceStorageAuctionSuccessfulIterator, error) {

	var sellerRule []interface{}
	for _, sellerItem := range seller {
		sellerRule = append(sellerRule, sellerItem)
	}
	var buyerRule []interface{}
	for _, buyerItem := range buyer {
		buyerRule = append(buyerRule, buyerItem)
	}
	var auctionIdRule []interface{}
	for _, auctionIdItem := range auctionId {
		auctionIdRule = append(auctionIdRule, auctionIdItem)
	}

	logs, sub, err := _IMarketplaceStorage.contract.FilterLogs(opts, "AuctionSuccessful", sellerRule, buyerRule, auctionIdRule)
	if err != nil {
		return nil, err
	}
	return &IMarketplaceStorageAuctionSuccessfulIterator{contract: _IMarketplaceStorage.contract, event: "AuctionSuccessful", logs: logs, sub: sub}, nil
}

// WatchAuctionSuccessful is a free log subscription operation binding the contract event 0x5ed769e966210f5d06d68e75831871131255cee237cb8edf78565aae170475ea.
//
// Solidity: event AuctionSuccessful(address indexed seller, address indexed buyer, bytes32 indexed auctionId, uint256 totalPrice)
func (_IMarketplaceStorage *IMarketplaceStorageFilterer) WatchAuctionSuccessful(opts *bind.WatchOpts, sink chan<- *IMarketplaceStorageAuctionSuccessful, seller []common.Address, buyer []common.Address, auctionId [][32]byte) (event.Subscription, error) {

	var sellerRule []interface{}
	for _, sellerItem := range seller {
		sellerRule = append(sellerRule, sellerItem)
	}
	var buyerRule []interface{}
	for _, buyerItem := range buyer {
		buyerRule = append(buyerRule, buyerItem)
	}
	var auctionIdRule []interface{}
	for _, auctionIdItem := range auctionId {
		auctionIdRule = append(auctionIdRule, auctionIdItem)
	}

	logs, sub, err := _IMarketplaceStorage.contract.WatchLogs(opts, "AuctionSuccessful", sellerRule, buyerRule, auctionIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IMarketplaceStorageAuctionSuccessful)
				if err := _IMarketplaceStorage.contract.UnpackLog(event, "AuctionSuccessful", log); err != nil {
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

// ParseAuctionSuccessful is a log parse operation binding the contract event 0x5ed769e966210f5d06d68e75831871131255cee237cb8edf78565aae170475ea.
//
// Solidity: event AuctionSuccessful(address indexed seller, address indexed buyer, bytes32 indexed auctionId, uint256 totalPrice)
func (_IMarketplaceStorage *IMarketplaceStorageFilterer) ParseAuctionSuccessful(log types.Log) (*IMarketplaceStorageAuctionSuccessful, error) {
	event := new(IMarketplaceStorageAuctionSuccessful)
	if err := _IMarketplaceStorage.contract.UnpackLog(event, "AuctionSuccessful", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IMarketplaceStorageBidSuccessfulIterator is returned from FilterBidSuccessful and is used to iterate over the raw logs and unpacked data for BidSuccessful events raised by the IMarketplaceStorage contract.
type IMarketplaceStorageBidSuccessfulIterator struct {
	Event *IMarketplaceStorageBidSuccessful // Event containing the contract specifics and raw log

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
func (it *IMarketplaceStorageBidSuccessfulIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IMarketplaceStorageBidSuccessful)
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
		it.Event = new(IMarketplaceStorageBidSuccessful)
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
func (it *IMarketplaceStorageBidSuccessfulIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IMarketplaceStorageBidSuccessfulIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IMarketplaceStorageBidSuccessful represents a BidSuccessful event raised by the IMarketplaceStorage contract.
type IMarketplaceStorageBidSuccessful struct {
	Bidder     common.Address
	AuctionId  [32]byte
	BlindedBid [32]byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterBidSuccessful is a free log retrieval operation binding the contract event 0x327e5a459c964bcfa4f6d71b32a04a53b0002417b0924a8057d9f12e688e0bec.
//
// Solidity: event BidSuccessful(address indexed bidder, bytes32 indexed auctionId, bytes32 blindedBid)
func (_IMarketplaceStorage *IMarketplaceStorageFilterer) FilterBidSuccessful(opts *bind.FilterOpts, bidder []common.Address, auctionId [][32]byte) (*IMarketplaceStorageBidSuccessfulIterator, error) {

	var bidderRule []interface{}
	for _, bidderItem := range bidder {
		bidderRule = append(bidderRule, bidderItem)
	}
	var auctionIdRule []interface{}
	for _, auctionIdItem := range auctionId {
		auctionIdRule = append(auctionIdRule, auctionIdItem)
	}

	logs, sub, err := _IMarketplaceStorage.contract.FilterLogs(opts, "BidSuccessful", bidderRule, auctionIdRule)
	if err != nil {
		return nil, err
	}
	return &IMarketplaceStorageBidSuccessfulIterator{contract: _IMarketplaceStorage.contract, event: "BidSuccessful", logs: logs, sub: sub}, nil
}

// WatchBidSuccessful is a free log subscription operation binding the contract event 0x327e5a459c964bcfa4f6d71b32a04a53b0002417b0924a8057d9f12e688e0bec.
//
// Solidity: event BidSuccessful(address indexed bidder, bytes32 indexed auctionId, bytes32 blindedBid)
func (_IMarketplaceStorage *IMarketplaceStorageFilterer) WatchBidSuccessful(opts *bind.WatchOpts, sink chan<- *IMarketplaceStorageBidSuccessful, bidder []common.Address, auctionId [][32]byte) (event.Subscription, error) {

	var bidderRule []interface{}
	for _, bidderItem := range bidder {
		bidderRule = append(bidderRule, bidderItem)
	}
	var auctionIdRule []interface{}
	for _, auctionIdItem := range auctionId {
		auctionIdRule = append(auctionIdRule, auctionIdItem)
	}

	logs, sub, err := _IMarketplaceStorage.contract.WatchLogs(opts, "BidSuccessful", bidderRule, auctionIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IMarketplaceStorageBidSuccessful)
				if err := _IMarketplaceStorage.contract.UnpackLog(event, "BidSuccessful", log); err != nil {
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

// ParseBidSuccessful is a log parse operation binding the contract event 0x327e5a459c964bcfa4f6d71b32a04a53b0002417b0924a8057d9f12e688e0bec.
//
// Solidity: event BidSuccessful(address indexed bidder, bytes32 indexed auctionId, bytes32 blindedBid)
func (_IMarketplaceStorage *IMarketplaceStorageFilterer) ParseBidSuccessful(log types.Log) (*IMarketplaceStorageBidSuccessful, error) {
	event := new(IMarketplaceStorageBidSuccessful)
	if err := _IMarketplaceStorage.contract.UnpackLog(event, "BidSuccessful", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IMarketplaceStorageOrderCancelledIterator is returned from FilterOrderCancelled and is used to iterate over the raw logs and unpacked data for OrderCancelled events raised by the IMarketplaceStorage contract.
type IMarketplaceStorageOrderCancelledIterator struct {
	Event *IMarketplaceStorageOrderCancelled // Event containing the contract specifics and raw log

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
func (it *IMarketplaceStorageOrderCancelledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IMarketplaceStorageOrderCancelled)
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
		it.Event = new(IMarketplaceStorageOrderCancelled)
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
func (it *IMarketplaceStorageOrderCancelledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IMarketplaceStorageOrderCancelledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IMarketplaceStorageOrderCancelled represents a OrderCancelled event raised by the IMarketplaceStorage contract.
type IMarketplaceStorageOrderCancelled struct {
	Id  [32]byte
	Raw types.Log // Blockchain specific contextual infos
}

// FilterOrderCancelled is a free log retrieval operation binding the contract event 0x5152abf959f6564662358c2e52b702259b78bac5ee7842a0f01937e670efcc7d.
//
// Solidity: event OrderCancelled(bytes32 indexed id)
func (_IMarketplaceStorage *IMarketplaceStorageFilterer) FilterOrderCancelled(opts *bind.FilterOpts, id [][32]byte) (*IMarketplaceStorageOrderCancelledIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _IMarketplaceStorage.contract.FilterLogs(opts, "OrderCancelled", idRule)
	if err != nil {
		return nil, err
	}
	return &IMarketplaceStorageOrderCancelledIterator{contract: _IMarketplaceStorage.contract, event: "OrderCancelled", logs: logs, sub: sub}, nil
}

// WatchOrderCancelled is a free log subscription operation binding the contract event 0x5152abf959f6564662358c2e52b702259b78bac5ee7842a0f01937e670efcc7d.
//
// Solidity: event OrderCancelled(bytes32 indexed id)
func (_IMarketplaceStorage *IMarketplaceStorageFilterer) WatchOrderCancelled(opts *bind.WatchOpts, sink chan<- *IMarketplaceStorageOrderCancelled, id [][32]byte) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _IMarketplaceStorage.contract.WatchLogs(opts, "OrderCancelled", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IMarketplaceStorageOrderCancelled)
				if err := _IMarketplaceStorage.contract.UnpackLog(event, "OrderCancelled", log); err != nil {
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

// ParseOrderCancelled is a log parse operation binding the contract event 0x5152abf959f6564662358c2e52b702259b78bac5ee7842a0f01937e670efcc7d.
//
// Solidity: event OrderCancelled(bytes32 indexed id)
func (_IMarketplaceStorage *IMarketplaceStorageFilterer) ParseOrderCancelled(log types.Log) (*IMarketplaceStorageOrderCancelled, error) {
	event := new(IMarketplaceStorageOrderCancelled)
	if err := _IMarketplaceStorage.contract.UnpackLog(event, "OrderCancelled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IMarketplaceStorageOrderCreatedIterator is returned from FilterOrderCreated and is used to iterate over the raw logs and unpacked data for OrderCreated events raised by the IMarketplaceStorage contract.
type IMarketplaceStorageOrderCreatedIterator struct {
	Event *IMarketplaceStorageOrderCreated // Event containing the contract specifics and raw log

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
func (it *IMarketplaceStorageOrderCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IMarketplaceStorageOrderCreated)
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
		it.Event = new(IMarketplaceStorageOrderCreated)
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
func (it *IMarketplaceStorageOrderCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IMarketplaceStorageOrderCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IMarketplaceStorageOrderCreated represents a OrderCreated event raised by the IMarketplaceStorage contract.
type IMarketplaceStorageOrderCreated struct {
	OrderId    [32]byte
	AssetId    *big.Int
	Seller     common.Address
	NftAddress common.Address
	PriceInWei *big.Int
	ExpiredAt  *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterOrderCreated is a free log retrieval operation binding the contract event 0x84c66c3f7ba4b390e20e8e8233e2a516f3ce34a72749e4f12bd010dfba238039.
//
// Solidity: event OrderCreated(bytes32 indexed orderId, uint256 indexed assetId, address indexed seller, address nftAddress, uint256 priceInWei, uint256 expiredAt)
func (_IMarketplaceStorage *IMarketplaceStorageFilterer) FilterOrderCreated(opts *bind.FilterOpts, orderId [][32]byte, assetId []*big.Int, seller []common.Address) (*IMarketplaceStorageOrderCreatedIterator, error) {

	var orderIdRule []interface{}
	for _, orderIdItem := range orderId {
		orderIdRule = append(orderIdRule, orderIdItem)
	}
	var assetIdRule []interface{}
	for _, assetIdItem := range assetId {
		assetIdRule = append(assetIdRule, assetIdItem)
	}
	var sellerRule []interface{}
	for _, sellerItem := range seller {
		sellerRule = append(sellerRule, sellerItem)
	}

	logs, sub, err := _IMarketplaceStorage.contract.FilterLogs(opts, "OrderCreated", orderIdRule, assetIdRule, sellerRule)
	if err != nil {
		return nil, err
	}
	return &IMarketplaceStorageOrderCreatedIterator{contract: _IMarketplaceStorage.contract, event: "OrderCreated", logs: logs, sub: sub}, nil
}

// WatchOrderCreated is a free log subscription operation binding the contract event 0x84c66c3f7ba4b390e20e8e8233e2a516f3ce34a72749e4f12bd010dfba238039.
//
// Solidity: event OrderCreated(bytes32 indexed orderId, uint256 indexed assetId, address indexed seller, address nftAddress, uint256 priceInWei, uint256 expiredAt)
func (_IMarketplaceStorage *IMarketplaceStorageFilterer) WatchOrderCreated(opts *bind.WatchOpts, sink chan<- *IMarketplaceStorageOrderCreated, orderId [][32]byte, assetId []*big.Int, seller []common.Address) (event.Subscription, error) {

	var orderIdRule []interface{}
	for _, orderIdItem := range orderId {
		orderIdRule = append(orderIdRule, orderIdItem)
	}
	var assetIdRule []interface{}
	for _, assetIdItem := range assetId {
		assetIdRule = append(assetIdRule, assetIdItem)
	}
	var sellerRule []interface{}
	for _, sellerItem := range seller {
		sellerRule = append(sellerRule, sellerItem)
	}

	logs, sub, err := _IMarketplaceStorage.contract.WatchLogs(opts, "OrderCreated", orderIdRule, assetIdRule, sellerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IMarketplaceStorageOrderCreated)
				if err := _IMarketplaceStorage.contract.UnpackLog(event, "OrderCreated", log); err != nil {
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

// ParseOrderCreated is a log parse operation binding the contract event 0x84c66c3f7ba4b390e20e8e8233e2a516f3ce34a72749e4f12bd010dfba238039.
//
// Solidity: event OrderCreated(bytes32 indexed orderId, uint256 indexed assetId, address indexed seller, address nftAddress, uint256 priceInWei, uint256 expiredAt)
func (_IMarketplaceStorage *IMarketplaceStorageFilterer) ParseOrderCreated(log types.Log) (*IMarketplaceStorageOrderCreated, error) {
	event := new(IMarketplaceStorageOrderCreated)
	if err := _IMarketplaceStorage.contract.UnpackLog(event, "OrderCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IMarketplaceStorageOrderSuccessfulIterator is returned from FilterOrderSuccessful and is used to iterate over the raw logs and unpacked data for OrderSuccessful events raised by the IMarketplaceStorage contract.
type IMarketplaceStorageOrderSuccessfulIterator struct {
	Event *IMarketplaceStorageOrderSuccessful // Event containing the contract specifics and raw log

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
func (it *IMarketplaceStorageOrderSuccessfulIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IMarketplaceStorageOrderSuccessful)
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
		it.Event = new(IMarketplaceStorageOrderSuccessful)
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
func (it *IMarketplaceStorageOrderSuccessfulIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IMarketplaceStorageOrderSuccessfulIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IMarketplaceStorageOrderSuccessful represents a OrderSuccessful event raised by the IMarketplaceStorage contract.
type IMarketplaceStorageOrderSuccessful struct {
	Id    [32]byte
	Buyer common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterOrderSuccessful is a free log retrieval operation binding the contract event 0x56e5f4ff73854a4423571427b069b47d6509b44bb78145207a634f2dd720916b.
//
// Solidity: event OrderSuccessful(bytes32 indexed id, address indexed buyer)
func (_IMarketplaceStorage *IMarketplaceStorageFilterer) FilterOrderSuccessful(opts *bind.FilterOpts, id [][32]byte, buyer []common.Address) (*IMarketplaceStorageOrderSuccessfulIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var buyerRule []interface{}
	for _, buyerItem := range buyer {
		buyerRule = append(buyerRule, buyerItem)
	}

	logs, sub, err := _IMarketplaceStorage.contract.FilterLogs(opts, "OrderSuccessful", idRule, buyerRule)
	if err != nil {
		return nil, err
	}
	return &IMarketplaceStorageOrderSuccessfulIterator{contract: _IMarketplaceStorage.contract, event: "OrderSuccessful", logs: logs, sub: sub}, nil
}

// WatchOrderSuccessful is a free log subscription operation binding the contract event 0x56e5f4ff73854a4423571427b069b47d6509b44bb78145207a634f2dd720916b.
//
// Solidity: event OrderSuccessful(bytes32 indexed id, address indexed buyer)
func (_IMarketplaceStorage *IMarketplaceStorageFilterer) WatchOrderSuccessful(opts *bind.WatchOpts, sink chan<- *IMarketplaceStorageOrderSuccessful, id [][32]byte, buyer []common.Address) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var buyerRule []interface{}
	for _, buyerItem := range buyer {
		buyerRule = append(buyerRule, buyerItem)
	}

	logs, sub, err := _IMarketplaceStorage.contract.WatchLogs(opts, "OrderSuccessful", idRule, buyerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IMarketplaceStorageOrderSuccessful)
				if err := _IMarketplaceStorage.contract.UnpackLog(event, "OrderSuccessful", log); err != nil {
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

// ParseOrderSuccessful is a log parse operation binding the contract event 0x56e5f4ff73854a4423571427b069b47d6509b44bb78145207a634f2dd720916b.
//
// Solidity: event OrderSuccessful(bytes32 indexed id, address indexed buyer)
func (_IMarketplaceStorage *IMarketplaceStorageFilterer) ParseOrderSuccessful(log types.Log) (*IMarketplaceStorageOrderSuccessful, error) {
	event := new(IMarketplaceStorageOrderSuccessful)
	if err := _IMarketplaceStorage.contract.UnpackLog(event, "OrderSuccessful", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IMarketplaceStorageRevealFailedIterator is returned from FilterRevealFailed and is used to iterate over the raw logs and unpacked data for RevealFailed events raised by the IMarketplaceStorage contract.
type IMarketplaceStorageRevealFailedIterator struct {
	Event *IMarketplaceStorageRevealFailed // Event containing the contract specifics and raw log

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
func (it *IMarketplaceStorageRevealFailedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IMarketplaceStorageRevealFailed)
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
		it.Event = new(IMarketplaceStorageRevealFailed)
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
func (it *IMarketplaceStorageRevealFailedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IMarketplaceStorageRevealFailedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IMarketplaceStorageRevealFailed represents a RevealFailed event raised by the IMarketplaceStorage contract.
type IMarketplaceStorageRevealFailed struct {
	Fake      bool
	AuctionId [32]byte
	Value     *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterRevealFailed is a free log retrieval operation binding the contract event 0x8491bb9f45e55b89e41be77fb6f559dcedeaed2ed85a61f74d039b2f1389d4a3.
//
// Solidity: event RevealFailed(bool fake, bytes32 indexed auctionId, uint256 value)
func (_IMarketplaceStorage *IMarketplaceStorageFilterer) FilterRevealFailed(opts *bind.FilterOpts, auctionId [][32]byte) (*IMarketplaceStorageRevealFailedIterator, error) {

	var auctionIdRule []interface{}
	for _, auctionIdItem := range auctionId {
		auctionIdRule = append(auctionIdRule, auctionIdItem)
	}

	logs, sub, err := _IMarketplaceStorage.contract.FilterLogs(opts, "RevealFailed", auctionIdRule)
	if err != nil {
		return nil, err
	}
	return &IMarketplaceStorageRevealFailedIterator{contract: _IMarketplaceStorage.contract, event: "RevealFailed", logs: logs, sub: sub}, nil
}

// WatchRevealFailed is a free log subscription operation binding the contract event 0x8491bb9f45e55b89e41be77fb6f559dcedeaed2ed85a61f74d039b2f1389d4a3.
//
// Solidity: event RevealFailed(bool fake, bytes32 indexed auctionId, uint256 value)
func (_IMarketplaceStorage *IMarketplaceStorageFilterer) WatchRevealFailed(opts *bind.WatchOpts, sink chan<- *IMarketplaceStorageRevealFailed, auctionId [][32]byte) (event.Subscription, error) {

	var auctionIdRule []interface{}
	for _, auctionIdItem := range auctionId {
		auctionIdRule = append(auctionIdRule, auctionIdItem)
	}

	logs, sub, err := _IMarketplaceStorage.contract.WatchLogs(opts, "RevealFailed", auctionIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IMarketplaceStorageRevealFailed)
				if err := _IMarketplaceStorage.contract.UnpackLog(event, "RevealFailed", log); err != nil {
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

// ParseRevealFailed is a log parse operation binding the contract event 0x8491bb9f45e55b89e41be77fb6f559dcedeaed2ed85a61f74d039b2f1389d4a3.
//
// Solidity: event RevealFailed(bool fake, bytes32 indexed auctionId, uint256 value)
func (_IMarketplaceStorage *IMarketplaceStorageFilterer) ParseRevealFailed(log types.Log) (*IMarketplaceStorageRevealFailed, error) {
	event := new(IMarketplaceStorageRevealFailed)
	if err := _IMarketplaceStorage.contract.UnpackLog(event, "RevealFailed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IMarketplaceStorageRevealSuccessfulIterator is returned from FilterRevealSuccessful and is used to iterate over the raw logs and unpacked data for RevealSuccessful events raised by the IMarketplaceStorage contract.
type IMarketplaceStorageRevealSuccessfulIterator struct {
	Event *IMarketplaceStorageRevealSuccessful // Event containing the contract specifics and raw log

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
func (it *IMarketplaceStorageRevealSuccessfulIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IMarketplaceStorageRevealSuccessful)
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
		it.Event = new(IMarketplaceStorageRevealSuccessful)
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
func (it *IMarketplaceStorageRevealSuccessfulIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IMarketplaceStorageRevealSuccessfulIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IMarketplaceStorageRevealSuccessful represents a RevealSuccessful event raised by the IMarketplaceStorage contract.
type IMarketplaceStorageRevealSuccessful struct {
	Fake       bool
	AuctionId  [32]byte
	Value      *big.Int
	BlindedBid [32]byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterRevealSuccessful is a free log retrieval operation binding the contract event 0xc69be17b446ade79dcd9f07e13f1c8493dd4a1940caf6318d617fbd25d256d9b.
//
// Solidity: event RevealSuccessful(bool fake, bytes32 indexed auctionId, uint256 value, bytes32 blindedBid)
func (_IMarketplaceStorage *IMarketplaceStorageFilterer) FilterRevealSuccessful(opts *bind.FilterOpts, auctionId [][32]byte) (*IMarketplaceStorageRevealSuccessfulIterator, error) {

	var auctionIdRule []interface{}
	for _, auctionIdItem := range auctionId {
		auctionIdRule = append(auctionIdRule, auctionIdItem)
	}

	logs, sub, err := _IMarketplaceStorage.contract.FilterLogs(opts, "RevealSuccessful", auctionIdRule)
	if err != nil {
		return nil, err
	}
	return &IMarketplaceStorageRevealSuccessfulIterator{contract: _IMarketplaceStorage.contract, event: "RevealSuccessful", logs: logs, sub: sub}, nil
}

// WatchRevealSuccessful is a free log subscription operation binding the contract event 0xc69be17b446ade79dcd9f07e13f1c8493dd4a1940caf6318d617fbd25d256d9b.
//
// Solidity: event RevealSuccessful(bool fake, bytes32 indexed auctionId, uint256 value, bytes32 blindedBid)
func (_IMarketplaceStorage *IMarketplaceStorageFilterer) WatchRevealSuccessful(opts *bind.WatchOpts, sink chan<- *IMarketplaceStorageRevealSuccessful, auctionId [][32]byte) (event.Subscription, error) {

	var auctionIdRule []interface{}
	for _, auctionIdItem := range auctionId {
		auctionIdRule = append(auctionIdRule, auctionIdItem)
	}

	logs, sub, err := _IMarketplaceStorage.contract.WatchLogs(opts, "RevealSuccessful", auctionIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IMarketplaceStorageRevealSuccessful)
				if err := _IMarketplaceStorage.contract.UnpackLog(event, "RevealSuccessful", log); err != nil {
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

// ParseRevealSuccessful is a log parse operation binding the contract event 0xc69be17b446ade79dcd9f07e13f1c8493dd4a1940caf6318d617fbd25d256d9b.
//
// Solidity: event RevealSuccessful(bool fake, bytes32 indexed auctionId, uint256 value, bytes32 blindedBid)
func (_IMarketplaceStorage *IMarketplaceStorageFilterer) ParseRevealSuccessful(log types.Log) (*IMarketplaceStorageRevealSuccessful, error) {
	event := new(IMarketplaceStorageRevealSuccessful)
	if err := _IMarketplaceStorage.contract.UnpackLog(event, "RevealSuccessful", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
