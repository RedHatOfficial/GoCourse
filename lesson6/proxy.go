package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func main() {
	target := url.URL{
		Host:   "localhost:8080",
		Scheme: "http",
	}
	http.Handle("/", httputil.NewSingleHostReverseProxy(&target))
	log.Printf("Proxy started")
	http.ListenAndServe("localhost:8081", nil)
}
