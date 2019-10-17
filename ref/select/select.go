package main

import (
	"log"
	"time"
)

func main() {
	tk1 := time.Tick(1 * time.Second)
	tk2 := time.Tick(1 * time.Second)

	for {
		select {
		case t := <-tk1:
			log.Println("tk1 =>", t)
		case t := <-tk2:
			log.Println("\ttk2 >>", t)
		}
	}
}
