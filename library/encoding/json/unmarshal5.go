package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	b := []byte(`{"Name":"Wednesday", "school": {"small":"hansan", "middle":"tangtian", "high":"jiuzhong"},"Age":6,"Parents":["Gomez","Morticia"]}`)
	var f interface{}
	if err := json.Unmarshal(b, &f); err != nil {
		panic(err)
	}

	fmt.Printf("%#v\n\n", f)

	m := f.(map[string]interface{})
	for k, v := range m {
		switch vv := v.(type) {
		case string:
			fmt.Println(k, "is string", vv)
		case float64:
			fmt.Println(k, "is float64", vv)
		case []interface{}:
			fmt.Println(k, "is an array:")
			for i, u := range vv {
				fmt.Println(i, u)
			}
		case map[string]interface{}:
			fmt.Println(k, "is an map:")
			for key,val := range vv {
				fmt.Println(key, "=>", val)
			}
		default:
			fmt.Println(k, "is of a type I don't know how to handle")
		}
	}

}
