package main

import "fmt"

func main() {
	var m0 map[string]bool
	fmt.Printf("%#v, len=%d\n", m0, len(m0))

	m1 := make(map[string]bool)
	fmt.Printf("%#v, len=%d\n", m1, len(m1))

	//m2:= make(map[string]bool) // error, can not compare map with ==
	//fmt.Printf("m1=m2, result=%t", m1==m2)
	fmt.Printf("unkey value: %#v\n", m0["x"])

	v, ok := m0["x"]
	fmt.Printf("v=%v,ok=%v\n", v, ok)

	// map中key值不存在，默认为其零值
	m3 := make(map[int]struct{ num int })
	v3, ok3 := m3[10]
	fmt.Printf("v3=%v,ok3=%v\n", v3, ok3)

}
