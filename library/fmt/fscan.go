package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	var input1, input2 string
	for {
		n, err := fmt.Fscan(os.Stdin, &input1, &input2)
		if err != nil {
			log.Println(err)
		}
		fmt.Printf("fscan >> n=%d, input1=%s, input2=%s\n", n, input1, input2)
	}
}
