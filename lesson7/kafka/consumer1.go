package main

import (
	"fmt"
	"gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"
)

const (
	server   = "localhost:9092"
	topic    = "upload"
	group_id = "group1"
)

func main() {
	consumer, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": server,
		"group.id":          group_id,
		"auto.offset.reset": "earliest",
	})
	defer consumer.Close()

	if err != nil {
		panic(err)
	}

	consumer.SubscribeTopics([]string{topic}, nil)

	for {
		message, err := consumer.ReadMessage(-1)
		if err == nil {
			fmt.Printf("Message on %s: %s %s\n", message.TopicPartition, string(message.Key), string(message.Value))
		} else {
			fmt.Printf("Consumer error: %v (%v)\n", err, message)
		}
	}
}
