package precompile

import (
	"errors"
	"math"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	linearmodel "github.com/pa-m/sklearn/linear_model"
	"gonum.org/v1/gonum/mat"
)

type ContractFitConfig struct {
	BlockTimestamp *big.Int `json:"blockTimestamp"`
}

func (c ContractFitConfig) Address() common.Address {
	return ContractFitAddress
}

func (c ContractFitConfig) Timestamp() *big.Int {
	return c.BlockTimestamp
}

func (c ContractFitConfig) IsDisabled() bool {
	return false
}

func (c ContractFitConfig) Equal(config StatefulPrecompileConfig) bool {
	return false
}

func (c ContractFitConfig) Configure(config ChainConfig, db StateDB, context BlockContext) {
}

func (c ContractFitConfig) Contract() StatefulPrecompiledContract {
	return ContractFitPrecompile
}

var (
	ContractFitPrecompile StatefulPrecompiledContract = createFitPrecompile(ContractFitAddress)
	fitSignature                                      = crypto.Keccak256([]byte("fit(uint256,int256[][],int256[][])"))[:4]
)

func MakeFitArgs() abi.Arguments {
	return abi.Arguments{
		{
			Name: "fitType",
			Type: mustType("uint256"),
		},
		{
			Name: "input",
			Type: mustType("int256[][]"),
		},
		{
			Name: "output",
			Type: mustType("int256[][]"),
		},
	}
}

func MakeFitReturnArgs() abi.Arguments {
	return abi.Arguments{
		{
			Name: "ret",
			Type: mustType("int256[][]"),
		},
	}
}

func specificFit(
	evm PrecompileAccessibleState,
	callerAddr common.Address,
	addr common.Address,
	input []byte,
	suppliedGas uint64,
	readOnly bool,
) (retb []byte, remainingGas uint64, err error) {
	inputCopy := make([]byte, len(input))
	copy(inputCopy, input)

	valsI, err := MakeFitArgs().UnpackValues(inputCopy)
	if err != nil {
		return nil, suppliedGas, err
	}

	vals, ok := valsI[0].([]*big.Int)
	if !ok {
		return nil, suppliedGas, errors.New("invalid type")
	}
	if len(vals)%2 != 0 {
		return nil, suppliedGas, errors.New("invalid parity of len(vals)")
	}

	xys3 := big.NewInt(0)
	for j := 0; j < len(vals); j += 2 {
		addend := big.NewInt(0)
		addend.Mul(vals[j], vals[j+1])
		xys3.Add(xys3, addend)
	}
	xys3.Mul(xys3, big.NewInt(3))
	xsys := big.NewInt(0)
	xs := big.NewInt(0)
	ys := big.NewInt(0)
	for j := 0; j < len(vals); j += 2 {
		xs.Add(xs, vals[j])
		ys.Add(ys, vals[j+1])
	}
	xsys.Mul(xs, ys)
	num := big.NewInt(0)
	num.Sub(xys3, xsys)
	denom := big.NewInt(0)
	x2s3 := big.NewInt(0)
	for j := 0; j < len(vals); j += 2 {
		addend := big.NewInt(0)
		addend.Mul(vals[j], vals[j])
		x2s3.Add(x2s3, addend)
	}
	x2s3.Mul(x2s3, big.NewInt(3))
	xs2 := big.NewInt(0)
	xs2.Mul(xs, xs)
	denom.Sub(x2s3, xs2)
	perc := big.NewInt(0)
	perc.Mul(num, big.NewInt(100))
	perc.Div(perc, denom)
	// fitb := 100 * (3*(v1*v2+v3*v4+v5*v6) - (v1+v3+v5)*(v2+v4+v6)) / (3*(v1*v1+v3*v3+v5*v5) - (v1+v3+v5)*(v1+v3+v5))
	// fita := 100*(v2+v4+v6)/3 - fitb*(v1+v3+v5)/3

	retb, err = MakeFitReturnArgs().PackValues([]interface{}{perc})
	if err != nil {
		return nil, suppliedGas, err
	}

	return retb, suppliedGas, nil
}

func fit(
	evm PrecompileAccessibleState,
	callerAddr common.Address,
	addr common.Address,
	input []byte,
	suppliedGas uint64,
	readOnly bool,
) (retb []byte, remainingGas uint64, err error) {
	// TODO: Figure out gas cost as a function of size of input matrices

	inputCopy := make([]byte, len(input))
	copy(inputCopy, input)

	valsI, err := MakeFitArgs().UnpackValues(inputCopy)
	if err != nil {
		return nil, suppliedGas, err
	}

	fitType, ok := valsI[0].(*big.Int)
	if !ok {
		return nil, suppliedGas, errors.New("invalid fit type")
	}

	X, ok := valsI[1].([][]*big.Int)
	matX := buildMatrix(X)
	if !ok {
		return nil, suppliedGas, errors.New("invalid input dataset type")
	}

	Y, ok := valsI[2].([][]*big.Int)
	matY := buildMatrix(Y)
	if !ok {
		return nil, suppliedGas, errors.New("invalid output dataset type")
	}

	numRowsY, _ := matY.Dims()
	numRowsX, _ := matX.Dims()
	if numRowsY != numRowsX {
		return nil, suppliedGas, errors.New("number of rows of X and Y must be equal")
	}

	var result mat.Dense

	switch uint(fitType.Uint64()) {
	case 0:
		// Ordinary Least Squares (OLS)
		regr := linearmodel.NewLinearRegression()
		regr.Fit(matX, matY)
		result = *regr.Coef
		// Round up to nearest 5 decimal places (picked out of hat)
		roundedDecimals := 5
		cols, rows := result.Dims()
		for row := 0; row < rows; row++ {
			for col := 0; col < cols; col++ {
				result.Set(row, col, math.Round(result.At(row, col)*math.Pow10(roundedDecimals))/math.Pow10(roundedDecimals))
			}
		}
	case 1:
		// Least Absolute Shrinkage and Selection Operator (LASSO)
		regr := linearmodel.NewRidge()
		regr.Alpha = 0.25
		regr.Tol = 1e-2
		regr.Normalize = false
		regr.Fit(matX, matY)
		result = *regr.Coef
	default:
		// Ordinary Least Squares (OLS)
		regr := linearmodel.NewLinearRegression()
		regr.Fit(matX, matY)
		result = *regr.Coef
		// Round up to nearest 5 decimal places (picked out of hat)
		roundedDecimals := 5
		cols, rows := result.Dims()
		for row := 0; row < rows; row++ {
			for col := 0; col < cols; col++ {
				result.Set(row, col, math.Round(result.At(row, col)*math.Pow10(roundedDecimals))/math.Pow10(roundedDecimals))
			}
		}
	}

	nr, _ := result.Dims()
	var res [][]*big.Int
	for row := 0; row < nr; row++ {
		rowVals := result.RawRowView(row)
		solRow := make([]*big.Int, len(rowVals))
		for i, rowVal := range rowVals {
			rowEntry, _ := big.NewFloat(rowVal).Int64()
			solRow[i] = big.NewInt(rowEntry)
		}
		res = append(res, solRow)
	}

	ret, err := MakeFitReturnArgs().PackValues([]interface{}{res})
	if err != nil {
		return nil, suppliedGas, err
	}

	return ret, suppliedGas, nil
}

func createFitPrecompile(precompileAddr common.Address) StatefulPrecompiledContract {
	// Return the contract with no fallback function.
	return newStatefulPrecompileWithFunctionSelectors(nil, []*statefulPrecompileFunction{newStatefulPrecompileFunction(fitSignature, fit)})
}
