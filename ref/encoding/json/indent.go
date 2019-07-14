package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"os"
)

func main() {
	type Road struct {
		Name   string
		Number int
	}
	roads := []*Road{
		{"Diamond Fork", 29},
		{"Sheep Creek", 51},
	}
	fmt.Printf("%#v\n", roads)

	b, err := json.Marshal(roads)
	if err != nil {
		log.Fatal(err)
	}

	var out bytes.Buffer
	json.Indent(&out, b, "", "  ")
	out.WriteTo(os.Stdout)
}
