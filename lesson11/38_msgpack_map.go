package main

import (
	"log"
	"os"

	"github.com/ugorji/go/codec"
)

const filename = "/tmp/map.bin"

func main() {
	fout, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0o600)
	if err != nil {
		log.Panic(err)
	}
	defer fout.Close()

	log.Print("Output file created")

	var handler codec.MsgpackHandle

	encoder := codec.NewEncoder(fout, &handler)

	log.Print("Encoder created")

	var m map[string]int = make(map[string]int)
	m["foo"] = 1
	m["bar"] = 2

	err = encoder.Encode(m)
	if err != nil {
		log.Panic(err)
	}

	log.Print("Done")
}
