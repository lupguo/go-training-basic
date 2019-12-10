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
