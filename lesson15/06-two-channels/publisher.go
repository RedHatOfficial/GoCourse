package main

import (
	"fmt"
	"log"
	"time"

	nats "github.com/nats-io/nats.go"
)

const Subject = "test1"
const Control = "test2"

func main() {
	conn, err := nats.Connect(nats.DefaultURL)

	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()

	fmt.Println("Connected")

	econn, err2 := nats.NewEncodedConn(conn, nats.DEFAULT_ENCODER)

	if err2 != nil {
		log.Fatal(err2)
	}

	defer econn.Close()

	cconn, err3 := nats.NewEncodedConn(conn, nats.DEFAULT_ENCODER)

	if err3 != nil {
		log.Fatal(err3)
	}

	defer cconn.Close()

	data_channel := make(chan string)
	econn.BindSendChan(Subject, data_channel)

	fmt.Println("Data channel created")

	control_channel := make(chan string)
	cconn.BindSendChan(Control, control_channel)

	fmt.Println("Control channel created")

	data_channel <- "Hello World #1"
	data_channel <- "Hello World #2"
	data_channel <- "Hello World #3"
	data_channel <- "EXIT"
	time.Sleep(2 * time.Second)
	control_channel <- "EXIT"

	conn.Flush()

	fmt.Println("All messages sent")
}
