// see https://stackoverflow.com/questions/37469995/goroutines-are-cooperatively-scheduled-does-that-mean-that-goroutines-that-don
package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func prt(x int, wg *sync.WaitGroup) {
	if x%2 == 0 {
		time.Sleep(2 * time.Second)
	}
	fmt.Println(x)
	wg.Done()
}

func main() {
	runtime.GOMAXPROCS(1)
	wg := &sync.WaitGroup{}
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go prt(i, wg)
	}
	wg.Wait()
	//sign := make(chan os.Signal)
	//signal.Notify(sign, os.Interrupt)
	//<-sign
}
