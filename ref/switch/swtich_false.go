package main

import (
	"fmt"
)

func main() {
	SwitchFalse()
	SwitchFalseFn()
}

func SwitchFalse() {
	switch false {
	case false:
		fmt.Println(false)
	case true:
		fmt.Println(true)
	}
}

func SwitchFalseFn() {
	switch f(); {
	case false:
		fmt.Println(false)
	case true:
		fmt.Println(true)
	}
}

func f() bool {
	return false
}
