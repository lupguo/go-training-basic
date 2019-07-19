package main

import (
	"fmt"
	"regexp"
)

func main() {
	str := "aa/bb/cc"
	re:= regexp.MustCompile("[^/]*")
	fmt.Printf("%#v", re.FindAllString(str, -1))
}
