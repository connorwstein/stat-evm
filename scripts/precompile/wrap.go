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
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"message\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"res\",\"type\":\"int256\"}],\"name\":\"Debug\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"fitted\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getFitted\",\"outputs\":[{\"internalType\":\"int256[][]\",\"name\":\"\",\"type\":\"int256[][]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getIPFSCoeffs\",\"outputs\":[{\"internalType\":\"int256[][]\",\"name\":\"\",\"type\":\"int256[][]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getIPFSIntercepts\",\"outputs\":[{\"internalType\":\"int256[]\",\"name\":\"\",\"type\":\"int256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getLastSample\",\"outputs\":[{\"internalType\":\"int256[]\",\"name\":\"\",\"type\":\"int256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMatrixMulti\",\"outputs\":[{\"internalType\":\"int256[][]\",\"name\":\"\",\"type\":\"int256[][]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"ipfsCoeffs\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"ipfsIntercepts\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ipfsMomentRes\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"med\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"moment\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"predicted\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"product\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"sample\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"fitType\",\"type\":\"uint256\"},{\"internalType\":\"int256[][]\",\"name\":\"X\",\"type\":\"int256[][]\"},{\"internalType\":\"int256[][]\",\"name\":\"Y\",\"type\":\"int256[][]\"}],\"name\":\"testFit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"ipfsHash\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"fitType\",\"type\":\"uint256\"}],\"name\":\"testIPFSFit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"ipfsHash\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"moment\",\"type\":\"uint256\"}],\"name\":\"testIPFSMoment\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int256[][]\",\"name\":\"a\",\"type\":\"int256[][]\"},{\"internalType\":\"int256[][]\",\"name\":\"b\",\"type\":\"int256[][]\"}],\"name\":\"testMatrixMult\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"vals\",\"type\":\"uint256[]\"}],\"name\":\"testMedian\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"v1\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"v2\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"v3\",\"type\":\"uint256[]\"}],\"name\":\"testMoment\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int256[][]\",\"name\":\"dependentVars\",\"type\":\"int256[][]\"}],\"name\":\"testPrediction\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"distributionType\",\"type\":\"uint256\"},{\"internalType\":\"int256\",\"name\":\"a\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"b\",\"type\":\"int256\"},{\"internalType\":\"uint256\",\"name\":\"numSamples\",\"type\":\"uint256\"}],\"name\":\"testSampler\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x6080604052730300000000000000000000000000000000000001600960006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550730300000000000000000000000000000000000004600a60006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550730300000000000000000000000000000000000006600b60006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550730300000000000000000000000000000000000005600c60006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550730300000000000000000000000000000000000008600d60006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550730300000000000000000000000000000000000009600e60006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550730300000000000000000000000000000000000007600f60006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555034801561026357600080fd5b50612044806102736000396000f3fe608060405234801561001057600080fd5b50600436106101425760003560e01c8063b7aec839116100b8578063d0148a671161007c578063d0148a6714610377578063df98490114610393578063f16d968c146103b1578063f213a5a1146103cf578063f505da8a146103eb578063fc784aa11461040757610142565b8063b7aec839146102bd578063ba88c3b7146102db578063baa40cc11461030b578063c44073ca1461033b578063cd53c7471461035957610142565b80636ba875391161010a5780636ba87539146101e95780637529916e1461020557806382bb60a71461022357806382f044991461025357806383b6033c1461028357806395697c0c1461029f57610142565b806310fa18781461014757806312f758a014610177578063385d763f146101935780633b1c29ef146101b15780635fd70772146101cd575b600080fd5b610161600480360381019061015c9190610f79565b610425565b60405161016e9190610fbf565b60405180910390f35b610191600480360381019061018c9190611133565b610449565b005b61019b6104f4565b6040516101a891906111cd565b60405180910390f35b6101cb60048036038101906101c691906113b8565b6104fa565b005b6101e760048036038101906101e291906114e5565b6105b7565b005b61020360048036038101906101fe9190611541565b61065f565b005b61020d61071f565b60405161021a919061168a565b60405180910390f35b61023d600480360381019061023891906116ac565b610777565b60405161024a9190610fbf565b60405180910390f35b61026d600480360381019061026891906116ac565b6107b4565b60405161027a9190610fbf565b60405180910390f35b61029d600480360381019061029891906116ec565b6107f1565b005b6102a7610896565b6040516102b49190611866565b60405180910390f35b6102c5610935565b6040516102d29190610fbf565b60405180910390f35b6102f560048036038101906102f091906116ac565b61093b565b6040516103029190610fbf565b60405180910390f35b61032560048036038101906103209190610f79565b610978565b6040516103329190610fbf565b60405180910390f35b61034361099c565b6040516103509190611866565b60405180910390f35b610361610a3b565b60405161036e91906111cd565b60405180910390f35b610391600480360381019061038c9190611888565b610a41565b005b61039b610b7f565b6040516103a891906111cd565b60405180910390f35b6103b9610b85565b6040516103c6919061168a565b60405180910390f35b6103e960048036038101906103e491906118d1565b610bdd565b005b610405600480360381019061040091906114e5565b610ca0565b005b61040f610d81565b60405161041c9190611866565b60405180910390f35b6002818154811061043557600080fd5b906000526020600020016000915090505481565b600b60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166320db7f0f8484846040518463ffffffff1660e01b81526004016104a8939291906119f6565b602060405180830381865afa1580156104c5573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906104e99190611a50565b600181905550505050565b60075481565b600c60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166359c6945683836040518363ffffffff1660e01b8152600401610557929190611a7d565b600060405180830381865afa158015610574573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f8201168201806040525081019061059d9190611c15565b600390805190602001906105b2929190610e20565b505050565b600d60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663a967f50383836040518363ffffffff1660e01b8152600401610614929190611ce6565b602060405180830381865afa158015610631573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906106559190611a50565b6007819055505050565b600f60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663d3a5d0db8484846040518463ffffffff1660e01b81526004016106be93929190611d16565b600060405180830381865afa1580156106db573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f820116820180604052508101906107049190611c15565b60049080519060200190610719929190610e20565b50505050565b6060600680548060200260200160405190810160405280929190818152602001828054801561076d57602002820191906000526020600020905b815481526020019060010190808311610759575b5050505050905090565b6005828154811061078757600080fd5b90600052602060002001818154811061079f57600080fd5b90600052602060002001600091509150505481565b600482815481106107c457600080fd5b9060005260206000200181815481106107dc57600080fd5b90600052602060002001600091509150505481565b600960009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663e8787f41826040518263ffffffff1660e01b815260040161084c9190611d5b565b602060405180830381865afa158015610869573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061088d9190611a50565b60008190555050565b60606005805480602002602001604051908101604052809291908181526020016000905b8282101561092c5783829060005260206000200180548060200260200160405190810160405280929190818152602001828054801561091857602002820191906000526020600020905b815481526020019060010190808311610904575b5050505050815260200190600101906108ba565b50505050905090565b60085481565b6003828154811061094b57600080fd5b90600052602060002001818154811061096357600080fd5b90600052602060002001600091509150505481565b6006818154811061098857600080fd5b906000526020600020016000915090505481565b60606004805480602002602001604051908101604052809291908181526020016000905b82821015610a3257838290600052602060002001805480602002602001604051908101604052809291908181526020018280548015610a1e57602002820191906000526020600020905b815481526020019060010190808311610a0a575b5050505050815260200190600101906109c0565b50505050905090565b60005481565b6000600c60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166359c6945683610a8a610896565b6040518363ffffffff1660e01b8152600401610aa7929190611a7d565b600060405180830381865afa158015610ac4573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f82011682018060405250810190610aed9190611c15565b90506000610af961071f565b9050620f424081600081518110610b1357610b12611d7d565b5b6020026020010151610b259190611e0a565b620f424083600081518110610b3d57610b3c611d7d565b5b6020026020010151600081518110610b5857610b57611d7d565b5b6020026020010151610b6a9190611e0a565b610b749190611e74565b600881905550505050565b60015481565b60606002805480602002602001604051908101604052809291908181526020018280548015610bd357602002820191906000526020600020905b815481526020019060010190808311610bbf575b5050505050905090565b600a60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663cca718bf858585856040518563ffffffff1660e01b8152600401610c3e9493929190611f08565b600060405180830381865afa158015610c5b573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f82011682018060405250810190610c849190611f4d565b60029080519060200190610c99929190610e80565b5050505050565b600e60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663d80fc60583836040518363ffffffff1660e01b8152600401610cfd929190611ce6565b600060405180830381865afa158015610d1a573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f82011682018060405250810190610d439190611f96565b6005600060066000849190509080519060200190610d62929190610e80565b50839190509080519060200190610d7a929190610e20565b5050505050565b60606003805480602002602001604051908101604052809291908181526020016000905b82821015610e1757838290600052602060002001805480602002602001604051908101604052809291908181526020018280548015610e0357602002820191906000526020600020905b815481526020019060010190808311610def575b505050505081526020019060010190610da5565b50505050905090565b828054828255906000526020600020908101928215610e6f579160200282015b82811115610e6e578251829080519060200190610e5e929190610e80565b5091602001919060010190610e40565b5b509050610e7c9190610ecd565b5090565b828054828255906000526020600020908101928215610ebc579160200282015b82811115610ebb578251825591602001919060010190610ea0565b5b509050610ec99190610ef1565b5090565b5b80821115610eed5760008181610ee49190610f0e565b50600101610ece565b5090565b5b80821115610f0a576000816000905550600101610ef2565b5090565b5080546000825590600052602060002090810190610f2c9190610ef1565b50565b6000604051905090565b600080fd5b600080fd5b6000819050919050565b610f5681610f43565b8114610f6157600080fd5b50565b600081359050610f7381610f4d565b92915050565b600060208284031215610f8f57610f8e610f39565b5b6000610f9d84828501610f64565b91505092915050565b6000819050919050565b610fb981610fa6565b82525050565b6000602082019050610fd46000830184610fb0565b92915050565b600080fd5b6000601f19601f8301169050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b61102882610fdf565b810181811067ffffffffffffffff8211171561104757611046610ff0565b5b80604052505050565b600061105a610f2f565b9050611066828261101f565b919050565b600067ffffffffffffffff82111561108657611085610ff0565b5b602082029050602081019050919050565b600080fd5b60006110af6110aa8461106b565b611050565b905080838252602082019050602084028301858111156110d2576110d1611097565b5b835b818110156110fb57806110e78882610f64565b8452602084019350506020810190506110d4565b5050509392505050565b600082601f83011261111a57611119610fda565b5b813561112a84826020860161109c565b91505092915050565b60008060006060848603121561114c5761114b610f39565b5b600061115a86828701610f64565b935050602084013567ffffffffffffffff81111561117b5761117a610f3e565b5b61118786828701611105565b925050604084013567ffffffffffffffff8111156111a8576111a7610f3e565b5b6111b486828701611105565b9150509250925092565b6111c781610f43565b82525050565b60006020820190506111e260008301846111be565b92915050565b600067ffffffffffffffff82111561120357611202610ff0565b5b602082029050602081019050919050565b600067ffffffffffffffff82111561122f5761122e610ff0565b5b602082029050602081019050919050565b61124981610fa6565b811461125457600080fd5b50565b60008135905061126681611240565b92915050565b600061127f61127a84611214565b611050565b905080838252602082019050602084028301858111156112a2576112a1611097565b5b835b818110156112cb57806112b78882611257565b8452602084019350506020810190506112a4565b5050509392505050565b600082601f8301126112ea576112e9610fda565b5b81356112fa84826020860161126c565b91505092915050565b6000611316611311846111e8565b611050565b9050808382526020820190506020840283018581111561133957611338611097565b5b835b8181101561138057803567ffffffffffffffff81111561135e5761135d610fda565b5b80860161136b89826112d5565b8552602085019450505060208101905061133b565b5050509392505050565b600082601f83011261139f5761139e610fda565b5b81356113af848260208601611303565b91505092915050565b600080604083850312156113cf576113ce610f39565b5b600083013567ffffffffffffffff8111156113ed576113ec610f3e565b5b6113f98582860161138a565b925050602083013567ffffffffffffffff81111561141a57611419610f3e565b5b6114268582860161138a565b9150509250929050565b600080fd5b600067ffffffffffffffff8211156114505761144f610ff0565b5b61145982610fdf565b9050602081019050919050565b82818337600083830152505050565b600061148861148384611435565b611050565b9050828152602081018484840111156114a4576114a3611430565b5b6114af848285611466565b509392505050565b600082601f8301126114cc576114cb610fda565b5b81356114dc848260208601611475565b91505092915050565b600080604083850312156114fc576114fb610f39565b5b600083013567ffffffffffffffff81111561151a57611519610f3e565b5b611526858286016114b7565b925050602061153785828601610f64565b9150509250929050565b60008060006060848603121561155a57611559610f39565b5b600061156886828701610f64565b935050602084013567ffffffffffffffff81111561158957611588610f3e565b5b6115958682870161138a565b925050604084013567ffffffffffffffff8111156115b6576115b5610f3e565b5b6115c28682870161138a565b9150509250925092565b600081519050919050565b600082825260208201905092915050565b6000819050602082019050919050565b61160181610fa6565b82525050565b600061161383836115f8565b60208301905092915050565b6000602082019050919050565b6000611637826115cc565b61164181856115d7565b935061164c836115e8565b8060005b8381101561167d5781516116648882611607565b975061166f8361161f565b925050600181019050611650565b5085935050505092915050565b600060208201905081810360008301526116a4818461162c565b905092915050565b600080604083850312156116c3576116c2610f39565b5b60006116d185828601610f64565b92505060206116e285828601610f64565b9150509250929050565b60006020828403121561170257611701610f39565b5b600082013567ffffffffffffffff8111156117205761171f610f3e565b5b61172c84828501611105565b91505092915050565b600081519050919050565b600082825260208201905092915050565b6000819050602082019050919050565b600082825260208201905092915050565b600061177d826115cc565b6117878185611761565b9350611792836115e8565b8060005b838110156117c35781516117aa8882611607565b97506117b58361161f565b925050600181019050611796565b5085935050505092915050565b60006117dc8383611772565b905092915050565b6000602082019050919050565b60006117fc82611735565b6118068185611740565b93508360208202850161181885611751565b8060005b85811015611854578484038952815161183585826117d0565b9450611840836117e4565b925060208a0199505060018101905061181c565b50829750879550505050505092915050565b6000602082019050818103600083015261188081846117f1565b905092915050565b60006020828403121561189e5761189d610f39565b5b600082013567ffffffffffffffff8111156118bc576118bb610f3e565b5b6118c88482850161138a565b91505092915050565b600080600080608085870312156118eb576118ea610f39565b5b60006118f987828801610f64565b945050602061190a87828801611257565b935050604061191b87828801611257565b925050606061192c87828801610f64565b91505092959194509250565b600081519050919050565b600082825260208201905092915050565b6000819050602082019050919050565b61196d81610f43565b82525050565b600061197f8383611964565b60208301905092915050565b6000602082019050919050565b60006119a382611938565b6119ad8185611943565b93506119b883611954565b8060005b838110156119e95781516119d08882611973565b97506119db8361198b565b9250506001810190506119bc565b5085935050505092915050565b6000606082019050611a0b60008301866111be565b8181036020830152611a1d8185611998565b90508181036040830152611a318184611998565b9050949350505050565b600081519050611a4a81610f4d565b92915050565b600060208284031215611a6657611a65610f39565b5b6000611a7484828501611a3b565b91505092915050565b60006040820190508181036000830152611a9781856117f1565b90508181036020830152611aab81846117f1565b90509392505050565b600081519050611ac381611240565b92915050565b6000611adc611ad784611214565b611050565b90508083825260208201905060208402830185811115611aff57611afe611097565b5b835b81811015611b285780611b148882611ab4565b845260208401935050602081019050611b01565b5050509392505050565b600082601f830112611b4757611b46610fda565b5b8151611b57848260208601611ac9565b91505092915050565b6000611b73611b6e846111e8565b611050565b90508083825260208201905060208402830185811115611b9657611b95611097565b5b835b81811015611bdd57805167ffffffffffffffff811115611bbb57611bba610fda565b5b808601611bc88982611b32565b85526020850194505050602081019050611b98565b5050509392505050565b600082601f830112611bfc57611bfb610fda565b5b8151611c0c848260208601611b60565b91505092915050565b600060208284031215611c2b57611c2a610f39565b5b600082015167ffffffffffffffff811115611c4957611c48610f3e565b5b611c5584828501611be7565b91505092915050565b600081519050919050565b600082825260208201905092915050565b60005b83811015611c98578082015181840152602081019050611c7d565b83811115611ca7576000848401525b50505050565b6000611cb882611c5e565b611cc28185611c69565b9350611cd2818560208601611c7a565b611cdb81610fdf565b840191505092915050565b60006040820190508181036000830152611d008185611cad565b9050611d0f60208301846111be565b9392505050565b6000606082019050611d2b60008301866111be565b8181036020830152611d3d81856117f1565b90508181036040830152611d5181846117f1565b9050949350505050565b60006020820190508181036000830152611d758184611998565b905092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b6000611e1582610fa6565b9150611e2083610fa6565b925082611e3057611e2f611dac565b5b600160000383147f800000000000000000000000000000000000000000000000000000000000000083141615611e6957611e68611ddb565b5b828205905092915050565b6000611e7f82610fa6565b9150611e8a83610fa6565b9250817f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff03831360008312151615611ec557611ec4611ddb565b5b817f8000000000000000000000000000000000000000000000000000000000000000038312600083121615611efd57611efc611ddb565b5b828201905092915050565b6000608082019050611f1d60008301876111be565b611f2a6020830186610fb0565b611f376040830185610fb0565b611f4460608301846111be565b95945050505050565b600060208284031215611f6357611f62610f39565b5b600082015167ffffffffffffffff811115611f8157611f80610f3e565b5b611f8d84828501611b32565b91505092915050565b60008060408385031215611fad57611fac610f39565b5b600083015167ffffffffffffffff811115611fcb57611fca610f3e565b5b611fd785828601611be7565b925050602083015167ffffffffffffffff811115611ff857611ff7610f3e565b5b61200485828601611b32565b915050925092905056fea264697066735822122010921d1203d54be0bfdf8da4324830541a73015559948f575d4c8c72f6553dcc64736f6c634300080f0033",
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

