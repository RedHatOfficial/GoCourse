package main

import (
	. "github.com/onsi/gomega"
	"testing"
)

func TestFactorialForZero(t *testing.T) {
	g := NewGomegaWithT(t)
	g.Expect(Factorial(0)).To(BeNumerically("==", 1))
}

func TestFactorialForOne(t *testing.T) {
	g := NewGomegaWithT(t)
	g.Expect(Factorial(1)).To(BeNumerically("==", 1))
}

func TestFactorialForTen(t *testing.T) {
	g := NewGomegaWithT(t)
	g.Expect(Factorial(10)).To(BeNumerically("==", 3628800))
}

func TestFactorialForSmallNumber(t *testing.T) {
	g := NewGomegaWithT(t)
	g.Expect(Factorial(5)).To(SatisfyAll(BeNumerically(">=", 10), BeNumerically("<=", 10000)))
}
