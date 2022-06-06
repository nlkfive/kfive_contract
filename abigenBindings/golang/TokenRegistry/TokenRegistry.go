// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package TokenRegistry

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

// TokenRegistryMetaData contains all meta data concerning the TokenRegistry contract.
var TokenRegistryMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"offchain\",\"type\":\"bytes32\"}],\"name\":\"__addNewToken\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"a\",\"type\":\"address\"}],\"name\":\"addAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"admin\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"a\",\"type\":\"address\"}],\"name\":\"removeAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"tokenMapping\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"decimals\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"valid\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"decimals\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"offchain\",\"type\":\"bytes32\"}],\"name\":\"addNewToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"tokenIsExisted\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"getBalanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"getTokenByAddr\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60806040523480156200001157600080fd5b50336000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555062000063336200006960201b60201c565b620002bc565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614620000fa576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401620000f19062000237565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614156200016d576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401620001649062000215565b60405180910390fd5b60018060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff02191690831515021790555050565b6000620001d6600e8362000259565b9150620001e3826200026a565b602082019050919050565b6000620001fd600f8362000259565b91506200020a8262000293565b602082019050919050565b600060208201905081810360008301526200023081620001c7565b9050919050565b600060208201905081810360008301526200025281620001ee565b9050919050565b600082825260208201905092915050565b7f61646472657373206973206e696c000000000000000000000000000000000000600082015250565b7f6e6f742073757065722061646d696e0000000000000000000000000000000000600082015250565b61145d80620002cc6000396000f3fe608060405234801561001057600080fd5b50600436106100935760003560e01c806363a846f81161006657806363a846f8146101305780637048027514610160578063705718281461017c57806399a32482146101ae578063ba27f50b146101ca57610093565b80631785f53c146100985780632c54de4f146100b457806353290b44146100d057806361eeeb8c14610100575b600080fd5b6100b260048036038101906100ad9190610cfd565b6101fe565b005b6100ce60048036038101906100c99190610d62565b610357565b005b6100ea60048036038101906100e59190610d26565b6103ed565b6040516100f79190611139565b60405180910390f35b61011a60048036038101906101159190610cfd565b610480565b604051610127919061105e565b60405180910390f35b61014a60048036038101906101459190610cfd565b6104de565b604051610157919061105e565b60405180910390f35b61017a60048036038101906101759190610cfd565b6104fe565b005b61019660048036038101906101919190610cfd565b610656565b6040516101a593929190611094565b60405180910390f35b6101c860048036038101906101c39190610dc5565b6107dd565b005b6101e460048036038101906101df9190610cfd565b6109f4565b6040516101f5959493929190610ffd565b60405180910390f35b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461028c576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610283906110f9565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614156102fc576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016102f3906110d9565b60405180910390fd5b6000600160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff02191690831515021790555050565b8373ffffffffffffffffffffffffffffffffffffffff166323b872dd8484846040518463ffffffff1660e01b815260040161039493929190610fc6565b602060405180830381600087803b1580156103ae57600080fd5b505af11580156103c2573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906103e69190610e6c565b5050505050565b60008273ffffffffffffffffffffffffffffffffffffffff166370a08231836040518263ffffffff1660e01b81526004016104289190610fab565b60206040518083038186803b15801561044057600080fd5b505afa158015610454573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906104789190610e95565b905092915050565b600080600260008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002090508060030160019054906101000a900460ff16915050919050565b60016020528060005260406000206000915054906101000a900460ff1681565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461058c576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610583906110f9565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614156105fc576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016105f3906110d9565b60405180910390fd5b60018060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff02191690831515021790555050565b606080600080600260008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020905080600101816002018260030160009054906101000a900460ff168280546106c390611267565b80601f01602080910402602001604051908101604052809291908181526020018280546106ef90611267565b801561073c5780601f106107115761010080835404028352916020019161073c565b820191906000526020600020905b81548152906001019060200180831161071f57829003601f168201915b5050505050925081805461074f90611267565b80601f016020809104026020016040519081016040528092919081815260200182805461077b90611267565b80156107c85780601f1061079d576101008083540402835291602001916107c8565b820191906000526020600020905b8154815290600101906020018083116107ab57829003601f168201915b50505050509150935093509350509193909250565b600160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16610869576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161086090611119565b60405180910390fd5b60006040518060a001604052808773ffffffffffffffffffffffffffffffffffffffff1681526020018681526020018581526020018460ff16815260200160011515815250905080600260008873ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008201518160000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506020820151816001019080519060200190610953929190610b74565b506040820151816002019080519060200190610970929190610b74565b5060608201518160030160006101000a81548160ff021916908360ff16021790555060808201518160030160016101000a81548160ff0219169083151502179055509050507f662298be6d0168dcaa626ce9f521c3992f79a1ae8a49dd23a143dd4ebf7d4b4a826040516109e49190611079565b60405180910390a1505050505050565b60026020528060005260406000206000915090508060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690806001018054610a3d90611267565b80601f0160208091040260200160405190810160405280929190818152602001828054610a6990611267565b8015610ab65780601f10610a8b57610100808354040283529160200191610ab6565b820191906000526020600020905b815481529060010190602001808311610a9957829003601f168201915b505050505090806002018054610acb90611267565b80601f0160208091040260200160405190810160405280929190818152602001828054610af790611267565b8015610b445780601f10610b1957610100808354040283529160200191610b44565b820191906000526020600020905b815481529060010190602001808311610b2757829003601f168201915b5050505050908060030160009054906101000a900460ff16908060030160019054906101000a900460ff16905085565b828054610b8090611267565b90600052602060002090601f016020900481019282610ba25760008555610be9565b82601f10610bbb57805160ff1916838001178555610be9565b82800160010185558215610be9579182015b82811115610be8578251825591602001919060010190610bcd565b5b509050610bf69190610bfa565b5090565b5b80821115610c13576000816000905550600101610bfb565b5090565b6000610c2a610c2584611179565b611154565b905082815260208101848484011115610c4257600080fd5b610c4d848285611225565b509392505050565b600081359050610c64816113b4565b92915050565b600081519050610c79816113cb565b92915050565b600081359050610c8e816113e2565b92915050565b600082601f830112610ca557600080fd5b8135610cb5848260208601610c17565b91505092915050565b600081359050610ccd816113f9565b92915050565b600081519050610ce2816113f9565b92915050565b600081359050610cf781611410565b92915050565b600060208284031215610d0f57600080fd5b6000610d1d84828501610c55565b91505092915050565b60008060408385031215610d3957600080fd5b6000610d4785828601610c55565b9250506020610d5885828601610c55565b9150509250929050565b60008060008060808587031215610d7857600080fd5b6000610d8687828801610c55565b9450506020610d9787828801610c55565b9350506040610da887828801610c55565b9250506060610db987828801610cbe565b91505092959194509250565b600080600080600060a08688031215610ddd57600080fd5b6000610deb88828901610c55565b955050602086013567ffffffffffffffff811115610e0857600080fd5b610e1488828901610c94565b945050604086013567ffffffffffffffff811115610e3157600080fd5b610e3d88828901610c94565b9350506060610e4e88828901610ce8565b9250506080610e5f88828901610c7f565b9150509295509295909350565b600060208284031215610e7e57600080fd5b6000610e8c84828501610c6a565b91505092915050565b600060208284031215610ea757600080fd5b6000610eb584828501610cd3565b91505092915050565b610ec7816111c6565b82525050565b610ed6816111d8565b82525050565b610ee5816111e4565b82525050565b6000610ef6826111aa565b610f0081856111b5565b9350610f10818560208601611234565b610f1981611328565b840191505092915050565b6000610f31600e836111b5565b9150610f3c82611339565b602082019050919050565b6000610f54600f836111b5565b9150610f5f82611362565b602082019050919050565b6000610f776009836111b5565b9150610f828261138b565b602082019050919050565b610f968161120e565b82525050565b610fa581611218565b82525050565b6000602082019050610fc06000830184610ebe565b92915050565b6000606082019050610fdb6000830186610ebe565b610fe86020830185610ebe565b610ff56040830184610f8d565b949350505050565b600060a0820190506110126000830188610ebe565b81810360208301526110248187610eeb565b905081810360408301526110388186610eeb565b90506110476060830185610f9c565b6110546080830184610ecd565b9695505050505050565b60006020820190506110736000830184610ecd565b92915050565b600060208201905061108e6000830184610edc565b92915050565b600060608201905081810360008301526110ae8186610eeb565b905081810360208301526110c28185610eeb565b90506110d16040830184610f9c565b949350505050565b600060208201905081810360008301526110f281610f24565b9050919050565b6000602082019050818103600083015261111281610f47565b9050919050565b6000602082019050818103600083015261113281610f6a565b9050919050565b600060208201905061114e6000830184610f8d565b92915050565b600061115e61116f565b905061116a8282611299565b919050565b6000604051905090565b600067ffffffffffffffff821115611194576111936112f9565b5b61119d82611328565b9050602081019050919050565b600081519050919050565b600082825260208201905092915050565b60006111d1826111ee565b9050919050565b60008115159050919050565b6000819050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b600060ff82169050919050565b82818337600083830152505050565b60005b83811015611252578082015181840152602081019050611237565b83811115611261576000848401525b50505050565b6000600282049050600182168061127f57607f821691505b60208210811415611293576112926112ca565b5b50919050565b6112a282611328565b810181811067ffffffffffffffff821117156112c1576112c06112f9565b5b80604052505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6000601f19601f8301169050919050565b7f61646472657373206973206e696c000000000000000000000000000000000000600082015250565b7f6e6f742073757065722061646d696e0000000000000000000000000000000000600082015250565b7f6e6f742061646d696e0000000000000000000000000000000000000000000000600082015250565b6113bd816111c6565b81146113c857600080fd5b50565b6113d4816111d8565b81146113df57600080fd5b50565b6113eb816111e4565b81146113f657600080fd5b50565b6114028161120e565b811461140d57600080fd5b50565b61141981611218565b811461142457600080fd5b5056fea26469706673582212200b68888efba669ad7263461ef88e1fff19502f71c6894492c591894b9a76f9ce64736f6c63430008040033",
}

