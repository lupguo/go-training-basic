package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"runtime/pprof"
)

// 归并驱动函数
func msort(A []int, N int) []int {
	// 1. 递归退出条件
	if len(A) <= 1 {
		return A
	}
	// 2. 分治，分成两个列表(左和右列表),分别针对左列表和右列表进行分治排序，得到有序数组
	left := msort(A[:N/2], len(A[:N/2]))
	right := msort(A[N/2:], len(A[N/2:]))

	// 3. 归并，合并两个有序数组
	return mergeSortedList(left, right)
}

// 归并有序列表函数
func mergeSortedList(x, y []int) []int {
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
	fmt.Println(msort(list, len(list)))

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
