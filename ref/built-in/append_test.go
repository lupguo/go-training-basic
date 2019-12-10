package built_in

import (
	"fmt"
	"testing"
)

func TestAppend(t *testing.T)  {
	s0 := []int{0, 0}
	s1 := append(s0, 2)                // append a single element     s1 == []int{0, 0, 2}
	s2 := append(s1, 3, 5, 7)          // append multiple elements    s2 == []int{0, 0, 2, 3, 5, 7}
	s3 := append(s2, s0...)            // append a slice              s3 == []int{0, 0, 2, 3, 5, 7, 0, 0}
	s4 := append(s3[3:6], s3[2:]...)   // append overlapping slice    s4 == []int{3, 5, 7, 2, 3, 5, 7, 0, 0}

	fmt.Printf("s0(%p)=%[1]v\n", s0)
	fmt.Printf("s1(%p)=%[1]v\n", s1)
	fmt.Printf("s2(%p)=%[1]v\n", s2)
	fmt.Printf("s3(%p)=%[1]v\n", s3)
	fmt.Printf("s4(%p)=%[1]v\n", s4)
}
