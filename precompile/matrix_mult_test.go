package precompile

import (
	"fmt"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/require"
	"gotest.tools/assert"
)

func TestMatrixMult(t *testing.T) {
	numRows := 3

	a_row := make([]*big.Int, numRows)
	for j := 0; j < numRows; j++ { // columns
		a_row[j] = big.NewInt(int64(j + 1))
	}
	a := [][]*big.Int{a_row}

	b := make([][]*big.Int, numRows)
	for i := 0; i < numRows; i++ { // rows
		b[i] = []*big.Int{big.NewInt(int64(i + 1))}
	}
	fmt.Printf("a: %dx%d, b: %dx%d\n", len(a), len(a[0]), len(b), len(b[0]))

	input, err := MakeMatrixMultArgs().PackValues([]interface{}{a, b})
	require.NoError(t, err)

	ret, _, err := matrixMult(nil, common.Address{}, common.Address{}, input, 0, false)
	require.NoError(t, err)

	v, err := MakeMatrixMultRetArgs().UnpackValues(ret)
	require.NoError(t, err)

	res := v[0].([][]*big.Int)
	fmt.Println(res)
	assert.Equal(t, res[0][0].String(), "14")
}
