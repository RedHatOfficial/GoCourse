package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/hello", func(w http.ResponseWriter, req *http.Request) {
		name, ok := req.URL.Query()["name"]

		if ok && len(name[0]) > 0 {
			fmt.Fprintf(w, "Hello %s\n", name[0])
		} else {
			fmt.Fprintf(w, "Say your name")
		}
	})
	http.ListenAndServe("localhost:8080", nil)
}
