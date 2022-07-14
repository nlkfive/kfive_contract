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
	StartTime     *big.Int
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
	StartTime     *big.Int
	MinIncrement  *big.Int
}

// MarketplaceStorageMetaData contains all meta data concerning the MarketplaceStorage contract.
var MarketplaceStorageMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"AssetNotExisted\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AssetUnvailable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AuctionEnded\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidBiddingPrice\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidExpiredTime\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidMkpSender\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidReveal\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidTime\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidWithdraw\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotBidYet\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotRunning\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotStarted\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotWinner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OrderExpired\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RewardGranted\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"TooEarly\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"TooLate\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"canceller\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"auctionId\",\"type\":\"bytes32\"}],\"name\":\"AuctionCancelledSuccessful\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"bidder\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"auctionId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"AuctionRefundSuccessful\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"blindAuctionId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"blindedBid\",\"type\":\"bytes32\"}],\"name\":\"BlindAuctionBidSuccessful\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"assetOwner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"nftAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"blindAuctionId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"assetId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"biddingEnd\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"revealEnd\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"startPriceInWei\",\"type\":\"uint256\"}],\"name\":\"BlindAuctionCreatedSuccessful\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"auctionHighestBidder\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"auctionId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"assetId\",\"type\":\"uint256\"}],\"name\":\"GrantAuctionRewardSuccessful\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"who\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"}],\"name\":\"OrderCancelled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"orderId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"assetId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"nftAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"priceInWei\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"expiredAt\",\"type\":\"uint256\"}],\"name\":\"OrderCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"buyer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"}],\"name\":\"OrderSuccessful\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"publicAuctionId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"bidValue\",\"type\":\"uint256\"}],\"name\":\"PublicAuctionBidSuccessful\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"assetOwner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"nftAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"publicAuctionId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"assetId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"biddingEnd\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"startPriceInWei\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"minIncrement\",\"type\":\"uint256\"}],\"name\":\"PublicAuctionCreatedSuccessful\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"blindAuctionId\",\"type\":\"bytes32\"}],\"name\":\"RevealSuccessful\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PAUSER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getRoleMember\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleMemberCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"orderByNftAsset\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"orderId\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"nftAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expiredAt\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_publicAuctionMarketplace\",\"type\":\"address\"}],\"name\":\"updatePublicAuctionMarketplace\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_blindAuctionMarketplace\",\"type\":\"address\"}],\"name\":\"updateBlindAuctionMarketplace\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_orderMarketplace\",\"type\":\"address\"}],\"name\":\"updateOrderMarketplace\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"nftAsset\",\"type\":\"bytes32\"}],\"name\":\"assetIsAvailable\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"publicAuctionId\",\"type\":\"bytes32\"}],\"name\":\"publicAuctionIsExisted\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"nftAsset\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"publicAuctionId\",\"type\":\"bytes32\"}],\"name\":\"publicAuctionIsEnded\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"nftAsset\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"publicAuctionId\",\"type\":\"bytes32\"}],\"name\":\"publicAuctionIsRunning\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"assetOwner\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"nftAsset\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"publicAuctionId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"biddingEnd\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startPriceInWei\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minIncrement\",\"type\":\"uint256\"}],\"name\":\"createPublicAuction\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"publicAuctionId\",\"type\":\"bytes32\"}],\"name\":\"getPublicAuction\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"highestBid\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"highestBidder\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"biddingEnd\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minIncrement\",\"type\":\"uint256\"}],\"internalType\":\"structIPublicAuction.PublicAuction\",\"name\":\"publicAuction\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"nftAsset\",\"type\":\"bytes32\"}],\"name\":\"endPublicAuction\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"bidder\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"highestBid\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"publicAuctionId\",\"type\":\"bytes32\"}],\"name\":\"updateHighestBidPublicAuction\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"blindAuctionId\",\"type\":\"bytes32\"}],\"name\":\"blindAuctionIsExisted\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"nftAsset\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"blindAuctionId\",\"type\":\"bytes32\"}],\"name\":\"blindAuctionIsEnded\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"nftAsset\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"blindAuctionId\",\"type\":\"bytes32\"}],\"name\":\"blindAuctionIsRunning\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"assetOwner\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"nftAsset\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"blindAuctionId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"biddingEnd\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"revealEnd\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startPriceInWei\",\"type\":\"uint256\"}],\"name\":\"createBlindAuction\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"blindAuctionId\",\"type\":\"bytes32\"}],\"name\":\"getBlindAuction\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"biddingEnd\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"revealEnd\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"highestBid\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"highestBidder\",\"type\":\"address\"}],\"internalType\":\"structIBlindAuction.BlindAuction\",\"name\":\"blindAuction\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"nftAsset\",\"type\":\"bytes32\"}],\"name\":\"endBlindAuction\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"bidder\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"highestBid\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"blindAuctionId\",\"type\":\"bytes32\"}],\"name\":\"updateHighestBidBlindAuction\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"nftAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"assetId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"orderId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"priceInWei\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expiredAt\",\"type\":\"uint256\"}],\"name\":\"createOrder\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"nftAsset\",\"type\":\"bytes32\"}],\"name\":\"getOrder\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"orderId\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"nftAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expiredAt\",\"type\":\"uint256\"}],\"internalType\":\"structIOrder.Order\",\"name\":\"order\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"nftAsset\",\"type\":\"bytes32\"}],\"name\":\"deleteOrder\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60806040523480156200001157600080fd5b5062000032620000266200010a60201b60201c565b6200011260201b60201c565b6000600260146101000a81548160ff021916908315150217905550620000716000801b620000656200010a60201b60201c565b620001d860201b60201c565b620000b27fa49807205ce4d355092ef5a8a18f56e8913cf4a201fbe287825b095693c21775620000a66200010a60201b60201c565b620001d860201b60201c565b620001047f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a7fa49807205ce4d355092ef5a8a18f56e8913cf4a201fbe287825b095693c21775620001ee60201b60201c565b620005b6565b600033905090565b6000600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905081600260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b620001ea82826200025160201b60201c565b5050565b600062000201836200031f60201b60201c565b905081600080858152602001908152602001600020600101819055508181847fbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff60405160405180910390a4505050565b6200026882826200033e60201b62001fe21760201c565b7fa49807205ce4d355092ef5a8a18f56e8913cf4a201fbe287825b095693c217758214806200029957506000801b82145b156200031b57620002d67f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a826200033e60201b62001fe21760201c565b6000801b8214156200031a57620003197fa49807205ce4d355092ef5a8a18f56e8913cf4a201fbe287825b095693c21775826200033e60201b62001fe21760201c565b5b5b5050565b6000806000838152602001908152602001600020600101549050919050565b6200035582826200038660201b620020161760201c565b6200038181600160008581526020019081526020016000206200047760201b620020f61790919060201c565b505050565b620003988282620004af60201b60201c565b6200047357600160008084815260200190815260200160002060000160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff021916908315150217905550620004186200010a60201b60201c565b73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16837f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a45b5050565b6000620004a7836000018373ffffffffffffffffffffffffffffffffffffffff1660001b6200051960201b60201c565b905092915050565b600080600084815260200190815260200160002060000160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16905092915050565b60006200052d83836200059360201b60201c565b620005885782600001829080600181540180825580915050600190039060005260206000200160009091909190915055826000018054905083600101600084815260200190815260200160002081905550600190506200058d565b600090505b92915050565b600080836001016000848152602001908152602001600020541415905092915050565b613d9d80620005c66000396000f3fe608060405234801561001057600080fd5b506004361061023c5760003560e01c80638456cb591161013b578063ae97d559116100b8578063deaf324a1161007c578063deaf324a1461070f578063e63ab1e91461072b578063e7c847aa14610749578063ef2d8b5514610765578063f2fde38b146107815761023c565b8063ae97d55914610647578063bc4124c514610663578063ca15c87314610693578063cd23b172146106c3578063d547741f146106f35761023c565b806391d14854116100ff57806391d14854146105a5578063930ca094146105d55780639a91d9c0146105f15780639d90ce3b1461060d578063a217fddf146106295761023c565b80638456cb5914610501578063866244481461050b57806387a61cbd1461053b5780638da5cb5b146105575780639010d07c146105755761023c565b80634a43c5b9116101c95780635c975abb1161018d5780635c975abb1461045b5780636ed3be5314610479578063715018a6146104a95780637584f03d146104b357806375b238fc146104e35761023c565b80634a43c5b91461039357806351b7e09d146103c35780635778472a146103df57806357a45b771461040f5780635a0affbb1461043f5761023c565b8063248a9ca311610210578063248a9ca3146102ed5780632f2ff15d1461031d57806336568abe146103395780633a3999d9146103555780633f4ba83a146103895761023c565b80622cb6a31461024157806301ffc9a7146102715780630a7f8977146102a157806315c97238146102d1575b600080fd5b61025b60048036038101906102569190613256565b61079d565b604051610268919061373d565b60405180910390f35b61028b600480360381019061028691906132ce565b6107bd565b604051610298919061373d565b60405180910390f35b6102bb60048036038101906102b691906131f1565b610837565b6040516102c8919061373d565b60405180910390f35b6102eb60048036038101906102e691906131f1565b61089e565b005b610307600480360381019061030291906131f1565b6109e0565b6040516103149190613758565b60405180910390f35b6103376004803603810190610332919061321a565b6109ff565b005b610353600480360381019061034e919061321a565b610a20565b005b61036f600480360381019061036a91906131f1565b610aa3565b604051610380959493929190613773565b60405180910390f35b610391610b19565b005b6103ad60048036038101906103a89190613256565b610b4e565b6040516103ba919061373d565b60405180910390f35b6103dd60048036038101906103d89190613052565b610b6e565b005b6103f960048036038101906103f491906131f1565b610c25565b60405161040691906138c4565b60405180910390f35b610429600480360381019061042491906131f1565b610d1c565b60405161043691906138df565b60405180910390f35b610459600480360381019061045491906131f1565b610e32565b005b610463610f74565b604051610470919061373d565b60405180910390f35b610493600480360381019061048e91906131f1565b610f8b565b6040516104a091906138a8565b60405180910390f35b6104b16110a1565b005b6104cd60048036038101906104c891906131f1565b611129565b6040516104da919061373d565b60405180910390f35b6104eb61114f565b6040516104f89190613758565b60405180910390f35b610509611173565b005b610525600480360381019061052091906131f1565b6111a8565b604051610532919061373d565b60405180910390f35b610555600480360381019061055091906131f1565b6111ce565b005b61055f611378565b60405161056c9190613722565b60405180910390f35b61058f600480360381019061058a9190613292565b6113a2565b60405161059c9190613722565b60405180910390f35b6105bf60048036038101906105ba919061321a565b6113d1565b6040516105cc919061373d565b60405180910390f35b6105ef60048036038101906105ea919061307b565b61143b565b005b61060b60048036038101906106069190613052565b61169d565b005b610627600480360381019061062291906131a2565b611754565b005b610631611895565b60405161063e9190613758565b60405180910390f35b610661600480360381019061065c9190613104565b61189c565b005b61067d60048036038101906106789190613256565b611a65565b60405161068a919061373d565b60405180910390f35b6106ad60048036038101906106a891906131f1565b611a86565b6040516106ba91906138fb565b60405180910390f35b6106dd60048036038101906106d89190613256565b611aaa565b6040516106ea919061373d565b60405180910390f35b61070d6004803603810190610708919061321a565b611acb565b005b61072960048036038101906107249190613104565b611aec565b005b610733611cb5565b6040516107409190613758565b60405180910390f35b610763600480360381019061075e91906131a2565b611cd9565b005b61077f600480360381019061077a9190613052565b611e27565b005b61079b60048036038101906107969190613052565b611ede565b005b600081600760008581526020019081526020016000205414905092915050565b60007ffd922689000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916827bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19161480610830575061082f82612126565b5b9050919050565b60008060001b600760008481526020019081526020016000205414801561087357506000801b6009600084815260200190815260200160002054145b801561089757506000801b600a600084815260200190815260200160002060000154145b9050919050565b6108a6610f74565b156108e6576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016108dd90613848565b60405180910390fd5b600560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff168073ffffffffffffffffffffffffffffffffffffffff166109286121a0565b73ffffffffffffffffffffffffffffffffffffffff1614610975576040517f2d0b96e300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6000801b600960008481526020019081526020016000205414156109c5576040517fa0e9298400000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60096000838152602001908152602001600020600090555050565b6000806000838152602001908152602001600020600101549050919050565b610a08826109e0565b610a11816121a8565b610a1b83836121bc565b505050565b610a286121a0565b73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614610a95576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610a8c90613888565b60405180910390fd5b610a9f8282612260565b5050565b600a6020528060005260406000206000915090508060000154908060010160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060020160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060030154908060040154905085565b7f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a610b43816121a8565b610b4b612304565b50565b600081600960008581526020019081526020016000205414905092915050565b610b76610f74565b15610bb6576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610bad90613848565b60405180910390fd5b7fa49807205ce4d355092ef5a8a18f56e8913cf4a201fbe287825b095693c21775610be0816121a8565b81600560006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505050565b610c2d612eb8565b600a60008381526020019081526020016000206040518060a0016040529081600082015481526020016001820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016002820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001600382015481526020016004820154815250509050919050565b610d24612f16565b6006600083815260200190815260200160002060405180610100016040529081600082015481526020016001820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001600282015481526020016003820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016004820154815260200160058201548152602001600682015481526020016007820154815250509050919050565b610e3a610f74565b15610e7a576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610e7190613848565b60405180910390fd5b600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff168073ffffffffffffffffffffffffffffffffffffffff16610ebc6121a0565b73ffffffffffffffffffffffffffffffffffffffff1614610f09576040517f2d0b96e300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6000801b60076000848152602001908152602001600020541415610f59576040517fa0e9298400000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60076000838152602001908152602001600020600090555050565b6000600260149054906101000a900460ff16905090565b610f93612f8a565b6008600083815260200190815260200160002060405180610100016040529081600082015481526020016001820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200160028201548152602001600382015481526020016004820154815260200160058201548152602001600682015481526020016007820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815250509050919050565b6110a96121a0565b73ffffffffffffffffffffffffffffffffffffffff166110c7611378565b73ffffffffffffffffffffffffffffffffffffffff161461111d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161111490613868565b60405180910390fd5b61112760006123a6565b565b60008060001b600660008481526020019081526020016000206000015414159050919050565b7fa49807205ce4d355092ef5a8a18f56e8913cf4a201fbe287825b095693c2177581565b7f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a61119d816121a8565b6111a561246c565b50565b60008060001b600860008481526020019081526020016000206000015414159050919050565b6111d6610f74565b15611216576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161120d90613848565b60405180910390fd5b600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff168073ffffffffffffffffffffffffffffffffffffffff166112586121a0565b73ffffffffffffffffffffffffffffffffffffffff16146112a5576040517f2d0b96e300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6000801b600a60008481526020019081526020016000206000015414156112f8576040517f84791e5200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600a60008381526020019081526020016000206000808201600090556001820160006101000a81549073ffffffffffffffffffffffffffffffffffffffff02191690556002820160006101000a81549073ffffffffffffffffffffffffffffffffffffffff02191690556003820160009055600482016000905550505050565b6000600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b60006113c9826001600086815260200190815260200160002061250f90919063ffffffff16565b905092915050565b600080600084815260200190815260200160002060000160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16905092915050565b600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff168073ffffffffffffffffffffffffffffffffffffffff1661147d6121a0565b73ffffffffffffffffffffffffffffffffffffffff16146114ca576040517f2d0b96e300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6114d2610f74565b15611512576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161150990613848565b60405180910390fd5b600086866040516020016115279291906136bc565b60405160208183030381529060405280519060200120905061154881610837565b61157e576040517fc0a24f6700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6040518060a001604052808681526020018973ffffffffffffffffffffffffffffffffffffffff1681526020018873ffffffffffffffffffffffffffffffffffffffff16815260200185815260200184815250600a60008381526020019081526020016000206000820151816000015560208201518160010160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060408201518160020160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060608201518160030155608082015181600401559050505050505050505050565b6116a5610f74565b156116e5576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016116dc90613848565b60405180910390fd5b7fa49807205ce4d355092ef5a8a18f56e8913cf4a201fbe287825b095693c2177561170f816121a8565b81600360006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505050565b61175c610f74565b1561179c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161179390613848565b60405180910390fd5b600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff168073ffffffffffffffffffffffffffffffffffffffff166117de6121a0565b73ffffffffffffffffffffffffffffffffffffffff161461182b576040517f2d0b96e300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6000600660008481526020019081526020016000209050838160020181905550848160030160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505050505050565b6000801b81565b6118a4610f74565b156118e4576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016118db90613848565b60405180910390fd5b600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff168073ffffffffffffffffffffffffffffffffffffffff166119266121a0565b73ffffffffffffffffffffffffffffffffffffffff1614611973576040517f2d0b96e300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b61197c87610837565b6119b2576040517fc0a24f6700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8560076000898152602001908152602001600020819055506000600660008881526020019081526020016000209050868160000181905550888160010160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550858160040181905550848160050181905550838160060181905550828160070181905550848160020181905550505050505050505050565b60008160076000858152602001908152602001600020541415905092915050565b6000611aa360016000848152602001908152602001600020612529565b9050919050565b60008160096000858152602001908152602001600020541415905092915050565b611ad4826109e0565b611add816121a8565b611ae78383612260565b505050565b611af4610f74565b15611b34576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611b2b90613848565b60405180910390fd5b600560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff168073ffffffffffffffffffffffffffffffffffffffff16611b766121a0565b73ffffffffffffffffffffffffffffffffffffffff1614611bc3576040517f2d0b96e300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b611bcc87610837565b611c02576040517fc0a24f6700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8560096000898152602001908152602001600020819055506000600860008881526020019081526020016000209050868160000181905550888160010160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550848160030181905550838160040181905550858160020181905550828160050181905550828160060181905550505050505050505050565b7f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a81565b611ce1610f74565b15611d21576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611d1890613848565b60405180910390fd5b600560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff168073ffffffffffffffffffffffffffffffffffffffff16611d636121a0565b73ffffffffffffffffffffffffffffffffffffffff1614611db0576040517f2d0b96e300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6000600860008481526020019081526020016000209050611dd4816003015461253e565b838160060181905550848160070160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505050505050565b611e2f610f74565b15611e6f576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611e6690613848565b60405180910390fd5b7fa49807205ce4d355092ef5a8a18f56e8913cf4a201fbe287825b095693c21775611e99816121a8565b81600460006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505050565b611ee66121a0565b73ffffffffffffffffffffffffffffffffffffffff16611f04611378565b73ffffffffffffffffffffffffffffffffffffffff1614611f5a576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611f5190613868565b60405180910390fd5b611f847fa49807205ce4d355092ef5a8a18f56e8913cf4a201fbe287825b095693c21775826121bc565b611f916000801b826121bc565b611fc27fa49807205ce4d355092ef5a8a18f56e8913cf4a201fbe287825b095693c21775611fbd611378565b612260565b611fd66000801b611fd1611378565b612260565b611fdf81612585565b50565b611fec8282612016565b61201181600160008581526020019081526020016000206120f690919063ffffffff16565b505050565b61202082826113d1565b6120f257600160008084815260200190815260200160002060000160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055506120976121a0565b73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16837f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a45b5050565b600061211e836000018373ffffffffffffffffffffffffffffffffffffffff1660001b61267d565b905092915050565b60007f5a05180f000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916827bffffffffffffffffffffffffffffffffffffffffffffffffffffffff191614806121995750612198826126ed565b5b9050919050565b600033905090565b6121b9816121b46121a0565b612767565b50565b6121c68282611fe2565b7fa49807205ce4d355092ef5a8a18f56e8913cf4a201fbe287825b095693c217758214806121f657506000801b82145b1561225c576122257f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a82611fe2565b6000801b82141561225b5761225a7fa49807205ce4d355092ef5a8a18f56e8913cf4a201fbe287825b095693c2177582611fe2565b5b5b5050565b61226a8282612804565b7fa49807205ce4d355092ef5a8a18f56e8913cf4a201fbe287825b095693c2177582148061229a57506000801b82145b15612300576122c97f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a82612804565b6000801b8214156122ff576122fe7fa49807205ce4d355092ef5a8a18f56e8913cf4a201fbe287825b095693c2177582611fe2565b5b5b5050565b61230c610f74565b61234b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161234290613808565b60405180910390fd5b6000600260146101000a81548160ff0219169083151502179055507f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa61238f6121a0565b60405161239c9190613722565b60405180910390a1565b6000600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905081600260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b612474610f74565b156124b4576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016124ab90613848565b60405180910390fd5b6001600260146101000a81548160ff0219169083151502179055507f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a2586124f86121a0565b6040516125059190613722565b60405180910390a1565b600061251e8360000183612838565b60001c905092915050565b600061253782600001612889565b9050919050565b80421161258257806040517f2a35a32400000000000000000000000000000000000000000000000000000000815260040161257991906138fb565b60405180910390fd5b50565b61258d6121a0565b73ffffffffffffffffffffffffffffffffffffffff166125ab611378565b73ffffffffffffffffffffffffffffffffffffffff1614612601576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016125f890613868565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff161415612671576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161266890613828565b60405180910390fd5b61267a816123a6565b50565b6000612689838361289a565b6126e25782600001829080600181540180825580915050600190039060005260206000200160009091909190915055826000018054905083600101600084815260200190815260200160002081905550600190506126e7565b600090505b92915050565b60007f7965db0b000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916827bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19161480612760575061275f826128bd565b5b9050919050565b61277182826113d1565b612800576127968173ffffffffffffffffffffffffffffffffffffffff166014612927565b6127a48360001c6020612927565b6040516020016127b59291906136e8565b6040516020818303038152906040526040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016127f791906137c6565b60405180910390fd5b5050565b61280e8282612c21565b6128338160016000858152602001908152602001600020612d0290919063ffffffff16565b505050565b6000826000018281548110612876577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b9060005260206000200154905092915050565b600081600001805490509050919050565b600080836001016000848152602001908152602001600020541415905092915050565b60007f01ffc9a7000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916827bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916149050919050565b60606000600283600261293a9190613993565b612944919061393d565b67ffffffffffffffff811115612983577f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6040519080825280601f01601f1916602001820160405280156129b55781602001600182028036833780820191505090505b5090507f300000000000000000000000000000000000000000000000000000000000000081600081518110612a13577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a9053507f780000000000000000000000000000000000000000000000000000000000000081600181518110612a9d577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a90535060006001846002612add9190613993565b612ae7919061393d565b90505b6001811115612bd3577f3031323334353637383961626364656600000000000000000000000000000000600f861660108110612b4f577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b1a60f81b828281518110612b8c577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a905350600485901c945080612bcc90613ad2565b9050612aea565b5060008414612c17576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612c0e906137e8565b60405180910390fd5b8091505092915050565b612c2b82826113d1565b15612cfe57600080600084815260200190815260200160002060000160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff021916908315150217905550612ca36121a0565b73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16837ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b60405160405180910390a45b5050565b6000612d2a836000018373ffffffffffffffffffffffffffffffffffffffff1660001b612d32565b905092915050565b60008083600101600084815260200190815260200160002054905060008114612eac576000600182612d6491906139ed565b9050600060018660000180549050612d7c91906139ed565b9050818114612e37576000866000018281548110612dc3577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b9060005260206000200154905080876000018481548110612e0d577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b90600052602060002001819055508387600101600083815260200190815260200160002081905550505b85600001805480612e71577f4e487b7100000000000000000000000000000000000000000000000000000000600052603160045260246000fd5b600190038181906000526020600020016000905590558560010160008681526020019081526020016000206000905560019350505050612eb2565b60009150505b92915050565b6040518060a0016040528060008019168152602001600073ffffffffffffffffffffffffffffffffffffffff168152602001600073ffffffffffffffffffffffffffffffffffffffff16815260200160008152602001600081525090565b60405180610100016040528060008019168152602001600073ffffffffffffffffffffffffffffffffffffffff16815260200160008152602001600073ffffffffffffffffffffffffffffffffffffffff168152602001600081526020016000815260200160008152602001600081525090565b60405180610100016040528060008019168152602001600073ffffffffffffffffffffffffffffffffffffffff1681526020016000815260200160008152602001600081526020016000815260200160008152602001600073ffffffffffffffffffffffffffffffffffffffff1681525090565b60008135905061300d81613d0b565b92915050565b60008135905061302281613d22565b92915050565b60008135905061303781613d39565b92915050565b60008135905061304c81613d50565b92915050565b60006020828403121561306457600080fd5b600061307284828501612ffe565b91505092915050565b60008060008060008060c0878903121561309457600080fd5b60006130a289828a01612ffe565b96505060206130b389828a01612ffe565b95505060406130c489828a0161303d565b94505060606130d589828a01613013565b93505060806130e689828a0161303d565b92505060a06130f789828a0161303d565b9150509295509295509295565b600080600080600080600060e0888a03121561311f57600080fd5b600061312d8a828b01612ffe565b975050602061313e8a828b01613013565b965050604061314f8a828b01613013565b95505060606131608a828b0161303d565b94505060806131718a828b0161303d565b93505060a06131828a828b0161303d565b92505060c06131938a828b0161303d565b91505092959891949750929550565b6000806000606084860312156131b757600080fd5b60006131c586828701612ffe565b93505060206131d68682870161303d565b92505060406131e786828701613013565b9150509250925092565b60006020828403121561320357600080fd5b600061321184828501613013565b91505092915050565b6000806040838503121561322d57600080fd5b600061323b85828601613013565b925050602061324c85828601612ffe565b9150509250929050565b6000806040838503121561326957600080fd5b600061327785828601613013565b925050602061328885828601613013565b9150509250929050565b600080604083850312156132a557600080fd5b60006132b385828601613013565b92505060206132c48582860161303d565b9150509250929050565b6000602082840312156132e057600080fd5b60006132ee84828501613028565b91505092915050565b61330081613a21565b82525050565b61330f81613a21565b82525050565b61332661332182613a21565b613afc565b82525050565b61333581613a33565b82525050565b61334481613a3f565b82525050565b61335381613a3f565b82525050565b600061336482613916565b61336e8185613921565b935061337e818560208601613a9f565b61338781613b59565b840191505092915050565b600061339d82613916565b6133a78185613932565b93506133b7818560208601613a9f565b80840191505092915050565b60006133d0602083613921565b91506133db82613b77565b602082019050919050565b60006133f3601483613921565b91506133fe82613ba0565b602082019050919050565b6000613416602683613921565b915061342182613bc9565b604082019050919050565b6000613439601083613921565b915061344482613c18565b602082019050919050565b600061345c602083613921565b915061346782613c41565b602082019050919050565b600061347f601783613932565b915061348a82613c6a565b601782019050919050565b60006134a2601183613932565b91506134ad82613c93565b601182019050919050565b60006134c5602f83613921565b91506134d082613cbc565b604082019050919050565b610100820160008201516134f2600085018261333b565b50602082015161350560208501826132f7565b5060408201516135186040850182613687565b50606082015161352b6060850182613687565b50608082015161353e6080850182613687565b5060a082015161355160a0850182613687565b5060c082015161356460c0850182613687565b5060e082015161357760e08501826132f7565b50505050565b60a082016000820151613593600085018261333b565b5060208201516135a660208501826132f7565b5060408201516135b960408501826132f7565b5060608201516135cc6060850182613687565b5060808201516135df6080850182613687565b50505050565b610100820160008201516135fc600085018261333b565b50602082015161360f60208501826132f7565b5060408201516136226040850182613687565b50606082015161363560608501826132f7565b5060808201516136486080850182613687565b5060a082015161365b60a0850182613687565b5060c082015161366e60c0850182613687565b5060e082015161368160e0850182613687565b50505050565b61369081613a95565b82525050565b61369f81613a95565b82525050565b6136b66136b182613a95565b613b20565b82525050565b60006136c88285613315565b6014820191506136d882846136a5565b6020820191508190509392505050565b60006136f382613472565b91506136ff8285613392565b915061370a82613495565b91506137168284613392565b91508190509392505050565b60006020820190506137376000830184613306565b92915050565b6000602082019050613752600083018461332c565b92915050565b600060208201905061376d600083018461334a565b92915050565b600060a082019050613788600083018861334a565b6137956020830187613306565b6137a26040830186613306565b6137af6060830185613696565b6137bc6080830184613696565b9695505050505050565b600060208201905081810360008301526137e08184613359565b905092915050565b60006020820190508181036000830152613801816133c3565b9050919050565b60006020820190508181036000830152613821816133e6565b9050919050565b6000602082019050818103600083015261384181613409565b9050919050565b600060208201905081810360008301526138618161342c565b9050919050565b600060208201905081810360008301526138818161344f565b9050919050565b600060208201905081810360008301526138a1816134b8565b9050919050565b6000610100820190506138be60008301846134db565b92915050565b600060a0820190506138d9600083018461357d565b92915050565b6000610100820190506138f560008301846135e5565b92915050565b60006020820190506139106000830184613696565b92915050565b600081519050919050565b600082825260208201905092915050565b600081905092915050565b600061394882613a95565b915061395383613a95565b9250827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0382111561398857613987613b2a565b5b828201905092915050565b600061399e82613a95565b91506139a983613a95565b9250817fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff04831182151516156139e2576139e1613b2a565b5b828202905092915050565b60006139f882613a95565b9150613a0383613a95565b925082821015613a1657613a15613b2a565b5b828203905092915050565b6000613a2c82613a75565b9050919050565b60008115159050919050565b6000819050919050565b60007fffffffff0000000000000000000000000000000000000000000000000000000082169050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b60005b83811015613abd578082015181840152602081019050613aa2565b83811115613acc576000848401525b50505050565b6000613add82613a95565b91506000821415613af157613af0613b2a565b5b600182039050919050565b6000613b0782613b0e565b9050919050565b6000613b1982613b6a565b9050919050565b6000819050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b6000601f19601f8301169050919050565b60008160601b9050919050565b7f537472696e67733a20686578206c656e67746820696e73756666696369656e74600082015250565b7f5061757361626c653a206e6f7420706175736564000000000000000000000000600082015250565b7f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160008201527f6464726573730000000000000000000000000000000000000000000000000000602082015250565b7f5061757361626c653a2070617573656400000000000000000000000000000000600082015250565b7f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572600082015250565b7f416363657373436f6e74726f6c3a206163636f756e7420000000000000000000600082015250565b7f206973206d697373696e6720726f6c6520000000000000000000000000000000600082015250565b7f416363657373436f6e74726f6c3a2063616e206f6e6c792072656e6f756e636560008201527f20726f6c657320666f722073656c660000000000000000000000000000000000602082015250565b613d1481613a21565b8114613d1f57600080fd5b50565b613d2b81613a3f565b8114613d3657600080fd5b50565b613d4281613a49565b8114613d4d57600080fd5b50565b613d5981613a95565b8114613d6457600080fd5b5056fea26469706673582212201f0d3cbaf1e40371e0054e19b307154b315a80f4bd015734b28822957b50d85964736f6c63430008040033",
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
// Solidity: function getBlindAuction(bytes32 blindAuctionId) view returns((bytes32,address,uint256,uint256,uint256,uint256,uint256,address) blindAuction)
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
// Solidity: function getBlindAuction(bytes32 blindAuctionId) view returns((bytes32,address,uint256,uint256,uint256,uint256,uint256,address) blindAuction)
func (_MarketplaceStorage *MarketplaceStorageSession) GetBlindAuction(blindAuctionId [32]byte) (IBlindAuctionBlindAuction, error) {
	return _MarketplaceStorage.Contract.GetBlindAuction(&_MarketplaceStorage.CallOpts, blindAuctionId)
}

