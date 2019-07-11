package main

import (
	"io/ioutil"
	"log"
	"os"
)

func main() {
	for i:=0; i<5; i++ {
		tmpfile, err := ioutil.TempFile("", "example.*.txt")
		if err != nil {
			log.Println(err)
			continue
		}
		log.Println("create temp file: "+tmpfile.Name())

		defer func(no int, filename string) {
			log.Printf("try to remove #%d => %s\n", no, filename)
			if err := os.Remove(filename); err != nil {
				log.Println(err)
			}
		}(i, tmpfile.Name())
	}
}