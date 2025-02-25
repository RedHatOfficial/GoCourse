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

	for i := 1; i < 10; i++ {
		message := fmt.Sprintf("Hello World #%d", i)
		err2 := conn.Publish(Subject, []byte(message))
		fmt.Println("Published", message)

		if err2 != nil {
			log.Fatal(err2)
		}

		conn.Flush()
	}

	fmt.Println("All messages sent")
}
