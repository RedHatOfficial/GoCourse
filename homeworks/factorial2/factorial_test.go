package main

import (
	"testing"
)

type factorialEntry struct {
	n          int64
	expected   int64
	shouldFail bool
}

func TestFactorial(t *testing.T) {
	var entries = []factorialEntry{
		{0, 1, false},
		{1, 1, false},
		{2, 2, false},
		{3, 6, false},
		{9, 362880, false},
		{10, 3628800, false},
		{20, 2432902008176640000, false},
		{-1, 0, true},
		{-100, 99, true},
	}
	for _, entry := range entries {
		computed, err := factorial(entry.n)
		switch {
		case err != nil:
			if entry.shouldFail {
				t.Logf("invalid input detected correctly: %d", entry.n)
			} else {
				t.Errorf("error returned for valid input: %d", entry.n)
			}
		case computed != entry.expected:
			t.Errorf("%d! != %d, but %d", entry.n, computed, entry.expected)
		default:
			t.Logf("factorial computer correctly for input %d", entry.n)
		}
	}
}
