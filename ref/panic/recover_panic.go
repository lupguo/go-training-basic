// https://golang.org/ref/spec#Handling_panics
//
package main

import (
	"log"
)

type Exception struct {
	code int
	msg  string
}

func NewException(code int, msg string) *Exception {
	return &Exception{
		code,
		msg,
	}
}

func main() {
	protect(work)
}

func work() {
	log.Println("do work...")
	protect(subWork)
	log.Println("do work remainder...")
	panic(NewException(10, "work panic"))
}

func subWork()  {
	log.Println("\tdo sub work...")
	panic(NewException(100, "sub work panic"))
}

func protect(g func()) {
	defer func() {
		if x := recover(); x != nil {
			switch v := x.(type) {
			case *Exception:
				log.Printf("\tpanic recover! run time panic: code: %d, msg: %s", v.code, v.msg)
			default:
				log.Printf("\tpanic recover! run time panic: %v", v)
			}
		}
		log.Println("done") // Println executes normally even if there is a panic
	}()
	log.Println("protect start")
	g()
}
