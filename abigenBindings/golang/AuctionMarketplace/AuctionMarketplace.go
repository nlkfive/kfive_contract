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
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"acceptedToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"marketplaceStorage\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"beneficary\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"ownerCutPerMillion\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"Bidded\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidBiddingEnd\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidContract\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidCut\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidPrice\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidRevealEnd\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotExisted\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotRunning\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"}],\"name\":\"TooEarly\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"}],\"name\":\"TooLate\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Unauthorized\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Unavailable\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"AddedBlackList\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"who\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"auctionId\",\"type\":\"bytes32\"}],\"name\":\"AuctionCancelled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"nftAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"auctionId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"assetId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"biddingEnd\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"revealEnd\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"startPriceInWei\",\"type\":\"uint256\"}],\"name\":\"AuctionCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"auctionId\",\"type\":\"bytes32\"}],\"name\":\"AuctionEnded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"bidder\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"auctionId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"deposit\",\"type\":\"uint256\"}],\"name\":\"AuctionRefund\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"buyer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"auctionId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalPrice\",\"type\":\"uint256\"}],\"name\":\"AuctionSuccessful\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"bidder\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"auctionId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"blindedBid\",\"type\":\"bytes32\"}],\"name\":\"BidSuccessful\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"ownerCutPerMillion\",\"type\":\"uint256\"}],\"name\":\"ChangedOwnerCutPerMillion\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"publicationFee\",\"type\":\"uint256\"}],\"name\":\"ChangedPublicationFee\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_marketplaceStorage\",\"type\":\"address\"}],\"name\":\"MarketplaceStorageUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"RemovedBlackList\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"fake\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"revealer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"auctionId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"RevealFailed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"fake\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"revealer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"auctionId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"blindedBid\",\"type\":\"bytes32\"}],\"name\":\"RevealSuccessful\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"CANCEL_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ERC721_Interface\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"IMarketplaceStorage_Interface\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PAUSER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"acceptedToken\",\"outputs\":[{\"internalType\":\"contractIBEP20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_evilUser\",\"type\":\"address\"}],\"name\":\"addBlackList\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"beneficary\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_maker\",\"type\":\"address\"}],\"name\":\"getBlackListStatus\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getRoleMember\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleMemberCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"isBlackListed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"marketplaceStorage\",\"outputs\":[{\"internalType\":\"contractIMarketplaceStorage\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minStageDuration\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ownerCutPerMillion\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"publicationFeeInWei\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_clearedUser\",\"type\":\"address\"}],\"name\":\"removeBlackList\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_beneficary\",\"type\":\"address\"}],\"name\":\"setBeneficary\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"}],\"name\":\"setMinStageDuration\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_ownerCutPerMillion\",\"type\":\"uint256\"}],\"name\":\"setOwnerCutPerMillion\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_publicationFee\",\"type\":\"uint256\"}],\"name\":\"setPublicationFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_marketplaceStorage\",\"type\":\"address\"}],\"name\":\"updateStorageAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"nftAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"assetId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startPriceInWei\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"biddingEnd\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"revealEnd\",\"type\":\"uint256\"}],\"name\":\"createAuction\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"nftAsset\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"auctionId\",\"type\":\"bytes32\"}],\"name\":\"cancelAuction\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"nftAsset\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"auctionId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"blindedBid\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"deposit\",\"type\":\"uint256\"}],\"name\":\"bidAuction\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"nftAsset\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"auctionId\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[]\",\"name\":\"_values\",\"type\":\"uint256[]\"},{\"internalType\":\"bool[]\",\"name\":\"_fake\",\"type\":\"bool[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"_secret\",\"type\":\"bytes32[]\"}],\"name\":\"revealBid\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"auctionId\",\"type\":\"bytes32\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"auctionId\",\"type\":\"bytes32\"}],\"name\":\"pendingReturns\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"nftAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"assetId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"auctionId\",\"type\":\"bytes32\"}],\"name\":\"closeAuction\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"auctionId\",\"type\":\"bytes32\"}],\"name\":\"checkExisted\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"nftAsset\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"auctionId\",\"type\":\"bytes32\"}],\"name\":\"checkRunning\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60806040523480156200001157600080fd5b50604051620072fe380380620072fe8339818101604052810190620000379190620014cb565b838284838180620000698173ffffffffffffffffffffffffffffffffffffffff166200050060201b620028431760201c565b620000ab576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401620000a290620017f7565b60405180910390fd5b60008190508073ffffffffffffffffffffffffffffffffffffffff166301ffc9a76345ad86c260e01b6040518263ffffffff1660e01b8152600401620000f2919062001750565b60206040518083038186803b1580156200010b57600080fd5b505afa15801562000120573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019062000146919062001537565b62000188576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016200017f90620017b3565b60405180910390fd5b82600260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505050506000600260146101000a81548160ff02191690831515021790555062000207620001fb6200051360201b60201c565b620005c660201b60201c565b620002276200021b6200051360201b60201c565b6200068c60201b60201c565b620002538473ffffffffffffffffffffffffffffffffffffffff166200050060201b620028431760201c565b6200028a576040517f6eefed2000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b83600560006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555081600260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550620003306000801b620003246200051360201b60201c565b6200085360201b60201c565b620003717fa49807205ce4d355092ef5a8a18f56e8913cf4a201fbe287825b095693c21775620003656200051360201b60201c565b6200085360201b60201c565b620003b27f9f959e00d95122f5cbd677010436cf273ef535b86b056afc172852144b9491d7620003a66200051360201b60201c565b6200085360201b60201c565b620003f37f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a620003e76200051360201b60201c565b6200085360201b60201c565b620004457f9f959e00d95122f5cbd677010436cf273ef535b86b056afc172852144b9491d77fa49807205ce4d355092ef5a8a18f56e8913cf4a201fbe287825b095693c217756200086960201b60201c565b620004977f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a7fa49807205ce4d355092ef5a8a18f56e8913cf4a201fbe287825b095693c217756200086960201b60201c565b620004a881620008cc60201b60201c565b610e1060088190555082600960006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550505050505050505062001c26565b600080823b905060008111915050919050565b60003073ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161415620005bf57600080368080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f820116905080830192505050505050509050600080369050905073ffffffffffffffffffffffffffffffffffffffff818301511692505050620005c3565b3390505b90565b6000600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905081600360006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b6200069c6200051360201b60201c565b73ffffffffffffffffffffffffffffffffffffffff16620006c26200099160201b60201c565b73ffffffffffffffffffffffffffffffffffffffff16146200071b576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401620007129062001819565b60405180910390fd5b6200074d7f9f959e00d95122f5cbd677010436cf273ef535b86b056afc172852144b9491d7826200085360201b60201c565b6200077f7f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a826200085360201b60201c565b620007946000801b826200085360201b60201c565b620007d57f9f959e00d95122f5cbd677010436cf273ef535b86b056afc172852144b9491d7620007c96200099160201b60201c565b620009bb60201b60201c565b620008167f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a6200080a6200099160201b60201c565b620009bb60201b60201c565b6200083a6000801b6200082e6200099160201b60201c565b620009bb60201b60201c565b620008508162000a0460201b620028561760201c565b50565b62000865828262000b1a60201b60201c565b5050565b60006200087c8362000b6260201b60201c565b905081600080858152602001908152602001600020600101819055508181847fbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff60405160405180910390a4505050565b7fa49807205ce4d355092ef5a8a18f56e8913cf4a201fbe287825b095693c217756200090e81620009026200051360201b60201c565b62000b8160201b60201c565b620f424082106200094b576040517f7e0d5ce600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b816006819055507ffa406a120a9e7f2b332bfb7a43d3bf1c3f079262202907a6b69c94b2821a02c66006546040516200098591906200183b565b60405180910390a15050565b6000600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b620009cc8262000b6260201b60201c565b620009ed81620009e16200051360201b60201c565b62000b8160201b60201c565b620009ff838362000c4560201b60201c565b505050565b62000a146200051360201b60201c565b73ffffffffffffffffffffffffffffffffffffffff1662000a3a6200099160201b60201c565b73ffffffffffffffffffffffffffffffffffffffff161462000a93576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040162000a8a9062001819565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16141562000b06576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040162000afd90620017d5565b60405180910390fd5b62000b1781620005c660201b60201c565b50565b62000b31828262000c8d60201b6200294e1760201c565b62000b5d816001600085815260200190815260200160002062000d7e60201b62002a2e1790919060201c565b505050565b6000806000838152602001908152602001600020600101549050919050565b62000b93828262000db660201b60201c565b62000c415762000bc68173ffffffffffffffffffffffffffffffffffffffff16601462000e2060201b62002a5e1760201c565b62000be18360001c602062000e2060201b62002a5e1760201c565b60405160200162000bf49291906200170e565b6040516020818303038152906040526040517f08c379a000000000000000000000000000000000000000000000000000000000815260040162000c3891906200176d565b60405180910390fd5b5050565b62000c5c82826200112f60201b62002d581760201c565b62000c8881600160008581526020019081526020016000206200122160201b62002e391790919060201c565b505050565b62000c9f828262000db660201b60201c565b62000d7a57600160008084815260200190815260200160002060000160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff02191690831515021790555062000d1f6200051360201b60201c565b73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16837f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a45b5050565b600062000dae836000018373ffffffffffffffffffffffffffffffffffffffff1660001b6200125960201b60201c565b905092915050565b600080600084815260200190815260200160002060000160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16905092915050565b60606000600283600262000e359190620018dc565b62000e4191906200187f565b67ffffffffffffffff81111562000e81577f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6040519080825280601f01601f19166020018201604052801562000eb45781602001600182028036833780820191505090505b5090507f30000000000000000000000000000000000000000000000000000000000000008160008151811062000f13577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a9053507f78000000000000000000000000000000000000000000000000000000000000008160018151811062000f9e577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a9053506000600184600262000fe09190620018dc565b62000fec91906200187f565b90505b6001811115620010de577f3031323334353637383961626364656600000000000000000000000000000000600f86166010811062001056577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b1a60f81b82828151811062001094577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a905350600485901c945080620010d69062001a24565b905062000fef565b506000841462001125576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016200111c9062001791565b60405180910390fd5b8091505092915050565b62001141828262000db660201b60201c565b156200121d57600080600084815260200190815260200160002060000160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff021916908315150217905550620011c26200051360201b60201c565b73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16837ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b60405160405180910390a45b5050565b600062001251836000018373ffffffffffffffffffffffffffffffffffffffff1660001b620012d360201b60201c565b905092915050565b60006200126d83836200146360201b60201c565b620012c8578260000182908060018154018082558091505060019003906000526020600020016000909190919091505582600001805490508360010160008481526020019081526020016000208190555060019050620012cd565b600090505b92915050565b60008083600101600084815260200190815260200160002054905060008114620014575760006001826200130891906200193d565b90506000600186600001805490506200132291906200193d565b9050818114620013e05760008660000182815481106200136b577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b9060005260206000200154905080876000018481548110620013b6577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b90600052602060002001819055508387600101600083815260200190815260200160002081905550505b856000018054806200141b577f4e487b7100000000000000000000000000000000000000000000000000000000600052603160045260246000fd5b6001900381819060005260206000200160009055905585600101600086815260200190815260200160002060009055600193505050506200145d565b60009150505b92915050565b600080836001016000848152602001908152602001600020541415905092915050565b600081519050620014978162001bd8565b92915050565b600081519050620014ae8162001bf2565b92915050565b600081519050620014c58162001c0c565b92915050565b60008060008060808587031215620014e257600080fd5b6000620014f28782880162001486565b9450506020620015058782880162001486565b9350506040620015188782880162001486565b92505060606200152b87828801620014b4565b91505092959194509250565b6000602082840312156200154a57600080fd5b60006200155a848285016200149d565b91505092915050565b6200156e8162001998565b82525050565b6000620015818262001858565b6200158d818562001863565b93506200159f818560208601620019ee565b620015aa8162001a82565b840191505092915050565b6000620015c28262001858565b620015ce818562001874565b9350620015e0818560208601620019ee565b80840191505092915050565b6000620015fb60208362001863565b9150620016088262001a93565b602082019050919050565b600062001622600f8362001863565b91506200162f8262001abc565b602082019050919050565b60006200164960268362001863565b9150620016568262001ae5565b604082019050919050565b60006200167060108362001863565b91506200167d8262001b34565b602082019050919050565b60006200169760208362001863565b9150620016a48262001b5d565b602082019050919050565b6000620016be60178362001874565b9150620016cb8262001b86565b601782019050919050565b6000620016e560118362001874565b9150620016f28262001baf565b601182019050919050565b6200170881620019e4565b82525050565b60006200171b82620016af565b9150620017298285620015b5565b91506200173682620016d6565b9150620017448284620015b5565b91508190509392505050565b600060208201905062001767600083018462001563565b92915050565b6000602082019050818103600083015262001789818462001574565b905092915050565b60006020820190508181036000830152620017ac81620015ec565b9050919050565b60006020820190508181036000830152620017ce8162001613565b9050919050565b60006020820190508181036000830152620017f0816200163a565b9050919050565b60006020820190508181036000830152620018128162001661565b9050919050565b60006020820190508181036000830152620018348162001688565b9050919050565b6000602082019050620018526000830184620016fd565b92915050565b600081519050919050565b600082825260208201905092915050565b600081905092915050565b60006200188c82620019e4565b91506200189983620019e4565b9250827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff03821115620018d157620018d062001a53565b5b828201905092915050565b6000620018e982620019e4565b9150620018f683620019e4565b9250817fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff048311821515161562001932576200193162001a53565b5b828202905092915050565b60006200194a82620019e4565b91506200195783620019e4565b9250828210156200196d576200196c62001a53565b5b828203905092915050565b60006200198582620019c4565b9050919050565b60008115159050919050565b60007fffffffff0000000000000000000000000000000000000000000000000000000082169050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b60005b8381101562001a0e578082015181840152602081019050620019f1565b8381111562001a1e576000848401525b50505050565b600062001a3182620019e4565b9150600082141562001a485762001a4762001a53565b5b600182039050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b6000601f19601f8301169050919050565b7f537472696e67733a20686578206c656e67746820696e73756666696369656e74600082015250565b7f496e76616c69642073746f726167650000000000000000000000000000000000600082015250565b7f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160008201527f6464726573730000000000000000000000000000000000000000000000000000602082015250565b7f496e76616c696420636f6e747261637400000000000000000000000000000000600082015250565b7f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572600082015250565b7f416363657373436f6e74726f6c3a206163636f756e7420000000000000000000600082015250565b7f206973206d697373696e6720726f6c6520000000000000000000000000000000600082015250565b62001be38162001978565b811462001bef57600080fd5b50565b62001bfd816200198c565b811462001c0957600080fd5b50565b62001c1781620019e4565b811462001c2357600080fd5b50565b6156c88062001c366000396000f3fe608060405234801561001057600080fd5b50600436106102735760003560e01c80638bafe55911610151578063ca15c873116100c3578063e47d606011610087578063e47d6060146106fc578063e4997dc51461072c578063e63ab1e914610748578063f2fde38b14610766578063f602188414610782578063fa07aa2b1461079e57610273565b8063ca15c8731461065a578063cbfda1c51461068a578063d547741f146106a8578063d6710e23146106c4578063daa3d985146106e057610273565b8063961c9ae411610115578063961c9ae4146105aa578063a01f79d4146105c6578063a217fddf146105e4578063ae4f119814610602578063af8996f114610620578063b42cf9291461063c57610273565b80638bafe559146104f25780638da5cb5b146105105780638e19899e1461052e5780639010d07c1461054a57806391d148541461057a57610273565b80633f4ba83a116101ea578063715018a6116101ae578063715018a61461046c57806375b238fc146104765780638456cb591461049457806385ec96a51461049e57806387b468ae146104ba578063889e2129146104d657610273565b80633f4ba83a146103d8578063451c3d80146103e257806346b3aec61461040057806359bf1abe1461041e5780635c975abb1461044e57610273565b806313c27ca71161023c57806313c27ca71461031857806319dad16d14610336578063248a9ca3146103525780632b4c32be146103825780632f2ff15d146103a057806336568abe146103bc57610273565b8062c5c3c31461027857806301ffc9a714610294578063076f6dc7146102c45780630db3cc05146102e05780630ecb93c0146102fc575b600080fd5b610292600480360381019061028d9190614558565b6107ce565b005b6102ae60048036038101906102a99190614757565b6108b2565b6040516102bb9190614db0565b60405180910390f35b6102de60048036038101906102d991906145bd565b61092c565b005b6102fa60048036038101906102f59190614417565b610a1c565b005b61031660048036038101906103119190614417565b610a93565b005b610320610ba1565b60405161032d9190614e63565b60405180910390f35b610350600480360381019061034b91906147a9565b610bc5565b005b61036c60048036038101906103679190614558565b610c77565b6040516103799190614e63565b60405180910390f35b61038a610c96565b6040516103979190614ea7565b60405180910390f35b6103ba60048036038101906103b59190614581565b610ca1565b005b6103d660048036038101906103d19190614581565b610cca565b005b6103e0610d4d565b005b6103ea610d8a565b6040516103f79190614ec2565b60405180910390f35b610408610db0565b6040516104159190614edd565b60405180910390f35b61043860048036038101906104339190614417565b610dd6565b6040516104459190614db0565b60405180910390f35b610456610e2c565b6040516104639190614db0565b60405180910390f35b610474610e43565b005b61047e610ecb565b60405161048b9190614e63565b60405180910390f35b61049c610eef565b005b6104b860048036038101906104b391906145f9565b610f2c565b005b6104d460048036038101906104cf91906146b8565b6113a5565b005b6104f060048036038101906104eb9190614417565b6116b5565b005b6104fa6116f4565b6040516105079190614ea7565b60405180910390f35b6105186116ff565b6040516105259190614b1b565b60405180910390f35b61054860048036038101906105439190614558565b611729565b005b610564600480360381019061055f919061471b565b6118da565b6040516105719190614b1b565b60405180910390f35b610594600480360381019061058f9190614581565b611909565b6040516105a19190614db0565b60405180910390f35b6105c460048036038101906105bf91906144b8565b611973565b005b6105ce6120dc565b6040516105db9190614fda565b60405180910390f35b6105ec6120e2565b6040516105f99190614e63565b60405180910390f35b61060a6120e9565b6040516106179190614fda565b60405180910390f35b61063a600480360381019061063591906147a9565b6120ef565b005b610644612165565b6040516106519190614fda565b60405180910390f35b610674600480360381019061066f9190614558565b61216b565b6040516106819190614fda565b60405180910390f35b61069261218f565b60405161069f9190614b1b565b60405180910390f35b6106c260048036038101906106bd9190614581565b6121b5565b005b6106de60048036038101906106d99190614469565b6121de565b005b6106fa60048036038101906106f591906145bd565b612268565b005b61071660048036038101906107119190614417565b6124eb565b6040516107239190614db0565b60405180910390f35b61074660048036038101906107419190614417565b61250b565b005b610750612619565b60405161075d9190614e63565b60405180910390f35b610780600480360381019061077b9190614417565b61263d565b005b61079c600480360381019061079791906147a9565b61279c565b005b6107b860048036038101906107b39190614558565b6127d9565b6040516107c59190614fda565b60405180910390f35b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16631a8b4169826040518263ffffffff1660e01b81526004016108299190614e63565b60206040518083038186803b15801561084157600080fd5b505afa158015610855573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610879919061452f565b6108af576040517fafdd489000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b50565b60007f5a05180f000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916827bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19161480610925575061092482612e69565b5b9050919050565b610935816107ce565b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663bc66c29f83836040518363ffffffff1660e01b8152600401610992929190614e7e565b60206040518083038186803b1580156109aa57600080fd5b505afa1580156109be573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906109e2919061452f565b610a18576040517f03b5e41300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5050565b7fa49807205ce4d355092ef5a8a18f56e8913cf4a201fbe287825b095693c21775610a4e81610a49612ee3565b612f94565b81600960006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505050565b610a9b612ee3565b73ffffffffffffffffffffffffffffffffffffffff16610ab96116ff565b73ffffffffffffffffffffffffffffffffffffffff1614610b0f576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610b0690614f9a565b60405180910390fd5b6001600460008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055507f42e160154868087d6bfdc0ca23d96a1c1cfa32f1b72ba9ba27b69b98a0d819dc81604051610b969190614b1b565b60405180910390a150565b7f9f959e00d95122f5cbd677010436cf273ef535b86b056afc172852144b9491d781565b7fa49807205ce4d355092ef5a8a18f56e8913cf4a201fbe287825b095693c21775610bf781610bf2612ee3565b612f94565b620f42408210610c33576040517f7e0d5ce600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b816006819055507ffa406a120a9e7f2b332bfb7a43d3bf1c3f079262202907a6b69c94b2821a02c6600654604051610c6b9190614fda565b60405180910390a15050565b6000806000838152602001908152602001600020600101549050919050565b6380ac58cd60e01b81565b610caa82610c77565b610cbb81610cb6612ee3565b612f94565b610cc58383613031565b505050565b610cd2612ee3565b73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614610d3f576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610d3690614fba565b60405180910390fd5b610d498282613065565b5050565b7f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a610d7f81610d7a612ee3565b612f94565b610d87613099565b50565b600560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6000600460008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff169050919050565b6000600260149054906101000a900460ff16905090565b610e4b612ee3565b73ffffffffffffffffffffffffffffffffffffffff16610e696116ff565b73ffffffffffffffffffffffffffffffffffffffff1614610ebf576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610eb690614f9a565b60405180910390fd5b610ec9600061313b565b565b7fa49807205ce4d355092ef5a8a18f56e8913cf4a201fbe287825b095693c2177581565b7f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a610f2181610f1c612ee3565b612f94565b610f29613201565b50565b610f34610e2c565b15610f74576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610f6b90614f7a565b60405180910390fd5b6000600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166315924b5b866040518263ffffffff1660e01b8152600401610fd19190614e63565b60e06040518083038186803b158015610fe957600080fd5b505afa158015610ffd573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906110219190614780565b905061103081606001516132a4565b600061103a612ee3565b90506000855190508085511461104f57600080fd5b8084511461105c57600080fd5b6000600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663c064bc508a8a6040518363ffffffff1660e01b81526004016110bb929190614e7e565b60206040518083038186803b1580156110d357600080fd5b505afa1580156110e7573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061110b919061452f565b9050600080600090505b838110156112a35761128e61127f89838151811061115c577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60200260200101518061116c5750845b876112388d86815181106111a9577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60200260200101518d87815181106111ea577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60200260200101518d888151811061122b577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60200260200101516132eb565b8d8681518110611271577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60200260200101518f613321565b836134da90919063ffffffff16565b9150808061129b9061532e565b915050611115565b50600081111561139957600560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663a9059cbb85836040518363ffffffff1660e01b815260040161130a929190614d50565b602060405180830381600087803b15801561132457600080fd5b505af1158015611338573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061135c919061452f565b507f6be22eb0241a0575f8b866ff176552cf71078977d6d30e501dd2d326c8d57114848a8360405161139093929190614d19565b60405180910390a15b50505050505050505050565b6113ad610e2c565b156113ed576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016113e490614f7a565b60405180910390fd5b6113f7848461092c565b6000600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166315924b5b856040518263ffffffff1660e01b81526004016114549190614e63565b60e06040518083038186803b15801561146c57600080fd5b505afa158015611480573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906114a49190614780565b6060015190506114b3816134f0565b60006114bd612ee3565b9050600560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166323b872dd8230866040518463ffffffff1660e01b815260040161151e93929190614c13565b602060405180830381600087803b15801561153857600080fd5b505af115801561154c573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611570919061452f565b506000600a600087815260200190815260200160002060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000868152602001908152602001600020541461160c576040517f089be2e000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b82600a600087815260200190815260200160002060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000868152602001908152602001600020819055507f327e5a459c964bcfa4f6d71b32a04a53b0002417b0924a8057d9f12e688e0bec8186866040516116a593929190614ce2565b60405180910390a1505050505050565b7fa49807205ce4d355092ef5a8a18f56e8913cf4a201fbe287825b095693c217756116e7816116e2612ee3565b612f94565b6116f082613537565b5050565b6345ad86c260e01b81565b6000600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b611731610e2c565b15611771576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161176890614f7a565b60405180910390fd5b600061177c826127d9565b90506000611788612ee3565b905060008211156118d5576000600b600085815260200190815260200160002060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002081905550600560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663a9059cbb82846040518363ffffffff1660e01b8152600401611846929190614d50565b602060405180830381600087803b15801561186057600080fd5b505af1158015611874573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611898919061452f565b507f6be22eb0241a0575f8b866ff176552cf71078977d6d30e501dd2d326c8d571148184846040516118cc93929190614d19565b60405180910390a15b505050565b600061190182600160008681526020019081526020016000206135b290919063ffffffff16565b905092915050565b600080600084815260200190815260200160002060000160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16905092915050565b846119938173ffffffffffffffffffffffffffffffffffffffff16612843565b6119c9576040517f6eefed2000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8073ffffffffffffffffffffffffffffffffffffffff166301ffc9a76380ac58cd60e01b6040518263ffffffff1660e01b8152600401611a099190614ea7565b60206040518083038186803b158015611a2157600080fd5b505afa158015611a35573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611a59919061452f565b611a6257600080fd5b611a6a610e2c565b15611aaa576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611aa190614f7a565b60405180910390fd5b6000611ac1600854426134da90919063ffffffff16565b841015611afa576040517f18c71d1a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b611b0f600854856134da90919063ffffffff16565b831015611b48576040517f4f2766ac00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6000851415611b82576040517ebfc92100000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6000611b8c612ee3565b905060008890508073ffffffffffffffffffffffffffffffffffffffff16636352211e896040518263ffffffff1660e01b8152600401611bcc9190614fda565b60206040518083038186803b158015611be457600080fd5b505afa158015611bf8573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611c1c9190614440565b92508273ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16141580611d9f57503073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1663081812fc8a6040518263ffffffff1660e01b8152600401611ca49190614fda565b60206040518083038186803b158015611cbc57600080fd5b505afa158015611cd0573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611cf49190614440565b73ffffffffffffffffffffffffffffffffffffffff161480611d9d57508073ffffffffffffffffffffffffffffffffffffffff1663e985e9c584306040518363ffffffff1660e01b8152600401611d4c929190614b36565b60206040518083038186803b158015611d6457600080fd5b505afa158015611d78573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611d9c919061452f565b5b155b15611dd6576040517f82b4290000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16630a7f89778a8a604051602001611e27929190614a19565b604051602081830303815290604052805190602001206040518263ffffffff1660e01b8152600401611e599190614e63565b60206040518083038186803b158015611e7157600080fd5b505afa158015611e85573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611ea9919061452f565b611edf576040517fa3b8915f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5060006007541115611fc257600560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166323b872dd82600960009054906101000a900473ffffffffffffffffffffffffffffffffffffffff166007546040518463ffffffff1660e01b8152600401611f6e93929190614c13565b602060405180830381600087803b158015611f8857600080fd5b505af1158015611f9c573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611fc0919061452f565b505b5060004282888a89604051602001611fde959493929190614a7f565b604051602081830303815290604052805190602001209050600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166384b04aa9838a8a858a8a8d6040518863ffffffff1660e01b815260040161205d9796959493929190614c4a565b600060405180830381600087803b15801561207757600080fd5b505af115801561208b573d6000803e3d6000fd5b505050507faf4bad306f14d5c908e5b871fa54296ed7d4f29b8092bf5062de6813c247e54b8289838a89898c6040516120ca9796959493929190614ba4565b60405180910390a15050505050505050565b60065481565b6000801b81565b60075481565b7fa49807205ce4d355092ef5a8a18f56e8913cf4a201fbe287825b095693c217756121218161211c612ee3565b612f94565b816007819055507fe7fa8737293f41b5dfa0d5c3e552860a06275bed7015581b083c7be7003308ba6007546040516121599190614fda565b60405180910390a15050565b60085481565b6000612188600160008481526020019081526020016000206135cc565b9050919050565b600960009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6121be82610c77565b6121cf816121ca612ee3565b612f94565b6121d98383613065565b505050565b6121e6610e2c565b15612226576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161221d90614f7a565b60405180910390fd5b61222f816107ce565b61226383838585604051602001612247929190614a19565b60405160208183030381529060405280519060200120846135e1565b505050565b612270610e2c565b156122b0576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016122a790614f7a565b60405180910390fd5b6122b9816107ce565b6122c3828261092c565b60006122cd612ee3565b90506000600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166315924b5b846040518263ffffffff1660e01b815260040161232c9190614e63565b60e06040518083038186803b15801561234457600080fd5b505afa158015612358573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061237c9190614780565b6000015190508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614806123e957506123e87f9f959e00d95122f5cbd677010436cf273ef535b86b056afc172852144b9491d76123e3612ee3565b611909565b5b61241f576040517f82b4290000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663ac84b7bc856040518263ffffffff1660e01b815260040161247a9190614e63565b600060405180830381600087803b15801561249457600080fd5b505af11580156124a8573d6000803e3d6000fd5b505050507f4283fca6e4b7cf10576815155991f12d5ef6adb5bcacdf255d3da62fddbf84ea82846040516124dd929190614cb9565b60405180910390a150505050565b60046020528060005260406000206000915054906101000a900460ff1681565b612513612ee3565b73ffffffffffffffffffffffffffffffffffffffff166125316116ff565b73ffffffffffffffffffffffffffffffffffffffff1614612587576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161257e90614f9a565b60405180910390fd5b6000600460008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055507fd7e9ec6e6ecd65492dce6bf513cd6867560d49544421d0783ddf06e76c24470c8160405161260e9190614b1b565b60405180910390a150565b7f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a81565b612645612ee3565b73ffffffffffffffffffffffffffffffffffffffff166126636116ff565b73ffffffffffffffffffffffffffffffffffffffff16146126b9576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016126b090614f9a565b60405180910390fd5b6126e37f9f959e00d95122f5cbd677010436cf273ef535b86b056afc172852144b9491d782613b37565b61270d7f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a82613b37565b61271a6000801b82613b37565b61274b7f9f959e00d95122f5cbd677010436cf273ef535b86b056afc172852144b9491d76127466116ff565b6121b5565b61277c7f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a6127776116ff565b6121b5565b6127906000801b61278b6116ff565b6121b5565b61279981612856565b50565b7fa49807205ce4d355092ef5a8a18f56e8913cf4a201fbe287825b095693c217756127ce816127c9612ee3565b612f94565b816008819055505050565b60006127e4826107ce565b600b60008381526020019081526020016000206000612801612ee3565b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020549050919050565b600080823b905060008111915050919050565b61285e612ee3565b73ffffffffffffffffffffffffffffffffffffffff1661287c6116ff565b73ffffffffffffffffffffffffffffffffffffffff16146128d2576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016128c990614f9a565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff161415612942576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161293990614f5a565b60405180910390fd5b61294b8161313b565b50565b6129588282611909565b612a2a57600160008084815260200190815260200160002060000160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055506129cf612ee3565b73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16837f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a45b5050565b6000612a56836000018373ffffffffffffffffffffffffffffffffffffffff1660001b613b45565b905092915050565b606060006002836002612a71919061514c565b612a7b91906150c5565b67ffffffffffffffff811115612aba577f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6040519080825280601f01601f191660200182016040528015612aec5781602001600182028036833780820191505090505b5090507f300000000000000000000000000000000000000000000000000000000000000081600081518110612b4a577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a9053507f780000000000000000000000000000000000000000000000000000000000000081600181518110612bd4577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a90535060006001846002612c14919061514c565b612c1e91906150c5565b90505b6001811115612d0a577f3031323334353637383961626364656600000000000000000000000000000000600f861660108110612c86577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b1a60f81b828281518110612cc3577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a905350600485901c945080612d03906152d3565b9050612c21565b5060008414612d4e576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612d4590614f1a565b60405180910390fd5b8091505092915050565b612d628282611909565b15612e3557600080600084815260200190815260200160002060000160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff021916908315150217905550612dda612ee3565b73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16837ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b60405160405180910390a45b5050565b6000612e61836000018373ffffffffffffffffffffffffffffffffffffffff1660001b613bb5565b905092915050565b60007f7965db0b000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916827bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19161480612edc5750612edb82613d3b565b5b9050919050565b60003073ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161415612f8d57600080368080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f820116905080830192505050505050509050600080369050905073ffffffffffffffffffffffffffffffffffffffff818301511692505050612f91565b3390505b90565b612f9e8282611909565b61302d57612fc38173ffffffffffffffffffffffffffffffffffffffff166014612a5e565b612fd18360001c6020612a5e565b604051602001612fe2929190614a45565b6040516020818303038152906040526040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016130249190614ef8565b60405180910390fd5b5050565b61303b828261294e565b6130608160016000858152602001908152602001600020612a2e90919063ffffffff16565b505050565b61306f8282612d58565b6130948160016000858152602001908152602001600020612e3990919063ffffffff16565b505050565b6130a1610e2c565b6130e0576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016130d790614f3a565b60405180910390fd5b6000600260146101000a81548160ff0219169083151502179055507f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa613124612ee3565b6040516131319190614b1b565b60405180910390a1565b6000600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905081600360006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b613209610e2c565b15613249576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161324090614f7a565b60405180910390fd5b6001600260146101000a81548160ff0219169083151502179055507f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a25861328d612ee3565b60405161329a9190614b1b565b60405180910390a1565b8042116132e857806040517f2a35a3240000000000000000000000000000000000000000000000000000000081526004016132df9190614fda565b60405180910390fd5b50565b600083838360405160200161330293929190614ade565b6040516020818303038152906040528051906020012090509392505050565b600080600a600084815260200190815260200160002060008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600086815260200190815260200160002054905060008114156133d8577f255ef988657daaeb7f921a0c776fa6efb221748fb64e61ab1febb0fd37f67175878785876040516133c69493929190614dcb565b60405180910390a160009150506134d1565b6000600a600085815260200190815260200160002060008873ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008781526020019081526020016000208190555061345281836134da90919063ffffffff16565b9150861580156134625750838110155b80156134755750613474838786613da5565b5b156134905761348d848361402190919063ffffffff16565b91505b7fabef59dc3ae014d197fad42649c58d34bfc816d3e1a7f26ca32c13611b13e7a187878587896040516134c7959493929190614e10565b60405180910390a1505b95945050505050565b600081836134e891906150c5565b905092915050565b80421061353457806040517f691e568200000000000000000000000000000000000000000000000000000000815260040161352b9190614fda565b60405180910390fd5b50565b80600260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055507f19e07a962d0e1cab8b7d4be06b77a7225ca7ebb744994cb6546fdd815590dc9d816040516135a79190614b1b565b60405180910390a150565b60006135c18360000183614037565b60001c905092915050565b60006135da82600001614088565b9050919050565b6135eb828261092c565b6000600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166315924b5b836040518263ffffffff1660e01b81526004016136489190614e63565b60e06040518083038186803b15801561366057600080fd5b505afa158015613674573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906136989190614780565b90506136a781608001516132a4565b60008160200151905060008260000151905060008360a00151905060008890508073ffffffffffffffffffffffffffffffffffffffff16636352211e896040518263ffffffff1660e01b81526004016137009190614fda565b60206040518083038186803b15801561371857600080fd5b505afa15801561372c573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906137509190614440565b73ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff16146137b4576040517f82b4290000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663ac84b7bc886040518263ffffffff1660e01b815260040161380f9190614e63565b600060405180830381600087803b15801561382957600080fd5b505af115801561383d573d6000803e3d6000fd5b50505050600073ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff1614613af4576000806006541115613980576138ab620f424061389d6006548661409990919063ffffffff16565b6140af90919063ffffffff16565b9050600560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663a9059cbb600960009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16836040518363ffffffff1660e01b815260040161392c929190614d50565b602060405180830381600087803b15801561394657600080fd5b505af115801561395a573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061397e919061452f565b505b600560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663a9059cbb856139d2848761402190919063ffffffff16565b6040518363ffffffff1660e01b81526004016139ef929190614d50565b602060405180830381600087803b158015613a0957600080fd5b505af1158015613a1d573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190613a41919061452f565b508173ffffffffffffffffffffffffffffffffffffffff166342842e0e85878c6040518463ffffffff1660e01b8152600401613a7f93929190614c13565b600060405180830381600087803b158015613a9957600080fd5b505af1158015613aad573d6000803e3d6000fd5b505050507f5ed769e966210f5d06d68e75831871131255cee237cb8edf78565aae170475ea84868986604051613ae69493929190614b5f565b60405180910390a150613b2c565b7f447d0298e6e2716a4343195759c35b277a60b76e628284d93262fa9869db992186604051613b239190614e63565b60405180910390a15b505050505050505050565b613b418282613031565b5050565b6000613b5183836140c5565b613baa578260000182908060018154018082558091505060019003906000526020600020016000909190919091505582600001805490508360010160008481526020019081526020016000208190555060019050613baf565b600090505b92915050565b60008083600101600084815260200190815260200160002054905060008114613d2f576000600182613be791906151a6565b9050600060018660000180549050613bff91906151a6565b9050818114613cba576000866000018281548110613c46577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b9060005260206000200154905080876000018481548110613c90577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b90600052602060002001819055508387600101600083815260200190815260200160002081905550505b85600001805480613cf4577f4e487b7100000000000000000000000000000000000000000000000000000000600052603160045260246000fd5b600190038181906000526020600020016000905590558560010160008681526020019081526020016000206000905560019350505050613d35565b60009150505b92915050565b60007f01ffc9a7000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916827bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916149050919050565b600080600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166315924b5b866040518263ffffffff1660e01b8152600401613e039190614e63565b60e06040518083038186803b158015613e1b57600080fd5b505afa158015613e2f573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190613e539190614780565b90508060c0015183111580613e6c57508060a001518311155b15613e7b57600091505061401a565b600073ffffffffffffffffffffffffffffffffffffffff16816020015173ffffffffffffffffffffffffffffffffffffffff1614613f7f57613f228160a00151600b6000846040015181526020019081526020016000206000846020015173ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020546134da90919063ffffffff16565b600b6000836040015181526020019081526020016000206000836020015173ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055505b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663572f7d5e858584604001516040518463ffffffff1660e01b8152600401613fe293929190614d79565b600060405180830381600087803b158015613ffc57600080fd5b505af1158015614010573d6000803e3d6000fd5b5050505060019150505b9392505050565b6000818361402f91906151a6565b905092915050565b6000826000018281548110614075577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b9060005260206000200154905092915050565b600081600001805490509050919050565b600081836140a7919061514c565b905092915050565b600081836140bd919061511b565b905092915050565b600080836001016000848152602001908152602001600020541415905092915050565b60006140fb6140f68461501a565b614ff5565b9050808382526020820190508285602086028201111561411a57600080fd5b60005b8581101561414a578161413088826142d4565b84526020840193506020830192505060018101905061411d565b5050509392505050565b600061416761416284615046565b614ff5565b9050808382526020820190508285602086028201111561418657600080fd5b60005b858110156141b6578161419c88826142fe565b845260208401935060208301925050600181019050614189565b5050509392505050565b60006141d36141ce84615072565b614ff5565b905080838252602082019050828560208602820111156141f257600080fd5b60005b85811015614222578161420888826143ed565b8452602084019350602083019250506001810190506141f5565b5050509392505050565b60008135905061423b8161561f565b92915050565b6000815190506142508161561f565b92915050565b600082601f83011261426757600080fd5b81356142778482602086016140e8565b91505092915050565b600082601f83011261429157600080fd5b81356142a1848260208601614154565b91505092915050565b600082601f8301126142bb57600080fd5b81356142cb8482602086016141c0565b91505092915050565b6000813590506142e381615636565b92915050565b6000815190506142f881615636565b92915050565b60008135905061430d8161564d565b92915050565b6000815190506143228161564d565b92915050565b60008135905061433781615664565b92915050565b600060e0828403121561434f57600080fd5b61435960e0614ff5565b9050600061436984828501614241565b600083015250602061437d84828501614241565b602083015250604061439184828501614313565b60408301525060606143a584828501614402565b60608301525060806143b984828501614402565b60808301525060a06143cd84828501614402565b60a08301525060c06143e184828501614402565b60c08301525092915050565b6000813590506143fc8161567b565b92915050565b6000815190506144118161567b565b92915050565b60006020828403121561442957600080fd5b60006144378482850161422c565b91505092915050565b60006020828403121561445257600080fd5b600061446084828501614241565b91505092915050565b60008060006060848603121561447e57600080fd5b600061448c8682870161422c565b935050602061449d868287016143ed565b92505060406144ae868287016142fe565b9150509250925092565b600080600080600060a086880312156144d057600080fd5b60006144de8882890161422c565b95505060206144ef888289016143ed565b9450506040614500888289016143ed565b9350506060614511888289016143ed565b9250506080614522888289016143ed565b9150509295509295909350565b60006020828403121561454157600080fd5b600061454f848285016142e9565b91505092915050565b60006020828403121561456a57600080fd5b6000614578848285016142fe565b91505092915050565b6000806040838503121561459457600080fd5b60006145a2858286016142fe565b92505060206145b38582860161422c565b9150509250929050565b600080604083850312156145d057600080fd5b60006145de858286016142fe565b92505060206145ef858286016142fe565b9150509250929050565b600080600080600060a0868803121561461157600080fd5b600061461f888289016142fe565b9550506020614630888289016142fe565b945050604086013567ffffffffffffffff81111561464d57600080fd5b614659888289016142aa565b935050606086013567ffffffffffffffff81111561467657600080fd5b61468288828901614256565b925050608086013567ffffffffffffffff81111561469f57600080fd5b6146ab88828901614280565b9150509295509295909350565b600080600080608085870312156146ce57600080fd5b60006146dc878288016142fe565b94505060206146ed878288016142fe565b93505060406146fe878288016142fe565b925050606061470f878288016143ed565b91505092959194509250565b6000806040838503121561472e57600080fd5b600061473c858286016142fe565b925050602061474d858286016143ed565b9150509250929050565b60006020828403121561476957600080fd5b600061477784828501614328565b91505092915050565b600060e0828403121561479257600080fd5b60006147a08482850161433d565b91505092915050565b6000602082840312156147bb57600080fd5b60006147c9848285016143ed565b91505092915050565b6147db816151da565b82525050565b6147f26147ed826151da565b615377565b82525050565b614801816151ec565b82525050565b614818614813826151ec565b615389565b82525050565b614827816151f8565b82525050565b61483e614839826151f8565b61539b565b82525050565b61484d81615202565b82525050565b61485c81615258565b82525050565b61486b8161527c565b82525050565b600061487c8261509e565b61488681856150a9565b93506148968185602086016152a0565b61489f81615460565b840191505092915050565b60006148b58261509e565b6148bf81856150ba565b93506148cf8185602086016152a0565b80840191505092915050565b60006148e86020836150a9565b91506148f38261548b565b602082019050919050565b600061490b6014836150a9565b9150614916826154b4565b602082019050919050565b600061492e6026836150a9565b9150614939826154dd565b604082019050919050565b60006149516010836150a9565b915061495c8261552c565b602082019050919050565b60006149746020836150a9565b915061497f82615555565b602082019050919050565b60006149976017836150ba565b91506149a28261557e565b601782019050919050565b60006149ba6011836150ba565b91506149c5826155a7565b601182019050919050565b60006149dd602f836150a9565b91506149e8826155d0565b604082019050919050565b6149fc8161524e565b82525050565b614a13614a0e8261524e565b6153b7565b82525050565b6000614a2582856147e1565b601482019150614a358284614a02565b6020820191508190509392505050565b6000614a508261498a565b9150614a5c82856148aa565b9150614a67826149ad565b9150614a7382846148aa565b91508190509392505050565b6000614a8b8288614a02565b602082019150614a9b82876147e1565b601482019150614aab8286614a02565b602082019150614abb82856147e1565b601482019150614acb8284614a02565b6020820191508190509695505050505050565b6000614aea8286614a02565b602082019150614afa8285614807565b600182019150614b0a828461482d565b602082019150819050949350505050565b6000602082019050614b3060008301846147d2565b92915050565b6000604082019050614b4b60008301856147d2565b614b5860208301846147d2565b9392505050565b6000608082019050614b7460008301876147d2565b614b8160208301866147d2565b614b8e604083018561481e565b614b9b60608301846149f3565b95945050505050565b600060e082019050614bb9600083018a6147d2565b614bc660208301896147d2565b614bd3604083018861481e565b614be060608301876149f3565b614bed60808301866149f3565b614bfa60a08301856149f3565b614c0760c08301846149f3565b98975050505050505050565b6000606082019050614c2860008301866147d2565b614c3560208301856147d2565b614c4260408301846149f3565b949350505050565b600060e082019050614c5f600083018a6147d2565b614c6c60208301896147d2565b614c7960408301886149f3565b614c86606083018761481e565b614c9360808301866149f3565b614ca060a08301856149f3565b614cad60c08301846149f3565b98975050505050505050565b6000604082019050614cce60008301856147d2565b614cdb602083018461481e565b9392505050565b6000606082019050614cf760008301866147d2565b614d04602083018561481e565b614d11604083018461481e565b949350505050565b6000606082019050614d2e60008301866147d2565b614d3b602083018561481e565b614d4860408301846149f3565b949350505050565b6000604082019050614d6560008301856147d2565b614d7260208301846149f3565b9392505050565b6000606082019050614d8e60008301866147d2565b614d9b60208301856149f3565b614da8604083018461481e565b949350505050565b6000602082019050614dc560008301846147f8565b92915050565b6000608082019050614de060008301876147f8565b614ded60208301866147d2565b614dfa604083018561481e565b614e0760608301846149f3565b95945050505050565b600060a082019050614e2560008301886147f8565b614e3260208301876147d2565b614e3f604083018661481e565b614e4c60608301856149f3565b614e59608083018461481e565b9695505050505050565b6000602082019050614e78600083018461481e565b92915050565b6000604082019050614e93600083018561481e565b614ea0602083018461481e565b9392505050565b6000602082019050614ebc6000830184614844565b92915050565b6000602082019050614ed76000830184614853565b92915050565b6000602082019050614ef26000830184614862565b92915050565b60006020820190508181036000830152614f128184614871565b905092915050565b60006020820190508181036000830152614f33816148db565b9050919050565b60006020820190508181036000830152614f53816148fe565b9050919050565b60006020820190508181036000830152614f7381614921565b9050919050565b60006020820190508181036000830152614f9381614944565b9050919050565b60006020820190508181036000830152614fb381614967565b9050919050565b60006020820190508181036000830152614fd3816149d0565b9050919050565b6000602082019050614fef60008301846149f3565b92915050565b6000614fff615010565b905061500b82826152fd565b919050565b6000604051905090565b600067ffffffffffffffff82111561503557615034615431565b5b602082029050602081019050919050565b600067ffffffffffffffff82111561506157615060615431565b5b602082029050602081019050919050565b600067ffffffffffffffff82111561508d5761508c615431565b5b602082029050602081019050919050565b600081519050919050565b600082825260208201905092915050565b600081905092915050565b60006150d08261524e565b91506150db8361524e565b9250827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff038211156151105761510f6153d3565b5b828201905092915050565b60006151268261524e565b91506151318361524e565b92508261514157615140615402565b5b828204905092915050565b60006151578261524e565b91506151628361524e565b9250817fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff048311821515161561519b5761519a6153d3565b5b828202905092915050565b60006151b18261524e565b91506151bc8361524e565b9250828210156151cf576151ce6153d3565b5b828203905092915050565b60006151e58261522e565b9050919050565b60008115159050919050565b6000819050919050565b60007fffffffff0000000000000000000000000000000000000000000000000000000082169050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b60006152638261526a565b9050919050565b60006152758261522e565b9050919050565b60006152878261528e565b9050919050565b60006152998261522e565b9050919050565b60005b838110156152be5780820151818401526020810190506152a3565b838111156152cd576000848401525b50505050565b60006152de8261524e565b915060008214156152f2576152f16153d3565b5b600182039050919050565b61530682615460565b810181811067ffffffffffffffff8211171561532557615324615431565b5b80604052505050565b60006153398261524e565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82141561536c5761536b6153d3565b5b600182019050919050565b6000615382826153a5565b9050919050565b6000615394826153c1565b9050919050565b6000819050919050565b60006153b08261547e565b9050919050565b6000819050919050565b60006153cc82615471565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6000601f19601f8301169050919050565b60008160f81b9050919050565b60008160601b9050919050565b7f537472696e67733a20686578206c656e67746820696e73756666696369656e74600082015250565b7f5061757361626c653a206e6f7420706175736564000000000000000000000000600082015250565b7f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160008201527f6464726573730000000000000000000000000000000000000000000000000000602082015250565b7f5061757361626c653a2070617573656400000000000000000000000000000000600082015250565b7f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572600082015250565b7f416363657373436f6e74726f6c3a206163636f756e7420000000000000000000600082015250565b7f206973206d697373696e6720726f6c6520000000000000000000000000000000600082015250565b7f416363657373436f6e74726f6c3a2063616e206f6e6c792072656e6f756e636560008201527f20726f6c657320666f722073656c660000000000000000000000000000000000602082015250565b615628816151da565b811461563357600080fd5b50565b61563f816151ec565b811461564a57600080fd5b50565b615656816151f8565b811461566157600080fd5b50565b61566d81615202565b811461567857600080fd5b50565b6156848161524e565b811461568f57600080fd5b5056fea2646970667358221220a3c86f6ddcc9000a9fd32b8539933a53a59eaa7cafefdc369724a5dda665a56464736f6c63430008040033",
}

