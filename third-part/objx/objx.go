// https://github.com/stretchr/objx
package main

import (
	"fmt"
	"github.com/stretchr/objx"
)

func main() {
	// Use MustFromJSON to make an objx.Map from some JSON
	m := objx.MustFromJSON(`{"name": "Mat", "age": 30, "male":true, "height": 1.85}`)

	// Get the details
	name := m.Get("name").Str()
	age := m.Get("age").Int()
	male := m.Get("male").Bool(false)
	height := m.Get("height").Float32(0)
	hobby := m.Get("hobby").Str()

	fmt.Printf("name:%s,age(%d),male(%t),height(%0.2f),hobby:%s\n", name, age, male, height, hobby)

	for key, val := range m {
		fmt.Printf("key=%s, val=%v\n", key, val)
	}
}
