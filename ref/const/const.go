package main

import "fmt"

type A struct {

}

const (
	a         = 1
	b float64 = 1.86
	c         = "test"
	d         = -1
)

func main() {
	fmt.Printf("%T,%T,%T,%T", a, b, c, d)
}