// TokenRegistryABI is the input ABI used to generate the binding from.
// Deprecated: Use TokenRegistryMetaData.ABI instead.
var TokenRegistryABI = TokenRegistryMetaData.ABI

// TokenRegistryBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use TokenRegistryMetaData.Bin instead.
var TokenRegistryBin = TokenRegistryMetaData.Bin

// DeployTokenRegistry deploys a new Ethereum contract, binding an instance of TokenRegistry to it.
func DeployTokenRegistry(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *TokenRegistry, error) {
	parsed, err := TokenRegistryMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(TokenRegistryBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &TokenRegistry{TokenRegistryCaller: TokenRegistryCaller{contract: contract}, TokenRegistryTransactor: TokenRegistryTransactor{contract: contract}, TokenRegistryFilterer: TokenRegistryFilterer{contract: contract}}, nil
}

// TokenRegistry is an auto generated Go binding around an Ethereum contract.
type TokenRegistry struct {
	TokenRegistryCaller     // Read-only binding to the contract
	TokenRegistryTransactor // Write-only binding to the contract
	TokenRegistryFilterer   // Log filterer for contract events
}

// TokenRegistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type TokenRegistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenRegistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TokenRegistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenRegistryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TokenRegistryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenRegistrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TokenRegistrySession struct {
	Contract     *TokenRegistry    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TokenRegistryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TokenRegistryCallerSession struct {
	Contract *TokenRegistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// TokenRegistryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TokenRegistryTransactorSession struct {
	Contract     *TokenRegistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// TokenRegistryRaw is an auto generated low-level Go binding around an Ethereum contract.
type TokenRegistryRaw struct {
	Contract *TokenRegistry // Generic contract binding to access the raw methods on
}

// TokenRegistryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TokenRegistryCallerRaw struct {
	Contract *TokenRegistryCaller // Generic read-only contract binding to access the raw methods on
}

// TokenRegistryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TokenRegistryTransactorRaw struct {
	Contract *TokenRegistryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTokenRegistry creates a new instance of TokenRegistry, bound to a specific deployed contract.
func NewTokenRegistry(address common.Address, backend bind.ContractBackend) (*TokenRegistry, error) {
	contract, err := bindTokenRegistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TokenRegistry{TokenRegistryCaller: TokenRegistryCaller{contract: contract}, TokenRegistryTransactor: TokenRegistryTransactor{contract: contract}, TokenRegistryFilterer: TokenRegistryFilterer{contract: contract}}, nil
}

// NewTokenRegistryCaller creates a new read-only instance of TokenRegistry, bound to a specific deployed contract.
func NewTokenRegistryCaller(address common.Address, caller bind.ContractCaller) (*TokenRegistryCaller, error) {
	contract, err := bindTokenRegistry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TokenRegistryCaller{contract: contract}, nil
}

// NewTokenRegistryTransactor creates a new write-only instance of TokenRegistry, bound to a specific deployed contract.
func NewTokenRegistryTransactor(address common.Address, transactor bind.ContractTransactor) (*TokenRegistryTransactor, error) {
	contract, err := bindTokenRegistry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TokenRegistryTransactor{contract: contract}, nil
}

// NewTokenRegistryFilterer creates a new log filterer instance of TokenRegistry, bound to a specific deployed contract.
func NewTokenRegistryFilterer(address common.Address, filterer bind.ContractFilterer) (*TokenRegistryFilterer, error) {
	contract, err := bindTokenRegistry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TokenRegistryFilterer{contract: contract}, nil
}

// bindTokenRegistry binds a generic wrapper to an already deployed contract.
func bindTokenRegistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TokenRegistryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TokenRegistry *TokenRegistryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TokenRegistry.Contract.TokenRegistryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TokenRegistry *TokenRegistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TokenRegistry.Contract.TokenRegistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TokenRegistry *TokenRegistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TokenRegistry.Contract.TokenRegistryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TokenRegistry *TokenRegistryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TokenRegistry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TokenRegistry *TokenRegistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TokenRegistry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TokenRegistry *TokenRegistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TokenRegistry.Contract.contract.Transact(opts, method, params...)
}

