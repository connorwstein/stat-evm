package precompile

import (
	"crypto/md5"
	"encoding/binary"
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"math"
	"math/big"
	"net/http"
	"strconv"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"golang.org/x/exp/rand"
	"gonum.org/v1/gonum/stat"
	"gonum.org/v1/gonum/stat/distuv"
)

type ContractIPFSPredictPriceConfig struct {
	BlockTimestamp *big.Int `json:"blockTimestamp"`
}

func (c ContractIPFSPredictPriceConfig) Address() common.Address {
	return ContractIPFSPredictPriceAddress
}

func (c ContractIPFSPredictPriceConfig) Timestamp() *big.Int {
	return c.BlockTimestamp
}

func (c ContractIPFSPredictPriceConfig) IsDisabled() bool {
	return false
}

func (c ContractIPFSPredictPriceConfig) Equal(config StatefulPrecompileConfig) bool {
	return false
}

func (c ContractIPFSPredictPriceConfig) Configure(config ChainConfig, db StateDB, context BlockContext) {
}

func (c ContractIPFSPredictPriceConfig) Contract() StatefulPrecompiledContract {
	return ContractIPFSPredictPricePrecompile
}

var (
	ContractIPFSPredictPricePrecompile StatefulPrecompiledContract = createIPFSPredictPricePrecompile(ContractIPFSPredictPriceAddress)
	IPFSPredictPriceSignature                                      = crypto.Keccak256([]byte("getIPFSPredictPrice(string,uint256,uint256)"))[:4]
)

func mustIPFSPredictPriceType(ts string) abi.Type {
	ty, _ := abi.NewType(ts, "", nil)
	return ty
}

func MakeIPFSPredictPriceArgs() abi.Arguments {
	return abi.Arguments{
		{
			Name: "v1",
			Type: mustIPFSPredictPriceType("string"),
		},
		{
			Name: "v2",
			Type: mustIPFSPredictPriceType("uint256"),
		},
		{
			Name: "v3",
			Type: mustIPFSPredictPriceType("uint256"),
		},
	}
}

func MakeIPFSPredictPriceRetArgs() abi.Arguments {
	return abi.Arguments{
		{
			Name: "ret",
			Type: mustIPFSPredictPriceType("uint256"),
		},
	}
}

func predictPriceIPFS(
	evm PrecompileAccessibleState,
	callerAddr common.Address,
	addr common.Address,
	input []byte,
	suppliedGas uint64,
	readOnly bool,
) (ret []byte, remainingGas uint64, err error) {
	inputCopy := make([]byte, len(input))
	copy(inputCopy, input)
	vals, err := MakeIPFSPredictPriceArgs().UnpackValues(inputCopy)
	if err != nil {
		return nil, suppliedGas, err
	}
	if len(vals) != 3 {
		return nil, suppliedGas, errors.New("Invalid number of args")
	}
	ipfsHash, ok := vals[0].(string)
	fmt.Println(ipfsHash)
	if !ok {
		return nil, suppliedGas, errors.New("invalid val")
	}
	v1, ok := vals[1].(*big.Int)
	if !ok {
		return nil, suppliedGas, errors.New("invalid val")
	}
	v2, ok := vals[2].(*big.Int)
	if !ok {
		return nil, suppliedGas, errors.New("invalid val")
	}
	if err != nil {
		fmt.Println(err)
	}

	data, err := http.Get(fmt.Sprintf("http://127.0.0.1:8080/ipfs/%s", ipfsHash))
	if err != nil {
		// TODO: return custom error to drop tx and maybe do some custom "failed data query billing from the sender"?
		return nil, remainingGas, err
	}

	r := csv.NewReader(data.Body)
	rows, err := r.ReadAll()
	if err != nil {
		return nil, suppliedGas, err
	}
	if err := data.Body.Close(); err != nil {
		return nil, suppliedGas, err
	}
	var price_history []float64
	for _, i := range rows {
		if _, err := strconv.ParseFloat(i[1], 64); err != nil {
			continue
		}
		floatNum, _ := strconv.ParseFloat(i[1], 32)
		price_history = append(price_history, floatNum)
	}
	// price_history = append(price_history, 33)
	// price_history = append(price_history, 35)
	// price_history = append(price_history, 32)
	// price_history = append(price_history, 35)
	// price_history = append(price_history, 32)
	// price_history = append(price_history, 32)

	a1 := price_history[1:]

	a2 := price_history[:len(price_history)-1]

	var ret_val []float64
	for i := 0; i < len(a1); i++ {
		ret_val = append(ret_val, a1[i]/a2[i]-1)
	}
	mu := stat.Mean(ret_val, nil)

	sigma := stat.StdDev(ret_val, nil)

	spot := price_history[len(price_history)-1]

	// Get random seed from hash of caller address, receiver address, and block timestamp
	h := md5.New()
	io.WriteString(h, callerAddr.String()+addr.String())
	randomSeed := rand.NewSource(binary.BigEndian.Uint64(h.Sum(nil)))

	dist := distuv.Normal{
		Mu:    mu,    // Mean of the normal distribution
		Sigma: sigma, // Standard deviation of the normal distribution
		Src:   randomSeed,
	}

	nTimeStep := v1.Int64()
	nSamples := v2.Int64()
	targetTime := nTimeStep
	var samples = make([]float64, nSamples+1)
	var pricePredictions = make([]float64, nTimeStep+1)
	for x := range samples {
		for i := range pricePredictions {
			pricePredictions[i] = 0
		}
		for i, _ := range pricePredictions {
			random := dist.Rand()
			if i == 0 {
				pricePredictions[i] = spot
			} else {
				res := math.Sqrt(1) * sigma * random
				val := mu + res + math.Log(pricePredictions[i-1])
				pricePredictions[i] = math.Exp(val)
			}
		}
		samples[x] = pricePredictions[targetTime]
		fmt.Println("samples", samples)
	}

	prediction := stat.Mean(samples, nil)
	bigval := new(big.Float).SetFloat64(prediction)
	intPrediction, _ := bigval.Int64()
	result := new(big.Int).SetInt64(intPrediction)
	fmt.Println("predict", prediction)
	ret, err = MakeIPFSPredictPriceRetArgs().PackValues([]interface{}{result})
	if err != nil {
		return nil, suppliedGas, err
	}
	return ret, suppliedGas, nil
}

func createIPFSPredictPricePrecompile(precompileAddr common.Address) StatefulPrecompiledContract {
	// Return new contract with no fallback function.
	return newStatefulPrecompileWithFunctionSelectors(nil, []*statefulPrecompileFunction{newStatefulPrecompileFunction(IPFSPredictPriceSignature, predictPriceIPFS)})
}
