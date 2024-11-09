package main

import (
	"fmt"
	"log"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

/*
// START1 OMIT
$ go get github.com/confluentinc/confluent-kafka-go/kafka
// END1 OMIT
*/
func main() {
	// START2 OMIT
	producer, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": "localhost:9092"})
	if err != nil {
		log.Fatalln("Failed to create producer:", err)
	}
	defer producer.Close()

	message := &kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &"test-topic", Partition: kafka.PartitionAny},
		Value:          []byte("Hello, Kafka from Confluent Kafka Go!"),
	}

	producer.Produce(message, nil)
	producer.Flush(1000)

	fmt.Println("Message produced successfully")
	// END2 OMIT
}
