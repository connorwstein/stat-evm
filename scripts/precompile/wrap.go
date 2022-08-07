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
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"message\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"res\",\"type\":\"int256\"}],\"name\":\"Debug\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"fitted\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getFitted\",\"outputs\":[{\"internalType\":\"int256[][]\",\"name\":\"\",\"type\":\"int256[][]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getIPFSCoeffs\",\"outputs\":[{\"internalType\":\"int256[][]\",\"name\":\"\",\"type\":\"int256[][]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getIPFSIntercepts\",\"outputs\":[{\"internalType\":\"int256[]\",\"name\":\"\",\"type\":\"int256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getLastSample\",\"outputs\":[{\"internalType\":\"int256[]\",\"name\":\"\",\"type\":\"int256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMatrixMulti\",\"outputs\":[{\"internalType\":\"int256[][]\",\"name\":\"\",\"type\":\"int256[][]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"ipfsCoeffs\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"ipfsIntercepts\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ipfsMomentRes\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"med\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"moment\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"product\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"sample\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"fitType\",\"type\":\"uint256\"},{\"internalType\":\"int256[][]\",\"name\":\"X\",\"type\":\"int256[][]\"},{\"internalType\":\"int256[][]\",\"name\":\"Y\",\"type\":\"int256[][]\"}],\"name\":\"testFit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"ipfsHash\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"fitType\",\"type\":\"uint256\"}],\"name\":\"testIPFSFit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"ipfsHash\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"moment\",\"type\":\"uint256\"}],\"name\":\"testIPFSMoment\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int256[][]\",\"name\":\"a\",\"type\":\"int256[][]\"},{\"internalType\":\"int256[][]\",\"name\":\"b\",\"type\":\"int256[][]\"}],\"name\":\"testMatrixMult\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"vals\",\"type\":\"uint256[]\"}],\"name\":\"testMedian\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"v1\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"v2\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"v3\",\"type\":\"uint256[]\"}],\"name\":\"testMoment\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"distributionType\",\"type\":\"uint256\"},{\"internalType\":\"int256\",\"name\":\"a\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"b\",\"type\":\"int256\"},{\"internalType\":\"uint256\",\"name\":\"numSamples\",\"type\":\"uint256\"}],\"name\":\"testSampler\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x6080604052730300000000000000000000000000000000000001600860006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550730300000000000000000000000000000000000004600960006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550730300000000000000000000000000000000000006600a60006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550730300000000000000000000000000000000000005600b60006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550730300000000000000000000000000000000000008600c60006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550730300000000000000000000000000000000000009600d60006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550730300000000000000000000000000000000000007600e60006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555034801561026357600080fd5b50611cdc806102736000396000f3fe608060405234801561001057600080fd5b506004361061012c5760003560e01c806395697c0c116100ad578063df98490111610071578063df98490114610343578063f16d968c14610361578063f213a5a11461037f578063f505da8a1461039b578063fc784aa1146103b75761012c565b806395697c0c14610289578063ba88c3b7146102a7578063baa40cc1146102d7578063c44073ca14610307578063cd53c747146103255761012c565b80636ba87539116100f45780636ba87539146101d35780637529916e146101ef57806382bb60a71461020d57806382f044991461023d57806383b6033c1461026d5761012c565b806310fa18781461013157806312f758a014610161578063385d763f1461017d5780633b1c29ef1461019b5780635fd70772146101b7575b600080fd5b61014b60048036038101906101469190610de5565b6103d5565b6040516101589190610e2b565b60405180910390f35b61017b60048036038101906101769190610f9f565b6103f9565b005b6101856104a4565b6040516101929190611039565b60405180910390f35b6101b560048036038101906101b09190611224565b6104aa565b005b6101d160048036038101906101cc9190611351565b610567565b005b6101ed60048036038101906101e891906113ad565b61060f565b005b6101f76106cf565b60405161020491906114f6565b60405180910390f35b61022760048036038101906102229190611518565b610727565b6040516102349190610e2b565b60405180910390f35b61025760048036038101906102529190611518565b610764565b6040516102649190610e2b565b60405180910390f35b61028760048036038101906102829190611558565b6107a1565b005b610291610846565b60405161029e91906116d2565b60405180910390f35b6102c160048036038101906102bc9190611518565b6108e5565b6040516102ce9190610e2b565b60405180910390f35b6102f160048036038101906102ec9190610de5565b610922565b6040516102fe9190610e2b565b60405180910390f35b61030f610946565b60405161031c91906116d2565b60405180910390f35b61032d6109e5565b60405161033a9190611039565b60405180910390f35b61034b6109eb565b6040516103589190611039565b60405180910390f35b6103696109f1565b60405161037691906114f6565b60405180910390f35b610399600480360381019061039491906116f4565b610a49565b005b6103b560048036038101906103b09190611351565b610b0c565b005b6103bf610bed565b6040516103cc91906116d2565b60405180910390f35b600281815481106103e557600080fd5b906000526020600020016000915090505481565b600a60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166320db7f0f8484846040518463ffffffff1660e01b815260040161045893929190611819565b602060405180830381865afa158015610475573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906104999190611873565b600181905550505050565b60075481565b600b60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166359c6945683836040518363ffffffff1660e01b81526004016105079291906118a0565b600060405180830381865afa158015610524573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f8201168201806040525081019061054d9190611a38565b60039080519060200190610562929190610c8c565b505050565b600c60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663a967f50383836040518363ffffffff1660e01b81526004016105c4929190611b09565b602060405180830381865afa1580156105e1573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906106059190611873565b6007819055505050565b600e60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663d3a5d0db8484846040518463ffffffff1660e01b815260040161066e93929190611b39565b600060405180830381865afa15801561068b573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f820116820180604052508101906106b49190611a38565b600490805190602001906106c9929190610c8c565b50505050565b6060600680548060200260200160405190810160405280929190818152602001828054801561071d57602002820191906000526020600020905b815481526020019060010190808311610709575b5050505050905090565b6005828154811061073757600080fd5b90600052602060002001818154811061074f57600080fd5b90600052602060002001600091509150505481565b6004828154811061077457600080fd5b90600052602060002001818154811061078c57600080fd5b90600052602060002001600091509150505481565b600860009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663e8787f41826040518263ffffffff1660e01b81526004016107fc9190611b7e565b602060405180830381865afa158015610819573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061083d9190611873565b60008190555050565b60606005805480602002602001604051908101604052809291908181526020016000905b828210156108dc578382906000526020600020018054806020026020016040519081016040528092919081815260200182805480156108c857602002820191906000526020600020905b8154815260200190600101908083116108b4575b50505050508152602001906001019061086a565b50505050905090565b600382815481106108f557600080fd5b90600052602060002001818154811061090d57600080fd5b90600052602060002001600091509150505481565b6006818154811061093257600080fd5b906000526020600020016000915090505481565b60606004805480602002602001604051908101604052809291908181526020016000905b828210156109dc578382906000526020600020018054806020026020016040519081016040528092919081815260200182805480156109c857602002820191906000526020600020905b8154815260200190600101908083116109b4575b50505050508152602001906001019061096a565b50505050905090565b60005481565b60015481565b60606002805480602002602001604051908101604052809291908181526020018280548015610a3f57602002820191906000526020600020905b815481526020019060010190808311610a2b575b5050505050905090565b600960009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663cca718bf858585856040518563ffffffff1660e01b8152600401610aaa9493929190611ba0565b600060405180830381865afa158015610ac7573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f82011682018060405250810190610af09190611be5565b60029080519060200190610b05929190610cec565b5050505050565b600d60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16635466c0b083836040518363ffffffff1660e01b8152600401610b69929190611b09565b600060405180830381865afa158015610b86573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f82011682018060405250810190610baf9190611c2e565b6005600060066000849190509080519060200190610bce929190610cec565b50839190509080519060200190610be6929190610c8c565b5050505050565b60606003805480602002602001604051908101604052809291908181526020016000905b82821015610c8357838290600052602060002001805480602002602001604051908101604052809291908181526020018280548015610c6f57602002820191906000526020600020905b815481526020019060010190808311610c5b575b505050505081526020019060010190610c11565b50505050905090565b828054828255906000526020600020908101928215610cdb579160200282015b82811115610cda578251829080519060200190610cca929190610cec565b5091602001919060010190610cac565b5b509050610ce89190610d39565b5090565b828054828255906000526020600020908101928215610d28579160200282015b82811115610d27578251825591602001919060010190610d0c565b5b509050610d359190610d5d565b5090565b5b80821115610d595760008181610d509190610d7a565b50600101610d3a565b5090565b5b80821115610d76576000816000905550600101610d5e565b5090565b5080546000825590600052602060002090810190610d989190610d5d565b50565b6000604051905090565b600080fd5b600080fd5b6000819050919050565b610dc281610daf565b8114610dcd57600080fd5b50565b600081359050610ddf81610db9565b92915050565b600060208284031215610dfb57610dfa610da5565b5b6000610e0984828501610dd0565b91505092915050565b6000819050919050565b610e2581610e12565b82525050565b6000602082019050610e406000830184610e1c565b92915050565b600080fd5b6000601f19601f8301169050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b610e9482610e4b565b810181811067ffffffffffffffff82111715610eb357610eb2610e5c565b5b80604052505050565b6000610ec6610d9b565b9050610ed28282610e8b565b919050565b600067ffffffffffffffff821115610ef257610ef1610e5c565b5b602082029050602081019050919050565b600080fd5b6000610f1b610f1684610ed7565b610ebc565b90508083825260208201905060208402830185811115610f3e57610f3d610f03565b5b835b81811015610f675780610f538882610dd0565b845260208401935050602081019050610f40565b5050509392505050565b600082601f830112610f8657610f85610e46565b5b8135610f96848260208601610f08565b91505092915050565b600080600060608486031215610fb857610fb7610da5565b5b6000610fc686828701610dd0565b935050602084013567ffffffffffffffff811115610fe757610fe6610daa565b5b610ff386828701610f71565b925050604084013567ffffffffffffffff81111561101457611013610daa565b5b61102086828701610f71565b9150509250925092565b61103381610daf565b82525050565b600060208201905061104e600083018461102a565b92915050565b600067ffffffffffffffff82111561106f5761106e610e5c565b5b602082029050602081019050919050565b600067ffffffffffffffff82111561109b5761109a610e5c565b5b602082029050602081019050919050565b6110b581610e12565b81146110c057600080fd5b50565b6000813590506110d2816110ac565b92915050565b60006110eb6110e684611080565b610ebc565b9050808382526020820190506020840283018581111561110e5761110d610f03565b5b835b81811015611137578061112388826110c3565b845260208401935050602081019050611110565b5050509392505050565b600082601f83011261115657611155610e46565b5b81356111668482602086016110d8565b91505092915050565b600061118261117d84611054565b610ebc565b905080838252602082019050602084028301858111156111a5576111a4610f03565b5b835b818110156111ec57803567ffffffffffffffff8111156111ca576111c9610e46565b5b8086016111d78982611141565b855260208501945050506020810190506111a7565b5050509392505050565b600082601f83011261120b5761120a610e46565b5b813561121b84826020860161116f565b91505092915050565b6000806040838503121561123b5761123a610da5565b5b600083013567ffffffffffffffff81111561125957611258610daa565b5b611265858286016111f6565b925050602083013567ffffffffffffffff81111561128657611285610daa565b5b611292858286016111f6565b9150509250929050565b600080fd5b600067ffffffffffffffff8211156112bc576112bb610e5c565b5b6112c582610e4b565b9050602081019050919050565b82818337600083830152505050565b60006112f46112ef846112a1565b610ebc565b9050828152602081018484840111156113105761130f61129c565b5b61131b8482856112d2565b509392505050565b600082601f83011261133857611337610e46565b5b81356113488482602086016112e1565b91505092915050565b6000806040838503121561136857611367610da5565b5b600083013567ffffffffffffffff81111561138657611385610daa565b5b61139285828601611323565b92505060206113a385828601610dd0565b9150509250929050565b6000806000606084860312156113c6576113c5610da5565b5b60006113d486828701610dd0565b935050602084013567ffffffffffffffff8111156113f5576113f4610daa565b5b611401868287016111f6565b925050604084013567ffffffffffffffff81111561142257611421610daa565b5b61142e868287016111f6565b9150509250925092565b600081519050919050565b600082825260208201905092915050565b6000819050602082019050919050565b61146d81610e12565b82525050565b600061147f8383611464565b60208301905092915050565b6000602082019050919050565b60006114a382611438565b6114ad8185611443565b93506114b883611454565b8060005b838110156114e95781516114d08882611473565b97506114db8361148b565b9250506001810190506114bc565b5085935050505092915050565b600060208201905081810360008301526115108184611498565b905092915050565b6000806040838503121561152f5761152e610da5565b5b600061153d85828601610dd0565b925050602061154e85828601610dd0565b9150509250929050565b60006020828403121561156e5761156d610da5565b5b600082013567ffffffffffffffff81111561158c5761158b610daa565b5b61159884828501610f71565b91505092915050565b600081519050919050565b600082825260208201905092915050565b6000819050602082019050919050565b600082825260208201905092915050565b60006115e982611438565b6115f381856115cd565b93506115fe83611454565b8060005b8381101561162f5781516116168882611473565b97506116218361148b565b925050600181019050611602565b5085935050505092915050565b600061164883836115de565b905092915050565b6000602082019050919050565b6000611668826115a1565b61167281856115ac565b935083602082028501611684856115bd565b8060005b858110156116c057848403895281516116a1858261163c565b94506116ac83611650565b925060208a01995050600181019050611688565b50829750879550505050505092915050565b600060208201905081810360008301526116ec818461165d565b905092915050565b6000806000806080858703121561170e5761170d610da5565b5b600061171c87828801610dd0565b945050602061172d878288016110c3565b935050604061173e878288016110c3565b925050606061174f87828801610dd0565b91505092959194509250565b600081519050919050565b600082825260208201905092915050565b6000819050602082019050919050565b61179081610daf565b82525050565b60006117a28383611787565b60208301905092915050565b6000602082019050919050565b60006117c68261175b565b6117d08185611766565b93506117db83611777565b8060005b8381101561180c5781516117f38882611796565b97506117fe836117ae565b9250506001810190506117df565b5085935050505092915050565b600060608201905061182e600083018661102a565b818103602083015261184081856117bb565b9050818103604083015261185481846117bb565b9050949350505050565b60008151905061186d81610db9565b92915050565b60006020828403121561188957611888610da5565b5b60006118978482850161185e565b91505092915050565b600060408201905081810360008301526118ba818561165d565b905081810360208301526118ce818461165d565b90509392505050565b6000815190506118e6816110ac565b92915050565b60006118ff6118fa84611080565b610ebc565b9050808382526020820190506020840283018581111561192257611921610f03565b5b835b8181101561194b578061193788826118d7565b845260208401935050602081019050611924565b5050509392505050565b600082601f83011261196a57611969610e46565b5b815161197a8482602086016118ec565b91505092915050565b600061199661199184611054565b610ebc565b905080838252602082019050602084028301858111156119b9576119b8610f03565b5b835b81811015611a0057805167ffffffffffffffff8111156119de576119dd610e46565b5b8086016119eb8982611955565b855260208501945050506020810190506119bb565b5050509392505050565b600082601f830112611a1f57611a1e610e46565b5b8151611a2f848260208601611983565b91505092915050565b600060208284031215611a4e57611a4d610da5565b5b600082015167ffffffffffffffff811115611a6c57611a6b610daa565b5b611a7884828501611a0a565b91505092915050565b600081519050919050565b600082825260208201905092915050565b60005b83811015611abb578082015181840152602081019050611aa0565b83811115611aca576000848401525b50505050565b6000611adb82611a81565b611ae58185611a8c565b9350611af5818560208601611a9d565b611afe81610e4b565b840191505092915050565b60006040820190508181036000830152611b238185611ad0565b9050611b32602083018461102a565b9392505050565b6000606082019050611b4e600083018661102a565b8181036020830152611b60818561165d565b90508181036040830152611b74818461165d565b9050949350505050565b60006020820190508181036000830152611b9881846117bb565b905092915050565b6000608082019050611bb5600083018761102a565b611bc26020830186610e1c565b611bcf6040830185610e1c565b611bdc606083018461102a565b95945050505050565b600060208284031215611bfb57611bfa610da5565b5b600082015167ffffffffffffffff811115611c1957611c18610daa565b5b611c2584828501611955565b91505092915050565b60008060408385031215611c4557611c44610da5565b5b600083015167ffffffffffffffff811115611c6357611c62610daa565b5b611c6f85828601611a0a565b925050602083015167ffffffffffffffff811115611c9057611c8f610daa565b5b611c9c85828601611955565b915050925092905056fea2646970667358221220df8ca323223a3b8d3b7229d3145dd2284e7ebbd9d2d8f31c131c10b2b00ff18264736f6c634300080f0033",
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

