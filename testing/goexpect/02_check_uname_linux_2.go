package main

import (
	"log"
	"regexp"
	"time"

	"github.com/google/goexpect"
)

func main() {
	child, _, err := expect.Spawn("uname", -1)
	if err != nil {
		log.Fatal(err)
	}

	defer child.Close()

	linuxRe := regexp.MustCompile("Linux")
	s, match, err := child.Expect(linuxRe, time.Second)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Found: %s", s)
	log.Printf("Matches: %v", match)
}
