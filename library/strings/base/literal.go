package main

import (
	"go-example/helper"
)

func main() {
	strs := []string{
		`Hel,好,<br>\b\n`,
		`Hel,好,<br>
[]
`,
		"Hel,好,<br>\b\n|[]",
	}

	tw, show := helper.TwStdoutShow()
	defer tw.Flush()

	show("Index", "String")
	for i, v := range strs {
		show(i, v)
	}
}
