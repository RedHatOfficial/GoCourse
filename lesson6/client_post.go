package main

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
)

type Params struct {
	Hello string
}

func main() {
	buffer := new(bytes.Buffer)
	params := Params{Hello: "world"}
	json.NewEncoder(buffer).Encode(params)
	httpResp, err := http.Post("http://localhost:8080/hello", "application/json", buffer)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Status: %s", httpResp.Status)
	io.Copy(os.Stdout, httpResp.Body)
}
