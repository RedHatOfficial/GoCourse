package main

func main() {
	array := [4]string{"one", "two", "three", "four"}

	for index, item := range array {
		println(index, item)
	}
}
