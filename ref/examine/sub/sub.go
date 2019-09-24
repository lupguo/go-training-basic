package sub

import "fmt"

func init() {
	fmt.Println("sub init")
}

func CallSubFunc() {
	fmt.Println("sub function call")
}

func Add(l ...int) int {
	var s int
	for _, v := range l {
		s += v
	}
	return s
}

// 9
func SwitchData(i int) {
	switch i {
	case 1:
	case 2, 3, 4:
		fallthrough
	case 5 / 8:
	}
	switch {
	case i > 1:
	case i > 5 && i < 3:
	case i == 2 || i == 4:
	default:
	}
}

// 15.

type MyInt int

func (a MyInt) Add1(b MyInt) MyInt {
	return a+b
}

func (a *MyInt) Add2(b MyInt) MyInt {
	return *a+b
}
func (a *MyInt) Add3(b *MyInt) MyInt {
	return *a+*b
}
func (a MyInt) Add4(b *MyInt) MyInt {
	return a+*b
}

