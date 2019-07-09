package main

import (
	"fmt"
)

type PeopleA interface {
	Speak(string) string
}

type Stduent struct{}

// 尝试将Stduent转为非指针接收器，对比差异
func (stu *Stduent) Speak(think string) (talk string) {
	if think == "bitch" {
		talk = "You are a good boy"
	} else {
		talk = "hi"
	}
	return
}

func main() {
	var peo PeopleA = &Stduent{} // ok, 支持隐式转换，
	//var peo PeopleA = Stduent{} // bad, 因为接收者stu需要指针类型，但传递了结构体对象
	fmt.Println(peo.Speak("bitch"))
}
