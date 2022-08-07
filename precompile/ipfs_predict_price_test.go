package precompile

import (
	"fmt"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/require"
)

func TestIPFSPredictPrice(t *testing.T) {
	input, err := MakeIPFSPredictPriceArgs().Pack("QmZbV8nB1uQf5NqhsXB8RhGNJpUr4v92Br4jw6a7MdExvz", big.NewInt(10), big.NewInt(10))
	require.NoError(t, err)
	prediction, _, _ := predictPriceIPFS(nil, common.Address{}, common.Address{}, input, 100_000, true)
	fmt.Println("prediction", prediction)
	require.NoError(t, err)
	t.Log(prediction)
}
