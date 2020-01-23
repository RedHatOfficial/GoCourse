package main

// Add is exported function to be tested
func Add(x int, y int) int {
	return x + y
}

func main() {
	println(Add(1, 2))
}
