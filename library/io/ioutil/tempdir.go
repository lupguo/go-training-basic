package main

import (
	"io/ioutil"
	"log"
	"os"
	"path"
)

func main() {
	var tmpdir string
	for i := 0; i < 10; i++ {
		dir, err := ioutil.TempDir("", "tkstorm_")
		if err != nil {
			log.Println(err)
			continue
		}
		log.Printf("tempdir: %s\n", dir)
		// tmpdir set
		if tmpdir == "" {
			tmpdir = path.Dir(dir)
		}
	}

	// remove tmpdir
	if err := os.RemoveAll(tmpdir); err != nil {
		log.Printf("removeall %s fail: %s\n", tmpdir, err)
	}
}
