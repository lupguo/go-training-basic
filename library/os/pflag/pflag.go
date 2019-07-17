package main

import (
	"fmt"
	"github.com/spf13/pflag"
)

func main()  {
	name := "verbose"
	short := name[:1]

	fs := pflag.NewFlagSet("Example", pflag.ContinueOnError)
	fs.BoolP(name, short, false, "verbose output")

	// len(short) must be == 1
	flag := fs.ShorthandLookup(short)

	fmt.Println(flag.Name)
}
