package main

import (
	"github.com/Arafatk/glot"
)

func main() {
	plot, err := glot.NewPlot(2, false, false)
	if err != nil {
		panic(err)
	}

	plot.AddPointGroup("Measured data", "points", []int32{1, 2, 4, 8, 9, 8, 4, 2, 1})
	plot.SavePlot("glot01.png")
}
