package main

import "fmt"

func main() {
	s1 := "test龍𓀀𓀁𓀂𓀃𓀄"
	i := len(s1)
	fmt.Println(i)
	fmt.Printf("%#v\n", s1)
	for k, v := range s1 {
		fmt.Printf("%d=>%c, %[2]d, %[2]U, %T\n", k, v, v)
	}

	fmt.Println(&s1) // &s1[0] 非法

	s2 := []byte(s1)
	fmt.Printf("%p, %p\n", &s2, &s2[0])
}
