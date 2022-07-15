// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package BEP20KFIVECrowdsale

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

// BEP20KFIVECrowdsaleMetaData contains all meta data concerning the BEP20KFIVECrowdsale contract.
var BEP20KFIVECrowdsaleMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"rate\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"wallet\",\"type\":\"address\"},{\"internalType\":\"contractIBEP20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"contractIBEP20\",\"name\":\"acceptToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"cap\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"openingTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"closingTime\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"prevClosingTime\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newClosingTime\",\"type\":\"uint256\"}],\"name\":\"TimedCrowdsaleExtended\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"purchaser\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"beneficiary\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"TokensPurchased\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"beneficiary\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"buyTokens\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"cap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"capReached\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"newReceiver\",\"type\":\"address\"}],\"name\":\"changeWallet\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"closingTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"_to\",\"type\":\"address\"}],\"name\":\"destroySmartContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"hasClosed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isOpen\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"openingTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token\",\"outputs\":[{\"internalType\":\"contractIBEP20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"wallet\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"weiRaised\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60806040523480156200001157600080fd5b5060405162002761380380620027618339818101604052810190620000379190620004f4565b818184898989896200005e62000052620003e360201b60201c565b620003eb60201b60201c565b600180819055506000600260006101000a81548160ff02191690831515021790555060008411620000c6576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401620000bd906200075c565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff16141562000139576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016200013090620006d4565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff161415620001ac576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401620001a39062000718565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614156200021f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040162000216906200077e565b60405180910390fd5b8360058190555082600460006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555081600260016101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555080600360006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550505050506000811162000333576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016200032a906200073a565b60405180910390fd5b80600781905550504282101562000381576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016200037890620006b2565b60405180910390fd5b818111620003c6576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401620003bd90620006f6565b60405180910390fd5b8160088190555080600981905550505050505050505050620009d0565b600033905090565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050816000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b600081519050620004c08162000982565b92915050565b600081519050620004d7816200099c565b92915050565b600081519050620004ee81620009b6565b92915050565b600080600080600080600060e0888a0312156200051057600080fd5b6000620005208a828b01620004dd565b9750506020620005338a828b01620004af565b9650506040620005468a828b01620004c6565b9550506060620005598a828b01620004c6565b94505060806200056c8a828b01620004dd565b93505060a06200057f8a828b01620004dd565b92505060c0620005928a828b01620004dd565b91505092959891949750929550565b6000620005b0603383620007a0565b9150620005bd8262000817565b604082019050919050565b6000620005d7600f83620007a0565b9150620005e48262000866565b602082019050919050565b6000620005fe603783620007a0565b91506200060b826200088f565b604082019050919050565b600062000625600d83620007a0565b91506200063282620008de565b602082019050919050565b60006200064c601983620007a0565b9150620006598262000907565b602082019050919050565b600062000673600c83620007a0565b9150620006808262000930565b602082019050919050565b60006200069a601483620007a0565b9150620006a78262000959565b602082019050919050565b60006020820190508181036000830152620006cd81620005a1565b9050919050565b60006020820190508181036000830152620006ef81620005c8565b9050919050565b600060208201905081810360008301526200071181620005ef565b9050919050565b60006020820190508181036000830152620007338162000616565b9050919050565b6000602082019050818103600083015262000755816200063d565b9050919050565b60006020820190508181036000830152620007778162000664565b9050919050565b6000602082019050818103600083015262000799816200068b565b9050919050565b600082825260208201905092915050565b6000620007be82620007ed565b9050919050565b6000620007d282620007ed565b9050919050565b6000620007e682620007b1565b9050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b7f54696d656443726f776473616c653a206f70656e696e672074696d652069732060008201527f6265666f72652063757272656e742074696d6500000000000000000000000000602082015250565b7f496e76616c696420616464726573730000000000000000000000000000000000600082015250565b7f54696d656443726f776473616c653a206f70656e696e672074696d652069732060008201527f6e6f74206265666f726520636c6f73696e672074696d65000000000000000000602082015250565b7f496e76616c696420746f6b656e00000000000000000000000000000000000000600082015250565b7f43617070656443726f776473616c653a20636170206973203000000000000000600082015250565b7f496e76616c696420726174650000000000000000000000000000000000000000600082015250565b7f496e76616c69642061636365707420746f6b656e000000000000000000000000600082015250565b6200098d81620007c5565b81146200099957600080fd5b50565b620009a781620007d9565b8114620009b357600080fd5b50565b620009c1816200080d565b8114620009cd57600080fd5b50565b611d8180620009e06000396000f3fe6080604052600436106101145760003560e01c80634f935945116100a05780638da5cb5b116100645780638da5cb5b1461032657806398b9a2dc14610351578063b7a8807c1461037a578063f2fde38b146103a5578063fc0c546a146103ce57610114565b80634f93594514610277578063521eb273146102a25780635c975abb146102cd578063715018a6146102f85780638456cb591461030f57610114565b806339df43ff116100e757806339df43ff146101b65780633f4ba83a146101df5780634042b66f146101f657806347535d7b146102215780634b6753bc1461024c57610114565b80630752881a146101195780631515bc2b146101355780632c4e722e14610160578063355274ea1461018b575b600080fd5b610133600480360381019061012e919061146c565b6103f9565b005b34801561014157600080fd5b5061014a6105e8565b6040516101579190611771565b60405180910390f35b34801561016c57600080fd5b506101756105f4565b60405161018291906118e7565b60405180910390f35b34801561019757600080fd5b506101a06105fe565b6040516101ad91906118e7565b60405180910390f35b3480156101c257600080fd5b506101dd60048036038101906101d89190611443565b610608565b005b3480156101eb57600080fd5b506101f461095b565b005b34801561020257600080fd5b5061020b610a28565b60405161021891906118e7565b60405180910390f35b34801561022d57600080fd5b50610236610a32565b6040516102439190611771565b60405180910390f35b34801561025857600080fd5b50610261610a4d565b60405161026e91906118e7565b60405180910390f35b34801561028357600080fd5b5061028c610a57565b6040516102999190611771565b60405180910390f35b3480156102ae57600080fd5b506102b7610a6b565b6040516102c491906116cd565b60405180910390f35b3480156102d957600080fd5b506102e2610a95565b6040516102ef9190611771565b60405180910390f35b34801561030457600080fd5b5061030d610aac565b005b34801561031b57600080fd5b50610324610b34565b005b34801561033257600080fd5b5061033b610c02565b60405161034891906116b2565b60405180910390f35b34801561035d57600080fd5b5061037860048036038101906103739190611443565b610c2b565b005b34801561038657600080fd5b5061038f610ceb565b60405161039c91906118e7565b60405180910390f35b3480156103b157600080fd5b506103cc60048036038101906103c7919061141a565b610cf5565b005b3480156103da57600080fd5b506103e3610ded565b6040516103f0919061178c565b60405180910390f35b6002600154141561043f576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610436906118c7565b60405180910390fd5b600260018190555061044f610a95565b1561048f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161048690611847565b60405180910390fd5b670de0b6b3a76400008110156104da576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016104d190611827565b60405180910390fd5b60006104f7670de0b6b3a764000083610e1790919063ffffffff16565b90506105038382610e2d565b600061050e82610e3b565b90506105348461052f670de0b6b3a764000085610e5990919063ffffffff16565b610e6f565b61053e8483610f47565b61055382600654610f4b90919063ffffffff16565b6006819055506105638482610f61565b8373ffffffffffffffffffffffffffffffffffffffff16610582610f6f565b73ffffffffffffffffffffffffffffffffffffffff167f6faf93231a456e552dbc9961f58d9713ee4f2e69d15f1975b050ef0911053a7b84846040516105c9929190611902565b60405180910390a36105db8483610f77565b5050600180819055505050565b60006009544211905090565b6000600554905090565b6000600754905090565b610610610f6f565b73ffffffffffffffffffffffffffffffffffffffff1661062e610c02565b73ffffffffffffffffffffffffffffffffffffffff1614610684576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161067b90611887565b60405180910390fd5b6000600260019054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166370a08231306040518263ffffffff1660e01b81526004016106e191906116b2565b60206040518083038186803b1580156106f957600080fd5b505afa15801561070d573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061073191906114d1565b9050600260019054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663a9059cbb83836040518363ffffffff1660e01b81526004016107909291906116e8565b602060405180830381600087803b1580156107aa57600080fd5b505af11580156107be573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906107e291906114a8565b506000600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166370a08231306040518263ffffffff1660e01b815260040161084091906116b2565b60206040518083038186803b15801561085857600080fd5b505afa15801561086c573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061089091906114d1565b9050600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663a9059cbb84836040518363ffffffff1660e01b81526004016108ef9291906116e8565b602060405180830381600087803b15801561090957600080fd5b505af115801561091d573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061094191906114a8565b508273ffffffffffffffffffffffffffffffffffffffff16ff5b610963610a95565b6109a2576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610999906117a7565b60405180910390fd5b6109aa610f6f565b73ffffffffffffffffffffffffffffffffffffffff166109c8610c02565b73ffffffffffffffffffffffffffffffffffffffff1614610a1e576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610a1590611887565b60405180910390fd5b610a26610f7b565b565b6000600654905090565b60006008544210158015610a4857506009544211155b905090565b6000600954905090565b6000600754610a64610a28565b1015905090565b6000600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b6000600260009054906101000a900460ff16905090565b610ab4610f6f565b73ffffffffffffffffffffffffffffffffffffffff16610ad2610c02565b73ffffffffffffffffffffffffffffffffffffffff1614610b28576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610b1f90611887565b60405180910390fd5b610b32600061101d565b565b610b3c610a95565b15610b7c576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610b7390611847565b60405180910390fd5b610b84610f6f565b73ffffffffffffffffffffffffffffffffffffffff16610ba2610c02565b73ffffffffffffffffffffffffffffffffffffffff1614610bf8576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610bef90611887565b60405180910390fd5b610c006110e1565b565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b610c33610f6f565b73ffffffffffffffffffffffffffffffffffffffff16610c51610c02565b73ffffffffffffffffffffffffffffffffffffffff1614610ca7576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610c9e90611887565b60405180910390fd5b80600460006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b6000600854905090565b610cfd610f6f565b73ffffffffffffffffffffffffffffffffffffffff16610d1b610c02565b73ffffffffffffffffffffffffffffffffffffffff1614610d71576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610d6890611887565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff161415610de1576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610dd8906117c7565b60405180910390fd5b610dea8161101d565b50565b6000600260019054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b60008183610e259190611992565b905092915050565b610e378282611184565b5050565b6000610e5260055483610e5990919063ffffffff16565b9050919050565b60008183610e6791906119c3565b905092915050565b600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166323b872dd83600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16846040518463ffffffff1660e01b8152600401610ef093929190611711565b602060405180830381600087803b158015610f0a57600080fd5b505af1158015610f1e573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610f4291906114a8565b505050565b5050565b60008183610f59919061193c565b905092915050565b610f6b82826111d9565b5050565b600033905090565b5050565b610f83610a95565b610fc2576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610fb9906117a7565b60405180910390fd5b6000600260006101000a81548160ff0219169083151502179055507f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa611006610f6f565b60405161101391906116b2565b60405180910390a1565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050816000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b6110e9610a95565b15611129576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161112090611847565b60405180910390fd5b6001600260006101000a81548160ff0219169083151502179055507f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a25861116d610f6f565b60405161117a91906116b2565b60405180910390a1565b61118c610a32565b6111cb576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016111c2906117e7565b60405180910390fd5b6111d5828261128d565b5050565b600260019054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663a9059cbb83836040518363ffffffff1660e01b8152600401611236929190611748565b602060405180830381600087803b15801561125057600080fd5b505af1158015611264573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061128891906114a8565b505050565b61129782826112f9565b6007546112b4826112a6610a28565b610f4b90919063ffffffff16565b11156112f5576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016112ec90611807565b60405180910390fd5b5050565b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff161415611369576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161136090611867565b60405180910390fd5b60008114156113ad576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016113a4906118a7565b60405180910390fd5b5050565b6000813590506113c081611cef565b92915050565b6000813590506113d581611d06565b92915050565b6000815190506113ea81611d1d565b92915050565b6000813590506113ff81611d34565b92915050565b60008151905061141481611d34565b92915050565b60006020828403121561142c57600080fd5b600061143a848285016113b1565b91505092915050565b60006020828403121561145557600080fd5b6000611463848285016113c6565b91505092915050565b6000806040838503121561147f57600080fd5b600061148d858286016113b1565b925050602061149e858286016113f0565b9150509250929050565b6000602082840312156114ba57600080fd5b60006114c8848285016113db565b91505092915050565b6000602082840312156114e357600080fd5b60006114f184828501611405565b91505092915050565b61150381611a77565b82525050565b61151281611a2f565b82525050565b61152181611a1d565b82525050565b61153081611a41565b82525050565b61153f81611a89565b82525050565b600061155260148361192b565b915061155d82611b2f565b602082019050919050565b600061157560268361192b565b915061158082611b58565b604082019050919050565b600061159860188361192b565b91506115a382611ba7565b602082019050919050565b60006115bb601d8361192b565b91506115c682611bd0565b602082019050919050565b60006115de600e8361192b565b91506115e982611bf9565b602082019050919050565b600061160160108361192b565b915061160c82611c22565b602082019050919050565b600061162460138361192b565b915061162f82611c4b565b602082019050919050565b600061164760208361192b565b915061165282611c74565b602082019050919050565b600061166a60078361192b565b915061167582611c9d565b602082019050919050565b600061168d601f8361192b565b915061169882611cc6565b602082019050919050565b6116ac81611a6d565b82525050565b60006020820190506116c76000830184611518565b92915050565b60006020820190506116e26000830184611509565b92915050565b60006040820190506116fd60008301856114fa565b61170a60208301846116a3565b9392505050565b60006060820190506117266000830186611518565b61173360208301856114fa565b61174060408301846116a3565b949350505050565b600060408201905061175d6000830185611518565b61176a60208301846116a3565b9392505050565b60006020820190506117866000830184611527565b92915050565b60006020820190506117a16000830184611536565b92915050565b600060208201905081810360008301526117c081611545565b9050919050565b600060208201905081810360008301526117e081611568565b9050919050565b600060208201905081810360008301526118008161158b565b9050919050565b60006020820190508181036000830152611820816115ae565b9050919050565b60006020820190508181036000830152611840816115d1565b9050919050565b60006020820190508181036000830152611860816115f4565b9050919050565b6000602082019050818103600083015261188081611617565b9050919050565b600060208201905081810360008301526118a08161163a565b9050919050565b600060208201905081810360008301526118c08161165d565b9050919050565b600060208201905081810360008301526118e081611680565b9050919050565b60006020820190506118fc60008301846116a3565b92915050565b600060408201905061191760008301856116a3565b61192460208301846116a3565b9392505050565b600082825260208201905092915050565b600061194782611a6d565b915061195283611a6d565b9250827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0382111561198757611986611ad1565b5b828201905092915050565b600061199d82611a6d565b91506119a883611a6d565b9250826119b8576119b7611b00565b5b828204905092915050565b60006119ce82611a6d565b91506119d983611a6d565b9250817fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0483118215151615611a1257611a11611ad1565b5b828202905092915050565b6000611a2882611a4d565b9050919050565b6000611a3a82611a4d565b9050919050565b60008115159050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b6000611a8282611aad565b9050919050565b6000611a9482611a9b565b9050919050565b6000611aa682611a4d565b9050919050565b6000611ab882611abf565b9050919050565b6000611aca82611a4d565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b7f5061757361626c653a206e6f7420706175736564000000000000000000000000600082015250565b7f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160008201527f6464726573730000000000000000000000000000000000000000000000000000602082015250565b7f54696d656443726f776473616c653a206e6f74206f70656e0000000000000000600082015250565b7f43617070656443726f776473616c653a20636170206578636565646564000000600082015250565b7f496e76616c696420616d6f756e74000000000000000000000000000000000000600082015250565b7f5061757361626c653a2070617573656400000000000000000000000000000000600082015250565b7f496e76616c69642062656e656669636961727900000000000000000000000000600082015250565b7f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572600082015250565b7f496e76616c696400000000000000000000000000000000000000000000000000600082015250565b7f5265656e7472616e637947756172643a207265656e7472616e742063616c6c00600082015250565b611cf881611a1d565b8114611d0357600080fd5b50565b611d0f81611a2f565b8114611d1a57600080fd5b50565b611d2681611a41565b8114611d3157600080fd5b50565b611d3d81611a6d565b8114611d4857600080fd5b5056fea26469706673582212201aa630c214d5cc91f61865f24df831824ab5b9c44fbc098e1ace790c8f5554e364736f6c63430008040033",
}

