package main

import (
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

const points = 400

func main() {
	plot := NewPlot(3)
	defer plot.Close()

	var pointsX [points]float64
	var pointsY [points]float64
	for i := 0; i < points; i++ {
		t := float64(i) * 8.0 * math.Pi / points
		pointsX[i] = math.Sin(t)
		pointsY[i] = math.Cos(t)
	}

	z := 0.0
	function1 := func(u, v float64) float64 {
		z = z + 1.0
		return z
	}

	plot.AddFunc3d("spiral", "points", pointsX[:], pointsY[:], function1)

	plot.SetTitle("Plot #15")
	plot.SetXLabel("t")
	plot.SetYLabel("y")

	plot.Save("glot15.png")
}
