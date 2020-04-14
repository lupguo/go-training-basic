package main

import (
	"fmt"
)

var a string
var ch = make(chan int, 10)

func main() {
	//changeA()
	hello()
}

func f() {
	print(a)
	//close(ch)
	ch <- 0
}

func hello() {
	a = "hey, man"
	//time.Sleep(time.Second)
	go f()
	<-ch
}

func changeA() {
	go func() {
		a = "hello"
	}()
	//time.Sleep(time.Second)
	fmt.Println("a=>", a)
}
