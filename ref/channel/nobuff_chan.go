package main

import (
	"fmt"
	"time"
)

func main() {
	// buffer channel
	bufch := make(chan int, 1)
	bufch <- 88

	// panic deadlock
	nbch := make(chan int)
	// must set no buffer channel receiver, handle deadlock
	go func() {
		close(bufch)
		fmt.Println(<-nbch,<-nbch, <-bufch, <-bufch, <-bufch)
	}()
	nbch <- 100
	close(nbch)
	time.Sleep(time.Second)
}
