package main

import (
	"fmt"
	_ "go-example/ref/initfunc/pkg1"
	"go-example/ref/initfunc/pkg2"
)

func init() {
	fmt.Println("execute main init")
}

func main() {
	pkg2.Show()
	fmt.Println("execute main() function")
}
