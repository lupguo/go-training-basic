package main

import "fmt"

type student struct {
	Name string
	Age  int
}

func pase_student() {
	m := make(map[string]*student)
	stus := []student{
		{Name: "zhou", Age: 24},
		{Name: "li", Age: 23},
		{Name: "wang", Age: 22},
	}
	for k, stu := range stus {
		m[stu.Name] = &stus[k]
	}

	fmt.Printf("%+v\n", m)
	for k, s := range m {
		fmt.Printf("name:%s, stu:%+v\n", k, s)
	}
}

func main() {
	pase_student()
}
