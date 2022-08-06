package precompile

import (
	"encoding/csv"
	"errors"
	"fmt"
	"log"
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

type historicalData struct {
	Date  string
	Price string
}

func readCsvFile(filePath string) [][]string {
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Unable to read input file "+filePath, err)
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal("Unable to parse file as CSV for "+filePath, err)
	}

	return records
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
	csvFile := readCsvFile("./ETHUSD.csv")

	fmt.Println(v1)
	fmt.Println(csvFile)
	return ret, suppliedGas, nil
}

func createPredictPricePrecompile(precompileAddr common.Address) StatefulPrecompiledContract {
	// Return new contract with no fallback function.
	return newStatefulPrecompileWithFunctionSelectors(nil, []*statefulPrecompileFunction{newStatefulPrecompileFunction(predictPriceSignature, predictPrice)})
}
