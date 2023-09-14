package main

import (
	"fmt"
	"time"
)

func worker(data, quit, ack chan int) {
	for {
		select {
		case x := <-data:
			fmt.Println("start working on", x)
			time.Sleep(1 * time.Second)
			fmt.Println("finish working on", x)
		case <-quit:
			fmt.Println("finishing...")
			time.Sleep(4 * time.Second)
			fmt.Println("...done")
			ack <- 1
			return
		}
	}
}

func main() {
	data := make(chan int)
	quit := make(chan int)
	ack := make(chan int)

	go worker(data, quit, ack)

	data <- 1
	quit <- 1
	<-ack
	fmt.Println("All done")
}
