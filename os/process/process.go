package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	childPs, err := os.StartProcess("ls", []string{"-al"}, nil)
	if err != nil {
		log.Fatalln(err)
	}

	stats, err := childPs.Wait()
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("%v", stats)
}