// Predicted is a free data retrieval call binding the contract method 0xb7aec839.
//
// Solidity: function predicted() view returns(int256)
func (_Main *MainCaller) Predicted(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "predicted")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Predicted is a free data retrieval call binding the contract method 0xb7aec839.
//
// Solidity: function predicted() view returns(int256)
func (_Main *MainSession) Predicted() (*big.Int, error) {
	return _Main.Contract.Predicted(&_Main.CallOpts)
}

// Predicted is a free data retrieval call binding the contract method 0xb7aec839.
//
// Solidity: function predicted() view returns(int256)
func (_Main *MainCallerSession) Predicted() (*big.Int, error) {
	return _Main.Contract.Predicted(&_Main.CallOpts)
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

// TestPrediction is a paid mutator transaction binding the contract method 0xd0148a67.
//
// Solidity: function testPrediction(int256[][] dependentVars) returns()
func (_Main *MainTransactor) TestPrediction(opts *bind.TransactOpts, dependentVars [][]*big.Int) (*types.Transaction, error) {
	return _Main.contract.Transact(opts, "testPrediction", dependentVars)
}

// TestPrediction is a paid mutator transaction binding the contract method 0xd0148a67.
//
// Solidity: function testPrediction(int256[][] dependentVars) returns()
func (_Main *MainSession) TestPrediction(dependentVars [][]*big.Int) (*types.Transaction, error) {
	return _Main.Contract.TestPrediction(&_Main.TransactOpts, dependentVars)
}

// TestPrediction is a paid mutator transaction binding the contract method 0xd0148a67.
//
// Solidity: function testPrediction(int256[][] dependentVars) returns()
func (_Main *MainTransactorSession) TestPrediction(dependentVars [][]*big.Int) (*types.Transaction, error) {
	return _Main.Contract.TestPrediction(&_Main.TransactOpts, dependentVars)
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

