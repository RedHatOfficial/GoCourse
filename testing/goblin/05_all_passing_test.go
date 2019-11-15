package main

import (
	. "github.com/franela/goblin"
	"testing"
)

func Test(t *testing.T) {
	g := Goblin(t)
	g.Describe("Adder", func() {
		g.It("Should add 1+1 ", func() {
			g.Assert(1 + 1).Equal(2)
		})
		g.It("Should add 2+2 ", func() {
			g.Assert(2 + 2).Equal(4)
		})
		g.It("Should add 10+20", func() {
			g.Assert(10 + 20).Equal(30)
		})
	})
}
