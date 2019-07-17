package main

import (
	"fmt"
	"os"
	"strconv"
	"text/tabwriter"
)

const (
	sizeB int = 1 << (10 * iota)
	sizeK
	sizeM
	sizeG
	sizeT
)

func main() {
	tw := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
	defer tw.Flush()
	fmt.Fprintf(tw, "String Size\t Int Size\n---\t---\n")
	showSize := func(strSize string) {
		fmt.Fprintf(tw, "%s\t%d\tbytes\n", strSize, strToIntSize(strSize))
	}
	strSizes := []string{
		"+1024",
		"+1024b",
		"+1024B",
		"+1024k",
		"+1024K",
		"+1024m",
		"+1024M",
		"+1024g",
		"+1024G",
		"+1024t",
		"+1024T",
		"10k",
		"10M",
	}

	for _, v := range strSizes {
		showSize(v)
	}
}

func strToIntSize(strSize string) int {
	var n, sizeUnit int
	var size []rune
	sizeUnit = sizeB
	for _, c := range strSize {
		switch {
		case c == '+', c == '-':
			continue
		case c >= '0' && c <= '9':
			size = append(size, c)
		case c == 'b', c == 'B':
			continue
		case c == 'k', c == 'K':
			sizeUnit = sizeK
		case c == 'm', c == 'M':
			sizeUnit = sizeM
		case c == 'g', c == 'G':
			sizeUnit = sizeG
		case c == 't', c == 'T':
			sizeUnit = sizeT
		}
	}

	// strconv to int size
	n, err := strconv.Atoi(string(size))
	if err != nil {
		panic(err)
	}

	return n * sizeUnit
}
