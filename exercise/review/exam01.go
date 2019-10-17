package main

import (
	"errors"
	"fmt"
)

func add(args ...int) int {
	sum := 0
	for _, arg := range args {
		sum += arg
	}
	return sum
}

type Slice []int

var ERR_ELEM_NT_EXIST = errors.New("ELEM NOT EXIST")

func (s *Slice) Remove(value interface{}) error {
	for i, v := range *s {
		if isEqual(value, v) {
			*s = append((*s)[:i], (*s)[i+1:]...)
			return nil
		}
	}
	return ERR_ELEM_NT_EXIST
}

func isEqual(value interface{}, i int) bool {
	return i == value
}

func main() {
	add([]int{1, 2, 3}...)

	//var x interface{} = nil
	//var y *struct{ x int } = nil
	//s := make([]int, 0)

	var sl Slice = []int{0, 1, 2, 3, 4, 5}
	sl.Remove(3)
	fmt.Println(sl)

	//x := []int{1, 2, 3, 4, 5, 6}

	// method
	var a Integer = 1
	var b Integer = 2
	var i interface{} = &a
	sum := i.(*Integer).Add(b)
	fmt.Println(sum)

	var ia IA
	var ib IB
	var ic IC
	ia, ib = ib, ia

	ia = ic
	//ic = ia //fail ia 没有实现ic的接口的一些方法

	var bb bool
	if bb == false {
	}

	// chan
	var ch1 chan int
	ch1 = make(chan int, 10)
	close(ch1)
	//ch := make(chan int)
	//ch1 <- 1
	fmt.Println(<-ch1)
	//fmt.Println(ch1, ch)
	//ch <- 1

	// map cap()函数仅支持数组、slice、通道、指向数组的指针
	mysliceP := &sl
	fmt.Println(cap(*mysliceP))
	//mymap := make(map[string]int,10)
	//cap(mymap)

}

type Integer int

func (a *Integer) AddP(b Integer) Integer {
	return *a + b
}

//type Integer int
func (a Integer) Add(b Integer) Integer {
	return a + b
}

type IA interface {
	method()
}
type IB interface {
	method()
}
type IC interface {
	method()
	ic()
}
