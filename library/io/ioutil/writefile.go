package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

func main() {
	// create temp dir
	dir, err := ioutil.TempDir("", "example")
	if err != nil {
		log.Fatal(err)
	}
	// clean temp dir
	defer os.RemoveAll(dir) // clean up

	// write to tmp file
	tmpfn := filepath.Join(dir, "tmpfile")
	content := []byte("temporary file's content")
	if err := ioutil.WriteFile(tmpfn, content, 0666); err != nil {
		log.Fatal(err)
	}

	// read tmp file
	fc, err := ioutil.ReadFile(tmpfn)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", fc)

}
