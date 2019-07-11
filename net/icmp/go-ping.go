package icmp

import (
	"fmt"
	"github.com/sparrc/go-ping"
)

func GopingExp()  {
	pinger, err := ping.NewPinger("www.google.com")
	if err != nil {
		panic(err)
	}
	pinger.Count = 3
	pinger.Run() // blocks until finished
	stats := pinger.Statistics() // get send/receive/rtt stats

	fmt.Printf("%#v", stats)
}
