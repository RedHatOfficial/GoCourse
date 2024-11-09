package main

import (
	"fmt"
	"log"

	"github.com/Shopify/sarama"
)

func main() {
	// START OMIT
	consumer, err := sarama.NewConsumer([]string{"localhost:9092"}, nil)
	if err != nil {
		log.Fatalln("Failed to start Sarama consumer:", err)
	}
	defer consumer.Close()

	partitionConsumer, err := consumer.ConsumePartition("test-topic", 0, sarama.OffsetNewest)
	if err != nil {
		log.Fatalln("Failed to start partition consumer:", err)
	}
	defer partitionConsumer.Close()

	for msg := range partitionConsumer.Messages() {
		fmt.Printf("Consumed message: %s\n", string(msg.Value))
	}
	// END OMIT
}
