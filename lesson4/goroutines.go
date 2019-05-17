package main

import (
	"fmt"
	"time"
)

func f(s string) {
	for i := 0; i < 5; i++ {
		fmt.Println(s)
		time.Sleep(500 * time.Millisecond)
	}
}

func main() {
	go f("first")
	go f("second")
	f("main")
}
