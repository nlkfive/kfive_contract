// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package Crowdsale

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

// CrowdsaleMetaData contains all meta data concerning the Crowdsale contract.
var CrowdsaleMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"__rate\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"__wallet\",\"type\":\"address\"},{\"internalType\":\"contractIBEP20\",\"name\":\"__token\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"purchaser\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"beneficiary\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"TokensPurchased\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"},{\"inputs\":[],\"name\":\"token\",\"outputs\":[{\"internalType\":\"contractIBEP20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"wallet\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"weiRaised\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"beneficiary\",\"type\":\"address\"}],\"name\":\"buyTokens\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"_to\",\"type\":\"address\"}],\"name\":\"destroySmartContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"newReceiver\",\"type\":\"address\"}],\"name\":\"changeWallet\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60806040523480156200001157600080fd5b5060405162001d9838038062001d98833981810160405281019062000037919062000348565b620000576200004b6200023760201b60201c565b6200023f60201b60201c565b600180819055506000600260006101000a81548160ff02191690831515021790555060008311620000bf576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401620000b69062000457565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16141562000132576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401620001299062000413565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff161415620001a5576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016200019c9062000435565b60405180910390fd5b8260048190555081600360006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555080600260016101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550505050620005b9565b600033905090565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050816000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b60008151905062000314816200056b565b92915050565b6000815190506200032b8162000585565b92915050565b60008151905062000342816200059f565b92915050565b6000806000606084860312156200035e57600080fd5b60006200036e8682870162000331565b9350506020620003818682870162000303565b925050604062000394868287016200031a565b9150509250925092565b6000620003ad600f8362000479565b9150620003ba82620004f0565b602082019050919050565b6000620003d4600d8362000479565b9150620003e18262000519565b602082019050919050565b6000620003fb600c8362000479565b9150620004088262000542565b602082019050919050565b600060208201905081810360008301526200042e816200039e565b9050919050565b600060208201905081810360008301526200045081620003c5565b9050919050565b600060208201905081810360008301526200047281620003ec565b9050919050565b600082825260208201905092915050565b60006200049782620004c6565b9050919050565b6000620004ab82620004c6565b9050919050565b6000620004bf826200048a565b9050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b7f496e76616c696420616464726573730000000000000000000000000000000000600082015250565b7f496e76616c696420746f6b656e00000000000000000000000000000000000000600082015250565b7f496e76616c696420726174650000000000000000000000000000000000000000600082015250565b62000576816200049e565b81146200058257600080fd5b50565b6200059081620004b2565b81146200059c57600080fd5b50565b620005aa81620004e6565b8114620005b657600080fd5b50565b6117cf80620005c96000396000f3fe6080604052600436106100c65760003560e01c8063715018a61161007f57806398b9a2dc1161005957806398b9a2dc14610227578063ec8ac4d814610250578063f2fde38b1461026c578063fc0c546a14610295576100dd565b8063715018a6146101ce5780638456cb59146101e55780638da5cb5b146101fc576100dd565b80632c4e722e146100e257806339df43ff1461010d5780633f4ba83a146101365780634042b66f1461014d578063521eb273146101785780635c975abb146101a3576100dd565b366100dd576100db6100d66102c0565b6102c8565b005b600080fd5b3480156100ee57600080fd5b506100f7610492565b6040516101049190611387565b60405180910390f35b34801561011957600080fd5b50610134600480360381019061012f9190610fdc565b61049c565b005b34801561014257600080fd5b5061014b610690565b005b34801561015957600080fd5b5061016261075d565b60405161016f9190611387565b60405180910390f35b34801561018457600080fd5b5061018d610767565b60405161019a91906111e4565b60405180910390f35b3480156101af57600080fd5b506101b8610791565b6040516101c59190611251565b60405180910390f35b3480156101da57600080fd5b506101e36107a8565b005b3480156101f157600080fd5b506101fa610830565b005b34801561020857600080fd5b506102116108fe565b60405161021e91906111c9565b60405180910390f35b34801561023357600080fd5b5061024e60048036038101906102499190610fdc565b610927565b005b61026a60048036038101906102659190610fb3565b6102c8565b005b34801561027857600080fd5b50610293600480360381019061028e9190610fb3565b6109e7565b005b3480156102a157600080fd5b506102aa610adf565b6040516102b7919061126c565b60405180910390f35b600033905090565b6002600154141561030e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161030590611367565b60405180910390fd5b600260018190555061031e610791565b1561035e576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610355906112e7565b60405180910390fd5b633b9aca003410156103a5576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161039c906112c7565b60405180910390fd5b60006103be633b9aca0034610b0990919063ffffffff16565b90506103ca8282610b1f565b60006103d582610bd7565b90506103ec82600554610bf590919063ffffffff16565b6005819055506103fc8382610c0b565b8273ffffffffffffffffffffffffffffffffffffffff1661041b6102c0565b73ffffffffffffffffffffffffffffffffffffffff167f6faf93231a456e552dbc9961f58d9713ee4f2e69d15f1975b050ef0911053a7b84846040516104629291906113a2565b60405180910390a36104748383610c19565b61047c610c1d565b6104868383610c88565b50506001808190555050565b6000600454905090565b6104a46102c0565b73ffffffffffffffffffffffffffffffffffffffff166104c26108fe565b73ffffffffffffffffffffffffffffffffffffffff1614610518576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161050f90611327565b60405180910390fd5b6000600260019054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166370a08231306040518263ffffffff1660e01b815260040161057591906111c9565b60206040518083038186803b15801561058d57600080fd5b505afa1580156105a1573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906105c5919061102e565b9050600260019054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663a9059cbb83836040518363ffffffff1660e01b81526004016106249291906111ff565b602060405180830381600087803b15801561063e57600080fd5b505af1158015610652573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906106769190611005565b508173ffffffffffffffffffffffffffffffffffffffff16ff5b610698610791565b6106d7576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016106ce90611287565b60405180910390fd5b6106df6102c0565b73ffffffffffffffffffffffffffffffffffffffff166106fd6108fe565b73ffffffffffffffffffffffffffffffffffffffff1614610753576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161074a90611327565b60405180910390fd5b61075b610c8c565b565b6000600554905090565b6000600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b6000600260009054906101000a900460ff16905090565b6107b06102c0565b73ffffffffffffffffffffffffffffffffffffffff166107ce6108fe565b73ffffffffffffffffffffffffffffffffffffffff1614610824576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161081b90611327565b60405180910390fd5b61082e6000610d2e565b565b610838610791565b15610878576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161086f906112e7565b60405180910390fd5b6108806102c0565b73ffffffffffffffffffffffffffffffffffffffff1661089e6108fe565b73ffffffffffffffffffffffffffffffffffffffff16146108f4576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016108eb90611327565b60405180910390fd5b6108fc610df2565b565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b61092f6102c0565b73ffffffffffffffffffffffffffffffffffffffff1661094d6108fe565b73ffffffffffffffffffffffffffffffffffffffff16146109a3576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161099a90611327565b60405180910390fd5b80600360006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b6109ef6102c0565b73ffffffffffffffffffffffffffffffffffffffff16610a0d6108fe565b73ffffffffffffffffffffffffffffffffffffffff1614610a63576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610a5a90611327565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff161415610ad3576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610aca906112a7565b60405180910390fd5b610adc81610d2e565b50565b6000600260019054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b60008183610b179190611432565b905092915050565b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff161415610b8f576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610b8690611307565b60405180910390fd5b6000811415610bd3576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610bca90611347565b60405180910390fd5b5050565b6000610bee60045483610e9590919063ffffffff16565b9050919050565b60008183610c0391906113dc565b905092915050565b610c158282610eab565b5050565b5050565b600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166108fc349081150290604051600060405180830381858888f19350505050158015610c85573d6000803e3d6000fd5b50565b5050565b610c94610791565b610cd3576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610cca90611287565b60405180910390fd5b6000600260006101000a81548160ff0219169083151502179055507f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa610d176102c0565b604051610d2491906111c9565b60405180910390a1565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050816000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b610dfa610791565b15610e3a576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610e31906112e7565b60405180910390fd5b6001600260006101000a81548160ff0219169083151502179055507f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258610e7e6102c0565b604051610e8b91906111c9565b60405180910390a1565b60008183610ea39190611463565b905092915050565b600260019054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663a9059cbb83836040518363ffffffff1660e01b8152600401610f08929190611228565b602060405180830381600087803b158015610f2257600080fd5b505af1158015610f36573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610f5a9190611005565b505050565b600081359050610f6e8161173d565b92915050565b600081359050610f8381611754565b92915050565b600081519050610f988161176b565b92915050565b600081519050610fad81611782565b92915050565b600060208284031215610fc557600080fd5b6000610fd384828501610f5f565b91505092915050565b600060208284031215610fee57600080fd5b6000610ffc84828501610f74565b91505092915050565b60006020828403121561101757600080fd5b600061102584828501610f89565b91505092915050565b60006020828403121561104057600080fd5b600061104e84828501610f9e565b91505092915050565b61106081611517565b82525050565b61106f816114cf565b82525050565b61107e816114bd565b82525050565b61108d816114e1565b82525050565b61109c81611529565b82525050565b60006110af6014836113cb565b91506110ba826115cf565b602082019050919050565b60006110d26026836113cb565b91506110dd826115f8565b604082019050919050565b60006110f5600e836113cb565b915061110082611647565b602082019050919050565b60006111186010836113cb565b915061112382611670565b602082019050919050565b600061113b6013836113cb565b915061114682611699565b602082019050919050565b600061115e6020836113cb565b9150611169826116c2565b602082019050919050565b60006111816007836113cb565b915061118c826116eb565b602082019050919050565b60006111a4601f836113cb565b91506111af82611714565b602082019050919050565b6111c38161150d565b82525050565b60006020820190506111de6000830184611075565b92915050565b60006020820190506111f96000830184611066565b92915050565b60006040820190506112146000830185611057565b61122160208301846111ba565b9392505050565b600060408201905061123d6000830185611075565b61124a60208301846111ba565b9392505050565b60006020820190506112666000830184611084565b92915050565b60006020820190506112816000830184611093565b92915050565b600060208201905081810360008301526112a0816110a2565b9050919050565b600060208201905081810360008301526112c0816110c5565b9050919050565b600060208201905081810360008301526112e0816110e8565b9050919050565b600060208201905081810360008301526113008161110b565b9050919050565b600060208201905081810360008301526113208161112e565b9050919050565b6000602082019050818103600083015261134081611151565b9050919050565b6000602082019050818103600083015261136081611174565b9050919050565b6000602082019050818103600083015261138081611197565b9050919050565b600060208201905061139c60008301846111ba565b92915050565b60006040820190506113b760008301856111ba565b6113c460208301846111ba565b9392505050565b600082825260208201905092915050565b60006113e78261150d565b91506113f28361150d565b9250827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0382111561142757611426611571565b5b828201905092915050565b600061143d8261150d565b91506114488361150d565b925082611458576114576115a0565b5b828204905092915050565b600061146e8261150d565b91506114798361150d565b9250817fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff04831182151516156114b2576114b1611571565b5b828202905092915050565b60006114c8826114ed565b9050919050565b60006114da826114ed565b9050919050565b60008115159050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b60006115228261154d565b9050919050565b60006115348261153b565b9050919050565b6000611546826114ed565b9050919050565b60006115588261155f565b9050919050565b600061156a826114ed565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b7f5061757361626c653a206e6f7420706175736564000000000000000000000000600082015250565b7f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160008201527f6464726573730000000000000000000000000000000000000000000000000000602082015250565b7f496e76616c696420616d6f756e74000000000000000000000000000000000000600082015250565b7f5061757361626c653a2070617573656400000000000000000000000000000000600082015250565b7f496e76616c69642062656e656669636961727900000000000000000000000000600082015250565b7f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572600082015250565b7f496e76616c696400000000000000000000000000000000000000000000000000600082015250565b7f5265656e7472616e637947756172643a207265656e7472616e742063616c6c00600082015250565b611746816114bd565b811461175157600080fd5b50565b61175d816114cf565b811461176857600080fd5b50565b611774816114e1565b811461177f57600080fd5b50565b61178b8161150d565b811461179657600080fd5b5056fea26469706673582212204c4e46c62ee4ce318705fad83999b7e0bb9e3aa0c0f569050ff540a6a147f47664736f6c63430008040033",
}

