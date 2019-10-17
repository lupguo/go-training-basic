// 冒泡排序：相邻元素比较，每一轮冒泡排序，将使得一个元素到最大或最小有序位置
package main

import (
	"fmt"
	"go-example/algorithm/mysort/internal/nlist"
)

func bubbleSort(A []int, N int) []int {
	// 外层循环，指定循环趟数，默认认为至多需要N-1(可以通过flag无交换标识优化退出外层循环趟数，表示所有列表已有序)
	for i := 0; i < N-1; i++ {
		isSorted := true
		// 内层循环，从第一个元素向后比较，彼此相邻元素进行相互比较A[j] vs A[j+1]
		for j := 0; j < N-1-i; j++ {
			if A[j+1] < A[j] { // 将大元素冒泡上去，即按从小大大排序，每轮从0开始，将大元素冒泡上去
				A[j], A[j+1] = A[j+1], A[j]
				isSorted = false
			}
		}
		if isSorted {
			return A
		}
	}
	return A
}

func main() {
	a := nlist.GetList(1e3, true)
	fmt.Println(bubbleSort(a, len(a)))
}
