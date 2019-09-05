package main

import "fmt"

var (
	arr  = [3]int{1, 2}
	parr = &arr
	ss   = []int{100, 200}
	str  = "test12"
	mm   = map[string]int{
		"red":   1,
		"green": 2,
		"blue":  3,
	}
)

func main() {
	fmt.Printf("%#v,%#v,%#v,%#v,%#v\n", arr, parr, ss, str, mm)

	// low:high:max
	a := [5]int{1, 2, 3, 4, 5}
	//s := a[1:4]
	t := a[1:4:5]
	fmt.Println(a[1:4][:], t[2:])
}
