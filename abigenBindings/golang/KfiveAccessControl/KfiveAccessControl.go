// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package KfiveAccessControl

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

// KfiveAccessControlMetaData contains all meta data concerning the KfiveAccessControl contract.
var KfiveAccessControlMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PAUSER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getRoleMember\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleMemberCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60806040523480156200001157600080fd5b5062000032620000266200010a60201b60201c565b6200011260201b60201c565b6000600260146101000a81548160ff021916908315150217905550620000716000801b620000656200010a60201b60201c565b620001d860201b60201c565b620000b27fa49807205ce4d355092ef5a8a18f56e8913cf4a201fbe287825b095693c21775620000a66200010a60201b60201c565b620001d860201b60201c565b620001047f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a7fa49807205ce4d355092ef5a8a18f56e8913cf4a201fbe287825b095693c21775620001ee60201b60201c565b620005b6565b600033905090565b6000600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905081600260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b620001ea82826200025160201b60201c565b5050565b600062000201836200031f60201b60201c565b905081600080858152602001908152602001600020600101819055508181847fbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff60405160405180910390a4505050565b6200026882826200033e60201b620007c51760201c565b7fa49807205ce4d355092ef5a8a18f56e8913cf4a201fbe287825b095693c217758214806200029957506000801b82145b156200031b57620002d67f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a826200033e60201b620007c51760201c565b6000801b8214156200031a57620003197fa49807205ce4d355092ef5a8a18f56e8913cf4a201fbe287825b095693c21775826200033e60201b620007c51760201c565b5b5b5050565b6000806000838152602001908152602001600020600101549050919050565b6200035582826200038660201b620007f91760201c565b6200038181600160008581526020019081526020016000206200047760201b620008d91790919060201c565b505050565b620003988282620004af60201b60201c565b6200047357600160008084815260200190815260200160002060000160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff021916908315150217905550620004186200010a60201b60201c565b73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16837f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a45b5050565b6000620004a7836000018373ffffffffffffffffffffffffffffffffffffffff1660001b6200051960201b60201c565b905092915050565b600080600084815260200190815260200160002060000160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16905092915050565b60006200052d83836200059360201b60201c565b620005885782600001829080600181540180825580915050600190039060005260206000200160009091909190915055826000018054905083600101600084815260200190815260200160002081905550600190506200058d565b600090505b92915050565b600080836001016000848152602001908152602001600020541415905092915050565b611eb380620005c66000396000f3fe608060405234801561001057600080fd5b506004361061010b5760003560e01c80638456cb59116100a2578063a217fddf11610071578063a217fddf14610280578063ca15c8731461029e578063d547741f146102ce578063e63ab1e9146102ea578063f2fde38b146103085761010b565b80638456cb59146101f85780638da5cb5b146102025780639010d07c1461022057806391d14854146102505761010b565b80633f4ba83a116100de5780633f4ba83a146101a85780635c975abb146101b2578063715018a6146101d057806375b238fc146101da5761010b565b806301ffc9a714610110578063248a9ca3146101405780632f2ff15d1461017057806336568abe1461018c575b600080fd5b61012a600480360381019061012591906116f8565b610324565b6040516101379190611934565b60405180910390f35b61015a60048036038101906101559190611657565b61039e565b604051610167919061194f565b60405180910390f35b61018a60048036038101906101859190611680565b6103bd565b005b6101a660048036038101906101a19190611680565b6103de565b005b6101b0610461565b005b6101ba610496565b6040516101c79190611934565b60405180910390f35b6101d86104ad565b005b6101e2610535565b6040516101ef919061194f565b60405180910390f35b610200610559565b005b61020a61058e565b6040516102179190611919565b60405180910390f35b61023a600480360381019061023591906116bc565b6105b8565b6040516102479190611919565b60405180910390f35b61026a60048036038101906102659190611680565b6105e7565b6040516102779190611934565b60405180910390f35b610288610651565b604051610295919061194f565b60405180910390f35b6102b860048036038101906102b39190611657565b610658565b6040516102c59190611a4c565b60405180910390f35b6102e860048036038101906102e39190611680565b61067c565b005b6102f261069d565b6040516102ff919061194f565b60405180910390f35b610322600480360381019061031d919061162e565b6106c1565b005b60007f5a05180f000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916827bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19161480610397575061039682610909565b5b9050919050565b6000806000838152602001908152602001600020600101549050919050565b6103c68261039e565b6103cf81610983565b6103d98383610997565b505050565b6103e6610a3b565b73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614610453576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161044a90611a2c565b60405180910390fd5b61045d8282610a43565b5050565b7f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a61048b81610983565b610493610ae7565b50565b6000600260149054906101000a900460ff16905090565b6104b5610a3b565b73ffffffffffffffffffffffffffffffffffffffff166104d361058e565b73ffffffffffffffffffffffffffffffffffffffff1614610529576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161052090611a0c565b60405180910390fd5b6105336000610b89565b565b7fa49807205ce4d355092ef5a8a18f56e8913cf4a201fbe287825b095693c2177581565b7f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a61058381610983565b61058b610c4f565b50565b6000600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b60006105df8260016000868152602001908152602001600020610cf290919063ffffffff16565b905092915050565b600080600084815260200190815260200160002060000160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16905092915050565b6000801b81565b600061067560016000848152602001908152602001600020610d0c565b9050919050565b6106858261039e565b61068e81610983565b6106988383610a43565b505050565b7f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a81565b6106c9610a3b565b73ffffffffffffffffffffffffffffffffffffffff166106e761058e565b73ffffffffffffffffffffffffffffffffffffffff161461073d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161073490611a0c565b60405180910390fd5b6107677fa49807205ce4d355092ef5a8a18f56e8913cf4a201fbe287825b095693c2177582610997565b6107746000801b82610997565b6107a57fa49807205ce4d355092ef5a8a18f56e8913cf4a201fbe287825b095693c217756107a061058e565b610a43565b6107b96000801b6107b461058e565b610a43565b6107c281610d21565b50565b6107cf82826107f9565b6107f481600160008581526020019081526020016000206108d990919063ffffffff16565b505050565b61080382826105e7565b6108d557600160008084815260200190815260200160002060000160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff02191690831515021790555061087a610a3b565b73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16837f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a45b5050565b6000610901836000018373ffffffffffffffffffffffffffffffffffffffff1660001b610e19565b905092915050565b60007f7965db0b000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916827bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916148061097c575061097b82610e89565b5b9050919050565b6109948161098f610a3b565b610ef3565b50565b6109a182826107c5565b7fa49807205ce4d355092ef5a8a18f56e8913cf4a201fbe287825b095693c217758214806109d157506000801b82145b15610a3757610a007f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a826107c5565b6000801b821415610a3657610a357fa49807205ce4d355092ef5a8a18f56e8913cf4a201fbe287825b095693c21775826107c5565b5b5b5050565b600033905090565b610a4d8282610f90565b7fa49807205ce4d355092ef5a8a18f56e8913cf4a201fbe287825b095693c21775821480610a7d57506000801b82145b15610ae357610aac7f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a82610f90565b6000801b821415610ae257610ae17fa49807205ce4d355092ef5a8a18f56e8913cf4a201fbe287825b095693c21775826107c5565b5b5b5050565b610aef610496565b610b2e576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610b25906119ac565b60405180910390fd5b6000600260146101000a81548160ff0219169083151502179055507f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa610b72610a3b565b604051610b7f9190611919565b60405180910390a1565b6000600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905081600260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b610c57610496565b15610c97576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610c8e906119ec565b60405180910390fd5b6001600260146101000a81548160ff0219169083151502179055507f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258610cdb610a3b565b604051610ce89190611919565b60405180910390a1565b6000610d018360000183610fc4565b60001c905092915050565b6000610d1a82600001611015565b9050919050565b610d29610a3b565b73ffffffffffffffffffffffffffffffffffffffff16610d4761058e565b73ffffffffffffffffffffffffffffffffffffffff1614610d9d576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610d9490611a0c565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff161415610e0d576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610e04906119cc565b60405180910390fd5b610e1681610b89565b50565b6000610e258383611026565b610e7e578260000182908060018154018082558091505060019003906000526020600020016000909190919091505582600001805490508360010160008481526020019081526020016000208190555060019050610e83565b600090505b92915050565b60007f01ffc9a7000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916827bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916149050919050565b610efd82826105e7565b610f8c57610f228173ffffffffffffffffffffffffffffffffffffffff166014611049565b610f308360001c6020611049565b604051602001610f419291906118df565b6040516020818303038152906040526040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610f83919061196a565b60405180910390fd5b5050565b610f9a8282611343565b610fbf816001600085815260200190815260200160002061142490919063ffffffff16565b505050565b6000826000018281548110611002577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b9060005260206000200154905092915050565b600081600001805490509050919050565b600080836001016000848152602001908152602001600020541415905092915050565b60606000600283600261105c9190611ae4565b6110669190611a8e565b67ffffffffffffffff8111156110a5577f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6040519080825280601f01601f1916602001820160405280156110d75781602001600182028036833780820191505090505b5090507f300000000000000000000000000000000000000000000000000000000000000081600081518110611135577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a9053507f7800000000000000000000000000000000000000000000000000000000000000816001815181106111bf577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a905350600060018460026111ff9190611ae4565b6112099190611a8e565b90505b60018111156112f5577f3031323334353637383961626364656600000000000000000000000000000000600f861660108110611271577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b1a60f81b8282815181106112ae577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a905350600485901c9450806112ee90611c23565b905061120c565b5060008414611339576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016113309061198c565b60405180910390fd5b8091505092915050565b61134d82826105e7565b1561142057600080600084815260200190815260200160002060000160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055506113c5610a3b565b73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16837ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b60405160405180910390a45b5050565b600061144c836000018373ffffffffffffffffffffffffffffffffffffffff1660001b611454565b905092915050565b600080836001016000848152602001908152602001600020549050600081146115ce5760006001826114869190611b3e565b905060006001866000018054905061149e9190611b3e565b90508181146115595760008660000182815481106114e5577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b906000526020600020015490508087600001848154811061152f577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b90600052602060002001819055508387600101600083815260200190815260200160002081905550505b85600001805480611593577f4e487b7100000000000000000000000000000000000000000000000000000000600052603160045260246000fd5b6001900381819060005260206000200160009055905585600101600086815260200190815260200160002060009055600193505050506115d4565b60009150505b92915050565b6000813590506115e981611e21565b92915050565b6000813590506115fe81611e38565b92915050565b60008135905061161381611e4f565b92915050565b60008135905061162881611e66565b92915050565b60006020828403121561164057600080fd5b600061164e848285016115da565b91505092915050565b60006020828403121561166957600080fd5b6000611677848285016115ef565b91505092915050565b6000806040838503121561169357600080fd5b60006116a1858286016115ef565b92505060206116b2858286016115da565b9150509250929050565b600080604083850312156116cf57600080fd5b60006116dd858286016115ef565b92505060206116ee85828601611619565b9150509250929050565b60006020828403121561170a57600080fd5b600061171884828501611604565b91505092915050565b61172a81611b72565b82525050565b61173981611b84565b82525050565b61174881611b90565b82525050565b600061175982611a67565b6117638185611a72565b9350611773818560208601611bf0565b61177c81611c7c565b840191505092915050565b600061179282611a67565b61179c8185611a83565b93506117ac818560208601611bf0565b80840191505092915050565b60006117c5602083611a72565b91506117d082611c8d565b602082019050919050565b60006117e8601483611a72565b91506117f382611cb6565b602082019050919050565b600061180b602683611a72565b915061181682611cdf565b604082019050919050565b600061182e601083611a72565b915061183982611d2e565b602082019050919050565b6000611851602083611a72565b915061185c82611d57565b602082019050919050565b6000611874601783611a83565b915061187f82611d80565b601782019050919050565b6000611897601183611a83565b91506118a282611da9565b601182019050919050565b60006118ba602f83611a72565b91506118c582611dd2565b604082019050919050565b6118d981611be6565b82525050565b60006118ea82611867565b91506118f68285611787565b91506119018261188a565b915061190d8284611787565b91508190509392505050565b600060208201905061192e6000830184611721565b92915050565b60006020820190506119496000830184611730565b92915050565b6000602082019050611964600083018461173f565b92915050565b60006020820190508181036000830152611984818461174e565b905092915050565b600060208201905081810360008301526119a5816117b8565b9050919050565b600060208201905081810360008301526119c5816117db565b9050919050565b600060208201905081810360008301526119e5816117fe565b9050919050565b60006020820190508181036000830152611a0581611821565b9050919050565b60006020820190508181036000830152611a2581611844565b9050919050565b60006020820190508181036000830152611a45816118ad565b9050919050565b6000602082019050611a6160008301846118d0565b92915050565b600081519050919050565b600082825260208201905092915050565b600081905092915050565b6000611a9982611be6565b9150611aa483611be6565b9250827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff03821115611ad957611ad8611c4d565b5b828201905092915050565b6000611aef82611be6565b9150611afa83611be6565b9250817fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0483118215151615611b3357611b32611c4d565b5b828202905092915050565b6000611b4982611be6565b9150611b5483611be6565b925082821015611b6757611b66611c4d565b5b828203905092915050565b6000611b7d82611bc6565b9050919050565b60008115159050919050565b6000819050919050565b60007fffffffff0000000000000000000000000000000000000000000000000000000082169050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b60005b83811015611c0e578082015181840152602081019050611bf3565b83811115611c1d576000848401525b50505050565b6000611c2e82611be6565b91506000821415611c4257611c41611c4d565b5b600182039050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b6000601f19601f8301169050919050565b7f537472696e67733a20686578206c656e67746820696e73756666696369656e74600082015250565b7f5061757361626c653a206e6f7420706175736564000000000000000000000000600082015250565b7f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160008201527f6464726573730000000000000000000000000000000000000000000000000000602082015250565b7f5061757361626c653a2070617573656400000000000000000000000000000000600082015250565b7f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572600082015250565b7f416363657373436f6e74726f6c3a206163636f756e7420000000000000000000600082015250565b7f206973206d697373696e6720726f6c6520000000000000000000000000000000600082015250565b7f416363657373436f6e74726f6c3a2063616e206f6e6c792072656e6f756e636560008201527f20726f6c657320666f722073656c660000000000000000000000000000000000602082015250565b611e2a81611b72565b8114611e3557600080fd5b50565b611e4181611b90565b8114611e4c57600080fd5b50565b611e5881611b9a565b8114611e6357600080fd5b50565b611e6f81611be6565b8114611e7a57600080fd5b5056fea264697066735822122059e62badf4a2acd0033b47799a8ebedf124e63a3ddcb878ffc36865e3f88b27864736f6c63430008040033",
}

