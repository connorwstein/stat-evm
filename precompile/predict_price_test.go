package precompile

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/require"
)

func TestPredictPrice(t *testing.T) {
	input, err := MakePredictPriceArgs().Pack(big.NewInt(10), big.NewInt(10))
	require.NoError(t, err)
	med, _, _ := predictPrice(nil, common.Address{}, common.Address{}, input, 100_000, true)

	require.NoError(t, err)
	t.Log(med)
}
