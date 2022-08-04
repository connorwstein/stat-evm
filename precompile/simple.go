package precompile

import (
	"errors"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
)

type ContractSimpleConfig struct {
	BlockTimestamp *big.Int `json:"blockTimestamp"`
}

func (c ContractSimpleConfig) Address() common.Address {
	return ContractMedianAddress
}

func (c ContractSimpleConfig) Timestamp() *big.Int {
	return c.BlockTimestamp
}

func (c ContractSimpleConfig) IsDisabled() bool {
	return false
}

func (c ContractSimpleConfig) Equal(config StatefulPrecompileConfig) bool {
	return false
}

func (c ContractSimpleConfig) Configure(config ChainConfig, db StateDB, context BlockContext) {
}

func (c ContractSimpleConfig) Contract() StatefulPrecompiledContract {
	return ContractMedianPrecompile
}

var (
	ContractSimplePrecompile StatefulPrecompiledContract = createSimplePrecompile(ContractSimpleAddress)
	simpleSignature = CalculateFunctionSelector("getSimple(uint256)")
)

func mustSimpleType(ts string) abi.Type {
	ty, _ := abi.NewType(ts, "", nil)
	return ty
}

func MakeSimpleArgs() abi.Arguments {
	return abi.Arguments{
		{
			Name: "v1",
			Type: mustSimpleType("uint256"),
		},
	}
}

func MakeSimpleRetArgs() abi.Arguments {
	return abi.Arguments{
		{
			Name: "ret",
			Type: mustType("uint256"),
		},
	}
}

func createSimplePrecompile(precompileAddr common.Address) StatefulPrecompiledContract {
	f := func(
		evm PrecompileAccessibleState,
		callerAddr common.Address,
		addr common.Address,
		input []byte,
		suppliedGas uint64,
		readOnly bool,
	) (ret []byte, remainingGas uint64, err error) {
		inputCopy := make([]byte, len(input))
		copy(inputCopy, input)

		var errb [32]byte
		errb[31] = 0xaa

		vals, err := MakeArgs().UnpackValues(inputCopy)
		if err != nil {
			return errb[:], suppliedGas, err
		}
		if len(vals) != 1 {
			return errb[:], suppliedGas, errors.New("invalid vals")
		}

		v1, ok := vals[0].(*big.Int)
		if !ok {
			return errb[:], suppliedGas, errors.New("invalid val")
		}

		ret, err = MakeSimpleRetArgs().PackValues([]interface{}{v1})
		if err != nil {
			return errb[:], suppliedGas, err
		}

		return ret, suppliedGas, nil
	}

	funcGetSimple := newStatefulPrecompileFunction(simpleSignature, f)
	
	// Return new contract with no fallback function.
	return newStatefulPrecompileWithFunctionSelectors(nil, []*statefulPrecompileFunction{funcGetSimple})
}
