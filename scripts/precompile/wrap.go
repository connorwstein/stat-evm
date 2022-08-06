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
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"message\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"res\",\"type\":\"int256\"}],\"name\":\"Debug\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"fitRes\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getLastSample\",\"outputs\":[{\"internalType\":\"int256[]\",\"name\":\"\",\"type\":\"int256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMatrixMulti\",\"outputs\":[{\"internalType\":\"int256[][]\",\"name\":\"\",\"type\":\"int256[][]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"med\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"sample\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"vals\",\"type\":\"uint256[]\"}],\"name\":\"testFit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int256[][]\",\"name\":\"a\",\"type\":\"int256[][]\"},{\"internalType\":\"int256[][]\",\"name\":\"b\",\"type\":\"int256[][]\"}],\"name\":\"testMatrixMult\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"vals\",\"type\":\"uint256[]\"}],\"name\":\"testMedian\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"distributionType\",\"type\":\"uint256\"},{\"internalType\":\"int256\",\"name\":\"a\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"b\",\"type\":\"int256\"},{\"internalType\":\"uint256\",\"name\":\"numSamples\",\"type\":\"uint256\"}],\"name\":\"testSampler\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x6080604052730300000000000000000000000000000000000001600460006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550730300000000000000000000000000000000000004600560006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550730300000000000000000000000000000000000005600660006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550730300000000000000000000000000000000000006600760006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555034801561016457600080fd5b506111d0806101746000396000f3fe608060405234801561001057600080fd5b50600436106100935760003560e01c8063bdee25fb11610066578063bdee25fb1461011c578063cd53c7471461013a578063f16d968c14610158578063f213a5a114610176578063fc784aa11461019257610093565b806310fa1878146100985780633b1c29ef146100c857806355817639146100e457806383b6033c14610100575b600080fd5b6100b260048036038101906100ad91906106fa565b6101b0565b6040516100bf9190610740565b60405180910390f35b6100e260048036038101906100dd91906109c1565b6101d4565b005b6100fe60048036038101906100f99190610afc565b610291565b005b61011a60048036038101906101159190610afc565b610336565b005b6101246103db565b6040516101319190610b54565b60405180910390f35b6101426103e1565b60405161014f9190610b54565b60405180910390f35b6101606103e7565b60405161016d9190610c2d565b60405180910390f35b610190600480360381019061018b9190610c4f565b61043f565b005b61019a610502565b6040516101a79190610de7565b60405180910390f35b600181815481106101c057600080fd5b906000526020600020016000915090505481565b600660009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166359c6945683836040518363ffffffff1660e01b8152600401610231929190610e09565b600060405180830381865afa15801561024e573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f820116820180604052508101906102779190610fa1565b6002908051906020019061028c9291906105a1565b505050565b600760009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166392091876826040518263ffffffff1660e01b81526004016102ec91906110a8565b602060405180830381865afa158015610309573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061032d91906110df565b60038190555050565b600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663e8787f41826040518263ffffffff1660e01b815260040161039191906110a8565b602060405180830381865afa1580156103ae573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906103d291906110df565b60008190555050565b60035481565b60005481565b6060600180548060200260200160405190810160405280929190818152602001828054801561043557602002820191906000526020600020905b815481526020019060010190808311610421575b5050505050905090565b600560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663cca718bf858585856040518563ffffffff1660e01b81526004016104a0949392919061110c565b600060405180830381865afa1580156104bd573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f820116820180604052508101906104e69190611151565b600190805190602001906104fb929190610601565b5050505050565b60606002805480602002602001604051908101604052809291908181526020016000905b828210156105985783829060005260206000200180548060200260200160405190810160405280929190818152602001828054801561058457602002820191906000526020600020905b815481526020019060010190808311610570575b505050505081526020019060010190610526565b50505050905090565b8280548282559060005260206000209081019282156105f0579160200282015b828111156105ef5782518290805190602001906105df929190610601565b50916020019190600101906105c1565b5b5090506105fd919061064e565b5090565b82805482825590600052602060002090810192821561063d579160200282015b8281111561063c578251825591602001919060010190610621565b5b50905061064a9190610672565b5090565b5b8082111561066e5760008181610665919061068f565b5060010161064f565b5090565b5b8082111561068b576000816000905550600101610673565b5090565b50805460008255906000526020600020908101906106ad9190610672565b50565b6000604051905090565b600080fd5b600080fd5b6000819050919050565b6106d7816106c4565b81146106e257600080fd5b50565b6000813590506106f4816106ce565b92915050565b6000602082840312156107105761070f6106ba565b5b600061071e848285016106e5565b91505092915050565b6000819050919050565b61073a81610727565b82525050565b60006020820190506107556000830184610731565b92915050565b600080fd5b6000601f19601f8301169050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6107a982610760565b810181811067ffffffffffffffff821117156107c8576107c7610771565b5b80604052505050565b60006107db6106b0565b90506107e782826107a0565b919050565b600067ffffffffffffffff82111561080757610806610771565b5b602082029050602081019050919050565b600080fd5b600067ffffffffffffffff82111561083857610837610771565b5b602082029050602081019050919050565b61085281610727565b811461085d57600080fd5b50565b60008135905061086f81610849565b92915050565b60006108886108838461081d565b6107d1565b905080838252602082019050602084028301858111156108ab576108aa610818565b5b835b818110156108d457806108c08882610860565b8452602084019350506020810190506108ad565b5050509392505050565b600082601f8301126108f3576108f261075b565b5b8135610903848260208601610875565b91505092915050565b600061091f61091a846107ec565b6107d1565b9050808382526020820190506020840283018581111561094257610941610818565b5b835b8181101561098957803567ffffffffffffffff8111156109675761096661075b565b5b80860161097489826108de565b85526020850194505050602081019050610944565b5050509392505050565b600082601f8301126109a8576109a761075b565b5b81356109b884826020860161090c565b91505092915050565b600080604083850312156109d8576109d76106ba565b5b600083013567ffffffffffffffff8111156109f6576109f56106bf565b5b610a0285828601610993565b925050602083013567ffffffffffffffff811115610a2357610a226106bf565b5b610a2f85828601610993565b9150509250929050565b600067ffffffffffffffff821115610a5457610a53610771565b5b602082029050602081019050919050565b6000610a78610a7384610a39565b6107d1565b90508083825260208201905060208402830185811115610a9b57610a9a610818565b5b835b81811015610ac45780610ab088826106e5565b845260208401935050602081019050610a9d565b5050509392505050565b600082601f830112610ae357610ae261075b565b5b8135610af3848260208601610a65565b91505092915050565b600060208284031215610b1257610b116106ba565b5b600082013567ffffffffffffffff811115610b3057610b2f6106bf565b5b610b3c84828501610ace565b91505092915050565b610b4e816106c4565b82525050565b6000602082019050610b696000830184610b45565b92915050565b600081519050919050565b600082825260208201905092915050565b6000819050602082019050919050565b610ba481610727565b82525050565b6000610bb68383610b9b565b60208301905092915050565b6000602082019050919050565b6000610bda82610b6f565b610be48185610b7a565b9350610bef83610b8b565b8060005b83811015610c20578151610c078882610baa565b9750610c1283610bc2565b925050600181019050610bf3565b5085935050505092915050565b60006020820190508181036000830152610c478184610bcf565b905092915050565b60008060008060808587031215610c6957610c686106ba565b5b6000610c77878288016106e5565b9450506020610c8887828801610860565b9350506040610c9987828801610860565b9250506060610caa878288016106e5565b91505092959194509250565b600081519050919050565b600082825260208201905092915050565b6000819050602082019050919050565b600082825260208201905092915050565b6000610cfe82610b6f565b610d088185610ce2565b9350610d1383610b8b565b8060005b83811015610d44578151610d2b8882610baa565b9750610d3683610bc2565b925050600181019050610d17565b5085935050505092915050565b6000610d5d8383610cf3565b905092915050565b6000602082019050919050565b6000610d7d82610cb6565b610d878185610cc1565b935083602082028501610d9985610cd2565b8060005b85811015610dd55784840389528151610db68582610d51565b9450610dc183610d65565b925060208a01995050600181019050610d9d565b50829750879550505050505092915050565b60006020820190508181036000830152610e018184610d72565b905092915050565b60006040820190508181036000830152610e238185610d72565b90508181036020830152610e378184610d72565b90509392505050565b600081519050610e4f81610849565b92915050565b6000610e68610e638461081d565b6107d1565b90508083825260208201905060208402830185811115610e8b57610e8a610818565b5b835b81811015610eb45780610ea08882610e40565b845260208401935050602081019050610e8d565b5050509392505050565b600082601f830112610ed357610ed261075b565b5b8151610ee3848260208601610e55565b91505092915050565b6000610eff610efa846107ec565b6107d1565b90508083825260208201905060208402830185811115610f2257610f21610818565b5b835b81811015610f6957805167ffffffffffffffff811115610f4757610f4661075b565b5b808601610f548982610ebe565b85526020850194505050602081019050610f24565b5050509392505050565b600082601f830112610f8857610f8761075b565b5b8151610f98848260208601610eec565b91505092915050565b600060208284031215610fb757610fb66106ba565b5b600082015167ffffffffffffffff811115610fd557610fd46106bf565b5b610fe184828501610f73565b91505092915050565b600081519050919050565b600082825260208201905092915050565b6000819050602082019050919050565b61101f816106c4565b82525050565b60006110318383611016565b60208301905092915050565b6000602082019050919050565b600061105582610fea565b61105f8185610ff5565b935061106a83611006565b8060005b8381101561109b5781516110828882611025565b975061108d8361103d565b92505060018101905061106e565b5085935050505092915050565b600060208201905081810360008301526110c2818461104a565b905092915050565b6000815190506110d9816106ce565b92915050565b6000602082840312156110f5576110f46106ba565b5b6000611103848285016110ca565b91505092915050565b60006080820190506111216000830187610b45565b61112e6020830186610731565b61113b6040830185610731565b6111486060830184610b45565b95945050505050565b600060208284031215611167576111666106ba565b5b600082015167ffffffffffffffff811115611185576111846106bf565b5b61119184828501610ebe565b9150509291505056fea26469706673582212207da31cb902e20c7d6a7a0ec5fce19096d545b0ddedf2c4abb076f8d0485a283964736f6c634300080f0033",
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

