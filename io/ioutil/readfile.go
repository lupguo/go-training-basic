package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	content, err := ioutil.ReadFile("testdata/config/config.yml")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("File contents: %s", content)

}