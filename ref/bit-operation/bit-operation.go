package main

import "fmt"

func main() {
	A := 7
	B := 2
	fmt.Printf("A(%03b),B(%03b)\n", A, B)
	fmt.Printf("A(%03b)&B(%03b)%03b\n", A, B, A&B)
	fmt.Printf("A(%03b)|B(%03b)%03b\n", A, B, A|B)
	fmt.Printf("A(%03b)^B(%03b)%03b\n", A, B, A^B)
	// << 左移
	fmt.Printf("A(%03b)<<0=%08b\n", A, A<<0)
	fmt.Printf("A(%03b)<<1=%08b\n", A, A<<1)
	fmt.Printf("A(%03b)<<2=%08b\n", A, A<<2)
	// >> 右移
	fmt.Printf("A(%03b)>>0=%03b\n", A, A>>0)
	fmt.Printf("A(%03b)>>1=%03b\n", A, A>>1)
	fmt.Printf("A(%03b)>>2=%03b\n", A, A>>2)
	fmt.Printf("A(%03b)>>3=%03b\n", A, A>>3)
	fmt.Printf("A(%03b)>>4=%03b\n", A, A>>4)
	// &^
	fmt.Printf("A(%03b)&^0=%03b\n", A, A&^0)
	fmt.Printf("A(%03b)&^1=%03b\n", A, A&^1)
	fmt.Printf("A(%03b)&^2=%03b\n", A, A&^2)
	fmt.Printf("A(%03b)&^3=%03b\n", A, A&^3)
	fmt.Printf("A(%03b)&^A(%03b)=%03b\n", A, A, A&^A)

	// Readable
	mod := 7
	const (
		Executable = 1 << iota
		Writeable
		Readable
	)
	fmt.Printf("mod remove read permission: File Mod(%03b)&^Readable(%03b)=%03b\n", mod, Readable, mod&^Readable)
	fmt.Printf("mod remove write permission: File Mod(%03b)&^Writeable(%03b)=%03b\n", mod, Writeable, mod&^Writeable)
}
