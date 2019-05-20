package main

import "fmt"

func f(c chan int) {
	a := <-c
	fmt.Println("Received", a)
}

func main() {
	c := make(chan int)
	c <- 5
	go f(c)
}
