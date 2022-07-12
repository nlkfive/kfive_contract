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

// MarketplaceStorageMetaData contains all meta data concerning the MarketplaceStorage contract.
var MarketplaceStorageMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"AssetNotExisted\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AssetUnvailable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AuctionAlreadyEnded\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidMkpSender\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"TooEarly\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"TooLate\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"bidder\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"blindAuctionId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"blindedBid\",\"type\":\"bytes32\"}],\"name\":\"BlindAuctionBidSuccessful\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"canceller\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"blindAuctionId\",\"type\":\"bytes32\"}],\"name\":\"BlindAuctionCancelled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"nftAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"blindAuctionId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"assetId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"biddingEnd\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"revealEnd\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"startPriceInWei\",\"type\":\"uint256\"}],\"name\":\"BlindAuctionCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"blindAuctionId\",\"type\":\"bytes32\"}],\"name\":\"BlindAuctionEnded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"bidder\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"blindAuctionId\",\"type\":\"bytes32\"}],\"name\":\"BlindAuctionRefund\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"buyer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"blindAuctionId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalPrice\",\"type\":\"uint256\"}],\"name\":\"BlindAuctionSuccessful\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"who\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"}],\"name\":\"OrderCancelled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"orderId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"assetId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"nftAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"priceInWei\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"expiredAt\",\"type\":\"uint256\"}],\"name\":\"OrderCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"buyer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"}],\"name\":\"OrderSuccessful\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"bidder\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"publicAuctionId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"bidValue\",\"type\":\"uint256\"}],\"name\":\"PublicAuctionBidSuccessful\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"canceller\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"publicAuctionId\",\"type\":\"bytes32\"}],\"name\":\"PublicAuctionCancelled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"nftAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"publicAuctionId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"assetId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"biddingEnd\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"startPriceInWei\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"minIncrement\",\"type\":\"uint256\"}],\"name\":\"PublicAuctionCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"publicAuctionId\",\"type\":\"bytes32\"}],\"name\":\"PublicAuctionEnded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"bidder\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"publicAuctionId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"PublicAuctionRefund\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"buyer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"publicAuctionId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalPrice\",\"type\":\"uint256\"}],\"name\":\"PublicAuctionSuccessful\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"blindAuctionId\",\"type\":\"bytes32\"}],\"name\":\"RevealSuccessful\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PAUSER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getRoleMember\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleMemberCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"orderByNftAsset\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"orderId\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"nftAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expiredAt\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_publicAuctionMarketplace\",\"type\":\"address\"}],\"name\":\"updatePublicAuctionMarketplace\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_blindAuctionMarketplace\",\"type\":\"address\"}],\"name\":\"updateBlindAuctionMarketplace\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_orderMarketplace\",\"type\":\"address\"}],\"name\":\"updateOrderMarketplace\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"nftAsset\",\"type\":\"bytes32\"}],\"name\":\"assetIsAvailable\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"publicAuctionId\",\"type\":\"bytes32\"}],\"name\":\"publicAuctionIsExisted\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"nftAsset\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"publicAuctionId\",\"type\":\"bytes32\"}],\"name\":\"publicAuctionIsEnded\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"nftAsset\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"publicAuctionId\",\"type\":\"bytes32\"}],\"name\":\"publicAuctionIsRunning\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"assetOwner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"nftAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"assetId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"publicAuctionId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"biddingEnd\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startPriceInWei\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minIncrement\",\"type\":\"uint256\"}],\"name\":\"createPublicAuction\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"publicAuctionId\",\"type\":\"bytes32\"}],\"name\":\"getPublicAuction\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"highestBid\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"highestBidder\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"biddingEnd\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minIncrement\",\"type\":\"uint256\"}],\"internalType\":\"structIPublicAuction.PublicAuction\",\"name\":\"publicAuction\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"nftAsset\",\"type\":\"bytes32\"}],\"name\":\"endPublicAuction\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"bidder\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"highestBid\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"publicAuctionId\",\"type\":\"bytes32\"}],\"name\":\"updateHighestBidPublicAuction\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"blindAuctionId\",\"type\":\"bytes32\"}],\"name\":\"blindAuctionIsExisted\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"nftAsset\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"blindAuctionId\",\"type\":\"bytes32\"}],\"name\":\"blindAuctionIsEnded\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"nftAsset\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"blindAuctionId\",\"type\":\"bytes32\"}],\"name\":\"blindAuctionIsRunning\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"assetOwner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"nftAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"assetId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"blindAuctionId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"biddingEnd\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"revealEnd\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startPriceInWei\",\"type\":\"uint256\"}],\"name\":\"createBlindAuction\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"blindAuctionId\",\"type\":\"bytes32\"}],\"name\":\"getBlindAuction\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"biddingEnd\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"revealEnd\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"highestBid\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"highestBidder\",\"type\":\"address\"}],\"internalType\":\"structIBlindAuction.BlindAuction\",\"name\":\"blindAuction\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"nftAsset\",\"type\":\"bytes32\"}],\"name\":\"endBlindAuction\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"bidder\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"highestBid\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"blindAuctionId\",\"type\":\"bytes32\"}],\"name\":\"updateHighestBidBlindAuction\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"nftAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"assetId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"orderId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"priceInWei\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expiredAt\",\"type\":\"uint256\"}],\"name\":\"createOrder\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"nftAsset\",\"type\":\"bytes32\"}],\"name\":\"getOrder\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"orderId\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"nftAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expiredAt\",\"type\":\"uint256\"}],\"internalType\":\"structIOrder.Order\",\"name\":\"order\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"nftAsset\",\"type\":\"bytes32\"}],\"name\":\"deleteOrder\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60806040523480156200001157600080fd5b5062000032620000266200010a60201b60201c565b6200011260201b60201c565b6000600260146101000a81548160ff021916908315150217905550620000716000801b620000656200010a60201b60201c565b620001d860201b60201c565b620000b27fa49807205ce4d355092ef5a8a18f56e8913cf4a201fbe287825b095693c21775620000a66200010a60201b60201c565b620001d860201b60201c565b620001047f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a7fa49807205ce4d355092ef5a8a18f56e8913cf4a201fbe287825b095693c21775620001ee60201b60201c565b620005b6565b600033905090565b6000600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905081600260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b620001ea82826200025160201b60201c565b5050565b600062000201836200031f60201b60201c565b905081600080858152602001908152602001600020600101819055508181847fbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff60405160405180910390a4505050565b6200026882826200033e60201b620020111760201c565b7fa49807205ce4d355092ef5a8a18f56e8913cf4a201fbe287825b095693c217758214806200029957506000801b82145b156200031b57620002d67f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a826200033e60201b620020111760201c565b6000801b8214156200031a57620003197fa49807205ce4d355092ef5a8a18f56e8913cf4a201fbe287825b095693c21775826200033e60201b620020111760201c565b5b5b5050565b6000806000838152602001908152602001600020600101549050919050565b6200035582826200038660201b620020451760201c565b6200038181600160008581526020019081526020016000206200047760201b620021251790919060201c565b505050565b620003988282620004af60201b60201c565b6200047357600160008084815260200190815260200160002060000160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff021916908315150217905550620004186200010a60201b60201c565b73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16837f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a45b5050565b6000620004a7836000018373ffffffffffffffffffffffffffffffffffffffff1660001b6200051960201b60201c565b905092915050565b600080600084815260200190815260200160002060000160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16905092915050565b60006200052d83836200059360201b60201c565b620005885782600001829080600181540180825580915050600190039060005260206000200160009091909190915055826000018054905083600101600084815260200190815260200160002081905550600190506200058d565b600090505b92915050565b600080836001016000848152602001908152602001600020541415905092915050565b613d9280620005c66000396000f3fe608060405234801561001057600080fd5b506004361061023c5760003560e01c806375b238fc1161013b578063a217fddf116100b8578063d7e4a3071161007c578063d7e4a3071461070f578063e63ab1e91461072b578063e7c847aa14610749578063ef2d8b5514610765578063f2fde38b146107815761023c565b8063a217fddf14610645578063bc4124c514610663578063ca15c87314610693578063cd23b172146106c3578063d547741f146106f35761023c565b80639010d07c116100ff5780639010d07c1461059157806391d14854146105c1578063930ca094146105f15780639a91d9c01461060d5780639d90ce3b146106295761023c565b806375b238fc146104ff5780638456cb591461051d578063866244481461052757806387a61cbd146105575780638da5cb5b146105735761023c565b80634a43c5b9116101c95780635c975abb1161018d5780635c975abb1461045b5780636af4c895146104795780636ed3be5314610495578063715018a6146104c55780637584f03d146104cf5761023c565b80634a43c5b91461039357806351b7e09d146103c35780635778472a146103df57806357a45b771461040f5780635a0affbb1461043f5761023c565b8063248a9ca311610210578063248a9ca3146102ed5780632f2ff15d1461031d57806336568abe146103395780633a3999d9146103555780633f4ba83a146103895761023c565b80622cb6a31461024157806301ffc9a7146102715780630a7f8977146102a157806315c97238146102d1575b600080fd5b61025b60048036038101906102569190613275565b61079d565b6040516102689190613734565b60405180910390f35b61028b600480360381019061028691906132ed565b6107bd565b6040516102989190613734565b60405180910390f35b6102bb60048036038101906102b69190613210565b610837565b6040516102c89190613734565b60405180910390f35b6102eb60048036038101906102e69190613210565b61089e565b005b61030760048036038101906103029190613210565b6109e0565b604051610314919061374f565b60405180910390f35b61033760048036038101906103329190613239565b6109ff565b005b610353600480360381019061034e9190613239565b610a20565b005b61036f600480360381019061036a9190613210565b610aa3565b60405161038095949392919061376a565b60405180910390f35b610391610b19565b005b6103ad60048036038101906103a89190613275565b610b4e565b6040516103ba9190613734565b60405180910390f35b6103dd60048036038101906103d89190613071565b610b6e565b005b6103f960048036038101906103f49190613210565b610c25565b60405161040691906138ba565b60405180910390f35b61042960048036038101906104249190613210565b610d1c565b60405161043691906138d5565b60405180910390f35b61045960048036038101906104549190613210565b610e27565b005b610463610f69565b6040516104709190613734565b60405180910390f35b610493600480360381019061048e9190613123565b610f80565b005b6104af60048036038101906104aa9190613210565b611165565b6040516104bc919061389f565b60405180910390f35b6104cd611270565b005b6104e960048036038101906104e49190613210565b6112f8565b6040516104f69190613734565b60405180910390f35b61050761131e565b604051610514919061374f565b60405180910390f35b610525611342565b005b610541600480360381019061053c9190613210565b611377565b60405161054e9190613734565b60405180910390f35b610571600480360381019061056c9190613210565b61139d565b005b61057b611547565b6040516105889190613719565b60405180910390f35b6105ab60048036038101906105a691906132b1565b611571565b6040516105b89190613719565b60405180910390f35b6105db60048036038101906105d69190613239565b6115a0565b6040516105e89190613734565b60405180910390f35b61060b6004803603810190610606919061309a565b61160a565b005b61062760048036038101906106229190613071565b61186c565b005b610643600480360381019061063e91906131c1565b611923565b005b61064d611a71565b60405161065a919061374f565b60405180910390f35b61067d60048036038101906106789190613275565b611a78565b60405161068a9190613734565b60405180910390f35b6106ad60048036038101906106a89190613210565b611a99565b6040516106ba91906138f0565b60405180910390f35b6106dd60048036038101906106d89190613275565b611abd565b6040516106ea9190613734565b60405180910390f35b61070d60048036038101906107089190613239565b611ade565b005b61072960048036038101906107249190613123565b611aff565b005b610733611ce4565b604051610740919061374f565b60405180910390f35b610763600480360381019061075e91906131c1565b611d08565b005b61077f600480360381019061077a9190613071565b611e56565b005b61079b60048036038101906107969190613071565b611f0d565b005b600081600760008581526020019081526020016000205414905092915050565b60007f30baaa08000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916827bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19161480610830575061082f82612155565b5b9050919050565b60008060001b600760008481526020019081526020016000205414801561087357506000801b6009600084815260200190815260200160002054145b801561089757506000801b600a600084815260200190815260200160002060000154145b9050919050565b6108a6610f69565b156108e6576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016108dd9061383f565b60405180910390fd5b600560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff168073ffffffffffffffffffffffffffffffffffffffff166109286121cf565b73ffffffffffffffffffffffffffffffffffffffff1614610975576040517f2d0b96e300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6000801b600960008481526020019081526020016000205414156109c5576040517fd02e774d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60096000838152602001908152602001600020600090555050565b6000806000838152602001908152602001600020600101549050919050565b610a08826109e0565b610a11816121d7565b610a1b83836121eb565b505050565b610a286121cf565b73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614610a95576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610a8c9061387f565b60405180910390fd5b610a9f828261228f565b5050565b600a6020528060005260406000206000915090508060000154908060010160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060020160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060030154908060040154905085565b7f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a610b43816121d7565b610b4b612333565b50565b600081600960008581526020019081526020016000205414905092915050565b610b76610f69565b15610bb6576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610bad9061383f565b60405180910390fd5b7fa49807205ce4d355092ef5a8a18f56e8913cf4a201fbe287825b095693c21775610be0816121d7565b81600560006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505050565b610c2d612ee7565b600a60008381526020019081526020016000206040518060a0016040529081600082015481526020016001820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016002820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001600382015481526020016004820154815250509050919050565b610d24612f45565b600660008381526020019081526020016000206040518060e0016040529081600082015481526020016001820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001600282015481526020016003820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200160048201548152602001600582015481526020016006820154815250509050919050565b610e2f610f69565b15610e6f576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610e669061383f565b60405180910390fd5b600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff168073ffffffffffffffffffffffffffffffffffffffff16610eb16121cf565b73ffffffffffffffffffffffffffffffffffffffff1614610efe576040517f2d0b96e300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6000801b60076000848152602001908152602001600020541415610f4e576040517fd02e774d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60076000838152602001908152602001600020600090555050565b6000600260149054906101000a900460ff16905090565b610f88610f69565b15610fc8576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610fbf9061383f565b60405180910390fd5b600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff168073ffffffffffffffffffffffffffffffffffffffff1661100a6121cf565b73ffffffffffffffffffffffffffffffffffffffff1614611057576040517f2d0b96e300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6000878760405160200161106c9291906136b3565b60405160208183030381529060405280519060200120905061108d81610837565b6110c3576040517fc0a24f6700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8560076000838152602001908152602001600020819055506000600660008881526020019081526020016000209050868160000181905550898160010160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555085816004018190555084816005018190555083816006018190555050505050505050505050565b61116d612fb1565b600860008381526020019081526020016000206040518060e0016040529081600082015481526020016001820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001600282015481526020016003820154815260200160048201548152602001600582015481526020016006820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815250509050919050565b6112786121cf565b73ffffffffffffffffffffffffffffffffffffffff16611296611547565b73ffffffffffffffffffffffffffffffffffffffff16146112ec576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016112e39061385f565b60405180910390fd5b6112f660006123d5565b565b60008060001b600660008481526020019081526020016000206000015414159050919050565b7fa49807205ce4d355092ef5a8a18f56e8913cf4a201fbe287825b095693c2177581565b7f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a61136c816121d7565b61137461249b565b50565b60008060001b600860008481526020019081526020016000206000015414159050919050565b6113a5610f69565b156113e5576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016113dc9061383f565b60405180910390fd5b600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff168073ffffffffffffffffffffffffffffffffffffffff166114276121cf565b73ffffffffffffffffffffffffffffffffffffffff1614611474576040517f2d0b96e300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6000801b600a60008481526020019081526020016000206000015414156114c7576040517f84791e5200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600a60008381526020019081526020016000206000808201600090556001820160006101000a81549073ffffffffffffffffffffffffffffffffffffffff02191690556002820160006101000a81549073ffffffffffffffffffffffffffffffffffffffff02191690556003820160009055600482016000905550505050565b6000600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b6000611598826001600086815260200190815260200160002061253e90919063ffffffff16565b905092915050565b600080600084815260200190815260200160002060000160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16905092915050565b600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff168073ffffffffffffffffffffffffffffffffffffffff1661164c6121cf565b73ffffffffffffffffffffffffffffffffffffffff1614611699576040517f2d0b96e300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6116a1610f69565b156116e1576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016116d89061383f565b60405180910390fd5b600086866040516020016116f69291906136b3565b60405160208183030381529060405280519060200120905061171781610837565b61174d576040517fc0a24f6700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6040518060a001604052808681526020018973ffffffffffffffffffffffffffffffffffffffff1681526020018873ffffffffffffffffffffffffffffffffffffffff16815260200185815260200184815250600a60008381526020019081526020016000206000820151816000015560208201518160010160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060408201518160020160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060608201518160030155608082015181600401559050505050505050505050565b611874610f69565b156118b4576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016118ab9061383f565b60405180910390fd5b7fa49807205ce4d355092ef5a8a18f56e8913cf4a201fbe287825b095693c217756118de816121d7565b81600360006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505050565b61192b610f69565b1561196b576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016119629061383f565b60405180910390fd5b600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff168073ffffffffffffffffffffffffffffffffffffffff166119ad6121cf565b73ffffffffffffffffffffffffffffffffffffffff16146119fa576040517f2d0b96e300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6000600660008481526020019081526020016000209050611a1e8160040154612558565b838160020181905550848160030160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505050505050565b6000801b81565b60008160076000858152602001908152602001600020541415905092915050565b6000611ab66001600084815260200190815260200160002061259f565b9050919050565b60008160096000858152602001908152602001600020541415905092915050565b611ae7826109e0565b611af0816121d7565b611afa838361228f565b505050565b611b07610f69565b15611b47576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611b3e9061383f565b60405180910390fd5b600560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff168073ffffffffffffffffffffffffffffffffffffffff16611b896121cf565b73ffffffffffffffffffffffffffffffffffffffff1614611bd6576040517f2d0b96e300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60008787604051602001611beb9291906136b3565b604051602081830303815290604052805190602001209050611c0c81610837565b611c42576040517fc0a24f6700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8560096000838152602001908152602001600020819055506000600860008881526020019081526020016000209050868160000181905550898160010160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555085816002018190555084816003018190555083816004018190555050505050505050505050565b7f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a81565b611d10610f69565b15611d50576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611d479061383f565b60405180910390fd5b600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff168073ffffffffffffffffffffffffffffffffffffffff16611d926121cf565b73ffffffffffffffffffffffffffffffffffffffff1614611ddf576040517f2d0b96e300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6000600860008481526020019081526020016000209050611e038160020154612558565b838160050181905550848160060160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505050505050565b611e5e610f69565b15611e9e576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611e959061383f565b60405180910390fd5b7fa49807205ce4d355092ef5a8a18f56e8913cf4a201fbe287825b095693c21775611ec8816121d7565b81600460006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505050565b611f156121cf565b73ffffffffffffffffffffffffffffffffffffffff16611f33611547565b73ffffffffffffffffffffffffffffffffffffffff1614611f89576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611f809061385f565b60405180910390fd5b611fb37fa49807205ce4d355092ef5a8a18f56e8913cf4a201fbe287825b095693c21775826121eb565b611fc06000801b826121eb565b611ff17fa49807205ce4d355092ef5a8a18f56e8913cf4a201fbe287825b095693c21775611fec611547565b61228f565b6120056000801b612000611547565b61228f565b61200e816125b4565b50565b61201b8282612045565b612040816001600085815260200190815260200160002061212590919063ffffffff16565b505050565b61204f82826115a0565b61212157600160008084815260200190815260200160002060000160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055506120c66121cf565b73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16837f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a45b5050565b600061214d836000018373ffffffffffffffffffffffffffffffffffffffff1660001b6126ac565b905092915050565b60007f5a05180f000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916827bffffffffffffffffffffffffffffffffffffffffffffffffffffffff191614806121c857506121c78261271c565b5b9050919050565b600033905090565b6121e8816121e36121cf565b612796565b50565b6121f58282612011565b7fa49807205ce4d355092ef5a8a18f56e8913cf4a201fbe287825b095693c2177582148061222557506000801b82145b1561228b576122547f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a82612011565b6000801b82141561228a576122897fa49807205ce4d355092ef5a8a18f56e8913cf4a201fbe287825b095693c2177582612011565b5b5b5050565b6122998282612833565b7fa49807205ce4d355092ef5a8a18f56e8913cf4a201fbe287825b095693c217758214806122c957506000801b82145b1561232f576122f87f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a82612833565b6000801b82141561232e5761232d7fa49807205ce4d355092ef5a8a18f56e8913cf4a201fbe287825b095693c2177582612011565b5b5b5050565b61233b610f69565b61237a576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612371906137ff565b60405180910390fd5b6000600260146101000a81548160ff0219169083151502179055507f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa6123be6121cf565b6040516123cb9190613719565b60405180910390a1565b6000600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905081600260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b6124a3610f69565b156124e3576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016124da9061383f565b60405180910390fd5b6001600260146101000a81548160ff0219169083151502179055507f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a2586125276121cf565b6040516125349190613719565b60405180910390a1565b600061254d8360000183612867565b60001c905092915050565b80421061259c57806040517f691e568200000000000000000000000000000000000000000000000000000000815260040161259391906138f0565b60405180910390fd5b50565b60006125ad826000016128b8565b9050919050565b6125bc6121cf565b73ffffffffffffffffffffffffffffffffffffffff166125da611547565b73ffffffffffffffffffffffffffffffffffffffff1614612630576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016126279061385f565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614156126a0576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016126979061381f565b60405180910390fd5b6126a9816123d5565b50565b60006126b883836128c9565b612711578260000182908060018154018082558091505060019003906000526020600020016000909190919091505582600001805490508360010160008481526020019081526020016000208190555060019050612716565b600090505b92915050565b60007f7965db0b000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916827bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916148061278f575061278e826128ec565b5b9050919050565b6127a082826115a0565b61282f576127c58173ffffffffffffffffffffffffffffffffffffffff166014612956565b6127d38360001c6020612956565b6040516020016127e49291906136df565b6040516020818303038152906040526040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161282691906137bd565b60405180910390fd5b5050565b61283d8282612c50565b6128628160016000858152602001908152602001600020612d3190919063ffffffff16565b505050565b60008260000182815481106128a5577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b9060005260206000200154905092915050565b600081600001805490509050919050565b600080836001016000848152602001908152602001600020541415905092915050565b60007f01ffc9a7000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916827bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916149050919050565b6060600060028360026129699190613988565b6129739190613932565b67ffffffffffffffff8111156129b2577f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6040519080825280601f01601f1916602001820160405280156129e45781602001600182028036833780820191505090505b5090507f300000000000000000000000000000000000000000000000000000000000000081600081518110612a42577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a9053507f780000000000000000000000000000000000000000000000000000000000000081600181518110612acc577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a90535060006001846002612b0c9190613988565b612b169190613932565b90505b6001811115612c02577f3031323334353637383961626364656600000000000000000000000000000000600f861660108110612b7e577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b1a60f81b828281518110612bbb577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a905350600485901c945080612bfb90613ac7565b9050612b19565b5060008414612c46576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612c3d906137df565b60405180910390fd5b8091505092915050565b612c5a82826115a0565b15612d2d57600080600084815260200190815260200160002060000160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff021916908315150217905550612cd26121cf565b73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16837ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b60405160405180910390a45b5050565b6000612d59836000018373ffffffffffffffffffffffffffffffffffffffff1660001b612d61565b905092915050565b60008083600101600084815260200190815260200160002054905060008114612edb576000600182612d9391906139e2565b9050600060018660000180549050612dab91906139e2565b9050818114612e66576000866000018281548110612df2577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b9060005260206000200154905080876000018481548110612e3c577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b90600052602060002001819055508387600101600083815260200190815260200160002081905550505b85600001805480612ea0577f4e487b7100000000000000000000000000000000000000000000000000000000600052603160045260246000fd5b600190038181906000526020600020016000905590558560010160008681526020019081526020016000206000905560019350505050612ee1565b60009150505b92915050565b6040518060a0016040528060008019168152602001600073ffffffffffffffffffffffffffffffffffffffff168152602001600073ffffffffffffffffffffffffffffffffffffffff16815260200160008152602001600081525090565b6040518060e0016040528060008019168152602001600073ffffffffffffffffffffffffffffffffffffffff16815260200160008152602001600073ffffffffffffffffffffffffffffffffffffffff1681526020016000815260200160008152602001600081525090565b6040518060e0016040528060008019168152602001600073ffffffffffffffffffffffffffffffffffffffff16815260200160008152602001600081526020016000815260200160008152602001600073ffffffffffffffffffffffffffffffffffffffff1681525090565b60008135905061302c81613d00565b92915050565b60008135905061304181613d17565b92915050565b60008135905061305681613d2e565b92915050565b60008135905061306b81613d45565b92915050565b60006020828403121561308357600080fd5b60006130918482850161301d565b91505092915050565b60008060008060008060c087890312156130b357600080fd5b60006130c189828a0161301d565b96505060206130d289828a0161301d565b95505060406130e389828a0161305c565b94505060606130f489828a01613032565b935050608061310589828a0161305c565b92505060a061311689828a0161305c565b9150509295509295509295565b600080600080600080600060e0888a03121561313e57600080fd5b600061314c8a828b0161301d565b975050602061315d8a828b0161301d565b965050604061316e8a828b0161305c565b955050606061317f8a828b01613032565b94505060806131908a828b0161305c565b93505060a06131a18a828b0161305c565b92505060c06131b28a828b0161305c565b91505092959891949750929550565b6000806000606084860312156131d657600080fd5b60006131e48682870161301d565b93505060206131f58682870161305c565b925050604061320686828701613032565b9150509250925092565b60006020828403121561322257600080fd5b600061323084828501613032565b91505092915050565b6000806040838503121561324c57600080fd5b600061325a85828601613032565b925050602061326b8582860161301d565b9150509250929050565b6000806040838503121561328857600080fd5b600061329685828601613032565b92505060206132a785828601613032565b9150509250929050565b600080604083850312156132c457600080fd5b60006132d285828601613032565b92505060206132e38582860161305c565b9150509250929050565b6000602082840312156132ff57600080fd5b600061330d84828501613047565b91505092915050565b61331f81613a16565b82525050565b61332e81613a16565b82525050565b61334561334082613a16565b613af1565b82525050565b61335481613a28565b82525050565b61336381613a34565b82525050565b61337281613a34565b82525050565b60006133838261390b565b61338d8185613916565b935061339d818560208601613a94565b6133a681613b4e565b840191505092915050565b60006133bc8261390b565b6133c68185613927565b93506133d6818560208601613a94565b80840191505092915050565b60006133ef602083613916565b91506133fa82613b6c565b602082019050919050565b6000613412601483613916565b915061341d82613b95565b602082019050919050565b6000613435602683613916565b915061344082613bbe565b604082019050919050565b6000613458601083613916565b915061346382613c0d565b602082019050919050565b600061347b602083613916565b915061348682613c36565b602082019050919050565b600061349e601783613927565b91506134a982613c5f565b601782019050919050565b60006134c1601183613927565b91506134cc82613c88565b601182019050919050565b60006134e4602f83613916565b91506134ef82613cb1565b604082019050919050565b60e082016000820151613510600085018261335a565b5060208201516135236020850182613316565b506040820151613536604085018261367e565b506060820151613549606085018261367e565b50608082015161355c608085018261367e565b5060a082015161356f60a085018261367e565b5060c082015161358260c0850182613316565b50505050565b60a08201600082015161359e600085018261335a565b5060208201516135b16020850182613316565b5060408201516135c46040850182613316565b5060608201516135d7606085018261367e565b5060808201516135ea608085018261367e565b50505050565b60e082016000820151613606600085018261335a565b5060208201516136196020850182613316565b50604082015161362c604085018261367e565b50606082015161363f6060850182613316565b506080820151613652608085018261367e565b5060a082015161366560a085018261367e565b5060c082015161367860c085018261367e565b50505050565b61368781613a8a565b82525050565b61369681613a8a565b82525050565b6136ad6136a882613a8a565b613b15565b82525050565b60006136bf8285613334565b6014820191506136cf828461369c565b6020820191508190509392505050565b60006136ea82613491565b91506136f682856133b1565b9150613701826134b4565b915061370d82846133b1565b91508190509392505050565b600060208201905061372e6000830184613325565b92915050565b6000602082019050613749600083018461334b565b92915050565b60006020820190506137646000830184613369565b92915050565b600060a08201905061377f6000830188613369565b61378c6020830187613325565b6137996040830186613325565b6137a6606083018561368d565b6137b3608083018461368d565b9695505050505050565b600060208201905081810360008301526137d78184613378565b905092915050565b600060208201905081810360008301526137f8816133e2565b9050919050565b6000602082019050818103600083015261381881613405565b9050919050565b6000602082019050818103600083015261383881613428565b9050919050565b600060208201905081810360008301526138588161344b565b9050919050565b600060208201905081810360008301526138788161346e565b9050919050565b60006020820190508181036000830152613898816134d7565b9050919050565b600060e0820190506138b460008301846134fa565b92915050565b600060a0820190506138cf6000830184613588565b92915050565b600060e0820190506138ea60008301846135f0565b92915050565b6000602082019050613905600083018461368d565b92915050565b600081519050919050565b600082825260208201905092915050565b600081905092915050565b600061393d82613a8a565b915061394883613a8a565b9250827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0382111561397d5761397c613b1f565b5b828201905092915050565b600061399382613a8a565b915061399e83613a8a565b9250817fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff04831182151516156139d7576139d6613b1f565b5b828202905092915050565b60006139ed82613a8a565b91506139f883613a8a565b925082821015613a0b57613a0a613b1f565b5b828203905092915050565b6000613a2182613a6a565b9050919050565b60008115159050919050565b6000819050919050565b60007fffffffff0000000000000000000000000000000000000000000000000000000082169050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b60005b83811015613ab2578082015181840152602081019050613a97565b83811115613ac1576000848401525b50505050565b6000613ad282613a8a565b91506000821415613ae657613ae5613b1f565b5b600182039050919050565b6000613afc82613b03565b9050919050565b6000613b0e82613b5f565b9050919050565b6000819050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b6000601f19601f8301169050919050565b60008160601b9050919050565b7f537472696e67733a20686578206c656e67746820696e73756666696369656e74600082015250565b7f5061757361626c653a206e6f7420706175736564000000000000000000000000600082015250565b7f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160008201527f6464726573730000000000000000000000000000000000000000000000000000602082015250565b7f5061757361626c653a2070617573656400000000000000000000000000000000600082015250565b7f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572600082015250565b7f416363657373436f6e74726f6c3a206163636f756e7420000000000000000000600082015250565b7f206973206d697373696e6720726f6c6520000000000000000000000000000000600082015250565b7f416363657373436f6e74726f6c3a2063616e206f6e6c792072656e6f756e636560008201527f20726f6c657320666f722073656c660000000000000000000000000000000000602082015250565b613d0981613a16565b8114613d1457600080fd5b50565b613d2081613a34565b8114613d2b57600080fd5b50565b613d3781613a3e565b8114613d4257600080fd5b50565b613d4e81613a8a565b8114613d5957600080fd5b5056fea26469706673582212203f5c64b626bad79fddbc81fb053527cebd9a63361e08a32e56a36b523727f7cf64736f6c63430008040033",
}

