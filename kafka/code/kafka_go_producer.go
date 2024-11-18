package main

import (
	"context"
	"log"
)

/*
// START1 OMIT
$ go get github.com/segmentio/kafka-go
// END1 OMIT
*/
func main() {
	// START2 OMIT
	writer := kafka.NewWriter(kafka.WriterConfig{
		Brokers: []string{"localhost:9092"},
		Topic:   "test-topic",
	})
	defer writer.Close()

	err := writer.WriteMessages(context.Background(), kafka.Message{
		Value: []byte("Hello, Kafka from kafka-go!"),
	})
	if err != nil {
		log.Fatalln("Failed to produce message:", err)
	}

	log.Println("Message produced successfully")
	// END2 OMIT
}
