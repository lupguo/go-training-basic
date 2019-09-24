package main

func main() {

}

// Receive can only receive by no Pointer/Interface type
type P *int
type Op uintptr

func (pp Op) Opp(i int)  {

}
func (pp P) PP(i int)  {

}

type K interface{}

func (kk K) kk(i int)  {

}
type Ok struct{}




