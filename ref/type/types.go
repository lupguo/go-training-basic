package main

import "fmt"

type (
	A1 = string
	A2 = A1
)

type (
	B1 string
	B2 B1
	B3 []B1
	B4 B3
)

func (s B1) String() string {
	return fmt.Sprintf("B1 type =>%v", string(s))
}

type BInf interface {
}

func main() {
	// interface define
	vars := []BInf{
		B1("hello"),
		100,
		"hi",
		1.86,
	}

	// type switch
	_ = func(v BInf) {
		switch vv := v.(type) {
		case int:
			fmt.Println("int type:", vv)
		case B1:
			fmt.Println("B1 type:", vv)
		default:
			fmt.Println("unknown type", vv)
		}
	}

	for _, v := range vars {
		fmt.Printf("%p\n", &v)
	}
}
