package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// Params is a type that represents data to be send in the request
type Params struct {
	Hello string
}

func main() {
	http.HandleFunc("/hello", func(w http.ResponseWriter, req *http.Request) {
		var params Params
		defer req.Body.Close()
		err := json.NewDecoder(req.Body).Decode(&params)
		if err != nil {
			http.Error(w, fmt.Sprintf("Error parsing body: %s", err), http.StatusBadRequest)
			return
		}
		fmt.Fprintf(w, "Hello %s", params.Hello)
	})
	http.ListenAndServe("localhost:8080", nil)
	if err := http.ListenAndServe("localhost:8080", nil); err != nil {
		log.Fatal(err)
	}
}
