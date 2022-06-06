// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package KFIVECrowdsale

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

// KFIVECrowdsaleMetaData contains all meta data concerning the KFIVECrowdsale contract.
var KFIVECrowdsaleMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"rate\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"wallet\",\"type\":\"address\"},{\"internalType\":\"contractIBEP20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"cap\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"openingTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"closingTime\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"prevClosingTime\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newClosingTime\",\"type\":\"uint256\"}],\"name\":\"TimedCrowdsaleExtended\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"purchaser\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"beneficiary\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"TokensPurchased\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"beneficiary\",\"type\":\"address\"}],\"name\":\"buyTokens\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"cap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"capReached\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"newReceiver\",\"type\":\"address\"}],\"name\":\"changeWallet\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"closingTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"_to\",\"type\":\"address\"}],\"name\":\"destroySmartContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"hasClosed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isOpen\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"openingTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token\",\"outputs\":[{\"internalType\":\"contractIBEP20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"wallet\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"weiRaised\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Bin: "0x60806040523480156200001157600080fd5b50604051620023c3380380620023c383398181016040528101906200003791906200043d565b8181848888886200005d620000516200032c60201b60201c565b6200033460201b60201c565b600180819055506000600260006101000a81548160ff02191690831515021790555060008311620000c5576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401620000bc9062000667565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16141562000138576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016200012f90620005df565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff161415620001ab576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401620001a29062000623565b60405180910390fd5b8260048190555081600360006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555080600260016101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550505050600081116200027d576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401620002749062000645565b60405180910390fd5b806006819055505042821015620002cb576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401620002c290620005bd565b60405180910390fd5b81811162000310576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401620003079062000601565b60405180910390fd5b8160078190555080600881905550505050505050505062000890565b600033905090565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050816000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b600081519050620004098162000842565b92915050565b60008151905062000420816200085c565b92915050565b600081519050620004378162000876565b92915050565b60008060008060008060c087890312156200045757600080fd5b60006200046789828a0162000426565b96505060206200047a89828a01620003f8565b95505060406200048d89828a016200040f565b9450506060620004a089828a0162000426565b9350506080620004b389828a0162000426565b92505060a0620004c689828a0162000426565b9150509295509295509295565b6000620004e260338362000689565b9150620004ef8262000700565b604082019050919050565b600062000509600f8362000689565b915062000516826200074f565b602082019050919050565b60006200053060378362000689565b91506200053d8262000778565b604082019050919050565b600062000557600d8362000689565b91506200056482620007c7565b602082019050919050565b60006200057e60198362000689565b91506200058b82620007f0565b602082019050919050565b6000620005a5600c8362000689565b9150620005b28262000819565b602082019050919050565b60006020820190508181036000830152620005d881620004d3565b9050919050565b60006020820190508181036000830152620005fa81620004fa565b9050919050565b600060208201905081810360008301526200061c8162000521565b9050919050565b600060208201905081810360008301526200063e8162000548565b9050919050565b6000602082019050818103600083015262000660816200056f565b9050919050565b60006020820190508181036000830152620006828162000596565b9050919050565b600082825260208201905092915050565b6000620006a782620006d6565b9050919050565b6000620006bb82620006d6565b9050919050565b6000620006cf826200069a565b9050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b7f54696d656443726f776473616c653a206f70656e696e672074696d652069732060008201527f6265666f72652063757272656e742074696d6500000000000000000000000000602082015250565b7f496e76616c696420616464726573730000000000000000000000000000000000600082015250565b7f54696d656443726f776473616c653a206f70656e696e672074696d652069732060008201527f6e6f74206265666f726520636c6f73696e672074696d65000000000000000000602082015250565b7f496e76616c696420746f6b656e00000000000000000000000000000000000000600082015250565b7f43617070656443726f776473616c653a20636170206973203000000000000000600082015250565b7f496e76616c696420726174650000000000000000000000000000000000000000600082015250565b6200084d81620006ae565b81146200085957600080fd5b50565b6200086781620006c2565b81146200087357600080fd5b50565b6200088181620006f6565b81146200088d57600080fd5b50565b611b2380620008a06000396000f3fe6080604052600436106101185760003560e01c8063521eb273116100a057806398b9a2dc1161006457806398b9a2dc14610350578063b7a8807c14610379578063ec8ac4d8146103a4578063f2fde38b146103c0578063fc0c546a146103e95761012f565b8063521eb273146102a15780635c975abb146102cc578063715018a6146102f75780638456cb591461030e5780638da5cb5b146103255761012f565b80633f4ba83a116100e75780633f4ba83a146101de5780634042b66f146101f557806347535d7b146102205780634b6753bc1461024b5780634f935945146102765761012f565b80631515bc2b146101345780632c4e722e1461015f578063355274ea1461018a57806339df43ff146101b55761012f565b3661012f5761012d610128610414565b61041c565b005b600080fd5b34801561014057600080fd5b506101496105e6565b6040516101569190611513565b60405180910390f35b34801561016b57600080fd5b506101746105f2565b6040516101819190611689565b60405180910390f35b34801561019657600080fd5b5061019f6105fc565b6040516101ac9190611689565b60405180910390f35b3480156101c157600080fd5b506101dc60048036038101906101d79190611258565b610606565b005b3480156101ea57600080fd5b506101f36107fa565b005b34801561020157600080fd5b5061020a6108c7565b6040516102179190611689565b60405180910390f35b34801561022c57600080fd5b506102356108d1565b6040516102429190611513565b60405180910390f35b34801561025757600080fd5b506102606108ec565b60405161026d9190611689565b60405180910390f35b34801561028257600080fd5b5061028b6108f6565b6040516102989190611513565b60405180910390f35b3480156102ad57600080fd5b506102b661090a565b6040516102c391906114a6565b60405180910390f35b3480156102d857600080fd5b506102e1610934565b6040516102ee9190611513565b60405180910390f35b34801561030357600080fd5b5061030c61094b565b005b34801561031a57600080fd5b506103236109d3565b005b34801561033157600080fd5b5061033a610aa1565b604051610347919061148b565b60405180910390f35b34801561035c57600080fd5b5061037760048036038101906103729190611258565b610aca565b005b34801561038557600080fd5b5061038e610b8a565b60405161039b9190611689565b60405180910390f35b6103be60048036038101906103b9919061122f565b61041c565b005b3480156103cc57600080fd5b506103e760048036038101906103e2919061122f565b610b94565b005b3480156103f557600080fd5b506103fe610c8c565b60405161040b919061152e565b60405180910390f35b600033905090565b60026001541415610462576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161045990611669565b60405180910390fd5b6002600181905550610472610934565b156104b2576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016104a9906115e9565b60405180910390fd5b633b9aca003410156104f9576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016104f0906115c9565b60405180910390fd5b6000610512633b9aca0034610cb690919063ffffffff16565b905061051e8282610ccc565b600061052982610cda565b905061054082600554610cf890919063ffffffff16565b6005819055506105508382610d0e565b8273ffffffffffffffffffffffffffffffffffffffff1661056f610414565b73ffffffffffffffffffffffffffffffffffffffff167f6faf93231a456e552dbc9961f58d9713ee4f2e69d15f1975b050ef0911053a7b84846040516105b69291906116a4565b60405180910390a36105c88383610d1c565b6105d0610d20565b6105da8383610d8b565b50506001808190555050565b60006008544211905090565b6000600454905090565b6000600654905090565b61060e610414565b73ffffffffffffffffffffffffffffffffffffffff1661062c610aa1565b73ffffffffffffffffffffffffffffffffffffffff1614610682576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161067990611629565b60405180910390fd5b6000600260019054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166370a08231306040518263ffffffff1660e01b81526004016106df919061148b565b60206040518083038186803b1580156106f757600080fd5b505afa15801561070b573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061072f91906112aa565b9050600260019054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663a9059cbb83836040518363ffffffff1660e01b815260040161078e9291906114c1565b602060405180830381600087803b1580156107a857600080fd5b505af11580156107bc573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906107e09190611281565b508173ffffffffffffffffffffffffffffffffffffffff16ff5b610802610934565b610841576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161083890611549565b60405180910390fd5b610849610414565b73ffffffffffffffffffffffffffffffffffffffff16610867610aa1565b73ffffffffffffffffffffffffffffffffffffffff16146108bd576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016108b490611629565b60405180910390fd5b6108c5610d8f565b565b6000600554905090565b600060075442101580156108e757506008544211155b905090565b6000600854905090565b60006006546109036108c7565b1015905090565b6000600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b6000600260009054906101000a900460ff16905090565b610953610414565b73ffffffffffffffffffffffffffffffffffffffff16610971610aa1565b73ffffffffffffffffffffffffffffffffffffffff16146109c7576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016109be90611629565b60405180910390fd5b6109d16000610e31565b565b6109db610934565b15610a1b576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610a12906115e9565b60405180910390fd5b610a23610414565b73ffffffffffffffffffffffffffffffffffffffff16610a41610aa1565b73ffffffffffffffffffffffffffffffffffffffff1614610a97576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610a8e90611629565b60405180910390fd5b610a9f610ef5565b565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b610ad2610414565b73ffffffffffffffffffffffffffffffffffffffff16610af0610aa1565b73ffffffffffffffffffffffffffffffffffffffff1614610b46576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610b3d90611629565b60405180910390fd5b80600360006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b6000600754905090565b610b9c610414565b73ffffffffffffffffffffffffffffffffffffffff16610bba610aa1565b73ffffffffffffffffffffffffffffffffffffffff1614610c10576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610c0790611629565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff161415610c80576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610c7790611569565b60405180910390fd5b610c8981610e31565b50565b6000600260019054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b60008183610cc49190611734565b905092915050565b610cd68282610f98565b5050565b6000610cf160045483610fed90919063ffffffff16565b9050919050565b60008183610d0691906116de565b905092915050565b610d188282611003565b5050565b5050565b600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166108fc349081150290604051600060405180830381858888f19350505050158015610d88573d6000803e3d6000fd5b50565b5050565b610d97610934565b610dd6576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610dcd90611549565b60405180910390fd5b6000600260006101000a81548160ff0219169083151502179055507f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa610e1a610414565b604051610e27919061148b565b60405180910390a1565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050816000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b610efd610934565b15610f3d576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610f34906115e9565b60405180910390fd5b6001600260006101000a81548160ff0219169083151502179055507f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258610f81610414565b604051610f8e919061148b565b60405180910390a1565b610fa06108d1565b610fdf576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610fd690611589565b60405180910390fd5b610fe982826110b7565b5050565b60008183610ffb9190611765565b905092915050565b600260019054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663a9059cbb83836040518363ffffffff1660e01b81526004016110609291906114ea565b602060405180830381600087803b15801561107a57600080fd5b505af115801561108e573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906110b29190611281565b505050565b6110c18282611123565b6006546110de826110d06108c7565b610cf890919063ffffffff16565b111561111f576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611116906115a9565b60405180910390fd5b5050565b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff161415611193576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161118a90611609565b60405180910390fd5b60008114156111d7576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016111ce90611649565b60405180910390fd5b5050565b6000813590506111ea81611a91565b92915050565b6000813590506111ff81611aa8565b92915050565b60008151905061121481611abf565b92915050565b60008151905061122981611ad6565b92915050565b60006020828403121561124157600080fd5b600061124f848285016111db565b91505092915050565b60006020828403121561126a57600080fd5b6000611278848285016111f0565b91505092915050565b60006020828403121561129357600080fd5b60006112a184828501611205565b91505092915050565b6000602082840312156112bc57600080fd5b60006112ca8482850161121a565b91505092915050565b6112dc81611819565b82525050565b6112eb816117d1565b82525050565b6112fa816117bf565b82525050565b611309816117e3565b82525050565b6113188161182b565b82525050565b600061132b6014836116cd565b9150611336826118d1565b602082019050919050565b600061134e6026836116cd565b9150611359826118fa565b604082019050919050565b60006113716018836116cd565b915061137c82611949565b602082019050919050565b6000611394601d836116cd565b915061139f82611972565b602082019050919050565b60006113b7600e836116cd565b91506113c28261199b565b602082019050919050565b60006113da6010836116cd565b91506113e5826119c4565b602082019050919050565b60006113fd6013836116cd565b9150611408826119ed565b602082019050919050565b60006114206020836116cd565b915061142b82611a16565b602082019050919050565b60006114436007836116cd565b915061144e82611a3f565b602082019050919050565b6000611466601f836116cd565b915061147182611a68565b602082019050919050565b6114858161180f565b82525050565b60006020820190506114a060008301846112f1565b92915050565b60006020820190506114bb60008301846112e2565b92915050565b60006040820190506114d660008301856112d3565b6114e3602083018461147c565b9392505050565b60006040820190506114ff60008301856112f1565b61150c602083018461147c565b9392505050565b60006020820190506115286000830184611300565b92915050565b6000602082019050611543600083018461130f565b92915050565b600060208201905081810360008301526115628161131e565b9050919050565b6000602082019050818103600083015261158281611341565b9050919050565b600060208201905081810360008301526115a281611364565b9050919050565b600060208201905081810360008301526115c281611387565b9050919050565b600060208201905081810360008301526115e2816113aa565b9050919050565b60006020820190508181036000830152611602816113cd565b9050919050565b60006020820190508181036000830152611622816113f0565b9050919050565b6000602082019050818103600083015261164281611413565b9050919050565b6000602082019050818103600083015261166281611436565b9050919050565b6000602082019050818103600083015261168281611459565b9050919050565b600060208201905061169e600083018461147c565b92915050565b60006040820190506116b9600083018561147c565b6116c6602083018461147c565b9392505050565b600082825260208201905092915050565b60006116e98261180f565b91506116f48361180f565b9250827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0382111561172957611728611873565b5b828201905092915050565b600061173f8261180f565b915061174a8361180f565b92508261175a576117596118a2565b5b828204905092915050565b60006117708261180f565b915061177b8361180f565b9250817fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff04831182151516156117b4576117b3611873565b5b828202905092915050565b60006117ca826117ef565b9050919050565b60006117dc826117ef565b9050919050565b60008115159050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b60006118248261184f565b9050919050565b60006118368261183d565b9050919050565b6000611848826117ef565b9050919050565b600061185a82611861565b9050919050565b600061186c826117ef565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b7f5061757361626c653a206e6f7420706175736564000000000000000000000000600082015250565b7f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160008201527f6464726573730000000000000000000000000000000000000000000000000000602082015250565b7f54696d656443726f776473616c653a206e6f74206f70656e0000000000000000600082015250565b7f43617070656443726f776473616c653a20636170206578636565646564000000600082015250565b7f496e76616c696420616d6f756e74000000000000000000000000000000000000600082015250565b7f5061757361626c653a2070617573656400000000000000000000000000000000600082015250565b7f496e76616c69642062656e656669636961727900000000000000000000000000600082015250565b7f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572600082015250565b7f496e76616c696400000000000000000000000000000000000000000000000000600082015250565b7f5265656e7472616e637947756172643a207265656e7472616e742063616c6c00600082015250565b611a9a816117bf565b8114611aa557600080fd5b50565b611ab1816117d1565b8114611abc57600080fd5b50565b611ac8816117e3565b8114611ad357600080fd5b50565b611adf8161180f565b8114611aea57600080fd5b5056fea26469706673582212207ab1db9771b162fdfb32ab3731db0d1229558ebb89a61c10bf46a9fffcec3eb464736f6c63430008040033",
}

