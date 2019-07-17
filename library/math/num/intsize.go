package main

import "fmt"

func main() {
	for i := uint(0); i < 65; i++ {
		fmt.Printf("1<<%d=%d\n", i, 1<<i)
	}
}
