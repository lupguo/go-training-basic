package main

import "fmt"

func main() {
	s1 := make([]int, 0, 10)

	for i:=10; i>0; i-- {
		s1[10-i] = i
	}
	fmt.Println(s1)

	//s1[0] = 1 // panic

	//s2 := make([]int) // error
	//m1 := make(map[string]int)
}
