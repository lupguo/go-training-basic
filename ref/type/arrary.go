package main

import (
	"fmt"
	"math/rand"
)

func main() {
	arr1 := [3]int{}
	var (
		arr2 [2 * 2]int
		arr3 [5]int
		arr4 [5]float64
		arr5 [2][3]int
	)

	arr6 := [...]int{3, 5, 6}

	for i, n := range rand.Perm(10) {
		if i < 5 {
			arr3[i] = n
		}
	}
	fmt.Println(arr1, arr2, arr3, arr4, arr5, arr6)
}
