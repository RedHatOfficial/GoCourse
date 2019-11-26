package main

import (
	"log"
	"regexp"
	"time"

	"github.com/google/goexpect"
)

func main() {
	child, _, err := expect.Spawn("python", 2*time.Second)
	if err != nil {
		log.Fatal(err)
	}

	defer child.Close()

	_, m, err := child.Expect(regexp.MustCompile("Python ([23])"), 2*time.Second)

	err = child.Send("quit()\n")
	if err != nil {
		log.Fatal(err)
	}
	version := m[1]
	log.Println("Python version:", version)
}