// GetBlindAuction is a free data retrieval call binding the contract method 0x6ed3be53.
//
// Solidity: function getBlindAuction(bytes32 blindAuctionId) view returns((bytes32,address,uint256,uint256,uint256,uint256,uint256,address) blindAuction)
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
// Solidity: function getPublicAuction(bytes32 publicAuctionId) view returns((bytes32,address,uint256,address,uint256,uint256,uint256,uint256) publicAuction)
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
// Solidity: function getPublicAuction(bytes32 publicAuctionId) view returns((bytes32,address,uint256,address,uint256,uint256,uint256,uint256) publicAuction)
func (_MarketplaceStorage *MarketplaceStorageSession) GetPublicAuction(publicAuctionId [32]byte) (IPublicAuctionPublicAuction, error) {
	return _MarketplaceStorage.Contract.GetPublicAuction(&_MarketplaceStorage.CallOpts, publicAuctionId)
}

// GetPublicAuction is a free data retrieval call binding the contract method 0x57a45b77.
//
// Solidity: function getPublicAuction(bytes32 publicAuctionId) view returns((bytes32,address,uint256,address,uint256,uint256,uint256,uint256) publicAuction)
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

// CreateBlindAuction is a paid mutator transaction binding the contract method 0xdeaf324a.
//
// Solidity: function createBlindAuction(address assetOwner, bytes32 nftAsset, bytes32 blindAuctionId, uint256 startTime, uint256 biddingEnd, uint256 revealEnd, uint256 startPriceInWei) returns()
func (_MarketplaceStorage *MarketplaceStorageTransactor) CreateBlindAuction(opts *bind.TransactOpts, assetOwner common.Address, nftAsset [32]byte, blindAuctionId [32]byte, startTime *big.Int, biddingEnd *big.Int, revealEnd *big.Int, startPriceInWei *big.Int) (*types.Transaction, error) {
	return _MarketplaceStorage.contract.Transact(opts, "createBlindAuction", assetOwner, nftAsset, blindAuctionId, startTime, biddingEnd, revealEnd, startPriceInWei)
}

