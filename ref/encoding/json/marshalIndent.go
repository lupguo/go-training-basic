package main

import (
	"encoding/json"
	"fmt"
	"log"
)

func main() {
	data := map[string]int{
		"a": 1,
		"b": 2,
	}

	json, err := json.MarshalIndent(data, "", "")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(json))
}
