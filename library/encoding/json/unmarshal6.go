package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	b := []byte(`{"Name":"Wednesday","Age":6,"Parents":["Gomez","Morticia"]}`)

	// Unmarshal在幕后分配了一个新的切片, 这是Unmarshal如何使用支持的引用类型（指针，切片和映射）的典型特征。
	type FamilyMember struct {
		Name    string
		Age     int
		Parents []string
	}

	var m FamilyMember
	fmt.Printf("%#v\n",m)
	if err := json.Unmarshal(b, &m); err != nil {
		panic(err)
	}
	fmt.Printf("%#v\n",m)

	// 指针引用
	type Bar int
	type Foo struct {
		Bar *Bar
	}
	var mc Foo
	fmt.Printf("%#v\n",mc)
	c := []byte(`{"Bar": 100, "Foo":20}`)
	if err := json.Unmarshal(c, &mc); err != nil {
		panic(err)
	}
	fmt.Printf("%#v, %#v\n",mc, *mc.Bar)
}