// CrowdsaleABI is the input ABI used to generate the binding from.
// Deprecated: Use CrowdsaleMetaData.ABI instead.
var CrowdsaleABI = CrowdsaleMetaData.ABI

// CrowdsaleBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use CrowdsaleMetaData.Bin instead.
var CrowdsaleBin = CrowdsaleMetaData.Bin

// DeployCrowdsale deploys a new Ethereum contract, binding an instance of Crowdsale to it.
func DeployCrowdsale(auth *bind.TransactOpts, backend bind.ContractBackend, __rate *big.Int, __wallet common.Address, __token common.Address) (common.Address, *types.Transaction, *Crowdsale, error) {
	parsed, err := CrowdsaleMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(CrowdsaleBin), backend, __rate, __wallet, __token)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Crowdsale{CrowdsaleCaller: CrowdsaleCaller{contract: contract}, CrowdsaleTransactor: CrowdsaleTransactor{contract: contract}, CrowdsaleFilterer: CrowdsaleFilterer{contract: contract}}, nil
}

// Crowdsale is an auto generated Go binding around an Ethereum contract.
type Crowdsale struct {
	CrowdsaleCaller     // Read-only binding to the contract
	CrowdsaleTransactor // Write-only binding to the contract
	CrowdsaleFilterer   // Log filterer for contract events
}

