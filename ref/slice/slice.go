package main

import (
	"fmt"
	"os"
	"os/signal"
	"time"
)

func main() {
	sl1 := make([]int, 2, 4)
	stop := make(chan struct{})

	// sl1 is always 4
	go func(s []int, stop chan struct{}) {
		for {
			select {
			case <-stop:
				return
			default:
				fmt.Printf("souce cap of slice: %d \n", cap(s))
				time.Sleep(500 * time.Millisecond)
			}
		}
	}(sl1, stop)

	for i := 0; i < 10000000; i++ {
		// append to new slice (memory copy)
		sl1 = append(sl1, i)
		if i%500 == 0 {
			time.Sleep(200 * time.Millisecond)
			fmt.Printf("%p, cap = %d\n", sl1, cap(sl1))
		}
	}
	// ctrl + c exit
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt)
	if v, ok := <-c ; ok {
		fmt.Printf("%#v, %#v\n", v, ok)
		stop <- struct{}{}
	}
}
