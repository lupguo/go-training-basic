package string

import (
	"testing"
)

func TestString(t *testing.T) {
	s := "Hey man, 你好!"
	for i := 0; i < len(s); i++ {
		t.Logf("s[%d]=%c", i, s[i])
	}
	for i, c := range s {
		t.Logf("s[%d]=%c", i, c)
	}
	t.Logf("%s%s", s[9:12], s[12:15])
}

func TestStringBitMove(t *testing.T) {
	c := '中'
	s := "hello"
	t.Logf("%c, %[1]v, %016[1]b, %v, %016[2]b", c<<1, c)
	t.Logf("%c, %[1]v, %0[1]X", c-1)
	t.Logf("%c, %[1]v, %0[1]X, len(c)=%d", c, len(s))
	t.Log(s + "world")
}