// GetIPFSCoeffs is a free data retrieval call binding the contract method 0x95697c0c.
//
// Solidity: function getIPFSCoeffs() view returns(int256[][])
func (_Main *MainCaller) GetIPFSCoeffs(opts *bind.CallOpts) ([][]*big.Int, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "getIPFSCoeffs")

	if err != nil {
		return *new([][]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([][]*big.Int)).(*[][]*big.Int)

	return out0, err

}

// GetIPFSCoeffs is a free data retrieval call binding the contract method 0x95697c0c.
//
// Solidity: function getIPFSCoeffs() view returns(int256[][])
func (_Main *MainSession) GetIPFSCoeffs() ([][]*big.Int, error) {
	return _Main.Contract.GetIPFSCoeffs(&_Main.CallOpts)
}

// GetIPFSCoeffs is a free data retrieval call binding the contract method 0x95697c0c.
//
// Solidity: function getIPFSCoeffs() view returns(int256[][])
func (_Main *MainCallerSession) GetIPFSCoeffs() ([][]*big.Int, error) {
	return _Main.Contract.GetIPFSCoeffs(&_Main.CallOpts)
}

// GetIPFSIntercepts is a free data retrieval call binding the contract method 0x7529916e.
//
// Solidity: function getIPFSIntercepts() view returns(int256[])
func (_Main *MainCaller) GetIPFSIntercepts(opts *bind.CallOpts) ([]*big.Int, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "getIPFSIntercepts")

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetIPFSIntercepts is a free data retrieval call binding the contract method 0x7529916e.
//
// Solidity: function getIPFSIntercepts() view returns(int256[])
func (_Main *MainSession) GetIPFSIntercepts() ([]*big.Int, error) {
	return _Main.Contract.GetIPFSIntercepts(&_Main.CallOpts)
}

