package main

import "fmt"

func main() {
	i := 1
	fmt.Printf("%p\n",&i)
	defer func() {
		fmt.Printf("in defer 2: %p\n",&i)
	}()
}
