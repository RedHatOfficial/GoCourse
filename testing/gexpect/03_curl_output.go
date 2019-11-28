package main

import (
	"log"

	"github.com/ThomasRooney/gexpect"
)

func main() {
	child, err := gexpect.Spawn("curl -X HEAD -v github.com")
	if err != nil {
		log.Fatal(err)
	}
	err = child.Expect("Location: https://github.com/")
	if err != nil {
		log.Fatal(err)
	}
	child.Wait()
}
