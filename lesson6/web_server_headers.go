package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/hello", func(w http.ResponseWriter, req *http.Request) {
		content_type := req.Header.Get("Content-type")
		if content_type != "application/json" {
			http.Error(w, fmt.Sprintf("Unexpected content %s", content_type), http.StatusBadRequest)
		}
		// parse ...
	})
	http.ListenAndServe("localhost:8080", nil)
}
