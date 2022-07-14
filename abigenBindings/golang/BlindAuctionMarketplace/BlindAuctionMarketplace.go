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
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"acceptedToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"marketplaceStorage\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"beneficary\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"ownerCutPerMillion\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AuctionEnded\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidBiddingPrice\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidContract\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidCut\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidPrice\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidReveal\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidTime\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidWithdraw\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotBidYet\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotExisted\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotRunning\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotStarted\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RewardGranted\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"TooEarly\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"TooLate\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Unauthorized\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Unavailable\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"AddedBlackList\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"canceller\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"auctionId\",\"type\":\"bytes32\"}],\"name\":\"AuctionCancelledSuccessful\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"bidder\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"auctionId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"AuctionRefundSuccessful\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"blindAuctionId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"blindedBid\",\"type\":\"bytes32\"}],\"name\":\"BlindAuctionBidSuccessful\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"assetOwner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"nftAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"blindAuctionId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"assetId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"biddingEnd\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"revealEnd\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"startPriceInWei\",\"type\":\"uint256\"}],\"name\":\"BlindAuctionCreatedSuccessful\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"ownerCutPerMillion\",\"type\":\"uint256\"}],\"name\":\"ChangedOwnerCutPerMillion\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"publicationFee\",\"type\":\"uint256\"}],\"name\":\"ChangedPublicationFee\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"auctionHighestBidder\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"auctionId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"assetId\",\"type\":\"uint256\"}],\"name\":\"GrantAuctionRewardSuccessful\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_marketplaceStorage\",\"type\":\"address\"}],\"name\":\"MarketplaceStorageUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"RemovedBlackList\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"blindAuctionId\",\"type\":\"bytes32\"}],\"name\":\"RevealSuccessful\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"CANCEL_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ERC721_Interface\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PAUSER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"acceptedToken\",\"outputs\":[{\"internalType\":\"contractIBEP20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_evilUser\",\"type\":\"address\"}],\"name\":\"addBlackList\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"beneficary\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_maker\",\"type\":\"address\"}],\"name\":\"getBlackListStatus\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getRoleMember\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleMemberCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"isBlackListed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"marketplaceStorage\",\"outputs\":[{\"internalType\":\"contractIMarketplaceStorage\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minStageDuration\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ownerCutPerMillion\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"publicationFeeInWei\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_clearedUser\",\"type\":\"address\"}],\"name\":\"removeBlackList\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_beneficary\",\"type\":\"address\"}],\"name\":\"setBeneficary\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"}],\"name\":\"setMinStageDuration\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_ownerCutPerMillion\",\"type\":\"uint256\"}],\"name\":\"setOwnerCutPerMillion\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_publicationFee\",\"type\":\"uint256\"}],\"name\":\"setPublicationFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_marketplaceStorage\",\"type\":\"address\"}],\"name\":\"updateStorageAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"nftAddress\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"blindAuctionId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"nftAsset\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"assetId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startPriceInWei\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"biddingEnd\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"revealEnd\",\"type\":\"uint256\"}],\"name\":\"createAuction\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"nftAsset\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"blindAuctionId\",\"type\":\"bytes32\"}],\"name\":\"cancelAuction\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"nftAsset\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"blindAuctionId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"blindedBid\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"deposit\",\"type\":\"uint256\"}],\"name\":\"bid\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"nftAddress\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"nftAsset\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"blindAuctionId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"assetId\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"blindAuctionId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"nftAsset\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"secret\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[]\",\"name\":\"_values\",\"type\":\"uint256[]\"}],\"name\":\"reveal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"blindAuctionId\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"bidder\",\"type\":\"address\"}],\"name\":\"getBlindBid\",\"outputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"\",\"type\":\"bytes32[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"blindAuctionId\",\"type\":\"bytes32\"}],\"name\":\"checkExisted\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"nftAsset\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"blindAuctionId\",\"type\":\"bytes32\"}],\"name\":\"checkRunning\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60806040523480156200001157600080fd5b5060405162006d9638038062006d968339818101604052810190620000379190620013ef565b838284838180620000698173ffffffffffffffffffffffffffffffffffffffff166200042760201b62002baf1760201c565b620000ab576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401620000a29062001678565b60405180910390fd5b600081905082600260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505050506000600260146101000a81548160ff0219169083151502179055506200012f620001236200044a60201b60201c565b620004fd60201b60201c565b6200014f620001436200044a60201b60201c565b620005c360201b60201c565b6200017b8473ffffffffffffffffffffffffffffffffffffffff166200042760201b62002baf1760201c565b620001b2576040517f6eefed2000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b83600560006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555081600260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550620002586000801b6200024c6200044a60201b60201c565b6200078a60201b60201c565b620002997fa49807205ce4d355092ef5a8a18f56e8913cf4a201fbe287825b095693c217756200028d6200044a60201b60201c565b6200078a60201b60201c565b620002da7f9f959e00d95122f5cbd677010436cf273ef535b86b056afc172852144b9491d7620002ce6200044a60201b60201c565b6200078a60201b60201c565b6200031b7f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a6200030f6200044a60201b60201c565b6200078a60201b60201c565b6200036d7f9f959e00d95122f5cbd677010436cf273ef535b86b056afc172852144b9491d77fa49807205ce4d355092ef5a8a18f56e8913cf4a201fbe287825b095693c21775620007a060201b60201c565b620003bf7f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a7fa49807205ce4d355092ef5a8a18f56e8913cf4a201fbe287825b095693c21775620007a060201b60201c565b620003d0816200080360201b60201c565b600160088190555082600960006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550505050505050505062001a2c565b6000808273ffffffffffffffffffffffffffffffffffffffff163b119050919050565b60003073ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161415620004f657600080368080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f820116905080830192505050505050509050600080369050905073ffffffffffffffffffffffffffffffffffffffff818301511692505050620004fa565b3390505b90565b6000600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905081600360006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b620005d36200044a60201b60201c565b73ffffffffffffffffffffffffffffffffffffffff16620005f9620008b860201b60201c565b73ffffffffffffffffffffffffffffffffffffffff161462000652576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040162000649906200169a565b60405180910390fd5b620006847f9f959e00d95122f5cbd677010436cf273ef535b86b056afc172852144b9491d7826200078a60201b60201c565b620006b67f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a826200078a60201b60201c565b620006cb6000801b826200078a60201b60201c565b6200070c7f9f959e00d95122f5cbd677010436cf273ef535b86b056afc172852144b9491d762000700620008b860201b60201c565b620008e260201b60201c565b6200074d7f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a62000741620008b860201b60201c565b620008e260201b60201c565b620007716000801b62000765620008b860201b60201c565b620008e260201b60201c565b62000787816200091b60201b62002bd21760201c565b50565b6200079c828262000a3160201b60201c565b5050565b6000620007b38362000a7960201b60201c565b905081600080858152602001908152602001600020600101819055508181847fbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff60405160405180910390a4505050565b7fa49807205ce4d355092ef5a8a18f56e8913cf4a201fbe287825b095693c21775620008358162000a9860201b60201c565b620f4240821062000872576040517f7e0d5ce600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b816006819055507ffa406a120a9e7f2b332bfb7a43d3bf1c3f079262202907a6b69c94b2821a02c6600654604051620008ac9190620016bc565b60405180910390a15050565b6000600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b620008f38262000a7960201b60201c565b620009048162000a9860201b60201c565b62000916838362000abc60201b60201c565b505050565b6200092b6200044a60201b60201c565b73ffffffffffffffffffffffffffffffffffffffff1662000951620008b860201b60201c565b73ffffffffffffffffffffffffffffffffffffffff1614620009aa576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401620009a1906200169a565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16141562000a1d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040162000a149062001656565b60405180910390fd5b62000a2e81620004fd60201b60201c565b50565b62000a48828262000b0460201b62002cca1760201c565b62000a74816001600085815260200190815260200160002062000bf560201b62002daa1790919060201c565b505050565b6000806000838152602001908152602001600020600101549050919050565b62000ab98162000aad6200044a60201b60201c565b62000c2d60201b60201c565b50565b62000ad3828262000cf160201b62002dda1760201c565b62000aff816001600085815260200190815260200160002062000de360201b62002ebb1790919060201c565b505050565b62000b16828262000e1b60201b60201c565b62000bf157600160008084815260200190815260200160002060000160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff02191690831515021790555062000b966200044a60201b60201c565b73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16837f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a45b5050565b600062000c25836000018373ffffffffffffffffffffffffffffffffffffffff1660001b62000e8560201b60201c565b905092915050565b62000c3f828262000e1b60201b60201c565b62000ced5762000c728173ffffffffffffffffffffffffffffffffffffffff16601462000eff60201b62002eeb1760201c565b62000c8d8360001c602062000eff60201b62002eeb1760201c565b60405160200162000ca0929190620015ce565b6040516020818303038152906040526040517f08c379a000000000000000000000000000000000000000000000000000000000815260040162000ce4919062001610565b60405180910390fd5b5050565b62000d03828262000e1b60201b60201c565b1562000ddf57600080600084815260200190815260200160002060000160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff02191690831515021790555062000d846200044a60201b60201c565b73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16837ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b60405160405180910390a45b5050565b600062000e13836000018373ffffffffffffffffffffffffffffffffffffffff1660001b6200120e60201b60201c565b905092915050565b600080600084815260200190815260200160002060000160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16905092915050565b600062000e9983836200139e60201b60201c565b62000ef457826000018290806001815401808255809150506001900390600052602060002001600090919091909150558260000180549050836001016000848152602001908152602001600020819055506001905062000ef9565b600090505b92915050565b60606000600283600262000f1491906200175d565b62000f20919062001700565b67ffffffffffffffff81111562000f60577f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6040519080825280601f01601f19166020018201604052801562000f935781602001600182028036833780820191505090505b5090507f30000000000000000000000000000000000000000000000000000000000000008160008151811062000ff2577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a9053507f7800000000000000000000000000000000000000000000000000000000000000816001815181106200107d577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a90535060006001846002620010bf91906200175d565b620010cb919062001700565b90505b6001811115620011bd577f3031323334353637383961626364656600000000000000000000000000000000600f86166010811062001135577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b1a60f81b82828151811062001173577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a905350600485901c945080620011b5906200186d565b9050620010ce565b506000841462001204576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401620011fb9062001634565b60405180910390fd5b8091505092915050565b6000808360010160008481526020019081526020016000205490506000811462001392576000600182620012439190620017be565b90506000600186600001805490506200125d9190620017be565b90508181146200131b576000866000018281548110620012a6577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b9060005260206000200154905080876000018481548110620012f1577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b90600052602060002001819055508387600101600083815260200190815260200160002081905550505b8560000180548062001356577f4e487b7100000000000000000000000000000000000000000000000000000000600052603160045260246000fd5b60019003818190600052602060002001600090559055856001016000868152602001908152602001600020600090556001935050505062001398565b60009150505b92915050565b600080836001016000848152602001908152602001600020541415905092915050565b600081519050620013d281620019f8565b92915050565b600081519050620013e98162001a12565b92915050565b600080600080608085870312156200140657600080fd5b60006200141687828801620013c1565b94505060206200142987828801620013c1565b93505060406200143c87828801620013c1565b92505060606200144f87828801620013d8565b91505092959194509250565b60006200146882620016d9565b620014748185620016e4565b93506200148681856020860162001837565b6200149181620018cb565b840191505092915050565b6000620014a982620016d9565b620014b58185620016f5565b9350620014c781856020860162001837565b80840191505092915050565b6000620014e2602083620016e4565b9150620014ef82620018dc565b602082019050919050565b600062001509602683620016e4565b9150620015168262001905565b604082019050919050565b600062001530601083620016e4565b91506200153d8262001954565b602082019050919050565b600062001557602083620016e4565b915062001564826200197d565b602082019050919050565b60006200157e601783620016f5565b91506200158b82620019a6565b601782019050919050565b6000620015a5601183620016f5565b9150620015b282620019cf565b601182019050919050565b620015c8816200182d565b82525050565b6000620015db826200156f565b9150620015e982856200149c565b9150620015f68262001596565b91506200160482846200149c565b91508190509392505050565b600060208201905081810360008301526200162c81846200145b565b905092915050565b600060208201905081810360008301526200164f81620014d3565b9050919050565b600060208201905081810360008301526200167181620014fa565b9050919050565b60006020820190508181036000830152620016938162001521565b9050919050565b60006020820190508181036000830152620016b58162001548565b9050919050565b6000602082019050620016d36000830184620015bd565b92915050565b600081519050919050565b600082825260208201905092915050565b600081905092915050565b60006200170d826200182d565b91506200171a836200182d565b9250827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff038211156200175257620017516200189c565b5b828201905092915050565b60006200176a826200182d565b915062001777836200182d565b9250817fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0483118215151615620017b357620017b26200189c565b5b828202905092915050565b6000620017cb826200182d565b9150620017d8836200182d565b925082821015620017ee57620017ed6200189c565b5b828203905092915050565b600062001806826200180d565b9050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b60005b83811015620018575780820151818401526020810190506200183a565b8381111562001867576000848401525b50505050565b60006200187a826200182d565b915060008214156200189157620018906200189c565b5b600182039050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b6000601f19601f8301169050919050565b7f537472696e67733a20686578206c656e67746820696e73756666696369656e74600082015250565b7f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160008201527f6464726573730000000000000000000000000000000000000000000000000000602082015250565b7f496e76616c696420636f6e747261637400000000000000000000000000000000600082015250565b7f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572600082015250565b7f416363657373436f6e74726f6c3a206163636f756e7420000000000000000000600082015250565b7f206973206d697373696e6720726f6c6520000000000000000000000000000000600082015250565b62001a0381620017f9565b811462001a0f57600080fd5b50565b62001a1d816200182d565b811462001a2957600080fd5b50565b61535a8062001a3c6000396000f3fe608060405234801561001057600080fd5b506004361061025d5760003560e01c80638456cb5911610146578063b42cf929116100c3578063e47d606011610087578063e47d6060146106c0578063e4997dc5146106f0578063e63ab1e91461070c578063f2fde38b1461072a578063f3764ffa14610746578063f6021884146107625761025d565b8063b42cf9291461061c578063ca15c8731461063a578063cbfda1c51461066a578063d547741f14610688578063daa3d985146106a45761025d565b80639655fe1a1161010a5780639655fe1a1461058a578063a01f79d4146105a6578063a217fddf146105c4578063ae4f1198146105e2578063af8996f1146106005761025d565b80638456cb59146104e6578063889e2129146104f05780638da5cb5b1461050c5780639010d07c1461052a57806391d148541461055a5761025d565b80632f2ff15d116101df57806346b3aec6116101a357806346b3aec61461043657806359bf1abe146104545780635c975abb14610484578063715018a6146104a25780637544b61b146104ac57806375b238fc146104c85761025d565b80632f2ff15d146103a657806336568abe146103c25780633af4b422146103de5780633f4ba83a1461040e578063451c3d80146104185761025d565b80630ecb93c0116102265780630ecb93c01461030257806313c27ca71461031e57806319dad16d1461033c578063248a9ca3146103585780632b4c32be146103885761025d565b8062c5c3c31461026257806301ffc9a71461027e578063076f6dc7146102ae57806307c8e9cb146102ca5780630db3cc05146102e6575b600080fd5b61027c60048036038101906102779190614390565b61077e565b005b6102986004803603810190610293919061454b565b610862565b6040516102a59190614b92565b60405180910390f35b6102c860048036038101906102c391906143f5565b6108dc565b005b6102e460048036038101906102df9190614431565b6109cc565b005b61030060048036038101906102fb9190614200565b610fa0565b005b61031c60048036038101906103179190614200565b61100f565b005b61032661111d565b6040516103339190614bad565b60405180910390f35b6103566004803603810190610351919061459e565b611141565b005b610372600480360381019061036d9190614390565b6111eb565b60405161037f9190614bad565b60405180910390f35b61039061120a565b60405161039d9190614bf1565b60405180910390f35b6103c060048036038101906103bb91906143b9565b611215565b005b6103dc60048036038101906103d791906143b9565b611236565b005b6103f860048036038101906103f391906143b9565b6112b9565b6040516104059190614b70565b60405180910390f35b610416611362565b005b610420611397565b60405161042d9190614c0c565b60405180910390f35b61043e6113bd565b60405161044b9190614c27565b60405180910390f35b61046e60048036038101906104699190614200565b6113e3565b60405161047b9190614b92565b60405180910390f35b61048c611439565b6040516104999190614b92565b60405180910390f35b6104aa611450565b005b6104c660048036038101906104c191906144ac565b6114d8565b005b6104d06118b2565b6040516104dd9190614bad565b60405180910390f35b6104ee6118d6565b005b61050a60048036038101906105059190614200565b61190b565b005b610514611942565b60405161052191906148da565b60405180910390f35b610544600480360381019061053f919061450f565b61196c565b60405161055191906148da565b60405180910390f35b610574600480360381019061056f91906143b9565b61199b565b6040516105819190614b92565b60405180910390f35b6105a4600480360381019061059f91906142b5565b611a05565b005b6105ae6120d3565b6040516105bb9190614d24565b60405180910390f35b6105cc6120d9565b6040516105d99190614bad565b60405180910390f35b6105ea6120e0565b6040516105f79190614d24565b60405180910390f35b61061a6004803603810190610615919061459e565b6120e6565b005b610624612154565b6040516106319190614d24565b60405180910390f35b610654600480360381019061064f9190614390565b61215a565b6040516106619190614d24565b60405180910390f35b61067261217e565b60405161067f91906148da565b60405180910390f35b6106a2600480360381019061069d91906143b9565b6121a4565b005b6106be60048036038101906106b991906143f5565b6121c5565b005b6106da60048036038101906106d59190614200565b61244d565b6040516106e79190614b92565b60405180910390f35b61070a60048036038101906107059190614200565b61246d565b005b61071461257b565b6040516107219190614bad565b60405180910390f35b610744600480360381019061073f9190614200565b61259f565b005b610760600480360381019061075b9190614252565b6126fe565b005b61077c6004803603810190610777919061459e565b612b7a565b005b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166386624448826040518263ffffffff1660e01b81526004016107d99190614bad565b60206040518083038186803b1580156107f157600080fd5b505afa158015610805573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906108299190614367565b61085f576040517fafdd489000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b50565b60007f5a05180f000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916827bffffffffffffffffffffffffffffffffffffffffffffffffffffffff191614806108d557506108d4826131e5565b5b9050919050565b6108e58161077e565b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16634a43c5b983836040518363ffffffff1660e01b8152600401610942929190614bc8565b60206040518083038186803b15801561095a57600080fd5b505afa15801561096e573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906109929190614367565b6109c8576040517f03b5e41300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5050565b6109d4611439565b15610a14576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610a0b90614cc4565b60405180910390fd5b610a1e83856108dc565b6000600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16636ed3be53866040518263ffffffff1660e01b8152600401610a7b9190614bad565b6101006040518083038186803b158015610a9457600080fd5b505afa158015610aa8573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610acc9190614574565b9050610adb816060015161325f565b610ae881608001516132a6565b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663cd23b17285876040518363ffffffff1660e01b8152600401610b45929190614bc8565b60206040518083038186803b158015610b5d57600080fd5b505afa158015610b71573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610b959190614367565b15610bcc576040517fa0e9298400000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6000610bd66132ed565b90506000600b600088815260200190815260200160002060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205490506000845190506000821480610c425750600081145b15610c79576040517f45c109aa00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6000600a60008a815260200190815260200160002060008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020805480602002602001604051908101604052809291908181526020018280548015610d1557602002820191906000526020600020905b815481526020019060010190808311610d01575b5050505050905080518214610d56576040517f9ea6d12700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6000805b83811015610eb157828181518110610d9b577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b6020026020010151888281518110610ddc577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60200260200101518a604051602001610df69291906148ae565b60405160208183030381529060405280519060200120148015610e585750878181518110610e4d577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b602002602001015182105b15610ea157878181518110610e96577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b602002602001015191505b610eaa8161339e565b9050610d5a565b508084118015610ec457508560c0015181115b15610f5b57600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663e7c847aa86838d6040518463ffffffff1660e01b8152600401610f2893929190614b39565b600060405180830381600087803b158015610f4257600080fd5b505af1158015610f56573d6000803e3d6000fd5b505050505b7f61d4ab60470c97b371c5f91be47873283386bb3af4de728e8795ef5318e013d6858b604051610f8c9291906149d3565b60405180910390a150505050505050505050565b7fa49807205ce4d355092ef5a8a18f56e8913cf4a201fbe287825b095693c21775610fca816133ab565b81600960006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505050565b6110176132ed565b73ffffffffffffffffffffffffffffffffffffffff16611035611942565b73ffffffffffffffffffffffffffffffffffffffff161461108b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161108290614ce4565b60405180910390fd5b6001600460008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055507f42e160154868087d6bfdc0ca23d96a1c1cfa32f1b72ba9ba27b69b98a0d819dc8160405161111291906148da565b60405180910390a150565b7f9f959e00d95122f5cbd677010436cf273ef535b86b056afc172852144b9491d781565b7fa49807205ce4d355092ef5a8a18f56e8913cf4a201fbe287825b095693c2177561116b816133ab565b620f424082106111a7576040517f7e0d5ce600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b816006819055507ffa406a120a9e7f2b332bfb7a43d3bf1c3f079262202907a6b69c94b2821a02c66006546040516111df9190614d24565b60405180910390a15050565b6000806000838152602001908152602001600020600101549050919050565b6380ac58cd60e01b81565b61121e826111eb565b611227816133ab565b61123183836133bf565b505050565b61123e6132ed565b73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16146112ab576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016112a290614d04565b60405180910390fd5b6112b582826133f3565b5050565b6060600a600084815260200190815260200160002060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002080548060200260200160405190810160405280929190818152602001828054801561135557602002820191906000526020600020905b815481526020019060010190808311611341575b5050505050905092915050565b7f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a61138c816133ab565b611394613427565b50565b600560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6000600460008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff169050919050565b6000600260149054906101000a900460ff16905090565b6114586132ed565b73ffffffffffffffffffffffffffffffffffffffff16611476611942565b73ffffffffffffffffffffffffffffffffffffffff16146114cc576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016114c390614ce4565b60405180910390fd5b6114d660006134c9565b565b6114e0611439565b15611520576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161151790614cc4565b60405180910390fd5b61152a84846108dc565b6000600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16636ed3be53856040518263ffffffff1660e01b81526004016115879190614bad565b6101006040518083038186803b1580156115a057600080fd5b505afa1580156115b4573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906115d89190614574565b90506115e7816040015161325f565b6115f481606001516132a6565b60006115fe6132ed565b90508160a00151831161163d576040517f7532083f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600b600086815260200190815260200160002060008273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020548311156117f857600560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166323b872dd8230600b60008a815260200190815260200160002060008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054876117319190614ed1565b6040518463ffffffff1660e01b815260040161174f9392919061499c565b602060405180830381600087803b15801561176957600080fd5b505af115801561177d573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906117a19190614367565b5082600b600087815260200190815260200160002060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055505b600a600086815260200190815260200160002060008273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000208490806001815401808255809150506001900390600052602060002001600090919091909150557f42412bd3e1b349d0fc9c3518c85d09f611398d80842c5d3e36449b76dca01b6b8186866040516118a2939291906149fc565b60405180910390a1505050505050565b7fa49807205ce4d355092ef5a8a18f56e8913cf4a201fbe287825b095693c2177581565b7f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a611900816133ab565b61190861358f565b50565b7fa49807205ce4d355092ef5a8a18f56e8913cf4a201fbe287825b095693c21775611935816133ab565b61193e82613632565b5050565b6000600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b600061199382600160008681526020019081526020016000206136ad90919063ffffffff16565b905092915050565b600080600084815260200190815260200160002060000160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16905092915050565b87611a258173ffffffffffffffffffffffffffffffffffffffff16612baf565b611a5b576040517f6eefed2000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8073ffffffffffffffffffffffffffffffffffffffff166301ffc9a76380ac58cd60e01b6040518263ffffffff1660e01b8152600401611a9b9190614bf1565b60206040518083038186803b158015611ab357600080fd5b505afa158015611ac7573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611aeb9190614367565b611af457600080fd5b611afc611439565b15611b3c576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611b3390614cc4565b60405180910390fd5b600042851080611b5f575083611b5d600854876136c790919063ffffffff16565b115b15611b96576040517f6f7eac2600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6000861415611bd0576040517ebfc92100000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6000611bda6132ed565b905060008b90508073ffffffffffffffffffffffffffffffffffffffff16636352211e8a6040518263ffffffff1660e01b8152600401611c1a9190614d24565b60206040518083038186803b158015611c3257600080fd5b505afa158015611c46573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611c6a9190614229565b92508273ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16141580611ded57503073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1663081812fc8b6040518263ffffffff1660e01b8152600401611cf29190614d24565b60206040518083038186803b158015611d0a57600080fd5b505afa158015611d1e573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611d429190614229565b73ffffffffffffffffffffffffffffffffffffffff161480611deb57508073ffffffffffffffffffffffffffffffffffffffff1663e985e9c584306040518363ffffffff1660e01b8152600401611d9a9291906148f5565b60206040518083038186803b158015611db257600080fd5b505afa158015611dc6573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611dea9190614367565b5b155b15611e24576040517f82b4290000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16630a7f89778b6040518263ffffffff1660e01b8152600401611e7f9190614bad565b60206040518083038186803b158015611e9757600080fd5b505afa158015611eab573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611ecf9190614367565b611f05576040517fa3b8915f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5060006007541115611fe857600560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166323b872dd82600960009054906101000a900473ffffffffffffffffffffffffffffffffffffffff166007546040518463ffffffff1660e01b8152600401611f949392919061499c565b602060405180830381600087803b158015611fae57600080fd5b505af1158015611fc2573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611fe69190614367565b505b50600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663deaf324a828a8c8989898d6040518863ffffffff1660e01b81526004016120509796959493929190614a33565b600060405180830381600087803b15801561206a57600080fd5b505af115801561207e573d6000803e3d6000fd5b505050507fa1605c4cefd3351cfe80d4cdc106707dea9aff20e3095109748a5f710ecf52d4818b8b8a8989898d6040516120bf98979695949392919061491e565b60405180910390a150505050505050505050565b60065481565b6000801b81565b60075481565b7fa49807205ce4d355092ef5a8a18f56e8913cf4a201fbe287825b095693c21775612110816133ab565b816007819055507fe7fa8737293f41b5dfa0d5c3e552860a06275bed7015581b083c7be7003308ba6007546040516121489190614d24565b60405180910390a15050565b60085481565b6000612177600160008481526020019081526020016000206136dd565b9050919050565b600960009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6121ad826111eb565b6121b6816133ab565b6121c083836133f3565b505050565b6121cd611439565b1561220d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161220490614cc4565b60405180910390fd5b61221782826108dc565b6000600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16636ed3be53836040518263ffffffff1660e01b81526004016122749190614bad565b6101006040518083038186803b15801561228d57600080fd5b505afa1580156122a1573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906122c59190614574565b90506122d481606001516132a6565b60006122de6132ed565b90508073ffffffffffffffffffffffffffffffffffffffff16826020015173ffffffffffffffffffffffffffffffffffffffff16148061234b575061234a7f9f959e00d95122f5cbd677010436cf273ef535b86b056afc172852144b9491d76123456132ed565b61199b565b5b612381576040517f82b4290000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166315c97238856040518263ffffffff1660e01b81526004016123dc9190614bad565b600060405180830381600087803b1580156123f657600080fd5b505af115801561240a573d6000803e3d6000fd5b505050507fa021239bb373fef19aa7c7ef798961ab038ebf95bd85d546b731d25855db190e818460405161243f9291906149d3565b60405180910390a150505050565b60046020528060005260406000206000915054906101000a900460ff1681565b6124756132ed565b73ffffffffffffffffffffffffffffffffffffffff16612493611942565b73ffffffffffffffffffffffffffffffffffffffff16146124e9576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016124e090614ce4565b60405180910390fd5b6000600460008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055507fd7e9ec6e6ecd65492dce6bf513cd6867560d49544421d0783ddf06e76c24470c8160405161257091906148da565b60405180910390a150565b7f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a81565b6125a76132ed565b73ffffffffffffffffffffffffffffffffffffffff166125c5611942565b73ffffffffffffffffffffffffffffffffffffffff161461261b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161261290614ce4565b60405180910390fd5b6126457f9f959e00d95122f5cbd677010436cf273ef535b86b056afc172852144b9491d7826136f2565b61266f7f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a826136f2565b61267c6000801b826136f2565b6126ad7f9f959e00d95122f5cbd677010436cf273ef535b86b056afc172852144b9491d76126a8611942565b6121a4565b6126de7f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a6126d9611942565b6121a4565b6126f26000801b6126ed611942565b6121a4565b6126fb81612bd2565b50565b612706611439565b15612746576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161273d90614cc4565b60405180910390fd5b61274f8261077e565b6000600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16636ed3be53846040518263ffffffff1660e01b81526004016127ac9190614bad565b6101006040518083038186803b1580156127c557600080fd5b505afa1580156127d9573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906127fd9190614574565b9050600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663cd23b17285856040518363ffffffff1660e01b815260040161285c929190614bc8565b60206040518083038186803b15801561287457600080fd5b505afa158015612888573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906128ac9190614367565b61293e57600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166315c97238856040518263ffffffff1660e01b815260040161290b9190614bad565b600060405180830381600087803b15801561292557600080fd5b505af1158015612939573d6000803e3d6000fd5b505050505b60006129486132ed565b90506000600b600086815260200190815260200160002060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054905060008114156129da576040517f35411d8b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8260e0015173ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff161415612a30578260c0015181612a229190614ed1565b9050612a2f878686613700565b5b600560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663a9059cbb83836040518363ffffffff1660e01b8152600401612a8d929190614b10565b602060405180830381600087803b158015612aa757600080fd5b505af1158015612abb573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612adf9190614367565b506000600b600087815260200190815260200160002060008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055507f5f4b88832be3746851e4d0ce6129f89fe39ddcf5a4b0204699724c459baf7d9d828683604051612b6993929190614aa2565b60405180910390a150505050505050565b7fa49807205ce4d355092ef5a8a18f56e8913cf4a201fbe287825b095693c21775612ba4816133ab565b816008819055505050565b6000808273ffffffffffffffffffffffffffffffffffffffff163b119050919050565b612bda6132ed565b73ffffffffffffffffffffffffffffffffffffffff16612bf8611942565b73ffffffffffffffffffffffffffffffffffffffff1614612c4e576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612c4590614ce4565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff161415612cbe576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612cb590614ca4565b60405180910390fd5b612cc7816134c9565b50565b612cd4828261199b565b612da657600160008084815260200190815260200160002060000160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff021916908315150217905550612d4b6132ed565b73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16837f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a45b5050565b6000612dd2836000018373ffffffffffffffffffffffffffffffffffffffff1660001b613c38565b905092915050565b612de4828261199b565b15612eb757600080600084815260200190815260200160002060000160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff021916908315150217905550612e5c6132ed565b73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16837ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b60405160405180910390a45b5050565b6000612ee3836000018373ffffffffffffffffffffffffffffffffffffffff1660001b613ca8565b905092915050565b606060006002836002612efe9190614e77565b612f089190614df0565b67ffffffffffffffff811115612f47577f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6040519080825280601f01601f191660200182016040528015612f795781602001600182028036833780820191505090505b5090507f300000000000000000000000000000000000000000000000000000000000000081600081518110612fd7577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a9053507f780000000000000000000000000000000000000000000000000000000000000081600181518110613061577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a905350600060018460026130a19190614e77565b6130ab9190614df0565b90505b6001811115613197577f3031323334353637383961626364656600000000000000000000000000000000600f861660108110613113577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b1a60f81b828281518110613150577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a905350600485901c94508061319090615010565b90506130ae565b50600084146131db576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016131d290614c64565b60405180910390fd5b8091505092915050565b60007f7965db0b000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916827bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19161480613258575061325782613e2e565b5b9050919050565b8042116132a357806040517f2a35a32400000000000000000000000000000000000000000000000000000000815260040161329a9190614d24565b60405180910390fd5b50565b8042106132ea57806040517f691e56820000000000000000000000000000000000000000000000000000000081526004016132e19190614d24565b60405180910390fd5b50565b60003073ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16141561339757600080368080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f820116905080830192505050505050509050600080369050905073ffffffffffffffffffffffffffffffffffffffff81830151169250505061339b565b3390505b90565b6000600182019050919050565b6133bc816133b76132ed565b613e98565b50565b6133c98282612cca565b6133ee8160016000858152602001908152602001600020612daa90919063ffffffff16565b505050565b6133fd8282612dda565b6134228160016000858152602001908152602001600020612ebb90919063ffffffff16565b505050565b61342f611439565b61346e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161346590614c84565b60405180910390fd5b6000600260146101000a81548160ff0219169083151502179055507f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa6134b26132ed565b6040516134bf91906148da565b60405180910390a1565b6000600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905081600360006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b613597611439565b156135d7576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016135ce90614cc4565b60405180910390fd5b6001600260146101000a81548160ff0219169083151502179055507f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a25861361b6132ed565b60405161362891906148da565b60405180910390a1565b80600260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055507f19e07a962d0e1cab8b7d4be06b77a7225ca7ebb744994cb6546fdd815590dc9d816040516136a291906148da565b60405180910390a150565b60006136bc8360000183613f35565b60001c905092915050565b600081836136d59190614df0565b905092915050565b60006136eb82600001613f86565b9050919050565b6136fc82826133bf565b5050565b6000600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16636ed3be53846040518263ffffffff1660e01b815260040161375d9190614bad565b6101006040518083038186803b15801561377657600080fd5b505afa15801561378a573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906137ae9190614574565b905060008160200151905060008590508073ffffffffffffffffffffffffffffffffffffffff16636352211e856040518263ffffffff1660e01b81526004016137f79190614d24565b60206040518083038186803b15801561380f57600080fd5b505afa158015613823573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906138479190614229565b73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16146138ab576040517f82b4290000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60008360e001519050600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16141561391b576040517f946ef90200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6000808560c00151905060006006541115613a2f5761395a620f424061394c60065484613f9790919063ffffffff16565b613fad90919063ffffffff16565b9150600560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663a9059cbb600960009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16846040518363ffffffff1660e01b81526004016139db929190614b10565b602060405180830381600087803b1580156139f557600080fd5b505af1158015613a09573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190613a2d9190614367565b505b600560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663a9059cbb86613a818585613fc390919063ffffffff16565b6040518363ffffffff1660e01b8152600401613a9e929190614b10565b602060405180830381600087803b158015613ab857600080fd5b505af1158015613acc573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190613af09190614367565b508373ffffffffffffffffffffffffffffffffffffffff166342842e0e86858a6040518463ffffffff1660e01b8152600401613b2e9392919061499c565b600060405180830381600087803b158015613b4857600080fd5b505af1158015613b5c573d6000803e3d6000fd5b50505050600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663e7c847aa6000808b6040518463ffffffff1660e01b8152600401613bc093929190614ad9565b600060405180830381600087803b158015613bda57600080fd5b505af1158015613bee573d6000803e3d6000fd5b505050507fa307d88e125ed19583b02332c5bd74f0323774248b3f9cac5047c2433d9d8ddc838989604051613c2593929190614aa2565b60405180910390a1505050505050505050565b6000613c448383613fd9565b613c9d578260000182908060018154018082558091505060019003906000526020600020016000909190919091505582600001805490508360010160008481526020019081526020016000208190555060019050613ca2565b600090505b92915050565b60008083600101600084815260200190815260200160002054905060008114613e22576000600182613cda9190614ed1565b9050600060018660000180549050613cf29190614ed1565b9050818114613dad576000866000018281548110613d39577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b9060005260206000200154905080876000018481548110613d83577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b90600052602060002001819055508387600101600083815260200190815260200160002081905550505b85600001805480613de7577f4e487b7100000000000000000000000000000000000000000000000000000000600052603160045260246000fd5b600190038181906000526020600020016000905590558560010160008681526020019081526020016000206000905560019350505050613e28565b60009150505b92915050565b60007f01ffc9a7000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916827bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916149050919050565b613ea2828261199b565b613f3157613ec78173ffffffffffffffffffffffffffffffffffffffff166014612eeb565b613ed58360001c6020612eeb565b604051602001613ee6929190614874565b6040516020818303038152906040526040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401613f289190614c42565b60405180910390fd5b5050565b6000826000018281548110613f73577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b9060005260206000200154905092915050565b600081600001805490509050919050565b60008183613fa59190614e77565b905092915050565b60008183613fbb9190614e46565b905092915050565b60008183613fd19190614ed1565b905092915050565b600080836001016000848152602001908152602001600020541415905092915050565b600061400f61400a84614d64565b614d3f565b9050808382526020820190508285602086028201111561402e57600080fd5b60005b8581101561405e578161404488826141d6565b845260208401935060208301925050600181019050614031565b5050509392505050565b600081359050614077816152b1565b92915050565b60008151905061408c816152b1565b92915050565b600082601f8301126140a357600080fd5b81356140b3848260208601613ffc565b91505092915050565b6000815190506140cb816152c8565b92915050565b6000813590506140e0816152df565b92915050565b6000815190506140f5816152df565b92915050565b60008135905061410a816152f6565b92915050565b6000610100828403121561412357600080fd5b61412e610100614d3f565b9050600061413e848285016140e6565b60008301525060206141528482850161407d565b6020830152506040614166848285016141eb565b604083015250606061417a848285016141eb565b606083015250608061418e848285016141eb565b60808301525060a06141a2848285016141eb565b60a08301525060c06141b6848285016141eb565b60c08301525060e06141ca8482850161407d565b60e08301525092915050565b6000813590506141e58161530d565b92915050565b6000815190506141fa8161530d565b92915050565b60006020828403121561421257600080fd5b600061422084828501614068565b91505092915050565b60006020828403121561423b57600080fd5b60006142498482850161407d565b91505092915050565b6000806000806080858703121561426857600080fd5b600061427687828801614068565b9450506020614287878288016140d1565b9350506040614298878288016140d1565b92505060606142a9878288016141d6565b91505092959194509250565b600080600080600080600080610100898b0312156142d257600080fd5b60006142e08b828c01614068565b98505060206142f18b828c016140d1565b97505060406143028b828c016140d1565b96505060606143138b828c016141d6565b95505060806143248b828c016141d6565b94505060a06143358b828c016141d6565b93505060c06143468b828c016141d6565b92505060e06143578b828c016141d6565b9150509295985092959890939650565b60006020828403121561437957600080fd5b6000614387848285016140bc565b91505092915050565b6000602082840312156143a257600080fd5b60006143b0848285016140d1565b91505092915050565b600080604083850312156143cc57600080fd5b60006143da858286016140d1565b92505060206143eb85828601614068565b9150509250929050565b6000806040838503121561440857600080fd5b6000614416858286016140d1565b9250506020614427858286016140d1565b9150509250929050565b6000806000806080858703121561444757600080fd5b6000614455878288016140d1565b9450506020614466878288016140d1565b9350506040614477878288016140d1565b925050606085013567ffffffffffffffff81111561449457600080fd5b6144a087828801614092565b91505092959194509250565b600080600080608085870312156144c257600080fd5b60006144d0878288016140d1565b94505060206144e1878288016140d1565b93505060406144f2878288016140d1565b9250506060614503878288016141d6565b91505092959194509250565b6000806040838503121561452257600080fd5b6000614530858286016140d1565b9250506020614541858286016141d6565b9150509250929050565b60006020828403121561455d57600080fd5b600061456b848285016140fb565b91505092915050565b6000610100828403121561458757600080fd5b600061459584828501614110565b91505092915050565b6000602082840312156145b057600080fd5b60006145be848285016141d6565b91505092915050565b60006145d3838361465b565b60208301905092915050565b6145e881614f05565b82525050565b60006145f982614da0565b6146038185614dc3565b935061460e83614d90565b8060005b8381101561463f57815161462688826145c7565b975061463183614db6565b925050600181019050614612565b5085935050505092915050565b61465581614f17565b82525050565b61466481614f23565b82525050565b61467381614f23565b82525050565b61468a61468582614f23565b61506b565b82525050565b61469981614f2d565b82525050565b6146a881614f83565b82525050565b6146b781614fa7565b82525050565b6146c681614fcb565b82525050565b60006146d782614dab565b6146e18185614dd4565b93506146f1818560208601614fdd565b6146fa8161510c565b840191505092915050565b600061471082614dab565b61471a8185614de5565b935061472a818560208601614fdd565b80840191505092915050565b6000614743602083614dd4565b915061474e8261511d565b602082019050919050565b6000614766601483614dd4565b915061477182615146565b602082019050919050565b6000614789602683614dd4565b91506147948261516f565b604082019050919050565b60006147ac601083614dd4565b91506147b7826151be565b602082019050919050565b60006147cf602083614dd4565b91506147da826151e7565b602082019050919050565b60006147f2601783614de5565b91506147fd82615210565b601782019050919050565b6000614815601183614de5565b915061482082615239565b601182019050919050565b6000614838602f83614dd4565b915061484382615262565b604082019050919050565b61485781614f79565b82525050565b61486e61486982614f79565b615075565b82525050565b600061487f826147e5565b915061488b8285614705565b915061489682614808565b91506148a28284614705565b91508190509392505050565b60006148ba828561485d565b6020820191506148ca8284614679565b6020820191508190509392505050565b60006020820190506148ef60008301846145df565b92915050565b600060408201905061490a60008301856145df565b61491760208301846145df565b9392505050565b600061010082019050614934600083018b6145df565b614941602083018a6145df565b61494e604083018961466a565b61495b606083018861484e565b614968608083018761484e565b61497560a083018661484e565b61498260c083018561484e565b61498f60e083018461484e565b9998505050505050505050565b60006060820190506149b160008301866145df565b6149be60208301856145df565b6149cb604083018461484e565b949350505050565b60006040820190506149e860008301856145df565b6149f5602083018461466a565b9392505050565b6000606082019050614a1160008301866145df565b614a1e602083018561466a565b614a2b604083018461466a565b949350505050565b600060e082019050614a48600083018a6145df565b614a55602083018961466a565b614a62604083018861466a565b614a6f606083018761484e565b614a7c608083018661484e565b614a8960a083018561484e565b614a9660c083018461484e565b98975050505050505050565b6000606082019050614ab760008301866145df565b614ac4602083018561466a565b614ad1604083018461484e565b949350505050565b6000606082019050614aee60008301866145df565b614afb60208301856146bd565b614b08604083018461466a565b949350505050565b6000604082019050614b2560008301856145df565b614b32602083018461484e565b9392505050565b6000606082019050614b4e60008301866145df565b614b5b602083018561484e565b614b68604083018461466a565b949350505050565b60006020820190508181036000830152614b8a81846145ee565b905092915050565b6000602082019050614ba7600083018461464c565b92915050565b6000602082019050614bc2600083018461466a565b92915050565b6000604082019050614bdd600083018561466a565b614bea602083018461466a565b9392505050565b6000602082019050614c066000830184614690565b92915050565b6000602082019050614c21600083018461469f565b92915050565b6000602082019050614c3c60008301846146ae565b92915050565b60006020820190508181036000830152614c5c81846146cc565b905092915050565b60006020820190508181036000830152614c7d81614736565b9050919050565b60006020820190508181036000830152614c9d81614759565b9050919050565b60006020820190508181036000830152614cbd8161477c565b9050919050565b60006020820190508181036000830152614cdd8161479f565b9050919050565b60006020820190508181036000830152614cfd816147c2565b9050919050565b60006020820190508181036000830152614d1d8161482b565b9050919050565b6000602082019050614d39600083018461484e565b92915050565b6000614d49614d5a565b9050614d55828261503a565b919050565b6000604051905090565b600067ffffffffffffffff821115614d7f57614d7e6150dd565b5b602082029050602081019050919050565b6000819050602082019050919050565b600081519050919050565b600081519050919050565b6000602082019050919050565b600082825260208201905092915050565b600082825260208201905092915050565b600081905092915050565b6000614dfb82614f79565b9150614e0683614f79565b9250827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff03821115614e3b57614e3a61507f565b5b828201905092915050565b6000614e5182614f79565b9150614e5c83614f79565b925082614e6c57614e6b6150ae565b5b828204905092915050565b6000614e8282614f79565b9150614e8d83614f79565b9250817fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0483118215151615614ec657614ec561507f565b5b828202905092915050565b6000614edc82614f79565b9150614ee783614f79565b925082821015614efa57614ef961507f565b5b828203905092915050565b6000614f1082614f59565b9050919050565b60008115159050919050565b6000819050919050565b60007fffffffff0000000000000000000000000000000000000000000000000000000082169050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b6000614f8e82614f95565b9050919050565b6000614fa082614f59565b9050919050565b6000614fb282614fb9565b9050919050565b6000614fc482614f59565b9050919050565b6000614fd682614f79565b9050919050565b60005b83811015614ffb578082015181840152602081019050614fe0565b8381111561500a576000848401525b50505050565b600061501b82614f79565b9150600082141561502f5761502e61507f565b5b600182039050919050565b6150438261510c565b810181811067ffffffffffffffff82111715615062576150616150dd565b5b80604052505050565b6000819050919050565b6000819050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6000601f19601f8301169050919050565b7f537472696e67733a20686578206c656e67746820696e73756666696369656e74600082015250565b7f5061757361626c653a206e6f7420706175736564000000000000000000000000600082015250565b7f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160008201527f6464726573730000000000000000000000000000000000000000000000000000602082015250565b7f5061757361626c653a2070617573656400000000000000000000000000000000600082015250565b7f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572600082015250565b7f416363657373436f6e74726f6c3a206163636f756e7420000000000000000000600082015250565b7f206973206d697373696e6720726f6c6520000000000000000000000000000000600082015250565b7f416363657373436f6e74726f6c3a2063616e206f6e6c792072656e6f756e636560008201527f20726f6c657320666f722073656c660000000000000000000000000000000000602082015250565b6152ba81614f05565b81146152c557600080fd5b50565b6152d181614f17565b81146152dc57600080fd5b50565b6152e881614f23565b81146152f357600080fd5b50565b6152ff81614f2d565b811461530a57600080fd5b50565b61531681614f79565b811461532157600080fd5b5056fea26469706673582212207aabd14785a4e23e1603aab929a4014ee50cedecea7b1865bd3af34ca771a6e464736f6c63430008040033",
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

