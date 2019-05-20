package main

import "fmt"

func f(c chan int) {
	a := <-c
	fmt.Println("Received", a)
}

func main() {
	c := make(chan int)
	go f(c)
	c <- 5
}
