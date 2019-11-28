package main

import (
	"log"

	"github.com/ThomasRooney/gexpect"
)

func main() {
	child, err := gexpect.Spawn("uname")
	if err != nil {
		log.Fatal(err)
	}
	err = child.Expect("BSD")
	if err != nil {
		log.Fatal(err)
	}
	child.Wait()
}
