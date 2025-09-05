package main

import (
	"log"
	"os"

	"github.com/ugorji/go/codec"
)

const filename = "/tmp/nil.bin"

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

	err = encoder.Encode(nil)
	if err != nil {
		log.Panic(err)
	}

	log.Print("Done")
}
