package main

import (
	"fmt"
	"time"
)

func f(c chan int) {
	if i, ok := <-c; ok {
		fmt.Println("Read number:", i)
	} else {
		fmt.Println("Channel closed")
	}
}

func main() {
	c := make(chan int)
	go f(c)
	close(c)
	time.Sleep(time.Second)
}
