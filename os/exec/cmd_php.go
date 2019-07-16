package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

func main() {
	cmd := exec.Command("php", "-v")
	cmd.Stdout = os.Stdout
	fmt.Printf("%+v\n", cmd)

	if err:= cmd.Run(); err != nil {
		log.Fatalln(err)
	}
	log.Println("exec finished.", )
}
