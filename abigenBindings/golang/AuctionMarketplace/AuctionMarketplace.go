// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package AuctionMarketplace

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

// AuctionMarketplaceMetaData contains all meta data concerning the AuctionMarketplace contract.
var AuctionMarketplaceMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"acceptedToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"marketplaceStorage\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"beneficary\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"ownerCutPerMillion\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"Bidded\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidBiddingEnd\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidContract\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidCut\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidPrice\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidRevealEnd\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotExisted\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotRunning\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"}],\"name\":\"TooEarly\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"}],\"name\":\"TooLate\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Unauthorized\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Unavailable\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"AddedBlackList\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"auctionId\",\"type\":\"bytes32\"}],\"name\":\"AuctionCancelled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"nftAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"auctionId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"assetId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"biddingEnd\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"revealEnd\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"startPriceInWei\",\"type\":\"uint256\"}],\"name\":\"AuctionCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"auctionId\",\"type\":\"bytes32\"}],\"name\":\"AuctionEnded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"bidder\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"auctionId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"deposit\",\"type\":\"uint256\"}],\"name\":\"AuctionRefund\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"buyer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"auctionId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalPrice\",\"type\":\"uint256\"}],\"name\":\"AuctionSuccessful\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"bidder\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"auctionId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"blindedBid\",\"type\":\"bytes32\"}],\"name\":\"BidSuccessful\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"ownerCutPerMillion\",\"type\":\"uint256\"}],\"name\":\"ChangedOwnerCutPerMillion\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"publicationFee\",\"type\":\"uint256\"}],\"name\":\"ChangedPublicationFee\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_marketplaceStorage\",\"type\":\"address\"}],\"name\":\"MarketplaceStorageUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"RemovedBlackList\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"fake\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"revealer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"auctionId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"RevealFailed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"fake\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"revealer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"auctionId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"blindedBid\",\"type\":\"bytes32\"}],\"name\":\"RevealSuccessful\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"CANCEL_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ERC721_Interface\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"IMarketplaceStorage_Interface\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PAUSER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"acceptedToken\",\"outputs\":[{\"internalType\":\"contractIBEP20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_evilUser\",\"type\":\"address\"}],\"name\":\"addBlackList\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"beneficary\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_maker\",\"type\":\"address\"}],\"name\":\"getBlackListStatus\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getRoleMember\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleMemberCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"isBlackListed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"marketplaceStorage\",\"outputs\":[{\"internalType\":\"contractIMarketplaceStorage\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minStageDuration\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ownerCutPerMillion\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"publicationFeeInWei\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_clearedUser\",\"type\":\"address\"}],\"name\":\"removeBlackList\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_beneficary\",\"type\":\"address\"}],\"name\":\"setBeneficary\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"}],\"name\":\"setMinStageDuration\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_ownerCutPerMillion\",\"type\":\"uint256\"}],\"name\":\"setOwnerCutPerMillion\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_publicationFee\",\"type\":\"uint256\"}],\"name\":\"setPublicationFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_marketplaceStorage\",\"type\":\"address\"}],\"name\":\"updateStorageAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"nftAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"assetId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startPriceInWei\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"biddingEnd\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"revealEnd\",\"type\":\"uint256\"}],\"name\":\"createAuction\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"nftAsset\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"auctionId\",\"type\":\"bytes32\"}],\"name\":\"cancelAuction\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"nftAsset\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"auctionId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"blindedBid\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"deposit\",\"type\":\"uint256\"}],\"name\":\"bidAuction\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"nftAsset\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"auctionId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[]\",\"name\":\"_values\",\"type\":\"uint256[]\"},{\"internalType\":\"bool[]\",\"name\":\"_fake\",\"type\":\"bool[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"_secret\",\"type\":\"bytes32[]\"}],\"name\":\"revealBid\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"auctionId\",\"type\":\"bytes32\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"auctionId\",\"type\":\"bytes32\"}],\"name\":\"pendingReturns\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"nftAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"assetId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"auctionId\",\"type\":\"bytes32\"}],\"name\":\"closeAuction\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"auctionId\",\"type\":\"bytes32\"}],\"name\":\"checkExisted\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"nftAsset\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"auctionId\",\"type\":\"bytes32\"}],\"name\":\"checkRunning\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// AuctionMarketplaceABI is the input ABI used to generate the binding from.
// Deprecated: Use AuctionMarketplaceMetaData.ABI instead.
var AuctionMarketplaceABI = AuctionMarketplaceMetaData.ABI

// AuctionMarketplace is an auto generated Go binding around an Ethereum contract.
type AuctionMarketplace struct {
	AuctionMarketplaceCaller     // Read-only binding to the contract
	AuctionMarketplaceTransactor // Write-only binding to the contract
	AuctionMarketplaceFilterer   // Log filterer for contract events
}

// AuctionMarketplaceCaller is an auto generated read-only Go binding around an Ethereum contract.
type AuctionMarketplaceCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AuctionMarketplaceTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AuctionMarketplaceTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AuctionMarketplaceFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AuctionMarketplaceFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AuctionMarketplaceSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AuctionMarketplaceSession struct {
	Contract     *AuctionMarketplace // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// AuctionMarketplaceCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AuctionMarketplaceCallerSession struct {
	Contract *AuctionMarketplaceCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// AuctionMarketplaceTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AuctionMarketplaceTransactorSession struct {
	Contract     *AuctionMarketplaceTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// AuctionMarketplaceRaw is an auto generated low-level Go binding around an Ethereum contract.
type AuctionMarketplaceRaw struct {
	Contract *AuctionMarketplace // Generic contract binding to access the raw methods on
}

// AuctionMarketplaceCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AuctionMarketplaceCallerRaw struct {
	Contract *AuctionMarketplaceCaller // Generic read-only contract binding to access the raw methods on
}

// AuctionMarketplaceTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AuctionMarketplaceTransactorRaw struct {
	Contract *AuctionMarketplaceTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAuctionMarketplace creates a new instance of AuctionMarketplace, bound to a specific deployed contract.
