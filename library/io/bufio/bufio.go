package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		txt := scanner.Text()
		if strings.ToLower(txt) == "quit" {
			break
		} else {
			fmt.Println(txt) // Println will add back the final '\n'
		}
	}

	fmt.Println("Go on...")
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
}
