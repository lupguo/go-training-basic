package main

import (
	"log"
	"time"
)

func main() {
	over := time.After(3 * time.Second)
	done := false
	i := 0
	defer log.Println("start...")
	for !done {
		select {
		case <-over:
			log.Println("over")
			done = true
		default:
			defer func() {
				log.Println("in defer #", i)
				i++
			}()
			time.Sleep(time.Second)
		}
	}
}
