package main

import "fmt"

func main() {
	l1 := []int{1, 2, 3}
	l2 := []int{4, 5}
	fmt.Println(l1, l2)
	fmt.Printf("%T, %[1]p,%[1]v ; %T, %[2]p,%[2]v\n", l1, l2)
	l1, l2 = l2, l1
	fmt.Println(l1, l2)
	fmt.Printf("%T, %[1]p,%[1]v ; %T, %[2]p,%[2]v\n", l1, l2)
}
