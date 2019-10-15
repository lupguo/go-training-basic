package main

import (
	"fmt"
	"os"
	"text/tabwriter"
)

func main() {
	tw := tabwriter.NewWriter(os.Stdout, 4, 4, 4, ' ', 0)
	defer tw.Flush()
	fmt.Fprintln(tw, "ID\tVar1\tVar2\tJudge\tReal Rs")

	type eqs struct {
		id     int
		v1, v2 interface{}
		jug    bool
	}

	// interface
	a1 := new(int)
	a2 := a1  //a1 == a2
	a3 := &a1 //a1 != a3

	// slice
	s1 := []int{1, 2, 3}
	s2 := s1
	_ = s2

	// map
	map1 := make(map[string]bool)
	map2 := make(map[string]bool)
	_, _ = map1, map2

	// values judge
	vals := []eqs{
		{1, 97, 'a', false},
		{2, new(int), new(int), false},
		//{3,[]int{1}, []int{1}, false}, //output panic
		//{4,s1, s2, false}, // panic
		{5, a1, a2, true},
		{6, a1, a3, false},
		{61, a1, *a3, true},
		{7, &s1, &s1[0], true}, // :(, judge is error
		//{8, map1, map2, false}, // panic: runtime error: comparing uncomparable type map[string]bool
	}

	// loops
	for _, q := range vals {
		fmt.Fprintf(tw, "#%d\t%#v\t%v\t%v\t%v\n", q.id, q.v1, q.v2, q.jug, q.v1 == q.v2)
	}
}
