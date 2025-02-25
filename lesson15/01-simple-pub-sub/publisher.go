package main

import (
	"fmt"

	nats "github.com/nats-io/nats.go"
)

const Subject = "test1"

func main() {
	conn, _ := nats.Connect(nats.DefaultURL)
	fmt.Printf("Connected to %s\n", nats.DefaultURL)

	conn.Publish(Subject, []byte("Hello World"))
	fmt.Println("Message published")

	conn.Flush()
	fmt.Println("Done")
}
