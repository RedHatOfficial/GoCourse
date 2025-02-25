package main

import (
	"fmt"
	"log"

	nats "github.com/nats-io/nats.go"
)

const Subject = "test1"

func main() {
	conn, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()
	fmt.Printf("Connected to %s\n", nats.DefaultURL)

	err2 := conn.Publish(Subject, []byte("Hello World"))
	if err2 != nil {
		log.Fatal(err2)
	}
	fmt.Println("Message published")

	conn.Flush()
	fmt.Println("Done")
}
