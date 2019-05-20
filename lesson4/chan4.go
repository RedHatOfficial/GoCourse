package main

import (
	"fmt"
	"time"
)


func f(c chan int) {
	a := <-c
	fmt.Println("Received", a)
}

func main() {
	c := make(chan int, 1)
	go f(c)
	c <- 5
	time.Sleep(time.Second)
}
