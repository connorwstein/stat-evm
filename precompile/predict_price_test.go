package precompile

import (
	"fmt"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/require"
)

func TestPredictPrice(t *testing.T) {
	input, err := MakePredictPriceArgs().Pack("ETHUSD")
	require.NoError(t, err)
	med, s, s2 := predictPrice(nil, common.Address{}, common.Address{}, input, 100_000, true)
	fmt.Println("s:", s)
	fmt.Println("s2:", s2)
	require.NoError(t, err)
	t.Log(med)
}
