// https://golang.org/ref/spec#Switch_statements
// The "fallthrough" statement is not permitted in a type switch.
package main

import (
	"log"
	"math/rand"
	"time"
)

func main() {
	v := []interface{}{
		1,
		"hello",
		1.24,
		func(int) float64 { return 2.00 },
		false,
		nil,
		map[string]int{},
		[]int{},
	}
	log.Printf("%#v\n", v)

	for _, v := range v {
		checkType(v)
	}

	// type test
	log.Println("type by if check >>")
	runTypeTest()
}

func checkType(x interface{}) {
	switch i := x.(type) {
	case nil:
		log.Println("x is nil") // type of i is type of x (interface{})
	case int:
		log.Printf("i type=%T, val=%[1]d", i) // type of i is int
	case float64:
		log.Printf("i type=%T, val=%0.2[1]f\n", i) // type of i is float64
	case func(int) float64:
		log.Printf("i type=%T, val=%[1]p\n", i) // type of i is func(int) float64
	case bool, string:
		log.Println("type is bool or string") // type of i is type of x (interface{})
	default:
		log.Println("don't know the type") // type of i is type of x (interface{})
	}
}

func runTypeTest() {
	rand.Seed(time.Now().UnixNano())
	var v interface{} // x is evaluated exactly once
	randVal := []interface{}{
		1,
		"hello",
		1.24,
		func(int) float64 { return 2.00 },
		false,
		nil,
		map[string]int{},
		[]int{},
	}
	v = randVal[rand.Intn(len(randVal))]

	if v == nil {
		i := v // type of i is type of x (interface{})
		log.Println("x is nil", i)
	} else if i, isInt := v.(int); isInt {
		log.Println(i) // type of i is int
	} else if i, isFloat64 := v.(float64); isFloat64 {
		log.Println(i) // type of i is float64
	} else if i, isFunc := v.(func(int) float64); isFunc {
		log.Println(i) // type of i is func(int) float64
	} else {
		_, isBool := v.(bool)
		_, isString := v.(string)
		if isBool || isString {
			i := v // type of i is type of x (interface{})
			log.Println("type is bool or string", i)
		} else {
			i := v // type of i is type of x (interface{})
			log.Println("don't know the type", i)
		}
	}
}
