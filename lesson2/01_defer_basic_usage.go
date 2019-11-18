package main

import "fmt"

func onFinish() {
	fmt.Println("Finished")
}

func main() {
	defer onFinish()

	for i := 10; i >= 0; i-- {
		fmt.Printf("%2d\n", i)
	}
	fmt.Println("Finishing main() function")
}
