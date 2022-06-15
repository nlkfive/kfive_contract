// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package BEP20Crowdsale

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

// BEP20CrowdsaleMetaData contains all meta data concerning the BEP20Crowdsale contract.
var BEP20CrowdsaleMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"__rate\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"__wallet\",\"type\":\"address\"},{\"internalType\":\"contractIBEP20\",\"name\":\"__token\",\"type\":\"address\"},{\"internalType\":\"contractIBEP20\",\"name\":\"__acceptToken\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"purchaser\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"beneficiary\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"TokensPurchased\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token\",\"outputs\":[{\"internalType\":\"contractIBEP20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"wallet\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"weiRaised\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"_to\",\"type\":\"address\"}],\"name\":\"destroySmartContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"newReceiver\",\"type\":\"address\"}],\"name\":\"changeWallet\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"beneficiary\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"buyTokens\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
	Bin: "0x60806040523480156200001157600080fd5b5060405162002133380380620021338339818101604052810190620000379190620003fd565b620000576200004b620002ec60201b60201c565b620002f460201b60201c565b600180819055506000600260006101000a81548160ff02191690831515021790555060008411620000bf576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401620000b69062000549565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff16141562000132576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401620001299062000505565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff161415620001a5576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016200019c9062000527565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16141562000218576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016200020f906200056b565b60405180910390fd5b8360058190555082600460006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555081600260016101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555080600360006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050505050620006f6565b600033905090565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050816000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b600081519050620003c981620006a8565b92915050565b600081519050620003e081620006c2565b92915050565b600081519050620003f781620006dc565b92915050565b600080600080608085870312156200041457600080fd5b60006200042487828801620003e6565b94505060206200043787828801620003b8565b93505060406200044a87828801620003cf565b92505060606200045d87828801620003cf565b91505092959194509250565b600062000478600f836200058d565b9150620004858262000604565b602082019050919050565b60006200049f600d836200058d565b9150620004ac826200062d565b602082019050919050565b6000620004c6600c836200058d565b9150620004d38262000656565b602082019050919050565b6000620004ed6014836200058d565b9150620004fa826200067f565b602082019050919050565b60006020820190508181036000830152620005208162000469565b9050919050565b60006020820190508181036000830152620005428162000490565b9050919050565b600060208201905081810360008301526200056481620004b7565b9050919050565b600060208201905081810360008301526200058681620004de565b9050919050565b600082825260208201905092915050565b6000620005ab82620005da565b9050919050565b6000620005bf82620005da565b9050919050565b6000620005d3826200059e565b9050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b7f496e76616c696420616464726573730000000000000000000000000000000000600082015250565b7f496e76616c696420746f6b656e00000000000000000000000000000000000000600082015250565b7f496e76616c696420726174650000000000000000000000000000000000000000600082015250565b7f496e76616c69642061636365707420746f6b656e000000000000000000000000600082015250565b620006b381620005b2565b8114620006bf57600080fd5b50565b620006cd81620005c6565b8114620006d957600080fd5b50565b620006e781620005fa565b8114620006f357600080fd5b50565b611a2d80620007066000396000f3fe6080604052600436106100c25760003560e01c80635c975abb1161007f5780638da5cb5b116100595780638da5cb5b146101fd57806398b9a2dc14610228578063f2fde38b14610251578063fc0c546a1461027a576100c2565b80635c975abb146101a4578063715018a6146101cf5780638456cb59146101e6576100c2565b80630752881a146100c75780632c4e722e146100e357806339df43ff1461010e5780633f4ba83a146101375780634042b66f1461014e578063521eb27314610179575b600080fd5b6100e160048036038101906100dc91906111f0565b6102a5565b005b3480156100ef57600080fd5b506100f8610494565b60405161010591906115e5565b60405180910390f35b34801561011a57600080fd5b50610135600480360381019061013091906111c7565b61049e565b005b34801561014357600080fd5b5061014c6107f1565b005b34801561015a57600080fd5b506101636108be565b60405161017091906115e5565b60405180910390f35b34801561018557600080fd5b5061018e6108c8565b60405161019b919061140b565b60405180910390f35b3480156101b057600080fd5b506101b96108f2565b6040516101c691906114af565b60405180910390f35b3480156101db57600080fd5b506101e4610909565b005b3480156101f257600080fd5b506101fb610991565b005b34801561020957600080fd5b50610212610a5f565b60405161021f91906113f0565b60405180910390f35b34801561023457600080fd5b5061024f600480360381019061024a91906111c7565b610a88565b005b34801561025d57600080fd5b506102786004803603810190610273919061119e565b610b48565b005b34801561028657600080fd5b5061028f610c40565b60405161029c91906114ca565b60405180910390f35b600260015414156102eb576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016102e2906115c5565b60405180910390fd5b60026001819055506102fb6108f2565b1561033b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161033290611545565b60405180910390fd5b670de0b6b3a7640000811015610386576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161037d90611525565b60405180910390fd5b60006103a3670de0b6b3a764000083610c6a90919063ffffffff16565b90506103af8382610c80565b60006103ba82610d38565b90506103e0846103db670de0b6b3a764000085610d5690919063ffffffff16565b610d6c565b6103ea8483610e44565b6103ff82600654610e4890919063ffffffff16565b60068190555061040f8482610e5e565b8373ffffffffffffffffffffffffffffffffffffffff1661042e610e6c565b73ffffffffffffffffffffffffffffffffffffffff167f6faf93231a456e552dbc9961f58d9713ee4f2e69d15f1975b050ef0911053a7b8484604051610475929190611600565b60405180910390a36104878483610e74565b5050600180819055505050565b6000600554905090565b6104a6610e6c565b73ffffffffffffffffffffffffffffffffffffffff166104c4610a5f565b73ffffffffffffffffffffffffffffffffffffffff161461051a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161051190611585565b60405180910390fd5b6000600260019054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166370a08231306040518263ffffffff1660e01b815260040161057791906113f0565b60206040518083038186803b15801561058f57600080fd5b505afa1580156105a3573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906105c79190611255565b9050600260019054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663a9059cbb83836040518363ffffffff1660e01b8152600401610626929190611426565b602060405180830381600087803b15801561064057600080fd5b505af1158015610654573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610678919061122c565b506000600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166370a08231306040518263ffffffff1660e01b81526004016106d691906113f0565b60206040518083038186803b1580156106ee57600080fd5b505afa158015610702573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906107269190611255565b9050600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663a9059cbb84836040518363ffffffff1660e01b8152600401610785929190611426565b602060405180830381600087803b15801561079f57600080fd5b505af11580156107b3573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906107d7919061122c565b508273ffffffffffffffffffffffffffffffffffffffff16ff5b6107f96108f2565b610838576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161082f906114e5565b60405180910390fd5b610840610e6c565b73ffffffffffffffffffffffffffffffffffffffff1661085e610a5f565b73ffffffffffffffffffffffffffffffffffffffff16146108b4576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016108ab90611585565b60405180910390fd5b6108bc610e78565b565b6000600654905090565b6000600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b6000600260009054906101000a900460ff16905090565b610911610e6c565b73ffffffffffffffffffffffffffffffffffffffff1661092f610a5f565b73ffffffffffffffffffffffffffffffffffffffff1614610985576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161097c90611585565b60405180910390fd5b61098f6000610f1a565b565b6109996108f2565b156109d9576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016109d090611545565b60405180910390fd5b6109e1610e6c565b73ffffffffffffffffffffffffffffffffffffffff166109ff610a5f565b73ffffffffffffffffffffffffffffffffffffffff1614610a55576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610a4c90611585565b60405180910390fd5b610a5d610fde565b565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b610a90610e6c565b73ffffffffffffffffffffffffffffffffffffffff16610aae610a5f565b73ffffffffffffffffffffffffffffffffffffffff1614610b04576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610afb90611585565b60405180910390fd5b80600460006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b610b50610e6c565b73ffffffffffffffffffffffffffffffffffffffff16610b6e610a5f565b73ffffffffffffffffffffffffffffffffffffffff1614610bc4576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610bbb90611585565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff161415610c34576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610c2b90611505565b60405180910390fd5b610c3d81610f1a565b50565b6000600260019054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b60008183610c789190611690565b905092915050565b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff161415610cf0576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610ce790611565565b60405180910390fd5b6000811415610d34576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610d2b906115a5565b60405180910390fd5b5050565b6000610d4f60055483610d5690919063ffffffff16565b9050919050565b60008183610d6491906116c1565b905092915050565b600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166323b872dd83600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16846040518463ffffffff1660e01b8152600401610ded9392919061144f565b602060405180830381600087803b158015610e0757600080fd5b505af1158015610e1b573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610e3f919061122c565b505050565b5050565b60008183610e56919061163a565b905092915050565b610e688282611081565b5050565b600033905090565b5050565b610e806108f2565b610ebf576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610eb6906114e5565b60405180910390fd5b6000600260006101000a81548160ff0219169083151502179055507f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa610f03610e6c565b604051610f1091906113f0565b60405180910390a1565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050816000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b610fe66108f2565b15611026576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161101d90611545565b60405180910390fd5b6001600260006101000a81548160ff0219169083151502179055507f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a25861106a610e6c565b60405161107791906113f0565b60405180910390a1565b600260019054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663a9059cbb83836040518363ffffffff1660e01b81526004016110de929190611486565b602060405180830381600087803b1580156110f857600080fd5b505af115801561110c573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611130919061122c565b505050565b6000813590506111448161199b565b92915050565b600081359050611159816119b2565b92915050565b60008151905061116e816119c9565b92915050565b600081359050611183816119e0565b92915050565b600081519050611198816119e0565b92915050565b6000602082840312156111b057600080fd5b60006111be84828501611135565b91505092915050565b6000602082840312156111d957600080fd5b60006111e78482850161114a565b91505092915050565b6000806040838503121561120357600080fd5b600061121185828601611135565b925050602061122285828601611174565b9150509250929050565b60006020828403121561123e57600080fd5b600061124c8482850161115f565b91505092915050565b60006020828403121561126757600080fd5b600061127584828501611189565b91505092915050565b61128781611775565b82525050565b6112968161172d565b82525050565b6112a58161171b565b82525050565b6112b48161173f565b82525050565b6112c381611787565b82525050565b60006112d6601483611629565b91506112e18261182d565b602082019050919050565b60006112f9602683611629565b915061130482611856565b604082019050919050565b600061131c600e83611629565b9150611327826118a5565b602082019050919050565b600061133f601083611629565b915061134a826118ce565b602082019050919050565b6000611362601383611629565b915061136d826118f7565b602082019050919050565b6000611385602083611629565b915061139082611920565b602082019050919050565b60006113a8600783611629565b91506113b382611949565b602082019050919050565b60006113cb601f83611629565b91506113d682611972565b602082019050919050565b6113ea8161176b565b82525050565b6000602082019050611405600083018461129c565b92915050565b6000602082019050611420600083018461128d565b92915050565b600060408201905061143b600083018561127e565b61144860208301846113e1565b9392505050565b6000606082019050611464600083018661129c565b611471602083018561127e565b61147e60408301846113e1565b949350505050565b600060408201905061149b600083018561129c565b6114a860208301846113e1565b9392505050565b60006020820190506114c460008301846112ab565b92915050565b60006020820190506114df60008301846112ba565b92915050565b600060208201905081810360008301526114fe816112c9565b9050919050565b6000602082019050818103600083015261151e816112ec565b9050919050565b6000602082019050818103600083015261153e8161130f565b9050919050565b6000602082019050818103600083015261155e81611332565b9050919050565b6000602082019050818103600083015261157e81611355565b9050919050565b6000602082019050818103600083015261159e81611378565b9050919050565b600060208201905081810360008301526115be8161139b565b9050919050565b600060208201905081810360008301526115de816113be565b9050919050565b60006020820190506115fa60008301846113e1565b92915050565b600060408201905061161560008301856113e1565b61162260208301846113e1565b9392505050565b600082825260208201905092915050565b60006116458261176b565b91506116508361176b565b9250827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff03821115611685576116846117cf565b5b828201905092915050565b600061169b8261176b565b91506116a68361176b565b9250826116b6576116b56117fe565b5b828204905092915050565b60006116cc8261176b565b91506116d78361176b565b9250817fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff04831182151516156117105761170f6117cf565b5b828202905092915050565b60006117268261174b565b9050919050565b60006117388261174b565b9050919050565b60008115159050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b6000611780826117ab565b9050919050565b600061179282611799565b9050919050565b60006117a48261174b565b9050919050565b60006117b6826117bd565b9050919050565b60006117c88261174b565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b7f5061757361626c653a206e6f7420706175736564000000000000000000000000600082015250565b7f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160008201527f6464726573730000000000000000000000000000000000000000000000000000602082015250565b7f496e76616c696420616d6f756e74000000000000000000000000000000000000600082015250565b7f5061757361626c653a2070617573656400000000000000000000000000000000600082015250565b7f496e76616c69642062656e656669636961727900000000000000000000000000600082015250565b7f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572600082015250565b7f496e76616c696400000000000000000000000000000000000000000000000000600082015250565b7f5265656e7472616e637947756172643a207265656e7472616e742063616c6c00600082015250565b6119a48161171b565b81146119af57600080fd5b50565b6119bb8161172d565b81146119c657600080fd5b50565b6119d28161173f565b81146119dd57600080fd5b50565b6119e98161176b565b81146119f457600080fd5b5056fea2646970667358221220a5b8004caf65b3eb69a8295cd9ad217b20987f88e18b3c3b8266f87a036f89fb64736f6c63430008040033",
}

