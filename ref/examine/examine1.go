package main

import (
	"fmt"
	"go-example/ref/examine/sub"
)

type exam struct {
	subId int
}

type empty struct{}

type errInf interface {
	show()
}

// 8 error const set to reference type
//const c1  =  errors.New("const 1")

// 9 init call order
func init() {
	fmt.Println("init 1")
}

func init() {
	fmt.Println("init 2")
	sub.CallSubFunc()
}

// 12. main can not set input variable and return values
func main() {
	// 1
	var p *int
	fmt.Println(p, p == nil)

	// 1.5
	s1, s2 := `ab`, `cd`
	cs := s1 + s2
	fmt.Println(cs)

	// 2. redefine p pointer
	p = new(int)
	*p = 1
	e := new(exam)
	ept := new(empty)
	fmt.Println(p, *p, p == nil, e, e == nil, ept, ept == nil)

	// 3. struct pointer
	p2 := &struct {
		name string
	}{"t"}
	// (&p).name is error
	fmt.Println(p2.name, &p2.name, (*p2).name, )

	// 4
	csa := `ab` + `cd`
	fmt.Println(csa, `ab`+`cd`)

	// 5 add fun call
	fmt.Println(sub.Add(1, 2), sub.Add(1, 2, 3), sub.Add([]int{1, 2, 3}...))

	// 6 type convert
	type oInt int
	type oInf interface{}
	num1 := 100
	var num2 oInf
	jnum := oInt(num1) // type convert
	num2 = jnum
	num3, ok := num2.(oInt) // only interface can assert type,
	fmt.Printf("%T-%[1]T,%T-%[2]T,%T-%[3]T,%v\n", num1, num2, num3, ok)

	// 7 bool check
	// true and false are the two untyped boolean values.
	var b bool
	b = true
	b = (1 == 2)
	b = bool(false)
	//b = 1
	//b = bool(1)
	_ = b

	// 10 nil setting
	// can not assign nil without explicit type
	// var x = nil //error
	// var x string = nil // string default is ""
	var x1 error = nil
	var x2 interface{} = x1
	_, _ = x1, x2

	// 11. slice
	//sl := make([]int) // miss len when make
	sl2 := make([]int, 0)
	fmt.Printf("%#v, len=%d, cap=%d\n", sl2, len(sl2), cap(sl2))

	// 13. slice remove elem
	ss := sub.Slice{0, 1, 2, 3, 4, 5}
	fmt.Println(ss)
	ssR, _ := ss.Remove(2)
	fmt.Println(ss, *ssR)

	// 14. slice define
	sl1 := []int{1, 2, 3}
	sl1 = []int{
		1, 2,
		3, 4,
	}
	//sl1 = []int{
	//	1, 2 // error: need trailing comma before newline
	//}
	_ = sl1

	// 15. check method call
	var ia, ib sub.MyInt = 1, 2
	var pa, pb *sub.MyInt
	pa, pb = &ia, &ib
	// input real param type must
	fmt.Println(ia.Add1(ib), ia.Add2(ib), ia.Add3(pb), ia.Add4(pb))
	fmt.Println(pa.Add1(ib), pa.Add2(ib), pa.Add3(pb), pa.Add4(pb))

}
