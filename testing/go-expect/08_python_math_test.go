package main

import (
	"fmt"
	"os"
	"os/exec"
	"testing"
	"time"

	expect "github.com/Netflix/go-expect"
)

func TestPythonInterpreter(t *testing.T) {
	console, err := expect.NewConsole(expect.WithStdout(os.Stdout))
	if err != nil {
		t.Fatal(err)
	}
	defer console.Close()
	t.Log("Console created")

	command := exec.Command("python")
	command.Stdin = console.Tty()
	command.Stdout = console.Tty()
	command.Stderr = console.Tty()

	err = command.Start()
	if err != nil {
		t.Fatal(err)
	}

	t.Log("Python interpreter started")
	time.Sleep(time.Second)

	str, err := console.Expect(expect.String("Python 2", "Python 3"), expect.WithTimeout(100*time.Millisecond))
	if err != nil {
		t.Fatal("Python not detected")
	}
	t.Log("Python interpreter detected: " + str)

	for i := uint(1); i < 10; i++ {
		console.SendLine(fmt.Sprintf("2**%d", i))
		_, err = console.Expectf("%d", 1<<i)
		if err != nil {
			t.Fatal("Math is wrong!")
		}
		t.Logf("Math is ok for input %d", i)
	}

	console.SendLine("quit()")

	err = command.Wait()
	if err != nil {
		t.Fatal(err)
	}
	t.Log("Done")
}
