package help

import (
	"log"
)

// 数组转通道
func AsChan(vs ...int) <-chan int {
	c := make(chan int)
	go func() {
		for _, v := range vs {
			c <- v
		}
		log.Println("close asChan...")
		close(c)
	}()
	return c
}
