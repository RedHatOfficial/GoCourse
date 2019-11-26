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

	br, err := child.ExpectBatch([]expect.Batcher{
		&expect.BCas{[]expect.Caser{
			&expect.Case{R: regexp.MustCompile("Python 2"), T: expect.OK()},
			&expect.Case{R: regexp.MustCompile("Python 3"), T: expect.OK()}}},
		&expect.BExp{R: ">>> "},
		&expect.BSnd{S: "6*7\n"},
		&expect.BExp{R: "42"},
		&expect.BExp{R: ">>> "},
		&expect.BSnd{S: "quit()\n"}},
		2*time.Second)

	log.Println("OK")
	for _, b := range br {
		log.Println(b.Output)
	}
}
