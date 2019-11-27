package main

import (
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

	if str == "Python 2" {
		console.SendLine("print 1,2,3")
		_, err = console.ExpectString("1 2 3")
		if err != nil {
			t.Fatal("print statement failure")
		}
		t.Log("print statement works as expected")
		_, err = console.ExpectString(">>> ")
		if err != nil {
			t.Fatal("prompt is not displayed")
		}
	} else {
		console.SendLine("print(1,2,3)")
		_, err = console.ExpectString("1 2 3")
		if err != nil {
			t.Fatal("print function failure")
		}
		t.Log("print function works as expected")
		_, err = console.ExpectString(">>> ")
		if err != nil {
			t.Fatal("prompt is not displayed")
		}
	}

	console.SendLine("quit()")

	err = command.Wait()
	if err != nil {
		t.Fatal(err)
	}
	t.Log("Done")
}
