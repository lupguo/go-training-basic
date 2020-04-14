package main

import (
	"fmt"
	"os"
)

func init() {
	defer fmt.Println("init defer print")
}

func main() {
	defer fmt.Println("the main func last print") // deferred functions are not run.
	os.Exit(0)
}
