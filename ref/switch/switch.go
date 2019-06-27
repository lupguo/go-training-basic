package main

import (
	"fmt"
)

func main() {
	sw2()
}

// fallthrough +1024unit
func sw2() {
	for _, c := range "+1024k" {
		switch {
		case c == '+' || c == '-':
			fallthrough
		case c >= '0' && c <= '9':
			fmt.Printf("%c", c)
		case c == 'k' || c == 'M' || c == 'G':
			fmt.Printf("unit")
		default:
			fmt.Println("\tout range")
		}
	}
}

// comma, equal fallthrough +1024unit
func sw3() {
	for _, c := range "+1024k" {
		switch {
		case c == '+', c == '-', c >= '0' && c <= '9':
			fmt.Printf("%c", c)
		case c == 'k' || c == 'M' || c == 'G':
			fmt.Printf("unit")
		default:
			fmt.Println("\tout range")
		}
	}
}

func sw1() {
	var nums string
forloop:
	for _, c := range "+1024k" {
		fmt.Printf("%c, %[1]T\n", c)
		switch {
		case c == '+' || c == '-':
			fmt.Println("\tcompare symbol")
		case c >= '0' && c <= '9':
			nums += string(c)
			fmt.Println("\tnumber")
		case c == 'k' || c == 'M' || c == 'G':
			fmt.Println("\tsize unit")
		default:
			fmt.Println("\tout range")
			break forloop
		}
	}
	fmt.Println(nums)
}
