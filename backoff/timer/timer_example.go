package main

import (
	"fmt"
	"time"
)

func main() {

	timer := time.NewTimer(1 * time.Second)
	atch := time.After(10 * time.Second)
	over := make(chan string)

	// after 5 seconds close the timer
	go func() {
		time.AfterFunc(3*time.Second, func() {
			timer.Stop()
		})
	}()

	// go on print time.Timer or stop the goroutine when time after 10 seconds
	go func() {
		count := 5
		for {
			select {
			case at := <-atch:
				fmt.Println(at, "time.After Launch", count)
				over <- "over"
				return
			case t := <-timer.C:
				timer.Reset(1 * time.Second)
				fmt.Println(t, "time.Timer", count)
				if count--; count == 0 {
					over <- "over"
					return
				}
			}
		}
	}()

	fmt.Println(<-over)
	//fmt.Println(runtime.NumGoroutine())
}