// KFIVECrowdsaleABI is the input ABI used to generate the binding from.
// Deprecated: Use KFIVECrowdsaleMetaData.ABI instead.
var KFIVECrowdsaleABI = KFIVECrowdsaleMetaData.ABI

// KFIVECrowdsaleBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use KFIVECrowdsaleMetaData.Bin instead.
var KFIVECrowdsaleBin = KFIVECrowdsaleMetaData.Bin

// DeployKFIVECrowdsale deploys a new Ethereum contract, binding an instance of KFIVECrowdsale to it.
func DeployKFIVECrowdsale(auth *bind.TransactOpts, backend bind.ContractBackend, rate *big.Int, wallet common.Address, token common.Address, cap *big.Int, openingTime *big.Int, closingTime *big.Int) (common.Address, *types.Transaction, *KFIVECrowdsale, error) {
	parsed, err := KFIVECrowdsaleMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(KFIVECrowdsaleBin), backend, rate, wallet, token, cap, openingTime, closingTime)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &KFIVECrowdsale{KFIVECrowdsaleCaller: KFIVECrowdsaleCaller{contract: contract}, KFIVECrowdsaleTransactor: KFIVECrowdsaleTransactor{contract: contract}, KFIVECrowdsaleFilterer: KFIVECrowdsaleFilterer{contract: contract}}, nil
}

