package main

import (
	"fmt"
	"path/filepath"
)

func main()  {
	fmt.Println(filepath.EvalSymlinks("/tmp/static.php"))

	fmt.Println(filepath.FromSlash("/ab/../://.c/d"))

	fmt.Println(filepath.Glob("/data/www/learn/php/static/*.php"))

	fmt.Println(filepath.Abs("/a/../b"))
}