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

	fmt.Println("Connected")

	econn, err2 := nats.NewEncodedConn(conn, nats.DEFAULT_ENCODER)

	if err2 != nil {
		log.Fatal(err2)
	}

	defer econn.Close()

	channel := make(chan string)
	econn.BindSendChan(Subject, channel)

	fmt.Println("Channel created")

	// poslat 100 zprav
	for i := 1; i < 100; i++ {
		message := fmt.Sprintf("Hello World #%d", i)
		channel <- message
		conn.Flush()
	}

	fmt.Println("All messages sent")
}
