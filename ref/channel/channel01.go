package main

import (
	"fmt"
	"time"
)

func greet(c chan string) {
	fmt.Println("Hello " + <-c + "!")
	time.Sleep(1*time.Second)
	fmt.Println("Not running. Because main goroutine is out!")
}

func main() {
	c := make(chan string) //
	go greet(c)
	c <- "John"
	fmt.Println("main() stopped")
}
