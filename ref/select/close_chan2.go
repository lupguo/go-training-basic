package main

import (
	"fmt"
	"time"
)

func main() {
	timeout := time.After(time.Second)
	over := false
	for !over {
		select {
		case <-timeout:
			over = true
			fmt.Println("elapse 2 second")
		default:
			fmt.Println("default...")
			time.Sleep(100 * time.Millisecond)
		}
	}
}
