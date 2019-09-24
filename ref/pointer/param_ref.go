package main

import "fmt"

type Slice []int

type Map map[string]int

type Struct struct {
	id   int
	name string
}

func main() {
	myslice := Slice{1, 2, 3, 4, 5}
	fmt.Printf("%T->%[1]v->%[1]p->%[1]x, %T->%[2]v->%[2]p\n", &myslice, myslice)
	show(&myslice, myslice)

	mymap := Map{"id": 1, "score": 100}
	fmt.Printf("%T->%[1]v->%[1]p->%[1]x,\n%T->%[2]v->%[2]p->%[2]x,\n %T->%[3]v->%[3]p\n", &mymap, *(&mymap), mymap)

	mystruct := Struct{
		id:   80,
		name: "ck",
	}
	fmt.Printf("%T->%[1]v->%[1]p->%[1]x,%T->%[2]v->%[2]x\n", &mystruct, mystruct)
	show(&mystruct, mystruct)
}

func show(s1 interface{}, s2 interface{}) {
	var v interface{}
	switch k := s1.(type) {
	case *Slice:
		v = k
		//v = *(s1.(*Slice))
	case *Map:
		v = k
		//v = *(s1.(*Map))
	case *Struct:
		v = k
		//v = *(s1.(*Struct))
	}
	fmt.Printf("FUNCALL V: %T->%[1]v->%#v\n", &v, v)
	fmt.Printf("FUNCALL: %T->%[1]v->%[1]p->%[1]x, %T->%[2]v->%[2]p->%[2]x, %T->%[3]v->%[3]p\n", s1, v, s2)
}
