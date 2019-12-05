package main

import (
	. "github.com/onsi/gomega"
	"testing"
)

func TestCollecions(t *testing.T) {
	g := NewGomegaWithT(t)
	c := []int{1, 2, 3, 4, 5}
	g.Expect(c).To(HaveLen(5))
	g.Expect(c).To(HaveCap(5))
	g.Expect(c).To(ContainElement(1))
	g.Expect(c).To(Not(ContainElement(42)))
}
