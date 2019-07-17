package main

import (
	"fmt"
	"unicode/utf8"
)

func main()  {
	const nihongo = "民国10年\xe2\x8c\x98"
	for i, w := 0, 0; i < len(nihongo); i += w {
		runeValue, width := utf8.DecodeRuneInString(nihongo[i:])
		fmt.Printf("%#U starts at byte position %d, witdh=%d\n", runeValue, i, width)
		w = width
	}

	const charPiece1 = "\xe2\x8c\x98"
	const charPiece2 = `\xe2\x8c\x98`
	fmt.Println(charPiece1, charPiece2)


	const zhwm = "民国10年\xe2\x8c\x98"
	for index, runeValue := range zhwm {
		fmt.Printf("%d, %U : %#[2]U \n", index,  runeValue)
	}

	fmt.Println("\xe2\x8c\x98")

	const placeOfInterest = `⌘`

	fmt.Printf("plain string: ")
	fmt.Printf("%s", placeOfInterest)
	fmt.Printf("\n")

	fmt.Printf("quoted string: ")
	fmt.Printf("%q", placeOfInterest)
	fmt.Printf("\n")

	fmt.Printf("hex bytes: ")
	for i := 0; i < len(placeOfInterest); i++ {
		fmt.Printf("%x ", placeOfInterest[i])
	}
	fmt.Printf("\n")


	fmt.Printf("\n\n")

	const sample = "\xbd\xb2\x3d\xbc\x20\xe2\x8c\x98"

	for i := 0; i < len(sample); i++ {
		fmt.Printf("%x ", sample[i])
	}
	fmt.Println()

	fmt.Printf("% x\n", sample)

	fmt.Printf("%q\n", sample)

	fmt.Printf("%+q\n", sample)

	sampleSlice := []byte("\xbd\xb2=\xbc ⌘")
	fmt.Printf("% x\n", sampleSlice)
	fmt.Printf("%q\n", sampleSlice)
	fmt.Printf("%+q\n", sampleSlice)

}
