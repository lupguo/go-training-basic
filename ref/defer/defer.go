package main

import "fmt"

func main() {
	// prints 3 2 1 0 before surrounding function returns
	for i := 0; i <= 3; i++ {
		defer fmt.Print(i)
	}

	// print 42
	fmt.Print(f())

	// print 12384
	s := NewSlice()
	defer s.Add(1).Add(2).Add(3).Add(4)
	s.Add(8)

	// running defer function first, then panic function
	panic("Panic Opoos!")
}

// f returns 42
func f() (result int) {
	defer func() {
		// result is accessed after it was set to 6 by the return statement
		result *= 7
	}()
	return 6
}

type Slice []int

func NewSlice() Slice {
	return make(Slice, 0)
}

func (s *Slice) Add(elem int) *Slice {
	*s = append(*s, elem)
	fmt.Print(elem)
	return s
}
