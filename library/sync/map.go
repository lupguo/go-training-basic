package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	//smap()

	omap()
}

func keyName(d int) string {
	return fmt.Sprintf("key_%d", d )
}

// sync.Map类型使用场景：
// 1. 当给定键的条目只写入一次但是多次读取时，如仅在增长的缓存中
// 2. 当多个goroutine读取，写入和覆盖不相交的键集的条目时。
// 在这两种情况下，与使用单独的互斥锁或RWMutex配对的Go地图相比，使用Map可以显着减少锁争用。
func smap()  {
	var m sync.Map
	//mm := make(map[string]int)

	for i := 0; i < 10000; i++ {
		go func(i int) {
			m.Store(keyName(i), i)
		}(i)
	}
	// be sure all map key set ok
	time.Sleep(time.Second)

	fail := 0
	for i := 0; i < 10000; i++ {
		v, ok := m.Load(keyName(i))
		if ok {
			fmt.Printf("key_%d=%d\n", i, v)
		} else {
			fail++
		}
	}
	fmt.Println("load fail times ", fail)
}

// 1. map并发写，导致竟态问题：fatal error: concurrent map writes
// 2. 可以基于sync.Mutex互斥锁进行临界区锁定(另外可以基于单个goroutine更新）
// 3. 可以基于sync.Map支持并发安全Map
func omap()  {
	//var wl sync.Mutex
	m := make(map[string]int)

	for i := 0; i < 10000; i++ {
		go func(i int) {
			//wl.Lock()
			m[keyName(i)] = i
			//wl.Unlock()
		}(i)
	}
	// be sure all map key set ok
	time.Sleep(time.Second)

	fail := 0
	for i := 0; i < 10000; i++ {
		v, ok := m[keyName(i)]
		if ok {
			fmt.Printf("key_%d=%d\n", i, v)
		} else {
			fail++
		}
	}
	fmt.Println("load fail times ", fail)
}
