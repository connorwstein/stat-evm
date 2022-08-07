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
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"message\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"res\",\"type\":\"int256\"}],\"name\":\"Debug\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"fitted\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getFitted\",\"outputs\":[{\"internalType\":\"int256[][]\",\"name\":\"\",\"type\":\"int256[][]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getLastSample\",\"outputs\":[{\"internalType\":\"int256[]\",\"name\":\"\",\"type\":\"int256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMatrixMulti\",\"outputs\":[{\"internalType\":\"int256[][]\",\"name\":\"\",\"type\":\"int256[][]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"med\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"moment\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"prediction\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"product\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"sample\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"fitType\",\"type\":\"uint256\"},{\"internalType\":\"int256[][]\",\"name\":\"X\",\"type\":\"int256[][]\"},{\"internalType\":\"int256[][]\",\"name\":\"Y\",\"type\":\"int256[][]\"}],\"name\":\"testFit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int256[][]\",\"name\":\"a\",\"type\":\"int256[][]\"},{\"internalType\":\"int256[][]\",\"name\":\"b\",\"type\":\"int256[][]\"}],\"name\":\"testMatrixMult\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"vals\",\"type\":\"uint256[]\"}],\"name\":\"testMedian\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"v1\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"v2\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"v3\",\"type\":\"uint256[]\"}],\"name\":\"testMoment\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"steps\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"samples\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"targetTime\",\"type\":\"uint256\"}],\"name\":\"testPrediction\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"distributionType\",\"type\":\"uint256\"},{\"internalType\":\"int256\",\"name\":\"a\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"b\",\"type\":\"int256\"},{\"internalType\":\"uint256\",\"name\":\"numSamples\",\"type\":\"uint256\"}],\"name\":\"testSampler\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x6080604052730300000000000000000000000000000000000001600660006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550730300000000000000000000000000000000000004600760006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550730300000000000000000000000000000000000006600860006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550730300000000000000000000000000000000000005600960006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550730300000000000000000000000000000000000007600a60006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550730300000000000000000000000000000000000008600b60006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555034801561020e57600080fd5b506118008061021e6000396000f3fe608060405234801561001057600080fd5b50600436106100f55760003560e01c8063ba88c3b711610097578063f16d968c11610066578063f16d968c14610272578063f213a5a114610290578063fb7f772b146102ac578063fc784aa1146102c8576100f5565b8063ba88c3b7146101e8578063c44073ca14610218578063cd53c74714610236578063df98490114610254576100f5565b80633b1c29ef116100d35780633b1c29ef146101645780636ba875391461018057806382f044991461019c57806383b6033c146101cc576100f5565b806310fa1878146100fa57806312f758a01461012a5780632372c91614610146575b600080fd5b610114600480360381019061010f9190610ac0565b6102e6565b6040516101219190610b06565b60405180910390f35b610144600480360381019061013f9190610c7a565b61030a565b005b61014e6103b5565b60405161015b9190610d14565b60405180910390f35b61017e60048036038101906101799190610eff565b6103bb565b005b61019a60048036038101906101959190610f77565b610478565b005b6101b660048036038101906101b19190611002565b610538565b6040516101c39190610b06565b60405180910390f35b6101e660048036038101906101e19190611042565b610575565b005b61020260048036038101906101fd9190611002565b61061a565b60405161020f9190610b06565b60405180910390f35b610220610657565b60405161022d919061120b565b60405180910390f35b61023e6106f6565b60405161024b9190610d14565b60405180910390f35b61025c6106fc565b6040516102699190610d14565b60405180910390f35b61027a610702565b604051610287919061129c565b60405180910390f35b6102aa60048036038101906102a591906112be565b61075a565b005b6102c660048036038101906102c19190611325565b61081d565b005b6102d06108c8565b6040516102dd919061120b565b60405180910390f35b600281815481106102f657600080fd5b906000526020600020016000915090505481565b600860009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166320db7f0f8484846040518463ffffffff1660e01b815260040161036993929190611436565b602060405180830381865afa158015610386573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906103aa9190611490565b600181905550505050565b60055481565b600960009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166359c6945683836040518363ffffffff1660e01b81526004016104189291906114bd565b600060405180830381865afa158015610435573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f8201168201806040525081019061045e9190611655565b60039080519060200190610473929190610967565b505050565b600a60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663d3a5d0db8484846040518463ffffffff1660e01b81526004016104d79392919061169e565b600060405180830381865afa1580156104f4573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f8201168201806040525081019061051d9190611655565b60049080519060200190610532929190610967565b50505050565b6004828154811061054857600080fd5b90600052602060002001818154811061056057600080fd5b90600052602060002001600091509150505481565b600660009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663e8787f41826040518263ffffffff1660e01b81526004016105d091906116e3565b602060405180830381865afa1580156105ed573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906106119190611490565b60008190555050565b6003828154811061062a57600080fd5b90600052602060002001818154811061064257600080fd5b90600052602060002001600091509150505481565b60606004805480602002602001604051908101604052809291908181526020016000905b828210156106ed578382906000526020600020018054806020026020016040519081016040528092919081815260200182805480156106d957602002820191906000526020600020905b8154815260200190600101908083116106c5575b50505050508152602001906001019061067b565b50505050905090565b60005481565b60015481565b6060600280548060200260200160405190810160405280929190818152602001828054801561075057602002820191906000526020600020905b81548152602001906001019080831161073c575b5050505050905090565b600760009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663cca718bf858585856040518563ffffffff1660e01b81526004016107bb9493929190611705565b600060405180830381865afa1580156107d8573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f82011682018060405250810190610801919061174a565b600290805190602001906108169291906109c7565b5050505050565b600b60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16634eaf19a78484846040518463ffffffff1660e01b815260040161087c93929190611793565b602060405180830381865afa158015610899573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906108bd9190611490565b600581905550505050565b60606003805480602002602001604051908101604052809291908181526020016000905b8282101561095e5783829060005260206000200180548060200260200160405190810160405280929190818152602001828054801561094a57602002820191906000526020600020905b815481526020019060010190808311610936575b5050505050815260200190600101906108ec565b50505050905090565b8280548282559060005260206000209081019282156109b6579160200282015b828111156109b55782518290805190602001906109a59291906109c7565b5091602001919060010190610987565b5b5090506109c39190610a14565b5090565b828054828255906000526020600020908101928215610a03579160200282015b82811115610a025782518255916020019190600101906109e7565b5b509050610a109190610a38565b5090565b5b80821115610a345760008181610a2b9190610a55565b50600101610a15565b5090565b5b80821115610a51576000816000905550600101610a39565b5090565b5080546000825590600052602060002090810190610a739190610a38565b50565b6000604051905090565b600080fd5b600080fd5b6000819050919050565b610a9d81610a8a565b8114610aa857600080fd5b50565b600081359050610aba81610a94565b92915050565b600060208284031215610ad657610ad5610a80565b5b6000610ae484828501610aab565b91505092915050565b6000819050919050565b610b0081610aed565b82525050565b6000602082019050610b1b6000830184610af7565b92915050565b600080fd5b6000601f19601f8301169050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b610b6f82610b26565b810181811067ffffffffffffffff82111715610b8e57610b8d610b37565b5b80604052505050565b6000610ba1610a76565b9050610bad8282610b66565b919050565b600067ffffffffffffffff821115610bcd57610bcc610b37565b5b602082029050602081019050919050565b600080fd5b6000610bf6610bf184610bb2565b610b97565b90508083825260208201905060208402830185811115610c1957610c18610bde565b5b835b81811015610c425780610c2e8882610aab565b845260208401935050602081019050610c1b565b5050509392505050565b600082601f830112610c6157610c60610b21565b5b8135610c71848260208601610be3565b91505092915050565b600080600060608486031215610c9357610c92610a80565b5b6000610ca186828701610aab565b935050602084013567ffffffffffffffff811115610cc257610cc1610a85565b5b610cce86828701610c4c565b925050604084013567ffffffffffffffff811115610cef57610cee610a85565b5b610cfb86828701610c4c565b9150509250925092565b610d0e81610a8a565b82525050565b6000602082019050610d296000830184610d05565b92915050565b600067ffffffffffffffff821115610d4a57610d49610b37565b5b602082029050602081019050919050565b600067ffffffffffffffff821115610d7657610d75610b37565b5b602082029050602081019050919050565b610d9081610aed565b8114610d9b57600080fd5b50565b600081359050610dad81610d87565b92915050565b6000610dc6610dc184610d5b565b610b97565b90508083825260208201905060208402830185811115610de957610de8610bde565b5b835b81811015610e125780610dfe8882610d9e565b845260208401935050602081019050610deb565b5050509392505050565b600082601f830112610e3157610e30610b21565b5b8135610e41848260208601610db3565b91505092915050565b6000610e5d610e5884610d2f565b610b97565b90508083825260208201905060208402830185811115610e8057610e7f610bde565b5b835b81811015610ec757803567ffffffffffffffff811115610ea557610ea4610b21565b5b808601610eb28982610e1c565b85526020850194505050602081019050610e82565b5050509392505050565b600082601f830112610ee657610ee5610b21565b5b8135610ef6848260208601610e4a565b91505092915050565b60008060408385031215610f1657610f15610a80565b5b600083013567ffffffffffffffff811115610f3457610f33610a85565b5b610f4085828601610ed1565b925050602083013567ffffffffffffffff811115610f6157610f60610a85565b5b610f6d85828601610ed1565b9150509250929050565b600080600060608486031215610f9057610f8f610a80565b5b6000610f9e86828701610aab565b935050602084013567ffffffffffffffff811115610fbf57610fbe610a85565b5b610fcb86828701610ed1565b925050604084013567ffffffffffffffff811115610fec57610feb610a85565b5b610ff886828701610ed1565b9150509250925092565b6000806040838503121561101957611018610a80565b5b600061102785828601610aab565b925050602061103885828601610aab565b9150509250929050565b60006020828403121561105857611057610a80565b5b600082013567ffffffffffffffff81111561107657611075610a85565b5b61108284828501610c4c565b91505092915050565b600081519050919050565b600082825260208201905092915050565b6000819050602082019050919050565b600081519050919050565b600082825260208201905092915050565b6000819050602082019050919050565b6110ec81610aed565b82525050565b60006110fe83836110e3565b60208301905092915050565b6000602082019050919050565b6000611122826110b7565b61112c81856110c2565b9350611137836110d3565b8060005b8381101561116857815161114f88826110f2565b975061115a8361110a565b92505060018101905061113b565b5085935050505092915050565b60006111818383611117565b905092915050565b6000602082019050919050565b60006111a18261108b565b6111ab8185611096565b9350836020820285016111bd856110a7565b8060005b858110156111f957848403895281516111da8582611175565b94506111e583611189565b925060208a019950506001810190506111c1565b50829750879550505050505092915050565b600060208201905081810360008301526112258184611196565b905092915050565b600082825260208201905092915050565b6000611249826110b7565b611253818561122d565b935061125e836110d3565b8060005b8381101561128f57815161127688826110f2565b97506112818361110a565b925050600181019050611262565b5085935050505092915050565b600060208201905081810360008301526112b6818461123e565b905092915050565b600080600080608085870312156112d8576112d7610a80565b5b60006112e687828801610aab565b94505060206112f787828801610d9e565b935050604061130887828801610d9e565b925050606061131987828801610aab565b91505092959194509250565b60008060006060848603121561133e5761133d610a80565b5b600061134c86828701610aab565b935050602061135d86828701610aab565b925050604061136e86828701610aab565b9150509250925092565b600081519050919050565b600082825260208201905092915050565b6000819050602082019050919050565b6113ad81610a8a565b82525050565b60006113bf83836113a4565b60208301905092915050565b6000602082019050919050565b60006113e382611378565b6113ed8185611383565b93506113f883611394565b8060005b8381101561142957815161141088826113b3565b975061141b836113cb565b9250506001810190506113fc565b5085935050505092915050565b600060608201905061144b6000830186610d05565b818103602083015261145d81856113d8565b9050818103604083015261147181846113d8565b9050949350505050565b60008151905061148a81610a94565b92915050565b6000602082840312156114a6576114a5610a80565b5b60006114b48482850161147b565b91505092915050565b600060408201905081810360008301526114d78185611196565b905081810360208301526114eb8184611196565b90509392505050565b60008151905061150381610d87565b92915050565b600061151c61151784610d5b565b610b97565b9050808382526020820190506020840283018581111561153f5761153e610bde565b5b835b81811015611568578061155488826114f4565b845260208401935050602081019050611541565b5050509392505050565b600082601f83011261158757611586610b21565b5b8151611597848260208601611509565b91505092915050565b60006115b36115ae84610d2f565b610b97565b905080838252602082019050602084028301858111156115d6576115d5610bde565b5b835b8181101561161d57805167ffffffffffffffff8111156115fb576115fa610b21565b5b8086016116088982611572565b855260208501945050506020810190506115d8565b5050509392505050565b600082601f83011261163c5761163b610b21565b5b815161164c8482602086016115a0565b91505092915050565b60006020828403121561166b5761166a610a80565b5b600082015167ffffffffffffffff81111561168957611688610a85565b5b61169584828501611627565b91505092915050565b60006060820190506116b36000830186610d05565b81810360208301526116c58185611196565b905081810360408301526116d98184611196565b9050949350505050565b600060208201905081810360008301526116fd81846113d8565b905092915050565b600060808201905061171a6000830187610d05565b6117276020830186610af7565b6117346040830185610af7565b6117416060830184610d05565b95945050505050565b6000602082840312156117605761175f610a80565b5b600082015167ffffffffffffffff81111561177e5761177d610a85565b5b61178a84828501611572565b91505092915050565b60006060820190506117a86000830186610d05565b6117b56020830185610d05565b6117c26040830184610d05565b94935050505056fea2646970667358221220d3cac845bd77710e6acb26bc92fc1fa1bf72e8e39d46e9dd00ce9ee5438de15164736f6c634300080f0033",
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

