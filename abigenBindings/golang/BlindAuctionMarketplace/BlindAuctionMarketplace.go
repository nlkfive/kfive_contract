// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package BlindAuctionMarketplace

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

// BlindAuctionMarketplaceMetaData contains all meta data concerning the BlindAuctionMarketplace contract.
var BlindAuctionMarketplaceMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"acceptedToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"marketplaceStorage\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"beneficary\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"ownerCutPerMillion\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AuctionEnded\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidBiddingEnd\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidBiddingPrice\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidContract\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidCut\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidPrice\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidWithdraw\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotBidYet\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotExisted\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotRunning\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RewardGranted\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"TooEarly\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"TooLate\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Unauthorized\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Unavailable\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"AddedBlackList\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"canceller\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"auctionId\",\"type\":\"bytes32\"}],\"name\":\"AuctionCancelledSuccessful\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"bidder\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"auctionId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"AuctionRefundSuccessful\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"blindAuctionId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"blindedBid\",\"type\":\"bytes32\"}],\"name\":\"BlindAuctionBidSuccessful\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"assetOwner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"nftAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"blindAuctionId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"assetId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"biddingEnd\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"revealEnd\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"startPriceInWei\",\"type\":\"uint256\"}],\"name\":\"BlindAuctionCreatedSuccessful\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"ownerCutPerMillion\",\"type\":\"uint256\"}],\"name\":\"ChangedOwnerCutPerMillion\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"publicationFee\",\"type\":\"uint256\"}],\"name\":\"ChangedPublicationFee\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"auctionHighestBidder\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"auctionId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"assetId\",\"type\":\"uint256\"}],\"name\":\"GrantAuctionRewardSuccessful\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_marketplaceStorage\",\"type\":\"address\"}],\"name\":\"MarketplaceStorageUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"RemovedBlackList\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"blindAuctionId\",\"type\":\"bytes32\"}],\"name\":\"RevealSuccessful\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"CANCEL_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ERC721_Interface\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PAUSER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"acceptedToken\",\"outputs\":[{\"internalType\":\"contractIBEP20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_evilUser\",\"type\":\"address\"}],\"name\":\"addBlackList\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"beneficary\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_maker\",\"type\":\"address\"}],\"name\":\"getBlackListStatus\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getRoleMember\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleMemberCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"isBlackListed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"marketplaceStorage\",\"outputs\":[{\"internalType\":\"contractIMarketplaceStorage\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minStageDuration\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ownerCutPerMillion\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"publicationFeeInWei\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_clearedUser\",\"type\":\"address\"}],\"name\":\"removeBlackList\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_beneficary\",\"type\":\"address\"}],\"name\":\"setBeneficary\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"}],\"name\":\"setMinStageDuration\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_ownerCutPerMillion\",\"type\":\"uint256\"}],\"name\":\"setOwnerCutPerMillion\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_publicationFee\",\"type\":\"uint256\"}],\"name\":\"setPublicationFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_marketplaceStorage\",\"type\":\"address\"}],\"name\":\"updateStorageAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"nftAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"assetId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startPriceInWei\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"biddingEnd\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"revealEnd\",\"type\":\"uint256\"}],\"name\":\"createBlindAuction\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"nftAsset\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"blindAuctionId\",\"type\":\"bytes32\"}],\"name\":\"cancelBlindAuction\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"nftAsset\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"blindAuctionId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"blindedBid\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"deposit\",\"type\":\"uint256\"}],\"name\":\"bidBlindAuction\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"blindAuctionId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"assetId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"nftAddress\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"nftAsset\",\"type\":\"bytes32\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"blindAuctionId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"nftAsset\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[]\",\"name\":\"_values\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes32\",\"name\":\"secret\",\"type\":\"bytes32\"}],\"name\":\"RevealBid\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"auctionId\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"bidder\",\"type\":\"address\"}],\"name\":\"getBlindBid\",\"outputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"\",\"type\":\"bytes32[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"auctionId\",\"type\":\"bytes32\"}],\"name\":\"checkExisted\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"auctionId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"nftAsset\",\"type\":\"bytes32\"}],\"name\":\"checkEnded\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"nftAsset\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"auctionId\",\"type\":\"bytes32\"}],\"name\":\"checkRunning\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60806040523480156200001157600080fd5b5060405162006f8f38038062006f8f8339818101604052810190620000379190620013f0565b838284838180620000698173ffffffffffffffffffffffffffffffffffffffff166200042860201b62002c941760201c565b620000ab576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401620000a29062001679565b60405180910390fd5b600081905082600260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505050506000600260146101000a81548160ff0219169083151502179055506200012f620001236200044b60201b60201c565b620004fe60201b60201c565b6200014f620001436200044b60201b60201c565b620005c460201b60201c565b6200017b8473ffffffffffffffffffffffffffffffffffffffff166200042860201b62002c941760201c565b620001b2576040517f6eefed2000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b83600560006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555081600260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550620002586000801b6200024c6200044b60201b60201c565b6200078b60201b60201c565b620002997fa49807205ce4d355092ef5a8a18f56e8913cf4a201fbe287825b095693c217756200028d6200044b60201b60201c565b6200078b60201b60201c565b620002da7f9f959e00d95122f5cbd677010436cf273ef535b86b056afc172852144b9491d7620002ce6200044b60201b60201c565b6200078b60201b60201c565b6200031b7f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a6200030f6200044b60201b60201c565b6200078b60201b60201c565b6200036d7f9f959e00d95122f5cbd677010436cf273ef535b86b056afc172852144b9491d77fa49807205ce4d355092ef5a8a18f56e8913cf4a201fbe287825b095693c21775620007a160201b60201c565b620003bf7f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a7fa49807205ce4d355092ef5a8a18f56e8913cf4a201fbe287825b095693c21775620007a160201b60201c565b620003d0816200080460201b60201c565b610e1060088190555082600960006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550505050505050505062001a2d565b6000808273ffffffffffffffffffffffffffffffffffffffff163b119050919050565b60003073ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161415620004f757600080368080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f820116905080830192505050505050509050600080369050905073ffffffffffffffffffffffffffffffffffffffff818301511692505050620004fb565b3390505b90565b6000600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905081600360006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b620005d46200044b60201b60201c565b73ffffffffffffffffffffffffffffffffffffffff16620005fa620008b960201b60201c565b73ffffffffffffffffffffffffffffffffffffffff161462000653576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016200064a906200169b565b60405180910390fd5b620006857f9f959e00d95122f5cbd677010436cf273ef535b86b056afc172852144b9491d7826200078b60201b60201c565b620006b77f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a826200078b60201b60201c565b620006cc6000801b826200078b60201b60201c565b6200070d7f9f959e00d95122f5cbd677010436cf273ef535b86b056afc172852144b9491d762000701620008b960201b60201c565b620008e360201b60201c565b6200074e7f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a62000742620008b960201b60201c565b620008e360201b60201c565b620007726000801b62000766620008b960201b60201c565b620008e360201b60201c565b62000788816200091c60201b62002cb71760201c565b50565b6200079d828262000a3260201b60201c565b5050565b6000620007b48362000a7a60201b60201c565b905081600080858152602001908152602001600020600101819055508181847fbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff60405160405180910390a4505050565b7fa49807205ce4d355092ef5a8a18f56e8913cf4a201fbe287825b095693c21775620008368162000a9960201b60201c565b620f4240821062000873576040517f7e0d5ce600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b816006819055507ffa406a120a9e7f2b332bfb7a43d3bf1c3f079262202907a6b69c94b2821a02c6600654604051620008ad9190620016bd565b60405180910390a15050565b6000600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b620008f48262000a7a60201b60201c565b620009058162000a9960201b60201c565b62000917838362000abd60201b60201c565b505050565b6200092c6200044b60201b60201c565b73ffffffffffffffffffffffffffffffffffffffff1662000952620008b960201b60201c565b73ffffffffffffffffffffffffffffffffffffffff1614620009ab576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401620009a2906200169b565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16141562000a1e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040162000a159062001657565b60405180910390fd5b62000a2f81620004fe60201b60201c565b50565b62000a49828262000b0560201b62002daf1760201c565b62000a75816001600085815260200190815260200160002062000bf660201b62002e8f1790919060201c565b505050565b6000806000838152602001908152602001600020600101549050919050565b62000aba8162000aae6200044b60201b60201c565b62000c2e60201b60201c565b50565b62000ad4828262000cf260201b62002ebf1760201c565b62000b00816001600085815260200190815260200160002062000de460201b62002fa01790919060201c565b505050565b62000b17828262000e1c60201b60201c565b62000bf257600160008084815260200190815260200160002060000160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff02191690831515021790555062000b976200044b60201b60201c565b73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16837f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a45b5050565b600062000c26836000018373ffffffffffffffffffffffffffffffffffffffff1660001b62000e8660201b60201c565b905092915050565b62000c40828262000e1c60201b60201c565b62000cee5762000c738173ffffffffffffffffffffffffffffffffffffffff16601462000f0060201b62002fd01760201c565b62000c8e8360001c602062000f0060201b62002fd01760201c565b60405160200162000ca1929190620015cf565b6040516020818303038152906040526040517f08c379a000000000000000000000000000000000000000000000000000000000815260040162000ce5919062001611565b60405180910390fd5b5050565b62000d04828262000e1c60201b60201c565b1562000de057600080600084815260200190815260200160002060000160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff02191690831515021790555062000d856200044b60201b60201c565b73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16837ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b60405160405180910390a45b5050565b600062000e14836000018373ffffffffffffffffffffffffffffffffffffffff1660001b6200120f60201b60201c565b905092915050565b600080600084815260200190815260200160002060000160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16905092915050565b600062000e9a83836200139f60201b60201c565b62000ef557826000018290806001815401808255809150506001900390600052602060002001600090919091909150558260000180549050836001016000848152602001908152602001600020819055506001905062000efa565b600090505b92915050565b60606000600283600262000f1591906200175e565b62000f21919062001701565b67ffffffffffffffff81111562000f61577f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6040519080825280601f01601f19166020018201604052801562000f945781602001600182028036833780820191505090505b5090507f30000000000000000000000000000000000000000000000000000000000000008160008151811062000ff3577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a9053507f7800000000000000000000000000000000000000000000000000000000000000816001815181106200107e577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a90535060006001846002620010c091906200175e565b620010cc919062001701565b90505b6001811115620011be577f3031323334353637383961626364656600000000000000000000000000000000600f86166010811062001136577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b1a60f81b82828151811062001174577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a905350600485901c945080620011b6906200186e565b9050620010cf565b506000841462001205576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401620011fc9062001635565b60405180910390fd5b8091505092915050565b6000808360010160008481526020019081526020016000205490506000811462001393576000600182620012449190620017bf565b90506000600186600001805490506200125e9190620017bf565b90508181146200131c576000866000018281548110620012a7577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b9060005260206000200154905080876000018481548110620012f2577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b90600052602060002001819055508387600101600083815260200190815260200160002081905550505b8560000180548062001357577f4e487b7100000000000000000000000000000000000000000000000000000000600052603160045260246000fd5b60019003818190600052602060002001600090559055856001016000868152602001908152602001600020600090556001935050505062001399565b60009150505b92915050565b600080836001016000848152602001908152602001600020541415905092915050565b600081519050620013d381620019f9565b92915050565b600081519050620013ea8162001a13565b92915050565b600080600080608085870312156200140757600080fd5b60006200141787828801620013c2565b94505060206200142a87828801620013c2565b93505060406200143d87828801620013c2565b92505060606200145087828801620013d9565b91505092959194509250565b60006200146982620016da565b620014758185620016e5565b93506200148781856020860162001838565b6200149281620018cc565b840191505092915050565b6000620014aa82620016da565b620014b68185620016f6565b9350620014c881856020860162001838565b80840191505092915050565b6000620014e3602083620016e5565b9150620014f082620018dd565b602082019050919050565b60006200150a602683620016e5565b9150620015178262001906565b604082019050919050565b600062001531601083620016e5565b91506200153e8262001955565b602082019050919050565b600062001558602083620016e5565b915062001565826200197e565b602082019050919050565b60006200157f601783620016f6565b91506200158c82620019a7565b601782019050919050565b6000620015a6601183620016f6565b9150620015b382620019d0565b601182019050919050565b620015c9816200182e565b82525050565b6000620015dc8262001570565b9150620015ea82856200149d565b9150620015f78262001597565b91506200160582846200149d565b91508190509392505050565b600060208201905081810360008301526200162d81846200145c565b905092915050565b600060208201905081810360008301526200165081620014d4565b9050919050565b600060208201905081810360008301526200167281620014fb565b9050919050565b60006020820190508181036000830152620016948162001522565b9050919050565b60006020820190508181036000830152620016b68162001549565b9050919050565b6000602082019050620016d46000830184620015be565b92915050565b600081519050919050565b600082825260208201905092915050565b600081905092915050565b60006200170e826200182e565b91506200171b836200182e565b9250827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff038211156200175357620017526200189d565b5b828201905092915050565b60006200176b826200182e565b915062001778836200182e565b9250817fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0483118215151615620017b457620017b36200189d565b5b828202905092915050565b6000620017cc826200182e565b9150620017d9836200182e565b925082821015620017ef57620017ee6200189d565b5b828203905092915050565b600062001807826200180e565b9050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b60005b83811015620018585780820151818401526020810190506200183b565b8381111562001868576000848401525b50505050565b60006200187b826200182e565b915060008214156200189257620018916200189d565b5b600182039050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b6000601f19601f8301169050919050565b7f537472696e67733a20686578206c656e67746820696e73756666696369656e74600082015250565b7f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160008201527f6464726573730000000000000000000000000000000000000000000000000000602082015250565b7f496e76616c696420636f6e747261637400000000000000000000000000000000600082015250565b7f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572600082015250565b7f416363657373436f6e74726f6c3a206163636f756e7420000000000000000000600082015250565b7f206973206d697373696e6720726f6c6520000000000000000000000000000000600082015250565b62001a0481620017fa565b811462001a1057600080fd5b50565b62001a1e816200182e565b811462001a2a57600080fd5b50565b6155528062001a3d6000396000f3fe608060405234801561001057600080fd5b50600436106102685760003560e01c806359bf1abe11610151578063ae4f1198116100c3578063d547741f11610087578063d547741f146106e7578063e47d606014610703578063e4997dc514610733578063e63ab1e91461074f578063f2fde38b1461076d578063f60218841461078957610268565b8063ae4f119814610641578063af8996f11461065f578063b42cf9291461067b578063ca15c87314610699578063cbfda1c5146106c957610268565b8063889e212911610115578063889e21291461056b5780638da5cb5b146105875780639010d07c146105a557806391d14854146105d5578063a01f79d414610605578063a217fddf1461062357610268565b806359bf1abe146104eb5780635c975abb1461051b578063715018a61461053957806375b238fc146105435780638456cb591461056157610268565b80632b4c32be116101ea57806336568abe116101ae57806336568abe1461043d5780633af4b422146104595780633f4ba83a14610489578063451c3d801461049357806346b3aec6146104b157806356479889146104cf57610268565b80632b4c32be146103af5780632f2ff15d146103cd578063314b3ee4146103e95780633414bd2614610405578063347ea9d61461042157610268565b80630ecb93c0116102315780630ecb93c01461030d57806313c27ca71461032957806319dad16d14610347578063222bbd8714610363578063248a9ca31461037f57610268565b8062c5c3c31461026d57806301ffc9a714610289578063026967e0146102b9578063076f6dc7146102d55780630db3cc05146102f1575b600080fd5b61028760048036038101906102829190614419565b6107a5565b005b6102a3600480360381019061029e9190614637565b610889565b6040516102b09190614d10565b60405180910390f35b6102d360048036038101906102ce9190614379565b610903565b005b6102ef60048036038101906102ea919061447e565b61101e565b005b61030b60048036038101906103069190614327565b61110e565b005b61032760048036038101906103229190614327565b61117d565b005b61033161128b565b60405161033e9190614d2b565b60405180910390f35b610361600480360381019061035c9190614689565b6112af565b005b61037d600480360381019061037891906144ba565b611359565b005b61039960048036038101906103949190614419565b6118b4565b6040516103a69190614d2b565b60405180910390f35b6103b76118d3565b6040516103c49190614d6f565b60405180910390f35b6103e760048036038101906103e29190614442565b6118de565b005b61040360048036038101906103fe9190614535565b6118ff565b005b61041f600480360381019061041a919061447e565b611ccb565b005b61043b600480360381019061043691906145d4565b611f52565b005b61045760048036038101906104529190614442565b6123e0565b005b610473600480360381019061046e9190614442565b612463565b6040516104809190614cee565b60405180910390f35b61049161250c565b005b61049b612541565b6040516104a89190614d8a565b60405180910390f35b6104b9612567565b6040516104c69190614da5565b60405180910390f35b6104e960048036038101906104e4919061447e565b61258d565b005b61050560048036038101906105009190614327565b612674565b6040516105129190614d10565b60405180910390f35b6105236126ca565b6040516105309190614d10565b60405180910390f35b6105416126e1565b005b61054b612769565b6040516105589190614d2b565b60405180910390f35b61056961278d565b005b61058560048036038101906105809190614327565b6127c2565b005b61058f6127f9565b60405161059c9190614a67565b60405180910390f35b6105bf60048036038101906105ba9190614598565b612823565b6040516105cc9190614a67565b60405180910390f35b6105ef60048036038101906105ea9190614442565b612852565b6040516105fc9190614d10565b60405180910390f35b61060d6128bc565b60405161061a9190614ea2565b60405180910390f35b61062b6128c2565b6040516106389190614d2b565b60405180910390f35b6106496128c9565b6040516106569190614ea2565b60405180910390f35b61067960048036038101906106749190614689565b6128cf565b005b61068361293d565b6040516106909190614ea2565b60405180910390f35b6106b360048036038101906106ae9190614419565b612943565b6040516106c09190614ea2565b60405180910390f35b6106d1612967565b6040516106de9190614a67565b60405180910390f35b61070160048036038101906106fc9190614442565b61298d565b005b61071d60048036038101906107189190614327565b6129ae565b60405161072a9190614d10565b60405180910390f35b61074d60048036038101906107489190614327565b6129ce565b005b610757612adc565b6040516107649190614d2b565b60405180910390f35b61078760048036038101906107829190614327565b612b00565b005b6107a3600480360381019061079e9190614689565b612c5f565b005b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166386624448826040518263ffffffff1660e01b81526004016108009190614d2b565b60206040518083038186803b15801561081857600080fd5b505afa15801561082c573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061085091906143f0565b610886576040517fafdd489000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b50565b60007f5a05180f000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916827bffffffffffffffffffffffffffffffffffffffffffffffffffffffff191614806108fc57506108fb826132ca565b5b9050919050565b846109238173ffffffffffffffffffffffffffffffffffffffff16612c94565b610959576040517f6eefed2000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8073ffffffffffffffffffffffffffffffffffffffff166301ffc9a76380ac58cd60e01b6040518263ffffffff1660e01b81526004016109999190614d6f565b60206040518083038186803b1580156109b157600080fd5b505afa1580156109c5573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906109e991906143f0565b6109f257600080fd5b6109fa6126ca565b15610a3a576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610a3190614e42565b60405180910390fd5b6000610a516008544261334490919063ffffffff16565b841015610a8a576040517f18c71d1a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6000851415610ac4576040517ebfc92100000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6000610ace61335a565b905060008890508073ffffffffffffffffffffffffffffffffffffffff16636352211e896040518263ffffffff1660e01b8152600401610b0e9190614ea2565b60206040518083038186803b158015610b2657600080fd5b505afa158015610b3a573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610b5e9190614350565b92508273ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16141580610ce157503073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1663081812fc8a6040518263ffffffff1660e01b8152600401610be69190614ea2565b60206040518083038186803b158015610bfe57600080fd5b505afa158015610c12573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610c369190614350565b73ffffffffffffffffffffffffffffffffffffffff161480610cdf57508073ffffffffffffffffffffffffffffffffffffffff1663e985e9c584306040518363ffffffff1660e01b8152600401610c8e929190614a82565b60206040518083038186803b158015610ca657600080fd5b505afa158015610cba573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610cde91906143f0565b5b155b15610d18576040517f82b4290000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16630a7f89778a8a604051602001610d69929190614976565b604051602081830303815290604052805190602001206040518263ffffffff1660e01b8152600401610d9b9190614d2b565b60206040518083038186803b158015610db357600080fd5b505afa158015610dc7573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610deb91906143f0565b610e21576040517fa3b8915f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5060006007541115610f0457600560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166323b872dd82600960009054906101000a900473ffffffffffffffffffffffffffffffffffffffff166007546040518463ffffffff1660e01b8152600401610eb093929190614b1a565b602060405180830381600087803b158015610eca57600080fd5b505af1158015610ede573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610f0291906143f0565b505b5060004282888a89604051602001610f209594939291906149dc565b604051602081830303815290604052805190602001209050600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663d7e4a307838a8a858a8a8d6040518863ffffffff1660e01b8152600401610f9f9796959493929190614b51565b600060405180830381600087803b158015610fb957600080fd5b505af1158015610fcd573d6000803e3d6000fd5b505050507fde549e5fcbd4d262f76b1f3a6362777a2e669ce8a2a7b5460e00c7a38341baf78289838a89898c60405161100c9796959493929190614aab565b60405180910390a15050505050505050565b611027816107a5565b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16634a43c5b983836040518363ffffffff1660e01b8152600401611084929190614d46565b60206040518083038186803b15801561109c57600080fd5b505afa1580156110b0573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906110d491906143f0565b61110a576040517f03b5e41300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5050565b7fa49807205ce4d355092ef5a8a18f56e8913cf4a201fbe287825b095693c217756111388161340b565b81600960006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505050565b61118561335a565b73ffffffffffffffffffffffffffffffffffffffff166111a36127f9565b73ffffffffffffffffffffffffffffffffffffffff16146111f9576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016111f090614e62565b60405180910390fd5b6001600460008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055507f42e160154868087d6bfdc0ca23d96a1c1cfa32f1b72ba9ba27b69b98a0d819dc816040516112809190614a67565b60405180910390a150565b7f9f959e00d95122f5cbd677010436cf273ef535b86b056afc172852144b9491d781565b7fa49807205ce4d355092ef5a8a18f56e8913cf4a201fbe287825b095693c217756112d98161340b565b620f42408210611315576040517f7e0d5ce600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b816006819055507ffa406a120a9e7f2b332bfb7a43d3bf1c3f079262202907a6b69c94b2821a02c660065460405161134d9190614ea2565b60405180910390a15050565b6113616126ca565b156113a1576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161139890614e42565b60405180910390fd5b6113ab838561101e565b6000600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16636ed3be53866040518263ffffffff1660e01b81526004016114089190614d2b565b60e06040518083038186803b15801561142057600080fd5b505afa158015611434573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906114589190614660565b9050611467816040015161341f565b6114748160600151613466565b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663cd23b17285876040518363ffffffff1660e01b81526004016114d1929190614d46565b60206040518083038186803b1580156114e957600080fd5b505afa1580156114fd573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061152191906143f0565b15611558576040517fa0e9298400000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600061156261335a565b90506000600b600088815260200190815260200160002060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054905060008551905060008214806115ce5750600081145b15611605576040517f45c109aa00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60005b81811015611870576116ed87828151811061164c577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b602002602001015187600a60008d815260200190815260200160002060008873ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002084815481106116dd577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b90600052602060002001546134ad565b1561185d5786818151811061172b577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b602002602001015183101580156117855750868181518110611776577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60200260200101518560a00151105b1561185c57600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663e7c847aa85898481518110611802577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60200260200101518c6040518463ffffffff1660e01b815260040161182993929190614cb7565b600060405180830381600087803b15801561184357600080fd5b505af1158015611857573d6000803e3d6000fd5b505050505b5b8080611868906151e9565b915050611608565b507f61d4ab60470c97b371c5f91be47873283386bb3af4de728e8795ef5318e013d683896040516118a2929190614bc0565b60405180910390a15050505050505050565b6000806000838152602001908152602001600020600101549050919050565b6380ac58cd60e01b81565b6118e7826118b4565b6118f08161340b565b6118fa83836134d5565b505050565b6119076126ca565b15611947576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161193e90614e42565b60405180910390fd5b611951848461101e565b6000600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16636ed3be53856040518263ffffffff1660e01b81526004016119ae9190614d2b565b60e06040518083038186803b1580156119c657600080fd5b505afa1580156119da573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906119fe9190614660565b9050611a0d8160400151613466565b6000611a1761335a565b905081608001518311611a56576040517f7532083f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600b600086815260200190815260200160002060008273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054831115611c1157600560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166323b872dd8230600b60008a815260200190815260200160002060008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205487611b4a919061504f565b6040518463ffffffff1660e01b8152600401611b6893929190614b1a565b602060405180830381600087803b158015611b8257600080fd5b505af1158015611b96573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611bba91906143f0565b5082600b600087815260200190815260200160002060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055505b600a600086815260200190815260200160002060008273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000208490806001815401808255809150506001900390600052602060002001600090919091909150557f42412bd3e1b349d0fc9c3518c85d09f611398d80842c5d3e36449b76dca01b6b818686604051611cbb93929190614be9565b60405180910390a1505050505050565b611cd36126ca565b15611d13576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611d0a90614e42565b60405180910390fd5b611d1d828261101e565b6000600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16636ed3be53836040518263ffffffff1660e01b8152600401611d7a9190614d2b565b60e06040518083038186803b158015611d9257600080fd5b505afa158015611da6573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611dca9190614660565b9050611dd98160400151613466565b6000611de361335a565b90508073ffffffffffffffffffffffffffffffffffffffff16826020015173ffffffffffffffffffffffffffffffffffffffff161480611e505750611e4f7f9f959e00d95122f5cbd677010436cf273ef535b86b056afc172852144b9491d7611e4a61335a565b612852565b5b611e86576040517f82b4290000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166315c97238856040518263ffffffff1660e01b8152600401611ee19190614d2b565b600060405180830381600087803b158015611efb57600080fd5b505af1158015611f0f573d6000803e3d6000fd5b505050507fa021239bb373fef19aa7c7ef798961ab038ebf95bd85d546b731d25855db190e8184604051611f44929190614bc0565b60405180910390a150505050565b611f5a6126ca565b15611f9a576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611f9190614e42565b60405180910390fd5b611fa3846107a5565b6000600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16636ed3be53866040518263ffffffff1660e01b81526004016120009190614d2b565b60e06040518083038186803b15801561201857600080fd5b505afa15801561202c573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906120509190614660565b9050600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663cd23b17283876040518363ffffffff1660e01b81526004016120af929190614d46565b60206040518083038186803b1580156120c757600080fd5b505afa1580156120db573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906120ff91906143f0565b1561211657612111816060015161341f565b6121a4565b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166315c97238836040518263ffffffff1660e01b81526004016121719190614d2b565b600060405180830381600087803b15801561218b57600080fd5b505af115801561219f573d6000803e3d6000fd5b505050505b60006121ae61335a565b90506000600b600088815260200190815260200160002060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205490506000811415612240576040517f35411d8b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8260c0015173ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff161415612296578260a0015181612288919061504f565b9050612295858789613509565b5b600560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663a9059cbb83836040518363ffffffff1660e01b81526004016122f3929190614c8e565b602060405180830381600087803b15801561230d57600080fd5b505af1158015612321573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061234591906143f0565b506000600b600089815260200190815260200160002060008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055507f5f4b88832be3746851e4d0ce6129f89fe39ddcf5a4b0204699724c459baf7d9d8288836040516123cf93929190614c20565b60405180910390a150505050505050565b6123e861335a565b73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614612455576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161244c90614e82565b60405180910390fd5b61245f8282613a4b565b5050565b6060600a600084815260200190815260200160002060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000208054806020026020016040519081016040528092919081815260200182805480156124ff57602002820191906000526020600020905b8154815260200190600101908083116124eb575b5050505050905092915050565b7f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a6125368161340b565b61253e613a7f565b50565b600560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663cd23b17282846040518363ffffffff1660e01b81526004016125ea929190614d46565b60206040518083038186803b15801561260257600080fd5b505afa158015612616573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061263a91906143f0565b612670576040517fafdd489000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5050565b6000600460008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff169050919050565b6000600260149054906101000a900460ff16905090565b6126e961335a565b73ffffffffffffffffffffffffffffffffffffffff166127076127f9565b73ffffffffffffffffffffffffffffffffffffffff161461275d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161275490614e62565b60405180910390fd5b6127676000613b21565b565b7fa49807205ce4d355092ef5a8a18f56e8913cf4a201fbe287825b095693c2177581565b7f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a6127b78161340b565b6127bf613be7565b50565b7fa49807205ce4d355092ef5a8a18f56e8913cf4a201fbe287825b095693c217756127ec8161340b565b6127f582613c8a565b5050565b6000600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b600061284a8260016000868152602001908152602001600020613d0590919063ffffffff16565b905092915050565b600080600084815260200190815260200160002060000160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16905092915050565b60065481565b6000801b81565b60075481565b7fa49807205ce4d355092ef5a8a18f56e8913cf4a201fbe287825b095693c217756128f98161340b565b816007819055507fe7fa8737293f41b5dfa0d5c3e552860a06275bed7015581b083c7be7003308ba6007546040516129319190614ea2565b60405180910390a15050565b60085481565b600061296060016000848152602001908152602001600020613d1f565b9050919050565b600960009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b612996826118b4565b61299f8161340b565b6129a98383613a4b565b505050565b60046020528060005260406000206000915054906101000a900460ff1681565b6129d661335a565b73ffffffffffffffffffffffffffffffffffffffff166129f46127f9565b73ffffffffffffffffffffffffffffffffffffffff1614612a4a576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612a4190614e62565b60405180910390fd5b6000600460008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055507fd7e9ec6e6ecd65492dce6bf513cd6867560d49544421d0783ddf06e76c24470c81604051612ad19190614a67565b60405180910390a150565b7f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a81565b612b0861335a565b73ffffffffffffffffffffffffffffffffffffffff16612b266127f9565b73ffffffffffffffffffffffffffffffffffffffff1614612b7c576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612b7390614e62565b60405180910390fd5b612ba67f9f959e00d95122f5cbd677010436cf273ef535b86b056afc172852144b9491d782613d34565b612bd07f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a82613d34565b612bdd6000801b82613d34565b612c0e7f9f959e00d95122f5cbd677010436cf273ef535b86b056afc172852144b9491d7612c096127f9565b61298d565b612c3f7f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a612c3a6127f9565b61298d565b612c536000801b612c4e6127f9565b61298d565b612c5c81612cb7565b50565b7fa49807205ce4d355092ef5a8a18f56e8913cf4a201fbe287825b095693c21775612c898161340b565b816008819055505050565b6000808273ffffffffffffffffffffffffffffffffffffffff163b119050919050565b612cbf61335a565b73ffffffffffffffffffffffffffffffffffffffff16612cdd6127f9565b73ffffffffffffffffffffffffffffffffffffffff1614612d33576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612d2a90614e62565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff161415612da3576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612d9a90614e22565b60405180910390fd5b612dac81613b21565b50565b612db98282612852565b612e8b57600160008084815260200190815260200160002060000160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff021916908315150217905550612e3061335a565b73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16837f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a45b5050565b6000612eb7836000018373ffffffffffffffffffffffffffffffffffffffff1660001b613d42565b905092915050565b612ec98282612852565b15612f9c57600080600084815260200190815260200160002060000160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff021916908315150217905550612f4161335a565b73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16837ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b60405160405180910390a45b5050565b6000612fc8836000018373ffffffffffffffffffffffffffffffffffffffff1660001b613db2565b905092915050565b606060006002836002612fe39190614ff5565b612fed9190614f6e565b67ffffffffffffffff81111561302c577f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6040519080825280601f01601f19166020018201604052801561305e5781602001600182028036833780820191505090505b5090507f3000000000000000000000000000000000000000000000000000000000000000816000815181106130bc577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a9053507f780000000000000000000000000000000000000000000000000000000000000081600181518110613146577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a905350600060018460026131869190614ff5565b6131909190614f6e565b90505b600181111561327c577f3031323334353637383961626364656600000000000000000000000000000000600f8616601081106131f8577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b1a60f81b828281518110613235577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a905350600485901c9450806132759061518e565b9050613193565b50600084146132c0576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016132b790614de2565b60405180910390fd5b8091505092915050565b60007f7965db0b000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916827bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916148061333d575061333c82613f38565b5b9050919050565b600081836133529190614f6e565b905092915050565b60003073ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16141561340457600080368080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f820116905080830192505050505050509050600080369050905073ffffffffffffffffffffffffffffffffffffffff818301511692505050613408565b3390505b90565b61341c8161341761335a565b613fa2565b50565b80421161346357806040517f2a35a32400000000000000000000000000000000000000000000000000000000815260040161345a9190614ea2565b60405180910390fd5b50565b8042106134aa57806040517f691e56820000000000000000000000000000000000000000000000000000000081526004016134a19190614ea2565b60405180910390fd5b50565b6000816134ba858561403f565b14156134c957600190506134ce565b600090505b9392505050565b6134df8282612daf565b6135048160016000858152602001908152602001600020612e8f90919063ffffffff16565b505050565b6000600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16636ed3be53836040518263ffffffff1660e01b81526004016135669190614d2b565b60e06040518083038186803b15801561357e57600080fd5b505afa158015613592573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906135b69190614660565b905060008160c00151905060008260200151905060008360a00151905060008790508073ffffffffffffffffffffffffffffffffffffffff16636352211e886040518263ffffffff1660e01b81526004016136119190614ea2565b60206040518083038186803b15801561362957600080fd5b505afa15801561363d573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906136619190614350565b73ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff16146136c5576040517f82b4290000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff1614806137005750600082145b15613737576040517f946ef90200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60008060065411156138425761376d620f424061375f6006548661407290919063ffffffff16565b61408890919063ffffffff16565b9050600560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663a9059cbb600960009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16836040518363ffffffff1660e01b81526004016137ee929190614c8e565b602060405180830381600087803b15801561380857600080fd5b505af115801561381c573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061384091906143f0565b505b600560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663a9059cbb85613894848761409e90919063ffffffff16565b6040518363ffffffff1660e01b81526004016138b1929190614c8e565b602060405180830381600087803b1580156138cb57600080fd5b505af11580156138df573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061390391906143f0565b508173ffffffffffffffffffffffffffffffffffffffff166342842e0e85878b6040518463ffffffff1660e01b815260040161394193929190614b1a565b600060405180830381600087803b15801561395b57600080fd5b505af115801561396f573d6000803e3d6000fd5b50505050600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663e7c847aa6000808a6040518463ffffffff1660e01b81526004016139d393929190614c57565b600060405180830381600087803b1580156139ed57600080fd5b505af1158015613a01573d6000803e3d6000fd5b505050507fa307d88e125ed19583b02332c5bd74f0323774248b3f9cac5047c2433d9d8ddc85888a604051613a3893929190614c20565b60405180910390a1505050505050505050565b613a558282612ebf565b613a7a8160016000858152602001908152602001600020612fa090919063ffffffff16565b505050565b613a876126ca565b613ac6576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401613abd90614e02565b60405180910390fd5b6000600260146101000a81548160ff0219169083151502179055507f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa613b0a61335a565b604051613b179190614a67565b60405180910390a1565b6000600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905081600360006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b613bef6126ca565b15613c2f576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401613c2690614e42565b60405180910390fd5b6001600260146101000a81548160ff0219169083151502179055507f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258613c7361335a565b604051613c809190614a67565b60405180910390a1565b80600260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055507f19e07a962d0e1cab8b7d4be06b77a7225ca7ebb744994cb6546fdd815590dc9d81604051613cfa9190614a67565b60405180910390a150565b6000613d1483600001836140b4565b60001c905092915050565b6000613d2d82600001614105565b9050919050565b613d3e82826134d5565b5050565b6000613d4e8383614116565b613da7578260000182908060018154018082558091505060019003906000526020600020016000909190919091505582600001805490508360010160008481526020019081526020016000208190555060019050613dac565b600090505b92915050565b60008083600101600084815260200190815260200160002054905060008114613f2c576000600182613de4919061504f565b9050600060018660000180549050613dfc919061504f565b9050818114613eb7576000866000018281548110613e43577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b9060005260206000200154905080876000018481548110613e8d577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b90600052602060002001819055508387600101600083815260200190815260200160002081905550505b85600001805480613ef1577f4e487b7100000000000000000000000000000000000000000000000000000000600052603160045260246000fd5b600190038181906000526020600020016000905590558560010160008681526020019081526020016000206000905560019350505050613f32565b60009150505b92915050565b60007f01ffc9a7000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916827bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916149050919050565b613fac8282612852565b61403b57613fd18173ffffffffffffffffffffffffffffffffffffffff166014612fd0565b613fdf8360001c6020612fd0565b604051602001613ff09291906149a2565b6040516020818303038152906040526040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016140329190614dc0565b60405180910390fd5b5050565b60008282604051602001614054929190614a3b565b60405160208183030381529060405280519060200120905092915050565b600081836140809190614ff5565b905092915050565b600081836140969190614fc4565b905092915050565b600081836140ac919061504f565b905092915050565b60008260000182815481106140f2577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b9060005260206000200154905092915050565b600081600001805490509050919050565b600080836001016000848152602001908152602001600020541415905092915050565b600061414c61414784614ee2565b614ebd565b9050808382526020820190508285602086028201111561416b57600080fd5b60005b8581101561419b578161418188826142fd565b84526020840193506020830192505060018101905061416e565b5050509392505050565b6000813590506141b4816154a9565b92915050565b6000815190506141c9816154a9565b92915050565b600082601f8301126141e057600080fd5b81356141f0848260208601614139565b91505092915050565b600081519050614208816154c0565b92915050565b60008135905061421d816154d7565b92915050565b600081519050614232816154d7565b92915050565b600081359050614247816154ee565b92915050565b600060e0828403121561425f57600080fd5b61426960e0614ebd565b9050600061427984828501614223565b600083015250602061428d848285016141ba565b60208301525060406142a184828501614312565b60408301525060606142b584828501614312565b60608301525060806142c984828501614312565b60808301525060a06142dd84828501614312565b60a08301525060c06142f1848285016141ba565b60c08301525092915050565b60008135905061430c81615505565b92915050565b60008151905061432181615505565b92915050565b60006020828403121561433957600080fd5b6000614347848285016141a5565b91505092915050565b60006020828403121561436257600080fd5b6000614370848285016141ba565b91505092915050565b600080600080600060a0868803121561439157600080fd5b600061439f888289016141a5565b95505060206143b0888289016142fd565b94505060406143c1888289016142fd565b93505060606143d2888289016142fd565b92505060806143e3888289016142fd565b9150509295509295909350565b60006020828403121561440257600080fd5b6000614410848285016141f9565b91505092915050565b60006020828403121561442b57600080fd5b60006144398482850161420e565b91505092915050565b6000806040838503121561445557600080fd5b60006144638582860161420e565b9250506020614474858286016141a5565b9150509250929050565b6000806040838503121561449157600080fd5b600061449f8582860161420e565b92505060206144b08582860161420e565b9150509250929050565b600080600080608085870312156144d057600080fd5b60006144de8782880161420e565b94505060206144ef8782880161420e565b935050604085013567ffffffffffffffff81111561450c57600080fd5b614518878288016141cf565b92505060606145298782880161420e565b91505092959194509250565b6000806000806080858703121561454b57600080fd5b60006145598782880161420e565b945050602061456a8782880161420e565b935050604061457b8782880161420e565b925050606061458c878288016142fd565b91505092959194509250565b600080604083850312156145ab57600080fd5b60006145b98582860161420e565b92505060206145ca858286016142fd565b9150509250929050565b600080600080608085870312156145ea57600080fd5b60006145f88782880161420e565b9450506020614609878288016142fd565b935050604061461a878288016141a5565b925050606061462b8782880161420e565b91505092959194509250565b60006020828403121561464957600080fd5b600061465784828501614238565b91505092915050565b600060e0828403121561467257600080fd5b60006146808482850161424d565b91505092915050565b60006020828403121561469b57600080fd5b60006146a9848285016142fd565b91505092915050565b60006146be838361475d565b60208301905092915050565b6146d381615083565b82525050565b6146ea6146e582615083565b615232565b82525050565b60006146fb82614f1e565b6147058185614f41565b935061471083614f0e565b8060005b8381101561474157815161472888826146b2565b975061473383614f34565b925050600181019050614714565b5085935050505092915050565b61475781615095565b82525050565b614766816150a1565b82525050565b614775816150a1565b82525050565b61478c614787826150a1565b615244565b82525050565b61479b816150ab565b82525050565b6147aa81615101565b82525050565b6147b981615125565b82525050565b6147c881615149565b82525050565b60006147d982614f29565b6147e38185614f52565b93506147f381856020860161515b565b6147fc816152f7565b840191505092915050565b600061481282614f29565b61481c8185614f63565b935061482c81856020860161515b565b80840191505092915050565b6000614845602083614f52565b915061485082615315565b602082019050919050565b6000614868601483614f52565b91506148738261533e565b602082019050919050565b600061488b602683614f52565b915061489682615367565b604082019050919050565b60006148ae601083614f52565b91506148b9826153b6565b602082019050919050565b60006148d1602083614f52565b91506148dc826153df565b602082019050919050565b60006148f4601783614f63565b91506148ff82615408565b601782019050919050565b6000614917601183614f63565b915061492282615431565b601182019050919050565b600061493a602f83614f52565b91506149458261545a565b604082019050919050565b614959816150f7565b82525050565b61497061496b826150f7565b615260565b82525050565b600061498282856146d9565b601482019150614992828461495f565b6020820191508190509392505050565b60006149ad826148e7565b91506149b98285614807565b91506149c48261490a565b91506149d08284614807565b91508190509392505050565b60006149e8828861495f565b6020820191506149f882876146d9565b601482019150614a08828661495f565b602082019150614a1882856146d9565b601482019150614a28828461495f565b6020820191508190509695505050505050565b6000614a47828561495f565b602082019150614a57828461477b565b6020820191508190509392505050565b6000602082019050614a7c60008301846146ca565b92915050565b6000604082019050614a9760008301856146ca565b614aa460208301846146ca565b9392505050565b600060e082019050614ac0600083018a6146ca565b614acd60208301896146ca565b614ada604083018861476c565b614ae76060830187614950565b614af46080830186614950565b614b0160a0830185614950565b614b0e60c0830184614950565b98975050505050505050565b6000606082019050614b2f60008301866146ca565b614b3c60208301856146ca565b614b496040830184614950565b949350505050565b600060e082019050614b66600083018a6146ca565b614b7360208301896146ca565b614b806040830188614950565b614b8d606083018761476c565b614b9a6080830186614950565b614ba760a0830185614950565b614bb460c0830184614950565b98975050505050505050565b6000604082019050614bd560008301856146ca565b614be2602083018461476c565b9392505050565b6000606082019050614bfe60008301866146ca565b614c0b602083018561476c565b614c18604083018461476c565b949350505050565b6000606082019050614c3560008301866146ca565b614c42602083018561476c565b614c4f6040830184614950565b949350505050565b6000606082019050614c6c60008301866146ca565b614c7960208301856147bf565b614c86604083018461476c565b949350505050565b6000604082019050614ca360008301856146ca565b614cb06020830184614950565b9392505050565b6000606082019050614ccc60008301866146ca565b614cd96020830185614950565b614ce6604083018461476c565b949350505050565b60006020820190508181036000830152614d0881846146f0565b905092915050565b6000602082019050614d25600083018461474e565b92915050565b6000602082019050614d40600083018461476c565b92915050565b6000604082019050614d5b600083018561476c565b614d68602083018461476c565b9392505050565b6000602082019050614d846000830184614792565b92915050565b6000602082019050614d9f60008301846147a1565b92915050565b6000602082019050614dba60008301846147b0565b92915050565b60006020820190508181036000830152614dda81846147ce565b905092915050565b60006020820190508181036000830152614dfb81614838565b9050919050565b60006020820190508181036000830152614e1b8161485b565b9050919050565b60006020820190508181036000830152614e3b8161487e565b9050919050565b60006020820190508181036000830152614e5b816148a1565b9050919050565b60006020820190508181036000830152614e7b816148c4565b9050919050565b60006020820190508181036000830152614e9b8161492d565b9050919050565b6000602082019050614eb76000830184614950565b92915050565b6000614ec7614ed8565b9050614ed382826151b8565b919050565b6000604051905090565b600067ffffffffffffffff821115614efd57614efc6152c8565b5b602082029050602081019050919050565b6000819050602082019050919050565b600081519050919050565b600081519050919050565b6000602082019050919050565b600082825260208201905092915050565b600082825260208201905092915050565b600081905092915050565b6000614f79826150f7565b9150614f84836150f7565b9250827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff03821115614fb957614fb861526a565b5b828201905092915050565b6000614fcf826150f7565b9150614fda836150f7565b925082614fea57614fe9615299565b5b828204905092915050565b6000615000826150f7565b915061500b836150f7565b9250817fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff04831182151516156150445761504361526a565b5b828202905092915050565b600061505a826150f7565b9150615065836150f7565b9250828210156150785761507761526a565b5b828203905092915050565b600061508e826150d7565b9050919050565b60008115159050919050565b6000819050919050565b60007fffffffff0000000000000000000000000000000000000000000000000000000082169050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b600061510c82615113565b9050919050565b600061511e826150d7565b9050919050565b600061513082615137565b9050919050565b6000615142826150d7565b9050919050565b6000615154826150f7565b9050919050565b60005b8381101561517957808201518184015260208101905061515e565b83811115615188576000848401525b50505050565b6000615199826150f7565b915060008214156151ad576151ac61526a565b5b600182039050919050565b6151c1826152f7565b810181811067ffffffffffffffff821117156151e0576151df6152c8565b5b80604052505050565b60006151f4826150f7565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8214156152275761522661526a565b5b600182019050919050565b600061523d8261524e565b9050919050565b6000819050919050565b600061525982615308565b9050919050565b6000819050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6000601f19601f8301169050919050565b60008160601b9050919050565b7f537472696e67733a20686578206c656e67746820696e73756666696369656e74600082015250565b7f5061757361626c653a206e6f7420706175736564000000000000000000000000600082015250565b7f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160008201527f6464726573730000000000000000000000000000000000000000000000000000602082015250565b7f5061757361626c653a2070617573656400000000000000000000000000000000600082015250565b7f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572600082015250565b7f416363657373436f6e74726f6c3a206163636f756e7420000000000000000000600082015250565b7f206973206d697373696e6720726f6c6520000000000000000000000000000000600082015250565b7f416363657373436f6e74726f6c3a2063616e206f6e6c792072656e6f756e636560008201527f20726f6c657320666f722073656c660000000000000000000000000000000000602082015250565b6154b281615083565b81146154bd57600080fd5b50565b6154c981615095565b81146154d457600080fd5b50565b6154e0816150a1565b81146154eb57600080fd5b50565b6154f7816150ab565b811461550257600080fd5b50565b61550e816150f7565b811461551957600080fd5b5056fea2646970667358221220bf894ebf2e874d5840eaa4a8c95ccee3277d328b34e3a5db7e21d0488b1863b764736f6c63430008040033",
}