func NewAuctionMarketplace(address common.Address, backend bind.ContractBackend) (*AuctionMarketplace, error) {
	contract, err := bindAuctionMarketplace(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AuctionMarketplace{AuctionMarketplaceCaller: AuctionMarketplaceCaller{contract: contract}, AuctionMarketplaceTransactor: AuctionMarketplaceTransactor{contract: contract}, AuctionMarketplaceFilterer: AuctionMarketplaceFilterer{contract: contract}}, nil
}

// NewAuctionMarketplaceCaller creates a new read-only instance of AuctionMarketplace, bound to a specific deployed contract.
func NewAuctionMarketplaceCaller(address common.Address, caller bind.ContractCaller) (*AuctionMarketplaceCaller, error) {
	contract, err := bindAuctionMarketplace(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AuctionMarketplaceCaller{contract: contract}, nil
}

// NewAuctionMarketplaceTransactor creates a new write-only instance of AuctionMarketplace, bound to a specific deployed contract.
func NewAuctionMarketplaceTransactor(address common.Address, transactor bind.ContractTransactor) (*AuctionMarketplaceTransactor, error) {
	contract, err := bindAuctionMarketplace(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AuctionMarketplaceTransactor{contract: contract}, nil
}

// NewAuctionMarketplaceFilterer creates a new log filterer instance of AuctionMarketplace, bound to a specific deployed contract.
func NewAuctionMarketplaceFilterer(address common.Address, filterer bind.ContractFilterer) (*AuctionMarketplaceFilterer, error) {
	contract, err := bindAuctionMarketplace(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AuctionMarketplaceFilterer{contract: contract}, nil
}

// bindAuctionMarketplace binds a generic wrapper to an already deployed contract.
func bindAuctionMarketplace(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AuctionMarketplaceABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AuctionMarketplace *AuctionMarketplaceRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AuctionMarketplace.Contract.AuctionMarketplaceCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AuctionMarketplace *AuctionMarketplaceRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AuctionMarketplace.Contract.AuctionMarketplaceTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AuctionMarketplace *AuctionMarketplaceRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AuctionMarketplace.Contract.AuctionMarketplaceTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AuctionMarketplace *AuctionMarketplaceCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AuctionMarketplace.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AuctionMarketplace *AuctionMarketplaceTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AuctionMarketplace.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AuctionMarketplace *AuctionMarketplaceTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AuctionMarketplace.Contract.contract.Transact(opts, method, params...)
}

// ADMINROLE is a free data retrieval call binding the contract method 0x75b238fc.
//
// Solidity: function ADMIN_ROLE() view returns(bytes32)
func (_AuctionMarketplace *AuctionMarketplaceCaller) ADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _AuctionMarketplace.contract.Call(opts, &out, "ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ADMINROLE is a free data retrieval call binding the contract method 0x75b238fc.
//
// Solidity: function ADMIN_ROLE() view returns(bytes32)
func (_AuctionMarketplace *AuctionMarketplaceSession) ADMINROLE() ([32]byte, error) {
	return _AuctionMarketplace.Contract.ADMINROLE(&_AuctionMarketplace.CallOpts)
}

// ADMINROLE is a free data retrieval call binding the contract method 0x75b238fc.
//
// Solidity: function ADMIN_ROLE() view returns(bytes32)
func (_AuctionMarketplace *AuctionMarketplaceCallerSession) ADMINROLE() ([32]byte, error) {
	return _AuctionMarketplace.Contract.ADMINROLE(&_AuctionMarketplace.CallOpts)
}

// CANCELROLE is a free data retrieval call binding the contract method 0x13c27ca7.
//
// Solidity: function CANCEL_ROLE() view returns(bytes32)
func (_AuctionMarketplace *AuctionMarketplaceCaller) CANCELROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _AuctionMarketplace.contract.Call(opts, &out, "CANCEL_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// CANCELROLE is a free data retrieval call binding the contract method 0x13c27ca7.
//
// Solidity: function CANCEL_ROLE() view returns(bytes32)
func (_AuctionMarketplace *AuctionMarketplaceSession) CANCELROLE() ([32]byte, error) {
	return _AuctionMarketplace.Contract.CANCELROLE(&_AuctionMarketplace.CallOpts)
}

// CANCELROLE is a free data retrieval call binding the contract method 0x13c27ca7.
//
// Solidity: function CANCEL_ROLE() view returns(bytes32)
func (_AuctionMarketplace *AuctionMarketplaceCallerSession) CANCELROLE() ([32]byte, error) {
	return _AuctionMarketplace.Contract.CANCELROLE(&_AuctionMarketplace.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_AuctionMarketplace *AuctionMarketplaceCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _AuctionMarketplace.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_AuctionMarketplace *AuctionMarketplaceSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _AuctionMarketplace.Contract.DEFAULTADMINROLE(&_AuctionMarketplace.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_AuctionMarketplace *AuctionMarketplaceCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _AuctionMarketplace.Contract.DEFAULTADMINROLE(&_AuctionMarketplace.CallOpts)
}

// ERC721Interface is a free data retrieval call binding the contract method 0x2b4c32be.
//
// Solidity: function ERC721_Interface() view returns(bytes4)
func (_AuctionMarketplace *AuctionMarketplaceCaller) ERC721Interface(opts *bind.CallOpts) ([4]byte, error) {
	var out []interface{}
	err := _AuctionMarketplace.contract.Call(opts, &out, "ERC721_Interface")

	if err != nil {
		return *new([4]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)

	return out0, err

}

// ERC721Interface is a free data retrieval call binding the contract method 0x2b4c32be.
//
// Solidity: function ERC721_Interface() view returns(bytes4)
func (_AuctionMarketplace *AuctionMarketplaceSession) ERC721Interface() ([4]byte, error) {
	return _AuctionMarketplace.Contract.ERC721Interface(&_AuctionMarketplace.CallOpts)
}

// ERC721Interface is a free data retrieval call binding the contract method 0x2b4c32be.
//
// Solidity: function ERC721_Interface() view returns(bytes4)
func (_AuctionMarketplace *AuctionMarketplaceCallerSession) ERC721Interface() ([4]byte, error) {
	return _AuctionMarketplace.Contract.ERC721Interface(&_AuctionMarketplace.CallOpts)
}

// IMarketplaceStorageInterface is a free data retrieval call binding the contract method 0x8bafe559.
//
// Solidity: function IMarketplaceStorage_Interface() view returns(bytes4)
func (_AuctionMarketplace *AuctionMarketplaceCaller) IMarketplaceStorageInterface(opts *bind.CallOpts) ([4]byte, error) {
	var out []interface{}
	err := _AuctionMarketplace.contract.Call(opts, &out, "IMarketplaceStorage_Interface")

	if err != nil {
		return *new([4]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)

	return out0, err

}

// IMarketplaceStorageInterface is a free data retrieval call binding the contract method 0x8bafe559.
//
// Solidity: function IMarketplaceStorage_Interface() view returns(bytes4)
func (_AuctionMarketplace *AuctionMarketplaceSession) IMarketplaceStorageInterface() ([4]byte, error) {
	return _AuctionMarketplace.Contract.IMarketplaceStorageInterface(&_AuctionMarketplace.CallOpts)
}

// IMarketplaceStorageInterface is a free data retrieval call binding the contract method 0x8bafe559.
//
// Solidity: function IMarketplaceStorage_Interface() view returns(bytes4)
func (_AuctionMarketplace *AuctionMarketplaceCallerSession) IMarketplaceStorageInterface() ([4]byte, error) {
	return _AuctionMarketplace.Contract.IMarketplaceStorageInterface(&_AuctionMarketplace.CallOpts)
}

// PAUSERROLE is a free data retrieval call binding the contract method 0xe63ab1e9.
//
// Solidity: function PAUSER_ROLE() view returns(bytes32)
func (_AuctionMarketplace *AuctionMarketplaceCaller) PAUSERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _AuctionMarketplace.contract.Call(opts, &out, "PAUSER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// PAUSERROLE is a free data retrieval call binding the contract method 0xe63ab1e9.
//
// Solidity: function PAUSER_ROLE() view returns(bytes32)
func (_AuctionMarketplace *AuctionMarketplaceSession) PAUSERROLE() ([32]byte, error) {
	return _AuctionMarketplace.Contract.PAUSERROLE(&_AuctionMarketplace.CallOpts)
}

// PAUSERROLE is a free data retrieval call binding the contract method 0xe63ab1e9.
//
// Solidity: function PAUSER_ROLE() view returns(bytes32)
func (_AuctionMarketplace *AuctionMarketplaceCallerSession) PAUSERROLE() ([32]byte, error) {
	return _AuctionMarketplace.Contract.PAUSERROLE(&_AuctionMarketplace.CallOpts)
}

// AcceptedToken is a free data retrieval call binding the contract method 0x451c3d80.
//
// Solidity: function acceptedToken() view returns(address)
func (_AuctionMarketplace *AuctionMarketplaceCaller) AcceptedToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AuctionMarketplace.contract.Call(opts, &out, "acceptedToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AcceptedToken is a free data retrieval call binding the contract method 0x451c3d80.
//
// Solidity: function acceptedToken() view returns(address)
func (_AuctionMarketplace *AuctionMarketplaceSession) AcceptedToken() (common.Address, error) {
	return _AuctionMarketplace.Contract.AcceptedToken(&_AuctionMarketplace.CallOpts)
}

// AcceptedToken is a free data retrieval call binding the contract method 0x451c3d80.
//
// Solidity: function acceptedToken() view returns(address)
func (_AuctionMarketplace *AuctionMarketplaceCallerSession) AcceptedToken() (common.Address, error) {
	return _AuctionMarketplace.Contract.AcceptedToken(&_AuctionMarketplace.CallOpts)
}

// Beneficary is a free data retrieval call binding the contract method 0xcbfda1c5.
//
// Solidity: function beneficary() view returns(address)
func (_AuctionMarketplace *AuctionMarketplaceCaller) Beneficary(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AuctionMarketplace.contract.Call(opts, &out, "beneficary")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Beneficary is a free data retrieval call binding the contract method 0xcbfda1c5.
//
// Solidity: function beneficary() view returns(address)
func (_AuctionMarketplace *AuctionMarketplaceSession) Beneficary() (common.Address, error) {
	return _AuctionMarketplace.Contract.Beneficary(&_AuctionMarketplace.CallOpts)
}

// Beneficary is a free data retrieval call binding the contract method 0xcbfda1c5.
//
// Solidity: function beneficary() view returns(address)
func (_AuctionMarketplace *AuctionMarketplaceCallerSession) Beneficary() (common.Address, error) {
	return _AuctionMarketplace.Contract.Beneficary(&_AuctionMarketplace.CallOpts)
}

// CheckExisted is a free data retrieval call binding the contract method 0x00c5c3c3.
//
// Solidity: function checkExisted(bytes32 auctionId) view returns()
func (_AuctionMarketplace *AuctionMarketplaceCaller) CheckExisted(opts *bind.CallOpts, auctionId [32]byte) error {
	var out []interface{}
	err := _AuctionMarketplace.contract.Call(opts, &out, "checkExisted", auctionId)

	if err != nil {
		return err
	}

	return err

}

// CheckExisted is a free data retrieval call binding the contract method 0x00c5c3c3.
//
// Solidity: function checkExisted(bytes32 auctionId) view returns()
func (_AuctionMarketplace *AuctionMarketplaceSession) CheckExisted(auctionId [32]byte) error {
	return _AuctionMarketplace.Contract.CheckExisted(&_AuctionMarketplace.CallOpts, auctionId)
}

// CheckExisted is a free data retrieval call binding the contract method 0x00c5c3c3.
//
// Solidity: function checkExisted(bytes32 auctionId) view returns()
func (_AuctionMarketplace *AuctionMarketplaceCallerSession) CheckExisted(auctionId [32]byte) error {
	return _AuctionMarketplace.Contract.CheckExisted(&_AuctionMarketplace.CallOpts, auctionId)
}

// CheckRunning is a free data retrieval call binding the contract method 0x076f6dc7.
//
// Solidity: function checkRunning(bytes32 nftAsset, bytes32 auctionId) view returns()
func (_AuctionMarketplace *AuctionMarketplaceCaller) CheckRunning(opts *bind.CallOpts, nftAsset [32]byte, auctionId [32]byte) error {
	var out []interface{}
	err := _AuctionMarketplace.contract.Call(opts, &out, "checkRunning", nftAsset, auctionId)

	if err != nil {
		return err
	}

	return err

}

// CheckRunning is a free data retrieval call binding the contract method 0x076f6dc7.
//
// Solidity: function checkRunning(bytes32 nftAsset, bytes32 auctionId) view returns()
func (_AuctionMarketplace *AuctionMarketplaceSession) CheckRunning(nftAsset [32]byte, auctionId [32]byte) error {
	return _AuctionMarketplace.Contract.CheckRunning(&_AuctionMarketplace.CallOpts, nftAsset, auctionId)
}

// CheckRunning is a free data retrieval call binding the contract method 0x076f6dc7.
//
// Solidity: function checkRunning(bytes32 nftAsset, bytes32 auctionId) view returns()
func (_AuctionMarketplace *AuctionMarketplaceCallerSession) CheckRunning(nftAsset [32]byte, auctionId [32]byte) error {
	return _AuctionMarketplace.Contract.CheckRunning(&_AuctionMarketplace.CallOpts, nftAsset, auctionId)
}

// GetBlackListStatus is a free data retrieval call binding the contract method 0x59bf1abe.
//
// Solidity: function getBlackListStatus(address _maker) view returns(bool)
func (_AuctionMarketplace *AuctionMarketplaceCaller) GetBlackListStatus(opts *bind.CallOpts, _maker common.Address) (bool, error) {
	var out []interface{}
	err := _AuctionMarketplace.contract.Call(opts, &out, "getBlackListStatus", _maker)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// GetBlackListStatus is a free data retrieval call binding the contract method 0x59bf1abe.
//
// Solidity: function getBlackListStatus(address _maker) view returns(bool)
func (_AuctionMarketplace *AuctionMarketplaceSession) GetBlackListStatus(_maker common.Address) (bool, error) {
	return _AuctionMarketplace.Contract.GetBlackListStatus(&_AuctionMarketplace.CallOpts, _maker)
}

// GetBlackListStatus is a free data retrieval call binding the contract method 0x59bf1abe.
//
// Solidity: function getBlackListStatus(address _maker) view returns(bool)
func (_AuctionMarketplace *AuctionMarketplaceCallerSession) GetBlackListStatus(_maker common.Address) (bool, error) {
	return _AuctionMarketplace.Contract.GetBlackListStatus(&_AuctionMarketplace.CallOpts, _maker)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_AuctionMarketplace *AuctionMarketplaceCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _AuctionMarketplace.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_AuctionMarketplace *AuctionMarketplaceSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _AuctionMarketplace.Contract.GetRoleAdmin(&_AuctionMarketplace.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_AuctionMarketplace *AuctionMarketplaceCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _AuctionMarketplace.Contract.GetRoleAdmin(&_AuctionMarketplace.CallOpts, role)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_AuctionMarketplace *AuctionMarketplaceCaller) GetRoleMember(opts *bind.CallOpts, role [32]byte, index *big.Int) (common.Address, error) {
	var out []interface{}
	err := _AuctionMarketplace.contract.Call(opts, &out, "getRoleMember", role, index)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_AuctionMarketplace *AuctionMarketplaceSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _AuctionMarketplace.Contract.GetRoleMember(&_AuctionMarketplace.CallOpts, role, index)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_AuctionMarketplace *AuctionMarketplaceCallerSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _AuctionMarketplace.Contract.GetRoleMember(&_AuctionMarketplace.CallOpts, role, index)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_AuctionMarketplace *AuctionMarketplaceCaller) GetRoleMemberCount(opts *bind.CallOpts, role [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _AuctionMarketplace.contract.Call(opts, &out, "getRoleMemberCount", role)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_AuctionMarketplace *AuctionMarketplaceSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _AuctionMarketplace.Contract.GetRoleMemberCount(&_AuctionMarketplace.CallOpts, role)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_AuctionMarketplace *AuctionMarketplaceCallerSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _AuctionMarketplace.Contract.GetRoleMemberCount(&_AuctionMarketplace.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_AuctionMarketplace *AuctionMarketplaceCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _AuctionMarketplace.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_AuctionMarketplace *AuctionMarketplaceSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _AuctionMarketplace.Contract.HasRole(&_AuctionMarketplace.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_AuctionMarketplace *AuctionMarketplaceCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _AuctionMarketplace.Contract.HasRole(&_AuctionMarketplace.CallOpts, role, account)
}

// IsBlackListed is a free data retrieval call binding the contract method 0xe47d6060.
//
// Solidity: function isBlackListed(address ) view returns(bool)
func (_AuctionMarketplace *AuctionMarketplaceCaller) IsBlackListed(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _AuctionMarketplace.contract.Call(opts, &out, "isBlackListed", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsBlackListed is a free data retrieval call binding the contract method 0xe47d6060.
//
// Solidity: function isBlackListed(address ) view returns(bool)
func (_AuctionMarketplace *AuctionMarketplaceSession) IsBlackListed(arg0 common.Address) (bool, error) {
	return _AuctionMarketplace.Contract.IsBlackListed(&_AuctionMarketplace.CallOpts, arg0)
}

// IsBlackListed is a free data retrieval call binding the contract method 0xe47d6060.
//
// Solidity: function isBlackListed(address ) view returns(bool)
func (_AuctionMarketplace *AuctionMarketplaceCallerSession) IsBlackListed(arg0 common.Address) (bool, error) {
	return _AuctionMarketplace.Contract.IsBlackListed(&_AuctionMarketplace.CallOpts, arg0)
}

// MarketplaceStorage is a free data retrieval call binding the contract method 0x46b3aec6.
//
// Solidity: function marketplaceStorage() view returns(address)
func (_AuctionMarketplace *AuctionMarketplaceCaller) MarketplaceStorage(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AuctionMarketplace.contract.Call(opts, &out, "marketplaceStorage")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// MarketplaceStorage is a free data retrieval call binding the contract method 0x46b3aec6.
//
// Solidity: function marketplaceStorage() view returns(address)
func (_AuctionMarketplace *AuctionMarketplaceSession) MarketplaceStorage() (common.Address, error) {
	return _AuctionMarketplace.Contract.MarketplaceStorage(&_AuctionMarketplace.CallOpts)
}

// MarketplaceStorage is a free data retrieval call binding the contract method 0x46b3aec6.
//
// Solidity: function marketplaceStorage() view returns(address)
func (_AuctionMarketplace *AuctionMarketplaceCallerSession) MarketplaceStorage() (common.Address, error) {
	return _AuctionMarketplace.Contract.MarketplaceStorage(&_AuctionMarketplace.CallOpts)
}

// MinStageDuration is a free data retrieval call binding the contract method 0xb42cf929.
//
// Solidity: function minStageDuration() view returns(uint256)
func (_AuctionMarketplace *AuctionMarketplaceCaller) MinStageDuration(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AuctionMarketplace.contract.Call(opts, &out, "minStageDuration")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinStageDuration is a free data retrieval call binding the contract method 0xb42cf929.
//
// Solidity: function minStageDuration() view returns(uint256)
func (_AuctionMarketplace *AuctionMarketplaceSession) MinStageDuration() (*big.Int, error) {
	return _AuctionMarketplace.Contract.MinStageDuration(&_AuctionMarketplace.CallOpts)
}

// MinStageDuration is a free data retrieval call binding the contract method 0xb42cf929.
//
// Solidity: function minStageDuration() view returns(uint256)
func (_AuctionMarketplace *AuctionMarketplaceCallerSession) MinStageDuration() (*big.Int, error) {
	return _AuctionMarketplace.Contract.MinStageDuration(&_AuctionMarketplace.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_AuctionMarketplace *AuctionMarketplaceCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AuctionMarketplace.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_AuctionMarketplace *AuctionMarketplaceSession) Owner() (common.Address, error) {
	return _AuctionMarketplace.Contract.Owner(&_AuctionMarketplace.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_AuctionMarketplace *AuctionMarketplaceCallerSession) Owner() (common.Address, error) {
	return _AuctionMarketplace.Contract.Owner(&_AuctionMarketplace.CallOpts)
}

// OwnerCutPerMillion is a free data retrieval call binding the contract method 0xa01f79d4.
//
// Solidity: function ownerCutPerMillion() view returns(uint256)
func (_AuctionMarketplace *AuctionMarketplaceCaller) OwnerCutPerMillion(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AuctionMarketplace.contract.Call(opts, &out, "ownerCutPerMillion")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// OwnerCutPerMillion is a free data retrieval call binding the contract method 0xa01f79d4.
//
// Solidity: function ownerCutPerMillion() view returns(uint256)
func (_AuctionMarketplace *AuctionMarketplaceSession) OwnerCutPerMillion() (*big.Int, error) {
	return _AuctionMarketplace.Contract.OwnerCutPerMillion(&_AuctionMarketplace.CallOpts)
}

// OwnerCutPerMillion is a free data retrieval call binding the contract method 0xa01f79d4.
//
// Solidity: function ownerCutPerMillion() view returns(uint256)
func (_AuctionMarketplace *AuctionMarketplaceCallerSession) OwnerCutPerMillion() (*big.Int, error) {
	return _AuctionMarketplace.Contract.OwnerCutPerMillion(&_AuctionMarketplace.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_AuctionMarketplace *AuctionMarketplaceCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _AuctionMarketplace.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_AuctionMarketplace *AuctionMarketplaceSession) Paused() (bool, error) {
	return _AuctionMarketplace.Contract.Paused(&_AuctionMarketplace.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_AuctionMarketplace *AuctionMarketplaceCallerSession) Paused() (bool, error) {
	return _AuctionMarketplace.Contract.Paused(&_AuctionMarketplace.CallOpts)
}

// PendingReturns is a free data retrieval call binding the contract method 0xfa07aa2b.
//
// Solidity: function pendingReturns(bytes32 auctionId) view returns(uint256)
func (_AuctionMarketplace *AuctionMarketplaceCaller) PendingReturns(opts *bind.CallOpts, auctionId [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _AuctionMarketplace.contract.Call(opts, &out, "pendingReturns", auctionId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PendingReturns is a free data retrieval call binding the contract method 0xfa07aa2b.
//
// Solidity: function pendingReturns(bytes32 auctionId) view returns(uint256)
func (_AuctionMarketplace *AuctionMarketplaceSession) PendingReturns(auctionId [32]byte) (*big.Int, error) {
	return _AuctionMarketplace.Contract.PendingReturns(&_AuctionMarketplace.CallOpts, auctionId)
}

// PendingReturns is a free data retrieval call binding the contract method 0xfa07aa2b.
//
// Solidity: function pendingReturns(bytes32 auctionId) view returns(uint256)
func (_AuctionMarketplace *AuctionMarketplaceCallerSession) PendingReturns(auctionId [32]byte) (*big.Int, error) {
	return _AuctionMarketplace.Contract.PendingReturns(&_AuctionMarketplace.CallOpts, auctionId)
}

// PublicationFeeInWei is a free data retrieval call binding the contract method 0xae4f1198.
//
// Solidity: function publicationFeeInWei() view returns(uint256)
func (_AuctionMarketplace *AuctionMarketplaceCaller) PublicationFeeInWei(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AuctionMarketplace.contract.Call(opts, &out, "publicationFeeInWei")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PublicationFeeInWei is a free data retrieval call binding the contract method 0xae4f1198.
//
// Solidity: function publicationFeeInWei() view returns(uint256)
func (_AuctionMarketplace *AuctionMarketplaceSession) PublicationFeeInWei() (*big.Int, error) {
	return _AuctionMarketplace.Contract.PublicationFeeInWei(&_AuctionMarketplace.CallOpts)
}

// PublicationFeeInWei is a free data retrieval call binding the contract method 0xae4f1198.
//
// Solidity: function publicationFeeInWei() view returns(uint256)
func (_AuctionMarketplace *AuctionMarketplaceCallerSession) PublicationFeeInWei() (*big.Int, error) {
	return _AuctionMarketplace.Contract.PublicationFeeInWei(&_AuctionMarketplace.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_AuctionMarketplace *AuctionMarketplaceCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _AuctionMarketplace.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_AuctionMarketplace *AuctionMarketplaceSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _AuctionMarketplace.Contract.SupportsInterface(&_AuctionMarketplace.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_AuctionMarketplace *AuctionMarketplaceCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _AuctionMarketplace.Contract.SupportsInterface(&_AuctionMarketplace.CallOpts, interfaceId)
}

// AddBlackList is a paid mutator transaction binding the contract method 0x0ecb93c0.
//
// Solidity: function addBlackList(address _evilUser) returns()
func (_AuctionMarketplace *AuctionMarketplaceTransactor) AddBlackList(opts *bind.TransactOpts, _evilUser common.Address) (*types.Transaction, error) {
	return _AuctionMarketplace.contract.Transact(opts, "addBlackList", _evilUser)
}

// AddBlackList is a paid mutator transaction binding the contract method 0x0ecb93c0.
//
// Solidity: function addBlackList(address _evilUser) returns()
func (_AuctionMarketplace *AuctionMarketplaceSession) AddBlackList(_evilUser common.Address) (*types.Transaction, error) {
	return _AuctionMarketplace.Contract.AddBlackList(&_AuctionMarketplace.TransactOpts, _evilUser)
}

// AddBlackList is a paid mutator transaction binding the contract method 0x0ecb93c0.
//
// Solidity: function addBlackList(address _evilUser) returns()
func (_AuctionMarketplace *AuctionMarketplaceTransactorSession) AddBlackList(_evilUser common.Address) (*types.Transaction, error) {
	return _AuctionMarketplace.Contract.AddBlackList(&_AuctionMarketplace.TransactOpts, _evilUser)
}

// BidAuction is a paid mutator transaction binding the contract method 0x87b468ae.
//
// Solidity: function bidAuction(bytes32 nftAsset, bytes32 auctionId, bytes32 blindedBid, uint256 deposit) returns()
func (_AuctionMarketplace *AuctionMarketplaceTransactor) BidAuction(opts *bind.TransactOpts, nftAsset [32]byte, auctionId [32]byte, blindedBid [32]byte, deposit *big.Int) (*types.Transaction, error) {
	return _AuctionMarketplace.contract.Transact(opts, "bidAuction", nftAsset, auctionId, blindedBid, deposit)
}

// BidAuction is a paid mutator transaction binding the contract method 0x87b468ae.
//
// Solidity: function bidAuction(bytes32 nftAsset, bytes32 auctionId, bytes32 blindedBid, uint256 deposit) returns()
func (_AuctionMarketplace *AuctionMarketplaceSession) BidAuction(nftAsset [32]byte, auctionId [32]byte, blindedBid [32]byte, deposit *big.Int) (*types.Transaction, error) {
	return _AuctionMarketplace.Contract.BidAuction(&_AuctionMarketplace.TransactOpts, nftAsset, auctionId, blindedBid, deposit)
}

// BidAuction is a paid mutator transaction binding the contract method 0x87b468ae.
//
// Solidity: function bidAuction(bytes32 nftAsset, bytes32 auctionId, bytes32 blindedBid, uint256 deposit) returns()
func (_AuctionMarketplace *AuctionMarketplaceTransactorSession) BidAuction(nftAsset [32]byte, auctionId [32]byte, blindedBid [32]byte, deposit *big.Int) (*types.Transaction, error) {
	return _AuctionMarketplace.Contract.BidAuction(&_AuctionMarketplace.TransactOpts, nftAsset, auctionId, blindedBid, deposit)
}

// CancelAuction is a paid mutator transaction binding the contract method 0xdaa3d985.
//
// Solidity: function cancelAuction(bytes32 nftAsset, bytes32 auctionId) returns()
func (_AuctionMarketplace *AuctionMarketplaceTransactor) CancelAuction(opts *bind.TransactOpts, nftAsset [32]byte, auctionId [32]byte) (*types.Transaction, error) {
	return _AuctionMarketplace.contract.Transact(opts, "cancelAuction", nftAsset, auctionId)
}

// CancelAuction is a paid mutator transaction binding the contract method 0xdaa3d985.
//
// Solidity: function cancelAuction(bytes32 nftAsset, bytes32 auctionId) returns()
func (_AuctionMarketplace *AuctionMarketplaceSession) CancelAuction(nftAsset [32]byte, auctionId [32]byte) (*types.Transaction, error) {
	return _AuctionMarketplace.Contract.CancelAuction(&_AuctionMarketplace.TransactOpts, nftAsset, auctionId)
}

// CancelAuction is a paid mutator transaction binding the contract method 0xdaa3d985.
//
// Solidity: function cancelAuction(bytes32 nftAsset, bytes32 auctionId) returns()
func (_AuctionMarketplace *AuctionMarketplaceTransactorSession) CancelAuction(nftAsset [32]byte, auctionId [32]byte) (*types.Transaction, error) {
	return _AuctionMarketplace.Contract.CancelAuction(&_AuctionMarketplace.TransactOpts, nftAsset, auctionId)
}

// CloseAuction is a paid mutator transaction binding the contract method 0xd6710e23.
//
// Solidity: function closeAuction(address nftAddress, uint256 assetId, bytes32 auctionId) returns()
func (_AuctionMarketplace *AuctionMarketplaceTransactor) CloseAuction(opts *bind.TransactOpts, nftAddress common.Address, assetId *big.Int, auctionId [32]byte) (*types.Transaction, error) {
	return _AuctionMarketplace.contract.Transact(opts, "closeAuction", nftAddress, assetId, auctionId)
}

// CloseAuction is a paid mutator transaction binding the contract method 0xd6710e23.
//
// Solidity: function closeAuction(address nftAddress, uint256 assetId, bytes32 auctionId) returns()
func (_AuctionMarketplace *AuctionMarketplaceSession) CloseAuction(nftAddress common.Address, assetId *big.Int, auctionId [32]byte) (*types.Transaction, error) {
	return _AuctionMarketplace.Contract.CloseAuction(&_AuctionMarketplace.TransactOpts, nftAddress, assetId, auctionId)
}

// CloseAuction is a paid mutator transaction binding the contract method 0xd6710e23.
//
// Solidity: function closeAuction(address nftAddress, uint256 assetId, bytes32 auctionId) returns()
func (_AuctionMarketplace *AuctionMarketplaceTransactorSession) CloseAuction(nftAddress common.Address, assetId *big.Int, auctionId [32]byte) (*types.Transaction, error) {
	return _AuctionMarketplace.Contract.CloseAuction(&_AuctionMarketplace.TransactOpts, nftAddress, assetId, auctionId)
}

// CreateAuction is a paid mutator transaction binding the contract method 0x961c9ae4.
//
// Solidity: function createAuction(address nftAddress, uint256 assetId, uint256 startPriceInWei, uint256 biddingEnd, uint256 revealEnd) returns()
func (_AuctionMarketplace *AuctionMarketplaceTransactor) CreateAuction(opts *bind.TransactOpts, nftAddress common.Address, assetId *big.Int, startPriceInWei *big.Int, biddingEnd *big.Int, revealEnd *big.Int) (*types.Transaction, error) {
	return _AuctionMarketplace.contract.Transact(opts, "createAuction", nftAddress, assetId, startPriceInWei, biddingEnd, revealEnd)
}

// CreateAuction is a paid mutator transaction binding the contract method 0x961c9ae4.
//
// Solidity: function createAuction(address nftAddress, uint256 assetId, uint256 startPriceInWei, uint256 biddingEnd, uint256 revealEnd) returns()
func (_AuctionMarketplace *AuctionMarketplaceSession) CreateAuction(nftAddress common.Address, assetId *big.Int, startPriceInWei *big.Int, biddingEnd *big.Int, revealEnd *big.Int) (*types.Transaction, error) {
	return _AuctionMarketplace.Contract.CreateAuction(&_AuctionMarketplace.TransactOpts, nftAddress, assetId, startPriceInWei, biddingEnd, revealEnd)
}

// CreateAuction is a paid mutator transaction binding the contract method 0x961c9ae4.
//
// Solidity: function createAuction(address nftAddress, uint256 assetId, uint256 startPriceInWei, uint256 biddingEnd, uint256 revealEnd) returns()
func (_AuctionMarketplace *AuctionMarketplaceTransactorSession) CreateAuction(nftAddress common.Address, assetId *big.Int, startPriceInWei *big.Int, biddingEnd *big.Int, revealEnd *big.Int) (*types.Transaction, error) {
	return _AuctionMarketplace.Contract.CreateAuction(&_AuctionMarketplace.TransactOpts, nftAddress, assetId, startPriceInWei, biddingEnd, revealEnd)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_AuctionMarketplace *AuctionMarketplaceTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _AuctionMarketplace.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_AuctionMarketplace *AuctionMarketplaceSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _AuctionMarketplace.Contract.GrantRole(&_AuctionMarketplace.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_AuctionMarketplace *AuctionMarketplaceTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _AuctionMarketplace.Contract.GrantRole(&_AuctionMarketplace.TransactOpts, role, account)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_AuctionMarketplace *AuctionMarketplaceTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AuctionMarketplace.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_AuctionMarketplace *AuctionMarketplaceSession) Pause() (*types.Transaction, error) {
	return _AuctionMarketplace.Contract.Pause(&_AuctionMarketplace.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_AuctionMarketplace *AuctionMarketplaceTransactorSession) Pause() (*types.Transaction, error) {
	return _AuctionMarketplace.Contract.Pause(&_AuctionMarketplace.TransactOpts)
}

// RemoveBlackList is a paid mutator transaction binding the contract method 0xe4997dc5.
//
// Solidity: function removeBlackList(address _clearedUser) returns()
func (_AuctionMarketplace *AuctionMarketplaceTransactor) RemoveBlackList(opts *bind.TransactOpts, _clearedUser common.Address) (*types.Transaction, error) {
	return _AuctionMarketplace.contract.Transact(opts, "removeBlackList", _clearedUser)
}

// RemoveBlackList is a paid mutator transaction binding the contract method 0xe4997dc5.
//
// Solidity: function removeBlackList(address _clearedUser) returns()
func (_AuctionMarketplace *AuctionMarketplaceSession) RemoveBlackList(_clearedUser common.Address) (*types.Transaction, error) {
	return _AuctionMarketplace.Contract.RemoveBlackList(&_AuctionMarketplace.TransactOpts, _clearedUser)
}

// RemoveBlackList is a paid mutator transaction binding the contract method 0xe4997dc5.
//
// Solidity: function removeBlackList(address _clearedUser) returns()
func (_AuctionMarketplace *AuctionMarketplaceTransactorSession) RemoveBlackList(_clearedUser common.Address) (*types.Transaction, error) {
	return _AuctionMarketplace.Contract.RemoveBlackList(&_AuctionMarketplace.TransactOpts, _clearedUser)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_AuctionMarketplace *AuctionMarketplaceTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AuctionMarketplace.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_AuctionMarketplace *AuctionMarketplaceSession) RenounceOwnership() (*types.Transaction, error) {
	return _AuctionMarketplace.Contract.RenounceOwnership(&_AuctionMarketplace.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_AuctionMarketplace *AuctionMarketplaceTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _AuctionMarketplace.Contract.RenounceOwnership(&_AuctionMarketplace.TransactOpts)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_AuctionMarketplace *AuctionMarketplaceTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _AuctionMarketplace.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_AuctionMarketplace *AuctionMarketplaceSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _AuctionMarketplace.Contract.RenounceRole(&_AuctionMarketplace.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_AuctionMarketplace *AuctionMarketplaceTransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _AuctionMarketplace.Contract.RenounceRole(&_AuctionMarketplace.TransactOpts, role, account)
}

// RevealBid is a paid mutator transaction binding the contract method 0x85ec96a5.
//
// Solidity: function revealBid(bytes32 nftAsset, bytes32 auctionId, uint256[] _values, bool[] _fake, bytes32[] _secret) returns()
func (_AuctionMarketplace *AuctionMarketplaceTransactor) RevealBid(opts *bind.TransactOpts, nftAsset [32]byte, auctionId [32]byte, _values []*big.Int, _fake []bool, _secret [][32]byte) (*types.Transaction, error) {
	return _AuctionMarketplace.contract.Transact(opts, "revealBid", nftAsset, auctionId, _values, _fake, _secret)
}

// RevealBid is a paid mutator transaction binding the contract method 0x85ec96a5.
//
// Solidity: function revealBid(bytes32 nftAsset, bytes32 auctionId, uint256[] _values, bool[] _fake, bytes32[] _secret) returns()
func (_AuctionMarketplace *AuctionMarketplaceSession) RevealBid(nftAsset [32]byte, auctionId [32]byte, _values []*big.Int, _fake []bool, _secret [][32]byte) (*types.Transaction, error) {
	return _AuctionMarketplace.Contract.RevealBid(&_AuctionMarketplace.TransactOpts, nftAsset, auctionId, _values, _fake, _secret)
}

// RevealBid is a paid mutator transaction binding the contract method 0x85ec96a5.
//
// Solidity: function revealBid(bytes32 nftAsset, bytes32 auctionId, uint256[] _values, bool[] _fake, bytes32[] _secret) returns()
func (_AuctionMarketplace *AuctionMarketplaceTransactorSession) RevealBid(nftAsset [32]byte, auctionId [32]byte, _values []*big.Int, _fake []bool, _secret [][32]byte) (*types.Transaction, error) {
	return _AuctionMarketplace.Contract.RevealBid(&_AuctionMarketplace.TransactOpts, nftAsset, auctionId, _values, _fake, _secret)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_AuctionMarketplace *AuctionMarketplaceTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _AuctionMarketplace.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_AuctionMarketplace *AuctionMarketplaceSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _AuctionMarketplace.Contract.RevokeRole(&_AuctionMarketplace.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_AuctionMarketplace *AuctionMarketplaceTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _AuctionMarketplace.Contract.RevokeRole(&_AuctionMarketplace.TransactOpts, role, account)
}

// SetBeneficary is a paid mutator transaction binding the contract method 0x0db3cc05.
//
// Solidity: function setBeneficary(address _beneficary) returns()
func (_AuctionMarketplace *AuctionMarketplaceTransactor) SetBeneficary(opts *bind.TransactOpts, _beneficary common.Address) (*types.Transaction, error) {
	return _AuctionMarketplace.contract.Transact(opts, "setBeneficary", _beneficary)
}

// SetBeneficary is a paid mutator transaction binding the contract method 0x0db3cc05.
//
// Solidity: function setBeneficary(address _beneficary) returns()
func (_AuctionMarketplace *AuctionMarketplaceSession) SetBeneficary(_beneficary common.Address) (*types.Transaction, error) {
	return _AuctionMarketplace.Contract.SetBeneficary(&_AuctionMarketplace.TransactOpts, _beneficary)
}

// SetBeneficary is a paid mutator transaction binding the contract method 0x0db3cc05.
//
// Solidity: function setBeneficary(address _beneficary) returns()
func (_AuctionMarketplace *AuctionMarketplaceTransactorSession) SetBeneficary(_beneficary common.Address) (*types.Transaction, error) {
	return _AuctionMarketplace.Contract.SetBeneficary(&_AuctionMarketplace.TransactOpts, _beneficary)
}

// SetMinStageDuration is a paid mutator transaction binding the contract method 0xf6021884.
//
// Solidity: function setMinStageDuration(uint256 duration) returns()
func (_AuctionMarketplace *AuctionMarketplaceTransactor) SetMinStageDuration(opts *bind.TransactOpts, duration *big.Int) (*types.Transaction, error) {
	return _AuctionMarketplace.contract.Transact(opts, "setMinStageDuration", duration)
}

// SetMinStageDuration is a paid mutator transaction binding the contract method 0xf6021884.
//
// Solidity: function setMinStageDuration(uint256 duration) returns()
func (_AuctionMarketplace *AuctionMarketplaceSession) SetMinStageDuration(duration *big.Int) (*types.Transaction, error) {
	return _AuctionMarketplace.Contract.SetMinStageDuration(&_AuctionMarketplace.TransactOpts, duration)
}

// SetMinStageDuration is a paid mutator transaction binding the contract method 0xf6021884.
//
// Solidity: function setMinStageDuration(uint256 duration) returns()
func (_AuctionMarketplace *AuctionMarketplaceTransactorSession) SetMinStageDuration(duration *big.Int) (*types.Transaction, error) {
	return _AuctionMarketplace.Contract.SetMinStageDuration(&_AuctionMarketplace.TransactOpts, duration)
}

// SetOwnerCutPerMillion is a paid mutator transaction binding the contract method 0x19dad16d.
//
// Solidity: function setOwnerCutPerMillion(uint256 _ownerCutPerMillion) returns()
func (_AuctionMarketplace *AuctionMarketplaceTransactor) SetOwnerCutPerMillion(opts *bind.TransactOpts, _ownerCutPerMillion *big.Int) (*types.Transaction, error) {
	return _AuctionMarketplace.contract.Transact(opts, "setOwnerCutPerMillion", _ownerCutPerMillion)
}

// SetOwnerCutPerMillion is a paid mutator transaction binding the contract method 0x19dad16d.
//
// Solidity: function setOwnerCutPerMillion(uint256 _ownerCutPerMillion) returns()
func (_AuctionMarketplace *AuctionMarketplaceSession) SetOwnerCutPerMillion(_ownerCutPerMillion *big.Int) (*types.Transaction, error) {
	return _AuctionMarketplace.Contract.SetOwnerCutPerMillion(&_AuctionMarketplace.TransactOpts, _ownerCutPerMillion)
}

// SetOwnerCutPerMillion is a paid mutator transaction binding the contract method 0x19dad16d.
//
// Solidity: function setOwnerCutPerMillion(uint256 _ownerCutPerMillion) returns()
func (_AuctionMarketplace *AuctionMarketplaceTransactorSession) SetOwnerCutPerMillion(_ownerCutPerMillion *big.Int) (*types.Transaction, error) {
	return _AuctionMarketplace.Contract.SetOwnerCutPerMillion(&_AuctionMarketplace.TransactOpts, _ownerCutPerMillion)
}

// SetPublicationFee is a paid mutator transaction binding the contract method 0xaf8996f1.
//
// Solidity: function setPublicationFee(uint256 _publicationFee) returns()
func (_AuctionMarketplace *AuctionMarketplaceTransactor) SetPublicationFee(opts *bind.TransactOpts, _publicationFee *big.Int) (*types.Transaction, error) {
	return _AuctionMarketplace.contract.Transact(opts, "setPublicationFee", _publicationFee)
}

// SetPublicationFee is a paid mutator transaction binding the contract method 0xaf8996f1.
//
// Solidity: function setPublicationFee(uint256 _publicationFee) returns()
func (_AuctionMarketplace *AuctionMarketplaceSession) SetPublicationFee(_publicationFee *big.Int) (*types.Transaction, error) {
	return _AuctionMarketplace.Contract.SetPublicationFee(&_AuctionMarketplace.TransactOpts, _publicationFee)
}

// SetPublicationFee is a paid mutator transaction binding the contract method 0xaf8996f1.
//
// Solidity: function setPublicationFee(uint256 _publicationFee) returns()
func (_AuctionMarketplace *AuctionMarketplaceTransactorSession) SetPublicationFee(_publicationFee *big.Int) (*types.Transaction, error) {
	return _AuctionMarketplace.Contract.SetPublicationFee(&_AuctionMarketplace.TransactOpts, _publicationFee)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_AuctionMarketplace *AuctionMarketplaceTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _AuctionMarketplace.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_AuctionMarketplace *AuctionMarketplaceSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _AuctionMarketplace.Contract.TransferOwnership(&_AuctionMarketplace.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_AuctionMarketplace *AuctionMarketplaceTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _AuctionMarketplace.Contract.TransferOwnership(&_AuctionMarketplace.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_AuctionMarketplace *AuctionMarketplaceTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AuctionMarketplace.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_AuctionMarketplace *AuctionMarketplaceSession) Unpause() (*types.Transaction, error) {
	return _AuctionMarketplace.Contract.Unpause(&_AuctionMarketplace.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_AuctionMarketplace *AuctionMarketplaceTransactorSession) Unpause() (*types.Transaction, error) {
	return _AuctionMarketplace.Contract.Unpause(&_AuctionMarketplace.TransactOpts)
}

// UpdateStorageAddress is a paid mutator transaction binding the contract method 0x889e2129.
//
// Solidity: function updateStorageAddress(address _marketplaceStorage) returns()
func (_AuctionMarketplace *AuctionMarketplaceTransactor) UpdateStorageAddress(opts *bind.TransactOpts, _marketplaceStorage common.Address) (*types.Transaction, error) {
	return _AuctionMarketplace.contract.Transact(opts, "updateStorageAddress", _marketplaceStorage)
}

// UpdateStorageAddress is a paid mutator transaction binding the contract method 0x889e2129.
//
// Solidity: function updateStorageAddress(address _marketplaceStorage) returns()
func (_AuctionMarketplace *AuctionMarketplaceSession) UpdateStorageAddress(_marketplaceStorage common.Address) (*types.Transaction, error) {
	return _AuctionMarketplace.Contract.UpdateStorageAddress(&_AuctionMarketplace.TransactOpts, _marketplaceStorage)
}

// UpdateStorageAddress is a paid mutator transaction binding the contract method 0x889e2129.
//
// Solidity: function updateStorageAddress(address _marketplaceStorage) returns()
func (_AuctionMarketplace *AuctionMarketplaceTransactorSession) UpdateStorageAddress(_marketplaceStorage common.Address) (*types.Transaction, error) {
	return _AuctionMarketplace.Contract.UpdateStorageAddress(&_AuctionMarketplace.TransactOpts, _marketplaceStorage)
}

// Withdraw is a paid mutator transaction binding the contract method 0x8e19899e.
//
// Solidity: function withdraw(bytes32 auctionId) returns()
func (_AuctionMarketplace *AuctionMarketplaceTransactor) Withdraw(opts *bind.TransactOpts, auctionId [32]byte) (*types.Transaction, error) {
	return _AuctionMarketplace.contract.Transact(opts, "withdraw", auctionId)
}

// Withdraw is a paid mutator transaction binding the contract method 0x8e19899e.
//
// Solidity: function withdraw(bytes32 auctionId) returns()
func (_AuctionMarketplace *AuctionMarketplaceSession) Withdraw(auctionId [32]byte) (*types.Transaction, error) {
	return _AuctionMarketplace.Contract.Withdraw(&_AuctionMarketplace.TransactOpts, auctionId)
}

// Withdraw is a paid mutator transaction binding the contract method 0x8e19899e.
//
// Solidity: function withdraw(bytes32 auctionId) returns()
func (_AuctionMarketplace *AuctionMarketplaceTransactorSession) Withdraw(auctionId [32]byte) (*types.Transaction, error) {
	return _AuctionMarketplace.Contract.Withdraw(&_AuctionMarketplace.TransactOpts, auctionId)
}

// AuctionMarketplaceAddedBlackListIterator is returned from FilterAddedBlackList and is used to iterate over the raw logs and unpacked data for AddedBlackList events raised by the AuctionMarketplace contract.
type AuctionMarketplaceAddedBlackListIterator struct {
	Event *AuctionMarketplaceAddedBlackList // Event containing the contract specifics and raw log

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
func (it *AuctionMarketplaceAddedBlackListIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AuctionMarketplaceAddedBlackList)
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
		it.Event = new(AuctionMarketplaceAddedBlackList)
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
func (it *AuctionMarketplaceAddedBlackListIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AuctionMarketplaceAddedBlackListIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AuctionMarketplaceAddedBlackList represents a AddedBlackList event raised by the AuctionMarketplace contract.
type AuctionMarketplaceAddedBlackList struct {
	User common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterAddedBlackList is a free log retrieval operation binding the contract event 0x42e160154868087d6bfdc0ca23d96a1c1cfa32f1b72ba9ba27b69b98a0d819dc.
//
// Solidity: event AddedBlackList(address _user)
func (_AuctionMarketplace *AuctionMarketplaceFilterer) FilterAddedBlackList(opts *bind.FilterOpts) (*AuctionMarketplaceAddedBlackListIterator, error) {

	logs, sub, err := _AuctionMarketplace.contract.FilterLogs(opts, "AddedBlackList")
	if err != nil {
		return nil, err
	}
	return &AuctionMarketplaceAddedBlackListIterator{contract: _AuctionMarketplace.contract, event: "AddedBlackList", logs: logs, sub: sub}, nil
}

// WatchAddedBlackList is a free log subscription operation binding the contract event 0x42e160154868087d6bfdc0ca23d96a1c1cfa32f1b72ba9ba27b69b98a0d819dc.
//
// Solidity: event AddedBlackList(address _user)
func (_AuctionMarketplace *AuctionMarketplaceFilterer) WatchAddedBlackList(opts *bind.WatchOpts, sink chan<- *AuctionMarketplaceAddedBlackList) (event.Subscription, error) {

	logs, sub, err := _AuctionMarketplace.contract.WatchLogs(opts, "AddedBlackList")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AuctionMarketplaceAddedBlackList)
				if err := _AuctionMarketplace.contract.UnpackLog(event, "AddedBlackList", log); err != nil {
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
func (_AuctionMarketplace *AuctionMarketplaceFilterer) ParseAddedBlackList(log types.Log) (*AuctionMarketplaceAddedBlackList, error) {
	event := new(AuctionMarketplaceAddedBlackList)
	if err := _AuctionMarketplace.contract.UnpackLog(event, "AddedBlackList", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AuctionMarketplaceAuctionCancelledIterator is returned from FilterAuctionCancelled and is used to iterate over the raw logs and unpacked data for AuctionCancelled events raised by the AuctionMarketplace contract.
type AuctionMarketplaceAuctionCancelledIterator struct {
	Event *AuctionMarketplaceAuctionCancelled // Event containing the contract specifics and raw log

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
func (it *AuctionMarketplaceAuctionCancelledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AuctionMarketplaceAuctionCancelled)
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
		it.Event = new(AuctionMarketplaceAuctionCancelled)
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
func (it *AuctionMarketplaceAuctionCancelledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AuctionMarketplaceAuctionCancelledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AuctionMarketplaceAuctionCancelled represents a AuctionCancelled event raised by the AuctionMarketplace contract.
type AuctionMarketplaceAuctionCancelled struct {
	AuctionId [32]byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterAuctionCancelled is a free log retrieval operation binding the contract event 0xd88fab4b08bf76f15cf2d6e03e382acf1edd6790ab82967e406abac37db20288.
//
// Solidity: event AuctionCancelled(bytes32 auctionId)
func (_AuctionMarketplace *AuctionMarketplaceFilterer) FilterAuctionCancelled(opts *bind.FilterOpts) (*AuctionMarketplaceAuctionCancelledIterator, error) {

	logs, sub, err := _AuctionMarketplace.contract.FilterLogs(opts, "AuctionCancelled")
	if err != nil {
		return nil, err
	}
	return &AuctionMarketplaceAuctionCancelledIterator{contract: _AuctionMarketplace.contract, event: "AuctionCancelled", logs: logs, sub: sub}, nil
}

// WatchAuctionCancelled is a free log subscription operation binding the contract event 0xd88fab4b08bf76f15cf2d6e03e382acf1edd6790ab82967e406abac37db20288.
//
// Solidity: event AuctionCancelled(bytes32 auctionId)
func (_AuctionMarketplace *AuctionMarketplaceFilterer) WatchAuctionCancelled(opts *bind.WatchOpts, sink chan<- *AuctionMarketplaceAuctionCancelled) (event.Subscription, error) {

	logs, sub, err := _AuctionMarketplace.contract.WatchLogs(opts, "AuctionCancelled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AuctionMarketplaceAuctionCancelled)
				if err := _AuctionMarketplace.contract.UnpackLog(event, "AuctionCancelled", log); err != nil {
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
// Solidity: event AuctionCancelled(bytes32 auctionId)
func (_AuctionMarketplace *AuctionMarketplaceFilterer) ParseAuctionCancelled(log types.Log) (*AuctionMarketplaceAuctionCancelled, error) {
	event := new(AuctionMarketplaceAuctionCancelled)
	if err := _AuctionMarketplace.contract.UnpackLog(event, "AuctionCancelled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AuctionMarketplaceAuctionCreatedIterator is returned from FilterAuctionCreated and is used to iterate over the raw logs and unpacked data for AuctionCreated events raised by the AuctionMarketplace contract.
type AuctionMarketplaceAuctionCreatedIterator struct {
	Event *AuctionMarketplaceAuctionCreated // Event containing the contract specifics and raw log

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
func (it *AuctionMarketplaceAuctionCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AuctionMarketplaceAuctionCreated)
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
		it.Event = new(AuctionMarketplaceAuctionCreated)
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
func (it *AuctionMarketplaceAuctionCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AuctionMarketplaceAuctionCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AuctionMarketplaceAuctionCreated represents a AuctionCreated event raised by the AuctionMarketplace contract.
type AuctionMarketplaceAuctionCreated struct {
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
func (_AuctionMarketplace *AuctionMarketplaceFilterer) FilterAuctionCreated(opts *bind.FilterOpts) (*AuctionMarketplaceAuctionCreatedIterator, error) {

	logs, sub, err := _AuctionMarketplace.contract.FilterLogs(opts, "AuctionCreated")
	if err != nil {
		return nil, err
	}
	return &AuctionMarketplaceAuctionCreatedIterator{contract: _AuctionMarketplace.contract, event: "AuctionCreated", logs: logs, sub: sub}, nil
}

// WatchAuctionCreated is a free log subscription operation binding the contract event 0xaf4bad306f14d5c908e5b871fa54296ed7d4f29b8092bf5062de6813c247e54b.
//
// Solidity: event AuctionCreated(address seller, address nftAddress, bytes32 auctionId, uint256 assetId, uint256 biddingEnd, uint256 revealEnd, uint256 startPriceInWei)
func (_AuctionMarketplace *AuctionMarketplaceFilterer) WatchAuctionCreated(opts *bind.WatchOpts, sink chan<- *AuctionMarketplaceAuctionCreated) (event.Subscription, error) {

	logs, sub, err := _AuctionMarketplace.contract.WatchLogs(opts, "AuctionCreated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AuctionMarketplaceAuctionCreated)
				if err := _AuctionMarketplace.contract.UnpackLog(event, "AuctionCreated", log); err != nil {
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
func (_AuctionMarketplace *AuctionMarketplaceFilterer) ParseAuctionCreated(log types.Log) (*AuctionMarketplaceAuctionCreated, error) {
	event := new(AuctionMarketplaceAuctionCreated)
	if err := _AuctionMarketplace.contract.UnpackLog(event, "AuctionCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AuctionMarketplaceAuctionEndedIterator is returned from FilterAuctionEnded and is used to iterate over the raw logs and unpacked data for AuctionEnded events raised by the AuctionMarketplace contract.
type AuctionMarketplaceAuctionEndedIterator struct {
	Event *AuctionMarketplaceAuctionEnded // Event containing the contract specifics and raw log

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
func (it *AuctionMarketplaceAuctionEndedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AuctionMarketplaceAuctionEnded)
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
		it.Event = new(AuctionMarketplaceAuctionEnded)
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
func (it *AuctionMarketplaceAuctionEndedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AuctionMarketplaceAuctionEndedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AuctionMarketplaceAuctionEnded represents a AuctionEnded event raised by the AuctionMarketplace contract.
type AuctionMarketplaceAuctionEnded struct {
	AuctionId [32]byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterAuctionEnded is a free log retrieval operation binding the contract event 0x447d0298e6e2716a4343195759c35b277a60b76e628284d93262fa9869db9921.
//
// Solidity: event AuctionEnded(bytes32 auctionId)
func (_AuctionMarketplace *AuctionMarketplaceFilterer) FilterAuctionEnded(opts *bind.FilterOpts) (*AuctionMarketplaceAuctionEndedIterator, error) {

	logs, sub, err := _AuctionMarketplace.contract.FilterLogs(opts, "AuctionEnded")
	if err != nil {
		return nil, err
	}
	return &AuctionMarketplaceAuctionEndedIterator{contract: _AuctionMarketplace.contract, event: "AuctionEnded", logs: logs, sub: sub}, nil
}

// WatchAuctionEnded is a free log subscription operation binding the contract event 0x447d0298e6e2716a4343195759c35b277a60b76e628284d93262fa9869db9921.
//
// Solidity: event AuctionEnded(bytes32 auctionId)
func (_AuctionMarketplace *AuctionMarketplaceFilterer) WatchAuctionEnded(opts *bind.WatchOpts, sink chan<- *AuctionMarketplaceAuctionEnded) (event.Subscription, error) {

	logs, sub, err := _AuctionMarketplace.contract.WatchLogs(opts, "AuctionEnded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AuctionMarketplaceAuctionEnded)
				if err := _AuctionMarketplace.contract.UnpackLog(event, "AuctionEnded", log); err != nil {
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
func (_AuctionMarketplace *AuctionMarketplaceFilterer) ParseAuctionEnded(log types.Log) (*AuctionMarketplaceAuctionEnded, error) {
	event := new(AuctionMarketplaceAuctionEnded)
	if err := _AuctionMarketplace.contract.UnpackLog(event, "AuctionEnded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AuctionMarketplaceAuctionRefundIterator is returned from FilterAuctionRefund and is used to iterate over the raw logs and unpacked data for AuctionRefund events raised by the AuctionMarketplace contract.
type AuctionMarketplaceAuctionRefundIterator struct {
	Event *AuctionMarketplaceAuctionRefund // Event containing the contract specifics and raw log

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
func (it *AuctionMarketplaceAuctionRefundIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AuctionMarketplaceAuctionRefund)
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
		it.Event = new(AuctionMarketplaceAuctionRefund)
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
func (it *AuctionMarketplaceAuctionRefundIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AuctionMarketplaceAuctionRefundIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AuctionMarketplaceAuctionRefund represents a AuctionRefund event raised by the AuctionMarketplace contract.
type AuctionMarketplaceAuctionRefund struct {
	Bidder    common.Address
	AuctionId [32]byte
	Deposit   *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterAuctionRefund is a free log retrieval operation binding the contract event 0x6be22eb0241a0575f8b866ff176552cf71078977d6d30e501dd2d326c8d57114.
//
// Solidity: event AuctionRefund(address bidder, bytes32 auctionId, uint256 deposit)
func (_AuctionMarketplace *AuctionMarketplaceFilterer) FilterAuctionRefund(opts *bind.FilterOpts) (*AuctionMarketplaceAuctionRefundIterator, error) {

	logs, sub, err := _AuctionMarketplace.contract.FilterLogs(opts, "AuctionRefund")
	if err != nil {
		return nil, err
	}
	return &AuctionMarketplaceAuctionRefundIterator{contract: _AuctionMarketplace.contract, event: "AuctionRefund", logs: logs, sub: sub}, nil
}

// WatchAuctionRefund is a free log subscription operation binding the contract event 0x6be22eb0241a0575f8b866ff176552cf71078977d6d30e501dd2d326c8d57114.
//
// Solidity: event AuctionRefund(address bidder, bytes32 auctionId, uint256 deposit)
func (_AuctionMarketplace *AuctionMarketplaceFilterer) WatchAuctionRefund(opts *bind.WatchOpts, sink chan<- *AuctionMarketplaceAuctionRefund) (event.Subscription, error) {

	logs, sub, err := _AuctionMarketplace.contract.WatchLogs(opts, "AuctionRefund")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AuctionMarketplaceAuctionRefund)
				if err := _AuctionMarketplace.contract.UnpackLog(event, "AuctionRefund", log); err != nil {
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
func (_AuctionMarketplace *AuctionMarketplaceFilterer) ParseAuctionRefund(log types.Log) (*AuctionMarketplaceAuctionRefund, error) {
	event := new(AuctionMarketplaceAuctionRefund)
	if err := _AuctionMarketplace.contract.UnpackLog(event, "AuctionRefund", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AuctionMarketplaceAuctionSuccessfulIterator is returned from FilterAuctionSuccessful and is used to iterate over the raw logs and unpacked data for AuctionSuccessful events raised by the AuctionMarketplace contract.
type AuctionMarketplaceAuctionSuccessfulIterator struct {
	Event *AuctionMarketplaceAuctionSuccessful // Event containing the contract specifics and raw log

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
func (it *AuctionMarketplaceAuctionSuccessfulIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AuctionMarketplaceAuctionSuccessful)
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
		it.Event = new(AuctionMarketplaceAuctionSuccessful)
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
func (it *AuctionMarketplaceAuctionSuccessfulIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AuctionMarketplaceAuctionSuccessfulIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AuctionMarketplaceAuctionSuccessful represents a AuctionSuccessful event raised by the AuctionMarketplace contract.
type AuctionMarketplaceAuctionSuccessful struct {
	Seller     common.Address
	Buyer      common.Address
	AuctionId  [32]byte
	TotalPrice *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterAuctionSuccessful is a free log retrieval operation binding the contract event 0x5ed769e966210f5d06d68e75831871131255cee237cb8edf78565aae170475ea.
//
// Solidity: event AuctionSuccessful(address seller, address buyer, bytes32 auctionId, uint256 totalPrice)
func (_AuctionMarketplace *AuctionMarketplaceFilterer) FilterAuctionSuccessful(opts *bind.FilterOpts) (*AuctionMarketplaceAuctionSuccessfulIterator, error) {

	logs, sub, err := _AuctionMarketplace.contract.FilterLogs(opts, "AuctionSuccessful")
	if err != nil {
		return nil, err
	}
	return &AuctionMarketplaceAuctionSuccessfulIterator{contract: _AuctionMarketplace.contract, event: "AuctionSuccessful", logs: logs, sub: sub}, nil
}

// WatchAuctionSuccessful is a free log subscription operation binding the contract event 0x5ed769e966210f5d06d68e75831871131255cee237cb8edf78565aae170475ea.
//
// Solidity: event AuctionSuccessful(address seller, address buyer, bytes32 auctionId, uint256 totalPrice)
func (_AuctionMarketplace *AuctionMarketplaceFilterer) WatchAuctionSuccessful(opts *bind.WatchOpts, sink chan<- *AuctionMarketplaceAuctionSuccessful) (event.Subscription, error) {

	logs, sub, err := _AuctionMarketplace.contract.WatchLogs(opts, "AuctionSuccessful")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AuctionMarketplaceAuctionSuccessful)
				if err := _AuctionMarketplace.contract.UnpackLog(event, "AuctionSuccessful", log); err != nil {
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
func (_AuctionMarketplace *AuctionMarketplaceFilterer) ParseAuctionSuccessful(log types.Log) (*AuctionMarketplaceAuctionSuccessful, error) {
	event := new(AuctionMarketplaceAuctionSuccessful)
	if err := _AuctionMarketplace.contract.UnpackLog(event, "AuctionSuccessful", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AuctionMarketplaceBidSuccessfulIterator is returned from FilterBidSuccessful and is used to iterate over the raw logs and unpacked data for BidSuccessful events raised by the AuctionMarketplace contract.
type AuctionMarketplaceBidSuccessfulIterator struct {
	Event *AuctionMarketplaceBidSuccessful // Event containing the contract specifics and raw log

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
func (it *AuctionMarketplaceBidSuccessfulIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AuctionMarketplaceBidSuccessful)
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
		it.Event = new(AuctionMarketplaceBidSuccessful)
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
func (it *AuctionMarketplaceBidSuccessfulIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AuctionMarketplaceBidSuccessfulIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AuctionMarketplaceBidSuccessful represents a BidSuccessful event raised by the AuctionMarketplace contract.
type AuctionMarketplaceBidSuccessful struct {
	Bidder     common.Address
	AuctionId  [32]byte
	BlindedBid [32]byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterBidSuccessful is a free log retrieval operation binding the contract event 0x327e5a459c964bcfa4f6d71b32a04a53b0002417b0924a8057d9f12e688e0bec.
//
// Solidity: event BidSuccessful(address bidder, bytes32 auctionId, bytes32 blindedBid)
func (_AuctionMarketplace *AuctionMarketplaceFilterer) FilterBidSuccessful(opts *bind.FilterOpts) (*AuctionMarketplaceBidSuccessfulIterator, error) {

	logs, sub, err := _AuctionMarketplace.contract.FilterLogs(opts, "BidSuccessful")
	if err != nil {
		return nil, err
	}
	return &AuctionMarketplaceBidSuccessfulIterator{contract: _AuctionMarketplace.contract, event: "BidSuccessful", logs: logs, sub: sub}, nil
}

// WatchBidSuccessful is a free log subscription operation binding the contract event 0x327e5a459c964bcfa4f6d71b32a04a53b0002417b0924a8057d9f12e688e0bec.
//
// Solidity: event BidSuccessful(address bidder, bytes32 auctionId, bytes32 blindedBid)
func (_AuctionMarketplace *AuctionMarketplaceFilterer) WatchBidSuccessful(opts *bind.WatchOpts, sink chan<- *AuctionMarketplaceBidSuccessful) (event.Subscription, error) {

	logs, sub, err := _AuctionMarketplace.contract.WatchLogs(opts, "BidSuccessful")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AuctionMarketplaceBidSuccessful)
				if err := _AuctionMarketplace.contract.UnpackLog(event, "BidSuccessful", log); err != nil {
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
func (_AuctionMarketplace *AuctionMarketplaceFilterer) ParseBidSuccessful(log types.Log) (*AuctionMarketplaceBidSuccessful, error) {
	event := new(AuctionMarketplaceBidSuccessful)
	if err := _AuctionMarketplace.contract.UnpackLog(event, "BidSuccessful", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AuctionMarketplaceChangedOwnerCutPerMillionIterator is returned from FilterChangedOwnerCutPerMillion and is used to iterate over the raw logs and unpacked data for ChangedOwnerCutPerMillion events raised by the AuctionMarketplace contract.
type AuctionMarketplaceChangedOwnerCutPerMillionIterator struct {
	Event *AuctionMarketplaceChangedOwnerCutPerMillion // Event containing the contract specifics and raw log

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
func (it *AuctionMarketplaceChangedOwnerCutPerMillionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AuctionMarketplaceChangedOwnerCutPerMillion)
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
		it.Event = new(AuctionMarketplaceChangedOwnerCutPerMillion)
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
func (it *AuctionMarketplaceChangedOwnerCutPerMillionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AuctionMarketplaceChangedOwnerCutPerMillionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AuctionMarketplaceChangedOwnerCutPerMillion represents a ChangedOwnerCutPerMillion event raised by the AuctionMarketplace contract.
type AuctionMarketplaceChangedOwnerCutPerMillion struct {
	OwnerCutPerMillion *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterChangedOwnerCutPerMillion is a free log retrieval operation binding the contract event 0xfa406a120a9e7f2b332bfb7a43d3bf1c3f079262202907a6b69c94b2821a02c6.
//
// Solidity: event ChangedOwnerCutPerMillion(uint256 ownerCutPerMillion)
func (_AuctionMarketplace *AuctionMarketplaceFilterer) FilterChangedOwnerCutPerMillion(opts *bind.FilterOpts) (*AuctionMarketplaceChangedOwnerCutPerMillionIterator, error) {

	logs, sub, err := _AuctionMarketplace.contract.FilterLogs(opts, "ChangedOwnerCutPerMillion")
	if err != nil {
		return nil, err
	}
	return &AuctionMarketplaceChangedOwnerCutPerMillionIterator{contract: _AuctionMarketplace.contract, event: "ChangedOwnerCutPerMillion", logs: logs, sub: sub}, nil
}

// WatchChangedOwnerCutPerMillion is a free log subscription operation binding the contract event 0xfa406a120a9e7f2b332bfb7a43d3bf1c3f079262202907a6b69c94b2821a02c6.
//
// Solidity: event ChangedOwnerCutPerMillion(uint256 ownerCutPerMillion)
func (_AuctionMarketplace *AuctionMarketplaceFilterer) WatchChangedOwnerCutPerMillion(opts *bind.WatchOpts, sink chan<- *AuctionMarketplaceChangedOwnerCutPerMillion) (event.Subscription, error) {

	logs, sub, err := _AuctionMarketplace.contract.WatchLogs(opts, "ChangedOwnerCutPerMillion")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AuctionMarketplaceChangedOwnerCutPerMillion)
				if err := _AuctionMarketplace.contract.UnpackLog(event, "ChangedOwnerCutPerMillion", log); err != nil {
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

// ParseChangedOwnerCutPerMillion is a log parse operation binding the contract event 0xfa406a120a9e7f2b332bfb7a43d3bf1c3f079262202907a6b69c94b2821a02c6.
//
// Solidity: event ChangedOwnerCutPerMillion(uint256 ownerCutPerMillion)
func (_AuctionMarketplace *AuctionMarketplaceFilterer) ParseChangedOwnerCutPerMillion(log types.Log) (*AuctionMarketplaceChangedOwnerCutPerMillion, error) {
	event := new(AuctionMarketplaceChangedOwnerCutPerMillion)
	if err := _AuctionMarketplace.contract.UnpackLog(event, "ChangedOwnerCutPerMillion", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AuctionMarketplaceChangedPublicationFeeIterator is returned from FilterChangedPublicationFee and is used to iterate over the raw logs and unpacked data for ChangedPublicationFee events raised by the AuctionMarketplace contract.
type AuctionMarketplaceChangedPublicationFeeIterator struct {
	Event *AuctionMarketplaceChangedPublicationFee // Event containing the contract specifics and raw log

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
func (it *AuctionMarketplaceChangedPublicationFeeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AuctionMarketplaceChangedPublicationFee)
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
		it.Event = new(AuctionMarketplaceChangedPublicationFee)
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
func (it *AuctionMarketplaceChangedPublicationFeeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AuctionMarketplaceChangedPublicationFeeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AuctionMarketplaceChangedPublicationFee represents a ChangedPublicationFee event raised by the AuctionMarketplace contract.
type AuctionMarketplaceChangedPublicationFee struct {
	PublicationFee *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterChangedPublicationFee is a free log retrieval operation binding the contract event 0xe7fa8737293f41b5dfa0d5c3e552860a06275bed7015581b083c7be7003308ba.
//
// Solidity: event ChangedPublicationFee(uint256 publicationFee)
func (_AuctionMarketplace *AuctionMarketplaceFilterer) FilterChangedPublicationFee(opts *bind.FilterOpts) (*AuctionMarketplaceChangedPublicationFeeIterator, error) {

	logs, sub, err := _AuctionMarketplace.contract.FilterLogs(opts, "ChangedPublicationFee")
	if err != nil {
		return nil, err
	}
	return &AuctionMarketplaceChangedPublicationFeeIterator{contract: _AuctionMarketplace.contract, event: "ChangedPublicationFee", logs: logs, sub: sub}, nil
}

// WatchChangedPublicationFee is a free log subscription operation binding the contract event 0xe7fa8737293f41b5dfa0d5c3e552860a06275bed7015581b083c7be7003308ba.
//
// Solidity: event ChangedPublicationFee(uint256 publicationFee)
func (_AuctionMarketplace *AuctionMarketplaceFilterer) WatchChangedPublicationFee(opts *bind.WatchOpts, sink chan<- *AuctionMarketplaceChangedPublicationFee) (event.Subscription, error) {

	logs, sub, err := _AuctionMarketplace.contract.WatchLogs(opts, "ChangedPublicationFee")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AuctionMarketplaceChangedPublicationFee)
				if err := _AuctionMarketplace.contract.UnpackLog(event, "ChangedPublicationFee", log); err != nil {
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

// ParseChangedPublicationFee is a log parse operation binding the contract event 0xe7fa8737293f41b5dfa0d5c3e552860a06275bed7015581b083c7be7003308ba.
//
// Solidity: event ChangedPublicationFee(uint256 publicationFee)
func (_AuctionMarketplace *AuctionMarketplaceFilterer) ParseChangedPublicationFee(log types.Log) (*AuctionMarketplaceChangedPublicationFee, error) {
	event := new(AuctionMarketplaceChangedPublicationFee)
	if err := _AuctionMarketplace.contract.UnpackLog(event, "ChangedPublicationFee", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AuctionMarketplaceMarketplaceStorageUpdatedIterator is returned from FilterMarketplaceStorageUpdated and is used to iterate over the raw logs and unpacked data for MarketplaceStorageUpdated events raised by the AuctionMarketplace contract.
type AuctionMarketplaceMarketplaceStorageUpdatedIterator struct {
	Event *AuctionMarketplaceMarketplaceStorageUpdated // Event containing the contract specifics and raw log

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
func (it *AuctionMarketplaceMarketplaceStorageUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AuctionMarketplaceMarketplaceStorageUpdated)
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
		it.Event = new(AuctionMarketplaceMarketplaceStorageUpdated)
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
func (it *AuctionMarketplaceMarketplaceStorageUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AuctionMarketplaceMarketplaceStorageUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AuctionMarketplaceMarketplaceStorageUpdated represents a MarketplaceStorageUpdated event raised by the AuctionMarketplace contract.
type AuctionMarketplaceMarketplaceStorageUpdated struct {
	MarketplaceStorage common.Address
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterMarketplaceStorageUpdated is a free log retrieval operation binding the contract event 0x19e07a962d0e1cab8b7d4be06b77a7225ca7ebb744994cb6546fdd815590dc9d.
//
// Solidity: event MarketplaceStorageUpdated(address _marketplaceStorage)
func (_AuctionMarketplace *AuctionMarketplaceFilterer) FilterMarketplaceStorageUpdated(opts *bind.FilterOpts) (*AuctionMarketplaceMarketplaceStorageUpdatedIterator, error) {

	logs, sub, err := _AuctionMarketplace.contract.FilterLogs(opts, "MarketplaceStorageUpdated")
	if err != nil {
		return nil, err
	}
	return &AuctionMarketplaceMarketplaceStorageUpdatedIterator{contract: _AuctionMarketplace.contract, event: "MarketplaceStorageUpdated", logs: logs, sub: sub}, nil
}

// WatchMarketplaceStorageUpdated is a free log subscription operation binding the contract event 0x19e07a962d0e1cab8b7d4be06b77a7225ca7ebb744994cb6546fdd815590dc9d.
//
// Solidity: event MarketplaceStorageUpdated(address _marketplaceStorage)
func (_AuctionMarketplace *AuctionMarketplaceFilterer) WatchMarketplaceStorageUpdated(opts *bind.WatchOpts, sink chan<- *AuctionMarketplaceMarketplaceStorageUpdated) (event.Subscription, error) {

	logs, sub, err := _AuctionMarketplace.contract.WatchLogs(opts, "MarketplaceStorageUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AuctionMarketplaceMarketplaceStorageUpdated)
				if err := _AuctionMarketplace.contract.UnpackLog(event, "MarketplaceStorageUpdated", log); err != nil {
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

// ParseMarketplaceStorageUpdated is a log parse operation binding the contract event 0x19e07a962d0e1cab8b7d4be06b77a7225ca7ebb744994cb6546fdd815590dc9d.
//
// Solidity: event MarketplaceStorageUpdated(address _marketplaceStorage)
func (_AuctionMarketplace *AuctionMarketplaceFilterer) ParseMarketplaceStorageUpdated(log types.Log) (*AuctionMarketplaceMarketplaceStorageUpdated, error) {
	event := new(AuctionMarketplaceMarketplaceStorageUpdated)
	if err := _AuctionMarketplace.contract.UnpackLog(event, "MarketplaceStorageUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AuctionMarketplaceOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the AuctionMarketplace contract.
type AuctionMarketplaceOwnershipTransferredIterator struct {
	Event *AuctionMarketplaceOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *AuctionMarketplaceOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AuctionMarketplaceOwnershipTransferred)
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
		it.Event = new(AuctionMarketplaceOwnershipTransferred)
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
func (it *AuctionMarketplaceOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AuctionMarketplaceOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AuctionMarketplaceOwnershipTransferred represents a OwnershipTransferred event raised by the AuctionMarketplace contract.
type AuctionMarketplaceOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_AuctionMarketplace *AuctionMarketplaceFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*AuctionMarketplaceOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _AuctionMarketplace.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &AuctionMarketplaceOwnershipTransferredIterator{contract: _AuctionMarketplace.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_AuctionMarketplace *AuctionMarketplaceFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *AuctionMarketplaceOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _AuctionMarketplace.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AuctionMarketplaceOwnershipTransferred)
				if err := _AuctionMarketplace.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_AuctionMarketplace *AuctionMarketplaceFilterer) ParseOwnershipTransferred(log types.Log) (*AuctionMarketplaceOwnershipTransferred, error) {
	event := new(AuctionMarketplaceOwnershipTransferred)
	if err := _AuctionMarketplace.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AuctionMarketplacePausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the AuctionMarketplace contract.
type AuctionMarketplacePausedIterator struct {
	Event *AuctionMarketplacePaused // Event containing the contract specifics and raw log

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
func (it *AuctionMarketplacePausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AuctionMarketplacePaused)
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
		it.Event = new(AuctionMarketplacePaused)
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
func (it *AuctionMarketplacePausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AuctionMarketplacePausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AuctionMarketplacePaused represents a Paused event raised by the AuctionMarketplace contract.
type AuctionMarketplacePaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_AuctionMarketplace *AuctionMarketplaceFilterer) FilterPaused(opts *bind.FilterOpts) (*AuctionMarketplacePausedIterator, error) {

	logs, sub, err := _AuctionMarketplace.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &AuctionMarketplacePausedIterator{contract: _AuctionMarketplace.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_AuctionMarketplace *AuctionMarketplaceFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *AuctionMarketplacePaused) (event.Subscription, error) {

	logs, sub, err := _AuctionMarketplace.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AuctionMarketplacePaused)
				if err := _AuctionMarketplace.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_AuctionMarketplace *AuctionMarketplaceFilterer) ParsePaused(log types.Log) (*AuctionMarketplacePaused, error) {
	event := new(AuctionMarketplacePaused)
	if err := _AuctionMarketplace.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AuctionMarketplaceRemovedBlackListIterator is returned from FilterRemovedBlackList and is used to iterate over the raw logs and unpacked data for RemovedBlackList events raised by the AuctionMarketplace contract.
type AuctionMarketplaceRemovedBlackListIterator struct {
	Event *AuctionMarketplaceRemovedBlackList // Event containing the contract specifics and raw log

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
func (it *AuctionMarketplaceRemovedBlackListIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AuctionMarketplaceRemovedBlackList)
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
		it.Event = new(AuctionMarketplaceRemovedBlackList)
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
func (it *AuctionMarketplaceRemovedBlackListIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AuctionMarketplaceRemovedBlackListIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AuctionMarketplaceRemovedBlackList represents a RemovedBlackList event raised by the AuctionMarketplace contract.
type AuctionMarketplaceRemovedBlackList struct {
	User common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterRemovedBlackList is a free log retrieval operation binding the contract event 0xd7e9ec6e6ecd65492dce6bf513cd6867560d49544421d0783ddf06e76c24470c.
//
// Solidity: event RemovedBlackList(address _user)
func (_AuctionMarketplace *AuctionMarketplaceFilterer) FilterRemovedBlackList(opts *bind.FilterOpts) (*AuctionMarketplaceRemovedBlackListIterator, error) {

	logs, sub, err := _AuctionMarketplace.contract.FilterLogs(opts, "RemovedBlackList")
	if err != nil {
		return nil, err
	}
	return &AuctionMarketplaceRemovedBlackListIterator{contract: _AuctionMarketplace.contract, event: "RemovedBlackList", logs: logs, sub: sub}, nil
}

// WatchRemovedBlackList is a free log subscription operation binding the contract event 0xd7e9ec6e6ecd65492dce6bf513cd6867560d49544421d0783ddf06e76c24470c.
//
// Solidity: event RemovedBlackList(address _user)
func (_AuctionMarketplace *AuctionMarketplaceFilterer) WatchRemovedBlackList(opts *bind.WatchOpts, sink chan<- *AuctionMarketplaceRemovedBlackList) (event.Subscription, error) {

	logs, sub, err := _AuctionMarketplace.contract.WatchLogs(opts, "RemovedBlackList")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AuctionMarketplaceRemovedBlackList)
				if err := _AuctionMarketplace.contract.UnpackLog(event, "RemovedBlackList", log); err != nil {
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
func (_AuctionMarketplace *AuctionMarketplaceFilterer) ParseRemovedBlackList(log types.Log) (*AuctionMarketplaceRemovedBlackList, error) {
	event := new(AuctionMarketplaceRemovedBlackList)
	if err := _AuctionMarketplace.contract.UnpackLog(event, "RemovedBlackList", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AuctionMarketplaceRevealFailedIterator is returned from FilterRevealFailed and is used to iterate over the raw logs and unpacked data for RevealFailed events raised by the AuctionMarketplace contract.
type AuctionMarketplaceRevealFailedIterator struct {
	Event *AuctionMarketplaceRevealFailed // Event containing the contract specifics and raw log

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
func (it *AuctionMarketplaceRevealFailedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AuctionMarketplaceRevealFailed)
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
		it.Event = new(AuctionMarketplaceRevealFailed)
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
func (it *AuctionMarketplaceRevealFailedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AuctionMarketplaceRevealFailedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AuctionMarketplaceRevealFailed represents a RevealFailed event raised by the AuctionMarketplace contract.
type AuctionMarketplaceRevealFailed struct {
	Fake      bool
	Revealer  common.Address
	AuctionId [32]byte
	Value     *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterRevealFailed is a free log retrieval operation binding the contract event 0x255ef988657daaeb7f921a0c776fa6efb221748fb64e61ab1febb0fd37f67175.
//
// Solidity: event RevealFailed(bool fake, address revealer, bytes32 auctionId, uint256 value)
func (_AuctionMarketplace *AuctionMarketplaceFilterer) FilterRevealFailed(opts *bind.FilterOpts) (*AuctionMarketplaceRevealFailedIterator, error) {

	logs, sub, err := _AuctionMarketplace.contract.FilterLogs(opts, "RevealFailed")
	if err != nil {
		return nil, err
	}
	return &AuctionMarketplaceRevealFailedIterator{contract: _AuctionMarketplace.contract, event: "RevealFailed", logs: logs, sub: sub}, nil
}

// WatchRevealFailed is a free log subscription operation binding the contract event 0x255ef988657daaeb7f921a0c776fa6efb221748fb64e61ab1febb0fd37f67175.
//
// Solidity: event RevealFailed(bool fake, address revealer, bytes32 auctionId, uint256 value)
func (_AuctionMarketplace *AuctionMarketplaceFilterer) WatchRevealFailed(opts *bind.WatchOpts, sink chan<- *AuctionMarketplaceRevealFailed) (event.Subscription, error) {

	logs, sub, err := _AuctionMarketplace.contract.WatchLogs(opts, "RevealFailed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AuctionMarketplaceRevealFailed)
				if err := _AuctionMarketplace.contract.UnpackLog(event, "RevealFailed", log); err != nil {
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
func (_AuctionMarketplace *AuctionMarketplaceFilterer) ParseRevealFailed(log types.Log) (*AuctionMarketplaceRevealFailed, error) {
	event := new(AuctionMarketplaceRevealFailed)
	if err := _AuctionMarketplace.contract.UnpackLog(event, "RevealFailed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AuctionMarketplaceRevealSuccessfulIterator is returned from FilterRevealSuccessful and is used to iterate over the raw logs and unpacked data for RevealSuccessful events raised by the AuctionMarketplace contract.
type AuctionMarketplaceRevealSuccessfulIterator struct {
	Event *AuctionMarketplaceRevealSuccessful // Event containing the contract specifics and raw log

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
func (it *AuctionMarketplaceRevealSuccessfulIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AuctionMarketplaceRevealSuccessful)
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
		it.Event = new(AuctionMarketplaceRevealSuccessful)
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
func (it *AuctionMarketplaceRevealSuccessfulIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AuctionMarketplaceRevealSuccessfulIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AuctionMarketplaceRevealSuccessful represents a RevealSuccessful event raised by the AuctionMarketplace contract.
type AuctionMarketplaceRevealSuccessful struct {
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
func (_AuctionMarketplace *AuctionMarketplaceFilterer) FilterRevealSuccessful(opts *bind.FilterOpts) (*AuctionMarketplaceRevealSuccessfulIterator, error) {

	logs, sub, err := _AuctionMarketplace.contract.FilterLogs(opts, "RevealSuccessful")
	if err != nil {
		return nil, err
	}
	return &AuctionMarketplaceRevealSuccessfulIterator{contract: _AuctionMarketplace.contract, event: "RevealSuccessful", logs: logs, sub: sub}, nil
}

// WatchRevealSuccessful is a free log subscription operation binding the contract event 0xabef59dc3ae014d197fad42649c58d34bfc816d3e1a7f26ca32c13611b13e7a1.
//
// Solidity: event RevealSuccessful(bool fake, address revealer, bytes32 auctionId, uint256 value, bytes32 blindedBid)
func (_AuctionMarketplace *AuctionMarketplaceFilterer) WatchRevealSuccessful(opts *bind.WatchOpts, sink chan<- *AuctionMarketplaceRevealSuccessful) (event.Subscription, error) {

	logs, sub, err := _AuctionMarketplace.contract.WatchLogs(opts, "RevealSuccessful")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AuctionMarketplaceRevealSuccessful)
				if err := _AuctionMarketplace.contract.UnpackLog(event, "RevealSuccessful", log); err != nil {
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
func (_AuctionMarketplace *AuctionMarketplaceFilterer) ParseRevealSuccessful(log types.Log) (*AuctionMarketplaceRevealSuccessful, error) {
	event := new(AuctionMarketplaceRevealSuccessful)
	if err := _AuctionMarketplace.contract.UnpackLog(event, "RevealSuccessful", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AuctionMarketplaceRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the AuctionMarketplace contract.
type AuctionMarketplaceRoleAdminChangedIterator struct {
	Event *AuctionMarketplaceRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *AuctionMarketplaceRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AuctionMarketplaceRoleAdminChanged)
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
		it.Event = new(AuctionMarketplaceRoleAdminChanged)
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
func (it *AuctionMarketplaceRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AuctionMarketplaceRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AuctionMarketplaceRoleAdminChanged represents a RoleAdminChanged event raised by the AuctionMarketplace contract.
type AuctionMarketplaceRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_AuctionMarketplace *AuctionMarketplaceFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*AuctionMarketplaceRoleAdminChangedIterator, error) {

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

	logs, sub, err := _AuctionMarketplace.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &AuctionMarketplaceRoleAdminChangedIterator{contract: _AuctionMarketplace.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_AuctionMarketplace *AuctionMarketplaceFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *AuctionMarketplaceRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _AuctionMarketplace.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AuctionMarketplaceRoleAdminChanged)
				if err := _AuctionMarketplace.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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
func (_AuctionMarketplace *AuctionMarketplaceFilterer) ParseRoleAdminChanged(log types.Log) (*AuctionMarketplaceRoleAdminChanged, error) {
	event := new(AuctionMarketplaceRoleAdminChanged)
	if err := _AuctionMarketplace.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AuctionMarketplaceRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the AuctionMarketplace contract.
type AuctionMarketplaceRoleGrantedIterator struct {
	Event *AuctionMarketplaceRoleGranted // Event containing the contract specifics and raw log

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
func (it *AuctionMarketplaceRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AuctionMarketplaceRoleGranted)
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
		it.Event = new(AuctionMarketplaceRoleGranted)
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
func (it *AuctionMarketplaceRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AuctionMarketplaceRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AuctionMarketplaceRoleGranted represents a RoleGranted event raised by the AuctionMarketplace contract.
type AuctionMarketplaceRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_AuctionMarketplace *AuctionMarketplaceFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*AuctionMarketplaceRoleGrantedIterator, error) {

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

	logs, sub, err := _AuctionMarketplace.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &AuctionMarketplaceRoleGrantedIterator{contract: _AuctionMarketplace.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_AuctionMarketplace *AuctionMarketplaceFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *AuctionMarketplaceRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _AuctionMarketplace.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AuctionMarketplaceRoleGranted)
				if err := _AuctionMarketplace.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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
func (_AuctionMarketplace *AuctionMarketplaceFilterer) ParseRoleGranted(log types.Log) (*AuctionMarketplaceRoleGranted, error) {
	event := new(AuctionMarketplaceRoleGranted)
	if err := _AuctionMarketplace.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AuctionMarketplaceRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the AuctionMarketplace contract.
type AuctionMarketplaceRoleRevokedIterator struct {
	Event *AuctionMarketplaceRoleRevoked // Event containing the contract specifics and raw log

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
func (it *AuctionMarketplaceRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AuctionMarketplaceRoleRevoked)
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
		it.Event = new(AuctionMarketplaceRoleRevoked)
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
func (it *AuctionMarketplaceRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AuctionMarketplaceRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AuctionMarketplaceRoleRevoked represents a RoleRevoked event raised by the AuctionMarketplace contract.
type AuctionMarketplaceRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_AuctionMarketplace *AuctionMarketplaceFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*AuctionMarketplaceRoleRevokedIterator, error) {

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

	logs, sub, err := _AuctionMarketplace.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &AuctionMarketplaceRoleRevokedIterator{contract: _AuctionMarketplace.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_AuctionMarketplace *AuctionMarketplaceFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *AuctionMarketplaceRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _AuctionMarketplace.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AuctionMarketplaceRoleRevoked)
				if err := _AuctionMarketplace.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
func (_AuctionMarketplace *AuctionMarketplaceFilterer) ParseRoleRevoked(log types.Log) (*AuctionMarketplaceRoleRevoked, error) {
	event := new(AuctionMarketplaceRoleRevoked)
	if err := _AuctionMarketplace.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AuctionMarketplaceUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the AuctionMarketplace contract.
type AuctionMarketplaceUnpausedIterator struct {
	Event *AuctionMarketplaceUnpaused // Event containing the contract specifics and raw log

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
func (it *AuctionMarketplaceUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AuctionMarketplaceUnpaused)
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
		it.Event = new(AuctionMarketplaceUnpaused)
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
func (it *AuctionMarketplaceUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AuctionMarketplaceUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AuctionMarketplaceUnpaused represents a Unpaused event raised by the AuctionMarketplace contract.
type AuctionMarketplaceUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_AuctionMarketplace *AuctionMarketplaceFilterer) FilterUnpaused(opts *bind.FilterOpts) (*AuctionMarketplaceUnpausedIterator, error) {

	logs, sub, err := _AuctionMarketplace.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &AuctionMarketplaceUnpausedIterator{contract: _AuctionMarketplace.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_AuctionMarketplace *AuctionMarketplaceFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *AuctionMarketplaceUnpaused) (event.Subscription, error) {

	logs, sub, err := _AuctionMarketplace.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AuctionMarketplaceUnpaused)
				if err := _AuctionMarketplace.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
func (_AuctionMarketplace *AuctionMarketplaceFilterer) ParseUnpaused(log types.Log) (*AuctionMarketplaceUnpaused, error) {
	event := new(AuctionMarketplaceUnpaused)
	if err := _AuctionMarketplace.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