// FitRes is a free data retrieval call binding the contract method 0xbdee25fb.
//
// Solidity: function fitRes() view returns(uint256)
func (_Main *MainCaller) FitRes(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "fitRes")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FitRes is a free data retrieval call binding the contract method 0xbdee25fb.
//
// Solidity: function fitRes() view returns(uint256)
func (_Main *MainSession) FitRes() (*big.Int, error) {
	return _Main.Contract.FitRes(&_Main.CallOpts)
}

// FitRes is a free data retrieval call binding the contract method 0xbdee25fb.
//
// Solidity: function fitRes() view returns(uint256)
func (_Main *MainCallerSession) FitRes() (*big.Int, error) {
	return _Main.Contract.FitRes(&_Main.CallOpts)
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

// TestFit is a paid mutator transaction binding the contract method 0x55817639.
//
// Solidity: function testFit(uint256[] vals) returns()
func (_Main *MainTransactor) TestFit(opts *bind.TransactOpts, vals []*big.Int) (*types.Transaction, error) {
	return _Main.contract.Transact(opts, "testFit", vals)
}

// TestFit is a paid mutator transaction binding the contract method 0x55817639.
//
// Solidity: function testFit(uint256[] vals) returns()
func (_Main *MainSession) TestFit(vals []*big.Int) (*types.Transaction, error) {
	return _Main.Contract.TestFit(&_Main.TransactOpts, vals)
}

// TestFit is a paid mutator transaction binding the contract method 0x55817639.
//
// Solidity: function testFit(uint256[] vals) returns()
func (_Main *MainTransactorSession) TestFit(vals []*big.Int) (*types.Transaction, error) {
	return _Main.Contract.TestFit(&_Main.TransactOpts, vals)
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

