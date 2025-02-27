package main

import (
	"context"
	"fmt"
	"time"

	"github.com/nats-io/nats.go"
	"github.com/nats-io/nats.go/jetstream"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	conn, _ := nats.Connect(nats.DefaultURL)
	fmt.Printf("Connected to %s\n", nats.DefaultURL)

	js, _ := jetstream.New(conn)

	stream, _ := js.CreateStream(ctx, jetstream.StreamConfig{
		Name:     "TEST",
		Subjects: []string{"TEST.*"},
	})

	consumer, _ := stream.CreateOrUpdateConsumer(ctx, jetstream.ConsumerConfig{
		Durable:   "CONS",
		AckPolicy: jetstream.AckExplicitPolicy,
	})

	it, _ := consumer.Messages()
	for {
		msg, _ := it.Next()
		msg.Ack()
		fmt.Printf("Received message: %s\n", string(msg.Data()))
	}
	it.Stop()
}
