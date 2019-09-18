package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	var urls = []string{
		"http://www.golang.org/",
		"http://www.google.com/",
		"http://www.somestupidname.com/",
	}
	for _, url := range urls {
		// Increment the WaitGroup counter.
		wg.Add(1)
		// Launch a goroutine to fetch the URL.
		go func(url string) {
			// Decrement the counter when the goroutine completes.
			defer wg.Done()
			// Fetch the URL.
			handleResponse(http.Get(url))
		}(url)
	}
	// Wait for all HTTP fetches to complete.
	wg.Wait()
}

func handleResponse(resp *http.Response, err error)  {
	r := resp.Request
	fmt.Printf("Host:%s, Uri:%s, Status Code:%d, Body Size:%d\n", r.Host, r.RequestURI, resp.StatusCode, resp.ContentLength)
	fmt.Printf("%#v\n",r)
	io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close()
	fmt.Printf("%#v\n",resp)
}