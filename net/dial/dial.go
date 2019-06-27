package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"time"
)

func main() {
	c, err := net.Dial("tcp", ":2000")
	if err != nil {
		log.Fatalln(err)
	}
	defer c.Close()

	fmt.Printf("using network %s, local address %s, connect to %s...\n", c.RemoteAddr().Network(), c.LocalAddr().String(), c.RemoteAddr().String())



	// read bg
	go func(c net.Conn) {
		for {
			wn, err := io.Copy(os.Stdout, c)
			if err != nil {
				log.Fatalln(err)
			}
			fmt.Printf("\tread from server total size: %d\n", wn)
			time.Sleep(500 * time.Millisecond)
		}
	}(c)

	reader := bufio.NewReader(os.Stdin)
	// write
	for {
		rb, err := reader.ReadBytes('\n')
		if err != nil {
			log.Println(err)
			continue
		}
		wn, err := c.Write(rb)
		if err != nil {
			log.Println(err)
			continue
		}
		log.Printf("client write to conn total size: %d\n", wn)
	}
}