// BlindAuctionMarketplaceABI is the input ABI used to generate the binding from.
// Deprecated: Use BlindAuctionMarketplaceMetaData.ABI instead.
var BlindAuctionMarketplaceABI = BlindAuctionMarketplaceMetaData.ABI

// BlindAuctionMarketplaceBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use BlindAuctionMarketplaceMetaData.Bin instead.
var BlindAuctionMarketplaceBin = BlindAuctionMarketplaceMetaData.Bin

// DeployBlindAuctionMarketplace deploys a new Ethereum contract, binding an instance of BlindAuctionMarketplace to it.
func DeployBlindAuctionMarketplace(auth *bind.TransactOpts, backend bind.ContractBackend, acceptedToken common.Address, marketplaceStorage common.Address, beneficary common.Address, ownerCutPerMillion *big.Int) (common.Address, *types.Transaction, *BlindAuctionMarketplace, error) {
	parsed, err := BlindAuctionMarketplaceMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(BlindAuctionMarketplaceBin), backend, acceptedToken, marketplaceStorage, beneficary, ownerCutPerMillion)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &BlindAuctionMarketplace{BlindAuctionMarketplaceCaller: BlindAuctionMarketplaceCaller{contract: contract}, BlindAuctionMarketplaceTransactor: BlindAuctionMarketplaceTransactor{contract: contract}, BlindAuctionMarketplaceFilterer: BlindAuctionMarketplaceFilterer{contract: contract}}, nil
}

