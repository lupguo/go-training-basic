package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func main() {

	tick := time.NewTicker(1 * time.Second)
	interrupt := make(chan string)

	// input anything to interrupt ticker
	go func() {
		input := bufio.NewScanner(os.Stdin)
		for input.Scan() {
			tick.Stop()
			interrupt <- "over"
		}
	}()

	count := 5
printPoint:
	for {
		select {
		case t := <-tick.C:
			fmt.Println(t, count)
			if count--; count == 0 {
				break printPoint
			}
		case <-interrupt:
			break printPoint
		}
	}

	fmt.Println("over")
}
