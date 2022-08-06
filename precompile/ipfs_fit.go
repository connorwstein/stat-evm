package precompile

import (
	"encoding/csv"
	"fmt"
	"math/big"
	"net/http"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/vm"
	"github.com/ethereum/go-ethereum/crypto"
)

type ContractIPFSFitConfig struct {
	BlockTimestamp *big.Int `json:"blockTimestamp"`
}

func (c ContractIPFSFitConfig) Address() common.Address {
	return ContractIPFSFitAddress
}

func (c ContractIPFSFitConfig) Timestamp() *big.Int {
	return c.BlockTimestamp
}

func (c ContractIPFSFitConfig) IsDisabled() bool {
	return false
}

func (c ContractIPFSFitConfig) Equal(config StatefulPrecompileConfig) bool {
	return false
}

func (c ContractIPFSFitConfig) Configure(config ChainConfig, db StateDB, context BlockContext) {
}

func (c ContractIPFSFitConfig) Contract() StatefulPrecompiledContract {
	return ContractIPFSFitPrecompile
}

var (
	ContractIPFSFitPrecompile StatefulPrecompiledContract = createIPFSFitPrecompile(ContractIPFSFitAddress)
	ipfsFitSignature                                      = crypto.Keccak256([]byte("getFit(string,uint256)"))[:4]
)

func MakeIPFSFitArgs() abi.Arguments {
	return abi.Arguments{
		{
			Name: "ipfsHash",
			Type: mustType("string"),
		},
		{
			Name: "fitType",
			Type: mustType("uint256"),
		},
	}
}

func MakeIPFSFitRetArgs() abi.Arguments {
	return abi.Arguments{
		{
			Name: "ret",
			Type: mustType("uint256"),
		},
	}
}

func getIPFSFit(
	evm PrecompileAccessibleState,
	callerAddr common.Address,
	addr common.Address,
	input []byte,
	suppliedGas uint64,
	readOnly bool,
) (ret []byte, remainingGas uint64, err error) {
	inputCopy := make([]byte, len(input))
	copy(inputCopy, input)
	valsI, err := MakeIPFSFitArgs().UnpackValues(inputCopy)
	if err != nil {
		return nil, suppliedGas, err
	}
	ipfsHash, ok := valsI[0].(string)
	if !ok {
		return nil, suppliedGas, vm.ErrExecutionReverted
	}
	data, err := http.Get(fmt.Sprintf("http://127.0.0.1:8080/ipfs/%s", ipfsHash))
	if err != nil {
		// TODO: return custom error to drop tx and maybe do some custom "failed data query billing from the sender"?
		return nil, remainingGas, err
	}
	r := csv.NewReader(data.Body)
	//rows, err := r.ReadAll()
	fmt.Println(r)
	ret, err = MakeRetArgs().PackValues([]interface{}{big.NewInt(int64(10))})
	if err != nil {
		return nil, suppliedGas, err
	}

	return ret, suppliedGas, nil
}

func createIPFSFitPrecompile(precompileAddr common.Address) StatefulPrecompiledContract {
	// Construct the contract with no fallback function.
	contract := newStatefulPrecompileWithFunctionSelectors(nil, []*statefulPrecompileFunction{newStatefulPrecompileFunction(ipfsFitSignature, getIPFSFit)})
	return contract
}
