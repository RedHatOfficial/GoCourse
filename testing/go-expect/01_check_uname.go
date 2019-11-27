package main

import (
	"log"
	"os"
	"os/exec"
	"time"

	expect "github.com/Netflix/go-expect"
)

func main() {
	console, err := expect.NewConsole(expect.WithStdout(os.Stdout))
	if err != nil {
		log.Fatal(err)
	}
	defer console.Close()

	command := exec.Command("uname")
	command.Stdin = console.Tty()
	command.Stdout = console.Tty()
	command.Stderr = console.Tty()

	err = command.Start()
	if err != nil {
		log.Fatal(err)
	}

	time.Sleep(time.Second)
	console.ExpectString("Linux")

	err = command.Wait()
	if err != nil {
		log.Fatal(err)
	}
}
