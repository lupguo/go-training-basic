package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strconv"
)

func main() {
	// Listen on TCP port 2000 on all available unicast and
	// anycast IP addresses of the local system.
	listen, err := net.Listen("tcp", ":2000")
	if err != nil {
		log.Fatal(err)
	}
	defer listen.Close()

	// Print listen message
	addr := listen.Addr()
	fmt.Printf("network:%s, listen on:%s...\n", addr.Network(), addr.String())

	// Wait for a connection.
	for {
		conn, err := listen.Accept()
		if err != nil {
			log.Fatal(err)
		}
		// Handle the connection in a new goroutine.
		// The loop then returns to accepting, so that
		// multiple connections may be served concurrently.
		go func(c net.Conn) {
			defer c.Close()

			// Client Coming
			log.Printf("client %s coming..", c.RemoteAddr().String())

			// Echo all incoming data.
			reader := bufio.NewReader(c)
			writer := bufio.NewWriter(c)
			for {
				rb, err := reader.ReadBytes('\n')
				if err != nil {
					log.Println("reader.ReadBytes() error,", err)
					return
				}
				log.Printf("Client say:%s\n", strconv.Quote(string(rb)))

				// Write to Client
				wb, err := writer.WriteString(fmt.Sprintf("Server answer:%s", rb))
				if err != nil {
					log.Println(err)
					return
				}

				// Last log
				log.Printf("server write to conn size: %d bytes\n", wb)
			}
		}(conn)
	}
}

