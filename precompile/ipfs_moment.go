package precompile

import (
	"encoding/csv"
	"fmt"
	"math/big"
	"net/http"
	"strconv"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/vm"
	"github.com/ethereum/go-ethereum/crypto"
)

type ContractIPFSMomentConfig struct {
	BlockTimestamp *big.Int `json:"blockTimestamp"`
}

func (c ContractIPFSMomentConfig) Address() common.Address {
	return ContractIPFSMomentAddress
}

func (c ContractIPFSMomentConfig) Timestamp() *big.Int {
	return c.BlockTimestamp
}

func (c ContractIPFSMomentConfig) IsDisabled() bool {
	return false
}

func (c ContractIPFSMomentConfig) Equal(config StatefulPrecompileConfig) bool {
	return false
}

func (c ContractIPFSMomentConfig) Configure(config ChainConfig, db StateDB, context BlockContext) {
}

func (c ContractIPFSMomentConfig) Contract() StatefulPrecompiledContract {
	return ContractIPFSMomentPrecompile
}

var (
	ContractIPFSMomentPrecompile StatefulPrecompiledContract = createIPFSMomentPrecompile(ContractIPFSMomentAddress)
	ipfsMomentSignature                                      = crypto.Keccak256([]byte("getMoment(string,uint256)"))[:4]
)

func MakeIPFSMomentArgs() abi.Arguments {
	return abi.Arguments{
		{
			Name: "ipfsHash",
			Type: mustType("string"),
		},
		{
			Name: "moment",
			Type: mustType("uint256"),
		},
	}
}

func MakeIPFSMomentRetArgs() abi.Arguments {
	return abi.Arguments{
		{
			Name: "ret",
			Type: mustType("uint256"),
		},
	}
}

func getIPFSMoment(
	evm PrecompileAccessibleState,
	callerAddr common.Address,
	addr common.Address,
	input []byte,
	suppliedGas uint64,
	readOnly bool,
) (ret []byte, remainingGas uint64, err error) {
	inputCopy := make([]byte, len(input))
	copy(inputCopy, input)

	valsI, err := MakeIPFSMomentArgs().UnpackValues(inputCopy)
	if err != nil {
		return nil, suppliedGas, err
	}
	ipfsHash, ok := valsI[0].(string)
	if !ok {
		return nil, suppliedGas, vm.ErrExecutionReverted
	}
	moment, ok := valsI[1].(*big.Int)
	if !ok {
		return nil, suppliedGas, vm.ErrExecutionReverted
	}
	data, err := http.Get(fmt.Sprintf("http://127.0.0.1:8080/ipfs/%s", ipfsHash))
	if err != nil {
		// TODO: return custom error to drop tx and maybe do some custom "failed data query billing from the sender"?
		return nil, remainingGas, err
	}
	r := csv.NewReader(data.Body)
	rows, err := r.ReadAll()
	var rawMoment int64
	for row := 1; row < len(rows); row++ { // Assume always first row is headings?
		for col := 0; col < len(rows[row]); col++ {
			i, err := strconv.Atoi(rows[row][col])
			if err != nil {
				// TODO: return custom error to drop tx and maybe do some custom "failed data query billing from the sender"?
				return nil, remainingGas, err
			}
			rawMoment += big.NewInt(0).Exp(big.NewInt(int64(i)), moment, nil).Int64()
		}
	}
	ret, err = MakeRetArgs().PackValues([]interface{}{big.NewInt(rawMoment)})
	if err != nil {
		return nil, suppliedGas, err
	}

	return ret, suppliedGas, nil
}

func createIPFSMomentPrecompile(precompileAddr common.Address) StatefulPrecompiledContract {
	// Construct the contract with no fallback function.
	contract := newStatefulPrecompileWithFunctionSelectors(nil, []*statefulPrecompileFunction{newStatefulPrecompileFunction(ipfsMomentSignature, getIPFSMoment)})
	return contract
}
