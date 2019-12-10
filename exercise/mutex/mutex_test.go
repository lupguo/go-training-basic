package mutex

import (
	"sync"
	"testing"
)

func TestSum01(t *testing.T) {
	sum := 0
	var wg sync.WaitGroup
	for i := 0; i < 5000; i++ {
		wg.Add(1)
		go func(wg *sync.WaitGroup) {
			defer wg.Done()
			sum ++
		}(&wg)
	}
	wg.Wait()
	t.Log("After 1000 operations, sum is ", sum)
}

func TestSum02(t *testing.T) {
	sum := 0
	var wg sync.WaitGroup
	var mu sync.Mutex
	for i := 0; i < 5000; i++ {
		wg.Add(1)
		go func(wg *sync.WaitGroup) {
			defer func() {
				wg.Done()
				mu.Unlock()
			}()
			mu.Lock()
			sum ++
		}(&wg)
	}
	wg.Wait()
	t.Log("After 1000 operations, sum is ", sum)
}
