package main

import (
	"fmt"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	// timeout setting
	timeout := make(chan bool, 1)
	go func() {
		time.Sleep(1 * time.Second)
		timeout <- true
	}()

	// exec
	done := make(chan struct{})
	go func() {
		time.Sleep(time.Duration(rand.Intn(3)+1) * 500 * time.Millisecond)
		done <- struct{}{}
	}()

	// select
	select {
	case <-timeout:
		fmt.Println("exec timeout")
	case <-done:
		fmt.Println("exec done")
	}
}
