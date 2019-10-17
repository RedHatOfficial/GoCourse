package main

import (
	"fmt"
	"net/http"
	"encoding/json"
)

func main() {
	http.HandleFunc("/hello", func(w http.ResponseWriter, req *http.Request) {
		var params map[string]interface{}
		err := json.NewDecoder(req.Body).Decode(&params)
		if err != nil {
			fmt.Printf("Error parsing body: %s", err)
			return
		}
		fmt.Printf("The body is: %#v", params)
	})
	http.ListenAndServe("localhost:8080", nil)
}
