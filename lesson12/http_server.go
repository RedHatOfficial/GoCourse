package main

import (
	"fmt"
	"net/http"
)

func main() {
	const address = ":8080"

	const directory = http.Dir(".")

	fmt.Println("Starting HTTP server on address", address)
	err := http.ListenAndServe(address, http.FileServer(directory))

	if err != nil {
		fmt.Println("Failed to start server", err)
		return
	}
}
