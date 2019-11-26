package main

import (
	"log"
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
		&expect.BExp{R: ">>> "},
		&expect.BSnd{S: "1+2\n"},
		&expect.BExp{R: "3"},
		&expect.BExp{R: ">>> "},
		&expect.BSnd{S: "6*7\n"},
		&expect.BExp{R: "42"},
		&expect.BExp{R: ">>> "},
		&expect.BSnd{S: "quit()\n"}},
		2*time.Second)

	if err != nil {
		log.Fatal(err)
	}
	log.Println("OK")
	for _, b := range br {
		log.Println(b.Output)
	}
}
