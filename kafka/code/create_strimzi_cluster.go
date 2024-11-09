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
$ kfk clusters --create \
--cluster kafka-cluster \
--replicas 1 --zk-replicas 1 \
--namespace default \
--yes
// END OMIT
*/
func main() {
	cmd := exec.Command("kfk", "clusters", "--create",
		"--cluster", "kafka-cluster",
		"--replicas", "1", "--zk-replicas", "1",
		"--namespace", "default", "--yes")

	var stdBuffer bytes.Buffer
	mw := io.MultiWriter(os.Stdout, &stdBuffer)

	cmd.Stdout = mw
	cmd.Stderr = mw

	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}

	//log.Println(stdBuffer.String())
}
