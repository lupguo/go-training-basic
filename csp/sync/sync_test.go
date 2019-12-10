package sync

import (
	"sync"
	"testing"
)

func TestSync(t *testing.T) {
	count := 0
	var wg sync.WaitGroup
	//var mu sync.Mutex
	ch := make(chan int, 10)
	for i := 0; i < 5000; i++ {
		wg.Add(1)
		go func() {
			defer func() {
				wg.Done()
				//mu.Unlock()
			}()
			//mu.Lock()
			ch <- 1
		}()
	}
	wg.Wait()
	for range ch {
		count++
	}
	t.Logf("count %d", count)
}
