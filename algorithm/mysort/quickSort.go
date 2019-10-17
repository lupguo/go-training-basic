package main

import (
	"fmt"
	"go-example/algorithm/mysort/internal/nlist"
)

// 快排启动函数，驱动例程
func qSort(A []int, N int) []int {
	myQuickSort(A, 0, N-1)
	return A
}

// 快排函数
// 		1. 递归基准，上下界相同
// 		2. partition函数，进行[left, right]区间分区
//		3. 分别针对左、右区间进行快排
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

// 分区函数，基于pivot基准点，将列表A，划分为在left区间到right区间，
//		1. 选取一个pivot基准元素(假定选取最右元素A[right]为基准点）
//		2. 申请i, j 游标，分别指向列表的最左(i=left)和最右元素的前一个元素(j=right-1)，注意left和right右边可能不是从0和N-2开始
//		3. 开始循环，循环内，先持续迭代移动i(i++), 直至A[i] > pivot，暂停；开始迭代向前移动j(j--) ，直至A[j] < pivot，暂停;
//			若i<j，则交换i,j，否则退出循环
//
func partition(A []int, left, right int) (index int) {
	// 选取列表最后一个元素为pivot(其他方式有取中值、随机等）
	pivot := midPivot(A, left, right)

	// 从[left,right-1]区间内，选取i,j游标，依次参照pivot基准点进行对照和调整，调整后满足：小数区间|基准点|大数区间
	//  - 注意这里的坑: 需要保证i,j的有效性，left-1实际已越界，在内存for循环通过初试(i++)确保不会越界(不这样，会导致同数据列表死循环，持续互换)
	//  (需要考虑3类情况，逆序、顺序、无差异数列)
	i, j := left, right-1
	for {
		// i从左往右，在满足不超过j且值小于基准值下，依次加1，否则找到一个大于pivot的元素，准备和右侧交换
		for ;A[i] < pivot; i++ {
		}
		// j从右往左，在满足大于i且值大于基准值下，依次减1，否则找到一个小于pivot的元素，准备和左侧交换
		for ;j > 0 && pivot < A[j]; j-- {
		}
		// 在满足i < j情况下，才进行i,j元素交换，否则退出循环
		if i < j {
			A[i], A[j] = A[j], A[i]
		} else {
			break
		}
	}
	// 将最右项pivot，与放到i的位置对调
	A[i], A[right] = A[right], A[i]
	return i
}

func midPivot(A []int, left int, right int) int {
	mid := (right - left) / 2
	if A[left] > A[mid] {
		A[left], A[mid] = A[mid], A[left]
	}
	if A[mid] > A[right] {
		A[mid], A[right] = A[right], A[mid]
	}
	if A[left] > A[mid] {
		A[left], A[mid] = A[mid], A[left]
	}

	// 交换mid中值和最后一个元素
	A[mid], A[right] = A[right], A[mid]
	return A[right]
}

func main() {
	//list := nlist.GetList(1e2, true)
	list := nlist.GetList(1e3, false)
	//list := nlist.GetConsList(1e1)
	//fmt.Println(list)
	fmt.Println(qSort(list, len(list)))
}
