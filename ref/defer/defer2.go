package main

import "fmt"

func trace(s string) string {
	fmt.Println("entering:", s)
	return s
}

func un(s string) {
	fmt.Println("leaving:", s)
}

func a() {
	// trace("a") run immediately, so trace func run first, print "entering a", and return string "a", then defer un("a")
	defer un(trace("a"))
	fmt.Println("in a")
}

func b() {
	// trace("b") run immediately, so trace func run first, print "entering b", and return string "b", then defer un("b")
	defer un(trace("b"))
	fmt.Println("in b")
	a()
}

func main() {
	b()
	//output
	// entering b , defer un("b")
	// in b
	// a()
	// entering a, defer un("a")
	// in a
	// leaving a ->defer un("a")
	// leaving b ->defer un("b")
}
