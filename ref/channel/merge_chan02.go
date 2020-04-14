package main

import (
	"fmt"
	"go-example/ref/channel/help"
	"sync"
)

func main() {
	a := help.AsChan(1, 2, 3, 4)
	b := help.AsChan(5, 6, 7, 8)
	c := help.AsChan(9, 10, 11, 12)
	mc := mergeChs(a, b, c)
	for v := range mc {
		fmt.Println(v)
	}
}

func mergeChs(cs ...<-chan int) <-chan int {
	mc := make(chan int)
	wg := &sync.WaitGroup{}
	for _, c := range cs {
		wg.Add(1)
		go func(c <-chan int) {
			for v := range c {
				mc <- v
			}
			wg.Done()
		}(c)
	}
	go func() {
		wg.Wait()
		close(mc)
	}()
	return mc
}
