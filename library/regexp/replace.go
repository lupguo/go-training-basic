package main

import (
	"fmt"
	"regexp"
)

func main() {
	fmt.Println("ReplaceAllLiteralString >> ")
	re := regexp.MustCompile(`a(x*)b`)
	fmt.Println(re.ReplaceAllLiteralString("-ab-axxb-", "T"))
	fmt.Println(re.ReplaceAllLiteralString("-ab-axxb-", "$1"))
	fmt.Println(re.ReplaceAllLiteralString("-ab-axxb-", "${1}"))

	fmt.Println("ReplaceAllString >> ")
	fmt.Println(re.ReplaceAllString("-ab-axxb-", "T"))
	fmt.Println(re.ReplaceAllString("-ab-axxb-", "$1"))
	fmt.Println(re.ReplaceAllString("-ab-axxb-", "$1W")) // $1W not found
	fmt.Println(re.ReplaceAllString("-ab-axxb-", "${1}W"))
}