// KFIVECrowdsale is an auto generated Go binding around an Ethereum contract.
type KFIVECrowdsale struct {
	KFIVECrowdsaleCaller     // Read-only binding to the contract
	KFIVECrowdsaleTransactor // Write-only binding to the contract
	KFIVECrowdsaleFilterer   // Log filterer for contract events
}

// KFIVECrowdsaleCaller is an auto generated read-only Go binding around an Ethereum contract.
type KFIVECrowdsaleCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KFIVECrowdsaleTransactor is an auto generated write-only Go binding around an Ethereum contract.
type KFIVECrowdsaleTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KFIVECrowdsaleFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type KFIVECrowdsaleFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KFIVECrowdsaleSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type KFIVECrowdsaleSession struct {
	Contract     *KFIVECrowdsale   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// KFIVECrowdsaleCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type KFIVECrowdsaleCallerSession struct {
	Contract *KFIVECrowdsaleCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// KFIVECrowdsaleTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type KFIVECrowdsaleTransactorSession struct {
	Contract     *KFIVECrowdsaleTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// KFIVECrowdsaleRaw is an auto generated low-level Go binding around an Ethereum contract.
type KFIVECrowdsaleRaw struct {
	Contract *KFIVECrowdsale // Generic contract binding to access the raw methods on
}

// KFIVECrowdsaleCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type KFIVECrowdsaleCallerRaw struct {
	Contract *KFIVECrowdsaleCaller // Generic read-only contract binding to access the raw methods on
}

// KFIVECrowdsaleTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type KFIVECrowdsaleTransactorRaw struct {
	Contract *KFIVECrowdsaleTransactor // Generic write-only contract binding to access the raw methods on
}