// BEP20CrowdsaleABI is the input ABI used to generate the binding from.
// Deprecated: Use BEP20CrowdsaleMetaData.ABI instead.
var BEP20CrowdsaleABI = BEP20CrowdsaleMetaData.ABI

// BEP20CrowdsaleBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use BEP20CrowdsaleMetaData.Bin instead.
var BEP20CrowdsaleBin = BEP20CrowdsaleMetaData.Bin

// DeployBEP20Crowdsale deploys a new Ethereum contract, binding an instance of BEP20Crowdsale to it.
func DeployBEP20Crowdsale(auth *bind.TransactOpts, backend bind.ContractBackend, __rate *big.Int, __wallet common.Address, __token common.Address, __acceptToken common.Address) (common.Address, *types.Transaction, *BEP20Crowdsale, error) {
	parsed, err := BEP20CrowdsaleMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(BEP20CrowdsaleBin), backend, __rate, __wallet, __token, __acceptToken)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &BEP20Crowdsale{BEP20CrowdsaleCaller: BEP20CrowdsaleCaller{contract: contract}, BEP20CrowdsaleTransactor: BEP20CrowdsaleTransactor{contract: contract}, BEP20CrowdsaleFilterer: BEP20CrowdsaleFilterer{contract: contract}}, nil
}

