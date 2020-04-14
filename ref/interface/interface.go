package main

import (
	"bytes"
	"io"
)

const debug = true

func main() {

	var buf *bytes.Buffer
	if debug {
		buf = new(bytes.Buffer) // enable collection of output
		for i := 0; i < 1e6; i++ {
			buf.WriteByte(byte(i))
		}
	}
	f(buf) // NOTE: subtly incorrect!
	if debug {
		// ...use buf...
	}
}

// If out is non-nil, output will be written to it.
func f(out io.Writer) {
	// ...do something...
	if out != nil {
		out.Write([]byte("done!\n"))
	}
}
