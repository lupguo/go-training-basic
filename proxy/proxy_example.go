package main

import (
	"context"
	"fmt"
	"golang.org/x/net/proxy"
	"io/ioutil"
	"log"
	"net"
	"net/http"
)

var (
	ADDR = "127.0.0.1:10443"
	URL  = "https://google.com"
)

func main() {
	// socks5 dialer
	dialer, err := proxy.SOCKS5("tcp", ADDR, nil, proxy.Direct)
	if err != nil {
		log.Fatalln("can't connect to the proxy:", err)
	}

	// make socks5 as the dialer
	httpTransport := &http.Transport{}
	httpTransport.DialContext = func(ctx context.Context, network, addr string) (conn net.Conn, e error) {
		return dialer.Dial(network, addr)
	}

	// create http client
	httpProxyClient := &http.Client{
		Transport: httpTransport,
	}

	// issue http request by socks5
	resp, err := httpProxyClient.Get(URL)
	if err != nil {
		log.Fatalln("http proxy client get fail:", err)
	}
	defer resp.Body.Close()

	// ok, print content
	byteData, _ := ioutil.ReadAll(resp.Body)
	fmt.Printf("%s", byteData)
}