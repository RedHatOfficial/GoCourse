package main

import (
	"fmt"
	"log"

	"github.com/nats-io/nats.go"
)

const StreamName = "TEST"
const StreamSubjects = "TEST.*"

func main() {
	conn, _ := nats.Connect(nats.DefaultURL)
	fmt.Printf("Connected to %s\n", nats.DefaultURL)

	js, err := conn.JetStream(nats.PublishAsyncMaxPending(256))
	if err != nil {
		log.Fatal(err)
	}

	stream, err := js.StreamInfo(StreamName)

	if stream == nil {
		fmt.Printf("Creating stream: %s\n", StreamName)

		_, err = js.AddStream(&nats.StreamConfig{
			Name:     StreamName,
			Subjects: []string{StreamSubjects},
		})
		if err != nil {
			log.Fatal(err)
		}
	}

	_, err = js.Publish("TEST.xyz", []byte("foo bar baz"))
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Printf("Published")
	}
	conn.Close()

}

func CreateStream(jetStream nats.JetStreamContext) error {
	return nil
}
