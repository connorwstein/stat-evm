package precompile

import (
	"context"
	"errors"
	"fmt"
	"log"
	"math"
	"math/big"
	"os"
	"strconv"

	"github.com/rocketlaunchr/dataframe-go"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/rocketlaunchr/dataframe-go/imports"
	"gonum.org/v1/gonum/stat"
	"gonum.org/v1/gonum/stat/distuv"
)

type ContractPredictPriceConfig struct {
	BlockTimestamp *big.Int `json:"blockTimestamp"`
}

func (c ContractPredictPriceConfig) Address() common.Address {
	return ContractPredictPriceAddress
}

func (c ContractPredictPriceConfig) Timestamp() *big.Int {
	return c.BlockTimestamp
}

func (c ContractPredictPriceConfig) IsDisabled() bool {
	return false
}

func (c ContractPredictPriceConfig) Equal(config StatefulPrecompileConfig) bool {
	return false
}

func (c ContractPredictPriceConfig) Configure(config ChainConfig, db StateDB, context BlockContext) {
}

func (c ContractPredictPriceConfig) Contract() StatefulPrecompiledContract {
	return ContractPredictPricePrecompile
}

var (
	ContractPredictPricePrecompile StatefulPrecompiledContract = createPredictPricePrecompile(ContractPredictPriceAddress)
	predictPriceSignature                                      = crypto.Keccak256([]byte("getPredictPrice(uint256)"))[:4]
)

func mustPredictPriceType(ts string) abi.Type {
	ty, _ := abi.NewType(ts, "", nil)
	return ty
}

func MakePredictPriceArgs() abi.Arguments {
	return abi.Arguments{
		{
			Name: "numOfSamples",
			Type: mustType("uint256"),
		},
		{
			Name: "stepToPredict",
			Type: mustType("uint256"),
		},
	}
}

func MakePredictPriceRetArgs() abi.Arguments {
	return abi.Arguments{
		{
			Name: "predictedPrice",
			Type: mustPredictPriceType("uint256"),
		},
	}
}

type SeriesFloat64 struct {
	values []float64
}

func (s *SeriesFloat64) copy() *SeriesFloat64 {
	if len(s.values) == 0 {
		return &SeriesFloat64{
			values: []float64{},
		}
	}
	// Copy slice
	x := s.values[0:len(s.values)]
	newSlice := append(x[:0:0], x...)
	return &SeriesFloat64{
		values: newSlice,
	}
}
func predictPrice(
	evm PrecompileAccessibleState,
	callerAddr common.Address,
	addr common.Address,
	input []byte,
	suppliedGas uint64,
	readOnly bool,
) (ret []byte, remainingGas uint64, err error) {
	inputCopy := make([]byte, len(input))
	copy(inputCopy, input)
	vals, err := MakePredictPriceArgs().UnpackValues(inputCopy)
	if err != nil {
		return nil, suppliedGas, err
	}
	if len(vals) != 2 {
		return nil, suppliedGas, errors.New("Invalid number of args")
	}

	v1, ok := vals[0].(*big.Int)
	if !ok {
		return nil, suppliedGas, errors.New("invalid val")
	}
	v2, ok := vals[0].(*big.Int)
	if !ok {
		return nil, suppliedGas, errors.New("invalid val")
	}
	if err != nil {
		fmt.Println(err)
	}
	ctx := context.TODO()
	csvfile, err := os.Open("./ETHUSD.csv")
	if err != nil {
		log.Fatal(err)
	}
	df, err := imports.LoadFromCSV(ctx, csvfile)

	iterator := df.ValuesIterator(dataframe.ValuesOptions{0, 1, true})
	df.Lock()
	var output []float64
	var last_not_null_val float64

	for {
		row, vals, _ := iterator()
		if row == nil {
			break
		}
		if vals[4] == "null" {
			output = append(output, last_not_null_val)
			continue
		}

		f := vals[4].(string)
		floatNum, _ := strconv.ParseFloat(f, 32)
		last_not_null_val = floatNum
		output = append(output, floatNum)
	}
	df.Unlock()

	a1 := output[1:]

	a2 := output[:len(output)-1]
	var ret_val []float64
	for i := 0; i < len(a1); i++ {
		ret_val = append(ret_val, a1[i]/a2[i]-1)
	}

	mu := stat.Mean(ret_val, nil)
	fmt.Println("mu", mu)

	sigma := stat.StdDev(ret_val, nil)
	fmt.Println("sigma", sigma)

	spot := output[len(output)-1]
	fmt.Println("spot", spot)

	nTimeStep := v1.Int64()
	nSamples := v2.Int64()
	var outerPricePredictions = make([]float64, nSamples+1)
	var pricePredictions = make([]float64, nTimeStep+1)
	for x := range pricePredictions {
		for i := range pricePredictions {
			pricePredictions[i] = 0
		}
		for i, _ := range pricePredictions {
			dist := distuv.Normal{
				Mu:    mu,    // Mean of the normal distribution
				Sigma: sigma, // Standard deviation of the normal distribution
			}
			random := dist.Rand()
			if i == 0 {
				pricePredictions[i] = spot
			} else {
				res := math.Sqrt(1) * sigma * random
				val := mu + res + math.Log(pricePredictions[i-1])
				fmt.Println("2 calc val is", val)
				pricePredictions[i] = math.Exp(val)
			}
		}
		outerPricePredictions[x] = stat.Mean(pricePredictions, nil)
		fmt.Println("predicted price path is", pricePredictions)

	}
	fmt.Println("outerPricePredictions", stat.Mean(outerPricePredictions, nil))

	return ret, suppliedGas, nil
}

func createPredictPricePrecompile(precompileAddr common.Address) StatefulPrecompiledContract {
	// Return new contract with no fallback function.
	return newStatefulPrecompileWithFunctionSelectors(nil, []*statefulPrecompileFunction{newStatefulPrecompileFunction(predictPriceSignature, predictPrice)})
}
