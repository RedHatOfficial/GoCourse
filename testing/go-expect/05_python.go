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

	command := exec.Command("python")
	command.Stdin = console.Tty()
	command.Stdout = console.Tty()
	command.Stderr = console.Tty()

	err = command.Start()
	if err != nil {
		log.Fatal(err)
	}

	time.Sleep(time.Second)
	console.ExpectString(">>> ")

	console.SendLine("1+2")
	console.ExpectString("3")
	console.ExpectString(">>> ")

	console.SendLine("6*7")
	console.ExpectString("42")
	console.ExpectString(">>> ")

	console.SendLine("quit()")

	err = command.Wait()
	if err != nil {
		log.Fatal(err)
	}
}