// BEP20KFIVECrowdsaleABI is the input ABI used to generate the binding from.
// Deprecated: Use BEP20KFIVECrowdsaleMetaData.ABI instead.
var BEP20KFIVECrowdsaleABI = BEP20KFIVECrowdsaleMetaData.ABI

// BEP20KFIVECrowdsaleBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use BEP20KFIVECrowdsaleMetaData.Bin instead.
var BEP20KFIVECrowdsaleBin = BEP20KFIVECrowdsaleMetaData.Bin

// DeployBEP20KFIVECrowdsale deploys a new Ethereum contract, binding an instance of BEP20KFIVECrowdsale to it.
func DeployBEP20KFIVECrowdsale(auth *bind.TransactOpts, backend bind.ContractBackend, rate *big.Int, wallet common.Address, token common.Address, acceptToken common.Address, cap *big.Int, openingTime *big.Int, closingTime *big.Int) (common.Address, *types.Transaction, *BEP20KFIVECrowdsale, error) {
	parsed, err := BEP20KFIVECrowdsaleMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(BEP20KFIVECrowdsaleBin), backend, rate, wallet, token, acceptToken, cap, openingTime, closingTime)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &BEP20KFIVECrowdsale{BEP20KFIVECrowdsaleCaller: BEP20KFIVECrowdsaleCaller{contract: contract}, BEP20KFIVECrowdsaleTransactor: BEP20KFIVECrowdsaleTransactor{contract: contract}, BEP20KFIVECrowdsaleFilterer: BEP20KFIVECrowdsaleFilterer{contract: contract}}, nil
}

