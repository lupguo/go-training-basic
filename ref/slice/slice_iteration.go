package main

import "fmt"

func main() {
	ps := []interface{}{1,2,3,4,5}
	fn := func(p ...interface{}) {
		fmt.Printf("%T, %[1]v\n", p)
		for k, v := range p {
			fmt.Println(k,"=>",v)
		}
	}
	fn(ps...)
}
