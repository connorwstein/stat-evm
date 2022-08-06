package precompile

import (
	"fmt"
	"io/ioutil"
	"math/big"
	"net/http"
	"testing"

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

	input, err := MakeIPFSMomentArgs().PackValues([]interface{}{"QmbZCvBrHFDCpSzuJWsQyjNCh6VMBeKGWQtQPYwmUgZxk9", big.NewInt(1)})
	require.NoError(t, err)
	ret, _, err := getIPFSMoment(nil, common.Address{}, common.Address{}, input, 0, false)
	require.NoError(t, err)
	res, err := MakeIPFSMomentRetArgs().UnpackValues(ret)
	require.NoError(t, err)
	fmt.Println(res)
}
