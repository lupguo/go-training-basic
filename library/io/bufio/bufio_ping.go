package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
)

func main() {
	pingCmd := exec.Command("ping", "www.tkstorm.com")
	pingStdout, err := pingCmd.StdoutPipe()
	if err != nil {
		log.Fatalln(err)
	}
	if err := pingCmd.Start(); err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(pingStdout)
	for scanner.Scan() {
		log.Println(scanner.Text()) // Println will add back the final '\n'
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading ping stdout:", err)
	}
}