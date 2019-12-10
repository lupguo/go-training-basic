package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	//recv()
	randBinary()
}

func randBinary() {
	c := make(chan int)
	i := 0
	go func() {
		for n := range c {
			fmt.Printf("%d", n)
			time.Sleep(50 * time.Millisecond)
			if i++; i%50 == 0 {
				fmt.Printf("\n")
			}
		}
	}()
	for { // randBinary random sequence of bits to c
		select {
		case c <- 0: // note: no statement, no fallthrough, no folding of cases
		case c <- 1:
		default:
		}
	}
}

func recv() {
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