// NewKFIVECrowdsale creates a new instance of KFIVECrowdsale, bound to a specific deployed contract.
func NewKFIVECrowdsale(address common.Address, backend bind.ContractBackend) (*KFIVECrowdsale, error) {
	contract, err := bindKFIVECrowdsale(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &KFIVECrowdsale{KFIVECrowdsaleCaller: KFIVECrowdsaleCaller{contract: contract}, KFIVECrowdsaleTransactor: KFIVECrowdsaleTransactor{contract: contract}, KFIVECrowdsaleFilterer: KFIVECrowdsaleFilterer{contract: contract}}, nil
}

// NewKFIVECrowdsaleCaller creates a new read-only instance of KFIVECrowdsale, bound to a specific deployed contract.
func NewKFIVECrowdsaleCaller(address common.Address, caller bind.ContractCaller) (*KFIVECrowdsaleCaller, error) {
	contract, err := bindKFIVECrowdsale(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &KFIVECrowdsaleCaller{contract: contract}, nil
}

// NewKFIVECrowdsaleTransactor creates a new write-only instance of KFIVECrowdsale, bound to a specific deployed contract.
func NewKFIVECrowdsaleTransactor(address common.Address, transactor bind.ContractTransactor) (*KFIVECrowdsaleTransactor, error) {
	contract, err := bindKFIVECrowdsale(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &KFIVECrowdsaleTransactor{contract: contract}, nil
}

// NewKFIVECrowdsaleFilterer creates a new log filterer instance of KFIVECrowdsale, bound to a specific deployed contract.
func NewKFIVECrowdsaleFilterer(address common.Address, filterer bind.ContractFilterer) (*KFIVECrowdsaleFilterer, error) {
	contract, err := bindKFIVECrowdsale(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &KFIVECrowdsaleFilterer{contract: contract}, nil
}

// bindKFIVECrowdsale binds a generic wrapper to an already deployed contract.
func bindKFIVECrowdsale(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(KFIVECrowdsaleABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_KFIVECrowdsale *KFIVECrowdsaleRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _KFIVECrowdsale.Contract.KFIVECrowdsaleCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_KFIVECrowdsale *KFIVECrowdsaleRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _KFIVECrowdsale.Contract.KFIVECrowdsaleTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_KFIVECrowdsale *KFIVECrowdsaleRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _KFIVECrowdsale.Contract.KFIVECrowdsaleTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_KFIVECrowdsale *KFIVECrowdsaleCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _KFIVECrowdsale.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_KFIVECrowdsale *KFIVECrowdsaleTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _KFIVECrowdsale.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_KFIVECrowdsale *KFIVECrowdsaleTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _KFIVECrowdsale.Contract.contract.Transact(opts, method, params...)
}

// Cap is a free data retrieval call binding the contract method 0x355274ea.
//
// Solidity: function cap() view returns(uint256)
func (_KFIVECrowdsale *KFIVECrowdsaleCaller) Cap(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _KFIVECrowdsale.contract.Call(opts, &out, "cap")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Cap is a free data retrieval call binding the contract method 0x355274ea.
//
// Solidity: function cap() view returns(uint256)
func (_KFIVECrowdsale *KFIVECrowdsaleSession) Cap() (*big.Int, error) {
	return _KFIVECrowdsale.Contract.Cap(&_KFIVECrowdsale.CallOpts)
}

// Cap is a free data retrieval call binding the contract method 0x355274ea.
//
// Solidity: function cap() view returns(uint256)
func (_KFIVECrowdsale *KFIVECrowdsaleCallerSession) Cap() (*big.Int, error) {
	return _KFIVECrowdsale.Contract.Cap(&_KFIVECrowdsale.CallOpts)
}

// CapReached is a free data retrieval call binding the contract method 0x4f935945.
//
// Solidity: function capReached() view returns(bool)
func (_KFIVECrowdsale *KFIVECrowdsaleCaller) CapReached(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _KFIVECrowdsale.contract.Call(opts, &out, "capReached")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CapReached is a free data retrieval call binding the contract method 0x4f935945.
//
// Solidity: function capReached() view returns(bool)
func (_KFIVECrowdsale *KFIVECrowdsaleSession) CapReached() (bool, error) {
	return _KFIVECrowdsale.Contract.CapReached(&_KFIVECrowdsale.CallOpts)
}

// CapReached is a free data retrieval call binding the contract method 0x4f935945.
//
// Solidity: function capReached() view returns(bool)
func (_KFIVECrowdsale *KFIVECrowdsaleCallerSession) CapReached() (bool, error) {
	return _KFIVECrowdsale.Contract.CapReached(&_KFIVECrowdsale.CallOpts)
}

// ClosingTime is a free data retrieval call binding the contract method 0x4b6753bc.
//
// Solidity: function closingTime() view returns(uint256)
func (_KFIVECrowdsale *KFIVECrowdsaleCaller) ClosingTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _KFIVECrowdsale.contract.Call(opts, &out, "closingTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ClosingTime is a free data retrieval call binding the contract method 0x4b6753bc.
//
// Solidity: function closingTime() view returns(uint256)
func (_KFIVECrowdsale *KFIVECrowdsaleSession) ClosingTime() (*big.Int, error) {
	return _KFIVECrowdsale.Contract.ClosingTime(&_KFIVECrowdsale.CallOpts)
}

// ClosingTime is a free data retrieval call binding the contract method 0x4b6753bc.
//
// Solidity: function closingTime() view returns(uint256)
func (_KFIVECrowdsale *KFIVECrowdsaleCallerSession) ClosingTime() (*big.Int, error) {
	return _KFIVECrowdsale.Contract.ClosingTime(&_KFIVECrowdsale.CallOpts)
}

// HasClosed is a free data retrieval call binding the contract method 0x1515bc2b.
//
// Solidity: function hasClosed() view returns(bool)
func (_KFIVECrowdsale *KFIVECrowdsaleCaller) HasClosed(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _KFIVECrowdsale.contract.Call(opts, &out, "hasClosed")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasClosed is a free data retrieval call binding the contract method 0x1515bc2b.
//
// Solidity: function hasClosed() view returns(bool)
func (_KFIVECrowdsale *KFIVECrowdsaleSession) HasClosed() (bool, error) {
	return _KFIVECrowdsale.Contract.HasClosed(&_KFIVECrowdsale.CallOpts)
}

// HasClosed is a free data retrieval call binding the contract method 0x1515bc2b.
//
// Solidity: function hasClosed() view returns(bool)
func (_KFIVECrowdsale *KFIVECrowdsaleCallerSession) HasClosed() (bool, error) {
	return _KFIVECrowdsale.Contract.HasClosed(&_KFIVECrowdsale.CallOpts)
}

// IsOpen is a free data retrieval call binding the contract method 0x47535d7b.
//
// Solidity: function isOpen() view returns(bool)
func (_KFIVECrowdsale *KFIVECrowdsaleCaller) IsOpen(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _KFIVECrowdsale.contract.Call(opts, &out, "isOpen")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsOpen is a free data retrieval call binding the contract method 0x47535d7b.
//
// Solidity: function isOpen() view returns(bool)
func (_KFIVECrowdsale *KFIVECrowdsaleSession) IsOpen() (bool, error) {
	return _KFIVECrowdsale.Contract.IsOpen(&_KFIVECrowdsale.CallOpts)
}

// IsOpen is a free data retrieval call binding the contract method 0x47535d7b.
//
// Solidity: function isOpen() view returns(bool)
func (_KFIVECrowdsale *KFIVECrowdsaleCallerSession) IsOpen() (bool, error) {
	return _KFIVECrowdsale.Contract.IsOpen(&_KFIVECrowdsale.CallOpts)
}

// OpeningTime is a free data retrieval call binding the contract method 0xb7a8807c.
//
// Solidity: function openingTime() view returns(uint256)
func (_KFIVECrowdsale *KFIVECrowdsaleCaller) OpeningTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _KFIVECrowdsale.contract.Call(opts, &out, "openingTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// OpeningTime is a free data retrieval call binding the contract method 0xb7a8807c.
//
// Solidity: function openingTime() view returns(uint256)
func (_KFIVECrowdsale *KFIVECrowdsaleSession) OpeningTime() (*big.Int, error) {
	return _KFIVECrowdsale.Contract.OpeningTime(&_KFIVECrowdsale.CallOpts)
}

// OpeningTime is a free data retrieval call binding the contract method 0xb7a8807c.
//
// Solidity: function openingTime() view returns(uint256)
func (_KFIVECrowdsale *KFIVECrowdsaleCallerSession) OpeningTime() (*big.Int, error) {
	return _KFIVECrowdsale.Contract.OpeningTime(&_KFIVECrowdsale.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_KFIVECrowdsale *KFIVECrowdsaleCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _KFIVECrowdsale.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_KFIVECrowdsale *KFIVECrowdsaleSession) Owner() (common.Address, error) {
	return _KFIVECrowdsale.Contract.Owner(&_KFIVECrowdsale.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_KFIVECrowdsale *KFIVECrowdsaleCallerSession) Owner() (common.Address, error) {
	return _KFIVECrowdsale.Contract.Owner(&_KFIVECrowdsale.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_KFIVECrowdsale *KFIVECrowdsaleCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _KFIVECrowdsale.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_KFIVECrowdsale *KFIVECrowdsaleSession) Paused() (bool, error) {
	return _KFIVECrowdsale.Contract.Paused(&_KFIVECrowdsale.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_KFIVECrowdsale *KFIVECrowdsaleCallerSession) Paused() (bool, error) {
	return _KFIVECrowdsale.Contract.Paused(&_KFIVECrowdsale.CallOpts)
}

// Rate is a free data retrieval call binding the contract method 0x2c4e722e.
//
// Solidity: function rate() view returns(uint256)
func (_KFIVECrowdsale *KFIVECrowdsaleCaller) Rate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _KFIVECrowdsale.contract.Call(opts, &out, "rate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Rate is a free data retrieval call binding the contract method 0x2c4e722e.
//
// Solidity: function rate() view returns(uint256)
func (_KFIVECrowdsale *KFIVECrowdsaleSession) Rate() (*big.Int, error) {
	return _KFIVECrowdsale.Contract.Rate(&_KFIVECrowdsale.CallOpts)
}

// Rate is a free data retrieval call binding the contract method 0x2c4e722e.
//
// Solidity: function rate() view returns(uint256)
func (_KFIVECrowdsale *KFIVECrowdsaleCallerSession) Rate() (*big.Int, error) {
	return _KFIVECrowdsale.Contract.Rate(&_KFIVECrowdsale.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_KFIVECrowdsale *KFIVECrowdsaleCaller) Token(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _KFIVECrowdsale.contract.Call(opts, &out, "token")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_KFIVECrowdsale *KFIVECrowdsaleSession) Token() (common.Address, error) {
	return _KFIVECrowdsale.Contract.Token(&_KFIVECrowdsale.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_KFIVECrowdsale *KFIVECrowdsaleCallerSession) Token() (common.Address, error) {
	return _KFIVECrowdsale.Contract.Token(&_KFIVECrowdsale.CallOpts)
}

// Wallet is a free data retrieval call binding the contract method 0x521eb273.
//
// Solidity: function wallet() view returns(address)
func (_KFIVECrowdsale *KFIVECrowdsaleCaller) Wallet(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _KFIVECrowdsale.contract.Call(opts, &out, "wallet")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Wallet is a free data retrieval call binding the contract method 0x521eb273.
//
// Solidity: function wallet() view returns(address)
func (_KFIVECrowdsale *KFIVECrowdsaleSession) Wallet() (common.Address, error) {
	return _KFIVECrowdsale.Contract.Wallet(&_KFIVECrowdsale.CallOpts)
}

// Wallet is a free data retrieval call binding the contract method 0x521eb273.
//
// Solidity: function wallet() view returns(address)
func (_KFIVECrowdsale *KFIVECrowdsaleCallerSession) Wallet() (common.Address, error) {
	return _KFIVECrowdsale.Contract.Wallet(&_KFIVECrowdsale.CallOpts)
}

// WeiRaised is a free data retrieval call binding the contract method 0x4042b66f.
//
// Solidity: function weiRaised() view returns(uint256)
func (_KFIVECrowdsale *KFIVECrowdsaleCaller) WeiRaised(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _KFIVECrowdsale.contract.Call(opts, &out, "weiRaised")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WeiRaised is a free data retrieval call binding the contract method 0x4042b66f.
//
// Solidity: function weiRaised() view returns(uint256)
func (_KFIVECrowdsale *KFIVECrowdsaleSession) WeiRaised() (*big.Int, error) {
	return _KFIVECrowdsale.Contract.WeiRaised(&_KFIVECrowdsale.CallOpts)
}

// WeiRaised is a free data retrieval call binding the contract method 0x4042b66f.
//
// Solidity: function weiRaised() view returns(uint256)
func (_KFIVECrowdsale *KFIVECrowdsaleCallerSession) WeiRaised() (*big.Int, error) {
	return _KFIVECrowdsale.Contract.WeiRaised(&_KFIVECrowdsale.CallOpts)
}

// BuyTokens is a paid mutator transaction binding the contract method 0xec8ac4d8.
//
// Solidity: function buyTokens(address beneficiary) payable returns()
func (_KFIVECrowdsale *KFIVECrowdsaleTransactor) BuyTokens(opts *bind.TransactOpts, beneficiary common.Address) (*types.Transaction, error) {
	return _KFIVECrowdsale.contract.Transact(opts, "buyTokens", beneficiary)
}

// BuyTokens is a paid mutator transaction binding the contract method 0xec8ac4d8.
//
// Solidity: function buyTokens(address beneficiary) payable returns()
func (_KFIVECrowdsale *KFIVECrowdsaleSession) BuyTokens(beneficiary common.Address) (*types.Transaction, error) {
	return _KFIVECrowdsale.Contract.BuyTokens(&_KFIVECrowdsale.TransactOpts, beneficiary)
}

// BuyTokens is a paid mutator transaction binding the contract method 0xec8ac4d8.
//
// Solidity: function buyTokens(address beneficiary) payable returns()
func (_KFIVECrowdsale *KFIVECrowdsaleTransactorSession) BuyTokens(beneficiary common.Address) (*types.Transaction, error) {
	return _KFIVECrowdsale.Contract.BuyTokens(&_KFIVECrowdsale.TransactOpts, beneficiary)
}

// ChangeWallet is a paid mutator transaction binding the contract method 0x98b9a2dc.
//
// Solidity: function changeWallet(address newReceiver) returns()
func (_KFIVECrowdsale *KFIVECrowdsaleTransactor) ChangeWallet(opts *bind.TransactOpts, newReceiver common.Address) (*types.Transaction, error) {
	return _KFIVECrowdsale.contract.Transact(opts, "changeWallet", newReceiver)
}

// ChangeWallet is a paid mutator transaction binding the contract method 0x98b9a2dc.
//
// Solidity: function changeWallet(address newReceiver) returns()
func (_KFIVECrowdsale *KFIVECrowdsaleSession) ChangeWallet(newReceiver common.Address) (*types.Transaction, error) {
	return _KFIVECrowdsale.Contract.ChangeWallet(&_KFIVECrowdsale.TransactOpts, newReceiver)
}

// ChangeWallet is a paid mutator transaction binding the contract method 0x98b9a2dc.
//
// Solidity: function changeWallet(address newReceiver) returns()
func (_KFIVECrowdsale *KFIVECrowdsaleTransactorSession) ChangeWallet(newReceiver common.Address) (*types.Transaction, error) {
	return _KFIVECrowdsale.Contract.ChangeWallet(&_KFIVECrowdsale.TransactOpts, newReceiver)
}

// DestroySmartContract is a paid mutator transaction binding the contract method 0x39df43ff.
//
// Solidity: function destroySmartContract(address _to) returns()
func (_KFIVECrowdsale *KFIVECrowdsaleTransactor) DestroySmartContract(opts *bind.TransactOpts, _to common.Address) (*types.Transaction, error) {
	return _KFIVECrowdsale.contract.Transact(opts, "destroySmartContract", _to)
}

// DestroySmartContract is a paid mutator transaction binding the contract method 0x39df43ff.
//
// Solidity: function destroySmartContract(address _to) returns()
func (_KFIVECrowdsale *KFIVECrowdsaleSession) DestroySmartContract(_to common.Address) (*types.Transaction, error) {
	return _KFIVECrowdsale.Contract.DestroySmartContract(&_KFIVECrowdsale.TransactOpts, _to)
}

// DestroySmartContract is a paid mutator transaction binding the contract method 0x39df43ff.
//
// Solidity: function destroySmartContract(address _to) returns()
func (_KFIVECrowdsale *KFIVECrowdsaleTransactorSession) DestroySmartContract(_to common.Address) (*types.Transaction, error) {
	return _KFIVECrowdsale.Contract.DestroySmartContract(&_KFIVECrowdsale.TransactOpts, _to)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_KFIVECrowdsale *KFIVECrowdsaleTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _KFIVECrowdsale.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_KFIVECrowdsale *KFIVECrowdsaleSession) Pause() (*types.Transaction, error) {
	return _KFIVECrowdsale.Contract.Pause(&_KFIVECrowdsale.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_KFIVECrowdsale *KFIVECrowdsaleTransactorSession) Pause() (*types.Transaction, error) {
	return _KFIVECrowdsale.Contract.Pause(&_KFIVECrowdsale.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_KFIVECrowdsale *KFIVECrowdsaleTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _KFIVECrowdsale.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_KFIVECrowdsale *KFIVECrowdsaleSession) RenounceOwnership() (*types.Transaction, error) {
	return _KFIVECrowdsale.Contract.RenounceOwnership(&_KFIVECrowdsale.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_KFIVECrowdsale *KFIVECrowdsaleTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _KFIVECrowdsale.Contract.RenounceOwnership(&_KFIVECrowdsale.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_KFIVECrowdsale *KFIVECrowdsaleTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _KFIVECrowdsale.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_KFIVECrowdsale *KFIVECrowdsaleSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _KFIVECrowdsale.Contract.TransferOwnership(&_KFIVECrowdsale.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_KFIVECrowdsale *KFIVECrowdsaleTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _KFIVECrowdsale.Contract.TransferOwnership(&_KFIVECrowdsale.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_KFIVECrowdsale *KFIVECrowdsaleTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _KFIVECrowdsale.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_KFIVECrowdsale *KFIVECrowdsaleSession) Unpause() (*types.Transaction, error) {
	return _KFIVECrowdsale.Contract.Unpause(&_KFIVECrowdsale.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_KFIVECrowdsale *KFIVECrowdsaleTransactorSession) Unpause() (*types.Transaction, error) {
	return _KFIVECrowdsale.Contract.Unpause(&_KFIVECrowdsale.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_KFIVECrowdsale *KFIVECrowdsaleTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _KFIVECrowdsale.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_KFIVECrowdsale *KFIVECrowdsaleSession) Receive() (*types.Transaction, error) {
	return _KFIVECrowdsale.Contract.Receive(&_KFIVECrowdsale.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_KFIVECrowdsale *KFIVECrowdsaleTransactorSession) Receive() (*types.Transaction, error) {
	return _KFIVECrowdsale.Contract.Receive(&_KFIVECrowdsale.TransactOpts)
}

// KFIVECrowdsaleOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the KFIVECrowdsale contract.
type KFIVECrowdsaleOwnershipTransferredIterator struct {
	Event *KFIVECrowdsaleOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *KFIVECrowdsaleOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KFIVECrowdsaleOwnershipTransferred)
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
		it.Event = new(KFIVECrowdsaleOwnershipTransferred)
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
func (it *KFIVECrowdsaleOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KFIVECrowdsaleOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KFIVECrowdsaleOwnershipTransferred represents a OwnershipTransferred event raised by the KFIVECrowdsale contract.
type KFIVECrowdsaleOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_KFIVECrowdsale *KFIVECrowdsaleFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*KFIVECrowdsaleOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _KFIVECrowdsale.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &KFIVECrowdsaleOwnershipTransferredIterator{contract: _KFIVECrowdsale.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_KFIVECrowdsale *KFIVECrowdsaleFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *KFIVECrowdsaleOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _KFIVECrowdsale.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KFIVECrowdsaleOwnershipTransferred)
				if err := _KFIVECrowdsale.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_KFIVECrowdsale *KFIVECrowdsaleFilterer) ParseOwnershipTransferred(log types.Log) (*KFIVECrowdsaleOwnershipTransferred, error) {
	event := new(KFIVECrowdsaleOwnershipTransferred)
	if err := _KFIVECrowdsale.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// KFIVECrowdsalePausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the KFIVECrowdsale contract.
type KFIVECrowdsalePausedIterator struct {
	Event *KFIVECrowdsalePaused // Event containing the contract specifics and raw log

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
func (it *KFIVECrowdsalePausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KFIVECrowdsalePaused)
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
		it.Event = new(KFIVECrowdsalePaused)
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
func (it *KFIVECrowdsalePausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KFIVECrowdsalePausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KFIVECrowdsalePaused represents a Paused event raised by the KFIVECrowdsale contract.
type KFIVECrowdsalePaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_KFIVECrowdsale *KFIVECrowdsaleFilterer) FilterPaused(opts *bind.FilterOpts) (*KFIVECrowdsalePausedIterator, error) {

	logs, sub, err := _KFIVECrowdsale.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &KFIVECrowdsalePausedIterator{contract: _KFIVECrowdsale.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_KFIVECrowdsale *KFIVECrowdsaleFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *KFIVECrowdsalePaused) (event.Subscription, error) {

	logs, sub, err := _KFIVECrowdsale.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KFIVECrowdsalePaused)
				if err := _KFIVECrowdsale.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_KFIVECrowdsale *KFIVECrowdsaleFilterer) ParsePaused(log types.Log) (*KFIVECrowdsalePaused, error) {
	event := new(KFIVECrowdsalePaused)
	if err := _KFIVECrowdsale.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// KFIVECrowdsaleTimedCrowdsaleExtendedIterator is returned from FilterTimedCrowdsaleExtended and is used to iterate over the raw logs and unpacked data for TimedCrowdsaleExtended events raised by the KFIVECrowdsale contract.
type KFIVECrowdsaleTimedCrowdsaleExtendedIterator struct {
	Event *KFIVECrowdsaleTimedCrowdsaleExtended // Event containing the contract specifics and raw log

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
func (it *KFIVECrowdsaleTimedCrowdsaleExtendedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KFIVECrowdsaleTimedCrowdsaleExtended)
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
		it.Event = new(KFIVECrowdsaleTimedCrowdsaleExtended)
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
func (it *KFIVECrowdsaleTimedCrowdsaleExtendedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KFIVECrowdsaleTimedCrowdsaleExtendedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KFIVECrowdsaleTimedCrowdsaleExtended represents a TimedCrowdsaleExtended event raised by the KFIVECrowdsale contract.
type KFIVECrowdsaleTimedCrowdsaleExtended struct {
	PrevClosingTime *big.Int
	NewClosingTime  *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterTimedCrowdsaleExtended is a free log retrieval operation binding the contract event 0x46711e222f558a07afd26e5e71b48ecb0a8b2cdcd40faeb1323e05e2c76a2f32.
//
// Solidity: event TimedCrowdsaleExtended(uint256 prevClosingTime, uint256 newClosingTime)
func (_KFIVECrowdsale *KFIVECrowdsaleFilterer) FilterTimedCrowdsaleExtended(opts *bind.FilterOpts) (*KFIVECrowdsaleTimedCrowdsaleExtendedIterator, error) {

	logs, sub, err := _KFIVECrowdsale.contract.FilterLogs(opts, "TimedCrowdsaleExtended")
	if err != nil {
		return nil, err
	}
	return &KFIVECrowdsaleTimedCrowdsaleExtendedIterator{contract: _KFIVECrowdsale.contract, event: "TimedCrowdsaleExtended", logs: logs, sub: sub}, nil
}

// WatchTimedCrowdsaleExtended is a free log subscription operation binding the contract event 0x46711e222f558a07afd26e5e71b48ecb0a8b2cdcd40faeb1323e05e2c76a2f32.
//
// Solidity: event TimedCrowdsaleExtended(uint256 prevClosingTime, uint256 newClosingTime)
func (_KFIVECrowdsale *KFIVECrowdsaleFilterer) WatchTimedCrowdsaleExtended(opts *bind.WatchOpts, sink chan<- *KFIVECrowdsaleTimedCrowdsaleExtended) (event.Subscription, error) {

	logs, sub, err := _KFIVECrowdsale.contract.WatchLogs(opts, "TimedCrowdsaleExtended")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KFIVECrowdsaleTimedCrowdsaleExtended)
				if err := _KFIVECrowdsale.contract.UnpackLog(event, "TimedCrowdsaleExtended", log); err != nil {
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
func (_KFIVECrowdsale *KFIVECrowdsaleFilterer) ParseTimedCrowdsaleExtended(log types.Log) (*KFIVECrowdsaleTimedCrowdsaleExtended, error) {
	event := new(KFIVECrowdsaleTimedCrowdsaleExtended)
	if err := _KFIVECrowdsale.contract.UnpackLog(event, "TimedCrowdsaleExtended", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// KFIVECrowdsaleTokensPurchasedIterator is returned from FilterTokensPurchased and is used to iterate over the raw logs and unpacked data for TokensPurchased events raised by the KFIVECrowdsale contract.
type KFIVECrowdsaleTokensPurchasedIterator struct {
	Event *KFIVECrowdsaleTokensPurchased // Event containing the contract specifics and raw log

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
func (it *KFIVECrowdsaleTokensPurchasedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KFIVECrowdsaleTokensPurchased)
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
		it.Event = new(KFIVECrowdsaleTokensPurchased)
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
func (it *KFIVECrowdsaleTokensPurchasedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KFIVECrowdsaleTokensPurchasedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KFIVECrowdsaleTokensPurchased represents a TokensPurchased event raised by the KFIVECrowdsale contract.
type KFIVECrowdsaleTokensPurchased struct {
	Purchaser   common.Address
	Beneficiary common.Address
	Value       *big.Int
	Amount      *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterTokensPurchased is a free log retrieval operation binding the contract event 0x6faf93231a456e552dbc9961f58d9713ee4f2e69d15f1975b050ef0911053a7b.
//
// Solidity: event TokensPurchased(address indexed purchaser, address indexed beneficiary, uint256 value, uint256 amount)
func (_KFIVECrowdsale *KFIVECrowdsaleFilterer) FilterTokensPurchased(opts *bind.FilterOpts, purchaser []common.Address, beneficiary []common.Address) (*KFIVECrowdsaleTokensPurchasedIterator, error) {

	var purchaserRule []interface{}
	for _, purchaserItem := range purchaser {
		purchaserRule = append(purchaserRule, purchaserItem)
	}
	var beneficiaryRule []interface{}
	for _, beneficiaryItem := range beneficiary {
		beneficiaryRule = append(beneficiaryRule, beneficiaryItem)
	}

	logs, sub, err := _KFIVECrowdsale.contract.FilterLogs(opts, "TokensPurchased", purchaserRule, beneficiaryRule)
	if err != nil {
		return nil, err
	}
	return &KFIVECrowdsaleTokensPurchasedIterator{contract: _KFIVECrowdsale.contract, event: "TokensPurchased", logs: logs, sub: sub}, nil
}

// WatchTokensPurchased is a free log subscription operation binding the contract event 0x6faf93231a456e552dbc9961f58d9713ee4f2e69d15f1975b050ef0911053a7b.
//
// Solidity: event TokensPurchased(address indexed purchaser, address indexed beneficiary, uint256 value, uint256 amount)
func (_KFIVECrowdsale *KFIVECrowdsaleFilterer) WatchTokensPurchased(opts *bind.WatchOpts, sink chan<- *KFIVECrowdsaleTokensPurchased, purchaser []common.Address, beneficiary []common.Address) (event.Subscription, error) {

	var purchaserRule []interface{}
	for _, purchaserItem := range purchaser {
		purchaserRule = append(purchaserRule, purchaserItem)
	}
	var beneficiaryRule []interface{}
	for _, beneficiaryItem := range beneficiary {
		beneficiaryRule = append(beneficiaryRule, beneficiaryItem)
	}

	logs, sub, err := _KFIVECrowdsale.contract.WatchLogs(opts, "TokensPurchased", purchaserRule, beneficiaryRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KFIVECrowdsaleTokensPurchased)
				if err := _KFIVECrowdsale.contract.UnpackLog(event, "TokensPurchased", log); err != nil {
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
func (_KFIVECrowdsale *KFIVECrowdsaleFilterer) ParseTokensPurchased(log types.Log) (*KFIVECrowdsaleTokensPurchased, error) {
	event := new(KFIVECrowdsaleTokensPurchased)
	if err := _KFIVECrowdsale.contract.UnpackLog(event, "TokensPurchased", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// KFIVECrowdsaleUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the KFIVECrowdsale contract.
type KFIVECrowdsaleUnpausedIterator struct {
	Event *KFIVECrowdsaleUnpaused // Event containing the contract specifics and raw log

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
func (it *KFIVECrowdsaleUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KFIVECrowdsaleUnpaused)
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
		it.Event = new(KFIVECrowdsaleUnpaused)
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
func (it *KFIVECrowdsaleUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KFIVECrowdsaleUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KFIVECrowdsaleUnpaused represents a Unpaused event raised by the KFIVECrowdsale contract.
type KFIVECrowdsaleUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_KFIVECrowdsale *KFIVECrowdsaleFilterer) FilterUnpaused(opts *bind.FilterOpts) (*KFIVECrowdsaleUnpausedIterator, error) {

	logs, sub, err := _KFIVECrowdsale.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &KFIVECrowdsaleUnpausedIterator{contract: _KFIVECrowdsale.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_KFIVECrowdsale *KFIVECrowdsaleFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *KFIVECrowdsaleUnpaused) (event.Subscription, error) {

	logs, sub, err := _KFIVECrowdsale.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KFIVECrowdsaleUnpaused)
				if err := _KFIVECrowdsale.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
func (_KFIVECrowdsale *KFIVECrowdsaleFilterer) ParseUnpaused(log types.Log) (*KFIVECrowdsaleUnpaused, error) {
	event := new(KFIVECrowdsaleUnpaused)
	if err := _KFIVECrowdsale.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
