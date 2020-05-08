package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/hello", func(w http.ResponseWriter, req *http.Request) {
		contentType := req.Header.Get("Content-type")
		if contentType != "application/json" {
			http.Error(w, fmt.Sprintf("Unexpected content %s", contentType), http.StatusBadRequest)
		}
		// parse ...
	})
	http.ListenAndServe("localhost:8080", nil)
}
