package main

import (
	"io/ioutil"
	"log"
	"os"
)

func main() {
	content := []byte("temporary file's content")
	// generate temp file and open it
	tmpfile, err := ioutil.TempFile("", "example")
	if err != nil {
		log.Fatal(err)
	}
	// defer clean up, if defer not set input param, when tmpfile set to nil, using tmpfile.Name() will got fatal
	defer func(name string) {
		log.Println("clean up, remove temp file.")
		if err := os.Remove(name); err != nil {
			log.Fatal(err)
		}
	}(tmpfile.Name())

	// write
	log.Println("write to file " + tmpfile.Name())
	if _, err := tmpfile.Write(content); err != nil {
		log.Fatal(err)
	}
	// close
	if err := tmpfile.Close(); err != nil {
		log.Fatal(err)
	}

	// update tmpfile, just for debug
	tmpfile = nil
}
