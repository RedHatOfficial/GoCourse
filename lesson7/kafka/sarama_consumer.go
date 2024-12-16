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
	consumer, err := sarama.NewConsumer([]string{KafkaConnectionString}, nil)

	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Connected to %s", KafkaConnectionString)

	defer func() {
		if err := consumer.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	partitionConsumer, err := consumer.ConsumePartition(KafkaTopic, 0, sarama.OffsetNewest)
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		if err := partitionConsumer.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	consumed := 0
	for {
		msg := <-partitionConsumer.Messages()
		log.Printf("Consumed message offset %d: %s:%s", msg.Offset, msg.Key, msg.Value)
		consumed++
	}

	log.Printf("Consumed: %d", consumed)
	log.Print("Done")
}
