package main

import (
	"fmt"
	"go-example/helper"
	"math"
	"unicode"
)

func main() {
	vals := []interface{}{
		0007,
		0x61,
		'a',
		`abc`,
		"abc",
		[]byte(`abc`),
		127,
		64,
		127 & 64,
		127 | 64,
		127 ^ 64,
		64 + 1,
		64 + 2,
		64 + 3,
		1 << 0,
		1 << 1,
		1 << 2,
		1 << 3,
		1 << 4,
		1 << 5,
		10 << 0,
		10 << 1,
		10 << 2,
		10 << 3,
		10 << 4,
		10 << 5,
		100 << 0,
		100 << 1,
		100 << 2,
		100 << 3,
		100 << 4,
		100 << 5,
		1 >> 0,
		1 >> 1,
		1 >> 2,
		1 >> 3,
		1 >> 4,
		1 >> 5,
		3 + 4i,
		3.5 + 4.2i,
		`\u2680`,
		`\U2691`,
		'\u2691',
		'\U00002691',
		03233,
		0xe411,
		1<<31 - 1,
		math.MaxUint32,
	}

	tw, show := helper.TwStdoutShow()
	show("index", "type", "val", "binary")
	for i, v := range vals {
		switch x := v.(type) {
		case string:
			v = []byte(x)
		}
		show(i, fmt.Sprintf("%T", v), fmt.Sprintf("%#v", v), fmt.Sprintf("%016b", v))
	}
	tw.Flush()

	// rune print
	fmt.Println("rune print...")
	for _, v := range vals {
		switch x := v.(type) {
		case rune:
			fmt.Printf("%q\n", rune(x))
		case int:
			if unicode.IsLetter(rune(x)) {
				fmt.Printf("%q\n", rune(x))
			}
		}
	}
}
