package main

import (
	"context"
	"fmt"
	"log"
)

func main() {
	// START OMIT
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{"localhost:9092"},
		Topic:   "test-topic",
		GroupID: "test-group",
	})
	defer reader.Close()

	for {
		msg, err := reader.ReadMessage(context.Background())
		if err != nil {
			log.Fatalln("Failed to consume message:", err)
		}
		fmt.Printf("Consumed message: %s\n", string(msg.Value))
	}
	// END OMIT
}
