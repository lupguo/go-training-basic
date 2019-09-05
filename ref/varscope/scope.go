package main

import (
	"fmt"
	"math/rand"
)

type (
	VPk int
	vpk float64
)

var (
	i VPk = VPk(rand.Intn(10))
	v vpk = vpk(rand.Float64() * 1.86)
)

func scopeA() {
	fs := func() {
		fmt.Println(i, v)
	}
	fs()
}

func scopeB() {
	var i2 interface{}
	if v > 1 {
		type vbig int
		i2 = vbig(v)
	} else {
		type vsmall float64
		i2 = vsmall(v)
	}
	fmt.Printf("type=>%#v, %[1]T", i2)
}

func main() {
	scopeA()
	scopeB()
}
