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
			&expect.Case{R: regexp.MustCompile("Python 3"), T: expect.OK()}}}}, time.Second)
	if err != nil {
		log.Fatal(err)
	}

	err = child.Send("quit()\n")
	if err != nil {
		log.Fatal(err)
	}
	log.Println("OK")
	for _, b := range br {
		log.Println(b.Output)
	}
}