// CheckExisted is a free data retrieval call binding the contract method 0x00c5c3c3.
//
// Solidity: function checkExisted(bytes32 blindAuctionId) view returns()
func (_BlindAuctionMarketplace *BlindAuctionMarketplaceCaller) CheckExisted(opts *bind.CallOpts, blindAuctionId [32]byte) error {
	var out []interface{}
	err := _BlindAuctionMarketplace.contract.Call(opts, &out, "checkExisted", blindAuctionId)

	if err != nil {
		return err
	}

	return err

}

// CheckExisted is a free data retrieval call binding the contract method 0x00c5c3c3.
//
// Solidity: function checkExisted(bytes32 blindAuctionId) view returns()
func (_BlindAuctionMarketplace *BlindAuctionMarketplaceSession) CheckExisted(blindAuctionId [32]byte) error {
	return _BlindAuctionMarketplace.Contract.CheckExisted(&_BlindAuctionMarketplace.CallOpts, blindAuctionId)
}

// CheckExisted is a free data retrieval call binding the contract method 0x00c5c3c3.
//
// Solidity: function checkExisted(bytes32 blindAuctionId) view returns()
func (_BlindAuctionMarketplace *BlindAuctionMarketplaceCallerSession) CheckExisted(blindAuctionId [32]byte) error {
	return _BlindAuctionMarketplace.Contract.CheckExisted(&_BlindAuctionMarketplace.CallOpts, blindAuctionId)
}

