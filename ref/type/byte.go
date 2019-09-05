package main

import "fmt"

func main() {
	bs1 := []byte("hey")
	b2 := 'T'
	b3 := 'K'
	s1 := "man!"

	fmt.Printf("%T,%[1]v\n", string(bs1)+s1)
	fmt.Printf("%T,%[1]v, %[1]q, %#[1]U\n", b2+b3)
	fmt.Printf("%q, %#[1]v, %[1]U, %#[1]U", '\u2691')
}
