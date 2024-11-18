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
$ kfk topics --create \
--topic orders \
--partitions 1 --replication-factor 1 \
-c kafka-cluster -n default
// END OMIT
*/
func main() {
	cmd := exec.Command("kfk", "topics", "--create",
		"--topic", "orders",
		"--partitions", "1", "--replication-factor", "1",
		"-c", "kafka-cluster", "-n", "default")

	var stdBuffer bytes.Buffer
	mw := io.MultiWriter(os.Stdout, &stdBuffer)

	cmd.Stdout = mw
	cmd.Stderr = mw

	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}

	//log.Println(stdBuffer.String())

}
