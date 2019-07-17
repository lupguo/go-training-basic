// https://golang.org/ref/spec#Type_assertions
package main

import (
	"io"
	"log"
)

func main() {
	var x interface{} = 7 // x has dynamic type int and value 7
	i := x.(int)          // i has type int and value 7
	log.Printf("x.(int)=%d\n", i)
	if j, ok := x.(string); ok {
		log.Println(j)
	} else {
		log.Printf("x.(string)=%t", ok)
	}

	// I
	s := new(Si)
	f(s)
}

type I interface {
	m()
}

type Si struct{}

func (s *Si) m() {}
func (s *Si) Read(p []byte) (n int, err error)  {
	return
}

func f(y I) {
	s := y.(*Si)       // illegal: string does not implement I (missing method m)
	r := y.(io.Reader) // r has type io.Reader and the dynamic type of y must implement both I and io.Reader
	log.Printf("%T, %T\n",s, r)
}
