package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan int)
	go func() {
		time.Sleep(3 * time.Second)
		close(ch1)
		//close(ch1) // panic: close of closed channel
	}()
	<-ch1
	<-ch1 // receive zero val
	fmt.Println("block done")

	if v, ok := <-ch1; ok {
		fmt.Println("channel is open, val ", v)
	} else {
		fmt.Println("channel is closed.")
	}
}
