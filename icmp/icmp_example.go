package icmp

import (
	"fmt"
	"golang.org/x/net/icmp"
	"golang.org/x/net/ipv4"
	"log"
	"net"
	"os"
	"time"
)

func Send() {
	ipConn, err := icmp.ListenPacket("udp4", "0.0.0.0")
	if err != nil {
		log.Fatal(err, "icmp.ListenPacket")
	}
	defer ipConn.Close()

	// deadline
	err = ipConn.SetDeadline(time.Now().Add(1 * time.Second))
	if err != nil {
		log.Fatal(err, "ipConn.SetDeadline")
	}

	// message
	sendMsg := icmp.Message{
		Type: ipv4.ICMPTypeEcho,
		Code: 0,
		Body: &icmp.Echo{
			ID:   os.Getpid() & 0xffff,
			Seq:  0,
			Data: []byte("PING-PONG"),
		},
	}
	wb, err := sendMsg.Marshal(nil)
	if err != nil {
		log.Fatal(err, "sendMsg.Marshal")
	}

	// dns
	addrs, err := net.LookupHost("tkstorm.com")
	if err != nil {
		log.Println(err, "net.LookupHost")
		return
	}
	fmt.Println(addrs)

	// send message
	start := time.Now()
	if _, err := ipConn.WriteTo(wb, &net.UDPAddr{IP: net.ParseIP(addrs[0])}); err != nil {
		log.Println(err, "ipConn.WriteTo")
		return
	}

	// read response
	rb := make([]byte, 1500)
	n, peer, err := ipConn.ReadFrom(rb)
	if err != nil {
		log.Println(err, "ipConn.ReadFrom")
		return
	}

	// parse buffer
	recMsg, err := icmp.ParseMessage(1, rb[:n])
	if err != nil {
		log.Fatalln(err, "icmp.ParseMessage")
	}

	elapse := time.Since(start).Nanoseconds() / 1e6

	switch recMsg.Type {
	case ipv4.ICMPTypeEchoReply: //icmp is echo type
		fmt.Printf("ICMP SEND: %+v\n", sendMsg)
		fmt.Printf("ICMP RECEIVE FROM %s, %s: %+v\n", peer.String(), peer.Network(), recMsg)
		fmt.Printf("ELPASE TIME: %dms\n", elapse)
	default:
		fmt.Printf("got %+v; want echo reply", recMsg)
	}
}
