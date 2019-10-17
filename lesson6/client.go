package main

import (
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	httpResp, err := http.Get("http://localhost:8080/hello")
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Status: %s", httpResp.Status)
	io.Copy(os.Stdout, httpResp.Body)

}
