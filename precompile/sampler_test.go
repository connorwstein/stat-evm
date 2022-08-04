package precompile

import (
	"fmt"
	"math/big"
	"testing"
	"reflect"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/stretchr/testify/require"
	"gonum.org/v1/gonum/stat/distuv"
)

func TestSampler(t *testing.T) {
	suppliedGas := 500_000

	// a, err := MakeSamplerArgs().PackValues([]interface{}{big.NewInt(0), big.NewInt(1)})
	a, err := MakeSamplerArgs().PackValues([]interface{}{big.NewInt(0), big.NewInt(1000000000000000000)})
	require.NoError(t, err)
	fmt.Println("Encoded packed values", hexutil.Encode(a[:]))
	vals, err := MakeSamplerArgs().UnpackValues(a)
	require.NoError(t, err)
	fmt.Println("unpacked values", vals)

	var errb [32]byte

	if err != nil {
		fmt.Println(errb[:], suppliedGas, err)
	}
	if len(vals) != 2 {
		fmt.Println(errb[:], suppliedGas)
	}

	v1, ok := vals[0].(*big.Int)
	if !ok {
		fmt.Println(errb[:], suppliedGas)
	}

	v2, ok := vals[1].(*big.Int)
	if !ok {
		fmt.Println(errb[:], suppliedGas)
	}

	fmt.Println("v1:", reflect.TypeOf(v1))
	fmt.Println("v2:", reflect.TypeOf(v2))
	
	dist := distuv.Normal{
		Mu:    float64(v1.Int64()), // Mean of the normal distribution
		Sigma: float64(v2.Int64()), // Standard deviation of the normal distribution
	}
	sample := dist.Rand()

	fmt.Println("sample:", sample)

	// res := new(big.Float).SetInt(big.NewInt(int64(math.Pow(10, 18))))
	bigval := new(big.Float).SetFloat64(sample)
	// bigval.Mul(bigval, res)
	fmt.Println("bigval:", bigval)

	f, _ := bigval.Uint64()
	result := new(big.Int).SetUint64(f)
	fmt.Println("result:", result)
	ret, err := MakeSamplerRetArgs().PackValues([]interface{}{result})
	if err != nil {
		fmt.Println(errb[:], suppliedGas, err)
	}
	
	fmt.Println(ret, suppliedGas)
}