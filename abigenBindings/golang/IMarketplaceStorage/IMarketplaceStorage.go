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

// IBlindAuctionBlindAuction is an auto generated low-level Go binding around an user-defined struct.
type IBlindAuctionBlindAuction struct {
	Id            [32]byte
	Seller        common.Address
	BiddingEnd    *big.Int
	RevealEnd     *big.Int
	StartPrice    *big.Int
	HighestBid    *big.Int
	HighestBidder common.Address
}

// IOrderOrder is an auto generated low-level Go binding around an user-defined struct.
type IOrderOrder struct {
	OrderId    [32]byte
	Seller     common.Address
	NftAddress common.Address
	Price      *big.Int
	ExpiredAt  *big.Int
}

// IPublicAuctionPublicAuction is an auto generated low-level Go binding around an user-defined struct.
type IPublicAuctionPublicAuction struct {
	Id            [32]byte
	Seller        common.Address
	HighestBid    *big.Int
	HighestBidder common.Address
	BiddingEnd    *big.Int
	StartPrice    *big.Int
	MinIncrement  *big.Int
}

// IMarketplaceStorageMetaData contains all meta data concerning the IMarketplaceStorage contract.
var IMarketplaceStorageMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"AssetNotExisted\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AssetUnvailable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AuctionAlreadyEnded\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidMkpSender\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"bidder\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"blindAuctionId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"blindedBid\",\"type\":\"bytes32\"}],\"name\":\"BlindAuctionBidSuccessful\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"canceller\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"blindAuctionId\",\"type\":\"bytes32\"}],\"name\":\"BlindAuctionCancelled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"nftAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"blindAuctionId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"assetId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"biddingEnd\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"revealEnd\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"startPriceInWei\",\"type\":\"uint256\"}],\"name\":\"BlindAuctionCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"blindAuctionId\",\"type\":\"bytes32\"}],\"name\":\"BlindAuctionEnded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"bidder\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"blindAuctionId\",\"type\":\"bytes32\"}],\"name\":\"BlindAuctionRefund\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"buyer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"blindAuctionId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalPrice\",\"type\":\"uint256\"}],\"name\":\"BlindAuctionSuccessful\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"who\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"}],\"name\":\"OrderCancelled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"orderId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"assetId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"nftAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"priceInWei\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"expiredAt\",\"type\":\"uint256\"}],\"name\":\"OrderCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"buyer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"}],\"name\":\"OrderSuccessful\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"bidder\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"publicAuctionId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"bidValue\",\"type\":\"uint256\"}],\"name\":\"PublicAuctionBidSuccessful\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"canceller\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"publicAuctionId\",\"type\":\"bytes32\"}],\"name\":\"PublicAuctionCancelled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"nftAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"publicAuctionId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"assetId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"biddingEnd\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"startPriceInWei\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"minIncrement\",\"type\":\"uint256\"}],\"name\":\"PublicAuctionCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"publicAuctionId\",\"type\":\"bytes32\"}],\"name\":\"PublicAuctionEnded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"bidder\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"publicAuctionId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"PublicAuctionRefund\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"buyer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"publicAuctionId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalPrice\",\"type\":\"uint256\"}],\"name\":\"PublicAuctionSuccessful\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"blindAuctionId\",\"type\":\"bytes32\"}],\"name\":\"RevealSuccessful\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"nftAsset\",\"type\":\"bytes32\"}],\"name\":\"assetIsAvailable\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"publicAuctionId\",\"type\":\"bytes32\"}],\"name\":\"publicAuctionIsExisted\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"existed\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"nftAsset\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"publicAuctionId\",\"type\":\"bytes32\"}],\"name\":\"publicAuctionIsEnded\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"ended\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"nftAsset\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"publicAuctionId\",\"type\":\"bytes32\"}],\"name\":\"publicAuctionIsRunning\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"existed\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"assetOwner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"nftAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"assetId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"publicAuctionId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"biddingEnd\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startPriceInWei\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minIncrement\",\"type\":\"uint256\"}],\"name\":\"createPublicAuction\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"publicAuctionId\",\"type\":\"bytes32\"}],\"name\":\"getPublicAuction\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"highestBid\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"highestBidder\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"biddingEnd\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minIncrement\",\"type\":\"uint256\"}],\"internalType\":\"structIPublicAuction.PublicAuction\",\"name\":\"publicAuction\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"nftAsset\",\"type\":\"bytes32\"}],\"name\":\"endPublicAuction\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"bidder\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"highestBid\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"publicAuctionId\",\"type\":\"bytes32\"}],\"name\":\"updateHighestBidPublicAuction\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"blindAuctionId\",\"type\":\"bytes32\"}],\"name\":\"blindAuctionIsExisted\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"existed\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"nftAsset\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"blindAuctionId\",\"type\":\"bytes32\"}],\"name\":\"blindAuctionIsEnded\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"ended\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"nftAsset\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"blindAuctionId\",\"type\":\"bytes32\"}],\"name\":\"blindAuctionIsRunning\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"existed\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"assetOwner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"nftAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"assetId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"blindAuctionId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"biddingEnd\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"revealEnd\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startPriceInWei\",\"type\":\"uint256\"}],\"name\":\"createBlindAuction\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"blindAuctionId\",\"type\":\"bytes32\"}],\"name\":\"getBlindAuction\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"biddingEnd\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"revealEnd\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"highestBid\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"highestBidder\",\"type\":\"address\"}],\"internalType\":\"structIBlindAuction.BlindAuction\",\"name\":\"blindAuction\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"nftAsset\",\"type\":\"bytes32\"}],\"name\":\"endBlindAuction\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"bidder\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"highestBid\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"blindAuctionId\",\"type\":\"bytes32\"}],\"name\":\"updateHighestBidBlindAuction\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"nftAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"assetId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"orderId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"priceInWei\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expiredAt\",\"type\":\"uint256\"}],\"name\":\"createOrder\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"nftAsset\",\"type\":\"bytes32\"}],\"name\":\"getOrder\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"orderId\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"nftAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expiredAt\",\"type\":\"uint256\"}],\"internalType\":\"structIOrder.Order\",\"name\":\"order\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"nftAsset\",\"type\":\"bytes32\"}],\"name\":\"deleteOrder\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
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

