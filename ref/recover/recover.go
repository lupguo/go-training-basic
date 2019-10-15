package main

import "fmt"

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("main recover:", r)
		}
	}()

	fmt.Println("do main goroutine")

	fmt.Println(funOne())

	panic("panic hello")
}

func funOne() int {
	defer func() int {
		fmt.Println("	fun one recover:",recover())
		return 60
	}()
	defer func() int {
		fmt.Println("	func one defer again")
		return 80
	}()
	fmt.Println("	do fun one, tobe panic")
	panic("fun one panic")
	return 100
}