// KfiveAccessControlABI is the input ABI used to generate the binding from.
// Deprecated: Use KfiveAccessControlMetaData.ABI instead.
var KfiveAccessControlABI = KfiveAccessControlMetaData.ABI

// KfiveAccessControlBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use KfiveAccessControlMetaData.Bin instead.
var KfiveAccessControlBin = KfiveAccessControlMetaData.Bin

// DeployKfiveAccessControl deploys a new Ethereum contract, binding an instance of KfiveAccessControl to it.
func DeployKfiveAccessControl(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *KfiveAccessControl, error) {
	parsed, err := KfiveAccessControlMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(KfiveAccessControlBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &KfiveAccessControl{KfiveAccessControlCaller: KfiveAccessControlCaller{contract: contract}, KfiveAccessControlTransactor: KfiveAccessControlTransactor{contract: contract}, KfiveAccessControlFilterer: KfiveAccessControlFilterer{contract: contract}}, nil
}

// KfiveAccessControl is an auto generated Go binding around an Ethereum contract.
type KfiveAccessControl struct {
	KfiveAccessControlCaller     // Read-only binding to the contract
	KfiveAccessControlTransactor // Write-only binding to the contract
	KfiveAccessControlFilterer   // Log filterer for contract events
}

// KfiveAccessControlCaller is an auto generated read-only Go binding around an Ethereum contract.
type KfiveAccessControlCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KfiveAccessControlTransactor is an auto generated write-only Go binding around an Ethereum contract.
type KfiveAccessControlTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KfiveAccessControlFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type KfiveAccessControlFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KfiveAccessControlSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type KfiveAccessControlSession struct {
	Contract     *KfiveAccessControl // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// KfiveAccessControlCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type KfiveAccessControlCallerSession struct {
	Contract *KfiveAccessControlCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// KfiveAccessControlTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type KfiveAccessControlTransactorSession struct {
	Contract     *KfiveAccessControlTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// KfiveAccessControlRaw is an auto generated low-level Go binding around an Ethereum contract.
type KfiveAccessControlRaw struct {
	Contract *KfiveAccessControl // Generic contract binding to access the raw methods on
}

// KfiveAccessControlCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type KfiveAccessControlCallerRaw struct {
	Contract *KfiveAccessControlCaller // Generic read-only contract binding to access the raw methods on
}

// KfiveAccessControlTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type KfiveAccessControlTransactorRaw struct {
	Contract *KfiveAccessControlTransactor // Generic write-only contract binding to access the raw methods on
}

