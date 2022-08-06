package precompile

import (
	"fmt"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/require"
)

func TestFit(t *testing.T) {
	fitType := big.NewInt(0) // 0 = OLS, 1 = LASSO

	var x [][]*big.Int
	var y [][]*big.Int
	for i := 0; i < 100; i++ {
		val := big.NewInt(int64(i + 1))
		val_times_two := new(big.Int).Mul(val, big.NewInt(2))
		x = append(x, []*big.Int{val})
		y = append(y, []*big.Int{val_times_two})
	}

	input, err := MakeFitArgs().PackValues([]interface{}{fitType, x, y})
	require.NoError(t, err)

	ret, _, err := fit(nil, common.Address{}, common.Address{}, input, 0, false)
	require.NoError(t, err)
	fmt.Println(ret)

	vals, err := MakeFitReturnArgs().UnpackValues(ret)
	require.NoError(t, err)

	fmt.Println(vals[0].([][]*big.Int)) // Should return [[2]]
}
