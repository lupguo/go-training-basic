package main

import (
	"fmt"
	"log"
)

func main() {
	var input string
	n, err := fmt.Scanln(&input)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("n=%d, input=%s\n", n, input)
}
