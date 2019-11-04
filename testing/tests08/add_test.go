package main

import "testing"

func TestAdd(t *testing.T) {
	result := add(1, 2)
	if result != 3 {
		t.Error("1+2 should be 3, got ", result, "instead")
	}
}
