package main

import "fmt"

func main() {
	slice := []int{0,1,2,3}
	fmt.Printf("slice[0:0]=%v, len=%v, cap=%v\n", slice[0:0], len(slice[0:0]), cap(slice[0:0]))
	fmt.Printf("slice[0:1]=%v, len=%v, cap=%v\n", slice[0:1], len(slice[0:1]), cap(slice[0:1]))
}
