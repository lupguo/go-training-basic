package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	var jsonBlob = []byte(`{"name":"UserA","nick":"NickA","age":13,"likes":["ball","skating"],
"friends":[{"F1":"FName1"},{"F2":"FName2"}],"dead":false}`)
	//vv := make(map[string]interface{})
	var vv interface{}
	err := json.Unmarshal(jsonBlob, &vv)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%#v\n", vv)
}
