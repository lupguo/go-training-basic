package main

import "fmt"

func calc(index string, a, b int) int {
	ret := a + b
	fmt.Println(index, a, b, ret)
	return ret
}

func main() {
	a := 1
	b := 2
	// defer在定义时候，执行值计算:
	// 1. defer calc("1", 1, calc("10", 1, 2))
	// a. 打印 10 1 2 3
	// defer calc("1", 1, 3)
	defer calc("1", a, calc("10", a, b))

	a = 0
	// 2. defer calc("2", 0, calc("20", 0, 2))
	// b. 打印 20 0 2 2
	// defer calc("2", 0, 2)
	defer calc("2", a, calc("20", a, b))
	b = 1

	// 3. defer LIFO
	// c. 打印 calc("2", 0, 2): 2 0 2 2
	// d. 打印 calc("1", 1, 3): 1 1 3 4
}