// BEP20KFIVECrowdsale is an auto generated Go binding around an Ethereum contract.
type BEP20KFIVECrowdsale struct {
	BEP20KFIVECrowdsaleCaller     // Read-only binding to the contract
	BEP20KFIVECrowdsaleTransactor // Write-only binding to the contract
	BEP20KFIVECrowdsaleFilterer   // Log filterer for contract events
}

// BEP20KFIVECrowdsaleCaller is an auto generated read-only Go binding around an Ethereum contract.
type BEP20KFIVECrowdsaleCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BEP20KFIVECrowdsaleTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BEP20KFIVECrowdsaleTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BEP20KFIVECrowdsaleFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BEP20KFIVECrowdsaleFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BEP20KFIVECrowdsaleSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BEP20KFIVECrowdsaleSession struct {
	Contract     *BEP20KFIVECrowdsale // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// BEP20KFIVECrowdsaleCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BEP20KFIVECrowdsaleCallerSession struct {
	Contract *BEP20KFIVECrowdsaleCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// BEP20KFIVECrowdsaleTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BEP20KFIVECrowdsaleTransactorSession struct {
	Contract     *BEP20KFIVECrowdsaleTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// BEP20KFIVECrowdsaleRaw is an auto generated low-level Go binding around an Ethereum contract.
type BEP20KFIVECrowdsaleRaw struct {
	Contract *BEP20KFIVECrowdsale // Generic contract binding to access the raw methods on
}

// BEP20KFIVECrowdsaleCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BEP20KFIVECrowdsaleCallerRaw struct {
	Contract *BEP20KFIVECrowdsaleCaller // Generic read-only contract binding to access the raw methods on
}

// BEP20KFIVECrowdsaleTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BEP20KFIVECrowdsaleTransactorRaw struct {
	Contract *BEP20KFIVECrowdsaleTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBEP20KFIVECrowdsale creates a new instance of BEP20KFIVECrowdsale, bound to a specific deployed contract.
