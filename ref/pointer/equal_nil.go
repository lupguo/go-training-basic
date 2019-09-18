package main

import "fmt"

func main() {
	var x, y, z int
	fmt.Printf("%v, %[1]q, %`123[1]x\n", &x)
	fmt.Printf("%p %x %q\n", &x, &y, &z)
	fmt.Printf("%v, %v, %v", &x == &x, &x == &y, &x == &z)
}
