package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	// If the file doesn't exist, create it, or append to the file
	f, err := os.OpenFile("/tmp/error.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0755)
	if err != nil {
		log.Fatal(err)
	}
	if _, err := f.Write([]byte("appended some data aadddddd\n")); err != nil {
		log.Fatal(err)
	}
	if err := f.Close(); err != nil {
		log.Fatal(err)
	}

	// 当前执行文件
	fmt.Println(os.Executable())

	// 当前pid，ppid
	fmt.Println("pid=", os.Getpid(), "ppid=", os.Getppid())
}
