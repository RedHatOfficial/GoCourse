package main

import (
	"log"

	"github.com/Shopify/sarama"
)

const (
	KafkaConnectionString = "localhost:9092"

	KafkaTopic = "test-topic"
)

func main() {
	producer, err := sarama.NewSyncProducer([]string{KafkaConnectionString}, nil)

	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Connected to %s", KafkaConnectionString)

	defer func() {
		if err := producer.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	msg := &sarama.ProducerMessage{Topic: KafkaTopic, Value: sarama.StringEncoder("testing 123")}
	partition, offset, err := producer.SendMessage(msg)
	if err != nil {
		log.Printf("FAILED to send message: %s\n", err)
	} else {
		log.Printf("> message sent to partition %d at offset %d\n", partition, offset)
	}

	log.Print("Done")
}
