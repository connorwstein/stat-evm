// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package main

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

// MainMetaData contains all meta data concerning the Main contract.
var MainMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"message\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"res\",\"type\":\"int256\"}],\"name\":\"Debug\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"fitted\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getFitted\",\"outputs\":[{\"internalType\":\"int256[][]\",\"name\":\"\",\"type\":\"int256[][]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getLastSample\",\"outputs\":[{\"internalType\":\"int256[]\",\"name\":\"\",\"type\":\"int256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMatrixMulti\",\"outputs\":[{\"internalType\":\"int256[][]\",\"name\":\"\",\"type\":\"int256[][]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ipfsMomentRes\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"med\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"moment\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"product\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"sample\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"fitType\",\"type\":\"uint256\"},{\"internalType\":\"int256[][]\",\"name\":\"X\",\"type\":\"int256[][]\"},{\"internalType\":\"int256[][]\",\"name\":\"Y\",\"type\":\"int256[][]\"}],\"name\":\"testFit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"ipfsHash\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"moment\",\"type\":\"uint256\"}],\"name\":\"testIPFSMoment\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int256[][]\",\"name\":\"a\",\"type\":\"int256[][]\"},{\"internalType\":\"int256[][]\",\"name\":\"b\",\"type\":\"int256[][]\"}],\"name\":\"testMatrixMult\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"vals\",\"type\":\"uint256[]\"}],\"name\":\"testMedian\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"v1\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"v2\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"v3\",\"type\":\"uint256[]\"}],\"name\":\"testMoment\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"distributionType\",\"type\":\"uint256\"},{\"internalType\":\"int256\",\"name\":\"a\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"b\",\"type\":\"int256\"},{\"internalType\":\"uint256\",\"name\":\"numSamples\",\"type\":\"uint256\"}],\"name\":\"testSampler\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x6080604052730300000000000000000000000000000000000001600660006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550730300000000000000000000000000000000000004600760006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550730300000000000000000000000000000000000006600860006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550730300000000000000000000000000000000000005600960006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550730300000000000000000000000000000000000008600a60006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550730300000000000000000000000000000000000007600b60006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555034801561020e57600080fd5b5061193c8061021e6000396000f3fe608060405234801561001057600080fd5b50600436106100f55760003560e01c806383b6033c11610097578063df98490111610066578063df98490114610270578063f16d968c1461028e578063f213a5a1146102ac578063fc784aa1146102c8576100f5565b806383b6033c146101e8578063ba88c3b714610204578063c44073ca14610234578063cd53c74714610252576100f5565b80633b1c29ef116100d35780633b1c29ef146101645780635fd70772146101805780636ba875391461019c57806382f04499146101b8576100f5565b806310fa1878146100fa57806312f758a01461012a578063385d763f14610146575b600080fd5b610114600480360381019061010f9190610abd565b6102e6565b6040516101219190610b03565b60405180910390f35b610144600480360381019061013f9190610c77565b61030a565b005b61014e6103b5565b60405161015b9190610d11565b60405180910390f35b61017e60048036038101906101799190610efc565b6103bb565b005b61019a60048036038101906101959190611029565b610478565b005b6101b660048036038101906101b19190611085565b610520565b005b6101d260048036038101906101cd9190611110565b6105e0565b6040516101df9190610b03565b60405180910390f35b61020260048036038101906101fd9190611150565b61061d565b005b61021e60048036038101906102199190611110565b6106c2565b60405161022b9190610b03565b60405180910390f35b61023c6106ff565b6040516102499190611319565b60405180910390f35b61025a61079e565b6040516102679190610d11565b60405180910390f35b6102786107a4565b6040516102859190610d11565b60405180910390f35b6102966107aa565b6040516102a391906113aa565b60405180910390f35b6102c660048036038101906102c191906113cc565b610802565b005b6102d06108c5565b6040516102dd9190611319565b60405180910390f35b600281815481106102f657600080fd5b906000526020600020016000915090505481565b600860009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166320db7f0f8484846040518463ffffffff1660e01b8152600401610369939291906114f1565b602060405180830381865afa158015610386573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906103aa919061154b565b600181905550505050565b60055481565b600960009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166359c6945683836040518363ffffffff1660e01b8152600401610418929190611578565b600060405180830381865afa158015610435573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f8201168201806040525081019061045e9190611710565b60039080519060200190610473929190610964565b505050565b600a60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663a967f50383836040518363ffffffff1660e01b81526004016104d59291906117e1565b602060405180830381865afa1580156104f2573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610516919061154b565b6005819055505050565b600b60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663d3a5d0db8484846040518463ffffffff1660e01b815260040161057f93929190611811565b600060405180830381865afa15801561059c573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f820116820180604052508101906105c59190611710565b600490805190602001906105da929190610964565b50505050565b600482815481106105f057600080fd5b90600052602060002001818154811061060857600080fd5b90600052602060002001600091509150505481565b600660009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663e8787f41826040518263ffffffff1660e01b81526004016106789190611856565b602060405180830381865afa158015610695573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906106b9919061154b565b60008190555050565b600382815481106106d257600080fd5b9060005260206000200181815481106106ea57600080fd5b90600052602060002001600091509150505481565b60606004805480602002602001604051908101604052809291908181526020016000905b828210156107955783829060005260206000200180548060200260200160405190810160405280929190818152602001828054801561078157602002820191906000526020600020905b81548152602001906001019080831161076d575b505050505081526020019060010190610723565b50505050905090565b60005481565b60015481565b606060028054806020026020016040519081016040528092919081815260200182805480156107f857602002820191906000526020600020905b8154815260200190600101908083116107e4575b5050505050905090565b600760009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663cca718bf858585856040518563ffffffff1660e01b81526004016108639493929190611878565b600060405180830381865afa158015610880573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f820116820180604052508101906108a991906118bd565b600290805190602001906108be9291906109c4565b5050505050565b60606003805480602002602001604051908101604052809291908181526020016000905b8282101561095b5783829060005260206000200180548060200260200160405190810160405280929190818152602001828054801561094757602002820191906000526020600020905b815481526020019060010190808311610933575b5050505050815260200190600101906108e9565b50505050905090565b8280548282559060005260206000209081019282156109b3579160200282015b828111156109b25782518290805190602001906109a29291906109c4565b5091602001919060010190610984565b5b5090506109c09190610a11565b5090565b828054828255906000526020600020908101928215610a00579160200282015b828111156109ff5782518255916020019190600101906109e4565b5b509050610a0d9190610a35565b5090565b5b80821115610a315760008181610a289190610a52565b50600101610a12565b5090565b5b80821115610a4e576000816000905550600101610a36565b5090565b5080546000825590600052602060002090810190610a709190610a35565b50565b6000604051905090565b600080fd5b600080fd5b6000819050919050565b610a9a81610a87565b8114610aa557600080fd5b50565b600081359050610ab781610a91565b92915050565b600060208284031215610ad357610ad2610a7d565b5b6000610ae184828501610aa8565b91505092915050565b6000819050919050565b610afd81610aea565b82525050565b6000602082019050610b186000830184610af4565b92915050565b600080fd5b6000601f19601f8301169050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b610b6c82610b23565b810181811067ffffffffffffffff82111715610b8b57610b8a610b34565b5b80604052505050565b6000610b9e610a73565b9050610baa8282610b63565b919050565b600067ffffffffffffffff821115610bca57610bc9610b34565b5b602082029050602081019050919050565b600080fd5b6000610bf3610bee84610baf565b610b94565b90508083825260208201905060208402830185811115610c1657610c15610bdb565b5b835b81811015610c3f5780610c2b8882610aa8565b845260208401935050602081019050610c18565b5050509392505050565b600082601f830112610c5e57610c5d610b1e565b5b8135610c6e848260208601610be0565b91505092915050565b600080600060608486031215610c9057610c8f610a7d565b5b6000610c9e86828701610aa8565b935050602084013567ffffffffffffffff811115610cbf57610cbe610a82565b5b610ccb86828701610c49565b925050604084013567ffffffffffffffff811115610cec57610ceb610a82565b5b610cf886828701610c49565b9150509250925092565b610d0b81610a87565b82525050565b6000602082019050610d266000830184610d02565b92915050565b600067ffffffffffffffff821115610d4757610d46610b34565b5b602082029050602081019050919050565b600067ffffffffffffffff821115610d7357610d72610b34565b5b602082029050602081019050919050565b610d8d81610aea565b8114610d9857600080fd5b50565b600081359050610daa81610d84565b92915050565b6000610dc3610dbe84610d58565b610b94565b90508083825260208201905060208402830185811115610de657610de5610bdb565b5b835b81811015610e0f5780610dfb8882610d9b565b845260208401935050602081019050610de8565b5050509392505050565b600082601f830112610e2e57610e2d610b1e565b5b8135610e3e848260208601610db0565b91505092915050565b6000610e5a610e5584610d2c565b610b94565b90508083825260208201905060208402830185811115610e7d57610e7c610bdb565b5b835b81811015610ec457803567ffffffffffffffff811115610ea257610ea1610b1e565b5b808601610eaf8982610e19565b85526020850194505050602081019050610e7f565b5050509392505050565b600082601f830112610ee357610ee2610b1e565b5b8135610ef3848260208601610e47565b91505092915050565b60008060408385031215610f1357610f12610a7d565b5b600083013567ffffffffffffffff811115610f3157610f30610a82565b5b610f3d85828601610ece565b925050602083013567ffffffffffffffff811115610f5e57610f5d610a82565b5b610f6a85828601610ece565b9150509250929050565b600080fd5b600067ffffffffffffffff821115610f9457610f93610b34565b5b610f9d82610b23565b9050602081019050919050565b82818337600083830152505050565b6000610fcc610fc784610f79565b610b94565b905082815260208101848484011115610fe857610fe7610f74565b5b610ff3848285610faa565b509392505050565b600082601f8301126110105761100f610b1e565b5b8135611020848260208601610fb9565b91505092915050565b600080604083850312156110405761103f610a7d565b5b600083013567ffffffffffffffff81111561105e5761105d610a82565b5b61106a85828601610ffb565b925050602061107b85828601610aa8565b9150509250929050565b60008060006060848603121561109e5761109d610a7d565b5b60006110ac86828701610aa8565b935050602084013567ffffffffffffffff8111156110cd576110cc610a82565b5b6110d986828701610ece565b925050604084013567ffffffffffffffff8111156110fa576110f9610a82565b5b61110686828701610ece565b9150509250925092565b6000806040838503121561112757611126610a7d565b5b600061113585828601610aa8565b925050602061114685828601610aa8565b9150509250929050565b60006020828403121561116657611165610a7d565b5b600082013567ffffffffffffffff81111561118457611183610a82565b5b61119084828501610c49565b91505092915050565b600081519050919050565b600082825260208201905092915050565b6000819050602082019050919050565b600081519050919050565b600082825260208201905092915050565b6000819050602082019050919050565b6111fa81610aea565b82525050565b600061120c83836111f1565b60208301905092915050565b6000602082019050919050565b6000611230826111c5565b61123a81856111d0565b9350611245836111e1565b8060005b8381101561127657815161125d8882611200565b975061126883611218565b925050600181019050611249565b5085935050505092915050565b600061128f8383611225565b905092915050565b6000602082019050919050565b60006112af82611199565b6112b981856111a4565b9350836020820285016112cb856111b5565b8060005b8581101561130757848403895281516112e88582611283565b94506112f383611297565b925060208a019950506001810190506112cf565b50829750879550505050505092915050565b6000602082019050818103600083015261133381846112a4565b905092915050565b600082825260208201905092915050565b6000611357826111c5565b611361818561133b565b935061136c836111e1565b8060005b8381101561139d5781516113848882611200565b975061138f83611218565b925050600181019050611370565b5085935050505092915050565b600060208201905081810360008301526113c4818461134c565b905092915050565b600080600080608085870312156113e6576113e5610a7d565b5b60006113f487828801610aa8565b945050602061140587828801610d9b565b935050604061141687828801610d9b565b925050606061142787828801610aa8565b91505092959194509250565b600081519050919050565b600082825260208201905092915050565b6000819050602082019050919050565b61146881610a87565b82525050565b600061147a838361145f565b60208301905092915050565b6000602082019050919050565b600061149e82611433565b6114a8818561143e565b93506114b38361144f565b8060005b838110156114e45781516114cb888261146e565b97506114d683611486565b9250506001810190506114b7565b5085935050505092915050565b60006060820190506115066000830186610d02565b81810360208301526115188185611493565b9050818103604083015261152c8184611493565b9050949350505050565b60008151905061154581610a91565b92915050565b60006020828403121561156157611560610a7d565b5b600061156f84828501611536565b91505092915050565b6000604082019050818103600083015261159281856112a4565b905081810360208301526115a681846112a4565b90509392505050565b6000815190506115be81610d84565b92915050565b60006115d76115d284610d58565b610b94565b905080838252602082019050602084028301858111156115fa576115f9610bdb565b5b835b81811015611623578061160f88826115af565b8452602084019350506020810190506115fc565b5050509392505050565b600082601f83011261164257611641610b1e565b5b81516116528482602086016115c4565b91505092915050565b600061166e61166984610d2c565b610b94565b9050808382526020820190506020840283018581111561169157611690610bdb565b5b835b818110156116d857805167ffffffffffffffff8111156116b6576116b5610b1e565b5b8086016116c3898261162d565b85526020850194505050602081019050611693565b5050509392505050565b600082601f8301126116f7576116f6610b1e565b5b815161170784826020860161165b565b91505092915050565b60006020828403121561172657611725610a7d565b5b600082015167ffffffffffffffff81111561174457611743610a82565b5b611750848285016116e2565b91505092915050565b600081519050919050565b600082825260208201905092915050565b60005b83811015611793578082015181840152602081019050611778565b838111156117a2576000848401525b50505050565b60006117b382611759565b6117bd8185611764565b93506117cd818560208601611775565b6117d681610b23565b840191505092915050565b600060408201905081810360008301526117fb81856117a8565b905061180a6020830184610d02565b9392505050565b60006060820190506118266000830186610d02565b818103602083015261183881856112a4565b9050818103604083015261184c81846112a4565b9050949350505050565b600060208201905081810360008301526118708184611493565b905092915050565b600060808201905061188d6000830187610d02565b61189a6020830186610af4565b6118a76040830185610af4565b6118b46060830184610d02565b95945050505050565b6000602082840312156118d3576118d2610a7d565b5b600082015167ffffffffffffffff8111156118f1576118f0610a82565b5b6118fd8482850161162d565b9150509291505056fea2646970667358221220ffa64e0b92e85d2751711411b14c92d84bd6bbc5fb5128404da86e4f93f447ad64736f6c634300080f0033",
}

