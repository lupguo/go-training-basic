package main

import "fmt"

func main() {
	s1 := make([]int, 10, 100)      // slice with len(s) == 10, cap(s) == 100
	s2 := make([]int, 1e3)          // slice with len(s) == cap(s) == 1000
	//s3 := make([]int, 1<<63)        // illegal: len(s) is not representable by a value of type int
	//s4 := make([]int, 10, 0)        // illegal: len(s) > cap(s)
	s3 := make([]int, 1<<5-1)        // illegal: len(s) is not representable by a value of type int
	s4 := make([]int, 10, 10)        // illegal: len(s) > cap(s)
	c5 := make(chan int, 10)        // channel with a buffer size of 10
	m6 := make(map[string]int, 100) // map with initial space for approximately 100 elements

	fmt.Println(s1, s2, s3, s4, c5, m6)
}
