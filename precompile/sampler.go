package precompile

import (
	"errors"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"gonum.org/v1/gonum/stat/distuv"
)

type ContractSamplerConfig struct {
	BlockTimestamp *big.Int `json:"blockTimestamp"`
}

func (c ContractSamplerConfig) Address() common.Address {
	return ContractSamplerAddress
}

func (c ContractSamplerConfig) Timestamp() *big.Int {
	return c.BlockTimestamp
}

func (c ContractSamplerConfig) IsDisabled() bool {
	return false
}

func (c ContractSamplerConfig) Equal(config StatefulPrecompileConfig) bool {
	return false
}

func (c ContractSamplerConfig) Configure(config ChainConfig, db StateDB, context BlockContext) {
}

func (c ContractSamplerConfig) Contract() StatefulPrecompiledContract {
	return ContractSamplerPrecompile
}

var (
	ContractSamplerPrecompile StatefulPrecompiledContract = createSamplerPrecompile(ContractSamplerAddress)
	samplerSignature = CalculateFunctionSelector("getSample(uint256,uint256)")
)

func mustSamplerType(ts string) abi.Type {
	ty, _ := abi.NewType(ts, "", nil)
	return ty
}

func MakeSamplerArgs() abi.Arguments {
	return abi.Arguments{
		{
			Name: "v1",
			Type: mustSamplerType("uint256"),
		},
		{
			Name: "v2",
			Type: mustSamplerType("uint256"),
		},
	}
}

func MakeSamplerRetArgs() abi.Arguments {
	return abi.Arguments{
		{
			Name: "ret",
			Type: mustType("uint256"),
		},
	}
}

func createSamplerPrecompile(precompileAddr common.Address) StatefulPrecompiledContract {
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
		// errb[31] = 0xaa

		vals, err := MakeSamplerArgs().UnpackValues(inputCopy)
		if err != nil {
			return errb[:], suppliedGas, err
		}
		if len(vals) != 2 {
			return errb[:], suppliedGas, errors.New("invalid vals")
		}

		v1, ok := vals[0].(*big.Int)
		if !ok {
			return errb[:], suppliedGas, errors.New("invalid val")
		}

		v2, ok := vals[1].(*big.Int)
		if !ok {
			return errb[:], suppliedGas, errors.New("invalid val")
		}

		dist := distuv.Normal{
			Mu:    float64(v1.Int64()), // Mean of the normal distribution
			Sigma: float64(v2.Int64()), // Standard deviation of the normal distribution
		}
		sample := dist.Rand()

		bigval := new(big.Float).SetFloat64(sample)
		// bigval = new(big.Float).SetFloat64(1924570923.124124) // Hardcoding a value works

		f, _ := bigval.Uint64()
		result := new(big.Int).SetUint64(f)
		ret, err = MakeSamplerRetArgs().PackValues([]interface{}{result})
		if err != nil {
			return errb[:], suppliedGas, err
		}

		return ret, suppliedGas, nil
	}

	funcGetSampler := newStatefulPrecompileFunction(samplerSignature, f)

	// Return new contract with no fallback function.
	return newStatefulPrecompileWithFunctionSelectors(nil, []*statefulPrecompileFunction{funcGetSampler})
}
