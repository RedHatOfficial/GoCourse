package main

import (
	"bytes"
	"io"
	"log"
	"os"
	"os/exec"
)

/*
// START OMIT
$ kfk clusters --describe \
--cluster kafka-cluster \
--namespace default
// END OMIT
*/
func main() {
	cmd := exec.Command("kfk", "clusters", "--describe",
		"--cluster", "kafka-cluster",
		"--namespace", "default")

	var stdBuffer bytes.Buffer
	mw := io.MultiWriter(os.Stdout, &stdBuffer)

	cmd.Stdout = mw
	cmd.Stderr = mw

	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}

	//log.Println(stdBuffer.String())
}
