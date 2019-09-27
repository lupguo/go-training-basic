package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int, 100)
	fmt.Printf("1) len=%d, cap=%d\n", len(ch1), cap(ch1)) // 0, 0
	fmt.Printf("2) len=%d, cap=%d\n", len(ch2), cap(ch2)) // 0, 1
	ch2 <- 1
	fmt.Printf("3) len=%d, cap=%d\n", len(ch2), cap(ch2)) // 1, 1
	//close(ch2)
	<-ch2
	fmt.Printf("4) len=%d, cap=%d\n", len(ch2), cap(ch2)) // 0, 1

	go func() {
		for i := 0; ; i++ {
			ch2 <- i
			time.Sleep(3 * time.Second)
		}
	}()

	for v := range ch2 {
		fmt.Printf("v=>%d\n", v)
	}
}
