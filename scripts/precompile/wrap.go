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
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"message\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"res\",\"type\":\"int256\"}],\"name\":\"Debug\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"getLastSample\",\"outputs\":[{\"internalType\":\"int256[]\",\"name\":\"\",\"type\":\"int256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMatrixMulti\",\"outputs\":[{\"internalType\":\"int256[][]\",\"name\":\"\",\"type\":\"int256[][]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ipfsMomentRes\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"med\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"moment\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"sample\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"ipfsHash\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"moment\",\"type\":\"uint256\"}],\"name\":\"testIPFSMoment\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int256[][]\",\"name\":\"a\",\"type\":\"int256[][]\"},{\"internalType\":\"int256[][]\",\"name\":\"b\",\"type\":\"int256[][]\"}],\"name\":\"testMatrixMult\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"vals\",\"type\":\"uint256[]\"}],\"name\":\"testMedian\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"v1\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"v2\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"v3\",\"type\":\"uint256[]\"}],\"name\":\"testMoment\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"distributionType\",\"type\":\"uint256\"},{\"internalType\":\"int256\",\"name\":\"a\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"b\",\"type\":\"int256\"},{\"internalType\":\"uint256\",\"name\":\"numSamples\",\"type\":\"uint256\"}],\"name\":\"testSampler\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x6080604052730300000000000000000000000000000000000001600560006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550730300000000000000000000000000000000000004600660006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550730300000000000000000000000000000000000006600760006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550730300000000000000000000000000000000000005600860006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550730300000000000000000000000000000000000008600960006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055503480156101b957600080fd5b5061156d806101c96000396000f3fe608060405234801561001057600080fd5b50600436106100a95760003560e01c806383b6033c1161007157806383b6033c14610150578063cd53c7471461016c578063df9849011461018a578063f16d968c146101a8578063f213a5a1146101c6578063fc784aa1146101e2576100a9565b806310fa1878146100ae57806312f758a0146100de578063385d763f146100fa5780633b1c29ef146101185780635fd7077214610134575b600080fd5b6100c860048036038101906100c391906107fe565b610200565b6040516100d59190610844565b60405180910390f35b6100f860048036038101906100f391906109b8565b610224565b005b6101026102cf565b60405161010f9190610a52565b60405180910390f35b610132600480360381019061012d9190610c3d565b6102d5565b005b61014e60048036038101906101499190610d6a565b610392565b005b61016a60048036038101906101659190610dc6565b61043a565b005b6101746104df565b6040516101819190610a52565b60405180910390f35b6101926104e5565b60405161019f9190610a52565b60405180910390f35b6101b06104eb565b6040516101bd9190610ecd565b60405180910390f35b6101e060048036038101906101db9190610eef565b610543565b005b6101ea610606565b6040516101f79190611087565b60405180910390f35b6002818154811061021057600080fd5b906000526020600020016000915090505481565b600760009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166320db7f0f8484846040518463ffffffff1660e01b815260040161028393929190611167565b602060405180830381865afa1580156102a0573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906102c491906111c1565b600181905550505050565b60045481565b600860009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166359c6945683836040518363ffffffff1660e01b81526004016103329291906111ee565b600060405180830381865afa15801561034f573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f820116820180604052508101906103789190611386565b6003908051906020019061038d9291906106a5565b505050565b600960009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663a967f50383836040518363ffffffff1660e01b81526004016103ef929190611457565b602060405180830381865afa15801561040c573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061043091906111c1565b6004819055505050565b600560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663e8787f41826040518263ffffffff1660e01b81526004016104959190611487565b602060405180830381865afa1580156104b2573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906104d691906111c1565b60008190555050565b60005481565b60015481565b6060600280548060200260200160405190810160405280929190818152602001828054801561053957602002820191906000526020600020905b815481526020019060010190808311610525575b5050505050905090565b600660009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663cca718bf858585856040518563ffffffff1660e01b81526004016105a494939291906114a9565b600060405180830381865afa1580156105c1573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f820116820180604052508101906105ea91906114ee565b600290805190602001906105ff929190610705565b5050505050565b60606003805480602002602001604051908101604052809291908181526020016000905b8282101561069c5783829060005260206000200180548060200260200160405190810160405280929190818152602001828054801561068857602002820191906000526020600020905b815481526020019060010190808311610674575b50505050508152602001906001019061062a565b50505050905090565b8280548282559060005260206000209081019282156106f4579160200282015b828111156106f35782518290805190602001906106e3929190610705565b50916020019190600101906106c5565b5b5090506107019190610752565b5090565b828054828255906000526020600020908101928215610741579160200282015b82811115610740578251825591602001919060010190610725565b5b50905061074e9190610776565b5090565b5b8082111561077257600081816107699190610793565b50600101610753565b5090565b5b8082111561078f576000816000905550600101610777565b5090565b50805460008255906000526020600020908101906107b19190610776565b50565b6000604051905090565b600080fd5b600080fd5b6000819050919050565b6107db816107c8565b81146107e657600080fd5b50565b6000813590506107f8816107d2565b92915050565b600060208284031215610814576108136107be565b5b6000610822848285016107e9565b91505092915050565b6000819050919050565b61083e8161082b565b82525050565b60006020820190506108596000830184610835565b92915050565b600080fd5b6000601f19601f8301169050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6108ad82610864565b810181811067ffffffffffffffff821117156108cc576108cb610875565b5b80604052505050565b60006108df6107b4565b90506108eb82826108a4565b919050565b600067ffffffffffffffff82111561090b5761090a610875565b5b602082029050602081019050919050565b600080fd5b600061093461092f846108f0565b6108d5565b905080838252602082019050602084028301858111156109575761095661091c565b5b835b81811015610980578061096c88826107e9565b845260208401935050602081019050610959565b5050509392505050565b600082601f83011261099f5761099e61085f565b5b81356109af848260208601610921565b91505092915050565b6000806000606084860312156109d1576109d06107be565b5b60006109df868287016107e9565b935050602084013567ffffffffffffffff811115610a00576109ff6107c3565b5b610a0c8682870161098a565b925050604084013567ffffffffffffffff811115610a2d57610a2c6107c3565b5b610a398682870161098a565b9150509250925092565b610a4c816107c8565b82525050565b6000602082019050610a676000830184610a43565b92915050565b600067ffffffffffffffff821115610a8857610a87610875565b5b602082029050602081019050919050565b600067ffffffffffffffff821115610ab457610ab3610875565b5b602082029050602081019050919050565b610ace8161082b565b8114610ad957600080fd5b50565b600081359050610aeb81610ac5565b92915050565b6000610b04610aff84610a99565b6108d5565b90508083825260208201905060208402830185811115610b2757610b2661091c565b5b835b81811015610b505780610b3c8882610adc565b845260208401935050602081019050610b29565b5050509392505050565b600082601f830112610b6f57610b6e61085f565b5b8135610b7f848260208601610af1565b91505092915050565b6000610b9b610b9684610a6d565b6108d5565b90508083825260208201905060208402830185811115610bbe57610bbd61091c565b5b835b81811015610c0557803567ffffffffffffffff811115610be357610be261085f565b5b808601610bf08982610b5a565b85526020850194505050602081019050610bc0565b5050509392505050565b600082601f830112610c2457610c2361085f565b5b8135610c34848260208601610b88565b91505092915050565b60008060408385031215610c5457610c536107be565b5b600083013567ffffffffffffffff811115610c7257610c716107c3565b5b610c7e85828601610c0f565b925050602083013567ffffffffffffffff811115610c9f57610c9e6107c3565b5b610cab85828601610c0f565b9150509250929050565b600080fd5b600067ffffffffffffffff821115610cd557610cd4610875565b5b610cde82610864565b9050602081019050919050565b82818337600083830152505050565b6000610d0d610d0884610cba565b6108d5565b905082815260208101848484011115610d2957610d28610cb5565b5b610d34848285610ceb565b509392505050565b600082601f830112610d5157610d5061085f565b5b8135610d61848260208601610cfa565b91505092915050565b60008060408385031215610d8157610d806107be565b5b600083013567ffffffffffffffff811115610d9f57610d9e6107c3565b5b610dab85828601610d3c565b9250506020610dbc858286016107e9565b9150509250929050565b600060208284031215610ddc57610ddb6107be565b5b600082013567ffffffffffffffff811115610dfa57610df96107c3565b5b610e068482850161098a565b91505092915050565b600081519050919050565b600082825260208201905092915050565b6000819050602082019050919050565b610e448161082b565b82525050565b6000610e568383610e3b565b60208301905092915050565b6000602082019050919050565b6000610e7a82610e0f565b610e848185610e1a565b9350610e8f83610e2b565b8060005b83811015610ec0578151610ea78882610e4a565b9750610eb283610e62565b925050600181019050610e93565b5085935050505092915050565b60006020820190508181036000830152610ee78184610e6f565b905092915050565b60008060008060808587031215610f0957610f086107be565b5b6000610f17878288016107e9565b9450506020610f2887828801610adc565b9350506040610f3987828801610adc565b9250506060610f4a878288016107e9565b91505092959194509250565b600081519050919050565b600082825260208201905092915050565b6000819050602082019050919050565b600082825260208201905092915050565b6000610f9e82610e0f565b610fa88185610f82565b9350610fb383610e2b565b8060005b83811015610fe4578151610fcb8882610e4a565b9750610fd683610e62565b925050600181019050610fb7565b5085935050505092915050565b6000610ffd8383610f93565b905092915050565b6000602082019050919050565b600061101d82610f56565b6110278185610f61565b93508360208202850161103985610f72565b8060005b8581101561107557848403895281516110568582610ff1565b945061106183611005565b925060208a0199505060018101905061103d565b50829750879550505050505092915050565b600060208201905081810360008301526110a18184611012565b905092915050565b600081519050919050565b600082825260208201905092915050565b6000819050602082019050919050565b6110de816107c8565b82525050565b60006110f083836110d5565b60208301905092915050565b6000602082019050919050565b6000611114826110a9565b61111e81856110b4565b9350611129836110c5565b8060005b8381101561115a57815161114188826110e4565b975061114c836110fc565b92505060018101905061112d565b5085935050505092915050565b600060608201905061117c6000830186610a43565b818103602083015261118e8185611109565b905081810360408301526111a28184611109565b9050949350505050565b6000815190506111bb816107d2565b92915050565b6000602082840312156111d7576111d66107be565b5b60006111e5848285016111ac565b91505092915050565b600060408201905081810360008301526112088185611012565b9050818103602083015261121c8184611012565b90509392505050565b60008151905061123481610ac5565b92915050565b600061124d61124884610a99565b6108d5565b905080838252602082019050602084028301858111156112705761126f61091c565b5b835b8181101561129957806112858882611225565b845260208401935050602081019050611272565b5050509392505050565b600082601f8301126112b8576112b761085f565b5b81516112c884826020860161123a565b91505092915050565b60006112e46112df84610a6d565b6108d5565b905080838252602082019050602084028301858111156113075761130661091c565b5b835b8181101561134e57805167ffffffffffffffff81111561132c5761132b61085f565b5b80860161133989826112a3565b85526020850194505050602081019050611309565b5050509392505050565b600082601f83011261136d5761136c61085f565b5b815161137d8482602086016112d1565b91505092915050565b60006020828403121561139c5761139b6107be565b5b600082015167ffffffffffffffff8111156113ba576113b96107c3565b5b6113c684828501611358565b91505092915050565b600081519050919050565b600082825260208201905092915050565b60005b838110156114095780820151818401526020810190506113ee565b83811115611418576000848401525b50505050565b6000611429826113cf565b61143381856113da565b93506114438185602086016113eb565b61144c81610864565b840191505092915050565b60006040820190508181036000830152611471818561141e565b90506114806020830184610a43565b9392505050565b600060208201905081810360008301526114a18184611109565b905092915050565b60006080820190506114be6000830187610a43565b6114cb6020830186610835565b6114d86040830185610835565b6114e56060830184610a43565b95945050505050565b600060208284031215611504576115036107be565b5b600082015167ffffffffffffffff811115611522576115216107c3565b5b61152e848285016112a3565b9150509291505056fea264697066735822122071f5d38f55f569171272206a412a641b67a14cee1ea97b212e7934db831506c664736f6c634300080f0033",
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

