package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"
)

var reqId int

func init() {
	rand.Seed(time.Now().UnixNano())
}

const (
	listenAddr = ":2333"
)

// 用于调试PHP curl_multi_select
func main() {
	// srv serverMux
	srvMux := &MyHandler{srvMux: http.NewServeMux()}

	// web srv
	srv := &http.Server{
		Addr:              listenAddr,
		Handler:           srvMux,
		MaxHeaderBytes:    100,
		TLSNextProto:      nil,
		ConnState:         nil,
		ErrorLog:          nil,
	}

	// srv setting
	srv.SetKeepAlivesEnabled(true)

	// srv starting
	log.Fatal(srv.ListenAndServe())
}

type MyHandler struct {
	srvMux *http.ServeMux
}

func (m *MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	m.srvMux.HandleFunc("/", defHandler)
}

func defHandler(w http.ResponseWriter, r *http.Request) {
	//reqId++
	// 请求日志
	log.Printf("RemoteAddr:%s, RequestUri:%s\n",r.RemoteAddr, r.RequestURI)
	// 模拟耗时
	time.Sleep(time.Duration(rand.Intn(10)+1) / 10 * time.Second)
	fmt.Fprintf(w, "URI:%s, ReqId=%d", r.RequestURI, reqId)
}
