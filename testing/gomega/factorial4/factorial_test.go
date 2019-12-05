package main

import (
	. "github.com/onsi/gomega"
	"testing"
)

func TestFactorialForNegativeValue(t *testing.T) {
	g := NewGomegaWithT(t)
	_, err := Factorial(-1)
	g.Expect(err).Should(HaveOccurred())
}

func TestFactorialForZero(t *testing.T) {
	g := NewGomegaWithT(t)
	result, err := Factorial(0)
	g.Expect(err).ShouldNot(HaveOccurred())
	g.Expect(result).To(BeNumerically("==", 1))
}

func TestFactorialForOne(t *testing.T) {
	g := NewGomegaWithT(t)
	result, err := Factorial(0)
	g.Expect(err).Should(Succeed())
	g.Expect(result).To(BeNumerically("==", 1))
}

func TestFactorialForTen(t *testing.T) {
	g := NewGomegaWithT(t)
	g.Expect(Factorial(10)).To(BeNumerically("==", 3628800))
}

func TestFactorialForSmallNumber(t *testing.T) {
	g := NewGomegaWithT(t)
	g.Expect(Factorial(5)).To(SatisfyAll(BeNumerically(">=", 10), BeNumerically("<=", 10000)))
}
