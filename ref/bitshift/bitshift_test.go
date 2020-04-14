package bitshift

import (
	"fmt"
	"reflect"
	"testing"
)

func TestBitShift(t *testing.T) {
	var n = 10
	const N int = 10

	x := (1 << n) / 100
	y := (1 << N) / 100
	fmt.Println(x, y)

	var p byte = (1 << n) / 100
	var q byte = (1 << N) / 100
	fmt.Println(p, q)

	fmt.Printf("%032b, %032b", byte(1<<n), byte((1<<N)/100))
	fmt.Println(reflect.TypeOf(N), reflect.New(reflect.TypeOf(N)).String())
}
