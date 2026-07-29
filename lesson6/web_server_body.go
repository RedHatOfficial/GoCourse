package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/hello", func(w http.ResponseWriter, req *http.Request) {
		defer req.Body.Close()
		fmt.Println("The body is:")
		io.Copy(os.Stdout, req.Body)
	})
	http.ListenAndServe("localhost:8080", nil)
	if err := http.ListenAndServe("localhost:8080", nil); err != nil {
		log.Fatal(err)
	}
}
