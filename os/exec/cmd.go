// simulate cmd run `$ echo "some input"|tr a-z A-Z`, result store in bytes.Buffer
package main

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
	"strings"

)
func main() {
	cmd := exec.Command("tr", "a-z", "A-Z")
	cmd.Stdin = strings.NewReader("some input")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("in all caps: %q\n", out.String())
}