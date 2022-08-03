package precompile

import (
	"errors"
	"fmt"
	"math"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
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
	_ StatefulPrecompileConfig = &ContractXChainECRecoverConfig{}
	// Singleton StatefulPrecompiledContract for XChain ECRecover.
	ContractSamplerPrecompile StatefulPrecompiledContract = createMedianPrecompile(ContractSamplerAddress)
	samplerSignature                                      = CalculateFunctionSelector("getSample(uint256,uint256)")
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
	//cocos,taste of thai express,sangcan indian
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
		addr common.Address, input []byte, suppliedGas uint64, readOnly bool) (ret []byte, remainingGas uint64, err error) {
		
		fmt.Println("input", hexutil.Encode(input))
		inputCopy := make([]byte, len(input))
		copy(inputCopy, input)

		var errb [32]byte

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

		// valsI := []*big.Int{v1, v2}
		dist := distuv.Normal{
			Mu:    float64(v1.Int64()), // Mean of the normal distribution
			Sigma: float64(v2.Int64()), // Standard deviation of the normal distribution
		}
		sample := dist.Rand()
		res := new(big.Float)
		bigval := new(big.Float).SetFloat64(sample)
		res.SetInt(big.NewInt(int64(math.Pow(10, 256))))
		bigval.Mul(bigval, res)
		result := new(big.Int)
		// f, _ := bigval.Uint64()
		result.SetUint64(10) // f
		ret, err = MakeSamplerRetArgs().PackValues([]interface{}{result})
		if err != nil {
			return errb[:], suppliedGas, err
		}
		return ret, suppliedGas, nil
	}

	funcGetSampler := newStatefulPrecompileFunction(samplerSignature, f)
	// Construct the contract with no fallback function.
	contract := newStatefulPrecompileWithFunctionSelectors(nil, []*statefulPrecompileFunction{funcGetSampler})
	return contract
}
