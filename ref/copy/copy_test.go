package main

import (
	"runtime"
	"testing"
)

var a = func() []int {
	s := make([]int, 1e6)
	for i := 0; i < 1e3; i++ {
		s[i] = i
	}
	return s
}()

func copyA() {
	b := make([]int, len(a))
	copy(b, a)
}

func copyB() {
	b := make([]int, len(a))
	b = append([]int{}, a...)
	//b = append(a[:0:0], a...)
	_ = b
}

func BenchmarkCopyCompare(b *testing.B) {
	runtime.GOMAXPROCS(4)
	tests := []struct {
		name string
		fn   func()
	}{
		{"copyA", copyA},
		{"copyB", copyB},
	}
	for _, test := range tests {
		b.Run(test.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				test.fn()
			}
		})
	}
}
