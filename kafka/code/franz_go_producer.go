package main

import (
	"context"
	"log"

	"github.com/twmb/franz-go/pkg/kgo"
)

/*
// START1 OMIT
$ go get github.com/twmb/franz-go/pkg/kgo
// END1 OMIT
*/
func main() {
	// START2 OMIT
	client, err := kgo.NewClient(
		kgo.SeedBrokers("localhost:9092"),
	)
	if err != nil {
		log.Fatalln("Failed to create Franz-go client:", err)
	}
	defer client.Close()

	record := &kgo.Record{Topic: "test-topic", Value: []byte("Hello, Kafka from Franz-go!")}
	client.Produce(context.Background(), record, func(r *kgo.Record, err error) {
		if err != nil {
			log.Printf("Failed to produce message: %v\n", err)
		} else {
			log.Println("Message produced successfully")
		}
	})
	// END2 OMIT
}
