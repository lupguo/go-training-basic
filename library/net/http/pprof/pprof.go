package main

import (
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
	"time"
)

func main() {

	go func() {
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()

	for i := 0; i < 60; i++ {
		st := time.Now()
		f := fib(i)
		et := time.Since(st)
		fmt.Printf("fib(%d)=%d elapse: %d ms\n", i, f, et.Nanoseconds()/1e6)
	}

}

func fib(n int) int {
	if n < 2 {
		return n
	}
	return fib(n-1) + fib(n-2)
}
