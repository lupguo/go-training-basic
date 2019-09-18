package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	noLock()

	lock()
	// Output
	// no mutex lock num= 9731
	// using mutex lock num= 10000
}

func noLock() {
	num := 0
	for i:=0; i< 10000; i++ {
		go func() {
			num++
		}()
	}

	time.Sleep(time.Second)
	fmt.Println("no mutex lock num=", num)
}

func lock()  {
	var m  sync.Mutex
	num := 0
	for i:=0; i< 10000; i++ {
		go func() {
			m.Lock()
			num++
			m.Unlock()
		}()
	}

	time.Sleep(time.Second)
	fmt.Println("using mutex lock num=", num)
}