// Prediction is a free data retrieval call binding the contract method 0x2372c916.
//
// Solidity: function prediction() view returns(uint256)
func (_Main *MainCaller) Prediction(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "prediction")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Prediction is a free data retrieval call binding the contract method 0x2372c916.
//
// Solidity: function prediction() view returns(uint256)
func (_Main *MainSession) Prediction() (*big.Int, error) {
	return _Main.Contract.Prediction(&_Main.CallOpts)
}

// Prediction is a free data retrieval call binding the contract method 0x2372c916.
//
// Solidity: function prediction() view returns(uint256)
func (_Main *MainCallerSession) Prediction() (*big.Int, error) {
	return _Main.Contract.Prediction(&_Main.CallOpts)
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

// TestPrediction is a paid mutator transaction binding the contract method 0xfb7f772b.
//
// Solidity: function testPrediction(uint256 steps, uint256 samples, uint256 targetTime) returns()
func (_Main *MainTransactor) TestPrediction(opts *bind.TransactOpts, steps *big.Int, samples *big.Int, targetTime *big.Int) (*types.Transaction, error) {
	return _Main.contract.Transact(opts, "testPrediction", steps, samples, targetTime)
}

// TestPrediction is a paid mutator transaction binding the contract method 0xfb7f772b.
//
// Solidity: function testPrediction(uint256 steps, uint256 samples, uint256 targetTime) returns()
func (_Main *MainSession) TestPrediction(steps *big.Int, samples *big.Int, targetTime *big.Int) (*types.Transaction, error) {
	return _Main.Contract.TestPrediction(&_Main.TransactOpts, steps, samples, targetTime)
}

// TestPrediction is a paid mutator transaction binding the contract method 0xfb7f772b.
//
// Solidity: function testPrediction(uint256 steps, uint256 samples, uint256 targetTime) returns()
func (_Main *MainTransactorSession) TestPrediction(steps *big.Int, samples *big.Int, targetTime *big.Int) (*types.Transaction, error) {
	return _Main.Contract.TestPrediction(&_Main.TransactOpts, steps, samples, targetTime)
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

