package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	close(ch)
	done := false
	for !done {
		select {
		case _, ok := <-ch:
			if !ok {
				fmt.Println("ch is closed")
				time.Sleep(50 * time.Millisecond)
				// set chan to nil for block the channel
				ch = nil
				done = true
			}
		}
	}
}
