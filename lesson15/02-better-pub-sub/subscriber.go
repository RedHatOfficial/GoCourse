package main

import (
	"fmt"
	"log"
	"sync"

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

	wg := sync.WaitGroup{}
	wg.Add(1)

	sub, err2 := conn.Subscribe(Subject, func(m *nats.Msg) {
		fmt.Printf("Received a message: %s\n", string(m.Data))
		wg.Done()
	})

	if err2 != nil {
		log.Fatal(err2)
	}

	fmt.Println("Subscribed", sub)

	wg.Wait()
	fmt.Println("Finished waiting for message")

	err3 := sub.Unsubscribe()
	if err3 != nil {
		log.Fatal(err3)
	}

	fmt.Println("Unsubscribed")
}
