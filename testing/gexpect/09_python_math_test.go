package main

import (
	"fmt"
	"testing"
	"time"

	"github.com/ThomasRooney/gexpect"
)

func expectOutput(t *testing.T, child *gexpect.ExpectSubprocess, output string) {
	err := child.ExpectTimeout(output, time.Second)
	if err != nil {
		t.Fatal(err)
	}
}

func expectPrompt(t *testing.T, child *gexpect.ExpectSubprocess) {
	expectOutput(t, child, ">>> ")
}

func sendCommand(t *testing.T, child *gexpect.ExpectSubprocess, command string) {
	err := child.SendLine(command)
	if err != nil {
		t.Fatal(err)
	}
}

func TestPythonInterpreter(t *testing.T) {
	child, err := gexpect.Spawn("python")
	if err != nil {
		t.Fatal(err)
	}
	t.Log("Python interpreter started")

	strs, err := child.ExpectRegexFind("Python [23]")
	if err != nil {
		t.Fatal("Python not detected")
	}
	t.Log("Python interpreter detected: " + strs[0])

	for i := uint(1); i < 10; i++ {
		sendCommand(t, child, fmt.Sprintf("2**%d", i))
		expectOutput(t, child, fmt.Sprintf("%d", 1<<i))
		t.Logf("Math is ok for input %d", i)
	}

	expectPrompt(t, child)
	sendCommand(t, child, "quit()")

	child.Wait()
}
