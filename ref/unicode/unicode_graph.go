package main

import (
	"fmt"
	"unicode"
)

func main() {
	for n, i := 0, 0; i < 1e6; i++ {
		r := rune(i)
		if ! unicode.IsGraphic(r) || unicode.Is(unicode.Han, r) {
			continue
		}
		fmt.Printf("%#U ", r)
		if i % 0x8 == 0 {
			fmt.Println()
		}
		if n++; n>50000 {
			return
		}
	}
}
