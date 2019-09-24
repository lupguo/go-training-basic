package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	var jsonBlob = []byte(`{"name":"UserA","nick":"NickA","age":13,"likes":["ball","skating"],
"friends":[{"F1":"FName1"},{"F2":"FName2"}],"dead":false}`)
	//f := make(map[string]interface{})
	var f interface{}
	err := json.Unmarshal(jsonBlob, &f)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%#v\n", f)
	m := f.(map[string]interface{})
	for k, v := range m {
		switch vv := v.(type) {
		case bool:
			fmt.Println(k, "is bool", vv)
		case string:
			fmt.Println(k, "is string", vv)
		case float64:
			fmt.Println(k, "is float64", vv)
		case []interface{}:
			fmt.Println(k, "is an array:")
			for i, u := range vv {
				fmt.Println(i, u)
			}
		default:
			fmt.Println(k, "is of a type I don't know how to handle")
		}
	}
}
