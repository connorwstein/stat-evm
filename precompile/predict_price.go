package precompile

import (
	"encoding/csv"
	"errors"
	"fmt"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
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
	predictPriceSignature                                      = crypto.Keccak256([]byte("getPredictPrice(string)"))[:4]
)

func mustPredictPriceType(ts string) abi.Type {
	ty, _ := abi.NewType(ts, "", nil)
	return ty
}

func MakePredictPriceArgs() abi.Arguments {
	return abi.Arguments{
		{
			Name: "v1",
			Type: mustPredictPriceType("string"),
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

type historicalData struct {
	Date  string
	Price string
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
	csvFile, err := os.Open("./ETHUSD.csv")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened CSV file")
	defer csvFile.Close()
	fmt.Println(v1)

	csvLines, err := csv.NewReader(csvFile).ReadAll()
	if err != nil {
		fmt.Println(err)
	}
	for _, line := range csvLines {
		data := historicalData{
			Date:  line[0],
			Price: line[1],
		}
		fmt.Println(data.Date + " " + data.Price)
	}

	return ret, suppliedGas, nil
}

func createPredictPricePrecompile(precompileAddr common.Address) StatefulPrecompiledContract {
	// Return new contract with no fallback function.
	return newStatefulPrecompileWithFunctionSelectors(nil, []*statefulPrecompileFunction{newStatefulPrecompileFunction(predictPriceSignature, predictPrice)})
}
