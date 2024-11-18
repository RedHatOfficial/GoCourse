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
$ open ~/Applications/GoLand.app
// END OMIT
*/
func main() {
	cmd := exec.Command("/usr/bin/open", "/Users/mabulgu/Applications/GoLand.app")

	var stdBuffer bytes.Buffer
	mw := io.MultiWriter(os.Stdout, &stdBuffer)

	cmd.Stdout = mw
	cmd.Stderr = mw

	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}

	//log.Println(stdBuffer.String())

}