// BlindAuctionMarketplace is an auto generated Go binding around an Ethereum contract.
type BlindAuctionMarketplace struct {
	BlindAuctionMarketplaceCaller     // Read-only binding to the contract
	BlindAuctionMarketplaceTransactor // Write-only binding to the contract
	BlindAuctionMarketplaceFilterer   // Log filterer for contract events
}

// BlindAuctionMarketplaceCaller is an auto generated read-only Go binding around an Ethereum contract.
type BlindAuctionMarketplaceCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BlindAuctionMarketplaceTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BlindAuctionMarketplaceTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BlindAuctionMarketplaceFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BlindAuctionMarketplaceFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BlindAuctionMarketplaceSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BlindAuctionMarketplaceSession struct {
	Contract     *BlindAuctionMarketplace // Generic contract binding to set the session for
	CallOpts     bind.CallOpts            // Call options to use throughout this session
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// BlindAuctionMarketplaceCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BlindAuctionMarketplaceCallerSession struct {
	Contract *BlindAuctionMarketplaceCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                  // Call options to use throughout this session
}

// BlindAuctionMarketplaceTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BlindAuctionMarketplaceTransactorSession struct {
	Contract     *BlindAuctionMarketplaceTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                  // Transaction auth options to use throughout this session
}

// BlindAuctionMarketplaceRaw is an auto generated low-level Go binding around an Ethereum contract.
type BlindAuctionMarketplaceRaw struct {
	Contract *BlindAuctionMarketplace // Generic contract binding to access the raw methods on
}

// BlindAuctionMarketplaceCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BlindAuctionMarketplaceCallerRaw struct {
	Contract *BlindAuctionMarketplaceCaller // Generic read-only contract binding to access the raw methods on
}

// BlindAuctionMarketplaceTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BlindAuctionMarketplaceTransactorRaw struct {
	Contract *BlindAuctionMarketplaceTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBlindAuctionMarketplace creates a new instance of BlindAuctionMarketplace, bound to a specific deployed contract.
