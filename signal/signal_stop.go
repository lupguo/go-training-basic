package main

import (
	"fmt"
	"os"
	"os/signal"
	"time"
)

func main() {
	// over
	over := make(chan struct{})

	// ticker
	tk := time.NewTicker(1 * time.Second)
	go func() {
		for nt := range tk.C {
			fmt.Println(nt.Format("2006/01/02 15:04:05"))
		}
	}()

	// listen for ctrl-C signal
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for _ = range c {
			tk.Stop()
			over <- struct{}{}
		}
	}()

	<-over
}
