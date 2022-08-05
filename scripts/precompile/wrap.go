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
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"message\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"res\",\"type\":\"uint256\"}],\"name\":\"Debug\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"getMatrixMulti\",\"outputs\":[{\"internalType\":\"int256[][]\",\"name\":\"\",\"type\":\"int256[][]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"last\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sample\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int256[][]\",\"name\":\"a\",\"type\":\"int256[][]\"},{\"internalType\":\"int256[][]\",\"name\":\"b\",\"type\":\"int256[][]\"}],\"name\":\"testMatrixMult\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"v1\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"v2\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"v3\",\"type\":\"uint256\"}],\"name\":\"testMedian\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"v1\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"v2\",\"type\":\"uint256\"}],\"name\":\"testSampler\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60806040526000805560018055730300000000000000000000000000000000000001600360006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550730300000000000000000000000000000000000004600460006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550730300000000000000000000000000000000000005600560006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555034801561011757600080fd5b50610ef5806101276000396000f3fe608060405234801561001057600080fd5b50600436106100625760003560e01c80633b1c29ef1461006757806347799da81461008357806359646bee146100a1578063af5e8e3d146100bd578063c9482a5d146100d9578063fc784aa1146100f7575b600080fd5b610081600480360381019061007c91906107d3565b610115565b005b61008b6101d2565b6040516100989190610864565b60405180910390f35b6100bb60048036038101906100b691906108ab565b6101d8565b005b6100d760048036038101906100d291906108fe565b6102f3565b005b6100e161039b565b6040516100ee919061094d565b60405180910390f35b6100ff6103a1565b60405161010c9190610ae8565b60405180910390f35b600560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166359c6945683836040518363ffffffff1660e01b8152600401610172929190610b0a565b600060405180830381865afa15801561018f573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f820116820180604052508101906101b89190610ca2565b600290805190602001906101cd929190610440565b505050565b60005481565b7f3c5ad147104e56be34a9176a6692f7df8d2f4b29a5af06bc6b98970d329d6577836040516102079190610d48565b60405180910390a1600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663d96539a08484846040518463ffffffff1660e01b815260040161026e93929190610d76565b602060405180830381865afa15801561028b573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906102af9190610dc2565b6000819055507f3c5ad147104e56be34a9176a6692f7df8d2f4b29a5af06bc6b98970d329d65776000546040516102e69190610e3b565b60405180910390a1505050565b600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663f1c8716283836040518363ffffffff1660e01b8152600401610350929190610e69565b602060405180830381865afa15801561036d573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906103919190610e92565b6001819055505050565b60015481565b60606002805480602002602001604051908101604052809291908181526020016000905b828210156104375783829060005260206000200180548060200260200160405190810160405280929190818152602001828054801561042357602002820191906000526020600020905b81548152602001906001019080831161040f575b5050505050815260200190600101906103c5565b50505050905090565b82805482825590600052602060002090810192821561048f579160200282015b8281111561048e57825182908051906020019061047e9291906104a0565b5091602001919060010190610460565b5b50905061049c91906104ed565b5090565b8280548282559060005260206000209081019282156104dc579160200282015b828111156104db5782518255916020019190600101906104c0565b5b5090506104e99190610511565b5090565b5b8082111561050d5760008181610504919061052e565b506001016104ee565b5090565b5b8082111561052a576000816000905550600101610512565b5090565b508054600082559060005260206000209081019061054c9190610511565b50565b6000604051905090565b600080fd5b600080fd5b600080fd5b6000601f19601f8301169050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6105b182610568565b810181811067ffffffffffffffff821117156105d0576105cf610579565b5b80604052505050565b60006105e361054f565b90506105ef82826105a8565b919050565b600067ffffffffffffffff82111561060f5761060e610579565b5b602082029050602081019050919050565b600080fd5b600067ffffffffffffffff8211156106405761063f610579565b5b602082029050602081019050919050565b6000819050919050565b61066481610651565b811461066f57600080fd5b50565b6000813590506106818161065b565b92915050565b600061069a61069584610625565b6105d9565b905080838252602082019050602084028301858111156106bd576106bc610620565b5b835b818110156106e657806106d28882610672565b8452602084019350506020810190506106bf565b5050509392505050565b600082601f83011261070557610704610563565b5b8135610715848260208601610687565b91505092915050565b600061073161072c846105f4565b6105d9565b9050808382526020820190506020840283018581111561075457610753610620565b5b835b8181101561079b57803567ffffffffffffffff81111561077957610778610563565b5b80860161078689826106f0565b85526020850194505050602081019050610756565b5050509392505050565b600082601f8301126107ba576107b9610563565b5b81356107ca84826020860161071e565b91505092915050565b600080604083850312156107ea576107e9610559565b5b600083013567ffffffffffffffff8111156108085761080761055e565b5b610814858286016107a5565b925050602083013567ffffffffffffffff8111156108355761083461055e565b5b610841858286016107a5565b9150509250929050565b6000819050919050565b61085e8161084b565b82525050565b60006020820190506108796000830184610855565b92915050565b6108888161084b565b811461089357600080fd5b50565b6000813590506108a58161087f565b92915050565b6000806000606084860312156108c4576108c3610559565b5b60006108d286828701610896565b93505060206108e386828701610896565b92505060406108f486828701610896565b9150509250925092565b6000806040838503121561091557610914610559565b5b600061092385828601610896565b925050602061093485828601610896565b9150509250929050565b61094781610651565b82525050565b6000602082019050610962600083018461093e565b92915050565b600081519050919050565b600082825260208201905092915050565b6000819050602082019050919050565b600081519050919050565b600082825260208201905092915050565b6000819050602082019050919050565b6109c981610651565b82525050565b60006109db83836109c0565b60208301905092915050565b6000602082019050919050565b60006109ff82610994565b610a09818561099f565b9350610a14836109b0565b8060005b83811015610a45578151610a2c88826109cf565b9750610a37836109e7565b925050600181019050610a18565b5085935050505092915050565b6000610a5e83836109f4565b905092915050565b6000602082019050919050565b6000610a7e82610968565b610a888185610973565b935083602082028501610a9a85610984565b8060005b85811015610ad65784840389528151610ab78582610a52565b9450610ac283610a66565b925060208a01995050600181019050610a9e565b50829750879550505050505092915050565b60006020820190508181036000830152610b028184610a73565b905092915050565b60006040820190508181036000830152610b248185610a73565b90508181036020830152610b388184610a73565b90509392505050565b600081519050610b508161065b565b92915050565b6000610b69610b6484610625565b6105d9565b90508083825260208201905060208402830185811115610b8c57610b8b610620565b5b835b81811015610bb55780610ba18882610b41565b845260208401935050602081019050610b8e565b5050509392505050565b600082601f830112610bd457610bd3610563565b5b8151610be4848260208601610b56565b91505092915050565b6000610c00610bfb846105f4565b6105d9565b90508083825260208201905060208402830185811115610c2357610c22610620565b5b835b81811015610c6a57805167ffffffffffffffff811115610c4857610c47610563565b5b808601610c558982610bbf565b85526020850194505050602081019050610c25565b5050509392505050565b600082601f830112610c8957610c88610563565b5b8151610c99848260208601610bed565b91505092915050565b600060208284031215610cb857610cb7610559565b5b600082015167ffffffffffffffff811115610cd657610cd561055e565b5b610ce284828501610c74565b91505092915050565b600082825260208201905092915050565b7f5061727431000000000000000000000000000000000000000000000000000000600082015250565b6000610d32600583610ceb565b9150610d3d82610cfc565b602082019050919050565b60006040820190508181036000830152610d6181610d25565b9050610d706020830184610855565b92915050565b6000606082019050610d8b6000830186610855565b610d986020830185610855565b610da56040830184610855565b949350505050565b600081519050610dbc8161087f565b92915050565b600060208284031215610dd857610dd7610559565b5b6000610de684828501610dad565b91505092915050565b7f5061727432000000000000000000000000000000000000000000000000000000600082015250565b6000610e25600583610ceb565b9150610e3082610def565b602082019050919050565b60006040820190508181036000830152610e5481610e18565b9050610e636020830184610855565b92915050565b6000604082019050610e7e6000830185610855565b610e8b6020830184610855565b9392505050565b600060208284031215610ea857610ea7610559565b5b6000610eb684828501610b41565b9150509291505056fea2646970667358221220585385b1a575d420d2a3f5f1fcbfc0c72c185e18fd3378fc8e315e26567533b164736f6c634300080f0033",
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

