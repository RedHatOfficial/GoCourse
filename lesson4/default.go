package main

import (
	"fmt"
	"time"
)

func main() {
	quit := make(chan int)
	go func() {
		for {
			select {
			case <-quit:
				return
			default:
				fmt.Println("Sleeping")
			}
			time.Sleep(time.Millisecond * 500)
		}
	}()
	time.Sleep(time.Second * 2)
	quit <- 1
}