// AuctionMarketplaceABI is the input ABI used to generate the binding from.
// Deprecated: Use AuctionMarketplaceMetaData.ABI instead.
var AuctionMarketplaceABI = AuctionMarketplaceMetaData.ABI

// AuctionMarketplaceBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use AuctionMarketplaceMetaData.Bin instead.
var AuctionMarketplaceBin = AuctionMarketplaceMetaData.Bin

// DeployAuctionMarketplace deploys a new Ethereum contract, binding an instance of AuctionMarketplace to it.
func DeployAuctionMarketplace(auth *bind.TransactOpts, backend bind.ContractBackend, acceptedToken common.Address, marketplaceStorage common.Address, beneficary common.Address, ownerCutPerMillion *big.Int) (common.Address, *types.Transaction, *AuctionMarketplace, error) {
	parsed, err := AuctionMarketplaceMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(AuctionMarketplaceBin), backend, acceptedToken, marketplaceStorage, beneficary, ownerCutPerMillion)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &AuctionMarketplace{AuctionMarketplaceCaller: AuctionMarketplaceCaller{contract: contract}, AuctionMarketplaceTransactor: AuctionMarketplaceTransactor{contract: contract}, AuctionMarketplaceFilterer: AuctionMarketplaceFilterer{contract: contract}}, nil
}

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
	Who       common.Address
	AuctionId [32]byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterAuctionCancelled is a free log retrieval operation binding the contract event 0x4283fca6e4b7cf10576815155991f12d5ef6adb5bcacdf255d3da62fddbf84ea.
//
// Solidity: event AuctionCancelled(address who, bytes32 auctionId)
func (_AuctionMarketplace *AuctionMarketplaceFilterer) FilterAuctionCancelled(opts *bind.FilterOpts) (*AuctionMarketplaceAuctionCancelledIterator, error) {

	logs, sub, err := _AuctionMarketplace.contract.FilterLogs(opts, "AuctionCancelled")
	if err != nil {
		return nil, err
	}
	return &AuctionMarketplaceAuctionCancelledIterator{contract: _AuctionMarketplace.contract, event: "AuctionCancelled", logs: logs, sub: sub}, nil
}

// WatchAuctionCancelled is a free log subscription operation binding the contract event 0x4283fca6e4b7cf10576815155991f12d5ef6adb5bcacdf255d3da62fddbf84ea.
//
// Solidity: event AuctionCancelled(address who, bytes32 auctionId)
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

// ParseAuctionCancelled is a log parse operation binding the contract event 0x4283fca6e4b7cf10576815155991f12d5ef6adb5bcacdf255d3da62fddbf84ea.
//
// Solidity: event AuctionCancelled(address who, bytes32 auctionId)
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
