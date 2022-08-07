package main

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/ethclient"
)

func confirmFail(ec *ethclient.Client, txHash common.Hash) {
	for {
		_, pending, err2 := ec.TransactionByHash(context.Background(), txHash)
		fmt.Println(pending, err2)
		if err2 != nil {
			return
		}
		// fmt.Println("Pending:", pending)
		// fmt.Println("Error:", err2)
		if !pending {
			receipt, err3 := ec.TransactionReceipt(context.Background(), txHash)
			fmt.Println(err3, receipt.GasUsed, receipt.Status)
			break
		}
		time.Sleep(time.Second)
	}
}

func confirm(ec *ethclient.Client, txHash common.Hash) {
	for {
		_, pending, err2 := ec.TransactionByHash(context.Background(), txHash)
		//panicErr(err2)
		fmt.Println("Pending:", pending)
		fmt.Println("Error:", err2)
		if !pending || (err2 != nil && err2.Error() != "not found") {
			receipt, err3 := ec.TransactionReceipt(context.Background(), txHash)
			fmt.Println(err3, receipt.GasUsed, receipt.Status)
			break
		}
		time.Sleep(time.Second)
	}
}

func panicErr(err error) {
	if err != nil {
		panic(err)
	}
}

func start_events(ec *ethclient.Client, addr common.Address) {
	ec, err := ethclient.Dial("ws://127.0.0.1:27044/ext/bc/t5jkcEANP8EXobp3MRGeDiBZpuQ4Rp8359Vd5vVaagdgwdW4u/ws")
	query := ethereum.FilterQuery{
		Addresses: []common.Address{addr},
	}

	logs := make(chan types.Log)
	sub, err := ec.SubscribeFilterLogs(context.Background(), query, logs)
	if err != nil {
		log.Fatal(err)
	}

	for {
		select {
		case err := <-sub.Err():
			log.Fatal(err)
		case vLog := <-logs:
			fmt.Println(vLog) // pointer to event log
			fmt.Println()
		}
	}
}

