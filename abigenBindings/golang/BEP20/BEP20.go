// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package BEP20

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

// BEP20MetaData contains all meta data concerning the BEP20 contract.
var BEP20MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"tokenName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"tokenSymbol\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"tokenDecimals\",\"type\":\"uint8\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60806040523480156200001157600080fd5b5060405162001ebe38038062001ebe8339818101604052810190620000379190620002b2565b620000576200004b620000ad60201b60201c565b620000b560201b60201c565b82600690805190602001906200006f92919062000179565b5081600590805190602001906200008892919062000179565b5080600460006101000a81548160ff021916908360ff160217905550505050620004d1565b600033905090565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050816000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b8280546200018790620003dc565b90600052602060002090601f016020900481019282620001ab5760008555620001f7565b82601f10620001c657805160ff1916838001178555620001f7565b82800160010185558215620001f7579182015b82811115620001f6578251825591602001919060010190620001d9565b5b5090506200020691906200020a565b5090565b5b80821115620002255760008160009055506001016200020b565b5090565b6000620002406200023a8462000363565b6200033a565b9050828152602081018484840111156200025957600080fd5b62000266848285620003a6565b509392505050565b600082601f8301126200028057600080fd5b81516200029284826020860162000229565b91505092915050565b600081519050620002ac81620004b7565b92915050565b600080600060608486031215620002c857600080fd5b600084015167ffffffffffffffff811115620002e357600080fd5b620002f1868287016200026e565b935050602084015167ffffffffffffffff8111156200030f57600080fd5b6200031d868287016200026e565b925050604062000330868287016200029b565b9150509250925092565b60006200034662000359565b905062000354828262000412565b919050565b6000604051905090565b600067ffffffffffffffff82111562000381576200038062000477565b5b6200038c82620004a6565b9050602081019050919050565b600060ff82169050919050565b60005b83811015620003c6578082015181840152602081019050620003a9565b83811115620003d6576000848401525b50505050565b60006002820490506001821680620003f557607f821691505b602082108114156200040c576200040b62000448565b5b50919050565b6200041d82620004a6565b810181811067ffffffffffffffff821117156200043f576200043e62000477565b5b80604052505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6000601f19601f8301169050919050565b620004c28162000399565b8114620004ce57600080fd5b50565b6119dd80620004e16000396000f3fe608060405234801561001057600080fd5b50600436106101005760003560e01c8063893d20e811610097578063a457c2d711610066578063a457c2d7146102b3578063a9059cbb146102e3578063dd62ed3e14610313578063f2fde38b1461034357610100565b8063893d20e8146102295780638da5cb5b1461024757806395d89b4114610265578063a0712d681461028357610100565b8063313ce567116100d3578063313ce567146101a157806339509351146101bf57806370a08231146101ef578063715018a61461021f57610100565b806306fdde0314610105578063095ea7b31461012357806318160ddd1461015357806323b872dd14610171575b600080fd5b61010d61035f565b60405161011a9190611456565b60405180910390f35b61013d60048036038101906101389190611251565b6103f1565b60405161014a919061143b565b60405180910390f35b61015b61040f565b6040516101689190611558565b60405180910390f35b61018b60048036038101906101869190611202565b610419565b604051610198919061143b565b60405180910390f35b6101a96104f2565b6040516101b69190611573565b60405180910390f35b6101d960048036038101906101d49190611251565b610509565b6040516101e6919061143b565b60405180910390f35b6102096004803603810190610204919061119d565b6105bc565b6040516102169190611558565b60405180910390f35b610227610605565b005b61023161068d565b60405161023e9190611420565b60405180910390f35b61024f61069c565b60405161025c9190611420565b60405180910390f35b61026d6106c5565b60405161027a9190611456565b60405180910390f35b61029d6004803603810190610298919061128d565b610757565b6040516102aa919061143b565b60405180910390f35b6102cd60048036038101906102c89190611251565b6107ef565b6040516102da919061143b565b60405180910390f35b6102fd60048036038101906102f89190611251565b6108bc565b60405161030a919061143b565b60405180910390f35b61032d600480360381019061032891906111c6565b6108da565b60405161033a9190611558565b60405180910390f35b61035d6004803603810190610358919061119d565b610961565b005b60606006805461036e90611688565b80601f016020809104026020016040519081016040528092919081815260200182805461039a90611688565b80156103e75780601f106103bc576101008083540402835291602001916103e7565b820191906000526020600020905b8154815290600101906020018083116103ca57829003601f168201915b5050505050905090565b60006104056103fe610a59565b8484610a61565b6001905092915050565b6000600354905090565b6000610426848484610c2c565b6104e784610432610a59565b6104e28560405180606001604052806028815260200161193560289139600260008b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000610498610a59565b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054610eba9092919063ffffffff16565b610a61565b600190509392505050565b6000600460009054906101000a900460ff16905090565b60006105b2610516610a59565b846105ad8560026000610527610a59565b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008973ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054610f0f90919063ffffffff16565b610a61565b6001905092915050565b6000600160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020549050919050565b61060d610a59565b73ffffffffffffffffffffffffffffffffffffffff1661062b61069c565b73ffffffffffffffffffffffffffffffffffffffff1614610681576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610678906114f8565b60405180910390fd5b61068b6000610f25565b565b600061069761069c565b905090565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b6060600580546106d490611688565b80601f016020809104026020016040519081016040528092919081815260200182805461070090611688565b801561074d5780601f106107225761010080835404028352916020019161074d565b820191906000526020600020905b81548152906001019060200180831161073057829003601f168201915b5050505050905090565b6000610761610a59565b73ffffffffffffffffffffffffffffffffffffffff1661077f61069c565b73ffffffffffffffffffffffffffffffffffffffff16146107d5576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016107cc906114f8565b60405180910390fd5b6107e66107e0610a59565b83610fe9565b60019050919050565b60006108b26107fc610a59565b846108ad856040518060600160405280602581526020016119836025913960026000610826610a59565b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008a73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054610eba9092919063ffffffff16565b610a61565b6001905092915050565b60006108d06108c9610a59565b8484610c2c565b6001905092915050565b6000600260008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054905092915050565b610969610a59565b73ffffffffffffffffffffffffffffffffffffffff1661098761069c565b73ffffffffffffffffffffffffffffffffffffffff16146109dd576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016109d4906114f8565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff161415610a4d576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610a44906114b8565b60405180910390fd5b610a5681610f25565b50565b600033905090565b600073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff161415610ad1576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610ac890611498565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff161415610b41576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610b3890611538565b60405180910390fd5b80600260008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055508173ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92583604051610c1f9190611558565b60405180910390a3505050565b600073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff161415610c9c576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610c9390611478565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff161415610d0c576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610d0390611518565b60405180910390fd5b610d788160405180606001604052806026815260200161195d60269139600160008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054610eba9092919063ffffffff16565b600160008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002081905550610e0d81600160008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054610f0f90919063ffffffff16565b600160008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055508173ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef83604051610ead9190611558565b60405180910390a3505050565b6000838311158290610f02576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610ef99190611456565b60405180910390fd5b5082840390509392505050565b60008183610f1d91906115aa565b905092915050565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050816000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff161415611059576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611050906114d8565b60405180910390fd5b61106e81600354610f0f90919063ffffffff16565b6003819055506110c681600160008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054610f0f90919063ffffffff16565b600160008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055508173ffffffffffffffffffffffffffffffffffffffff16600073ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef836040516111679190611558565b60405180910390a35050565b60008135905061118281611906565b92915050565b6000813590506111978161191d565b92915050565b6000602082840312156111af57600080fd5b60006111bd84828501611173565b91505092915050565b600080604083850312156111d957600080fd5b60006111e785828601611173565b92505060206111f885828601611173565b9150509250929050565b60008060006060848603121561121757600080fd5b600061122586828701611173565b935050602061123686828701611173565b925050604061124786828701611188565b9150509250925092565b6000806040838503121561126457600080fd5b600061127285828601611173565b925050602061128385828601611188565b9150509250929050565b60006020828403121561129f57600080fd5b60006112ad84828501611188565b91505092915050565b6112bf81611600565b82525050565b6112ce81611612565b82525050565b60006112df8261158e565b6112e98185611599565b93506112f9818560208601611655565b61130281611718565b840191505092915050565b600061131a602583611599565b915061132582611729565b604082019050919050565b600061133d602483611599565b915061134882611778565b604082019050919050565b6000611360602683611599565b915061136b826117c7565b604082019050919050565b6000611383601f83611599565b915061138e82611816565b602082019050919050565b60006113a6602083611599565b91506113b18261183f565b602082019050919050565b60006113c9602383611599565b91506113d482611868565b604082019050919050565b60006113ec602283611599565b91506113f7826118b7565b604082019050919050565b61140b8161163e565b82525050565b61141a81611648565b82525050565b600060208201905061143560008301846112b6565b92915050565b600060208201905061145060008301846112c5565b92915050565b6000602082019050818103600083015261147081846112d4565b905092915050565b600060208201905081810360008301526114918161130d565b9050919050565b600060208201905081810360008301526114b181611330565b9050919050565b600060208201905081810360008301526114d181611353565b9050919050565b600060208201905081810360008301526114f181611376565b9050919050565b6000602082019050818103600083015261151181611399565b9050919050565b60006020820190508181036000830152611531816113bc565b9050919050565b60006020820190508181036000830152611551816113df565b9050919050565b600060208201905061156d6000830184611402565b92915050565b60006020820190506115886000830184611411565b92915050565b600081519050919050565b600082825260208201905092915050565b60006115b58261163e565b91506115c08361163e565b9250827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff038211156115f5576115f46116ba565b5b828201905092915050565b600061160b8261161e565b9050919050565b60008115159050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b600060ff82169050919050565b60005b83811015611673578082015181840152602081019050611658565b83811115611682576000848401525b50505050565b600060028204905060018216806116a057607f821691505b602082108114156116b4576116b36116e9565b5b50919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b6000601f19601f8301169050919050565b7f42455032303a207472616e736665722066726f6d20746865207a65726f20616460008201527f6472657373000000000000000000000000000000000000000000000000000000602082015250565b7f42455032303a20617070726f76652066726f6d20746865207a65726f2061646460008201527f7265737300000000000000000000000000000000000000000000000000000000602082015250565b7f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160008201527f6464726573730000000000000000000000000000000000000000000000000000602082015250565b7f42455032303a206d696e7420746f20746865207a65726f206164647265737300600082015250565b7f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572600082015250565b7f42455032303a207472616e7366657220746f20746865207a65726f206164647260008201527f6573730000000000000000000000000000000000000000000000000000000000602082015250565b7f42455032303a20617070726f766520746f20746865207a65726f20616464726560008201527f7373000000000000000000000000000000000000000000000000000000000000602082015250565b61190f81611600565b811461191a57600080fd5b50565b6119268161163e565b811461193157600080fd5b5056fe42455032303a207472616e7366657220616d6f756e74206578636565647320616c6c6f77616e636542455032303a207472616e7366657220616d6f756e7420657863656564732062616c616e636542455032303a2064656372656173656420616c6c6f77616e63652062656c6f77207a65726fa264697066735822122098a1a7f19d7fe01dbf849ea059732c5689b23dd6c71c4aaa796c1a9181506fb964736f6c63430008040033",
}

