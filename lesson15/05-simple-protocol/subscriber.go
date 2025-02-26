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

	econn, err2 := nats.NewEncodedConn(conn, nats.DEFAULT_ENCODER)

	if err2 != nil {
		log.Fatal(err)
	}

	defer econn.Close()

	channel := make(chan string)
	econn.BindRecvChan(Subject, channel)

	fmt.Println("Channel created")

	fmt.Println("Waiting for messages")
	for {
		message := <-channel
		fmt.Println(message)
		if message == "EXIT" {
			fmt.Println("Received EXIT message...")
			break
		}
	}
	fmt.Println("Done")
}
