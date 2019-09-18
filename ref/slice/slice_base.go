package main

import (
	"fmt"
	"math/rand"
	"os"
	"sort"
	"text/tabwriter"
)

func main() {
	s := rand.Perm(10)
	sort.Ints(s)

	sc := make([]int, 10)
	cpnums := copy(sc, s[3:])
	fmt.Println("Copy from s to sc:", cpnums)
	copy(sc, sc[1:7])

	copy(s, s[1:5])
	// change
	ss := [][]int{
		s,
		s[0:],
		s[0:1],
		s[0:2],
		s[:2],
		s[1:],
		s[1:2],
		s[2:2],
		s[5:],
		s[:8],
		sc,
	}

	// s
	tw := tabwriter.NewWriter(os.Stdout, 4,4,2,' ', tabwriter.TabIndent)
	defer tw.Flush()
	fmt.Fprintf(tw, "Len(s)\tCap(s)\tVals\n")

	// show
	for _, v := range ss {
		fmt.Fprintf(tw, "%d\t%d\t%d\n", len(v), cap(v), v)
	}
}
