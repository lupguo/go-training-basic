package main

import "fmt"

func main() {
	s1 := make([]int, 0, 20)
	s1 = append(s1, []int{1,2,3,4,5}...)
	fmt.Printf("len(s1)=%d, cap(s1)=%d\n", len(s1), cap(s1))
	fmt.Println(s1[len(s1):],s1[len(s1)-1:],s1[len(s1)-5:],)
	fmt.Println(s1)
}
