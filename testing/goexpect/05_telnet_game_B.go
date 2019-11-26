package main

import (
	"log"
	"time"

	"github.com/google/goexpect"
)

func main() {
	child, _, err := expect.Spawn("telnet zombiemud.org", 1*time.Second)
	if err != nil {
		log.Fatal(err)
	}

	defer child.Close()

	_, err = child.ExpectBatch([]expect.Batcher{
		&expect.BExp{R: "Your choice or name:"},
		&expect.BSnd{S: "d\n"},
		&expect.BExp{R: "Ok, see you later!"}}, 1*time.Second)

	if err != nil {
		log.Fatal(err)
	}
	log.Println("OK")
}
