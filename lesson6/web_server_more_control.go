package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	mux := http.NewServeMux()
	addr := "localhost:8080"

	mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		log.Printf("Handling request at %s", req.URL.Path)
		fmt.Fprintf(w, "Hello from %s\n", req.URL.Path)
	})
	server := &http.Server{
		Addr:           addr,
		Handler:        mux,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	log.Printf("Running server at %s", addr)
	server.ListenAndServe()
}
