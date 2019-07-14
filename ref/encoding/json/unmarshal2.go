package main

import (
	"encoding/json"
	"log"
)

func main() {
	var jsonBlob = []byte(`{"name":"UserA","nick":"NA","age":13,"likes":["ball","skating"],"friends":[{"F1":"FName1"},{"F2":"FName2"}],"dead":false}`)
	vv := make(map[string]interface{})
	err := json.Unmarshal(jsonBlob, &vv)
	if err != nil {
		log.Fatalln(err)
	}
	log.Printf("%+v\n", vv)
}
