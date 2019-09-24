package main

import (
	"fmt"
)

// 使用两个 goroutine 交替打印序列，一个 goroutinue 打印数字，
// 另外一个goroutine打印字母， 最终效果如下 12AB34CD56EF78GH910IJ 。
func main() {
	chC := make(chan struct{})
	chN := make(chan struct{})
	fin := make(chan struct{})

	// print num
	go func() {
		for i := 1; i <= 10; i = i + 2 {
			fmt.Printf("%d%d", i, i+1)
			chC <- struct{}{}
			<-chN
		}
		fin <- struct{}{}
	}()

	// print char
	go func() {
		for c := 'A'; c <= 'J'; c = c + 2 {
			<-chC
			fmt.Printf("%c%c", c, c+1)
			chN <- struct{}{}
		}
	}()

	<-fin
}
