package main

import "fmt"

func onFinish(i int) {
	fmt.Printf("Defer #%2d\n", i)
}

func main() {
	for i := 0; i <= 10; i++ {
		defer onFinish(i)
	}
	fmt.Println("Finishing main() function")
}
