package main

import "os"

func main() {
	file, _ := os.Open("file.txt")
	file.Close()
}
