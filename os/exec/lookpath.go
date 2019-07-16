package main

import (
	"fmt"
	"log"
	"os/exec"
)

func main() {
	path, err := exec.LookPath("fortune")
	if err != nil {
		log.Fatal("installing fortune is in your future; ", err)
	}
	fmt.Printf("fortune is available at %s\n", path)
}