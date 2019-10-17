package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/hello", func(w http.ResponseWriter, req *http.Request) {
		fmt.Println("The body is:")
		io.Copy(os.Stdout, req.Body)
	})
	http.ListenAndServe("localhost:8080", nil)
}
