package main

import (
	"fmt"
	"regexp"
)

func main() {
	str := "aa/bb@cc|dd"
	re := regexp.MustCompile(`[/@|]|(@[a-z]+)|_([a-z]+)`)
	fmt.Printf("NumSubexp Num : %d\n", re.NumSubexp())
	fmt.Printf("%#v\n", re.Split(str, -1))
	fmt.Printf("%#v\n", re.Split(str, 1))
	fmt.Printf("%#v\n", re.Split(str, 2))
	fmt.Printf("%#v\n", re.Split(str, 3))
	fmt.Printf("%#v\n", re.Split(str, 4))
}