// MainABI is the input ABI used to generate the binding from.
// Deprecated: Use MainMetaData.ABI instead.
var MainABI = MainMetaData.ABI

// MainBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use MainMetaData.Bin instead.
var MainBin = MainMetaData.Bin

// DeployMain deploys a new Ethereum contract, binding an instance of Main to it.
func DeployMain(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Main, error) {
	parsed, err := MainMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(MainBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Main{MainCaller: MainCaller{contract: contract}, MainTransactor: MainTransactor{contract: contract}, MainFilterer: MainFilterer{contract: contract}}, nil
}

// Main is an auto generated Go binding around an Ethereum contract.
type Main struct {
	MainCaller     // Read-only binding to the contract
	MainTransactor // Write-only binding to the contract
	MainFilterer   // Log filterer for contract events
}

// MainCaller is an auto generated read-only Go binding around an Ethereum contract.
type MainCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MainTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MainTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MainFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MainFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MainSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MainSession struct {
	Contract     *Main             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MainCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MainCallerSession struct {
	Contract *MainCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// MainTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MainTransactorSession struct {
	Contract     *MainTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MainRaw is an auto generated low-level Go binding around an Ethereum contract.
type MainRaw struct {
	Contract *Main // Generic contract binding to access the raw methods on
}

// MainCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MainCallerRaw struct {
	Contract *MainCaller // Generic read-only contract binding to access the raw methods on
}

// MainTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MainTransactorRaw struct {
	Contract *MainTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMain creates a new instance of Main, bound to a specific deployed contract.
func NewMain(address common.Address, backend bind.ContractBackend) (*Main, error) {
	contract, err := bindMain(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Main{MainCaller: MainCaller{contract: contract}, MainTransactor: MainTransactor{contract: contract}, MainFilterer: MainFilterer{contract: contract}}, nil
}

// NewMainCaller creates a new read-only instance of Main, bound to a specific deployed contract.
func NewMainCaller(address common.Address, caller bind.ContractCaller) (*MainCaller, error) {
	contract, err := bindMain(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MainCaller{contract: contract}, nil
}

// NewMainTransactor creates a new write-only instance of Main, bound to a specific deployed contract.
func NewMainTransactor(address common.Address, transactor bind.ContractTransactor) (*MainTransactor, error) {
	contract, err := bindMain(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MainTransactor{contract: contract}, nil
}

// NewMainFilterer creates a new log filterer instance of Main, bound to a specific deployed contract.
func NewMainFilterer(address common.Address, filterer bind.ContractFilterer) (*MainFilterer, error) {
	contract, err := bindMain(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MainFilterer{contract: contract}, nil
}

// bindMain binds a generic wrapper to an already deployed contract.
func bindMain(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MainABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Main *MainRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Main.Contract.MainCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Main *MainRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Main.Contract.MainTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Main *MainRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Main.Contract.MainTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Main *MainCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Main.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Main *MainTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Main.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Main *MainTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Main.Contract.contract.Transact(opts, method, params...)
}

// Fitted is a free data retrieval call binding the contract method 0x82f04499.
//
// Solidity: function fitted(uint256 , uint256 ) view returns(int256)
func (_Main *MainCaller) Fitted(opts *bind.CallOpts, arg0 *big.Int, arg1 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "fitted", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Fitted is a free data retrieval call binding the contract method 0x82f04499.
//
// Solidity: function fitted(uint256 , uint256 ) view returns(int256)
func (_Main *MainSession) Fitted(arg0 *big.Int, arg1 *big.Int) (*big.Int, error) {
	return _Main.Contract.Fitted(&_Main.CallOpts, arg0, arg1)
}

// Fitted is a free data retrieval call binding the contract method 0x82f04499.
//
// Solidity: function fitted(uint256 , uint256 ) view returns(int256)
func (_Main *MainCallerSession) Fitted(arg0 *big.Int, arg1 *big.Int) (*big.Int, error) {
	return _Main.Contract.Fitted(&_Main.CallOpts, arg0, arg1)
}

// GetFitted is a free data retrieval call binding the contract method 0xc44073ca.
//
// Solidity: function getFitted() view returns(int256[][])
func (_Main *MainCaller) GetFitted(opts *bind.CallOpts) ([][]*big.Int, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "getFitted")

	if err != nil {
		return *new([][]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([][]*big.Int)).(*[][]*big.Int)

	return out0, err

}

// GetFitted is a free data retrieval call binding the contract method 0xc44073ca.
//
// Solidity: function getFitted() view returns(int256[][])
func (_Main *MainSession) GetFitted() ([][]*big.Int, error) {
	return _Main.Contract.GetFitted(&_Main.CallOpts)
}

// GetFitted is a free data retrieval call binding the contract method 0xc44073ca.
//
// Solidity: function getFitted() view returns(int256[][])
func (_Main *MainCallerSession) GetFitted() ([][]*big.Int, error) {
	return _Main.Contract.GetFitted(&_Main.CallOpts)
}

// GetLastSample is a free data retrieval call binding the contract method 0xf16d968c.
//
// Solidity: function getLastSample() view returns(int256[])
func (_Main *MainCaller) GetLastSample(opts *bind.CallOpts) ([]*big.Int, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "getLastSample")

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetLastSample is a free data retrieval call binding the contract method 0xf16d968c.
//
// Solidity: function getLastSample() view returns(int256[])
func (_Main *MainSession) GetLastSample() ([]*big.Int, error) {
	return _Main.Contract.GetLastSample(&_Main.CallOpts)
}

// GetLastSample is a free data retrieval call binding the contract method 0xf16d968c.
//
// Solidity: function getLastSample() view returns(int256[])
func (_Main *MainCallerSession) GetLastSample() ([]*big.Int, error) {
	return _Main.Contract.GetLastSample(&_Main.CallOpts)
}

// GetMatrixMulti is a free data retrieval call binding the contract method 0xfc784aa1.
//
// Solidity: function getMatrixMulti() view returns(int256[][])
func (_Main *MainCaller) GetMatrixMulti(opts *bind.CallOpts) ([][]*big.Int, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "getMatrixMulti")

	if err != nil {
		return *new([][]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([][]*big.Int)).(*[][]*big.Int)

	return out0, err

}

// GetMatrixMulti is a free data retrieval call binding the contract method 0xfc784aa1.
//
// Solidity: function getMatrixMulti() view returns(int256[][])
func (_Main *MainSession) GetMatrixMulti() ([][]*big.Int, error) {
	return _Main.Contract.GetMatrixMulti(&_Main.CallOpts)
}

// GetMatrixMulti is a free data retrieval call binding the contract method 0xfc784aa1.
//
// Solidity: function getMatrixMulti() view returns(int256[][])
func (_Main *MainCallerSession) GetMatrixMulti() ([][]*big.Int, error) {
	return _Main.Contract.GetMatrixMulti(&_Main.CallOpts)
}

// IpfsMomentRes is a free data retrieval call binding the contract method 0x385d763f.
//
// Solidity: function ipfsMomentRes() view returns(uint256)
func (_Main *MainCaller) IpfsMomentRes(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "ipfsMomentRes")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// IpfsMomentRes is a free data retrieval call binding the contract method 0x385d763f.
//
// Solidity: function ipfsMomentRes() view returns(uint256)
func (_Main *MainSession) IpfsMomentRes() (*big.Int, error) {
	return _Main.Contract.IpfsMomentRes(&_Main.CallOpts)
}

// IpfsMomentRes is a free data retrieval call binding the contract method 0x385d763f.
//
// Solidity: function ipfsMomentRes() view returns(uint256)
func (_Main *MainCallerSession) IpfsMomentRes() (*big.Int, error) {
	return _Main.Contract.IpfsMomentRes(&_Main.CallOpts)
}

// Med is a free data retrieval call binding the contract method 0xcd53c747.
//
// Solidity: function med() view returns(uint256)
func (_Main *MainCaller) Med(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "med")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Med is a free data retrieval call binding the contract method 0xcd53c747.
//
// Solidity: function med() view returns(uint256)
func (_Main *MainSession) Med() (*big.Int, error) {
	return _Main.Contract.Med(&_Main.CallOpts)
}

// Med is a free data retrieval call binding the contract method 0xcd53c747.
//
// Solidity: function med() view returns(uint256)
func (_Main *MainCallerSession) Med() (*big.Int, error) {
	return _Main.Contract.Med(&_Main.CallOpts)
}

// Moment is a free data retrieval call binding the contract method 0xdf984901.
//
// Solidity: function moment() view returns(uint256)
func (_Main *MainCaller) Moment(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "moment")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Moment is a free data retrieval call binding the contract method 0xdf984901.
//
// Solidity: function moment() view returns(uint256)
func (_Main *MainSession) Moment() (*big.Int, error) {
	return _Main.Contract.Moment(&_Main.CallOpts)
}

// Moment is a free data retrieval call binding the contract method 0xdf984901.
//
// Solidity: function moment() view returns(uint256)
func (_Main *MainCallerSession) Moment() (*big.Int, error) {
	return _Main.Contract.Moment(&_Main.CallOpts)
}

// Product is a free data retrieval call binding the contract method 0xba88c3b7.
//
// Solidity: function product(uint256 , uint256 ) view returns(int256)
func (_Main *MainCaller) Product(opts *bind.CallOpts, arg0 *big.Int, arg1 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "product", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Product is a free data retrieval call binding the contract method 0xba88c3b7.
//
// Solidity: function product(uint256 , uint256 ) view returns(int256)
func (_Main *MainSession) Product(arg0 *big.Int, arg1 *big.Int) (*big.Int, error) {
	return _Main.Contract.Product(&_Main.CallOpts, arg0, arg1)
}

// Product is a free data retrieval call binding the contract method 0xba88c3b7.
//
// Solidity: function product(uint256 , uint256 ) view returns(int256)
func (_Main *MainCallerSession) Product(arg0 *big.Int, arg1 *big.Int) (*big.Int, error) {
	return _Main.Contract.Product(&_Main.CallOpts, arg0, arg1)
}

// Sample is a free data retrieval call binding the contract method 0x10fa1878.
//
// Solidity: function sample(uint256 ) view returns(int256)
func (_Main *MainCaller) Sample(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "sample", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Sample is a free data retrieval call binding the contract method 0x10fa1878.
//
// Solidity: function sample(uint256 ) view returns(int256)
func (_Main *MainSession) Sample(arg0 *big.Int) (*big.Int, error) {
	return _Main.Contract.Sample(&_Main.CallOpts, arg0)
}

// Sample is a free data retrieval call binding the contract method 0x10fa1878.
//
// Solidity: function sample(uint256 ) view returns(int256)
func (_Main *MainCallerSession) Sample(arg0 *big.Int) (*big.Int, error) {
	return _Main.Contract.Sample(&_Main.CallOpts, arg0)
}

// TestFit is a paid mutator transaction binding the contract method 0x6ba87539.
//
// Solidity: function testFit(uint256 fitType, int256[][] X, int256[][] Y) returns()
func (_Main *MainTransactor) TestFit(opts *bind.TransactOpts, fitType *big.Int, X [][]*big.Int, Y [][]*big.Int) (*types.Transaction, error) {
	return _Main.contract.Transact(opts, "testFit", fitType, X, Y)
}

// TestFit is a paid mutator transaction binding the contract method 0x6ba87539.
//
// Solidity: function testFit(uint256 fitType, int256[][] X, int256[][] Y) returns()
func (_Main *MainSession) TestFit(fitType *big.Int, X [][]*big.Int, Y [][]*big.Int) (*types.Transaction, error) {
	return _Main.Contract.TestFit(&_Main.TransactOpts, fitType, X, Y)
}

// TestFit is a paid mutator transaction binding the contract method 0x6ba87539.
//
// Solidity: function testFit(uint256 fitType, int256[][] X, int256[][] Y) returns()
func (_Main *MainTransactorSession) TestFit(fitType *big.Int, X [][]*big.Int, Y [][]*big.Int) (*types.Transaction, error) {
	return _Main.Contract.TestFit(&_Main.TransactOpts, fitType, X, Y)
}

// TestIPFSMoment is a paid mutator transaction binding the contract method 0x5fd70772.
//
// Solidity: function testIPFSMoment(string ipfsHash, uint256 moment) returns()
func (_Main *MainTransactor) TestIPFSMoment(opts *bind.TransactOpts, ipfsHash string, moment *big.Int) (*types.Transaction, error) {
	return _Main.contract.Transact(opts, "testIPFSMoment", ipfsHash, moment)
}

// TestIPFSMoment is a paid mutator transaction binding the contract method 0x5fd70772.
//
// Solidity: function testIPFSMoment(string ipfsHash, uint256 moment) returns()
func (_Main *MainSession) TestIPFSMoment(ipfsHash string, moment *big.Int) (*types.Transaction, error) {
	return _Main.Contract.TestIPFSMoment(&_Main.TransactOpts, ipfsHash, moment)
}

// TestIPFSMoment is a paid mutator transaction binding the contract method 0x5fd70772.
//
// Solidity: function testIPFSMoment(string ipfsHash, uint256 moment) returns()
func (_Main *MainTransactorSession) TestIPFSMoment(ipfsHash string, moment *big.Int) (*types.Transaction, error) {
	return _Main.Contract.TestIPFSMoment(&_Main.TransactOpts, ipfsHash, moment)
}

// TestMatrixMult is a paid mutator transaction binding the contract method 0x3b1c29ef.
//
// Solidity: function testMatrixMult(int256[][] a, int256[][] b) returns()
func (_Main *MainTransactor) TestMatrixMult(opts *bind.TransactOpts, a [][]*big.Int, b [][]*big.Int) (*types.Transaction, error) {
	return _Main.contract.Transact(opts, "testMatrixMult", a, b)
}

// TestMatrixMult is a paid mutator transaction binding the contract method 0x3b1c29ef.
//
// Solidity: function testMatrixMult(int256[][] a, int256[][] b) returns()
func (_Main *MainSession) TestMatrixMult(a [][]*big.Int, b [][]*big.Int) (*types.Transaction, error) {
	return _Main.Contract.TestMatrixMult(&_Main.TransactOpts, a, b)
}

// TestMatrixMult is a paid mutator transaction binding the contract method 0x3b1c29ef.
//
// Solidity: function testMatrixMult(int256[][] a, int256[][] b) returns()
func (_Main *MainTransactorSession) TestMatrixMult(a [][]*big.Int, b [][]*big.Int) (*types.Transaction, error) {
	return _Main.Contract.TestMatrixMult(&_Main.TransactOpts, a, b)
}

// TestMedian is a paid mutator transaction binding the contract method 0x83b6033c.
//
// Solidity: function testMedian(uint256[] vals) returns()
func (_Main *MainTransactor) TestMedian(opts *bind.TransactOpts, vals []*big.Int) (*types.Transaction, error) {
	return _Main.contract.Transact(opts, "testMedian", vals)
}

// TestMedian is a paid mutator transaction binding the contract method 0x83b6033c.
//
// Solidity: function testMedian(uint256[] vals) returns()
func (_Main *MainSession) TestMedian(vals []*big.Int) (*types.Transaction, error) {
	return _Main.Contract.TestMedian(&_Main.TransactOpts, vals)
}

// TestMedian is a paid mutator transaction binding the contract method 0x83b6033c.
//
// Solidity: function testMedian(uint256[] vals) returns()
func (_Main *MainTransactorSession) TestMedian(vals []*big.Int) (*types.Transaction, error) {
	return _Main.Contract.TestMedian(&_Main.TransactOpts, vals)
}

// TestMoment is a paid mutator transaction binding the contract method 0x12f758a0.
//
// Solidity: function testMoment(uint256 v1, uint256[] v2, uint256[] v3) returns()
func (_Main *MainTransactor) TestMoment(opts *bind.TransactOpts, v1 *big.Int, v2 []*big.Int, v3 []*big.Int) (*types.Transaction, error) {
	return _Main.contract.Transact(opts, "testMoment", v1, v2, v3)
}

// TestMoment is a paid mutator transaction binding the contract method 0x12f758a0.
//
// Solidity: function testMoment(uint256 v1, uint256[] v2, uint256[] v3) returns()
func (_Main *MainSession) TestMoment(v1 *big.Int, v2 []*big.Int, v3 []*big.Int) (*types.Transaction, error) {
	return _Main.Contract.TestMoment(&_Main.TransactOpts, v1, v2, v3)
}

// TestMoment is a paid mutator transaction binding the contract method 0x12f758a0.
//
// Solidity: function testMoment(uint256 v1, uint256[] v2, uint256[] v3) returns()
func (_Main *MainTransactorSession) TestMoment(v1 *big.Int, v2 []*big.Int, v3 []*big.Int) (*types.Transaction, error) {
	return _Main.Contract.TestMoment(&_Main.TransactOpts, v1, v2, v3)
}

// TestSampler is a paid mutator transaction binding the contract method 0xf213a5a1.
//
// Solidity: function testSampler(uint256 distributionType, int256 a, int256 b, uint256 numSamples) returns()
func (_Main *MainTransactor) TestSampler(opts *bind.TransactOpts, distributionType *big.Int, a *big.Int, b *big.Int, numSamples *big.Int) (*types.Transaction, error) {
	return _Main.contract.Transact(opts, "testSampler", distributionType, a, b, numSamples)
}

// TestSampler is a paid mutator transaction binding the contract method 0xf213a5a1.
//
// Solidity: function testSampler(uint256 distributionType, int256 a, int256 b, uint256 numSamples) returns()
func (_Main *MainSession) TestSampler(distributionType *big.Int, a *big.Int, b *big.Int, numSamples *big.Int) (*types.Transaction, error) {
	return _Main.Contract.TestSampler(&_Main.TransactOpts, distributionType, a, b, numSamples)
}

// TestSampler is a paid mutator transaction binding the contract method 0xf213a5a1.
//
// Solidity: function testSampler(uint256 distributionType, int256 a, int256 b, uint256 numSamples) returns()
func (_Main *MainTransactorSession) TestSampler(distributionType *big.Int, a *big.Int, b *big.Int, numSamples *big.Int) (*types.Transaction, error) {
	return _Main.Contract.TestSampler(&_Main.TransactOpts, distributionType, a, b, numSamples)
}

// MainDebugIterator is returned from FilterDebug and is used to iterate over the raw logs and unpacked data for Debug events raised by the Main contract.
type MainDebugIterator struct {
	Event *MainDebug // Event containing the contract specifics and raw log

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
func (it *MainDebugIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MainDebug)
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
		it.Event = new(MainDebug)
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
func (it *MainDebugIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MainDebugIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MainDebug represents a Debug event raised by the Main contract.
type MainDebug struct {
	Message string
	Res     *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterDebug is a free log retrieval operation binding the contract event 0x08a92484e616affef9d47a59453c54d1381a7d985ca52d3c76e2ced243e3c9e8.
//
// Solidity: event Debug(string message, int256 res)
func (_Main *MainFilterer) FilterDebug(opts *bind.FilterOpts) (*MainDebugIterator, error) {

	logs, sub, err := _Main.contract.FilterLogs(opts, "Debug")
	if err != nil {
		return nil, err
	}
	return &MainDebugIterator{contract: _Main.contract, event: "Debug", logs: logs, sub: sub}, nil
}

// WatchDebug is a free log subscription operation binding the contract event 0x08a92484e616affef9d47a59453c54d1381a7d985ca52d3c76e2ced243e3c9e8.
//
// Solidity: event Debug(string message, int256 res)
func (_Main *MainFilterer) WatchDebug(opts *bind.WatchOpts, sink chan<- *MainDebug) (event.Subscription, error) {

	logs, sub, err := _Main.contract.WatchLogs(opts, "Debug")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MainDebug)
				if err := _Main.contract.UnpackLog(event, "Debug", log); err != nil {
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

// ParseDebug is a log parse operation binding the contract event 0x08a92484e616affef9d47a59453c54d1381a7d985ca52d3c76e2ced243e3c9e8.
//
// Solidity: event Debug(string message, int256 res)
func (_Main *MainFilterer) ParseDebug(log types.Log) (*MainDebug, error) {
	event := new(MainDebug)
	if err := _Main.contract.UnpackLog(event, "Debug", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

