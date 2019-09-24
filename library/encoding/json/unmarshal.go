package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	type Animal struct {
		Name  string
		Order string
		desc  string
	}
	var jsonBlob = []byte(`[
	{"Name": "Platypus", "Order": "Monotremata"},
	{"Name": "Quoll",    "Order": "Dasyuromorphia"}
]`)
	var animals []Animal
	err := json.Unmarshal(jsonBlob, &animals)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Printf("%#v", animals)
}
