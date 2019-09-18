package main

import (
	"fmt"
)

// fib returns a channel which transports fibonacci numbers
func fib(length int) <-chan int {
	// make buffered channel
	c := make(chan int, length)

	// run generation concurrently
	go func() {
		for i, j, k := 0, 1, 0; k < length; i, j, k = i+j, i, k+1 {
			c <- i
		}
		close(c)
	}()

	// return channel
	return c
}

func main() {
	// read 10 fibonacci numbers from channel returned by `fib` function
	for fn := range fib(20) {
		fmt.Println("Current fibonacci number is", fn)
	}
}
