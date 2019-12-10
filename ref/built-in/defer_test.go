package built_in

import (
	"testing"
)

// Test Exit
// ...
func TestDefer(t *testing.T) {
	defer A(t)
	defer func() {
		t.Log("Test Exit")
	}()
}

// X
// Exec A++
// Exec A
func A(t *testing.T) {
	defer t.Log("Exec A")
	defer func() {
		t.Log("Exec A++")
	}()
	t.Log("X")
}
