package main

import "fmt"

func main() {
	s1 := []int{2: 10}
	fmt.Println(s1, len(s1), cap(s1))

	s2 := []int{3: 7, 9, 10: 100, 0: 99}
	fmt.Println(s2, len(s2), cap(s2))

	sx := []byte{1023: 'x'}
	sy := []byte{1023: 'y'}
	fmt.Println(sx, len(sx))
	fmt.Println(sy, len(sy))
}
