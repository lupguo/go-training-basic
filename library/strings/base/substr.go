package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s := "Hello, プロWorldグラム"
	for i, r := range s {
		fmt.Println(i, r)
	}

	// substr 1
	fmt.Println(`substr 1 >>`)
	for i := 7; i < len(s); i++ {
		fmt.Println(s[:i], "...")
	}

	// substr 1.5
	fmt.Println(`substr 1.5 >>`)
	for i := 7; i < len(s); {
		j := i + 1
		for ; ; j++ {
			if utf8.Valid([]byte(s[i:j])) {
				fmt.Println(s[i:j])
				break
			}
		}
		i = j
	}

	// substr 2
	fmt.Println(`substr 2 >> `)
	for i := 7; i < len(s); {
		j := i + 1
		for ; !utf8.Valid([]byte(s[i:j])); j++ {
			continue
		}
		i = j
		fmt.Println(s[:i], "...")
	}

	// substr 3
	fmt.Println(`substr 3 >> `)
	for i, _ := range s {
		if i > 7 {
			fmt.Println(s[:i], "...")
		}
	}

	// substr func
	fmt.Println(`substr func >>`)
	substr := func(str string, len int, tail string) string {
		rstr := []rune(str)
		return fmt.Sprintf("%s%s", string(rstr[:len]), tail)
	}
	for i:=7; i<len([]rune(s)); i++ {
		fmt.Println(substr(s, i, ".."))
	}
}
