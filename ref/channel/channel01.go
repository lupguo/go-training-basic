package main

import (
	"fmt"
	"time"
)

func greet(c chan string) {
	fmt.Println("Hello " + <-c + "!")

	// finish channel block, main goroutine will go on process
	time.Sleep(1*time.Second)
	fmt.Println("Not running. Because main goroutine is out!")
}

func main() {
	fmt.Println("main() started")
	c := make(chan string)

	go greet(c)

	c <- "John"
	fmt.Println("main() stopped")
}
