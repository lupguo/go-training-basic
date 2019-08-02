package main

import (
	"encoding/json"
	"fmt"
	"log"
)

func main() {
	type Message struct {
		Name string
		Body string
		Time int64 // default value
	}

	b := []byte(`{"Name":"Bob","Food":"Pickle"}`)
	var m Message
	if err := json.Unmarshal(b, &m); err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("%+v", m)

}
