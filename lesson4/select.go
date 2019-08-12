package main

import "fmt"

func main() {
	c := make(chan int)
	quit := make (chan int)
	go func() {
		for {
			select {
			case x := <-c:
				fmt.Println(x)
			case <-quit:
				quit <- 1
				return
			}
		}
	}()
	quit <- 1
	<-quit
	fmt.Println("All done")
}
