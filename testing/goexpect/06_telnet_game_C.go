package main

import (
	"log"
	"time"

	"github.com/google/goexpect"
)

func main() {
	child, _, err := expect.Spawn("telnet zombiemud.org", 2*time.Second)
	if err != nil {
		log.Fatal(err)
	}

	defer child.Close()

	br, err := child.ExpectBatch([]expect.Batcher{
		&expect.BExp{R: "Your choice or name:"},
		&expect.BSnd{S: "d\n"},
		&expect.BExp{R: "Ok, see you later!"}}, 2*time.Second)

	if err != nil {
		log.Fatal(err)
	}
	log.Println("OK")
	for _, b := range br {
		log.Println(b.Idx, b.Output)
	}
}
