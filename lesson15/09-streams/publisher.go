package main

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/nats-io/nats.go"
	"github.com/nats-io/nats.go/jetstream"
)

const StreamName = "TEST"
const StreamSubjects = "TEST.*"

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	conn, _ := nats.Connect(nats.DefaultURL)
	fmt.Printf("Connected to %s\n", nats.DefaultURL)

	js, _ := jetstream.New(conn)

	for i := 0; i < 100; i++ {
		js.Publish(ctx, "TEST.new", []byte("hello message "+strconv.Itoa(i)))
		fmt.Printf("Published message #%d\n", i)
		time.Sleep(1 * time.Second)
	}

	conn.Close()
}
