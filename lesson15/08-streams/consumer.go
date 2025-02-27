package main

import (
	"fmt"
	"log"

	"github.com/nats-io/nats.go"
)

func main() {
	conn, _ := nats.Connect(nats.DefaultURL)
	fmt.Printf("Connected to %s\n", nats.DefaultURL)

	js, err := conn.JetStream(nats.PublishAsyncMaxPending(256))
	if err != nil {
		log.Fatal(err)
	}

	_, err = js.Subscribe("TEST.xyz", func(m *nats.Msg) {
		err := m.Ack()

		if err != nil {
			log.Println("Unable to Ack", err)
			return
		}

		message := string(m.Data)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("Consumed %s\n", message)
	})

	if err != nil {
		log.Println("Subscribe failed")
		return
	}
	conn.Close()
}
