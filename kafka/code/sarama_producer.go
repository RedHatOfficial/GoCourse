package main

import (
	"fmt"
	"log"

	"github.com/Shopify/sarama"
)

/*
// START1 OMIT
$ go get github.com/Shopify/sarama
// END1 OMIT
*/
func main() {
	// START2 OMIT
	producer, err := sarama.NewSyncProducer([]string{"localhost:9092"}, nil)
	if err != nil {
		log.Fatalln("Failed to start Sarama producer:", err)
	}
	defer producer.Close()

	message := &sarama.ProducerMessage{
		Topic: "test-topic",
		Value: sarama.StringEncoder("Hello, Kafka from Sarama!"),
	}

	_, _, err = producer.SendMessage(message)
	if err != nil {
		log.Fatalln("Failed to produce message:", err)
	}

	fmt.Println("Message produced successfully")
	// END2 OMIT
}
