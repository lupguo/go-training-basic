package main

import (
	"fmt"
)

type PeopleB interface {
	Show()
}

type StudentB struct{}

func (stu *StudentB) Show() {

}

func live() PeopleB {
	var stu *StudentB // live()函数执行，初始化一个stu=nil的零值，返回Student的零值
	return stu
}

var in interface{}

func main() {
	//	fmt.Printf("%+v\n",nil != nil) 两个nil无法比较
	fmt.Println(in == nil) // 为true
	in = new(StudentB)
	fmt.Println(in == nil) // 为false

	if live() == nil { // live()返回一个Student指针类型零值，与nil不相等
		fmt.Println("AAAAAAA")
	} else { //
		fmt.Println("BBBBBBB")
	}
}
