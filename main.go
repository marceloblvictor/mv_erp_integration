package main

import (
	"log"
	"net/http"
)

func main() {
	log.Fatal("test")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, World!"))
	})
}

// Start the HTTP server
