package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "Hello from level 0: %s\n", req.URL.Path)
	})
	http.HandleFunc("/one", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "Hello from level 1: %s\n", req.URL.Path)
	})
	http.HandleFunc("/one/two", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "Hello from level 2: %s\n", req.URL.Path)
	})
	http.ListenAndServe("localhost:8080" , nil)
}
