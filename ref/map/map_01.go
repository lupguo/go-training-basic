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
}
