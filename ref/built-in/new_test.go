package built_in

import (
	"fmt"
	"testing"
)

func TestNew(t *testing.T) {
	type S struct {
		a int;
		b float64
	}
	p := new(S)
	fmt.Printf("%p, %v", p, &p)
}