func main() {
	ec, err := ethclient.Dial("http://127.0.0.1:11918/ext/bc/2YTfhNo7gDm3pdnfdtboQMcKQ6UDrWBX1MTdcuBF5Ch1Wqqrj5/rpc")
	panicErr(err)

	b, err := ec.ChainID(context.Background())
	panicErr(err)

	key, err := crypto.HexToECDSA("56289e99c94b6912bfc12adc093c9b51124f0dc54ac7a766b2bc5ccf558d8027")
	panicErr(err)

	user, err := bind.NewKeyedTransactorWithChainID(key, b)
	panicErr(err)

	addr, deployTx, testContract, err := DeployMain(user, ec)
	panicErr(err)

	fmt.Println("contract address", addr)

	confirm(ec, deployTx.Hash())

	user.GasLimit = 10_000_000
	//tx, err := testContract.TestMedian(user, []*big.Int{big.NewInt(5), big.NewInt(10), big.NewInt(3), big.NewInt(7), big.NewInt(12)})
	//panicErr(err)
	//confirm(ec, tx.Hash())
	//// fmt.Println("Tx hash (median):", tx.Hash())
	//l, err := testContract.Med(nil)
	//panicErr(err)
	//fmt.Println("median", l, err)
	//
	//sampler_tx, err := testContract.TestSampler(user, big.NewInt(0), big.NewInt(0), big.NewInt(1000000000000000000), big.NewInt(5))
	//panicErr(err)
	//confirm(ec, sampler_tx.Hash())
	//// fmt.Println("Tx hash (sampler):", sampler_tx.Hash())
	//
	//sampler_res, err := testContract.GetLastSample(nil)
	//panicErr(err)
	//fmt.Println("sample result", sampler_res, err)
	//
	//moment_tx, err := testContract.TestMoment(user, big.NewInt(2), []*big.Int{big.NewInt(23), big.NewInt(31), big.NewInt(23)}, []*big.Int{big.NewInt(1), big.NewInt(1), big.NewInt(1)})
	//panicErr(err)
	//confirm(ec, moment_tx.Hash())
	//// fmt.Println("Tx hash (moment_tx):", moment_tx.Hash())
	//
	//moment_res, err := testContract.Moment(nil)
	//panicErr(err)
	//fmt.Println("moment result", moment_res, err)
	//
	//var a [][]*big.Int
	//for i := 0; i < 1; i++ {
	//	a = append(a, []*big.Int{big.NewInt(1), big.NewInt(2), big.NewInt(3)})
	//}
	//var b2 [][]*big.Int
	//for i := 0; i < 3; i++ {
	//	b2 = append(b2, []*big.Int{big.NewInt(int64(i + 1))})
	//}
	//matrix_tx, err := testContract.TestMatrixMult(user, a, b2)
	//panicErr(err)
	//confirm(ec, matrix_tx.Hash())
	//fmt.Println("Tx hash (matrix):", matrix_tx.Hash())
	//
	//vals, err := testContract.GetMatrixMulti(nil)
	//panicErr(err)
	//fmt.Println("Matrix results", vals)
	//
	//ipfs_moment_tx, err := testContract.TestIPFSMoment(user, "QmbZCvBrHFDCpSzuJWsQyjNCh6VMBeKGWQtQPYwmUgZxk9", big.NewInt(1))
	//panicErr(err)
	//confirm(ec, ipfs_moment_tx.Hash())
	//// fmt.Println("Tx hash (matrix):", matrix_tx.Hash())
	//res, err := testContract.IpfsMomentRes(nil)
	//panicErr(err)
	//fmt.Println("IPFS results", res)
	//fmt.Println(vals)
	//
	//var x [][]*big.Int
	//var y [][]*big.Int
	//for i := 0; i < 100; i++ {
	//	val := big.NewInt(int64(i + 1))
	//	val_times_two := new(big.Int).Mul(val, big.NewInt(2))
	//	x = append(x, []*big.Int{val})
	//	y = append(y, []*big.Int{val_times_two})
	//}
	//
	//fitType := big.NewInt(0) // 0 = OLS, 1 = LASSO
	//fit_tx, err := testContract.TestFit(user, fitType, x, y)
	//panicErr(err)
	//confirm(ec, fit_tx.Hash())
	//fmt.Println("Tx hash (fit):", fit_tx.Hash())
	//
	//fit_vals, err := testContract.GetFitted(nil)
	//panicErr(err)
	//fmt.Println(fit_vals)

	// Data does not exist should error
	ipfs_fit_tx_fail, err := testContract.TestIPFSFit(user, "asldflkjasdf", big.NewInt(0))
	fmt.Println(ipfs_fit_tx_fail, err)

	//confirmFail(ec, ipfs_fit_tx_fail.Hash())
	//_, pending, err2 := ec.TransactionByHash(context.Background(), ipfs_fit_tx_fail.Hash())
	//fmt.Println("non-existent data rejected from mempool")

	ipfs_fit_tx, err := testContract.TestIPFSFit(user, "QmQqzbr5c7mLUpWRjNZZfpY7GHhZuJJcsSKKFMhXJ3hrQt", big.NewInt(0))
	panicErr(err)
	confirm(ec, ipfs_fit_tx.Hash())
	coeffs, err := testContract.GetIPFSCoeffs(nil)
	panicErr(err)
	fmt.Println(coeffs)
	intercepts, err := testContract.GetIPFSIntercepts(nil)
	panicErr(err)
	fmt.Println(intercepts)

	// longitude,latitude,housing_median_age,total_rooms,total_bedrooms,population,households,median_income,median_house_value
	// -122, 38, 21, 7099, 1106, 2401, 1138, 8
	var idv = [][]*big.Int{
		{big.NewInt(-122), big.NewInt(38), big.NewInt(21), big.NewInt(7099), big.NewInt(1106), big.NewInt(2401), big.NewInt(1138), big.NewInt(8)},
	}
	pred_tx, err := testContract.TestPrediction(user, idv)
	panicErr(err)
	confirm(ec, pred_tx.Hash())
	pred, err := testContract.Predicted(nil)
	fmt.Println(pred.Int64())
}
