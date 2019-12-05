package main

import (
	. "github.com/onsi/gomega"
	"testing"
)

func TestFactorialForZero(t *testing.T) {
	RegisterTestingT(t)
	Ω(Factorial(0)).To(Equal(int64(1)))
}

func TestFactorialForOne(t *testing.T) {
	RegisterTestingT(t)
	Ω(Factorial(1)).To(Equal(int64(1)))
}

func TestFactorialForTen(t *testing.T) {
	RegisterTestingT(t)
	Ω(Factorial(10)).To(Equal(int64(3628800)))
}
