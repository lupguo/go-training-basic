package main

import (
	"github.com/pkg/browser"
	"go-example/helper"
	"os"
	"strings"
)

const download = helper.TestDataPath + "/html/download_new.html"

func main() {
	//ExampleOpenURL()
	//ExampleOpenFile()
	//ExampleOpenReader()
	ExampleOpenReaderFile()
}

func ExampleOpenFile() {
	browser.OpenFile(download)
}

func ExampleOpenReader() {
	// https://github.com/rust-lang/rust/issues/13871
	const quote = `There was a night when winds from unknown spaces
whirled us irresistibly into limitless vacum beyond all thought and entity.
Perceptions of the most maddeningly untransmissible sort thronged upon us;
perceptions of infinity which at the time convulsed us with joy, yet which
are now partly lost to my memory and partly incapable of presentation to others.`
	r := strings.NewReader(quote)
	browser.OpenReader(r)
}

func ExampleOpenReaderFile()  {
	file, _ := os.Open(download)
	browser.OpenReader(file)
}

func ExampleOpenURL() {
	const url = "http://golang.org/"
	browser.OpenURL(url)
}