// MarketplaceStorageABI is the input ABI used to generate the binding from.
// Deprecated: Use MarketplaceStorageMetaData.ABI instead.
var MarketplaceStorageABI = MarketplaceStorageMetaData.ABI

// MarketplaceStorageBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use MarketplaceStorageMetaData.Bin instead.
var MarketplaceStorageBin = MarketplaceStorageMetaData.Bin

// DeployMarketplaceStorage deploys a new Ethereum contract, binding an instance of MarketplaceStorage to it.
func DeployMarketplaceStorage(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *MarketplaceStorage, error) {
	parsed, err := MarketplaceStorageMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(MarketplaceStorageBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &MarketplaceStorage{MarketplaceStorageCaller: MarketplaceStorageCaller{contract: contract}, MarketplaceStorageTransactor: MarketplaceStorageTransactor{contract: contract}, MarketplaceStorageFilterer: MarketplaceStorageFilterer{contract: contract}}, nil
}

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

// BlindAuctionIsEnded is a free data retrieval call binding the contract method 0xcd23b172.
//
// Solidity: function blindAuctionIsEnded(bytes32 nftAsset, bytes32 blindAuctionId) view returns(bool)
func (_MarketplaceStorage *MarketplaceStorageCaller) BlindAuctionIsEnded(opts *bind.CallOpts, nftAsset [32]byte, blindAuctionId [32]byte) (bool, error) {
	var out []interface{}
	err := _MarketplaceStorage.contract.Call(opts, &out, "blindAuctionIsEnded", nftAsset, blindAuctionId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// BlindAuctionIsEnded is a free data retrieval call binding the contract method 0xcd23b172.
//
// Solidity: function blindAuctionIsEnded(bytes32 nftAsset, bytes32 blindAuctionId) view returns(bool)
func (_MarketplaceStorage *MarketplaceStorageSession) BlindAuctionIsEnded(nftAsset [32]byte, blindAuctionId [32]byte) (bool, error) {
	return _MarketplaceStorage.Contract.BlindAuctionIsEnded(&_MarketplaceStorage.CallOpts, nftAsset, blindAuctionId)
}

// BlindAuctionIsEnded is a free data retrieval call binding the contract method 0xcd23b172.
//
// Solidity: function blindAuctionIsEnded(bytes32 nftAsset, bytes32 blindAuctionId) view returns(bool)
func (_MarketplaceStorage *MarketplaceStorageCallerSession) BlindAuctionIsEnded(nftAsset [32]byte, blindAuctionId [32]byte) (bool, error) {
	return _MarketplaceStorage.Contract.BlindAuctionIsEnded(&_MarketplaceStorage.CallOpts, nftAsset, blindAuctionId)
}

// BlindAuctionIsExisted is a free data retrieval call binding the contract method 0x86624448.
//
// Solidity: function blindAuctionIsExisted(bytes32 blindAuctionId) view returns(bool)
func (_MarketplaceStorage *MarketplaceStorageCaller) BlindAuctionIsExisted(opts *bind.CallOpts, blindAuctionId [32]byte) (bool, error) {
	var out []interface{}
	err := _MarketplaceStorage.contract.Call(opts, &out, "blindAuctionIsExisted", blindAuctionId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// BlindAuctionIsExisted is a free data retrieval call binding the contract method 0x86624448.
//
// Solidity: function blindAuctionIsExisted(bytes32 blindAuctionId) view returns(bool)
func (_MarketplaceStorage *MarketplaceStorageSession) BlindAuctionIsExisted(blindAuctionId [32]byte) (bool, error) {
	return _MarketplaceStorage.Contract.BlindAuctionIsExisted(&_MarketplaceStorage.CallOpts, blindAuctionId)
}

// BlindAuctionIsExisted is a free data retrieval call binding the contract method 0x86624448.
//
// Solidity: function blindAuctionIsExisted(bytes32 blindAuctionId) view returns(bool)
func (_MarketplaceStorage *MarketplaceStorageCallerSession) BlindAuctionIsExisted(blindAuctionId [32]byte) (bool, error) {
	return _MarketplaceStorage.Contract.BlindAuctionIsExisted(&_MarketplaceStorage.CallOpts, blindAuctionId)
}

// BlindAuctionIsRunning is a free data retrieval call binding the contract method 0x4a43c5b9.
//
// Solidity: function blindAuctionIsRunning(bytes32 nftAsset, bytes32 blindAuctionId) view returns(bool)
func (_MarketplaceStorage *MarketplaceStorageCaller) BlindAuctionIsRunning(opts *bind.CallOpts, nftAsset [32]byte, blindAuctionId [32]byte) (bool, error) {
	var out []interface{}
	err := _MarketplaceStorage.contract.Call(opts, &out, "blindAuctionIsRunning", nftAsset, blindAuctionId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// BlindAuctionIsRunning is a free data retrieval call binding the contract method 0x4a43c5b9.
//
// Solidity: function blindAuctionIsRunning(bytes32 nftAsset, bytes32 blindAuctionId) view returns(bool)
func (_MarketplaceStorage *MarketplaceStorageSession) BlindAuctionIsRunning(nftAsset [32]byte, blindAuctionId [32]byte) (bool, error) {
	return _MarketplaceStorage.Contract.BlindAuctionIsRunning(&_MarketplaceStorage.CallOpts, nftAsset, blindAuctionId)
}

// BlindAuctionIsRunning is a free data retrieval call binding the contract method 0x4a43c5b9.
//
// Solidity: function blindAuctionIsRunning(bytes32 nftAsset, bytes32 blindAuctionId) view returns(bool)
func (_MarketplaceStorage *MarketplaceStorageCallerSession) BlindAuctionIsRunning(nftAsset [32]byte, blindAuctionId [32]byte) (bool, error) {
	return _MarketplaceStorage.Contract.BlindAuctionIsRunning(&_MarketplaceStorage.CallOpts, nftAsset, blindAuctionId)
}

// GetBlindAuction is a free data retrieval call binding the contract method 0x6ed3be53.
//
// Solidity: function getBlindAuction(bytes32 blindAuctionId) view returns((bytes32,address,uint256,uint256,uint256,uint256,address) blindAuction)
func (_MarketplaceStorage *MarketplaceStorageCaller) GetBlindAuction(opts *bind.CallOpts, blindAuctionId [32]byte) (IBlindAuctionBlindAuction, error) {
	var out []interface{}
	err := _MarketplaceStorage.contract.Call(opts, &out, "getBlindAuction", blindAuctionId)

	if err != nil {
		return *new(IBlindAuctionBlindAuction), err
	}

	out0 := *abi.ConvertType(out[0], new(IBlindAuctionBlindAuction)).(*IBlindAuctionBlindAuction)

	return out0, err

}

// GetBlindAuction is a free data retrieval call binding the contract method 0x6ed3be53.
//
// Solidity: function getBlindAuction(bytes32 blindAuctionId) view returns((bytes32,address,uint256,uint256,uint256,uint256,address) blindAuction)
func (_MarketplaceStorage *MarketplaceStorageSession) GetBlindAuction(blindAuctionId [32]byte) (IBlindAuctionBlindAuction, error) {
	return _MarketplaceStorage.Contract.GetBlindAuction(&_MarketplaceStorage.CallOpts, blindAuctionId)
}

// GetBlindAuction is a free data retrieval call binding the contract method 0x6ed3be53.
//
// Solidity: function getBlindAuction(bytes32 blindAuctionId) view returns((bytes32,address,uint256,uint256,uint256,uint256,address) blindAuction)
func (_MarketplaceStorage *MarketplaceStorageCallerSession) GetBlindAuction(blindAuctionId [32]byte) (IBlindAuctionBlindAuction, error) {
	return _MarketplaceStorage.Contract.GetBlindAuction(&_MarketplaceStorage.CallOpts, blindAuctionId)
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

// GetPublicAuction is a free data retrieval call binding the contract method 0x57a45b77.
//
// Solidity: function getPublicAuction(bytes32 publicAuctionId) view returns((bytes32,address,uint256,address,uint256,uint256,uint256) publicAuction)
func (_MarketplaceStorage *MarketplaceStorageCaller) GetPublicAuction(opts *bind.CallOpts, publicAuctionId [32]byte) (IPublicAuctionPublicAuction, error) {
	var out []interface{}
	err := _MarketplaceStorage.contract.Call(opts, &out, "getPublicAuction", publicAuctionId)

	if err != nil {
		return *new(IPublicAuctionPublicAuction), err
	}

	out0 := *abi.ConvertType(out[0], new(IPublicAuctionPublicAuction)).(*IPublicAuctionPublicAuction)

	return out0, err

}

// GetPublicAuction is a free data retrieval call binding the contract method 0x57a45b77.
//
// Solidity: function getPublicAuction(bytes32 publicAuctionId) view returns((bytes32,address,uint256,address,uint256,uint256,uint256) publicAuction)
func (_MarketplaceStorage *MarketplaceStorageSession) GetPublicAuction(publicAuctionId [32]byte) (IPublicAuctionPublicAuction, error) {
	return _MarketplaceStorage.Contract.GetPublicAuction(&_MarketplaceStorage.CallOpts, publicAuctionId)
}

// GetPublicAuction is a free data retrieval call binding the contract method 0x57a45b77.
//
// Solidity: function getPublicAuction(bytes32 publicAuctionId) view returns((bytes32,address,uint256,address,uint256,uint256,uint256) publicAuction)
func (_MarketplaceStorage *MarketplaceStorageCallerSession) GetPublicAuction(publicAuctionId [32]byte) (IPublicAuctionPublicAuction, error) {
	return _MarketplaceStorage.Contract.GetPublicAuction(&_MarketplaceStorage.CallOpts, publicAuctionId)
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

// PublicAuctionIsEnded is a free data retrieval call binding the contract method 0xbc4124c5.
//
// Solidity: function publicAuctionIsEnded(bytes32 nftAsset, bytes32 publicAuctionId) view returns(bool)
func (_MarketplaceStorage *MarketplaceStorageCaller) PublicAuctionIsEnded(opts *bind.CallOpts, nftAsset [32]byte, publicAuctionId [32]byte) (bool, error) {
	var out []interface{}
	err := _MarketplaceStorage.contract.Call(opts, &out, "publicAuctionIsEnded", nftAsset, publicAuctionId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// PublicAuctionIsEnded is a free data retrieval call binding the contract method 0xbc4124c5.
//
// Solidity: function publicAuctionIsEnded(bytes32 nftAsset, bytes32 publicAuctionId) view returns(bool)
func (_MarketplaceStorage *MarketplaceStorageSession) PublicAuctionIsEnded(nftAsset [32]byte, publicAuctionId [32]byte) (bool, error) {
	return _MarketplaceStorage.Contract.PublicAuctionIsEnded(&_MarketplaceStorage.CallOpts, nftAsset, publicAuctionId)
}

// PublicAuctionIsEnded is a free data retrieval call binding the contract method 0xbc4124c5.
//
// Solidity: function publicAuctionIsEnded(bytes32 nftAsset, bytes32 publicAuctionId) view returns(bool)
func (_MarketplaceStorage *MarketplaceStorageCallerSession) PublicAuctionIsEnded(nftAsset [32]byte, publicAuctionId [32]byte) (bool, error) {
	return _MarketplaceStorage.Contract.PublicAuctionIsEnded(&_MarketplaceStorage.CallOpts, nftAsset, publicAuctionId)
}

// PublicAuctionIsExisted is a free data retrieval call binding the contract method 0x7584f03d.
//
// Solidity: function publicAuctionIsExisted(bytes32 publicAuctionId) view returns(bool)
func (_MarketplaceStorage *MarketplaceStorageCaller) PublicAuctionIsExisted(opts *bind.CallOpts, publicAuctionId [32]byte) (bool, error) {
	var out []interface{}
	err := _MarketplaceStorage.contract.Call(opts, &out, "publicAuctionIsExisted", publicAuctionId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// PublicAuctionIsExisted is a free data retrieval call binding the contract method 0x7584f03d.
//
// Solidity: function publicAuctionIsExisted(bytes32 publicAuctionId) view returns(bool)
func (_MarketplaceStorage *MarketplaceStorageSession) PublicAuctionIsExisted(publicAuctionId [32]byte) (bool, error) {
	return _MarketplaceStorage.Contract.PublicAuctionIsExisted(&_MarketplaceStorage.CallOpts, publicAuctionId)
}

// PublicAuctionIsExisted is a free data retrieval call binding the contract method 0x7584f03d.
//
// Solidity: function publicAuctionIsExisted(bytes32 publicAuctionId) view returns(bool)
func (_MarketplaceStorage *MarketplaceStorageCallerSession) PublicAuctionIsExisted(publicAuctionId [32]byte) (bool, error) {
	return _MarketplaceStorage.Contract.PublicAuctionIsExisted(&_MarketplaceStorage.CallOpts, publicAuctionId)
}

// PublicAuctionIsRunning is a free data retrieval call binding the contract method 0x002cb6a3.
//
// Solidity: function publicAuctionIsRunning(bytes32 nftAsset, bytes32 publicAuctionId) view returns(bool)
func (_MarketplaceStorage *MarketplaceStorageCaller) PublicAuctionIsRunning(opts *bind.CallOpts, nftAsset [32]byte, publicAuctionId [32]byte) (bool, error) {
	var out []interface{}
	err := _MarketplaceStorage.contract.Call(opts, &out, "publicAuctionIsRunning", nftAsset, publicAuctionId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// PublicAuctionIsRunning is a free data retrieval call binding the contract method 0x002cb6a3.
//
// Solidity: function publicAuctionIsRunning(bytes32 nftAsset, bytes32 publicAuctionId) view returns(bool)
func (_MarketplaceStorage *MarketplaceStorageSession) PublicAuctionIsRunning(nftAsset [32]byte, publicAuctionId [32]byte) (bool, error) {
	return _MarketplaceStorage.Contract.PublicAuctionIsRunning(&_MarketplaceStorage.CallOpts, nftAsset, publicAuctionId)
}

// PublicAuctionIsRunning is a free data retrieval call binding the contract method 0x002cb6a3.
//
// Solidity: function publicAuctionIsRunning(bytes32 nftAsset, bytes32 publicAuctionId) view returns(bool)
func (_MarketplaceStorage *MarketplaceStorageCallerSession) PublicAuctionIsRunning(nftAsset [32]byte, publicAuctionId [32]byte) (bool, error) {
	return _MarketplaceStorage.Contract.PublicAuctionIsRunning(&_MarketplaceStorage.CallOpts, nftAsset, publicAuctionId)
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

// CreateBlindAuction is a paid mutator transaction binding the contract method 0xd7e4a307.
//
// Solidity: function createBlindAuction(address assetOwner, address nftAddress, uint256 assetId, bytes32 blindAuctionId, uint256 biddingEnd, uint256 revealEnd, uint256 startPriceInWei) returns()
func (_MarketplaceStorage *MarketplaceStorageTransactor) CreateBlindAuction(opts *bind.TransactOpts, assetOwner common.Address, nftAddress common.Address, assetId *big.Int, blindAuctionId [32]byte, biddingEnd *big.Int, revealEnd *big.Int, startPriceInWei *big.Int) (*types.Transaction, error) {
	return _MarketplaceStorage.contract.Transact(opts, "createBlindAuction", assetOwner, nftAddress, assetId, blindAuctionId, biddingEnd, revealEnd, startPriceInWei)
}

// CreateBlindAuction is a paid mutator transaction binding the contract method 0xd7e4a307.
//
// Solidity: function createBlindAuction(address assetOwner, address nftAddress, uint256 assetId, bytes32 blindAuctionId, uint256 biddingEnd, uint256 revealEnd, uint256 startPriceInWei) returns()
func (_MarketplaceStorage *MarketplaceStorageSession) CreateBlindAuction(assetOwner common.Address, nftAddress common.Address, assetId *big.Int, blindAuctionId [32]byte, biddingEnd *big.Int, revealEnd *big.Int, startPriceInWei *big.Int) (*types.Transaction, error) {
	return _MarketplaceStorage.Contract.CreateBlindAuction(&_MarketplaceStorage.TransactOpts, assetOwner, nftAddress, assetId, blindAuctionId, biddingEnd, revealEnd, startPriceInWei)
}

// CreateBlindAuction is a paid mutator transaction binding the contract method 0xd7e4a307.
//
// Solidity: function createBlindAuction(address assetOwner, address nftAddress, uint256 assetId, bytes32 blindAuctionId, uint256 biddingEnd, uint256 revealEnd, uint256 startPriceInWei) returns()
func (_MarketplaceStorage *MarketplaceStorageTransactorSession) CreateBlindAuction(assetOwner common.Address, nftAddress common.Address, assetId *big.Int, blindAuctionId [32]byte, biddingEnd *big.Int, revealEnd *big.Int, startPriceInWei *big.Int) (*types.Transaction, error) {
	return _MarketplaceStorage.Contract.CreateBlindAuction(&_MarketplaceStorage.TransactOpts, assetOwner, nftAddress, assetId, blindAuctionId, biddingEnd, revealEnd, startPriceInWei)
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

// CreatePublicAuction is a paid mutator transaction binding the contract method 0x6af4c895.
//
// Solidity: function createPublicAuction(address assetOwner, address nftAddress, uint256 assetId, bytes32 publicAuctionId, uint256 biddingEnd, uint256 startPriceInWei, uint256 minIncrement) returns()
func (_MarketplaceStorage *MarketplaceStorageTransactor) CreatePublicAuction(opts *bind.TransactOpts, assetOwner common.Address, nftAddress common.Address, assetId *big.Int, publicAuctionId [32]byte, biddingEnd *big.Int, startPriceInWei *big.Int, minIncrement *big.Int) (*types.Transaction, error) {
	return _MarketplaceStorage.contract.Transact(opts, "createPublicAuction", assetOwner, nftAddress, assetId, publicAuctionId, biddingEnd, startPriceInWei, minIncrement)
}

// CreatePublicAuction is a paid mutator transaction binding the contract method 0x6af4c895.
//
// Solidity: function createPublicAuction(address assetOwner, address nftAddress, uint256 assetId, bytes32 publicAuctionId, uint256 biddingEnd, uint256 startPriceInWei, uint256 minIncrement) returns()
func (_MarketplaceStorage *MarketplaceStorageSession) CreatePublicAuction(assetOwner common.Address, nftAddress common.Address, assetId *big.Int, publicAuctionId [32]byte, biddingEnd *big.Int, startPriceInWei *big.Int, minIncrement *big.Int) (*types.Transaction, error) {
	return _MarketplaceStorage.Contract.CreatePublicAuction(&_MarketplaceStorage.TransactOpts, assetOwner, nftAddress, assetId, publicAuctionId, biddingEnd, startPriceInWei, minIncrement)
}

// CreatePublicAuction is a paid mutator transaction binding the contract method 0x6af4c895.
//
// Solidity: function createPublicAuction(address assetOwner, address nftAddress, uint256 assetId, bytes32 publicAuctionId, uint256 biddingEnd, uint256 startPriceInWei, uint256 minIncrement) returns()
func (_MarketplaceStorage *MarketplaceStorageTransactorSession) CreatePublicAuction(assetOwner common.Address, nftAddress common.Address, assetId *big.Int, publicAuctionId [32]byte, biddingEnd *big.Int, startPriceInWei *big.Int, minIncrement *big.Int) (*types.Transaction, error) {
	return _MarketplaceStorage.Contract.CreatePublicAuction(&_MarketplaceStorage.TransactOpts, assetOwner, nftAddress, assetId, publicAuctionId, biddingEnd, startPriceInWei, minIncrement)
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

// EndBlindAuction is a paid mutator transaction binding the contract method 0x15c97238.
//
// Solidity: function endBlindAuction(bytes32 nftAsset) returns()
func (_MarketplaceStorage *MarketplaceStorageTransactor) EndBlindAuction(opts *bind.TransactOpts, nftAsset [32]byte) (*types.Transaction, error) {
	return _MarketplaceStorage.contract.Transact(opts, "endBlindAuction", nftAsset)
}

// EndBlindAuction is a paid mutator transaction binding the contract method 0x15c97238.
//
// Solidity: function endBlindAuction(bytes32 nftAsset) returns()
func (_MarketplaceStorage *MarketplaceStorageSession) EndBlindAuction(nftAsset [32]byte) (*types.Transaction, error) {
	return _MarketplaceStorage.Contract.EndBlindAuction(&_MarketplaceStorage.TransactOpts, nftAsset)
}

// EndBlindAuction is a paid mutator transaction binding the contract method 0x15c97238.
//
// Solidity: function endBlindAuction(bytes32 nftAsset) returns()
func (_MarketplaceStorage *MarketplaceStorageTransactorSession) EndBlindAuction(nftAsset [32]byte) (*types.Transaction, error) {
	return _MarketplaceStorage.Contract.EndBlindAuction(&_MarketplaceStorage.TransactOpts, nftAsset)
}

// EndPublicAuction is a paid mutator transaction binding the contract method 0x5a0affbb.
//
// Solidity: function endPublicAuction(bytes32 nftAsset) returns()
func (_MarketplaceStorage *MarketplaceStorageTransactor) EndPublicAuction(opts *bind.TransactOpts, nftAsset [32]byte) (*types.Transaction, error) {
	return _MarketplaceStorage.contract.Transact(opts, "endPublicAuction", nftAsset)
}

// EndPublicAuction is a paid mutator transaction binding the contract method 0x5a0affbb.
//
// Solidity: function endPublicAuction(bytes32 nftAsset) returns()
func (_MarketplaceStorage *MarketplaceStorageSession) EndPublicAuction(nftAsset [32]byte) (*types.Transaction, error) {
	return _MarketplaceStorage.Contract.EndPublicAuction(&_MarketplaceStorage.TransactOpts, nftAsset)
}

// EndPublicAuction is a paid mutator transaction binding the contract method 0x5a0affbb.
//
// Solidity: function endPublicAuction(bytes32 nftAsset) returns()
func (_MarketplaceStorage *MarketplaceStorageTransactorSession) EndPublicAuction(nftAsset [32]byte) (*types.Transaction, error) {
	return _MarketplaceStorage.Contract.EndPublicAuction(&_MarketplaceStorage.TransactOpts, nftAsset)
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

// UpdateBlindAuctionMarketplace is a paid mutator transaction binding the contract method 0x51b7e09d.
//
// Solidity: function updateBlindAuctionMarketplace(address _blindAuctionMarketplace) returns()
func (_MarketplaceStorage *MarketplaceStorageTransactor) UpdateBlindAuctionMarketplace(opts *bind.TransactOpts, _blindAuctionMarketplace common.Address) (*types.Transaction, error) {
	return _MarketplaceStorage.contract.Transact(opts, "updateBlindAuctionMarketplace", _blindAuctionMarketplace)
}

// UpdateBlindAuctionMarketplace is a paid mutator transaction binding the contract method 0x51b7e09d.
//
// Solidity: function updateBlindAuctionMarketplace(address _blindAuctionMarketplace) returns()
func (_MarketplaceStorage *MarketplaceStorageSession) UpdateBlindAuctionMarketplace(_blindAuctionMarketplace common.Address) (*types.Transaction, error) {
	return _MarketplaceStorage.Contract.UpdateBlindAuctionMarketplace(&_MarketplaceStorage.TransactOpts, _blindAuctionMarketplace)
}

// UpdateBlindAuctionMarketplace is a paid mutator transaction binding the contract method 0x51b7e09d.
//
// Solidity: function updateBlindAuctionMarketplace(address _blindAuctionMarketplace) returns()
func (_MarketplaceStorage *MarketplaceStorageTransactorSession) UpdateBlindAuctionMarketplace(_blindAuctionMarketplace common.Address) (*types.Transaction, error) {
	return _MarketplaceStorage.Contract.UpdateBlindAuctionMarketplace(&_MarketplaceStorage.TransactOpts, _blindAuctionMarketplace)
}

// UpdateHighestBidBlindAuction is a paid mutator transaction binding the contract method 0xe7c847aa.
//
// Solidity: function updateHighestBidBlindAuction(address bidder, uint256 highestBid, bytes32 blindAuctionId) returns()
func (_MarketplaceStorage *MarketplaceStorageTransactor) UpdateHighestBidBlindAuction(opts *bind.TransactOpts, bidder common.Address, highestBid *big.Int, blindAuctionId [32]byte) (*types.Transaction, error) {
	return _MarketplaceStorage.contract.Transact(opts, "updateHighestBidBlindAuction", bidder, highestBid, blindAuctionId)
}

// UpdateHighestBidBlindAuction is a paid mutator transaction binding the contract method 0xe7c847aa.
//
// Solidity: function updateHighestBidBlindAuction(address bidder, uint256 highestBid, bytes32 blindAuctionId) returns()
func (_MarketplaceStorage *MarketplaceStorageSession) UpdateHighestBidBlindAuction(bidder common.Address, highestBid *big.Int, blindAuctionId [32]byte) (*types.Transaction, error) {
	return _MarketplaceStorage.Contract.UpdateHighestBidBlindAuction(&_MarketplaceStorage.TransactOpts, bidder, highestBid, blindAuctionId)
}

// UpdateHighestBidBlindAuction is a paid mutator transaction binding the contract method 0xe7c847aa.
//
// Solidity: function updateHighestBidBlindAuction(address bidder, uint256 highestBid, bytes32 blindAuctionId) returns()
func (_MarketplaceStorage *MarketplaceStorageTransactorSession) UpdateHighestBidBlindAuction(bidder common.Address, highestBid *big.Int, blindAuctionId [32]byte) (*types.Transaction, error) {
	return _MarketplaceStorage.Contract.UpdateHighestBidBlindAuction(&_MarketplaceStorage.TransactOpts, bidder, highestBid, blindAuctionId)
}

// UpdateHighestBidPublicAuction is a paid mutator transaction binding the contract method 0x9d90ce3b.
//
// Solidity: function updateHighestBidPublicAuction(address bidder, uint256 highestBid, bytes32 publicAuctionId) returns()
func (_MarketplaceStorage *MarketplaceStorageTransactor) UpdateHighestBidPublicAuction(opts *bind.TransactOpts, bidder common.Address, highestBid *big.Int, publicAuctionId [32]byte) (*types.Transaction, error) {
	return _MarketplaceStorage.contract.Transact(opts, "updateHighestBidPublicAuction", bidder, highestBid, publicAuctionId)
}

// UpdateHighestBidPublicAuction is a paid mutator transaction binding the contract method 0x9d90ce3b.
//
// Solidity: function updateHighestBidPublicAuction(address bidder, uint256 highestBid, bytes32 publicAuctionId) returns()
func (_MarketplaceStorage *MarketplaceStorageSession) UpdateHighestBidPublicAuction(bidder common.Address, highestBid *big.Int, publicAuctionId [32]byte) (*types.Transaction, error) {
	return _MarketplaceStorage.Contract.UpdateHighestBidPublicAuction(&_MarketplaceStorage.TransactOpts, bidder, highestBid, publicAuctionId)
}

// UpdateHighestBidPublicAuction is a paid mutator transaction binding the contract method 0x9d90ce3b.
//
// Solidity: function updateHighestBidPublicAuction(address bidder, uint256 highestBid, bytes32 publicAuctionId) returns()
func (_MarketplaceStorage *MarketplaceStorageTransactorSession) UpdateHighestBidPublicAuction(bidder common.Address, highestBid *big.Int, publicAuctionId [32]byte) (*types.Transaction, error) {
	return _MarketplaceStorage.Contract.UpdateHighestBidPublicAuction(&_MarketplaceStorage.TransactOpts, bidder, highestBid, publicAuctionId)
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

// UpdatePublicAuctionMarketplace is a paid mutator transaction binding the contract method 0xef2d8b55.
//
// Solidity: function updatePublicAuctionMarketplace(address _publicAuctionMarketplace) returns()
func (_MarketplaceStorage *MarketplaceStorageTransactor) UpdatePublicAuctionMarketplace(opts *bind.TransactOpts, _publicAuctionMarketplace common.Address) (*types.Transaction, error) {
	return _MarketplaceStorage.contract.Transact(opts, "updatePublicAuctionMarketplace", _publicAuctionMarketplace)
}

// UpdatePublicAuctionMarketplace is a paid mutator transaction binding the contract method 0xef2d8b55.
//
// Solidity: function updatePublicAuctionMarketplace(address _publicAuctionMarketplace) returns()
func (_MarketplaceStorage *MarketplaceStorageSession) UpdatePublicAuctionMarketplace(_publicAuctionMarketplace common.Address) (*types.Transaction, error) {
	return _MarketplaceStorage.Contract.UpdatePublicAuctionMarketplace(&_MarketplaceStorage.TransactOpts, _publicAuctionMarketplace)
}

// UpdatePublicAuctionMarketplace is a paid mutator transaction binding the contract method 0xef2d8b55.
//
// Solidity: function updatePublicAuctionMarketplace(address _publicAuctionMarketplace) returns()
func (_MarketplaceStorage *MarketplaceStorageTransactorSession) UpdatePublicAuctionMarketplace(_publicAuctionMarketplace common.Address) (*types.Transaction, error) {
	return _MarketplaceStorage.Contract.UpdatePublicAuctionMarketplace(&_MarketplaceStorage.TransactOpts, _publicAuctionMarketplace)
}

// MarketplaceStorageBlindAuctionBidSuccessfulIterator is returned from FilterBlindAuctionBidSuccessful and is used to iterate over the raw logs and unpacked data for BlindAuctionBidSuccessful events raised by the MarketplaceStorage contract.
type MarketplaceStorageBlindAuctionBidSuccessfulIterator struct {
	Event *MarketplaceStorageBlindAuctionBidSuccessful // Event containing the contract specifics and raw log

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
func (it *MarketplaceStorageBlindAuctionBidSuccessfulIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketplaceStorageBlindAuctionBidSuccessful)
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
		it.Event = new(MarketplaceStorageBlindAuctionBidSuccessful)
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
func (it *MarketplaceStorageBlindAuctionBidSuccessfulIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketplaceStorageBlindAuctionBidSuccessfulIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketplaceStorageBlindAuctionBidSuccessful represents a BlindAuctionBidSuccessful event raised by the MarketplaceStorage contract.
type MarketplaceStorageBlindAuctionBidSuccessful struct {
	Bidder         common.Address
	BlindAuctionId [32]byte
	BlindedBid     [32]byte
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterBlindAuctionBidSuccessful is a free log retrieval operation binding the contract event 0x42412bd3e1b349d0fc9c3518c85d09f611398d80842c5d3e36449b76dca01b6b.
//
// Solidity: event BlindAuctionBidSuccessful(address bidder, bytes32 blindAuctionId, bytes32 blindedBid)
func (_MarketplaceStorage *MarketplaceStorageFilterer) FilterBlindAuctionBidSuccessful(opts *bind.FilterOpts) (*MarketplaceStorageBlindAuctionBidSuccessfulIterator, error) {

	logs, sub, err := _MarketplaceStorage.contract.FilterLogs(opts, "BlindAuctionBidSuccessful")
	if err != nil {
		return nil, err
	}
	return &MarketplaceStorageBlindAuctionBidSuccessfulIterator{contract: _MarketplaceStorage.contract, event: "BlindAuctionBidSuccessful", logs: logs, sub: sub}, nil
}

// WatchBlindAuctionBidSuccessful is a free log subscription operation binding the contract event 0x42412bd3e1b349d0fc9c3518c85d09f611398d80842c5d3e36449b76dca01b6b.
//
// Solidity: event BlindAuctionBidSuccessful(address bidder, bytes32 blindAuctionId, bytes32 blindedBid)
func (_MarketplaceStorage *MarketplaceStorageFilterer) WatchBlindAuctionBidSuccessful(opts *bind.WatchOpts, sink chan<- *MarketplaceStorageBlindAuctionBidSuccessful) (event.Subscription, error) {

	logs, sub, err := _MarketplaceStorage.contract.WatchLogs(opts, "BlindAuctionBidSuccessful")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketplaceStorageBlindAuctionBidSuccessful)
				if err := _MarketplaceStorage.contract.UnpackLog(event, "BlindAuctionBidSuccessful", log); err != nil {
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
func (_MarketplaceStorage *MarketplaceStorageFilterer) ParseBlindAuctionBidSuccessful(log types.Log) (*MarketplaceStorageBlindAuctionBidSuccessful, error) {
	event := new(MarketplaceStorageBlindAuctionBidSuccessful)
	if err := _MarketplaceStorage.contract.UnpackLog(event, "BlindAuctionBidSuccessful", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MarketplaceStorageBlindAuctionCancelledIterator is returned from FilterBlindAuctionCancelled and is used to iterate over the raw logs and unpacked data for BlindAuctionCancelled events raised by the MarketplaceStorage contract.
type MarketplaceStorageBlindAuctionCancelledIterator struct {
	Event *MarketplaceStorageBlindAuctionCancelled // Event containing the contract specifics and raw log

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
func (it *MarketplaceStorageBlindAuctionCancelledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketplaceStorageBlindAuctionCancelled)
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
		it.Event = new(MarketplaceStorageBlindAuctionCancelled)
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
func (it *MarketplaceStorageBlindAuctionCancelledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketplaceStorageBlindAuctionCancelledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketplaceStorageBlindAuctionCancelled represents a BlindAuctionCancelled event raised by the MarketplaceStorage contract.
type MarketplaceStorageBlindAuctionCancelled struct {
	Canceller      common.Address
	BlindAuctionId [32]byte
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterBlindAuctionCancelled is a free log retrieval operation binding the contract event 0x620e9f21de284210359a265815fc84d3f6f4686df7a3841c7ef048a50ea80216.
//
// Solidity: event BlindAuctionCancelled(address canceller, bytes32 blindAuctionId)
func (_MarketplaceStorage *MarketplaceStorageFilterer) FilterBlindAuctionCancelled(opts *bind.FilterOpts) (*MarketplaceStorageBlindAuctionCancelledIterator, error) {

	logs, sub, err := _MarketplaceStorage.contract.FilterLogs(opts, "BlindAuctionCancelled")
	if err != nil {
		return nil, err
	}
	return &MarketplaceStorageBlindAuctionCancelledIterator{contract: _MarketplaceStorage.contract, event: "BlindAuctionCancelled", logs: logs, sub: sub}, nil
}

// WatchBlindAuctionCancelled is a free log subscription operation binding the contract event 0x620e9f21de284210359a265815fc84d3f6f4686df7a3841c7ef048a50ea80216.
//
// Solidity: event BlindAuctionCancelled(address canceller, bytes32 blindAuctionId)
func (_MarketplaceStorage *MarketplaceStorageFilterer) WatchBlindAuctionCancelled(opts *bind.WatchOpts, sink chan<- *MarketplaceStorageBlindAuctionCancelled) (event.Subscription, error) {

	logs, sub, err := _MarketplaceStorage.contract.WatchLogs(opts, "BlindAuctionCancelled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketplaceStorageBlindAuctionCancelled)
				if err := _MarketplaceStorage.contract.UnpackLog(event, "BlindAuctionCancelled", log); err != nil {
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
func (_MarketplaceStorage *MarketplaceStorageFilterer) ParseBlindAuctionCancelled(log types.Log) (*MarketplaceStorageBlindAuctionCancelled, error) {
	event := new(MarketplaceStorageBlindAuctionCancelled)
	if err := _MarketplaceStorage.contract.UnpackLog(event, "BlindAuctionCancelled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MarketplaceStorageBlindAuctionCreatedIterator is returned from FilterBlindAuctionCreated and is used to iterate over the raw logs and unpacked data for BlindAuctionCreated events raised by the MarketplaceStorage contract.
type MarketplaceStorageBlindAuctionCreatedIterator struct {
	Event *MarketplaceStorageBlindAuctionCreated // Event containing the contract specifics and raw log

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
func (it *MarketplaceStorageBlindAuctionCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketplaceStorageBlindAuctionCreated)
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
		it.Event = new(MarketplaceStorageBlindAuctionCreated)
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
func (it *MarketplaceStorageBlindAuctionCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketplaceStorageBlindAuctionCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketplaceStorageBlindAuctionCreated represents a BlindAuctionCreated event raised by the MarketplaceStorage contract.
type MarketplaceStorageBlindAuctionCreated struct {
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
func (_MarketplaceStorage *MarketplaceStorageFilterer) FilterBlindAuctionCreated(opts *bind.FilterOpts) (*MarketplaceStorageBlindAuctionCreatedIterator, error) {

	logs, sub, err := _MarketplaceStorage.contract.FilterLogs(opts, "BlindAuctionCreated")
	if err != nil {
		return nil, err
	}
	return &MarketplaceStorageBlindAuctionCreatedIterator{contract: _MarketplaceStorage.contract, event: "BlindAuctionCreated", logs: logs, sub: sub}, nil
}

// WatchBlindAuctionCreated is a free log subscription operation binding the contract event 0xe64fd4ba0702f46ae298a3be33f16f223a5e06556977c42dee952f976abfed83.
//
// Solidity: event BlindAuctionCreated(address seller, address nftAddress, bytes32 blindAuctionId, uint256 assetId, uint256 biddingEnd, uint256 revealEnd, uint256 startPriceInWei)
func (_MarketplaceStorage *MarketplaceStorageFilterer) WatchBlindAuctionCreated(opts *bind.WatchOpts, sink chan<- *MarketplaceStorageBlindAuctionCreated) (event.Subscription, error) {

	logs, sub, err := _MarketplaceStorage.contract.WatchLogs(opts, "BlindAuctionCreated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketplaceStorageBlindAuctionCreated)
				if err := _MarketplaceStorage.contract.UnpackLog(event, "BlindAuctionCreated", log); err != nil {
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
func (_MarketplaceStorage *MarketplaceStorageFilterer) ParseBlindAuctionCreated(log types.Log) (*MarketplaceStorageBlindAuctionCreated, error) {
	event := new(MarketplaceStorageBlindAuctionCreated)
	if err := _MarketplaceStorage.contract.UnpackLog(event, "BlindAuctionCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MarketplaceStorageBlindAuctionEndedIterator is returned from FilterBlindAuctionEnded and is used to iterate over the raw logs and unpacked data for BlindAuctionEnded events raised by the MarketplaceStorage contract.
type MarketplaceStorageBlindAuctionEndedIterator struct {
	Event *MarketplaceStorageBlindAuctionEnded // Event containing the contract specifics and raw log

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
func (it *MarketplaceStorageBlindAuctionEndedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketplaceStorageBlindAuctionEnded)
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
		it.Event = new(MarketplaceStorageBlindAuctionEnded)
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
func (it *MarketplaceStorageBlindAuctionEndedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketplaceStorageBlindAuctionEndedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketplaceStorageBlindAuctionEnded represents a BlindAuctionEnded event raised by the MarketplaceStorage contract.
type MarketplaceStorageBlindAuctionEnded struct {
	BlindAuctionId [32]byte
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterBlindAuctionEnded is a free log retrieval operation binding the contract event 0xbe3b74f797d49d96897f669a8b62fc0d77fa4cd6cf999a1b08bc31e7f4faf219.
//
// Solidity: event BlindAuctionEnded(bytes32 blindAuctionId)
func (_MarketplaceStorage *MarketplaceStorageFilterer) FilterBlindAuctionEnded(opts *bind.FilterOpts) (*MarketplaceStorageBlindAuctionEndedIterator, error) {

	logs, sub, err := _MarketplaceStorage.contract.FilterLogs(opts, "BlindAuctionEnded")
	if err != nil {
		return nil, err
	}
	return &MarketplaceStorageBlindAuctionEndedIterator{contract: _MarketplaceStorage.contract, event: "BlindAuctionEnded", logs: logs, sub: sub}, nil
}

// WatchBlindAuctionEnded is a free log subscription operation binding the contract event 0xbe3b74f797d49d96897f669a8b62fc0d77fa4cd6cf999a1b08bc31e7f4faf219.
//
// Solidity: event BlindAuctionEnded(bytes32 blindAuctionId)
func (_MarketplaceStorage *MarketplaceStorageFilterer) WatchBlindAuctionEnded(opts *bind.WatchOpts, sink chan<- *MarketplaceStorageBlindAuctionEnded) (event.Subscription, error) {

	logs, sub, err := _MarketplaceStorage.contract.WatchLogs(opts, "BlindAuctionEnded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketplaceStorageBlindAuctionEnded)
				if err := _MarketplaceStorage.contract.UnpackLog(event, "BlindAuctionEnded", log); err != nil {
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
func (_MarketplaceStorage *MarketplaceStorageFilterer) ParseBlindAuctionEnded(log types.Log) (*MarketplaceStorageBlindAuctionEnded, error) {
	event := new(MarketplaceStorageBlindAuctionEnded)
	if err := _MarketplaceStorage.contract.UnpackLog(event, "BlindAuctionEnded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MarketplaceStorageBlindAuctionRefundIterator is returned from FilterBlindAuctionRefund and is used to iterate over the raw logs and unpacked data for BlindAuctionRefund events raised by the MarketplaceStorage contract.
type MarketplaceStorageBlindAuctionRefundIterator struct {
	Event *MarketplaceStorageBlindAuctionRefund // Event containing the contract specifics and raw log

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
func (it *MarketplaceStorageBlindAuctionRefundIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketplaceStorageBlindAuctionRefund)
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
		it.Event = new(MarketplaceStorageBlindAuctionRefund)
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
func (it *MarketplaceStorageBlindAuctionRefundIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketplaceStorageBlindAuctionRefundIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketplaceStorageBlindAuctionRefund represents a BlindAuctionRefund event raised by the MarketplaceStorage contract.
type MarketplaceStorageBlindAuctionRefund struct {
	Bidder         common.Address
	BlindAuctionId [32]byte
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterBlindAuctionRefund is a free log retrieval operation binding the contract event 0xadada2afe0e9306e1b93faf56e74b041719e68f722de4f550d5a57e39610b404.
//
// Solidity: event BlindAuctionRefund(address bidder, bytes32 blindAuctionId)
func (_MarketplaceStorage *MarketplaceStorageFilterer) FilterBlindAuctionRefund(opts *bind.FilterOpts) (*MarketplaceStorageBlindAuctionRefundIterator, error) {

	logs, sub, err := _MarketplaceStorage.contract.FilterLogs(opts, "BlindAuctionRefund")
	if err != nil {
		return nil, err
	}
	return &MarketplaceStorageBlindAuctionRefundIterator{contract: _MarketplaceStorage.contract, event: "BlindAuctionRefund", logs: logs, sub: sub}, nil
}

// WatchBlindAuctionRefund is a free log subscription operation binding the contract event 0xadada2afe0e9306e1b93faf56e74b041719e68f722de4f550d5a57e39610b404.
//
// Solidity: event BlindAuctionRefund(address bidder, bytes32 blindAuctionId)
func (_MarketplaceStorage *MarketplaceStorageFilterer) WatchBlindAuctionRefund(opts *bind.WatchOpts, sink chan<- *MarketplaceStorageBlindAuctionRefund) (event.Subscription, error) {

	logs, sub, err := _MarketplaceStorage.contract.WatchLogs(opts, "BlindAuctionRefund")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketplaceStorageBlindAuctionRefund)
				if err := _MarketplaceStorage.contract.UnpackLog(event, "BlindAuctionRefund", log); err != nil {
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
func (_MarketplaceStorage *MarketplaceStorageFilterer) ParseBlindAuctionRefund(log types.Log) (*MarketplaceStorageBlindAuctionRefund, error) {
	event := new(MarketplaceStorageBlindAuctionRefund)
	if err := _MarketplaceStorage.contract.UnpackLog(event, "BlindAuctionRefund", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MarketplaceStorageBlindAuctionSuccessfulIterator is returned from FilterBlindAuctionSuccessful and is used to iterate over the raw logs and unpacked data for BlindAuctionSuccessful events raised by the MarketplaceStorage contract.
type MarketplaceStorageBlindAuctionSuccessfulIterator struct {
	Event *MarketplaceStorageBlindAuctionSuccessful // Event containing the contract specifics and raw log

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
func (it *MarketplaceStorageBlindAuctionSuccessfulIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketplaceStorageBlindAuctionSuccessful)
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
		it.Event = new(MarketplaceStorageBlindAuctionSuccessful)
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
func (it *MarketplaceStorageBlindAuctionSuccessfulIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketplaceStorageBlindAuctionSuccessfulIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketplaceStorageBlindAuctionSuccessful represents a BlindAuctionSuccessful event raised by the MarketplaceStorage contract.
type MarketplaceStorageBlindAuctionSuccessful struct {
	Seller         common.Address
	Buyer          common.Address
	BlindAuctionId [32]byte
	TotalPrice     *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterBlindAuctionSuccessful is a free log retrieval operation binding the contract event 0x2258179acb0769137f4ff532afd72c91199ae597abded9936acb61a2e66fc142.
//
// Solidity: event BlindAuctionSuccessful(address seller, address buyer, bytes32 blindAuctionId, uint256 totalPrice)
func (_MarketplaceStorage *MarketplaceStorageFilterer) FilterBlindAuctionSuccessful(opts *bind.FilterOpts) (*MarketplaceStorageBlindAuctionSuccessfulIterator, error) {

	logs, sub, err := _MarketplaceStorage.contract.FilterLogs(opts, "BlindAuctionSuccessful")
	if err != nil {
		return nil, err
	}
	return &MarketplaceStorageBlindAuctionSuccessfulIterator{contract: _MarketplaceStorage.contract, event: "BlindAuctionSuccessful", logs: logs, sub: sub}, nil
}

// WatchBlindAuctionSuccessful is a free log subscription operation binding the contract event 0x2258179acb0769137f4ff532afd72c91199ae597abded9936acb61a2e66fc142.
//
// Solidity: event BlindAuctionSuccessful(address seller, address buyer, bytes32 blindAuctionId, uint256 totalPrice)
func (_MarketplaceStorage *MarketplaceStorageFilterer) WatchBlindAuctionSuccessful(opts *bind.WatchOpts, sink chan<- *MarketplaceStorageBlindAuctionSuccessful) (event.Subscription, error) {

	logs, sub, err := _MarketplaceStorage.contract.WatchLogs(opts, "BlindAuctionSuccessful")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketplaceStorageBlindAuctionSuccessful)
				if err := _MarketplaceStorage.contract.UnpackLog(event, "BlindAuctionSuccessful", log); err != nil {
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
func (_MarketplaceStorage *MarketplaceStorageFilterer) ParseBlindAuctionSuccessful(log types.Log) (*MarketplaceStorageBlindAuctionSuccessful, error) {
	event := new(MarketplaceStorageBlindAuctionSuccessful)
	if err := _MarketplaceStorage.contract.UnpackLog(event, "BlindAuctionSuccessful", log); err != nil {
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

// MarketplaceStoragePublicAuctionBidSuccessfulIterator is returned from FilterPublicAuctionBidSuccessful and is used to iterate over the raw logs and unpacked data for PublicAuctionBidSuccessful events raised by the MarketplaceStorage contract.
type MarketplaceStoragePublicAuctionBidSuccessfulIterator struct {
	Event *MarketplaceStoragePublicAuctionBidSuccessful // Event containing the contract specifics and raw log

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
func (it *MarketplaceStoragePublicAuctionBidSuccessfulIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketplaceStoragePublicAuctionBidSuccessful)
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
		it.Event = new(MarketplaceStoragePublicAuctionBidSuccessful)
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
func (it *MarketplaceStoragePublicAuctionBidSuccessfulIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketplaceStoragePublicAuctionBidSuccessfulIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketplaceStoragePublicAuctionBidSuccessful represents a PublicAuctionBidSuccessful event raised by the MarketplaceStorage contract.
type MarketplaceStoragePublicAuctionBidSuccessful struct {
	Bidder          common.Address
	PublicAuctionId [32]byte
	BidValue        *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterPublicAuctionBidSuccessful is a free log retrieval operation binding the contract event 0x9f541988b3aac74ac43a6f4b330ceed668244b74461b90850d199a38a172f7c2.
//
// Solidity: event PublicAuctionBidSuccessful(address bidder, bytes32 publicAuctionId, uint256 bidValue)
func (_MarketplaceStorage *MarketplaceStorageFilterer) FilterPublicAuctionBidSuccessful(opts *bind.FilterOpts) (*MarketplaceStoragePublicAuctionBidSuccessfulIterator, error) {

	logs, sub, err := _MarketplaceStorage.contract.FilterLogs(opts, "PublicAuctionBidSuccessful")
	if err != nil {
		return nil, err
	}
	return &MarketplaceStoragePublicAuctionBidSuccessfulIterator{contract: _MarketplaceStorage.contract, event: "PublicAuctionBidSuccessful", logs: logs, sub: sub}, nil
}

// WatchPublicAuctionBidSuccessful is a free log subscription operation binding the contract event 0x9f541988b3aac74ac43a6f4b330ceed668244b74461b90850d199a38a172f7c2.
//
// Solidity: event PublicAuctionBidSuccessful(address bidder, bytes32 publicAuctionId, uint256 bidValue)
func (_MarketplaceStorage *MarketplaceStorageFilterer) WatchPublicAuctionBidSuccessful(opts *bind.WatchOpts, sink chan<- *MarketplaceStoragePublicAuctionBidSuccessful) (event.Subscription, error) {

	logs, sub, err := _MarketplaceStorage.contract.WatchLogs(opts, "PublicAuctionBidSuccessful")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketplaceStoragePublicAuctionBidSuccessful)
				if err := _MarketplaceStorage.contract.UnpackLog(event, "PublicAuctionBidSuccessful", log); err != nil {
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
func (_MarketplaceStorage *MarketplaceStorageFilterer) ParsePublicAuctionBidSuccessful(log types.Log) (*MarketplaceStoragePublicAuctionBidSuccessful, error) {
	event := new(MarketplaceStoragePublicAuctionBidSuccessful)
	if err := _MarketplaceStorage.contract.UnpackLog(event, "PublicAuctionBidSuccessful", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MarketplaceStoragePublicAuctionCancelledIterator is returned from FilterPublicAuctionCancelled and is used to iterate over the raw logs and unpacked data for PublicAuctionCancelled events raised by the MarketplaceStorage contract.
type MarketplaceStoragePublicAuctionCancelledIterator struct {
	Event *MarketplaceStoragePublicAuctionCancelled // Event containing the contract specifics and raw log

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
func (it *MarketplaceStoragePublicAuctionCancelledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketplaceStoragePublicAuctionCancelled)
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
		it.Event = new(MarketplaceStoragePublicAuctionCancelled)
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
func (it *MarketplaceStoragePublicAuctionCancelledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketplaceStoragePublicAuctionCancelledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketplaceStoragePublicAuctionCancelled represents a PublicAuctionCancelled event raised by the MarketplaceStorage contract.
type MarketplaceStoragePublicAuctionCancelled struct {
	Canceller       common.Address
	PublicAuctionId [32]byte
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterPublicAuctionCancelled is a free log retrieval operation binding the contract event 0x87cbb8dceaaacfbe7c9d99abf2a9c85f0249d860e329f60869079b204c7ad00d.
//
// Solidity: event PublicAuctionCancelled(address canceller, bytes32 publicAuctionId)
func (_MarketplaceStorage *MarketplaceStorageFilterer) FilterPublicAuctionCancelled(opts *bind.FilterOpts) (*MarketplaceStoragePublicAuctionCancelledIterator, error) {

	logs, sub, err := _MarketplaceStorage.contract.FilterLogs(opts, "PublicAuctionCancelled")
	if err != nil {
		return nil, err
	}
	return &MarketplaceStoragePublicAuctionCancelledIterator{contract: _MarketplaceStorage.contract, event: "PublicAuctionCancelled", logs: logs, sub: sub}, nil
}

// WatchPublicAuctionCancelled is a free log subscription operation binding the contract event 0x87cbb8dceaaacfbe7c9d99abf2a9c85f0249d860e329f60869079b204c7ad00d.
//
// Solidity: event PublicAuctionCancelled(address canceller, bytes32 publicAuctionId)
func (_MarketplaceStorage *MarketplaceStorageFilterer) WatchPublicAuctionCancelled(opts *bind.WatchOpts, sink chan<- *MarketplaceStoragePublicAuctionCancelled) (event.Subscription, error) {

	logs, sub, err := _MarketplaceStorage.contract.WatchLogs(opts, "PublicAuctionCancelled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketplaceStoragePublicAuctionCancelled)
				if err := _MarketplaceStorage.contract.UnpackLog(event, "PublicAuctionCancelled", log); err != nil {
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
func (_MarketplaceStorage *MarketplaceStorageFilterer) ParsePublicAuctionCancelled(log types.Log) (*MarketplaceStoragePublicAuctionCancelled, error) {
	event := new(MarketplaceStoragePublicAuctionCancelled)
	if err := _MarketplaceStorage.contract.UnpackLog(event, "PublicAuctionCancelled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MarketplaceStoragePublicAuctionCreatedIterator is returned from FilterPublicAuctionCreated and is used to iterate over the raw logs and unpacked data for PublicAuctionCreated events raised by the MarketplaceStorage contract.
type MarketplaceStoragePublicAuctionCreatedIterator struct {
	Event *MarketplaceStoragePublicAuctionCreated // Event containing the contract specifics and raw log

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
func (it *MarketplaceStoragePublicAuctionCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketplaceStoragePublicAuctionCreated)
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
		it.Event = new(MarketplaceStoragePublicAuctionCreated)
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
func (it *MarketplaceStoragePublicAuctionCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketplaceStoragePublicAuctionCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketplaceStoragePublicAuctionCreated represents a PublicAuctionCreated event raised by the MarketplaceStorage contract.
type MarketplaceStoragePublicAuctionCreated struct {
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
func (_MarketplaceStorage *MarketplaceStorageFilterer) FilterPublicAuctionCreated(opts *bind.FilterOpts) (*MarketplaceStoragePublicAuctionCreatedIterator, error) {

	logs, sub, err := _MarketplaceStorage.contract.FilterLogs(opts, "PublicAuctionCreated")
	if err != nil {
		return nil, err
	}
	return &MarketplaceStoragePublicAuctionCreatedIterator{contract: _MarketplaceStorage.contract, event: "PublicAuctionCreated", logs: logs, sub: sub}, nil
}

// WatchPublicAuctionCreated is a free log subscription operation binding the contract event 0xbdca6148a24d8e6e2d4ced0a6e168095869e61ea26b82332620abe8ba3ba5bd9.
//
// Solidity: event PublicAuctionCreated(address seller, address nftAddress, bytes32 publicAuctionId, uint256 assetId, uint256 biddingEnd, uint256 startPriceInWei, uint256 minIncrement)
func (_MarketplaceStorage *MarketplaceStorageFilterer) WatchPublicAuctionCreated(opts *bind.WatchOpts, sink chan<- *MarketplaceStoragePublicAuctionCreated) (event.Subscription, error) {

	logs, sub, err := _MarketplaceStorage.contract.WatchLogs(opts, "PublicAuctionCreated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketplaceStoragePublicAuctionCreated)
				if err := _MarketplaceStorage.contract.UnpackLog(event, "PublicAuctionCreated", log); err != nil {
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
func (_MarketplaceStorage *MarketplaceStorageFilterer) ParsePublicAuctionCreated(log types.Log) (*MarketplaceStoragePublicAuctionCreated, error) {
	event := new(MarketplaceStoragePublicAuctionCreated)
	if err := _MarketplaceStorage.contract.UnpackLog(event, "PublicAuctionCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MarketplaceStoragePublicAuctionEndedIterator is returned from FilterPublicAuctionEnded and is used to iterate over the raw logs and unpacked data for PublicAuctionEnded events raised by the MarketplaceStorage contract.
type MarketplaceStoragePublicAuctionEndedIterator struct {
	Event *MarketplaceStoragePublicAuctionEnded // Event containing the contract specifics and raw log

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
func (it *MarketplaceStoragePublicAuctionEndedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketplaceStoragePublicAuctionEnded)
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
		it.Event = new(MarketplaceStoragePublicAuctionEnded)
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
func (it *MarketplaceStoragePublicAuctionEndedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketplaceStoragePublicAuctionEndedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketplaceStoragePublicAuctionEnded represents a PublicAuctionEnded event raised by the MarketplaceStorage contract.
type MarketplaceStoragePublicAuctionEnded struct {
	PublicAuctionId [32]byte
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterPublicAuctionEnded is a free log retrieval operation binding the contract event 0x156c2754a62667a51625de4ceb04df6640d97d06de8c89bd0bbb33307f144e42.
//
// Solidity: event PublicAuctionEnded(bytes32 publicAuctionId)
func (_MarketplaceStorage *MarketplaceStorageFilterer) FilterPublicAuctionEnded(opts *bind.FilterOpts) (*MarketplaceStoragePublicAuctionEndedIterator, error) {

	logs, sub, err := _MarketplaceStorage.contract.FilterLogs(opts, "PublicAuctionEnded")
	if err != nil {
		return nil, err
	}
	return &MarketplaceStoragePublicAuctionEndedIterator{contract: _MarketplaceStorage.contract, event: "PublicAuctionEnded", logs: logs, sub: sub}, nil
}

// WatchPublicAuctionEnded is a free log subscription operation binding the contract event 0x156c2754a62667a51625de4ceb04df6640d97d06de8c89bd0bbb33307f144e42.
//
// Solidity: event PublicAuctionEnded(bytes32 publicAuctionId)
func (_MarketplaceStorage *MarketplaceStorageFilterer) WatchPublicAuctionEnded(opts *bind.WatchOpts, sink chan<- *MarketplaceStoragePublicAuctionEnded) (event.Subscription, error) {

	logs, sub, err := _MarketplaceStorage.contract.WatchLogs(opts, "PublicAuctionEnded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketplaceStoragePublicAuctionEnded)
				if err := _MarketplaceStorage.contract.UnpackLog(event, "PublicAuctionEnded", log); err != nil {
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
func (_MarketplaceStorage *MarketplaceStorageFilterer) ParsePublicAuctionEnded(log types.Log) (*MarketplaceStoragePublicAuctionEnded, error) {
	event := new(MarketplaceStoragePublicAuctionEnded)
	if err := _MarketplaceStorage.contract.UnpackLog(event, "PublicAuctionEnded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MarketplaceStoragePublicAuctionRefundIterator is returned from FilterPublicAuctionRefund and is used to iterate over the raw logs and unpacked data for PublicAuctionRefund events raised by the MarketplaceStorage contract.
type MarketplaceStoragePublicAuctionRefundIterator struct {
	Event *MarketplaceStoragePublicAuctionRefund // Event containing the contract specifics and raw log

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
func (it *MarketplaceStoragePublicAuctionRefundIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketplaceStoragePublicAuctionRefund)
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
		it.Event = new(MarketplaceStoragePublicAuctionRefund)
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
func (it *MarketplaceStoragePublicAuctionRefundIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketplaceStoragePublicAuctionRefundIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketplaceStoragePublicAuctionRefund represents a PublicAuctionRefund event raised by the MarketplaceStorage contract.
type MarketplaceStoragePublicAuctionRefund struct {
	Bidder          common.Address
	PublicAuctionId [32]byte
	Value           *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterPublicAuctionRefund is a free log retrieval operation binding the contract event 0x4fac96064e645402e0d4854b9549caba0169e3252c0e08a305eb60f498c88911.
//
// Solidity: event PublicAuctionRefund(address bidder, bytes32 publicAuctionId, uint256 value)
func (_MarketplaceStorage *MarketplaceStorageFilterer) FilterPublicAuctionRefund(opts *bind.FilterOpts) (*MarketplaceStoragePublicAuctionRefundIterator, error) {

	logs, sub, err := _MarketplaceStorage.contract.FilterLogs(opts, "PublicAuctionRefund")
	if err != nil {
		return nil, err
	}
	return &MarketplaceStoragePublicAuctionRefundIterator{contract: _MarketplaceStorage.contract, event: "PublicAuctionRefund", logs: logs, sub: sub}, nil
}

// WatchPublicAuctionRefund is a free log subscription operation binding the contract event 0x4fac96064e645402e0d4854b9549caba0169e3252c0e08a305eb60f498c88911.
//
// Solidity: event PublicAuctionRefund(address bidder, bytes32 publicAuctionId, uint256 value)
func (_MarketplaceStorage *MarketplaceStorageFilterer) WatchPublicAuctionRefund(opts *bind.WatchOpts, sink chan<- *MarketplaceStoragePublicAuctionRefund) (event.Subscription, error) {

	logs, sub, err := _MarketplaceStorage.contract.WatchLogs(opts, "PublicAuctionRefund")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketplaceStoragePublicAuctionRefund)
				if err := _MarketplaceStorage.contract.UnpackLog(event, "PublicAuctionRefund", log); err != nil {
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
func (_MarketplaceStorage *MarketplaceStorageFilterer) ParsePublicAuctionRefund(log types.Log) (*MarketplaceStoragePublicAuctionRefund, error) {
	event := new(MarketplaceStoragePublicAuctionRefund)
	if err := _MarketplaceStorage.contract.UnpackLog(event, "PublicAuctionRefund", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MarketplaceStoragePublicAuctionSuccessfulIterator is returned from FilterPublicAuctionSuccessful and is used to iterate over the raw logs and unpacked data for PublicAuctionSuccessful events raised by the MarketplaceStorage contract.
type MarketplaceStoragePublicAuctionSuccessfulIterator struct {
	Event *MarketplaceStoragePublicAuctionSuccessful // Event containing the contract specifics and raw log

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
func (it *MarketplaceStoragePublicAuctionSuccessfulIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketplaceStoragePublicAuctionSuccessful)
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
		it.Event = new(MarketplaceStoragePublicAuctionSuccessful)
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
func (it *MarketplaceStoragePublicAuctionSuccessfulIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketplaceStoragePublicAuctionSuccessfulIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketplaceStoragePublicAuctionSuccessful represents a PublicAuctionSuccessful event raised by the MarketplaceStorage contract.
type MarketplaceStoragePublicAuctionSuccessful struct {
	Seller          common.Address
	Buyer           common.Address
	PublicAuctionId [32]byte
	TotalPrice      *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterPublicAuctionSuccessful is a free log retrieval operation binding the contract event 0xd5133f9e5285bd0635ccfe53ae8978b96ef34a7bf16a0a696a411dc669322cb6.
//
// Solidity: event PublicAuctionSuccessful(address seller, address buyer, bytes32 publicAuctionId, uint256 totalPrice)
func (_MarketplaceStorage *MarketplaceStorageFilterer) FilterPublicAuctionSuccessful(opts *bind.FilterOpts) (*MarketplaceStoragePublicAuctionSuccessfulIterator, error) {

	logs, sub, err := _MarketplaceStorage.contract.FilterLogs(opts, "PublicAuctionSuccessful")
	if err != nil {
		return nil, err
	}
	return &MarketplaceStoragePublicAuctionSuccessfulIterator{contract: _MarketplaceStorage.contract, event: "PublicAuctionSuccessful", logs: logs, sub: sub}, nil
}

// WatchPublicAuctionSuccessful is a free log subscription operation binding the contract event 0xd5133f9e5285bd0635ccfe53ae8978b96ef34a7bf16a0a696a411dc669322cb6.
//
// Solidity: event PublicAuctionSuccessful(address seller, address buyer, bytes32 publicAuctionId, uint256 totalPrice)
func (_MarketplaceStorage *MarketplaceStorageFilterer) WatchPublicAuctionSuccessful(opts *bind.WatchOpts, sink chan<- *MarketplaceStoragePublicAuctionSuccessful) (event.Subscription, error) {

	logs, sub, err := _MarketplaceStorage.contract.WatchLogs(opts, "PublicAuctionSuccessful")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketplaceStoragePublicAuctionSuccessful)
				if err := _MarketplaceStorage.contract.UnpackLog(event, "PublicAuctionSuccessful", log); err != nil {
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
func (_MarketplaceStorage *MarketplaceStorageFilterer) ParsePublicAuctionSuccessful(log types.Log) (*MarketplaceStoragePublicAuctionSuccessful, error) {
	event := new(MarketplaceStoragePublicAuctionSuccessful)
	if err := _MarketplaceStorage.contract.UnpackLog(event, "PublicAuctionSuccessful", log); err != nil {
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
	Seller         common.Address
	BlindAuctionId [32]byte
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterRevealSuccessful is a free log retrieval operation binding the contract event 0x61d4ab60470c97b371c5f91be47873283386bb3af4de728e8795ef5318e013d6.
//
// Solidity: event RevealSuccessful(address seller, bytes32 blindAuctionId)
func (_MarketplaceStorage *MarketplaceStorageFilterer) FilterRevealSuccessful(opts *bind.FilterOpts) (*MarketplaceStorageRevealSuccessfulIterator, error) {

	logs, sub, err := _MarketplaceStorage.contract.FilterLogs(opts, "RevealSuccessful")
	if err != nil {
		return nil, err
	}
	return &MarketplaceStorageRevealSuccessfulIterator{contract: _MarketplaceStorage.contract, event: "RevealSuccessful", logs: logs, sub: sub}, nil
}

// WatchRevealSuccessful is a free log subscription operation binding the contract event 0x61d4ab60470c97b371c5f91be47873283386bb3af4de728e8795ef5318e013d6.
//
// Solidity: event RevealSuccessful(address seller, bytes32 blindAuctionId)
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

// ParseRevealSuccessful is a log parse operation binding the contract event 0x61d4ab60470c97b371c5f91be47873283386bb3af4de728e8795ef5318e013d6.
//
// Solidity: event RevealSuccessful(address seller, bytes32 blindAuctionId)
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
