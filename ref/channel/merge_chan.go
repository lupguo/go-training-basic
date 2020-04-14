package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"
)

func main() {
	a := asChan(1, 3, 5, 7)
	b := asChan(2, 4, 6, 8, 9, 10, 100, 1000)
	mc := merge(a, b)
	for v := range mc {
		fmt.Println(v)
	}
}

// 数组转通道
func asChan(vs ...int) <-chan int {
	c := make(chan int)
	go func() {
		for _, v := range vs {
			c <- v
			time.Sleep(time.Duration(rand.Intn(50)) * time.Millisecond)
		}
		fmt.Println("close asChan...")
		close(c)
	}()
	return c
}

// 合并通道
func merge(a, b <-chan int) <-chan int {
	c := make(chan int)
	go func() {
		defer close(c)
		adone, bdone := false, false
		for !adone || !bdone {
			select {
			case v, ok := <-a:
				if !ok {
					fmt.Println("a is done")
					a = nil
					adone = true
					continue
				}
				fmt.Println("receive v from <-a")
				c <- v
			case v, ok := <-b:
				if !ok {
					log.Println("b is done")
					b = nil
					bdone = true
					continue
				}
				fmt.Println("receive v from <-b")
				c <- v
			}
		}
	}()
	return c
}
