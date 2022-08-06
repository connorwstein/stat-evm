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

// MomentMetaData contains all meta data concerning the Moment contract.
var MomentMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"x\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"y\",\"type\":\"uint256[]\"},{\"internalType\":\"uint64\",\"name\":\"momentIndex\",\"type\":\"uint64\"}],\"name\":\"getMoment\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"947d649a": "getMoment(uint256[],uint256[],uint64)",
	},
}

// MomentABI is the input ABI used to generate the binding from.
// Deprecated: Use MomentMetaData.ABI instead.
var MomentABI = MomentMetaData.ABI

// Deprecated: Use MomentMetaData.Sigs instead.
// MomentFuncSigs maps the 4-byte function signature to its string representation.
var MomentFuncSigs = MomentMetaData.Sigs

// Moment is an auto generated Go binding around an Ethereum contract.
type Moment struct {
	MomentCaller     // Read-only binding to the contract
	MomentTransactor // Write-only binding to the contract
	MomentFilterer   // Log filterer for contract events
}

// MomentCaller is an auto generated read-only Go binding around an Ethereum contract.
type MomentCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MomentTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MomentTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MomentFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MomentFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MomentSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MomentSession struct {
	Contract     *Moment           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MomentCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MomentCallerSession struct {
	Contract *MomentCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// MomentTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MomentTransactorSession struct {
	Contract     *MomentTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MomentRaw is an auto generated low-level Go binding around an Ethereum contract.
type MomentRaw struct {
	Contract *Moment // Generic contract binding to access the raw methods on
}

// MomentCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MomentCallerRaw struct {
	Contract *MomentCaller // Generic read-only contract binding to access the raw methods on
}

// MomentTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MomentTransactorRaw struct {
	Contract *MomentTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMoment creates a new instance of Moment, bound to a specific deployed contract.
func NewMoment(address common.Address, backend bind.ContractBackend) (*Moment, error) {
	contract, err := bindMoment(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Moment{MomentCaller: MomentCaller{contract: contract}, MomentTransactor: MomentTransactor{contract: contract}, MomentFilterer: MomentFilterer{contract: contract}}, nil
}

// NewMomentCaller creates a new read-only instance of Moment, bound to a specific deployed contract.
func NewMomentCaller(address common.Address, caller bind.ContractCaller) (*MomentCaller, error) {
	contract, err := bindMoment(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MomentCaller{contract: contract}, nil
}

// NewMomentTransactor creates a new write-only instance of Moment, bound to a specific deployed contract.
func NewMomentTransactor(address common.Address, transactor bind.ContractTransactor) (*MomentTransactor, error) {
	contract, err := bindMoment(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MomentTransactor{contract: contract}, nil
}

// NewMomentFilterer creates a new log filterer instance of Moment, bound to a specific deployed contract.
func NewMomentFilterer(address common.Address, filterer bind.ContractFilterer) (*MomentFilterer, error) {
	contract, err := bindMoment(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MomentFilterer{contract: contract}, nil
}

// bindMoment binds a generic wrapper to an already deployed contract.
func bindMoment(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MomentABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Moment *MomentRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Moment.Contract.MomentCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Moment *MomentRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Moment.Contract.MomentTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Moment *MomentRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Moment.Contract.MomentTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Moment *MomentCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Moment.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Moment *MomentTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Moment.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Moment *MomentTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Moment.Contract.contract.Transact(opts, method, params...)
}

// GetMoment is a free data retrieval call binding the contract method 0x947d649a.
//
// Solidity: function getMoment(uint256[] x, uint256[] y, uint64 momentIndex) view returns(uint256)
func (_Moment *MomentCaller) GetMoment(opts *bind.CallOpts, x []*big.Int, y []*big.Int, momentIndex uint64) (*big.Int, error) {
	var out []interface{}
	err := _Moment.contract.Call(opts, &out, "getMoment", x, y, momentIndex)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMoment is a free data retrieval call binding the contract method 0x947d649a.
//
// Solidity: function getMoment(uint256[] x, uint256[] y, uint64 momentIndex) view returns(uint256)
func (_Moment *MomentSession) GetMoment(x []*big.Int, y []*big.Int, momentIndex uint64) (*big.Int, error) {
	return _Moment.Contract.GetMoment(&_Moment.CallOpts, x, y, momentIndex)
}

// GetMoment is a free data retrieval call binding the contract method 0x947d649a.
//
// Solidity: function getMoment(uint256[] x, uint256[] y, uint64 momentIndex) view returns(uint256)
func (_Moment *MomentCallerSession) GetMoment(x []*big.Int, y []*big.Int, momentIndex uint64) (*big.Int, error) {
	return _Moment.Contract.GetMoment(&_Moment.CallOpts, x, y, momentIndex)
}

// TestMetaData contains all meta data concerning the Test contract.
var TestMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"moment\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"x\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"y\",\"type\":\"uint256[]\"},{\"internalType\":\"uint64\",\"name\":\"momentIndex\",\"type\":\"uint64\"}],\"name\":\"testMoment\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"df984901": "moment()",
		"e9d6b2bc": "testMoment(uint256[],uint256[],uint64)",
	},
	Bin: "0x6080604052600180546001600160a01b03191673030000000000000000000000000000000000000217905534801561003657600080fd5b506102e1806100466000396000f3fe608060405234801561001057600080fd5b50600436106100365760003560e01c8063df9849011461003b578063e9d6b2bc14610056575b600080fd5b61004460005481565b60405190815260200160405180910390f35b610069610064366004610199565b61006b565b005b600154604051634a3eb24d60e11b81526001600160a01b039091169063947d649a9061009f90869086908690600401610252565b602060405180830381865afa1580156100bc573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906100e09190610292565b600055505050565b634e487b7160e01b600052604160045260246000fd5b600082601f83011261010f57600080fd5b8135602067ffffffffffffffff8083111561012c5761012c6100e8565b8260051b604051601f19603f83011681018181108482111715610151576101516100e8565b60405293845285810183019383810192508785111561016f57600080fd5b83870191505b8482101561018e57813583529183019190830190610175565b979650505050505050565b6000806000606084860312156101ae57600080fd5b833567ffffffffffffffff808211156101c657600080fd5b6101d2878388016100fe565b945060208601359150808211156101e857600080fd5b6101f4878388016100fe565b935060408601359150808216821461020b57600080fd5b50809150509250925092565b600081518084526020808501945080840160005b838110156102475781518752958201959082019060010161022b565b509495945050505050565b6060815260006102656060830186610217565b82810360208401526102778186610217565b91505067ffffffffffffffff83166040830152949350505050565b6000602082840312156102a457600080fd5b505191905056fea26469706673582212200e74bed091980a33bb6a8fb68cb43439abd11b17d4a6282dad86f513c4880bb164736f6c634300080f0033",
}

// TestABI is the input ABI used to generate the binding from.
// Deprecated: Use TestMetaData.ABI instead.
var TestABI = TestMetaData.ABI

// Deprecated: Use TestMetaData.Sigs instead.
// TestFuncSigs maps the 4-byte function signature to its string representation.
var TestFuncSigs = TestMetaData.Sigs

// TestBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use TestMetaData.Bin instead.
var TestBin = TestMetaData.Bin

// DeployTest deploys a new Ethereum contract, binding an instance of Test to it.
func DeployTest(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Test, error) {
	parsed, err := TestMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(TestBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Test{TestCaller: TestCaller{contract: contract}, TestTransactor: TestTransactor{contract: contract}, TestFilterer: TestFilterer{contract: contract}}, nil
}

// Test is an auto generated Go binding around an Ethereum contract.
type Test struct {
	TestCaller     // Read-only binding to the contract
	TestTransactor // Write-only binding to the contract
	TestFilterer   // Log filterer for contract events
}

// TestCaller is an auto generated read-only Go binding around an Ethereum contract.
type TestCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TestTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TestFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TestSession struct {
	Contract     *Test             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TestCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TestCallerSession struct {
	Contract *TestCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// TestTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TestTransactorSession struct {
	Contract     *TestTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TestRaw is an auto generated low-level Go binding around an Ethereum contract.
type TestRaw struct {
	Contract *Test // Generic contract binding to access the raw methods on
}

// TestCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TestCallerRaw struct {
	Contract *TestCaller // Generic read-only contract binding to access the raw methods on
}

// TestTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TestTransactorRaw struct {
	Contract *TestTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTest creates a new instance of Test, bound to a specific deployed contract.
func NewTest(address common.Address, backend bind.ContractBackend) (*Test, error) {
	contract, err := bindTest(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Test{TestCaller: TestCaller{contract: contract}, TestTransactor: TestTransactor{contract: contract}, TestFilterer: TestFilterer{contract: contract}}, nil
}

// NewTestCaller creates a new read-only instance of Test, bound to a specific deployed contract.
func NewTestCaller(address common.Address, caller bind.ContractCaller) (*TestCaller, error) {
	contract, err := bindTest(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TestCaller{contract: contract}, nil
}

// NewTestTransactor creates a new write-only instance of Test, bound to a specific deployed contract.
func NewTestTransactor(address common.Address, transactor bind.ContractTransactor) (*TestTransactor, error) {
	contract, err := bindTest(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TestTransactor{contract: contract}, nil
}

// NewTestFilterer creates a new log filterer instance of Test, bound to a specific deployed contract.
func NewTestFilterer(address common.Address, filterer bind.ContractFilterer) (*TestFilterer, error) {
	contract, err := bindTest(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TestFilterer{contract: contract}, nil
}

// bindTest binds a generic wrapper to an already deployed contract.
func bindTest(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TestABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Test *TestRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Test.Contract.TestCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Test *TestRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Test.Contract.TestTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Test *TestRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Test.Contract.TestTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Test *TestCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Test.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Test *TestTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Test.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Test *TestTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Test.Contract.contract.Transact(opts, method, params...)
}

// Moment is a free data retrieval call binding the contract method 0xdf984901.
//
// Solidity: function moment() view returns(uint256)
func (_Test *TestCaller) Moment(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Test.contract.Call(opts, &out, "moment")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Moment is a free data retrieval call binding the contract method 0xdf984901.
//
// Solidity: function moment() view returns(uint256)
func (_Test *TestSession) Moment() (*big.Int, error) {
	return _Test.Contract.Moment(&_Test.CallOpts)
}

// Moment is a free data retrieval call binding the contract method 0xdf984901.
//
// Solidity: function moment() view returns(uint256)
func (_Test *TestCallerSession) Moment() (*big.Int, error) {
	return _Test.Contract.Moment(&_Test.CallOpts)
}

// TestMoment is a paid mutator transaction binding the contract method 0xe9d6b2bc.
//
// Solidity: function testMoment(uint256[] x, uint256[] y, uint64 momentIndex) returns()
func (_Test *TestTransactor) TestMoment(opts *bind.TransactOpts, x []*big.Int, y []*big.Int, momentIndex uint64) (*types.Transaction, error) {
	return _Test.contract.Transact(opts, "testMoment", x, y, momentIndex)
}

// TestMoment is a paid mutator transaction binding the contract method 0xe9d6b2bc.
//
// Solidity: function testMoment(uint256[] x, uint256[] y, uint64 momentIndex) returns()
func (_Test *TestSession) TestMoment(x []*big.Int, y []*big.Int, momentIndex uint64) (*types.Transaction, error) {
	return _Test.Contract.TestMoment(&_Test.TransactOpts, x, y, momentIndex)
}

// TestMoment is a paid mutator transaction binding the contract method 0xe9d6b2bc.
//
// Solidity: function testMoment(uint256[] x, uint256[] y, uint64 momentIndex) returns()
func (_Test *TestTransactorSession) TestMoment(x []*big.Int, y []*big.Int, momentIndex uint64) (*types.Transaction, error) {
	return _Test.Contract.TestMoment(&_Test.TransactOpts, x, y, momentIndex)
}