// Last is a free data retrieval call binding the contract method 0x47799da8.
//
// Solidity: function last() view returns(uint256)
func (_Main *MainCaller) Last(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "last")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Last is a free data retrieval call binding the contract method 0x47799da8.
//
// Solidity: function last() view returns(uint256)
func (_Main *MainSession) Last() (*big.Int, error) {
	return _Main.Contract.Last(&_Main.CallOpts)
}

// Last is a free data retrieval call binding the contract method 0x47799da8.
//
// Solidity: function last() view returns(uint256)
func (_Main *MainCallerSession) Last() (*big.Int, error) {
	return _Main.Contract.Last(&_Main.CallOpts)
}

// Sample is a free data retrieval call binding the contract method 0xc9482a5d.
//
// Solidity: function sample() view returns(int256)
func (_Main *MainCaller) Sample(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "sample")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Sample is a free data retrieval call binding the contract method 0xc9482a5d.
//
// Solidity: function sample() view returns(int256)
func (_Main *MainSession) Sample() (*big.Int, error) {
	return _Main.Contract.Sample(&_Main.CallOpts)
}

// Sample is a free data retrieval call binding the contract method 0xc9482a5d.
//
// Solidity: function sample() view returns(int256)
func (_Main *MainCallerSession) Sample() (*big.Int, error) {
	return _Main.Contract.Sample(&_Main.CallOpts)
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

// TestMedian is a paid mutator transaction binding the contract method 0x59646bee.
//
// Solidity: function testMedian(uint256 v1, uint256 v2, uint256 v3) returns()
func (_Main *MainTransactor) TestMedian(opts *bind.TransactOpts, v1 *big.Int, v2 *big.Int, v3 *big.Int) (*types.Transaction, error) {
	return _Main.contract.Transact(opts, "testMedian", v1, v2, v3)
}

// TestMedian is a paid mutator transaction binding the contract method 0x59646bee.
//
// Solidity: function testMedian(uint256 v1, uint256 v2, uint256 v3) returns()
func (_Main *MainSession) TestMedian(v1 *big.Int, v2 *big.Int, v3 *big.Int) (*types.Transaction, error) {
	return _Main.Contract.TestMedian(&_Main.TransactOpts, v1, v2, v3)
}

// TestMedian is a paid mutator transaction binding the contract method 0x59646bee.
//
// Solidity: function testMedian(uint256 v1, uint256 v2, uint256 v3) returns()
func (_Main *MainTransactorSession) TestMedian(v1 *big.Int, v2 *big.Int, v3 *big.Int) (*types.Transaction, error) {
	return _Main.Contract.TestMedian(&_Main.TransactOpts, v1, v2, v3)
}

// TestSampler is a paid mutator transaction binding the contract method 0xaf5e8e3d.
//
// Solidity: function testSampler(uint256 v1, uint256 v2) returns()
func (_Main *MainTransactor) TestSampler(opts *bind.TransactOpts, v1 *big.Int, v2 *big.Int) (*types.Transaction, error) {
	return _Main.contract.Transact(opts, "testSampler", v1, v2)
}

// TestSampler is a paid mutator transaction binding the contract method 0xaf5e8e3d.
//
// Solidity: function testSampler(uint256 v1, uint256 v2) returns()
func (_Main *MainSession) TestSampler(v1 *big.Int, v2 *big.Int) (*types.Transaction, error) {
	return _Main.Contract.TestSampler(&_Main.TransactOpts, v1, v2)
}

// TestSampler is a paid mutator transaction binding the contract method 0xaf5e8e3d.
//
// Solidity: function testSampler(uint256 v1, uint256 v2) returns()
func (_Main *MainTransactorSession) TestSampler(v1 *big.Int, v2 *big.Int) (*types.Transaction, error) {
	return _Main.Contract.TestSampler(&_Main.TransactOpts, v1, v2)
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

// FilterDebug is a free log retrieval operation binding the contract event 0x3c5ad147104e56be34a9176a6692f7df8d2f4b29a5af06bc6b98970d329d6577.
//
// Solidity: event Debug(string message, uint256 res)
func (_Main *MainFilterer) FilterDebug(opts *bind.FilterOpts) (*MainDebugIterator, error) {

	logs, sub, err := _Main.contract.FilterLogs(opts, "Debug")
	if err != nil {
		return nil, err
	}
	return &MainDebugIterator{contract: _Main.contract, event: "Debug", logs: logs, sub: sub}, nil
}

// WatchDebug is a free log subscription operation binding the contract event 0x3c5ad147104e56be34a9176a6692f7df8d2f4b29a5af06bc6b98970d329d6577.
//
// Solidity: event Debug(string message, uint256 res)
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

// ParseDebug is a log parse operation binding the contract event 0x3c5ad147104e56be34a9176a6692f7df8d2f4b29a5af06bc6b98970d329d6577.
//
// Solidity: event Debug(string message, uint256 res)
func (_Main *MainFilterer) ParseDebug(log types.Log) (*MainDebug, error) {
	event := new(MainDebug)
	if err := _Main.contract.UnpackLog(event, "Debug", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

