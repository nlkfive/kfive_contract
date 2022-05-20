// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package MarketplaceStorage

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

// MarketplaceStorageMetaData contains all meta data concerning the MarketplaceStorage contract.
var MarketplaceStorageMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"AssetNotExisted\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AssetUnvailable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AuctionAlreadyEnded\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidMkpSender\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotRunning\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"}],\"name\":\"TooEarly\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"}],\"name\":\"TooLate\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"who\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"auctionId\",\"type\":\"bytes32\"}],\"name\":\"AuctionCancelled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"nftAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"auctionId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"assetId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"biddingEnd\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"revealEnd\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"startPriceInWei\",\"type\":\"uint256\"}],\"name\":\"AuctionCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"auctionId\",\"type\":\"bytes32\"}],\"name\":\"AuctionEnded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"bidder\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"auctionId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"deposit\",\"type\":\"uint256\"}],\"name\":\"AuctionRefund\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"buyer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"auctionId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalPrice\",\"type\":\"uint256\"}],\"name\":\"AuctionSuccessful\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"bidder\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"auctionId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"blindedBid\",\"type\":\"bytes32\"}],\"name\":\"BidSuccessful\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"who\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"}],\"name\":\"OrderCancelled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"orderId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"assetId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"nftAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"priceInWei\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"expiredAt\",\"type\":\"uint256\"}],\"name\":\"OrderCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"buyer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"}],\"name\":\"OrderSuccessful\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"fake\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"revealer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"auctionId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"RevealFailed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"fake\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"revealer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"auctionId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"blindedBid\",\"type\":\"bytes32\"}],\"name\":\"RevealSuccessful\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PAUSER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getRoleMember\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleMemberCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"orderByNftAsset\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"orderId\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"nftAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expiredAt\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_auctionMarketplace\",\"type\":\"address\"}],\"name\":\"updateAuctionMarketplace\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_orderMarketplace\",\"type\":\"address\"}],\"name\":\"updateOrderMarketplace\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"nftAsset\",\"type\":\"bytes32\"}],\"name\":\"assetIsAvailable\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"auctionId\",\"type\":\"bytes32\"}],\"name\":\"auctionIsExisted\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"nftAsset\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"auctionId\",\"type\":\"bytes32\"}],\"name\":\"auctionIsEnded\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"nftAsset\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"auctionId\",\"type\":\"bytes32\"}],\"name\":\"auctionIsRunning\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"assetOwner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"nftAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"assetId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"auctionId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"biddingEnd\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"revealEnd\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startPriceInWei\",\"type\":\"uint256\"}],\"name\":\"createAuction\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"auctionId\",\"type\":\"bytes32\"}],\"name\":\"getAuction\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"highestBidder\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"biddingEnd\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"revealEnd\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"highestBid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startPrice\",\"type\":\"uint256\"}],\"internalType\":\"structIAuction.Auction\",\"name\":\"order\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"nftAsset\",\"type\":\"bytes32\"}],\"name\":\"auctionEnded\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"bidder\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"highestBid\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"auctionId\",\"type\":\"bytes32\"}],\"name\":\"updateHighestBid\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"nftAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"assetId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"orderId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"priceInWei\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expiredAt\",\"type\":\"uint256\"}],\"name\":\"createOrder\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"nftAsset\",\"type\":\"bytes32\"}],\"name\":\"getOrder\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"orderId\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"nftAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expiredAt\",\"type\":\"uint256\"}],\"internalType\":\"structIOrder.Order\",\"name\":\"order\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"nftAsset\",\"type\":\"bytes32\"}],\"name\":\"deleteOrder\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// MarketplaceStorageABI is the input ABI used to generate the binding from.
// Deprecated: Use MarketplaceStorageMetaData.ABI instead.
var MarketplaceStorageABI = MarketplaceStorageMetaData.ABI

// MarketplaceStorage is an auto generated Go binding around an Ethereum contract.
type MarketplaceStorage struct {
	MarketplaceStorageCaller     // Read-only binding to the contract
	MarketplaceStorageTransactor // Write-only binding to the contract
	MarketplaceStorageFilterer   // Log filterer for contract events
}

// MarketplaceStorageCaller is an auto generated read-only Go binding around an Ethereum contract.
type MarketplaceStorageCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MarketplaceStorageTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MarketplaceStorageTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MarketplaceStorageFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MarketplaceStorageFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MarketplaceStorageSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MarketplaceStorageSession struct {
	Contract     *MarketplaceStorage // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// MarketplaceStorageCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MarketplaceStorageCallerSession struct {
	Contract *MarketplaceStorageCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// MarketplaceStorageTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MarketplaceStorageTransactorSession struct {
	Contract     *MarketplaceStorageTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// MarketplaceStorageRaw is an auto generated low-level Go binding around an Ethereum contract.
type MarketplaceStorageRaw struct {
	Contract *MarketplaceStorage // Generic contract binding to access the raw methods on
}

// MarketplaceStorageCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MarketplaceStorageCallerRaw struct {
	Contract *MarketplaceStorageCaller // Generic read-only contract binding to access the raw methods on
}

// MarketplaceStorageTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MarketplaceStorageTransactorRaw struct {
	Contract *MarketplaceStorageTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMarketplaceStorage creates a new instance of MarketplaceStorage, bound to a specific deployed contract.
