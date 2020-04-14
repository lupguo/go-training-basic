package multi_assign

import (
	"testing"
)

func TestMultiAssign(t *testing.T) {
	var z = make([]int, 1)
	z = []int{1, 2, 3, 4, 5}
	t.Log(z)

	var y []int
	y = []int{1, 2, 3, 4}
	t.Log(y)

	var a []int = nil
	a = []int{0}
	a, a[0] = []int{1, 2}, 9
	t.Log(a)

	var s = make([]int, 10)
	s, s[0] = []int{1, 2}, 9
	t.Log(s)

	x := []int{123}
	x, x[0], x, x[0] = nil, 456, []int{7, 8, 9}, 233
	t.Log(x)
}
