package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	var input1, input2 string
	n, err := fmt.Fscanln(os.Stdin, &input1, &input2)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("fscanln >> n=%d, input1=%s, input2=%s\n", n, input1, input2)
}
