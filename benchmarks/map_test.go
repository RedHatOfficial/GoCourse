package main

import (
	"testing"
)

type key struct {
	ID      int
	payload [100]byte
}

type value struct{}

func BenchmarkInsertIntoPreallocatedMap(b *testing.B) {
	m := make(map[key]value, b.N)

	for i := 0; i < b.N; i++ {
		k := key{
			ID: i,
		}
		m[k] = value{}
	}
}

func BenchmarkInsertIntoEmptyMap(b *testing.B) {
	m := map[key]value{}

	for i := 0; i < b.N; i++ {
		k := key{
			ID: i,
		}
		m[k] = value{}
	}
}
