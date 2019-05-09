package main

import "fmt"

func on_finish() {
	fmt.Println("Finished")
}

func main() {
	defer on_finish()

	for i := 10; i >= 0; i-- {
		fmt.Printf("%2d\n", i)
	}
	fmt.Println("Finishing main() function")
}
