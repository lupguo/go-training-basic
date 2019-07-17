package main

import (
	"fmt"
	"net"
)

func main() {
	ipv4Addr := net.ParseIP("192.168.1.1")

	fmt.Println(ipv4Addr.String())
	fmt.Println(ipv4Addr.To4().String())
	fmt.Println(ipv4Addr.To16().String())

	fmt.Println(ipv4Addr.DefaultMask())

	mrsh, _ := ipv4Addr.MarshalText()
	fmt.Printf("%v", mrsh)
}
