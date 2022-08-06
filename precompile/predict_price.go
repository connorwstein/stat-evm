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

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/rocketlaunchr/dataframe-go"
	"github.com/rocketlaunchr/dataframe-go/forecast/interpolation"
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
			Name: "v1",
			Type: mustPredictPriceType("uint256"),
		},
	}
}

func MakePredictPriceRetArgs() abi.Arguments {
	return abi.Arguments{
		{
			Name: "ret",
			Type: mustType("uint256"),
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
	if len(vals) != 1 {
		return nil, suppliedGas, errors.New("Invalid number of args")
	}

	v1, ok := vals[0].(*big.Int)
	if !ok {
		return nil, suppliedGas, errors.New("invalid val")
	}
	fmt.Println(v1)
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
	for {
		row, vals, _ := iterator()
		if row == nil {
			break
		}

		f := vals[4].(string)
		floatNum, _ := strconv.ParseFloat(f, 32)
		output = append(output, floatNum)
	}
	df.Unlock()
	opts := interpolation.InterpolateOptions{
		Method:        interpolation.ForwardFill{},
		FillDirection: interpolation.Forward,
		Limit:         &[]int{1}[0],
		FillRegion:    nil,
		InPlace:       true,
	}
	interpolated_data, err := interpolation.Interpolate(ctx, output, opts)
	fmt.Println("interpolated_data", interpolated_data)
	sigma := stat.Mean(output, nil)
	fmt.Println("Sigma", sigma)
	a1 := output[1:]

	//TODO INTERPOLATION
	// GET MU
	// GET SIGMA
	a2 := output[:len(output)-1]
	var ret_val []float64
	for i := 0; i < len(a1); i++ {
		ret_val = append(ret_val, a1[i]/a2[i]-1)
	}
	fmt.Println("ret_val", ret_val)
	mu := stat.Mean(ret_val, nil)
	fmt.Println("mu", mu)

	spot := output[len(output)-1]
	fmt.Println("Spot is", spot)

	nTimeStep := 10
	var pricePredictions []float64
	pricePredictions_ := make([]float64, nTimeStep)
	pricePredictions_ = append(pricePredictions, math.Log(spot))
	pricePredictions_[0] = math.Log(spot)
	for i, v := range pricePredictions_ {
		dist := distuv.Normal{
			Mu:    mu,    // Mean of the normal distribution
			Sigma: sigma, // Standard deviation of the normal distribution
		}
		random := dist.Rand()
		if i == 0 {
			prevPrice := pricePredictions_[i]
			fmt.Println("Prev val is", prevPrice)
			fmt.Println("rand val is", random)
			res := math.Sqrt(1) * sigma * random
			fmt.Println("calc mu is", mu)
			fmt.Println("calc sigma is", sigma)
			fmt.Println("calc res is", res)
			val := mu + res + prevPrice
			fmt.Println("calc val is", val)
			pricePredictions_ = append(pricePredictions_, val)
		} else {
			prevPrice := pricePredictions_[i-1]
			fmt.Println("Prev val is", prevPrice)
			res := math.Sqrt(1) * sigma * random
			val := mu + res + prevPrice
			pricePredictions_ = append(pricePredictions_, val)
		}
		fmt.Println("here", v, i)
	}
	fmt.Println(pricePredictions_)
	return ret, suppliedGas, nil
}

func createPredictPricePrecompile(precompileAddr common.Address) StatefulPrecompiledContract {
	// Return new contract with no fallback function.
	return newStatefulPrecompileWithFunctionSelectors(nil, []*statefulPrecompileFunction{newStatefulPrecompileFunction(predictPriceSignature, predictPrice)})
}
