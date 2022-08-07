package precompile

import (
	"fmt"
	"io/ioutil"
	"math/big"
	"net/http"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/require"
)

func TestIPFS(t *testing.T) {
	// ipfs pin
	a, err := http.Get("http://127.0.0.1:8080/ipfs/QmbZCvBrHFDCpSzuJWsQyjNCh6VMBeKGWQtQPYwmUgZxk9")
	require.NoError(t, err)
	b, err := ioutil.ReadAll(a.Body)
	require.NoError(t, err)
	t.Log(string(b))

	input, err := MakeIPFSMomentArgs().PackValues([]interface{}{"bafybeihwbgtlimdj3fnopubripcds4wtn5yedvlmxag7l6wzya5bzps45i/housing.csv", big.NewInt(1)})
	require.NoError(t, err)
	s := time.Now()
	ret, _, err := getIPFSMoment(nil, common.Address{}, common.Address{}, input, 0, false)
	require.NoError(t, err)
	fmt.Println(time.Since(s))
	res, err := MakeIPFSMomentRetArgs().UnpackValues(ret)
	require.NoError(t, err)
	fmt.Println(res)
}
