package main

import (
	"log"
	"runtime"
	"time"
)

// Send the sequence 2, 3, 4, … to channel 'ch'.
func generate(ch chan<- int) {
	for i := 2; i<100 ; i++ {
		//log.Printf("generate ch postion: %p\n", ch)
		ch <- i // Send 'i' to channel 'ch'.
	}
	close(ch)
}

// Copy the values from channel 'src' to channel 'dst',
// removing those divisible by 'prime'.
func filter(src <-chan int, dst chan<- int, prime int) {
	j := 10000
	for i := range src { // Loop over values received from 'src'.
		//log.Printf("filter(%d) src => %p, dst => %p\n", j, src, dst)
		if i%prime != 0 {
			dst <- i // Send 'i' to channel 'dst'.
		}
		j++
	}
	close(dst)
}

// The prime sieve: Daisy-chain filter processes together.
func sieve() {
	ch := make(chan int) // Create a new channel.
	//log.Printf("sieve ch postion: %p\n", ch)
	go generate(ch) // Start generate() as a subprocess.
	for i := 0; ; i++ {
		prime, ok := <-ch
		if !ok {
			return
		}
		log.Println("prime =>", prime)
		ch1 := make(chan int)
		go filter(ch, ch1, prime)
		//log.Printf("sieve (%d) ch => %p, ch1 => %p\n", i, ch, ch1)
		ch = ch1
		time.Sleep(100 * time.Millisecond)
		log.Println("--------------")
	}
}

func status() {
	for {
		log.Printf("NumGoroutine=%d\n", runtime.NumGoroutine())
		time.Sleep(1000 * time.Millisecond)
	}
}

func init() {
	go status()
}

func main() {
	sieve()
}
