package main

import (
	"fmt"
	"strconv"
)

func main() {

	var v int64 = 127

	fmt.Println(strconv.FormatInt(v, 2))
	fmt.Printf("%08b\t", v)
	fmt.Printf("%08X\t", v)
	fmt.Printf("%04X\t\n", v)

	s := strconv.QuoteToASCII(`"Fran & 
Freddie's Diner	☺"`) // there is a tab character inside the string literal
	fmt.Println(s)
	s2 := strconv.Quote(`"Fran & 
Freddie's Diner	☺"`) // there is a tab character inside the string literal
	fmt.Println(s2)
	s3 := strconv.QuoteToGraphic("\u263a") // there is a tab character inside the string literal
	fmt.Println(s3)
}
