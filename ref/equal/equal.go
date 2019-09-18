package main

import (
	"fmt"
	"os"
	"text/tabwriter"
)

func main() {
	tw := tabwriter.NewWriter(os.Stdout, 4, 4, 4, ' ', 0)
	defer tw.Flush()
	fmt.Fprintln(tw, "Var1\tVar2\tJudge\tRunning Equal Rs")

	type eqs struct {
		v1, v2 interface{}
		jug    bool
	}

	// interface
	a1 := new(int)
	a2 := a1
	a3 := &a1

	// s1
	s1 := []int{1}
	s2 := s1
	_ = s2

	// values add
	vals := []eqs{
		{97, 'a', false},
		{new(int), new(int), false},
		//{[]int{1}, []int{1}, false}, //output panic
		//{s1, s2, false}, // panic
		{a1, a2, true},
		{a1, a3, false},
		{&s1, &s1[0],true}, //
	}

	// loops
	for _, q := range vals {
		fmt.Fprintf(tw, "%v\t%v\t%v\t%v\n", q.v1, q.v2, q.jug, q.v1 == q.v2)
	}
}
