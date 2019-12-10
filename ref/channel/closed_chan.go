package main

import "fmt"

func main() {
	ch := make(chan int)
	close(ch)
	for i := 0; i < 5; i++ {
		fmt.Println(<-ch)
	}
	ch = make(chan int)
	go func() {
		ch <- 50
	}()
	if v, ok := <-ch; ok { // ok 为true表示还可以接收，false表示不可接收，表示通道被关闭
		fmt.Println("receive ", v)
	}
	close(ch)
	ch <- 100 // panic: send on closed channel
}
