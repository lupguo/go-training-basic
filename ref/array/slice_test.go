package array_test

import (
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

// Append的扩展速度
func TestSlice(t *testing.T) {
	var a []int
	cpm := make(map[int]int)
	for i := 0; i < 1e6; i++ {
		a = append(a, i*2)
		if _, ok := cpm[cap(a)]; !ok {
			cpm[cap(a)] = cap(a)
			t.Logf("len(a)=%d, cap(a)=%d, size=2^%f", len(a), cap(a), math.Log2(float64(cap(a))))
		}
	}
}

func TestSubSlice(t *testing.T) {
	a := [5]int{1, 2, 3, 4, 5}
	sa := a[:]
	s := a[1:4]
	t.Logf("&a[0](%p):%v", &a[0], a[0])
	t.Logf("&a[1](%p):%v", &a[1], a[1])
	t.Logf("&a(%p):%v", &a, a)
	t.Logf("a[:], a(%p):%v", sa, sa)
	t.Logf("a[1:4], a(%p):%v", s, s)
	// s[low:high], from s[low] iterate high-low elems
	t.Logf("a[2:3]: %v", a[2:3]) // 从下标为2的元素(a[2])开始，包含1个元素(3-2)，因此为[3]
	t.Logf("a[1:4]: %v", a[1:4]) // 从下班为1的元素(a[1])开始，包含3个元素(4-1)，因此为[2,3,4]
	// t.Log(a[1:6]) // panic，下标越界(从a[1]开始，但a[6]不存在)
	t.Logf("a[5:]:%v", a[5:]) //
	t.Logf("a[5]:%v", a[5:])  //
	t.Log(a[1:5][2:][2:])     //
}

func TestLowHigh(t *testing.T) {
	s := []int{1, 2, 3, 4, 5}
	t.Logf("s[1:3]=%v, len=%d, cap=%d", s[1:3], len(s[1:3]), cap(s[1:3]))         // 包含从s[1]开始，长度为3-1两个元素[2,3]，cap没有声明，则cap(s[1:])
	t.Logf("s[1:3:5]=%v, len=%d, cap=%d", s[1:3:5], len(s[1:3:5]), cap(s[1:3:5])) // len为3-1=2, cap为5-1=4
}

func TestEqual(t *testing.T) {
	s := [...]int{1, 2, 3, 4, 5} //注意s变量是数组，需要通过取地址符取其地址
	s1 := s[:]
	s2 := s1
	s3 := s[:4]
	t.Logf("address s(%p), s1(%p), s2(%p), s3(%p)", &s, s1, s2, s3)
	//if s1 == s2 { // compile error
	//}
	// 基于reflect.DeepEqual所得
	assert.Equal(t, s1, s2, "slice s1 should equal s2")
	assert.NotEqual(t, s1, s3, "s1 not equal s3")
}