// CreateBlindAuction is a paid mutator transaction binding the contract method 0xdeaf324a.
//
// Solidity: function createBlindAuction(address assetOwner, bytes32 nftAsset, bytes32 blindAuctionId, uint256 startTime, uint256 biddingEnd, uint256 revealEnd, uint256 startPriceInWei) returns()
func (_MarketplaceStorage *MarketplaceStorageSession) CreateBlindAuction(assetOwner common.Address, nftAsset [32]byte, blindAuctionId [32]byte, startTime *big.Int, biddingEnd *big.Int, revealEnd *big.Int, startPriceInWei *big.Int) (*types.Transaction, error) {
	return _MarketplaceStorage.Contract.CreateBlindAuction(&_MarketplaceStorage.TransactOpts, assetOwner, nftAsset, blindAuctionId, startTime, biddingEnd, revealEnd, startPriceInWei)
}

// CreateBlindAuction is a paid mutator transaction binding the contract method 0xdeaf324a.
//
// Solidity: function createBlindAuction(address assetOwner, bytes32 nftAsset, bytes32 blindAuctionId, uint256 startTime, uint256 biddingEnd, uint256 revealEnd, uint256 startPriceInWei) returns()
func (_MarketplaceStorage *MarketplaceStorageTransactorSession) CreateBlindAuction(assetOwner common.Address, nftAsset [32]byte, blindAuctionId [32]byte, startTime *big.Int, biddingEnd *big.Int, revealEnd *big.Int, startPriceInWei *big.Int) (*types.Transaction, error) {
	return _MarketplaceStorage.Contract.CreateBlindAuction(&_MarketplaceStorage.TransactOpts, assetOwner, nftAsset, blindAuctionId, startTime, biddingEnd, revealEnd, startPriceInWei)
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

// CreatePublicAuction is a paid mutator transaction binding the contract method 0xae97d559.
//
// Solidity: function createPublicAuction(address assetOwner, bytes32 nftAsset, bytes32 publicAuctionId, uint256 biddingEnd, uint256 startPriceInWei, uint256 startTime, uint256 minIncrement) returns()
func (_MarketplaceStorage *MarketplaceStorageTransactor) CreatePublicAuction(opts *bind.TransactOpts, assetOwner common.Address, nftAsset [32]byte, publicAuctionId [32]byte, biddingEnd *big.Int, startPriceInWei *big.Int, startTime *big.Int, minIncrement *big.Int) (*types.Transaction, error) {
	return _MarketplaceStorage.contract.Transact(opts, "createPublicAuction", assetOwner, nftAsset, publicAuctionId, biddingEnd, startPriceInWei, startTime, minIncrement)
}

// CreatePublicAuction is a paid mutator transaction binding the contract method 0xae97d559.
//
// Solidity: function createPublicAuction(address assetOwner, bytes32 nftAsset, bytes32 publicAuctionId, uint256 biddingEnd, uint256 startPriceInWei, uint256 startTime, uint256 minIncrement) returns()
func (_MarketplaceStorage *MarketplaceStorageSession) CreatePublicAuction(assetOwner common.Address, nftAsset [32]byte, publicAuctionId [32]byte, biddingEnd *big.Int, startPriceInWei *big.Int, startTime *big.Int, minIncrement *big.Int) (*types.Transaction, error) {
	return _MarketplaceStorage.Contract.CreatePublicAuction(&_MarketplaceStorage.TransactOpts, assetOwner, nftAsset, publicAuctionId, biddingEnd, startPriceInWei, startTime, minIncrement)
}

// CreatePublicAuction is a paid mutator transaction binding the contract method 0xae97d559.
//
// Solidity: function createPublicAuction(address assetOwner, bytes32 nftAsset, bytes32 publicAuctionId, uint256 biddingEnd, uint256 startPriceInWei, uint256 startTime, uint256 minIncrement) returns()
func (_MarketplaceStorage *MarketplaceStorageTransactorSession) CreatePublicAuction(assetOwner common.Address, nftAsset [32]byte, publicAuctionId [32]byte, biddingEnd *big.Int, startPriceInWei *big.Int, startTime *big.Int, minIncrement *big.Int) (*types.Transaction, error) {
	return _MarketplaceStorage.Contract.CreatePublicAuction(&_MarketplaceStorage.TransactOpts, assetOwner, nftAsset, publicAuctionId, biddingEnd, startPriceInWei, startTime, minIncrement)
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

// MarketplaceStorageAuctionCancelledSuccessfulIterator is returned from FilterAuctionCancelledSuccessful and is used to iterate over the raw logs and unpacked data for AuctionCancelledSuccessful events raised by the MarketplaceStorage contract.
type MarketplaceStorageAuctionCancelledSuccessfulIterator struct {
	Event *MarketplaceStorageAuctionCancelledSuccessful // Event containing the contract specifics and raw log

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
func (it *MarketplaceStorageAuctionCancelledSuccessfulIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketplaceStorageAuctionCancelledSuccessful)
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
		it.Event = new(MarketplaceStorageAuctionCancelledSuccessful)
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
func (it *MarketplaceStorageAuctionCancelledSuccessfulIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketplaceStorageAuctionCancelledSuccessfulIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketplaceStorageAuctionCancelledSuccessful represents a AuctionCancelledSuccessful event raised by the MarketplaceStorage contract.
type MarketplaceStorageAuctionCancelledSuccessful struct {
	Canceller common.Address
	AuctionId [32]byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterAuctionCancelledSuccessful is a free log retrieval operation binding the contract event 0xa021239bb373fef19aa7c7ef798961ab038ebf95bd85d546b731d25855db190e.
//
// Solidity: event AuctionCancelledSuccessful(address canceller, bytes32 auctionId)
func (_MarketplaceStorage *MarketplaceStorageFilterer) FilterAuctionCancelledSuccessful(opts *bind.FilterOpts) (*MarketplaceStorageAuctionCancelledSuccessfulIterator, error) {

	logs, sub, err := _MarketplaceStorage.contract.FilterLogs(opts, "AuctionCancelledSuccessful")
	if err != nil {
		return nil, err
	}
	return &MarketplaceStorageAuctionCancelledSuccessfulIterator{contract: _MarketplaceStorage.contract, event: "AuctionCancelledSuccessful", logs: logs, sub: sub}, nil
}

// WatchAuctionCancelledSuccessful is a free log subscription operation binding the contract event 0xa021239bb373fef19aa7c7ef798961ab038ebf95bd85d546b731d25855db190e.
//
// Solidity: event AuctionCancelledSuccessful(address canceller, bytes32 auctionId)
func (_MarketplaceStorage *MarketplaceStorageFilterer) WatchAuctionCancelledSuccessful(opts *bind.WatchOpts, sink chan<- *MarketplaceStorageAuctionCancelledSuccessful) (event.Subscription, error) {

	logs, sub, err := _MarketplaceStorage.contract.WatchLogs(opts, "AuctionCancelledSuccessful")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketplaceStorageAuctionCancelledSuccessful)
				if err := _MarketplaceStorage.contract.UnpackLog(event, "AuctionCancelledSuccessful", log); err != nil {
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

// ParseAuctionCancelledSuccessful is a log parse operation binding the contract event 0xa021239bb373fef19aa7c7ef798961ab038ebf95bd85d546b731d25855db190e.
//
// Solidity: event AuctionCancelledSuccessful(address canceller, bytes32 auctionId)
func (_MarketplaceStorage *MarketplaceStorageFilterer) ParseAuctionCancelledSuccessful(log types.Log) (*MarketplaceStorageAuctionCancelledSuccessful, error) {
	event := new(MarketplaceStorageAuctionCancelledSuccessful)
	if err := _MarketplaceStorage.contract.UnpackLog(event, "AuctionCancelledSuccessful", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MarketplaceStorageAuctionRefundSuccessfulIterator is returned from FilterAuctionRefundSuccessful and is used to iterate over the raw logs and unpacked data for AuctionRefundSuccessful events raised by the MarketplaceStorage contract.
type MarketplaceStorageAuctionRefundSuccessfulIterator struct {
	Event *MarketplaceStorageAuctionRefundSuccessful // Event containing the contract specifics and raw log

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
func (it *MarketplaceStorageAuctionRefundSuccessfulIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketplaceStorageAuctionRefundSuccessful)
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
		it.Event = new(MarketplaceStorageAuctionRefundSuccessful)
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
func (it *MarketplaceStorageAuctionRefundSuccessfulIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketplaceStorageAuctionRefundSuccessfulIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketplaceStorageAuctionRefundSuccessful represents a AuctionRefundSuccessful event raised by the MarketplaceStorage contract.
type MarketplaceStorageAuctionRefundSuccessful struct {
	Bidder    common.Address
	AuctionId [32]byte
	Value     *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterAuctionRefundSuccessful is a free log retrieval operation binding the contract event 0x5f4b88832be3746851e4d0ce6129f89fe39ddcf5a4b0204699724c459baf7d9d.
//
// Solidity: event AuctionRefundSuccessful(address bidder, bytes32 auctionId, uint256 value)
func (_MarketplaceStorage *MarketplaceStorageFilterer) FilterAuctionRefundSuccessful(opts *bind.FilterOpts) (*MarketplaceStorageAuctionRefundSuccessfulIterator, error) {

	logs, sub, err := _MarketplaceStorage.contract.FilterLogs(opts, "AuctionRefundSuccessful")
	if err != nil {
		return nil, err
	}
	return &MarketplaceStorageAuctionRefundSuccessfulIterator{contract: _MarketplaceStorage.contract, event: "AuctionRefundSuccessful", logs: logs, sub: sub}, nil
}

// WatchAuctionRefundSuccessful is a free log subscription operation binding the contract event 0x5f4b88832be3746851e4d0ce6129f89fe39ddcf5a4b0204699724c459baf7d9d.
//
// Solidity: event AuctionRefundSuccessful(address bidder, bytes32 auctionId, uint256 value)
func (_MarketplaceStorage *MarketplaceStorageFilterer) WatchAuctionRefundSuccessful(opts *bind.WatchOpts, sink chan<- *MarketplaceStorageAuctionRefundSuccessful) (event.Subscription, error) {

	logs, sub, err := _MarketplaceStorage.contract.WatchLogs(opts, "AuctionRefundSuccessful")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketplaceStorageAuctionRefundSuccessful)
				if err := _MarketplaceStorage.contract.UnpackLog(event, "AuctionRefundSuccessful", log); err != nil {
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

// ParseAuctionRefundSuccessful is a log parse operation binding the contract event 0x5f4b88832be3746851e4d0ce6129f89fe39ddcf5a4b0204699724c459baf7d9d.
//
// Solidity: event AuctionRefundSuccessful(address bidder, bytes32 auctionId, uint256 value)
func (_MarketplaceStorage *MarketplaceStorageFilterer) ParseAuctionRefundSuccessful(log types.Log) (*MarketplaceStorageAuctionRefundSuccessful, error) {
	event := new(MarketplaceStorageAuctionRefundSuccessful)
	if err := _MarketplaceStorage.contract.UnpackLog(event, "AuctionRefundSuccessful", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
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
	Sender         common.Address
	BlindAuctionId [32]byte
	BlindedBid     [32]byte
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterBlindAuctionBidSuccessful is a free log retrieval operation binding the contract event 0x42412bd3e1b349d0fc9c3518c85d09f611398d80842c5d3e36449b76dca01b6b.
//
// Solidity: event BlindAuctionBidSuccessful(address sender, bytes32 blindAuctionId, bytes32 blindedBid)
func (_MarketplaceStorage *MarketplaceStorageFilterer) FilterBlindAuctionBidSuccessful(opts *bind.FilterOpts) (*MarketplaceStorageBlindAuctionBidSuccessfulIterator, error) {

	logs, sub, err := _MarketplaceStorage.contract.FilterLogs(opts, "BlindAuctionBidSuccessful")
	if err != nil {
		return nil, err
	}
	return &MarketplaceStorageBlindAuctionBidSuccessfulIterator{contract: _MarketplaceStorage.contract, event: "BlindAuctionBidSuccessful", logs: logs, sub: sub}, nil
}

// WatchBlindAuctionBidSuccessful is a free log subscription operation binding the contract event 0x42412bd3e1b349d0fc9c3518c85d09f611398d80842c5d3e36449b76dca01b6b.
//
// Solidity: event BlindAuctionBidSuccessful(address sender, bytes32 blindAuctionId, bytes32 blindedBid)
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
// Solidity: event BlindAuctionBidSuccessful(address sender, bytes32 blindAuctionId, bytes32 blindedBid)
func (_MarketplaceStorage *MarketplaceStorageFilterer) ParseBlindAuctionBidSuccessful(log types.Log) (*MarketplaceStorageBlindAuctionBidSuccessful, error) {
	event := new(MarketplaceStorageBlindAuctionBidSuccessful)
	if err := _MarketplaceStorage.contract.UnpackLog(event, "BlindAuctionBidSuccessful", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MarketplaceStorageBlindAuctionCreatedSuccessfulIterator is returned from FilterBlindAuctionCreatedSuccessful and is used to iterate over the raw logs and unpacked data for BlindAuctionCreatedSuccessful events raised by the MarketplaceStorage contract.
type MarketplaceStorageBlindAuctionCreatedSuccessfulIterator struct {
	Event *MarketplaceStorageBlindAuctionCreatedSuccessful // Event containing the contract specifics and raw log

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
func (it *MarketplaceStorageBlindAuctionCreatedSuccessfulIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketplaceStorageBlindAuctionCreatedSuccessful)
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
		it.Event = new(MarketplaceStorageBlindAuctionCreatedSuccessful)
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
func (it *MarketplaceStorageBlindAuctionCreatedSuccessfulIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketplaceStorageBlindAuctionCreatedSuccessfulIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketplaceStorageBlindAuctionCreatedSuccessful represents a BlindAuctionCreatedSuccessful event raised by the MarketplaceStorage contract.
type MarketplaceStorageBlindAuctionCreatedSuccessful struct {
	AssetOwner      common.Address
	NftAddress      common.Address
	BlindAuctionId  [32]byte
	AssetId         *big.Int
	StartTime       *big.Int
	BiddingEnd      *big.Int
	RevealEnd       *big.Int
	StartPriceInWei *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterBlindAuctionCreatedSuccessful is a free log retrieval operation binding the contract event 0xa1605c4cefd3351cfe80d4cdc106707dea9aff20e3095109748a5f710ecf52d4.
//
// Solidity: event BlindAuctionCreatedSuccessful(address assetOwner, address nftAddress, bytes32 blindAuctionId, uint256 assetId, uint256 startTime, uint256 biddingEnd, uint256 revealEnd, uint256 startPriceInWei)
func (_MarketplaceStorage *MarketplaceStorageFilterer) FilterBlindAuctionCreatedSuccessful(opts *bind.FilterOpts) (*MarketplaceStorageBlindAuctionCreatedSuccessfulIterator, error) {

	logs, sub, err := _MarketplaceStorage.contract.FilterLogs(opts, "BlindAuctionCreatedSuccessful")
	if err != nil {
		return nil, err
	}
	return &MarketplaceStorageBlindAuctionCreatedSuccessfulIterator{contract: _MarketplaceStorage.contract, event: "BlindAuctionCreatedSuccessful", logs: logs, sub: sub}, nil
}

// WatchBlindAuctionCreatedSuccessful is a free log subscription operation binding the contract event 0xa1605c4cefd3351cfe80d4cdc106707dea9aff20e3095109748a5f710ecf52d4.
//
// Solidity: event BlindAuctionCreatedSuccessful(address assetOwner, address nftAddress, bytes32 blindAuctionId, uint256 assetId, uint256 startTime, uint256 biddingEnd, uint256 revealEnd, uint256 startPriceInWei)
func (_MarketplaceStorage *MarketplaceStorageFilterer) WatchBlindAuctionCreatedSuccessful(opts *bind.WatchOpts, sink chan<- *MarketplaceStorageBlindAuctionCreatedSuccessful) (event.Subscription, error) {

	logs, sub, err := _MarketplaceStorage.contract.WatchLogs(opts, "BlindAuctionCreatedSuccessful")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketplaceStorageBlindAuctionCreatedSuccessful)
				if err := _MarketplaceStorage.contract.UnpackLog(event, "BlindAuctionCreatedSuccessful", log); err != nil {
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

// ParseBlindAuctionCreatedSuccessful is a log parse operation binding the contract event 0xa1605c4cefd3351cfe80d4cdc106707dea9aff20e3095109748a5f710ecf52d4.
//
// Solidity: event BlindAuctionCreatedSuccessful(address assetOwner, address nftAddress, bytes32 blindAuctionId, uint256 assetId, uint256 startTime, uint256 biddingEnd, uint256 revealEnd, uint256 startPriceInWei)
func (_MarketplaceStorage *MarketplaceStorageFilterer) ParseBlindAuctionCreatedSuccessful(log types.Log) (*MarketplaceStorageBlindAuctionCreatedSuccessful, error) {
	event := new(MarketplaceStorageBlindAuctionCreatedSuccessful)
	if err := _MarketplaceStorage.contract.UnpackLog(event, "BlindAuctionCreatedSuccessful", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MarketplaceStorageGrantAuctionRewardSuccessfulIterator is returned from FilterGrantAuctionRewardSuccessful and is used to iterate over the raw logs and unpacked data for GrantAuctionRewardSuccessful events raised by the MarketplaceStorage contract.
type MarketplaceStorageGrantAuctionRewardSuccessfulIterator struct {
	Event *MarketplaceStorageGrantAuctionRewardSuccessful // Event containing the contract specifics and raw log

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
func (it *MarketplaceStorageGrantAuctionRewardSuccessfulIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketplaceStorageGrantAuctionRewardSuccessful)
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
		it.Event = new(MarketplaceStorageGrantAuctionRewardSuccessful)
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
func (it *MarketplaceStorageGrantAuctionRewardSuccessfulIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketplaceStorageGrantAuctionRewardSuccessfulIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketplaceStorageGrantAuctionRewardSuccessful represents a GrantAuctionRewardSuccessful event raised by the MarketplaceStorage contract.
type MarketplaceStorageGrantAuctionRewardSuccessful struct {
	AuctionHighestBidder common.Address
	AuctionId            [32]byte
	AssetId              *big.Int
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterGrantAuctionRewardSuccessful is a free log retrieval operation binding the contract event 0xa307d88e125ed19583b02332c5bd74f0323774248b3f9cac5047c2433d9d8ddc.
//
// Solidity: event GrantAuctionRewardSuccessful(address auctionHighestBidder, bytes32 auctionId, uint256 assetId)
func (_MarketplaceStorage *MarketplaceStorageFilterer) FilterGrantAuctionRewardSuccessful(opts *bind.FilterOpts) (*MarketplaceStorageGrantAuctionRewardSuccessfulIterator, error) {

	logs, sub, err := _MarketplaceStorage.contract.FilterLogs(opts, "GrantAuctionRewardSuccessful")
	if err != nil {
		return nil, err
	}
	return &MarketplaceStorageGrantAuctionRewardSuccessfulIterator{contract: _MarketplaceStorage.contract, event: "GrantAuctionRewardSuccessful", logs: logs, sub: sub}, nil
}

// WatchGrantAuctionRewardSuccessful is a free log subscription operation binding the contract event 0xa307d88e125ed19583b02332c5bd74f0323774248b3f9cac5047c2433d9d8ddc.
//
// Solidity: event GrantAuctionRewardSuccessful(address auctionHighestBidder, bytes32 auctionId, uint256 assetId)
func (_MarketplaceStorage *MarketplaceStorageFilterer) WatchGrantAuctionRewardSuccessful(opts *bind.WatchOpts, sink chan<- *MarketplaceStorageGrantAuctionRewardSuccessful) (event.Subscription, error) {

	logs, sub, err := _MarketplaceStorage.contract.WatchLogs(opts, "GrantAuctionRewardSuccessful")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketplaceStorageGrantAuctionRewardSuccessful)
				if err := _MarketplaceStorage.contract.UnpackLog(event, "GrantAuctionRewardSuccessful", log); err != nil {
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

// ParseGrantAuctionRewardSuccessful is a log parse operation binding the contract event 0xa307d88e125ed19583b02332c5bd74f0323774248b3f9cac5047c2433d9d8ddc.
//
// Solidity: event GrantAuctionRewardSuccessful(address auctionHighestBidder, bytes32 auctionId, uint256 assetId)
func (_MarketplaceStorage *MarketplaceStorageFilterer) ParseGrantAuctionRewardSuccessful(log types.Log) (*MarketplaceStorageGrantAuctionRewardSuccessful, error) {
	event := new(MarketplaceStorageGrantAuctionRewardSuccessful)
	if err := _MarketplaceStorage.contract.UnpackLog(event, "GrantAuctionRewardSuccessful", log); err != nil {
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
	Sender          common.Address
	PublicAuctionId [32]byte
	BidValue        *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterPublicAuctionBidSuccessful is a free log retrieval operation binding the contract event 0x9f541988b3aac74ac43a6f4b330ceed668244b74461b90850d199a38a172f7c2.
//
// Solidity: event PublicAuctionBidSuccessful(address sender, bytes32 publicAuctionId, uint256 bidValue)
func (_MarketplaceStorage *MarketplaceStorageFilterer) FilterPublicAuctionBidSuccessful(opts *bind.FilterOpts) (*MarketplaceStoragePublicAuctionBidSuccessfulIterator, error) {

	logs, sub, err := _MarketplaceStorage.contract.FilterLogs(opts, "PublicAuctionBidSuccessful")
	if err != nil {
		return nil, err
	}
	return &MarketplaceStoragePublicAuctionBidSuccessfulIterator{contract: _MarketplaceStorage.contract, event: "PublicAuctionBidSuccessful", logs: logs, sub: sub}, nil
}

// WatchPublicAuctionBidSuccessful is a free log subscription operation binding the contract event 0x9f541988b3aac74ac43a6f4b330ceed668244b74461b90850d199a38a172f7c2.
//
// Solidity: event PublicAuctionBidSuccessful(address sender, bytes32 publicAuctionId, uint256 bidValue)
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
// Solidity: event PublicAuctionBidSuccessful(address sender, bytes32 publicAuctionId, uint256 bidValue)
func (_MarketplaceStorage *MarketplaceStorageFilterer) ParsePublicAuctionBidSuccessful(log types.Log) (*MarketplaceStoragePublicAuctionBidSuccessful, error) {
	event := new(MarketplaceStoragePublicAuctionBidSuccessful)
	if err := _MarketplaceStorage.contract.UnpackLog(event, "PublicAuctionBidSuccessful", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MarketplaceStoragePublicAuctionCreatedSuccessfulIterator is returned from FilterPublicAuctionCreatedSuccessful and is used to iterate over the raw logs and unpacked data for PublicAuctionCreatedSuccessful events raised by the MarketplaceStorage contract.
type MarketplaceStoragePublicAuctionCreatedSuccessfulIterator struct {
	Event *MarketplaceStoragePublicAuctionCreatedSuccessful // Event containing the contract specifics and raw log

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
func (it *MarketplaceStoragePublicAuctionCreatedSuccessfulIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketplaceStoragePublicAuctionCreatedSuccessful)
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
		it.Event = new(MarketplaceStoragePublicAuctionCreatedSuccessful)
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
func (it *MarketplaceStoragePublicAuctionCreatedSuccessfulIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketplaceStoragePublicAuctionCreatedSuccessfulIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketplaceStoragePublicAuctionCreatedSuccessful represents a PublicAuctionCreatedSuccessful event raised by the MarketplaceStorage contract.
type MarketplaceStoragePublicAuctionCreatedSuccessful struct {
	AssetOwner      common.Address
	NftAddress      common.Address
	PublicAuctionId [32]byte
	AssetId         *big.Int
	BiddingEnd      *big.Int
	StartPriceInWei *big.Int
	StartTime       *big.Int
	MinIncrement    *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterPublicAuctionCreatedSuccessful is a free log retrieval operation binding the contract event 0xd3c25d0f5cc0da484b0fc13f082a5e7da800d235500c5904281dbdbb9e22ba90.
//
// Solidity: event PublicAuctionCreatedSuccessful(address assetOwner, address nftAddress, bytes32 publicAuctionId, uint256 assetId, uint256 biddingEnd, uint256 startPriceInWei, uint256 startTime, uint256 minIncrement)
func (_MarketplaceStorage *MarketplaceStorageFilterer) FilterPublicAuctionCreatedSuccessful(opts *bind.FilterOpts) (*MarketplaceStoragePublicAuctionCreatedSuccessfulIterator, error) {

	logs, sub, err := _MarketplaceStorage.contract.FilterLogs(opts, "PublicAuctionCreatedSuccessful")
	if err != nil {
		return nil, err
	}
	return &MarketplaceStoragePublicAuctionCreatedSuccessfulIterator{contract: _MarketplaceStorage.contract, event: "PublicAuctionCreatedSuccessful", logs: logs, sub: sub}, nil
}

// WatchPublicAuctionCreatedSuccessful is a free log subscription operation binding the contract event 0xd3c25d0f5cc0da484b0fc13f082a5e7da800d235500c5904281dbdbb9e22ba90.
//
// Solidity: event PublicAuctionCreatedSuccessful(address assetOwner, address nftAddress, bytes32 publicAuctionId, uint256 assetId, uint256 biddingEnd, uint256 startPriceInWei, uint256 startTime, uint256 minIncrement)
func (_MarketplaceStorage *MarketplaceStorageFilterer) WatchPublicAuctionCreatedSuccessful(opts *bind.WatchOpts, sink chan<- *MarketplaceStoragePublicAuctionCreatedSuccessful) (event.Subscription, error) {

	logs, sub, err := _MarketplaceStorage.contract.WatchLogs(opts, "PublicAuctionCreatedSuccessful")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketplaceStoragePublicAuctionCreatedSuccessful)
				if err := _MarketplaceStorage.contract.UnpackLog(event, "PublicAuctionCreatedSuccessful", log); err != nil {
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

// ParsePublicAuctionCreatedSuccessful is a log parse operation binding the contract event 0xd3c25d0f5cc0da484b0fc13f082a5e7da800d235500c5904281dbdbb9e22ba90.
//
// Solidity: event PublicAuctionCreatedSuccessful(address assetOwner, address nftAddress, bytes32 publicAuctionId, uint256 assetId, uint256 biddingEnd, uint256 startPriceInWei, uint256 startTime, uint256 minIncrement)
func (_MarketplaceStorage *MarketplaceStorageFilterer) ParsePublicAuctionCreatedSuccessful(log types.Log) (*MarketplaceStoragePublicAuctionCreatedSuccessful, error) {
	event := new(MarketplaceStoragePublicAuctionCreatedSuccessful)
	if err := _MarketplaceStorage.contract.UnpackLog(event, "PublicAuctionCreatedSuccessful", log); err != nil {
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
	Sender         common.Address
	BlindAuctionId [32]byte
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterRevealSuccessful is a free log retrieval operation binding the contract event 0x61d4ab60470c97b371c5f91be47873283386bb3af4de728e8795ef5318e013d6.
//
// Solidity: event RevealSuccessful(address sender, bytes32 blindAuctionId)
func (_MarketplaceStorage *MarketplaceStorageFilterer) FilterRevealSuccessful(opts *bind.FilterOpts) (*MarketplaceStorageRevealSuccessfulIterator, error) {

	logs, sub, err := _MarketplaceStorage.contract.FilterLogs(opts, "RevealSuccessful")
	if err != nil {
		return nil, err
	}
	return &MarketplaceStorageRevealSuccessfulIterator{contract: _MarketplaceStorage.contract, event: "RevealSuccessful", logs: logs, sub: sub}, nil
}

// WatchRevealSuccessful is a free log subscription operation binding the contract event 0x61d4ab60470c97b371c5f91be47873283386bb3af4de728e8795ef5318e013d6.
//
// Solidity: event RevealSuccessful(address sender, bytes32 blindAuctionId)
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
// Solidity: event RevealSuccessful(address sender, bytes32 blindAuctionId)
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
