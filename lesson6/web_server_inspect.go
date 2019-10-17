package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		requestDump, err := httputil.DumpRequest(req, true)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(string(requestDump))

	})
	http.ListenAndServe("localhost:8080" , nil)
}
