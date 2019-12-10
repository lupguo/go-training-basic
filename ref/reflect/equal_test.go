package reflect

import (
	"reflect"
	"testing"
)

func TestDeepEqual(t *testing.T) {
	m1 := map[int]string{1: "one", 2: "two", 3: "three",}
	m2 := map[int]string{1: "one", 2: "two", 3: "three",}
	t.Logf("m1 equal m2, result %v", reflect.DeepEqual(m1, m2))

	s1 := []int{1, 2}
	s2 := []int{1, 2}
	s3 := []int{2, 1}
	t.Log(reflect.DeepEqual(s1, s2), reflect.DeepEqual(s1, s3))

	a1 := [...]int{1, 2}
	a2 := [...]int{1, 2}
	//a3 := [...]int{1, 2, 3}
	a4 := [...]int{2, 1}
	if a1 == a2 { //数组值可比较
		t.Logf("a1==a2: %v", a1 == a2)
	}
	if a1 == a4 {
		t.Logf("a1==a4: %v", a1 == a4)
	}
//	if a1 == a3 { // 类型不同无法匹配
//	}
}
