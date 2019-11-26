package main

import (
	"log"
	"regexp"
	"time"

	"github.com/google/goexpect"
)

func main() {
	child, errChannel, err := expect.Spawn("python", 2*time.Second)
	if err != nil {
		log.Fatal(err)
	}

	defer child.Close()

	_, err = child.ExpectBatch([]expect.Batcher{
		&expect.BCas{[]expect.Caser{
			&expect.Case{R: regexp.MustCompile("Python 2"), T: expect.OK()},
			&expect.Case{R: regexp.MustCompile("Python 3"), T: expect.OK()}}},
		&expect.BExp{R: ">>> "},
		&expect.BSnd{S: "import sys\n"},
		&expect.BExp{R: ">>> "},
		&expect.BSnd{S: "sys.exit(1)\n"}},
		2*time.Second)

	log.Println("OK")

	err = <-errChannel
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Exit: success")
}