// CheckRunning is a free data retrieval call binding the contract method 0x076f6dc7.
//
// Solidity: function checkRunning(bytes32 nftAsset, bytes32 blindAuctionId) view returns()
func (_BlindAuctionMarketplace *BlindAuctionMarketplaceCaller) CheckRunning(opts *bind.CallOpts, nftAsset [32]byte, blindAuctionId [32]byte) error {
	var out []interface{}
	err := _BlindAuctionMarketplace.contract.Call(opts, &out, "checkRunning", nftAsset, blindAuctionId)

	if err != nil {
		return err
	}

	return err

}

// CheckRunning is a free data retrieval call binding the contract method 0x076f6dc7.
//
// Solidity: function checkRunning(bytes32 nftAsset, bytes32 blindAuctionId) view returns()
func (_BlindAuctionMarketplace *BlindAuctionMarketplaceSession) CheckRunning(nftAsset [32]byte, blindAuctionId [32]byte) error {
	return _BlindAuctionMarketplace.Contract.CheckRunning(&_BlindAuctionMarketplace.CallOpts, nftAsset, blindAuctionId)
}

// CheckRunning is a free data retrieval call binding the contract method 0x076f6dc7.
//
// Solidity: function checkRunning(bytes32 nftAsset, bytes32 blindAuctionId) view returns()
func (_BlindAuctionMarketplace *BlindAuctionMarketplaceCallerSession) CheckRunning(nftAsset [32]byte, blindAuctionId [32]byte) error {
	return _BlindAuctionMarketplace.Contract.CheckRunning(&_BlindAuctionMarketplace.CallOpts, nftAsset, blindAuctionId)
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
// Solidity: function getBlindBid(bytes32 blindAuctionId, address bidder) view returns(bytes32[])
func (_BlindAuctionMarketplace *BlindAuctionMarketplaceCaller) GetBlindBid(opts *bind.CallOpts, blindAuctionId [32]byte, bidder common.Address) ([][32]byte, error) {
	var out []interface{}
	err := _BlindAuctionMarketplace.contract.Call(opts, &out, "getBlindBid", blindAuctionId, bidder)

	if err != nil {
		return *new([][32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][32]byte)).(*[][32]byte)

	return out0, err

}

// GetBlindBid is a free data retrieval call binding the contract method 0x3af4b422.
//
// Solidity: function getBlindBid(bytes32 blindAuctionId, address bidder) view returns(bytes32[])
func (_BlindAuctionMarketplace *BlindAuctionMarketplaceSession) GetBlindBid(blindAuctionId [32]byte, bidder common.Address) ([][32]byte, error) {
	return _BlindAuctionMarketplace.Contract.GetBlindBid(&_BlindAuctionMarketplace.CallOpts, blindAuctionId, bidder)
}

// GetBlindBid is a free data retrieval call binding the contract method 0x3af4b422.
//
// Solidity: function getBlindBid(bytes32 blindAuctionId, address bidder) view returns(bytes32[])
func (_BlindAuctionMarketplace *BlindAuctionMarketplaceCallerSession) GetBlindBid(blindAuctionId [32]byte, bidder common.Address) ([][32]byte, error) {
	return _BlindAuctionMarketplace.Contract.GetBlindBid(&_BlindAuctionMarketplace.CallOpts, blindAuctionId, bidder)
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

// Bid is a paid mutator transaction binding the contract method 0x7544b61b.
//
// Solidity: function bid(bytes32 nftAsset, bytes32 blindAuctionId, bytes32 blindedBid, uint256 deposit) returns()
func (_BlindAuctionMarketplace *BlindAuctionMarketplaceTransactor) Bid(opts *bind.TransactOpts, nftAsset [32]byte, blindAuctionId [32]byte, blindedBid [32]byte, deposit *big.Int) (*types.Transaction, error) {
	return _BlindAuctionMarketplace.contract.Transact(opts, "bid", nftAsset, blindAuctionId, blindedBid, deposit)
}

// Bid is a paid mutator transaction binding the contract method 0x7544b61b.
//
// Solidity: function bid(bytes32 nftAsset, bytes32 blindAuctionId, bytes32 blindedBid, uint256 deposit) returns()
func (_BlindAuctionMarketplace *BlindAuctionMarketplaceSession) Bid(nftAsset [32]byte, blindAuctionId [32]byte, blindedBid [32]byte, deposit *big.Int) (*types.Transaction, error) {
	return _BlindAuctionMarketplace.Contract.Bid(&_BlindAuctionMarketplace.TransactOpts, nftAsset, blindAuctionId, blindedBid, deposit)
}

// Bid is a paid mutator transaction binding the contract method 0x7544b61b.
//
// Solidity: function bid(bytes32 nftAsset, bytes32 blindAuctionId, bytes32 blindedBid, uint256 deposit) returns()
func (_BlindAuctionMarketplace *BlindAuctionMarketplaceTransactorSession) Bid(nftAsset [32]byte, blindAuctionId [32]byte, blindedBid [32]byte, deposit *big.Int) (*types.Transaction, error) {
	return _BlindAuctionMarketplace.Contract.Bid(&_BlindAuctionMarketplace.TransactOpts, nftAsset, blindAuctionId, blindedBid, deposit)
}

// CancelAuction is a paid mutator transaction binding the contract method 0xdaa3d985.
//
// Solidity: function cancelAuction(bytes32 nftAsset, bytes32 blindAuctionId) returns()
func (_BlindAuctionMarketplace *BlindAuctionMarketplaceTransactor) CancelAuction(opts *bind.TransactOpts, nftAsset [32]byte, blindAuctionId [32]byte) (*types.Transaction, error) {
	return _BlindAuctionMarketplace.contract.Transact(opts, "cancelAuction", nftAsset, blindAuctionId)
}

// CancelAuction is a paid mutator transaction binding the contract method 0xdaa3d985.
//
// Solidity: function cancelAuction(bytes32 nftAsset, bytes32 blindAuctionId) returns()
func (_BlindAuctionMarketplace *BlindAuctionMarketplaceSession) CancelAuction(nftAsset [32]byte, blindAuctionId [32]byte) (*types.Transaction, error) {
	return _BlindAuctionMarketplace.Contract.CancelAuction(&_BlindAuctionMarketplace.TransactOpts, nftAsset, blindAuctionId)
}

// CancelAuction is a paid mutator transaction binding the contract method 0xdaa3d985.
//
// Solidity: function cancelAuction(bytes32 nftAsset, bytes32 blindAuctionId) returns()
func (_BlindAuctionMarketplace *BlindAuctionMarketplaceTransactorSession) CancelAuction(nftAsset [32]byte, blindAuctionId [32]byte) (*types.Transaction, error) {
	return _BlindAuctionMarketplace.Contract.CancelAuction(&_BlindAuctionMarketplace.TransactOpts, nftAsset, blindAuctionId)
}

// CreateAuction is a paid mutator transaction binding the contract method 0x9655fe1a.
//
// Solidity: function createAuction(address nftAddress, bytes32 blindAuctionId, bytes32 nftAsset, uint256 assetId, uint256 startPriceInWei, uint256 startTime, uint256 biddingEnd, uint256 revealEnd) returns()
func (_BlindAuctionMarketplace *BlindAuctionMarketplaceTransactor) CreateAuction(opts *bind.TransactOpts, nftAddress common.Address, blindAuctionId [32]byte, nftAsset [32]byte, assetId *big.Int, startPriceInWei *big.Int, startTime *big.Int, biddingEnd *big.Int, revealEnd *big.Int) (*types.Transaction, error) {
	return _BlindAuctionMarketplace.contract.Transact(opts, "createAuction", nftAddress, blindAuctionId, nftAsset, assetId, startPriceInWei, startTime, biddingEnd, revealEnd)
}

// CreateAuction is a paid mutator transaction binding the contract method 0x9655fe1a.
//
// Solidity: function createAuction(address nftAddress, bytes32 blindAuctionId, bytes32 nftAsset, uint256 assetId, uint256 startPriceInWei, uint256 startTime, uint256 biddingEnd, uint256 revealEnd) returns()
func (_BlindAuctionMarketplace *BlindAuctionMarketplaceSession) CreateAuction(nftAddress common.Address, blindAuctionId [32]byte, nftAsset [32]byte, assetId *big.Int, startPriceInWei *big.Int, startTime *big.Int, biddingEnd *big.Int, revealEnd *big.Int) (*types.Transaction, error) {
	return _BlindAuctionMarketplace.Contract.CreateAuction(&_BlindAuctionMarketplace.TransactOpts, nftAddress, blindAuctionId, nftAsset, assetId, startPriceInWei, startTime, biddingEnd, revealEnd)
}

// CreateAuction is a paid mutator transaction binding the contract method 0x9655fe1a.
//
// Solidity: function createAuction(address nftAddress, bytes32 blindAuctionId, bytes32 nftAsset, uint256 assetId, uint256 startPriceInWei, uint256 startTime, uint256 biddingEnd, uint256 revealEnd) returns()
func (_BlindAuctionMarketplace *BlindAuctionMarketplaceTransactorSession) CreateAuction(nftAddress common.Address, blindAuctionId [32]byte, nftAsset [32]byte, assetId *big.Int, startPriceInWei *big.Int, startTime *big.Int, biddingEnd *big.Int, revealEnd *big.Int) (*types.Transaction, error) {
	return _BlindAuctionMarketplace.Contract.CreateAuction(&_BlindAuctionMarketplace.TransactOpts, nftAddress, blindAuctionId, nftAsset, assetId, startPriceInWei, startTime, biddingEnd, revealEnd)
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

// Reveal is a paid mutator transaction binding the contract method 0x07c8e9cb.
//
// Solidity: function reveal(bytes32 blindAuctionId, bytes32 nftAsset, bytes32 secret, uint256[] _values) returns()
func (_BlindAuctionMarketplace *BlindAuctionMarketplaceTransactor) Reveal(opts *bind.TransactOpts, blindAuctionId [32]byte, nftAsset [32]byte, secret [32]byte, _values []*big.Int) (*types.Transaction, error) {
	return _BlindAuctionMarketplace.contract.Transact(opts, "reveal", blindAuctionId, nftAsset, secret, _values)
}

// Reveal is a paid mutator transaction binding the contract method 0x07c8e9cb.
//
// Solidity: function reveal(bytes32 blindAuctionId, bytes32 nftAsset, bytes32 secret, uint256[] _values) returns()
func (_BlindAuctionMarketplace *BlindAuctionMarketplaceSession) Reveal(blindAuctionId [32]byte, nftAsset [32]byte, secret [32]byte, _values []*big.Int) (*types.Transaction, error) {
	return _BlindAuctionMarketplace.Contract.Reveal(&_BlindAuctionMarketplace.TransactOpts, blindAuctionId, nftAsset, secret, _values)
}

// Reveal is a paid mutator transaction binding the contract method 0x07c8e9cb.
//
// Solidity: function reveal(bytes32 blindAuctionId, bytes32 nftAsset, bytes32 secret, uint256[] _values) returns()
func (_BlindAuctionMarketplace *BlindAuctionMarketplaceTransactorSession) Reveal(blindAuctionId [32]byte, nftAsset [32]byte, secret [32]byte, _values []*big.Int) (*types.Transaction, error) {
	return _BlindAuctionMarketplace.Contract.Reveal(&_BlindAuctionMarketplace.TransactOpts, blindAuctionId, nftAsset, secret, _values)
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

// Withdraw is a paid mutator transaction binding the contract method 0xf3764ffa.
//
// Solidity: function withdraw(address nftAddress, bytes32 nftAsset, bytes32 blindAuctionId, uint256 assetId) returns()
func (_BlindAuctionMarketplace *BlindAuctionMarketplaceTransactor) Withdraw(opts *bind.TransactOpts, nftAddress common.Address, nftAsset [32]byte, blindAuctionId [32]byte, assetId *big.Int) (*types.Transaction, error) {
	return _BlindAuctionMarketplace.contract.Transact(opts, "withdraw", nftAddress, nftAsset, blindAuctionId, assetId)
}

// Withdraw is a paid mutator transaction binding the contract method 0xf3764ffa.
//
// Solidity: function withdraw(address nftAddress, bytes32 nftAsset, bytes32 blindAuctionId, uint256 assetId) returns()
func (_BlindAuctionMarketplace *BlindAuctionMarketplaceSession) Withdraw(nftAddress common.Address, nftAsset [32]byte, blindAuctionId [32]byte, assetId *big.Int) (*types.Transaction, error) {
	return _BlindAuctionMarketplace.Contract.Withdraw(&_BlindAuctionMarketplace.TransactOpts, nftAddress, nftAsset, blindAuctionId, assetId)
}

// Withdraw is a paid mutator transaction binding the contract method 0xf3764ffa.
//
// Solidity: function withdraw(address nftAddress, bytes32 nftAsset, bytes32 blindAuctionId, uint256 assetId) returns()
func (_BlindAuctionMarketplace *BlindAuctionMarketplaceTransactorSession) Withdraw(nftAddress common.Address, nftAsset [32]byte, blindAuctionId [32]byte, assetId *big.Int) (*types.Transaction, error) {
	return _BlindAuctionMarketplace.Contract.Withdraw(&_BlindAuctionMarketplace.TransactOpts, nftAddress, nftAsset, blindAuctionId, assetId)
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
	StartTime       *big.Int
	BiddingEnd      *big.Int
	RevealEnd       *big.Int
	StartPriceInWei *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterBlindAuctionCreatedSuccessful is a free log retrieval operation binding the contract event 0xa1605c4cefd3351cfe80d4cdc106707dea9aff20e3095109748a5f710ecf52d4.
//
// Solidity: event BlindAuctionCreatedSuccessful(address assetOwner, address nftAddress, bytes32 blindAuctionId, uint256 assetId, uint256 startTime, uint256 biddingEnd, uint256 revealEnd, uint256 startPriceInWei)
func (_BlindAuctionMarketplace *BlindAuctionMarketplaceFilterer) FilterBlindAuctionCreatedSuccessful(opts *bind.FilterOpts) (*BlindAuctionMarketplaceBlindAuctionCreatedSuccessfulIterator, error) {

	logs, sub, err := _BlindAuctionMarketplace.contract.FilterLogs(opts, "BlindAuctionCreatedSuccessful")
	if err != nil {
		return nil, err
	}
	return &BlindAuctionMarketplaceBlindAuctionCreatedSuccessfulIterator{contract: _BlindAuctionMarketplace.contract, event: "BlindAuctionCreatedSuccessful", logs: logs, sub: sub}, nil
}

// WatchBlindAuctionCreatedSuccessful is a free log subscription operation binding the contract event 0xa1605c4cefd3351cfe80d4cdc106707dea9aff20e3095109748a5f710ecf52d4.
//
// Solidity: event BlindAuctionCreatedSuccessful(address assetOwner, address nftAddress, bytes32 blindAuctionId, uint256 assetId, uint256 startTime, uint256 biddingEnd, uint256 revealEnd, uint256 startPriceInWei)
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

// ParseBlindAuctionCreatedSuccessful is a log parse operation binding the contract event 0xa1605c4cefd3351cfe80d4cdc106707dea9aff20e3095109748a5f710ecf52d4.
//
// Solidity: event BlindAuctionCreatedSuccessful(address assetOwner, address nftAddress, bytes32 blindAuctionId, uint256 assetId, uint256 startTime, uint256 biddingEnd, uint256 revealEnd, uint256 startPriceInWei)
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
