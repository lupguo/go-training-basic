// 归并排序：
// 	1. 递归: 对分区间A[:N/2], A[N/2:]分别递归使用归并排序，得到有序序列sortedLeft, sortedRight
//		递归退出条件，仅有单个元素，直接返回有序
// 	2. 合并: merge(l,r)，合并两个有序列表：
// 		申请两个游标i,j分别指示两个列表头部元素，依次迭代比较i,j, 退出条件未i或者j已经抵达列表末尾
//
package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"runtime/pprof"
)

// 归并驱动函数
func mergeSort(A []int, N int) []int {
	// 1. 递归退出条件
	if len(A) <= 1 {
		return A
	}
	// 2. 分治，分成两个列表(左和右列表),分别针对左列表和右列表进行分治排序，得到有序数组
	sortedLeft := mergeSort(A[:N/2], len(A[:N/2]))
	sortedRight := mergeSort(A[N/2:], len(A[N/2:]))

	// 3. 归并，合并两个有序数组
	return merge(sortedLeft, sortedRight)
}

// 归并有序列表函数
func merge(x, y []int) []int {
	lenx, leny := len(x), len(y)
	z := make([]int, lenx+leny) // 此处可以优化，可以尝试将z数组放入入参，避免每次make申请
	var i, j int
	for i < lenx && j < leny {
		if x[i] < y[j] {
			z[i+j] = x[i]
			i++
		} else {
			z[i+j] = y[j]
			j++
		}
	}
	if len(x[i:]) > 0 {
		copy(z[i+j:], x[i:])
	} else {
		copy(z[i+j:], y[j:])
	}
	return z
}

var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")
var memprofile = flag.String("memprofile", "", "write memory profile to this file")

func main() {
	// cpu
	flag.Parse()
	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal(err)
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}

	list := []int{1, 3, 4, 5, 8, 2, 6, 7, 9}
	fmt.Println(list)
	fmt.Println(mergeSort(list, len(list)))

	// memory
	if *memprofile != "" {
		f, err := os.Create(*memprofile)
		if err != nil {
			log.Fatal(err)
		}
		pprof.WriteHeapProfile(f)
		f.Close()
		return
	}
}
