package main

import (
	"fmt"
	"log"
)

func main() {
	var input1, input2 string
	for {
		n, err := fmt.Scan(&input1, &input2)
		if err != nil {
			log.Println(err)
		}
		fmt.Printf("n=%d, input1=%s, input2=%s\n", n, input1, input2)
	}
}