// CrowdsaleCaller is an auto generated read-only Go binding around an Ethereum contract.
type CrowdsaleCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CrowdsaleTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CrowdsaleTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CrowdsaleFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CrowdsaleFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CrowdsaleSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CrowdsaleSession struct {
	Contract     *Crowdsale        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CrowdsaleCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CrowdsaleCallerSession struct {
	Contract *CrowdsaleCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// CrowdsaleTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CrowdsaleTransactorSession struct {
	Contract     *CrowdsaleTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// CrowdsaleRaw is an auto generated low-level Go binding around an Ethereum contract.
type CrowdsaleRaw struct {
	Contract *Crowdsale // Generic contract binding to access the raw methods on
}

// CrowdsaleCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CrowdsaleCallerRaw struct {
	Contract *CrowdsaleCaller // Generic read-only contract binding to access the raw methods on
}

// CrowdsaleTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CrowdsaleTransactorRaw struct {
	Contract *CrowdsaleTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCrowdsale creates a new instance of Crowdsale, bound to a specific deployed contract.
func NewCrowdsale(address common.Address, backend bind.ContractBackend) (*Crowdsale, error) {
	contract, err := bindCrowdsale(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Crowdsale{CrowdsaleCaller: CrowdsaleCaller{contract: contract}, CrowdsaleTransactor: CrowdsaleTransactor{contract: contract}, CrowdsaleFilterer: CrowdsaleFilterer{contract: contract}}, nil
}

// NewCrowdsaleCaller creates a new read-only instance of Crowdsale, bound to a specific deployed contract.
func NewCrowdsaleCaller(address common.Address, caller bind.ContractCaller) (*CrowdsaleCaller, error) {
	contract, err := bindCrowdsale(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CrowdsaleCaller{contract: contract}, nil
}

// NewCrowdsaleTransactor creates a new write-only instance of Crowdsale, bound to a specific deployed contract.
func NewCrowdsaleTransactor(address common.Address, transactor bind.ContractTransactor) (*CrowdsaleTransactor, error) {
	contract, err := bindCrowdsale(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CrowdsaleTransactor{contract: contract}, nil
}

// NewCrowdsaleFilterer creates a new log filterer instance of Crowdsale, bound to a specific deployed contract.
func NewCrowdsaleFilterer(address common.Address, filterer bind.ContractFilterer) (*CrowdsaleFilterer, error) {
	contract, err := bindCrowdsale(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CrowdsaleFilterer{contract: contract}, nil
}

// bindCrowdsale binds a generic wrapper to an already deployed contract.
func bindCrowdsale(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(CrowdsaleABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Crowdsale *CrowdsaleRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Crowdsale.Contract.CrowdsaleCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Crowdsale *CrowdsaleRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Crowdsale.Contract.CrowdsaleTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Crowdsale *CrowdsaleRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Crowdsale.Contract.CrowdsaleTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Crowdsale *CrowdsaleCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Crowdsale.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Crowdsale *CrowdsaleTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Crowdsale.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Crowdsale *CrowdsaleTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Crowdsale.Contract.contract.Transact(opts, method, params...)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Crowdsale *CrowdsaleCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Crowdsale.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Crowdsale *CrowdsaleSession) Owner() (common.Address, error) {
	return _Crowdsale.Contract.Owner(&_Crowdsale.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Crowdsale *CrowdsaleCallerSession) Owner() (common.Address, error) {
	return _Crowdsale.Contract.Owner(&_Crowdsale.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Crowdsale *CrowdsaleCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Crowdsale.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Crowdsale *CrowdsaleSession) Paused() (bool, error) {
	return _Crowdsale.Contract.Paused(&_Crowdsale.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Crowdsale *CrowdsaleCallerSession) Paused() (bool, error) {
	return _Crowdsale.Contract.Paused(&_Crowdsale.CallOpts)
}

// Rate is a free data retrieval call binding the contract method 0x2c4e722e.
//
// Solidity: function rate() view returns(uint256)
func (_Crowdsale *CrowdsaleCaller) Rate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Crowdsale.contract.Call(opts, &out, "rate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Rate is a free data retrieval call binding the contract method 0x2c4e722e.
//
// Solidity: function rate() view returns(uint256)
func (_Crowdsale *CrowdsaleSession) Rate() (*big.Int, error) {
	return _Crowdsale.Contract.Rate(&_Crowdsale.CallOpts)
}

// Rate is a free data retrieval call binding the contract method 0x2c4e722e.
//
// Solidity: function rate() view returns(uint256)
func (_Crowdsale *CrowdsaleCallerSession) Rate() (*big.Int, error) {
	return _Crowdsale.Contract.Rate(&_Crowdsale.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_Crowdsale *CrowdsaleCaller) Token(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Crowdsale.contract.Call(opts, &out, "token")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_Crowdsale *CrowdsaleSession) Token() (common.Address, error) {
	return _Crowdsale.Contract.Token(&_Crowdsale.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_Crowdsale *CrowdsaleCallerSession) Token() (common.Address, error) {
	return _Crowdsale.Contract.Token(&_Crowdsale.CallOpts)
}

// Wallet is a free data retrieval call binding the contract method 0x521eb273.
//
// Solidity: function wallet() view returns(address)
func (_Crowdsale *CrowdsaleCaller) Wallet(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Crowdsale.contract.Call(opts, &out, "wallet")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Wallet is a free data retrieval call binding the contract method 0x521eb273.
//
// Solidity: function wallet() view returns(address)
func (_Crowdsale *CrowdsaleSession) Wallet() (common.Address, error) {
	return _Crowdsale.Contract.Wallet(&_Crowdsale.CallOpts)
}

// Wallet is a free data retrieval call binding the contract method 0x521eb273.
//
// Solidity: function wallet() view returns(address)
func (_Crowdsale *CrowdsaleCallerSession) Wallet() (common.Address, error) {
	return _Crowdsale.Contract.Wallet(&_Crowdsale.CallOpts)
}

// WeiRaised is a free data retrieval call binding the contract method 0x4042b66f.
//
// Solidity: function weiRaised() view returns(uint256)
func (_Crowdsale *CrowdsaleCaller) WeiRaised(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Crowdsale.contract.Call(opts, &out, "weiRaised")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WeiRaised is a free data retrieval call binding the contract method 0x4042b66f.
//
// Solidity: function weiRaised() view returns(uint256)
func (_Crowdsale *CrowdsaleSession) WeiRaised() (*big.Int, error) {
	return _Crowdsale.Contract.WeiRaised(&_Crowdsale.CallOpts)
}

// WeiRaised is a free data retrieval call binding the contract method 0x4042b66f.
//
// Solidity: function weiRaised() view returns(uint256)
func (_Crowdsale *CrowdsaleCallerSession) WeiRaised() (*big.Int, error) {
	return _Crowdsale.Contract.WeiRaised(&_Crowdsale.CallOpts)
}

// BuyTokens is a paid mutator transaction binding the contract method 0xec8ac4d8.
//
// Solidity: function buyTokens(address beneficiary) payable returns()
func (_Crowdsale *CrowdsaleTransactor) BuyTokens(opts *bind.TransactOpts, beneficiary common.Address) (*types.Transaction, error) {
	return _Crowdsale.contract.Transact(opts, "buyTokens", beneficiary)
}

// BuyTokens is a paid mutator transaction binding the contract method 0xec8ac4d8.
//
// Solidity: function buyTokens(address beneficiary) payable returns()
func (_Crowdsale *CrowdsaleSession) BuyTokens(beneficiary common.Address) (*types.Transaction, error) {
	return _Crowdsale.Contract.BuyTokens(&_Crowdsale.TransactOpts, beneficiary)
}

// BuyTokens is a paid mutator transaction binding the contract method 0xec8ac4d8.
//
// Solidity: function buyTokens(address beneficiary) payable returns()
func (_Crowdsale *CrowdsaleTransactorSession) BuyTokens(beneficiary common.Address) (*types.Transaction, error) {
	return _Crowdsale.Contract.BuyTokens(&_Crowdsale.TransactOpts, beneficiary)
}

// ChangeWallet is a paid mutator transaction binding the contract method 0x98b9a2dc.
//
// Solidity: function changeWallet(address newReceiver) returns()
func (_Crowdsale *CrowdsaleTransactor) ChangeWallet(opts *bind.TransactOpts, newReceiver common.Address) (*types.Transaction, error) {
	return _Crowdsale.contract.Transact(opts, "changeWallet", newReceiver)
}

// ChangeWallet is a paid mutator transaction binding the contract method 0x98b9a2dc.
//
// Solidity: function changeWallet(address newReceiver) returns()
func (_Crowdsale *CrowdsaleSession) ChangeWallet(newReceiver common.Address) (*types.Transaction, error) {
	return _Crowdsale.Contract.ChangeWallet(&_Crowdsale.TransactOpts, newReceiver)
}

// ChangeWallet is a paid mutator transaction binding the contract method 0x98b9a2dc.
//
// Solidity: function changeWallet(address newReceiver) returns()
func (_Crowdsale *CrowdsaleTransactorSession) ChangeWallet(newReceiver common.Address) (*types.Transaction, error) {
	return _Crowdsale.Contract.ChangeWallet(&_Crowdsale.TransactOpts, newReceiver)
}

// DestroySmartContract is a paid mutator transaction binding the contract method 0x39df43ff.
//
// Solidity: function destroySmartContract(address _to) returns()
func (_Crowdsale *CrowdsaleTransactor) DestroySmartContract(opts *bind.TransactOpts, _to common.Address) (*types.Transaction, error) {
	return _Crowdsale.contract.Transact(opts, "destroySmartContract", _to)
}

// DestroySmartContract is a paid mutator transaction binding the contract method 0x39df43ff.
//
// Solidity: function destroySmartContract(address _to) returns()
func (_Crowdsale *CrowdsaleSession) DestroySmartContract(_to common.Address) (*types.Transaction, error) {
	return _Crowdsale.Contract.DestroySmartContract(&_Crowdsale.TransactOpts, _to)
}

// DestroySmartContract is a paid mutator transaction binding the contract method 0x39df43ff.
//
// Solidity: function destroySmartContract(address _to) returns()
func (_Crowdsale *CrowdsaleTransactorSession) DestroySmartContract(_to common.Address) (*types.Transaction, error) {
	return _Crowdsale.Contract.DestroySmartContract(&_Crowdsale.TransactOpts, _to)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Crowdsale *CrowdsaleTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Crowdsale.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Crowdsale *CrowdsaleSession) Pause() (*types.Transaction, error) {
	return _Crowdsale.Contract.Pause(&_Crowdsale.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Crowdsale *CrowdsaleTransactorSession) Pause() (*types.Transaction, error) {
	return _Crowdsale.Contract.Pause(&_Crowdsale.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Crowdsale *CrowdsaleTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Crowdsale.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Crowdsale *CrowdsaleSession) RenounceOwnership() (*types.Transaction, error) {
	return _Crowdsale.Contract.RenounceOwnership(&_Crowdsale.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Crowdsale *CrowdsaleTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Crowdsale.Contract.RenounceOwnership(&_Crowdsale.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Crowdsale *CrowdsaleTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Crowdsale.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Crowdsale *CrowdsaleSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Crowdsale.Contract.TransferOwnership(&_Crowdsale.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Crowdsale *CrowdsaleTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Crowdsale.Contract.TransferOwnership(&_Crowdsale.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Crowdsale *CrowdsaleTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Crowdsale.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Crowdsale *CrowdsaleSession) Unpause() (*types.Transaction, error) {
	return _Crowdsale.Contract.Unpause(&_Crowdsale.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Crowdsale *CrowdsaleTransactorSession) Unpause() (*types.Transaction, error) {
	return _Crowdsale.Contract.Unpause(&_Crowdsale.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Crowdsale *CrowdsaleTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Crowdsale.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Crowdsale *CrowdsaleSession) Receive() (*types.Transaction, error) {
	return _Crowdsale.Contract.Receive(&_Crowdsale.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Crowdsale *CrowdsaleTransactorSession) Receive() (*types.Transaction, error) {
	return _Crowdsale.Contract.Receive(&_Crowdsale.TransactOpts)
}

// CrowdsaleOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Crowdsale contract.
type CrowdsaleOwnershipTransferredIterator struct {
	Event *CrowdsaleOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *CrowdsaleOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrowdsaleOwnershipTransferred)
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
		it.Event = new(CrowdsaleOwnershipTransferred)
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
func (it *CrowdsaleOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrowdsaleOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrowdsaleOwnershipTransferred represents a OwnershipTransferred event raised by the Crowdsale contract.
type CrowdsaleOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Crowdsale *CrowdsaleFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*CrowdsaleOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Crowdsale.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &CrowdsaleOwnershipTransferredIterator{contract: _Crowdsale.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Crowdsale *CrowdsaleFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *CrowdsaleOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Crowdsale.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrowdsaleOwnershipTransferred)
				if err := _Crowdsale.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Crowdsale *CrowdsaleFilterer) ParseOwnershipTransferred(log types.Log) (*CrowdsaleOwnershipTransferred, error) {
	event := new(CrowdsaleOwnershipTransferred)
	if err := _Crowdsale.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrowdsalePausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the Crowdsale contract.
type CrowdsalePausedIterator struct {
	Event *CrowdsalePaused // Event containing the contract specifics and raw log

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
func (it *CrowdsalePausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrowdsalePaused)
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
		it.Event = new(CrowdsalePaused)
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
func (it *CrowdsalePausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrowdsalePausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrowdsalePaused represents a Paused event raised by the Crowdsale contract.
type CrowdsalePaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Crowdsale *CrowdsaleFilterer) FilterPaused(opts *bind.FilterOpts) (*CrowdsalePausedIterator, error) {

	logs, sub, err := _Crowdsale.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &CrowdsalePausedIterator{contract: _Crowdsale.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Crowdsale *CrowdsaleFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *CrowdsalePaused) (event.Subscription, error) {

	logs, sub, err := _Crowdsale.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrowdsalePaused)
				if err := _Crowdsale.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_Crowdsale *CrowdsaleFilterer) ParsePaused(log types.Log) (*CrowdsalePaused, error) {
	event := new(CrowdsalePaused)
	if err := _Crowdsale.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrowdsaleTokensPurchasedIterator is returned from FilterTokensPurchased and is used to iterate over the raw logs and unpacked data for TokensPurchased events raised by the Crowdsale contract.
type CrowdsaleTokensPurchasedIterator struct {
	Event *CrowdsaleTokensPurchased // Event containing the contract specifics and raw log

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
func (it *CrowdsaleTokensPurchasedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrowdsaleTokensPurchased)
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
		it.Event = new(CrowdsaleTokensPurchased)
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
func (it *CrowdsaleTokensPurchasedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrowdsaleTokensPurchasedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrowdsaleTokensPurchased represents a TokensPurchased event raised by the Crowdsale contract.
type CrowdsaleTokensPurchased struct {
	Purchaser   common.Address
	Beneficiary common.Address
	Value       *big.Int
	Amount      *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterTokensPurchased is a free log retrieval operation binding the contract event 0x6faf93231a456e552dbc9961f58d9713ee4f2e69d15f1975b050ef0911053a7b.
//
// Solidity: event TokensPurchased(address indexed purchaser, address indexed beneficiary, uint256 value, uint256 amount)
func (_Crowdsale *CrowdsaleFilterer) FilterTokensPurchased(opts *bind.FilterOpts, purchaser []common.Address, beneficiary []common.Address) (*CrowdsaleTokensPurchasedIterator, error) {

	var purchaserRule []interface{}
	for _, purchaserItem := range purchaser {
		purchaserRule = append(purchaserRule, purchaserItem)
	}
	var beneficiaryRule []interface{}
	for _, beneficiaryItem := range beneficiary {
		beneficiaryRule = append(beneficiaryRule, beneficiaryItem)
	}

	logs, sub, err := _Crowdsale.contract.FilterLogs(opts, "TokensPurchased", purchaserRule, beneficiaryRule)
	if err != nil {
		return nil, err
	}
	return &CrowdsaleTokensPurchasedIterator{contract: _Crowdsale.contract, event: "TokensPurchased", logs: logs, sub: sub}, nil
}

// WatchTokensPurchased is a free log subscription operation binding the contract event 0x6faf93231a456e552dbc9961f58d9713ee4f2e69d15f1975b050ef0911053a7b.
//
// Solidity: event TokensPurchased(address indexed purchaser, address indexed beneficiary, uint256 value, uint256 amount)
func (_Crowdsale *CrowdsaleFilterer) WatchTokensPurchased(opts *bind.WatchOpts, sink chan<- *CrowdsaleTokensPurchased, purchaser []common.Address, beneficiary []common.Address) (event.Subscription, error) {

	var purchaserRule []interface{}
	for _, purchaserItem := range purchaser {
		purchaserRule = append(purchaserRule, purchaserItem)
	}
	var beneficiaryRule []interface{}
	for _, beneficiaryItem := range beneficiary {
		beneficiaryRule = append(beneficiaryRule, beneficiaryItem)
	}

	logs, sub, err := _Crowdsale.contract.WatchLogs(opts, "TokensPurchased", purchaserRule, beneficiaryRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrowdsaleTokensPurchased)
				if err := _Crowdsale.contract.UnpackLog(event, "TokensPurchased", log); err != nil {
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
func (_Crowdsale *CrowdsaleFilterer) ParseTokensPurchased(log types.Log) (*CrowdsaleTokensPurchased, error) {
	event := new(CrowdsaleTokensPurchased)
	if err := _Crowdsale.contract.UnpackLog(event, "TokensPurchased", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrowdsaleUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the Crowdsale contract.
type CrowdsaleUnpausedIterator struct {
	Event *CrowdsaleUnpaused // Event containing the contract specifics and raw log

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
func (it *CrowdsaleUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrowdsaleUnpaused)
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
		it.Event = new(CrowdsaleUnpaused)
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
func (it *CrowdsaleUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrowdsaleUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrowdsaleUnpaused represents a Unpaused event raised by the Crowdsale contract.
type CrowdsaleUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Crowdsale *CrowdsaleFilterer) FilterUnpaused(opts *bind.FilterOpts) (*CrowdsaleUnpausedIterator, error) {

	logs, sub, err := _Crowdsale.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &CrowdsaleUnpausedIterator{contract: _Crowdsale.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Crowdsale *CrowdsaleFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *CrowdsaleUnpaused) (event.Subscription, error) {

	logs, sub, err := _Crowdsale.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrowdsaleUnpaused)
				if err := _Crowdsale.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
func (_Crowdsale *CrowdsaleFilterer) ParseUnpaused(log types.Log) (*CrowdsaleUnpaused, error) {
	event := new(CrowdsaleUnpaused)
	if err := _Crowdsale.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
