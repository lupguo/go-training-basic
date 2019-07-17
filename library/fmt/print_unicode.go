package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	seed := time.Now().UnixNano()
	r := rand.New(rand.NewSource(seed))

	for rndInt := range RandIntCh(r,5) {
		fmt.Printf("% d", rndInt)
	}
}

// rand int chan
type RndCh chan int

func RandIntCh(r *rand.Rand, n int) <-chan int {
	ch := make(RndCh)

	// signal control, only 10 currency goroutine
	signal := make(chan struct{}, 10)

	for i := 0; i < n; i++ {
		signal <- struct{}{}
		go func(r *rand.Rand) {
			ch <- r.Intn(1e5)
			<-signal
		}(r)
	}

	return ch
}