func NewMarketplaceStorage(address common.Address, backend bind.ContractBackend) (*MarketplaceStorage, error) {
	contract, err := bindMarketplaceStorage(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MarketplaceStorage{MarketplaceStorageCaller: MarketplaceStorageCaller{contract: contract}, MarketplaceStorageTransactor: MarketplaceStorageTransactor{contract: contract}, MarketplaceStorageFilterer: MarketplaceStorageFilterer{contract: contract}}, nil
}

// NewMarketplaceStorageCaller creates a new read-only instance of MarketplaceStorage, bound to a specific deployed contract.
func NewMarketplaceStorageCaller(address common.Address, caller bind.ContractCaller) (*MarketplaceStorageCaller, error) {
	contract, err := bindMarketplaceStorage(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MarketplaceStorageCaller{contract: contract}, nil
}

// NewMarketplaceStorageTransactor creates a new write-only instance of MarketplaceStorage, bound to a specific deployed contract.
func NewMarketplaceStorageTransactor(address common.Address, transactor bind.ContractTransactor) (*MarketplaceStorageTransactor, error) {
	contract, err := bindMarketplaceStorage(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MarketplaceStorageTransactor{contract: contract}, nil
}

// NewMarketplaceStorageFilterer creates a new log filterer instance of MarketplaceStorage, bound to a specific deployed contract.
func NewMarketplaceStorageFilterer(address common.Address, filterer bind.ContractFilterer) (*MarketplaceStorageFilterer, error) {
	contract, err := bindMarketplaceStorage(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MarketplaceStorageFilterer{contract: contract}, nil
}

// bindMarketplaceStorage binds a generic wrapper to an already deployed contract.
func bindMarketplaceStorage(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MarketplaceStorageABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MarketplaceStorage *MarketplaceStorageRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MarketplaceStorage.Contract.MarketplaceStorageCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MarketplaceStorage *MarketplaceStorageRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MarketplaceStorage.Contract.MarketplaceStorageTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MarketplaceStorage *MarketplaceStorageRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MarketplaceStorage.Contract.MarketplaceStorageTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MarketplaceStorage *MarketplaceStorageCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MarketplaceStorage.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MarketplaceStorage *MarketplaceStorageTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MarketplaceStorage.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MarketplaceStorage *MarketplaceStorageTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MarketplaceStorage.Contract.contract.Transact(opts, method, params...)
}

// ADMINROLE is a free data retrieval call binding the contract method 0x75b238fc.
//
// Solidity: function ADMIN_ROLE() view returns(bytes32)
func (_MarketplaceStorage *MarketplaceStorageCaller) ADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _MarketplaceStorage.contract.Call(opts, &out, "ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ADMINROLE is a free data retrieval call binding the contract method 0x75b238fc.
//
// Solidity: function ADMIN_ROLE() view returns(bytes32)
func (_MarketplaceStorage *MarketplaceStorageSession) ADMINROLE() ([32]byte, error) {
	return _MarketplaceStorage.Contract.ADMINROLE(&_MarketplaceStorage.CallOpts)
}

// ADMINROLE is a free data retrieval call binding the contract method 0x75b238fc.
//
// Solidity: function ADMIN_ROLE() view returns(bytes32)
func (_MarketplaceStorage *MarketplaceStorageCallerSession) ADMINROLE() ([32]byte, error) {
	return _MarketplaceStorage.Contract.ADMINROLE(&_MarketplaceStorage.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_MarketplaceStorage *MarketplaceStorageCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _MarketplaceStorage.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_MarketplaceStorage *MarketplaceStorageSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _MarketplaceStorage.Contract.DEFAULTADMINROLE(&_MarketplaceStorage.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_MarketplaceStorage *MarketplaceStorageCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _MarketplaceStorage.Contract.DEFAULTADMINROLE(&_MarketplaceStorage.CallOpts)
}

// PAUSERROLE is a free data retrieval call binding the contract method 0xe63ab1e9.
//
// Solidity: function PAUSER_ROLE() view returns(bytes32)
func (_MarketplaceStorage *MarketplaceStorageCaller) PAUSERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _MarketplaceStorage.contract.Call(opts, &out, "PAUSER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// PAUSERROLE is a free data retrieval call binding the contract method 0xe63ab1e9.
//
// Solidity: function PAUSER_ROLE() view returns(bytes32)
func (_MarketplaceStorage *MarketplaceStorageSession) PAUSERROLE() ([32]byte, error) {
	return _MarketplaceStorage.Contract.PAUSERROLE(&_MarketplaceStorage.CallOpts)
}

// PAUSERROLE is a free data retrieval call binding the contract method 0xe63ab1e9.
//
// Solidity: function PAUSER_ROLE() view returns(bytes32)
func (_MarketplaceStorage *MarketplaceStorageCallerSession) PAUSERROLE() ([32]byte, error) {
	return _MarketplaceStorage.Contract.PAUSERROLE(&_MarketplaceStorage.CallOpts)
}

// AssetIsAvailable is a free data retrieval call binding the contract method 0x0a7f8977.
//
// Solidity: function assetIsAvailable(bytes32 nftAsset) view returns(bool)
func (_MarketplaceStorage *MarketplaceStorageCaller) AssetIsAvailable(opts *bind.CallOpts, nftAsset [32]byte) (bool, error) {
	var out []interface{}
	err := _MarketplaceStorage.contract.Call(opts, &out, "assetIsAvailable", nftAsset)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// AssetIsAvailable is a free data retrieval call binding the contract method 0x0a7f8977.
//
// Solidity: function assetIsAvailable(bytes32 nftAsset) view returns(bool)
func (_MarketplaceStorage *MarketplaceStorageSession) AssetIsAvailable(nftAsset [32]byte) (bool, error) {
	return _MarketplaceStorage.Contract.AssetIsAvailable(&_MarketplaceStorage.CallOpts, nftAsset)
}

// AssetIsAvailable is a free data retrieval call binding the contract method 0x0a7f8977.
//
// Solidity: function assetIsAvailable(bytes32 nftAsset) view returns(bool)
func (_MarketplaceStorage *MarketplaceStorageCallerSession) AssetIsAvailable(nftAsset [32]byte) (bool, error) {
	return _MarketplaceStorage.Contract.AssetIsAvailable(&_MarketplaceStorage.CallOpts, nftAsset)
}

// AuctionIsEnded is a free data retrieval call binding the contract method 0xc064bc50.
//
// Solidity: function auctionIsEnded(bytes32 nftAsset, bytes32 auctionId) view returns(bool)
func (_MarketplaceStorage *MarketplaceStorageCaller) AuctionIsEnded(opts *bind.CallOpts, nftAsset [32]byte, auctionId [32]byte) (bool, error) {
	var out []interface{}
	err := _MarketplaceStorage.contract.Call(opts, &out, "auctionIsEnded", nftAsset, auctionId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// AuctionIsEnded is a free data retrieval call binding the contract method 0xc064bc50.
//
// Solidity: function auctionIsEnded(bytes32 nftAsset, bytes32 auctionId) view returns(bool)
func (_MarketplaceStorage *MarketplaceStorageSession) AuctionIsEnded(nftAsset [32]byte, auctionId [32]byte) (bool, error) {
	return _MarketplaceStorage.Contract.AuctionIsEnded(&_MarketplaceStorage.CallOpts, nftAsset, auctionId)
}

// AuctionIsEnded is a free data retrieval call binding the contract method 0xc064bc50.
//
// Solidity: function auctionIsEnded(bytes32 nftAsset, bytes32 auctionId) view returns(bool)
func (_MarketplaceStorage *MarketplaceStorageCallerSession) AuctionIsEnded(nftAsset [32]byte, auctionId [32]byte) (bool, error) {
	return _MarketplaceStorage.Contract.AuctionIsEnded(&_MarketplaceStorage.CallOpts, nftAsset, auctionId)
}

// AuctionIsExisted is a free data retrieval call binding the contract method 0x1a8b4169.
//
// Solidity: function auctionIsExisted(bytes32 auctionId) view returns(bool)
func (_MarketplaceStorage *MarketplaceStorageCaller) AuctionIsExisted(opts *bind.CallOpts, auctionId [32]byte) (bool, error) {
	var out []interface{}
	err := _MarketplaceStorage.contract.Call(opts, &out, "auctionIsExisted", auctionId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// AuctionIsExisted is a free data retrieval call binding the contract method 0x1a8b4169.
//
// Solidity: function auctionIsExisted(bytes32 auctionId) view returns(bool)
func (_MarketplaceStorage *MarketplaceStorageSession) AuctionIsExisted(auctionId [32]byte) (bool, error) {
	return _MarketplaceStorage.Contract.AuctionIsExisted(&_MarketplaceStorage.CallOpts, auctionId)
}

// AuctionIsExisted is a free data retrieval call binding the contract method 0x1a8b4169.
//
// Solidity: function auctionIsExisted(bytes32 auctionId) view returns(bool)
func (_MarketplaceStorage *MarketplaceStorageCallerSession) AuctionIsExisted(auctionId [32]byte) (bool, error) {
	return _MarketplaceStorage.Contract.AuctionIsExisted(&_MarketplaceStorage.CallOpts, auctionId)
}

// AuctionIsRunning is a free data retrieval call binding the contract method 0xbc66c29f.
//
// Solidity: function auctionIsRunning(bytes32 nftAsset, bytes32 auctionId) view returns(bool)
func (_MarketplaceStorage *MarketplaceStorageCaller) AuctionIsRunning(opts *bind.CallOpts, nftAsset [32]byte, auctionId [32]byte) (bool, error) {
	var out []interface{}
	err := _MarketplaceStorage.contract.Call(opts, &out, "auctionIsRunning", nftAsset, auctionId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// AuctionIsRunning is a free data retrieval call binding the contract method 0xbc66c29f.
//
// Solidity: function auctionIsRunning(bytes32 nftAsset, bytes32 auctionId) view returns(bool)
func (_MarketplaceStorage *MarketplaceStorageSession) AuctionIsRunning(nftAsset [32]byte, auctionId [32]byte) (bool, error) {
	return _MarketplaceStorage.Contract.AuctionIsRunning(&_MarketplaceStorage.CallOpts, nftAsset, auctionId)
}

// AuctionIsRunning is a free data retrieval call binding the contract method 0xbc66c29f.
//
// Solidity: function auctionIsRunning(bytes32 nftAsset, bytes32 auctionId) view returns(bool)
func (_MarketplaceStorage *MarketplaceStorageCallerSession) AuctionIsRunning(nftAsset [32]byte, auctionId [32]byte) (bool, error) {
	return _MarketplaceStorage.Contract.AuctionIsRunning(&_MarketplaceStorage.CallOpts, nftAsset, auctionId)
}

// GetAuction is a free data retrieval call binding the contract method 0x15924b5b.
//
// Solidity: function getAuction(bytes32 auctionId) view returns((address,address,bytes32,uint256,uint256,uint256,uint256) order)
func (_MarketplaceStorage *MarketplaceStorageCaller) GetAuction(opts *bind.CallOpts, auctionId [32]byte) (IAuctionAuction, error) {
	var out []interface{}
	err := _MarketplaceStorage.contract.Call(opts, &out, "getAuction", auctionId)

	if err != nil {
		return *new(IAuctionAuction), err
	}

	out0 := *abi.ConvertType(out[0], new(IAuctionAuction)).(*IAuctionAuction)

	return out0, err

}

// GetAuction is a free data retrieval call binding the contract method 0x15924b5b.
//
// Solidity: function getAuction(bytes32 auctionId) view returns((address,address,bytes32,uint256,uint256,uint256,uint256) order)
func (_MarketplaceStorage *MarketplaceStorageSession) GetAuction(auctionId [32]byte) (IAuctionAuction, error) {
	return _MarketplaceStorage.Contract.GetAuction(&_MarketplaceStorage.CallOpts, auctionId)
}

// GetAuction is a free data retrieval call binding the contract method 0x15924b5b.
//
// Solidity: function getAuction(bytes32 auctionId) view returns((address,address,bytes32,uint256,uint256,uint256,uint256) order)
func (_MarketplaceStorage *MarketplaceStorageCallerSession) GetAuction(auctionId [32]byte) (IAuctionAuction, error) {
	return _MarketplaceStorage.Contract.GetAuction(&_MarketplaceStorage.CallOpts, auctionId)
}

// GetOrder is a free data retrieval call binding the contract method 0x5778472a.
//
// Solidity: function getOrder(bytes32 nftAsset) view returns((bytes32,address,address,uint256,uint256) order)
func (_MarketplaceStorage *MarketplaceStorageCaller) GetOrder(opts *bind.CallOpts, nftAsset [32]byte) (IOrderOrder, error) {
	var out []interface{}
	err := _MarketplaceStorage.contract.Call(opts, &out, "getOrder", nftAsset)

	if err != nil {
		return *new(IOrderOrder), err
	}

	out0 := *abi.ConvertType(out[0], new(IOrderOrder)).(*IOrderOrder)

	return out0, err

}

// GetOrder is a free data retrieval call binding the contract method 0x5778472a.
//
// Solidity: function getOrder(bytes32 nftAsset) view returns((bytes32,address,address,uint256,uint256) order)
func (_MarketplaceStorage *MarketplaceStorageSession) GetOrder(nftAsset [32]byte) (IOrderOrder, error) {
	return _MarketplaceStorage.Contract.GetOrder(&_MarketplaceStorage.CallOpts, nftAsset)
}

// GetOrder is a free data retrieval call binding the contract method 0x5778472a.
//
// Solidity: function getOrder(bytes32 nftAsset) view returns((bytes32,address,address,uint256,uint256) order)
func (_MarketplaceStorage *MarketplaceStorageCallerSession) GetOrder(nftAsset [32]byte) (IOrderOrder, error) {
	return _MarketplaceStorage.Contract.GetOrder(&_MarketplaceStorage.CallOpts, nftAsset)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_MarketplaceStorage *MarketplaceStorageCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _MarketplaceStorage.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_MarketplaceStorage *MarketplaceStorageSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _MarketplaceStorage.Contract.GetRoleAdmin(&_MarketplaceStorage.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_MarketplaceStorage *MarketplaceStorageCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _MarketplaceStorage.Contract.GetRoleAdmin(&_MarketplaceStorage.CallOpts, role)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_MarketplaceStorage *MarketplaceStorageCaller) GetRoleMember(opts *bind.CallOpts, role [32]byte, index *big.Int) (common.Address, error) {
	var out []interface{}
	err := _MarketplaceStorage.contract.Call(opts, &out, "getRoleMember", role, index)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_MarketplaceStorage *MarketplaceStorageSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _MarketplaceStorage.Contract.GetRoleMember(&_MarketplaceStorage.CallOpts, role, index)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_MarketplaceStorage *MarketplaceStorageCallerSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _MarketplaceStorage.Contract.GetRoleMember(&_MarketplaceStorage.CallOpts, role, index)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_MarketplaceStorage *MarketplaceStorageCaller) GetRoleMemberCount(opts *bind.CallOpts, role [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _MarketplaceStorage.contract.Call(opts, &out, "getRoleMemberCount", role)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_MarketplaceStorage *MarketplaceStorageSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _MarketplaceStorage.Contract.GetRoleMemberCount(&_MarketplaceStorage.CallOpts, role)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_MarketplaceStorage *MarketplaceStorageCallerSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _MarketplaceStorage.Contract.GetRoleMemberCount(&_MarketplaceStorage.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_MarketplaceStorage *MarketplaceStorageCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _MarketplaceStorage.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_MarketplaceStorage *MarketplaceStorageSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _MarketplaceStorage.Contract.HasRole(&_MarketplaceStorage.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_MarketplaceStorage *MarketplaceStorageCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _MarketplaceStorage.Contract.HasRole(&_MarketplaceStorage.CallOpts, role, account)
}

// OrderByNftAsset is a free data retrieval call binding the contract method 0x3a3999d9.
//
// Solidity: function orderByNftAsset(bytes32 ) view returns(bytes32 orderId, address seller, address nftAddress, uint256 price, uint256 expiredAt)
func (_MarketplaceStorage *MarketplaceStorageCaller) OrderByNftAsset(opts *bind.CallOpts, arg0 [32]byte) (struct {
	OrderId    [32]byte
	Seller     common.Address
	NftAddress common.Address
	Price      *big.Int
	ExpiredAt  *big.Int
}, error) {
	var out []interface{}
	err := _MarketplaceStorage.contract.Call(opts, &out, "orderByNftAsset", arg0)

	outstruct := new(struct {
		OrderId    [32]byte
		Seller     common.Address
		NftAddress common.Address
		Price      *big.Int
		ExpiredAt  *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.OrderId = *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	outstruct.Seller = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.NftAddress = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)
	outstruct.Price = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.ExpiredAt = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// OrderByNftAsset is a free data retrieval call binding the contract method 0x3a3999d9.
//
// Solidity: function orderByNftAsset(bytes32 ) view returns(bytes32 orderId, address seller, address nftAddress, uint256 price, uint256 expiredAt)
func (_MarketplaceStorage *MarketplaceStorageSession) OrderByNftAsset(arg0 [32]byte) (struct {
	OrderId    [32]byte
	Seller     common.Address
	NftAddress common.Address
	Price      *big.Int
	ExpiredAt  *big.Int
}, error) {
	return _MarketplaceStorage.Contract.OrderByNftAsset(&_MarketplaceStorage.CallOpts, arg0)
}

// OrderByNftAsset is a free data retrieval call binding the contract method 0x3a3999d9.
//
// Solidity: function orderByNftAsset(bytes32 ) view returns(bytes32 orderId, address seller, address nftAddress, uint256 price, uint256 expiredAt)
func (_MarketplaceStorage *MarketplaceStorageCallerSession) OrderByNftAsset(arg0 [32]byte) (struct {
	OrderId    [32]byte
	Seller     common.Address
	NftAddress common.Address
	Price      *big.Int
	ExpiredAt  *big.Int
}, error) {
	return _MarketplaceStorage.Contract.OrderByNftAsset(&_MarketplaceStorage.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_MarketplaceStorage *MarketplaceStorageCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MarketplaceStorage.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_MarketplaceStorage *MarketplaceStorageSession) Owner() (common.Address, error) {
	return _MarketplaceStorage.Contract.Owner(&_MarketplaceStorage.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_MarketplaceStorage *MarketplaceStorageCallerSession) Owner() (common.Address, error) {
	return _MarketplaceStorage.Contract.Owner(&_MarketplaceStorage.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_MarketplaceStorage *MarketplaceStorageCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _MarketplaceStorage.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_MarketplaceStorage *MarketplaceStorageSession) Paused() (bool, error) {
	return _MarketplaceStorage.Contract.Paused(&_MarketplaceStorage.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_MarketplaceStorage *MarketplaceStorageCallerSession) Paused() (bool, error) {
	return _MarketplaceStorage.Contract.Paused(&_MarketplaceStorage.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_MarketplaceStorage *MarketplaceStorageCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _MarketplaceStorage.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_MarketplaceStorage *MarketplaceStorageSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _MarketplaceStorage.Contract.SupportsInterface(&_MarketplaceStorage.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_MarketplaceStorage *MarketplaceStorageCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _MarketplaceStorage.Contract.SupportsInterface(&_MarketplaceStorage.CallOpts, interfaceId)
}

// AuctionEnded is a paid mutator transaction binding the contract method 0xac84b7bc.
//
// Solidity: function auctionEnded(bytes32 nftAsset) returns()
func (_MarketplaceStorage *MarketplaceStorageTransactor) AuctionEnded(opts *bind.TransactOpts, nftAsset [32]byte) (*types.Transaction, error) {
	return _MarketplaceStorage.contract.Transact(opts, "auctionEnded", nftAsset)
}

// AuctionEnded is a paid mutator transaction binding the contract method 0xac84b7bc.
//
// Solidity: function auctionEnded(bytes32 nftAsset) returns()
func (_MarketplaceStorage *MarketplaceStorageSession) AuctionEnded(nftAsset [32]byte) (*types.Transaction, error) {
	return _MarketplaceStorage.Contract.AuctionEnded(&_MarketplaceStorage.TransactOpts, nftAsset)
}

// AuctionEnded is a paid mutator transaction binding the contract method 0xac84b7bc.
//
// Solidity: function auctionEnded(bytes32 nftAsset) returns()
func (_MarketplaceStorage *MarketplaceStorageTransactorSession) AuctionEnded(nftAsset [32]byte) (*types.Transaction, error) {
	return _MarketplaceStorage.Contract.AuctionEnded(&_MarketplaceStorage.TransactOpts, nftAsset)
}

// CreateAuction is a paid mutator transaction binding the contract method 0x84b04aa9.
//
// Solidity: function createAuction(address assetOwner, address nftAddress, uint256 assetId, bytes32 auctionId, uint256 biddingEnd, uint256 revealEnd, uint256 startPriceInWei) returns()
func (_MarketplaceStorage *MarketplaceStorageTransactor) CreateAuction(opts *bind.TransactOpts, assetOwner common.Address, nftAddress common.Address, assetId *big.Int, auctionId [32]byte, biddingEnd *big.Int, revealEnd *big.Int, startPriceInWei *big.Int) (*types.Transaction, error) {
	return _MarketplaceStorage.contract.Transact(opts, "createAuction", assetOwner, nftAddress, assetId, auctionId, biddingEnd, revealEnd, startPriceInWei)
}

// CreateAuction is a paid mutator transaction binding the contract method 0x84b04aa9.
//
// Solidity: function createAuction(address assetOwner, address nftAddress, uint256 assetId, bytes32 auctionId, uint256 biddingEnd, uint256 revealEnd, uint256 startPriceInWei) returns()
func (_MarketplaceStorage *MarketplaceStorageSession) CreateAuction(assetOwner common.Address, nftAddress common.Address, assetId *big.Int, auctionId [32]byte, biddingEnd *big.Int, revealEnd *big.Int, startPriceInWei *big.Int) (*types.Transaction, error) {
	return _MarketplaceStorage.Contract.CreateAuction(&_MarketplaceStorage.TransactOpts, assetOwner, nftAddress, assetId, auctionId, biddingEnd, revealEnd, startPriceInWei)
}

// CreateAuction is a paid mutator transaction binding the contract method 0x84b04aa9.
//
// Solidity: function createAuction(address assetOwner, address nftAddress, uint256 assetId, bytes32 auctionId, uint256 biddingEnd, uint256 revealEnd, uint256 startPriceInWei) returns()
func (_MarketplaceStorage *MarketplaceStorageTransactorSession) CreateAuction(assetOwner common.Address, nftAddress common.Address, assetId *big.Int, auctionId [32]byte, biddingEnd *big.Int, revealEnd *big.Int, startPriceInWei *big.Int) (*types.Transaction, error) {
	return _MarketplaceStorage.Contract.CreateAuction(&_MarketplaceStorage.TransactOpts, assetOwner, nftAddress, assetId, auctionId, biddingEnd, revealEnd, startPriceInWei)
}

// CreateOrder is a paid mutator transaction binding the contract method 0x930ca094.
//
// Solidity: function createOrder(address seller, address nftAddress, uint256 assetId, bytes32 orderId, uint256 priceInWei, uint256 expiredAt) returns()
func (_MarketplaceStorage *MarketplaceStorageTransactor) CreateOrder(opts *bind.TransactOpts, seller common.Address, nftAddress common.Address, assetId *big.Int, orderId [32]byte, priceInWei *big.Int, expiredAt *big.Int) (*types.Transaction, error) {
	return _MarketplaceStorage.contract.Transact(opts, "createOrder", seller, nftAddress, assetId, orderId, priceInWei, expiredAt)
}

// CreateOrder is a paid mutator transaction binding the contract method 0x930ca094.
//
// Solidity: function createOrder(address seller, address nftAddress, uint256 assetId, bytes32 orderId, uint256 priceInWei, uint256 expiredAt) returns()
func (_MarketplaceStorage *MarketplaceStorageSession) CreateOrder(seller common.Address, nftAddress common.Address, assetId *big.Int, orderId [32]byte, priceInWei *big.Int, expiredAt *big.Int) (*types.Transaction, error) {
	return _MarketplaceStorage.Contract.CreateOrder(&_MarketplaceStorage.TransactOpts, seller, nftAddress, assetId, orderId, priceInWei, expiredAt)
}

// CreateOrder is a paid mutator transaction binding the contract method 0x930ca094.
//
// Solidity: function createOrder(address seller, address nftAddress, uint256 assetId, bytes32 orderId, uint256 priceInWei, uint256 expiredAt) returns()
func (_MarketplaceStorage *MarketplaceStorageTransactorSession) CreateOrder(seller common.Address, nftAddress common.Address, assetId *big.Int, orderId [32]byte, priceInWei *big.Int, expiredAt *big.Int) (*types.Transaction, error) {
	return _MarketplaceStorage.Contract.CreateOrder(&_MarketplaceStorage.TransactOpts, seller, nftAddress, assetId, orderId, priceInWei, expiredAt)
}

// DeleteOrder is a paid mutator transaction binding the contract method 0x87a61cbd.
//
// Solidity: function deleteOrder(bytes32 nftAsset) returns()
func (_MarketplaceStorage *MarketplaceStorageTransactor) DeleteOrder(opts *bind.TransactOpts, nftAsset [32]byte) (*types.Transaction, error) {
	return _MarketplaceStorage.contract.Transact(opts, "deleteOrder", nftAsset)
}

// DeleteOrder is a paid mutator transaction binding the contract method 0x87a61cbd.
//
// Solidity: function deleteOrder(bytes32 nftAsset) returns()
func (_MarketplaceStorage *MarketplaceStorageSession) DeleteOrder(nftAsset [32]byte) (*types.Transaction, error) {
	return _MarketplaceStorage.Contract.DeleteOrder(&_MarketplaceStorage.TransactOpts, nftAsset)
}

// DeleteOrder is a paid mutator transaction binding the contract method 0x87a61cbd.
//
// Solidity: function deleteOrder(bytes32 nftAsset) returns()
func (_MarketplaceStorage *MarketplaceStorageTransactorSession) DeleteOrder(nftAsset [32]byte) (*types.Transaction, error) {
	return _MarketplaceStorage.Contract.DeleteOrder(&_MarketplaceStorage.TransactOpts, nftAsset)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_MarketplaceStorage *MarketplaceStorageTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _MarketplaceStorage.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_MarketplaceStorage *MarketplaceStorageSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _MarketplaceStorage.Contract.GrantRole(&_MarketplaceStorage.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_MarketplaceStorage *MarketplaceStorageTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _MarketplaceStorage.Contract.GrantRole(&_MarketplaceStorage.TransactOpts, role, account)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_MarketplaceStorage *MarketplaceStorageTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MarketplaceStorage.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_MarketplaceStorage *MarketplaceStorageSession) Pause() (*types.Transaction, error) {
	return _MarketplaceStorage.Contract.Pause(&_MarketplaceStorage.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_MarketplaceStorage *MarketplaceStorageTransactorSession) Pause() (*types.Transaction, error) {
	return _MarketplaceStorage.Contract.Pause(&_MarketplaceStorage.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_MarketplaceStorage *MarketplaceStorageTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MarketplaceStorage.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_MarketplaceStorage *MarketplaceStorageSession) RenounceOwnership() (*types.Transaction, error) {
	return _MarketplaceStorage.Contract.RenounceOwnership(&_MarketplaceStorage.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_MarketplaceStorage *MarketplaceStorageTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _MarketplaceStorage.Contract.RenounceOwnership(&_MarketplaceStorage.TransactOpts)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_MarketplaceStorage *MarketplaceStorageTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _MarketplaceStorage.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_MarketplaceStorage *MarketplaceStorageSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _MarketplaceStorage.Contract.RenounceRole(&_MarketplaceStorage.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_MarketplaceStorage *MarketplaceStorageTransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _MarketplaceStorage.Contract.RenounceRole(&_MarketplaceStorage.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_MarketplaceStorage *MarketplaceStorageTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _MarketplaceStorage.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_MarketplaceStorage *MarketplaceStorageSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _MarketplaceStorage.Contract.RevokeRole(&_MarketplaceStorage.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_MarketplaceStorage *MarketplaceStorageTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _MarketplaceStorage.Contract.RevokeRole(&_MarketplaceStorage.TransactOpts, role, account)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_MarketplaceStorage *MarketplaceStorageTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _MarketplaceStorage.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_MarketplaceStorage *MarketplaceStorageSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _MarketplaceStorage.Contract.TransferOwnership(&_MarketplaceStorage.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_MarketplaceStorage *MarketplaceStorageTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _MarketplaceStorage.Contract.TransferOwnership(&_MarketplaceStorage.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_MarketplaceStorage *MarketplaceStorageTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MarketplaceStorage.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_MarketplaceStorage *MarketplaceStorageSession) Unpause() (*types.Transaction, error) {
	return _MarketplaceStorage.Contract.Unpause(&_MarketplaceStorage.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_MarketplaceStorage *MarketplaceStorageTransactorSession) Unpause() (*types.Transaction, error) {
	return _MarketplaceStorage.Contract.Unpause(&_MarketplaceStorage.TransactOpts)
}

// UpdateAuctionMarketplace is a paid mutator transaction binding the contract method 0x3dd328d1.
//
// Solidity: function updateAuctionMarketplace(address _auctionMarketplace) returns()
func (_MarketplaceStorage *MarketplaceStorageTransactor) UpdateAuctionMarketplace(opts *bind.TransactOpts, _auctionMarketplace common.Address) (*types.Transaction, error) {
	return _MarketplaceStorage.contract.Transact(opts, "updateAuctionMarketplace", _auctionMarketplace)
}

// UpdateAuctionMarketplace is a paid mutator transaction binding the contract method 0x3dd328d1.
//
// Solidity: function updateAuctionMarketplace(address _auctionMarketplace) returns()
func (_MarketplaceStorage *MarketplaceStorageSession) UpdateAuctionMarketplace(_auctionMarketplace common.Address) (*types.Transaction, error) {
	return _MarketplaceStorage.Contract.UpdateAuctionMarketplace(&_MarketplaceStorage.TransactOpts, _auctionMarketplace)
}

// UpdateAuctionMarketplace is a paid mutator transaction binding the contract method 0x3dd328d1.
//
// Solidity: function updateAuctionMarketplace(address _auctionMarketplace) returns()
func (_MarketplaceStorage *MarketplaceStorageTransactorSession) UpdateAuctionMarketplace(_auctionMarketplace common.Address) (*types.Transaction, error) {
	return _MarketplaceStorage.Contract.UpdateAuctionMarketplace(&_MarketplaceStorage.TransactOpts, _auctionMarketplace)
}

// UpdateHighestBid is a paid mutator transaction binding the contract method 0x572f7d5e.
//
// Solidity: function updateHighestBid(address bidder, uint256 highestBid, bytes32 auctionId) returns()
func (_MarketplaceStorage *MarketplaceStorageTransactor) UpdateHighestBid(opts *bind.TransactOpts, bidder common.Address, highestBid *big.Int, auctionId [32]byte) (*types.Transaction, error) {
	return _MarketplaceStorage.contract.Transact(opts, "updateHighestBid", bidder, highestBid, auctionId)
}

// UpdateHighestBid is a paid mutator transaction binding the contract method 0x572f7d5e.
//
// Solidity: function updateHighestBid(address bidder, uint256 highestBid, bytes32 auctionId) returns()
func (_MarketplaceStorage *MarketplaceStorageSession) UpdateHighestBid(bidder common.Address, highestBid *big.Int, auctionId [32]byte) (*types.Transaction, error) {
	return _MarketplaceStorage.Contract.UpdateHighestBid(&_MarketplaceStorage.TransactOpts, bidder, highestBid, auctionId)
}

// UpdateHighestBid is a paid mutator transaction binding the contract method 0x572f7d5e.
//
// Solidity: function updateHighestBid(address bidder, uint256 highestBid, bytes32 auctionId) returns()
func (_MarketplaceStorage *MarketplaceStorageTransactorSession) UpdateHighestBid(bidder common.Address, highestBid *big.Int, auctionId [32]byte) (*types.Transaction, error) {
	return _MarketplaceStorage.Contract.UpdateHighestBid(&_MarketplaceStorage.TransactOpts, bidder, highestBid, auctionId)
}

// UpdateOrderMarketplace is a paid mutator transaction binding the contract method 0x9a91d9c0.
//
// Solidity: function updateOrderMarketplace(address _orderMarketplace) returns()
func (_MarketplaceStorage *MarketplaceStorageTransactor) UpdateOrderMarketplace(opts *bind.TransactOpts, _orderMarketplace common.Address) (*types.Transaction, error) {
	return _MarketplaceStorage.contract.Transact(opts, "updateOrderMarketplace", _orderMarketplace)
}

// UpdateOrderMarketplace is a paid mutator transaction binding the contract method 0x9a91d9c0.
//
// Solidity: function updateOrderMarketplace(address _orderMarketplace) returns()
func (_MarketplaceStorage *MarketplaceStorageSession) UpdateOrderMarketplace(_orderMarketplace common.Address) (*types.Transaction, error) {
	return _MarketplaceStorage.Contract.UpdateOrderMarketplace(&_MarketplaceStorage.TransactOpts, _orderMarketplace)
}

// UpdateOrderMarketplace is a paid mutator transaction binding the contract method 0x9a91d9c0.
//
// Solidity: function updateOrderMarketplace(address _orderMarketplace) returns()
func (_MarketplaceStorage *MarketplaceStorageTransactorSession) UpdateOrderMarketplace(_orderMarketplace common.Address) (*types.Transaction, error) {
	return _MarketplaceStorage.Contract.UpdateOrderMarketplace(&_MarketplaceStorage.TransactOpts, _orderMarketplace)
}

// MarketplaceStorageAuctionCancelledIterator is returned from FilterAuctionCancelled and is used to iterate over the raw logs and unpacked data for AuctionCancelled events raised by the MarketplaceStorage contract.
type MarketplaceStorageAuctionCancelledIterator struct {
	Event *MarketplaceStorageAuctionCancelled // Event containing the contract specifics and raw log

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
func (it *MarketplaceStorageAuctionCancelledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketplaceStorageAuctionCancelled)
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
		it.Event = new(MarketplaceStorageAuctionCancelled)
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
func (it *MarketplaceStorageAuctionCancelledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketplaceStorageAuctionCancelledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketplaceStorageAuctionCancelled represents a AuctionCancelled event raised by the MarketplaceStorage contract.
type MarketplaceStorageAuctionCancelled struct {
	Who       common.Address
	AuctionId [32]byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterAuctionCancelled is a free log retrieval operation binding the contract event 0x4283fca6e4b7cf10576815155991f12d5ef6adb5bcacdf255d3da62fddbf84ea.
//
// Solidity: event AuctionCancelled(address who, bytes32 auctionId)
func (_MarketplaceStorage *MarketplaceStorageFilterer) FilterAuctionCancelled(opts *bind.FilterOpts) (*MarketplaceStorageAuctionCancelledIterator, error) {

	logs, sub, err := _MarketplaceStorage.contract.FilterLogs(opts, "AuctionCancelled")
	if err != nil {
		return nil, err
	}
	return &MarketplaceStorageAuctionCancelledIterator{contract: _MarketplaceStorage.contract, event: "AuctionCancelled", logs: logs, sub: sub}, nil
}

// WatchAuctionCancelled is a free log subscription operation binding the contract event 0x4283fca6e4b7cf10576815155991f12d5ef6adb5bcacdf255d3da62fddbf84ea.
//
// Solidity: event AuctionCancelled(address who, bytes32 auctionId)
func (_MarketplaceStorage *MarketplaceStorageFilterer) WatchAuctionCancelled(opts *bind.WatchOpts, sink chan<- *MarketplaceStorageAuctionCancelled) (event.Subscription, error) {

	logs, sub, err := _MarketplaceStorage.contract.WatchLogs(opts, "AuctionCancelled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketplaceStorageAuctionCancelled)
				if err := _MarketplaceStorage.contract.UnpackLog(event, "AuctionCancelled", log); err != nil {
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

// ParseAuctionCancelled is a log parse operation binding the contract event 0x4283fca6e4b7cf10576815155991f12d5ef6adb5bcacdf255d3da62fddbf84ea.
//
// Solidity: event AuctionCancelled(address who, bytes32 auctionId)
func (_MarketplaceStorage *MarketplaceStorageFilterer) ParseAuctionCancelled(log types.Log) (*MarketplaceStorageAuctionCancelled, error) {
	event := new(MarketplaceStorageAuctionCancelled)
	if err := _MarketplaceStorage.contract.UnpackLog(event, "AuctionCancelled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MarketplaceStorageAuctionCreatedIterator is returned from FilterAuctionCreated and is used to iterate over the raw logs and unpacked data for AuctionCreated events raised by the MarketplaceStorage contract.
type MarketplaceStorageAuctionCreatedIterator struct {
	Event *MarketplaceStorageAuctionCreated // Event containing the contract specifics and raw log

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
func (it *MarketplaceStorageAuctionCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketplaceStorageAuctionCreated)
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
		it.Event = new(MarketplaceStorageAuctionCreated)
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
func (it *MarketplaceStorageAuctionCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketplaceStorageAuctionCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketplaceStorageAuctionCreated represents a AuctionCreated event raised by the MarketplaceStorage contract.
type MarketplaceStorageAuctionCreated struct {
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
// Solidity: event AuctionCreated(address seller, address nftAddress, bytes32 auctionId, uint256 assetId, uint256 biddingEnd, uint256 revealEnd, uint256 startPriceInWei)
func (_MarketplaceStorage *MarketplaceStorageFilterer) FilterAuctionCreated(opts *bind.FilterOpts) (*MarketplaceStorageAuctionCreatedIterator, error) {

	logs, sub, err := _MarketplaceStorage.contract.FilterLogs(opts, "AuctionCreated")
	if err != nil {
		return nil, err
	}
	return &MarketplaceStorageAuctionCreatedIterator{contract: _MarketplaceStorage.contract, event: "AuctionCreated", logs: logs, sub: sub}, nil
}

// WatchAuctionCreated is a free log subscription operation binding the contract event 0xaf4bad306f14d5c908e5b871fa54296ed7d4f29b8092bf5062de6813c247e54b.
//
// Solidity: event AuctionCreated(address seller, address nftAddress, bytes32 auctionId, uint256 assetId, uint256 biddingEnd, uint256 revealEnd, uint256 startPriceInWei)
func (_MarketplaceStorage *MarketplaceStorageFilterer) WatchAuctionCreated(opts *bind.WatchOpts, sink chan<- *MarketplaceStorageAuctionCreated) (event.Subscription, error) {

	logs, sub, err := _MarketplaceStorage.contract.WatchLogs(opts, "AuctionCreated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketplaceStorageAuctionCreated)
				if err := _MarketplaceStorage.contract.UnpackLog(event, "AuctionCreated", log); err != nil {
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
// Solidity: event AuctionCreated(address seller, address nftAddress, bytes32 auctionId, uint256 assetId, uint256 biddingEnd, uint256 revealEnd, uint256 startPriceInWei)
func (_MarketplaceStorage *MarketplaceStorageFilterer) ParseAuctionCreated(log types.Log) (*MarketplaceStorageAuctionCreated, error) {
	event := new(MarketplaceStorageAuctionCreated)
	if err := _MarketplaceStorage.contract.UnpackLog(event, "AuctionCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MarketplaceStorageAuctionEndedIterator is returned from FilterAuctionEnded and is used to iterate over the raw logs and unpacked data for AuctionEnded events raised by the MarketplaceStorage contract.
type MarketplaceStorageAuctionEndedIterator struct {
	Event *MarketplaceStorageAuctionEnded // Event containing the contract specifics and raw log

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
func (it *MarketplaceStorageAuctionEndedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketplaceStorageAuctionEnded)
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
		it.Event = new(MarketplaceStorageAuctionEnded)
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
func (it *MarketplaceStorageAuctionEndedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketplaceStorageAuctionEndedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketplaceStorageAuctionEnded represents a AuctionEnded event raised by the MarketplaceStorage contract.
type MarketplaceStorageAuctionEnded struct {
	AuctionId [32]byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterAuctionEnded is a free log retrieval operation binding the contract event 0x447d0298e6e2716a4343195759c35b277a60b76e628284d93262fa9869db9921.
//
// Solidity: event AuctionEnded(bytes32 auctionId)
func (_MarketplaceStorage *MarketplaceStorageFilterer) FilterAuctionEnded(opts *bind.FilterOpts) (*MarketplaceStorageAuctionEndedIterator, error) {

	logs, sub, err := _MarketplaceStorage.contract.FilterLogs(opts, "AuctionEnded")
	if err != nil {
		return nil, err
	}
	return &MarketplaceStorageAuctionEndedIterator{contract: _MarketplaceStorage.contract, event: "AuctionEnded", logs: logs, sub: sub}, nil
}

// WatchAuctionEnded is a free log subscription operation binding the contract event 0x447d0298e6e2716a4343195759c35b277a60b76e628284d93262fa9869db9921.
//
// Solidity: event AuctionEnded(bytes32 auctionId)
func (_MarketplaceStorage *MarketplaceStorageFilterer) WatchAuctionEnded(opts *bind.WatchOpts, sink chan<- *MarketplaceStorageAuctionEnded) (event.Subscription, error) {

	logs, sub, err := _MarketplaceStorage.contract.WatchLogs(opts, "AuctionEnded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketplaceStorageAuctionEnded)
				if err := _MarketplaceStorage.contract.UnpackLog(event, "AuctionEnded", log); err != nil {
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
// Solidity: event AuctionEnded(bytes32 auctionId)
func (_MarketplaceStorage *MarketplaceStorageFilterer) ParseAuctionEnded(log types.Log) (*MarketplaceStorageAuctionEnded, error) {
	event := new(MarketplaceStorageAuctionEnded)
	if err := _MarketplaceStorage.contract.UnpackLog(event, "AuctionEnded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MarketplaceStorageAuctionRefundIterator is returned from FilterAuctionRefund and is used to iterate over the raw logs and unpacked data for AuctionRefund events raised by the MarketplaceStorage contract.
type MarketplaceStorageAuctionRefundIterator struct {
	Event *MarketplaceStorageAuctionRefund // Event containing the contract specifics and raw log

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
func (it *MarketplaceStorageAuctionRefundIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketplaceStorageAuctionRefund)
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
		it.Event = new(MarketplaceStorageAuctionRefund)
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
func (it *MarketplaceStorageAuctionRefundIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketplaceStorageAuctionRefundIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketplaceStorageAuctionRefund represents a AuctionRefund event raised by the MarketplaceStorage contract.
type MarketplaceStorageAuctionRefund struct {
	Bidder    common.Address
	AuctionId [32]byte
	Deposit   *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterAuctionRefund is a free log retrieval operation binding the contract event 0x6be22eb0241a0575f8b866ff176552cf71078977d6d30e501dd2d326c8d57114.
//
// Solidity: event AuctionRefund(address bidder, bytes32 auctionId, uint256 deposit)
func (_MarketplaceStorage *MarketplaceStorageFilterer) FilterAuctionRefund(opts *bind.FilterOpts) (*MarketplaceStorageAuctionRefundIterator, error) {

	logs, sub, err := _MarketplaceStorage.contract.FilterLogs(opts, "AuctionRefund")
	if err != nil {
		return nil, err
	}
	return &MarketplaceStorageAuctionRefundIterator{contract: _MarketplaceStorage.contract, event: "AuctionRefund", logs: logs, sub: sub}, nil
}

// WatchAuctionRefund is a free log subscription operation binding the contract event 0x6be22eb0241a0575f8b866ff176552cf71078977d6d30e501dd2d326c8d57114.
//
// Solidity: event AuctionRefund(address bidder, bytes32 auctionId, uint256 deposit)
func (_MarketplaceStorage *MarketplaceStorageFilterer) WatchAuctionRefund(opts *bind.WatchOpts, sink chan<- *MarketplaceStorageAuctionRefund) (event.Subscription, error) {

	logs, sub, err := _MarketplaceStorage.contract.WatchLogs(opts, "AuctionRefund")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketplaceStorageAuctionRefund)
				if err := _MarketplaceStorage.contract.UnpackLog(event, "AuctionRefund", log); err != nil {
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
// Solidity: event AuctionRefund(address bidder, bytes32 auctionId, uint256 deposit)
func (_MarketplaceStorage *MarketplaceStorageFilterer) ParseAuctionRefund(log types.Log) (*MarketplaceStorageAuctionRefund, error) {
	event := new(MarketplaceStorageAuctionRefund)
	if err := _MarketplaceStorage.contract.UnpackLog(event, "AuctionRefund", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MarketplaceStorageAuctionSuccessfulIterator is returned from FilterAuctionSuccessful and is used to iterate over the raw logs and unpacked data for AuctionSuccessful events raised by the MarketplaceStorage contract.
type MarketplaceStorageAuctionSuccessfulIterator struct {
	Event *MarketplaceStorageAuctionSuccessful // Event containing the contract specifics and raw log

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
func (it *MarketplaceStorageAuctionSuccessfulIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketplaceStorageAuctionSuccessful)
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
		it.Event = new(MarketplaceStorageAuctionSuccessful)
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
func (it *MarketplaceStorageAuctionSuccessfulIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketplaceStorageAuctionSuccessfulIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketplaceStorageAuctionSuccessful represents a AuctionSuccessful event raised by the MarketplaceStorage contract.
type MarketplaceStorageAuctionSuccessful struct {
	Seller     common.Address
	Buyer      common.Address
	AuctionId  [32]byte
	TotalPrice *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterAuctionSuccessful is a free log retrieval operation binding the contract event 0x5ed769e966210f5d06d68e75831871131255cee237cb8edf78565aae170475ea.
//
// Solidity: event AuctionSuccessful(address seller, address buyer, bytes32 auctionId, uint256 totalPrice)
func (_MarketplaceStorage *MarketplaceStorageFilterer) FilterAuctionSuccessful(opts *bind.FilterOpts) (*MarketplaceStorageAuctionSuccessfulIterator, error) {

	logs, sub, err := _MarketplaceStorage.contract.FilterLogs(opts, "AuctionSuccessful")
	if err != nil {
		return nil, err
	}
	return &MarketplaceStorageAuctionSuccessfulIterator{contract: _MarketplaceStorage.contract, event: "AuctionSuccessful", logs: logs, sub: sub}, nil
}

// WatchAuctionSuccessful is a free log subscription operation binding the contract event 0x5ed769e966210f5d06d68e75831871131255cee237cb8edf78565aae170475ea.
//
// Solidity: event AuctionSuccessful(address seller, address buyer, bytes32 auctionId, uint256 totalPrice)
func (_MarketplaceStorage *MarketplaceStorageFilterer) WatchAuctionSuccessful(opts *bind.WatchOpts, sink chan<- *MarketplaceStorageAuctionSuccessful) (event.Subscription, error) {

	logs, sub, err := _MarketplaceStorage.contract.WatchLogs(opts, "AuctionSuccessful")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketplaceStorageAuctionSuccessful)
				if err := _MarketplaceStorage.contract.UnpackLog(event, "AuctionSuccessful", log); err != nil {
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
// Solidity: event AuctionSuccessful(address seller, address buyer, bytes32 auctionId, uint256 totalPrice)
func (_MarketplaceStorage *MarketplaceStorageFilterer) ParseAuctionSuccessful(log types.Log) (*MarketplaceStorageAuctionSuccessful, error) {
	event := new(MarketplaceStorageAuctionSuccessful)
	if err := _MarketplaceStorage.contract.UnpackLog(event, "AuctionSuccessful", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MarketplaceStorageBidSuccessfulIterator is returned from FilterBidSuccessful and is used to iterate over the raw logs and unpacked data for BidSuccessful events raised by the MarketplaceStorage contract.
type MarketplaceStorageBidSuccessfulIterator struct {
	Event *MarketplaceStorageBidSuccessful // Event containing the contract specifics and raw log

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
func (it *MarketplaceStorageBidSuccessfulIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketplaceStorageBidSuccessful)
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
		it.Event = new(MarketplaceStorageBidSuccessful)
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
func (it *MarketplaceStorageBidSuccessfulIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketplaceStorageBidSuccessfulIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketplaceStorageBidSuccessful represents a BidSuccessful event raised by the MarketplaceStorage contract.
type MarketplaceStorageBidSuccessful struct {
	Bidder     common.Address
	AuctionId  [32]byte
	BlindedBid [32]byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterBidSuccessful is a free log retrieval operation binding the contract event 0x327e5a459c964bcfa4f6d71b32a04a53b0002417b0924a8057d9f12e688e0bec.
//
// Solidity: event BidSuccessful(address bidder, bytes32 auctionId, bytes32 blindedBid)
func (_MarketplaceStorage *MarketplaceStorageFilterer) FilterBidSuccessful(opts *bind.FilterOpts) (*MarketplaceStorageBidSuccessfulIterator, error) {

	logs, sub, err := _MarketplaceStorage.contract.FilterLogs(opts, "BidSuccessful")
	if err != nil {
		return nil, err
	}
	return &MarketplaceStorageBidSuccessfulIterator{contract: _MarketplaceStorage.contract, event: "BidSuccessful", logs: logs, sub: sub}, nil
}

// WatchBidSuccessful is a free log subscription operation binding the contract event 0x327e5a459c964bcfa4f6d71b32a04a53b0002417b0924a8057d9f12e688e0bec.
//
// Solidity: event BidSuccessful(address bidder, bytes32 auctionId, bytes32 blindedBid)
func (_MarketplaceStorage *MarketplaceStorageFilterer) WatchBidSuccessful(opts *bind.WatchOpts, sink chan<- *MarketplaceStorageBidSuccessful) (event.Subscription, error) {

	logs, sub, err := _MarketplaceStorage.contract.WatchLogs(opts, "BidSuccessful")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketplaceStorageBidSuccessful)
				if err := _MarketplaceStorage.contract.UnpackLog(event, "BidSuccessful", log); err != nil {
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
// Solidity: event BidSuccessful(address bidder, bytes32 auctionId, bytes32 blindedBid)
func (_MarketplaceStorage *MarketplaceStorageFilterer) ParseBidSuccessful(log types.Log) (*MarketplaceStorageBidSuccessful, error) {
	event := new(MarketplaceStorageBidSuccessful)
	if err := _MarketplaceStorage.contract.UnpackLog(event, "BidSuccessful", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MarketplaceStorageOrderCancelledIterator is returned from FilterOrderCancelled and is used to iterate over the raw logs and unpacked data for OrderCancelled events raised by the MarketplaceStorage contract.
type MarketplaceStorageOrderCancelledIterator struct {
	Event *MarketplaceStorageOrderCancelled // Event containing the contract specifics and raw log

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
func (it *MarketplaceStorageOrderCancelledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketplaceStorageOrderCancelled)
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
		it.Event = new(MarketplaceStorageOrderCancelled)
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
func (it *MarketplaceStorageOrderCancelledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketplaceStorageOrderCancelledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketplaceStorageOrderCancelled represents a OrderCancelled event raised by the MarketplaceStorage contract.
type MarketplaceStorageOrderCancelled struct {
	Who common.Address
	Id  [32]byte
	Raw types.Log // Blockchain specific contextual infos
}

// FilterOrderCancelled is a free log retrieval operation binding the contract event 0x35974c4230d53fb4c6e8553fd900c88ba92747dbc689a79bcd6ba755cb936985.
//
// Solidity: event OrderCancelled(address who, bytes32 id)
func (_MarketplaceStorage *MarketplaceStorageFilterer) FilterOrderCancelled(opts *bind.FilterOpts) (*MarketplaceStorageOrderCancelledIterator, error) {

	logs, sub, err := _MarketplaceStorage.contract.FilterLogs(opts, "OrderCancelled")
	if err != nil {
		return nil, err
	}
	return &MarketplaceStorageOrderCancelledIterator{contract: _MarketplaceStorage.contract, event: "OrderCancelled", logs: logs, sub: sub}, nil
}

// WatchOrderCancelled is a free log subscription operation binding the contract event 0x35974c4230d53fb4c6e8553fd900c88ba92747dbc689a79bcd6ba755cb936985.
//
// Solidity: event OrderCancelled(address who, bytes32 id)
func (_MarketplaceStorage *MarketplaceStorageFilterer) WatchOrderCancelled(opts *bind.WatchOpts, sink chan<- *MarketplaceStorageOrderCancelled) (event.Subscription, error) {

	logs, sub, err := _MarketplaceStorage.contract.WatchLogs(opts, "OrderCancelled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketplaceStorageOrderCancelled)
				if err := _MarketplaceStorage.contract.UnpackLog(event, "OrderCancelled", log); err != nil {
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
func (_MarketplaceStorage *MarketplaceStorageFilterer) ParseOrderCancelled(log types.Log) (*MarketplaceStorageOrderCancelled, error) {
	event := new(MarketplaceStorageOrderCancelled)
	if err := _MarketplaceStorage.contract.UnpackLog(event, "OrderCancelled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MarketplaceStorageOrderCreatedIterator is returned from FilterOrderCreated and is used to iterate over the raw logs and unpacked data for OrderCreated events raised by the MarketplaceStorage contract.
type MarketplaceStorageOrderCreatedIterator struct {
	Event *MarketplaceStorageOrderCreated // Event containing the contract specifics and raw log

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
func (it *MarketplaceStorageOrderCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketplaceStorageOrderCreated)
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
		it.Event = new(MarketplaceStorageOrderCreated)
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
func (it *MarketplaceStorageOrderCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketplaceStorageOrderCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketplaceStorageOrderCreated represents a OrderCreated event raised by the MarketplaceStorage contract.
type MarketplaceStorageOrderCreated struct {
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
func (_MarketplaceStorage *MarketplaceStorageFilterer) FilterOrderCreated(opts *bind.FilterOpts) (*MarketplaceStorageOrderCreatedIterator, error) {

	logs, sub, err := _MarketplaceStorage.contract.FilterLogs(opts, "OrderCreated")
	if err != nil {
		return nil, err
	}
	return &MarketplaceStorageOrderCreatedIterator{contract: _MarketplaceStorage.contract, event: "OrderCreated", logs: logs, sub: sub}, nil
}

// WatchOrderCreated is a free log subscription operation binding the contract event 0x84c66c3f7ba4b390e20e8e8233e2a516f3ce34a72749e4f12bd010dfba238039.
//
// Solidity: event OrderCreated(bytes32 orderId, uint256 assetId, address seller, address nftAddress, uint256 priceInWei, uint256 expiredAt)
func (_MarketplaceStorage *MarketplaceStorageFilterer) WatchOrderCreated(opts *bind.WatchOpts, sink chan<- *MarketplaceStorageOrderCreated) (event.Subscription, error) {

	logs, sub, err := _MarketplaceStorage.contract.WatchLogs(opts, "OrderCreated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketplaceStorageOrderCreated)
				if err := _MarketplaceStorage.contract.UnpackLog(event, "OrderCreated", log); err != nil {
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
func (_MarketplaceStorage *MarketplaceStorageFilterer) ParseOrderCreated(log types.Log) (*MarketplaceStorageOrderCreated, error) {
	event := new(MarketplaceStorageOrderCreated)
	if err := _MarketplaceStorage.contract.UnpackLog(event, "OrderCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MarketplaceStorageOrderSuccessfulIterator is returned from FilterOrderSuccessful and is used to iterate over the raw logs and unpacked data for OrderSuccessful events raised by the MarketplaceStorage contract.
type MarketplaceStorageOrderSuccessfulIterator struct {
	Event *MarketplaceStorageOrderSuccessful // Event containing the contract specifics and raw log

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
func (it *MarketplaceStorageOrderSuccessfulIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketplaceStorageOrderSuccessful)
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
		it.Event = new(MarketplaceStorageOrderSuccessful)
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
func (it *MarketplaceStorageOrderSuccessfulIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketplaceStorageOrderSuccessfulIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketplaceStorageOrderSuccessful represents a OrderSuccessful event raised by the MarketplaceStorage contract.
type MarketplaceStorageOrderSuccessful struct {
	Id     [32]byte
	Buyer  common.Address
	Seller common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterOrderSuccessful is a free log retrieval operation binding the contract event 0xa3fe2a6ffb1ad0de1c1ee8e3513d205b99980c40af9de82e502095245debabbb.
//
// Solidity: event OrderSuccessful(bytes32 id, address buyer, address seller)
func (_MarketplaceStorage *MarketplaceStorageFilterer) FilterOrderSuccessful(opts *bind.FilterOpts) (*MarketplaceStorageOrderSuccessfulIterator, error) {

	logs, sub, err := _MarketplaceStorage.contract.FilterLogs(opts, "OrderSuccessful")
	if err != nil {
		return nil, err
	}
	return &MarketplaceStorageOrderSuccessfulIterator{contract: _MarketplaceStorage.contract, event: "OrderSuccessful", logs: logs, sub: sub}, nil
}

// WatchOrderSuccessful is a free log subscription operation binding the contract event 0xa3fe2a6ffb1ad0de1c1ee8e3513d205b99980c40af9de82e502095245debabbb.
//
// Solidity: event OrderSuccessful(bytes32 id, address buyer, address seller)
func (_MarketplaceStorage *MarketplaceStorageFilterer) WatchOrderSuccessful(opts *bind.WatchOpts, sink chan<- *MarketplaceStorageOrderSuccessful) (event.Subscription, error) {

	logs, sub, err := _MarketplaceStorage.contract.WatchLogs(opts, "OrderSuccessful")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketplaceStorageOrderSuccessful)
				if err := _MarketplaceStorage.contract.UnpackLog(event, "OrderSuccessful", log); err != nil {
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
func (_MarketplaceStorage *MarketplaceStorageFilterer) ParseOrderSuccessful(log types.Log) (*MarketplaceStorageOrderSuccessful, error) {
	event := new(MarketplaceStorageOrderSuccessful)
	if err := _MarketplaceStorage.contract.UnpackLog(event, "OrderSuccessful", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MarketplaceStorageOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the MarketplaceStorage contract.
type MarketplaceStorageOwnershipTransferredIterator struct {
	Event *MarketplaceStorageOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *MarketplaceStorageOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketplaceStorageOwnershipTransferred)
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
		it.Event = new(MarketplaceStorageOwnershipTransferred)
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
func (it *MarketplaceStorageOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketplaceStorageOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketplaceStorageOwnershipTransferred represents a OwnershipTransferred event raised by the MarketplaceStorage contract.
type MarketplaceStorageOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_MarketplaceStorage *MarketplaceStorageFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*MarketplaceStorageOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _MarketplaceStorage.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &MarketplaceStorageOwnershipTransferredIterator{contract: _MarketplaceStorage.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_MarketplaceStorage *MarketplaceStorageFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *MarketplaceStorageOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _MarketplaceStorage.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketplaceStorageOwnershipTransferred)
				if err := _MarketplaceStorage.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_MarketplaceStorage *MarketplaceStorageFilterer) ParseOwnershipTransferred(log types.Log) (*MarketplaceStorageOwnershipTransferred, error) {
	event := new(MarketplaceStorageOwnershipTransferred)
	if err := _MarketplaceStorage.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MarketplaceStoragePausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the MarketplaceStorage contract.
type MarketplaceStoragePausedIterator struct {
	Event *MarketplaceStoragePaused // Event containing the contract specifics and raw log

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
func (it *MarketplaceStoragePausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketplaceStoragePaused)
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
		it.Event = new(MarketplaceStoragePaused)
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
func (it *MarketplaceStoragePausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketplaceStoragePausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketplaceStoragePaused represents a Paused event raised by the MarketplaceStorage contract.
type MarketplaceStoragePaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_MarketplaceStorage *MarketplaceStorageFilterer) FilterPaused(opts *bind.FilterOpts) (*MarketplaceStoragePausedIterator, error) {

	logs, sub, err := _MarketplaceStorage.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &MarketplaceStoragePausedIterator{contract: _MarketplaceStorage.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_MarketplaceStorage *MarketplaceStorageFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *MarketplaceStoragePaused) (event.Subscription, error) {

	logs, sub, err := _MarketplaceStorage.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketplaceStoragePaused)
				if err := _MarketplaceStorage.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_MarketplaceStorage *MarketplaceStorageFilterer) ParsePaused(log types.Log) (*MarketplaceStoragePaused, error) {
	event := new(MarketplaceStoragePaused)
	if err := _MarketplaceStorage.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MarketplaceStorageRevealFailedIterator is returned from FilterRevealFailed and is used to iterate over the raw logs and unpacked data for RevealFailed events raised by the MarketplaceStorage contract.
type MarketplaceStorageRevealFailedIterator struct {
	Event *MarketplaceStorageRevealFailed // Event containing the contract specifics and raw log

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
func (it *MarketplaceStorageRevealFailedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketplaceStorageRevealFailed)
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
		it.Event = new(MarketplaceStorageRevealFailed)
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
func (it *MarketplaceStorageRevealFailedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketplaceStorageRevealFailedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketplaceStorageRevealFailed represents a RevealFailed event raised by the MarketplaceStorage contract.
type MarketplaceStorageRevealFailed struct {
	Fake      bool
	Revealer  common.Address
	AuctionId [32]byte
	Value     *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterRevealFailed is a free log retrieval operation binding the contract event 0x255ef988657daaeb7f921a0c776fa6efb221748fb64e61ab1febb0fd37f67175.
//
// Solidity: event RevealFailed(bool fake, address revealer, bytes32 auctionId, uint256 value)
func (_MarketplaceStorage *MarketplaceStorageFilterer) FilterRevealFailed(opts *bind.FilterOpts) (*MarketplaceStorageRevealFailedIterator, error) {

	logs, sub, err := _MarketplaceStorage.contract.FilterLogs(opts, "RevealFailed")
	if err != nil {
		return nil, err
	}
	return &MarketplaceStorageRevealFailedIterator{contract: _MarketplaceStorage.contract, event: "RevealFailed", logs: logs, sub: sub}, nil
}

// WatchRevealFailed is a free log subscription operation binding the contract event 0x255ef988657daaeb7f921a0c776fa6efb221748fb64e61ab1febb0fd37f67175.
//
// Solidity: event RevealFailed(bool fake, address revealer, bytes32 auctionId, uint256 value)
func (_MarketplaceStorage *MarketplaceStorageFilterer) WatchRevealFailed(opts *bind.WatchOpts, sink chan<- *MarketplaceStorageRevealFailed) (event.Subscription, error) {

	logs, sub, err := _MarketplaceStorage.contract.WatchLogs(opts, "RevealFailed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketplaceStorageRevealFailed)
				if err := _MarketplaceStorage.contract.UnpackLog(event, "RevealFailed", log); err != nil {
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

// ParseRevealFailed is a log parse operation binding the contract event 0x255ef988657daaeb7f921a0c776fa6efb221748fb64e61ab1febb0fd37f67175.
//
// Solidity: event RevealFailed(bool fake, address revealer, bytes32 auctionId, uint256 value)
func (_MarketplaceStorage *MarketplaceStorageFilterer) ParseRevealFailed(log types.Log) (*MarketplaceStorageRevealFailed, error) {
	event := new(MarketplaceStorageRevealFailed)
	if err := _MarketplaceStorage.contract.UnpackLog(event, "RevealFailed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MarketplaceStorageRevealSuccessfulIterator is returned from FilterRevealSuccessful and is used to iterate over the raw logs and unpacked data for RevealSuccessful events raised by the MarketplaceStorage contract.
type MarketplaceStorageRevealSuccessfulIterator struct {
	Event *MarketplaceStorageRevealSuccessful // Event containing the contract specifics and raw log

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
func (it *MarketplaceStorageRevealSuccessfulIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketplaceStorageRevealSuccessful)
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
		it.Event = new(MarketplaceStorageRevealSuccessful)
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
func (it *MarketplaceStorageRevealSuccessfulIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketplaceStorageRevealSuccessfulIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketplaceStorageRevealSuccessful represents a RevealSuccessful event raised by the MarketplaceStorage contract.
type MarketplaceStorageRevealSuccessful struct {
	Fake       bool
	Revealer   common.Address
	AuctionId  [32]byte
	Value      *big.Int
	BlindedBid [32]byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterRevealSuccessful is a free log retrieval operation binding the contract event 0xabef59dc3ae014d197fad42649c58d34bfc816d3e1a7f26ca32c13611b13e7a1.
//
// Solidity: event RevealSuccessful(bool fake, address revealer, bytes32 auctionId, uint256 value, bytes32 blindedBid)
func (_MarketplaceStorage *MarketplaceStorageFilterer) FilterRevealSuccessful(opts *bind.FilterOpts) (*MarketplaceStorageRevealSuccessfulIterator, error) {

	logs, sub, err := _MarketplaceStorage.contract.FilterLogs(opts, "RevealSuccessful")
	if err != nil {
		return nil, err
	}
	return &MarketplaceStorageRevealSuccessfulIterator{contract: _MarketplaceStorage.contract, event: "RevealSuccessful", logs: logs, sub: sub}, nil
}

// WatchRevealSuccessful is a free log subscription operation binding the contract event 0xabef59dc3ae014d197fad42649c58d34bfc816d3e1a7f26ca32c13611b13e7a1.
//
// Solidity: event RevealSuccessful(bool fake, address revealer, bytes32 auctionId, uint256 value, bytes32 blindedBid)
func (_MarketplaceStorage *MarketplaceStorageFilterer) WatchRevealSuccessful(opts *bind.WatchOpts, sink chan<- *MarketplaceStorageRevealSuccessful) (event.Subscription, error) {

	logs, sub, err := _MarketplaceStorage.contract.WatchLogs(opts, "RevealSuccessful")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketplaceStorageRevealSuccessful)
				if err := _MarketplaceStorage.contract.UnpackLog(event, "RevealSuccessful", log); err != nil {
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

// ParseRevealSuccessful is a log parse operation binding the contract event 0xabef59dc3ae014d197fad42649c58d34bfc816d3e1a7f26ca32c13611b13e7a1.
//
// Solidity: event RevealSuccessful(bool fake, address revealer, bytes32 auctionId, uint256 value, bytes32 blindedBid)
func (_MarketplaceStorage *MarketplaceStorageFilterer) ParseRevealSuccessful(log types.Log) (*MarketplaceStorageRevealSuccessful, error) {
	event := new(MarketplaceStorageRevealSuccessful)
	if err := _MarketplaceStorage.contract.UnpackLog(event, "RevealSuccessful", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MarketplaceStorageRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the MarketplaceStorage contract.
type MarketplaceStorageRoleAdminChangedIterator struct {
	Event *MarketplaceStorageRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *MarketplaceStorageRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketplaceStorageRoleAdminChanged)
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
		it.Event = new(MarketplaceStorageRoleAdminChanged)
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
func (it *MarketplaceStorageRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketplaceStorageRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketplaceStorageRoleAdminChanged represents a RoleAdminChanged event raised by the MarketplaceStorage contract.
type MarketplaceStorageRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_MarketplaceStorage *MarketplaceStorageFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*MarketplaceStorageRoleAdminChangedIterator, error) {

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

	logs, sub, err := _MarketplaceStorage.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &MarketplaceStorageRoleAdminChangedIterator{contract: _MarketplaceStorage.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_MarketplaceStorage *MarketplaceStorageFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *MarketplaceStorageRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _MarketplaceStorage.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketplaceStorageRoleAdminChanged)
				if err := _MarketplaceStorage.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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
func (_MarketplaceStorage *MarketplaceStorageFilterer) ParseRoleAdminChanged(log types.Log) (*MarketplaceStorageRoleAdminChanged, error) {
	event := new(MarketplaceStorageRoleAdminChanged)
	if err := _MarketplaceStorage.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MarketplaceStorageRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the MarketplaceStorage contract.
type MarketplaceStorageRoleGrantedIterator struct {
	Event *MarketplaceStorageRoleGranted // Event containing the contract specifics and raw log

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
func (it *MarketplaceStorageRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketplaceStorageRoleGranted)
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
		it.Event = new(MarketplaceStorageRoleGranted)
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
func (it *MarketplaceStorageRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketplaceStorageRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketplaceStorageRoleGranted represents a RoleGranted event raised by the MarketplaceStorage contract.
type MarketplaceStorageRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_MarketplaceStorage *MarketplaceStorageFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*MarketplaceStorageRoleGrantedIterator, error) {

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

	logs, sub, err := _MarketplaceStorage.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &MarketplaceStorageRoleGrantedIterator{contract: _MarketplaceStorage.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_MarketplaceStorage *MarketplaceStorageFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *MarketplaceStorageRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _MarketplaceStorage.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketplaceStorageRoleGranted)
				if err := _MarketplaceStorage.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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
func (_MarketplaceStorage *MarketplaceStorageFilterer) ParseRoleGranted(log types.Log) (*MarketplaceStorageRoleGranted, error) {
	event := new(MarketplaceStorageRoleGranted)
	if err := _MarketplaceStorage.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MarketplaceStorageRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the MarketplaceStorage contract.
type MarketplaceStorageRoleRevokedIterator struct {
	Event *MarketplaceStorageRoleRevoked // Event containing the contract specifics and raw log

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
func (it *MarketplaceStorageRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketplaceStorageRoleRevoked)
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
		it.Event = new(MarketplaceStorageRoleRevoked)
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
func (it *MarketplaceStorageRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketplaceStorageRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketplaceStorageRoleRevoked represents a RoleRevoked event raised by the MarketplaceStorage contract.
type MarketplaceStorageRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_MarketplaceStorage *MarketplaceStorageFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*MarketplaceStorageRoleRevokedIterator, error) {

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

	logs, sub, err := _MarketplaceStorage.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &MarketplaceStorageRoleRevokedIterator{contract: _MarketplaceStorage.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_MarketplaceStorage *MarketplaceStorageFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *MarketplaceStorageRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _MarketplaceStorage.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketplaceStorageRoleRevoked)
				if err := _MarketplaceStorage.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
func (_MarketplaceStorage *MarketplaceStorageFilterer) ParseRoleRevoked(log types.Log) (*MarketplaceStorageRoleRevoked, error) {
	event := new(MarketplaceStorageRoleRevoked)
	if err := _MarketplaceStorage.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MarketplaceStorageUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the MarketplaceStorage contract.
type MarketplaceStorageUnpausedIterator struct {
	Event *MarketplaceStorageUnpaused // Event containing the contract specifics and raw log

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
func (it *MarketplaceStorageUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketplaceStorageUnpaused)
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
		it.Event = new(MarketplaceStorageUnpaused)
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
func (it *MarketplaceStorageUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketplaceStorageUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketplaceStorageUnpaused represents a Unpaused event raised by the MarketplaceStorage contract.
type MarketplaceStorageUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_MarketplaceStorage *MarketplaceStorageFilterer) FilterUnpaused(opts *bind.FilterOpts) (*MarketplaceStorageUnpausedIterator, error) {

	logs, sub, err := _MarketplaceStorage.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &MarketplaceStorageUnpausedIterator{contract: _MarketplaceStorage.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_MarketplaceStorage *MarketplaceStorageFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *MarketplaceStorageUnpaused) (event.Subscription, error) {

	logs, sub, err := _MarketplaceStorage.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketplaceStorageUnpaused)
				if err := _MarketplaceStorage.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
func (_MarketplaceStorage *MarketplaceStorageFilterer) ParseUnpaused(log types.Log) (*MarketplaceStorageUnpaused, error) {
	event := new(MarketplaceStorageUnpaused)
	if err := _MarketplaceStorage.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