// BEP20Crowdsale is an auto generated Go binding around an Ethereum contract.
type BEP20Crowdsale struct {
	BEP20CrowdsaleCaller     // Read-only binding to the contract
	BEP20CrowdsaleTransactor // Write-only binding to the contract
	BEP20CrowdsaleFilterer   // Log filterer for contract events
}

// BEP20CrowdsaleCaller is an auto generated read-only Go binding around an Ethereum contract.
type BEP20CrowdsaleCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BEP20CrowdsaleTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BEP20CrowdsaleTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BEP20CrowdsaleFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BEP20CrowdsaleFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BEP20CrowdsaleSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BEP20CrowdsaleSession struct {
	Contract     *BEP20Crowdsale   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BEP20CrowdsaleCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BEP20CrowdsaleCallerSession struct {
	Contract *BEP20CrowdsaleCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// BEP20CrowdsaleTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BEP20CrowdsaleTransactorSession struct {
	Contract     *BEP20CrowdsaleTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// BEP20CrowdsaleRaw is an auto generated low-level Go binding around an Ethereum contract.
type BEP20CrowdsaleRaw struct {
	Contract *BEP20Crowdsale // Generic contract binding to access the raw methods on
}

// BEP20CrowdsaleCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BEP20CrowdsaleCallerRaw struct {
	Contract *BEP20CrowdsaleCaller // Generic read-only contract binding to access the raw methods on
}

// BEP20CrowdsaleTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BEP20CrowdsaleTransactorRaw struct {
	Contract *BEP20CrowdsaleTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBEP20Crowdsale creates a new instance of BEP20Crowdsale, bound to a specific deployed contract.
func NewBEP20Crowdsale(address common.Address, backend bind.ContractBackend) (*BEP20Crowdsale, error) {
	contract, err := bindBEP20Crowdsale(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BEP20Crowdsale{BEP20CrowdsaleCaller: BEP20CrowdsaleCaller{contract: contract}, BEP20CrowdsaleTransactor: BEP20CrowdsaleTransactor{contract: contract}, BEP20CrowdsaleFilterer: BEP20CrowdsaleFilterer{contract: contract}}, nil
}

// NewBEP20CrowdsaleCaller creates a new read-only instance of BEP20Crowdsale, bound to a specific deployed contract.
func NewBEP20CrowdsaleCaller(address common.Address, caller bind.ContractCaller) (*BEP20CrowdsaleCaller, error) {
	contract, err := bindBEP20Crowdsale(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BEP20CrowdsaleCaller{contract: contract}, nil
}

// NewBEP20CrowdsaleTransactor creates a new write-only instance of BEP20Crowdsale, bound to a specific deployed contract.
func NewBEP20CrowdsaleTransactor(address common.Address, transactor bind.ContractTransactor) (*BEP20CrowdsaleTransactor, error) {
	contract, err := bindBEP20Crowdsale(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BEP20CrowdsaleTransactor{contract: contract}, nil
}

// NewBEP20CrowdsaleFilterer creates a new log filterer instance of BEP20Crowdsale, bound to a specific deployed contract.
func NewBEP20CrowdsaleFilterer(address common.Address, filterer bind.ContractFilterer) (*BEP20CrowdsaleFilterer, error) {
	contract, err := bindBEP20Crowdsale(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BEP20CrowdsaleFilterer{contract: contract}, nil
}

// bindBEP20Crowdsale binds a generic wrapper to an already deployed contract.
func bindBEP20Crowdsale(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BEP20CrowdsaleABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BEP20Crowdsale *BEP20CrowdsaleRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BEP20Crowdsale.Contract.BEP20CrowdsaleCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BEP20Crowdsale *BEP20CrowdsaleRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BEP20Crowdsale.Contract.BEP20CrowdsaleTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BEP20Crowdsale *BEP20CrowdsaleRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BEP20Crowdsale.Contract.BEP20CrowdsaleTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BEP20Crowdsale *BEP20CrowdsaleCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BEP20Crowdsale.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BEP20Crowdsale *BEP20CrowdsaleTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BEP20Crowdsale.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BEP20Crowdsale *BEP20CrowdsaleTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BEP20Crowdsale.Contract.contract.Transact(opts, method, params...)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BEP20Crowdsale *BEP20CrowdsaleCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BEP20Crowdsale.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BEP20Crowdsale *BEP20CrowdsaleSession) Owner() (common.Address, error) {
	return _BEP20Crowdsale.Contract.Owner(&_BEP20Crowdsale.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BEP20Crowdsale *BEP20CrowdsaleCallerSession) Owner() (common.Address, error) {
	return _BEP20Crowdsale.Contract.Owner(&_BEP20Crowdsale.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_BEP20Crowdsale *BEP20CrowdsaleCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _BEP20Crowdsale.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_BEP20Crowdsale *BEP20CrowdsaleSession) Paused() (bool, error) {
	return _BEP20Crowdsale.Contract.Paused(&_BEP20Crowdsale.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_BEP20Crowdsale *BEP20CrowdsaleCallerSession) Paused() (bool, error) {
	return _BEP20Crowdsale.Contract.Paused(&_BEP20Crowdsale.CallOpts)
}

// Rate is a free data retrieval call binding the contract method 0x2c4e722e.
//
// Solidity: function rate() view returns(uint256)
func (_BEP20Crowdsale *BEP20CrowdsaleCaller) Rate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BEP20Crowdsale.contract.Call(opts, &out, "rate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Rate is a free data retrieval call binding the contract method 0x2c4e722e.
//
// Solidity: function rate() view returns(uint256)
func (_BEP20Crowdsale *BEP20CrowdsaleSession) Rate() (*big.Int, error) {
	return _BEP20Crowdsale.Contract.Rate(&_BEP20Crowdsale.CallOpts)
}

// Rate is a free data retrieval call binding the contract method 0x2c4e722e.
//
// Solidity: function rate() view returns(uint256)
func (_BEP20Crowdsale *BEP20CrowdsaleCallerSession) Rate() (*big.Int, error) {
	return _BEP20Crowdsale.Contract.Rate(&_BEP20Crowdsale.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_BEP20Crowdsale *BEP20CrowdsaleCaller) Token(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BEP20Crowdsale.contract.Call(opts, &out, "token")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_BEP20Crowdsale *BEP20CrowdsaleSession) Token() (common.Address, error) {
	return _BEP20Crowdsale.Contract.Token(&_BEP20Crowdsale.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_BEP20Crowdsale *BEP20CrowdsaleCallerSession) Token() (common.Address, error) {
	return _BEP20Crowdsale.Contract.Token(&_BEP20Crowdsale.CallOpts)
}

// Wallet is a free data retrieval call binding the contract method 0x521eb273.
//
// Solidity: function wallet() view returns(address)
func (_BEP20Crowdsale *BEP20CrowdsaleCaller) Wallet(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BEP20Crowdsale.contract.Call(opts, &out, "wallet")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Wallet is a free data retrieval call binding the contract method 0x521eb273.
//
// Solidity: function wallet() view returns(address)
func (_BEP20Crowdsale *BEP20CrowdsaleSession) Wallet() (common.Address, error) {
	return _BEP20Crowdsale.Contract.Wallet(&_BEP20Crowdsale.CallOpts)
}

// Wallet is a free data retrieval call binding the contract method 0x521eb273.
//
// Solidity: function wallet() view returns(address)
func (_BEP20Crowdsale *BEP20CrowdsaleCallerSession) Wallet() (common.Address, error) {
	return _BEP20Crowdsale.Contract.Wallet(&_BEP20Crowdsale.CallOpts)
}

// WeiRaised is a free data retrieval call binding the contract method 0x4042b66f.
//
// Solidity: function weiRaised() view returns(uint256)
func (_BEP20Crowdsale *BEP20CrowdsaleCaller) WeiRaised(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BEP20Crowdsale.contract.Call(opts, &out, "weiRaised")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WeiRaised is a free data retrieval call binding the contract method 0x4042b66f.
//
// Solidity: function weiRaised() view returns(uint256)
func (_BEP20Crowdsale *BEP20CrowdsaleSession) WeiRaised() (*big.Int, error) {
	return _BEP20Crowdsale.Contract.WeiRaised(&_BEP20Crowdsale.CallOpts)
}

// WeiRaised is a free data retrieval call binding the contract method 0x4042b66f.
//
// Solidity: function weiRaised() view returns(uint256)
func (_BEP20Crowdsale *BEP20CrowdsaleCallerSession) WeiRaised() (*big.Int, error) {
	return _BEP20Crowdsale.Contract.WeiRaised(&_BEP20Crowdsale.CallOpts)
}

// BuyTokens is a paid mutator transaction binding the contract method 0x0752881a.
//
// Solidity: function buyTokens(address beneficiary, uint256 value) payable returns()
func (_BEP20Crowdsale *BEP20CrowdsaleTransactor) BuyTokens(opts *bind.TransactOpts, beneficiary common.Address, value *big.Int) (*types.Transaction, error) {
	return _BEP20Crowdsale.contract.Transact(opts, "buyTokens", beneficiary, value)
}

// BuyTokens is a paid mutator transaction binding the contract method 0x0752881a.
//
// Solidity: function buyTokens(address beneficiary, uint256 value) payable returns()
func (_BEP20Crowdsale *BEP20CrowdsaleSession) BuyTokens(beneficiary common.Address, value *big.Int) (*types.Transaction, error) {
	return _BEP20Crowdsale.Contract.BuyTokens(&_BEP20Crowdsale.TransactOpts, beneficiary, value)
}

// BuyTokens is a paid mutator transaction binding the contract method 0x0752881a.
//
// Solidity: function buyTokens(address beneficiary, uint256 value) payable returns()
func (_BEP20Crowdsale *BEP20CrowdsaleTransactorSession) BuyTokens(beneficiary common.Address, value *big.Int) (*types.Transaction, error) {
	return _BEP20Crowdsale.Contract.BuyTokens(&_BEP20Crowdsale.TransactOpts, beneficiary, value)
}

// ChangeWallet is a paid mutator transaction binding the contract method 0x98b9a2dc.
//
// Solidity: function changeWallet(address newReceiver) returns()
func (_BEP20Crowdsale *BEP20CrowdsaleTransactor) ChangeWallet(opts *bind.TransactOpts, newReceiver common.Address) (*types.Transaction, error) {
	return _BEP20Crowdsale.contract.Transact(opts, "changeWallet", newReceiver)
}

// ChangeWallet is a paid mutator transaction binding the contract method 0x98b9a2dc.
//
// Solidity: function changeWallet(address newReceiver) returns()
func (_BEP20Crowdsale *BEP20CrowdsaleSession) ChangeWallet(newReceiver common.Address) (*types.Transaction, error) {
	return _BEP20Crowdsale.Contract.ChangeWallet(&_BEP20Crowdsale.TransactOpts, newReceiver)
}

// ChangeWallet is a paid mutator transaction binding the contract method 0x98b9a2dc.
//
// Solidity: function changeWallet(address newReceiver) returns()
func (_BEP20Crowdsale *BEP20CrowdsaleTransactorSession) ChangeWallet(newReceiver common.Address) (*types.Transaction, error) {
	return _BEP20Crowdsale.Contract.ChangeWallet(&_BEP20Crowdsale.TransactOpts, newReceiver)
}

// DestroySmartContract is a paid mutator transaction binding the contract method 0x39df43ff.
//
// Solidity: function destroySmartContract(address _to) returns()
func (_BEP20Crowdsale *BEP20CrowdsaleTransactor) DestroySmartContract(opts *bind.TransactOpts, _to common.Address) (*types.Transaction, error) {
	return _BEP20Crowdsale.contract.Transact(opts, "destroySmartContract", _to)
}

// DestroySmartContract is a paid mutator transaction binding the contract method 0x39df43ff.
//
// Solidity: function destroySmartContract(address _to) returns()
func (_BEP20Crowdsale *BEP20CrowdsaleSession) DestroySmartContract(_to common.Address) (*types.Transaction, error) {
	return _BEP20Crowdsale.Contract.DestroySmartContract(&_BEP20Crowdsale.TransactOpts, _to)
}

// DestroySmartContract is a paid mutator transaction binding the contract method 0x39df43ff.
//
// Solidity: function destroySmartContract(address _to) returns()
func (_BEP20Crowdsale *BEP20CrowdsaleTransactorSession) DestroySmartContract(_to common.Address) (*types.Transaction, error) {
	return _BEP20Crowdsale.Contract.DestroySmartContract(&_BEP20Crowdsale.TransactOpts, _to)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_BEP20Crowdsale *BEP20CrowdsaleTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BEP20Crowdsale.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_BEP20Crowdsale *BEP20CrowdsaleSession) Pause() (*types.Transaction, error) {
	return _BEP20Crowdsale.Contract.Pause(&_BEP20Crowdsale.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_BEP20Crowdsale *BEP20CrowdsaleTransactorSession) Pause() (*types.Transaction, error) {
	return _BEP20Crowdsale.Contract.Pause(&_BEP20Crowdsale.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_BEP20Crowdsale *BEP20CrowdsaleTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BEP20Crowdsale.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_BEP20Crowdsale *BEP20CrowdsaleSession) RenounceOwnership() (*types.Transaction, error) {
	return _BEP20Crowdsale.Contract.RenounceOwnership(&_BEP20Crowdsale.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_BEP20Crowdsale *BEP20CrowdsaleTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _BEP20Crowdsale.Contract.RenounceOwnership(&_BEP20Crowdsale.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_BEP20Crowdsale *BEP20CrowdsaleTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _BEP20Crowdsale.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_BEP20Crowdsale *BEP20CrowdsaleSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _BEP20Crowdsale.Contract.TransferOwnership(&_BEP20Crowdsale.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_BEP20Crowdsale *BEP20CrowdsaleTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _BEP20Crowdsale.Contract.TransferOwnership(&_BEP20Crowdsale.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_BEP20Crowdsale *BEP20CrowdsaleTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BEP20Crowdsale.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_BEP20Crowdsale *BEP20CrowdsaleSession) Unpause() (*types.Transaction, error) {
	return _BEP20Crowdsale.Contract.Unpause(&_BEP20Crowdsale.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_BEP20Crowdsale *BEP20CrowdsaleTransactorSession) Unpause() (*types.Transaction, error) {
	return _BEP20Crowdsale.Contract.Unpause(&_BEP20Crowdsale.TransactOpts)
}

// BEP20CrowdsaleOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the BEP20Crowdsale contract.
type BEP20CrowdsaleOwnershipTransferredIterator struct {
	Event *BEP20CrowdsaleOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *BEP20CrowdsaleOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BEP20CrowdsaleOwnershipTransferred)
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
		it.Event = new(BEP20CrowdsaleOwnershipTransferred)
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
func (it *BEP20CrowdsaleOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BEP20CrowdsaleOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BEP20CrowdsaleOwnershipTransferred represents a OwnershipTransferred event raised by the BEP20Crowdsale contract.
type BEP20CrowdsaleOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_BEP20Crowdsale *BEP20CrowdsaleFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*BEP20CrowdsaleOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _BEP20Crowdsale.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &BEP20CrowdsaleOwnershipTransferredIterator{contract: _BEP20Crowdsale.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_BEP20Crowdsale *BEP20CrowdsaleFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *BEP20CrowdsaleOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _BEP20Crowdsale.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BEP20CrowdsaleOwnershipTransferred)
				if err := _BEP20Crowdsale.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_BEP20Crowdsale *BEP20CrowdsaleFilterer) ParseOwnershipTransferred(log types.Log) (*BEP20CrowdsaleOwnershipTransferred, error) {
	event := new(BEP20CrowdsaleOwnershipTransferred)
	if err := _BEP20Crowdsale.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BEP20CrowdsalePausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the BEP20Crowdsale contract.
type BEP20CrowdsalePausedIterator struct {
	Event *BEP20CrowdsalePaused // Event containing the contract specifics and raw log

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
func (it *BEP20CrowdsalePausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BEP20CrowdsalePaused)
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
		it.Event = new(BEP20CrowdsalePaused)
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
func (it *BEP20CrowdsalePausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BEP20CrowdsalePausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BEP20CrowdsalePaused represents a Paused event raised by the BEP20Crowdsale contract.
type BEP20CrowdsalePaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_BEP20Crowdsale *BEP20CrowdsaleFilterer) FilterPaused(opts *bind.FilterOpts) (*BEP20CrowdsalePausedIterator, error) {

	logs, sub, err := _BEP20Crowdsale.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &BEP20CrowdsalePausedIterator{contract: _BEP20Crowdsale.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_BEP20Crowdsale *BEP20CrowdsaleFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *BEP20CrowdsalePaused) (event.Subscription, error) {

	logs, sub, err := _BEP20Crowdsale.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BEP20CrowdsalePaused)
				if err := _BEP20Crowdsale.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_BEP20Crowdsale *BEP20CrowdsaleFilterer) ParsePaused(log types.Log) (*BEP20CrowdsalePaused, error) {
	event := new(BEP20CrowdsalePaused)
	if err := _BEP20Crowdsale.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BEP20CrowdsaleTokensPurchasedIterator is returned from FilterTokensPurchased and is used to iterate over the raw logs and unpacked data for TokensPurchased events raised by the BEP20Crowdsale contract.
type BEP20CrowdsaleTokensPurchasedIterator struct {
	Event *BEP20CrowdsaleTokensPurchased // Event containing the contract specifics and raw log

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
func (it *BEP20CrowdsaleTokensPurchasedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BEP20CrowdsaleTokensPurchased)
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
		it.Event = new(BEP20CrowdsaleTokensPurchased)
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
func (it *BEP20CrowdsaleTokensPurchasedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BEP20CrowdsaleTokensPurchasedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BEP20CrowdsaleTokensPurchased represents a TokensPurchased event raised by the BEP20Crowdsale contract.
type BEP20CrowdsaleTokensPurchased struct {
	Purchaser   common.Address
	Beneficiary common.Address
	Value       *big.Int
	Amount      *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterTokensPurchased is a free log retrieval operation binding the contract event 0x6faf93231a456e552dbc9961f58d9713ee4f2e69d15f1975b050ef0911053a7b.
//
// Solidity: event TokensPurchased(address indexed purchaser, address indexed beneficiary, uint256 value, uint256 amount)
func (_BEP20Crowdsale *BEP20CrowdsaleFilterer) FilterTokensPurchased(opts *bind.FilterOpts, purchaser []common.Address, beneficiary []common.Address) (*BEP20CrowdsaleTokensPurchasedIterator, error) {

	var purchaserRule []interface{}
	for _, purchaserItem := range purchaser {
		purchaserRule = append(purchaserRule, purchaserItem)
	}
	var beneficiaryRule []interface{}
	for _, beneficiaryItem := range beneficiary {
		beneficiaryRule = append(beneficiaryRule, beneficiaryItem)
	}

	logs, sub, err := _BEP20Crowdsale.contract.FilterLogs(opts, "TokensPurchased", purchaserRule, beneficiaryRule)
	if err != nil {
		return nil, err
	}
	return &BEP20CrowdsaleTokensPurchasedIterator{contract: _BEP20Crowdsale.contract, event: "TokensPurchased", logs: logs, sub: sub}, nil
}

// WatchTokensPurchased is a free log subscription operation binding the contract event 0x6faf93231a456e552dbc9961f58d9713ee4f2e69d15f1975b050ef0911053a7b.
//
// Solidity: event TokensPurchased(address indexed purchaser, address indexed beneficiary, uint256 value, uint256 amount)
func (_BEP20Crowdsale *BEP20CrowdsaleFilterer) WatchTokensPurchased(opts *bind.WatchOpts, sink chan<- *BEP20CrowdsaleTokensPurchased, purchaser []common.Address, beneficiary []common.Address) (event.Subscription, error) {

	var purchaserRule []interface{}
	for _, purchaserItem := range purchaser {
		purchaserRule = append(purchaserRule, purchaserItem)
	}
	var beneficiaryRule []interface{}
	for _, beneficiaryItem := range beneficiary {
		beneficiaryRule = append(beneficiaryRule, beneficiaryItem)
	}

	logs, sub, err := _BEP20Crowdsale.contract.WatchLogs(opts, "TokensPurchased", purchaserRule, beneficiaryRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BEP20CrowdsaleTokensPurchased)
				if err := _BEP20Crowdsale.contract.UnpackLog(event, "TokensPurchased", log); err != nil {
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

// ParseTokensPurchased is a log parse operation binding the contract event 0x6faf93231a456e552dbc9961f58d9713ee4f2e69d15f1975b050ef0911053a7b.
//
// Solidity: event TokensPurchased(address indexed purchaser, address indexed beneficiary, uint256 value, uint256 amount)
func (_BEP20Crowdsale *BEP20CrowdsaleFilterer) ParseTokensPurchased(log types.Log) (*BEP20CrowdsaleTokensPurchased, error) {
	event := new(BEP20CrowdsaleTokensPurchased)
	if err := _BEP20Crowdsale.contract.UnpackLog(event, "TokensPurchased", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BEP20CrowdsaleUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the BEP20Crowdsale contract.
type BEP20CrowdsaleUnpausedIterator struct {
	Event *BEP20CrowdsaleUnpaused // Event containing the contract specifics and raw log

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
func (it *BEP20CrowdsaleUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BEP20CrowdsaleUnpaused)
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
		it.Event = new(BEP20CrowdsaleUnpaused)
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
func (it *BEP20CrowdsaleUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BEP20CrowdsaleUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BEP20CrowdsaleUnpaused represents a Unpaused event raised by the BEP20Crowdsale contract.
type BEP20CrowdsaleUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_BEP20Crowdsale *BEP20CrowdsaleFilterer) FilterUnpaused(opts *bind.FilterOpts) (*BEP20CrowdsaleUnpausedIterator, error) {

	logs, sub, err := _BEP20Crowdsale.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &BEP20CrowdsaleUnpausedIterator{contract: _BEP20Crowdsale.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_BEP20Crowdsale *BEP20CrowdsaleFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *BEP20CrowdsaleUnpaused) (event.Subscription, error) {

	logs, sub, err := _BEP20Crowdsale.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BEP20CrowdsaleUnpaused)
				if err := _BEP20Crowdsale.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
func (_BEP20Crowdsale *BEP20CrowdsaleFilterer) ParseUnpaused(log types.Log) (*BEP20CrowdsaleUnpaused, error) {
	event := new(BEP20CrowdsaleUnpaused)
	if err := _BEP20Crowdsale.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
