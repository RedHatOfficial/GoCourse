package main

import (
	. "github.com/franela/goblin"
	"testing"
)

func TestAdder(t *testing.T) {
	g := Goblin(t)
	g.Describe("Adder", func() {
		g.Describe("Positive numbers", func() {
			g.It("Should add two numbers ", func() {
				g.Assert(1 + 1).Equal(2)
			})
			g.It("Should add two numbers", func() {
				g.Assert(2 + 2).Equal(4)
			})
			g.It("Should add two numbers", func() {
				g.Assert(10 + 20).Equal(30)
			})
		})
		g.Describe("Negative numbers", func() {
			g.It("Should add two numbers ", func() {
				g.Assert(1 + -1).Equal(0)
			})
			g.It("Should add two numbers", func() {
				g.Assert(2 + -4).Equal(-2)
			})
			g.It("Should add two numbers", func() {
				g.Assert(10 + -20).Equal(-10)
			})
		})
	})
}

func TestMultiplier(t *testing.T) {
	g := Goblin(t)
	g.Describe("Multiplier", func() {
		g.Describe("Positive numbers", func() {
			g.It("Should multiply two numbers ", func() {
				g.Assert(1 * 1).Equal(1)
			})
			g.It("Should multiply two numbers", func() {
				g.Assert(2 * 2).Equal(4)
			})
			g.It("Should multiply two numbers", func() {
				g.Assert(10 * 20).Equal(200)
			})
		})
		g.Describe("Negative numbers", func() {
			g.It("Should multiply two numbers ", func() {
				g.Assert(1 * -1).Equal(-1)
			})
			g.It("Should multiply two numbers", func() {
				g.Assert(2 * -4).Equal(-8)
			})
			g.It("Should multiply two numbers", func() {
				g.Assert(10 * -20).Equal(-200)
			})
		})
	})
}