func NewBlindAuctionMarketplace(address common.Address, backend bind.ContractBackend) (*BlindAuctionMarketplace, error) {
	contract, err := bindBlindAuctionMarketplace(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BlindAuctionMarketplace{BlindAuctionMarketplaceCaller: BlindAuctionMarketplaceCaller{contract: contract}, BlindAuctionMarketplaceTransactor: BlindAuctionMarketplaceTransactor{contract: contract}, BlindAuctionMarketplaceFilterer: BlindAuctionMarketplaceFilterer{contract: contract}}, nil
}

// NewBlindAuctionMarketplaceCaller creates a new read-only instance of BlindAuctionMarketplace, bound to a specific deployed contract.
func NewBlindAuctionMarketplaceCaller(address common.Address, caller bind.ContractCaller) (*BlindAuctionMarketplaceCaller, error) {
	contract, err := bindBlindAuctionMarketplace(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BlindAuctionMarketplaceCaller{contract: contract}, nil
}

// NewBlindAuctionMarketplaceTransactor creates a new write-only instance of BlindAuctionMarketplace, bound to a specific deployed contract.
func NewBlindAuctionMarketplaceTransactor(address common.Address, transactor bind.ContractTransactor) (*BlindAuctionMarketplaceTransactor, error) {
	contract, err := bindBlindAuctionMarketplace(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BlindAuctionMarketplaceTransactor{contract: contract}, nil
}

// NewBlindAuctionMarketplaceFilterer creates a new log filterer instance of BlindAuctionMarketplace, bound to a specific deployed contract.
func NewBlindAuctionMarketplaceFilterer(address common.Address, filterer bind.ContractFilterer) (*BlindAuctionMarketplaceFilterer, error) {
	contract, err := bindBlindAuctionMarketplace(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BlindAuctionMarketplaceFilterer{contract: contract}, nil
}

// bindBlindAuctionMarketplace binds a generic wrapper to an already deployed contract.
func bindBlindAuctionMarketplace(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BlindAuctionMarketplaceABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BlindAuctionMarketplace *BlindAuctionMarketplaceRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BlindAuctionMarketplace.Contract.BlindAuctionMarketplaceCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BlindAuctionMarketplace *BlindAuctionMarketplaceRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BlindAuctionMarketplace.Contract.BlindAuctionMarketplaceTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BlindAuctionMarketplace *BlindAuctionMarketplaceRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BlindAuctionMarketplace.Contract.BlindAuctionMarketplaceTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BlindAuctionMarketplace *BlindAuctionMarketplaceCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BlindAuctionMarketplace.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BlindAuctionMarketplace *BlindAuctionMarketplaceTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BlindAuctionMarketplace.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BlindAuctionMarketplace *BlindAuctionMarketplaceTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BlindAuctionMarketplace.Contract.contract.Transact(opts, method, params...)
}

// ADMINROLE is a free data retrieval call binding the contract method 0x75b238fc.
//
// Solidity: function ADMIN_ROLE() view returns(bytes32)
func (_BlindAuctionMarketplace *BlindAuctionMarketplaceCaller) ADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _BlindAuctionMarketplace.contract.Call(opts, &out, "ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ADMINROLE is a free data retrieval call binding the contract method 0x75b238fc.
//
// Solidity: function ADMIN_ROLE() view returns(bytes32)
func (_BlindAuctionMarketplace *BlindAuctionMarketplaceSession) ADMINROLE() ([32]byte, error) {
	return _BlindAuctionMarketplace.Contract.ADMINROLE(&_BlindAuctionMarketplace.CallOpts)
}

// ADMINROLE is a free data retrieval call binding the contract method 0x75b238fc.
//
// Solidity: function ADMIN_ROLE() view returns(bytes32)
func (_BlindAuctionMarketplace *BlindAuctionMarketplaceCallerSession) ADMINROLE() ([32]byte, error) {
	return _BlindAuctionMarketplace.Contract.ADMINROLE(&_BlindAuctionMarketplace.CallOpts)
}

// CANCELROLE is a free data retrieval call binding the contract method 0x13c27ca7.
//
// Solidity: function CANCEL_ROLE() view returns(bytes32)
func (_BlindAuctionMarketplace *BlindAuctionMarketplaceCaller) CANCELROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _BlindAuctionMarketplace.contract.Call(opts, &out, "CANCEL_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// CANCELROLE is a free data retrieval call binding the contract method 0x13c27ca7.
//
// Solidity: function CANCEL_ROLE() view returns(bytes32)
func (_BlindAuctionMarketplace *BlindAuctionMarketplaceSession) CANCELROLE() ([32]byte, error) {
	return _BlindAuctionMarketplace.Contract.CANCELROLE(&_BlindAuctionMarketplace.CallOpts)
}

// CANCELROLE is a free data retrieval call binding the contract method 0x13c27ca7.
//
// Solidity: function CANCEL_ROLE() view returns(bytes32)
func (_BlindAuctionMarketplace *BlindAuctionMarketplaceCallerSession) CANCELROLE() ([32]byte, error) {
	return _BlindAuctionMarketplace.Contract.CANCELROLE(&_BlindAuctionMarketplace.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_BlindAuctionMarketplace *BlindAuctionMarketplaceCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _BlindAuctionMarketplace.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_BlindAuctionMarketplace *BlindAuctionMarketplaceSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _BlindAuctionMarketplace.Contract.DEFAULTADMINROLE(&_BlindAuctionMarketplace.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_BlindAuctionMarketplace *BlindAuctionMarketplaceCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _BlindAuctionMarketplace.Contract.DEFAULTADMINROLE(&_BlindAuctionMarketplace.CallOpts)
}

// ERC721Interface is a free data retrieval call binding the contract method 0x2b4c32be.
//
// Solidity: function ERC721_Interface() view returns(bytes4)
func (_BlindAuctionMarketplace *BlindAuctionMarketplaceCaller) ERC721Interface(opts *bind.CallOpts) ([4]byte, error) {
	var out []interface{}
	err := _BlindAuctionMarketplace.contract.Call(opts, &out, "ERC721_Interface")

	if err != nil {
		return *new([4]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)

	return out0, err

}

// ERC721Interface is a free data retrieval call binding the contract method 0x2b4c32be.
//
// Solidity: function ERC721_Interface() view returns(bytes4)
func (_BlindAuctionMarketplace *BlindAuctionMarketplaceSession) ERC721Interface() ([4]byte, error) {
	return _BlindAuctionMarketplace.Contract.ERC721Interface(&_BlindAuctionMarketplace.CallOpts)
}

// ERC721Interface is a free data retrieval call binding the contract method 0x2b4c32be.
//
// Solidity: function ERC721_Interface() view returns(bytes4)
func (_BlindAuctionMarketplace *BlindAuctionMarketplaceCallerSession) ERC721Interface() ([4]byte, error) {
	return _BlindAuctionMarketplace.Contract.ERC721Interface(&_BlindAuctionMarketplace.CallOpts)
}

// PAUSERROLE is a free data retrieval call binding the contract method 0xe63ab1e9.
//
// Solidity: function PAUSER_ROLE() view returns(bytes32)
func (_BlindAuctionMarketplace *BlindAuctionMarketplaceCaller) PAUSERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _BlindAuctionMarketplace.contract.Call(opts, &out, "PAUSER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// PAUSERROLE is a free data retrieval call binding the contract method 0xe63ab1e9.
//
// Solidity: function PAUSER_ROLE() view returns(bytes32)
func (_BlindAuctionMarketplace *BlindAuctionMarketplaceSession) PAUSERROLE() ([32]byte, error) {
	return _BlindAuctionMarketplace.Contract.PAUSERROLE(&_BlindAuctionMarketplace.CallOpts)
}

// PAUSERROLE is a free data retrieval call binding the contract method 0xe63ab1e9.
//
// Solidity: function PAUSER_ROLE() view returns(bytes32)
func (_BlindAuctionMarketplace *BlindAuctionMarketplaceCallerSession) PAUSERROLE() ([32]byte, error) {
	return _BlindAuctionMarketplace.Contract.PAUSERROLE(&_BlindAuctionMarketplace.CallOpts)
}

// AcceptedToken is a free data retrieval call binding the contract method 0x451c3d80.
//
// Solidity: function acceptedToken() view returns(address)
func (_BlindAuctionMarketplace *BlindAuctionMarketplaceCaller) AcceptedToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BlindAuctionMarketplace.contract.Call(opts, &out, "acceptedToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AcceptedToken is a free data retrieval call binding the contract method 0x451c3d80.
//
// Solidity: function acceptedToken() view returns(address)
func (_BlindAuctionMarketplace *BlindAuctionMarketplaceSession) AcceptedToken() (common.Address, error) {
	return _BlindAuctionMarketplace.Contract.AcceptedToken(&_BlindAuctionMarketplace.CallOpts)
}

// AcceptedToken is a free data retrieval call binding the contract method 0x451c3d80.
//
// Solidity: function acceptedToken() view returns(address)
func (_BlindAuctionMarketplace *BlindAuctionMarketplaceCallerSession) AcceptedToken() (common.Address, error) {
	return _BlindAuctionMarketplace.Contract.AcceptedToken(&_BlindAuctionMarketplace.CallOpts)
}

// Beneficary is a free data retrieval call binding the contract method 0xcbfda1c5.
//
// Solidity: function beneficary() view returns(address)
func (_BlindAuctionMarketplace *BlindAuctionMarketplaceCaller) Beneficary(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BlindAuctionMarketplace.contract.Call(opts, &out, "beneficary")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Beneficary is a free data retrieval call binding the contract method 0xcbfda1c5.
//
// Solidity: function beneficary() view returns(address)
func (_BlindAuctionMarketplace *BlindAuctionMarketplaceSession) Beneficary() (common.Address, error) {
	return _BlindAuctionMarketplace.Contract.Beneficary(&_BlindAuctionMarketplace.CallOpts)
}

// Beneficary is a free data retrieval call binding the contract method 0xcbfda1c5.
//
// Solidity: function beneficary() view returns(address)
func (_BlindAuctionMarketplace *BlindAuctionMarketplaceCallerSession) Beneficary() (common.Address, error) {
	return _BlindAuctionMarketplace.Contract.Beneficary(&_BlindAuctionMarketplace.CallOpts)
}

// CheckEnded is a free data retrieval call binding the contract method 0x56479889.
//
// Solidity: function checkEnded(bytes32 auctionId, bytes32 nftAsset) view returns()
func (_BlindAuctionMarketplace *BlindAuctionMarketplaceCaller) CheckEnded(opts *bind.CallOpts, auctionId [32]byte, nftAsset [32]byte) error {
	var out []interface{}
	err := _BlindAuctionMarketplace.contract.Call(opts, &out, "checkEnded", auctionId, nftAsset)

	if err != nil {
		return err
	}

	return err

}

// CheckEnded is a free data retrieval call binding the contract method 0x56479889.
//
// Solidity: function checkEnded(bytes32 auctionId, bytes32 nftAsset) view returns()
func (_BlindAuctionMarketplace *BlindAuctionMarketplaceSession) CheckEnded(auctionId [32]byte, nftAsset [32]byte) error {
	return _BlindAuctionMarketplace.Contract.CheckEnded(&_BlindAuctionMarketplace.CallOpts, auctionId, nftAsset)
}

// CheckEnded is a free data retrieval call binding the contract method 0x56479889.
//
// Solidity: function checkEnded(bytes32 auctionId, bytes32 nftAsset) view returns()
func (_BlindAuctionMarketplace *BlindAuctionMarketplaceCallerSession) CheckEnded(auctionId [32]byte, nftAsset [32]byte) error {
	return _BlindAuctionMarketplace.Contract.CheckEnded(&_BlindAuctionMarketplace.CallOpts, auctionId, nftAsset)
}

// CheckExisted is a free data retrieval call binding the contract method 0x00c5c3c3.
//
// Solidity: function checkExisted(bytes32 auctionId) view returns()
func (_BlindAuctionMarketplace *BlindAuctionMarketplaceCaller) CheckExisted(opts *bind.CallOpts, auctionId [32]byte) error {
	var out []interface{}
	err := _BlindAuctionMarketplace.contract.Call(opts, &out, "checkExisted", auctionId)

	if err != nil {
		return err
	}

	return err

}

// CheckExisted is a free data retrieval call binding the contract method 0x00c5c3c3.
//
// Solidity: function checkExisted(bytes32 auctionId) view returns()
func (_BlindAuctionMarketplace *BlindAuctionMarketplaceSession) CheckExisted(auctionId [32]byte) error {
	return _BlindAuctionMarketplace.Contract.CheckExisted(&_BlindAuctionMarketplace.CallOpts, auctionId)
}

// CheckExisted is a free data retrieval call binding the contract method 0x00c5c3c3.
//
// Solidity: function checkExisted(bytes32 auctionId) view returns()
func (_BlindAuctionMarketplace *BlindAuctionMarketplaceCallerSession) CheckExisted(auctionId [32]byte) error {
	return _BlindAuctionMarketplace.Contract.CheckExisted(&_BlindAuctionMarketplace.CallOpts, auctionId)
}

// CheckRunning is a free data retrieval call binding the contract method 0x076f6dc7.
//
// Solidity: function checkRunning(bytes32 nftAsset, bytes32 auctionId) view returns()
func (_BlindAuctionMarketplace *BlindAuctionMarketplaceCaller) CheckRunning(opts *bind.CallOpts, nftAsset [32]byte, auctionId [32]byte) error {
	var out []interface{}
	err := _BlindAuctionMarketplace.contract.Call(opts, &out, "checkRunning", nftAsset, auctionId)

	if err != nil {
		return err
	}

	return err

}

// CheckRunning is a free data retrieval call binding the contract method 0x076f6dc7.
//
// Solidity: function checkRunning(bytes32 nftAsset, bytes32 auctionId) view returns()
func (_BlindAuctionMarketplace *BlindAuctionMarketplaceSession) CheckRunning(nftAsset [32]byte, auctionId [32]byte) error {
	return _BlindAuctionMarketplace.Contract.CheckRunning(&_BlindAuctionMarketplace.CallOpts, nftAsset, auctionId)
}

// CheckRunning is a free data retrieval call binding the contract method 0x076f6dc7.
//
// Solidity: function checkRunning(bytes32 nftAsset, bytes32 auctionId) view returns()
func (_BlindAuctionMarketplace *BlindAuctionMarketplaceCallerSession) CheckRunning(nftAsset [32]byte, auctionId [32]byte) error {
	return _BlindAuctionMarketplace.Contract.CheckRunning(&_BlindAuctionMarketplace.CallOpts, nftAsset, auctionId)
}

// GetBlackListStatus is a free data retrieval call binding the contract method 0x59bf1abe.
//
// Solidity: function getBlackListStatus(address _maker) view returns(bool)
func (_BlindAuctionMarketplace *BlindAuctionMarketplaceCaller) GetBlackListStatus(opts *bind.CallOpts, _maker common.Address) (bool, error) {
	var out []interface{}
	err := _BlindAuctionMarketplace.contract.Call(opts, &out, "getBlackListStatus", _maker)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// GetBlackListStatus is a free data retrieval call binding the contract method 0x59bf1abe.
//
// Solidity: function getBlackListStatus(address _maker) view returns(bool)
func (_BlindAuctionMarketplace *BlindAuctionMarketplaceSession) GetBlackListStatus(_maker common.Address) (bool, error) {
	return _BlindAuctionMarketplace.Contract.GetBlackListStatus(&_BlindAuctionMarketplace.CallOpts, _maker)
}

// GetBlackListStatus is a free data retrieval call binding the contract method 0x59bf1abe.
//
// Solidity: function getBlackListStatus(address _maker) view returns(bool)
func (_BlindAuctionMarketplace *BlindAuctionMarketplaceCallerSession) GetBlackListStatus(_maker common.Address) (bool, error) {
	return _BlindAuctionMarketplace.Contract.GetBlackListStatus(&_BlindAuctionMarketplace.CallOpts, _maker)
}

// GetBlindBid is a free data retrieval call binding the contract method 0x3af4b422.
//
// Solidity: function getBlindBid(bytes32 auctionId, address bidder) view returns(bytes32[])
func (_BlindAuctionMarketplace *BlindAuctionMarketplaceCaller) GetBlindBid(opts *bind.CallOpts, auctionId [32]byte, bidder common.Address) ([][32]byte, error) {
	var out []interface{}
	err := _BlindAuctionMarketplace.contract.Call(opts, &out, "getBlindBid", auctionId, bidder)

	if err != nil {
		return *new([][32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][32]byte)).(*[][32]byte)

	return out0, err

}

// GetBlindBid is a free data retrieval call binding the contract method 0x3af4b422.
//
// Solidity: function getBlindBid(bytes32 auctionId, address bidder) view returns(bytes32[])
func (_BlindAuctionMarketplace *BlindAuctionMarketplaceSession) GetBlindBid(auctionId [32]byte, bidder common.Address) ([][32]byte, error) {
	return _BlindAuctionMarketplace.Contract.GetBlindBid(&_BlindAuctionMarketplace.CallOpts, auctionId, bidder)
}

// GetBlindBid is a free data retrieval call binding the contract method 0x3af4b422.
//
// Solidity: function getBlindBid(bytes32 auctionId, address bidder) view returns(bytes32[])
func (_BlindAuctionMarketplace *BlindAuctionMarketplaceCallerSession) GetBlindBid(auctionId [32]byte, bidder common.Address) ([][32]byte, error) {
	return _BlindAuctionMarketplace.Contract.GetBlindBid(&_BlindAuctionMarketplace.CallOpts, auctionId, bidder)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_BlindAuctionMarketplace *BlindAuctionMarketplaceCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _BlindAuctionMarketplace.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_BlindAuctionMarketplace *BlindAuctionMarketplaceSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _BlindAuctionMarketplace.Contract.GetRoleAdmin(&_BlindAuctionMarketplace.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_BlindAuctionMarketplace *BlindAuctionMarketplaceCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _BlindAuctionMarketplace.Contract.GetRoleAdmin(&_BlindAuctionMarketplace.CallOpts, role)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_BlindAuctionMarketplace *BlindAuctionMarketplaceCaller) GetRoleMember(opts *bind.CallOpts, role [32]byte, index *big.Int) (common.Address, error) {
	var out []interface{}
	err := _BlindAuctionMarketplace.contract.Call(opts, &out, "getRoleMember", role, index)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_BlindAuctionMarketplace *BlindAuctionMarketplaceSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _BlindAuctionMarketplace.Contract.GetRoleMember(&_BlindAuctionMarketplace.CallOpts, role, index)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_BlindAuctionMarketplace *BlindAuctionMarketplaceCallerSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _BlindAuctionMarketplace.Contract.GetRoleMember(&_BlindAuctionMarketplace.CallOpts, role, index)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_BlindAuctionMarketplace *BlindAuctionMarketplaceCaller) GetRoleMemberCount(opts *bind.CallOpts, role [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _BlindAuctionMarketplace.contract.Call(opts, &out, "getRoleMemberCount", role)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_BlindAuctionMarketplace *BlindAuctionMarketplaceSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _BlindAuctionMarketplace.Contract.GetRoleMemberCount(&_BlindAuctionMarketplace.CallOpts, role)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_BlindAuctionMarketplace *BlindAuctionMarketplaceCallerSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _BlindAuctionMarketplace.Contract.GetRoleMemberCount(&_BlindAuctionMarketplace.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_BlindAuctionMarketplace *BlindAuctionMarketplaceCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _BlindAuctionMarketplace.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_BlindAuctionMarketplace *BlindAuctionMarketplaceSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _BlindAuctionMarketplace.Contract.HasRole(&_BlindAuctionMarketplace.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_BlindAuctionMarketplace *BlindAuctionMarketplaceCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _BlindAuctionMarketplace.Contract.HasRole(&_BlindAuctionMarketplace.CallOpts, role, account)
}

// IsBlackListed is a free data retrieval call binding the contract method 0xe47d6060.
//
// Solidity: function isBlackListed(address ) view returns(bool)
func (_BlindAuctionMarketplace *BlindAuctionMarketplaceCaller) IsBlackListed(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _BlindAuctionMarketplace.contract.Call(opts, &out, "isBlackListed", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsBlackListed is a free data retrieval call binding the contract method 0xe47d6060.
//
// Solidity: function isBlackListed(address ) view returns(bool)
func (_BlindAuctionMarketplace *BlindAuctionMarketplaceSession) IsBlackListed(arg0 common.Address) (bool, error) {
	return _BlindAuctionMarketplace.Contract.IsBlackListed(&_BlindAuctionMarketplace.CallOpts, arg0)
}

// IsBlackListed is a free data retrieval call binding the contract method 0xe47d6060.
//
// Solidity: function isBlackListed(address ) view returns(bool)
func (_BlindAuctionMarketplace *BlindAuctionMarketplaceCallerSession) IsBlackListed(arg0 common.Address) (bool, error) {
	return _BlindAuctionMarketplace.Contract.IsBlackListed(&_BlindAuctionMarketplace.CallOpts, arg0)
}

// MarketplaceStorage is a free data retrieval call binding the contract method 0x46b3aec6.
//
// Solidity: function marketplaceStorage() view returns(address)
func (_BlindAuctionMarketplace *BlindAuctionMarketplaceCaller) MarketplaceStorage(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BlindAuctionMarketplace.contract.Call(opts, &out, "marketplaceStorage")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// MarketplaceStorage is a free data retrieval call binding the contract method 0x46b3aec6.
//
// Solidity: function marketplaceStorage() view returns(address)
func (_BlindAuctionMarketplace *BlindAuctionMarketplaceSession) MarketplaceStorage() (common.Address, error) {
	return _BlindAuctionMarketplace.Contract.MarketplaceStorage(&_BlindAuctionMarketplace.CallOpts)
}

// MarketplaceStorage is a free data retrieval call binding the contract method 0x46b3aec6.
//
// Solidity: function marketplaceStorage() view returns(address)
func (_BlindAuctionMarketplace *BlindAuctionMarketplaceCallerSession) MarketplaceStorage() (common.Address, error) {
	return _BlindAuctionMarketplace.Contract.MarketplaceStorage(&_BlindAuctionMarketplace.CallOpts)
}

// MinStageDuration is a free data retrieval call binding the contract method 0xb42cf929.
//
// Solidity: function minStageDuration() view returns(uint256)
func (_BlindAuctionMarketplace *BlindAuctionMarketplaceCaller) MinStageDuration(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BlindAuctionMarketplace.contract.Call(opts, &out, "minStageDuration")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinStageDuration is a free data retrieval call binding the contract method 0xb42cf929.
//
// Solidity: function minStageDuration() view returns(uint256)
func (_BlindAuctionMarketplace *BlindAuctionMarketplaceSession) MinStageDuration() (*big.Int, error) {
	return _BlindAuctionMarketplace.Contract.MinStageDuration(&_BlindAuctionMarketplace.CallOpts)
}

// MinStageDuration is a free data retrieval call binding the contract method 0xb42cf929.
//
// Solidity: function minStageDuration() view returns(uint256)
func (_BlindAuctionMarketplace *BlindAuctionMarketplaceCallerSession) MinStageDuration() (*big.Int, error) {
	return _BlindAuctionMarketplace.Contract.MinStageDuration(&_BlindAuctionMarketplace.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BlindAuctionMarketplace *BlindAuctionMarketplaceCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BlindAuctionMarketplace.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BlindAuctionMarketplace *BlindAuctionMarketplaceSession) Owner() (common.Address, error) {
	return _BlindAuctionMarketplace.Contract.Owner(&_BlindAuctionMarketplace.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BlindAuctionMarketplace *BlindAuctionMarketplaceCallerSession) Owner() (common.Address, error) {
	return _BlindAuctionMarketplace.Contract.Owner(&_BlindAuctionMarketplace.CallOpts)
}

// OwnerCutPerMillion is a free data retrieval call binding the contract method 0xa01f79d4.
//
// Solidity: function ownerCutPerMillion() view returns(uint256)
func (_BlindAuctionMarketplace *BlindAuctionMarketplaceCaller) OwnerCutPerMillion(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BlindAuctionMarketplace.contract.Call(opts, &out, "ownerCutPerMillion")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// OwnerCutPerMillion is a free data retrieval call binding the contract method 0xa01f79d4.
//
// Solidity: function ownerCutPerMillion() view returns(uint256)
func (_BlindAuctionMarketplace *BlindAuctionMarketplaceSession) OwnerCutPerMillion() (*big.Int, error) {
	return _BlindAuctionMarketplace.Contract.OwnerCutPerMillion(&_BlindAuctionMarketplace.CallOpts)
}

// OwnerCutPerMillion is a free data retrieval call binding the contract method 0xa01f79d4.
//
// Solidity: function ownerCutPerMillion() view returns(uint256)
func (_BlindAuctionMarketplace *BlindAuctionMarketplaceCallerSession) OwnerCutPerMillion() (*big.Int, error) {
	return _BlindAuctionMarketplace.Contract.OwnerCutPerMillion(&_BlindAuctionMarketplace.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_BlindAuctionMarketplace *BlindAuctionMarketplaceCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _BlindAuctionMarketplace.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_BlindAuctionMarketplace *BlindAuctionMarketplaceSession) Paused() (bool, error) {
	return _BlindAuctionMarketplace.Contract.Paused(&_BlindAuctionMarketplace.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_BlindAuctionMarketplace *BlindAuctionMarketplaceCallerSession) Paused() (bool, error) {
	return _BlindAuctionMarketplace.Contract.Paused(&_BlindAuctionMarketplace.CallOpts)
}

// PublicationFeeInWei is a free data retrieval call binding the contract method 0xae4f1198.
//
// Solidity: function publicationFeeInWei() view returns(uint256)
func (_BlindAuctionMarketplace *BlindAuctionMarketplaceCaller) PublicationFeeInWei(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BlindAuctionMarketplace.contract.Call(opts, &out, "publicationFeeInWei")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PublicationFeeInWei is a free data retrieval call binding the contract method 0xae4f1198.
//
// Solidity: function publicationFeeInWei() view returns(uint256)
func (_BlindAuctionMarketplace *BlindAuctionMarketplaceSession) PublicationFeeInWei() (*big.Int, error) {
	return _BlindAuctionMarketplace.Contract.PublicationFeeInWei(&_BlindAuctionMarketplace.CallOpts)
}

// PublicationFeeInWei is a free data retrieval call binding the contract method 0xae4f1198.
//
// Solidity: function publicationFeeInWei() view returns(uint256)
func (_BlindAuctionMarketplace *BlindAuctionMarketplaceCallerSession) PublicationFeeInWei() (*big.Int, error) {
	return _BlindAuctionMarketplace.Contract.PublicationFeeInWei(&_BlindAuctionMarketplace.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_BlindAuctionMarketplace *BlindAuctionMarketplaceCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _BlindAuctionMarketplace.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_BlindAuctionMarketplace *BlindAuctionMarketplaceSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _BlindAuctionMarketplace.Contract.SupportsInterface(&_BlindAuctionMarketplace.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_BlindAuctionMarketplace *BlindAuctionMarketplaceCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _BlindAuctionMarketplace.Contract.SupportsInterface(&_BlindAuctionMarketplace.CallOpts, interfaceId)
}

// RevealBid is a paid mutator transaction binding the contract method 0x222bbd87.
//
// Solidity: function RevealBid(bytes32 blindAuctionId, bytes32 nftAsset, uint256[] _values, bytes32 secret) returns()
func (_BlindAuctionMarketplace *BlindAuctionMarketplaceTransactor) RevealBid(opts *bind.TransactOpts, blindAuctionId [32]byte, nftAsset [32]byte, _values []*big.Int, secret [32]byte) (*types.Transaction, error) {
	return _BlindAuctionMarketplace.contract.Transact(opts, "RevealBid", blindAuctionId, nftAsset, _values, secret)
}

// RevealBid is a paid mutator transaction binding the contract method 0x222bbd87.
//
// Solidity: function RevealBid(bytes32 blindAuctionId, bytes32 nftAsset, uint256[] _values, bytes32 secret) returns()
func (_BlindAuctionMarketplace *BlindAuctionMarketplaceSession) RevealBid(blindAuctionId [32]byte, nftAsset [32]byte, _values []*big.Int, secret [32]byte) (*types.Transaction, error) {
	return _BlindAuctionMarketplace.Contract.RevealBid(&_BlindAuctionMarketplace.TransactOpts, blindAuctionId, nftAsset, _values, secret)
}

// RevealBid is a paid mutator transaction binding the contract method 0x222bbd87.
//
// Solidity: function RevealBid(bytes32 blindAuctionId, bytes32 nftAsset, uint256[] _values, bytes32 secret) returns()
func (_BlindAuctionMarketplace *BlindAuctionMarketplaceTransactorSession) RevealBid(blindAuctionId [32]byte, nftAsset [32]byte, _values []*big.Int, secret [32]byte) (*types.Transaction, error) {
	return _BlindAuctionMarketplace.Contract.RevealBid(&_BlindAuctionMarketplace.TransactOpts, blindAuctionId, nftAsset, _values, secret)
}

// AddBlackList is a paid mutator transaction binding the contract method 0x0ecb93c0.
//
// Solidity: function addBlackList(address _evilUser) returns()
func (_BlindAuctionMarketplace *BlindAuctionMarketplaceTransactor) AddBlackList(opts *bind.TransactOpts, _evilUser common.Address) (*types.Transaction, error) {
	return _BlindAuctionMarketplace.contract.Transact(opts, "addBlackList", _evilUser)
}

// AddBlackList is a paid mutator transaction binding the contract method 0x0ecb93c0.
//
// Solidity: function addBlackList(address _evilUser) returns()
func (_BlindAuctionMarketplace *BlindAuctionMarketplaceSession) AddBlackList(_evilUser common.Address) (*types.Transaction, error) {
	return _BlindAuctionMarketplace.Contract.AddBlackList(&_BlindAuctionMarketplace.TransactOpts, _evilUser)
}

// AddBlackList is a paid mutator transaction binding the contract method 0x0ecb93c0.
//
// Solidity: function addBlackList(address _evilUser) returns()
func (_BlindAuctionMarketplace *BlindAuctionMarketplaceTransactorSession) AddBlackList(_evilUser common.Address) (*types.Transaction, error) {
	return _BlindAuctionMarketplace.Contract.AddBlackList(&_BlindAuctionMarketplace.TransactOpts, _evilUser)
}

// BidBlindAuction is a paid mutator transaction binding the contract method 0x314b3ee4.
//
// Solidity: function bidBlindAuction(bytes32 nftAsset, bytes32 blindAuctionId, bytes32 blindedBid, uint256 deposit) returns()
func (_BlindAuctionMarketplace *BlindAuctionMarketplaceTransactor) BidBlindAuction(opts *bind.TransactOpts, nftAsset [32]byte, blindAuctionId [32]byte, blindedBid [32]byte, deposit *big.Int) (*types.Transaction, error) {
	return _BlindAuctionMarketplace.contract.Transact(opts, "bidBlindAuction", nftAsset, blindAuctionId, blindedBid, deposit)
}

// BidBlindAuction is a paid mutator transaction binding the contract method 0x314b3ee4.
//
// Solidity: function bidBlindAuction(bytes32 nftAsset, bytes32 blindAuctionId, bytes32 blindedBid, uint256 deposit) returns()
func (_BlindAuctionMarketplace *BlindAuctionMarketplaceSession) BidBlindAuction(nftAsset [32]byte, blindAuctionId [32]byte, blindedBid [32]byte, deposit *big.Int) (*types.Transaction, error) {
	return _BlindAuctionMarketplace.Contract.BidBlindAuction(&_BlindAuctionMarketplace.TransactOpts, nftAsset, blindAuctionId, blindedBid, deposit)
}

// BidBlindAuction is a paid mutator transaction binding the contract method 0x314b3ee4.
//
// Solidity: function bidBlindAuction(bytes32 nftAsset, bytes32 blindAuctionId, bytes32 blindedBid, uint256 deposit) returns()
func (_BlindAuctionMarketplace *BlindAuctionMarketplaceTransactorSession) BidBlindAuction(nftAsset [32]byte, blindAuctionId [32]byte, blindedBid [32]byte, deposit *big.Int) (*types.Transaction, error) {
	return _BlindAuctionMarketplace.Contract.BidBlindAuction(&_BlindAuctionMarketplace.TransactOpts, nftAsset, blindAuctionId, blindedBid, deposit)
}

// CancelBlindAuction is a paid mutator transaction binding the contract method 0x3414bd26.
//
// Solidity: function cancelBlindAuction(bytes32 nftAsset, bytes32 blindAuctionId) returns()
func (_BlindAuctionMarketplace *BlindAuctionMarketplaceTransactor) CancelBlindAuction(opts *bind.TransactOpts, nftAsset [32]byte, blindAuctionId [32]byte) (*types.Transaction, error) {
	return _BlindAuctionMarketplace.contract.Transact(opts, "cancelBlindAuction", nftAsset, blindAuctionId)
}

// CancelBlindAuction is a paid mutator transaction binding the contract method 0x3414bd26.
//
// Solidity: function cancelBlindAuction(bytes32 nftAsset, bytes32 blindAuctionId) returns()
func (_BlindAuctionMarketplace *BlindAuctionMarketplaceSession) CancelBlindAuction(nftAsset [32]byte, blindAuctionId [32]byte) (*types.Transaction, error) {
	return _BlindAuctionMarketplace.Contract.CancelBlindAuction(&_BlindAuctionMarketplace.TransactOpts, nftAsset, blindAuctionId)
}

// CancelBlindAuction is a paid mutator transaction binding the contract method 0x3414bd26.
//
// Solidity: function cancelBlindAuction(bytes32 nftAsset, bytes32 blindAuctionId) returns()
func (_BlindAuctionMarketplace *BlindAuctionMarketplaceTransactorSession) CancelBlindAuction(nftAsset [32]byte, blindAuctionId [32]byte) (*types.Transaction, error) {
	return _BlindAuctionMarketplace.Contract.CancelBlindAuction(&_BlindAuctionMarketplace.TransactOpts, nftAsset, blindAuctionId)
}

// CreateBlindAuction is a paid mutator transaction binding the contract method 0x026967e0.
//
// Solidity: function createBlindAuction(address nftAddress, uint256 assetId, uint256 startPriceInWei, uint256 biddingEnd, uint256 revealEnd) returns()
func (_BlindAuctionMarketplace *BlindAuctionMarketplaceTransactor) CreateBlindAuction(opts *bind.TransactOpts, nftAddress common.Address, assetId *big.Int, startPriceInWei *big.Int, biddingEnd *big.Int, revealEnd *big.Int) (*types.Transaction, error) {
	return _BlindAuctionMarketplace.contract.Transact(opts, "createBlindAuction", nftAddress, assetId, startPriceInWei, biddingEnd, revealEnd)
}

// CreateBlindAuction is a paid mutator transaction binding the contract method 0x026967e0.
//
// Solidity: function createBlindAuction(address nftAddress, uint256 assetId, uint256 startPriceInWei, uint256 biddingEnd, uint256 revealEnd) returns()
func (_BlindAuctionMarketplace *BlindAuctionMarketplaceSession) CreateBlindAuction(nftAddress common.Address, assetId *big.Int, startPriceInWei *big.Int, biddingEnd *big.Int, revealEnd *big.Int) (*types.Transaction, error) {
	return _BlindAuctionMarketplace.Contract.CreateBlindAuction(&_BlindAuctionMarketplace.TransactOpts, nftAddress, assetId, startPriceInWei, biddingEnd, revealEnd)
}

// CreateBlindAuction is a paid mutator transaction binding the contract method 0x026967e0.
//
// Solidity: function createBlindAuction(address nftAddress, uint256 assetId, uint256 startPriceInWei, uint256 biddingEnd, uint256 revealEnd) returns()
func (_BlindAuctionMarketplace *BlindAuctionMarketplaceTransactorSession) CreateBlindAuction(nftAddress common.Address, assetId *big.Int, startPriceInWei *big.Int, biddingEnd *big.Int, revealEnd *big.Int) (*types.Transaction, error) {
	return _BlindAuctionMarketplace.Contract.CreateBlindAuction(&_BlindAuctionMarketplace.TransactOpts, nftAddress, assetId, startPriceInWei, biddingEnd, revealEnd)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_BlindAuctionMarketplace *BlindAuctionMarketplaceTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _BlindAuctionMarketplace.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_BlindAuctionMarketplace *BlindAuctionMarketplaceSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _BlindAuctionMarketplace.Contract.GrantRole(&_BlindAuctionMarketplace.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_BlindAuctionMarketplace *BlindAuctionMarketplaceTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _BlindAuctionMarketplace.Contract.GrantRole(&_BlindAuctionMarketplace.TransactOpts, role, account)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_BlindAuctionMarketplace *BlindAuctionMarketplaceTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BlindAuctionMarketplace.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_BlindAuctionMarketplace *BlindAuctionMarketplaceSession) Pause() (*types.Transaction, error) {
	return _BlindAuctionMarketplace.Contract.Pause(&_BlindAuctionMarketplace.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_BlindAuctionMarketplace *BlindAuctionMarketplaceTransactorSession) Pause() (*types.Transaction, error) {
	return _BlindAuctionMarketplace.Contract.Pause(&_BlindAuctionMarketplace.TransactOpts)
}

// RemoveBlackList is a paid mutator transaction binding the contract method 0xe4997dc5.
//
// Solidity: function removeBlackList(address _clearedUser) returns()
func (_BlindAuctionMarketplace *BlindAuctionMarketplaceTransactor) RemoveBlackList(opts *bind.TransactOpts, _clearedUser common.Address) (*types.Transaction, error) {
	return _BlindAuctionMarketplace.contract.Transact(opts, "removeBlackList", _clearedUser)
}

// RemoveBlackList is a paid mutator transaction binding the contract method 0xe4997dc5.
//
// Solidity: function removeBlackList(address _clearedUser) returns()
func (_BlindAuctionMarketplace *BlindAuctionMarketplaceSession) RemoveBlackList(_clearedUser common.Address) (*types.Transaction, error) {
	return _BlindAuctionMarketplace.Contract.RemoveBlackList(&_BlindAuctionMarketplace.TransactOpts, _clearedUser)
}

// RemoveBlackList is a paid mutator transaction binding the contract method 0xe4997dc5.
//
// Solidity: function removeBlackList(address _clearedUser) returns()
func (_BlindAuctionMarketplace *BlindAuctionMarketplaceTransactorSession) RemoveBlackList(_clearedUser common.Address) (*types.Transaction, error) {
	return _BlindAuctionMarketplace.Contract.RemoveBlackList(&_BlindAuctionMarketplace.TransactOpts, _clearedUser)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_BlindAuctionMarketplace *BlindAuctionMarketplaceTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BlindAuctionMarketplace.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_BlindAuctionMarketplace *BlindAuctionMarketplaceSession) RenounceOwnership() (*types.Transaction, error) {
	return _BlindAuctionMarketplace.Contract.RenounceOwnership(&_BlindAuctionMarketplace.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_BlindAuctionMarketplace *BlindAuctionMarketplaceTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _BlindAuctionMarketplace.Contract.RenounceOwnership(&_BlindAuctionMarketplace.TransactOpts)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_BlindAuctionMarketplace *BlindAuctionMarketplaceTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _BlindAuctionMarketplace.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_BlindAuctionMarketplace *BlindAuctionMarketplaceSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _BlindAuctionMarketplace.Contract.RenounceRole(&_BlindAuctionMarketplace.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_BlindAuctionMarketplace *BlindAuctionMarketplaceTransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _BlindAuctionMarketplace.Contract.RenounceRole(&_BlindAuctionMarketplace.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_BlindAuctionMarketplace *BlindAuctionMarketplaceTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _BlindAuctionMarketplace.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_BlindAuctionMarketplace *BlindAuctionMarketplaceSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _BlindAuctionMarketplace.Contract.RevokeRole(&_BlindAuctionMarketplace.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_BlindAuctionMarketplace *BlindAuctionMarketplaceTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _BlindAuctionMarketplace.Contract.RevokeRole(&_BlindAuctionMarketplace.TransactOpts, role, account)
}

// SetBeneficary is a paid mutator transaction binding the contract method 0x0db3cc05.
//
// Solidity: function setBeneficary(address _beneficary) returns()
func (_BlindAuctionMarketplace *BlindAuctionMarketplaceTransactor) SetBeneficary(opts *bind.TransactOpts, _beneficary common.Address) (*types.Transaction, error) {
	return _BlindAuctionMarketplace.contract.Transact(opts, "setBeneficary", _beneficary)
}

// SetBeneficary is a paid mutator transaction binding the contract method 0x0db3cc05.
//
// Solidity: function setBeneficary(address _beneficary) returns()
func (_BlindAuctionMarketplace *BlindAuctionMarketplaceSession) SetBeneficary(_beneficary common.Address) (*types.Transaction, error) {
	return _BlindAuctionMarketplace.Contract.SetBeneficary(&_BlindAuctionMarketplace.TransactOpts, _beneficary)
}

// SetBeneficary is a paid mutator transaction binding the contract method 0x0db3cc05.
//
// Solidity: function setBeneficary(address _beneficary) returns()
func (_BlindAuctionMarketplace *BlindAuctionMarketplaceTransactorSession) SetBeneficary(_beneficary common.Address) (*types.Transaction, error) {
	return _BlindAuctionMarketplace.Contract.SetBeneficary(&_BlindAuctionMarketplace.TransactOpts, _beneficary)
}

// SetMinStageDuration is a paid mutator transaction binding the contract method 0xf6021884.
//
// Solidity: function setMinStageDuration(uint256 duration) returns()
func (_BlindAuctionMarketplace *BlindAuctionMarketplaceTransactor) SetMinStageDuration(opts *bind.TransactOpts, duration *big.Int) (*types.Transaction, error) {
	return _BlindAuctionMarketplace.contract.Transact(opts, "setMinStageDuration", duration)
}

// SetMinStageDuration is a paid mutator transaction binding the contract method 0xf6021884.
//
// Solidity: function setMinStageDuration(uint256 duration) returns()
func (_BlindAuctionMarketplace *BlindAuctionMarketplaceSession) SetMinStageDuration(duration *big.Int) (*types.Transaction, error) {
	return _BlindAuctionMarketplace.Contract.SetMinStageDuration(&_BlindAuctionMarketplace.TransactOpts, duration)
}

// SetMinStageDuration is a paid mutator transaction binding the contract method 0xf6021884.
//
// Solidity: function setMinStageDuration(uint256 duration) returns()
func (_BlindAuctionMarketplace *BlindAuctionMarketplaceTransactorSession) SetMinStageDuration(duration *big.Int) (*types.Transaction, error) {
	return _BlindAuctionMarketplace.Contract.SetMinStageDuration(&_BlindAuctionMarketplace.TransactOpts, duration)
}

// SetOwnerCutPerMillion is a paid mutator transaction binding the contract method 0x19dad16d.
//
// Solidity: function setOwnerCutPerMillion(uint256 _ownerCutPerMillion) returns()
func (_BlindAuctionMarketplace *BlindAuctionMarketplaceTransactor) SetOwnerCutPerMillion(opts *bind.TransactOpts, _ownerCutPerMillion *big.Int) (*types.Transaction, error) {
	return _BlindAuctionMarketplace.contract.Transact(opts, "setOwnerCutPerMillion", _ownerCutPerMillion)
}

// SetOwnerCutPerMillion is a paid mutator transaction binding the contract method 0x19dad16d.
//
// Solidity: function setOwnerCutPerMillion(uint256 _ownerCutPerMillion) returns()
func (_BlindAuctionMarketplace *BlindAuctionMarketplaceSession) SetOwnerCutPerMillion(_ownerCutPerMillion *big.Int) (*types.Transaction, error) {
	return _BlindAuctionMarketplace.Contract.SetOwnerCutPerMillion(&_BlindAuctionMarketplace.TransactOpts, _ownerCutPerMillion)
}

// SetOwnerCutPerMillion is a paid mutator transaction binding the contract method 0x19dad16d.
//
// Solidity: function setOwnerCutPerMillion(uint256 _ownerCutPerMillion) returns()
func (_BlindAuctionMarketplace *BlindAuctionMarketplaceTransactorSession) SetOwnerCutPerMillion(_ownerCutPerMillion *big.Int) (*types.Transaction, error) {
	return _BlindAuctionMarketplace.Contract.SetOwnerCutPerMillion(&_BlindAuctionMarketplace.TransactOpts, _ownerCutPerMillion)
}

// SetPublicationFee is a paid mutator transaction binding the contract method 0xaf8996f1.
//
// Solidity: function setPublicationFee(uint256 _publicationFee) returns()
func (_BlindAuctionMarketplace *BlindAuctionMarketplaceTransactor) SetPublicationFee(opts *bind.TransactOpts, _publicationFee *big.Int) (*types.Transaction, error) {
	return _BlindAuctionMarketplace.contract.Transact(opts, "setPublicationFee", _publicationFee)
}

// SetPublicationFee is a paid mutator transaction binding the contract method 0xaf8996f1.
//
// Solidity: function setPublicationFee(uint256 _publicationFee) returns()
func (_BlindAuctionMarketplace *BlindAuctionMarketplaceSession) SetPublicationFee(_publicationFee *big.Int) (*types.Transaction, error) {
	return _BlindAuctionMarketplace.Contract.SetPublicationFee(&_BlindAuctionMarketplace.TransactOpts, _publicationFee)
}

// SetPublicationFee is a paid mutator transaction binding the contract method 0xaf8996f1.
//
// Solidity: function setPublicationFee(uint256 _publicationFee) returns()
func (_BlindAuctionMarketplace *BlindAuctionMarketplaceTransactorSession) SetPublicationFee(_publicationFee *big.Int) (*types.Transaction, error) {
	return _BlindAuctionMarketplace.Contract.SetPublicationFee(&_BlindAuctionMarketplace.TransactOpts, _publicationFee)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_BlindAuctionMarketplace *BlindAuctionMarketplaceTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _BlindAuctionMarketplace.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_BlindAuctionMarketplace *BlindAuctionMarketplaceSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _BlindAuctionMarketplace.Contract.TransferOwnership(&_BlindAuctionMarketplace.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_BlindAuctionMarketplace *BlindAuctionMarketplaceTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _BlindAuctionMarketplace.Contract.TransferOwnership(&_BlindAuctionMarketplace.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_BlindAuctionMarketplace *BlindAuctionMarketplaceTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BlindAuctionMarketplace.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_BlindAuctionMarketplace *BlindAuctionMarketplaceSession) Unpause() (*types.Transaction, error) {
	return _BlindAuctionMarketplace.Contract.Unpause(&_BlindAuctionMarketplace.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_BlindAuctionMarketplace *BlindAuctionMarketplaceTransactorSession) Unpause() (*types.Transaction, error) {
	return _BlindAuctionMarketplace.Contract.Unpause(&_BlindAuctionMarketplace.TransactOpts)
}

// UpdateStorageAddress is a paid mutator transaction binding the contract method 0x889e2129.
//
// Solidity: function updateStorageAddress(address _marketplaceStorage) returns()
func (_BlindAuctionMarketplace *BlindAuctionMarketplaceTransactor) UpdateStorageAddress(opts *bind.TransactOpts, _marketplaceStorage common.Address) (*types.Transaction, error) {
	return _BlindAuctionMarketplace.contract.Transact(opts, "updateStorageAddress", _marketplaceStorage)
}

// UpdateStorageAddress is a paid mutator transaction binding the contract method 0x889e2129.
//
// Solidity: function updateStorageAddress(address _marketplaceStorage) returns()
func (_BlindAuctionMarketplace *BlindAuctionMarketplaceSession) UpdateStorageAddress(_marketplaceStorage common.Address) (*types.Transaction, error) {
	return _BlindAuctionMarketplace.Contract.UpdateStorageAddress(&_BlindAuctionMarketplace.TransactOpts, _marketplaceStorage)
}

// UpdateStorageAddress is a paid mutator transaction binding the contract method 0x889e2129.
//
// Solidity: function updateStorageAddress(address _marketplaceStorage) returns()
func (_BlindAuctionMarketplace *BlindAuctionMarketplaceTransactorSession) UpdateStorageAddress(_marketplaceStorage common.Address) (*types.Transaction, error) {
	return _BlindAuctionMarketplace.Contract.UpdateStorageAddress(&_BlindAuctionMarketplace.TransactOpts, _marketplaceStorage)
}

// Withdraw is a paid mutator transaction binding the contract method 0x347ea9d6.
//
// Solidity: function withdraw(bytes32 blindAuctionId, uint256 assetId, address nftAddress, bytes32 nftAsset) returns()
func (_BlindAuctionMarketplace *BlindAuctionMarketplaceTransactor) Withdraw(opts *bind.TransactOpts, blindAuctionId [32]byte, assetId *big.Int, nftAddress common.Address, nftAsset [32]byte) (*types.Transaction, error) {
	return _BlindAuctionMarketplace.contract.Transact(opts, "withdraw", blindAuctionId, assetId, nftAddress, nftAsset)
}

// Withdraw is a paid mutator transaction binding the contract method 0x347ea9d6.
//
// Solidity: function withdraw(bytes32 blindAuctionId, uint256 assetId, address nftAddress, bytes32 nftAsset) returns()
func (_BlindAuctionMarketplace *BlindAuctionMarketplaceSession) Withdraw(blindAuctionId [32]byte, assetId *big.Int, nftAddress common.Address, nftAsset [32]byte) (*types.Transaction, error) {
	return _BlindAuctionMarketplace.Contract.Withdraw(&_BlindAuctionMarketplace.TransactOpts, blindAuctionId, assetId, nftAddress, nftAsset)
}

// Withdraw is a paid mutator transaction binding the contract method 0x347ea9d6.
//
// Solidity: function withdraw(bytes32 blindAuctionId, uint256 assetId, address nftAddress, bytes32 nftAsset) returns()
func (_BlindAuctionMarketplace *BlindAuctionMarketplaceTransactorSession) Withdraw(blindAuctionId [32]byte, assetId *big.Int, nftAddress common.Address, nftAsset [32]byte) (*types.Transaction, error) {
	return _BlindAuctionMarketplace.Contract.Withdraw(&_BlindAuctionMarketplace.TransactOpts, blindAuctionId, assetId, nftAddress, nftAsset)
}

// BlindAuctionMarketplaceAddedBlackListIterator is returned from FilterAddedBlackList and is used to iterate over the raw logs and unpacked data for AddedBlackList events raised by the BlindAuctionMarketplace contract.
type BlindAuctionMarketplaceAddedBlackListIterator struct {
	Event *BlindAuctionMarketplaceAddedBlackList // Event containing the contract specifics and raw log

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
func (it *BlindAuctionMarketplaceAddedBlackListIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BlindAuctionMarketplaceAddedBlackList)
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
		it.Event = new(BlindAuctionMarketplaceAddedBlackList)
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
func (it *BlindAuctionMarketplaceAddedBlackListIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BlindAuctionMarketplaceAddedBlackListIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BlindAuctionMarketplaceAddedBlackList represents a AddedBlackList event raised by the BlindAuctionMarketplace contract.
type BlindAuctionMarketplaceAddedBlackList struct {
	User common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterAddedBlackList is a free log retrieval operation binding the contract event 0x42e160154868087d6bfdc0ca23d96a1c1cfa32f1b72ba9ba27b69b98a0d819dc.
//
// Solidity: event AddedBlackList(address _user)
func (_BlindAuctionMarketplace *BlindAuctionMarketplaceFilterer) FilterAddedBlackList(opts *bind.FilterOpts) (*BlindAuctionMarketplaceAddedBlackListIterator, error) {

	logs, sub, err := _BlindAuctionMarketplace.contract.FilterLogs(opts, "AddedBlackList")
	if err != nil {
		return nil, err
	}
	return &BlindAuctionMarketplaceAddedBlackListIterator{contract: _BlindAuctionMarketplace.contract, event: "AddedBlackList", logs: logs, sub: sub}, nil
}

// WatchAddedBlackList is a free log subscription operation binding the contract event 0x42e160154868087d6bfdc0ca23d96a1c1cfa32f1b72ba9ba27b69b98a0d819dc.
//
// Solidity: event AddedBlackList(address _user)
func (_BlindAuctionMarketplace *BlindAuctionMarketplaceFilterer) WatchAddedBlackList(opts *bind.WatchOpts, sink chan<- *BlindAuctionMarketplaceAddedBlackList) (event.Subscription, error) {

	logs, sub, err := _BlindAuctionMarketplace.contract.WatchLogs(opts, "AddedBlackList")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BlindAuctionMarketplaceAddedBlackList)
				if err := _BlindAuctionMarketplace.contract.UnpackLog(event, "AddedBlackList", log); err != nil {
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
func (_BlindAuctionMarketplace *BlindAuctionMarketplaceFilterer) ParseAddedBlackList(log types.Log) (*BlindAuctionMarketplaceAddedBlackList, error) {
	event := new(BlindAuctionMarketplaceAddedBlackList)
	if err := _BlindAuctionMarketplace.contract.UnpackLog(event, "AddedBlackList", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BlindAuctionMarketplaceAuctionCancelledSuccessfulIterator is returned from FilterAuctionCancelledSuccessful and is used to iterate over the raw logs and unpacked data for AuctionCancelledSuccessful events raised by the BlindAuctionMarketplace contract.
type BlindAuctionMarketplaceAuctionCancelledSuccessfulIterator struct {
	Event *BlindAuctionMarketplaceAuctionCancelledSuccessful // Event containing the contract specifics and raw log

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
func (it *BlindAuctionMarketplaceAuctionCancelledSuccessfulIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BlindAuctionMarketplaceAuctionCancelledSuccessful)
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
		it.Event = new(BlindAuctionMarketplaceAuctionCancelledSuccessful)
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
func (it *BlindAuctionMarketplaceAuctionCancelledSuccessfulIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BlindAuctionMarketplaceAuctionCancelledSuccessfulIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BlindAuctionMarketplaceAuctionCancelledSuccessful represents a AuctionCancelledSuccessful event raised by the BlindAuctionMarketplace contract.
type BlindAuctionMarketplaceAuctionCancelledSuccessful struct {
	Canceller common.Address
	AuctionId [32]byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterAuctionCancelledSuccessful is a free log retrieval operation binding the contract event 0xa021239bb373fef19aa7c7ef798961ab038ebf95bd85d546b731d25855db190e.
//
// Solidity: event AuctionCancelledSuccessful(address canceller, bytes32 auctionId)
func (_BlindAuctionMarketplace *BlindAuctionMarketplaceFilterer) FilterAuctionCancelledSuccessful(opts *bind.FilterOpts) (*BlindAuctionMarketplaceAuctionCancelledSuccessfulIterator, error) {

	logs, sub, err := _BlindAuctionMarketplace.contract.FilterLogs(opts, "AuctionCancelledSuccessful")
	if err != nil {
		return nil, err
	}
	return &BlindAuctionMarketplaceAuctionCancelledSuccessfulIterator{contract: _BlindAuctionMarketplace.contract, event: "AuctionCancelledSuccessful", logs: logs, sub: sub}, nil
}

// WatchAuctionCancelledSuccessful is a free log subscription operation binding the contract event 0xa021239bb373fef19aa7c7ef798961ab038ebf95bd85d546b731d25855db190e.
//
// Solidity: event AuctionCancelledSuccessful(address canceller, bytes32 auctionId)
func (_BlindAuctionMarketplace *BlindAuctionMarketplaceFilterer) WatchAuctionCancelledSuccessful(opts *bind.WatchOpts, sink chan<- *BlindAuctionMarketplaceAuctionCancelledSuccessful) (event.Subscription, error) {

	logs, sub, err := _BlindAuctionMarketplace.contract.WatchLogs(opts, "AuctionCancelledSuccessful")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BlindAuctionMarketplaceAuctionCancelledSuccessful)
				if err := _BlindAuctionMarketplace.contract.UnpackLog(event, "AuctionCancelledSuccessful", log); err != nil {
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
func (_BlindAuctionMarketplace *BlindAuctionMarketplaceFilterer) ParseAuctionCancelledSuccessful(log types.Log) (*BlindAuctionMarketplaceAuctionCancelledSuccessful, error) {
	event := new(BlindAuctionMarketplaceAuctionCancelledSuccessful)
	if err := _BlindAuctionMarketplace.contract.UnpackLog(event, "AuctionCancelledSuccessful", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BlindAuctionMarketplaceAuctionRefundSuccessfulIterator is returned from FilterAuctionRefundSuccessful and is used to iterate over the raw logs and unpacked data for AuctionRefundSuccessful events raised by the BlindAuctionMarketplace contract.
type BlindAuctionMarketplaceAuctionRefundSuccessfulIterator struct {
	Event *BlindAuctionMarketplaceAuctionRefundSuccessful // Event containing the contract specifics and raw log

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
func (it *BlindAuctionMarketplaceAuctionRefundSuccessfulIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BlindAuctionMarketplaceAuctionRefundSuccessful)
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
		it.Event = new(BlindAuctionMarketplaceAuctionRefundSuccessful)
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
func (it *BlindAuctionMarketplaceAuctionRefundSuccessfulIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BlindAuctionMarketplaceAuctionRefundSuccessfulIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BlindAuctionMarketplaceAuctionRefundSuccessful represents a AuctionRefundSuccessful event raised by the BlindAuctionMarketplace contract.
type BlindAuctionMarketplaceAuctionRefundSuccessful struct {
	Bidder    common.Address
	AuctionId [32]byte
	Value     *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterAuctionRefundSuccessful is a free log retrieval operation binding the contract event 0x5f4b88832be3746851e4d0ce6129f89fe39ddcf5a4b0204699724c459baf7d9d.
//
// Solidity: event AuctionRefundSuccessful(address bidder, bytes32 auctionId, uint256 value)
func (_BlindAuctionMarketplace *BlindAuctionMarketplaceFilterer) FilterAuctionRefundSuccessful(opts *bind.FilterOpts) (*BlindAuctionMarketplaceAuctionRefundSuccessfulIterator, error) {

	logs, sub, err := _BlindAuctionMarketplace.contract.FilterLogs(opts, "AuctionRefundSuccessful")
	if err != nil {
		return nil, err
	}
	return &BlindAuctionMarketplaceAuctionRefundSuccessfulIterator{contract: _BlindAuctionMarketplace.contract, event: "AuctionRefundSuccessful", logs: logs, sub: sub}, nil
}

// WatchAuctionRefundSuccessful is a free log subscription operation binding the contract event 0x5f4b88832be3746851e4d0ce6129f89fe39ddcf5a4b0204699724c459baf7d9d.
//
// Solidity: event AuctionRefundSuccessful(address bidder, bytes32 auctionId, uint256 value)
func (_BlindAuctionMarketplace *BlindAuctionMarketplaceFilterer) WatchAuctionRefundSuccessful(opts *bind.WatchOpts, sink chan<- *BlindAuctionMarketplaceAuctionRefundSuccessful) (event.Subscription, error) {

	logs, sub, err := _BlindAuctionMarketplace.contract.WatchLogs(opts, "AuctionRefundSuccessful")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BlindAuctionMarketplaceAuctionRefundSuccessful)
				if err := _BlindAuctionMarketplace.contract.UnpackLog(event, "AuctionRefundSuccessful", log); err != nil {
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
func (_BlindAuctionMarketplace *BlindAuctionMarketplaceFilterer) ParseAuctionRefundSuccessful(log types.Log) (*BlindAuctionMarketplaceAuctionRefundSuccessful, error) {
	event := new(BlindAuctionMarketplaceAuctionRefundSuccessful)
	if err := _BlindAuctionMarketplace.contract.UnpackLog(event, "AuctionRefundSuccessful", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BlindAuctionMarketplaceBlindAuctionBidSuccessfulIterator is returned from FilterBlindAuctionBidSuccessful and is used to iterate over the raw logs and unpacked data for BlindAuctionBidSuccessful events raised by the BlindAuctionMarketplace contract.
type BlindAuctionMarketplaceBlindAuctionBidSuccessfulIterator struct {
	Event *BlindAuctionMarketplaceBlindAuctionBidSuccessful // Event containing the contract specifics and raw log

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
func (it *BlindAuctionMarketplaceBlindAuctionBidSuccessfulIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BlindAuctionMarketplaceBlindAuctionBidSuccessful)
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
		it.Event = new(BlindAuctionMarketplaceBlindAuctionBidSuccessful)
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
func (it *BlindAuctionMarketplaceBlindAuctionBidSuccessfulIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BlindAuctionMarketplaceBlindAuctionBidSuccessfulIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BlindAuctionMarketplaceBlindAuctionBidSuccessful represents a BlindAuctionBidSuccessful event raised by the BlindAuctionMarketplace contract.
type BlindAuctionMarketplaceBlindAuctionBidSuccessful struct {
	Sender         common.Address
	BlindAuctionId [32]byte
	BlindedBid     [32]byte
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterBlindAuctionBidSuccessful is a free log retrieval operation binding the contract event 0x42412bd3e1b349d0fc9c3518c85d09f611398d80842c5d3e36449b76dca01b6b.
//
// Solidity: event BlindAuctionBidSuccessful(address sender, bytes32 blindAuctionId, bytes32 blindedBid)
func (_BlindAuctionMarketplace *BlindAuctionMarketplaceFilterer) FilterBlindAuctionBidSuccessful(opts *bind.FilterOpts) (*BlindAuctionMarketplaceBlindAuctionBidSuccessfulIterator, error) {

	logs, sub, err := _BlindAuctionMarketplace.contract.FilterLogs(opts, "BlindAuctionBidSuccessful")
	if err != nil {
		return nil, err
	}
	return &BlindAuctionMarketplaceBlindAuctionBidSuccessfulIterator{contract: _BlindAuctionMarketplace.contract, event: "BlindAuctionBidSuccessful", logs: logs, sub: sub}, nil
}

// WatchBlindAuctionBidSuccessful is a free log subscription operation binding the contract event 0x42412bd3e1b349d0fc9c3518c85d09f611398d80842c5d3e36449b76dca01b6b.
//
// Solidity: event BlindAuctionBidSuccessful(address sender, bytes32 blindAuctionId, bytes32 blindedBid)
func (_BlindAuctionMarketplace *BlindAuctionMarketplaceFilterer) WatchBlindAuctionBidSuccessful(opts *bind.WatchOpts, sink chan<- *BlindAuctionMarketplaceBlindAuctionBidSuccessful) (event.Subscription, error) {

	logs, sub, err := _BlindAuctionMarketplace.contract.WatchLogs(opts, "BlindAuctionBidSuccessful")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BlindAuctionMarketplaceBlindAuctionBidSuccessful)
				if err := _BlindAuctionMarketplace.contract.UnpackLog(event, "BlindAuctionBidSuccessful", log); err != nil {
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
func (_BlindAuctionMarketplace *BlindAuctionMarketplaceFilterer) ParseBlindAuctionBidSuccessful(log types.Log) (*BlindAuctionMarketplaceBlindAuctionBidSuccessful, error) {
	event := new(BlindAuctionMarketplaceBlindAuctionBidSuccessful)
	if err := _BlindAuctionMarketplace.contract.UnpackLog(event, "BlindAuctionBidSuccessful", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BlindAuctionMarketplaceBlindAuctionCreatedSuccessfulIterator is returned from FilterBlindAuctionCreatedSuccessful and is used to iterate over the raw logs and unpacked data for BlindAuctionCreatedSuccessful events raised by the BlindAuctionMarketplace contract.
type BlindAuctionMarketplaceBlindAuctionCreatedSuccessfulIterator struct {
	Event *BlindAuctionMarketplaceBlindAuctionCreatedSuccessful // Event containing the contract specifics and raw log

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
func (it *BlindAuctionMarketplaceBlindAuctionCreatedSuccessfulIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BlindAuctionMarketplaceBlindAuctionCreatedSuccessful)
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
		it.Event = new(BlindAuctionMarketplaceBlindAuctionCreatedSuccessful)
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
func (it *BlindAuctionMarketplaceBlindAuctionCreatedSuccessfulIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BlindAuctionMarketplaceBlindAuctionCreatedSuccessfulIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BlindAuctionMarketplaceBlindAuctionCreatedSuccessful represents a BlindAuctionCreatedSuccessful event raised by the BlindAuctionMarketplace contract.
type BlindAuctionMarketplaceBlindAuctionCreatedSuccessful struct {
	AssetOwner      common.Address
	NftAddress      common.Address
	BlindAuctionId  [32]byte
	AssetId         *big.Int
	BiddingEnd      *big.Int
	RevealEnd       *big.Int
	StartPriceInWei *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterBlindAuctionCreatedSuccessful is a free log retrieval operation binding the contract event 0xde549e5fcbd4d262f76b1f3a6362777a2e669ce8a2a7b5460e00c7a38341baf7.
//
// Solidity: event BlindAuctionCreatedSuccessful(address assetOwner, address nftAddress, bytes32 blindAuctionId, uint256 assetId, uint256 biddingEnd, uint256 revealEnd, uint256 startPriceInWei)
func (_BlindAuctionMarketplace *BlindAuctionMarketplaceFilterer) FilterBlindAuctionCreatedSuccessful(opts *bind.FilterOpts) (*BlindAuctionMarketplaceBlindAuctionCreatedSuccessfulIterator, error) {

	logs, sub, err := _BlindAuctionMarketplace.contract.FilterLogs(opts, "BlindAuctionCreatedSuccessful")
	if err != nil {
		return nil, err
	}
	return &BlindAuctionMarketplaceBlindAuctionCreatedSuccessfulIterator{contract: _BlindAuctionMarketplace.contract, event: "BlindAuctionCreatedSuccessful", logs: logs, sub: sub}, nil
}

// WatchBlindAuctionCreatedSuccessful is a free log subscription operation binding the contract event 0xde549e5fcbd4d262f76b1f3a6362777a2e669ce8a2a7b5460e00c7a38341baf7.
//
// Solidity: event BlindAuctionCreatedSuccessful(address assetOwner, address nftAddress, bytes32 blindAuctionId, uint256 assetId, uint256 biddingEnd, uint256 revealEnd, uint256 startPriceInWei)
func (_BlindAuctionMarketplace *BlindAuctionMarketplaceFilterer) WatchBlindAuctionCreatedSuccessful(opts *bind.WatchOpts, sink chan<- *BlindAuctionMarketplaceBlindAuctionCreatedSuccessful) (event.Subscription, error) {

	logs, sub, err := _BlindAuctionMarketplace.contract.WatchLogs(opts, "BlindAuctionCreatedSuccessful")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BlindAuctionMarketplaceBlindAuctionCreatedSuccessful)
				if err := _BlindAuctionMarketplace.contract.UnpackLog(event, "BlindAuctionCreatedSuccessful", log); err != nil {
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

// ParseBlindAuctionCreatedSuccessful is a log parse operation binding the contract event 0xde549e5fcbd4d262f76b1f3a6362777a2e669ce8a2a7b5460e00c7a38341baf7.
//
// Solidity: event BlindAuctionCreatedSuccessful(address assetOwner, address nftAddress, bytes32 blindAuctionId, uint256 assetId, uint256 biddingEnd, uint256 revealEnd, uint256 startPriceInWei)
func (_BlindAuctionMarketplace *BlindAuctionMarketplaceFilterer) ParseBlindAuctionCreatedSuccessful(log types.Log) (*BlindAuctionMarketplaceBlindAuctionCreatedSuccessful, error) {
	event := new(BlindAuctionMarketplaceBlindAuctionCreatedSuccessful)
	if err := _BlindAuctionMarketplace.contract.UnpackLog(event, "BlindAuctionCreatedSuccessful", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BlindAuctionMarketplaceChangedOwnerCutPerMillionIterator is returned from FilterChangedOwnerCutPerMillion and is used to iterate over the raw logs and unpacked data for ChangedOwnerCutPerMillion events raised by the BlindAuctionMarketplace contract.
type BlindAuctionMarketplaceChangedOwnerCutPerMillionIterator struct {
	Event *BlindAuctionMarketplaceChangedOwnerCutPerMillion // Event containing the contract specifics and raw log

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
func (it *BlindAuctionMarketplaceChangedOwnerCutPerMillionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BlindAuctionMarketplaceChangedOwnerCutPerMillion)
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
		it.Event = new(BlindAuctionMarketplaceChangedOwnerCutPerMillion)
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
func (it *BlindAuctionMarketplaceChangedOwnerCutPerMillionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BlindAuctionMarketplaceChangedOwnerCutPerMillionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BlindAuctionMarketplaceChangedOwnerCutPerMillion represents a ChangedOwnerCutPerMillion event raised by the BlindAuctionMarketplace contract.
type BlindAuctionMarketplaceChangedOwnerCutPerMillion struct {
	OwnerCutPerMillion *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterChangedOwnerCutPerMillion is a free log retrieval operation binding the contract event 0xfa406a120a9e7f2b332bfb7a43d3bf1c3f079262202907a6b69c94b2821a02c6.
//
// Solidity: event ChangedOwnerCutPerMillion(uint256 ownerCutPerMillion)
func (_BlindAuctionMarketplace *BlindAuctionMarketplaceFilterer) FilterChangedOwnerCutPerMillion(opts *bind.FilterOpts) (*BlindAuctionMarketplaceChangedOwnerCutPerMillionIterator, error) {

	logs, sub, err := _BlindAuctionMarketplace.contract.FilterLogs(opts, "ChangedOwnerCutPerMillion")
	if err != nil {
		return nil, err
	}
	return &BlindAuctionMarketplaceChangedOwnerCutPerMillionIterator{contract: _BlindAuctionMarketplace.contract, event: "ChangedOwnerCutPerMillion", logs: logs, sub: sub}, nil
}

// WatchChangedOwnerCutPerMillion is a free log subscription operation binding the contract event 0xfa406a120a9e7f2b332bfb7a43d3bf1c3f079262202907a6b69c94b2821a02c6.
//
// Solidity: event ChangedOwnerCutPerMillion(uint256 ownerCutPerMillion)
func (_BlindAuctionMarketplace *BlindAuctionMarketplaceFilterer) WatchChangedOwnerCutPerMillion(opts *bind.WatchOpts, sink chan<- *BlindAuctionMarketplaceChangedOwnerCutPerMillion) (event.Subscription, error) {

	logs, sub, err := _BlindAuctionMarketplace.contract.WatchLogs(opts, "ChangedOwnerCutPerMillion")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BlindAuctionMarketplaceChangedOwnerCutPerMillion)
				if err := _BlindAuctionMarketplace.contract.UnpackLog(event, "ChangedOwnerCutPerMillion", log); err != nil {
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
func (_BlindAuctionMarketplace *BlindAuctionMarketplaceFilterer) ParseChangedOwnerCutPerMillion(log types.Log) (*BlindAuctionMarketplaceChangedOwnerCutPerMillion, error) {
	event := new(BlindAuctionMarketplaceChangedOwnerCutPerMillion)
	if err := _BlindAuctionMarketplace.contract.UnpackLog(event, "ChangedOwnerCutPerMillion", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BlindAuctionMarketplaceChangedPublicationFeeIterator is returned from FilterChangedPublicationFee and is used to iterate over the raw logs and unpacked data for ChangedPublicationFee events raised by the BlindAuctionMarketplace contract.
type BlindAuctionMarketplaceChangedPublicationFeeIterator struct {
	Event *BlindAuctionMarketplaceChangedPublicationFee // Event containing the contract specifics and raw log

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
func (it *BlindAuctionMarketplaceChangedPublicationFeeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BlindAuctionMarketplaceChangedPublicationFee)
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
		it.Event = new(BlindAuctionMarketplaceChangedPublicationFee)
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
func (it *BlindAuctionMarketplaceChangedPublicationFeeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BlindAuctionMarketplaceChangedPublicationFeeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BlindAuctionMarketplaceChangedPublicationFee represents a ChangedPublicationFee event raised by the BlindAuctionMarketplace contract.
type BlindAuctionMarketplaceChangedPublicationFee struct {
	PublicationFee *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterChangedPublicationFee is a free log retrieval operation binding the contract event 0xe7fa8737293f41b5dfa0d5c3e552860a06275bed7015581b083c7be7003308ba.
//
// Solidity: event ChangedPublicationFee(uint256 publicationFee)
func (_BlindAuctionMarketplace *BlindAuctionMarketplaceFilterer) FilterChangedPublicationFee(opts *bind.FilterOpts) (*BlindAuctionMarketplaceChangedPublicationFeeIterator, error) {

	logs, sub, err := _BlindAuctionMarketplace.contract.FilterLogs(opts, "ChangedPublicationFee")
	if err != nil {
		return nil, err
	}
	return &BlindAuctionMarketplaceChangedPublicationFeeIterator{contract: _BlindAuctionMarketplace.contract, event: "ChangedPublicationFee", logs: logs, sub: sub}, nil
}

// WatchChangedPublicationFee is a free log subscription operation binding the contract event 0xe7fa8737293f41b5dfa0d5c3e552860a06275bed7015581b083c7be7003308ba.
//
// Solidity: event ChangedPublicationFee(uint256 publicationFee)
func (_BlindAuctionMarketplace *BlindAuctionMarketplaceFilterer) WatchChangedPublicationFee(opts *bind.WatchOpts, sink chan<- *BlindAuctionMarketplaceChangedPublicationFee) (event.Subscription, error) {

	logs, sub, err := _BlindAuctionMarketplace.contract.WatchLogs(opts, "ChangedPublicationFee")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BlindAuctionMarketplaceChangedPublicationFee)
				if err := _BlindAuctionMarketplace.contract.UnpackLog(event, "ChangedPublicationFee", log); err != nil {
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
func (_BlindAuctionMarketplace *BlindAuctionMarketplaceFilterer) ParseChangedPublicationFee(log types.Log) (*BlindAuctionMarketplaceChangedPublicationFee, error) {
	event := new(BlindAuctionMarketplaceChangedPublicationFee)
	if err := _BlindAuctionMarketplace.contract.UnpackLog(event, "ChangedPublicationFee", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BlindAuctionMarketplaceGrantAuctionRewardSuccessfulIterator is returned from FilterGrantAuctionRewardSuccessful and is used to iterate over the raw logs and unpacked data for GrantAuctionRewardSuccessful events raised by the BlindAuctionMarketplace contract.
type BlindAuctionMarketplaceGrantAuctionRewardSuccessfulIterator struct {
	Event *BlindAuctionMarketplaceGrantAuctionRewardSuccessful // Event containing the contract specifics and raw log

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
func (it *BlindAuctionMarketplaceGrantAuctionRewardSuccessfulIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BlindAuctionMarketplaceGrantAuctionRewardSuccessful)
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
		it.Event = new(BlindAuctionMarketplaceGrantAuctionRewardSuccessful)
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
func (it *BlindAuctionMarketplaceGrantAuctionRewardSuccessfulIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BlindAuctionMarketplaceGrantAuctionRewardSuccessfulIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BlindAuctionMarketplaceGrantAuctionRewardSuccessful represents a GrantAuctionRewardSuccessful event raised by the BlindAuctionMarketplace contract.
type BlindAuctionMarketplaceGrantAuctionRewardSuccessful struct {
	AuctionHighestBidder common.Address
	AuctionId            [32]byte
	AssetId              *big.Int
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterGrantAuctionRewardSuccessful is a free log retrieval operation binding the contract event 0xa307d88e125ed19583b02332c5bd74f0323774248b3f9cac5047c2433d9d8ddc.
//
// Solidity: event GrantAuctionRewardSuccessful(address auctionHighestBidder, bytes32 auctionId, uint256 assetId)
func (_BlindAuctionMarketplace *BlindAuctionMarketplaceFilterer) FilterGrantAuctionRewardSuccessful(opts *bind.FilterOpts) (*BlindAuctionMarketplaceGrantAuctionRewardSuccessfulIterator, error) {

	logs, sub, err := _BlindAuctionMarketplace.contract.FilterLogs(opts, "GrantAuctionRewardSuccessful")
	if err != nil {
		return nil, err
	}
	return &BlindAuctionMarketplaceGrantAuctionRewardSuccessfulIterator{contract: _BlindAuctionMarketplace.contract, event: "GrantAuctionRewardSuccessful", logs: logs, sub: sub}, nil
}

// WatchGrantAuctionRewardSuccessful is a free log subscription operation binding the contract event 0xa307d88e125ed19583b02332c5bd74f0323774248b3f9cac5047c2433d9d8ddc.
//
// Solidity: event GrantAuctionRewardSuccessful(address auctionHighestBidder, bytes32 auctionId, uint256 assetId)
func (_BlindAuctionMarketplace *BlindAuctionMarketplaceFilterer) WatchGrantAuctionRewardSuccessful(opts *bind.WatchOpts, sink chan<- *BlindAuctionMarketplaceGrantAuctionRewardSuccessful) (event.Subscription, error) {

	logs, sub, err := _BlindAuctionMarketplace.contract.WatchLogs(opts, "GrantAuctionRewardSuccessful")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BlindAuctionMarketplaceGrantAuctionRewardSuccessful)
				if err := _BlindAuctionMarketplace.contract.UnpackLog(event, "GrantAuctionRewardSuccessful", log); err != nil {
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
func (_BlindAuctionMarketplace *BlindAuctionMarketplaceFilterer) ParseGrantAuctionRewardSuccessful(log types.Log) (*BlindAuctionMarketplaceGrantAuctionRewardSuccessful, error) {
	event := new(BlindAuctionMarketplaceGrantAuctionRewardSuccessful)
	if err := _BlindAuctionMarketplace.contract.UnpackLog(event, "GrantAuctionRewardSuccessful", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BlindAuctionMarketplaceMarketplaceStorageUpdatedIterator is returned from FilterMarketplaceStorageUpdated and is used to iterate over the raw logs and unpacked data for MarketplaceStorageUpdated events raised by the BlindAuctionMarketplace contract.
type BlindAuctionMarketplaceMarketplaceStorageUpdatedIterator struct {
	Event *BlindAuctionMarketplaceMarketplaceStorageUpdated // Event containing the contract specifics and raw log

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
func (it *BlindAuctionMarketplaceMarketplaceStorageUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BlindAuctionMarketplaceMarketplaceStorageUpdated)
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
		it.Event = new(BlindAuctionMarketplaceMarketplaceStorageUpdated)
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
func (it *BlindAuctionMarketplaceMarketplaceStorageUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BlindAuctionMarketplaceMarketplaceStorageUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BlindAuctionMarketplaceMarketplaceStorageUpdated represents a MarketplaceStorageUpdated event raised by the BlindAuctionMarketplace contract.
type BlindAuctionMarketplaceMarketplaceStorageUpdated struct {
	MarketplaceStorage common.Address
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterMarketplaceStorageUpdated is a free log retrieval operation binding the contract event 0x19e07a962d0e1cab8b7d4be06b77a7225ca7ebb744994cb6546fdd815590dc9d.
//
// Solidity: event MarketplaceStorageUpdated(address _marketplaceStorage)
func (_BlindAuctionMarketplace *BlindAuctionMarketplaceFilterer) FilterMarketplaceStorageUpdated(opts *bind.FilterOpts) (*BlindAuctionMarketplaceMarketplaceStorageUpdatedIterator, error) {

	logs, sub, err := _BlindAuctionMarketplace.contract.FilterLogs(opts, "MarketplaceStorageUpdated")
	if err != nil {
		return nil, err
	}
	return &BlindAuctionMarketplaceMarketplaceStorageUpdatedIterator{contract: _BlindAuctionMarketplace.contract, event: "MarketplaceStorageUpdated", logs: logs, sub: sub}, nil
}

// WatchMarketplaceStorageUpdated is a free log subscription operation binding the contract event 0x19e07a962d0e1cab8b7d4be06b77a7225ca7ebb744994cb6546fdd815590dc9d.
//
// Solidity: event MarketplaceStorageUpdated(address _marketplaceStorage)
func (_BlindAuctionMarketplace *BlindAuctionMarketplaceFilterer) WatchMarketplaceStorageUpdated(opts *bind.WatchOpts, sink chan<- *BlindAuctionMarketplaceMarketplaceStorageUpdated) (event.Subscription, error) {

	logs, sub, err := _BlindAuctionMarketplace.contract.WatchLogs(opts, "MarketplaceStorageUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BlindAuctionMarketplaceMarketplaceStorageUpdated)
				if err := _BlindAuctionMarketplace.contract.UnpackLog(event, "MarketplaceStorageUpdated", log); err != nil {
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
func (_BlindAuctionMarketplace *BlindAuctionMarketplaceFilterer) ParseMarketplaceStorageUpdated(log types.Log) (*BlindAuctionMarketplaceMarketplaceStorageUpdated, error) {
	event := new(BlindAuctionMarketplaceMarketplaceStorageUpdated)
	if err := _BlindAuctionMarketplace.contract.UnpackLog(event, "MarketplaceStorageUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BlindAuctionMarketplaceOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the BlindAuctionMarketplace contract.
type BlindAuctionMarketplaceOwnershipTransferredIterator struct {
	Event *BlindAuctionMarketplaceOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *BlindAuctionMarketplaceOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BlindAuctionMarketplaceOwnershipTransferred)
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
		it.Event = new(BlindAuctionMarketplaceOwnershipTransferred)
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
func (it *BlindAuctionMarketplaceOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BlindAuctionMarketplaceOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BlindAuctionMarketplaceOwnershipTransferred represents a OwnershipTransferred event raised by the BlindAuctionMarketplace contract.
type BlindAuctionMarketplaceOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_BlindAuctionMarketplace *BlindAuctionMarketplaceFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*BlindAuctionMarketplaceOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _BlindAuctionMarketplace.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &BlindAuctionMarketplaceOwnershipTransferredIterator{contract: _BlindAuctionMarketplace.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_BlindAuctionMarketplace *BlindAuctionMarketplaceFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *BlindAuctionMarketplaceOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _BlindAuctionMarketplace.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BlindAuctionMarketplaceOwnershipTransferred)
				if err := _BlindAuctionMarketplace.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_BlindAuctionMarketplace *BlindAuctionMarketplaceFilterer) ParseOwnershipTransferred(log types.Log) (*BlindAuctionMarketplaceOwnershipTransferred, error) {
	event := new(BlindAuctionMarketplaceOwnershipTransferred)
	if err := _BlindAuctionMarketplace.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BlindAuctionMarketplacePausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the BlindAuctionMarketplace contract.
type BlindAuctionMarketplacePausedIterator struct {
	Event *BlindAuctionMarketplacePaused // Event containing the contract specifics and raw log

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
func (it *BlindAuctionMarketplacePausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BlindAuctionMarketplacePaused)
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
		it.Event = new(BlindAuctionMarketplacePaused)
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
func (it *BlindAuctionMarketplacePausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BlindAuctionMarketplacePausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BlindAuctionMarketplacePaused represents a Paused event raised by the BlindAuctionMarketplace contract.
type BlindAuctionMarketplacePaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_BlindAuctionMarketplace *BlindAuctionMarketplaceFilterer) FilterPaused(opts *bind.FilterOpts) (*BlindAuctionMarketplacePausedIterator, error) {

	logs, sub, err := _BlindAuctionMarketplace.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &BlindAuctionMarketplacePausedIterator{contract: _BlindAuctionMarketplace.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_BlindAuctionMarketplace *BlindAuctionMarketplaceFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *BlindAuctionMarketplacePaused) (event.Subscription, error) {

	logs, sub, err := _BlindAuctionMarketplace.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BlindAuctionMarketplacePaused)
				if err := _BlindAuctionMarketplace.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_BlindAuctionMarketplace *BlindAuctionMarketplaceFilterer) ParsePaused(log types.Log) (*BlindAuctionMarketplacePaused, error) {
	event := new(BlindAuctionMarketplacePaused)
	if err := _BlindAuctionMarketplace.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BlindAuctionMarketplaceRemovedBlackListIterator is returned from FilterRemovedBlackList and is used to iterate over the raw logs and unpacked data for RemovedBlackList events raised by the BlindAuctionMarketplace contract.
type BlindAuctionMarketplaceRemovedBlackListIterator struct {
	Event *BlindAuctionMarketplaceRemovedBlackList // Event containing the contract specifics and raw log

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
func (it *BlindAuctionMarketplaceRemovedBlackListIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BlindAuctionMarketplaceRemovedBlackList)
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
		it.Event = new(BlindAuctionMarketplaceRemovedBlackList)
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
func (it *BlindAuctionMarketplaceRemovedBlackListIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BlindAuctionMarketplaceRemovedBlackListIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BlindAuctionMarketplaceRemovedBlackList represents a RemovedBlackList event raised by the BlindAuctionMarketplace contract.
type BlindAuctionMarketplaceRemovedBlackList struct {
	User common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterRemovedBlackList is a free log retrieval operation binding the contract event 0xd7e9ec6e6ecd65492dce6bf513cd6867560d49544421d0783ddf06e76c24470c.
//
// Solidity: event RemovedBlackList(address _user)
func (_BlindAuctionMarketplace *BlindAuctionMarketplaceFilterer) FilterRemovedBlackList(opts *bind.FilterOpts) (*BlindAuctionMarketplaceRemovedBlackListIterator, error) {

	logs, sub, err := _BlindAuctionMarketplace.contract.FilterLogs(opts, "RemovedBlackList")
	if err != nil {
		return nil, err
	}
	return &BlindAuctionMarketplaceRemovedBlackListIterator{contract: _BlindAuctionMarketplace.contract, event: "RemovedBlackList", logs: logs, sub: sub}, nil
}

// WatchRemovedBlackList is a free log subscription operation binding the contract event 0xd7e9ec6e6ecd65492dce6bf513cd6867560d49544421d0783ddf06e76c24470c.
//
// Solidity: event RemovedBlackList(address _user)
func (_BlindAuctionMarketplace *BlindAuctionMarketplaceFilterer) WatchRemovedBlackList(opts *bind.WatchOpts, sink chan<- *BlindAuctionMarketplaceRemovedBlackList) (event.Subscription, error) {

	logs, sub, err := _BlindAuctionMarketplace.contract.WatchLogs(opts, "RemovedBlackList")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BlindAuctionMarketplaceRemovedBlackList)
				if err := _BlindAuctionMarketplace.contract.UnpackLog(event, "RemovedBlackList", log); err != nil {
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
func (_BlindAuctionMarketplace *BlindAuctionMarketplaceFilterer) ParseRemovedBlackList(log types.Log) (*BlindAuctionMarketplaceRemovedBlackList, error) {
	event := new(BlindAuctionMarketplaceRemovedBlackList)
	if err := _BlindAuctionMarketplace.contract.UnpackLog(event, "RemovedBlackList", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BlindAuctionMarketplaceRevealSuccessfulIterator is returned from FilterRevealSuccessful and is used to iterate over the raw logs and unpacked data for RevealSuccessful events raised by the BlindAuctionMarketplace contract.
type BlindAuctionMarketplaceRevealSuccessfulIterator struct {
	Event *BlindAuctionMarketplaceRevealSuccessful // Event containing the contract specifics and raw log

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
func (it *BlindAuctionMarketplaceRevealSuccessfulIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BlindAuctionMarketplaceRevealSuccessful)
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
		it.Event = new(BlindAuctionMarketplaceRevealSuccessful)
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
func (it *BlindAuctionMarketplaceRevealSuccessfulIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BlindAuctionMarketplaceRevealSuccessfulIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BlindAuctionMarketplaceRevealSuccessful represents a RevealSuccessful event raised by the BlindAuctionMarketplace contract.
type BlindAuctionMarketplaceRevealSuccessful struct {
	Sender         common.Address
	BlindAuctionId [32]byte
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterRevealSuccessful is a free log retrieval operation binding the contract event 0x61d4ab60470c97b371c5f91be47873283386bb3af4de728e8795ef5318e013d6.
//
// Solidity: event RevealSuccessful(address sender, bytes32 blindAuctionId)
func (_BlindAuctionMarketplace *BlindAuctionMarketplaceFilterer) FilterRevealSuccessful(opts *bind.FilterOpts) (*BlindAuctionMarketplaceRevealSuccessfulIterator, error) {

	logs, sub, err := _BlindAuctionMarketplace.contract.FilterLogs(opts, "RevealSuccessful")
	if err != nil {
		return nil, err
	}
	return &BlindAuctionMarketplaceRevealSuccessfulIterator{contract: _BlindAuctionMarketplace.contract, event: "RevealSuccessful", logs: logs, sub: sub}, nil
}

// WatchRevealSuccessful is a free log subscription operation binding the contract event 0x61d4ab60470c97b371c5f91be47873283386bb3af4de728e8795ef5318e013d6.
//
// Solidity: event RevealSuccessful(address sender, bytes32 blindAuctionId)
func (_BlindAuctionMarketplace *BlindAuctionMarketplaceFilterer) WatchRevealSuccessful(opts *bind.WatchOpts, sink chan<- *BlindAuctionMarketplaceRevealSuccessful) (event.Subscription, error) {

	logs, sub, err := _BlindAuctionMarketplace.contract.WatchLogs(opts, "RevealSuccessful")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BlindAuctionMarketplaceRevealSuccessful)
				if err := _BlindAuctionMarketplace.contract.UnpackLog(event, "RevealSuccessful", log); err != nil {
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
func (_BlindAuctionMarketplace *BlindAuctionMarketplaceFilterer) ParseRevealSuccessful(log types.Log) (*BlindAuctionMarketplaceRevealSuccessful, error) {
	event := new(BlindAuctionMarketplaceRevealSuccessful)
	if err := _BlindAuctionMarketplace.contract.UnpackLog(event, "RevealSuccessful", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BlindAuctionMarketplaceRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the BlindAuctionMarketplace contract.
type BlindAuctionMarketplaceRoleAdminChangedIterator struct {
	Event *BlindAuctionMarketplaceRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *BlindAuctionMarketplaceRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BlindAuctionMarketplaceRoleAdminChanged)
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
		it.Event = new(BlindAuctionMarketplaceRoleAdminChanged)
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
func (it *BlindAuctionMarketplaceRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BlindAuctionMarketplaceRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BlindAuctionMarketplaceRoleAdminChanged represents a RoleAdminChanged event raised by the BlindAuctionMarketplace contract.
type BlindAuctionMarketplaceRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_BlindAuctionMarketplace *BlindAuctionMarketplaceFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*BlindAuctionMarketplaceRoleAdminChangedIterator, error) {

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

	logs, sub, err := _BlindAuctionMarketplace.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &BlindAuctionMarketplaceRoleAdminChangedIterator{contract: _BlindAuctionMarketplace.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_BlindAuctionMarketplace *BlindAuctionMarketplaceFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *BlindAuctionMarketplaceRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _BlindAuctionMarketplace.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BlindAuctionMarketplaceRoleAdminChanged)
				if err := _BlindAuctionMarketplace.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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
func (_BlindAuctionMarketplace *BlindAuctionMarketplaceFilterer) ParseRoleAdminChanged(log types.Log) (*BlindAuctionMarketplaceRoleAdminChanged, error) {
	event := new(BlindAuctionMarketplaceRoleAdminChanged)
	if err := _BlindAuctionMarketplace.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BlindAuctionMarketplaceRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the BlindAuctionMarketplace contract.
type BlindAuctionMarketplaceRoleGrantedIterator struct {
	Event *BlindAuctionMarketplaceRoleGranted // Event containing the contract specifics and raw log

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
func (it *BlindAuctionMarketplaceRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BlindAuctionMarketplaceRoleGranted)
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
		it.Event = new(BlindAuctionMarketplaceRoleGranted)
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
func (it *BlindAuctionMarketplaceRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BlindAuctionMarketplaceRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BlindAuctionMarketplaceRoleGranted represents a RoleGranted event raised by the BlindAuctionMarketplace contract.
type BlindAuctionMarketplaceRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_BlindAuctionMarketplace *BlindAuctionMarketplaceFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*BlindAuctionMarketplaceRoleGrantedIterator, error) {

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

	logs, sub, err := _BlindAuctionMarketplace.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &BlindAuctionMarketplaceRoleGrantedIterator{contract: _BlindAuctionMarketplace.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_BlindAuctionMarketplace *BlindAuctionMarketplaceFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *BlindAuctionMarketplaceRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _BlindAuctionMarketplace.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BlindAuctionMarketplaceRoleGranted)
				if err := _BlindAuctionMarketplace.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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
func (_BlindAuctionMarketplace *BlindAuctionMarketplaceFilterer) ParseRoleGranted(log types.Log) (*BlindAuctionMarketplaceRoleGranted, error) {
	event := new(BlindAuctionMarketplaceRoleGranted)
	if err := _BlindAuctionMarketplace.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BlindAuctionMarketplaceRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the BlindAuctionMarketplace contract.
type BlindAuctionMarketplaceRoleRevokedIterator struct {
	Event *BlindAuctionMarketplaceRoleRevoked // Event containing the contract specifics and raw log

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
func (it *BlindAuctionMarketplaceRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BlindAuctionMarketplaceRoleRevoked)
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
		it.Event = new(BlindAuctionMarketplaceRoleRevoked)
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
func (it *BlindAuctionMarketplaceRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BlindAuctionMarketplaceRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BlindAuctionMarketplaceRoleRevoked represents a RoleRevoked event raised by the BlindAuctionMarketplace contract.
type BlindAuctionMarketplaceRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_BlindAuctionMarketplace *BlindAuctionMarketplaceFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*BlindAuctionMarketplaceRoleRevokedIterator, error) {

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

	logs, sub, err := _BlindAuctionMarketplace.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &BlindAuctionMarketplaceRoleRevokedIterator{contract: _BlindAuctionMarketplace.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_BlindAuctionMarketplace *BlindAuctionMarketplaceFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *BlindAuctionMarketplaceRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _BlindAuctionMarketplace.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BlindAuctionMarketplaceRoleRevoked)
				if err := _BlindAuctionMarketplace.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
func (_BlindAuctionMarketplace *BlindAuctionMarketplaceFilterer) ParseRoleRevoked(log types.Log) (*BlindAuctionMarketplaceRoleRevoked, error) {
	event := new(BlindAuctionMarketplaceRoleRevoked)
	if err := _BlindAuctionMarketplace.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BlindAuctionMarketplaceUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the BlindAuctionMarketplace contract.
type BlindAuctionMarketplaceUnpausedIterator struct {
	Event *BlindAuctionMarketplaceUnpaused // Event containing the contract specifics and raw log

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
func (it *BlindAuctionMarketplaceUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BlindAuctionMarketplaceUnpaused)
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
		it.Event = new(BlindAuctionMarketplaceUnpaused)
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
func (it *BlindAuctionMarketplaceUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BlindAuctionMarketplaceUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BlindAuctionMarketplaceUnpaused represents a Unpaused event raised by the BlindAuctionMarketplace contract.
type BlindAuctionMarketplaceUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_BlindAuctionMarketplace *BlindAuctionMarketplaceFilterer) FilterUnpaused(opts *bind.FilterOpts) (*BlindAuctionMarketplaceUnpausedIterator, error) {

	logs, sub, err := _BlindAuctionMarketplace.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &BlindAuctionMarketplaceUnpausedIterator{contract: _BlindAuctionMarketplace.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_BlindAuctionMarketplace *BlindAuctionMarketplaceFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *BlindAuctionMarketplaceUnpaused) (event.Subscription, error) {

	logs, sub, err := _BlindAuctionMarketplace.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BlindAuctionMarketplaceUnpaused)
				if err := _BlindAuctionMarketplace.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
func (_BlindAuctionMarketplace *BlindAuctionMarketplaceFilterer) ParseUnpaused(log types.Log) (*BlindAuctionMarketplaceUnpaused, error) {
	event := new(BlindAuctionMarketplaceUnpaused)
	if err := _BlindAuctionMarketplace.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}