package main

import (
	. "github.com/franela/goblin"
	"testing"
)

func Test(t *testing.T) {
	g := Goblin(t)
	g.Describe("Adder", func() {
		g.It("Should add two numbers ", func() {
			g.Assert(1 + 1).Equal(2)
		})
		g.It("Should add two numbers", func() {
			g.Assert(1 + 1).Equal(5)
		})
		g.It("Should substract two numbers")
		g.Xit("Should add two numbers, excluded ", func() {
			g.Assert(3 + 1).Equal(4)
		})
	})
}
