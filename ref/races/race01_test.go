package races

import (
	"log"
	"testing"
)

func TestPrintNum(t *testing.T) {
	sum := 0
	getCh := func(n int) <-chan int {
		ch := make(chan int)
		go func(n int) {
			for i := 0; i < n; i++ {
				ch <- 1
			}
			close(ch)
		}(n)
		return ch
	}

	// sum
	for range getCh(1000) {
		sum++
	}
	log.Println(sum)
}
