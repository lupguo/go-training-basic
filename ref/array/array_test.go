package array_test

import (
	"fmt"
	"testing"
)

// 数组比较
// 1. 长度不相同，则异常；
// 2. 若长度相同，元素不同，两个数组不相等
// 3. 若长度相同，元素相同，但排序不一致，两个数组也不相等
// 4. 仅两个数组长度一致、元素一致、且元素排序也相同则两个数组相等
func TestArray(t *testing.T) {
	a1 := [...]int{1, 2}
	a2 := [...]int{2, 1}
	a3 := [...]int{1, 3}
	b := [...]int{1, 2, 3}
	c := [...]int{1, 2, 3}
	fmt.Printf("%v==%v result: %v\n", a1, a2, a1 == a2)
	fmt.Printf("%v==%v result: %v\n", a1, a3, a1 == a3)
	//fmt.Printf("a2==b result: %v\n", a2 == b) // 编译错误，两个数组长度不相等
	fmt.Printf("%v==%v result: %v\n", b, c, b == c)
}

// 多维数组，内部维数需要确定，仅最后一维可以利用...替换
func TestBidArray(t *testing.T) {
	a := [...][3]int{{1, 2}, {3, 4}, {2, 3, 5}}
	for i, b := range a {
		t.Logf("i=%d, b=%v", i, b)
		for j, b := range b {
			t.Logf("j=%d, b=%v", j, b)
		}
	}
}

func TestThreeBidArray(t *testing.T) {
	a := [...][3][4]int{{{1, 2, 3}, {3, 4}, {6, 7}}, {{1, 3, 4, 6}}}
	for i, x := range a {
		t.Logf("i=%d, x=%v", i, x)
		for j, y := range x {
			t.Logf("\tj=%d, y=%d", j, y)
			for k, z := range y {
				t.Logf("\t\tk=%d, z=%d", k, z)
			}
		}
	}
}

func sum(ops ...int) int {
	ret := 0
	for _, op := range ops {
		ret += op
	}
	return ret
}

func TestSum(t *testing.T) {
	t.Log(sum(1,2,3))
	t.Log(sum(1,2,3,4))
}
