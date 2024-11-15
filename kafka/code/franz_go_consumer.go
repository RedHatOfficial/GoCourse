package main

import (
	"context"
	"fmt"
	"log"

	"github.com/twmb/franz-go/pkg/kgo"
)

func main() {
	// START OMIT
	client, err := kgo.NewClient(
		kgo.SeedBrokers("localhost:9092"),
		kgo.ConsumerGroup("test-group"),
		kgo.ConsumeTopics("test-topic"),
	)
	if err != nil {
		log.Fatalln("Failed to create Franz-go client:", err)
	}
	defer client.Close()

	for {
		fetches := client.PollFetches(context.Background())
		fetches.EachRecord(func(record *kgo.Record) {
			fmt.Printf("Consumed message: %s\n", string(record.Value))
		})
	}
	// END OMIT
}
