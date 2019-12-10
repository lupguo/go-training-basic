package interface_var

import (
	"testing"
)

type Coder interface{}

type GoProgrammer struct {
	Lang string
}
type CProgrammer struct {
	Lang string
}

func TestVar(t *testing.T) {
	var prog Coder
	prog = GoProgrammer{"Go"}
	prog = CProgrammer{"C"}
	switch v := prog.(type) {
	case GoProgrammer:
		t.Logf("type(%T)%#[1]v, pointer(%[1]p)", &prog)
		t.Logf("value(%v)", prog)
		t.Log(v.Lang)
	case CProgrammer:
		t.Logf("type(%T)%#[1]v, pointer(%[1]s)", prog)
		t.Logf("value(%v)", prog)
		t.Log(v.Lang)
	}
}
