package main

import "fmt"

func main() {
	// 39
	var slice1 []int
	slice2 := make([]int, 3)
	slice3 := make([]int, 0)
	slice1 = append(slice1, 1, 2, 3)
	slice2 = append(slice2, 0,1,2,3,4,5,6,7)
	slice3 = append(slice3, 1,2,3,4)
	fmt.Println(slice1, slice2, slice3)

	// 42 panic dereference nil pointer
	//var i *int
	//fmt.Println(*i)

	// 43 parse error: cap cannot used for cap
	//mymap := make(map[string]int)
	//fmt.Println(cap(mymap))

	// var error
	// 12b := c


}
