package main

import "fmt"

func qsort(A []int, N int) []int {
	myQuickSort(A, 0, N-1)
	return A
}

// 分区函数，针对列表A，在l区间到r区间，选项一个pivot进行对比后分区
func partition(A []int, l, r int) (index int) {
	// 选取列表最后一个元素为pivot(其他方式有取中值、随机等）
	pivot := A[r]
	i, j := l, r-1
	for {
		// i从左往右，在满足不超过j且值小于基准值下，依次加1，否则找到一个大于pivot的元素，准备和右侧交换
		for A[i] < pivot {
			i++
		}
		// j从右往左，在满足大于i且值大于基准值下，依次减1，否则找到一个小于pivot的元素，准备和左侧交换
		for pivot < A[j] {
			j--
		}
		// 在满足i < j情况下，才进行i,j元素交换，否则退出循环
		if i < j {
			A[i], A[j] = A[j], A[i]
		} else {
			break
		}

	}
	// 将pivot项放到i=j的位置
	A[i], A[r] = A[r], A[i]
	return i
}

func myQuickSort(A []int, left, right int) {
	// 1. 快排递归返回条件
	if left >= right {
		return
	}

	// 2. 分区: 选取一个基准值pivot(最后元素、随机元素、小中大中值)，进行partition分区，
	// 	返回分割点(确保在分割点左边为小元素，右边为大元素)
	p := partition(A, left, right)

	// 3. 针对分区点左边元素进行快速排序 && 针对分区点右边元素进行快速排序
	myQuickSort(A, 0, p-1)
	myQuickSort(A, p+1, right)
}

func main() {
	list := []int{1, 3, 4, 5, 8, 2, 9, 7, 6}
	fmt.Println(list)
	fmt.Println(qsort(list, len(list)))
}
