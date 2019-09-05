package main

import "fmt"

//A
type A struct {
}

func (a A) show() {
	fmt.Println("A")
}

func (a *A) pshow() {
	fmt.Println("*A")
}

//B
type B struct {
	A
}

//func (b B) show() {
//	fmt.Println("B")
//}
//
//func (b *B) pshow() {
//	fmt.Println("*B")
//}

//C
type C struct {
	*A
}

//func (c C) show() {
//	fmt.Println("C")
//}
//func (c *C) pshow() {
//	fmt.Println("*C")
//}

func main() {
	a, b, c := A{}, B{}, C{}
	pa, pb, pc := new(A), new(B), new(C)

	fmt.Println("show::")
	a.show()
	b.show() //b是结构体变量，B类似是包含A结构体，非指针引用，可以调用接收器为A或者*A类型的接收器
	//c.show() c是结构体变量，但C类型是包含对A结构体的指针引用，无法直接调用接收器为A的引用
	a.pshow()
	b.pshow() //b是结构体变量，B类似是包含A结构体，非指针引用，可以调用接收器为A或者*A类型的接收器
	c.pshow()

	fmt.Println("pointer show::")
	pa.show()
	pb.show()  //pb是结构体指针变量，B类似是包含A结构体，非指针引用，可以调用接收器为A或者*A类型的接收器
	//pc.show()
	pa.pshow()
	pb.pshow() //pb是结构体指针变量，B类似是包含A结构体，非指针引用，可以调用接收器为A或者*A类型的接收器
	pc.pshow()
}
