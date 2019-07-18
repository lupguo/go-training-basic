package main

import "fmt"

func main() {
	ucs := []rune{
		'\u2680',
		'\u2681',
		'\u2682',
		'\u2683',
		'\u2684',
		'\u2685',
	}

	fmt.Println("%q")
	for _, u := range ucs {
		fmt.Printf("%q ", u)
	}
	fmt.Println("\n%#q")
	for _, u := range ucs {
		fmt.Printf("%#q ", u)
	}
	fmt.Println("\n%+q")
	for _, u := range ucs {
		fmt.Printf("%+q ", u)
	}
	fmt.Println("\n%#U")
	for _, u := range ucs {
		fmt.Printf("%#U ", u)
	}
}
