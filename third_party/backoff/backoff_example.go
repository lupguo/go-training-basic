package main

import (
	"errors"
	"fmt"
	"github.com/cenkalti/backoff"
	"math/rand"
	"time"
)

// Seed
func init() {
	rand.Seed(time.Now().UnixNano())
}

// rpc result
type RpcResult struct {
	code int
	msg  string
	data interface{}
}

// common rpc func
func doRpc() (*RpcResult, error) {

	fmt.Println("try rpc executing... ", time.Now().Format("15:04:05"))

	if rand.Intn(10) < int(SuccessProb*10) {
		return &RpcResult{100, "result is ok", "ok"}, nil
	} else {
		return nil, errors.New("rand error")
	}
}

// Retry times
var MaxRetry uint64 = 3
var SuccessProb = 0.5

// doRpc with backoff + retry
func DoRpcRetry() (*RpcResult, error) {

	rpcrs := new(RpcResult)

	opfn := func() (rpcerr error) {
		rpcrs, rpcerr = doRpc()
		return rpcerr
	}

	// exponential back off
	bkf := backoff.WithMaxRetries(backoff.NewExponentialBackOff(), MaxRetry)
	err := backoff.Retry(opfn, bkf)
	if err != nil {
		return nil, err
	}
	return rpcrs, err
}

func main() {
	data, err := DoRpcRetry()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%#v", *data)
	}
}