// NewKfiveAccessControl creates a new instance of KfiveAccessControl, bound to a specific deployed contract.
func NewKfiveAccessControl(address common.Address, backend bind.ContractBackend) (*KfiveAccessControl, error) {
	contract, err := bindKfiveAccessControl(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &KfiveAccessControl{KfiveAccessControlCaller: KfiveAccessControlCaller{contract: contract}, KfiveAccessControlTransactor: KfiveAccessControlTransactor{contract: contract}, KfiveAccessControlFilterer: KfiveAccessControlFilterer{contract: contract}}, nil
}

// NewKfiveAccessControlCaller creates a new read-only instance of KfiveAccessControl, bound to a specific deployed contract.
func NewKfiveAccessControlCaller(address common.Address, caller bind.ContractCaller) (*KfiveAccessControlCaller, error) {
	contract, err := bindKfiveAccessControl(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &KfiveAccessControlCaller{contract: contract}, nil
}

// NewKfiveAccessControlTransactor creates a new write-only instance of KfiveAccessControl, bound to a specific deployed contract.
func NewKfiveAccessControlTransactor(address common.Address, transactor bind.ContractTransactor) (*KfiveAccessControlTransactor, error) {
	contract, err := bindKfiveAccessControl(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &KfiveAccessControlTransactor{contract: contract}, nil
}

// NewKfiveAccessControlFilterer creates a new log filterer instance of KfiveAccessControl, bound to a specific deployed contract.
func NewKfiveAccessControlFilterer(address common.Address, filterer bind.ContractFilterer) (*KfiveAccessControlFilterer, error) {
	contract, err := bindKfiveAccessControl(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &KfiveAccessControlFilterer{contract: contract}, nil
}

// bindKfiveAccessControl binds a generic wrapper to an already deployed contract.
func bindKfiveAccessControl(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(KfiveAccessControlABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_KfiveAccessControl *KfiveAccessControlRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _KfiveAccessControl.Contract.KfiveAccessControlCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_KfiveAccessControl *KfiveAccessControlRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _KfiveAccessControl.Contract.KfiveAccessControlTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_KfiveAccessControl *KfiveAccessControlRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _KfiveAccessControl.Contract.KfiveAccessControlTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_KfiveAccessControl *KfiveAccessControlCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _KfiveAccessControl.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_KfiveAccessControl *KfiveAccessControlTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _KfiveAccessControl.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_KfiveAccessControl *KfiveAccessControlTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _KfiveAccessControl.Contract.contract.Transact(opts, method, params...)
}

// ADMINROLE is a free data retrieval call binding the contract method 0x75b238fc.
//
// Solidity: function ADMIN_ROLE() view returns(bytes32)
func (_KfiveAccessControl *KfiveAccessControlCaller) ADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _KfiveAccessControl.contract.Call(opts, &out, "ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ADMINROLE is a free data retrieval call binding the contract method 0x75b238fc.
//
// Solidity: function ADMIN_ROLE() view returns(bytes32)
func (_KfiveAccessControl *KfiveAccessControlSession) ADMINROLE() ([32]byte, error) {
	return _KfiveAccessControl.Contract.ADMINROLE(&_KfiveAccessControl.CallOpts)
}

// ADMINROLE is a free data retrieval call binding the contract method 0x75b238fc.
//
// Solidity: function ADMIN_ROLE() view returns(bytes32)
func (_KfiveAccessControl *KfiveAccessControlCallerSession) ADMINROLE() ([32]byte, error) {
	return _KfiveAccessControl.Contract.ADMINROLE(&_KfiveAccessControl.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_KfiveAccessControl *KfiveAccessControlCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _KfiveAccessControl.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_KfiveAccessControl *KfiveAccessControlSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _KfiveAccessControl.Contract.DEFAULTADMINROLE(&_KfiveAccessControl.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_KfiveAccessControl *KfiveAccessControlCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _KfiveAccessControl.Contract.DEFAULTADMINROLE(&_KfiveAccessControl.CallOpts)
}

// PAUSERROLE is a free data retrieval call binding the contract method 0xe63ab1e9.
//
// Solidity: function PAUSER_ROLE() view returns(bytes32)
func (_KfiveAccessControl *KfiveAccessControlCaller) PAUSERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _KfiveAccessControl.contract.Call(opts, &out, "PAUSER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// PAUSERROLE is a free data retrieval call binding the contract method 0xe63ab1e9.
//
// Solidity: function PAUSER_ROLE() view returns(bytes32)
func (_KfiveAccessControl *KfiveAccessControlSession) PAUSERROLE() ([32]byte, error) {
	return _KfiveAccessControl.Contract.PAUSERROLE(&_KfiveAccessControl.CallOpts)
}

// PAUSERROLE is a free data retrieval call binding the contract method 0xe63ab1e9.
//
// Solidity: function PAUSER_ROLE() view returns(bytes32)
func (_KfiveAccessControl *KfiveAccessControlCallerSession) PAUSERROLE() ([32]byte, error) {
	return _KfiveAccessControl.Contract.PAUSERROLE(&_KfiveAccessControl.CallOpts)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_KfiveAccessControl *KfiveAccessControlCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _KfiveAccessControl.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_KfiveAccessControl *KfiveAccessControlSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _KfiveAccessControl.Contract.GetRoleAdmin(&_KfiveAccessControl.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_KfiveAccessControl *KfiveAccessControlCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _KfiveAccessControl.Contract.GetRoleAdmin(&_KfiveAccessControl.CallOpts, role)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_KfiveAccessControl *KfiveAccessControlCaller) GetRoleMember(opts *bind.CallOpts, role [32]byte, index *big.Int) (common.Address, error) {
	var out []interface{}
	err := _KfiveAccessControl.contract.Call(opts, &out, "getRoleMember", role, index)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_KfiveAccessControl *KfiveAccessControlSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _KfiveAccessControl.Contract.GetRoleMember(&_KfiveAccessControl.CallOpts, role, index)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_KfiveAccessControl *KfiveAccessControlCallerSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _KfiveAccessControl.Contract.GetRoleMember(&_KfiveAccessControl.CallOpts, role, index)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_KfiveAccessControl *KfiveAccessControlCaller) GetRoleMemberCount(opts *bind.CallOpts, role [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _KfiveAccessControl.contract.Call(opts, &out, "getRoleMemberCount", role)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_KfiveAccessControl *KfiveAccessControlSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _KfiveAccessControl.Contract.GetRoleMemberCount(&_KfiveAccessControl.CallOpts, role)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_KfiveAccessControl *KfiveAccessControlCallerSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _KfiveAccessControl.Contract.GetRoleMemberCount(&_KfiveAccessControl.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_KfiveAccessControl *KfiveAccessControlCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _KfiveAccessControl.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_KfiveAccessControl *KfiveAccessControlSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _KfiveAccessControl.Contract.HasRole(&_KfiveAccessControl.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_KfiveAccessControl *KfiveAccessControlCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _KfiveAccessControl.Contract.HasRole(&_KfiveAccessControl.CallOpts, role, account)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_KfiveAccessControl *KfiveAccessControlCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _KfiveAccessControl.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_KfiveAccessControl *KfiveAccessControlSession) Owner() (common.Address, error) {
	return _KfiveAccessControl.Contract.Owner(&_KfiveAccessControl.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_KfiveAccessControl *KfiveAccessControlCallerSession) Owner() (common.Address, error) {
	return _KfiveAccessControl.Contract.Owner(&_KfiveAccessControl.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_KfiveAccessControl *KfiveAccessControlCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _KfiveAccessControl.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_KfiveAccessControl *KfiveAccessControlSession) Paused() (bool, error) {
	return _KfiveAccessControl.Contract.Paused(&_KfiveAccessControl.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_KfiveAccessControl *KfiveAccessControlCallerSession) Paused() (bool, error) {
	return _KfiveAccessControl.Contract.Paused(&_KfiveAccessControl.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_KfiveAccessControl *KfiveAccessControlCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _KfiveAccessControl.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_KfiveAccessControl *KfiveAccessControlSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _KfiveAccessControl.Contract.SupportsInterface(&_KfiveAccessControl.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_KfiveAccessControl *KfiveAccessControlCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _KfiveAccessControl.Contract.SupportsInterface(&_KfiveAccessControl.CallOpts, interfaceId)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_KfiveAccessControl *KfiveAccessControlTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _KfiveAccessControl.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_KfiveAccessControl *KfiveAccessControlSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _KfiveAccessControl.Contract.GrantRole(&_KfiveAccessControl.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_KfiveAccessControl *KfiveAccessControlTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _KfiveAccessControl.Contract.GrantRole(&_KfiveAccessControl.TransactOpts, role, account)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_KfiveAccessControl *KfiveAccessControlTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _KfiveAccessControl.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_KfiveAccessControl *KfiveAccessControlSession) Pause() (*types.Transaction, error) {
	return _KfiveAccessControl.Contract.Pause(&_KfiveAccessControl.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_KfiveAccessControl *KfiveAccessControlTransactorSession) Pause() (*types.Transaction, error) {
	return _KfiveAccessControl.Contract.Pause(&_KfiveAccessControl.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_KfiveAccessControl *KfiveAccessControlTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _KfiveAccessControl.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_KfiveAccessControl *KfiveAccessControlSession) RenounceOwnership() (*types.Transaction, error) {
	return _KfiveAccessControl.Contract.RenounceOwnership(&_KfiveAccessControl.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_KfiveAccessControl *KfiveAccessControlTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _KfiveAccessControl.Contract.RenounceOwnership(&_KfiveAccessControl.TransactOpts)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_KfiveAccessControl *KfiveAccessControlTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _KfiveAccessControl.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_KfiveAccessControl *KfiveAccessControlSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _KfiveAccessControl.Contract.RenounceRole(&_KfiveAccessControl.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_KfiveAccessControl *KfiveAccessControlTransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _KfiveAccessControl.Contract.RenounceRole(&_KfiveAccessControl.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_KfiveAccessControl *KfiveAccessControlTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _KfiveAccessControl.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_KfiveAccessControl *KfiveAccessControlSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _KfiveAccessControl.Contract.RevokeRole(&_KfiveAccessControl.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_KfiveAccessControl *KfiveAccessControlTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _KfiveAccessControl.Contract.RevokeRole(&_KfiveAccessControl.TransactOpts, role, account)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_KfiveAccessControl *KfiveAccessControlTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _KfiveAccessControl.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_KfiveAccessControl *KfiveAccessControlSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _KfiveAccessControl.Contract.TransferOwnership(&_KfiveAccessControl.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_KfiveAccessControl *KfiveAccessControlTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _KfiveAccessControl.Contract.TransferOwnership(&_KfiveAccessControl.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_KfiveAccessControl *KfiveAccessControlTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _KfiveAccessControl.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_KfiveAccessControl *KfiveAccessControlSession) Unpause() (*types.Transaction, error) {
	return _KfiveAccessControl.Contract.Unpause(&_KfiveAccessControl.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_KfiveAccessControl *KfiveAccessControlTransactorSession) Unpause() (*types.Transaction, error) {
	return _KfiveAccessControl.Contract.Unpause(&_KfiveAccessControl.TransactOpts)
}

// KfiveAccessControlOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the KfiveAccessControl contract.
type KfiveAccessControlOwnershipTransferredIterator struct {
	Event *KfiveAccessControlOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *KfiveAccessControlOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KfiveAccessControlOwnershipTransferred)
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
		it.Event = new(KfiveAccessControlOwnershipTransferred)
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
func (it *KfiveAccessControlOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KfiveAccessControlOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KfiveAccessControlOwnershipTransferred represents a OwnershipTransferred event raised by the KfiveAccessControl contract.
type KfiveAccessControlOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_KfiveAccessControl *KfiveAccessControlFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*KfiveAccessControlOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _KfiveAccessControl.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &KfiveAccessControlOwnershipTransferredIterator{contract: _KfiveAccessControl.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_KfiveAccessControl *KfiveAccessControlFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *KfiveAccessControlOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _KfiveAccessControl.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KfiveAccessControlOwnershipTransferred)
				if err := _KfiveAccessControl.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_KfiveAccessControl *KfiveAccessControlFilterer) ParseOwnershipTransferred(log types.Log) (*KfiveAccessControlOwnershipTransferred, error) {
	event := new(KfiveAccessControlOwnershipTransferred)
	if err := _KfiveAccessControl.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// KfiveAccessControlPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the KfiveAccessControl contract.
type KfiveAccessControlPausedIterator struct {
	Event *KfiveAccessControlPaused // Event containing the contract specifics and raw log

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
func (it *KfiveAccessControlPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KfiveAccessControlPaused)
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
		it.Event = new(KfiveAccessControlPaused)
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
func (it *KfiveAccessControlPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KfiveAccessControlPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KfiveAccessControlPaused represents a Paused event raised by the KfiveAccessControl contract.
type KfiveAccessControlPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_KfiveAccessControl *KfiveAccessControlFilterer) FilterPaused(opts *bind.FilterOpts) (*KfiveAccessControlPausedIterator, error) {

	logs, sub, err := _KfiveAccessControl.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &KfiveAccessControlPausedIterator{contract: _KfiveAccessControl.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_KfiveAccessControl *KfiveAccessControlFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *KfiveAccessControlPaused) (event.Subscription, error) {

	logs, sub, err := _KfiveAccessControl.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KfiveAccessControlPaused)
				if err := _KfiveAccessControl.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_KfiveAccessControl *KfiveAccessControlFilterer) ParsePaused(log types.Log) (*KfiveAccessControlPaused, error) {
	event := new(KfiveAccessControlPaused)
	if err := _KfiveAccessControl.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// KfiveAccessControlRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the KfiveAccessControl contract.
type KfiveAccessControlRoleAdminChangedIterator struct {
	Event *KfiveAccessControlRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *KfiveAccessControlRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KfiveAccessControlRoleAdminChanged)
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
		it.Event = new(KfiveAccessControlRoleAdminChanged)
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
func (it *KfiveAccessControlRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KfiveAccessControlRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KfiveAccessControlRoleAdminChanged represents a RoleAdminChanged event raised by the KfiveAccessControl contract.
type KfiveAccessControlRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_KfiveAccessControl *KfiveAccessControlFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*KfiveAccessControlRoleAdminChangedIterator, error) {

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

	logs, sub, err := _KfiveAccessControl.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &KfiveAccessControlRoleAdminChangedIterator{contract: _KfiveAccessControl.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_KfiveAccessControl *KfiveAccessControlFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *KfiveAccessControlRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _KfiveAccessControl.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KfiveAccessControlRoleAdminChanged)
				if err := _KfiveAccessControl.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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
func (_KfiveAccessControl *KfiveAccessControlFilterer) ParseRoleAdminChanged(log types.Log) (*KfiveAccessControlRoleAdminChanged, error) {
	event := new(KfiveAccessControlRoleAdminChanged)
	if err := _KfiveAccessControl.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// KfiveAccessControlRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the KfiveAccessControl contract.
type KfiveAccessControlRoleGrantedIterator struct {
	Event *KfiveAccessControlRoleGranted // Event containing the contract specifics and raw log

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
func (it *KfiveAccessControlRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KfiveAccessControlRoleGranted)
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
		it.Event = new(KfiveAccessControlRoleGranted)
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
func (it *KfiveAccessControlRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KfiveAccessControlRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KfiveAccessControlRoleGranted represents a RoleGranted event raised by the KfiveAccessControl contract.
type KfiveAccessControlRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_KfiveAccessControl *KfiveAccessControlFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*KfiveAccessControlRoleGrantedIterator, error) {

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

	logs, sub, err := _KfiveAccessControl.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &KfiveAccessControlRoleGrantedIterator{contract: _KfiveAccessControl.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_KfiveAccessControl *KfiveAccessControlFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *KfiveAccessControlRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _KfiveAccessControl.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KfiveAccessControlRoleGranted)
				if err := _KfiveAccessControl.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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
func (_KfiveAccessControl *KfiveAccessControlFilterer) ParseRoleGranted(log types.Log) (*KfiveAccessControlRoleGranted, error) {
	event := new(KfiveAccessControlRoleGranted)
	if err := _KfiveAccessControl.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// KfiveAccessControlRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the KfiveAccessControl contract.
type KfiveAccessControlRoleRevokedIterator struct {
	Event *KfiveAccessControlRoleRevoked // Event containing the contract specifics and raw log

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
func (it *KfiveAccessControlRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KfiveAccessControlRoleRevoked)
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
		it.Event = new(KfiveAccessControlRoleRevoked)
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
func (it *KfiveAccessControlRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KfiveAccessControlRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KfiveAccessControlRoleRevoked represents a RoleRevoked event raised by the KfiveAccessControl contract.
type KfiveAccessControlRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_KfiveAccessControl *KfiveAccessControlFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*KfiveAccessControlRoleRevokedIterator, error) {

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

	logs, sub, err := _KfiveAccessControl.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &KfiveAccessControlRoleRevokedIterator{contract: _KfiveAccessControl.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_KfiveAccessControl *KfiveAccessControlFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *KfiveAccessControlRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _KfiveAccessControl.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KfiveAccessControlRoleRevoked)
				if err := _KfiveAccessControl.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
func (_KfiveAccessControl *KfiveAccessControlFilterer) ParseRoleRevoked(log types.Log) (*KfiveAccessControlRoleRevoked, error) {
	event := new(KfiveAccessControlRoleRevoked)
	if err := _KfiveAccessControl.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// KfiveAccessControlUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the KfiveAccessControl contract.
type KfiveAccessControlUnpausedIterator struct {
	Event *KfiveAccessControlUnpaused // Event containing the contract specifics and raw log

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
func (it *KfiveAccessControlUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KfiveAccessControlUnpaused)
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
		it.Event = new(KfiveAccessControlUnpaused)
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
func (it *KfiveAccessControlUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KfiveAccessControlUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KfiveAccessControlUnpaused represents a Unpaused event raised by the KfiveAccessControl contract.
type KfiveAccessControlUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_KfiveAccessControl *KfiveAccessControlFilterer) FilterUnpaused(opts *bind.FilterOpts) (*KfiveAccessControlUnpausedIterator, error) {

	logs, sub, err := _KfiveAccessControl.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &KfiveAccessControlUnpausedIterator{contract: _KfiveAccessControl.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_KfiveAccessControl *KfiveAccessControlFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *KfiveAccessControlUnpaused) (event.Subscription, error) {

	logs, sub, err := _KfiveAccessControl.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KfiveAccessControlUnpaused)
				if err := _KfiveAccessControl.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
func (_KfiveAccessControl *KfiveAccessControlFilterer) ParseUnpaused(log types.Log) (*KfiveAccessControlUnpaused, error) {
	event := new(KfiveAccessControlUnpaused)
	if err := _KfiveAccessControl.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
