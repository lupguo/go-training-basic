package main

import (
	"log"
	"math/rand"
	"time"
)

func main() {
	// switch default
	rand.Seed(time.Now().UnixNano())
	tag := rand.Intn(10)
	switch tag {
	default:
		log.Println("default...")
	case 0, 1, 2, 3:
		log.Println("0~3")
	case 4, 5, 6, 7:
		log.Println("4~7")
	}

	// switch simple stmt
	absX := func() int {
		var x int
		defer func() {
			log.Println("rand x=", x)
		}()
		f := func() int {
			return rand.Intn(10) - rand.Intn(10)
		}
		switch x = f(); { // missing switch expression means "true"
		case x < 0:
			return -x
		default:
			return x
		}
	}
	log.Println("absX()=", absX())

	// switch true
	x, y, z := rand.Intn(10), rand.Intn(10), rand.Intn(10)
	switch {
	default:
		log.Println("not match!!")
	case x < y:
		log.Printf("x=%d,y=%d, x<y\n", x ,y)
	case x < z:
		log.Printf("x=%d,z=%d, x<z\n", x ,z)
	case x == 4:
		log.Println("x==4")
	}
}
