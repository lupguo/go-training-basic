package main

import "fmt"

func main() {
	// init
	for i, k := 0, 100; i < 5; i++ {
		fmt.Printf("i=%d,k=%d\n", i, k)
		k--
	}

}
