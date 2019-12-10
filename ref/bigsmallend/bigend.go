package main

import "fmt"

func main() {
	b := []byte("AB")
	fmt.Printf("&str=%p, &A=%p, &B=%p\n",b, &b[0], &b[1])
}
