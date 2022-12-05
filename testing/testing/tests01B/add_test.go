package main_test

import "tests01"
import "testing"

func TestAdd(t *testing.T) {
	result := main.Add(1, 2)
	if result != 3 {
		t.Error("1+2 should be 3, got ", result, "instead")
	}
}
