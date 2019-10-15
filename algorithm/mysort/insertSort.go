package main

import "fmt"

func insertSort(A []int, N int) []int {
	for i := 1; i < N; i++ { //N-1轮插入检测
		tmp := A[i]                        //待插入数元素
		j := i                             // 迭代计数器
		for ; j > 0 && A[j-1] > tmp; j-- { //j从最大元素开始，往前检测，满足条件:j未迭代完或前一个元素位置大于<待插入数元素>，则前一个元素后移
			A[j] = A[j-1]
		}
		A[j] = tmp //流出的空位，即为满足插入的位置(将待插入元素)插入到指定位置
	}
	return A
}

func main() {
	A := []int{1, 3, 4, 6, 8, 5, 7, 9, 2}
	fmt.Println(A)
	fmt.Println(insertSort(A, len(A)))
}
