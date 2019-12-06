package main

import (
	"regexp"
	"testing"
	"time"

	"github.com/google/goexpect"
)

func TestPythonInterpreter(t *testing.T) {
	child, _, err := expect.Spawn("python", 2*time.Second)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("Python interpreter has been started")

	defer child.Close()

	_, m, err := child.Expect(regexp.MustCompile("Python ([23])"), 2*time.Second)
	if err != nil {
		log.Fatal(err)
	}

	err = child.Send("quit()\n")
	if err != nil {
		t.Fatal(err)
	}

	if len(m) < 1 {
		t.Fatal("No match (should not happen")
	}
	version := m[1]
	t.Log("Detected Python version:", version)
}
