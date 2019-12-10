package procon

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
	"time"
)

type Work struct {
	wid int
}

func goRecord(t *testing.T) {
	tk := time.NewTicker(time.Millisecond * 100)
	go func() {
		for {
			select {
			case tic := <-tk.C:
				t.Logf("%s, goroutines numbers: %d", tic, runtime.NumGoroutine())
			}
		}
	}()
}

func TestProCon(t *testing.T) {
	goRecord(t)
	logicCpu := runtime.NumCPU()
	t.Logf("logic cpu numbers: %d", logicCpu)
	runtime.GOMAXPROCS(logicCpu)
	var ch = producer()
	for w := range ch {
		doWork(w)
	}
}

func doWork(work Work) {
	fmt.Println("do working =>", work.wid)
}

func producer() chan Work {
	wch := make(chan Work)
	var wg sync.WaitGroup
	go func() {
		wg.Wait()
		close(wch)
	}()
	for i := 0; i < 10000; i++ {
		wg.Add(1)
		go func(wid int) {
			defer func() {
				wg.Done()
			}()
			wch <- Work{wid: wid}
		}(i)
	}
	return wch
}
