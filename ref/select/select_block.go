package main

import "fmt"

func main() {
	ok := make(chan int)
	select {
	case <-ok:
	}
	fmt.Println("over")
}