// BlindAuctionIsEnded is a free data retrieval call binding the contract method 0xcd23b172.
//
// Solidity: function blindAuctionIsEnded(bytes32 nftAsset, bytes32 blindAuctionId) view returns(bool ended)
func (_IMarketplaceStorage *IMarketplaceStorageCaller) BlindAuctionIsEnded(opts *bind.CallOpts, nftAsset [32]byte, blindAuctionId [32]byte) (bool, error) {
	var out []interface{}
	err := _IMarketplaceStorage.contract.Call(opts, &out, "blindAuctionIsEnded", nftAsset, blindAuctionId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// BlindAuctionIsEnded is a free data retrieval call binding the contract method 0xcd23b172.
//
// Solidity: function blindAuctionIsEnded(bytes32 nftAsset, bytes32 blindAuctionId) view returns(bool ended)
func (_IMarketplaceStorage *IMarketplaceStorageSession) BlindAuctionIsEnded(nftAsset [32]byte, blindAuctionId [32]byte) (bool, error) {
	return _IMarketplaceStorage.Contract.BlindAuctionIsEnded(&_IMarketplaceStorage.CallOpts, nftAsset, blindAuctionId)
}

// BlindAuctionIsEnded is a free data retrieval call binding the contract method 0xcd23b172.
//
// Solidity: function blindAuctionIsEnded(bytes32 nftAsset, bytes32 blindAuctionId) view returns(bool ended)
func (_IMarketplaceStorage *IMarketplaceStorageCallerSession) BlindAuctionIsEnded(nftAsset [32]byte, blindAuctionId [32]byte) (bool, error) {
	return _IMarketplaceStorage.Contract.BlindAuctionIsEnded(&_IMarketplaceStorage.CallOpts, nftAsset, blindAuctionId)
}

// BlindAuctionIsExisted is a free data retrieval call binding the contract method 0x86624448.
//
// Solidity: function blindAuctionIsExisted(bytes32 blindAuctionId) view returns(bool existed)
func (_IMarketplaceStorage *IMarketplaceStorageCaller) BlindAuctionIsExisted(opts *bind.CallOpts, blindAuctionId [32]byte) (bool, error) {
	var out []interface{}
	err := _IMarketplaceStorage.contract.Call(opts, &out, "blindAuctionIsExisted", blindAuctionId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// BlindAuctionIsExisted is a free data retrieval call binding the contract method 0x86624448.
//
// Solidity: function blindAuctionIsExisted(bytes32 blindAuctionId) view returns(bool existed)
func (_IMarketplaceStorage *IMarketplaceStorageSession) BlindAuctionIsExisted(blindAuctionId [32]byte) (bool, error) {
	return _IMarketplaceStorage.Contract.BlindAuctionIsExisted(&_IMarketplaceStorage.CallOpts, blindAuctionId)
}

// BlindAuctionIsExisted is a free data retrieval call binding the contract method 0x86624448.
//
// Solidity: function blindAuctionIsExisted(bytes32 blindAuctionId) view returns(bool existed)
func (_IMarketplaceStorage *IMarketplaceStorageCallerSession) BlindAuctionIsExisted(blindAuctionId [32]byte) (bool, error) {
	return _IMarketplaceStorage.Contract.BlindAuctionIsExisted(&_IMarketplaceStorage.CallOpts, blindAuctionId)
}

// BlindAuctionIsRunning is a free data retrieval call binding the contract method 0x4a43c5b9.
//
// Solidity: function blindAuctionIsRunning(bytes32 nftAsset, bytes32 blindAuctionId) view returns(bool existed)
func (_IMarketplaceStorage *IMarketplaceStorageCaller) BlindAuctionIsRunning(opts *bind.CallOpts, nftAsset [32]byte, blindAuctionId [32]byte) (bool, error) {
	var out []interface{}
	err := _IMarketplaceStorage.contract.Call(opts, &out, "blindAuctionIsRunning", nftAsset, blindAuctionId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// BlindAuctionIsRunning is a free data retrieval call binding the contract method 0x4a43c5b9.
//
// Solidity: function blindAuctionIsRunning(bytes32 nftAsset, bytes32 blindAuctionId) view returns(bool existed)
func (_IMarketplaceStorage *IMarketplaceStorageSession) BlindAuctionIsRunning(nftAsset [32]byte, blindAuctionId [32]byte) (bool, error) {
	return _IMarketplaceStorage.Contract.BlindAuctionIsRunning(&_IMarketplaceStorage.CallOpts, nftAsset, blindAuctionId)
}

// BlindAuctionIsRunning is a free data retrieval call binding the contract method 0x4a43c5b9.
//
// Solidity: function blindAuctionIsRunning(bytes32 nftAsset, bytes32 blindAuctionId) view returns(bool existed)
func (_IMarketplaceStorage *IMarketplaceStorageCallerSession) BlindAuctionIsRunning(nftAsset [32]byte, blindAuctionId [32]byte) (bool, error) {
	return _IMarketplaceStorage.Contract.BlindAuctionIsRunning(&_IMarketplaceStorage.CallOpts, nftAsset, blindAuctionId)
}

// GetBlindAuction is a free data retrieval call binding the contract method 0x6ed3be53.
//
// Solidity: function getBlindAuction(bytes32 blindAuctionId) view returns((bytes32,address,uint256,uint256,uint256,uint256,address) blindAuction)
func (_IMarketplaceStorage *IMarketplaceStorageCaller) GetBlindAuction(opts *bind.CallOpts, blindAuctionId [32]byte) (IBlindAuctionBlindAuction, error) {
	var out []interface{}
	err := _IMarketplaceStorage.contract.Call(opts, &out, "getBlindAuction", blindAuctionId)

	if err != nil {
		return *new(IBlindAuctionBlindAuction), err
	}

	out0 := *abi.ConvertType(out[0], new(IBlindAuctionBlindAuction)).(*IBlindAuctionBlindAuction)

	return out0, err

}

// GetBlindAuction is a free data retrieval call binding the contract method 0x6ed3be53.
//
// Solidity: function getBlindAuction(bytes32 blindAuctionId) view returns((bytes32,address,uint256,uint256,uint256,uint256,address) blindAuction)
func (_IMarketplaceStorage *IMarketplaceStorageSession) GetBlindAuction(blindAuctionId [32]byte) (IBlindAuctionBlindAuction, error) {
	return _IMarketplaceStorage.Contract.GetBlindAuction(&_IMarketplaceStorage.CallOpts, blindAuctionId)
}

// GetBlindAuction is a free data retrieval call binding the contract method 0x6ed3be53.
//
// Solidity: function getBlindAuction(bytes32 blindAuctionId) view returns((bytes32,address,uint256,uint256,uint256,uint256,address) blindAuction)
func (_IMarketplaceStorage *IMarketplaceStorageCallerSession) GetBlindAuction(blindAuctionId [32]byte) (IBlindAuctionBlindAuction, error) {
	return _IMarketplaceStorage.Contract.GetBlindAuction(&_IMarketplaceStorage.CallOpts, blindAuctionId)
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

// GetPublicAuction is a free data retrieval call binding the contract method 0x57a45b77.
//
// Solidity: function getPublicAuction(bytes32 publicAuctionId) view returns((bytes32,address,uint256,address,uint256,uint256,uint256) publicAuction)
func (_IMarketplaceStorage *IMarketplaceStorageCaller) GetPublicAuction(opts *bind.CallOpts, publicAuctionId [32]byte) (IPublicAuctionPublicAuction, error) {
	var out []interface{}
	err := _IMarketplaceStorage.contract.Call(opts, &out, "getPublicAuction", publicAuctionId)

	if err != nil {
		return *new(IPublicAuctionPublicAuction), err
	}

	out0 := *abi.ConvertType(out[0], new(IPublicAuctionPublicAuction)).(*IPublicAuctionPublicAuction)

	return out0, err

}

// GetPublicAuction is a free data retrieval call binding the contract method 0x57a45b77.
//
// Solidity: function getPublicAuction(bytes32 publicAuctionId) view returns((bytes32,address,uint256,address,uint256,uint256,uint256) publicAuction)
func (_IMarketplaceStorage *IMarketplaceStorageSession) GetPublicAuction(publicAuctionId [32]byte) (IPublicAuctionPublicAuction, error) {
	return _IMarketplaceStorage.Contract.GetPublicAuction(&_IMarketplaceStorage.CallOpts, publicAuctionId)
}

// GetPublicAuction is a free data retrieval call binding the contract method 0x57a45b77.
//
// Solidity: function getPublicAuction(bytes32 publicAuctionId) view returns((bytes32,address,uint256,address,uint256,uint256,uint256) publicAuction)
func (_IMarketplaceStorage *IMarketplaceStorageCallerSession) GetPublicAuction(publicAuctionId [32]byte) (IPublicAuctionPublicAuction, error) {
	return _IMarketplaceStorage.Contract.GetPublicAuction(&_IMarketplaceStorage.CallOpts, publicAuctionId)
}

// PublicAuctionIsEnded is a free data retrieval call binding the contract method 0xbc4124c5.
//
// Solidity: function publicAuctionIsEnded(bytes32 nftAsset, bytes32 publicAuctionId) view returns(bool ended)
func (_IMarketplaceStorage *IMarketplaceStorageCaller) PublicAuctionIsEnded(opts *bind.CallOpts, nftAsset [32]byte, publicAuctionId [32]byte) (bool, error) {
	var out []interface{}
	err := _IMarketplaceStorage.contract.Call(opts, &out, "publicAuctionIsEnded", nftAsset, publicAuctionId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// PublicAuctionIsEnded is a free data retrieval call binding the contract method 0xbc4124c5.
//
// Solidity: function publicAuctionIsEnded(bytes32 nftAsset, bytes32 publicAuctionId) view returns(bool ended)
func (_IMarketplaceStorage *IMarketplaceStorageSession) PublicAuctionIsEnded(nftAsset [32]byte, publicAuctionId [32]byte) (bool, error) {
	return _IMarketplaceStorage.Contract.PublicAuctionIsEnded(&_IMarketplaceStorage.CallOpts, nftAsset, publicAuctionId)
}

// PublicAuctionIsEnded is a free data retrieval call binding the contract method 0xbc4124c5.
//
// Solidity: function publicAuctionIsEnded(bytes32 nftAsset, bytes32 publicAuctionId) view returns(bool ended)
func (_IMarketplaceStorage *IMarketplaceStorageCallerSession) PublicAuctionIsEnded(nftAsset [32]byte, publicAuctionId [32]byte) (bool, error) {
	return _IMarketplaceStorage.Contract.PublicAuctionIsEnded(&_IMarketplaceStorage.CallOpts, nftAsset, publicAuctionId)
}

// PublicAuctionIsExisted is a free data retrieval call binding the contract method 0x7584f03d.
//
// Solidity: function publicAuctionIsExisted(bytes32 publicAuctionId) view returns(bool existed)
func (_IMarketplaceStorage *IMarketplaceStorageCaller) PublicAuctionIsExisted(opts *bind.CallOpts, publicAuctionId [32]byte) (bool, error) {
	var out []interface{}
	err := _IMarketplaceStorage.contract.Call(opts, &out, "publicAuctionIsExisted", publicAuctionId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// PublicAuctionIsExisted is a free data retrieval call binding the contract method 0x7584f03d.
//
// Solidity: function publicAuctionIsExisted(bytes32 publicAuctionId) view returns(bool existed)
func (_IMarketplaceStorage *IMarketplaceStorageSession) PublicAuctionIsExisted(publicAuctionId [32]byte) (bool, error) {
	return _IMarketplaceStorage.Contract.PublicAuctionIsExisted(&_IMarketplaceStorage.CallOpts, publicAuctionId)
}

// PublicAuctionIsExisted is a free data retrieval call binding the contract method 0x7584f03d.
//
// Solidity: function publicAuctionIsExisted(bytes32 publicAuctionId) view returns(bool existed)
func (_IMarketplaceStorage *IMarketplaceStorageCallerSession) PublicAuctionIsExisted(publicAuctionId [32]byte) (bool, error) {
	return _IMarketplaceStorage.Contract.PublicAuctionIsExisted(&_IMarketplaceStorage.CallOpts, publicAuctionId)
}

// PublicAuctionIsRunning is a free data retrieval call binding the contract method 0x002cb6a3.
//
// Solidity: function publicAuctionIsRunning(bytes32 nftAsset, bytes32 publicAuctionId) view returns(bool existed)
func (_IMarketplaceStorage *IMarketplaceStorageCaller) PublicAuctionIsRunning(opts *bind.CallOpts, nftAsset [32]byte, publicAuctionId [32]byte) (bool, error) {
	var out []interface{}
	err := _IMarketplaceStorage.contract.Call(opts, &out, "publicAuctionIsRunning", nftAsset, publicAuctionId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// PublicAuctionIsRunning is a free data retrieval call binding the contract method 0x002cb6a3.
//
// Solidity: function publicAuctionIsRunning(bytes32 nftAsset, bytes32 publicAuctionId) view returns(bool existed)
func (_IMarketplaceStorage *IMarketplaceStorageSession) PublicAuctionIsRunning(nftAsset [32]byte, publicAuctionId [32]byte) (bool, error) {
	return _IMarketplaceStorage.Contract.PublicAuctionIsRunning(&_IMarketplaceStorage.CallOpts, nftAsset, publicAuctionId)
}

// PublicAuctionIsRunning is a free data retrieval call binding the contract method 0x002cb6a3.
//
// Solidity: function publicAuctionIsRunning(bytes32 nftAsset, bytes32 publicAuctionId) view returns(bool existed)
func (_IMarketplaceStorage *IMarketplaceStorageCallerSession) PublicAuctionIsRunning(nftAsset [32]byte, publicAuctionId [32]byte) (bool, error) {
	return _IMarketplaceStorage.Contract.PublicAuctionIsRunning(&_IMarketplaceStorage.CallOpts, nftAsset, publicAuctionId)
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

// CreateBlindAuction is a paid mutator transaction binding the contract method 0xd7e4a307.
//
// Solidity: function createBlindAuction(address assetOwner, address nftAddress, uint256 assetId, bytes32 blindAuctionId, uint256 biddingEnd, uint256 revealEnd, uint256 startPriceInWei) returns()
func (_IMarketplaceStorage *IMarketplaceStorageTransactor) CreateBlindAuction(opts *bind.TransactOpts, assetOwner common.Address, nftAddress common.Address, assetId *big.Int, blindAuctionId [32]byte, biddingEnd *big.Int, revealEnd *big.Int, startPriceInWei *big.Int) (*types.Transaction, error) {
	return _IMarketplaceStorage.contract.Transact(opts, "createBlindAuction", assetOwner, nftAddress, assetId, blindAuctionId, biddingEnd, revealEnd, startPriceInWei)
}

// CreateBlindAuction is a paid mutator transaction binding the contract method 0xd7e4a307.
//
// Solidity: function createBlindAuction(address assetOwner, address nftAddress, uint256 assetId, bytes32 blindAuctionId, uint256 biddingEnd, uint256 revealEnd, uint256 startPriceInWei) returns()
func (_IMarketplaceStorage *IMarketplaceStorageSession) CreateBlindAuction(assetOwner common.Address, nftAddress common.Address, assetId *big.Int, blindAuctionId [32]byte, biddingEnd *big.Int, revealEnd *big.Int, startPriceInWei *big.Int) (*types.Transaction, error) {
	return _IMarketplaceStorage.Contract.CreateBlindAuction(&_IMarketplaceStorage.TransactOpts, assetOwner, nftAddress, assetId, blindAuctionId, biddingEnd, revealEnd, startPriceInWei)
}

// CreateBlindAuction is a paid mutator transaction binding the contract method 0xd7e4a307.
//
// Solidity: function createBlindAuction(address assetOwner, address nftAddress, uint256 assetId, bytes32 blindAuctionId, uint256 biddingEnd, uint256 revealEnd, uint256 startPriceInWei) returns()
func (_IMarketplaceStorage *IMarketplaceStorageTransactorSession) CreateBlindAuction(assetOwner common.Address, nftAddress common.Address, assetId *big.Int, blindAuctionId [32]byte, biddingEnd *big.Int, revealEnd *big.Int, startPriceInWei *big.Int) (*types.Transaction, error) {
	return _IMarketplaceStorage.Contract.CreateBlindAuction(&_IMarketplaceStorage.TransactOpts, assetOwner, nftAddress, assetId, blindAuctionId, biddingEnd, revealEnd, startPriceInWei)
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

// CreatePublicAuction is a paid mutator transaction binding the contract method 0x6af4c895.
//
// Solidity: function createPublicAuction(address assetOwner, address nftAddress, uint256 assetId, bytes32 publicAuctionId, uint256 biddingEnd, uint256 startPriceInWei, uint256 minIncrement) returns()
func (_IMarketplaceStorage *IMarketplaceStorageTransactor) CreatePublicAuction(opts *bind.TransactOpts, assetOwner common.Address, nftAddress common.Address, assetId *big.Int, publicAuctionId [32]byte, biddingEnd *big.Int, startPriceInWei *big.Int, minIncrement *big.Int) (*types.Transaction, error) {
	return _IMarketplaceStorage.contract.Transact(opts, "createPublicAuction", assetOwner, nftAddress, assetId, publicAuctionId, biddingEnd, startPriceInWei, minIncrement)
}

// CreatePublicAuction is a paid mutator transaction binding the contract method 0x6af4c895.
//
// Solidity: function createPublicAuction(address assetOwner, address nftAddress, uint256 assetId, bytes32 publicAuctionId, uint256 biddingEnd, uint256 startPriceInWei, uint256 minIncrement) returns()
func (_IMarketplaceStorage *IMarketplaceStorageSession) CreatePublicAuction(assetOwner common.Address, nftAddress common.Address, assetId *big.Int, publicAuctionId [32]byte, biddingEnd *big.Int, startPriceInWei *big.Int, minIncrement *big.Int) (*types.Transaction, error) {
	return _IMarketplaceStorage.Contract.CreatePublicAuction(&_IMarketplaceStorage.TransactOpts, assetOwner, nftAddress, assetId, publicAuctionId, biddingEnd, startPriceInWei, minIncrement)
}

// CreatePublicAuction is a paid mutator transaction binding the contract method 0x6af4c895.
//
// Solidity: function createPublicAuction(address assetOwner, address nftAddress, uint256 assetId, bytes32 publicAuctionId, uint256 biddingEnd, uint256 startPriceInWei, uint256 minIncrement) returns()
func (_IMarketplaceStorage *IMarketplaceStorageTransactorSession) CreatePublicAuction(assetOwner common.Address, nftAddress common.Address, assetId *big.Int, publicAuctionId [32]byte, biddingEnd *big.Int, startPriceInWei *big.Int, minIncrement *big.Int) (*types.Transaction, error) {
	return _IMarketplaceStorage.Contract.CreatePublicAuction(&_IMarketplaceStorage.TransactOpts, assetOwner, nftAddress, assetId, publicAuctionId, biddingEnd, startPriceInWei, minIncrement)
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

// EndBlindAuction is a paid mutator transaction binding the contract method 0x15c97238.
//
// Solidity: function endBlindAuction(bytes32 nftAsset) returns()
func (_IMarketplaceStorage *IMarketplaceStorageTransactor) EndBlindAuction(opts *bind.TransactOpts, nftAsset [32]byte) (*types.Transaction, error) {
	return _IMarketplaceStorage.contract.Transact(opts, "endBlindAuction", nftAsset)
}

// EndBlindAuction is a paid mutator transaction binding the contract method 0x15c97238.
//
// Solidity: function endBlindAuction(bytes32 nftAsset) returns()
func (_IMarketplaceStorage *IMarketplaceStorageSession) EndBlindAuction(nftAsset [32]byte) (*types.Transaction, error) {
	return _IMarketplaceStorage.Contract.EndBlindAuction(&_IMarketplaceStorage.TransactOpts, nftAsset)
}

// EndBlindAuction is a paid mutator transaction binding the contract method 0x15c97238.
//
// Solidity: function endBlindAuction(bytes32 nftAsset) returns()
func (_IMarketplaceStorage *IMarketplaceStorageTransactorSession) EndBlindAuction(nftAsset [32]byte) (*types.Transaction, error) {
	return _IMarketplaceStorage.Contract.EndBlindAuction(&_IMarketplaceStorage.TransactOpts, nftAsset)
}

// EndPublicAuction is a paid mutator transaction binding the contract method 0x5a0affbb.
//
// Solidity: function endPublicAuction(bytes32 nftAsset) returns()
func (_IMarketplaceStorage *IMarketplaceStorageTransactor) EndPublicAuction(opts *bind.TransactOpts, nftAsset [32]byte) (*types.Transaction, error) {
	return _IMarketplaceStorage.contract.Transact(opts, "endPublicAuction", nftAsset)
}

// EndPublicAuction is a paid mutator transaction binding the contract method 0x5a0affbb.
//
// Solidity: function endPublicAuction(bytes32 nftAsset) returns()
func (_IMarketplaceStorage *IMarketplaceStorageSession) EndPublicAuction(nftAsset [32]byte) (*types.Transaction, error) {
	return _IMarketplaceStorage.Contract.EndPublicAuction(&_IMarketplaceStorage.TransactOpts, nftAsset)
}

// EndPublicAuction is a paid mutator transaction binding the contract method 0x5a0affbb.
//
// Solidity: function endPublicAuction(bytes32 nftAsset) returns()
func (_IMarketplaceStorage *IMarketplaceStorageTransactorSession) EndPublicAuction(nftAsset [32]byte) (*types.Transaction, error) {
	return _IMarketplaceStorage.Contract.EndPublicAuction(&_IMarketplaceStorage.TransactOpts, nftAsset)
}

// UpdateHighestBidBlindAuction is a paid mutator transaction binding the contract method 0xe7c847aa.
//
// Solidity: function updateHighestBidBlindAuction(address bidder, uint256 highestBid, bytes32 blindAuctionId) returns()
func (_IMarketplaceStorage *IMarketplaceStorageTransactor) UpdateHighestBidBlindAuction(opts *bind.TransactOpts, bidder common.Address, highestBid *big.Int, blindAuctionId [32]byte) (*types.Transaction, error) {
	return _IMarketplaceStorage.contract.Transact(opts, "updateHighestBidBlindAuction", bidder, highestBid, blindAuctionId)
}

// UpdateHighestBidBlindAuction is a paid mutator transaction binding the contract method 0xe7c847aa.
//
// Solidity: function updateHighestBidBlindAuction(address bidder, uint256 highestBid, bytes32 blindAuctionId) returns()
func (_IMarketplaceStorage *IMarketplaceStorageSession) UpdateHighestBidBlindAuction(bidder common.Address, highestBid *big.Int, blindAuctionId [32]byte) (*types.Transaction, error) {
	return _IMarketplaceStorage.Contract.UpdateHighestBidBlindAuction(&_IMarketplaceStorage.TransactOpts, bidder, highestBid, blindAuctionId)
}

// UpdateHighestBidBlindAuction is a paid mutator transaction binding the contract method 0xe7c847aa.
//
// Solidity: function updateHighestBidBlindAuction(address bidder, uint256 highestBid, bytes32 blindAuctionId) returns()
func (_IMarketplaceStorage *IMarketplaceStorageTransactorSession) UpdateHighestBidBlindAuction(bidder common.Address, highestBid *big.Int, blindAuctionId [32]byte) (*types.Transaction, error) {
	return _IMarketplaceStorage.Contract.UpdateHighestBidBlindAuction(&_IMarketplaceStorage.TransactOpts, bidder, highestBid, blindAuctionId)
}

// UpdateHighestBidPublicAuction is a paid mutator transaction binding the contract method 0x9d90ce3b.
//
// Solidity: function updateHighestBidPublicAuction(address bidder, uint256 highestBid, bytes32 publicAuctionId) returns()
func (_IMarketplaceStorage *IMarketplaceStorageTransactor) UpdateHighestBidPublicAuction(opts *bind.TransactOpts, bidder common.Address, highestBid *big.Int, publicAuctionId [32]byte) (*types.Transaction, error) {
	return _IMarketplaceStorage.contract.Transact(opts, "updateHighestBidPublicAuction", bidder, highestBid, publicAuctionId)
}

// UpdateHighestBidPublicAuction is a paid mutator transaction binding the contract method 0x9d90ce3b.
//
// Solidity: function updateHighestBidPublicAuction(address bidder, uint256 highestBid, bytes32 publicAuctionId) returns()
func (_IMarketplaceStorage *IMarketplaceStorageSession) UpdateHighestBidPublicAuction(bidder common.Address, highestBid *big.Int, publicAuctionId [32]byte) (*types.Transaction, error) {
	return _IMarketplaceStorage.Contract.UpdateHighestBidPublicAuction(&_IMarketplaceStorage.TransactOpts, bidder, highestBid, publicAuctionId)
}

// UpdateHighestBidPublicAuction is a paid mutator transaction binding the contract method 0x9d90ce3b.
//
// Solidity: function updateHighestBidPublicAuction(address bidder, uint256 highestBid, bytes32 publicAuctionId) returns()
func (_IMarketplaceStorage *IMarketplaceStorageTransactorSession) UpdateHighestBidPublicAuction(bidder common.Address, highestBid *big.Int, publicAuctionId [32]byte) (*types.Transaction, error) {
	return _IMarketplaceStorage.Contract.UpdateHighestBidPublicAuction(&_IMarketplaceStorage.TransactOpts, bidder, highestBid, publicAuctionId)
}

// IMarketplaceStorageBlindAuctionBidSuccessfulIterator is returned from FilterBlindAuctionBidSuccessful and is used to iterate over the raw logs and unpacked data for BlindAuctionBidSuccessful events raised by the IMarketplaceStorage contract.
type IMarketplaceStorageBlindAuctionBidSuccessfulIterator struct {
	Event *IMarketplaceStorageBlindAuctionBidSuccessful // Event containing the contract specifics and raw log

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
func (it *IMarketplaceStorageBlindAuctionBidSuccessfulIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IMarketplaceStorageBlindAuctionBidSuccessful)
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
		it.Event = new(IMarketplaceStorageBlindAuctionBidSuccessful)
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
func (it *IMarketplaceStorageBlindAuctionBidSuccessfulIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IMarketplaceStorageBlindAuctionBidSuccessfulIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IMarketplaceStorageBlindAuctionBidSuccessful represents a BlindAuctionBidSuccessful event raised by the IMarketplaceStorage contract.
type IMarketplaceStorageBlindAuctionBidSuccessful struct {
	Bidder         common.Address
	BlindAuctionId [32]byte
	BlindedBid     [32]byte
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterBlindAuctionBidSuccessful is a free log retrieval operation binding the contract event 0x42412bd3e1b349d0fc9c3518c85d09f611398d80842c5d3e36449b76dca01b6b.
//
// Solidity: event BlindAuctionBidSuccessful(address bidder, bytes32 blindAuctionId, bytes32 blindedBid)
func (_IMarketplaceStorage *IMarketplaceStorageFilterer) FilterBlindAuctionBidSuccessful(opts *bind.FilterOpts) (*IMarketplaceStorageBlindAuctionBidSuccessfulIterator, error) {

	logs, sub, err := _IMarketplaceStorage.contract.FilterLogs(opts, "BlindAuctionBidSuccessful")
	if err != nil {
		return nil, err
	}
	return &IMarketplaceStorageBlindAuctionBidSuccessfulIterator{contract: _IMarketplaceStorage.contract, event: "BlindAuctionBidSuccessful", logs: logs, sub: sub}, nil
}

// WatchBlindAuctionBidSuccessful is a free log subscription operation binding the contract event 0x42412bd3e1b349d0fc9c3518c85d09f611398d80842c5d3e36449b76dca01b6b.
//
// Solidity: event BlindAuctionBidSuccessful(address bidder, bytes32 blindAuctionId, bytes32 blindedBid)
func (_IMarketplaceStorage *IMarketplaceStorageFilterer) WatchBlindAuctionBidSuccessful(opts *bind.WatchOpts, sink chan<- *IMarketplaceStorageBlindAuctionBidSuccessful) (event.Subscription, error) {

	logs, sub, err := _IMarketplaceStorage.contract.WatchLogs(opts, "BlindAuctionBidSuccessful")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IMarketplaceStorageBlindAuctionBidSuccessful)
				if err := _IMarketplaceStorage.contract.UnpackLog(event, "BlindAuctionBidSuccessful", log); err != nil {
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

// ParseBlindAuctionBidSuccessful is a log parse operation binding the contract event 0x42412bd3e1b349d0fc9c3518c85d09f611398d80842c5d3e36449b76dca01b6b.
//
// Solidity: event BlindAuctionBidSuccessful(address bidder, bytes32 blindAuctionId, bytes32 blindedBid)
func (_IMarketplaceStorage *IMarketplaceStorageFilterer) ParseBlindAuctionBidSuccessful(log types.Log) (*IMarketplaceStorageBlindAuctionBidSuccessful, error) {
	event := new(IMarketplaceStorageBlindAuctionBidSuccessful)
	if err := _IMarketplaceStorage.contract.UnpackLog(event, "BlindAuctionBidSuccessful", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IMarketplaceStorageBlindAuctionCancelledIterator is returned from FilterBlindAuctionCancelled and is used to iterate over the raw logs and unpacked data for BlindAuctionCancelled events raised by the IMarketplaceStorage contract.
type IMarketplaceStorageBlindAuctionCancelledIterator struct {
	Event *IMarketplaceStorageBlindAuctionCancelled // Event containing the contract specifics and raw log

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
func (it *IMarketplaceStorageBlindAuctionCancelledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IMarketplaceStorageBlindAuctionCancelled)
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
		it.Event = new(IMarketplaceStorageBlindAuctionCancelled)
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
func (it *IMarketplaceStorageBlindAuctionCancelledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IMarketplaceStorageBlindAuctionCancelledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IMarketplaceStorageBlindAuctionCancelled represents a BlindAuctionCancelled event raised by the IMarketplaceStorage contract.
type IMarketplaceStorageBlindAuctionCancelled struct {
	Canceller      common.Address
	BlindAuctionId [32]byte
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterBlindAuctionCancelled is a free log retrieval operation binding the contract event 0x620e9f21de284210359a265815fc84d3f6f4686df7a3841c7ef048a50ea80216.
//
// Solidity: event BlindAuctionCancelled(address canceller, bytes32 blindAuctionId)
func (_IMarketplaceStorage *IMarketplaceStorageFilterer) FilterBlindAuctionCancelled(opts *bind.FilterOpts) (*IMarketplaceStorageBlindAuctionCancelledIterator, error) {

	logs, sub, err := _IMarketplaceStorage.contract.FilterLogs(opts, "BlindAuctionCancelled")
	if err != nil {
		return nil, err
	}
	return &IMarketplaceStorageBlindAuctionCancelledIterator{contract: _IMarketplaceStorage.contract, event: "BlindAuctionCancelled", logs: logs, sub: sub}, nil
}

// WatchBlindAuctionCancelled is a free log subscription operation binding the contract event 0x620e9f21de284210359a265815fc84d3f6f4686df7a3841c7ef048a50ea80216.
//
// Solidity: event BlindAuctionCancelled(address canceller, bytes32 blindAuctionId)
func (_IMarketplaceStorage *IMarketplaceStorageFilterer) WatchBlindAuctionCancelled(opts *bind.WatchOpts, sink chan<- *IMarketplaceStorageBlindAuctionCancelled) (event.Subscription, error) {

	logs, sub, err := _IMarketplaceStorage.contract.WatchLogs(opts, "BlindAuctionCancelled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IMarketplaceStorageBlindAuctionCancelled)
				if err := _IMarketplaceStorage.contract.UnpackLog(event, "BlindAuctionCancelled", log); err != nil {
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

// ParseBlindAuctionCancelled is a log parse operation binding the contract event 0x620e9f21de284210359a265815fc84d3f6f4686df7a3841c7ef048a50ea80216.
//
// Solidity: event BlindAuctionCancelled(address canceller, bytes32 blindAuctionId)
func (_IMarketplaceStorage *IMarketplaceStorageFilterer) ParseBlindAuctionCancelled(log types.Log) (*IMarketplaceStorageBlindAuctionCancelled, error) {
	event := new(IMarketplaceStorageBlindAuctionCancelled)
	if err := _IMarketplaceStorage.contract.UnpackLog(event, "BlindAuctionCancelled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IMarketplaceStorageBlindAuctionCreatedIterator is returned from FilterBlindAuctionCreated and is used to iterate over the raw logs and unpacked data for BlindAuctionCreated events raised by the IMarketplaceStorage contract.
type IMarketplaceStorageBlindAuctionCreatedIterator struct {
	Event *IMarketplaceStorageBlindAuctionCreated // Event containing the contract specifics and raw log

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
func (it *IMarketplaceStorageBlindAuctionCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IMarketplaceStorageBlindAuctionCreated)
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
		it.Event = new(IMarketplaceStorageBlindAuctionCreated)
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
func (it *IMarketplaceStorageBlindAuctionCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IMarketplaceStorageBlindAuctionCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IMarketplaceStorageBlindAuctionCreated represents a BlindAuctionCreated event raised by the IMarketplaceStorage contract.
type IMarketplaceStorageBlindAuctionCreated struct {
	Seller          common.Address
	NftAddress      common.Address
	BlindAuctionId  [32]byte
	AssetId         *big.Int
	BiddingEnd      *big.Int
	RevealEnd       *big.Int
	StartPriceInWei *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterBlindAuctionCreated is a free log retrieval operation binding the contract event 0xe64fd4ba0702f46ae298a3be33f16f223a5e06556977c42dee952f976abfed83.
//
// Solidity: event BlindAuctionCreated(address seller, address nftAddress, bytes32 blindAuctionId, uint256 assetId, uint256 biddingEnd, uint256 revealEnd, uint256 startPriceInWei)
func (_IMarketplaceStorage *IMarketplaceStorageFilterer) FilterBlindAuctionCreated(opts *bind.FilterOpts) (*IMarketplaceStorageBlindAuctionCreatedIterator, error) {

	logs, sub, err := _IMarketplaceStorage.contract.FilterLogs(opts, "BlindAuctionCreated")
	if err != nil {
		return nil, err
	}
	return &IMarketplaceStorageBlindAuctionCreatedIterator{contract: _IMarketplaceStorage.contract, event: "BlindAuctionCreated", logs: logs, sub: sub}, nil
}

// WatchBlindAuctionCreated is a free log subscription operation binding the contract event 0xe64fd4ba0702f46ae298a3be33f16f223a5e06556977c42dee952f976abfed83.
//
// Solidity: event BlindAuctionCreated(address seller, address nftAddress, bytes32 blindAuctionId, uint256 assetId, uint256 biddingEnd, uint256 revealEnd, uint256 startPriceInWei)
func (_IMarketplaceStorage *IMarketplaceStorageFilterer) WatchBlindAuctionCreated(opts *bind.WatchOpts, sink chan<- *IMarketplaceStorageBlindAuctionCreated) (event.Subscription, error) {

	logs, sub, err := _IMarketplaceStorage.contract.WatchLogs(opts, "BlindAuctionCreated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IMarketplaceStorageBlindAuctionCreated)
				if err := _IMarketplaceStorage.contract.UnpackLog(event, "BlindAuctionCreated", log); err != nil {
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

// ParseBlindAuctionCreated is a log parse operation binding the contract event 0xe64fd4ba0702f46ae298a3be33f16f223a5e06556977c42dee952f976abfed83.
//
// Solidity: event BlindAuctionCreated(address seller, address nftAddress, bytes32 blindAuctionId, uint256 assetId, uint256 biddingEnd, uint256 revealEnd, uint256 startPriceInWei)
func (_IMarketplaceStorage *IMarketplaceStorageFilterer) ParseBlindAuctionCreated(log types.Log) (*IMarketplaceStorageBlindAuctionCreated, error) {
	event := new(IMarketplaceStorageBlindAuctionCreated)
	if err := _IMarketplaceStorage.contract.UnpackLog(event, "BlindAuctionCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IMarketplaceStorageBlindAuctionEndedIterator is returned from FilterBlindAuctionEnded and is used to iterate over the raw logs and unpacked data for BlindAuctionEnded events raised by the IMarketplaceStorage contract.
type IMarketplaceStorageBlindAuctionEndedIterator struct {
	Event *IMarketplaceStorageBlindAuctionEnded // Event containing the contract specifics and raw log

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
func (it *IMarketplaceStorageBlindAuctionEndedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IMarketplaceStorageBlindAuctionEnded)
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
		it.Event = new(IMarketplaceStorageBlindAuctionEnded)
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
func (it *IMarketplaceStorageBlindAuctionEndedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IMarketplaceStorageBlindAuctionEndedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IMarketplaceStorageBlindAuctionEnded represents a BlindAuctionEnded event raised by the IMarketplaceStorage contract.
type IMarketplaceStorageBlindAuctionEnded struct {
	BlindAuctionId [32]byte
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterBlindAuctionEnded is a free log retrieval operation binding the contract event 0xbe3b74f797d49d96897f669a8b62fc0d77fa4cd6cf999a1b08bc31e7f4faf219.
//
// Solidity: event BlindAuctionEnded(bytes32 blindAuctionId)
func (_IMarketplaceStorage *IMarketplaceStorageFilterer) FilterBlindAuctionEnded(opts *bind.FilterOpts) (*IMarketplaceStorageBlindAuctionEndedIterator, error) {

	logs, sub, err := _IMarketplaceStorage.contract.FilterLogs(opts, "BlindAuctionEnded")
	if err != nil {
		return nil, err
	}
	return &IMarketplaceStorageBlindAuctionEndedIterator{contract: _IMarketplaceStorage.contract, event: "BlindAuctionEnded", logs: logs, sub: sub}, nil
}

// WatchBlindAuctionEnded is a free log subscription operation binding the contract event 0xbe3b74f797d49d96897f669a8b62fc0d77fa4cd6cf999a1b08bc31e7f4faf219.
//
// Solidity: event BlindAuctionEnded(bytes32 blindAuctionId)
func (_IMarketplaceStorage *IMarketplaceStorageFilterer) WatchBlindAuctionEnded(opts *bind.WatchOpts, sink chan<- *IMarketplaceStorageBlindAuctionEnded) (event.Subscription, error) {

	logs, sub, err := _IMarketplaceStorage.contract.WatchLogs(opts, "BlindAuctionEnded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IMarketplaceStorageBlindAuctionEnded)
				if err := _IMarketplaceStorage.contract.UnpackLog(event, "BlindAuctionEnded", log); err != nil {
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

// ParseBlindAuctionEnded is a log parse operation binding the contract event 0xbe3b74f797d49d96897f669a8b62fc0d77fa4cd6cf999a1b08bc31e7f4faf219.
//
// Solidity: event BlindAuctionEnded(bytes32 blindAuctionId)
func (_IMarketplaceStorage *IMarketplaceStorageFilterer) ParseBlindAuctionEnded(log types.Log) (*IMarketplaceStorageBlindAuctionEnded, error) {
	event := new(IMarketplaceStorageBlindAuctionEnded)
	if err := _IMarketplaceStorage.contract.UnpackLog(event, "BlindAuctionEnded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IMarketplaceStorageBlindAuctionRefundIterator is returned from FilterBlindAuctionRefund and is used to iterate over the raw logs and unpacked data for BlindAuctionRefund events raised by the IMarketplaceStorage contract.
type IMarketplaceStorageBlindAuctionRefundIterator struct {
	Event *IMarketplaceStorageBlindAuctionRefund // Event containing the contract specifics and raw log

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
func (it *IMarketplaceStorageBlindAuctionRefundIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IMarketplaceStorageBlindAuctionRefund)
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
		it.Event = new(IMarketplaceStorageBlindAuctionRefund)
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
func (it *IMarketplaceStorageBlindAuctionRefundIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IMarketplaceStorageBlindAuctionRefundIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IMarketplaceStorageBlindAuctionRefund represents a BlindAuctionRefund event raised by the IMarketplaceStorage contract.
type IMarketplaceStorageBlindAuctionRefund struct {
	Bidder         common.Address
	BlindAuctionId [32]byte
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterBlindAuctionRefund is a free log retrieval operation binding the contract event 0xadada2afe0e9306e1b93faf56e74b041719e68f722de4f550d5a57e39610b404.
//
// Solidity: event BlindAuctionRefund(address bidder, bytes32 blindAuctionId)
func (_IMarketplaceStorage *IMarketplaceStorageFilterer) FilterBlindAuctionRefund(opts *bind.FilterOpts) (*IMarketplaceStorageBlindAuctionRefundIterator, error) {

	logs, sub, err := _IMarketplaceStorage.contract.FilterLogs(opts, "BlindAuctionRefund")
	if err != nil {
		return nil, err
	}
	return &IMarketplaceStorageBlindAuctionRefundIterator{contract: _IMarketplaceStorage.contract, event: "BlindAuctionRefund", logs: logs, sub: sub}, nil
}

// WatchBlindAuctionRefund is a free log subscription operation binding the contract event 0xadada2afe0e9306e1b93faf56e74b041719e68f722de4f550d5a57e39610b404.
//
// Solidity: event BlindAuctionRefund(address bidder, bytes32 blindAuctionId)
func (_IMarketplaceStorage *IMarketplaceStorageFilterer) WatchBlindAuctionRefund(opts *bind.WatchOpts, sink chan<- *IMarketplaceStorageBlindAuctionRefund) (event.Subscription, error) {

	logs, sub, err := _IMarketplaceStorage.contract.WatchLogs(opts, "BlindAuctionRefund")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IMarketplaceStorageBlindAuctionRefund)
				if err := _IMarketplaceStorage.contract.UnpackLog(event, "BlindAuctionRefund", log); err != nil {
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

// ParseBlindAuctionRefund is a log parse operation binding the contract event 0xadada2afe0e9306e1b93faf56e74b041719e68f722de4f550d5a57e39610b404.
//
// Solidity: event BlindAuctionRefund(address bidder, bytes32 blindAuctionId)
func (_IMarketplaceStorage *IMarketplaceStorageFilterer) ParseBlindAuctionRefund(log types.Log) (*IMarketplaceStorageBlindAuctionRefund, error) {
	event := new(IMarketplaceStorageBlindAuctionRefund)
	if err := _IMarketplaceStorage.contract.UnpackLog(event, "BlindAuctionRefund", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IMarketplaceStorageBlindAuctionSuccessfulIterator is returned from FilterBlindAuctionSuccessful and is used to iterate over the raw logs and unpacked data for BlindAuctionSuccessful events raised by the IMarketplaceStorage contract.
type IMarketplaceStorageBlindAuctionSuccessfulIterator struct {
	Event *IMarketplaceStorageBlindAuctionSuccessful // Event containing the contract specifics and raw log

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
func (it *IMarketplaceStorageBlindAuctionSuccessfulIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IMarketplaceStorageBlindAuctionSuccessful)
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
		it.Event = new(IMarketplaceStorageBlindAuctionSuccessful)
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
func (it *IMarketplaceStorageBlindAuctionSuccessfulIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IMarketplaceStorageBlindAuctionSuccessfulIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IMarketplaceStorageBlindAuctionSuccessful represents a BlindAuctionSuccessful event raised by the IMarketplaceStorage contract.
type IMarketplaceStorageBlindAuctionSuccessful struct {
	Seller         common.Address
	Buyer          common.Address
	BlindAuctionId [32]byte
	TotalPrice     *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterBlindAuctionSuccessful is a free log retrieval operation binding the contract event 0x2258179acb0769137f4ff532afd72c91199ae597abded9936acb61a2e66fc142.
//
// Solidity: event BlindAuctionSuccessful(address seller, address buyer, bytes32 blindAuctionId, uint256 totalPrice)
func (_IMarketplaceStorage *IMarketplaceStorageFilterer) FilterBlindAuctionSuccessful(opts *bind.FilterOpts) (*IMarketplaceStorageBlindAuctionSuccessfulIterator, error) {

	logs, sub, err := _IMarketplaceStorage.contract.FilterLogs(opts, "BlindAuctionSuccessful")
	if err != nil {
		return nil, err
	}
	return &IMarketplaceStorageBlindAuctionSuccessfulIterator{contract: _IMarketplaceStorage.contract, event: "BlindAuctionSuccessful", logs: logs, sub: sub}, nil
}

// WatchBlindAuctionSuccessful is a free log subscription operation binding the contract event 0x2258179acb0769137f4ff532afd72c91199ae597abded9936acb61a2e66fc142.
//
// Solidity: event BlindAuctionSuccessful(address seller, address buyer, bytes32 blindAuctionId, uint256 totalPrice)
func (_IMarketplaceStorage *IMarketplaceStorageFilterer) WatchBlindAuctionSuccessful(opts *bind.WatchOpts, sink chan<- *IMarketplaceStorageBlindAuctionSuccessful) (event.Subscription, error) {

	logs, sub, err := _IMarketplaceStorage.contract.WatchLogs(opts, "BlindAuctionSuccessful")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IMarketplaceStorageBlindAuctionSuccessful)
				if err := _IMarketplaceStorage.contract.UnpackLog(event, "BlindAuctionSuccessful", log); err != nil {
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

// ParseBlindAuctionSuccessful is a log parse operation binding the contract event 0x2258179acb0769137f4ff532afd72c91199ae597abded9936acb61a2e66fc142.
//
// Solidity: event BlindAuctionSuccessful(address seller, address buyer, bytes32 blindAuctionId, uint256 totalPrice)
func (_IMarketplaceStorage *IMarketplaceStorageFilterer) ParseBlindAuctionSuccessful(log types.Log) (*IMarketplaceStorageBlindAuctionSuccessful, error) {
	event := new(IMarketplaceStorageBlindAuctionSuccessful)
	if err := _IMarketplaceStorage.contract.UnpackLog(event, "BlindAuctionSuccessful", log); err != nil {
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
	Who common.Address
	Id  [32]byte
	Raw types.Log // Blockchain specific contextual infos
}

// FilterOrderCancelled is a free log retrieval operation binding the contract event 0x35974c4230d53fb4c6e8553fd900c88ba92747dbc689a79bcd6ba755cb936985.
//
// Solidity: event OrderCancelled(address who, bytes32 id)
func (_IMarketplaceStorage *IMarketplaceStorageFilterer) FilterOrderCancelled(opts *bind.FilterOpts) (*IMarketplaceStorageOrderCancelledIterator, error) {

	logs, sub, err := _IMarketplaceStorage.contract.FilterLogs(opts, "OrderCancelled")
	if err != nil {
		return nil, err
	}
	return &IMarketplaceStorageOrderCancelledIterator{contract: _IMarketplaceStorage.contract, event: "OrderCancelled", logs: logs, sub: sub}, nil
}

// WatchOrderCancelled is a free log subscription operation binding the contract event 0x35974c4230d53fb4c6e8553fd900c88ba92747dbc689a79bcd6ba755cb936985.
//
// Solidity: event OrderCancelled(address who, bytes32 id)
func (_IMarketplaceStorage *IMarketplaceStorageFilterer) WatchOrderCancelled(opts *bind.WatchOpts, sink chan<- *IMarketplaceStorageOrderCancelled) (event.Subscription, error) {

	logs, sub, err := _IMarketplaceStorage.contract.WatchLogs(opts, "OrderCancelled")
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

// ParseOrderCancelled is a log parse operation binding the contract event 0x35974c4230d53fb4c6e8553fd900c88ba92747dbc689a79bcd6ba755cb936985.
//
// Solidity: event OrderCancelled(address who, bytes32 id)
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
// Solidity: event OrderCreated(bytes32 orderId, uint256 assetId, address seller, address nftAddress, uint256 priceInWei, uint256 expiredAt)
func (_IMarketplaceStorage *IMarketplaceStorageFilterer) FilterOrderCreated(opts *bind.FilterOpts) (*IMarketplaceStorageOrderCreatedIterator, error) {

	logs, sub, err := _IMarketplaceStorage.contract.FilterLogs(opts, "OrderCreated")
	if err != nil {
		return nil, err
	}
	return &IMarketplaceStorageOrderCreatedIterator{contract: _IMarketplaceStorage.contract, event: "OrderCreated", logs: logs, sub: sub}, nil
}

// WatchOrderCreated is a free log subscription operation binding the contract event 0x84c66c3f7ba4b390e20e8e8233e2a516f3ce34a72749e4f12bd010dfba238039.
//
// Solidity: event OrderCreated(bytes32 orderId, uint256 assetId, address seller, address nftAddress, uint256 priceInWei, uint256 expiredAt)
func (_IMarketplaceStorage *IMarketplaceStorageFilterer) WatchOrderCreated(opts *bind.WatchOpts, sink chan<- *IMarketplaceStorageOrderCreated) (event.Subscription, error) {

	logs, sub, err := _IMarketplaceStorage.contract.WatchLogs(opts, "OrderCreated")
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
// Solidity: event OrderCreated(bytes32 orderId, uint256 assetId, address seller, address nftAddress, uint256 priceInWei, uint256 expiredAt)
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
	Id     [32]byte
	Buyer  common.Address
	Seller common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterOrderSuccessful is a free log retrieval operation binding the contract event 0xa3fe2a6ffb1ad0de1c1ee8e3513d205b99980c40af9de82e502095245debabbb.
//
// Solidity: event OrderSuccessful(bytes32 id, address buyer, address seller)
func (_IMarketplaceStorage *IMarketplaceStorageFilterer) FilterOrderSuccessful(opts *bind.FilterOpts) (*IMarketplaceStorageOrderSuccessfulIterator, error) {

	logs, sub, err := _IMarketplaceStorage.contract.FilterLogs(opts, "OrderSuccessful")
	if err != nil {
		return nil, err
	}
	return &IMarketplaceStorageOrderSuccessfulIterator{contract: _IMarketplaceStorage.contract, event: "OrderSuccessful", logs: logs, sub: sub}, nil
}

// WatchOrderSuccessful is a free log subscription operation binding the contract event 0xa3fe2a6ffb1ad0de1c1ee8e3513d205b99980c40af9de82e502095245debabbb.
//
// Solidity: event OrderSuccessful(bytes32 id, address buyer, address seller)
func (_IMarketplaceStorage *IMarketplaceStorageFilterer) WatchOrderSuccessful(opts *bind.WatchOpts, sink chan<- *IMarketplaceStorageOrderSuccessful) (event.Subscription, error) {

	logs, sub, err := _IMarketplaceStorage.contract.WatchLogs(opts, "OrderSuccessful")
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

// ParseOrderSuccessful is a log parse operation binding the contract event 0xa3fe2a6ffb1ad0de1c1ee8e3513d205b99980c40af9de82e502095245debabbb.
//
// Solidity: event OrderSuccessful(bytes32 id, address buyer, address seller)
func (_IMarketplaceStorage *IMarketplaceStorageFilterer) ParseOrderSuccessful(log types.Log) (*IMarketplaceStorageOrderSuccessful, error) {
	event := new(IMarketplaceStorageOrderSuccessful)
	if err := _IMarketplaceStorage.contract.UnpackLog(event, "OrderSuccessful", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IMarketplaceStoragePublicAuctionBidSuccessfulIterator is returned from FilterPublicAuctionBidSuccessful and is used to iterate over the raw logs and unpacked data for PublicAuctionBidSuccessful events raised by the IMarketplaceStorage contract.
type IMarketplaceStoragePublicAuctionBidSuccessfulIterator struct {
	Event *IMarketplaceStoragePublicAuctionBidSuccessful // Event containing the contract specifics and raw log

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
func (it *IMarketplaceStoragePublicAuctionBidSuccessfulIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IMarketplaceStoragePublicAuctionBidSuccessful)
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
		it.Event = new(IMarketplaceStoragePublicAuctionBidSuccessful)
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
func (it *IMarketplaceStoragePublicAuctionBidSuccessfulIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IMarketplaceStoragePublicAuctionBidSuccessfulIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IMarketplaceStoragePublicAuctionBidSuccessful represents a PublicAuctionBidSuccessful event raised by the IMarketplaceStorage contract.
type IMarketplaceStoragePublicAuctionBidSuccessful struct {
	Bidder          common.Address
	PublicAuctionId [32]byte
	BidValue        *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterPublicAuctionBidSuccessful is a free log retrieval operation binding the contract event 0x9f541988b3aac74ac43a6f4b330ceed668244b74461b90850d199a38a172f7c2.
//
// Solidity: event PublicAuctionBidSuccessful(address bidder, bytes32 publicAuctionId, uint256 bidValue)
func (_IMarketplaceStorage *IMarketplaceStorageFilterer) FilterPublicAuctionBidSuccessful(opts *bind.FilterOpts) (*IMarketplaceStoragePublicAuctionBidSuccessfulIterator, error) {

	logs, sub, err := _IMarketplaceStorage.contract.FilterLogs(opts, "PublicAuctionBidSuccessful")
	if err != nil {
		return nil, err
	}
	return &IMarketplaceStoragePublicAuctionBidSuccessfulIterator{contract: _IMarketplaceStorage.contract, event: "PublicAuctionBidSuccessful", logs: logs, sub: sub}, nil
}

// WatchPublicAuctionBidSuccessful is a free log subscription operation binding the contract event 0x9f541988b3aac74ac43a6f4b330ceed668244b74461b90850d199a38a172f7c2.
//
// Solidity: event PublicAuctionBidSuccessful(address bidder, bytes32 publicAuctionId, uint256 bidValue)
func (_IMarketplaceStorage *IMarketplaceStorageFilterer) WatchPublicAuctionBidSuccessful(opts *bind.WatchOpts, sink chan<- *IMarketplaceStoragePublicAuctionBidSuccessful) (event.Subscription, error) {

	logs, sub, err := _IMarketplaceStorage.contract.WatchLogs(opts, "PublicAuctionBidSuccessful")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IMarketplaceStoragePublicAuctionBidSuccessful)
				if err := _IMarketplaceStorage.contract.UnpackLog(event, "PublicAuctionBidSuccessful", log); err != nil {
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

// ParsePublicAuctionBidSuccessful is a log parse operation binding the contract event 0x9f541988b3aac74ac43a6f4b330ceed668244b74461b90850d199a38a172f7c2.
//
// Solidity: event PublicAuctionBidSuccessful(address bidder, bytes32 publicAuctionId, uint256 bidValue)
func (_IMarketplaceStorage *IMarketplaceStorageFilterer) ParsePublicAuctionBidSuccessful(log types.Log) (*IMarketplaceStoragePublicAuctionBidSuccessful, error) {
	event := new(IMarketplaceStoragePublicAuctionBidSuccessful)
	if err := _IMarketplaceStorage.contract.UnpackLog(event, "PublicAuctionBidSuccessful", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IMarketplaceStoragePublicAuctionCancelledIterator is returned from FilterPublicAuctionCancelled and is used to iterate over the raw logs and unpacked data for PublicAuctionCancelled events raised by the IMarketplaceStorage contract.
type IMarketplaceStoragePublicAuctionCancelledIterator struct {
	Event *IMarketplaceStoragePublicAuctionCancelled // Event containing the contract specifics and raw log

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
func (it *IMarketplaceStoragePublicAuctionCancelledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IMarketplaceStoragePublicAuctionCancelled)
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
		it.Event = new(IMarketplaceStoragePublicAuctionCancelled)
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
func (it *IMarketplaceStoragePublicAuctionCancelledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IMarketplaceStoragePublicAuctionCancelledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IMarketplaceStoragePublicAuctionCancelled represents a PublicAuctionCancelled event raised by the IMarketplaceStorage contract.
type IMarketplaceStoragePublicAuctionCancelled struct {
	Canceller       common.Address
	PublicAuctionId [32]byte
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterPublicAuctionCancelled is a free log retrieval operation binding the contract event 0x87cbb8dceaaacfbe7c9d99abf2a9c85f0249d860e329f60869079b204c7ad00d.
//
// Solidity: event PublicAuctionCancelled(address canceller, bytes32 publicAuctionId)
func (_IMarketplaceStorage *IMarketplaceStorageFilterer) FilterPublicAuctionCancelled(opts *bind.FilterOpts) (*IMarketplaceStoragePublicAuctionCancelledIterator, error) {

	logs, sub, err := _IMarketplaceStorage.contract.FilterLogs(opts, "PublicAuctionCancelled")
	if err != nil {
		return nil, err
	}
	return &IMarketplaceStoragePublicAuctionCancelledIterator{contract: _IMarketplaceStorage.contract, event: "PublicAuctionCancelled", logs: logs, sub: sub}, nil
}

// WatchPublicAuctionCancelled is a free log subscription operation binding the contract event 0x87cbb8dceaaacfbe7c9d99abf2a9c85f0249d860e329f60869079b204c7ad00d.
//
// Solidity: event PublicAuctionCancelled(address canceller, bytes32 publicAuctionId)
func (_IMarketplaceStorage *IMarketplaceStorageFilterer) WatchPublicAuctionCancelled(opts *bind.WatchOpts, sink chan<- *IMarketplaceStoragePublicAuctionCancelled) (event.Subscription, error) {

	logs, sub, err := _IMarketplaceStorage.contract.WatchLogs(opts, "PublicAuctionCancelled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IMarketplaceStoragePublicAuctionCancelled)
				if err := _IMarketplaceStorage.contract.UnpackLog(event, "PublicAuctionCancelled", log); err != nil {
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

// ParsePublicAuctionCancelled is a log parse operation binding the contract event 0x87cbb8dceaaacfbe7c9d99abf2a9c85f0249d860e329f60869079b204c7ad00d.
//
// Solidity: event PublicAuctionCancelled(address canceller, bytes32 publicAuctionId)
func (_IMarketplaceStorage *IMarketplaceStorageFilterer) ParsePublicAuctionCancelled(log types.Log) (*IMarketplaceStoragePublicAuctionCancelled, error) {
	event := new(IMarketplaceStoragePublicAuctionCancelled)
	if err := _IMarketplaceStorage.contract.UnpackLog(event, "PublicAuctionCancelled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IMarketplaceStoragePublicAuctionCreatedIterator is returned from FilterPublicAuctionCreated and is used to iterate over the raw logs and unpacked data for PublicAuctionCreated events raised by the IMarketplaceStorage contract.
type IMarketplaceStoragePublicAuctionCreatedIterator struct {
	Event *IMarketplaceStoragePublicAuctionCreated // Event containing the contract specifics and raw log

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
func (it *IMarketplaceStoragePublicAuctionCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IMarketplaceStoragePublicAuctionCreated)
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
		it.Event = new(IMarketplaceStoragePublicAuctionCreated)
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
func (it *IMarketplaceStoragePublicAuctionCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IMarketplaceStoragePublicAuctionCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IMarketplaceStoragePublicAuctionCreated represents a PublicAuctionCreated event raised by the IMarketplaceStorage contract.
type IMarketplaceStoragePublicAuctionCreated struct {
	Seller          common.Address
	NftAddress      common.Address
	PublicAuctionId [32]byte
	AssetId         *big.Int
	BiddingEnd      *big.Int
	StartPriceInWei *big.Int
	MinIncrement    *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterPublicAuctionCreated is a free log retrieval operation binding the contract event 0xbdca6148a24d8e6e2d4ced0a6e168095869e61ea26b82332620abe8ba3ba5bd9.
//
// Solidity: event PublicAuctionCreated(address seller, address nftAddress, bytes32 publicAuctionId, uint256 assetId, uint256 biddingEnd, uint256 startPriceInWei, uint256 minIncrement)
func (_IMarketplaceStorage *IMarketplaceStorageFilterer) FilterPublicAuctionCreated(opts *bind.FilterOpts) (*IMarketplaceStoragePublicAuctionCreatedIterator, error) {

	logs, sub, err := _IMarketplaceStorage.contract.FilterLogs(opts, "PublicAuctionCreated")
	if err != nil {
		return nil, err
	}
	return &IMarketplaceStoragePublicAuctionCreatedIterator{contract: _IMarketplaceStorage.contract, event: "PublicAuctionCreated", logs: logs, sub: sub}, nil
}

// WatchPublicAuctionCreated is a free log subscription operation binding the contract event 0xbdca6148a24d8e6e2d4ced0a6e168095869e61ea26b82332620abe8ba3ba5bd9.
//
// Solidity: event PublicAuctionCreated(address seller, address nftAddress, bytes32 publicAuctionId, uint256 assetId, uint256 biddingEnd, uint256 startPriceInWei, uint256 minIncrement)
func (_IMarketplaceStorage *IMarketplaceStorageFilterer) WatchPublicAuctionCreated(opts *bind.WatchOpts, sink chan<- *IMarketplaceStoragePublicAuctionCreated) (event.Subscription, error) {

	logs, sub, err := _IMarketplaceStorage.contract.WatchLogs(opts, "PublicAuctionCreated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IMarketplaceStoragePublicAuctionCreated)
				if err := _IMarketplaceStorage.contract.UnpackLog(event, "PublicAuctionCreated", log); err != nil {
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

// ParsePublicAuctionCreated is a log parse operation binding the contract event 0xbdca6148a24d8e6e2d4ced0a6e168095869e61ea26b82332620abe8ba3ba5bd9.
//
// Solidity: event PublicAuctionCreated(address seller, address nftAddress, bytes32 publicAuctionId, uint256 assetId, uint256 biddingEnd, uint256 startPriceInWei, uint256 minIncrement)
func (_IMarketplaceStorage *IMarketplaceStorageFilterer) ParsePublicAuctionCreated(log types.Log) (*IMarketplaceStoragePublicAuctionCreated, error) {
	event := new(IMarketplaceStoragePublicAuctionCreated)
	if err := _IMarketplaceStorage.contract.UnpackLog(event, "PublicAuctionCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IMarketplaceStoragePublicAuctionEndedIterator is returned from FilterPublicAuctionEnded and is used to iterate over the raw logs and unpacked data for PublicAuctionEnded events raised by the IMarketplaceStorage contract.
type IMarketplaceStoragePublicAuctionEndedIterator struct {
	Event *IMarketplaceStoragePublicAuctionEnded // Event containing the contract specifics and raw log

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
func (it *IMarketplaceStoragePublicAuctionEndedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IMarketplaceStoragePublicAuctionEnded)
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
		it.Event = new(IMarketplaceStoragePublicAuctionEnded)
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
func (it *IMarketplaceStoragePublicAuctionEndedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IMarketplaceStoragePublicAuctionEndedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IMarketplaceStoragePublicAuctionEnded represents a PublicAuctionEnded event raised by the IMarketplaceStorage contract.
type IMarketplaceStoragePublicAuctionEnded struct {
	PublicAuctionId [32]byte
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterPublicAuctionEnded is a free log retrieval operation binding the contract event 0x156c2754a62667a51625de4ceb04df6640d97d06de8c89bd0bbb33307f144e42.
//
// Solidity: event PublicAuctionEnded(bytes32 publicAuctionId)
func (_IMarketplaceStorage *IMarketplaceStorageFilterer) FilterPublicAuctionEnded(opts *bind.FilterOpts) (*IMarketplaceStoragePublicAuctionEndedIterator, error) {

	logs, sub, err := _IMarketplaceStorage.contract.FilterLogs(opts, "PublicAuctionEnded")
	if err != nil {
		return nil, err
	}
	return &IMarketplaceStoragePublicAuctionEndedIterator{contract: _IMarketplaceStorage.contract, event: "PublicAuctionEnded", logs: logs, sub: sub}, nil
}

// WatchPublicAuctionEnded is a free log subscription operation binding the contract event 0x156c2754a62667a51625de4ceb04df6640d97d06de8c89bd0bbb33307f144e42.
//
// Solidity: event PublicAuctionEnded(bytes32 publicAuctionId)
func (_IMarketplaceStorage *IMarketplaceStorageFilterer) WatchPublicAuctionEnded(opts *bind.WatchOpts, sink chan<- *IMarketplaceStoragePublicAuctionEnded) (event.Subscription, error) {

	logs, sub, err := _IMarketplaceStorage.contract.WatchLogs(opts, "PublicAuctionEnded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IMarketplaceStoragePublicAuctionEnded)
				if err := _IMarketplaceStorage.contract.UnpackLog(event, "PublicAuctionEnded", log); err != nil {
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

// ParsePublicAuctionEnded is a log parse operation binding the contract event 0x156c2754a62667a51625de4ceb04df6640d97d06de8c89bd0bbb33307f144e42.
//
// Solidity: event PublicAuctionEnded(bytes32 publicAuctionId)
func (_IMarketplaceStorage *IMarketplaceStorageFilterer) ParsePublicAuctionEnded(log types.Log) (*IMarketplaceStoragePublicAuctionEnded, error) {
	event := new(IMarketplaceStoragePublicAuctionEnded)
	if err := _IMarketplaceStorage.contract.UnpackLog(event, "PublicAuctionEnded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IMarketplaceStoragePublicAuctionRefundIterator is returned from FilterPublicAuctionRefund and is used to iterate over the raw logs and unpacked data for PublicAuctionRefund events raised by the IMarketplaceStorage contract.
type IMarketplaceStoragePublicAuctionRefundIterator struct {
	Event *IMarketplaceStoragePublicAuctionRefund // Event containing the contract specifics and raw log

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
func (it *IMarketplaceStoragePublicAuctionRefundIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IMarketplaceStoragePublicAuctionRefund)
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
		it.Event = new(IMarketplaceStoragePublicAuctionRefund)
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
func (it *IMarketplaceStoragePublicAuctionRefundIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IMarketplaceStoragePublicAuctionRefundIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IMarketplaceStoragePublicAuctionRefund represents a PublicAuctionRefund event raised by the IMarketplaceStorage contract.
type IMarketplaceStoragePublicAuctionRefund struct {
	Bidder          common.Address
	PublicAuctionId [32]byte
	Value           *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterPublicAuctionRefund is a free log retrieval operation binding the contract event 0x4fac96064e645402e0d4854b9549caba0169e3252c0e08a305eb60f498c88911.
//
// Solidity: event PublicAuctionRefund(address bidder, bytes32 publicAuctionId, uint256 value)
func (_IMarketplaceStorage *IMarketplaceStorageFilterer) FilterPublicAuctionRefund(opts *bind.FilterOpts) (*IMarketplaceStoragePublicAuctionRefundIterator, error) {

	logs, sub, err := _IMarketplaceStorage.contract.FilterLogs(opts, "PublicAuctionRefund")
	if err != nil {
		return nil, err
	}
	return &IMarketplaceStoragePublicAuctionRefundIterator{contract: _IMarketplaceStorage.contract, event: "PublicAuctionRefund", logs: logs, sub: sub}, nil
}

// WatchPublicAuctionRefund is a free log subscription operation binding the contract event 0x4fac96064e645402e0d4854b9549caba0169e3252c0e08a305eb60f498c88911.
//
// Solidity: event PublicAuctionRefund(address bidder, bytes32 publicAuctionId, uint256 value)
func (_IMarketplaceStorage *IMarketplaceStorageFilterer) WatchPublicAuctionRefund(opts *bind.WatchOpts, sink chan<- *IMarketplaceStoragePublicAuctionRefund) (event.Subscription, error) {

	logs, sub, err := _IMarketplaceStorage.contract.WatchLogs(opts, "PublicAuctionRefund")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IMarketplaceStoragePublicAuctionRefund)
				if err := _IMarketplaceStorage.contract.UnpackLog(event, "PublicAuctionRefund", log); err != nil {
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

// ParsePublicAuctionRefund is a log parse operation binding the contract event 0x4fac96064e645402e0d4854b9549caba0169e3252c0e08a305eb60f498c88911.
//
// Solidity: event PublicAuctionRefund(address bidder, bytes32 publicAuctionId, uint256 value)
func (_IMarketplaceStorage *IMarketplaceStorageFilterer) ParsePublicAuctionRefund(log types.Log) (*IMarketplaceStoragePublicAuctionRefund, error) {
	event := new(IMarketplaceStoragePublicAuctionRefund)
	if err := _IMarketplaceStorage.contract.UnpackLog(event, "PublicAuctionRefund", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IMarketplaceStoragePublicAuctionSuccessfulIterator is returned from FilterPublicAuctionSuccessful and is used to iterate over the raw logs and unpacked data for PublicAuctionSuccessful events raised by the IMarketplaceStorage contract.
type IMarketplaceStoragePublicAuctionSuccessfulIterator struct {
	Event *IMarketplaceStoragePublicAuctionSuccessful // Event containing the contract specifics and raw log

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
func (it *IMarketplaceStoragePublicAuctionSuccessfulIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IMarketplaceStoragePublicAuctionSuccessful)
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
		it.Event = new(IMarketplaceStoragePublicAuctionSuccessful)
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
func (it *IMarketplaceStoragePublicAuctionSuccessfulIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IMarketplaceStoragePublicAuctionSuccessfulIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IMarketplaceStoragePublicAuctionSuccessful represents a PublicAuctionSuccessful event raised by the IMarketplaceStorage contract.
type IMarketplaceStoragePublicAuctionSuccessful struct {
	Seller          common.Address
	Buyer           common.Address
	PublicAuctionId [32]byte
	TotalPrice      *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterPublicAuctionSuccessful is a free log retrieval operation binding the contract event 0xd5133f9e5285bd0635ccfe53ae8978b96ef34a7bf16a0a696a411dc669322cb6.
//
// Solidity: event PublicAuctionSuccessful(address seller, address buyer, bytes32 publicAuctionId, uint256 totalPrice)
func (_IMarketplaceStorage *IMarketplaceStorageFilterer) FilterPublicAuctionSuccessful(opts *bind.FilterOpts) (*IMarketplaceStoragePublicAuctionSuccessfulIterator, error) {

	logs, sub, err := _IMarketplaceStorage.contract.FilterLogs(opts, "PublicAuctionSuccessful")
	if err != nil {
		return nil, err
	}
	return &IMarketplaceStoragePublicAuctionSuccessfulIterator{contract: _IMarketplaceStorage.contract, event: "PublicAuctionSuccessful", logs: logs, sub: sub}, nil
}

// WatchPublicAuctionSuccessful is a free log subscription operation binding the contract event 0xd5133f9e5285bd0635ccfe53ae8978b96ef34a7bf16a0a696a411dc669322cb6.
//
// Solidity: event PublicAuctionSuccessful(address seller, address buyer, bytes32 publicAuctionId, uint256 totalPrice)
func (_IMarketplaceStorage *IMarketplaceStorageFilterer) WatchPublicAuctionSuccessful(opts *bind.WatchOpts, sink chan<- *IMarketplaceStoragePublicAuctionSuccessful) (event.Subscription, error) {

	logs, sub, err := _IMarketplaceStorage.contract.WatchLogs(opts, "PublicAuctionSuccessful")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IMarketplaceStoragePublicAuctionSuccessful)
				if err := _IMarketplaceStorage.contract.UnpackLog(event, "PublicAuctionSuccessful", log); err != nil {
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

// ParsePublicAuctionSuccessful is a log parse operation binding the contract event 0xd5133f9e5285bd0635ccfe53ae8978b96ef34a7bf16a0a696a411dc669322cb6.
//
// Solidity: event PublicAuctionSuccessful(address seller, address buyer, bytes32 publicAuctionId, uint256 totalPrice)
func (_IMarketplaceStorage *IMarketplaceStorageFilterer) ParsePublicAuctionSuccessful(log types.Log) (*IMarketplaceStoragePublicAuctionSuccessful, error) {
	event := new(IMarketplaceStoragePublicAuctionSuccessful)
	if err := _IMarketplaceStorage.contract.UnpackLog(event, "PublicAuctionSuccessful", log); err != nil {
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
	Seller         common.Address
	BlindAuctionId [32]byte
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterRevealSuccessful is a free log retrieval operation binding the contract event 0x61d4ab60470c97b371c5f91be47873283386bb3af4de728e8795ef5318e013d6.
//
// Solidity: event RevealSuccessful(address seller, bytes32 blindAuctionId)
func (_IMarketplaceStorage *IMarketplaceStorageFilterer) FilterRevealSuccessful(opts *bind.FilterOpts) (*IMarketplaceStorageRevealSuccessfulIterator, error) {

	logs, sub, err := _IMarketplaceStorage.contract.FilterLogs(opts, "RevealSuccessful")
	if err != nil {
		return nil, err
	}
	return &IMarketplaceStorageRevealSuccessfulIterator{contract: _IMarketplaceStorage.contract, event: "RevealSuccessful", logs: logs, sub: sub}, nil
}

// WatchRevealSuccessful is a free log subscription operation binding the contract event 0x61d4ab60470c97b371c5f91be47873283386bb3af4de728e8795ef5318e013d6.
//
// Solidity: event RevealSuccessful(address seller, bytes32 blindAuctionId)
func (_IMarketplaceStorage *IMarketplaceStorageFilterer) WatchRevealSuccessful(opts *bind.WatchOpts, sink chan<- *IMarketplaceStorageRevealSuccessful) (event.Subscription, error) {

	logs, sub, err := _IMarketplaceStorage.contract.WatchLogs(opts, "RevealSuccessful")
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

// ParseRevealSuccessful is a log parse operation binding the contract event 0x61d4ab60470c97b371c5f91be47873283386bb3af4de728e8795ef5318e013d6.
//
// Solidity: event RevealSuccessful(address seller, bytes32 blindAuctionId)
func (_IMarketplaceStorage *IMarketplaceStorageFilterer) ParseRevealSuccessful(log types.Log) (*IMarketplaceStorageRevealSuccessful, error) {
	event := new(IMarketplaceStorageRevealSuccessful)
	if err := _IMarketplaceStorage.contract.UnpackLog(event, "RevealSuccessful", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
