// 选择排序：已排序区间|未排序区间 (和插入排序类似, 但选择排序会每次从未排序区间找到最小元素，放入已排序区间末尾)
//		已排序区间扩张: A[0], A[1],... A[N-1]
//		外层只需要经过N-1轮内部比较即可以完成排序(临界点i=N-2, 内循环j=N-1)
package main

import "fmt"

func selectSort(A []int, N int) []int {
	var min int
	for i:=0; i< N-1; i++ {
		min = i
		for j:=i+1; j< N; j++ {
			if A[j] < A[min] {
				min = j
			}
		}
		A[i], A[min] = A[min], A[i]
	}
	return A
}

func main() {
	a := []int{3,2,1,5,6,8,9,7,4,0}
	fmt.Println(selectSort(a, len(a)))
}
