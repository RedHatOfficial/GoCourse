package main

import (
	"fmt"
	"github.com/Arafatk/glot"
	"math"
)

type Plot struct {
	*glot.Plot
}

func (plot *Plot) Save(filename string) {
	plot.Cmd("set terminal pngcairo")
	plot.Cmd("set output '" + filename + "'")
	plot.Cmd("replot")
}

func NewPlot(dimensions int) *Plot {
	plot, err := glot.NewPlot(dimensions, false, false)
	if err != nil {
		panic(err)
	}
	return &Plot{plot}

}

const points = 100

func main() {
	plot := NewPlot(2)
	defer plot.Close()

	pts := make([][]float64, 2)
	for i := 0; i < 2; i++ {
		pts[i] = make([]float64, points)
	}
	for i := 0; i < points; i++ {
		x := float64(i) * 2.0 * math.Pi / points
		pts[0][i] = x
		pts[1][i] = math.Sin(x)
	}
	plot.AddPointGroup("sin t", "lines", pts)

	plot.SetTitle("Plot #8")
	plot.SetXLabel("t")
	plot.SetYLabel("sin t")

	plot.Save("glot08.png")
}
