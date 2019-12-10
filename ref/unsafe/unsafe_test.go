package unsafe

import (
	"testing"
	"unsafe"
)

func TestErrorConvert(t *testing.T) {
	i := 10
	t.Logf("%#v", unsafe.Pointer(&i))
	f := (*float64)(unsafe.Pointer(&i))
	t.Logf("%#v", f)
	t.Logf("%T, %v", *f, *f)
}

type MyInt int

func TestAliasConvert(t *testing.T) {
	a := []int{1, 2, 3}
	ma := *(*[]MyInt)(unsafe.Pointer(&a))
	t.Logf("%T, %[1]v", ma)
}
