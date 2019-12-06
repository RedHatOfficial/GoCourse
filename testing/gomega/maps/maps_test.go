package main

import (
	. "github.com/onsi/gomega"
	"testing"
)

func TestCollecions(t *testing.T) {
	g := NewGomegaWithT(t)

	m := make(map[string]int)
	m["one"] = 1
	m["two"] = 2
	m["three"] = 3

	g.Expect(m).To(HaveLen(3))
	g.Expect(m).To(ContainElement(1))
	g.Expect(m).To(Not(ContainElement(42)))
	g.Expect(m).To(HaveKey("two"))
	g.Expect(m).To(Not(HaveKey("forty two")))
}
