package main

import "fmt"

// A
type A struct{}
func (a A) show() {
	fmt.Println("A call by show()")
}
func (a *A) ptShow() {
	fmt.Println("*A call by ptShow()")
}
// B
type B struct {
	A
}
// C
type C struct {
	*A
}

func main() {
	fmt.Println("test a, pa show:")
	a, pa := A{}, new(A)
	a.show()
	a.ptShow()
	pa.show()
	pa.ptShow()
	fmt.Printf("type of a(%T), pa(%T)\n", a, pa)

	//type B struct {
	//	A
	//}
	fmt.Println("\ntest b, pb show:")
	b, pb := B{}, new(B)
	b.show()
	b.ptShow()
	pb.show()
	pb.ptShow()

	//type C struct {
	//	*A
	//}
	fmt.Println("\ntest c, pc show:")
	c, pc := C{}, new(C)
	//c.show()	//panic c为C类型
	c.ptShow()
	//pc.show()	//panic pc为*C类型，
	pc.ptShow()
	fmt.Printf("type of c.A(%T), pc.A(%T)\n", c.A, pc.A)
}
