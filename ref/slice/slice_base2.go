package main

import "fmt"

func main() {

	// panic: runtime error: index out of range
	//s1 := make([]int, 0, 10)
	//for i:=10; i>0; i-- {
	//	s1[10-i] = i
	//}
	//fmt.Println(s1)

	//s1[0] = 1 // panic

	//s2 := make([]int) // error, slice must have length

	s := []int{1, 2, 3}
	ss := s[1:]
	fmt.Printf("%#p, %#p, %#p, %#p\n", ss, s, &ss[0], &s[0])

	for i := 100; i < 120; i++ {
		ss = append(ss, i)
		fmt.Printf("[cap(ss)=%d], %v\n", cap(ss), ss)
	}

	for k := range ss {
		ss[k] += 10
	}
	fmt.Println(ss)
}
