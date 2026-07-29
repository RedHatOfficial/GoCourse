package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/hello", func(w http.ResponseWriter, req *http.Request) {
		defer req.Body.Close()
		var params map[string]interface{}
		err := json.NewDecoder(req.Body).Decode(&params)
		if err != nil {
			fmt.Printf("Error parsing body: %s", err)
			http.Error(w, fmt.Sprintf("Error parsing body: %s", err), http.StatusBadRequest)
			return
		}
		fmt.Printf("The body is: %#v", params)
	})
	http.ListenAndServe("localhost:8080", nil)
	if err := http.ListenAndServe("localhost:8080", nil); err != nil {
		log.Fatal(err)
	}
}