// BEP20ABI is the input ABI used to generate the binding from.
// Deprecated: Use BEP20MetaData.ABI instead.
var BEP20ABI = BEP20MetaData.ABI

// BEP20Bin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use BEP20MetaData.Bin instead.
var BEP20Bin = BEP20MetaData.Bin

// DeployBEP20 deploys a new Ethereum contract, binding an instance of BEP20 to it.
func DeployBEP20(auth *bind.TransactOpts, backend bind.ContractBackend, tokenName string, tokenSymbol string, tokenDecimals uint8) (common.Address, *types.Transaction, *BEP20, error) {
	parsed, err := BEP20MetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(BEP20Bin), backend, tokenName, tokenSymbol, tokenDecimals)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &BEP20{BEP20Caller: BEP20Caller{contract: contract}, BEP20Transactor: BEP20Transactor{contract: contract}, BEP20Filterer: BEP20Filterer{contract: contract}}, nil
}

// BEP20 is an auto generated Go binding around an Ethereum contract.
type BEP20 struct {
	BEP20Caller     // Read-only binding to the contract
	BEP20Transactor // Write-only binding to the contract
	BEP20Filterer   // Log filterer for contract events
}

// BEP20Caller is an auto generated read-only Go binding around an Ethereum contract.
type BEP20Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BEP20Transactor is an auto generated write-only Go binding around an Ethereum contract.
type BEP20Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BEP20Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BEP20Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BEP20Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BEP20Session struct {
	Contract     *BEP20            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BEP20CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BEP20CallerSession struct {
	Contract *BEP20Caller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// BEP20TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BEP20TransactorSession struct {
	Contract     *BEP20Transactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BEP20Raw is an auto generated low-level Go binding around an Ethereum contract.
type BEP20Raw struct {
	Contract *BEP20 // Generic contract binding to access the raw methods on
}

// BEP20CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BEP20CallerRaw struct {
	Contract *BEP20Caller // Generic read-only contract binding to access the raw methods on
}

// BEP20TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BEP20TransactorRaw struct {
	Contract *BEP20Transactor // Generic write-only contract binding to access the raw methods on
}

