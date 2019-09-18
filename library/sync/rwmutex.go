package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	rwLocker()

	wLocker()
}

func rwLocker() {
	defer func(start time.Time) {
		fmt.Printf("read locker time elapsed %v\n", time.Since(start))
	}(time.Now())

	var rwl sync.RWMutex
	for i := 0; i < 100000; i++ {
		go func() {
			rwl.RLock()
			rwl.RUnlock()
		}()
	}
}

func wLocker() {
	defer func(start time.Time) {
		fmt.Printf("write locker time elapsed %v\n", time.Since(start))
	}(time.Now())

	var wl sync.Mutex
	for i := 0; i < 100000; i++ {
		go func() {
			wl.Lock()
			wl.Unlock()
		}()
	}
}
