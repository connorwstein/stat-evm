package precompile

import (
	"fmt"
	"math/big"
	"strconv"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/require"
)

func TestIPFSFit(t *testing.T) {
	// Fit type = 0, OLS
	input, err := MakeIPFSFitArgs().PackValues([]interface{}{"QmQqzbr5c7mLUpWRjNZZfpY7GHhZuJJcsSKKFMhXJ3hrQt", big.NewInt(0)}) // housing2.csv
	require.NoError(t, err)
	s := time.Now()
	ret, _, err := getIPFSFit(nil, common.Address{}, common.Address{}, input, 0, false)
	require.NoError(t, err)
	fmt.Println(time.Since(s))
	res, err := MakeIPFSFitRetArgs().UnpackValues(ret)
	require.NoError(t, err)
	weights := res[0].([][]*big.Int)
	intercept := res[1].([]*big.Int)
	// -122.22,37.86,21.0,7099.0,1106.0,2401.0,1138.0,8.3014 = 358500.0 // 3735255
	//vals := []float64{-122.22, 37.86, 21.0, 7099.0, 1106.0, 2401.0, 1138.0, 8.3014}
	// -122.26,37.82,40.0,624.0,195.0,423.0,160.0,0.9506,187500.0
	//vals := []float64{-122.26, 37.82, 40.0, 624.0, 195.0, 423.0, 160.0, 0.9506} // 3735293
	vals := []float64{-121.04, 39.08, 8.0, 2870.0, 526.0, 1307.0, 451.0, 3.463} // ,201700.0
	sum := float64(0)
	for i, x := range weights {
		f, _ := strconv.ParseFloat(fmt.Sprintf("%v", x[0].Int64()), 64)
		sum += vals[i] * f
		fmt.Println(x[0])
	}
	f2, _ := strconv.ParseFloat(fmt.Sprintf("%v", intercept[0].Int64()), 64)
	fmt.Println("sum", sum+f2)
}