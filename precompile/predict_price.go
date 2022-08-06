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
		fmt.Println(vals[0], vals[4])
		f := vals[4].(string)
		floatNum, _ := strconv.ParseFloat(f, 32)
		output = append(output, floatNum)
	}
	df.Unlock()
	sigma := stat.Mean(output, nil)
	fmt.Println("Mu", sigma)
	variance := stat.Variance(output, nil)
	mu := math.Sqrt(variance)
	fmt.Println("Sigma", mu)
	spot := output[len(output)-1]
	fmt.Println("Spot is", spot)
	dist := distuv.Normal{
		Mu:    mu,    // Mean of the normal distribution
		Sigma: sigma, // Standard deviation of the normal distribution
	}
	random := dist.Rand()

	nTimeStep := 10
	counter := 0
	var pricePredictions []float64
	pricePredictions = append(pricePredictions, math.Log(spot))
	for counter < nTimeStep {
		res := math.Sqrt(1) * sigma * random
		val := mu + res
		pricePredictions = append(pricePredictions, val)
		counter++
	}
	fmt.Println("Got res num", pricePredictions)
	return ret, suppliedGas, nil
}

func createPredictPricePrecompile(precompileAddr common.Address) StatefulPrecompiledContract {
	// Return new contract with no fallback function.
	return newStatefulPrecompileWithFunctionSelectors(nil, []*statefulPrecompileFunction{newStatefulPrecompileFunction(predictPriceSignature, predictPrice)})
}