// Admin is a free data retrieval call binding the contract method 0x63a846f8.
//
// Solidity: function admin(address ) view returns(bool)
func (_TokenRegistry *TokenRegistryCaller) Admin(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _TokenRegistry.contract.Call(opts, &out, "admin", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Admin is a free data retrieval call binding the contract method 0x63a846f8.
//
// Solidity: function admin(address ) view returns(bool)
func (_TokenRegistry *TokenRegistrySession) Admin(arg0 common.Address) (bool, error) {
	return _TokenRegistry.Contract.Admin(&_TokenRegistry.CallOpts, arg0)
}

// Admin is a free data retrieval call binding the contract method 0x63a846f8.
//
// Solidity: function admin(address ) view returns(bool)
func (_TokenRegistry *TokenRegistryCallerSession) Admin(arg0 common.Address) (bool, error) {
	return _TokenRegistry.Contract.Admin(&_TokenRegistry.CallOpts, arg0)
}

// GetBalanceOf is a free data retrieval call binding the contract method 0x53290b44.
//
// Solidity: function getBalanceOf(address token, address addr) view returns(uint256)
func (_TokenRegistry *TokenRegistryCaller) GetBalanceOf(opts *bind.CallOpts, token common.Address, addr common.Address) (*big.Int, error) {
	var out []interface{}
	err := _TokenRegistry.contract.Call(opts, &out, "getBalanceOf", token, addr)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBalanceOf is a free data retrieval call binding the contract method 0x53290b44.
//
// Solidity: function getBalanceOf(address token, address addr) view returns(uint256)
func (_TokenRegistry *TokenRegistrySession) GetBalanceOf(token common.Address, addr common.Address) (*big.Int, error) {
	return _TokenRegistry.Contract.GetBalanceOf(&_TokenRegistry.CallOpts, token, addr)
}

// GetBalanceOf is a free data retrieval call binding the contract method 0x53290b44.
//
// Solidity: function getBalanceOf(address token, address addr) view returns(uint256)
func (_TokenRegistry *TokenRegistryCallerSession) GetBalanceOf(token common.Address, addr common.Address) (*big.Int, error) {
	return _TokenRegistry.Contract.GetBalanceOf(&_TokenRegistry.CallOpts, token, addr)
}

// GetTokenByAddr is a free data retrieval call binding the contract method 0x70571828.
//
// Solidity: function getTokenByAddr(address token) view returns(string, string, uint8)
func (_TokenRegistry *TokenRegistryCaller) GetTokenByAddr(opts *bind.CallOpts, token common.Address) (string, string, uint8, error) {
	var out []interface{}
	err := _TokenRegistry.contract.Call(opts, &out, "getTokenByAddr", token)

	if err != nil {
		return *new(string), *new(string), *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)
	out1 := *abi.ConvertType(out[1], new(string)).(*string)
	out2 := *abi.ConvertType(out[2], new(uint8)).(*uint8)

	return out0, out1, out2, err

}

// GetTokenByAddr is a free data retrieval call binding the contract method 0x70571828.
//
// Solidity: function getTokenByAddr(address token) view returns(string, string, uint8)
func (_TokenRegistry *TokenRegistrySession) GetTokenByAddr(token common.Address) (string, string, uint8, error) {
	return _TokenRegistry.Contract.GetTokenByAddr(&_TokenRegistry.CallOpts, token)
}

// GetTokenByAddr is a free data retrieval call binding the contract method 0x70571828.
//
// Solidity: function getTokenByAddr(address token) view returns(string, string, uint8)
func (_TokenRegistry *TokenRegistryCallerSession) GetTokenByAddr(token common.Address) (string, string, uint8, error) {
	return _TokenRegistry.Contract.GetTokenByAddr(&_TokenRegistry.CallOpts, token)
}

// TokenIsExisted is a free data retrieval call binding the contract method 0x61eeeb8c.
//
// Solidity: function tokenIsExisted(address token) view returns(bool)
func (_TokenRegistry *TokenRegistryCaller) TokenIsExisted(opts *bind.CallOpts, token common.Address) (bool, error) {
	var out []interface{}
	err := _TokenRegistry.contract.Call(opts, &out, "tokenIsExisted", token)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// TokenIsExisted is a free data retrieval call binding the contract method 0x61eeeb8c.
//
// Solidity: function tokenIsExisted(address token) view returns(bool)
func (_TokenRegistry *TokenRegistrySession) TokenIsExisted(token common.Address) (bool, error) {
	return _TokenRegistry.Contract.TokenIsExisted(&_TokenRegistry.CallOpts, token)
}

// TokenIsExisted is a free data retrieval call binding the contract method 0x61eeeb8c.
//
// Solidity: function tokenIsExisted(address token) view returns(bool)
func (_TokenRegistry *TokenRegistryCallerSession) TokenIsExisted(token common.Address) (bool, error) {
	return _TokenRegistry.Contract.TokenIsExisted(&_TokenRegistry.CallOpts, token)
}

// TokenMapping is a free data retrieval call binding the contract method 0xba27f50b.
//
// Solidity: function tokenMapping(address ) view returns(address token, string symbol, string name, uint8 decimals, bool valid)
func (_TokenRegistry *TokenRegistryCaller) TokenMapping(opts *bind.CallOpts, arg0 common.Address) (struct {
	Token    common.Address
	Symbol   string
	Name     string
	Decimals uint8
	Valid    bool
}, error) {
	var out []interface{}
	err := _TokenRegistry.contract.Call(opts, &out, "tokenMapping", arg0)

	outstruct := new(struct {
		Token    common.Address
		Symbol   string
		Name     string
		Decimals uint8
		Valid    bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Token = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Symbol = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.Name = *abi.ConvertType(out[2], new(string)).(*string)
	outstruct.Decimals = *abi.ConvertType(out[3], new(uint8)).(*uint8)
	outstruct.Valid = *abi.ConvertType(out[4], new(bool)).(*bool)

	return *outstruct, err

}

// TokenMapping is a free data retrieval call binding the contract method 0xba27f50b.
//
// Solidity: function tokenMapping(address ) view returns(address token, string symbol, string name, uint8 decimals, bool valid)
func (_TokenRegistry *TokenRegistrySession) TokenMapping(arg0 common.Address) (struct {
	Token    common.Address
	Symbol   string
	Name     string
	Decimals uint8
	Valid    bool
}, error) {
	return _TokenRegistry.Contract.TokenMapping(&_TokenRegistry.CallOpts, arg0)
}

// TokenMapping is a free data retrieval call binding the contract method 0xba27f50b.
//
// Solidity: function tokenMapping(address ) view returns(address token, string symbol, string name, uint8 decimals, bool valid)
func (_TokenRegistry *TokenRegistryCallerSession) TokenMapping(arg0 common.Address) (struct {
	Token    common.Address
	Symbol   string
	Name     string
	Decimals uint8
	Valid    bool
}, error) {
	return _TokenRegistry.Contract.TokenMapping(&_TokenRegistry.CallOpts, arg0)
}

// AddAdmin is a paid mutator transaction binding the contract method 0x70480275.
//
// Solidity: function addAdmin(address a) returns()
func (_TokenRegistry *TokenRegistryTransactor) AddAdmin(opts *bind.TransactOpts, a common.Address) (*types.Transaction, error) {
	return _TokenRegistry.contract.Transact(opts, "addAdmin", a)
}

// AddAdmin is a paid mutator transaction binding the contract method 0x70480275.
//
// Solidity: function addAdmin(address a) returns()
func (_TokenRegistry *TokenRegistrySession) AddAdmin(a common.Address) (*types.Transaction, error) {
	return _TokenRegistry.Contract.AddAdmin(&_TokenRegistry.TransactOpts, a)
}

// AddAdmin is a paid mutator transaction binding the contract method 0x70480275.
//
// Solidity: function addAdmin(address a) returns()
func (_TokenRegistry *TokenRegistryTransactorSession) AddAdmin(a common.Address) (*types.Transaction, error) {
	return _TokenRegistry.Contract.AddAdmin(&_TokenRegistry.TransactOpts, a)
}

// AddNewToken is a paid mutator transaction binding the contract method 0x99a32482.
//
// Solidity: function addNewToken(address token, string symbol, string name, uint8 decimals, bytes32 offchain) returns()
func (_TokenRegistry *TokenRegistryTransactor) AddNewToken(opts *bind.TransactOpts, token common.Address, symbol string, name string, decimals uint8, offchain [32]byte) (*types.Transaction, error) {
	return _TokenRegistry.contract.Transact(opts, "addNewToken", token, symbol, name, decimals, offchain)
}

// AddNewToken is a paid mutator transaction binding the contract method 0x99a32482.
//
// Solidity: function addNewToken(address token, string symbol, string name, uint8 decimals, bytes32 offchain) returns()
func (_TokenRegistry *TokenRegistrySession) AddNewToken(token common.Address, symbol string, name string, decimals uint8, offchain [32]byte) (*types.Transaction, error) {
	return _TokenRegistry.Contract.AddNewToken(&_TokenRegistry.TransactOpts, token, symbol, name, decimals, offchain)
}

// AddNewToken is a paid mutator transaction binding the contract method 0x99a32482.
//
// Solidity: function addNewToken(address token, string symbol, string name, uint8 decimals, bytes32 offchain) returns()
func (_TokenRegistry *TokenRegistryTransactorSession) AddNewToken(token common.Address, symbol string, name string, decimals uint8, offchain [32]byte) (*types.Transaction, error) {
	return _TokenRegistry.Contract.AddNewToken(&_TokenRegistry.TransactOpts, token, symbol, name, decimals, offchain)
}

// RemoveAdmin is a paid mutator transaction binding the contract method 0x1785f53c.
//
// Solidity: function removeAdmin(address a) returns()
func (_TokenRegistry *TokenRegistryTransactor) RemoveAdmin(opts *bind.TransactOpts, a common.Address) (*types.Transaction, error) {
	return _TokenRegistry.contract.Transact(opts, "removeAdmin", a)
}

// RemoveAdmin is a paid mutator transaction binding the contract method 0x1785f53c.
//
// Solidity: function removeAdmin(address a) returns()
func (_TokenRegistry *TokenRegistrySession) RemoveAdmin(a common.Address) (*types.Transaction, error) {
	return _TokenRegistry.Contract.RemoveAdmin(&_TokenRegistry.TransactOpts, a)
}

// RemoveAdmin is a paid mutator transaction binding the contract method 0x1785f53c.
//
// Solidity: function removeAdmin(address a) returns()
func (_TokenRegistry *TokenRegistryTransactorSession) RemoveAdmin(a common.Address) (*types.Transaction, error) {
	return _TokenRegistry.Contract.RemoveAdmin(&_TokenRegistry.TransactOpts, a)
}

// TransferToken is a paid mutator transaction binding the contract method 0x2c54de4f.
//
// Solidity: function transferToken(address token, address from, address to, uint256 amount) returns()
func (_TokenRegistry *TokenRegistryTransactor) TransferToken(opts *bind.TransactOpts, token common.Address, from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TokenRegistry.contract.Transact(opts, "transferToken", token, from, to, amount)
}

// TransferToken is a paid mutator transaction binding the contract method 0x2c54de4f.
//
// Solidity: function transferToken(address token, address from, address to, uint256 amount) returns()
func (_TokenRegistry *TokenRegistrySession) TransferToken(token common.Address, from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TokenRegistry.Contract.TransferToken(&_TokenRegistry.TransactOpts, token, from, to, amount)
}

// TransferToken is a paid mutator transaction binding the contract method 0x2c54de4f.
//
// Solidity: function transferToken(address token, address from, address to, uint256 amount) returns()
func (_TokenRegistry *TokenRegistryTransactorSession) TransferToken(token common.Address, from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TokenRegistry.Contract.TransferToken(&_TokenRegistry.TransactOpts, token, from, to, amount)
}

// TokenRegistryAddNewTokenIterator is returned from FilterAddNewToken and is used to iterate over the raw logs and unpacked data for AddNewToken events raised by the TokenRegistry contract.
type TokenRegistryAddNewTokenIterator struct {
	Event *TokenRegistryAddNewToken // Event containing the contract specifics and raw log

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
func (it *TokenRegistryAddNewTokenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenRegistryAddNewToken)
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
		it.Event = new(TokenRegistryAddNewToken)
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
func (it *TokenRegistryAddNewTokenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenRegistryAddNewTokenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenRegistryAddNewToken represents a AddNewToken event raised by the TokenRegistry contract.
type TokenRegistryAddNewToken struct {
	Offchain [32]byte
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterAddNewToken is a free log retrieval operation binding the contract event 0x662298be6d0168dcaa626ce9f521c3992f79a1ae8a49dd23a143dd4ebf7d4b4a.
//
// Solidity: event __addNewToken(bytes32 offchain)
func (_TokenRegistry *TokenRegistryFilterer) FilterAddNewToken(opts *bind.FilterOpts) (*TokenRegistryAddNewTokenIterator, error) {

	logs, sub, err := _TokenRegistry.contract.FilterLogs(opts, "__addNewToken")
	if err != nil {
		return nil, err
	}
	return &TokenRegistryAddNewTokenIterator{contract: _TokenRegistry.contract, event: "__addNewToken", logs: logs, sub: sub}, nil
}

// WatchAddNewToken is a free log subscription operation binding the contract event 0x662298be6d0168dcaa626ce9f521c3992f79a1ae8a49dd23a143dd4ebf7d4b4a.
//
// Solidity: event __addNewToken(bytes32 offchain)
func (_TokenRegistry *TokenRegistryFilterer) WatchAddNewToken(opts *bind.WatchOpts, sink chan<- *TokenRegistryAddNewToken) (event.Subscription, error) {

	logs, sub, err := _TokenRegistry.contract.WatchLogs(opts, "__addNewToken")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenRegistryAddNewToken)
				if err := _TokenRegistry.contract.UnpackLog(event, "__addNewToken", log); err != nil {
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

// ParseAddNewToken is a log parse operation binding the contract event 0x662298be6d0168dcaa626ce9f521c3992f79a1ae8a49dd23a143dd4ebf7d4b4a.
//
// Solidity: event __addNewToken(bytes32 offchain)
func (_TokenRegistry *TokenRegistryFilterer) ParseAddNewToken(log types.Log) (*TokenRegistryAddNewToken, error) {
	event := new(TokenRegistryAddNewToken)
	if err := _TokenRegistry.contract.UnpackLog(event, "__addNewToken", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
