package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	rwm := &sync.RWMutex{}
	wg := sync.WaitGroup{}
	wg.Add(6)

	go func(rwm *sync.RWMutex) {
		rwm.Lock()
		fmt.Println("Goroutine Get Write lock")
		time.Sleep(1 * time.Second)
		rwm.Unlock()
		fmt.Println("Goroutine Release Write Unlock")
		wg.Done()
	}(rwm)

	for i := 1; i <= 10; i++ {
		go func(i int) {
			rwm.RLock()
			fmt.Println("Goroutine", i, "Get Read lock")
			time.Sleep(500 * time.Millisecond)
			rwm.RUnlock()
			fmt.Println("Goroutine", i, "Release Read lock")
			wg.Done()
		}(i)
	}

	rwm.Lock()
	fmt.Println("Main Goroutine Get Write lock")
	rwm.Unlock()
	fmt.Println("Main Goroutine Release Write lock")
	wg.Wait()
}
