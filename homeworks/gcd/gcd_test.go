package main

import "testing"

type gcdEntry struct {
	x        uint
	y        uint
	expected uint
}

func TestGCD(t *testing.T) {
	var entries = []gcdEntry{
		{1, 1, 1},
		{12, 8, 4},
		{1000, 1, 1},
		{1000, 500, 500},
		{1224, 992, 8},
		{112909, 115249, 1}, // primes
	}
	for _, entry := range entries {
		computed := gcd(entry.x, entry.y)
		if computed != entry.expected {
			t.Errorf("GCD(%d, %d) != %d, but %d", entry.x, entry.y, computed, entry.expected)
		}
		t.Logf("GCD ok for input %d and %d", entry.x, entry.y)
	}
}