// GetIPFSIntercepts is a free data retrieval call binding the contract method 0x7529916e.
//
// Solidity: function getIPFSIntercepts() view returns(int256[])
func (_Main *MainCallerSession) GetIPFSIntercepts() ([]*big.Int, error) {
	return _Main.Contract.GetIPFSIntercepts(&_Main.CallOpts)
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

// IpfsCoeffs is a free data retrieval call binding the contract method 0x82bb60a7.
//
// Solidity: function ipfsCoeffs(uint256 , uint256 ) view returns(int256)
func (_Main *MainCaller) IpfsCoeffs(opts *bind.CallOpts, arg0 *big.Int, arg1 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "ipfsCoeffs", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// IpfsCoeffs is a free data retrieval call binding the contract method 0x82bb60a7.
//
// Solidity: function ipfsCoeffs(uint256 , uint256 ) view returns(int256)
func (_Main *MainSession) IpfsCoeffs(arg0 *big.Int, arg1 *big.Int) (*big.Int, error) {
	return _Main.Contract.IpfsCoeffs(&_Main.CallOpts, arg0, arg1)
}

// IpfsCoeffs is a free data retrieval call binding the contract method 0x82bb60a7.
//
// Solidity: function ipfsCoeffs(uint256 , uint256 ) view returns(int256)
func (_Main *MainCallerSession) IpfsCoeffs(arg0 *big.Int, arg1 *big.Int) (*big.Int, error) {
	return _Main.Contract.IpfsCoeffs(&_Main.CallOpts, arg0, arg1)
}

// IpfsIntercepts is a free data retrieval call binding the contract method 0xbaa40cc1.
//
// Solidity: function ipfsIntercepts(uint256 ) view returns(int256)
func (_Main *MainCaller) IpfsIntercepts(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "ipfsIntercepts", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// IpfsIntercepts is a free data retrieval call binding the contract method 0xbaa40cc1.
//
// Solidity: function ipfsIntercepts(uint256 ) view returns(int256)
func (_Main *MainSession) IpfsIntercepts(arg0 *big.Int) (*big.Int, error) {
	return _Main.Contract.IpfsIntercepts(&_Main.CallOpts, arg0)
}

// IpfsIntercepts is a free data retrieval call binding the contract method 0xbaa40cc1.
//
// Solidity: function ipfsIntercepts(uint256 ) view returns(int256)
func (_Main *MainCallerSession) IpfsIntercepts(arg0 *big.Int) (*big.Int, error) {
	return _Main.Contract.IpfsIntercepts(&_Main.CallOpts, arg0)
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

// TestIPFSFit is a paid mutator transaction binding the contract method 0xf505da8a.
//
// Solidity: function testIPFSFit(string ipfsHash, uint256 fitType) returns()
func (_Main *MainTransactor) TestIPFSFit(opts *bind.TransactOpts, ipfsHash string, fitType *big.Int) (*types.Transaction, error) {
	return _Main.contract.Transact(opts, "testIPFSFit", ipfsHash, fitType)
}

// TestIPFSFit is a paid mutator transaction binding the contract method 0xf505da8a.
//
// Solidity: function testIPFSFit(string ipfsHash, uint256 fitType) returns()
func (_Main *MainSession) TestIPFSFit(ipfsHash string, fitType *big.Int) (*types.Transaction, error) {
	return _Main.Contract.TestIPFSFit(&_Main.TransactOpts, ipfsHash, fitType)
}

// TestIPFSFit is a paid mutator transaction binding the contract method 0xf505da8a.
//
// Solidity: function testIPFSFit(string ipfsHash, uint256 fitType) returns()
func (_Main *MainTransactorSession) TestIPFSFit(ipfsHash string, fitType *big.Int) (*types.Transaction, error) {
	return _Main.Contract.TestIPFSFit(&_Main.TransactOpts, ipfsHash, fitType)
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

