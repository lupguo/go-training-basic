package main

import (
	"github.com/gopherjs/gopherjs/js"
)

func main() {
	js.Global.Call("console.log", "Hello, Gopher Js!")
}