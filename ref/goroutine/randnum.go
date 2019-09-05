// We task will find the rand number smaller than 10 in [0, 1e5) range
// main process create 10 goroutine for range randChan int,
// and each RandIntCh channel background is muti currency process write to it
package main

import (
	"fmt"
	"go-example/helper"
	"log"
	"math/rand"
	"sync"
	"time"
)

func init() {
	// stat
	helper.StatusLog(helper.ToFile, 200*time.Millisecond)
}

func main() {
	Run()
}

func Run() {
	seed := time.Now().UnixNano()
	r := rand.New(rand.NewSource(seed))

	// run 10 goroutine for generate rand int, and each goroutine RandIntCh func is currency process
	wg := sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(idx int) {
			log.Printf("Run rand generate idx %d\n", idx)

			for rnd := range RandIntToCh(r, 100000, 1000, 1e6) {
				if rnd < 100 {
					fmt.Printf("\nFind a number => % d\n", rnd)
				} else {
					fmt.Printf("%d ", rnd)
				}
			}
			wg.Done()
		}(i)
	}
	// wait for all goroutine finish
	wg.Wait()
}

// RndCh rand int chan
type RndCh chan int

// generate rand int ch
func RandIntToCh(r *rand.Rand, routines int, semaphore int, limit int) RndCh {
	rch := make(RndCh)
	// close rch when done
	wg := sync.WaitGroup{}

	// signal control, only 10 currency goroutine
	sema := make(chan struct{}, semaphore)
	for i := 0; i < routines; i++ {
		wg.Add(1)
		go func() {
			sema <- struct{}{}
			// simulate more heavy work is doing
			time.Sleep(100 * time.Millisecond)
			rch <- r.Intn(limit)
			<-sema
			wg.Done()
		}()
	}

	// wait and close rch
	go func() {
		wg.Wait()
		close(rch)
	}()

	return rch
}
