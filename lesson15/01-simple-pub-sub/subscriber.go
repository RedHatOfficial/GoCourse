package main

import (
	"fmt"
	"sync"

	nats "github.com/nats-io/nats.go"
)

const Subject = "test1"

func main() {
	conn, _ := nats.Connect(nats.DefaultURL)
	fmt.Printf("Connected to %s\n", nats.DefaultURL)

	wg := sync.WaitGroup{}
	wg.Add(1)

	conn.Subscribe(Subject, func(m *nats.Msg) {
		fmt.Printf("Received a message: %s\n", string(m.Data))
		wg.Done()
	})
	wg.Wait()
	fmt.Println("Finished waiting for message")
}