func NewBEP20KFIVECrowdsale(address common.Address, backend bind.ContractBackend) (*BEP20KFIVECrowdsale, error) {
	contract, err := bindBEP20KFIVECrowdsale(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BEP20KFIVECrowdsale{BEP20KFIVECrowdsaleCaller: BEP20KFIVECrowdsaleCaller{contract: contract}, BEP20KFIVECrowdsaleTransactor: BEP20KFIVECrowdsaleTransactor{contract: contract}, BEP20KFIVECrowdsaleFilterer: BEP20KFIVECrowdsaleFilterer{contract: contract}}, nil
}

// NewBEP20KFIVECrowdsaleCaller creates a new read-only instance of BEP20KFIVECrowdsale, bound to a specific deployed contract.
func NewBEP20KFIVECrowdsaleCaller(address common.Address, caller bind.ContractCaller) (*BEP20KFIVECrowdsaleCaller, error) {
	contract, err := bindBEP20KFIVECrowdsale(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BEP20KFIVECrowdsaleCaller{contract: contract}, nil
}

// NewBEP20KFIVECrowdsaleTransactor creates a new write-only instance of BEP20KFIVECrowdsale, bound to a specific deployed contract.
func NewBEP20KFIVECrowdsaleTransactor(address common.Address, transactor bind.ContractTransactor) (*BEP20KFIVECrowdsaleTransactor, error) {
	contract, err := bindBEP20KFIVECrowdsale(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BEP20KFIVECrowdsaleTransactor{contract: contract}, nil
}

// NewBEP20KFIVECrowdsaleFilterer creates a new log filterer instance of BEP20KFIVECrowdsale, bound to a specific deployed contract.
func NewBEP20KFIVECrowdsaleFilterer(address common.Address, filterer bind.ContractFilterer) (*BEP20KFIVECrowdsaleFilterer, error) {
	contract, err := bindBEP20KFIVECrowdsale(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BEP20KFIVECrowdsaleFilterer{contract: contract}, nil
}

// bindBEP20KFIVECrowdsale binds a generic wrapper to an already deployed contract.
func bindBEP20KFIVECrowdsale(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BEP20KFIVECrowdsaleABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BEP20KFIVECrowdsale *BEP20KFIVECrowdsaleRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BEP20KFIVECrowdsale.Contract.BEP20KFIVECrowdsaleCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BEP20KFIVECrowdsale *BEP20KFIVECrowdsaleRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BEP20KFIVECrowdsale.Contract.BEP20KFIVECrowdsaleTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BEP20KFIVECrowdsale *BEP20KFIVECrowdsaleRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BEP20KFIVECrowdsale.Contract.BEP20KFIVECrowdsaleTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BEP20KFIVECrowdsale *BEP20KFIVECrowdsaleCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BEP20KFIVECrowdsale.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BEP20KFIVECrowdsale *BEP20KFIVECrowdsaleTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BEP20KFIVECrowdsale.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BEP20KFIVECrowdsale *BEP20KFIVECrowdsaleTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BEP20KFIVECrowdsale.Contract.contract.Transact(opts, method, params...)
}

// Cap is a free data retrieval call binding the contract method 0x355274ea.
//
// Solidity: function cap() view returns(uint256)
func (_BEP20KFIVECrowdsale *BEP20KFIVECrowdsaleCaller) Cap(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BEP20KFIVECrowdsale.contract.Call(opts, &out, "cap")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Cap is a free data retrieval call binding the contract method 0x355274ea.
//
// Solidity: function cap() view returns(uint256)
func (_BEP20KFIVECrowdsale *BEP20KFIVECrowdsaleSession) Cap() (*big.Int, error) {
	return _BEP20KFIVECrowdsale.Contract.Cap(&_BEP20KFIVECrowdsale.CallOpts)
}

// Cap is a free data retrieval call binding the contract method 0x355274ea.
//
// Solidity: function cap() view returns(uint256)
func (_BEP20KFIVECrowdsale *BEP20KFIVECrowdsaleCallerSession) Cap() (*big.Int, error) {
	return _BEP20KFIVECrowdsale.Contract.Cap(&_BEP20KFIVECrowdsale.CallOpts)
}

// CapReached is a free data retrieval call binding the contract method 0x4f935945.
//
// Solidity: function capReached() view returns(bool)
func (_BEP20KFIVECrowdsale *BEP20KFIVECrowdsaleCaller) CapReached(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _BEP20KFIVECrowdsale.contract.Call(opts, &out, "capReached")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CapReached is a free data retrieval call binding the contract method 0x4f935945.
//
// Solidity: function capReached() view returns(bool)
func (_BEP20KFIVECrowdsale *BEP20KFIVECrowdsaleSession) CapReached() (bool, error) {
	return _BEP20KFIVECrowdsale.Contract.CapReached(&_BEP20KFIVECrowdsale.CallOpts)
}

// CapReached is a free data retrieval call binding the contract method 0x4f935945.
//
// Solidity: function capReached() view returns(bool)
func (_BEP20KFIVECrowdsale *BEP20KFIVECrowdsaleCallerSession) CapReached() (bool, error) {
	return _BEP20KFIVECrowdsale.Contract.CapReached(&_BEP20KFIVECrowdsale.CallOpts)
}

// ClosingTime is a free data retrieval call binding the contract method 0x4b6753bc.
//
// Solidity: function closingTime() view returns(uint256)
func (_BEP20KFIVECrowdsale *BEP20KFIVECrowdsaleCaller) ClosingTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BEP20KFIVECrowdsale.contract.Call(opts, &out, "closingTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ClosingTime is a free data retrieval call binding the contract method 0x4b6753bc.
//
// Solidity: function closingTime() view returns(uint256)
func (_BEP20KFIVECrowdsale *BEP20KFIVECrowdsaleSession) ClosingTime() (*big.Int, error) {
	return _BEP20KFIVECrowdsale.Contract.ClosingTime(&_BEP20KFIVECrowdsale.CallOpts)
}

// ClosingTime is a free data retrieval call binding the contract method 0x4b6753bc.
//
// Solidity: function closingTime() view returns(uint256)
func (_BEP20KFIVECrowdsale *BEP20KFIVECrowdsaleCallerSession) ClosingTime() (*big.Int, error) {
	return _BEP20KFIVECrowdsale.Contract.ClosingTime(&_BEP20KFIVECrowdsale.CallOpts)
}

// HasClosed is a free data retrieval call binding the contract method 0x1515bc2b.
//
// Solidity: function hasClosed() view returns(bool)
func (_BEP20KFIVECrowdsale *BEP20KFIVECrowdsaleCaller) HasClosed(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _BEP20KFIVECrowdsale.contract.Call(opts, &out, "hasClosed")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasClosed is a free data retrieval call binding the contract method 0x1515bc2b.
//
// Solidity: function hasClosed() view returns(bool)
func (_BEP20KFIVECrowdsale *BEP20KFIVECrowdsaleSession) HasClosed() (bool, error) {
	return _BEP20KFIVECrowdsale.Contract.HasClosed(&_BEP20KFIVECrowdsale.CallOpts)
}

// HasClosed is a free data retrieval call binding the contract method 0x1515bc2b.
//
// Solidity: function hasClosed() view returns(bool)
func (_BEP20KFIVECrowdsale *BEP20KFIVECrowdsaleCallerSession) HasClosed() (bool, error) {
	return _BEP20KFIVECrowdsale.Contract.HasClosed(&_BEP20KFIVECrowdsale.CallOpts)
}

// IsOpen is a free data retrieval call binding the contract method 0x47535d7b.
//
// Solidity: function isOpen() view returns(bool)
func (_BEP20KFIVECrowdsale *BEP20KFIVECrowdsaleCaller) IsOpen(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _BEP20KFIVECrowdsale.contract.Call(opts, &out, "isOpen")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsOpen is a free data retrieval call binding the contract method 0x47535d7b.
//
// Solidity: function isOpen() view returns(bool)
func (_BEP20KFIVECrowdsale *BEP20KFIVECrowdsaleSession) IsOpen() (bool, error) {
	return _BEP20KFIVECrowdsale.Contract.IsOpen(&_BEP20KFIVECrowdsale.CallOpts)
}

// IsOpen is a free data retrieval call binding the contract method 0x47535d7b.
//
// Solidity: function isOpen() view returns(bool)
func (_BEP20KFIVECrowdsale *BEP20KFIVECrowdsaleCallerSession) IsOpen() (bool, error) {
	return _BEP20KFIVECrowdsale.Contract.IsOpen(&_BEP20KFIVECrowdsale.CallOpts)
}

// OpeningTime is a free data retrieval call binding the contract method 0xb7a8807c.
//
// Solidity: function openingTime() view returns(uint256)
func (_BEP20KFIVECrowdsale *BEP20KFIVECrowdsaleCaller) OpeningTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BEP20KFIVECrowdsale.contract.Call(opts, &out, "openingTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// OpeningTime is a free data retrieval call binding the contract method 0xb7a8807c.
//
// Solidity: function openingTime() view returns(uint256)
func (_BEP20KFIVECrowdsale *BEP20KFIVECrowdsaleSession) OpeningTime() (*big.Int, error) {
	return _BEP20KFIVECrowdsale.Contract.OpeningTime(&_BEP20KFIVECrowdsale.CallOpts)
}

// OpeningTime is a free data retrieval call binding the contract method 0xb7a8807c.
//
// Solidity: function openingTime() view returns(uint256)
func (_BEP20KFIVECrowdsale *BEP20KFIVECrowdsaleCallerSession) OpeningTime() (*big.Int, error) {
	return _BEP20KFIVECrowdsale.Contract.OpeningTime(&_BEP20KFIVECrowdsale.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BEP20KFIVECrowdsale *BEP20KFIVECrowdsaleCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BEP20KFIVECrowdsale.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BEP20KFIVECrowdsale *BEP20KFIVECrowdsaleSession) Owner() (common.Address, error) {
	return _BEP20KFIVECrowdsale.Contract.Owner(&_BEP20KFIVECrowdsale.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BEP20KFIVECrowdsale *BEP20KFIVECrowdsaleCallerSession) Owner() (common.Address, error) {
	return _BEP20KFIVECrowdsale.Contract.Owner(&_BEP20KFIVECrowdsale.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_BEP20KFIVECrowdsale *BEP20KFIVECrowdsaleCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _BEP20KFIVECrowdsale.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_BEP20KFIVECrowdsale *BEP20KFIVECrowdsaleSession) Paused() (bool, error) {
	return _BEP20KFIVECrowdsale.Contract.Paused(&_BEP20KFIVECrowdsale.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_BEP20KFIVECrowdsale *BEP20KFIVECrowdsaleCallerSession) Paused() (bool, error) {
	return _BEP20KFIVECrowdsale.Contract.Paused(&_BEP20KFIVECrowdsale.CallOpts)
}

// Rate is a free data retrieval call binding the contract method 0x2c4e722e.
//
// Solidity: function rate() view returns(uint256)
func (_BEP20KFIVECrowdsale *BEP20KFIVECrowdsaleCaller) Rate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BEP20KFIVECrowdsale.contract.Call(opts, &out, "rate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Rate is a free data retrieval call binding the contract method 0x2c4e722e.
//
// Solidity: function rate() view returns(uint256)
func (_BEP20KFIVECrowdsale *BEP20KFIVECrowdsaleSession) Rate() (*big.Int, error) {
	return _BEP20KFIVECrowdsale.Contract.Rate(&_BEP20KFIVECrowdsale.CallOpts)
}

// Rate is a free data retrieval call binding the contract method 0x2c4e722e.
//
// Solidity: function rate() view returns(uint256)
func (_BEP20KFIVECrowdsale *BEP20KFIVECrowdsaleCallerSession) Rate() (*big.Int, error) {
	return _BEP20KFIVECrowdsale.Contract.Rate(&_BEP20KFIVECrowdsale.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_BEP20KFIVECrowdsale *BEP20KFIVECrowdsaleCaller) Token(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BEP20KFIVECrowdsale.contract.Call(opts, &out, "token")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_BEP20KFIVECrowdsale *BEP20KFIVECrowdsaleSession) Token() (common.Address, error) {
	return _BEP20KFIVECrowdsale.Contract.Token(&_BEP20KFIVECrowdsale.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_BEP20KFIVECrowdsale *BEP20KFIVECrowdsaleCallerSession) Token() (common.Address, error) {
	return _BEP20KFIVECrowdsale.Contract.Token(&_BEP20KFIVECrowdsale.CallOpts)
}

// Wallet is a free data retrieval call binding the contract method 0x521eb273.
//
// Solidity: function wallet() view returns(address)
func (_BEP20KFIVECrowdsale *BEP20KFIVECrowdsaleCaller) Wallet(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BEP20KFIVECrowdsale.contract.Call(opts, &out, "wallet")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Wallet is a free data retrieval call binding the contract method 0x521eb273.
//
// Solidity: function wallet() view returns(address)
func (_BEP20KFIVECrowdsale *BEP20KFIVECrowdsaleSession) Wallet() (common.Address, error) {
	return _BEP20KFIVECrowdsale.Contract.Wallet(&_BEP20KFIVECrowdsale.CallOpts)
}

// Wallet is a free data retrieval call binding the contract method 0x521eb273.
//
// Solidity: function wallet() view returns(address)
func (_BEP20KFIVECrowdsale *BEP20KFIVECrowdsaleCallerSession) Wallet() (common.Address, error) {
	return _BEP20KFIVECrowdsale.Contract.Wallet(&_BEP20KFIVECrowdsale.CallOpts)
}

// WeiRaised is a free data retrieval call binding the contract method 0x4042b66f.
//
// Solidity: function weiRaised() view returns(uint256)
func (_BEP20KFIVECrowdsale *BEP20KFIVECrowdsaleCaller) WeiRaised(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BEP20KFIVECrowdsale.contract.Call(opts, &out, "weiRaised")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WeiRaised is a free data retrieval call binding the contract method 0x4042b66f.
//
// Solidity: function weiRaised() view returns(uint256)
func (_BEP20KFIVECrowdsale *BEP20KFIVECrowdsaleSession) WeiRaised() (*big.Int, error) {
	return _BEP20KFIVECrowdsale.Contract.WeiRaised(&_BEP20KFIVECrowdsale.CallOpts)
}

// WeiRaised is a free data retrieval call binding the contract method 0x4042b66f.
//
// Solidity: function weiRaised() view returns(uint256)
func (_BEP20KFIVECrowdsale *BEP20KFIVECrowdsaleCallerSession) WeiRaised() (*big.Int, error) {
	return _BEP20KFIVECrowdsale.Contract.WeiRaised(&_BEP20KFIVECrowdsale.CallOpts)
}

// BuyTokens is a paid mutator transaction binding the contract method 0x0752881a.
//
// Solidity: function buyTokens(address beneficiary, uint256 value) payable returns()
func (_BEP20KFIVECrowdsale *BEP20KFIVECrowdsaleTransactor) BuyTokens(opts *bind.TransactOpts, beneficiary common.Address, value *big.Int) (*types.Transaction, error) {
	return _BEP20KFIVECrowdsale.contract.Transact(opts, "buyTokens", beneficiary, value)
}

// BuyTokens is a paid mutator transaction binding the contract method 0x0752881a.
//
// Solidity: function buyTokens(address beneficiary, uint256 value) payable returns()
func (_BEP20KFIVECrowdsale *BEP20KFIVECrowdsaleSession) BuyTokens(beneficiary common.Address, value *big.Int) (*types.Transaction, error) {
	return _BEP20KFIVECrowdsale.Contract.BuyTokens(&_BEP20KFIVECrowdsale.TransactOpts, beneficiary, value)
}

// BuyTokens is a paid mutator transaction binding the contract method 0x0752881a.
//
// Solidity: function buyTokens(address beneficiary, uint256 value) payable returns()
func (_BEP20KFIVECrowdsale *BEP20KFIVECrowdsaleTransactorSession) BuyTokens(beneficiary common.Address, value *big.Int) (*types.Transaction, error) {
	return _BEP20KFIVECrowdsale.Contract.BuyTokens(&_BEP20KFIVECrowdsale.TransactOpts, beneficiary, value)
}

// ChangeWallet is a paid mutator transaction binding the contract method 0x98b9a2dc.
//
// Solidity: function changeWallet(address newReceiver) returns()
func (_BEP20KFIVECrowdsale *BEP20KFIVECrowdsaleTransactor) ChangeWallet(opts *bind.TransactOpts, newReceiver common.Address) (*types.Transaction, error) {
	return _BEP20KFIVECrowdsale.contract.Transact(opts, "changeWallet", newReceiver)
}

// ChangeWallet is a paid mutator transaction binding the contract method 0x98b9a2dc.
//
// Solidity: function changeWallet(address newReceiver) returns()
func (_BEP20KFIVECrowdsale *BEP20KFIVECrowdsaleSession) ChangeWallet(newReceiver common.Address) (*types.Transaction, error) {
	return _BEP20KFIVECrowdsale.Contract.ChangeWallet(&_BEP20KFIVECrowdsale.TransactOpts, newReceiver)
}

// ChangeWallet is a paid mutator transaction binding the contract method 0x98b9a2dc.
//
// Solidity: function changeWallet(address newReceiver) returns()
func (_BEP20KFIVECrowdsale *BEP20KFIVECrowdsaleTransactorSession) ChangeWallet(newReceiver common.Address) (*types.Transaction, error) {
	return _BEP20KFIVECrowdsale.Contract.ChangeWallet(&_BEP20KFIVECrowdsale.TransactOpts, newReceiver)
}

// DestroySmartContract is a paid mutator transaction binding the contract method 0x39df43ff.
//
// Solidity: function destroySmartContract(address _to) returns()
func (_BEP20KFIVECrowdsale *BEP20KFIVECrowdsaleTransactor) DestroySmartContract(opts *bind.TransactOpts, _to common.Address) (*types.Transaction, error) {
	return _BEP20KFIVECrowdsale.contract.Transact(opts, "destroySmartContract", _to)
}

// DestroySmartContract is a paid mutator transaction binding the contract method 0x39df43ff.
//
// Solidity: function destroySmartContract(address _to) returns()
func (_BEP20KFIVECrowdsale *BEP20KFIVECrowdsaleSession) DestroySmartContract(_to common.Address) (*types.Transaction, error) {
	return _BEP20KFIVECrowdsale.Contract.DestroySmartContract(&_BEP20KFIVECrowdsale.TransactOpts, _to)
}

// DestroySmartContract is a paid mutator transaction binding the contract method 0x39df43ff.
//
// Solidity: function destroySmartContract(address _to) returns()
func (_BEP20KFIVECrowdsale *BEP20KFIVECrowdsaleTransactorSession) DestroySmartContract(_to common.Address) (*types.Transaction, error) {
	return _BEP20KFIVECrowdsale.Contract.DestroySmartContract(&_BEP20KFIVECrowdsale.TransactOpts, _to)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_BEP20KFIVECrowdsale *BEP20KFIVECrowdsaleTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BEP20KFIVECrowdsale.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_BEP20KFIVECrowdsale *BEP20KFIVECrowdsaleSession) Pause() (*types.Transaction, error) {
	return _BEP20KFIVECrowdsale.Contract.Pause(&_BEP20KFIVECrowdsale.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_BEP20KFIVECrowdsale *BEP20KFIVECrowdsaleTransactorSession) Pause() (*types.Transaction, error) {
	return _BEP20KFIVECrowdsale.Contract.Pause(&_BEP20KFIVECrowdsale.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_BEP20KFIVECrowdsale *BEP20KFIVECrowdsaleTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BEP20KFIVECrowdsale.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_BEP20KFIVECrowdsale *BEP20KFIVECrowdsaleSession) RenounceOwnership() (*types.Transaction, error) {
	return _BEP20KFIVECrowdsale.Contract.RenounceOwnership(&_BEP20KFIVECrowdsale.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_BEP20KFIVECrowdsale *BEP20KFIVECrowdsaleTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _BEP20KFIVECrowdsale.Contract.RenounceOwnership(&_BEP20KFIVECrowdsale.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_BEP20KFIVECrowdsale *BEP20KFIVECrowdsaleTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _BEP20KFIVECrowdsale.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_BEP20KFIVECrowdsale *BEP20KFIVECrowdsaleSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _BEP20KFIVECrowdsale.Contract.TransferOwnership(&_BEP20KFIVECrowdsale.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_BEP20KFIVECrowdsale *BEP20KFIVECrowdsaleTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _BEP20KFIVECrowdsale.Contract.TransferOwnership(&_BEP20KFIVECrowdsale.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_BEP20KFIVECrowdsale *BEP20KFIVECrowdsaleTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BEP20KFIVECrowdsale.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_BEP20KFIVECrowdsale *BEP20KFIVECrowdsaleSession) Unpause() (*types.Transaction, error) {
	return _BEP20KFIVECrowdsale.Contract.Unpause(&_BEP20KFIVECrowdsale.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_BEP20KFIVECrowdsale *BEP20KFIVECrowdsaleTransactorSession) Unpause() (*types.Transaction, error) {
	return _BEP20KFIVECrowdsale.Contract.Unpause(&_BEP20KFIVECrowdsale.TransactOpts)
}

// BEP20KFIVECrowdsaleOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the BEP20KFIVECrowdsale contract.
type BEP20KFIVECrowdsaleOwnershipTransferredIterator struct {
	Event *BEP20KFIVECrowdsaleOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *BEP20KFIVECrowdsaleOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BEP20KFIVECrowdsaleOwnershipTransferred)
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
		it.Event = new(BEP20KFIVECrowdsaleOwnershipTransferred)
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
func (it *BEP20KFIVECrowdsaleOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BEP20KFIVECrowdsaleOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BEP20KFIVECrowdsaleOwnershipTransferred represents a OwnershipTransferred event raised by the BEP20KFIVECrowdsale contract.
type BEP20KFIVECrowdsaleOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_BEP20KFIVECrowdsale *BEP20KFIVECrowdsaleFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*BEP20KFIVECrowdsaleOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _BEP20KFIVECrowdsale.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &BEP20KFIVECrowdsaleOwnershipTransferredIterator{contract: _BEP20KFIVECrowdsale.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_BEP20KFIVECrowdsale *BEP20KFIVECrowdsaleFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *BEP20KFIVECrowdsaleOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _BEP20KFIVECrowdsale.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BEP20KFIVECrowdsaleOwnershipTransferred)
				if err := _BEP20KFIVECrowdsale.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_BEP20KFIVECrowdsale *BEP20KFIVECrowdsaleFilterer) ParseOwnershipTransferred(log types.Log) (*BEP20KFIVECrowdsaleOwnershipTransferred, error) {
	event := new(BEP20KFIVECrowdsaleOwnershipTransferred)
	if err := _BEP20KFIVECrowdsale.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BEP20KFIVECrowdsalePausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the BEP20KFIVECrowdsale contract.
type BEP20KFIVECrowdsalePausedIterator struct {
	Event *BEP20KFIVECrowdsalePaused // Event containing the contract specifics and raw log

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
func (it *BEP20KFIVECrowdsalePausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BEP20KFIVECrowdsalePaused)
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
		it.Event = new(BEP20KFIVECrowdsalePaused)
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
func (it *BEP20KFIVECrowdsalePausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BEP20KFIVECrowdsalePausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BEP20KFIVECrowdsalePaused represents a Paused event raised by the BEP20KFIVECrowdsale contract.
type BEP20KFIVECrowdsalePaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_BEP20KFIVECrowdsale *BEP20KFIVECrowdsaleFilterer) FilterPaused(opts *bind.FilterOpts) (*BEP20KFIVECrowdsalePausedIterator, error) {

	logs, sub, err := _BEP20KFIVECrowdsale.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &BEP20KFIVECrowdsalePausedIterator{contract: _BEP20KFIVECrowdsale.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_BEP20KFIVECrowdsale *BEP20KFIVECrowdsaleFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *BEP20KFIVECrowdsalePaused) (event.Subscription, error) {

	logs, sub, err := _BEP20KFIVECrowdsale.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BEP20KFIVECrowdsalePaused)
				if err := _BEP20KFIVECrowdsale.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_BEP20KFIVECrowdsale *BEP20KFIVECrowdsaleFilterer) ParsePaused(log types.Log) (*BEP20KFIVECrowdsalePaused, error) {
	event := new(BEP20KFIVECrowdsalePaused)
	if err := _BEP20KFIVECrowdsale.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BEP20KFIVECrowdsaleTimedCrowdsaleExtendedIterator is returned from FilterTimedCrowdsaleExtended and is used to iterate over the raw logs and unpacked data for TimedCrowdsaleExtended events raised by the BEP20KFIVECrowdsale contract.
type BEP20KFIVECrowdsaleTimedCrowdsaleExtendedIterator struct {
	Event *BEP20KFIVECrowdsaleTimedCrowdsaleExtended // Event containing the contract specifics and raw log

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
func (it *BEP20KFIVECrowdsaleTimedCrowdsaleExtendedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BEP20KFIVECrowdsaleTimedCrowdsaleExtended)
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
		it.Event = new(BEP20KFIVECrowdsaleTimedCrowdsaleExtended)
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
func (it *BEP20KFIVECrowdsaleTimedCrowdsaleExtendedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BEP20KFIVECrowdsaleTimedCrowdsaleExtendedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BEP20KFIVECrowdsaleTimedCrowdsaleExtended represents a TimedCrowdsaleExtended event raised by the BEP20KFIVECrowdsale contract.
type BEP20KFIVECrowdsaleTimedCrowdsaleExtended struct {
	PrevClosingTime *big.Int
	NewClosingTime  *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterTimedCrowdsaleExtended is a free log retrieval operation binding the contract event 0x46711e222f558a07afd26e5e71b48ecb0a8b2cdcd40faeb1323e05e2c76a2f32.
//
// Solidity: event TimedCrowdsaleExtended(uint256 prevClosingTime, uint256 newClosingTime)
func (_BEP20KFIVECrowdsale *BEP20KFIVECrowdsaleFilterer) FilterTimedCrowdsaleExtended(opts *bind.FilterOpts) (*BEP20KFIVECrowdsaleTimedCrowdsaleExtendedIterator, error) {

	logs, sub, err := _BEP20KFIVECrowdsale.contract.FilterLogs(opts, "TimedCrowdsaleExtended")
	if err != nil {
		return nil, err
	}
	return &BEP20KFIVECrowdsaleTimedCrowdsaleExtendedIterator{contract: _BEP20KFIVECrowdsale.contract, event: "TimedCrowdsaleExtended", logs: logs, sub: sub}, nil
}

// WatchTimedCrowdsaleExtended is a free log subscription operation binding the contract event 0x46711e222f558a07afd26e5e71b48ecb0a8b2cdcd40faeb1323e05e2c76a2f32.
//
// Solidity: event TimedCrowdsaleExtended(uint256 prevClosingTime, uint256 newClosingTime)
func (_BEP20KFIVECrowdsale *BEP20KFIVECrowdsaleFilterer) WatchTimedCrowdsaleExtended(opts *bind.WatchOpts, sink chan<- *BEP20KFIVECrowdsaleTimedCrowdsaleExtended) (event.Subscription, error) {

	logs, sub, err := _BEP20KFIVECrowdsale.contract.WatchLogs(opts, "TimedCrowdsaleExtended")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BEP20KFIVECrowdsaleTimedCrowdsaleExtended)
				if err := _BEP20KFIVECrowdsale.contract.UnpackLog(event, "TimedCrowdsaleExtended", log); err != nil {
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

// ParseTimedCrowdsaleExtended is a log parse operation binding the contract event 0x46711e222f558a07afd26e5e71b48ecb0a8b2cdcd40faeb1323e05e2c76a2f32.
//
// Solidity: event TimedCrowdsaleExtended(uint256 prevClosingTime, uint256 newClosingTime)
func (_BEP20KFIVECrowdsale *BEP20KFIVECrowdsaleFilterer) ParseTimedCrowdsaleExtended(log types.Log) (*BEP20KFIVECrowdsaleTimedCrowdsaleExtended, error) {
	event := new(BEP20KFIVECrowdsaleTimedCrowdsaleExtended)
	if err := _BEP20KFIVECrowdsale.contract.UnpackLog(event, "TimedCrowdsaleExtended", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BEP20KFIVECrowdsaleTokensPurchasedIterator is returned from FilterTokensPurchased and is used to iterate over the raw logs and unpacked data for TokensPurchased events raised by the BEP20KFIVECrowdsale contract.
type BEP20KFIVECrowdsaleTokensPurchasedIterator struct {
	Event *BEP20KFIVECrowdsaleTokensPurchased // Event containing the contract specifics and raw log

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
func (it *BEP20KFIVECrowdsaleTokensPurchasedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BEP20KFIVECrowdsaleTokensPurchased)
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
		it.Event = new(BEP20KFIVECrowdsaleTokensPurchased)
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
func (it *BEP20KFIVECrowdsaleTokensPurchasedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BEP20KFIVECrowdsaleTokensPurchasedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BEP20KFIVECrowdsaleTokensPurchased represents a TokensPurchased event raised by the BEP20KFIVECrowdsale contract.
type BEP20KFIVECrowdsaleTokensPurchased struct {
	Purchaser   common.Address
	Beneficiary common.Address
	Value       *big.Int
	Amount      *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterTokensPurchased is a free log retrieval operation binding the contract event 0x6faf93231a456e552dbc9961f58d9713ee4f2e69d15f1975b050ef0911053a7b.
//
// Solidity: event TokensPurchased(address indexed purchaser, address indexed beneficiary, uint256 value, uint256 amount)
func (_BEP20KFIVECrowdsale *BEP20KFIVECrowdsaleFilterer) FilterTokensPurchased(opts *bind.FilterOpts, purchaser []common.Address, beneficiary []common.Address) (*BEP20KFIVECrowdsaleTokensPurchasedIterator, error) {

	var purchaserRule []interface{}
	for _, purchaserItem := range purchaser {
		purchaserRule = append(purchaserRule, purchaserItem)
	}
	var beneficiaryRule []interface{}
	for _, beneficiaryItem := range beneficiary {
		beneficiaryRule = append(beneficiaryRule, beneficiaryItem)
	}

	logs, sub, err := _BEP20KFIVECrowdsale.contract.FilterLogs(opts, "TokensPurchased", purchaserRule, beneficiaryRule)
	if err != nil {
		return nil, err
	}
	return &BEP20KFIVECrowdsaleTokensPurchasedIterator{contract: _BEP20KFIVECrowdsale.contract, event: "TokensPurchased", logs: logs, sub: sub}, nil
}

// WatchTokensPurchased is a free log subscription operation binding the contract event 0x6faf93231a456e552dbc9961f58d9713ee4f2e69d15f1975b050ef0911053a7b.
//
// Solidity: event TokensPurchased(address indexed purchaser, address indexed beneficiary, uint256 value, uint256 amount)
func (_BEP20KFIVECrowdsale *BEP20KFIVECrowdsaleFilterer) WatchTokensPurchased(opts *bind.WatchOpts, sink chan<- *BEP20KFIVECrowdsaleTokensPurchased, purchaser []common.Address, beneficiary []common.Address) (event.Subscription, error) {

	var purchaserRule []interface{}
	for _, purchaserItem := range purchaser {
		purchaserRule = append(purchaserRule, purchaserItem)
	}
	var beneficiaryRule []interface{}
	for _, beneficiaryItem := range beneficiary {
		beneficiaryRule = append(beneficiaryRule, beneficiaryItem)
	}

	logs, sub, err := _BEP20KFIVECrowdsale.contract.WatchLogs(opts, "TokensPurchased", purchaserRule, beneficiaryRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BEP20KFIVECrowdsaleTokensPurchased)
				if err := _BEP20KFIVECrowdsale.contract.UnpackLog(event, "TokensPurchased", log); err != nil {
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
func (_BEP20KFIVECrowdsale *BEP20KFIVECrowdsaleFilterer) ParseTokensPurchased(log types.Log) (*BEP20KFIVECrowdsaleTokensPurchased, error) {
	event := new(BEP20KFIVECrowdsaleTokensPurchased)
	if err := _BEP20KFIVECrowdsale.contract.UnpackLog(event, "TokensPurchased", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BEP20KFIVECrowdsaleUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the BEP20KFIVECrowdsale contract.
type BEP20KFIVECrowdsaleUnpausedIterator struct {
	Event *BEP20KFIVECrowdsaleUnpaused // Event containing the contract specifics and raw log

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
func (it *BEP20KFIVECrowdsaleUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BEP20KFIVECrowdsaleUnpaused)
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
		it.Event = new(BEP20KFIVECrowdsaleUnpaused)
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
func (it *BEP20KFIVECrowdsaleUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BEP20KFIVECrowdsaleUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BEP20KFIVECrowdsaleUnpaused represents a Unpaused event raised by the BEP20KFIVECrowdsale contract.
type BEP20KFIVECrowdsaleUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_BEP20KFIVECrowdsale *BEP20KFIVECrowdsaleFilterer) FilterUnpaused(opts *bind.FilterOpts) (*BEP20KFIVECrowdsaleUnpausedIterator, error) {

	logs, sub, err := _BEP20KFIVECrowdsale.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &BEP20KFIVECrowdsaleUnpausedIterator{contract: _BEP20KFIVECrowdsale.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_BEP20KFIVECrowdsale *BEP20KFIVECrowdsaleFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *BEP20KFIVECrowdsaleUnpaused) (event.Subscription, error) {

	logs, sub, err := _BEP20KFIVECrowdsale.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BEP20KFIVECrowdsaleUnpaused)
				if err := _BEP20KFIVECrowdsale.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
func (_BEP20KFIVECrowdsale *BEP20KFIVECrowdsaleFilterer) ParseUnpaused(log types.Log) (*BEP20KFIVECrowdsaleUnpaused, error) {
	event := new(BEP20KFIVECrowdsaleUnpaused)
	if err := _BEP20KFIVECrowdsale.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
