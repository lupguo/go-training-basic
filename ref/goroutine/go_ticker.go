package main

import (
	"log"
	"time"
)

func main() {
	timeCh := make(chan time.Time)
	// fin
	fin := make(chan struct{})

	// create time tick
	// timer chan =event=> fin chan =event=> ticker chan
	tk := time.NewTicker(1 * time.Second)
	tm := time.NewTimer(5 * time.Second)
	go func() {
		for {
			select {
			case t := <-tk.C:
				timeCh <- t
			//case <-tm.C:
			//	log.Println("receive timer over event, close time chan and stop time ticker chan.")
			//	close(timeCh)
			//	tk.Stop()
			//}
			case <-tm.C:
				log.Println("receive timer over event, send an fin event.")
				fin <- struct{}{}
				close(fin)
			}
		}
	}()

	// when can not add to for select case
	go func() {
		if _, ok := <-fin; ok {
			log.Println("receive fin, do other finish clean work!")
			tk.Stop()
			close(timeCh)
		}
	}()

	log.Println("Time show:")
	for t := range timeCh {
		log.Printf("%s\n", t.Format("2006/01/02 15:04:05"))
	}
}
