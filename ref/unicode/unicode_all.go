package main

import (
	"fmt"
	"log"
	"unicode"
)

func main() {
	//printBA()
	printRange(0x13000, 0x13500)
}

func printRange(begin, end int) {
	if end - begin > 5000 {
		log.Fatalln("too many print")
	}
	for i := begin; i < end; i++ {
		if ! unicode.IsGraphic(rune(i)) {
			return
		}
		fmt.Printf("%#U ", i)
		if i%5 == 0 {
			fmt.Println()
		}
	}
}

func printBAKey(r rune) {
	for i := r - 100; i < r+100; i++ {
		fmt.Printf("%c", i)
	}
}

func printBA() {
	for i := '𓀁' - 10; i < '𓀁'+100; i++ {
		fmt.Printf("%#U ", i)
	}
}

func printU() {
	for i := 0; i < 80000; i = i + 5 {
		fmt.Printf("%c ", rune(i))
		if i%1000 == 0 {
			fmt.Println()
		}
	}
}
