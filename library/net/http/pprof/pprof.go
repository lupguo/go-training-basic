package main

import (
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
	"strconv"
	"time"
)

func main() {
	// pprof goroutine
	go func() {
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()

	// 1. http.HandleFunc
	http.HandleFunc("/fib", func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		n, _ := strconv.Atoi(r.FormValue("n"))
		fmt.Fprintf(w, "fib(%d)=%d", n, fib(n))
	})
	// 2. http server handler
	http.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "%s %s", r.RequestURI, time.Now())
	}))
	// 3. http.handlerFunc + http.Handle
	dfHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("dfHandler")
	})
	http.Handle("/dh", dfHandler)

	log.Fatal(http.ListenAndServe(":8703", nil))

	//for i := 0; i < 60; i++ {
	//	st := time.Now()
	//	f := fib(i)
	//	et := time.Since(st)
	//	fmt.Printf("fib(%d)=%d elapse: %d ms\n", i, f, et.Nanoseconds()/1e6)
	//}
}

func fib(n int) int {
	if n < 2 {
		return n
	}
	return fib(n-1) + fib(n-2)
}
