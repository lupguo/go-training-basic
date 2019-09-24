package main

import "fmt"

func main() {
	list := NewIntList([]int{1, 2, 3, 4, 5})
	fmt.Println(list.Sum())
}

type IntList struct {
	Value int
	Tail  *IntList
}

func NewIntList(num []int) *IntList {
	list := new(IntList)
	list.init(num)
	return list
}

func (list *IntList) add(elem int, pos *IntList) {
	last := list
	if last == nil {
		last = &IntList{
			Value: elem,
		}
	}
	for last.Tail != nil {
		last = last.Tail
	}
}

func (list *IntList) init(nums []int) {
	for _, v := range nums {
		list.add(v, list)
	}
	fmt.Println(list)
}

// Sum returns the sum of the list elements.
func (list *IntList) Sum() int {
	if list == nil {
		return 0
	}
	return list.Value + list.Tail.Sum()
}
