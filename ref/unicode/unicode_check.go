package main

import (
	"fmt"
	"unicode"
)

func main() {
	checkRune('龍')
	checkRune('𓀄')
}

func checkRune(r rune) {
	var msg string
	switch {
	default:
		msg = "no check..."
	case unicode.IsGraphic(r):
		msg = "is graphic"
	case unicode.IsDigit(r):
		msg = "is digit"
	case unicode.IsLetter(r):
		msg = "is letter"
	case unicode.IsPrint(r):
		msg = "is print"
	case unicode.IsSymbol(r):
		msg = "is symbol"

	}

	switch {
	case unicode.Is(unicode.Han, r):
		msg += " is hanyu"
	}

	fmt.Printf("%#U %s\n", r, msg)
}
