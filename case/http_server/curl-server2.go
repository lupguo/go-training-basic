package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"
)

// 用于调试PHP curl_multi_select
func main() {
	http.HandleFunc("/curl/", func(w http.ResponseWriter, r *http.Request) {
		now := time.Now()
		time.Sleep(100*time.Millisecond)
		//time.Sleep(*time.Second)
		fmt.Fprintf(w,"URI:%s,TIME ESAPSED:%s\n", r.RequestURI, time.Since(now))
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		v, ok := w.(http.Flusher)
		if ok {
			log.Println("http.flusher")
		}
		//w.Header().Set("x-request-id", "ok")
		for i := 0; i < 10; i++ {
			log.Printf("URI:%s, RemoteAddr:%v\n", r.RequestURI, r.RemoteAddr)
			now := time.Now()
			time.Sleep(time.Duration(rand.Intn(10)+1) * 300 * time.Millisecond)
			fmt.Fprintf(w,"URI:%s,TIME ESAPSED:%s\n", r.RequestURI, time.Since(now))
			v.Flush()
		}
	})

	log.Fatal(http.ListenAndServe(":2334", nil))
}
