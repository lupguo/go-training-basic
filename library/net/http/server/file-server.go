package main

import (
	"log"
	"net/http"
)

func main() {
	// Simple static webserver:
	log.Fatal(http.ListenAndServe(":8085", http.FileServer(http.Dir("/usr/share/doc"))))
}