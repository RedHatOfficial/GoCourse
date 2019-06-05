package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int)
	go func() {
		for i := range c {
			fmt.Println(i)
		}
		fmt.Println("Channel closed, f() over")
	}()
	for i := 0; i < 5; i++ {
		c <- i*2
	}
	close(c)
	time.Sleep(time.Second)
}