// NewBEP20 creates a new instance of BEP20, bound to a specific deployed contract.
func NewBEP20(address common.Address, backend bind.ContractBackend) (*BEP20, error) {
	contract, err := bindBEP20(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BEP20{BEP20Caller: BEP20Caller{contract: contract}, BEP20Transactor: BEP20Transactor{contract: contract}, BEP20Filterer: BEP20Filterer{contract: contract}}, nil
}

// NewBEP20Caller creates a new read-only instance of BEP20, bound to a specific deployed contract.
func NewBEP20Caller(address common.Address, caller bind.ContractCaller) (*BEP20Caller, error) {
	contract, err := bindBEP20(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BEP20Caller{contract: contract}, nil
}

// NewBEP20Transactor creates a new write-only instance of BEP20, bound to a specific deployed contract.
func NewBEP20Transactor(address common.Address, transactor bind.ContractTransactor) (*BEP20Transactor, error) {
	contract, err := bindBEP20(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BEP20Transactor{contract: contract}, nil
}

// NewBEP20Filterer creates a new log filterer instance of BEP20, bound to a specific deployed contract.
func NewBEP20Filterer(address common.Address, filterer bind.ContractFilterer) (*BEP20Filterer, error) {
	contract, err := bindBEP20(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BEP20Filterer{contract: contract}, nil
}

// bindBEP20 binds a generic wrapper to an already deployed contract.
func bindBEP20(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BEP20ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BEP20 *BEP20Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BEP20.Contract.BEP20Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BEP20 *BEP20Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BEP20.Contract.BEP20Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BEP20 *BEP20Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BEP20.Contract.BEP20Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BEP20 *BEP20CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BEP20.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BEP20 *BEP20TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BEP20.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BEP20 *BEP20TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BEP20.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_BEP20 *BEP20Caller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _BEP20.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_BEP20 *BEP20Session) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _BEP20.Contract.Allowance(&_BEP20.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_BEP20 *BEP20CallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _BEP20.Contract.Allowance(&_BEP20.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_BEP20 *BEP20Caller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _BEP20.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_BEP20 *BEP20Session) BalanceOf(account common.Address) (*big.Int, error) {
	return _BEP20.Contract.BalanceOf(&_BEP20.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_BEP20 *BEP20CallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _BEP20.Contract.BalanceOf(&_BEP20.CallOpts, account)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_BEP20 *BEP20Caller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _BEP20.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_BEP20 *BEP20Session) Decimals() (uint8, error) {
	return _BEP20.Contract.Decimals(&_BEP20.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_BEP20 *BEP20CallerSession) Decimals() (uint8, error) {
	return _BEP20.Contract.Decimals(&_BEP20.CallOpts)
}

// GetOwner is a free data retrieval call binding the contract method 0x893d20e8.
//
// Solidity: function getOwner() view returns(address)
func (_BEP20 *BEP20Caller) GetOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BEP20.contract.Call(opts, &out, "getOwner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetOwner is a free data retrieval call binding the contract method 0x893d20e8.
//
// Solidity: function getOwner() view returns(address)
func (_BEP20 *BEP20Session) GetOwner() (common.Address, error) {
	return _BEP20.Contract.GetOwner(&_BEP20.CallOpts)
}

// GetOwner is a free data retrieval call binding the contract method 0x893d20e8.
//
// Solidity: function getOwner() view returns(address)
func (_BEP20 *BEP20CallerSession) GetOwner() (common.Address, error) {
	return _BEP20.Contract.GetOwner(&_BEP20.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_BEP20 *BEP20Caller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _BEP20.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_BEP20 *BEP20Session) Name() (string, error) {
	return _BEP20.Contract.Name(&_BEP20.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_BEP20 *BEP20CallerSession) Name() (string, error) {
	return _BEP20.Contract.Name(&_BEP20.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BEP20 *BEP20Caller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BEP20.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BEP20 *BEP20Session) Owner() (common.Address, error) {
	return _BEP20.Contract.Owner(&_BEP20.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BEP20 *BEP20CallerSession) Owner() (common.Address, error) {
	return _BEP20.Contract.Owner(&_BEP20.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_BEP20 *BEP20Caller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _BEP20.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_BEP20 *BEP20Session) Symbol() (string, error) {
	return _BEP20.Contract.Symbol(&_BEP20.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_BEP20 *BEP20CallerSession) Symbol() (string, error) {
	return _BEP20.Contract.Symbol(&_BEP20.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_BEP20 *BEP20Caller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BEP20.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_BEP20 *BEP20Session) TotalSupply() (*big.Int, error) {
	return _BEP20.Contract.TotalSupply(&_BEP20.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_BEP20 *BEP20CallerSession) TotalSupply() (*big.Int, error) {
	return _BEP20.Contract.TotalSupply(&_BEP20.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_BEP20 *BEP20Transactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BEP20.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_BEP20 *BEP20Session) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BEP20.Contract.Approve(&_BEP20.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_BEP20 *BEP20TransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BEP20.Contract.Approve(&_BEP20.TransactOpts, spender, amount)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_BEP20 *BEP20Transactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _BEP20.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_BEP20 *BEP20Session) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _BEP20.Contract.DecreaseAllowance(&_BEP20.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_BEP20 *BEP20TransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _BEP20.Contract.DecreaseAllowance(&_BEP20.TransactOpts, spender, subtractedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_BEP20 *BEP20Transactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _BEP20.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_BEP20 *BEP20Session) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _BEP20.Contract.IncreaseAllowance(&_BEP20.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_BEP20 *BEP20TransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _BEP20.Contract.IncreaseAllowance(&_BEP20.TransactOpts, spender, addedValue)
}

// Mint is a paid mutator transaction binding the contract method 0xa0712d68.
//
// Solidity: function mint(uint256 amount) returns(bool)
func (_BEP20 *BEP20Transactor) Mint(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _BEP20.contract.Transact(opts, "mint", amount)
}

// Mint is a paid mutator transaction binding the contract method 0xa0712d68.
//
// Solidity: function mint(uint256 amount) returns(bool)
func (_BEP20 *BEP20Session) Mint(amount *big.Int) (*types.Transaction, error) {
	return _BEP20.Contract.Mint(&_BEP20.TransactOpts, amount)
}

// Mint is a paid mutator transaction binding the contract method 0xa0712d68.
//
// Solidity: function mint(uint256 amount) returns(bool)
func (_BEP20 *BEP20TransactorSession) Mint(amount *big.Int) (*types.Transaction, error) {
	return _BEP20.Contract.Mint(&_BEP20.TransactOpts, amount)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_BEP20 *BEP20Transactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BEP20.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_BEP20 *BEP20Session) RenounceOwnership() (*types.Transaction, error) {
	return _BEP20.Contract.RenounceOwnership(&_BEP20.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_BEP20 *BEP20TransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _BEP20.Contract.RenounceOwnership(&_BEP20.TransactOpts)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_BEP20 *BEP20Transactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BEP20.contract.Transact(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_BEP20 *BEP20Session) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BEP20.Contract.Transfer(&_BEP20.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_BEP20 *BEP20TransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BEP20.Contract.Transfer(&_BEP20.TransactOpts, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_BEP20 *BEP20Transactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BEP20.contract.Transact(opts, "transferFrom", sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_BEP20 *BEP20Session) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BEP20.Contract.TransferFrom(&_BEP20.TransactOpts, sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_BEP20 *BEP20TransactorSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _BEP20.Contract.TransferFrom(&_BEP20.TransactOpts, sender, recipient, amount)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_BEP20 *BEP20Transactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _BEP20.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_BEP20 *BEP20Session) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _BEP20.Contract.TransferOwnership(&_BEP20.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_BEP20 *BEP20TransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _BEP20.Contract.TransferOwnership(&_BEP20.TransactOpts, newOwner)
}

// BEP20ApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the BEP20 contract.
type BEP20ApprovalIterator struct {
	Event *BEP20Approval // Event containing the contract specifics and raw log

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
func (it *BEP20ApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BEP20Approval)
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
		it.Event = new(BEP20Approval)
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
func (it *BEP20ApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BEP20ApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BEP20Approval represents a Approval event raised by the BEP20 contract.
type BEP20Approval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_BEP20 *BEP20Filterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*BEP20ApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _BEP20.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &BEP20ApprovalIterator{contract: _BEP20.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_BEP20 *BEP20Filterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *BEP20Approval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _BEP20.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BEP20Approval)
				if err := _BEP20.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_BEP20 *BEP20Filterer) ParseApproval(log types.Log) (*BEP20Approval, error) {
	event := new(BEP20Approval)
	if err := _BEP20.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BEP20OwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the BEP20 contract.
type BEP20OwnershipTransferredIterator struct {
	Event *BEP20OwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *BEP20OwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BEP20OwnershipTransferred)
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
		it.Event = new(BEP20OwnershipTransferred)
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
func (it *BEP20OwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BEP20OwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BEP20OwnershipTransferred represents a OwnershipTransferred event raised by the BEP20 contract.
type BEP20OwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_BEP20 *BEP20Filterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*BEP20OwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _BEP20.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &BEP20OwnershipTransferredIterator{contract: _BEP20.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_BEP20 *BEP20Filterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *BEP20OwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _BEP20.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BEP20OwnershipTransferred)
				if err := _BEP20.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_BEP20 *BEP20Filterer) ParseOwnershipTransferred(log types.Log) (*BEP20OwnershipTransferred, error) {
	event := new(BEP20OwnershipTransferred)
	if err := _BEP20.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BEP20TransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the BEP20 contract.
type BEP20TransferIterator struct {
	Event *BEP20Transfer // Event containing the contract specifics and raw log

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
func (it *BEP20TransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BEP20Transfer)
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
		it.Event = new(BEP20Transfer)
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
func (it *BEP20TransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BEP20TransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BEP20Transfer represents a Transfer event raised by the BEP20 contract.
type BEP20Transfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_BEP20 *BEP20Filterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*BEP20TransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _BEP20.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &BEP20TransferIterator{contract: _BEP20.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_BEP20 *BEP20Filterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *BEP20Transfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _BEP20.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BEP20Transfer)
				if err := _BEP20.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_BEP20 *BEP20Filterer) ParseTransfer(log types.Log) (*BEP20Transfer, error) {
	event := new(BEP20Transfer)
	if err := _BEP20.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
