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
	linearmodel "github.com/pa-m/sklearn/linear_model"
)

type ContractIPFSFitConfig struct {
	BlockTimestamp *big.Int `json:"blockTimestamp"`
}

func (c ContractIPFSFitConfig) Address() common.Address {
	return ContractIPFSFitAddress
}

func (c ContractIPFSFitConfig) Timestamp() *big.Int {
	return c.BlockTimestamp
}

func (c ContractIPFSFitConfig) IsDisabled() bool {
	return false
}

func (c ContractIPFSFitConfig) Equal(config StatefulPrecompileConfig) bool {
	return false
}

func (c ContractIPFSFitConfig) Configure(config ChainConfig, db StateDB, context BlockContext) {
}

func (c ContractIPFSFitConfig) Contract() StatefulPrecompiledContract {
	return ContractIPFSFitPrecompile
}

var (
	ContractIPFSFitPrecompile StatefulPrecompiledContract = createIPFSFitPrecompile(ContractIPFSFitAddress)
	ipfsFitSignature                                      = crypto.Keccak256([]byte("getFit(string,uint256)"))[:4]
)

func MakeIPFSFitArgs() abi.Arguments {
	return abi.Arguments{
		{
			Name: "ipfsHash",
			Type: mustType("string"),
		},
		{
			Name: "fitType",
			Type: mustType("uint256"),
		},
	}
}

func MakeIPFSFitRetArgs() abi.Arguments {
	return abi.Arguments{
		{
			Name: "coeff",
			Type: mustType("int256[][]"),
		},
		{
			Name: "intercept",
			Type: mustType("int256[]"),
		},
	}
}

func parse(rows [][]string) ([][]float64, [][]float64, error) {
	var X [][]float64                      // row major independent
	var Y [][]float64                      // row major dependent
	for row := 1; row < len(rows); row++ { // Assume always first row is headings?
		var rowVals []float64
		skipRow := false
		for col := 0; col < (len(rows[row]) - 1); col++ { // all but last column
			if rows[row][col] == "" {
				// Throw away row if empty value
				skipRow = true
				break
			}
			f, err := strconv.ParseFloat(rows[row][col], 64)
			if err != nil {
				return nil, nil, err
			}
			rowVals = append(rowVals, f)
		}
		if skipRow {
			continue
		}
		X = append(X, rowVals)

		// Parse dependent val (last column)
		f, err := strconv.ParseFloat(rows[row][len(rows[row])-1], 64)
		if err != nil {
			return nil, nil, err
		}
		Y = append(Y, []float64{f})
	}
	return X, Y, nil
	//matX := buildMatrixFromFloat(X)
	//matY := buildMatrixFromFloat(Y)
	//r1, c1 := matX.Dims()
	//r2, c2 := matY.Dims()
}

func getIPFSFit(
	evm PrecompileAccessibleState,
	callerAddr common.Address,
	addr common.Address,
	input []byte,
	suppliedGas uint64,
	readOnly bool,
) (ret []byte, remainingGas uint64, err error) {
	inputCopy := make([]byte, len(input))
	copy(inputCopy, input)

	valsI, err := MakeIPFSFitArgs().UnpackValues(inputCopy)
	if err != nil {
		return nil, suppliedGas, err
	}

	ipfsHash, ok := valsI[0].(string)
	if !ok {
		return nil, suppliedGas, vm.ErrExecutionReverted
	}
	data, err := http.Get(fmt.Sprintf("http://127.0.0.1:8080/ipfs/%s", ipfsHash))
	if err != nil {
		// TODO: return custom error to drop tx and maybe do some custom "failed data query billing from the sender"?
		return nil, suppliedGas, err
	}
	r := csv.NewReader(data.Body)
	rows, err := r.ReadAll()
	if err != nil {
		return nil, suppliedGas, err
	}
	if err := data.Body.Close(); err != nil {
		return nil, suppliedGas, err
	}
	X, Y, err := parse(rows)
	if err != nil {
		return nil, suppliedGas, err
	}
	matX := buildMatrixFromFloat(X)
	matY := buildMatrixFromFloat(Y)
	// Compress into matrix
	regr := linearmodel.NewLinearRegression()
	regr.Fit(matX, matY)
	fmt.Println("receipt")
	coeff := *regr.Coef
	fmt.Printf("intercept %f\n", regr.Intercept.RawRowView(0)[0])
	rows2, cols2 := coeff.Dims()
	for row := 0; row < rows2; row++ {
		for col := 0; col < cols2; col++ {
			coeff.Set(row, col, coeff.At(row, col)*1_000_000)
		}
	}
	nr, _ := coeff.Dims()
	var res [][]*big.Int
	for row := 0; row < nr; row++ {
		rowVals := coeff.RawRowView(row)
		solRow := make([]*big.Int, len(rowVals))
		for i, rowVal := range rowVals {
			rowEntry, _ := big.NewFloat(rowVal).Int64()
			solRow[i] = big.NewInt(rowEntry)
		}
		res = append(res, solRow)
	}
	// TODO: maybe generalize to
	rowEntry, _ := big.NewFloat(regr.Intercept.RawRowView(0)[0] * 1_000_000).Int64() // Scale up instead of truncate?
	ret, err = MakeIPFSFitRetArgs().PackValues([]interface{}{res, []*big.Int{big.NewInt(rowEntry)}})
	if err != nil {
		return nil, suppliedGas, err
	}

	return ret, suppliedGas, nil
}

func createIPFSFitPrecompile(precompileAddr common.Address) StatefulPrecompiledContract {
	// Construct the contract with no fallback function.
	contract := newStatefulPrecompileWithFunctionSelectors(nil, []*statefulPrecompileFunction{newStatefulPrecompileFunction(ipfsFitSignature, getIPFSFit)})
	return contract
}
