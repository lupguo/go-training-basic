// output {"Foo":"Foo","FoO":"FoO","FOO":"FOO"}, struct field not export will be ignore
package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	v := struct {
		foo string
		foO string
		Foo string
		FoO string
		FOO string
	}{
		"foo",
		"foO",
		"Foo",
		"FoO",
		"FOO",
	}

	if d, err := json.Marshal(v); err != nil {
		panic(err)
	}else {
		fmt.Printf("%s\n", d)
	}
}
