package main

import (
	"fmt"
	"log"

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
	econn.BindRecvChan(Subject, data_channel)

	fmt.Println("Data channel created")

	control_channel := make(chan string)
	cconn.BindRecvChan(Control, control_channel)

	fmt.Println("Control channel created")

MESSAGE_LOOP:
	for {
		select {
		case message := <-data_channel:
			fmt.Println("Received data message", message)
		case control := <-control_channel:
			fmt.Println("Received control message", control)
			if control == "EXIT" {
				break MESSAGE_LOOP
			}
		}
		fmt.Println("--------")
	}
}
