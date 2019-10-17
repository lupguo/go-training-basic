package main

import "fmt"

// 合并两个有序数组列表
func mergeArrayList(A, B []int) []int {
	lenA, lenB := len(A), len(B)
	list := make([]int, lenA+lenB)
	var i, j int
	for i < lenA && j < lenB {
		if A[i] <= B[j] {
			list[i+j] = A[i]
			i++
		} else {
			list[i+j] = B[j]
			j++
		}
	}
	if i == lenA && j < lenB {
		copy(list[i+j:], B[j:])
	} else if i < lenA && j == lenB {
		copy(list[i+j:], A[i:])
	}
	return list
}

func main() {
	a := []int{1, 3, 5}
	b := []int{2, 4, 6, 8, 10}
	fmt.Println(mergeArrayList(a, b))
}
